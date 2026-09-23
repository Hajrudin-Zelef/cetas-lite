---
id: etape6-phasef2-fpga/00-fpga/44-cross-references-to-other-phase-f-files
title: "44. Cross-references to other Phase F files"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: reference
actors: ["AMD", "Intel", "United States"]
dates: []
keywords: ["amd", "gpu", "intel", "memory", "quantization", "research", "throughput", "training"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [586, 633]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: b691ab7ff6bfa54b1956cbf5b911011a7b80dd1c0896b45560bf79d9a567d20f
---

# 44. Cross-references to other Phase F files

## 44. Cross-references to other Phase F files

- F1 (NIC/DPU/SmartNIC): overlaps on SN1000/U45, Napatech, Exablaze, Pensando, MANA/Azure Boost — F2 covers the FPGA silicon and cards; F1 covers the NIC/DPU product view `[unverified]`
- F3 (RoT/TPM/HSM/TEE): Mach-N2 PQC, PolarFire/N3070X HW RoT, bitstream security — FPGA security features feed F3 `[unverified]`
- F4 (virtio/SR-IOV/DPDK/SPDK): OpenNIC/Corundum DPDK PMDs, vDPA on SN1000, SDNet/P4 — software data-plane complements in F4 `[unverified]`

## 45. Altera software stack

- **Quartus Prime**: Altera's FPGA design software (Lite/Standard/Pro editions historically); Pro targets Agilex/Stratix 10 high-end flows `[official]`
- **Quartus Prime 25.1** validated with Serial Lite IV IP 5.5.2 (Apr 2025) `[official]`
- **oneAPI Base Toolkit + OpenCL BSP**: heterogeneous programming path for OFS boards (Hitek NC200 lists oneAPI/OpenCL BSP support) `[official]`
- **Nios V**: Altera's RISC-V soft processor family (m/g/c variants) for control-plane logic in fabric `[official]`
- Only Quartus-generated bitstreams load on production Altera devices — no third-party bitstream path `[official]`

## 46. AMD software ecosystem for Versal/Alveo

- **Vivado ML**: RTL design suite; 2025.1 added Spartan UltraScale+ and next-gen Versal; 2026.1 adds tiered licensing, NoC/QoR improvements `[secondary]`
- **Vitis Unified**: software platform for Arm/AIE application development; 2026.1 bundles Vivado 2026.1 `[secondary]`
- **Vitis AI**: quantization/compiler/runtime stack for deploying DNNs to Versal AI Engines and Alveo cards `[official]`
- **FINN**: open-source (Xilinx Research) dataflow compiler for quantized neural networks on FPGAs; integrated with Alveo UL3524-class flows `[independent]`
- **PYNQ**: Python-on-Zynq/Versal open-source framework for rapid prototyping on Arm+FPGA SoCs `[independent]`
- **Vitis Networking P4 (VitisNetP4)**: P4-programmable networking IP for Alveo/Versal pipelines `[official]`
- **Embedded Development Framework (EDF)**: Yocto-based Linux, drivers for video/GPU/fast memory on Versal Prime Gen 2 `[vendor-reported]`

## 47. FPGA market structure context

- The FPGA market is a duopoly: **AMD (Xilinx) and Altera** together hold the large majority of data-center/high-end FPGA share; Lattice leads low-power/small, Microchip in secure/midrange, Achronix in high-bandwidth niches `[secondary]`
- Data-center FPGA demand drivers: SmartNIC/DPU offload, video transcode, trading, emulation/prototyping, telco L1 — not general-purpose AI training `[secondary]`
- FPGA competition with structured ASICs/eASICs (e.g., Intel's former eASIC line) and with ASSPs for networking offload `[unverified]`
- Supply longevity matters: AMD 15-year lifecycle (Versal Premium Gen 2 MoP), Altera M-Series to 2035 — industrial/defense procurement drivers `[vendor-reported]`
- No 2026 market-size dollar figure is quoted here — treat any such figure found elsewhere as needing a named analyst source `[unverified]`

## 48. Emulation and prototyping (VP1902 context)

- Pre-silicon emulation is a major high-end FPGA demand driver independent of data-center deployment `[secondary]`
- **VP1902**: 18.5M logic cells / 60B gates, 5.6 Tbps measured throughput, 8× faster debug vs VU19P (vendor claim), no AI Engines `[vendor-reported]`
- EDA partners: Cadence, Siemens, Synopsys flows supported on VP1902 `[vendor-reported]`
- Competing emulation platforms (Synopsys HAPS/ZeBu, Cadence Palladium/Protium) use custom or FPGA-based engines — context for VP1902's market `[secondary]`
- Sampling to early-access customers began Q3 of launch year; production followed H1 next year `[vendor-reported]`

## 49. Defense, aerospace, and rugged FPGAs

- **XQR Versal AI Core**: Class B space qualification, Raytheon customer, reprogrammable on-orbit `[secondary]`
- **RT PolarFire**: radiation-tolerant Microchip family for space (detail gap) `[unverified]`
- Agilex 7/9 target radar, electronic warfare, and secure comms (Direct RF series explicitly) `[vendor-reported]`
- Versal Premium Gen 2 MoP targets aerospace/defense: secure comms, LEO satellites, radar; extended/industrial temperature grades `[vendor-reported]`
- CNSA 2.0 PQC (Mach-N2) aligns with US government post-quantum procurement mandates converting security into a buying gate `[secondary]`

