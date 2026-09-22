---
id: ai-industry-kb-2026-wave6/03-qwen-and-alibaba/figures-and-metrics
title: "Figures and metrics"
domain: qwen-and-alibaba
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Baseten", "China", "DeepSeek", "Google", "Meta", "Moonshot", "Nvidia", "OpenRouter"]
dates: ["2025-06-05", "2026-02-03", "2026-02-15", "2026-02-16", "2026-03-30", "2026-04-02", "2026-04-16", "2026-04-20", "2026-04-22", "2026-05-20", "2026-06-30", "2026-07-15", "2026-07-21", "2026-08-03", "2026-08-22", "2026-08-26", "2026-09", "2026-09-02", "2026-09-17", "2026-09-22"]
keywords: ["agent", "apache", "benchmark", "benchmarks", "blackwell", "capex", "cost", "decode", "deepseek", "embedding", "kimi", "leaderboard"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1258, 1359]
section: "§3. Qwen and Alibaba"
delta_of: ai-industry-kb-2026
sha256: c2401187d4c8df4f7075264a32fbad7c0d8f6a26b4c2e1680cc34c84cd0e0907
---

# Figures and metrics

## Figures and metrics

| Model | Date | Size (total/active) | License/access | Provenance |
|---|---|---|---|---|
| Qwen3-Coder-Next | 2026-02-03/04 | 80B / ~3B | Apache 2.0 | [VENDOR] |
| Qwen3.5-397B-A17B | 2026-02-16 | 397B / 17B | open | [SECONDARY] |
| Qwen3.6-35B-A3B | 2026-04-16 | 35B / ~3B | open | [SECONDARY] |
| Qwen3.6-27B | 2026-04-22 | 27B | open | [SECONDARY] |
| Qwen3.8-Flash-Next | 2026-08-26 | 125B / 6B | open | [SECONDARY] |
| Qwen3.8-Max-0902 | 2026-09-02 | snapshot (not new base) | custom license | [SECONDARY] |
| Qwen3.6-Plus / Max-Preview | 2026-04-02 / 04-20 | — | closed API | [SECONDARY] |
| Qwen3.7-Max / Plus | 2026-05-20 / 06-01 | — | closed API | [SECONDARY] |

| Benchmark | Score | Date/version | Provenance |
|---|---|---|---|
| Qwen3.8-Max (0902), AA Intelligence Index | 45 | v4.3, reported 2026-09-17 | [SECONDARY] |
| Qwen3.8-Max DeepSWE | 56.6 (vendor harness) / 69.3 (secondary) | 2026, different harnesses | [VENDOR/SECONDARY] |
| Qwen3-0.6B cumulative HF downloads | 22.7M | community leaderboard, Sept 2026 | [COMMUNITY] |

**No cross-version comparison:** the 45 is a v4.3 figure on the 0902 checkpoint; Baseten's "58" for Qwen3.8-Max is the AA percentage-presentation scale, not the 53-point Index scale — never mix the two. [DIRECTIONAL]

### Qwen download scale (September 2026)
- 460+ Qwen models released; **>3 billion cumulative global downloads**; 300,000+ derivative models. [VENDOR via Alibaba press, Reuters-cited]
- First seven months of 2026: **~2.045 billion HF downloads** — roughly **4.9× Google's and 9× Meta's**; 151,448 derivatives (4.7× Meta's Llama count). [VENDOR]
- HF's *State of Open Source: Spring 2026* independently puts China at **41% of downloads**; the often-cited 1.2%→30% OpenRouter figure is **not** from that report (misattribution correction stands). [SECONDARY]


### New verified metrics — expansion

| Item | Value | Date | Provenance |
|---|---|---|---|
| Qwen3.5-397B-A17B params | 397B / 17B active; 512 experts (10+1 active) | 2026-02-15/16 | [SECONDARY] |
| Qwen3.5-397B-A17B context | 262K native, extensible to 1,010,000 | 2026-02 | [SECONDARY] |
| Qwen3.5-397B-A17B languages | 201 languages and dialects | 2026-02 | [SECONDARY] |
| Qwen3.5-397B-A17B decode speed | 8.6–19× faster than Qwen3-Max | 2026-02 | [SECONDARY] |
| Qwen3.5-397B-A17B license | Apache 2.0 | 2026-02 | [SECONDARY] |
| Qwen3.5-397B-A17B MMLU-Pro | 87.8 (vendor claim) | 2026-02 | [SECONDARY] |
| Qwen3.5-397B-A17B AIME26 | 91.3 (vendor claim) | 2026-02 | [SECONDARY] |
| Qwen3.5-397B-A17B LiveCodeBench v6 | 83.6 (vendor claim) | 2026-02 | [SECONDARY] |
| Qwen3.5-397B-A17B local serving | 259 tok/s single-user on 4× RTX PRO 6000 | 2026-06 | [COMMUNITY] |
| Qwen3.6-Max-Preview release | 2026-04-20, closed-weights | 2026-04-20 | [SECONDARY] |
| Qwen3.6-Max-Preview AA Index | 52 | 2026-04 | [SECONDARY] |
| Qwen3.6-35B-A3B SWE-Verified | 73.4% | 2026 | [SECONDARY] |
| Qwen3.6-35B-A3B LiveCodeBench v6 | 80.4% | 2026 | [SECONDARY] |
| Qwen3.7-Max release | Apsara Conference 2026-05-20 | 2026-05-20 | [SECONDARY] |
| Qwen3.7-Max AA Index v4.0 | 56.6 (+4.8 over Qwen3.6-Max-Preview) | 2026-05 | [SECONDARY] |
| Qwen3.7-Max SWE-Verified | 80.4 (vendor claim) | 2026-05 | [SECONDARY] |
| Qwen3.7-Max context | 1M input, 65,536 max output | 2026-05 | [SECONDARY] |
| Qwen3.7-Max pricing | ¥12/¥36 (Model Studio); $2.50/$7.50 (OpenRouter) | 2026-05 | [SECONDARY] |
| Qwen3.7-Max 35-hour demo | >1,000 tool calls, 10× kernel improvement | 2026-05-20 | [SECONDARY] |
| Zhenwu M890 | 3× predecessor; >650 external customers | 2026-05 | [SECONDARY] |
| Panjiu AL128 | 128 accelerators, PB/s bandwidth | 2026-05-20 | [SECONDARY] |
| Qwen3.8-Max params | 2.4T / 95B active | 2026-08-03 | [SECONDARY] |
| Qwen3.8-Max pricing | $2.00/$6.00 per 1M (intl); ¥12/¥36 (China) | 2026-08-03 | [SECONDARY] |
| Qwen3.8-Max PaperBench | 93.0 (vendor claim) | 2026-08 | [SECONDARY] |
| Qwen3.8-Max Terminal-Bench 2.1 | 86.6 (vendor claim) | 2026-08 | [SECONDARY] |
| Qwen3.8-Max DeepSWE 1.1 | 56.6 vs Qwen3.7-Max 21.6 (vendor) | 2026-08 | [SECONDARY] |
| Qwen3.8-Max hardware floor | ~450GB quant; 4.9TB BF16; 72 Blackwell Ultra | 2026-08 | [SECONDARY] |
| Qwen3.8 open vs paid AA Index | 57.7 vs 58.1 (v4.1.1) | 2026-08 | [SECONDARY] |
| Qwen Image 3.0 prompt | 4,500 tokens max | 2026-07-21 | [SECONDARY] |
| Qwen Image 3.0 text | legible to 10px, 12 languages | 2026-07-21 | [SECONDARY] |
| Qwen Image 3.0 pricing | from $0.03/image; Pro i2i $0.075 | 2026-08 | [SECONDARY] |
| Alibaba AI Cloud Q2 2026 | RMB48.44B (+45%); AI >1/3 of cloud | 2026-06-30 | [SECONDARY] |
| Alibaba AI products | RMB12.38B, 12th triple-digit quarter | 2026-06-30 | [SECONDARY] |
| Alibaba capex Q2 2026 | RMB67.68B (+75%) | 2026-06-30 | [SECONDARY] |
| Alibaba net income Q2 2026 | −75% to RMB10.44B | 2026-06-30 | [SECONDARY] |
| Omdia China AI cloud share | 38.1% (#1) | 2025 | [SECONDARY] |
| Qwen app agent shutdown | 2026-07-15, no migration path | 2026-07-15 | [SECONDARY] |
| Qwen3.5-Omni | Thinker-Talker MoE, 256K, 113-lang ASR | 2026-03-30 | [SECONDARY] |
| Qwen3.5-Omni-Plus | 215 SOTA audio/video (vendor claim) | 2026-03-30 | [SECONDARY] |
| Qwen3-Coder training | 7.5T tokens, 70% code, 358 langs | 2026 | [SECONDARY] |
| Qwen3-Coder SWE-bench | 69.6% OpenHands 500-turn; 67.0% standard | 2026 | [SECONDARY] |
| Qwen3-Embedding-8B MTEB | 70.58 multilingual; 80.68 code | 2025-06-05 | [SECONDARY] |
| Qwen3-Embedding throughput | 0.6B 2,100 docs/s; 8B 650 docs/s | 2026 | [COMMUNITY] |
| Qwen3.5-397B-A17B pricing | $0.60/$3.60 per 1M (third-party) | 2026 | [SECONDARY] |
| Qwen3.5-397B-A17B AA Index | 45 (vs Kimi K2.5 46, DeepSeek V3.2 42) | 2026 | [SECONDARY] |
| Qwen3.7-Max HLE | 41.4 (vendor claim) | 2026-05 | [SECONDARY] |
| Qwen3.7-Max Text Arena | #13, Elo 1475 | 2026-05 | [SECONDARY] |
| Qwen3.8-Max params | 2.4T / ~95B active | 2026-08-03 | [SECONDARY] |
| Qwen3.8-Max context | ~1M input, 131K output, 262K reasoning budget | 2026-08-03 | [SECONDARY] |
| Qwen3.8-Max price | $2.00 / $6.00; cache $0.25 / $0.17 | 2026-08-03 | [SECONDARY] |
| Qwen3.8-Max China price | ¥12 / ¥36 | 2026-08 | [SECONDARY] |
| Qwen3.8-Max PaperBench | 93.0 (vendor claim) | 2026-08 | [SECONDARY] |
| Qwen3.8-Max TB 2.1 | 86.6 (vendor claim) | 2026-08 | [SECONDARY] |
| Qwen3.8-Max DeepSWE 1.1 | 56.6 vs Qwen3.7-Max 21.6 (vendor claim) | 2026-08 | [SECONDARY] |
| Qwen3.8-Max AA v4.3 | 40 → 45 (0902) | 2026-08/09 | [SECONDARY] |
| Qwen3.8-Max cost/task | $2.67 → $5.41 (0902) | 2026-09 | [SECONDARY] |
| Qwen3.8 open weights AA v4.1.1 | 57.7 vs 58.1 paid (near tie) | 2026-08-22 | [SECONDARY] |
| Qwen Image 3.0 prompt limit | 4,500 tokens (up from ~1,000) | 2026-07-21 | [SECONDARY] |
| Qwen Image 3.0 benchmarks | none published | 2026-07 | [SECONDARY] |
| Qwen3-Coder-480B SWE-Verified | 69.6% (348/500, OpenHands) | 2025-08 | [SECONDARY] |
| Qwen3-Coder-480B price floor | $0.30 / $1.00 (Deepinfra) | 2026-09-22 | [SECONDARY] |
| Qwen3-Embedding-8B MTEB | 70.58, No.1 multilingual (as of 2025-06-05) | 2025-06 | [SECONDARY] |
| Alibaba AI Cloud revenue | RMB48.44B (~$7.14B), +45% YoY | Q2 2026 | [SECONDARY] |
| Alibaba AI products | RMB12.38B, 12th qtr triple-digit growth | Q2 2026 | [SECONDARY] |
| Alibaba capex | RMB67.68B (~$10B), +75% YoY | Q2 2026 | [SECONDARY] |
| Alibaba net income | RMB10.44B, −75% YoY | Q2 2026 | [SECONDARY] |
| Alibaba AI Labs EBITA loss | −RMB13.86B on RMB3.34B revenue | Q2 2026 | [SECONDARY] |
| Alibaba cloud (Sep report) | $5.6B, +34%; rationing access | Q2 FY2026 | [SECONDARY] |

---

