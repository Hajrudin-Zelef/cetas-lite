---
id: collect-261001-ia-llm/ia-llm/etape4-trackb-local-inference-3
title: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Huawei", "Intel", "MiniMax", "Z.ai"]
dates: ["2026-03-18", "2026-08-18", "2026-08-27", "2026-09-07"]
keywords: ["inference", "llama", "llama.cpp", "apache", "ascend", "attention", "decode", "deepseek", "diffusion", "distribution", "flash attention", "glm"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [114, 149]
sha256: 4d38469b0e7a3cf413fb31104800ca9bcbaef58c230010611c578e050235ae2e
---

# Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)

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
## 4. Performance: flash attention, speculative decoding, measured pp/tg

### 4.1 Flash attention & core kernels (2026)
- **Sparse flash attention** landed in mainline: CUDA + ggml sparse-FA for DeepSeek-V4/GLM (#27970), Metal sparse FA (#28098), `ggml_flash_attn_ext_set_n_kv_max` scheduling API (ggml v0.23.0) [official] (v0.4.0).
- Dense-FA tuning: CUDA XOR-swizzle of K/V smem fp16 tiles; per-device (Q, NE) fa-vec tunings across the whole Apple line; MoE-weighted expert reduction fusion extended to speculative decoding; Metal dequantized quantized KV→F16 before FA for large batches; Vulkan dequant-Q8_0-KV-once in coopmat1 [official] (v0.2.0/v0.4.0).
- KV-cache machinery: n-gram history lookup in the sequence position index; early sequence-scan termination; non-contiguous restore optimization; per-layer expert routing; `--n-cpu-ffn`; `--fit` auto layer distribution (moved from server to common, stream-aware) [official] (v0.3.0/v0.4.0).
- **DFlash2** (block-diffusion speculative decoding, inco.ai): 2B Apache-2.0 draft released **2026-08-18**; llama.cpp engine support merged **2026-08-27** via PR #27342 (Jian Chen, per community doc); converter fix for NVFP4 scales (#28000); fused DFlash encoder into KV-cache injection [official+secondary] (v0.4.0; https://runaihome.com/blog/qwen38-27b-dflash2-speculative-decoding-guide-2026/; https://github.com/nathanw1014/strix-halo-llamacpp/blob/HEAD/docs/dflash2-strix.md).

### 4.2 Speculative decoding modes in 2026
| Mode | Status |
|---|---|
| Draft-model (EAGLE-style) | Long-standing `--spec-type` support |
| **draft-mtp** (multi-token prediction heads) | First-class; native flag `--spec-type draft-mtp`; upstream added MTP for GLM-4.5-Air (v0.3.0); ik supports MTP for GLM-4.x, Qwen 3.5/3.6, Gemma 4, GLM 5, Step 3.7, Qwen 3.8 [official/secondary] |
| **DFlash/DFlash2** | Block-diffusion: drafter predicts a whole block (block_size=16) in one pass from extracted target hidden states; DFlash2 adds conv+selector for higher acceptance; replaces MTP [official/secondary] |
| **DSpark** | DFlash + semi-autoregressive Markov head + confidence head for DeepSeek-V4-class models; upstream supports DSpark for Nemotron 3.5, bailingmoe3, LFM2 [official] (v0.2.0/v0.3.0/v0.4.0) |
| ik self-spec | n-gram and suffix self-speculative decoding [secondary] (ik README) |

