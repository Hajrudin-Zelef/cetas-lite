---
id: ai-industry-kb-2026/08-quantization-model-formats/part-16
title: "8. Quantization & Model Formats (part 16)"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Nvidia", "Z.ai", "vLLM"]
dates: []
keywords: ["quantization", "awq", "benchmark", "blackwell", "compute", "consumer", "deepseek", "fp4", "fp8", "glm", "gptq", "int4"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4337, 4345]
section: "8. Quantization & Model Formats"
sha256: 8563283e696af99ba95d8f48abbcb4f1fbdc7adc79c73c5f5f43d6a5fb1c1a44
---

# 8. Quantization & Model Formats (part 16)

- `llm-compressor` (vLLM project) is the one-command pipeline for FP8/AWQ/GPTQ/MXFP4/NVFP4: FP8 needs no calibration; MoE-aware `--moe` keeps router/gate layers unquantized for stability; one command emits a Hugging Face repo with a generated model card (compressed-tensors safetensors).
- NVFP4 safetensors checkpoints carry `hf_quant_config.json`; vLLM falls back automatically to the Marlin path on pre-Blackwell (the `VLLM_NVFP4_GEMM_BACKEND=marlin` override is deprecated in favor of `--linear-backend`).
- ModelOpt 0.46.0 added **NVFP4 Four-Over-Six (4/6)** and MXFP8 checkpoint support in vLLM; ModelOpt quantization schemes cover INT8, FP8, INT4, NF4, NVFP4 (weight-only, channel, block).
- **Pruning and quantization compose multiplicatively**: Minitron structured pruning + FP8 PTQ on Nemotron-3-Nano-30B = 2.6× throughput and 2.6× memory; **REAP (50% expert pruning) + AutoRound W4A16 took GLM-4.7 from 700 GB to 92 GB (~6.5× total compression)**.
- Pre-quantized NVFP4 examples: DeepSeek-R1-0528, Llama 3, FLUX.1-dev, Qwen3.6-27B/35B-A3B, `ornith-ai/Ornith-1.5-35B-A3B-NVFP4`.
- **DeepSeek-V4-Flash-NVFP4-FP8-MTP hardware scope**: B200/B300 native; on consumer Blackwell vLLM falls back to the Marlin BF16 kernel — use the W4A16 sibling (~159 GB) there instead of the 172 GB NVFP4 build.
- **W4A16 vs W4A4 — check the model card before downloading**: NVFP4 W4A16 runs on Ampere/Hopper via Marlin (`float4_e2m1f` supported even at sm_86) — the memory win without native FP4 compute; **W4A4 needs Blackwell** (native FP4 activations).
- HQQ third-party benchmark figures (~2× speedup on H100/A100 at 4-bit, ~11–13% perplexity increase on Llama-2) come from a community research compilation, not peer-reviewed publication — treat as directional [DIRECTIONAL].

