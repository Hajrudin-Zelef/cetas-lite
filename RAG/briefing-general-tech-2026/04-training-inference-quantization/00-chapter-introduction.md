---
id: briefing-general-tech-2026/04-training-inference-quantization/00-chapter-introduction
title: "Training vs inference architectures and quantization — introduction"
domain: training-inference-quantization
role: deep-dive
task: architecture
actors: ["AMD", "BitNet", "Google", "MLCommons", "Nvidia", "PrismML", "UALink"]
dates: ["2026-06", "2026-09", "2026-09-22"]
keywords: ["inference", "quantization", "training", "awq", "benchmark", "bitnet", "blackwell", "cost per token", "decode", "disaggregated", "disaggregated serving", "fp8"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g05"
source_lines: [4229, 4301]
sha256: 6233929b4831187f7946e3d41b34c909ce25a012b409697a91b5dfa85ed6eca9
---

# Training vs inference architectures and quantization — introduction

<a id="g05"></a>
## 5. Training vs inference architectures and quantization

If one structural shift defined AI infrastructure in 2026, it is this:
training and inference stopped being the same business. For most of the
2023–2025 boom, the two workloads ran on roughly the same hardware — Nvidia
H100/H200 racks, identical memory architectures, identical networking,
differing only in how many of them you bought. By September 2026, the stack
has visibly bifurcated. Training racks became denser, hotter and more power-
hungry than ever, built around high-bandwidth HBM and tightly coupled
NVLink/UALink scale-up fabrics, with power envelopes climbing from the
~100–120 kW class of the Blackwell GB200 NVL72 to the ~190–230 kW of the Rubin
NVL72. Inference stacks, meanwhile, broke apart along a different axis:
prefill separated from decode, memory architectures diversified toward large,
cheaper pools — including LPDDR5X, dense on-chip SRAM, and even packaging-free
designs with no HBM at all — and serving moved out of the hyperscaler core
toward colocation facilities, telco sites and the edge.

This chapter documents that bifurcation and its economic engine: quantization.
The two halves of the chapter are really one story told twice. The first half
(Sections 5.1–5.6) describes the architectural split — what training racks
look like, what inference racks look like, and why cost per token became the
metric that arbitrates between them. The second half (Sections 5.7–5.14)
describes the bit-level techniques that moved that metric more than any
hardware change in 2026: McKinsey's 85–95% cost-per-token reduction figure,
FP8 as the safe production tier, INT4/NVFP4/MXFP4 as the aggressive tier, the
AWQ-vs-GPTQ method contest and its Marlin kernels, and the sub-4-bit frontier
— BitNet's 1.58-bit ternary training, PrismML's Bonsai compression of a 27B
model into 3.9 GB, and Google's TurboQuant 3-bit KV cache now shipping as an
SGLang flag.

A note on method, consistent with the rest of this dossier. Vendor performance
claims are labeled as vendor claims. Single-source or single-configuration
figures are flagged as such. Dates after September 22, 2026 are given as
announced or planned, not as fact. And where the dossier's verification
process cleared an earlier suspicion of hallucination — McKinsey's 85–95%
figure, PrismML Bonsai's existence, TurboQuant's paper trail — that clearance
is recorded in the text, because the reader deserves to know which numbers
survived adversarial checking.

| Chapter roadmap | Sections |
|---|---|
| The split | 5.1 The great specialization: training vs inference |
| The benchmark mirror | 5.2 MLPerf v6.1 as an illustration of the shift |
| The software of disaggregation | 5.3 Nvidia Dynamo and disaggregated serving |
| The training rack | 5.4 Training racks: dense, HBM, high power |
| The inference rack | 5.5 Inference racks: disaggregated, LPDDR, edge |
| The metric | 5.6 Cost per token as the decisive metric |
| The economics | 5.7 Quantization: the economics (McKinsey, June 2026) |
| The safe tier | 5.8 FP8: the safe production standard |
| The aggressive tier | 5.9 INT4, NVFP4, MXFP4: the aggressive standard |
| The methods | 5.10 AWQ vs GPTQ, Marlin kernels |
| The frontier I | 5.11 BitNet and the 1.58-bit frontier |
| The frontier II | 5.12 PrismML Bonsai: 27B models in a few gigabytes |
| The frontier III | 5.13 TurboQuant: 3-bit KV cache |
| The runtimes | 5.14 Runtimes: vLLM, TensorRT-LLM, SGLang, llama.cpp |
| The synthesis | 5.15 What quantization changes for the industry |

A note on scope, so the reader knows what this chapter does not claim. It does
not cover training-time quantization: the techniques here are inference
techniques, because training's synchronous gradient exchange has no
compression escape hatch comparable to weight quantization — you cannot
quantize a gradient the way you quantize a weight, which is one more reason
the training rack stayed dense while the inference rack diverged. It does not
cover the 2025–2026 refinements of the quantization-method lineage (SpinQuant
and successors): they were not verified for this dossier and are left out
rather than asserted. And where verification failed — FP8 native support on
AMD MI300+, the fine benchmark figures for Bonsai, the "~95% retention" single
table for NVFP4 — the text says so instead of smoothing it over. The anti-
fabrication rule is the chapter's load-bearing wall: every date, price, name
and figure below survived a verification pass, and everything else is labeled
as analysis, vendor claim, or open question.

