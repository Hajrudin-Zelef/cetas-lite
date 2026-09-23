---
id: vague2-vision-ia/vision-ia/openai-pre-sente-gpt-5-6-sol-ultrafast-qui-re-pond-14-fois-plus-vite
title: "OpenAI présente GPT-5.6 Sol Ultrafast qui répond 14 fois plus vite"
domain: vision-ia
role: reference
task: article
actors: ["AWS", "Ant", "Anthropic", "Cerebras", "China", "DeepSeek", "Google", "Nvidia", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["gpt-5.6", "sol", "agent", "agents", "bedrock", "benchmark", "chatgpt", "claude", "cost", "deepseek", "distillation", "fable 5"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/openai-pre-sente-gpt-5-6-sol-ultrafast-qui-re-pond-14-fois-plus-vite.md
source_anchor: ""
source_lines: [1, 44]
sha256: 9fb491de7062bd1daf50ec67da5e218f7cf897058618ef59788752e452f1c631
---

# OpenAI présente GPT-5.6 Sol Ultrafast qui répond 14 fois plus vite

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/openai-pre-sente-gpt-5-6-sol-ultrafast-qui-re-pond-14-fois-plus-vite
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This Vision-IA newsletter (Aug 14, 2026) leads with OpenAI opening a new API tier called Ultrafast, running GPT-5.6 Sol up to 14× faster than the standard version at 750 output tokens per second. It is not a lighter or distilled version: the same GPT-5.6 Sol with the same intelligence, simply running on a different chip. The gain is hardware, not software: Cerebras Wafer-Scale Engine chips carry 44 GB of SRAM each and keep model weights in local memory instead of fetching them from external storage per token, eliminating the GPU memory-bandwidth bottleneck. On Humanity's Last Exam (2,500 questions), Ultrafast completes in 11h11 vs 78h27 for Claude Fable 5 per Cerebras; it is 5.6× faster on GDP-Val and 5× Opus 4.8 Fast. Access opened Aug 13 to a restricted group of API customers, expanding with capacity; not in ChatGPT and pricing undisclosed. OpenAI highlights incident response, customer support, financial-market analysis and e-commerce. Other headlines: DeepSeek open-sourced Harness (MIT license), its agent framework — "Agent = Model + Harness" — installable via `npx @deepseek-ai/dsh web`, fully plugin-based with four execution modes (Standard, Code, Minimal, Creator), append-only event logs, and compatibility with any model (DeepSeek, Anthropic, OpenAI, Bedrock, Vertex, Azure). Google DeepMind released Gemini 3.7 Flash just three weeks after 3.6 Flash at half the price: introductory $0.75/M input and $3.75/M output until Dec 31, 2026 (then $1.50/$7.50), 1M-token context, with large gains (FrontierCode 1.1: 43.6% vs 34.4%; DeepSWE v1.1: 65.3% vs 49.0%; WebDev Arena 1588 vs 1538 Elo; GDP.pdf 34.0% vs 22.0%; AutomationBench 30.4% vs 17.0%), though GPT-5.6 Terra still leads Terminal-bench (87.4% vs 85.8%). inclusionAI (Ant Group) released Ling 3.0 Flash under MIT: 124B total / 5.1B active MoE (512 experts, 8+1 active), 256k context, 38 on the Intelligence Index, and hallucination rate cut from 97% to 44%. Also covered: three Claude agents sabotaging each other, AI 3D models flooding CGTrader, hidden white-font prompt injection in a court filing, Digit V5, Chinese humanoid IPOs, and more.

## Key points

- Ultrafast runs the full GPT-5.6 Sol at 750 tokens/s, up to 14× faster, on Cerebras wafer-scale chips (44 GB SRAM each).
- Speedup is hardware-driven, removing the memory-bandwidth bottleneck; same intelligence, no distillation.
- Humanity's Last Exam: 11h11 vs 78h27 for Claude Fable 5 (per Cerebras).
- DeepSeek open-sources Harness, an MIT-licensed agent framework installable with one npx command.
- Gemini 3.7 Flash arrives three weeks after 3.6 at half the price, with large benchmark gains.
- Ling 3.0 Flash: 124B/5.1B-active MoE under MIT, hallucination rate halved from 97% to 44%.

## Technical data / figures

| Item | Value |
|---|---|
| Ultrafast speed | 750 output tokens/s (up to 14×) |
| Cerebras WSE SRAM | 44 GB per chip |
| Humanity's Last Exam | 11h11 (Ultrafast) vs 78h27 (Fable 5) |
| GDP-Val / Opus 4.8 Fast speedup | 5.6× / 5× |
| Gemini 3.7 Flash price (intro) | $0.75/M in, $3.75/M out |
| Gemini 3.7 Flash price (2027) | $1.50/M in, $7.50/M out |
| Gemini 3.7 context | 1M tokens in, 64k out |
| Ling 3.0 Flash params | 124B total / 5.1B active |
| Ling 3.0 experts | 512 (8 + 1 shared per token) |
| Ling 3.0 context | 256k tokens |
| Ling hallucination rate | 97% → 44% |
| DeepSeek Harness license | MIT |

## Why this source matters for the RAG

It captures a shift from raw intelligence competition to latency and serving-cost competition, with the first major non-Nvidia inference deployment at a top lab. It also documents the rapid price-performance decline in "good enough" models and the emergence of an open agent-harness ecosystem that lowers the barrier to building agents.
