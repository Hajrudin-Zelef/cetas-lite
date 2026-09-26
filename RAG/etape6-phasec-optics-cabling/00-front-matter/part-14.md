---
id: etape6-phasec-optics-cabling/00-front-matter/part-14
title: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure (part 14)"
domain: front-matter
role: reference
task: reference
actors: ["Cohere", "Microsoft", "Nvidia", "TSMC"]
dates: ["2026-02-10", "2026-03", "2026-06", "2026-07-14"]
keywords: ["optics", "backlog", "cpo", "disclosure", "dsp", "ethernet", "gpu", "ipo", "latency", "lpo", "npo", "nvidia"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [420, 439]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: c9b45147c8ceb66dc0be14dd4a1b4916de17120e1a782444f7f10e977267644a
---

# Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure (part 14)

#### 3.3.2 Module vendors shipping LPO
- **Eoptolink**: launched 800G LPO at OFC 2023 (Mar 6, 2023) [official: https://www.eoptolink.com/news?start=15]; family: 800G OSFP/QSFP112-DD DR8, 2×FR4, AOC; demoed industry-first 200G/lambda LPO 800G OSFP DR4 (SiPh PIC, 4×200G) at OFC 2024 [official: https://www.eoptolink.com/news/13-new-products?start=5]. Portfolio also includes LPO and LRO variants for AI/ML [secondary: https://www.ainvest.com/news/eoptolink-technology-pioneering-800g-optical-transceiver-revolution-delivering-explosive-growth-h1-2025-2508/]. 800G "already in high-volume deployment" per company at OFC 2026 [secondary: https://semiconductorinsight.com/report/optical-module-package-market/].
- **InnoLight (Zhongji Innolight)**: "wide range of low power, low latency 400G, 800G, and 1.6T LPO solutions" with SiPh platform; founding member of LPO MSA [official via Alphawave press: silicon.co.uk link above]. XPO and NPO products in "custom development or R&D refinement" (H1 2026 results); 1.6T SiPh modules in volume ramp; order backlog covers all of 2026 [vendor-reported: https://finance.biggo.com/news/2b25909c-bd1d-49fd-8c51-a9c7996700d5]. World's largest optical interconnect supplier by revenue 2021–2025 per CIC data in HK IPO prospectus (21.2% global share 2025; 28.1% high-speed datacom segment) [vendor-reported via prospectus: https://finance.biggo.com/news/9890cf93-e2c0-459a-b6d6-29423c08b793].
- **AOI (Applied Optoelectronics)**: demoed **1.6T OSFP DR8 LPO** and **800G OSFP DR8 LPO in an 800G switch (CEI-112G Linear)** at ECOC 2024 with OIF; also CEI-224G Linear demo [official: https://www.globenewswire.com/news-release/2024/09/24/2952013/0/en/AOI-Leads-Industry-with-1-6T-OSFP-DR8-LPO-Technology-at-ECOC-2024.html]. Also showed co-packaging architecture concepts. AAOI–Microsoft relationship "built on" LPO per one analyst note [secondary: https://github.com/47tzp4ydc9-cmyk/life-os/blob/HEAD/investment-os/narrative/research/silicon-photonics.md — **unverified**]. Note: AAOI first volume 800G only in Q1 2026, "late" vs peers [secondary: equity-watch note — treat as analyst commentary].
- **Coherent**: demoed **800G-DR8 LPO transceiver** (8×100G PAM4, OSFP, vertically integrated optics, "excellent pre-FEC BER without a DSP") at ECOC 2023 [official: https://www.globenewswire.com/en/news-release/2023/10/02/2752425/11543/en/Coherent-to-Demonstrate-Next-Generation-Transceiver-and-Semiconductor-Laser-Technology-for-800G-and-1-6T-Transmission-in-AI-Networks-at-ECOC-2023.html]. ECOC 2026: unveiled **PhotonLink** platform; Coherent 10-K FY2026 names CPO, co-packaged optics, SiPh, and NVIDIA (13×) — strongest Tier-1 CPO disclosure among component vendors; "deep engagements across both CPO and NPO" but no scale-up date quoted [secondary: https://github.com/victorhwn7255/stocks-wiki/blob/HEAD/wiki/themes/CPO-platform-battle.md].
- **Accelink**: demonstrated/sampled **1.6T OSFP224 DR8** transceivers at/around OFC 2026 [secondary: https://tspasemiconductor.substack.com/p/ofc-2026-outlook-ai-data-center-optical]. LPO MSA founding member. No specific 800G LPO product name sourced — **gap flagged**.
- **Hisense**: LPO MSA founding member [official]. No specific 800G LPO product sourced — **gap flagged**.
- **Source Photonics, Sumitomo Electric**: named in task; Sumitomo collaborates with Eoptolink on MCF [secondary: ainvest link above]. No LPO product names sourced — **gap flagged**.
- **FS (FS Inc./FS.com)**: launched **800G OSFP DR8 LPO** (finned top), 8.5 W max, Aug 2025 announcement [vendor-reported: https://aithority.com/machine-learning/fs-launches-800g-lpo-module-a-power-efficiency-and-latency-optimized-solution-for-ai-hpc-data-centers/; https://www.thefastmode.com/technology-solutions/44199-fs-launches-800g-lpo-module-to-power-next-gen-ai-data-centers].
- **AICPLIGHT (third-party)**: 800GBASE-DR8 OSFP LPO, 8 W, CMIS 5.0 / LPO MSA 1.0 / OSFP MSA compliant [vendor-reported: medium.com/@aicplight888].
- **Cisco**: Feb 2026 (Cisco Live EMEA) launched **800G LPO** optics alongside 1.6T pluggables for G300-based N9000/8000 systems — signal processing handled in G300 SerDes [secondary/independent: https://www.theregister.com/on-prem/2026/02/10/cisco-unveils-1024t-silicon-one-g300-switch-chip/4836608; https://www.reuters.com/business/media-telecom/cisco-unveils-new-ai-networking-chip-taking-broadcom-and-nvidia-2026-02-10/].
- **Adtran**: **LiteWave800** 800G DR8 LPO, OSFP, ~1 pJ/bit (~0.8 W claimed) using single-mode VCSELs — March 2026 announcement [vendor-reported via https://ascentoptics.com/blog/800g-power-consumption/].

#### 3.3.3 NVIDIA/Mellanox optics position
- NVIDIA first co-proposed LPO with MACOM in 2022 [secondary: naddod]. NVIDIA's **Spectrum-4 switches + BlueField-3** use in-house 100G PAM4 SerDes with "Direct Drive" optical modules — no DSP or CDR in the module — bringing 400G/800G module power "down to single-digit values from over ten watts" [secondary: https://www.naddod.com/blog/nvidia-spectrum-x-solution-benefits-and-product-components].
- **Spectrum-X Ethernet Photonics**: CPO-based; 200G/lane SerDes; 1.6T port power claimed 25 W → 9 W; SN6800 "super-switch" (4 ASICs, 409.6 Tb/s) [secondary: https://markets.financialcontent.com/ibtimes/article/tokenring-2026-1-20-nvidias-spectrum-x-ethernet-photonics-powering-the-million-gpu-era-with-light-speed-efficiency]. NVIDIA developed its own SiPh engines with micro-ring modulators (MRMs), TSMC fab partner, high-power lasers, detachable fiber connectors; covers Quantum-X (InfiniBand) and Spectrum-X (Ethernet) [secondary: https://www.nextplatform.com/connect/2025/03/18/nvidia-weaves-silicon-photonics-into-infiniband-and-ethernet/1642057].
- Shipping status conflict: NVIDIA CTO Gilad Shainer publicly stated (June 2026) that CPO is "already shipping and ramping H2 2026," disputing SemiAnalysis yield concerns [secondary: https://github.com/47tzp4ydc9-cmyk/life-os/blob/HEAD/investment-os/narrative/research/silicon-photonics.md]. Meanwhile investor-facing synthesis expects large-scale CPO only 2028–2030 (Yole) and LightCounting: "another two years before CPO is really shipping in volumes" (Sept 2026) [independent: https://www.eetimes.com/ai-demand-reshapes-optical-connectivity-and-photonics-roadmaps/]. **Flag: NVIDIA shipping claims vs analyst volume forecasts conflict — treat volume status as [unverified].**
- Supply lock-up: NVIDIA deployed >$6.5B into photonics suppliers in 2026 YTD, incl. $2B each into Coherent and Lumentum with multiyear commitments [secondary: github life-os research note — **unverified**, treat as analyst claim].
- Note: one research note asserts "InnoLight + Eoptolink ~60% of NVIDIA 800G volume" [secondary: https://github.com/roachx92/equity-watch/blob/HEAD/tickers/AAOI/reports/2026-07-14.md — unverified analyst claim].

