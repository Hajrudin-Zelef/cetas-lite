---
id: etape4-trackb-local-inference/00-local-inference/1-executive-summary
title: "1. Executive summary"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["Anthropic", "Apple", "Meta", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-01", "2026-07", "2026-09"]
keywords: ["agent", "agentic", "attention", "benchmark", "chatgpt", "claude", "compute", "funding", "gguf", "gpu", "inference", "llama"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [277, 321]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 3d61cddfe09a74e9a4ed49d544ebfe5427f249f0c3e9c2d140f82ea2737e6282
---

# 1. Executive summary

## 1. Executive summary

- Ollama remains the dominant local-LLM runner in 2026: #1 most-starred conversational-AI open-source project (178,535 stars, Aug 2026 [secondary]), ~8.9M monthly developers, 85% of Fortune 500 [vendor-reported via TechCrunch, 9 Jul 2026], ~1M installs/week [vendor-reported].
- Funding: **$65M Series B led by Theory Ventures announced 9 July 2026**; $88M raised total (prior $15M Series A, Benchmark) [independent].
- 2026 release cadence: v0.21 (Apr) → v0.30 (Jun, engine rebuild on llama.cpp + MLX) → v0.32 (Jul–Aug, Muse Glimmer, coding-agent "launch" integrations) → v0.33 (late Aug) → **v0.34.2 (15 Sep 2026, latest at cut-off)** [official/secondary].
- Strategic pivot of 2026: from pure local runner to **local + cloud hybrid** — Ollama Cloud (`:cloud` models, per-token pricing, off-peak discounts), **coding-agent integrations** (`ollama launch claude/codex/pi/openclaw/hermes/dsh/muse`), and first-party **desktop apps** wired into ChatGPT Desktop and Claude Desktop.
- Pricing shifted during 2026 from GPU-time billing (still reported Jul 2026) to **transparent per-token pricing with usage credits**: Free $0 / Pro $20/mo ($60 usage) / Max $100/mo ($300 usage) / Team $500/mo ($1,000 shared) [secondary, Sep 2026]; Team pricing data conflicts with a seat-based figure (see §7, flagged).
- Competition: LM Studio (GUI/desktop lane, proprietary, paid enterprise) and vLLM (server-throughput lane, up to 2.6x throughput claims) are the two named challengers; 2026 coverage consistently ranks Ollama #1 for developer API use but notes weak concurrency and no built-in auth.

---

## 2. Company, team and funding

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

