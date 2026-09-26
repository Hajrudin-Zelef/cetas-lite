---
id: etape4-trackb-local-inference/00-local-inference/open-questions-uncertainties-flagged-for-follow-up
title: "OPEN QUESTIONS / UNCERTAINTIES (flagged for follow-up)"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "Apple", "Google", "Intel", "Meta", "Nvidia", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-08-16", "2026-09-07"]
keywords: ["agent", "agentic", "amd", "benchmark", "blackwell", "compute", "cost", "fine-tuning", "fp4", "gguf", "glm", "inference"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [923, 969]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 946e3b8cc666f761ed8df62e7bddf4374559a1f04237ca119c4e45d73b460b96
---

# OPEN QUESTIONS / UNCERTAINTIES (flagged for follow-up)

## OPEN QUESTIONS / UNCERTAINTIES (flagged for follow-up)

1. Exact LM Studio download/user counts — no vendor-published figure found; "millions of downloads" is secondary only. [unverified]
2. LM Studio Enterprise tier and Secure Cloud pricing — not published at cutoff. [unverified]
3. DFlash / DSpark assistant drafters in LM Studio 0.4.22 — changelog-confirmed names only; architectural details not verified. [unverified]
4. IPEX-LLM archival (Jan 28, 2026) — from a community research doc; not independently confirmed against Intel's repo at cutoff. [unverified]
5. Snapdragon X2 Elite / X2 Elite Extreme specs (80 TOPS claim) — single secondary source. [unverified]
6. DGX Spark "2.5×" post-launch software speedup and GPT-OSS-120B 58.8 tok/s (vLLM/MXFP4) figures — vendor/community-reported, not independently reproduced here. [vendor-reported/secondary]
7. Street prices (RTX 5090 ~$2,500–$3,799; RX 9070 XT ~$740) move with GDDR7/memory supply; figures are dated snapshots (Aug–Sep 2026). [secondary]
8. AMD's listed price for the Ryzen AI Halo developer platform ($3,999) and ASRock AI BOX-A395 pricing (unannounced) — from secondary coverage. [secondary]

---

## APPENDIX A — 2026 TIMELINE (LM Studio / local hardware / GGUF-relevant)

| Date | Event | Provenance |
|---|---|---|
| Jan 28, 2026 | LM Studio 0.4.0: GUI/core split (llmster daemon), continuous batching, native REST API `/api/v1/*`, UI refresh | [official] (https://lmstudio.ai/blog/0.4.0) |
| Jan 28, 2026 (reported) | Intel IPEX-LLM archived (read-only) for "known security issues" | [unverified] (https://github.com/chriscorbell/llm-server/blob/HEAD/docs/research/2026-09-07-engines-on-battlemage.md) |
| Feb 23, 2026 | NVIDIA raises DGX Spark Founders Edition MSRP $3,999 → $4,699 (+18%), citing LPDDR5x supply constraints | [secondary] (http://gpusmith.com/articles/en/pdfs/nvidia-dgx-spark-review-specs-performance.pdf) |
| Feb 27, 2026 | LM Studio 0.4.6: LM Link (Tailscale partnership) introduced | [official] (https://lmstudio.ai/changelog/lmstudio/lmstudio-v0.4.6) |
| Mar 17, 2026 | ASRock Industrial announces AI BOX-A395 (Ryzen AI Max+ 395, 128 GB LPDDR5x-8000) | [secondary] (https://videocardz.com/newz/asrock-industrial-launches-ai-box-a395-with-ryzen-ai-max-395-and-128gb-lpddr5x) |
| ~Mar–Apr 2026 | llama.cpp Q8_0 reorder fix merged for SYCL (Apr 7); Intel llm-scaler-vllm B70 support official since 0.14.0-b8.2 | [secondary] (https://github.com/chriscorbell/llm-server/blob/HEAD/docs/research/2026-09-07-engines-on-battlemage.md) |
| Jun 4, 2026 | LM Studio 0.4.16 ships with Locally iPhone/iPad app + LM Link remote access | [secondary] (https://www.digitalapplied.com/blog/lm-studio-locally-lm-link-iphone-local-llm-2026) |
| Jun 2026 | DGX Spark multi-node clustering (up to 4 units) via Cluster Assistant; software claims up to 2.5× vs launch on key workloads | [secondary] (https://emarque.co/collections/nvidia-dgx-spark) |
| Jul 16, 2026 | LM Studio Bionic announced (agentic app; local + Secure Cloud execution) | [secondary] (https://9to5mac.com/2026/07/16/lm-studio-expands-beyond-chat-with-bionic-a-new-ai-agent-app-for-open-models/) |
| Jul 2026 (announced) | Jetson T3000 / T2000 modules announced; Q1 2027 availability | [independent] (http://aiwiki.ai/wiki/jetson_thor) |
| ~Jul 2026 | FP4 inference "being integrated into llama.cpp's CUDA backend" (NVIDIA Blackwell NVFP4) | [secondary] (https://www.compute-market.com/blog/rx-9070-xt-vs-rtx-5060-ti-local-ai-2026) |
| Aug 26, 2026 | GLM-5.3-Flash (320B MoE, 18B active, 1M context) added to Bionic Secure Cloud | [secondary] (https://9to5mac.com/2026/08/26/lm-studio-adds-glm-5-3-flash-to-bionic-with-image-support-and-1m-token-context/) |
| Aug 27, 2026 | Bionic 1.1.0: 2–2.75× faster prompt processing on M5 Macs; Muse Glimmer tool calling for MLX | [official] (https://lmstudio.ai/changelog) |
| Aug 28, 2026 | LM Studio 0.4.22: DFlash/DSpark/MTP assistant drafters (speculative decoding) | [official] (https://lmstudio.ai/changelog/lmstudio/lmstudio-v0.4.22) |
| Sep 8–19, 2026 | Bionic 1.1.2 (Linux builds) → 1.1.4 → 1.1.5 (Splash engine by Inco AI for Qwen3.8 on Mac) | [official] (https://lmstudio.ai/changelog) |
| Sep 9, 2026 | **LM Studio 0.4.24 (latest at cutoff)**: advanced llama.cpp GGUF-load overrides, drafter heuristics, `/api/v1/chat` image fix | [official] (https://lmstudio.ai/changelog/lmstudio/lmstudio-v0.4.24) |

## APPENDIX B — ADDITIONAL HARDWARE NOTES

- **Mac Studio (M3 Ultra)** remains the bandwidth king of unified-memory machines in 2026 coverage: up to 512 GB unified memory at up to ~800–819 GB/s — vs DGX Spark's 273 GB/s; frequently cited as Spark's main capacity-class rival alongside Strix Halo. [secondary] (http://gpusmith.com/articles/en/pdfs/nvidia-dgx-spark-review-specs-performance.pdf)
- **Jetson Orin family** runs JetPack 6.x (Ubuntu 22.04, aarch64) with CUDA/cuDNN/TensorRT at no extra software cost; Jetson Nano (Maxwell) discontinued; Xavier NX/AGX Xavier approaching EOL. [independent] (https://github.com/alpininsight/capi-provider-ssh/blob/HEAD/docs/roadmap/nvidia-jetson-edge-devices.md)
- **RTX 5060 Ti 16 GB** ($429 MSRP, 16 GB GDDR7, 448 GB/s, 180 W): 5th-gen Tensor Cores with FP4; the CUDA-ecosystem budget pick at 16 GB. [independent] (https://www.compute-market.com/blog/rx-9070-xt-vs-rtx-5060-ti-local-ai-2026), (https://localaimaster.com/blog/rx-9070-xt-local-ai)
- **Used market:** RTX 3090 (24 GB) at $699–$999 is the community value buy for 24 GB VRAM (fine-tuning with Unsloth, 32B inference). [secondary] (https://www.compute-market.com/blog/rx-9070-xt-vs-rtx-5060-ti-local-ai-2026)
- **MoE efficiency note (2026):** on bandwidth-limited unified-memory systems, MoE models (gpt-oss-20B, Qwen3-30B-A3B, LFM2-24B-A2B) are dramatically faster than dense models of similar total size — e.g., gpt-oss:20b at ~92 tok/s on a 16 GB RX 9070 XT, where a dense 27B Q4 would be unusable. [independent] (https://github.com/hirokuze/local-llm-benchmark-rx9070xt), (https://localaimaster.com/blog/rx-9070-xt-local-ai)
- **KV-cache quantization** is now standard practice for fitting bigger contexts: 8-bit KV (`--cache-type-k q8_0 --cache-type-v q8_0`) saves ~50% VRAM; vllm-mlx documents 4-bit/8-bit/fp16 KV options. [independent] (https://midwest.social/post/48528597), (https://github.com/anisoptera/vllm-mlx-upstream)
- **Multi-token prediction (MTP) speculative decoding** is the 2026 speed lever: community reports 46 tok/s on Qwen3.6-27B IQ3_M (RX 9070 XT, 62.7% draft acceptance); Google MTP drafters for Gemma; LM Studio 0.4.22+ ships DFlash/DSpark/MTP assistant-drafter support; Bionic enables MTP for more models. [independent/official] (https://midwest.social/post/48528597), (https://lmstudio.ai/changelog/lmstudio/lmstudio-v0.4.22)
- **oMLX** (18.8k GitHub stars, Aug 2026) — a macOS-native multi-model inference server with tiered KV cache (hot blocks in unified RAM, cold spilled to SSD) — illustrates the Apple-Silicon power-user niche between llama.cpp and GUI tools like LM Studio. [independent] (https://github.com/hongsw/clawfit/blob/HEAD/docs/research-watch/2026-08-16-omlx-apple-silicon-llm-inference-server-ssd-kv-cache.md)

## APPENDIX C — LM STUDIO MODEL-CATALOG & ECOSYSTEM NOTES

