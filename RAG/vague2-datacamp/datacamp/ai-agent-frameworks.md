---
id: vague2-datacamp/datacamp/ai-agent-frameworks
title: "Frameworks d’agents IA : bâtir des systèmes plus intelligents avec les bons outils"
domain: datacamp
role: reference
task: article
actors: ["Google", "Microsoft", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["agent", "agents", "cost", "memory", "open source", "reasoning", "research", "tool use"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/ai-agent-frameworks.md
source_anchor: ""
source_lines: [1, 69]
sha256: 4c9babe9ff9cc5f0228fad00b9257160fe217bbb685ec166200e25a55e449a82
---

# Frameworks d’agents IA : bâtir des systèmes plus intelligents avec les bons outils

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/ai-agent-frameworks
- **Site** : DataCamp
- **Type** : Article / Guide
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This guide explains AI agent frameworks: pre-equipped toolkits that handle memory, orchestration, and the essentials so teams don't build agents from scratch. AI agents are programs that can perceive, plan, and act—analyzing a goal, breaking it into steps, and acting alone or together. Building one from scratch requires managing memory, task planning, tool wiring, coordination, and error handling, which frameworks simplify.

Frameworks provide: **architecture** (agent interaction structure), **memory** (short/long-term recall), **models** (LLMs as the core for understanding/reasoning/action), **toolkits** (APIs, search engines, interpreters), an **orchestration layer** (task coordination and multi-agent collaboration), and **integrations** (LangChain, OpenAI, Azure, Slack, etc.).

An agent's decision cycle: **task planning** (decompose the request into steps), **function calling** (choose tools/APIs), **execution** (run code, fetch data, send email, write a response), and a **feedback loop** (evaluate results, update memory, adjust or ask for clarification). Systems can be **single-agent** (independent, one foundational model) or **multi-agent** (multiple collaborating/competing agents, potentially with different models). Common framework features: persistent memory, retrieval-augmented generation (RAG), tool use, and inter-agent collaboration.

The article profiles nine popular frameworks:
1. **CrewAI** — role-based ("Researcher", "Writer", "Analyst") teams; simple, open source, compatible with major LLMs and integrated RAG; ideal for customer-service bots or research-heavy automation; can feel constrained for highly flexible roles.
2. **LangGraph** — graph-theory-based, supports loops/branches/detours for non-linear workflows (travel assistants, insurance comparisons); integrated with LangChain; steep learning curve.
3. **AutoGen** (Microsoft) — layered architecture (Core, AgentChat, Extensions), with AutoGen Studio low-code environment; handles async messaging; useful with Azure/Teams; can feel over-engineered for simple projects.
4. **Agno (by Phidata)** — minimalist, clean Python, integrated cloud deployment; great for fast prototyping; smaller community.
5. **Atomic Agents** — clear, modular architecture with total control, no "magic" orchestration; for maintainability/performance (enterprise tools, R&D); not for quick prototypes.
6. **OpenAI Agents SDK** — production-ready evolution of Swarm, lightweight, provider-agnostic (OpenAI Responses/Chat Completions + 100+ LLMs); key idea is "handoffs"; good for privacy-sensitive local/on-prem apps.
7. **LlamaIndex** — began as GPT Index for structured data; excellent at extracting insights from documents/tables/databases and handling query transitions; weaker at orchestration than CrewAI/LangGraph.
8. **Semantic Kernel** (Microsoft) — enterprise framework compatible with Python, C#, Java, integrating LLMs into ERP/CRM; handles memory, planning, skills; best for strict existing-system integration; can be heavy for small teams.
9. **Google ADK (Agent Development Kit)** — open source for modular agents, strong orchestration and tool use, supports the "agent-as-tool" paradigm and evaluation workflows.

The article then offers selection criteria (ease of use, workflow complexity, customization, integration needs, deployment environment, performance/scalability, security/privacy) with a quick-reference table, and concludes there is no universal winner—start small, learn each framework's strengths, and combine building blocks (e.g., LangChain + LangGraph). Recommendations: CrewAI or Agno for speed without complexity; LangGraph or AutoGen for more control/power; Semantic Kernel for large enterprises with strict requirements.

## Key points

- AI agent frameworks are toolkits providing architecture, memory, models, toolkits, orchestration, and integrations.
- Agent decision cycle: task planning → function calling → execution → feedback loop.
- Single-agent systems use one model for defined goals; multi-agent systems coordinate specialized agents.
- Nine frameworks covered: CrewAI, LangGraph, AutoGen, Agno, Atomic Agents, OpenAI Agents SDK, LlamaIndex, Semantic Kernel, Google ADK.
- Best beginner-friendly: CrewAI, Agno; complex workflows: LangGraph, AutoGen; deep customization: Atomic Agents.
- Enterprise integration/security: Semantic Kernel; privacy/local: OpenAI Agents SDK.
- No universal winner; combine frameworks as needed; watch cost, bugs, security, and ethics at scale.

## Technical data / figures

| Factor | Best choices |
| --- | --- |
| Beginner-friendly | CrewAI, Agno |
| Complex workflows | LangGraph, AutoGen |
| Deep customization | Atomic Agents, LangGraph |
| API/tool integration | Semantic Kernel, CrewAI |
| Cloud deployment | Agno, Semantic Kernel |
| Local/privacy-focused | OpenAI's Agent SDK |
| Scale applications | AutoGen, LangGraph |

| Framework | Key trait |
| --- | --- |
| CrewAI | Role-based teams, simple, RAG included |
| LangGraph | Graph-based, loops/branches, LangChain ecosystem |
| AutoGen | Microsoft, layered, low-code Studio, async messaging |
| Agno | Minimalist Python, cloud deploy |
| Atomic Agents | Modular, full control, no hidden orchestration |
| OpenAI Agents SDK | Lightweight, provider-agnostic, handoffs |
| LlamaIndex | Data/document retrieval strength |
| Semantic Kernel | Enterprise (Python/C#/Java), ERP/CRM integration |
| Google ADK | Modular, agent-as-tool, evaluation workflows |

## Why this source matters for the RAG

It provides a broad, structured survey of the AI agent framework landscape with selection criteria and use-case fit, useful for architecture and tooling questions. It also explains agent fundamentals (decision cycle, single vs multi-agent) that anchor more advanced agent content.
