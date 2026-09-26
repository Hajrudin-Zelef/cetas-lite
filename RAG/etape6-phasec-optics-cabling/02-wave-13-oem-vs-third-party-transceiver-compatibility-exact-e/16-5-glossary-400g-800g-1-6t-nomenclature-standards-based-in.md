---
id: etape6-phasec-optics-cabling/02-wave-13-oem-vs-third-party-transceiver-compatibility-exact-e/16-5-glossary-400g-800g-1-6t-nomenclature-standards-based-in
title: "16.5 Glossary — 400G/800G/1.6T nomenclature [standards-based/independent]"
domain: wave-13-oem-vs-third-party-transceiver-compatibility-exact-e
role: deep-dive
task: regulation
actors: ["Nvidia", "United States"]
dates: []
keywords: ["cpo", "dsp", "ethernet", "lpo", "npo", "nvidia", "optics", "serdes"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2788, 2830]
section: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior"
sha256: 8417af3b0dbf57bd80618960477982ba5ae52b0483956d5dfb487540821b055c
---

# 16.5 Glossary — 400G/800G/1.6T nomenclature [standards-based/independent]

| Product | Retailer | Price | Stock note | URL |
|---|---|---|---|---|
| 800G OSFP DR8 500m SiPh | QSFPTEK | US$699 | 3 pcs, dated 13 Sep 2026 | https://www.qsfptek.com/product/103578.html |
| 800G OSFP DR8 100m (NVIDIA-coded) | QSFPTEK | US$799 | 19 pcs, dated 9 Sep 2026 | https://www.qsfptek.com/product/103592.html |
| 800G OSFP SR8 (NVIDIA-coded) | FS.com | US$879 | — | https://www.fs.com/c/osfp-200-400-800g-4089 |
| 800G QSFP-DD DR8 500m | QSFPTEK | US$1,021.90 | 2 pcs, dated 14 Sep 2026, 137 sold | https://www.qsfptek.com/product/102570.html |
| 800G OSFP DR8 (Arista-coded) | FS.com | US$1,319 | ~1K sold | https://www.fs.com/c/osfp-200-400-800g-4089 |
| 800G QSFP-DD DR8 (Cisco-coded) | FS.com | US$1,429 | — | https://www.fs.com/c/osfp-200-400-800g-4089 |
| 800G QSFP-DD800 SR8 D.858HG.005.MP | FlexOptix | EUR 1,599.26 | 2 pcs | https://www.flexoptix.net/en/transceiver |
| 800G QSFP-DD800 2xDR4 D.134HG2.05 | FlexOptix | EUR 1,604.90 | 3 pcs, +10 ~Dec 9 2026 | https://www.flexoptix.net/en/transceiver |
| 800G QSFP-DD DR8 (twin-port IB) | ATGBICS/DigiKey mp | US$1,641 | 3,255 in stock | digikey.com marketplace listing |
| 800G OSFP DR8 (IB) | ATGBICS direct | £1,959 | — | https://www.atgbics.com/800g-osfp-dr8-transceiver-500m-osfp-800g-dr8-flat-top-optical-infiniband-at |
| 800G OSFP DR8 (Juniper eq., Axiom) | SHI | US$2,866 | MSRP US$3,128 | https://www.shi.com/product/50881435/ |
| 800G OSFP DR8 | ATGBICS/DigiKey mp | US$3,067 | 669 in stock, ~14-day shipment | https://www.digikey.com/en/products/detail/atgbics/OSFP-800G-DR8-FLT-MSA-AT/25582709 |
| 800G OSFP DR8 | Approved Networks | US$2,275 | list price, article 17 Jul 2024; current price not found | https://convergedigest.com/approved-networks-adds-osfp-800g-dr8-transceiver/?amp=1 [secondary] |

**Reference anchors (other speeds, same date):** 100G QSFP28 SR4 — FS.com US$99.00 (188.6K sold), QSFPTEK US$39.90, FlexOptix Q.851HG.02 from EUR 73.18 (12,449 pcs) [vendor-reported/official]; 400G QSFP-DD DR4 — FS.com Arista-coded US$549 / Cisco-coded US$749, QSFPTEK Arista-compatible US$449.90, PRO-OPTICS generic SiPh US$699.00, LightOptics US$519.50, FlexOptix D.134HG.05 from EUR 490.60 (1,657 pcs) [vendor-reported/official]; NVIDIA ConnectX-8 1-port 800G OSFP NIC — FS.com US$2,519 [vendor-reported] (https://www.FS.com/c/nvidia-ethernet-nics-4014).
**Outlier excluded from ladder:** Westbury Photonics £182 "Sold out" with inconsistent specs ("500m across MMF") — [unverified], not comparable.

### 16.5 Glossary — 400G/800G/1.6T nomenclature [standards-based/independent]

- **DR4** — 500m single-mode, 4 parallel fibers per direction (MPO-12), 1310nm; 400G-DR4 = 4×100G PAM4, 800G-DR8 = 8×100G PAM4 [standards-based; see Waves 1–2].
- **DR8 / DR8+** — 800G parallel-fiber variants; "+" denotes extended reach options per vendor [vendor-reported].
- **FR4 / FR8** — 2km duplex-LC/LC variants (CWDM4 grid for FR4); 800G-FR8 = 8×100G [standards-based].
- **SR4 / SR8** — short-reach multimode (850nm VCSEL), 50–100m; 400G-SR4 (MPO-12), 400G-SR8 (MPO-16), 800G-SR8 (MPO-16) [standards-based].
- **LR4 / LR8 / ER4** — 10km / 10km / 40km duplex single-mode variants [standards-based].
- **XDR8** — extended-dynamic-range 800G variant (Arista-coded QSFP-DD XDR8 listed at FS.com US$1,999 — Wave 13 §7.3) [vendor-reported].
- **VSR4** — very-short-reach 400G variant (ProLabs Arista QDD-400G-VSR4-C) [official] (prolabs.com).
- **SR4.2 / BiDi** — bidirectional variants using two wavelengths over duplex fiber (e.g., 850/900nm) [standards-based; see Wave 1 QSFP-40/100-SRBD].
- **SWDM4** — shortwave WDM, 4 wavelengths over duplex OM3/OM4 [standards-based].
- **SiPh** — silicon photonics: photonic integrated circuits in silicon; appears in retail SKUs (QSFPTEK, FS QDD-DR4-400G-Si, PRO-OPTICS) [vendor-reported].
- **LPO** — linear-drive pluggable optics (no DSP); **DSP** — DSP-based retimed optics; **CPO/NPO** — co-packaged / near-packaged optics (see Wave 3 for full treatment) [standards-based].
- **OSFP / QSFP-DD / QSFP112 / OSFP-XD** — form factors; 800G uses OSFP and QSFP-DD; 1.6T targets OSFP-XD (evidence: not found — Wave 15 §15.6) [standards-based].
- **224G SerDes** — 224 Gb/s electrical lane signaling underpinning 1.6T (8×200G) per IEEE P802.3dj track (Wave 15 §15.6) [independent].
- **AEC** — active electrical cable (retimed copper); **ACC** — active copper cable; **AOC** — active optical cable (see Wave 6 for full treatment) [standards-based].
- **CMIS** — Common Management Interface Specification for 400G+ modules (10Gtek cites CMIS 5.0 for 800G AEC) [official] (store.10gtek.com).
- **MPO-12 / MPO-16 / MPO-24** — multi-fiber push-on connectors; 800G SR8 uses MPO-16/APC (FS QDD-400G-SR8 lists MPO-16/APC; QSFPTEK 800G DR8 dual-MPO-12 patterns) [vendor-reported].
- **APC vs UPC** — angled vs ultra physical contact polish; APC mitigates back-reflection that destabilizes laser transmitters (Wave 13 §5) [independent].

*End of Wave 16. Phase C complete: 16 waves, append-only, earlier sections untouched.*

---

