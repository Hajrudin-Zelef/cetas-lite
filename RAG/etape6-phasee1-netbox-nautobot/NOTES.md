# NOTES — corpus `etape6-phasee1-netbox-nautobot`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-netbox-nautobot` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)` — 750 lignes, 11 chunks.

- Wave 1 — NetBox core: version line and platform features (2024–2026)
- Wave 2 — NetBox Labs commercial ecosystem: Cloud, Assurance, Diode, Discovery
- Wave 3 — NetBox plugins ecosystem
- Wave 4 — Nautobot core: version line and platform features (2023–2026)
- Wave 5 — Nautobot Apps: marketplace and flagship automation apps
- Wave 6 — Head-to-head: NetBox vs Nautobot (2026 view)
- Wave 7 — Adjacent IPAM/DCIM tools
- Wave 8 — Real deployments, adoption signals, and company metrics
- Wave 10 — NetBox Labs portfolio expansion: AI, discovery, and ingestion
- Wave 11 — Deployment & operations: NetBox and Nautobot in production
- Wave 12 — Migration tooling and cross-pollination
- Wave 13 — AI assistants and professional services
- Wave 14 — Consolidated reference tables
- Wave 15 — Gaps, conflicts, and unverified claims log
- Wave 16 — NetBox Labs in 2026: platform, funding, and community signals
- Wave 17 — Glossary of terms used in this file
- Wave 18 — API and automation-engine deep dive (release-note-grounded)
- Wave 19 — Selection guide: when the collected evidence favors which platform `[secondary]`
- Wave 20 — Consolidated source index (all URLs cited in this file)
- Wave 21 — Timeline: NetBox & Nautobot, January 2024 → September 2026 `[official][vendor-reported]`
- Wave 22 — Follow-up research seeds (post-cutoff work for a later pass)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
