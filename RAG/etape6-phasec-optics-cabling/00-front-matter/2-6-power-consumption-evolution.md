---
id: etape6-phasec-optics-cabling/00-front-matter/2-6-power-consumption-evolution
title: "2.6 Power consumption evolution"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "China", "Cohere", "Nvidia", "TSMC"]
dates: ["2025-03", "2026-03"]
keywords: ["2nm", "3nm", "backlog", "cost", "cpo", "dsp", "lpo", "npo", "nvidia", "optics", "wavelength"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [305, 329]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: a3e61222263df3abd372345783795fac687f92e5c2a8d73d2a61386961d4acce
---

# 2.6 Power consumption evolution

- **CWDM4** — 4 center wavelengths **1271, 1291, 1311, 1331 nm** (20 nm spacing; members of ITU-T G.694.2 CWDM grid which spans 1271–1611 nm in 18 channels) [independent — ITU-T G.694.2 via Corning app note]. Used by 100G CWDM4 (2014 CWDM4 MSA), 400G-FR4/LR4, 800G-2FR4/2LR4 (2 groups). 100G QSFP28 CWDM4: 4×25G NRZ, uncooled DFB, <3.5 W [vendor-reported — DawnrayTech]. 400G FR4: 4×100G PAM4, uncooled DML/DFB [vendor-reported]. Host FEC required for CWDM4 links [vendor-reported]. Sources: https://www.fibermall.com/blog/100g-qsfp28-lr4-psm4-and-cwdm4-modules.htm ; https://www.corning.com/catalog/coc/documents/application-engineering-notes/AEN133.pdf
- **LAN-WDM** — 4 dense-ish wavelengths near 1310 nm used by IEEE 400GBASE-LR4/100GBASE-LR4: **1295, 1300, 1304, 1309 nm** nominal (precisely 1295.56 / 1300.05 / 1304.58 / 1309.14 nm per IEEE 802.3ba) [independent — IEEE 802.3ba / widely cited]; cooled EML transmitters; LR4/ER4/FR8/LR8 applications [vendor-reported — FS/L-P].
- **SWDM4 / SR4.2** — 4 short wavelengths over MMF: **850, 880, 910, 940 nm** (SWDM4); 400G SR4.2 uses 850/910 nm two-λ variant, 150 m on OM5 [independent — widely cited; FS table lists 850/910 nm for SR4.2]. OM5 wideband MMF optimized for 850–950 nm [independent].
- **BiDi** — single-fiber bidirectional: 850 nm / 900 nm pair for MMF BiDi (e.g., 100G-SR1.2) and 1270/1330 nm for SMF BiDi classes — **[unverified: not confirmed by a source in this pass; commonly cited in industry literature, flag as unverified]**.
- **850 nm VCSEL vs 1310 nm**: VCSELs = MMF, short reach (≤100 m), lowest cost, used in SR8/SR4/SR4.2 [vendor-reported]; 1310 nm = SMF parallel (DR4/DR8/XDR8), DFB/EML [vendor-reported]; wavelength-MUXed 1310-region (CWDM4/LAN-WDM) = duplex-fiber 2–10 km [vendor-reported].
- **Silicon photonics (SiPh)**: InnoLight 1.6T SiPh modules in volume ramp (2026) [secondary]; Coherent SiPh PIC + pure-silicon MZM for 400G/lane (3.2T path) [official]; NVIDIA TSMC COUPE SiPh for CPO switches [official]; Marvell 200G/lane SiPh CPO platform [official]; Yole 2026 report segments module forecasts by optical technology: VCSEL, DML, EML, SiPh, plus emerging TFLN/BTO/organic modulators [secondary — optics.org].

### 2.6 Power consumption evolution

- **400G-DR4 (DSP-based)**: ≤9–10 W (FS QSFP-DD) [vendor-reported]; 8–10 W (Ascent) [vendor-reported]; "up to ~12 W vendor-dependent" (network-switch.com) [vendor-reported] — **[flagged: ranges differ by vendor/generation]**.
- **800G-DR8 (DSP-based)**: ≤16.5 W QSFP-DD; ≤16.5 W OSFP (DSP); OSFP quoted "~15 W" [vendor-reported — FS, expertini]; 800G 2FR4 ≤17 W; 800G 2LR4 ≤18 W [vendor-reported — FS].
- **LPO (linear pluggable optics)**: 800G DR8-LPO **≤8.5 W** (no DSP) — roughly half the DSP module power [vendor-reported — FS].
- **Trend**: per-module power rises with lane rate (~9–10 W @ 400G → ~16–17 W @ 800G DSP), while power-per-bit improves (25 mW/G → ~20 mW/G). DSP chip itself is ~6–8 W per 800G module [secondary — XenoSpectrum]; removing it (LPO/CPO/NPO) is the main lever: NPO ~5–7 W/800G, CPO ~4–5 W/800G per Broadcom [secondary — XenoSpectrum].
- **Form-factor thermal ceilings (upper thresholds)**:
  - QSFP-DD: 7–12 W (early/400G era) [secondary]; QSFP-DD800 vendor modules rated to 18 W [vendor-reported — FS 800G guide]; QSFP-DD1600 typical 20–28 W [secondary — fibermall].
  - OSFP: 12–15 W (400G era) [secondary]; 800G OSFP vendor modules rated to 18 W [vendor-reported — FS]; OSFP-XD designed to ~40 W [secondary].
  - **[flagged conflict: different sources quote different ceiling numbers because they refer to different MSA revisions/generations — 12 W (QSFP-DD 400G) vs 18 W (QSFP-DD800) vs 20–28 W (QSFP-DD1600)]**.
- NVIDIA CPO claim: 1.6T port power 25 W (pluggable) → 9 W (CPO), ~5× efficiency gain [official — NVIDIA].

### 2.7 Key vendors & positioning

**DSP makers**: Broadcom (Sian3 3nm, Sian2M 5nm for 800G/1.6T) [secondary]; Marvell (Nova/Nova 2 1.6T, Ara 3nm 1.6T platform, Libra 2nm 800G coherent, 38 ECOC 2026 demos) [official]; Inphi (acquired by Marvell 2021 — Inphi DSPs cited in FS 400G tables) [vendor-reported]; NVIDIA does not sell standalone DSPs but drives CPO/SiPh ecosystem [official].
**Module makers**: InnoLight (largest share ~27% [secondary], 1.6T SiPh ramping, backlog through 2026) [secondary]; Coherent (vertically integrated — VCSEL/InP/SiPh/lasers, OFC+ECOC 2026 demos) [official]; Eoptolink (1.6T DR4 400G-λ demo OFC 2026, 800G in volume) [secondary]; AOI (800G ramp, 500k/mo 800G+1.6T target end-2026) [secondary]; Accelink, Hisense Broadband, Source Photonics, Sumitomo — major players per market-report vendor lists but **no 2026-specific announcements captured in this pass [gap]**; Lumentum (lasers, named as beneficiary if China-sourced modules restricted) [secondary]; Cisco (400G ZR/ZR+ coherent QSFP-DD; market-report vendor) [secondary].
**Notable 2026 announcements** (dated): Marvell ECOC 2nm demos (21 Sept 2026) [official]; Broadcom OFC NPO 3.2T VCSEL (March 2026) [secondary]; Coherent OFC 1.6T/3.2T/XPO MSA (17 March 2026) [official]; Coherent ECOC 3.2T/6.4T-NPO (Sept 2026) [secondary]; Eoptolink OFC 1.6T DR4 (March 2026) [secondary]; NVIDIA GTC CPO switches (19 March 2025, shipping into 2026) [official]; InnoLight H1 2026 results (Aug 2026) [secondary]; FS 800G LPO module launch (28 Dec 2025) [secondary].

