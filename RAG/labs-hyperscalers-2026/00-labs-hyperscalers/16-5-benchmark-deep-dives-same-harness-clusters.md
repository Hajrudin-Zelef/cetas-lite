---
id: labs-hyperscalers-2026/00-labs-hyperscalers/16-5-benchmark-deep-dives-same-harness-clusters
title: "16.5 Benchmark deep-dives — same-harness clusters"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: benchmark
actors: ["Anthropic", "Google", "Meta", "Microsoft", "Mistral", "OpenAI", "Sakana", "xAI"]
dates: []
keywords: ["benchmark", "agent", "agentic", "agi", "astra", "fable 5", "fugu", "gemini", "gpt-5.6", "gpt-6", "grok", "grok 4"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [2312, 2372]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 1cb7899dda33d46b2ae913b10ba6f0f3ce7d4fc676cd05380be6d6b7f3d32d2b
---

# 16.5 Benchmark deep-dives — same-harness clusters

### 16.5 Benchmark deep-dives — same-harness clusters

**SWE-bench lineage (vendor-reported unless noted; harness versions differ):**
| Model | Benchmark | Score | Provenance |
|---|---|---|---|
| GPT-5.3-Codex | SWE-Bench Pro | 56.8% | [vendor-reported] |
| Muse Spark 1.3 | DeepSWE v1.1 | 75.4% | [vendor-reported] (vs GPT-6 Astra 74.1% same table) |
| GPT-6 Astra | DeepSWE v1.1 | 74.1% | [vendor-reported] |
| Grok 4.5 | SWE-Bench Pro | 64.7% | [vendor-reported] |
| Grok 4.6 | DeepSWE v1.1 | 65.9% | [vendor-reported] |
| Grok 4.6 | FrontierCode v1.1 Ext | 61.3% | [vendor-reported] |
| Grok 4.7 | DeepSWE v1.1 (high-effort) | 71.0% | [vendor-reported] |
| Mistral Medium 3.5 | SWE-Bench Verified | 77.6% | [vendor-reported] |
| Sakana Fugu Ultra v2 | DeepSWE | 74.3 | [vendor-reported] |
| Anthropic Opus 5 | SWE-bench (harness A) | 96–97% | [vendor-reported — conflicts with 74.8% below] |
| Anthropic Opus 5 | SWE-bench (harness B) | 74.8% | [vendor-reported — likely different harness/config; see §18] |

**Terminal-Bench lineage (vendor-reported; 2.0 vs 2.1 vs 3.0 vs 4.0 NOT comparable):**
| Model | Version | Score |
|---|---|---|
| GPT-5.3-Codex | 2.0 | 77.3% |
| GPT-5.5 | 2.0 | 82.7% |
| GPT-5.6 Sol | 2.1 | 88.8% (91.9% ultra) |
| Gemini 3.5 Flash | 2.1 | 76.2% |
| Gemini 3.6 Flash | 2.1 | 78.0% |
| Grok 4.5 | 2.1 | 83.3% |
| Grok 4.6 | 3.0 | 26% |
| GPT-6 Astra | 4.0 | 57.9% |
| Fable 5 (vendor) | vendor harness | 88.0% |
| Fable 5 (Vals AI independent) | independent re-run | 80.52% |

**Agentic / computer-use (vendor-reported):**
| Model | Benchmark | Score |
|---|---|---|
| GPT-5.5 | OSWorld-Verified | 78.7% |
| GPT-6 Astra | OSWorld 2.0 | 72.6% |
| Gemini 3.6 Flash | OSWorld-Verified | 83.0% |
| Grok 4.3-era | AA-Briefcase Elo | 1,546 (4.20) → 1,546-class |
| Grok 4.6 | AA-Briefcase Elo | 1,577 |
| Grok 4.7 | AA-Briefcase Elo | 1,657 |
| Grok 4.5 | AutomationBench-AA | 51% (#1 at launch) |

**Reasoning / knowledge (vendor-reported):**
| Model | Benchmark | Score |
|---|---|---|
| Gemini 3.1 Pro | ARC-AGI-2 | 77.1% |
| Gemini 3.5 Flash | ARC-AGI-2 | 77.1% |
| Gemini 3.5 Flash | MCP Atlas | 83.6% |
| Gemini 3.5 Flash | CharXiv Reasoning | 84.2% |
| Gemini 3.6 Flash | CharXiv Reasoning | 85.2% (no tools) / 89.4% (with tools) |
| Gemini 3.6 Flash | GDM-MRCR v2 | 91.8% (128k) / 54.0% (1M) |
| MAI-Thinking-1 | AIME 2025 / 2026 | 97.0% / 94.5% |
| Mistral Small 4 | GPQA Diamond | 71.2% |
| Mistral Small 4 | MMLU-Pro | 78.0% |
| Grok 4.6 | GDPVal-AA v2 Elo | 1,753 |
| Gemini 3.6 Flash | GDPVal-AA v2 Elo | 1,421 |

**Coding-agent leaderboards (independent, dated):**
- Artificial Analysis Coding Agent Index: Grok 4.7 = 56 (+9 vs 4.6) on Sep 21, 2026 first pass. [independent]
- CursorBench: Grok 4.6 v3.2 69.9% → Grok 4.7 v4.0 46.3% (version change; not comparable). [vendor-reported]

