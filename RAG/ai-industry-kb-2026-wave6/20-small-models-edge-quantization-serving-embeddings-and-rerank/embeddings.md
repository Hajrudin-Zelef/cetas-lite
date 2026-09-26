---
id: ai-industry-kb-2026-wave6/20-small-models-edge-quantization-serving-embeddings-and-rerank/embeddings
title: "Embeddings"
domain: small-models-edge-quantization-serving-embeddings-and-rerank
role: deep-dive
task: quantization
actors: ["Alibaba", "Cohere", "Google", "Nvidia", "OpenAI", "vLLM"]
dates: ["2025-04", "2025-06", "2026-01-15", "2026-02", "2026-02-18", "2026-02-23", "2026-03", "2026-05-07", "2026-05-12", "2026-08-13"]
keywords: ["embedding", "embeddings", "agentic", "apache", "benchmark", "cohere", "gemini", "inference", "license", "llama", "llama.cpp", "moe"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9680, 9691]
section: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers"
delta_of: ai-industry-kb-2026
sha256: ebbcb954b79d8b6174e380e50b86d764056de4485388c7ddf8cbe5296cd07e67
---

# Embeddings

### Embeddings
- **2026-01-15** — Voyage-4 launched: voyage-4 / -large / -lite / -nano plus voyage-multimodal-3.5; first production MoE embedding model; shared embedding space; Matryoshka dims 256/512/1024/2048; MongoDB Atlas GA the same day [SECONDARY]; voyage-4-nano open-weight Apache 2.0 (~600M params, Qwen-based, 32K context) [SECONDARY]; **2026-08-13** — voyage-code-4: trained on issue-fixing PRs; +27.5% on an agentic code-retrieval benchmark, +14.0% across 28 datasets vs voyage-code-3, third of the price [COMMUNITY single-source].
- **2026-02-23** — jina-embeddings-v5 announced (HF org listing 2026-02-18): **small = 677M, nano = 239M — the press copy inverted these** [SECONDARY]; best-in-class among comparable sizes on MMTEB; open weights on HF; self-host via vLLM/llama.cpp/MLX; Elastic Inference Service; license **CC BY-NC 4.0** [SECONDARY]; **2026-05-07** — v5-omni (v5-omni-small/nano, multimodal text+image) [SECONDARY].
- Qwen3-Embedding-8B (June 2025 model): the 2026 delta is **continued MTEB Multilingual lead at 70.58**; 4096-dim MRL; 32K context; Apache 2.0; 100+ languages; 0.6B/4B siblings via Ollama; MTEB-Code 80.68 [SECONDARY].
- MTEB 2026 state: Gemini Embedding 001 English 68.32; gemini-embedding-2 preview March 2026 (multimodal); Voyage family high 60s; EmbeddingGemma-300M on-device; OpenAI text-embedding-3-large 64.6% English [SECONDARY].
- NVIDIA llama-nemotron-embed-1b-v2: February 2026, NVIDIA Open Model License [SECONDARY]; NV-Embed-v2 72.31 (March-2026 MTEB snapshot) [SECONDARY]; Llama-Embed-Nemotron-8B multilingual #1 provisional [SECONDARY].
- Cohere Embed 4 (April 2025) / Rerank 4 (late 2025, 32K context, Fast/Pro tiers): 2025 models; the 2026 event is FedRAMP High on **2026-05-12** [SECONDARY]; Rerank 4 context 32K [SECONDARY] vs 4K [COMMUNITY] — logged, unresolved [DIRECTIONAL].
- No verified 2026 ColBERT successor (jina-colbert-v2, 2024, latest found); no BGE-M3 successor (the GTE line continues via Qwen3-Embedding) [DIRECTIONAL].


### New verified facts — expansion

