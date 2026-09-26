---
id: etape6-phasef2-fpga/00-fpga/8-packet-processing-and-p4-on-fpgas
title: "8. Packet processing and P4 on FPGAs"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: reference
actors: ["AMD", "Intel", "United States"]
dates: ["2025-06", "2026-09-22"]
keywords: ["accelerator", "amd", "gpus", "intel", "latency", "pricing", "research"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [145, 179]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: 48e02d70f4c7ff8cc43cf69b71b4ed05edc35b2e192d41932289c812444b297e
---

# 8. Packet processing and P4 on FPGAs

- **Napatech**: FPGA-based SmartNIC/NIC vendor (Link/NT series); partnerships with Intel/Altera and Silicom observed; **no evidence found that Intel acquired Napatech** — any "Napatech (now Intel)" claim encountered in briefs is **incorrect/unverified** and must not be repeated as fact `[unverified]`
- **Silicom**: white-box partner integrating Napatech FPGA SmartNIC technology into server adapters `[secondary]` — source: https://www.eeworldonline.com/fpga-based-software-works-with-intel-powered-smartnics/ `[secondary]`
- **Accolade Technology**: FPGA-based packet-capture/processing adapters (ANIC series) for network monitoring, recording, and inline security — listed here as a required vendor; detailed 2026 product/pricing research is a **gap** (not yet collected) `[unverified]`
- **Cisco / Exablaze**: Exablaze rebranded under Cisco — **ExaNIC** products map to **Nexus SmartNIC/NIC** names, **ExaLINK** to **Nexus 3550**; ExaNIC targets ultra-low-latency trading and high-performance packet capture `[official]` — source: https://www.cisco.com/site/us/en/about/corporate-development/acquisitions/exablaze/index.html `[official]`
- **BittWare (Molex)**: FPGA accelerator/SmartNIC boards across AMD/Xilinx, Altera, and Achronix devices; **SmartNIC Shell** for building FPGA-powered 100G NICs with DPDK support, match-action processing, Xilinx SDNet/P4, and timestamping `[vendor-reported]` — source: https://embeddedcomputing.com/application/networking-5g/bittware-announces-smartnic-shell-for-building-fpga-powered-100g-nics-supports-dpdk-xilinx-sdnet-p4-programming-user-customizations-and-timestamping `[secondary]`
- **Hitek Systems NC200** (Agilex, OFS-based, up to 400G QSFP-DD) — see §1 `[official]`

## 8. Packet processing and P4 on FPGAs

- **Corundum**: open-source FPGA NIC framework targeting up to **100 Gb/s**; custom DMA engine + Linux driver; scatter/gather, PTP timestamping, **10,000+ queues**, user-logic regions `[independent]` — sources: http://bittware-molex.com/files/corundum_ip_datasheet.pdf, https://leed.sysnet.ucsd.edu/papers/CorundumPaper.pdf, https://github.com/corundum/corundum/wiki/Frequently-Asked-Questions `[independent]`
- Corundum vs **OpenNIC** (per Corundum FAQ): OpenNIC uses AMD/Xilinx **QDMA** hard IP and benefits from existing drivers/DPDK PMD; Corundum uses a **custom DMA** engine with broader board support; both can host P4-generated packet-processing pipelines `[independent]`
- **P4 on FPGA**: AMD/Xilinx **SDNet** compiles P4 to FPGA logic; BittWare SmartNIC Shell exposes SDNet/P4 programmability `[vendor-reported]`; P4→FPGA flows also exist via open tooling (P4C backends, Corundum match-action examples) `[independent]`
- **DPDK on FPGA NICs**: BittWare SmartNIC Shell ships DPDK support; OpenNIC ships a DPDK poll-mode driver; Corundum has DPDK integration work in its ecosystem `[independent]`
- PTP/1588 hardware timestamping is a standard FPGA-NIC feature (Hitek NC200 lists full PTP/1588 sync support) `[official]`

## 9. Toolchains — vendor flows

- **Altera Quartus Prime**: Agilex/Stratix/Cyclone synthesis, place-and-route, timing closure, and signed production bitstreams; only path to production bitstreams on modern Altera devices `[official]`
- **AMD Vivado / Vitis 2026.1**: Vitis 2026.1 bundles Vivado 2026.1 (vitis, v++, aiecompiler, bootgen, sdtgen, lopper, emconfigutil, aarch64-linux-gnu-g++) — current as of 2026 per AMD tutorial repos `[secondary]` — source: https://github.com/xilinx/vitis-tutorials/blob/HEAD/Vitis_Platform_Creation/Design_Tutorials/05_Edge_VEK280_DFX/README.md `[secondary]`
- **Vivado Design Suite 2026.1**: new **tiered licensing model**, new device support, major NoC and QoR improvements, ease-of-use enhancements for Versal adaptive SoCs `[secondary]` — source: https://www.dataweek.co.za/25105r `[secondary]`
- **Vivado 2025.1** (30 June 2025): added AMD **Spartan UltraScale+** and next-gen Versal device support; FMAX/QoR enhancements for Versal SSIT devices; unified selective device installer for all Versal devices; AXI switch IP; DFX summary GUI `[secondary]`
- **Achronix ACE**: place-and-route toolchain for Speedster7t (separate from Quartus/Vivado) `[official]`
- **Lattice Radiant / Propel**: Nexus/Avant design software `[official]`; **Microchip Libero SoC**: PolarFire/PolarFire SoC flow `[official]`

## 10. Pricing snapshots (as observed 2026-09-22)

- AMD Alveo V80 MSRP: **US$9,495** `[vendor-reported]`
- AMD Alveo U280: launch US$5,995 (standard) / US$9,995 (premium, 2018); secondary market ~US$999–1,450 (Sept 2026) `[secondary]`
- AMD Alveo U250: launch from US$8,995 (2018); secondary market ~US$2,250–5,479 (Sept 2026) `[secondary]`
- AMD Alveo U200: secondary market ~US$1,350 (Sept 2026) `[secondary]`
- AMD Alveo U30: secondary market ~US$99–280 (Sept 2026) `[secondary]`
- Alinx AXU-series third-party eval boards: ~US$900–1,050 (Sept 2026) `[secondary]`
- **Gaps**: no current OEM list prices found for Altera Agilex 5/7/9 dev kits, AMD Versal eval kits (VEK280/VCK190/VMK180), Lattice Avant/Nexus 2 kits, Microchip PolarFire kits, or Achronix VectorPath cards — vendor quote/authorized-distributor pricing only `[unverified]`
- Enterprise FPGA cards (V80-class) sit in the **US$5k–15k** band at MSRP, comparable to high-end data-center GPUs per-card but with very different performance profiles per workload `[unverified]`

