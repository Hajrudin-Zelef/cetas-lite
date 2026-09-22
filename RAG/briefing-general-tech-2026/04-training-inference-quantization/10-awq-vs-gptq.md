---
id: briefing-general-tech-2026/04-training-inference-quantization/10-awq-vs-gptq
title: "AWQ vs GPTQ, Marlin kernels"
domain: training-inference-quantization
role: deep-dive
task: quantization
actors: []
dates: ["2026-06-28"]
keywords: ["awq", "gptq", "decode", "gpu", "inference", "int4", "perplexity", "quantization", "training", "vllm"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g05-10"
source_lines: [5033, 5127]
sha256: 618b563a64c45825f61395d8b9d694e27ab4260bc85c727ef05296e7b8fc50b2
---

# AWQ vs GPTQ, Marlin kernels

<a id="g05-10"></a>
### 5.10 AWQ vs GPTQ, Marlin kernels

Beneath the format tiers sits the method layer: *how* the weights actually get
to 4 bits. A format is a container; a quantization method is the algorithm
that decides which information survives the compression. Two post-training
quantization methods dominated the conversation through 2026 — GPTQ, the 2023
pioneer, and AWQ (Activation-aware Weight Quantization) — and the year's
verdict was decisive: AWQ supplanted GPTQ for new releases (Digital Applied,
June 28, 2026).

The technical difference is worth understanding because it explains why the
verdict went the way it did. GPTQ quantizes layer by layer, minimizing the
reconstruction error of each layer's outputs — a sound, general approach that
treats all weights as equally worthy of protection. AWQ starts from a
different observation: not all weights matter equally, and the ones that
matter can be identified *through the activations*. A small fraction of
weights — roughly 1%, the "salient" weights — are disproportionately
responsible for the activation outliers that propagate through the network;
quantize those carelessly and the error compounds through every subsequent
layer. AWQ's move is selective protection: identify the ~1% of salient weights
via activation statistics, keep them at higher precision, and compress the
remaining 99% aggressively. Protect the 1% that matters; compress the 99% that
doesn't. The method's name is the method: activation-aware.

The quality figures are the industry's rule of thumb, and the dossier presents
them as exactly that — a rule of thumb for budgeting, not a guarantee per
model. AWQ typically costs 0.5–1.5% in perplexity degradation; INT4
quantization generally lands in the 2–5% range depending on model and task;
overall, practitioners budget 0.5–2 points of quality loss for a well-executed
4-bit conversion. (Perplexity here is the standard language-modeling quality
metric — how surprised the model is by held-out text — and a 1% degradation
is, for most production uses, invisible against the noise of prompt
engineering and sampling.) The speedup side: roughly 3× AWQ inference speedup
versus FP16 is the figure in circulation — to be contextualized, because it
depends on batch size, sequence length, and whether the workload is memory-
bound enough for the bandwidth saving to translate fully into throughput. A
single-stream latency-sensitive deployment will see less than 3×; a batched
throughput deployment on memory-bound hardware will see it or more.

Beneath the method layer sits the kernel layer, and 2026 confirmed that the
kernel layer is where theoretical savings become realized tokens per second.
Marlin — and its successor Machete — are hand-tuned GPU kernels in vLLM that
execute 4-bit matrix multiplication on Ampere and Hopper hardware: they fuse
dequantization with the multiply-accumulate, keep the working set in the
fastest memory tiers, and eliminate the overhead that would otherwise eat the
bandwidth saving. Quantization without a matching kernel is a smaller file
that runs no faster; Marlin is the reason AWQ's 4-bit weights actually decode
faster rather than merely occupying less VRAM. The kernel-method-format stack
is the full picture: AWQ decides *what* survives, INT4 is the *container*,
Marlin is the *execution*.

AWQ's breadth of runtime support confirms the verdict: beyond vLLM, AWQ is
supported in LMDeploy and in Hugging Face's TGI (Text Generation Inference),
making it the most broadly served 4-bit method in the 2026 landscape. GPTQ
remains supported everywhere — no deployed pipeline was ripped out — but new
releases, new model ports, and new serving deployments chose AWQ through 2026.

The lineage deserves a sentence for context, because the method layer has a
settled history beneath the 2026 contest. The 2023–2024 foundation is
confirmed and uncontroversial: SmoothQuant's handling of activation outliers
via per-channel scaling migration, GPTQ's layer-wise quantization, AWQ's
activation-aware salient-weight protection, QuaRot's structured rotation
transforms that make weights and activations jointly quantization-friendly.
The 2025–2026 refinements in this lineage (SpinQuant and its successors) were
not verified for this dossier and are deliberately left out — the dossier's
method is to assert only what was checked. What matters for 2026 is the
outcome, not the full genealogy: AWQ won the method contest for new releases,
Marlin-class kernels made the wins real on Ampere and Hopper, and the method
layer is now stable enough that operators choose formats and set quality gates
rather than auditioning algorithms.

| Method / kernel | Role | 2026 status |
|---|---|---|
| GPTQ | Layer-wise PTQ pioneer (2023); minimizes per-layer reconstruction error | Supplanted by AWQ for new releases; still supported everywhere |
| AWQ | Activation-aware: protects ~1% salient weights (found via activation statistics), rest to 4-bit | Dominant method for new releases (Digital Applied, 28/06/2026); vLLM, LMDeploy, TGI |
| Marlin / Machete | Hand-tuned 4-bit kernels in vLLM (fused dequant + matmul) | Ampere/Hopper; where theoretical bandwidth savings become real throughput |
| Quality rule of thumb | — | 0.5–2 pts loss typical for well-executed 4-bit; AWQ 0.5–1.5% perplexity; INT4 generally 2–5% by model/task |
| Speedup figure | — | ~3× AWQ vs FP16 in circulation — context-dependent (batch, sequence length, memory-bound-ness) |
| Lineage (settled) | SmoothQuant, GPTQ, AWQ, QuaRot (2023–2024) | Confirmed; 2025–26 refinements (SpinQuant…) not verified here |

The kernel deserves its "last mile" reputation for a concrete reason:
dequantization overhead. A 4-bit weight can't be multiplied directly by most
hardware — it must be unpacked to a wider type first, and a naive
implementation pays that unpacking cost on every single multiply-accumulate,
erasing the bandwidth saving that justified quantization in the first place.
Marlin-class kernels fuse the unpacking with the matrix multiplication:
weights stream from memory in 4-bit form (the bandwidth win is preserved), get
expanded in registers or shared memory at the last possible moment, and
accumulate in full precision. The fusion is the entire trick, and it's why the
method-format-kernel stack has to be evaluated as a unit: AWQ's salient-weight
protection is only as good as INT4's container, and INT4's container is only
as fast as Marlin's fused execution. In 2026, operators who benchmarked
formats without the matching kernels measured the container, not the system.

