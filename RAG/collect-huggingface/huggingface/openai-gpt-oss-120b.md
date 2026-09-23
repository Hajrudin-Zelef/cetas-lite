---
id: collect-huggingface/huggingface/openai-gpt-oss-120b
title: "gpt-oss-120b - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["AMD", "Alibaba", "Hugging Face", "Nvidia", "OpenAI", "vLLM"]
dates: ["2026-09-23"]
keywords: ["agentic", "amd", "apache", "benchmarks", "fine-tuning", "gpu", "latency", "license", "memory", "moe", "mxfp4", "nvidia"]
source: docs/RAG/Collect RAG/03_huggingface/openai-gpt-oss-120b.md
source_anchor: ""
source_lines: [1, 46]
sha256: 5d3660388ef4d84b693c477b7734393ff95dbbaf130c1fc658f435faedff1a33
---

# gpt-oss-120b - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/openai/gpt-oss-120b
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

gpt-oss-120b is OpenAI's flagship open-weight language model, released under the permissive Apache 2.0 license and designed for powerful reasoning, agentic tasks, and versatile developer use cases. It is a Mixture-of-Experts (MoE) model with 117B total parameters and 5.1B active parameters, sized to fit on a single 80GB GPU such as an NVIDIA H100 or AMD MI300X. The companion gpt-oss-20b targets lower latency and local use. Both models were post-trained with MXFP4 quantization of the MoE weights, which is why gpt-oss-120b runs on one 80GB GPU and gpt-oss-20b within 16GB of memory; all evaluations were performed with the same MXFP4 quantization. The models were trained on OpenAI's harmony response format and must only be used with it, or they will not work correctly. Key capabilities include configurable reasoning effort (low, medium, high), full chain-of-thought access for debugging (not intended for end users), fine-tunability, and native agentic capabilities such as function calling, web browsing, Python code execution, and structured outputs. Reasoning level is set in the system prompt (e.g., "Reasoning: high"). Supported runtimes include Transformers, vLLM (with a dedicated gpt-oss wheel), PyTorch/Triton reference implementations, Ollama, and LM Studio. On the HF evaluation results, GPQA Diamond is reported at 80.81 and SWE-bench Verified (reasoning medium) at 52.6; reasoning-medium scores of 73.1 (no tools) and 73.5 (with tools) are also listed. The model can be fine-tuned on a single H100 node, and it has 4.87M downloads in the last month with 130+ community quantizations.

## Key points

- OpenAI open-weight MoE reasoning model: 117B total / 5.1B active parameters.
- Fits on a single 80GB GPU (H100, MI300X); post-trained with MXFP4 quantization.
- Apache 2.0 license, no copyleft or patent risk, commercially usable.
- Must be used with the harmony response format; configurable reasoning effort (low/medium/high).
- Agentic: function calling, web browsing, Python execution, structured outputs.
- Runtimes: Transformers, vLLM, PyTorch/Triton, Ollama, LM Studio.
- Benchmarks: GPQA Diamond 80.81, SWE-bench Verified 52.6 (medium reasoning).

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 117B |
| Active parameters | 5.1B |
| Architecture | Mixture-of-Experts (MoE), text generation |
| Quantization | Native MXFP4 (8-bit) MoE weights; BF16/U8 tensor types |
| Hardware target | Single 80GB GPU (H100 / MI300X) |
| License | Apache 2.0 |
| Response format | OpenAI harmony (mandatory) |
| Reasoning levels | Low, medium, high |
| GPQA Diamond | 80.81 |
| SWE-bench Verified | 52.6 (reasoning medium) |
| Fine-tuning | Single H100 node |
| Downloads last month | 4,865,014 |
| Paper | arXiv:2508.10925 |

## Why this source matters for the RAG

This is the canonical card for OpenAI's flagship open-weight model, anchoring the gpt-oss family in the knowledge base and providing authoritative data on MoE sizing, MXFP4 quantization, harmony formatting and agentic capabilities. It is essential for comparisons against Qwen3.6, Gemma 4 and NVIDIA Nemotron entries.
