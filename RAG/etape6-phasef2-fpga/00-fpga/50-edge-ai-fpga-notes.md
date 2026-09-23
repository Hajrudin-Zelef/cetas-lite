---
id: etape6-phasef2-fpga/00-fpga/50-edge-ai-fpga-notes
title: "50. Edge-AI FPGA notes"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Alibaba", "Baidu", "China"]
dates: []
keywords: ["amd", "asic", "aws", "benchmark", "compute", "ethernet", "gpu", "hbm", "hyperscaler", "inference", "mcp", "memory"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [634, 689]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: b35a54bef92273169b50e735645fd6199a6ea40cfc1a9818fd0549596caa2bee
---

# 50. Edge-AI FPGA notes

## 50. Edge-AI FPGA notes

- **Versal AI Edge Gen 2**: AIE-ML engines for edge inference; VEK280 eval board (xcve2802) `[secondary]`
- **Lattice sensAI** stack: edge-AI solution stack on Nexus/Avant (person detection, object counting class workloads) `[vendor-reported]`
- **PolarFire SoC + SmartHLS**: low-power deterministic inference at the far edge `[official]`
- **PolarFire Ethernet Sensor Bridge rev 2.0**: 4-camera 10GbE sensor aggregation for edge AI (Aug 2026) `[vendor-reported]`
- FPGAs compete at the edge on determinism and sensor-fusion flexibility vs NPUs/ASICs on raw TOPS/W `[unverified]`

## 51. Glossary — acronyms used in this file

- ACAP: Adaptive Compute Acceleration Platform (AMD Versal marketing term) `[official]`
- AFU: Acceleration Functional Unit (OFS workload region) `[official]`
- AIE / AIE-ML: AI Engine / AI Engine-ML (Versal hardened vector/AI tiles) `[official]`
- ALM: Adaptive Logic Module (Altera logic block) `[official]`
- BSP: Board Support Package `[official]`
- CXL: Compute Express Link `[official]`
- DFX: Dynamic Function eXchange (partial reconfiguration) `[official]`
- DPDK: Data Plane Development Kit `[independent]`
- FIM: FPGA Interface Manager (OFS shell) `[official]`
- HBM: High Bandwidth Memory `[official]`
- HFT: High-Frequency Trading `[secondary]`
- HLS: High-Level Synthesis `[official]`
- LUT: Look-Up Table `[official]`
- MCP: Model Context Protocol (open standard linking LLMs to tools) `[independent]`
- MLP: Machine Learning Processor (Achronix Speedster7t hardened block) `[official]`
- MoP: Memory on Package `[vendor-reported]`
- NoC: Network on Chip `[official]`
- OPAE: Open Programmable Acceleration Engine (OFS SDK) `[official]`
- OFS: Open FPGA Stack `[official]`
- P4: Programming Protocol-independent Packet Processors language `[independent]`
- PQC: Post-Quantum Cryptography `[official]`
- PTP: Precision Time Protocol (IEEE 1588) `[official]`
- QDMA: Queue-based DMA (AMD/Xilinx hardened IP) `[official]`
- RoT: Root of Trust `[official]`
- SDNet: AMD/Xilinx P4-to-FPGA compiler toolchain `[official]`
- SerDes: Serializer/Deserializer `[official]`
- SLC: System Logic Cell (Lattice) `[official]`
- UEC: Ultra Ethernet Consortium `[secondary]`
- UPF: User Plane Function (5G packet core) `[official]`
- vDPA: vHost Data Path Acceleration `[official]`
- vRAN / ORAN: virtualized / Open Radio Access Network `[secondary]`
- XCVR: Transceiver `[official]`

## 52. Open questions for future research passes

- Confirm AWS F1 EOL against the official aws/aws-fpga repository and capture any successor (F2?) `[unverified]`
- Capture Accolade Technology 2026 portfolio, pricing, and software in detail `[unverified]`
- Verify Versal Premium Gen 2 and Agilex 9 Direct RF production-shipment status post-planned dates `[unverified]`
- Obtain OEM/dev-kit list prices: Agilex/Quartus kits, Versal eval kits, Lattice Nexus 2/Avant kits, PolarFire kits, VectorPath cards `[unverified]`
- Find an independent 2026 benchmark or TCO study (FPGA vs ASIC vs GPU) `[unverified]`
- Check NIST CMVP for Nexus 2 FIPS 140-3 Level 2 certification status `[unverified]`
- Track Lattice Prompt adoption and any AMD/Altera AI-assistant responses `[unverified]`
- Track Chinese hyperscaler (Alibaba, Tencent, Baidu) 2026 FPGA programs `[unverified]`
- Clarify Alveo U45 specifications and availability (successor positioning vs SN1000) `[unverified]`
- Watch 800G/1.6T FPGA SmartNIC shells and UEC adoption on FPGA NICs `[unverified]`

