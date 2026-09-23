---
id: collect-opencode-docs/opencode-docs/index
title: "Introduction - opencode documentation"
domain: opencode-docs
role: reference
task: documentation
actors: []
dates: ["2026-09-23"]
keywords: ["agent", "agents"]
source: docs/RAG/Collect RAG/04_opencode_docs/index.md
source_anchor: ""
source_lines: [1, 60]
sha256: c6032dcb49caaf9ff7f95d7ecbbdfc7f7d030b9adbc6d07cd2b5fd05330208d3
---

# Introduction - opencode documentation

## Metadata

- **Source** : https://opencode.ai/docs/fr
- **Site** : opencode.ai
- **Type** : Documentation
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is the **Introduction** page of the OpenCode documentation — the official getting-started guide. OpenCode is described as an **open-source AI coding agent** available as a **terminal-based interface, a desktop application, or an IDE extension**. The page walks a new user through prerequisites, installation, provider configuration, project initialization, and core usage.

**Prerequisites:** a modern terminal emulator (WezTerm or Alacritty, cross-platform; Ghostty or Kitty, Linux/macOS) and API keys for the LLM providers you intend to use.

**Installation:** the simplest method is the install script `curl -fsSL https://opencode.ai/install | bash`. Alternatives include Node package managers (`npm install -g opencode-ai`, `bun`, `pnpm`, `yarn global add opencode-ai`), Homebrew (`brew install anomalyco/tap/opencode` — the OpenCode tap is recommended over the less frequently updated official formula), Arch Linux (`sudo pacman -S opencode` stable, or `paru -S opencode-bin` from AUR), Windows (WSL recommended; Chocolatey `choco install opencode`, Scoop `scoop install opencode`, NPM, Mise `mise use -g github:anomalyco/opencode`, Docker `docker run -it --rm ghcr.io/anomalyco/opencode`), or prebuilt binaries from Releases.

**Configuration:** any LLM provider can be used by configuring its API keys. Beginners are advised to use **OpenCode Zen**, a curated list of tested models. Run `/connect` in the TUI, select `opencode`, go to `opencode.ai/auth`, sign in, add billing, copy the API key and paste it. Other providers are also selectable.

**Initialization:** navigate to the project (`cd /path/to/project`), run `opencode`, then run `/init`. OpenCode analyzes the project and creates an **AGENTS.md** file at the project root — this file should be committed to Git as it helps OpenCode understand structure and coding patterns.

**Usage:** you can ask OpenCode to explain the codebase (use `@` for fuzzy file search, e.g. `How is authentication handled in @packages/functions/src/api/index.ts`). To add features, it is recommended to first create a plan using **Plan Mode** (toggle with `Tab`, indicator bottom-right), iterate on the plan (images can be drag-and-dropped into the terminal), then switch back to **Build Mode** with `Tab` and request changes. For simple changes you can ask directly, referencing patterns in other files. Changes can be reverted with `/undo` (repeatable) and reapplied with `/redo`. Conversations can be shared with `/share`, which creates and copies a link; conversations are **not shared by default**. Finally, users are encouraged to personalize OpenCode via themes, keybinds, code formatters, custom commands, and the OpenCode Config.

## Key points

- OpenCode is an **open-source AI coding agent** for terminal, desktop, and IDE.
- Install via `curl -fsSL https://opencode.ai/install | bash`, npm/bun/pnpm/yarn, Homebrew, Arch, Chocolatey, Scoop, Mise, or Docker.
- On Windows, **WSL is recommended** for best performance and full feature compatibility.
- Configure providers with `/connect`; beginners are directed to OpenCode Zen (curated tested models).
- `/init` generates an **AGENTS.md** file that should be committed to Git.
- Use `Tab` to toggle **Plan Mode** vs **Build Mode**; `@` performs fuzzy file search.
- `/undo` and `/redo` revert/reapply changes; `/share` creates a conversation link (opt-in, not default).

## Technical data / figures

| Item | Value |
| --- | --- |
| Install script | `curl -fsSL https://opencode.ai/install \| bash` |
| npm | `npm install -g opencode-ai` |
| Homebrew | `brew install anomalyco/tap/opencode` |
| Arch (stable) | `sudo pacman -S opencode` |
| Arch (AUR) | `paru -S opencode-bin` |
| Windows (Choco) | `choco install opencode` |
| Windows (Scoop) | `scoop install opencode` |
| Mise | `mise use -g github:anomalyco/opencode` |
| Docker | `docker run -it --rm ghcr.io/anomalyco/opencode` |
| Connect command | `/connect` |
| Auth URL | `opencode.ai/auth` |
| Init command | `/init` |
| Generated file | `AGENTS.md` |
| Mode toggle | `Tab` (Plan / Build) |
| Undo / Redo | `/undo`, `/redo` |
| Share | `/share` |
| File search | `@` |

## Why this source matters for the RAG

It is the canonical entry point for installing, configuring and using OpenCode, so it underpins most onboarding and quick-start questions. It also documents essential commands (`/connect`, `/init`, `/undo`, `/redo`, `/share`), modes, and the AGENTS.md workflow that the rest of the documentation builds upon.
