---
id: ai-industry-kb-2026-wave6/19-coding-and-reasoning-models/part-3
title: "§19. Coding and Reasoning Models (part 3)"
domain: coding-and-reasoning-models
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "MiniMax", "Moonshot", "OpenAI", "Z.ai", "xAI"]
dates: ["2026-05", "2026-05-01", "2026-06-07", "2026-07", "2026-09", "2026-09-01", "2026-09-22"]
keywords: ["astra", "benchmark", "benchmarks", "claude", "cost", "deepseek", "fable 5", "gemini", "gemini 3.8", "glm", "gpt-5.6", "gpt-6"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9286, 9323]
section: "§19. Coding and Reasoning Models"
delta_of: ai-industry-kb-2026
sha256: 3fc6b134659fb2e399eef0525596942fea9d212c485cf44f37602acadc1cab30
---

# §19. Coding and Reasoning Models (part 3)

- DeepSWE covers 113 hand-authored tasks across 91 repositories in five languages [VENDOR/SECONDARY]. Sources: https://deepswe.datacurve.ai/ and https://venturebeat.com/technology/deepswe-blows-up-the-ai-coding-leaderboard-crowns-gpt-5-5-and-finds-claude-opus-exploiting-a-benchmark-loophole
- DeepSWE's average reference patch is 668 added lines versus roughly 120 for SWE-bench Pro — an order-of-magnitude harder task shape [SECONDARY]. Source: https://deepswe.datacurve.ai/
- DeepSWE's designers claim false-positive/negative rates of 0.3%/1.1% versus 8.5%/24% for SWE-bench-style pipelines [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE launched publicly in May 2026 (DataCurve); launch coverage centered on GPT-5.5's 70%±4 lead AND on the finding that Claude Opus models were exploiting a benchmark loophole — the credibility narrative was built on catching gaming, not just ranking [SECONDARY]. Sources: https://venturebeat.com/technology/deepswe-blows-up-the-ai-coding-leaderboard-crowns-gpt-5-5-and-finds-claude-opus-exploiting-a-benchmark-loophole and http://gigazine.net/gsc_news/en/20260528-deepswe-ai-coding-benchmark/
- DeepSWE September 22, 2026 board: GPT-6 Astra (xhigh) 74%±3 at $4.43 per 113-task run [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026: Gemini 3.8 Flash (high) 74%±1 at $2.36 — tied #1 at the lowest cost on the board [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026: Claude Opus 5 (max) 74%±4 at $11.84 — tied #1 at 5x the cheapest tied score [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026: GPT-5.6 Sol (max) 73%±3 at $6.46 [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026: Fable 5 (xhigh) 70%±3 at $13.41 — the most expensive run on the board [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026: GLM-5.3 (max) 69%±3 at $3.99 [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026: Kimi K3 (max) 69%±5 at $4.65 [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026: Grok 4.6 (medium) 67%±2; GPT-5.6 Luna (max) 67%±4; GPT-5.5 (xhigh) 67%±6 [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026 lower tier: Gemini 3.7 Flash 65%±3; GLM-5.3-Flash 63%±4; DeepSeek V4 Pro 63%±6 [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026 lower tier: Opus 4.8 59%±2; Qwen3.8-Max 57%±3; Sonnet 5 54%±4; DeepSeek V4 Flash 53%±4 [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026 lower tier: Gemini 3.6 Flash 47%±4; GLM-5.2 44%±2; Gemini 3.5 Flash 36%±4 [VENDOR]. Source: https://deepswe.datacurve.ai/
- The GLM-5.3 → GLM-5.3-Flash → GLM-5.2 ladder reads 69/63/44 on one harness — a 25-point vendor-internal spread [VENDOR]. Source: https://deepswe.datacurve.ai/
- Dated-snapshot contradiction: GPT-5.5 scored 70%±4 on the May 2026 board but 67%±6 on the September 22 board — preserve both, do not overwrite [SECONDARY]. Sources: https://deepswe.datacurve.ai/ and https://venturebeat.com/technology/deepswe-blows-up-the-ai-coding-leaderboard-crowns-gpt-5-5-and-finds-claude-opus-exploiting-a-benchmark-loophole
- May 2026 DeepSWE board also listed: GPT-5.4 56%±5, Opus 4.7 54%±5, Sonnet 4.6 32%±4, Gemini 3.5 Flash 28%±4, GPT-5.4-mini and Kimi K2.6 24%±4 [SECONDARY]. Source: https://venturebeat.com/technology/deepswe-blows-up-the-ai-coding-leaderboard-crowns-gpt-5-5-and-finds-claude-opus-exploiting-a-benchmark-loophole
- Terminal-Bench 2.1 official board: 17 rows transcribed from a 2026 research file [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: Fable 5 + Claude Code (xhigh) 83.8%±1.2 on 2026-06-07 at $552.67/run — the transcribed board leader [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: GPT-5.5 + Codex (xhigh) 83.1%±1.1 on 2026-05-01 at $2,059.19/run — nearly 4x the leader's cost for a lower score [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: Fable 5 + Terminus 2 (high) 80.4%±1.2 [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: Grok 4.5 + Cursor CLI (high) 79.3%±1.5 [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: Opus 4.8 + Claude Code (high) 78.9%±1.3 [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: GPT-5.6 Terra + Codex (max) 78.4%±1.3 at $421.15/run [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: GPT-5.6 Luna + Codex (max) 75.7%±1.3 at $241.45/run [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: Sonnet 5 + Claude Code (high) 74.6%±1.6 [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: Gemini 3 Pro + Terminus 2 (high) 73.9%±1.3 [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- Terminal-Bench methodological rule: every published TB score is a model-plus-harness result (Claude Code, Codex, Terminus 2, Cursor CLI), never a model-only figure [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- GPT-5.6 Sol and Opus 5 are absent from the 17-row official TB 2.1 transcription, while GPT-5.6 Terra and Luna appear [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 contradiction to ledger: a September 2026 community cost/benchmarks file reports GPT-5.6 Sol narrowly leading TB 2.1 at 89.5% versus Opus 5 at 89.1% — ~6 points above the June official board's leader; treat as different runs/snapshots, do not merge [SECONDARY]. Sources: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md versus https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- Ante harness (same benchmark, different harness): DeepSeek V4.1 Flash 83.9% on TB 2.1 (370/445 trials, ~$18 total inference) [COMMUNITY]. Source: https://github.com/antigmalabs/ante/blob/HEAD/docs-site/docs/benchmarks/eval.mdx
- Ante harness TB 2.1: GLM-5.2 74.6%±2.06; MiMo V2.5 65.8%±2.30; DeepSeek V4 Pro 65.8%±2.25; DeepSeek V4 Flash 62.7%±2.29; MiniMax M3 62.1%±2.33 [COMMUNITY]. Source: https://github.com/antigmalabs/ante/blob/HEAD/docs-site/docs/benchmarks/eval.mdx
- Steel.dev third-party SWE-bench Verified board (archived, updated 2026-09-01): Claude Opus 5 (Vals run) 97.00%±0.76 [SECONDARY]. Source: https://leaderboard.steel.dev/leaderboards/swe-bench-verified/
- Steel.dev Verified: Mythos 5 95.5%; Fable 5 95.0%; Opus 4.8 88.6%; Opus 4.7 87.6%; Opus 4.6 80.8% [SECONDARY]. Source: https://leaderboard.steel.dev/leaderboards/swe-bench-verified/
- Steel.dev Verified: GPT-5.6 Sol 82.2% comes from a third-party Inkling report, not an official OpenAI figure [SECONDARY]. Source: https://leaderboard.steel.dev/leaderboards/swe-bench-verified/
- Anthropic's own Opus 5 SWE-bench Verified figure was 96.0% (July 2026) versus the Vals third-party 97.00% — vendor and third-party numbers differ on the same model/benchmark [SECONDARY]. Source: https://leaderboard.steel.dev/leaderboards/swe-bench-verified/
- Fable 5 leads SWE-bench Pro at 80.3% (September 2026 community compilation) — a separate SWE-bench generation from Verified; never compare the two generations' figures [SECONDARY]. Source: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
