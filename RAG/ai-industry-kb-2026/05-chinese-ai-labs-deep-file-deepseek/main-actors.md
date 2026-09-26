---
id: ai-industry-kb-2026/05-chinese-ai-labs-deep-file-deepseek/main-actors
title: "Main actors"
domain: chinese-ai-labs-deep-file-deepseek
role: deep-dive
task: actor-profile
actors: ["China", "DeepSeek", "SGLang", "vLLM"]
dates: ["2024-12-26", "2025-01-20", "2025-03-25", "2025-05-28", "2025-08-21", "2025-09-22", "2025-09-29", "2025-12-01", "2026-02", "2026-02-11", "2026-02-15", "2026-03", "2026-03-09", "2026-04-24", "2026-06", "2026-06-02", "2026-06-15", "2026-08-12", "2026-09-10"]
keywords: ["apache", "attention", "attribution", "benchmark", "benchmarks", "cost", "cybersecurity", "deepseek", "fp4", "fp8", "inference", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1953, 2001]
section: "5. Chinese AI Labs — Deep File: DeepSeek"
sha256: 3694d57b8959b01542c19da519d90db8230cfc3daf4c5649676f7b20dce6bac3
---

# Main actors

## Main actors

- **DeepSeek** (Hangzhou lab): releases V4-Lite, V4 Preview, V4-Pro-0813, V4.1-Flash; writes the API changelog and pricing pages used as primary-adjacent evidence; ships MIT weights on HF/ModelScope. In 2026 it ships first and explains later — stealth GA, expiring-id beta, website statement removed next day.
- **Peking University**: named in 36kr's March reporting as DeepSeek's co-developer of the Engram conditional memory system — rumor-phase attribution, never confirmed by DeepSeek material; appears again in V4.1 secondary coverage of the 196B module.
- **36kr** (Chinese tech outlet): origin of the V4-Lite ~200B figure and the Engram March-2026 narrative (relayed by awesomeagents); the sole source for the 200B figure — no official count exists.
- **Vals.ai**: dated pre-release evaluation of V4-Pro-0813 (2026-08-12) — the primary anchor for the +9.48-point figure and SWE-bench Verified 96.40%; #18 at eval time, 12th in mid-September press (live-index rank drift).
- **Aikido Security** (Belgian cybersecurity firm): niche finding that V4-Pro-0813 led all tested models in vulnerability-detection count, with poor precision — reported via coinlive, SCMP, it-daily.
- **Digitimes**: reported (June 2, 2026) the permanent 75% flagship API price cut from June 1 — the dating anchor that unwinds the brief's "June 1 = V4.1" conflation.
- **Artificial Analysis** (Intelligence Index): independent composite; V4-Pro-0813 at 53 (v4.1-era, +8 over April preview) vs ~36 on the rebased v4.3 scale — the methodology-version discipline the RAG must keep.
- **cellcog**: reconstructed DeepSeek's release cadence (V4 Preview → Flash-0731 → Pro-0813 → Flash-Vision-Exp → V4.1-Flash) and the fastest-structural-update interval read; dates the expiring-id beta as a DeepSeek first.
- **buzzgrewal** (Medium analysis): documented that the official 58-page V4 report mentions "Engram" and "O(1)" zero times; laid out the CSA/HCA hybrid attention design — the evidence that kills the rumor-phase memory narrative for V4.
- **oliver-foster (Medium)** and **getshint**: covered the V4.1-Flash beta mechanics — the <100-word official user-group notice and the expiring `deepseek-v4.1-flash-expires-on-0910` probe id; the tell that this was a probe, not a product.
- **davidborish.com / magicshot.ai / APIMaster**: GA-day coverage — vendor claims (fourfold memory-cost cut, "Stronger, Faster, More Accessible"), $0.15/$0.60 pricing, and the conflicting V4-Pro retirement notices (APIMaster carries the revised notice keeping V4 Pro service unchanged).
- **orcarouter.ai**: day-0 inference-stack analysis (vLLM preview container Sept 9; SGLang + Miles day-0 Sept 10; Ollama within a day) with the caution that day-0 "support" was preview-plus-integration, not stable pip releases.
- **miraflow.ai**: synthesis of the 0813 GA (AA-index +8, ~3.6× price-increase framing [DIRECTIONAL]) and a CED explainer for V4.1.
- **SCMP / coinlive / nationpress / it-daily**: mid-September coverage — stealth GA mechanics (statement removed next day), 12th-on-Vals-Index rank drift, cybersecurity finding, pricing criticism ("somewhat expensive" vs 33-cent market input median).
- **techi.com**: documented the deepseek-chat/reasoner alias retirement (15:59 UTC cutoff) and migration behavior; Void editor docs and gokin provider integration corroborate.
- **ai-pricelog / roninforge / BenchLM**: changelog and price-backfill infrastructure carrying DeepSeek's own announcement text and list prices ($0.435/$0.87 from June 1; MIT license directory listing).
- **nxcode / ez0000001000000 / kissapi.ai / geeky-gadgets / aiforautomation.io**: March 2026 rumor/leak-phase coverage — the NDA/soft-launch framing and (since-retired) Apache 2.0 speculation; nxcode's March-5 FAQ describing V4 as "just landed" exemplifies trackers posting ahead of primary confirmation.
- **yorozuipsc** (Japanese enterprise report, updated June 2026): uses only DeepSeek's official news page as source of truth — "the latest model release is DeepSeek-V4 Preview, 2026-04-24" — a useful primary-anchoring example.
- **vLLM / SGLang / Miles / Ollama**: day-0 inference ecosystem for V4.1-Flash; SGLang + Miles joint post unusually detailed about which architecture parts are expensive.
- **hermes-workspace (June 15, 2026 scout)**: secondary spec sheet — "1.6T MoE, 49B active, true 1M context ... MIT license, FP4/FP8 mixed precision, 384K output window" — confirms the V4 line's mid-2026 shape.
- **Techpresso**: documented the DeepSWE harness sensitivity (65.5%–74.2% swing on the same V4.1-Flash checkpoint) — the independent caution on vendor launch benchmarks.
- **flowtivity.ai**: September GA-week benchmark coverage of V4.1-Flash.
- **temperaturezero.com** (2026-09-10): V4.1-Flash inference and KV-cache benchmark analysis — the 890 B/token claim's secondary technical read.
- **aicraftjournal.com**: GA-week coverage of peak pricing ($0.30/$1.20), MIT license, and the Sept-14 V4-Pro question.
- **gotechpresso.com**: GA-week V4.1-Flash coverage.
- **miraflow.ai (CED explainer)**: Causal Encoder-Decoder explained for the 2026 V4.1 release — the CED-expansion anchor.
- **datanorth.ai**: V4.1-Flash release coverage.
- **aiwith.me**: V4.1-Flash launch coverage.
- **gnana70 (Medium)**: V4.1-Flash architecture/benchmarks/real-cost analysis.

## Timeline and context

### Pre-2026 lineage (context anchor)

- 2024-12-26 — DeepSeek-V3 GA. 2025-01-20 — DeepSeek-R1 GA. 2025-03-25 — V3-0324 checkpoint. 2025-05-28 — R1-0528 checkpoint. 2025-08-21 — V3.1 GA. 2025-09-22 — V3.1-Terminus checkpoint. 2025-09-29 — V3.2-Exp. 2025-12-01 — V3.2 GA.
- **2026-02-11** — Silent production upgrade: V3.2-class model context 128K → 1M tokens, detected by the community (awesomeagents) — the start of DeepSeek's 2026 "quiet rollout" pattern.

### Phase 1 — Rumor (February 2026)

- **~2026-02-15**: the first V4 mentions circulate — naming confusion, multiple release windows (mid-February, Lunar New Year, early March) pass with no launch (nxcode FAQ). Secondary blogs speculate the flagship at ~1T params, Apache 2.0 license, Engram "O(1) memory" with "97% needle-in-a-haystack at 1M" and "1M context costs roughly the same as 128K" (buzzgrewal's summary of the rumor cycle). **State clearly: this is a pre-reporting/rumor phase, NOT a release** — no model, no weights, no price, no license exists yet.
- The "97% needle-in-a-haystack at 1M" and "1M context costs roughly the same as 128K" claims are rumor-cycle statements — they never appear in any DeepSeek-dated material.
- The ~1T-param flagship speculation of February (kissapi.ai-style coverage) never matches the April reality (1.6T Pro / 284B Flash) — rumor-phase scale guesses are residue, not precursors.
- Consolidation rule from this phase: any undated figure (param count, license, architecture term) traced to February–March 2026 secondary coverage is rumor-phase residue until anchored to a DeepSeek-dated announcement.
- "Compressed Expert Dispatch" is most plausibly an AI-generated expansion of "CED" produced during this rumor cycle (quoted-phrase search: zero relevant hits) — the consolidation kill-list starts here.

### Phase 2 — Leak / soft-preview (2026-03-09)

