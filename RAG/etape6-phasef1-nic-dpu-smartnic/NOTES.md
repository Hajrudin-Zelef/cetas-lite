# NOTES — corpus `etape6-phasef1-nic-dpu-smartnic`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-nic-dpu-smartnic` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)` — 756 lignes, 12 chunks.

- Wave 1 — NVIDIA BlueField family: BF-2, BF-3, BF-4
- Wave 2 — DOCA software framework
- Wave 3 — AMD Pensando portfolio
- Wave 4 — Intel IPU line
- Wave 5 — Marvell OCTEON, Broadcom Stingray, Fungible, other DPU vendors
- Wave 6 — NVIDIA ConnectX NIC line (DPU-adjacent, 2026)
- Wave 7 — Offloads deep-dive
- Wave 8 — Competitive landscape and adoption (2026)
- Wave 9 — Open items, conflicts, glossary, source index
- Wave 10 — OEM integration and form factors
- Wave 11 — Ultra Ethernet 1.0 deep-dive (transport context)
- Wave 12 — Master comparison tables
- Wave 13 — Pricing recap matrix (all 2026 snapshots)
- Wave 14 — DOCA release detail, P4 ecosystem, consolidation correction
- Wave 15 — AI fabric architecture and deployment framework
- Wave 16 — Broadcom: Thor 2, Thor Ultra, NetXtreme E-Series, Stingray status
- Wave 17 — Marvell OCTEON 10 SKUs, remaining vendors, lineage timelines
- Wave 18 — 2026 chronology, optics/form factors, management & security
- Wave 19 — Generational spec tables (quick-reference)
- Wave 20 — Pre-deployment validation checklist (field notes)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
