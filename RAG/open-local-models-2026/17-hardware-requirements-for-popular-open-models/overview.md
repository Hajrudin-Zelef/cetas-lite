---
id: open-local-models-2026/17-hardware-requirements-for-popular-open-models/overview
title: "2. HARDWARE REQUIREMENTS FOR POPULAR OPEN MODELS"
domain: hardware-requirements-for-popular-open-models
role: deep-dive
task: hardware
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Mistral", "Nvidia", "SGLang", "Z.ai", "vLLM"]
dates: []
keywords: ["agents", "amd", "blackwell", "consumer", "decode", "deepseek", "fp4", "fp8", "gguf", "glm", "gpu", "gpus"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [834, 909]
section: "2. HARDWARE REQUIREMENTS FOR POPULAR OPEN MODELS"
sha256: fac49437427b3fbcb036b1f2f1e439552c11943ee7d9925c8b7d8cbba2f7e82e
---

# 2. HARDWARE REQUIREMENTS FOR POPULAR OPEN MODELS

## 2.1 The VRAM math

```
VRAM ≈ Parameters (B) × Bytes per parameter + KV-cache + runtime overhead
```

| Quantization | Bytes/param | 7B | 13B | 32B/34B | 70B | Quality vs FP16* |
|---|---|---|---|---|---|---|
| FP16/BF16 | 2.00 | 14 GB | 26 GB | 64–68 GB | 140 GB | baseline |
| Q8_0 (8-bit) | ~1.06 | 7.5 GB | ~14 GB | 34–36 GB | 74 GB | ~99% (near-identical) |
| Q6_K | ~0.81 | 5.7 GB | — | — | 57 GB | ~97% (excellent) |
| Q5_K_M | ~0.69 | 4.8 GB | — | — | 48 GB | ~95% (good) |
| **Q4_K_M** | **~0.56** | **3.9 GB** | **~8 GB** | **18–20 GB** | **39 GB** | **~90% (acceptable)** |
| Q3_K_M | ~0.44 | 3.1 GB | — | — | 31 GB | ~80% (noticeable loss) |
| Q2_K | ~0.31 | 2.2 GB | — | — | 22 GB | ~70% (significant loss) |
| FP8 (E4M3) | 1.0 | 7 GB | — | — | 70 GB | ~100% (near-lossless) |

*\*Approximate, task-dependent. Chat is forgiving; coding/reasoning less so. Community rule: "Before dropping to Q3 to fit a model, ask: would a smaller model at Q5 give better results? The answer is usually yes."*

**KV cache — the underestimated factor:** grows linearly with context and batch. For a 70B GQA model: ~10 GB extra at 32K context (FP16 KV); ~24.6 GB at 75K tokens. Doubling context doubles KV cache. FP8 KV cache is the cheapest 2× concurrency lever.

**CPU offload fallback:** if weights exceed VRAM, layers spill to system RAM — works but 5–20× slower for offloaded layers. Partial offload (10–20%) is often acceptable.

Sources: https://willitrunai.com/blog/vram-requirements-for-ai-models · https://willitrunai.com/blog/quantization-q4-q8-fp16-explained · https://github.com/casteldazur/awesome-local-ai/blob/HEAD/guides/vram-requirements.md · https://www.alekseialeinikov.com/en/blog/topics/ai/quantization-explained-run-70b-model-consumer-hardware-2026 · https://github.com/drajb/gekro/blob/HEAD/apps/web/src/content/apps/gpu-vram-calculator.md

## 2.2 What fits where (2026 hardware classes)

### Consumer GPUs (24 GB — RTX 4090 class)
- Comfortable: up to **~32–34B at Q4_K_M** (18–20 GB).
- 70B at Q4_K_M (39 GB) needs significant CPU offload on 24 GB — usable but slow.
- 8B models at Q4 (~5 GB): ~95 tok/s reported on RTX 4090-class (llama.cpp eval).
- Note: during decode, **memory bandwidth** (not just capacity) dominates speed once the model fits.

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

