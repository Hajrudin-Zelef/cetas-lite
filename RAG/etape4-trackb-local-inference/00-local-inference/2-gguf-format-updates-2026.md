---
id: etape4-trackb-local-inference/00-local-inference/2-gguf-format-updates-2026
title: "2. GGUF format updates (2026)"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: quantization
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Huawei", "Intel", "MiniMax", "Unsloth", "Z.ai"]
dates: ["2026-03-18", "2026-09", "2026-09-07", "2026-09-14"]
keywords: ["gguf", "ascend", "attention", "benchmark", "benchmarks", "bitnet", "decode", "deepseek", "glm", "gpu", "gpus", "inference"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [82, 132]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 74bf948ff59b50956b7e2d9f42ead23a7676f3ed2c571911ec6f1748cf476398
---

# 2. GGUF format updates (2026)

## 2. GGUF format updates (2026)

### 2.1 Mainline llama.cpp: NO new weight-quantization types in 2026
The authoritative evidence is the current-master server README (September 2026): allowed KV-cache types are still exactly `f32, f16, bf16, q8_0, q4_0, q4_1, iq4_nl, q5_0, q5_1` [official] (https://github.com/ggml-org/llama.cpp/blob/master/tools/server/README.md). The weight-quant family remains the established three tiers:

| Family | Examples | Calibrated? | Quality ranking (per a 2026 community taxonomy) |
|---|---|---|---|
| Legacy | Q4_0, Q4_1, Q5_0, Q5_1, Q8_0 | No | Worst-per-bit, but Q8_0 near-lossless [secondary] |
| K-quants | Q2_K … Q6_K (Q4_K_M still the mainstream default) | No (data-free) | Middle [secondary] |
| I-quants | IQ1_S … IQ4_XS/IQ4_NL (non-linear lattice codebooks + importance matrix) | Yes (imatrix) | Better per-bit, especially < 4 bpw [secondary] |

(https://github.com/noonghunna/club-3090/blob/HEAD/docs/QUANTIZATION.md)

The only 2026 "format-adjacent" additions in mainline are:
- **NVFP4 scales** in the DFlash2 speculative-decoding path (dflash: pass missing NVFP4 scales to attention ops, #28000, v0.4.0 window) [official] (v0.4.0 changelog).
- **TQ1_0/TQ2_0** (ternary BitNet quants) exist in mainline but pre-date 2026 (v0.2.0 changelog marks `tq2_0` as not supported on SYCL — i.e., it exists, just unsupported there) [official] (v0.2.0 changelog).
- **DFlash2/NVFP4 draft metadata**: new GGUF metadata keys for the DFlash2 drafter (`conv_kernel_size`, `conv_group_size`, `selector_rank`, `selector_top_k`) [secondary] (https://github.com/wszhoho/llama-cpp-turboquant-dflash2 — describes the format; upstream merged DFlash2 via #27342).
- Vocab handling: **integer tokenizer scores** supported (v0.1.2, #27260); DFlash draft conversion reuses the *target's* vocab via `--target-model-dir` and adds `_vocab_hparams()` so special-token ids follow the target config (fixes HunYuan pad/eod ids) [official] (releases page, PR #28890 description).

### 2.2 New-in-2026 quantization types are FORK-EXCLUSIVE (verified)
- **IQK quants** (`IQ4_KS`, `IQ5_KS`, `IQ4_K`, `IQ2_K`, `IQ2_KS`, `IQ3_KT`, `IQ4_KT`, `IQ1_S`) — ik_llama.cpp only. Refined grids + imatrix + co-designed CUDA/CPU kernels; widely regarded as the best quality-per-bit in the GGUF world [secondary] (https://github.com/noonghunna/club-3090/blob/HEAD/docs/QUANTIZATION.md; https://github.com/ikawrakow/ik_llama.cpp). Note: ik's README warns **not** to use Unsloth `_XL` models containing f16 tensors, and that K-quants lack CUDA row-interleaved implementations so `-rtr` hurts hybrid GPU/CPU for them [secondary] (ik README).
- **TurboQuant** (`TURBO2_0`=enum 43, `TURBO3_0`=44, `TURBO4_0`=45; WHT rotation + PolarQuant, based on arXiv 2504.19874/ICLR 2026; fused into FA kernels as turbo KV types; InnerQ channel equalization) — **wszhoho/llama-cpp-turboquant-dflash2** fork only [secondary] (https://github.com/wszhoho/llama-cpp-turboquant-dflash2). (Strix Halo lab notes confirm turbo KV types were CPU-only upstream-incompatible and the setup later retired the fork for upstream Vulkan [secondary] https://github.com/bogdan-d/strix-halo-setup.)
- **ROCmFPX** (GGML types 100–107, "GGML_TYPE_COUNT=108" in that build) — **artomyuan/llama.cpp-rocm** fork only; fixes "invalid type 101" validation in `llama-quantize` (2026-09-14 changelog) [secondary] (https://github.com/artomyuan/llama.cpp-rocm/blob/HEAD/CHANGELOG.en.md).
- **APEX I-Compact** quants appear in community benchmarks (Qwen3.6 35B MoE "APEX I-Compact") but no verified upstream/FHF format spec was located — treat as [unverified] in the GGUF-format sense [secondary] (https://github.com/luksamuk/ai-dotfiles/blob/HEAD/llama-swap/docs/BINARIES.md).

### 2.3 Quality-per-bit evidence (2026)
- Unified evaluation on Llama-3.1-8B-Instruct (arXiv 2601.14277, 2026): PPL increase vs FP16 — Q8_0 +0.1%, Q6_K +0.4%, Q5_K_M +1.1%, Q4_K_M +3.3%, Q4_K_S +4.1%, Q3_K_M +8.7%, Q3_K_S +22.4% [independent] (https://arxiv.org/pdf/2601.14277v1). Note: 2026-trained models lose measurably more per quant step than 2023-era models [secondary] (https://github.com/pradeepgudipati/gguf-switchboard/blob/HEAD/docs/QUANT_SCORING.md).
- Quesma benchmark (Aug 2026) of Qwen3.8-27B: **Q4_K_M matches BF16** on GPQA Diamond, IFBench, Terminal-Bench 2.1; Q2 drops; IQ1 collapses to chance [secondary] (https://github.com/chriscorbell/llm-server/blob/HEAD/docs/research/2026-09-07-b70-and-qwen38-27b.md).

---
## 3. Backends: status in 2026

| Backend | Status / 2026 highlights |
|---|---|
| **CUDA** | First-class. XOR-swizzle flash-attn K/V smem fp16 tiles (#25635); tuned MMVQ→MMQ decode crossover per HW/quant (#26079, #28285 for SM87); fused MoE weighted expert reduction extended to specdec + topk-router fusion (#27621, #25952); fast `mm_ids_helper` for any `n_expert_used` (#27978); **concurrent streams per split for multi-GPU** (#28198); CUDA 12 (12.8 libs) and CUDA 13 (13.4 libs) official builds; **Windows arm64 CUDA 13.4 (preview)** added [official] (release pages + v0.4.0 changelog). |
| **Metal** | First-class. Massive 2026 per-device **fa-vec tunings**: M1, M2, M2 Pro, M2 Max, M1 Ultra, M3, M3 Pro, M3 Max, M3 Ultra, **A18 Pro (MacBook Neo)**, M5+ — plus **Metal 4.0 tensor API enabled on M5+/A19+** (#27461); sparse FA (#28098); per-op source split + parallel compilation (v0.3.0/ggml 0.22); top-k radix; quantized-KV dequant to F16 before FA for large batches; bin-kernel consolidation (#19390, Feb 2026) [official] (v0.4.0, v0.2.0 changelogs; pydist changelog). |
| **Vulkan** | Strong community/mainline backend. IQ3_S mat-vec for batch>4 (#27449); bfloat16 extension only if supported (#28155); Q8_0 KV dequant in coopmat1 (#25494); tiled transpose; **Intel Xe FA optimization kernels (Xe-LPG Plus / Xe2 / Xe3, #24406)**; mat-vec row tuning for **Strix Halo** batched inference (#27909); top_k radix for k≥1024 (Qwen 3.8 Flash Next, #28032); Vulkan FA MMQ fp32 Q-quant [official] (releases page, v0.2.0/v0.4.0). Reported ~4–5% uplift from Mesa 26.0.1→26.0.2 on Strix Halo [secondary] (https://github.com/getnyrex/strix-halo-guide/blob/HEAD/RESEARCH.md). |
| **ROCm/HIP** | Official builds at **ROCm 10.0** (updated in v0.4.0 window, #27803); radix TOP_K for long rows (#27466); RDNA3 MMQ config tuning; Q2_0 dot-product path for gfx1201 (#26753). Known pain: ROCm 7.2 regression — 73% perf loss with `GGML_HIP_ROCWMMA_FATTN=ON` on upstream (Feb 28, 2026); Qwen3.5 hang on gfx1151 during load_tensors (open, 2026-03-18; workaround `--flash-attn off --n-gpu-layers 1`); kernel 6.19.4 misdetects gfx1151 as gfx1100 [secondary] (https://github.com/getnyrex/strix-halo-guide/blob/HEAD/RESEARCH.md). |
| **SYCL** | Intel GPUs. Fused rms_norm+mul+add / add+add chains (#27610); Q2_K reordered MMVQ + ESIMD kernels; Q5_K ESIMD; peer-to-peer copy API (#27550); MKL FA refactor; host-pinned 2 GiB cap (#27559); FP32/FP16 official builds. `tq2_0` marked unsupported [official] (v0.2.0/v0.3.0/v0.4.0 changelogs). |
| **OpenCL** | Adreno-focused (+Intel Xe-LP tuning). Quant lm_head / decode GEMV and medium-batch GEMM for spec-decoding/MTP (#26477); MoE expert-bias folding into epilogue for gpt-oss (#26431); fused SSM scan (Mamba-2) (#26439); deterministic MoE expert scatter (#26464); Xe-LP TG/PP tuning (#26438); Adreno A6x/A7x compiler workarounds [official] (v0.2.0/v0.3.0/v0.4.0). |
| **OpenVINO** | Still **[In Progress]** in the supported-backend matrix; official Windows/Linux builds ship with OpenVINO 2026.3/2026.3.1; CI had failing OpenVINO tests disabled in the v0.4.0 window (#28347) [official] (README backend table via selfhost.directory; v0.4.0 changelog). |
| **Hexagon (Snapdragon)** | New mobile-tier backend with dedicated release builds (Linux arm64 + Android arm64: CPU/Adreno/Hexagon NPU) and setup guides. 2026 work: MUL_MAT/MUL_MAT_ID fusion, F16 unary ops, device discovery + on-demand sessions, FA HMX queue fixes, CPY fence fix; added to `docs/backend/ops.md` (#28263) [official] (releases page, v0.4.0/v0.2.0). |
| **WebGPU** | All-targets; misc fixes (ARGSORT/TOP_K inf, tensor_get offset, rope offset support, overlapping-src mulmat) [official] (v0.2.0/v0.3.0/v0.4.0). |
| **CPU (x86/ARM/RISC-V)** | AVX/AVX2/AVX512/AMX; ARM NEON + KleidiAI (SME2 F32 GEMV added; KleidiAI macOS build listed DISABLED in latest assets — status unclear, flag as [unverified]); RVV/ZVFH/ZICBOP/ZIHINTPAUSE; **AVX2: large-batch PP speedup for IQ models** (#27402); SpacemiT IME kernels (#27961); s390x fixes [official] (v0.2.0/v0.4.0). Zen4 PP speedups for IQ2_KS/IQ4_KS/IQ5_KS are ik-fork-only (#428) [secondary] (ik README). |
| **RPC / distributed** | **Apple RDMA** added as RPC transport (#26421); RPC event/async backend APIs (#18626); buffer serialization fixes (#26500); `use_count` fusion (v0.2.0) [official] (v0.4.0/v0.2.0). Note: Pi-cluster RPC measurements show ~25× regression vs single-node on 1 GbE [independent] (https://github.com/hellomatik-org/distributed-llama/raw/cfbdcd4c16659a7f61edc0413b05e09b37f74060/paper/main_en.pdf). |
| Others listed | BLAS, BLIS, **CANN** (Ascend NPU), **MUSA** (Moore Threads GPU), IBM zDNN, ZenDNN, VirtGPU [official] (backend table). |

### 3.1 Model-architecture support added in 2026 (notable)
Qwen3.5 dense+MoE (b7973, Feb 2026) → Qwen3-Next/Qwen3.5-MoE fused delta-net (ik, Mar 2026) → Qwen3.6-27B → **Qwen3.8-Flash-Next** (`qwen4exp`, v0.4.0, upstream) → DeepSeek-V4 (sparse FA, DSpark, MLA), GLM-4.5-Air MTP, GLM-5/5.2, Nemotron-3-Puzzle / Nemotron 3.5 DSpark, Gemma 4 (+vision fixes), **Step3.5-Flash** (b7964), **Step 3.7**, bailingmoe3 (DSpark), LFM2/LFM2-MoE (+DSpark), dots3-note (DSA-ISWA KV), MiniMax-01, HunyuanOCR 1.5/v1.0 (+DFlash support #28890), GraniteSWA/GraniteMoeSWA, nanbeige4.2-3B, Qwen3-TTS-0.6B fix, DeepSeek-V4-Flash-Vision-Exp [official] (release pages v0.1.2–v0.4.0; pydist changelog; ik README). **Known issue:** Qwen3.8-27B emits silent EOS past ~130K prompt tokens on CUDA/CPU — open issue #27756 (Aug 26, 2026), suspected GDN recurrent-state error [secondary] (https://github.com/chriscorbell/llm-server/blob/HEAD/docs/research/2026-09-07-b70-and-qwen38-27b.md).

---
