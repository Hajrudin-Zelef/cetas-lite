---
id: ai-industry-kb-2026/08-quantization-model-formats/main-actors-continued
title: "Main actors (continued)"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "AWS", "Alibaba", "Apple", "DeepSeek", "Falcon", "Google", "Hugging Face", "Intel", "Meta", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "Qualcomm", "SGLang", "TII", "Unsloth", "vLLM"]
dates: ["2023-09", "2026-02-13", "2026-02-20", "2026-05", "2026-05-11", "2026-05-13", "2026-05-28", "2026-07", "2026-07-15", "2026-08-15", "2026-08-18", "2026-08-27", "2026-09-22"]
keywords: ["amd", "apache", "attention", "attribution", "awq", "aws", "benchmarks", "bitnet", "blackwell", "compute", "consumer", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4420, 4461]
section: "8. Quantization & Model Formats"
sha256: 96d2533582a7f5e3d0577a4981016961f3268a52cd7fc78f84bb6de6fdade2fb
---

# Main actors (continued)

- Qwen-Coder 32B FP8: RTX 6000 Ada ~100–120 tok/s (1.7–2× vs FP16); RTX 4090 ~60–80 tok/s (emulated, no speedup).
- Llama 3.3 70B, A100 80GB, batch=1 (estimated): Q4_K_M ~20–30 tok/s at ~42 GB VRAM; Q5_K_M ~28–36 tok/s at ~48 GB; Q8_0 ~20–25 tok/s at ~75 GB; FP16 needs 2× A100 (~30–40 tok/s, 140 GB) — Q4_K_M on one A100 costs roughly half per hour of FP16 on two.
- DeepSeek-V3.2 NVFP4+TP2 (GB300): **7,360 TGS prefill-only**, 2,816 TGS mixed (ISL=2k/OSL=1k); DeepSeek-R1 NVFP4+EP2 (2× GB300): **22,476 TGS prefill-only**, 3,072 TGS mixed — 8× prefill, 10–20× mixed-context vs Hopper.
- Nemotron-3.5-Lightning-30B-A3B-NVFP4 on vLLM 0.27.1 (2026-08-18): L40S 1345.4 tok/s @ $0.182/1M out tokens; **RTX PRO 6000 Blackwell 2541.1 tok/s** @ $0.239/1M; A100 2100.8 tok/s @ $0.182/1M [COMMUNITY].
- RTX PRO 6000 shootout (Qwen3.8-Flash-Next, vLLM 0.17.0rc1, TP4, Sept 2026): AWQ 3519 vs NVFP4 3232/3220 tok/s at C=128 with MTP; 2796 vs 2294/2291 without MTP — AWQ won decode at every concurrency on this rig [COMMUNITY].
- Qwen3.6-35B-A3B NVFP4 GEMM on consumer Blackwell (s0me1-dev SM120 patches): 175 tok/s [COMMUNITY].
- 2× RTX 5090, Qwen3-8B NVFP4, concurrency 128–256: **4,400+ TPS at $0.002/MTok**; 380M tokens/day under $800/year electricity [COMMUNITY].
- Private-Blackwell RAG (vLLM + AIPerf): NVFP4 **1.6× throughput vs BF16**, 41% energy reduction, 2–4% quality loss [COMMUNITY/SECONDARY].
- INT4 KV + fused flash-attention kernel: **3.9× kernel speedup** to the bandwidth roofline; attention MAE 3e-08 vs FP32; 42/42 GPU tests passing [COMMUNITY].
- NVIDIA official figures: NVFP4 ≈ 2× FP8 GEMM compute; B200 FP4 ≈ 8× FP16 tensor-core throughput [VENDOR]; FP8 30–50% speedup over FP16 on H100.
- Hardware sizing rules: **RTX PRO 6000 (Blackwell, 96 GB)** holds 70B+ models in FP4 on a single card; **RTX 5090 (32 GB)** is the budget entry point for 14B–32B FP4 inference; Ornith-1.5-35B-A3B-NVFP4 loads on an 8 GB-class Ada GPU via the Marlin fallback (~4.5 bits/value, W4A16-class speed) vs 21.7 GB for the GGUF Q4_K_M of the same model.

## Main actors (continued)

- **NVIDIA** — NVFP4 format owner; tcgen05 FP4 tensor cores on Blackwell (SM100/103); claimed ~2× FP8 GEMM compute and B200 ≈ 8× FP16 [VENDOR]; ModelOpt (0.46.0, 2026-08-18: Triton fast path 34×, LSQ/Dual-LSQ, NVFP4 4/6, MXFP8); DeepSeek-R1-FP4 (671B) production proof; Kimi-K2.6-NVFP4 posted 2026-05-13; Nemotron-3-Nano-30B NVFP4 via QAD (99.4% of BF16) [VENDOR]; official NVFP4 guidance (3.5× vs FP16, 1.8× vs FP8, <1% loss).
- **OpenAI** — GPT-OSS ships natively in MXFP4 (2025-08): GPT-OSS-120B on a single H100, GPT-OSS-20B in 16 GB; MXFP4 credited with ~75% inference-cost reduction.
- **Open Compute Project (OCP)** — MX specification v1.0 (September 2023): AMD, Arm, Intel, Meta, Microsoft, NVIDIA, Qualcomm; E2M1 blocks of 32 with E8M0 scales; the open standard behind MXFP4.
- **AMD** — Instinct MI355X: Dell measured up to 6.1× inference throughput with MXFP4 on PowerEdge XE9785L; MI350 partial MXFP4 support; AMD Quark on ROCm paths.
- **Intel** — AutoRound: strongest measured 4-bit PTQ (lm-eval 30/32 vs GPTQ, 27/32 vs AWQ; 91.4% top-1 logprob agreement); **OpenVINO 2026.2**: INT4 KV-cache compression on Intel GPU (BY_CHANNEL keys group-16, BY_TOKEN values on paged attention).
- **Microsoft Research** — BitNet: natively trained 1.58-bit ternary models (b1.58 2B4T, MIT); at 3B matches FP16 LLaMA perplexity/zero-shot; the 100B-on-single-CPU claim remains theoretical.
- **Moonshot AI** — Kimi K2.6 (1T total, 32B active): ships INT4 natively via QAT (~2× speedup); NVIDIA's Kimi-K2.6-NVFP4 (2026-05-13) makes it a Blackwell-format reference.
- **DeepSeek** — DeepSeek-R1-FP4 (671B: MMLU −0.1, GSM8K −0.2 vs FP8 base); DeepSeek-V3.2/R1 NVFP4 on GB300 (vLLM blog 2026-02-13; R1: 22,476 TGS prefill-only on 2× GB300); DeepSeek-V4: FP4 (MXFP4) QAT on MoE experts + indexer QK path; community DeepSeek-V4-Flash-NVFP4-FP8-MTP quant (May 2026, 172 GB).
- **Unsloth** — NVFP4 export since July 2026 ("Dynamic Unsloth NVFP4 Quants": gemma-4-31B-it-NVFP4, gemma-4-26B-A4B-it-NVFP4, gemma-4-12b-it-NVFP4; hybrid NVFP4-GGUF containers such as Qwen3.6 NVFP4-GGUF, accelerated only via Unsloth kernels on Blackwell); 1,374-model Hugging Face catalog (GGUF + NVFP4).
- **PrismML** — Bonsai-27B (2026-07-15, Apache 2.0): first 27B model on a phone; 1-bit "Binary g128" (3.9 GB, >90% retention, 11 tok/s iPhone 17 Pro) and ternary (7.2 GB, >95% retention); near-lossless 4-bit KV enabling 262K context on-device.
- **Meta** — Qwen3.8-27B (2026-08-15) as the Dynamic v3.0 launch vehicle; Meta Muse covered by v3.0 quants.
- **Alibaba / Qwen** — Qwen-Coder 32B: the FP8 accuracy reference datapoint (79–80% vs ~80% FP16); Qwen3.6-27B/35B-A3B NVFP4 pre-quantized checkpoints.
- **Hugging Face** — absorbed ggml.ai (announced 2026-02-20; MIT license and team autonomy preserved), cementing GGUF as the de-facto local format; 6,500+ AWQ repos on the Hub; optimum-quanto in maintenance mode (confirmed 2026-09-22, redirects to bitsandbytes/torchAO); quantization overview lists GPTQModel as the GPTQ entry.
- **Red Hat** — independently evaluated TurboQuant and drove its vLLM integration (May 2026; PR #38479): FP8 remains the best KV-cache default; publishes quantized checkpoints (e.g. Qwen3.5-122B-A10B-NVFP4).
- **Google** — TurboQuant: the original method (arXiv:2504.19874, ICLR 2026); attribution reconciled against earlier Red-Hat-as-inventor wording.
- **Community / independent actors** — srmiles (first independent Dynamic v3.0 replication test, 2026-08-27); MarcoPizeta (vLLM/SGLang/EXL3 shootout, Sept 2026); rtx6kpro (AWQ > NVFP4 decode shootout, Sept 2026); avarok (NVFP4 W4A4 MoE on GB10, production-ready); canada-quant (2026-05-28 Marlin-fallback measurements); 0xsero (Blackwell dtype matrices); club-3090 (consumer-Blackwell dtype matrix); massed-compute (Nemotron NVFP4 benchmarks, 2026-08-18); TII (Falcon3-10B-Instruct-1.58bit, Falcon-E 1B/3B ternary); mobiusml (HQQ data-free quantization, gemlite Triton kernels with MXFP4/NVFP4 dynamic).

- **TII** — Falcon3-10B-Instruct-1.58bit (~2 GB GGUF) and Falcon-E 1B/3B (natively trained ternary from scratch on ~1.5T tokens; 665 MB / 999 MB disk).
- **Dell** — measured up to 6.1× inference throughput with MXFP4 on PowerEdge XE9785L + AMD Instinct MI355X [VENDOR].
- **Artefact2** — GGUF quantization overview: the community-standard per-type KL/PPL table measured on real hardware [COMMUNITY].
- **mradermacher** — per-type model cards converging on the IQ-quant ladder ("IQ-quants are often preferable over similar sized non-IQ quants"; DeepSeek-V2.5 236B-class i1-IQ4_XS = 125.7 GB at ~0.8% perplexity increase) [COMMUNITY].
- **Intel Arc Pro B70 cookbook (sergiob)** — independent AutoRound-vs-GPTQ logprob-parity test (Ornith-1.5-35B-A3B, vLLM-XPU): AutoRound 91.4% top-1 agreement / KL 0.277 vs GPTQ 90.5% / 0.319 [COMMUNITY].
- **Red Hat AI engineers (Eldar Kurtic, Michael Goin, Alexandre Marques)** — TurboQuant vLLM evaluation, published 2026-05-11; drove the vLLM integration (PR #38479).
- **0xsero / club-3090 / massed-compute / canada-quant / srmiles / s0me1-dev** — community dtype matrices (NVFP4 KV datacenter-only; consumer-Blackwell W4A16-vs-W4A4 guidance), throughput/$ benchmarks, the 2026-05-28 Marlin-fallback measurement, the first Dynamic v3.0 independent test, and SM120 NVFP4 GEMM patches [COMMUNITY].
- **NextPlatform** — AWS Blackwell rental economics analysis: FP8 halves cost-per-teraflop vs FP16; FP4 halves it again [ANALYST].

## Timeline and context (continued)

