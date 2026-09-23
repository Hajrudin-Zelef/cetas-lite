---
id: collect-huggingface/huggingface/xiaomimimo-mimo-v2-6-distill-qwen-9b
title: "MiMo-V2.6-Distill-Qwen-9B - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "SGLang", "Xiaomi"]
dates: ["2026-09-23"]
keywords: ["qwen", "agent", "agentic", "benchmark", "benchmarks", "cyber", "cybersecurity", "distillation", "fine-tuning", "license", "mit license", "reasoning"]
source: docs/RAG/Collect RAG/03_huggingface/XiaomiMiMo-MiMo-V2.6-Distill-Qwen-9B.md
source_anchor: ""
source_lines: [1, 50]
sha256: 17bda5cc57a06bc82ec49231f3fd0a74e207e5f53233d3a842e8683b359dcb2c
---

# MiMo-V2.6-Distill-Qwen-9B - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/XiaomiMiMo/MiMo-V2.6-Distill-Qwen-9B
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

MiMo-V2.6-Distill-Qwen-9B is a 9B dense agentic model developed by Xiaomi MiMo through supervised fine-tuning of Qwen3.5-9B on MiMo-generated data, covering coding, general-purpose agent tasks, visual coding, and cybersecurity. It is released as an SFT checkpoint intended as a starting point for open research in agentic reinforcement learning. Architecture is based on Qwen3.5-9B (dense, Transformers/AutoModelForMultimodalLM, Image-Text-to-Text pipeline), with the MiMo v2.6 chat template and thinking enabled via chat_template_kwargs.

Evaluation (vs Qwen3.5-9B baseline): SWE-Bench Verified avg@3 61.1 (vs 60.0), SWE-Bench Pro avg@3 44.6 (vs 32.0), MiMo Code (mini) avg@3 51.6 (vs 19.5), MiMo Cyber (mini) avg@3 31.3 (vs 5.7), AutomationBench v1.0.6 avg@1 30.3 (vs 5.0), Terminal Bench 2.1 avg@1 37.1 (vs 27.0), Toolathlon-Verified avg@1 35.2 (vs 25.9), OfficeQA avg@1 19.5 (vs 9.0), JobBench avg@1 18.3 (vs 2.6), MiMo General (mini) avg@1 62.2 (vs 28.5), MiMo Visual Coding (mini) avg@1 64.0 (vs 61.7).

Training data: weighted SFT mixture of 77.4B total tokens including 27.2B loss-bearing tokens, split across Code (23.2B, 29.9%), Cyber (11.0B, 14.2%), General (22.0B, 28.5%), and Visual (21.2B, 27.4%). License: MIT. Weights in BF16 (9B params). Quickstart uses recent SGLang with Qwen3.5 support (--reasoning-parser mimo), enabling thinking explicitly; also usable with Transformers without trust_remote_code. Community has produced 31 quantizations, 7 finetunes, and 2 merges.

## Key points

- 9B dense agentic model: SFT of Qwen3.5-9B on MiMo-generated data.
- Covers coding, general agent tasks, visual coding, and cybersecurity.
- Released as an SFT checkpoint to bootstrap open research in agentic RL.
- Large gains over the Qwen3.5-9B base across all domains (e.g., MiMo Cyber 5.7 to 31.3).
- 77.4B-token weighted SFT mixture with 27.2B loss-bearing tokens.
- MIT license; BF16 weights.
- Run with recent SGLang (Qwen3.5 support) + mimo reasoning/tool-call parsers; thinking via chat_template_kwargs.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | XiaomiMiMo |
| Model name | MiMo-V2.6-Distill-Qwen-9B |
| Base model | Qwen/Qwen3.5-9B |
| Architecture | Dense (Qwen3.5), Image-Text-to-Text |
| Total params | 9B |
| Active params | 9B (dense) |
| Context length | Inherited from Qwen3.5-9B |
| Training | SFT on MiMo-generated data (77.4B tokens, 27.2B loss-bearing) |
| License | MIT |
| Weight formats | BF16 |
| Benchmarks | SWE-Bench Verified 61.1, SWE-Bench Pro 44.6, MiMo General (mini) 62.2, Visual Coding (mini) 64.0 |
| Deployment | SGLang (mimo parsers), Transformers |
| Downloads/month | 3,253 |

## Why this source matters for the RAG

This card documents a distilled, accessible 9B agentic checkpoint that makes the MiMo-V2.6 RL capabilities available for local research, with detailed per-domain benchmark deltas and training-data composition. It is essential reference material for retrieval on distillation, agentic SFT data mixing, and compact agent models.
