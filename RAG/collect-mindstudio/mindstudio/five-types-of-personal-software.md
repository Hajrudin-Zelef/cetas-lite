---
id: collect-mindstudio/mindstudio/five-types-of-personal-software
title: "The Five Types of Software You Can Build with AI (And How to Pick One)"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "Apple", "OpenAI", "Z.ai"]
dates: ["2026-08", "2026-09-23"]
keywords: ["agent", "claude", "glm", "research"]
source: docs/RAG/Collect RAG/02_mindstudio/five-types-of-personal-software.md
source_anchor: ""
source_lines: [1, 54]
sha256: 32da02e2f4ef3b6edbb7cc05d97d4cb20378474f2b23e11f1fa41e31116f8e10
---

# The Five Types of Software You Can Build with AI (And How to Pick One)

## Metadata

- **Source**: https://www.mindstudio.ai/blog/five-types-of-personal-software
- **Site**: MindStudio
- **Type**: Article
- **Language**: en
- **Verification status**: ✅ reachable
- **Collection date**: 2026-09-23

## Full summary

This article presents a framework for choosing the right software shape before building with AI. Nearly every personal software idea falls into one of five shapes: a **local tool** (runs on one computer), a **web app** (opens from a link), a **native phone app**, a **background service** (no screen), or a **hardware project** (connects software to the physical world). Determining which shape an idea needs, before touching any AI builder or coding tool, determines the tooling and complexity involved.

**Why shape matters before picking a tool.** Most people start a software project by picking a tool first — opening an AI builder and typing a description. Every idea has a shape determined by where its data comes from, who needs to see it, what device it needs to reach, and what happens if it fails. Example: a ferry-arrival tracker listening to a ship's radio broadcast needs a physical receiver near an antenna, a small always-on computer, and a simple phone display; a shared household maintenance tracker needs a database two people can edit, photo storage, and a browser interface. Neither needs an app store listing or millions-of-users infrastructure. Once the desired change is described, the shape is usually implied.

**Local tool.** Runs on one computer, stores data in files or a small local database. No login system, no server, often no website. Simplest shape; fits a document organizer, private research assistant, or photo rename/sort utility. Best when only one person uses it and it never needs phone access or sharing.

**Web app.** Opens from a link, runs in a browser, can be pinned to a phone's home screen. The recommended default for most personal software: one version works across laptop, iPhone, and Android, and can be shared without app store review. This is the shape hosted AI builders are best at — **Lovable, Replit, and Bolt** let builders describe what they want in plain language, then generate a working preview and publish, all in a pre-configured environment. Lovable suits clean web apps on the shortest path from idea to working interface. It presents a real decision: let the tool manage its own database (Lovable Cloud) or connect a separate database service (**Supabase**, built on the Postgres standard) — the easier path keeps everything in one service; the deliberate path keeps data portable.

**Native app.** What most people picture when they hear "app," but it requires more work: app store submission, platform-specific development, ongoing maintenance across updates. Native development is only necessary when the software depends on a phone capability a web app can't reliably access: persistent push notifications, Bluetooth connections, background location tracking, or NFC. Otherwise, start with a web app.

**Background service.** Software with no screen. Wakes on a schedule or event, does a job, sends results elsewhere. Examples: checking a public record every morning, processing a file when it arrives, alerting when a sensor crosses a threshold. Often the simplest projects in practice — take data from one place, move it to another, on a timer or trigger — and frequently the backbone of a larger personal project, quietly collecting data that a web app later displays.

**Hardware project.** Software living close to the physical world — reading a signal, controlling a device, capturing sensor data. Small computers that recur: a general-purpose single-board computer (real processing power near the sensor), a simple microcontroller board (narrow jobs like reading one sensor), and home automation platforms (connecting existing smart devices).

**Choosing between Lovable, Replit, Codex, and Claude Code.** These solve different problems despite being compared as interchangeable. **Lovable, Replit, and Bolt are hosted builders**: AI coding agent + ready-made environment + publish button. Lovable leans toward fast, clean web apps. Replit is a stronger fit when a project needs more than a browser interface (scripting language, always-running server, scheduled job) because it keeps environment, database, and publishing in one service. Bolt offers a similar full-stack path (interface, database, authentication, storage). The tradeoff for all three: convenience in exchange for dependence on a single service. For more control, builders separate every piece: a coding agent (Codex or Claude Code) working on files in **GitHub**, a dedicated database (**Supabase**), and hosting (**Vercel**) — more accounts and moving parts, but any piece can be swapped without disturbing the rest.

**Two things people conflate.** The coding agent (the tool that reads files, writes changes, shows previews) and the underlying model powering it (which might come from OpenAI, Anthropic, or an open-source option like GLM). A coding agent and its model don't have to come from the same company; picking one doesn't lock in the other.

## Key points

- Five software shapes: local tools, web apps, native apps, background services, and hardware projects — each implying different tools and complexity.
- Web apps are the default starting point: cross-device, shareable, no app store approval or extra infrastructure.
- Native apps are only necessary for deep phone features: push notifications, Bluetooth, background location, NFC.
- Hosted builders (Lovable, Replit, Bolt) package a coding agent + ready-made environment so nontechnical builders never touch a command line.
- Advanced builds separate pieces (coding agent + GitHub + Supabase + Vercel), trading convenience for portability.
- Background services are simple, screenless timer/trigger jobs often serving as the backbone of larger projects.
- Hardware projects use single-board computers, microcontrollers, or home automation platforms.
- The coding agent and the underlying model are separate decisions — a harness like Codex or Claude Code can run different models.

## Technical data / figures

- Five shapes: local tool, web app, native app, background service, hardware project.
- Hosted builders: Lovable (fast clean web apps), Replit (more than browser: scripting, servers, scheduled jobs), Bolt (full-stack: interface, DB, auth, storage).
- Advanced stack: coding agent (Codex/Claude Code) + GitHub + Supabase (Postgres-based) + Vercel.
- Native-app triggers: push notifications, Bluetooth, background location, NFC.
- Hardware options: single-board computers, microcontroller boards, home automation platforms.
- Model/harness separation: e.g., Codex or Claude Code can run OpenAI, Anthropic, or open-source (GLM) models.

## Why this source matters for the RAG

Provides a clear, current (August 2026) decision framework for choosing software shape and build tools — with named platforms (Lovable, Replit, Bolt, Codex, Claude Code, Supabase, Vercel) and practical guidance. This is up-to-date, practical knowledge that helps the RAG answer accurately about AI software building choices, tool selection, and the coding-agent-vs-model distinction.
