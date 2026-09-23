# NOTES — corpus `etape6-phasef2-fpga`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-fpga` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Phase F2 — FPGA (Field-Programmable Gate Arrays)` — 754 lignes, 14 chunks.

- 1. Altera — post-Intel independence
- 2. AMD — Versal adaptive SoC families
- 3. AMD Alveo accelerator/SmartNIC cards
- 4. Achronix Speedster7t
- 5. Lattice Semiconductor — Nexus and Avant
- 6. Microchip PolarFire
- 7. FPGA SmartNIC vendors
- 8. Packet processing and P4 on FPGAs
- 9. Toolchains — vendor flows
- 10. Pricing snapshots (as observed 2026-09-22)
- 11. TCO: FPGA vs ASIC vs GPU
- 12. Adoption evidence
- 13. Open-source FPGA tooling
- 14. Conflicts, gaps, and uncertainty
- 15. Key URLs
- 16. Altera Agilex 7 — detailed family plan (device overview 683458, rev 2025.12.08)
- 17. AMD Versal AI Core — figures of merit (VC1502 → VC2802)
- 18. Napatech — NT400 and N3070X (Agilex-based SmartNICs)
- 19. Lattice — Nexus 2 detail, Certus-N2, Mach-N2, Avant 30/50
- 20. Microchip PolarFire — additional detail
- 21. FPGA SmartNIC ecosystem — additional vendors and notes
- 22. P4, DPDK, and packet-processing — extended
- 23. Hyperscaler and telco adoption — extended notes
- 24. Open-source tooling — extended
- 25. Pricing and TCO — extended snapshots
- 26. Conflicts, gaps, uncertainty (consolidated)
- 27. Additional key URLs (sections 16–24)
- 28. Agilex 5 — series detail
- 29. AMD Alveo media/video accelerators
- 30. Achronix — company and ecosystem detail
- 31. Lattice — Avant detail and legacy families
- 32. Microchip — tools and ecosystem
- 33. Exablaze / Cisco — ultra-low-latency NICs
- 34. Microsoft Catapult and FPGA DNN acceleration
- 35. Open-source tooling — extended inventory
- 36. Timeline — key FPGA events (2025-01 → 2026-09-22)
- 37. SKU / part-number index (quick reference)
- 38. Open FPGA Stack (OFS) — concepts
- 39. AMD adaptive-computing portfolio notes
- 40. Networking standards context for FPGA NICs
- 41. Security features on data-center FPGAs
- 42. Power and thermal envelopes (cards)
- 43. Methodology and limitations of this file
- 44. Cross-references to other Phase F files
- 45. Altera software stack
- 46. AMD software ecosystem for Versal/Alveo
- 47. FPGA market structure context
- 48. Emulation and prototyping (VP1902 context)
- 49. Defense, aerospace, and rugged FPGAs
- 50. Edge-AI FPGA notes
- 51. Glossary — acronyms used in this file
- 52. Open questions for future research passes
- 53. More Alveo cards — U55C, U50, V70-class notes
- 54. Other FPGA board/system vendors (context)
- 55. Where to find current FPGA pricing (guidance, not prices)
- 56. Document statistics and integrity
- 57. Flagship-device quick comparison (one line each)
- 58. Errata and revision note

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
