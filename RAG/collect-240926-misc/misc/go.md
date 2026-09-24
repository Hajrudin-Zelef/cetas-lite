---
id: collect-240926-misc/misc/go
title: "Go"
domain: opencode
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "LongCat", "Meta", "Microsoft", "MiniMax", "Moonshot", "Z.ai", "xAI"]
dates: ["2026-09-30"]
keywords: ["agent", "agents", "claude", "copilot", "cost", "deepseek", "glm", "gpt-5.6", "gpt-6", "gpu", "grok", "grok 4"]
source: docs/RAG/clean_en/misc/go.md
source_anchor: ""
source_lines: [1, 332]
sha256: 795137dd96bf457f91b0720a88a163d267093583323ccf8e6c63bba4b64f78c6
---

# Go

<!-- source: https://opencode.ai/docs/fr/go/ -->

# Go

Low-cost subscription for open coding models.

OpenCode Go is a low-cost subscription at **$10/month** that gives you reliable access to popular open coding models.

Go works like any other provider in OpenCode. You subscribe to OpenCode Go and get your API key. It's **completely optional** and you don't need to use it to use OpenCode.

It's designed primarily for international users and offers stable global access.

Open models have become really performant. They now achieve performance close to proprietary models for coding tasks. And since many providers can offer them competitively, they're generally much cheaper.

However, getting reliable, low-latency access to these models can be difficult. Providers vary in quality and availability.

To address this, we did several things:

1. We tested a selected group of open models and discussed with their teams the best way to run them.
2. We then worked with a few providers to make sure they were properly served.
3. Finally, we evaluated the performance of the model/provider combination and put together a list that we feel comfortable recommending.

OpenCode Go gives you access to these models for **$10/month**.

OpenCode Go works like any other provider in OpenCode.

1. You log in to **OpenCode Zen**, subscribe to Go and copy your API key.
2. You run the `/connect` command in the TUI, select `OpenCode Go` and paste your API key.
3. Run `/models` in the TUI to see the list of models available via Go.

The current list of models includes:

- **Grok 4.7**
- **Grok 4.6**
- **GLM-5.3-Flash**
- **GLM-5.3**
- **GLM-5.2**
- **GLM-5.1**
- **GPT 6 Luna**
- **GPT 5.6 Luna**
- **Kimi K3**
- **Kimi K2.7 Code**
- **Kimi K2.6**
- **LongCat-2.0**
- **MiMo-V2.6-Flash**
- **MiMo-V2.6-Pro**
- **MiMo-V2.5**
- **MiMo-V2.5-Pro**
- **MiniMax M3**
- **MiniMax M2.7**
- **Muse Spark 1.3 Contributor** (limited regions)
- **Muse Spark 1.2 Contributor** (limited regions)
- **Qwen3.8 Max**
- **Qwen3.8 Flash**
- **Qwen3.7 Max**
- **Qwen3.7 Plus**
- **Qwen3.6 Plus**
- **DeepSeek V4.1 Flash**
- **DeepSeek V4 Pro**
- **DeepSeek V4 Flash**
- **DeepSeek V4 Flash Vision Exp**
- **Hy4 preview**
- **Hy3**
- **Space Bunny Free** (for a limited time)

The list of models may change as we test and add new ones.

OpenCode Go is designed for OpenCode and other coding agents that produce similar types of requests. Traffic is monitored to detect abuse that degrades the experience of other users.

Your client must:

1. Send the usual traffic of a coding agent
2. Identify itself with its own user agent, such as `my-coding-agent/1.0`, rather
than with the generic name of an SDK or HTTP library.
3. Send a stable session ID in `x-opencode-session` for each conversation so that we can optimize routing and
prompt caching.

Besides OpenCode, the proper functioning of the following clients with OpenCode Go has been validated. However, we do not guarantee that they will continue to work in the future.

| Client | Session support | 
|---|---|
| **Hermes** | Builds containing PR #101864 send the header on main and auxiliary OpenCode requests. The fix was merged after v0.21.0; that version alone does not include it. | 
| **Claude Code** | Go recognizes its native session header. No custom header wrapper is needed. | 
| **Codex** | Go recognizes its native session header. Some versions and proxy configurations still omit it; keep the session header when forwarding requests. | 
| **ZCode** | Go recognizes its native session header. Our request regarding `x-opencode-session` remains open, but it is no longer necessary to specifically send this header. | 
| **Pi** | Current builds send session information for OpenCode. Update older installations. | 
| **jcode** | Upgrade to version **v0.81.6 or later**, which includes the session header fix. | 
| **Kilo Code CLI** | Builds containing PR #13752 restore OpenCode session headers. This fix concerns the CLI, not the VS Code extension. See issue #13723. | 

In the versions we examined, these clients do not support sessions or only partially support them. The associated reports track fixes and workarounds.

| Client | Status and tracking | 
|---|---|
| **DeepSeek Harness** | Session information is present for some model paths, but absent for others. We recognize its native header; it remains to be sent from all adapters. Discussion #5495. | 
| **GitHub Copilot Chat** | Automatic session header support is the subject of VS Code issue #334186. | 
| **Kimi Code** | Automatic session header support is the subject of issue #3506. | 
| **MiMo Code** | Issue #2317 has a proposed fix in PR #2327, which has not yet been merged. | 

Usage limits are defined as monthly dollar amounts. The table below shows the monthly limit and token prices for each model.

Each model is subject to the following usage limits: 5 hours — 20% of the monthly limit; weekly — 50%; and monthly — 100%.

For example, if a model has a monthly limit of $60, you can use up to:

- **5-hour limit** → $12 of usage
- **Weekly limit** → $30 of usage
- **Monthly limit** → $60 of usage

Token prices are listed per million tokens.

| Model | Input | Output | Cached Read | Cached Write | Monthly limit | 
|---|---|---|---|---|---|
| GLM-5.3-Flash | $0.15 | $0.50 | $0.03 | - | **$60** | 
| GLM-5.3 | $1.40 | $4.40 | $0.26 | - | **$15** | 
| GLM-5.2 | $1.40 | $4.40 | $0.26 | - | **$60** | 
| GLM-5.1 | $1.40 | $4.40 | $0.26 | - | **$60** | 
| Kimi K3 | $3.00 | $15.00 | $0.30 | - | **$15** | 
| Kimi K2.7 Code | $0.95 | $4.00 | $0.19 | - | **$60** | 
| Kimi K2.6 | $0.95 | $4.00 | $0.16 | - | **$60** | 
| LongCat-2.0 | $0.30 | $1.20 | $0.006 | - | **$60** | 
| MiMo-V2.6-Flash | $0.14 | $0.28 | $0.0028 | - | **$60** | 
| MiMo-V2.6-Pro | $0.435 | $0.87 | $0.003625 | - | **$15** | 
| MiMo-V2.5 | $0.14 | $0.28 | $0.0028 | - | **$60** | 
| MiMo-V2.5-Pro | $0.435 | $0.87 | $0.003625 | - | **$15** | 
| MiniMax M3 | $0.30 | $1.20 | $0.06 | - | **$60** | 
| MiniMax M2.7 | $0.30 | $1.20 | $0.06 | $0.375 | **$60** | 
| MiniMax M2.5 | $0.30 | $1.20 | $0.06 | $0.375 | **$60** | 
| Muse Spark 1.3 Contributor | $0.10 | $0.20 | $0.002 | - | **$60** | 
| Muse Spark 1.2 Contributor | $0.10 | $0.20 | $0.002 | - | **$60** | 
| Qwen3.8 Max | $2.00 | $6.00 | $0.25 | $2.50 | **$15** | 
| Qwen3.8 Flash | $0.15 | $0.47 | $0.016 | $0.20 | **$30** | 
| Qwen3.7 Max | $2.50 | $7.50 | $0.50 | $3.125 | **$30** | 
| Qwen3.7 Plus (â¤ 256K tokens) | $0.40 | $1.60 | $0.04 | $0.50 | **$60** | 
| Qwen3.7 Plus (> 256K tokens) | $1.20 | $4.80 | $0.12 | $1.50 | **$60** | 
| Qwen3.6 Plus (â¤ 256K tokens) | $0.50 | $3.00 | $0.05 | $0.625 | **$60** | 
| Qwen3.6 Plus (> 256K tokens) | $2.00 | $6.00 | $0.20 | $2.50 | **$60** | 
| DeepSeek V4.1 Flash (Off-Peak) | $0.15 | $0.60 | $0.003 | - | ~~$15~~**$60**4x Â· Ends Sept 27 | 
| DeepSeek V4.1 Flash (Peak) | $0.30 | $1.20 | $0.006 | - | ~~$15~~**$60**4x Â· Ends Sept 27 | 
| DeepSeek V4 Pro (Off-Peak) | $0.66 | $1.98 | $0.022 | - | **$15** | 
| DeepSeek V4 Pro (Peak) | $1.32 | $3.96 | $0.044 | - | **$15** | 
| DeepSeek V4 Flash (Off-Peak) | $0.15 | $0.60 | $0.003 | - | **$30** | 
| DeepSeek V4 Flash (Peak) | $0.30 | $1.20 | $0.006 | - | **$30** | 
| DeepSeek V4 Flash Vision Exp (Off-Peak) | $0.15 | $0.60 | $0.003 | - | **$15** | 
| DeepSeek V4 Flash Vision Exp (Peak) | $0.30 | $1.20 | $0.006 | - | **$15** | 
| Hy4 preview | $0.834 | $2.501 | $0.042 | - | **$30** | 
| Hy3 | $0.14 | $0.58 | $0.035 | - | **$60** | 
| Space Bunny Free | Free | Free | Free | - | **Unlimited**for a limited time | 
| Grok 4.7 (â¤ 200K tokens) | $2.00 | $6.00 | $0.50 | - | **$15** | 
| Grok 4.7 (> 200K tokens) | $4.00 | $12.00 | $1.00 | - | **$15** | 
| Grok 4.6 (â¤ 200K tokens) | $2.00 | $6.00 | $0.50 | - | **$15** | 
| Grok 4.6 (> 200K tokens) | $4.00 | $12.00 | $1.00 | - | **$15** | 
| GPT 6 Luna (â¤ 272K tokens) | $0.10 | $0.50 | $0.01 | $0.125 | **$15** | 
| GPT 6 Luna (> 272K tokens) | $0.20 | $0.75 | $0.02 | $0.25 | **$15** | 
| GPT 5.6 Luna (â¤ 272K tokens) | $0.20 | $1.20 | $0.02 | $0.25 | **$15** | 
| GPT 5.6 Luna (> 272K tokens) | $0.40 | $1.80 | $0.04 | $0.50 | **$15** | 

**Space Bunny Free:** Free for a limited time.

**DeepSeek V4.1 Flash / V4 Pro / V4 Flash / V4 Flash Vision Exp:** Peak hours are 01:00-04:00 and 06:00-10:00 UTC, Monday through Friday; all other hours, including weekends, are Off-Peak. Learn more.

**DeepSeek V4 Flash Vision Exp:** Images are converted into tokens based on their dimensions and billed as input tokens along with text tokens. Learn more.

The table below provides an estimate of the number of requests based on typical Go usage habits:

| Model | requests per 5 hours | requests per week | requests per month | 
|---|---|---|---|
| GLM-5.3-Flash | 6,320 | 15,790 | 31,580 | 
| GLM-5.3 | 220 | 540 | 1,080 | 
| GLM-5.2 | 880 | 2,150 | 4,300 | 
| GLM-5.1 | 880 | 2,150 | 4,300 | 
| Kimi K3 | 110 | 250 | 490 | 
| Kimi K2.7 Code | 1,350 | 3,380 | 6,750 | 
| Kimi K2.6 | 1,150 | 2,880 | 5,750 | 
| LongCat-2.0 | 11,400 | 28,600 | 57,200 | 
| MiMo-V2.6-Flash | 30,100 | 75,200 | 150,400 | 
| MiMo-V2.6-Pro | 3,250 | 8,150 | 16,300 | 
| MiMo-V2.5 | 30,100 | 75,200 | 150,400 | 
| MiMo-V2.5-Pro | 3,250 | 8,150 | 16,300 | 
| MiniMax M3 | 3,200 | 8,000 | 16,000 | 
| MiniMax M2.7 | 3,400 | 8,500 | 17,000 | 
| Muse Spark 1.3 Contributor | 45,300 | 113,300 | 226,600 | 
| Muse Spark 1.2 Contributor | 45,300 | 113,300 | 226,600 | 
| Qwen3.8 Max | 160 | 400 | 810 | 
| Qwen3.8 Flash | 5,400 | 13,500 | 27,000 | 
| Qwen3.7 Max | 170 | 420 | 840 | 
| Qwen3.7 Plus | 4,300 | 10,800 | 21,600 | 
| Qwen3.6 Plus | 3,300 | 8,200 | 16,300 | 
| DeepSeek V4.1 Flash 4x Â· Ends Sept 27 | ~~6,500~~**26,000** | ~~16,250~~**65,000** | ~~32,500~~**130,000** | 
| DeepSeek V4 Pro | 1,050 | 2,600 | 5,200 | 
| DeepSeek V4 Flash | 13,000 | 32,500 | 65,000 | 
| DeepSeek V4 Flash Vision Exp | 6,500 | 16,250 | 32,500 | 
| Hy4 preview | 1,350 | 3,380 | 6,770 | 
| Hy3 | 4,300 | 10,750 | 21,500 | 
| Space Bunny Free | Unlimited | Unlimited | Unlimited | 
| Grok 4.7 | 169 | 423 | 845 | 
| Grok 4.6 | 169 | 423 | 845 | 
| GPT 6 Luna | 4,230 | 10,560 | 21,130 | 
| GPT 5.6 Luna | 2,050 | 5,100 | 10,250 | 

The estimates use the following numbers of tokens per request; actual usage varies.

- Grok 4.7/4.6 — 390 input tokens, 32,500 cached, 120 output tokens per request
- GLM-5.3-Flash — 1,000 input tokens, 55,000 cached, 200 output tokens per request
- GLM-5.3/5.2/5.1 — 700 input tokens, 52,000 cached, 150 output tokens per request
- GPT 6 Luna — 1,000 input tokens, 50,000 cached, 220 output tokens per request
- GPT 5.6 Luna — 1,000 input tokens, 50,000 cached, 220 output tokens per request
- Kimi K3 — 1,050 input tokens, 76,500 cached, 300 output tokens per request
- Kimi K2.7/K2.6 — 870 input tokens, 55,000 cached, 200 output tokens per request
- LongCat-2.0 — 920 input tokens, 88,900 cached, 200 output tokens per request
- DeepSeek V4.1 Flash — 410 input tokens, 71,300 cached, 310 output tokens per request
- DeepSeek V4 Pro — 750 input tokens, 82,000 cached, 290 output tokens per request
- DeepSeek V4 Flash — 410 input tokens, 71,300 cached, 310 output tokens per request
- DeepSeek V4 Flash Vision Exp — 410 input tokens, 71,300 cached, 310 output tokens per request
- MiniMax M3 — 510 input tokens, 56,000 cached, 190 output tokens per request
- MiniMax M2.7 — 300 input tokens, 55,000 cached, 125 output tokens per request
- Muse Spark 1.3 Contributor — 620 input tokens, 71,400 cached, 300 output tokens per request
- Muse Spark 1.2 Contributor — 620 input tokens, 71,400 cached, 300 output tokens per request
- Qwen3.8 Max — 420 input tokens, 66,000 cached, 200 output tokens per request
- Qwen3.8 Flash — 600 input tokens, 58,000 cached, 200 output tokens per request
- Qwen3.7 Max — 420 input tokens, 66,000 cached, 200 output tokens per request
- Qwen3.7 Plus — 500 input tokens, 57,000 cached, 190 output tokens per request
- Qwen3.6 Plus — 500 input tokens, 57,000 cached, 190 output tokens per request
- Hy4 preview — 830 input tokens, 71,500 cached, 295 output tokens per request
- Hy3 — 830 input tokens, 71,500 cached, 295 output tokens per request
- MiMo-V2.6-Flash — 830 input tokens, 71,500 cached, 295 output tokens per request
- MiMo-V2.6-Pro — 790 input tokens, 86,000 cached, 305 output tokens per request
- MiMo-V2.5 — 830 input tokens, 71,500 cached, 295 output tokens per request
- MiMo-V2.5-Pro — 790 input tokens, 86,000 cached, 305 output tokens per request

You can track your current usage in the **console**.

Usage limits may change as we learn from early usage and feedback.

If you also have credits in your Zen balance, you can enable the **Use balance** option in the console. When enabled, Go will fall back to your Zen balance after you reach your usage limits instead of blocking requests.

With Go, you pay $10/month, and the included monthly usage varies by model.

For most models, we achieve this through volume discounts and reserved GPU capacity. We then pass those savings on to you in the form of higher monthly usage.

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
