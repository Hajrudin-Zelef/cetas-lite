---
id: briefing-general-tech-2026/04-training-inference-quantization/13-turboquant
title: "TurboQuant: 3-bit KV cache"
domain: training-inference-quantization
role: deep-dive
task: quantization
actors: ["BitNet", "Google", "PrismML"]
dates: ["2025-04"]
keywords: ["kv cache", "turboquant", "accelerator", "agentic", "bitnet", "decode", "inference", "lpddr", "perplexity", "quantization", "sglang", "training"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g05-13"
source_lines: [5295, 5400]
canonical_for: ["quantization"]
sha256: 681b4a81cdc28fcf7030cb428e8a9679048caf540c8b78370cd7fe5a66de39ee
---

# TurboQuant: 3-bit KV cache

<a id="g05-13"></a>
### 5.13 TurboQuant: 3-bit KV cache

Weights are only half the memory story in inference. The other half is the KV
cache — and in the long-context, agentic workloads that defined 2026 demand,
it is often the *binding* constraint. The mechanism is worth explaining
because it is the reason an entire second quantization frontier opened this
year. Every token a model generates requires re-attending to every previous
token; to avoid recomputing the whole history at each step, the serving stack
caches the key and value tensors for all past positions — the KV cache. The
cache grows linearly with context length, with layer count, and with batch
size. At the 128K–1M context lengths that agentic workflows made routine in
2026, a single long conversation's KV cache can exceed the model's own weight
footprint: you quantized the weights to 4 bits and the cache still eats your
memory budget at 16. Weight quantization without KV-cache quantization is half
a solution, and 2026 was the year the industry started solving the other half.

Google's TurboQuant is the flagship of that second frontier. The paper is real
and the trail is verified: *"TurboQuant: Online Vector Quantization with Near-
optimal Distortion Rate,"* arXiv 2504.19874 (April 2025), presented at ICLR
2026. The dossier records the clearance explicitly — an earlier hallucination
suspicion was checked and cleared — and adds the disambiguation that the
literature requires: TurboQuant is not to be confused with the QuaRot 4-bit
line or the KVQuant/KIVI lineage. It is its own method (online vector
quantization with near-optimal distortion-rate guarantees), and KIVI — the
prior KV-cache quantization baseline — is cited in the paper as a *beaten*
baseline, which positions TurboQuant as a generational step rather than an
incremental tweak.

The figures, from the Google Research blog: a 3-bit KV cache delivering at
least 6× KV memory reduction, and — for the 4-bit variant versus a 32-bit
baseline — up to 8× attention speedup on H100. The benchmarks: LongBench,
Needle-in-Haystack and RULER — the long-context retrieval and reasoning suites
where KV-cache fidelity actually matters, since a cache that corrupts distant
context would fail exactly these tests — run on Gemma and Mistral models. And
crucially for the "production, not paper" test: the method is integrated into
SGLang via the `--kv-cache-dtype turboquant` flag, which makes it a deployable
option in a production serving runtime rather than a result that lives only in
arXiv.

The nuance, as always with headline acceleration figures, is in the bounds —
and the dossier states them because unbounded figures mislead operators sizing
fleets. The ×8 attention speedup and the ≥6× memory reduction hold on specific
configurations: long contexts (where the cache dominates), the H100 (with its
specific memory hierarchy), the benchmarked model family. They are not
universal constants; a short-context workload will see less, a different
accelerator will see different. On quality: the paper's abstract claims
quality-neutrality at 3.5 bits per channel — a precise, bounded statement —
while the "3-bit with zero quality loss" phrasing that circulates in secondary
coverage is the marketing bound of that claim, not its precise form. Read
precisely, TurboQuant says: the KV cache, the half of inference memory that
weight quantization doesn't touch, can be compressed to roughly 3 bits with
negligible quality impact on long-context benchmarks. That is already a strong
claim; it doesn't need the marketing version.

The strategic read connects directly to Section 5.5's inference-rack story and
completes the 2026 inference-economics stack. Disaggregation put decode on
memory-dense hardware; the LPDDR bets made that memory cheap per gigabyte;
TurboQuant shrinks what each gigabyte has to hold by a factor of six. Each
step multiplies the last: cheaper memory × less memory per token of context ×
separately scaled decode fleets = the cost-per-token curve that Section 5.6
measures and Section 5.7 attributes mostly to model optimization. TurboQuant
is the 2026 proof that quantization's second frontier — after weights — is the
KV cache, that the frontier is already behind a runtime flag, and that the bit
ladder still has rungs the production stack hasn't fully climbed.

| TurboQuant datapoint | Detail |
|---|---|
| Paper | "TurboQuant: Online Vector Quantization with Near-optimal Distortion Rate," arXiv 2504.19874 (April 2025); presented at ICLR 2026 |
| What | 3-bit KV cache quantization (online vector quantization, near-optimal distortion rate) |
| Figures (Google Research blog) | ≥6× KV memory reduction (3-bit); ×8 attention on H100 for the 4-bit variant vs 32-bit baseline — on specific configs (long context, H100, benchmarked models), not universal |
| Benchmarks | LongBench, Needle-in-Haystack, RULER — long-context suites where cache fidelity is tested; Gemma and Mistral models |
| Lineage | Own method; KIVI cited as beaten baseline; not QuaRot 4-bit, not KVQuant/KIVI lineage |
| Quality bound | Paper abstract: quality-neutral at 3.5 bits/channel — "3-bit zero loss" is the marketing bound of the precise claim |
| Deployment | Integrated in SGLang (`--kv-cache-dtype turboquant`) — production flag, not just a paper |
| Hallucination check | Earlier suspicion cleared — paper, figures and integration verified |

Why was the KV cache solved second, years after weight quantization matured?
Because it's a harder problem in one specific way: weights are static, the
cache is dynamic. Weight quantization is done once, offline, with the full
model and calibration data available — you can spend hours getting it right.
KV-cache quantization must happen *online*, during inference, as each new
token's keys and values stream in: the quantizer sees each cache entry once,
in order, with no chance to revisit, and must compress it without knowing the
future queries that will attend to it. That's what "online vector
quantization" in TurboQuant's title means, and why "near-optimal distortion
rate" is the claim that matters — it's a guarantee about doing the best
possible under the one-pass constraint. The weights were the easy half of
inference memory; the cache was the half that had to be solved live. That it
was solved to 3 bits with negligible quality loss, inside a single ICLR cycle,
is why this section treats TurboQuant as a frontier result rather than an
incremental one.

Step back and the two frontiers of Sections 5.11–5.13 form a complete picture
of inference memory. The weight frontier (BitNet's 1.58-bit training, Bonsai's
3.9 GB compression of a 27B model) attacks the static half: the parameters,
loaded once, read every step. The cache frontier (TurboQuant's 3-bit KV cache)
attacks the dynamic half: the per-conversation state, written once per token,
read every subsequent step. They compose — a ternary-weight model with a 3-bit
cache is smaller twice over, in both halves — and they fail independently,
which is why the chapter treats them separately: weight quality and cache
fidelity are different eval problems with different benchmarks (perplexity
suites versus LongBench/Needle-in-Haystack/RULER). An operator in 2027 will
hold quality gates on both halves. The bit ladder has two rails; this chapter
climbed both.

