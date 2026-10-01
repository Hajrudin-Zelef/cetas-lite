---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-14
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["decode", "embedding", "embeddings", "valuation"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [2588, 2779]
sha256: 0eb7a17fb7ca9e21b565e54a20d7937fdfd3dc80be3158cd122008a2ec54de01
---

# Outils dev + ingénierie RAG (chunk & corpus)

def main() -> None:
    ap = argparse.ArgumentParser()
    ap.add_argument("--input", type=Path, required=True)    # dossier de .md
    ap.add_argument("--output", type=Path, required=True)   # chunks.jsonl
    ap.add_argument("--source", default="guides")
    args = ap.parse_args()
    args.output.parent.mkdir(parents=True, exist_ok=True)
    n_files = n_chunks = 0
    with args.output.open("w", encoding="utf-8") as f:
        for md_file in sorted(args.input.rglob("*.md")):
            md = md_file.read_text(encoding="utf-8")
            for ch in chunk_markdown_sections(md, args.source, md_file.stem):
                f.write(json.dumps(ch, ensure_ascii=False) + "\n")
                n_chunks += 1
            n_files += 1
    print(f"{n_files} fichiers → {n_chunks} chunks → {args.output}")

if __name__ == "__main__":
    main()
```

```bash
python scripts/chunk_md.py --input ~/workspace/user/files/ --output build/chunks.jsonl --source guides
head -c 600 build/chunks.jsonl
```

## 86. EXEMPLE COMPLET : chunker fixe avec overlap (secours)

Pour les articles web sans structure `##` (tes `articles_*.json`) :

```python
"""Chunker fixe à fenêtre token + overlap. Secours quand pas de structure."""
from __future__ import annotations
import argparse, hashlib, json
from pathlib import Path
import tiktoken

ENC = tiktoken.encoding_for_model("text-embedding-3-small")

def chunk_fixed(text: str, window: int = 500, overlap: int = 50) -> list[str]:
    toks = ENC.encode(text)
    stride = window - overlap
    return [ENC.decode(toks[i:i + window]) for i in range(0, len(toks), stride)]

def main() -> None:
    ap = argparse.ArgumentParser()
    ap.add_argument("--input", type=Path, required=True)   # articles.json
    ap.add_argument("--output", type=Path, required=True)  # chunks.jsonl
    ap.add_argument("--window", type=int, default=500)
    ap.add_argument("--overlap", type=int, default=50)
    args = ap.parse_args()
    data = json.loads(args.input.read_text(encoding="utf-8"))
    # data = [{"url":..., "title":..., "text":...}] — TON format articles_*.json
    with args.output.open("w", encoding="utf-8") as f:
        for art in data:
            for idx, part in enumerate(chunk_fixed(art["text"], args.window, args.overlap)):
                ch = {
                    "text": part,
                    "hash": hashlib.sha256(part.encode()).hexdigest(),
                    "metadata": {
                        "source": "articles",
                        "doc_title": art.get("title", ""),
                        "url": art.get("url", ""),
                        "chunk_index": idx,
                        "tokens": len(ENC.encode(part)),
                        "corpus_version": "v1",
                    },
                }
                f.write(json.dumps(ch, ensure_ascii=False) + "\n")
    print(f"OK → {args.output}")

if __name__ == "__main__":
    main()
```

## 87. Déduplication par hash (exacte)

```python
"""Filtre les chunks en double EXACT (même sha256)."""
import json, sys
from pathlib import Path

def dedup_exact(in_path: Path, out_path: Path) -> tuple[int, int]:
    seen, kept, total = set(), 0, 0
    with in_path.open(encoding="utf-8") as fin, out_path.open("w", encoding="utf-8") as fout:
        for line in fin:
            total += 1
            ch = json.loads(line)
            if ch["hash"] in seen:
                continue            # doublon exact → poubelle
            seen.add(ch["hash"])
            fout.write(line); kept += 1
    return kept, total - kept

if __name__ == "__main__":
    kept, dropped = dedup_exact(Path(sys.argv[1]), Path(sys.argv[2]))
    print(f"Gardés : {kept} — doublons exacts supprimés : {dropped}")
```

> **Ton cas réel** : les 12 620 commandes CLI Huawei contiennent des
> doublons (même commande documentée dans plusieurs chapitres, en-têtes
> répétés). La dédup exacte les élimine **avant** l'embedding → moins de
> tokens payés, moins de bruit au retrieval.

## 88. Tailles d'embeddings et stockage (calcul)

```text
text-embedding-3-small : 1536 dims × 4 octets (float32) = 6 144 o ≈ 6 Ko/vecteur

Formule : stockage_vecteurs ≈ N_chunks × 1536 × 4 octets
          stockage_total   ≈ × 2.5 à 3 (texte + métadonnées + index HNSW)

Exemples :
   10 000 chunks →  60 Mo de vecteurs → ~150-200 Mo au total (Postgres)
  100 000 chunks → 600 Mo de vecteurs → ~1.5-2 Go au total
1 000 000 chunks →   6 Go de vecteurs → ~15-20 Go au total

Coût embedding : N_chunks × 500 tokens × $0.02/1M
   10 000 chunks → 5M tokens → $0.10
  100 000 chunks → 50M tokens → $1.00
```

> Moralité : à ton échelle, le stockage et l'embedding sont **négligeables**.
> Le vrai coût, c'est la **qualité** (mauvais chunks = mauvaises réponses).

## 89. Tableau comparatif des stratégies

| Stratégie | Complexité | Coût | Précision | Quand l'utiliser |
|---|---|---|---|---|
| Fixe 512/50 | ★ | ★ | ★★ | articles web sans structure |
| **Sections `##`** | ★★ | ★ | ★★★ | **tes guides (défaut)** |
| Récursive | ★★ | ★ | ★★ | sections trop longues, HTML |
| Sémantique | ★★★ | ★★★ (3-5x) | ★★★ ? | si évaluation le justifie |
| Proposition | ★★★ | ★★ | ★★★ (factoid) | FAQ, fiches réflexes |
| Parent-enfant | ★★★ | ★★ | ★★★ | précision + contexte |

## 90. Pièges du chunking (10)

1. **Chunks orphelins** : un titre `##` seul ou une phrase coupée n'a aucun
   sens hors contexte. Toujours un en-tête de section dans le chunk.
2. **Doublons des CLI Huawei** (ton cas réel) : 12 620 commandes, beaucoup
   de redites inter-chapitres → dédup exacte **avant** embedding (§87).
3. **Compter en caractères** : voir §76 — troncature silencieuse.
4. **Couper les fences ```** : voir §82 — deux chunks morts.
5. **Zéro métadonnées** : impossible de filtrer par source/section après coup.
   Les métadonnées se mettent **au chunking**, jamais après.
6. **Overlap à 0** : cécité de frontière (§83). Minimum 10 %.
7. **Chunks géants « au cas où »** : 2000 tokens diluent l'embedding ;
   le LLM reçoit 3 chunks × 2000 = 6000 tokens de bruit.
8. **Mélanger les langues** dans un chunk (ex. : doc FR + exemple EN) :
   acceptable, mais la métadonnée `lang` doit refléter le majoritaire.
9. **Embedder les sommaires/TDM** : « 1. Intro … 2. Install … » n'apporte
   rien au retrieval et pollue. Filtre les sections < 50 tokens utiles.
10. **Re-chunker sans re-embedder** : changer la stratégie invalide TOUS
    les vecteurs → re-embed complet (pas cher : §71, mais à planifier §102).

## 91. Checklist chunking

- [ ] Tokenizer réel (`tiktoken`, `cl100k_base`) partout
- [ ] Stratégie par type de doc (sections pour guides, fixe pour articles)
- [ ] Fences et tableaux jamais coupés
- [ ] Overlap 10-20 %
- [ ] Métadonnées complètes (source, section, page, hash)
- [ ] Dédup exacte avant embedding
- [ ] Échantillon de 20 chunks **relu à la main** avant le batch complet

## 92. Pense-bête chunking

```bash
python scripts/chunk_md.py --input guides/ --output build/chunks.jsonl --source guides
python scripts/chunk_fixed.py --input articles.json --output build/chunks_articles.jsonl
python scripts/dedup_exact.py build/chunks.jsonl build/chunks_dedup.jsonl
```

---

## 93. Construire un corpus : vue d'ensemble

```text
COLLECTE → NETTOYAGE → NORMALISATION → DÉDUPLICATION → VERSIONING
   (§94)      (§95)         (§96)            (§97-99)        (§100)
      → SÉPARATION COLLECTIONS → ÉVALUATION → MAINTENANCE
           (§101)                  (§103)        (§104)
```

**Règle d'or :** garde TOUJOURS le brut (`corpus/brut/`, hors git) séparé
du nettoyé (`build/`). Quand ta logique de nettoyage s'améliore, tu
**re-joues** le pipeline sans re-scraper (le scraping est le plus cher en
temps et en risque de ban).

## 94. Collecte : tes scripts existants

