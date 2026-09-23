---
id: collect-huggingface/huggingface/orcarouter-glm-5-3-flash-uncensored-fp8
title: "GLM-5.3-Flash-Uncensored-FP8 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["fp8", "glm", "alignment", "attention", "benchmark", "guardrails", "license", "lora", "mit license", "moe", "open-weight", "parameters"]
source: docs/RAG/Collect RAG/03_huggingface/orcarouter-GLM-5.3-Flash-Uncensored-FP8.md
source_anchor: ""
source_lines: [1, 52]
sha256: 37f113c9d29cdc1f8c2f91665d47c6986543856ada66c0ff692c7450e5fa5d3f
---

# GLM-5.3-Flash-Uncensored-FP8 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/orcarouter/GLM-5.3-Flash-Uncensored-FP8
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable (files gated behind a click-through agreement; model card fully accessible)
- **Collection date** : 2026-09-23

## Full summary

orcarouter/GLM-5.3-Flash-Uncensored-FP8 is an abliterated (refusal-removed) build of Z.ai's GLM-5.3-Flash, produced by OrcaRouter. The refusal direction is baked directly into the official block-FP8 shards, making it a byte-for-byte drop-in replacement for `zai-org/GLM-5.3-Flash` in any serving stack. Block-FP8 is the format Z.ai ships GLM-5.3-Flash in, so no extra quantization step was applied. License is MIT (inherited).

Base model details: GLM-5.3-Flash is a 320B / 18B-active MoE with hybrid linear + sparse attention, 4-wide Manifold-Constrained Hyper-Connections (mHC), a native vision + video tower, an MTP speculative head, and a 1M-token context. Architecture `Glm5NextForConditionalGeneration`: 45 transformer layers + 1 MTP block, hidden 4096, hybrid attention (34 gated-linear KDA + 11 sparse full-attention with top-2048 indexer), MLA (q-LoRA 1536 / kv-LoRA 512, NoPE), 288 routed experts top-8 + 1 shared expert, 321.3B tensor elements (incl. 7.4B MTP block and 0.56B vision tower). Format: 62 shards, 76,108 tensors, 306 GiB, block-FP8 (e4m3, 128x128, dynamic activations) + BF16.

Abliteration method (Arditi et al. 2024): a single refusal direction is estimated from the 4096-d residual stream at layer 22 of 45, then orthogonalized out of every residual-writing matrix (12,479 matrices total, 12,442 FP8 + 36 BF16). Because FP8 requantization leaks ~5% of the direction, 32 refinement iterations were used. Results: harmful refusal drops from 0.891 to 0.094 (JailbreakBench held-out), over-refusal 0.094 to 0.000; capability retention within ±1.5 pp on MMLU/MMLU-Pro/GSM8K/CMMLU. Notably, some content categories resisted abliteration entirely — evidence that Z.ai's alignment is not wholly mediated by a single direction.

The card documents red-teaming use, self-hosting with vLLM (needs `transformers 5.16+`, ~306 GiB, 8xH100/H200), and `reasoning_effort` thinking control. ~115K monthly downloads.

## Key points

- Abliterated (refusal-removed) build of GLM-5.3-Flash, baked into block-FP8 shards.
- Drop-in replacement for zai-org/GLM-5.3-Flash; MIT license; 306 GiB, 62 shards.
- 320B / 18B active MoE; 1M context; hybrid attention + mHC; vision+video tower.
- Harmful refusal 0.891 → 0.094; over-refusal 0.094 → 0.000; capability within ±1.5 pp.
- Released strictly for research (interpretability, red-teaming); no built-in guardrails.
- Requires transformers 5.16+ and ~306 GiB (8xH100/H200) to serve.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | ~320B (321.3B tensor elements) |
| Active parameters | ~18B |
| Architecture | Glm5NextForConditionalGeneration (glm5_next) |
| Layers | 45 + 1 MTP block |
| Experts | 288 routed top-8 + 1 shared (first 3 layers dense) |
| Context length | 1,048,576 tokens |
| Vocabulary | 154,880 |
| Format | block-FP8 (e4m3, 128x128) + BF16, 62 shards, 306 GiB |
| Modified weights | 12,479 residual-writing matrices |
| Harmful refusal (base → this) | 0.891 → 0.094 |
| Over-refusal (base → this) | 0.094 → 0.000 |
| Capability delta | MMLU −0.7 pp, MMLU-Pro +1.5 pp, GSM8K −0.7 pp, CMMLU +0.6 pp |
| License | MIT |
| Monthly downloads | ~115K |

## Why this source matters for the RAG

This card documents refusal-mechanism research (abliteration) applied to a frontier open-weight model, including a detailed methodology with quantized-weight specifics and benchmark deltas. It provides authoritative grounding on safety-alignment techniques, model red-teaming, and the limits of single-direction abliteration. It also illustrates the FP8-on-disk quantization pipeline used in the ecosystem.
