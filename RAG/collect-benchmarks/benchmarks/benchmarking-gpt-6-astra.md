---
id: collect-benchmarks/benchmarks/benchmarking-gpt-6-astra
title: "Benchmarking GPT-6 Astra"
domain: benchmarks
role: reference
task: article
actors: ["Anthropic", "Meta", "OpenAI", "Z.ai", "xAI"]
dates: ["2026-09-09", "2026-09-23"]
keywords: ["astra", "benchmark", "gpt-6", "agent", "claude", "cost", "fable 5", "glm", "gpt-5.6", "grok", "grok 4", "leaderboard"]
source: docs/RAG/Collect RAG/06_benchmarks/benchmarking-gpt-6-astra.md
source_anchor: ""
source_lines: [1, 54]
sha256: 8af1a7d4e0b7465968a1e410e45c6fae241d7153e1329f958cd7a56b8946c5bf
---

# Benchmarking GPT-6 Astra

## Metadata

- **Source** : https://artificialanalysis.ai/articles/benchmarking-gpt-6-astra
- **Site** : ArtificialAnalysis
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This Artificial Analysis article, published September 9, 2026, evaluates GPT-6 Astra, which ties Claude Fable 5.1 for leadership on both flagship Artificial Analysis indices—the Intelligence Index and the Coding Agent Index—while doing so at substantially lower cost. On the Intelligence Index, GPT-6 Astra (max) scores 53, level with Claude Fable 5.1 (max with fallback) and 6 points above its predecessor GPT-5.6 Sol (max). Pricing rose to $10/$50 per million input/output tokens (2.5x GPT-5.6 Sol's $4/$20), with a 90% cache-read discount and 25% cache-write premium. Despite higher token prices, Astra matches Fable 5.1's Intelligence Index score at ~40% of the cost per task ($3.26 vs $7.63), because of dramatically lower token usage.

Methodologically, the article uses Pareto-frontier analysis across two axes: Intelligence Index vs Cost per Task, and Intelligence Index vs Output Tokens per Task. Every Astra reasoning effort level (low to max) sits on both frontiers. At max effort it uses 27k output tokens per task versus 78k for Claude Fable 5.1—about one third—for the same score. At max effort Astra is ~60% more expensive per task than GPT-5.6 Sol.

Notable evaluation results: Astra leads Terminal-Bench v4.0 at 59% (Fable 5.1 52%, GPT-5.6 Sol 40%), AutomationBench-AA at 69% (Grok 4.6 67%, GPT-5.6 Sol 60%), and GDP.pdf (passes all criteria on 31% of attempts vs 27%). In AA-Omniscience, hallucination rate halves from 92% to 51% at max effort while accuracy rises 4 points. AA-Briefcase gains ~90 Elo, with higher rubric and Analytical Quality Elo but lower Presentation Quality Elo, where GPT-5.6 Sol still leads. Conversely, GDPval-AA v2 drops ~45 Elo, attributed to Astra using far fewer turns (24 per task vs 45 for Sol, 60 for Fable 5.1/Opus 5).

On the Coding Agent Index, Astra scores 62 in Codex, tied with Claude Fable 5.1 in Claude Code (62) and ahead of Claude Opus 5 (60), GPT-5.6 Sol (55), and Muse Spark 1.3 (54). Its 7-point gain over Sol comes from Terminal-Bench v4.0 (56% vs 37%) and SWE-Atlas-QnA (62% vs 54%), partly offset by DeepSWE (68% vs 72%). At max effort it costs $7.09 per task in the Coding Agent Index, ~40% cheaper than Fable 5.1 and ~30% cheaper than Opus 5 for equal or higher scores.

## Key points

- GPT-6 Astra (max) scores 53 on the Intelligence Index, tying Claude Fable 5.1 and gaining 6 points on GPT-5.6 Sol.
- Ties the top Coding Agent Index score of 62 in Codex, ahead of Opus 5 (60), GPT-5.6 Sol (55), and Muse Spark 1.3 (54).
- Matches Fable 5.1 at ~40% of cost per task on Intelligence, ~40% cheaper on coding, driven by token efficiency.
- Halves hallucination rate (92% → 51%) on AA-Omniscience while raising accuracy 4 points.
- Leads Terminal-Bench v4.0 (59%), AutomationBench-AA (69%), and GDP.pdf (31% all-criteria).
- Regresses ~45 Elo on GDPval-AA v2, using only 24 turns per task vs 45–60 for rivals.
- Pricing 2.5x GPT-5.6 Sol: $10/$50 per M input/output tokens, 90% cache-read discount.

## Technical data / figures

| Metric | GPT-6 Astra (max) | Claude Fable 5.1 | GPT-5.6 Sol (max) | Others |
|--------|-------------------|-------------------|-------------------|--------|
| Intelligence Index | 53 | 53 (max w/ fallback) | 47 | — |
| Coding Agent Index | 62 (Codex) | 62 (Claude Code) | 55 | Opus 5: 60; Muse Spark 1.3: 54 |
| Cost/task (Intelligence) | $3.26 | $7.63 | — | — |
| Output tokens/task | 27k | 78k | — | — |
| Cost/task (Coding) | $7.09 | ~$11.8 (≈40% more) | ~$6.2 (≈15% less) | Opus 5 ~$10.1 (≈30% more) |
| AA-Omniscience hallucination | 51% | — | 92% | — |
| Terminal-Bench v4.0 | 59% | 52% | 40% | — |
| AutomationBench-AA | 69% | — | 60% | Grok 4.6: 67%; GLM-5.3: 62% |
| GDP.pdf (all criteria) | 31% | — | 27% | — |
| GDPval-AA v2 Elo | −45 vs Sol | — | baseline | — |
| AA-Briefcase Elo | +90 vs Sol | — | baseline | — |
| SWE-Atlas-QnA | 62% | — | 54% | — |
| DeepSWE | 68% | — | 72% | — |
| Price ($/M in–out) | $10 / $50 | — | $4 / $20 | — |

## Why this source matters for the RAG

It is a primary-source benchmark analysis of a frontier model, capturing cost-efficiency, token-efficiency, and hallucination trade-offs that raw leaderboard scores omit. The Pareto-frontier framing and detailed per-evaluation deltas make it valuable for comparing frontier models on both capability and economics.

