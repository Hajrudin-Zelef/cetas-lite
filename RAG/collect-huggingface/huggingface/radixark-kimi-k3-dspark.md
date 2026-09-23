---
id: collect-huggingface/huggingface/radixark-kimi-k3-dspark
title: "Kimi K3 DSpark speculator - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "Moonshot", "SGLang"]
dates: ["2026-09-23"]
keywords: ["kimi", "attention", "benchmarks", "distillation", "embedding", "embeddings", "gqa", "inference", "parameters", "safetensors", "sglang", "speculative decoding"]
source: docs/RAG/Collect RAG/03_huggingface/RadixArk-Kimi-K3-DSpark.md
source_anchor: ""
source_lines: [1, 48]
sha256: 30cacf15f459df1427d946dbbd8c014be401bb765a4c8ff024c4194dd8176a5e
---

# Kimi K3 DSpark speculator - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/RadixArk/Kimi-K3-DSpark
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

RadixArk/Kimi-K3-DSpark is a long-context DSpark speculator (draft model) for the Kimi K3 target model (moonshotai/Kimi-K3), enabling faster inference via speculative decoding at context lengths up to 1 million tokens. DSpark extends the DFlash parallel-draft backbone with a Markov logit-bias head and a per-position confidence head. The checkpoint was trained with SpecForge using hidden states from a live SGLang target engine, with the target weights (embedding/unembedding) intentionally excluded. It is a single-file BF16 safetensors of 2,249,289,601 parameters (hub reports ~2B), tagged Transformers/Qwen3 architecture with custom code and requiring `trust_remote_code=True`.

Specifications: the draft is 5 full-attention Qwen3-style GQA layers with hidden size 7168, 64 query heads / 16 KV heads, and `block_size=7`; verification width is 1 current token + 7 draft tokens; auxiliary target layers are [7, 23, 51, 67, 83]; training context was 65,536 tokens. Evaluation uses `acc_len` (SGLang's histogram-native request acceptance length): SWE-Rebench 4.6594, GSM8K 5.4176, MATH500 4.1329, HumanEval 5.5121, MBPP 5.1980, MT-Bench 3.9342, AIME26 2.9893, and RULER V2 1M 4.2553 (MK 4.4658 / MV 4.3081 / QA 3.9919 across prompts spanning 1,000,432–1,047,925 tokens). AIME26 acc_len is broken down by output-length bucket, reaching 4.9194 on 32K+ outputs.

Serving recipe uses SGLang: `sglang serve` with `--speculative-algorithm DSPARK`, `--speculative-draft-model-path RadixArk/Kimi-K3-DSpark`, `--speculative-dspark-block-size 7`, `--speculative-draft-attention-backend trtllm_mha`, `--enable-linear-replayssm-spec`, and `--context-length 1048576`. YaRN-16 is enabled by default in the draft config (`original_max_position_embeddings=65536`, `max_position_embeddings=1048576`). Training details: SpecForge online distillation with a frozen Kimi K3 target on SGLang; loss `0.1 CE + 0.9 L1 distillation + 1.0 confidence BCE`, decay gamma 4.0, 512 sampled anchors per sequence, block_size 7; topology 4 nodes × 4 GB300 (16 ranks), 2 × TP8 target replicas, DP2 sampler, FSDP16 SHARD_GRAD_OP, global batch 512.

## Key points

- DSpark speculator (DFlash + Markov logit-bias + confidence head) for Kimi K3.
- 2.25B-param BF16 draft; 5 Qwen3-style GQA layers; block_size 7; 1M context.
- Trained via SpecForge online distillation against a live SGLang Kimi K3 engine.
- acc_len: HumanEval 5.51, GSM8K 5.42, SWE-Rebench 4.66, MATH500 4.13, RULER V2 1M 4.26.
- SGLang serving recipe with DSPARK, trtllm_mha draft backend, YaRN-16 default.
- ~4.1M downloads/month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | RadixArk |
| Model name | Kimi-K3-DSpark |
| Type | DSpark speculator (speculative decoding draft) |
| Target model | moonshotai/Kimi-K3 |
| Params | 2,249,289,601 (BF16 single-file safetensors) |
| Draft layers | 5 (Qwen3 GQA, hidden 7168, 64q/16kv heads) |
| Block size / verification | 7 draft tokens + 1 current |
| Auxiliary target layers | [7, 23, 51, 67, 83] |
| Trained context | 65,536 tokens; inference up to 1M |
| acc_len highlights | HumanEval 5.5121, GSM8K 5.4176, SWE-Rebench 4.6594, MATH500 4.1329, MT-Bench 3.9342, RULER V2 1M 4.2553 |
| Training | SpecForge online distillation, 4×4 GB300, global batch 512 |
| Downloads/month | ~4,139,009 |

## Why this source matters for the RAG

This card documents a state-of-the-art DSpark speculative-decoding draft model for a 1M-context frontier model (Kimi K3), with exact architecture, training, and acceptance-length benchmarks. It is a valuable citable source on efficient long-context speculative decoding and SGLang DSPARK deployment for the RAG on frontier serving practices.
