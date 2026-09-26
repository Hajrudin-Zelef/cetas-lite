---
id: ai-industry-kb-2026/08-quantization-model-formats/unsloth-nvfp4-and-dynamic-generation-exports
title: "Unsloth NVFP4 and Dynamic-generation exports"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Hugging Face", "Intel", "Nvidia", "Perplexity", "SGLang", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-02-20", "2026-07", "2026-08", "2026-08-15", "2026-08-18", "2026-08-19", "2026-08-20", "2026-08-27", "2026-09-22"]
keywords: ["fp4", "nvfp4", "agentic", "amd", "awq", "benchmark", "blackwell", "compute", "cost", "decode", "deepseek", "distillation"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4302, 4336]
section: "8. Quantization & Model Formats"
sha256: bfab37e0199093508c930d532e30f8ecc648441095e8268bea35ce0c37921c12
---

# Unsloth NVFP4 and Dynamic-generation exports

- vLLM quantization→backend mapping (Isambard, August 2026): FP8 → `deep_gemm`/`hpc`/`flashinfer_cutlass`; INT4 AWQ/GPTQ → `marlin`; NVFP4 → `flashinfer_cutedsl`; W4A8 → `humming`.
- **CUTLASS native FP8 GEMM needs compute capability ≥ 89**; FP8 *weights* (W8A16) via Marlin-FP8 work from SM75.
- Hard-won deployment rules: Marlin rejects symmetric AWQ `uint4` with `zero_point=false` (AutoRound quants) on Ampere — use a compressed-tensors quant instead; Marlin needs layer dims divisible by 128 (multimodal vision MLPs with odd dims need Hopper/Machete); AWQ `quant_method` requires `--dtype float16` (bf16 rejected at config validation).
- On AMD, AWQ/GPTQ/Marlin remain NVIDIA-only — the ROCm paths use FP8 via llm-compressor, GGUF, or AMD Quark. SGLang also loads GGUF via `--load-format gguf`, but NVIDIA-only.
- **TensorRT-LLM** supports FP8, NVFP4 and MXFP4 via ModelOpt; compiled per-GPU; W4A4 needs Blackwell.
- **NVIDIA ModelOpt 0.46.0 (2026-08-18)**: NVFP4 and FP8 PTQ recipes with projection-output quantizers; **Learned Scale Quantization (LSQ) and Dual-LSQ** for quantization-aware distillation; fused Triton fast path for `local_hessian` NVFP4 weight-scale search — **~34× faster** than the Python reference sweep (bit-exact); **NVFP4 Four-Over-Six (4/6)** and MXFP8 checkpoint support in vLLM via `hf_quant_config.json` (`quantization="modelopt_fp4"`, `quantization="modelopt_mxfp8"`).
- **bitsandbytes** multi-backend refactor: **ROCm (AMD) and Intel CPU mature**, Intel XPU in progress, Apple Silicon expected Q4 2026/Q1 2027; **NF4 + QLoRA** remains the standard single-GPU fine-tuning path.
- **torchAO** is Hugging Face's recommended active alternative alongside bitsandbytes; torchao int4-weight is the only int4 path on Apple Silicon (Metal).
- Hugging Face's **optimum-quanto** is in **maintenance mode** (confirmed 2026-09-22): "Major new features or breaking changes are unlikely to be merged. For production-ready quantization features or active development, consider alternative projects such as bitsandbytes or torchAO."
- **gemlite** (mobiusml's Triton kernel library): 7–8× faster prefill and 3–6× faster decode vs default PyTorch AO kernels; ships **MXFP formats end-to-end** — A16W8_MXFP, A16W4_MXFP, A8W8_MXFP_dynamic, **A4W4_MXFP_dynamic**, **A4W4_NVFP_dynamic (Blackwell)** — plus HQQ weight-only (4/2-bit) and 1.58-bit ternary.
- **s0me1-dev Blackwell GeForce NVFP4 GEMM** patches for vLLM/FlashInfer/CUTLASS on SM120: Qwen3.6-35B-A3B at 175 tok/s [COMMUNITY].
- **ggml.ai joined Hugging Face — announced 2026-02-20** (Discussion ggml-org/llama.cpp #19759); projects stay 100% MIT-licensed, Gerganov's team keeps full technical leadership; llama.cpp passed **126,000+ GitHub stars** in 2026. The HF absorption cements GGUF as the de-facto local-inference format.

### Unsloth NVFP4 and Dynamic-generation exports

- **Since July 2026, Unsloth exports "Dynamic Unsloth NVFP4 Quants"** ("faster and more accurate") for the NVIDIA stack — `gemma-4-31B-it-NVFP4`, `gemma-4-26B-A4B-it-NVFP4`, `gemma-4-12b-it-NVFP4` — plus hybrid **NVFP4-GGUF** containers (Qwen3.6 NVFP4-GGUF: a GGUF container holding NVFP4-packed weights, accelerated only via Unsloth kernels on Blackwell).
- **Unsloth Dynamic v3.0 launched 2026-08-19** (alongside **Qwen3.8-27B**, launched 2026-08-15): claims **>10% better top-1 accuracy at the same size than every other provider** [VENDOR]; Qwen3.8-27B v3.0 quants recorded **5.1 million downloads in 5 days**; pure PTQ (no QAT, no QAD); imatrix published for community testing.
- UD variant ladder: `UD-TQ1_0` (~1-bit; DeepSeek-V3.1 671B → ~192 GB); `UD-IQ2_M` (~2-bit); `UD-Q2_K_XL` (~2.7-bit, recommended; GLM-5.1 → ~220 GB); `UD-Q3_K_XL` (~3-bit; Qwen3.8-27B ~92.4% top-1 BF16 agreement); `UD-Q4_K_XL` (~4-bit; Qwen3.8-27B ~96% top-1 BF16 agreement, 16.69 GiB) — agreement figures are third-party graph readings of Unsloth's fidelity curve, proxies only [VENDOR].
- First independent test (srmiles, 2026-08-27, Qwen3.8-27B): v3.0 measurably closer to the Q8_0 referee on mean/median/tail KLD *at smaller size* — corroborating the vendor claim's direction (KLD-based, not benchmark-score-based); but it loses on 99.0% Δp and top-1 agreement at 4-bit, **decodes 9.5% slower** at Q4_K_XL, and `UD-Q3_K_XL` wedged llama-server (token id −1; may be fixed in later re-uploads) [COMMUNITY].
- The Unsloth Hugging Face org publishes **1,374 models** (GGUF + NVFP4).

### Measuring quality: perplexity, KL divergence, benchmark reading

- **Perplexity delta vs the FP16 baseline** is the standard first check (tables in Figures). Rule of thumb: <0.1 ≈ indistinguishable, 0.1–0.2 ≈ production-safe, >0.5 ≈ noticeable.
- **KL divergence** between quantized and full-precision output distributions is the stricter test Unsloth cites for Dynamic v3.0 — it catches distribution shifts that perplexity averages hide. Artefact2's per-type table records median KL, q99 KL, fraction of top tokens differing, and ln(PPL ratio) per quant type — the pin-to-numbers reference.
- Divergence-300 @32 (Unsloth's Dynamic 3.0 metric) has **no independent replication** as of 2026-09-22; the srmiles test corroborated KLD direction vs a Q8_0 referee only, on one model, one hardware.
- **Benchmark retention** (% of FP16 score): Bonsai-27B ternary >95%, 1-bit >90%; GLM-5.2 Dynamic 2-bit ~82% top-1 accuracy at 84% smaller. Aggregate benchmark scores can lie — a format can hold MMLU while degrading long-form coherence — so per-task evaluation (especially coding: the Qwen-Coder FP8 data) is the reliable signal.
- Unsloth's third-party top-1 BF16 agreement figures (UD-Q4_K_XL ~96%, UD-Q3_K_XL ~92.4% on Qwen3.8-27B) are **readings of the published fidelity curve** — quant-fidelity proxies only; do not multiply them into downstream task scores [VENDOR].
- **QAT vs PTQ at INT4**: 3–8% drop (PTQ) vs <1% (QAT) — the single largest quality lever at 4-bit, at the cost of retraining compute and format lock-in.
- **Importance matrices (imatrix)** materially improve 1–3-bit quants: `llama-imatrix` computes per-weight importance from calibration text; `llama-quantize --imatrix` uses it during quantization; all IQ types need an imatrix for best results (the `i1-` prefix marks imatrix-weighted quants).
- Dynamic 3.0's imatrix: **>1.5M tokens from diverse sources**, explicitly refined for **agentic coding, chat, and multilingual performance**; Unsloth warns that "instruct models have unique chat templates, and using text-only calibration datasets is not effective for instruct models" — corroborated by the independent malaiwah analysis (2026-08-20).
- FP8 E4M3's maximum relative quantization error is 6.25% per value (half a step of the 3-bit mantissa); per-row/per-head scaling keeps average error well below that.

### Conversion and checkpoint notes

