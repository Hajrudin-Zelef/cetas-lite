---
id: collect-korben/korben/zcode-zai-glm-coding
title: "ZCode - L'app maison de z.ai pour coder avec GLM"
domain: korben
role: reference
task: article
actors: ["Anthropic", "Apple", "China", "Z.ai"]
dates: ["2026-09-23"]
keywords: ["glm", "agent", "agentic", "agents", "claude", "mcp", "pricing"]
source: docs/RAG/Collect RAG/01_korben/zcode-zai-glm-coding.md
source_anchor: ""
source_lines: [1, 50]
sha256: a501fb6ea9dd0ddb4ddcc90d262783b703a5cc6a56a2d885e388560baa21ecc8
---

# ZCode - L'app maison de z.ai pour coder avec GLM

## Metadata

- **Source** : https://korben.info/zcode-zai-glm-coding.html
- **Site** : Korben.info
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Following his article on GLM 5.2, Korben presents ZCode, the desktop application from the z.ai team (the Chinese lab formerly known as Zhipu AI) for coding with AI agents. It runs on Mac, Windows and Linux (beta for the latter) and is designed for long tasks. The idea is that you set "Goals" and the agent plans, executes and verifies its own work step by step while you follow the progress.

It provides a real environment with a file manager, terminal, Git panel and live browser preview. It speaks the MCP protocol, can launch several agents in parallel, and you can even drive your tasks remotely from WeChat or Feishu. Everything runs on the same GLM Coding Plan as the rest of their ecosystem, starting at $18 per month.

The article recalls that GLM 5.2 plugs directly into Claude Code because z.ai's API is Anthropic-compatible, with the small launcher shared the day before. But ZCode is mainly a good showcase for z.ai, to keep users in their app rather than using the competition. Nothing is mandatory: ZCode should be seen as one more option for coding with AI.

The app is very recent, and browsing their feedback repo, the author sees that it still suffers painful crashes on Apple Silicon, notably errors causing loss of unsaved work and a window that sometimes stays black even after restarting the app. That is a bit off-putting for something meant to handle long work sessions. He thinks it will quickly stabilize, given how well the Z.ai devs work. For existing GLM Coding Plan subscribers, trying it costs nothing to see if the Goals fit their workflow; otherwise, the API plus your usual tool will do the job just as well.

## Key points

- ZCode is z.ai's (ex-Zhipu AI) desktop app for coding with AI agents, on Mac, Windows and Linux (beta).
- Designed for long tasks: you set Goals and the agent plans, executes and verifies its work step by step.
- Includes file manager, terminal, Git panel and live browser preview; supports the MCP protocol.
- Can run several agents in parallel and be driven remotely via WeChat or Feishu.
- Runs on the same GLM Coding Plan as the rest of the ecosystem, from $18/month.
- GLM 5.2 plugs into Claude Code via z.ai's Anthropic-compatible API; ZCode is mainly a vendor showcase.
- Still unstable: crashes on Apple Silicon causing loss of unsaved work and black windows after restart.
- Author expects quick stabilization; API + usual tool remains a valid alternative.

## Technical data / figures

| Item | Value |
|---|---|
| Application | ZCode |
| Vendor | z.ai (formerly Zhipu AI) |
| Platforms | Mac, Windows, Linux (beta) |
| Core concept | Goals: agent plans, executes, verifies step by step |
| Features | File manager, terminal, Git panel, live browser preview, MCP protocol |
| Parallelism | Multiple agents in parallel |
| Remote control | WeChat, Feishu |
| Plan | GLM Coding Plan, from $18/month |
| Known issues | Crashes on Apple Silicon, lost unsaved work, black window after restart |
| API compatibility | Anthropic-compatible (works with Claude Code) |

## Why this source matters for the RAG

This article describes a vendor-native agentic coding application (ZCode) and its capabilities, pricing and current stability issues, complementing the GLM 5.2 coverage. It is useful for questions about AI coding agents, agentic workflows (Goals), MCP support, and cross-platform desktop AI tooling.
