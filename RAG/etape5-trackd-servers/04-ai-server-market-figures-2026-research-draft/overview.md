---
id: etape5-trackd-servers/04-ai-server-market-figures-2026-research-draft/overview
title: "AI Server Market Figures (2026) — Research Draft"
domain: ai-server-market-figures-2026-research-draft
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Alibaba", "Baidu", "ByteDance", "China", "Google", "Meta", "Microsoft", "Nvidia", "Oracle", "United States"]
dates: ["2025-10-30", "2026-01-20", "2026-04-15", "2026-08-03", "2026-09", "2026-09-22"]
keywords: ["research", "amd", "asic", "aws", "blackwell", "capex", "gpu", "gpus", "hbm", "helios", "inference", "nvidia"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [1008, 1053]
section: "AI Server Market Figures (2026) — Research Draft"
sha256: 52701d08d4290b1735944aadf100e74346e6f0d6f26d6c31b694f77711445a22
---

# AI Server Market Figures (2026) — Research Draft
**As of: September 22, 2026** | Language: English | Status: draft (étape 5)
**Coverage:** AI server market size, vendor shares, ODM share, GPU/general split, regional breakdowns, analyst reports, supply-chain notes.

**Provenance legend** (applied to every claim):
- `[official]` — vendor press releases / earnings disclosures / official analyst press releases
- `[vendor-reported]` — figures from vendor executives or investor materials
- `[independent]` — reputable third-party press/research interpretation of tracker data
- `[secondary]` — syndicated research firms / aggregators / trade press citing others
- `[unverified]` — blogs, LinkedIn posts, AI-generated summaries; treat with caution

> All figures are cited to their source as found in September 2026. Absolute dollar/unit totals differ sharply between sources because of **scope differences** (AI servers only vs. AI infrastructure vs. total server market; revenue vs. shipments; vendor revenue vs. end-user spend). See "Conflicting estimates" section.

---

## 1. Market size and growth — headline figures

### 1.1 TrendForce — AI server shipments and revenue
TrendForce's public AI-server forecasts evolved through 2025–2026 (each later release supersedes the earlier one):

- **October 30, 2025** press release [official]: [TrendForce](https://www.trendforce.com/presscenter/news/20251030-12762.html)
  - 2025: AI server shipment growth forecast "about 24%" (slightly lowered due to restrictions on NVIDIA H20 shipments to China and minor delays in GB300/B300 platform launches). [official]
  - 2025: AI server **revenue +~48% YoY**, driven by Blackwell platforms (GB200, GB300 full-rack solutions). [official]
  - 2026: AI server shipments to grow **"more than 20%"**; AI servers to account for **17% of overall server shipments**. [official]
  - 2026: AI server revenue to grow **"more than 30%"** and account for **74% of the overall server market value**. [official]
  - 2025: NVIDIA to maintain **~70% of the AI chip market share**. [official]
  - By 2026, ASIC shipments "expected to surpass those of GPUs," gradually reducing NVIDIA's share. [official] *(⚠️ tension with TrendForce's Jan-2026 figure that GPUs = 69.7% of AI server shipments — flagged in Conflicts.)*
  - HBM consumption expected to surge **70% in 2026**. [official]

- **January 20, 2026** press release [official]: [TrendForce](https://www.trendforce.com/presscenter/news/20260120-12887.html)
  - 2026: global AI server shipments forecast **>28% YoY**, outpacing overall server market growth of 12.8%. [official]
  - 2026: total global server shipments +12.8% YoY (AI inference boosting general-purpose server demand). [official]
  - In 2026, **GPUs remain the leading category at 69.7% of AI server shipments**; NVIDIA GB300-based systems expected to drive most shipments; VR200-based platforms to ramp gradually in 2H 2026. [official]
  - **ASIC-based AI servers: ~27.8% of shipments in 2026**, "the highest since 2023," with ASIC shipment growth outpacing GPU-based systems; Google investing more in own ASICs than most CSP competitors, positioned as a leading market player. [official]

- **April 15, 2026** press release (server industry, component constraints) [official]: [TrendForce](https://www.trendforce.com/presscenter/news/20260415-13013.html)
  - 2026 AI server shipment growth held at **~28% YoY**, ASIC-based systems expected to outpace GPU-based systems. [official]
  - ASIC share slightly revised down from nearly 28% to **~27%** (validation/tuning delays at players such as Meta and AWS). [official]
  - **Full-year 2026 total server shipment growth cut to ~13% YoY** (previously approaching 20%), because suppliers prioritize capacity for higher-margin AI server products and general-server component lead times lengthened. [official]

- **August 3, 2026** press release [official]: [TrendForce](https://www.trendforce.com/presscenter/news/20260803-13161.html)
  - 2026 AI server shipment forecast **revised upward from 28% to nearly 31% YoY**. [official]
  - Combined CapEx of the world's nine largest CSPs (Google, Amazon, Meta, Microsoft, Oracle, ByteDance, Tencent, Alibaba, Baidu) to rise **~90% YoY in 2026 to >US$886.7 billion**, with the five North American hyperscalers accounting for nearly 90%. [official]
  - 2026: major CSPs forecast to account for **76% of annual high-end AI server shipments**. [official] *(per TrendForce landing page: "By 2026, major CSPs are forecasted to account for 76% of the annual total shipments of high-end AI servers.")* [official] — https://www.trendforce.com/landingPage/ai_server
  - CSP AI-chip strategies 2026: Google expanding in-house TPUs through 2026–2027; AWS deploying NVIDIA GB300 as primary GPU platform in 2026 while expanding internal ASICs; Meta on NVIDIA GB/VR and AMD Helios rack-scale systems in 2026, accelerating proprietary AI ASICs in 2027. [official]

