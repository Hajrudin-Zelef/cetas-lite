---
id: labs-hyperscalers-2026/00-labs-hyperscalers/4-3-muse-glimmer-30b-open-weights-apache-2-0
title: "4.3 Muse Glimmer 30B — open weights (Apache 2.0)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "EU", "Fireworks AI", "Google", "Hugging Face", "Intel", "Meta", "Nebius", "Nvidia", "OpenAI", "SGLang", "Together AI", "United States", "vLLM", "xAI"]
dates: ["2025-06", "2026-04-08", "2026-07", "2026-07-07", "2026-07-09", "2026-07-12", "2026-08-05", "2026-08-10", "2026-09-02"]
keywords: ["apache", "muse", "open weights", "agent", "agentic", "amd", "astra", "benchmark", "benchmarks", "capex", "claude", "compute"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [694, 741]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 53021f3063fc1041f8f834f83d7f38ab399a138f6ecf0036bf74de72ba0bd6e4
---

# 4.3 Muse Glimmer 30B — open weights (Apache 2.0)

### 4.3 Muse Glimmer 30B — open weights (Apache 2.0)
- **Muse Glimmer 30B** released under **Apache 2.0** — Meta's open-weights offering in the window, continuing the Llama-era open strategy under the new Muse brand. [secondary: Track B]
- 30B parameters; positioned for on-device / edge / fine-tuning use cases. [secondary]

### 4.4 Capex, compute, MTIA
- **Meta capex** remained at historic highs through 2026 to fund AI data-center buildout (Track B; quarterly figures vary — pull from earnings for precision).
- **MTIA "Iris"** — Meta's next-gen AI training/inference chip program referenced in Track B (naming from roadmaps — [secondary]).
- Meta's compute fleet expansion (hundreds of thousands of GPUs) continued; **Nebius–Meta orders** noted in Track D (see §13).
- **Meta–Nebius**: Nebius reported Meta as a customer / order flow in 2026 (see §13.5). [secondary]

### 4.5 Products & partnerships
- **Meta AI** assistant continued rollout across Facebook, Instagram, WhatsApp, Messenger. [secondary]
- **Ray-Ban Meta / Oakley Meta** smart glasses — AI wearable line; 2026 iterations with deeper Muse integration (Track B).
- **Meta–Scale AI**: Meta's **$14.3B investment in Scale AI (June 2025, pre-window)** — Alexandr Wang joined Meta as Chief AI Officer; 2026 execution continued under the new **Meta Superintelligence Labs** structure. [secondary: Track B]
- **Meta–Nvidia**: continued GPU purchasing at scale. [secondary]
- **Meta–AMD / MTIA**: heterogeneous silicon strategy. [secondary]

### 4.6 Controversies & regulatory
- **EU AI Act / DSA**: Meta's AI products in scope; compliance posture through 2026 (Track B).
- **Llama-era data lawsuits** continued through 2026 (Track B — specific filings to verify).
- **"A Call for Collective Action on Cyber Defense"** (27 Aug 2026) — Meta signatory status not confirmed in sources returned (gap).

### 4.7 Meta benchmark snapshot (as reported)

| Model | Benchmark | Score | Provenance |
|---|---|---|---|
| Muse Spark 1.3 | DeepSWE v1.1 | 75.4% | [vendor-reported] |
| GPT-6 Astra (for comparison) | DeepSWE v1.1 | 74.1% | [vendor-reported] |

---
### 4.8 Muse family — detailed release record (from Track B)

Meta Superintelligence Labs (MSL), led by Chief AI Officer **Alexandr Wang** (ex-Scale AI, joined via the ~$14B Scale AI investment/deal in 2025), pivoted Meta from pure open-weights to a dual strategy: closed API models (Muse Spark) + open-weights local models (Muse Glimmer). This is Meta's **first-ever paid model API** — a deliberate strategic shift from the Llama tradition [secondary] https://mlq.ai/news/meta-launches-muse-spark-11-api-its-first-paid-ai-model-to-challenge-anthropic-and-openai/.

**Muse Spark 1.0 (April 8, 2026)**: first major Meta AI model since the Scale AI deal; announced by Mark Zuckerberg at the Meta AI event; no public developer API at launch [secondary — CNBC] https://www.cnbc.com/2026/04/08/meta-debuts-first-major-ai-model-since-14-billion-deal-to-bring-in-alexandr-wang.html. Reported benchmark framing (vendor-adjacent aggregators): 77.4% coding vs. top rivals, 59% on Terminal-Bench 2.0; Meta claims parity/superiority on agentic and frontend/visual tasks vs. Claude Opus 4.6, GPT-5.4/5.5, Gemini 3.1 Pro — treat individual numbers as [unverified]. Independent commentary notes strength in visual/UI reasoning, weaker on heavy backend programming [unverified/secondary].

**Muse Spark 1.1 (July 9, 2026) — first commercial developer API**: multimodal (text, image, video, PDF, audio in → text out) reasoning model for agentic tasks: tool use, computer use, coding, MCP servers, custom skills, multi-agent orchestration, repo-level code edits, built-in web-search grounding. Context window: **1M tokens** (up from 262K in 1.0) with active context management. **Pricing: $1.25 / 1M input, $4.25 / 1M output** (reasoning tokens billed as output); $20 one-time free credit; `reasoning_effort` parameter (minimal → xhigh). API US-only during public preview; endpoint `api.meta.com` — model id `muse-spark-1.1`; supports OpenAI Chat Completions and Anthropic Messages formats [vendor-reported/secondary] https://pondero.ai/news/2026-07-12-meta-muse-spark-1-1-launch-model-api/. Benchmarks as Meta presents them: 54 on the Artificial Analysis Intelligence Index (level with Grok 4.5); Meta claims parity with GPT-5.5 and Claude Opus 4.8 "across many agentic evals" [vendor-reported/secondary]. Launched same day in "Thinking" mode in the **Meta AI app** and on meta.ai; early API partners: Replit, Cline, Box [secondary]. Zuckerberg announced it on X — described as his first post in three years — calling it *"a strong agentic and coding model at a very low price"* [secondary]. Same week: **Muse Image** (July 7, 2026, MSL's first image-generation model) and **Muse Video** launched for creative workflows [secondary].

**Muse Spark 1.2 (August 5, 2026)**: second major version; shipped closed via the Meta Model API at **$1.25/1M input** (same headline pricing); positioned for complex software engineering. On Aug 10, Meta announced open weights for Spark 1.2 will be released "soon" under a permissive license — weights not yet published as of Sep 22, 2026 [secondary]. Specific benchmark deltas 1.1 → 1.2 not verified in fetched sources.

**Muse Spark 1.3 (September 2, 2026)**: noted in DataCamp's Muse Spark 1.1 writeup — no specifications, pricing, or benchmark details verified [unverified]. **DeepSWE v1.1 75.4%** (Meta-reported, beating GPT-6 Astra's 74.1%) — the only verified 1.3 datapoint [vendor-reported].

**Muse Glimmer (August 10, 2026) — the open-weights release**: **30B-parameter dense multimodal model** (29.6B by one count) from MSL, released as **open weights under the Apache 2.0 license** — fully free for commercial use — on Hugging Face. First Muse-family model with published weights [secondary, widely reported — NYT, Reuters, The Guardian, The Register, Bloomberg] https://www.datacamp.com/blog/muse-glimmer. Purpose-built for **local, always-on agentic workflows** on consumer hardware: quantized 4-bit builds fit **under 20 GB** (from ~55 GB), running on a **single 24 GB (or 32 GB) consumer GPU** or Mac-class unified memory; **120K–131K token context**; 1.8B-parameter perception/vision encoder (text + image in; text out only; video as frames; no audio); native runtime integrations: llama.cpp, MLX, ExecuTorch, Ollama, LM Studio, vLLM, SGLang; hardware optimizations listed for AMD, Arm, Dell, Intel, Nvidia. No first-party Meta-hosted API — self-host or third-party providers (e.g. Together AI, Fireworks AI) [secondary]. Training: logit distillation from the Muse Spark teacher line (Bloomberg: distilled from Muse Spark 1.2), long-context agentic data, RL [secondary]. Meta-reported benchmarks vs. size class: leads on **MCP-Atlas (75.5)** and **DeepSearch QA (74.6)** vs. Qwen3.6-27B and Gemma4-31B; trails Qwen3.6-27B on OSWorld-Verified and Terminal-Bench 2.1 [vendor-reported/secondary].

**Pricing context (July 2026 list prices, per secondary reporting)**: GPT-5.6 Luna $1.00/$6.00; GPT-5.6 Sol $5.00/$30.00; Claude Sonnet 5 $2.00/$10.00; GPT-5.6 Terra $2.50/$15.00. Muse Spark 1.1 positioned 50–75% cheaper than its closest rivals [secondary].

### 4.9 Zuckerberg's open-source essay (Aug 10, 2026)

