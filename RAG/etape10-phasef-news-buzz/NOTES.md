# NOTES — corpus `etape10-phasef-news-buzz`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-news-buzz` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`STEP 10 — Phase F: News, Buzz & Society (2026-02-01 → 2026-09-22)` — 754 lignes, 11 chunks.

- 0. Provenance legend
- 1. Month-by-month chronology
- 2. Deep dive — Cinema
- 3. Deep dive — Music
- 4. Deep dive — Gaming & interactive entertainment
- 5. Deep dive — Streaming & television
- 6. Deep dive — Pop culture, buzz & memes
- 7. Deep dive — Society, movements, legal cases, lifestyle
- 8. Notable deaths (Feb 1 → Sep 22, 2026)
- 9. Awards summary table
- 10. Recap tables
- 11. Conflict & gap register
- 12. Glossary
- 13. Source index (verbatim URLs)
- 14. Quick date index (every dated event in this file)
- 15. People of the period (Feb → Sep 2026)
- 16. Annotated source notes
- 17. Coverage completeness self-check
- 12b. Glossary — extended
- 18. Reading guide for RAG ingestion
- 19. Post-cutoff watchlist (excluded from coverage; for a future update)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
