---
id: ai-industry-kb-2026-wave6/03-qwen-and-alibaba/overview
title: "§3. Qwen and Alibaba"
domain: qwen-and-alibaba
role: deep-dive
task: actor-profile
actors: ["Alibaba", "China", "Meta", "SGLang"]
dates: ["2026-02-03", "2026-02-16", "2026-02-24", "2026-03-30", "2026-04", "2026-04-02", "2026-04-16", "2026-04-20", "2026-04-22", "2026-05-20", "2026-06-01", "2026-07", "2026-08", "2026-08-12", "2026-08-26", "2026-09-02", "2026-09-17"]
keywords: ["qwen", "apache", "attention", "benchmark", "leaderboard", "license", "moe", "omni", "open source", "pricing", "revenue", "sglang"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1086, 1139]
section: "§3. Qwen and Alibaba"
delta_of: ai-industry-kb-2026
sha256: 1ae26d55c03f68842fef7b06dab04e9f2e6b3fc9738b247c6c131f11b1898d9a
---

# §3. Qwen and Alibaba

Keywords: qwen3.5 397b, qwen3.5 flash, qwen3.5 omni, qwen3.6 plus, qwen3.6 max preview, qwen3.6 35b, qwen3.6 27b, qwen3.7 max, qwen3.7 plus, qwen3.7 flash, qwen3.8 flash next, qwen3.8 max 0902, qwen3 coder next, gated deltanet, 800k coding tasks, qwen4 exp, qwen 4 in training, alibaba qwen, apache 2.0 qwen, custom license qwen3.8 max, 100m mau, qwen downloads

## Summary

Alibaba's Qwen line in 2026 moved fast through the 3.5, 3.6, 3.7, and 3.8 generations, mixing open Apache-2.0 releases with closed flagship tiers. The verified chronology: **Qwen3.5-397B-A17B (2026-02-16)**; medium open models plus hosted Flash (2026-02-24/25); **Qwen3.5-Omni (2026-03-30)**; the Qwen3.6 line across April (Plus closed 2026-04-02, 35B-A3B open 2026-04-16, Max-Preview closed 2026-04-20, 27B open 2026-04-22); the Qwen3.7 line (Max closed 2026-05-20, Plus closed 2026-06-01, Flash July snapshot/release); **Qwen3.8-Flash-Next 125B/6B (2026-08-26)**; and the **Qwen3.8-Max-0902 snapshot (2026-09-02)** — a checkpoint refresh, not a new base model. [SECONDARY]

Standalone highlights: **Qwen3-Coder-Next (2026-02-03/04)** — 80B total, ~3B active, Apache 2.0, Gated DeltaNet hybrid attention, 262K context, trained on 800K+ verifiable coding tasks [VENDOR]; the Qwen3.8-Max custom license (name-display above 100M MAU/$20M revenue; separate commercial license for MaaS above $50M) [VENDOR]; and `qwen4_exp`, an experimental SGLang checkpoint that is **not** a Qwen 4 announcement — Alibaba stated Qwen 4 was still in training. [VENDOR]

## Key dated facts

### Qwen3.5 line
- **2026-02-16** — Qwen3.5-397B-A17B released (open, 397B total / 17B active MoE). [SECONDARY]
- **2026-02-24/25** — Medium open Qwen3.5 models released alongside hosted **Qwen3.5-Flash**. [SECONDARY]
- **2026-03-30** — Qwen3.5-Omni released: text/audio/video omni-capable line, multilingual and captioning upgrades. [SECONDARY]

### Qwen3.6 line (April 2026)
- **2026-04-02** — Qwen3.6-Plus released (closed/API tier). [SECONDARY]
- **2026-04-16** — Qwen3.6-35B-A3B released (open, 35B total / ~3B active). [SECONDARY]
- **2026-04-20** — Qwen3.6-Max-Preview released (closed). [SECONDARY]
- **2026-04-22** — Qwen3.6-27B released (open). [SECONDARY]

### Qwen3.7 line
- **2026-05-20** — Qwen3.7-Max released (closed). [SECONDARY]
- **2026-06-01** — Qwen3.7-Plus released (closed). [SECONDARY]
- **July 2026** — Qwen3.7-Flash: July snapshot/release cadence. [SECONDARY]
- The vendor plan that Qwen3.7-Plus "will be open source" **remained unfulfilled in the observation window** — a stated intent, not a release. [VENDOR]

### Qwen3.8 line
- **2026-08-26** — Qwen3.8-Flash-Next released (125B total / 6B active). [SECONDARY]
- **2026-09-02** — Qwen3.8-Max-0902: a dated **snapshot/checkpoint of Qwen3.8-Max**, not a new base model. Reported at AA Intelligence Index v4.3 = 45 [SECONDARY, r/LocalLLaMA via alextech, 2026-09-17], reclaiming the China lead on that board — checkpoint-dated, not methodology-independent. [SECONDARY]
- Qwen3.8-Max API pricing reported at $2.00/$6.00 per M input/output (2026-08-12); pricing discussion is intentionally shallow here — the section's license terms are the durable fact. [SECONDARY]

### Qwen3-Coder-Next (2026-02-03/04)
- **2026-02-03/04** — Qwen3-Coder-Next released: 80B total / ~3B active, **Apache 2.0**, Gated DeltaNet hybrid attention, 262K context, trained on **800K+ verifiable coding tasks**. [VENDOR]
- The Gated DeltaNet hybrid architecture and the 800K-task coding corpus are the differentiators against the general Qwen3.6 line. [VENDOR]

### Licensing and the Max tier
- Qwen3 / 3.5 / 3.6 / 3.8 (up to 27B-class) are **Apache 2.0** [VENDOR]; Qwen3.8-Max ships under a **custom license**: name-display required above 100M MAU or $20M monthly revenue; MaaS/AI-assistant businesses above $50M trailing revenue need a separate commercial license; internal use is exempt. [VENDOR]
- The custom Max license is the "flagship has the most strings" case inside Alibaba's own line (see wave6/05 Part 5-A1 for the lab-by-lab map). [DIRECTIONAL]

### qwen4_exp (August 2026)
- `qwen4_exp` appeared as an **experimental checkpoint (SGLang day-0)** — community-tracked, not vendor-announced. [SECONDARY]
- Alibaba's stated position: **Qwen 4 was still in training**; `qwen4_exp` is not a Qwen 4 release and must not be cited as one. [VENDOR]

### Qwen3.8-Max benchmark and adoption (secondary context)
- Qwen3.8-Max DeepSWE: **56.6 [VENDOR] vs 69.3 [SECONDARY]** — different harnesses; provenance-tagged, not averaged. The corpus's standing rule: never merge vendor-harness and third-party-harness numbers. [VENDOR/SECONDARY]
- Adoption: **Reuters, Airbnb, and Pinterest** have adopted Qwen models [SECONDARY]; a community "Global LLM Download Leaderboard" (HF, Sept 2026) shows **Qwen3-0.6B at 22.7M cumulative downloads** — small models dominate raw counts. [COMMUNITY]
- Qwen3.7-Plus API pricing reported at **$0.276/M input** (Sept 2026, Global ≤256K non-thinking) — dated, and cheap enough to sit in the bottom quartile of the corpus price table. [SECONDARY]


### New verified facts — expansion

