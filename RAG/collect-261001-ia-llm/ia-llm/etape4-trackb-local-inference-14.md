---
id: collect-261001-ia-llm/ia-llm/etape4-trackb-local-inference-14
title: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Apple", "Intel", "Microsoft", "Nvidia", "TensorRT-LLM", "Unsloth"]
dates: ["2025-03", "2026-06"]
keywords: ["inference", "llama", "llama.cpp", "amd", "benchmarks", "blackwell", "compute", "consumer", "datacenter", "decode", "fine-tuning", "fp4"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [785, 816]
sha256: 39ac161e660b434e2b9032d67507d1740fca7308d13f29cfa91478d23af67218
---

# Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)

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

### 2.3 AMD

**Radeon RX 9070 XT** (RDNA4, gfx1201):
- 16 GB GDDR6, 640 GB/s bandwidth, MSRP $599; street from ~$740 (early Aug 2026). Board power 304 W. [independent] (https://localaimaster.com/blog/rx-9070-xt-local-ai)
- Local inference: ROCm backend on Linux (gfx1150/gfx1201 target, ROCm 6.3+; ROCm 7 in newer coverage) delivers 35–42 tok/s on 8B models; Vulkan backend works cross-platform (~15% slower); llama.cpp's official ROCm Docker image natively supports RDNA4 (gfx1201). Community run: Qwen3.6-27B IQ3_M at 46 tok/s decode with MTP speculative decoding (llama.cpp + ROCm Docker, 16k context, 8-bit KV). [independent] (https://www.compute-market.com/blog/rx-9070-xt-vs-rtx-5060-ti-local-ai-2026), (https://midwest.social/post/48528597)
- vs RTX 5060 Ti 16GB ($429 MSRP): 9070 XT's 43% bandwidth advantage shows up ~linearly in tok/s (~52 vs ~33 tok/s on 14B Q4); the 5060 Ti buys you CUDA (ExLlamaV2, TensorRT-LLM, mainstream fine-tuning) and 124 W lower power. Verdict from independent coverage: choose 9070 XT for Linux/fastest-16GB-inference or gaming; choose 5060 Ti for the CUDA ecosystem. [independent] (https://localaimaster.com/blog/rx-9070-xt-local-ai)
- No FP4 support on any AMD consumer GPU (NVIDIA-exclusive). [secondary] (https://www.compute-market.com/blog/rx-9070-xt-vs-rtx-5060-ti-local-ai-2026)
- Easiest path for AMD users: LM Studio with Vulkan backend (Windows & Linux) — no ROCm setup; ROCm only to squeeze the last bit of speed. [independent] (https://github.com/bokiko/localist/blob/HEAD/guides/amd-gpu.md)
- 24 GB alternative: RX 7900 XTX — runs 32B-class models a 16 GB card cannot hold; the community "spend more, get 24 GB" option. [secondary] (https://localaimaster.com/blog/rx-9070-xt-local-ai)

**Ryzen AI Max+ 395 "Strix Halo"** (16× Zen 5, 40-CU Radeon 8060S iGPU / gfx1151, up to 128 GB LPDDR5x-8000 unified):
- The Windows answer to Apple Silicon for local LLM: 128 GB unified memory with ~256 GB/s theoretical (~215 GB/s measured) bus. Massive capacity, modest bandwidth — token generation is bandwidth-bound, so expect single-digit tok/s on dense 70B, but huge models fit: MoE models fly (Qwen3 30B MoE ~66–72 tok/s; LFM2-24B-A2B ~109 tok/s; ~100 tok/s on 30B MoE per setup guides). 70B (Q4) ~15–20 tok/s (ROCm/rocWMMA, fits entirely in memory). [independent] (https://github.com/nelisw/framework-strix-halo-llm-setup), (https://codersera.com/blog/amd-strix-halo-ryzen-ai-max-local-llm-setup-2026/), (https://digitalarchitects.hr/insights/amd-ryzen-ai-max-395-local-llm/)
- Practical setup (2026): allocate ~96 GB to iGPU via AMD Adrenalin (Windows) or GTT kernel parameter + small BIOS framebuffer (Linux, kernel ≥ 6.16.9/6.18); run LM Studio / Ollama / llama.cpp on Vulkan backend (most reliable as of mid-2026; ROCm HIP path improving but had stability issues; community benchmarks show Vulkan outperforming HIP). [independent] (https://codersera.com/blog/amd-strix-halo-ryzen-ai-max-local-llm-setup-2026/), (https://digitalarchitects.hr/insights/amd-ryzen-ai-max-395-local-llm/)
- Price points: Framework Desktop ~$1,999 (128 GB), GMKtec mini PCs ~$1,499–$2,500, AMD Ryzen AI Halo dev platform $3,999; ASRock Industrial AI BOX-A395 announced Mar 17, 2026 (200×100×232 mm; pricing unannounced). [secondary] (https://digitalarchitects.hr/insights/amd-ryzen-ai-max-395-local-llm/), (https://videocardz.com/newz/asrock-industrial-launches-ai-box-a395-with-ryzen-ai-max-395-and-128gb-lpddr5x)
- Price delta vs DGX Spark: Strix Halo platforms generally undercut Spark by $1,000–$2,000 at comparable 128 GB. [secondary] (http://gpusmith.com/articles/en/pdfs/nvidia-dgx-spark-review-specs-performance.pdf)

**ROCm for local inference (2026 status):** supported on RDNA4 (gfx1200/gfx1201) and Strix Halo (gfx1151); official llama.cpp ROCm Docker images; LM Studio ships ROCm builds. Maturing but still behind CUDA in ecosystem (no Unsloth, no TensorRT-LLM equivalent); fine-tuning on ROCm via PyTorch works but slower than CUDA. [independent] (https://www.compute-market.com/blog/rx-9070-xt-vs-rtx-5060-ti-local-ai-2026), (https://localaimaster.com/blog/rx-9070-xt-local-ai)

### 2.4 Intel

