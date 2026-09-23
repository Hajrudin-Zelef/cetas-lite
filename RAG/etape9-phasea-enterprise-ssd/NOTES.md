# NOTES — corpus `etape9-phasea-enterprise-ssd`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-enterprise-ssd` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 9 — Enterprise SSD Hardware (Phase A)` — 759 lignes, 13 chunks.

- 0. Scope, method, provenance
- 1. 2026 market and pricing context
- 2. Samsung datacenter SSDs
- 3. Kioxia enterprise SSDs
- 4. Micron enterprise SSDs
- 5. Solidigm enterprise SSDs
- 6. SK hynix enterprise SSDs
- 7. Western Digital Ultrastar DC SN861
- 8. Phison Pascari — the controller vendor's own SSD brand
- 9. SSD controllers (enterprise)
- 10. NAND flash deep dive (2026)
- 11. Endurance, reliability, and data protection
- 12. Prices, availability, and the homelab angle (2026)
- 13. Comparison matrix (Gen5 flagships, top SKUs)
- 14. Gaps, conflicts, and unverified claims register
- 15. Glossary
- 16. Source index (verbatim URLs)
- 17. Workload decision guide
- 18. Form factors and interface math
- 19. DRAM cache vs HMB and enterprise firmware
- 20. Endurance worked examples (5-year warranty)
- 21. Vendor detail expansions
- 22. Homelab buyer's checklist (expanded)
- 23. AI workloads and QLC economics
- 24. NAND industry structure (2026)
- 25. Controller market notes
- 26. Retail/consumer price divergence (context)
- 27. $/TB and $/TBW economics (dated snapshots, 2026)
- 28. Gen4 vs Gen5 purchase decision (2026)
- 29. NVMe 2.0 enterprise feature deep dive
- 30. Warranty and support notes
- 31. One-line cheat sheet per drive family
- 32. How to read a datacenter SSD datasheet (method notes)
- 33. Firmware update and fleet management
- 34. Security deep dive
- 35. Thermal and power design numbers
- 36. Form-factor availability matrix
- 37. Related RAG cross-references
- 15b. Glossary additions

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
