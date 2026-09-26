---
id: etape6-phasec-optics-cabling/00-front-matter/2-8-standards-bodies-status-sept-2026
title: "2.8 Standards bodies status (Sept 2026)"
domain: front-matter
role: reference
task: reference
actors: ["China", "Cohere", "Lambda"]
dates: ["2024-03", "2024-05", "2024-06", "2026-03", "2026-06", "2026-07", "2026-09-22"]
keywords: ["asic", "cost", "cpo", "dsp", "energy", "ethernet", "lpo", "npo", "optics", "research", "revenue", "training"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [330, 367]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 631e504c7e80efe010a1c1653e0dd22ba0320e66163eca221e8191acbae2b188
---

# 2.8 Standards bodies status (Sept 2026)

### 2.8 Standards bodies status (Sept 2026)

- **IEEE P802.3dj** (200/400/800G + 1.6T, 200G/lane): in-progress; PAR projected RevCom Mar 2026 [official]; Nokia expected ratification Oct 2026 (2025 deck) [vendor-reported]; as of 22 Sept 2026 still unratified, with ECOC 2026 interoperability showcases running [secondary]. **Tag final date [unverified]**.
- **IEEE 802.3df-2024**: published 15 March 2024 [official] — current 800G baseline.
- **IEEE 802.3 400 Gb/s-per-lane Signaling Study Group**: chartered **13 March 2026** for 400G/lane electrical + SMF optics to 500 m (maps to 3.2T-class Ethernet timing) [independent — IEEE 802.3 site via GitHub mirror notes]. Also **P802.3ds** (200 Gb/s per wavelength MMF PHYs Task Force) active [independent]; IEEE "802.3 Ethernet Interconnect for AI" assessment started Jan (2026) [secondary].
- **OIF CEI-224G**: **CEI-224G-Linear project launched at OIF Q2 2024 meeting (Athens, 7–9 May 2024)** [official — OIF/BusinessWire] — supports 224G full-linear optical modules (LPO/CPO/NPO) for Ethernet/UEC/AI/ML; TP1/TP1a, TP4/TP4a electrical specs; built on CEI-112G-Linear methodology. Also CMIS smart-module + out-of-band link-training white papers and EEI vendor-requirements doc released at same meeting [official]. Sources: https://www.businesswire.com/news/home/20240613908066/en/OIF-Q2-Technical-and-MAE-Committees-Meeting-Wraps-with-CEI-224G-Linear-Project-Launch-New-CMIS-White-Papers-and-Requirements-for-Energy-Efficient-Interfaces ; https://convergedigest.com/oif-advances-cei-224g-linear-project/?amp=1
- **CPO-related**: OCP CPO workstream — 300-page blueprint white paper published; 20-company coalition; **Q4 2026 deadline for first CPO specifications**; OIF published "Implementation Agreement for a 3.2 Tb/s Co-Packaged (CPO) Module" (referenced June 2026) [secondary — techtimes Aug 2026; NTMM PDF]; Coherent FlexConnect (detachable FAU for CPO) introduced ECOC 2026 [secondary]; Yole: large-scale CPO deployment realistically 2028–2030, not 2026–27 [secondary — via investment research citing Yole]. NPO emerging as interim architecture (field-replaceable, no DSP) [secondary — XenoSpectrum].
- **MSA revisions**: QSFP-DD HW Spec Rev 7.1 (June 2024, QSFP-DD1600) [secondary]; OSFP MSA Rev 5.1 (50/100/200G-class lanes) [secondary]; SFP-DD MSA (2017) [secondary]; CMIS 5.x (QSFP-DD800/1600 modules; CMIS 5.1 cited in DAC datasheet) [vendor-reported]; 100G Lambda MSA (single-lambda DR/FR/LR) [vendor-reported]; CWDM4 MSA 2014 [secondary].

### 2.9 Price trend data (2026, market reports)

- **LightCounting — "Ethernet Optics" report, March 2026** [official LightCounting site]: analyzes AI-cluster optics; 2022–2025 historical shipments + **2026–2031 forecasts for units, prices, and sales** across 100+ product categories (100/200/400/800G, 1.6T, 3.2T, retimed, LPO/LRO, CPO/NPO by reach/form factor) for Cloud/Enterprise/Telecom [official]. LightCounting (Apr 2024): **40% sales growth 2024, >20% 2025, double-digit growth 2026–2027** for optical Ethernet transceivers, soft-landing risk after [secondary — IEEE ComSoc]. Source: https://www.lightcounting.com/document/march-2026-ethernet-optics/382/toc ; https://techblog.comsoc.org/2024/04/03/lightcounting-optical-ethernet-transceiver-sales-will-increase-by-40-in-2024/
- **Yole Group — "Optical Transceivers for Datacom and Telecom 2026"**: market "poised to exceed **$110BN by 2031**" [secondary — optics.org]; datacom transceiver market from **~$10B (2021) to forecast $112B (2031), ~35% CAGR**; CPO optical-engine revenue forecast **$0.6B (2026) → $112.1B (2031)** with scale-up ~94% [secondary — Yole via semiconductor-today, Sept 2026]; InP/laser supply constraints flagged [secondary]. Sources: https://optics.org/news/optical-transceiver-market-poised-to-exceed-110bn-by-2031---report ; https://www.semiconductor-today.com/news_items/2026/sep/yole-070926.shtml
- **800G unit economics (2025 base)** [secondary — MarketPublishers, July 2026]: global 800G data-center transceiver market **$8,022M (2025) → $27,421M (2032), 19.3% CAGR 2026–2032**; production ~13.23M units (2025), **average price ~$620/unit (2025)**; industry gross margins 26–44% [secondary — lower-tier source, treat ASP as indicative]. Source: https://pdf.marketpublishers.com/lpinfo/global-800g-data-center-transceiver-market-lp.pdf
- **ASP erosion**: vertically integrated Chinese suppliers compress BoM; "-2.2% drag on CAGR" from ASP erosion cited [secondary — market report via eastmnweeklynews]. 100G CWDM4 costs 30–40% less than LR4 (component basis) [vendor-reported]. LR4 30–50% premium over FR4 (400G) [vendor-reported].
- **GM Insights (2026)**: 400–800 Gbps tier $2,207.3M → $13,460.3M (19.48% CAGR); 800G-and-above tier $907.7M (2025) → $7,691.6M (2035), 23.10% CAGR [secondary]. Source: https://www.gminsights.com/industry-analysis/optical-transceiver-market

### 2.10 Conflicts & gaps (Wave 2, explicit)
1. **[CONFLICT]** QSFP+ release year: 2013 (FS) vs 2012 (QSFPTEK).
2. **[CONFLICT]** 400G-LR4 power: 12–15 W (Ascent) vs ≤9–10 W (FS).
3. **[CONFLICT]** 400G-DR4 power: ≤9 W (FS) vs 8–10 W (Ascent) vs ~12 W (network-switch.com).
4. **[CONFLICT]** 800G-SR8 MMF reach: 30m/50m OM3/OM4 (FS) vs "up to 100 m" (generic guides).
5. **[CONFLICT]** QSFP-DD/OSFP power ceilings (12/15/18/21/28 W) refer to different MSA generations — must cite generation with number.
6. **[UNVERIFIED]** Exact IEEE 802.3dj ratification date; FCC ban on China-sourced transceivers; 400G/lane timeline; BiDi wavelength pairs; 800G-FR8/LR8 per-lane 200G details (inferred from naming, not sourced).
7. **[GAP — not found]** 2026-specific announcements for Hisense Broadband, Accelink, Source Photonics, Sumitomo; detailed Yole/Cignal AI/Dell'Oro 2026 ASP tables (paywalled).
8. **Non-comparable figures**: FS "≤X W" are maximum ratings per vendor SKU, not typical; Ascent "typical" ranges vs MarketPublishers ASP ($620 avg across all 800G types/reaches) cannot be compared to type-specific prices.

---

## Wave 3 — DSP vs LPO vs CPO/NPO (research date 2026-09-22)

### 3.1 Definitions

#### 3.1.1 DSP-based (retimed) modules
- A conventional high-speed optical module contains a Digital Signal Processor (DSP) ASIC that performs signal equalization, retiming, and compensation (clock-data recovery) to counteract attenuation and distortion in long electrical traces [secondary: https://naddod.medium.com/optical-interconnect-technology-analysis-lpo-npo-cpo-bd9b3488fb10].
- In fully retimed pluggables, the module DSP "owns" FFE/DFE/CDR on both sides of the optics [secondary: https://github.com/farrox/short-reach-optics/blob/HEAD/docs/ch5-channel-equalization-ctle-ffe-dfe-and-dsp.md]. In a 400G module, the 7nm DSP consumes ~4W, roughly 50% of module power; DSP BOM cost is ~20–40% of a 400G module [secondary: https://www.lemmymorgan.com/what-is-linear-drive-pluggable-optics/]. **Note: this is an older, low-credibility source; treat figures as [unverified] for current generations.**
- At 800G, each of 8 lanes carries 106.25 Gbps electrical signaling which the module's internal DSP converts to/from the optical domain [secondary: https://roboticsandautomationnews.com/2026/09/22/osfp-modules-the-complete-guide-to-400g-800g-and-1-6t-optical-transceivers-for-ai-and-hyperscale-data-centers/104982/].

