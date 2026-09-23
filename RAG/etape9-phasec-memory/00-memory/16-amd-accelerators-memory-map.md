---
id: etape9-phasec-memory/00-memory/16-amd-accelerators-memory-map
title: "16. AMD accelerators: memory map"
domain: step-9-phase-c-server-accelerator-memory-hardware
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Cerebras", "Falcon", "Google", "Intel", "Meta", "Microsoft", "Nvidia", "Samsung"]
dates: []
keywords: ["accelerator", "amd", "memory", "aws", "blackwell", "compute", "cost", "custom silicon", "disaggregated", "fp4", "fp8", "gpu"]
source: docs/RAG/etape9_phaseC_memory.md
source_anchor: ""
source_lines: [295, 355]
section: "Step 9 — Phase C: Server & Accelerator Memory Hardware"
sha256: 0447f4619f1d7aec4a3b064fee68ce72723cfe86d2a043c5c156718fae1c620a
---

# 16. AMD accelerators: memory map

## 16. AMD accelerators: memory map

| Accelerator | Memory | Bandwidth | Notes |
|---|---|---|---|
| MI300X (CDNA 3) | 192GB HBM3 | 5.3 TB/s | 750W; shipping |
| MI325X (CDNA 3) | 256GB HBM3e | 6.0 TB/s | 750W+; shipping |
| MI350X (CDNA 4) | 288GB HBM3e | 8.0 TB/s | 1000W air; FP4/FP6/FP8; shipping |
| MI355X (CDNA 4) | 288GB HBM3e | 8.0 TB/s | 1400W liquid; shipping |
| MI400 / MI455X (2026) | 432GB **HBM4** | 19.6 TB/s | 40 PF FP4; 300 GB/s scale-out/GPU |
| Helios rack (H2 2026) | 72× MI455X | — | + Venice EPYC + Vulcano NIC |
[secondary]

- MI355X DLC "Orv3" rack: 2OU, 128 GPUs, 36TB HBM3e, 2.6 exaflops FP4 (96-GPU EIA version: 27TB)
  [secondary].
- AMD publicly names **Samsung as primary HBM4 partner for MI455X** (diversification strategy); Micron is
  the dual-source candidate, extending the Samsung+Micron dual-source used on MI350 [secondary].
- AMD claims MI400X is 10× MI300X-class performance and positions 432GB HBM4 / 19.6 TB/s against
  NVIDIA's Rubin (288GB / 13 TB/s) on memory-intensive workloads; treat vendor performance claims as
  [vendor-reported].

## 17. Intel Gaudi 3: the HBM2e outlier

- Gaudi 3: **128GB HBM2e** in eight stacks at up to **3.7 TB/s** (3.67 TB/s per Tom's Hardware),
  64 Tensor Processor Cores + 8 Matrix Multiplication Engines, 96MB on-die SRAM (19.2 TB/s),
  24× 200GbE RDMA ports, PCIe Gen5 x16; 900W OAM (HL-325L) / 600W PCIe AIC [official][secondary].
- 1.835 PFLOPS FP8/BF16 matrix compute; native FP8 (E4M3 + E5M2), BF16/FP16/TF32/FP32 [official].
- Intel positions Gaudi 3 on lower price/TCO vs H100/H200 rather than peak performance; IBM Cloud offers
  Gaudi 3 virtual-server profiles [secondary].
- Intel's next-gen "Falcon Shores" / Gaudi roadmap beyond Gaudi 3: no confirmed 2026 memory configuration
  found — flagged as a gap [unverified].

## 18. Other accelerators and custom silicon

- Google TPUs: custom ASICs stayed on HBM3e through current shipping generations; Google's Ironwood TPU
  per-chip HBM bandwidth is reported below a single Rubin GPU's HBM4 bandwidth; HBM4 planned for future
  iterations [secondary].
- AWS Trainium (Annapurna Labs): first partner for NVIDIA's NVHBM base-die technology on Trainium4
  (with NVLink Fusion) [secondary].
- Meta MTIA and Microsoft Maia-class custom silicon: HBM-equipped; detailed 2026 memory configurations
  not captured — gap [unverified].
- Cerebras WSE-3 (wafer-scale): uses on-wafer SRAM (44GB) rather than HBM — architecturally distinct;
  memory capacity is the binding constraint vs HBM-based GPUs [secondary].

## 19. GDDR6 / GDDR7

- JEDEC published the GDDR7 standard: PAM3 signaling (1.5 bits/cycle vs NRZ's 1), doubling GDDR6's
  bandwidth; up to 32 Gb/s per pin at launch (48 Gb/s spec max), 192 GB/s per device; four 8-bit channels
  (vs GDDR6's two 16-bit), 32n prefetch; 1.2V; 266-ball FBGA [secondary].
- First-gen GDDR7: 16Gb dies (2GB); 24Gb dies (3GB) followed; 32Gb (4GB) and 48Gb (6GB) densities
  reportedly planned within two years; 3GB modules remain very expensive in 2026, blamed for RTX 5050
  9GB / RTX 50 SUPER refresh delays [secondary].
- Bandwidth math at 32 Gb/s: 128-bit → 512 GB/s; 192-bit → 768 GB/s; 256-bit → 1024 GB/s; 384-bit →
  1536 GB/s; 512-bit → 2048 GB/s [secondary].
- Adoption: NVIDIA GeForce RTX 50 "Blackwell" series first to ship GDDR7; RTX PRO 6000 uses 32× 3GB
  modules (96GB) [secondary].
- AI angle: NVIDIA **Rubin CPX** uses **128GB of GDDR7** — deliberately trading HBM's bandwidth/cost for
  GDDR7's cost efficiency on the *context/prefill* phase of disaggregated inference (millions of input
  tokens); GDDR7 roadmap to 48 GT/s → 192 GB/s per device [secondary].
- Micron roadmap: GDDR7 speed bump to 36 Gb/s expected early 2026 (possibly "GDDR7X" branding) with
  24Gb+ dies [vendor-reported].

