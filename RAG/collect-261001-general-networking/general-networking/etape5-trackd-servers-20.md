---
id: collect-261001-general-networking/general-networking/etape5-trackd-servers-20
title: "Step 5 — Server Vendors + AI Server Market + Data-Center Networking"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Google", "Nvidia", "OpenAI", "Samsung", "TSMC"]
dates: ["2026-05", "2026-28-07"]
keywords: ["amd", "asic", "backlog", "capex", "compute", "cowos", "dram", "gpu", "hbm", "hbm3", "hbm4", "memory"]
source: docs/RAG/collect-261001-general-networking/etape5_trackD_servers.md
source_anchor: ""
source_lines: [1277, 1327]
sha256: 0e19331fc78ff893671f775cc5e1ab79d3b097186b2bd7e24f81d5dfba350aa7
---

# Step 5 — Server Vendors + AI Server Market + Data-Center Networking

| Date | Publisher | Title / content | Link | Provenance |
|---|---|---|---|---|
| Jan 20, 2026 | TrendForce | Global AI server shipments forecast >28% YoY in 2026, rising ASIC share | [presscenter/news/20260120-12887](https://www.trendforce.com/presscenter/news/20260120-12887.html) | [official] |
| Apr 15, 2026 | TrendForce | 2026 server shipments +13% YoY on component delays; AI servers ~28% | [presscenter/news/20260415-13013](https://www.trendforce.com/presscenter/news/20260415-13013.html) | [official] |
| Aug 3, 2026 | TrendForce | AI server shipments forecast raised to nearly 31% YoY; CSP CapEx +90% | [presscenter/news/20260803-13161](https://www.trendforce.com/presscenter/news/20260803-13161.html) | [official] |
| Oct 30, 2025 | TrendForce | >20% AI server shipment growth in 2026; revenue +30%+, 74% of server value | [presscenter/news/20251030-12762](https://www.trendforce.com/presscenter/news/20251030-12762.html) | [official] |
| Jun 15, 2026 | IDC | Worldwide Server Market Revenue Surpasses $122B in Q1 2026 | [idc.com](https://www.idc.com/resource-center/press-releases/1q26-server-tracker/) | [official] |
| Jul 21, 2026 | IDC | Worldwide Quarterly AI Infrastructure Tracker (2026 forecast $497B) — via summary | [quantumrun.com](https://www.quantumrun.com/consulting/global-ai-market-size/) | [secondary] |
| ~Sept 2026 | IDC | Q2 2026 server tracker: $166.3B record — via StorageReview | [storagereview.com](https://www.storagereview.com/news/idc-external-storage-tracker-q2-2026-10-3-billion-as-server-market-tops-166-billion) | [independent] |
| Feb 2026 | Gartner | IT spending $6.15T; data center systems $653.4B (+31.7%); servers +36.9% | [hpcwire.com](https://www.hpcwire.com/aiwire/2026/02/04/gartner-forecasts-worldwide-it-spending-to-grow-10-8-in-2026-totaling-6-15t/) | [secondary] |
| Jul 2026 | Gartner | IT spending revised to $6.37T; data center systems $822B (+62.5%) | [techbriefly.com](https://techbriefly.com/2026/07/28/2026-global-it-spending-forecast-6-37-trillion/) | [secondary] |
| May 2026 | Gartner | Worldwide AI spending $2.59T in 2026 (+47%) | via [techbriefly.com](https://techbriefly.com/2026/07/28/2026-global-it-spending-forecast-6-37-trillion/) | [secondary] |
| Jul 30, 2026 | Omdia | 2026 semiconductor forecast raised to +94.1% YoY on AI memory | [financialcontent.com](https://markets.financialcontent.com/fatpitch.valueinvestingnews/article/bizwire-2026-7-30-omdia-ai-demand-drives-941-surge-in-semiconductor-forecast-for-2026) | [secondary] |
| Oct 2025 | DIGITIMES Research | Global server shipment forecast 2026 and beyond (CAGR ~5% 2025–2030) | [digitimes.com](https://www.digitimes.com/reports/item.php?id=20240227RS400) (summary) | [secondary] |
| Feb 6, 2026 | DIGITIMES | AI spillover lifts 2026 general-purpose server shipments 10% (paywalled) | [digitimes.com](https://www.digitimes.com/news/a20260206PD221/demand-shipments-2026-general-purpose-server-growth.html) | [secondary] |
| Sept 21, 2026 | Next Platform | "Accelerated Server Sales Are Themselves Accelerating" (IDC Q2 analysis) | [nextplatform.com](https://www.nextplatform.com/compute/2026/09/21/accelerated-server-sales-are-themselves-accelerating/5297993) | [independent] |
| Sept 2026 | Data Center Knowledge | AI Server Market Update: Vendors Shift from Silicon to Services | [datacenterknowledge.com](https://www.datacenterknowledge.com/servers/ai-server-market-update-vendors-shift-from-silicon-to-services) | [secondary] |
| Jul 2026 | GPUSmith | NVIDIA AI Server OEM Comparison: Supermicro vs Dell vs HPE (PDF) | [gpusmith.com](http://gpusmith.com/articles/en/pdfs/nvidia-ai-server-oem-comparison.pdf) | [secondary] |

---

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

