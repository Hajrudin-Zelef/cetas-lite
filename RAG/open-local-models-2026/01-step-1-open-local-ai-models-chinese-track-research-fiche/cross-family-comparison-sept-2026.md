---
id: open-local-models-2026/01-step-1-open-local-ai-models-chinese-track-research-fiche/cross-family-comparison-sept-2026
title: "CROSS-FAMILY COMPARISON (Sept 2026)"
domain: step-1-open-local-ai-models-chinese-track-research-fiche
role: deep-dive
task: reference
actors: ["Alibaba", "China", "DeepSeek", "Huawei", "Meta", "MiniMax", "Moonshot", "OpenRouter", "Xiaomi", "Z.ai", "xAI"]
dates: ["2026-05"]
keywords: ["agent", "apache", "attention", "benchmarks", "datacenter", "deepseek", "fp4", "glm", "gpu", "grok", "grok 4", "inference"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [198, 279]
section: "STEP 1 — Open / Local AI Models (Chinese track) — Research Fiche"
sha256: 17c637cb89d4af888ecc2170b9d7c133119ca274a70137dbed9a298c7a758fbc
---

# CROSS-FAMILY COMPARISON (Sept 2026)

## CROSS-FAMILY COMPARISON (Sept 2026)

| Model | Release | Params (total/active) | Context | License | API $/M in/out | SWE-bench Verified | AA Index |
|---|---|---|---|---|---|---|---|
| Qwen3.5-397B-A17B | Feb 2026 | 397B / 17B | 262K→1M+ | Apache 2.0 | — | 76.2 | — |
| Qwen3.6-27B | Apr 2026 | 27B dense | 256K | Apache 2.0 | — | 77.2 | — |
| Qwen3.6-Max-Preview | Apr 2026 | 35B / 3B | 260K | Proprietary | — | — | 52 (v4.1.1) |
| Qwen3.7-Max | May 2026 | >1T | 1M | Proprietary | $2.50/$7.50 | — | 56.6 (v4.1.1) |
| Qwen3.8-Max | Aug 2026 | 2.4T / 95B | 1M | Custom (open) | $2 / $6 | — (Pro 67.7) | 45 (v4.3) |
| DeepSeek V4-Pro | Apr 2026 | 1.6T / 49B | 1M | MIT | $0.66/$1.98 | — | 36 |
| DeepSeek V4-Flash | Apr 2026 | 284B / 13B | 1M | MIT | $0.22/$0.66 | — | 35 |
| DeepSeek V4.1 Flash | Sep 2026 | 552B+196B / 8–16B | 1M | MIT | $0.15/$0.60 | — | ~39–40 |
| Kimi K2.6 | Apr 2026 | 1T / 32B | 256K | Modified MIT | $0.95/$4.00 | 80.2 | 54 |
| Kimi K3 | Jul 2026 | 2.8T / ~16×896e | 1M | Custom ($20M MaaS gate) | $3 / $15 | — | 44 (v4.3) |
| GLM-5.2 | Jun 2026 | 753B / 40B | 1M | MIT | $1.40/$4.40 | — (Pro 62.1) | 51 |
| GLM-5.3 | Aug 2026 | 753B / 40B | 1M | Custom ($10B review) | ~$2.15 blended | — | 45 (v4.3) |
| GLM-5.3-Flash | Aug 2026 | 320B / 18B | 1M | MIT | $0.15/$0.50 | — | 57 (v4.1.1) |
| MiMo-V2.5-Pro | Apr 2026 | 1.02T / 42B | 1M | MIT | $1.00/$3.00* | — (Pro 57.2) | — |
| MiMo-V2.6-Pro | Sep 2026 | 1.02T / 42B | 1M | MIT | $0.435/$0.87 | — (DeepSWE 71.9) | 46 (v4.3) |
| MiMo-V2.6-Flash | Sep 2026 | 309B / 15B | 1M | MIT | $0.14/$0.28 | — | — |
| MiniMax M2.5 | Feb 2026 | 229B / 10B | 196K | MINIMAX license | $0.30/$1.20 | 80.2 | — |
| MiniMax M2.7 | Mar 2026 | 230B / 10B | 200K | Non-commercial | $0.30/$1.20 | — (Pro 56.22) | 50 |
| MiniMax M3 | Jun 2026 | 428B / 23B | 1M | Community ($20M gate) | $0.30/$1.20 | 80.5 | — |
| Jev (TypeSafe) | Sep 2026 | n/a (decision model) | 64K | Proprietary | $0.042/in, out free | n/a | n/a |

\*V2.5 launch-era pricing; later revised downward.

## STRUCTURAL FINDINGS (for RAG metadata)
1. **License tightening is universal** across Chinese open-weight families: MIT/Apache → custom licenses with revenue-gated MaaS clauses ($20M: Kimi K3, MiniMax M3; $10B security review: GLM-5.3; territorial: MiniMax H3). Counter-example: GLM-5.3-Flash stayed plain MIT.
2. **Sovereign-hardware narrative**: GLM-5.3-Flash serves entirely on ~100K domestic Chinese accelerators (Cambricon/Huawei/Moore Threads) — built in ~13 days with Infra-Agent automation; MiMo-V2.6's RL ran partly on Chinese chips.
3. **Staged open-weight release**: GLM-5.3's 2-week safety hold is the first explicit pause→harden→ship precedent.
4. **1M context is table stakes** for flagships (K3, GLM-5.x, M3, MiMo-V2.x, Qwen Max, DeepSeek V4.x); Kimi K2.x (256K) and MiniMax M2.x (~200K) are the exceptions.
5. **"Open" now means datacenter-scale**: flagship open models are trillion-parameter MoEs; small dense models (Qwen3.6/3.8-27B) serve the single-GPU tier.
6. **Architectural trend**: off plain attention — Gated DeltaNet hybrids (Alibaba), compressed-sparse attention + FP4 KV (DeepSeek), KDA linear attention (Moonshot, Z.ai), MiniMax Sparse Attention, encoder-decoder (DeepSeek V4.1 Flash), Engram conditional memory (DeepSeek).
7. **Meta out**: Llama 5 delayed to 2027, Meta pivoted to closed Muse Spark — Chinese labs own the open frontier.

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
