# NOTES — corpus `etape9-phasec-memory`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-memory` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 9 — Phase C: Server & Accelerator Memory Hardware` — 753 lignes, 12 chunks.

- 0. Scope, method and provenance
- 1. Server DRAM landscape, 2026
- 2. DDR5 deep-dive: speeds, architecture, modules
- 3. MRDIMM: multiplexed-rank DIMMs and the road to 12,800 MT/s
- 4. Server DIMM taxonomy: UDIMM vs RDIMM vs LRDIMM vs 3DS
- 5. Platform memory architectures, 2026
- 6. DDR6: status, architecture, timeline
- 7. LPDDR5X and LPDDR6
- 8. CAMM2 and LPCAMM2 form factors
- 9. ECC: on-die vs side-band
- 10. Advanced ECC and RAS: Chipkill, SDDC, ADDDC, mirroring
- 11. HBM generations: the specification table
- 12. HBM2e in 2026: legacy but still deployed
- 13. HBM3 / HBM3e: the 2023–2026 workhorse
- 14. HBM4: the 2026 transition
- 15. NVIDIA accelerators: memory map
- 16. AMD accelerators: memory map
- 17. Intel Gaudi 3: the HBM2e outlier
- 18. Other accelerators and custom silicon
- 19. GDDR6 / GDDR7
- 20. HBM supply chain and market, 2026
- 21. DRAM pricing signals, 2026
- 22. CXL: memory expansion and pooling
- 23. CXL as the Optane successor (positioning)
- 24. Intel Optane: discontinued
- 25. NVDIMM and other persistent-memory options, 2026
- 26. AI memory sizing: inference
- 27. AI memory sizing: training and fine-tuning
- 28. The GPU memory hierarchy (why HBM exists)
- 29. VRAM-per-model quick reference (FP16 weights + headroom)
- 30. Decision guide: server DRAM in 2026
- 31. HBM power, thermals and packaging
- 32. The base-die revolution (2025–2026)
- 33. Server memory form factors, consolidated
- 33B. DRAM process nodes, 2026
- 33C. Memory technology comparison matrix
- 33D. Server memory power and efficiency
- 34. What to watch after 2026-09-22
- 35. Conflicts, gaps and non-comparable figures
- 36. Glossary
- 37. Source index (verbatim URLs)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
