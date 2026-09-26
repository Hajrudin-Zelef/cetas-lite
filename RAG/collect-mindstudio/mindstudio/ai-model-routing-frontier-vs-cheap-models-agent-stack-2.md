---
id: collect-mindstudio/mindstudio/ai-model-routing-frontier-vs-cheap-models-agent-stack-2
title: "AI Model Routing: When to Use Frontier Models vs Cheap Models in Your Agent Stack"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["agent", "agents", "claude", "compute", "cost", "latency", "pricing", "reasoning", "voice"]
source: docs/RAG/Collect RAG/02_mindstudio/ai-model-routing-frontier-vs-cheap-models-agent-stack.md
source_anchor: ""
source_lines: [26, 54]
sha256: 725f95189a06505bacc83c578ff8a7af58e11713f73cff16318a2e45b54cee6d
---

# AI Model Routing: When to Use Frontier Models vs Cheap Models in Your Agent Stack

Building a routing layer in practice: (1) define a task taxonomy — list the distinct task types and make an initial assignment per type (default cheap, default frontier, or "evaluate before routing"); (2) build a router — often a small prompt that takes the incoming task and outputs a routing decision; cheap models work fine since the routing task itself is well-defined; use rule-based logic (regex, token count thresholds, input type detection) to avoid LLM calls entirely for obvious cases; (3) add fallback logic — define what "good enough" means per task type, build validation checks (structured output parsing, confidence thresholds, output length checks), and route to frontier if the cheap model's output fails; (4) log everything — track which model handled each task, whether fallback triggered, and output quality; this data tells you whether routing logic is working; (5) iterate based on cost-quality data after a week of production data — tasks routed to frontier that almost never fail cheap checks can shift down; cheap tasks with high fallback rates may need to go straight to frontier.

Multi-agent stacks: routing decisions multiply across every agent in the stack, making agent-level model assignment even more important. The orchestrator-worker pattern: one frontier model as orchestrator breaks tasks into subtasks and assigns them to worker agents; worker agents use cheap models for execution; the orchestrator only re-enters when a worker hits a problem it can't handle — concentrating frontier spend at the coordination layer where it has the highest leverage. Specialist agents with model tiers: assign tiers based on what each specialist does, not just what it is (a "researcher" agent might use cheap models for structured search result parsing and frontier models only for synthesizing findings). Agent-to-agent communication (structured data, status updates, task results) almost never needs a frontier model — well-formatted, predictable messages; cheap models handle them fine, or skip LLM calls entirely and pass data directly.

Common mistakes: routing by model reputation, not task fit (Claude Opus isn't better than Haiku for extracting a phone number from a form); routing all complex-sounding tasks to frontier ("analyze the sentiment of these 100 reviews" sounds like analysis but is actually well-defined classification); no fallback logic (static routing without fallback is fragile — a cheap model that usually gets it right will occasionally fail); ignoring latency requirements (frontier models 5–15 seconds vs cheap models under 500ms for short completions — route toward cheap for latency-sensitive contexts like live chat, voice interfaces, interactive tools); not logging routing decisions (without data on which model handled which tasks and how often fallback triggered, optimization is guesswork).

## Key points

- Frontier models are built for novelty, ambiguity, and complex reasoning; cheap models excel at well-defined, structured, repeated tasks.
- The cost difference between frontier and cheap models is typically 30–200x per token.
- At scale: 100,000 calls/day = ~$500/day all-frontier, ~$15/day all-cheap, $40–60/day with smart hybrid routing (~$165,000/year difference).
- "Frontier for planning, cheap for execution" is the most reliable agent-stack routing strategy.
- Routing strategies: static task-type, prompt complexity scoring, cascading/fallback, ensemble verification, speculative execution.
- ~60–80% of agent workloads can be handled well by cheap models; the remaining 20–40% benefits meaningfully from frontier.
- Build a routing layer: task taxonomy, router, fallback logic, logging, iterate on data.

## Technical data / figures

- Frontier pricing: ~$5–75 per M tokens; latency 5–30s.
- Cheap pricing: often under $0.50 per M; self-hosted ≈ compute cost; latency often <500ms.
- Cost gaps: Claude 3.5 Sonnet $3/M vs Haiku 3.5 $0.08/M = 37x; GPT-4o $5/M vs GPT-4o mini $0.15/M = 33x.
- 100,000 calls/day at 100 tokens/call: all-frontier ~$500/day; all-cheap ~$15/day; hybrid ~$40–60/day; ~$165,000/year difference.
- Frontier signals: ambiguity, multi-step reasoning, no clear template, outputs hard to validate automatically.
- Cheap-model territory: extraction, classification, formatting, translation, simple summarization, well-scoped RAG.

## Why this source matters for the RAG

Delivers the frontier-vs-cheap mental model, quantified cost gaps (30–200x), a 100K-calls/day cost example, and a complete routing-layer build guide — core for RAG on model routing and agent-stack cost architecture. The routing strategies (cascade, speculative, ensemble) and orchestrator-worker pattern are directly reusable.

