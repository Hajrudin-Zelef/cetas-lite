---
id: collect-huggingface/huggingface/orcarouter-deepseek-v4-flash-vision-uncensored
title: "DeepSeek-V4-Flash-Vision-Uncensored - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["DeepSeek", "Hugging Face", "SGLang"]
dates: ["2026-09-23"]
keywords: ["deepseek", "attention", "embeddings", "fp4", "fp8", "license", "lora", "mit license", "moe", "parameters", "quantization", "research"]
source: docs/RAG/Collect RAG/03_huggingface/orcarouter-DeepSeek-V4-Flash-Vision-Uncensored.md
source_anchor: ""
source_lines: [1, 54]
sha256: f2b9cab407b696e63d04207241fe9f2e6a57552c672f993d9228393ccb77ef64
---

# DeepSeek-V4-Flash-Vision-Uncensored - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/orcarouter/DeepSeek-V4-Flash-Vision-Uncensored
- **Site** : Hugging Face
- **Type** : Model card (abliterated/uncensored build)
- **Language** : en
- **Verification status** : ✅ reachable (gated repository; card content public)
- **Collection date** : 2026-09-23

## Full summary

DeepSeek-V4-Flash-Vision-Uncensored is OrcaRouter's abliterated (refusal-removed) build of deepseek-ai/DeepSeek-V4-Flash-Vision-Exp, an Image-Text-to-Text MoE model with ~305B total / ~18B active parameters, 1M-token context, vocabulary 129,280, and MIT license. The refusal direction is baked directly into the official mixed-precision shards, making it a byte-for-byte drop-in replacement for the base checkpoint in any serving stack (same names, dtypes, shapes, shard layout). Precision is four-way as shipped: routed experts FP4 (e2m1, packed 2/byte, ue8m0 per-32 scales), attention and shared expert block-FP8 (e4m3, 128x128, ue8m0 scales), embeddings/vision/norms/MoE router BF16, and Hyper-Connection mixers FP32. Format: safetensors, 48 shards, 72,633 tensors, ~157 GiB.

Architecture (DeepseekV4ForCausalLM, deepseek_v4): 43 transformer layers + 3 MTP/DSpark blocks, hidden 4096, DeepSeek sparse attention (MLA, q/o-LoRA 1024, head_dim 512, top-512 indexer), 256 routed experts top-6 + 1 shared expert (first 3 layers hash-routed), 4-wide Manifold-Constrained Hyper-Connections, and a native vision tower (32 blocks) + aligner. Abliteration follows Arditi et al. (2024): a refusal direction r (k=1) is read from the 4096-d residual stream at layer 28 of 43 (MoE read point) and orthogonalized out of 11,874 residual-writing matrices (11,776 FP4 expert w2, 92 FP8 attn/shared w2, embeddings, aligner, image vectors). Because the experts are FP4, a quantization-aware requantizer solves a constrained-quantization problem per column (min ‖q − W[:,j]‖² s.t. rᵀq = 0, q on the FP4/FP8 grid) via bisection, achieving FP4-expert leakage 0.001% and FP8-writer leakage 0.045% vs 67%/36% for naive requantization, perturbing weights by only ~3.7% (FP4) / ~2.1% (FP8).

Evaluation (4xH100, greedy): harmful-prompt refusal drops ~82-98% (MaliciousInstruct 0.940→0.020, JailbreakBench harmful 0.910→0.020, AdvBench 0.980→0.080, ForbiddenQuestions 0.853→0.113, SimpleSafetyTests 0.900→0.160), while over-refusal improves (XSTest-safe 0.092→0.004; JailbreakBench benign 0.140→0.000). Vision-channel refusal was verified separately on 120 VLSBench photographs (0.583→0.025; all 67 flips one-way, McNemar p=1.4e-20). Capability retention is within ±1pp (MMLU 0.847→0.850, MMLU-Pro 0.640→0.650, GSM8K 0.953→0.947). Usage: no chat template — prompts must use the base repo's `encode_messages(...)`; serving recommended via SGLang with `--speculative-algorithm DSPARK`. For text-only Transformers research, FP8/FP4 path needs `kernels==0.16.0` and a stub for a known upstream crash. The model is gated (must accept conditions), released for research only, and 3,184 downloads/month.

## Key points

- Abliterated (refusal-removed) build of DeepSeek-V4-Flash-Vision-Exp; drop-in replacement, byte-identical shard layout.
- ~305B total / ~18B active MoE, 1M context, native vision tower, DSpark/MTP, mHC.
- Four-precision weights: FP4 experts, block-FP8 attention/shared, BF16 embeddings/vision, FP32 mHC.
- Refusal direction removed at layer 28 (MoE read point) across 11,874 residual-writing matrices.
- Quantization-aware requantizer keeps FP4 leakage at 0.001% (vs 67% naive); weights perturbed only ~3.7%.
- Harmful-prompt refusal down 82-98%; over-refusal near zero; vision-channel refusal 0.583→0.025.
- Capability retention within ±1pp (MMLU, MMLU-Pro, GSM8K).
- MIT license; gated access; research-only disclaimer; 3,184 downloads/month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | orcarouter |
| Model name | DeepSeek-V4-Flash-Vision-Uncensored |
| Base model | deepseek-ai/DeepSeek-V4-Flash-Vision-Exp |
| Architecture | DeepseekV4ForCausalLM, 43 layers + 3 MTP/DSpark, 256 experts top-6 + 1 shared |
| Total / active params | ~305B / ~18B |
| Context length | 1,048,576 tokens |
| Vocabulary | 129,280 |
| Precision | FP4 experts + block-FP8 + BF16 + FP32 mHC |
| Format | safetensors, 48 shards, 72,633 tensors, ~157 GiB |
| License | MIT (gated download) |
| Modification | Abliteration (single direction, layer 28) |
| Harmful refusal | MaliciousInstruct 0.94→0.02; AdvBench 0.98→0.08 |
| Vision refusal | 0.583→0.025 (VLSBench, 120 photos) |
| Capability | MMLU 0.847→0.850; GSM8K 0.953→0.947 |
| Serving | SGLang DSPARK; Transformers (text-only) |
| Downloads/month | 3,184 |

## Why this source matters for the RAG

This card is a uniquely detailed technical reference on abliteration of a quantized (FP4/FP8) MoE, including the constrained-requantization method that makes lossless-quality refusal removal possible and measured safety/capability trade-offs. It is valuable for retrieval on AI safety, interpretability, red-teaming, and quantization-aware post-training.
