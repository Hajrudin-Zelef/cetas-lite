---
id: etape4-trackb-local-inference/00-local-inference/v0-34-0-8-september-2026-official-github-release
title: "v0.34.0 — 8 September 2026 [official: GitHub release]"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["AMD", "Anthropic", "Apple", "DeepSeek", "Meta", "Nvidia", "OpenAI"]
dates: ["2026-07", "2026-08", "2026-08-10", "2026-08-12", "2026-08-13", "2026-08-18", "2026-09"]
keywords: ["agent", "agentic", "amd", "benchmarks", "chatgpt", "claude", "decode", "deepseek", "gguf", "llama", "llama.cpp", "multimodal"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [322, 368]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: a5162b26f0b36b0b6e5d9fc1563de389b661873b9ca89de00aed6bedabe6773d
---

# v0.34.0 — 8 September 2026 [official: GitHub release]

### v0.34.0 — 8 September 2026 [official: GitHub release]
- **Ollama models can now be used directly in ChatGPT Desktop** — "keep your existing workflow while running open models"; setup available from the Ollama app on macOS [official].
- Improved structured-output performance on Apple Silicon [official].
- Support for OpenAI-compatible client **tool search** and **response compaction** [official].
- Source: https://github.com/ollama/ollama/releases/tag/v0.34.0

### v0.33.3 — early September 2026 [official: GitHub release]
- `gemma4` now supports **images and audio on the MLX engine** [official].
- Report cached prompt tokens [official].
- Honor GGUF model-defined default parameters [official].
- MLX, MLX-C, llama.cpp updates [official].
- Source: https://github.com/ollama/ollama/releases/tag/v0.33.3

### v0.33.0 — 26 August 2026 [secondary: bokiko/localist news]
- **Claude Desktop**: developers can configure Claude Desktop to work with Ollama as a third-party gateway provider [secondary].
- Source: https://github.com/bokiko/localist/blob/HEAD/news/2026-08.md (references https://github.com/ollama/ollama/releases/tag/v0.33.0)

### v0.32.15 — 21 August 2026 [secondary: bokiko/localist news]
- New desktop onboarding flow on first launch [secondary].
- Caches resolved model metadata between requests; time-to-first-token roughly halved (TTFT ~995 ms → ~524 ms in benchmarks) [secondary].
- Source: https://github.com/bokiko/localist/blob/HEAD/news/2026-08.md (references https://github.com/ollama/ollama/releases/tag/v0.32.15)

### v0.32.11 — 18 August 2026 [secondary: boardwire-ai article of official release]
- `ollama launch dsh` supports **DeepSeek Harness** (DeepSeek's open-source agent harness) [secondary].
- `ollama launch muse` supports **Muse Code** (Meta's agentic coding CLI) [secondary].
- The **OpenAI-compatible Responses API now supports web search** [secondary].
- Muse Glimmer template updates [secondary].
- Source: https://github.com/jonas-is-coding/boardwire-ai/blob/HEAD/articles/2026-08-18-ollama-v0-32-11-b60cd68a161b.md (source: https://github.com/ollama/ollama/releases/tag/v0.32.11)

### v0.32.10-rc0 — 13 August 2026 [secondary: boardwire-ai article of official RC]
- Kernel-fusion optimization for ModelOpt checkpoints applying float32 global scale: on an M5 Max, `qwen3.6:27b` prefill **703 → 769 t/s (+7.9%)**, `muse-glimmer:30b` prefill **790 → 843 t/s (+6.7%)**; greedy outputs byte-identical; speculative decode unchanged within noise [secondary].
- Source: https://github.com/jonas-is-coding/boardwire-ai/blob/HEAD/articles/2026-08-13-ollama-v0-32-10-rc0-a8ebf902d5a6.md

### v0.32.8 — 12 August 2026 [secondary: boardwire-ai article of official release]
- **Muse Glimmer on all platforms**: NVIDIA, AMD and additional platforms (previously Apple Silicon only) [secondary].
- Source: https://github.com/jonas-is-coding/boardwire-ai/blob/HEAD/articles/2026-08-12-ollama-v0-32-8-05c0630e94fc.md

### v0.32.7 — 10 August 2026 [secondary: boardwire-ai article of official release]
- **Muse Glimmer**: Meta's newest open model, the first released by **Meta Superintelligence Labs**; 30B multimodal model purpose-built for agent workloads that run locally [secondary].
- Initial support via Ollama's **MLX engine on Apple Silicon** with **DFlash** speculative decoding and image input [secondary].
- Usage: `ollama run muse-glimmer` / `ollama run muse-glimmer:30b-mlx`; powers Claude Code, Codex, Pi, OpenClaw, Hermes via `ollama launch <agent> --model muse-glimmer` [secondary].
- Source: https://github.com/jonas-is-coding/boardwire-ai/blob/HEAD/articles/2026-08-10-ollama-v0-32-7-21fdc04ea90a.md

### v0.32.5 — 28 July 2026 [secondary: bokiko/localist news]
- Fixed an MLX Metal bug that could reduce output quality for **NVFP4** models, particularly Laguna [secondary].
- Source: https://github.com/bokiko/localist/blob/HEAD/news/2026-07.md (references https://github.com/ollama/ollama/releases/tag/v0.32.5)

