---
id: ai-industry-kb-2026-wave6/20-small-models-edge-quantization-serving-embeddings-and-rerank/part-5
title: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers (part 5)"
domain: small-models-edge-quantization-serving-embeddings-and-rerank
role: deep-dive
task: quantization
actors: ["Alibaba", "Cohere", "Google"]
dates: ["2025-06", "2026-08-09"]
keywords: ["embedding", "embeddings", "quantization", "reranker", "apache", "awq", "benchmark", "benchmarks", "cohere", "cost", "distillation", "fp8"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9727, 9756]
section: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers"
delta_of: ai-industry-kb-2026
sha256: bbefd53746387279a8606ca52935b3079a68e2d91b2a3bfbd99c2031cc53dfa1
---

# §20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers (part 5)

- MTEB leaderboard state 2026 (community tracker): 56+ datasets, 8 task categories, MMTEB ceiling score 72.32, 15 models tracked [SECONDARY]. Source: https://www.codesota.com/benchmarks/mteb
- KaLM-Embedding-Gemma3-12B-2511 is the credible MMTEB ceiling model: 11.76B parameters, 3840 dimensions, custom community license, higher memory/indexing cost — treat as the upper bound, not the default production pick [SECONDARY]. Source: https://www.codesota.com/benchmarks/mteb
- Qwen3-Embedding official STS table (Qwen repo): 8B scores 73.84 overall (75.00/76.97/80.08/84.23/66.99/78.21/63.53 across sub-benchmarks); 4B 72.27; 0.6B 66.33; context 32,768; Apache 2.0 [VENDOR]. Source: https://github.com/QwenLM/Qwen3-Embedding
- Qwen3-Embedding-8B held #1 on MTEB multilingual at 70.58 as of June 2025 per a secondary compilation — the 70.58 (multilingual aggregate) and 73.84 (STS table) figures are different cuts, not a contradiction [SECONDARY]. Source: https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
- gte-Qwen2-7B-instruct: 7.6B parameters, 71.62 STS aggregate (72.19/75.77/66.06/81.16/69.24/75.70/65.20); gte-Qwen2-1.5B-instruct: 67.12 [SECONDARY]. Source: https://github.com/QwenLM/Qwen3-Embedding
- bge-multilingual-gemma2: 9B parameters, 67.64 STS aggregate — below gte-Qwen2-7B-instruct despite larger size [SECONDARY]. Source: https://github.com/QwenLM/Qwen3-Embedding
- BAAI/bge-m3: 568M parameters, MTEB-BR 0.6157 — the production-baseline workhorse, often winning after quantization and reranking despite losing benchmark points [SECONDARY]. Sources: https://github.com/tardellirs/colibri-embed and https://www.codesota.com/benchmarks/mteb
- multilingual-e5-large-instruct: 0.6B, 58.08 STS aggregate [SECONDARY]. Source: https://github.com/QwenLM/Qwen3-Embedding
- EmbeddingGemma-300M: 308M parameters, 768→512/256/128 via Matryoshka, MTEB English v2 69.67, 100+ languages, quantization-aware training holding RAM under 200MB — tops the under-500M class [SECONDARY]. Source: https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md
- Colibri (~157M, Brazilian-Portuguese): MTEB-BR 0.6501, beating its own embeddinggemma-300m base (0.6490) at half the size and matching models up to ~10x larger — vocab trim → multi-teacher distillation (Qwen3-4B+8B) → model soup [COMMUNITY]. Source: https://github.com/tardellirs/colibri-embed
- Cohere Embed v4: 256/512/1024/1536 Matryoshka dimensions, 128K context, proprietary license, multimodal — unified text/image/interleaved text+image vectors, supersedes v3 [SECONDARY]. Source: https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
- gemini-embedding-001: 3072 dimensions truncatable to 1536/768/256, 68.3 MTEB multilingual, Google's unified successor to text-embedding-004 [SECONDARY]. Source: https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
- voyage-3-large: Matryoshka 256/512/1024/2048, 32K context, proprietary, native output_dtype for int8/uint8/binary at embed time, MTEB 65.1 [SECONDARY]. Source: https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
- jina-embeddings-v3: 1024 down to 32 via Matryoshka, 8K context, 89 languages, task-specific LoRA adapters, CC BY-NC 4.0 (commercial license available), MTEB 65.5 [SECONDARY]. Source: https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
- jina-embeddings-v4: 3.8B, 2048 MRL, ViDoRe 84.11 dense / 90.17 late, ViDoRe-v3 57.52 — but Qwen Research non-commercial license disqualifies it for commercial pipelines [SECONDARY]. Source: https://github.com/t0msilver/vidtheque/blob/HEAD/research/multimodal-embedding-2026-08-09.md
- jina-embeddings-v5-omni-small: 1.74B, 1024 MRL (32–1024), CC BY-NC-4.0, ViDoRe-in-MIEB 79.08 — closest v4 rival, same license disqualification [SECONDARY]. Source: https://github.com/t0msilver/vidtheque/blob/HEAD/research/multimodal-embedding-2026-08-09.md
- Qwen3-VL-Embedding-2B: Apache 2.0, 2048 dims (64–2048 MRL), MMEB-V2 VisDoc 79.2, ViDoRe-v1 84.4 / v2 65.3, ~4.4GB bf16 — the recommended text→frame pick in one 2026 evaluation [SECONDARY]. Source: https://github.com/t0msilver/vidtheque/blob/HEAD/research/multimodal-embedding-2026-08-09.md
- Qwen3-VL-Embedding-8B: Apache 2.0, 4096 dims, VisDoc 82.4, ViDoRe-v1 87.2 / v2 69.9 — VRAM-disqualified (~16GB) in the same evaluation [SECONDARY]. Source: https://github.com/t0msilver/vidtheque/blob/HEAD/research/multimodal-embedding-2026-08-09.md
- text-embedding-3-large: 3072 dims, 8K context, proprietary, MTEB 64.6; text-embedding-3-small: 1536 (Matryoshka to 512), MTEB 62.3 [SECONDARY]. Source: https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
- nomic-embed-text: 137M, 768 MRL, MTEB English v2 62.39, 274MB — Apache 2.0 open alternative [SECONDARY]. Source: https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md
- all-MiniLM-L6-v2: 23M, 384 dims, ~56 older-MTEB, 46MB — still the floor baseline [SECONDARY]. Source: https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md
- Jina reranker v3: 0.6B on Qwen3-0.6B backbone, 28 layers, "last but not late interaction" listwise architecture, up to 64 documents in 131K context [VENDOR]. Source: https://huggingface.co/jinaai/jina-reranker-v3
- Jina reranker v3 scores: BEIR 61.94, MIRACL 66.83, MKQA 67.92, CoIR 70.64 — versus BGE reranker v2 m3 (BEIR 56.51, MIRACL 69.32), Qwen3-Reranker-0.6B (BEIR 56.28), Qwen3-Reranker-4B (BEIR 61.16, CoIR 73.91) [VENDOR]. Source: https://huggingface.co/jinaai/jina-reranker-v3
- Qwen3-Reranker official table: 0.6B MTEB-R 65.80 / CMTEB-R 71.31 / MMTEB-R 66.36; 4B 69.76 / 75.94 / 72.74; 8B 69.02 / 77.45 / 72.94 — evaluated on top-100 candidates from Qwen3-Embedding-0.6B [VENDOR]. Source: https://github.com/QwenLM/Qwen3-Embedding
- Reranker serving interfaces: /v1/rerank returns sorted Jina/Cohere-shaped output; /v1/score preserves input order [COMMUNITY]. Source: https://github.com/agentculture/lobes-cli/blob/HEAD/docs/qwen3-reranker-0.6b.md
- One local M5 test reported BGE v2 at 84ms/100 candidates versus Qwen3 4B at 312ms and hosted Cohere at 198ms — a single independent test, do not generalize [COMMUNITY]. Source: http://contracollective.com/blog/bge-reranker-v2-vs-cohere-rerank-3-vs-qwen3-reranker-m5-max-mlx-2026
- MLX-side reranker support recognizes three kinds: sequence-classifier, generative Qwen3, and Qwen3-VL multimodal rerankers [COMMUNITY]. Source: https://github.com/lablup/mlxcel/blob/HEAD/TECHNICAL_REPORTS/1417-v1-rerank-three-reranker-kinds-20260826.en.md
- GGUF quantization ladder (community reference): Q4_K_M ≈ 4.5 bits/weight (llama.cpp sweet spot), Q5_K_M ~5.5, Q6_K ~6.5–6.6, Q8_0 ~8–8.5 [COMMUNITY]. Sources: https://github.com/kekzl/imp/blob/HEAD/docs/quantization.md and https://github.com/gguf-org/ggk/blob/HEAD/docs/editor/quantizer.md
- FP8 is 8 bits/weight; NVFP4 ≈ 4 bits plus per-16 FP8 microscale/tensor scale, practical ~0.56 byte/parameter; MXFP4 ≈ 4–4.5 bits with 32-value blocks, practical ~0.53 byte/parameter [COMMUNITY]. Source: https://github.com/kubesimplify/website/blob/HEAD/content/blog/day-4-quantization-demystified-bf16-fp8-nvfp4-mxfp4-int4-gguf-and-why-it-all-matters.md
- AWQ is W4A16 (4-bit weights, FP16 activations, usually group size 128); GPTQ supports 2/3/4/8-bit weight-only paths [COMMUNITY]. Source: https://github.com/kekzl/imp/blob/HEAD/docs/quantization.md
