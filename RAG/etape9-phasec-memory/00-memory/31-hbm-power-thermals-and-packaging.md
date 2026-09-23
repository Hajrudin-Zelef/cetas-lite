---
id: etape9-phasec-memory/00-memory/31-hbm-power-thermals-and-packaging
title: "31. HBM power, thermals and packaging"
domain: step-9-phase-c-server-accelerator-memory-hardware
role: deep-dive
task: hardware
actors: ["AWS", "China", "Nvidia", "Samsung", "TSMC"]
dates: ["2026-07-04"]
keywords: ["hbm", "packaging", "accelerator", "dram", "gpu", "hbm4", "lpddr5x", "memory", "nvidia", "nvlink"]
source: docs/RAG/etape9_phaseC_memory.md
source_anchor: ""
source_lines: [544, 590]
section: "Step 9 — Phase C: Server & Accelerator Memory Hardware"
sha256: 1d337cfe6495484e07c152bb1f40fce11eeff75b08dbb96e7f875ffa352236ba
---

# 31. HBM power, thermals and packaging

## 31. HBM power, thermals and packaging

- HBM4 projected at ~75W per stack (KAIST roadmap); HBM8 projected at 180W per stack by 2038 — thermals
  are a first-order limiter for stack height and per-GPU stack counts [secondary].
- Practical consequence: 12-high HBM3e stacks (B300's 288GB) and 16-high HBM4 stacks push
  co-packaging, advanced cooling (direct liquid cold plates are now standard on B200/B300-class systems)
  and interposer design [secondary].
- Base-die logic growth (Section 14) moves heat sources *into* the stack; custom base dies (NVHBM-class)
  trade XPU die area for stack complexity [secondary].

## 32. The base-die revolution (2025–2026)

- HBM4's 2048-bit interface and 32 channels grew the control/interface logic in the base die to the point
  where DRAM-process logic is no longer optimal — hence all three vendors outsourcing base dies to TSMC
  logic processes [secondary].
- NVIDIA NVHBM (Aug 2026): memory controller moves into the base die; +30% bandwidth, −15% HBM power,
  +25% XPU die area freed vs standard HBM4E; first partner Amazon Annapurna (Trainium4 + NVLink Fusion)
  [secondary].
- HBM4e adds customizability: base dies tailored per application (AI, HPC, networking) rather than one
  standard product [secondary].
- Strategic read: TSMC now extends from logic contract manufacturing into memory-based components — HBM
  is described as a new growth curve for TSMC [secondary].

## 33. Server memory form factors, consolidated

| Form factor | Use | Status 2026 |
|---|---|---|
| DDR5 RDIMM (288-pin) | Standard server DIMM | Volume; 6400 MT/s mainstream, 8000 MT/s emerging |
| DDR5 LRDIMM / 3DS | Max capacity/server | Volume; up to 256GB+/module |
| MRDIMM | Bandwidth-optimized server | 8800 MT/s Q1 2027; 12,800 JEDEC Gen2 |
| E3.S CXL Type 3 | Memory expansion/pooling | Volume (Samsung/SK hynix/Micron), 96–256GB |
| LPCAMM2 | Laptop/edge modular LPDDR5X | Shipping (Lenovo 2026, others) |
| CAMM2 (DDR5) | Desktop/laptop/industrial | Sampling/early adoption |
| HBM (2e/3/3e/4) | Accelerator in-package | HBM3e volume; HBM4 ramping H2 2026 |

## 33B. DRAM process nodes, 2026

- Samsung: 6th-generation **1c** DRAM technology applied to HBM4 development ahead of competitors;
  Counterpoint notes Samsung may reallocate its expanding 1c process to meet demand [secondary].
- Micron: **1β (1-beta)** process underpins the 32Gb monolithic die in 128GB DDR5 RDIMMs (45% better bit
  density vs prior node); Hiroshima plant leads 1β production and is introducing EUV for **1γ** DRAM, with
  a new ~28,000 m² cleanroom broken ground July 4, 2026 [vendor-reported][secondary].
- SK hynix: pushing production and raising sales targets; 1c-class processes in the mix for HBM4-era
  output [secondary].
- Chinese vendor CXMT "may surprise on the upside" per Counterpoint — a watch item for supply diversification
  [secondary].

