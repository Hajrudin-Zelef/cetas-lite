---
id: collect-261001-general-networking/general-networking/state-of-open-models-summer-2026-observations-2
title: "state-of-open-models-summer-2026-observations"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "China", "DeepSeek", "Google", "Hugging Face", "Meta", "Moonshot", "Unsloth", "Z.ai"]
dates: []
keywords: ["agent", "agentic", "agents", "apache", "attention", "claude", "consumer", "deepseek", "disclosure", "fine-tuning", "gguf", "glm"]
source: docs/RAG/collect-261001-general-networking/state-of-open-models-summer-2026-observations.md
source_anchor: ""
source_lines: [55, 104]
sha256: 71769b0cc1fb395a345c80ae94ec5c4d9e5e9ee44e5216e4325dc221eaf58c41
---

# state-of-open-models-summer-2026-observations

By this measure, Qwen has become one of the largest foundations in the open model ecosystem. Qwen-based models now account for **151,448 derivatives** on the Hub, 2.6× Meta’s total footprint and 4.7× the Llama repositories specifically. Google follows with 82,506 derivatives. The third-largest source is Unsloth, a community account publishing quantized and fine-tuning-ready builds, many of which further extend the Qwen ecosystem.

Qwen derivatives have increased at roughly **180–210 new repositories per day** throughout the first seven months of 2026, showing that adoption is not driven only by individual launches. Qwen has become part of the default workflow for developers deciding what models to fine-tune and deploy.

Several factors contributed to this position. **First, consistency.** Qwen has maintained a regular release cadence, continuously updating its model family rather than relying on occasional flagship releases. **Second, coverage.** It publishes models across a wide range of sizes and use cases, allowing developers to stay within the same ecosystem whether they need a small local model or a larger deployment model. **Third, openness.** Apache 2.0 licensing reduces friction for modification, redistribution, and commercial use.

These factors reinforce each other. A broad model family attracts more developers; more developers create more derivatives; and those derivatives make the ecosystem more attractive to future users.

This position was built largely by the community. The **151,448 derivatives** represent downstream work created by other developers, not releases produced by Qwen itself. Even among the **28,531 GGUF conversions** of Qwen models on the Hub, Qwen published only 54.

Among models that declare a parameter count, those under 1B take 83% of all-time downloads and everything above 100B takes 1%. Restricting to downloads accumulated in 2026 changes nothing: 3% of the volume goes to models above 70B. This is the March finding that has held up most cleanly, for the same reason as before, small models are the only ones that run on the hardware most developers actually have.

So how does a trillion-parameter model reach anyone at all? Through llama.cpp.

In February the ggml team joined Hugging Face, with the project remaining fully open-source, community-governed and in the same technical direction. What changed is that the most important project in local inference now has durable resources behind it.

The ceiling moved with llama.cpp. The July snapshot carries GGUF builds of DeepSeek-V4-Flash at roughly 284B parameters and Kimi-K3 at roughly 2.8 trillion. Local inference used to mean an 8B model on a laptop. It now means a trillion-parameter mixture-of-experts spread across a few consumer machines, which is the alternative route the frontier did not have a year ago, and the reason a frontier-first release strategy is viable at all.

And that route runs on Qwen: 39.6 million GGUF downloads a month, nearly twice Gemma's 20.8 million and more than five times Llama's 7.5 million. The Llama gap is not a supply problem, Llama-derived GGUF repositories slightly outnumber Qwen's. Same shelf space, a fifth of the traffic.

Model repositories grew 21.5% over these seven months. Several things around them grew several times faster.

Repositories declaring the gguf library rose 464%, lerobot 194% and Apple's mlx148%, against 16% for transformers and peft and 21% for diffusers. The modelling core is growing at roughly the platform average. The layer that decides where a model can physically run local inference formats, Apple silicon, robot control stacks, is growing three to seven times faster than that.

Across the ten largest model families, the labs behind these models publish very few official GGUF conversions. Yet GGUF versions are often the ones used by developers running models locally. Providing an official conversion at release, documenting quantization choices, and signing the artifacts would require limited additional effort. Rather than maintaining this workflow internally, labs could collaborate with existing ecosystem contributors such as Unsloth. Doing so would narrow the gap between the weights tested by model creators and the versions adopted by the broader community.

We could not have written this section in March, because the instrument did not exist. The agent-usage dataset, published in July, records the agent/<name> token that coding agents send when they call the Hub through huggingface_hub or the hf CLI — searching for models, pushing datasets, running Jobs, creating Spaces. For the first time we can see how much agent traffic the Hub receives and which harnesses it comes from.

Claude Code led July with 44.4%, but a single month conceals the real finding: it held 67.8% in April and 64% in May, while Codex climbed steadily from 10.4% to 20.8%. This is a market with no incumbent, where one release or one changed default can move half the traffic in a month.

The second finding is the unregistered row. Nearly a quarter of agent-tagged traffic in July came from harnesses not yet named in the dataset, and in May that figure was 59.8%. Between April and July more than a dozen new client identifiers appeared. New entrants are arriving faster than any registry can name them — which is itself the finding.

We spent much of the year building for this reader rather than only for human browsers. Papers began serving machine-readable Markdown in March. April brought agent traces as a first-class dataset type and an agents.md endpoint on every Gradio Space, so an agent can read a Space's API and call it directly. July brought the hf_fs tool on our MCP server, exposing repositories, storage, docs and papers through a single interface in just over a thousand tokens, alongside attachable sandboxes for secure execution. The same consolidation happened at the protocol layer, with MCP moving into the Linux Foundation's Agentic AI Foundation.

Then, in July, an agent stopped being a reader and became an intruder. What appears to be the first documented case of an autonomous agent running a sustained intrusion on its own initiative happened to us. While our team tried to use frontier closed models to analyze the captured attack code, their safety guardrails declined the work. The analysis was completed in the end on a quantized open model GLM-5.2 running on our own infrastructure. We published a disclosure and a full technical timeline.

Compared to the spring report, the geographical rebalancing of power continues to accelerate. While U.S. open source models continue to be competitive, the race between several Chinese frontier model labs draws strong attention. Many likes on these frontier models point to what excites the community the most, and growth opportunity for companies leveraging the attention for valuations.

However, the AI race is not only sprints, but also a marathon; tools like llama.cpp helps deploying the big models locally, but a broad model family and its adoption is still the key, to build a positive feedback loop between developers, publisher and future users. Models to be embedded in the infrastructure and being part of the ecosystem, may lead to a commercially sound exit at the end of the tunnel.

In the end, with agents being the number 1 user on HF Hub for the first time, the next report may look very different.

In AI, a few months can reshape the ecosystem.

This analysis is based on activity observed on the Hugging Face Hub during the first seven months of 2026.

The metrics used in this report, including downloads, likes, derivatives, and model releases, represent different aspects of ecosystem activity. They should not be interpreted as direct measures of model quality, commercial adoption, or overall market share.

Downloads indicate usage within the Hub ecosystem, but they do not capture API usage, private deployments, or models distributed through other channels.

