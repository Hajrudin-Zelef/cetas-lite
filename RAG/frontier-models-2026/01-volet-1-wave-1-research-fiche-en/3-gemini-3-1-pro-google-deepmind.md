---
id: frontier-models-2026/01-volet-1-wave-1-research-fiche-en/3-gemini-3-1-pro-google-deepmind
title: "3. Gemini 3.1 Pro (Google DeepMind)"
domain: volet-1-wave-1-research-fiche-en
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Google", "Microsoft", "OpenAI"]
dates: ["2025-01", "2026-02", "2026-02-19", "2026-03-26", "2026-05-04"]
keywords: ["gemini", "agentic", "agents", "agi", "benchmark", "benchmarks", "claude", "compute", "consumer", "copilot", "cost", "multimodal"]
source: docs/RAG/Grands titres IA modèlesEN.md
source_anchor: ""
source_lines: [96, 146]
section: "VOLET 1 — Wave 1 Research Fiche (EN)"
sha256: 6c84a439106d0a0ee437abb8912fc2c98c1a445f8d74a1c6c7159182f98278ad
---

# 3. Gemini 3.1 Pro (Google DeepMind)

## 3. Gemini 3.1 Pro (Google DeepMind)

### Release & announcement
- **Released February 19, 2026 — in Preview** (`gemini-3.1-pro-preview`), ~3 months after Gemini 3 Pro (Nov 2025) — a fast turnaround by frontier standards.
- Preview → staged rollout: developer launch in Feb; consumer rollout accelerated May 4, 2026 in the Gemini app (expanded limits for AI Pro/Ultra subscribers) and NotebookLM for paid users. Gemini 3 Pro was discontinued on Vertex AI on March 26, 2026, pushing enterprise customers onto the 3.1 track.

### Architecture & technical specs
- **Architecture:** Mixture-of-Experts multimodal reasoning model with extended chain-of-thought ("thinking" / reasoning levels: Minimal, Low, Medium, High — configurable reasoning compute per request).
- **Context:** 1,048,576 input tokens (~1M; ≈1,500 A4 pages) → suitable for whole-codebase/long-context RAG. **Output:** 65,536 tokens per response.
- **Multimodality (native):** text, images, video, audio, PDF, code repositories as first-class inputs (e.g., ~45 min video with audio, 10 videos per prompt); text output only (no image/audio generation in preview). Live API not supported.
- Knowledge cutoff: January 2025 (API docs).
- Output speed: ~104–116 tokens/sec (well above peer median of ~72 t/s).
- **Dedicated agentic endpoint:** `gemini-3.1-pro-preview-customtools` — optimized for custom tool calls (view_file, search_code, bash) and multi-step agentic pipelines.

### Key benchmarks (reported values, Feb 2026)
| Benchmark | Score |
|---|---|
| **ARC-AGI-2** | **77.1%** — more than double Gemini 3 Pro (31.1%) |
| **SWE-Bench Verified** | **80.6%** (Google's figure; ~75% on third-party BenchLM standardization) |
| GPQA Diamond | 94.3% |
| MMMLU | 92.6% |
| MMMU / MMMU-Pro | ~62–64% / ~80–84% (independent trackers) |
| APEX-Agents | 33.5% (vs Claude Opus 4.6's 29.8%) |
- Headline claim: **1st place on 13 of 16** evaluated benchmarks (Google).

### API pricing & availability
- **$2.00 / 1M input, $12.00 / 1M output** for prompts ≤200K tokens; **$4.00 / $18.00** above 200K — **unchanged from Gemini 3 Pro**, making it a "free performance upgrade." Context caching supported (up to ~75% cost reduction); cached input ~$0.20/1M.
- Consumer: bundled in Google AI Pro and AI Ultra subscriptions (regional pricing varies).
- Channels: Gemini app, AI Studio, Gemini API, Vertex AI, Gemini CLI, Google Antigravity platform, Android Studio, NotebookLM, GitHub Copilot / VS Code extensions.

### Agentic / coding use cases
- **Creative & agentic coding:** code-based SVG animations, 3D visualizations, live aerospace dashboards, interactive design prototypes from text prompts alone.
- **Long-context RAG workflows:** entire codebases, contracts, research corpora in a single pass.
- Enterprise: Batch API, context caching, function calling, code execution, search grounding, structured outputs, URL context.

### What distinguished it in 2026
- **The reasoning-leap release at flat pricing**: generational jump (ARC-AGI-2 doubled, SWE-Bench 80.6%) with **zero price increase** — and the multimodal/reasoning flagship anchoring Google's entire 2026 ecosystem (Antigravity, AI Studio, NotebookLM, Android Studio), with a purpose-built agentic tool-calling endpoint.

---

## Comparison summary (February 2026 wave)

| Dimension | Claude Sonnet 5 "Fennec" | GPT-5.3-Codex | Gemini 3.1 Pro |
|---|---|---|---|
| **Release** | Feb 3, 2026 | Feb 5, 2026 | Feb 19, 2026 (Preview) |
| **Context** | 1M tokens | 400K tokens | 1M tokens |
| **SWE-Bench Verified** | **82.1%** | 56.8% (SWE-Bench Pro) | 80.6% |
| **Terminal-Bench 2.0** | 80.4% (2.1) | **77.3%** | — |
| **Headline reasoning** | Dev Team orchestration | Self-recursive training | ARC-AGI-2 **77.1%** |
| **Input / Output $ per 1M** | $3 / $15 | $1.75 / $14 | $2 / $12 |
| **Standout** | Agentic autonomy, best price/perf | Terminal workflows + LTS stability | Multimodal reasoning at flat pricing |
