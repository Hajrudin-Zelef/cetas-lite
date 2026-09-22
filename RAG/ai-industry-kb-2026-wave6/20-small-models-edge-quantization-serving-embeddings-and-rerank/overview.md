---
id: ai-industry-kb-2026-wave6/20-small-models-edge-quantization-serving-embeddings-and-rerank/overview
title: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers"
domain: small-models-edge-quantization-serving-embeddings-and-rerank
role: deep-dive
task: quantization
actors: ["Alibaba", "Apple", "Cohere", "DeepSeek", "Google", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM", "xAI"]
dates: ["2024-10", "2025-04", "2025-06", "2025-08", "2026-01-15", "2026-02", "2026-02-18", "2026-02-23", "2026-03", "2026-03-02", "2026-03-31", "2026-04", "2026-04-03", "2026-04-07", "2026-04-16", "2026-04-29", "2026-05-07", "2026-05-12", "2026-05-16", "2026-05-28", "2026-06-08", "2026-06-27", "2026-07-11", "2026-07-25", "2026-08", "2026-08-13", "2026-08-26", "2026-09"]
keywords: ["embedding", "embeddings", "quantization", "agentic", "apache", "attention", "benchmark", "blackwell", "cohere", "decode", "deepseek", "fp4"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9639, 9689]
section: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers"
sha256: 04f3027dadbcaffab68c4f40879db42edfe9b13704e89a2b30c4605186394263
---

# §20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers

Keywords: small language models, edge AI, quantization, vLLM, SGLang, Ollama, llama.cpp, embeddings, rerankers, Qwen3.5, Liquid LFM2.5, SmolLM3, Apple Foundation Models, Granite, NVFP4, MXFP4, GGUF, Voyage-4, jina-embeddings-v5, Matryoshka

## Summary
- **Qwen3.5 small line** (2026-03-02): 9B/4B/2B/0.8B, Apache 2.0, 262,144 native context extensible to ~1,010,000, hybrid Gated DeltaNet/attention [VENDOR].
- **Liquid AI LFM2.5**: 8B-A1B (2026-05-28), 2.6B (date not pinned), 230M (2026-06-27); LFM Open License v1.0 limits commercial use to organizations under $10M annual revenue [SECONDARY].
- **SmolLM3-3B availability is established** (Azure guide April 2026), but the release date is not pinned — do not cite it as a February–September 2026 launch [DIRECTIONAL].
- **Apple Foundation Models 3** announced at WWDC26 (2026-06-08): Core ~3B dense on-device + Core Advanced 20B sparse; proprietary, SDK-only [SECONDARY].
- **No Phi-5, no 2026 Command R7B successor, no xAI small/open model** in the corpus [DIRECTIONAL]; IBM Granite 4.1 (2026-04-29) and 4.2 (~August 2026), Apache 2.0 [SECONDARY]; Ministral 3 (2026) stays [UNVERIFIED] single-source.
- Serving: vLLM V1 lockdown (v0.16.0, 2026-03-02) then Model Runner V2 default (v0.25.0, 2026-07-11); SGLang day-zero Qwen3.8 hybrid support; Ollama 0.19.0 MLX backend (2026-03-31) with gains measured specifically on Qwen3.5-35B-A3B [SECONDARY].
- Embeddings: Voyage-4 (2026-01-15, first production MoE embedding model, nano open-weight Apache 2.0); jina-embeddings-v5 (small 677M / nano 239M — the press copy inverted them, CC BY-NC 4.0); Qwen3-Embedding-8B (June 2025 model; the 2026 delta is continued MTEB Multilingual lead at 70.58) [SECONDARY].
- No verified 2026 ColBERT or BGE-M3 successor; Cohere Embed 4/Rerank 4 are 2025 models whose 2026 event is FedRAMP High (2026-05-12) [SECONDARY].

## Key dated facts
### Small language models
- **2026-03-02** — Qwen3.5 line released: 9B/4B/2B/0.8B; 262,144 native context, extensible to roughly 1,010,000; hybrid Gated DeltaNet + attention; Apache 2.0; natively multimodal (text/image/video in) [VENDOR]; the 0.8B: 24 layers, 6× hybrid blocks, MTP-trained [VENDOR]; HF Base checkpoint: Qwen/Qwen3.5-0.8B-Base [COMMUNITY]; the 3.6-series QwenLM/Qwen3.6 repo is the adjacent lineage [SECONDARY].
- **2026-05-28** — Liquid AI LFM2.5-8B-A1B: 8.3B total / 1.5B active, 128K context, 9 languages; day-one llama.cpp/MLX/vLLM/SGLang/ONNX; 253 tok/s on M5 Max; under 6GB [SECONDARY].
- 2026 — Liquid LFM2.5-2.6B: Raspberry Pi-grade device pitch; release date **not pinned** in the corpus [UNVERIFIED].
- **2026-06-27** — Liquid LFM2.5-230M: 32K context, 293–375MB, 213 tok/s on S25 Ultra [SECONDARY]; LFM Open License v1.0: commercial use only below $10M annual revenue [SECONDARY].
- SmolLM3-3B: fully open recipe (~11.2T tokens, think/no-think modes, 6 languages, 64K extensible to 128K via YaRN) [SECONDARY]; availability established by an Azure deployment guide dated 2026-04-07, but **no release date found** — do not cite as a February–September 2026 launch [DIRECTIONAL].
- **2026-06-08** — Apple Foundation Models 3 at WWDC26: AFM 3 Core ~3B dense on-device (8K context), Core Advanced 20B sparse (1–4B active, NAND-streamed), PrivateCloudCompute server tier (32K), Adapter seam (Session 339); proprietary, SDK-only [SECONDARY].
- **2026-04-29** — IBM Granite 4.1: 3B/8B/30B, Apache 2.0, ~15T tokens, GGUF available; 3B: HumanEval 81.71 / MBPP 71.16; 30B: 512K context [SECONDARY]; Granite 4.2 around **August 2026** [SECONDARY].
- No Phi-5 in the corpus (Azure's 2026 post covers the Phi-4 family as current) [SECONDARY]; no 2026 Command R7B successor (c4ai-command-r7b-12-2024 current) [SECONDARY]; no xAI small/open model (Grok 2.5, August 2025, latest found) [SECONDARY]; Ministral 3 (2026) single-source via aigazine.net [SECONDARY-only] — kept **[UNVERIFIED]**, do not merge with the October 2024 Ministral 3B/8B line.
- Gemma 4 edge QAT builds (gemma4:e2b-it-qat 4.3GB, e4b-it-qat 6.1GB in Ollama) — pointer to §14 [DIRECTIONAL].
- Adjacent community sightings (do not promote without corroboration): Arcee Maestro 7B preview, microsoft/Fara1.5-9B [COMMUNITY].

### Quantization formats and tooling
- **2026-03** — llama.cpp adds GGML_TYPE_NVFP4/MXFP4 (PR #20730); **2026-06** — GPU NVFP4 on sm_120/121 via CUDA dp4a + native FP4 (PRs #20644/#22196), plus SYCL/Vulkan paths [SECONDARY]; new sub-3-bit types Q1_0 (block 128) / Q2_0 (block 64); ternary and Bonsai 1–2-bit experimentation [SECONDARY].
- **2026-04-03** — GPTQModel 6.0.3: ParoQuant, GGUF, FP8, EXL3, FOEM; PrismML/Bonsai 1-bit; Gemma4/GLM4 MoE Lite [SECONDARY]; **2026-04-16** — GPTQModel 6.1.0: JIT CUDA kernels (~300× wheel shrink), Marlin on Turing+, GLM 5/5.1 [SECONDARY].
- llm-compressor v0.13.0: REAP expert pruning; Humming arbitrary bit-widths (3/5/6/7; W2–W8 × A4/A8/A16); observer fusion; XPU support; Llama 4 W4A16; NVFP4 examples; AutoAWQ deprecated in favor of llm-compressor [SECONDARY].
- **2026-07** — Unsloth Qwen3.6-27B-NVFP4: 2.5× throughput claim, 24GB VRAM; official MXFP4/NVFP4 quantization maps [COMMUNITY].
- KV cache: Q8 mainstream on MLX/Ollama; Google TurboQuant 3-bit claim is [SECONDARY] single-source [SECONDARY].

### Serving engines
- **2026-03-02** — vLLM v0.16.0: V1 lockdown — V0 engine core removed, API server forced to V1 [SECONDARY].
- **2026-07-11** — vLLM v0.25.0: Model Runner V2 default for all dense models; legacy PagedAttention removed as default; Transformers backend as fast as native vLLM with FP8 MoE; 558 commits / 232 contributors; new models (GLM-5, DeepSeek-V3.2, MiniMax-M3 NVFP4); Streaming Parser Engine including Kimi k2.5/k2.6/k2.7 parser; universal speculative decoding (TLI) [SECONDARY].
- SGLang: **2026-04-07** — v0.5.10 (GLM-5 NVFP4+MTP stable on Blackwell) [SECONDARY]; **2026-05-16** — v0.5.12 (benchmarked config) [SECONDARY]; **2026-07-25** — v0.5.16 (cutlass NVFP4 MoE backend removed; triton/flashinfer_cutlass remain; flag renames) [COMMUNITY]; day-zero Qwen3.8 hybrid support **2026-08-26** (PRs #36497/#36585) [COMMUNITY].
- **2026-03-31** — Ollama 0.19.0: Apple Silicon switches to the MLX framework; +57% prefill (1,154→1,810) and +93% decode (57.8→112) measured on M5 Max with **Qwen3.5-35B-A3B (~3B active) — do not generalize to dense 35B-class models** [SECONDARY]; >32GB unified memory; Linux/Windows keep llama.cpp; 0.20/0.21 broaden MLX to Gemma 4 + mixed-precision [SECONDARY].
- **~2026-07** — TensorRT-LLM v1.3.0rc21: DeepSeek V4, MiniMax M3 (MXFP8+NVFP4), Gemma 4 12B Unified, Qwen3.5-VL MoE/Dense, Qwen3.6 NVFP4; /v1/embeddings dynamic batching; **breaking**: legacy TRT Python modules removed; SM120 W4A16 NVFP4 Marlin [SECONDARY].

### Embeddings
- **2026-01-15** — Voyage-4 launched: voyage-4 / -large / -lite / -nano plus voyage-multimodal-3.5; first production MoE embedding model; shared embedding space; Matryoshka dims 256/512/1024/2048; MongoDB Atlas GA the same day [SECONDARY]; voyage-4-nano open-weight Apache 2.0 (~600M params, Qwen-based, 32K context) [SECONDARY]; **2026-08-13** — voyage-code-4: trained on issue-fixing PRs; +27.5% on an agentic code-retrieval benchmark, +14.0% across 28 datasets vs voyage-code-3, third of the price [COMMUNITY single-source].
- **2026-02-23** — jina-embeddings-v5 announced (HF org listing 2026-02-18): **small = 677M, nano = 239M — the press copy inverted these** [SECONDARY]; best-in-class among comparable sizes on MMTEB; open weights on HF; self-host via vLLM/llama.cpp/MLX; Elastic Inference Service; license **CC BY-NC 4.0** [SECONDARY]; **2026-05-07** — v5-omni (v5-omni-small/nano, multimodal text+image) [SECONDARY].
- Qwen3-Embedding-8B (June 2025 model): the 2026 delta is **continued MTEB Multilingual lead at 70.58**; 4096-dim MRL; 32K context; Apache 2.0; 100+ languages; 0.6B/4B siblings via Ollama; MTEB-Code 80.68 [SECONDARY].
- MTEB 2026 state: Gemini Embedding 001 English 68.32; gemini-embedding-2 preview March 2026 (multimodal); Voyage family high 60s; EmbeddingGemma-300M on-device; OpenAI text-embedding-3-large 64.6% English [SECONDARY].
- NVIDIA llama-nemotron-embed-1b-v2: February 2026, NVIDIA Open Model License [SECONDARY]; NV-Embed-v2 72.31 (March-2026 MTEB snapshot) [SECONDARY]; Llama-Embed-Nemotron-8B multilingual #1 provisional [SECONDARY].
- Cohere Embed 4 (April 2025) / Rerank 4 (late 2025, 32K context, Fast/Pro tiers): 2025 models; the 2026 event is FedRAMP High on **2026-05-12** [SECONDARY]; Rerank 4 context 32K [SECONDARY] vs 4K [COMMUNITY] — logged, unresolved [DIRECTIONAL].
- No verified 2026 ColBERT successor (jina-colbert-v2, 2024, latest found); no BGE-M3 successor (the GTE line continues via Qwen3-Embedding) [DIRECTIONAL].


