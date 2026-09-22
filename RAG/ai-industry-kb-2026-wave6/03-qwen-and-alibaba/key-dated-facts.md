---
id: ai-industry-kb-2026-wave6/03-qwen-and-alibaba/key-dated-facts
title: "Key dated facts"
domain: qwen-and-alibaba
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "ByteDance", "California", "China", "DeepSeek", "Google", "Hugging Face", "Meta", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "Together AI", "United States", "Z.ai"]
dates: ["2025-06", "2026-02-03", "2026-02-15", "2026-02-16", "2026-02-24", "2026-03-30", "2026-03-31", "2026-04", "2026-04-02", "2026-04-16", "2026-04-20", "2026-04-22", "2026-05", "2026-05-18", "2026-05-20", "2026-06", "2026-06-01", "2026-06-30", "2026-07", "2026-07-15", "2026-07-19", "2026-07-21", "2026-08", "2026-08-03", "2026-08-05", "2026-08-12", "2026-08-14", "2026-08-20", "2026-08-26", "2026-09-02", "2026-09-17", "2026-09-22"]
keywords: ["agent", "agentic", "agents", "alignment", "apache", "attention", "attribution", "awq", "benchmark", "benchmarks", "blackwell", "capex"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1096, 1257]
section: "§3. Qwen and Alibaba"
sha256: 17ebfb57829984889f5b56602b1b9553a56aa2bcf6159d8b7ee90284633e31cb
---

# Key dated facts

## Key dated facts

### Qwen3.5 line
- **2026-02-16** — Qwen3.5-397B-A17B released (open, 397B total / 17B active MoE). [SECONDARY]
- **2026-02-24/25** — Medium open Qwen3.5 models released alongside hosted **Qwen3.5-Flash**. [SECONDARY]
- **2026-03-30** — Qwen3.5-Omni released: text/audio/video omni-capable line, multilingual and captioning upgrades. [SECONDARY]

### Qwen3.6 line (April 2026)
- **2026-04-02** — Qwen3.6-Plus released (closed/API tier). [SECONDARY]
- **2026-04-16** — Qwen3.6-35B-A3B released (open, 35B total / ~3B active). [SECONDARY]
- **2026-04-20** — Qwen3.6-Max-Preview released (closed). [SECONDARY]
- **2026-04-22** — Qwen3.6-27B released (open). [SECONDARY]

### Qwen3.7 line
- **2026-05-20** — Qwen3.7-Max released (closed). [SECONDARY]
- **2026-06-01** — Qwen3.7-Plus released (closed). [SECONDARY]
- **July 2026** — Qwen3.7-Flash: July snapshot/release cadence. [SECONDARY]
- The vendor plan that Qwen3.7-Plus "will be open source" **remained unfulfilled in the observation window** — a stated intent, not a release. [VENDOR]

### Qwen3.8 line
- **2026-08-26** — Qwen3.8-Flash-Next released (125B total / 6B active). [SECONDARY]
- **2026-09-02** — Qwen3.8-Max-0902: a dated **snapshot/checkpoint of Qwen3.8-Max**, not a new base model. Reported at AA Intelligence Index v4.3 = 45 [SECONDARY, r/LocalLLaMA via alextech, 2026-09-17], reclaiming the China lead on that board — checkpoint-dated, not methodology-independent. [SECONDARY]
- Qwen3.8-Max API pricing reported at $2.00/$6.00 per M input/output (2026-08-12); pricing discussion is intentionally shallow here — the section's license terms are the durable fact. [SECONDARY]

### Qwen3-Coder-Next (2026-02-03/04)
- **2026-02-03/04** — Qwen3-Coder-Next released: 80B total / ~3B active, **Apache 2.0**, Gated DeltaNet hybrid attention, 262K context, trained on **800K+ verifiable coding tasks**. [VENDOR]
- The Gated DeltaNet hybrid architecture and the 800K-task coding corpus are the differentiators against the general Qwen3.6 line. [VENDOR]

### Licensing and the Max tier
- Qwen3 / 3.5 / 3.6 / 3.8 (up to 27B-class) are **Apache 2.0** [VENDOR]; Qwen3.8-Max ships under a **custom license**: name-display required above 100M MAU or $20M monthly revenue; MaaS/AI-assistant businesses above $50M trailing revenue need a separate commercial license; internal use is exempt. [VENDOR]
- The custom Max license is the "flagship has the most strings" case inside Alibaba's own line (see wave6/05 Part 5-A1 for the lab-by-lab map). [DIRECTIONAL]

### qwen4_exp (August 2026)
- `qwen4_exp` appeared as an **experimental checkpoint (SGLang day-0)** — community-tracked, not vendor-announced. [SECONDARY]
- Alibaba's stated position: **Qwen 4 was still in training**; `qwen4_exp` is not a Qwen 4 release and must not be cited as one. [VENDOR]

### Qwen3.8-Max benchmark and adoption (secondary context)
- Qwen3.8-Max DeepSWE: **56.6 [VENDOR] vs 69.3 [SECONDARY]** — different harnesses; provenance-tagged, not averaged. The corpus's standing rule: never merge vendor-harness and third-party-harness numbers. [VENDOR/SECONDARY]
- Adoption: **Reuters, Airbnb, and Pinterest** have adopted Qwen models [SECONDARY]; a community "Global LLM Download Leaderboard" (HF, Sept 2026) shows **Qwen3-0.6B at 22.7M cumulative downloads** — small models dominate raw counts. [COMMUNITY]
- Qwen3.7-Plus API pricing reported at **$0.276/M input** (Sept 2026, Global ≤256K non-thinking) — dated, and cheap enough to sit in the bottom quartile of the corpus price table. [SECONDARY]


### New verified facts — expansion

### Qwen3.5-397B-A17B — 2026-02-15/16, Apache 2.0, hybrid Gated DeltaNet MoE [SECONDARY]
- **CONTRADICTION**: release date reported as **2026-02-15** (Puter developer docs) vs **2026-02-16** (Medium/dig.watch launch coverage). The model card is the tiebreaker to fetch; the corpus should use 2026-02-15/16 as a range until then. [SECONDARY]
- Architecture (per a Hugging Face community card transcribing the spec): **397B total / 17B active**, **512 routed experts with 10 routed + 1 shared active per token**, 60 layers in the pattern `15 × (3 × (Gated DeltaNet → MoE) → 1 × (Gated Attention → MoE))`, Gated DeltaNet with 64 V/16 QK linear-attention heads (dim 128), Gated Attention with 32 Q / 2 KV heads (dim 256, RoPE-64), expert intermediate dim 1024, embedding size 248320 (padded — consistent with the reported ~250k vocabulary). **Multi-step MTP** trained. [SECONDARY]
- Context: **262,144 natively, extensible to 1,010,000 tokens** — the 1M-scale extensibility is native to the architecture, not a serving hack. [SECONDARY]
- Positioning: **native vision-language foundation model via early-fusion training** on multimodal tokens; claimed **cross-generational parity with Qwen3** while beating Qwen3-VL on reasoning, coding, agents, and visual understanding; **8.6–19× faster decoding than Qwen3-Max** via the Gated DeltaNet + sparse MoE stack; **201 languages and dialects** (up from Qwen3's 119); **Apache 2.0** license with weights on Hugging Face and ModelScope. [SECONDARY]
- Hosted tier: **Qwen3.5-Plus** on Alibaba Cloud Model Studio — 1M context, built-in adaptive tool use. [SECONDARY]
- Vendor-claimed benchmarks (from the HF card, label as vendor claims): MMLU-Pro **87.8**, MMLU-Redux **94.9**, SuperGPQA **70.4**, C-Eval **93.0**, IFEval **92.6**, IFBench **76.5**, MultiChallenge **67.6**, GPQA **88.4**, HLE **28.7**, HLE-Verified **37.6**, LiveCodeBench v6 **83.6**, HMMT Feb 25 **94.8**, HMMT Nov 25 **92.7**, AIME26 **91.3**, BFCL-V4 **72.9**, TAU2-Bench **86.7**, MCP-Mark **46.1**, BrowseComp **69.0/78.6**, BrowseComp-zh **70.3**, WideSearch **74.0**, Seal-0 **46.9**, MMMLU **88.5**. The card tabulates Qwen3.5-397B-A17B against GPT-5.2, Claude 4.5 Opus, Gemini 3 Pro, Qwen3-Max-Thinking, and K2.5-1T-A32B. [SECONDARY]
- Community serving: on **4× RTX PRO 6000 (96GB)**, a local-inference recipe reports **259 tok/s single-user** (AWQ + b12x decode + EAGLE MTP), 180 tok/s NVFP4 + MTP, and 1550 tok/s at 64 concurrent users — evidence the 397B/17B footprint is servable on prosumer 4-GPU hardware. [COMMUNITY]
- **Release mechanics**: announced **2026-02-16** via the Qwen blog ("Qwen3.5: Towards Native Multimodal Agents"); GitHub QwenLM/Qwen3.5 news entry dated **2026-02-16** confirms "The first release includes a 397B-A17B MoE model." Third-party trackers: OpenRouter lists **February 16**, Artificial Analysis lists **February 17** — the 02-15/02-16 contradiction in the corpus narrows to a **02-16/02-17** range with the official blog as tiebreaker (02-16). [SECONDARY]
- **Pricing** (third-party): **$0.60 input / $3.60 output** per 1M (Together AI, Novita, OpenRouter); Puter docs list **Feb 15, 2026** release date (the 02-15 source). Max output **64–66K tokens**. Input: text/image/video; output: text. Function calling + structured output supported. [SECONDARY]
- **Competitive frame** (sci-tech-today): vs **Kimi K2.5** (1T/32B, $0.60/$3.00, Modified MIT, AA 46) vs **DeepSeek V3.2** (685B/37B, $0.28/$0.42, text-only) — Qwen3.5 at **AA 45**, the only one with **native vision-language + 201 languages + Apache 2.0**. [SECONDARY]
- **Architecture detail** (per SemiAnalysis inferencex-app model card): **3:1 Gated DeltaNet to Gated Attention hybrid stack**, **262K native context**, **Apache 2.0 weights**. [SECONDARY]
- **Qwen3-VL-235B-A22B** (companion vision model): **8.6× faster** than Qwen3-Max at 32K tokens, **19× faster** at 256K tokens — the efficiency story behind the native vision-language claim. [SECONDARY]

### Qwen3.6 — first closed-weights flagship, April 2026 [SECONDARY]
- **2026-04-20** — **Qwen3.6-Max-Preview** released: the first Qwen flagship to ship **closed-weights only**, breaking the Apache-2.0 pattern. Vendor-claimed top ranks on **six coding/agent benchmarks**: SWE-Bench Pro, Terminal-Bench 2.0, SkillsBench, QwenClawBench, QwenWebBench, SciCode. Artificial Analysis Intelligence Index: **52**. Specs: **260K context**, OpenAI + Anthropic API compatible, **`preserve_thinking`** for multi-turn agents. [SECONDARY]
- **Qwen3.6-35B-A3B FP8** — the **first open-weight release in the 3.6 series**: 35B total / **3B active**, Gated DeltaNet + gated-attention hybrid, vision input, **262K native context extensible to 1M**, `preserve_thinking` option, served in FP8 on Together AI. Benchmarks (Together AI model card): **SWE-bench Verified 73.4%**, LiveCodeBench v6 80.4%, Terminal-Bench 2.0 51.5%, MMMU 81.7%, MathVista (mini) 86.4%, GPQA Diamond 84.1%, HLE 20%, SciCode 36%, Terminal-Bench 2.1 45%. [SECONDARY]
- **Qwen3.6-Plus Preview** (2026-03-31): agentic-coding + reasoning flagship with 1M context — the March 30/31 double drop with Qwen3.5-Omni. [SECONDARY]

### Qwen3.7-Max — May 2026, 1M context, 35-hour agent tasks [SECONDARY]
- **2026-05-18** — Qwen3.7-Max-Preview and Qwen3.7-Plus-Preview land on Arena's leaderboard (Qwen's own X announcement: Alibaba now **#6 lab in text, #5 in vision**). **2026-05-20** — **Qwen3.7-Max** officially released at the **2026 Apsara Conference**. [SECONDARY]
- Agent positioning: designed for the agent era; claimed ability to **fully autonomously complete ultra-long-duration complex agent tasks lasting up to 35 hours**. [SECONDARY]
- Vendor-claimed benchmarks (marktechpost/wedoany compilations, label as vendor claims): **SWE-Verified 80.4** (vs Claude Opus-4.6 Max 80.8, DeepSeek-v4-Pro Max 80.6), **Terminal-Bench 2.0-Terminus 69.7** (vs DeepSeek-v4-pro-Max 67.9), **GPQA Diamond 92.4** (vs Opus-4.6 91.3), **HLE 41.4** (vs Opus-4.6 40.0), **MCP-Mark 60.8** (vs GLM-5.1 57.5), **MCP-Atlas 76.4** (vs Opus-4.6 75.8), **SpreadSheetBench-v1 87** (top tier), WMT24++ **85.8**, MAXIFE **89.2**. [SECONDARY]
- Independent: **Artificial Analysis Intelligence Index v4.0: 56.6** — +4.8 over Qwen3.6-Max-Preview (51.8), ahead of Gemini 3.5 Flash (55.3), behind GPT-5.5 (60.2), Claude Opus 4.7 (57.3), Gemini 3.1 Pro Preview (57.2). [SECONDARY]
- Text Arena: **#13 overall, Elo 1475**; category ranks #7 Math, #9 Expert Prompts, #9 Software/IT, #10 Coding. One careful-reading flag: AA-Omniscience raw accuracy dropped 7.6 points (37.7→30.1%) while hallucination rate fell 21.3 points (44.2→22.9%) — the model refuses more (attempt rate 67.3→48.0%, lowest among frontier models compared) rather than recalling more. [SECONDARY]
- YC-Bench (simulated year of startup operations): Qwen3.7-Max generated **$2.08M simulated revenue** vs $1.05M (Qwen3.6-Plus) vs $352K (Qwen3.5-Plus) — 2× and 5.9× generational jumps. Treat as vendor-selected simulation, not enterprise ROI. [SECONDARY]
- Licensing posture: **Qwen3.7-Plus to be open-sourced; Qwen3.7-Max stays proprietary** — the same two-tier playbook as 3.8. [SECONDARY]
- **Alibaba Cloud Summit 2026 (May 20, Hangzhou)**: Qwen3.7-Max launched alongside the **Zhenwu M890 AI chip** (3× predecessor performance) and the **Panjiu AL128 Supernode Server** (128 accelerators, PB/s internal bandwidth) — a full-stack reveal from silicon to model to cloud. Liu Weiguang (SVP Alibaba Cloud): "What we're building is China's AI factory." [SECONDARY]
- **The 35-hour agent demo**: Qwen3.7-Max was given a task brief on a **Zhenwu M890 chip it had never encountered in training**; working without human intervention, it ran **35 consecutive hours**, executed **>1,000 tool calls**, and delivered a **production-grade AI computing kernel outperforming the chip manufacturer's official version by 10×**. Optimized for OpenClaw, Hermes Agent, Claude Code, Qwen Paw, Qoder. [SECONDARY]
- **Pricing** (May 2026): Model Studio **¥12/¥36 per 1M** (~$1.71/$5.14); OpenRouter **$2.50/$7.50**; cache input **90% off** (¥1.2 / $0.25). **1M context** (doubled from 256K), **65,536 max output**. OpenAI-compatible endpoint (`qwen3.7-max`) + **native Anthropic Messages Protocol** (Claude Code & OpenClaw via `ANTHROPIC_BASE_URL` swap). Two API key types: `sk-` and `sk-sp-` (Token Plan). [SECONDARY]
- **Qianwen App integration** (May 22): Qwen3.7-Max in Qianwen App v6.9.7+, PC, Web — **free for all users**. [SECONDARY]
- **Policy shift context**: Qwen3.6-Max-Preview (April) and Qwen3.7-Max (May) are **API-only, no open weights** — Alibaba's first sustained turn toward the "frontier-closed, smaller-open" strategy. Alibaba killed the **free tier of Qwen Code** the prior month. [SECONDARY]

### Qwen3.8-Max — 2.4T/95B, GA 2026-08-03, $2/$6 pricing [SECONDARY]
- **July 2026** — previewed at the **World AI Conference**; preview began **2026-07-19**. **2026-08-03** — **general availability**. [SECONDARY]
- Spec: **2.4T total / ~95B active per token** MoE, **1M-token context** (991K max input, 983K with thinking; opper.ai lists 984K), **131K max output**, **262K reasoning budget**, text/image/video input, hybrid thinking mode (step-by-step or direct), native tool calling, structured output. [SECONDARY]
- Pricing (international): **$2.00 input / $6.00 output** per million, flat with no context-length tiers; implicit cache reads **$0.25**, explicit cache creation **$2.50**, explicit cache reads **$0.17**. China price list: **¥12 / ¥36** — ~20% below international, not a direct FX conversion. Rate limits: 2M tokens/min, 15K requests/min. [SECONDARY]
- Five built-in tools on the Responses API: **code_interpreter, web_search, web_extractor, t2i_search, i2i_search**. [SECONDARY]
- Vendor-claimed benchmarks (from Alibaba's published table): **Terminal-Bench 2.1 86.6** (vs Opus 4.8 / Fable 5 84.6; behind GPT-5.6 Sol-max 88.8), **SWE-bench Pro 67.7** (vs Fable 5 80.0), **FrontierSWE 73.5** (vs Fable 5 88.8), **PaperBench 93.0** (leads), **IFBench 82.8** (leads), **GPQA Diamond 92.6** (up marginally from Qwen3.7-Max's 92.4), **DeepSWE 1.1 56.6** (vs Qwen3.7-Max 21.6 — the clearest generational coding jump), **OSWorld-Verified 86.1**, **Parametric CAD Bench 91.5**, **OmniDocBench 1.5 92.1** (multimodal/vision rows lead). [SECONDARY]
- **Methodological caveats** worth keeping: the multimodal comparison table benchmarks against **Qwen3.7-Plus, not Qwen3.7-Max** (flatters the generational delta); Alibaba's own RL scaling curve **peaks at 0.725 near 4,000 training environments then declines** to 0.719 and 0.689 — diminishing returns at RL scale. [SECONDARY]
- Independent: **Artificial Analysis Intelligence Index v4.3: 40 (Aug) → 45 (0902)** — but cost per task rose **$2.67 → $5.41** (output tokens per task 63K → 108K; reasoning tokens 71K alone now exceed the prior version's total). Code Arena WebDev: **1691, rank #1** (0902), edging Claude Opus 5 Max (1687) — the 3-point gap is inside the noise band; the correct read is **parity at ~4× lower price** ($6 vs $25 output). Arena text: **5th, Elo 1496**. [SECONDARY]
- **Qwen3.8-Max-0902** snapshot (2026-09-02): same architecture and pricing, post-training focused on coding and long-horizon agents. [SECONDARY]
- **Open weights Qwen3.8-2.4T-A95B**: announced ~2026-08-14 (Qwen X announcement; "for the first time, Qwen3.8 brings a Qwen-Max-class model to open release"). Artificial Analysis Index **v4.1.1: 58.1 (paid Max) vs 57.7 (open)** — functionally indistinguishable on measured intelligence; cost per index task $1.13 vs $1.09; median 45 tok/s both. The paid service adds: **vision input & non-thinking support, 1M context by default, official built-in tools** (open model served as text-input only, 984K context). [SECONDARY]
- **CONTRADICTION / timing note**: one launch-week blog reported open weights "ship the week of August 10," while an independent August write-up found **no repository, license, or model card** at that time. The weights did materialize (~Aug 14) under a custom license — the sequence was announcement → verification gap → release, not a same-day drop. [SECONDARY]
- **GA launch details** (August 3, 2026): announced at **Hangzhou**; **2.4T parameters**, **1M context**, **5th in Text Arena, 2nd in Vision Arena**. Available via **Alibaba Cloud Model Studio**; also on **QwenWork** (Alibaba's all-in-one workplace AI agent platform). [SECONDARY]
- **Agent flagship positioning**: built on Qwen3.7-Max architecture, scaled up and extended to **multimodal input for the first time above the trillion-parameter mark**. In one test: **>10 days autonomously coding a self-evolving software harness** from scratch (user feedback, own tests, iterating code/previews/logs). In another: **reproduced an ML research paper from zero** — 33 rounds of GPU training over ~125 hours, 7,600 lines of code, then **18 improvement ideas outperforming the original paper**. Entered a **real online contest with 526 human teams: beat 87%**. [SECONDARY]
- **Competitive frame**: vs **Kimi K3** (Moonshot, 2.8T, launched prior month) — both text/image/video, 1M context. Qwen3.8-Max became the **highest-ranking Chinese model for text** on Arena.AI; trails Claude Fable 5 and three Opus variants. Vision: **2nd globally**, behind one Fable 5 variant. [SECONDARY]
- **Launch-day coincidence**: Qwen3.8-Max GA (Aug 3) landed the **same day as DeepSeek V4-Flash** — the two Chinese labs shipping flagships simultaneously. [SECONDARY]
- **Hardware floor**: cheapest usable quant ~**450GB combined RAM+VRAM**; lossless BF16 ~**4.9TB**; Alibaba's reference deployment: **72 Blackwell Ultra GPUs**. The **Qwen3.8-27B** sibling (~14–16GB VRAM at 4-bit, RTX 4090-runnable) is the developer-relevant release. [SECONDARY]

### Qwen3.8 licensing — exact custom-license terms [SECONDARY]
- The repo license field reads **`qwen3.8-max`**, not Apache-2.0: a custom Qwen-authored document titled **"Qwen3.8-Max License," Copyright 2026 Qwen**. Qwen3.8-Flash-Next ships under the **Qwen Community License Agreement 1.0** with the same operative thresholds. [SECONDARY]
- Clause-by-clause (per two independent license analyses and the HF license text they quote):
  1. **Attribution trigger** — if a commercial product/service exceeds **100 million monthly active users OR US$20 million in monthly revenue**, the model name must be **prominently displayed in the UI**. A branding requirement, not a payment. [SECONDARY]
  2. **Separate-license trigger** — if you or affiliates operate a **"Model as a Service" or "AI Work Assistant" business** with **aggregate revenue exceeding US$50 million over any consecutive 12 months**, you must **obtain a separate license from Qwen before commercial use**. [SECONDARY]
  3. **Carve-out** — **internal use is exempt**, provided the model, its outputs, and its capabilities are **not exposed to third parties**. [SECONDARY]
- **Notably absent**: no explicit restriction on using the model's outputs to train other models — a real divergence from several competing custom licenses. Also: **no revenue-share percentage anywhere** — it is a usage gate (negotiate a license past the threshold), not a royalty. [SECONDARY]
- **License comparison**: Qwen3.5-397B-A17B (Apache 2.0) vs Qwen3.8-Max (custom qwen3.8-max) — the same lab, six months apart, opposite licensing. The **permissiveness cycle** is the story: open when distributing, gated when monetizing. [DIRECTIONAL]
- **Qwen3.8-27B** (Apache 2.0, ~14–16GB VRAM at 4-bit): the **developer-relevant** open release — runs on RTX 4090, unlike the 2.4T flagship (450GB+ floor). [SECONDARY]
- Two-tier split: **Qwen3.8-27B is Apache 2.0**; the **2.4T flagship is Qwen3.8-Max License**. And the historical arc: Qwen-72B and Qwen2-72B shipped under near-identical custom terms (the same 100M-MAU trigger; Qwen2 went further with a non-compete clause), then **all of Qwen3 shipped Apache 2.0**, and Qwen3.8-Max walks it back. [SECONDARY]
- Community caution: derivative cards sometimes tag themselves Apache 2.0 because their **recipe/delta** is Apache-licensed — that metadata does **not** relicense the underlying Qwen-derived weights. [SECONDARY]

### Qwen Image 3.0 / 3.0 Pro — 2026-07-21, no benchmarks, no weights [SECONDARY]
- **CONTRADICTION**: release date reported as **2026-07-21** (majority of launch-week coverage) vs **2026-08-05** (siray.ai). The July 21 date is better attested. [SECONDARY]
- Headline spec: **4,500-token prompt input** (up from ~1,000 on Qwen-Image-2.0), vendor-demonstrated **text legibility down to 10 pixels**, **native support for 12 languages**, LaTeX math notation, single-pass dense layouts (nine-panel infographic grids, nested UI mockups, full academic-paper pages, simulated newspaper front pages), and live-data capability (e.g., a weather graphic for a specific city and date). [SECONDARY]
- **Transparency reversal**: no benchmark scores, no model card, no technical report, no parameter count, **no downloadable weights** — a sharp break from Qwen-Image 1.0 (Aug 2025, Apache-2.0 open weights + same-day technical report) and Qwen-Image 2.0 (May 2026, report + Alibaba's own eval leaderboard). Access via Qwen Chat and Alibaba Cloud APIs; reported invite-only at launch with no public API initially. [SECONDARY]
- Lineage context (predates 3.0, do not attribute to it): third-party trackers reported the Qwen-Image line at **88.32 GenEval, 83.84 GSO image-editing, 58.30 LongText-Bench English** (Jan 2026); Alibaba's own **Qwen-Image-Bench placed Qwen Image 2.0 Pro 5th overall**, behind OpenAI and Google models. [SECONDARY]
- Independent hands-on tests (secondary): strengths in dense text rendering and reference-based editing; gaps in landmark and historical-figure accuracy. [SECONDARY]
- **Positioning keyword**: 1.0 = "Precision"; 2.0 = "Precision, Variety, Completeness, Beauty, Authenticity"; 3.0 = **"Real"** (实) — organized into three pillars: **Rich Content, Authentic Details, Deep Knowledge**. Goal: "useful" rather than "good-looking" — deployable productivity tool for documents, interfaces, knowledge-dense visuals. [SECONDARY]
- **Tiers**: `qwen-image-3.0` (standard), `qwen-image-3.0-pro` (flagship). Resolution **512×512 to 2048×2048** (2K PNG). Reference images (editing): **1–3 per request**. Pricing: **from $0.03/image** (high-res); input $0.003/image; Pro i2i **$0.075/generation** (ModelsLab). Availability: Model Studio, Qwen Cloud, kie.ai, OpenArt, ModelsLab. [SECONDARY]
- **Launch sequence**: shipped **2026-07-21** behind invite-only wall; **opened to all Qwen AI platform users 2026-08-05** (OrcaRouter GA tracking). The 07-21/08-05 "contradiction" is **invite-only vs GA**, not a date conflict. [SECONDARY]
- **2.0 efficiency context**: Qwen-Image-2.0 (May 2026) shipped a 4-step variant (down from 40 steps); on Alibaba's arena, 2.0 landed just behind **GPT-Image-2** and **Nano Banana Pro**. [SECONDARY]

### Qwen3.5-Omni — 2026-03-30 [SECONDARY]
- Released **2026-03-30**: multimodal model accepting **text + image + audio + video**, **113 languages**, with "Vibe Coding" capability. [SECONDARY]
- **Architecture**: **Thinker-Talker + Hybrid-Attention MoE** (both Thinker and Talker MoE); **256K context**; **>10 hours continuous audio**; **400+ seconds 720p video** (1 FPS); trained on **>100M hours audio/video data**. [SECONDARY]
- **Variants**: **Plus** (30B-A3B MoE), **Flash** (lightweight MoE), **Light** (dense, open weights). **Qwen3.5-Omni-Plus-Realtime** for streaming. [SECONDARY]
- **Speech**: **113 languages/dialects recognition** (up from 19), **36 languages generation** (up from 10); multi-codebook codec for single-frame synthesis; **ARIA** (dynamic text-speech alignment in streaming decoding). [SECONDARY]
- **Benchmarks**: **215 SOTA** across audio/video understanding (vendor claim — treat with skepticism; self-selected benchmarks favor the releasing lab); Plus **surpasses Gemini 3.1 Pro on audio understanding**, matches on audio-visual. [SECONDARY]
- **Audio-Visual Vibe Coding**: camera + speech/gesture → functional website/game — "building through natural language and visual descriptions." [SECONDARY]
- **Availability**: Qwen Chat, **open weights on Hugging Face/ModelScope**, Alibaba Cloud API (Offline + Realtime). Technical report: arXiv:2604.15804. [SECONDARY]
- **Release cadence note**: Feb (Qwen3.5 397B, Qwen3-Coder-480B, Qwen3-Coder-Next) → Mar 30 (Qwen3.5-Omni) → Mar 31 (Qwen3.6-Plus Preview) — **five major releases in two months**; the Mar 30–31 back-to-back is "wartime shipping." [SECONDARY]

### Regulatory + pricing events — July 2026 [SECONDARY]
- **2026-07-15**: Alibaba's **Qwen app and ByteDance's Doubao simultaneously shut down user-created/custom AI agent features** as China's **'humanlike AI interaction services'** (anthropomorphic-AI/companion) rules took effect. **Qwen offered no migration path** — agent configurations and conversation histories **permanently deleted** (Doubao offered read-only access until Oct 15). Consumer-app change; models/API unaffected. [SECONDARY]
- **Early July 2026** (SCMP): Alibaba cut **Qwen3.7-Max by ~80%** and **Qwen3.7-Plus by ~60%** for international users on **Qoder** during off-peak hours (10pm–8am Beijing = US working hours) — an explicit play for US developer demand amid the Chinese token price war. [SECONDARY]

### Qwen3-Coder line — 480B/30B, SWE-bench Verified 69.6% [SECONDARY]
- **Qwen3-Coder-480B-A35B-Instruct**: 480B MoE / **35B active**, **256K context natively (1M with extrapolation)**, **Apache 2.0**, vendor claim of open-model SOTA on Agentic Coding, Agentic Browser-Use, Agentic Tool-Use, "comparable to Claude Sonnet 4." **SWE-bench Verified: 69.6% (348/500)** via OpenHands (Aug 2025 submission) — per-repo breakdown: django 165/231 (71.4%), sympy 53/75 (70.7%), sphinx 30/44 (68.2%), scikit-learn 27/32 (84.4%), pytest 15/19 (79.0%), xarray 17/22 (77.3%), requests 6/8 (75.0%), matplotlib 20/34 (58.8%), astropy 10/22 (45.5%), pylint 3/10 (30.0%). [SECONDARY]
- Lowest provider pricing (llm-stats.com, 2026-09-22): **$0.30 input / $1.00 output** (Deepinfra). [SECONDARY]
- **Qwen3-Coder-30B-A3B-Instruct**: SWE-bench Multilingual **~33.3–33.8%** (OpenHands/SWE-agent, NVIDIA NeMo skills eval). [SECONDARY]
- **Qwen3-Coder-Next**: smaller/faster hybrid architecture, reported **SWE-bench >70%** at much lower inference cost. [SECONDARY]
- **Training**: **7.5T tokens, 70% code, 358 programming languages**; post-training: SFT + RL (Code RL: execution-driven; Long-Horizon Agent RL: multi-turn planning/tool use); **Constitutional AI** for safety. Qwen built a **20,000-environment system on Alibaba Cloud** for agentic evaluation/training. [SECONDARY]
- **Qwen Code**: open-sourced CLI (forked from Gemini Code), Node.js/npm; integrates with **Claude Code** (DashScope proxy), **Cline**, Ollama, LMStudio, MLX-LM, llama.cpp, KTransformers. [SECONDARY]
- **Together AI pricing**: **$2.00/$2.00** per 1M input/output (480B). [SECONDARY]
- **SWE-bench harness note**: the 69.6% (348/500) is **OpenHands, 500-turn**; standard (non-500-turn) reported as **67.0%** — the harness/turn budget matters, never mix. [SECONDARY]

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

