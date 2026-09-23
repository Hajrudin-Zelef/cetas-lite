---
id: collect-huggingface/huggingface/unsloth-gpt-oss-120b-bf16
title: "gpt-oss-120b-BF16 (Unsloth) - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "OpenAI", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["agentic", "apache", "benchmark", "fine-tuning", "gguf", "gpu", "license", "memory", "moe", "mxfp4", "open-weight", "parameters"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-gpt-oss-120b-BF16.md
source_anchor: ""
source_lines: [1, 44]
sha256: d2dea396a359891f5d77a3ba52801f4310b8327fa0caef248932177eec4d7294
---

# gpt-oss-120b-BF16 (Unsloth) - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/gpt-oss-120b-BF16
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This repository by Unsloth AI provides the full-precision BF16 Safetensors version of OpenAI's gpt-oss-120b model, intended for users who want the unquantized weights rather than the MXFP4 native format or GGUF quantizations. The base model is OpenAI's flagship open-weight Mixture-of-Experts (MoE) language model with 117B total parameters and 5.1B active parameters, released under the permissive Apache 2.0 license and trained on the harmony response format, which must be used for correct operation. It is designed for powerful reasoning, agentic tasks, and general developer use, with configurable reasoning effort (low, medium, high), full chain-of-thought access, fine-tunability, and native agentic capabilities including function calling, web browsing, Python code execution, and structured outputs. The BF16 checkpoint is significantly larger than the MXFP4 original, so it requires more memory; the original MXFP4 format is what allows gpt-oss-120b to run on a single 80GB GPU such as an H100 or MI300X. Supported runtimes for this BF16 release include Transformers, vLLM, SGLang, Docker Model Runner, and Unsloth Desktop. Reasoning level is set via the system prompt (e.g., "Reasoning: high"). No dedicated benchmark table is included on this card; the model inherits the capabilities of the original gpt-oss-120b. The repository has 9 likes and 3,383 downloads in the last month, and is part of Unsloth's gpt-oss collection alongside GGUF, 4-bit and other 16-bit formats. It can be fine-tuned, and the larger model can be fine-tuned on a single H100 node.

## Key points

- Unsloth BF16 (full-precision) Safetensors release of OpenAI gpt-oss-120b.
- Base model: 117B total / 5.1B active MoE, Apache 2.0, harmony response format.
- For users needing unquantized weights; larger memory footprint than native MXFP4.
- Configurable reasoning effort (low/medium/high) and full chain-of-thought.
- Agentic: function calling, web browsing, Python execution, structured outputs.
- Runtimes: Transformers, vLLM, SGLang, Docker Model Runner, Unsloth Desktop.
- 3,383 downloads last month; 9 likes; part of Unsloth gpt-oss collection.

## Technical data / figures

| Attribute | Value |
|---|---|
| Base model | openai/gpt-oss-120b |
| Total / active params | 117B / 5.1B (MoE) |
| Precision | BF16 |
| Format | Safetensors |
| License | Apache 2.0 |
| Response format | OpenAI harmony (mandatory) |
| Reasoning levels | Low, medium, high |
| Runtime engines | Transformers, vLLM, SGLang, Docker Model Runner |
| Fine-tuning | Single H100 node |
| Downloads last month | 3,383 |
| Likes | 9 |

## Why this source matters for the RAG

This card documents the full-precision redistribution of OpenAI's largest open-weight model, which matters for teams that need unquantized weights for fine-tuning or numerical fidelity. It completes the gpt-oss coverage in the knowledge base alongside the MXFP4 originals and the Unsloth GGUF quants.
