---
id: collect-korben/korben/glm-5-2-retour-experience
title: "GLM 5.2 - Le premier modèle IA open source que je garde"
domain: korben
role: reference
task: article
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Hugging Face", "Moonshot", "Z.ai"]
dates: ["2026-06", "2026-09-23"]
keywords: ["glm", "open source", "benchmark", "claude", "consumer", "context window", "deepseek", "fable 5", "kimi", "leaderboard", "license", "llama"]
source: docs/RAG/Collect RAG/01_korben/glm-5-2-retour-experience.md
source_anchor: ""
source_lines: [1, 52]
sha256: 48bf021464c132148ccfb369f0a5880bfe7cb063b4085cfeec342ea1dfc97ac5
---

# GLM 5.2 - Le premier modèle IA open source que je garde

## Metadata

- **Source** : https://korben.info/glm-5-2-retour-experience.html
- **Site** : Korben.info
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The article is a first-hand, honest field report by Korben on GLM 5.2, the latest open-weights large language model from the Chinese lab Z.ai (formerly known as Zhipu AI), released in June 2026. The author states that this is the first time an open-weights model has genuinely satisfied him for his day-to-day coding work, after disappointing experiences with Qwen, Llama, Kimi and DeepSeek, which he says either produced bugs, endless evasive discussions, or conversations that degraded into Chinese characters.

GLM 5.2 is a very large model: 744 billion parameters using a Mixture-of-Experts (MoE) architecture, of which roughly 40 billion parameters are activated per token. A dedicated variant, glm-5.2[1m], offers a context window of up to 1 million tokens. It is published under the MIT license with weights downloadable on HuggingFace.

The key practical highlight is that Korben plugged GLM 5.2 directly into Claude Code. Because Z.ai's API is Anthropic-compatible, he only needs to point Claude Code at the Z.ai endpoint and provide his API key, after which GLM 5.2 responds as if it were Claude. He shares a small bash launcher script that exports ANTHROPIC_BASE_URL to https://api.z.ai/api/anthropic, sets the auth token, maps the default Sonnet and Opus models to "glm-5.2[1m]", and sets CLAUDE_CODE_AUTO_COMPACT_WINDOW to 1000000. His skills and scripts all work unchanged.

The one regret is the inability to run it locally: even quantized to 2-bit, the model requires around 240 GB of RAM, far beyond typical consumer hardware. The API is therefore the only realistic and affordable entry point.

On the Arena.ai front-end coding leaderboard, GLM 5.2 ranks second, just behind Fable 5, making it the first open-weights model at that level of the ranking, since everything above it is proprietary. Korben tempers the enthusiasm by not calling it a "Claude killer" but confirms it matches his subjective experience. He concludes that open-source AI could now genuinely enter his daily workflow. Z.ai offers a GLM Coding Plan starting at $18/month, integrated with Claude Code, Cline and about twenty similar tools.

## Key points

- GLM 5.2 is Z.ai's (ex-Zhipu AI) new open-weights model, released June 2026, MIT-licensed with weights on HuggingFace.
- Architecture: 744B parameters MoE, ~40B active per token, up to 1M-token context in the glm-5.2[1m] variant.
- It integrates directly into Claude Code via Z.ai's Anthropic-compatible API, using existing skills and scripts unchanged.
- First open-weights model the author keeps for real work after disappointing tests with Qwen, Llama, Kimi and DeepSeek.
- Ranks 2nd on the Arena.ai front-end code leaderboard, behind Fable 5, the first open-weights model at this level.
- Cannot run locally: ~240 GB RAM required even at 2-bit quantization; API is the only practical route.
- GLM Coding Plan starts at $18/month and works with Claude Code, Cline and ~20 other tools.

## Technical data / figures

| Item | Value |
|---|---|
| Model | GLM 5.2 |
| Lab | Z.ai (formerly Zhipu AI) |
| Release | June 2026 |
| Parameters | 744 billion (MoE) |
| Active parameters per token | ~40 billion |
| Context window | up to 1,000,000 tokens (glm-5.2[1m]) |
| License | MIT |
| Weights | HuggingFace |
| Local RAM requirement (2-bit) | ~240 GB |
| Arena.ai front-end code rank | 2nd (behind Fable 5) |
| Coding plan price | from $18/month |

## Why this source matters for the RAG

This article documents a concrete, real-world evaluation of a major open-weights model (GLM 5.2) and a practical integration pattern with Claude Code via an Anthropic-compatible API. It provides useful facts and figures (parameter count, context size, hardware requirements, pricing) and a subjective but well-qualified benchmark reference, making it valuable for questions about open-source LLMs, local deployment limits, and coding-assistant tooling.
