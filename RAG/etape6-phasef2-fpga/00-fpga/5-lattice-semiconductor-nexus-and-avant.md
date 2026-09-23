---
id: etape6-phasef2-fpga/00-fpga/5-lattice-semiconductor-nexus-and-avant
title: "5. Lattice Semiconductor — Nexus and Avant"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: reference
actors: ["AMD", "Intel", "United States"]
dates: ["2025-06", "2026-08-11", "2026-09-22"]
keywords: ["accelerator", "amd", "benchmarks", "compute", "ethernet", "gpus", "intel", "latency", "pricing", "research", "robotics", "serdes"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [128, 179]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: 7a906da108590b4a3f6d1323742c3492b2e5642006f72235355256fcada29e2a
---

# 5. Lattice Semiconductor — Nexus and Avant

## 5. Lattice Semiconductor — Nexus and Avant

- **Avant** is a 16 nm midrange platform: broadly **200K–500K logic cells**, up to **25G SerDes**, PCIe Gen4, external DDR support, targeting midrange edge processing `[vendor-reported]` — sources: https://www.cnx-software.com/news/lattice/ and https://www.eetimes.com/lattice-expands-into-mid-range-fpgas/ `[secondary]`
- Avant-E parts target midrange FPGA edge-processing tasks (video, industrial automation, robotics) `[secondary]` — source: https://www.hackster.io/news/lattice-launches-avant-fpga-platform-avant-e-parts-for-mid-range-fpga-edge-processing-tasks-3d67e117b47e.amp `[secondary]`
- **Nexus** platform targets small, low-power devices; **Nexus 2 / Certus-N2** material reports **16G SerDes**, PCIe Gen4, enhanced security features, sampling status `[vendor-reported]` — (vendor product pages; tag as vendor-reported)
- Lattice vendor comparisons (e.g., ~2.5× lower power, ~2× performance vs unnamed competitor baselines) must be treated as **[vendor-reported]** with unnamed baselines — do not present as independent benchmarks `[unverified]`
- Lattice's historical strength: low-power small-form-factor FPGAs (iCE40, ECP5, MachXO) for control-plane, sensor aggregation, and bridging — relevant to data-center board-management and SmartNIC sidecar roles `[secondary]`

## 6. Microchip PolarFire

- PolarFire family: up to **500K logic elements**, **12.7 Gb/s SerDes**, nonvolatile **28 nm** SONOS fabric, hardened security (cryptoprocessor, DPA countermeasures), PCIe Gen2, DDR3/4 `[vendor-reported]` — sources: https://embeddedcomputing.com/technology/analog-and-power/pcbs-components/mouser-product-of-the-week-microchip-technology-polarfire-fpgas and https://ww1.microchip.com/downloads/en/DeviceDoc/Polarfire_SOC_Product_Overview.pdf `[vendor-reported]`
- **PolarFire SoC** adds a 5-core RISC-V (SiFive) subsystem: 4× RV64GC application cores + 1× RV64IMAC monitor core for embedded/edge compute `[official]`
- **PolarFire Ethernet Sensor Bridge rev 2.0** announced **2026-08-11**: reported **60% smaller** than prior revision, supports **4 cameras** and **10GbE-based sensor transport** for edge-AI sensor connectivity `[vendor-reported]` — source: https://www.stocktitan.net/news/MCHP/microchip-advances-edge-ai-sensor-connectivity-with-rev-2-0-polar-becysa7ddxek.html `[secondary]`
- PolarFire positioning: lowest-power midrange FPGA with superior thermal/power profile and supply-chain longevity (defense, industrial, space) `[vendor-reported]`; independent power comparisons vs Agilex/Versal were not found in this research — gap `[unverified]`

## 7. FPGA SmartNIC vendors

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

