---
id: collect-261001-ia-llm/ia-llm/etape4-trackb-local-inference-17
title: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Apple", "DeepSeek", "Google", "Hugging Face", "Intel", "Meta", "Moonshot", "Nvidia", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-08-16", "2026-09-07", "2026-09-22"]
keywords: ["inference", "llama", "llama.cpp", "agent", "agentic", "benchmark", "blackwell", "compute", "cost", "deepseek", "fine-tuning", "fp4"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [938, 986]
sha256: 8edf75910427589eb3973dfd761bdb356b8f5a58be405322de744d55e379a395
---

# Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)

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

- The in-app catalog is the product's signature flow: search → Hugging Face pull → memory-fit estimate → download → chat; supports GGUF and MLX quants; no terminal required. [secondary] (https://medium.com/@nishilbhave/lm-studio-in-2026-download-models-run-local-llms-vs-ollama-b30567df17f8)
- 2026 catalog landmarks trace the open-model frontier: `openai/gpt-oss-20b` (120B sibling runs on 128 GB unified machines), Kimi K2.7 Code / K3 (Moonshot), GLM 5.2/5.3-Flash (Z.ai), Qwen3.5/3.6/3.8 (Alibaba), Gemma 4 (Google), DeepSeek V4 Pro (cloud), NVIDIA Nemotron-Nano-v2. [official/secondary]
- LM Link free tier: up to 5 devices; paid/enterprise details unpublished at cutoff. [secondary] (https://learning.christiandrapatz.de/lmstudio-en.pdf)
- Community position vs Ollama (2026): LM Studio = GUI-first, Ollama = CLI-first; both llama.cpp-based so raw tok/s nearly identical; Ollama added paid cloud tiers in 2026 (Pro $20/mo, Max $100/mo) while local stays free; LM Studio remains free locally with paid Secure Cloud credits. [secondary] (https://www.kunalganglani.com/blog/lm-studio-vs-ollama)
- vLLM vs llama.cpp at scale (Red Hat 2026 benchmark, cited in independent coverage): at 64 concurrent users vLLM generated ~44× more tokens/s than llama.cpp — i.e., llama.cpp/LM Studio is the single-user king, vLLM the serving king. [secondary] (https://medium.com/@nishilbhave/local-llms-in-2026-which-runtime-to-run-and-the-hardware-you-need-a88450dece2e)
- Menlo Ventures (2025, cited 2026): open-source models hold ~11% of enterprise LLM usage, down from 19% — local inference remains a niche vs hosted APIs, which frames LM Studio's audience. [secondary] (same)

---

*End of report. Research cutoff: September 22, 2026.*


---


## Master open-verification log (track B, carried forward)

