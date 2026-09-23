---
id: collect-huggingface/huggingface/xiaomimimo-mimo-v2-6-flash-rl
title: "MiMo-V2.6-Flash-RL - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "OpenRouter", "SGLang", "Xiaomi", "vLLM"]
dates: ["2026-09-23"]
keywords: ["agent", "agentic", "agents", "benchmarks", "context window", "cyber", "cybersecurity", "distillation", "fp8", "license", "mit license", "moe"]
source: docs/RAG/Collect RAG/03_huggingface/XiaomiMiMo-MiMo-V2.6-Flash-RL.md
source_anchor: ""
source_lines: [1, 57]
sha256: 78a572218fecba1ae6b99d392a932711d57d472b43ae0ec30455b3ad64b8e7c5
---

# MiMo-V2.6-Flash-RL - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/XiaomiMiMo/MiMo-V2.6-Flash-RL
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

MiMo-V2.6-Flash-RL is the efficiency-balanced checkpoint of Xiaomi's MiMo-V2.6 series, built to scale reinforcement learning toward self-improvement. It is a native omnimodal model handling text, image, video, and audio in one model, with a 1M-token context window for long repositories, tool traces, and multi-session agent runs. Architecture: sparse MoE with 309B total / 15B activated parameters; LLM backbone of 48 layers (39 SWA + 9 GA), hidden size 4096, 64 SWA heads (Q/KV 8) and 64 GA heads (Q/KV 4), head dims QK/V = 192/128, sliding window 128, 256 routed experts (8 activated), no shared experts, plus a 5-layer SWA MTP speculative decoder (DFlash-style, predicting 7 subsequent tokens). Vision encoder: 681M-param MiMo ViT (28 layers: 24 SWA + 4 full); audio: 308M AudioTokenizer (20 RVQ codebooks) + 127M audio patch encoder.

Training methodology: "You Only RL Once" — one mixed RL run across coding, general agents, visual, and cybersecurity, using fully asynchronous GRPO on very large batches (1,568 prompts x 16 rollouts per step). Groupwise Agentic Grading introduces Groupwise Reward Synthesis (GRS) and Groupwise Advantage Redistribution (GAR) for a self-improvement loop; Aligned RL cold-starts from self-correction; MOPD2 (Multi-Prefix Multi-Teacher On-Policy Distillation) extends capabilities to hard-to-verify tasks.

Benchmarks (close to Pro despite 3x fewer parameters): DeepSWE v1.1 67.9, ProgramBench 26.0, MiMo Code Bench 61.2, AutomationBench 52.3, Toolathlon-Verified 73.6, Terminal Bench 4.0 28.8, Terminal Bench 2.1 87.6, OSWorld-Verified 80.8, JobBench 61.2, CyberGym 95.1, MiMo Cyber Bench 77.2, MiMo VisualCoding 71.5. License: MIT. Weights in F32/BF16/F8_E4M3/U8. Deployment via SGLang (--tp 8, EAGLE speculative decoding) and vLLM (tensor-parallel-size 4). Recommended sampling temperature=1.0, top_p=0.95. Also on ModelScope, AI Studio, MiMo Code, Desktop, OpenRouter.

## Key points

- Efficiency-balanced checkpoint of the MiMo-V2.6 series (309B total / 15B activated).
- Native omnimodal: text, image, video, audio; 1M-token context.
- Hybrid SWA/GA backbone: 48 layers (39 SWA + 9 GA), 256 routed experts (8 active).
- "You Only RL Once" mixed RL with asynchronous GRPO on huge batches.
- GRS + GAR groupwise agentic grading closes a self-improvement loop.
- 5-layer MTP speculative decoder predicts 7 subsequent tokens.
- MIT license; F32/BF16/FP8/U8 weight formats.
- Deploy via SGLang (EAGLE) or vLLM; temp 1.0, top_p 0.95.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | XiaomiMiMo |
| Model name | MiMo-V2.6-Flash-RL |
| Architecture | Sparse MoE, hybrid SWA/GA |
| Total params | 309B |
| Activated params | 15B |
| Backbone layers | 48 (39 SWA + 9 GA) |
| Hidden size | 4096 |
| Experts | 256 routed (8 activated), no shared |
| Sliding window | 128 |
| Context length | 1M tokens |
| Modalities | Text, Image, Video, Audio |
| Vision encoder | MiMo ViT, 681M params |
| Audio encoders | 308M AudioTokenizer + 127M patch encoder |
| Speculative decoder | 5-layer SWA MTP |
| License | MIT |
| Weight formats | F32, BF16, F8_E4M3, U8 |
| Benchmarks | DeepSWE 67.9, Toolathlon-Verified 73.6, OSWorld 80.8, CyberGym 95.1 |
| Sampling | temperature=1.0, top_p=0.95 |
| Downloads/month | 13,243 |

## Why this source matters for the RAG

This card documents the efficiency-balanced sibling of MiMo-V2.6-Pro-RL, showing how a 15B-active omnimodal MoE approaches flagship agentic performance with the same RL methodology. It is key reference material for retrieval on large-scale RL training, omnimodal MoE models, and efficient long-context agent architectures.
