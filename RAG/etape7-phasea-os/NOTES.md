# NOTES — corpus `etape7-phasea-os`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-os` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 7 — Phase A: Server Operating Systems (Linux OS Layer)` — 775 lignes, 14 chunks.

- A0 — Scope, observation date, method and provenance legend
- A1 — Executive snapshot: the server Linux landscape in September 2026
- A2 — Master lifecycle table
- A3 — Debian 13 "Trixie": the current stable
- A4 — Debian 12 "Bookworm": current status and LTS transition
- A5 — Debian 11 "Bullseye": LTS just ended — action required
- A6 — Debian cloud images, installer and provisioning
- A7 — Ubuntu 26.04 LTS "Resolute Raccoon": the new LTS
- A8 — Ubuntu 24.04 LTS "Noble Numbat": the workhorse
- A9 — Older Ubuntu LTS: 22.04 and 20.04 status
- A10 — Ubuntu Pro, ESM and the Legacy add-on: pricing and structure
- A11 — Canonical Landscape: fleet management
- A12 — Ubuntu Core: the immutable IoT/edge variant
- A13 — RHEL 10: current enterprise major
- A14 — RHEL 9: the incumbent enterprise major
- A15 — RHEL 8 and 7: legacy tail
- A16 — RHEL licensing and pricing: the subscription model
- A17 — RHEL in-place upgrades: Leapp
- A18 — Rocky Linux 10 "Red Quartz": current community enterprise major
- A19 — Rocky Linux 9 "Blue Onyx" and 8 "Green Obsidian"
- A20 — CIQ: the commercial company behind Rocky
- A21 — AlmaLinux 10 "Purple Lion": the x86_64-v2 differentiator
- A22 — AlmaLinux 9 and 8
- A23 — AlmaLinux Kitten 10: preview, not production
- A24 — Rocky vs AlmaLinux: decision factors
- A25 — Migration tooling: CentOS → Rocky/Alma, and major upgrades
- A26 — Oracle Linux: the secondary comparison
- A27 — TuxCare: extended lifecycle and live patching vendor
- A28 — Kernel version master table
- A29 — Package management compared
- A30 — Live kernel patching compared
- A31 — Security patching cadence and embargoes
- A32 — CIS hardening per distribution
- A33 — FIPS 140-3 validation status
- A34 — DISA STIG coverage
- A35 — Common Criteria, SELinux and AppArmor
- A36 — Adoption and market share: read the caveats first
- A37 — Proxmox VE's Debian relationship
- A38 — Cloud and virtualization images matrix
- A39 — Server vs desktop relevance
- A40 — Minimal and immutable variants: Fedora CoreOS, openSUSE MicroOS, bootc
- A41 — Commercial pricing comparison matrix (date-labeled)
- A42 — Decision guide: which distribution for which role
- A43 — Gaps, conflicts and non-comparable figures register
- A44 — Glossary
- A45 — Source index (verbatim URLs, in order of first use)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
