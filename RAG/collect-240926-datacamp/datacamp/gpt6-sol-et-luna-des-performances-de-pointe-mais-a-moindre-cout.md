---
id: collect-240926-datacamp/datacamp/gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout
title: "gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout"
domain: datacamp
role: reference
task: reference
actors: ["AWS", "Anthropic", "Microsoft", "OpenAI", "OpenRouter"]
dates: []
keywords: ["luna", "sol", "agent", "agents", "alignment", "astra", "aws", "bedrock", "benchmark", "benchmarks", "chatgpt", "claude"]
source: docs/RAG/clean_en/datacamp/gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout.md
source_anchor: ""
source_lines: [1, 285]
sha256: aa5777209195523a590381b336c6fd7e23530d493a236980c3a0401214b5fd5c
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

\n
Independent data remains modest for now. Artificial Analysis places GPT‑6 Sol at 57 on its Coding Agent Index versus 55 for GPT‑5.6 Sol, and at 48 versus 47 on its Intelligence Index, with an estimated cost per task down from $1.99 to $1.06.

\n
### Factual reliability

\n
OpenAI's internal evaluation is built from de-identified ChatGPT conversations where users had reported a factual error from an earlier model. On this set, GPT‑6 Sol makes about half as many errors as GPT‑5.6 Sol, and GPT‑6 Luna at higher effort matches GPT‑5.6 Sol for about one-hundredth of its cost.

\n
Artificial Analysis's AA-Omniscience test points in the same direction, with one nuance: Sol's hallucination rate dropped from 92% for GPT‑5.6 Sol to 60%, but it only answered 83% of questions versus 99% for its predecessor, so part of the gain comes from refusing to answer. Luna's hallucination rate on the same test was 77%.

\n
### Computer use

\n
Astra remains OpenAI's best model for computer use, and the post confirms it. On offline OSWorld 2.0, GPT‑6 Sol at xhigh effort reaches 60.5% versus 60.3% for Claude Opus 5 at medium effort, at a cost per task roughly 80% lower. GPT‑6 Luna at maximum effort beats GPT‑5.6 Sol at medium effort for one-tenth of its cost.

\n
## Which level should you choose?

\n
Sol is the default choice for most developer and agent work, Luna for high volumes, and Astra for projects where a bad answer costs more than tokens. The GPT‑6 family now has three tiers sharing the same training recipe, with differences in depth, speed, and price.

\n
All three offer the same reasoning scale in the API: `none`, `low`, `medium`, `high`, `xhigh`, and `max` for Sol and Luna, with Astra starting at `low`. The higher tiers consume more tokens to reason, and most of OpenAI's headline results were obtained at xhigh or max.

GPT‑6 also allows changing effort mid-conversation without invalidating the prompt cache: an agent can thus do low-cost follow-ups at low and only escalate the difficult steps.

| Use case | Tier | Why |
|---|---|---|
| Multi-step agents on business applications | Sol | Leads on AutomationBench and Agents' Last Exam for a fraction of the cost per task of competitors |
| Coding agents producing mergeable changes | Sol | Substantial gain on FrontierCode vs GPT‑5.6 Sol; close to Fable 5 on DeepSWE at significantly lower cost |
| Fact-rich writing and research syntheses | Sol | About half as many factual errors as its predecessor |
| Classification, extraction, and routing at scale | Luna | $0.10 input and $0.50 output per million tokens, and parity with GPT‑5.6 Sol on factuality at high effort |
| Desktop use with a Free or Go subscription | Luna | The only GPT‑6 model available on these plans |
| Long-horizon computer use and the hardest end-to-end work | Astra | Still the best at OpenAI, at $10 input and $50 output per million tokens |

**A caveat about Luna:** Artificial Analysis measured 51,000 output tokens per task, versus 41,000 for GPT‑5.6 Luna; on workloads without an output cap, the price drop is therefore smaller than the label suggests.

## Testing GPT‑6 Sol and Luna: practical examples

\nThe benchmarks above come from OpenAI, so I ran a build task on all four models with a strictly identical prompt and tools.

\nThe task: a single-file HTML visualizer of Dijkstra's algorithm, animating an edge examination every 400 ms until the target stabilizes, with distinct node states, pause/step/possibly restart controls, and a final distances table. No build step, no dependencies, no network.

\nThe real test was in the data file. It contains three graphs:

\n
- \n
- A normal graph (all positive weights, reachable target) \n
- A graph where the target is unreachable \n
- A graph containing a negative weight, making Dijkstra inapplicable (infinite loop) \n

The prompt never mentions these traps. Detecting them is the goal, and this tests both the FrontierCode gain vs GPT‑5.6 and the decline in misleading claims about coding work.

\n
All four ran in OpenCode with the same tool surface (file read/edit only), same level of reasoning effort, one attempt, fresh session, prompt identical to the byte.

\n
Here is an example of what GPT-6 Sol produced for the normal graph:

The main findings:

\n
- \n
- **On the healthy graph, no difference.** All four delivered a working file, found the correct path at distance 24, fixed the nodes in order, and designed readable controls fitting on one screen. Zero errors in fidelity and layout. \n
- **The unreachable target didn't trap anyone either.** All four stopped on their own and left the two islands at infinity. \n
- **Scenario 3 makes the whole result.** All the models spotted the negative edge. Only GPT‑6 Sol treated it as a reason to stop: it indicated on screen that an undirected negative edge makes shortest paths undefined, and refused to run. \n

Let's look closer. GPT-6 Sol refused to run the algorithm because of the negative weight and displayed a highly visible red error message.

\n
GPT‑5.6 Sol flagged the negative weight in a warning, then still ran the algorithm, highlighting a path and filling the distances table as if the answer were valid. A warning next to a wrong answer is still a wrong answer. GPT‑5.6 Luna did the same thing.

\n
GPT‑6 Luna landed between the two: a banner naming the edge and declaring Dijkstra inapplicable, then a distances table at negative infinity. Defensible for an undirected negative edge, but confusing to read.

\n\n
So, is GPT‑6 a real step forward? Barely, and on a single point, even if it is important. On the build task itself, the generations are equivalent. The difference lies entirely in how each model handled an input it knew to be defective: this comes back to OpenAI's honesty promise, more than to coding skill. The two Lunas did not stand out.

\n
Another lesson from the test: when four models pass every milestone except one, the bar needs to be raised.

\n
## Pricing and availability of GPT‑6 Sol and Luna

\n
Here is the new pricing grid. Both cuts represent a 50% reduction compared with the respective GPT‑5.6 prices.

\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n
| Per 1M tokens | GPT‑6 Sol | GPT‑5.6 Sol | GPT‑6 Luna | GPT‑5.6 Luna | 
|---|---|---|---|---|
| Input | $2 | $4 | $0.10 | $0.20 | 
| Output | $10 | $20 | $0.50 | $1.20 | 

Cached input token reads get a 90% discount on GPT‑6; on long agent runs, the cache hit rate therefore matters as much as the listed price. OpenAI's launch post does not publish batch rates for these models, and both remain below the $10 (input) and $50 (output) of GPT‑6 Astra.

\n
The only access conditions mentioned concern the subscription and product surface, where the picture becomes clearer.

\n
## How to access GPT‑6 Sol and Luna?

\n
GPT‑6 Sol and Luna are already available on many platforms:

- 
ChatGPT Work and Codex for Plus, Pro, Business, Enterprise, and Edu subscribers, with Luna also accessible to Free and Go users in the desktop app.- 
On the API, they are available as `gpt-6-sol` and `gpt-6-luna`. OpenAI notes that the rollout in ChatGPT is spread over the day, so it may take a while to appear for everyone.- 
In addition, both models are available via OpenRouter (as `openai/gpt-6-sol` and `openai/gpt-6-luna`), in GitHub Copilot, on Azure AI Foundry, and on AWS Bedrock.

On the API, they are available as `gpt-6-sol` and `gpt-6-luna`. OpenAI notes that the rollout in ChatGPT is spread over the day, so it may take a while to appear for everyone. A minimal call looks like this:

`from openai import OpenAI\n\nclient = OpenAI()\nresponse = client.responses.create(\n    model=\"gpt-6-sol\",\n    reasoning={\"effort\": \"high\"},\n    input=\"List the open pull requests that touch billing code and summarize the risk of each.\",\n)\nprint(response.output_text)`

If you are migrating from GPT‑5.6, OpenAI's GPT‑6 model recommendations, written for Astra, advise removing `temperature` and `top_p`, starting at `low` if you were using `none` or `minimal`, otherwise keeping your current effort level, and replacing `prompt_cache_retention` with `prompt_cache_options.ttl` set to 30 minutes.

## To conclude

Sol and Luna take the training recipe from the flagship Astra model to make the lower tier significantly more affordable and faster. OpenAI's communication explicitly cites Anthropic's models and claims gains on cost-adjusted benchmarks. The comparison is relevant, because Opus 5.5 also focuses on efficiency.

My take: if you run agents or coding workloads on GPT‑5.6 Sol or Luna, switch now. Same effort scale, half the price, and in our test, the only thing GPT‑6 Sol did differently from its predecessor was refuse to give a confident answer to a question the algorithm could not solve. If you are on Claude, the cost-adjusted figures argue for running your own comparison rather than taking OpenAI at its word.

If you want to build on these models, we recommend the OpenAI Fundamentals skill track, which covers the API end to end in 15 hours.

I am a writer and editor in the field of data science. I am particularly interested in linear algebra, statistics, R, and so on. I also play a lot of chess!

**Editor-in-Chief, Data Science at DataCamp |** **I am passionate about forecasting and development using APIs.**

## FAQ on GPT‑6 Sol and Luna

### What are GPT‑6 Sol and Luna, and what is their relationship to GPT‑6 Astra?

These are two new variants in the GPT‑6 family, positioned below Astra (OpenAI's top-tier model) in terms of price and capability. They use training methods similar to Astra, but are optimized to reduce costs and improve latency, in order to meet everyday needs rather than the most demanding projects for which Astra is reserved.

### How are they cheaper than the previous generation?

Both models cost 50% less than their GPT‑5.6 promotional rates. Concretely: Sol goes from $4/$20 to $2/$10 per million input/output tokens, and Luna from $0.20/$1.20 to $0.10/$0.50 per million tokens.

### How do they compare with competitors like Claude?

OpenAI claims strong cost-efficiency — for example, on AutomationBench, GPT‑6 Sol at high effort outperforms Claude Opus 5 for a fraction of its cost per task. Similar comparisons are presented on coding benchmarks (FrontierCode, DeepSWE) and computer use (OSWorld), generally highlighting better or comparable scores at a significantly lower cost. Note that these are results reported by OpenAI, and that competitor figures come from public sources rather than tests conducted by OpenAI.

### What improvements have been made to caching and alignment?

The caching improvements aim for higher hit rates by default (with up to a 90% discount on cached input tokens), as well as new tools such as a cache dashboard and reasoning/tooling effort settings that preserve the cache. On alignment, both models show improved honesty indicators (e.g., fewer misleading claims about coding work) compared with their GPT‑5.6 predecessors.

### When and where can I access these models?

They are available now in ChatGPT Work and Codex for Plus, Pro, Business, Enterprise, and Edu subscribers, with Luna also accessible to Free/Go users in the desktop app. Via the API, they are accessible as `gpt-6-sol` and `gpt-6-luna`. They are not yet available in the standard consumer ChatGPT offering, and the rollout is happening gradually throughout the day.

### Should I use GPT‑6 Sol or GPT‑6 Luna?

Use Sol for agent workflows, coding on real repositories, and highly factual professional work; it is the tier on which OpenAI's results on AutomationBench, DeepSWE, and factuality rest. Use Luna for high-volume, cost-sensitive tasks (classification, extraction, routing), where its $0.10 input and $0.50 output per million tokens matter more than peak capability. Luna is also the only GPT‑6 model available to Free and Go users in the ChatGPT desktop app.
