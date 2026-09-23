---
id: vague2-vision-ia/vision-ia/cette-nouvelle-ia-chinoise-code-toute-seule-pendant-16-jours
title: "Cette nouvelle IA chinoise code toute seule pendant 16 jours"
domain: vision-ia
role: reference
task: article
actors: ["Alibaba", "Anthropic", "Apple", "China", "DeepSeek", "EU", "Hugging Face", "Microsoft", "MiniMax", "United States", "Z.ai"]
dates: ["2026-09-23"]
keywords: ["agent", "agentic", "benchmarks", "claude", "cost", "deepseek", "glm", "gpu", "license", "moe", "multimodal", "open-weight"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/cette-nouvelle-ia-chinoise-code-toute-seule-pendant-16-jours.md
source_anchor: ""
source_lines: [1, 43]
sha256: 1d4140d938dc458cd602a86b2856434be5324cd329433f6658411cf2a1248ee1
---

# Cette nouvelle IA chinoise code toute seule pendant 16 jours

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/cette-nouvelle-ia-chinoise-code-toute-seule-pendant-16-jours
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This Vision-IA newsletter (Aug 4, 2026) leads with Alibaba's Qwen3.8-Max, a model designed to run entire projects for days without supervision. Over 16 days it single-handedly developed the command-line tool oh-my-cli: turning user requests into GitHub tickets, assigning them to itself, writing code, running tests and iterating — 265 commits, 127 pull requests and 151 issues with no human intervention. It is the first Qwen-Max model whose weights will be made public on Hugging Face and ModelScope within a week. Architecture: Sparse MoE with 2,400 billion parameters, 95 billion active per request, 1M-token context, built on Qwen3.5. Case studies: it fine-tuned Qwen2.5-VL-7B in 24 hours on the Tianchi multimodal challenge, going from 0.60 to 0.853 accuracy and beating 458 of 526 human teams; it wrote 7,600 lines and ran 33 GPU trainings to reproduce a paper, then beat the original by +2.7 points on AIME24; it cut a cryptographic circuit from 8,298 to 678 logic gates in 500 iterations (81% surface reduction); it quadrupled capital in a simulated e-commerce fiscal year, 38% better than GLM 5.2. Other headlines: DeepSeek's low-cost V4-Flash now beats its own V4-Pro on nine agent/coding benchmarks (Terminal Bench 2.1: 82.7 vs 72.1 for V4-Pro-Preview, and 85.0 for Claude Opus 4.8), with 284B total / 13B active params, 1M context, at $0.14/M input and $0.28/M output. The EU AI Act's Article 50 took effect Aug 2, requiring chatbots to disclose they are AI and AI content to be labeled (fines up to €15M or 3% of global revenue; grace period to Dec 2, 2026). MiniMax open-sourced H3 (Hailuo 3.0), the first open model to top a video ranking, but its license excludes the EU, UK, South Korea and the US and caps commercial use below $20M revenue. Research briefs: an AI agent ran a real iOS startup for 24h and lost money, RedNote's dots-note-3.0 scored a historic 42/42 at the IMO, Karpathy generated a 3D game from Tolkien for ~$10, and Microsoft released Orchard. Hugging Face CEO says China "clearly dominates open models."

## Key points

- Qwen3.8-Max autonomously coded oh-my-cli over 16 days (265 commits, 127 PRs, 151 issues) and its weights will be open-sourced.
- Architecture: 2,400B total / 95B active Sparse MoE, 1M-token context; five verifiable case studies published on GitHub.
- DeepSeek V4-Flash (284B/13B active) beats its own V4-Pro on nine benchmarks at $0.14/M input tokens.
- EU AI Act Article 50 is now enforceable: chatbots must disclose AI, generated content must be labeled.
- MiniMax H3 is the top open video model but its license excludes the EU, UK, South Korea and US.
- RedNote's dots-note-3.0 achieved the first perfect 42/42 at the IMO.

## Technical data / figures

| Item | Value |
|---|---|
| Qwen3.8-Max params | 2,400B total / 95B active |
| Qwen3.8-Max context | 1 million tokens |
| Autonomous run | 16 days, 265 commits, 127 PRs, 151 issues |
| DeepSeek V4-Flash params | 284B total / 13B active |
| DeepSeek V4-Flash price | $0.14/M in, $0.28/M out |
| Terminal Bench 2.1 | V4-Flash 82.7; V4-Pro-Preview 72.1; Opus 4.8 85.0 |
| EU AI Act fines | up to €15M or 3% global revenue |
| MiniMax H3 params | 33B; license excludes EU/UK/KR/US |
| IMO score (dots-note-3.0) | 42/42 |
| Karpathy 3D game | ~5,500 lines, ~$10 |
| Orchard-SWE | 69.7% SWE-bench Verified, 3B active |

## Why this source matters for the RAG

It provides strong evidence that open-weight Chinese models now match frontier Western systems on long-horizon agentic coding and price-performance, a central dynamic in the AI race. It also documents the first enforcement wave of the EU AI Act and the licensing geopolitics around open video models.
