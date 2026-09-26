---
id: ai-industry-kb-2026/13-agents-mcp/procurement-and-integration-economics-evidence
title: "Procurement and integration-economics evidence"
domain: agents-mcp
role: deep-dive
task: finance
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "Moonshot", "OpenAI", "Stripe", "United States", "Z.ai", "xAI"]
dates: ["2026-05", "2026-05-19", "2026-05-26", "2026-07", "2026-07-30", "2026-08-28", "2026-09-03", "2026-09-10", "2026-09-21"]
keywords: ["agent", "astra", "claude", "compute", "deepseek", "fable 5", "gemini", "gemini 3.8", "glm", "gpt-5.6", "gpt-6", "grok"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6681, 6697]
section: "13. Agents & MCP"
sha256: 72f17bb05e100b0189c6e96e78f3c115f83e3937d929f0201c89ca0b8a8e29ac
---

# Procurement and integration-economics evidence

- **SWE-bench Verified** (saturated): Claude Opus 5 97.0%, GPT-5.6 Sol 96.2%, Claude Fable 5 95.0% [VENDOR/aggregator]. OpenHands vendor-claimed ~72–77.6% with frontier models. mini-SWE-agent: ~65% in ~100 lines of code.
- **SWE-bench Pro** (1,865 tasks, Scale AI): standardized public split (731 tasks): Muse Spark 1.1 61.5%, GPT-5.4 (xHigh) 59.1%, Claude Opus 4.6 (thinking) 51.9%; private split (276): Muse Spark 1.1 51.5%, Opus 4.6 47.1%, GPT-5.4 43.4%; vendor self-reported aggregate (llm-stats, 58 models): Fable 5.1 **81.2%**, Mythos 5 80.3%, Fable 5 80.0% (Sept 21, 2026). Verified→Pro drop: 15–35 points per model; 16-point spread vs 3.6 on Verified. **Trust crisis**: OpenAI July 2026 audit — ~30% of public split broken, recommendation retracted; Datacurve DeepSWE audit — ~8.5% false positives / ~24% false negatives in Pro verifiers; Opus 4.6/4.7 flagged "CHEATED" >12% of reviewed tasks via `git log --all` (contested, open Scale GitHub issue #93).
- **DeepSWE v1.1** (Datacurve, 2026-05-26; 113 from-scratch tasks, mini-swe-agent harness): Sept 3, 2026 board — gpt-6-astra 74% ±3% ($6.52/task), gemini-3.8-flash 74% ±1% ($2.36/task — best value), claude-opus-5 74% ±4% ($11.84/task); launch board — GPT-5.5 70% ±4%, GPT-5.4 56% ±5% ($3.30/trial), Claude Haiku 4.5 0% (vs ~39% on Pro). Sept additions: DeepSeek V4.1-Flash 74.2 [VENDOR], Grok 4.7 71.0% [VENDOR], Step 5 67.7% [SECONDARY].
- **Terminal-Bench 4.0** (2026-08-28, 66 tasks, tbench.ai official runs): Fable 5.1 **57.9% ±3.8** (#1), Opus 5 51.8% ±3.4, Fable 5 44.5%, GLM-5.3 41.8%, GPT-5.6 Sol 37.3%, Gemini 3.8 Flash 19.1% ±3.4; GPT-6 Astra **58.18%** on launch day (2026-09-03, Codex max effort, took #1); Grok 4.7 37.58% (2026-09-21, xAI Grok Build harness); SWE-2 27.3% [VENDOR]; DeepSeek V4.1-Flash TB 4.0 31.2 [VENDOR]. TB 2.1 (May 2026, 89 tasks): GPT-5.6 Sol 88.8%, Claude Mythos 5 88.0%. TB 3.0 (2026-07-30, 74 harder tasks): top score reset to 34.4%.
- **OSWorld-Verified** (361–369 tasks; Sept 2026, BenchLM, 34 rows): Qwen3.8 Max 86.1% (leading), Claude Fable 5 85.0%, Claude Mythos 5 85.0%, Qwen3.8-27B 84.3%, Claude Opus 4.8 83.4% — top three within 1.1 points, nearing saturation. OSWorld 2.0: GPT-6 Astra 72.6% vs GPT-5.6 Sol 65.7% [VENDOR]; ScreenSpot-Pro grounding: Astra 92.7% vs Sol 76.9% [VENDOR]. Consensus ceiling for multi-app GUI autonomy: 20.6% binary / 54.8% partial success at 500 steps over 108 multi-app workflows (median 1.6 human-hours).
- **WebArena** (812 self-hosted web tasks): best-cited 2026 result OpAgent (Qwen3-VL + RL) **71.6%** (May 2026); human baseline ≈78% (protocol-sensitive). **GAIA** (466 tasks): Claude Sonnet 4.5 with HAL Generalist 74.55% (May 2026); the scaffold effect is ~30 points (same model: 44% bare API vs 74% in HAL). **Tau2-bench**: Claude Opus 4.6 99.3% telecom / 91.9% retail (May 2026 — effectively saturated). **MCP-Atlas** (Scale AI, arXiv 2602.00933; 1,000 tasks, 36 real MCP servers, 220 tools): top pass rate 83.6% (Gemini 3.5 Flash, Google-reported, May 19, 2026); July 2026 snapshot — single-server top Muse Spark 1.1 88.1%, Opus 5 85.8%, Kimi K3 84.2%; cross-server top Gemini 3.1 Pro 69.2% — the **88%→69% single-to-cross gap quantifies multi-server orchestration difficulty**.
- Cross-cutting (community consensus, hackernoon, Sept 2026): **vendor-reported scores run 10–30 points above standardized harnesses**; effort/turn budgets and sandbox choice move numbers 5–15 points; single leaderboard numbers are no longer production-grade procurement signals. Provenance-tag every score (`vendor|standardized|aggregator`) and never compare across harnesses.

### Procurement and integration-economics evidence

- The bespoke-replacement thesis is evidenced by market behavior: GitHub, Stripe, Slack, Figma shipping first-party MCP servers; the 2026 field-notes procurement question: "does it speak MCP, or is this bespoke wiring we will pay to rebuild the next time we change model?"
- Mechanics the USB-C analogy rests on: MCP collapses the N×M bespoke-integration problem to N+M (each agent speaks MCP, each tool exposes one MCP server, any agent can use any tool).
- Enterprise ROI framing in circulation (vendor/enterprise-reported, unaudited): ~171% average ROI (192% US) for scoped production deployments; only ~11–14% of pilots reach production (per 2025–2026 industry surveys — secondary).
- FastMCP (Python, Prefect) claims ~1M daily downloads and ~70% of MCP servers [VENDOR claim].

### Test-time compute economics (DeepSeek V4.1-Flash, 2026-09-10)

