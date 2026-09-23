---
id: collect-huggingface/huggingface/xiaomimimo-mimo-v2-6-pro-rl
title: "MiMo-V2.6-Pro-RL - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Anthropic", "Hugging Face", "OpenAI", "OpenRouter", "SGLang", "Xiaomi", "vLLM"]
dates: ["2026-09-23"]
keywords: ["agent", "agentic", "agents", "attention", "benchmarks", "claude", "compute", "cybersecurity", "distillation", "fp8", "gpt-5.6", "license"]
source: docs/RAG/Collect RAG/03_huggingface/XiaomiMiMo-MiMo-V2.6-Pro-RL.md
source_anchor: ""
source_lines: [1, 57]
sha256: d50599e33072ab60df36ad5fd49e86011cdc6c6f12c805fde68a631430421d46
---

# MiMo-V2.6-Pro-RL - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/XiaomiMiMo/MiMo-V2.6-Pro-RL
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

MiMo-V2.6-Pro-RL is the flagship checkpoint of Xiaomi's MiMo-V2.6 series, built to scale reinforcement learning toward self-improvement by scaling RL compute, environment diversity, and grader compute together. It is a native omnimodal model handling text, image, video, and audio in a single model, with 1M-token context for long repositories, tool traces, and multi-session agent runs. Architecture: sparse MoE with 1.02T total / 42B activated parameters; the LLM backbone has 70 layers (60 sliding-window attention SWA + 10 global attention GA), hidden size 6144, 128 Q/KV heads, sliding window 128, 384 routed experts (8 activated), no shared experts, plus a 5-layer SWA MTP speculative decoder (DFlash-style, predicting 7 subsequent tokens). Vision encoder is a 681M-param MiMo ViT (28 layers: 24 SWA + 4 full); audio uses a 308M AudioTokenizer (20 RVQ codebooks) plus a 127M audio patch encoder.

Training methodology highlights: "You Only RL Once" — one mixed RL run across coding, general agents, visual, and cybersecurity tasks, using fully asynchronous GRPO on large batches (1,568 prompts x 16 rollouts per step, billions of tokens per update). Groupwise Agentic Grading introduces Groupwise Reward Synthesis (GRS) and Groupwise Advantage Redistribution (GAR) for a self-improvement loop. Aligned RL cold-starts from self-correction with environment hardening against reward hacking, and MOPD2 (Multi-Prefix Multi-Teacher On-Policy Distillation) extends capabilities to hard-to-verify tasks.

Benchmarks: DeepSWE v1.1 71.9 (vs Claude Opus 5 74.0, GPT-5.6 Sol 73.0), ProgramBench 26.5, MiMo Code Bench 63.2, Toolathlon-Verified 76.9, OSWorld-Verified 82.0, JobBench 62.0, CyberGym 94.0, SEC Bench Pro 66.3, MiMo VisualCoding 72.3. License: MIT. Weights in BF16/F32/F8_E4M3/U8. Deployment via SGLang (e.g., --tp 16 --dp 2, EAGLE speculative decoding) and vLLM (tensor-parallel-size 8, MiMo-V2.5 recipe). Recommended sampling temperature=1.0, top_p=0.95. Also on ModelScope, AI Studio, MiMo Code, Desktop, OpenRouter.

## Key points

- Flagship of the MiMo-V2.6 series scaling RL toward self-improvement.
- Native omnimodal: text, image, video, audio; 1M-token context.
- Sparse MoE: 1.02T total / 42B activated; 384 routed experts (8 active).
- "You Only RL Once" — single mixed RL run across coding, agents, visual, cybersecurity.
- Groupwise Reward Synthesis (GRS) + Groupwise Advantage Redistribution (GAR) agentic grading.
- 5-layer MTP speculative decoder (7 tokens predicted per forward pass).
- MIT license; BF16/F32/FP8/U8 weight formats.
- Deploy via SGLang (EAGLE speculative decoding) or vLLM; temp 1.0, top_p 0.95.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | XiaomiMiMo |
| Model name | MiMo-V2.6-Pro-RL |
| Architecture | Sparse MoE, hybrid SWA/GA |
| Total params | 1.02T |
| Activated params | 42B |
| Backbone layers | 70 (60 SWA + 10 GA) |
| Hidden size | 6144 |
| Experts | 384 routed (8 activated), no shared |
| Sliding window | 128 |
| Context length | 1M tokens |
| Modalities | Text, Image, Video, Audio |
| Vision encoder | MiMo ViT, 681M params |
| Audio encoders | 308M AudioTokenizer + 127M patch encoder |
| Speculative decoder | 5-layer SWA MTP |
| License | MIT |
| Weight formats | BF16, F32, F8_E4M3, U8 |
| Benchmarks | DeepSWE 71.9, Toolathlon-Verified 76.9, OSWorld 82.0, CyberGym 94.0 |
| Sampling | temperature=1.0, top_p=0.95 |
| Downloads/month | 4,070 |

## Why this source matters for the RAG

This card captures a state-of-the-art open-weight omnimodal agent model and its novel RL training methodology (GRS/GAR, You Only RL Once, MOPD2), making it critical reference material for retrieval on large-scale RL, agentic grading, speculative decoding, and omnimodal MoE architectures.
