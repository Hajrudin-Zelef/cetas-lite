---
id: ai-industry-kb-2026-wave6/07-minimax/native-multimodality-and-computer-use
title: "Native multimodality and computer use"
domain: minimax
role: deep-dive
task: funding-deals
actors: ["Anthropic", "MiniMax", "OpenRouter"]
dates: ["2026-01", "2026-02-13", "2026-06-01"]
keywords: ["multimodal", "agent", "agentic", "alignment", "benchmarks", "claude", "cost", "fp8", "inference", "latency", "license", "memory"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3191, 3224]
section: "§7. MiniMax"
delta_of: ai-industry-kb-2026
sha256: 961d2e20efc2d993ef2b4212ad273f5102c5c1c5dbaa577f6cc4a993c1e9b173
---

# Native multimodality and computer use

### Native multimodality and computer use
- M3 accepts **images and video as input** and can **drive a computer** (computer use) — MiniMax's demo: the model **opens a local ERP system and enters invoices in batches** autonomously, moving between applications [VENDOR via secondary] (pasqualepillitteri.it).
- Multimodality built in from the start: the entire data pipeline was rebuilt to **over 100T training tokens with images and text aligned from the first step** — deep alignment of textual and visual semantic spaces lets the model treat a screenshot and a sentence as parts of the same reasoning [VENDOR via secondary] (pasqualepillitteri.it).
- Agentic endurance cases: autonomously **reproduced a paper accepted at ICLR 2025 — working nearly 12 hours, producing 18 commits and 23 experimental figures**; optimized a CUDA kernel lifting FP8 GEMM hardware utilization from **7.6% to 71.3% (9.4× gain) after 147 submissions and ~2,000 tool calls over 24 hours** [VENDOR via secondary] (pasqualepillitteri.it).

### MiniMax Agent (product alongside M3)
- Three new features: **persistent memory** ("Session Note" working pad — saves facts and intermediate results, summarizes older history to stay within token limits); **skills that evolve** (repeated complex tasks become custom skills, reusable later); **unified billing** (everything through the Token Plan — text, images, voice, music draw from one pool) [SECONDARY] (pasqualepillitteri.it).

### Pricing, API tiers, Token Plan (pasqualepillitteri.it; launch-day)
- **OpenRouter launch price: $0.30 input / $1.20 output per M tokens**, with a **50% debut discount** [SECONDARY].
- Official API pricing: **two tiers by context length** — standard up to **512,000 tokens**, higher tier beyond; **two service levels** (standard and priority for latency-sensitive industrial use) [SECONDARY].
- **Thinking mode** for complex reasoning, **non-thinking mode** for fast responses [SECONDARY].
- **Token Plan subscription: Plus $20/month, Max $50/month, Ultra $120/month** — growing token allowances shared across all modes [SECONDARY].
- **No particular geographic restrictions**: reachable via MiniMax APIs and aggregators like OpenRouter; pay-as-you-go and MiniMax Agent work the same everywhere [SECONDARY].
- The self-hosting argument: once weights are downloadable, in-house inference meets data-handling requirements without sending sensitive code overseas [SECONDARY analysis].


### New verified facts — expansion (continued — company, lineup, M2.5, M2.7)

### Company and model lineup
- Legal name: **Shanghai Xiyu Jizhi Technology Co., Ltd.**; founded in **Shanghai in late 2021**; builds in-house full-modality foundation models across text, video, voice, music, and images [COMMUNITY] (gameworkerkim/vibe-investing MiniMax coding guide).
- Listed on the Hong Kong Stock Exchange as **0100.HK in January 2026**; serves **over 200 million cumulative users across 200+ countries** [COMMUNITY] (vibe-investing guide).
- Flagship lineup (per the community guide): **M2.1** (text coding-focused, 197K ctx, 13+ languages, low cost, open-weight); **M2.5** (text agent, 197K, 230B/10B MoE, open-weight); **M2.7** (text agent, 205K, recursive self-improve, open-weight); **M3 (2026-06-01**, text+multimodal, 1M, MSA, open-weight planned); **Hailuo 2.3** (video, 1080p/10s, API only); **Speech 2.6 / Music 2.6** (voice/music, **40 languages, 250ms latency**, API only) [COMMUNITY] (vibe-investing guide).
- **M3 release date: 2026-06-01** per the guide — a precise date for the existing §7's M2→M3 arc [COMMUNITY] (vibe-investing guide).

### MiniMax M2.5 — release and spec
- **Released February 13, 2026** under the **MIT license** as an open-weights model [SECONDARY] (the-decoder.com, 2026-02-13).
- **230B total / 10B active** MoE; **197K context** [COMMUNITY] (vibe-investing guide).
- Trained using **reinforcement learning across hundreds of thousands of complex environments**; the model **learned to optimize its actions through its own planning** rather than only planning when prompted — aimed at strong results with low token consumption on long-running tasks [VENDOR via secondary] (the-decoder.com).
- Handles **Word, Excel, and PowerPoint files** quickly per the vendor [VENDOR via secondary] (the-decoder.com).
- Vendor benchmarks (all [VENDOR]): **SWE-Bench Verified 80.2%**; **Droid platform 79.7%**; **OpenCode 76.1%**; **BrowseComp 76.3%**; **BFCL (agent tool use) 76.8%** [SECONDARY reporting vendor] (the-decoder.com; vibe-investing guide).
- Price/performance: **80.2% SWE-bench Verified — only 1.8pp behind Claude Opus 4.7 (82.0%) — at roughly 1/17 the price** [COMMUNITY analysis] (vibe-investing guide).
- Framed by press as **"intelligence too cheap to meter"** — the pricing-pressure moment for Western labs [SECONDARY] (the-decoder.com).

