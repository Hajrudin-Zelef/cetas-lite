# NOTES — corpus `etape6-phased1-fabrics-spine-leaf`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-fabrics-spine-leaf` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs` — 750 lignes, 13 chunks.

- Wave 1 — Spine/leaf & Clos fundamentals (design principles, oversubscription, scaling math, failure domains, cabling)
- Wave 2 — Core/distribution/access vs spine-leaf: fit, trade-offs, migration
- Wave 3 — Vendor reference designs
- Wave 4 — Fabric-level underlay essentials (numbered vs unnumbered, cabling plans)
- Wave 5 — Real deployments: scale examples, references, costs
- Wave 6 — Design decision checklist (synthesis)
- Coverage audit & open items
- Wave 7 — Clos sizing worked examples with published sources
- Wave 8 — Cisco deep-dive: verified scalability numbers and the 2026 ACI-vs-EVPN debate
- Wave 9 — Arista deep-dive: 800G hardware and AVD data model
- Wave 10 — NVIDIA deep-dive: Spectrum-X launch facts and Cumulus reference design
- Wave 11 — Dell and Aruba deep-dive
- Wave 12 — Juniper deep-dive: 5-stage EVPN-VXLAN JVD with Apstra
- Wave 13 — Switch-silicon radix update: Tomahawk 6 systems (2025–2026)
- Wave 14 — Hyperscaler deployments deep-dive
- Wave 15 — Fabric economics: whitebox TCO and market data
- Wave 16 — Updated coverage audit & open items
- Wave 17 — Cisco 2026 update: Silicon One G300, Nexus One, Hyperfabric, and a fabric security note
- Wave 18 — AI backend fabric topologies: rail-only vs rail-optimized, ROD vs RUD
- Wave 19 — Fabric-level multi-homing options (vPC / VSX / MC-LAG / ESI)
- Wave 20 — Deterministic cabling plan: worked example
- Wave 21 — Final coverage audit & open items
- Wave 22 — Vendor reference-design comparison matrix (synthesis)
- Wave 23 — Oversubscription planning worksheet (worked profiles)
- Wave 24 — Deployment notes and field patterns
- Wave 25 — Final coverage audit (supersedes Wave 21)
- Wave 26 — Switch-silicon generational table (radix context for fabric design)
- Wave 27 — Fabric glossary (terms used in this file)
- Wave 28 — Source index (key URLs consulted)
- Wave 29 — Brownfield migration: decision matrix and phase plan
- Wave 30 — Worked example: 64-leaf fabric bill of materials and acceptance checks
- Wave 31 — Key numbers at a glance & tag audit

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
