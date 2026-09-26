---
id: ai-industry-kb-2026-wave6/02-deepseek/v4-pro-0813-release-state-and-pricing-details-secondary
title: "V4-Pro-0813 — release-state and pricing details [SECONDARY]"
domain: deepseek
role: deep-dive
task: pricing
actors: ["Anthropic", "DeepSeek", "Moonshot", "OpenAI", "OpenRouter"]
dates: ["2026-08", "2026-08-12", "2026-08-13"]
keywords: ["pricing", "agent", "attention", "benchmark", "benchmarks", "compute", "context window", "cost", "decode", "deepseek", "fable 5", "fp4"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [657, 679]
section: "§2. DeepSeek"
delta_of: ai-industry-kb-2026
sha256: 7c8817d9cfd6fd65019fb159c3cf5d3e1ba4599c033a499e535834cbedf64de4
---

# V4-Pro-0813 — release-state and pricing details [SECONDARY]

### V4-Pro-0813 — release-state and pricing details [SECONDARY]
- **Release-state discrepancy**: the `deepseek-chat`/`deepseek-reasoner` API surface began resolving to the new checkpoint around **2026-08-12**, but the official checkpoint identifier and GA date are **2026-08-13** (0813). Corpus entries must date events by the source that claims them: "API began routing 08-12" vs "0813 GA." [SECONDARY]
- Spec: 1.6T MoE with **49B active per token**; the DSpark-augmented checkpoint adds spec tokens toward a **~1.7T physical checkpoint size** — the 1.6T/1.7T pair is a vendor/compute-vs-physical accounting nuance, not two models. Throughput reported at **~78.1 tokens/s**. [SECONDARY]
- Official API pricing after the 0813 release: **$1.32 input / $3.96 output** per million, with **$0.044 cached-input** and **half-price off-peak** during **01:00–04:00 and 06:00–10:00 UTC**. [SECONDARY]
- The 0813 GA landed alongside Moonshot's **Kimi K3** — one secondary comparison (webpronews) framed the two as the August 2026 frontier contenders. V4-Pro was DeepSeek's flagship for exactly four weeks (0813 → 0914 retirement), a tenure so short the Kimi K3 comparison was the only major head-to-head it ever got. [DIRECTIONAL]
- Vendor-claimed agent benchmarks: **Terminal-Bench 2.1: 87.9**, **CyberGym: 83.3**, **DeepSWE: 62.7** — all vendor claims until independently replicated. [SECONDARY]
- **Launch coverage** (Reuters, 2026-08-13): V4-Pro-0813 priced at **$1.32 input / $3.96 output** — ~9× input and ~14× output vs V4 Flash ($0.14/$0.28) — as DeepSeek "seeks to turn stronger benchmark performance into a premium flagship offering." The piece notes the awkward backstory: V4 Flash-0731 had **unexpectedly outperformed the April preview V4 Pro** in independent tests, forcing the official Pro to prove it was actually the flagship. [SECONDARY]
- **Pricing discrepancy**: WCCFTech and others quote **$0.435 input / $0.87 output** — ~3× below the Reuters $1.32/$3.96. The two sets are best read as different pricing tiers/windows (peak vs off-peak vs cached), not a correction; the corpus must not pick one as "the" price without the tier label. [SECONDARY]
- **Competitive context** (WCCFTech): OpenAI had just launched a price war discounting GPT-5.6 Luna 80% ($1→$0.20 input, $6→$1.20 output); DeepSeek's $0.14/$0.28 V4-Flash-0731 "eviscerated any comparative price advantage." V4-Pro-0813's Terminal-Bench 87.9 sits within a decimal point of **Fable 5's 88.0** at roughly **57× lower output-token cost**. DeepSeek was **2nd only to Anthropic in July token volume** (before V4-Flash-0731) — "I'd imagine it will be 1st for August." [SECONDARY]
- **DeepLearning.ai The Batch** (the strongest secondary): V4-Pro-0813 — 1M input / 384K output at 78.1 tok/s; 1.6T/49B (+1.7T with DSpark); adjustable reasoning (none/low/high/max, defaults high); tool calls; context caching; **53 on AA Intelligence Index (3rd among open weights)**; 10th of 115 on Arena WebDev; MIT weights; **open-source agent harness in developer preview**; API prices **increased for all models** with the release. Undisclosed: new training data/methods, knowledge cutoff. [SECONDARY]
- **Launch mechanics** (Simon Willison, via community wiki): API-only on Aug 12, **no official announcement page** — OpenRouter was the access point; benchmarks released to the **Official DeepSeek WeChat Group**, copied to a Reddit post (deleted as "low-effort"), then an **ASCII-art table on Hacker News** — "consistent with DeepSeek's pattern of informal checkpoint announcements." Willison observed **markedly different output across reasoning levels** (low/medium/high) on his pelican SVG test — variance he "had not noticed with any other model." [SECONDARY]
- **CONTRADICTION**: Artificial Analysis Index readings for V4-Pro-0813 conflict between **53** and **36** across sources — nearly certainly different AA Index revisions/dates. Per corpus policy (never mix AA Index versions), neither replaces the other; document both with their version/date or omit until confirmed. [SECONDARY]

### V4.1-Flash — architecture internals [SECONDARY]
- Depth split: **40 transformer layers = 20-layer causal encoder + 20-layer decoder** — the causal encoder-decoder structure the wave3 verification track labeled "CED." [SECONDARY]
- MoE: **1 shared + 384 routed experts, 6 routed active per token** — a much wider expert pool than V4 Flash with a modest active count, pushing specialization over per-token compute. [SECONDARY]
- **Size contradiction**: vendor headline = **552B compute backbone**; community counting adds a **196B Engram conditional memory** module for a **748B physical-checkpoint** number. This is a real accounting contradiction (compute vs physical), not two interchangeable sizes — the corpus must keep the two figures on separate ledger lines. [SECONDARY]
- Active parameters differ by phase: **8B active in prefill, 16B in decode**. [SECONDARY]
- Attention: **CSA2 (compressed sparse attention v2)**. KV cache: **FP4 at ~890 bytes/token — under 1GB for 1M tokens** — the engineering fact that makes the 1M context window deployable. [SECONDARY]
- Context/output: native **1M input, 384K output**; image input reported as possible (not firmly confirmed). [SECONDARY]
- Pricing (peak): **$0.30 input / $1.20 output**; off-peak **$0.15 / $0.60**; cache **$0.006 peak / $0.003 off-peak** — among the cheapest frontier-grade agent APIs in the corpus. [SECONDARY]
- Vendor-claimed benchmarks (label as claims): GPQA-Diamond **90.9**, Terminal-Bench 2.1 **90.6**, DeepSWE **74.2**, HLE **36.8**, CyberGym **88.1**. [SECONDARY]

