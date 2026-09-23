---
id: collect-mindstudio/mindstudio/ai-lab-revenue-per-megawatt-economics
title: "How OpenAI and Anthropic Turn Compute Into Profit"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "Nvidia", "OpenAI", "SpaceX"]
dates: ["2026-09-23"]
keywords: ["compute", "capex", "cost", "gpu", "gpus", "inference", "nvidia", "open-weight", "optics", "pricing", "revenue", "training"]
source: docs/RAG/Collect RAG/02_mindstudio/ai-lab-revenue-per-megawatt-economics.md
source_anchor: ""
source_lines: [1, 56]
sha256: 5989af956894bbd53cc625dbc44e32ae0c1cf78dd36ea80cb02b4c9e9eeef808
---

# How OpenAI and Anthropic Turn Compute Into Profit

## Metadata

- **Source** : https://www.mindstudio.ai/blog/ai-lab-revenue-per-megawatt-economics
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article examines the revenue-per-megawatt economics behind the 2026 shift of OpenAI and Anthropic from venture-funded losses to profitability. Revenue per megawatt measures how much money a data center's worth of compute generates versus what it costs to run. According to SemiAnalysis founder Dylan Patel, the base cost of running a megawatt of AI compute sits around $10 to $15 million a year. Older models barely covered that cost; newer ones, like Anthropic's latest Opus release, reportedly generate as high as $50 million in revenue per megawatt, turning a money-losing operation into one of the highest-margin businesses in tech.

For most of their existence OpenAI and Anthropic burned investor cash to acquire users. Serving early models like GPT-4 on older Nvidia Hopper GPUs reportedly generated negative gross margin. That changed as model quality and pricing caught up to compute cost. Patel points to Anthropic crossing into profitability in Q2, with OpenAI believed to have followed in Q3, helped by coding-focused products like Codex and newer model releases. The mechanism: the base cost of a megawatt stayed roughly flat at $10–15M/year while the revenue that megawatt can generate climbed sharply. Spend $10 on inference capacity, generate $50, reinvest the difference into training the next model — a compounding loop.

Because a lab generates several times more revenue per megawatt than the underlying compute costs, it can outbid almost anyone competing for the same GPUs, TPUs, or data center capacity. Patel notes OpenAI and Anthropic each controlled roughly 2 gigawatts of compute at the start of the year; by year's end both are above 5 gigawatts (a 3–4x increase). Together they are absorbing around 30% of all incremental compute added globally this year, projected to climb to 40–50% next year. New entrants such as SpaceX reportedly build compute capacity to lease directly to the labs since they can pay the highest marginal price. Both labs are also moving to build and deploy their own chips and infrastructure.

The physical supply chain reacts far slower than lab revenue grows. Compute requires chip fabrication tools, cleanrooms, power infrastructure, and specialized components like the mirrors inside ASML's EUV lithography machines. Patel describes a rough model where ~$6 billion of fab-level capex produces enough tooling for a gigawatt of annual compute capacity, which can generate on the order of $100 billion in end AI revenue per year — an enormous discrepancy (roughly 100x by his accounting). Component makers like Carl Zeiss (optics for ASML) are ramping toward parts for around 100 EUV tools by decade's end, a target already revised upward. Patel calls this a "whip" effect: the signal for more capacity takes a long time to travel from labs back to equipment makers.

Even if combined lab revenue reaches hundreds of billions next year, total AI infrastructure capex is projected at around $2 trillion — a mismatch operating cash flow alone cannot close, so outside capital remains essential. Patel argues this is healthy, since growing businesses want capex ahead of current returns. The more interesting question is compute pricing: if labs can profitably pay $25M, $40M, or more per megawatt while general-purpose renters operate at $10–15M, the market bifurcates. It is already relatively straightforward for smaller operators to rent GPU capacity, run open-weight models, and profit at the baseline price, which is pushing the floor price of compute upward. Labs with far higher revenue per megawatt have more room to keep paying more, concentrating global compute in a few hands.

## Key points

- Anthropic turned cash-flow positive in Q2 2026; OpenAI believed to have followed in Q3.
- Revenue per megawatt: compute costs ~$10–15M/year per MW; Anthropic reportedly earns up to $50M per MW.
- The cost-revenue gap lets labs self-fund a growing share of training runs.
- OpenAI and Anthropic each grew from ~2 GW to >5 GW of compute in 2026; together absorbing ~30% of incremental global compute, heading to 40–50%.
- Newer chip generations pack 3–5x more performance per watt, making each new gigawatt worth more.
- Physical supply chain (chip tools, fabs, power) scales far slower, creating a "whip" effect and a ~100x build-vs-revenue gap.
- Total AI infrastructure capex ~$1T this year, >$2T by 2028; outside capital still essential.
- Compute market is bifurcating, pushing the floor price of compute upward.

## Technical data / figures

| Item | Value |
|---|---|
| Compute cost per MW | ~$10–15M / year |
| Anthropic revenue per MW | up to ~$50M |
| Anthropic profitability | Q2 2026 (cash-flow positive) |
| OpenAI profitability | Believed Q3 2026 |
| Compute capacity start of year | ~2 GW each |
| Compute capacity end of year | >5 GW each |
| Share of incremental global compute | ~30% (2026) → 40–50% (2027) |
| Fab capex for 1 GW/year | ~$6B |
| End AI revenue per GW/year | ~$100B |
| Build-vs-revenue discrepancy | ~100x |
| EUV tool parts target (Carl Zeiss) | ~100 tools by end of decade |
| Total AI infra capex 2026 | >$1 trillion |
| Total AI infra capex 2028 | >$2 trillion |

## Why this source matters for the RAG

It provides the core economic model (revenue per megawatt) explaining AI lab profitability, compute concentration, and supply-chain bottlenecks — essential context for understanding infrastructure investment and compute pricing. It also explains why open-weight operators can profit at baseline compute prices, tying directly to the local/cloud economics theme.

