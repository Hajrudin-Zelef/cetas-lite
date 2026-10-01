---
id: collect-261001-ia-llm/ia-llm/harness-scaffold-and-the-ai-agent-terms-worth-getting-right-2
title: "harness-scaffold-and-the-ai-agent-terms-worth-getting-right"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft"]
dates: []
keywords: ["agent", "agents", "claude", "copilot", "gemini", "mcp", "memory", "tool use", "training"]
source: docs/RAG/collect-261001-ia-llm/harness-scaffold-and-the-ai-agent-terms-worth-getting-right.md
source_anchor: ""
source_lines: [57, 97]
sha256: a913941836096779933d4e054394014168e2c91bbd2fdb984725d36624c74523
---

# harness-scaffold-and-the-ai-agent-terms-worth-getting-right

An agent called by another agent to handle a specific subtask. It has its own model and scaffold, reasons independently, and returns a result. The calling agent doesn't need to know how it works internally. This is what separates a **sub-agent** from a **tool** (a function call) or a **skill** (packaged knowledge): a sub-agent can itself reason, use tools, and call further sub-agents. The calling agent is sometimes called an **orchestrator**.

The terms above apply whether you're training or deploying. These four are specific to training, where the agent runs through tasks, gets scored, and its model's weights get updated. Every RL training system for LLMs is built around the same pipeline:

The environment is anything you can interact with: a stateful object that takes an action as input, updates its internal state, and returns an observation. In the LLM context, actions are typically tool calls. A filesystem is a simple example: the action `touch foo.txt` updates the state by creating the file, and the observation might be the updated file listing. Definitions vary across frameworks.

We recently published a dedicated guide on this, so rather than compress it here, see The Ultimate Guide to RL Environments for a complete breakdown of types, frameworks, and examples.

The trainer is what makes the agent better: it runs many agent episodes, scores the results and uses them to update the inner model's weights. TRL's GRPOTrainer is a concrete example: a single class that handles episode generation, reward scoring, and weight updates.

A rollout is one full agent run from start to finish: what the agent saw, what it did, and what reward it got at each step. It's also called a *trajectory* or a *trace*, depending on the context. This is the raw data RL algorithms learn from.

The score that tells the training algorithm whether the model is getting better. It can be *verifiable* (tests pass/fail, answer matches), or *learned* (human preferences, LLM-as-judge), *sparse* (one score at the end of an episode), or *dense* (a score at each step). This is what the trainer uses to actually update the inner model's weights. For a thorough breakdown of each type, see the Reward Architecture section in Adithya's guide.

**Rubrics** break the reward into explicit dimensions with weights, rather than a single number. OpenEnv and Verifiers implement rubrics as objects you can combine (`WeightedSum`, `Sequential`, `Gate`).

*If any definition feels imprecise or you've encountered a term we've missed, we'd love to hear from you.*

*Thanks to Pedro Cuenca, Quentin Gallouédec, Shaun Smith, and Adithya S Kolavi for reviewing this post.*

The definition of Agent in popular use changes every year :)

We really leaned into Autonomy in this post from last year, but this year it seems to be all about model + harnesses: https://huggingface.co/blog/ethics-soc-7

All in all, the more precise and technically specific the term, the better!

Coding Agent, or Agent CLI – a command-line application, used to launch and manage agent instances (e.g. Opencode, Claude Code, Copilot CLI, Gemini CLI, etc)

Agent Definition – a file (Markdown/JSON) that describes how an agent behaves; serves as a blueprint from which instances are created (like a class in Java)

Agent Instance – a running process created from an Agent Definition (like an object instantiated from a class in Java).

Agent / AI Agent – avoid; use the specific term (CLI, Definition, or Instance) depending on what you mean; not to forget that “Agent” can also mean the CI pipeline agent or a human agent

it's a standardization protocol that makes the Tool Use → Harness connection much cleaner and more interoperable. The glossary describes what agents do; MCP is a specification for how they do the tool-calling part reliably across different vendors and implementations.

Really useful glossary, and I especially appreciate the attempt to separate the model, harness, scaffold, tools, context, and training loop.

One concern I have is with standardizing on “scaffolding” as the term for the behavior-defining layer around the model. To me, “scaffolding” already feels quite overloaded, specifically in software projects. It is used extensively in code generation, and more narrowly by framework generators to create routes, controllers, models, migrations, components, tests, and other boilerplate structures.

I wonder whether “rigging” might be a better candidate. It still relates naturally to the metaphor of a harness, but is less already-standardized in adjacent software and AI contexts. In the horse/cart sense, rigging is the arrangement that attaches the harness to the thing being pulled and makes the whole setup usable. That feels close to what this layer does for an agent: system prompt, tool descriptions, response formats, memory/context rules, and other configurations that connect the model to the operational environment in a particular way.
