---
id: etape6-phasef2-fpga/00-fpga/34-microsoft-catapult-and-fpga-dnn-acceleration
title: "34. Microsoft Catapult and FPGA DNN acceleration"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Intel", "Microsoft"]
dates: ["2025-04-01", "2025-09-12", "2025-12-20", "2026-01-14", "2026-02-26", "2026-05-19", "2026-05-29", "2026-06-09", "2026-07-06", "2026-08-11", "2026-09-16", "2026-09-22"]
keywords: ["acquisition", "amd", "aws", "ethernet", "hbm", "hyperscaler", "inference", "intel", "latency", "lpddr5x", "mcp"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [479, 534]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: e763664e405fcb16b67622cc1dd4752c59ba8b023d6aa07000399b1a0e72df74
---

# 34. Microsoft Catapult and FPGA DNN acceleration

## 34. Microsoft Catapult and FPGA DNN acceleration

- **Project Catapult** (2010s): FPGAs in every new Azure server for Bing ranking acceleration — the original hyperscale FPGA deployment `[secondary]`
- **Catapult v2 / Brainwave**: Stratix 10 FPGAs for DNN inference (e.g., ResNet-50) with ms-scale latency claims `[secondary]`
- **Azure Boost + MANA** (2023): FPGA-powered network/storage/virtualization offload; current estate is heterogeneous (FPGA + in-house DPU + Pensando) `[secondary]`
- Takeaway: Microsoft remains the canonical hyperscaler FPGA adopter, but the offload mix has diversified beyond pure FPGA `[secondary]`

## 35. Open-source tooling — extended inventory

- **Yosys**: open synthesis (Verilog → netlist); the foundation of open FPGA flows `[independent]`
- **nextpnr**: open place-and-route; targets iCE40, ECP5, Gowin, Xilinx 7-series (via X-Ray) `[independent]`
- **Project IceStorm / Trellis / X-Ray / Apicula**: reverse-engineered bitstream documentation enabling open flows per family `[independent]`
- **F4PGA (ex-SymbiFlow)**: open toolchain targeting Xilinx 7-series via X-Ray database `[independent]`
- **GHDL**: open VHDL simulator; **Verilator**: fast Verilog/SystemVerilog simulator; **cocotb**: Python testbench framework — all standard in open FPGA verification `[independent]`
- **LiteX / Migen**: Python SoC builder producing portable RTL; popular with ECP5/iCE40 open-hardware boards (OrangeCrab, Fomu, TinyFPGA) `[independent]`
- **openFPGALoader 1.1.1**: universal JTAG/SPI programmer across vendors `[independent]`
- **Corundum**: the most deployment-relevant open FPGA IP — 100G NIC framework with Linux driver `[independent]`
- **What open flows cannot do**: hardened 100G+ MACs, PCIe Gen5/6, HBM/CXL, vendor timing signoff, encrypted bitstreams, Agilex/Versal-class devices `[unverified]`
- Ubuntu 24.04 distro packages lag upstream (build from source for current support) `[secondary]`

## 36. Timeline — key FPGA events (2025-01 → 2026-09-22)

- 2025-04-01: Altera announces Agilex 7 **M-Series** in production, shipping worldwide, lifecycle to 2035 `[vendor-reported]`
- 2025-09-12/15: Silver Lake closes **51%** Altera acquisition; Intel retains 49% `[secondary]`
- 2026-01-14: **Exegy acquires NovaSparks** (FPGA trading consolidation) `[secondary]`
- 2026-02-26: Altera telecom push — NEC (Agilex 7), MTI (Agilex 5 E-Series), Wisig radios `[vendor-reported]`
- 2026-05-19: Achronix Speedster7t datasheet **v2.3** released `[official]`
- 2026-05-29: AMD launches Versal Prime Gen 2 compact SKUs (2VM3104/3254/3454) `[secondary]`
- 2026-06-09: Altera introduces **Agilex 9 Direct RF** (AGRW039); production planned Q3 2026 `[vendor-reported]`
- 2026-07-06: AMD Versal Premium Gen 2 **MoP** (32 GB LPDDR5X) announced; sampling end-2026 planned `[vendor-reported]`
- 2026-08-11: Microchip **PolarFire Ethernet Sensor Bridge rev 2.0** announced `[vendor-reported]`
- 2026-09-16: Lattice launches **Mach-N2** (PQC secure control) and **Lattice Prompt** (AI/MCP design assistant) `[secondary]`
- Undated (observed 2026): AMD steers new SmartNIC designs from SN1000 family to **Alveo U45** `[official]`
- Claimed but unverified: AWS F1 EOL **2025-12-20** (single third-party source) `[unverified]`

## 37. SKU / part-number index (quick reference)

- **AGRW039**: Altera Agilex 9 Direct RF flagship SoC FPGA (3.9M LEs, 64 GSa/s RF) `[vendor-reported]`
- **AGF027** (AGFA/AGFB variants): Altera Agilex 7 F-Series top device, 2.7M LEs, up to 58G XCVRs `[official]`
- **AGI027**: Altera Agilex 7 I-Series, 2.7M LEs, up to 112G XCVRs, R-Tile PCIe Gen5 `[official]`
- **AGIC041**: Intel/Altera Agilex device in Napatech N3070X (1× F-tile, 3× R-tile, 4M LEs) `[secondary]`
- **2VP3422 / 2VP3522 / 2VP3622**: AMD Versal Premium Gen 2 (+ MoP variants with 32 GB LPDDR5X) `[vendor-reported]`
- **2VM3104 / 2VM3254 / 2VM3454 / 2VM3654**: AMD Versal Prime Gen 2 compact SoCs `[vendor-reported]`
- **VC1902 / VC2802**: AMD Versal AI Core (400 AI Engines class; 405 INT8x4 TOPs on VC2802) `[secondary]`
- **XCVC1902**: Versal AI Core device on VCK5000 inference card `[secondary]`
- **XCV80**: Versal HBM device on Alveo V80 (32 GB HBM2e) `[vendor-reported]`
- **XCVH1782**: Versal HBM device, up to 819 GB/s HBM2e `[vendor-reported]`
- **VP1902**: Versal Premium emulation SoC, 18.5M logic cells `[vendor-reported]`
- **XQRVC1902**: radiation-tolerant Versal AI Core for space (Class B) `[secondary]`
- **AC7t800**: Achronix Speedster7t flagship (711K LEs, 864 MLPs) `[vendor-reported]`
- **LN2-CTxx**: Lattice Certus-N2 ordering family (Nexus 2); LN2-CT20E/CT16E variants `[official]`
- **A-SN1022-P4**: AMD Alveo SN1000 SmartNIC part number `[official]`
- **N3070XP12 / N3070XP34**: Napatech 400G SmartNIC variants (FH ½ / ¾ length) `[secondary]`
- **xcve2802-vsvh1760-2MP-e-S**: Versal AI Edge device on VEK280 eval board `[secondary]`
- **XCU26**: Xilinx UltraScale+ FPGA powering Alveo SN1000 `[vendor-reported]`

