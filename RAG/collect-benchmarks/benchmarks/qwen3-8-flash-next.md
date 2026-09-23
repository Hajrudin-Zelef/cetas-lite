---
id: collect-benchmarks/benchmarks/qwen3-8-flash-next
title: "Qwen3.8-Flash-Next Benchmark Scores & Performance"
domain: benchmarks
role: reference
task: benchmark
actors: ["Alibaba"]
dates: ["2026-09-22", "2026-09-23"]
keywords: ["benchmark", "agentic", "agents", "benchmarks", "context window", "multimodal", "open-weight", "reasoning"]
source: docs/RAG/Collect RAG/06_benchmarks/qwen3-8-flash-next.md
source_anchor: ""
source_lines: [1, 70]
sha256: 77739341a8105fae78bebeaad0322dc2e07c2fcc0c790e90243c835a1cc96205
---

# Qwen3.8-Flash-Next Benchmark Scores & Performance

## Metadata

- **Source** : https://benchlm.ai/models/qwen3-8-flash-next
- **Site** : BenchLM
- **Type** : Benchmark
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This BenchLM model page profiles Qwen3.8-Flash-Next, an experimental-preview "Next" variant of Alibaba's Qwen3.8 Flash family. The model is open-weight and classified as a reasoning model with a 262K-token context window. BenchLM assigns it an overall score of 60.67/100, ranking it #42 out of 505 tracked AI models. The page notes that coverage is partial: only 37 of the 481 benchmarks BenchLM tracks are populated for this model, so the overall score should be treated as conservative.

The strongest results cluster in vision and grounded multimodal evaluations. MathVision scores 90.6% (95.7% with Python), CharXiv reaches 90.6% (84.6% without tools), RealWorldQA 88.5%, and AA-MMMU-Pro 79.8%. Knowledge performance is solid, with GPQA/GPQA-Diamond at 91.7% and AA-GPQA Diamond at 92.3%, though the harder Humanity's Last Exam sits at 35.9% (38.0% on AA-HLE). The Artificial Analysis Intelligence Index is listed at 39.8%. Hallucination behavior is weak: the AA-Omniscience Index is negative (-9.7%), with 24.5% accuracy and a 45.3% hallucination rate.

Agentic results are mixed. AndroidWorld is high at 84.5%, and CoWorkBench (73.9%) and Toolathlon-Verified (73.5%) are respectable, but OSWorld 2.0 is only 19.4% and Agents' Last Exam 51.2%. GDPval-AA shows 55.6% normalized (1648 Elo). On coding, LiveCodeBench v6 is very strong at 91.9% and SWE Multilingual 81%, while SWE-bench Pro is 62.5%, DeepSWE 58.7%, NL2Repo 48.1%, and AA-SciCode 50.6%; the AA Coding Index is 73.0%. Reasoning is led by AA-LCR at 79.7%, with CritPt weak at 11.1%, and instruction following (IFBench) at 81.3%.

Within the Alibaba family, Qwen3.8 Max leads at 72.02, followed by Qwen3.7 Max (62.89); Qwen3.8-Flash-Next (60.67) sits above Qwen3.8-27B (55.63) and Qwen3.7 Flash (47.44). The page was last updated September 22, 2026.

## Key points

- Overall BenchLM score 60.67/100, ranked #42 of 505; partial coverage (37/481 benchmarks) makes the score conservative.
- Open-weight reasoning model by Alibaba with a 262K context window; experimental-preview "Next" variant.
- Multimodal/vision is the standout area: MathVision 90.6%, CharXiv 90.6%, RealWorldQA 88.5%.
- Coding is strong on LiveCodeBench v6 (91.9%) and SWE Multilingual (81%), weaker on NL2Repo (48.1%).
- Weak hallucination profile: AA-Omniscience Index -9.7%, 45.3% hallucination rate.
- Ranks below Qwen3.8 Max (72.02) and Qwen3.7 Max (62.89) but above Qwen3.8-27B (55.63) and Qwen3.7 Flash (47.44).

## Technical data / figures

| Benchmark | Score |
|-----------|-------|
| Overall (BenchLM) | 60.67/100 (#42 of 505) |
| Context window | 262K |
| Coverage | 37 of 481 benchmarks |
| Artificial Analysis Intelligence Index | 39.8% |
| AA Coding Index | 73.0% |
| LiveCodeBench v6 | 91.9% |
| SWE-bench Pro | 62.5% |
| SWE Multilingual | 81% |
| DeepSWE | 58.7% |
| NL2Repo | 48.1% |
| AA-SciCode | 50.6% |
| GPQA / GPQA-Diamond | 91.7% |
| AA-GPQA Diamond | 92.3% |
| HLE / HLE w/o tools | 35.9% / 35.9% |
| AA-HLE | 38.0% |
| AA-Omniscience Index | -9.7% |
| AA-Omniscience Accuracy | 24.5% |
| AA-Omniscience Hallucination Rate | 45.3% |
| AndroidWorld | 84.5% |
| CoWorkBench | 73.9% |
| Toolathlon-Verified | 73.5% |
| OSWorld 2.0 | 19.4% |
| Agents' Last Exam | 51.2% |
| GDPval-AA (normalized / Elo) | 55.6% / 1648 |
| MathVision (w/ Python) | 90.6% (95.7%) |
| CharXiv (w/o tools) | 90.6% (84.6%) |
| RealWorldQA | 88.5% |
| AA-MMMU-Pro | 79.8% |
| AA-LCR | 79.7% |
| CritPt | 11.1% |
| IFBench | 81.3% |

## Why this source matters for the RAG

It provides a dense, structured benchmark snapshot of an open-weight Qwen3.8 "Flash-Next" reasoning model, useful for cross-model retrieval and family-internal comparison. The wide coverage across agentic, coding, multimodal, reasoning, and knowledge categories makes it a strong reference node for capability profiling and hallucination analysis.

