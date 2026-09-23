# NOTES — corpus `etape9-phased-data-protection-raid`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-data-protection-raid` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 9 — Phase D: Data Protection & RAID Hardware` — 750 lignes, 11 chunks.

- D1 — Data Loss Protection (DLP): what "protection" means at drive level
- D2 — End-to-end data path protection: T10 PI / DIF / DIX and NVMe
- D3 — Firmware resilience: NIST SP 800-193 and storage-specific controls
- D4 — Secure erase and sanitize: retiring drives without leaking data
- D5 — Reliability metrics: MTBF, AFR, UBER, and how to read them
- D6 — Wear leveling, garbage collection, over-provisioning, TRIM
- D7 — SMART monitoring and thermal management
- D8 — Hardware RAID controllers: 2026 landscape
- D9 — RAID levels: math, write penalties, and rebuild reality
- D10 — HBAs, tri-mode, and expanders
- D11 — NVMe hardware RAID: GRAID SupremeRAID and the software alternative
- D12 — JBOD / IT mode for ZFS and Ceph: why "HBA mode" matters
- D13 — Decision matrices: hardware RAID vs ZFS vs Ceph erasure coding
- D14 — Tape in 2026: LTO-9/LTO-10 status, LTFS, and market signals
- D15 — Optical archival: status in 2026
- D16 — Operational playbook: protecting data day to day
- D17 — Worked numerical examples (reproducible arithmetic)
- D18 — Gap and conflict register
- D19 — Glossary
- D20 — Source index (verbatim URLs)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
