---
id: vague2-nerdykings/nerdykings/deepseek-harness-fin-claude-code-1
title: "DeepSeek Harness : La Fin De Claude Code ?"
domain: nerdykings
role: reference
task: article
actors: ["Anthropic", "DeepSeek"]
dates: ["2026-09-23"]
keywords: ["claude", "deepseek", "agent", "agentic", "agents", "benchmarks", "license", "memory", "mit license", "open source", "reasoning", "sandbox"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/deepseek-harness-fin-claude-code.md
source_anchor: ""
source_lines: [1, 40]
sha256: 8186d8ca999070ebea5210a5a6b34291eb36f4856a59041344182797e0fee55d
---

# DeepSeek Harness : La Fin De Claude Code ?

## Metadata

- **Source** : https://www.nerdykings.com/blog/deepseek-harness-fin-claude-code.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

DeepSeek released a project that could become much more important than a simple new coding agent: **Harness**. At first glance it looks like an open-source clone of Claude Code — it reads files, runs terminal commands, uses tools, and works autonomously on a project. But what makes it interesting lies elsewhere: DeepSeek built an extremely modular environment where almost every component of the agent can be replaced — the model, the tools, the permissions, session management, and even the loop controlling the agent's behavior. The author installed Harness, tested its Creator mode, and evaluates whether this "everything is a plugin" approach holds up.

What is a "harness"? It's an essential part of AI agents that's rarely discussed. When using Claude Code, it feels like all capabilities come directly from Claude, but the model is only part of the system: you must build a whole environment around it — give it file and terminal access, provide tools, manage context and permissions, retrieve action results, and above all determine how it chains steps to accomplish a task. That system around the model is the harness. The model provides reasoning; the harness turns reasoning into concrete action by deciding what the agent can do and how. This distinction matters: two agents using the exact same model can get very different results simply because their harness is designed differently. Recent work even showed that keeping an identical DeepSeek model but improving its harness could strongly boost performance on some tasks — not that a good harness makes a bad model intelligent, but that agent performance now depends as much on the environment built around the model as on the model itself.

"Everything is a plugin" — Harness architecture. Many agentic systems have a relatively fixed core to which you connect different tools. Harness goes much further, making the architecture itself modular. It relies on a framework called **Cortex**, which assembles the runtime's parts as plugins. When DeepSeek says "everything is a plugin," it isn't just a way to install a few extra extensions: the adapter communicating with the model can be replaced, as can the tool registry, the permissions system, session persistence, the sandbox, sub-agents, and even the main loop orchestrating the agent's behavior. Usually, an AI agent "plugin" means giving it a new capability; here DeepSeek allows modifying much more deeply how the agent itself is built. Concretely, you could have an agent using DeepSeek as the model with your own tools, permissions, and environment — and if another model becomes better or cheaper tomorrow, you could swap it while keeping much of the existing infrastructure. Harness separates three distinct things: the model (reasoning), the harness (orchestrating the agent and its tools), and the environment (files, terminal, network, services). That's why Harness looks more like infrastructure for building your own agent than a simple Claude Code clone.

Creator mode: the author asked the agent to modify itself. Harness offers four modes: **standard** (code, edit files, run commands, web search, sub-agents, task planning — used 90% of the time); **PTC**, where the agent chains several tools in one TypeScript program rather than one by one (for complex tasks); **minimal**, deliberately limited without the big toolbox (more useful for benchmarks); and **Creator**, designed to build your own agent — it has standard capabilities but can also inspect the runtime internals, test plugins, and create custom presets. Testing Creator, the author asked it to create a plugin adding a real-time clock in the top-right corner of the interface. Moments later the clock appeared. He then asked it to make the clock draggable by mouse without losing functionality. The agent re-read the source of the plugin it had just created, examined the overlay system, and delivered a floating draggable clock. That's the idea behind "everything is a plugin": instead of only using an agent as designed by its developer, you can adapt its environment to your needs directly from the agent itself. In Creator mode, the agent can create and modify its own **Cordis** plugins — new tool, runtime, interface.

Does it really replace Claude Code? On paper both can work on code, manipulate files, use a terminal, and perform tasks autonomously, but their philosophy differs. Claude Code is above all an integrated product: Anthropic controls the model and much of the environment, making it simple to use and highly optimized. DeepSeek Harness is more an infrastructure for building and customizing your own agentic environment. It's open source under **MIT license**, not tied to a single model, and its architecture allows much deeper replacement of system components. The author doesn't think it replaces Claude Code today: Harness just launched and DeepSeek officially calls it a **developer preview**, with possible incompatible changes. Claude Code is also much more mature on security: Harness's sandbox can limit filesystem modifications, but DeepSeek itself says it doesn't by itself restrict network or process visibility. So for an immediately ready-to-use performant coding agent, Claude Code remains the obvious choice; for building and deeply modifying your own agentic environment, Harness offers something much more open.

Installation (summary): install Node.js if needed; search "DeepSeek Harness" and go to the official site to copy the install command; paste and run it in the terminal; retrieve the generated local address and open it in a browser; enter your DeepSeek API key, create a workspace, choose your mode, and go.

The author's view: the value isn't the promise of replacing Claude Code tomorrow — DeepSeek doesn't even claim that. The architectural bet matters: as models become increasingly interchangeable, the real battle shifts to the layer built around them — tools, permissions, memory, control loop. This echoes DeepSeek's open-source strategy on V4: make infrastructure open and hackable rather than locking everything in a finished product. Today, for daily reliable use, Claude Code remains largely ahead. But for developers who want to understand — and above all control — what happens under an agent's hood, an MIT framework where absolutely everything is a plugin is a real signal.

## Key points

- DeepSeek **Harness** is an open-source (MIT), highly modular agent environment, not just a coding-agent clone.
- Core concept: a **harness** surrounds the model and turns reasoning into action (tools, permissions, context, control loop).
- "Everything is a plugin" via the **Cortex** framework: model adapter, tools, permissions, sessions, sandbox, sub-agents, and main loop are all replaceable.
- Three separated layers: model (reasoning), harness (orchestration), environment (files/terminal/network).
- Four modes: standard, PTC (TypeScript tool chaining), minimal, and Creator (build your own agent).
- Creator mode tested successfully: created a real-time clock plugin and made it draggable by re-reading its own source.
- Versus Claude Code: integrated optimized product vs open hackable infrastructure; Harness is a **developer preview** with possible breaking changes.
- Security caveat: Harness sandbox limits filesystem changes but not network or process visibility by itself.

## Technical data / figures

