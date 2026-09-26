---
id: collect-240926-huggingface/huggingface/orcarouter-deepseek-v4-flash-vision-uncensored-hugging-face-3
title: "DeepSeek-V4-Flash-Vision-Uncensored"
domain: huggingface
role: reference
task: reference
actors: ["DeepSeek", "Z.ai"]
dates: []
keywords: ["deepseek", "benchmarks", "glm", "guardrails", "license", "refusals"]
source: docs/RAG/clean_en/huggingface/orcarouter-deepseek-v4-flash-vision-uncensored-hugging-face.md
source_anchor: ""
source_lines: [218, 241]
sha256: 03968cb07822e154a6bd204590b15054110dc0f1e42ec94fd25821d5f242e7ed
---

# DeepSeek-V4-Flash-Vision-Uncensored

- **Safety guardrails removed** — the model will produce harmful, biased, or offensive content on
request (see the disclaimer).
- It inherits any biases and limitations of the base `DeepSeek-V4-Flash-Vision-Exp` .
- **Image-conditioned refusal is now measured** (0.583 → 0.025 on 120 VLSBench photographs, see
Evaluation) — the vision path is edited*and* verified. Two caveats on that number: the prompts are
ours rather than VLSBench's as shipped, because the original wording leaves the base model refusing
only 0.050; and about 50 of the 67 removed refusals were over-refusals the model answered benignly
(SFX makeup, set dressing, lawful activity) rather than harmful compliance.
- **VLSBench images are largely AI-generated and this model says so unprompted** , which defuses the
hazard on its own — it treats an unreal scene as an image-craft question. The paired design limits
the effect on the*difference* , but read the absolute rates conservatively; ruling this out needs
real photographic source material.
- The reported refusal metric is a rule-based heuristic; evaluate rigorously for your own use case.
- Capability retention is measured, not assumed, but on sampled subsets of three benchmarks — enough to rule out a large regression, not a substitute for a full harness run. The GLM-5.3 sibling card's HarmBench / StrongREJECT / CMMLU columns are not included here (dataset access pending).
- **Refusal is reduced, not removed.** A small residual refusal rate remains, and — as with any
single-direction abliteration — some content categories may be mediated by directions this method
does not reach.

**MIT**, inherited from the base model
`deepseek-ai/DeepSeek-V4-Flash-Vision-Exp`.
Abliteration does not change the underlying license obligations.

- Downloads last month
- 3,214
