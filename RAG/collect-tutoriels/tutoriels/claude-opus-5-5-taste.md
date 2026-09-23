---
id: collect-tutoriels/tutoriels/claude-opus-5-5-taste
title: "Claude Opus 5.5 vs GPT-6 Astra : un test gustatif qui a dépassé les benchmarks"
domain: tutoriels
role: reference
task: article
actors: ["Anthropic", "OpenAI"]
dates: ["2026-09", "2026-09-23"]
keywords: ["astra", "benchmark", "benchmarks", "claude", "gpt-6", "opus 5", "cost", "fable 5", "gpt-5.6", "opus 4", "parameters", "pricing"]
source: docs/RAG/Collect RAG/07_tutoriels/claude-opus-5-5-taste.md
source_anchor: ""
source_lines: [1, 76]
sha256: fd8db2906362d0c9b367717e6cf59e6d220b813cfc4007c73dad177ad9d9f87d
---

# Claude Opus 5.5 vs GPT-6 Astra : un test gustatif qui a dépassé les benchmarks

## Metadata

- **Source** : https://www.orcarouter.ai/fr/blog/claude-opus-5-5-taste
- **Site** : OrcaRouter
- **Type** : Article
- **Language** : fr
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This OrcaRouter article by Gideon Frost (published 23 September 2026) examines the apparent contradiction in a practitioner's social post about Claude Opus 5.5: someone asked it to make a short video about a rival lab solving the Navier-Stokes equations, published the result, and wrote that the model is "very nice," has "excellent taste," and is the first Claude since Opus 4.7 they actually want to use intensively — yet they still keep GPT-6 Astra and GPT-5.6 Sol as their main engines because their work is heavily research-oriented. The article frames this as three claims pulling in different directions: a model can be the most pleasant to work with without being the one you route production traffic to. The interesting question is not which model is best, but which of the two conclusions survives contact with the numbers, and which turns out to be a claim about taste.

The article is explicitly a practitioner's report, not a benchmark campaign — a signal about how the model feels in creative, open-ended work, not proof of capability. All figures are labelled with who produced them. It explains that a "taste test" can and cannot tell you certain things: the artifact matters. Making a video about a math result is not a binary pass/fail coding task; there is no compile step, no test suite, no Terminal-Bench harness. When a practitioner says a model has "excellent taste," they are describing judgment about scope and emphasis in work where the specification is deliberately vague. Taste is what you want when exploring, not when auditing 200,000 lines of code and justifying a bill.

The article then compares two sets of scores. On raw published benchmarks, Claude Opus 5.5 leads GPT-6 Astra in most cases, but the two most-cited sources disagree on the magnitude. Anthropic's launch table places Claude Opus 5.5 at 66.4% on Terminal-Bench 4.0, with GPT-6 Astra at 57.9% and Claude Fable 5.1 at 55.8% — an 8.5-point vendor-announced lead. Artificial Analysis, running its own harness, measured Claude Opus 5.5 at 59.6% on the same benchmark: tied with GPT-6 Astra at xhigh effort, and about 11 points above Claude Opus 5. The lead is real in both readings; its magnitude is not. The article lists benchmarks where the models swap places: Terminal-Bench 4.0, Humanity's Last Exam with tools, GDPval-AA v2.1 (where Claude Opus 5.5 leads 1,846 vs 1,542 Elo, from independent Artificial Analysis), Terminal-Bench-Science 0.1 (Astra wins 64.6% vs 58.7%), AutomationBench (Astra 41.4% vs 40.0%), and FrontierCode v1.1 (Claude 54.4% vs Astra 53.3%). Research-intensive workloads are where GPT-6 Astra holds up.

A second, less flattering reason for weak margins: cross-vendor comparisons are not run with the same parameters. Claude Opus 5.5's published figures come from xhigh reasoning effort versus GPT-6 Astra at high. Artificial Analysis places both at xhigh, and the coding gap disappears — the 8.5-point lead becomes a tie. The token bill tells the same story: Anthropic's own launch materials indicate Claude Opus 5.5 spends on the order of 119,000 tokens per problem, about 84,000 of them on thinking, while GPT-6 Astra solves comparable problems in about 27,000 tokens. Thinking tokens bill as output tokens. A model using four times the tokens at one-fifth the output price costs about the same.

On pricing, Claude Opus 5.5 is the cheaper model: $4.00/M input, $20.00/M output, $0.20/M cached input, $5.00/M cache writes, 1M context, 128k max output. GPT-6 Astra: $10.00/M input, $50.00/M output, $1.00/M cached input, $12.50/M cache writes; beyond 272,000 input tokens in a single request, the entire request is re-billed at $20.00 input / $100.00 output, with a batch tier at $5.00/$25.00. So Claude Opus 5.5 is 2.5x cheaper on input and output, with cached input five times cheaper. But the token-usage caveat and GPT-6 Astra's long-context pricing are the two traps. The article explains why the research-heavy user did not switch: their workloads are long, open-ended, tool-using, and expensive per run, where Astra dominates science and automation leaderboards and uses a fraction of the thinking tokens. The defensible conclusion is a routing decision, not a conversion.

The article ends by noting GPT-6 Astra is available on OrcaRouter at OpenAI's own list price with 0% markup (changes arrive same-day), while Claude Opus 5.5 is available via Anthropic's own API and third-party platforms (not served by OrcaRouter). It recommends placing both behind one key and endpoint and routing by workload, with automatic failover, and closes with two things to watch: whether the effort-setting asymmetry resolves, and whether Claude Opus 5.5's thinking-token overhead shrinks.

## Key points

- Claude Opus 5.5 is praised for "excellent taste" in open-ended creative work, yet the practitioner keeps GPT-6 Astra as the main engine.
- Vendor vs independent benchmark disagreement: 66.4% vs 59.6% on Terminal-Bench 4.0.
- Effort-setting asymmetry (xhigh vs high) explains much of the claimed 8.5-point lead; at matched xhigh, it becomes a tie.
- Claude Opus 5.5 uses ~119,000 tokens/problem (~84,000 thinking) vs GPT-6 Astra ~27,000 tokens.
- Pricing: Claude Opus 5.5 is 2.5x cheaper on input/output and 5x cheaper on cached input.
- GPT-6 Astra's long-context surcharge kicks in above 272,000 input tokens, re-billing the entire request.
- Conclusion is a routing decision: judgment-bound work to Claude, long research loops to Astra.
- Taste is real capability but is exactly what benchmark axes do not capture.

## Technical data / figures

Benchmarks:

| Benchmark | Claude Opus 5.5 | GPT-6 Astra |
| --- | --- | --- |
| Terminal-Bench 4.0 (Anthropic) | 66.4% | 57.9% |
| Terminal-Bench 4.0 (Artificial Analysis) | 59.6% | 57.9% |
| Humanity's Last Exam with tools | 67.7% (vendor) / 61.4% (independent) | 57.2% |
| GDPval-AA v2.1 (Elo) | 1,846 | 1,542 |
| Terminal-Bench-Science 0.1 | 58.7% | 64.6% |
| AutomationBench | 40.0% | 41.4% |
| FrontierCode v1.1 | 54.4% | 53.3% |
| Claude Fable 5.1 Terminal-Bench 4.0 | — | 55.8% |

Token usage:

| Model | Tokens/problem | Thinking tokens |
| --- | --- | --- |
| Claude Opus 5.5 | ~119,000 | ~84,000 |
| GPT-6 Astra | ~27,000 | — |

Pricing (per million tokens):

| Item | Claude Opus 5.5 | GPT-6 Astra |
| --- | --- | --- |
| Input | $4.00 | $10.00 |
| Output | $20.00 | $50.00 |
| Cached input | $0.20 | $1.00 |
| Cache writes | $5.00 | $12.50 |
| Context | 1M | 272K threshold |
| Max output | 128k | — |
| Batch tier | — | $5.00 / $25.00 |
| Long-context (>272k) | — | $20.00 / $100.00 |

Other figures: Anthropic internal fact-check test reportedly passes 16 of 18 attempts, versus 0 of 18 for Claude Fable 5.1 and Claude Opus 5 (vendor claim, not reproduced).

## Why this source matters for the RAG

It provides a nuanced, source-labelled comparison of two frontier models that separates benchmark claims from practical routing and cost decisions, including effort-setting asymmetry and token-efficiency traps. It is valuable for RAG questions about model selection, benchmark reliability, pricing, and the distinction between measurable capability and subjective "taste."
