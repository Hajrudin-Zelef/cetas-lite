---
id: etape4-trackb-local-inference/00-local-inference/earlier-2026-febaug-from-b-tags-selected-milestones-secondar
title: "Earlier 2026 (Feb–Aug, from `b` tags) — selected milestones [secondary]/[official]"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["Alibaba", "Unsloth"]
dates: ["2026-02-13", "2026-08-31", "2026-09", "2026-09-07", "2026-09-14"]
keywords: ["attention", "benchmark", "benchmarks", "bitnet", "gguf", "gpu", "llama", "llama.cpp", "moe", "nvfp4", "quantization", "research"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [75, 113]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 5188983e2edb72a7e8247319f5e8db7daa57552d213c196e35694b559c6f8979
---

# Earlier 2026 (Feb–Aug, from `b` tags) — selected milestones [secondary]/[official]

### Earlier 2026 (Feb–Aug, from `b` tags) — selected milestones [secondary]/[official]
- **b7964** (Feb 2026): Step3.5-Flash support (#19283). **b7973**: Qwen3.5 dense and MoE support (community-driven, PR #19435; author notes "Llama.cpp missing all the zero-day releases", built on the common-delta-net PR #19125) [secondary] (https://pypi.org/project/llama-cpp-pydist/0.31.0/).
- **b8018** (2026-02-13): pydist sync point with 44 upstream commits [secondary] (same).
- Community build records (Apr 2026, Strix Halo): b8119 (kyuz0 custom), b8299 official (+40% prompt speed over b8119), b8461 (kyuz0 Vulkan RADV) [secondary] (https://github.com/valentijnvenus/strix-halo-setup).
- llama-cpp-python fork (JamePeng) synced to upstream commit `9723942` on 2026-08-31 with `load_mode` + Qwen3.8-Flash-Next + DFlash2 NVFP4 fix [secondary] (https://github.com/tao71-ai/llama-cpp-python-jamepeng/blob/HEAD/CHANGELOG.md).

---
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

