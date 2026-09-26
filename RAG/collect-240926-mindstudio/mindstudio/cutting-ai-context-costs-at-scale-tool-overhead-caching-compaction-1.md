---
id: collect-240926-mindstudio/mindstudio/cutting-ai-context-costs-at-scale-tool-overhead-caching-compaction-1
title: "cutting-ai-context-costs-at-scale-tool-overhead-caching-compaction"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["cost", "agent", "agentic", "agents", "context window", "mcp", "memory", "pruning"]
source: docs/RAG/clean_en/mindstudio/cutting-ai-context-costs-at-scale-tool-overhead-caching-compaction.md
source_anchor: ""
source_lines: [1, 62]
sha256: 767dbb435f7410aa37de592fe6bd561659f5b9ba3759c61637c9852f32122489
---

# cutting-ai-context-costs-at-scale-tool-overhead-caching-compaction

<!-- source: https://www.mindstudio.ai/blog/advanced-context-engineering-token-savings -->

## Why do API costs explode even when your prompts stay short?

Because every request to a large language model resends the entire conversation history, not just your new message. The model has no persistent memory between calls, so each turn packages up everything that came before, tool definitions, prior responses, system instructions, and sends it again. That resent material is called “reused input,” and at scale it dwarfs anything a person actually typed. One heavy Codex user tracking a single working day recorded 3.77 billion tokens moving through the workspace, with 3.59 billion of them (about 96%) being reused input across 143 threads. Nobody typed that. The mechanics of how chat-based LLMs simulate memory did.

This matters more as usage scales up, not less. Basic habits like trimming prompts and starting fresh threads help, but they run into diminishing returns once you’re running multi-tool agents, long debugging sessions, or automated pipelines. At that point the real costs live in three places: what tools you load, how the context window gets managed over long tasks, and whether anything sits between your app and the model to catch waste before it’s billed.

## TL;DR

- **Reused input dominates token bills** because every API call resends the full conversation history, so later messages in a thread cost far more than the first one, even when nothing new is added.
- **Tool definitions are billed as input before the model does anything** , and Anthropic has published figures showing a typical multi-server setup (GitHub, Slack, Sentry, Grafana) can burn roughly 55,000 tokens on tool descriptions alone before a task starts.
- **Compaction and context editing let long-running tasks shed old material** without a full restart: OpenAI’s Codex compaction carries forward summarized state, while Anthropic’s context editing strips out stale tool results and thinking blocks between requests.
- **Neither compaction nor context editing is free of tradeoffs** , since both depend on approximated versions of earlier turns rather than the original text, so debugging accuracy can drift on very long sessions.
- **Loading fewer tools per task is a direct, controllable lever** that doesn’t require waiting on model providers to fix the underlying inefficiency.
- **Middleware that intercepts requests before they reach the model** is the most powerful lever available, because it can filter, cache, or reroute calls at the infrastructure level rather than relying on prompt discipline alone.

## Remy doesn't build the plumbing. It inherits it.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

## How does tool-definition overhead actually work?

When you connect an AI agent to external tools (via something like MCP servers, function calling, or plugin registries), each tool comes with a schema: a name, a description of what it does, when to use it, and what arguments it accepts. That schema is not optional metadata sitting off to the side. It gets serialized into the model’s input context on every single call, whether or not the model ends up using that tool.

Anthropic’s own published analysis found that a fairly ordinary setup, a handful of connected servers like GitHub, Slack, Sentry, and Grafana, can consume around 55,000 tokens in tool definitions before the model has processed a single instruction from the user. That’s overhead paid on every turn of every conversation using that configuration, regardless of task complexity.

The practical fix is loading only the tools a given job actually needs, rather than connecting every available integration by default and letting the model sort it out. Current models aren’t yet reliably good at ignoring irrelevant tools cheaply. That may improve as models get better at pruning their own context, but as of now it’s a cost that sits squarely on whoever configures the agent.

## What is compaction and when should you use it?

Compaction is a technique for long-running agent sessions where the system periodically summarizes or condenses earlier parts of the conversation instead of carrying the full raw history forward. OpenAI’s Codex supports this for extended work sessions: rather than resending every prior exchange verbatim, it carries forward a compressed state that preserves the necessary context while using fewer tokens on subsequent calls.

This matters most in situations where restarting a clean thread isn’t practical, for example when you’re deep into debugging a system and the model needs a decision from you mid-task. Killing the thread and starting over would lose the working state you’ve built up. Compaction lets the session continue without paying full price for every previous message on every future call.

The tradeoff is that compaction relies on a summarized, approximated version of earlier turns rather than the original text. For most conversational work that’s a fine trade. For tasks requiring exact recall of something said many turns ago, like a specific error message or an exact code snippet, an approximation can lose fidelity that a full context window would have preserved.

## What is context editing and how is it different?

Anthropic’s context editing works on a related but distinct problem: clearing out old tool results and intermediate thinking blocks between requests, rather than summarizing the whole conversation. Long agentic tasks often involve calling tools repeatedly and getting back verbose results (file contents, search hits, command output) that were useful in the moment but become dead weight once the task moves past them.

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

Context editing strips that stale material out before the next request goes to the model, keeping the working context lighter without discarding the actual conversational thread the way starting a new task would. Like compaction, it depends on the system correctly deciding what’s safe to drop, which means you’re trusting an automated process to distinguish “no longer needed” from “might matter later.”

For advanced users, the practical implication is the same in both cases: know roughly where your context window sits at any point in a long session, and understand that once compaction or context editing kicks in, you’re no longer working from a verbatim record. That’s usually fine. Occasionally it isn’t, and it’s worth knowing which mode you’re in when accuracy on a specific detail really matters.

## Is prompt caching a solution to this problem?

Prompt caching addresses a related but separate issue: cost, not context size. When a portion of your input (a system prompt, a tool schema, a large reference document) stays identical across multiple calls, prompt caching lets the provider charge less for reprocessing that unchanged portion instead of billing it at full input-token rates every time. It doesn’t shrink your context window or reduce how much the model has to attend to. It reduces what you pay for the parts that haven’t changed.

This makes caching most valuable when you have stable, reusable context, like a consistent set of tool definitions or a long reference document that gets queried repeatedly, sitting alongside a smaller amount of turn-specific content. It’s less useful for conversations where the bulk of the context changes every turn, since there’s nothing stable to cache.

Caching and compaction/context editing solve different halves of the same problem: caching makes reused tokens cheaper, while compaction and context editing make there be fewer reused tokens in the first place. Using both together, where applicable, compounds the savings.

## Why does middleware interception matter for scale?

