# NOTES — corpus `etape10-phasee-news-world`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-news-world` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Etape 10E — World News: Wars, Geopolitics, Elections, Economy, Fortunes (February 2026 – 22 September 2026)` — 750 lignes, 13 chunks.

- 1. Scope, method and provenance legend
- 2. Month-by-month chronology (February → September 2026)
- 3. Thematic deep-dives
- 4. Gaps, conflicts, and unverified claims register
- 5. Glossary
- 6. Source index (exact URLs as returned by research)
- 7. Annotated source guide (what each source supports)
- 8. Key figures at a glance (all dated; see main text for provenance)
- 9. Notable attributed quotes (verbatim as reported)
- 10. Research log and methodological notes
- 11. Names index (roles as reported)
- 12. Version and corrections note

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
