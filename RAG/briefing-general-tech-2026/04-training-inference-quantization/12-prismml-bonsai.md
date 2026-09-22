---
id: briefing-general-tech-2026/04-training-inference-quantization/12-prismml-bonsai
title: "PrismML Bonsai: 27B models in a few gigabytes"
domain: training-inference-quantization
role: deep-dive
task: quantization
actors: ["Apple", "BitNet", "Intel", "PrismML"]
dates: ["2026-07-14", "2026-07-16", "2026-08-03"]
keywords: ["awq", "benchmark", "bitnet", "decode", "gpu", "inference", "nvfp4", "training"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g05-12"
source_lines: [5210, 5294]
sha256: d8822ba298eca69c2be518e42adf5909c1fd9f56139259d6d4b0bdeed7ad97d6
---

# PrismML Bonsai: 27B models in a few gigabytes

<a id="g05-12"></a>
### 5.12 PrismML Bonsai: 27B models in a few gigabytes

BitNet proved that 1.58-bit training was possible as a research program;
PrismML's Bonsai, launched July 14, 2026, asked the commercial question — what
if you take an existing frontier-class open model and compress it to ternary
*after the fact*, no native ternary training required? The artifact it
produced is the most striking compression datapoint of 2026: a 1-bit/ternary
compression of Qwen3.6 27B in which the 1-bit variant weighs 3.9 GB — 1.125
bits per weight — and the ternary variant 5.9 GB. Restate that slowly, because
the scale is what matters: a 27-billion-parameter model, of the class that
normally requires a datacenter GPU's memory just to load, in a file smaller
than a 4K movie. The release was Apache 2.0 licensed, with GGUF, MLX and AWQ
builds — the three distribution formats of the open-weights ecosystem,
covering CPU inference, Apple silicon, and the AWQ serving path respectively —
and PrismML claimed roughly 11 tokens per second on an iPhone: a phone running
a 27B-class model locally, no server round-trip.

The numbers deserve the dossier's standard honesty treatment, applied
carefully because this is exactly the kind of release where enthusiasm outruns
verification. What is confirmed: the existence of the release, dated July 14,
2026 — and the dossier records that an earlier suspicion of hallucination
around Bonsai was cleared by verification; the launch is real. The orders of
magnitude are confirmed: 3.9 GB and 5.9 GB for 27 billion parameters are
arithmetic facts about the released files, not marketing claims. What should
be handled with care: the fine benchmark figures — quality retention at 1.125
bits per weight is the entire ballgame, and secondary coverage does not
establish it — and the ~11 tokens/s iPhone claim, which is a vendor figure on
unspecified test conditions. The coverage itself was mostly secondary and of
uneven quality: Motley Fool (August 3, 2026), Business Standard (July 16,
2026). A CNBC mention of very early Apple talks is unconfirmed and should be
treated as rumor, not signal — "very early talks" is the precise category of
claim that evaporates on contact with verification.

The distinction that matters most is between "the file exists" and "it answers
well." Bonsai is a compression artifact, not a serving benchmark: the release
proves that a 27B model can be *represented* in 3.9 GB, not that the 3.9 GB
model *performs* at 27B quality on any particular eval suite. The gap between
representation and performance is where the careful reader should keep their
skepticism — and where, symmetrically, the release's defenders would point out
that even partial quality retention at these bitrates redraws the deployment
map.

Bonsai's strategic significance does not depend on its benchmarks holding up
perfectly, and this is why the dossier covers it at length despite the
reserves. It is the first commercial artifact to put a 27B-class model inside
phone-scale memory using 1-bit/ternary compression of an *existing* model —
collapsing the distinction between BitNet's research frontier (native ternary
training, Section 5.11) and the product frontier (take a model people already
use, compress it, ship it). If the quality holds at even a fraction of the
implicit claim, the implication for the inference-rack story of Section 5.5 is
direct and large: the "large cheap memory" answer to decode economics gets
radically cheaper when the model itself is a few gigabytes, the edge-
deployment story stops being about small models and starts being about big
ones, and the cost-per-token curve of Section 5.6 takes another step down that
no hardware purchase was needed to unlock. Watch this space with the reserves
noted above — but watch it, because the direction (existing models, extreme
compression, phone-scale deployment) is where the bit ladder points.

| Bonsai datapoint | Detail (with reserve flags) |
|---|---|
| Launch | 14/07/2026 — verified real (earlier hallucination suspicion cleared) |
| What | 1-bit/ternary compression of Qwen3.6 27B (post-training compression of an existing model — not native ternary training) |
| Sizes | 1-bit variant 3.9 GB (1.125 bits/weight); ternary 5.9 GB — orders of magnitude confirmed as file facts |
| License / builds | Apache 2.0; GGUF, MLX, AWQ builds (CPU, Apple silicon, AWQ serving path) |
| iPhone claim | ~11 tokens/s on iPhone — vendor claim, unspecified conditions, handle with care |
| Coverage | Motley Fool 3/08/2026, Business Standard 16/07/2026 — mostly secondary, uneven quality |
| Apple talks | CNBC mention of "very early" talks — unconfirmed, treat as rumor |
| Core reserve | Compression artifact, not serving benchmark: representation proven, quality retention at 1.125 bits/weight not established by verified sources |

The right way to think about Bonsai's reserves is as an evaluation gap with a
known shape. Representation — "the 27B model fits in 3.9 GB" — is verified;
performance — "the 3.9 GB model answers like a 27B model" — is what would need
a public eval suite (MMLU-class knowledge, long-context retrieval, instruction
following, ideally the same harnesses the base Qwen3.6 27B was measured on) to
establish. Until that suite is published and reproduced, the rational stance
is the dossier's: the compression is real and remarkable, the quality
retention is unproven. Note what would *not* close the gap: more secondary
coverage, more vendor demos, more social-media benchmarks. Only reproducible
evals on standard harnesses. This is the general discipline the chapter
applies everywhere — the Intel KLD measurement counts because it's a
measurement on a named model; the "~95% NVFP4 retention" doesn't quite count
because it's a single table — applied here to the most exciting artifact in
the section.

