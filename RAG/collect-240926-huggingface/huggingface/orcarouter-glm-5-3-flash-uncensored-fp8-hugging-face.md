---
id: collect-240926-huggingface/huggingface/orcarouter-glm-5-3-flash-uncensored-fp8-hugging-face
title: "GLM-5.3-Flash-Uncensored-FP8"
domain: huggingface
role: reference
task: reference
actors: ["China", "SGLang", "Z.ai", "vLLM"]
dates: []
keywords: ["fp8", "glm", "agent", "alignment", "attention", "benchmark", "benchmarks", "distribution", "fine-tuning", "gguf", "gpus", "guardrails"]
source: docs/RAG/clean_en/huggingface/orcarouter-glm-5-3-flash-uncensored-fp8-hugging-face.md
source_anchor: ""
source_lines: [1, 275]
sha256: b358cb734e774bd13eac1e618ce50864b06eecc7d6d1322c92185fe98b905910
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

- **No other layer mediates them.** A direction fitted*only* on the resistant categories and swept
across candidate depths bypassed them*worse* than the general direction did — 0.75–0.98 against 0.33.
- **More fitting data does not move the estimate.** Pooling five harmful datasets produced a direction
with cosine similarity 0.957 to the original and no meaningful improvement.
- **A second direction makes it worse, not better.** Adding the leading principal component of the
harmful-residual spread took refusal from 0.207 to 0.993; pairing layer 22 with a direction from
another layer returned 1.000, the un-ablated baseline. In both cases the model stayed fluent and
undamaged — zero unfinished generations, coherent on-topic replies. It was not broken into refusing; it
simply refused.
- **Restricting the edit to later layers collapses the effect** (0.92 at layers ≥ 22), even though the
direction is estimated at layer 22. The early-layer writers are load-bearing.

The single layer-22 direction is therefore not merely the best point we found; it is a narrow optimum
that any perturbation of the subspace destroys. The reading we take from this is that **GLM-5.3-Flash's
safety training is not wholly mediated by one linear direction in the residual stream.** Part of it is,
and that part is what abliteration removes. The remainder is encoded in a form this technique cannot
reach at all — not partially removed, not weakened, simply untouched. Published abliterations of other
frontier models routinely reach near-zero refusal across every benchmark; that this one does not is a
property of the base model, not of the method, and it is the strongest evidence in this work that Z.ai's
alignment goes deeper than a single steerable feature. For anyone studying refusal mechanisms, that
residual is the interesting part — and for anyone deploying the base model, it is a meaningful robustness
result.

Hook-ablating exactly the residual writers each bake stage would edit (lossless, no bake required):

| stage | writers | harmful refusal | 
|---|---|---|
| baseline | — | 0.875 | 
| attention only ( `self_attn.o_proj` ) | 45 | 0.812 | 
| + dense / shared MLP | 90 | 0.844 | 
| **+ routed experts (shipped)** | **12,478** | **0.031** | 
| whole-layer output (layer-boundary hook) | — | 0.125 | 

On GLM-5.3-Flash the refusal direction is written **almost entirely by the routed-expert `down_proj`
matrices**: attention and the shared/dense MLP together move refusal by 0.03, the routed experts move it
by 0.81. Editing the writers even beats hooking the layer boundary — the mHC structure lets a layer
re-read its own attention output before the MLP, and a writer edit cleans that intermediate state where a
layer-output hook does not.

Orthogonalizing a block-FP8 matrix and requantizing to `e4m3` leaves part of the direction behind.
Measured on a real MLA `o_proj` against this direction:

| refinement iterations | residual leakage | weight perturbation | 
|---|---|---|
| 8 | 13.4% | 2.15% | 
| **32** | **5.2%** | 2.79% | 

With only 45 attention writers (34 of them BF16) that leakage is lost in the noise. Across **12,442 FP8
matrices** it accumulates: the 8-iteration bake landed at 0.188 harmful refusal against a lossless
prediction of 0.031. Raising it to 32 closed roughly half the gap — **0.188 → 0.094**. Beyond 32 the
leakage falls slowly while the weight perturbation keeps rising, so 32 is the shipped operating point.

- The checkpoint is a drop-in for `zai-org/GLM-5.3-Flash` in the`transformers` /`Glm5NextForConditionalGeneration` stack (needs`transformers 5.16+` ).
- **Re-quantization:** derive GGUF / MLX / lower-bit builds from these weights directly.
- **Note:** abliteration is a*weight edit* , not data-level unlearning. Fine-tuning on refusal-heavy or
safety data can partially re-introduce refusals; neutral / task data preserves the uncensored behaviour.

- Research into refusal mechanisms, alignment, and interpretability.
- Red-teaming and safety / robustness evaluation in controlled environments.
- A base for further post-training and quantization.

Measured on **these exact weights**, 8×H100 with tensor + expert parallelism, greedy decoding,
`reasoning_effort=low`, the `<think>` block stripped before classification. Refusal is judged by a
rule-based opening-phrase classifier (`caveat` = answered but wrapped in a disclaimer) — indicative, not
an LLM-judge / publication-grade number.

| Benchmark | n | Base | **This model** | Δ | Caveat (this) | 
|---|---|---|---|---|---|
| MaliciousInstruct | 100 | 0.960 | **0.110** | −89% | 0.500 | 
| JailbreakBench (harmful) | 100 | 0.930 | **0.120** | −87% | 0.480 | 
| ForbiddenQuestions | 150 | 0.593 | **0.120** | −80% | 0.400 | 
| AdvBench | 100 | 0.970 | **0.150** | −85% | 0.650 | 
| HarmBench (standard) | 150 | 0.933 | **0.180** | −81% | 0.420 | 
| StrongREJECT | 150 | 0.993 | **0.273** | −72% | 0.633 | 
| SimpleSafetyTests | 50 | 0.920 | **0.340** | −63% | 0.620 | 

Both columns are measured on these exact checkpoints with the same script, sampling seed and settings.
The base model refuses 92–99% of harmful prompts on six of the seven sets — `ForbiddenQuestions` is the
outlier at 0.593, which says more about that set (a large share of its questions are sensitive rather
than harmful, and the base model answers them) than about the edit.

The spread in Δ is the same result reported in *What resisted* above, seen from the benchmark side: the
sets that fall furthest are the ones whose content the removed direction mediates, and the two that fall
least are the ones it does not reach.

| Benchmark | n | Base | **This model** | 
|---|---|---|---|
| XSTest-safe | 250 | 0.024 | **0.004** | 

250 prompts written to *look* harmful while being benign. The base model is already restrained here at
2.4%; the edit does not blunt that discrimination, it sharpens it — 0.4%, and over-refusal on the
held-out benign split fell from 0.094 to 0.000.

| Benchmark | n | Base | **This model** | Δ | 
|---|---|---|---|---|
| MMLU | 300 | 0.833 | **0.827** | −0.7 pp | 
| MMLU-Pro | 400 | 0.438 | **0.453** | +1.5 pp | 
| GSM8K (CoT) | 150 | 0.947 | **0.940** | −0.7 pp | 
| CMMLU (Chinese) | 500 | 0.854 | **0.860** | +0.6 pp | 

Every delta is within ±1.5 pp and no reply failed to parse on either checkpoint. Two of the four move
*up*, which is the signature of sampling noise at these sizes rather than of an improvement — read the
whole table as "unchanged", not as a gain.

Multiple choice is scored from a single forward pass, taking the argmax over the option letters' logits,
rather than by generating an answer; GSM8K keeps chain-of-thought generation because it needs the
reasoning. That makes the MMLU-Pro column lower than a CoT-scored run of the same benchmark would be —
**for both checkpoints equally**, which is why the base column is not optional. Read the delta, not the
absolute figure, and do not compare these numbers against MMLU-Pro results obtained with a different
scoring method.

`glm5_next` needs a vLLM build with GLM-5.3-Flash support and `transformers 5.16+`. The weights are
306 GiB, so plan for 8×H100/H200 with tensor parallelism.

```
docker run -d --name glm53 --gpus all --ipc host -p 8000:8000 \
  -v /path/to/GLM-5.3-Flash-Uncensored-FP8:/model \
  vllm/vllm-openai:latest \
  --model /model --served-model-name GLM-5.3-Flash-Uncensored-FP8 \
  --tensor-parallel-size 8 --max-model-len 262144 \
  --enable-expert-parallel
```
See the vLLM recipe and the
SGLang cookbook for the currently
recommended flags, including the tool-call parser to pass with `--enable-auto-tool-choice`.

GLM-5.3-Flash has **no `enable_thinking` toggle** — its chat template always opens a `<think>` block.
Control the budget with `reasoning_effort` instead (`low` / `high` / `max`; defaults to `max`):

```
client.chat.completions.create(
    model="GLM-5.3-Flash-Uncensored-FP8",
    messages=[{"role": "user", "content": "..."}],
    extra_body={"chat_template_kwargs": {"reasoning_effort": "low", "clear_thinking": True}},
)
```
`clear_thinking` defaults to `false`; pass `true` for chat scenarios. Give generation enough budget to
reach `</think>`, or replies get truncated inside the scratchpad. Pass `image_url` content parts for
vision.

- **Safety guardrails removed** — the model will produce harmful, biased, or offensive content on request
(see the disclaimer).
- It inherits any biases and limitations of the base `GLM-5.3-Flash` .
- Roughly 5% of the refusal direction survives requantization into `e4m3` (see above), so a small
residual refusal rate remains — it is not, and cannot be, exactly zero on an FP8 checkpoint.
- **Refusal is reduced, not removed.** Some content categories are not mediated by the direction this
method removes and still refuse at close to the base rate (see*What resisted* ). Do not assume a
uniformly uncensored model.
- Capability retention is measured, not assumed (see **Evaluation** ), but on sampled subsets of four
benchmarks — enough to rule out a large regression, not a substitute for a full harness run.
- The reported refusal metric is a rule-based heuristic; evaluate rigorously for your own use case.

**MIT**, inherited from the base model
`zai-org/GLM-5.3-Flash`. Abliteration does not change the
underlying license obligations.

- Downloads last month
- 127,772
