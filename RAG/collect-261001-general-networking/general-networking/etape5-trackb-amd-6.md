---
id: collect-261001-general-networking/general-networking/etape5-trackb-amd-6
title: "Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Google", "Meta", "Microsoft", "Nvidia", "OpenAI", "Oracle", "TSMC", "vLLM"]
dates: []
keywords: ["amd", "gpu", "gpus", "2nm", "benchmarks", "compute", "copilot", "fp4", "hbm3", "hbm4", "helios", "inference"]
source: docs/RAG/collect-261001-general-networking/etape5_trackB_amd.md
source_anchor: ""
source_lines: [194, 255]
sha256: 552ffd99552220c2a8d7187c8af61073b304c0346e14604fb06eb62a4c8ab287
---

# Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)

### 8.1 Meta — 6 GW multi-generation deal (announced ~Feb 25, 2026)
- AMD and Meta signed a multi-year agreement for up to **6 gigawatts** of AI infrastructure across multiple Instinct GPU generations **[independent — https://hostingjournalist.com/news/amd-lands-meta-as-major-ai-chip-customer]**.
- First phase: a **custom MI450-architecture GPU** variant (optimized for Meta's software stack and low-latency inference) + 6th Gen EPYC "Venice" CPUs + Helios rack architecture, on ROCm. **Shipments for the first 1 GW begin H2 2026**; remaining 5 GW roll out in tranches **through 2031** **[independent/secondary — https://hostingjournalist.com/news/amd-lands-meta-as-major-ai-chip-customer; https://markets.financialcontent.com/ms.intelvalue/article/marketminute-2026-2-25-amd-secures-historic-6-gigawatt-ai-infrastructure-deal-with-meta-platforms]**.
- Financial structure: AMD issued Meta **performance-based warrants for up to 160 million AMD shares** (~10% of the company) vesting in stages per deployed gigawatt; final tranche vests only if AMD's share price hits **$600** **[secondary — https://markets.financialcontent.com/ms.intelvalue/article/marketminute-2026-2-25-amd-secures-historic-6-gigawatt-ai-infrastructure-deal-with-meta-platforms]**.
- Deal value: Wolfe Research analysts estimated **$15–20B per gigawatt**, implying **$90–120B** in potential AMD revenue over the full 6 GW (Meta was already an AMD customer, so not all incremental) **[secondary — https://www.neowin.net/news/amd-lands-massive-100-billion-gpu-deal-with-meta-for-ai-data-centers/]**.
- Co-development: Helios rack architecture **co-designed by AMD and Meta through OCP**; Forrest Norrod (AMD DCGM EVP/GM): "Helios reflects the spirit of open collaboration that defines the OCP community" **[secondary — https://techsabado.com/2025/10/28/tech-news-amd-ibm-zyphra-forge-new-path-for-ai-infrastructure/; https://hostingjournalist.com/news/amd-lands-meta-as-major-ai-chip-customer]**.
- ⚠️ The same secondary sources note Meta signed a hardware deal with NVIDIA a week prior — the AMD deal is part of a multi-vendor diversification strategy **[secondary — https://www.neowin.net/news/amd-lands-massive-100-billion-gpu-deal-with-meta-for-ai-data-centers/]**.

### 8.2 OpenAI — 6 GW deal (announced Oct 2025; executing in 2026)
- AMD and OpenAI reached a **6 GW multi-generation agreement** (announced Oct 2025; OpenAI statement **Oct 6**): deployments start with the **Instinct MI450 series + rack-scale (Helios) solutions**, extending to future generations; **first 1 GW begins H2 2026**, tied to the Stargate buildout **[secondary — https://en.tmtpost.com/post/7721864]**.
- Same warrant structure as Meta: up to **160M AMD shares at $0.01**, vesting on deployment milestones (first tranche at 1 GW), share-price targets, and OpenAI technical/commercial milestones **[secondary — same]**.
- Collaboration history: "began with the MI300X and continued with the MI350X series" **[secondary — same]**.

### 8.3 Oracle (Oct 2025 announcement; deployment from Q3 2026)
- 50,000 AMD GPUs on OCI starting calendar Q3 2026, expansion 2027+: MI300X first, then MI355X GA; earlier 30,000-unit MI355X order (Mar 2025) **[secondary — https://www.mitrade.com/insights/news/live-news/article-3-1194409-20251015; https://www.nasdaq.com/articles/oracle-recently-delivered-incredible-news-advanced-micro-devices-amd-stock-investors]**.

### 8.4 Other 2026 ecosystem items
- **Red Hat:** Red Hat AI 3 certified on Instinct; joint vLLM development; AMD GPU Operator on OpenShift; RHEL AI bare-metal on MI300X **[secondary — https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/amd-deep-dive.md]**.
- **Cirrascale:** first neocloud to commit to Helios + MI455X/MI430X (Jul 14, 2026) **[official — https://www.businesswire.com/news/home/20260714481734/en/Cirrascale-Advances-Open-AI-Infrastructure-with-AMD-Helios-Rackscale-Solution-and-AMD-Instinct-MI400-Series-GPUs]**.
- **AMD AI DevDay 2026** (San Francisco) and **Advancing AI 2026** (Jul 22–23, San Francisco) as the year's developer/customer showcases **[secondary — https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/amd-deep-dive.md]**.
- **IBM/Zyphra** (Oct 2025, context): MI300X training cluster on IBM Cloud + Pensando Pollara 400 NICs **[secondary — https://techsabado.com/2025/10/28/tech-news-amd-ibm-zyphra-forge-new-path-for-ai-infrastructure/]**.

---

## 9. SPEC SUMMARY TABLE

| Product | Architecture / Process | Memory | Peak compute (dense) | Power | Status Sep 2026 |
|---|---|---|---|---|---|
| MI300X | CDNA 3 | 192 GB HBM3, 5.3 TB/s | ~1.3 PFLOPS FP16 | 750 W | Mature; ~$6/hr cloud; full-capacity pricing signals |
| MI325X | CDNA 3 | 256 GB HBM3E, 6 TB/s | 2.61 PFLOPS FP16 | 1,000 W | Shipping |
| MI350X | CDNA 4 / TSMC N3P | 288 GB HBM3E, 8 TB/s | 18.4 PFLOPS FP4 | 1,000 W | GA H2 2025; $8.60/hr OCI |
| MI355X | CDNA 4 / TSMC N3P | 288 GB HBM3E, 8 TB/s | 20.1 PFLOPS FP4 | 1,400 W | Flagship shipping; Oracle 30k order |
| MI455X | CDNA 5 / TSMC 2nm | 432 GB HBM4, 19.6 TB/s | 40 PFLOPS FP4 | TBD (liquid) | Volume H2 2026 |
| EPYC 9965 (Turin) | Zen 5c | DDR5-6400 12ch | 192c/384t | 500 W | Adopted (Google, OVHcloud) |
| EPYC Venice (9006) | Zen 6 / TSMC 2nm | DDR5-8000 16ch | 256c/512t | TBD | Volume production from Jul 2026 |
| Ryzen AI 9 HX 475 | Zen 5/5c + RDNA 3.5 + XDNA 2 (60 TOPS) | LPDDR5X-8533 | 12c/24t | 28 W | Announced CES 2026 (Gorgon Point) |
| Ryzen AI Max+ 395 | Zen 5 + RDNA 3.5 40CU + XDNA 2 (50 TOPS) | 128 GB unified | 16c/32t | 45–120 W | Shipping; ~30 mini-PC/laptop designs |

---

## 10. OPEN VERIFICATION ITEMS

1. **MI350 vs B200 head-to-heads** — AMD's 10–30% claims are vendor-reported; independent third-party benchmarks at scale remain thin as of Sep 22, 2026.
2. **MI300X/MI350 list prices** — AMD does not publish; all figures are cloud $/hr or community estimates. Older "~$15,000" claims have no verified 2026 source.
3. **MI400/MI455X final specs and exact volume timing** — H2 2026 volume is AMD's stated schedule; actual ramp and TDP figures unconfirmed.
4. **MI430X** — sovereign-AI/HPC variant; specs not published by AMD as of Sep 2026.
5. **"MI450" vs "MI455X" naming** — sources mix the two; MI455X is the flagship rack part; Meta's "custom MI450" variant details unconfirmed.
6. **MI500 (2027)** — announced at Analyst Day; no specs.
7. **Helios bandwidth figures** — 260 TBps (Register) vs 43 TB/s scale-out (Cirrascale/BusinessWire) vs per-GPU 1.8 TB/s/dir (STH): marketing maxima; apples-to-oranges.
8. **Oracle 30,000 MI355X order** — from a 2025 earnings call via secondary reporting; fulfillment status unconfirmed.
9. **Venice pricing** — no MSRP published as of Sep 2026.
10. **Turin 2026 pricing trends** — no 2026 SKU/pricing updates located in this research.
11. **"First Copilot+ desktop CPU"** — thinly sourced; AMD's Gorgon Point desktop APUs' certification details unconfirmed.
12. **Gorgon Halo / Medusa Halo** — rumor-grade (community research repo).
13. **Venice CCD core type** — resolved post-Advancing AI 2026 (Zen 6 / Zen 6c split), but per-SKU core maps incomplete.
14. **Wolfe Research $90–120B Meta-deal estimate** — analyst projection, not committed revenue.
15. **Spheron $3.59/hr MI300X vs $2.65/hr H100** — live marketplace snapshot, Sep 22, 2026; supply-driven, moves constantly.

---

## 11. COLLECTION METADATA

