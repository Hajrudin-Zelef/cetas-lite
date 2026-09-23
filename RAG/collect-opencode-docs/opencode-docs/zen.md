---
id: collect-opencode-docs/opencode-docs/zen
title: "Zen - opencode documentation"
domain: opencode-docs
role: reference
task: documentation
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "United States", "Z.ai", "xAI"]
dates: ["2026-09-23"]
keywords: ["agents", "astra", "benchmark", "claude", "cost", "deepseek", "fable 5", "gemini", "gemini 3.8", "glm", "grok", "grok 4"]
source: docs/RAG/Collect RAG/04_opencode_docs/zen.md
source_anchor: ""
source_lines: [1, 56]
sha256: a37bcc7f3b36d6c0472a187edffd559a9542973d41a03c39ab5c16fae4d92f1a
---

# Zen - opencode documentation

## Metadata

- **Source** : https://opencode.ai/docs/fr/zen/
- **Site** : opencode.ai
- **Type** : Documentation
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

**OpenCode Zen** is a curated, tested and verified list of models provided by the OpenCode team. It behaves like any other provider in OpenCode: you sign in to OpenCode Zen, add billing details, copy your API key, run `/connect` in the TUI, select OpenCode Zen, paste the key, then run `/models` to see recommended models. Billing is per request and you can top up credits. It is entirely optional.

**Background:** although many models exist, only a few work well as coding agents, and providers are configured very differently, producing highly variable performance and quality. Using an aggregator like OpenRouter does not guarantee the best version of a model. OpenCode therefore tested selected models, discussed optimal serving with their teams, worked with chosen providers, compared model/provider combinations, and published a confidently recommended list. Zen is the AI gateway that grants access to these models.

**Endpoints:** models are exposed under `https://opencode.ai/zen/v1/...` — `/responses` (@ai-sdk/openai), `/messages` (@ai-sdk/anthropic), `/chat/completions` (@ai-sdk/openai-compatible), and `/models/<model>` (@ai-sdk/google for Gemini). The catalog includes GPT 6 Astra/Sol/Luna, GPT 5.6 Sol/Terra/Luna, GPT 5.5/Pro, GPT 5.4 family, GPT 5.3/5.2/5.1 Codex variants, GPT 5 family, Claude Fable 5.1/5, Opus 5.5/5/4.8/4.7/4.6/4.5, Sonnet 5/4.6/4.5, Haiku 4.5, Gemini 3.8/3.7/3.6/3.5 Flash, Gemini 3.1 Pro, Gemini 3 Flash, Grok 4.7/4.6/4.5, Grok Build 0.1, Muse Spark 1.3/1.2, Qwen3.8 Flash, Qwen3.7 Max/Plus, Qwen3.6/3.5 Plus, DeepSeek V4.1 Flash/Pro/Flash/Flash Vision Exp, MiniMax M3/M2.7/M2.5, GLM 5.3 Flash/5.3/5.2/5.1/5, Kimi K3/K2.7 Code/K2.6/K2.5, plus free models (Big Pickle, MiMo-V2.6-Flash Free, MiMo-V2.5 Free, Ling 3.0 Flash Fin Free, Nemotron 3 Ultra/3.5 Lightning Free, Muse Spark 1.3 Contributor Free, Jev 1.13 Free). Config model IDs use `opencode/<model-id>` (e.g. `opencode/gpt-5.5`).

**Jev** is a System One model by TypeSafe AI that returns structured decisions instead of text, evaluating typed questions (yes/no `noul`, multiple-choice `choice`, rating-grid `score`) at `https://opencode.ai/zen/v1/systemone`.

**Pricing** is pay-per-use, listed per 1M tokens (e.g. MiniMax M3 $0.30/$1.20; Claude Opus 5 $5/$25; GPT 6 Astra $10/$50; GPT 6 Luna $0.10/$0.50). Low-cost models (Haiku, Nano, Flash) may appear in usage history because OpenCode uses them to generate session titles. Credit card fees are passed at cost (4.4% + $0.30 per transaction).

**Auto-reload:** if balance drops below $5, Zen auto-reloads $20 (configurable/disableable). **Monthly limits** can be set per workspace and per member. **Deprecated models** include GPT 5.2/5.1/5 Codex, Claude Opus 4.1, Sonnet 4, Haiku 3.5, Gemini 3 Pro, MiniMax M2.5/M2.1, GLM 5/4.7/4.6, Kimi K2.5/K2 Thinking/K2, Qwen3 Coder 480B. **Privacy:** all models hosted in the US; providers use zero-retention and no training except free models (Big Pickle, MiMo free tiers, Ling Fin Free, Nemotron NVIDIA endpoints for trial only, Muse Spark Contributor Free). OpenAI/Anthropic APIs retain requests 30 days.

**Teams:** workspaces are free in beta; roles are **Admin** (models, members, API keys, billing) and **Member** (own keys only); admins set per-member spend limits and enable/disable specific models. **Bring Your Own Key:** you can use your own OpenAI/Anthropic keys while accessing other Zen models, billed directly by the provider. Goals: benchmark best model/provider combos, offer highest quality without degrading performance, pass through price drops at cost, and avoid lock-in.

## Key points

- OpenCode Zen is a **curated AI gateway** of tested models, usable like any OpenCode provider.
- Setup: sign in → add billing → copy API key → `/connect` → `/models`; billing per request.
- API base `https://opencode.ai/zen/v1/` with `/responses`, `/messages`, `/chat/completions`, `/models/<model>`; config IDs use `opencode/<model-id>`.
- **Jev** is a structured-decision model at `/v1/systemone` supporting `noul`, `choice`, `score` questions.
- Auto-reload triggers $20 when balance < $5; monthly limits per workspace/member; deprecated models listed.
- Teams: free beta, Admin/Member roles, per-member spend limits and per-model enable/disable.
- BYOK supported for OpenAI/Anthropic; privacy is US-hosted, zero-retention except free/contributor models.

## Technical data / figures

| Item | Value |
| --- | --- |
| API base | `https://opencode.ai/zen/v1/` |
| Endpoints | `/responses`, `/messages`, `/chat/completions`, `/models/<model>` |
| Model ID format | `opencode/<model-id>` |
| Models endpoint | `https://opencode.ai/zen/v1/models` |
| Jev endpoint | `https://opencode.ai/zen/v1/systemone` |
| Auto-reload threshold | Balance < $5 → reload $20 |
| Card fees | 4.4% + $0.30 per transaction |
| Example pricing | MiniMax M3 $0.30/$1.20; Claude Opus 5 $5/$25; GPT 6 Astra $10/$50 |
| Roles | Admin, Member |
| Data retention | 0-day (most), 30 days (OpenAI/Anthropic) |
| Deprecated example | GPT 5.2 Codex (Jul 23, 2026), Claude Opus 4.1 (Aug 5, 2026) |

## Why this source matters for the RAG

It details OpenCode Zen's catalog, endpoints, pricing, billing rules, team management and privacy policies, making it a central reference for model access and cost questions. It also clarifies structured-decision usage via Jev and the distinction between free, contributor and standard models.
