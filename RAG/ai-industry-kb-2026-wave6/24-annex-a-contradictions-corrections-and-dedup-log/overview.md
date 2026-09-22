---
id: ai-industry-kb-2026-wave6/24-annex-a-contradictions-corrections-and-dedup-log/overview
title: "Annex A — Contradictions, Corrections, and Dedup Log"
domain: appendix
role: appendix
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "ByteDance", "Cohere", "DeepSeek", "EU", "Falcon", "Google", "Groq", "Hugging Face", "LongCat", "Meituan", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Poolside", "SpaceX", "Z.ai", "vLLM", "xAI"]
dates: ["2025-06", "2025-08", "2025-08-13", "2025-09-22", "2025-09-30", "2025-12", "2025-12-09", "2025-12-31", "2026-01", "2026-01-05", "2026-01-14", "2026-02-20", "2026-02-27", "2026-03-13", "2026-04", "2026-04-07", "2026-04-20", "2026-04-22", "2026-04-26", "2026-04-28", "2026-04-29", "2026-05-07", "2026-06-01", "2026-07-02", "2026-07-09", "2026-07-10", "2026-08-20", "2026-08-28", "2026-09", "2026-09-01", "2026-09-03", "2026-09-20", "2026-09-21", "2026-09-22", "2026-09-24", "2026-10-23", "2026-10-27"]
keywords: ["acquisition", "agi", "apache", "arr", "astra", "aws", "benchmark", "claude", "cohere", "compute", "consumer", "context window"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [11547, 11695]
section: "Annex A — Contradictions, Corrections, and Dedup Log"
sha256: 19e4be10ccedb729dd1b5446bea29e56eb8908c089fa45d2b84672c53710ff8d
---

# Annex A — Contradictions, Corrections, and Dedup Log

This annex records every contradiction found across the five wave-6 source files, the resolution applied during consolidation, retracted or non-existent entries, and the dedup journal (what was dropped, merged, or reduced to a pointer).

## A.1 Contradictions resolved

| # | Claim(s) in conflict | Resolution applied | Provenance |
|---|---|---|---|
| 1 | DeepSeek's CED = "Compressed Expert Dispatch" vs "Causal Encoder-Decoder" | "Causal Encoder-Decoder" is confirmed by DeepSeek's own Hugging Face README. "Compressed Expert Dispatch" is a plausible-but-false expansion — rejected everywhere. | [VENDOR] |
| 2 | DeepSeek V4-Lite 200B "release" (Oct 2025) | Never shipped. SEO-bait noise; carried only as [COMMUNITY] rumor in §2, excluded from the §23 inventory. | [COMMUNITY] |
| 3 | DeepSeek R2 "released April 2026, 32B, MIT, AIME 92.7%" | **Not released as of 2026-09-22.** The April-2026-release narrative is classified as fabrication-risk [UNVERIFIED]; owned by §2 with a one-line quarantine. | [UNVERIFIED] |
| 4 | GLM-5.1 date: 2026-04-07 vs 2026-05-07 | 2026-04-07 retained (stray May-7 line rejected). Recorded as unresolved-adjacent in the §23 row. | [SECONDARY] |
| 5 | GLM-5.2 vs GLM-5.3 parameter counts: 744B / 743B / 753.3B | Same weights, three countings: vendor headline (744B), vLLM config rounding (743B), Hugging Face tensor count (753.3B). Kept as one base in §4. | [VENDOR/SECONDARY] |
| 6 | MiniMax M3 date: 2026-03-13 vs 2026-06-01 | 2026-06-01. March-13 claims collide with the M2.7 rollout; M2 outlier date 2026-10-23 corrected to 2026-10-27. | [VENDOR] |
| 7 | Mistral Large 3 date: Dec 2025 vs Apr 2026 | **December 2025.** The real April 2026 release is Medium 3.5 (2026-04-29/30); Medium 3.1 is an August 2025 refresh. | [VENDOR] |
| 8 | Devstral 2 date: 2026 vs Dec 2025 | **2025-12-09/10** (+ Devstral Small 2, Vibe CLI). 2026 dating rejected. | [VENDOR] |
| 9 | Kimi K2.6 date: Apr 13 vs Apr 20 2026 | **2026-04-20.** Apr-13 variant rejected. | [VENDOR] |
| 10 | LongCat Flash-Thinking origin: Jan 2026 vs Sep 2025 | Original release **2025-09-22/23**; January 2026 is the Thinking-2601 refresh (canonical effective date **2026-01-14** per the Meituan changelog, vs Jan-16 coverage). | [VENDOR/SECONDARY] |
| 11 | "Owl Alpha" token volume: billions vs trillions | **~10.1 trillion tokens/month** (community-observed stealth volume), not billions. | [COMMUNITY] |
| 12 | MiMo V2.5-Pro vs V2.6-Pro (both 1.02T/42B) | Two **distinct** generations sharing the parameter profile: V2.5-Pro (2026-04-22, 1M context) and V2.6-Pro (2026-09-21/22, "Live RL", AA v4.3 46.32). | [VENDOR] |
| 13 | Grok Code Fast 1 date: 2026 vs Aug 2025 | **August 2025.** 2026 dating rejected. | [VENDOR] |
| 14 | Grok 5 availability | **Not shipped as of 2026-09-22.** Carried as [UNVERIFIED] only. | [UNVERIFIED] |
| 15 | Nemotron 4 (1T) status | Reported **in development** (The Information/Reuters, Aug 11 2026) — not announced, no date. | [SECONDARY] |
| 16 | "No Falcon releases since Falcon 2" | **Contradicted:** Falcon-H1R/H1 Arabic (2026-01-05), Falcon Perception + OCR (~2026-04), Falcon Perception-300M (2026-07) all shipped. | [SECONDARY] |
| 17 | "NVIDIA acquired Poolside" | **Contradicted:** $6B non-exclusive license of the **Model Factory** + $1B equity (2026-08-20) — "not an acquisition and not an acquihire"; Laguna itself was not licensed. | [SECONDARY] |
| 18 | Sora 2 date: 2026 vs 2025 | **2025-09-30.** Consumer Sora wound down 2026-04-26; Sora API decommissioned 2026-09-24 (single-source-family, needs corroboration). | [VENDOR/SECONDARY] |
| 19 | Runway Gen-5 / ElevenLabs v4 | Gen-5 **not shipped** as of 2026-09-22 (Gen-4.5 Dec 2025 is the flagship); ElevenLabs v4 **previewed only** (ElevenSummit Warsaw, Jun 2026). | [SECONDARY] |
| 20 | ByteDance Seed-Coder date: 2026 vs 2025 | **June 2025.** 2026 dating rejected. | [SECONDARY] |
| 21 | "Ministral 3" (2026) | Single-source, [UNVERIFIED]; **not** merged with the October-2024 Ministral 3B/8B line. | [UNVERIFIED] |
| 22 | Jina v5 model sizes | small = **677M**, nano = **239M**; several press writeups inverted the two. CC BY-NC 4.0. | [VENDOR] |
| 23 | Anthropic's two $30B figures | Kept strictly separate: Nov 2025 **Azure $30B compute commitment** vs Feb 2026 **$30B Series G**. | [SECONDARY] |
| 24 | NVIDIA → OpenAI $30B announcement | Formal announcement **2026-02-27** at a $730B pre-money; the 2026-02-20 FT item was the leak report, not the announcement. | [SECONDARY] |
| 25 | Ideogram 4 "Apache 2.0" | Shorthand contradicted: **code/pipeline** Apache 2.0; **weights** under a non-commercial Ideogram agreement. | [VENDOR] |
| 26 | Hunyuan3D 3.0 "open" | Free **hosted** access — not open weights. | [VENDOR] |
| 27 | Terminal-Bench 4.0 timing | Announced **2026-08-28** (tbench.ai); leaderboard circulation **early September 2026**. Both milestones kept; TB 4.0 never compared with TB 2.1. | [VENDOR] |
| 28 | Ollama 0.19 MLX gains (+57% prefill / +93% decode) | Measured specifically on **Qwen3.5-35B-A3B** (~3B active, NVFP4) — not generalized to dense 35B models. | [VENDOR] |
| 29 | "Claude 4.62" | Zero hits across sources; treated as a garbled "4.6" data-entry error — dropped from §§13 and §23. | [CONTRADICTED] |
| 30 | Gemini 2.0 Ultra | Never existed — no verified release record; dropped everywhere. | [CONTRADICTED] |
| 31 | "OpenAI shelved flagship Astra" (Medium, Aug 7 2026) | [UNVERIFIED]; conflicts with the **2026-09-03 GPT-6 Astra** launch. Not used as fact. | [UNVERIFIED] |
| 32 | Midjourney V8.1 alpha date | Apr 14 vs Apr 30 2026 — both kept, flagged as a discrepancy. | [SECONDARY] |
| 33 | Cohere Rerank 4 context | 32K vs 4K across sources — both kept, flagged. | [SECONDARY] |
| 34 | ERNIE 5.0/5.1 pricing | $0.85/$3.40 vs $0.59/$2.65 — unresolved; presented as-is, source unclear whether version-driven or market-driven. | [SECONDARY] |
| 35 | MiMo-V2.6-Pro AA 46.32 vs Qwen3.8-Max 45 | Two [SECONDARY] snapshots from different dates — kept separate, not merged into one ranking. | [SECONDARY] |
| 36 | Terminal-Bench 4.0 leader | Scale-official (Fable 5.1 57.9%) vs BenchLM vs alextech snapshots — kept separate with attributions. | [SECONDARY/VENDOR] |

## A.2 Retracted / non-existent entries

The following names appeared in source material but are **not real releases** as of 2026-09-22 and are excluded from the §23 master inventory (or marked non-existent there):

- **Gemini 2.0 Ultra** — never existed (no release record). [CONTRADICTED]
- **Claude 4.62** — zero hits; garbled "4.6". [CONTRADICTED]
- **DeepSeek V4-Lite 200B** — SEO noise; never shipped. [COMMUNITY]
- **DeepSeek R2** — not released as of 2026-09-22. [UNVERIFIED]
- **Grok 5** — not shipped as of 2026-09-22. [UNVERIFIED]
- **Nemotron 4 (1T)** — in development only, not announced. [SECONDARY]
- **Gemini 3.5 Pro** — internal/unreleased only. [SECONDARY]
- **Gemini Omni Pro** — teased only. [UNVERIFIED]
- **Runway Gen-5** — not shipped as of 2026-09-22. [SECONDARY]
- **ElevenLabs v4** — previewed only, not GA. [SECONDARY]
- **Phi-5** — no verified release; no named 2026 Microsoft coding model. [UNVERIFIED]
- **Ministral 3 (2026)** — single-source, not merged with the 2024 line. [UNVERIFIED]
- **Gemma TTS / Gemma 4.5** — phantom; dropped. [CONTRADICTED]
- **"Preview Vision" (Gemini 3.1)** — no separate product; folded into the 3.1 Pro Preview line. [SECONDARY]
- **qwen4_exp** — experimental checkpoint, not a Qwen 4 announcement. [COMMUNITY]
- **PlayAI / PlayHT** — dead: shut down 2025-12-31 after Meta's July-2025 acqui-hire; Groq retired the engine. [SECONDARY]

## A.3 Dedup journal

What was dropped, merged, or reduced to a pointer during consolidation — and where the surviving detail lives:

1. **DeepSeek R2 / V4 chronology** — owned by §2. §§17, 19, 20, 23 carry at most one line each.
2. **MiniMax M3 full detail** — owned by §7. §§19, 22, 23 carry one-liners / table rows.
3. **MiMo V2.6-Pro full detail** — owned by §8. §22 carries the row; §23 the row.
4. **License deep-map** — owned by §21. §§2–20 carry per-model license one-liners only; §9 quotes the Llama EU clause verbatim (the clause's primary citation home).
5. **Pricing table** — owned by §22. §§2–20 carry isolated price facts; none carry the full table.
6. **Tencent Hy4 Preview + openPangu 2.0** — §6's deep blocks were merged into §17 (announce/release dates, license split, co-design framing, the Apache-2.0-above-190B metric); §6 keeps a compact pointer plus its metrics-table rows.
7. **July-2026 coalitions** (July 24 open-weights letter; July 27 Open Secure AI Alliance) — full treatment in §21; §§12, 15 carry the event-angle facts (July 21 breach disclosure, GLM 5.2 forensic role); all other sections carry one-line signatory pointers at most.
8. **Kling pricing** — contradicted across sources; dropped entirely (not carried in §§18 or 22).
9. **"OpenAI/Google later signed the July 24 letter"** — [SECONDARY] community-only; not repeated as fact anywhere.
10. **Vals 52.37% (+9.48) for DeepSeek V4-Pro-0813** — benchmark unspecified; carried with the caveat, never used as a board score.
11. **TB 4.0 vs TB 2.1 / AA percentage vs 53-point scales** — never mixed; every score carries its version or snapshot date.
12. **Anthropic $30B Azure vs $30B Series G** — never merged (see A.1 #23).
13. **Main-KB overlap** — anything marked "covered in main KB" in the sources was reduced to a one-line cross-reference, not rewritten.
14. **01.AI exit** — carried in §17 as [SECONDARY] with the transparency note that the wave-6 Part-1 source attaches no dedicated URL to it.
15. **Step 5 Preview / openPangu 2.0 / Ling 3.0 Flash licenses** — left [UNVERIFIED] or unlabeled rather than inferred as permissive.



## A.4 Expansion-round contradictions (2026-09-22)

The 23-section expansion (8,129 new lines, ~1,500 new source URLs) surfaced the following additional contradictions. None were resolved by fiat; each is recorded with its competing claims and provenance.

| # | Claim(s) in conflict | Status / handling | Provenance |
|---|---|---|---|
| 45 | GPT-5.6 launch: 2026-07-02 vs OpenAI launch page 2026-07-10 (Copilot 2026-07-09) | Recorded as dated variants; not merged. | [SECONDARY] |
| 46 | GPT-5.6 Sol pricing: $4/$20 vs $5/$30 | Competing rows kept; basis of the $5/$30 figure unknown. | [VENDOR] |
| 47 | GPT-5.4 context window: 272K vs 1.05M | Competing claims preserved. | [VENDOR] |
| 48 | Astra ARC-AGI-3: 99.9% (vendor harness) vs 62.7% (standard harness) | Harness-bound; never compared. | [VENDOR] |
| 49 | Opus 4.6 surcharge: $7.50 vs $10/M | Competing figures preserved. | [VENDOR] |
| 50 | Sonnet 4.6 max output: 16K vs 64K | Competing figures preserved. | [SECONDARY] |
| 51 | Fable 5 context: 128K vs 1M | Competing claims preserved. | [VENDOR] |
| 52 | Fable cache-read pricing: $1 → $0.25 on 2026-09-01 | Resolved as a dated price change, not a contradiction. | [VENDOR] |
| 53 | Gemini 3.1 Pro pricing: $2/$12 vs $1.50/$9 | Competing rows preserved. | [VENDOR] |
| 54 | Nano Banana 2 context: 1M vs 131K | Competing claims preserved. | [SECONDARY] |
| 55 | Veo 3.1 Fast: $0.10 vs $0.15/s | Competing figures preserved. | [SECONDARY] |
| 56 | AI Ultra pricing: $249.99 vs $124.99 vs ~$100 | Regional/tier variants; kept as competing rows. | [VENDOR/SECONDARY] |
| 57 | GLM Coding Plan: $18/$72/$160 vs $18–$168 vs promo $3–10/$30/~$80 | Competing rows preserved. | [VENDOR] |
| 58 | "MTP" expansion: Multi-Token Prediction vs Multimodal-Task-Prompt | Source-dependent; both documented. | [SECONDARY] |
| 59 | Kimi K3 license: bespoke Kimi K3 License vs "Modified MIT" | Resolved: 4 sources confirm the bespoke K3 License; "Modified MIT" describes identical terms. K2 line was Modified MIT; K3 moved to bespoke. | [VENDOR] |
| 60 | Moonshot valuation: $10B / $18B / $20B+ / $30B | Different dates and funding stages; not comparable. | [SECONDARY] |
| 61 | MiniMax M2.5 context: 192K vs 196K | Competing figures preserved. | [VENDOR] |
| 62 | M3 MSA speedups: 9.7×/15.6× vs 9×/15× vs 14.2×/7.6× | Different test setups; never compared. | [VENDOR] |
| 63 | MiMo V2-Omni context: 256K vs 262K | Catalog discrepancy preserved. | [SECONDARY] |
| 64 | Mistral revenue: €200M (Le Monde) vs ~$400M ARR | Competing figures preserved. | [SECONDARY] |
| 65 | SpaceX/xAI acquisition date: Feb 2026 vs "newly announced" Sept 2026 | Competing datings preserved. | [SECONDARY] |
| 66 | Mistral Medium 3.1: 2025-08-13 (catalog) vs 08-12 (docs); Medium 3.5: 2026-04-28 vs 04-29/30 | Source-dependent datings preserved. | [VENDOR] |
| 67 | Grok 4.6 AA Index: 54 (v4.1) vs 61 (v4.1.1 per AWS) | Methodology versions; never combined. | [SECONDARY] |
| 68 | DeepSeek V4.1 Flash Terminal-Bench 2.1: 90.6% (codingfleet) vs 83.8% (official board) | Competing boards preserved. | [COMMUNITY] |
| 69 | Terminal-Bench 4.0: Fable 5.1 57.9% (official) vs Mythos 5.1 60.9% (Anthropic) | Source-bound; never compared. | [VENDOR] |
| 70 | Leading SWE-bench Pro: 61.5% (Scale) vs 81.2% (CodingFleet) vs 80.0% (llm-stats) | Three different boards; never mixed. | [SECONDARY] |
| 71 | Step 3.7 Flash: May 28 vs 29; pricing $0.16/$0.92 vs $0.20/$1.15 | Competing datings and rows preserved. | [VENDOR] |
| 72 | Tencent Hy4: 770B vs ~780B | Competing figures preserved. | [VENDOR] |
| 73 | MiniMax IPO proceeds: HK$4.8B vs HK$5.54B | Competing figures preserved. | [SECONDARY] |
| 74 | ERNIE pricing: $0.85/$3.40 vs $0.59/$2.65 vs $3/$12 | Competing rows preserved. | [VENDOR] |
| 75 | Veo 3.1: 8s vs 12s duration; free tier contested | Competing claims preserved. | [VENDOR/SECONDARY] |
| 76 | Kling 3.0: Feb 4 vs 5; 4K native at launch vs Apr 23 | Competing datings preserved. | [SECONDARY] |
| 77 | Seedance resolution: 480p/720p vs 1080p vs 4K marketing | Competing claims preserved. | [VENDOR] |
| 78 | Midjourney V8.1 / V8.2 release dates | Competing datings preserved. | [SECONDARY] |
| 79 | FLUX.2 Flex: $0.05 vs $0.06/$0.12 | Competing rows preserved. | [VENDOR] |
| 80 | Suno v5.5 vs v6 / v6-wild | Competing version claims preserved. | [SECONDARY] |
| 81 | ElevenLabs v4: preview-only vs claimed GA | GA unverified as of 2026-09-20; carried as [UNVERIFIED]. | [UNVERIFIED] |
| 82 | Command A+: 24B vs 25B active parameters | Competing figures preserved. | [SECONDARY] |
| 83 | Cohere Rerank 4: 4K vs 32K context | Competing claims preserved. | [SECONDARY] |
| 84 | Command A: 128K vs 256K (mirror surfaces) | Competing claims preserved. | [SECONDARY] |
| 85 | Nemotron 3 Super: 12B vs 12.7B active | Competing figures preserved. | [VENDOR] |
| 86 | Cosmos 3: May 31 vs Jun 1, 2026 | Competing datings preserved. | [VENDOR] |
| 87 | Alpamayo 2 Super: 32B vs 34B | Competing figures preserved. | [SECONDARY] |
| 88 | Nemotron 3 Lightning: 30B/3B vs 31.6B/~3.6B | Competing figures preserved. | [VENDOR] |

## A.5 Expansion merge log (2026-09-22)

- 23 expansion files (8,129 lines total) merged as strict delta-only content into §§1–23; no base-section text was rewritten.
- `# EXPANSION §N` title lines were dropped; each `## NEW X` marker was remapped to a `### New verified … — expansion` block inside the corresponding base subsection (Key dated facts, Figures and metrics, Timeline and context, Implications, Sources and URLs); "continued" markers became `### … (continued — topic)` blocks in the same target subsection.
- Unknown `##` markers (none of the five) were demoted to `###` rather than promoted to section level.
- Provenance honesty: 248 explicit single-source markers from the MiMo/Meta/Mistral/xAI research round were preserved verbatim; single-source claims remain labeled [UNVERIFIED] or (single X coverage) rather than silently elevated.
- Two-independent-source compliance is strong on newly researched material but was not audited claim-by-claim on every line; Annex A.4 records every contradiction the workers found.



