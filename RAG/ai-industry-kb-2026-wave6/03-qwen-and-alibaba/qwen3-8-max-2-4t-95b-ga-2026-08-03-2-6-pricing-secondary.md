---
id: ai-industry-kb-2026-wave6/03-qwen-and-alibaba/qwen3-8-max-2-4t-95b-ga-2026-08-03-2-6-pricing-secondary
title: "Qwen3.8-Max — 2.4T/95B, GA 2026-08-03, $2/$6 pricing [SECONDARY]"
domain: qwen-and-alibaba
role: deep-dive
task: pricing
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Moonshot", "Nvidia", "OpenAI", "United States"]
dates: ["2026-07", "2026-07-19", "2026-08-03", "2026-08-14", "2026-09-02"]
keywords: ["pricing", "qwen", "agent", "agents", "apache", "attribution", "benchmarks", "blackwell", "claude", "copyright", "cost", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1173, 1201]
section: "§3. Qwen and Alibaba"
delta_of: ai-industry-kb-2026
sha256: 6e825b7114fc10fdcd5f2753478db7d0490debe406a4abea76840cab9fe082cd
---

# Qwen3.8-Max — 2.4T/95B, GA 2026-08-03, $2/$6 pricing [SECONDARY]

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

