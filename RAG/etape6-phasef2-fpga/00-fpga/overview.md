---
id: etape6-phasef2-fpga/00-fpga/overview
title: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: reference
actors: ["AMD", "Intel"]
dates: ["2025-09-12", "2026-06", "2026-09-22"]
keywords: ["accelerator", "acquisition", "amd", "asic", "benchmark", "cost", "dsp", "gpu", "inference", "intel", "lpddr5x", "memory"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [1, 54]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: a712f395039abeb4032a545aa2fedc8ac49930ea7ff45cbc02ce3ff105d77e94
---

# Phase F2 — FPGA (Field-Programmable Gate Arrays)

- **Scope**: Reconfigurable accelerator landscape for data centers: FPGA vendors and flagship families, FPGA-based accelerator/SmartNIC cards, packet-processing and P4 ecosystems, current pricing snapshots, TCO versus ASIC/GPU, deployment evidence (hyperscalers, telecom/Open RAN, HFT), and open-source FPGA tooling.
- **Observation date**: 2026-09-22 (all facts "current" are as observed on this date).
- **Method**: Read-only public-web research via search snippets and opened pages; no logins, no vendor contact, no purchases.
- **Provenance tags**: Every factual statement carries exactly one tag — `[official]` (vendor/official documentation), `[vendor-reported]` (vendor press/marketing claims), `[independent]` (independent benchmark/lab review), `[secondary]` (news/press/distributor/third-party reporting), `[unverified]` (unconfirmed, ambiguous, or inference-from-context).
- **Identifier policy**: SKUs, prices, URLs, versions, dates are copied from sources; nothing invented. Where a figure conflicts between sources, both are shown.
- **Language**: English (per project requirement).
- **Coverage warning**: This file tracks fast-moving product lines; vendor claims are kept as `[vendor-reported]` and launch dates as planned unless the vendor confirms shipping.

## 1. Altera — post-Intel independence

- Silver Lake completed acquisition of a **51%** stake in Altera on **2025-09-12/15**, with Intel retaining **49%**; Altera now presents itself as the world's largest pure-play independent FPGA provider `[secondary]` — sources: https://www.digitimes.com/news/a20250916VL200/intel-altera-silver-lake-fpga-partnership.html and https://www.edge-ai-vision.com/2025/09/altera-closes-silver-lake-investment-to-become-worlds-largest-pure-play-fpga-solutions-provider/ `[secondary]`
- Altera is headquartered as an independent company post-close; product branding has moved from "Intel Agilex" to "Altera Agilex" across datasheets and docs.altera.com `[official]`
- Open FPGA Stack (OFS) documentation and product briefs are now hosted on docs.altera.com, continuing the previously Intel-published OFS ecosystem `[official]` — source: https://docs.altera.com/api/khub/documents/hP~tPGLMHx9jBkbaZI50DQ/content `[official]`
- OFS board catalog (latest crawl observed ~189 days ago) lists partner SmartNIC/HPC cards including the Hitek Systems Agilex Low-Profile NC200 PCIe card: AGF006–AGF027 device SKUs, PCIe 4.0 x16, QSFP-DD up to 400G (8×56G PAM4), 24 GB DDR4, 75–100 W variants, part numbers AGF-NC200-B74-01 and AGF-NC200-B74-03 `[official]`

### 1.1 Agilex 5

- Agilex 5 E-Series used by MTI in a 4T4R macro radio unit delivering 25GbE fronthaul `[vendor-reported]` — source: https://www.businesswire.com/news/home/20260226443864/en/Alteras-Programmable-Solutions-Accelerate-the-Next-Generation-of-5G-Advanced-and-6G-Wideband-Radios `[vendor-reported]`
- Agilex 5 targets power- and cost-optimized midrange applications (edge AI, industrial, video) as the value tier below Agilex 7 `[official]`

### 1.2 Agilex 7

- NEC deployed Agilex 7 FPGAs in a sub-6GHz Massive MIMO 5G radio and reported **>40% power reduction** versus the prior implementation `[vendor-reported]` — source: https://www.businesswire.com/news/home/20260226443864/en/Alteras-Programmable-Solutions-Accelerate-the-Next-Generation-of-5G-Advanced-and-6G-Wideband-Radios `[vendor-reported]`
- Agilex 7 remains Altera's high-performance mainstream tier (transceiver-rich variants up to 116G) for networking, data center, and defense workloads `[official]`

### 1.3 Agilex 9 (Direct RF)

- Altera introduced the Agilex 9 Direct RF series (AGRW039 flagship) in June 2026; reported headline specs: **3.9M logic elements / 1.3M ALMs**, **12,300 DSP/AI blocks**, **18.4 TFLOPS single-precision**, integrated 64 Gsample/s RF data converters, up to **24× 58G transceivers**, hard 400GbE/200GbE and 300G Interlaken `[vendor-reported]` — sources: https://www.cnx-software.com/2026/06/09/altera-agilex-9-direct-rf-series-agrw039-soc-fpga-targets-high-performance-rf-systems/ and https://www.businesswire.com/news/home/20260608517153/en/Altera-Introduces-Next-Generation-Agilex-9-Direct-RF-Series-SoC-FPGA-to-Power-the-Future-of-High-Performance-RF-Systems `[vendor-reported]`
- Agilex 9 Direct RF engineering samples were reported available at announcement; production silicon and development kits were planned for **Q3 2026** (treat as planned/vendor-reported until confirmed shipped) `[vendor-reported]`
- Target applications: high-performance RF systems, radar, electronic warfare, satellite communications, and wideband radio `[vendor-reported]`

### 1.4 Agilex 3 (low end)

- Agilex 3 B-Series and C-Series are power- and cost-optimized devices in compact form factors targeting board monitoring/management, video/vision, protocol expansion, portable imaging, sensor fusion, drives, robotics I/O `[vendor-reported]` — source: https://it-online.co.za/2023/09/18/intel-expands-fpga-portfolio/ `[secondary]`
- Intel PSG reported 35% YoY revenue growth in Q2 2023 (third consecutive record quarter) — historical context for the portfolio's scale pre-spinout `[vendor-reported]`

## 2. AMD — Versal adaptive SoC families

### 2.1 Versal Premium Gen 2 (standard + MoP)

- Flagship SKUs **2VP3422 / 2VP3522 / 2VP3622**; top part reports up to **3,273K system logic cells**, **1,496K LUTs**, **7,616 DSP engines** `[vendor-reported]` — sources: https://www.cnx-software.com/2024/11/14/amd-versal-premium-gen2-soc-fpga-family-features-arm-cortex-a72-r5f-cores-high-end-fpga-fabric-pcie-gen6-cxl-3-1-interfaces/ and https://www.cnx-software.com/2026/07/06/amd-versal-premium-gen-2-mop-adaptive-soc-integrates-32gb-lpddr5x-to-shrink-board-footprint-by-60/ `[secondary]`
- Dual PCIe **Gen6 x8** controllers with **CXL 3.1** via CPM6 hard IP `[vendor-reported]`
- Networking: 100G multirate SerDes plus up to **5× 600GbE MACs** on top parts; hard **400G crypto engines** delivering up to 800 Gb/s line-rate encryption `[vendor-reported]`
- 56–72 **GTM2 transceivers at 112G PAM4** `[vendor-reported]`
- Security: PCIe IDE, AES-GCM-256/128 crypto engines `[vendor-reported]`
- MoP (memory-on-package) variants integrate **32 GB LPDDR5X at 9000 Mb/s** for up to **288 GB/s**, shrinking board footprint by a vendor-claimed **60%** `[vendor-reported]`
- MoP package example **VMVA3575** (55 × 57.5 mm, 0.92 mm pitch); extended (0–110 °C) and industrial (−40–110 °C) grades; AMD **15-year product lifecycle** coverage `[vendor-reported]`
- Availability: samples reported early 2026, production expected **H2 2026** (planned; not confirmed shipped as of 2026-09-22) `[vendor-reported]`; MoP sampling end 2026, MoP production H2 2027 (planned) `[vendor-reported]`
- Early-access design tools for Premium Gen 2 were announced for **Q2 2025** `[vendor-reported]`

### 2.2 Versal Prime Gen 2

