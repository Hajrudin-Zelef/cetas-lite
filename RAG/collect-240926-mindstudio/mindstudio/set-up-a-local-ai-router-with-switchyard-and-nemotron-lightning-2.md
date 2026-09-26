---
id: collect-240926-mindstudio/mindstudio/set-up-a-local-ai-router-with-switchyard-and-nemotron-lightning-2
title: "set-up-a-local-ai-router-with-switchyard-and-nemotron-lightning"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "Moonshot", "Nvidia", "OpenRouter"]
dates: []
keywords: ["agents", "benchmark", "claude", "cost", "kimi", "nvidia", "reasoning", "tokens per second"]
source: docs/RAG/clean_en/mindstudio/set-up-a-local-ai-router-with-switchyard-and-nemotron-lightning.md
source_anchor: ""
source_lines: [76, 110]
sha256: ec5cc2dc5c15d2edd2e774d22a67cdf3b05af326130a5905948949463b08cd72
---

# set-up-a-local-ai-router-with-switchyard-and-nemotron-lightning

The economics depend heavily on the workload, but the general pattern is that frontier-only setups pay a steep premium for the last few percentage points of task completion. In one comparison referenced from an internal evaluation, a frontier model like Claude Opus completed roughly 80% of tasks in a benchmark set but at a cost of around $180 for that run. Pairing a frontier model with a faster workhorse model like Nemotron 3.5 Lightning through a router can achieve similar completion performance at substantially lower total cost, because only the tasks that genuinely need frontier reasoning get escalated to it.

## What does this look like in a real workflow?

One practical example involves network alert analysis, where incoming data contains sensitive details like hostnames, IP addresses, and account numbers. In a privacy-focused setup, Nemotron 3.5 Lightning runs locally and generates a compact digest of each alert, stripped of personally identifying information. Only the sanitized, combined digest gets sent to a hosted model like Kimi K3 for deeper root-cause analysis. The sensitive data never leaves local hardware.

A second variant applies escalation routing directly to alert analysis: every alert starts with the local model, and a stage router decides, based on signals already present in the conversation, whether to escalate specific alerts to the more capable hosted model.

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

In both cases, routing decisions happen once per task rather than once per conversational turn. Routing per turn is expensive, breaks caching, and adds classifier overhead on every single exchange, so the practical pattern is to decide the route at the start of a task and stick with it.

## Frequently Asked Questions

### What’s the difference between SwitchYard and a service like OpenRouter?

OpenRouter provides a single API endpoint that gives access to multiple hosted models, which is a form of routing but runs entirely through OpenRouter’s infrastructure. SwitchYard is open-source and designed to run on your own infrastructure, including local models, giving you direct control over targets, strategies, and data flow.

### Do I need a DGX Spark to run Nemotron 3.5 Lightning locally?

The demonstrated setup used a DGX Spark, which is NVIDIA’s local AI development hardware, to achieve around 71 tokens per second. The model can potentially run on other local hardware capable of serving a 30 billion parameter mixture-of-experts model, though performance will vary with the hardware used.

### Which SwitchYard routing strategy should I start with?

Stage routing is a reasonable default since it reads signals already present in the conversation and costs nothing extra to run. LLM classification and escalation add value when you need a dedicated decision point, but both introduce an additional model call, so they’re worth reserving for cases where free signals aren’t enough.

### Does using a router mean sacrificing output quality?

Not necessarily. The goal is to match task complexity to model capability rather than downgrading everything. Simple tasks route to faster, cheaper models, while genuinely complex tasks still escalate to frontier models, which can preserve overall task completion rates while cutting total spend.

### Can a router help with data privacy, not just cost?

Yes. Routing sensitive data to a local model first, then stripping identifying details before sending a summary to a hosted API, is one of the four core reasons organizations adopt routing layers, alongside cost, speed, and specialization.
