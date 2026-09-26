---
id: collect-240926-mindstudio/mindstudio/what-is-inkling-thinking-machines-labs-first-open-weight-multimodal-ai-model-1
title: "what-is-inkling-thinking-machines-labs-first-open-weight-multimodal-ai-model"
domain: mindstudio
role: reference
task: reference
actors: ["China", "Mistral", "OpenAI", "United States", "Z.ai"]
dates: ["2024-09"]
keywords: ["multimodal", "open-weight", "attention", "benchmark", "benchmarks", "compute", "funding", "glm", "inference", "memory", "mistral", "mixture of experts"]
source: docs/RAG/clean_en/mindstudio/what-is-inkling-thinking-machines-labs-first-open-weight-multimodal-ai-model.md
source_anchor: ""
source_lines: [1, 106]
sha256: 714f14db3a56f4f9f6f0d270a0a21ad50be9760c0b357fa77bbe5e226a69526c
---

# what-is-inkling-thinking-machines-labs-first-open-weight-multimodal-ai-model

<!-- source: https://www.mindstudio.ai/blog/what-is-inkling-thinking-machines-labs-open-weight-model -->

## Mira Murati’s First Shot at Open AI

When Mira Murati left OpenAI in September 2024, the AI world paid attention. She’d spent years as the company’s CTO and had a front-row seat to what closed, proprietary AI development looks like at the highest level. So when her new company, Thinking Machines Lab, released its first model — Inkling — with open weights, the contrast was hard to miss.

Inkling is a 952B parameter multimodal AI model. It can process both text and images. It’s open-weight, meaning the model weights are publicly available for download and use. And it represents the first tangible product from a company that’s been one of the most-watched AI startups since it launched.

This article covers what Inkling is, how its architecture works, what the benchmarks show, how it compares to models like GLM 5.2, and why any of this matters for people actually building with AI.

## What Thinking Machines Lab Is Building

Thinking Machines Lab (TML) was founded in early 2025. The company has positioned itself around a specific thesis: that highly capable AI models should be open and accessible, not locked behind APIs controlled by a handful of companies.

Murati has been clear that this isn’t just a product decision — it’s a philosophical one. Open-weight models can be audited, fine-tuned, run locally, and deployed without ongoing API dependencies. That’s a meaningful difference from how the largest labs have typically operated.

## Remy is new. The platform isn't.

Remy is the latest expression of years of platform work. Not a hastily wrapped LLM.

The company attracted significant funding and talent quickly. Several researchers and engineers from top AI labs followed Murati to TML, which accelerated how fast the team could move from founding to a working model release.

Inkling is the first model to come out of that effort. It’s a research-grade multimodal system, but it’s also been released with practical deployment in mind.

## Inkling’s Architecture: 952B Parameters and MoE

The headline number is 952 billion parameters. That’s a large model by any measure, but the raw parameter count doesn’t tell the full story.

### Mixture of Experts Design

Inkling uses a Mixture of Experts (MoE) architecture. In MoE models, not all parameters are active for every inference call. Instead, a routing mechanism selects a subset of “expert” networks to handle each input. Only a fraction of the total parameter count is active during any single forward pass.

This design has two practical advantages:

- **Lower compute per inference** : Because only a portion of the network runs at any given time, serving the model requires less memory and compute than a dense model of the same total size.
- **Specialization** : Different expert networks can develop stronger capabilities in different domains — code, reasoning, language understanding — without those specializations interfering with each other.

MoE has become a standard approach for frontier models at scale. GPT-4 is widely believed to use a similar architecture. Mistral’s Mixtral models brought MoE to the open-weight space at a smaller scale. Inkling pushes that further.

### Multimodal from the Ground Up

Inkling isn’t a text model with a vision module bolted on. It was designed as a multimodal system, meaning image and language understanding are part of the same architecture rather than separate stacks connected by adapters.

This matters for tasks that require reasoning across both modalities at once — reading a chart, interpreting a diagram, understanding a UI screenshot, or analyzing an image with accompanying text. Models that treat vision as an add-on often struggle with these tasks because the two modalities don’t deeply inform each other.

Inkling handles both image and text inputs natively, with outputs in text.

## Benchmark Performance

Benchmarks for Inkling show strong performance across standard evaluation categories, particularly in reasoning, visual understanding, and general knowledge tasks.

### Where Inkling Performs Well

On reasoning-heavy benchmarks — MMLU, GPQA, and related tasks — Inkling’s scores are competitive with much more resource-intensive closed models. The MoE architecture appears to help here: routing inputs to specialized expert clusters seems to improve performance on structured reasoning tasks.

On multimodal benchmarks like MMBench and MathVista (which tests visual mathematical reasoning), Inkling holds up well compared to models in the same weight class and several above it.

Code generation benchmarks like HumanEval show solid but not exceptional results. Inkling performs meaningfully better than average, but it’s not the strongest model in this specific category.

### What the Numbers Mean in Practice

Benchmark scores are proxies, not guarantees. A model that tops MMLU doesn’t automatically write better emails or analyze your business data more accurately. What benchmark performance does signal is:

- The model generalizes well across domains
- It’s been trained on diverse, high-quality data
- The architecture doesn’t have obvious weak spots that might surface in real tasks

Inkling’s benchmark profile suggests it’s a capable general-purpose model, not a narrow specialist.

## How Inkling Compares to GLM 5.2

The meta description for this article specifically mentions GLM 5.2, so let’s address that directly.

GLM 5.2 comes from Zhipu AI, a Beijing-based lab with ties to Tsinghua University. The GLM series has been one of the more prominent open-weight model families to emerge outside of the US and European AI ecosystem. GLM models are strong at Chinese-language tasks and competitive on general English benchmarks.

### Where They Differ

| Feature | Inkling | GLM 5.2 | 
|---|---|---|
| Parameter count | 952B (MoE) | Smaller dense architecture | 
| Multimodal | Yes (native) | Yes | 
| Open weights | Yes | Yes | 
| Primary strength | General reasoning + vision | Chinese/English bilingual, general tasks | 
| Origin | Thinking Machines Lab (US) | Zhipu AI (China) | 
| Architecture | Mixture of Experts | Dense transformer | 

The two models serve somewhat different primary audiences. GLM 5.2 is particularly strong for teams working in Chinese or with Chinese-language data. Inkling’s primary design goals seem to be English-language reasoning and vision tasks, though it handles multiple languages.

In terms of raw capability on English benchmarks, Inkling generally scores higher. But “higher benchmark scores” doesn’t mean Inkling is the right choice for every use case. If your workflow is multilingual or you have specific deployment constraints that favor GLM’s architecture, that matters more than aggregate scores.

### Best for Comparisons

**Inkling is best for**: Large-scale deployments where English-language reasoning and vision tasks are central, teams who want open weights with a large active-parameter count, and applications that need native multimodal capability.

**GLM 5.2 is best for**: Teams working heavily in Chinese, organizations with existing integrations into Zhipu’s ecosystem, and deployments where a dense (rather than MoE) architecture is preferred.

## Why Open-Weight Matters Here

It’s worth stepping back to explain why open-weight models are significant beyond the technical details.

### What “Open-Weight” Actually Means

Open-weight means the trained model weights are publicly released. Anyone can download them, run them locally, fine-tune them on custom data, or deploy them on their own infrastructure.

This is different from:

