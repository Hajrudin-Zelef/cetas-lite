---
id: ai-industry-kb-2026-wave6/20-small-models-edge-quantization-serving-embeddings-and-rerank/part-8
title: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers (part 8)"
domain: small-models-edge-quantization-serving-embeddings-and-rerank
role: deep-dive
task: quantization
actors: ["Alibaba", "China", "Google", "Hugging Face", "Microsoft", "Mistral", "OpenAI", "TensorRT-LLM", "Unsloth", "vLLM"]
dates: ["2026-04", "2026-07"]
keywords: ["embedding", "quantization", "reranker", "agent", "agentic", "apache", "attention", "awq", "benchmark", "claude", "compute", "copilot"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9795, 9828]
section: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers"
delta_of: ai-industry-kb-2026
sha256: a52d12d4283d11a118107e2834aa651bad13ac859d96f74815b6c8668a6d6145
---

# §20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers (part 8)

- Mistral's data-training posture (2026): Vibe input/output used for training by default unless opted out; Vibe Enterprise opted out by default; Studio Free may use data; pay-as-you-go can opt out; Zero Data Retention is separate and approval-based [SECONDARY]. Source: https://www.aipricing.guru/mistral-ai-pricing/
- Mistral's "latest" aliases (e.g. mistral-medium-latest) can silently point to newer implementations — pinning model IDs matters for reproducibility [SECONDARY]. Source: https://aiworldtoday.com/guides/mistral-ai-pricing


- Phi-4 at 14B outperformed GPT-4 on some STEM reasoning tasks while requiring a fraction of the infrastructure [SECONDARY]. Source: https://www.turing.com/blog/exploring-phi-4
- Phi-4 ships 128K-token support, function calling, and reasoning-optimized variants per a 2026 enterprise review [SECONDARY]. Source: https://www.turing.com/blog/exploring-phi-4
- Production Phi-4-reasoning note: whole L40S GPU, minReplicas 1, no scale-to-zero, 600s timeout, vLLM v0.20.2 [COMMUNITY]. Source: https://github.com/ualberta-rcg/aleph/blob/HEAD/models/phi-4-reasoning/CLAUDE.md
- The same deployment sets VLLM_ATTENTION_BACKEND=TRITON_ATTN_VLLM_V1 on L40S (SM89) — attention backend pinning is part of the 2026 small-model serving recipe [COMMUNITY]. Source: https://github.com/ualberta-rcg/aleph/blob/HEAD/models/phi-4-reasoning/CLAUDE.md
- Phi-4-mini-reasoning exists as a variant ("like Phi-4-mini + reasoning") alongside Phi-3-mini-4k/128k, Phi-3-medium-4k/128k, and Phi-3-small-8k in the community architecture table [COMMUNITY]. Source: https://github.com/labscommunity/cascadia/blob/HEAD/docs/architectures/phi.md
- Phi-4-mini config specifics: vocab_size 200064, sliding_window 262144, partial_rotary_factor 0.75 [COMMUNITY]. Source: https://github.com/labscommunity/cascadia/blob/HEAD/docs/architectures/phi.md
- Gemma 3 270M was trained on 6T tokens versus 2T for the 1B — the smallest Gemma 3 got 3x the tokens of its larger sibling (overtrained tiny) [SECONDARY]. Source: https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_3.md
- A 1M-context GGUF quant of Qwen3-Coder-30B-A3B-Instruct exists via Unsloth — the coder variant's long context survives quantization [SECONDARY]. Source: https://huggingface.co/unsloth/Qwen3-Coder-30B-A3B-Instruct-1M-GGUF
- SmolVLM (companion to SmolLM3): compact multimodal models for visual QA, image description, visual storytelling, multiple images per conversation, on-device [SECONDARY]. Source: https://github.com/huggingface/smollm/
- SmolLM3's full training configs are published at huggingface.co/datasets/HuggingFaceTB/smollm3-configs — the "engineering blueprint" release [SECONDARY]. Source: https://www.xugj520.cn/en/archives/smollm3-compact-multilingual-ai.html
- Community packaging: SmolLM3-3B runs behind Gradio/Pinokio installers (pierrunoyt) with extended-thinking toggle, temperature/top-p/max-tokens controls [COMMUNITY]. Source: https://github.com/pierrunoyt/smollm3-3b-pinokio-gradio
- Qwen3-Embedding paper citation: arXiv 2506.05176, "Qwen3 Embedding: Advancing Text Embedding and Reranking Through Foundation Models" [VENDOR]. Source: https://github.com/QwenLM/Qwen3-Embedding
- Qwen3-Reranker-8B: CMTEB-R 77.45 (best Chinese rerank in the table), MMTEB-R 72.94, MLDR 70.19, FollowIR 8.05; 4B takes FollowIR at 14.84 [VENDOR]. Source: https://github.com/QwenLM/Qwen3-Embedding
- Code-retrieval reranker gap: Qwen3-Reranker-4B MTEB-Code 81.20 versus BGE-reranker-v2-m3 41.38 — a 40-point split on code tasks [SECONDARY]. Source: https://github.com/QwenLM/Qwen3-Embedding
- MTEB-BR open-model table (22 Brazilian-Portuguese tasks) [COMMUNITY]: Linq-Embed-Mistral 7B 0.6473 | text-embedding-3-large 0.6449 | SFR-Embedding-2_R 7B 0.6397 | gte-Qwen2-7B-instruct 0.6392 | microsoft/harrier-oss-v1-27b 0.6390 | bge-m3 0.6157. Source: https://github.com/tardellirs/colibri-embed
- microsoft/harrier-oss-v1-27b: a 27B Microsoft embedding model appearing on the MTEB-BR board at 0.6390 — a new model name for the corpus [COMMUNITY]. Source: https://github.com/tardellirs/colibri-embed
- Qwen2.5-14B-Instruct-AWQ confirmed Apache 2.0 — quantized redistributions keep the upstream license and are described as safe for quantized production [COMMUNITY]. Source: https://github.com/ufal/atrium-project/issues/9
- Colibri serving details: fp16 branch plus fast ONNX path (~2x CPU throughput), Matryoshka 768/512/256/128 [COMMUNITY]. Source: https://github.com/tardellirs/colibri-embed
- EmbeddingGemma CoreML conversion toolchain handles palettization, blockwise quantization, and .mlpackage → .mlmodelc compilation [COMMUNITY]. Source: https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md
- Mistral model-role positioning (2026): Small = cost-sensitive general workloads; Large = open-weight flagship general; Medium = long-running agentic/coding; Codestral = frequent code completion [SECONDARY]. Source: https://aiworldtoday.com/guides/mistral-ai-pricing
- Leanstral, an open Lean 4 code agent from Mistral, is currently free [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026
- Mistral price-history tracking began April 2026; flat charts stay hidden until a change is detected [SECONDARY]. Source: https://www.aipricing.guru/mistral-ai-pricing/
- Qwen3.7 Turbo: ~1380 ELO, ~450 tok/s, $0.12/M input, Apache 2.0 [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- Qwen3.7 Plus: 72B total / 18B active, ~1420 ELO, ~88% GPQA, ~280 tok/s, $0.32–0.42/M, Apache 2.0 [SECONDARY]. Sources: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf and https://github.com/kzinmr/ai-topics/blob/HEAD/wiki/comparisons/llm-api-pricing.md
- Qwen3.6 Max Preview: 1446 ELO, $7.80/M input [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- Qwen3.7 Max Thinking: BenchLM 77.4/100 in the July 2026 ranking compilation [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- Copilot+ background AI: "tens of AI data models always running in the background collecting tons of information real-time right on the device" per Counterpoint's Neil Shah — the privacy pitch is local processing, not no processing [SECONDARY]. Source: https://www.computerworld.com/article/2116724/microsoft-launches-ai-powered-copilot-pcs.html?ref=aileap&utm_source=aileap&utm_medium=referral
- Microsoft's NPU economics claim: "What once required compute that cost thousands of dollars can now be done on a device with options that cost hundreds" [VENDOR]. Source: https://news.microsoft.com/source/features/ai/how-the-npu-is-paving-the-way-toward-a-more-intelligent-windows/
- Gaps not verified in this pass: Llama 3.2 1B/3B specs, Gemini Nano 2026 status, and TensorRT-LLM 2026 release notes were not pulled — do not treat their absence here as absence of releases [DIRECTIONAL].

