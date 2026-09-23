---
id: collect-mindstudio/mindstudio/what-is-inkling-thinking-machines-labs-open-weight-model
title: "What Is Inkling? Thinking Machines Labs' First Open-Weight Multimodal AI Model"
domain: mindstudio
role: reference
task: article
actors: ["China", "OpenAI", "United States", "Z.ai"]
dates: ["2024-09", "2026-09-23"]
keywords: ["multimodal", "open-weight", "benchmark", "benchmarks", "compute", "cost", "distribution", "fine-tuning", "glm", "inference", "memory", "mixture of experts"]
source: docs/RAG/Collect RAG/02_mindstudio/what-is-inkling-thinking-machines-labs-open-weight-model.md
source_anchor: ""
source_lines: [1, 49]
sha256: 1f70f069c6c6d58b6d95612328a1e68b51d34db03602d89396f909f09e106a64
---

# What Is Inkling? Thinking Machines Labs' First Open-Weight Multimodal AI Model

## Metadata

- **Source** : https://www.mindstudio.ai/blog/what-is-inkling-thinking-machines-labs-open-weight-model
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

**Inkling** is the first model from **Thinking Machines Lab (TML)**, the company founded by **Mira Murati** (former OpenAI CTO, who left OpenAI in September 2024). It is a **952-billion-parameter multimodal AI model** that can process both text and images (outputs in text). It is **open-weight** — model weights are publicly available for download and use. TML was founded in early 2025 around the thesis that highly capable AI models should be open and accessible, not locked behind APIs. Murati frames this as a philosophical decision: open-weight models can be audited, fine-tuned, run locally, and deployed without ongoing API dependencies.

Architecture: **Mixture of Experts (MoE)** — a routing mechanism selects a subset of expert networks per input, so only a fraction of the 952B parameters are active during any forward pass. Advantages: lower compute/memory per inference than a dense model of the same size, and specialization (experts develop strengths in code, reasoning, language without interfering). It's multimodal from the ground up — image and language understanding are part of the same architecture, not separate stacks connected by adapters — which matters for tasks requiring reasoning across both modalities (reading charts, diagrams, UI screenshots, images with text).

Benchmarks: strong on reasoning-heavy benchmarks (**MMLU, GPQA**), competitive with more resource-intensive closed models; on multimodal benchmarks like **MMBench and MathVista** holds up well against models in its weight class and above; on code generation (**HumanEval**) solid but not exceptional. The article notes benchmark scores are proxies, not guarantees, and that Inkling's profile suggests a capable general-purpose model, not a narrow specialist.

Comparison vs **GLM 5.2** (Zhipu AI, Beijing, Tsinghua ties; smaller dense architecture, strong Chinese-language): Inkling generally scores higher on English benchmarks; GLM 5.2 has advantages in Chinese-language tasks, existing Zhipu ecosystem integrations, and a dense (non-MoE) architecture. Inkling is best for English-language reasoning/vision at scale with open weights; GLM 5.2 for Chinese-focused work and deployments preferring a dense architecture.

"Open-weight" defined: trained weights publicly released (download, run locally, fine-tune) — distinct from open-source (would include training code/data) and free API access (weights unavailable). Enterprise implications: data privacy, customization via fine-tuning, cost control, vendor independence. The release signals TML's competitive strategy (open weights as distribution/ecosystem building) and is positioned as a first step, with larger systems in development. Platform access (e.g., MindStudio) allows using Inkling without self-hosting.

## Key points

- Inkling: Thinking Machines Lab's first model — 952B-parameter, open-weight, native multimodal (text+image), built by Mira Murati's post-OpenAI venture.
- MoE architecture gives high capability at lower per-inference compute; multimodal from the ground up.
- Competitive with closed frontier models on reasoning (MMLU, GPQA) and multimodal (MMBench, MathVista) benchmarks; HumanEval solid but not exceptional.
- vs GLM 5.2: Inkling higher on English benchmarks; GLM 5.2 better for Chinese tasks, uses a dense architecture.
- Open-weight ≠ open-source ≠ free API; enterprise benefits: privacy, fine-tuning, cost control, no vendor lock-in.
- Positioned as TML's first step, not a final product.

## Technical data / figures

| Feature | Inkling | GLM 5.2 |
|---|---|---|
| Parameter count | 952B (MoE) | smaller dense architecture |
| Multimodal | Yes (native) | Yes |
| Open weights | Yes | Yes |
| Primary strength | general reasoning + vision | Chinese/English bilingual, general tasks |
| Origin | Thinking Machines Lab (US) | Zhipu AI (China) |
| Architecture | Mixture of Experts | Dense transformer |

- Benchmarks: MMLU, GPQA (reasoning); MMBench, MathVista (multimodal); HumanEval (code).
- Access: download/self-host weights, fine-tune, or via multi-model platforms.

## Why this source matters for the RAG

It documents a major open-weight frontier-scale multimodal release and its MoE architecture, benchmarks, and comparison against GLM 5.2 — useful for evaluating open-weight vs proprietary options. It also clarifies the open-weight/open-source distinction and enterprise benefits (privacy, fine-tuning, cost control), which are core to local/hybrid AI strategy.
