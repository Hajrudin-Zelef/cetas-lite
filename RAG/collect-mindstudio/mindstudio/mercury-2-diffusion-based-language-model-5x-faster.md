---
id: collect-mindstudio/mindstudio/mercury-2-diffusion-based-language-model-5x-faster
title: "What Is Mercury 2? The Diffusion-Based Language Model That Runs 5x Faster Than Claude Haiku"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "Google", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["claude", "diffusion", "benchmarks", "compute", "cost", "gemini", "gpu", "inference", "latency", "llama", "reasoning", "throughput"]
source: docs/RAG/Collect RAG/02_mindstudio/mercury-2-diffusion-based-language-model-5x-faster.md
source_anchor: ""
source_lines: [1, 51]
sha256: e02743291b51ade3baa5f5a907c7fd307dffaff40a7e221afe1cafdabc6473d7
---

# What Is Mercury 2? The Diffusion-Based Language Model That Runs 5x Faster Than Claude Haiku

## Metadata

- **Source** : https://www.mindstudio.ai/blog/mercury-2-diffusion-based-language-model-5x-faster
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article explains Mercury 2, a diffusion-based language model from Inception Labs that produces text in parallel rather than token-by-token, benchmarking around five times faster than Claude Haiku in throughput.

Inception Labs was founded by researchers with roots at Stanford, built around the thesis that autoregressive token-by-token generation has fundamental speed limits that can't be engineered away. Mercury 2 is their second-generation release, building on the original Mercury Coder models with improved quality and broader capability.

How diffusion works for text: in image diffusion, noise is added to images during training and the model learns to reverse it; inference runs denoising repeatedly from random noise. Text is discrete, so diffusion language models use masked diffusion — tokens are randomly masked with a special [MASK] token during training and the model predicts all masked positions simultaneously. At inference, Mercury 2 starts with a fully masked sequence and iteratively refines it, committing some tokens while refining others across a set number of denoising steps. The critical difference from autoregressive generation: multiple tokens can be predicted in parallel each step — it's not blocked waiting for token N before predicting N+1.

Architecture improvements in Mercury 2: a more efficient denoising schedule (high-quality outputs in fewer iterations than generation one); better calibration on natural language tasks (expanded beyond code to instruction following, reasoning, and NLG); and a scalable architecture available at multiple parameter scales.

Speed benchmarks: the "5x faster than Claude Haiku" claim is measured in output tokens-per-second (throughput), not time-to-first-token. Haiku is itself one of Anthropic's faster models, so beating it 5x puts Mercury 2 in a genuinely different generation-speed category. Distinction: latency (time to first token) vs throughput (tokens/second once generating) — autoregressive models can have low first-token latency; diffusion needs at least one full denoising pass before output. Mercury 2 wins decisively on sustained throughput (long outputs: code generation, long-form drafts, document processing).

Quality: Inception Labs published benchmarks showing Mercury 2 competitive with Claude Haiku and similar-scale autoregressive models on standard coding and reasoning benchmarks. It is not ahead of larger models like Claude 3.5 Sonnet or GPT-4o on reasoning-heavy tasks.

Where it fits: strong for code completion/generation, batch processing pipelines, cost-sensitive high-volume applications (higher tokens/s at same compute = lower cost/token), and real-time generation in interfaces. Autoregressive models still hold an edge on complex multi-step reasoning, precise instruction following (structured outputs, complex schemas), and latency-critical first-token applications.

Broader picture: autoregressive dominance (GPT, Claude, Llama, Gemini) stems from well-understood training recipes, predictable scaling laws, and massive infrastructure — not theoretical optimality. Diffusion models have more favorable scaling for long outputs since many tokens are refined in parallel per denoising step. The likely outcome: diffusion models become the right choice for a specific class of throughput-constrained use cases rather than replacing autoregressive models wholesale.

## Key points

- Mercury 2 is a diffusion-based LLM that refines entire output sequences in parallel instead of generating tokens one at a time.
- ~5x throughput vs Claude Haiku (output tokens/second) — driven by parallel token refinement matching GPU hardware.
- Competitive quality with similar-scale autoregressive models; not positioned to beat frontier models on complex reasoning.
- Best for code generation, batch processing, and high-volume text workflows; autoregressive still leads on deep reasoning and strict instruction following.
- Latency (first token) and throughput (tokens/s) are distinct metrics — Mercury 2's edge is sustained throughput.

## Technical data / figures

| Metric | Value |
|---|---|
| Speed claim | ~5x Claude Haiku (output tokens/s) |
| Generation mechanism | Masked diffusion, parallel token refinement |
| Key advantage | Sustained throughput over long outputs |
| Where it loses | Complex multi-step reasoning, precise formatting, time-to-first-token |
| Company | Inception Labs (Stanford-rooted founders) |
| Family | Mercury → Mercury Coder → Mercury 2 (multiple scales) |

## Why this source matters for the RAG

Documents a novel generation architecture (diffusion LLMs) with concrete throughput benchmarks — relevant for choosing generation models in high-volume RAG pipelines where per-token cost and output speed dominate. Clarifies the latency-vs-throughput distinction useful for production model selection.
