---
id: etape9-phasee-storage-market/00-storage-market/e22-2026-timeline-of-market-events
title: "E22 — 2026 timeline of market events"
domain: step-9-phase-e-storage-memory-market-2026
role: deep-dive
task: reference
actors: ["AWS", "China", "Google", "Nvidia", "Samsung", "United States"]
dates: ["2026-02-18", "2026-03-24", "2026-04-18", "2026-06-24", "2026-07-10", "2026-07-19", "2026-07-21", "2026-07-22", "2026-08-05", "2026-08-15", "2026-08-31", "2026-09", "2026-09-07", "2026-09-15", "2026-09-16", "2026-09-18", "2026-09-22"]
keywords: ["benchmark", "capex", "consumer", "cost", "dram", "foundry", "gpu", "gpus", "hbm", "hbm4", "inference", "ipo"]
source: docs/RAG/etape9_phaseE_storage_market.md
source_anchor: ""
source_lines: [420, 476]
section: "Step 9 Phase E — Storage & Memory Market 2026"
sha256: b764b513e2e7f757a9372dfde6cf8c9dbb142613cfdd74999f3cb874524af6e3
---

# E22 — 2026 timeline of market events

## E22 — 2026 timeline of market events

- **2026-02-18:** Western Digital completes final exit from SanDisk — 7.5M shares (~$3.1B) via debt-for-equity exchange; WDC becomes pure-play HDD, SanDisk fully independent [secondary].
- **2026-03-24:** (Adjacent) TechEmpower Round 24 repo archived — cited in sibling files; no direct storage link, but marks the benchmark drought of 2026 [secondary].
- **2026-04-18:** Google Research's TurboQuant (KV-cache 3–4-bit compression) published at ICLR 2026 — the most-cited answer to the inference memory wall [secondary].
- **2026-05:** SK hynix crosses $1T market cap on AI-memory exposure [secondary].
- **2026-06-24:** SK hynix "choosing DDR5 profits over HBM4 ramp" analysis — commodity DRAM margins ≥ HBM margins, reinforcing the dual shortage [secondary].
- **Q2 2026:** Enterprise SSD pricing power "strengthening" — 70–75% QoQ price increases; all 2026 HBM production sold out; DRAM quarterly revenue $154.73B (+59.5% QoQ) [secondary].
- **2026-07-10:** gpusmith LLM inference sizing guide — H200 at $30–40K/card, $3.72–10.60/GPU-hour rental (Jan 2026 figures) [secondary].
- **2026-07-19:** Seoul Economic Daily — Samsung CMM-D 3.0 (CXL 3.2) mass production targeted end of 2026 [secondary].
- **2026-07-21:** WDC–Kioxia merger talks reportedly restart; WDC +12.51% to $548.39 [secondary].
- **2026-07-22:** Reports: WDC's 2026 HDD capacity sold out, contracts through 2028 [secondary].
- **2026-08-05:** TrendForce updates 3Q26 PC-client OEM SSD contract prices (1 TB avg $321) — increases "narrowing" as buyers hit tolerance [secondary].
- **2026-08-15:** Mitrade assessment — WDC–Kioxia merger probability "remains low" (Bain exited, SK hynix veto stands) [secondary].
- **2026-08-31:** Seoul Economic Daily — spot HBM at 4–5× contract; Samsung reviewing S5 foundry→memory conversion [secondary].
- **2026-09-07:** TechTimes — Samsung/SK hynix below 10 days DRAM inventory [secondary].
- **2026-09-15:** Memeburn local-AI memory guide — 24 GB VRAM as the serious-tier bar [secondary].
- **2026-09-16:** Spiceworks thread — 10-year-old Dell enterprise SSDs still healthy in SMB duty [secondary].
- **2026-09-18:** MemoryPriceChart — DDR5 index +397% vs Aug 2025, SSD index +158%; DDR5 $18.44/GB, SSD $159.94/TB medians [secondary]. Reuters via TheStreet — SanDisk $93.9B in commitments, shares ~470% YoY [secondary]. Reuters — CXMT aiming at NAND entry [secondary].
- **2026-09-22:** Research cutoff for this file.

## E23 — Regional price notes

- **United States:** deepest used-enterprise market (eBay, ServerPartDeals, TechMikeNY); 2026 HDD capacity sold out at WDC with contracts through 2028; US Commerce signals holding the Chinese-memory-chip ban — a structural tailwind for domestic supply chains [secondary].
- **Europe:** retail DDR5/SSD indices (MemoryPriceChart, Sweden-based) are the cleanest public price time-series; refurbisher market (e.g., server-parts.eu) emphasizes ISO 9001 + 3–5-year warranties [secondary].
- **Asia-Pacific:** fab/construction news concentrated in Japan (Kitakami $31B SanDisk–Kioxia), Korea (Yongin/Honam, Pyeongtaek conversions), China (Wuhan Phase 3, Xian plant pivots); YMTC's 10–20% price edge is most felt in China-domestic channels [secondary].
- **Currency effects:** KRW/JPY weakness in 2026 flatters Korean/Japanese vendors' USD revenue while raising their imported-equipment costs — one reason vendor USD revenue shares move faster than bit shares [secondary].
- **Gray-market caution:** iDRAC/iLO license keys and OEM sleds are the two most counterfeited/fragile items in cross-region used purchases; confirm license status in writing before paying [secondary].

## E24 — TCO worked examples ($/TB/year)

- **Example A — 40 TB usable media pool (RAIDZ2, 6× 8 TB):** used 8 TB HDDs at $100/drive = $600 + ~30 W idle (~$40/year power) + HBA $60 → ~$700 first year ≈ **$17.50/TB/year**; SSD equivalent (used enterprise 3.84 TB × 12, RAIDZ2) ≈ $3,600+ → **~$90/TB/year** — 5× for IOPS the workload doesn't need [secondary].
- **Example B — 10 TB Proxmox VM datastore (mirror):** 2× used 3.84 TB enterprise SATA SSD at ~$300 each = $600 ≈ **$78/TB capex**, 5-year expected service → **~$16/TB/year** before power; vs 2× new 4 TB consumer NVMe at ~$400 each = $800 with lower endurance but simpler integration [secondary].
- **Example C — 64 GB DDR5 workstation kit (Sept 2026):** 64 × $18.44 ≈ **$1,180** — more than the SSD, motherboard, and PSU combined in many 2026 builds; the used-DDR4 route (server platform) cuts this to ~$100–200 [secondary].
- **Example D — Checkpoint storage for a 70B fine-tune:** 700 GB/checkpoint × 8 retained = 5.6 TB on NVMe scratch (one-time capex) + S3-class object for archive at ~$20/TB/month → **~$112/month** retained — trivial next to GPU rental ($3.72–10.60/GPU-hour × hundreds of GPUs) [secondary].
- **TCO rule:** in 2026, $/TB/year favors HDD for bulk, used enterprise flash for active tiers, and new consumer TLC for single-drive simplicity; $/GB/year overwhelmingly favors keeping data *out* of DRAM/HBM [secondary].

## E25 — Watch list (post-cutoff)

- W1. SK hynix M15X and Micron Idaho fab ramps — first "meaningful" new DRAM supply, expected mid-2027.
- W2. Samsung Pyeongtaek P5 — 2028 production target.
- W3. HBM4 volume ramp Q3 2026 → Q1 2027 and Rubin-platform allocation splits.
- W4. Samsung CMM-D 3.0 CXL 3.2 mass production (targeted end of 2026) and Google/NVIDIA Vera CXL deployments.
- W5. WDC–Kioxia merger talks — SK hynix veto status.
- W6. YMTC mainland IPO and enterprise-SSD push; CXMT NAND entry.
- W7. SanDisk–Kioxia Kitakami $31B build milestones through 2032.
- W8. PCIe Gen6 enterprise SSD street pricing (Samsung PM1763, Micron 9650) once qualified.
- W9. SanDisk HBF consortium standardization progress (SK hynix + SanDisk).
- W10. MemoryPriceChart DDR5/SSD indices — whether the September 2026 acceleration (+1.7%/+3.9% weekly) sustains into Q4.

## E26 — Methodology note: how to read 2026 storage prices

- **Four price regimes, never mixed:** (1) fab/wafer economics (die $/wafer, bit growth), (2) contract prices (TrendForce OEM indices — large-buyer, quarterly), (3) retail street (MemoryPriceChart baskets, Newegg/Amazon — small-buyer, weekly), (4) secondary market (eBay/ServerPartDeals — depreciation-driven, per-listing). A claim like "SSDs cost $X/TB" is meaningless without the regime [secondary].
- **Indices cited here:** MemoryPriceChart (retail, chain-linked, Aug 2025 backcast; six DDR5 kits / six SSDs), TrendForce PC-client OEM SSD contract (quarterly buyer survey), VDURA flash-volatility index (enterprise 30 TB class). Each has selection bias; direction is more reliable than level [secondary].
- **Why ratios move faster than prices:** the 30 TB SSD:HDD ratio went 7× → 23× → 18.6× within a year mostly because *enterprise SSD* prices moved — HDD $/TB barely budged (~$40/TB throughout) [secondary].
- **Survivorship in community data:** homelab blogs report drives that survived; failed drives don't get blog posts. Treat anecdotal longevity (10-year-old SSDs, 0%-wear pulls) as existence proofs, not failure-rate data [secondary].
- **Vendor guidance vs street:** vendor $/TB figures (e.g., WD's 7.4–9.9× premium framing) describe portfolio positioning, not transaction prices — always paired here with an independent or street observation [vendor-reported].
