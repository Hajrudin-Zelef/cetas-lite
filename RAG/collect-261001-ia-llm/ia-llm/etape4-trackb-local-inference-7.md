---
id: collect-261001-ia-llm/ia-llm/etape4-trackb-local-inference-7
title: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Anthropic", "Apple", "DeepSeek", "Meta", "Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-01", "2026-07", "2026-08", "2026-08-10", "2026-08-12", "2026-08-13", "2026-08-18", "2026-09"]
keywords: ["inference", "llama", "llama.cpp", "agent", "agentic", "amd", "attention", "benchmark", "benchmarks", "chatgpt", "claude", "compute"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [290, 364]
sha256: 82d93d56d20a7ceba875995870a4ba8e18d5b91370d716f32213961f6e416431
---

# Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)

- **Founders:** Jeff Morgan (CEO) and Michael Chiang. Previously built Kitematic, acquired by Docker; then built Docker Desktop. Ollama launched 2023 [independent: TechCrunch, 9 Jul 2026]. URL: https://techcrunch.com/2026/07/09/popular-open-source-ai-developer-tool-ollama-raises-65m-grows-to-nearly-9m-users/
- **Series B:** $65M, announced **9 July 2026**, led by Theory Ventures [independent: TechCrunch].
- **Total raised:** $88M. Prior round: $15M Series A led by Benchmark's Peter Fenton (Fenton sits on the board) [independent: TechCrunch]. Morgan and Fenton declined to discuss revenue or valuation [independent].
- **Team size:** 14 employees (stated Jul 2026) [vendor-reported via TechCrunch].
- **Business thesis (2026):** the proving point for monetization came around January 2026 when OpenClaw became popular — "larger open models suddenly became able to do these agentic tasks, like coding" — pushing deep-pocketed enterprises and AI app-layer startups toward open models for daily workloads while reserving closed models for as-needed use [vendor-reported: Morgan via TechCrunch]. Benchmark's Fenton frames every high-inference-spend company as having a "vital existential project" to move to open-weight models [independent].
- **Cloud model:** Ollama hosts larger models on its "neocloud" via subscription tiers (free to $100/month at Jul 2026); usage then tracked by GPU time, not tokens [independent: TechCrunch, 9 Jul 2026]. *Note: this billing model was superseded by per-token billing by Sep 2026 — see §6.*
- **Founder position on cloud vs. open source:** Morgan calls cloud an evolution of the mission ("too big to run on your own computer. So we said, 'Hey, let's help find the compute for that'"); Fenton: "Nothing has changed for the core product that's free on the desktop" [independent: TechCrunch].
- **Community tension:** ~a year before Jul 2026, blog/social posts complained the cloud business was drawing attention from the free project, citing Ollama as an example of the "enshittification" of dev tools [independent: TechCrunch, referencing 2025 HN/Reddit/blog posts].
- **Competitive funding context (Jul 2026):** open-source inference providers also raising — Inferact (maker of vLLM) and RadixArk (maker of SGLang) named alongside Ollama, OpenClaw/NanoClaw, and Arcee [independent: TechCrunch].

---

## 3. Release timeline 2026 (latest first)

> Ollama ships "multiple/month" [secondary]. Version numbers below are all from release notes; RC tags noted where the source only confirmed an RC.

### v0.34.2 — 15 September 2026 [secondary: dev.to summary of official changelog]
- First-run setup when running `ollama`: option to sign in or continue locally; state shared with the desktop app on macOS and Windows [secondary].
- `ollama://apps` deep link opens the Apps page of the desktop app directly (macOS, Windows) [secondary].
- Fix: "Fixed excessive memory growth during long generations with MLX speculative decoding" [secondary].
- llama.cpp update (no version detail in changelog) [secondary].
- Source: http://dev.to/jtorchia/ollama-v0342-vs-v0341-a-targeted-mlx-fix-not-the-same-regression-mkh (summarizes https://github.com/ollama/ollama/releases/tag/v0.34.2)

### v0.34.1 — 14 September 2026 [official via newreleases.io mirror of GitHub release]
- MLX safetensors `ollama create` is **no longer experimental**; creating GGUF models now requires llama.cpp tooling for safetensor conversion and quantization [official].
- "Improved MLX memory handling on Apple Silicon" [official].
- Runaway repeat-token detection now requires **100 repeat tokens** before triggering (fewer false positives, e.g. OCR) [official].
- `/api/tags` much faster on large model libraries: **3.1 s → 294 ms cold** in testing; model capabilities reported consistently [official].
- Deprecated `typical_p`: can no longer be set on new models; existing GGUF models retain support [official].
- MLX and llama.cpp updates [official].
- Source: https://newreleases.io/project/github/ollama/ollama/release/v0.34.1 (mirror of https://github.com/ollama/ollama/releases/tag/v0.34.1)

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

