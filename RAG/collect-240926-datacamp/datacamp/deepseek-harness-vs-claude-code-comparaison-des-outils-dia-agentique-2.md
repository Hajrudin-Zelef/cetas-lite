---
id: collect-240926-datacamp/datacamp/deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique-2
title: "Curriculum"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "DeepSeek"]
dates: []
keywords: ["agents", "alignment", "benchmark", "benchmarks", "claude", "compute", "cost", "deepseek", "latency", "license", "pricing", "sandbox"]
source: docs/RAG/clean_en/datacamp/deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique.md
source_anchor: ""
source_lines: [120, 250]
sha256: f28436ad261a226983c027afd8b3522dc27ddd0ea8a7d8f0a3f5cc216326cfcc
---

# Curriculum

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

