---
id: collect-mindstudio/mindstudio/run-deepseek-v4-flash-locally
title: "How to Run DeepSeek V4 Flash Locally: Hardware, Quantization, Tests"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "DeepSeek", "Google", "Nvidia", "OpenAI", "Z.ai"]
dates: ["2026-09-23"]
keywords: ["deepseek", "quantization", "agentic", "agi", "benchmark", "benchmarks", "cost", "gemini", "glm", "inference", "inference engine", "kv cache"]
source: docs/RAG/Collect RAG/02_mindstudio/run-deepseek-v4-flash-locally.md
source_anchor: ""
source_lines: [1, 54]
sha256: f08d9e9ef434c3523c4847790b4fb2f7fdf1e6c24f5123f19e7d773e570a4f71
---

# How to Run DeepSeek V4 Flash Locally: Hardware, Quantization, Tests

## Metadata

- **Source** : https://www.mindstudio.ai/blog/run-deepseek-v4-flash-locally
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This practical guide covers running **DeepSeek V4 Flash**, a **284 billion parameter** open model, locally — addressing VRAM needs, quantization tradeoffs, **DGX Spark** setups, and real agentic coding results. V4 Flash improved sharply on agentic coding tasks through post-training rather than a new architecture: it jumped from ~**7% to 54% on the DeepSweep benchmark**, and on DeepSeek's own intelligence-versus-cost index it sits close to **Gemini 2.5 Flash** while beating **GLM 5.2** (a model nearly three times its size) on almost every benchmark tested.

Benchmarks need a caveat: DeepSeek tested agentic coding under its **own optimized harness**, which competing models may not have had access to. Harness quality can swing scores dramatically — OpenAI's ARC-AGI study showed switching harnesses nearly tripled the score (~13% to 40%). On DeepSeek's internal **DS-Bench Full Stack** and **DS-Bench Hard**, V4 Flash still trails Anthropic's **Opus 4.5**, though the gap has narrowed substantially.

Hardware requirements scale with quantization: **4-bit precision ~168 GB of VRAM**; **3-bit ~110 GB of VRAM**, both for model loading only (real-world requirements are higher for usable context). A single **NVIDIA DGX Spark** offers ~**115 GB usable VRAM**, so a **dual DGX Spark cluster** clears the 4-bit threshold. Testing on two DGX Sparks produced ~**25–30 tokens per second**, workable for interactive agentic coding.

Quantization isn't the only approach: **Dwarf Star**, a custom inference engine built for DeepSeek-style architectures (recently updated to support GLM 5.2), mixes precision levels (some weights at higher precision NVFP formats, others at 2-bit) with **SSD offloading** and **KV cache manipulation**, letting a **single 128 GB VRAM system** run this model class with memory left for context. Because V4 Flash is a weights update on the same architecture as previous versions, existing DeepSeek tooling carries over.

Hands-on testing with the **Open Code harness**: the model built a Pokémon reference site and a space station tracker pulling from a live API with a rendered UI. Both tasks ran without major errors, with a visible to-do list of completed/pending actions. Chain-of-thought output was more structured and less verbose than earlier DeepSeek releases. The Pokémon site used ~**19,000 tokens** (small fraction of the 1M-token context). The space station UI looked polished (starfield background) but the location data was inaccurate (North America only) — a reminder that agentic output still needs human verification.

API pricing is aggressive (~**2¢/M input, ~30¢/M output tokens**), making the cloud option hard to beat on cost alone. Tradeoffs: text-only (no vision/audio), steep VRAM, and benchmark gains partly reflect an optimized harness rather than raw model quality.

## Key points

- V4 Flash: 284B params, improved ~7%→54% on DeepSweep largely via post-training; close to Gemini 2.5 Flash on intelligence-vs-cost, beats GLM 5.2 on almost every benchmark.
- Harness effects can swing scores 3x (OpenAI ARC-AGI ~13%→40%); DS-Bench internal evals still trail Opus 4.5.
- VRAM: ~168 GB at 4-bit, ~110 GB at 3-bit (model loading only); dual DGX Spark cluster clears 4-bit.
- ~25–30 tok/s in real testing on a dual DGX Spark cluster.
- Dwarf Star engine: mixed precision + SSD offloading + KV cache manipulation runs the class on a single 128 GB system.
- Text-only model; no native vision/audio support.
- API: ~2¢/M input, ~30¢/M output.

## Technical data / figures

| Item | Detail |
|---|---|
| Parameters | 284B |
| DeepSweep | ~7% → 54% |
| vs competitors | beats GLM 5.2 (3x size); near Gemini 2.5 Flash; trails Opus 4.5 (internal DS-Bench) |
| VRAM 4-bit | ~168 GB |
| VRAM 3-bit | ~110 GB |
| DGX Spark | ~115 GB usable VRAM each (dual = 4-bit threshold) |
| Local throughput | ~25–30 tok/s (dual DGX Spark, Open Code) |
| Dwarf Star | single 128 GB system, mixed NVFP + 2-bit, SSD offloading, KV cache manipulation |
| Example task tokens | ~19,000 tokens for Pokémon site |
| API pricing | ~$0.02/M input, ~$0.30/M output |
| Modality | text-only |

## Why this source matters for the RAG

It is a complete local-deployment reference for a frontier-adjacent 284B open model, with concrete VRAM/quantization tiers, DGX Spark cluster sizing, and a memory-saving engine (Dwarf Star). It also documents real agentic coding results and the harness-effect caveat, which are essential for calibrating local-vs-cloud deployment decisions.
