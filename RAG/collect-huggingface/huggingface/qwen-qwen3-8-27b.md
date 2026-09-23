---
id: collect-huggingface/huggingface/qwen-qwen3-8-27b
title: "Qwen3.8-27B - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "Meta", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["qwen", "agent", "agentic", "agents", "apache", "attention", "benchmark", "benchmarks", "embedding", "license", "multimodal", "muse"]
source: docs/RAG/Collect RAG/03_huggingface/Qwen-Qwen3.8-27B.md
source_anchor: ""
source_lines: [1, 53]
sha256: 27bb3b39babf60e6810c4c21d8d667eb9a7a7c05a978ebdaee5a9abedb5d956c
---

# Qwen3.8-27B - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Qwen/Qwen3.8-27B
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Qwen3.8-27B is the most capable dense model in the Qwen3.8 generation of the Qwen open-model family, built on the architectural foundation of Qwen3.5. It is a native vision-language causal model (27B params, hidden dim 5120, 64 layers, 248,320 padded token embedding) that understands images and videos, with flexible thinking control. It is licensed Apache-2.0, tagged qwen3_5/conversational, has 28B params on the Hub (BF16) with a chat template, and 6,912,469 monthly downloads. The hidden layout is 16 × (3 × (Gated DeltaNet → FFN) → 1 × (Gated Attention → FFN)). Gated DeltaNet: 48 linear-attention heads for V, 16 for QK, head dim 128. Gated Attention: 24 heads for Q, 4 for KV, head dim 256, RoPE dim 64. FFN intermediate 17,408. MTP (Multi-Token Prediction) is trained with multiple steps. Context length is 262,144 natively, extensible to 1,000,000 tokens.

Thinking control is central: thinking mode is on by default and can be disabled per request; reasoning depth is tunable via `reasoning_effort` (`xhigh` default, `medium`, `low`); reasoning context from history is retained via `preserve_thinking` (enabled by default). Recommended sampling: thinking mode temp 1.0, top_p 0.95, top_k 20; instruct/non-thinking mode temp 0.7, top_p 0.80, presence_penalty 1.5. Recommended max output lengths within 1M context: 262,144 tokens reasoning, 131,072 final response. For contexts beyond 262,144 tokens, static YaRN RoPE scaling is advised (config and vLLM/SGLang/TokenSpeed command-line examples provided), noting it can affect short-text performance. For hour-scale video, the `longest_edge` video preprocessor setting can be raised (to 469,762,048 for ~224k video tokens).

Benchmarks (vs Qwen3.6-27B, Qwen3.7-Plus, Muse Glimmer-30B, Opus4.6 Max) highlight text: Terminal Bench 2.1 73.0, SWE-bench Pro 61.7, NL2Repo-Bench 42.3, DeepSWE 1.1 42.2, QwenSWEBench 79.0, CoWorkBench 70.7, JobBench 33.4, Agents' Last Exam 20.4, IFBench 79.5, GPQA Diamond 89.2, HLE 30.8, LiveCodeBench v6 90.3. VL: OSWorld-Verified 84.3, WebArena-Verified 64.8, AndroidWorld 81.9, RecreationBench 47.1, ClawEval-MM 57.4, SWE-MM 38.6, Vision2Web 62.9, MathVision 94.6 (with CI), BabyVision 85.6 (with CI), CharXiv 90.2 (with CI), OmniDocBench 1.5 91.1, RealWorldQA 85.9, ERQA 65.5. Deployment: Transformers, vLLM, SGLang, TokenSpeed, and Qwen Cloud (1M context, built-in tools, coming soon). API examples cover text, image, video input, non-thinking mode, and disabling preserved thinking.

## Key points

- Dense 27B native vision-language model (image + video understanding), Apache-2.0.
- 64 layers, Gated DeltaNet + Gated Attention hybrid layout, MTP multi-token prediction.
- Native context 262,144; extensible to 1M via YaRN.
- Thinking on by default; `reasoning_effort` (xhigh/medium/low) and `preserve_thinking` controls.
- Strong agentic/coding: Terminal Bench 2.1 73.0, SWE-bench Pro 61.7, DeepSWE 1.1 42.2, LiveCodeBench v6 90.3.
- VL/computer-use: OSWorld-Verified 84.3, WebArena-Verified 64.8, AndroidWorld 81.9.
- Served via Transformers, vLLM, SGLang, TokenSpeed, Qwen Cloud.
- 6.9M downloads/month; 1,206 quantized derivatives.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | Qwen |
| Model name | Qwen3.8-27B |
| Type | Causal LM with vision encoder (Image-Text-to-Text) |
| Parameters | 27B (hub 28B) |
| Layers | 64 (hidden 5120, FFN 17,408) |
| Embedding | 248,320 (padded) |
| Attention | Gated DeltaNet (48V/16QK, d128) + Gated Attention (24Q/4KV, d256) |
| MTP | Multi-Token Prediction |
| Context length | 262,144 native; up to 1,000,000 with YaRN |
| License | Apache-2.0 |
| Weight format | BF16 |
| Thinking | default on; reasoning_effort xhigh/medium/low; preserve_thinking |
| Key benchmarks | SWE-bench Pro 61.7; GPQA-D 89.2; OSWorld-V 84.3; WebArena-V 64.8; AndroidWorld 81.9; OmniDocBench 91.1 |
| Serving | Transformers, vLLM, SGLang, TokenSpeed, Qwen Cloud |
| Downloads/month | 6,912,469 |

## Why this source matters for the RAG

This card is a comprehensive reference for a modern dense multimodal agent-capable model, covering hybrid attention architecture, thinking-effort controls, YaRN long-context scaling, and agentic benchmark data. It is essential for retrieval on Qwen model family specs, local deployment, and agent/VL evaluation.
