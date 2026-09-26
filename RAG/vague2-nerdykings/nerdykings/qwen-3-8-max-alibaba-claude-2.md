---
id: vague2-nerdykings/nerdykings/qwen-3-8-max-alibaba-claude-2
title: "Qwen 3.8 Max : Alibaba Vient-Il Vraiment De Rattraper Claude ?"
domain: nerdykings
role: reference
task: article
actors: ["Alibaba", "Anthropic", "China", "Moonshot", "United States"]
dates: []
keywords: ["claude", "qwen", "attention", "benchmark", "fable 5", "gpus", "kimi", "mixture of experts", "moe", "multimodal", "open-weight", "parameters"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/qwen-3-8-max-alibaba-claude.md
source_anchor: ""
source_lines: [32, 69]
sha256: 96c731348d83acfe822728b70ff6212368cd36d7b7b6620e1a0913667267f25d
---

# Qwen 3.8 Max : Alibaba Vient-Il Vraiment De Rattraper Claude ?

- **Qwen 3.8 Max Preview** (Alibaba) is claimed to be the 2nd best model worldwide, behind Claude Fable 5.
- Architecture: **Mixture of Experts** with **2.4 trillion parameters** (active parameter count not disclosed).
- Long-context innovation: combines classic attention with **gated delta net** for a compact summarized internal state.
- Multimodal: handles text, images, videos, and complex documents in one workflow.
- Especially strong at front-end, animations, simulations, and turning visual references into functional apps.
- King Bench: **65/80** — one point behind Claude Fable 5, ahead of Sonnet 4.8, Kimi K3, and GPT.
- Real tests validated: hand-drawn portfolio → website; 3D archery game; Forest Robot; Mesopotamian scene (excellent).
- Main flaw: over-thinks simple tasks; deep "thinking" mode can cause long waits without clear instructions.
- Still a preview; weights promised but a >1 TB model requires multiple professional GPUs.
- Practical use will mostly be via Alibaba's cloud or specialized providers.

## Technical data / figures

| Item | Value |
|---|---|
| Model | Qwen 3.8 Max Preview |
| Developer | Alibaba |
| Architecture | Mixture of Experts |
| Total parameters | 2,400 billion (2.4T) |
| Active parameters | Not disclosed |
| King Bench score | 65 / 80 |
| Ranking claim | 2nd behind Claude Fable 5 |
| Beaten in benchmark | Claude Sonnet 4.8, Kimi K3, GPT |
| Long-context mechanism | Gated delta net + classic attention |
| Modalities | Text, image, video, documents |
| Size (open-weight, compressed) | >1 terabyte; multiple pro GPUs |
| Status | Preview |

- Key concepts: **MoE**, **gated delta net**, **long context**, **visual programming**, **thinking mode**
- Real-world tests: portfolio, 3D archery, Forest Robot, Mesopotamian 3D scene

## Why this source matters for the RAG

This article provides a hands-on evaluation of a major Chinese frontier model with architectural details, independent benchmark data, and concrete usability trade-offs. It is valuable for a RAG knowledge base on multimodal models, MoE long-context designs, front-end code generation, and the US-China model race.

## Source URL

https://www.nerdykings.com/blog/qwen-3-8-max-alibaba-claude.html
