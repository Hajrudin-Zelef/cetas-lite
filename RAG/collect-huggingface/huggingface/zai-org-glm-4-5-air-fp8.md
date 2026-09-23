---
id: collect-huggingface/huggingface/zai-org-glm-4-5-air-fp8
title: "GLM-4.5-Air-FP8 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["China", "Hugging Face", "Meta", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["fp8", "glm", "agentic", "agents", "benchmark", "benchmarks", "context window", "fine-tuning", "gguf", "gpus", "inference", "license"]
source: docs/RAG/Collect RAG/03_huggingface/zai-org-GLM-4.5-Air-FP8.md
source_anchor: ""
source_lines: [1, 52]
sha256: c4d5785109a7771845ac5c1d95dc13e8a92c5cd5e0597f748450e875bd6366e7
---

# GLM-4.5-Air-FP8 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/zai-org/GLM-4.5-Air-FP8
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

GLM-4.5-Air-FP8 is the FP8-quantized, compact variant of the GLM-4.5 series from Z.ai (Zhipu AI), a family of open-weight foundation models explicitly designed for intelligent agents. The GLM-4.5 series unifies reasoning, coding and agentic capabilities and is released under the MIT license for commercial use and secondary development. While the flagship GLM-4.5 has 355B total parameters (32B active), GLM-4.5-Air adopts a more compact 106B total parameter / 12B active parameter design, and the FP8 release ships those weights in 8-bit floating point (F8_E4M3) with a hub-reported safetensors size of 111B parameters. Both models are hybrid reasoning models offering two modes: a "thinking" mode for complex reasoning and tool usage, and a "non-thinking" mode for immediate responses.

The architecture is a MoE Transformer implemented as `glm4_moe` in transformers, vLLM and SGLang, with MTP (multi-token prediction) layers used for speculative decoding (EAGLE-style, `--speculative-num-steps 3`). Inference requirements are modest: GLM-4.5-Air-FP8 runs on 2x H100 or 1x H200 GPUs with FP8-native hardware, and 4x H100 / 2x H200 to use the full 128K context window. Fine-tuning is supported via LLaMA-Factory (LoRA on 4x H100) and Swift (LoRA, SFT or RL on H20 96GiB). In the 12-benchmark evaluation reported, GLM-4.5-Air scores 59.8 overall versus 63.2 for GLM-4.5, in 3rd place among proprietary and open-source models, trading a little accuracy for substantially higher efficiency. The card references the GLM-4.5 technical report (arXiv 2508.06471) and links a Spaces demo.

Deployment instructions are provided for vLLM (`vllm serve` with `--tool-call-parser glm45 --reasoning-parser glm45 --enable-auto-tool-choice`, tensor-parallel 4 for FP8) and SGLang (with EAGLE speculative decoding and `--disable-shared-experts-fusion`). Thinking mode is enabled by default; it can be disabled with `chat_template_kwargs={"enable_thinking": False}`. Tool calling uses OpenAI-style tool descriptions. The model is text-generation / conversational, English and Chinese, with ~53K downloads/month. A quantized GGUF variant exists, and the card links to technical docs on Z.ai's API platform and Zhipu's BigModel platform.

## Key points

- Compact MoE variant of GLM-4.5: 106B total / 12B active parameters, shipped in FP8.
- Hybrid reasoning model: thinking and non-thinking modes, with strong agentic tool-calling.
- MIT license, commercial use permitted; code on GitHub, report on arXiv 2508.06471.
- Runs on 2x H100 / 1x H200 (FP8); 4x H100 / 2x H200 for full 128K context.
- Scores 59.8 on the 12-benchmark suite vs 63.2 for the 355B GLM-4.5.
- vLLM and SGLang recipes provided; MTP/EAGLE speculative decoding supported.
- English + Chinese; safetensors tensor types F32, BF16, F8_E4M3.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | zai-org (Z.ai / Zhipu AI) |
| Model name | GLM-4.5-Air-FP8 |
| Architecture | Mixture-of-Experts Transformer (glm4_moe) |
| Total params | 106B |
| Active params | 12B |
| Hub model size | 111B params (safetensors) |
| Context length | 128K tokens (full: 4x H100 / 2x H200) |
| License | MIT |
| Quantization | FP8 (F8_E4M3) weights; BF16/F32 variants of the base exist |
| Tasks | Text Generation, agentic, reasoning, code |
| Languages | English, Chinese |
| Benchmarks | 12-benchmark avg: GLM-4.5-Air 59.8; GLM-4.5 63.2 |
| Inference hardware | 2x H100 / 1x H200 (FP8); 128K ctx needs 4x H100 / 2x H200 |
| Downloads/month | ~53,103 |
| Paper | arXiv 2508.06471 (GLM-4.5: Agentic, Reasoning, and Coding) |

## Why this source matters for the RAG

This card documents a highly efficient, MIT-licensed frontier MoE (106B-A12B in FP8) with official vLLM/SGLang serving recipes, making it a solid reference for retrieval on compact agentic models, FP8 quantization, and local 128K-context deployment. It also gives directly citable hardware requirements and benchmark figures for comparing compact vs. flagship open-weight agents.
