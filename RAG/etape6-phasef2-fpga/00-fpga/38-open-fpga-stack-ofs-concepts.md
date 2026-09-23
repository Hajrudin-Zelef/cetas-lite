---
id: etape6-phasef2-fpga/00-fpga/38-open-fpga-stack-ofs-concepts
title: "38. Open FPGA Stack (OFS) — concepts"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "Baidu", "China", "Intel", "United States"]
dates: ["2026-09-22"]
keywords: ["accelerator", "amd", "compute", "cost", "ethernet", "gpus", "hbm", "hyperscaler", "intel", "memory", "pricing", "research"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [535, 585]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: bd1b4bed33f6aba2194bf780632f5a09ba6d9b43f0872d5cdb0b9421c4acffe2
---

# 38. Open FPGA Stack (OFS) — concepts

## 38. Open FPGA Stack (OFS) — concepts

- OFS is the open-source infrastructure (now under Altera docs) for building FPGA acceleration platforms: FIM (FPGA Interface Manager) + AFU (Accelerator Functional Unit) regions `[official]`
- **OPAE** (Open Programmable Acceleration Engine): software SDK/libraries and upstreamed kernel drivers for OFS platforms `[official]`
- Reference FIMs use a modular "take and tailor" approach for SmartNIC, vRAN, NFV, FSI workloads `[official]`
- OFS board catalog includes third-party cards (e.g., Hitek NC200) with OFS software + oneAPI/OpenCL BSP support `[official]`
- Intel's **IPU F2000X** was the first production OFS-based adapter (2023) `[secondary]`
- OFS vs vendor-proprietary shells: OFS trades some performance tuning for portability and open drivers `[unverified]`

## 39. AMD adaptive-computing portfolio notes

- **Spartan UltraScale+**: cost-optimized AMD FPGA family; added to Vivado 2025.1 device support `[secondary]`
- Legacy families still in the field: Artix-7, Kintex-7/UltraScale, Virtex UltraScale+, Zynq-7000/UltraScale+ MPSoC — relevant for installed-base and secondary-market pricing `[secondary]`
- **Zynq UltraScale+ MPSoC**: Arm + FPGA SoC widely used in embedded vision, industrial, and aerospace (predecessor concept to Versal) `[secondary]`
- Versal spans AI Core / AI Edge / Prime / Premium / HBM / Premium Gen 2 — segmentation by compute, power, and memory bandwidth `[vendor-reported]`
- Vivado 2026.1 tiered licensing: new commercial model for Versal-era tools (details beyond the headline not captured) `[secondary]`

## 40. Networking standards context for FPGA NICs

- **400G Ethernet** is now the mainstream FPGA SmartNIC target (NT400, N3070X, V80 QSFP56) `[vendor-reported]`
- **800G** shells exist in the Altera OFS ecosystem context (per scope); 1.6T Ethernet is the next industry milestone — FPGA SerDes at 112G PAM4 today, 224G in development industry-wide `[unverified]`
- **Ultra Ethernet (UEC)**: noted on Napatech N3070X as a roadmap item — relevant to AI-scale networking `[secondary]`
- **CXL 2.0/3.1**: Agilex 7 M-Series (CXL 2.0), N3070X (CXL 2.0 mount option), Versal Premium Gen 2 (CXL 3.1) — FPGA as CXL-attached memory/compute expansion `[vendor-reported]`
- **PTP/1588** hardware timestamping: table stakes for FPGA NICs (Hitek NC200, Corundum, ExaNIC) `[official]`

## 41. Security features on data-center FPGAs

- Bitstream encryption/authentication: standard on Agilex, Versal, Nexus 2, PolarFire `[official]`
- **Secure boot + HW RoT**: N3070X offers HW RoT option; PolarFire/Mach-N2 emphasize RoT for board control `[secondary]`
- **Post-quantum crypto**: Lattice Mach-N2 ships CNSA 2.0 PQC (ML-DSA, ML-KEM, LMS, XMSS) — first FPGA family observed with standardized PQC `[vendor-reported]`
- AES-GCM-256/128 hard crypto engines on Versal Premium Gen 2 (800 Gb/s line-rate claim) `[vendor-reported]`
- Side-channel resistance and TRNG on Nexus 2; DPA countermeasures on PolarFire `[official]`
- Relevance: links FPGA track to Phase F3 (RoT/TPM/HSM) — FPGAs increasingly carry platform-security functions `[unverified]`

## 42. Power and thermal envelopes (cards)

- Alveo SN1000: **75 W** (PCIe edge powered), FHHL `[vendor-reported]`
- Napatech N3070X: max **150 W**, passive cooling, PCIe 150W-ATX aux connector `[secondary]`
- Alveo V80: **190 W** `[vendor-reported]`
- Hitek NC200: 75 W (edge) / ~100 W (6-pin) variants `[official]`
- General: FPGA accelerator cards span 75–300 W like GPUs; perf/W claims are workload-specific and vendor-reported `[unverified]`

## 43. Methodology and limitations of this file

- Research conducted 2026-09-22 via public web search snippets and opened pages; no vendor briefings, no purchases, no logins `[unverified]`
- Distributor/broker prices (eBay, global-ic) are snapshots, not OEM list prices, and vary by seller/condition/quantity `[secondary]`
- Vendor performance claims (TOPs, power reduction, user-density) are reported as `[vendor-reported]` with baselines where known; none independently benchmarked here `[unverified]`
- Planned dates (shipments, sampling, production) are explicitly marked; "current" means "as observed 2026-09-22" `[unverified]`
- Chinese-market FPGA vendors (Gowin, Efinix is TW/US) and hyperscaler internal programs are lightly covered — known gaps `[unverified]`
- Accolade Technology, RT PolarFire detail, and Alibaba/Baidu/Tencent 2026 FPGA status remain open gaps `[unverified]`

