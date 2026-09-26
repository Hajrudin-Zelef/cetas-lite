# NOTES — corpus `collect-opencode-docs`

## Génération

Mode **`files`** (`RAG/_tools/build_rag.py`) : **une fiche = un chunk**, copiée verbatim
de `docs/RAG/Collect RAG/04_opencode_docs`. Chaque fiche suit le template `_TEMPLATE.md` de la collecte (titre H1,
Metadata, Full summary, Key points, Technical data, Why this source matters).

Chaque chunk reçoit un en-tête YAML avec `source` = chemin de la fiche source,
`source_lines: [1, N]` (bornes de la fiche), `task` dérivée du champ `Type`
(article, model-card, benchmark, tutorial, documentation, review).

- **6 fiches** · **14 chunks** (chunker v2, garde-fou 8000 car. : les fiches
  `config` et `providers` — ajoutées le 2026-09-26 — sont découpées, les 4 autres
  tiennent en un chunk).
- `00_INDEX.md` et `_TEMPLATE.md` (à la racine de la collecte) sont exclus du chunking
  (scaffolding de collecte, pas un contenu à interroger). Les sources restent intacts
  sur le disque.

## Contenu

opencode — 4 fiches collectées le 2026-09-23, statut ✅ reachable sauf
mention. Chaque fiche : URL source, site, type, langue, résumé structuré (300–800 mots),
points clés, données techniques, valeur ajoutée pour le RAG.

## Fiabilité

- `verify_rag.py` mode `files` : une fiche = un chunk, sha256 de chaque fiche == chunk,
  corps verbatim dans le chunk, `source_files` cohérent.
- Aucune réécriture du contenu des fiches ; seul ajout = l'en-tête YAML + éventuel H1.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
