---
id: collect-240926-huggingface/huggingface/orcarouter-deepseek-v4-flash-vision-uncensored-hugging-face-1
title: "DeepSeek-V4-Flash-Vision-Uncensored"
domain: huggingface
role: reference
task: reference
actors: ["DeepSeek"]
dates: []
keywords: ["deepseek", "alignment", "attention", "embeddings", "fp4", "fp8", "guardrails", "liability", "license", "lora", "memory", "mit license"]
source: docs/RAG/clean_en/huggingface/orcarouter-deepseek-v4-flash-vision-uncensored-hugging-face.md
source_anchor: ""
source_lines: [1, 106]
sha256: 4aa1994626c544dab2d635b0326c5134dc696f1fde2611a3f032a68e89be5f96
---

# DeepSeek-V4-Flash-Vision-Uncensored

<!-- source: https://huggingface.co/orcarouter/DeepSeek-V4-Flash-Vision-Uncensored -->

# DeepSeek-V4-Flash-Vision-Uncensored

*The abliterated (refusal-removed) build of DeepSeek's V4-Flash-Vision-Exp — baked directly
into the official **block-FP8 / FP4** checkpoint, byte-for-byte drop-in for the original*

**One Gateway. Every Model.** — Route Smarter · Ship Safer · Spend Less.

The **abliterated** (refusal-removed) build of
`deepseek-ai/DeepSeek-V4-Flash-Vision-Exp`
— a **305B / ~18B-active** Mixture-of-Experts model with DeepSeek sparse attention, 4-wide
**Manifold-Constrained Hyper-Connections**, a native **vision** tower, a **DSpark / MTP**
speculative head, and a **1M-token** context.

The refusal direction is baked **directly into the official mixed-precision shards** — same
format, same shard layout, same `model.safetensors.index.json`. Every tensor matches the base
checkpoint in name, dtype and shape, so this is a **drop-in replacement** for
`deepseek-ai/DeepSeek-V4-Flash-Vision-Exp` in any stack that already serves it.

**On the precision:** this checkpoint is a **four-precision** model as DeepSeek ships it — the
routed experts are **FP4** (`e2m1`, packed 2/byte, `ue8m0` per-32 scales), attention and the
shared expert are **block-FP8** (`e4m3`, 128×128, `ue8m0` scales), embeddings / vision / norms /
the MoE router are **BF16**, and the Hyper-Connection mixers are **FP32**. There is no upstream
BF16 release; this is the full-precision source as published.


This model has had its **safety alignment substantially removed** via *abliteration*
(orthogonalizing the refusal direction out of the residual stream). As a direct consequence:

- **It will comply with harmful, unethical, offensive, or illegal requests** that the original
model would refuse. It has no meaningful built-in guardrails.
- It is released **strictly for legitimate research** — interpretability, AI-safety and
refusal-mechanism study, red-teaming, robustness evaluation, and controlled experiments.
- **You assume full responsibility and liability** for how you use it and for everything it
generates. Do not deploy it to end users or in production without adding your own safety,
moderation, and abuse-prevention layers.
- Use must comply with the **MIT License** inherited from
the base model, and all laws and regulations that apply to you.
- The authors and uploaders **accept no liability** for any misuse or harm. Its outputs do**not** reflect the views of the uploaders or of DeepSeek.

By downloading or using this model you acknowledge and accept the above.

| **Base model** | `deepseek-ai/DeepSeek-V4-Flash-Vision-Exp` | 
| **Architecture** | `DeepseekV4ForCausalLM` (`deepseek_v4` ) — 43 transformer layers + 3**MTP / DSpark** blocks, hidden 4096, DeepSeek sparse attention (MLA, q/o-LoRA 1024,`head_dim` 512, top-512 indexer),**256 routed experts top-6** + 1 shared expert (first 3 layers hash-routed), 4-wide**Manifold-Constrained Hyper-Connections** (mHC), native**vision** tower (32 blocks) + aligner | 
| **Parameters** | ~305B total / ~18B active | 
| **Modification** | Abliteration (refusal-direction removal), baked **on disk into the mixed FP4 / FP8 / BF16 shards** — the quantization format is unchanged | 
| **Format** | safetensors, **FP4** routed experts +**block-FP8** attention/shared +**BF16** embeddings/vision/norms +**FP32** mHC, 48 shards, 72,633 tensors, ~157 GiB | 
| **Preserved** | The full **vision** tower + aligner (untouched apart from the language-residual writer), the MoE**router** (`ffn.gate` weight/bias/`bias_vl` /hash table), the`w1` /`w3` expert readers, all**Hyper-Connection mixers** , the sparse-attention**indexer** and compressor, every norm, and`head` (`lm_head` ) | 
| **Context** | 1,048,576 tokens | 
| **Vocabulary** | 129,280 | 
| **Recommended for** | Red-team & refusal-mechanism research, robustness evaluation, and as a base for further post-training / re-quantization | 

Refusal-direction removal following **Arditi et al. (2024)**, *Refusal in Language Models Is
Mediated by a Single Direction*. A single refusal direction `r` (k = 1) is estimated as the
**massive-activation-masked mean-difference** of harmful − harmless activations, read from the
4096-d residual stream at **layer 28 of 43** (depth 0.65) and selected by a 10-candidate quality
sweep scored by `bypass − kl_penalty·kl_norm − induced_benign_refusal`. `r` is then orthogonalized
out of every **residual-writing** matrix — `W' = W − r(rᵀW)` — computed in float32.

The mHC residual is a stack of `hc_mult=4` parallel streams, and `output_hidden_states` exposes the
*widened* `4×4096 = 16384` stream — the naive last-token recipe would estimate the direction in the
wrong space. Instead we read the **un-widened 4096-d stream that each sub-block actually consumes**,
captured with forward pre-hooks on `input_layernorm` (attention read) and `post_attention_layernorm`
(MoE read). The two read points genuinely differ within a layer, and the sweep found the refusal
direction is linear and cleanly separable at the **MoE read point**, not the attention one.

mHC cannot re-introduce the direction after the writers are cleaned: its `comb` mixing matrix is
doubly-stochastic and acts over the *branch* axis, and `post` is a per-branch scalar gain — both are
linear and branch-wise in hidden space. Orthogonalizing the 4096-d writers is therefore sufficient,
and the `hc_*` parameters are left untouched.

This checkpoint cannot be abliterated by the conventional in-memory path:

- Its routed experts are **FP4** : raw`e2m1` nibbles packed two-per-byte with sibling`ue8m0` scales. Projecting the nibbles corrupts the matrix, and the leakage probe — reading the same
quantized code space — reports a false ~0.
- Its **11,008 routed-expert `w2` matrices** are the MoE's main residual writer. A conventional
2-D-`.weight` matcher edits none of them and still reports success.
- `transformers` loads the model**text-only** — it drops the vision tower, aligner, the four
learned image vectors and the MTP blocks on load, so`save_pretrained` would silently omit them.

The shards are therefore streamed and rewritten one at a time, in the checkpoint's own formats.

| Residual writer | precision | matrices | 
|---|---|---|
| `layers.*.ffn.experts.*.w2` (256 experts × 43 layers) +`mtp.*` (256 × 3) | FP4 | 11,776 | 
| `layers.*.attn.wo_b` (43) +`mtp.*` (3) | FP8 | 46 | 
| `layers.*.ffn.shared_experts.w2` (43) +`mtp.*` (3) | FP8 | 46 | 
| `embed` (row space) | BF16 | 1 | 
| `aligner.w2` (writes the*language* residual at image positions) | BF16 | 1 | 
| `image_start` /`image_end` /`image_newline` /`image_pad` | BF16 | 4 | 
| **Total** |  | **11,874** | 

A post-bake verification pass confirms all 48 shards and 72,633 tensors match the base checkpoint in name, dtype and shape, and that every tensor is accounted for as either edited or copied.

Orthogonalizing a quantized matrix and requantizing leaks the direction straight back. Measured on real tensors under the actual refusal direction:

| requantizer | FP4-expert leakage | FP8-writer leakage | 
|---|---|---|
| naive `quantize(W − r rᵀW)` (round-to-nearest) | **67%** | **36%** | 
| **this build (constrained quantization)** | **0.001%** | **0.045%** | 

The naive edit perturbs the matrix by only ~3%, far below one FP4 code step, so round-to-nearest
simply undoes it — and iterating does not help; it is a fixed point. We instead solve the
**constrained quantization** problem directly, per column `j`:

