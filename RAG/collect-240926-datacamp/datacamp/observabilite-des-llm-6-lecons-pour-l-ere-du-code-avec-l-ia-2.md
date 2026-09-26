---
id: collect-240926-datacamp/datacamp/observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia-2
title: "observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "Mistral"]
dates: []
keywords: ["agent", "agents", "claude", "cost", "reasoning", "sandbox"]
source: docs/RAG/clean_en/datacamp/observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia.md
source_anchor: ""
source_lines: [117, 160]
sha256: 66f22c5bc3d6b74a4db4729aa132160432e332cabd11176ad184438b8298fca2
---

# observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia

Cost is the twin constraint of security, and containing the price of putting large language models into production is becoming a discipline in its own right. The answer Lê-Quôc presented at DASH: Datadog's Agent Console.

Ask a developer which model they need: often, they'll name the most powerful (and most expensive) one. Sometimes that's the right choice, but a large share of the work is generic enough that a cheaper, faster model handles it just as well. Distinguishing between them requires analyzing an organization's agent trajectories: which tools they call, how often they succeed, until patterns emerge.

These patterns become heuristics rather than rules: a frontier model like the latest Claude Opus or the GPT models for planning, an economical model like Claude Haiku for generating tests.

| Task | Model tier | Why |
|---|---|---|
| Planning and complex reasoning | Frontier model (e.g., Claude Opus, GPT) | The best reasoning capability pays for itself here |
| Routine, generic code | Mid-tier (e.g., Claude Sonnet, GPT-mini) | Good enough, and far cheaper to run frequently |
| Test generation and simple transformations | Fast and cheap (e.g., Claude Haiku, GPT-nano) | Speed and price win as long as quality holds |

The underlying principle concerns ownership of the decision. If you reduce cost to a single number, you get what Lê-Quôc calls "very low actionability": either everyone cuts spending, and you kill useful work, or everyone keeps going, and the company can't keep up. He prefers putting the data in front of the developers and SREs who choose the models.

## Lesson 6: learn how to learn

Asked what new engineers should study, Lê-Quôc gives an answer that sounds old, but isn't.

You have to learn how to learn.

Alexis Lê-Quôc, CTO at Datadog

Models are the most patient tutors ever invented, capable of explaining anything at any pace — a level of access once reserved for a privileged few. But a tutor is only useful if you question it. The skill is knowing what to ask and how to verify the answer.

He recommends understanding computing layer by layer rather than treating it as magic. Take a scheduler, a load balancer, a sandbox, and ask a model to explain how it works, then dig deeper:

- What does this term mean?
- How do you measure it?
- What are the mathematical foundations?
- How do you know if it's working well?

Studying the classics this way is deliberately slow. He compares it to learning an instrument: you can listen to music all day, but to play the piano, you have to put your hands on the keyboard.

Same goes for code written by AI. Vibe coding is great, he says, as long as you come back to it and ask why it worked: why that architecture choice, do better approaches exist, what was it based on. The goal is not to write less code with AI, but to better understand the code you now produce in far greater quantity.

## In conclusion

Lê-Quôc's central message: the loop hasn't changed, but the speed has. From now on, no human can observe closely enough at AI's cadence: monitoring, and a growing share of construction, pass to agents that neither tire nor panic.

He advocates treating observability as a control plane, not as a collection of graphs. If agents write, test, deliver, and operate software, they need the same grounding in real production data as good engineers, with in addition a person who keeps judgment and the stop button. Datadog positions observability as the layer that makes this balance safe.

The skill expected of engineers is clear: read systems through their behavior in production, not only through their source. To anchor this habit, our Machine Learning in Production skill track is a good starting point.

**Data Science Editor-in-Chief at DataCamp |** **I am passionate about forecasting and development using APIs.**
