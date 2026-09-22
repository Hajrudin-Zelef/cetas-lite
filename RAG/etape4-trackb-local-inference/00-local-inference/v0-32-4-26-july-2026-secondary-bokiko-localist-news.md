---
id: etape4-trackb-local-inference/00-local-inference/v0-32-4-26-july-2026-secondary-bokiko-localist-news
title: "v0.32.4 — 26 July 2026 [secondary: bokiko/localist news]"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "Cohere"]
dates: ["2026-04", "2026-05", "2026-05-29", "2026-06", "2026-06-01", "2026-06-17", "2026-07", "2026-07-13"]
keywords: ["accelerator", "agent", "claude", "context window", "gguf", "gpus", "inference", "llama", "llama.cpp", "packaging", "research", "speculative decoding"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [369, 417]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 7b418e6344147331982deb8eabc5a777fc7df6d1747b6d71c45cbb11cc3a57ec
---

# v0.32.4 — 26 July 2026 [secondary: bokiko/localist news]

### v0.32.4 — 26 July 2026 [secondary: bokiko/localist news]
- Support **Laguna on Apple GPUs via the MLX engine** [secondary].
- Quantize draft-model output heads at the requested type when creating speculative-decoding drafts [secondary].
- Source: https://github.com/bokiko/localist/blob/HEAD/news/2026-07.md (references https://github.com/ollama/ollama/releases/tag/v0.32.4)

### v0.32.3 — 24 July 2026 [secondary: bokiko/localist news]
- Fixed model downloads that stall before sending data [secondary].
- Integration fixes: restored Claude Code Channels, fixed Anthropic thinking streams, Hermes Desktop respects `--force-build` [secondary].
- Source: https://github.com/bokiko/localist/blob/HEAD/news/2026-07.md (references https://github.com/ollama/ollama/releases/tag/v0.32.3)

### v0.32.0 — 13 July 2026 [secondary: boardwire-ai article of official release]
- `create`: select the **qwen3.5 parser and renderer** for Qwen3.5/Next (PR #17078) [secondary].
- `launch`: warn before old agent models (PR #17063); new agent UI in cmd (PR #17017) [secondary].
- Source: https://github.com/jonas-is-coding/boardwire-ai/blob/HEAD/articles/2026-07-13-ollama-v0-32-0-720f6826f66d.md

### v0.30.11 — June 2026 [official: GitHub release v0.30.11-rc0]
- `launch`: thinking-capability detection for opencode; **auto-install Claude Code / opencode when missing** [official].
- `launch/codex`: detect model drift when Codex App UI switches [official].
- Hardware: fix inverted iGPU/dGPU Vulkan classification on Windows hybrid graphics; add sm_86 to cuda_v13_windows preset; Jetson CC 87 for CUDA v13; use host Vulkan loader on Windows; CUDA JIT packaging fix for MLX [official].
- Engine: unify/tune speculative decoding in mlxrunner; llama.cpp version update [official].
- Docs: redesign docs landing and integrations overview; document max think level [official].
- Source: https://github.com/ollama/ollama/releases/tag/v0.30.11-rc0

### v0.30.9 — 17 June 2026 [secondary: boardwire-ai article of official release]
- **Cohere2Moe architecture support** [secondary].
- Fixed LFM2 parser/render for cases where thinking was not emitted [secondary].
- Fixed bug where `ollama launch claude` and other coding-agent/assistant use cases would only output one token [secondary].
- Ollama returns an error if a single message exceeds the current context window [secondary].
- Source: https://github.com/jonas-is-coding/boardwire-ai/blob/HEAD/articles/2026-06-17-ollama-v0-30-9-de557a3a864d.md

### v0.30.0 — June 2026 (RCs from late May 2026) [secondary: boardwire-ai articles of official RCs]
- **Architecture change**: directly supports **llama.cpp** instead of building on top of GGML; enables GGUF file-format compatibility [secondary].
- **MLX** used to accelerate model inference on Apple Silicon [secondary].
- Known issues at RC: `laguna-xs.2` not yet supported on Windows/Linux; `llama3.2-vision` not yet supported; `nomic-embed-text` now lowercases inputs per model card (behavior change vs. prior versions) [secondary].
- Install pins: `OLLAMA_VERSION=0.30.0-rc31` (Mac/Linux install.sh, Windows install.ps1) [secondary].
- Sources: https://github.com/jonas-is-coding/boardwire-ai/blob/HEAD/articles/2026-06-01-ollama-v0-30-0-977511cd989c.md ; https://github.com/jonas-is-coding/boardwire-ai/blob/HEAD/articles/2026-05-29-ollama-v0-30-0-b1fd508aac93.md

### v0.21.2-rc0 — 23 April 2026 [secondary: AppSelfHost]
- **Bundled OpenClaw web search at launch** — web search available out of the box at Ollama start, no extra config or sidecar [secondary].
- Expanded docs covering **structured outputs** and **Ollama Cloud** [secondary].
- GitHub stars at the time: **169,752** [secondary].
- Source: https://appselfhost.com/ollama-v0-21-2-rc0-released-bundled-web-search-and-structured-outputs-docs-land-in-new-rc/

### Flagged uncertainty — v0.31.x
- A third-party research doc lists Ollama "v0.31.2 / v0.31.3 (Apr 22 2026)" [secondary: geeks-accelerator/ollama-herd research], which conflicts with version ordering (v0.21.x in late April, v0.30.x in June, v0.32.0 on 13 Jul 2026). The date is likely a snapshot error; v0.31.x presumably shipped late June / early July 2026, but **this is [unverified]** — no release-note content for v0.31.x was recovered.
- Source: https://github.com/geeks-accelerator/ollama-herd/blob/HEAD/docs/research/mlx-vs-ollama-adoption-2026.md

---

