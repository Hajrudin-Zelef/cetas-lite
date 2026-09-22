---
id: ai-industry-kb-2026-wave6/23-master-model-inventory/nvidia-cohere-jina-others
title: "NVIDIA / Cohere / Jina / Others"
domain: master-model-inventory
role: deep-dive
task: reference
actors: ["Alibaba", "Apple", "China", "Cohere", "DeepSeek", "Falcon", "Microsoft", "Mistral", "Nvidia", "OpenAI", "Sakana", "Unsloth", "Z.ai"]
dates: ["2025-03-20", "2026-02-16", "2026-03-30", "2026-04-02", "2026-04-16", "2026-04-22", "2026-07-21", "2026-07-24", "2026-07-31", "2026-08-12", "2026-08-13", "2026-09-01", "2026-09-02", "2026-09-10", "2026-09-14"]
keywords: ["cohere", "nvidia", "apache", "astra", "awq", "benchmarks", "consumer", "decode", "deepseek", "embedding", "embeddings", "fugu"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [11096, 11167]
section: "§23. Master Model Inventory"
delta_of: ai-industry-kb-2026
sha256: 17b98f819a2b7bfb56c2f0ad2623814f6a182d19025acf62546f535a04b69603
---

# NVIDIA / Cohere / Jina / Others

### NVIDIA / Cohere / Jina / Others

- **NVIDIA Nemotron 3 / 3.5**: Ultra 550B/55B; NVIDIA Open Model License [SECONDARY]. Source: wave6/02-model-weights-delta.md
- **NVIDIA nemotron-speech-streaming-en-0.6b**: carries NOML NOTICE-file requirement [COMMUNITY]. Source: https://github.com/mikeschirtzinger/silent-notetaker/blob/HEAD/docs/research/model-licenses.md
- **Cohere Command A+** (218B/25B): Apache 2.0 [SECONDARY]. Source: wave6/02-model-weights-delta.md
- **Cohere Embed v4**: proprietary; retains input_type (query vs document) [SECONDARY]. Source: https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md
- **Jina v3** (CC BY-NC 4.0) / **v4** (Qwen Research, non-commercial) / **v5-omni-small** (CC BY-NC-4.0, 0.6B, 20K context): non-commercial tier [SECONDARY]. Sources: https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md and https://www.codesota.com/benchmarks/mteb
- **Jina reranker v3** (0.6B): BEIR 61.94/MIRACL 66.83/MKQA 67.92/CoIR 70.64 [VENDOR]. Source: https://huggingface.co/jinaai/jina-reranker-v3
- **KaLM-Embedding-Gemma3-12B-2511** (11.76B, 3840 dims): 72.32 MMTEB ceiling; custom license [SECONDARY]. Source: https://www.codesota.com/benchmarks/mteb
- **BGE-large-en-v1.5**: Apache 2.0; MTEB 64.23 (2024 figure — date it) [SECONDARY]. Source: https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
- **nomic-embed-text** (137M): Apache 2.0; 62.39 MTEB [SECONDARY]. Source: https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
- **Colibri-0.3B** (Portuguese): 72.71 MTEB-BR aggregate, beats its 300M base at half size; Matryoshka + ONNX [COMMUNITY]. Source: https://github.com/tardellirs/colibri-embed
- **Sakana Fugu**: named in corpus; specs UNVERIFIED [UNVERIFIED].
- **Tencent Hy3**: named in corpus; specs UNVERIFIED [UNVERIFIED].
- **Falcon-H1** (2026-01/2026-07 releases): releases verified, license/specs UNVERIFIED in this pass [UNVERIFIED]. Source: wave6/02-model-weights-delta.md
- **GPT-6 Astra**: also listed under OpenAI — no duplicate entry; cross-reference only.

### Additional entries (variants, embeddings, rerankers, retired)

- **Qwen3-0.6B-GGUF** (Unsloth): community quant of Qwen3-0.6B; Apache 2.0 inherited; purpose-built variants (e.g. Heimdall) exist [SECONDARY]. Sources: https://huggingface.co/unsloth/Qwen3-30B-A3B-GGUF and https://github.com/mikeschirtzinger/silent-notetaker/blob/HEAD/docs/research/model-licenses.md
- **Qwen3-30B-A3B MLX 4-bit**: Apple-silicon port; Apache 2.0 [SECONDARY]. Source: https://huggingface.co/Qwen/Qwen3-30B-A3B-MLX-4bit
- **Qwen3-235B-A22B-Instruct-2507**: Apache 2.0 confirmed [COMMUNITY]. Source: https://github.com/ufal/atrium-project/issues/9
- **onnx-community Qwen3-0.6B/1.7B-ONNX**: Apache 2.0 inherited via README YAML [COMMUNITY]. Source: https://github.com/mikeschirtzinger/silent-notetaker/blob/HEAD/docs/research/model-licenses.md
- **Qwen2.5-14B-Instruct-AWQ**: Apache 2.0; quantized redistribution [COMMUNITY]. Source: https://github.com/ufal/atrium-project/issues/9
- **DeepSeek V4 (alias)**: retired 2026-07-24; superseded by V4 Pro / V4 Flash / V4.1 Flash — do not cite "V4" without a date [SECONDARY]. Source: wave6/03-model-weights-wave3.md
- **GLM-4.5 / GLM-4.7-FlashX**: named in cache-pricing catalog; licenses UNVERIFIED [SECONDARY]. Source: https://github.com/2389-research/dippin-lang/commit/437a8a2889b92a4842582195f2a525836ad93645
- **Phi-3 family** (mini-4k/128k, medium-4k/128k, small-8k): prior generation; superseded by Phi-4 [COMMUNITY]. Source: https://github.com/labscommunity/cascadia/blob/HEAD/docs/architectures/phi.md
- **Phi-4-mini-reasoning**: mini + reasoning variant [COMMUNITY]. Source: https://github.com/labscommunity/cascadia/blob/HEAD/docs/architectures/phi.md
- **Llama 3.2 1B/3B**: specs not re-verified in this pass — entry reserved, fields UNVERIFIED [UNVERIFIED].
- **Mistral OCR 4** ($4/1K pages OCR; $5/1K Document AI) vs **Mistral OCR** ($2/1K, legacy $8): version-specific pricing [SECONDARY]. Sources: https://curlscape.com/blog/mistral-api-pricing-2026 and https://aiworldtoday.com/guides/mistral-ai-pricing
- **Le Chat** (Free $0 / Pro $14.99/mo / Team $24/user/mo): Mistral's consumer surface [SECONDARY]. Source: https://www.smashingapps.com/mistral-ai-review/
- **BGE-reranker-v2-m3**: MTEB-R 57.03 / CMTEB-R 72.16 (best Chinese) / MTEB-Code 41.38 [SECONDARY]. Source: https://github.com/QwenLM/Qwen3-Embedding
- **gte-multilingual-reranker-base** (0.3B): MTEB-R 59.51 / CMTEB-R 74.08 / MLDR 66.33 [SECONDARY]. Source: https://github.com/QwenLM/Qwen3-Embedding
- **Jina-multilingual-reranker-v2-base** (0.3B): MTEB-R 58.22 / CMTEB-R 63.37 / MMTEB-R 63.73 [SECONDARY]. Source: https://github.com/QwenLM/Qwen3-Embedding
- **ritrieve_zh_v1** (0.3B): 72.71 STS aggregate; strongest Chinese-subset scores in Qwen table [SECONDARY]. Source: https://github.com/QwenLM/Qwen3-Embedding
- **SFR-Embedding-2_R** (7B): 0.6397 MTEB-BR [COMMUNITY]. Source: https://github.com/tardellirs/colibri-embed
- **Linq-Embed-Mistral** (7B): 0.6473 MTEB-BR [COMMUNITY]. Source: https://github.com/tardellirs/colibri-embed
- **gte-Qwen2-7B-instruct**: 0.6392 MTEB-BR [COMMUNITY]. Source: https://github.com/tardellirs/colibri-embed
- **microsoft/harrier-oss-v1-27b** (27B): 0.6390 MTEB-BR — new Microsoft embedding model name [COMMUNITY]. Source: https://github.com/tardellirs/colibri-embed
- **OpenAI text-embedding-3-large** (0.6449 MTEB-BR) / **text-embedding-3-small** (62.3 MTEB): proprietary [SECONDARY]. Source: https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
- **voyage-3-large** (65.1 MTEB): proprietary [SECONDARY]. Source: https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
- **multilingual-e5-large-instruct** (58.08 MTEB): license UNVERIFIED [SECONDARY]. Source: https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
- **all-MiniLM-L6-v2** (~56 MTEB): the legacy baseline everything is measured against [SECONDARY]. Source: https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
- **FunAudioLLM/SenseVoiceSmall**: FunASR Model Open Source License v1.1 — commercial scope ambiguous, do not treat as MIT [COMMUNITY]. Source: https://github.com/mikeschirtzinger/silent-notetaker/blob/HEAD/docs/research/model-licenses.md

## Figures and metrics

### DeepSeek

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| DeepSeek V4 family (V4-Pro, V4-Flash, V4-Flash-Vision-Exp) | 2026 | 552B MoE (+196B Engram, ~748B total); 8B prefill / 16B decode [SECONDARY] | 1M (384K max output) | MIT · open [VENDOR] | Active; V4-Pro API traffic routes to V4.1 Flash rates (2026-09-14) [SECONDARY] |
| DeepSeek V4-Pro-0813 | 2026-08-13 | — | — | MIT [VENDOR] | Retired into V4.1 routing [SECONDARY] |
| DeepSeek V4-Flash-0731 | 2026-07-31 | — | — | MIT [VENDOR] | Active; 4.7M HF downloads (2026-09-01) [COMMUNITY] |
| DeepSeek V4.1-Flash | 2026-09-10 (GA) | — | — | MIT · open [VENDOR] | Active; peak/off-peak pricing [VENDOR] |
| DeepSeek R1-1776 | 2025-03-20 | — | — | MIT · open | Fine-tune (censorship-removed R1) [SECONDARY] |

### Qwen / Alibaba

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Qwen3.5-397B-A17B | 2026-02-16 | 397B/17B MoE [SECONDARY] | — | Apache 2.0 [VENDOR] | Active |
| Qwen3.6-35B-A3B | 2026-04-16 | 35B/3B [SECONDARY] | — | Apache 2.0 [VENDOR] | Active |
| Qwen3.6-27B | 2026-04-22 | 27B dense | 1M | Apache 2.0 [VENDOR] | Active |
| Qwen3-Next-80B-A3B | ~2026-04-02 | 80B/3B MoE [SECONDARY] | 1M | Apache 2.0 [VENDOR] | Preview ("Next" arch) |
| Qwen3.5-Omni (Light open variant) | 2026-03-30 | — | 256K | Light license [UNVERIFIED] | Active; API-only reading disputed [SECONDARY] |
| Qwen3.8-27B | 2026 | 27.8B dense [SECONDARY] | — | Apache 2.0 [VENDOR] | Active |
| Qwen3.8-Max | 2026 (weights 2026-08-12) | 2.4T/95B [SECONDARY] | 1M | Custom license (text-only weights) [VENDOR] | Active |
| Qwen3.8-Max-0902 | 2026-09-02 | 2.4T (same base) [SECONDARY] | 1M | Custom license | Refresh, not a new base [SECONDARY] |
| Qwen Image 3.0 / 3.0 Pro | 2026-07-21 | — | 4,500-token prompt input [VENDOR] | Closed/proprietary | API-only; no weights [SECONDARY] |
| qwen4_exp | 2026-08 | — | — | Internal | Experimental checkpoint; not Qwen 4 [SECONDARY] |

