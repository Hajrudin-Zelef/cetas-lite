---
id: collect-mindstudio/mindstudio/gpt-5-6-sol-orchestrator-cheaper-sub-agent-models-2
title: "How to Use GPT-5.6 Sol as an Orchestrator with Cheaper Sub-Agent Models"
domain: mindstudio
role: reference
task: article
actors: ["OpenAI"]
dates: []
keywords: ["agent", "gpt-5.6", "sol", "agents", "cost", "latency", "luna", "terra"]
source: docs/RAG/Collect RAG/02_mindstudio/gpt-5-6-sol-orchestrator-cheaper-sub-agent-models.md
source_anchor: ""
source_lines: [32, 56]
sha256: ebe18b1165155363682e4d499223adfe0b69e352c487391d4922039e239bdc96
---

# How to Use GPT-5.6 Sol as an Orchestrator with Cheaper Sub-Agent Models

Practical notes: sub-agents can call other sub-agents (hierarchical or nested multi-agent architecture — Luna could act as a mid-tier orchestrator for its own Terra sub-tasks), but start flat (one orchestrator, multiple sub-agents) until there's a clear reason to add hierarchy. There's no fixed answer for how many sub-agents an orchestrator should manage, but workflows with more than 8–10 concurrent sub-tasks often benefit from chunking work into phases. Build error handling into workflow logic — flag for retry, use a fallback approach, or escalate to human review — rather than assuming it will always work cleanly.

## Key points

- Using the most powerful model for everything wastes spend on capability you don't use.
- Orchestrator/sub-agent pattern: GPT-5.6 Sol plans, routes, and synthesizes; Luna handles mid-complexity tasks; Terra handles simple high-volume tasks.
- Cost reduction of 60–80% per run is achievable by limiting Sol to orchestration and judgment.
- Sol's orchestrator strengths: task decomposition, routing decisions, synthesis/quality control, ambiguity handling.
- Sub-agent prompts must be self-contained; the orchestrator review pass catches errors and hallucinations.
- Heuristic: tasks with defined input/output and limited edge cases belong on Luna or Terra.
- ~8–10 concurrent sub-tasks is a practical ceiling; chunk into phases and define exit conditions.

## Technical data / figures

- Cost impact: orchestrating every step vs only orchestration/judgment → 60–80% per-run reduction.
- Tier mapping: Sol (orchestration, judgment, synthesis); Luna (summarization, classification, drafting, extraction); Terra (templating, formatting, simple lookups, boolean decisions, validation).
- Orchestrator prompt output: JSON array with task_id, description, model_tier (sol/luna/terra), input_data, expected_output_format; depends_on field for dependencies.
- Workflow steps: entry point → route to sub-agents → dependencies → review pass → exit conditions.
- Practical ceiling: 8–10 concurrent sub-tasks; parallelism for latency; retry/fallback/human-escalation error handling.
- Monitoring: cost per run, retry rate, quality scores, latency by step.

## Why this source matters for the RAG

Provides a concrete, reusable architecture for orchestrator/sub-agent cost optimization, including a working system-prompt template, task-tier mapping, and step-by-step build instructions — directly applicable to RAG content on multi-agent orchestration and token-cost reduction. The 60–80% savings figure and the Sol/Luna/Terra tiering are highly citable.

