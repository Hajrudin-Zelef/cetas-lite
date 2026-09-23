# NOTES — corpus `etape10-phaseb-news-mobile`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-news-mobile` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 10 — Phase B: Mobile & Devices News (February → September 2026)` — 751 lignes, 13 chunks.

- B0. Scope, method, provenance
- B1. Chronology — February 2026
- B2. Chronology — March 2026 (MWC Barcelona, March 2–5)
- B3. Chronology — April–May 2026
- B4. Chronology — June 2026
- B5. Chronology — July 2026 (Galaxy Unpacked, London, July 22)
- B6. Chronology — August 2026 (Made by Google, August 12)
- B7. Chronology — September 2026
- B8. Deep dive — Apple
- B9. Deep dive — Samsung
- B10. Deep dive — Google
- B11. Deep dive — Chinese & other brands
- B12. Chipsets — the 2026 silicon story
- B13. Wearables & smart glasses
- B14. Tablets & e-readers
- B15. Mobile OS & ecosystem
- B16. Connectivity — 5G, satellite, 6G
- B17. Flagship comparison tables (Sept 2026)
- B18. Market data & industry notes
- B19. Reviews consensus (Feb–Sept 2026)
- B20. Gaps & conflicts register
- B21. Glossary
- B22. Source index (verbatim URLs)
- B23. Mid-range & budget wave (Feb–Sept 2026)
- B24. Mobile security & safety (Feb–Sept 2026)
- B25. Carriers, eSIM & retail
- B26. Foldable market analysis (2026)
- B27. Camera technology trends (2026)
- B28. Battery & charging trends (2026)
- B29. What to watch (post-cutoff)
- B30. Key-dates timeline (Feb → Sep 22, 2026)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
