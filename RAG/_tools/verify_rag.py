#!/usr/bin/env python3
"""Independent verifier for every RAG corpus (does not trust build_rag.py).

For each corpus in RAG/manifest.json, using only the generated files:
  1. the manifest ranges form an exact cover of the source (1..n, no gap/overlap);
  2. every chunk's source slice appears verbatim in its generated file;
  3. every source anchor appears exactly once across the corpus;
  4. manifest <-> files consistency (path exists, sha256 matches source slice).

Exit code 0 = OK, 1 = failure.
"""

from __future__ import annotations

import hashlib
import json
import re
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
RAG_ROOT = ROOT / "RAG"

fail: list[str] = []


def check(cond, msg):
    print(("  ok   " if cond else "  FAIL ") + msg)
    if not cond:
        fail.append(msg)


def verify_files(corpus: dict, man: dict, chunks: list):
    """Corpus `files` : une source = un dossier de fiches. Une fiche courte
    fait un chunk verbatim ; une fiche longue (> borne) est decoupee en
    tranches contigues couvrant exactement son fichier source."""
    slug = corpus["corpus"]
    base = ROOT / man["source_dir"]
    srcs = sorted(p for p in base.rglob("*.md") if not p.name.startswith("_"))
    print(f"\n=== {slug} === ({len(srcs)} fiches, {len(chunks)} chunks)")

    by_src = {}
    for c in chunks:
        by_src.setdefault(c["source"], []).append(c)

    print("1) chaque fiche source a au moins un chunk, aucun chunk orphelin")
    want = {str(p.relative_to(ROOT)) for p in srcs}
    check(set(by_src) == want, "chaque fiche a au moins un chunk")

    print("2) contenu verbatim + sha256 (par tranche)")
    all_ok = True
    for src, cs in by_src.items():
        lines = (ROOT / src).read_text(encoding="utf-8").splitlines(keepends=True)
        n = len(lines)
        ordered = sorted(cs, key=lambda c: c["source_lines"][0])
        cursor = 1
        for c in ordered:
            a, b = c["source_lines"]
            if a != cursor or b < a or b > n:
                all_ok = False
                check(False, f"{c['path']} range [{a},{b}] non contigu (attendu {cursor})")
            cursor = b + 1
            slice_text = "".join(lines[a - 1:b])
            content = (RAG_ROOT / c["path"]).read_text(encoding="utf-8")
            if content.count(slice_text) != 1:
                all_ok = False
                check(False, f"{c['path']} ne contient pas sa tranche verbatim")
            if hashlib.sha256(slice_text.encode("utf-8")).hexdigest() != c["sha256"]:
                all_ok = False
                check(False, f"{c['path']} sha256 mismatch")
        if cursor != n + 1:
            all_ok = False
            check(False, f"{src} non couvert jusqu'a la ligne {n} (cursor={cursor})")
    check(all_ok, "chaque fiche source couverte exactement par ses tranches")

    print("3) manifest <-> fichiers")
    check(man["chunk_count"] == len(chunks), "chunk_count consistent")
    check(man.get("source_files") == len(srcs), "source_files consistent")
    expect = hashlib.sha256("".join(
        f"{c['source']}\0{c['source_lines'][0]}\0{c['sha256']}\n"
        for c in sorted(chunks, key=lambda c: (c["source"], c["source_lines"][0]))
    ).encode("utf-8")).hexdigest()
    check(man["source_sha256"] == expect, "global source sha256 matches")
    check(all((RAG_ROOT / c["path"]).exists() for c in chunks), "every path exists")

    print("4) corpus relationship")
    check("delta_of" not in man, "no delta_of without a delta relationship")


def verify(corpus: dict):
    slug = corpus["corpus"]
    man = json.loads((RAG_ROOT / slug / "manifest.json").read_text(encoding="utf-8"))
    chunks = man["chunks"]
    if man.get("source_dir"):
        return verify_files(corpus, man, chunks)
    src_path = ROOT / corpus["source"]
    src = src_path.read_text(encoding="utf-8")
    lines = src.splitlines(keepends=True)
    n = len(lines)
    print(f"\n=== {slug} === ({n} lines, {len(chunks)} chunks)")

    print("1) exact cover")
    ordered = sorted(chunks, key=lambda c: c["source_lines"][0])
    check(ordered[0]["source_lines"][0] == 1, "starts at line 1")
    check(ordered[-1]["source_lines"][1] == n, f"ends at line {n}")
    cursor = 1
    contiguous = True
    for c in ordered:
        a, b = c["source_lines"]
        if not (a == cursor and b >= a):
            contiguous = False
            check(False, f"{c['path']} range [{a},{b}] not contiguous")
        cursor = b + 1
    check(contiguous, "all ranges contiguous")

    print("2) source slice present verbatim in file")
    anchor_count = {}
    for c in chunks:
        a, b = c["source_lines"]
        slice_text = "".join(lines[a - 1:b])
        content = (RAG_ROOT / c["path"]).read_text(encoding="utf-8")
        check(content.count(slice_text) == 1, f"{c['path']} contains its slice once")
        digest = hashlib.sha256(slice_text.encode("utf-8")).hexdigest()
        check(digest == c["sha256"], f"{c['path']} sha256 matches")
        for anc in re.findall(r'<a id="([^"]+)">', slice_text):
            anchor_count[anc] = anchor_count.get(anc, 0) + 1

    print("3) anchors unique")
    src_anchors = re.findall(r'<a id="([^"]+)">', src)
    check(len(src_anchors) == len(set(src_anchors)), f"{len(src_anchors)} source anchors unique")
    check(set(anchor_count) == set(src_anchors), "all source anchors present")
    check(all(v == 1 for v in anchor_count.values()), "each anchor found exactly once")

    print("4) manifest <-> files")
    check(man["chunk_count"] == len(chunks), "chunk_count consistent")
    check(man["source_sha256"] == hashlib.sha256(src.encode("utf-8")).hexdigest(),
          "global source sha256 matches")
    check(all((RAG_ROOT / c["path"]).exists() for c in chunks), "every path exists")

    print("5) corpus relationship")
    base = man.get("delta_of")
    if base:
        check(man.get("relationship") == "delta", f"declared delta of `{base}`")
        check((RAG_ROOT / base / "manifest.json").exists(), f"base corpus `{base}` exists")
        check(all(c.get("delta_of") == base for c in chunks),
              "every chunk carries delta_of")
        bman = json.loads((RAG_ROOT / base / "manifest.json").read_text(encoding="utf-8"))
        bh = {c["sha256"] for c in bman["chunks"]}
        dup = [c["path"] for c in chunks if c["sha256"] in bh]
        check(not dup, f"no chunk duplicates a `{base}` chunk ({len(dup)} found)")
    else:
        check("delta_of" not in man, "no delta_of without a delta relationship")


def main():
    gman = json.loads((RAG_ROOT / "manifest.json").read_text(encoding="utf-8"))
    for corpus in gman["corpora"]:
        verify(corpus)
    print()
    if fail:
        print(f"FAILED: {len(fail)} check(s)")
        sys.exit(1)
    print("ALL CHECKS PASSED")


if __name__ == "__main__":
    main()
