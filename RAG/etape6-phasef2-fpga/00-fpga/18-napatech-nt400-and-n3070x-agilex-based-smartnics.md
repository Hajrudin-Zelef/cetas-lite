---
id: etape6-phasef2-fpga/00-fpga/18-napatech-nt400-and-n3070x-agilex-based-smartnics
title: "18. Napatech — NT400 and N3070X (Agilex-based SmartNICs)"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: reference
actors: ["AMD", "Intel", "TSMC", "United States"]
dates: []
keywords: ["accelerator", "amd", "asic", "chiplet", "compute", "dsp", "ethernet", "full-duplex", "inference", "int4", "intel", "memory"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [291, 317]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: 84bbfc8655c50e41b2c9d36635a08aef821550480359bc7ff1de01f48a6776cb
---

# 18. Napatech — NT400 and N3070X (Agilex-based SmartNICs)

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

