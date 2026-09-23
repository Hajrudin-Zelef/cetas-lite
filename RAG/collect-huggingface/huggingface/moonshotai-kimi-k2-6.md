---
id: collect-huggingface/huggingface/moonshotai-kimi-k2-6
title: "Kimi K2.6 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Anthropic", "Google", "Hugging Face", "Moonshot", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["kimi", "agent", "agentic", "agents", "attention", "benchmark", "benchmarks", "claude", "context window", "gemini", "int4", "license"]
source: docs/RAG/Collect RAG/03_huggingface/moonshotai-Kimi-K2.6.md
source_anchor: ""
source_lines: [1, 51]
sha256: 5e811403f37726c65117cbcd472472d99bc29fa50bbb6e75de12eb1aaf7da872
---

# Kimi K2.6 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/moonshotai/Kimi-K2.6
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Kimi K2.6 is Moonshot AI's open-source, native multimodal agentic model advancing practical capabilities in long-horizon coding, coding-driven design, proactive autonomous execution, and swarm-based task orchestration. It is a 1T-total-parameter MoE with 32B activated parameters: 61 layers (1 dense), MLA attention (hidden 7168, 64 heads), MoE hidden 2048 per expert, 384 experts with 8 selected per token plus 1 shared expert, SwiGLU activation, 160K vocabulary, a 256K context window, and a MoonViT vision encoder (400M params). It natively takes image and video input (video chat is experimental, official-API only). Native INT4 quantization follows the Kimi K2-Thinking method. License is a Modified MIT. Downloads ~441K/month.

Evaluation compares K2.6 with GPT-5.4 (xhigh), Claude Opus 4.6 (max), Gemini 3.1 Pro (thinking high) and Kimi K2.5. Key results: HLE-Full 34.7 (54.0 with tools), AIME 2026 96.4, HMMT 2026 92.7, IMO-AnswerBench 86.0, GPQA-Diamond 90.5, Terminal-Bench 2.0 66.7, SWE-Bench Pro 58.6, SWE-Bench Multilingual 76.7, SWE-Bench Verified 80.2, LiveCodeBench v6 89.6, BrowseComp 83.2 (86.3 Agent Swarm), DeepSearchQA F1 92.5, Toolathlon 50.0, MCPMark 55.9, Claw Eval pass@3 80.9, OSWorld-Verified 73.1, MMMU-Pro 79.4, CharXiv 80.4 (86.7 w/ python), MathVision 87.4 (93.2 w/ python). Evaluation settings: thinking mode enabled, temperature 1.0, top-p 1.0, 262,144-token context; coding scores averaged over 10 runs; vision over 3 runs; HLE max-gen 98,304 tokens.

Usage: the model offers Thinking and Instant modes (recommended temperature 1.0 thinking / 0.6 instant, top_p 0.95; instant mode via `{'chat_template_kwargs': {"thinking": False}}` or API `extra_body={'thinking': {'type': 'disabled'}}`), supports `preserve_thinking` (retains reasoning across multi-turn interactions, disabled by default), and interleaved thinking + multi-step tool calls. Deployment is via vLLM, SGLang or KTransformers, reusing the Kimi K2.5 deployment method (same architecture); `transformers` must be `>=4.57.1, <5.0.0`. It works best with the Kimi Code CLI agent framework. API access at platform.moonshot.ai with OpenAI/Anthropic-compatible endpoints and a vendor verifier tool. Reference paper: Kimi K2.5 technical report (arXiv 2602.02276).

## Key points

- 1T-total / 32B-active MoE, 384 experts (8/token), MLA, 61 layers, 256K context.
- Native multimodal: text, image, video; MoonViT vision encoder (400M).
- Agent swarm orchestration (300 sub-agents, 4,000 coordinated steps).
- Thinking + Instant modes; `preserve_thinking`; interleaved thinking/tool calls.
- Native INT4 quantization (K2-Thinking method); Modified MIT license.
- Strong results: SWE-bench Verified 80.2, Terminal-Bench 2.0 66.7, BrowseComp 83.2, GPQA-D 90.5, AIME 2026 96.4.
- Serve via vLLM / SGLang / KTransformers; transformers >=4.57.1,<5.0.0.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | moonshotai |
| Model name | Kimi-K2.6 |
| Architecture | MoE, 61 layers (1 dense), MLA, hidden 7168, 64 heads |
| Total params | 1T |
| Active params | 32B |
| Experts | 384 (8/token) + 1 shared |
| Context length | 256K (262,144) |
| Vision encoder | MoonViT (400M) |
| Quantization | Native INT4 |
| Vocab | 160K |
| License | modified-mit |
| Key benchmarks | AIME 2026 96.4, GPQA-D 90.5, SWE-bench Verified 80.2, SWE-bench Multilingual 76.7, Terminal-Bench 2.0 66.7, LiveCodeBench v6 89.6, BrowseComp 83.2, HLE-Full 34.7 (54.0 w/ tools) |
| Paper | Kimi K2.5 (arXiv 2602.02276) |
| Downloads/month | ~440,966 |

## Why this source matters for the RAG

This card documents a frontier open-weight multimodal agentic model (1T-A32B, 256K context, native INT4) with dense benchmark data across reasoning, coding, agentic and vision tasks plus detailed API/thinking-mode usage. It is a key up-to-date reference for retrieval on Moonshot's K2.x line, multimodal agent design, and native INT4 quantization.
