# NOTES — corpus `etape9-phasee-storage-market`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-storage-market` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 9 Phase E — Storage & Memory Market 2026` — 476 lignes, 8 chunks.

- E0 — Scope, date, method, provenance legend
- E1 — The 2026 memory market in one picture
- E2 — NAND price trends: contract, spot, and street
- E3 — Vendor landscape 2026: who makes the memory
- E4 — M&A and corporate events 2025–2026
- E5 — HBM economics 2026: the memory that prices the GPU
- E6 — CXL: the "post-HBM battleground"
- E7 — The used/refurbished enterprise market (homelab angle)
- E8 — Firmware lock-in and platform friction (Dell/HPE)
- E9 — Buying guide: new vs used in 2026
- E10 — QLC vs TLC, DRAM-less pitfalls, and workload matching
- E11 — Counterfeit and misrepresented-drive detection
- E12 — AI storage sizing I: training checkpoints
- E13 — AI storage sizing II: inference KV-cache and VRAM
- E14 — Cost per GB: VRAM vs HBM vs DDR5 vs NAND
- E15 — Computational storage and open firmware: 2026 status
- E16 — Enterprise SSD roadmap: PCIe Gen6 and beyond
- E17 — Conflicts, gaps, and unverified claims register
- E18 — Glossary
- E19 — Source index (verbatim URLs)
- E3b — NAND technology race 2026: layers, density, and the China factor
- E5b — HBM generation specification table (JEDEC + vendor data)
- E7b — Used-market deep dive: SKUs, patterns, and price anchors
- E12b — Checkpoint math for MoE and 400B+ models
- E13b — KV-cache worked examples by model class
- E20 — Buyer FAQ (2026 market edition)
- E21 — Decision matrices
- E22 — 2026 timeline of market events
- E23 — Regional price notes
- E24 — TCO worked examples ($/TB/year)
- E25 — Watch list (post-cutoff)
- E26 — Methodology note: how to read 2026 storage prices

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
