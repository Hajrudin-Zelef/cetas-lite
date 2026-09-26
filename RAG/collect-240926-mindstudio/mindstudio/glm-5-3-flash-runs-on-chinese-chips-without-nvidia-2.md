---
id: collect-240926-mindstudio/mindstudio/glm-5-3-flash-runs-on-chinese-chips-without-nvidia-2
title: "glm-5-3-flash-runs-on-chinese-chips-without-nvidia"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "China", "Nvidia", "OpenAI", "OpenRouter", "Z.ai"]
dates: []
keywords: ["glm", "nvidia", "benchmark", "claude", "cost", "gpt-5.6", "gpus", "inference", "open weights", "opus 4", "parameters"]
source: docs/RAG/clean_en/mindstudio/glm-5-3-flash-runs-on-chinese-chips-without-nvidia.md
source_anchor: ""
source_lines: [52, 68]
sha256: 39951673ab8860256c71d15d797896da0926b72cf2798b4b1d5e02d93af5f70e
---

# glm-5-3-flash-runs-on-chinese-chips-without-nvidia

It’s an open-weights language model from the Chinese AI lab ZAI, using a mixture-of-experts architecture with 320 billion total parameters and 18 billion active parameters per token. It’s positioned as a faster, cheaper “flash” version within the GLM 5.3 model family.

### Did GLM 5.3 Flash really run without any Nvidia chips?

According to reporting cited from semi analysis, ZAI served GLM 5.3 Flash at over 100 trillion tokens per day on a cluster of Chinese-made AI chips with a serving stack and interconnect co-designed for that hardware, without using Nvidia GPUs. This is ZAI’s own reported claim and hasn’t been independently verified in the source material.

### How does GLM 5.3 Flash compare in cost to Claude or GPT models?

Reported per-task costs put GLM 5.3 Flash at around 9 cents, compared to about $3.14 for Claude Opus 4.5 and 95 cents for GPT-5.6 Soul, while scoring only a few points lower than Opus 4.5 on the artificial analysis intelligence index.

### Can I run GLM 5.3 Flash myself?

Since it’s released as open weights, it can be accessed through ZAI’s API, through OpenRouter, or self-hosted by anyone with sufficient hardware, and plugged into any tool supporting an OpenAI-compatible API.

### Why does Chinese chip inference at scale matter beyond this one model?

It suggests that a full domestic alternative to Nvidia-based AI infrastructure, chips, interconnect, and serving software together, can support frontier-adjacent models economically at production scale, not just in isolated benchmark tests.
