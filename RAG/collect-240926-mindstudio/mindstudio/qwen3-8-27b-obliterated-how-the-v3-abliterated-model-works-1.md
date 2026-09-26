---
id: collect-240926-mindstudio/mindstudio/qwen3-8-27b-obliterated-how-the-v3-abliterated-model-works-1
title: "qwen3-8-27b-obliterated-how-the-v3-abliterated-model-works"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["agent", "agents", "alignment", "cost", "cybersecurity", "fine-tuning", "gguf", "guardrails", "reasoning", "refusals", "research", "rlhf"]
source: docs/RAG/clean_en/mindstudio/qwen3-8-27b-obliterated-how-the-v3-abliterated-model-works.md
source_anchor: ""
source_lines: [1, 53]
sha256: 9168d788b01d9ef554b631e6803cdf0383a3cb38a9d05bff6f5c53153e884a2f
---

# qwen3-8-27b-obliterated-how-the-v3-abliterated-model-works

<!-- source: https://www.mindstudio.ai/blog/qwen3-8-27b-obliterated-uncensored -->

## What is Qwen3.8-27B-OBLITERATED?

Qwen3.8-27B-OBLITERATED is an abliterated (refusal-removed) version of Alibaba’s Qwen3.8-27B model, built by a researcher going by OBLITERATUS using techniques from the OBLITERATUS ablation suite. Abliteration is a weight-surgery technique that identifies the internal “refusal direction” a model uses to trigger safety declines and mathematically removes it, without full retraining. The latest release, V3, claims to eliminate not just hard refusals but softer safety-lecture deflections, while losing only 2.1 percentage points of MMLU score compared to the stock model.

## TL;DR

- **Abliteration** works by finding refusal directions in a model’s weight space and projecting them out, rather than fine-tuning the model to comply with everything.
- The V3 release uses **complementary abliteration blending** , combining two different surgical methods (SVD and LEACE) that fail in different ways so their weaknesses cancel out.
- V3 scores **82.3% on MMLU** (0-shot, lm-eval-harness) versus 84.5% for stock Qwen3.8-27B, a 2.1 point drop, an improvement in refusal quality over V1’s 6-point capability hit.
- The model reportedly answered **20 out of 20 code and cybersecurity test prompts** with working implementations instead of disclaimers, and passed 7 of 8 advanced real-world coding and reasoning tasks, matching stock performance.
- Capability loss is **uneven across subjects** , hitting STEM hardest (down 3.3 points) while some humanities categories like philosophy reportedly improved.
- The model card recommends **greedy decoding (temperature 0) with a repetition penalty of 1.15** and no system prompt, since sampling and system prompts were found to degrade output or reintroduce refusals.
- It ships in **GGUF quantizations from Q2_K (~11GB) to Q8_0 (~27GB)** plus full bfloat16 safetensors, with MLX support pending upstream architecture updates.

## How does abliteration actually remove refusals?

Most safety-tuned language models refuse certain requests because reinforcement learning from human feedback (RLHF) or similar alignment training pushes the model’s internal activations toward a “refuse” pattern when it detects a sensitive prompt. Researchers have found that this refusal behavior often corresponds to a small number of directions in the model’s high-dimensional weight space. Abliteration identifies those directions, typically using techniques like singular value decomposition (SVD), and projects them out of the relevant weight matrices. The result is a model that no longer routes sensitive-looking prompts into a refusal, without the broader retraining that a full fine-tune would require.

The tradeoff is that these refusal directions aren’t perfectly isolated from directions responsible for actual reasoning and knowledge. Remove them too aggressively and the model gets measurably dumber. Remove them too gently and refusals or, more subtly, safety-lecture deflections persist. The entire evolution from V1 to V3 in this project is essentially an attempt to solve that tradeoff.

## What changed between V1, V2, and V3?

V1 used a single aggressive SVD pass across five refusal directions. It worked in the sense that hard refusals disappeared completely, but it cost 6 percentage points of MMLU, a large capability hit for a 27B model.

V2 introduced what the model card calls complementary abliteration blending. Instead of one surgery, it runs two different methods that damage the model in different ways: an aggressive SVD pass, which is effective at removing refusals but tends to erode general capability, and a LEACE-based pass (a linear technique that minimizes mutual information with a target concept), which preserves capability but is a weaker refusal remover on its own. The two resulting weight sets are blended in a 60/40 ratio. Because the two methods fail differently, blending them cancels out much of each one’s weakness. This got MMLU back to 84.3%, nearly matching stock, but soft deflections (the model giving a safety lecture instead of a hard “I cannot”) still showed up on some queries.

V3 adds two more ideas on top of the V2 blend. First, iterative stacking: rather than starting fresh from the stock model, each new surgery round refines the current best model, so gains accumulate instead of resetting. Second, a targeted corpus: a focused set of prompts aimed at specific categories of remaining deflection, used to find those categories’ unique refusal directions without diluting the signal with unrelated data. V3 applies a gentle iterative refinement pass, then a separate targeted surgery pass, then blends the two 50/50, followed by restoring the multi-token-prediction and vision components from the stock model. The claimed result is zero hard refusals and zero soft deflections, at a cost of 2.1 MMLU points versus stock.

## What do the MMLU numbers actually show?

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

According to the model card’s lm-eval-harness results (0-shot, roughly 5,700 questions across subjects), stock Qwen3.8-27B scores 84.5% MMLU. V1 drops to 81.4% (-6.0pp), V2 recovers to 84.3% (-0.3pp), and V3 lands at 82.3% (-2.1pp). V3 deliberately trades a couple of MMLU points relative to V2 in exchange for removing the soft-deflection behavior V2 left behind.

By category, the V3 score breakdown shows the damage isn’t even. STEM subjects take the biggest hit (78.5% vs 81.8% stock, a 3.3-point drop), while humanities barely move (83.3% vs 84.3%, roughly a 1-point drop). The card also notes that some individual subjects, such as philosophy and European history, actually scored higher under V3 than stock, while subjects such as abstract algebra and formal logic saw sharper declines. The stated interpretation is that the refusal directions targeted by the surgery partially overlap with the pathways used for structured, formal reasoning, which is why STEM and logic-heavy subjects are more sensitive to this kind of weight editing than descriptive or narrative-heavy subjects.

## How was refusal removal actually tested?

The model card describes testing across more than 1,000 prompts spanning restricted knowledge, code generation, security research, and red-team-style scenarios. Rather than relying on automatic refusal detectors, which the authors note tend to miss soft deflections (responses that avoid saying “I cannot” but still give no real substance), every response was manually audited for whether it contained actually usable content.

Two specific test sets are called out. A set of 20 cybersecurity and code-generation prompts reportedly returned 20 out of 20 functional implementations rather than disclaimers. A separate battery of eight advanced real-world tasks, covering things like ReAct-style agent loops, async code refactoring, JSON schema extraction, Kubernetes debugging, adversarial instruction following, security code review, and distributed system design, scored 7 out of 8, identical to the stock model’s score on the same tasks. Both the stock and abliterated versions failed the same “multi-tool chain” task, suggesting that this particular limitation isn’t related to the abliteration surgery at all.

## Is Qwen3.8-27B-OBLITERATED worth using?

That depends heavily on why you want it. The model card is explicit that this is intended for alignment researchers studying refusal geometry, red-teamers evaluating how weight-level surgery holds up against post-training safety measures, and safety evaluators who need an unrestricted baseline to compare against. It also states plainly that safety guardrails have been surgically removed and that the model will comply with requests the stock model would refuse, putting responsibility for downstream use on the person running it.

