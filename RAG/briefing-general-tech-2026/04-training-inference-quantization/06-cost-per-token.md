---
id: briefing-general-tech-2026/04-training-inference-quantization/06-cost-per-token
title: "Cost per token as the decisive metric"
domain: training-inference-quantization
role: deep-dive
task: economics
actors: ["Amazon", "MLCommons"]
dates: ["2025-11"]
keywords: ["cost per token", "agentic", "decode", "disaggregated", "gptq", "gpu", "gpus", "inference", "int4", "llama", "lpddr", "mlperf"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g05-6"
source_lines: [4717, 4793]
canonical_for: ["cost-per-token"]
sha256: 4563c3fe3eebcc5023395a64e76185501e50af6ca62a298c8eafc8515e5ec0fd
---

# Cost per token as the decisive metric

<a id="g05-6"></a>
### 5.6 Cost per token as the decisive metric

Every architectural choice in this chapter — disaggregation, LPDDR pools,
telco deployment, the quantization techniques that fill the second half — is
ultimately arbitrated by one number: cost per token. In 2026 this metric
completed its move from engineering dashboards to the top of operator budgets,
and the reason is compositional: the makeup of AI spending flipped, and the
flip made per-token cost the only number that matters at scale.

McKinsey's AI Economics Report (2025, cited via DataMites) puts inference at
65–90% of total LLM operating cost. The intuition behind the range is
straightforward. Training is a one-time capital event: expensive, spectacular,
but bounded — you pay it once per model generation and then it is sunk.
Inference is a permanent operating line: every query, every agent loop, every
token generated for every user, forever, metered against the fleet. At the
usage scales of 2026 — hundreds of millions of users, agentic workloads that
multiply tokens per task by ten or a hundred — the operating line swallows the
capital event. When two-thirds to nine-tenths of a model's lifetime cost is
incurred one token at a time, the only hardware question worth asking is how
cheaply each token can be produced, and every other metric — peak FLOPS,
memory capacity, interconnect bandwidth — is instrumental to that one.

That is why the apparently disparate stories of this chapter's first half all
point the same direction. MLPerf's scale-out records (Section 5.2) measure
tokens per second at cluster scale. Dynamo's disaggregation (Section 5.3) lets
each phase be priced independently. The LPDDR bets (Section 5.5) attack the
memory-cost term of the token equation directly. None of these would matter
much in a world where training dominated spend; in a world where inference is
65–90% of operating cost, each of them is a margin story.

A concrete, independently measured datapoint makes the economics tangible at
the level of a single deployment decision. On AWS pricing (November 2025),
serving Llama-2-13B on an A10 GPU cost $8.12 per million tokens in FP16 — and
$3.10 per million tokens in INT4 with GPTQ quantization. Same model, same
cloud, same GPU, one technique change: a 62% reduction in per-token cost. That
is the kind of margin that re-prices an entire product line — the difference
between a feature that loses money at scale and one that prints it — and it
came from compressing weights from 16 bits to 4, not from a new chip, a new
datacenter, or a new model.

The honest boundaries on that datapoint are worth stating plainly, because the
dossier's method requires it. It is a single configuration — one model
(Llama-2-13B), one GPU (A10), one cloud (AWS), one date (November 2025), one
method (GPTQ). It should not be generalized into "quantization saves 62%" as a
universal law; different models, newer GPUs with native low-precision
acceleration, and already-optimized baselines will show different deltas. But
its direction is confirmed by every larger study, including the McKinsey
analysis of the next section: model optimization, dominated by quantization,
is the single largest lever on per-token cost available to operators in 2026.
And the agentic multiplier sharpens the point — when a task goes from one
model call to fifty, the per-token cost is multiplied fifty times while the
training cost stays sunk, which is why inference's share of spend keeps
climbing and why the quantization chapters below are, economically, the most
important in this dossier's infrastructure coverage.

| Cost datapoint | Source | Figure |
|---|---|---|
| Inference share of total LLM operating cost | McKinsey AI Economics Report 2025 (via DataMites) | 65–90% |
| Llama-2-13B on A10, FP16 | AWS pricing, Nov 2025 | $8.12 / M tokens |
| Llama-2-13B on A10, INT4 GPTQ | AWS pricing, Nov 2025 | $3.10 / M tokens (−62%) |
| Scope note | — | Single configuration; direction confirmed by larger studies, magnitude not universal |

A short analytical note, because the dossier's readers will encounter both
figures in the wild: throughput is not cost, though they rhyme. The 2–4×
throughput gain from FP16→INT4 measures tokens per second on fixed hardware;
the −62% per-token cost figure measures dollars per million tokens on a cloud
bill. Throughput becomes cost only through utilization — a 4× faster fleet
serving the same demand at the same utilization is 4× cheaper per token, but a
fleet that absorbs the throughput as headroom for traffic spikes has bought
latency insurance, not cost reduction. Operators in 2026 learned to manage
both dials: quantize for throughput, then decide — via autoscaling, bin-
packing and disaggregated decode fleets — how much of the throughput to
convert into margin and how much to hold as quality-of-service. The McKinsey
85–95% is a cost figure; the 2–4× is a throughput figure; the industry's job
is the conversion between them.

