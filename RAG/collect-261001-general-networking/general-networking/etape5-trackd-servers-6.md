---
id: collect-261001-general-networking/general-networking/etape5-trackd-servers-6
title: "Step 5 — Server Vendors + AI Server Market + Data-Center Networking"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: ["2026-09-22"]
keywords: ["acquisition", "amd", "backlog", "blackwell", "compute", "disclosure", "gpu", "gpus", "inference", "intel", "liquid cooling", "memory"]
source: docs/RAG/collect-261001-general-networking/etape5_trackD_servers.md
source_anchor: ""
source_lines: [456, 519]
sha256: 80a8e1e268b1f30087f05f23527d43609ff33be17fa5efd27f61b1dabd1b8d2c
---

# Step 5 — Server Vendors + AI Server Market + Data-Center Networking

### Dell
- https://www.dell.com/en-us/dt/corporate/newsroom/announcements/detailpage.press-releases~usa~2026~2~dell-technologies-delivers-fourth-quarter-and-full-year-fiscal-2026-results.htm
- https://www.dell.com/en-us/dt/corporate/newsroom/announcements/detailpage.press-releases~usa~2026~05~dell-technologies-delivers-first-quarter-fiscal-2027-financial-results.htm
- https://www.morningstar.com/news/business-wire/20260901574850/dell-technologies-delivers-second-quarter-fiscal-2027-financial-results
- https://WWW.ZACKS.COM/stock/news/2877003/dell-q4-earnings-beat-estimates-revenues-rise-yy-shares-up?art_rec=blog-analyst_blog-up_next-ID01-img-2877003
- https://www.blocksandfiles.com/flash/2026/02/27/ai-server-frenzy-fuels-record-revenues-for-dell/4092727
- https://www.zacks.com/stock/news/2928734/dells-q1-earnings-call-centers-on-ai-supply-and-demand
- https://pulse2.com/dell-technologies-q2-revenue-hits-record-47-billion-as-ai-server-backlog-reaches-95-billion-and-fy27-outlook-rises-to-192-billion/
- https://sharesify.com/dell-q2-fy2027-ai-demand-surges-but-expectations-are-now-sky-high/
- https://techx.pk/dell-q2-fy2027-ai-server-orders-backlog-95-billion/
- https://finance.biggo.com/news/US_DELL_2026-09-08
- https://finance.biggo.com/news/US_DELL_2026-05-28
- https://www.techpowerup.com/336994/dell-technologies-unveils-next-generation-enterprise-ai-solutions-with-nvidia
- https://www.smetechguru.co.za/dell-technologies-unveils-next-generation-enterprise-ai-solutions-with-nvidia/
- https://hothardware.com/news/dell-announces-poweredge-ai-servers
- https://digitalisationworld.com/news/67716/dell-technologies-expands-dell-ai-factory
- https://www.storagenewsletter.com/2024/05/27/dell-technologies-world-expanded-dell-ai-factory-with-nvidia-to-turbocharge-ai-adoption/
- https://d2649.cms.socastsrm.com/2025/02/14/dell-nears-deal-to-sell-5-billion-in-ai-servers-to-xai-bloomberg-news-reports/
- https://www.sdxcentral.com/news/musks-xai-considering-second-data-center-5bn-dell-chip-deal/
- https://www.brandiconimage.com/2025/02/dell-close-to-finalizing-5-billion-deal.html
- https://en.tmtpost.com/post/7462003
- https://www.mitrade.com/insights/news/live-news/article-3-641092-20250215

*End of report — compiled 2026-09-22. Work was read-only; nothing sent externally.*


---

## §2 HPE + Giga Computing + ASUS + MSI + xFusion

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

