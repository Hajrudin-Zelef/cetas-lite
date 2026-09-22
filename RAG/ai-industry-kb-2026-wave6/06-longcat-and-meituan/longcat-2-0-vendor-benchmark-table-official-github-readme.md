---
id: ai-industry-kb-2026-wave6/06-longcat-and-meituan/longcat-2-0-vendor-benchmark-table-official-github-readme
title: "LongCat-2.0 — vendor benchmark table (official GitHub README)"
domain: longcat-and-meituan
role: deep-dive
task: benchmark
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "LongCat", "Moonshot", "Nvidia", "OpenAI", "Z.ai"]
dates: ["2025-09-22", "2026-01-14", "2026-06-12", "2026-06-29", "2026-06-30", "2026-08-28"]
keywords: ["benchmark", "agent", "apache", "asic", "claude", "decode", "deepseek", "embedding", "gemini", "glm", "kimi", "leaderboard"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2824, 2910]
section: "§6. LongCat and Meituan"
delta_of: ai-industry-kb-2026
sha256: edbcf6ff1275c88255f106c258894ec145507a51805e705b9b4b00acb24b59dd
---

# LongCat-2.0 — vendor benchmark table (official GitHub README)

### LongCat-2.0 — vendor benchmark table (official GitHub README)
- All LongCat-2.0 scores below measured **in-house under a unified harness** unless marked `*` (= cited from the compared model's official report); `-` = no comparable public score [VENDOR] (github.com/meituan-longcat/LongCat-2.0).
- **Code Agent**: Terminal-Bench **2.1** — 70.8 (vs GPT-5.5 73.8*, Claude Opus 4.7 71.7*, Claude Opus 4.8 78.9*, Gemini 3.1 Pro 70.7*); SWE-bench **Pro** — 59.5 (vs GPT-5.5 58.6*, Claude Opus 4.6 57.3*, Opus 4.7 64.3*, Opus 4.8 69.2*, Gemini 3.1 Pro 54.2*); SWE-bench **Multilingual** — 77.3 (vs Opus 4.6 77.8*, Opus 4.7 80.5*, Opus 4.8 84.8*, Gemini 3.1 Pro 76.9*) [VENDOR].
- **General Agent**: FORTE 73.2 (vs GPT-5.5 77.8, Opus 4.8 77.2); BrowseComp 79.9 (vs GPT-5.5 84.4*, Opus 4.8 84.3*); RWSearch 78.8 (vs GPT-5.5 85.3, Opus 4.8 77.3) [VENDOR].
- **Foundational**: IFEval 90.0; Writing Bench 83.8; IMO-AnswerBench 81.8; GPQA-diamond 88.9 (vs GPT-5.5 93.6*, Gemini 3.1 Pro 94.3*) [VENDOR].
- The **Terminal-Bench version question is now resolved: 2.1** per the vendor README [VENDOR] — do not carry the unversioned "Terminal-Bench 70.8" phrasing forward.
- SWE-bench **Pro 59.5 edges GPT-5.5's 58.6** (vendor's in-house run vs GPT-5.5's official-report figure — note the asymmetric provenance before citing as a "win") [VENDOR with caveat].

### LongCat-2.0 — secondary benchmark and long-context figures
- Long-context retrieval: **Needle-In-A-Haystack at 1M tokens maintains >94% retrieval up to the full million tokens**, described as rivaling Gemini 2.0 Flash (2M proprietary) out to 500K [SECONDARY] (ayinedjimi-consultants.fr PDF — single source, mark [UNVERIFIED] for procurement).
- Secondary comparison table (methodology unattributed, treat as [SECONDARY] with unclear provenance): LongCat-2.0 — MMLU **88.4%**, HumanEval+ **91.7%**, GPQA **68.2%**, MATH **84.6%**, 1M context; vs DeepSeek-V4-Pro (88.1/89.4/70.1/87.3), Kimi K2.7 (87.6/90.2/66.8/83.9), GLM-5.2 (85.9/87.1/63.4/80.7), Llama 4 405B (86.2/85.8/64.1/82.4), Qwen3 235B MoE (87.1/88.6/65.9/83.1) — all 128K context [SECONDARY] (ayinedjimi-consultants.fr PDF).
- Reading of that table: LongCat-2.0 leads the open-source group on HumanEval+ (coding) and is the only 1M-context entry; DeepSeek-V4-Pro keeps a small edge on GPQA and MATH (pure reasoning) [DIRECTIONAL].

### Pricing and platform
- API pricing: **~$0.75/M input / $2.95/M output** tokens — positioned as undercutting GPT-5.5 and Claude Sonnet 5 [SECONDARY] (cryptobriefing.com, 2026-07).
- Platform changelog: **https://longcat.chat/platform/docs/ChangeLog.html** is the vendor's running API/platform change log [VENDOR] (longcat.chat).
- Existing §6 records the **May 29 six-model sunset** and **Owl Alpha**; the exact six-model retirement list was not recovered in this research pass — flagged as a remaining gap, not filled here [DIRECTIONAL].

## Figures and metrics

| Model | Date | Size (total/active) | Context | License | Provenance |
|---|---|---|---|---|---|
| LongCat-Flash-Thinking | 2025-09-22/23 | — | — | — | [VENDOR] |
| LongCat-Flash-Thinking-2601 | 2026-01-14 (changelog) | — | — | — | [VENDOR] |
| LongCat-Flash-Omni | ~2025-11 | 560B / 27B | 128K | MIT | [VENDOR] |
| LongCat-2.0 | 2026-06-30 (reveal); 07-05 (weights) | 1.6T / ~48B | 1M | MIT | [VENDOR] |
| Tencent Hy4 Preview | 2026-08-28 | 770B / 49B | >1M | Apache 2.0 | [VENDOR] |
| openPangu 2.0 Pro | 2026-06-12 (announce); 06-30 (release) | 505B / 18B | 512K | not verified | [SECONDARY] |
| openPangu 2.0 Flash | 2026-06-12 (announce); 06-30 (release) | 92B / 6B | 512K | not verified | [SECONDARY] |
| Owl Alpha (stealth) | revealed 2026-06-29 | — | — | — | [VENDOR] |

| Metric | Figure | Provenance |
|---|---|---|
| Owl Alpha reported traffic | ~10.1 trillion tokens/month | [SECONDARY] |
| LongCat-2.0 training hardware (claimed) | 50,000+ domestic accelerators, no NVIDIA | [VENDOR] (unverified) |
| 2026 Apache-2.0 releases above 190B params | Hy4 770B · openPangu Pro 505B · Step 3.7 Flash 198B · dots3-note 280B | [SECONDARY] |


### New verified metrics — expansion (continued — LongCat-2.0 architecture figures)

- Zero-computation experts + PID bias control; 33B–56B dynamic (avg 48B); ScMoE [SECONDARY].
- LSA: 3 indexing methods (streaming-aware, cross-layer, hierarchical); evolution of DSA [SECONDARY].
- 135B N-gram module; MOPD 3-teacher fusion; 6D parallelism; CPP/SP prefill, KVP/large EP decode [SECONDARY].
- Muon customized at scale; all-gather CP for 1M ctx; async load balancing [SECONDARY].
- Training tokens: 35T+ (tamago) vs "over 30T" (cryptobriefing) — variance recorded [SECONDARY].


### New verified metrics — expansion (continued — 2601/ZigZag figures)

- 2601: 560B / 27B avg active; AIME-25 up to 100.0; LCB 82.8; τ²-Telecom 99.3; BrowseComp-zh up to 77.7 [VENDOR/SECONDARY].
- ZigZag: ~1.5× e2e speedup; 1M context via YaRN; 1024-token/layer effective span [SECONDARY].
- MLX 5.5-bit: ppl 1.141; 23 tok/s single on M3 Ultra 512GB [COMMUNITY].


### New verified metrics — expansion

### LongCat-2.0 vendor benchmark table (official README, [VENDOR])
| Benchmark | LongCat-2.0 | Best cited rival |
|---|---|---|
| Terminal-Bench **2.1** | 70.8 | Claude Opus 4.8 78.9* |
| SWE-bench **Pro** | 59.5 | Claude Opus 4.8 69.2* / GPT-5.5 58.6* |
| SWE-bench Multilingual | 77.3 | Claude Opus 4.8 84.8* |
| FORTE | 73.2 | GPT-5.5 77.8 |
| BrowseComp | 79.9 | GPT-5.5 84.4* |
| RWSearch | 78.8 | GPT-5.5 85.3 |
| IFEval | 90.0 | Gemini 3.1 Pro 96.1 |
| Writing Bench | 83.8 | Claude Opus 4.7 85.3 |
| IMO-AnswerBench | 81.8 | Gemini 3.1 Pro 90.0 |
| GPQA-diamond | 88.9 | Gemini 3.1 Pro 94.3* |

(`*` = cited from the rival's official report; LongCat-2.0 figures measured in-house under a unified harness. Keep this provenance split visible — it is not a like-for-like leaderboard.)

### Pricing
- **$0.75/M input / $2.95/M output** — undercut positioning vs GPT-5.5 and Claude Sonnet 5 [SECONDARY] (cryptobriefing.com).
- Output/input price ratio ≈ **3.9×** [DIRECTIONAL].

### Long-context and secondary figures
- NIAH @ 1M tokens: **>94% retrieval** to the full million [SECONDARY, single source] (ayinedjimi-consultants.fr PDF).
- Secondary table: MMLU 88.4 / HumanEval+ 91.7 / GPQA 68.2 / MATH 84.6 [SECONDARY, unattributed methodology] (ayinedjimi-consultants.fr PDF).
- LongCat-Flash (original): MMLU **89.71** [VENDOR]; throughput **>100 tok/s** [SECONDARY]; 560B / 18.6–31.3B dynamic (avg ~27B) / 128K ctx [SECONDARY].
- LongCat-Flash-Prover: MiniF2F-Test **97.1%** (72 attempts); ProverBench **70.8%**; PutnamBench **41.5%** (≤220 attempts) [VENDOR].

### Scale figures
- **1.6T total / 33–56B dynamic active (avg ~48B)**; **1M native context**; MIT license [SECONDARY] (marktechpost.com; cryptobriefing.com).
- **50,000-card domestic ASIC cluster, zero NVIDIA** [VENDOR via secondary] (cryptobriefing.com; marktechpost.com).
- **135B-parameter N-gram embedding module** (~100× embedding-space expansion) [SECONDARY] (venturebeat.com).

