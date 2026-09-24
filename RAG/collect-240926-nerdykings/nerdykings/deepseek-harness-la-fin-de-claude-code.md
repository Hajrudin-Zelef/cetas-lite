---
id: collect-240926-nerdykings/nerdykings/deepseek-harness-la-fin-de-claude-code
title: "DeepSeek Harness: The End Of Claude Code?"
domain: nerdykings
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google"]
dates: []
keywords: ["claude", "deepseek", "agent", "agentic", "agents", "benchmarks", "license", "memory", "mit license", "open source", "reasoning", "sandbox"]
source: docs/RAG/clean_en/nerdykings/deepseek-harness-la-fin-de-claude-code.md
source_anchor: ""
source_lines: [1, 51]
sha256: 04c0907013ce2289f344d1d3432167e4a9f0b193ee033eadd1f87e18bcea7504
---

# DeepSeek Harness: The End Of Claude Code?

<!-- source: https://www.nerdykings.com/blog/deepseek-harness-fin-claude-code.html -->

# DeepSeek Harness: The End Of Claude Code?

DeepSeek has just released a project that could become far more important than just another coding agent: **Harness**. At first glance, it could easily be mistaken for an open source clone of Claude Code — it reads your files, runs commands in the terminal, uses tools, and works autonomously on a project. Except what makes it truly interesting lies elsewhere. DeepSeek has built an extremely modular environment in which virtually every component of the agent can be replaced — the model, the tools, the permissions, session management, and even the loop that controls the agent's behavior. I installed Harness, tested its Creator mode, and I'll tell you whether this "everything is a plugin" approach really holds up.

## What is a "harness," exactly?

To understand the appeal of the project, you first need to understand what a **harness** is — an essential part of AI agents that is ultimately discussed fairly little. When you use Claude Code, you might easily get the impression that all these capabilities come directly from Claude. Yet the model is only one part of the system: you need to build an entire environment around it to allow it to act concretely — give it access to files and the terminal, provide it with tools, manage its context and permissions, retrieve the results of its actions, and above all determine how it will chain all these steps together to accomplish a task. It is precisely this entire system built around the model that is called a harness.

So we can view the model as the part that provides the reasoning, while the harness turns that reasoning into concrete action by determining what the agent can do and how it will do it. And this distinction matters: two agents using exactly the same model can achieve very different results, simply because their harness is not designed the same way. Recent work has even shown that by keeping an identical DeepSeek model but improving its harness, performance could greatly improve on certain tasks. This doesn't mean a good harness makes a bad model intelligent, but it shows that an agent's performance now depends as much on the environment built around the model as on the model itself.

## "Everything is a plugin" — the architecture of Harness

It is precisely on this layer that DeepSeek Harness tries to bring something new. In many agentic systems, you have a relatively fixed core to which you can then connect different tools. With Harness, DeepSeek goes much further by making the architecture itself modular. The project is based on a framework called **Cortex**, which allows the different parts of the runtime to be assembled in the form of plugins.

When DeepSeek claims that *"everything is a plugin,"* it's not just a way of installing a few extra extensions. The adapter that communicates with the model can be replaced, as can the tool registry, the permission system, session persistence, the sandbox, sub-agents, and even the main loop that orchestrates the agent's behavior. This is probably the most important point to understand: usually, when we talk about plugins for an AI agent, we're mainly talking about giving it a new capability. Here, DeepSeek allows for much deeper modification of the way the agent itself is built.

Concretely, you could have an agent that uses DeepSeek as its model, with your own tools, your own permissions, and your own environment. And if tomorrow another model becomes better or cheaper, you could replace it while keeping much of the infrastructure already built. Harness therefore separates three very distinct things: the model that provides the reasoning, the harness that orchestrates the agent and its tools, and then the environment — files, terminal, network, services — that the agent can access. This is why Harness looks more like an infrastructure for building your own agent than a simple clone of Claude Code.

## Creator mode: I asked the agent to modify itself

Once launched, Harness offers four modes. The **standard mode** — coding, modifying files, running commands, searching the web, using sub-agents, planning tasks — the one you'll use in 90% of cases. The **PTC mode**, where the agent chains several tools in a single TypeScript program instead of doing things one by one, rather for complex tasks with many successive actions. The **minimal mode**, deliberately limited, without the big toolbox — more useful for benchmarks than for daily use. And finally, the **Creator mode**, designed for building your own agent: it has the capabilities of the standard mode, but can also inspect the internal workings of the runtime, test plugins, and create custom presets.

This is the one I wanted to actually test. I asked it: *"create a plugin that adds a real-time clock to the top right corner of the interface"*. A few moments later, the clock was there, in the top right. Nice, but not exactly what I wanted — so I asked it again to modify its own plugin so that the clock could be moved with the mouse, by drag and drop, without losing its current functionality. The agent went back to read the source code of the plugin it had just created, examined the overlay system, and delivered a floating clock that you can indeed drag around the interface. That's exactly the idea behind "everything is a plugin": instead of only using an agent as it was designed by its developer, you can start adapting your environment to your own needs, directly from the agent itself.

## Does it really replace Claude Code?

On paper, both can work on code, manipulate files, use a terminal, and perform tasks autonomously. But their philosophy is quite different. **Claude Code** is first and foremost an integrated product: Anthropic controls the model and much of the environment around it, which gives you something very easy to use and already highly optimized. **DeepSeek Harness**, on the other hand, is more of an infrastructure that lets you build and customize your own agentic environment. The project is open source under the MIT license, it isn't tied to a single model, and its architecture allows you to replace the various building blocks of the system much more deeply.

But personally, I don't think we can say it replaces Claude Code today. Harness has just come out, and DeepSeek still officially presents it as a **developer preview**, with incompatible changes that may still happen. Claude Code also remains much more mature on the security side: the sandbox shipped with Harness can limit modifications to the file system, but DeepSeek itself points out that it doesn't by itself restrict the network or process visibility. So if your goal is simply to have a performant coding agent that's ready to use immediately, Claude Code remains the most obvious choice. If, on the other hand, you want to build your own agentic environment and be able to deeply replace or modify its various building blocks, Harness offers something much more open.

## Installation, in summary

For those who want to try it themselves, installation is really simple:

- Install **Node.js** if you haven't already
- Search for "DeepSeek Harness" on Google and go to the official site to copy the installation command
- Paste it into your terminal and run it
- Grab the generated local address and open it in your browser
- Enter your DeepSeek API key, create a workspace, choose your mode, and off you go

## My take

What strikes me about Harness isn't the promise of replacing Claude Code tomorrow morning — it clearly isn't, and DeepSeek doesn't even claim that. What matters is the architectural bet behind it: as models become more and more interchangeable, the real battle shifts to the layer you build around them — the tools, the permissions, the memory, the control loop. It's exactly the same underlying logic found in DeepSeek's open source strategy for V4: making the infrastructure open and tinker-friendly rather than locking everything into a finished product.

Let's be honest: today, for everyday use, a reliable and hassle-free coding agent, Claude Code is still largely ahead. But for developers who want to understand — and above all control — what happens under the hood of an agent, having an MIT framework where absolutely everything is a plugin is a real signal. One to watch in the coming months, especially if DeepSeek keeps up the update pace we already know it for on its models.

### 🛠️ Tools you can try related to this article

A selection of my tested tools, relevant for going further.
