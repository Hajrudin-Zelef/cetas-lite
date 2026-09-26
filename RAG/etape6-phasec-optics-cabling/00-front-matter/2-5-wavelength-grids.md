---
id: etape6-phasec-optics-cabling/00-front-matter/2-5-wavelength-grids
title: "2.5 Wavelength grids"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "China", "Cohere", "Nvidia", "TSMC"]
dates: ["2025-03", "2026-03"]
keywords: ["wavelength", "2nm", "3nm", "backlog", "cpo", "dsp", "ethernet", "gpus", "npo", "nvidia", "optics", "research"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [286, 304]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 7b38c2eb9b49a96551c34c241e20adbf377b9258fcb2ec3bf493321d82889b04
---

# 2.5 Wavelength grids

- **Standards path**: IEEE P802.3dj defining 200G/lane building blocks for 200/400/800GbE and 1.6TbE; expected completion "second half of 2026" [secondary — Ethernet Alliance via NetworkWorld]; ratification expected Oct 2026 [vendor-reported — Nokia]; **status as of today: NOT yet published — tag final ratification date [unverified]**.
- **Electrical prerequisite**: 200G-class SerDes. 800G AI lossless-link demo at ECOC 2026 exercises link negotiation, Link Layer Retry (LLR) and Credit-Based Flow Control (CBFC) "running alongside IEEE P802.3dj" [secondary — Converge Digest].
- **Form factors for 1.6T**:
  - QSFP-DD1600: 8×200G PAM4, 20–28 W typical, Rev 7.0/7.1 MSA [secondary — fibermall].
  - OSFP1600: OSFP MSA Rev 5.1, 8×200G-class lanes [secondary — roboticsandautomationnews].
  - OSFP-XD: 16×100G lanes (reuses 100G SerDes ecosystem), ~40 W target, forward path to 3.2T via 16×200G [secondary — Medium/aicplight].
- **Vendor announcements** (all dated):
  - **Marvell** — Nova 2: "industry's first PAM4 optical DSP integrating 200 Gbps electrical and optical interfaces," 8×200G electrical + 8×200G optical, backward compat up to 3 generations [official — Marvell press, ~Jan 2025 via Nasdaq repost]. 5nm 200G/lane Nova (2023), 3nm 1.6T Ara platform (2024), **2nm 800G Libra DSP announced earlier in 2026** [secondary — BizTech Weekly]. At **ECOC 2026 (Sept 20–24)**: 38 demos incl. industry-first 2nm 400G/lane optical PAM4 (path to 3.2T), 2nm 800G ZR/ZR+ OSFP pluggable with MACsec (Libra DSP), 2nm 1.6T ZR and coherent-lite O-band demos, 102.4T CPO platform on 200G/lane SiPh [official — Marvell press release 21 Sept 2026 via BusinessWire]. No availability dates given [secondary]. Sources: https://www.businesswire.com/news/home/20260920935708/en/Marvell-to-Showcase-Industry-First-2nm-Optical-Technology-Demos-for-AI-Data-Center-Infrastructure-at-ECOC-2026 ; https://www.nasdaq.com/press-release/marvell-extends-1-6t-connectivity-leadership-with-industrys-first-pam4-optical-dsp
  - **Broadcom** — Sian3 (3nm) and Sian2M (5nm) DSPs for 800G/1.6T transceivers [secondary — ainvest, Apr 2025]. Bailly: 51.2T CPO Ethernet switch (8× 6.4T SiPh optical engines + Tomahawk 5), ~70% power reduction claimed [official — Broadcom via Nasdaq repost]. At **OFC 2026 (March)**: 3.2T VCSEL-based NPO (near-packaged optics) product line, ~1 pJ/bit claimed, 18 engines, 73.7 Tbps escape bandwidth [secondary — XenoSpectrum, Aug 2026]. Sources: https://www.nasdaq.com/press-release/broadcom-delivers-industrys-first-51-2-tbps-co-packaged-optics-ethernet-switch ; https://xenospectrum.com/en/near-packaged-optics-npo-cpo-interim-architecture-2026/
  - **NVIDIA** — Spectrum-X Photonics / Quantum-X Photonics CPO silicon-photonic switches (GTC, **19 March 2025**) [official]: 128×800G or 512×800G (100T/400T configs); Quantum-X 144×800G InfiniBand on 200G SerDes; built on TSMC COUPE (65nm EIC + PIC, SoIC-X); ecosystem incl. Coherent, Corning, Foxconn, Lumentum, Senko, **and pluggable partners Coherent, Eoptolink, Fabrinet, Innolight**; Quantum-X availability "later this year" (2025), Spectrum-X in 2026 [official — NVIDIA newsroom]. Claimed: 3.5× power efficiency, 63× signal integrity, 10× resiliency, 4× fewer lasers vs pluggables [official]. Sources: http://nvidianews.nvidia.com/news/nvidia-spectrum-x-co-packaged-optics-networking-switches-ai-factories ; https://www.tomshardware.com/networking/nvidias-silicon-photonics-based-1-6-tb-s-switch-platforms-enable-clusters-with-millions-of-gpus
  - **Coherent** — at **OFC 2026 (announced 17 March 2026)** [official — GlobeNewswire]: multiple 1.6T transceivers across SiPh PIC, high-power InP CW laser, 200G InP EML, 200G GaAs VCSEL, with "three different DSP solutions from three industry leaders"; 400G/lane PAM4 links for 3.2T (differential EML + pure-silicon MZM); new multi-lane **XPO pluggable MSA form factor** aimed at 12.8T and beyond. At **ECOC 2026**: 3.2T pluggable (OSFP-size, dual 1.6T paths, 8× 425G PAM4 line lanes, differential EMLs), 6.4T NPO optical engine (SiPh + fiber attach + ELSFP), CPO components FlexConnect/OptiAlign/OptiFocus, PhotonLink [secondary — Fibre Systems]. Sources: https://www.globenewswire.com/news-release/2026/03/17/3257277/11543/en/Coherent-Demonstrates-Technologies-for-Next-Generation-Pluggable-Transceivers-at-OFC-2026.html ; https://www.fibre-systems.com/article/ecoc-2026-coherents-ai-optical-innovations
  - **InnoLight (Zhongji Innolight)** — H1 2026 results: **"1.6T silicon photonics modules have entered the volume ramp phase, with quarterly shipments climbing and becoming the core product driving revenue growth"**; order backlog covers all of 2026 with some 2027 orders; XPO and NPO products in custom development/R&D [secondary — BigGo Finance, Aug 2026]; R&D H1 ~1.15B yuan (+96.9% YoY) [secondary]. InnoLight cited as ~27% global optical-module market share [secondary — Wccftech, Aug 2026]. **FCC reportedly "mulls banning China-sourced optical transceivers"** — action not confirmed; tag **[unverified]** [secondary — Wccftech]. Sources: https://finance.biggo.com/news/2b25909c-bd1d-49fd-8c51-a9c7996700d5 ; https://wccftech.com/the-fcc-mulls-banning-china-sourced-optical-transceivers-threatening-innolights-27-global-market-share-as-coherent-and-lumentum-prepare-to-pounce/
  - **Eoptolink** — at **OFC 2026 (March 2026)**: demonstrated 1.6T DR4 optical transceiver "using 400G-per-lambda technology"; 100G-per-lambda 800G products already in high-volume deployment; 200G-per-lambda 1.6T ramping; also unveiled 12.8 Tbps liquid-cooled pluggable-optics concept for AI DCs [secondary — Semiconductor Insight]. Source: https://semiconductorinsight.com/report/optical-module-package-market/
  - **Applied Optoelectronics (AOI)** — 800G "expects 800G to become its largest data-center revenue driver in 2026" (from Q2 2026); 800G production capacity ~90k units/month end-2025; **target >500,000 units/month of 800G and 1.6T combined by end-2026**; demand expected to exceed capacity through mid-2027; new hyperscale customer qualifying 800G and 1.6T [secondary — Zacks, 10 March 2026]. Source: https://www.zacks.com/stock/news/2881638/will-aaois-800g-transceivers-accelerate-its-data-center-momentum
  - **Hisense Broadband, Accelink, Sumitomo, Source Photonics** — no specific 2026 1.6T announcements found in this research pass **[gap — not found]**; note Yole 2026 report analyzes "emerging Chinese vendors Terahop, Eoptolink, and Accelink" alongside Coherent/Nvidia/Lumentum [secondary — optics.org].
- **1.6T production timeline consensus**: engineering samples in customer validation; volume production **late 2026 through 2027** [secondary — roboticsandautomationnews, Sept 2026].

### 2.5 Wavelength grids

