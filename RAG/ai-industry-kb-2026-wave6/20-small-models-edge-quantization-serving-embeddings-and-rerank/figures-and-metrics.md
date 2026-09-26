---
id: ai-industry-kb-2026-wave6/20-small-models-edge-quantization-serving-embeddings-and-rerank/figures-and-metrics
title: "Figures and metrics"
domain: small-models-edge-quantization-serving-embeddings-and-rerank
role: deep-dive
task: quantization
actors: ["Alibaba", "Apple", "Cohere", "Microsoft", "Mistral", "Nvidia", "SGLang", "TensorRT-LLM", "Unsloth", "vLLM"]
dates: ["2026-01-15", "2026-02-18", "2026-02-23", "2026-03-02", "2026-03-31", "2026-04-03", "2026-04-07", "2026-04-16", "2026-04-29", "2026-05", "2026-05-07", "2026-05-12", "2026-05-16", "2026-05-28", "2026-06-08", "2026-06-27", "2026-07-11", "2026-07-25", "2026-08-13", "2026-08-26"]
keywords: ["apache", "benchmarks", "cohere", "decode", "embedding", "embeddings", "fp8", "gemini", "gguf", "gpu", "int4", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9829, 9907]
section: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers"
delta_of: ai-industry-kb-2026
sha256: 5cdd64d2fb99e542d81053d34c7a5a2b24cf3e4614956963c6d2aa40350fc7a6
---

# Figures and metrics

## Figures and metrics
| Small model | Release | Params | License / terms |
|---|---|---|---|
| Qwen3.5 (4 sizes) | 2026-03-02 | 9B/4B/2B/0.8B | Apache 2.0 [VENDOR] |
| LFM2.5-8B-A1B | 2026-05-28 | 8.3B/1.5B | LFM Open License ($10M cap) [SECONDARY] |
| LFM2.5-230M | 2026-06-27 | 230M | LFM Open License ($10M cap) [SECONDARY] |
| LFM2.5-2.6B | 2026 (date unpinned) | 2.6B | LFM Open License ($10M cap) [UNVERIFIED] |
| SmolLM3-3B | date not pinned | 3B | fully open recipe [SECONDARY] |
| Apple AFM 3 Core | 2026-06-08 | ~3B dense | proprietary, SDK-only [SECONDARY] |
| Apple AFM 3 Core Advanced | 2026-06-08 | 20B sparse | proprietary, SDK-only [SECONDARY] |
| IBM Granite 4.1 | 2026-04-29 | 3B/8B/30B | Apache 2.0 [SECONDARY] |
| "Ministral 3" | 2026 | — | — [UNVERIFIED] |

| Embedding model | Event | License |
|---|---|---|
| Voyage-4 | 2026-01-15 | API; nano Apache 2.0 [SECONDARY] |
| voyage-code-4 | 2026-08-13 | API [COMMUNITY] |
| jina-embeddings-v5 | 2026-02-23 | CC BY-NC 4.0 [SECONDARY] |
| jina-embeddings-v5-omni | 2026-05-07 | CC BY-NC 4.0 [SECONDARY] |
| Qwen3-Embedding-8B | 2025-06 (2026: MTEB lead 70.58) | Apache 2.0 [SECONDARY] |
| llama-nemotron-embed-1b-v2 | 2026-02 | NVIDIA Open Model License [SECONDARY] |

- jina-embeddings-v5 size correction: small 677M / nano 239M (press copy inverted) [SECONDARY].
- Ollama 0.19.0 gains are Qwen3.5-35B-A3B-specific: +57% prefill, +93% decode on M5 Max [SECONDARY].
- vLLM v0.25.0: 558 commits / 232 contributors in the release window [SECONDARY].
- Unsloth Qwen3.6-27B-NVFP4: 2.5× throughput claim, 24GB VRAM [COMMUNITY].


### New verified metrics — expansion

- Qwen3 dense spec table (secondary compilation) [SECONDARY]: 0.6B/1.7B/4B → 32K context; 8B/14B/32B → 128K; 30B-A3B → 30.5B total/3.3B active, 48 layers, 128 experts/8 active, 32,768 native → 131,072 YaRN. Source: http://dev.to/best_codes/qwen-3-benchmarks-comparisons-model-specifications-and-more-4hoa
- Gemma 3 spec table [SECONDARY]: 270M 6T tokens/32K | 1B 2T/32K | 4B 4T/128K | 12B 12T/128K | 27B 14T/128K; cutoff Aug 2024. Source: https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_3.md
- SmolLM3 mode-split benchmarks [SECONDARY]: AIME 2025 36.7/9.3 | LiveCodeBench 30.0/15.2 | GPQA Diamond 41.7/35.7 (think/no-think). Source: https://www.xugj520.cn/en/archives/smollm3-compact-multilingual-ai.html
- 2026 small-model comparison table (2025-era figures — date carefully) [SECONDARY]: Phi-4 14B — MMLU 84.8, HumanEval 82.6, MATH 80.4, 16K, MIT | Gemma 3 12B — 83.2/78.4/75.1, 128K | Mistral Small 3.1 24B — 82.7/79.1/71.3, 128K, Apache 2.0 | Llama 3.2 11B — 80.1/72.3/68.9, 128K, Llama 3 | Qwen 2.5 14B — 82.9/80.7/79.8, 128K, Apache 2.0 | Gemma 3 27B — 87.3/84.2/83.6, 128K. Source: https://ayinedjimi-consultants.fr/static/pdf/small-language-models-2026.pdf
- MTEB 2026 ceiling snapshot [SECONDARY]: MMTEB ceiling 72.32 (KaLM-Embedding-Gemma3-12B-2511) | Qwen3-Embedding-8B 70.58 multilingual (#1 Jun 2025) | EmbeddingGemma-300M 69.67 (MTEB-eng-v2, <500M class top) | gemini-embedding-001 68.3 multilingual | jina-embeddings-v3 65.5 | voyage-3-large 65.1 | text-embedding-3-large 64.6 | nomic-embed-text 62.39 | text-embedding-3-small 62.3 | all-MiniLM-L6-v2 ~56. Sources: https://www.codesota.com/benchmarks/mteb, https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md, https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md
- Qwen3-Embedding STS aggregates (official) [VENDOR]: 8B 73.84 | 4B 72.27 | 0.6B 66.33. Source: https://github.com/QwenLM/Qwen3-Embedding
- Reranker comparison [VENDOR]: Jina reranker v3 (0.6B) BEIR 61.94 / MIRACL 66.83 / MKQA 67.92 / CoIR 70.64 | Qwen3-Reranker-4B BEIR 61.16 / CoIR 73.91 | Qwen3-Reranker-0.6B BEIR 56.28 | BGE reranker v2 m3 BEIR 56.51 / MIRACL 69.32. Sources: https://huggingface.co/jinaai/jina-reranker-v3 and https://github.com/QwenLM/Qwen3-Embedding
- Qwen3-Reranker retrieval-subset table (official) [VENDOR]: 0.6B — MTEB-R 65.80, CMTEB-R 71.31, MMTEB-R 66.36, MLDR 50.26 | 4B — 69.76, 75.94, 72.74, 69.97 | 8B — 69.02, 77.45, 72.94, 70.19. Source: https://github.com/QwenLM/Qwen3-Embedding
- Quantization bits/weight ladder [COMMUNITY]: Q4_K_M ~4.5 | Q5_K_M ~5.5 | Q6_K ~6.5–6.6 | Q8_0 ~8–8.5 | FP8 8 | NVFP4 ~4 (+microscale, ~0.56 B/param) | MXFP4 ~4–4.5 (~0.53 B/param). Sources: https://github.com/kekzl/imp/blob/HEAD/docs/quantization.md, https://github.com/gguf-org/ggk/blob/HEAD/docs/editor/quantizer.md, https://github.com/kubesimplify/website/blob/HEAD/content/blog/day-4-quantization-demystified-bf16-fp8-nvfp4-mxfp4-int4-gguf-and-why-it-all-matters.md

## Main actors
- **Alibaba/Qwen** — Qwen3.5 small line (Apache 2.0); Qwen3-Embedding-8B MTEB Multilingual lead [VENDOR].
- **Liquid AI** — LFM2.5 family; $10M-revenue-capped open license [SECONDARY].
- **Apple** — Foundation Models 3 (WWDC26); proprietary on-device + PrivateCloudCompute [SECONDARY].
- **IBM** — Granite 4.1/4.2 (Apache 2.0) [SECONDARY].
- **Microsoft** — Phi-4 family current; no Phi-5 [SECONDARY].
- **vLLM / SGLang / Ollama / llama.cpp** — the serving-and-quantization stack: V1 lockdown, Model Runner V2, MLX backend, NVFP4/MXFP4/Q1_0/Q2_0 [SECONDARY].
- **Voyage AI** — Voyage-4 MoE embeddings; open-weight nano [SECONDARY].
- **Jina AI** — v5 text + v5-omni embeddings (CC BY-NC 4.0) [SECONDARY].
- **NVIDIA** — llama-nemotron-embed-1b-v2; TensorRT-LLM v1.3.0rc21 [SECONDARY].
- **Cohere** — Embed 4 / Rerank 4 (2025); FedRAMP High May 2026 [SECONDARY].

## Timeline and context
- **2026-01-15** — Voyage-4 launched [SECONDARY].
- **2026-02-18** — jina-embeddings-v5-text repositories dated (HF) [SECONDARY].
- **2026-02-23** — jina-embeddings-v5 announced [SECONDARY].
- **2026-03-02** — Qwen3.5 small line; vLLM v0.16.0 (V1 lockdown) [VENDOR].
- **2026-03-31** — Ollama 0.19.0 (MLX backend) [SECONDARY].
- **2026-04-03** — GPTQModel 6.0.3 [SECONDARY].
- **2026-04-07** — SGLang v0.5.10; SmolLM3-3B availability evidenced [SECONDARY].
- **2026-04-16** — GPTQModel 6.1.0 [SECONDARY].
- **2026-04-29** — IBM Granite 4.1 [SECONDARY].
- **2026-05-07** — jina-embeddings-v5-omni [SECONDARY].
- **2026-05-12** — Cohere FedRAMP High [SECONDARY].
- **2026-05-16** — SGLang v0.5.12 [SECONDARY].
- **2026-05-28** — Liquid LFM2.5-8B-A1B [SECONDARY].
- **2026-06** — GPU NVFP4 in llama.cpp (PRs #20644/#22196) [SECONDARY].
- **2026-06-08** — Apple Foundation Models 3 (WWDC26) [SECONDARY].
- **2026-06-27** — Liquid LFM2.5-230M [SECONDARY].
- **2026-07** — Unsloth Qwen3.6-27B-NVFP4; TensorRT-LLM v1.3.0rc21 [SECONDARY].
- **2026-07-11** — vLLM v0.25.0 (Model Runner V2 default) [SECONDARY].
- **2026-07-25** — SGLang v0.5.16 [COMMUNITY].
- **2026-08-13** — voyage-code-4 [COMMUNITY].
- **2026-08-26** — SGLang day-zero Qwen3.8 hybrid support [COMMUNITY].
- **~2026-08** — IBM Granite 4.2 [SECONDARY].


### New verified timeline entries — expansion

