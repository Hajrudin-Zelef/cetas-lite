---
id: etape9-phasec-memory/00-memory/20-hbm-supply-chain-and-market-2026
title: "20. HBM supply chain and market, 2026"
domain: step-9-phase-c-server-accelerator-memory-hardware
role: deep-dive
task: hardware
actors: ["China", "Microsoft", "Nvidia", "Samsung"]
dates: ["2025-08", "2025-11", "2026-04", "2026-07", "2026-07-04", "2026-08", "2026-09-17", "2026-09-18"]
keywords: ["hbm", "capex", "compute", "consumer", "cost", "dram", "foundry", "gpu", "hbm4", "hyperscaler", "latency", "memory"]
source: docs/RAG/etape9_phaseC_memory.md
source_anchor: ""
source_lines: [356, 425]
section: "Step 9 — Phase C: Server & Accelerator Memory Hardware"
sha256: 23ce223984e6107efca7925fbf31331ef5cc7392a39e00b086a96921089ba4f3
---

# 20. HBM supply chain and market, 2026

## 20. HBM supply chain and market, 2026

- Three suppliers only: SK hynix (~70% HBM market share per Korean press), Samsung, Micron [secondary].
- 2026 capacity estimates (TrendForce composite): SK hynix ~150–170k wafers/month (12-Hi equivalent),
  Samsung ~110–130k, Micron ~50–60k — combined ~31–36M wafer-equivalents/year, roughly 900M–1B 12-Hi
  stacks, the physical ceiling for ~7.5–8M Rubin-class GPU shipments [secondary].
- Demand vs supply: 2026 top-5 hyperscaler capex ~$680–740B, GPU ~30–35% of capex (~$210–260B), HBM
  ~50–60% of GPU BOM → HBM-side demand ~$100–150B vs industry HBM revenue estimate ~$80–100B: demand
  exceeds supply by ~20–50% — the math behind "sold out through 2028" [secondary].
- 2026 HBM production sold out months in advance to AI data-center customers; because HBM and consumer
  DRAM/GDDR share finite wafer capacity, the shortage pushed conventional DRAM and GDDR prices sharply
  higher too [secondary].
- Samsung–SK hynix competition is NVIDIA's leverage: Samsung offered 12-layer HBM3e for the NVIDIA H20
  (China) at 20–30% below SK hynix pricing, pending qualification; Jensen Huang publicly stated "we need
  Samsung, and we need SK hynix" [secondary].
- Export-price data (Korea International Trade Association): average HBM export unit price $76.13 end of
  July 2026 (+9.5% MoM, first time above $70); Q2 2026 average $60.22 vs $40.94 in Q1 as HBM4 shipments
  began [secondary].
- HBM is ~3–4× more wafer-intensive per gigabyte than conventional DRAM [secondary].
- Capacity expansions: Samsung reviewing conversion of its Pyeongtaek S5 foundry line to memory; SK hynix
  expanding Yongin/Honam and weighing a Japan JV plant (possibly with Kioxia); Micron broke ground on a
  new ~28,000 m² cleanroom at Hiroshima (July 4, 2026) for 1γ DRAM and next-gen AI DRAM/HBM [secondary].

## 21. DRAM pricing signals, 2026

- Retail (MemoryPriceChart, September 18, 2026): chain-linked DDR5 index +397% since August 2025; basket
  median **$18.44/GB** (mean $18.83/GB) across six 32/64GB desktop DDR5 kits; SSD index +158% over the same
  period — a 239-point divergence between DRAM and NAND pricing [secondary].
- Secondary-market snapshot (August 2026): DDR4 ~$8.11/GB, DDR5 ~$18.10/GB [secondary].
- GameGPU (September 17, 2026): 64GB DDR5-5600 kits rose from $191 (Aug 2025) to $1,118 (Aug 2026), +485%;
  32GB DDR5-6000 kits from <$90 to $370–530 [secondary].
- Framework (vendor pass-through pricing): $10/GB for DDR5 laptop modules (weighted-average-cost basis),
  raised repeatedly through 2025–2026 on supplier volatility [vendor-reported].
- Contract/spot: DDR5 16Gb (2GB×8) average $47.07 in July 2026 (~$2.94/Gb); DDR3 4Gb $12.75 (~$3.19/Gb) —
  legacy DDR3 temporarily priced above DDR5 per-gigabit [secondary].
- Server DRAM commands a premium over desktop UDIMM pricing above (server RDIMM $/GB runs higher due to
  RCD/PMIC/ECC and lower volumes); exact 2026 server-RDIMM $/GB contract figures were not captured —
  gap [unverified].
- Outlook: analysts see the shortage lasting to 2028, possibly 2030; "no point waiting for prices to
  drop" is the 2026 consensus for buyers [secondary].

## 22. CXL: memory expansion and pooling

- CXL (Compute Express Link) runs on the PCIe physical layer and adds cache-coherent device-attached
  memory (CXL.mem) plus cache semantics (CXL.cache) on top of CXL.io [secondary].
- Version timeline: CXL 1.0/1.1 (2019, PCIe 5.0) → CXL 2.0 (2020: pooling, switched topology) → CXL 3.0
  (2022: PCIe 6.0, fabric-attached memory) → CXL 3.1 (Nov 2023: TEE-IO security, GFAM — Global Fabric
  Attached Memory, breaking tree-only topologies) → CXL 4.0 expected 2026 (PCIe 7.0 backbone; IP-only
  from Synopsys/Cadence so far) [secondary].
- 2026 commercial reality: volume production is still centered on **CXL 2.0-era Type 3 memory expansion**;
  CXL 3.x specs are published but CPUs, switches, devices, OSes and management software must mature
  together before large-scale 3.x deployment [secondary].
- Type 3 market: estimated **$1.8–2.5B in 2026**, driven by hyperscalers and AI infra providers; CXL
  pooling has demonstrated up to 50% higher effective memory utilization [secondary].
- Vendor products (all E3.S, PCIe Gen5 x8, 96–256GB): Samsung CMM-D (expansion), CMM-B (box pooling
  appliance), CMM-H (DRAM+NAND hybrid), CXL-PNM (processing-near-memory); SK hynix CMM-DDR5 and CMM-Ax
  (memory+compute); Micron CZ120/CZ122 [secondary].
- Add-in cards: SMART Modular CXL AICs (Type 3, PCIe Gen5 FHHL) take standard DDR5 RDIMMs — 4-DIMM (up to
  2TB with 512GB RDIMMs) or 8-DIMM (up to 4TB), 64 GB/s total bandwidth [vendor-reported].
- Latency budget reference (community-compiled, April 2026): local DDR5 ~80–90ns; CXL 2.0 attached DRAM
  (no switch) ~170–250ns at ~50 GB/s (PCIe 5.0 x16); CXL 3.x switched pool ~250–500ns at ~100 GB/s
  (PCIe 6.0 x16) — versus NUMA remote-socket DRAM at ~200–250ns [secondary].
- Cloud milestone: Microsoft Azure ran the first cloud CXL deployment (CXL 2.0) in November 2025
  [secondary].
- Near-memory compute: XCENA + Samsung's MX1 (Hot Chips 2026) — a Type 3 device combining up to 2TB DDR5,
  SSD-backed byte-addressable capacity and **3,072 RISC-V cores** for RAG search, KV-cache access and
  memory-bound analytics [secondary].
- CXL 3.x switch silicon: Astera Labs Atlas 3, Marvell_structured interconnect and Montage M88MX6852
  sampling 2025–2026 [secondary].

