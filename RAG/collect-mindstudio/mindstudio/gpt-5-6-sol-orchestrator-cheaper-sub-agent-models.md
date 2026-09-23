---
id: collect-mindstudio/mindstudio/gpt-5-6-sol-orchestrator-cheaper-sub-agent-models
title: "How to Use GPT-5.6 Sol as an Orchestrator with Cheaper Sub-Agent Models"
domain: mindstudio
role: reference
task: article
actors: ["OpenAI"]
dates: ["2026-09-23"]
keywords: ["agent", "gpt-5.6", "sol", "agents", "cost", "latency", "luna", "reasoning", "terra", "throughput", "training"]
source: docs/RAG/Collect RAG/02_mindstudio/gpt-5-6-sol-orchestrator-cheaper-sub-agent-models.md
source_anchor: ""
source_lines: [1, 56]
sha256: 5f4dd2d41f52f519f723ea42125747274f6c63333c0b7f2ef0be8a78e336c763
---

# How to Use GPT-5.6 Sol as an Orchestrator with Cheaper Sub-Agent Models

## Metadata

- **Source** : https://www.mindstudio.ai/blog/gpt-5-6-sol-orchestrator-cheaper-sub-agent-models
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This guide explains why using the most powerful model for every task is a mistake, and how to use GPT-5.6 Sol as an orchestrator with cheaper sub-agents (Luna and Terra) to cut costs while keeping output quality high.

Core argument: most of what a complex AI workflow actually does doesn't require that level of intelligence. Summarizing a document, formatting a response, extracting structured data from a form — these tasks don't need a model that excels at deep reasoning and multi-step planning. Routing every sub-task through GPT-5.6 Sol when cheaper models handle them just as well means paying a premium for capability you're not using. The smarter approach: let GPT-5.6 Sol do what it's genuinely good at — orchestrating, planning, and making decisions — while delegating execution to lower-cost models like Luna or Terra.

Orchestrator/sub-agent architecture: separating the "thinking" from the "doing." The orchestrator receives a high-level goal, breaks it into discrete tasks, decides which agent or model should handle each, and synthesizes results into a coherent final output. Sub-agents execute specific, bounded tasks and return results without needing to understand the broader context. Analogy: a project manager coordinating a team of specialists.

Why the pattern works: LLMs are priced based on capability, which correlates with model size and training cost. GPT-5.6 Sol's strength lies in complex reasoning, nuanced judgment, and long-context planning. Luna and Terra are lighter, faster models that handle high-throughput tasks well at a fraction of the cost. When task complexity is aligned with model capability: Sol handles orchestration, intent parsing, quality checks, and edge-case decisions; Luna handles mid-complexity tasks (summarization, rewriting, classification, structured extraction); Terra handles simpler repetitive tasks (formatting, templating, basic lookups). Cost argument: orchestrating a workflow with GPT-5.6 Sol at every step versus only for orchestration and judgment can reduce per-run costs by 60–80% depending on workflow structure. At scale — thousands of runs per week — that difference is substantial.

GPT-5.6 Sol's orchestrator strengths: task decomposition (reliably breaks a complex user request into logical subtasks with the right level of granularity); routing decisions (assesses which model tier is appropriate for each subtask based on complexity, expected output format, and sensitivity — routing logic is where much cost efficiency is won or lost); synthesis and quality control (evaluates sub-agent results for consistency, catches errors, decides whether to pass output downstream or loop back for revision); handling ambiguity (handles underspecified requests better than lighter models, filling in reasonable assumptions or asking targeted clarifying questions before kicking off sub-tasks). What it shouldn't be doing: bulk content generation after a plan is established, simple data transformation or extraction, template-based formatting, repetitive classification at high volume — these are exactly the tasks Luna and Terra are suited for.

Task-to-model mapping: GPT-5.6 Sol (orchestration layer) = interpreting user intent from ambiguous/complex prompts, generating the initial task plan and deciding execution sequence, selecting which sub-agent handles each step, reviewing sub-agent outputs for accuracy and coherence, handling exceptions/conflicts/unexpected results, final output assembly and high-stakes rewriting. Luna (mid-tier execution) = summarizing long documents/conversations, rewriting/reformatting content based on established guidelines, multi-label classification with moderate complexity, extracting structured data from semi-structured inputs, drafting email/message responses within a defined template, intermediate reasoning steps that don't require full model depth. Terra (high-volume simple execution) = single-label classification (sentiment, category), filling in templates with provided data, simple data lookups and string transformations, boolean decisions based on clear criteria, input validation and formatting normalization.

Structuring the workflow (5 steps): (1) Define the workflow's entry point and goal — pass input to GPT-5.6 Sol with a clear orchestration prompt describing the overall goal, listing available sub-agents and what they're good for, specifying the output format for the task plan (JSON works well), and defining constraints. A minimal orchestrator system prompt example: "You are an orchestrator. Given the user's request, produce a task plan as a JSON array. Each task should include: task_id, description, model_tier (sol/luna/terra), input_data, and expected_output_format. Available agents: Luna — summarization, classification, drafting, extraction; Terra — templating, formatting, simple lookups. Handle ambiguous inputs by stating your assumptions. Route to Sol only tasks that require judgment or synthesis." (2) Parse the task plan and route to sub-agents — branching logic if model_tier == "luna" or "terra". (3) Handle task dependencies — include a depends_on field listing task_ids that must complete first so the workflow runner sequences them correctly. (4) Collect and return results to the orchestrator — the orchestrator checks format, identifies incomplete/inconsistent results, flags tasks for retry, and synthesizes all outputs into the final deliverable. (5) Define exit conditions — all tasks complete without errors, final output passes a defined quality check, max retry count reached, or explicit human approval — to prevent infinite loops or partial outputs that look complete.

Common mistakes: over-routing to Sol (cost savings disappear; build explicit routing rules into the orchestrator prompt and test routing decisions with a sample set of inputs before deploying at scale; heuristic: clearly defined input format, predictable output format, limited edge cases → Luna or Terra); under-specifying sub-agent prompts (must be self-contained — every instruction needed should be in the prompt passed to them, not assumed from the broader context); skipping the review pass (tempting to save tokens, but consistency errors and hallucinations get caught in review; cutting it creates more rework than it saves); ignoring latency (running tasks sequentially through multiple model calls adds latency; map which tasks can run in parallel — parallelism is one of the biggest levers for keeping multi-agent workflows fast).

Measuring and optimizing: cost per run (track how many tokens each model tier consumes per execution; if Sol handles a disproportionate share, revisit routing logic); task retry rate (high retry rates on a sub-agent usually mean the prompt is underspecified or the task is more complex than anticipated — consider bumping those tasks up to Luna); output quality scores (define a quality rubric and score a sample of runs regularly; if quality drops when shifting tasks to cheaper models, you've over-optimized); latency by workflow step (if the review pass is consistently slow, batch multiple sub-task results into a single review call).

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

