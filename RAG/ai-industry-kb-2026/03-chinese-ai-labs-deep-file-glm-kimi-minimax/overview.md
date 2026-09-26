---
id: ai-industry-kb-2026/03-chinese-ai-labs-deep-file-glm-kimi-minimax/overview
title: "3. Chinese AI Labs — Deep File: GLM, Kimi, MiniMax"
domain: chinese-ai-labs-deep-file-glm-kimi-minimax
role: deep-dive
task: actor-profile
actors: ["Anthropic", "ByteDance", "China", "Hugging Face", "MiniMax", "Moonshot", "OpenRouter", "United States", "Z.ai"]
dates: ["2026-04-13", "2026-04-20", "2026-06-13", "2026-06-16", "2026-06-17", "2026-06-24", "2026-07", "2026-07-31", "2026-08", "2026-08-10", "2026-08-14", "2026-08-18", "2026-08-20", "2026-08-26", "2026-08-28", "2026-09-22"]
keywords: ["glm", "kimi", "agent", "agents", "attention", "benchmark", "benchmarks", "claude", "compute", "cost", "cyber", "fable 5"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [997, 1038]
section: "3. Chinese AI Labs — Deep File: GLM, Kimi, MiniMax"
sha256: dbbca084bdba4064ac5a25c5b7cd365e4aeb6d3daab43ddac0249b7b04714ad5
---

# 3. Chinese AI Labs — Deep File: GLM, Kimi, MiniMax
Keywords: GLM-5.2, GLM-5.3, Z.ai, Zhipu AI, IndexShare, parameter-count reconciliation, MTP block, FP8 shorthand, Modified MIT license, Kimi K2.6, Moonshot AI, agent swarm, Kimi K3, Stable LatentMoE, Quantile Balancing, MiniMax H3, Hailuo 3.0, Hailuo 03, MiniMax 2.7T rumor, ByteDance Seed 2.1 Turbo, Volcano Engine FORCE, Terminal-Bench 3.0, DeepSWE v1.1, CyberGym, ExploitBench, OpenVuln, VulnHunter, GLM Coding Plan, ZCode, AI tigers

## Summary

This deep file consolidates the verification record for the Chinese-lab models in scope for the final knowledge base: **GLM-5.2 and GLM-5.3** (Z.ai), **Kimi K2.6 and K3** (Moonshot AI), and **MiniMax H3 / Hailuo 3.0** plus the MiniMax business context and the MiniMax 2.7-trillion-parameter LLM rumor. ByteDance's Seed 2.1 Turbo is covered here as the Chinese-lab-adjacent frontier item because its release record lives in the same verification pass; Seedance 2.5 and the Seed 2.1 base/Pro tiers remain canonical elsewhere (wave3/06 §1.3). The record below prefers the primary technical reports (Moonshot's Kimi K3 tech report) and config-derived evidence over press re-reporting, and every vendor-sourced figure carries its label.

Six brief claims were adjudicated in the verification pass (wave3.1/03), with the following consolidated verdicts:

1. **GLM-5.2 release — VERIFIED**, with the parameter count settled as a three-way reconciliation, not a single number: **753B total / ~40B active** is canonical; Z.ai's own "744B-A40B" is an FP8-build/VRAM shorthand, not the architectural total. MIT open weights, native 1M-token context, IndexShare sparse attention. The claim "best open-weight text model" is valid **only at the release moment** [DIRECTIONAL].
2. **GLM-5.3 — VERIFIED in full detail**: announced 2026-08-14, same 743B MoE base as GLM-5.2, all gains from scaled post-training (more RL environments, task diversity, longer trajectories, more RL compute); Terminal-Bench 3.0 **28.3**, DeepSWE v1.1 **66.9**, CyberGym **84.5**, ExploitBench **54.4** [VENDOR]; launch access via GLM Coding Plan ($18/mo) and ZCode; weights held ~2 weeks for safety hardening (the cyber capability grew faster than Z.ai expected) and shipped 2026-08-28/29 under a bespoke GLM-5.3 License. Always-on reasoning (low/high/max) is a breaking change from 5.2. Text-only, like 5.2.
3. **Kimi K2.6 — VERIFIED with a release-phase correction**: the brief's 2026-04-13 date marks the start of an eight-day preview; general availability was **2026-04-20/21**. Specs confirmed: 1T total / 32B active, 384 experts (8 routed + 1 shared per token), Modified MIT, 256K/262,144 context. The brief's "near Opus 4.6" **understates** the cited benchmark: SWE-Bench Pro 58.6% vs Claude Opus 4.6's 53.4% — Kimi K2.6 was **ahead**, not close.
4. **Kimi K3 — CROSS-REFERENCE**: full chronology lives in wave3/02-openweight-chronology.md (Claim 4). This file keeps only the architecture detail from the Moonshot technical report (Stable LatentMoE, Quantile Balancing, MXFP4, SiTU-GLU) and the flagged source conflicts — per the dedup rule, nothing from the wave3 chronology is recopied.
5. **MiniMax — NAME ALIAS RESOLVED**: MiniMax H3, Hailuo 3.0, and "Hailuo 03" are **one model, three names** — the consolidation must merge them into a single entry, not three. Released 2026-07-31 (date ±1 day across sources). Model details are canonical in wave2.1/07-multimodal-agents-delta.md §1.1; this file records the alias resolution, the Community License territorial restrictions, and the [UNVERIFIED] MiniMax 2.7T LLM rumor (Reuters, early July 2026).
6. **Seed 2.1 Turbo — VERIFIED as a real model with a dating correction**: ByteDance's official June 23/24 FORCE launch materials announce Seed 2.1, Seed 2.1 Pro, and a preview — **no Turbo variant named**; the Turbo's presence at FORCE rests on a single June 24 trade report. Western trackers only logged it on 2026-08-10/12 (first-seen dates, not release dates). Pin **2026-06-24** as launch with a single-source caveat. Specs verified: 262K context, text/image/video in → text out, reasoning plus tool calling, closed/proprietary, $0.50/$2.50 per M tokens. **No independent benchmarks** [UNVERIFIED]; parameter count undisclosed.

Two standing watch items for future RAG updates: the MiniMax 2.7T LLM (no confirmation by 2026-09-22) and the confidential Moonshot HK IPO filing rumor (~$3B target, single secondary, carried from wave2.1/05) — both stay [UNVERIFIED].

### Evidence grades used in this file

- **[PRIMARY]**: the Moonshot Kimi K3 technical report; Reuters launch-day reporting; model cards and Hugging Face repository metadata.
- **[VENDOR]**: all Z.ai benchmark and pricing figures (GLM-5.2/5.3), MiniMax's 2K-cost claim, Moonshot's ~2.5× scaling-efficiency claim. BetaNews explicitly notes the GLM-5.3 figures were not independently verified in this pass.
- **[COMMUNITY]**: config-derived research notes (the GLM-5.2 parameter reconciliation), practitioner write-ups, architecture inspections of the K3 report, engineering journals (Seed 2.1 Turbo dating).
- **[UNVERIFIED]**: the MiniMax 2.7T LLM rumor, the Moonshot HK IPO filing rumor, Seed 2.1 Turbo benchmarks/parameters/license, pre-2026-08-14 GLM-5.3 mentions.
- **[DIRECTIONAL]**: time-bound superlatives ("best open-weight text model" at GLM-5.2's release moment; "beats Opus 4.8" benchmark-dependently).
- **What this file does not own**: Kimi K3 chronology and its "beats Opus 4.8" evidence (wave3/02 Claim 4); MiniMax H3 full model detail (wave2.1/07 §1.1); GLM-5.3-Flash (wave1/07 §2.4); Seedance 2.5 (wave3/06 §1.3); the OpenVuln narrative (wave1/07 §2.4). This file is the canonical home for the GLM-5.2 parameter reconciliation, the GLM-5.3 rumor-phase chronology, the Kimi K2.6 release record, the MiniMax alias resolution, and the Seed 2.1 Turbo dating record.

## Key dated facts

### GLM (Z.ai)

- **2026-06-13/16 — GLM-5.2 announcement window.** Z.ai announces GLM-5.2; the US export-control order suspending Anthropic's Fable 5 / Mythos 5 landed the day before GLM-5.2 shipped, and launch coverage positioned GLM-5.2 as the open-weight counter-move.
- **2026-06-16 — GLM-5.2 weights released** (pinned date; ±1-day source variance — the Hugging Face card is dated 2026-06-17). MIT open weights. This is the canonical release date for the KB; wave3/02's finding matches, no contradiction.
- **Early August 2026 — GLM-5.3 rumor phase.** No official confirmation from Z.ai; only analyst and press speculation that a GLM-5.x successor was imminent. Per the brief's dating note: anything dated in this window must be marked [UNVERIFIED] rumor, not fact.
- **2026-08-14 — GLM-5.3 announced** (a Friday). First confirmed sighting across Unite.AI, SiliconANGLE, Axios, BetaNews. From this date, claims are attributable to Z.ai's own release materials.
- **2026-08-18 — GLM-5.3 API pricing opens**: $1.40/$4.40 per M input/output tokens; cached input $0.26, matching GLM-5.2.
- **2026-08-20 — "Ox Alpha" tops OpenRouter usage charts anonymously**; retrospectively confirmed by Z.ai (to Bloomberg) as GLM-5.3-Flash — canonical home wave1/07-frontier-vs-open-weight.md §2.4, not duplicated here.
- **2026-08-26 — GLM-5.3-Flash announced**: 320B/18B MoE, MIT, $0.15/M (canonical home wave1/07 §2.4).
- **2026-08-28/29 — GLM-5.3 weights ship** (`zai-org/GLM-5.3`, 753B/40B, bespoke GLM-5.3 License), ending the ~2-week safety hold — the first weight hold in the GLM line. Resolves the wave1 open question as a delay, not a policy reversal.

### Kimi (Moonshot AI)

