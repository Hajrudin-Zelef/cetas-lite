# NOTES — corpus `etape6-phasea-vendors-dc`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

La source contient 2 titres H1 : un dossier par document H1 (un entête éventuel devient `00-front-matter`).

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 6 — Phase A: Enterprise Data-Center Switching Vendors` — 2313 lignes, 38 chunks.

- Dell Networking + NVIDIA Networking (Mellanox heritage) + Aruba (HPE) + Cisco Meraki
- Master timeline (2026)
- §1 Dell Networking
- §2 NVIDIA Networking (Mellanox heritage)
- §3 Aruba (HPE Networking)
- §4 Cisco Meraki
- §5 Cross-vendor comparison (2026)
- §6 Market context (2026)
- §7 Verification log (open items)
- §8 Sources (verbatim URLs)
- Supplementary / Complementary Research Pass — 2026-09-22
- Supplementary / Complementary Research Pass #2 — 2026-09-22
- Supplementary / Complementary Research Pass — round 2 (2026-09-22)
- Supplementary / Complementary Research Pass — round 3a: Cisco FY2026 results, NVLink-6 scale-up, HPE Q3 FY2026, Meraki SKUs/API, Dell N3200 OS matrix (2026-09-22)
- Supplementary / Complementary Research Pass — round 3 (2026-09-22)
- Supplementary / Complementary Research Pass — round 3b (2026-09-22): Meraki pricing & licensing tiers, Cisco C9350/C9610, IOS XE 17.18, BlueField-4, Quantum-X800 detail, Cumulus 5.18
- Supplementary / Complementary Research Pass — round 4 (2026-09-22)
- Supplementary / Complementary Research Pass — round 4 (2026-09-22)
- Supplementary / Complementary Research Pass — round 4 (2026-09-22)
- Supplementary / Complementary Research Pass #3 — 2026-09-22
- Supplementary / Complementary Research Pass — round 4 (2026-09-22)
- Supplementary / Complementary Research Pass — round 4 (2026-09-22)
- Supplementary / Complementary Research Pass — round 4 (2026-09-22)
- Supplementary / Complementary Research Pass — round 4 addendum (2026-09-22, "QFX5140 + HPE Helios fabric")
- R5-A. Cisco C9550 Series Fixed Core Smart Switches — full datasheet specifications (NEW detail)
- R5-B. Dell PowerSwitch Z9964/Z9864 — official radix and two-tier engineering detail (NEW detail)
- R5-C. Meraki MS licensing — official SKU matrix and tier definitions (NEW detail)
- R5-D. NVIDIA networking — Spectrum-X platform composition and roadmap naming note (NEW detail)
- 15. Addendum — NVIDIA networking at xAI Colossus, ConnectX/BlueField roadmap, Meraki price points (September 22, 2026)
- Supplementary / Complementary Research Pass — round 4 (2026-09-22)
- 16. Addendum — Meraki Cloud Management IOS XE 17.18.x release timeline, Aruba CX 10040 street pricing, Broadcom Trident5-X12 silicon context, Dell SFM for SONiC official spec (September 22, 2026)
- Supplementary / Complementary Research Pass — round 5 (2026-09-22)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
