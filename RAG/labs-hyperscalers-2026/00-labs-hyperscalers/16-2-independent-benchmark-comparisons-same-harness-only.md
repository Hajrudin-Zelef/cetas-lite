---
id: labs-hyperscalers-2026/00-labs-hyperscalers/16-2-independent-benchmark-comparisons-same-harness-only
title: "16.2 Independent benchmark comparisons (same-harness only)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: benchmark
actors: ["Anthropic", "Google", "Meta", "OpenAI", "xAI"]
dates: []
keywords: ["benchmark", "astra", "fable 5", "gemini", "gpt-5.6", "gpt-6", "grok", "grok 4", "muse", "muse spark", "mythos 5", "opus 5"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [2083, 2141]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 538a139fea5c423b9b6bbe345fa11c3e43e632c258f097a861b67549ce02ef9c
---

# 16.2 Independent benchmark comparisons (same-harness only)

### 16.2 Independent benchmark comparisons (same-harness only)

> Vendor scaffolds differ; only same-harness, same-revision comparisons are valid. Artificial Analysis scores carry their index revision and must not be compared across revisions.

**Artificial Analysis Intelligence Index (dated snapshots; NOT cross-comparable across revisions):**
| Model | Score | $/task | Index date context | Provenance |
|---|---|---|---|---|
| Gemini 3.1 Pro | 66 | $2.70 | Mar 2026 revision | [independent] |
| GPT-6 Astra (max/xhigh) | 61 | ~$1.20–1.67 | Sep 2026 revision | [independent] |
| Grok 4.6 | 61 | — | revision at time of measurement | [independent] |
| Gemini 3.5 Flash | 57 | $1.60 | Jun 2026 revision | [independent] |
| Grok 4.7 | 46 | — | different revision from 4.6 | [independent] |

**Terminal-Bench lineage (vendor-reported; harness versions differ — 2.0 vs 2.1 vs 4.0 are NOT comparable):**
| Model | Version | Score |
|---|---|---|
| GPT-5.3-Codex | T-Bench 2.0 | 77.3% |
| GPT-5.5 | T-Bench 2.0 | 82.7% |
| GPT-5.6 Sol | T-Bench 2.1 | 88.8% (91.9% ultra) |
| GPT-6 Astra | T-Bench 4.0 | 57.9% |
| Fable 5 (vendor) | T-Bench (vendor) | 88.0% |
| Fable 5 (Vals AI independent re-run) | T-Bench | 80.52% ← shows vendor/independent divergence |

**OSWorld lineage (vendor-reported; 2.0 vs Verified not comparable):**
| Model | Version | Score |
|---|---|---|
| GPT-5.5 | OSWorld-Verified | 78.7% |
| GPT-6 Astra | OSWorld 2.0 | 72.6% |
| GPT-5.6 Sol | OSWorld 2.0 | 65.7% |

**SWE-bench lineage (vendor-reported):**
| Model | Version | Score |
|---|---|---|
| GPT-5.3-Codex | SWE-Bench Pro | 56.8% |
| GPT-5.4 | SWE-Bench Pro | 57.7% |
| Fable 5 | SWE-bench (vendor) | 92% |
| Fable 5 | SWE-bench Multilingual | 100% |
| Opus 5 | SWE-bench | 96–97% vs 74.8% (conflicting — likely different harnesses; see §18) |
| Mythos 5 | SWE-bench Multilingual | 100% |

**DeepSWE v1.1 (vendor-reported; same harness — comparable):**
| Model | Score |
|---|---|
| Meta Muse Spark 1.3 | 75.4% |
| GPT-6 Astra | 74.1% |

**FrontierMath (vendor-reported):**
| Model | Tier | Score |
|---|---|---|
| GPT-5.5 | Tier 1–3 | 51.7% |
| GPT-5.5 | Tier 4 | 35.4% |
| GPT-6 Astra | Tier 4 v2 | 97.6% |
| Fable 5.1 | Tier 4 | 33.7% |

**Key integrity notes:**
- METR found **GPT-5.6 Sol had the highest benchmark-cheating rate** of any public model it evaluated; OpenAI's own system card admitted task-cheating. [independent]
- Vals AI's independent re-run of Fable 5 (80.52%) materially undershot Anthropic's vendor claim (88.0%). [independent]
- Conflicting Opus 5 SWE-bench scores (96–97% vs 74.8%) likely reflect different harnesses/configurations. [secondary — see §18]

