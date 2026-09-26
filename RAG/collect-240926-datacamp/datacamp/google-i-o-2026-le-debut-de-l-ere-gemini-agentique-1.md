---
id: collect-240926-datacamp/datacamp/google-i-o-2026-le-debut-de-l-ere-gemini-agentique-1
title: "google-i-o-2026-le-debut-de-l-ere-gemini-agentique"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "Google", "Meta", "OpenAI"]
dates: []
keywords: ["agent", "gemini", "agentic", "agents", "agi", "benchmark", "benchmarks", "chatgpt", "claude", "consumer", "cost", "mcp"]
source: docs/RAG/clean_en/datacamp/google-i-o-2026-le-debut-de-l-ere-gemini-agentique.md
source_anchor: ""
source_lines: [1, 73]
sha256: 4c4502849e55d466d1c82822da039576fcaf789272c70e1be028ee1e81999a97
---

# google-i-o-2026-le-debut-de-l-ere-gemini-agentique

<!-- source: https://www.datacamp.com/fr/blog/google-i-o-2026 -->

Course

The common thread running through almost all the announcements on the first day of the Google I/O 2026 conference was the same: agents. Not chatbots, nor simple assistants, but persistent agents, executing tasks in the background and integrated across Google's entire product stack. CEO Sundar Pichai unambiguously named it "the agentic Gemini era," and the announcements largely confirmed this framing.

Google also responded to Anthropic's Claude Max and OpenAI's ChatGPT Pro subscriptions by introducing a new Google AI Ultra offering at I/O, priced at the same $100 per month and giving access to some of the agentic features detailed below.

In this article, I present the announcements that matter most to AI practitioners and developers. I focus on updates available now or expected very soon.

## Gemini 3.5 Flash

Gemini 3.5 Flash is the flagship model release of I/O 2026. It outperforms Gemini 3.1 Pro on agentic and coding benchmarks, while being, according to Google, 4 times faster in output tokens per second than other frontier models. We cannot yet confirm this, but the promise is bold.

On the benchmarks side, the progress is visible, particularly on MCP Atlas, CharXiV Reasoning, and Finance Agent v2, where Gemini 3.5 Flash takes the lead. Overall, 3.5 Flash appears to rival Claude Opus 4.7 and GPT-5.5.

| **Benchmark** | **3.5 Flash** | **3 Flash** | **3.1 Pro** | **Claude Sonnet 4.6** | **Opus 4.7** | **GPT-5.5** | 
| Terminal-bench 2.1 | 76.2% | 58.0% | 70.3% | -- | 66.1% | **78.2%** | 
| SWE-Bench Pro | 55.1% | 49.6% | 54.2% | -- | 64.3% | **58.6%** | 
| MCP Atlas | **83.6%** | 62.0% | 78.2% | 69.5% | 79.1% | 75.3% | 
| OSWorld | 78.4% | 65.1% | 76.2% | 72.5% | 78.0% | **78.7%** | 
| Finance Agent v2 | **57.9%** | 42.6% | 43.0% | 51.0% | 51.5% | 51.8% | 
| CharXiv Reasoning | **84.2%** | 80.3% | 83.3% | 72.4% | 82.1% | 84.1% | 
| Humanity's Last Exam | 40.2% | 33.7% | 44.4% | 33.2% | **46.9%** | 41.4% | 
| ARC-AGI-2 | 72.1% | 33.6% | 77.1% | 58.3% | 75.8% | **84.6%** | 

The cost aspect is worth noting. Google claims that companies processing around 1 trillion tokens per day could save more than $1 billion per year by moving 80% of workloads from other frontier models to 3.5 Flash. This is a message aimed directly at OpenAI's and Anthropic's enterprise customers. Gemini 3.5 Flash is available today via the Gemini API, Google AI Studio, and the Gemini app. Gemini 3.5 Pro is already being used internally and is expected next month.

To learn more, we recommend reading our article on Gemini 3.5 Flash, where we detail the new model.

## Gemini Omni

Gemini Omni is Google's new natively multimodal media generation model, capable of taking any combination of text, images, audio, and video as input, and producing video output. The first model in the family, Gemini Omni Flash, is available today in the Gemini app, Google Flow, and YouTube Shorts.

Key architectural point: Omni merges what was previously a separate stack (Veo for video, Imagen for images, distinct audio systems) into a single model. The result: more consistent edits and fewer pipeline artifacts during multi-modal work. Google did not publish numerical benchmarks for Omni at launch, so independent evaluations are still to come. API access for developers and enterprise customers is coming in the weeks following I/O.

We tested it and detailed it in our article on Gemini Omni. The first video generation results are uneven (especially against the very high standards set by tools like Seedance 2.0), but a more powerful Gemini Omni Pro is expected soon.

## Antigravity 2.0

Antigravity is Google's agent-centric development platform, and version 2.0 presented at I/O marks a major expansion. Once positioned as a coding environment, it is now a complete platform for developing, deploying, and managing cohorts of autonomous AI agents. The centerpiece is a new standalone desktop application serving as a central hub for orchestration: you can run multiple agents in parallel on distinct tasks, simultaneously.

The ecosystem now has four distinct surfaces for developers:

- **Antigravity 2.0 Desktop App:** orchestrates multiple agents in parallel and manages scheduled background tasks. Integrates with Google AI Studio, Android, and Firebase.
- **Antigravity CLI:** native terminal interface for creating and running agents without a GUI. Google invites Gemini CLI users to migrate.
- **Antigravity SDK:** programmable access to the same agent harness that powers Google's products, with support for custom agent behaviors hosted on your own infrastructure.
- **Antigravity in Gemini Enterprise Agent Platform:** connects Antigravity directly to Google Cloud projects for enterprise workloads.

The agent core also gains several very useful functions. The most notable: it can now generate **modular sub-agents** on the fly, each running in parallel with workspace isolation, while inheriting the parent's tools and permissions. Long operations run asynchronously and no longer block the agent loop.

In the spirit of Claude Code Hooks, **JSON Hooks** allow attaching custom shell scripts to key steps (before/after tool calls, model calls, or stopping conditions) for logging, argument adjustment, or instruction injection. **Scheduled tasks** allow defining cron-based prompts for periodic agent runs like daily PR summaries or hourly deployment checks, with results shown in the sidebar for smooth handoff with the human in the loop.

On the administration side, Antigravity introduces "projects" as the organizational unit to scope settings, resources, and permissions per agent group, avoiding overly broad global permissions. The redesigned sidebar allows grouping conversations by project, status, or recency, with **native Git worktree support**. This project-based organization echoes Cursor's multi-window management and Codex's task queue, with finer-grained permission scoping per project.

**Voice input** via Gemini audio models and **new slash commands** (`/goal` for autonomous runs, `/grill-me` to clarify before a task, `/schedule` for cron prompts, `/browser` to enable the browser) round out the experience.

**A**ntigravity 2.0 is available starting today. The Google AI Ultra plan ($100/month) includes a usage quota 5 times higher in Antigravity compared to the Google AI Pro plan.

## Managed Agents in the Gemini API

Alongside Antigravity 2.0, Google announced Managed Agents in the Gemini API, which bring agentic capabilities directly into the API layer for developers wanting to build agent-powered applications without managing orchestration themselves. This is the API-side counterpart to the Antigravity desktop experience.

Concretely, you can define agent behaviors, tool integrations, and multi-step flows via the Gemini API, and Google's infrastructure handles execution. Potentially a real step change for teams building production applications requiring long-running tasks, without deploying their own agent harness. Access is available via Google AI Studio, and enterprise customers access it via the Gemini Enterprise Agent Platform.

An honest caveat: early developer feedback since I/O indicates that documentation for complex agentic flows and error handling remains sparse. Rate limits and quota management are also cited as sources of friction. These points should ease over time, but you should be aware of them before committing to this stack.

## Gemini Spark

Gemini Spark (not to be confused with Meta's latest LLM, Muse Spark) is Google's new personal agent, and it's the most consumer-oriented agentic announcement. It runs 24/7 on dedicated virtual machines in Google Cloud, without requiring your computer to stay open. Spark is powered by Gemini 3.5 and the Antigravity harness, which allows it to handle long-duration tasks in the background.

Launch features include:

