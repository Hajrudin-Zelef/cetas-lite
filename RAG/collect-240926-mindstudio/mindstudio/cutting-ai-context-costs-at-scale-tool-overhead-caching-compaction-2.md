---
id: collect-240926-mindstudio/mindstudio/cutting-ai-context-costs-at-scale-tool-overhead-caching-compaction-2
title: "cutting-ai-context-costs-at-scale-tool-overhead-caching-compaction"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["cost", "agent", "agentic", "agents", "context window", "memory"]
source: docs/RAG/clean_en/mindstudio/cutting-ai-context-costs-at-scale-tool-overhead-caching-compaction.md
source_anchor: ""
source_lines: [63, 89]
sha256: abbf46e3b265d2958a475ebfc3f0067db6d3e925328beeac7b7552057a3700f4
---

# cutting-ai-context-costs-at-scale-tool-overhead-caching-compaction

Everything above (tool loading, compaction, context editing, caching) still operates within the assumption that requests go straight from your application to the model provider. A middleware layer sits between the two, intercepting requests before they’re sent and applying rules, filters, or routing logic that the model itself has no part in.

This matters because it moves optimization out of “hope the model or the user remembers to be efficient” and into infrastructure that enforces efficiency regardless of what any individual prompt does. For teams running many agents or automated pipelines rather than a single chat interface, this is usually where the largest, most consistent savings come from, because it doesn’t rely on habit or discipline at the point of use. It can catch redundant tool loads, route requests to cheaper models when appropriate, or block retries that would otherwise resend a failed multi-thousand-token context for a second and third time.

For solo users chatting with an assistant, habits and built-in provider features like compaction and caching cover most of the ground. For teams or builders running agentic systems at real volume, the interception layer is what keeps costs proportional to actual work rather than to how many turns a conversation happened to take.

## Frequently Asked Questions

### What is reused input in the context of LLM API costs?

Reused input is the portion of a request that consists of previously sent conversation history, tool definitions, or system instructions being resent because the model has no memory between calls. It typically makes up the vast majority of tokens in any extended conversation or agent session.

### Does prompt caching reduce my context window usage?

No. Prompt caching reduces the cost of reprocessing unchanged content, but the full content is still sent and still counts toward the model’s context window limit. It’s a billing optimization, not a context-size optimization.

### What’s the difference between compaction and context editing?

Compaction, as supported by OpenAI’s Codex, summarizes and carries forward a condensed state across a long session. Context editing, as supported by Anthropic, removes specific stale elements like old tool results and thinking blocks between requests. Both reduce token load on long tasks but through different mechanisms.

### How many tokens do tool definitions typically cost?

Anthropic has published figures showing a typical setup with several connected tool servers, such as GitHub, Slack, Sentry, and Grafana, can consume around 55,000 tokens in tool definitions alone before any task-specific processing begins.

### Is it worth building custom middleware just for token savings?

It depends on scale. For occasional use, habit changes and provider-native features like caching and compaction usually suffice. For teams running many concurrent agents or automated workflows, middleware that intercepts and filters requests before they reach the model tends to produce larger and more consistent savings than relying on per-conversation discipline alone.
