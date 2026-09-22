---
id: briefing-general-tech-2026/04-training-inference-quantization/07-quantization-economics
title: "Quantization: the economics (McKinsey, June 2026)"
domain: training-inference-quantization
role: deep-dive
task: quantization
actors: []
dates: ["2026-06"]
keywords: ["quantization", "cost per token", "decode", "disaggregated", "distillation", "fp4", "fp8", "inference", "int4", "pruning"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g05-7"
source_lines: [4794, 4873]
canonical_for: ["quantization"]
sha256: 922353ca257ab97f9e3bf797edb2050925d806b4778749653d339ac07ca2a050
---

# Quantization: the economics (McKinsey, June 2026)

<a id="g05-7"></a>
### 5.7 Quantization: the economics (McKinsey, June 2026)

The authoritative economic framing of the quantization wave came from
McKinsey's June 2026 report *"The technology shifts reducing AI inference
costs,"* part of the firm's Frontiers of compute series. Read directly on
mckinsey.com — and this dossier records the verification explicitly, because
an early suspicion that the headline figure might have been a hallucination
was cleared by direct reading of the report text — the report's central claim
is striking enough to quote verbatim: **"model optimization reduced cost per
token by 85 to 95 percent…"**

Eighty-five to ninety-five percent. Not a marginal efficiency gain but an
order-of-magnitude-level reduction, attributed to model optimization
techniques with quantization — weight and activation precision reduction — and
pruning as the named components. It is the single most dramatic number in the
2026 inference-economics literature, and it reframes the entire hardware
conversation: the largest lever on per-token cost in the last two years was
not cheaper chips, not cheaper power, not better datacenters, but making
models smaller in bits.

The mechanism is simple enough to state precisely, because it falls directly
out of the memory-bound nature of inference established in Section 5.1. Every
decode step streams the model's weights from memory to the compute units; the
time and energy cost of that step is dominated by memory traffic, not
arithmetic. Halving the bits per weight therefore halves the memory footprint
*and* the memory traffic per token — and since memory traffic is the binding
constraint, the throughput and cost follow the bits down. The report's
headline performance figure for the step from FP16 to INT4 is a 2–4×
throughput gain, consistent with the arithmetic: one quarter of the bits per
weight means roughly four times the effective memory bandwidth for the same
hardware, minus the overhead of dequantizing on the fly (the overhead the
Marlin-class kernels of Section 5.10 exist to minimize).

The 85–95% figure needs honest contextualization, and this is the most common
misreading of the report to guard against. The figure describes what model
optimization achieved over its measured window — a period in which the
baseline moved from unoptimized FP16/FP32 serving, through early quantization,
to the quantized, pruned, distilled pipelines of 2026. It is a *cumulative,
bundle-level* number: quantization plus pruning plus the distillation and
architectural efficiencies folded into "model optimization." It is not what a
single INT4 conversion delivers on top of an already-optimized 2025 stack —
that delta is closer to the 2–4× throughput figure, or the −62% single-
configuration datapoint of Section 5.6. Conflating the cumulative figure with
the marginal one produces either euphoria ("quantization cuts 90%!") or
cynicism ("the 90% is a lie") — both wrong. Read correctly, the report says:
the industry's per-token cost fell by an order of magnitude as optimization
matured, and quantization was the dominant term in that maturation.

That reading is what makes the rest of this chapter a taxonomy of the dominant
term. Sections 5.8 through 5.14 walk the bit ladder from the safe FP8 tier
down through INT4 and the FP4 family to the 1.58-bit frontier and the 3-bit KV
cache — each rung a concrete technique, each with verified dates, measured
quality costs, and named hardware support. The McKinsey number is the *why*;
the formats are the *how*.

| McKinsey "technology shifts reducing AI inference costs" (June 2026) | Detail |
|---|---|
| Series | Frontiers of compute |
| Headline quote (verified by direct reading) | "model optimization reduced cost per token by 85 to 95 percent…" |
| Components named | Quantization (weight + activation precision reduction), pruning |
| FP16→INT4 throughput | 2–4× |
| Hallucination check | Early suspicion cleared — the figure is real, read on mckinsey.com |
| Correct reading | Cumulative, bundle-level figure over the optimization window — not the marginal gain of one INT4 conversion on an optimized stack |

It is worth unpacking what "model optimization" bundles in the McKinsey
accounting, because the bundle has three terms and only one of them is this
chapter's subject. Quantization — reducing the bits per weight and activation
— is the dominant term and the one with the cleanest economics: fewer bits,
less memory traffic, proportional throughput. Pruning — removing weights or
structures entirely — is the second named term: it reduces the work rather
than the precision of the work, and composes with quantization (a pruned,
quantized model is smaller twice over). The third term, folded into the
report's framing though not named in the headline quote, is the ambient
improvement of the serving stack itself: better batching, better kernels,
disaggregated architectures — the software compounding that Sections 5.3 and
5.14 describe. The 85–95% is the bundle's cumulative effect; this chapter's
claim is only that quantization is the largest term in it, which the FP16→INT4
2–4× figure and the production format map both support.

