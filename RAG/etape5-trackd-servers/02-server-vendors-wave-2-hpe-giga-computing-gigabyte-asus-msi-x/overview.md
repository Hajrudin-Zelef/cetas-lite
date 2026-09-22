---
id: etape5-trackd-servers/02-server-vendors-wave-2-hpe-giga-computing-gigabyte-asus-msi-x/overview
title: "Server Vendors — Wave 2: HPE + Giga Computing (Gigabyte) + ASUS + MSI + xFusion"
domain: server-vendors-wave-2-hpe-giga-computing-gigabyte-asus-msi-x
role: deep-dive
task: reference
actors: ["AMD", "Intel", "Nvidia", "Oracle", "United States"]
dates: ["2025-07-02", "2025-07-28", "2026-07-31", "2026-08", "2026-09-02", "2026-09-22"]
keywords: ["acquisition", "amd", "backlog", "blackwell", "compute", "cost", "disclosure", "gpu", "gpus", "hbm", "hyperscaler", "inference"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [486, 542]
section: "Server Vendors — Wave 2: HPE + Giga Computing (Gigabyte) + ASUS + MSI + xFusion"
sha256: b2fa824f813e5420c37b9f48013c77516b5b128ad91f6d6c9e9798b89fac0667
---

# Server Vendors — Wave 2: HPE + Giga Computing (Gigabyte) + ASUS + MSI + xFusion
**As of: September 22, 2026** | **RAG collection draft — Step 5 (volet 1, etape 5)**

> Provenance legend used on every factual claim:
> `[official]` = vendor's own press release / datasheet / store page
> `[vendor-reported]` = vendor disclosure via earnings deck, call, or investor filing (management commentary)
> `[independent]` = reputable independent outlet (ServeTheHome, The Register, IDC, MLCommons, etc.)
> `[secondary]` = press/analyst/blog aggregation of primary sources
> `[unverified]` = low-confidence claim, single weak source, or gap explicitly flagged

---

## 1. HPE (Hewlett Packard Enterprise)

### 1.1 AI server lineup (2026)

**HPE ProLiant Compute XD685 (8-GPU node)**
- Support for **8× NVIDIA Blackwell Ultra (B300 HGX)**, **8× NVIDIA Blackwell (B200)**, **8× NVIDIA H200 Tensor Core**, **or 8× AMD Instinct MI355X** accelerators [official — HPE Store UK product page, SKU #S4Q28A, https://buy.hpe.com/uk/en/compute/rack-servers/proliant-compute-xd600-servers/proliant-compute-xd600-server/hpe-proliant-compute-xd685-air-cooling-server/p/s4q28a].
- Direct liquid cooling (DLC) available for **all** GPU options; air cooling available for H200 configurations only [official — same HPE store page].
- Modular 5U chassis for DLC environments / 6U chassis for air-cooled configurations; 2× 5th Gen AMD EPYC processors (up to 5.0 GHz, 32–160 cores) [official — HPE datasheet].
- 24× DDR5-6400 RDIMM slots, 12 memory channels per CPU, ECC; HPE iLO 6 management with silicon root of trust / Zero Trust [official — HPE store page].
- The XD685 was originally announced for 5th Gen AMD EPYC AI training workloads (announcement covered by VarIndia) [secondary — https://varindia.com/news/hpe-unveils-new-solutions-to-boost-ai-model-training].
- ⚠️ **[unverified/gap]**: The task mentions **ProLiant Compute XD680**; the official HPE store source consulted lists only the XD685 with B300/B200/H200/MI355X options. The Register (Nov 2024) previously reported an XD680 variant configured with 8× Intel Gaudi3 accelerators [independent — The Register, https://www.theregister.com/on-prem/2024/11/13/hpe-crams-224-nvidia-blackwell-gpus-into-latest-cray-ex/]. Whether XD680 remains a live 2026 SKU with updated GPU options was not confirmed — flagged as a gap.

**HPE Cray XD670 (HPC/AI, Intel CPU base)**
- 5U chassis, single node; **8× NVIDIA H200 SXM (700 W TDP, 141 GB HBM3e each)** or H100; 2× 5th Gen Intel Xeon Scalable (up to 400 W TDP); 32× DDR5 DIMM slots (8-channel/socket, up to 5600 MT/s); PCIe Gen5; optional plug-and-play **direct liquid cooling**; integrated storage [official — HPE Cray XD670 datasheet, https://www.hpe.com/psnow/generateDDS/HPE%20Cray%20XD670%20data%20sheet-PSN1014737116TREN.pdf].
- MLPerf Inference v5.0 top rankings cited in Subaru win [vendor-reported].

**HPE Cray XD675 (Blackwell)** ⚠️ **[unverified/gap]**: The task brief lists Cray XD675 with NVIDIA B200, but the research surfaced no 2026-dated primary source confirming an XD675 SKU. The Register (Nov 2024) described Cray EX Blackwell blades (224 GPUs per cabinet, shipping late 2025) [independent], not an XD675. Treat XD675 as an unconfirmed product name.

**Liquid cooling** — HPE markets decades of liquid-cooling deployment leadership; DLC options span XD685 (all GPU options), Cray XD670, and rack-scale solutions [official].

### 1.2 Juniper Networks acquisition — status as of Sept 2026

- Deal **closed July 2, 2025**; Juniper stock (JNPR) ceased trading on NYSE and the company was integrated into HPE [secondary — https://www.levelfields.ai/news/hpe-finalizes-14-billion-juniper-networks-acquisition; https://channellife.ca/story/hpe-completes-juniper-networks-deal-to-boost-ai-cloud-focus]. Deal value widely cited as **~$14 billion** [secondary — multiple outlets].
- **August 2026 legal milestone**: U.S. District Judge Casey Pitts dismissed a challenge by a Colorado-led coalition of state attorneys general against the DOJ settlement, issuing a 41-page ruling finding the states had not shown the amended final judgment was against the public interest. This gave HPE final court approval to operate as an integrated entity [secondary — https://www.ainvest.com/news/hpe-clears-juniper-deal-ai-hardware-stocks-rally-2608/]. Date calibration: the article was "last updated 39 days ago" on ~Sept 22, 2026 → ruling **mid-August 2026**.
- DOJ settlement terms (agreed 2025): HPE **divested its Instant On wireless networking business** and **auctioned a license for Juniper's Mist AI software** [secondary — ainvest article above].
- Integration: former Juniper CEO **Rami Rahim** leads the combined **HPE Networking** business unit (Juniper portfolio + HPE Aruba); networking business doubled in size; networking now accounts for **over 50% of HPE's total operating income** [secondary — levelfields]. HPE says integration is **ahead of plan**, targeting **$600M annualized cost synergies by end of fiscal 2028**; net leverage fell to **1.8×** (below the 2× target) more than a year early [secondary — Zacks, https://www.zacks.com/stock/news/2984471/hpe-q3-earnings-surpass-expectations-revenues-rise-yy].
- **AI networking implications**: "Networks for AI" orders of **$0.7B in Q3 FY2026** (quarter ended July 31, 2026), **$2.2B cumulative**, with HPE raising its fiscal-2026 target to **$2.5–$3B** [vendor-reported via Zacks — https://www.zacks.com/stock/news/2984471/hpe-q3-earnings-surpass-expectations-revenues-rise-yy]. Reported networking revenue was $2.89B (+74.9% YoY as reported, including Juniper; **+10% normalized** for the acquisition) — the two figures must not be conflated [vendor-reported via fourweekmba/payney analyses].

### 1.3 AI customer wins (2026)
- **$3.5 billion inferencing deal with an unnamed hyperscaler**, awarded **after quarter-end** (i.e., August 2026), disclosed on HPE's Q3 FY2026 earnings deck and call [vendor-reported — HPE Q3 FY26 earnings slides, https://files.quartr.com/conference-calls/cdd89295c3957b4612350653f2c0618f-2026-09-02-20-15-50.pdf].
- **Oracle gigawatt-scale networking deal**: expanded arrangement for routers/switches supporting Oracle's AI data center buildout, announced Sept 2, 2026; includes **HPE warrants issued to Oracle** — share count, strike price, conditions, and dilution all **undisclosed** [secondary — SiliconANGLE, https://siliconangle.com/2026/09/02/ai-demand-lifts-hpe-and-netapp-to-record-quarters-investors-sell-both/].
- **Subaru (2025 win, still in-market)**: HPE Cray XD670 with NVIDIA H200 GPUs for next-gen EyeSight AI driver-assist development (image recognition training/inference), on-prem deployment + HPE Tech Care; announced **July 28, 2025** [official — Business Wire via StockTitan, https://www.stocktitan.net/news/HPE/subaru-selects-hpe-to-accelerate-ai-development-for-next-generation-85xppynm4hze.html].

### 1.4 AI revenue figures — Q3 FY2026 earnings (reported Sept 2, 2026; quarter ended July 31, 2026)
- Total revenue **$12.2B**, +34% YoY (above high end of guidance); non-GAAP diluted EPS **$1.11**; non-GAAP gross margin record **40.4%** [vendor-reported — press release + SEC 8-K, via https://fourweekmba.com/ai-hpe-q3-fy2026-ai-supply-constraint-memory-bottleneck/].
- **AI systems orders: $2.4B** in the quarter, up **>30% sequentially**; **AI systems backlog: $6.8B** (+14% sequentially) [vendor-reported — HPE earnings deck via Zacks, https://www.zacks.com/stock/news/2984417/hpe-q3-earnings-call-highlights-ai-demand-networking-shift].
- **Total AI backlog (AI Systems + Networks for AI): record $7.6B** exiting the quarter [vendor-reported — same sources; also https://payney.com/blog/hpe-q3-2026-122b-revenue-fy26-eps-outlook-raised/].
- Segment: **Cloud & AI revenue $9.04B, +25.4% YoY**, operating margin 17%; **server revenue $6.77B, +35.3%**; traditional server orders +75% YoY (AI-ready configs, higher ASPs); Private Cloud AI orders up **triple digits** YoY [vendor-reported — Zacks, https://www.zacks.com/stock/news/2987783/hpe-surges-26-in-3-months-is-it-the-right-time-to-buy-the-stock].
- Supply constraint: management attributed the gap between normalized orders (+42%) and revenue (+34%) to **DDR5, DDR4, NAND, HBM, and wafer capacity** — expected to persist into fiscal 2027 [vendor-reported].
- Guidance: FY2026 revenue growth raised to **34–37%** (~$46.5B midpoint), non-GAAP EPS **$3.75–$3.85**, free cash flow ≥**$3.75B**; Q4 FY26 revenue guided **$13.9–$14.8B**; FY2027 framework raised to **13–17% revenue growth**, EPS **$4.4–$4.6**, FCF ≥**$5B** [vendor-reported — Zacks].
- Market context: IDC Q2 2026 — HPE ranked **#4 server vendor worldwide**, **$5.87B revenue, 3.5% share, +46.0% YoY** [independent — IDC via StorageReview, https://www.storagereview.com/news/idc-external-storage-tracker-q2-2026-10-3-billion-as-server-market-tops-166-billion].
- ⚠️ **Price gap**: no public per-server pricing found for XD685/XD670/Cray systems (enterprise quotes only).

---

