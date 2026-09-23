---
id: collect-huggingface/huggingface/zai-org-glm-5-fp8
title: "GLM-5-FP8 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Anthropic", "China", "DeepSeek", "Google", "Hugging Face", "Mistral", "Moonshot", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["fp8", "glm", "agentic", "attention", "benchmark", "benchmarks", "claude", "cost", "deepseek", "gemini", "kimi", "license"]
source: docs/RAG/Collect RAG/03_huggingface/zai-org-GLM-5-FP8.md
source_anchor: ""
source_lines: [1, 52]
sha256: 422d6a648a4ce82d37c7e1854ff1bbf671af921b8a64dd4ba9cdae1587a48ebf
---

# GLM-5-FP8 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/zai-org/GLM-5-FP8
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

GLM-5-FP8 is the FP8-quantized release of GLM-5, the next-generation open-weight model from Z.ai (Zhipu AI) targeted at complex systems engineering and long-horizon agentic tasks. GLM-5 scales from the GLM-4.5's 355B parameters (32B active) to 744B total parameters (40B active), and increases pre-training data from 23T to 28.5T tokens. The hub reports a safetensors size of 754B parameters (F32, BF16 and F8_E4M3 tensor types). A key architectural innovation is the integration of DeepSeek Sparse Attention (DSA), which drastically reduces deployment cost while preserving long-context capacity; the architecture tag is `glm_moe_dsa`. The model is released under the MIT license, with an arXiv technical report (2602.15763, "GLM-5: from Vibe Coding to Agentic Engineering") and a GitHub repository.

Post-training relies on a novel asynchronous reinforcement-learning infrastructure called "slime" (developed at THUDM), which improves RL training throughput and enables finer-grained iterations. The card's benchmark table compares GLM-5 with GLM-4.7, DeepSeek-V3.2, Kimi K2.5, Claude Opus 4.5, Gemini 3 Pro and GPT-5.2 (xhigh). Notable scores: HLE 30.5 (50.4 with tools), GPQA-Diamond 86.0, SWE-bench Verified 77.8, SWE-bench Multilingual 73.3, Terminal-Bench 2.0 56.2 (verified 60.7/61.1), CyberGym 43.2, BrowseComp 62.0 (75.9 with context management), τ²-Bench 89.7, MCP-Atlas 67.8, Tool-Decathlon 38.0, and Vending Bench 2 $4,432.12. An evaluation result for GPQA-Diamond (85.35) is also reported on the hub. Detailed footnotes document evaluation protocols (max generation 131,072 tokens for HLE; OpenHands harness with 200K context for SWE-bench; Terminus 2 with 128K context for Terminal-Bench 2.0; Claude Code 2.1.x for CyberGym/MCP-Atlas; independent runs by Andon Labs for Vending Bench 2).

Deployment is supported by vLLM (v0.19.0+), SGLang (v0.5.10+), KTransformers (v0.5.3+), Transformers (v0.5.4+) and xLLM (v0.8.0+). Example serve commands use tensor-parallel size 8 with MTP speculative decoding and `--tool-call-parser glm47 --reasoning-parser glm45 --enable-auto-tool-choice`. It is a conversational text-generation model in English and Chinese, with ~50K downloads/month.

## Key points

- 744B total / 40B active MoE; FP8 release (F8_E4M3), ~754B safetensors params.
- Integrates DeepSeek Sparse Attention (DSA) for lower deployment cost at long context.
- Targets long-horizon agentic engineering; trained on 28.5T tokens.
- Trained with slime, a novel asynchronous RL infrastructure.
- Strong agentic results: CyberGym 43.2, Terminal-Bench 2.0 56.2, BrowseComp (ctx manage) 75.9, τ²-Bench 89.7.
- MIT license; arXiv 2602.15763; deployable via vLLM, SGLang, KTransformers, xLLM.
- English + Chinese; ~50K downloads/month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | zai-org (Z.ai / Zhipu AI) |
| Model name | GLM-5-FP8 |
| Architecture | MoE Transformer with DeepSeek Sparse Attention (glm_moe_dsa) |
| Total params | 744B |
| Active params | 40B |
| Hub model size | 754B params (safetensors) |
| Pre-training data | 28.5T tokens |
| License | MIT |
| Quantization | FP8 (F8_E4M3); also F32 / BF16 tensors |
| Tasks | Text Generation, conversational, agentic |
| Languages | English, Chinese |
| Key benchmarks | HLE 30.5 (50.4 w/ tools), GPQA-D 86.0, SWE-bench V 77.8, Terminal-Bench 2.0 56.2, BrowseComp 62.0 (75.9 w/ ctx), τ²-Bench 89.7 |
| Serving | vLLM ≥0.19.0, SGLang ≥0.5.10, KTransformers, Transformers, xLLM |
| Downloads/month | ~50,193 |
| Paper | arXiv 2602.15763 |

## Why this source matters for the RAG

This card is a key up-to-date reference for an open frontier agentic MoE (744B-A40B FP8) incorporating DeepSeek Sparse Attention, documenting state-of-the-art long-horizon agentic benchmarks and multi-engine deployment recipes. It provides dense, citable figures on scaling (params, tokens), RL infrastructure (slime), and agentic evaluation methodology.
