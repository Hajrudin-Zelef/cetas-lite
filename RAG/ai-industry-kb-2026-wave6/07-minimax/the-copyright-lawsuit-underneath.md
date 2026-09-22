---
id: ai-industry-kb-2026-wave6/07-minimax/the-copyright-lawsuit-underneath
title: "The copyright lawsuit underneath"
domain: minimax
role: deep-dive
task: actor-profile
actors: ["Anthropic", "ByteDance", "China", "EU", "Google", "MiniMax", "OpenAI", "OpenRouter", "United States"]
dates: ["2025-09", "2026-01", "2026-02-13", "2026-05-26", "2026-06", "2026-06-01", "2026-06-07"]
keywords: ["copyright", "lawsuit", "agent", "agentic", "alignment", "attention", "benchmark", "benchmarks", "claude", "compute", "context window", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3150, 3224]
section: "§7. MiniMax"
delta_of: ai-industry-kb-2026
sha256: 7946cb4ead18d0abed47cd47e7b25f64bc4c424ecf5d31996d6a8d3a31a5df8f
---

# The copyright lawsuit underneath

### The copyright lawsuit underneath
- **Disney, Universal, and Warner Bros. Discovery jointly sued MiniMax and its Hailuo platform in September 2025**, alleging the company trained its video/image models on **unauthorized copies of their characters** and that Hailuo could generate recognizable **Spider-Man, Darth Vader, and Shrek** from simple text prompts [SECONDARY].
- MiniMax sought dismissal on jurisdictional grounds (US courts have no authority over a Chinese company); on **May 26, 2026, a federal judge rejected that argument** and let the case proceed into **full discovery** [SECONDARY].
- **H3 runs on the same Hailuo infrastructure named in the lawsuit** [SECONDARY].

### Pricing and market reaction
- MiniMax's claim: **2K generation costs less than a third of mainstream competitors; 768p costs less than half of what competitors charge at 720p** (competitors unnamed) [VENDOR via secondary].
- Independent breakdown: **H3 2K ~$7.80/min vs Kling 3.0 ~$20.16/min (1080p) vs ByteDance Seedance 2.0 ~$22.45/min (1080p)** — "lines up reasonably well" with MiniMax's claim; treat exact dollars as directional (one explainer, not a controlled benchmark) [SECONDARY].
- AA independent scores: **H3 ranked 1st globally in video editing at launch**; **top 3 in text-to-video and image-to-video**; trailed **Google's Gemini Omni Flash** on T2V and **Gemini Omni Flash + Seedance 2.0** on I2V specifically [SECONDARY].
- Stock: MiniMax shares climbed **roughly 13% intraday on heavy turnover** on the H3 news; **ByteDance rolled out Seedance 2.5 the same week** [SECONDARY].

### Limitations (all [SECONDARY])
- **15 seconds is a hard ceiling** per clip; Seedance 2.5's 30-second clips are a real advantage for narrative work.
- 2K requires an API call even for local open-weights users.
- US/EU/UK/SK users need an **individual license** for legal self-hosting.
- Every API call routes through **infrastructure in China**, subject to China's National Intelligence Law.
- Vendor pricing claims/demos are promotional; **AA scores reflect the hosted API pipeline, not a locally run H3-Base render**.
- **No independent security audit of the released weights**; training-data provenance not independently verified.


### New verified facts — expansion (continued — M3 launch day, stock story, Kilo audit, MSA internals, MiniMax Agent)

### Launch day and the stock market (pasqualepillitteri.it, June 2026)
- **June 1, 2026**: MiniMax stock **climbed more than 5%** on the Hong Kong Stock Exchange, touching **907.5 Hong Kong dollars**, then **collapsed to a 12.38% loss within a few hours** on turnover of **1.49 billion Hong Kong dollars** — classic "sell the news" behavior [SECONDARY] (pasqualepillitteri.it).
- Same-day headline: M3 billed as the **first open-weights system to combine frontier coding, a 1M-token context window, and native multimodality** [SECONDARY] (pasqualepillitteri.it).
- Vendor claim at launch: **59.0% SWE-Bench Pro** — per MiniMax above **GPT-5.5 and Gemini 3.1 Pro**, close to **Claude Opus 4.7** [VENDOR via secondary] (pasqualepillitteri.it).
- **Open-weights caveat at launch**: at announcement the model was accessible via **API and MiniMax Agent**; actual weights were **not yet public** — MiniMax promised weights plus the full technical report **within roughly 10 days** of launch. The article frames it as a statement of intent, not a delivered fact [SECONDARY] (pasqualepillitteri.it).

### Kilo independent audit (June 7, 2026)
- The Kilo team gave **the same code and the same prompt** to **MiniMax M3 and Claude Opus 4.8**: find **17 real bugs** deliberately planted in a webhook delivery service (TypeScript, Bun, SQLite) [SECONDARY] (pasqualepillitteri.it).
- Result: **M3 found 13/17 bugs at $0.07** in 5m 03s; **Opus 4.8 (medium) found 13/17 at $1.30** (3m 53s); **Opus 4.8 (high) 13/17 at $1.93** (4m 34s); **Opus 4.8 (xhigh) 15/17 at $2.03** (7m 26s); **Opus 4.8 (max) 15/17 at $3.39** (9m 24s) [SECONDARY].
- Cost ratios: Opus at mid-tier costs **18–27× more** for the same 13 bugs; at max it wins by 2 bugs at **up to 48× the per-run cost** [SECONDARY].
- What M3 missed (caught by Opus at higher reasoning levels): **invalid JSON returning a 500 error, database setup running at import time, async callback inside a synchronous transaction** — subtle bugs separating quick review from exhaustive review [SECONDARY].
- The article's practical read: for high-volume/repeated audits M3 delivers the same harvest at a fraction of the price; for maximum coverage with unconstrained budget, Opus 4.8 at full reasoning stays ahead [SECONDARY].

### MSA internals ("KV outer gather Q")
- MSA = **MiniMax Sparse Attention**; each token attends only to the tokens that matter; compute per token at 1M context **drops to 1/20th of the previous generation** [VENDOR via secondary] (pasqualepillitteri.it).
- Decoding **up to 15.6× faster**, prefill **more than 9× faster** than M2 on very long contexts [VENDOR via secondary] (pasqualepillitteri.it).
- **Why MSA instead of lightning attention**: MiniMax's previous models used "lightning attention" (a linear variant); M3 returns to sparse attention with a new implementation called **"KV outer gather Q"**, in which each memory block is **read only once** and accesses are **contiguous** — claimed at **more than 4× the gain of Flash-Sparse-Attention**, a reference implementation [VENDOR via secondary] (pasqualepillitteri.it).
- Inference efficiency as the control point: "whoever controls this layer controls the costs" — the same lesson the article ties to antirez's local-inference work [SECONDARY analysis] (pasqualepillitteri.it).

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

