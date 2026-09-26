---
id: ai-industry-kb-2026-wave6/22-pricing-and-the-benchmark-landscape/figures-and-metrics
title: "Figures and metrics"
domain: pricing-and-the-benchmark-landscape
role: deep-dive
task: pricing
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Meta", "Mistral", "Moonshot", "OpenAI", "Poolside", "Z.ai", "xAI"]
dates: ["2026-02-19", "2026-04", "2026-04-23", "2026-07-09", "2026-07-12", "2026-07-16", "2026-07-21", "2026-08-10", "2026-08-11", "2026-08-18", "2026-08-19", "2026-08-21", "2026-08-26", "2026-08-28", "2026-09-01", "2026-09-03", "2026-09-04", "2026-09-07", "2026-09-10", "2026-09-14", "2026-09-17", "2026-09-21", "2026-09-22", "2026-11-21"]
keywords: ["astra", "benchmark", "benchmarks", "claude", "cost", "deepseek", "fable 5", "gemini", "gemini 3.8", "glm", "gpt-5.6", "gpt-6"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10693, 10766]
section: "§22. Pricing and the Benchmark Landscape"
delta_of: ai-industry-kb-2026
sha256: 5aac275cb3077bf726a7122365d3560a50545edceb39234414e0afa029cc2b94
---

# Figures and metrics

## Figures and metrics

- **Input spread: 111×** ($0.09 Laguna S 2.1 → $10.00 GPT-6 Astra/Fable 5.1); intra-generation spread: **20×** (Sol/Luna).
- Cheapest frontier API: V4.1 Flash off-peak $0.15/$0.60; cheapest per Index task: $0.27 [SECONDARY].
- AA Index v4.3 (2026-09-07): Fable 5.1 = Astra 53 · Opus 5 51 · Spark 1.3 48 · Qwen3.8-Max (0902) 45 · GLM-5.3 44.9 · Kimi K3 43.8 · GLM-5.3-Flash 42 · V4.1-Flash 40 · V4-Pro-0813 36.
- LMArena text (Sept 2026): Fable 5 ~1525 · Opus 5 1522 · GPT-5.6 Sol 1514 · Kimi K3 ~1500 · Llama 4 Maverick 1352. Coding arena: Kimi K3 1,679.
- TB 4.0 (official): Fable 5.1 57.9% · Opus 5 51.8% · GLM-5.3 41.8% · GPT-5.6 Sol 37.3% · Grok 4.7 26% · V4.1 Flash 27%.
- SWE-bench Verified: Opus 5 97.0% (saturated). SWE-bench Pro (official): Muse Spark 1.1 61.5%. FACTS Suite: Gemini 3 Pro 68.8 (ceiling ~70%).
- Cache-read discounts reach ~97.5% off input (Anthropic $0.25/M).


### New verified metrics — expansion

- Input-price spread across the 2026 frontier (per 1M tokens): GLM-5.3-Flash $0.15 → Gemini 3.8 Flash Lite $0.30 → Mistral Small 4 $0.15 / Gemini 3.7 Flash $0.75 → Mistral Large 3 $0.50 → GLM-5.3 $1.40 → Kimi K3 $3 → Qwen3.7 Max Thinking $4 → GPT-5.6 Terra $8 → Gemini 3.8 Flash Pro $4.50 → Fable 5 $10 → GPT-5.6 Max $9 → Gemini 3.8 Flash Max $9 → GPT-5.5 Pro $15 — a 100x input-price range across usable frontier models [SECONDARY].
- Cache-read discounts: K3 10x ($3→$0.30), GLM-5.3 ~5.4x ($1.40→$0.26), Gemini Flash 10x ($0.75→$0.075), GPT-5.6 explicit write 1.25x surcharge [SECONDARY].
- Batch discounts: OpenAI batch 0.5x; Anthropic batch 0.5x — non-urgent workloads halve in price on both platforms [SECONDARY].
- Long-context penalty: OpenAI 2x input / 1.5x output above 272K on GPT-5.4/5.5/5.6 — a 1M-token Sol request bills like 2M input tokens [SECONDARY].
- Terminal-Bench 2.1 cost-per-run extremes: $18 total inference for DeepSeek V4.1 Flash (370/445 trials) versus $2,059.19/run for GPT-5.5 + Codex xhigh — benchmark cost and benchmark score are independent axes [SECONDARY].
- DeepSWE cost-per-point (Sept 2026): Gemini 3.8 Flash high 74% at $2.36 versus Fable 5 xhigh 70% at $13.41 — the price-performance frontier favors the cheap [SECONDARY].
- Mistral price history: tracking since April 2026; Large 3's $2/$6 → $0.50/$1.50 move is the largest 2026 frontier price cut documented here [SECONDARY].


- DeepSWE price-per-point, derived (Sept 2026 board): Gemini 3.8 Flash high $2.36/74% = ~$0.032 per point; GLM-5.3 max $3.99/69% = ~$0.058; Kimi K3 max $4.65/69% = ~$0.067; GPT-6 Astra xhigh $4.43/74% = ~$0.060; GPT-5.6 Sol max $6.46/73% = ~$0.088; Fable 5 xhigh $13.41/70% = ~$0.19 — a 6x spread per benchmark point [DIRECTIONAL — derived from sourced board figures]. Source: https://deepswe.datacurve.ai/
- Terminal-Bench 2.1 cost parity: DeepSeek V4.1 Flash 83.9% at ~$18 total inference versus GPT-5.5 + Codex xhigh 83.1% at $2,059.19/run — equal scores, ~114x cost ratio [DIRECTIONAL — derived]. Sources: https://github.com/antigmalabs/ante/blob/HEAD/docs-site/docs/benchmarks/eval.mdx and https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- Terminal-Bench 2.1 official-board cost ladder: Fable 5 + Claude Code xhigh $552.67/run (83.8%) | GPT-5.6 Terra + Codex max $421.15 (78.4%) | GPT-5.6 Luna + Codex max $241.45 (75.7%) — cost falls with the tier, scores roughly follow [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- Contradiction-resolution status: (1) GPT-5.6 Sol $4/$20 vs $5/$30 — UNRESOLVED, check OpenAI live page; (2) Sonnet 5 $2/$10 vs $3/$15 — UNRESOLVED, check Anthropic live page; (3) Gemini Flash $3.75 vs $7.50 — RESOLVED (promo vs scheduled); (4) Mistral Large 3 $2/$6 vs $0.50/$1.50 — RESOLVED (stale vs current); (5) Qwen3.7 Plus $0.32 vs $0.42 — UNRESOLVED provider variance; (6) Mistral Nemo $0.02/$0.10 vs $0.15/$0.15 — UNRESOLVED; (7) K3 $20M revenue prong — UNRESOLVED [SECONDARY].
- Price-verification freshness: the newest confirmations are 2026-09-22 (Mistral, Gemini tiers); the oldest structural rules are 2026-07-12 (OpenAI 272K) — a 10-week window; prices move inside it [SECONDARY].


- Price-check freshness ledger: OpenAI 10 weeks old (2026-07-12), Claude 6 weeks (2026-08-11), cache landscape 6 weeks (2026-08-10/14), Mistral/Gemini current (2026-09-22) — the §22 table's rows have different ages; the oldest structural rule (272K) is also the most stable [SECONDARY].
- Unresolved contradictions carried forward: GPT-5.6 Sol ($4/$20 vs $5/$30), Sonnet 5 ($2/$10 vs $3/$15), Qwen3.7 Plus ($0.32 vs $0.42), Mistral Nemo ($0.02/$0.10 vs $0.15/$0.15), K3 revenue prong — five open items for the next verification pass [SECONDARY].


- Output-price spread: 240x ($0.50 → $120) exceeds the input-price spread (100x) — output tokens are where the 2026 price war is steepest [DIRECTIONAL — derived].
- Effective-price ranking at 90% cache hit: GLM-5.3 $0.37 < Kimi K3 $0.57 < Gemini 3.7 Flash promo $0.14 (0.9×$0.075+0.1×$0.75) — caching reshuffles the leaderboard [DIRECTIONAL — derived].

## Main actors

- **Artificial Analysis** — Intelligence Index v4.3 (2026-09-07); three revisions in one week; $/Index-task cost metric.
- **LMArena / Arena** — $1.7B valuation, ~$100M annualized-revenue evaluation business, funded by OpenAI/Google/Anthropic; text and coding boards.
- **Scale** — SWE-bench Pro official standardized public board (731 tasks).
- **Vals** — independent verified leaderboards (SWE-bench Verified 2026-08-19).
- **Google DeepMind** — FACTS Benchmark Suite operator.
- **Zapier** — co-builder of AutomationBench-AA (657 held-out business-workflow tasks).
- **DeepSeek** — cheapest frontier-class API (V4.1 Flash off-peak).
- **OpenAI** — three-tier 5.6 pricing (20× Sol/Luna spread); GPT-6 Astra ($10/$50, 272K cliff).
- **Anthropic** — TB 4.0 leader (Fable 5.1 57.9% official); LMArena text leader; cancelled the Sonnet 5 price rise.
- **Moonshot** — Kimi K3 coding-arena #1 (1,679 Elo), free quota gateway.
- **xAI** — Grok 4.7 at $2/$6 (200K cliff).
- **Resellers (e.g. Qubax)** — up to ~91% off retail; the effective price floor [SECONDARY].

## Timeline and context

- **2026-02-19** — Gemini 3.1 Pro pricing ($2/$12 ≤200K, $4/$18 above) [VENDOR].
- **2026-04-23** — GPT-5.5 / 5.5 Pro pricing ($5/$30, $30/$180) [VENDOR].
- **2026-07-09** — GPT-5.6 Luna GA at $1/$6; falls to $0.20/$1.20 by early September [VENDOR/SECONDARY].
- **2026-07-16** — Kimi K3 API pricing ($3/$15) [VENDOR].
- **2026-07-21** — Poolside Laguna S 2.1 ($0.09/$0.18) [VENDOR].
- **2026-08-11** — Sonnet 5's planned Sept-1 rise to $3/$15 cancelled (intro $2/$10 permanent) [SECONDARY].
- **2026-08-18** — GLM-5.3 pricing ($1.40/$4.40) [VENDOR].
- **2026-08-19** — Vals SWE-bench Verified board: Opus 5 97.0% (saturation documented) [SECONDARY].
- **2026-08-21** — OpenAI cuts GPT-5.6 Sol 20% in / 33% out (promo through 2026-11-21) [VENDOR].
- **2026-08-26** — GLM-5.3-Flash pricing ($0.15 input) [VENDOR].
- **2026-08-28** — Terminal-Bench 4.0 announced (tbench.ai) [SECONDARY].
- **2026-09-01/02** — Terminal-Bench 4.0 released; official runs: Fable 5.1 57.9% #1 [SECONDARY].
- **2026-09-03** — AA Index v4.1.1; GPT-6 Astra launches ($10/$50, 272K cliff) [VENDOR/SECONDARY].
- **2026-09-04** — AA Index v4.2 (rebase; GPQA-Diamond dropped) [SECONDARY].
- **2026-09-07** — AA Index v4.3 (TB v4.0, AutomationBench-AA, 45% private weighting): Fable 5.1 = GPT-6 Astra = 53 [SECONDARY].
- **2026-09-10** — DeepSeek V4.1 Flash GA (peak/off-peak pricing) [VENDOR].
- **2026-09-14** — V4-Pro API traffic routes to V4.1 Flash at V4.1 rates [SECONDARY].
- **2026-09-17** — Qwen3.8-Max (0902) reported at AA v4.3 45, reclaiming China's lead [SECONDARY].
- **2026-09-21** — Grok 4.7 launch ($2/$6 ≤200K, $4/$12 above) [VENDOR].


