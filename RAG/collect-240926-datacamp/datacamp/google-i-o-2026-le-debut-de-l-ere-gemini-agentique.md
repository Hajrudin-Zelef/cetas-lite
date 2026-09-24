---
id: collect-240926-datacamp/datacamp/google-i-o-2026-le-debut-de-l-ere-gemini-agentique
title: "google-i-o-2026-le-debut-de-l-ere-gemini-agentique"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "Apple", "Google", "Meta", "Mistral", "Nvidia", "OpenAI", "United States"]
dates: []
keywords: ["agent", "gemini", "agentic", "agents", "agi", "benchmark", "benchmarks", "chatgpt", "claude", "compute", "consumer", "cost"]
source: docs/RAG/clean_en/datacamp/google-i-o-2026-le-debut-de-l-ere-gemini-agentique.md
source_anchor: ""
source_lines: [1, 146]
sha256: b24cae93538421aeafcf573391772d79fbcaa62f7527d6ba4068815cb8fe8e3c
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

- Integration with **Google tools (Workspace, Gmail, Calendar)** from the start, with third-party tool support via MCP in the coming weeks.
- Interaction via the **Gemini app**, then soon by email and chat.
- Operating directly in **Chrome** as an agentic browsing layer, expected later this summer.
- Live tracking of task progress via **Android Halo**, a new interface space on Android, planned for later this year.

The comparison with OpenAI's agent ecosystem and Anthropic's tool-use capabilities is inevitable. Spark's differentiator is persistent 24/7 execution on Google Cloud infrastructure, combined with deep integration into Google's productivity suite. If your work already lives in Google Workspace, that's a real plus. Otherwise, the value proposition is less obvious.

Privacy is a legitimate concern here. An agent that continuously monitors your inbox, calendar, and documents raises real questions about data residency and compliance in regulated industries. One of the questions I had, for example, was: "What happens to the agent's memory when an employee leaves?" Google has not yet provided detailed answers.

Spark is rolling out this week to trusted testers, with a beta for Google AI Ultra subscribers ($100/month) in the United States the following week. We cover it in detail in our article on Gemini Spark.

## Search agents and AI mode

AI mode in Search was introduced at the last I/O. A year later, it has surpassed one billion monthly active users. Google is going further with two new agentic capabilities.

The first: **information agents in Search**: personal background agents that you configure to monitor topics and surface relevant information at the right time. They will roll out this summer, first for Google AI Pro and Ultra subscribers.

The second: a **generative interface in Search**, powered by Gemini 3.5 Flash and Antigravity. Search will now build custom layouts, interactive visuals, and even persistent dashboards or mini-apps for complex, long-running queries. Generative interface capabilities will arrive free for everyone this summer. Persistent dashboards and custom apps will first be available to Pro and Ultra subscribers in the United States.

This legitimately worries publishers and SEO professionals (as AI Overview and AI mode had already done). When AI-generated answers fully resolve the query in Search, there's no longer any reason to click through to the source. We've already seen it: AI Overviews and the first rollout of AI mode led to marked traffic drops. Google still hasn't proposed a clear revenue-sharing mechanism or traffic guarantee for publishers whose content feeds these answers.

## Google Flow

Google Flow, introduced at I/O 2025 as an AI filmmaking tool, is taking a major step forward with three key developments:

- **Smarter planning agent.** The updated Flow agent can now plan and reason about multi-step creative projects. You provide your inputs (for example a concept, reference images, a draft script) and it helps you go from brainstorming to creation and then to editing, all within the same environment. The new agent is available today for everyone.
- **Native video with Gemini Omni.** Flow now handles video generation and editing natively via the Omni model. You can describe in natural language the changes to make to a clip in your reel and iterate conversationally. Character consistency also improves: identity and voice are preserved across scenes, which is useful for a short film or a campaign with recurring characters.
- **Vibe coding for custom tools.** Instead of being limited to the tools shipped with Flow, you can now create your own directly on the platform. Google showed examples like custom video effects, hand-drawn animation tools, or text overlay workflows, without leaving Flow.

Together, these updates position Flow beyond a simple creative assistant. It is becoming a platform for building creative workflows, with a mobile app now in beta on Android and soon on iOS.

## SynthID expansion

SynthID, Google's invisible watermarking system for AI, has already marked more than 100 billion images and videos and 60,000 cumulative years of audio assets since its launch three years ago. The most important announcement at I/O isn't the scale, but the partners: OpenAI, Kakao, and Eleven Labs are adopting SynthID alongside Nvidia, which signed on last year.

Cross-industry adoption makes all the difference. A watermarking standard is only useful if it is widespread enough for "unwatermarked" to become a meaningful signal. Google is also extending Content Credentials verification (the C2PA standard) to Search and Chrome, in order to indicate whether content comes from an AI or a camera and whether it has been edited with generative tools. The combination of SynthID + C2PA adds two independent layers of provenance: a relevant approach, knowing that it is easy to remove either one in isolation.

## Honorable mentions

Several other I/O announcements deserve a quick look:

- **Docs Live:** a new voice feature for Google Docs that lets you pour out your ideas verbally and let Gemini structure them into a document. Rolling out this summer for subscribers, with voice features also planned for Gmail and Keep at the same time.
- **Google Pics:** a new AI image creation and editing tool, built on the Nano Banana model, which treats each element as an individual object rather than a flat image. Available now for trusted testers, rolling out this summer for Google AI Pro and Ultra subscribers.
- **Android Halo:** a new interface space on Android to view in real time the updates and progress of agent tasks like Gemini Spark. Arriving later this year.
- **Daily Brief:** a ready-to-use agent in the Gemini app that compiles a personalized morning summary from your inbox, calendar, and tasks, with suggested next steps. No separate pricing announced; expected as an integral part of the Gemini app experience.
- **TPU 8t and 8i:** Google's 8th-generation TPUs adopt a dual-chip approach: 8t optimized for large-scale pre-training (nearly 3 times the raw compute power of the previous generation, scalable to more than 1 million TPUs worldwide) and 8i optimized for inference. Both offer up to 2 times more performance per watt than the previous generation.
- **Gemini for Science:** a set of AI tools connecting Antigravity to more than 30 major life sciences databases. Science Skills is available today on GitHub and directly in Antigravity.

## Final thoughts

Google I/O 2026 is betting on agents as the main trajectory of AI, with Gemini 3.5 Flash and Antigravity 2.0 as the underlying infrastructure for almost everything else. What you can use right now: Gemini 3.5 Flash (via the Gemini API and AI Studio), the new Flow agent, Gemini Omni Flash, and the Antigravity 2.0 desktop app. Gemini Spark, the Search agents, and the generative Search interface will arrive over the summer, mainly reserved for the new AI Ultra offering at $100/month (at least at launch).

For me, the Antigravity upgrade is the most interesting, because it **operates on two levels simultaneously: as a standalone application for developers, it competes directly with Codex and Claude Code; as a platform, its underlying ADK and its Managed Agents API challenge orchestration frameworks like LangChain, AutoGen, and OpenAI's Agents SDK. The integration with Gemini and the Google Cloud deployment layer are the differentiators (and the lock-in risk) on both fronts.**

## Google I/O 2026: FAQ

### How does Gemini 3.5 Flash compare to GPT-5.5 and Claude Opus 4.7?

Gemini 3.5 Flash leads on several agentic benchmarks such as MCP Atlas (83.6%) and Finance Agent v2 (57.9%), while GPT-5.5 is ahead on SWE-Bench Pro and ARC-AGI-2. Claude Opus 4.7 remains the best on Humanity's Last Exam (46.9%). Key takeaway: it rivals cutting-edge models while running faster and much cheaper at scale, as its name suggests. A more powerful Pro variant is expected soon.

### How does Google Antigravity differ from Claude Code or Codex?

Google Antigravity 2.0 is an agent-centric development platform that lets you orchestrate multiple AI agents in parallel via a desktop app, a CLI, an SDK, and an enterprise API. Unlike Claude Code (a terminal-oriented coding agent) or Codex (a task-queue-based system), Antigravity offers finer-grained per-project permission scoping, sub-agent creation, and direct integration with Google Cloud and Firebase. Its dual role as a developer tool and a platform SDK makes it closer to an orchestration framework than a simple coding assistant.

### Is the $100/month Google AI Ultra subscription worth it compared to ChatGPT Pro or Claude Max?

All three plans are priced at $100/month, but the value depends on your ecosystem. Google AI Ultra's differentiator is access to Gemini Spark (a persistent 24/7 agent), a 5x higher Antigravity usage quota, and deep integration with Google Workspace. If your workflow already relies on Gmail, Docs, and Calendar, Ultra has a natural advantage. If you're mainly looking for coding help or API-level flexibility, ChatGPT Pro or Claude Max may suit you better.

### What is Gemini Omni and how does it handle video generation?

Gemini Omni is Google's natively multimodal model that accepts any mix of text, images, audio, and video as input, and produces video output. It unifies previously separate systems (Veo for video, Imagen for images) into a single model, which should improve editing consistency across modalities. The first version, Omni Flash, is available now, with a more powerful Omni Pro expected soon. No independent benchmarks have been published yet, so real-world quality remains to be evaluated.

**Data Science Editor-in-Chief at DataCamp |** **I am passionate about forecasting and development using APIs.**
