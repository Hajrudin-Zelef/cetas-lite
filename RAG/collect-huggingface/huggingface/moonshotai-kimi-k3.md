---
id: collect-huggingface/huggingface/moonshotai-kimi-k3
title: "Kimi K3 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Anthropic", "Hugging Face", "Moonshot", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["kimi", "agent", "agentic", "attention", "benchmark", "benchmarks", "claude", "context window", "fable 5", "glm", "gpt-5.6", "license"]
source: docs/RAG/Collect RAG/03_huggingface/moonshotai-Kimi-K3.md
source_anchor: ""
source_lines: [1, 53]
sha256: c7cd6ea56813c4119750de057481f2ae40b23f65b7c89cdfcf81d2078970bcba
---

# Kimi K3 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/moonshotai/Kimi-K3
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Kimi K3 is Moonshot AI's open-weight, native multimodal agentic model and the world's first open 3T-class model. It is a 2.8T-parameter MoE built on two new mechanisms — Kimi Delta Attention (KDA) and Attention Residuals (AttnRes) — and scales MoE sparsity with a Stable LatentMoE framework that activates 16 of 896 experts (104B active parameters), yielding roughly a 2.5x improvement in scaling efficiency over Kimi K2. It natively understands text, images and video within one model and supports a 1-million-token context window. Architecture summary: 93 layers (1 dense + 69 KDA + 24 Gated MLA), attention hidden 7168, 96 heads, latent MoE dim 3584, per-expert hidden 3072, 2 shared experts, vocab 160K, SiTU-GLU activation, MoonViT-V2 vision encoder (401M params). Quantization-aware training uses MXFP4 weights with MXFP8 activations. License: Kimi K3 License (custom). Downloads ~1.86M/month.

Evaluation is extensive across reasoning, coding, agentic and vision benchmarks, with a `max` reasoning effort setting. Highlights: GPQA Diamond 93.5, HLE-Full 43.5/56.0 (w/ tools), AIME-style 90s+, DeepSWE 67.5, ProgramBench 77.8, Terminal-Bench 2.1 88.3, FrontierSWE 81.2, SWE-Marathon 42.0, BrowseComp 91.2, DeepSearchQA F1 95.0, MCPMark-Verified 94.5, Toolathlon-Verified 76.5, OSWorld-Verified 84.8, Video-MME 90.0, MMMU-Pro 81.6/83.4, CharXiv 84.8/91.3, MathVision 94.3/97.8, ZeroBench 23.0/41.0 — beating proprietary frontier models (Claude Fable 5, GPT-5.6 Sol, Claude Opus 4.8, GLM-5.2) on many tasks. Detailed footnotes describe evaluation harnesses (Kimi Code, Claude Code, Codex, Terminus 2) and settings.

Deployment: official API via platform.kimi.ai (`kimi-k3`, OpenAI/Anthropic-compatible); recommended engines vLLM (recipes), SGLang (cookbook) and TokenSpeed. Kimi K3 always has thinking enabled and returns `reasoning_content`; thinking effort is set via `reasoning_effort` ("low"/"high"/"max", default "max"). It was trained in preserved-thinking-history mode, so multi-turn/tool calls require passing back the full assistant message including `reasoning_content` and `tool_calls`. Best used with the Kimi Code CLI agent framework.

## Key points

- World's first open 3T-class model: 2.8T total / 104B active (16/896 experts).
- Built on Kimi Delta Attention (KDA) + Attention Residuals + Stable LatentMoE.
- Native multimodal (text/image/video) with 1M-token context; MoonViT-V2 encoder.
- MXFP4 weights / MXFP8 activations via quantization-aware training.
- Frontier agentic results: Terminal-Bench 2.1 88.3, BrowseComp 91.2, DeepSWE 67.5, MCPMark-Verified 94.5.
- Thinking always on; `reasoning_effort` low/high/max; preserved-thinking required.
- Kimi K3 License; vLLM / SGLang / TokenSpeed deployment.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | moonshotai |
| Model name | Kimi-K3 |
| Architecture | MoE (KDA + Gated MLA + AttnRes), 93 layers |
| Total params | 2.8T |
| Active params | 104B |
| Experts | 896 (16/token) + 2 shared |
| Context length | 1,048,576 tokens |
| Attention | KDA & Gated MLA, hidden 7168, 96 heads |
| Activation | SiTU-GLU |
| Vision encoder | MoonViT-V2 (401M) |
| Quantization | MXFP4 weights / MXFP8 activations (QAT) |
| Vocab | 160K |
| License | kimi-k3 |
| Key benchmarks | GPQA-D 93.5, DeepSWE 67.5, Terminal-Bench 2.1 88.3, BrowseComp 91.2, Video-MME 90.0, MathVision 94.3/97.8 |
| Modalities | Text, Image, Video |
| Downloads/month | ~1,860,092 |

## Why this source matters for the RAG

This card is a primary source for a frontier open-weight multimodal 1M-context agentic model, providing complete architecture parameters, extensive benchmark tables, quantization details and serving guidance. It anchors up-to-date knowledge on 3T-class MoE design, multimodal reasoning, and long-horizon agentic evaluation for the RAG.
