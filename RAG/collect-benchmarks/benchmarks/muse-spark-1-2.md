---
id: collect-benchmarks/benchmarks/muse-spark-1-2
title: "Muse Spark 1.2: Benchmarks and analysis"
domain: benchmarks
role: reference
task: article
actors: ["Anthropic", "Meta", "Moonshot", "OpenAI", "SpaceX", "United States", "xAI"]
dates: ["2026-08-05", "2026-09-23"]
keywords: ["benchmarks", "muse", "muse spark", "agentic", "claude", "context window", "cost", "fable 5", "gpt-5.6", "grok", "grok 4", "kimi"]
source: docs/RAG/Collect RAG/06_benchmarks/muse-spark-1-2.md
source_anchor: ""
source_lines: [1, 64]
sha256: 2a8568b599e7931aa0ae174816025d5e0b1d3561169bf873e9597f03dc16c709
---

# Muse Spark 1.2: Benchmarks and analysis

## Metadata

- **Source** : https://artificialanalysis.ai/articles/muse-spark-1-2
- **Site** : ArtificialAnalysis
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This Artificial Analysis article, published August 5, 2026, analyzes Meta's Muse Spark 1.2, which scores 54 on the Artificial Analysis Intelligence Index. It is Meta's third release in four months and markedly improves agentic knowledge-work capabilities over prior versions, placing Meta effectively tied with SpaceXAI for third among US labs. Muse Spark 1.2 (xhigh) rises 3 points over Muse Spark 1.1 (51) and 11 points over Muse Spark 1.0 (43, April). It lands effectively tied with GPT-5.5 (xhigh, 55) and Grok 4.5 (high, 54), narrowly behind Claude Opus 5 (max, 61), Claude Fable 5 (max w/ fallback, 60), GPT-5.6 Sol (max, 59), and Kimi K3 (max, 57). Meta shared early access with Artificial Analysis for benchmarking.

The 3-point Index gain is concentrated in agentic evaluations. GDPval-AA v2 Elo jumps 260 points to 1631 (from 1371), ranking #5 overall and ahead of Claude Opus 4.8 (max, 1588) but behind Claude Opus 5 (1852), GPT-5.6 Sol (1730), and Kimi K3 (1685). GDPval-AA v2 is the leading metric for general agentic performance, measuring realistic knowledge-work tasks (presentations, analysis) with shell access and web browsing via the Stirrup reference harness. Terminal-Bench v2.1 gains 2 points (78% to 80%) and τ³-Banking rises 2 points (25% to 27%).

On cost, Muse Spark 1.2 costs $0.40 per Intelligence Index task at unchanged pricing of $1.25/$4.25 per million input/output tokens (cache hits $0.15/M). It is among the most cost-efficient models at its intelligence level—only Grok 4.5 (high, $0.37) and GPT-5.6 Sol (medium, $0.39) are cheaper in its cluster, while GPT-5.6 Terra (max, $0.51), Kimi K3 (max, $0.86), and GPT-5.5 (xhigh, $1.18) cost more. The increase over Muse Spark 1.1's $0.29 per task stems from higher token usage: input tokens up ~53% and output tokens up ~36%, concentrated in GDPval-AA v2.

On AA-Omniscience, the score rose from 18 to 22 as hallucination rate fell 10 points (38% to 28%) and attempt rate dropped from 82% to 67%; accuracy slipped from 41% to 38% due to heavy abstention. Scientific reasoning was largely unchanged: CritPt gained 3 points (15% to 18%), SciCode fell 2 points (58% to 56%), and Humanity's Last Exam fell 1 point (45% to 44%). The model retains a 1M-token context window and is available on Meta's first-party API at launch.

## Key points

- Muse Spark 1.2 scores 54 on the Intelligence Index, up 3 from 1.1 (51) and 11 from 1.0 (43).
- Effectively tied with GPT-5.5 (55) and Grok 4.5 (54); behind Opus 5 (61), Fable 5 (60), GPT-5.6 Sol (59), Kimi K3 (57).
- Agentic gains dominate: GDPval-AA v2 +260 Elo to 1631 (#5 overall), Terminal-Bench v2.1 +2 (80%), τ³-Banking +2 (27%).
- Cost $0.40 per Intelligence Index task at $1.25/$4.25 per M tokens; token usage up ~53% input / ~36% output vs 1.1.
- AA-Omniscience rises 18 → 22 via abstention: hallucination 38% → 28%, attempt rate 82% → 67%, accuracy 41% → 38%.
- Scientific reasoning mostly flat: CritPt +3 (18%), SciCode −2 (56%), HLE −1 (44%).
- Retains 1M-token context; available on Meta's first-party API at launch.

## Technical data / figures

| Metric | Muse Spark 1.2 | Muse Spark 1.1 | Comparison |
|--------|----------------|----------------|------------|
| Intelligence Index | 54 | 51 | 1.0: 43 |
| GDPval-AA v2 Elo | 1631 (#5) | 1371 | Opus 5: 1852; GPT-5.6 Sol: 1730; Kimi K3: 1685; Opus 4.8: 1588 |
| Terminal-Bench v2.1 | 80% | 78% | — |
| τ³-Banking | 27% | 25% | — |
| AA-Omniscience Index | 22 | 18 | — |
| Hallucination rate | 28% | 38% | — |
| Attempt rate | 67% | 82% | — |
| Omniscience accuracy | 38% | 41% | — |
| CritPt | 18% | 15% | — |
| SciCode | 56% | 58% | — |
| Humanity's Last Exam | 44% | 45% | — |
| Cost per Index task | $0.40 | $0.29 | Grok 4.5: $0.37; GPT-5.6 Sol (med): $0.39 |
| Pricing ($/M in–out) | $1.25 / $4.25 | unchanged | cache hits $0.15/M |
| Context window | 1M | 1M | — |

| Intelligence Index leaderboard | Score |
|--------------------------------|-------|
| Claude Opus 5 (max) | 61 |
| Claude Fable 5 (max w/ fallback) | 60 |
| GPT-5.6 Sol (max) | 59 |
| Kimi K3 (max) | 57 |
| GPT-5.5 (xhigh) | 55 |
| Muse Spark 1.2 (xhigh) | 54 |
| Grok 4.5 (high) | 54 |

## Why this source matters for the RAG

It provides a primary-source, version-over-version analysis of Meta's Muse Spark family, isolating where gains occur (agentic knowledge work) versus where they don't (scientific reasoning). The explicit cost-per-task and abstention/hallucination trade-off data are valuable for evaluating reliability and economics alongside raw capability.

