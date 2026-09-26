---
id: collect-240926-mindstudio/mindstudio/deepseek-vision-vs-claude-sonnet-4-6-vs-gemini-flash-3-which-vision-model-uses-1-1
title: "deepseek-vision-vs-claude-sonnet-4-6-vs-gemini-flash-3-which-vision-model-uses-1"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "OpenAI"]
dates: ["2024-03", "2024-10", "2024-12", "2025-01", "2025-10", "2026-04-29"]
keywords: ["claude", "deepseek", "gemini", "agent", "agents", "attention", "benchmark", "benchmarks", "compute", "cost", "distillation", "fine-tuning"]
source: docs/RAG/clean_en/mindstudio/deepseek-vision-vs-claude-sonnet-4-6-vs-gemini-flash-3-which-vision-model-uses-1.md
source_anchor: ""
source_lines: [1, 69]
sha256: 42699c6ef88b9712c5e2b0bde0a4b711a6c6722cf6c8ea2d3520a6930746ea03
---

# deepseek-vision-vs-claude-sonnet-4-6-vs-gemini-flash-3-which-vision-model-uses-1

<!-- source: https://www.mindstudio.ai/blog/deepseek-vision-vs-claude-sonnet-4-6-vs-gemini-flash-3-kv-cache-comparison -->

## The Vision Model You’re Paying 10x Too Much to Run

DeepSeek’s new vision model uses roughly 90 KV cache entries for an 80×80 image. Claude Sonnet 4.6 uses around 870. Gemini Flash 3 uses approximately 1,000. That’s not a rounding error — that’s a structural difference in how these models represent visual information, and it has direct implications for what you pay every time you process an image.

If you’re choosing a vision model for a production workload right now, that number matters more than almost any benchmark score.

The paper behind this is titled “Thinking with Visual Primitives.” It was published, then became difficult to find. The model itself started rolling out in limited form on April 29, 2026, alongside DeepSeek’s fast and expert modes in the app and on the web. The paper is hard to locate; the efficiency numbers are not.

## Why KV Cache Is the Number That Actually Matters

Most vision model comparisons lead with accuracy benchmarks. That’s fine for research papers. For production systems, the KV cache size is closer to the real cost driver.

The KV cache is what the model holds in memory while processing a sequence. Larger KV cache means more memory, more compute, more cost per inference. When you’re running thousands of image queries per day, a 10x difference in cache size is a 10x difference in what you’re paying to serve those queries.

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

This is why the 80×80 comparison is the right anchor. It’s a controlled, apples-to-apples test: same image resolution, same task, three different models. DeepSeek comes in at ~90 entries. Sonnet 4.6 comes in at ~870. Gemini Flash 3 lands at ~1,000. The gap is not marginal.

The underlying reason for DeepSeek’s efficiency is architectural, not accidental. Their custom vision transformer — they call it the DeepSeek Vision Transformer — applies a three-stage compression pipeline. A 756×756 image first becomes 2,916 patch tokens via 14×14 patches. A 3×3 spatial compression along the channel dimension reduces that to 324 tokens. Then a compressed sparse attention mechanism from the V4 paper compresses the KV cache by another factor of four. End result: 81 entries. Total compression ratio from raw pixels to KV cache: approximately 7,000x.

That’s not a trick. That’s an architecture decision made consistently across multiple model generations.

## Three Models, Three Different Philosophies

### DeepSeek Vision: Efficiency as a First Principle

DeepSeek’s vision work didn’t start with this model. The lineage runs back to March 2024 with DeepSeek VL, through Janus in October 2024 (which decoupled visual encoders for understanding versus generation), through VL2 in December 2024 (which ported mixture-of-experts and multi-head latent attention into vision), through Janus Pro 7B in January 2025, and through the DeepSeek OCR paper in October 2025.

That OCR paper is worth pausing on. The framing was strange — they called it OCR but the actual idea was: take 1,000 text tokens, render them as an image, encode the image, get back 100 vision tokens that reconstruct the original text at 97% accuracy. Ten-to-one compression on long context. Andrej Karpathy’s reaction: “the tokenizer must go, pixel may be better inputs to language models than text.”

That’s the through-line. Every model in this lineage asks the same question: what’s the cheapest representation that still works? The current vision model is the answer applied to spatial reasoning.

The language backbone is DeepSeek V4 Flash — a 284B parameter mixture-of-experts model with 13B active parameters at inference. Frontier-grade reasoning at a fraction of the activation cost. The vision encoder supports arbitrary resolution. The whole system is designed to be cheap to run.

The technique the paper introduces — “visual primitives” — is also worth understanding in this context. When the model reasons about an image, it can output inline tokens in the format `<ref>label</ref><box>x1,y1,x2,y2</box>`. These are special vocabulary tokens embedded directly in the chain of thought. Not function calling. Not a separate tool. The model literally points to things as it reasons, the same way a human uses a finger to track objects in a dense scene. Counting people in a team photo, tracing a path through a maze, distinguishing a Chihuahua from a muffin — all of these become more reliable when the model can anchor its reasoning to coordinates rather than language descriptions alone.

The training pipeline that produces this behavior is five stages: multimodal pre-training, then separate supervised fine-tuning for grounding (boxes) and pointing (points), then GRPO reinforcement learning with three reward heads (format, quality, accuracy) applied to each, then a unified RFT that merges both specialists, then on-policy distillation into a single student model. Two specialists trained separately, then consolidated. It’s an elegant design.

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

### Claude Sonnet 4.6: Capable, But Expensive Per Token

Sonnet 4.6’s ~870 KV cache entries for the same 80×80 image reflects a different set of architectural priorities. Anthropic’s vision models are strong — particularly on document understanding, complex visual QA, and tasks that require integrating visual and textual reasoning over long contexts. If you’re building something where accuracy on nuanced visual tasks is the primary constraint and cost is secondary, Sonnet 4.6 is a reasonable choice.

But ~870 entries versus ~90 is a real cost difference. If you’re running image processing at scale — document pipelines, visual inspection workflows, anything that processes hundreds or thousands of images — you’re paying roughly 10x more per image for the same input. That’s not a small consideration.

For teams evaluating GPT-5.4 vs Claude Opus 4.6 on document processing, the KV cache efficiency question is increasingly part of the calculus, not just benchmark accuracy.

### Gemini Flash 3: Fast, But Not Cheap on Vision

Gemini Flash 3 at ~1,000 KV cache entries is the least efficient of the three on this dimension. That’s somewhat surprising given that Flash variants are typically positioned as the cost-efficient option in Google’s lineup. The efficiency story here is about latency and throughput, not about memory footprint per image.

Gemini Flash 3 is still ahead of DeepSeek on raw count QA tasks, according to the paper’s benchmarks. If your use case is straightforward visual question answering — “how many items are in this image,” “what text appears here” — Flash 3 remains competitive. But on topological reasoning tasks, the gap is significant: DeepSeek scores 67% on maze navigation versus 49% for Gemini Flash 3. That’s not a close race.

The 1,000-entry KV cache also means that at scale, Gemini Flash 3 is the most expensive of the three to run on vision tasks, despite being positioned as the budget option for text.

## What the Benchmarks Actually Show (and What They Don’t)

The paper is honest in a way that deserves acknowledgment. There’s a footnote that reads: “reported scores cover only a subset of evaluation dimensions directly relevant to the research focus of this paper and are therefore not indicative of the model’s overall capabilities.”

