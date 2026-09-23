---
id: collect-huggingface/huggingface/openai-gpt-oss-20b
title: "gpt-oss-20b - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "OpenAI", "vLLM"]
dates: ["2026-09-23"]
keywords: ["agentic", "apache", "benchmarks", "consumer", "fine-tuning", "gguf", "gpu", "inference", "latency", "license", "memory", "moe"]
source: docs/RAG/Collect RAG/03_huggingface/openai-gpt-oss-20b.md
source_anchor: ""
source_lines: [1, 46]
sha256: e6dec09af5876dcb8e6d86b70241a11671d56521d94d4fd45317e5512484f79c
---

# gpt-oss-20b - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/openai/gpt-oss-20b
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

gpt-oss-20b is OpenAI's smaller open-weight language model, part of the gpt-oss series released under the permissive Apache 2.0 license. It is a Mixture-of-Experts (MoE) model with 21B total parameters and 3.6B active parameters, designed for lower latency, local, and specialized use cases. Its larger sibling, gpt-oss-120b (117B total / 5.1B active), targets production and high-reasoning workloads on a single 80GB GPU. Both models were post-trained with native MXFP4 quantization of the MoE weights, allowing gpt-oss-20b to run within 16GB of memory. They were trained on OpenAI's harmony response format and should only be used with it, as they will not work correctly otherwise. Highlights include configurable reasoning effort (low, medium, high), full chain-of-thought access (for debugging, not for end users), fine-tunability, and native agentic capabilities such as function calling, web browsing, Python code execution, and structured outputs. Reasoning effort is set in the system prompt, e.g., "Reasoning: high". Supported inference paths include Transformers (pipeline or `transformers serve`), vLLM (dedicated gpt-oss wheel), PyTorch/Triton reference implementations, Ollama (`ollama pull gpt-oss:20b`), and LM Studio. The model can be fine-tuned on consumer hardware, unlike the larger model which needs a single H100 node. On the HF evaluation results, GPQA Diamond is reported at 58.59 and SWE-bench Verified (reasoning medium) at 53.2, with reasoning-medium scores of 66.0 (no tools) and 67.1 (with tools). It reports 6.69M downloads in the last month and 236 community quantizations.

## Key points

- OpenAI open-weight MoE reasoning model: 21B total / 3.6B active parameters.
- Runs within 16GB of memory thanks to native MXFP4 quantization of MoE weights.
- Apache 2.0 license; fine-tunable even on consumer hardware.
- Must use the harmony response format; configurable reasoning effort (low/medium/high).
- Agentic: function calling, web browsing, Python execution, structured outputs.
- Runtimes: Transformers, vLLM, PyTorch/Triton, Ollama, LM Studio.
- Benchmarks: GPQA Diamond 58.59, SWE-bench Verified 53.2 (medium reasoning).

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 21B |
| Active parameters | 3.6B |
| Architecture | Mixture-of-Experts (MoE), text generation |
| Quantization | Native MXFP4 (8-bit) MoE weights; BF16/U8 tensor types |
| Memory footprint | Within 16GB |
| License | Apache 2.0 |
| Response format | OpenAI harmony (mandatory) |
| Reasoning levels | Low, medium, high |
| GPQA Diamond | 58.59 |
| SWE-bench Verified | 53.2 (reasoning medium) |
| Fine-tuning | Consumer hardware |
| Downloads last month | 6,687,853 |
| Paper | arXiv:2508.10925 |

## Why this source matters for the RAG

This card documents the accessible, locally runnable member of OpenAI's open-weight family, making it a key reference for consumer-hardware deployment and low-latency reasoning. It pairs with the gpt-oss-120b and unsloth GGUF entries to cover the full gpt-oss deployment spectrum.
