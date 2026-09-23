# NOTES — corpus `etape7-phasee-storage-software`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-storage-software` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)` — 751 lignes, 11 chunks.

- File header
- A. Ceph — release train 2024–2026
- B. Ceph — architecture and core concepts
- C. Ceph — BlueStore internals and OSD planning
- D. rclone — "rsync for cloud storage"
- E. TrueNAS — SCALE vs CORE in 2026
- F. OpenMediaVault (OMV)
- G. Unraid
- H. Appliance vendors: Synology, QNAP, Asustor
- I. ZFS / OpenZFS deep-dive
- J. Decision guides and comparison
- K. Conflicts, gaps and unverified claims
- L. Source index (verbatim URLs)
- M. Ceph — operations, deployment and ecosystem (expansion)
- N. NAS vendors and ZFS — further detail (expansion)
- O. Further operational detail (final expansion)
- P. Final additions
- Q. Quick-reference checklists
- R. Version-pinning reference (cutoff 2026-09-22)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
