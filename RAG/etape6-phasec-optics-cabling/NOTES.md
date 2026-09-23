# NOTES — corpus `etape6-phasec-optics-cabling`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

La source contient 3 titres H1 : un dossier par document H1 (un entête éventuel devient `00-front-matter`).

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure` — 3012 lignes, 52 chunks.

- Wave 1 — FS.com catalog & pricing (research date 2026-09-22)
- Wave 2 — Optical transceiver landscape & standards (research date 2026-09-22)
- Wave 3 — DSP vs LPO vs CPO/NPO (research date 2026-09-22)
- Wave 4 — Fiber-optic infrastructure (research date 2026-09-22)
- Wave 5 — Copper structured cabling: Cat6A/Cat8 (research date 2026-09-22)
- Wave 6 — DAC vs AOC vs AEC: Direct-Attach Cable Technologies
- Wave 7 — FS.com Optics Catalog & Pricing (Second Pass: SKUs, DAC/AOC Gaps, Coding, Warranty)
- Wave 8 — Fiber-Optic Infrastructure: OS2/OM3/OM4/OM5, MPO/MTP, Polarity, Trunks, Cassettes, Testing
- Wave 9 — Copper Structured Cabling: Categories, OOB Roles, Pricing, PoE, Testing, Trends
- Wave 10 — The 800G Switching Ecosystem (as of September 2026)
- Wave 11 — 800G NICs/DPUs and AI-Cluster Deployments
- 0. Evidence legend and method notes
- 1. Coding mechanics — EEPROM fields, checksums, standards, OUI validation, recoding
- 2. Programmers and coding tools
- 3. Vendor ecosystem
- 4. Warranty models — matrix, conflicts, exclusions, MTBF
- 5. OEM lock-in behavior and support implications
- 6. Representative price gaps (with dates and comparability flags)
- 7. Legal and gray areas
- 8. Enterprise/hyperscale practice
- 9. Conflicts, unverified claims, and gaps register
- 10. Source ledger (key URLs, by category)
- 11. Coverage check against the eight requested areas
- How to read this report
- 1. Exact error messages (verbatim quotes)
- 2. EEPROM / identification mechanics
- 3. OEM support / compatibility policy
- 4. Third-party vendors: coding, testing, warranty claims
- 5. DOM/DDM failure diagnosis and failure modes
- 6. FEC
- 7. Prices (dated snapshots, observed 2026-09-22)
- 8. Practical compatibility notes
- 9. Explicit verification gaps — DO NOT quote as verified
- Wave 14 — Vendor compatibility matrix: third-party coded optics vs OEM platforms (research date 2026-09-22)
- Wave 15 — 800G ecosystem readiness, silicon photonics, 1.6T signals, availability, alternatives & consolidated gaps (research date 2026-09-22)
- Wave 16 — DAC/AOC coding, lab validation playbook, buyer's decision framework, 800G price ladder & glossary (research date 2026-09-22)
- Wave 17 — Reference tables: coded-SKU examples, OEM part numbers, 800G catalog detail, retailer price matrices, programming workflow (research date 2026-09-22)
- Wave 18 — Phase C coverage verification & provenance audit (research date 2026-09-22)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
