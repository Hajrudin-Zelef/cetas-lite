# NOTES — corpus `etape10-phasec-news-pc-mac`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-news-pc-mac` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 10 Phase C — PC & Mac News (February → September 2026)` — 751 lignes, 14 chunks.

- 0. Executive Summary
- 1. Chronology — February 2026
- 2. Chronology — March 2026
- 3. Chronology — April 2026
- 4. Chronology — May 2026
- 5. Chronology — June 2026
- 6. Chronology — July 2026
- 7. Chronology — August 2026
- 8. Chronology — September 2026 (through 09-22)
- 9. Deep Dive — CPUs
- 10. Deep Dive — GPUs
- 11. Deep Dive — Apple Mac & macOS
- 12. Deep Dive — Windows & Microsoft
- 13. Deep Dive — Laptops, Desktops, Handhelds
- 14. Deep Dive — Gaming
- 15. Market Data & Price Tables
- 16. Conflicts & Gaps Register
- 17. Glossary
- 18. Source Index (verbatim URLs)
- 19. January 2026 Context (pre-scope background)
- 20. Deep Dive — AI PC & Copilot+ (2026)
- 21. 2026 CPU Landscape — Comparison Table
- 22. 2026 GPU Landscape — What You Could Actually Buy
- 23. 2026 Game Release Calendar (verified items only)
- 24. The Memory Crisis — Cross-Cutting Impact on PC & Mac
- 25. Watch List — Q4 2026 and Beyond
- 26. Key Takeaways for RAG Indexing
- 27. Windows 11 in 2026 — The "Fix Windows 11" Year
- 28. Nintendo 2026 — Full News Roundup
- 29. Consolidated Dated Event Log (Feb → Sep 2026)
- 30. Key People & Outlets Index
- 31. Methodology Note
- 32. Deep Dive — Intel in 2026 (Panther Lake → Nova Lake)
- 33. Deep Dive — AMD in 2026 (AM5 Longevity → Zen 6)
- 34. Deep Dive — Apple in 2026 (M5 Ultra, M6, Mac Lineup)
- 35. Quarterly Recap — The Three Acts of 2026
- 36. 2026 Street Prices vs MSRP (GPU market)
- 37. Questions This File Answers (RAG retrieval aid)
- 38. Source Quality Map (outlets cited in this file)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
