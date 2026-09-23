# NOTES — corpus `etape10-phaseg-news-science`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-news-science` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 10 Phase G — Science & Planet News (February → September 2026)` — 752 lignes, 12 chunks.

- 1. Month-by-month chronology (February → September 2026)
- 2. Deep dive — Human & robotic spaceflight 2026
- 3. Deep dive — Astronomy 2026
- 4. Deep dive — Climate & extreme weather 2026
- 5. Deep dive — Health & medicine 2026
- 6. Deep dive — Energy science, oceans & environment
- 7. Gaps & conflicts register
- 8. Glossary
- 9. Source index (verbatim URLs)
- 10. Spaceflight technical snapshots (appendix)
- 11. Numerical records & superlatives 2026 (Feb → Sept)
- 12. Forward watchlist (October → December 2026)
- 13. Source-reliability notes (methodology appendix)
- 14. Expanded month notes (second pass — additional detail)
- 15. JWST 2026: the complete discovery log
- 16. Climate: ENSO and ocean dynamics in 2026
- 17. The 2026 fusion landscape: company-by-company
- 18. Health: the 2026 Alzheimer's pipeline
- 19. Cross-cutting theme: AI in science (2026)
- 20. People of 2026
- 21. Starship Flight 12 engineering dossier (May 22, 2026)
- 22. Structured disagreements of 2026 (reader's guide)
- 23. Chapter summaries for retrieval
- 24. Final glossary additions
- 25. Notable 2026 publications & reports log
- 26. Agency & dataset fact sheets (what each source measured)
- 27. Verified 2026 flight & event chronology (one-line index)
- 28. Milestone watchboard: announced 2026 targets vs status
- 30. Quick-reference key-dates card (2026)
- 29. Editorial close: what this file does and doesn't claim

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
