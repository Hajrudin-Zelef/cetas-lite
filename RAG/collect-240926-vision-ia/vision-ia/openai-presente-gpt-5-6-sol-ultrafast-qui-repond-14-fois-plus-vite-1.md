---
id: collect-240926-vision-ia/vision-ia/openai-presente-gpt-5-6-sol-ultrafast-qui-repond-14-fois-plus-vite-1
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AWS", "Alibaba", "Ant", "Anthropic", "Cerebras", "China", "DeepSeek", "Google", "Hugging Face", "Moonshot", "Nvidia", "OpenAI"]
dates: ["2026-03", "2026-12-31", "2027-01-01"]
keywords: ["research", "agent", "agents", "attention", "bedrock", "benchmark", "chatgpt", "claude", "compute", "context window", "cost", "cost per token"]
source: docs/RAG/clean_en/vision-ia/openai-presente-gpt-5-6-sol-ultrafast-qui-repond-14-fois-plus-vite.md
source_anchor: ""
source_lines: [1, 72]
sha256: f2eddf93769e6fbbba170d979bbcecad41ba444c4a75acde2d3d4f50d9ff5c6f
---

# 🧠 **RESEARCH**

<!-- source: https://vision-ia.beehiiv.com/p/openai-pre-sente-gpt-5-6-sol-ultrafast-qui-re-pond-14-fois-plus-vite -->

OpenAI opens a new tier of its API in preview, called Ultrafast, which runs GPT-5.6 Sol up to **14 times faster** than the standard version, at **750 output tokens per second**. This is not a lighter or distilled version: it's the same GPT-5.6 Sol, with the same level of intelligence, simply placed on a different type of chip.

**Key takeaways:**

- The gain doesn't come from software but from hardware. The **Cerebras** Wafer-Scale Engine chips pack **44 GB of SRAM** each and keep the model's weights in local memory, instead of fetching them from external storage with every token generated. The classic bottleneck of GPU inference, memory bandwidth, disappears.

- On Humanity's Last Exam and its **2,500 questions**, Ultrafast completes the test in **11h11**, versus **78h27** for Claude Fable 5, according to measurements published by Cerebras.

- Same source: **5.6 times** faster on the GDP-Val benchmark, **5 times** the speed of Opus 4.8 in Fast mode.

- Access open since **August 13** to a limited group of API customers, with gradual expansion depending on available capacity. It's not in ChatGPT yet, and pricing hasn't been disclosed.

- The use cases highlighted by OpenAI: incident response, customer support, financial market analysis, e-commerce. Anything that handles waiting poorly.

**Why it matters:** for two years, the race was about model intelligence, and waiting in front of a blinking cursor was part of the scenery. Here, nothing changes on the reasoning side, only latency collapses. This is precisely what was missing for an AI agent to go from background task to real-time tool: an assistant that responds in one second isn't used like an assistant that responds in fifteen. Incidentally, the compute no longer comes from Nvidia, and this is the first deployment of this scale at a leading lab.

DeepSeek releases in developer preview **Harness**, the open-source framework under MIT license on which its own agents run, source code included. The principle boils down to an equation put forward by the company: "Agent = Model + Harness." The model provides the intelligence, the harness handles everything else, meaning the environment, tools, sessions, and the ability to work for a long time without derailing.

**In detail:**

- Everything is a plugin: models, tools, skills, sessions, sandboxes, storage, loops, scheduling, and even the interface. A kernel called **Cordis** handles the mounting, unmounting, and dependencies between these building blocks, which can be swapped by changing the configuration, without ever touching the source code.

- **Four execution modes**: Standard (all the tooling), Code (the agent writes code to orchestrate multiple rounds of tool calls), Minimal (a shell and a file editor, to evaluate a model without crutches), and Creator (live runtime inspection, in-memory plugin testing, building new modes).

- Each session produces an append-only event log that records everything the model sees: system prompts, reasoning, tool calls and results, sub-agent scheduling, context injections. You can resume, duplicate, search, and replay an execution from that same stream.

- **Compatible with any model**: DeepSeek, Anthropic, OpenAI, Bedrock, Vertex, Azure, and any OpenAI-compatible endpoint.

- You can try it right away: install Node.js, run `npx @deepseek-ai/dsh web`, and the interface runs locally. The full repository is on GitHub under `deepseek-ai/deepseek-harness`.

**The context:** this is the first serious open-source rival to Claude Code, and it arrives on the same day DeepSeek also releases its **V4-Pro** model, whose API prices increase on **August 16** with differentiated peak and off-peak pricing in China. The topic is generating a lot of reaction in the tech community, less for raw performance that hasn't yet been quantified than for what it signifies: the hard part of an agent is no longer the model, it's the plumbing around it, and DeepSeek has just put it on the table.

Google DeepMind releases Gemini 3.7 Flash **just three weeks** after Gemini 3.6 Flash, and launches it at **half the price** of its predecessor. This multimodal model (text, image, audio, video) targets coding, web development, and agent workflows, with clear progress across all published evaluations.

**A few key figures:**

- Introductory pricing until December 31, 2026: **$0.75 per million input tokens**, **$3.75 for output**. On January 1, 2027, it goes to $1.50 and $7.50.

- Context window of **1 million input tokens**, output capped at **64,000 tokens**, adjustable thinking effort to balance quality, cost, and latency. Knowledge cutoff March 2026.

- Versus 3.6 Flash: FrontierCode 1.1 at **43.6%** versus 34.4%, DeepSWE v1.1 at **65.3%** versus 49.0%, WebDev Arena at **1588 Elo** versus 1538, GDP.pdf (complex document processing) at **34.0%** versus 22.0%, AutomationBench (real business workflows) at **30.4%** versus 17.0%.

- It doesn't win everywhere: on Terminal-bench, GPT-5.6 Terra keeps the lead with **87.4%** versus 85.8%.

- Where to try it: Google AI Studio, the Gemini API, Android Studio, Antigravity, and starting today in Gemini Spark for Google AI Pro and Ultra subscribers in more than **160 countries**.

Google is accompanying the launch with compelling demos: a playable 3D game generated from a simple text prompt, with Nano Banana creating characters, objects, and textures in real time, or an annual report in PDF transformed into an interactive web page with live charts.

**What it changes:** the pace, above all. Three weeks between two generations, with the price halved and scores up by ten to fifteen points, that means one very concrete thing for you: the cost of "good enough" AI is falling faster than its performance is rising. What was expensive to automate a month ago costs almost nothing today.

inclusionAI, the AI lab of Ant Group, is releasing Ling 3.0 Flash under the MIT license, with weights downloadable on Hugging Face. It is a Mixture-of-Experts model: **124 billion parameters in total**, but only **5.1 billion active** for each token generated. In other words, the compute cost of a small model with a good portion of the knowledge of a large one.

**The key points:**

- Hybrid architecture BailingMoeV3: **512 experts**, 8 of which are activated per token plus a shared expert, Kimi Delta linear attention combined with Multi-head Latent Attention, native context window of **256,000 tokens**.

- **38 points** on the Artificial Analysis Intelligence Index, on par with Qwen3.6 27B, but still far from the open leader DeepSeek V4 Flash and its 52 points. No smaller model reaches its score.

- The real progress is elsewhere: on the AA Omniscience test, the hallucination rate drops from **97% to 44%** compared with the previous version. The model now refuses far more often to answer when it does not have a reliable answer.

- Gains are also clear on agent tasks, notably the t3-Bench Banking benchmark.

- Available in fp8, fp4, int4 and **GGUF** on Hugging Face, hosted on DeepInfra and the inclusionAI API, free for now on Kilo Code. Cost per token lower than Qwen3.6 27B, even if it consumes more tokens on complex tasks.

**The impact to remember:** the GGUF format plus 5 billion active parameters is exactly the combination that makes it possible to run this model on a properly equipped personal machine, without a subscription, without an API, without your data leaving your home. And the number that matters is not the benchmark score but the hallucination rate cut by more than half in one generation: a model that knows how to say "I don't know" is infinitely more useful in daily life than a model that gains three points on a test.

# 🧠 **RESEARCH**

**Three Claude agents on the same project, and it's a turf war**

