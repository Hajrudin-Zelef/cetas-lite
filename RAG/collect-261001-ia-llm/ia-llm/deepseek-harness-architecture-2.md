---
id: collect-261001-ia-llm/ia-llm/deepseek-harness-architecture-2
title: "DeepSeek Harness Architecture"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "consumer", "memory", "sandbox", "settlement"]
source: docs/RAG/collect-261001-ia-llm/deepseek-harness-architecture.md
source_anchor: ""
source_lines: [83, 136]
sha256: e852801037237789cb14fd81dca8799573030ae0c2e56ea890433d6315387c44
---

# DeepSeek Harness Architecture

```
turn/start
  claim next-step input plus one queued message
  assemble prompt sections + tool schemas; project runtime context
  -> agent/pre-step                   reject | enter(messages, startsRequestSeries?)
     reject, or a first enter rewritten empty -> close the turn with no step
     step/start
     agent/request -> prepareCall (cancellation commits neither system nor users)
     reconcile system/message using the prepared call capability
     append entered messages as user/message; log request/header and request/context as needed
     derive and freeze model history from the log
     stream the bound prepared call -> llm/stream -> agent/assistant-stream start
       agent/assistant-stream chunk*
       assistant/message | assistant/attempt -> agent/assistant-stream end
     tool/call* -> tools/pre-execute -> tools/execute -> tools/post-execute -> tool/result*
     step/end
     tools owe another request, or next-step input arrived -> claim -> next step
  -> agent/turn-stopping
turn/end
```
`turn/*`, `step/*`, `system/message`, `user/message`, `assistant/message`, `assistant/attempt`, and `tool/*` are durable session events; the rest are live extension points across three domains. `agent/assistant-stream` publishes process-local start, transient chunk, and end frames. The loop commits the complete compact stream as one message or log-only attempt before a committed end frame, and the Web Session-follow adapter is the live event's only remote consumer. `agent/pre-step`, `agent/request`, `llm/stream`, and the three `tools/*` events are waterfalls, whose listeners must call `next()` to delegate; `agent/turn-stopping` is serial and has no `next()`.

One inbox feeds the driver; injected context waits for a waking message. AgentLoop’s durable `inbox` projection exposes pending input without live Agents.

`agent/pre-step` decides the accepted input. Listeners may rewrite or reject claimed messages; a rejected or empty first claim closes a durable turn without a step. An enter decision may set `startsRequestSeries`: the loop logs a fresh `request/header` (reason `series`, or `change` with `startsSeries: true` when the envelope also changed). Wrapping listeners preserve that declaration with `{ ...decision, messages }`. After assembly and `step/start`, `agent/request` and `prepareCall()` resolve the actual route before the system prompt and accepted users are committed; cancellation during either async phase commits neither. Prepared call capabilities govern prompt admission, not preceding `request/context`. Every attempt synchronously reconciles the same rendered assembly, appends users only on the first attempt, logs header/context as needed, and derives and freezes the request before streaming the bound call. Retries do not repeat assembly or `agent/pre-step`. Surface replacements and image-offload decisions after attachment start a new request series, including during the first resumed pre-step; unchanged resume continues the series. The first admitted step reserves the system head before user messages even for an empty prompt (no wire message). The prompt travels only as `system/message` history: an empty rendering clears all active system nodes, leaving no old prompt model-visible; capable routes append non-empty updates after cached history, including supported tool updates; incapable routes and new request series consolidate non-empty prompt text at the first system node, with logged empty replacements for non-empty later system nodes (decision; decision rule).

The loop sends immutable requests with live cancellation and reuses freeze evidence for frozen identities; agent-loop owns construction and cancellation causes.

Failed steps record missing tool results.

Details: the sequence diagram, the tool pipeline, and cancellation and error recovery.

## Session log 

The session log is the source of the context the model sees. `deriveMessages()` projects model history from it. Each `assistant/message` embeds the exact compact timed stream that produced its assembled content; `assistant/attempt` retains settled failed, retried, cancelled, and stream-error attempts without adding model history. Fork, resume, transcripts, telemetry, and persistence all derive from these durable settlements, while live UI incrementality comes from `agent/assistant-stream`; a hard process loss before settlement leaves no durable attempt stream (decision).

Session consumers know only the current logical format. Header-only `stat` and `list` rescan each Session directory, select its numerically highest canonical generation, and translate a supported historical header without loading events or publishing a successor. A stored-session `open` selects that same generation, refuses a future version, or decodes and composes the static adjacent migration chain once before returning validated current logical events. A read open uses that in-memory result without publishing a successor; a write open first encodes, verifies, and exclusively publishes the final version-named successor beside the unchanged source. Ordinary repair of an unsealed interrupted tail remains a handle consumer responsibility; migration inserts a missing interrupted `turn/end` only for the bounded released restart already sealed by a later `turn/start`. JSONL v0 uses `session.jsonl[.zstd]`, v1 and later use lowercase `session.vN.jsonl[.zstd]`, and committed generation paths are never renamed, replaced, or deleted. The JSONL provider owns physical framing, compression, generation selection, and exclusive publication, while each adjacent migration package owns exactly one `vN -> vN+1` step (decision).

**Model-visible means logged.** A runtime invariant checks model requests are reconstructable from the log. New model-visible inputs require session events. Plugins changing existing message content register pure message projections; detached readers supply the same definitions explicitly. Tool changes are capability-independent; Session tool history supplies provider declarations.

**Projection seam.** `dsh-session-projection` owns `ctx.sessionProjections`: registered units fold committed events incrementally, host consumers read one typed state with `stateOf()`, and carriers batch cropped client views with `snapshot()`. A host reader either requires this service during activation or fails explicitly when the registry or required key is absent. Contributors may retain `ctx.inject(['sessionProjections'], ...)` registration without silently defaulting a missing host value. The agent loop registers shared `turnBoundary` state for its readers (decision).

## Capability seams 

A **seam** is a swappable capability with three roles: a **Service Definition** declaring the interface, a **Service Provider** implementing it, and a **Consumer** using it, commonly a model-facing tool. A package may combine roles, but one role alone is not a seam; adding a capability means designing all three (capability graph).

Seams are why one provider swap changes the whole product. Filesystem and subprocess providers share one execution world, so pointing them at a remote sandbox moves Bash, PTY, and LSP with them, with no provider forks. Subagent providers vary just as widely behind one interface, from a fresh child agent to a delegated turn in another product.

Experimental Agent Teams is a published opt-in coordination seam on `ctx.agentTeams`, with a durable roster, task board, and mailbox layered over continuable subagents.

## Where new behavior goes 

New behavior attaches to a documented extension point. Changing the loop itself updates this map.

