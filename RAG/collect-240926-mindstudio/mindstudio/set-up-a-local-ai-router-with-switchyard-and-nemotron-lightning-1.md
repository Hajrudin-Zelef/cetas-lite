---
id: collect-240926-mindstudio/mindstudio/set-up-a-local-ai-router-with-switchyard-and-nemotron-lightning-1
title: "set-up-a-local-ai-router-with-switchyard-and-nemotron-lightning"
domain: mindstudio
role: reference
task: reference
actors: ["Moonshot", "Nvidia"]
dates: []
keywords: ["agentic", "context window", "cost", "kimi", "latency", "moe", "nvidia", "pricing", "reasoning", "speculative decoding", "throughput", "tokens per second"]
source: docs/RAG/clean_en/mindstudio/set-up-a-local-ai-router-with-switchyard-and-nemotron-lightning.md
source_anchor: ""
source_lines: [1, 75]
sha256: 5c759ea9cd9686797019ca1a2a3764b6a938de50b57871a3cee55476ad424448
---

# set-up-a-local-ai-router-with-switchyard-and-nemotron-lightning

<!-- source: https://www.mindstudio.ai/blog/setup-switchyard-nemotron-local-router -->

## What is a local AI router, and why would you build one?

A local AI router is a piece of infrastructure that sits between your application and multiple language models, deciding in real time which model should handle a given request. Instead of sending every task to the same frontier model, the router looks at signals like task complexity, latency requirements, or data sensitivity, and picks the cheapest or fastest model that can still do the job. NVIDIA’s SwitchYard is an open-source router built on top of RouteLLM that adds a classifier, a proxy layer, and several routing strategies, and it can run entirely on your own hardware alongside a local model like Nemotron 3.5 Lightning.

## TL;DR

- **SwitchYard** is NVIDIA’s open-source routing layer, built on RouteLLM, that adds a classifier, a proxy for request plumbing, and support for tracking conversation state across sessions.
- **Nemotron 3.5 Lightning** is a 30 billion parameter mixture-of-experts model from NVIDIA designed for high throughput rather than raw intelligence, using a “latent MoE” approach that projects tokens into a smaller latent space before routing them to active experts.
- Routing exists for four main reasons: **cost** ,**speed** ,**specialization** , and**privacy** , and each reason comes with different tradeoffs for how a routing strategy should be configured.
- SwitchYard ships with four routing strategies: **random** (fixed traffic split for A/B testing),**LLM classifier** (a model decides the tier),**stage router** (reads existing conversation signals for free), and**escalation** (starts cheap and upgrades on trouble).
- On DGX Spark hardware, Nemotron 3.5 Lightning ran a reasoning task in about 20 seconds locally, compared to roughly 50 seconds for Kimi K3 over an API, on a comparable workload.
- Combining a workhorse model like Nemotron Lightning with a frontier model like Kimi K3 through SwitchYard can approximate the completion quality of an expensive frontier-only setup at a fraction of the cost.
- Routing decisions should happen once per task or conversation, not per turn, because per-turn routing breaks caching and adds unnecessary classifier overhead.

## Remy is new. The platform isn't.

Remy is the latest expression of years of platform work. Not a hastily wrapped LLM.

## Why do agentic systems need a router at all?

Most people default to using the most capable model available for every task, especially when paying a flat subscription instead of per-token API pricing. That habit gets expensive fast once you switch to API billing, because even trivial tasks like writing a commit message get routed through a frontier-level model by default.

A router fixes this by matching task complexity to model capability. There are four distinct motivations for doing this, and each has its own constraints:

- **Cost** : token spend adds up quickly at scale, and many workflows don’t need frontier intelligence for every step.
- **Speed** : agentic workflows that take hundreds of steps can’t afford to wait on a slow model at each one, even if that model is smarter.
- **Specialization** : the best all-around model isn’t necessarily the best model for a specific task, and many organizations use fine-tuned models built for narrow jobs.
- **Privacy** : sensitive data may need to stay on local infrastructure, with only sanitized summaries sent to a hosted API.

A single routing layer rarely satisfies all four goals equally well, which is why SwitchYard supports multiple strategies rather than one fixed approach.

## What is Nemotron 3.5 Lightning, and why pair it with a router?

Nemotron 3.5 Lightning is part of NVIDIA’s Nemotron model family, built specifically for high-throughput workloads rather than maximum reasoning depth. It’s a 30 billion parameter mixture-of-experts model with 30 total experts, and architecturally it’s a hybrid of Mamba and transformer components. That hybrid design helps it preserve a 1 million token context window more efficiently than a pure transformer would.

The model uses what NVIDIA calls “latent MoE”: incoming tokens get projected into a smaller latent space before being routed to one of the active experts, which the model’s creators claim improves accuracy per byte processed. Lightning also supports speculative decoding techniques including MTP, along with Dlash and DSpark, which let it process multiple candidate tokens per pass instead of one at a time.

On an artificial intelligence capability index, Lightning lands on par with GPT-OSS, but at meaningfully higher throughput. Running locally on a DGX Spark, it produced about 71 tokens per second on a single stream, with multi-token prediction accepting roughly 70% of its speculative guesses. In a head-to-head test on a 1,400-token reasoning task, the local Lightning model finished in about 20 seconds, while Kimi K3 over an API took roughly 50 seconds for a comparable task. Lightning isn’t trying to out-think a frontier model here. It’s built for jobs where speed and local execution matter more than peak intelligence.

## How does SwitchYard’s routing configuration actually work?

SwitchYard is organized around a few core primitives you configure directly:

**Profiles** define the overall routing behavior, such as whether you’re using a stage router or an escalation router.

### Everyone else built a construction worker.

We built the contractor.

    One file at a time.

UI, API, database, deploy.

**Targets** are the individual models available to the router. You give each one a label (for example, “capable” and “efficient”), specify which model to use, and point to its API provider. A local model like Nemotron 3.5 Lightning running on DGX Spark hardware can sit alongside a hosted model like Kimi K3 as two separate targets.

Once profiles and targets are set, your application simply points at the profile you want to use, and SwitchYard handles dispatching requests to the right target based on the strategy you’ve chosen.

## What routing strategies does SwitchYard support?

SwitchYard ships with four strategies, and you choose one per route in the configuration:

**Random** splits traffic across targets on a fixed ratio. It sounds pointless on its own, but it’s useful for A/B testing or establishing a cost/quality baseline before switching to something smarter.

**LLM classifier** uses a second model to read the incoming request and decide which tier should handle it. The classifier takes the request content as input and outputs a tier selection. This adds an extra model call to every request, so it carries a real cost.

**Stage router** skips the extra model call entirely and instead reads signals already present in the conversation, such as whether tool calls are failing or whether code edits are landing successfully. Because these signals are already available, this strategy costs essentially nothing extra to run.

**Escalation** is a variant of the classifier approach. Every task starts on the cheaper model, and a judge model watches the session. When it detects real trouble, it escalates the session to the more capable model and keeps it there for the rest of that session rather than escalating and de-escalating turn by turn.

Cost matters here as much as capability: random and stage routing are essentially free to run, LLM classification adds a call on every request, and escalation only costs one judge call until a session locks into its tier.

## How much can a router actually save on frontier model costs?

