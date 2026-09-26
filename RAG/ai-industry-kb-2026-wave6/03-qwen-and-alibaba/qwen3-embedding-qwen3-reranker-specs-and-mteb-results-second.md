---
id: ai-industry-kb-2026-wave6/03-qwen-and-alibaba/qwen3-embedding-qwen3-reranker-specs-and-mteb-results-second
title: "Qwen3-Embedding / Qwen3-Reranker — specs and MTEB results [SECONDARY]"
domain: qwen-and-alibaba
role: deep-dive
task: actor-profile
actors: ["Alibaba", "California", "China", "Google", "United States"]
dates: ["2025-06", "2026-06", "2026-06-30", "2026-08-20", "2026-09-22"]
keywords: ["embedding", "qwen", "reranker", "capex", "compute", "consumer", "disclosure", "distribution", "fine-tuning", "gemini", "gpu", "inference"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1237, 1257]
section: "§3. Qwen and Alibaba"
delta_of: ai-industry-kb-2026
sha256: 170273429dd12f4f83957ba5113a145c7922a079d98c2076c4dbd318cf93f579
---

# Qwen3-Embedding / Qwen3-Reranker — specs and MTEB results [SECONDARY]

### Qwen3-Embedding / Qwen3-Reranker — specs and MTEB results [SECONDARY]
- Sizes: **0.6B / 4B / 8B** for both embedding and reranker (embedding: 28/36/36 layers, 32K sequence, 1024/2560/4096 dims; **MRL support, instruction-aware**). [SECONDARY]
- Vendor results (Qwen3-Embedding technical report, arXiv:2506.05176): **Qwen3-Embedding-8B ranked No.1 on the MTEB multilingual leaderboard (70.58)** as of June 2025; 4B/8B best overall on MMTEB; the 0.6B model "only lags behind" Gemini-Embedding despite 0.6B params; full STS table (8B: e.g., 73.84 / 75.00 / 76.97 / 80.08 across columns vs multilingual-e5-large-instruct 58.08 / 58.24 / 69.80 / 48.23). [SECONDARY]
- Reranker results: Qwen3-Reranker-4B **69.76** MTEB-R / 75.94 CMTEB-R; Qwen3-Reranker-8B 69.02 / **77.45** CMTEB-R / **72.94** MMTEB-R; Qwen3-Reranker-0.6B 65.80 / 71.31 — all beating BGE-reranker-v2-m3 (57.03 / 72.16) and Jina multilingual reranker v2 base (58.22 / 63.37). [SECONDARY]
- **MTEB Code**: Qwen3-Embedding-8B **80.68** — surpassing Gemini-Embedding (proprietary SOTA at the time). [SECONDARY]
- **Training recipe** (technical report): synthetic multilingual/multitask relevance data from Qwen3-Instruct for stage-1 unsupervised; high-quality small-scale supervised stage-2; rerankers: SFT + **model merging**; instruction-aware (1–5% boost with task instructions; English instructions best for multilingual). [SECONDARY]
- **Throughput** (ChunkHound consumer-GPU test, 10k Python files): 0.6B **2,100 docs/sec** (4.8s), 4B **1,200 docs/sec** (8.3s), 8B **650 docs/sec** (15.4s) — 4B recommended for speed/accuracy tradeoff. [COMMUNITY]
- **Caveat**: leaderboard #1 ≠ universal superiority — MTEB was designed because embedding performance varies by task; production choice depends on language distribution, domain, latency, vector count. [DIRECTIONAL]
- Community usage: chunkhound recommends the 4B as the speed/accuracy tradeoff (0.6B: 2,100 docs/s; 4B: 1,200 docs/s; 8B: 650 docs/s on consumer GPU). [COMMUNITY]

### Alibaba cloud/AI revenue and capex — June quarter 2026 [SECONDARY]
- **2026-08-20 earnings** (quarter ended 2026-06-30, per the businesswire release and secondary compilations): **AI Cloud and Compute Services revenue RMB48.44B (~$7.14B), +45% YoY** — external-customer revenue also +45%; **AI-related products RMB12.38B (~$1.82B), 12th consecutive quarter of triple-digit growth** (~26% of AI Cloud sales); cloud adjusted EBITA **+133% to RMB5.63B** (margin ~12% vs ~7% a year earlier). Group revenue RMB268.95B (+9%). [SECONDARY]
- The bill: group operating income **−57% to RMB15.16B**; net income attributable to ordinary shareholders **−75% to RMB10.44B**; capex **+75% to RMB67.68B (~$10B)** on AI infrastructure and compute capacity; AI Labs and Applications lost **RMB13.86B adjusted EBITA** (~2.5× the AI Cloud profit) on RMB3.34B revenue (+16%) — frontier-model, Qwen inference, and consumer-AI spend. [SECONDARY]
- **2026-09** (DCD, Q2 FY2026): cloud unit revenue **+34% YoY to $5.6B**, AI-related products triple-digit for the **9th consecutive quarter** — the 9-vs-12 discrepancy is a **fiscal-quarter vs calendar-quarter basis difference**, not a data conflict. CEO Eddie Wu: "demand is accelerating"; Alibaba is **rationing server access** — full-stack cloud customers get priority over single-GPU renters; supply-chain undersupply expected for **"two to three years"**; the Oct-2025-claimed **GPU pooling system** reportedly cut GPU use **82%**. [SECONDARY]
- Hardware hedge: Alibaba unveiled a **new proprietary T-Head AI chip** (~2026-09-22) and stated plans for **5–10 trillion parameter models** — 2–4× the Qwen3.8-Max scale — plus exploration of **recursive self-improvement**. Omdia: Alibaba Cloud **#1 in China's AI cloud market, 38.1% share** (2025). [SECONDARY]
- **AI revenue-share milestone**: June quarter 2026 — **AI now drives >1/3 of cloud revenue** (first explicit disclosure). AI-related products: **RMB12.38B/quarter, ~$7.3B annualized run rate**. [SECONDARY]
- **US DoD designation**: June 2026 — designated a **military-linked firm** by the US Defense Department; Alibaba challenged via **lawsuit in California federal court**. [SECONDARY]
- **Zhenwu M890 adoption**: **>650 external customers across 20+ industries** (autonomous driving, internet, financial services) via Alibaba Cloud — training, fine-tuning, inference. [SECONDARY]

---

