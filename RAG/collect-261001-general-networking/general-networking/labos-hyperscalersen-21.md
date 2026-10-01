---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-21
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "China", "Fireworks AI", "Hugging Face", "Intel", "Meta", "Nvidia", "OpenAI", "SGLang", "Together AI", "United States", "vLLM", "xAI"]
dates: ["2025-07", "2026-03", "2026-04-29", "2026-07", "2026-07-07", "2026-07-09", "2026-07-12", "2026-08-05", "2026-08-10", "2026-09-02"]
keywords: ["agent", "agentic", "amd", "apache", "astra", "benchmark", "benchmarks", "capex", "claude", "compute", "consumer", "context window"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [730, 758]
sha256: 97d77bacfe81d95dfbad9845614865f3f49495a4775ec52c66e2c636d1217a94
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

**Muse Spark 1.1 (July 9, 2026) — first commercial developer API**: multimodal (text, image, video, PDF, audio in → text out) reasoning model for agentic tasks: tool use, computer use, coding, MCP servers, custom skills, multi-agent orchestration, repo-level code edits, built-in web-search grounding. Context window: **1M tokens** (up from 262K in 1.0) with active context management. **Pricing: $1.25 / 1M input, $4.25 / 1M output** (reasoning tokens billed as output); $20 one-time free credit; `reasoning_effort` parameter (minimal → xhigh). API US-only during public preview; endpoint `api.meta.com` — model id `muse-spark-1.1`; supports OpenAI Chat Completions and Anthropic Messages formats [vendor-reported/secondary] https://pondero.ai/news/2026-07-12-meta-muse-spark-1-1-launch-model-api/. Benchmarks as Meta presents them: 54 on the Artificial Analysis Intelligence Index (level with Grok 4.5); Meta claims parity with GPT-5.5 and Claude Opus 4.8 "across many agentic evals" [vendor-reported/secondary]. Launched same day in "Thinking" mode in the **Meta AI app** and on meta.ai; early API partners: Replit, Cline, Box [secondary]. Zuckerberg announced it on X — described as his first post in three years — calling it *"a strong agentic and coding model at a very low price"* [secondary]. Same week: **Muse Image** (July 7, 2026, MSL's first image-generation model) and **Muse Video** launched for creative workflows [secondary].

**Muse Spark 1.2 (August 5, 2026)**: second major version; shipped closed via the Meta Model API at **$1.25/1M input** (same headline pricing); positioned for complex software engineering. On Aug 10, Meta announced open weights for Spark 1.2 will be released "soon" under a permissive license — weights not yet published as of Sep 22, 2026 [secondary]. Specific benchmark deltas 1.1 → 1.2 not verified in fetched sources.

**Muse Spark 1.3 (September 2, 2026)**: noted in DataCamp's Muse Spark 1.1 writeup — no specifications, pricing, or benchmark details verified [unverified]. **DeepSWE v1.1 75.4%** (Meta-reported, beating GPT-6 Astra's 74.1%) — the only verified 1.3 datapoint [vendor-reported].

**Muse Glimmer (August 10, 2026) — the open-weights release**: **30B-parameter dense multimodal model** (29.6B by one count) from MSL, released as **open weights under the Apache 2.0 license** — fully free for commercial use — on Hugging Face. First Muse-family model with published weights [secondary, widely reported — NYT, Reuters, The Guardian, The Register, Bloomberg] https://www.datacamp.com/blog/muse-glimmer. Purpose-built for **local, always-on agentic workflows** on consumer hardware: quantized 4-bit builds fit **under 20 GB** (from ~55 GB), running on a **single 24 GB (or 32 GB) consumer GPU** or Mac-class unified memory; **120K–131K token context**; 1.8B-parameter perception/vision encoder (text + image in; text out only; video as frames; no audio); native runtime integrations: llama.cpp, MLX, ExecuTorch, Ollama, LM Studio, vLLM, SGLang; hardware optimizations listed for AMD, Arm, Dell, Intel, Nvidia. No first-party Meta-hosted API — self-host or third-party providers (e.g. Together AI, Fireworks AI) [secondary]. Training: logit distillation from the Muse Spark teacher line (Bloomberg: distilled from Muse Spark 1.2), long-context agentic data, RL [secondary]. Meta-reported benchmarks vs. size class: leads on **MCP-Atlas (75.5)** and **DeepSearch QA (74.6)** vs. Qwen3.6-27B and Gemma4-31B; trails Qwen3.6-27B on OSWorld-Verified and Terminal-Bench 2.1 [vendor-reported/secondary].

**Pricing context (July 2026 list prices, per secondary reporting)**: GPT-5.6 Luna $1.00/$6.00; GPT-5.6 Sol $5.00/$30.00; Claude Sonnet 5 $2.00/$10.00; GPT-5.6 Terra $2.50/$15.00. Muse Spark 1.1 positioned 50–75% cheaper than its closest rivals [secondary].

### 4.9 Zuckerberg's open-source essay (Aug 10, 2026)

- Zuckerberg published a **~6,500-word essay, "The Future is for Everyone,"** on Meta's website on Aug 10, 2026, simultaneous with the Muse Glimmer release [secondary — Reuters Breakingviews; The Guardian].
- Core arguments: everyone should have a personalized "superintelligence" agent; concentrated AI control (by companies, governments, or AI itself) is the biggest risk; **open weights "will be the best way to protect safety and security over time"**; US should loosen restrictions on AI training data to keep US open-weight labs competitive with China; build energy capacity faster [secondary].
- Practical commitments: independent safety review board to approve safety rules before future model releases; **$1B community fund** for communities near Meta data centers; 2026 infra spend $130–145B; offer to share some internal AI training data with officials [secondary].
- Context: shift from his July 2025 stance that Meta likely wouldn't open-source all superintelligence models; follows the Jan 2026 Llama 4 benchmark controversy (Zuckerberg acknowledged the LM Arena test-version vs. shipped-version mismatch — specifics [unverified]) [secondary/unverified].
- The letter's "exceptionally capable personal agent for everyone" and "fully private mode where even Meta cannot see data" do not exist in any shipped Meta product — quote the letter as direction, not capability [secondary].

### 4.10 FAIR / Meta AI research milestones 2026

- **"Beyond Language Modeling" (March 2026)** — large-scale study on building native multimodal foundation models from scratch (no language-pretrained backbone), using the Transfusion framework; findings: RAE encoder built on SigLIP 2 beats VAE encoders/raw pixels; mixed text+vision training is synergistic; MoE architectures naturally specialize experts by modality [secondary].
- **JASCO** — joint audio-and-symbolic conditioning model for temporally controlled text-to-music generation; paper + sample page released; inference code slated for AudioCraft repo under MIT [official — Meta AI blog].
- **NeuralSet (April 29, 2026)** — Python package unifying neural recordings (fMRI, M/EEG, spikes) with deep learning: single PyTorch-ready DataLoader, HuggingFace embeddings support [secondary].
- **AI Research Preference Models / RPMs (Sep 6, 2026)** — FAIR + Oxford + UCL paper: RPMs rank unexecuted ML-experiment candidates before spending GPU hours; scaffold and AIRS-Bench open source; backbone Qwen3.6-27B [secondary].
- **SAM 3** — next-gen "Promptable Concept Segmentation" (2× claimed gain over prior systems); research paper submitted to ICLR 2026, models NOT yet publicly released [unverified].
- **Byte-model fact-check (Sep 2026)** — a viral X thread misrepresented Meta's Dec-2024 Byte Latent Transformer as new 2026 work; the actual new paper (UW + Meta FAIR: "Breaking the Token Ceiling") shows byte models keep improving with compute and can exceed token models' ceiling under distillation with ~1/6th the data [secondary].

### 4.11 Meta infrastructure: capex, chips, data centers

