---
id: collect-240926-mindstudio/mindstudio/how-xai-s-cursor-deal-quietly-built-grok-into-a-real-contender-2
title: "how-xai-s-cursor-deal-quietly-built-grok-into-a-real-contender"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "OpenAI", "SpaceX", "xAI"]
dates: []
keywords: ["grok", "acquisition", "agent", "benchmark", "claude", "compute", "consumer", "cost", "gpt-5.6", "gpu", "grok 4", "reasoning"]
source: docs/RAG/clean_en/mindstudio/how-xai-s-cursor-deal-quietly-built-grok-into-a-real-contender.md
source_anchor: ""
source_lines: [58, 90]
sha256: 58d33b62a363de6f5363680acb9490ed363e083df13d85055d379ae3e3304941
---

# how-xai-s-cursor-deal-quietly-built-grok-into-a-real-contender

xAI’s version of this flywheel started later but moved fast once Cursor was in the fold. Grok 4.5 became a favorite as a sub-agent model inside Cursor for many developers because it was fast and capable. That real-world usage, plus Cursor’s existing data reserves, fed directly into training Grok 4.6. xAI’s own materials describe using Grok 4.5 to help regenerate training trajectories (across reasoning tasks, agent workflows, and domains like software engineering and general knowledge work) used to train Grok 4.6. That’s a live example of one generation of a model helping construct the training data for its successor, a pattern now common across frontier labs even if it doesn’t amount to fully autonomous self-improvement.

## What is Grokbot, and why does it matter for xAI’s strategy?

Grok 4.6 doesn’t only show up in Cursor. It also powers Grokbot, a newer xAI product aimed at a broader, non-technical audience rather than developers. Grokbot reportedly strips out model selection entirely and hides any visible code, presenting users only with finished outputs like documents and presentations. The pitch is that most valuable knowledge work is ultimately produced through code, so a strong coding model can power general productivity tasks even for people who never see a line of it.

## Remy doesn't build the plumbing. It inherits it.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

That’s a meaningful shift from where xAI started. Early framing around Grok emphasized a “maximally truthful” chatbot positioned against perceived bias in other assistants. That framing didn’t translate into developer or enterprise adoption, and xAI lost ground to Anthropic and OpenAI for an extended stretch. The pivot toward coding-first infrastructure, paired with consumer-facing products like Grokbot that hide the technical complexity, looks like an attempt to capture both ends of the market at once: developers through Cursor, and general users through Grokbot.

## Frequently Asked Questions

### Did xAI actually acquire Cursor, or just partner with it?

Reporting on the deal describes it as an acquisition connected to the broader xAI/SpaceX corporate structure, with Cursor’s coding data and product becoming part of xAI’s stack rather than a simple partnership.

### Is Grok 4.6 better than GPT-5.6 or Claude Opus?

It depends on the benchmark. Grok 4.6 reportedly leads on some evaluations like GDPval and Harvey Lab, but trails GPT-5.6 Sol Max and Claude Opus on others, including Deep Sweet, a benchmark many developers treat as a good proxy for real-world coding feel.

### Why would Anthropic buy compute from a competitor like xAI?

Anthropic has repeatedly underestimated demand for its models and run into capacity constraints. xAI built a very large GPU fleet quickly and had spare capacity, making a compute deal commercially useful for both sides even though the two labs compete directly on models.

### What is the difference between Cursor and Grokbot as products?

Cursor is a developer-focused coding editor. Grokbot is aimed at a broader, less technical audience, hiding model selection and code entirely and surfacing only finished outputs like documents.

### Is Grok 4.6 cheap compared to other frontier models?

Yes, on a per-token basis. At $2 per million input tokens and $6 per million output tokens, it undercuts GPT-5.6 Sol and Fable significantly, though its cost-per-completed-task rose compared to Grok 4.5 as the model got more capable.
