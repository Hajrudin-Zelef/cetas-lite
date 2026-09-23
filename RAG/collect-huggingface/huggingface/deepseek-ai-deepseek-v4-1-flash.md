---
id: collect-huggingface/huggingface/deepseek-ai-deepseek-v4-1-flash
title: "DeepSeek-V4.1-Flash - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["DeepSeek", "Hugging Face", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["deepseek", "agentic", "attention", "benchmark", "benchmarks", "context window", "cost", "decode", "distillation", "fp4", "fp8", "inference"]
source: docs/RAG/Collect RAG/03_huggingface/deepseek-ai-DeepSeek-V4.1-Flash.md
source_anchor: ""
source_lines: [1, 55]
sha256: 28583e40fcb04d2f64f237be7e47986cfc688c62d008b9004bfd1ebc30368d93
---

# DeepSeek-V4.1-Flash - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/deepseek-ai/DeepSeek-V4.1-Flash
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

DeepSeek-V4.1-Flash is a multimodal Mixture-of-Experts (MoE) model from DeepSeek-AI focused on pushing the limits of KV cache compression for input-heavy agentic workloads. It has a 552B-parameter backbone and supports contexts of up to one million tokens, natively processing images and text to generate text autoregressively. Its headline efficiency feature is a Causal Encoder-Decoder (CED) architecture: a 40-layer Transformer split into a 20-layer causal encoder followed by a 20-layer decoder, where the decoder's global KV cache is projected from the encoder's final hidden states instead of each decoder layer's own states. This lets the model activate only 8B parameters per token during prefill and 16B during decode.

Two further innovations are highlighted. Compressed Sparse Attention 2 (CSA2) assigns each attention layer one of three static modes (Full, Reindex, Reuse) to share main KV and indexer K across layers and reuse Top-K sparse-attention indices; a Hierarchical Sparse Indexer bounds deeper indexer cost independently of context length. Combined with FP4 main KV caching (E2M1 with E4M3 scales), the global KV cache drops to 890 bytes/token (roughly 1/4 of DeepSeek-V4-Flash and 437x smaller than DeepSeek-V1). SWA Bounded Replay reconstructs missing SWA KV states by replaying only the most recent n_win tokens, reducing persistent KV footprint to ~1/8. Additional components include Single-Pass mHC, Engram conditional memory (196B parameters, token-based sparse lookup), and DSpark speculative decoding. Each MoE layer uses 1 shared expert and 384 routed experts, activating 6 routed experts per token. A DeepSeek-ViT vision encoder (2D-RoPE, 3x3 pixel-unshuffle) plus a two-layer MLP projector handle images.

The model was pre-trained from scratch on a 45T-token multimodal corpus (sparse attention at 64K, extended to 1M at 34T tokens), and post-trained with SFT → RL → on-policy distillation. It supports continuously controllable reasoning effort (integer 1–100). License is MIT. Benchmark strengths include Codeforces rating 3471, Terminal-Bench 2.1 90.6, DeepSWE v1.1 74.2, GPQA Diamond 90.9, and MMMU-Pro 56.5. Recommended sampling: temperature 1.0, top_p 0.95/1.0, 1M context, max_tokens >= 256K. No Jinja chat template is shipped; instead an `encoding` Python reference and the `deepseek-recipe` Rust toolkit are provided. vLLM and SGLang launch instructions are included.

## Key points

- Multimodal MoE with 552B backbone parameters; activates 8B/token during prefill and 16B/token during decode.
- 1M-token context window; native image + text input, autoregressive text output.
- Novel Causal Encoder-Decoder (CED) architecture with CSA2 compressed sparse attention and FP4 KV cache (890 bytes/token).
- Engram conditional memory (196B) and DSpark speculative decoding.
- Pre-trained on 45T tokens; post-trained via SFT → RL → on-policy distillation.
- Controllable reasoning effort (integer 1–100) trading cost for accuracy.
- MIT license; FP8/BF16 weights; 763B params reported on the hub (incl. auxiliary modules).
- Strong agentic/coding results: Codeforces 3471, Terminal-Bench 2.1 90.6, DeepSWE v1.1 74.2.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | deepseek-ai |
| Model name | DeepSeek-V4.1-Flash |
| Architecture | Causal Encoder-Decoder (CED) Transformer, MoE |
| Backbone params | 552B |
| Active params | 8B (prefill) / 16B (decode) |
| Reported hub size | 763B params (safetensors) |
| MoE layout | 1 shared + 384 routed experts/layer, 6 routed active/token |
| Context length | up to 1M tokens |
| KV cache | 890 bytes/token (FP4 E2M1, E4M3 scales) |
| Pre-training data | 45T tokens (multimodal) |
| Modalities | Image-Text-to-Text |
| License | MIT |
| Weight formats | BF16, F32, F8_E4M3, I8 (FP8) |
| Reasoning effort | 1–100 (continuous) |
| Recommended sampling | temp 1.0, top_p 0.95/1.0, max_tokens ≥ 256K |
| Key benchmarks | GPQA-D 90.9, Codeforces 3471, Terminal-Bench 2.1 90.6, DeepSWE v1.1 74.2, MMMU-Pro 56.5 |
| Downloads/month | 570,909 |

## Why this source matters for the RAG

This is a frontier open-weight multimodal MoE model card documenting state-of-the-art KV cache compression and long-context agentic performance, making it a high-value reference for retrieval on efficient inference, MoE design, and 1M-context multimodal reasoning. Its detailed architecture (CED, CSA2, Engram, DSpark) and benchmark tables provide dense, citable technical facts.
