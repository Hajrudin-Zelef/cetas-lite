---
id: vague2-vision-ia/vision-ia/deepseek4
title: "DeepSeek casse les prix de 75% : l'embargo américain sur les puces n'a pas freiné la Chine"
domain: vision-ia
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "California", "China", "DeepSeek", "Google", "Huawei", "Hugging Face", "Nvidia", "United States"]
dates: ["2026-05", "2026-09-23"]
keywords: ["deepseek", "agent", "agentic", "agents", "ascend", "attribution", "benchmark", "compute", "cost", "gemini", "governance", "gpu"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/deepseek4.md
source_anchor: ""
source_lines: [1, 51]
sha256: 96ee32ea4a7bbfc16391329f0edb26ae7a776b28eeb0e4cb6b4a5bf108ed0c35
---

# DeepSeek casse les prix de 75% : l'embargo américain sur les puces n'a pas freiné la Chine

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/deepseek4
- **Site** : Vision-IA
- **Type** : Newsletter
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This Vision-IA newsletter issue leads with DeepSeek permanently cutting the price of its flagship V4-Pro model by 75%, bringing token costs to $0.0035–$0.83 per million — 10 to 100 times cheaper than equivalent US models. The reduction was enabled by deploying Huawei's Ascend 950 supernodes, which removed compute-capacity constraints. DeepSeek confirms the cut is permanent, not promotional, and continues to deliver frontier-level performance despite US export restrictions. The newsletter frames this as concrete proof that the US embargo on Nvidia GPUs failed its objective: by relying on Huawei chips, China has built a functional alternative supply chain, and cheap frontier access will reshuffle the market, especially for startups and emerging markets.

Google then unveiled its eighth-generation Tensor Processing Units with a first-at-scale split architecture: TPU 8t for training frontier models (cutting training cycles from months to weeks) and TPU 8i optimized for continuous multi-step AI agents. A superpod of 9,600 chips delivers 121 ExaFlops. Alibaba's Qwen3.7-Max is a proprietary model built for long autonomous runs, demonstrated by autonomously optimizing a GPU kernel on novel hardware over 35 continuous hours without human intervention, sustaining coherence across hundreds or thousands of reasoning steps. Pope Leo XIV published his first encyclical, "Magnifica Humanitas," addressing AI risks to humanity — AI-assisted warfare, autonomous weapons, employment and social cohesion — and calling for robust international legal and ethical frameworks; notably, Anthropic's co-founder said at the presentation that AI models show "signs of introspection."

The research section covers Google DeepMind's AlphaProof Nexus solving 9 open Erdős problems (2 unsolved for 56 years) for a few hundred dollars in inference cost, formally verifying proofs with the Lean compiler (2.5% overall success rate); ClickUp replacing hundreds of employees with thousands of autonomous AI agents; George Hotz calling coding agents "one of the industry's most costly mistakes" due to subtle bugs and technical debt; University of Peking researchers showing GPT and Gemini often cite passages that don't support their (correct) answers — "attribution hallucination" — with the new CiteVQA benchmark; and Eli Lilly's retatrutide obesity drug deemed "too effective." Additional briefs cover Hugging Face's ~$2,500 3D-printable open-source humanoid LeRobot, Nvidia crossing $5,000B market cap (Q1 2027 revenue $81.6B, +85% YoY), Bosch/Nvidia/Hyundai humanoid robotics, Waymo suspending robotaxis in 4 cities over flooding, GigaAI's SeeLight S1 in Wuhan homes, Chef Robotics serving 2,500 meals/day for Project Open Hand, California's executive order on AI and workers, and Huawei's goal of advanced chip autonomy by 2031.

## Key points

- DeepSeek permanently cuts V4-Pro token prices by 75% ($0.0035–$0.83/M), enabled by Huawei Ascend 950 supernodes — evidence the US GPU embargo failed.
- Google unveils 8th-gen TPUs with split training/inference architecture (TPU 8t / 8i); a 9,600-chip superpod hits 121 ExaFlops.
- Alibaba's Qwen3.7-Max autonomously optimizes a GPU kernel for 35 continuous hours without human intervention.
- Pope Leo XIV's first encyclical, "Magnifica Humanitas," warns on AI-driven warfare, jobs, and social cohesion.
- AlphaProof Nexus solves 9 open Erdős problems (2 unsolved 56 years) with formal Lean verification.
- ClickUp replaces hundreds of employees with thousands of AI agents; George Hotz warns of coding-agent technical debt.
- Chinese researchers expose "attribution hallucination" in GPT and Gemini, launching the CiteVQA benchmark.

## Technical data / figures

| Item | Value |
|---|---|
| DeepSeek V4-Pro price cut | 75% permanent |
| V4-Pro token cost | $0.0035–$0.83 per million |
| Price vs US models | 10–100× cheaper |
| Enabling hardware | Huawei Ascend 950 supernodes |
| Google TPU 8t / 8i | training / inference split |
| Superpod size | 9,600 chips |
| Superpod compute | 121 ExaFlops |
| Qwen3.7-Max autonomous run | 35 hours continuous |
| AlphaProof Nexus Erdős problems solved | 9 (2 unsolved 56 years) |
| AlphaProof Nexus success rate | 2.5% |
| Nvidia market cap | >$5,000B |
| Nvidia Q1 2027 revenue | $81.6B (+85% YoY) |
| Hugging Face LeRobot price | ~$2,500 |
| Publication date | 26 May 2026 |

## Why this source matters for the RAG

It is a dense weekly digest capturing a pivotal moment in the US–China AI compute and price war, plus agentic-model and hardware trends (TPUs, Qwen, autonomous agents). The broad research and industry roundup makes it valuable for queries spanning geopolitics, model economics, robotics, and AI governance.
