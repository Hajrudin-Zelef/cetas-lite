---
id: ai-industry-kb-2026/08-quantization-model-formats/kv-cache-quantization-fp8-mainstream-int4-practical-fp4-limi
title: "KV-cache quantization — FP8 mainstream, INT4 practical, FP4 limited"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["Alibaba", "Apple", "Falcon", "Google", "Intel", "Meta", "Microsoft", "Nvidia", "SGLang", "TII", "vLLM"]
dates: ["2026-01", "2026-02", "2026-03", "2026-04", "2026-05-11", "2026-07", "2026-07-15", "2026-08", "2026-09"]
keywords: ["fp4", "fp8", "int4", "quantization", "apache", "attention", "benchmark", "benchmarks", "bitnet", "blackwell", "cost", "decode"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4272, 4301]
section: "8. Quantization & Model Formats"
sha256: 3666ce6742d914939c73e9fae85b9a6f1f72410ac351a9c3ac244d2ec0995b17
---

# KV-cache quantization — FP8 mainstream, INT4 practical, FP4 limited

- **BitNet (Microsoft Research)**: natively trained ternary weights {−1, 0, +1}, i.e. log2(3) ≈ **1.58 bits per parameter**; matrix multiplications become additions/subtractions; activations 8-bit (W1.58A8); custom LUT-based ternary kernels (`bitnet.cpp`).
- Measured: a 2B model fits in **0.4 GB** (vs 4.8 GB FP16); CPU inference at **29 ms/token**; **2.37–6.17× speedup over llama.cpp on x86 CPUs**; **71.9–82.2% energy reduction**.
- At 3B parameters, BitNet b1.58 **matches FP16 LLaMA in perplexity and zero-shot accuracy** while using 3.55× less GPU memory and running 2.71× faster; at 70B, projected throughput is **8.9× higher** than the FP16 baseline.
- Reference model: **BitNet b1.58 2B4T** (2B params, trained on 4 trillion tokens, MIT license); a January 2026 CPU-optimization update added another 1.15–2.1×.
- The headline claim — a theoretical **100B model running on a single CPU at human reading speed (5–7 tok/s)** — remains **theoretical: no such model exists yet** [PARTIALLY VERIFIED].
- Ecosystem reality check (March 2026 deep-dive): 35,000+ GitHub stars but only ~3 active Microsoft maintainers, 269+ open issues; no API server docs, no Docker images, no pip install. Third-party ternary models: TII's **Falcon3-10B-Instruct-1.58bit** (~2 GB GGUF `i2_s`) and **Falcon-E 1B/3B** (natively trained ternary, ~1.5T tokens; 665 MB / 999 MB disk).
- **Bonsai-27B (PrismML, announced 2026-07-15, Apache 2.0)**: derived from Qwen3.6-27B, built on an "Intelligence Density" philosophy — maximum capability per unit of memory/power/footprint.
- **1-bit Bonsai-27B**: "Binary g128" — each weight a single sign bit per group of 128 sharing one FP16 scale → **1.125 bits/weight**; weights occupy **3.9 GB** (vs 54 GB FP16 baseline and 17.6 GB for a conventional 4-bit quant at true 5.2 bits/weight); **>90% of full-precision benchmark performance**; runs at **11 tok/s on an iPhone 17 Pro**.
- **Ternary Bonsai-27B**: 1.58-bit {−1, 0, +1} weights in GGUF `Q2_0_g128` (2-bit slots, FP16 group-wise scaling); deployed size **~7.2 GB** (5.9 GB ideal at 1.71 bits/weight); **>95% of full-precision benchmark performance**.
- **262K-token context practical on-device**: hybrid-attention backbone (~75% linear / ~25% full attention) plus **near-lossless 4-bit KV-cache quantization** — full-attention cache grows on only 16 of 64 layers (~4.3 GB at the full 262K window); measured 100K-token context at 11.6–12.2 GB peak, **~6.8 GB with 4-bit KV**, full 262K window at ~9.4 GB peak.
- **DSpark speculative-decoding drafter**: lossless **1.37× decode speedup** on the CUDA path; ships as a 4-bit pack (1.79 GB Q4_1). Custom low-bit kernels: llama.cpp fork (CUDA + Metal), MLX fork (Apple Silicon), mlx-swift fork (iOS/macOS) — packed weights consumed directly, never expanded to FP16.
- Independent verification (July 2026): a 27B Bonsai ran on a **16 GB M4 Mac mini at 4.174 GB peak RAM and ~21.5 tok/s**, fully in memory — "a 27-billion-parameter model using less memory than a 7B."
- The vLLM GGUF plugin has an experimental `prism-ternary` branch adding Prism Q1_0/Q2_0 group-128 tensor layouts to serve Bonsai 1.7B/4B/27B (tested on RTX 2070 sm_75 and RTX 5060 Ti sm_120).
- **INT2 and below**: 8× memory reduction but **5–15% quality loss**; viable only for narrow tasks — not production-ready for general use as of September 2026. AQLM is cited in 2026 reports as best-in-class at 2-bit but remains the research reference with no production kernel ecosystem [UNVERIFIED].

### KV-cache quantization — FP8 mainstream, INT4 practical, FP4 limited

- **FP8 KV mainstream (production)**: see the Red Hat/vLLM 2026-05-11 finding above; practitioner guides (April/August 2026) call FP8 "the 2026 production default on H100/Blackwell" for weights and KV alike.
- **TurboQuant (Google, ICLR 2026, arXiv:2504.19874)** — detailed in §7: online vector quantization for the KV cache (randomized Hadamard rotation + Lloyd-Max scalar quantization, ~2 bits per coordinate, no calibration); Red Hat independently evaluated it and drove the vLLM integration (PR #38479; presets `turboquant_k8v4`, `turboquant_4bit_nc`, `turboquant_k3v4_nc`, `turboquant_3bit_nc`); study verdict: k8v4 offers no significant advantage over FP8, 4bit-nc is a viable memory-for-throughput tradeoff (up to 3.4× capacity, 1–4 point accuracy cost), aggressive 3-bit variants should be avoided without thorough validation (up to 20-point accuracy drops).
- **OpenVINO 2026.2 added INT4 KV-cache compression on Intel GPU** (INT8 was already the default): on paged attention, keys use per-channel scales (`BY_CHANNEL`, group size 16) and values use per-token scales (`BY_TOKEN`); on SDPA (non-paged) both are per-token, with higher accuracy risk; `i4`/`u4` spellings are equivalent.
- **KVQuant**: per-channel key quantization, **pre-RoPE key quantization** (quantize K before RoPE, then dequantize and apply RoPE), non-uniform per-layer bit allocation (early layers 8-bit, late layers 2–4-bit), per-vector dense-and-sparse (outliers at FP32/FP16, rest INT4; 5–10% sparse overhead).
- **SAW-INT4** (arXiv:2604.19157, April 2026): system-aware 4-bit KV for real serving; token-wise asymmetric INT4 as the foundation (4× BF16 footprint), block-diagonal Hadamard rotation to kill channel-wise outliers without calibration; naive INT4 is shown to degrade severely from channel-range heterogeneity.
- **DeltaKV** (arXiv:2602.08005, February 2026): residual-based compression to **29% of original size (~3.4×)** with near-lossless LongBench/SCBench accuracy, plus Sparse-vLLM (custom engine, 2× vLLM throughput) managing sparse KV layouts.
- **ARKV** (arXiv:2603.08727, February 2026): adaptive per-layer per-token precision across Original/Quantized/Evicted states; **4× memory reduction at ~97% baseline accuracy** on Llama3/Qwen3 long-context benchmarks.
- Field rule: **INT4 KV gives 4× reduction with 1–3% degradation** on standard benchmarks — the same profile as INT4 weights; at INT4, the KV-vs-weights crossover for an 8B model moves from 128K to ~512K context.
- An independent INT4 KV + fused flash-attention CUDA kernel project (2026): keys-only INT4, per-page (256-token) per-channel asymmetric scales (**+2 dB SNR over whole-sequence scales**), nibble-packed storage, **3.9× kernel speedup** to the bandwidth roofline, attention MAE 3e-08 vs FP32, 42/42 GPU tests passing [COMMUNITY].
- **FP4 KV is the limited tier**: SGLang documents FP4/NVFP4 KV-cache paths but they are **Blackwell-gated** (native FP4 tensor cores on B200) and several paths remain experimental; "mainstream generalization" of FP8/FP4 KV overstates the picture — FP8 is the production reality, FP4 is emerging.

### Production serving facts: engines and kernel gates (September 2026)

