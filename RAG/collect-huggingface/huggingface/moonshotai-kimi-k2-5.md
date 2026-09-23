---
id: collect-huggingface/huggingface/moonshotai-kimi-k2-5
title: "Kimi K2.5 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "Moonshot", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["kimi", "agent", "agentic", "agents", "attention", "benchmark", "benchmarks", "claude", "deepseek", "gemini", "int4", "license"]
source: docs/RAG/Collect RAG/03_huggingface/moonshotai-Kimi-K2.5.md
source_anchor: ""
source_lines: [1, 52]
sha256: 5829a0cd6f9a137326b4a2f0e85110a43acd5027778a686eca332cb5263d262c
---

# Kimi K2.5 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/moonshotai/Kimi-K2.5
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Kimi K2.5 is Moonshot AI's open-source, native multimodal agentic model built through continual pre-training on ~15 trillion mixed visual and text tokens atop Kimi-K2-Base. It integrates vision and language understanding with agentic capabilities, instant and thinking modes, and conversational and agentic paradigms. It is a 1T-total/32B-active MoE: 61 layers (1 dense), MLA attention (hidden 7168, 64 heads), MoE hidden 2048 per expert, 384 experts with 8 selected per token plus 1 shared expert, SwiGLU activation, 160K vocabulary, 256K context, and a MoonViT vision encoder (400M params). Native INT4 quantization follows the Kimi K2-Thinking method. License: Modified MIT. Downloads ~313K/month. Changelog (2026.1.29) notes removal of the default system prompt and fixing the `<|media_begin|>` token in the chat template.

Key features: native multimodality (visual knowledge, cross-modal reasoning, vision-grounded tool use), coding with vision (generate code from UI designs and video workflows), and Agent Swarm (self-directed coordinated swarm execution with dynamically instantiated, domain-specific sub-agents). Evaluation (thinking mode, temp 1.0, top-p 0.95, 256K ctx) vs GPT-5.2, Claude Opus 4.5, Gemini 3 Pro, DeepSeek V3.2, Qwen3-VL-235B-A22B-Thinking. Key results: HLE-Full 30.1 (50.2 w/ tools), AIME 2025 96.1, HMMT 2025 95.4, IMO-AnswerBench 81.8, GPQA-Diamond 87.6, MMLU-Pro 87.1, MMMU-Pro 78.5, CharXiv 77.5, MathVision 84.2, OCRBench 92.3, OmniDocBench 1.5 88.8, WorldVQA 46.3, VideoMME 87.4, SWE-bench Verified 76.8, SWE-bench Pro 50.7, SWE-bench Multilingual 73.0, Terminal Bench 2.0 50.8, PaperBench 63.5, LiveCodeBench v6 85.0, BrowseComp 60.6 (74.9 ctx manage, 78.4 Agent Swarm), DeepSearchQA 77.1, LongBench v2 61.0, AA-LCR 70.0.

Usage: Thinking and Instant modes (temp 1.0 thinking / 0.6 instant, top_p 0.95; instant via `{'chat_template_kwargs': {"thinking": False}}`), image and video input (video experimental, official API only), interleaved thinking and multi-step tool calls, and best integration with Kimi Code CLI. Deployment: vLLM, SGLang, KTransformers; transformers >= 4.57.1. Reference: "Kimi K2.5: Visual Agentic Intelligence" (arXiv 2602.02276).

## Key points

- 1T-total / 32B-active MoE; 61 layers, MLA, 384 experts (8/token), 256K context.
- Continual pre-training on ~15T visual + text tokens from Kimi-K2-Base.
- Native multimodal (image/video), Agent Swarm, vision-grounded coding.
- Thinking + Instant modes; interleaved thinking and tool calls.
- Native INT4 quantization; Modified MIT license.
- Strong vision + agentic results: VideoMME 87.4, OCRBench 92.3, BrowseComp 78.4 (swarm), SWE-bench Verified 76.8.
- Serve via vLLM / SGLang / KTransformers; transformers >=4.57.1.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | moonshotai |
| Model name | Kimi-K2.5 |
| Architecture | MoE, 61 layers (1 dense), MLA, hidden 7168, 64 heads |
| Total params | 1T |
| Active params | 32B |
| Experts | 384 (8/token) + 1 shared |
| Context length | 256K |
| Vision encoder | MoonViT (400M) |
| Continual pre-training | ~15T visual + text tokens |
| Quantization | Native INT4 |
| Vocab | 160K |
| License | modified-mit |
| Key benchmarks | AIME 2025 96.1, GPQA-D 87.6, MMLU-Pro 87.1, VideoMME 87.4, OCRBench 92.3, SWE-bench Verified 76.8, BrowseComp 78.4 (swarm), HLE-Full 30.1 (50.2 w/ tools) |
| Paper | arXiv 2602.02276 |
| Downloads/month | ~312,782 |

## Why this source matters for the RAG

This card is the primary reference for Kimi K2.5, an open 1T-class multimodal agentic model with a 15T-token continual-pretraining recipe, dense vision/coding/agentic benchmark tables, and full deployment/usage guidance. It anchors up-to-date RAG knowledge on multimodal MoE design, agent swarms, and native INT4 quantization.
