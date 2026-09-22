---
id: labs-hyperscalers-2026/00-labs-hyperscalers/4-3-muse-glimmer-30b-open-weights-apache-2-0
title: "4.3 Muse Glimmer 30B — open weights (Apache 2.0)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "Broadcom", "China", "EU", "Fireworks AI", "Google", "Hugging Face", "Intel", "Meta", "Microsoft", "Nebius", "Nvidia", "OpenAI", "SGLang", "TSMC", "Together AI", "United States", "vLLM", "xAI"]
dates: ["2025-04", "2025-06", "2025-07", "2026-03", "2026-04-08", "2026-04-29", "2026-07", "2026-07-07", "2026-07-09", "2026-07-12", "2026-07-28", "2026-08-05", "2026-08-10", "2026-09", "2026-09-02", "2026-09-18"]
keywords: ["apache", "muse", "open weights", "accelerator", "agent", "agentic", "amd", "astra", "benchmark", "benchmarks", "blackwell", "capex"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [694, 777]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: ee2d43d4c3b4805f755f6bdc6e6b6f7f6abfecef15c1eee66307c28dbd0acf7a
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

- **Capex guidance 2026:** initially $115–135B; raised (April) to **$125–145B** including finance-lease payments; Q2 reporting raised the floor again to **$130–145B**. ~70% of capex is now AI infrastructure. Q2 2026 revenue **$60.8B (+28% YoY)**; profit missed expectations; shares fell on margin concerns [secondary].
- **Compute targets (Reuters, July 9, 2026, internal comms):** ~7 GW of compute capacity by end of 2026 (1 GW deployed H1 2026 + 2.5 GW planned H2), doubling to **14 GW in 2027** [secondary].
- **MTIA "Iris":** next MTIA accelerator reportedly begins production **September 2026**; co-developed with **Broadcom**, manufactured by **TSMC**; plan to refresh chips roughly every 6 months through 2027 [secondary — Reuters via DataFLOQ].
- **AMD deal (Feb 2026):** up to **$60B over 5 years** for AMD AI chips, per Reuters reporting (also described as "parallel 6 gigawatts of AMD GPUs") [secondary — primary Reuters URL not directly fetched].
- **Nvidia:** remains Meta's primary GPU supplier; training compute reportedly on 500,000+ Blackwell B200-class GPUs [unverified]; Meta also rents large Google TPU clusters [secondary/unverified].
- **BlackRock JV (July 28, 2026):** 1-GW AI data-center campus in **El Paso, Texas**, ~$14B total cost; BlackRock funds hold 80% equity, Meta 20%; capacity online from 2028. Separate 1-GW facility in **Alberta** announced alongside [secondary].
- **Tulsa:** $1B AI-optimized data center announced (1,000+ construction jobs, ~100 permanent); water-efficient cooling, clean-energy matching [secondary].
- **Corning (Jan 2026):** multi-year **$6B** agreement for fiber-optic cables for data centers [secondary].
- **"Tents":** Meta experimenting with lightweight rapidly-deployable data-center structures — permanent facilities take 12–24 months vs. weeks for tents; signal is speed-to-capacity [secondary].
- Power strategy: direct gas-turbine deals, long-term nuclear PPAs, on-site battery storage; behind-the-meter generation toward GW-scale campuses [secondary].
- IDC (via secondary): global AI infrastructure spend projected **$497B in 2026 (+56% YoY)** [secondary].

### 4.12 Llama status — detailed (from Track B)

- **Last verified Llama-family release:** Llama 4 Scout / Maverick, April 2025. No Llama 4.x point release verified in Feb–Sep 2026.
- **Llama 5: UNVERIFIED — conflicting claims. Do not treat as a shipped product fact.** Several low-quality aggregator pages claim Llama 5 launched **April 8, 2026** at a "Meta AI Connect summit" with "600B+ parameters" and "5M-token context," citing a CNBC article whose actual title references the Muse Spark line, not a Llama launch [unverified]. A separate aggregator narrative ("Llama 5 Avocado," Aug–Sep 2026) says Meta's next open model project finished basic training in early 2026 but launch was pushed back [unverified]. A structured Meta-Llama tracking page (verification stamp 2026-09-18) states explicitly: *"No Llama 4.1, 4.5 or 5 exists on any first-party source"* [secondary].
- **AI pricing tracker (Aug 10, 2026 checkpoint):** Meta announced it will *resume* open-source model releases "soon" but published no model name, checkpoint, license, hardware target, or endpoint; hosted Llama API pricing remains third-party inference providers only — Meta publishes weights, not a first-party token price [secondary] https://www.aipricing.guru/meta-llama-pricing/.
- **Meta AI app / Reality Labs:** Meta AI app and meta.ai added "Thinking" mode powered by Muse Spark 1.1 (July 9, 2026) [vendor-reported/secondary]. No verified 2026 Reality Labs AI-specific launches found in fetched sources. **Meta Connect 2026 scheduled Sept 23–24, 2026** — after this report's cutoff; re-verify post-Connect [secondary].

