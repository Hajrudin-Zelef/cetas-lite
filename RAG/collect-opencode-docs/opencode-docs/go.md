---
id: collect-opencode-docs/opencode-docs/go
title: "Go - opencode documentation"
domain: opencode-docs
role: reference
task: documentation
actors: ["Alibaba", "Anthropic", "DeepSeek", "LongCat", "Meta", "Microsoft", "MiniMax", "Moonshot", "OpenAI", "Z.ai", "xAI"]
dates: ["2026-09-23"]
keywords: ["agent", "claude", "copilot", "cost", "deepseek", "glm", "grok", "grok 4", "kimi", "latency", "luna", "muse"]
source: docs/RAG/Collect RAG/04_opencode_docs/go.md
source_anchor: ""
source_lines: [1, 56]
sha256: c0287f7f73fe12f73d3a9a64955ae1ced9a1258a62d56cc094c7162206ddfa70
---

# Go - opencode documentation

## Metadata

- **Source** : https://opencode.ai/docs/fr/go/
- **Site** : opencode.ai
- **Type** : Documentation
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

**OpenCode Go** is a low-cost subscription (USD $10/month) providing reliable, low-latency access to popular **open coding models**. It works like any other provider in OpenCode: you subscribe via OpenCode Zen, obtain an API key, run `/connect` in the TUI, select OpenCode Go, paste the key, then run `/models` to list available models. It is entirely optional and designed mainly for international users needing stable global access. Only one member per workspace can subscribe.

**Background:** open models now approach proprietary performance on coding tasks and are usually much cheaper, but reliable low-latency access is hard to obtain because providers vary in quality and availability. OpenCode tested a curated set of models, worked with selected providers to ensure correct serving, and evaluated model/provider combinations into a recommended list.

**Model list (can change):** Grok 4.7, Grok 4.6, GLM-5.3-Flash, GLM-5.3, GLM-5.2, GLM-5.1, GPT 5.6 Luna, Kimi K3, Kimi K2.7 Code, Kimi K2.6, LongCat-2.0, MiMo-V2.6-Flash/Pro, MiMo-V2.5/Pro, MiniMax M3/M2.7, Muse Spark 1.3/1.2 Contributor (limited regions), Qwen3.8 Max/Flash, Qwen3.7 Max/Plus, Qwen3.6 Plus, DeepSeek V4.1 Flash, V4 Pro, V4 Flash, V4 Flash Vision Exp, Hy4 preview, Hy3.

**Client requirements:** traffic must look like a coding agent, identify with a custom user-agent (e.g. `my-coding-agent/1.0`), and send a stable session ID in `x-opencode-session`. Validated clients include Hermes (PR #101864), Claude Code, Codex, ZCode, Pi, jcode (v0.81.6+), Kilo Code CLI (PR #13752). Problematic clients: DeepSeek Harness (discussion #5495), GitHub Copilot Chat (VS Code #334186), Kimi Code (#3506), MiMo Code (#2317 / PR #2327).

**Usage limits** are monthly dollar amounts, with sub-limits: 5-hour window = 20% of monthly, weekly = 50%, monthly = 100%. Per-model input/output/cached-read prices (per million tokens) and monthly limits are tabulated (e.g. GLM-5.3-Flash $0.15/$0.50, $60; Kimi K3 $3.00/$15.00, $15; Qwen3.8 Max $2.00/$6.00, $15; DeepSeek V4.1 Flash off-peak $0.15/$0.60, $30). Estimated request counts per 5h/week/month are also provided, using assumptions like Grok 4.7 (390 input, 32,500 cached, 120 output tokens/request) and GLM-5.3-Flash (1,000 / 55,000 / 200). DeepSeek peak hours are 01:00–04:00 and 06:00–10:00 UTC weekdays; all other times off-peak. If limits are hit, free models remain usable; with Zen balance, the "Use balance" option lets Go fall back to balance.

**Endpoints:** models are reachable at `https://opencode.ai/zen/go/v1/...` (`/responses` for @ai-sdk/openai, `/chat/completions` for @ai-sdk/openai-compatible, `/messages` for @ai-sdk/anthropic). Config model IDs use `opencode-go/<model-id>` (e.g. `opencode-go/kimi-k3`). The full model list is at `https://opencode.ai/zen/go/v1/models`. Privacy: most models are not used for training with 0-day retention; Grok 4.7/4.6 and GPT 5.6 Luna retain up to 30 days; Muse Spark Contributor models are used for training (no ZDR); DeepSeek ZDR renewed monthly (valid until 30 Sep 2026).

## Key points

- OpenCode Go is a **$10/month** optional subscription for curated **open coding models**, used like any OpenCode provider.
- Setup: subscribe via Zen → get API key → `/connect` → `/models`; one subscriber per workspace.
- Clients must send a custom user-agent and a stable `x-opencode-session` header.
- Usage limits are monthly dollar amounts, split 20% (5h) / 50% (weekly) / 100% (monthly).
- API base: `https://opencode.ai/zen/go/v1/` with model IDs formatted `opencode-go/<model-id>`.
- Free models remain available after limits; Zen balance fallback can be enabled.
- Privacy varies: 0-day retention for most; 30-day for Grok/GPT Luna; training opt-in for Muse Spark Contributor.

## Technical data / figures

| Item | Value |
| --- | --- |
| Price | $10 / month |
| Setup command | `/connect` (select OpenCode Go) |
| Model listing command | `/models` |
| Session header | `x-opencode-session` |
| User-agent example | `my-coding-agent/1.0` |
| API base | `https://opencode.ai/zen/go/v1/` |
| Endpoint patterns | `/responses`, `/chat/completions`, `/messages` |
| Model ID format | `opencode-go/<model-id>` |
| Models endpoint | `https://opencode.ai/zen/go/v1/models` |
| Limit windows | 5h = 20%, weekly = 50%, monthly = 100% |
| Example model | GLM-5.3-Flash: $0.15 in / $0.50 out, $60 monthly |
| Example model | Kimi K3: $3.00 in / $15.00 out, $15 monthly |
| DeepSeek peak hours | 01:00–04:00 & 06:00–10:00 UTC (Mon–Fri) |

## Why this source matters for the RAG

It fully documents OpenCode's low-cost subscription offering, including pricing, model catalog, API endpoints, request/session requirements and privacy terms. This enables accurate answers about how to enable Go, how limits are computed, and how to integrate it via the OpenAI/Anthropic-compatible endpoints.
