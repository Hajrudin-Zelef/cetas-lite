---
id: etape9-phasee-storage-market/00-storage-market/overview
title: "Step 9 Phase E — Storage & Memory Market 2026"
domain: step-9-phase-e-storage-memory-market-2026
role: deep-dive
task: reference
actors: ["China", "Samsung", "TSMC"]
dates: ["2025-08", "2026-05", "2026-06", "2026-08", "2026-08-05", "2026-09-18", "2026-09-22"]
keywords: ["memory", "benchmark", "capex", "consumer", "cost", "dram", "dram price", "foundry", "hbm", "hbm4", "inference", "kv cache"]
source: docs/RAG/etape9_phaseE_storage_market.md
source_anchor: ""
source_lines: [1, 57]
section: "Step 9 Phase E — Storage & Memory Market 2026"
sha256: a0bf4599fd8f603229164607c7dc4f5256f5d2425fa131533ce9fd3a4560d21b
---

# Step 9 Phase E — Storage & Memory Market 2026

## E0 — Scope, date, method, provenance legend

- **Scope:** Storage and memory *market* 2026 — NAND/DRAM price trends, vendor landscape and M&A, HBM economics, the used/refurbished enterprise market for homelabs, AI storage sizing (training checkpoints, inference KV cache), computational storage, and buying guides. Hardware specs (form factors, NVMe, ECC) are covered in sibling Step 9 phase files; this file covers the money, the market structure, and the buyer.
- **Research cutoff:** 2026-09-22. All prices, versions, and market events are dated at or before this cutoff.
- **Method:** Read-only web research via `browser.search` and `browser.open`. Prices are street/retail observations and contract-price indices reported by named sources; they are not quotes and vary by region and week. No vendor data sheets were scraped directly; vendor figures below are cited from the search sources shown.
- **Provenance legend:**
  - `[official]` — vendor press release, spec page, or earnings material.
  - `[vendor-reported]` — vendor-sourced claim relayed by a third party (benchmark, press interview).
  - `[independent]` — independent lab/test publication (e.g., ServeTheHome, Tom's Hardware, Artificial Analysis).
  - `[secondary]` — press or analyst reporting, forums, community docs.
  - `[unverified]` — single-source, rumor-grade, or otherwise uncorroborated at cutoff.
- **Cross-references:** Step 9 sibling files cover SSD hardware (form factors, NVMe, DWPD), memory hardware (DDR/HBM/CXL/ECC), and data-protection hardware (RAID, DLP/PLP). This file does not re-spec them.
- **Non-comparability warning:** Contract prices, spot prices, retail street prices, and auction prices are different measurement regimes. Comparing a TrendForce contract average with an eBay "Buy It Now" listing is invalid. Indices cited below are named with their methodology.

## E1 — The 2026 memory market in one picture

- 2026 is a **seller's market for DRAM and a tight, enterprise-skewed market for NAND** — the AI build-out absorbed wafer capacity while hyperscalers locked in long-term agreements, leaving consumer and spot buyers paying the premium [secondary].
- MemoryPriceChart (independent price index, Kalmar, Sweden) reported on 2026-09-18 that its chain-linked DDR5 index stood **397% above its August 2025 backcast**, while its SSD index was 158% above the same baseline — a 239-percentage-point gap between retail DRAM and retail NAND pricing dynamics [secondary]. Source: https://lifestyle.washingtonguardian.com/story/883602/ram-and-ssd-prices-diverge-as-ddr5-index-rises-397-versus-158-for-ssds/
- The index's latest published weekly point showed a **median of $18.44/GB (mean $18.83/GB) for retail DDR5 kits** (six eligible 32/64 GB desktop kits, four brands) and a **median of $159.94/TB (mean $168.83/TB) for 1–2 TB mainstream PCIe 4.0 NVMe SSDs** [secondary]. Implication: a 64 GB DDR5 kit at ~$18.44/GB costs roughly **$1,180** at retail — system memory, not storage, is now the budget-killer in 2026 builds [secondary].
- The DDR5 index rose 1.7% week-over-week and the SSD index 3.9% week-over-week at the 2026-09-18 observation — the run-up is ongoing at cutoff, not over [secondary].
- TrendForce's Q2 2026 DRAM revenue report put **global DRAM revenue at $154.73 billion for a single quarter (April–June 2026), +59.5% quarter-on-quarter** — a figure described as "unimaginable as recently as 2024" [secondary]. Source: https://www.techtimes.com/articles/326837/20260907/memory-runs-dry-samsung-sk-hynix-drop-below-10-day-supply-hbm4-devours-capacity.htm
- TrendForce Q2 2026 DRAM revenue shares: **Samsung 39.4% ($60.98B, +63.4% QoQ), SK hynix 24.9% ($38.59B, +37.9%), Micron 23.3% ($36B, +65.5%)** — Micron posted the sharpest revenue increase by concentrating capacity on higher-margin server DRAM [secondary]. Even Taiwanese mature-node players (Nanya +68.3%, Winbond +75.8% QoQ) rode the shortage by absorbing demand the big three could not serve [secondary].
- On the NAND side, Pulse reported **Q1 2026 NAND revenue of $46 billion — exceeding total full-year 2023 NAND revenue** — driven by enterprise/AI SSD demand and premium pricing [secondary]. Source: https://www.pulse.bot/hardware/news/nand-chip-revenue-in-q1-2026-was-higher-than-the-total-revenue-for-the-entire-year-of-2023-7fb5514b-82ac-4f27-9d56-92e8d8b72062/
- An investor-oriented sector note (GitHub, updated ~Aug 2026) assessed the NAND pricing-power trend as **strengthening for enterprise SSD (70–75% QoQ price increases in Q2 2026)** on a supply deficit (20–22% demand growth vs 15–17% supply growth, all 2026 production sold out, shift to long-term agreements), and **weakening medium-term for consumer SSD** as YMTC capacity additions (~150K WSPM current, doubling potential) pressure consumer ASPs [secondary]. Source: https://github.com/jameswong2011/investmentvault/blob/HEAD/Sectors/NAND%20Memory%20&%20Storage.md
- The same note put the **1Tb TLC NAND die spot price at $4.80 → $10.70 within six months**, calling NAND a "peak cycle" with structural shortage extending into 2027 and no new fab output before H2 2028 [secondary].
- Structural forecast consensus across sources: **no meaningful relief before mid-2027** (SK hynix M15X and Micron Idaho fab), Samsung Pyeongtaek P5 not until 2028 [secondary]. Source: https://www.techtimes.com/articles/319016/20260624/sk-hynix-dethroned-samsung-after-26-years-now-choosing-ddr5-profits-over-hbm4-ramp.htm

## E2 — NAND price trends: contract, spot, and street

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

