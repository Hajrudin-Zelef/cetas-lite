---
id: collect-240926-mindstudio/mindstudio/deepseek-s-thinking-with-visual-primitives-5-technical-breakthroughs-in-the-pape-2
title: "deepseek-s-thinking-with-visual-primitives-5-technical-breakthroughs-in-the-pape"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google"]
dates: ["2025-10"]
keywords: ["deepseek", "attention", "claude", "cost", "gemini", "kv cache", "multimodal", "qwen", "reasoning", "training"]
source: docs/RAG/clean_en/mindstudio/deepseek-s-thinking-with-visual-primitives-5-technical-breakthroughs-in-the-pape.md
source_anchor: ""
source_lines: [78, 133]
sha256: bcdde793c770a0269435f389190a67167a2dfd3f2144cc007d7b759d60f2ae6c
---

# deepseek-s-thinking-with-visual-primitives-5-technical-breakthroughs-in-the-pape

DeepSeek OCR (October 2025): the real conceptual precursor. Take 1,000 text tokens, render them as an image, encode the image, get back 100 vision tokens that reconstruct the original text at 97% accuracy. Ten times compression on long context. Andrej Karpathy’s reaction: “the tokenizer must go, pixel may be better inputs to language models than text.” That quote circulated widely, and it pointed at something real — the assumption that text tokens are the natural input format for language models may not survive contact with better vision architectures.

The visual primitives paper is the next step in that lineage. OCR said: compress text into pixels. Visual primitives says: make spatial coordinates first-class tokens in reasoning. Both are about representation efficiency. Both are about finding the cheapest form that preserves the information you actually need.

If you’re building multimodal AI applications and following this space, understanding the architecture decisions behind models like Gemma 4 — which also supports arbitrary resolution and native vision — gives useful context for how different labs are converging on similar problems from different directions.

## 5. Three Admitted Limitations That Most Coverage Will Skip

The model has three limitations the paper explicitly acknowledges.

First: resolution-bound. Fine-grain scenes can still fail. The compression pipeline is aggressive — a 756×756 image ends up at 81 KV cache entries — and at some point that compression loses detail that matters. Dense scenes with small objects are still a problem.

Second: visual primitives mode has to be triggered explicitly. The model doesn’t auto-decide when to use it. This is a significant practical limitation. A model that can point but doesn’t know when to point is a model that requires the user to know when pointing would help. That’s a burden that should eventually be internalized.

## 
      Plans first.
      *Then code.*
    

    Remy writes the spec, manages the build, and ships the app.

Third: point-based topological reasoning doesn’t generalize well across scenarios. The maze navigation numbers are strong, but the paper acknowledges the approach doesn’t transfer cleanly to all spatial reasoning tasks. It’s not a universal solution.

These are honest admissions. They also sketch the obvious roadmap: better resolution handling, automatic mode selection, and broader generalization. The limitations tell you where the next paper is coming from.

## The Compression Architecture Underneath All of This

The efficiency story deserves its own mention because it’s what makes the visual primitives approach viable at scale.

The custom vision transformer — DeepSeek calls it the DeepSeek Vision Transformer — supports arbitrary resolution using 14×14 patches. A 756×756 image produces about 571,000 pixels, which becomes 2,916 patch tokens. A 3×3 spatial compression along the channel dimension takes nine adjacent patches into one, bringing it to 324 tokens. Then compressed sparse attention from the V4 paper compresses the KV cache by another factor of four. Final result: 81 KV cache entries for the entire image.

That’s approximately 7,000× total compression from raw pixels to KV cache entries.

For an 80×80 image, the comparison is stark: DeepSeek uses roughly 90 KV cache entries, Claude Sonnet 4.6 uses around 870, and Gemini Flash 3 uses around 1,000. About 10× more efficient on the same image. Which means roughly one-tenth the cost to run.

This matters for the visual primitives technique specifically because chain-of-thought reasoning with inline bounding boxes is token-intensive. The model is emitting `<ref><box>` pairs throughout its reasoning trace. If the image representation itself is expensive, the whole approach becomes prohibitive. The compression pipeline is what makes the reasoning approach affordable.

When you’re building applications that chain vision models with other tools — the kind of multi-step workflows where MindStudio’s visual builder lets you connect 200+ models and 1,000+ integrations without writing orchestration code — the per-image cost difference between 90 and 870 KV cache entries compounds fast across thousands of requests.

## The Deployment Reality

As of April 29, DeepSeek started rolling out vision mode in the app and on the web alongside fast and expert modes — a limited test, not a full release. The paper itself is hard to find. The model behind it is in gradual rollout.

The three limitations are real constraints on current usefulness. A model that requires explicit triggering of visual primitives mode, that struggles with fine-grain scenes, and that doesn’t generalize its topological reasoning across all scenarios is not a drop-in replacement for anything.

But the direction is clear. Inline spatial tokens in chain-of-thought reasoning is a better design than language-only spatial description. The five-stage training pipeline with separate specialists and three reward heads is a more principled approach than single-pass SFT. The compression architecture makes it economically viable.

The paper was published and pulled. That’s unusual. But the ideas in it are documented, the architecture is described, and the model is rolling out. The details above are what matter — not the publication status.

For anyone building vision-heavy applications, the question worth sitting with is this: if the reference gap is real — if language genuinely can’t point — then every multimodal model that reasons purely in text is working around a structural limitation. Visual primitives is one answer to that. It probably won’t be the last.

If you’re thinking about how to build production applications on top of models like this, tools like Remy take a different approach to the development layer: you write a spec in annotated markdown, and it compiles into a complete TypeScript stack — backend, database, auth, deployment. The spec is the source of truth; the code is derived output. The abstraction level keeps rising.

The comparison between Claude Sonnet 4.6 and other frontier models is worth tracking alongside DeepSeek’s vision work — the capability gaps between labs are shifting faster than most deployment decisions account for. And if you’re evaluating smaller models for edge deployment alongside cloud-based vision models, Qwen 3.5’s approach to running locally on phones offers a useful contrast in how different labs are thinking about the efficiency-capability tradeoff.

The visual primitives paper may be hard to find. The ideas aren’t.
