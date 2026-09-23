---
id: etape6-phasef2-fpga/00-fpga/21-fpga-smartnic-ecosystem-additional-vendors-and-notes
title: "21. FPGA SmartNIC ecosystem — additional vendors and notes"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Alibaba", "Baidu", "Intel", "Microsoft", "United States"]
dates: ["2025-12-20"]
keywords: ["acquisition", "agentic", "amd", "asic", "aws", "distribution", "gpu", "hbm", "hyperscaler", "inference", "intel", "latency"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [336, 385]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: 6ebb1288857710a2c8dc56f1dc4788a17adc07986dca8bd8d40c4bede6162dd1
---

# 21. FPGA SmartNIC ecosystem — additional vendors and notes

## 21. FPGA SmartNIC ecosystem — additional vendors and notes

- **Intel/Altera IPU F2000X**: first production adapter based on Intel's IPU platform released to open source as part of OFS (2023) — historical anchor for the Intel/Altera infrastructure-processing roadmap `[secondary]` — source: https://it-online.co.za/2023/09/18/intel-expands-fpga-portfolio/ `[secondary]`
- **Silicom + Napatech**: Silicom integrates Napatech FPGA SmartNIC technology into white-box server adapters (OEM/ODM channel) `[secondary]`
- **Accolade Technology** (gap, restated): FPGA-based ANIC adapters for lossless packet capture, filtering, deduplication, and inline security; no 2026 portfolio/pricing captured — needs dedicated research `[unverified]`
- **Cisco Nexus SmartNIC/NIC (ex-ExaNIC)**: ultra-low-latency NICs for HFT and deterministic packet capture; ExaLINK → Nexus 3550 switching line `[official]`
- **BittWare**: boards across AMD/Xilinx, Altera, and Achronix silicon; SmartNIC Shell adds DPDK, match-action, P4/SDNet, timestamping (see §7) `[vendor-reported]`
- **Hitek NC200** (Agilex, OFS): SmartNIC/data-center/HPC/vRAN target; PTP/1588 with external clock/sync inputs `[official]`
- "Napatech (now Intel)" correction (restated for RAG retrieval): **no acquisition found**; Napatech remains independent and partners with Intel/Altera on Agilex-based cards `[unverified]`

## 22. P4, DPDK, and packet-processing — extended

- **OpenNIC** (AMD open-source NIC shell): QDMA-based, DPDK poll-mode driver available, P4 via SDNet-generated pipelines `[independent]`
- **Corundum** (§8): custom DMA + Linux driver, 10k+ queues, PTP; research platform from UCSD (Corundum paper) with ongoing 2026 ecosystem activity `[independent]`
- **P4→FPGA via SDNet**: AMD/Xilinx SDNet compiles P4_14/P4_16 to FPGA IP; BittWare SmartNIC Shell exposes it for 100G NIC builds `[vendor-reported]`
- **DPDK on FPGA**: OpenNIC DPDK PMD, BittWare SmartNIC Shell DPDK support, Corundum DPDK integration — DPDK is the common software API across FPGA NIC shells `[independent]`
- **Match-action processing**: BittWare SmartNIC Shell and Napatech Link-* suites expose programmable match-action pipelines without full RTL design `[vendor-reported]`
- **PTP/1588**: hardware timestamping standard on FPGA NICs (Hitek NC200 full sync support; Corundum PTP) `[official]`
- **vDPA/virtio**: SN1000 supports virtio and ef100 with vDPA for virtualized data paths `[vendor-reported]`

## 23. Hyperscaler and telco adoption — extended notes

- **Microsoft Catapult v2 / Azure Boost**: MANA FPGA acceleration under Azure Boost offloads networking/storage/virtualization; later strategy adds in-house DPU silicon and broader AMD Pensando SmartNIC use — heterogeneous offload estate, not FPGA-only `[secondary]`
- **AWS F1** (historical spec): f1.2xlarge = 1 FPGA; f1.16xlarge = 8 FPGAs, 64 GiB DDR4 ECC per FPGA, 400 Gb/s bidirectional ring; AFI (Amazon FPGA Image) distribution model `[official]`; EOL claim 2025-12-20 **unverified** — check https://github.com/aws/aws-fpga `[unverified]`
- **Alibaba / Tencent / Baidu**: FPGA acceleration historically used in search/recommendation and networking; no 2026 status captured — gap `[unverified]`
- **Open RAN / vRAN**: FPGAs used for L1/L2 offload (LDPC/polar coding, beamforming, fronthaul eCPRI): Altera Agilex (NEC, MTI, Wisig), AMD Versal, Achronix Speedster7t all claim ORAN/vRAN sockets `[vendor-reported]`
- **5G UPF**: Napatech NT400 offload claims 90×/7× user-density advantages (vendor analysis) — see §18 `[vendor-reported]`
- **HFT**: Alveo UL3524, Exegy/NovaSparks/Enyx consolidation, ExaNIC/Nexus SmartNICs — FPGA remains the standard for sub-microsecond trading pipelines `[secondary]`
- **Emulation/prototyping**: VP1902 (18.5M LCs) targets pre-silicon verification of AI/automotive/5.0 ASICs — a major FPGA demand driver independent of deployment acceleration `[vendor-reported]`

## 24. Open-source tooling — extended

- **Yosys + nextpnr**: the standard open synthesis/P&R pair; 2026 commits active across the ecosystem `[independent]` — source: https://github.com/vibeic/awesome-open-ic/blob/HEAD/docs/design-tools.md `[secondary]`
- Bitstream projects: **IceStorm** (iCE40), **Trellis** (ECP5), **X-Ray** (Xilinx 7-series), **Apicula** (Gowin) `[independent]`
- **openFPGALoader 1.1.1**: universal programmer (Xilinx, Altera/Intel, Lattice, Gowin, Efinix, Anlogic, Cologne Chip); Linux/Windows/macOS/OpenBSD; Intel/Altera SVF+RBF via `quartus_cpf` `[independent]` — sources: https://github.com/trabucayre/openfpgaloader/blob/HEAD/README.md and https://github.com/trabucayre/openfpgaloader/blob/HEAD/doc/vendors/intel.rst `[independent]`
- **LiteX / Migen**: Python-based SoC builder generating RTL for open and vendor flows; widely used with Lattice ECP5 and Xilinx 7-series in open-hardware projects `[independent]`
- **Corundum** as open NIC IP: production-grade 100G FPGA NIC framework with driver — the most deployment-relevant open FPGA networking artifact `[independent]`
- Distro lag: Ubuntu 24.04 ships substantially older nextpnr/openFPGALoader than upstream (0.11.1 / 1.1.1 noted Aug 2026) — build from source for current devices `[secondary]` — source: https://github.com/machdyne/zeitlos/blob/HEAD/docs/toolchain.md `[secondary]`
- **Hard limits of open flows**: no hardened IP (400G MACs, PCIe Gen5/6, HBM, CXL), no vendor timing signoff, no encrypted production bitstreams, no support for Agilex/Versal/Speedster7t/Avant/PolarFire — vendor tools mandatory for production on high-end families `[unverified]`
- **Lattice Prompt** (Sep 2026) is the first vendor AI/MCP-based FPGA design assistant observed — agentic IDE → Radiant orchestration; 10× productivity vendor-claimed `[vendor-reported]`

## 25. Pricing and TCO — extended snapshots

- Alveo V80: **US$9,495** MSRP `[vendor-reported]`
- Alveo U280/U250/U200/U30: launch vs secondary-market spreads show 3–10× depreciation on older cards (U250: $8,995 launch → ~$2,250–5,479 resale) `[secondary]`
- Agilex 7 AGF027 distributor reference: ~**US$20k–37.8k** per 3pcs (broker, Hong Kong) — illustrates high-end FPGA silicon pricing; not OEM MSRP `[secondary]`
- No OEM list prices found for: Agilex 5/7/9 dev kits, Versal VEK280/VCK190/VMK180, Lattice Nexus 2/Avant kits, PolarFire kits, VectorPath S7t-VG6 — distributor quote only `[unverified]`
- TCO arguments heard from vendors: CPU-core reclaim (SmartNIC offload), power reduction (NEC >40%), ASIC-NRE avoidance, deterministic latency vs GPU — all `[vendor-reported]`, none independently benchmarked in this pass `[unverified]`
- Missing: any 2026 independent $/Gbps, $/inference, or perf/W study across Agilex/Versal/Speedster7t — **research gap** `[unverified]`

