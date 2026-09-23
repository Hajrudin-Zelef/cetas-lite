---
id: etape9-phasec-memory/00-memory/overview
title: "Step 9 — Phase C: Server & Accelerator Memory Hardware"
domain: step-9-phase-c-server-accelerator-memory-hardware
role: deep-dive
task: hardware
actors: ["California", "Intel", "SGLang", "Samsung", "TensorRT-LLM", "United States", "vLLM"]
dates: ["2026-08", "2026-09-22"]
keywords: ["accelerator", "memory", "dram", "dram price", "hbm", "inference", "intel", "lawsuit", "lpddr5x", "power delivery", "pricing", "research"]
source: docs/RAG/etape9_phaseC_memory.md
source_anchor: ""
source_lines: [1, 59]
section: "Step 9 — Phase C: Server & Accelerator Memory Hardware"
sha256: b6dd56596d42a3307e242eb9b213bf2ee67eac601c02e7ea9cc275e703a4177e
---

# Step 9 — Phase C: Server & Accelerator Memory Hardware

**Research cut-off: 2026-09-22. English. Sole writer; no other workspace files modified.**

## 0. Scope, method and provenance

This file covers server and accelerator memory *hardware* (Step 9 track C): DDR5 server DIMMs and MRDIMMs,
the DDR6 timeline, LPDDR5X/6, CAMM2/LPCAMM2 form factors, HBM generations and which accelerators use them,
GDDR6/GDDR7, CXL memory expansion and pooling, ECC/RAS, the end of Intel Optane persistent memory, AI
memory sizing math (inference and training), and 2026 memory pricing and supply signals. Cross-cut: storage-class
memory software (Ceph, ZFS) was covered in Step 7E; networking fabrics in Step 6; inference *servers* (vLLM,
SGLang, TensorRT-LLM) in Step 4 — this file stays on the hardware/memory-subsystem angle.

Method: read-only web research (search + page text), current through 2026-09-22. Provenance tags on factual
lines: `[official]` vendor/standards-body primary source, `[vendor-reported]` vendor marketing/press claims,
`[independent]` independent measurement or community evidence, `[secondary]` press/analyst/research aggregators,
`[unverified]` could not be corroborated or comes from a weak source. No SKUs, prices, versions or URLs were
invented; gaps, conflicts and non-comparable figures are flagged in-line and in the register (Section 36).

---

## 1. Server DRAM landscape, 2026

- DDR5 is now the default server memory technology: DDR5 accounted for over 80% of server DRAM shipments,
  with DDR4 falling below 20% and possibly heading toward discontinuation [secondary].
- The defining story of 2026 is a severe DRAM shortage driven by AI: hyperscalers buying record volumes of
  memory for AI server farms, plus suppliers reallocating wafer capacity from DDR to HBM (HBM die area is
  roughly twice a standard DDR chip; HBM production accounts for about 25% of global silicon wafer
  production) [secondary].
- Industry experts estimate DRAM price surges continuing until shortages last, expected until 2028, with
  further hikes expected in Q4 2026; Samsung has warned 2027 could be worse than 2026 [secondary].
- Taiwan's Digitimes reported (August 2026) that DRAM and HBM production allocation for 2027 is already
  sold out across the three main manufacturers (Samsung, SK hynix, Micron) [secondary].
- Counterpoint Research: memory prices rose ~50% in 2025, were forecast to rise ~30% in Q4 2025 and possibly
  another ~20% early 2026; suppliers are racing to meet demand with DRAM production increases exceeding 20%
  in 2026 [secondary].
- The three dominant DRAM suppliers (Samsung, SK hynix, Micron) face a US class-action lawsuit filed in the
  Northern District of California alleging coordinated production since 2022 (31→8 weeks of inventory,
  alleged 700% price increase over four years) [secondary].

## 2. DDR5 deep-dive: speeds, architecture, modules

- JEDEC DDR5 data rates span 4800 to 8400 MT/s (4.8 to 8.4 GT/s), delivering 50%+ more bandwidth than DDR4's
  1.6–3.2 GT/s ceiling; DDR5-4800 delivers 38.4 GB/s per module vs 25.6 GB/s for DDR4-3200 [secondary].
- Architecture changes vs DDR4: dual independent 32-bit sub-channels per DIMM, 8 bank groups / 32 banks
  (vs 4/16), burst length 16 (vs 8), on-module PMIC for power delivery, SPD hub and temperature sensor,
  and **mandatory on-die ECC** on every DDR5 chip [secondary].
- Die densities in 2026: 16Gb and 24Gb in volume, 32Gb monolithic dies shipping (Micron 1β-based 32Gb die
  underpins 128GB RDIMMs), with projections up to 64Gb per die in the DDR5 generation [secondary][vendor-reported].
- Server RDIMM capacities in 2026: mainstream 16/32/64/96/128GB RDIMMs; 3DS (through-silicon-via stacked)
  parts reach 256GB per module (Micron demonstrated the world's first 512GB DDR5 module, enabling a claimed
  12TB of RAM in one server) [vendor-reported][secondary].
- Pin compatibility: DDR5 long DIMM keeps 288 pins (keyed differently from DDR4); DDR5 SO-DIMM has 262 pins
  [secondary].
- Voltage: DDR5 operates at 1.1V (vs 1.2V DDR4), with the PMIC moved onto the module itself for cleaner
  power delivery at high speeds [secondary].
- Signal integrity at 6400+ MT/s forces 1 DIMM-per-channel (1DPC) operation on current platforms to hold
  rated speeds; 2DPC operation typically steps down [secondary].

