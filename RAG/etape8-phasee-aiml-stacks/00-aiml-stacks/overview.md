---
id: etape8-phasee-aiml-stacks/00-aiml-stacks/overview
title: "Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)"
domain: step-8-phase-e-ai-ml-software-stacks-developer-facing
role: deep-dive
task: reference
actors: ["SGLang", "vLLM"]
dates: ["2026-09-22"]
keywords: ["benchmark", "gpu", "inference", "llama", "llama.cpp", "quantization", "research", "sglang", "tpu", "training", "vllm"]
source: docs/RAG/etape8_phaseE_aiml_stacks.md
source_anchor: ""
source_lines: [1, 47]
section: "Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)"
sha256: a346b5d713a4ed8ccf7e86fe58f9318e487cffd3b62555800afb44c1a573ec4a
---

# Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)

> **Cutoff:** 2026-09-22. **Status:** research snapshot, one writer, sole path
> `~/workspace/rag_collect/etape8_phaseE_aiml_stacks.md`. **Line budget:** ≥750 lines.
> All Markdown deliverables are in English per project rule.

## Scope

This file covers the developer-facing AI/ML software stack: deep-learning
frameworks, model libraries and hubs, LLM application frameworks, vector
databases, kernel compilers and model formats, quantization tooling, GPU/TPU
programming interfaces, distributed-training systems, and experiment/eval
tooling. It is the software counterpart to the hardware and infrastructure
layers documented elsewhere in the project.

Out of scope here, covered in other steps and cross-referenced where relevant:
inference servers (`vLLM`, `SGLang`, `Ollama`, `llama.cpp`) — see Step 4;
`pgvector` operational detail — see Step 7F; GPU/CPU hardware and
interconnects — see Step 6 phases; training-infrastructure operations
(Slurm/Kubernetes scheduling) — see Step 6.

## Methodology and provenance

- Official project sources (GitHub repositories, documentation, PyPI/NuGet
  package registries) were preferred where reachable.
- Secondary press, vendor blogs, and community mirrors supply context where
  official pages could not be checked.
- Every factual line carries exactly one provenance tag from the allowed set:
  `[official]`, `[vendor-reported]`, `[independent]`, `[secondary]`,
  `[unverified]`.
- Conflicts, gaps, and unverifiable claims are recorded in the
  Conflicts & Gaps Register (Section 13), not silently resolved.
- No version, URL, date, price, or benchmark figure is invented; where a
  figure appears it is quoted from an identified source with its tag.

## Provenance legend

- `[official]` — first-party source: project repository, documentation, or
  package registry.
- `[vendor-reported]` — vendor (company) claim not independently confirmed.
- `[independent]` — third-party measurement or study with stated methodology.
- `[secondary]` — press, blogs, mirrors, or documentation not from the project.
- `[unverified]` — plausibly correct but not confirmed against an identified
  source as of the cutoff.

---

