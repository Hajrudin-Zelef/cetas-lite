---
id: vague2-datacamp/datacamp/deepseek-harness-vs-claude-code
title: "DeepSeek Harness vs Claude Code : qu'est-ce qui change avec le même modèle ?"
domain: datacamp
role: reference
task: article
actors: ["AWS", "Anthropic", "DeepSeek", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["claude", "deepseek", "agent", "agents", "bedrock", "benchmark", "benchmarks", "cost", "license", "pricing", "sandbox", "sonnet 5"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/deepseek-harness-vs-claude-code.md
source_anchor: ""
source_lines: [1, 57]
sha256: 8191a2f60e6ebcc26d5d4251b9e6e4549493b8b4040c04ca7fbdb9910eb0e329
---

# DeepSeek Harness vs Claude Code : qu'est-ce qui change avec le même modèle ?

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/deepseek-harness-vs-claude-code
- **Site** : DataCamp
- **Type** : Article (comparatif + étude de cas)
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp comparison argues that most coding-agent comparisons miss the real differentiator: how far a developer can replace the agent runtime's components. **DeepSeek Harness** exposes model adapters, storage, sandboxes, and the agent loop as replaceable plugins (via Cordis), while **Claude Code** ships its own loop around Claude and offers extension points *around* that fixed loop.

The author ran a controlled case study: the same broken TypeScript repository and the same model (**Claude Sonnet 5**) given to both tools. The repo was a "habit-streak" library with one faulty line that measured calendar days by rounding elapsed milliseconds, causing three of ten frozen tests to fail. Dependencies were pre-installed before timing; each agent received the same instruction. Permission setups were not equivalent: Harness started in `workspace-write` but its Windows sandbox could not launch, so the web UI requested approval three times and `danger-full-access` was granted; Claude Code used `acceptEdits` and could only run the test command via bash.

**Results:** both produced a byte-identical patch (replacing a `Math.round((later - earlier) / DAY_MS)` calculation with the unused `daysBetween` helper) and passed 10/10 tests. Claude Code reported 55.0 seconds; DeepSeek Harness took 125.4 seconds on Windows, with approval prompts and shell retries explaining most of the gap. The author stresses this is a single Windows run with Harness in release-candidate state, not a stable benchmark.

**Architecture differences:** Harness's agent loop is a plugin (`@deepseek-ai/dsh-agent-loop`) replaceable via `cordis.patch.yml`; it supports DeepSeek, Anthropic, OpenAI, Bedrock, Vertex, Azure, and local OpenAI-compatible endpoints. Claude Code runs only Claude (directly, or via Bedrock/Vertex). Harness logs sessions as an append-only typed event journal with a "trajectory" view (resume, fork, search, replay); Claude Code uses JSONL transcripts, checkpoints, and OpenTelemetry. Harness sandboxing uses bubblewrap/Landlock (Linux), Seatbelt (macOS), ACL-restricted token (Windows), blocking commands if the sandbox fails; Claude Code offers permission modes and allow/ask/deny rules plus pre-edit file snapshots for Git-free rollback. Claude Code's Auto mode is now the built-in start mode on Pro/Max/Team in terminal and VS Code, using a background classifier.

**Cost:** Harness has no license fee but model and infrastructure costs apply (DeepSeek API uses peak/off-peak pricing); the test run cost ~$0.11 at public Anthropic API rates. Claude Code is included in paid Claude plans (Pro $20/mo; Max $100 or $200/mo), and the test displayed $0.349 using subscription credentials — not a direct price comparison. Harness's published benchmarks used Minimal mode, not Standard, and do not measure Claude Code. Harness is developer preview; teams should pin versions.

**Limitations:** Harness requires manual provider/model config, its Windows sandbox failed, and its Python SDK has no Windows wheel (Linux x64/arm64, macOS arm64 only). Claude Code is Claude-only, has a non-replaceable loop, needs paid access, and compacts long sessions (persistent rules belong in `CLAUDE.md`).

**Verdict:** start with Claude Code for ordinary application work; choose Harness when modifying or inspecting the runtime *is* the task; use both if that split reflects your reality.

## Key points

- Core distinction: replaceable runtime (Harness) vs packaged loop with extensions (Claude Code).
- Same model (Claude Sonnet 5), same repo: both produced a byte-identical patch and 10/10 tests.
- Timing: Claude Code 55.0s vs Harness 125.4s on Windows (approvals/shell retries, not model differences).
- Harness runs many providers/local endpoints; Claude Code runs only Claude.
- Harness sessions = append-only typed event journal with trajectory view; Claude Code = JSONL + OpenTelemetry.
- Harness is MIT developer preview with breaking changes; Claude Code is production software.
- DeepSeek's published benchmarks used Minimal mode, so they do not establish Standard-mode performance.
- Harness reads both `CLAUDE.md` and `AGENTS.md`; its Python SDK has no Windows wheel.

## Technical data / figures

| Dimension | DeepSeek Harness | Claude Code |
|---|---|---|
| Type | Runtime with replaceable components | Packaged coding agent |
| License/status | MIT, developer preview | Proprietary, production |
| Models | DeepSeek, Anthropic, OpenAI, Bedrock, Vertex, Azure, local | Claude (direct/Bedrock/Vertex) |
| Loop replaceable | Yes (plugin) | No (extensions around it) |
| Execution history | Append-only typed event journal | JSONL transcripts, checkpoints, OpenTelemetry |
| Surfaces | Local web UI, headless CLI, Python SDK | Terminal, IDE, desktop, web, Slack, CI |
| Sandbox | bubblewrap/Landlock, Seatbelt, Windows ACL token | Permission modes, allow/ask/deny, file snapshots |
| Test time (Windows) | 125.4 s | 55.0 s |
| Test cost | ~$0.11 (public API rates) | $0.349 (subscription credentials) |
| Pricing | Free license; provider bills | Claude Pro $20/mo; Max $100–200/mo |
| Version tested | v0.1.0-rc.7 | Production |

## Why this source matters for the RAG

It provides a rare same-model, same-repo empirical comparison that isolates the runtime from the model, a key concept in agent engineering. It is valuable for questions about DeepSeek Harness, Claude Code, runtime replaceability, agent sandboxing, and real-world agent benchmarking caveats.
