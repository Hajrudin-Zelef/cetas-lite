---
id: collect-benchmarks/benchmarks/mimo-v2-6-pro
title: "MiMo-V2.6-Pro Benchmark Scores & Performance"
domain: benchmarks
role: reference
task: benchmark
actors: ["ExploitGym", "Xiaomi"]
dates: ["2026-09-22", "2026-09-23"]
keywords: ["benchmark", "agent", "agentic", "agents", "benchmarks", "context window", "multimodal", "open-weight", "reasoning"]
source: docs/RAG/Collect RAG/06_benchmarks/mimo-v2-6-pro.md
source_anchor: ""
source_lines: [1, 67]
sha256: b27dbcff22a0e22b232b70e6ef3e4681d0583f4a04990113e347fce7021e67a9
---

# MiMo-V2.6-Pro Benchmark Scores & Performance

## Metadata

- **Source** : https://benchlm.ai/models/mimo-v2-6-pro
- **Site** : BenchLM
- **Type** : Benchmark
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This BenchLM page covers MiMo-V2.6-Pro, Xiaomi's open-weight reasoning model with a 1M-token context window. It is the "pro" variant of the MiMo-V2.6 family and a sibling to MiMo-V2.6-Flash, with an earlier related model in MiMo-V2.5-Pro. The page reports 28 source-displayable benchmark rows but no public overall score—the overall score is listed as "Coming soon" and the model remains unranked. BenchLM notes partial coverage (28 of 481 benchmarks), so any eventual aggregate would be conservative. The page was last updated September 22, 2026.

The strongest results appear in coding and agentic security. DeepSWE reaches 71.9%, and AA-SciCode 60.9%, while ProgramBench is lower at 26.5%. Terminal-Bench 2.1 is exceptionally high at 89.9%, though Terminal-Bench 4.0 is only 34.90%. CyberGym is 94.0% and OSWorld-Verified 82%, but ExploitGym is weak at 17.8%. Toolathlon-Verified (76.9%) and JobBench (62.0%) are solid, while AutomationBench is 53.1%, Agents' Last Exam 31.6%, and GDP.pdf 19.2%. GDPval-AA reports 58.7% normalized and 1673 Elo. AA-Briefcase Elo is 1522, AA-AutomationBench 58.6%.

Knowledge and reasoning are mixed. The Artificial Analysis Intelligence Index is 46.3%. AA-LCR is strong at 86.3%, but CritPt is only 26.6%. AA-HLE is 49.4%. The hallucination profile is comparatively favorable: AA-Omniscience Index +8.4% (positive), with 34.8% accuracy and a 40.6% hallucination rate. The single multimodal entry is Design Arena Website at 1331 Elo.

For context, within Xiaomi's lineup MiMo-V2.5 leads at 55.56, followed by MiMo-V2.5-Pro (54.04) and MiMo-V2-Pro (53.79); MiMo-V2.6-Flash has no computed score yet, and MiMo-V2-Flash is 40.16. The absence of a headline score means MiMo-V2.6-Pro cannot yet be ranked against the broader 505-model index.

## Key points

- Xiaomi open-weight reasoning model with a 1M-token context window; "pro" tier of MiMo-V2.6.
- No public overall score yet ("Coming soon"), unranked; only 28 of 481 benchmarks covered.
- Very high agentic/security marks: CyberGym 94.0%, Terminal-Bench 2.1 89.9%, OSWorld-Verified 82%.
- Coding led by DeepSWE 71.9% and AA-SciCode 60.9%; ProgramBench low at 26.5%.
- AA Intelligence Index 46.3%; AA-LCR 86.3% but CritPt weak at 26.6%.
- Positive AA-Omniscience Index (+8.4%) with 40.6% hallucination rate.
- Sits below Xiaomi's MiMo-V2.5 (55.56) and MiMo-V2.5-Pro (54.04) in the family ranking.

## Technical data / figures

| Benchmark | Score |
|-----------|-------|
| Overall (BenchLM) | Coming soon / Unranked |
| Context window | 1M |
| Coverage | 28 of 481 benchmarks |
| Artificial Analysis Intelligence Index | 46.3% |
| DeepSWE | 71.9% |
| ProgramBench | 26.5% |
| Terminal-Bench 2.1 | 89.9% |
| Terminal-Bench 4.0 | 34.90% |
| AA-SciCode | 60.9% |
| CyberGym | 94.0% |
| ExploitGym | 17.8% |
| OSWorld-Verified | 82% |
| Toolathlon-Verified | 76.9% |
| JobBench | 62.0% |
| AutomationBench | 53.1% |
| AA-AutomationBench | 58.6% |
| Agents' Last Exam | 31.6% |
| GDP.pdf | 19.2% |
| GDPval-AA (normalized / Elo) | 58.7% / 1673 |
| AA-Briefcase Elo | 1522 |
| AA-LCR | 86.3% |
| CritPt | 26.6% |
| AA-HLE | 49.4% |
| AA-Omniscience Index | 8.4% |
| AA-Omniscience Accuracy | 34.8% |
| AA-Omniscience Hallucination Rate | 40.6% |
| Design Arena Website | 1331 |

## Why this source matters for the RAG

It documents a large-context, open-weight Xiaomi model whose aggregate score is not yet published, making the raw per-benchmark rows especially valuable for tracking an emerging model family. The strong security/terminal-agent scores and positive hallucination index offer differentiated signals for capability and reliability retrieval.

