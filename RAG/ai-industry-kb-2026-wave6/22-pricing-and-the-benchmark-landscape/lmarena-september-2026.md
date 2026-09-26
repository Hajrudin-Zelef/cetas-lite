---
id: ai-industry-kb-2026-wave6/22-pricing-and-the-benchmark-landscape/lmarena-september-2026
title: "LMArena, September 2026"
domain: pricing-and-the-benchmark-landscape
role: deep-dive
task: pricing
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Meta", "Moonshot", "OpenAI", "Z.ai", "xAI"]
dates: ["2026-07", "2026-08-19", "2026-08-28", "2026-09", "2026-09-03", "2026-09-04", "2026-09-07", "2026-09-17", "2026-09-21"]
keywords: ["agent", "agentic", "agents", "astra", "benchmark", "claude", "cost", "deepseek", "fable 5", "gemini", "gemini 3.8", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10561, 10612]
section: "§22. Pricing and the Benchmark Landscape"
delta_of: ai-industry-kb-2026
sha256: 0831bd68c86db4fc012810bbe2b052d984e8fb00dc5e9fdd0d80b8352b4981b6
---

# LMArena, September 2026

- Published **2026-09-07** — the third revision in under a week (v4.1.1 on 2026-09-03 → v4.2 on 2026-09-04 → v4.3 on 2026-09-07), an accelerated rollout of changes planned for Index v5 [SECONDARY].
- Methodology changes: **Terminal-Bench upgraded v2.1→v4.0**; **τ³-Banking replaced by AutomationBench-AA** (built with Zapier: 657 held-out business-workflow tasks; any guardrail violation zeroes the task); **private-test weighting 40%→45%** [SECONDARY]. Category weights unchanged: Agents 30%, Coding 20%, General 30%, Scientific Reasoning 20%.
- Leaders (v4.3): **Claude Fable 5.1 (max w/ fallback) = GPT-6 Astra (max) = 53 (tie)**; Claude Opus 5 51; Muse Spark 1.3 48 [SECONDARY]. Cost-per-task at the top: Astra $3.26 vs Fable $7.63 (57% cheaper for the same index score) [SECONDARY].
- Open-weight scores (v4.3): **Qwen3.8-Max (0902) 45** [SECONDARY, r/LocalLLaMA via alextech, 2026-09-17 — the 0902 checkpoint reclaimed the China lead], GLM-5.3 44.9, Kimi K3 43.8, GLM-5.3-Flash 42, DeepSeek V4.1-Flash 40, DeepSeek V4-Pro-0813 36 [SECONDARY].
- Later report: MiMo-V2.6-Pro scored **46.32** on the AA Index v4.3 (announced 2026-09-21/22, after the Sept 7 snapshot) — highest open-weight figure reported, ahead of Kimi K3 and Qwen3.8-Max [SECONDARY, unite.ai/VentureBeat]; treat as a separate dated claim, not a merged re-ranking.
- **Scale caution**: one tracker reports Qwen3.8-Max at **58** and Kimi K3 (max) at **60** "on the Intelligence Index," and benchlm's Sept-21 AA mirror shows GPT-5.6 Sol at **58.9%** — a different AA presentation scale (percentage view) than the 53-point Index view. **Never mix the 53-point and 58.9%-style figures.**
- **Standing rule**: scores from v4.1/v4.2 are not comparable to v4.3 — GPQA-Diamond was dropped as solved in the rebase; absolute scores fell. "Any production decision made using last week's leaderboard was made on a methodology that no longer applies" [SECONDARY, techtimes].

### LMArena, September 2026

- Text board (Sept 2026 snapshot): **Claude Fable 5 ~1525 (#1)**, Opus 5 1522, GPT-5.6 Sol 1514, Opus 4.8 1512, Grok 4.5 1499, Gemini 3.1 Pro Preview 1500, **Kimi K3 ~1500**, GLM-5.2 1483, Sonnet 5 1479, DeepSeek V4 Pro 1462, Qwen 3.7 Max 1455, V4 Flash 1441, Llama 4 Maverick 1352 [SECONDARY]. Volume: 10M+ evaluations/month.
- **Coding (Frontend Code) arena: Kimi K3 #1 at 1,679 Elo** — the first open model to top a board outright [SECONDARY].
- Gaming/neutrality concerns are now mainstream: (1) models optimize specifically for Arena — human vote biases reward polish over task success ("if you pick a model off LMArena for a coding agent, you are optimising for the wrong thing") [SECONDARY]; (2) conflict of interest — Arena is a **$1.7B-valuation** startup ($150M Series A; $250M raised in 7 months) backed by **OpenAI, Google, and Anthropic**, and a **~$100M annualized-revenue evaluation business** [SECONDARY]; (3) Arena's factuality-weighted ranking (25% factual-accuracy weight over 2M+ labeled claims) moved GPT-5.5 **up 13 places to #7** and cratered Muse Spark **13 places to #20**, with Fable 5 slipping to #2 [SECONDARY, cryptobriefing ~July 2026] — the default Elo ordering is preference, not truth.

### SWE-bench

- SWE-bench **Verified** (Vals leaderboard, 2026-08-19): **Claude Opus 5 97.0%**, GPT-5.6 Sol Max 96.2%, Claude Fable 5 Max 95.0% — **effectively saturated** as a frontier differentiator [SECONDARY].
- SWE-bench **Pro** (Scale's official standardized public board): **Muse Spark 1.1 + mini-SWE-agent 61.5%** leads the public board [SECONDARY].
- Vendor-reported, non-comparable harness numbers: Claude Fable 5 80.0%, GPT-5.6 Sol 64.6% (OpenAI model card) [SECONDARY] — **never mix vendor-aggregate with the Scale public board**.
- The three number classes must never be mixed: (1) Scale public standardized, 731 tasks; (2) Scale commercial/private standardized, 276 tasks; (3) vendor-aggregate/custom harness.

### Terminal-Bench 4.0

- Released **Sept 1–2, 2026** (announced by tbench.ai **2026-08-28**): 66 tasks (19 fixed, 8 removed as saturated/refusal-prone/publicly solved), mini-SWE-agent harness, 3 runs/task, 5 trials/task, 8-hour agent timeout [SECONDARY].
- Official tbench.ai runs: **Claude Fable 5.1 57.9%±3.8 (#1)**, Opus 5 51.8%, Fable 5 44.5%, **GLM-5.3 41.8%**, GPT-5.6 Sol 37.3%, Gemini 3.8 Flash 19.1%±3.4 [SECONDARY].
- Anthropic's own TB-4.0 numbers: **Mythos 5.1 60.9%**, Fable 5.1 55.8%, Opus 5 52.3% [VENDOR].
- BenchLM's Sept-21 mirror: **GPT-6 Astra 58.18%** leads its snapshot; alextech (Sept 22): GPT-6 Astra 60%, Fable 5.1 55%, **DeepSeek V4.1 Flash 27%**, **Grok 4.7 26%** [SECONDARY]. Different snapshots/dates; confidence intervals overlap — attribute per source.
- **TB 4.0 scores are NOT comparable with TB 2.1** (where V4.1 Flash took #1 at 90.6% on the Sept-11 update) — different task set, harness, and difficulty.

### Vals AI and FACTS

- Vals (vals.ai) is the 2026-era independent benchmark operator running verified leaderboards; its SWE-bench Verified board (2026-08-19) is the source for the saturation figures above [SECONDARY].
- Google DeepMind's **FACTS Benchmark Suite** (announced early 2026; initial leaderboard run Dec 2025): four dimensions — Parametric, Search, Multimodal, Grounding v2; 3,513 public examples plus a private Kaggle holdout; LLM-judge ensemble (models rate their own outputs 3.23pp higher on average) [VENDOR/SECONDARY].
- Initial run: **Gemini 3 Pro 68.8** (Search 83.8, Parametric 76.4, Multimodal 46.1), Gemini 2.5 Pro 62.1, GPT-5 61.8, Grok 4 53.6, Claude 4.5 Opus 51.3 [VENDOR/SECONDARY].
- **The ~70% factuality ceiling: no model has broken 70%** — the headline finding [SECONDARY, VentureBeat et al.]. For RAG builders the Search sub-score (Gemini 3 Pro 83.8) is the procurement-relevant number; Multimodal sub-scores (~46–47% even at the top) are the caution flag.

### Who leads what, as of September 2026 (version-pinned)

| Category (version pinned) | Leader |
|---|---|
| Composite intelligence (AA Index **v4.3**, 2026-09-07) | **Fable 5.1 = GPT-6 Astra (53)** — tie; Astra 57% cheaper per task [SECONDARY] |
| Open-weight composite (AA **v4.3**) | Qwen3.8-Max (0902) 45 [SECONDARY] > GLM-5.3 44.9 > Kimi K3 43.8 > GLM-5.3-Flash 42 > V4.1-Flash 40 > V4-Pro-0813 36 |
| Human preference text (LMArena, Sept 2026) | Claude Fable 5 (~1525) [SECONDARY] |
| Human preference coding (LMArena Frontend Code) | Kimi K3 (1,679) — first open #1 [SECONDARY] |
| Agentic terminal (Terminal-Bench **4.0**, official) | Claude Fable 5.1 (57.9%); vendor-self-reported: Mythos 5.1 60.9% [SECONDARY/VENDOR] |
| Issue-resolution coding (SWE-bench Verified) | Saturated — Opus 5 97.0%; live board is SWE-bench Pro (Muse Spark 1.1 61.5% official) [SECONDARY] |
| Factuality/grounding (FACTS Suite, Dec 2025 run) | Gemini 3 Pro (68.8); ceiling ~70% unbroken [VENDOR/SECONDARY] |
| Cost per unit of intelligence (AA $/Index-task, Sept 2026) | DeepSeek V4.1 Flash ($0.27) — cheapest measured [SECONDARY] |


### New verified facts — expansion

