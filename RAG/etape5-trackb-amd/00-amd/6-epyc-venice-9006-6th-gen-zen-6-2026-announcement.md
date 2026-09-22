---
id: etape5-trackb-amd/00-amd/6-epyc-venice-9006-6th-gen-zen-6-2026-announcement
title: "6. EPYC VENICE 9006 (6th Gen, Zen 6) — 2026 announcement"
domain: step-5-track-b-amd-hardware-instinct-gpus-epyc-ryzen-cpus
role: deep-dive
task: hardware
actors: ["AMD", "China", "Google", "Intel", "Meta", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "Oracle", "TSMC", "vLLM"]
dates: ["2026-07", "2026-09-22"]
keywords: ["2nm", "agentic", "amd", "benchmark", "benchmarks", "chiplet", "compute", "copilot", "datacenter", "fp4", "gpu", "gpus"]
source: docs/RAG/etape5_trackB_amd.md
source_anchor: ""
source_lines: [152, 260]
section: "Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)"
sha256: 1a863a776814b9b51680e6ce445d938d0c81459290123331ce06ba887629374c
---

# 6. EPYC VENICE 9006 (6th Gen, Zen 6) — 2026 announcement

## 6. EPYC VENICE 9006 (6th Gen, Zen 6) — 2026 announcement

### 6.1 Timeline
- **CES 2026 (Jan 6, 2026):** Lisa Su held up a bare (no-IHS) Venice sample in her keynote, revealing TSMC **InFO_oS** advanced chiplet packaging — 2 I/O dies surrounded by 2 rows of 4 CCDs **[independent — https://abit.ee/en/hard/processors/amd-epyc-venice-processor-server-zen-6-ces-2026-chiplet-infoos-tsmc-256-cores-en; https://www.guru3d.com/story/amd-unveils-256core-epyc-venice-for-helios-mi455x-ai-racks-platform/]**.
- **Advancing AI 2026 (Jul 22–23, 2026), San Francisco:** AMD formally unveiled the 6th Gen EPYC "Venice" (EPYC 9006 series); announced as **entering volume production on TSMC 2nm** — the first HPC/server chip on 2nm **[independent — https://www.martincid.com/technology-sv/amd-epyc-venice-256-cores-2nm/; https://finance.biggo.com/news/1fa76c25-95b5-4fd8-a771-5dd63c39e1a2]**. AMD CTO Mark Papermaster confirmed the July 2026 launch window, starting with server/datacenter EPYC **[independent — https://www.club386.com/amd-confirms-zen-6-launch-date/]**.
- EPYC roadmap through 2030 also published: **Venice (Zen 6) 2026 → Florence (Zen 7) 2028 → Ravenna (Zen 8) 2030** **[secondary — https://www.kad8.com/hardware/amd-epyc-roadmap-venice-florence-and-ravenna-through-2030/]**.

### 6.2 Specs and variants
- **Up to 256 cores / 512 threads** per socket: 8 CCDs × 32 Zen 6c cores (density-optimized), or up to **96 standard Zen 6 cores** (192 threads) for single-thread leadership **[official via press — https://finance.biggo.com/news/1fa76c25-95b5-4fd8-a771-5dd63c39e1a2; https://www.martincid.com/technology-sv/amd-epyc-venice-256-cores-2nm/]**. ⚠️ Whether the 32-core CCDs are full Zen 6 or dense Zen 6c was initially ambiguous in CES coverage; the Advancing AI 2026 reveal clarified the Zen 6 / Zen 6c split.
- **3D V-Cache up to 1,152 MB L3** on stacked variants **[official via press — https://finance.biggo.com/news/1fa76c25-95b5-4fd8-a771-5dd63c39e1a2]**.
- New **SP7 socket**; **16 channels DDR5-8000** (~1.6 TB/s aggregate, ~3× Turin's ~614 GB/s); **MRDIMM** support reaching DDR5-12800-equivalent bandwidth **[official via press — https://www.martincid.com/technology-sv/amd-epyc-venice-256-cores-2nm/; https://finance.biggo.com/news/1fa76c25-95b5-4fd8-a771-5dd63c39e1a2]**.
- **PCIe Gen 6 (64 Gbps)**, 5th-gen Infinity Fabric **[official via press — https://www.martincid.com/technology-sv/amd-epyc-venice-256-cores-2nm/]**.
- TSMC 2nm GAA nanosheet: 10–15% higher perf at same power or 25–30% lower power at same perf, +15% density (TSMC figures) **[independent — https://www.martincid.com/technology-sv/amd-epyc-venice-256-cores-2nm/]**.
- Product lines (4): **EPYC 9006 SP7** (flagship 256c "agentic AI" part, 5 GHz boost, 128 PCIe Gen6 lanes); **EPYC 9006X SP7** (up to 96c, 5.15 GHz, 3× L3 per core, HPC/simulation); **EPYC 9006 SP8** (8–128c, edge/power-constrained); **EPYC 9006 LP "Verano"** (LPDDR5X, 24 channels, SOCAMM2 form factor, 112 Gbps CPU-to-GPU bandwidth for rack-scale AI chassis) **[official via press — https://www.martincid.com/technology-sv/amd-epyc-venice-256-cores-2nm/; https://www.tweaktown.com/news/112793/amd-announces-6th-gen-amd-epyc-venice-cpus-up-to-256-cores-and-512-threads-and-built-for-ai/index.html]**.
- AMD vendor claims: **+70% CPU throughput vs Turin Zen 5**; **2.2× AI inference throughput vs NVIDIA's "Vera" platform** at the 256-core tier; Venice **>3× faster than NVIDIA Vera** (per TweakTown headline) **[vendor-reported — https://www.martincid.com/technology-sv/amd-epyc-venice-256-cores-2nm/; https://www.tweaktown.com/news/112793/amd-announces-6th-gen-amd-epyc-venice-cpus-up-to-256-cores-and-512-threads-and-built-for-ai/index.html]**.
- Manufacturing: initial production at **TSMC Taiwan**, later also **TSMC Arizona** **[independent — https://www.club386.com/amd-confirms-zen-6-launch-date/]**.

---

## 7. RYZEN CLIENT AI — 2026 releases

### 7.1 Gorgon Point / Ryzen AI 400 series (announced CES 2026, Jan 6, 2026)
- AMD announced the **Ryzen AI 400 series mobile processors** (codename **Gorgon Point**) — a refresh of Strix Point (Ryzen AI 300): same **Zen 5 / Zen 5c CPU + RDNA 3.5 GPU + XDNA 2 NPU** formula, 4nm-class, with higher clocks and an upgraded NPU **[independent — https://www.techpowerup.com/338398/amd-ryzen-ai-9-hx-475-470-gorgon-point-apus-surface-in-shipping-manifests; https://abit.ee/en/hard/processors/amd-ryzen-ai-400-gorgon-point-processor-apu-zen-5-copilot-strix-halo-ces-2026-npu-en; https://wccftech.com/amd-confirms-zen-6-medusa-cpus-2027-next-gen-gaming-gpu-improved-ai-raytracing/]**.
- Flagship **Ryzen AI 9 HX 475**: 12c/24t (4× Zen 5 + 8× Zen 5c), up to **5.2 GHz** (vs 5.1 on Strix Point), 16-CU Radeon 890M, **L3 up to 36 MB** (vs 34), **NPU up to 60 TOPS** (vs 50 on Strix Point; 55+ TOPS in other configs), LPDDR5X-8533 support (vs 8000), 28 W default TDP **[official via press / secondary — https://abit.ee/en/hard/processors/amd-ryzen-ai-400-gorgon-point-processor-apu-zen-5-copilot-strix-halo-ces-2026-npu-en; https://www.tweaktown.com/news/104174/amds-next-gen-gorgon-point-apus-leaked-zen-5-rdna-3-refresh-expected-in-2026/index.html; https://www.techpowerup.com/338398/amd-ryzen-ai-9-hx-475-470-gorgon-point-apus-surface-in-shipping-manifests]**.
- Also surfaced in shipping manifests: **Ryzen AI 9 HX 470** (5.2 GHz), plus new **Ryzen AI 7 450** and **Ryzen 5/3** entry models **[secondary — https://www.techpowerup.com/338398/amd-ryzen-ai-9-hx-475-470-gorgon-point-apus-surface-in-shipping-manifests]**.
- **Desktop + enterprise expansion (MWC 2026, Mar 2):** **Ryzen AI PRO 400 series** (enterprise-hardened), and **Ryzen AI 7 450G / 5 440G desktop APUs** — positioned by one source as the first **Copilot+-certified desktop processors** (Gorgon Point silicon, XDNA 2 NPU ≥40 TOPS) **[secondary — https://markets.financialcontent.com/pentictonherald/article/marketminute-2026-3-6-amd-solidifies-ai-everywhere-strategy-with-massive-ryzen-ai-400-expansion-into-desktop-and-enterprise-markets; https://abit.ee/en/hard/processors/amd-ryzen-ai-400-gorgon-point-processor-apu-zen-5-copilot-strix-halo-ces-2026-npu-en]**. ⚠️ HP, Lenovo, Dell OEM design wins reported; the "first Copilot+ desktop" claim is thinly sourced — treat cautiously.
- AMD's claimed positioning: flagship Ryzen AI 9 HX 470 shows **1.3× multitasking, 1.7× content creation, +10% gaming** vs Intel Core Ultra 9 288V at comparable TDP (28 W vs 30 W); up to 24h laptop battery life **[vendor-reported — https://abit.ee/en/hard/processors/amd-ryzen-ai-400-gorgon-point-processor-apu-zen-5-copilot-strix-halo-ces-2026-npu-en]**.
- **Roadmap context:** Gorgon Point is a Zen 5 refresh; true next-gen **Medusa Point (Zen 6, N3-class)** arrives **2027** with >10× AI performance claims, plus **Medusa Baby** for mainstream (H2 2027); Strix Halo remains in-market through at least end-2027 **[secondary — https://www.tomshardware.com/pc-components/cpus/amd-mobile-cpu-roadmap-leak-claims-zen-6-arrives-in-2027; https://wccftech.com/amd-confirms-zen-6-medusa-cpus-2027-next-gen-gaming-gpu-improved-ai-raytracing/]**.

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

- **Research date:** September 22, 2026.
- **Search passes:** 8 web-search batches covering MI300/MI350 specs & benchmarks, MI400/Helios architecture, Oracle/Meta/OpenAI partnerships, EPYC Venice announcement, Ryzen Gorgon Point & Strix Halo, MI300X cloud pricing.
- **Source mix:** AMD announcements via press coverage (official-by-report), The Register, ServeTheHome, Tom's Hardware, TechPowerUp, CRN, TweakTown, Notebookcheck, HotHardware, BusinessWire (Cirrascale), plus secondary aggregators and community research repos (flagged as such).
- **Cross-references:** ROCm stack → `etape4_trackC_cuda_rocm_pytorch.md` (Step 4 Track C); NVIDIA comparison data → Step 3 Track D.
- **Nothing was sent externally; no live-browser visits; read-only web research + local file write.**
