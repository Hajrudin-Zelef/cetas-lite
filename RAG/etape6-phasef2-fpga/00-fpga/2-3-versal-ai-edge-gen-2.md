---
id: etape6-phasef2-fpga/00-fpga/2-3-versal-ai-edge-gen-2
title: "2.3 Versal AI Edge Gen 2"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: reference
actors: ["United States"]
dates: []
keywords: ["accelerator", "amd", "asic", "compute", "dsp", "gpu", "hbm", "memory"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [55, 78]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: 507aaae47f0694deecbfc04f5d03ba51dfdd63ece033d0780bf3bee8284a45ee
---

# 2.3 Versal AI Edge Gen 2

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

