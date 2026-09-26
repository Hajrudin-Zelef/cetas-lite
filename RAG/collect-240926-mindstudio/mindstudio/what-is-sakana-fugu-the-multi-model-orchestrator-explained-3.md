---
id: collect-240926-mindstudio/mindstudio/what-is-sakana-fugu-the-multi-model-orchestrator-explained-3
title: "what-is-sakana-fugu-the-multi-model-orchestrator-explained"
domain: mindstudio
role: reference
task: reference
actors: ["Sakana"]
dates: []
keywords: ["fugu", "sakana", "benchmark", "cost", "inference", "latency", "training"]
source: docs/RAG/clean_en/mindstudio/what-is-sakana-fugu-the-multi-model-orchestrator-explained.md
source_anchor: ""
source_lines: [203, 216]
sha256: 9e236aab3b1b198bf56762826ff764cc9f92c0a142b01f42eb3a0ea3c2f1138c
---

# what-is-sakana-fugu-the-multi-model-orchestrator-explained

### Can I build my own multi-model routing system?

Yes. Open-source projects like RouteLLM from LMSys provide frameworks for training routing classifiers. The core components are: a set of models at different capability/cost tiers, a dataset of queries labeled by which tier is appropriate, a classifier trained on that dataset, and an inference layer that routes based on classifier output. Commercial services like Martian and Unify offer managed routing without the build overhead.

## Key Takeaways

- **Sakana Fugu is a multi-model LLM orchestrator** that routes each prompt to the most appropriate model tier automatically, using a learned classifier.
- **Its two-tier design** separates fast, cheap model inference from capable, expensive inference — and routes queries based on predicted difficulty.
- **Benchmark results show meaningful cost savings** relative to always using the top-tier model, with accuracy that approaches full Tier 2 performance at reasonable routing thresholds.
- **Fugu works best at scale** — the cost and latency benefits compound with query volume, and the routing classifier performs best on stable, diverse query distributions.
- **Alternatives exist** — from open-source frameworks like RouteLLM to commercial routing services to manual model assignment in tools like MindStudio.
- **The core concept is sound** : for most real-world applications, a significant share of queries don’t need the most powerful model available. Smart routing captures that efficiency without sacrificing quality where it matters.

If you’re building AI applications and want to experiment with multi-model workflows without building routing infrastructure first, MindStudio’s no-code builder lets you assign different models to different workflow steps and test the results — a practical way to develop intuition for model selection before investing in automated routing.
