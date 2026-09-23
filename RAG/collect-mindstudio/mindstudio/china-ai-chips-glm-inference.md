---
id: collect-mindstudio/mindstudio/china-ai-chips-glm-inference
title: "GLM 5.3 Flash Runs on Chinese Chips Without Nvidia"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "China", "Google", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Z.ai"]
dates: ["2026-09-23"]
keywords: ["glm", "nvidia", "agentic", "benchmark", "claude", "cost", "cost per token", "gemini", "gpt-5.6", "gpus", "inference", "kimi"]
source: docs/RAG/Collect RAG/02_mindstudio/china-ai-chips-glm-inference.md
source_anchor: ""
source_lines: [1, 55]
sha256: 112e42c48858dc8bb22bc79c68b99a1ac0d476d9ab78943b7af489a3a7e84562
---

# GLM 5.3 Flash Runs on Chinese Chips Without Nvidia

## Metadata

- **Source** : https://www.mindstudio.ai/blog/china-ai-chips-glm-inference
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article examines ZAI's claim that it served more than 100 trillion tokens per day of its open-weights model GLM 5.3 Flash running entirely on Chinese-made AI chips, with no Nvidia hardware in the stack. The reporting is cited from semi analysis via a source video. GLM 5.3 Flash is a mixture-of-experts model with 320 billion total parameters and 18 billion active per token, positioned as a faster, cheaper "flash" variant in the GLM 5.3 family. The headline significance is that a near-frontier-capable model, at genuine production volume, ran on domestic silicon rather than the GPUs that have dominated AI infrastructure for a decade.

The article frames inference at scale as the real bottleneck determining who can deploy AI cheaply and widely. Training a model once is one problem; serving it to millions of users daily at a low enough cost per token to sustain a business is a harder one. ZAI's reported claim is that Chinese chip stacks can now handle that harder problem for a model competitive with mainstream frontier systems on coding and agentic tasks. The key enabler described is hardware and model co-design: chips, high-bandwidth interconnect, and a serving stack built together specifically for that hardware rather than porting an Nvidia-first pipeline. ZAI claims "hardware efficiency and per token cost comparable to mainstream Nvidia GPUs." The article cautions this is ZAI's own reported claim, not independently verified in the source material, but notes that 100T tokens/day is production-scale traffic, a far stronger signal than an isolated benchmark.

Performance comparisons: GLM 5.3 Flash scored 84.3 on Terminal Bench (ahead of full-size GLM 5.2), close to GPT-5.6 Tera and Gemini 3.7 Flash. On Deep Suite it scored 63.4, a sizable jump from GLM 5.2 and competitive with GPT-5.6 Tera. On GDPval (an OpenAI benchmark for real-world knowledge work) it reportedly ranked first by a large margin among compared models. Against the largest frontier models the gap is real but narrow: on the artificial analysis intelligence index, GLM 5.3 Flash scored 57 vs Claude Opus 4.5's 62 — notable given Opus 4.5 is estimated in the multi-trillion parameter range while Flash runs at 320B total / 18B active.

Cost per completed task is where the numbers get striking. Claude Opus 4.5 reportedly costs ~$3.14 per task on the intelligence index measure; GPT-5.6 Soul ~95 cents; Kimi K3 ~84 cents; GLM 5.3 Flash ~9 cents — roughly 2–3% of Opus 4.5's cost while landing only a handful of points behind. However, cheapest is not most efficient: GPT-5.6 Luna Max reportedly costs about half of GLM 5.3 Flash's price while scoring only a few points lower, and GLM 5.3 Flash reportedly uses ~47,000 output tokens on average per intelligence-benchmark task — more than double Luna Max's usage. Open weights mean it is available via ZAI's API, OpenRouter, or self-hosting, and multiple providers can compete on price. Informal testing showed mixed open-ended generation results (Soul produced more polished 3D scene renders; GLM 5.3 Flash produced a cleaner website in one test).

## Key points

- ZAI reportedly served 100+ trillion tokens/day of GLM 5.3 Flash entirely on Chinese-made AI chips, no Nvidia hardware.
- GLM 5.3 Flash is open weights: 320B total parameters, 18B active per token (MoE).
- Key enabler is full-stack co-design: chips, interconnect, and serving software optimized together.
- Cost per completed task ~9 cents vs ~$3.14 for Claude Opus 4.5, while scoring only a few points lower (57 vs 62 on the artificial analysis intelligence index).
- Benchmark scores: Terminal Bench 84.3, Deep Suite 63.4, first on GDPval among compared models.
- Not the most token-efficient: ~47,000 output tokens/task, more than double GPT-5.6 Luna Max, which costs about half.
- Claim is ZAI's own and not independently verified in the source material.
- Suggests a viable full-stack domestic alternative to Nvidia for frontier-adjacent inference at production scale.

## Technical data / figures

| Item | Value |
|---|---|
| Total parameters | 320B (MoE) |
| Active parameters per token | 18B |
| Reported serving volume | >100 trillion tokens/day |
| Hardware | Chinese-made AI chips, no Nvidia |
| Terminal Bench | 84.3 |
| Deep Suite | 63.4 |
| GDPval | Ranked first by large margin |
| Artificial analysis intelligence index | 57 (vs Claude Opus 4.5: 62) |
| Cost per task (GLM 5.3 Flash) | ~$0.09 |
| Cost per task (Claude Opus 4.5) | ~$3.14 |
| Cost per task (GPT-5.6 Soul) | ~$0.95 |
| Cost per task (Kimi K3) | ~$0.84 |
| Avg output tokens/task | ~47,000 |
| Availability | ZAI API, OpenRouter, self-host |

## Why this source matters for the RAG

It documents a major geopolitical and economic shift in AI inference — a production-scale, Nvidia-free serving stack for a frontier-adjacent open-weight model — with concrete cost and benchmark figures. It is essential context for evaluating open-weight model economics, hardware independence, and per-task cost comparisons in the RAG.

