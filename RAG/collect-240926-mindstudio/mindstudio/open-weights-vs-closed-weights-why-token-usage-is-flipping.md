---
id: collect-240926-mindstudio/mindstudio/open-weights-vs-closed-weights-why-token-usage-is-flipping
title: "open-weights-vs-closed-weights-why-token-usage-is-flipping"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Moonshot", "Nvidia", "OpenAI", "Perplexity", "United States"]
dates: []
keywords: ["open-weight", "agent", "agents", "benchmark", "benchmarks", "claude", "compute", "cost", "cost per token", "deepseek", "fine-tuning", "inference"]
source: docs/RAG/clean_en/mindstudio/open-weights-vs-closed-weights-why-token-usage-is-flipping.md
source_anchor: ""
source_lines: [1, 77]
sha256: cea814a4855ebc460d53358776957b784d69a9b61366a6d1a271729aa723e8f7
---

# open-weights-vs-closed-weights-why-token-usage-is-flipping

<!-- source: https://www.mindstudio.ai/blog/open-vs-closed-weight-model-token-share -->

## What’s actually happening in the token-share data?

Data from Vercel’s hosting platform, tracked from June to August, shows a clear crossover: open-weight models (freely downloadable, self-hostable, fine-tunable) are now processing more total tokens than closed-weight models like OpenAI’s GPT line and Anthropic’s Claude. The line for closed models has been sliding down while open-weight usage climbs. DeepSeek alone has passed Anthropic in token share, 25.2% versus 24.5%. But token volume and revenue are telling two different stories. Anthropic and OpenAI still pull in the overwhelming majority of dollars spent on AI inference, even as their share of raw token traffic shrinks.

## TL;DR

- **Token share has flipped** toward open-weight models, with DeepSeek now edging out Anthropic in percentage of tokens processed on Vercel’s platform.
- **Revenue concentration hasn’t budged** the same way: Anthropic captures roughly 64.6% of model spend versus DeepSeek’s 2.8%, even though their token shares are close.
- **Frontier intelligence carries a massive premium** because the gap between “very good” and “best available” translates into billions of dollars for use cases like trading, legal work, or coding at the edge of what’s possible.
- **Chinese open-weight models are now core infrastructure** for major US companies, with Harvey, Airbnb, Perplexity, and Cursor all building on models like Qwen, Kimi, or DeepSeek.
- **Cost per token isn’t the same as cost per task** , since some cheaper models need far more tokens to reach the same answer, narrowing the real-world savings.
- **Nvidia benefits either way** , because open-weight adoption still drives chip demand, an open-source token costs roughly the same compute to produce as a frontier one.
- **The market is splitting into three tiers** : cheap commodity open models carrying most volume, open-weight “state of the art specialists” capturing meaningful spend, and closed frontier generalists taking a small slice of volume but most of the revenue.

## Remy doesn't write the code. It manages the agents who do.

Remy runs the project. The specialists do the work. You work with the PM, not the implementers.

## Why is token usage shifting to open-weight models?

Cost and control are the two biggest drivers. Open-weight models are generally far cheaper to run than closed frontier models, and businesses can self-host them, fine-tune them on private data, and avoid sending sensitive information to a third party. China has been a major source of these models, releasing increasingly capable open-weight systems that anyone can download and adapt. As the total pool of AI tokens used across the industry keeps growing, a larger share of that growth is landing on open models rather than the closed APIs from OpenAI and Anthropic.

There’s also a structural reason: once a task doesn’t require the absolute best available model, there’s little incentive to pay premium prices. If a flash-tier open model can handle the bulk of everyday tasks, most users have no reason to reach for the most expensive option.

## Why does closed-weight revenue stay so concentrated?

Because the value of being marginally smarter than the competition is enormous in the tasks where it matters. Venture investor Gavin Baker’s framing, cited in the source discussion, is that frontier tokens from labs like Anthropic and OpenAI may represent only 10 to 25% of total tokens used, yet account for 60 to 90% of all economic value generated. The reasoning: for high-stakes use cases like high-frequency trading, legal analysis, or advanced coding, a small accuracy edge is worth paying dramatically more for. Anthropic’s annualized revenue run rate is reportedly above $65 billion, with OpenAI around $40 billion, more than every open-model provider combined.

The pricing gap illustrates this starkly. Claude’s higher-tier models can run around $50 per million output tokens, while a lighter open-weight model like DeepSeek’s flash-tier variant can cost closer to 18 cents per million output tokens. For most tasks, that cheaper model is more than adequate. For the tasks where it isn’t, the premium is easy to justify because the cost of a wrong answer outweighs the token bill.

## Is cost per token the right way to compare models?

Not fully. Investor Bindu Reddy’s point, raised in industry discussion, is that the real metric should be cost per completed task, not cost per token. Some models need significantly more tokens to reach the same conclusion as a competitor. In one comparison, Kimi K3 priced at roughly half the per-token cost of GPT 5.1 still ended up costing almost the same per completed task (84 cents versus 96 cents), because it required more tokens to get there. That gap matters for anyone budgeting AI spend based on sticker price rather than actual task completion cost.

## What does this mean for companies choosing between open and closed models?

## Remy doesn't build the plumbing. It inherits it.

Other agents wire up auth, databases, models, and integrations from scratch every time you ask them to build something.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

Model selection increasingly comes down to more than raw benchmark scores. Investor Martin Casado’s framing highlights cost, privacy, and product fit as the real decision drivers. Privacy in particular is pushing large, well-known companies toward open-weight models they can control end to end. Examples include Thomson Reuters’ Harvey (built on Kimi), Airbnb (using Qwen), Perplexity (using DeepSeek), and Cursor (using Kimi K2.5). These aren’t experimental startups, they’re major businesses choosing to build core products on Chinese open-weight models rather than closed US APIs.

Harvey’s own published results show why: after fine-tuning Kimi K3 into a legal-specific model, it performed competitively on legal benchmarks, ranking near the top on tasks like contract analysis and corporate law agent benchmarks. Fine-tuning on proprietary legal data, something only possible with open weights, let Harvey push a general-purpose open model into specialist territory without depending on a closed vendor to eventually offer the same capability.

## What are the real risks of relying on closed-weight platforms?

The core risk is platform dependency. When a business builds entirely on a closed model API, it’s exposed on several fronts. Data handed to the provider during normal usage can reveal a lot about how a company operates, not just customer information but internal workflows and strategy. If the provider ever decided to compete directly, it would already have a detailed view of the business it’s competing with. There’s also pricing risk (a provider can raise costs unilaterally) and availability risk (a provider can change terms, restrict access, or shut down a service). Open-weight models remove most of this risk: the business owns the weights, owns the fine-tuned data, and can move between dozens of inference providers or self-host, giving it real negotiating power instead of dependency on one or two vendors.

## How is the AI token economy likely to split going forward?

One useful framing, attributed to MIT economist Christian Catalini, splits the market into three tiers. At the bottom are cheap, generalist open-weight models, commodity-priced, carrying the most volume but a small share of total spend. In the middle are open-weight “state of the art specialist” models, fine-tuned for enterprise use cases, which carry meaningful spend despite being open. At the top sit closed, frontier generalist models from labs like OpenAI and Anthropic, a small share of total volume but the majority of revenue. That structure lines up with what the token-share and spend-share data already show: broad usage spreading toward open models while premium spend stays locked with a handful of closed-model providers.

## Frequently Asked Questions

### What’s the difference between open-weight and closed-weight AI models?

Open-weight models release their trained parameters publicly, so anyone can download, run, and fine-tune them on their own infrastructure. Closed-weight models like GPT and Claude are only accessible through a provider’s API, with the underlying weights kept private.

### Why is DeepSeek gaining token share on Anthropic?

DeepSeek’s models are priced far lower per token than Anthropic’s, and being open weight, they can run on many different providers or be self-hosted, making them attractive for high-volume, cost-sensitive workloads even as Anthropic retains most premium, high-value use cases.

### Does more token usage mean more revenue for open-weight providers?

No. Token volume and revenue are decoupled in this market. Open-weight models handle a growing share of total tokens, but closed frontier models still capture the large majority of total dollars spent, because premium tasks command premium prices.

### Why are US companies building products on Chinese open-weight models?

Cost is part of it, but privacy and control matter just as much. Open weights let companies fine-tune models on proprietary data without sending sensitive information to a third-party API, and they avoid depending entirely on a single closed-model vendor.

### Is cost per token a reliable way to compare AI models?

Not on its own. Some models require more tokens to complete the same task, so a lower per-token price doesn’t always translate into a lower total cost. Comparing cost per completed task gives a more accurate picture than raw token pricing.
