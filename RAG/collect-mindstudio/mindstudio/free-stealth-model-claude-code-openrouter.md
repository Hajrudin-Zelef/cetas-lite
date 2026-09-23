---
id: collect-mindstudio/mindstudio/free-stealth-model-claude-code-openrouter
title: "How to Run Claude Code for Free Using OpenRouter's Models"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "DeepSeek", "Google", "OpenRouter"]
dates: ["2026-09-23"]
keywords: ["claude", "agent", "agentic", "agents", "cost", "deepseek", "gemini", "memory", "research", "training"]
source: docs/RAG/Collect RAG/02_mindstudio/free-stealth-model-claude-code-openrouter.md
source_anchor: ""
source_lines: [1, 45]
sha256: 7f6808bf538881aa9ab40eb378c945eabd7ba50735bdb31d7e46c45e3e7f89d2
---

# How to Run Claude Code for Free Using OpenRouter's Models

## Metadata

- **Source** : https://www.mindstudio.ai/blog/free-stealth-model-claude-code-openrouter
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This guide explains how to run Claude Code without paying for Anthropic, by routing requests through OpenRouter to free models. The core mechanism is that Claude Code does not verify which model generates responses; it only needs an API endpoint speaking the correct protocol. By editing the `env` section of Claude Code's `settings.json` and swapping the Anthropic API key for an OpenRouter key, users can route requests to any model OpenRouter hosts, including free models.

OpenRouter is a gateway providing one API key for hundreds of models from different providers, including paid options (DeepSeek, Claude Opus, GPT, Gemini variants) and a free tier where input and output tokens are listed at $0. Free listings typically come with rate limits, daily caps, and the risk that providers pull the free tier without notice. A newly released free model, Stealth Ox Alpha (from a provider called "Stealth"), is highlighted. It is anonymous — no public confirmation of who built it, what it was trained on, or where it is hosted — a privacy concern for proprietary code.

Setup steps: create an OpenRouter account, generate an API key, search "free" on the models page and copy a model identifier (e.g. `stealth/ox-alpha`), open Claude Code's settings file (only works in an IDE like VS Code or terminal, not the desktop app which forces an Anthropic model), edit the `env` block replacing the API token and model fields, save, and start a new session. Because OpenRouter rotates free models, the "free models router" option auto-selects from currently free models.

Testing results: functionally the free model operated in the full Claude Code agent loop — it acknowledged goals, searched files, wrote code, called tools, and self-corrected. It built a multi-page landing site (product listings, shopping cart, brand styling, nutrition-fact detail pages) from a single prompt, and used previously stored scripts, API keys, skills, and memory files to pull real YouTube analytics and produce a quarterly/annual report as a Google Sheet. However, speed and reliability suffered: tasks taking minutes with paid Claude models stretched to 40 minutes or several hours. One run failed repeatedly with "upstream idle timeout exceeded"; another got stuck in a retry loop for ~45 minutes. Building a full front end took ~6 hours instead of a fraction of that with a paid frontier model.

The article concludes free models are usable for research, reporting, and knowledge work, but noticeably weaker for complex multi-step engineering and deep technical orchestration (coordinating sub-agents, large codebases, long dependency chains). The practical takeaway: most everyday requests don't need a top-tier model; the dividing line is task complexity — simple well-scoped requests can run on cheap/free infrastructure, while orchestrator-level work still justifies paying for Claude or GPT.

## Key points

- Claude Code can be pointed at OpenRouter by editing the `env` section of `settings.json`; this works only in IDE/terminal, not the desktop app.
- OpenRouter hosts many free models ($0 in/out tokens) with rate limits, daily caps, and volatile availability.
- Stealth Ox Alpha is an anonymous free model; origin, training data, and infrastructure undisclosed — a privacy risk for sensitive projects.
- The free model worked in the full agentic loop (tool calls, file edits, self-correction) and even used stored skills/memory, but was very slow (hours vs minutes).
- Failures included "upstream idle timeout exceeded" errors and retry loops; a front-end build took ~6 hours.
- Recommended use: knowledge work, research, reporting; not complex software builds or deep orchestration.
- Most everyday requests don't need frontier models; task complexity is the deciding factor.

## Technical data / figures

- Setup: OpenRouter account → API key → free model identifier (e.g. `stealth/ox-alpha`) → edit `env` in `settings.json`.
- Free-tier restrictions: rate limits, daily caps, possible removal without notice.
- Test task 1: multi-page landing site built from a single prompt (listings, cart, styling, detail pages).
- Test task 2: YouTube analytics pulled via stored scripts/keys → structured quarterly & annual report in Google Sheets.
- Test failures: "upstream idle timeout exceeded"; retry loop ~45 minutes; front-end build ~6 hours.

## Why this source matters for the RAG

Documents a practical, reproducible method for routing Claude Code through OpenRouter to free/third-party models, useful for cost-optimization RAG. It also provides empirical tradeoff data (speed, reliability, privacy) on free-tier and stealth models — relevant to model-routing and agent-cost knowledge bases.

