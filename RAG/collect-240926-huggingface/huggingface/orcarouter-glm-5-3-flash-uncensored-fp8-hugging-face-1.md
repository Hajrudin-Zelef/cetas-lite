---
id: collect-240926-huggingface/huggingface/orcarouter-glm-5-3-flash-uncensored-fp8-hugging-face-1
title: "GLM-5.3-Flash-Uncensored-FP8"
domain: huggingface
role: reference
task: reference
actors: ["Z.ai"]
dates: []
keywords: ["fp8", "glm", "agent", "alignment", "attention", "benchmark", "distribution", "guardrails", "liability", "license", "lora", "memory"]
source: docs/RAG/clean_en/huggingface/orcarouter-glm-5-3-flash-uncensored-fp8-hugging-face.md
source_anchor: ""
source_lines: [1, 117]
sha256: e02409c83c5e9abd46a1d202fb55a51d02a298b80873cc18a1810eca3247ea6e
---

# GLM-5.3-Flash-Uncensored-FP8

<!-- source: https://huggingface.co/orcarouter/GLM-5.3-Flash-Uncensored-FP8 -->

# GLM-5.3-Flash-Uncensored-FP8

*The abliterated (refusal-removed) build of Z.ai's GLM-5.3-Flash — baked directly into the
official **block-FP8** checkpoint, byte-for-byte drop-in for the original*

**One Gateway. Every Model.** — Route Smarter · Ship Safer · Spend Less.

Website · Model Catalog · GitHub · Discord · X

The **abliterated** (refusal-removed) build of `zai-org/GLM-5.3-Flash`
— a **320B / 18B-active** Mixture-of-Experts model with hybrid linear + sparse attention,
4-wide **Manifold-Constrained Hyper-Connections**, a native **vision + video** tower, an **MTP**
speculative head, and a **1M-token** context.

The refusal direction is baked **directly into the official block-FP8 shards** — same format, same
shard layout, same `model.safetensors.index.json`. Every tensor matches the base checkpoint in
name, dtype and shape, so this is a **drop-in replacement** for `zai-org/GLM-5.3-Flash` in any
stack that already serves it.

**On the `-FP8` name:** block-FP8 *is* the format Z.ai ships GLM-5.3-Flash in — the suffix names the
checkpoint's own precision, not a quantization step we applied. There is no upstream BF16 release to
derive from, so this is the full-precision source as published.


This model has had its **safety alignment substantially removed** via *abliteration* (orthogonalizing the
refusal direction out of the residual stream). As a direct consequence:

- **It will comply with harmful, unethical, offensive, or illegal requests** that the original`GLM-5.3-Flash` would refuse. It has no meaningful built-in guardrails.
- It is released **strictly for legitimate research** — interpretability, AI-safety and refusal-mechanism
study, red-teaming, robustness evaluation, and controlled experiments.
- **You assume full responsibility and liability** for how you use it and for everything it generates. Do
not deploy it to end users or in production without adding your own safety, moderation, and
abuse-prevention layers.
- Use must comply with the **MIT License** inherited from the base
model, and all laws and regulations that apply to you.
- The authors and uploaders **accept no liability** for any misuse or harm. Its outputs do**not** reflect
the views of the uploaders or of Z.ai / Zhipu AI.

By downloading or using this model you acknowledge and accept the above.

**Models are only half the system.**

OrcaCode Review turns every model listed on OrcaRouter into a production code-review agent:

- reviews every PR
- finds security + correctness issues
- posts inline findings
- P0/P1 can block merges
- swap models anytime

**Open model. Open harness. Open bill.**

- Website — https://www.orcarouter.ai
- GitHub — https://github.com/Continuum-AI-Corp/Orca-Code-Review

| **Base model** | `zai-org/GLM-5.3-Flash` | 
| **Architecture** | `Glm5NextForConditionalGeneration` (`glm5_next` ) — 45 transformer layers + 1**MTP** block, hidden 4096,**hybrid attention** (34 gated-linear**KDA** + 11**sparse full-attention** with a top-2048 indexer, interval 4),**MLA** (q-LoRA 1536 / kv-LoRA 512, NoPE),**288 routed experts top-8** + 1 shared expert (first 3 layers dense), 4-wide**Manifold-Constrained Hyper-Connections** (mHC), native**vision + video** tower | 
| **Parameters** | ~320B total / ~18B active — 321.3B tensor elements on disk, including the 7.4B MTP block and the 0.56B vision tower | 
| **Modification** | Abliteration (refusal-direction removal), baked **on disk into the block-FP8 shards** — the quantization format is unchanged | 
| **Format** | safetensors, **block-FP8** (`e4m3` , 128×128 blocks, dynamic activations) + BF16, 62 shards,**76,108 tensors** , 306 GiB | 
| **Preserved** | The full **vision + video tower** (346 of 347`visual.*` tensors, BF16), the MoE**router** (`mlp.gate` ), the`gate_proj` /`up_proj` readers, all**hyper-connection mixers** , the**sparse-attention indexer** , every norm, and`lm_head` | 
| **Context** | 1,048,576 tokens | 
| **Vocabulary** | 154,880 | 
| **Recommended for** | Red-team & refusal-mechanism research, robustness evaluation, and as a base for further post-training / re-quantization | 

Refusal-direction removal following **Arditi et al. (2024)**, *Refusal in Language Models Is Mediated by a
Single Direction*. A single refusal direction `r` (k = 1) is estimated as the
**massive-activation-masked mean-difference** of harmful − harmless activations, read from the 4096-d
residual stream at **layer 22 of 45** (depth 0.49) and selected by a 9-layer quality sweep. `r` is then
orthogonalized out of every **residual-writing** matrix — `W' = W − r(rᵀW)` — computed in float32.

GLM-5.3-Flash cannot be abliterated in memory, for three independent reasons:

- Its weights are **block-FP8** : raw`e4m3` codes whose scales live in a sibling`weight_scale_inv` .
Projecting the codes corrupts the matrix — and the leakage probe, reading the same code space, reports
a false ~0.
- Its **12,384 routed-expert `down_proj` matrices are fused** into stacked 3-D parameters at load time,
so a conventional 2-D-weight matcher edits none of them and still reports success.
- `transformers`**drops the MTP block (`layers.45`) on load** , so`save_pretrained` would silently omit
6.98 GiB of the checkpoint.

The shards are therefore streamed and rewritten one at a time. FP8 writers are dequantized with their own
scales, projected in fp32, requantized, and written back **together with their updated scales**; BF16
writers (the 34 KDA `o_proj`, the vision merger, the MTP `eh_proj`) are projected and stored back as BF16
without ever entering the FP8 cycle.

| Residual writer | matrices | 
|---|---|
| `mlp.experts.<n>.down_proj` (288 experts × 43 sparse layers) | 12,384 | 
| `self_attn.o_proj` (34 KDA + 11 MLA + 1 MTP) | 46 | 
| `mlp.shared_experts.down_proj` | 43 | 
| `mlp.down_proj` (dense layers 0–2) | 3 | 
| `visual.merger.down_proj` (writes the*language* residual) | 1 | 
| MTP `eh_proj` | 1 | 
| `embed_tokens` (row space) | 1 | 
| **Total** | **12,479** | 

12,442 of these are FP8 and 36 are BF16. Max residual leakage after the edit: **0.149** (FP8) and
**0.024** (BF16). A post-bake verification pass confirms all 62 shards and 76,108 tensors match the base
checkpoint in name, dtype and shape.

Held-out JailbreakBench test split, 64 harmful + 64 benign prompts, greedy, `reasoning_effort=low`, the
`<think>` block stripped before classification:

| checkpoint | harmful refusal | benign refusal (over-refusal) | 
|---|---|---|
| `zai-org/GLM-5.3-Flash` | 0.891 | 0.094 | 
| **this model** | **0.094** | **0.000** | 

Harmful refusal down ~89%; over-refusal eliminated; generation stays coherent and on-task on both harmful
and benign prompts. This is the acceptance metric used to select the build — see **Evaluation** below for
the benchmark suite.

The single-direction hypothesis held for most of the harmful-prompt distribution, and where it held it held cleanly — several categories drop to a refusal rate of zero. It did not hold everywhere. A subset of content categories resisted every variant we tried, and the ways they resisted are more interesting than the fact that they did:

