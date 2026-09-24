---
id: collect-240926-misc/misc/zen
title: "Zen"
domain: opencode
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "United States", "Z.ai", "xAI"]
dates: ["2026-02-06", "2026-02-16", "2026-03-06", "2026-03-09", "2026-03-15", "2026-05-14", "2026-06-15", "2026-07-23", "2026-08-05"]
keywords: ["agent", "agents", "astra", "benchmark", "claude", "cost", "deepseek", "fable 5", "gemini", "gemini 3.8", "glm", "gpt-5.6"]
source: docs/RAG/clean_en/misc/zen.md
source_anchor: ""
source_lines: [1, 309]
sha256: 02fa0a5ae1edc280191d2d7013858033aa8e1528764d8e8cf53dc972681b9c39
---

# Zen

<!-- source: https://opencode.ai/docs/fr/zen/ -->

# Zen

Organized list of models provided by OpenCode.

OpenCode Zen is a list of tested and verified models provided by the OpenCode team.

Zen works like any other provider in OpenCode. You log in to OpenCode Zen and get your API key. It's **completely optional** and you don't need to use it to use OpenCode.

There are a large number of models, but only a few work well as coding agents. Additionally, most providers are configured very differently; so you get very variable performance and quality.

So, if you are using a model via something like OpenRouter, you can never be sure of getting the best version of the model you want.

To fix this, we did several things:

1. We tested a selected group of models and discussed with their teams the best way to make them work.
2. We then worked with a few providers to ensure they were being served correctly.
3. Finally, we compared the model/provider combination and established a list that we confidently recommend.

OpenCode Zen is an AI gateway that gives you access to these models.

OpenCode Zen works like any other provider in OpenCode.

1. You log in to **OpenCode Zen** , add your billing information and copy your API key.
2. You run the `/connect` command in the TUI, select OpenCode Zen and paste your API key.
3. Run `/models` in the TUI to see the list of models we recommend.

Billing is per request and you can add credits to your account.

You can also access our models via the following API endpoints.

| Model | Model ID | Endpoint | AI SDK Package |
|---|---|---|---|---|
| GPT 6 Astra | gpt-6-astra | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 6 Sol | gpt-6-sol | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 6 Luna | gpt-6-luna | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.6 Sol | gpt-5.6-sol | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.6 Terra | gpt-5.6-terra | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.6 Luna | gpt-5.6-luna | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.5 | gpt-5.5 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.5 Pro | gpt-5.5-pro | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.4 | gpt-5.4 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.4 Pro | gpt-5.4-pro | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.4 Mini | gpt-5.4-mini | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.4 Nano | gpt-5.4-nano | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.3 Codex | gpt-5.3-codex | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.3 Codex Spark | gpt-5.3-codex-spark | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.2 | gpt-5.2 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.2 Codex | gpt-5.2-codex | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.1 | gpt-5.1 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.1 Codex | gpt-5.1-codex | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.1 Codex Max | gpt-5.1-codex-max | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5.1 Codex Mini | gpt-5.1-codex-mini | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5 | gpt-5 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5 Codex | gpt-5-codex | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| GPT 5 Nano | gpt-5-nano | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| Claude Fable 5.1 | claude-fable-5-1 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Claude Fable 5 | claude-fable-5 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Claude Opus 5.5 | claude-opus-5-5 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Claude Opus 5 | claude-opus-5 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Claude Opus 4.8 | claude-opus-4-8 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Claude Opus 4.7 | claude-opus-4-7 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Claude Opus 4.6 | claude-opus-4-6 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Claude Opus 4.5 | claude-opus-4-5 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Claude Sonnet 5 | claude-sonnet-5 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Claude Sonnet 4.6 | claude-sonnet-4-6 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Claude Sonnet 4.5 | claude-sonnet-4-5 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Claude Haiku 4.5 | claude-haiku-4-5 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Gemini 3.8 Flash | gemini-3.8-flash | `https://opencode.ai/zen/v1/models/gemini-3.8-flash` | `@ai-sdk/google` |
| Gemini 3.7 Flash | gemini-3.7-flash | `https://opencode.ai/zen/v1/models/gemini-3.7-flash` | `@ai-sdk/google` |
| Gemini 3.6 Flash | gemini-3.6-flash | `https://opencode.ai/zen/v1/models/gemini-3.6-flash` | `@ai-sdk/google` |
| Gemini 3.5 Flash | gemini-3.5-flash | `https://opencode.ai/zen/v1/models/gemini-3.5-flash` | `@ai-sdk/google` |
| Gemini 3.5 Flash Lite | gemini-3.5-flash-lite | `https://opencode.ai/zen/v1/models/gemini-3.5-flash-lite` | `@ai-sdk/google` |
| Gemini 3.1 Pro | gemini-3.1-pro | `https://opencode.ai/zen/v1/models/gemini-3.1-pro` | `@ai-sdk/google` |
| Gemini 3 Flash | gemini-3-flash | `https://opencode.ai/zen/v1/models/gemini-3-flash` | `@ai-sdk/google` |
| Grok 4.7 | grok-4.7 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| Grok 4.6 | grok-4.6 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| Grok 4.5 | grok-4.5 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| Grok Build 0.1 | grok-build-0.1 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| Muse Spark 1.3 | muse-spark-1.3 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| Muse Spark 1.2 | muse-spark-1.2 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |
| Qwen3.8 Flash | qwen3.8-flash | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Qwen3.7 Max | qwen3.7-max | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Qwen3.7 Plus | qwen3.7-plus | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Qwen3.6 Plus | qwen3.6-plus | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| Qwen3.5 Plus | qwen3.5-plus | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` |
| DeepSeek V4.1 Flash | deepseek-v4.1-flash | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| DeepSeek V4 Pro | deepseek-v4-pro | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| DeepSeek V4 Flash | deepseek-v4-flash | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| DeepSeek V4 Flash Vision Exp | deepseek-v4-flash-vision-exp | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| MiniMax M3 | minimax-m3 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| MiniMax M2.7 | minimax-m2.7 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| MiniMax M2.5 | minimax-m2.5 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| GLM 5.3 Flash | glm-5.3-flash | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| GLM 5.3 | glm-5.3 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| GLM 5.2 | glm-5.2 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| GLM 5.1 | glm-5.1 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| GLM 5 | glm-5 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Kimi K2.5 | kimi-k2.5 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Kimi K2.6 | kimi-k2.6 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Kimi K2.7 Code | kimi-k2.7-code | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Kimi K3 | kimi-k3 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Jev 1.13 | jev-1.13 | `https://opencode.ai/zen/v1/systemone` | - |
| Jev 1.13 Free | jev-1.13-free | `https://opencode.ai/zen/v1/systemone` | - |
| Big Pickle | big-pickle | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Space Bunny Free | space-bunny-free | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| MiMo-V2.6-Flash Free | mimo-v2.6-flash-free | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| MiMo-V2.5 Free | mimo-v2.5-free | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Ling 3.0 Flash Fin Free | ling-3.0-flash-fin-free | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Nemotron 3 Ultra Free | nemotron-3-ultra-free | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Nemotron 3.5 Lightning Free | nemotron-3.5-lightning-free | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Muse Spark 1.3 Contributor Free | muse-spark-1.3-contributor-free | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` |

The model id in your OpenCode configuration uses the format `opencode/<model-id>`. For example, for GPT 5.5, you would use `opencode/gpt-5.5` in your configuration.

You can retrieve the complete list of available models and their metadata from:

Jev is a System One model from TypeSafe AI designed to quickly make structured decisions. Instead of generating text, it evaluates a `state` from typed questions and returns values and probabilities directly usable by your code. It supports yes/no questions (`noul`), multiple choice (`choice`), and those based on a rating grid (`score`).

Use your OpenCode Zen API key with the endpoint `https://opencode.ai/zen/v1/systemone`. This example checks whether a support request is urgent:

You can ask multiple questions in a single request. Jev evaluates them in parallel and returns each response under the corresponding question ID:

Use `jev-1.13-free` instead of `jev-1.13` to use the free model offered for a limited time. See the TypeSafe AI documentation to learn more about question types and response fields.

We support a pay-as-you-go payment model. Below are the prices **per 1M tokens**.

| Model | Input | Output | Cached Read | Cached Write |
|---|---|---|---|---|
| Big Pickle | Free | Free | Free | - |
| Space Bunny Free | Free | Free | Free | - |
| MiMo-V2.6-Flash Free | Free | Free | Free | - |
| MiMo-V2.5 Free | Free | Free | Free | - |
| Ling 3.0 Flash Fin Free | Free | Free | Free | - |
| Nemotron 3 Ultra Free | Free | Free | Free | - |
| Nemotron 3.5 Lightning Free | Free | Free | Free | - |
| Muse Spark 1.3 Contributor Free | Free | Free | Free | - |
| Jev 1.13 Free | Free | Free | - | - |
| Jev 1.13 | $0.042 | Free | - | - |
| MiniMax M3 | $0.30 | $1.20 | $0.06 | - |
| MiniMax M2.7 | $0.30 | $1.20 | $0.06 | - |
| MiniMax M2.5 | $0.30 | $1.20 | $0.06 | - |
| GLM 5.3 Flash | $0.15 | $0.50 | $0.03 | - |
| GLM 5.3 | $1.40 | $4.40 | $0.26 | - |
| GLM 5.2 | $1.40 | $4.40 | $0.26 | - |
| GLM 5.1 | $1.40 | $4.40 | $0.26 | - |
| GLM 5 | $1.00 | $3.20 | $0.20 | - |
| Kimi K2.7 Code | $0.95 | $4.00 | $0.19 | - |
| Kimi K3 | $3.00 | $15.00 | $0.30 | - |
| Kimi K2.6 | $0.95 | $4.00 | $0.16 | - |
| Kimi K2.5 | $0.60 | $3.00 | $0.10 | - |
| Qwen3.8 Flash | $0.15 | $0.47 | $0.016 | $0.20 |
| Qwen3.7 Max | $2.50 | $7.50 | $0.50 | $3.125 |
| Qwen3.7 Plus | $0.40 | $1.60 | $0.04 | $0.50 |
| Qwen3.6 Plus | $0.50 | $3.00 | $0.05 | $0.625 |
| Qwen3.5 Plus | $0.20 | $1.20 | $0.02 | $0.25 |
| DeepSeek V4.1 Flash | $0.30 | $1.20 | $0.006 | - |
| DeepSeek V4 Pro | $1.74 | $3.48 | $0.145 | - |
| DeepSeek V4 Flash | $0.14 | $0.28 | $0.028 | - |
| DeepSeek V4 Flash Vision Exp | $0.14 | $0.28 | $0.028 | - |
| Claude Fable 5.1 | $10.00 | $50.00 | $0.25 | $12.50 |
| Claude Fable 5 | $10.00 | $50.00 | $1.00 | $12.50 |
| Claude Opus 5.5 | $4.00 | $20.00 | $0.20 | $5.00 |
| Claude Opus 5 | $5.00 | $25.00 | $0.50 | $6.25 |
| Claude Opus 4.8 | $5.00 | $25.00 | $0.50 | $6.25 |
| Claude Opus 4.7 | $5.00 | $25.00 | $0.50 | $6.25 |
| Claude Opus 4.6 | $5.00 | $25.00 | $0.50 | $6.25 |
| Claude Opus 4.5 | $5.00 | $25.00 | $0.50 | $6.25 |
| Claude Sonnet 5 | $2.00 | $10.00 | $0.20 | $2.50 |
| Claude Sonnet 4.6 | $3.00 | $15.00 | $0.30 | $3.75 |
| Claude Sonnet 4.5 (≤ 200K tokens) | $3.00 | $15.00 | $0.30 | $3.75 |
| Claude Sonnet 4.5 (> 200K tokens) | $6.00 | $22.50 | $0.60 | $7.50 |
| Claude Haiku 4.5 | $1.00 | $5.00 | $0.10 | $1.25 |
| Gemini 3.8 Flash | $1.50 | $7.50 | $0.15 | - |
| Gemini 3.7 Flash | $1.50 | $7.50 | $0.15 | - |
| Gemini 3.6 Flash | $1.50 | $7.50 | $0.15 | - |
| Gemini 3.5 Flash | $1.50 | $9.00 | $0.15 | - |
| Gemini 3.5 Flash Lite | $0.30 | $2.50 | $0.03 | - |
| Gemini 3.1 Pro (≤ 200K tokens) | $2.00 | $12.00 | $0.20 | - |
| Gemini 3.1 Pro (> 200K tokens) | $4.00 | $18.00 | $0.40 | - |
| Gemini 3 Flash | $0.50 | $3.00 | $0.05 | - |
| Grok 4.7 (≤ 200K tokens) | $2.00 | $6.00 | $0.50 | - |
| Grok 4.7 (> 200K tokens) | $4.00 | $12.00 | $1.00 | - |
| Grok 4.6 (≤ 200K tokens) | $2.00 | $6.00 | $0.50 | - |
| Grok 4.6 (> 200K tokens) | $4.00 | $12.00 | $1.00 | - |
| Grok 4.5 (≤ 200K tokens) | $2.00 | $6.00 | $0.30 | - |
| Grok 4.5 (> 200K tokens) | $4.00 | $12.00 | $0.60 | - |
| Grok Build 0.1 | $1.00 | $2.00 | $0.20 | - |
| Muse Spark 1.3 | $1.25 | $4.25 | $0.15 | - |
| Muse Spark 1.2 | $1.25 | $4.25 | $0.15 | - |
| GPT 6 Astra (≤ 272K tokens) | $10.00 | $50.00 | $1.00 | $12.50 |
| GPT 6 Astra (> 272K tokens) | $20.00 | $75.00 | $2.00 | $25.00 |
| GPT 6 Sol (≤ 272K tokens) | $2.00 | $10.00 | $0.20 | $2.50 |
| GPT 6 Sol (> 272K tokens) | $4.00 | $15.00 | $0.40 | $5.00 |
| GPT 6 Luna (≤ 272K tokens) | $0.10 | $0.50 | $0.01 | $0.125 |
| GPT 6 Luna (> 272K tokens) | $0.20 | $0.75 | $0.02 | $0.25 |
| GPT 5.6 Sol (≤ 272K tokens) | $4.00 | $20.00 | $0.40 | $5.00 |
| GPT 5.6 Sol (> 272K tokens) | $8.00 | $30.00 | $0.80 | $10.00 |
| GPT 5.6 Terra (≤ 272K tokens) | $2.00 | $12.00 | $0.20 | $2.50 |
| GPT 5.6 Terra (> 272K tokens) | $4.00 | $18.00 | $0.40 | $5.00 |
| GPT 5.6 Luna (≤ 272K tokens) | $0.20 | $1.20 | $0.02 | $0.25 |
| GPT 5.6 Luna (> 272K tokens) | $0.40 | $1.80 | $0.04 | $0.50 |
| GPT 5.5 (≤ 272K tokens) | $5.00 | $30.00 | $0.50 | - |
| GPT 5.5 (> 272K tokens) | $10.00 | $45.00 | $1.00 | - |
| GPT 5.5 Pro | $30.00 | $180.00 | $30.00 | - |
| GPT 5.4 (≤ 272K tokens) | $2.50 | $15.00 | $0.25 | - |
| GPT 5.4 (> 272K tokens) | $5.00 | $22.50 | $0.50 | - |
| GPT 5.4 Pro | $30.00 | $180.00 | $30.00 | - |
| GPT 5.4 Mini | $0.75 | $4.50 | $0.075 | - |
| GPT 5.4 Nano | $0.20 | $1.25 | $0.02 | - |
| GPT 5.3 Codex Spark | $1.75 | $14.00 | $0.175 | - |
| GPT 5.3 Codex | $1.75 | $14.00 | $0.175 | - |
| GPT 5.2 | $1.75 | $14.00 | $0.175 | - |
| GPT 5.2 Codex | $1.75 | $14.00 | $0.175 | - |
| GPT 5.1 | $1.07 | $8.50 | $0.107 | - |
| GPT 5.1 Codex | $1.07 | $8.50 | $0.107 | - |
| GPT 5.1 Codex Max | $1.25 | $10.00 | $0.125 | - |
| GPT 5.1 Codex Mini | $0.25 | $2.00 | $0.025 | - |
| GPT 5 | $1.07 | $8.50 | $0.107 | - |
| GPT 5 Codex | $1.07 | $8.50 | $0.107 | - |
| GPT 5 Nano | $0.05 | $0.40 | $0.005 | - |

**DeepSeek V4 Flash Vision Exp:** Images are converted into tokens based on their dimensions and billed as input tokens along with text tokens. Learn more.

You may notice low-cost models, such as Haiku, Nano, or Flash, in your usage history. OpenCode uses these models to generate session titles.

The free models:

- MiMo-V2.6-Flash Free is available on OpenCode for a limited time. The team is using this period to gather feedback and improve the model.
- MiMo-V2.5 Free is available on OpenCode for a limited time. The team is using this period to gather feedback and improve the model.
- Ling 3.0 Flash Fin Free is available on OpenCode for a limited time. The team is using this period to gather feedback and improve the model.
- Nemotron 3 Ultra Free is available on OpenCode for a limited time. The team is using this period to gather feedback and improve the model.
- Nemotron 3.5 Lightning Free is available on OpenCode for a limited time. The team is using this period to gather feedback and improve the model.
- Big Pickle is a free stealth model on OpenCode for a limited time. The team is using this period to gather feedback and improve the model.
- Space Bunny Free is a free stealth model on OpenCode for a limited time. Its provider applies a zero-retention policy and does not use your data to train models.
- Muse Spark 1.3 Contributor Free is available on OpenCode for a limited time. The team is using this period to gather feedback and improve the model.
- Jev 1.13 Free is available on OpenCode for a limited time.

Contact us if you have any questions.

If your balance drops below $5, Zen will automatically top up $20.

You can change the auto top-up amount. You can also disable auto top-up entirely.

You can also set a monthly usage limit for the entire workspace and for each member of your team.

For example, if you set a monthly usage limit of $20, Zen will not use more than $20 in a month. But if auto top-up is enabled, Zen may end up charging you more than $20 if your balance drops below $5.

| Model | Deprecation date | 
|---|---|
| GPT 5.2 Codex | July 23, 2026 | 
| GPT 5.1 Codex | July 23, 2026 | 
| GPT 5.1 Codex Max | July 23, 2026 | 
| GPT 5.1 Codex Mini | July 23, 2026 | 
| GPT 5 Codex | July 23, 2026 | 
| Claude Opus 4.1 | August 5, 2026 | 
| Claude Sonnet 4 | June 15, 2026 | 
| Claude Haiku 3.5 | February 16, 2026 | 
| Gemini 3 Pro | March 9, 2026 | 
| MiniMax M2.5 | August 5, 2026 | 
| MiniMax M2.1 | March 15, 2026 | 
| GLM 5 | May 14, 2026 | 
| GLM 4.7 | March 15, 2026 | 
| GLM 4.6 | March 15, 2026 | 
| Kimi K2.5 | August 5, 2026 | 
| Kimi K2 Thinking | March 6, 2026 | 
| Kimi K2 | March 6, 2026 | 
| Qwen3 Coder 480B | February 6, 2026 | 

All of our models are hosted in the US. Our providers follow a zero-retention policy and do not use your data for model training, with the following exceptions:

- Big Pickle: During its free period, the data collected may be used to improve the model.
- MiMo-V2.6-Flash Free: During its free period, the data collected may be used to improve the model.
- MiMo-V2.5 Free: During its free period, the data collected may be used to improve the model.
- Ling 3.0 Flash Fin Free: During its free period, the data collected may be used to improve the model.
- Nemotron 3 Ultra Free (free NVIDIA endpoints): For trial use only — do not send personal or confidential data. Your usage is logged for security purposes and to improve NVIDIA's products and services. Session data logged for improvement purposes is not linked to your identity or any persistent identifier. For more information about our data processing practices, see our Privacy Policy. By interacting with this endpoint, you consent to our collection, recording, and use of this information as well as the NVIDIA API Trial Terms of Service.
- Nemotron 3.5 Lightning Free (free NVIDIA endpoints): For trial use only — do not send personal or confidential data. Your usage is logged for security purposes and to improve NVIDIA's products and services. Session data logged for improvement purposes is not linked to your identity or any persistent identifier. For more information about our data processing practices, see our Privacy Policy. By interacting with this endpoint, you consent to our collection, recording, and use of this information as well as the NVIDIA API Trial Terms of Service.
- OpenAI APIs: Requests are retained for 30 days in accordance with OpenAI's Data Policies.
- Anthropic APIs: Requests are retained for 30 days in accordance with Anthropic's Data Policies.
- Muse Spark 1.3 Contributor Free: Heavily discounted token pricing in exchange for permission to use your prompts and completions to train future Meta models. Learn more.

Zen also works very well for teams. You can invite teammates, assign roles, select the models your team uses, and more.

Managing your workspace is currently free for teams as part of the beta version. We will share more details on pricing soon.

You can invite teammates to your workspace and assign roles:

- **Admin**: Manage models, members, API keys, and billing
- **Member**: Manage only their own API keys

Administrators can also set monthly spending limits for each member to keep costs under control.

Administrators can enable or disable specific models for the workspace. Requests made to a disabled model will return an error.

This is useful if you want to disable the use of a model that collects data.

You can use your own OpenAI or Anthropic API keys while accessing other models in Zen.

When you use your own keys, tokens are billed directly by the provider, not by Zen.

For example, your organization may already have a key for OpenAI or Anthropic and you want to use it instead of the one provided by Zen.

We created OpenCode Zen to:

1. **Benchmark** the best models/providers for coding agents.
2. Have access to the **highest quality** options without degrading performance or switching to cheaper providers.
3. Pass on any **price drops** by selling at cost; the only margin serves to cover our processing fees.
4. Have **no lock-in** by allowing you to use it with any other coding agent. And always allow you to use any other provider with OpenCode as well.
