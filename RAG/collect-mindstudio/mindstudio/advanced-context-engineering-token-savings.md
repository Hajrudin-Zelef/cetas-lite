---
id: collect-mindstudio/mindstudio/advanced-context-engineering-token-savings
title: "Cutting AI Context Costs at Scale: Tool Overhead, Caching, Compaction"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["cost", "agent", "agentic", "agents", "context window", "mcp", "memory"]
source: docs/RAG/Collect RAG/02_mindstudio/advanced-context-engineering-token-savings.md
source_anchor: ""
source_lines: [1, 49]
sha256: 8ab239f77c5dd3a5e8139928e50506c79748409be4dd30649b6195c82267a483
---

# Cutting AI Context Costs at Scale: Tool Overhead, Caching, Compaction

## Metadata

- **Source** : https://www.mindstudio.ai/blog/advanced-context-engineering-token-savings
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This technical article explains why LLM API costs explode even when prompts stay short, and details the mechanisms that drive cost: tool-definition overhead, compaction, context editing, prompt caching, and middleware interception.

Root cause: every request to an LLM resends the entire conversation history because models have no persistent memory between calls. Each turn packages everything before it (prior responses, tool definitions, system instructions), which is "reused input." At scale this dwarfs what a person actually typed. One heavy Codex user tracked a single working day with 3.77 billion tokens moving through the workspace, 3.59 billion of which (~96%) were reused input across 143 threads. Nobody typed that.

Tool-definition overhead: each tool connected to an agent (via MCP servers, function calling, plugin registries) comes with a schema (name, description, usage, arguments) that is serialized into the model's input context on every call whether or not the tool is used. Anthropic's published analysis found a fairly ordinary setup — GitHub, Slack, Sentry, Grafana — can consume ~55,000 tokens in tool definitions before the model processes a single instruction. The fix is loading only the tools a given job needs; current models aren't yet reliably good at ignoring irrelevant tools cheaply.

Compaction: for long-running agent sessions, the system periodically summarizes/condenses earlier conversation instead of carrying full raw history. OpenAI's Codex supports this for extended work sessions, carrying forward a compressed state. This matters when restarting a clean thread isn't practical (e.g., deep debugging where the model needs a decision mid-task). Tradeoff: relies on summarized/approximated versions of earlier turns, losing exact fidelity (specific error messages, exact code snippets).

Context editing: Anthropic's variant removes specific stale elements (old tool results, intermediate thinking blocks) between requests rather than summarizing the whole conversation. Long agentic tasks call tools repeatedly and get verbose results (file contents, search hits, command output) that become dead weight. Context editing strips that material out before the next request. Tradeoff: trusting an automated process to distinguish "no longer needed" from "might matter later."

Prompt caching: addresses cost, not context size. When a portion of input (system prompt, tool schema, large reference document) stays identical across calls, the provider charges less for reprocessing the unchanged portion. It doesn't shrink the context window; it reduces what you pay for parts that haven't changed. Most valuable with stable reusable context (consistent tool definitions, long reference documents); less useful when the bulk of context changes every turn. Caching makes reused tokens cheaper; compaction/context editing make there be fewer reused tokens — using both compounds savings.

Middleware interception: a layer between app and provider that intercepts requests and applies rules/filters/routing logic the model has no part in. This moves optimization from "hope the model/user remembers to be efficient" into infrastructure that enforces efficiency. It can catch redundant tool loads, route to cheaper models, block retries that resend failed multi-thousand-token contexts. For teams running many agents or automated pipelines, this is where the largest, most consistent savings come from. For solo users, habits and provider features (compaction, caching) cover most of the ground.

## Key points

- Reused input (full conversation history resent each call) dominates token bills; one Codex day: 3.77B tokens moved, 96% reused input across 143 threads.
- Tool definitions are billed as input on every call; Anthropic: ~55,000 tokens for a GitHub/Slack/Sentry/Grafana setup before a task starts.
- Load only needed tools — models aren't yet good at cheaply ignoring irrelevant tools.
- Compaction (OpenAI Codex): summarizes and carries forward compressed state for long sessions; loses exact fidelity.
- Context editing (Anthropic): strips stale tool results and thinking blocks between requests.
- Prompt caching cuts cost of reprocessing unchanged content; it's a billing optimization, not a context-size optimization.
- Middleware interception is the most powerful lever at scale: filter, cache, or reroute calls at the infrastructure level.

## Technical data / figures

- Codex workspace example: 3.77B tokens/day, 3.59B (≈96%) reused input, 143 threads.
- Anthropic tool-definition figure: ~55,000 tokens for a multi-server setup (GitHub, Slack, Sentry, Grafana).
- Compaction supported by: OpenAI Codex. Context editing supported by: Anthropic.
- Prompt caching: reduces cost of unchanged reprocessing; content still counts toward context window.
- Middleware: can catch redundant tool loads, route to cheaper models, block repeated retries of failed contexts.

## Why this source matters for the RAG

Provides the core mechanics and quantified figures behind AI token-cost blowups (reused input, tool-definition overhead) and the mitigation techniques (compaction, context editing, caching, middleware) — foundational for RAG entries on context engineering and token-savings strategy. The 55K-token tool overhead and 96% reused-input figures are directly citable.

