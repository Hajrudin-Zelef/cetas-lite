---
id: etape9-phasee-storage-market/00-storage-market/e3-vendor-landscape-2026-who-makes-the-memory
title: "E3 — Vendor landscape 2026: who makes the memory"
domain: step-9-phase-e-storage-memory-market-2026
role: deep-dive
task: reference
actors: ["China", "Samsung", "TSMC"]
dates: ["2026-05", "2026-08", "2026-08-05", "2026-09-18"]
keywords: ["memory", "capex", "consumer", "cost", "dram", "foundry", "hbm", "hbm4", "nand", "pricing", "revenue", "valuation"]
source: docs/RAG/etape9_phaseE_storage_market.md
source_anchor: ""
source_lines: [32, 59]
section: "Step 9 Phase E — Storage & Memory Market 2026"
sha256: 6d9dd28f6f68801127daee0e06a94907af953fe9c7e517617a9e65eb1491e74b
---

# E3 — Vendor landscape 2026: who makes the memory

- **Contract prices (TrendForce, 3Q26, PC-client OEM SSD):** 1 TB mSATA/M.2 TLC PCIe value-grade averaged **$321 (high $385, low $265)**; 512 GB averaged $201.13; 256 GB averaged $128.38. Updated 2026-08-05 [secondary]. Source: https://www.trendforce.com/price/flash/pcc_oem_ssd_contract
- TrendForce's 3Q26 note: growth in client SSD contract prices **slowed sharply** — surging costs from other component shortages plus excess system inventory made buyers cautious, while suppliers held prices firm by shifting capacity to enterprise SSDs; the stalemate keeps contract prices **plateauing at elevated levels** [secondary].
- TrendForce's 2Q26 note: AI/enterprise demand crowded out consumer-grade supply, a **structural supply shortage**, and suppliers drove "massive spikes in contract prices"; future increases were expected to narrow as cost pass-through hit buyer tolerance limits [secondary].
- **The enterprise SSD vs HDD gap widened violently in 2026.** VDURA's hybrid-storage report (via KAD8, Aug 2026): a **30 TB enterprise TLC SSD averaged ~$22,600 (~$753/TB)** vs a **30 TB enterprise HDD at ~$1,216 (~$40.5/TB)** — an **18.6× ratio in Q3 2026, reportedly reaching 23.2× earlier in 2026**, vs ~7× in Q3 2025 [secondary]. Source: https://www.kad8.com/storage/ssd-vs-hdd-in-2026-why-30tb-ssds-cost-up-to-23x-more/
- Western Digital's own framing: enterprise SSDs carry a **7.4×–9.9× premium over HDDs**, a gap holding above 5× through 2028 [vendor-reported]. VDURA's flash-volatility index pegged 30 TB QLC SSDs at **22.6× the cost of matching HDDs in Q1 2026** [secondary]. Source: https://www.webpronews.com/why-hdds-beat-ssds-for-bulk-storage-in-the-age-of-skyrocketing-nand-prices/
- Street-level illustration of the run-up: the **WD Black SN850X 1 TB NVMe, "once a $100 offering," was retailing above $230** in 2026 [secondary]. Source: https://www.pulse.bot/hardware/news/nand-chip-revenue-in-q1-2026-was-higher-than-the-total-revenue-for-the-entire-year-of-2023-7fb5514b-82ac-4f27-9d56-92e8d8b72062/
- Consumer-grade specifics from a 2026 homelab SSD roundup (prices as observed, ~$): WD Black SN8100 Gen5 2 TB TLC ~$280 ($140/TB); WD Black SN7100 Gen4 2 TB TLC ~$130 ($65/TB); Samsung 990 EVO Plus 2 TB TLC ~$195; Seagate FireCuda 540 Gen5 2 TB TLC ~$180 — all TLC, all with DRAM or HMB [secondary]. Source: https://botmonster.com/self-hosting/best-m2-nvme-ssds-homelab-2026/
- Note the apparent tension: MemoryPriceChart's mainstream-SSD basket median ($159.94/TB, Sept 2026) vs the homelab roundup's $65–$140/TB figures — different baskets, different weeks, different regions; treat as a range, not a contradiction [secondary]. See E17 conflicts.
- Enterprise street pricing (PCServerAndParts, Aug 2026): new NVMe U.2 pricing spanned **$300–$1,172/TB** (Kioxia CM7-R vs Micron 9550 Pro at the same 7.68 TB capacity — a 2.3× spread between generations/tiers, "not one drive at two prices") [secondary]. Source: https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/
- The "SSD prices rival gold" framing (WebProNews, Jan 2026): a top-tier 8 TB NVMe drive under 50 g costing over $1,000 exceeds gold's per-gram price — a rhetorical illustration of the valuation shift, not a purchasing metric [secondary]. Source: https://www.webpronews.com/ssd-prices-rival-gold-in-2026-due-to-nand-shortages-and-ai-boom/

## E3 — Vendor landscape 2026: who makes the memory

- **DRAM triopoly (volume producers): SK hynix, Samsung, Micron** — all three allocating capacity to HBM simultaneously, which is the structural cause of the commodity-DRAM constraint [secondary]. Source: https://www.ainvest.com/news/memory-shortage-market-wrong-2608/
- **SK hynix** crossed the **$1 trillion market-capitalization threshold in late May 2026** (one of three Asian companies alongside TSMC and Samsung at that level) on AI-memory exposure [secondary]. It holds the HBM lead: **~58% HBM market share as of Q1 2026**, leading in HBM3E and HBM4 qualifications [secondary].
- SK hynix's strategic 2026 move: **choosing DDR5 margin over HBM4 volume** — commodity DRAM margins caught or exceeded HBM margins in late 2025, so suppliers balance HBM against high-margin commodity in 2026, which "self-reinforces the dual shortage" [secondary]. Source: https://www.techtimes.com/articles/319016/20260624/sk-hynix-dethroned-samsung-after-26-years-now-choosing-ddr5-profits-over-hbm4-ramp.htm
- **Samsung:** 39.4% DRAM revenue share in Q2 2026, aided by early HBM4 mass production; but its HBM economics lagged (sharpest commodity-vs-HBM margin crossover), and long-term fixed-price HBM contracts delay ASP recognition [secondary]. ~21% HBM share in Q1 2026, described as "recovery mode; yield and qualification delays" [secondary]. Samsung is reviewing conversion of the **S5 foundry line at Pyeongtaek into memory facilities** as early as 2027 [secondary]. Source: https://en.sedaily.com/finance/2026/08/31/spot-hbm-prices-hit-5-times-contract-levels
- **Micron:** 23.3% DRAM revenue share in Q2 2026 (+65.5% QoQ, sharpest increase), ~21% HBM share, entire 2026 HBM supply sold out [secondary].
- **NAND vendors:** Samsung, SK hynix (incl. Solidigm), Micron, **SanDisk** (independent since Feb 2025, ticker SNDK), **Kioxia**, and China's **YMTC** (consumer share 8% → 13% as larger rivals chased enterprise margins) [secondary].
- **YMTC's consumer-share gain (8% → 13%)** is the consumer market's pressure valve: incumbents redirected wafers to enterprise, and YMTC filled the consumer gap [secondary].
- **CXMT** (Chinese DRAM/NAND startup) was reported on 2026-09-18 to be **aiming to join NAND flash production**, potentially adding another rival to the Samsung/SK hynix/Micron/YMTC field [secondary]. Source: https://www.thestreet.com/investing/sandisk-93-billion-in-deals-nand-cost-competition
- SanDisk's leverage: **$93.9 billion in long-term customer commitments/deals** reported Sept 2026, but the open question is conversion to delivered revenue — LTA visibility creates delivery obligations in a shortage [secondary].
- **Kioxia:** Japan's largest NAND maker; SK hynix is effectively its largest shareholder; SanDisk–Kioxia Flash Ventures JVs supply "substantially all" of SanDisk's wafers [secondary].
- Capex/fab news (2026): SanDisk and Kioxia confirmed **> $31 billion investment in Japan through 2032**, including a new memory-production facility at Kioxia's Kitakami plant (announced August 2026) [secondary]. SK hynix's Solidigm unit is **considering a U.S. NAND manufacturing facility** (no final decision at cutoff) [secondary]. SK hynix expanded production bases in Yongin and the Honam region and is weighing a Japan JV plant, possibly with Kioxia cooperation [secondary].
- NAND technology cadence: Samsung teased **7th-gen Z-NAND with GIDS for 2026** alongside its PCIe Gen6 enterprise SSD roadmap [secondary]. (Z-NAND/V-NAND layer-count specifics are in the sibling hardware file.)

## E4 — M&A and corporate events 2025–2026

