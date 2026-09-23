---
id: collect-mindstudio/mindstudio/dark-bloom-earnings-requirements
title: "Dark Bloom Earnings: RAM Requirements and Payout Mechanics Explained"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Apple", "OpenRouter", "Stripe"]
dates: ["2026-09-23"]
keywords: ["attention", "benchmark", "compute", "cost", "distribution", "inference", "memory", "open-weight", "qwen", "revenue", "throughput"]
source: docs/RAG/Collect RAG/02_mindstudio/dark-bloom-earnings-requirements.md
source_anchor: ""
source_lines: [1, 52]
sha256: 4d0c0562d37e0c29a252ae76d2d71cfa835c389d61d37104d1814a8332c917e5
---

# Dark Bloom Earnings: RAM Requirements and Payout Mechanics Explained

## Metadata

- **Source** : https://www.mindstudio.ai/blog/dark-bloom-earnings-requirements
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article focuses on the earnings, hardware requirements, and payout mechanics of Dark Bloom, a project that turns idle Apple silicon Macs into nodes in a distributed inference network. Instead of a company building a massive data center, individual users install software that lets other people's AI requests run on their machine, with the owner paid based on usage. It's a peer-to-peer model for AI compute, launched only days before generating significant attention, reportedly serving 4.5 billion tokens in about a week through its OpenRouter listing. It runs open-source models like Qwen, Gemma, and GPT-OSS.

On earnings: the clearest public number comes from Dark Bloom's own earnings estimator, which projects a monthly payout by hardware. For a Mac Studio running an M5 Ultra chip with 96GB of unified memory, the tool estimated around $37 per month. That's a projection, not guaranteed income, based on current inference demand and the machine's token-serving capacity; both variables can shift as more people join or demand for hosted models changes. The article is realistic: $37/month is not a job replacement or meaningful side hustle on its own, but the appeal is that it's largely passive — a Mac that would otherwise sit idle can generate small incremental income. The main offsetting cost is electricity, and Apple silicon is power-efficient, so the incremental utility cost is expected to be small relative to earnings, though exact figures depend on local electricity rates.

On hardware requirements: Dark Bloom currently requires at least 48GB of RAM to enroll as a provider. This wasn't arbitrary — it was raised after a surge in demand, used as a quality bar to keep performance consistent. Many Mac Studios, higher-configuration Mac minis, and higher-end MacBook Pros already meet or exceed it. During enrollment, the software runs a benchmark test measuring tokens-per-second throughput before fully activating a node, confirming the machine can serve models at acceptable speed before receiving real inference traffic. The verification model used is GPT-OSS 20B, a relatively compact open-weight model realistic to run on capable Apple hardware. The project has said the RAM minimum could come down over time as they refine the system and add more nuanced qualification methods.

On payout mechanics: getting paid requires connecting a Stripe account to your bank. Stripe acts as intermediary, meaning Dark Bloom does not have direct access to a provider's bank account — it can only push payments through Stripe's infrastructure, not pull funds or access banking credentials. This is a standard pattern for platforms paying many individual users without building their own banking integration. As of now, providers keep all revenue their machine generates, but this is explicitly called out as something that could change — a future version could introduce a revenue split where the platform takes a percentage, similar to many marketplace or gig-economy platforms after an initial growth phase.

On safety: this is the most-asked question given the software runs with device-level access. The codebase is publicly available for inspection. An AI-assisted code review reportedly found no hidden malware, no cryptocurrency mining, and no credential theft, and characterized the device management/access concern as low risk. Dark Bloom's bigger claim is architectural: a published white paper describes preventing the machine owner (who has physical and root access) from observing prompts or responses, running inference inside a single hardened Swift process using MLX Swift LM, avoiding subprocesses, local servers, or interprocess communication that could expose data. None of this eliminates risk entirely — the project is new and changes quickly — but the combination of open-source code, a published technical explanation, and Stripe-mediated payouts suggests a project trying to be transparent.

## Key points

- Dark Bloom rents idle Mac compute to run Qwen, Gemma, and GPT-OSS for other users.
- Official estimator: ~$37/month for a Mac Studio with M5 Ultra and 96GB memory.
- 48GB RAM minimum to enroll, raised from an earlier threshold due to demand.
- Payouts via connected Stripe account; Dark Bloom never has direct bank access.
- Enrollment includes a tokens-per-second benchmark (GPT-OSS 20B) before activation.
- Privacy architecture: hardened Swift process (MLX Swift LM) prevents owners from observing prompts/responses.
- Providers currently keep 100% of revenue; split may change.
- Reported 4.5 billion tokens served in about a week via OpenRouter.

## Technical data / figures

| Item | Value |
|---|---|
| Estimated earnings | ~$37/month (M5 Ultra, 96GB) |
| Min RAM | 48GB |
| Verification model | GPT-OSS 20B |
| Payout intermediary | Stripe |
| Token volume | 4.5 billion in ~1 week |
| Models | Qwen, Gemma, GPT-OSS |
| Privacy framework | MLX Swift LM (hardened Swift process) |
| Revenue split | 100% to providers (may change) |
| Distribution | OpenRouter listing |

## Why this source matters for the RAG

It supplies concrete payout, hardware, and safety details for distributed inference, complementing the network-overview article with the specific economics that would drive adoption. It is essential for anyone evaluating whether to run a provider node or analyze peer-to-peer compute markets.

