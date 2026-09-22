---
id: etape4-trackb-local-inference/00-local-inference/4-performance-flash-attention-speculative-decoding-measured-
title: "4. Performance: flash attention, speculative decoding, measured pp/tg"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: model-release
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Nvidia", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-08-18", "2026-08-19", "2026-08-27", "2026-09"]
keywords: ["attention", "flash attention", "speculative decoding", "amd", "apache", "bitnet", "consumer", "decode", "deepseek", "diffusion", "distribution", "gguf"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [133, 180]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: c933bd2f9dbe33760cb209a9ec2f26583799b835e156f3f26849790556250123
---

# 4. Performance: flash attention, speculative decoding, measured pp/tg

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

### 4.3 Published pp/tg numbers (as reported — do not compare across setups)
- **NVIDIA gpt-oss-20b on llama.cpp** (NVIDIA testing via llama.cpp): RTX 5090 **282 tok/s**, Mac M3 Ultra **116 tok/s**, AMD 7900 XTX **102 tok/s** [vendor-reported via independent] (https://www.techradar.com/ai-platforms-assistants/gpt-oss-20b-performance-faster-pc-rtx-nvidia).
- **DFlash2, Qwen3.8-27B** (Sept 7, 2026): 3.43× claimed on SGLang/vLLM at C=1; llama.cpp implementation measures **1.81×**; RTX 3090 on vLLM ≈ 2.6×. If MTP already runs, DFlash2 adds only a few percent [secondary] (https://runaihome.com/blog/qwen38-27b-dflash2-speculative-decoding-guide-2026/).
- **DFlash2 on Strix Halo (Vulkan/RADV, 2026-08-19)**: 2.23× at d0, 1.89× at 8k, **2.00× at 32k**; v1 collapses to 1.02× by 32k; decode base 11.81→10.54 t/s [secondary] (https://github.com/nathanw1014/strix-halo-llamacpp/blob/HEAD/docs/dflash2-strix.md).
- **ik_llama.cpp vs upstream** (RTX 3050 6GB hybrid, Qwen3.6 35B MoE APEX I-Compact + MTP): ik **91.2 pp / 40.4 tg** vs upstream 60.8–67.1 pp / 30.5 tg (+41% prompt); dense GPT-OSS 20B Q4_K_M upstream wins (66.8–104.6 pp / 31.8 tg vs ik 42.1–53.9 / 22.4). Rule of thumb: ik for MoE, upstream for dense [secondary] (https://github.com/luksamuk/ai-dotfiles/blob/HEAD/llama-swap/docs/BINARIES.md).
- **Strix Halo (Qwen3-30B-A3B, 130k ctx)**: ROCm pp512 41 t/s tg128 5 t/s; rocWMMA-tuned pp512 51 / tg128 13 t/s; Vulkan RADV pp512 17 / tg128 13 t/s. DGX Spark vs Strix Halo @120B tg32: **56 vs 47–53 t/s** [secondary] (https://github.com/getnyrex/strix-halo-guide/blob/HEAD/RESEARCH.md).
- **Strix Halo (gfx1151, 64GB, Vulkan, Apr 2026)**: kyuz0 RADV build — **351 pp / 19 tg**; official b8299 — **393 pp / 22 tg**; turboquant fork 8793 — 393 pp / 22 tg [secondary] (https://github.com/valentijnvenus/strix-halo-setup).
- **DGX Spark GB10 (llama.cpp, INT4)**: Llama 3.3 70B ≈ 3.8 t/s decode; GPT-OSS 120B (4-bit MoE) ≈ 32–33 t/s; general guidance <20B = 15–40 t/s [secondary] (https://github.com/dylancouzon/aie-talk/blob/HEAD/research/V-nvidia-arm.md).
- **Dual Xeon (2× QYFS/8480+, ik_llama.cpp, DeepSeek-R1 Q4_K_M)**: pp512 109.69 t/s, pp1000 105.46 t/s, tg128 8.74 t/s [secondary] (https://forums.servethehome.com/index.php?threads/es-xeon-discussion.5031/page-205).
- **Apple Silicon study (M2 Ultra, early 2026 arxiv 2511.05502)**: llama.cpp ≈150 t/s short-context only; MLX ≈230, MLC-LLM ≈190, Ollama 20–40 [independent] (https://arxiv.org/pdf/2511.05502v1.pdf).
- **llama.cpp vs Ollama**: 3–12% raw throughput advantage claimed for llama.cpp (video analysis, 2026-08) [secondary] (https://www.youtube.com/watch?v=MAkd1qmOuAI); vs vLLM on single 24 GB card: llama.cpp **120 t/s vs vLLM 19 t/s** (6–7×, CUDA-graphs memory argument) [secondary] (https://medium.com/data-science-collective/what-is-the-best-hardware-for-running-local-llms-in-2026-mac-vs-5090-vs-cloud-ff023b660442).
- Qwen3.8-Flash-Next ≈125B-A6B MoE + 51B "engram" table on Strix Halo 128GB (EngramHalo fork, UD-IQ4_XS): ≈28 t/s warm [secondary] (https://github.com/sypherin/strix-halo-setup/blob/HEAD/docs/qwen3.8-flash-next-engramhalo-stability.md).

---
## 5. ik_llama.cpp fork — status & headline features

**Repo:** https://github.com/ikawrakow/ik_llama.cpp (Iwan Kawrakow). Active in September 2026 (README references PRs up to ~#2369; changelog entries from 2026-03 merged `--fit` etc.) [secondary].

**What it is:** performance-focused fork of llama.cpp with better CPU and hybrid GPU/CPU performance, new SOTA quantization types, first-class BitNet support, and MLA/FlashMLA/fused-MoE improvements [secondary] (https://github.com/ikawrakow/ik_llama.cpp/blob/HEAD/README.md).

**Headline 2025→2026 features (from its own changelog):**
- **IQK quants**: IQ2_KS, IQ3_KT, IQ4_KT, IQ4_KS, IQ5_KS, IQ1_S — kernels co-designed with the grids; AVX-VNNI optimizations; Zen4 PP speedups for IQ2_KS/IQ4_KS/IQ5_KS (#428); faster IQ3_KT/IQ4_KT (#453); ~2% CUDA tg gain for iq2_ks (#468).
- **New split mode "graph"** for multi-GPU (PR #1022) — tensor parallelism at the GGML graph level; reported **3–4× speedup on 4× Tesla T4** vs layer/row splitting [secondary] (https://medium.com/@jagusztinl/llama-cpp-performance-breakthrough-for-multi-gpu-setups-04c83a66feb2).
- Fused delta-net for Qwen3-Next and Qwen3.5-MoE (#1315/#1333/#1362/#1373); Hadamard transforms for K/V cache (#1033/#1034/#1527); auto-fit offloaded tensors to VRAM for MoE+dense (#1501/#1504, merged Mar 2026, per ik fork docs), per-GPU fit margin (#1872).
- **MTP decoding**: GLM-4.x MoE (#1270), Qwen 3.5/3.6 (#1698/#1745), Gemma 4 (#1744), GLM 5 (#1890), Step 3.7 (#2250), Qwen 3.8 (#2369), incl. standalone MTP heads via `-md`.
- Self-speculative decoding (n-gram #1261, suffix #1646); DFlash initial support (#1970); DSpark initial support (#2280).
- Server extras: **OpenAI `/v1/responses` endpoint** (#1184), function-call support (#628), jinja template support (#677), expiring logit bias, string-ban, Adaptive-P sampler, multimodal vision in llama-mtmd-cli and llama-server, mikupad alt WebUI, MCP support (#1904), dynamic control-vector endpoints, on-demand tensor reload.
- Checkpoints for recurrent models (#1310/#1398); GLM-DSA indexer cache; GLM-5.2 vision hack (#2283).

**Why it matters:** the go-to engine for **CPU and hybrid CPU/GPU MoE inference** (DeepSeek-class models on consumer hardware); origin of IQK quants widely considered the best GGUF quality-per-bit; source of several ideas later re-adopted upstream (auto-fit/`--fit` logic). **Scope limits (explicit):** only CPU (AVX2+/NEON+) and CUDA (Turing+) are fully supported backends — ROCm/Vulkan/Metal issues are not handled [secondary] (ik README). Only use `-rtr` (row-interleaved repack) deliberately for k-quants, which lack CUDA row-interleaved implementations [secondary] (ik README).

