---
id: collect-261001-ia-llm/ia-llm/async-grpo-with-lora-across-hf-jobs-a-bucket-a-proxy-and-no-nccl-4
title: "every Job gets the same bucket at the same absolute path"
domain: ia-llm
role: reference
task: reference
actors: ["vLLM"]
dates: []
keywords: ["compute", "gpu", "gpus", "latency", "lora", "memory", "training", "vllm"]
source: docs/RAG/collect-261001-ia-llm/async-grpo-with-lora-across-hf-jobs-a-bucket-a-proxy-and-no-nccl.md
source_anchor: ""
source_lines: [212, 258]
sha256: 264c5b1dbdd6f630906a4c0c0ed1090a8e45036fdb33d79684157596b5fd2733
---

# every Job gets the same bucket at the same absolute path

*Figure 1. trackio run `r1-dp2`. Panels: `reward` with its 20-step rolling mean and 50-step block means, and `ratio` on a 0.99 to 1.01 axis. Reward climbs from 0.15 to 0.44 over 500 steps; `ratio` stays between 0.9993 and 1.0004 throughout.*

500 steps took 3 h 27 min. Mean reward goes from 0.145 over the first 20 steps to 0.438 over the last 20. More importantly for this test, `ratio` stays at 1.000 for every step! The policy served by vLLM always matches the one used by the trainer to score the rollout. This held across all 126 syncs. Mean staleness was 1.5 policy versions, against a maximum of 4. The trackio dashboard has the full curves.

We have our undeniable proof that LoRA AsyncGRPO works! Let's now see how we can improve our training runs by looking at the recent detailed AsyncGRPO metrics.

Async RL is a pipeline between training and generation. Making one side faster does nothing if the other side cannot keep up. Fortunately, in `AsyncGRPOTrainer` we've added enough timings and metrics to see this directly.

All the useful metrics are documented in the Logged metrics section. `perf/rollout_wait_s` tells us how long the trainer waits for samples. `rollout/backpressure_s` tells us how long generation waits for space in the rollout queue. They are diametrically opposed to each other and should not both be high. Together with the queue size, they tell us which side is slow.

We ran five experiments. Each one starts from a problem visible in the previous run's dashboard. Unless mentioned otherwise, the model, recipe and three-Job layout stay the same. The names below are the trackio run names.

We keep these four groups of metrics visible:

- **`perf/step_s`** and**`perf/fwd_bwd_s`** : how long an optimizer step takes, and how much of it is forward+backward. If a step takes nearly as long as the forward+backward, then the trainer is clearly compute-bound.
- **`perf/rollout_wait_s`** : how long the trainer sat waiting for samples before it could start a step. Near zero means generation is ahead of training. Samples are available immediately to be trained on.
- **`sample/rollout_queue_size`** against`queue_maxsize` : the buffer between the two sides. Full means generation is being throttled; empty means the trainer is starving.
- **`rollout/backpressure_s`** and**`rollout/score_block_s`** : how long the rollout worker sat blocked because that buffer was full. The worker is a two-stage pipeline: generation hands finished groups to a scoring stage, and scoring pushes scored samples into the rollout buffer. When the buffer is full, scoring cannot enqueue and blocks, which is`rollout/backpressure_s` . Scoring then stops draining its own input queue, so generation cannot hand over the next group either, which is`rollout/score_block_s` . Both are the same stall, seen first at the scoring stage and then propagated back to generation.

The diagnosis is simple: a full queue with zero rollout wait and high backpressure means the trainer is too slow. An empty queue with rising rollout wait and no backpressure means generation is too slow. Comparing `perf/mfu_wall_clock` with `perf/mfu_fwd_bwd` also shows how much time the trainer GPUs spend waiting instead of training.

*Figure 2. trackio run `r1-dp2`. Panels: `perf/step_s`, `perf/fwd_bwd_s`, `sample/rollout_queue_size`, `rollout/backpressure_s`. Step time and forward+backward overlap almost completely; the queue sits pinned near 476 of 512 and backpressure never drops below 11 s per rollout group: trainer-bound.*

`perf/step_s` is 22.9 s and `perf/fwd_bwd_s` is 21.9 s. Forward and backward take 96 % of the step time. The queue stays full and the trainer waits only 0.02 s for rollouts, and the rollout worker spends 15 seconds per group blocked by backpressure. The two vLLM replicas generate faster than the trainer consumes. The reported 4.6k tokens/s is not their actual limit; they simply have nowhere to put more output.

The batch metrics explain the terrible 3.9 % MFU. `batch/microbatches_per_step` is 64 and `batch/samples_per_row` is 1.0. Each rank processes one sequence of around 1.2k tokens, 64 times per step. This comes from the reference recipe's `per_device_train_batch_size=1`. For a 1.5B model on an H200, this is completely latency-bound.

The fix is not to change the batch size. We keep 128 completions per optimizer step and only change how they are laid out on the GPU: instead of one sequence per microbatch, we pack many sequences densely into each row. The trainer supports this through **token-budget batching**. With `token_budget > 0`, it packs several samples into one padding-free row per rank. An optimizer step processes `gradient_accumulation_steps` rows. We set `token_budget=16384` and `gradient_accumulation_steps=6`.

*Figure 3. trackio runs `r1-dp2` and `r1-dp2-tb16k` overlaid over their first 154 steps. Panels: `batch/samples_per_row`, `batch/microbatches_per_step`, `perf/step_s`, `perf/fwd_bwd_s`, `perf/mfu_fwd_bwd`, `rollout/generated_tok_s`. Packing takes samples per row from 1 to 13, microbatches from 64 to 6, step time from 23 s to 5.9 s, and generation from 4.2k to 27.5k tok/s with no change on the vLLM side.*

`batch/samples_per_row` goes from 1.0 to ~12.7 and the number of microbatches drops from 64 to 6. The rows are 95 % full! Forward and backward fall from 21.9 s to 5.6 s, while MFU rises from 3.9 % to 19 %. We now train on around 150 samples per step because the rows pack better than the mean-length estimate predicted.

Generation also jumps from 4.6k to 25k tokens/s, even though we changed nothing on the vLLM side. The queue is no longer constantly full, so the replicas can finally run. This is why we do not like optimizing pipeline stages in isolation. We need to be careful to evaluate the whole system, as a slow stage can hide the real performance of everything before it.

`perf/fwd_s` is 1.34 s while `perf/fwd_bwd_s` is 5.6 s. A normal backward costs roughly twice the forward, and with frozen base weights it should be closer to once. A ratio of 3.2 is suspicious.

The reason is that `AsyncGRPOConfig` defaults to `gradient_checkpointing=True`. Every microbatch recomputes its forward during the backward. This also explains why a 16k-token row only uses 25 GB on a 141 GB H200! It's a memory optimization for this trainer, but we don't need it in this specific case: the model is small enough to fit into VRAM with the activations kept for the backward.

*Figure 4. trackio runs `r1-dp2-tb16k` and `r1-dp2-tb16k-nockpt` overlaid over their first 134 steps. Panels: `perf/fwd_s`, `perf/fwd_bwd_s`, `perf/weight_sync_s`, `sample/rollout_queue_size`, `perf/rollout_wait_s`, `perf/mfu_fwd_bwd`. Forward+backward drops by one forward; the queue falls from ~420 to ~60, and rollout wait rises from 0.02 s to 0.5 s: the bottleneck shifts to generation.*

With `gradient_checkpointing=False`, forward and backward drop to 4.6 s, almost exactly one forward less, and MFU reaches 23 %. The queue now falls to 71 and rollout wait rises from 0.04 s to 0.6 s. The trainer consumes samples faster than two replicas generate them. We *successfully* moved the bottleneck to generation.

This exposes two more costs. A 7.6-second weight sync every four steps now takes 25 % of wall-clock time. It was only 8 % when each step took 23 seconds. Also, backward is still 2.5 times slower than forward. With frozen base weights, there are around 2 seconds per step that do not look like normal model math.

Since generation was now too slow, we added a third replica. We also reduced the adapter retry interval from 2 s to 0.5 s in the proxy and disabled `fsdp_reshard_after_forward` to check whether FSDP2 re-gathers caused the extra 2 seconds in backward.

