---
id: ai-industry-kb-2026/04-chinese-ai-labs-deep-file-qwen-seed-ling/part-2
title: "4. Chinese AI Labs — Deep File: Qwen, Seed, Ling (part 2)"
domain: chinese-ai-labs-deep-file-qwen-seed-ling
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "Apple", "ByteDance", "China", "MiniMax", "Moonshot", "Nvidia", "OpenRouter", "Z.ai"]
dates: ["2026-02-16", "2026-03-30", "2026-04-02", "2026-04-16", "2026-04-21", "2026-05-19", "2026-05-20", "2026-05-21", "2026-06-02", "2026-06-24", "2026-07-15", "2026-07-16", "2026-07-19", "2026-07-21", "2026-07-23", "2026-07-27", "2026-07-31", "2026-08-03", "2026-08-10", "2026-08-11", "2026-08-12", "2026-08-14", "2026-08-17", "2026-08-26", "2026-09-02", "2026-09-03", "2026-09-04", "2026-09-05", "2026-09-22"]
keywords: ["qwen", "agent", "agentic", "apache", "benchmark", "benchmarks", "claude", "fable 5", "fp8", "glm", "gpu", "kimi"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1403, 1434]
section: "4. Chinese AI Labs — Deep File: Qwen, Seed, Ling"
sha256: 5ca5443cedc0f6f3db0595d02f75e1b7754638820ebfc7ab926dd601386c981f
---

# 4. Chinese AI Labs — Deep File: Qwen, Seed, Ling (part 2)

- 2026-02-16 — Qwen3.5-397B-A17B (Apache 2.0 open weights, 397B/17B MoE, 262K ctx via YaRN→1M) + Qwen3.5-Plus (proprietary hosted, 1M ctx, tools). Canonical chronology: §2.
- 2026-03-30/31 — Qwen3.5-Omni (hosted API-only; Plus/Flash/Light; Thinker–Talker architecture; 113-language speech recognition). Date corrected from "April." Canonical chronology: §2.
- 2026-04-02 — Qwen3.6-Plus (proprietary API, agentic coding, 1M context). Canonical chronology: §2.
- 2026-04-16 — Qwen3.6-35B-A3B (open weights, 35B/3B MoE). Canonical chronology: §2.
- 2026-04-21 — Qwen3.6-Max-Preview (proprietary preview). Canonical chronology: §2.
- 2026-05-20 — Qwen3.7-Max unveiled (Hangzhou; API public 2026-05-21/23; proprietary, 1M ctx, 65K output). Canonical chronology: §2.
- 2026-05-19/2026-06-02 — Qwen3.7-Plus preview (2026-05-19) → official release (2026-06-02); proprietary agent-oriented multimodal model. Canonical chronology: §2.
- 2026-06-24 — ByteDance Seed 2.1 Turbo at Volcano Engine FORCE; 262K ctx; $0.50/$2.50 per M. Defensible date with single-source caveat on Turbo's presence at FORCE; Western catalog first-seen dates (2026-08-10/12) are aggregator artifacts, not releases.
- 2026-07-19 — Qwen3.8 announced (2.4T MoE, WAIC Shanghai); Qwen3.8-Max-Preview immediate via Token Plan at 10% pricing; vendor "just behind Claude Fable 5" [VENDOR].
- 2026-07-21 — Qwen Image 3.0 (closed, API-only; no weights/report/benchmarks at launch). Canonical chronology: §2.
- 2026-07-23 — InclusionAI Ling 3.0 Flash (124B/5.1B active MoE, 262K ctx; free tier on OpenRouter).
- 2026-07-27 — Qwen3.7 Flash public (per OpenRouter; QwenCloud snapshot 2026-07-15 is an identifier, not the launch date).
- 2026-08-03 — Qwen3.8-Max GA (2.4T/95B active [third-party-reported], 1M ctx, multimodal, $2/$6 per M; custom restrictive license).
- 2026-08-12 — Qwen3.8-Max downloadable weights (`Qwen/Qwen3.8-2.4T-A95B` + FP8; custom qwen3.8-max license; text-only, thinking-mode-only — strictly separate from the hosted multimodal/1M product).
- 2026-08-14 — Qwen3.8-27B weights (Apache 2.0; 27B dense; multimodal). Date corrected from "August 12."
- 2026-08-26 — Qwen3.8-Flash (open weights, 6B active) + Qwen3.8-Flash-Next (hosted only).
- 2026-09-02 — Qwen3.8-Max-0902 (new checkpoint; cross-ref §2 / wave 2.1).
- 2026-09-03 — InclusionAI Ling 3.0 Flash Fin (finance variant; $0.06/$0.18 per M).
- Date unpinned — Qwen3.6-27B dense variant exists in third-party benchmark tables (Terminal-Bench 2.1 60.7, AA Index 38) but no release date was pinned in research; flagged [DATE UNVERIFIED] in consolidation.
- 2026-07-15 — Beijing approved Apple Intelligence powered by Qwen in China (secondary, memeburn.com) — actor-context fact, not a model release.
- 2026-07-15 — QwenCloud snapshot identifier `qwen3.7-flash-2026-07-15` (snapshot, not public launch; public availability 2026-07-27 per OpenRouter).
- 2026-07-16/17 — Moonshot Kimi K3 launch window (2.8T), the immediate predecessor to the Qwen3.8 announcement; cross-ref §5 (Kimi line).
- 2026-07-23 — EqualOcean coverage of Qwen3.8 preview; Alibaba holds ~36% of Moonshot (MLQ) — the ownership link behind the July sequencing.
- 2026-07-31 — MiniMax H3 video model launch (Reuters) — the third July-cluster Chinese-lab release; cross-ref §7/§12, not duplicated here.
- 2026-08-11 — 302.AI head-to-head: Qwen Image 3.0 Pro vs GPT-Image-2 in infographic/layout scenarios [COMMUNITY].
- 2026-08-12 — NVIDIA deployment blog: "Alibaba released the open weights" + GB300 NVL72 reference deployment (>4,000 tok/s per GPU, FP8).
- 2026-08-10/12 — Western aggregator first-seen dates for Seed 2.1 Turbo (LLM Gateway PR #3580, NanoGPT, OpenRouter) — aggregator artifacts, not the release.
- 2026-08-17 — gigazine.net English coverage of the Qwen3.8-27B weight drop (dated three days after the 2026-08-14 official announcement).
- 2026-09-04 — AA Intelligence Index v4.3 composite: Qwen3.8-Max ≈ 40 (4th open) vs GLM-5.3 ≈ 45, Kimi K3 ≈ 44 [DIRECTIONAL].
- 2026-09-05 — Code Arena WebDev table (preliminary): Qwen3.8-Max-0902 4th — covered wave 2.1/05; recap only.
- 2026-09-22 — Seed 2.1 Turbo pricing $0.50/$2.50 per M verified current on OpenRouter/NanoGPT (as-of date for all pricing in this part unless otherwise sourced).

