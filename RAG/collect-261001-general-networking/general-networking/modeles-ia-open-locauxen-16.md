---
id: collect-261001-general-networking/general-networking/modeles-ia-open-locauxen-16
title: "ÉTAPE 1 — Open / Local AI Models (EN)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Mistral", "Nvidia", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM"]
dates: []
keywords: ["agents", "amd", "attention", "awq", "blackwell", "consumer", "datacenter", "decode", "deepseek", "fp4", "fp8", "gguf"]
source: docs/RAG/collect-261001-general-networking/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [869, 940]
sha256: 83dcab361c208eb5e85a38ec236fd46d42d61ef0d0b0ddf38389813704e4198c
---

# ÉTAPE 1 — Open / Local AI Models (EN)

### Prosumer single-GPU (32–48 GB — RTX 5090 32 GB, RTX 4090 Pro 48 GB)
- RTX 5090 (32 GB): 70B Q4 is tight; 27B-class models (e.g., Qwen3.8-27B) run comfortably; community verdict — "on today's best models a 5090 nearly matches the Pro 6000 for a fraction of the price" for dense models up to ~32B.
- RTX 4090 Pro (48 GB): 70B Q4_K_M (39 GB) fits with headroom; 70B Q5_K_M (48 GB) tight.

### Workstation flagship (96 GB — RTX PRO 6000 Blackwell, ~$10,000–11,000)
- **Llama 3.3 70B at Q8 (~74 GB):** fits with context headroom; community reports **35–50 tok/s** sustained.
- **70B at BF16:** full-quality, no quantization needed.
- 120B-class at FP8, 180B+ at INT4/FP4 on a single card.
- Qwen3.5-122B-A10B MoE (10B active) fits 96 GB.
- 2× cards (192 GB) reach larger frontier MoEs; 4× (384 GB) covers virtually every open-weight model released to date.
- Specs: 24,064 CUDA cores, 96 GB GDDR7 ECC, 1.79 TB/s bandwidth, 600W.

### Unified-memory desktops (DGX Spark / Strix Halo / Mac)
- **NVIDIA DGX Spark (GB10, 128 GB unified LPDDR5x, ~$4,699):** runs 128B-class models (e.g., Mistral Medium 128B) locally for far less than a Pro 6000 build.
- **AMD Strix Halo (128 GB, ~$2,799 desktop):** same class — 128B local inference at consumer-ish pricing.
- **Apple Silicon:**
  - M4 Max 128 GB: 70B Q4_K_M at ~10–15 tok/s; unified pool shared CPU/GPU, no PCIe copy.
  - M5 Max 128 GB (up to 614 GB/s): 70B Q4 comfortable; Llama 3.3 70B Q6 (~55 GB) leaves room for context + OS.
  - M5 Ultra up to **512 GB** (top config, ships late Oct 2026, expected >$10,000): first single Mac that could theoretically hold GLM-5.3 (744B MoE, ~239 GB at 2-bit) in unified memory — unbenchmarked as of Sept 2026.
  - Mac Studio M5 Max from $2,499 (128 GB); M5 Ultra from $5,499 (96 GB base).
  - **MLX is the 2026 unlock on Mac:** Ollama 0.19+ / LM Studio mlx-engine give 2–3× decode vs llama.cpp-Metal; Qwen3.6-35B-A3B MoE at 60–70+ tok/s on M5 Max-class hardware.

### Multi-GPU / node-scale (the 400B+ MoE tier)
Critical MoE fact: **only ~8–42B params are active per token, but ALL weights must be resident** — the router can select any expert, so nothing can be left off-card.

| Model | Checkpoint | VRAM floor | Practical config |
|---|---|---|---|
| **DeepSeek V4.1 Flash** (552B MoE; community analysis argues 748B incl. 196B Engram tables — disputed) | ~510 GB (FP8 weights + FP4 routed experts, 48 shards) | **~614 GB** (vLLM recipe) | 8× H200 (1,128 GB) recommended; GB200 NVL4 tray validated; 8× H100 80 GB (640 GB) fits with ~130 GB headroom — tight for production |
| **DeepSeek V4 Flash** (older, 284B) | 158 GB native (FP4+FP8) | ~170–175 GB | 2× RTX PRO 6000 (192 GB) min for native quality; Q4_K_M (~86–96 GB) fits a single 96 GB card — tight, 45–60 tok/s, ~5% quality degradation vs native |
| **GLM-5.3 Flash** (320B / 18B active) | ~306 GiB | 8-GPU node class | more headroom on 8× H100 than V4.1 Flash |
| **GLM-5.1** (744B / 40B active) | FP8 ≈ **800 GB** | — | 8× H200/H20 node |
| **GLM-5.3** (744B MoE) | ~239 GB at 2-bit | — | 512 GB Mac Studio M5 Ultra (theoretical); 4× RTX 3090/4090 documented community path |

- DeepSeek V4.1 Flash KV cache: ~890 bytes/token (FP4 cache) — 1M context ≈ 890 MB... note: at 1M tokens the cache is material (~0.9 GB) but was never the binding constraint; weights are.
- No GGUF/Ollama desktop conversion existed for V4.1 Flash at research time (checkpoint already ships FP8/FP4; little room left for further quantization).
- Engines: vLLM ≥ 0.30, SGLang day-zero build, NVIDIA Dynamo — Hopper/Blackwell and AMD MI350X/MI355X.

Sources: https://www.yottalabs.ai/post/deepseek-v4-1-flash-hardware-requirements-gpu-memory-2026 · https://github.com/dbirks/home-k8s/blob/HEAD/deepseek-v4-flash-hosting-notes.md · https://www.orcarouter.ai/blog/deepseek-v4-1-flash-local-deployment · https://www.orcarouter.ai/blog/deepseek-v4-1-flash-inference-stack · https://www.yottalabs.ai/post/deepseek-v4-1-flash-pricing-specs-v4-pro-routing-2026 · https://the-decoder.com/new-deepseek-model-v4-1-flash-cuts-memory-needs-for-ai-agents/ · https://www.promptquorum.com/local-llms/local-llm-hardware-guide-2026 · https://awesomeagents.ai/news/apple-m5-pro-max-70b-models-portable/ · https://www.newegg.com/insider/nvidia-rtx-pro-6000-blackwell-workstation-96gb-gddr7-for-serious-local-ai/ · https://mwgamers.com/blog/nvidia-rtx-pro-6000-monolith-future-gaming/

---

# 3. QUANTIZATION LANDSCAPE 2026

## 3.1 Format cheat sheet (practical bytes/param and quality story)

| Format | Practical bytes/param | Quality story | Ecosystem |
|---|---|---|---|
| BF16 | 2.0 | training + high-quality baseline | everything |
| FP8 (E4M3) | 1.0 | often ~identical to BF16 (~0.0 ppl drop) | server serving default (vLLM/SGLang) |
| **NVFP4** | ~0.56 | close to FP8 in NVIDIA examples; ~0.3–0.5 ppl drop; competitive with best INT4 | **Blackwell-native**; 16-elem blocks, UE4M3 scales |
| MXFP4 | ~0.53 | useful 4-bit microscaling; ~0.5–1.0 ppl drop; slightly behind NVFP4 | OCP MX standard; 32-elem blocks |
| **GGUF Q4_K_M** | ~0.56–0.65 | strong local-chat default | llama.cpp / Ollama / LM Studio |
| INT4 GPTQ/AWQ | ~0.5 + metadata | calibration-dependent; ~0.3–1.5 ppl drop | legacy/common 4-bit path — "only if that is all you have" |
| GGUF Q8_0 | ~1.06 | ~99% of FP16 | near-lossless local |
| GGUF IQ3/IQ2 (+imatrix) | ~0.35 / ~0.22 | quality drops fast below 3-bit | extreme compression |

*Rule of thumb from the 2026 doctrine: low-batch/single-user → weight-only INT4-class; high-batch serving → W8A8 FP8 (INT4 dequant makes it slower at high batch); KV pressure → FP8 KV cache first (cheapest 2× concurrency).*

## 3.2 What the community actually uses

- **Desktop/local chat:** GGUF **Q4_K_M** remains the default sweet spot — Ollama auto-selects it by VRAM. Q8_0 when VRAM allows. IQ4_XS approaches Q4_K_M quality at slightly less VRAM.
- **Server self-hosting:** **FP8** for <400B models (near-lossless, mature kernels); **NVFP4** where the artifact ships that way (DeepSeek V4.1 Flash ships FP8+FP4 natively; SGLang shipped day-zero GLM-5.2 support including NVIDIA's NVFP4 checkpoint for Blackwell).
- **AWQ/GPTQ:** declining — kept for older model checkpoints that only exist in those formats.

## 3.3 The Blackwell FP4 caveat (important for 2026 buyers)

- NVFP4 dequantization is **free on Blackwell Tensor Cores** (native format); INT4 schemes need a separate dequant kernel that costs throughput.
- **But:** as of 2026, **no off-the-shelf stack delivers NVFP4 weights + NVFP4 KV cache on consumer Blackwell (SM120/RTX 5090/PRO 6000/DGX Spark):** vLLM's NVFP4 path falls back to Marlin (40–50% perf loss); SGLang's attention backends fail for MoE on SM120; TensorRT-LLM lacks NVFP4 KV-cache support. Only custom CUTLASS kernels close the gap. SM100 (datacenter B200) and SM120 (consumer Blackwell) need different kernels — code compiled for `sm_100a` traps on consumer Blackwell.
- Independent test (QuTLASS/MR-GPTQ MXFP4 on RTX 5090): genuine **~4× GEMM speedup** at batch ≥128, but **end-to-end decode lost to BF16 outright** at batch 1–32 (20.4 vs 78.6 tok/s at batch 1) — FP4 wins at GEMM level, loses at memory-bound decode on consumer cards.

**Bottom line for RAG:** recommend NVFP4 artifacts on datacenter Blackwell/Hopper; on consumer Blackwell and Apple Silicon, GGUF Q4_K_M (llama.cpp/Ollama) or MLX 4-bit remain the practical defaults. The FP4 software story on consumer Blackwell is unfinished.

