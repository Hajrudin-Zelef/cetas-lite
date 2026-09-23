---
id: collect-huggingface/huggingface/poolside-laguna-s-2-1-dflash-nvfp4
title: "Laguna S 2.1-DFlash-NVFP4 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["nvfp4", "agentic", "benchmark", "benchmarks", "latency", "license", "moe", "open-weight", "safetensors", "sglang", "speculative decoding", "throughput"]
source: docs/RAG/Collect RAG/03_huggingface/poolside-Laguna-S-2.1-DFlash-NVFP4.md
source_anchor: ""
source_lines: [1, 48]
sha256: f6649e7483fccca135a11cea5d2a73969d80926b36408718b6b4ac7eee44ece9
---

# Laguna S 2.1-DFlash-NVFP4 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/poolside/Laguna-S-2.1-DFlash-NVFP4
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Laguna S 2.1-DFlash-NVFP4 is poolside's DFlash speculator (draft model) for the NVFP4 target poolside/Laguna-S-2.1-NVFP4. It is a 6-layer Laguna-style draft model in BF16 (hub-reported model size 1B params) that enables lower-latency serving of the 117.6B-A8.5B NVFP4 base model via speculative decoding. The speculator was trained with `e0630_rhiemann_baseline` SFT followed by DFlash Stage-2 for 15k steps, and the recommended serving setting is `num_speculative_tokens=7`. DFlash upstream support is in progress in vLLM (#46853), SGLang (#29446) and TRT-LLM (#15666); the intended target model is poolside/Laguna-S-2.1-NVFP4, and it can be paired with the base using `--speculative-config '{"model":"poolside/Laguna-S-2.1-DFlash-NVFP4","num_speculative_tokens":7,"method":"dflash"}'`.

The card reports benchmark results measured with TP=2, temperature=0 and `num_speculative_tokens=15`. Throughput speedups across concurrency levels (1/4/8/16) are: GSM8K 3.324x/2.498x/2.279x/2.302x, MATH-500 2.893x/2.174x/1.948x/1.965x, HumanEval 3.692x/2.742x/2.634x/2.626x, MBPP 2.426x/1.848x/1.719x/1.731x, MT-Bench 2.338x/1.831x/1.772x/1.672x. Acceptance lengths at concurrency 1–16 are: GSM8K 5.775–5.804, MATH-500 4.942–4.935, HumanEval 6.438–6.298, MBPP 4.171–4.123, MT-Bench 4.017–3.981. In practice, speedups are highest at low concurrency (up to ~3.7x on HumanEval) and remain above 1.7x even at concurrency 16, while acceptance length stays consistently in the 4–6.5 token range.

The repo is tagged as a "speculators" / "dflash" model (Safetensors, BF16) and belongs to the Laguna S 2.1 collection. Downloads are ~42,449/month. Like the related Laguna S 2.1 releases, it is part of poolside's ecosystem for serving the open-weight Laguna S 2.1 family locally, and the card links to the NVFP4 target model card and the release blog post.

## Key points

- 6-layer Laguna-style DFlash speculator (BF16, ~1B params) for the NVFP4 base.
- Trained via e0630_rhiemann_baseline SFT + DFlash Stage-2 (15k steps).
- Recommended `num_speculative_tokens=7`; benchmarks run at 15.
- Throughput speedups up to 3.69x (HumanEval, concurrency 1); ≥1.67x at concurrency 16.
- Acceptance length 4–6.5 tokens across benchmarks and concurrency.
- Upstream DFlash support in progress (vLLM #46853, SGLang #29446, TRT-LLM #15666).
- ~42K downloads/month; part of the Laguna S 2.1 collection.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | poolside |
| Model name | Laguna S 2.1-DFlash-NVFP4 |
| Architecture | Laguna-style DFlash speculator, 6 layers, BF16 |
| Params | ~1B (hub) |
| Target model | poolside/Laguna-S-2.1-NVFP4 |
| Recommended spec tokens | 7 |
| License | OpenMDW-1.1 (family) |
| Throughput speedup (TP=2, t=0, n=15) | GSM8K up to 3.324x; MATH-500 up to 2.893x; HumanEval up to 3.692x; MBPP up to 2.426x; MT-Bench up to 2.338x |
| Acceptance length | GSM8K ~5.8, MATH-500 ~4.9, HumanEval ~6.4, MBPP ~4.2, MT-Bench ~4.0 |
| Downloads/month | 42,449 |
| Upstream support | vLLM #46853, SGLang #29446, TRT-LLM #15666 |

## Why this source matters for the RAG

This card documents a concrete, benchmarked DFlash speculative-decoding draft model for a 4-bit frontier MoE, providing exact throughput speedups and acceptance-length data across five benchmarks and four concurrency levels. It is a compact, citable source on efficient speculative decoding for locally served open-weight agentic models.
