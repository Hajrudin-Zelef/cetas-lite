---
id: collect-240926-mindstudio/mindstudio/how-openai-and-anthropic-turn-compute-into-profit-1
title: "how-openai-and-anthropic-turn-compute-into-profit"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "Nvidia", "OpenAI", "SpaceX"]
dates: []
keywords: ["compute", "agent", "agents", "cost", "funding", "gpus", "inference", "nvidia", "optics", "pricing", "research", "revenue"]
source: docs/RAG/clean_en/mindstudio/how-openai-and-anthropic-turn-compute-into-profit.md
source_anchor: ""
source_lines: [1, 55]
sha256: 97ce7aa0695990a30c39665c35950ea90a922b4aab314213275fd19118e6f595
---

# how-openai-and-anthropic-turn-compute-into-profit

<!-- source: https://www.mindstudio.ai/blog/ai-lab-revenue-per-megawatt-economics -->

## What is revenue per megawatt and why does it matter?

Revenue per megawatt measures how much money a data center’s worth of compute generates compared to what it costs to run. For frontier AI labs, this single number explains the difference between years of venture-funded losses and a sudden, sharp turn toward profitability. According to SemiAnalysis founder Dylan Patel, the base cost of running a megawatt of AI compute sits around $10 to $15 million a year. Older models barely covered that cost. Newer ones, like Anthropic’s latest Opus release, are reportedly generating revenue as high as $50 million per megawatt, turning what used to be a money-losing operation into one of the highest-margin businesses in tech.

## TL;DR

- **Anthropic turned cash-flow positive in Q2** , and OpenAI is believed to have followed in Q3, marking a shift from purely venture-funded losses to revenue-funded growth.
- The economics flipped because **newer models generate far more revenue per megawatt** than the cost of the compute running them, with Anthropic reportedly seeing up to $50 million in revenue against $10 to $15 million in compute costs.
- That margin lets labs **self-fund a growing share of their own training runs** , reducing reliance on outside capital even as total spending keeps rising.
- OpenAI and Anthropic are absorbing an outsized share of new global compute, reportedly around **30% of incremental compute added this year, rising toward 40 to 50% next year** .
- Newer chip generations pack **three to five times more performance per watt** than prior hardware, so each new gigawatt of capacity is worth disproportionately more than older gigawatts.
- The **physical supply chain (chipmaking tools, fabs, power) reacts far slower** than lab revenue grows, creating a widening gap between what labs could spend and what the world can actually build.
- Total AI infrastructure capital spending is estimated at **over $1 trillion this year and more than $2 trillion by 2028** , with labs taking a growing slice of that pool.

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

## How did AI labs go from losing money to turning a profit?

For most of their existence, OpenAI and Anthropic operated the way most venture-backed startups do: burn investor cash to acquire users and improve the product, worry about margins later. Serving early models like GPT-4 on older Nvidia Hopper GPUs reportedly generated negative gross margin. Every query cost the company more than it brought in.

That changed as model quality and pricing caught up to the underlying compute cost. Patel points to Anthropic crossing into profitability in the second quarter, with OpenAI believed to have reached a similar point in the third quarter, helped by growth in coding-focused products like Codex and newer model releases. The mechanism is simple: the base cost of a megawatt of compute has stayed roughly flat at $10 to $15 million a year, while the revenue that megawatt can generate when running a frontier model has climbed sharply, reportedly reaching $50 million per megawatt in Anthropic’s case.

That gap between cost and revenue is what converts a lab from a cash-burning research lab into a self-funding business. Spend $10 on inference capacity, generate $50 in revenue, then reinvest the difference into training the next model. The loop compounds.

## Why does this let labs outbid everyone else for compute?

Once a company generates several times more revenue per megawatt than the underlying compute costs, it can afford to pay more for that compute than almost anyone else competing for the same GPUs, TPUs, or data center capacity. This is the core dynamic reshaping who controls global compute.

Patel notes that at the start of this year, OpenAI and Anthropic each controlled roughly 2 gigawatts of compute; by year’s end both are above 5 gigawatts, a three to four times increase. More strikingly, the two labs together are absorbing around 30% of all *incremental* compute added globally this year, a figure projected to climb to 40 to 50% next year as previously signed contracts come online. Some of that capacity is being built by new entrants such as SpaceX, which is reportedly constructing compute capacity to lease directly to the labs, since they can pay the highest marginal price for it. Both OpenAI and Anthropic are also moving to build and deploy their own chips and infrastructure rather than relying solely on rented capacity.

The logic is straightforward: if you can turn compute into revenue at multiples that dwarf what anyone else can achieve, you can always outbid the market. Patel argues this isn’t a temporary blip but an accelerating trend, one that could put half of the world’s incremental new compute in the hands of just two companies within roughly a year.

## Is the physical supply chain able to keep up?

## Remy doesn't build the plumbing. It inherits it.

Other agents wire up auth, databases, models, and integrations from scratch every time you ask them to build something.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

Not easily, and this is where the economics get strange. Compute doesn’t materialize instantly. It requires chip fabrication tools, cleanrooms, power infrastructure, and specialized components like the mirrors used inside ASML’s extreme ultraviolet lithography machines. Patel describes a rough model where roughly $6 billion of fab-level capital expenditure produces enough tooling to build a gigawatt of compute capacity annually, and that gigawatt can then generate on the order of $100 billion in end AI revenue per year once deployed and run over its useful life. Even after accounting for the many layers of cost in between (data centers, power, installation, and margin for every company in the supply chain), the discrepancy between what it costs to build the capacity and what it can earn once running is enormous, on the order of 100x by Patel’s rough accounting.

That gap is precisely why demand is racing ahead of supply. Everyone in the chain, from chipmakers to power providers, would need years of lead time to expand output to match where lab revenue is headed. Component makers like Carl Zeiss (which manufactures optics for ASML’s lithography tools) are reportedly ramping toward producing enough parts for around 100 EUV tools by the end of the decade, a target that has already been revised upward once as the scale of AI demand became clearer. But retooling a supply chain this specialized doesn’t happen overnight even when the money is available. Patel calls it a “whip” effect: the signal that more capacity is needed takes a long time to travel from the labs at the end of the chain back to the equipment makers at the start of it.

## Will lab cash flow alone be enough to fund this expansion?

Not yet, and probably not for a while. Even if combined lab revenue reaches into the hundreds of billions of dollars next year, total AI infrastructure capital expenditure is projected to be around $2 trillion. That’s a mismatch labs can’t close with operating cash flow alone, meaning outside capital, whether from investors, cloud partners, or new financing structures, remains essential to fund the buildout. Patel argues this isn’t necessarily a problem: healthy, growing businesses generally want capital expenditure to run ahead of current returns, since that’s how you build capacity for future growth rather than just current demand.

