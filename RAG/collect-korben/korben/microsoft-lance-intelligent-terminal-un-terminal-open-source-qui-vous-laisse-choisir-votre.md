---
id: collect-korben/korben/microsoft-lance-intelligent-terminal-un-terminal-open-source-qui-vous-laisse-choisir-votre
title: "Microsoft lance Intelligent Terminal, un terminal open source qui vous laisse choisir votre IA"
domain: korben
role: reference
task: article
actors: ["Anthropic", "Microsoft", "OpenAI"]
dates: ["2026-06", "2026-09-23"]
keywords: ["open source", "agent", "agents", "claude", "copilot", "license", "mit license"]
source: docs/RAG/Collect RAG/01_korben/microsoft-lance-intelligent-terminal-un-terminal-open-source-qui-vous-laisse-choisir-votre-ia.md
source_anchor: ""
source_lines: [1, 52]
sha256: 7a880b614e9e50479dbaf03d3c40c0f0f5728b7e6c53c38bc0990781236dcd46
---

# Microsoft lance Intelligent Terminal, un terminal open source qui vous laisse choisir votre IA

## Metadata

- **Source** : https://korben.info/microsoft-lance-intelligent-terminal-un-terminal-open-source-qui-vous-laisse-choisir-votre-ia.html
- **Site** : Korben.info
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The article reports that Microsoft presented, on 2 June during its Build 2026 developer conference, an experimental version of its terminal called Intelligent Terminal 0.1, an open-source software released under the MIT license that integrates AI assistants directly into the window where commands are typed.

For context, a terminal is the text interface developers and system administrators use daily to launch commands, compile code or pilot a remote server. Microsoft has long offered Windows Terminal, and Intelligent Terminal is a distinct variant, not a replacement. The central element is the "agent pane", a panel anchored to the side of the window that acts as a copilot. It constantly reads what is displayed in the terminal, and when a command fails, it detects the error and proposes an explanation or a fix. You can choose to receive a simple notification in the status bar, or let the assistant apply its corrections directly, depending on the configuration you define.

There is also a trick in the command palette: type a question mark followed by your request, and the tool launches a background task in a dedicated tab without blocking the current session, handy when delegating a long operation without interrupting work. On assistants, Microsoft made an open choice. By default it is GitHub Copilot CLI, the in-house assistant that works from the command line. But Intelligent Terminal accepts any agent compatible with the Agent Client Protocol (ACP), a standard that lets different AI assistants plug into the same tool. You can therefore use Claude Code, Anthropic's coding assistant, or Codex, OpenAI's equivalent, instead of Microsoft's solution, which is rare enough at the Windows maker to be noted.

The software installs from the Microsoft Store or via `winget install Microsoft.IntelligentTerminal`, and sits alongside Windows Terminal without replacing it. Microsoft states that if you do not want AI in your usual terminal, nothing changes for you. There is one detail that sticks, noted by Phoronix, the site specialized in Linux news: despite its open code under the MIT license, Intelligent Terminal currently only runs on Windows. It is also a 0.1 experimental version, which promises bugs and changes before any eventual stable release. The article concludes that an open-source terminal letting the user choose their AI instead of imposing its own is a rather well-thought-out gesture from Microsoft.

## Key points

- Microsoft presented Intelligent Terminal 0.1 at Build 2026 (2 June), an experimental open-source terminal under the MIT license.
- Core feature: the "agent pane", a side panel that reads terminal output, detects errors and proposes explanations or fixes.
- Users can choose a status-bar notification or let the assistant apply corrections directly.
- Command palette shortcut: `?` + request launches a background task in a dedicated tab without blocking the session.
- Default assistant is GitHub Copilot CLI, but any ACP (Agent Client Protocol)-compatible agent works, including Claude Code and Codex.
- Installs via Microsoft Store or `winget install Microsoft.IntelligentTerminal`; sits alongside Windows Terminal, not replacing it.
- Limitation: despite the MIT open code, it currently runs only on Windows (noted by Phoronix).
- Version 0.1 experimental; users without AI needs are unaffected.

## Technical data / figures

| Item | Value |
|---|---|
| Product | Intelligent Terminal 0.1 |
| Company | Microsoft |
| Presented | 2 June 2026 (Build 2026) |
| License | MIT |
| Status | Experimental (0.1) |
| Core feature | Agent pane (error detection, explanations, fixes) |
| Default AI | GitHub Copilot CLI |
| Standard | ACP (Agent Client Protocol) |
| Compatible agents | Claude Code, Codex, GitHub Copilot CLI, others |
| Installation | Microsoft Store or `winget install Microsoft.IntelligentTerminal` |
| Platform | Windows only (for now) |
| Relationship to Windows Terminal | Distinct variant, not a replacement |

## Why this source matters for the RAG

This article documents a notable open-source, AI-integrated terminal from Microsoft that supports multiple AI assistants via the ACP standard, with install methods and current platform limits. It is valuable for questions about AI-integrated developer tools, terminal assistants and ecosystem openness.
