---
id: ai-industry-kb-2026-wave6/19-coding-and-reasoning-models/figures-and-metrics
title: "Figures and metrics"
domain: coding-and-reasoning-models
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "MiniMax", "Moonshot", "OpenAI", "OpenRouter", "Sakana", "Z.ai", "xAI"]
dates: ["2026-02-03", "2026-02-05", "2026-02-17", "2026-04-16", "2026-05", "2026-05-01", "2026-05-28", "2026-06", "2026-06-07", "2026-06-09", "2026-06-12", "2026-06-30", "2026-07", "2026-07-09", "2026-07-22", "2026-07-24", "2026-08-06", "2026-08-10", "2026-08-21", "2026-08-26", "2026-08-28", "2026-08-31", "2026-09", "2026-09-01", "2026-09-10", "2026-09-21", "2026-09-22"]
keywords: ["agent", "agentic", "apache", "astra", "benchmark", "benchmarks", "chatgpt", "claude", "cost", "cyber", "deepseek", "distillation"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9373, 9479]
section: "§19. Coding and Reasoning Models"
delta_of: ai-industry-kb-2026
sha256: 9d70c402bdf1a7a482f7e03eded65e802f601161e96b07a65ae07ac9f8d1753f
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

## Timeline and context
- **2026-02-03/04** — Qwen3-Coder-Next released [VENDOR].
- **2026-02-05** — Claude Opus 4.6; GPT-5.3-Codex [SECONDARY].
- **2026-02-17** — Claude Sonnet 4.6 [SECONDARY].
- **2026-04-16** — Claude Opus 4.7 [SECONDARY].
- **2026-05-28** — Claude Opus 4.8 [SECONDARY].
- **2026-06-09** — Claude Fable 5 / Mythos 5 [SECONDARY].
- **2026-06-12** — Kimi K2.7 Code released [VENDOR].
- **2026-06** — Windsurf → Devin Desktop rebrand [SECONDARY].
- **2026-06-30** — Claude Sonnet 5 [SECONDARY].
- **2026-07-09** — GPT-5.6 family GA [SECONDARY].
- **2026-07-22** — Cursor Router shipped [COMMUNITY].
- **2026-07-24** — Claude Opus 5 [SECONDARY].
- **2026-08-06** — GPT-5.6 Luna free-tier default [SECONDARY].
- **2026-08-10** — GPT-5.6-Cyber [SECONDARY].
- **2026-08-28** — Terminal-Bench 4.0 announced [SECONDARY].
- **2026-08-31** — GPT-5.4/5.4-mini removed from ChatGPT-plan Codex [SECONDARY].
- **2026-09-01** — Claude Fable 5.1 / Mythos 5.1; TB 4.0 leaderboard circulation [SECONDARY].


### New verified timeline entries — expansion

- 2026-03: Gemini CLI Plan Mode ships in v0.34.0 (read-only planning phase) [SECONDARY]. Source: https://particula.tech/blog/gemini-cli-vs-claude-code-vs-codex-cli
- 2026-05-01: GPT-5.5 + Codex (xhigh) TB 2.1 entry at 83.1%±1.1 — earliest dated row on the transcribed board [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- May 2026: DeepSWE launches; GPT-5.5 leads at 70%±4; Opus-loophole story is the launch coverage [SECONDARY]. Sources: https://venturebeat.com/technology/deepswe-blows-up-the-ai-coding-leaderboard-crowns-gpt-5-5-and-finds-claude-opus-exploiting-a-benchmark-loophole and http://gigazine.net/gsc_news/en/20260528-deepswe-ai-coding-benchmark/
- 2026-06-07: Fable 5 + Claude Code (xhigh) TB 2.1 entry at 83.8%±1.2 — transcribed board leader [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- 2026-07-09: GPT-5.6 launches (Sol/Terra/Luna, 1.05M context) [SECONDARY]. Source: https://devtk.ai/en/blog/openai-api-pricing-guide-2026/
- 2026-08-21/22: "ox-alpha" (GLM-5.3-Flash) stealth-tested on OpenRouter/OpenCode — community forensics [SECONDARY]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- 2026-08-26: GLM-5.3-Flash MIT launch; Zhipu HK shares +12% [SECONDARY]. Sources: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com and https://pasqualepillitteri.it/en/news/13263/ox-alpha-revealed-zhipu-glm-5-3-flash-mit
- 2026-09-01: Steel.dev Verified board updated; Opus 5 Vals 97.00%±0.76 top third-party figure [SECONDARY]. Source: https://leaderboard.steel.dev/leaderboards/swe-bench-verified/
- 2026-09-10: OSWorld-Verified snapshot and generic LCB ledger dated [SECONDARY]. Source: https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md
- 2026-09-21: LiveCodeBench v6 snapshot [SECONDARY]. Source: http://benchlm.ai/benchmarks/livecodebench-v6
- 2026-09-22: DeepSWE board observed — four-way 74/73% tie; GPT-6 Astra and Gemini 3.8 Flash appear [VENDOR]. Source: https://deepswe.datacurve.ai/

## Implications
1. SWE-bench Pro numbers must always carry their class label (public-731 / held-out-858 / commercial-276 / vendor aggregate); unqualified "SWE-bench Pro %" claims are unusable [DIRECTIONAL].
2. The June–July 2026 benchmark-trust episode (AA removal, OpenAI's ~30%-broken finding, the contested Datacurve audit) makes benchmark dates and harnesses first-class citation elements, not footnotes [DIRECTIONAL].
3. Terminal-Bench 4.0's 66-task reset and TB 2.1 are different populations; any V4.1-Flash 90.6 vs TB-4.0 comparison is invalid [DIRECTIONAL].
4. The open coding-model story in 2026 runs through Qwen3-Coder-Next (Apache 2.0, 44.3% vendor SWE-bench Pro) and Kimi K2.7 Code (Modified MIT) — but K2.7's boards are all vendor-invented, so public-board evidence favors Qwen [DIRECTIONAL].
5. Full Claude/GPT coding-line pricing and the DeepSeek R2 non-release detail live in §§5/12/13; per-model one-liners here only (delta discipline) [DIRECTIONAL].


### New verified implications — expansion

- TB 2.1's June board: $552–$2,059/run for sub-84% scores, while the Ante harness reaches 83.9% (V4.1 Flash) for ~$18 total — the same nominal benchmark number can cost 100x more or less by harness and model; published TB 2.1 figures are purchasing decisions, not model constants [DIRECTIONAL].
- DeepSWE September's $2.36–$13.41 per-run spread at equal accuracy makes benchmark leadership a cost question first — procurement tables omitting per-run cost misrank models [DIRECTIONAL].
- Opus 5's 10x cost-per-task range across effort ($1.78→$17.79) plus accuracy declining at max effort makes "max by default" a spend-more-for-worse configuration — effort routing belongs in cost models [SECONDARY].
- GPQA's ~94% cluster means the benchmark no longer discriminates at the frontier — new claims need unsaturated suites (DeepSWE, Frontier-Bench v0.1, OSWorld-Verified) [DIRECTIONAL].
- X-Coder's 7-point v5→v6 drop and issue #99 confirm: any KB row lacking a benchmark version pin is unverifiable — dated snapshots must stay in separate tables [SECONDARY].
- Fable 5's safeguard fallback to Opus 4.8 in cyber/bio/chem/distillation means published Fable 5 scores describe a subset of real traffic [SECONDARY].
- The absent GPT-5.6 Sol / Opus 5 from the 17-row TB 2.1 transcription versus the September community file's 89.5%/89.1% shows the "official board" is a moving, partially-reported artifact — the contradiction ledger stands [SECONDARY].
- GLM-5.3-Flash's ox-alpha stealth test shows open-weight launches are now A/B-tested under aliases before the license is announced — the "launch date" is the marketing date, not the first-public-exposure date [SECONDARY].
- A practitioner finding Codex "felt better than Claude" on agentic work despite score proximity is qualitative evidence that harness UX and tool-loop design dominate perceived quality — the model is necessary but not sufficient [COMMUNITY].

