---
id: ai-industry-kb-2026-wave6/19-coding-and-reasoning-models/part-9
title: "§19. Coding and Reasoning Models (part 9)"
domain: coding-and-reasoning-models
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "MiniMax", "Moonshot", "OpenAI", "Poolside", "Z.ai", "xAI"]
dates: ["2026-06", "2026-09"]
keywords: ["agents", "benchmarks", "claude", "deepseek", "fable 5", "gemini", "glm", "gpt-5.6", "grok", "grok 4", "kimi", "leaderboard"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9563, 9581]
section: "§19. Coding and Reasoning Models"
delta_of: ai-industry-kb-2026
sha256: 25fc9821f9460e7be7ae20ea1a822b77185ed39c5547d9f185ef282f8adc26ac
---

# §19. Coding and Reasoning Models (part 9)

- SWE-bench Pro design: 1,865 tasks across 41 professional repositories, split into public, commercial (private), and held-out sets; fresh repositories rotate in partly to reset contamination [SECONDARY]. Source: https://www.morphllm.com/claude-benchmarks
- Two SWE-bench Pro score families exist and are NOT comparable: vendor-run (each lab's own scaffold, 15–30 points higher) versus Scale's SEAL standardized leaderboard (same harness for every model) [SECONDARY]. Source: https://www.morphllm.com/claude-benchmarks
- Vendor aggregate (llm-stats, September 2026, 57 models, up from 44 in August): Claude Fable 5 leads at 80.0% all-time high; Anthropic's own launch figure was 80.3% [SECONDARY]. Source: https://www.morphllm.com/swe-bench-pro
- September 2026 vendor board top ranks: Fable 5 80.0%, Claude Mythos Preview 77.8%, Claude Opus 4.8 69.2%, Qwen3.8 Max (2.4T) 67.7%, Grok 4.5 64.7%, GPT-5.6 Sol 64.6%, Claude Opus 4.7 64.3%, GPT-5.6 Terra 63.4%, Claude Sonnet 5 63.2%, GPT-5.6 Luna 62.7% [SECONDARY]. Source: https://www.morphllm.com/swe-bench-pro
- Open-weight ranks on the September vendor board: GLM-5.2 62.1% (top open-weight), Qwen3.8 Flash/Flash-Next 62.5%, Qwen3.8-27B 61.7%, Muse Spark 1.1 (Meta) 61.5%, Qwen3.7 Max 60.6%, Laguna S 2.1 (Poolside) 59.4%, MiniMax M3 59.0%, Kimi K2.6 58.6%, GLM-5.1 58.4%, MiMo V2.5 Pro 57.2%, MiniMax M2.7 56.2%, DeepSeek-V4-Pro-Max 55.4%, DeepSeek-V4-Flash-Max 52.6% [SECONDARY]. Sources: https://www.morphllm.com/swe-bench-pro and https://github.com/boreilly-dc/sota-reference/blob/HEAD/agents/open-models-coding-agents.md
- New model names surfacing on the Pro board: Qwen3.8 Max (2.4T), Qwen3.8-Flash-Next, Qwen3.8-27B, Tencent Hy4 preview (65.7%) and Hy3 (57.9%), Meta Muse Spark 1.1 (61.5%), Poolside Laguna S 2.1 (59.4%) [SECONDARY]. Source: https://www.morphllm.com/swe-bench-pro
- Qwen3.8 Max is listed as "Closed API (Alibaba)" — unlike Qwen3.7 Max which is open-weight; the 3.8 generation's openness is unverified [SECONDARY]. Source: https://futuretweets.com/swe-bench-pro-leaderboard-2026/
- Claude Fable 5.1 appears at 81.2% Pro (September 1 BenchLM update) — a newer point release above Fable 5's 80.0%; pricing $10/$50, 1.0M context [SECONDARY]. Source: https://futuretweets.com/swe-bench-pro-leaderboard-2026/
- Pro differentiates where Verified saturates: 16-point spread (Fable 5 80.3% vs GPT-5.6 Sol 64.6%) versus 3.6 points on Verified; the ranking REVERSES (Opus 5 leads Verified 97.0% vs Fable 5 95.0%, but Fable 5 leads Pro 80.3% vs Opus 5 79.2%) [SECONDARY]. Source: https://stackfutures.com/blog/swe-bench-pro-differentiates-frontier-models-verified-saturating-2026/
- GPT-5.6 Sol drops the most between generations: 96.2% Verified → 64.6% Pro (−31.6 pts), nearly twice either Anthropic model's drop — suggesting Sol is tuned to Verified's task structure [SECONDARY]. Source: https://stackfutures.com/blog/swe-bench-pro-differentiates-frontier-models-verified-saturating-2026/
- Standardized (Scale SEAL) Pro board: Muse Spark 1.1 (Meta) 61.5% and GPT-5.4 (xHigh) 59.1% lead — the 80-vs-61 gap to the vendor board is the scaffolding effect in one line [SECONDARY]. Source: https://localaimaster.com/models/swe-bench-explained-ai-benchmarks
- LiveCodeBench (Vals.ai standardized, June 2026): Claude Fable 5 89.78%, Gemini 3.1 Pro Preview 88.48%, GPT-5.2 Codex 87.99%, Claude Opus 4.8 87.82%, Gemini 3.5 Flash 87.60%, DeepSeek V4 Pro 87.48%, GPT-5.3 Codex 87.31%, Qwen3.7 Max 87.06%, Kimi K2.6 86.77%, Nemotron 3 Ultra 85.98%, Qwen3.6 Plus 85.95%, GLM-4.7 82.23%, MiniMax M3 82.15%, GLM-5/5.1 81.4–81.9% — the open–closed gap on pure algorithmic coding is now ~2 points [SECONDARY]. Source: https://github.com/boreilly-dc/sota-reference/blob/HEAD/agents/open-models-coding-agents.md
- Kimi K3 has no SWE-bench Pro entry yet — Moonshot published Terminal-Bench 2.1 (88.3%) and DeepSWE (67.5%) instead [SECONDARY]. Source: https://localaimaster.com/models/swe-bench-explained-ai-benchmarks
- Pro output-price ladder (September 2026, $/M): Fable 5 $50, Opus 4.8 $25, Opus 4.7 $25, Sol $30, Terra $12, Sonnet 5 $10, Luna $1.20, Qwen3.8 Max $4.95, Grok 4.5 $6.00, GLM-5.2 $2.40, Qwen3.8 Flash $0.47, Qwen3.8-27B $3.00, Muse Spark 1.1 $4.25, Qwen3.7 Max $3.75, Laguna S 2.1 $0.20, MiniMax M3 $1.10, Gemini 3.6 Flash $7.50, GPT-5.5 $30, Kimi K2.6 $3.50, GLM-5.1 $3.50, Hy3 $0.58, GPT-5.4 $15, GPT-5.3 Codex $14, MiniMax M2.7 $1.20, DeepSeek-V4-Pro-Max $2.60, Gemini 3.1 Pro $12, DeepSeek-V4-Flash-Max $0.18 [SECONDARY]. Source: https://www.morphllm.com/swe-bench-pro
- Month-over-month price moves on the Pro board (August → September 2026): DeepSeek-V4-Pro-Max $3.20 → $2.60, GLM-5.2 $3.00 → $2.40, GLM-5.1 $4.40 → $3.50, DeepSeek-V4-Flash-Max $0.20 → $0.18, Qwen3.8 Max n/a → $4.95 — dated snapshots, not errors [SECONDARY]. Source: https://www.morphllm.com/swe-bench-pro versus https://www.morphllm.com/ai-coding-benchmarks-2026
- Anthropic's launch figures versus aggregator: Fable 5 80.3% (Anthropic) vs 80.0% (llm-stats); Opus 5 79.2% (Anthropic launch, not yet in the aggregator) — vendor figures precede third-party listing [SECONDARY]. Sources: https://www.morphllm.com/claude-benchmarks and https://localaimaster.com/models/swe-bench-explained-ai-benchmarks

**Additional facts, sixth tranche (Terminal-Bench generations):**

