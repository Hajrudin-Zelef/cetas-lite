---
id: etape4-trackb-local-inference/00-local-inference/part-18
title: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio) (part 18)"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["Alibaba", "Microsoft", "Nvidia"]
dates: ["2025-03", "2026-06"]
keywords: ["llama", "benchmarks", "blackwell", "consumer", "datacenter", "decode", "fp4", "gpu", "gpus", "lpddr5x", "memory", "nvfp4"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [785, 796]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 4d66ae9f53c81678528f8e5d9cfa27b594fc04b39cf7a28876adb2c75ef12035
---

# Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio) (part 18)

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

