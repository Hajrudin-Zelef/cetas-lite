---
id: collect-240926-misc/misc/go-3
title: "Go"
domain: opencode
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "LongCat", "Meta", "MiniMax", "Moonshot", "Z.ai", "xAI"]
dates: ["2026-09-30"]
keywords: ["agent", "cost", "deepseek", "glm", "gpt-5.6", "gpt-6", "grok", "grok 4", "kimi", "luna", "muse", "muse spark"]
source: docs/RAG/clean_en/misc/go.md
source_anchor: ""
source_lines: [240, 332]
sha256: 86409c29f253c5179621a451a24e9a57170ce05ad99764a0f48bd6182c525361
---

# Go

For some models, we haven't yet had the opportunity to negotiate a discount or host them at a lower cost, either because the model is new or because its public pricing is already reduced.

For these models, you still get a little more than if you paid the model providers directly; that's why their included monthly usage is lower.

You can also access Go models via the following API endpoints.

| Model | Model ID | Endpoint | AI SDK Package |
|---|---|---|---|
| Grok 4.7 | grok-4.7 | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` |
| Grok 4.6 | grok-4.6 | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` |
| GPT 6 Luna | gpt-6-luna | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` |
| GPT 5.6 Luna | gpt-5.6-luna | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` |
| GLM-5.3-Flash | glm-5.3-flash | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| GLM-5.3 | glm-5.3 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| GLM-5.2 | glm-5.2 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| GLM-5.1 | glm-5.1 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Kimi K3 | kimi-k3 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Kimi K2.7 Code | kimi-k2.7-code | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Kimi K2.6 | kimi-k2.6 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| LongCat-2.0 | longcat-2.0 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| DeepSeek V4.1 Flash | deepseek-v4.1-flash | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| DeepSeek V4 Pro | deepseek-v4-pro | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| DeepSeek V4 Flash | deepseek-v4-flash | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| DeepSeek V4 Flash Vision Exp | deepseek-v4-flash-vision-exp | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| MiMo-V2.6-Flash | mimo-v2.6-flash | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| MiMo-V2.6-Pro | mimo-v2.6-pro | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| MiMo-V2.5 | mimo-v2.5 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| MiMo-V2.5-Pro | mimo-v2.5-pro | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| MiniMax M3 | minimax-m3 | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` |
| MiniMax M2.7 | minimax-m2.7 | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` |
| MiniMax M2.5 | minimax-m2.5 | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` |
| Muse Spark 1.3 Contributor | muse-spark-1.3-contributor | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` |
| Muse Spark 1.2 Contributor | muse-spark-1.2-contributor | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` |
| Qwen3.8 Max | qwen3.8-max | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` |
| Qwen3.8 Flash | qwen3.8-flash | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` |
| Qwen3.7 Max | qwen3.7-max | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` |
| Qwen3.7 Plus | qwen3.7-plus | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` |
| Qwen3.6 Plus | qwen3.6-plus | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` |
| Hy4 preview | hy4-preview | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Hy3 | hy3 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |
| Space Bunny Free | space-bunny-free | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` |

The model ID in your OpenCode configuration uses the format `opencode-go/<model-id>`. For example, for Kimi K3, you would use `opencode-go/kimi-k3` in your configuration.

You can retrieve the complete list of available models and their metadata from:

| Model | Model Training | Data Retention |
|---|---|---|
| Grok 4.7 | Not used | 30 days |
| Grok 4.6 | Not used | 30 days |
| GPT 6 Luna | Not used | 30 days |
| GPT 5.6 Luna | Not used | 30 days |
| GLM-5.3-Flash | Not used | 0 days |
| GLM-5.3 | Not used | 0 days |
| GLM-5.2 | Not used | 0 days |
| GLM-5.1 | Not used | 0 days |
| Kimi K3 | Not used | 0 days |
| Kimi K2.7 Code | Not used | 0 days |
| Kimi K2.6 | Not used | 0 days |
| LongCat-2.0 | Not used | 0 days |
| MiMo-V2.6-Pro | Not used | 0 days |
| MiMo-V2.6-Flash | Not used | 0 days |
| MiMo-V2.5-Pro | Not used | 0 days |
| MiMo-V2.5 | Not used | 0 days |
| Qwen3.8 Max | Not used | 0 days |
| Qwen3.8 Flash | Not used | 0 days |
| Qwen3.7 Max | Not used | 0 days |
| Qwen3.7 Plus | Not used | 0 days |
| Qwen3.6 Plus | Not used | 0 days |
| MiniMax M3 | Not used | 0 days |
| MiniMax M2.7 | Not used | 0 days |
| Muse Spark 1.3 Contributor | Yes | No ZDR |
| Muse Spark 1.2 Contributor | Yes | No ZDR |
| DeepSeek V4.1 Flash | Not used | 0 days |
| DeepSeek V4 Pro | Not used | 0 days |
| DeepSeek V4 Flash | Not used | 0 days |
| DeepSeek V4 Flash Vision Exp | Not used | 0 days |
| Hy4 preview | Not used | 0 days |
| Hy3 | Not used | 0 days |
| Space Bunny Free | Not used | 0 days |

- **Grok 4.7/4.6:** ZDR disables important API features that depend on stored data, including stateful Responses API, Files and Collections, and Batch API. Learn more.
- **GPT 6 Luna / GPT 5.6 Luna:** Abuse monitoring logs are generated for all use of API features and retained for a maximum of 30 days. Learn more.
- **Muse Spark 1.3 Contributor:** Heavily reduced token pricing in exchange for permission to use your prompts and completions to train future Meta models. Availability is limited to regions permitted by Meta's Geographic Usage Policy. Learn more.
- **Muse Spark 1.2 Contributor:** Heavily reduced token pricing in exchange for permission to use your prompts and completions to train future Meta models. Availability is limited to regions permitted by Meta's Geographic Usage Policy. Learn more.
- **DeepSeek:** The ZDR agreement is renewed each month. The current agreement is valid until September 30, 2026.

We created OpenCode Go to:

1. Make AI coding **accessible** to more people with a low-cost subscription.
2. Provide **reliable** access to the best open coding models.
3. Select models that are **tested and evaluated** for use as a coding agent.
4. Have **no exclusive lock-in** by allowing you to use any other provider with OpenCode as well.
