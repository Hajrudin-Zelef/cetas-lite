---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-13
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["decode", "embedding", "nvidia", "valuation"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [2406, 2587]
sha256: 777fff6897863291484dcd1cb97b1ca9b1abf68052b875b01ea8ccd319074cdb
---

# Outils dev + ingénierie RAG (chunk & corpus)

**Recommandation pour toi :** commence simple (sections §78). Passe au
parent-enfant si ton évaluation montre « bon chunk retrouvé mais réponse
incomplète ».

## 82. Ne JAMAIS couper un bloc de code ou un tableau en deux

C'est le piège n°1 sur la doc technique (vérifié : article dev.to 2026
« your chunker splits code fences »). Un bloc coupé = deux chunks
inutilisables.

```python
BLOC_INSECABLES = ("```",)  # fences markdown

def split_respecting_blocks(text: str, max_tokens: int) -> list[str]:
    """Découpe sur les paragraphes mais ne coupe jamais dans un fence ``` ni un tableau."""
    lines, chunks, cur, in_fence, in_table = text.splitlines(), [], [], False, False
    for ln in lines:
        s = ln.strip()
        if s.startswith("```"):
            in_fence = not in_fence
        in_table = s.startswith("|") and not in_fence
        # Si la ligne OUVRE un bloc et qu'on dépasse, on flush AVANT le bloc
        if (s.startswith("```") and not in_fence is False) or (in_table and not cur_table):
            pass  # (logique simplifiée — voir le chunker complet §85)
        cur.append(ln)
    ...
```

> **Implémentation réelle :** voir `chunk_markdown_sections` au §85, qui
> traite les fences comme des atomes et découpe les sections longues
> **entre** les paragraphes, jamais dans un bloc.

## 83. Overlap : combien et pourquoi

L'overlap combat la **« cécité de frontière »** : un fait à cheval sur la
coupure entre chunk 47 et 48 est invisible au retrieval (47 finit en plein
milieu d'une phrase, 48 commence sans contexte).

```text
Chunk 47 : ...la commande display interface brief affiche l'état...
Chunk 48 : ...des ports. Pour filtrer : display interface brief | include up
Requête "filtrer les ports up" → AUCUN des deux ne matche bien.
Avec 50 tokens d'overlap, le chunk 48 commence 50 tokens plus tôt → le fait
entier est dans UN chunk → match.
```

**Combien (vérifié 2026) :**
- **Minimum pratique : 10 %** (50 tokens sur 512).
- **Zone recommandée : 10-20 %** (NVIDIA FinanceBench : 15 % = meilleur
  ratio précision/coût).
- **Au-delà de 25 %** : rendements décroissants (tu dupliques du contexte
  que le modèle traite comme du bruit).

**Avec le chunking par sections (§78)** : l'overlap est moins critique
(les sections sont des unités naturelles), mais garde le **titre de
section** en en-tête de chaque sous-chunk = « overlap sémantique » gratuit.

## 84. Métadonnées par chunk : CRITIQUE pour filtrer

Un chunk sans métadonnées est **introuvable précisément**. Chaque chunk
porte :

| Champ | Exemple | Usage |
|---|---|---|
| `source` | `huawei_cli_ref` | `WHERE source = ...` (séparer collections §99) |
| `doc_title` | `NetEngine AR1000V Command Reference` | citation |
| `section` | `## Configuration iStack` | filtre + citation |
| `page` | `2317` | traçabilité PDF |
| `chunk_index` | `42` | ordre, reconstruction |
| `chunk_hash` | `sha256` | dédup exacte (§95) |
| `lang` | `fr` | filtre langue |
| `corpus_version` | `v1` | invalidation (§102) |

> **Astuce « fil d'Ariane »** : préfixe chaque chunk avec son chemin de
> section (`Guide Debian > 12. APT > 12.3 Pinning`) dans le texte embeddé
> **et** en métadonnées. Le retrieval gagne en précision car le contexte
> hiérarchique est dans le vecteur lui-même. (Garde aussi le texte brut
> sans préfixe dans une colonne séparée pour l'affichage propre.)

## 85. EXEMPLE COMPLET : chunker markdown par sections + métadonnées

`scripts/chunk_md.py` — **ton chunker principal** (testable en CI, §20) :

```python
"""Chunker markdown par sections ## avec métadonnées.
Usage : python scripts/chunk_md.py --input guides/ --output build/chunks.jsonl
"""
from __future__ import annotations
import argparse, hashlib, json, re
from pathlib import Path
import tiktoken

ENC = tiktoken.encoding_for_model("text-embedding-3-small")
MAX_TOKENS = 500       # taille cible d'un chunk (§77)
OVERLAP_TOKENS = 50    # 10 % (§83)

RE_H2 = re.compile(r"^##\s+(.*)$")

def count_tokens(text: str) -> int:
    return len(ENC.encode(text))

def chunk_hash(text: str) -> str:
    return hashlib.sha256(text.encode("utf-8")).hexdigest()

def split_long_section(title: str, body: str) -> list[str]:
    """Découpe une section trop longue en sous-chunks de ~500 tokens,
    SANS couper les fences ``` ni les tableaux."""
    lines = body.splitlines()
    chunks, cur, cur_tok = [], [], 0
    i = 0
    while i < len(lines):
        # Regroupe les lignes en "blocs insécables" (paragraphe / fence / tableau)
        if lines[i].strip().startswith("```"):
            j = i + 1
            while j < len(lines) and not lines[j].strip().startswith("```"):
                j += 1
            j = min(j + 1, len(lines))
            block = "\n".join(lines[i:j]); i = j
        elif lines[i].strip().startswith("|"):
            j = i
            while j < len(lines) and lines[j].strip().startswith("|"):
                j += 1
            block = "\n".join(lines[i:j]); i = j
        else:
            j = i
            while j < len(lines) and lines[j].strip() not in ("", "```") \
                    and not lines[j].strip().startswith("|"):
                j += 1
            block = "\n".join(lines[i:j]); i = max(j, i + 1)
            if not block.strip():
                continue
        bt = count_tokens(block)
        # Bloc seul > MAX_TOKENS (rare) : découpe dure au token près
        if bt > MAX_TOKENS:
            toks = ENC.encode(block)
            for k in range(0, len(toks), MAX_TOKENS - OVERLAP_TOKENS):
                chunks.append(ENC.decode(toks[k:k + MAX_TOKENS]))
            continue
        if cur and cur_tok + bt > MAX_TOKENS:
            chunks.append("\n".join(cur))
            # overlap : on repart avec la fin du chunk précédent
            tail_toks = ENC.encode("\n".join(cur))[-OVERLAP_TOKENS:]
            cur, cur_tok = [ENC.decode(tail_toks)], OVERLAP_TOKENS
        cur.append(block); cur_tok += bt
    if cur:
        chunks.append("\n".join(cur))
    # Chaque sous-chunk garde le titre de section en en-tête (§83)
    return [f"## {title}\n\n{c}" for c in chunks]

def chunk_markdown_sections(md: str, source: str, doc_title: str) -> list[dict]:
    """1 section ## = 1 chunk (ou N sous-chunks). Retourne des dicts
    {text, metadata} prêts pour l'embedding."""
    chunks, cur_title, cur_lines = [], "intro", []
    def flush():
        if not any(l.strip() for l in cur_lines):
            return
        body = "\n".join(cur_lines).strip()
        parts = [body] if count_tokens(body) <= MAX_TOKENS \
            else split_long_section(cur_title, body)
        for idx, part in enumerate(parts):
            text = part if part.startswith("##") else f"## {cur_title}\n\n{part}"
            chunks.append({
                "text": text,
                "hash": chunk_hash(text),
                "metadata": {
                    "source": source,
                    "doc_title": doc_title,
                    "section": cur_title,
                    "chunk_index": len(chunks),
                    "tokens": count_tokens(text),
                    "corpus_version": "v1",
                },
            })
    for line in md.splitlines():
        m = RE_H2.match(line)
        if m:
            flush(); cur_title, cur_lines = m.group(1).strip(), []
        else:
            cur_lines.append(line)
    flush()
    return chunks

