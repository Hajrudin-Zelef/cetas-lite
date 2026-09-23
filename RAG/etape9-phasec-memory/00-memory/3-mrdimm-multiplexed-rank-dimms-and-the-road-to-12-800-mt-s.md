---
id: etape9-phasec-memory/00-memory/3-mrdimm-multiplexed-rank-dimms-and-the-road-to-12-800-mt-s
title: "3. MRDIMM: multiplexed-rank DIMMs and the road to 12,800 MT/s"
domain: step-9-phase-c-server-accelerator-memory-hardware
role: deep-dive
task: hardware
actors: ["AMD", "Intel"]
dates: ["2025-10", "2025-11", "2026-09"]
keywords: ["accelerator", "agentic", "amd", "clearwater forest", "cost", "energy", "gpu", "gpus", "helios", "intel", "latency", "memory"]
source: docs/RAG/etape9_phaseC_memory.md
source_anchor: ""
source_lines: [60, 110]
section: "Step 9 — Phase C: Server & Accelerator Memory Hardware"
sha256: ff7be910b7cf81213a44b853eb75c6a1d71cb60c6aaed4735b523ffd070cc0cb
---

# 3. MRDIMM: multiplexed-rank DIMMs and the road to 12,800 MT/s

## 3. MRDIMM: multiplexed-rank DIMMs and the road to 12,800 MT/s

- Intel Xeon 6 added 8000 MT/s DDR5 RDIMM support (August–September 2026) on select Xeon 6700P "Granite
  Rapids" and Xeon 6+ "Clearwater Forest" SKUs, in 1DPC configurations; measured 3–6% workload gains vs
  6400 MT/s, up to ~20% more total memory bandwidth [secondary].
- The real payoff is **MRDIMM** (Multiplexed Rank DIMM): first-generation MRDIMMs target 8800 MT/s on
  select Xeon 6900P systems in Q1 2027; JEDEC published the MRDIMM Gen2 standard pushing server memory to
  **12,800 MT/s** for AI and data-center workloads [secondary].
- Longer-term MRDIMM targets: up to 17,600 MT/s expected around 2027, and up to 17,600 MT/s on future
  platforms in the 2029–2030 window, per Intel-side reporting [secondary].
- Memory interface chipsets enabling the ramp: Rambus unveiled Gen5 DDR5 RDIMM chipsets (Gen5 RCD +
  PMIC5030 + SPD hub + temperature sensor) for DDR5 RDIMM 8000, and MRCD/MDB chipsets for MRDIMM 12800 —
  an industry-first complete chipset for both [vendor-reported].
- Renesas pushed a Gen6 DDR5 registered clock driver (RCD) targeting 9600 MT/s for the next wave of
  RDIMMs (November 2025 announcement) [vendor-reported]; Montage Technology scaled Gen4 DDR5 RCD04 support
  to 7200 MT/s for mass-production RDIMMs (October 2025) [vendor-reported].
- Micron was the first manufacturer to ship 128GB DDR5 RDIMMs (32Gb monolithic die, 1β process) rated up
  to 8000 MT/s, validated from 5600 MT/s on all leading server platforms; Micron claims 45% better bit
  density, up to 24% better energy efficiency and up to 16% lower latency vs competitive 3DS TSV products,
  plus up to 28% faster AI training performance [vendor-reported].

## 4. Server DIMM taxonomy: UDIMM vs RDIMM vs LRDIMM vs 3DS

- **UDIMM (unbuffered)**: no register; used in workstations/entry servers; capacity limited, 1–2 DPC
  typically [secondary].
- **RDIMM (registered)**: register clock driver (RCD) buffers command/address; the standard server DIMM;
  **all DDR5 RDIMMs are side-band ECC modules** (80-bit bus: 2×(32 data + 8 ECC)) [secondary].
- **LRDIMM (load-reduced)**: memory buffer (data buffer) reduces electrical load per rank, enabling higher
  ranks per channel (used for max-capacity configurations, e.g., 8–12 ranks/channel equivalents)
  [secondary].
- **3DS (3D-stacked)**: TSV-stacked dies for maximum capacity per slot (e.g., 256GB DDR5 RDIMMs);
  higher latency and cost than monolithic-die modules [secondary].
- Practical rule of thumb 2026: balance capacity against speed — populating all channels at 1DPC with
  RDIMMs maximizes bandwidth; LRDIMM/3DS buys capacity at the cost of latency and (for MRDIMM-era
  chipsets) chipset compatibility [independent].

## 5. Platform memory architectures, 2026

- Intel Xeon 6 (Granite Rapids / Clearwater Forest): DDR5-6400 baseline, 8000 MT/s RDIMM support on
  select SKUs (1DPC), MRDIMM 8800 MT/s arriving with Xeon 6900P systems in Q1 2027 [secondary].
- Intel's messaging positions host memory bandwidth as a first-order AI design variable: agentic AI
  workloads push the host CPU to orchestrate models, data prep, storage access and inter-accelerator
  traffic, making memory bandwidth the difference between busy and stalled accelerators [vendor-reported].
- AMD EPYC 5th Gen "Turin" (Zen 5): up to 192 cores per socket, shipping; EPYC 6th Gen "Venice" expected
  H2 2026 with up to 256 cores / 512 threads, claimed 2.0x CPU-to-GPU bandwidth and 1.7x gen-over-gen
  performance uplift, plus a reported 1.6 TB/s of memory bandwidth [secondary][unverified for the 1.6 TB/s figure].
- Venice is designed into AMD's Helios rack (72× MI400-class GPUs + Venice CPUs) for H2 2026 [secondary].
- General server trend: core counts have leapt from 12 to 96 and now 128+ cores per socket while memory
  bandwidth scaled more slowly, making bandwidth-per-core the binding constraint for data-intensive
  workloads (AI, in-memory DBs, analytics) [secondary].

