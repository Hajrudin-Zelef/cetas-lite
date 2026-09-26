---
id: ai-industry-kb-2026/01-frontier-closed-model-releases/figures-and-metrics
title: "Figures and metrics"
domain: frontier-closed-model-releases
role: deep-dive
task: model-release
actors: ["Anthropic", "Glasswing", "Google", "Meta", "Moonshot", "OpenAI", "xAI"]
dates: ["2026-01", "2026-05", "2026-07", "2026-07-09"]
keywords: ["agent", "agentic", "agents", "agi", "benchmark", "claude", "compute", "cost", "fable 5", "foundry", "gemini", "gpt-5.6"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [328, 408]
section: "1. Frontier Closed-Model Releases"
sha256: 7b3f3de5909519032503909a5ec6a8ac1ad7e3a89846ef368f167c5fbe61ae2c
---

# Figures and metrics

## Figures and metrics

### Pricing table — closed frontier flagships (USD per million tokens, in / out, cached-read where disclosed)

| Model | In | Cached in | Out | Notes |
|---|---|---|---|---|
| GPT-5.6 Sol | $5 | $0.50 | $30 | Pro/Ultra compute modes, 4–16 parallel agents in Ultra |
| GPT-5.6 Terra | $2.50 | $0.25 | $15 | near-GPT-5.5 perf at half cost |
| GPT-5.6 Luna | $1 | $0.10 | $6 | cheapest major-lab model; 41.3% long-context recall |
| GPT-5.5 | $5 | — | $25–30 | GA May 2026 |
| Claude Opus 4.6 | $5 | — | $25 | fast mode $30/$150 |
| Claude Opus 4.7 | — | — | — | 41 days before 4.8 |
| Claude Opus 4.8 | $5 | — | $25 | flat vs 4.7 |
| Claude Fable 5 | $10 | — | $50 | metered usage credits from Jul 7; Fable 5.1 Sep 1 |
| Claude Mythos 5 | — | — | — | more restricted sibling; Project Glasswing re-approval Jun 26 |
| Claude Sonnet 4.6 | $3 | — | $15 | |
| Claude Sonnet 5 | $2/$10 intro → $3/$15 | — | — | intro pricing through Aug 31, 2026 |
| Claude Opus 5 | $5 | $0.50 | $25 | = Opus 4.8 price, half Fable 5 |
| Gemini 3.1 Pro | $2 [UNVERIFIED] | — | $12 [UNVERIFIED] | single secondary source (aiworkflows.tools) |
| Grok 4.5 | $2 | $0.50 | $6 | developer Jul 8, public Jul 9 |
| Grok 4.6 | $2 | $0.30→$0.50 | $6 | 2× for Fast variant |
| Grok 4.7 | $2 | — | $6 | Sep 21, 2026 |
| Muse Spark 1.2 | $1.25 | $0.15 | $4.25 | web search $2.50/1K calls |
| Grok STT 1.0 | — | — | — | billed on duration: $0.10/hr batch, $0.20/hr streaming |
| Grok TTS | — | — | — | $4.20/M characters |
| Gemini Omni Flash | — | — | — | $0.10/second of video output |
| Nano Banana 2 Lite | — | — | — | $0.034 per 1K images ($0.25/$1.50 per M tokens) |

### Benchmark table — headline scores with provenance tags

| Model | Score | Provenance |
|---|---|---|
| Claude Opus 4.6 | SWE-Bench Verified 80.8%; Terminal-Bench 2.0 65.4%; ARC AGI 2 68.8%; MRCR v2 76% @1M; Finance Agent 1606 Elo | community/secondary aggregation |
| Claude Sonnet 4.6 | SWE-Bench Verified 79.6%; OSWorld 72.5%; office Elo 1633 | community/secondary |
| GPT-5.5 | SWE-bench Verified 88.7% (by Apr 2026); AA 60.2 | aggregator / VENDOR-adjacent |
| Claude Opus 4.7 | SWE-bench Verified 87.6% (by Apr 2026); DeepSWE 54% | aggregator / standardized |
| Claude Opus 4.8 | AA Intelligence Index 61.4; SWE-bench Pro 69.2% | standardized / aggregator |
| Claude Fable 5 | SWE-bench Pro 80.3%; FrontierCode 29.3% | secondary — no primary citation located [UNVERIFIED] |
| Claude Sonnet 5 | SWE-bench Pro 63.2% | secondary |
| Claude Opus 5 | SWE-bench Verified 96.0%; SWE-bench Multimodal 59.4%; Frontier-Bench v0.1 43.3–44.3; HLE 56.3; TB 4.0 51.8% | VENDOR |
| Claude Fable 5.1 | TB 4.0 57.9% ±3.8 (#1); CursorBench 3.2 Max 73.4% | official runs (tbench.ai) / Cursor |
| Claude Mythos 5.1 | Frontier-Bench (Anthropic internal) 60.9%; TB 4.0 (Anthropic-run) 55.8% | VENDOR |
| GPT-5.6 Sol | TB-2.1 91.9% (Ultra mode); SWE-bench Verified 96.2%; TB 4.0 37.3% | VENDOR / aggregator / official |
| Gemini 3.1 Pro | ARC AGI 2 77.1% (agentic); GPQA Diamond 94.3% [UNVERIFIED] | vendor-announced / single secondary |
| Grok 4.5 | #1 agentic tool-use (AA Index) | vendor-announced |
| Grok 4.6 | AA Index 61; CursorBench v3.2 69.9% ($2.81/task); FrontierCode v1.1 Ext 61.3%; APEX-Agents 57.5%; DeepSWE 65.9–73%; TB v3.0 26–34.6% | VENDOR / harness-variant splits |
| Grok 4.7 | DeepSWE v1.1 71.0%; TB 4.0 37.58–38.0% | VENDOR (press) / harness-variant split |
| Muse Spark 1.2 | AA Index 57 (v4.1.1) vs 54 (independent eval) | methodology difference — keep both |
| Muse Spark 1.3 | AA Index 61 | Wave 2.1 record |
| Kimi K3 | surpassed Opus 4.8 on AA (July 2026) | Wave 2.1 delta record |
| Grok Imagine Image 2.0 | T2I 1,320 Elo (#2); Image Edit 1,439 (#2) — behind GPT-Image-2 (1,380 / 1,463) | vendor/community arena |
| Grok STT | phone-call entity benchmark 5.0% WER vs 12.0% ElevenLabs / 13.5% Deepgram / 21.3% AssemblyAI | VENDOR |

### Context windows and knowledge cutoffs

| Model | Context | Cutoff | Notes |
|---|---|---|---|
| Claude Opus 4.6 / Sonnet 4.6 | 1M (beta) | Opus 4.6 undocumented; Sonnet 4.6 Aug 2025 | 128K max output (Opus 4.6) |
| Claude Opus 4.8 | 1M default API (200K Foundry) | undocumented in consulted sources | |
| Claude Sonnet 5 | 1M | January 2026 | 128K out (300K batch extended-output beta) |
| Claude Opus 5 | 1M in / 128K out | undocumented in consulted sources | |
| GPT-5.6 Sol / Terra / Luna | ~1.05M | undocumented in consulted sources | 128K max output, vision |
| GPT-5.5 | undisclosed (~400K working estimate) | undocumented | [ESTIMATE] — OpenAI never disclosed |
| Gemini 3.1 Pro | 1M launch ([UNVERIFIED] 2M later) | undocumented | 64–66K output |
| Grok 4.2 Beta | 256K | undocumented | |
| Grok 4.5 / 4.6 | 500K | 4.6: Feb 1, 2026 | text+image in, text out |
| Muse Spark 1.2 | 1,048,576 (~1M) | undocumented in consulted sources | mandatory reasoning |
| Gemini Omni Flash | 10-s preview generations | — | video generation context |

### The "most competitive day" scorecard — 2026-07-09

| Lab | Release | Same-day evidence |
|---|---|---|
| OpenAI | GPT-5.6 Sol / Terra / Luna GA | pricing tiers consolidated Jul 10–12 (pureai.com, ailearningguides.com, innvobyte.com, bitrouter PR #674) |
| xAI | Grok 4.5 public GA | developer release Jul 8 → public Jul 9 (Musk X post chain) |
| Anthropic | Fable 5 available (restored Jul 1) | all three labs' flagships simultaneously available |
| Meta | Muse Spark 1.1 (first paid API) | Wave 3.1 lineage record |
| Press framing | "one of the most competitive single days in the sector's history" (pureai.com, Jul 10); "arguably the busiest single day in the history of AI models" (ailearningguides.com) | [COMMUNITY] editorial framing |

## Main actors

