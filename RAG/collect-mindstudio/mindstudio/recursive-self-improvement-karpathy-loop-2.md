---
id: collect-mindstudio/mindstudio/recursive-self-improvement-karpathy-loop-2
title: "What Is Recursive Self-Improvement in AI? The Karpathy Loop Explained"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "Google", "OpenAI"]
dates: ["2026-05"]
keywords: ["recursive self-improvement", "agent", "claude", "compute", "gemini", "reasoning", "research", "safeguards", "throughput"]
source: docs/RAG/Collect RAG/02_mindstudio/recursive-self-improvement-karpathy-loop.md
source_anchor: ""
source_lines: [45, 55]
sha256: 9b18eefe0e923818b23ba4b7df57cd32923d18c9e2ae9dd8f3fda54db472ea52
---

# What Is Recursive Self-Improvement in AI? The Karpathy Loop Explained

- Loop stages: Propose → Implement → Execute → Evaluate → Commit or Discard → Repeat.
- Standard assistant vs agent loop: initiation (no vs yes), runs code (sometimes vs yes), reads output (no vs yes), commits (no vs yes), iterates (only when asked vs automatically).
- Throughput: 10–20 meaningful iterations/day (human + assistant) vs hundreds (autonomous cloud-compute loop).
- Recommended models: Claude 3.5 Sonnet, Claude 3 Opus, GPT-4o, Gemini 1.5 Pro.
- RSI vs RL: RL updates weights via reward signals; the loop uses an LLM as a fixed reasoning engine (system-level RL, no weight updates).
- Safety constraints: no weight modification; human-defined objectives; auditable commits; bounded scope.
- Safeguards: scope-based human approval, full logging, rollback triggers on metric degradation, tightly scoped permissions.

## Why this source matters for the RAG

Provides current (May 2026) coverage of the Karpathy Loop — a concrete, named implementation of recursive self-improvement, including the Auto Research workflow, model recommendations, safety constraints, and the distinction from reactive coding assistants and reinforcement learning. This is precise, up-to-date knowledge that complements the other RSI articles and lets the RAG answer accurately on autonomous agent loops and RSI concepts.
