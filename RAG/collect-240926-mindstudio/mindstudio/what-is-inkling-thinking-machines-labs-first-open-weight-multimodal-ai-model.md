---
id: collect-240926-mindstudio/mindstudio/what-is-inkling-thinking-machines-labs-first-open-weight-multimodal-ai-model
title: "what-is-inkling-thinking-machines-labs-first-open-weight-multimodal-ai-model"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "China", "Google", "Meta", "Mistral", "OpenAI", "United States", "Z.ai"]
dates: ["2024-09"]
keywords: ["multimodal", "open-weight", "agents", "attention", "benchmark", "benchmarks", "claude", "compute", "cost", "distribution", "fine-tuning", "funding"]
source: docs/RAG/clean_en/mindstudio/what-is-inkling-thinking-machines-labs-first-open-weight-multimodal-ai-model.md
source_anchor: ""
source_lines: [1, 195]
sha256: b8acee80d7836085d29bbfbe2f092c8893580e244abfa6001f552fb76d2694ac
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

- **Open source** : True open source would include training code, data, and weights. Most “open-weight” models don’t fully release training data or complete training pipelines.
- **Open access API** : Some models are free to use via an API but the weights aren’t available. You can call them but can’t run or modify them yourself.

Inkling falls into the open-weight category. You can download and run the weights. You can fine-tune it. You’re not dependent on TML’s infrastructure.

### The Practical Implications

For enterprises, open weights mean:

- **Data privacy** : You can run inference on sensitive data without sending it to a third-party API
- **Customization** : Fine-tune on proprietary data to improve performance on specific tasks
- **Cost control** : Running your own inference can be cheaper at scale than paying per-token API costs
- **Vendor independence** : No lock-in to a single provider’s pricing or availability

These aren’t abstract benefits. They’re reasons why companies like Mistral, Meta (with Llama), and now Thinking Machines Lab have bet on open-weight releases as a strategy.

## 
      Plans first.
      *Then code.*
    

    Remy writes the spec, manages the build, and ships the app.

## Using Inkling and Other Frontier Models in Practice

Understanding what a model is capable of is one thing. Using it in a real workflow is another.

Most teams don’t deploy raw model weights directly into production. They need infrastructure around it — prompting systems, integrations with other tools, user interfaces, and reliability layers. That’s where platforms that abstract over model complexity become relevant.

## What Inkling’s Release Signals for the Broader AI Landscape

Inkling’s release is notable beyond its specs. It’s the first major signal of what Thinking Machines Lab is actually building.

### The Talent and Backing Argument

TML launched with credibility. Murati isn’t a first-time founder trying to compete with frontier labs — she ran product and engineering at one of the most advanced AI organizations in the world. The team she assembled has similar backgrounds.

That pedigree means Inkling isn’t a hobbyist project or a research artifact. It’s a deliberate product release from people who understand what frontier AI development requires. The 952B parameter MoE architecture isn’t an accident — it reflects specific design decisions about capability and deployability.

### Open Weight as a Competitive Strategy

Releasing Inkling as open-weight is a specific competitive move. It creates distribution. Anyone who downloads and deploys the model is now part of TML’s ecosystem, potentially contributing to fine-tuned variants, benchmarking data, and community feedback.

It also differentiates TML from OpenAI, Anthropic, and Google — none of whom release weights for their top models. If you believe that open-weight models will become the default for enterprise AI (rather than closed APIs), TML is positioning early.

### What Comes Next

Inkling is explicitly a first model, not a final one. TML has indicated ongoing development on larger and more capable systems. The question is whether Inkling’s open-weight release will generate the kind of community and ecosystem momentum that accelerates future development — similar to how Llama’s releases accelerated Meta’s position in the open model space.

## Frequently Asked Questions

### What is Inkling AI?

Inkling is the first open-weight AI model released by Thinking Machines Lab, the company founded by former OpenAI CTO Mira Murati. It’s a 952B parameter multimodal model capable of processing both text and image inputs. The model weights are publicly available, meaning anyone can download, run, and fine-tune it.

### Who made Inkling?

Inkling was made by Thinking Machines Lab (TML), an AI research company founded by Mira Murati in early 2025. Murati previously served as CTO of OpenAI. TML has brought together researchers and engineers from several leading AI organizations.

### What does “open-weight” mean for an AI model?

Open-weight means the trained parameters of the model are publicly released for download. Users can run the model locally, fine-tune it on their own data, and deploy it on private infrastructure — without depending on the model creator’s API or servers. This is distinct from open-source (which would also include training data and code) and from free API access (where you can use the model but not access or modify the weights).

### How does Inkling compare to GPT-4 or Claude?

## Remy doesn't write the code. It manages the agents who do.

Remy runs the project. The specialists do the work. You work with the PM, not the implementers.

Inkling is competitive with closed frontier models on many standard benchmarks, particularly reasoning and vision tasks. It doesn’t consistently outperform GPT-4 or Claude 3 Opus across all categories, but it holds its own in most general-purpose evaluations. The key practical difference is that Inkling’s weights are open — you can run it locally and customize it, which GPT-4 and Claude don’t allow.

### What is a Mixture of Experts model?

In a Mixture of Experts (MoE) model, the network is divided into specialized sub-networks called “experts.” A routing mechanism selects which experts handle each input, so not all parameters are active at once. This lets very large models (like Inkling’s 952B parameters) run inference at lower compute costs than an equivalently sized dense model, while still benefiting from the capacity that a large total parameter count provides.

### Can I use Inkling without downloading the weights myself?

Yes. You don’t need to self-host Inkling to use it. Platforms like MindStudio provide access to a wide range of models — including emerging open-weight models — through a single interface, without requiring you to manage your own model infrastructure. This is practical for teams that want to evaluate or build with new models quickly.

## Key Takeaways

- Inkling is Thinking Machines Lab’s first model — a 952B parameter, open-weight, multimodal AI built by Mira Murati’s post-OpenAI venture.
- It uses a Mixture of Experts architecture, which allows high capability at lower per-inference compute costs.
- Benchmark performance is competitive with closed frontier models, particularly in reasoning and vision tasks.
- Compared to GLM 5.2, Inkling scores higher on English-language benchmarks; GLM 5.2 has advantages in Chinese-language tasks and uses a dense architecture.
- Open-weight releases have practical enterprise advantages: data privacy, fine-tuning, cost control, and no vendor lock-in.
- TML is explicitly positioning Inkling as a first step, not a final product.

If you want to experiment with Inkling and other frontier models without managing infrastructure, MindStudio lets you access 200+ models in one place — no API keys required, free to start. It’s a practical way to test how a model like Inkling actually performs on your specific tasks before committing to a deployment strategy. You can also explore how AI agents built on MindStudio use multiple models in coordinated workflows to handle complex, multi-step tasks.
