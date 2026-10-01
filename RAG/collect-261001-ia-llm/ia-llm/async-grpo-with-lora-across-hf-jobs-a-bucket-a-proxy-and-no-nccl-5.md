---
id: collect-261001-ia-llm/ia-llm/async-grpo-with-lora-across-hf-jobs-a-bucket-a-proxy-and-no-nccl-5
title: "every Job gets the same bucket at the same absolute path"
domain: ia-llm
role: reference
task: reference
actors: ["Hugging Face", "vLLM"]
dates: ["2025-09"]
keywords: ["cost", "cost per token", "fine-tuning", "gpus", "inference", "lora", "throughput", "training", "vllm"]
source: docs/RAG/collect-261001-ia-llm/async-grpo-with-lora-across-hf-jobs-a-bucket-a-proxy-and-no-nccl.md
source_anchor: ""
source_lines: [259, 314]
sha256: 225df652f45331d44d48019530242cb85650e599372336a26afe85a64da05068
---

# every Job gets the same bucket at the same absolute path

*Figure 5. trackio runs `r1-dp2-tb16k-nockpt` and `r1-dp3-tb16k-nockpt` overlaid over their first 134 steps. Panels: `perf/weight_sync_s`, `rollout/generated_tok_s`, `rollout/inflight`, `perf/fwd_bwd_s`. Sync falls from 7.6 s to 5.8 s; generation and forward+backward do not move; `rollout/inflight` reads 128 in both runs, which is the cap the third replica ran into.*

Weight sync falls from 7.6 s to 5.8 s, so the shorter retry helps. Forward and backward stay at 4.6 s, which rules out resharding. Generation moves from 25k to only 26k tokens/s. The third replica does basically nothing.

The reason was sitting in `rollout/inflight`: 128 in every run. The proxy shows those requests split as 44 + 43 + 41 across the three replicas. `max_inflight_tasks` limits concurrency for the whole rollout worker, not per replica. A 1.5B model on an H200 processes 43 and 130 concurrent sequences at almost the same cost per token. Splitting 128 requests over three GPUs gives nearly the same throughput as splitting them over two.

So vLLM was not the limit. Our own client-side constant was. We had set it conservatively because we did not know how hundreds of long HTTPS requests would behave through the public Jobs proxy. At this point, 130,000 rollout completions had crossed it without a single transport error.

`max_inflight_tasks=384` and `queue_maxsize=768`, nothing else.

*Figure 6. All five trackio runs (`r1-dp2`, `r1-dp2-tb16k`, `r1-dp2-tb16k-nockpt`, `r1-dp3-tb16k-nockpt`, `r1-dp3-inflight384`) overlaid, x-axis in steps. Panels: `perf/step_s`, `reward`, `sample/rollout_queue_size`, `sample/staleness_mean`. Step time falls from 22.9 s to 4.8 s across the series while the reward curves stay on top of each other; the last run's queue refills to ~690 of 768 and its staleness settles at 2. Runs 2 to 4 were stopped early once the dashboard had answered the question.*

With 384 requests in flight, each replica gets 128. The queue quickly fills to around 690 out of 768 and stays there. Backpressure returns to 5 seconds and rollout wait falls to 0.03 seconds. Training is the bottleneck again! Forward and backward take 4.6 seconds, weight sync adds an amortized 1.5 seconds, and median step time is 4.8 seconds.

Mean staleness rises from 1.5 to 2.0 versions because samples wait longer in the larger queue. This is still below `max_staleness=4`, and `ratio` remains very close to 1.000.

| 500 steps | run 1 `r1-dp2` | run 5 `r1-dp3-inflight384` | 
|---|---|---|
| wall clock | 3 h 27 min | **53 min** | 
| `perf/step_s` , p50 | 22.9 s | 4.8 s | 
| `perf/fwd_bwd_s` , p50 | 21.9 s | 4.6 s | 
| `perf/mfu_fwd_bwd` | 3.9 % | 23.5 % | 
| `batch/samples_per_step` | 128 | 168 | 
| samples trained | 64 000 | 84 078 | 
| `perf/weight_sync_s` , p50 | 8.5 s | 6.2 s | 
| `sample/staleness_mean` | 1.5 | 2.0 | 
| reward, first 20 → last 20 steps | 0.145 → 0.438 | 0.145 → 0.416 | 

*Figure 7. trackio runs `r1-dp2` and `r1-dp3-inflight384`, reward against wall-clock minutes since the first optimizer step. Same recipe, same 500 steps, same final reward; run 5 gets there in 52 minutes instead of 3 h 26 min.*

The final run is 3.9× faster and trains on 31 % more samples, with basically the same reward curve. Packing, disabling checkpointing and raising the in-flight limit made the difference. In each case, the dashboard made this clear within the first ten minutes.

```
git clone https://github.com/AmineDiro/hfjobs-lora-buckets && cd hfjobs-lora-buckets
hf auth login
MAX_STEPS=20 RUN_TAG=smoke ./run_all.sh --wait        # ~15 min, three Jobs, cancels the servers when done
MAX_STEPS=500 ./run_all.sh --wait                     # run 1: the reference batch shape, ~3.5 h
TOKEN_BUDGET=16384 GRAD_ACCUM=6 GRADIENT_CHECKPOINTING=0 PROXY_LORA_RETRY_S=0.5 \
  MAX_INFLIGHT=384 QUEUE_MAXSIZE=768 MAX_STEPS=500 ./run_all.sh --wait   # run 5: same recipe, ~55 min
```
- John Schulman et al., LoRA Without Regret, Thinking Machines Lab, September 2025. The case that rank-1 LoRA matches full fine-tuning for policy-gradient RL, and why.
- TRL, `AsyncGRPOTrainer` and its logged metrics.
- TRL PR #7017: PEFT/LoRA support for `AsyncGRPOTrainer` with adapter-only vLLM sync, released in TRL v1.14.
- Hugging Face Jobs and Storage Buckets; `hf-mount` .
- `hf-mount-repro` : the two-script reproduction of the 30-second negative-cache stall.
- The trackio dashboard for every run in this post.
- Penghui Qi, Zichen Liu, Xiangxin Zhou, Tianyu Pang, Chao Du, Wee Sun Lee, Min Lin, Defeating the Training-Inference Mismatch via FP16, arXiv:2510.26788, 2025. Source of the Sanity dataset `sail/Sanity-Test-R1D-1.5B` and of the LoRA recipe,`sail-sg/Precision-RL` ,`oat/scripts/lora/bf16_grpo_tis_lora.sh` .

```
@article{qi2025precisionrl,
  title={Defeating the Training-Inference Mismatch via FP16},
  author={Qi, Penghui and Liu, Zichen and Zhou, Xiangxin and Pang, Tianyu and Du, Chao and Lee, Wee Sun and Lin, Min},
  journal={arXiv preprint arXiv:2510.26788},
  year={2025}
}
```
