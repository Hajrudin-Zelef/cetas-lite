---
id: etape5-trackc-huawei-intel/00-huawei-intel/2-chinese-ai-chips-cambricon-moore-threads-biren-metax
title: "2. CHINESE AI CHIPS — CAMBRICON, MOORE THREADS, BIREN, METAX"
domain: step-5-track-c-huawei-chinese-hardware-intel-cpus-feb-1-sep-
role: deep-dive
task: hardware
actors: ["AMD", "Alibaba", "Baidu", "ByteDance", "China", "DeepSeek", "Huawei", "IREN", "Nvidia"]
dates: ["2025-06", "2025-09", "2025-12-05", "2025-12-17", "2026-01-02", "2026-03", "2026-06"]
keywords: ["amd", "compute", "decode", "deepseek", "energy", "fp4", "fp8", "gpu", "hbm", "inference", "ipo", "lpddr"]
source: docs/RAG/etape5_trackC_huawei_intel.md
source_anchor: ""
source_lines: [57, 103]
section: "Step 5 — Track C: Huawei + Chinese Hardware + Intel CPUs (Feb 1 → Sep 22, 2026)"
sha256: 66b85ecf225a09eb8f8fe2f11faf694b66fdc4e6a2b79e1634d5f9ec3516fd5d
---

# 2. CHINESE AI CHIPS — CAMBRICON, MOORE THREADS, BIREN, METAX

## 2. CHINESE AI CHIPS — CAMBRICON, MOORE THREADS, BIREN, METAX

### 2.1 Cambricon — the production leader

- **2026 production plan (Bloomberg, Dec 2025):** Cambricon plans to deliver **~500,000 AI accelerators in 2026** — more than tripling the ~142,000 units Goldman Sachs estimates for 2025 — including up to **300,000 units of flagship Siyuan 590 and 690** **[independent — Bloomberg via TrendForce/TechRadar/i3investor]**.
- **Process:** SMIC **N+2 7nm**; yields on the 590/690 reportedly only **~20%** — four of five dies unusable — and SMIC advanced-node capacity is the binding constraint **[independent — Bloomberg via TrendForce]**.
- **Siyuan 590:** in volume production since **Q3 2024**; ~80% of NVIDIA A100 performance; gains over Siyuan 290 estimated at **15–20%** via more advanced HBM and added compute units (same N+2 node) **[independent — TrendForce; secondary — ainvest]**.
- **Siyuan 690:** still in testing; large-scale mass production may slip to **H2 2026** **[independent — TrendForce]**.
- **Commercial:** Cambricon held just ~1% of China's AI chip market in 2024 (behind NVIDIA, Huawei, AMD); revenue surged **14× in the September 2025 quarter**; market value up 9× since 2021; **RMB 4.6B revenue in the first three quarters of 2025** — multiples of other Chinese AI-chip startups **[independent — TrendForce; secondary — ainvest]**.
- **Customers:** ByteDance accounts for **>50% of current orders**; Alibaba in talks as a future customer **[independent — Bloomberg relay]**.
- ⚠️ **Cambricon denied the reports:** in a WeChat statement it said circulating media information on its products, customers and production forecasts is "all inaccurate," without specifying which outlets **[independent — communicationstoday citing Bloomberg]** — treat the 500K figure as **[secondary]** pending confirmation.

### 2.2 Moore Threads — first domestic-GPU IPO, now profitable (narrowly)

- **IPO:** listed on Shanghai STAR Market **December 5, 2025** as "China's first domestic GPU stock," raising **8 billion yuan (~$1.2B)**; shares rose **>500% in the first 8 trading sessions**, 5× on debut day **[secondary — CoinMarketCap; BigGo Finance]**.
- **2025 results:** revenue **RMB 1.45–1.52B** (+231–247% YoY); net loss narrowed to **RMB 950M–1.1B** from 1.6B a year earlier; SCMP, Jan 22, 2026 **[secondary — indexbox citing SCMP]**.
- **Q1 2026 (TrendForce, Apr 28, 2026, citing EE Times China):** revenue **RMB 738M (+155.35% YoY)**; first quarterly net profit **RMB 29.36M** attributable to shareholders (swinging from a loss) — though **ex-non-recurring profit was still a RMB 54M loss** (narrowed 60% YoY) **[independent — TrendForce]**.
- **Flagship MTT S5000** (4th-gen **Pinghu** architecture, mass production 2025): up to **1,000 TFLOPS dense AI compute** per card; full precision **FP8–FP64**; native hardware acceleration for Transformer and MoE architectures **[independent — TrendForce; secondary — TMTPost]**; designed to rival H100 for LLM training **[secondary — TMTPost; DIGITIMES]**.
- **March 2026:** **RMB 660M order** for the **KUAE intelligent-computing cluster** — a 10,000-GPU system "capable of end-to-end training of trillion-parameter models" **[independent — TrendForce]**.
- **MUSA Developer Conference (Dec 20–21, 2025):** unveiled **Huagang** next-gen architecture (FP4–FP64 full precision, +50% compute density, 10× energy efficiency, scaling to 100,000+ GPU clusters) underpinning **Huashan** (AI training+inference) and **Lushan** (graphics) products; **MUSA 5.0** full-stack platform; KUAE 2.0 cluster data at 10K-GPU scale: **60% training efficiency** (dense), **40%** (MoE), >90% effective training time, ~95% linear scaling efficiency; **DeepSeek R1 671B** run on S5000 with SiliconFlow: **>4,000 tokens/s prefill and >1,000 tokens/s decode per card** — all vendor-reported figures, treat as **[vendor-reported]** **[secondary — Jiemian Global]**.
- **Capital markets:** began Hong Kong listing process in **June 2026** ("A+H" dual platform); hired EY as auditor **[secondary — BigGo Finance]**.
- Founder/CEO **Zhang Jianzhong** (ex-NVIDIA China head); company founded 2020; 3-year R&D spend ~3.8B yuan vs cumulative revenue; top 2024 customer Baidu (42% of sales) **[secondary — TMTPost]**.

### 2.3 Biren — blockbuster Hong Kong IPO

- **Listed January 2, 2026** on HKEX (**HKG: 6082**); raised **HK$5.58B (~$717M)**; retail oversubscription **2,348×**; shares spiked **+119% intraday**, closed **+76%** at HK$19.60 offer price **[secondary — financialcontent/tokenring; TrendForce]**.
- **Products:** **BR100** positioned as the domestic alternative to NVIDIA's high-end chips; BR104 successor in pipeline (per IPO-era coverage) **[secondary]**.
- **Pre-IPO:** June 2025 round of **CNY 1.5B** led by state-backed Guangdong/Shanghai investors; pre-IPO valuation **CNY 14B**; backers include Qiming, IDG Capital, Hillhouse, Walden, Ping An **[independent — TrendForce]**.
- Founded 2019, fabless general-purpose GPU.

### 2.4 MetaX — STAR listing, then A+H

- **Listed December 17, 2025** on STAR Market (12 days after Moore Threads); raised **3.9B yuan (~$578.9M)**; IPO price 104.66 yuan/share; opening market cap **41.8B yuan (~$6B)**; P/S 56.4× vs peer average 127.4× **[secondary — MEXC; BigGo Finance]**.
- **Products:** **Xiyun C-series** chips; claims support for full-scale pretraining of **128B-parameter MoE models**; CUDA-compatible software stack **[secondary — TMTPost]**.
- **Financials (from IPO-era filings):** Q1 losses of 2.3B yuan; **8.9B yuan unsold inventory**; receivables at **192% of Q1 revenue** — flagged by analysts as turnover/payment risk **[secondary — TMTPost]**.
- **June 2026:** initiated Hong Kong listing ("A+H" dual capital platform) **[secondary — BigGo Finance]**.

### 2.5 Sector picture

- Listing wave: Moore Threads (STAR, Dec 5 2025), MetaX (STAR, Dec 17 2025), **Biren (HK, Jan 2 2026)**, **Iluvatar CoreX (HK, Jan 8 2026)**, **Enflame** STAR IPO accepted Jan 22, 2026 (first A-share IPO accepted for review in 2026) **[secondary — BigGo Finance]**.
- Top-five domestic GPU makers raised **>30B yuan in secondary offerings** (2026) **[secondary — BigGo Finance]**.
- Domestic substitution rate for AI chips: **2–3.6% (2022–2024) → projected 31% by 2029** **[secondary — BigGo Finance]**.
- **Common bottleneck:** all rely on SMIC N+2 as the most stable domestic advanced node; N+3/N+4 multi-patterning yields stuck ~20%; equipment/spare-part import restrictions limit capacity expansion **[independent — TrendForce]**.
- Memory (HBM/LPDDR) availability is a parallel constraint **[independent — Bloomberg relay via techplugged]**.

---

