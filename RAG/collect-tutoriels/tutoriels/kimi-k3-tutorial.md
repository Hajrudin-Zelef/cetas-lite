---
id: collect-tutoriels/tutoriels/kimi-k3-tutorial
title: "Kimi K3 : fonctionnalités, benchmarks, API et 5 exemples pratiques"
domain: tutoriels
role: reference
task: tutorial
actors: ["Moonshot", "OpenAI"]
dates: ["2026-07", "2026-09-23"]
keywords: ["benchmarks", "kimi", "agent", "context window", "cost", "multimodal", "parameters", "pricing", "reasoning", "tool calling"]
source: docs/RAG/Collect RAG/07_tutoriels/kimi-k3-tutorial.md
source_anchor: ""
source_lines: [1, 61]
sha256: 645b864d748f29ca82ae8f68660292e96ac676381452ea4e6e5c54fbd61a1df6
---

# Kimi K3 : fonctionnalités, benchmarks, API et 5 exemples pratiques

## Metadata

- **Source** : https://www.datacamp.com/fr/tutorial/kimi-k3-tutorial
- **Site** : DataCamp
- **Type** : Tutorial
- **Language** : fr
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This practical DataCamp tutorial (author Khalid Abdelaty, published 21 July 2026) walks through how to access and use Moonshot AI's Kimi K3, an open model released 16 July 2026 with 2.8 trillion parameters, a one-million-token context window, and native vision. K3 is Moonshot's largest open release, far beyond Kimi K2, and the first they describe as belonging to the 3-trillion-parameter class. The tutorial is the hands-on counterpart to DataCamp's separate launch-analysis blog post.

The article explains access routes: kimi.com (web and mobile), Kimi Work (desktop app for reports and dashboards), and Kimi Code, a terminal coding agent installed via npm under `@moonshot-ai/kimi-code` where the model is chosen with `/model`. K3 usage in Kimi Code requires a paid subscription, and the million-token window needs a higher tier. The API is OpenAI-SDK compatible, requiring Python 3.9+ and an API key from the Kimi platform console, with `base_url="https://api.moonshot.ai/v1"` and model name `kimi-k3`.

The core of the tutorial is five examples. Four use the API and show real token usage and cost, run on 17 July 2026 with the `kimi-k3` model; together they cost about 11 cents cold, or a few cents with caching. Example 1 covers streaming reasoning and final answer, where `reasoning_content` is delivered on a separate channel from `content` (488 output tokens, under one cent). Example 2 demonstrates tool calling with structured JSON output, using `tool_choice="required"`, returning an order summary (5 mechanical keyboards at $89 = $445). Example 3 shows dynamically loading a tool definition mid-conversation via a `system` message containing a `tools` field and no `content` (cheapest call, ~0.2 cents). Example 4 shows automatic context caching reducing long-context cost: a ~33,000-token knowledge base cost ~9.9 cents uncached versus ~1.1 cents cached on 32,512 prefix tokens, nearly a 9x factor, because cached input bills at $0.30/million vs $3.00/million. Example 5 sends a base64 data-URL screenshot to detect layout bugs, and K3 found several issues (offset card, badge over a number, overflow bar) though it missed a low-contrast subtitle; the call cost ~2 cents.

The tutorial closes with limits: only `reasoning_effort="max"` is available; sampling parameters (temperature, top_p, penalties) are locked; outputs can be long and costly; public image URLs are not supported via the API. It concludes K3 is a reasonable default for repo-scale analysis, repeated long-context calls, or multimodal engineering, while a smaller model is simpler for fast, low-cost chat.

## Key points

- Kimi K3: 2.8T parameters, 1M-token context, native vision, released 16 July 2026 by Moonshot AI.
- OpenAI-SDK compatible API; use base_url `https://api.moonshot.ai/v1` and model `kimi-k3`.
- Reasoning is streamed separately via `reasoning_content`; keep it for logs, show only `content` to users.
- `tool_choice="required"` forces at least one tool call; K3 supports dynamic mid-conversation tool loading.
- Automatic context caching cut a ~33k-token request from ~9.9 cents to ~1.1 cents (~9x).
- Only `reasoning_effort="max"`; temperature/top_p/penalties are locked; no public image URLs via API.
- K3 is not a replacement for elders: `kimi-k2.7-code` and `kimi-k2.6` remain better for pure speed tasks.

## Technical data / figures

| Item | Value |
| --- | --- |
| Model | `kimi-k3` |
| Parameters | 2.8 trillion |
| Context window | 1,048,576 tokens |
| Release date | 16 July 2026 |
| API base URL | `https://api.moonshot.ai/v1` |
| Python requirement | 3.9+ |
| SDK | `openai>=1.0` |
| Cached input price | $0.30 / million tokens |
| Uncached input price | $3.00 / million tokens |
| Example 1 (streaming) | 488 output tokens, <1 cent |
| Example 4 (33k tokens) | ~9.9 cents miss vs ~1.1 cents hit |
| Total 4 API examples | ~11 cents cold |

Model lineup:

| Model | Context window | Ideal for |
| --- | --- | --- |
| `kimi-k3` | 1,048,576 tokens | Flagship: long coding, vision, knowledge |
| `kimi-k2.7-code` | 262,144 tokens | Dedicated coding, faster high-speed option |
| `kimi-k2.6` | 262,144 tokens | General text, image and video chat |

Install command: `python -m pip install --upgrade "openai>=1.0"`

## Why this source matters for the RAG

It is a concrete, reproducible API tutorial with real token counts, costs, and working code for K3's tool calling, dynamic tool loading, context caching, and vision features. It provides the practical counterpart to K3's launch analysis, making it valuable for RAG tasks about model usage patterns, pricing, and limitations.
