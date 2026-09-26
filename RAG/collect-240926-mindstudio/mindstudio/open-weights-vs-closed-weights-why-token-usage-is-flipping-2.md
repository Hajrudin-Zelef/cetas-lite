---
id: collect-240926-mindstudio/mindstudio/open-weights-vs-closed-weights-why-token-usage-is-flipping-2
title: "open-weights-vs-closed-weights-why-token-usage-is-flipping"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "China", "DeepSeek", "OpenAI", "United States"]
dates: []
keywords: ["open-weight", "claude", "cost", "cost per token", "deepseek", "open weights", "parameters", "pricing", "revenue"]
source: docs/RAG/clean_en/mindstudio/open-weights-vs-closed-weights-why-token-usage-is-flipping.md
source_anchor: ""
source_lines: [55, 77]
sha256: b5122bb7a673767d7f70dcf812e3cb5fdf424ebbe4cbe04aa85ebeb6877c496a
---

# open-weights-vs-closed-weights-why-token-usage-is-flipping

One useful framing, attributed to MIT economist Christian Catalini, splits the market into three tiers. At the bottom are cheap, generalist open-weight models, commodity-priced, carrying the most volume but a small share of total spend. In the middle are open-weight “state of the art specialist” models, fine-tuned for enterprise use cases, which carry meaningful spend despite being open. At the top sit closed, frontier generalist models from labs like OpenAI and Anthropic, a small share of total volume but the majority of revenue. That structure lines up with what the token-share and spend-share data already show: broad usage spreading toward open models while premium spend stays locked with a handful of closed-model providers.

## Frequently Asked Questions

### What’s the difference between open-weight and closed-weight AI models?

Open-weight models release their trained parameters publicly, so anyone can download, run, and fine-tune them on their own infrastructure. Closed-weight models like GPT and Claude are only accessible through a provider’s API, with the underlying weights kept private.

### Why is DeepSeek gaining token share on Anthropic?

DeepSeek’s models are priced far lower per token than Anthropic’s, and being open weight, they can run on many different providers or be self-hosted, making them attractive for high-volume, cost-sensitive workloads even as Anthropic retains most premium, high-value use cases.

### Does more token usage mean more revenue for open-weight providers?

No. Token volume and revenue are decoupled in this market. Open-weight models handle a growing share of total tokens, but closed frontier models still capture the large majority of total dollars spent, because premium tasks command premium prices.

### Why are US companies building products on Chinese open-weight models?

Cost is part of it, but privacy and control matter just as much. Open weights let companies fine-tune models on proprietary data without sending sensitive information to a third-party API, and they avoid depending entirely on a single closed-model vendor.

### Is cost per token a reliable way to compare AI models?

Not on its own. Some models require more tokens to complete the same task, so a lower per-token price doesn’t always translate into a lower total cost. Comparing cost per completed task gives a more accurate picture than raw token pricing.
