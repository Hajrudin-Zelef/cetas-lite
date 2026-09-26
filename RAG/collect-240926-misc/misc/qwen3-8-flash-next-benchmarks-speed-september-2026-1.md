---
id: collect-240926-misc/misc/qwen3-8-flash-next-benchmarks-speed-september-2026-1
title: "qwen3-8-flash-next-benchmarks-speed-september-2026"
domain: benchlm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "Moonshot", "OpenAI", "Sakana"]
dates: []
keywords: ["benchmark", "benchmarks", "agentic", "agents", "astra", "claude", "cost", "deepseek", "fable 5", "fugu", "gemini", "gemini 3.8"]
source: docs/RAG/clean_en/misc/qwen3-8-flash-next-benchmarks-speed-september-2026.md
source_anchor: ""
source_lines: [1, 159]
sha256: 0a02014ba497c63a2beb40b6325299a3b1a2ed35fac3e8e9da3307ff36929ea8
---

# qwen3-8-flash-next-benchmarks-speed-september-2026

<!-- source: https://benchlm.ai/models/qwen3-8-flash-next -->

## Decision snapshot

Each value carries a field reference instead of floating alone. Markers compare this model with the current ranked and priced catalog; they are not absolute quality thresholds.

Capability

60.7/100

field median 50.4#42 of 196 ranked models

Public

#42of 196

Verified —

Price

Self-hosted; infrastructure cost varies

input median $0.97No comparable first-party hosted token rate

Speed

53tok/s

field median 91 tok/sFirst token 40.34 s

Context

262Ktokens

field median 256,000Maximum output length is tracked separately

## Strongest published evidence

Multimodal & Grounded ranks #8. Particularly strong for screenshots, documents, charts, and grounded multimodal workflows.

## Validate before choosing

24 published rows leave some tracked benchmark slots empty. No comparable first-party API token rate is published.

## Category score record

Scores and ranks appear only where published evidence can be displayed. The table keeps the score, weight, cohort, and evidence state together.

| Category scores, ranks, weighting, benchmark coverage, and evidence status |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|
| Category | Score | Rank | Percentile | Weight | Benchmarks | Evidence | 
|---|---|---|---|---|---|---|
| AgenticRank #23 of 105Percentile 79thWeight 22%6 benchmarksVerified | 57.1 | 6 benchmarks | Verified |  |  |  | 
| CodingRank #27 of 135Percentile 81stWeight 20%5 benchmarksVerified | 55.9 | 5 benchmarks | Verified |  |  |  | 
| ReasoningWeight 17%0 benchmarksNot measured | Not measured | 0 benchmarks | Not measured |  |  |  | 
| MultimodalRank #8 of 50Percentile 86thWeight 12%8 benchmarksVerified | 83.1 | 8 benchmarks | Verified |  |  |  | 
| KnowledgeRank #40 of 160Percentile 75thWeight 12%4 benchmarksVerified | 56.5 | 4 benchmarks | Verified |  |  |  | 
| MultilingualWeight 7%0 benchmarksNot measured | Not measured | 0 benchmarks | Not measured |  |  |  | 
| Inst. FollowingRank #33 of 124Percentile 74thWeight 5%1 benchmarkVerified | 87.2 | 1 benchmark | Verified |  |  |  | 
| MathWeight 5%0 benchmarksNot measured | Not measured | 0 benchmarks | Not measured |  |  |  | 

## How much of this is verified

Coverage is split by category so a strong number never hides a thin evidence base. Verified means the row is tied to a published source; provisional rows remain visible but separate.

1. Agentic6/6 verified
2. Coding5/5 verified
3. ReasoningNot measured
4. Multimodal8/8 verified
5. Knowledge4/4 verified
6. MultilingualNot measured
7. Inst. Following1/1 verified
8. MathNot measured

## Capability shape

Each axis shows percentile within that category’s eligible cohort. The comparison outline is the median of the six nearest public-score peers; a collapsed vertex means the category is not rank-eligible.

Qwen3.8-Flash-Next category percentile values

- Agentic79th percentile
- Coding81st percentile
- ReasoningNot eligible
- Multimodal86th percentile
- Knowledge75th percentile
- MultilingualNot eligible
- Instruction following74th percentile
- MathNot eligible

The dashed outline is median of 6 nearest peers.

### Eligible category ranks

1. Agentic#23/105
2. Coding#27/135
3. ReasoningNot ranked
4. Multimodal#8/50
5. Knowledge#40/160
6. MultilingualNot ranked
7. Inst. Following#33/124
8. MathNot ranked

## Benchmark ledger

Coding opens by default. The marker compares each value with the best source-verified result in the catalog; provisional leaders do not set the reference. Expand the remaining categories for every published row.

## Coding5 rows

| Coding benchmark values, best verified comparison, weight, and source status |  |  |  |  |  | 
|---|---|---|---|---|---|
| Benchmark | Score | Versus best verified row | Gap | Weight | Evidence | 
|---|---|---|---|---|---|
| SWE-bench Pro | Score62.5% | Versus best verified row Best verified: Claude Opus 5.5 · 89.9% | Gap27.4 behind | Weight26% ref. weight |  | 
| DeepSWE | Score58.7% | Versus best verified row Best verified: Muse Spark 1.3 · 75.4% | Gap16.7 behind | Weight15% ref. weight |  | 
| SWE Multilingual | Score81% | Versus best verified row Best verified: Claude Opus 5.5 · 93.9% | Gap12.9 behind | Weight5% ref. weight |  | 
| NL2Repo | Score48.1% | Versus best verified row Best verified: DeepSeek V4.1 Flash · 65.4% | Gap17.3 behind | WeightDisplay only |  | 
| LiveCodeBench v6 | Score91.9% | Versus best verified row Best verified: Sakana Fugu-Ultra · 93.2% | Gap1.3 behind | WeightDisplay only |  | 

## Agentic6 rows

| Agentic benchmark values, best verified comparison, weight, and source status |  |  |  |  |  | 
|---|---|---|---|---|---|
| Benchmark | Score | Versus best verified row | Gap | Weight | Evidence | 
|---|---|---|---|---|---|
| OSWorld 2.0 | Score19.4% | Versus best verified row Best verified: GPT-6 Astra · 72.6% | Gap53.2 behind | Weight10% ref. weight |  | 
| JobBench | Score55.7% | Versus best verified row Best verified: Muse Spark 1.3 · 64.9% | Gap9.2 behind | Weight5% ref. weight |  | 
| Agents' Last Exam | Score51.2% | Versus best verified row Best verified: GPT-6 Astra · 59.3% | Gap8.1 behind | Weight3% ref. weight |  | 
| Toolathlon-Verified | Score73.5% | Versus best verified row Best verified: Claude Opus 5 · 80.6% | Gap7.1 behind | Weight3% ref. weight |  | 
| CoWorkBench | Score73.9% | Versus best verified row Best verified: Qwen3.8-Omni-Flash · 75.3% | Gap1.4 behind | WeightDisplay only |  | 
| AndroidWorld | Score84.5% | Versus best verified row Best verified: Qwen3.8-Omni-Flash · 87.1% | Gap2.6 behind | WeightDisplay only |  | 

## Multimodal8 rows

| Multimodal benchmark values, best verified comparison, weight, and source status |  |  |  |  |  | 
|---|---|---|---|---|---|
| Benchmark | Score | Versus best verified row | Gap | Weight | Evidence | 
|---|---|---|---|---|---|
| CharXivCharXiv Reasoning | Score90.6% | Versus best verified row Best verified: Qwen3.8 Max · 93.5% | Gap2.9 behind | WeightWeighted 20% |  | 
| Vision2Web | Score64.0% | Versus best verified row Best verified: Qwen3.8 Max · 69.0% | Gap5 behind | WeightDisplay only |  | 
| ERQA | Score72.3% | Versus best verified row Best verified: Qwen3.8 Max · 77.8% | Gap5.5 behind | WeightDisplay only |  | 
| LVBench | Score76.6% | Versus best verified row Best verified: Gemini 3.8 Flash · 87.1% | Gap10.5 behind | WeightDisplay only |  | 
| RealWorldQA | Score88.5% | Versus best verified row Best verified: Qwen3.8-Flash-Next · 88.5% | GapBest verified | WeightDisplay only |  | 
| MathVision | Score90.6% | Versus best verified row Best verified: Qwen3.8 Max · 95.2% | Gap4.6 behind | WeightDisplay only |  | 
| MathVision w/ PythonMathVision with Python | Score95.7% | Versus best verified row Best verified: Kimi K3 · 97.8% | Gap2.1 behind | WeightDisplay only |  | 
| CharXiv w/o toolsCharXiv Reasoning without tools | Score84.6% | Versus best verified row Best verified: Claude Mythos 5 · 88.9% | Gap4.3 behind | WeightDisplay only |  | 

## Knowledge4 rows

| Knowledge benchmark values, best verified comparison, weight, and source status |  |  |  |  |  | 
|---|---|---|---|---|---|
| Benchmark | Score | Versus best verified row | Gap | Weight | Evidence | 
|---|---|---|---|---|---|
| HLEHumanity's Last Exam | Score35.9% | Versus best verified row Best verified: Claude Fable 5.1 · 65% | Gap29.1 behind | Weight44% ref. weight |  | 
| HLE w/o toolsHumanity's Last Exam without tools | Score35.9% | Versus best verified row Best verified: Claude Opus 5.5 · 64.4% | Gap28.5 behind | Weight7% ref. weight |  | 
| GPQAGraduate-Level Google-Proof Q&A | Score91.7% | Versus best verified row Best verified: GPT-6 Astra · 96% | Gap4.3 behind | Weight3% ref. weight |  | 
| GPQA-DGPQA Diamond | Score91.7% | Versus best verified row Best verified: GPT-6 Astra · 96.0% | Gap4.3 behind | WeightDisplay only |  | 

## Inst. Following1 row

