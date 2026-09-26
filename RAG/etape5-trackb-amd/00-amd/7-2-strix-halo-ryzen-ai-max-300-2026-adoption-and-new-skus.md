---
id: etape5-trackb-amd/00-amd/7-2-strix-halo-ryzen-ai-max-300-2026-adoption-and-new-skus
title: "7.2 Strix Halo (Ryzen AI Max 300) — 2026 adoption and new SKUs"
domain: step-5-track-b-amd-hardware-instinct-gpus-epyc-ryzen-cpus
role: deep-dive
task: hardware
actors: ["AMD", "China", "Meta", "Moonshot", "Nvidia", "OpenAI", "Oracle", "vLLM"]
dates: []
keywords: ["amd", "benchmark", "gpu", "gpus", "helios", "inference", "intel", "kimi", "latency", "llama", "llama.cpp", "lpddr5x"]
source: docs/RAG/etape5_trackB_amd.md
source_anchor: ""
source_lines: [181, 219]
section: "Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)"
sha256: 6c11ae4be0dbaa020c6988756feb34146b5337d2f3fecd6500fb093fab9f469b
---

# 7.2 Strix Halo (Ryzen AI Max 300) — 2026 adoption and new SKUs

### 7.2 Strix Halo (Ryzen AI Max 300) — 2026 adoption and new SKUs
- Strix Halo (Zen 5, up to 16c, 40-CU RDNA 3.5 Radeon 8060S, XDNA 2 50 TOPS, up to 128 GB unified LPDDR5X with up to 96 GB assignable as VRAM) launched Jan 2025; by 2026 it became the **de facto local-LLM workstation chip** **[secondary — https://www.techradar.com/pro/ryzen-ai-max-395-cpu-could-be-amds-sleeper-hit-against-nvidias-ai-dominance-as-nearly-30-strix-halo-mini-ai-workstation-models-hit-the-market-including-some-rather-funky-ones]**.
- **Adoption (2026):** nearly **30 mini AI workstation models** launched by 2026 — Beelink GTR9 Pro, Seaviv AideaStation R1, HP Z2 Mini G1a, GMKtec EVO-X2, Minisforum MS-S1 MAX, Abee AI Station 395 Max; laptops: Asus ROG Flow Z13, Chinese mobile-workstation designs (Sixunited, Linglong, Tianba). Price positioning **$1,800–$2,800** vs $21,000+ for multi-GPU AI servers **[secondary — TechRadar]**. ⚠️ Big-brand availability remains limited; "real-life tests remain scarce" per TechRadar.
- **New SKUs at CES 2026:** **Ryzen AI Max+ 392** and **Ryzen AI Max+ 388** — detuned CPU core counts from the Max+ 395 but **full 40-CU GPU** retained; the 392 surfaced in Geekbench (~2,917 single-core, ~12–15% behind the 395 in multi-core) **[independent — https://hothardware.com/news/amd-ces-2026?ref=thetechstreetnow.com; https://www.notebookcheck.net/New-AMD-Strix-Halo-Ryzen-AI-Max-392-stars-in-early-benchmark-after-CES-2026-debut.1204390.0.html]**.
- **Local-AI positioning:** AMD claims HP Z2 Mini G1a advantages of **1.5×/1.7× tokens-per-second-per-dollar** in LM Studio on GPT-OSS 20B/120B vs NVIDIA DGX Spark **[vendor-reported — https://hothardware.com/news/amd-ces-2026?ref=thetechstreetnow.com]**; community reports ~47–53 tok/s on 120B models vs DGX Spark's 56 tok/s at lower price **[secondary — https://github.com/getnyrex/strix-halo-guide/blob/HEAD/RESEARCH.md]**.
- **Framework Desktop** clusters: 4-node Strix Halo clusters running **1T-parameter Kimi K2.5** via distributed llama.cpp reported by community researchers **[secondary — https://github.com/getnyrex/strix-halo-guide/blob/HEAD/RESEARCH.md]**.
- **Ryzen AI Embedded (Jan 8, 2026):** new embedded portfolio — **P100** (Zen 5, up to 12c, RDNA 3.5, XDNA 2, 50 TOPS, 15–54 W, −40°C to +105°C) for in-vehicle/industrial; **X100** (Strix Halo silicon, up to 16c) for "physical AI"/autonomous systems **[official via press — https://videocardz.com/newz/amd-introduces-ryzen-ai-embedded-p100-strix-point-krackan-and-x100-strix-halo-series]**.
- Rumored roadmap: "Gorgon Halo" (Ryzen AI Max 400, same arch, higher clocks) Q4 2026; "Medusa Halo" with LPDDR6 (~80% more memory bandwidth) **[unverified — https://github.com/getnyrex/strix-halo-guide/blob/HEAD/RESEARCH.md]**.

---

## 8. AMD AI PARTNERSHIPS AND STRATEGIC DEALS — 2026

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

