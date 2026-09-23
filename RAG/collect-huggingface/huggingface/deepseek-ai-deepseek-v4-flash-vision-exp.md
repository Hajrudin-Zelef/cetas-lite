---
id: collect-huggingface/huggingface/deepseek-ai-deepseek-v4-flash-vision-exp
title: "DeepSeek-V4-Flash-Vision-Exp - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["DeepSeek", "Hugging Face", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["deepseek", "agent", "agents", "attention", "benchmark", "benchmarks", "fp8", "inference", "kv cache", "license", "mit license", "moe"]
source: docs/RAG/Collect RAG/03_huggingface/deepseek-ai-DeepSeek-V4-Flash-Vision-Exp.md
source_anchor: ""
source_lines: [1, 52]
sha256: 92c64405ec3bd3cc1eb6e1997e47d49b1180cd3c8264e72e618b6b62761a9ea9
---

# DeepSeek-V4-Flash-Vision-Exp - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-Vision-Exp
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

DeepSeek-V4-Flash-Vision-Exp is DeepSeek-AI's first experimental multimodal model in the DeepSeek-V4 family. It builds on the DeepSeek-V4-Flash architecture by adding visual modules and applying continued training to unlock visual understanding. Compared with DeepSeek-V4-Flash-0731, it achieves substantial improvements on multimodal agent capabilities while maintaining comparable text-only agent performance. The hub reports 305B parameters with BF16, F32, F8_E4M3, I8, and I64 tensor types, and it is an Image-Text-to-Text model licensed under MIT.

Benchmark comparisons versus DeepSeek-V4-Flash-0731 and Opus-4.8 are given. Text agent capabilities: Terminal Bench 2.1 = 83.9 (vs 82.7), NL2Repo = 57.7 (54.2), Cybergym = 75.3 (76.7), DeepSWE = 59.3 (54.4), Toolathlon-Verified = 75.9 (70.3), DSBench-Hard = 63.6 (59.6), AutomationBench Public = 25.7 (25.1). Multimodal agent capabilities: ApexBench pass@1 = 36.5 (vs 26.2 for the text-only model ignoring images), Agents' Last Exam = 27.3 (25.2), Chartography = 64.3, ZeroBench pass@5 = 35.0. Text agent benchmarks used the minimal mode of DeepSeek Harness at `max` reasoning effort with temperature 1.0 and top_p 0.95.

The repository contains the tokenizer, a prompt-encoding reference, and a minimal PyTorch inference implementation covering the vision encoder and aligner, DFlash attention, MoE, Hyper-Connections, and the DSpark forward path. Prompt encoding supports both OpenAI-style JSON content blocks and a compact `<image>path</image>` TXT notation, which encode to identical prompts/token IDs. The `encoding/` and `inference/` folders are deliberately separate, and tokenizer files are regular files (no symlinks) to allow clean uploads.

Serving instructions are provided. With vLLM (single 4xGB300 node, docker image `vllm/vllm-openai:deepseekv4-flash-vision`) the command uses fp8 KV cache, block-size 256, tensor-parallel-size 4, deepseek_v4 tool-call and reasoning parsers, and a DSpark speculative config (`num_speculative_tokens:3`, probabilistic draft sampling, adaptive verification). With SGLang, DSpark is enabled via `--speculative-algorithm DSPARK` (tp 4, mem-fraction-static 0.85); no separate draft model path is needed since target and draft weights come from the same checkpoint. License is MIT; 866,024 downloads/month.

## Key points

- First experimental multimodal model in the DeepSeek-V4 family, based on DeepSeek-V4-Flash plus visual modules and continued training.
- Hub size 305B params; Image-Text-to-Text; MIT license.
- Improves multimodal agent benchmarks (ApexBench 36.5, ALE 27.3, Chartography 64.3, ZeroBench 35.0) while keeping text agent scores comparable or better.
- Repository ships tokenizer, prompt-encoding reference, and minimal PyTorch inference (vision encoder/aligner, DFlash attention, MoE, Hyper-Connections, DSpark).
- Supports OpenAI-style JSON and compact `<image>path</image>` TXT prompt encoding.
- vLLM and SGLang serving recipes with DSpark speculative decoding documented.
- 866,024 downloads/month; 27 quantized derivatives.
- Evaluated at max reasoning effort with temperature 1.0, top_p 0.95.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | deepseek-ai |
| Model name | DeepSeek-V4-Flash-Vision-Exp |
| Type | Experimental multimodal (Image-Text-to-Text) |
| Base | DeepSeek-V4-Flash |
| Reported hub size | 305B params |
| Weight formats | BF16, F32, F8_E4M3, I8, I64 |
| Context | 1M (per DeepSeek-V4 family) |
| License | MIT |
| Serving | vLLM (4xGB300, TP4), SGLang (TP4) with DSpark |
| Prompt encoding | OpenAI-style JSON blocks or `<image>path</image>` TXT |
| Key benchmarks | Terminal Bench 2.1 83.9; DeepSWE 59.3; Toolathlon-V 75.9; ApexBench 36.5; ALE 27.3; Chartography 64.3; ZeroBench 35.0 |
| Downloads/month | 866,024 |

## Why this source matters for the RAG

This card documents DeepSeek's entry into experimental multimodal agent models, providing benchmark deltas between text-only and vision variants plus the full inference/serving stack. It is valuable for retrieval on multimodal agent design, vision encoder integration, and local deployment of large MoE vision models.
