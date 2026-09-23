---
id: etape9-phasec-memory/00-memory/6-ddr6-status-architecture-timeline
title: "6. DDR6: status, architecture, timeline"
domain: step-9-phase-c-server-accelerator-memory-hardware
role: deep-dive
task: hardware
actors: ["Nvidia", "Samsung"]
dates: ["2025-07", "2025-12", "2026-06", "2026-07", "2026-08", "2026-09-22"]
keywords: ["consumer", "cost", "dram", "lpddr", "lpddr5x", "memory", "nvidia", "parameters"]
source: docs/RAG/etape9_phaseC_memory.md
source_anchor: ""
source_lines: [111, 166]
section: "Step 9 — Phase C: Server & Accelerator Memory Hardware"
sha256: ea9a355866b0d103c6126c5b8eb579f21cc8b00741078053e5aa236dbda76fa8
---

# 6. DDR6: status, architecture, timeline

## 6. DDR6: status, architecture, timeline

- JEDEC completed a DDR6 draft specification in late 2024 defining electrical characteristics, signaling
  protocols and performance targets; final DDR6 1.0 ratification (originally targeted 2Q 2025) has slipped
  into 2026 as JEDEC refines timing and signaling parameters [secondary].
- As of May–June 2026 reporting, the DDR6 specification is still not finalized; Samsung, SK hynix and
  Micron are in early substrate and module development with substrate suppliers — commercial product
  specifications are not locked [secondary].
- Target performance: base data rate 8800 MT/s scaling to 17,600 MT/s; specialty overclocked modules may
  target 21,000 MT/s+ [secondary].
- Architecture: DDR6 replaces DDR5's dual 32-bit sub-channel design with a **quad 24-bit sub-channel**
  configuration for more parallelism and bandwidth efficiency, at the cost of stricter signal-integrity
  requirements [secondary].
- Realistic commercial window: early server/enterprise deployments and engineering samples in 2027–2028,
  mainstream consumer adoption 2029–2030; SK hynix's internal roadmap (SK AI Summit 2025) shows DDR6
  introduction around 2029–2030 [secondary].
- One outlier article claimed DDR6 would "debut in high-end devices in 2026 and mass adoption by 2027" —
  this conflicts with the broader 2026 reporting and JEDEC status; treat as [unverified].
- Industry reporting converges on DDR6 moving to the **CAMM2 connector** instead of traditional DIMM
  slots for signal-integrity reasons at 8800+ MT/s, with server platforms expected to lead [secondary].

## 7. LPDDR5X and LPDDR6

- LPDDR5X in modular form: Micron's LPCAMM2 brings LPDDR5X to desktops, laptops and data-center-adjacent
  use with up to 9600 MT/s transfer rates, ~50% higher peak bandwidth over a 128-bit interface, up to
  58% lower active power, up to 80% lower standby power, and 64% smaller footprint vs DDR5 SODIMMs
  [vendor-reported].
- LPDDR6 is the most mature DDR6-era standard: JEDEC published **JESD209-6** in July 2025, defining base
  data rates of 10,667 MT/s up to 14,400 MT/s for mobile/embedded/AI use [secondary].
- LPDDR6 is expected in products around 2026, targeting mobile devices, AI workloads and power-sensitive
  environments first; desktops/workstations follow later [secondary].
- Seismic supply-chain note (Counterpoint via EE Times, August 2026): **Nvidia's pivot to LPDDR** makes it
  a customer "on the scale of a major smartphone maker" — a demand shock the supply chain cannot easily
  absorb, compounding the conventional DRAM shortage [secondary].

## 8. CAMM2 and LPCAMM2 form factors

- CAMM2 (Compression Attached Memory Module, JEDEC standard): LGA-style compression connector instead of
  edge-connector DIMM slots; screws secure the module to the board; shorter traces improve signal integrity
  at high data rates; thinner profile [secondary].
- Two flavors: **DDR5 CAMM2** (DDR5 chips) and **LPCAMM2** (LPDDR5X chips) [secondary].
- Micron LPCAMM2: up to 9600 MT/s claimed, 128-bit interface, single-module replaces dual-SODIMM
  subsystems; Micron's product brief forecasts LPCAMM2 staying two speed grades ahead of SODIMMs through
  2026 (LPCAMM2 8500 vs SODIMM 6400) [vendor-reported].
- Adoption 2026: Lenovo's new-for-2026 T-series ThinkPad line ships 8533 MT/s LPDDR5X LPCAMM2 modules from
  multiple suppliers; iFixit notes the standard enables repairable/upgradable LPDDR5X laptops
  [secondary]. (Note: Micron exited the consumer/Crucial market in December 2025, but LPCAMM2 expansion
  continued via other suppliers [secondary].)
- Industrial/edge: Innodisk introduced DDR5 and LPDDR5X CAMM2 modules for rugged/industrial applications
  (sampling Q4 2025), citing vibration resistance and serviceability; Apacer launched LPDDR5 CAMM2 for
  edge AI in July 2026 [vendor-reported].
- TeamGroup T-Create Expert AI CAMM2 (LPDDR5X): up to 7500 MT/s CL28 (60 GB/s), 32/64GB modules, with
  6400 MT/s CL24 variants [vendor-reported].
- Server angle: no volume server platform has adopted CAMM2 as of 2026-09-22; it is positioned as the
  DDR6-era server connector, while CXL E3.S modules cover the current memory-expansion need [secondary].

