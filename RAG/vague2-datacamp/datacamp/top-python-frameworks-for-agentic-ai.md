---
id: vague2-datacamp/datacamp/top-python-frameworks-for-agentic-ai
title: "Top 15 frameworks Python pour créer des applications d'IA agentique complètes"
domain: datacamp
role: reference
task: article
actors: ["Anthropic", "Google", "Hugging Face", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["agent", "agentic", "agents", "claude", "cost", "gemini", "guardrails", "inference", "mcp", "memory", "open source", "research"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/top-python-frameworks-for-agentic-ai.md
source_anchor: ""
source_lines: [1, 62]
sha256: a3d6c6a9a2137c491a5970b3e0ef0ec25b0c9bf81105ad04a47b32accb55f846
---

# Top 15 frameworks Python pour créer des applications d'IA agentique complètes

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/top-python-frameworks-for-agentic-ai
- **Site** : DataCamp
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp article surveys 15 Python frameworks for building agentic AI applications, organized into five categories, with a short copy-paste Python example for each. The author (Abid Ali Awan) draws on hands-on experience building RAG apps, autonomous agents, multi-agent systems, and deployable AI tools.

**General-purpose agent frameworks.** *LangChain* connects LLMs to external data, tools, APIs, and app logic, now a full ecosystem including RAG, tool-using agents, workflows, integrations, monitoring, and evaluation; example uses `create_agent`. *LangGraph* structures apps as a graph of connected steps for controlled loops, tool calls, shared state, human validation, and multi-agent workflows, with persistence, streaming, checkpoints, and resume. *Agno* builds agents, multi-agent teams, and structured workflows, plus AgentOS for exposing agents as APIs and storing sessions/traces, self-hostable for data control. *Pydantic AI* offers a Python-native, type-safe experience (FastAPI-like) with Pydantic validation of tool arguments and structured outputs.

**Official provider SDKs.** *OpenAI Agents SDK* is a lightweight framework with a built-in agent loop, function tools, handoffs, guardrails, sessions, human-in-the-loop, and tracing visible in the OpenAI Dashboard. *Google Agent Development Kit (ADK)* is open source for building, evaluating, and deploying agents, best with Gemini but multi-provider. *Claude Agent SDK* (Python/TypeScript) reuses Claude Code's agent loop, tools, and context management for agents that read files, edit code, run commands, and use MCP tools.

**Multi-agent orchestration.** *CrewAI* defines role-based agents and crews with sequential or hierarchical processes. *MetaGPT* simulates a software company (PM, architect, project manager, engineer) to turn a brief into plans and code. *AgentScope* provides controllable, observable agents with ready-made agents, memory, parallel tool calls, tracing, and Studio visualization. *CAMEL-AI* is research-oriented for multi-agent societies, role-play, and simulation, supporting local models via Ollama, vLLM, and SGLang.

**Data/RAG frameworks.** *LlamaIndex* builds context-aware apps over private data (connectors, indexes, retrievers, query engines, agents, workflows). *Haystack* provides modular, production-ready RAG pipelines with explicit control over retrieval, prompting, and data flow.

**Lightweight/open-model.** *Hugging Face smolagents* is a minimal framework with CodeAgent (writes Python) and ToolCallingAgent, working well with Hugging Face models, Inference Providers, and local open models.

The article closes with a comparison table and advice: start with the simplest framework that meets the need, use official SDKs for quick starts, LangChain/LangGraph for flexibility, CrewAI for role teams, LlamaIndex/Haystack for RAG, and smolagents for lightweight open-model use.

## Key points

- 15 frameworks across 5 categories: general-purpose, official SDKs, multi-agent orchestration, data/RAG, and lightweight/open-model.
- LangChain = broad integrations; LangGraph = stateful graph orchestration and durable workflows.
- Official SDKs (OpenAI Agents, Google ADK, Claude Agent SDK) give fastest integration and tracing.
- CrewAI, MetaGPT, AgentScope, CAMEL-AI target multi-agent coordination and simulation.
- LlamaIndex and Haystack are the data/RAG specialists.
- smolagents targets minimal abstraction and local/open models.
- Agents cost more than single prompts: a single request can trigger 5-6 LLM calls due to looping.
- Evaluate non-deterministic agents with LLM-as-a-judge tools (Ragas, TruLens, DeepEval).

## Technical data / figures

| Framework | Main strength | Choose when |
|---|---|---|
| LangChain | Broad integration ecosystem | Need extensive integrations and flexibility |
| LangGraph | Stateful graph orchestration | Need controlled loops and durable workflows |
| Agno | Full agent platform (AgentOS) | Want to build, self-host, manage agent services |
| Pydantic AI | Type-safe Python development | Need validated tools/dependencies and structured output |
| OpenAI Agents SDK | Simple agent dev + tracing | Primarily using OpenAI models |
| Google ADK | Gemini-native development | Using Gemini/Google Cloud or multi-agent workflows |
| Claude Agent SDK | Computer/file/command tools | Building coding, research, or long-running agents with Claude |
| CrewAI | Role-based agent teams | Multiple specialized agents collaborate |
| MetaGPT | Software company simulation | Agents plan and generate full software projects |
| AgentScope | Observable, controllable agents | Need traceable agents/multi-agent apps |
| CAMEL-AI | Multi-agent research and simulation | Studying role-play or large agent systems |
| LlamaIndex | Data-connected, context-aware apps | Agents retrieve/reason over private documents |
| Haystack | Modular, transparent RAG pipelines | Want control over retrieval and prompting |
| smolagents | Lightweight agents for open models | Minimal abstraction, code agents, local models |

Evaluation tools mentioned: **Ragas**, **TruLens**, **DeepEval**. Metrics: factual grounding, tool-selection accuracy, response relevance.

## Why this source matters for the RAG

It is a compact, practical map of the Python agentic-AI framework ecosystem, including selection guidance and cost/evaluation caveats. It supports RAG queries comparing frameworks, SDK choices, RAG-vs-agentic-RAG distinctions, and implementation patterns.
