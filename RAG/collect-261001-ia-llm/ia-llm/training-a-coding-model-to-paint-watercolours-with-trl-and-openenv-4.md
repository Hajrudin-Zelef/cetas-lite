---
id: collect-261001-ia-llm/ia-llm/training-a-coding-model-to-paint-watercolours-with-trl-and-openenv-4
title: "training-a-coding-model-to-paint-watercolours-with-trl-and-openenv"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["training", "lora"]
source: docs/RAG/collect-261001-ia-llm/training-a-coding-model-to-paint-watercolours-with-trl-and-openenv.md
source_anchor: ""
source_lines: [280, 310]
sha256: 552aa5fdfdfdf4d696e6486b7b86008ce7eba16b2526aa47cc26ef9932bf6faf
---

# training-a-coding-model-to-paint-watercolours-with-trl-and-openenv

And underneath all of it sits the question this project cannot close. **178 paintings
made by models define what this trained model considers beautiful.** The pool is the bottleneck,
and it is the part of the pipeline with no principled answer.

For anyone reproducing this, there are two deliberate divergences from Narreddi's recipe, both argued earlier in the article:

- The pairwise judge draws its references half from `love` and half from`okay` , instead
of comparing against the top tier only, so a weak early policy still gets signal.
- One small sentence of craft in the prompt, paint each petal two or three times, a big pass first and a smaller, more opaque one inside it.

The rest, `all-linear` for the LoRA, infrastructure failures returning `None` instead
of 0.0, stopping the judge runs at step 110, are decisions I had to make along the way
because his blog does not specify them. His implementation is not published, so
I cannot tell whether they match his choices or diverge from them.

| artifact | where | 
|---|---|
| the recipe, and how to reproduce it | `02-watercolour/` | 
| the reference pool, with every source sketch | `watercolour-reference-pool` | 
| the environment, ready to duplicate | `watercolour-env` | 
| the HPSv3 scorer, ready to duplicate | `watercolour-hpsv3` | 
| the hps-only adapter and rollouts | `watercolour-grpo-hps-only` ·`watercolour-rollouts-hps-only` | 
| the judge-led adapter and rollouts | `watercolour-grpo-judge-led` ·`watercolour-rollouts-judge-led` | 
| the hps-led adapter and rollouts | `watercolour-grpo-hps-led` ·`watercolour-rollouts-hps-led` | 
| the gallery, every painting browsable | `watercolour-gallery` | 
| the training curves | live: `judge-led` ·`hps-led` ·`hps-only` , and the CSV files in`results/` | 
| all of it | Paint with Code | 

The per-rollout numbers in this article can be recomputed from the published datasets. Nothing here depends on any Space staying switched on.

The method and original idea are Surya Narreddi's. The library is Alejandro Campos Uribe's.
