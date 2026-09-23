# NOTES — corpus `etape10-phased-news-sport`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-news-sport` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 10D — Sports News (February 2026 → 22 September 2026)` — 935 lignes, 13 chunks.

- 0. Scope, method and provenance
- 1. February 2026
- 2. March 2026
- 3. April 2026
- 4. May 2026
- 5. June 2026
- 6. July 2026
- 7. August 2026
- 8. September 2026 (through the 22nd)
- 9. Deep dive — FIFA World Cup 2026
- 10. Deep dive — European club football 2025-26 and the summer window
- 11. Deep dive — NBA 2025-26
- 12. Deep dive — Winter Olympics Milano Cortina 2026
- 13. Deep dive — tennis 2026 (Grand Slams)
- 14. Formula 1, cycling, athletics and other sports
- 15. Sports games and product releases
- 16. Conflicts and gaps register
- 17. Glossary
- 18. Source index (verbatim URLs)
- 19. End of file

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
