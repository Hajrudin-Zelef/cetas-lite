---
id: etape4-trackb-local-inference/00-local-inference/1-6-pricing
title: "1.6 Pricing"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: pricing
actors: ["AMD", "Alibaba", "Apple", "Intel", "Nvidia", "TensorRT-LLM"]
dates: ["2026-07"]
keywords: ["pricing", "acquisition", "amd", "attention", "benchmark", "blackwell", "compute", "consumer", "decode", "flash attention", "fp4", "gguf"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [733, 784]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: c0205c2e8893618cb76bddf4e51f381961dc25fe9249a78531154a66513162b3
---

# 1.6 Pricing

- Inference engines: **llama.cpp** (primary, CPU/GPU, all platforms) and **Apple MLX** (Apple Silicon M1–M5); both are independently downloadable runtimes — engine updates ship separately from app updates (`lms runtime update`). [secondary cross-checked with official 0.4.0 install docs] (https://github.com/aterrylu/autonomos/blob/HEAD/docs/research/desktop-shells/lm-studio.md)
- Backend coverage on desktop: CUDA (NVIDIA), Metal (Apple), Vulkan (cross-vendor, incl. AMD/Intel); changelog entries also reference ROCm builds and multi-GPU selection bugs on CUDA 12 / ROCm / Vulkan. [official] (https://lmstudio.ai/changelog/lmstudio/lmstudio-v0.4.22)
- NPU: LM Studio runs on the user's GPU or NPU per third-party coverage; NPU-specific offload reported in community NPU laptop tests (see §2.4). Flag: exact NPU backends in LM Studio are not enumerated in official docs found at cutoff [secondary/unverified]. [secondary] (http://www.howtogeek.com/you-can-make-a-self-hosted-ai-server-with-lm-studio-040/)

### 1.6 Pricing

- **Desktop app and local use: free** (home and work). [secondary — widely reported, incl. LM Studio comparison coverage] (https://www.aixploria.com/en/lm-studio-ai/), (https://www.kunalganglani.com/blog/lm-studio-vs-ollama)
- One comparison site reports an **Enterprise tier** for teams needing LM Link multi-device workload routing and priority support — but this is not confirmed on LM Studio's public pricing pages at cutoff; treat as [unverified]. (https://www.kunalganglani.com/blog/lm-studio-vs-ollama)
- **LM Studio Secure Cloud**: pay-as-you-go via Cloud Credits; exact pricing "not published" as of July 2026 per a Bionic developer guide; "Bionic Pass" subscription announced as "coming soon". LM Link free tier covers up to 5 devices. [secondary] (https://learning.christiandrapatz.de/lmstudio-en.pdf)
- LM Studio acquisition of Locally AI (mobile, Adrien Grondin) — strategic push to iPhone/iPad on-device AI; iOS local models realistically 1B–3B at practical speeds. [secondary] (https://evermx.com/case/lm-studio-acquires-locally-ai-mobile-on-device-ai)

---

## PART 2 — LOCAL INFERENCE HARDWARE LANDSCAPE (2026)

### 2.1 Recommended GPUs / VRAM for 2026 flagship open models

Community-maintained VRAM math (llama.cpp/GGUF ecosystem) — the standard estimator is `VRAM ≈ params × bytes-per-weight × ~1.1–1.2` (KV cache + overhead):

| Model / quant | Approx weights | Practical VRAM incl. context | Source |
|---|---|---|---|
| 7–8B, Q4_K_M | ~5 GB | ~6 GB (entry: 8–12 GB card) | [independent] (https://github.com/high-altitude-ai/local-llm-rig) |
| 13–14B, Q4_K_M | ~8.7–9.1 GB | ~10 GB | [independent] same + (https://github.com/casteldazur/awesome-local-ai/blob/HEAD/guides/vram-requirements.md) |
| 32–34B, Q4_K_M | ~20 GB | ~22 GB | [independent] (https://github.com/high-altitude-ai/local-llm-rig) |
| 70B, Q4_K_M | ~40–43 GB file | ~45–48 GB loaded | [independent] (https://github.com/high-altitude-ai/local-llm-rig), (https://github.com/joyroy9454/aryvora/blob/HEAD/content/posts/run-ai-locally-ollama-llama-cpp-guide.md) |
| 70B, Q8_0 | ~78 GB | ~85 GB+ | [independent] (same) |
| 120B (gpt-oss-120B), MXFP4/NVFP4 | — | runs on DGX Spark 128GB at ~38 tok/s single user; ~505 tok/s batched on 2-unit cluster | [secondary] (https://emarque.co/collections/nvidia-dgx-spark) |

Key community takeaways [independent] (https://github.com/high-altitude-ai/local-llm-rig):
- Q4_K_M is the default sweet spot, not a compromise (small quality delta vs FP16, 4× VRAM saving).
- The painful gap is 24 GB → 48 GB: 24 GB comfortably runs 32B-class models; 70B at usable quality needs ~48 GB (2× cards or unified-memory silicon).
- Long context eats VRAM: 32k context on 8B adds gigabytes of KV cache; budget +20–30% for long-doc use.
- A 70B Q4 model "wants roughly 40 GB of memory", so a 24 GB GPU or a 64 GB+ Mac is the practical entry point for big models. [independent] (https://medium.com/@nishilbhave/local-llms-in-2026-which-runtime-to-run-and-the-hardware-you-need-a88450dece2e)

**GPU ladder for local inference in 2026 (community consensus):**
- 16 GB (RTX 5060 Ti 16GB, RX 9070 XT, Arc B580): up to ~13–14B at Q4, 20B MoE; 16 GB "enough for most use cases" but 30B+ models hit a ceiling. [independent] (https://www.compute-market.com/blog/rx-9070-xt-vs-rtx-5060-ti-local-ai-2026), (https://localaimaster.com/blog/rx-9070-xt-local-ai)
- 24 GB (RTX 4090, RTX 3090 used, RX 7900 XTX, Arc Pro B60): 32B-class comfortably; used RTX 3090 at $699–$999 is the community value pick for 24 GB. [secondary] (same)
- 32 GB (RTX 5090): 27B dense, 35B MoE comfortably; 70B only with aggressive quant/split offload. [independent] (https://github.com/sipeed/llmdev.guide/blob/HEAD/devices/nvidia-rtx-5090-32gb.md)
- 96 GB (RTX PRO 6000 Blackwell Workstation Edition): runs $8,500+; GPT-OSS-120B-class locally on one card. [secondary] (http://gpusmith.com/articles/en/pdfs/nvidia-dgx-spark-review-specs-performance.pdf)
- 128 GB unified (DGX Spark, Strix Halo systems, Mac Studio): the "run 70B+ without compromise" tier; 200B-class at FP4 on DGX Spark. [independent/secondary] (https://emarque.co/collections/nvidia-dgx-spark)

### 2.2 NVIDIA

**GeForce RTX 5090** (Blackwell GB202; released Jan 30, 2025, announced at CES Jan 6, 2025):
- 32 GB GDDR7, 512-bit bus, 1,792 GB/s (~1.79 TB/s) bandwidth; 21,760 CUDA cores; 680 5th-gen Tensor Cores; 104.8 TFLOPS FP32; 575 W TDP; MSRP $1,999 (Founders Edition). [independent] (https://www.runpod.io/articles/guides/nvidia-rtx-5090), (https://www.pcgamer.com/hardware/graphics-cards/nvidia-announces-the-rtx-50-series-led-by-the-usd1-999-rtx-5090-with-twice-the-performance-of-the-4090/)
- Street price reality: ~$2,500–$3,200 secondary market; one benchmark guide reports street ~$3,799 early 2026 due to GDDR7 constraints. Supply tight much of 2025–2026. [secondary] (https://github.com/sipeed/llmdev.guide/blob/HEAD/devices/nvidia-rtx-5090-32gb.md)
- Inference: highest decode throughput among consumer GPUs (bandwidth-bound decode); community numbers: Llama 7B Q4_0 ≈ 300 tok/s decode / ~15k tok/s prefill; Qwen3.5-35B-A3B int4 ≈ 194 tok/s (llama.cpp CUDA + Flash Attention). [independent] (same)
- NVFP4 hardware (dual-level scaling) is consumer-first on Blackwell; FP4 supported in TensorRT-LLM and being integrated into llama.cpp's CUDA backend as of July 2026 — an NVIDIA-exclusive advantage, no AMD consumer GPU supports FP4. [secondary] (https://www.compute-market.com/blog/rx-9070-xt-vs-rtx-5060-ti-local-ai-2026)

**RTX 5080**: 10,752 CUDA cores, 16 GB GDDR7, 960 GB/s, 360 W, MSRP $999. [independent] (https://www.pcgamer.com/hardware/graphics-cards/nvidia-announces-the-rtx-50-series-led-by-the-usd1-999-rtx-5090-with-twice-the-performance-of-the-4090/)
**RTX 5070 Ti / 5070**: 16 GB / 12 GB GDDR7; $749 / $549. [independent] (same)

