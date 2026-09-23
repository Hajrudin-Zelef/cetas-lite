---
id: vague2-datacamp/datacamp/gpt-6-sol-and-luna
title: "GPT-6 Sol et Luna : rendre les capacités d'Astra accessibles"
domain: datacamp
role: reference
task: article
actors: ["AWS", "Anthropic", "Microsoft", "OpenAI", "OpenRouter"]
dates: ["2026-09-23"]
keywords: ["astra", "gpt-6", "luna", "sol", "agent", "agents", "alignment", "aws", "bedrock", "benchmark", "benchmarks", "chatgpt"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/gpt-6-sol-and-luna.md
source_anchor: ""
source_lines: [1, 56]
sha256: 673ebe34329f8ee7c7ed88c59aa490cc7d19eda0dcbf9b71fe27e82a1d3a5852
---

# GPT-6 Sol et Luna : rendre les capacités d'Astra accessibles

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/gpt-6-sol-and-luna
- **Site** : DataCamp
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

OpenAI expanded the GPT-6 family with two models, GPT-6 Sol and GPT-6 Luna, both positioned below the flagship GPT-6 Astra and trained with the same recipe at roughly half the price of their GPT-5.6 predecessors. Anthropic released Claude Opus 5.5 the same day with a similar cost-and-speed positioning. The article frames Sol and Luna as deriving most of Astra's frontier capability (coding, computer use, etc.) at more affordable and faster tiers—analogous to how Opus 5.5 relates to Fable 5.1.

Key features include genuinely lower prices across the board (Sol's API price halved; Luna slightly below half), strong business-workflow results on AutomationBench, coding gains aligned to cost rather than raw accuracy, fewer factual errors (Sol ~2x fewer than GPT-5.6 Sol), a less verbose writing style inherited from Astra, smarter/cheaper prompt caching with ~90% discount on cached input reads plus new cache dashboards, and reduced misleading claims about their own work (lower deception rates on adversarial alignment evals).

Benchmarks are heavily cost-adjusted. On AutomationBench (Zapier's end-to-end agent test across 47 tools in sales, marketing, ops, support, finance, HR), GPT-6 Sol at xhigh scores 33.2% for $0.27/task, beating Claude Opus 5 at max effort (26.9%, 11.1x the cost) and even GPT-6 Astra at low effort (30.3%, 3.9x cost). On Agents' Last Exam, Sol at max reaches 56.4%, above Claude Opus 5's best at 60% lower cost per task. On DeepSWE v1.1, Sol at max scores 68.8% vs Claude Fable 5's 69.9% at xhigh (1.1 points behind, ~80% lower cost per task); Luna at max scores 66.6%. Independent data is still modest: Artificial Analysis puts Sol at 57 on its Coding Agent Index vs 55 for GPT-5.6 Sol, and 48 vs 47 on its Intelligence Index, with cost per task dropping from $1.99 to $1.06. On factual reliability, OpenAI's internal eval shows Sol ~2x fewer errors than GPT-5.6 Sol; Artificial Analysis's AA-Omniscience shows Sol's hallucination rate dropping from 92% to 60%, though it answered only 83% of questions vs 99% for its predecessor. On computer use, Astra remains OpenAI's best; GPT-6 Sol at xhigh reaches 60.5% on offline OSWorld 2.0 vs 60.3% for Claude Opus 5 at medium effort at ~80% lower cost.

The article includes a hands-on single-file Dijkstra visualizer test across four models, where all performed identically except on a deliberately flawed graph containing a negative edge: only GPT-6 Sol treated it as a stop condition and refused to run, while GPT-5.6 Sol and GPT-5.6 Luna warned but ran anyway, and GPT-6 Luna produced a confusing infinite-negative distance table. Pricing: Sol $2 input/$10 output per 1M tokens (from $4/$20), Luna $0.10/$0.50 (from $0.20/$1.20), with a 90% cached-input discount. Models are available via ChatGPT Work, Codex, the API (`gpt-6-sol`, `gpt-6-luna`), OpenRouter, GitHub Copilot, Azure AI Foundry, and AWS Bedrock.

## Key points

- GPT-6 Sol and Luna are the two tiers below Astra, same training recipe, at half the price of GPT-5.6 predecessors.
- Main gains are cost-per-task for agents and code, not raw capability; almost all OpenAI results are cost-adjusted.
- Sol makes ~2x fewer factual errors than GPT-5.6 Sol; Luna improves but trails Sol on reliability.
- Use Sol for agent workflows and coding; Luna for high-volume repetitive tasks; Astra only when the task justifies its price.
- Sol at xhigh tops AutomationBench (33.2%, $0.27/task), beating Claude Opus 5 at 11.1x the cost.
- On DeepSWE v1.1, Sol (68.8%) is 1.1 points behind Fable 5 (69.9%) at ~80% lower cost per task.
- Caching improvements: ~90% discount on cached input reads, new cache dashboards.
- Hands-on test: only GPT-6 Sol refused to run Dijkstra on a negative-edge graph; others warned but ran anyway.

## Technical data / figures

| Per 1M tokens | GPT-6 Sol | GPT-5.6 Sol | GPT-6 Luna | GPT-5.6 Luna |
| --- | --- | --- | --- | --- |
| Input | $2 | $4 | $0.10 | $0.20 |
| Output | $10 | $20 | $0.50 | $1.20 |

| Model (effort) | AutomationBench score | Cost per task |
| --- | --- | --- |
| GPT-6 Sol (xhigh) | 33.2% | $0.27 |
| GPT-6 Astra (low) | 30.3% | 3.9x GPT-6 Sol |
| Claude Fable 5.1 w/ Opus 5 fallback (max) | 31.4% | >8.9x GPT-6 Sol |
| Claude Opus 5 (max) | 26.9% | 11.1x GPT-6 Sol |

- Reasoning effort scale: `none`, `low`, `medium`, `high`, `xhigh`, `max` (Sol/Luna); Astra starts at `low`.
- API IDs: `gpt-6-sol`, `gpt-6-luna`; OpenRouter: `openai/gpt-6-sol`, `openai/gpt-6-luna`.
- Luna measured at 51,000 output tokens/task vs 41,000 for GPT-5.6 Luna (Artificial Analysis).
- DeepSWE v1.1: Sol (max) 68.8%, Luna (max) 66.6%, Fable 5 (xhigh) 69.9%.
- Agents' Last Exam: Sol (max) 56.4%; OSWorld 2.0 offline: Sol (xhigh) 60.5%.
- AA Coding Agent Index: Sol 57 vs GPT-5.6 Sol 55; AA cost per task $1.99 → $1.06.

## Why this source matters for the RAG

It documents the cost-optimized GPT-6 sub-tiers and the cost-adjusted benchmark methodology that dominates 2026 model releases, useful for advising on model selection by budget and workload. It also provides a practical reliability case study (negative-edge graph handling) illustrating the honesty/alignment improvements claimed by OpenAI.
