---
id: etape6-phasec-optics-cabling/00-front-matter/3-3-who-ships-what-as-of-sept-2026
title: "3.3 Who ships what (as of Sept 2026)"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "China", "Nvidia"]
dates: ["2025-03", "2026-03"]
keywords: ["asic", "benchmark", "cpo", "dpo", "dsp", "ethernet", "gpu", "latency", "lpo", "nvidia", "optics", "serdes"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [393, 413]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: b4148b573005795326f676a8d338418643996d85e0fe95ad167a29715086988a
---

# 3.3 Who ships what (as of Sept 2026)

#### 3.2.2 Power consumption (800G-class)
Sourced numbers (do not mix generations — flags noted):
- **DSP-based 800G**: 14–18 W baseline [secondary: https://ascentoptics.com/blog/800g-power-consumption/]; 12–15 W per 800G port [secondary: https://momoview.com/blog/en/posts/co-packaged-optics-cpo-silicon-photonics-industry-analysis-ai-interconnect-bottleneck-2026/]; 8–12 W [secondary: https://edgeoptic.com/blog/lpo-vs-dsp — conflicts with above; weaker source]. Broadcom BCM85812 5nm DSP: drives 800G SMF module power to "sub 11 W" and MMF to "sub 10 W" [official: https://www.broadcom.com/products/ethernet-connectivity/phy-and-poe/optical/bcm85812 and https://www.photonicsonline.com/doc/broadcom-introduces-industry-s-first-nm-g-lane-optical-pam-dsp-phy-tia-laser-driver-0001]. One Chinese market piece claims 800G multimode DSP module ">13 W" vs LPO "<4 W" — [unverified, weak source: https://finance.biggo.com/news/uei-yp4BOLsyMWM01Hme].
- **LPO 800G**: 5–8 W (40–50% reduction) [secondary: https://ascentoptics.com/blog/800g-power-consumption/]; 7–9 W (~60% of DSP baseline) [secondary: momoview]; 2.5–8 W across 400G/800G [secondary: edgeoptic — wide range, weak]; FS 800G DR8 OSFP LPO: max 8.5 W, "~50% lower than 800G DSP-based modules" [vendor-reported: https://aithority.com/machine-learning/fs-launches-800g-lpo-module-a-power-efficiency-and-latency-optimized-solution-for-ai-hpc-data-centers/]; Eoptolink 800G LPO: <8.0 W [vendor-reported: OIF demo PDF]; AICPLIGHT 800G DR8 LPO: 8 W, 50% below "traditional DPO" [vendor-reported: medium.com/@aicplight888].
- **Extreme vendor claim**: Adtran LiteWave800 (800G DR8 LPO, OSFP, single-mode VCSELs + in-house electronics) achieves "just 1 pJ/bit, or about 0.8 W" — claimed 12–18× lower than typical DSP-based optics, per Adtran March 2026 announcement [vendor-reported, marketing: https://ascentoptics.com/blog/800g-power-consumption/ — flagged as marketing; the 0.8 W figure conflicts with the 5–8 W typical LPO range and is not independently verified].
- **1.6T retimed vs LPO**: Semtech blog states retimed modules "23–25 W for a 1.6T DR-8" (text says "200G links" — ambiguous, likely means 200G/lane-class links, i.e. 1.6T; flagged as ambiguous), expected ~20 W with next-gen DSPs; LPO target ~10 W; RTLR/LRO ~16 W today. System-level: at 512 ports (Tomahawk-6 generation), LPO saves ~500 W at module level, up to 1 kW with cooling; double at 1024 ports [vendor-reported: https://blog.semtech.com/ai-date-center-basics-what-is-linear-pluggable-optics-lpo].
- **NVIDIA claim**: Spectrum-X Ethernet Photonics cuts power per 1.6T port from 25 W (traditional pluggable) to 9 W [secondary/vendor-adjacent: https://markets.financialcontent.com/ibtimes/article/tokenring-2026-1-20-nvidias-spectrum-x-ethernet-photonics-powering-the-million-gpu-era-with-light-speed-efficiency]. Per momoview's compilation: DSP→LPO cuts back-end network power ~16%; full CPO transition cuts transceiver power ~84% — figures attributed to NVIDIA [secondary].
- **Micas/Broadcom CPO benchmark (March 2025)**: fully-populated CPO switch showed >40% system power savings vs standard pluggable transceivers and >25% vs LPO-populated systems [vendor-reported: https://www.globenewswire.com/news-release/2025/03/17/3043828/0/en/Micas-Networks-Announces-Industry-s-First-51-2T-Co-Packaged-Optics-Network-Switch-System-Now-in-Volume-Production.html].
- **Cisco G300 + LPO**: Cisco would not disclose its 800G LPO module's power, but said G300 systems with LPO give ~30% reduction in switch power; pluggables "usually 10–20 W," LPO ~50% reduction [vendor-reported: https://www.theregister.com/on-prem/2026/02/10/cisco-unveils-1024t-silicon-one-g300-switch-chip/4836608].

#### 3.2.3 Thermal / system benefits
- Lower module power = lower thermal load → reduced switch/server cooling demand [vendor-reported: FS/aithority link above]; LPO cited as relieving thermal pressure in AI clusters [multiple secondary].

#### 3.2.4 Link budget / interoperability concerns with LPO
- LPO requires a host ASIC with advanced linear-PAM4 SerDes equalization; standard switch requirement listed as "Linear PAM4 support" vs "Standard" for DSP modules [secondary: https://edgeoptic.com/blog/lpo-vs-dsp].
- Reach is short: up to ~500 m (DR8) to 2 km (2×FR4) per Eoptolink spec [official: OIF demo PDF]; edgeoptic cites "up to 500 m or 2 km (ideally)" vs DSP ">2 km" [secondary]. One Medium piece claims LPO limited to ~50 m intra-rack — [unverified, conflicts with Eoptolink's official 500 m–2 km spec; flagged as unreliable].
- Interoperability: "Limited interoperability — can't be easily interconnected with DSP-based optics due to different signal handling approach" [secondary: edgeoptic]. Eoptolink's LPO supports triple-rate operation 106.25G/53.125G PAM4/25.78G NRZ [official: OIF demo PDF].
- Channel quality demands: LPO needs low-loss, pre-tested fiber plant [secondary]. Semtech notes high-loss channels or high-reflection environments may still favor retimed/RTLR [vendor-reported].

### 3.3 Who ships what (as of Sept 2026)

