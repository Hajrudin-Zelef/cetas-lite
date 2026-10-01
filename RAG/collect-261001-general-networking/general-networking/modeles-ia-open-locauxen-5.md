---
id: collect-261001-general-networking/general-networking/modeles-ia-open-locauxen-5
title: "ÉTAPE 1 — Open / Local AI Models (EN)"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "China", "DeepSeek", "Google", "Hugging Face", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Poolside", "Xiaomi", "xAI"]
dates: ["2025-04", "2025-04-05", "2026-04-08", "2026-07-09", "2026-08-05", "2026-08-10", "2026-09", "2026-09-02", "2026-09-08", "2026-09-22"]
keywords: ["agent", "apache", "attention", "benchmark", "benchmarks", "consumer", "context window", "deepseek", "fp4", "fp8", "gemini", "glm"]
source: docs/RAG/collect-261001-general-networking/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [234, 337]
sha256: 01d012a6e6bd4cef62548aa5b8394bb9c6dee7d1f16eee4a4046d294343f6d7c
---

# ÉTAPE 1 — Open / Local AI Models (EN)

## CAVEATS
- All vendor-reported benchmarks are self-reported unless labeled Artificial Analysis / independent harness. MiMo V2.6 figures were <24h old at research time.
- AA Index v4.1.x vs v4.3/v4.3.2 numbers are NOT comparable across versions — always record the version.
- Qwen3.8-Max open weights: weights confirmed shipped Aug 12, 2026 (`Qwen/Qwen3.8-2.4T-A95B`) but under custom license, text-only, reduced context.
- DeepSeek R2: unreleased/vaporware as of Sept 2026.
- MiniMax M2.5 license: HF frontmatter ("modified-mit") disagrees with the repo license file ("MINIMAX MODEL LICENSE") — trust the file.
- MiMo "$0.20/$0.70" pricing from the brief could not be verified — likely confusion with cache-read pricing or Grok 4.1 Fast.
- Kimi K3 "~104B active/token" appears in one source only — unverified; 896 experts / 16 active is multiply confirmed.
- Jev is American and proprietary — filed separately from the Chinese open track.

## KEY SOURCES
- Artificial Analysis: https://artificialanalysis.ai
- BenchLM.ai: https://benchlm.ai
- OpenRouter: https://openrouter.ai
- CheapestInference (State of Open Weights) · ChinaAPI / China AI Index
- https://venturebeat.com/technology/better-than-deepseek-xiaomis-mimo-v2-6-pro-debuts-as-the-top-open-weights-model-in-the-world-alongside-cheaper-v2-6-flash
- https://theoutpost.ai/news-story/xiaomi-mi-mo-v2-6-pro-debuts-as-top-open-source-ai-model-scoring-46-on-intelligence-index-31152/
- https://officechai.com/ai/xiaomi-mimo-v-2-6-pro-benchmarks/
- https://www.intelligentliving.co/qwen3-8-max-update-smarter/
- https://www.intelligentliving.co/xiaomi-mimo-v2-6-pro-cheapest-frontier/
- https://github.com/sampraszheng/yxz/blob/HEAD/wiki/synthesis/open-weight-llm-agent-stack-six-region.md
- https://www.swfte.com/ai/leaderboard
- https://en.wikipedia.org/wiki/Xiaomi_MiMo
- https://en.wikipedia.org/wiki/Jev_(AI_model)
- https://jev-agent.com/
- https://www.alibabacloud.com/en/press-room/alibaba-unveils-qwen3-8-max
- https://www.alizila.com/alibaba-unveils-qwen3-8-max-most-capable-flagship-model-to-date/
- https://explainx.ai/blog/qwen3-8-max-open-weights-live-hugging-face-august-2026
- https://www.siliconreport.com/alibaba-launches-qwen3-6-27b-and-says-its-dense-coding-model-tops-qwen3-5-397b-a17b-d15d82b7347c62a3
- https://www.morphllm.com/deepseek-v4
- https://www.orcarouter.ai/blog/deepseek-v4-1-flash-openrouter
- https://dataconomy.com/2026/09/11/deepseek-v4-1-flash-ultralow-token-pricing/
- https://theplanettools.ai/blog/deepseek-r2-never-shipped-what-deepseek-released-instead-2026
- https://openrouter.ai/moonshotai/kimi-k2
- https://www.marktechpost.com/2026/06/12/moonshot-ai-releases-kimi-k2-7-code-a-coding-model-reporting-21-8-on-kimi-code-bench-v2-over-k2-6/
- https://venturebeat.com/technology/kimi-k3s-full-weights-are-here-but-theyre-open-with-a-caveat-what-enterprises-should-know
- https://www.marktechpost.com/2026/08/26/z-ai-releases-glm-5-3-flash-a-320b-a18b-natively-multimodal-moe-with-a-1m-token-context/
- https://www.unite.ai/z-ai-details-glm-5-3-flash-inference-build-on-100-000-chinese-chips/
- https://cryptobriefing.com/zhipu-glm-53-flash-chinese-ai-chips/
- https://huggingface.co/MiniMaxAI/MiniMax-M3
- https://platform.minimax.io/docs/guides/pricing-paygo
- https://the-decoder.com/chinas-minimax-h3-is-the-first-open-model-to-top-an-ai-video-ranking/
- https://www.techtimes.com/articles/318622/20260618/minimax-m3-takes-open-weight-ai-lead-sparse-attention-architecture-now-verified.htm
- https://felloai.com/glm-5-2/
- https://felloai.com/minimax-m2-5-model/
- https://forum.devtalk.com/t/laya-the-os-version-of-jev-created-a-year-ago/249901
# STEP 1 — Open / Local AI Models — WESTERN TRACK (EN)

**Consolidated research fiche for RAG ingestion. Compiled 2026-09-22.**
*Coverage: Meta Llama · Meta Muse Spark · Poolside · Mistral AI · NVIDIA Nemotron · Microsoft Phi · Google Gemma · other notable Western open-weight releases through September 2026. Benchmark figures are vendor-reported unless marked independent.*

---

# PART 1 — META LLAMA FAMILY

## 1.1 Timeline overview (through September 2026)

| Date | Event |
|---|---|
| **2025-04-05** | Meta releases **Llama 4 Scout** and **Llama 4 Maverick** — first natively-multimodal, open-weight MoE models in the Llama family. Llama 4 **Behemoth** (288B-active teacher) announced as still training. |
| 2025-04 | **Benchmark controversy**: the LMArena entry `Llama-4-Maverick-03-26-Experimental` turns out to be an unreleased variant scoring far above the public weights; later publicly acknowledged. |
| 2025-06 | Meta forms **Meta Superintelligence Labs (MSL)** under Chief AI Officer **Alexandr Wang**, following the ~$14B Scale AI investment. |
| **2026-01** | Departing chief AI scientist Yann LeCun publicly acknowledges the Llama 4 benchmark manipulation. |
| **2026-04-08** | **Muse Spark** launches — first model from MSL, **closed-weight/proprietary** (Meta's first ever proprietary model), replacing the Llama brand. Internally codenamed **"Avocado"**. |
| 2026-07-09 | **Meta Model API** goes public and paid — Meta's first metered inference API. |
| 2026-07 | **Muse Spark 1.1** ships. |
| **2026-08-05** | **Muse Spark 1.2** + **Muse Code** (terminal coding agent, macOS/Linux beta). |
| **2026-08-10** | **Muse Glimmer 30B** — open weights under **Apache 2.0**; Meta announces **Muse Spark 1.2's weights will also be open-sourced**. Zuckerberg publishes an essay framing open weights as a safeguard against AI concentration. |
| 2026-09-02 | **Muse Spark 1.3** ships (xhigh public, max partner preview); paid API opened Sept 3. |
| 2026-09-08 | **"Muse" consumer app** launches — personal agent app on per-user Meta-hosted VMs. |
| Sept 2026 | Current frontier: **muse-spark-1.3**. **No Llama 5 exists** — the newest Meta model on OpenRouter is still `meta-llama/llama-4-maverick` (Apr 2025). Behemoth was never released. |

## 1.2 Llama 4 Scout

**Release date:** April 5, 2025

| Attribute | Value |
|---|---|
| Architecture | Mixture-of-Experts, natively multimodal (early fusion; text+image in) |
| Parameters | **109B total / 17B active per token, 16 experts** |
| Context window | **10M tokens** — largest open-weight context at release |
| Training data | ~40T tokens, text-focused corpus |
| Hardware | Fits on a single NVIDIA H100 (FP8) |
| License | **Llama 4 Community License** (gated; accept on Hugging Face) |
| HF weights | `meta-llama/Llama-4-Scout-17B-16E-Instruct` (+ `-Instruct-Original` native format, base `-17B-16E`); NVIDIA FP8/FP4 builds `nvidia/Llama-4-Scout-17B-16E-Instruct-FP8` |

**Benchmarks (Meta-reported, April 2025):** MMLU **79.6%**, MMLU Pro **74.3%**, GPQA Diamond **57.2%**, MATH **50.3%**, LiveCodeBench **32.8%**, MMMU **69.4%**, MGSM **90.6%**.
**Benchmarks (Artificial Analysis, independent):** Intelligence Index v4.3 **6.5** (earlier AA estimate **8**); GPQA Diamond **58.7%**, SciCode **21.3%**, Terminal-Bench 2.1 **3.7%**, AIME 2025 **14%**. Output speed **100.1 tok/s** median, TTFT **0.83s**.
**API pricing:** AA median across providers **$0.19 / $0.68** per 1M in/out tokens; OpenRouter-range ≈ **$0.15–$0.20 in / $0.50–$0.80 out**.
**What distinguishes it:** the 10M-token context window in an open-weight model (∼78× GPT-4o's 128K), single-H100 deployability, and best-in-class multimodal quality at its size tier (beat Gemma 3, Gemini 2.0 Flash-Lite, Mistral 3.1 per Meta).

## 1.3 Llama 4 Maverick

**Release date:** April 5, 2025

| Attribute | Value |
|---|---|
| Architecture | MoE, natively multimodal (early fusion); distilled partly from Behemoth |
| Parameters | **~400B total / 17B active per token, 128 experts** |
| Context window | **1M tokens** |
| Training data | ~22T tokens, multimodal incl. Meta content |
| License | **Llama 4 Community License** |
| HF weights | `meta-llama/Llama-4-Maverick-17B-128E-Instruct` (+ `-Original` native); Ollama `llama4:maverick` |

