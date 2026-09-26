---
id: collect-240926-mindstudio/mindstudio/deepseek-v4-flash-the-cheapest-frontier-level-open-model-yet-2
title: "deepseek-v4-flash-the-cheapest-frontier-level-open-model-yet"
domain: mindstudio
role: reference
task: reference
actors: ["DeepSeek", "OpenAI", "Z.ai"]
dates: []
keywords: ["deepseek", "agi", "benchmark", "benchmarks", "cost", "glm", "multimodal", "parameters", "pricing", "training"]
source: docs/RAG/clean_en/mindstudio/deepseek-v4-flash-the-cheapest-frontier-level-open-model-yet.md
source_anchor: ""
source_lines: [70, 82]
sha256: d1978edecaabf14a948d4725f08df004ec6d7d38b0f9586b3d6af6e9014faaac
---

# deepseek-v4-flash-the-cheapest-frontier-level-open-model-yet

On the benchmarks DeepSeek published, yes, on nearly every one, despite GLM 5.2 having roughly three times the parameters. This illustrates how much performance post-training can add without increasing model size.

### Why do harness-dependent scores matter for comparing models?

The harness (the scaffolding managing tool calls, retries, and task execution around a model) can swing benchmark scores dramatically on its own. OpenAI’s ARC-AGI comparison showed scores roughly tripling with an optimized harness versus the standard one, so cross-model comparisons using different harnesses aren’t fully apples-to-apples.

### Is DeepSeek V4 Flash multimodal?

No. It’s text-only with no image or vision input support. Multimodal alternatives exist for use cases that require processing images or audio, though they may score lower on pure coding benchmarks.

### What does DeepSeek V4 Flash cost through the API?

Pricing runs around 2 cents per million input tokens and close to 30 cents per million output tokens, placing it near the most cost-efficient end of currently available frontier-adjacent models.
