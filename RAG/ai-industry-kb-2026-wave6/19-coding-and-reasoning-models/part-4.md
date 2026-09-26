---
id: ai-industry-kb-2026-wave6/19-coding-and-reasoning-models/part-4
title: "§19. Coding and Reasoning Models (part 4)"
domain: coding-and-reasoning-models
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "Microsoft", "Nvidia", "OpenAI", "OpenRouter", "Sakana", "Z.ai"]
dates: ["2026-09-10", "2026-09-21"]
keywords: ["reasoning", "agents", "benchmark", "benchmarks", "chatgpt", "claude", "cost", "deepseek", "fable 5", "fugu", "gemini", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9324, 9353]
section: "§19. Coding and Reasoning Models"
delta_of: ai-industry-kb-2026
sha256: 8de2f8a2ad73093f32ada54205a8ec3e265d4836e7e538a871e7ad27b5260303
---

# §19. Coding and Reasoning Models (part 4)

- GLM-5.3-Flash DeepSWE v1.1 (vendor-reported, Z.ai): 63.4, ahead of GLM-5.2 at 46.2 and Claude Opus 4.8 at 58.0 — a 17-point generational jump within the Flash line [VENDOR]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- Z.ai head-to-head (vendor, Aug 2026): GLM-5.3-Flash leads Opus 4.8 on GDPVal-AA v2 (1773 vs 1582), DeepSWE v1.1 (63.4 vs 58.0), AutomationBench (48.8 vs 41.0) [VENDOR]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- Z.ai head-to-head: Opus 4.8 leads Flash on Terminal-Bench 2.1 (85.0 vs 84.3) and HLE with Tools (57.9 vs 55.3); Z.ai Code Bench v1.0 at max effort: Flash 29.0 vs Opus 29.5 [VENDOR]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- GLM-5.3-Flash was stealth-tested as "ox-alpha" on OpenRouter/OpenCode (community forensics Aug 21–22: tokenizer match, Z.AI error codes, Java stack trace) before the MIT launch — open weights can follow closed preview with no license continuity obligation [SECONDARY]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- GLM-5.3-Flash is "the most popular model on OpenRouter by token usage" — price drives adoption, adoption drives the benchmark sample [SECONDARY]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- GLM-5.3-Flash is described as the first 200B-class open-source reasoning model running natively on domestic (non-NVIDIA) chips — MIT plus non-NVIDIA hardware is the news [SECONDARY]. Source: https://pasqualepillitteri.it/en/news/13263/ox-alpha-revealed-zhipu-glm-5-3-flash-mit
- Zhipu's Hong Kong shares closed up more than 12% (HK$1,160) on the GLM-5.3-Flash/MIT launch news — markets price open-weight releases [SECONDARY]. Source: https://pasqualepillitteri.it/en/news/13263/ox-alpha-revealed-zhipu-glm-5-3-flash-mit
- Gemini 3.7 Flash vs 3.6 Flash (three weeks apart): FrontierCode 1.1 43.6% vs 34.4% [SECONDARY]. Source: https://toknow.ai/posts/gemini-3-7-flash-half-price-coding-workhorse/index.pdf
- Gemini 3.7 vs 3.6 Flash: DeepSWE v1.1 65.3% vs 49.0%; GDP.pdf 34.0% vs 22.0%; AutomationBench 30.4% vs 17.0% — the largest gains came in the cheap tier, not the flagship [SECONDARY]. Source: https://toknow.ai/posts/gemini-3-7-flash-half-price-coding-workhorse/index.pdf
- Gemini 3.7 Flash "targets coding and agents with a 50% introductory price cut" (VentureBeat) — the positioning is coding-first [SECONDARY]. Source: https://venturebeat.com/technology/googles-gemini-3-7-flash-targets-coding-and-agents-with-a-50-introductory-price-cut
- LiveCodeBench v6 snapshot 2026-09-21 (version-pinned): Sakana Fugu-Ultra 93.2; Sakana Fugu 92.9; Qwen3.8-Omni-Flash 92.6 [SECONDARY]. Source: http://benchlm.ai/benchmarks/livecodebench-v6
- LiveCodeBench generic ledger 2026-09-10 (NOT version-pinned — keep separate): Qwen3.7 Max 91.6; Qwen3.7 Plus 89.6; Solar Pro 4 87.8; GLM-4.7 84.9; Qwen3.6-27B 83.9; Qwen3.6-35B-A3B 80.4; Mercury 2 67.3; DeepSeek V3 37.6 [SECONDARY]. Source: https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md
- X-Coder 7B (Microsoft/Tsinghua): LiveCodeBench v5 62.9, v6 55.8 — a 7-point drop across the version boundary [SECONDARY]. Source: https://dataconomy.com/2026/01/27/microsoft-and-tsinghuas-x-coder-hits-62-9-pass-rate-on-livecodebench-v5/
- LiveCodeBench maintainers document the version-comparability problem in issue #99: scores under different task windows or eval scripts are not directly comparable even when the benchmark name matches [COMMUNITY]. Source: https://github.com/livecodebench/livecodebench/issues/99
- OSWorld-Verified snapshot 2026-09-10: Qwen3.8 Max 86.1 tops the board — an open-weight model ahead of every flagship on computer use [SECONDARY]. Source: https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md
- OSWorld-Verified: Fable 5 85.0; Mythos 5 85.0; Qwen3.8-27B 84.3; Opus 4.8 83.4; Gemini 3.6 Flash 83.0 [SECONDARY]. Source: https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md
- OSWorld-Verified: Holo3-35B-A3B 82.6; Sonnet 5 81.2; Muse Spark 1.1 80.8; Holo3-122B-A10B 78.8 [SECONDARY]. Source: https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md
- OSWorld-Verified corroboration (independent): Opus 4.8 83.4 matches exactly; Opus 4.7 82.8; GPT-5.5 78.7; Gemini 3.1 Pro 76.2 [SECONDARY]. Source: https://www.worthview.com/claude-opus-4-8-vs-gpt-5-5-vs-gemini-3-1-pro-benchmark-2026/
- GPQA Diamond is saturated: GPT-5.5 ~94.0; Gemini 3.1 Pro ~94.1–94.3; Opus 4.8 ~93.6 — gaps inside measurement noise [SECONDARY]. Sources: https://tech-insider.org/claude-vs-chatgpt-vs-gemini-2026/ and https://neuralcoretech.com/gpt-5-5-vs-claude-opus-4-7-vs-gemini-3-1-pro-2026-benchmark/
- GPQA older rows: GPT-5.2 Pro 93.2; Gemini 3 Pro 91.9; Opus 4.6 91.3; Sonnet 4.6 89.9; Opus 4.5 87.0 [SECONDARY]. Source: https://github.com/rolandtolnay/mindsystem/blob/HEAD/references/models/claude-sonnet-4-6.md
- Qwen3.7 Max Thinking: GPQA 92.4%, LM Arena ELO 1475, 235B/22B active, ~197 tok/s [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- Qwen3.7 Max Thinking: BenchLM 77.4/100 [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- Qwen3.7 Plus: ~1420 ELO, ~88% GPQA, 72B/18B active, ~280 tok/s; Turbo: ~1380 ELO, ~450 tok/s — the efficiency tier [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- Qwen3.6-27B: SWE-bench Verified 77.2%, GPQA Diamond 87.8%, AIME 2026 94.1% (secondary compilation) [SECONDARY]. Source: https://toknow.ai/posts/qwen36-deepseek-v4-china-open-weight-frontier-models/index.pdf
- AIME 2026 evidence is weak and conflicting: GLM-5.2 claimed at 99.2 (vendor-derived review), Qwen3.5-plus at 91.3 (research note), and a leaked DeepSeek V4 99.4 flagged as impossible/fabricated with Epoch reportedly rejecting associated FrontierMath claims — keep the 99.4 as contradiction-only, do not build an AIME table [UNVERIFIED]. Sources: https://www.aqalion.com/blog/glm-5v-turbo-review-benchmarks-vs-claude-gpt-gemini and https://github.com/jamoeight/claude-code-deep-research-v2/blob/HEAD/research/notes/7_benchmarks_capability.md
- All Claude v5 models expose five effort levels (low/medium/high/xhigh/max) with high as the API default [SECONDARY]. Sources: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md and https://lushbinary.com/blog/claude-opus-5-vs-fable-5-sonnet-5-model-selection-cost/
- Opus 5 ships a Fast mode at roughly 2.5x speed for twice the base price [SECONDARY]. Source: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- Frontier-Bench v0.1: Opus 5 scores 43.3% at max effort and 44.4% at xhigh — higher at xhigh than max [SECONDARY]. Source: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- Opus 5 cost per task (AA-Briefcase): ~$1.78 at low effort to ~$17.79 at max — a 10x range within one model ID [SECONDARY]. Source: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- No long-context premium on Claude 4.6+ / v5 models — a 900K-token request bills at the same per-token rate as a 9K one, unlike OpenAI's >272K 2x/1.5x rule [SECONDARY]. Source: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
