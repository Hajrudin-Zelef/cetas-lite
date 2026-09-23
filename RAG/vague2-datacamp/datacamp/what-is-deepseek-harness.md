---
id: vague2-datacamp/datacamp/what-is-deepseek-harness
title: "Qu'est-ce que DeepSeek Harness ? Le runtime d'agent où tout est plugin"
domain: datacamp
role: reference
task: article
actors: ["Anthropic", "DeepSeek", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["agent", "deepseek", "benchmarks", "claude", "license", "mcp", "memory", "open source", "reasoning", "sandbox", "tool calling"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/what-is-deepseek-harness.md
source_anchor: ""
source_lines: [1, 56]
sha256: 105e939ed669296d08e954de57c5022188424730a447baf88d9fd58495085639
---

# Qu'est-ce que DeepSeek Harness ? Le runtime d'agent où tout est plugin

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/what-is-deepseek-harness
- **Site** : DataCamp
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp article explains DeepSeek Harness (`dsh`), an open-source (MIT-licensed) agent runtime from DeepSeek AI. The core idea is **"Agent = Model + Harness"**: the model handles reasoning and generation, while the harness connects that reasoning to a real file system, terminal, tools, and session history so an agent can inspect files, edit code, run tests, and react to failures. It is built on **Cordis**, a plugin framework from the Koishi chatbot ecosystem (created by developer Shigma), and is described in the paper *A Programming Paradigm for Spatiotemporal Composability*.

Two common misconceptions are corrected: Harness is **not** an AI model (the runtime can use DeepSeek, Anthropic, OpenAI, or any OpenAI-compatible endpoint), and it goes **beyond a coding assistant** — Standard mode only resembles one; Minimal and Creator modes change the toolset.

**Architecture.** Cordis provides a shared service directory: plugins use stable keys like `ctx.tools`, `ctx.llm`, and `ctx.sessions` rather than importing provider code. The slogan "everything is a plugin" applies to model adapters, tools, sessions, sandboxes, storage, scheduling, the agent loop, and UI — though Cordis itself is the indispensable substrate. Two Cordis concepts matter: *spatial composability* (a plugin declares dependencies and activates/deactivates based on service availability) and *temporal composability* (removing a plugin reverts its tracked effects like event listeners and prompt sections, but not external actions like shell commands).

A running instance is a plugin tree. Two confusing levels: **runtime profiles** (`web`, `headless`, `sdk`, `sdk-minimal`, `acp`) control how the app launches and which plugin bundles load; **agent presets** (Standard, PTC, Minimal, Creator) control what an active session can use. The agent loop distinguishes a *step* (a model request plus tool calls) from a *turn* (zero or more steps). Sessions use an **append-only typed event journal** — not a chat array — enabling resume, fork, search, replay, and the Trajectory view. Tool execution flows through policy check → execution → result post-processing; **approval** (user confirmation) and **sandbox** (execution limits) are separate controls.

**Four modes:** Standard (general-purpose: file edits, shell, search, skills, planning, goals, subagents, workflows); PTC (programmatic tool calling — the model writes a program against a generated SDK, invoking tools via `run_code`, each still policy-checked; formerly "Code mode"); Minimal (only a persistent shell and a string-substitution file editor, used for model benchmarks); Creator (inspect the runtime and test Cordis plugins in memory).

**Comparisons:** Claude Code keeps its internal loop fixed while supporting project instructions, skills, hooks, MCP, subagents, and an Agent SDK. Codex has open-source CLI/App Server but extends via documented entry points. OpenCode is open-source with a client-server architecture but plugins extend a fixed server core. Harness makes the loop and session store replaceable.

**Limitations:** developer preview with breaking changes (Code→PTC rename, session API changes, removal of optional SQLite); more control means more complexity; local-first storage but external providers/MCP/web tools can send data out; no security audit has been conducted, and sandboxing/approvals reduce but do not guarantee isolation.

## Key points

- `dsh` is an open-source agent runtime, not a model; it can drive DeepSeek, Anthropic, OpenAI, or local models.
- Built on Cordis, a plugin framework making model adapters, tools, sessions, sandboxes, storage, loop, and UI replaceable.
- Agent = Model + Harness: the runtime determines what reasoning can access and do.
- Sessions are append-only typed event journals enabling resume, fork, search, replay, and Trajectory.
- Four modes: Standard, PTC (programmatic tool calling), Minimal (2 tools), Creator (runtime inspection).
- Approval and sandbox are distinct controls; sandboxing limits execution, prompts do not.
- Still a developer preview: breaking API changes, no security audit.
- Local-first storage, but connected providers/MCP/web tools can exfiltrate data.

## Technical data / figures

| Item | Detail |
|---|---|
| Abbreviation | `dsh` |
| License | MIT (open source) |
| Base framework | Cordis (from Koishi ecosystem, by Shigma) |
| Service keys | `ctx.tools`, `ctx.llm`, `ctx.sessions` |
| Runtime profiles | `web`, `headless`, `sdk`, `sdk-minimal`, `acp` |
| Agent presets/modes | Standard, PTC, Minimal, Creator |
| Minimal mode tools | Persistent shell + string-substitution file editor |
| Session model | Append-only typed event journal |
| Paper | "A Programming Paradigm for Spatiotemporal Composability" (arXiv:2608.25512) |
| Version note | Since 0.1.2, Web PTC no longer exposes generic `workflow` tool by default |

## Why this source matters for the RAG

It provides a detailed architectural reference for agent harnesses and runtime plugin design, clarifying the model-vs-runtime distinction central to modern agent tooling. It is valuable for answering questions on DeepSeek Harness, Cordis, agent sessions, and comparisons with Claude Code, Codex, and OpenCode.
