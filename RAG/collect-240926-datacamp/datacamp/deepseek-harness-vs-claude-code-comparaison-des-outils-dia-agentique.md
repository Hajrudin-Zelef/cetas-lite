---
id: collect-240926-datacamp/datacamp/deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique
title: "Curriculum"
domain: datacamp
role: reference
task: reference
actors: ["AWS", "Anthropic", "DeepSeek", "OpenAI", "vLLM", "xAI"]
dates: []
keywords: ["agent", "agentic", "agents", "alignment", "bedrock", "benchmark", "benchmarks", "claude", "compute", "cost", "deepseek", "grok"]
source: docs/RAG/clean_en/datacamp/deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique.md
source_anchor: ""
source_lines: [1, 301]
sha256: 0d50daeb3d0c54e11e67377f49e1fdcf887021cb62205c46f9e7ba0388c87fe6
---

# Curriculum

<!-- source: https://www.datacamp.com/fr/blog/deepseek-harness-vs-claude-code -->

# Curriculum

Most code agent comparisons boil down to a race for speed, price, number of tools, or benchmark scores. This framing misses the real question that distinguishes these two approaches: how far can a developer go in replacing the components of the agent runtime?

DeepSeek Harness and Claude Code answer this question at different levels. Harness exposes model adapters, storage, sandboxes, and the agent loop as replaceable plugins. Claude Code embeds its own loop around Claude and provides extension points for the workflow. Here, this boundary matters more than a simple comparison of DeepSeek and Claude models.

I gave both tools the same broken repository and the same Claude model to observe what the runtime changes. A single case study is not enough to decide between the products, but it shows why the harness is not a background detail. I focus on the differences in model choice, configuration, verification, execution logs, and cost.

## Key takeaways

- If the runtime is part of the work: Harness exposes the model adapter, storage, sandbox, and agent loop as replaceable plugins, with support for multiple model providers.
- If the focus is application code: Claude Code embeds more runtime and extends its workflow via Skills, hooks, MCP, and related features.
- Test with the same model: both produced a byte-identical patch and passed the original test suite. In this Windows trial, Claude Code reported 55.0 seconds; Harness recorded 125.4 seconds after three approval requests.
- Cost: DeepSeek Harness has no license fee; model usage is billed by the chosen provider, with possible infrastructure costs. Claude Code is included in paid Claude plans.
- Simple rule: start with Claude Code for everyday application work. Choose Harness when modifying or inspecting the runtime is part of the mandate.

## Introduction to AI agents

## DeepSeek Harness vs Claude Code: quick comparison

| Dimension | DeepSeek Harness | Claude Code | 
|---|---|---|
| What it is | Agent runtime with replaceable components | Packaged code agent | 
| License and status | MIT, developer preview, no stable version | Proprietary, production product | 
| Models | DeepSeek, Anthropic, OpenAI, Bedrock, Vertex, Azure, local | Claude, direct or via Bedrock and Vertex | 
| Replaceable loop | Yes, substitutable plugin extension point | No, extensions graft around it | 
| Unit of extension | Cordis plugin, any runtime capability | Plugin grouping skills, hooks, agents, MCP | 
| Execution history | Append-only typed event log, replayable | JSONL transcript, checkpoints, OpenTelemetry | 
| Surfaces | Local web interface, headless CLI, Python SDK | Terminal, IDE, desktop, web, Slack, CI | 
| Configuration | Requires provider setup and mastery of profiles | Lighter default configuration; permissions, hooks, MCP, and project optional | 

## What is DeepSeek Harness?

DeepSeek Harness is an open-source agent harness currently in developer preview. It provides the runtime in which models use tools and context, and it can run models from providers other than DeepSeek. Its README warns that updates may break existing configurations.

If you want to try it, start with our DeepSeek Harness tutorial.

### How DeepSeek Harness works

The agent loop is implemented as a plugin. Harness uses Cordis to compose model adapters, tools, sessions, storage, sandboxes, and the loop.

Cordis allows plugins to discover services and exchange events. Dependency changes load or unload plugins; hot reloading cleans up stale listeners and background tasks.

The default profile exposes the loop as a configuration. `dsh --profile headless --dump-default-config` includes this entry:

```
- id: agent-loop
  name: '@deepseek-ai/dsh-agent-loop'
  config:
    agents: []
```
This entry can be replaced via `cordis.patch.yml`, proof that the loop itself is not fixed.

Plugin panel showing runtime components. Video by the author.

### The four execution modes

DeepSeek Harness offers four distinct execution modes.

- **Standard**: full code agent mode
- **PTC** /**Code**: combines multiple tool calls into a single TypeScript program
- **Minimal**: keeps only persistent bash and str_replace_editor
- **Creator**: for inspecting and testing the runtime

DeepSeek used *Minimal* mode for its benchmarks.

## What is Claude Code?

Claude Code is Anthropic's proprietary code agent and an agentic harness for local or managed use via the terminal, IDE, desktop, web, Slack, and CI.

The CLI launches with `claude` in the project directory. This directory becomes the default file scope, and the session reads project instructions from `CLAUDE.md`. Users can expand file access or add external services afterward.

Anthropic presents these interfaces as ways to access Claude Code. Local sessions run on your machine. Cloud sessions run in managed environments or on servers operated by your organization. Remote Control allows you to drive local work from a browser.

The best starting points are our Claude Code tutorial and the Claude Code best practices guide.

### Built-in loop and extensions

Anthropic describes Claude Code as the agentic harness around Claude, with a repeated loop of context, action, and verification. Users can steer it during execution.

Claude Code stores sessions locally and compacts old context as the window fills. `CLAUDE.md` and automatic memory retain selected instructions and project details from one session to the next. Subagents use separate context windows and return summaries to the parent session.

### How Claude Code's extension layer works

Claude Code supports several extension mechanisms. `CLAUDE.md`, skills, hooks, MCP, subagents, plugins, Agent Teams, and the Agent SDK all extend its workflow.

These extensions operate around Claude Code's built-in loop. Hooks can enforce tool-call rules, while the Agent SDK provides tools and context management in code.

## DeepSeek Harness vs Claude Code: architecture and control

DeepSeek Harness exposes lower-level, replaceable runtime elements. Claude Code keeps its built-in loop and supports extensions around it.

### Replaceable runtime vs packaged agent

DeepSeek Harness treats the runtime as configurable infrastructure: model adapters, storage, sandboxes, and the loop can be replaced via profile entries. Claude Code keeps its built-in loop fixed and extends the workflow around it.

Harness can use this mechanism to swap model providers or other parts of the runtime. The word "plugin" does not have the same scope here:

- Claude Code plugins aggregate features around the loop.
- DeepSeek Harness plugins can define parts of the runtime itself.

### Model support and provider choice

DeepSeek Harness supports more model providers. It runs DeepSeek, Anthropic, OpenAI, clouds, and compatible local endpoints. Claude Code runs… Claude. In other words, DeepSeek Harness can host Claude, but Claude Code cannot host DeepSeek.

This is what allowed me to keep the same model during testing. Comparing a DeepSeek model in Harness to Claude in Claude Code changes two variables at once.

Provider choice involves configuring credentials and the endpoint, including exact model IDs. Claude Code controls the model family and most request parameters.

Switching providers does not guarantee identical behavior. Models can differ in tooling format, context size, or reasoning controls. Harness keeps plugin configuration accessible, but the provider adapter must remain compatible with the chosen endpoint.

### Session logs and traceability

DeepSeek Harness records prompts, context injections, tool calls, and permission decisions in an append-only event stream. Its "trajectory" view shows the origin of each record and supports resume, fork, search, and replay.

Claude Code stores JSONL transcripts, supports resume and fork, and emits OpenTelemetry traces.

Both expose session history, but DeepSeek Harness puts the execution path more front and center in the UI.

"Trajectory" view reconstructing a full run. Video by the author.

### Execution environments and permissions

DeepSeek Harness uses bubblewrap or Landlock on Linux, Seatbelt on macOS, and an ACL-restricted token on Windows. If the sandbox cannot start, Harness blocks the command. Claude Code offers permission modes as well as allow, ask, and deny rules. In this run, the Windows sandbox influenced the verification path and timing.

Access labels differ, making strict alignment of settings impossible.

- 
DeepSeek Harness permission presets include `read-only`, `workspace-write`, and `danger-full-access`.
- 
Claude Code separates automatic edits, manual approval, planning, and command access by rules. It also takes file snapshots before modification for rollback without Git.

On current Pro, Max, and Team sessions in the terminal and VS Code, Auto mode is Claude Code's built-in startup mode. A classifier reviews actions in the background. The test explicitly used `acceptEdits`; its permission behavior therefore reflects the test configuration rather than the current default.

This nuance matters when comparing the number of approvals.

The execution location also differs. Harness runs primarily on the machine hosting its web interface or a headless process. Claude Code can run locally, in Anthropic-managed cloud environments, or on self-hosted infrastructure. Remote Control adds a browser interface while execution stays local.

## Hands-on test: DeepSeek Harness vs Claude Code

To check whether these runtime differences affected execution, I gave both tools the same one-line TypeScript bug. If you only want the timed result, this is the section to read.

### Test setup: same model, same repository, same prompt

I wrote a TypeScript "habit-streak" library with one faulty line. It measured calendar days by rounding elapsed milliseconds.

- A check-in at 23:50 followed by another at 00:10 appeared to belong to the same day.
- A check-in at 08:00 followed by 20:00 the next day appeared to indicate a missed day.

Ten frozen tests revealed three failures.

Each new clone had its dependencies installed before timing. Both agents received the same instruction and used Claude Sonnet 5. The repository and prompt were fixed; tools, permissions, and sandbox behaviors remained specific to each product.

The permission configurations were not equivalent. DeepSeek Harness started in `workspace-write`, but its Windows sandbox failed to launch. The web interface requested approval, and I granted `danger-full-access` three times. Claude Code used `accept-edits` and could only run the test command via bash.

This targeted task used only local code and tests. The agents had to find and fix a faulty line, then run the suite. Dependencies were preinstalled before the timer, although Harness chose to rerun `npm install`.

### Results: identical patch, different approval paths

Claude Code reported 55.0 seconds. It ran the suite, found the bug, modified a source file, and reran the tests. Ten out of ten passed.

In this Windows test, DeepSeek Harness took 125.4 seconds. It found the same unused `daysBetween` helper, applied the same fix, and successfully ran the original suite.

The diffs were byte-identical:

```
-import { DAY_MS } from './dates.js';
+import { daysBetween } from './dates.js';
 
 function gapInDays(earlier: number, later: number): number {
-  return Math.round((later - earlier) / DAY_MS);
+  return daysBetween(earlier, later);
 }
```
The approval path and shell-side retries account for most of the time gap.

### Why DeepSeek Harness was slower on Windows

DeepSeek Harness's `workspace-write` sandbox failed to start on Windows, so the web interface requested a broader permission mode three times. After approval, Harness ran the original suite, modified `src/streak.ts`, then reran the suite. Ten out of ten tests passed.

This result says nothing about Harness's behavior on Linux or macOS. Since both runs used the same model ID, the time gap cannot come from a difference in model family. However, distinct permissions and tool paths shaped each run.

Both runs achieve ten passing tests. Image by the author.

### What the same-model test isolates

Sharing the same model ID did not make the requests identical. Each product provided its own system prompt, tool descriptions, context, and permission rules. Each of these inputs can influence the model's next response.

The setup therefore controls the model family better than a "DeepSeek model vs Claude" comparison, without isolating all variables.

I only ran one timed trial per configuration. Repeats, in randomized order, would be needed before treating times or tool counts as stable performance measurements.

## DeepSeek Harness vs Claude Code: costs, benchmarks, and testing limitations

A Windows case study cannot settle the general performance or pricing comparison. It does, however, show why benchmark details matter.

### Benchmark caveat: DeepSeek used Minimal mode

The scores published by DeepSeek used Minimal mode, not the Standard mode used in this comparison. Independent tests with another harness reported lower results. I did not reproduce them. Results in Minimal mode do not establish Standard mode performance and do not measure Claude Code.

### Cost: API calls vs Claude subscriptions

DeepSeek Harness has no license fee, but model and infrastructure costs apply. DeepSeek's API uses peak and off-peak pricing. Claude Code is included in paid Claude plans. Check Anthropic's pricing page before budgeting.

The run with Harness cost about $0.11 at Anthropic's public API rates. Claude Code showed $0.349 while using subscription credentials. Since these figures come from different billing systems, this is not a direct price comparison.

Claude's individual plans list Pro at $20 per month and Max at $100 or $200 per month, while DeepSeek Harness has no equivalent subscription; the model provider bills. Local models eliminate API fees but still consume machine resources.

Teams should also account for hosted sandboxes and self-hosted compute when those services are billed outside the model. These costs do not appear in the Harness license.

### Stability: developer preview vs production product

Harness was still in release candidate during testing. Teams should pin package versions, as updates can break profiles or saved sessions; Claude Code is already in production.

## Limitations of DeepSeek Harness vs Claude Code

Neither system covers all workflows. Their limitations stem from different design choices.

### DeepSeek Harness limitations

Harness exposes more runtime components, which also leaves more configuration and testing to the user:

- 
The project remains in developer preview, and updates can break profiles or saved sessions
- 
Provider credentials, model IDs, and endpoint rules require manual configuration
- 
Its `workspace-write` sandbox failed to start in this Windows case study and required three `danger-full-access` approvals
- 
The current Python SDK has no Windows wheel

The observation about the sandbox applies only to the Windows execution described above. I did not repeat the test on Linux or macOS.

The v0.1.0-rc.7 release notes mention a fix for persistent Bash latency and an update to `node-pty` for better PTY compatibility. This observation is therefore both version- and Windows-specific.

### Limitations of Claude Code

Claude Code supports more runtime decisions, but its fixed boundaries prevent certain types of experimentation:

- 
It runs Claude models and does not provide a path for other model families
- 
Its built-in session loop is not exposed as a replaceable component
- 
Claude Code requires paid access or separately billed API credits, with usage limits
- 
Long sessions can compact older context; so place persistent rules in `CLAUDE.md`

These limitations matter most when a task requires another model provider or a custom loop.

## DeepSeek Harness vs Claude Code: Which to Choose?

Set the model names aside for a moment. Is your task to finish application code, or to modify the runtime that does the work?

Choose DeepSeek Harness if the runtime itself is the project. It is suitable for multi-provider testing, custom loops or sandboxes, and work requiring detailed execution traces. Its preview status also means accepting version changes and configuration validation.

Choose Claude Code if you are working on an existing codebase and want the agent to explore files, modify code, and run project checks with less runtime configuration. In this Windows test, it ran the original suite without the three full-access approvals required by Harness.

Use both if that split reflects your reality. A team can use Claude Code for daily work and Harness for experimenting with prompts, tools, or the agent loop.

## Conclusion

Both tools can inspect a repository, modify files, and run commands. The difference lies in runtime control: Harness exposes its components as plugins, while Claude Code packs the built-in loop and lets you extend the workflow around it.

If I had to choose a starting point based on this test, I would start with Claude Code for everyday application work. I would go with Harness when it comes to modifying or inspecting the runtime, because that is the very purpose of the task. This is a Windows case study with Harness in preview, not a definitive ranking.

To go further, see our articles on Grok Build, Cursor, OpenCode, Codex, and other alternatives to Claude Code.

## DeepSeek Harness vs Claude Code: FAQ

### Can DeepSeek Harness reuse Claude Code project instructions?

**Yes. Harness reads `CLAUDE.md` and `AGENTS.md`, so repository instructions can be reused. Claude Code Skills remain separate.**

### Can DeepSeek Harness run Claude Code or Codex as a sub-agent?

**Yes. Preview builds include optional sub-agent provider bundles for Claude Code and Codex. Pin the Harness version before relying on their configuration.**

### Does Claude Code allow replacing the model with another?

No, it only runs Claude. Harness can host Claude, but Claude Code cannot host DeepSeek.

### Which platforms currently support the DeepSeek Harness Python SDK?

**The current wheel supports Linux on x64 or arm64 as well as recent versions of macOS arm64. There is no Windows wheel, and the composition example uses full filesystem access.**

### Does Claude Code have platform restrictions like the Harness Python SDK?

Not exactly. It has no native Windows sandbox, so it requires WSL2 for `/sandbox` on Windows, but the CLI itself works on macOS, Linux, and Windows.

### Can DeepSeek Harness Code mode run Python programs?

**Not with the provided backend. PTC mode (also called Code) recognizes programming languages, but the documented runtime currently offers a TypeScript backend.**

### Can DeepSeek Harness run local models?

**Yes. Add a custom provider exposing an OpenAI-compatible endpoint, such as a local Ollama or vLLM server, then register the model ID in DeepSeek Harness.**

I am a data engineer and community builder. I work on data pipelines, cloud, and AI tools, while writing practical and impactful tutorials for DataCamp and emerging developers.
