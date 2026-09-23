---
id: vague2-datacamp/datacamp/google-i-o-2026
title: "Google I/O 2026 : le début de l'ère Gemini agentique"
domain: datacamp
role: reference
task: article
actors: ["Anthropic", "Google", "Nvidia", "OpenAI", "United States"]
dates: ["2026-09-23"]
keywords: ["agent", "gemini", "agentic", "agents", "agi", "benchmark", "benchmarks", "chatgpt", "claude", "compute", "inference", "mcp"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/google-i-o-2026.md
source_anchor: ""
source_lines: [1, 64]
sha256: 3618d4add1c3440811a196ad2bf35e456ca07b8dc6f2eb28143ecef892b13f3f
---

# Google I/O 2026 : le début de l'ère Gemini agentique

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/google-i-o-2026
- **Site** : DataCamp
- **Type** : Article (actualité / keynote recap)
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp recap covers day one of Google I/O 2026, whose unifying theme was **agents** — persistent, background-executing agents integrated across Google's product stack. CEO Sundar Pichai called it "the agentic Gemini era." Google also introduced a **Google AI Ultra** tier at $100/month to counter Anthropic's Claude Max and OpenAI's ChatGPT Pro.

**Gemini 3.5 Flash** is the flagship model release, beating Gemini 3.1 Pro on agentic and coding benchmarks while claimed to be 4x faster in output tokens/second. It leads MCP Atlas (83.6%), Finance Agent v2 (57.9%), and CharXiv Reasoning (84.2%), competing with Claude Opus 4.7 and GPT-5.5. Google claims enterprises processing ~1 trillion tokens/day could save over $1B/year by moving 80% of workloads to 3.5 Flash. Available now via Gemini API, AI Studio, and the Gemini app; Gemini 3.5 Pro is used internally and expected next month.

**Gemini Omni** is a natively multimodal media-generation model taking any mix of text, images, audio, and video, producing video output. It merges previously separate stacks (Veo for video, Imagen for images, separate audio systems) into one model for more coherent edits. Gemini Omni Flash is available in the Gemini app, Google Flow, and YouTube Shorts; API access follows. No benchmarks published; a stronger Omni Pro is expected.

**Antigravity 2.0** expands Google's agent-centric dev platform from a coding environment into a full platform for developing, deploying, and managing autonomous agent cohorts, with a standalone desktop app as orchestration hub. Four surfaces: desktop app, CLI, SDK, and integration in Gemini Enterprise Agent Platform. New capabilities: on-the-fly modular subagents with workspace isolation inheriting parent tools/permissions; async long operations; **JSON Hooks** (shell scripts at lifecycle points); **scheduled tasks** (cron prompts); "projects" as an organizational unit for scoped permissions; native Git worktrees; voice input; new slash commands (`/goal`, `/grill-me`, `/schedule`, `/browser`). Available today; AI Ultra gives 5x the Antigravity usage of AI Pro.

**Managed Agents in the Gemini API** bring agentic capabilities into the API layer, letting developers define agent behaviors, tool integrations, and multi-step flows while Google handles execution. Available via AI Studio and Gemini Enterprise Agent Platform. Caveat: early developer feedback notes sparse documentation for complex flows/error handling and friction with rate limits/quotas.

**Gemini Spark** is Google's personal agent running 24/7 on dedicated VMs in Google Cloud (no need to keep your computer open), powered by Gemini 3.5 and the Antigravity harness. Launch features: Google Workspace/Gmail/Calendar integration with third-party MCP tools soon; interaction via Gemini app then email/chat; Chrome browsing layer later in summer; progress tracking via Android Halo. Rolled out to trusted testers, then US AI Ultra beta. Privacy concerns remain (data residency, agent memory when an employee leaves).

**Search agents and AI mode:** AI Mode passed 1 billion MAU. Two new agentic capabilities: information agents (background agents monitoring topics, summer for AI Pro/Ultra) and a generative Search interface powered by Gemini 3.5 Flash and Antigravity, building custom layouts, interactive visuals, persistent dashboards, and mini-apps (generative UI free for all this summer; persistent dashboards first for US Pro/Ultra). SEO/publisher traffic concerns persist.

**Google Flow** gets a smarter planning agent, native video with Gemini Omni (conversational editing, character/voice consistency), and vibe coding for custom tools; mobile beta on Android.

**SynthID** has watermarked 100B+ images/videos and 60,000 years of audio; new partners include OpenAI, Kakao, and Eleven Labs (plus Nvidia). Content Credentials (C2PA) verification extends to Search and Chrome.

**Honorable mentions:** Docs Live (voice), Google Pics (Nano Banana-based image tool), Android Halo, Daily Brief, **TPU 8t/8i** (8th-gen dual-chip: 8t for pre-training ~3x prior raw compute, scalable to 1M+ TPUs; 8i for inference; up to 2x performance/watt), and Gemini for Science (30+ life-science databases; Science Skills on GitHub and Antigravity).

The author finds Antigravity's upgrade most interesting because it competes on two levels: as a standalone app against Codex and Claude Code, and as a platform (ADK + Managed Agents API) against LangChain, AutoGen, and OpenAI Agents SDK.

## Key points

- I/O 2026 theme: "agentic Gemini era" — persistent background agents across Google's stack.
- Gemini 3.5 Flash: top on MCP Atlas, Finance Agent v2, CharXiv; claimed 4x faster output tokens; potential $1B/year savings at scale.
- Gemini Omni unifies Veo/Imagen/audio into one multimodal model producing video.
- Antigravity 2.0 becomes a full agent platform (desktop, CLI, SDK, enterprise) with subagents, JSON Hooks, cron tasks, project scoping, Git worktrees.
- Managed Agents API brings agent orchestration to the Gemini API layer.
- Gemini Spark runs 24/7 on Google Cloud VMs, deeply integrated with Workspace.
- Google AI Ultra launched at $100/month, matching Claude Max and ChatGPT Pro.
- TPU 8t/8i 8th-gen chips target pre-training and inference with up to 2x performance/watt.

## Technical data / figures

| Benchmark | 3.5 Flash | 3 Flash | 3.1 Pro | Claude Sonnet 4.6 | Opus 4.7 | GPT-5.5 |
|---|---|---|---|---|---|---|
| Terminal-bench 2.1 | 76.2% | 58.0% | 70.3% | -- | 66.1% | **78.2%** |
| SWE-Bench Pro | 55.1% | 49.6% | 54.2% | -- | 64.3% | **58.6%** |
| MCP Atlas | **83.6%** | 62.0% | 78.2% | 69.5% | 79.1% | 75.3% |
| OSWorld | 78.4% | 65.1% | 76.2% | 72.5% | 78.0% | **78.7%** |
| Finance Agent v2 | **57.9%** | 42.6% | 43.0% | 51.0% | 51.5% | 51.8% |
| CharXiv Reasoning | **84.2%** | 80.3% | 83.3% | 72.4% | 82.1% | 84.1% |
| Humanity's Last Exam | 40.2% | 33.7% | 44.4% | 33.2% | **46.9%** | 41.4% |
| ARC-AGI-2 | 72.1% | 33.6% | 77.1% | 58.3% | 75.8% | **84.6%** |

Other figures: Google AI Ultra $100/mo; AI Ultra gives 5x Antigravity usage vs AI Pro; AI Mode 1B+ MAU; SynthID 100B+ images/videos, 60,000 years audio; TPU 8t ~3x prior raw compute, up to 2x perf/watt, 1M+ TPU scale.

## Why this source matters for the RAG

It is a comprehensive, current record of Google's 2026 agent-first announcements, model benchmarks, and pricing, essential for questions on Gemini models, Antigravity, agent platforms, and the competitive AI landscape.
