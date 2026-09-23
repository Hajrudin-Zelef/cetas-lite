---
id: etape6-phasef2-fpga/00-fpga/11-tco-fpga-vs-asic-vs-gpu
title: "11. TCO: FPGA vs ASIC vs GPU"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Intel", "Microsoft"]
dates: ["2025-12-20", "2026-01-14"]
keywords: ["asic", "gpu", "accelerator", "acquisition", "amd", "aws", "benchmarks", "dsp", "gpus", "hbm", "inference", "intel"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [180, 234]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: c6526dd8578fe389950cb16033219c601fe516d2d4fc8b714c4b5c96ff7ca225
---

# 11. TCO: FPGA vs ASIC vs GPU

## 11. TCO: FPGA vs ASIC vs GPU

- FPGA value proposition per vendors: reconfigurability avoids ASIC NRE and respin risk; lower power than GPUs for fixed-function pipelines (e.g., NEC >40% power reduction claim on Agilex 7 radios) `[vendor-reported]`
- FPGA weaknesses vs ASIC: ~10–35× area/power overhead is the textbook range for equivalent logic — widely cited in literature but **no 2026 Agilex/Versal-specific independent comparison** was collected in this research `[unverified]`
- FPGA vs GPU for AI inference: Versal AI Edge/AIE-ML and Agilex DSP/AI blocks target deterministic low-latency inference; GPUs dominate throughput-per-dollar on dense training — workload-dependent, no single crossover point established here `[unverified]`
- SmartNIC TCO angle: FPGA SmartNICs (SN1000/U45, Napatech, Exablaze) offload vSwitch, crypto, and telemetry from host CPUs, reclaiming CPU cores for revenue workloads — the standard vendor TCO argument `[vendor-reported]`
- **Gap**: no independent 2026 TCO study (FPGA vs ASIC vs GPU, $/Gbps or $/inference) was found in this research pass; treat vendor TCO claims as marketing until independently benchmarked `[unverified]`

## 12. Adoption evidence

### 12.1 Microsoft — Catapult / Azure Boost

- Microsoft's **Catapult** program pioneered FPGA acceleration in Azure (Bing ranking, networking) on Altera/Intel FPGAs `[secondary]`
- **Azure Boost** (2023): FPGA-powered **MANA** network acceleration offloads networking/storage/virtualization from host CPUs `[secondary]` — source: https://www.theregister.com/off-prem/2023/11/21/microsoft-adds-fpga-powered-network-accelerator-to-azure/969552?td=keepreading `[secondary]`
- Later Microsoft strategy also includes an internally designed DPU and expanding AMD Pensando-based deployments — do **not** imply the entire 2026 Azure Boost estate is FPGA-only `[secondary]` — source: https://www.nextplatform.com/cloud/2026/07/20/microsoft-taps-amd-for-at-scale-ai-cpu-and-gpu-clusters/5275161 `[secondary]`

### 12.2 AWS FPGA — F1 status

- Historical F1 (2017 GA): up to **8 Xilinx UltraScale+ FPGAs** per f1.16xlarge, **64 GiB DDR4 ECC per FPGA**, **400 Gb/s** bidirectional ring interconnect `[official]` — source: https://aws.amazon.com/about-aws/whats-new/2017/04/amazon-ec2-f1-instances-customizable-fpgas-for-hardware-acceleration-are-now-generally-available/ `[official]`
- A third-party GitHub doc (coderhard/hngac-fpga, observed 2026) states **F1 reached end of life on 2025-12-20** with only grandfathered customers retaining launch eligibility — **not yet verified against AWS's official aws-fpga repository**; treat as `[unverified]` until confirmed `[unverified]` — source: https://github.com/coderhard/hngac-fpga/blob/HEAD/docs/setup-aws-fpga-ami.md `[secondary]`
- **Action**: verify F1 EOL against https://github.com/aws/aws-fpga before presenting as definitive `[unverified]`

### 12.3 Telecom / Open RAN

- NEC (Agilex 7, sub-6GHz Massive MIMO, >40% power cut), MTI (Agilex 5 E-Series, 4T4R macro RU, 25GbE fronthaul), Wisig (32TR n78 Massive MIMO ULPI radio + 8T8R n78 macrocell, India) — Altera 5G Advanced/6G radio momentum, Feb 2026 `[vendor-reported]`
- FPGAs (Agilex, Versal, Speedster7t) are used for L1/L2 offload, fronthaul, beamforming, and channel coding in vRAN/Open RAN architectures `[vendor-reported]`

### 12.4 High-frequency trading / ultra-low latency

- AMD Alveo UL3524 purpose-built for electronic trading (see §3.6) `[vendor-reported]`
- Exegy–NovaSparks merger (2026-01-14) consolidates FPGA market-data/normalization products `[secondary]`
- Cisco/Exablaze ExaNIC → Nexus SmartNIC/NIC for ultra-low-latency trading `[official]`

## 13. Open-source FPGA tooling

- **Yosys** (synthesis) + **nextpnr** (place-and-route) form the standard open flow; active 2026 commits observed across related projects `[independent]` — source: https://github.com/vibeic/awesome-open-ic/blob/HEAD/docs/design-tools.md `[secondary]`
- Family support via reverse-engineered bitstream projects: **Project IceStorm** (Lattice iCE40), **Project Trellis** (Lattice ECP5), **Project X-Ray** (Xilinx 7-series), **Project Apicula** (Gowin) `[independent]`
- nextpnr targets: iCE40, ECP5, Gowin, and some Xilinx 7-series devices — **no equivalent open flow for Agilex, Versal, Speedster7t, Avant, or PolarFire** `[independent]`
- **openFPGALoader** 1.1.1 (Aug 2026 project note) programs many boards/cables/FPGAs (Xilinx, Altera/Intel, Lattice, Gowin, Efinix, Anlogic, Cologne Chip) on Linux/Windows/macOS/OpenBSD; supports SVF/RBF flows for Intel/Altera via `quartus_cpf` `[independent]` — sources: https://github.com/trabucayre/openfpgaloader/blob/HEAD/README.md and https://github.com/trabucayre/openfpgaloader/blob/HEAD/doc/vendors/intel.rst `[independent]`
- Ubuntu 24.04 distro packages for nextpnr/openFPGALoader lag upstream substantially (nextpnr 0.11.1 / openFPGALoader 1.1.1 noted in an Aug 2026 project doc vs older distro packages) `[secondary]` — source: https://github.com/machdyne/zeitlos/blob/HEAD/docs/toolchain.md `[secondary]`
- **Corundum** (open-source 100G FPGA NIC framework, §8) demonstrates production-grade open FPGA networking IP `[independent]`
- **Limitations of open flows** (vs vendor tools): no support for hardened IP blocks (100G/400G MACs, PCIe Gen5/6, HBM controllers), no vendor timing models/signoff, no encrypted/signed production bitstreams, limited large-device capacity — open flows are viable for small/midrange Lattice/Gowin/Xilinx-7s designs, not for Agilex 7/9 or Versal-class production `[unverified]`
- Vendor flows remain mandatory for production on all high-end families covered here `[unverified]`

## 14. Conflicts, gaps, and uncertainty

- "Napatech (now Intel)" as seen in some briefs: **no acquisition evidence found**; do not repeat `[unverified]`
- AWS F1 EOL 2025-12-20: single third-party GitHub source; **needs official AWS confirmation** `[unverified]`
- Versal Premium Gen 2 production H2 2026 and MoP sampling end-2026/production H2-2027: **planned dates per vendor reporting**, not confirmed shipments `[vendor-reported]`
- Agilex 9 Direct RF production Q3 2026: **planned**, not confirmed shipped `[vendor-reported]`
- Accolade Technology: named in scope; detailed 2026 portfolio/pricing not yet researched — **gap** `[unverified]`
- Current OEM/dev-kit pricing for Agilex, Versal eval kits, Lattice, PolarFire, Achronix: **gap** (distributor-quote only) `[unverified]`
- Independent 2026 benchmarks (Agilex 9 vs Versal Premium Gen 2 vs Speedster7t; FPGA vs ASIC vs GPU TCO): **gap** — only vendor-reported figures collected `[unverified]`
- Lattice Avant/Nexus 2 detailed specs beyond the ranges given: verify against lattice.com datasheets before quoting exact SKUs `[unverified]`

