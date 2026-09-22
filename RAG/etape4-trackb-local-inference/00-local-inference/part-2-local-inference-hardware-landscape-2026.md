---
id: etape4-trackb-local-inference/00-local-inference/part-2-local-inference-hardware-landscape-2026
title: "PART 2 — LOCAL INFERENCE HARDWARE LANDSCAPE (2026)"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: hardware
actors: ["AMD", "Alibaba", "Microsoft", "Nvidia", "TensorRT-LLM"]
dates: ["2025-03", "2026-06", "2026-07"]
keywords: ["inference", "amd", "attention", "benchmark", "benchmarks", "blackwell", "compute", "consumer", "datacenter", "decode", "flash attention", "fp4"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [746, 796]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 1036961161c5b0d17220d3c49d1b4087a443153a08def7760b490ed3809cd5fb
---

# PART 2 — LOCAL INFERENCE HARDWARE LANDSCAPE (2026)

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

**DGX Spark** (GB10 Grace Blackwell Superchip; formerly "Project Digits"):
- Status: launched Oct 15, 2025 (global availability); announced at CES Jan 2025 as Project Digits, renamed at GTC March 2025. [secondary] (https://www.techedt.com/nvidia-launches-dgx-spark-personal-ai-supercomputer-on-15-october), (https://www.aitooldiscovery.com/ai-infra/nvidia-dgx-spark-explained)
- Specs: 20-core Arm CPU (10× Cortex-X925 + 10× Cortex-A725), Blackwell GPU w/ 6,144 CUDA cores, 5th-gen Tensor Cores, 1 PFLOP FP4 (sparse), **128 GB LPDDR5x unified memory, 273 GB/s**, 4 TB NVMe, 150×150×50.5 mm, 1.2 kg, DGX OS (Ubuntu 24.04-based). [independent] (https://www.tomshardware.com/pc-components/gpus/nvidia-dgx-spark-review/2), (https://peterfalkingham.com/2026/09/18/academic-tech-gigabyte-ai-top-nvidia-dgx-spark-review/)
- Pricing: $3,999 MSRP at launch (Oct 15, 2025); **raised to $4,699 on Feb 23, 2026** (+$700, ~18%) citing LPDDR5x supply constraints; no hardware change. [secondary, NVIDIA forum statement relayed] (http://gpusmith.com/articles/en/pdfs/nvidia-dgx-spark-review-specs-performance.pdf)
- Performance (independent): LMSYS — GPT-OSS 20B in Ollama: ~2,053 tok/s prefill, 49.7 tok/s decode (~1/4 of RTX Pro 6000, slower than a single RTX 5090 at 8,519/205); 70B models ~35–45 tok/s; GPT-OSS-120B ~38 tok/s single-user; batched Llama 3.1 8B at batch 32 ≈ 368 tok/s decode. **The 273 GB/s bandwidth, not the 1-PFLOP headline, determines throughput** — consistent reviewer consensus. [secondary summarizing LMSYS/The Register/Raschka] (same)
- 2026 software progress: CES 2026 enterprise update claims up to 2.5× on key workloads vs launch (NVIDIA cites 2.6× on Qwen-235B across 2 units with NVFP4 + speculative decoding); multi-node clustering up to 4 units via Cluster Assistant (June 2026 release); ~200B params FP4 per unit, ~405B with two units over ConnectX-7 (200 GbE). [secondary] (https://emarque.co/collections/nvidia-dgx-spark)
- Caveats reported: thermal/power-delivery issues on early units (John Carmack publicly reported power capping ~100W of rated 240W and reboots under load; NVIDIA forums confirmed a known platform issue); consumer-class Blackwell GPU causes software-compat friction vs datacenter cards. [secondary] (http://gpusmith.com/articles/en/pdfs/nvidia-dgx-spark-review-specs-performance.pdf)
- Also announced: **RTX Spark** family (Windows AI laptops/desktops on "RTX Spark Superchip", up to 1 PFLOP FP4, up to 128 GB unified, from ASUS/Dell/HP/Lenovo/Microsoft/MSI — Fall 2026 target; distinct from DGX Spark). [secondary] (https://emarque.co/collections/nvidia-dgx-spark)
- **DGX Station** (GB300 Grace Blackwell Ultra Desktop Superchip): 20 PFLOP, 784 GB unified memory; 1T+ param models on one unit — the step above Spark. [secondary] (https://www.aitooldiscovery.com/ai-infra/nvidia-dgx-spark-explained)
- **RTX PRO 6000 Blackwell Workstation Edition**: 96 GB GDDR7, $8,500+ (HotHardware cited at ~$10,000 before host workstation). [secondary] (http://gpusmith.com/articles/en/pdfs/nvidia-dgx-spark-review-specs-performance.pdf)
- ServeTheHome coverage exists (Jetson Thor / DGX Spark ecosystem); STH's Thor coverage noted preliminary specs and bandwidth-sensitive benchmarks. [independent] (http://aiwiki.ai/wiki/jetson_thor)

