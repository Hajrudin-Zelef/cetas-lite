---
id: ai-industry-kb-2026-wave6/19-coding-and-reasoning-models/figures-and-metrics
title: "Figures and metrics"
domain: coding-and-reasoning-models
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "MiniMax", "Moonshot", "OpenAI", "Sakana", "Z.ai", "xAI"]
dates: ["2026-05", "2026-05-01", "2026-06", "2026-06-07", "2026-09", "2026-09-01", "2026-09-10", "2026-09-21"]
keywords: ["agent", "agentic", "apache", "astra", "benchmark", "benchmarks", "chatgpt", "claude", "cost", "cyber", "deepseek", "fable 5"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9373, 9425]
section: "§19. Coding and Reasoning Models"
delta_of: ai-industry-kb-2026
sha256: 936f06766722f273f86d80c0b149014e121e2c317509da1535144bab46aac00d
---

# Figures and metrics

## Figures and metrics
| Model | Key coding figure | Class |
|---|---|---|
| Qwen3-Coder-Next | SWE-bench Pro 44.3% [VENDOR] | vendor-reported |
| DeepSeek V4.1-Flash | TB 2.1 90.6; DeepSWE v1.1 74.2 [SECONDARY] | standardized boards |
| Opus 4.8 | SWE-bench Pro 69.2% [VENDOR] | vendor aggregate |
| GPT-5.3-Codex | SWE-bench Pro public 56.8% [VENDOR] | vendor (public-class label) |
| Kimi K2.7 Code | Kimi Code Bench v2 62.0 [VENDOR] | vendor-invented |
| SWE-1.5 | SWE-bench Verified 40.08% [VENDOR] | vendor |
| Cursor Composer 2 | SWE-bench Multilingual 73.7% [VENDOR] | vendor (Multilingual ≠ Verified) |
| Fable 5.1 | TB 4.0 57.9%±3.8 [SECONDARY] | standardized |
| Opus 5 | TB 4.0 51.8 / 53.9 / 52.3 per source [SECONDARY] | standardized, snapshot variance |

SWE-bench Pro classes — never mix [DIRECTIONAL]:
| Class | Size | Example top scores |
|---|---|---|
| Scale standardized public | 731 tasks | GPT-5.4 xHigh 59.1%, Opus 4.6 51.9% [SECONDARY] |
| Held-out | 858 tasks | — [SECONDARY] |
| Scale commercial/private | 276 tasks | Opus 4.6 47.1%, GPT-5.4 43.4% [SECONDARY] |
| Vendor aggregate | separate class | Opus 4.8 69.2%, GLM-5.2 62.1%, MiniMax M3 59.0% [VENDOR] |


### New verified metrics — expansion

- DeepSWE September 2026 board with per-run cost and output tokens [VENDOR — official board]: GPT-6 Astra (xhigh) 74%±3, $4.43, 30k tokens, 29 steps | Gemini 3.8 Flash (high) 74%±1, $2.36, 143k, 166 | Opus 5 (max) 74%±4, $11.84, 118k, 99 | GPT-5.6 Sol (max) 73%±3, $6.46, 60k, 61 | Fable 5 (xhigh) 70%±3, $13.41, 80k, 68 | GLM-5.3 (max) 69%±3, $3.99, 80k, 124 | Kimi K3 (max) 69%±5, $4.65, 81k, 98. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026 lower board [VENDOR]: Grok 4.6 (medium) 67%±2 | GPT-5.6 Luna (max) 67%±4 | GPT-5.5 (xhigh) 67%±6 | Gemini 3.7 Flash 65%±3 | GLM-5.3-Flash 63%±4 | DeepSeek V4 Pro 63%±6 | Opus 4.8 59%±2 | Qwen3.8-Max 57%±3 | Sonnet 5 54%±4 | DeepSeek V4 Flash 53%±4 | Gemini 3.6 Flash 47%±4 | GLM-5.2 44%±2 | Gemini 3.5 Flash 36%±4. Source: https://deepswe.datacurve.ai/
- DeepSWE May 2026 original board (dated; do not merge) [SECONDARY]: GPT-5.5 70%±4 | GPT-5.4 56%±5 | Opus 4.7 54%±5 | Sonnet 4.6 32%±4 | Gemini 3.5 Flash 28%±4 | GPT-5.4-mini and Kimi K2.6 24%±4. Source: https://venturebeat.com/technology/deepswe-blows-up-the-ai-coding-leaderboard-crowns-gpt-5-5-and-finds-claude-opus-exploiting-a-benchmark-loophole
- Terminal-Bench 2.1 official board, 17 rows [SECONDARY]: Fable 5 + Claude Code (xhigh) 83.8%±1.2, 2026-06-07, $552.67 | GPT-5.5 + Codex (xhigh) 83.1%±1.1, 2026-05-01, $2,059.19 | Fable 5 + Terminus 2 (high) 80.4%±1.2 | Grok 4.5 + Cursor CLI (high) 79.3%±1.5 | Opus 4.8 + Claude Code (high) 78.9%±1.3 | GPT-5.6 Terra + Codex (max) 78.4%±1.3, $421.15 | GPT-5.6 Luna + Codex (max) 75.7%±1.3, $241.45 | Sonnet 5 + Claude Code (high) 74.6%±1.6 | Gemini 3 Pro + Terminus 2 (high) 73.9%±1.3. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- Terminal-Bench 2.1, Ante same-harness runs [COMMUNITY]: DeepSeek V4.1 Flash 83.9% (370/445, ~$18 total) | GLM-5.2 74.6%±2.06 | MiMo V2.5 65.8%±2.30 | DeepSeek V4 Pro 65.8%±2.25 | DeepSeek V4 Flash 62.7%±2.29 | MiniMax M3 62.1%±2.33. Source: https://github.com/antigmalabs/ante/blob/HEAD/docs-site/docs/benchmarks/eval.mdx
- SWE-bench Verified, Steel.dev archived 2026-09-01 [SECONDARY]: Opus 5 (Vals) 97.00%±0.76 | Mythos 5 95.5% | Fable 5 95.0% | Opus 4.8 88.6% | Opus 4.7 87.6% | GPT-5.6 Sol 82.2% (Inkling) | Opus 4.6 80.8%. Vendor point: Anthropic's own Opus 5 was 96.0% (Jul 2026). Source: https://leaderboard.steel.dev/leaderboards/swe-bench-verified/
- LiveCodeBench v6 2026-09-21 (pinned) [SECONDARY]: Fugu-Ultra 93.2 | Fugu 92.9 | Qwen3.8-Omni-Flash 92.6. Source: http://benchlm.ai/benchmarks/livecodebench-v6
- LiveCodeBench generic 2026-09-10 (unpinned — separate) [SECONDARY]: Qwen3.7 Max 91.6 | Qwen3.7 Plus 89.6 | Solar Pro 4 87.8 | GLM-4.7 84.9 | Qwen3.6-27B 83.9 | Qwen3.6-35B-A3B 80.4 | Mercury 2 67.3 | DeepSeek V3 37.6. Source: https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md
- OSWorld-Verified 2026-09-10 [SECONDARY]: Qwen3.8 Max 86.1 | Fable 5 85.0 | Mythos 5 85.0 | Qwen3.8-27B 84.3 | Opus 4.8 83.4 | Gemini 3.6 Flash 83.0 | Holo3-35B-A3B 82.6 | Sonnet 5 81.2 | Muse Spark 1.1 80.8 | Holo3-122B-A10B 78.8. Source: https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md
- GPQA Diamond saturation cluster [SECONDARY]: GPT-5.5 ~94.0 | Gemini 3.1 Pro ~94.1–94.3 | Opus 4.8 ~93.6 | GPT-5.2 Pro 93.2 | Gemini 3 Pro 91.9 | Opus 4.6 91.3 | Sonnet 4.6 89.9 | Opus 4.5 87.0. Sources: https://tech-insider.org/claude-vs-chatgpt-vs-gemini-2026/ and https://github.com/rolandtolnay/mindsystem/blob/HEAD/references/models/claude-sonnet-4-6.md
- GLM-5.3-Flash vs Opus 4.8 vendor head-to-head (Z.ai, Aug 2026) [VENDOR]: GDPVal-AA v2 1773 vs 1582 | DeepSWE v1.1 63.4 vs 58.0 | AutomationBench 48.8 vs 41.0 | TB 2.1 84.3 vs 85.0 | HLE w/ Tools 55.3 vs 57.9 | Z.ai Code Bench v1.0 (max) 29.0 vs 29.5. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- Gemini 3.7 vs 3.6 Flash deltas (three weeks) [SECONDARY]: FrontierCode 1.1 43.6% vs 34.4% | DeepSWE v1.1 65.3% vs 49.0% | GDP.pdf 34.0% vs 22.0% | AutomationBench 30.4% vs 17.0%. Source: https://toknow.ai/posts/gemini-3-7-flash-half-price-coding-workhorse/index.pdf
- Kimi K3 vs DeepSeek V4.1 Flash economics (OrcaRouter/AA, Sept 2026) [SECONDARY]: $2.00 vs $0.27 per task (7.4x) | Vals agentic $17.59 vs $0.41 (~40x) | 34.7 vs 214.4 tok/s | cache read $0.30 vs $0.003/M. Source: https://www.orcarouter.ai/blog/deepseek-v4-1-flash-vs-kimi-k3
- AA blended 7:2:1 $/1M and cost/task [SECONDARY]: K3 $2.31/$0.94 | GLM-5.2 $0.90/$0.32 | V4 Pro $0.18/$0.04. Source: https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
- DeepSeek official peak/off-peak cards per 1M (peak = weekday 01:00–04:00, 06:00–10:00 UTC) [SECONDARY]: V4-Flash (0731) peak $0.44 miss/$0.014 hit/$1.32 out, off-peak half | V4.1-Flash off-peak $0.15/$0.60, peak $0.30/$1.20. Sources: https://kimi-k2.org/blog/56-kimi-k3-vs-deepseek-v4-flash-0731 and https://www.orcarouter.ai/blog/deepseek-v4-1-flash-vs-kimi-k3
- Opus 5 cost-per-task (AA-Briefcase) [SECONDARY]: $1.78 (low) → $17.79 (max). Source: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md

## Main actors
- **Scale AI** — SWE-bench Pro operator (public / held-out / private boards) [SECONDARY].
- **Harbor** — Terminal-Bench 4.0 official harness [SECONDARY].
- **Artificial Analysis** — removed SWE-bench Pro from its Coding Agent Index mid-June 2026 [SECONDARY].
- **Datacurve** — DeepSWE operator; audit claims contested [COMMUNITY].
- **OpenAI** — found ~30% of SWE-bench Pro tasks broken; GPT-5.3-Codex; GPT-5.6-Cyber [SECONDARY].
- **Alibaba** — Qwen3-Coder-Next (Apache 2.0) [VENDOR].
- **Moonshot** — Kimi K2.7 Code (Modified MIT) [VENDOR].
- **DeepSeek** — V4.1-Flash coding lead on TB 2.1; R2 unreleased [SECONDARY].
- **Anthropic** — Claude coding line; TB 4.0 official leader (Fable 5.1) [SECONDARY].
- **Cognition** — Windsurf → Devin Desktop rebrand; SWE-1.5/SWE-1.6 [SECONDARY].

