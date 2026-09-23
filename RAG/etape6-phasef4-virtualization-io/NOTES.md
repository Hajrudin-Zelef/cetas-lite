# NOTES — corpus `etape6-phasef4-virtualization-io`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-virtualization-io` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Phase F4 — I/O Virtualization & CPU Acceleration Extensions` — 751 lignes, 12 chunks.

- Wave 1 — virtio: spec, transports, and cloud reality
- Wave 2 — SR-IOV: VFs, switchdev, and the live-migration trade-off
- Wave 3 — CPU virtualization extensions: VT-x/VT-d, AMD-V/AMD-Vi, ATS/PRI/PASID
- Wave 4 — AVX-512, AVX10, and AMX: SIMD for packet and AI work
- Wave 5 — DPDK & SPDK: 2026 releases, PMD coverage, and where they run
- Wave 6 — GPU virtualization: vGPU, MIG, SR-IOV graphics, and AI relevance
- Part C — Decision guides
- Part D — Conflicts, gaps, and open items
- Part E — Glossary (key terms)
- Part F — Source index (verbatim URLs)
- Wave 7 — vDPA and the virtio 1.4 admin-queue track (deep dive)
- Wave 8 — DPDK PMD ecosystem / OVS-DPDK and NVIDIA MIG profiles
- Wave 9 — SR-IOV operations, VT-d internals, and vhost-user protocol notes
- Part G — Research log
- Wave 10 — ARM virtualization, virtio-fs, and the rust-vmm ecosystem
- Part H — Sourced performance figures compendium (do not mix eras)
- Part F (continued) — additional source URLs
- Wave 11 — DPDK security offloads and NVMe-oF transport selection
- Part F (continued 2) — additional source URLs
- Wave 12 — Hyperscaler NIC programs (2026) and virtio-net feature reference
- Part F (continued 3) — additional source URLs
- Wave 13 — SPDK vhost targets: vhost-blk / vhost-scsi for VMs
- Part F (continued 4) — additional source URLs
- Part J — Delivery notes and known thin spots
- Appendix K — DPDK PMD / device quick lookup (all names sourced from DPDK docs)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
