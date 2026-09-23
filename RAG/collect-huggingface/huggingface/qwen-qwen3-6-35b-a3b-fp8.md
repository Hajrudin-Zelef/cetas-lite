---
id: collect-huggingface/huggingface/qwen-qwen3-6-35b-a3b-fp8
title: "Qwen3.6-35B-A3B-FP8 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["fp8", "qwen", "agentic", "agents", "apache", "attention", "benchmark", "benchmarks", "context window", "inference", "license", "moe"]
source: docs/RAG/Collect RAG/03_huggingface/Qwen-Qwen3.6-35B-A3B-FP8.md
source_anchor: ""
source_lines: [1, 56]
sha256: 4694329772ac95224fd9481ea27dcd158e1eb2ba5bd3e55268485a2a11e3c5ad
---

# Qwen3.6-35B-A3B-FP8 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Qwen/Qwen3.6-35B-A3B-FP8
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Qwen3.6-35B-A3B-FP8 is the FP8-quantized release of Qwen3.6-35B-A3B, the first open-weight variant of the Qwen3.6 series. It is a causal language model with a vision encoder (image-text-to-text) designed for agentic coding, repository-level reasoning, and general multimodal tasks. The release emphasizes stability and real-world utility, with two flagship upgrades: improved agentic coding (frontend workflows and repo-level reasoning) and "Thinking Preservation," a new option to retain reasoning context from historical messages to streamline iterative development.

Architecturally it is a Mixture-of-Experts (MoE) hybrid model with 35B total parameters and 3B activated per token. It has 40 layers, a hidden dimension of 2048, a padded vocabulary of 248,320, and a hybrid layout of 10 × (3 × (Gated DeltaNet → MoE) → 1 × (Gated Attention → MoE)). The MoE has 256 experts with 8 routed + 1 shared activated. It uses multi-step Multi-Token Prediction (MTP). Context length is 262,144 tokens natively, extensible up to 1,010,000 tokens via YaRN RoPE scaling.

The FP8 quantization is fine-grained with a block size of 128, and performance metrics are nearly identical to the original model. License is Apache-2.0. It is compatible with Hugging Face Transformers, vLLM (>=0.19.0), SGLang (>=0.5.10), and KTransformers.

Benchmarks: SWE-bench Verified 73.4, SWE-bench Multilingual 67.2, SWE-bench Pro 49.5, Terminal-Bench 2.0 51.5, MCPMark 37.0, MMLU-Pro 85.2, GPQA 86.0, AIME26 92.7, LiveCodeBench v6 80.4. Vision: MMMU 81.7, MathVista 86.4, OmniDocBench1.5 89.9, VideoMMMU 83.7. The model thinks by default (thinking mode) and supports `reasoning_effort` and `preserve_thinking`.

## Key points

- FP8-quantized MoE multimodal model: 35B total / 3B active parameters.
- Native context 262,144 tokens, extensible to 1,010,000 via YaRN.
- Agentic coding focus: SWE-bench Verified 73.4, Terminal-Bench 2.0 51.5.
- Thinking mode by default; supports `reasoning_effort` and `preserve_thinking`.
- Vision-language support (images, video); 256 experts (8 routed + 1 shared).
- Apache-2.0 license; served via vLLM, SGLang, KTransformers.
- Quantization is fine-grained FP8 (block size 128), near-lossless.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 35B |
| Active parameters | 3B |
| Architecture | MoE + Gated DeltaNet/Gated Attention + Vision Encoder |
| Layers | 40 |
| Hidden dimension | 2048 |
| Experts | 256 (8 routed + 1 shared) |
| Context length | 262,144 native; up to 1,010,000 with YaRN |
| Vocabulary | 248,320 (padded) |
| License | Apache-2.0 |
| Quantization | FP8 (fine-grained, block 128) |
| SWE-bench Verified | 73.4 |
| Terminal-Bench 2.0 | 51.5 |
| MMLU-Pro | 85.2 |
| GPQA | 86.0 |
| AIME26 | 92.7 |
| MMMU | 81.7 |
| Monthly downloads | ~8.6M |

## Why this source matters for the RAG

Qwen3.6-35B-A3B-FP8 represents a frontier open-weight agentic coding model with a very large context window and strong benchmark results, directly relevant to questions about coding agents and local deployment. The FP8 variant documents quantization trade-offs and serving commands, which are key for practical inference. It anchors the knowledge base on current Qwen3.6 capabilities.
