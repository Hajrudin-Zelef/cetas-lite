---
id: collect-huggingface/huggingface/poolside-laguna-s-2-1-nvfp4
title: "Laguna S 2.1-NVFP4 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Apple", "DeepSeek", "Hugging Face", "Nvidia", "SGLang", "vLLM"]
dates: ["2026-07", "2026-08", "2026-09-23"]
keywords: ["nvfp4", "agentic", "attention", "benchmark", "benchmarks", "consumer", "context window", "decode", "deepseek", "embeddings", "fp8", "gguf"]
source: docs/RAG/Collect RAG/03_huggingface/poolside-Laguna-S-2.1-NVFP4.md
source_anchor: ""
source_lines: [1, 50]
sha256: 7a44a801a8187e16421c5098a0302ba76dc80aacee467e81c22cff2133aff7ae
---

# Laguna S 2.1-NVFP4 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/poolside/Laguna-S-2.1-NVFP4
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Laguna S 2.1-NVFP4 is poolside's NVFP4 (NVIDIA 4-bit floating-point) quantized checkpoint of the Laguna S 2.1 Mixture-of-Experts model — a 117.6B total parameter / 8.5B activated parameter text-to-text model designed for agentic coding and long-horizon work on a local machine. The weights are roughly 71 GB in NVFP4. The architecture uses a mixed attention layout across 48 layers: 36 layers use Sliding Window Attention (512-token window) with per-head gating and 12 layers use global attention (3:1 SWA:global ratio), with softplus gating and per-layer rotary scales for fast inference and low KV cache requirements. KV cache is quantized to FP8 to reduce memory per token. It supports interleaved thinking between tool calls (native reasoning support, enable/disable thinking per request). Training comprised pre-training, post-training and RL stages with the Muon optimizer, using 256 experts plus 1 shared expert. The model is released under the OpenMDW-1.1 license and is an updated checkpoint (August 2026) that supersedes the earlier version of the repository.

The checkpoint ships configured for a 1,048,576-token (1M) context window; quantization was calibrated at 1M, and a 256K cap can be set via `config.json` (`rope_parameters` factor 32.0, attention_factor 1.3465735902799727, max_position_embeddings 262144). Sampling defaults in `generation_config.json` are authoritative (`top_k 20` eval-certified truncation). Benchmark results (as of 21 July 2026) show Laguna S 2.1 at Terminal-Bench 2.1 70.2%, SWE-bench Multilingual 78.5%, SWE-Bench Pro (Public Dataset) 59.4%, DeepSWE 40.4%, SWE Atlas (Codebase QnA) 46.2%, and Toolathlon Verified 49.7%, competitive with Tencent Hy3, Nemotron 3 Ultra, DeepSeek-V4-Pro Max and others.

Deployment is supported in vLLM (0.25.0+), Transformers and TRT-LLM (>=1.3.0rc16), with quantization auto-detected from `quantization_config`. NVFP4 does not run correctly on SGLang. For local machines, Ollama and llama.cpp (BF16 and Q4_K_M only) are recommended; on a DGX Spark (GB10) or Apple Silicon Mac Studio, Ollama runs the Q4_K_M GGUF (~75 GB) at ~12.6 tok/s on the Spark and ~17.6 on Apple Silicon. A detailed vLLM recipe on the DGX Spark gives native NVFP4 + DFlash speculative decoding (paired with poolside/Laguna-S-2.1-DFlash-NVFP4): prefill 600–800 tok/s, decode ~15 tok/s prose / 22–24 code, DFlash accepting 2.9–3.1 tokens/step. The card also covers controlling reasoning and responsible-use guidance.

## Key points

- 117.6B total / 8.5B active MoE; NVFP4 weights ~71 GB, aimed at local agentic coding.
- Mixed attention: 36 SWA layers (512-token window) + 12 global layers; FP8 KV cache.
- Native reasoning: interleaved thinking, per-request toggle.
- 1M context configured by default; 256K via config.json override.
- OpenMDW-1.1 license; updated August 2026 checkpoint.
- vLLM/Transformers/TRT-LLM supported; NVFP4 unsupported on SGLang.
- DGX Spark (GB10) / Mac Studio recipes: Q4_K_M ~75 GB, 12.6–17.6 tok/s.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | poolside |
| Model name | Laguna S 2.1-NVFP4 |
| Architecture | laguna MoE, 48 layers (36 SWA + 12 global), 256 experts + 1 shared |
| Total params | 117.6B |
| Active params | 8.5B |
| Hub model size | 118B params |
| Context length | 1,048,576 tokens (default); 262,144 via config |
| Quantization | NVFP4 (BF16 + U8 tensors on hub); FP8 KV cache |
| Optimizer | Muon |
| License | OpenMDW-1.1 |
| Key benchmarks | Terminal-Bench 2.1 70.2%, SWE-bench Multilingual 78.5%, SWE-Bench Pro 59.4%, DeepSWE 40.4%, SWE Atlas 46.2%, Toolathlon Verified 49.7% |
| Downloads/month | 439,146 |
| Serving | vLLM ≥0.25.0, Transformers, TRT-LLM ≥1.3.0rc16; Ollama, llama.cpp (BF16/Q4_K_M) |

## Why this source matters for the RAG

This card is a high-value reference for 4-bit (NVFP4) frontier MoE serving on consumer/edge hardware, combining full architecture details (mixed SWA/global attention, per-head gating, Muon training), a 1M context story, and extremely detailed DGX Spark/Apple Silicon deployment measurements. It also provides citable agentic-coding benchmark comparisons with other open frontier models.
