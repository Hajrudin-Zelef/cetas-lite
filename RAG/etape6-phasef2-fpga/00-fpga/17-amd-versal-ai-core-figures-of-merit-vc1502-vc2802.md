---
id: etape6-phasef2-fpga/00-fpga/17-amd-versal-ai-core-figures-of-merit-vc1502-vc2802
title: "17. AMD Versal AI Core — figures of merit (VC1502 → VC2802)"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: reference
actors: ["AMD", "Anthropic", "Intel", "TSMC", "United States"]
dates: ["2026-09-16"]
keywords: ["amd", "accelerator", "agentic", "asic", "chiplet", "claude", "compute", "dsp", "energy", "ethernet", "full-duplex", "inference"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [289, 335]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: f0d06ff4ee4132fe07028b0fef2a20449d837be01b9b44423458e270b99520a6
---

# 17. AMD Versal AI Core — figures of merit (VC1502 → VC2802)

## 17. AMD Versal AI Core — figures of merit (VC1502 → VC2802)

- Figures below from an AMD dealer-mirror PDF of the Versal AI Core series figures-of-merit table; tag all `[secondary]` — source: https://amd.nt-rt.ru/images/manuals/k3-1.pdf `[secondary]`
- AI Engine peak INT8x4: **101 / 101 / 100 / 133 / 202 / 405 TOPs** across VC1502 / VC1702 / VC1802 / VC1902 / VC2602 / VC2802 `[secondary]`
- AI Engine peak INT8: **101 / 101 / 100 / 133 / 101 / 202 TOPs** `[secondary]`
- AI Engine peak INT16: **25 / 25 / 25 / 33 / 25 / 51 TOPs**; FP32: **6 / 6 / 6 / 8 / 8 / 17 TFLOPs** `[secondary]`
- AI Engine peak SRAM bandwidth: **405 / 405 / 399 / 532 / 202 / 405 Tb/s** `[secondary]`
- DSP engine peak INT8: **9.1 / 9.1 / 11.0 / 13.6 / 6.8 / 9.1 TOPs**; INT24: 3.0 / 3.0 / 3.7 / 4.5 / 2.3 / 3.0 TOPs `[secondary]`
- Adaptable-engine peak INT4: **45 / 56 / 90 / 112 / 47 / 65 TOPs**; INT8: 12 / 14 / 23 / 29 / 12 / 17 TOPs `[secondary]`
- Scalar engines: Arm Cortex-A72 **18,942 DMIPs** (19,516 on VC2602/VC2802); Cortex-R5F 2,672 DMIPs `[secondary]`
- Memory/NoC: total SRAM bandwidth 92–188 Tb/s; transceiver bandwidth **2.48 Tb/s** (2.10 on VC2602/VC2802); NoC cross-sectional bandwidth 1.2–2.2 Tb/s; DDR4 up to 102.4 GB/s; LPDDR4 up to 136.5 GB/s `[secondary]`
- **VCK5000** inference card: XCVC1902, **400 AI Engines @ 1.25 GHz**, **1,968 DSP engines**, up to **145 TOPS INT8** for AI inference `[secondary]` — source: https://www.eenewseurope.com/en/ai-and-dsp-boost-for-fpga-signal-processing/ `[secondary]`
- Versal Premium with AI Engines (VP2502/VP2802): **472 AI accelerator cores** as chiplets over processor+FPGA substrate; **7,392 / 14,304 DSP engines**; up to **9 Tb/s** serial bandwidth; hardened 100G/600G Ethernet, 400G crypto engines, PCIe 5.0; vendor-claimed 4× signal-processing capacity vs 16 nm Virtex UltraScale+ VU13P `[vendor-reported]`
- **XQR Versal AI Core (space)**: Class B qualification (US DoD) for operation in space; **400 AI engines**, ~**900K logic cells**, 191 Mbits memory, TSMC 7 nm; reprogrammable on-orbit; **Raytheon** named as customer for next-gen space processors; shipping was expected early the year after announcement (2022 announcement — historical) `[secondary]` — sources: https://siliconangle.com/2022/11/15/amd-announces-reprogrammable-versal-ai-chips-ready-space-based-deployments/ and https://www.Theregister.Com/2022/11/15/amd_versal_space_chip/?td=keepreading `[secondary]`
- VP1902 detail: **18.5M logic cells / 60B logic gates**, measured throughput **5.6 Tbps** (~2× VU19P); **no AI Engines** (pure-compute emulation focus); volume shipments began Q3 of launch year `[secondary]` — source: https://www.hpcwire.com/2023/06/28/fpga-development-brings-amd-and-intel-competition-to-the-forefront/ `[secondary]`

## 18. Napatech — NT400 and N3070X (Agilex-based SmartNICs)

- **NT400**: based on the Intel FPGA SmartNIC **N6000-PL** platform; PCIe **Gen4 x16** host; **2× QSFP56** (up to 2× 200G, configurable 10/25/40/50/100/200G); full-duplex 2× 100Gbps host flow; sustains **400G total over tens of millions of flows** (vendor-reported) `[vendor-reported]` — source: https://www.napatech.com/media/press-releases/napatech-leverages-latest-intel-agilex-fpga-to-launch-industrys-first-400gbps-smartnic-solutions/ `[vendor-reported]`
- NT400 core: **Intel Agilex 7 FPGA F-Tile** chiplet with hardened configurable Ethernet stack 10G–400G; five supported configurations for price/performance/power trade-offs `[vendor-reported]`
- Software suites: **Link-Capture™** (monitoring/recording), **Link-Virtualization™** (cloud vSwitch data plane), **Link-Inline™** (5G UPF and inline apps); install-card-and-software model, no direct FPGA programming required `[vendor-reported]`
- **5G UPF offload** on NT400: up to **190M concurrent flows**, **>2M flows/sec** learning rate, **113M packets/sec** stateful, full 2× 200G wire speed at typical packet sizes (vendor-reported); ecosystem partners **A5G Networks, Advantech, Druid Software, Kontron** `[vendor-reported]` — source: https://www.napatech.com/media/press-releases/napatech-launches-5g-upf-offload-solution-on-its-intel-agilex-fpga-based-400gbps-smartnic/ `[vendor-reported]`
- Napatech UPF TCO claim: **90× more users/server than software UPF**, **7× more than a competing ASIC-based SmartNIC** in a representative operator use case — vendor analysis, treat as `[vendor-reported]`, not independently verified `[vendor-reported]`
- **N3070X** (newer 400G platform, per distributor-hosted datasheet mirror — tag `[secondary]`): Intel Agilex **AGIC041** (1× F-tile, 3× R-tile), **4M LEs / 1,356.1K ALMs / 17.1K M20Ks**, AES/SM4 crypto support `[secondary]` — source: https://cdn-adepci1.actonsoftware.com/acton/cdna/14951/f-d0dea8a4-3331-49df-910e-2faaebb4575f/1/10 `[secondary]`
- N3070X interfaces: **3× PCIe Gen5 x16** (32 GT/s), card-edge host, **CXL 2.0** (mount option), **2× QSFP-DD** (1× 400G / 2× 200G / 4× 100G), **4× 4 GB DDR4-2666 ECC** (options to 4× 8 GB), 2× 2048 Mbit boot flash, integrated BMC, **HW RoT (option)**, MCTP/PLDM manageability, passive cooling, max **150 W**, FH 1/2-length (366 g) and FH 3/4-length (429 g) variants `[secondary]`
- Napatech software roadmap for N3070X lists Link-Programmable™ supported, Link-Inline™/Link-Capture™/Link-Virtualization™ as directional roadmap; **Ultra Ethernet (UEC)** noted as roadmap item `[secondary]`

## 19. Lattice — Nexus 2 detail, Certus-N2, Mach-N2, Avant 30/50

- Nexus 2 platform (advance datasheet FPGA-DS-02122-0.73, ©2026): **TSMC 16 nm FinFET**, **65K–220K system logic cells**, 120–520 18×18 sysDSP multipliers, **4–12 Mb** EBR, timing closure up to **350 MHz** typical (EBR/DSP/clocks 625 MHz) `[official]` — source: https://latticesemi-prod-latticesemi-live.vercel.app/view_document?document_id=54437 `[official]`
- Nexus 2 interfaces: **2–8× 16G** multiprotocol SerDes, hardened **PCIe Gen1/2/3/4**, DDR4/LPDDR4 to **2400 Mbps**, DDR3L 1866, hardened MIPI D-PHY to 4.5 Gbps/lane and C-PHY to 7.98 Gbps/trio, up to **349** programmable sysI/O `[official]`
- Nexus 2 security: **AES-256-GCM**, ECDSA-521 / RSA4096 auth, anti-tamper, PUF/unique ID, user-data encryption, side-channel resistance, TRNG `[official]`
- Vendor comparisons for Nexus 2 (unnamed competitor baselines — keep `[vendor-reported]`): up to **3× lower power**, up to **10× more energy-efficient** edge sensor monitoring, up to **3.2× faster MIPI**, up to **10× faster configuration**, up to **5× smaller** form factor `[vendor-reported]` — source: https://www.edn.com/lattice-launches-small-size-fpga-platform/ `[secondary]`
- **Certus-N2** (first Nexus 2 family, LN2-CT): 64K–220K logic cells, up to 135K LUTs, up to **520** 18×18 multipliers, up to **12 Mb** embedded memory, **8× 16G** SerDes (128 Gbps total), up to 7 GPLLs, <**20 ms** instant-on via xSPI flash, packages down to **9×9 mm** (0.5/0.8/1.0 mm pitch), commercial + industrial grades, integrated AES-GCM/ECDSA/RSA engines; **sampling** as of Dec 2024 announcements `[vendor-reported]` — sources: https://www.cnx-software.com/2024/12/12/lattice-unveils-nexus-2-small-fpga-platform-lattice-avant-30-and-avant-50-mid-range-devices-updated-lattice-design-software-tools/ and https://www.hackster.io/news/lattice-semiconductor-launches-the-certus-n2-small-fpga-family-built-around-the-nexus-2-platform-f0e437012f1b.amp `[secondary]`
- Certus-N2 example solutions: industrial machine-vision platform (MIPI CSI-2/LVDS/sub-LVDS, custom ISP, 10GbE), automotive zonal controller/gateway, 10GbE frame-grabber with PCIe output and LPDDR4 frame buffer `[vendor-reported]`
- **Mach-N2** (announced **2026-09-16**, six days before observation date): secure-control FPGAs on Nexus 2, **65K–220K** SLCs across five devices, integrated nonvolatile flash, hardware Root of Trust, **CNSA 2.0-compliant post-quantum cryptography** (ML-DSA, LMS, XMSS auth; **ML-KEM** key exchange), crypto agility, 128–256 Mb config flash, **sub-30 ms** instant-on, up to 3 stored recovery images, PCIe 4.0 / 10G Ethernet / 16G SerDes; **available for order now**, samples shipped to compute/communications customers `[vendor-reported]` — source: https://futurumgroup.com/insights/lattice-mach-n2-turns-the-post-quantum-procurement-gate-into-an-fpga-moat/ `[secondary]`
- **Lattice Prompt** (announced 2026-09-16): free AI FPGA design assistant linking agentic IDEs/LLMs (**Claude Code, Cursor, VS Code**) to **Lattice Radiant** via open **Model Context Protocol (MCP**), orchestrating simulation, synthesis, place-and-route, timing analysis, bitstream generation; vendor-reported **10×+ productivity gains** `[vendor-reported]`
- **Avant 30 and Avant 50** (announced Dec 2024): new midrange entries expanding the Avant portfolio for connectivity and functional capacity; evaluation board listed "coming soon" (status per Sep 2026 article); **no public pricing confirmed** `[secondary]` — source: https://www.hackster.io/news/lattice-semiconductor-launches-the-certus-n2-small-fpga-family-built-around-the-nexus-2-platform-f0e437012f1b.amp `[secondary]`
- Lattice software: **Radiant** and **Propel** updated for Nexus 2/Certus-N2/Avant (RISC-V variants, improved debug, power calculation); solution stacks **sensAI** (edge AI), **mVision** (embedded vision), **Automate** (factory automation), **Drive** (automotive) `[vendor-reported]` — source: https://www.eeworldonline.com/pcie-gen-4-integration-added-to-small-form-factor-fpga/ `[secondary]`
- FIPS 140-3 **Level 2** compliance claimed for Nexus 2 AES-GCM/SHA3-512 functions (vendor claim; certification status should be verified in NIST CMVP listings before quoting as certified) `[vendor-reported]`

## 20. Microchip PolarFire — additional detail

- PolarFire SoC product overview: 5-core RISC-V subsystem (4× RV64GC + 1× RV64IMAC monitor), deterministic low-latency fabric, SECDED on memories, single-event-upset (SEU) immunity positioning for space/defense `[official]` — source: https://ww1.microchip.com/downloads/en/DeviceDoc/Polarfire_SOC_Product_Overview.pdf `[official]`
- PolarFire (non-SoC) focus: lowest-power midrange FPGA, **12.7 Gbps** SerDes, PCIe Gen2, DDR3/4, hardened crypto — suited to always-on, thermally constrained edge and comms workloads `[vendor-reported]`
- RT PolarFire (radiation-tolerant) extends the family into space applications — named here for completeness; detailed 2026 status not researched (gap) `[unverified]`

