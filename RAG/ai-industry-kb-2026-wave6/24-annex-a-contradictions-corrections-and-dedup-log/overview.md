---
id: ai-industry-kb-2026-wave6/24-annex-a-contradictions-corrections-and-dedup-log/overview
title: "Annex A — Contradictions, Corrections, and Dedup Log"
domain: appendix
role: appendix
task: reference
actors: ["Alibaba", "Anthropic", "ByteDance", "Cohere", "DeepSeek", "Falcon", "Google", "Groq", "Hugging Face", "LongCat", "Meituan", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Poolside", "Z.ai", "vLLM", "xAI"]
dates: ["2025-06", "2025-08", "2025-09-22", "2025-09-30", "2025-12", "2025-12-09", "2025-12-31", "2026-01", "2026-01-05", "2026-01-14", "2026-02-20", "2026-02-27", "2026-03-13", "2026-04", "2026-04-07", "2026-04-20", "2026-04-22", "2026-04-26", "2026-04-29", "2026-05-07", "2026-06-01", "2026-08-20", "2026-08-28", "2026-09", "2026-09-03", "2026-09-21", "2026-09-22", "2026-09-24", "2026-10-23", "2026-10-27"]
keywords: ["acquisition", "apache", "astra", "claude", "cohere", "compute", "consumer", "decode", "deepseek", "fable 5", "gemini", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [11547, 11616]
section: "Annex A — Contradictions, Corrections, and Dedup Log"
delta_of: ai-industry-kb-2026
sha256: cacfd4bf936c7e3b9ddf59059648b1187c114e29db8211240265e66434606483
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

