---
id: labs-hyperscalers-2026/00-labs-hyperscalers/10-independent-labs
title: "§10 — INDEPENDENT LABS"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "Baseten", "China", "Cohere", "DeepSeek", "EU", "Fireworks AI", "Hugging Face", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "Perplexity", "Sakana", "Stability AI", "StepFun", "Together AI", "Z.ai"]
dates: ["2025-02", "2025-07", "2025-10", "2026-03", "2026-06", "2026-07", "2026-09"]
keywords: ["acquisition", "agent", "agentic", "amd", "apache", "benchmark", "blackwell", "cohere", "consumer", "cost", "deepseek", "export controls"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1407, 1462]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 294b0dca43f402340daaa8acfc15beb272436316c636792902093e5ed6e168d3
---

# §10 — INDEPENDENT LABS

## §10 — INDEPENDENT LABS

> **2026 at a glance — independent labs:** Thinking Machines open-released Inkling (975B/41B, Jul 15); Sakana's Fugu orchestration went Max/Ultra v2 (Sep 11–12) routing around export controls; Cohere shipped Command A+ (Apache 2.0, May 20–21) and signed the Aleph Alpha merger (Sep 16); Perplexity passed $750M annualized revenue; smaller 2024-vintage labs went quiet.

### 10.1 Thinking Machines Lab (Mira Murati)
- **Inkling** — Thinking Machines' model release in the window. **Inkling is a model, not an assistant or codename** (correction preserved from Track C). [secondary]
- Architecture: **975B total parameters / 41B active MoE**, **open weights**. [secondary]
- One of the most significant open-weights releases of 2026 by parameter count. [secondary]

### 10.2 Cohere × Aleph Alpha — definitive merger agreement (16 Sept 2026)
- **Cohere and Aleph Alpha signed a definitive merger agreement on 16 September 2026** — [vendor-reported/press]; **regulatory approval pending** as of 22 Sept 2026. [secondary: Track C]
- Creates a transatlantic enterprise-AI champion (North America + EU sovereign-AI credentials). [secondary]

### 10.3 Sakana AI
- **Benchmark claims remain vendor-reported** and require caveats — Sakana's 2026 evaluation claims (Track C) should not be cited without the vendor-reported tag. [vendor-reported]
- Continued research output in evolutionary / nature-inspired AI. [secondary]

### 10.4 Other independent labs (2026 activity)
- **Cognition (Devin)** — agentic coding company; 2026 product/funding milestones (Track C).
- **Anysphere (Cursor)** — **acquisition status disputed/unverified**: reports of an OpenAI acquisition circulated (see §2.2 OpenAI uncertainties); **treat as [unverified]** — no confirmed close as of 22 Sept 2026. Cursor remained an independent product surface in wide use. [unverified]
- **Perplexity** — 2026 funding/product milestones; Comet browser; publisher partnerships and lawsuits (Track C).
- **Character.AI** — continued consumer AI companion scale; Noam Shazeer departed for OpenAI (18 June 2026 — see §2.4). [secondary]
- **Runway / Luma / Pika** — generative media models; 2026 releases (Track C).
- **ElevenLabs** — voice AI; 2026 milestones (Track C).
- **Sierra / Adept / Imbue** — agent startups; 2026 activity (Track C).
- ** poolside / Magic / Reflection** — coding-model labs; 2026 activity (Track C).
- **Liquid AI** — 2026 model releases (Track C).
- **Nous Research** — open-weights collective; 2026 releases (Track C).
- **Allen Institute (Ai2)** — OLMo lineage; 2026 open-model releases (Track C).
- **IBM Granite** — enterprise open models; 2026 iterations (Track C).
- **Hugging Face** — beyond the July breach (see §2.4): continued as the central open-model hub; **NVIDIA→Hugging Face definitive acquisition agreement 3 Sept 2026** (see §17). [independent/secondary]
- **Stability AI** — 2026 restructuring/releases (Track C).
- **Black Forest Labs (FLUX)** — image models; 2026 releases (Track C).
- **DeepSeek** — 2026 releases closely watched post-R1/V3 (Track C).
- **Qwen (Alibaba)** — Qwen 2026 releases; open-weights leadership. [secondary]
- **Kimi (Moonshot AI)** — long-context models; 2026 releases (Track C).
- **Zhipu (Z.AI) / MiniMax / StepFun** — Chinese labs; 2026 releases (Track C).

---
### 10.5 Thinking Machines Lab — detailed (from Track C)

**Founding context:** Founded by Mira Murati, former CTO of OpenAI (departed Sept 2024). Company incorporated/launch announced February 2025 (SiliconANGLE). Founding team drawn largely from former OpenAI researchers; post-launch co-founder departures: co-founder/CTO Barret Zoph and co-founder Luke Metz returned to OpenAI amid a collapsed follow-on fundraise (Jan 2026) [secondary]. Headcount ~200 after Jan 2026 departures [secondary — Revelio Labs via valueaddvc, mid-2026]. Leadership beyond Murati not independently verified here.

**Funding:** July 2025: ~$2B seed round at $12B valuation, led by Andreessen Horowitz; participants included Nvidia, Accel, ServiceNow, Cisco Ventures, AMD, Jane Street [independent] https://siliconangle.com/2025/07/15/thinking-machines-led-former-openai-cto-mira-murati-raises-2b-seed-funding/ — one of the largest seed rounds ever disclosed. March 2026: multi-year partnership/deal with Nvidia — significant investment (undisclosed terms) plus procurement of at least **1 GW of next-generation Nvidia processors** (Vera Rubin systems, deployment starting 2027) [independent]. Industry estimates ~$50B cost for 1 GW-class capacity [secondary]. Failed follow-on: talks for a round targeting $50–60B valuation collapsed Jan 2026 [secondary]. Sept 2026 (talks, NOT closed): raising $5–6B at ≥$40B pre-money valuation; Accel reportedly leading, Nvidia discussing ~$2.5B participation [secondary/unverified — Forkast, 4 Sept 2026]. Treat as unconfirmed.

**Products:**
- **Tinker** — fine-tuning platform / training API, launched October 2025 [vendor-reported/secondary]. Lets companies fine-tune base models on their own data and own the resulting model. Cited enterprise customer: hedge fund Bridgewater Associates, which reportedly fine-tuned an open model via Tinker, scoring 84.7% on (company-evaluated) financial reasoning tests [vendor-reported].
- **Inkling** (released 15 July 2026) — Thinking Machines' first in-house foundation model, and **a model, not an assistant or codename** [vendor-reported/secondary, multiple outlets] https://fourweekmba.com/ai-thinking-machines-inkling-open-weights-customization-strateg/:
  - MoE transformer; 975B total parameters, 41B active per forward pass; pretrained on 45T tokens (text, images, audio, video); up to 1M-token context; native text/image/audio reasoning, outputs text + code; controllable "thinking effort" dial.
  - **Open-weight release** on Hugging Face (Apache 2.0 per secondary coverage; includes a Blackwell-optimized checkpoint). The lab explicitly says Inkling is "not the strongest overall model available today" — positioned as a broad, balanced base for customization, flagship model for Tinker [vendor-reported].
  - Benchmark claims: e.g. matches Nvidia Nemotron 3 Ultra on one coding benchmark at ~1/3 the tokens [vendor-reported; independent evaluation not yet published].
  - Inference live on Tinker, Together AI, Fireworks, Modal, Databricks, Baseten [secondary].
  - Access/pricing: reported as free of charge on Tinker at launch [secondary]; no per-token list pricing located. Mark: **pricing unconfirmed**.
- **Inkling-Small** — previewed alongside Inkling: 276B total / 12B active parameters; full weights planned after testing [vendor-reported/secondary].
- No consumer assistant product announced as of Sept 2026.

