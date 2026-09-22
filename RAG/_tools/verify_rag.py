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


def verify(corpus: dict):
    slug = corpus["corpus"]
    src_path = ROOT / corpus["source"]
    src = src_path.read_text(encoding="utf-8")
    lines = src.splitlines(keepends=True)
    n = len(lines)
    man = json.loads((RAG_ROOT / slug / "manifest.json").read_text(encoding="utf-8"))
    chunks = man["chunks"]
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
