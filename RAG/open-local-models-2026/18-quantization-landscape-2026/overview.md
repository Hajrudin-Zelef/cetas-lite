---
id: open-local-models-2026/18-quantization-landscape-2026/overview
title: "3. QUANTIZATION LANDSCAPE 2026"
domain: quantization-landscape-2026
role: deep-dive
task: quantization
actors: ["Apple", "DeepSeek", "Nvidia", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM"]
dates: []
keywords: ["quantization", "attention", "awq", "blackwell", "consumer", "datacenter", "decode", "deepseek", "fp4", "fp8", "gguf", "glm"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [910, 944]
section: "3. QUANTIZATION LANDSCAPE 2026"
sha256: bd5aacc7e65079879f55d59d8183fab4d97e7c9946d4645e742ec731712913e4
---

# 3. QUANTIZATION LANDSCAPE 2026

## 3.1 Format cheat sheet (practical bytes/param and quality story)

| Format | Practical bytes/param | Quality story | Ecosystem |
|---|---|---|---|
| BF16 | 2.0 | training + high-quality baseline | everything |
| FP8 (E4M3) | 1.0 | often ~identical to BF16 (~0.0 ppl drop) | server serving default (vLLM/SGLang) |
| **NVFP4** | ~0.56 | close to FP8 in NVIDIA examples; ~0.3–0.5 ppl drop; competitive with best INT4 | **Blackwell-native**; 16-elem blocks, UE4M3 scales |
| MXFP4 | ~0.53 | useful 4-bit microscaling; ~0.5–1.0 ppl drop; slightly behind NVFP4 | OCP MX standard; 32-elem blocks |
| **GGUF Q4_K_M** | ~0.56–0.65 | strong local-chat default | llama.cpp / Ollama / LM Studio |
| INT4 GPTQ/AWQ | ~0.5 + metadata | calibration-dependent; ~0.3–1.5 ppl drop | legacy/common 4-bit path — "only if that is all you have" |
| GGUF Q8_0 | ~1.06 | ~99% of FP16 | near-lossless local |
| GGUF IQ3/IQ2 (+imatrix) | ~0.35 / ~0.22 | quality drops fast below 3-bit | extreme compression |

*Rule of thumb from the 2026 doctrine: low-batch/single-user → weight-only INT4-class; high-batch serving → W8A8 FP8 (INT4 dequant makes it slower at high batch); KV pressure → FP8 KV cache first (cheapest 2× concurrency).*

## 3.2 What the community actually uses

- **Desktop/local chat:** GGUF **Q4_K_M** remains the default sweet spot — Ollama auto-selects it by VRAM. Q8_0 when VRAM allows. IQ4_XS approaches Q4_K_M quality at slightly less VRAM.
- **Server self-hosting:** **FP8** for <400B models (near-lossless, mature kernels); **NVFP4** where the artifact ships that way (DeepSeek V4.1 Flash ships FP8+FP4 natively; SGLang shipped day-zero GLM-5.2 support including NVIDIA's NVFP4 checkpoint for Blackwell).
- **AWQ/GPTQ:** declining — kept for older model checkpoints that only exist in those formats.

## 3.3 The Blackwell FP4 caveat (important for 2026 buyers)

- NVFP4 dequantization is **free on Blackwell Tensor Cores** (native format); INT4 schemes need a separate dequant kernel that costs throughput.
- **But:** as of 2026, **no off-the-shelf stack delivers NVFP4 weights + NVFP4 KV cache on consumer Blackwell (SM120/RTX 5090/PRO 6000/DGX Spark):** vLLM's NVFP4 path falls back to Marlin (40–50% perf loss); SGLang's attention backends fail for MoE on SM120; TensorRT-LLM lacks NVFP4 KV-cache support. Only custom CUTLASS kernels close the gap. SM100 (datacenter B200) and SM120 (consumer Blackwell) need different kernels — code compiled for `sm_100a` traps on consumer Blackwell.
- Independent test (QuTLASS/MR-GPTQ MXFP4 on RTX 5090): genuine **~4× GEMM speedup** at batch ≥128, but **end-to-end decode lost to BF16 outright** at batch 1–32 (20.4 vs 78.6 tok/s at batch 1) — FP4 wins at GEMM level, loses at memory-bound decode on consumer cards.

**Bottom line for RAG:** recommend NVFP4 artifacts on datacenter Blackwell/Hopper; on consumer Blackwell and Apple Silicon, GGUF Q4_K_M (llama.cpp/Ollama) or MLX 4-bit remain the practical defaults. The FP4 software story on consumer Blackwell is unfinished.

Sources: https://github.com/chsasank/blackwell · https://github.com/0xsero/blackwell-gpu-wiki/blob/HEAD/docs/blackwell/nvfp4-deep-dive.md · https://github.com/kubesimplify/website/blob/HEAD/content/blog/day-4-quantization-demystified-bf16-fp8-nvfp4-mxfp4-int4-gguf-and-why-it-all-matters.md · https://github.com/codehalwell/fable-skills/blob/HEAD/skills/llm-inference-optimization/SKILL.md · https://github.com/notwitcheer/llm-bench-rig/blob/HEAD/reports/fp4-consumer-blackwell.md

---

