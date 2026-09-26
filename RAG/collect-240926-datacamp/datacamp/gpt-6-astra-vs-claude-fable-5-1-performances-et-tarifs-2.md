---
id: collect-240926-datacamp/datacamp/gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs-2
title: "gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "Irregular", "OpenAI"]
dates: []
keywords: ["astra", "claude", "gpt-6", "agent", "agents", "alignment", "benchmark", "cost", "cybersecurity", "exploit", "fable 5", "gpt-5.6"]
source: docs/RAG/clean_en/datacamp/gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs.md
source_anchor: ""
source_lines: [92, 177]
sha256: a547fa106f101def98200aa0688d742cb509fec6a9b7f782afeb79f47796882e
---

# gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs

### PC use and professional deliverables

Astra dominates this axis, especially because Anthropic hasn't published comparable figures. OpenAI reports 72.6% for Astra on the offline OSWorld 2.0 set against 70.2% for Claude Opus 5, and 92.7% on ScreenSpot-Pro against 87.3% for Claude Fable 5. The figures specific to Fable 5.1 (77.9% partial, 41.7% strict) come from a different version of the tasks and a different scoring rubric.

The artifact figures are less ambiguous: 95.9% on BenchCAD against 84.3%, and 41.4% against 31.4% on AutomationBench.

For agents that click through real software and produce slides or CAD, Astra is the better choice.

### Safety, cybersecurity, and alignment

Astra's alignment figures are the most underrated aspect of its launch. On OpenAI's internal benchmark for safe PC use, where a lower score is better, it reaches 2.4% against 9.5% for Fable 5.1.

In cybersecurity, it changes divisions: 100% on ExploitBench, and 86 FrontierCyber challenges solved out of 226 against 34 for GPT-5.6 Sol according to Irregular's independent tests. Anthropic unrestrained Fable 5.1 just enough to find vulnerabilities but not to exploit them.

Two nuances to keep in mind:

- OpenAI indicates that Astra's written reasoning is *harder* to monitor than Sol's, because it solves in fewer written steps.
- The UK's AISI observed Astra carrying out simulated supply-chain attacks in 2 out of 500 cases, even when the framework prohibited internet access, down from 60 out of 499 when the framework was ambiguous.

Astra is both the safest agent and the most capable attacker, hence the access control.

### Pricing: what you actually pay

The list prices are identical, so on the grid, the only difference between these two models comes from caching and long context. That framing only holds if both consume the same number of tokens for the same task, which isn't the case: hence the importance of the per-task measurements below, which are far more telling than the grid.

#### Token-by-token pricing

| Rate | GPT-6 Astra | Claude Fable 5.1 | 
|---|---|---|
| Input, per 1M tokens | $10.00 | $10.00 | 
| Output, per 1M tokens | $50.00 | $50.00 | 
| Cached input read, per 1M | $1.00 | $0.25 | 
| Cache write (5 min), per 1M | $12.50 | $12.50 | 
| Batch discount | 50% | 50% | 
| Rates beyond 272K input tokens | $20.00 input / $2.00 cache / $75.00 output | No surcharge | 

Input, output, cache writes, and batch discount align to the cent. Cache reads are the exception, with a 4x gap: Anthropic brings Fable 5.1's cache reads down to $0.25, while OpenAI charges $1.00 for Astra's, a 90% discount vs. input.

Beyond 272 K input tokens, the gap jumps to 8x, because OpenAI's surcharge also doubles cache reads to $2.00 at the same time as input, while Anthropic's pricing doc keeps the 1 M token window at the standard rate. Note how low the threshold is relative to Astra's announcement: 272 K is roughly a quarter of its 1.05 M window; any request beyond a quarter of its context is billed at the higher tier.

Cache reads are the rate that compounds over a long loop, because an agent sends back the same system, tool definitions, and repository context on every turn. Our prompt caching guide breaks down the difference between cache writes vs reads, if that's new to you.

#### What a real workload costs under the grid

This whole table is a catalog calculation on a fixed token shape, not a measured result. It answers: "if both models consumed identical tokens, what would the grid bill?" The next section answers the question of what they actually consume.

| Workload | GPT-6 Astra | Claude Fable 5.1 | Difference | 
|---|---|---|---|
| Balanced assistant: 1 M input / 250 K output | $22.50 | $22.50 | $0, 0 % | 
| Research, below threshold: 10 M input / 1 M output | $150 | $150 | $0, 0 % | 
| Research, above threshold: 10 M input / 1 M output | $275 | $150 | $125, 83% more for Astra | 
| Very cache-heavy loop: 100 K prefix, 1,000 reads | $201 | $126 | $75, 59% more for Astra | 

The formula is the same: (volume ÷ 1 M) × rate, summed over fresh input, cache reads and writes, and output. The "cache loop" row assumes a 100 K prefix written once then re-read 1,000 times, plus 5 K fresh input and 1 K output per request.

The "research above threshold" row is the most telling. Keep every request under 272 K input tokens, and both cost exactly $150. Push the same 10 M tokens through requests that each exceed 272 K, and Astra climbs to $275 while Fable 5.1 stays at $150, because Anthropic bills the 1 M window at the standard rate.

The very cache-heavy loop is the one many will encounter. At 1,000 cache reads of a 100 K prefix, the $0.75 per million cached tokens gap becomes $75 on an otherwise identical workload. Note the "shape": deliberately output-poor, 1 K of output per request for 100 K of cache reads. It's the only shape where Fable 5.1's cache advantage decides the bill.

#### What each task actually costs

As soon as you measure real token consumption instead of assuming it, the table flips. Artificial Analysis prices each model per task on the Intelligence Index, and its method already includes input, cache reads/writes, thinking, and response; Fable 5.1's cache advantage is therefore included.

| Effort level | GPT-6 Astra: score/cost per task | Claude Fable 5.1: score/cost per task | 
|---|---|---|
| max | 61 / $1.67 | 66 / $3.76 | 
| xhigh | 61 / $1.20 | 65 / $2.72 | 

At max effort, Fable 5.1 costs 2.25 times Astra for the same batch of tasks at identical list prices. Artificial Analysis reaches the same finding on its coding index, where Astra matches Fable 5 "for less than half the cost, thanks to significant token efficiency gains." Astra's scale drops to $0.46 per task at low effort.

Fable 5.1 isn't penalized by its grid here, but by its output volume. Artificial Analysis measured it at about 1.7x the output tokens of Fable 5 at max effort, hence a cost per task 20% higher than its predecessor despite the cache reduction. Without that reduction, it would be around $5.16: the discount really does help, but not enough.

You pay that surcharge for something. Fable 5.1 scores 5 points higher on the same index, and 4 points higher on xhigh. The question is whether 5 points are worth 2.25x: that's for you to judge, but the idea that identical list prices yield identical bills doesn't hold up against measurements.

OpenAI's claims point the same way, for whatever a vendor's self-reporting is worth: estimated API cost per task about 31% below Fable 5.1 on Terminal-Bench Science 0.1, 63% below on Terminal-Bench 4.0, and 86% below on BenchCAD, measured at the chosen effort settings.

## How GPT-6 Astra and Claude Fable 5.1 performed

Astra narrowly won our test overall, though both models delivered a working simulation on the first try.

### The test

I ran a hard task against both models under identical conditions, rather than several shallow ones. The test chosen here: a from-scratch physics simulation in a single HTML file, chosen because it exercises exactly what both vendors say they improved: sustained correctness over a long build, with no supporting library.

This instance is a refreshed version of the test. The published variant used a rotating square container; here, it's replaced by a rotating hexagon with a counter-rotating central obstacle and mass proportional to size, which increases difficulty while keeping a clear confinement to judge. Here is the prompt, pasted word for word into a new chat for each model:

