---
id: briefing-general-tech-2026/04-training-inference-quantization/09-int4-nvfp4-mxfp4
title: "INT4, NVFP4, MXFP4: the aggressive standard"
domain: training-inference-quantization
role: deep-dive
task: quantization
actors: ["Intel", "Nvidia", "OpenAI"]
dates: []
keywords: ["fp4", "int4", "mxfp4", "nvfp4", "awq", "benchmark", "blackwell", "decode", "fp8", "gptq", "quantization", "tensorrt"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g05-9"
source_lines: [4942, 5032]
sha256: a62ac178a68d39d58f361a1a84baa9951d949b0bb5d7a659f8817d418862a6e4
---

# INT4, NVFP4, MXFP4: the aggressive standard

<a id="g05-9"></a>
### 5.9 INT4, NVFP4, MXFP4: the aggressive standard

If FP8 is the format nobody questions, 4-bit is the format everybody
benchmarks — and in 2026 it became the aggressive production standard: the
tier operators deploy after the evals pass, where the quality cost is real but
measured, and the per-token economics are transformative. The 4-bit story of
2026 has three dialects — INT4, NVFP4, MXFP4 — and understanding their
differences is understanding the year's serving landscape.

**INT4** — integer 4-bit quantization, with per-group scaling — is the established member of the family, confirmed as a first-class deployment format across the production serving stack: TensorRT-LLM, Torch-TensorRT and vLLM documentation all treat it as standard. The industry's working rule, confirmed by the docs and by practice, is that 4-bit *with grouping* is the best quality-per-bit balance available for nearly all production cases. "Grouping" deserves a sentence of explanation because it is the difference between 4-bit working and 4-bit failing: instead of one scaling factor for an entire weight tensor — which forces tiny weights and huge weights to share the same 16 representable values — the tensor is divided into small groups (typically 32–128 weights), each with its own scale. Outliers get their own range; the bulk of weights get fine resolution around zero. The memory overhead of the per-group scales is negligible; the quality gain over per-tensor scaling is what makes 4-bit production-viable rather than a research curiosity.

The Blackwell generation added two 4-bit floating-point dialects, and their
2026 trajectories diverged in instructive ways. **NVFP4** — Nvidia's 4-bit
floating-point format — is supported in TensorRT-LLM and vLLM but runs on
Blackwell only: a hardware-gated format that ties the aggressive tier to the
newest fleet, and therefore to the operators who have already bought into the
Blackwell generation. **MXFP4** (microscaling FP4) is the open-standards
sibling — and it arrived in 2026 with the strongest possible endorsement:
OpenAI's gpt-oss ships in MXFP4, making a 4-bit microscaling format the native
distribution format of a frontier open-weights release. When a frontier lab
ships its weights in 4-bit, the format conversation is over at the
distribution layer: downstream tooling either supports MXFP4 or it doesn't
serve gpt-oss.

The numbers that justify the tier are straightforward. INT4 cuts weight memory
to one quarter of FP32 — a 75% reduction — and because decode is memory-
traffic-bound, the per-token economics follow the memory down: roughly
proportional throughput gains on memory-bound workloads, subject to the
dequantization overhead that the Marlin-class kernels (Section 5.10) minimize.
But 2026 also produced the honesty nuance that keeps this section from being a
marketing sheet: an Intel cookbook (2026) measured on Qwen3.5-35B that MXFP4
showed roughly 2× worse KLD than GPTQ-INT4 at the same bitwidth. KLD —
Kullback–Leibler divergence between the quantized and full-precision output
distributions — is a direct measure of how much the quantization distorted
what the model actually says, and on that measure, on that model, integer INT4
beat floating-point FP4 by a factor of two. The FP4 family is not uniformly
superior to INT4; which wins is hardware- and model-dependent. "4-bit" is not
one format but a family whose members must be benchmarked per deployment — a
conclusion the dossier states plainly because the alternative (assuming the
newest format wins) is how operators burn eval budgets.

The same caution applies to the "~95% quality retention" figure sometimes
cited for NVFP4: it comes from a single table and should be handled with care,
not quoted as an established constant. The more solid measured datapoint is
the bridge from Section 5.8: DeepSeek-R1-0528 lost at most 1% of quality going
FP8→NVFP4 via post-training quantization — a direct, model-specific
measurement rather than a generalized claim.

The 2026 read, in one paragraph: 4-bit with grouping is the aggressive
production standard, deployed after evals across TensorRT-LLM, Torch-TensorRT
and vLLM; NVFP4 and MXFP4 are rising fast on the Blackwell fleet, with MXFP4
carrying the gpt-oss distribution endorsement and NVFP4 carrying the
Blackwell-only gate; and the choice between FP4-family and integer INT4 is an
empirical question per model and per chip — the Intel KLD datapoint proves the
hierarchy isn't settled — not a matter of picking the newest acronym.

| 4-bit format | Nature | Support / notes |
|---|---|---|
| INT4 (+ grouping) | Integer 4-bit, per-group scales (typically 32–128 weights/group) | TensorRT-LLM, Torch-TensorRT, vLLM docs; best quality-per-bit balance for nearly all production cases |
| NVFP4 | Nvidia 4-bit floating point | TensorRT-LLM, vLLM; **Blackwell only** (hardware-gated); "~95% retention" from a single table — treat with caution; ≤1% loss measured FP8→NVFP4 on DeepSeek-R1-0528 via PTQ |
| MXFP4 | Microscaling FP4 (open standard) | OpenAI's gpt-oss ships in MXFP4 (distribution-layer endorsement); but Intel cookbook (2026, Qwen3.5-35B) measured ~2× worse KLD than GPTQ-INT4 at same bitwidth — FP4 vs INT4 superiority is hardware- and model-dependent |
| Memory arithmetic | — | INT4 = −75% weight memory vs FP32; decode throughput follows memory traffic down (minus dequant overhead) |

Why does grouping matter so much at 4 bits that the dossier calls it "the
whole game"? Consider the alternative. Per-tensor quantization assigns one
scale to millions of weights sharing 16 representable values: a single outlier
weight stretches the range, and every small weight collapses into the same two
or three buckets — catastrophic resolution loss where most of the weights
live. Per-group quantization (groups of 32–128) gives each neighborhood its
own scale: outliers get their own range without contaminating their neighbors,
and the dense cluster of small weights around zero gets fine resolution. The
cost is storing one scale per group — negligible against the weights
themselves. Microscaling (the "M" in MXFP4) pushes the same idea finer: block-
level scales with a shared exponent format, trading a little more metadata for
a little more fidelity. The pattern across INT4-grouped, NVFP4 and MXFP4 is
one idea at three granularities: precision is a local property, and 4 bits are
plenty when each neighborhood sets its own range.

The gpt-oss detail in Section 5.9 carries weight beyond the format itself,
because distribution format is a standards event. When a frontier lab ships
weights natively in MXFP4, every downstream tool — every serving runtime,
every fine-tuning stack, every edge converter — must either support MXFP4 or
not serve the model. That compels the ecosystem faster than any benchmark:
vLLM and TensorRT-LLM support it because they have to, and the GGUF/AWQ/MLX
converters follow. Contrast with NVFP4, which is Blackwell-gated and therefore
a fleet-upgrade question, not an ecosystem question. The 2026 lesson for
format watchers: track what the frontier labs *ship*, not just what the
hardware *supports* — distribution is where standards are actually set, and
gpt-oss set MXFP4's.

