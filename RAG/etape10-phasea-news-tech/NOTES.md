# NOTES — corpus `etape10-phasea-news-tech`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-news-tech` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 10A — Tech News: AI, Chips, Cloud, Cybersecurity (February → 22 September 2026)` — 767 lignes, 11 chunks.

- Provenance legend
- February 2026
- March 2026
- April 2026
- May 2026
- June 2026
- July 2026
- August 2026
- September 2026 (1–22 September)
- Cross-month thematic syntheses
- Gaps, conflicts and data-quality register
- Glossary
- Source index (verbatim URLs)
- Appendix A — February → September 2026 at a glance (timeline table)
- Appendix B — Frontier-model release ledger (Feb → 22 Sep 2026)
- Appendix C — Cloud-incident ledger (Jul–Aug 2026)
- Appendix D — Breach and ransomware ledger (Feb → 22 Sep 2026)
- Appendix E — Datacenter, capex and financing ledger (2026)
- Appendix F — Regulatory and policy calendar (2026)
- Appendix G — Event deep-dives (from the same sources, no new claims)
- Appendix H — Corroboration notes (items with 2+ independent sources in this corpus)
- Appendix I — Monthly analyst notes (synthesis, not new facts)
- Glossary (extended)
- Appendix J — Key figures: prices, fines, fees and counts (all sourced above)
- Appendix K — People and organizations index

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
