---
id: collect-240926-mindstudio/mindstudio/how-to-use-gpt-5-6-sol-as-an-orchestrator-with-cheaper-sub-agent-models-2
title: "how-to-use-gpt-5-6-sol-as-an-orchestrator-with-cheaper-sub-agent-models"
domain: mindstudio
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["agent", "sol", "agents", "cost", "gpt-5.6", "latency", "luna", "terra"]
source: docs/RAG/clean_en/mindstudio/how-to-use-gpt-5-6-sol-as-an-orchestrator-with-cheaper-sub-agent-models.md
source_anchor: ""
source_lines: [139, 252]
sha256: 1c2a52267a1bc94bda96265ce7166f6d9dc3a366acbabff640831794addb7b96
---

# how-to-use-gpt-5-6-sol-as-an-orchestrator-with-cheaper-sub-agent-models

In a no-code environment like MindStudio, this looks like a branching condition — if `model_tier == "luna"`, run the task through Luna; if `model_tier == "terra"`, route to Terra. Each branch passes the `input_data` and a task-specific prompt to the assigned model.

In a code-first setup, you’d parse the JSON and loop through tasks, calling the appropriate API endpoint for each tier.

### Step 3: Handle Task Dependencies

Some tasks can run in parallel. Others must wait for the output of a previous task. Your orchestration plan should flag these dependencies so your workflow runner knows to sequence them correctly.

Ask GPT-5.6 Sol to include a `depends_on` field in each task object, listing any `task_id` values that must complete first. This keeps your workflow from sending downstream tasks before their inputs are ready.

### Step 4: Collect and Return Results to the Orchestrator

Once sub-agents finish their tasks, pass all outputs back to GPT-5.6 Sol for review. The orchestrator should:

- Check that outputs match the expected format
- Identify any tasks that produced incomplete or inconsistent results
- Flag tasks for retry if needed
- Synthesize all outputs into the final deliverable

### Built like a system. Not vibe-coded.

Remy manages the project — every layer architected, not stitched together at the last second.

This review step is where having a capable orchestrator really pays off. Luna and Terra don’t have the context to judge whether their individual outputs fit together correctly — GPT-5.6 Sol does.

### Step 5: Define Exit Conditions

Decide upfront when the workflow is “done.” Options include:

- All tasks complete without errors
- Final output passes a defined quality check
- Maximum retry count reached
- Explicit human approval (for sensitive workflows)

Without clear exit conditions, orchestrated workflows can loop indefinitely or produce partial outputs that look complete.

## Common Mistakes and How to Avoid Them

### Over-Routing to Sol

If your orchestrator is handling tasks that Luna or Terra could do, your cost savings disappear. Build explicit routing rules into your orchestrator prompt and test routing decisions with a sample set of inputs before deploying at scale.

A useful heuristic: if a task has a clearly defined input format, a predictable output format, and limited edge cases, it belongs on Luna or Terra.

### Under-Specifying Sub-Agent Prompts

Sub-agents operate without the broader context the orchestrator holds. Their prompts need to be self-contained — every instruction they need to complete the task should be in the prompt passed to them, not assumed from the overall workflow.

If Terra is formatting a piece of text, the prompt should include the exact format specification, not a reference to “the formatting guidelines discussed earlier.”

### Skipping the Review Pass

It’s tempting to skip sending outputs back to GPT-5.6 Sol for review to save on tokens. In low-stakes workflows, that’s fine. But in anything customer-facing or consequential, the review pass is where consistency errors and hallucinations get caught. Cutting it tends to create more rework than it saves in token costs.

### Ignoring Latency

Running tasks sequentially through multiple model calls adds latency. For user-facing applications, map out which tasks can run in parallel and structure your workflow accordingly. Parallelism is one of the biggest levers for keeping multi-agent workflows fast.

## How MindStudio Makes This Architecture Easy to Build

Setting up an orchestrator/sub-agent system from scratch — managing API calls, routing logic, retries, and result aggregation — is non-trivial engineering work. MindStudio’s visual workflow builder handles most of the infrastructure so you can focus on the logic.

Within MindStudio, you can:

- Drop GPT-5.6 Sol into an orchestrator node with a custom system prompt
- Add Luna and Terra as separate agent nodes, each with their own prompts and configurations
- Use conditional branching to route tasks based on the orchestrator’s output
- Connect results back to the orchestrator for the review pass
- Run sub-agents in parallel where task dependencies allow

MindStudio gives you access to 200+ AI models in a single platform — no separate API keys, no account juggling. You can swap model tiers, test routing decisions, and iterate on prompts without touching infrastructure.

For teams running high-volume workflows, the cost difference between routing everything through a premium model versus using an orchestrator/sub-agent pattern can be significant. MindStudio’s built-in model access makes it straightforward to experiment with different model combinations until you find the right balance of cost and quality for your specific use case.

## Measuring and Optimizing the System

Once your orchestrated workflow is running, monitor these metrics:

**Cost per run.** Track how many tokens each model tier consumes per workflow execution. If Sol is handling a disproportionate share of tokens, revisit your routing logic.

**Task retry rate.** High retry rates on a specific sub-agent often mean the prompt is underspecified or the task is more complex than anticipated. Consider bumping those tasks up to Luna.

**Output quality scores.** Define a quality rubric for final outputs and score a sample of runs regularly. If quality drops when you shift tasks to cheaper models, you’ve likely over-optimized.

**Latency by workflow step.** Identify bottlenecks. If the orchestrator review pass is consistently slow, consider whether you can batch multiple sub-task results into a single review call rather than reviewing each individually.

## Frequently Asked Questions

### What’s the difference between an orchestrator and a sub-agent?

An orchestrator manages the overall task — it plans, routes, and synthesizes. A sub-agent executes a specific, bounded piece of work within that plan. The orchestrator needs to understand the full context; sub-agents only need to understand their individual task.

### How do I know which tasks to assign to cheaper models?

A useful test: if you gave the task to a competent but junior person who had only the information in the prompt (no broader context), could they complete it reliably? If yes, it’s probably appropriate for Luna or Terra. If the task requires judgment, context-sensitivity, or handling ambiguity, keep it with Sol.

### Does using sub-agents affect output quality?

When done correctly, the quality impact is minimal. The orchestrator’s synthesis pass catches errors and inconsistencies from sub-agents. The key is being honest about which tasks genuinely need high-capability models versus which ones just feel like they do.

### Can sub-agents call other sub-agents?

Yes, and this is sometimes called a hierarchical or nested multi-agent architecture. Luna could act as a mid-tier orchestrator for its own set of Terra sub-tasks. This adds complexity but can improve efficiency in very large workflows. Start flat (one orchestrator, multiple sub-agents) until you have a clear reason to add hierarchy.

### How many sub-agents should an orchestrator manage?

There’s no fixed answer, but practical limits exist. GPT-5.6 Sol can track a finite amount of context. Workflows with more than 8–10 concurrent sub-tasks often benefit from chunking the work into phases, with the orchestrator handling one phase at a time.

### What happens if a sub-agent fails or returns a bad output?

Your orchestrator should be prompted to handle this explicitly — either flagging the task for retry, using a fallback approach, or escalating to a human review. Build error handling into your workflow logic rather than assuming it will always work cleanly.

## Key Takeaways

