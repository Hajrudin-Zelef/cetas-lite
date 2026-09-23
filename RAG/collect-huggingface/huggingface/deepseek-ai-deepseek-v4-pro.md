---
id: collect-huggingface/huggingface/deepseek-ai-deepseek-v4-pro
title: "DeepSeek-V4-Pro - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["DeepSeek", "Hugging Face", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["deepseek", "agentic", "attention", "benchmark", "benchmarks", "context window", "distillation", "fp4", "fp8", "inference", "kv cache", "license"]
source: docs/RAG/Collect RAG/03_huggingface/deepseek-ai-DeepSeek-V4-Pro.md
source_anchor: ""
source_lines: [1, 53]
sha256: 8172ae9c95a1c69fe8891fc71a134445b4ccfad1b5b603f77a3245fa69bae7f1
---

# DeepSeek-V4-Pro - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

DeepSeek-V4-Pro is the flagship of the DeepSeek-V4 preview series: a Mixture-of-Experts (MoE) language model with 1.6T total parameters and 49B activated parameters, supporting a context length of one million tokens. It is accompanied by DeepSeek-V4-Flash (284B total / 13B activated, 1M context). The series introduces three key upgrades: a Hybrid Attention Architecture combining Compressed Sparse Attention (CSA) and Heavily Compressed Attention (HCA), which at 1M context reduces single-token inference FLOPs to only 27% and KV cache to 10% versus DeepSeek-V3.2; Manifold-Constrained Hyper-Connections (mHC) to strengthen residual connections and signal propagation; and the Muon optimizer for faster, more stable training. Both models are pre-trained on more than 32T tokens and post-trained with a two-stage paradigm — independent cultivation of domain experts (SFT + RL with GRPO) followed by unified consolidation via on-policy distillation.

The released checkpoint uses FP4 + FP8 mixed precision (MoE expert parameters in FP4, most others FP8), while the Base variants use FP8 mixed. On the instruct side, the model supports three reasoning modes: Non-think (fast responses), Think High (logical analysis), and Think Max (maximum reasoning with a special system prompt). DeepSeek-V4-Pro-Max is positioned as the best open-source model available today, with strong coding performance and a narrowed gap to closed models on reasoning and agentic tasks. Notable Max scores: MMLU-Pro 87.5, SimpleQA-Verified 57.9, GPQA Diamond 90.1, HLE 37.7, LiveCodeBench 93.5, Codeforces rating 3206, SWE Verified 80.6, Terminal Bench 2.0 67.9, BrowseComp 83.4, and GDPval-AA Elo 1554.

The release ships no Jinja chat template; an `encoding` folder provides Python scripts to convert OpenAI-compatible messages to model input and parse outputs (`encode_messages`, `parse_message_from_completion_text`, with a `thinking_mode` parameter). Recommended sampling is temperature 1.0, top_p 1.0; for Think Max, a context window of at least 384K tokens is advised. The license is MIT. The hub reports 1.6T params (BF16, I64, F32, F8_E4M3, I8) and 537,444 monthly downloads. arXiv 2606.19348.

## Key points

- Flagship MoE with 1.6T total / 49B activated parameters and 1M-token context.
- Hybrid Attention (CSA + HCA) cuts 1M-context FLOPs to 27% and KV cache to 10% vs DeepSeek-V3.2.
- Manifold-Constrained Hyper-Connections (mHC) and Muon optimizer for stability/convergence.
- Pre-trained on >32T tokens; two-stage post-training (domain experts then on-policy distillation).
- FP4 + FP8 mixed precision (experts FP4, others FP8); Base variant is FP8 mixed.
- Three reasoning modes: Non-think, Think High, Think Max.
- DeepSeek-V4-Pro-Max claims best open-source model; Codeforces 3206, LiveCodeBench 93.5, SWE Verified 80.6.
- MIT license; no Jinja template (dedicated `encoding` Python folder); temp 1.0, top_p 1.0.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | deepseek-ai |
| Model name | DeepSeek-V4-Pro |
| Architecture | MoE, hybrid attention (CSA + HCA), mHC |
| Total params | 1.6T |
| Activated params | 49B |
| Context length | 1M tokens |
| Pre-training data | >32T tokens |
| Precision | FP4 + FP8 mixed (experts FP4) |
| Optimizer | Muon |
| Reasoning modes | Non-think / Think High / Think Max |
| License | MIT |
| Sampling | temp 1.0, top_p 1.0; Think Max context ≥ 384K |
| Key benchmarks (Max) | MMLU-Pro 87.5; GPQA-D 90.1; LiveCodeBench 93.5; Codeforces 3206; SWE Verified 80.6; BrowseComp 83.4 |
| arXiv | 2606.19348 |
| Downloads/month | 537,444 |

## Why this source matters for the RAG

This is the flagship model card for the DeepSeek-V4 generation, documenting frontier-scale MoE architecture (hybrid attention, mHC, Muon) and extensive benchmark tables across reasoning modes. It is a central reference for retrieval on million-token context efficiency, open-weight frontier models, and reasoning-effort trade-offs.
