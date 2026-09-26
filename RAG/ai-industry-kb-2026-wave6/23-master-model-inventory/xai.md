---
id: ai-industry-kb-2026-wave6/23-master-model-inventory/xai
title: "xAI"
domain: master-model-inventory
role: deep-dive
task: reference
actors: ["Alibaba", "DeepSeek", "Google", "Hugging Face", "MiniMax", "Moonshot", "Nvidia", "OpenRouter", "Unsloth", "Z.ai", "vLLM", "xAI"]
dates: ["2025-03-13", "2025-04-29", "2026-04-20", "2026-07-27", "2026-07-31", "2026-08-14", "2026-08-26", "2026-09-22", "2026-12-31"]
keywords: ["agents", "apache", "attribution", "benchmark", "benchmarks", "cost", "deepseek", "embedding", "gemini", "gemini 3.8", "gguf", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [11005, 11054]
section: "§23. Master Model Inventory"
delta_of: ai-industry-kb-2026
sha256: 516e821728d083eaef7ad3d69761d337a0d9775c10ae66cfc0551ac637ce0cc0
---

# xAI

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

### Qwen

- **Qwen3 dense family** (2025-04-29): 0.6B/1.7B/4B (32K), 8B/14B/32B (128K); Apache 2.0; 30B-A3B MoE (30.5B/3.3B active, 48 layers, 128 experts/8 active, 32K→131K YaRN) [SECONDARY]. Source: http://dev.to/best_codes/qwen-3-benchmarks-comparisons-model-specifications-and-more-4hoa
- **Qwen3-Coder-30B-A3B-Instruct** (2026): 1M-context GGUF quant exists via Unsloth; Apache 2.0 [SECONDARY]. Source: https://huggingface.co/unsloth/Qwen3-Coder-30B-A3B-Instruct-1M-GGUF
- **Qwen3.5-9B / Qwen3.6-27B**: Apache 2.0 confirmed [COMMUNITY]. Source: https://github.com/ufal/atrium-project/issues/9
- **Qwen3.7 Max Thinking** (2026-07): ELO 1475 (#9); BenchLM 77.4/100; MIT; $4/$16 per 1M, $1.20 cached, $0.50 vector [SECONDARY]. Sources: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf and https://github.com/kzinmr/ai-topics/blob/HEAD/wiki/comparisons/llm-api-pricing.md
- **Qwen3.7 Plus**: 72B/18B active, ~1420 ELO, ~88% GPQA, ~280 tok/s, $0.32–0.42/M (provider variance UNRESOLVED), Apache 2.0 [SECONDARY]. Sources: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf and https://github.com/vLLM-X/ktransformers/blob/HEAD/doc/en/api/server.rst
- **Qwen3.7 Turbo**: ~1380 ELO, ~450 tok/s, $0.12/M, Apache 2.0 [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- **Qwen3.8-Max**: custom text-only weights (not Apache) [SECONDARY]. Source: wave6/03-model-weights-wave3.md
- **Qwen3-Embedding** (0.6B/4B/8B): 8B 70.58 multilingual MMTEB / 73.84 STS; Apache 2.0; paper arXiv 2506.05176 [VENDOR]. Source: https://github.com/QwenLM/Qwen3-Embedding
- **Qwen3-Reranker** (0.6B/4B/8B): 4B BEIR 61.16/CoIR 73.91; 8B CMTEB-R 77.45; Apache 2.0 [VENDOR]. Source: https://github.com/QwenLM/Qwen3-Embedding
- **Qwen3-VL-Embedding**: Apache 2.0; multimodal [VENDOR]. Source: https://github.com/QwenLM/Qwen3-Embedding
- **Qwen3.5-Omni Light**: license UNVERIFIED; API-only reading exists [UNVERIFIED].

### Z.ai

- **GLM-5.3** (2026-08-14): 744B/743B/753B counting conventions (same base as 5.2); closed at launch, weights staged behind safety review; Z.ai list $1.40/$0.26/$4.40 [SECONDARY]; Terminal-Bench 2.1 (Ante): 74.6%±2.06 [SECONDARY]; DeepSWE 2026-09-22: 69%±3 at $3.99 (max) [SECONDARY]. Sources: https://emergent.sh/learn/what-is-glm-5-3, https://deepswe.datacurve.ai/, https://github.com/antigmalabs/ante/blob/HEAD/docs-site/docs/benchmarks/eval.mdx
- **GLM-5.3-Flash** (2026-08-26): MIT, zai-org/GLM-5.3-Flash; 200B-class, runs on non-NVIDIA/domestic chips; Z.ai list $0.15/$0.03/$0.50; most popular model on OpenRouter by token usage; DeepSWE: 63%±4 [SECONDARY]. Sources: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com and https://deepswe.datacurve.ai/
- **GLM-5.2**: MIT; OpenRouter (2026-08-26) $0.60/$1.70; Terminal-Bench 2.1 (Ante): 74.6%±2.06 [SECONDARY]. Sources: https://github.com/prism-shadow/agenthub/blob/HEAD/changelog/0.4.8/2026-08-26-glm-5.3-flash-vision.md and https://www.orcarouter.ai/blog/deepseek-v4-1-flash-vs-kimi-k3
- **GLM-4.7**: named in corpus; license UNVERIFIED [UNVERIFIED].

### Moonshot

- **Kimi K3** (2026-07): Modified MIT (100M MAU attribution; $20M/mo revenue prong UNRESOLVED); weights committed by 2026-07-27; $3/$0.30/$15 per 1M; DeepSWE 2026-09-22: 69%±5 at $4.65 (max) [SECONDARY]. Sources: https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/ and https://deepswe.datacurve.ai/
- **Kimi K2.7 Code**: named in corpus (Kimi K2.7 Code vendor boards); license UNVERIFIED [UNVERIFIED].
- **Kimi K2.6** (2026-04-20): license UNVERIFIED — do not assume K3's Modified MIT [UNVERIFIED].

### MiniMax

