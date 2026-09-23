# NOTES — corpus `etape7-phased-proxmox-backup`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-proxmox-backup` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research` — 755 lignes, 13 chunks.

- 1. Proxmox VE — product overview
- 2. Proxmox VE 9.x release line (2025–2026)
- 3. Proxmox VE architecture: KVM, LXC, clustering, HA
- 4. Proxmox VE storage: Ceph, ZFS, LVM, PBS integration
- 5. Proxmox VE networking & security: SDN, firewall, identity
- 6. Proxmox VE licensing & subscription pricing (2026)
- 7. Proxmox Backup Server (PBS)
- 8. Proxmox Datacenter Manager (PDM)
- 9. Hypervisor alternatives
- 10. Veeam — Backup & Replication v12/v13
- 11. Open-source backup tools
- 12. Database backup: PostgreSQL and MySQL ecosystems
- 13. Ransomware-proof backup design (2026)
- 14. RTO/RPO guidance and restore testing
- 15. Comparison matrices
- 16. Gaps, conflicts and unverified claims
- 17. Glossary
- 18. Source index (verbatim URLs)
- 19. Proxmox VE operations deep-dive (supplement)
- 20. Proxmox Backup Server operations deep-dive (supplement)
- 21. Hypervisor alternatives deep-dive (supplement)
- 22. Veeam deep-dive (supplement)
- 23. Open-source backup tools deep-dive (supplement)
- 24. Database backup operations deep-dive (supplement)
- 25. Proxmox VE security hardening checklist (supplement)
- 26. Backup scheduling and retention policy examples
- 27. Capacity planning formulas and rules of thumb
- 28. Common errors and troubleshooting
- 29. Training, certification and community resources
- 30. Proxmox VE 9.x detailed changelog notes (supplement)
- 31. Platform comparison: decision factors in detail
- 32. Veeam v13 feature inventory (itemized)
- 33. Monitoring and alerting stack for Proxmox estates
- 34. Ransomware case-study lessons (patterns, 2024–2026)
- 35. Backup encryption key management
- 36. Command cheat-sheet (quick reference)
- 37. Version timeline (2024–2026)
- 38. Decision trees (quick selectors)
- 39. Open questions for future waves

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
