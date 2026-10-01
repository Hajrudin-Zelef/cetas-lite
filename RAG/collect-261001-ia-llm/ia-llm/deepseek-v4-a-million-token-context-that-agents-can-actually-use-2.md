---
id: collect-261001-ia-llm/ia-llm/deepseek-v4-a-million-token-context-that-agents-can-actually-use-2
title: "deepseek-v4-a-million-token-context-that-agents-can-actually-use"
domain: ia-llm
role: reference
task: reference
actors: ["DeepSeek"]
dates: []
keywords: ["agent", "deepseek", "benchmark", "context window", "parameters", "reasoning"]
source: docs/RAG/collect-261001-ia-llm/deepseek-v4-a-million-token-context-that-agents-can-actually-use.md
source_anchor: ""
source_lines: [67, 71]
sha256: 25165ca6c7ff702e70484c768dc5de7a889b25d36a64ef3a86d4366b75e2f65a
---

# deepseek-v4-a-million-token-context-that-agents-can-actually-use

Both instruct models support three reasoning modes: Non-think (fast, no chain of thought), Think High (explicit reasoning in `<think>` blocks), and Think Max (maximum reasoning effort with a dedicated system prompt). Think Max requires a context window of at least 384K tokens. The recommended sampling parameters across all modes are `temperature=1.0, top_p=1.0`.

The V4-Pro numbers on SWE Verified, MCPAtlas, and the internal R&D benchmark put it at parity with frontier closed models on agent tasks. The open question is how the community's tool harnesses adapt to the `|DSML|` schema and whether the interleaved thinking gains transfer to out-of-domain agent frameworks.

Figures in this blog post are from the technical report at DeepSeek_V4.pdf.
