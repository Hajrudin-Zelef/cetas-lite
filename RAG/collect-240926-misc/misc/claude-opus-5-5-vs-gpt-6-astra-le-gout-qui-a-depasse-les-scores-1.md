---
id: collect-240926-misc/misc/claude-opus-5-5-vs-gpt-6-astra-le-gout-qui-a-depasse-les-scores-1
title: "claude-opus-5-5-vs-gpt-6-astra-le-gout-qui-a-depasse-les-scores"
domain: orcarouter
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["astra", "claude", "gpt-6", "agentic", "benchmark", "benchmarks", "cybersecurity", "fable 5", "gpt-5.6", "leaderboard", "opus 4", "opus 5"]
source: docs/RAG/clean_en/misc/claude-opus-5-5-vs-gpt-6-astra-le-gout-qui-a-depasse-les-scores.md
source_anchor: ""
source_lines: [1, 56]
sha256: b2f7b7fc652ae0310de301af07e3b9f589f8f41cfcc27d9796717a9df108233b
---

# claude-opus-5-5-vs-gpt-6-astra-le-gout-qui-a-depasse-les-scores

<!-- source: https://www.orcarouter.ai/fr/blog/claude-opus-5-5-taste -->

Someone asked **Claude Opus 5.5** to make a short video about a rival lab solving the Navier-Stokes equations, published the result, and then wrote the sentence that makes this comparison interesting: the model is "really nice," it has "excellent taste," and it's the first Claude since Opus 4.7 that they actually want to use intensively — but they're still keeping **GPT-6 Astra** and GPT-5.6 Sol as their main engines, because their work is heavily research-oriented. That's three claims in a single post, and they pull in different directions. A model can be the most pleasant thing to work with without being the one you route production traffic to. The interesting question isn't which model is best. It's which of these two conclusions holds up against the numbers, and which turns out to be a statement about taste.

This post is a practitioner's report, not a benchmark campaign — treat it as a signal about how the model feels in creative, open-ended work, not as proof of its capabilities. All numerical data below is labeled with an indication of its source.

## What a taste test can and cannot tell you

The artifact matters here. Making a video about a mathematical result is not a programming task with a binary pass/fail criterion. There's no compilation step, no test suite, no Terminal-Bench harness evaluating whether the result is worth anything. The model had to choose an angle, decide what to show, keep the whole thing coherent, and stop. When a practitioner says a model has "excellent taste," that's what they're describing: a judgment about scope and emphasis that manifests in work where the specification is deliberately vague.

That's a real capability, and it's the one benchmark suites measure least well. It's also why the same person can praise a model without switching to it. Taste is what you want when you're exploring. It's not what you want when you have 200,000 lines of code to audit and an invoice to justify.

Anthropic's own launch framing points to the same faculty, in prose rather than video: Claude Opus 5.5 is presented as writing less "Claudish" — fewer preambles, fewer hedges, a greater tendency to state the main idea first. Independent readability evaluations support the direction of this claim, even where they diverge on magnitude. The vendor also reports that an internal fact-checking test passes 16 out of 18 attempts, versus zero out of 18 for Claude Fable 5.1 as well as Claude Opus 5; this figure is Anthropic's own, not reproduced, and should be read as a claim rather than a result.

## Two scoreboards, one disagreement that matters

On the raw published benchmarks, Claude Opus 5.5 leads GPT-6 Astra in the majority of cases. The problem is that the two most-cited sources don't agree on the size of the gap.

Anthropic's launch table places Claude Opus 5.5 at 66.4% on Terminal-Bench 4.0, with GPT-6 Astra at 57.9% and Claude Fable 5.1 at 55.8%. That's an 8.5-point lead announced by the vendor in agentic coding — the kind of margin that settles a debate in a marketing presentation. Artificial Analysis, which runs its own harness, measured Claude Opus 5.5 at 59.6% on the same benchmark: tied with GPT-6 Astra at xhigh effort, and about 11 points above Claude Opus 5. The lead is real in both readings. Its magnitude is not.

Where the two models trade places is more useful than where they don't:

• Terminal-Bench 4.0 (agentic coding) — Claude Opus 5.5 66.4% per Anthropic's own table, 59.6% per Artificial Analysis; GPT-6 Astra 57.9%.

• Humanity's Last Exam, with tools — Claude Opus 5.5: 67.7% per the vendor, 61.4% in independent measurement; GPT-6 Astra: 57.2%.

• GDPval-AA v2.1 (real knowledge work across 44 occupations) — Claude Opus 5.5 1,846 Elo; GPT-6 Astra 1,542 Elo. Both figures come from Artificial Analysis's independent leaderboard, not the vendor.

• Terminal-Bench-Science 0.1 — GPT-6 Astra 64.6% versus Claude Opus 5.5 58.7%. This is a clear win for Astra, and it's the science-flavored agentic benchmark.

• AutomationBench — GPT-6 Astra 41.4% versus Claude Opus 5.5 40.0%. Thin margin, but Astra again.

• FrontierCode v1.1 — Claude Opus 5.5 54.4% vs GPT-6 Astra 53.3%. Within one point.

Read that list the way the practitioner would. Research-heavy workloads — the ones with a scientific or literary component, the ones that run a long tool-use loop and require the model to keep its footing — are exactly where GPT-6 Astra holds its ground. The knowledge-work and coding leaderboards where Claude Opus 5.5 gallops ahead are the ones you'd show if your job were shipping software. The two people in that conversation are each reading their own benchmark correctly. They're just reading different benchmarks.

## The effort setting does more work than the model

There's a second, less flattering reason for the thin displayed margins. Cross-vendor comparisons aren't run with the same parameters.

The published figures for Claude Opus 5.5 come from an extra-high reasoning effort configuration (xhigh), versus GPT-6 Astra at high. Artificial Analysis's independent runs place both models at xhigh, and the coding gap disappears — the 8.5-point lead advertised by the vendor turns into a tie. This is not an accusation of wrongdoing; it is what happens when a vendor chooses the configuration that best showcases its model and an independent lab chooses the configuration that matches the competitor's. Before migrating a workload on the strength of a benchmark table, check which effort setting it was run at, because the answer is often "the flattering one."

The token bill tells the same story from a different angle. Anthropic's own launch materials state that Claude Opus 5.5 spends on the order of 119,000 tokens per problem, of which about 84,000 go to thinking, while GPT-6 Astra solves comparable problems in about 27,000 tokens. Thinking tokens are billed as output tokens. A model that uses four times as many tokens, at one-fifth the output price, comes out roughly even — and that's without counting the fact that the cheaper model is often the one you would have been willing to run at a higher effort setting anyway.

An additional caveat applies specifically to the Claude Opus 5.5 side: Anthropic's safety routing can direct certain cybersecurity- and biology-related requests down a more restrictive path, which lowers the published scores for those evaluations relative to what the underlying model would produce. If your work touches these domains, the gap between the benchmark and your experience will be real, and it will run in the direction of an optimistic benchmark.

## What a real task costs on each

Claude Opus 5.5 is the cheaper model on the price sheet, and the gap is not small:

• Claude Opus 5.5 — $4.00 per million input tokens, $20.00 per million output tokens, $0.20 per million cached input tokens, $5.00 per million cache writes. 1M token context, 128k max output.

• GPT-6 Astra — $10.00 per million input tokens, $50.00 per million output tokens, $1.00 per million cached input tokens, $12.50 per million cache writes. Beyond 272,000 input tokens in a single request, the entire request is rebilled at $20.00 input and $75.00 output; batch pricing is $5.00/$25.00.

So Claude Opus 5.5 is 2.5x cheaper on input and 2.5x cheaper on output, with cached input five times cheaper. If your tasks are similar in tokens, that's the whole decision, and the question of taste is just decoration.

