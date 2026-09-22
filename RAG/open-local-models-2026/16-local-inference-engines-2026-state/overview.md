---
id: open-local-models-2026/16-local-inference-engines-2026-state/overview
title: "1. LOCAL INFERENCE ENGINES — 2026 STATE"
domain: local-inference-engines-2026-state
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Anthropic", "Apple", "ByteDance", "Hugging Face", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "United States", "Z.ai", "vLLM"]
dates: ["2025-06", "2026-03", "2026-03-30", "2026-04", "2026-04-08", "2026-05", "2026-05-22", "2026-06", "2026-07-16"]
keywords: ["inference", "inference engine", "acquisition", "agent", "agentic", "agents", "amd", "attention", "benchmark", "claude", "compute", "copilot"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [779, 833]
section: "1. LOCAL INFERENCE ENGINES — 2026 STATE"
sha256: 2c328785175e636058fe4994477d4fc993309863ac19db03c795e0e7c1822a20
---

# 1. LOCAL INFERENCE ENGINES — 2026 STATE

## 1.1 Ollama

**What it is:** MIT-licensed local LLM daemon (REST API on port 11434, model management, minimal terminal chat). The default local backend for the ecosystem — powers Open WebUI, Continue.dev, AnythingLLM, Cline/Roo Code local setups, and most IDE/agent integrations. Model library at ollama.com/library has 100+ models; `ollama pull` downloads the recommended quantization for your hardware automatically (no manual Q4_K_M vs Q5_K_S choice).

**2026 developments:**
- **v0.19 (preview, March 30, 2026): MLX backend for Apple Silicon.** Rebuilt the Mac inference stack on Apple's MLX framework. Official Ollama benchmark on M5 Max running Qwen3.5–35B-A3B (NVFP4): prefill 1,154 → 1,810 tok/s (+57%), decode 58 → 112 tok/s (+93%); int4 pushes to 1,851 prefill / 134 decode tok/s. Community reports 2–3× faster decode vs the old Metal-via-llama.cpp path, TTFT roughly halved. llama.cpp remains the backend on Linux/Windows.
- **v0.24.0 (May 2026):** latest stable cited in 2026 guides; MLX backend now standard on Apple Silicon.
- **OpenAI-compatible `/v1/chat/completions` endpoint** — any OpenAI-API-speaking tool works unmodified. Also added **Anthropic Messages API compatibility**, so Claude-format tools work too.
- **GitHub Copilot integration (March 2026):** every code suggestion, chat prompt, and agentic workflow can route to local models on localhost:11434 — no API keys, no telemetry, no per-token charges. Structural shift for regulated industries (code never leaves the network perimeter).
- Time to first inference: ~5 min macOS/Linux, ~8 min Windows.

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

