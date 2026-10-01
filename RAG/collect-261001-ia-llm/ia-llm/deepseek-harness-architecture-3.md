---
id: collect-261001-ia-llm/ia-llm/deepseek-harness-architecture-3
title: "DeepSeek Harness Architecture"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/deepseek-harness-architecture.md
source_anchor: ""
source_lines: [137, 160]
sha256: c7cb1ec0f790a235ac6c944b41dca9df577e742c7e4ba07b39aaa822076ac0ca
---

# DeepSeek Harness Architecture

| Goal | Mechanism | 
|---|---|
| Add a model provider | register its adapter on `ctx.llm` | 
| Add a model-facing capability | register on `ctx.tools` ; its schema joins prompt assembly | 
| Give one session a different capability set | compose an agent preset; a service row there needs an `isolate` realm | 
| Add shell execution | register a `ctx.shell` backend; the local one spawns through`ctx.subprocess` | 
| Add persistent terminal execution | register a `ctx.terminals` backend plus`dsh-tool-terminal` | 
| Add a human command | register on `ctx.commands` ; it dispatches without a model turn | 
| Manage background jobs | register on `ctx.jobs` ;`job_*` tools read or stop jobs | 
| Start a Session from an external webhook | register a trusted rule on `ctx.webhookRuntime` and mount a provider adapter | 
| Add filesystem access or policy | register a `ctx.fs` provider or listen to`fs/*` events | 
| Confine spawned processes | use a `ctx.sandbox` backend; consumers wrap argv before spawning | 
| Intercept a request, tool, or turn | use its `agent/*` or`tools/*` event;`agent/turn-stopping` stops a turn | 
| Add model-facing context | call `agent.inject()` ; it lands in the next admitted request | 
| Add UI or editor integration | drive `ctx.agents` and render from`session/event` | 
| Add a Web Client Chat node | register a `ConversationNodeDefinition` + keyed renderer | 
| Add durable session state | extend `SessionEventMap` ; render and replay from the log | 
| Generate session titles | register the sole `ctx.sessionTitle` provider | 
| Manage a same-session objective | use `ctx.goals` ; continue through`agent/*` | 
| Fork a session at a turn boundary | `ctx.agents.create({ sessionId, seed, meta: { parentSession, seedLength } })` — only agent-loop-published sessions persist | 
| Store sessions in a new backend | implement `SessionPersistence` (`create` /`open` /`stat` /`list` /`export` ) over the shared handle scaffolding | 
| Scope a registration to one agent | use that agent's `agent.ctx` | 

The extension cookbook maps features to capabilities and indexes the step-by-step guides for packages, tools, LLM adapters, and settings pages. The Conversation subsystem owns Chat-node assembly.
