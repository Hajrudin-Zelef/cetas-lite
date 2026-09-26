---
id: collect-240926-datacamp/datacamp/gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout-2
title: "gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout"
domain: datacamp
role: reference
task: reference
actors: ["AWS", "Anthropic", "Microsoft", "OpenAI", "OpenRouter"]
dates: []
keywords: ["luna", "sol", "agent", "agents", "astra", "aws", "bedrock", "benchmarks", "chatgpt", "claude", "copilot", "cost"]
source: docs/RAG/clean_en/datacamp/gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout.md
source_anchor: ""
source_lines: [131, 246]
sha256: eabd5d8c33c4101e33713cbbaacc139981a6813afce08f5ce4561c46ba367e63
---

# gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout

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

