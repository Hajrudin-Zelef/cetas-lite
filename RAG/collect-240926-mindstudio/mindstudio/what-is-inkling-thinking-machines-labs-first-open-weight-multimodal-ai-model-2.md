---
id: collect-240926-mindstudio/mindstudio/what-is-inkling-thinking-machines-labs-first-open-weight-multimodal-ai-model-2
title: "what-is-inkling-thinking-machines-labs-first-open-weight-multimodal-ai-model"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "China", "Google", "Meta", "Mistral", "OpenAI", "Z.ai"]
dates: []
keywords: ["multimodal", "open-weight", "agents", "benchmark", "benchmarks", "claude", "compute", "cost", "distribution", "fine-tuning", "glm", "inference"]
source: docs/RAG/clean_en/mindstudio/what-is-inkling-thinking-machines-labs-first-open-weight-multimodal-ai-model.md
source_anchor: ""
source_lines: [107, 195]
sha256: 2ffcfd1113c70c772fa25df2c989852433df968c996bce7f47cf9324de2d3e71
---

# what-is-inkling-thinking-machines-labs-first-open-weight-multimodal-ai-model

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
