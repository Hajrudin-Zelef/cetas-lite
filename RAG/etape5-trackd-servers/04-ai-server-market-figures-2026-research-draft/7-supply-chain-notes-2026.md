---
id: etape5-trackd-servers/04-ai-server-market-figures-2026-research-draft/7-supply-chain-notes-2026
title: "7. Supply chain notes 2026"
domain: ai-server-market-figures-2026-research-draft
role: deep-dive
task: reference
actors: ["AMD", "Google", "Nvidia", "OpenAI", "Samsung", "TSMC"]
dates: ["2026-07"]
keywords: ["advisory", "amd", "asic", "backlog", "compute", "cowos", "datacenter", "dram", "energy", "gpu", "gpus", "hbm"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [1298, 1352]
section: "AI Server Market Figures (2026) — Research Draft"
sha256: 5e4974f1b0531fac941c073adb41fd80550cac977f7b151e5730d7c3daa68626
---

# 7. Supply chain notes 2026

## 7. Supply chain notes 2026

### 7.1 GPU allocation and pricing
- IDC: "supply — not demand — is the primary ceiling on near-term server market growth" (Q1 2026). [official]
- B200 street pricing under 2026 supply constraints: **$45–55K/unit** in 8-GPU volumes vs. $30–40K list; resale/residual value at **158% of launch price** (Silicon Data, Sept 2026). [unverified]
- Dell (Sept 2026): memory component availability is the **primary constraint** limiting total shipment volume; high demand keeping backlog elevated through year-end. [independent — earnings coverage]
- Supermicro CEO (Aug 2026): some Q4 revenue pushed by customer delays tied to **power, cooling, and networking infrastructure** — "the next bottleneck in AI may not be compute itself." [secondary]
- AIwire/Gartner (Feb 2026): demand from hyperscalers for AI-optimized servers shows no sign of plateauing; AI infrastructure investment "moved from cyclical to durable." [secondary]

### 7.2 HBM and memory
- **All three HBM suppliers' 2026 capacity sold out** — booked by NVIDIA/AMD/Google TPU by end of Q1 2026; shortage forecast to persist **through 2028** (sell-side). [unverified — Momoview industry analysis]
- TrendForce: HBM consumption **+70% in 2026**; HBM3→HBM3E→HBM4 per-stack pricing ≈ **$200→$300→$500** (TrendForce/Silicon Analysts estimates). [official / unverified]
- TrendForce: SK Hynix holds **>50% of global HBM market**, Samsung and Micron ~25% each. [secondary — via Notebookcheck on TrendForce]
- Customer allocation (Q1 2026, industry analysis): SK Hynix ~70% of NVIDIA Rubin HBM4; Samsung >60% of Google TPU HBM3E + AMD MI455X HBM4 partnership; Micron #2 NVIDIA slot + dual-source on AMD MI350. [unverified]
- Server DRAM contract prices: **+90–95% in Q1 2026**; **+50–55% QoQ in Q2 2026** (TrendForce); enterprise NAND +55–60% in Q1 2026; 16GB server memory modules doubled in six months; high-end enterprise SSDs nearly doubled. [secondary]
- Commercial DRAM prices up **~700% since 2022** (TrendForce via fiisual). [secondary]
- Micron: 1 unit of HBM3E consumes ~**3× the wafer capacity** of the same bits in DDR5; TrendForce: 1GB HBM ≈ 4GB standard DRAM in wafer area; revenue per wafer for HBM runs 3–5× DDR5. [secondary]
- Fulfillment: DRAM suppliers meeting only **75–80% of demand in 2H 2026**, possibly ~60% in 2027 (Meritz Securities); Micron meeting **50–65%** of key customer requests (CEO Sanjay Mehrotra). [secondary]
- SK Hynix: entire 2026 HBM output sold; signed supply agreements with OpenAI for Stargate, expected to **more than double the industry's total HBM requirements**. [secondary — Notebookcheck]
- Gartner analyst Shrish Pant: HBM wafer reallocation "very real," impacting the market "till the end of 2027"; meaningful price relief unlikely before late 2027. [secondary — TechTimes/Computex 2026]
- Samsung memory chief (Apr 2026 earnings): "significant shortages" across memory products expected **through at least 2027**. [secondary — TechTimes]
- TSMC CoWoS capacity: ~35K wafer starts/month (late 2024) → **120–130K/month by end of 2026** (~4× in <2 years), "and demand still outpaces it." [secondary — TechTimes]

### 7.3 Component lead times
- TrendForce (Apr 2026): suppliers prioritize AI-server allocation; general-server component lead times **extending to nearly one year** [official]:
  - PMICs: 21–26 weeks → **35–40 weeks** (8-inch BCD capacity allocated to AI PMICs; Samsung S7 8-inch fab shutdown squeezing supply). [official]
  - BMC chips (mature nodes): 11–16 weeks → **21–26 weeks**. [official]
  - PCBs, CPUs: constraints "already evident." [official]
- Global storage inventory (Feb 2026): only **~4 weeks** vs. 8–12 week safety line. [secondary]

### 7.4 Liquid cooling adoption
- **~22% of newly built facilities** using liquid cooling by 2026 (KAD industry overview); cold plate / direct-to-chip ≈ **65% of the liquid-cooling market** in 2026; default design for NVIDIA GB200/GB300. [secondary — KAD](https://www.kad8.com/server/data-center-liquid-cooling-for-ai-workloads-2026/)
- APAC liquid cooling market: **$1.7B (2025) → $2.1B (2026) → $17.5B (2036)**, 23.6% CAGR; direct-to-chip ~55% share in APAC 2026. [secondary — DataNext Research]
- AI datacenter liquid cooling (global): **$6.6B in 2026 → $38.4B by 2033**, 28.7% CAGR (Market Minds Advisory, Feb 2026). [secondary]
- AI datacenter liquid cooling: **$17.83B by 2036** at 16.9% CAGR (Future Market Insights); direct-to-chip ≈ **47% of revenue in 2026**; hyperscale AI datacenters ≈ **55% of deployment demand by 2026**. [secondary]
- Immersion: single-phase at 120–150 MW annual deployment rate in 2026–2027; immersion <2% of global DC cooling (by IT load) in 2025 → 12–18% by 2031–2035 (analyst roadmap). [secondary — Energy Solutions Intelligence]
- Investment signal: Supermicro cites liquid-cooled rack-scale systems (DCBBS) as a competitive differentiator with margin benefit; Foxconn/Quanta GB300 NVL72 racks ship with complete liquid cooling. [secondary]

---

## 8. Conflicting estimates — flagged explicitly

| # | Conflict | Sources |
|---|---|---|
| 1 | **2026 AI infrastructure spend**: $497B (+~56%, IDC AI Infrastructure Tracker July 2026) vs. **$487B** (IDC via Data Center Knowledge, Sept 2026). Likely different tracker vintages; both secondhand. | [secondary] vs [secondary] |
| 2 | **2026 AI server shipment growth**: >20% (TrendForce Oct 2025) → >28% (Jan 2026) → ~28% (Apr 2026) → **nearly 31%** (Aug 2026). Use the latest (Aug 2026) as current. | [official] series |
| 3 | **Total 2026 server shipment growth**: 12.8% (TrendForce Jan 2026) vs. **~13%** (Apr 2026, revised down from ~20% on component constraints) vs. 19.2% (DIGITIMES-adjacent compilation, unverified) vs. 12.8% cited again in invest-bud compilation. The 19.2% figure is an outlier — treat as [unverified]. | [official]/[unverified] |
| 4 | **ASIC vs. GPU in 2026**: TrendForce Oct 2025 says ASIC shipments "expected to surpass those of GPUs" by 2026; TrendForce Jan 2026 says GPUs = **69.7% of AI server shipments**. Possible scope mismatch (AI-chip units vs. AI-server shipments) — do not merge. | [official] vs [official] |
| 5 | **Dell FY2026 Q4 AI figures**: Data Center Knowledge: $9B AI server revenue (+342% YoY), $43B backlog; July-2026 compilation: +39% total sales with AI server revenue "expected to double next year." Different framings of the same quarter — the DCK numbers are more specific and consistent with the FY27 trajectory ($74B guide). | [secondary] vs [secondary] |
| 6 | **AI server market $ value 2026**: IDC AI-infrastructure $487–497B vs. syndicated AI-server estimates $157–224B vs. TrendForce revenue-implied growth. These measure different scopes (full AI infrastructure incl. storage/network vs. servers only vs. AI-server systems). Not directly comparable. | various |
| 7 | **Next Platform's IDC-derived unit math** ($140K avg × 840K units ≠ stated $137.35B total) is internally inconsistent; use IDC's official $87.4B/$27.5B revenue splits and treat unit/ASP derivations as estimates. | [independent] |
| 8 | **Foxconn "40% share"**: company claim ("global AI server market") vs. brokerages ("NVIDIA AI server racks") — different denominators; the latter projects 50% in 2027. | [vendor-reported] vs [secondary] |

---

