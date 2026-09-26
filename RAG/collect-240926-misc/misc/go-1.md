---
id: collect-240926-misc/misc/go-1
title: "Go"
domain: opencode
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "LongCat", "Meta", "Microsoft", "MiniMax", "Moonshot", "Z.ai", "xAI"]
dates: []
keywords: ["agent", "agents", "claude", "copilot", "cost", "deepseek", "glm", "grok", "grok 4", "kimi", "latency", "luna"]
source: docs/RAG/clean_en/misc/go.md
source_anchor: ""
source_lines: [1, 110]
sha256: f0deac98b4a421ac67c7c27c8713418e0179b3b7f6d77601ceb23bcf88e93ea6
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

