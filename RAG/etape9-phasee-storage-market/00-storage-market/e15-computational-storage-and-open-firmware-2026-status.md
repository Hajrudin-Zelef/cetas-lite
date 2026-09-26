---
id: etape9-phasee-storage-market/00-storage-market/e15-computational-storage-and-open-firmware-2026-status
title: "E15 — Computational storage and open firmware: 2026 status"
domain: step-9-phase-e-storage-memory-market-2026
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Google", "Meta", "Samsung"]
dates: ["2026-07"]
keywords: ["amd", "compute", "dram", "energy", "gpu", "hbm", "hbm4", "inference", "kv cache", "memory", "nand", "research"]
source: docs/RAG/etape9_phaseE_storage_market.md
source_anchor: ""
source_lines: [185, 203]
section: "Step 9 Phase E — Storage & Memory Market 2026"
sha256: 436ed9309aa3d389670fab74af3b3309914b025817d5ebbe269c1a350af63fbf
---

# E15 — Computational storage and open firmware: 2026 status

- **Retail DDR5 (Sept 2026): ~$18.44/GB** (MemoryPriceChart basket median) [secondary].
- **Retail NAND (Sept 2026): ~$0.16/GB** ($159.94/TB mainstream NVMe) — a **~115× gap** between retail DRAM and retail NAND per GB [secondary].
- **Enterprise SSD (Q3 2026): ~$0.75/GB** ($753/TB for 30 TB TLC) vs **enterprise HDD ~$0.04/GB** ($40.5/TB) [secondary].
- **Used enterprise SSD: ~$0.08–0.13/GB** ($78–127/TB) — roughly **6–9× cheaper per GB than new enterprise NVMe** ($0.30–1.17/GB) [secondary].
- **HBM (spot, Aug 2026): ~$58/GB contract-class (36 GB stack at $2,100 spot) up to ~$97/GB (HBM4 16-layer at $3,500)** — **3,000–5,000× retail DDR5 per GB**, priced as strategic allocation [secondary].
- **The GPU memory tax:** GDDR7-driven BOM inflation pushed the RTX 5090 from $1,999 MSRP to $4,329 street (mid-July 2026); VRAM is >80% of high-end GPU BOM [secondary]. For inference buyers, the $/GB ladder is: NAND < used enterprise flash < new enterprise flash < DDR5 < GDDR < HBM — and each rung is 2–100× the previous [secondary].
- **Why this matters for AI infra planning:** training clusters are sized by HBM capacity per GPU and checkpoint *bandwidth*; inference fleets are sized by **weights + KV cache per card**; bulk dataset storage stays on the cheapest tier that meets the ingest rate (HDD/object for cold, QLC/new-NVMe for warm, node-local NVMe for hot) [secondary].

## E15 — Computational storage and open firmware: 2026 status

- **Samsung SmartSSD: effectively mothballed.** The concept (2018) put NAND + HBM + RDIMM next to an AMD Xilinx FPGA inside the SSD for server-less compute; Gen2 launched 2022 [secondary]. By 2025 it had "all but disappeared from Samsung's portfolio" — still buyable on Amazon under the **AMD Xilinx brand at $517.70 for 3.84 TB**, but a Gen3 device whose novelty and complexity made it a hard sell; COVID-19 then generative AI (which demanded *capacity*, not in-drive compute) killed the business case [secondary]. Source: https://www.techradar.com/pro/samsung-and-amd-made-a-revolutionary-ssd-together-then-it-was-left-to-wither-in-the-shadows-and-nobody-knows-exactly-why
- **Assessment:** computational storage devices (CSDs) were "an interesting but niche market, closer to traditional servers" — nice, but without AI-hardware growth potential; Samsung mothballed after Gen2 despite 2022 claims of "great potential" [secondary].
- **What replaced the CSD story:** (1) **CXL-attached memory/NAND** (see E6) as the industry's chosen "compute near data" vehicle; (2) **DPUs/SmartNICs** doing storage offload on the network side (see Step 6 Phase F1); (3) **in-drive AI for the drive itself** (predictive failure, ZNS/FDP placement) rather than general compute [secondary].
- **NGD Systems:** the other notable CSD vendor — no 2026 product/market signal surfaced in this research pass; treat as dormant-or-acquired, flagged as a gap [unverified].
- **Open-source SSD firmware (OpenSSD / Cosmos+ OpenSSD):** the community FPGA-based open firmware platform remains a research/education vehicle (universities, FTL research); no evidence of 2026 commercial adoption or a new platform revision in this pass — gap flagged [unverified].
- **Open-channel SSDs / ZNS as the "open" successor:** the industry's openness energy moved from open firmware to **open interfaces** — NVMe ZNS (zone namespaces) and FDP (flexible data placement) expose placement control to the host without opening the FTL; hyperscalers (Meta, Google) drive ZNS adoption for write-amplification reduction. Details are in the sibling SSD-hardware file [secondary].

## E16 — Enterprise SSD roadmap: PCIe Gen6 and beyond

