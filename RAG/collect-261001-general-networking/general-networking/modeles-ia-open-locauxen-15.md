---
id: collect-261001-general-networking/general-networking/modeles-ia-open-locauxen-15
title: "ÉTAPE 1 — Open / Local AI Models (EN)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Alibaba", "Apple", "ByteDance", "Hugging Face", "Mistral", "Moonshot", "Nvidia", "OpenAI", "United States", "Z.ai", "vLLM"]
dates: ["2025-06", "2026-04", "2026-04-08", "2026-05-22", "2026-06", "2026-07-16"]
keywords: ["acquisition", "agent", "agentic", "agents", "amd", "attention", "benchmark", "compute", "consumer", "copilot", "decode", "fp8"]
source: docs/RAG/collect-261001-general-networking/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [792, 868]
sha256: c0395feed05505a7749d812953a83324c3d5cc5ac1775f86c4f17d41da8c43c7
---

# ÉTAPE 1 — Open / Local AI Models (EN)

**Positioning 2026:** the API-first headless local runner — best for terminal users, daemons, REST APIs, Docker workflows, repeatable model pulls, local agents/automation. "If you want a headless local model server for scripts, agents, Docker, CI jobs, or production-like APIs, choose Ollama."

Sources: https://medium.com/@tentenco/ollama-0-19-ships-mlx-backend-for-apple-silicon-local-ai-inference-gets-a-real-speed-bump-878b4928f680 · https://aifoss.dev/blog/llamafile-vs-ollama-vs-lm-studio-2026/ · https://medium.com/@_lukasz_/local-coding-models-in-2026-why-mlx-beat-ollama-on-apple-silicon-68118d806c40 · https://pooyagolchian.com/blog/github-copilot-ollama-agentic-local-llm-2026/ · https://localclaw.io/blog/ollama-vs-lm-studio-2026

## 1.2 llama.cpp

**What it is:** the C/C++ inference engine under the whole local-LLM world (Ollama, LM Studio's llama engine, and dozens of wrappers). GGML library parses GGUF weights and builds the compute graph. Supports Apple Silicon, x86, RISC-V, NVIDIA/AMD GPUs.

**2026 developments:**
- **Versioning:** builds b7931 → b10369 (Aug 2026); major version bump to **llama.cpp 0.2.0** with enforced semantic versioning (Aug 2026).
- **Backends:** CUDA, Metal (consolidated binary kernels, Feb 2026), **Vulkan** (full cross-vendor GPU path), **ROCm 6.4.4** (AMD — Qwen3-235B-A22B hit 1,132 tok/s on Radeon 8060S via ROCm, 3.47× over ROCm 7.0.1 in one reported test), SYCL, OpenCL, WebGPU, Hexagon.
- **Quantization:** 1.5-bit to 8-bit; **GGUF** is the standard checkpoint format; **IQ2/IQ3** quants with importance-matrix (`--imatrix`) calibration for quality retention at extreme compression.
- **MCP integration (finalized early 2026, PR #19546):** llama.cpp became an MCP host — local agents with tool use can run fully offline. Elevates it from inference engine to modular local AI platform.
- **Multimodal:** vision support via libmtmd in llama-server; `--mmproj-device` places the vision projector on a specific GPU.
- **Speculative decoding:** MTP (multi-token prediction) draft-model auto-discovery via `--models-dir`; DSpark speculative decoder accepts SpecForge-exported drafts.
- **Native pocket-TTS support (b10369, Aug 2026):** SEANet decoder rewrite (matmul + col2im), −80% per-frame generation time on CUDA, output correlation 0.999994 vs prior — sample-identical speedup.
- **Architecture support cadence:** rapid 2026 additions — Kimi-K3 (hybrid KDA linear + MLA attention, 1024-expert latent MoE), ByteDance BailingMoE3, IBM Granite SWA/MoE-SWA, Qwen3.5 dense + MoE, Step3.5-Flash.

**Key flags for VRAM control:** `--n-gpu-layers` (fine-grained GPU offload), `--no-mmap`, `--flash-attn`.

**Positioning 2026:** maximum flexibility — Apple Silicon native, embedded/edge, one-off CLI experiments, and the engine inside Ollama/LM Studio. Best single-stream latency (one benchmark: llama.cpp ROCm + MTP 39 tok/s single-stream vs vLLM 19.5). Not for multi-user production.

Sources: https://github.com/licjon/cl-llama-cpp/blob/HEAD/docs/upstream-digest.md · https://dasroot.net/posts/2026/01/running-local-llms-python-ollama-llama-cpp-transformers/ · https://aihaberleri.org/en/news/llamacpp-integrates-mcp-protocol-to-expand-local-ai-capabilities · https://dasroot.net/posts/2026/01/local-llm-deployment-ollama-llama-cpp/

## 1.3 LM Studio

**What it is:** desktop app (Mac/Windows/Linux) — the easiest GUI for discovering, downloading (built-in Hugging Face model browser), configuring, and chatting with GGUF models. One-click OpenAI-compatible local server (port 1234). Hardware auto-detection (CUDA/MPS), visible controls for context, GPU offload, sampling.

**2026 developments:**
- **LM Studio Bionic (launched July 16, 2026):** separate agentic app for Mac/Windows — an AI agent that accesses local files and autonomously does coding, document creation, research. Built-in open-model library; falls back to **LM Studio Secure Cloud** (US servers, Zero Data Retention) for heavy tasks — billed LM Studio account. Bionic added Kimi K3, then **GLM-5.3-Flash** (Aug 26, 2026) as cloud models. Ships a voice keyboard with local transcription (Mistral Voxtral).
- **MTP speculative decoding stable in v0.4.14 (May 22, 2026)** — materially faster inference on supported models; current stable **v0.4.17**.
- **MCP host since v0.3.17** (June 2025); v0.4.10 (April 2026) added OAuth 2.1 browser auth; ephemeral per-request MCPs (SSRF-guarded); fine-grained `allowed_tools` gating.
- **Locally acquisition (April 8, 2026):** extended the device story to iPhone/iPad — shipped as the Locally mobile app + **LM Link**, an end-to-end-encrypted bridge letting a phone drive large models on the desktop (v0.4.16, June 2026).
- **mlx-engine v1.8.5 (June 2026):** disk-backed KV-cache reuse + continuous batching for serious local agentic workloads.
- **Model Search by hardware:** tells you which models fit your Mac's specs.

**Positioning 2026:** "LM Studio is an app" (easiest desktop experience, HF GGUF browsing, non-technical users) vs "Ollama is a runtime" (headless, scriptable, Docker/CI). Note: Ollama now also offers optional cloud models — "using Ollama" no longer automatically means every request is local.

Sources: https://9to5mac.com/2026/07/16/lm-studio-expands-beyond-chat-with-bionic-a-new-ai-agent-app-for-open-models/ · https://9to5mac.com/2026/08/26/lm-studio-adds-glm-5-3-flash-to-bionic-with-image-support-and-1m-token-context/ · https://media.patentllm.org/news/local-ai/lm-studio-adds-mtp-speculative-decoding-qwen-3-6-gguf-quants-20260520 · https://localclaw.io/blog/ollama-vs-lm-studio-2026

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

