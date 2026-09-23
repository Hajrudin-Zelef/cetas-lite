---
id: etape9-phasee-storage-market/00-storage-market/e3b-nand-technology-race-2026-layers-density-and-the-china-f
title: "E3b — NAND technology race 2026: layers, density, and the China factor"
domain: step-9-phase-e-storage-memory-market-2026
role: deep-dive
task: reference
actors: ["AMD", "AWS", "China", "Google", "Intel", "Meta", "Nvidia", "Samsung", "United States"]
dates: ["2023-05"]
keywords: ["nand", "amd", "asic", "capex", "consumer", "cost", "datacenter", "dram", "gpu", "gpus", "hbm", "hbm3"]
source: docs/RAG/etape9_phaseE_storage_market.md
source_anchor: ""
source_lines: [308, 353]
section: "Step 9 Phase E — Storage & Memory Market 2026"
sha256: 3ecccd19726a5944a09b3d3827cac7471f0682d6f3277483e8dd7f62aa1ecd87
---

# E3b — NAND technology race 2026: layers, density, and the China factor

## E3b — NAND technology race 2026: layers, density, and the China factor

- **2026 layer-count leaderboard (mass production):** SK hynix **321L** (industry first past 300, "4D NAND") > Samsung **V9 286L** > Micron **G9 276L (B58R)** > Kioxia/SanDisk **BiCS8 218L** > YMTC **232L (Xtacking 3.0)** mass-producing **267L (Xtacking 4.0)** [secondary]. Source: http://hsj00.github.io/assets/pdf/3D_NAND_V9_V12_Technical_Report.pdf
- **The V9 generation (2025–2026):** layer range 232L (Kioxia BiCS7, YMTC) to 321L (SK hynix); dual-deck stacking standardized; TLC standard with QLC expanding; **3.2 Gbps I/O standardized**; bit density 21–28 Gb/mm² (TLC) [secondary].
- **SK hynix's 321L milestone:** first vendor past 300 layers, but at a cost — etch steps +20% and total process steps +30% vs V8 [secondary]. SK hynix also leads on **HBF (high-bandwidth flash)** — vertically stacking 3D NAND dies like HBM — and formed a **consortium with SanDisk to standardize HBF specifications globally** [secondary]. Source: https://m.koreaherald.com/article/10711815
- **Samsung V9 notes:** partial **Mo (molybdenum) gate** application reducing wordline resistance; **V9 QLC mass production reportedly delayed in H1 2026 on durability issues** [secondary]. V10 (400+ layers, triple-deck, hybrid bonding, 5.6 GT/s) shown at ISSCC 2025 with 1 Tb TLC die at 28 Gb/mm² [independent]. Source: https://www.tomshardware.com/pc-components/ssds/samsung-unveils-10th-gen-v-nand-400-layers-5-6-gt-s-and-hybrid-bonding
- **Micron G9:** 276L TLC, up to 3600 MT/s I/O, 21.0 Gb/mm² [secondary].
- **Kioxia/SanDisk BiCS8:** 218L, 2 Tb QLC die option, up to 3600 MT/s; BiCS9 (~332L) in development [secondary].
- **YMTC's 2026 surge (the market's biggest structural story after HBM):**
  - **14% of global NAND bit shipments in Q2 2026** (Counterpoint) — third place behind Samsung (25%) and SK hynix (22%), edging out Kioxia; bit shipments +22% YoY, +5% QoQ [secondary]. Source: https://www.tweaktown.com/news/113106/chinas-ymtc-has-climbed-to-third-place-in-global-nand-flash-shipments/index.html
  - Still trails in *revenue* due to consumer-heavy mix, but explicitly shifting toward enterprise SSDs and considering a mainland IPO (valuation expected >$40B) [secondary].
  - Mass-producing **267L Xtacking 4.0**, targeting 300+ layers next; ~200K WSPM across two Wuhan fabs, third fab adding 100K WSPM later in 2026, two more 50K fabs in pipeline → **~400K WSPM total, in Samsung's territory (390–450K)** [secondary].
  - Announced **three new production lines operational 2026–2027** (flagship Phase 3 Wuhan, H2 2026), with **>50% domestic equipment sourcing**; **10–20% price advantage** over Japanese/Korean rivals [secondary]. Source: https://www.pulse.bot/hardware/news/ymtc-expands-nand-and-dram-ambitions-with-new-fabs-despite-us-sanctions-pressure-edab59ce-3e39-4b05-9470-e6762dd62d77/
  - Phase 3 capacity split roughly evenly **NAND + DRAM** (testing LPDDR samples, pursuing TSV-based HBM stacks for AI) — YMTC is becoming a memory generalist, not just a NAND house [secondary].
- **Korean defensive moves:** Samsung halted 128L production at its Xian (China) plant (~40% of its NAND output) and is ramping 236L → 286L in 2026, roadmap beyond 400L; SK hynix raised Dalian (China) NAND capex >50% to ₩440.6B [secondary].
- **What the layer race means for buyers:** layer count is a cost-per-bit weapon, not a performance feature — higher layers → smaller dies → lower $/GB *at the fab*. In a shortage, vendors monetize it as margin rather than passing it through; the 1Tb TLC die spot doubling ($4.80 → $10.70) happened *during* the 300L ramp [secondary].
- **PLC watch:** 5-bit-per-cell NAND remains a lab/roadmap item (<500 P/E cycles estimated); no 2026 product signal found — do not plan around PLC [unverified].

## E5b — HBM generation specification table (JEDEC + vendor data)

| Generation | JEDEC / intro | Interface | Pin speed | Bandwidth/stack | Max capacity/stack | Typical deployment |
|---|---|---|---|---|---|---|
| HBM2E | 2018–2020 | 1024-bit | 2.5–3.2 Gbps | ~460 GB/s | 16 GB | NVIDIA A100, AMD MI250X, Intel Agilex 7 M-series FPGAs |
| HBM3 | Jan 2022 | 1024-bit | 6.4 Gbps | 819 GB/s | 24 GB | NVIDIA H100 |
| HBM3E | May 2023 (JESD238) | 1024-bit, 16 ch / 32 pseudo-ch | 9.2–12.4 Gbps | 1.15–1.33 TB/s | 24 GB (8H) / 36 GB (12H) | NVIDIA H200 (6×24 GB = 141 GB, 4.8 TB/s), AMD MI300X (8×24 GB = 192 GB, 5.3 TB/s), Intel Gaudi 3 |
| HBM4 | Apr 2025 (JESD270-4) | 2048-bit, 32 ch / 64 pseudo-ch | 6.4–12.8 Gbps (13 Gbps demo'd) | >2.0 TB/s (up to 3.3 TB/s) | 64 GB (16H, 32Gb layers) | NVIDIA Rubin (2026), AMD MI400-class [unverified for MI400] |
| HBM4E | Previewed Mar 2026 | 2048-bit | 16 Gbps | ~4.0 TB/s | TBD | Roadmap |

- HBM4 is an architectural overhaul: doubled interface width, **not backward-compatible** with HBM3/3E controllers; 1.05V core (vs 1.1V), ~60% efficiency gain over HBM2/2E, new Directed Refresh Management (DRFM) [secondary]. Source: https://blogs.sw.siemens.com/semiconductor-packaging/2026/04/24/hbm3e-hbm4-ic-design-guide/
- **Custom-ASIC caveat:** Google/Amazon/Meta AI ASICs largely stayed on **HBM3E through their current shipping generations** even as HBM4 ramped for NVIDIA/AMD GPUs — HBM4's per-chip bandwidth lead is currently a GPU-vendor advantage, not a universal baseline [secondary]. Source: http://gpusmith.com/articles/en/pdfs/hbm3e-vs-hbm4-ai-gpus.pdf
- **Why the 12H→16H stack transition worsens the shortage:** taller stacks consume more DRAM die area per GB of delivered HBM — the "reverse scaling" mechanism quantified in E5 [secondary].

## E7b — Used-market deep dive: SKUs, patterns, and price anchors

- **The classic homelab enterprise SSDs (recurring eBay/ServerPartDeals listings, 2025–2026):** Samsung PM883/PM893 (SATA), PM9A3/PM963 (NVMe U.2), Intel/Micron 5300/5400 Pro/Max (SATA), Intel DC P4510/P4610 (NVMe U.2), Kioxia/HGST SAS lines, Micron 9300/9400 Pro/Max (NVMe U.2). These are the parts the $78–127/TB band refers to [secondary].
- **Why read-intensive ("1 DWPD / RI") pulls dominate listings:** most datacenter flash does read-heavy work, so retired RI drives typically show single-digit percentage-used — the exact drives a homelab wants for VM datastores and media [secondary].
- **Write-intensive ("3–10 DWPD / WI") pulls:** rarer, pricier per TB, and the right buy only for SLOG/ZIL, database redo, or heavy ingest — and even then, new is the 2026 guidance for write-heavy tiers (see E9) [secondary].
- **SATA vs SAS vs NVMe on the used market:** SATA enterprise SSDs are the cheapest per TB and slot into any backplane; SAS needs a SAS HBA/controller; **NVMe U.2 needs U.2 bays or M.2/U.2 adapters plus PCIe bifurcation or tri-mode controllers** — factor the adapter/cable cost (~$20–60) into the $/TB math [secondary].
- **DDR4 ECC RDIMM (used, 2026):** the homelab RAM staple — 16 GB sticks $25–40, 32 GB sticks commonly $50–90 on eBay-class listings; 64 GB LRDIMMs higher but still far below new DDR5 RDIMM pricing [secondary]. Compatibility rule: match the server generation's rank/voltage support (e.g., Dell 13G/14G DDR4 2400/2666) — check the vendor QVL or community reports before bulk-buying [secondary].
- **The mini-PC vs used-server tradeoff (2026):** an MS-01-class mini PC idles at 13–35 W in silence, saving $120–200/year vs a rack server at average US power rates; the used server wins when the job needs hundreds of GB of cheap ECC RAM, 8–24 hot-swap bays, full-height GPU slots, redundant PSUs, and a real BMC — purchase price favors the server, running cost favors the mini PC [secondary]. Source: https://pcserverandparts.com/news/best-used-server-for-home-lab-2026/
- **Where the deals cluster:** corporate refresh cycles (3–5 years) dump matched lots — same SKU, same firmware, similar wear — which is ideal for RAID/ZFS vdevs; single-drive odd lots are cheaper per unit but complicate array building [secondary].
- **ServerPartDeals pattern:** enterprise-focused retailer with rotating stock of pulled drives, publishes per-drive SMART on request; typical premium over eBay auctions but with testing and return policy — the "verified middle" between eBay luck and new retail [secondary].
- **r/homelabsales pattern:** peer-to-peer, best for full servers and RAM lots; escrow via PayPal G&S is the community norm; prices typically 10–20% under eBay after fees [secondary].
- **What "recertified" means at different sellers:** manufacturer recertified (rare for enterprise SSDs) > refurbisher-tested with SMART report and 3–5-year warranty > "pulled working" (as-is, test on arrival) > "for parts." Price should descend in that order; if it doesn't, walk away [secondary].
- **Power-cost reality check for 24/7 spinning rust:** a 4-bay NAS with 8 TB HDDs at ~$80–120 used per drive plus ~25–40 W idle is still the cheapest $/TB/year for media in 2026 — the SSD premium only pays off where IOPS or latency matter [secondary].

