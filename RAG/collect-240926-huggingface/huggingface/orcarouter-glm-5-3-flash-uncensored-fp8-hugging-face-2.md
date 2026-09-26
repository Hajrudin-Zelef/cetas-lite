---
id: collect-240926-huggingface/huggingface/orcarouter-glm-5-3-flash-uncensored-fp8-hugging-face-2
title: "GLM-5.3-Flash-Uncensored-FP8"
domain: huggingface
role: reference
task: reference
actors: ["China", "SGLang", "Z.ai", "vLLM"]
dates: []
keywords: ["fp8", "glm", "alignment", "attention", "benchmark", "fine-tuning", "gguf", "gpus", "quantization", "reasoning", "refusals", "research"]
source: docs/RAG/clean_en/huggingface/orcarouter-glm-5-3-flash-uncensored-fp8-hugging-face.md
source_anchor: ""
source_lines: [118, 246]
sha256: ab05379b1275df5bc27f176948747fdbd951b72636b0290c6a7805e160e9f62b
---

# GLM-5.3-Flash-Uncensored-FP8

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

