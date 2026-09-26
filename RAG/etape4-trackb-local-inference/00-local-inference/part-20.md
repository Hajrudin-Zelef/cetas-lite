---
id: etape4-trackb-local-inference/00-local-inference/part-20
title: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio) (part 20)"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["AMD", "Apple", "Intel", "Microsoft", "Qualcomm"]
dates: []
keywords: ["inference", "llama", "llama.cpp", "amd", "benchmark", "copilot", "gpu", "intel", "memory", "mistral", "throughput"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [841, 849]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 885f7815beb6f6eca473e8a319aaafe66fbef4d0246d3128c6fd96a1ed0be090
---

# Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio) (part 20)

The honest 2026 picture: NPUs matter for Copilot+ OS features (Studio Effects, Recall-type workloads) and small on-device assistants — **not** for serious local LLM inference, where GPU/iGPU/CPU still dominate:

- **Qualcomm Snapdragon X Elite (Hexagon NPU, 45 TOPS):** a community benchmark (LM Studio 0.3 + Ollama/DirectML, Win 11 24H2, Llama 3.1 8B Q4_K_M) measured ~26–31 tok/s on Snapdragon X Elite (QNN) vs 19–24 tok/s on Ryzen AI 9 HX 375 (DirectML, partial NPU) vs 13–17 tok/s on Intel Core Ultra 7 258V (OpenVINO, partial) — but TOPS ≠ usable throughput; driver/runtime maturity decides. [secondary — single-blogger benchmark, treat as indicative] (https://medium.com/@neilandrews1983/best-npu-laptops-2026-llama-mistral-benchmark-rankings-for-uk-buyers-c041a1804219)
- **AMD Ryzen AI NPUs (XDNA 2, up to 50 TOPS):** Linux XDNA driver stack matured enough in 2026 to run LLMs on the NPU directly (Phoronix/Michael Larabel testing) via ONNX Runtime + Vitis AI — a milestone, but the `ryzenai` backend in community libraries was still "awaiting hardware validation" as of mid-2026. [secondary] (https://www.webpronews.com/amds-ryzen-ai-npus-can-now-run-llms-locally-on-linux-heres-what-that-means/), (https://github.com/quintindk/crier)
- **Intel Core Ultra NPU:** OpenVINO path; llama.cpp OpenVINO backend still "in progress". [independent] (https://github.com/lunncing/llama.cpp-adaptive-kv-streaming — backend table, via search result)
- **Apple Neural Engine:** used via Core ML / MLX; Apple's on-device AI story is GPU+Neural-Engine combined, not NPU-only. [secondary]
- **Snapdragon X2 generation:** one AI-laptop comparison reports a "Snapdragon X2 Elite Extreme" with 80 TOPS NPU (highest NPU score in its table) but weaker memory bandwidth/GPU for memory-intensive inference. [secondary — single source, unverified] (https://medium.com/@raulasla/the-ai-laptop-chip-guide-nobody-wrote-how-to-actually-compare-processors-in-2026-d63b4a9476f2)
- Bottom line: NPUs in 2026 are a complement for always-on/assistant workloads; local LLM inference still runs on GPU (CUDA/Metal/Vulkan), with the NPU contributing only where the runtime explicitly supports it (QNN/OpenVINO/Vitis AI). Flag any vendor TOPS figure as marketing input, not inference throughput.

