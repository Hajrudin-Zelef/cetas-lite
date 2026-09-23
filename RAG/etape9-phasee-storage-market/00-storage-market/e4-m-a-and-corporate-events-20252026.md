---
id: etape9-phasee-storage-market/00-storage-market/e4-m-a-and-corporate-events-20252026
title: "E4 — M&A and corporate events 2025–2026"
domain: step-9-phase-e-storage-memory-market-2026
role: deep-dive
task: reference
actors: ["AWS", "China", "Google", "Intel", "Meta", "Microsoft", "Nvidia", "Samsung", "United States"]
dates: ["2025-01-02", "2025-02-24", "2026-02-18", "2026-03", "2026-07", "2026-07-21", "2026-08-31", "2026-09-18"]
keywords: ["accelerator", "capex", "compute", "consumer", "cost", "datacenter", "distribution", "dram", "gpu", "gpus", "hbm", "hbm4"]
source: docs/RAG/etape9_phaseE_storage_market.md
source_anchor: ""
source_lines: [58, 113]
section: "Step 9 Phase E — Storage & Memory Market 2026"
sha256: e7a1edfa4a1d4b7205034813b71eb60a2725b9f1a280b93bb6761da17a0e1c86
---

# E4 — M&A and corporate events 2025–2026

## E4 — M&A and corporate events 2025–2026

- **Western Digital → SanDisk separation (completed 2025-02-24):** WD spun off its flash/NAND business into **SanDisk Corporation (NASDAQ: SNDK)** via pro-rata distribution on 2025-02-24; WDC retained HDD in full, SanDisk retained flash in full [secondary]. Source: https://github.com/mast3rkey/portfolio-hq/blob/HEAD/intelligence/WDC_SANDISK_COMPARISON.md
- **Final tie-cut (2026-02-18):** Western Digital exited its remaining stake — **7,513,019 SanDisk shares (~$3.1B)** — via a debt-for-equity exchange with JPMorgan and Bank of America affiliates, retiring long-term debt without the tax cost of a cash sale; WDC is now a **pure-play HDD manufacturer** under CEO Irving Tan, SanDisk fully autonomous under CEO David Goeckeler [secondary]. Source: https://markets.financialcontent.com/buffnews/article/marketminute-2026-2-18-western-digital-cuts-final-ties-with-sandisk-31-billion-exit-sends-shares-tumbling
- SanDisk's structural cost model (from its disclosures as summarized in a community evidence bundle): **substantially all flash wafers come from Flash Ventures JVs (49.9%-owned with Kioxia)**; SanDisk pays half of the JVs' fixed costs **regardless of output** and funds ~half of approved JV capex — a materially different fixed-cost exposure than WDC's HDD model [secondary].
- **SanDisk stock trajectory:** surged ~470% year-over-year to **$1,791.82 on 2026-09-18** (per Reuters, via TheStreet); "one of the year's top performers," up >140% since re-listing in early 2025 [secondary].
- **WDC–Kioxia merger talks (restarted ~July 2026, rumor-grade):** Western Digital shares jumped 12.51% to $548.39 on 2026-07-21 on reports that merger talks with Kioxia restarted; talks in 2021 and Oct 2023 had collapsed (2023 blocked by SK Hynix opposition) [secondary]. Source: https://www.tradingkey.com/analysis/stocks/us-stocks/262047263-western-digital-wdc-stock-surges-july-22-2026-kioxia-merger-hdd-sold-out-tradingkey
- Combined WDC(flash-era)+Kioxia NAND shipment share would be **~25%** per Counterpoint Q2 2026 data — surpassing SK hynix (22%) and matching Samsung — but the merger faces three obstacles: Bain Capital has exited its position, **SK Hynix's veto power remains**, and the valuation gap is widening; assessed probability of success is low [secondary]. Source: https://www.mitrade.com/insights/news/live-news/article-8-2002414-20260815
- WDC's HDD side: **2026 HDD production capacity reportedly already sold**, with contracts through 2028 — the HDD business is itself supply-constrained, which is why NAND/HDD substitution economics matter less than the ratio suggests [secondary].
- US policy context: Commerce signals that the **ban on Chinese memory chips will hold**, a tailwind for WDC/SanDisk and a constraint on YMTC/CXMT expansion into US supply chains [secondary].

## E5 — HBM economics 2026: the memory that prices the GPU

- **Why HBM dominates the market narrative:** HBM is 3–4× more wafer-intensive per gigabyte than commodity DRAM ("reverse scaling" — a commodity wafer yields ~3× the bits of an HBM3E 12-Hi wafer, ~4× at HBM4), so every wafer shifted to HBM tightens the entire memory market [secondary]. SemiAnalysis (Feb 2026) put DRAM **~7% below demand in 2026 and 2027**, with the HBM-specific shortfall widening 5% → 6% → 9% (2025→2027) [secondary]. Source: https://github.com/jameswong2011/investmentvault/blob/HEAD/Theses/000660%20-%20SK%20Hynix.md
- **Allocation shares (Q1 2026, latest public at cutoff):** SK hynix ~58%, Samsung ~21%, Micron ~21% of HBM market [secondary]. For NVIDIA's **Rubin platform (R200-class) first-12-months supply**, SemiAnalysis estimated **~60% SK hynix / ~30% Samsung / ~0% Micron**, with Korean samples at ~10 Gbps pin speed and SK hynix retaining a signal-integrity/jitter packaging lead [secondary].
- **2026 HBM production sold out months in advance** across all three suppliers to AI datacenter customers (NVIDIA, Microsoft, Google, Meta, Amazon) [secondary]. Source: http://gpusmith.com/articles/en/pdfs/ai-gpu-prices-rising-hbm-memory-shortage.pdf
- **Spot vs contract:** a 36 GB HBM3E product traded at **$2,100 on spot vs ~$500,000–700,000 KRW long-term contract** — 4–5× contract levels; 16-layer HBM4 (pre-mass-production) traded at **$3,500 spot** (MegaGrid Supply data via Seoul Economic Daily, 2026-08-31) [secondary]. Source: https://en.sedaily.com/finance/2026/08/31/spot-hbm-prices-hit-5-times-contract-levels
- Even contracted hyperscalers "are unable to receive all of their contracted volumes" — the spot market is a rationing mechanism, not a price-discovery one [secondary].
- HBM3E 12-Hi 2026 contract pricing was **tracking flat vs the prior -15–20% consensus** — i.e., the expected price decline did not happen [secondary].
- **HBM generation timing:** HBM3E dominated AI-accelerator supply through 2025 and into 2026; **HBM4 in meaningful volume expected only from Q3 2026** (TrendForce) [secondary]. NVIDIA's H200 uses six HBM3E stacks per unit; demand for that generation outstripped capacity even as HBM4 ramped [secondary].
- **Inventory signal (Sept 2026):** Samsung and SK hynix reportedly dropped **below 10 days of DRAM inventory** as HBM4 devoured capacity; the 12-die HBM3E → 16-die HBM4 transition makes this generation's displacement effect ~one-third worse per module [secondary]. Source: https://www.techtimes.com/articles/326837/20260907/memory-runs-dry-samsung-sk-hynix-drop-below-10-day-supply-hbm4-devours-capacity.htm
- **The GPU-price transmission:** VRAM (GDDR on consumer/pro cards) now accounts for **>80% of a high-end GPU's bill of materials**, up from a much smaller historical share; the RTX 5090 rose from its $1,999 launch price to **$4,329 on Amazon by mid-July 2026**, with some listings over $5,000; 16 GB+ VRAM cards broadly saw double-digit percentage hikes as board partners passed through GDDR7 cost increases [secondary]. Source: http://gpusmith.com/articles/en/pdfs/ai-gpu-prices-rising-hbm-memory-shortage.pdf
- **Cost-per-GB framing:** at $3,500 spot for a 16-layer HBM4 stack (~36 GB-class), HBM4 is on the order of **~$95–100/GB** — roughly 5,000× the $0.018/GB retail DDR5 figure; even contract HBM (~$500K–700K KRW per 36 GB stack) is orders of magnitude above commodity DRAM. HBM is priced as a strategic allocation, not a commodity [secondary].

## E6 — CXL: the "post-HBM battleground"

- As HBM and DDR approach capacity limits for AI workloads, Samsung, SK hynix, and Micron are racing on **Compute Express Link (CXL) memory** — a separate, lower-cost memory tier that expands capacity without additional GPUs/CPUs and enables pooling [secondary]. Source: https://www.trendforce.com/news/2026/07/21/news-samsung-reportedly-targets-2026-cxl-3-2-mass-production-sk-hynix-advances-new-ai-memory-architecture/
- **Samsung:** CMM-D 1.0 (CXL 1.1, 2022) → CMM-D 2.0 (CXL 2.0, 2023) → **CMM-D 3.0 (CXL 3.1, targeting 2026-12 mass production, with CXL 3.2 upgrade in progress)** — industry-first claim; CXL 3.2 introduces the **CXL Hot Page Monitoring Unit (CHMU)** for frequency-aware page placement [secondary]. Source: https://en.sedaily.com/finance/2026/07/19/cxl-memory-war-erupts-as-next-memory-battle-heats-up
- Samsung claims CXL memory can **increase total memory capacity by up to 50% and double memory bandwidth** vs DDR5-only systems, and support memory pooling across servers [vendor-reported]. Samsung verified CXL memory operations **in a real user environment with Red Hat** [vendor-reported]. Source: https://www.techradar.com/pro/samsung-reveals-major-step-forward-in-data-center-memory-supply
- Samsung plans to unveil **"Pangea v3" (CXL 3.2-based) within 2026**, and reported a 10× performance gain on a demo system [secondary]. Source: https://www.trendforce.com/news/2026/04/28/news-cxl-may-emerge-as-post-hbm-battleground-samsung-reportedly-presents-system-with-10x-performance-gain/
- **SK hynix:** first CXL DRAM (2022), CXL 2.0 product (2023, customer validation completed 2025), now advancing **second-generation CXL 3.0** products [secondary].
- **Micron:** unveiled its own CXL memory module in 2024 and is in the race [secondary].
- **Demand-side signals:** Google is **reported deploying CXL controllers in its datacenters** (The Information via Digital Today, March 2026); **NVIDIA plans CXL 3.1 support in its Vera CPU** (expected late 2026) — potentially the largest real-world CXL test to date [unverified for both; single-press-chain reporting].
- **WCCFTech (GMIF Innovation Summit 2025):** Samsung's Kevin Yoon said **CXL 3.1/PCIe 6.0 CMM-D products arrive in 2026** alongside the PM1763 Gen6 SSD [secondary]. Source: https://wccftech.com/samsung-512-tb-pcie-gen6-ssds-2027-innogrit-preps-gen6-ai-nvme-cxl-enterprise/
- CXL also absorbs part of the old "computational storage" story (see E15): memory-semantic SSDs (CMM-H/CXL-attached NAND) blur the line between storage and memory expansion — but as of cutoff, **CXL-attached NAND remains a vendor roadmap item, not a shippable product line** [unverified].

## E7 — The used/refurbished enterprise market (homelab angle)

- **Why it exists:** 3–5-year enterprise refresh cycles dump read-intensive SSDs and DDR4 ECC RDIMMs into the secondary market at prices decoupled from current NAND/DRAM contracts [secondary].
- **Price observations (PCServerAndParts, Aug 2026):**
  - Used/refurbished 3.84 TB enterprise SATA/SAS SSDs: **$78–$127/TB** vs **$300/TB+ for new NVMe U.2** at the same era — the used market prices off decommissioned pools, "which do not know what wafers cost this month" [secondary].
  - Used enterprise 8 TB 3.5" SATA 7.2K HDD: **$149.99 (~$19/TB)**; recertified 16 TB Exos X16: **$663.97 (~$41/TB)** [secondary].
  - Rule of thumb given: used enterprise SSD at $78/TB vs nearline SATA at $19–$25/TB is **3–4×, not the 18×** of new 30 TB drives vs new HDDs [secondary]. Source: https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/
- **Community price points (homelab planning guide, Sept 2026):** 2.5" SATA SSDs 240 GB $20–25 / 500 GB $35–45 / 1 TB $60–80; 3.5" HDDs 2 TB $20–30 used / 4 TB $40–60 / 8 TB $80–120 [secondary]. Source: https://github.com/jaded423/homelab/blob/HEAD/docs/homelab-planning-guide.md
- **Where homelabbers shop:** eBay (best selection), r/homelabsales, r/hardwareswap, TechMikeNY, NewServerLife, Amazon Renewed (pricier), **ServerPartDeals** (enterprise drives), plus local Craigslist/Facebook Marketplace and government/university surplus [secondary].
- **ECC RDIMM note (2026):** on used enterprise gear, **ECC RDIMMs are "the cheapest RAM you can buy anyway"** — take the protection because it costs nothing extra [secondary]. Source: https://pcserverandparts.com/news/best-used-server-for-home-lab-2026/
- **What to check before deploying a used SSD** (the 15-minute `smartctl` routine) [secondary]:
  - NVMe: **Percentage Used** (endurance consumed; 100 = rated life gone), Available Spare, Data Units Written, Power On Hours.
  - SAS: percentage-used endurance indicator. SATA: **Media_Wearout_Indicator** (counts down).
  - Vendor wear attributes are inconsistent — cross-check total bytes written against rated endurance (DWPD × capacity × 365 × warranty years) [secondary].
  - Example from a homelab blog: an Intel DC S3610 1.6 TB (MLC, TBW rated 10.7 PB) showed **0% wear after 44,123 power-on hours (~5 years)** — read-intensive enterprise drives often die of old age with endurance untouched [secondary]. Source: https://github.com/mkuthan/mkuthan.github.io/blob/HEAD/_posts/2025-01-02-homelab-upgrade.md
- **When used is the wrong call** (PCServerAndParts, 2026): write-heavy tiers, hardware on a vendor compatibility list, unmirrored single points of failure, and anywhere nobody will check wear [secondary].
- **Endurance math for buyers:** 1 DWPD × 3.84 TB × 365 × 5 years ≈ **7 PB rated lifetime writes**; file shares, VM datastores, backup targets and boot volumes "come nowhere near that"; reserve 10 DWPD write-intensive drives (bought new) for transaction logs and write caches [secondary].
- **Refurbished-seller checklist** (server-parts.eu): ISO 9001 certification, detailed SMART reports (TBW, DWPD, power cycles), **minimum 3-year warranty (ideally 5)**, verified origin (decommissioned from Tier-1 DC, not failed stock), and verified power-loss protection [vendor-reported]. Source: https://www.server-parts.eu/post/refurbished-enterprise-ssd
- Spiceworks community (Sept 2026): Dell enterprise SSDs running ~10 years in SMB file-server duty still healthy; advice — keep a compatible spare, run a third-party wear checker, and **test the restore path before migration** [secondary]. Source: https://community.spiceworks.com/t/sethoughts-on-re-using-ssds-for-a-server-upgrade/1257813

