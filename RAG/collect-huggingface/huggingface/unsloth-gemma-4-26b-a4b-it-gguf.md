---
id: collect-huggingface/huggingface/unsloth-gemma-4-26b-a4b-it-gguf
title: "gemma-4-26B-A4B-it-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Google", "Hugging Face", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "apache", "attention", "benchmark", "benchmarks", "context window", "inference", "license", "llama", "llama.cpp", "memory", "moe"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-gemma-4-26B-A4B-it-GGUF.md
source_anchor: ""
source_lines: [1, 55]
sha256: 34e433834d0451a4bee45b1b83b8afd65d23dc016d053cd799368f1d001d12b1
---

# gemma-4-26B-A4B-it-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/gemma-4-26B-A4B-it-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth's Dynamic 2.0 GGUF quantization repository for Google DeepMind's Gemma 4 26B A4B instruction-tuned model. Gemma 4 is a family of open multimodal models that handle text and image input and generate text output; this release lists four sizes (E2B, E4B, 26B A4B, and 31B), spanning dense and Mixture-of-Experts (MoE) architectures. The 26B A4B variant is a MoE model with 25.2B total parameters and 3.8B active parameters, 30 layers, a 1024-token sliding window, a 256K-token context window, a 262K vocabulary, and 8 active / 128 total experts plus 1 shared expert. It supports text and image (audio is limited to E2B/E4B in this card), with a ~550M-parameter vision encoder. The model uses hybrid attention interleaving local sliding-window attention with full global attention (final layer always global), and applies unified Keys/Values and Proportional RoPE (p-RoPE) on global layers to optimize long-context memory. Gemma 4 adds configurable thinking modes (enabled via `<|think|>` in the system prompt), native system-prompt support, variable image resolution (token budgets 70/140/280/560/1120), video-as-frames, and native function calling. Benchmark results for 26B A4B: MMLU Pro 82.6%, AIME 2026 (no tools) 88.3%, LiveCodeBench v6 77.1%, Codeforces ELO 1718, GPQA Diamond 82.3%, Tau2 68.2%, HLE no tools 8.7%, BigBench Extra Hard 64.8%, MMMLU 86.3%, MMMU Pro 73.8%, MATH-Vision 82.4%, and MRCR v2 8-needle 128k 44.1%. Recommended sampling: temperature 1.0, top_p 0.95, top_k 64. Quant sizes range from UD-IQ2_XXS (9.92GB) through UD-Q4_K_M (16.9GB), MXFP4_MOE (16.6GB), Q8_0 (26.9GB) to BF16 (50.5GB), with MTP drafter files included. License is Apache 2.0; 576,971 monthly downloads.

## Key points

- Unsloth Dynamic 2.0 GGUF quantization of Gemma 4 26B A4B instruction model (MoE).
- 25.2B total / 3.8B active parameters, 30 layers, 256K context, 8 active / 128 total experts (+1 shared).
- Hybrid attention (1024-token sliding window + global), unified K/V, p-RoPE; text + image modalities.
- Configurable thinking mode, native system prompts, variable image resolution, native function calling.
- Benchmarks: MMLU Pro 82.6%, GPQA-D 82.3%, AIME 2026 88.3%, LiveCodeBench v6 77.1%, MMMU Pro 73.8%.
- Quant sizes from UD-IQ2_XXS 9.92GB to BF16 50.5GB; MTP drafter files (Q8_0 462MB, BF16 855MB) included.
- Runs via llama.cpp, LM Studio, Ollama, vLLM, SGLang; recommended quant UD-Q4_K_M.
- Apache 2.0 license; recommended sampling temperature 1.0, top_p 0.95, top_k 64.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | gemma-4-26B-A4B-it-GGUF |
| Base model | google/gemma-4-26B-A4B-it |
| Architecture | Gemma 4 MoE, hybrid attention |
| Total params | 25.2B |
| Active params | 3.8B |
| Experts | 8 active / 128 total + 1 shared |
| Layers | 30 |
| Sliding window | 1024 tokens |
| Context length | 256K tokens |
| Vocabulary | 262K |
| Modalities | Text, Image |
| Quantization | Unsloth Dynamic 2.0 GGUF (2-bit to 16-bit) |
| Recommended quant | UD-Q4_K_M (16.9GB) |
| Smallest quant | UD-IQ2_XXS (9.92GB) |
| Largest quant | BF16 (50.5GB) |
| MTP drafters | Q8_0 (462MB), BF16 (855MB) |
| Sampling | temperature 1.0, top_p 0.95, top_k 64 |
| License | Apache 2.0 |
| Key benchmarks | MMLU Pro 82.6%; GPQA-D 82.3%; AIME 2026 88.3%; LiveCodeBench v6 77.1% |
| Downloads/month | 576,971 |

## Why this source matters for the RAG

This card documents a widely used local GGUF deployment of Gemma 4 26B A4B, showcasing Unsloth Dynamic 2.0 quantization for a hybrid-attention MoE. It is a key reference for efficient on-device MoE inference and quantized multimodal deployment.
