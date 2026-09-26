---
id: ai-industry-kb-2026/08-quantization-model-formats/sub-4-bit-bitnet-bonsai-and-the-ternary-frontier
title: "Sub-4-bit: BitNet, Bonsai, and the ternary frontier"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "Alibaba", "DeepSeek", "Intel", "Meta", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "Qualcomm", "TensorRT-LLM", "Unsloth", "vLLM"]
dates: ["2023-09", "2026-02-13", "2026-05", "2026-05-13", "2026-07", "2026-08-18"]
keywords: ["bitnet", "amd", "attention", "benchmarks", "blackwell", "compute", "consumer", "datacenter", "decode", "deepseek", "distillation", "embeddings"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4244, 4271]
section: "8. Quantization & Model Formats"
sha256: 33951332236818ed64835adf3bc6fc2a53b19427369ccc0aaacf5b557f846b3e
---

# Sub-4-bit: BitNet, Bonsai, and the ternary frontier

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

