---
id: collect-presse-fr/presse-fr/claude-opus-5-5-gpt-6-sol-numerama
title: "Claude Opus 5.5 et GPT-6 Sol sortent le même jour : Anthropic et OpenAI ont du mal à « ralentir » l'IA"
domain: presse-fr
role: reference
task: article
actors: ["Anthropic", "EU", "Google", "OpenAI", "SpaceX", "xAI"]
dates: ["2023-03", "2026-08", "2026-09", "2026-09-23"]
keywords: ["claude", "gpt-6", "opus 5", "sol", "agents", "astra", "benchmark", "benchmarks", "chatgpt", "cost", "cybersecurity", "fable 5"]
source: docs/RAG/Collect RAG/05_presse_fr/claude-opus-5-5-gpt-6-sol-numerama.md
source_anchor: ""
source_lines: [1, 64]
sha256: 627914f1d3a2afde07d0760fb202ec2ea293dbd3a054c30ad3a1654bbdee246a
---

# Claude Opus 5.5 et GPT-6 Sol sortent le même jour : Anthropic et OpenAI ont du mal à « ralentir » l'IA

## Metadata

- **Source** : https://www.numerama.com/tech/2338131-claude-opus-5-5-et-gpt-6-sol-sortent-le-meme-jour-anthropic-et-openai-ont-du-mal-a-ralentir-lia.html
- **Site** : Numerama
- **Type** : Article de presse
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Published 22 September 2026 by Nicolas Lellouche, this Numerama piece details the near-simultaneous launches of Anthropic's Claude Opus 5.5 and OpenAI's GPT-6 Sol and GPT-6 Luna. It notes that it is rare for the two sector leaders to launch simultaneously, and that they had not done so since 14 March 2023 (GPT-4 and the first Claude). The irony stressed throughout: both companies had called to slow the AI race the previous week, yet they launched new models at the same time. The key surprise is pricing: both firms cut API prices.

**Claude Opus 5.5** is presented as reaching the level of Claude Fable 5.1 on most tasks at a much lower price. On Terminal-Bench 4.0 (command-line programming) it scores 66.4%, versus 57.9% for GPT-6 Astra and 55.8% for Fable 5.1. On GDPval-AA (professional tasks across 44 occupations) it reaches 1,846 Elo, versus 1,735 for Fable 5.1 and 1,542 for Astra. GPT-6 Astra keeps the advantage on two tests: enterprise task automation (41.4% vs 40%) and scientific research (64.6% vs 58.7%). Opus costs 2.5× less than Fable yet surpasses it on most tests, which raises questions about the relevance of Anthropic's high-end model until Fable 5.5. Anthropic itself admits the real gap is smaller than the numbers suggest. Pricing: $4/$20 per million tokens (vs $5/$25 for Opus 5). Pro, Max and Team subscribers get raised 5-hour usage limits plus a free, storable limit reset (like OpenAI). Notably, Claude's effort selector now defaults to "Medium" instead of "High," and reasoning mode can no longer be disabled — part of the claimed savings stems from this, while benchmarks run at maximum effort. Opus 5.5 inherits Fable 5.1 safeguards: cybersecurity requests are redirected to Opus 4.8, biology requests to Opus 5.

A major point: **Opus 5.5 is the first Opus launched with content watermarking**, as required by the EU AI Act since 2 August 2026. Anthropic applies watermarking worldwide (unable to limit it by region). The detection tool is reserved for regulators, media, fact-checkers and researchers, and Anthropic warns it is not infallible proof.

**GPT-6 Sol and Luna** complete the GPT-6 family three weeks after Astra. Sol is the high/mid-range model; Luna is entry-level; Astra remains the most powerful. No Terra version for now (unlike GPT-5.6). Prices halve: Sol $2/$10, Luna $0.10/$0.50. On AutomationBench, Sol (xhigh) scores 33.2% at $0.27/task, ahead of Fable 5.1 (31.4%) and Opus 5 (26.9%, which costs 11× more per task). Anthropic reports Opus 5.5 at 40% on the same test — still ahead but pricier. On DeepSWE coding, Sol scores 68.8%, 1.1 points behind Fable 5 for ~80% less cost; Luna reaches 66.6%, level with Opus 5 at medium effort. OpenAI claims about half the factual errors of GPT-5.6 Sol per an internal evaluation. OpenAI also takes a jab: Fable 5.1's cost is underestimated because it excludes fallbacks to Opus 5, occurring on ~40% of tasks in that test — a mechanism Opus 5.5 now also uses.

Sol and Luna arrive the same day in ChatGPT Work and Codex for Plus, Pro, Business, Enterprise and Edu subscribers; Free and Go users get Luna in the desktop app. The day before, xAI (now part of SpaceX) launched **Grok 4.7** at $2/$6, one notch below Anthropic/OpenAI; Musk himself placed it near Claude Opus 5. Google is described as absent: Gemini 3.5 Pro was promised for June but poor performance may doom it; the company now talks about Gemini 4 in training.

## Key points

- Anthropic and OpenAI launched on the same day (22 Sept 2026); first simultaneous launch since 14 March 2023.
- Claude Opus 5.5 matches Fable 5.1 on most tasks and tops 7/9 benchmarks; beats GPT-6 Astra on 4/6 shared tests.
- Opus 5.5 pricing $4/$20 vs Opus 5's $5/$25; ~40% savings claimed; effort defaults to "Medium."
- Opus 5.5 is Anthropic's first Opus with EU AI Act content watermarking (applied worldwide).
- GPT-6 Sol ($2/$10) targets code/agents; GPT-6 Luna ($0.10/$0.50) targets high-volume tasks.
- Sol scores 33.2% on AutomationBench at $0.27/task vs Fable 5.1 (31.4%) and Opus 5 (26.9%).
- OpenAI alleges Fable 5.1 cost is understated due to Opus 5 fallbacks on ~40% of tasks.
- Grok 4.7 launched the day before at $2/$6; Google's Gemini 3.5 Pro may never ship, focus shifts to Gemini 4.

## Technical data / figures

**API pricing (per million tokens):**

| Model | Input | Output |
|---|---|---|
| GPT-6 Luna | $0.10 | $0.50 |
| GPT-6 Sol | $2 | $10 |
| Claude Opus 5.5 | $4 | $20 |
| GPT-6 Astra / Claude Fable 5.1 | $10 | $50 |
| Grok 4.7 | $2 | $6 |

**Benchmarks:**

| Benchmark | Claude Opus 5.5 | GPT-6 Astra | Claude Fable 5.1 | GPT-6 Sol | Claude Opus 5 |
|---|---|---|---|---|---|
| Terminal-Bench 4.0 | 66.4% | 57.9% | 55.8% | — | — |
| GDPval-AA (Elo) | 1,846 | 1,542 | 1,735 | — | — |
| Enterprise automation | 40% | 41.4% | — | — | — |
| Scientific research | 58.7% | 64.6% | — | — | — |
| AutomationBench | 40% (per Anthropic) | — | 31.4% | 33.2% (xhigh) | 26.9% |
| DeepSWE | — | — | (Sol 1.1 pts behind) | 68.8% | — |
| DeepSWE (Luna) | — | — | — | — | 66.6% (level with Opus 5 medium) |

- **EU AI Act watermarking obligation**: effective 2 August 2026.
- **AutomationBench cost/task**: Sol $0.27 vs Opus 5 (11× more).

## Why this source matters for the RAG

This is the most detailed French press account of the 22 September 2026 double launch, with granular benchmark scores, pricing tables, and the watermarking/regulatory angle. It is essential for questions on model comparison, API costs, EU AI Act compliance, and competitive dynamics among Anthropic, OpenAI, xAI and Google.
