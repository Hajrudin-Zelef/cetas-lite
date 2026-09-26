---
id: collect-240926-datacamp/datacamp/gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout-1
title: "gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["luna", "sol", "agent", "agents", "alignment", "astra", "benchmark", "benchmarks", "claude", "cost", "fable 5", "gpt-6"]
source: docs/RAG/clean_en/datacamp/gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout.md
source_anchor: ""
source_lines: [1, 130]
sha256: cc63afb7db69d6f9d79c1fdadc8e5e37b222e291cb2ffe2c5ca2384c71cfe69f
---

# gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout

<!-- source: https://www.datacamp.com/fr/blog/gpt-6-sol-and-luna -->

Tutorial

OpenAI has just expanded the GPT‑6 family with two new models: GPT‑6 Sol and GPT‑6 Luna, both positioned below the flagship GPT‑6 Astra, which we discussed earlier this month. Sol and Luna both benefit from significant price cuts (50% compared to their GPT‑5.6 rates), with the added bonus of gains in programming and workflow automation.

\n
Whether by coincidence of the calendar or not, Anthropic released Claude Opus 5.5 on the same day (today), with a similar positioning: the same training approach as the flagship model, optimized for cost and speed.

\n
To learn more about the reference model, see our GPT‑6 Astra API tutorial and our GPT‑6 Astra vs Claude Fable 5.1 comparison.

\n
## In brief

\n
- \n
- **GPT‑6 Sol and Luna are the two tiers below Astra**, trained with the same recipe and offered at half the price of their GPT‑5.6 predecessors. \n
- **The main gain concerns cost per task for agents and code**, more than raw capability. Almost all results reported by OpenAI are adjusted for cost. \n
- **GPT‑6 Sol makes about half as many factual errors as GPT‑5.6 Sol**; Luna also improves, but remains behind Sol in reliability. \n
- **Favor Sol for agent workflows and code, and Luna for high volumes of repetitive tasks**. Choose Astra only if the task justifies its price. \n
- **If you already use GPT‑5.6 Sol or Luna, switch for the price**; in our hands-on test, the capability gain was modest and manifested as honesty when faced with a faulty input. \n

## What are GPT‑6 Sol and Luna?

\n
Astra remains OpenAI's most capable and best-aligned model, reserved for the most demanding work. Sol and Luna aim to derive the essence of that same training approach, and much of the performance achieved in programming, computer use, etc., into more affordable and faster tiers.

\n
It can be likened to the same relationship Opus 5.5 has with Fable 5.1: not a new frontier, but near-state-of-the-art performance at a fraction of the cost.

\n
## Key features of GPT‑6 Sol and Luna

\n
A few points to remember:

\n
### Genuinely cheaper, across the board

\n
Sol's API price is halved. Luna's is even slightly less than half. You'll find more details in the pricing section.

\n
### Solid results on business workflows

\n
On AutomationBench, which evaluates agents on tasks in sales, marketing, operations, support, finance, etc., Sol at the highest effort level surpasses Claude Opus 5, at a tiny fraction of its cost per task. Luna improves markedly over its GPT‑5.6 predecessor while becoming significantly cheaper per task. OpenAI included a chart that I've reconstructed here:

\n
I noticed that the chart did not include the Opus 5.5 figures published by Anthropic in their Opus 5.5 announcement, so I added a new series. We can see that the latest versions of Sol and Luna remain the most efficient. Probably, if OpenAI didn't include it in its launch announcement, it's purely a matter of timing: Opus 5.5 came out an hour earlier.

\n
### Code gains that follow cost, not just accuracy

\n
On FrontierCode, Sol is described as roughly matching Claude Fable 5.1's best score, but at a much lower price. On another code benchmark (DeepSWE), Sol approaches Fable 5's best score with a steep discount on cost, and Luna is presented as comparable to Opus 5 and Fable 5 at intermediate effort levels.

\n
### Fewer factual errors

\n
OpenAI states that Sol roughly halves the error rate of its predecessor on an internal factuality evaluation based on real conversations in which users had reported errors. Luna also improves; OpenAI claims that at a higher effort level, it matches the factuality of GPT‑5.6 Sol for a fraction of its cost. This point was harder to verify.

\n
### A less verbose writing style

\n
Astra's more concise and less jargon-heavy communication style has been carried over to Sol and Luna. OpenAI has been continuously refining its models' conversational style since 4o was criticized for being sycophantic.

\n
### Smarter and cheaper caching for agents

\n
Beyond the price per token, OpenAI highlights improvements to prompt caching intended to help agents and long conversations reuse context more efficiently, with a discount of about 90% on cached input reads. New dashboards and diagnostics allow developers to identify where the cache is working or not.

\n
### Fewer misleading claims about their own work

\n
Both models are less likely than their GPT‑5.6 counterparts to misrepresent what they did on a coding task.

\n
OpenAI's alignment evaluations, conducted at the maximum effort level, deliberately create situations conducive to dishonesty, and Sol and Luna show lower deception rates than GPT‑5.6 Sol and Luna on tests of coding misinformation, broken research, reviewer circumvention, warning circumvention, and unauthorized interactions.

\n
OpenAI notes that these are adversarial configurations, not failure rates in everyday use, and publishes the full results in the GPT‑6 system card.

\n
## How do GPT‑6 Sol and Luna perform on benchmarks?

\n
OpenAI's communication relies heavily on cost-adjusted comparisons. I mentioned earlier Sol's performance on AutomationBench. OpenAI doesn't just say it's "better than Opus 5": it's better *and* about 11 times cheaper per task.

It should also be noted that, according to OpenAI's own figures, Astra remains ahead of Sol and Luna for computer-use tasks, and that some comparisons with Claude models use different effort levels (e.g., Sol at "xhigh" versus Opus 5 at "medium"), which really strengthens the efficiency arguments but makes raw capability comparisons trickier.

\n
With those nuances in mind, here is what OpenAI reports, grouped by the type of work each benchmark represents.

\n
### Business workflows and agents

\n
Sol's best result is on AutomationBench, Zapier's test for agents working end-to-end across 47 tools in sales, marketing, operations, support, finance, and HR. GPT‑6 Sol at the xhigh effort level scores 33.2% for $0.27 per task, ahead of Claude Opus 5 at maximum effort and even ahead of GPT‑6 Astra at low effort, for a fraction of Opus 5's cost per task.

\n
The Fable 5.1 line should be read with caution. OpenAI lists Claude Fable 5.1 with Opus 5 fallback just below Sol, but the fallback triggered on about 40% of tasks, and its cost is not included in the reported figure.

\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n
| Model (and effort) | AutomationBench score | Cost per task | 
|---|---|---|
| GPT‑6 Sol (xhigh) | 33.2% | $0.27 | 
| GPT‑6 Astra (low) | 30.3% | 3.9× GPT‑6 Sol | 
| Claude Fable 5.1 with Opus 5 fallback (max) | 31.4% | More than 8.9× GPT‑6 Sol (fallback cost not reported) | 
| Claude Opus 5 (max) | 26.9% | 11.1× GPT‑6 Sol | 

For Luna, the same test tells a story of generational progress. At high effort, GPT‑6 Luna gains 5.4 percentage points over GPT‑5.6 Luna, with cost per task down 58%.

\n
On Agents' Last Exam, which evaluates agents on long-horizon professional workflows across 55 sub-sectors, GPT‑6 Sol at maximum effort reaches 56.4%. OpenAI states this is above Claude Opus 5's best score in the evaluation, at a cost per task 60% lower.

\n
### Programming on real code repositories

\n
On DeepSWE v1.1, which evaluates agents on long-horizon software engineering tasks in real repositories, GPT‑6 Sol at maximum effort scores 68.8%. Claude Fable 5's best score in the evaluation is 69.9% at xhigh effort, so Sol sits 1.1 points behind, at a cost per task roughly 80% lower.

\n
GPT‑6 Luna at maximum effort reaches 66.6% on the same benchmark, which OpenAI presents as comparable to Claude Opus 5 and Fable 5 at medium effort. In these comparisons, Luna costs 93% less per task than Opus 5 and 96% less than Fable 5.

