---
id: ai-industry-kb-2026/08-quantization-model-formats/part-7
title: "8. Quantization & Model Formats (part 7)"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "Apple", "DeepSeek", "Intel", "Nvidia", "vLLM"]
dates: []
keywords: ["quantization", "amd", "awq", "blackwell", "consumer", "datacenter", "deepseek", "distribution", "fine-tuning", "fp4", "fp8", "gguf"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3998, 4012]
section: "8. Quantization & Model Formats"
sha256: fc44c0c6f80eb76c623af93c2c3b6d39063b5a338d07c560dde18cd0cb396832
---

# 8. Quantization & Model Formats (part 7)

| Workload | First choice | Fallback / note |
|---|---|---|
| Local CPU / laptop / Apple Silicon | GGUF Q4_K_M (imatrix) | Q4_K_S for tighter RAM; UD-Q2_K_XL extreme |
| Consumer NVIDIA (peak tok/s, interactive) | ExLlamaV3 + EXL3 | GGUF if distribution/portability matters |
| Consumer NVIDIA (compat, breadth) | GGUF via llama.cpp | Multimodal: `--mmproj-device` for vision projector |
| vLLM Hopper/Ada serving | FP8 static (llm-compressor or ModelOpt) | W4A16 NVFP4 via Marlin where memory-bound |
| vLLM Blackwell datacenter | NVFP4 W4A4 + FP8 E4M3 KV | DeepSeek-R1-FP4 as proof point (see 08b) |
| vLLM Blackwell consumer (GB10/RTX 5090) | NVFP4 W4A4 (no FP4 FLOPS win on SM120) | W4A16 Marlin path for quality-per-watt |
| AMD ROCm serving | FP8 via llm-compressor; GGUF | MXFP4 on MI355X 6.1× (TRT-LLM/ROCm plugin) |
| Intel XPU | AutoRound + vLLM-XPU | Logprob parity verified 91.4% top-1 agreement |
| Single-GPU fine-tuning | bitsandbytes NF4 + QLoRA | HQQ for calibration-free speed |
| Long-context (512K+) | INT4 KV + FP8 weights | see sibling part 08b for KV-cache methods |

- Notes: AWQ/GPTQ remain valid as **compat formats** for NVIDIA pre-Blackwell where no native FP8/FP4 checkpoint exists — but new quantization work should not target them. NVFP4 checkpoint naming: check for W4A16 vs W4A4 on the model card before downloading. For MoE pruning+quant pipelines, use ModelOpt's composed recipes (Minitron + FP8, 2.6× compose on Nemotron-3-Nano-30B).

