# NOTES — corpus `etape9-phaseb-form-factors-interfaces`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-form-factors-interfaces` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)` — 751 lignes, 12 chunks.

- 1. Why form factors still matter in 2026
- 2. M.2 — client-born, datacenter-boot
- 3. U.2 (SFF-8639) and U.3 (SFF-TA-1001) — the 2.5" NVMe era
- 4. EDSFF — the NVMe-native family
- 5. Other physical formats in the 2026 datacenter
- 6. Form-factor comparison matrix (2026 enterprise)
- 7. NVMe specification family: 2.0 → 2.4 (August 2026)
- 8. NVMe I/O command sets in 2026
- 9. NVMe over Fabrics — transports in 2026
- 10. Advanced NVMe features (hardware-relevant)
- 11. PCIe generations: the lane underneath everything
- 12. SAS and SATA in 2026 — not dead, repositioned
- 13. Backplanes, expanders, and enclosure management
- 14. Storage cabling and connectors
- 15. Server platform integration (2026)
- 16. Decision guide — picking form factor & interface
- 17. Gaps, conflicts, and unverified claims register
- 18. Glossary
- 19. Source index (verbatim URLs)
- 20. Interface timeline (1990 → 2026)
- 21. Drive-type ↔ bay-type interoperability matrix
- 22. PCIe lane budgeting example (Gen5 2P server)
- 23. Retimers, redrivers, and cable reach
- 24. CXL note (storage-adjacent)
- 25. Storage software stack checklist (hardware-adjacent)
- 26. Capacity points and endurance classes (2026)
- 27. Security features (drive-level)
- 28. Firmware and lifecycle management
- 29. Thermal and airflow notes
- 30. One-page cheat sheet (2026 enterprise storage hardware)
- 31. SKUs referenced in this file (quick table)
- 32. Document revision note

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
