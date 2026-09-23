---
id: etape6-phasef2-fpga/00-fpga/29-amd-alveo-media-video-accelerators
title: "29. AMD Alveo media/video accelerators"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: multimodal
actors: ["AMD", "TSMC", "United States"]
dates: ["2026-05-19"]
keywords: ["accelerator", "amd", "acquisition", "asic", "compute", "decode", "ethernet", "gpu", "inference", "latency", "pricing", "robotics"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [434, 478]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: 704aec449bd777f8f612c7c5f1e832f1bb7ba41bf7dcedbcef7ef6f970ca37a2
---

# 29. AMD Alveo media/video accelerators

## 29. AMD Alveo media/video accelerators

- **Alveo MA35D**: media-accelerator card with dual 5 nm video processing units (VPUs); AV1, H.264, H.265 encode/decode; targets live-streaming and video-transcode density `[vendor-reported]`
- MA35D positioning: up to 8K multi-channel transcode per card class, low power per stream vs CPU/GPU transcode (vendor claim) `[vendor-reported]`
- Alveo **U30** (older media card): secondary-market ~US$99–280 (Sep 2026); no current OEM price found `[secondary]`
- Video/transcode is one of the longest-running FPGA data-center workloads (alongside networking and search) `[secondary]`

## 30. Achronix — company and ecosystem detail

- Achronix (founded 2004, Santa Clara) is an independent FPGA vendor: standalone **Speedster7t** FPGAs plus **Speedcore** embedded-FPGA IP for ASIC/SoC integration `[secondary]`
- **Speedster7t AC7t800** (recap): TSMC 7 nm, 711K LEs, 864 MLPs, 12 Tb/s 2D NoC, 6× GDDR6 (1.5 Tb/s), 2× 400GbE, PCIe Gen5 x16 `[vendor-reported]`
- Speedster7t datasheet v2.3 released **2026-05-19** — current doc baseline `[official]`
- Toolchain: **ACE** (Achronix CAD Environment) for place-and-route; Synopsys Synplify Pro for synthesis `[official]`
- **VectorPath S7t-VG6** (BittWare): Speedster7t accelerator card, GDDR6, 1× QSFP-DD (400G or multi-100G) `[secondary]`
- Speedcore eFPGA: licensable fabric blocks integrated into customer ASICs — the "FPGA inside ASIC" model distinct from standalone FPGA cards `[official]`
- Achronix targets: AI/ML inference, smartNIC/DPU offload, 5G/ORAN, computational storage, test & measurement `[vendor-reported]`

## 31. Lattice — Avant detail and legacy families

- **Avant** platform (16 nm): midrange, 200K–500K logic cells class, up to 25G SerDes, PCIe Gen4, external DDR — Lattice's move up from small-FPGA roots `[vendor-reported]`
- **Avant-E** parts: midrange edge processing (video, industrial automation, robotics) `[secondary]`
- **Avant-G / Avant-X**: connectivity/compute-oriented Avant variants (naming per Lattice portfolio) `[official]`
- **Avant 30 / Avant 50** (Dec 2024): new midrange entries; no public pricing; eval board "coming soon" per Sep 2026 coverage `[secondary]`
- **Certus-N2** (Nexus 2, 16 nm): 64K–220K LCs, 8× 16G SerDes, PCIe Gen4, FIPS 140-3 L2-claimed security; sampling `[vendor-reported]`
- **Mach-N2** (Sep 16, 2026): secure-control family, CNSA 2.0 PQC, orderable now `[vendor-reported]`
- Legacy/adjacent: **iCE40 UltraPlus** (ultra-low-power sensor hub), **ECP5** (open-flow friendly, Trellis), **MachXO3D** (secure control predecessor), **CrossLink-NX** (MIPI/vision bridging) `[official]`
- Lattice's data-center relevance: board-management, control-plane, RoT-adjacent, and SmartNIC sidecar roles rather than primary compute acceleration `[unverified]`

## 32. Microchip — tools and ecosystem

- **Libero SoC**: Microchip's FPGA design suite for PolarFire/PolarFire SoC (synthesis, P&R, power, security flows) `[official]`
- **Icicle Kit**: PolarFire SoC evaluation/development kit with the 5-core RISC-V subsystem `[official]`
- **Mi-V ecosystem**: Microchip's RISC-V partner/program ecosystem around PolarFire SoC soft and hard CPU subsystems `[official]`
- **SmartHLS**: high-level synthesis flow for PolarFire (C/C++ to RTL) `[official]`
- PolarFire Ethernet Sensor Bridge rev 2.0 (Aug 11, 2026): 60% smaller, 4 cameras, 10GbE sensor transport `[vendor-reported]`
- RT PolarFire: radiation-tolerant variant for space — 2026 status not researched (gap) `[unverified]`

## 33. Exablaze / Cisco — ultra-low-latency NICs

- Exablaze (acquired by Cisco) built FPGA-based **ExaNIC** NICs for sub-microsecond trading and deterministic packet capture `[secondary]`
- ExaNIC X25/X100-class products: 10/25/100G ultra-low-latency NICs with hardware timestamping `[secondary]`
- Post-acquisition mapping: **ExaNIC → Nexus SmartNIC/NIC**, **ExaLINK → Nexus 3550** (Cisco naming) `[official]` — source: https://www.cisco.com/site/us/en/about/corporate-development/acquisitions/exablaze/index.html `[official]`
- Exablaze also offered **exaclock** precision-timing products for PTP-synchronized trading fabrics `[secondary]`
- The Exegy–NovaSparks deal (Jan 14, 2026) further consolidates this FPGA-trading vendor space `[secondary]`

