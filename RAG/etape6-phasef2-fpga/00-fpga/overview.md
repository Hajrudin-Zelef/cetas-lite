---
id: etape6-phasef2-fpga/00-fpga/overview
title: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: reference
actors: ["AMD", "Intel", "United States"]
dates: ["2025-09-12", "2026-06", "2026-09-22"]
keywords: ["accelerator", "acquisition", "amd", "asic", "benchmark", "compute", "cost", "dsp", "gpu", "hbm", "inference", "intel"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [1, 78]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: d6c32776f3e0fec746ff2570b9ff0d7fdd2ff3ca3f6b22afcf05bb7b87062054
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

- New compact SKUs **2VM3104 / 2VM3254 / 2VM3454** in **23×23 mm** (2VM3104, 2VM3254) and **29×29 mm** (2VM3454) packages; 2VM3654 released earlier shares the footprint for drop-in migration `[vendor-reported]` — source: https://www.cnx-software.com/2026/05/29/amd-launches-versal-prime-gen-2-2vm3454-2vm3254-and-2vm3104-adaptive-socs-in-compact-23x23mm-packages/ `[secondary]`
- Dual-core **Arm Cortex-A78AE + Cortex-R52** application/real-time processors `[vendor-reported]`
- Safety: **ASIL D / SIL 3** features integrated, removing the need for a separate safety controller in automotive/industrial designs `[vendor-reported]`
- High-speed connectivity (HSC) engine available on 2VM3454 `[vendor-reported]`
- Tools: **Vivado** for RTL, **Vitis** for software on the A78/R52 cores, **Embedded Development Framework (EDF, Yocto-based)** with ready Linux images and video/GPU/fast-memory drivers; bare-metal on real-time cores `[vendor-reported]`
- Availability: early-access tools already available for 2VM3654; 2VM3454 sampling later 2026 (planned); 2VM3254/2VM3104 expected **2027** (planned) `[vendor-reported]`

### 2.3 Versal AI Edge Gen 2

- Versal AI Edge Series Gen 2 targets edge AI with AIE-ML engines; VEK280 evaluation board carries **xcve2802-vsvh1760-2MP-e-S** `[secondary]` — source: https://github.com/xilinx/vitis-tutorials/blob/HEAD/Vitis_Platform_Creation/Design_Tutorials/05_Edge_VEK280_DFX/README.md `[secondary]`
- Vivado 2025.1 added new addressing GUI grouping equivalent address spaces for **Versal Prime Gen 2 and Versal AI Edge Gen 2** `[secondary]` — source: https://www.dataweek.co.za/25105r `[secondary]`

### 2.4 Versal HBM (XCVH1782)

- XCVH1782 integrates **HBM2e up to 819 GB/s** peak memory bandwidth for memory-bound acceleration `[vendor-reported]` — source: https://www.servethehome.com/amd-alveo-v80-versal-hbm-networking-accelerator/ `[secondary]`
- The Alveo V80 card is built on the **XCV80** Versal HBM device with 10,848 DSP slices and **32 GB HBM2e at 820 GB/s peak** `[vendor-reported]` — source: https://www.servethehome.com/new-faster-amd-alveo-v80-accelerator-with-hbm2e-and-fast-networking/ `[secondary]`
- Alveo V80 networking: **4× QSFP56** (2× 100G per cage), PCIe **Gen4 x16 or Gen5 x8**, **190 W** card power, reported MSRP **US$9,495** `[vendor-reported]` — sources: https://www.servethehome.com/new-faster-amd-alveo-v80-accelerator-with-hbm2e-and-fast-networking/ and https://wccftech.com/amd-announces-mass-production-of-the-alveo-v80-compute-accelerator-9495-price-tag/amp/ `[vendor-reported]`

### 2.5 Versal Premium VP1902 (emulation/prototyping)

- VP1902: **18.5M logic cells**; vendor claims 2× programmable-logic density and 2× aggregate I/O bandwidth vs prior-gen Virtex UltraScale+ VU19P; up to 8× faster debugging via programmable NoC (vendor comparisons, baseline named = VU19P) `[vendor-reported]` — source: https://www.techpowerup.com/310577/amd-introduces-worlds-largest-fpga-based-adaptive-soc-for-emulation-and-prototyping?amp `[secondary]`
- Target: pre-silicon emulation/prototyping of AI, automotive, Industry 5.0 ASIC/SoC designs; ecosystem with Cadence, Siemens, Synopsys EDA flows `[vendor-reported]`
- Sampling began Q3 (year of announcement per source) to early-access customers; production expected first half of the following year (historical; status beyond that not verified here) `[vendor-reported]`

