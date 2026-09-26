---
id: collect-240926-misc/misc/zen-4
title: "Zen"
domain: opencode
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "United States", "Z.ai"]
dates: ["2026-02-06", "2026-02-16", "2026-03-06", "2026-03-09", "2026-03-15", "2026-05-14", "2026-06-15", "2026-07-23", "2026-08-05"]
keywords: ["agent", "agents", "benchmark", "claude", "cost", "gemini", "glm", "kimi", "muse", "muse spark", "nvidia", "opus 4"]
source: docs/RAG/clean_en/misc/zen.md
source_anchor: ""
source_lines: [246, 309]
sha256: e75a5e56627a7d041794b944f78c801c8128dadf2816c32a32457207fe84e08b
---

# Zen

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
