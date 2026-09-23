---
id: vague2-datacamp/datacamp/claude-mythos-5
title: "Claude Mythos 5 : fonctionnalités, benchmarks et capacités"
domain: datacamp
role: reference
task: article
actors: ["Anthropic", "Glasswing", "Google", "Stripe", "United States"]
dates: ["2026-04", "2026-06", "2026-07", "2026-09-23"]
keywords: ["benchmark", "benchmarks", "claude", "mythos 5", "agent", "agentic", "cyber", "cybersecurity", "fable 5", "gemini", "guardrails", "memory"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/claude-mythos-5.md
source_anchor: ""
source_lines: [1, 59]
sha256: 43a45512fbe460efbf759917c238c4a4a0f7da1e8cedb4497c0c419ee489dba5
---

# Claude Mythos 5 : fonctionnalités, benchmarks et capacités

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/claude-mythos-5
- **Site** : DataCamp
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article covers Claude Mythos 5, Anthropic's most capable model, positioned above the Opus class in what Anthropic calls the "Mythos" tier. Anthropic launched two models on 9 June 2026: Claude Fable 5 (the public version of the Mythos class with conservative safety guardrails) and Claude Mythos 5 (the same underlying model with guardrails lifted for a restricted group of trusted partners). A 1 July 2026 update notes Fable 5 access was restored worldwide after an export-control order was lifted; Mythos 5 remains restricted to approved Project Glasswing partners.

Mythos 5 and Fable 5 share the same architecture; the difference is guardrails. Fable 5 has classifiers redirecting sensitive cybersecurity and biology requests to Claude Opus 4.8, while Mythos 5 lifts these for vetted partners. The name difference reflects guardrails, not capabilities. Mythos 5 is the second version of the Mythos tier, a direct update to Mythos Preview (April 2026, Project Glasswing, a US government cybersecurity collaboration).

**Capabilities**: autonomous software engineering at scale (Stripe reported it compressed months of engineering into days, executing a 50-million-line Ruby migration in one day; best-in-class on FrontierCode Diamond even at medium effort); security (Project Glasswing partners used Mythos Preview to identify 10,000+ major/critical vulnerabilities); drug design and protein engineering (Anthropic's internal team accelerated drug design ~10x; matched or beat human operators on 14 protein targets, 9 producing promising drug candidates); novel scientific hypotheses (first Anthropic model to regularly produce original hypotheses; scientists preferred its molecular biology hypotheses ~80% of the time; one E. coli protein mechanism hypothesis was independently corroborated); autonomous genomics research (over a week of largely autonomous work assembling single-cell data for millions of cells across 138 animal species, training a custom ML model that beat a recently published Science model while 100x smaller); vision and long-context (93.2% on CharXiv Reasoning with tools; file-based memory tripled its performance gain vs Opus 4.8; reached the final act of Slay the Spire three times as often).

**Benchmarks**: Mythos 5 leads or ties on nearly all benchmarks, with consistent gains over Opus 4.8. Key results vs Mythos Preview / Opus 4.8 / GPT 5.5 / Gemini 3.1 Pro: SWE-Bench Pro 80.3% vs 77.8%/69.2%/58.6%/54.2%; FrontierCode Diamond 29.3% (xhigh) vs 13.4%/5.7%; GDPval-AA 1932 vs 1890/1769/1314; GDP.pdf 29.8% vs 22.5%/24.9%/16.7%; Blueprint-Bench 2 38.6% vs 14.5%/36.2%/26.5%; AutomationBench 17.4% vs 15.5%/12.9%/9.6%; OSWorld-Verified 85.0% vs 85.4%/83.4%/78.7%/76.2%; Legal Agent Benchmark 13.3% vs 10.4%/2.1%/0.0%; HLE no tools 59.0%* vs 56.8%/49.8%/41.4%/44.4%; HLE with tools 64.5%* vs 64.7%/57.9%/52.2%/51.4%; BioMysteryBench hard 46.1%* vs 29.6%/40.0%; Terminal-Bench 2.1 88.0%* vs 82.7%/83.4%/70.7%; ExploitBench 78.0%* vs 69.0%/40.0%/34.0%; HealthBench Professional 66.0%* vs 64.7%/56.9%/51.8%. Asterisked scores differ more because Fable 5's classifiers redirect sensitive queries to Opus 4.8.

**Pricing/availability**: $10 input / $50 output per million tokens (less than half of Mythos Preview's $25/$125); API ID `claude-mythos-5`. Access limited to Project Glasswing partners (cyber guardrails lifted) and a small group of biomedical researchers (bio/chem guardrails lifted, cyber retained). A 30-day data retention policy applies to all Mythos-class traffic (not used for training, deleted after 30 days except for safety monitoring).

## Key points

- Claude Mythos 5 is Anthropic's most capable model, above the Opus class; same architecture as Fable 5, differing only in guardrails.
- Released 9 June 2026 alongside Fable 5; Mythos 5 restricted to approved Project Glasswing partners.
- Leads or ties nearly all benchmarks, with the largest gaps on ExploitBench (78.0% vs 40.0% Opus 4.8) and FrontierCode Diamond.
- SWE-Bench Pro 80.3%; HLE with tools 64.5%; Terminal-Bench 2.1 88.0%.
- Demonstrated drug design (~10x acceleration), novel scientific hypotheses, and autonomous genomics research.
- Pricing $10/$50 per million tokens (less than half of Mythos Preview); API ID `claude-mythos-5`.
- 30-day data retention policy for all Mythos-class traffic.

## Technical data / figures

| Category | Benchmark | Mythos 5 / Fable 5 | Mythos Preview | Opus 4.8 | GPT 5.5 | Gemini 3.1 Pro |
| --- | --- | --- | --- | --- | --- | --- |
| Agentic code | SWE-Bench Pro | 80.3% | 77.8% | 69.2% | 58.6% | 54.2% |
| Agentic code | FrontierCode (Diamond) | 29.3% (xhigh) | — | 13.4% (xhigh) | 5.7% (xhigh) | — |
| Knowledge work | GDPval-AA | 1932 | — | 1890 | 1769 | 1314 |
| Knowledge vision | GDP.pdf | 29.8% (no tools) | — | 22.5% | 24.9% | 16.7% |
| Spatial reasoning | Blueprint-Bench 2 | 38.6% | — | 14.5% | 36.2% | 26.5% |
| Tool use | AutomationBench | 17.4% | — | 15.5% | 12.9% | 9.6% |
| Computer use | OSWorld-Verified | 85.0% | 85.4% | 83.4% | 78.7% | 76.2% |
| Legal | Legal Agent Benchmark | 13.3% | — | 10.4% | 2.1% | 0.0% |
| Reasoning | HLE (no tools) | 59.0%* | 56.8% | 49.8% | 41.4% | 44.4% |
| Reasoning | HLE (with tools) | 64.5%* | 64.7% | 57.9% | 52.2% | 51.4% |
| Biology | BioMysteryBench (hard) | 46.1%* | 29.6% | 40.0% | — | — |
| Agentic code | Terminal-Bench 2.1 | 88.0%* | — | 82.7% | 83.4% | 70.7% |
| Cybersecurity | ExploitBench (Cap%) | 78.0%* | 69.0% | 40.0% | 34.0% | — |
| Health | HealthBench Professional | 66.0%* | 64.7% | 56.9% | 51.8% | — |

- Pricing: $10 input / $50 output per 1M tokens (Mythos Preview: $25/$125).
- API model ID: `claude-mythos-5`.
- Data retention: 30 days for all Mythos-class traffic.

## Why this source matters for the RAG

It documents Anthropic's highest-capability restricted model, its benchmark-leading results, and the guardrail-based distinction from the public Fable 5—useful for frontier-model and AI-safety questions. It also captures the access-control and data-retention nuances relevant to enterprise deployment.
