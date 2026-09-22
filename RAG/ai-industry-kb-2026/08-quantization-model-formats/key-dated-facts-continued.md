---
id: ai-industry-kb-2026/08-quantization-model-formats/key-dated-facts-continued
title: "Key dated facts (continued)"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Falcon", "Google", "Hugging Face", "Intel", "Meta", "Microsoft", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "Perplexity", "Qualcomm", "SGLang", "TII", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM"]
dates: ["2023-09", "2025-05", "2026-01", "2026-02", "2026-02-13", "2026-02-20", "2026-03", "2026-04", "2026-05", "2026-05-11", "2026-05-13", "2026-07", "2026-07-15", "2026-08", "2026-08-15", "2026-08-18", "2026-08-19", "2026-08-20", "2026-08-27", "2026-09", "2026-09-22"]
keywords: ["agentic", "amd", "apache", "attention", "awq", "benchmark", "benchmarks", "bitnet", "blackwell", "compute", "consumer", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4217, 4345]
section: "8. Quantization & Model Formats"
sha256: 706b3e4db32618355931ac083d01c21c3c7f6c95cac4197016c4837e948f9847
---

# Key dated facts (continued)

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

- **FP16→INT4 (GPTQ/AWQ) = 4× memory reduction, 1–3% quality loss** — the format that puts a 70B dense model on a single 48 GB GPU.
- AWQ (Lin et al., arXiv 2306.00978): scales weights by activation importance before quantization; W4A16, group size 128; consensus best INT4 PTQ quality; official kernel is the default vLLM path, Marlin/Machete when large-batch throughput matters.
- GPTQ: column-by-column inverse-Hessian error compensation; requires calibration (~5–30 minutes for a 7B model); slightly worse quality than AWQ on most benchmarks but widest hardware compatibility.
- Marlin family (fused dequantize+GEMM for weight-only 4-bit): **Marlin's floor is SM75 (Turing), not SM80** — audited in vLLM v0.26.0 source (`if device_capability < 75`); **Marlin-24 act-order kernels were removed from vLLM main**; Machete is SM90-exact (Hopper); Conch targets SM80+.
- QAT vs PTQ at INT4: **PTQ quality drop 3–8% vs QAT <1%**; cost is full retraining compute plus format lock-in. **Kimi K2.6 ships INT4 natively via QAT (~2× speedup)**; **DeepSeek-V4 applies FP4 (MXFP4) QAT to MoE experts plus the indexer QK path** — the parameter-heavy parts.
- 2026 tooling shakeout (detail in part 08a): **AutoAWQ is officially deprecated** — the AWQ workflow moved to llm-compressor; AutoGPTQ's successor is **GPTQModel** (5,000+ GPTQ repos on Hugging Face reference it); **AutoRound (Intel)** is the strongest measured 4-bit PTQ (beats GPTQ 30/32 configs and AWQ 27/32 on lm-eval; September 2026 independent logprob test: 91.4% top-1 agreement vs BF16, KL 0.277, vs GPTQ 90.5% / 0.319); **HQQ** owns the data-free/fast niche (1/2/3/4/8 bits, no calibration).
- AWQ deployment nuance (September 2026): vLLM and SGLang both ship awq/awq_marlin (and gptq_marlin) paths, and AWQ is described as "the more commonly used of the two" vs GPTQ on vLLM deployments — but workflow friction on 2026 models is real (AutoAWQ was archived in May 2025 and never got Qwen3.5 support; llm-compressor's pip release pins `transformers<=4.57.6`, predating Qwen3.5 support; AWQ's layer-by-layer scale search OOMs on a 40 GB A100 — fix: `offload_device=cpu`); and a community 4× RTX PRO 6000 Blackwell shootout had **AWQ beating both NVFP4 recipes on decode throughput at every concurrency** (C=128: AWQ 3519 vs 3232/3220 tok/s with MTP) [COMMUNITY, single rig — NVFP4 backend/kernel version-sensitive]. Net: dominant *compat* INT4 path, not always the fastest.

### FP4 — NVFP4 and MXFP4, the 2026 frontier

- Why FP4 beats INT4 at equal bit-width: floating-point spacing is logarithmic — high resolution near zero, lower at extremes — matching near-Gaussian weight distributions; a block-level shared scale adapts to local magnitude.
- **NVFP4 anatomy**: element FP4 E2M1, **block size 16** (finer granularity), per-block scale in **FP8 E4M3**, plus an additional per-tensor **FP32 global scale** (block scale capped at ±448); total ≈ 4 + 8/16 ≈ **4.5 bits/element**; NVIDIA-proprietary. NVIDIA's two-level scaling strategy preserves dynamic range far more faithfully than flat INT4.
- Native execution on **datacenter Blackwell** (SM100/SM103: B100/B200/GB200/GB300) via `tcgen05` FP4 tensor-core instructions — roughly **2× the FP8 GEMM compute**; gate: CUDA ≥ 12.8. Claimed B200 ≈ 8× FP16 tensor-core throughput [VENDOR].
- **Marlin weight-only fallback**: packed FP4 weights load on **SM75+** (Turing and later, including Ampere/Ada/Hopper) and dequantize inside the kernel — the memory win without the FP4 speed edge (W4A16-class throughput); prefer FP16 activations on this path (BF16 + Marlin has reported garbled output). The old `VLLM_NVFP4_GEMM_BACKEND` env var is deprecated, replaced by `--linear-backend`; vLLM now auto-selects CUTLASS/FlashInfer/Marlin at load.
- **DeepSeek-R1-FP4 (671B)**: MMLU **90.7 vs 90.8** on the FP8 base (−0.1), GSM8K **96.1 vs 96.3** (−0.2); NVIDIA Model Optimizer PTQ to FP4 E2M1 (weights + activations of linear operators; attention/embeddings at higher precision); the method itself is a **proprietary black box** — proves W4A4 solvable at massive scale, but the method is not open.
- **NVIDIA's Kimi-K2.6-NVFP4, posted 2026-05-13**: Moonshot AI's Kimi-K2.6 (1T total parameters, 32B active, 256K context, text/image/video), quantized with ModelOpt 0.44.0 via an INT4→BF16→NVFP4 path, served through vLLM, tested on B200.
- **NVIDIA's official NVFP4 guidance** (build.nvidia.com, DGX Spark guide): cut memory **~3.5× vs FP16** and **~1.8× vs FP8**, maintain accuracy close to FP8 (**usually <1% loss**), mixed-precision execution with accumulation in higher precision; a DGX-Station playbook quantizes DeepSeek-R1-Distill-Llama-8B to NVFP4 via TensorRT Model Optimizer.
- **Nemotron-3-Nano-30B NVFP4 via quantization-aware distillation (QAD)**: frozen BF16 teacher, KL divergence between token distributions — no replay of SFT/RL stages — reaches **99.4% of the BF16 baseline** on reasoning/coding benchmarks; arithmetic throughput 2–3×, memory ~1.8× smaller, "fourfold speed gains" [VENDOR].
- **DeepSeek-V3.2 / R1 NVFP4 on GB300 (SM103)** (vLLM project blog, 2026-02-13; stack: vLLM v0.14.1, CUDA 13.0): V3.2 NVFP4+TP2 → **7,360 TGS prefill-only**, 2,816 TGS mixed (ISL=2k/OSL=1k); R1 NVFP4+EP2 on 2× GB300 → **22,476 TGS prefill-only**, 3,072 TGS mixed; vs Hopper: **8× prefill, 10–20× mixed-context**.
- Private-Blackwell RAG study (vLLM + AIPerf): **NVFP4 gave 1.6× throughput vs BF16 alongside 41% energy reduction, with 2–4% model quality loss**; self-hosted costs $0.001–$0.04/MTok (electricity only), 40–200× below budget cloud APIs; hardware break-even under 4 months at 30M tokens/day [COMMUNITY/SECONDARY].
- **Nemotron-3.5-Lightning-30B-A3B-NVFP4 on vLLM 0.27.1** (massed-compute, 2026-08-18): L40S 1345.4 tok/s @ $0.182/1M out tokens; **RTX PRO 6000 Blackwell 2541.1 tok/s** @ $0.239/1M; A100 2100.8 tok/s @ $0.182/1M — L40S/A100 via W4A16 `humming` + `--quantization modelopt_fp4`, Blackwell on the native FP4 path + FP8 KV; native path wins throughput, best $/token tied on L40S/A100 [COMMUNITY].
- **The FP4 training paradox (2026)**: nanochat 560M pretraining on a single DGX Spark measured **BF16 ~17,500 tok/s (~11 days) vs NVFP4 ~13,000 tok/s (~14 days)** — for that recipe on that hardware BF16 was faster. **NVFP4 is an inference format; it is not a training default** [COMMUNITY].
- Consumer Blackwell (SM120): NVFP4 *weights* run (RTX 5090, SM120); native NVFP4 MoE matmul executes only on B200/B300 — on consumer Blackwell vLLM falls back to the Marlin BF16 kernel, forfeiting the FP4 FLOPS advantage (measured throughput parity with the smaller ~159 GB W4A16 sibling, recommended there). Stock vLLM **NVFP4 KV-cache/attention path is datacenter-only** (crashes on RTX 5090; stock consumer-Blackwell FP4-era KV is `fp8_e4m3`); a community route (FA2 prefill + XQA decode) reached genuine FP4 KV on sm_120 at **1.6× the fp8 KV pool** (July 2026).
- GB10 production datapoint (vLLM PR #29242, avarok): **NVFP4 W4A4 MoE on Blackwell GB10 is production-ready** — 4× memory reduction (**160 GB → 40 GB**), single-GPU 80B-parameter deployment, 128K context.
- **DeepSeek-V4-Flash-NVFP4-FP8-MTP** (community quant, May 2026): NVFP4-routed experts + FP8 block-128 attention + BF16 multi-token-prediction head retained for vLLM speculative decoding; **172 GB** total; native NVFP4 MoE matmul only on **B200/B300**.
- **Unsloth NVFP4 export since July 2026**: "Dynamic Unsloth NVFP4 Quants" — e.g. `gemma-4-31B-it-NVFP4`, `gemma-4-26B-A4B-it-NVFP4`, `gemma-4-12b-it-NVFP4` — plus hybrid **NVFP4-GGUF** containers (Qwen3.6 NVFP4-GGUF: a GGUF container holding NVFP4-packed weights, accelerated only via Unsloth kernels on Blackwell).
- Pre-quantized NVFP4 checkpoints: DeepSeek-R1-0528, Llama 3, FLUX.1-dev, Qwen3.6-27B/35B-A3B, `ornith-ai/Ornith-1.5-35B-A3B-NVFP4` — the latter loads on an **8 GB-class Ada GPU (SM89) via the Marlin fallback** (~4.5 bits/value) at W4A16-class speed, vs 21.7 GB for the GGUF Q4_K_M of the same model.
- The canonical 2026 multi-format production stack: **NVFP4 weights (~4.5 bits/elem) + FP8 E4M3 KV cache + BF16/FP8 prefill activations + BF16 decode activations + FP32 accumulators** (LayerNorm/softmax/final logits all FP32); SM100 has dedicated NVFP4→BF16 dequantization in `tcgen05`, SM120 does it in software; at FP8 and below overflow is a real concern (448 max for E4M3) — per-block scaling is the standard mitigation, and production FP4 does not generally use stochastic rounding.
- **MXFP4 (OCP microscaling standard)**: OCP MX specification v1.0 (**September 2023**), contributions from AMD, Arm, Intel, Meta, Microsoft, NVIDIA and Qualcomm — cross-vendor, unlike NVFP4. Format: **E2M1** elements in blocks of **32** sharing one **8-bit E8M0 scale**; a 32-element block needs 136 bits (17 bytes) vs 1,024 bits (128 bytes) FP32 — **7.5× storage reduction**; GPT-2 perplexity 18.4 (FP32) → 18.7; ResNet-50 top-1 75.9% vs 76.1% baseline.
- **OpenAI's GPT-OSS models ship natively in MXFP4 (2025-08)**: GPT-OSS-120B runs on a **single H100**, GPT-OSS-20B fits in **16 GB** of memory; credited with cutting OpenAI's inference costs by **~75%** for these models.
- **AMD datapoint**: Dell measured **up to 6.1× inference throughput** with MXFP4 on a PowerEdge XE9785L with **AMD Instinct MI355X** GPUs.
- `llm-compressor` supports MXFP4 with per-group quantization (group_size=32), fully dynamic activation quantization, and **no calibration data required** (round-to-nearest) — usable as an NVFP4 alternative when calibration data is unavailable.
- **MXFP4 = Hopper+ only**: gpt-oss-120b (MXFP4) runs on H100, **not on A100** (no Ampere MXFP4 kernel); NVFP4 W4A16 via Marlin does reach Ampere/Hopper (float4_e2m1f supported even at sm_86) — only W4A4 is Blackwell-only.
- **NVFP4 ≠ MXFP4**: block size 16 vs 32; per-block scale FP8 E4M3 + per-tensor FP32 scale vs E8M0; NVIDIA-proprietary vs open spec (partly supported by AMD MI350 / Intel Gaudi 3). The industry frequently conflates the two "FP4"s — they differ in precision and hardware paths.
- Qualification carried: FP4 is the production default **on Blackwell**; on Ampere/Hopper it rides the Marlin W4A16 fallback; FP4 KV cache remains experimental outside datacenter Blackwell.

### Sub-4-bit: BitNet, Bonsai, and the ternary frontier

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

- `llm-compressor` (vLLM project) is the one-command pipeline for FP8/AWQ/GPTQ/MXFP4/NVFP4: FP8 needs no calibration; MoE-aware `--moe` keeps router/gate layers unquantized for stability; one command emits a Hugging Face repo with a generated model card (compressed-tensors safetensors).
- NVFP4 safetensors checkpoints carry `hf_quant_config.json`; vLLM falls back automatically to the Marlin path on pre-Blackwell (the `VLLM_NVFP4_GEMM_BACKEND=marlin` override is deprecated in favor of `--linear-backend`).
- ModelOpt 0.46.0 added **NVFP4 Four-Over-Six (4/6)** and MXFP8 checkpoint support in vLLM; ModelOpt quantization schemes cover INT8, FP8, INT4, NF4, NVFP4 (weight-only, channel, block).
- **Pruning and quantization compose multiplicatively**: Minitron structured pruning + FP8 PTQ on Nemotron-3-Nano-30B = 2.6× throughput and 2.6× memory; **REAP (50% expert pruning) + AutoRound W4A16 took GLM-4.7 from 700 GB to 92 GB (~6.5× total compression)**.
- Pre-quantized NVFP4 examples: DeepSeek-R1-0528, Llama 3, FLUX.1-dev, Qwen3.6-27B/35B-A3B, `ornith-ai/Ornith-1.5-35B-A3B-NVFP4`.
- **DeepSeek-V4-Flash-NVFP4-FP8-MTP hardware scope**: B200/B300 native; on consumer Blackwell vLLM falls back to the Marlin BF16 kernel — use the W4A16 sibling (~159 GB) there instead of the 172 GB NVFP4 build.
- **W4A16 vs W4A4 — check the model card before downloading**: NVFP4 W4A16 runs on Ampere/Hopper via Marlin (`float4_e2m1f` supported even at sm_86) — the memory win without native FP4 compute; **W4A4 needs Blackwell** (native FP4 activations).
- HQQ third-party benchmark figures (~2× speedup on H100/A100 at 4-bit, ~11–13% perplexity increase on Llama-2) come from a community research compilation, not peer-reviewed publication — treat as directional [DIRECTIONAL].

