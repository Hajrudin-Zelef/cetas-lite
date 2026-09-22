---
id: ai-industry-kb-2026-wave6/23-master-model-inventory/overview
title: "§23. Master Model Inventory"
domain: master-model-inventory
role: deep-dive
task: reference
actors: ["Alibaba", "Anthropic", "Cohere", "DeepSeek", "Falcon", "Google", "Hugging Face", "LongCat", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Poolside", "SGLang", "Z.ai", "vLLM", "xAI"]
dates: ["2025-03-13", "2025-08", "2025-09-22", "2025-12", "2026-04", "2026-04-20", "2026-05", "2026-06-01", "2026-06-03", "2026-07-24", "2026-07-31", "2026-08-11", "2026-08-12", "2026-08-13", "2026-08-20", "2026-08-31", "2026-09-03", "2026-09-10", "2026-09-22", "2026-12-31"]
keywords: ["acquisition", "agents", "apache", "astra", "benchmark", "benchmarks", "claude", "cost", "deepseek", "embedding", "fable 5", "gemini"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10951, 11025]
section: "§23. Master Model Inventory"
sha256: 8be80214711422d4450448a3111fb274de9ab16e90aff46cd189c2d76fff772d
---

# §23. Master Model Inventory

Keywords: model inventory, DeepSeek, Qwen, GLM, Kimi, LongCat, MiniMax, MiMo, Llama, Mistral, Grok, GPT-6 Astra, Claude Fable, Gemini, Nemotron, Command A+, Falcon, Laguna, open weights

## Summary

This section is a condensed one-row-per-model master inventory of every model verified across wave6 sources 01–05. It adds no new facts — it condenses. Full per-model detail lives in §§2–17 (retained there by ownership rule); licensing analysis lives in §21; dated pricing and the benchmark landscape live in §22. Models verified as non-existent, unreleased, or contradicted are listed explicitly in the superseded/contradicted list below rather than as inventory rows. Key cross-source corrections are baked in: CED = "Causal Encoder-Decoder" [VENDOR]; DeepSeek V4-Lite 200B never shipped; V4-Pro-0813 released 2026-08-13; V4.1-Flash GA 2026-09-10; Qwen3.8-Max weights 2026-08-12 with the -0902 refresh; GLM-5.2 and 5.3 share one base (744B vendor headline / 753.3B HF tensor sum / ~743B vLLM rounding); Kimi K2.6 released 2026-04-20; LongCat-Flash-Thinking originally 2025-09-22/23; MiMo V2.5-Pro and V2.6-Pro are distinct generations sharing 1.02T/42B; MiniMax M3 released 2026-06-01; Mistral Large 3 is December 2025; NVIDIA licensed Poolside's Model Factory, not Laguna; Command A+ is Apache-2.0 open-weight, not fully open-source; Falcon was active in 2026; Gemma 4 12B launched 2026-06-03; Grok Code Fast 1 is August 2025; DeepSeek R2 remained unreleased as of 2026-09-22.

## Key dated facts

### Superseded, contradicted, or non-existent models

- **Gemini 2.0 Ultra** — never existed; no verified release record [CONTRADICTED].
- **Claude 4.62** — no trace in any source; treated as a data-entry error [CONTRADICTED].
- **DeepSeek V4-Lite 200B** — rumor/SEO noise; never shipped [CONTRADICTED].
- **DeepSeek R2** — NOT released as of 2026-09-22; "April 2026 32B MIT release" claims contradicted by multiple outlets [CONTRADICTED].
- **Grok 5** — not shipped by 2026-09-22 [CONTRADICTED as a release claim].
- **Nemotron 4 (1T)** — reported in development (2026-08-11); never announced or shipped; NVIDIA declined comment [SECONDARY].
- **Gemini 3.5 Pro** — internal/unreleased; no public release [UNVERIFIED as a release].
- **Gemini Omni Pro** — teased only; unverified [UNVERIFIED].
- **Runway Gen-5** — not shipped; flagship is Gen-4.5 (Dec 2025) [CONTRADICTED as a release claim].
- **ElevenLabs v4** — previewed (ElevenSummit Warsaw, Jun 2026); no GA, no endpoint as of Sept 2026; flagship is Eleven v3 [SECONDARY].
- **Phi-5** — no verified release found as of 2026-09-22.
- **Ministral 3 (2026)** — single-source only; do not merge with Ministral 3B/8B (Oct 2024) [UNVERIFIED].
- **qwen4_exp** — experimental/internal checkpoint (SGLang day-0, Aug 2026); NOT a Qwen 4 release [SECONDARY].
- "OpenAI shelved flagship Astra Aug 7" (Medium) — [UNVERIFIED]; conflicts with the documented GPT-6 Astra launch of 2026-09-03; not used as fact.
- "Compressed Expert Dispatch" as DeepSeek's CED — [CONTRADICTED]; CED = "Causal Encoder-Decoder" per DeepSeek's own HF README [VENDOR].
- "NVIDIA acquired Poolside" — [CONTRADICTED]; $6B non-exclusive Model Factory license + $1B equity (2026-08-20); "not an acquisition and not an acquihire" [SECONDARY].


### New verified facts — expansion

### OpenAI

- **GPT-6 Astra** (2026-09): flagship; DeepSWE 2026-09-22: 74%±3 at $4.43 (tied #1) [SECONDARY — official board]; public API price UNVERIFIED in this pass; status: current. Source: https://deepswe.datacurve.ai/
- **GPT-5.6 Sol** (2026): 400K context (1M tier); 272K whole-request repricing (2x input/1.5x output); price contradiction $4/$20 (Jul 2026) vs $5/$30 (Sep 2026) UNRESOLVED; DeepSWE 2026-09-22: 73%±3 at $6.46 [SECONDARY]; proprietary. Sources: https://deepswe.datacurve.ai/ and https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md
- **GPT-5.6 Terra** ($8/$24 per 1M, Jul 2026): Terminal-Bench 2.1 official board: 78.4%±1.3 at $421.15/run with Codex max [SECONDARY]; proprietary; status: current. Sources: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md and https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md
- **GPT-5.6 Luna** ($4/$12 per 1M, Jul 2026): Terminal-Bench 2.1: 75.7%±1.3 at $241.45/run with Codex max [SECONDARY]; proprietary; status: current. Sources: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md and https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md
- **GPT-5.5** (2026): DeepSWE May 2026: 70%±4; DeepSWE 2026-09-22: 67%±6 — dated snapshots, do not merge [SECONDARY]; Terminal-Bench 2.1: 83.1%±1.1 at $2,059.19/run with Codex xhigh [SECONDARY]; proprietary. Sources: https://deepswe.datacurve.ai/ and https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- **GPT-5.5 Pro** ($15/$120 per 1M, Jul 2026): highest-priced OpenAI tier documented; proprietary [SECONDARY]. Source: https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md
- **GPT-5.4**: 272K repricing family member; detailed specs not pulled in this pass — entry exists, fields UNVERIFIED [UNVERIFIED].
- **GPT-5.3-Codex / Codex model** ($12/$60 per 1M, Jul 2026): the Codex harness model; Terminal-Bench 2.1 scores are model+harness pairs [SECONDARY]. Source: https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md

### Anthropic

- **Claude Opus 5** (2026-07-24): $5/$25 per 1M; DeepSWE 2026-09-22: 74%±4 at $11.84 (max effort) [SECONDARY]; Terminal-Bench 2.1: Opus 4.8 (predecessor) 78.9%±1.3 with Claude Code high [SECONDARY]; proprietary; status: current flagship. Sources: https://deepswe.datacurve.ai/ and https://www.morphllm.com/claude-code-pricing
- **Claude Sonnet 5** (2026-07-24): price contradiction $2/$10 (intro, reported ended 2026-08-31 → $3/$15) vs $2/$10 permanent UNRESOLVED; proprietary [SECONDARY — contradiction]. Sources: https://www.morphllm.com/claude-code-pricing and https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- **Claude Fable 5** ($10/$50 per 1M): DeepSWE 2026-09-22: 70%±3 at $13.41 (xhigh) [SECONDARY]; Terminal-Bench 2.1: 83.8%±1.2 at $552.67/run with Claude Code xhigh [SECONDARY]; proprietary. Sources: https://deepswe.datacurve.ai/ and https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- **Claude Haiku 4.5** ($1/$5 per 1M): budget tier; launched alongside Opus 5/Sonnet 5 on 2026-07-24 [SECONDARY]; proprietary. Source: https://www.morphllm.com/claude-code-pricing
- **Claude Mythos 5 / 5.1**: named in corpus; specs, price, benchmarks UNVERIFIED in this pass [UNVERIFIED].
- **Claude Opus 4.8**: Terminal-Bench 2.1: 78.9%±1.3 with Claude Code high — the pre-Opus-5 reference point [SECONDARY]; superseded. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md

### Google

- **Gemini 3.8 Flash** (2026-09): four tiers — Standard $0.75/$0.075/$3.75, Pro $4.50/$0.45/$18.00 (2M context), Lite $0.30/$0.03/$1.20, Max $9/$0.90/$36; DeepSWE 2026-09-22: 74%±1 at $2.36 (high, tied #1) [SECONDARY]; proprietary; status: current price-performance leader. Sources: https://deepswe.datacurve.ai/ and https://github.com/aaronmarchant96-max/rei-ai/blob/HEAD/docs/CACHE_PRICING_LANDSCAPE.md
- **Gemini 3.7 Flash** (2026): intro pricing $0.75/$0.075/$3.75 through 2026-12-31, then $1.50/$0.15/$7.50; 50% introductory cut for coding/agents [SECONDARY]; proprietary. Sources: https://venturebeat.com/technology/googles-gemini-3-7-flash-targets-coding-and-agents-with-a-50-introductory-price-cut and https://devtk.ai/en/blog/gemini-api-pricing-guide-2026/
- **Gemini 3.6 Flash**: same intro/standard schedule as 3.7 Flash; the 3.6→3.7 delta is benchmark-measured in §19 [SECONDARY]; proprietary. Source: https://github.com/caudena/beam_weaver/blob/HEAD/docs/partners/google.md
- **Gemini 3.5 Flash**: older generation; $9/M output — do not confuse with 3.6/3.7 pricing [SECONDARY]; proprietary. Source: https://devtk.ai/en/blog/gemini-api-pricing-guide-2026/
- **Gemini 3.1 Pro / 2.5 Flash / 3.5 Flash-Lite**: named in corpus; 2026 pricing UNVERIFIED [UNVERIFIED].
- **Gemma 3** (2025-03-13): 1B (2T tokens/32K), 4B (4T/128K), 12B (12T/128K), 27B (14T/128K), 270M (6T/32K); Gemma Terms of Use; cutoff Aug 2024 [SECONDARY]. Source: https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_3.md
- **Gemma 4** (2026): gemma-4-31B-it and gemma-4-26B-A4B-it, Apache 2.0 — Google's license transition [COMMUNITY]. Source: https://github.com/ufal/atrium-project/issues/9
- **EmbeddingGemma-300M**: 69.67 MTEB-eng-v2 (<500M class top), Matryoshka 768/512/256/128, <200MB RAM via QAT; license UNVERIFIED [SECONDARY]. Sources: https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md and https://www.codesota.com/benchmarks/mteb

### xAI

- **Grok 4.5 / 4.6**: proprietary API-only; 2026 pricing UNVERIFIED in this pass [UNVERIFIED]. Source: https://www.aipricing.guru/xai-grok-pricing/
- **Grok Code Fast 1**: coding-specialized; specs UNVERIFIED [UNVERIFIED].

### DeepSeek

- **DeepSeek V4 Pro** (2026): MIT, weights on Hugging Face day one; DeepSWE 2026-09-22: 63%±6 [SECONDARY]; Terminal-Bench 2.1 (Ante harness): 65.8%±2.25 [SECONDARY]; price UNVERIFIED; status: current open flagship. Sources: https://deepswe.datacurve.ai/ and https://github.com/antigmalabs/ante/blob/HEAD/docs-site/docs/benchmarks/eval.mdx
- **DeepSeek V4 Flash** (2026): MIT; positioned as the cheaper/faster V4 [SECONDARY]; price UNVERIFIED. Source: https://kimi-k2.org/blog/56-kimi-k3-vs-deepseek-v4-flash-0731
- **DeepSeek V4.1 Flash** (2026-07-31, "0731"): MIT, 284B total/13B active, text; Terminal-Bench 2.1 (Ante): 83.9%, 370/445 trials, ~$18 total inference [SECONDARY]; DeepSWE 2026-09-22: 90.6% on TB 2.1 / 74.2 DeepSWE v1.1 per corpus [SECONDARY]; keep separate from generic V4 Flash. Sources: https://github.com/antigmalabs/ante/blob/HEAD/docs-site/docs/benchmarks/eval.mdx and https://kimi-k2.org/blog/56-kimi-k3-vs-deepseek-v4-flash-0731
- **DeepSeek V3**: prior generation, MIT; superseded by V4 line [SECONDARY]. Source: https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/

