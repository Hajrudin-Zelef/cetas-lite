---
id: ai-industry-kb-2026/08-quantization-model-formats/key-dated-facts-continued
title: "Key dated facts (continued)"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "Alibaba", "Hugging Face", "MiniMax", "Moonshot", "Nvidia", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM"]
dates: ["2026-02-20", "2026-05-11", "2026-07-06", "2026-08-21"]
keywords: ["acquisition", "awq", "benchmarks", "blackwell", "cost", "decode", "fp8", "gguf", "glm", "gptq", "gpu", "gpus"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4196, 4233]
section: "8. Quantization & Model Formats"
sha256: 095bfb83ca13c64ed956d6bf3fbc69397ccde53c8d69be622d0e5351d8b557cf
---

# Key dated facts (continued)

- [COMMUNITY] ModelOpt 0.45.0 deep dive (SMF Works; pruning+quant 2.6× compose, Nemotron-3-Nano-30B) — https://github.com/smfworks/aiclearinghouse-site/blob/HEAD/content/blog/2026-07-06-nvidia-model-optimizer-0-45-deep-dive.md
- [COMMUNITY] ModelOpt checkpoint formats in vLLM (FP8 variants, NVFP4, MXFP8; --linear-backend) — https://github.com/guqiong96/lvllm/blob/HEAD/docs/features/quantization/modelopt.md
- [VENDOR] optimum-quanto README (maintenance mode; HF redirects to bitsandbytes/torchAO) — https://github.com/huggingface/optimum-quanto/blob/HEAD/README.md
- [VENDOR] Hugging Face quantization overview (GPTQModel vs AutoGPTQ, bitsandbytes multi-backend, torchao Metal) — https://huggingface.co/docs/transformers/v4.49.0/quantization/overview
- [COMMUNITY] vLLM v0.26.0 capability gates (Marlin SM75 floor, Machete SM90-exact, Conch SM80+, CUTLASS FP8 ≥89) — https://github.com/adargit008/mcgyvr/blob/HEAD/okf/must-read/touching-engine.md
- [COMMUNITY] vLLM quant deployment rules (MXFP4 Hopper+ only, NVFP4 W4A16 via Marlin on Ampere, AWQ --dtype float16) — https://github.com/billyhargroveofficial/billy-skills-hub/blob/HEAD/cloud-gpu-vllm-uncensored/SKILL.md
- [COMMUNITY] vLLM backend selection (Isambard, Aug 2026; quant→backend mapping, W4A8 humming, NVFP4 flashinfer_cutedsl) — https://github.com/ai4ci/ivllm/blob/HEAD/design/references/vllm-backend-selection.md
- [VENDOR] HQQ (mobiusml; data-free, 1/2/3/4/8-bit, HQQ+, PEFT) — https://github.com/mobiusml/hqq/blob/HEAD/Readme.md
- [VENDOR] gemlite (mobiusml; Triton kernels, MXFP4/NVFP4 dynamic, 7–8× prefill, 3–6× decode) — https://deepwiki.com/mobiusml/gemlite
- [COMMUNITY] HQQ quantization benchmarks (ai-atlasforge; A100 numbers, INT8 KV for 8K+ ctx) — https://github.com/dragonshadows1978/ai-atlasforge/blob/HEAD/HQQ_QUANTIZATION_BENCHMARKS_RESEARCH.md
- [VENDOR] llama.cpp upstream digest (2026-08-21; v0.2.0 semver, Kimi K3, ROCm/Vulkan, DSpark, --mmproj-device) — https://github.com/licjon/cl-llama-cpp/blob/HEAD/docs/upstream-digest.md
- [COMMUNITY] llama.cpp ROCm 6.4.4 (dasroot; Qwen3-235B-A22B 1132 tok/s on Radeon 8060S, 3.47× vs ROCm 7.0.1) — https://dasroot.net/posts/2026/01/running-local-llms-python-ollama-llama-cpp-transformers/
- [COMMUNITY] Production quantization wiki (epyc-root; Q4_K_M production standard, REAP+AutoRound GLM-4.7 700→92GB) — https://github.com/pestopoppa/epyc-root/blob/HEAD/wiki/quantization.md
- [COMMUNITY] Quantization inference guide 2026 (GPTQ/AWQ derivations, INT3/INT2 pitfalls) — https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Inference_Optimization/quantization_inference.md
- [COMMUNITY] LLM quantization guide 2026 (ayinedjimi; 4-bit MMLU 2–4 pts, AWQ vs GPTQ perplexity) — https://ayinedjimi-consultants.fr/static/pdf/quantization-llm-2026-gguf-gptq.pdf
- [VENDOR] Simon Willison — ggml.ai joins Hugging Face (2026-02-20) — https://simonwillison.net/2026/Feb/20/ggmlai-joins-hugging-face/
- [VENDOR] GitHub — ggml.ai joins Hugging Face, llama.cpp Discussion #19759 — https://github.com/ggml-org/llama.cpp/discussions/19759
- [COMMUNITY] InsiderLLM — what the Hugging Face acquisition means — https://insiderllm.com/guides/llamacpp-hugging-face-ggml-acquisition/
- [UNVERIFIED] NVIDIA offered $12.9B to acquire Hugging Face (single-source secondary outlet) — no corroboration; claim treated as unverified
- [VENDOR] KDnuggets — Quantization and pruning methods to make your LLM leaner — https://www.kdnuggets.com/quantization-and-pruning-methods-to-make-your-llm-leaner

## Key dated facts (continued)

### FP8 — the 2026 production default (near-lossless)

- FP8 E4M3 weight quantization shows **<0.1 perplexity increase** on Llama-family and Qwen models; measured on Qwen-Coder 32B at **79–80% on coding tasks vs ~80% in FP16 (<1% drop)**; FP8 E5M2 is slightly worse and less stable (78–79%, −1–2%).
- 2026 production guidance: **FP16→FP8 on H100 gives 30–50% inference speedup at <0.1% quality loss** (reported perplexity delta 0.01); "the no-brainer quantization; every H100 deployment should use it" [COMMUNITY].
- Measured: Qwen-Coder 32B in FP8 runs **~100–120 tok/s on RTX 6000 Ada (1.7–2× vs FP16)** thanks to native Tensor Core support; on RTX 4090 (no native FP8) only ~60–80 tok/s (0.9–1.3×, emulated via casting — no speedup).
- VRAM: Qwen-Coder 32B FP8 uses **~32–36 GB (50% reduction vs ~65 GB FP16)**; summary rule: **FP8 = 1 byte/param, near-lossless on most models**.
- Two formats: **E4M3** (4 exponent, 3 mantissa bits — preferred for inference) and **E5M2** (5 exponent, 2 mantissa — wider range, less precision).
- **No calibration data required** (unlike GPTQ/AWQ); model authors increasingly publish pre-quantized FP8 checkpoints via llm-compressor / AutoFP8.
- FP8 KV cache has minimal generation-quality impact with per-head dynamic scaling (~3% scale-factor memory overhead per token); E4M3's maximum relative quantization error is 6.25% per value; main risk is outlier activations exceeding the [−448, 448] range getting clamped.
- Hardware: native FP8 Tensor Cores on H100/H200 (Hopper), L40S / RTX 6000 Ada (Ada), and all Blackwell GPUs; **not** on Ampere/A100 or RTX 4090 (emulated path only).
- vLLM has an FP8 kernel since 0.4.1+ (requires Ada-or-newer GPU); **Ollama and llama.cpp have no FP8 support** — near-lossless local inference uses GGUF Q6_K/Q8_0 instead.
- FP8 KV verified mainstream in production: the **Red Hat / vLLM study published 2026-05-11** (Llama-3.3-70B-Instruct, two Qwen3-30B-A3B models, MiniMax-M2.7) explicitly recommends `--kv-cache-dtype fp8` as the default — 2× capacity, no throughput cost, negligible accuracy loss, significantly better TTFT under burst load; vLLM, TensorRT-LLM and SGLang all ship FP8 KV-cache paths.

### INT4 — AWQ, GPTQ, Marlin kernels, QAT (cross-ref: full tooling detail in part 08a)

