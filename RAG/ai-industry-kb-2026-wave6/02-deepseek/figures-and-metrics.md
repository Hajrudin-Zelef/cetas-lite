---
id: ai-industry-kb-2026-wave6/02-deepseek/figures-and-metrics
title: "Figures and metrics"
domain: deepseek
role: deep-dive
task: actor-profile
actors: ["Anthropic", "DeepSeek", "Nvidia", "OpenAI", "vLLM"]
dates: ["2025-05", "2025-09-29", "2025-12-01", "2026-04-24", "2026-05", "2026-06", "2026-06-03", "2026-08-13", "2026-09-03", "2026-09-07", "2026-09-10", "2026-09-14"]
keywords: ["agent", "astra", "benchmark", "blackwell", "claude", "compute", "decode", "deepseek", "fable 5", "fp4", "fp8", "gpt-6"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [779, 869]
section: "§2. DeepSeek"
delta_of: ai-industry-kb-2026
sha256: 67657e7a07ed80801e68a06bae26063dbd038b39e6674f72abb89b2cfc3fa15b
---

# Figures and metrics

## Figures and metrics

| Model | Total params | Active params | Context | Output | Pricing (per M) | Provenance |
|---|---|---|---|---|---|---|
| V4-Pro-0813 | undisclosed in-window | undisclosed in-window | 1M (V4 line) | 384K (V4 line) | — | [VENDOR] |
| V4.1-Flash | 552B backbone + 196B Engram | 8B prefill / 16B decode | 1M | 384K | peak $0.30/$1.20; off-peak $0.15/$0.60; cache $0.006/$0.003 | [VENDOR] |
| R2 (claimed, not released) | 32B (claimed) | — | — | — | — | [UNVERIFIED] |

| Benchmark | Score | Date/version | Provenance |
|---|---|---|---|
| V4-Pro-0813, AA Intelligence Index | 36 | v4.2/v4.3-era | [SECONDARY] |
| V4.1-Flash, AA Intelligence Index | 40 (max effort) | v4.3, 2026-09-07 | [SECONDARY] |
| V4.1-Flash, Terminal-Bench 4.0 | 27% | Sept 2026, TB 4.0 methodology | [SECONDARY] |
| V4.1-Flash, Terminal-Bench 2.1-era | #1 at 90.6% (Sept-11 update) | TB 2.1 methodology | [SECONDARY] |
| V4-Pro-0813, Vals (unspecified bench) | 52.37% (+9.48) | 2026, benchmark unspecified | [SECONDARY] |
| V4.1-Flash, AA $/Index-task | $0.27 | Sept 2026 — cheapest measured | [SECONDARY] |

### Benchmark context (version-pinned)
- AA Index **v4.3** leaders (2026-09-07): Claude Fable 5.1 (max w/ fallback) and GPT-6 Astra (max) **tied at 53** — the ceiling against which V4.1-Flash's 40 and V4-Pro-0813's 36 sit. [SECONDARY]
- V4.1 Flash **took #1 on Terminal-Bench 2.1-era at 90.6%** (Sept-11 update) — before the TB 4.0 hard reset dropped it to 27%; the two scores belong to different task sets and harnesses. [SECONDARY]
- A Vals figure for V4-Pro-0813 (52.37%, +9.48) is cited in wave6/03 but the benchmark is unspecified — treat as [SECONDARY] with missing context, not as a board score. [SECONDARY]

**No cross-version comparison:** AA Index scores above belong to different methodology versions (v4.1.1 → v4.2 → v4.3 in the week of 2026-09-03 to 2026-09-07). The 36 vs 40 gap spans versions; do not present it as a clean improvement without the version caveat. TB 4.0 scores (27%) are not comparable with TB 2.1-era scores. [DIRECTIONAL]


### New verified metrics — expansion

| Item | Value | Date | Provenance |
|---|---|---|---|
| V3.2 total parameters | **671B vs 685B** (disputed) | 2025-12 | [SECONDARY] |
| V3.2 active parameters | ~37B | 2025-12 | [SECONDARY] |
| V3.2 context | 128K | 2025-12 | [SECONDARY] |
| V3.2-Exp AIME 2025 | 89.3 (vendor claim) | 2025-12 | [SECONDARY] |
| V3.2-Exp GPQA-Diamond | 79.9 (vs V3.1-Terminus 80.7) | 2025-12 | [SECONDARY] |
| V3.2-Exp Codeforces | 2121 (vendor claim) | 2025-12 | [SECONDARY] |
| V3.2 API price | $0.14/$0.28 vs $0.27/$0.40 (disputed) | 2026 | [SECONDARY] |
| V4 Pro total/active | 1.6T / 49B | 2026-04-24 | [SECONDARY] |
| V4 Flash total/active | 284B / ~13B | 2026-04-24 | [SECONDARY] |
| V4 Pro-0813 price | $1.32 in / $3.96 out; $0.044 cached | 2026-08-13 | [SECONDARY] |
| V4 Pro-0813 off-peak | 50% off, 01:00–04:00 & 06:00–10:00 UTC | 2026-08-13 | [SECONDARY] |
| V4 Pro-0813 throughput | ~78.1 tokens/s | 2026-08 | [SECONDARY] |
| V4 Pro-0813 TB 2.1 | 87.9 (vendor claim) | 2026-08 | [SECONDARY] |
| V4 Pro-0813 CyberGym | 83.3 (vendor claim) | 2026-08 | [SECONDARY] |
| V4 Pro-0813 DeepSWE | 62.7 (vendor claim) | 2026-08 | [SECONDARY] |
| V4 Pro-0813 AA Index | 53 vs 36 (disputed revisions) | 2026-08 | [SECONDARY] |
| V4.1-Flash layers | 40 (20 causal encoder + 20 decoder) | 2026 | [SECONDARY] |
| V4.1-Flash MoE | 1 shared + 384 routed, 6 active | 2026 | [SECONDARY] |
| V4.1-Flash size | 552B compute vs 748B physical (disputed) | 2026 | [SECONDARY] |
| V4.1-Flash active | 8B prefill / 16B decode | 2026 | [SECONDARY] |
| V4.1-Flash KV cache | FP4 ~890 bytes/token | 2026 | [SECONDARY] |
| V4.1-Flash price | $0.30/$1.20 peak; $0.15/$0.60 off-peak | 2026 | [SECONDARY] |
| V4.1-Flash GA | 2026-09-10 (beta 09-08, expires-on-0910 ID, 20 concurrent) | 2026-09-10 | [SECONDARY] |
| V4.1-Flash total | 552B backbone + 196B Engram = ~748B vs ~763B (disputed headline) | 2026-09 | [SECONDARY] |
| V4.1-Flash KV vs V1 | 437x (vs Nov-2023 V1); honest: 4x vs V4 Flash | 2026-09 | [SECONDARY] |
| V4.1-Flash training | from scratch on 45T tokens; 1M context at 34T mark | 2026-09 | [SECONDARY] |
| V4.1-Flash reasoning effort | continuously controllable 1-100 | 2026-09 | [SECONDARY] |
| V4.1-Flash observed serving | 215 tok/s out, p50 TTFT 842 ms, p95 2.40 s | 2026-09 | [SECONDARY] |
| V4.1-Flash DeepSWE v1.1 | 74.2 vs V4-Pro 62.7 vs V4-Flash 54.4 (vendor claims) | 2026-09 | [SECONDARY] |
| V4.1-Flash AutomationBench | 54.8 vs 43.2 vs 37.7 (vendor claims) | 2026-09 | [SECONDARY] |
| V4.1-Flash Agent's Last Exam | 31.8 vs 25.7 vs 25.2 (vendor claims) | 2026-09 | [SECONDARY] |
| V4.1-Flash 8xH200 serving | ~915 concurrent 1M-token users vs 40 for V3 | 2026-09 | [COMMUNITY] |
| V3.2-Speciale | gold-level IMO/CMO/IOI 2025, ICPC WF 2025 | 2025-12-01 | [SECONDARY] |
| V3.2 GRPO | >10% of pre-training compute | 2025-12 | [SECONDARY] |
| V3.2-Exp indexer | learned from 2.1B tokens; 100B-token fine-tune; 5 specialists distilled | 2025-09 | [SECONDARY] |
| V4.1-Flash revision | df42c109… (tech report 1,809,802 bytes) | 2026-09 | [COMMUNITY] |
| V4.1-Flash hardware floor | ~614 GB vLLM; 4×Blackwell or 8×H200 day-one | 2026-09 | [SECONDARY] |
| V4.1-Flash context | 1,048,576 tokens; 384K max output | 2026-09 | [SECONDARY] |
| V4-Pro-0813 AA Index | 53 (reasoning version; vs 40 V4 Flash) | 2026-08 | [SECONDARY] |
| V4-Pro-0813 Arena WebDev | 10th of 115 | 2026-08 | [SECONDARY] |
| V4-Pro-0813 price premium | 9× input / 14× output vs V4 Flash ($1.32/$3.96) | 2026-08-13 | [SECONDARY] |
| V4-Pro-0813 alt pricing | $0.435/$0.87 (tier/window-dependent — do not mix) | 2026-08 | [SECONDARY] |
| V4.1-Flash DGX Spark | single GB10, 73.8 GB hot experts (25.6%), DSpark ~1.5x | 2026-09 | [COMMUNITY] |
| V4 Pro retirement | 2026-09-14; traffic routed to V4.1-Flash at Flash pricing | 2026-09-14 | [SECONDARY] |
| V4.1-Flash pricing detail | $0.15/$0.60/$0.003 off-peak; $0.30/$1.20/$0.006 peak; weekends off-peak | 2026-09 | [SECONDARY] |
| V4.1-Flash GPQA-Diamond | 90.9 (vendor claim) | 2026 | [SECONDARY] |
| V4.1-Flash TB 2.1 | 90.6 (vendor claim) | 2026 | [SECONDARY] |
| V4.1-Flash DeepSWE | 74.2 (vendor claim) | 2026 | [SECONDARY] |
| V4.1-Flash HLE | 36.8 (vendor claim) | 2026 | [SECONDARY] |
| DeepSeek maiden raise | RMB50B (~$7.4B) targeted at RMB350–400B | 2026-06-03 | [SECONDARY] |
| DeepSeek first round closed | $7B, June 2026 | 2026-06 | [SECONDARY] |
| DeepSeek second round | ≥¥10B (~$1.4B) at $71–74B pre-money — frozen Sept 2026 | 2026-09 | [SECONDARY] |
| DeepSeek inference margin | ~85% (leaked investor notes, unverified) | 2026-09 | [SECONDARY] |
| R2 launch plan | May 2026 (one source says May 2025 — likely error) | 2026 | [SECONDARY] |
| V3.2-Exp release | 2025-09-29; $0.028/M input (50% cut) | 2025-09-29 | [SECONDARY] |
| V3.2 official + Speciale | 2025-12-01; input as low as $0.07/M | 2025-12-01 | [SECONDARY] |
| V3.2 self-hosting | 685B: ~700 GB VRAM FP8 (8–10× H100); ~386 GB 4-bit (5–6× H100) | 2025-12 | [SECONDARY] |

---

## Main actors

