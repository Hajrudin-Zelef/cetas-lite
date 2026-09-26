---
id: ai-industry-kb-2026/08-quantization-model-formats/tooling-verdict-table
title: "Tooling verdict table"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "Apple", "Intel", "Nvidia", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM"]
dates: ["2026-08-18", "2026-09"]
keywords: ["amd", "apache", "awq", "blackwell", "cost", "decode", "distillation", "fine-tuning", "fp8", "glm", "gptq", "gpu"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3918, 3942]
section: "8. Quantization & Model Formats"
sha256: 0a90a0ee0c245859f15a986aa60336b650a55c435142398cdeefd8bb37fbb6f9
---

# Tooling verdict table

- **AutoRound** (Intel Neural Compressor): signSGD-based rounding-optimization PTQ, tunes per-weight rounding against a small calibration set at GPTQ-comparable cost. Intel's lm-eval comparisons: superior to GPTQ (30/32 configs) and AWQ (27/32), OmniQuant (16/16). Independent September 2026 test (Intel Arc Pro B70 cookbook, Ornith-1.5-35B-A3B, vLLM-XPU): logprob parity vs BF16 — AutoRound **91.4% top-1 agreement, KL 0.277** vs GPTQ 90.5% / 0.319 — equal-or-best on all primary metrics. Composes with expert pruning: **REAP (50% expert pruning) + AutoRound W4A16 took GLM-4.7 from 700 GB to 92 GB (~6.5× total compression)**. Status: **strongest 4-bit PTQ**.
- **HQQ / gemlite** (mobiusml): data-free, no calibration, quantizes the largest models in minutes; 1/2/3/4/8 bits; linear dequantization compatible with optimized CUDA/Triton kernels; PEFT-compatible; recommended start `nbits=4, group_size=64, axis=1`. **HQQ+** adds trainable low-rank adapters to recover quality at lower bits. [COMMUNITY/DIRECTIONAL] ~2× speedup on H100/A100 at 4-bit with ~11–13% perplexity increase on Llama-2 (community research compilation, not peer-reviewed). **gemlite** (Triton kernel library): 7–8× faster prefill, 3–6× faster decode vs default PyTorch AO kernels; ships MXFP formats end-to-end — A16W8_MXFP, A16W4_MXFP, A8W8_MXFP_dynamic, **A4W4_MXFP_dynamic**, **A4W4_NVFP_dynamic (Blackwell)** — plus HQQ weight-only (4/2-bit) and 1.58-bit ternary; automatic kernel selection per matrix shape with GPU-specific autotuned configs.
- **AQLM** (Egiazarian et al., arXiv:2401.06166): consistently cited in 2026 reports as **best-in-class at 2-bit** [UNVERIFIED — evidence base remains the original 2024 paper plus derivative citations; no 2026 independent production comparison], without production tooling or kernel ecosystem — the research reference for the 2-bit tier.
- **NVIDIA ModelOpt** (`pip install nvidia-modelopt`, Apache 2.0): the toolkit turning a checkpoint into a deployable quantized artifact for TensorRT-LLM, vLLM, and SGLang. **0.46.0 (2026-08-18)** highlights: NVFP4 and FP8 **PTQ recipes** with projection-output quantizers; **Learned Scale Quantization (LSQ) and Dual-LSQ** for quantization-aware distillation; **fused Triton fast path for `local_hessian` NVFP4 weight-scale search — ~34× faster** (bit-exact); **NVFP4 Four-Over-Six (4/6)** and MXFP8 checkpoint support in vLLM via `hf_quant_config.json`. 0.45.0: **pruning and quantization compose multiplicatively** — Minitron structured pruning + FP8 PTQ on Nemotron-3-Nano-30B = 2.6× throughput and 2.6× memory. Supports INT8, FP8, INT4, NF4, NVFP4 (weight-only, channel, block schemes).
- **optimum-quanto: maintenance mode** — the definitive loser call. Its README: "Major new features or breaking changes are unlikely to be merged. For production-ready quantization features or active development, consider alternative projects such as bitsandbytes or torchAO." Remains functional (int2/int4/int8/float8 weights, int8/float8 activations, CUDA kernels) but no longer where HF invests.
- **bitsandbytes**: refactored to multi-backend — **ROCm (AMD) and Intel CPU implementations mature**, Intel XPU in progress, Apple Silicon expected Q4 2026/Q1 2027. NF4 + QLoRA remains the standard single-GPU fine-tuning path.
- **torchAO** (pytorch/ao): HF's recommended active alternative alongside bitsandbytes; torchao int4-weight is the only int4 path on Apple Silicon (Metal).
- The 4-bit PTQ crown moved from AWQ to AutoRound — anyone still calibrating with AutoAWQ is on a deprecated path; the 2026 recommendation is AutoRound (accuracy) or HQQ (speed/data-free), quantized through llm-compressor or ModelOpt.

### Tooling verdict table

| Tool | 2026 status | Niche |
|---|---|---|
| llm-compressor | **Active, recommended** | AWQ/FP8/NVFP4/MXFP4 for vLLM; MoE-aware `--moe` |
| AutoRound | **Strongest 4-bit PTQ** | Best measured accuracy; composes with pruning |
| HQQ / gemlite | Active | Data-free/fast quant; Triton kernels incl. MXFP4/NVFP4 |
| GPTQModel | **Active** | AutoGPTQ successor; 5,000+ HF repos |
| NVIDIA ModelOpt | **Active** | Blackwell-native; NVFP4/FP8 PTQ+QAT recipes |
| bitsandbytes | Active | NF4/QLoRA fine-tuning; multi-backend refactor |
| torchAO | Active | HF-recommended; Metal int4 |
| AutoAWQ | **Deprecated** | Migrated to llm-compressor |
| optimum-quanto | **Maintenance mode** | HF redirects to bitsandbytes/torchAO |
| AQLM | Research | Best-in-class 2-bit, no production ecosystem |
| ExLlamaV2 | **Archived** | Legacy EXL2 checkpoints only |

