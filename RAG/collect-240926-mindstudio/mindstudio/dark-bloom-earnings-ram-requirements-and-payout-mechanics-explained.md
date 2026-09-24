---
id: collect-240926-mindstudio/mindstudio/dark-bloom-earnings-ram-requirements-and-payout-mechanics-explained
title: "dark-bloom-earnings-ram-requirements-and-payout-mechanics-explained"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Apple", "OpenRouter", "Stripe"]
dates: []
keywords: ["attention", "benchmark", "compute", "cost", "gpu", "gpus", "inference", "memory", "open-weight", "qwen", "revenue", "throughput"]
source: docs/RAG/clean_en/mindstudio/dark-bloom-earnings-ram-requirements-and-payout-mechanics-explained.md
source_anchor: ""
source_lines: [1, 65]
sha256: 5b2c3864577297d9d684f2390a9edba985ed3b5d75e9fb08f61925c182564c1e
---

# dark-bloom-earnings-ram-requirements-and-payout-mechanics-explained

<!-- source: https://www.mindstudio.ai/blog/dark-bloom-earnings-requirements -->

## What is Dark Bloom and how does it pay you?

Dark Bloom is a new project that turns idle Apple silicon Macs into nodes in a distributed inference network. Instead of a company building a massive data center, individual users install a small piece of software that lets other people’s AI requests run on their machine. In exchange, the machine owner gets paid based on usage. It’s a peer-to-peer model for AI compute rather than a centralized one, and it launched only days before generating significant attention online, reportedly serving 4.5 billion tokens in about a week through its listing on OpenRouter.

## TL;DR

- Dark Bloom lets you **rent out idle Mac compute** to run open-source AI models like Qwen, Gemma, and GPT-OSS for other users.
- The official earnings calculator on dark bloom.dev estimated **roughly $37 a month** for a Mac Studio with an M5 Ultra chip and 96GB of memory.
- A **48GB RAM minimum** is currently required to enroll as a provider, a bar raised from an earlier threshold due to high demand.
- Payouts run through a **connected Stripe account** , meaning Dark Bloom itself never has direct access to your bank account.
- The project claims a **privacy-preserving architecture** using a hardened Swift process (MLX Swift LM) so the machine owner can’t observe the prompts or responses being processed.
- Setup happens through a **command-line install** tied to a Dark Bloom account, with a native Mac app described as coming later.
- Providers currently keep **100% of earned revenue** , though that split is explicitly described as something that could change in the future.

## How much can you actually earn running Dark Bloom?

The clearest public number comes from Dark Bloom’s own earnings estimator, which lets you plug in your hardware and see a projected monthly payout. For a Mac Studio running an M5 Ultra chip with 96GB of unified memory, the tool estimated around $37 per month. That figure is not guaranteed income. It’s a projection based on current demand for inference and the machine’s capacity to serve tokens, and both of those variables can shift as more people join the network or as demand for the models it hosts changes.

It’s worth being realistic about what $37 a month means. It’s not a replacement for a job or even a meaningful side hustle on its own. The appeal is that it’s largely passive: a Mac that would otherwise sit idle on a desk overnight or during the day can generate small, incremental income by serving inference requests for open-weight models like Qwen, Gemma, and GPT-OSS variants. The main offsetting cost is electricity, and Apple silicon machines are known for being power-efficient compared to typical GPU servers, so the incremental utility cost of running a provider node is expected to be small relative to the earnings, though the exact number depends on local electricity rates.

## What hardware do you need to qualify?

Dark Bloom currently requires at least 48GB of RAM to enroll as a provider. This isn’t an arbitrary number picked at launch. It was raised after the project saw a surge in demand from people wanting to join, and the team used it as a quality bar to keep performance consistent across the network. Many Mac Studios, Mac minis with higher configurations, and higher-end MacBook Pro models already meet or exceed this threshold, which is likely why the requirement was set there rather than lower.

During enrollment, the software runs a benchmark test to measure tokens-per-second throughput before fully activating a node. This step confirms the machine can actually serve models at an acceptable speed before it starts receiving real inference traffic from other users. The specific model used in this verification step has been GPT-OSS 20B, a relatively compact open-weight model that’s realistic to run locally on capable Apple hardware.

There’s an expectation, stated by the project itself, that the RAM minimum could come down over time as they refine the system and add more nuanced ways to qualify machines. For now, if your Mac has less than 48GB of memory, you’re not eligible to earn as a provider.

## How does the payout mechanism actually work?

Getting paid requires connecting a Stripe account to your bank. Stripe acts as the intermediary in this relationship, which means Dark Bloom does not have direct access to a provider’s bank account. The platform can only push payments to you through Stripe’s infrastructure, not pull funds or access banking credentials directly. This structure is a fairly standard pattern for platforms that need to pay out many individual users without building their own banking integration from scratch.

As of now, providers keep all the revenue their machine generates. That detail matters because it’s explicitly called out as something that could change down the line. A future version of Dark Bloom could introduce a revenue split where the platform takes a percentage, similar to how many marketplace or gig-economy platforms evolve after an initial growth phase where users keep the full upside.

## Is Dark Bloom safe to install on your Mac?

This is the question most people ask first, and it’s a fair one given that the software runs with device-level access on your machine. The codebase is publicly available, meaning anyone can inspect or audit it rather than trusting it blindly. An AI-assisted code review of the project reportedly found no hidden malware, no cryptocurrency mining, and no credential theft mechanisms in the code. The concern around device management and access was characterized as low risk in that review.

The bigger privacy claim Dark Bloom makes is architectural. The project published a white paper describing how it prevents the machine owner, who has physical and root access to the hardware, from observing the prompts or responses being processed on their device. The stated approach runs inference inside a single hardened Swift process using MLX Swift LM (Apple’s own machine learning framework for its silicon GPUs), avoiding subprocesses, local servers, or interprocess communication that could otherwise expose data in transit. In plain terms, the goal is that even though your Mac is doing the computational work, you as the owner can’t see what’s being asked or answered.

None of this eliminates risk entirely. The project is new, changes are happening quickly, and anyone installing software with system-level access should weigh that against the modest financial upside. But the combination of open-source code, a published technical explanation, and a Stripe-mediated payout system suggests a project trying to be transparent about how it works rather than obscuring its mechanics.

## Frequently Asked Questions

### How much RAM do I need to run Dark Bloom?

You currently need at least 48GB of RAM on your Mac to enroll as a provider. This threshold was raised from an earlier level due to high demand and is used as a baseline quality bar for the network.

### How much money can I realistically make with Dark Bloom?

Based on the project’s own earnings estimator, a Mac Studio with an M5 Ultra chip and 96GB of memory was projected to earn around $37 a month. Actual earnings will vary by hardware, uptime, and demand for inference at any given time.

### How do I get paid for running Dark Bloom?

Payouts go through a connected Stripe account linked to your bank. Stripe sits between you and Dark Bloom, so the platform can send payments but never has direct access to your bank account.

### Can Dark Bloom see what data is being processed on my Mac?

The project’s stated design runs inference in an isolated, hardened process meant to prevent the machine owner from observing the prompts or outputs passing through their hardware, according to its published white paper.

### Do I keep all the money my Mac earns?

As of now, providers keep 100% of the revenue their machine generates. The project has indicated this split could change in future versions.
