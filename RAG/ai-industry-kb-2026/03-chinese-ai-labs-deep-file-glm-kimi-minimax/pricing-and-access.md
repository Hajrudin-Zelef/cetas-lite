---
id: ai-industry-kb-2026/03-chinese-ai-labs-deep-file-glm-kimi-minimax/pricing-and-access
title: "Pricing and access"
domain: chinese-ai-labs-deep-file-glm-kimi-minimax
role: deep-dive
task: pricing
actors: ["Anthropic", "ByteDance", "China", "EU", "ExploitGym", "MiniMax", "Moonshot", "OpenAI", "OpenRouter", "SGLang", "United States", "Z.ai", "vLLM"]
dates: ["2026-04-20", "2026-06", "2026-06-16", "2026-06-23", "2026-06-24", "2026-07-27", "2026-08", "2026-08-02", "2026-08-18", "2026-08-28", "2026-09-22"]
keywords: ["pricing", "agent", "agentic", "benchmark", "benchmarks", "claude", "consumer", "cyber", "disclosure", "fine-tuning", "glm", "kimi"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1136, 1236]
section: "3. Chinese AI Labs — Deep File: GLM, Kimi, MiniMax"
sha256: 2b8f0a4897ff57419f6ed5f11ed94122dc2c66cc94552362c17030ae1891fed4
---

# Pricing and access

- Terminal-Bench 2.1: **81.0%** — first open model above 80% (vendor-adjacent secondary).
- SWE-Bench Pro: **62.1** vs 58.4 previous generation [VENDOR].
- Binding constraint: text-only (no vision).

#### GLM-5.3 (Z.ai-reported; BetaNews flags the figures as not independently verified)

- Terminal-Bench 3.0: **28.3** (up from 4.6 on GLM-5.2 — ~6×).
- DeepSWE v1.1: **66.9** (up from 46.2).
- FrontierSWE: **78.1**; AutomationBench v1.0.6: **48.2**.
- CyberGym: **84.5** (up from 77.2); ExploitBench: **54.4** (up from 24.4 — more than doubled); ExploitGym: **105/130** tasks under normalized 2h/6h budgets.

#### Kimi K2.6 (secondary, Moonshot-announcement-adjacent; not independently reproduced in this pass)

- SWE-Bench Pro: **58.6%** vs GPT-5.4 57.7% vs Claude Opus 4.6 **53.4%** — Kimi K2.6 is **ahead** of Opus 4.6; the brief's "close to Opus 4.6" undersells it. Correction recorded for consolidation.
- SWE-Bench Verified: **80.2%** (10-run average); SWE-Bench Multilingual: **76.7%**.
- Terminal-Bench 2.0: **66.7%**; LiveCodeBench v6: **89.6%**.
- HLE-Full with tools: **54.0%** vs GPT-5.4 52.1% vs Opus 4.6 53.0%.
- Reasoning: AIME 2026 **96.4%**; HMMT 2026 **92.7%**; GPQA-Diamond **90.5%**.
- Vision: MMMU-Pro **79.4%** (80.1% with Python tool use); MathVision **87.4%** (93.2% with Python).
- Agentic: OSWorld-Verified **73.1%**; BrowseComp **83.2%** single-agent / **86.3%** with agent swarm.

#### MiniMax H3

- #1 Video Editing with Audio on Artificial Analysis — first open model to top an AI video ranking.

### Pricing and access

- **GLM-5.3**: launch access through Z.ai's API and the **GLM Coding Plan** — **$18/mo** base tier (Pro $80, Max $168 reported) — plus the **ZCode** agent. API per-token pricing from 2026-08-18: **$1.40/$4.40** input/output per M tokens; cached input **$0.26**, matching GLM-5.2.
- **GLM Coding Plan tiers** (reported): $18/mo entry, $80/mo Pro, $168/mo Max — the subscription ladder through which Z.ai gated first access to GLM-5.3 before API per-token pricing opened on 2026-08-18. This is the 2026 Chinese-lab access pattern: subscription-first, API-second, weights-last (with the weight leg delayed here by the safety hold).
- **GLM-5.3 weights** (`zai-org/GLM-5.3`, 753B/40B) shipped 2026-08-28/29 under the **bespoke GLM-5.3 License** — note the license step-down from MIT on GLM-5.2 (and on GLM-5.3-Flash, which stayed MIT).
- **GLM-5.3-Flash**: 320B/18B MoE, MIT, $0.15/M (canonical wave1/07 §2.4; no duplication).
- **Kimi K2.6**: **$0.60/M input** via Moonshot API; deployable on vLLM, SGLang, KTransformers; HF `moonshotai/Kimi-K2.6`, Modified MIT.
- **Seed 2.1 Turbo**: **$0.50/$2.50** per M input/output tokens on OpenRouter/NanoGPT (verified current as of 2026-09-22).
- **MiniMax** [VENDOR, single source]: generating 2K video costs less than one-third of mainstream rival products.

### Vulnerability ledger — OpenVuln ("VulnHunter")

- Z.ai launched **OpenVuln** (internal name "VulnHunter"), a scanning service built on GLM-5.3 for public code repositories.
- The public disclosure ledger at **cvd.z.ai** credits GLM-5.3 with **2,436 vulnerability findings across 269 open-source projects** (1,097 critical/high; **53 assigned CVE numbers**; oldest flaw dated to **1981**; average bug hidden **26.6 years**).
- Already covered in wave1 §2.4 — this file only confirms the figures; do not duplicate.

### At-a-glance comparison (open-weight flagships, April–August 2026)

| Model | Release (weights) | Total / active | License | Context | Headline fact |
|---|---|---|---|---|---|
| Kimi K2.6 | 2026-04-20/21 (preview 04-13) | 1T / 32B | Modified MIT | 256K / 262,144 | SWE-Bench Pro 58.6% — ahead of Opus 4.6 (53.4%) |
| GLM-5.2 | 2026-06-16 (±1 day) | 753B / ~40B | MIT | 1M native | IndexShare; TB 2.1 81.0% — first open model above 80% |
| Kimi K3 | ~2026-07-27 (launch 07-16) | 2.8T / 104B | open weights | — | 896+2 experts, 16 active; Stable LatentMoE, Quantile Balancing |
| MiniMax H3 | 2026-08-02/03 (launch 07-31) | 33B dense | H3 Community License | multimodal video | #1 Video Editing with Audio on Artificial Analysis |
| GLM-5.3 | 2026-08-28/29 (announced 08-14) | 753B / ~40B (same base) | bespoke GLM-5.3 License | — | TB 3.0 28.3 (from 4.6) — post-training-only gains |

### GLM-5.2 → GLM-5.3 deltas (same 743B base — the post-training-scaling demonstration) [VENDOR]

| Benchmark | GLM-5.2 | GLM-5.3 | Change |
|---|---|---|---|
| Terminal-Bench 3.0 | 4.6 | 28.3 | ~6× |
| DeepSWE v1.1 | 46.2 | 66.9 | +20.7 pts |
| CyberGym | 77.2 | 84.5 | +7.3 pts |
| ExploitBench | 24.4 | 54.4 | more than doubled |
| ExploitGym | — | 105/130 (normalized 2h/6h budgets) | new metric |
| FrontierSWE | — | 78.1 | new metric |
| AutomationBench v1.0.6 | — | 48.2 | new metric |

(GLM-5.2's Terminal-Bench 2.1 81.0% is a separate harness-generation record — TB 2.1 and TB 3.0 scores are not comparable.)

### Licensing ladder (the 2026 Chinese-lab licensing pattern)

| Model | License | Notable term |
|---|---|---|
| GLM-5.2 | MIT | unrestricted open weights |
| GLM-5.3-Flash | MIT | 320B/18B MoE, $0.15/M |
| GLM-5.3 | bespoke GLM-5.3 License | step down from MIT after the cyber-driven weight hold |
| Kimi K2.6 | Modified MIT | commercial use, fine-tuning, resale, no royalties; operators above 100M MAU or $20M/mo revenue must display "Kimi K2.6" in the product UI |
| MiniMax H3 | MiniMax H3 Community License | **territorial exclusion**: US, EU, UK, South Korea barred from local deployment |

Pattern for the KB: the most capable 2026 artifacts trend toward bespoke, gated, or territorially restricted licenses (GLM-5.3, H3), while mid-tier or distilled variants stay MIT.

### Seed 2.1 family (June 2026)

| Variant | Announced | Context | Pricing | Provenance |
|---|---|---|---|---|
| Seed 2.1 (base) | 2026-06-23/24, FORCE | undisclosed | undisclosed | official ByteDance launch materials |
| Seed 2.1 Pro | 2026-06-23/24, FORCE | undisclosed | undisclosed | official; preview charted 8th on Code Arena (1539), "level with Opus 4.6" (chart-transcribed, caution) |
| Seed 2.1 Turbo | 2026-06-24, FORCE | 262K / 262,144 | $0.50/$2.50 per M | **single trade report**; official blog omits Turbo; Western trackers first-seen 08-10/12 |

All three: closed/proprietary, parameter counts undisclosed, text/image/video in → text out (Turbo verified; shared across the family per ByteDance materials). No independent benchmarks for any Turbo claim [UNVERIFIED].

### MiniMax naming merge (the alias table)

| Name | Where it appears | Status |
|---|---|---|
| MiniMax H3 | API model ID `MiniMax-H3`; MarkTechPost launch coverage; Reuters | canonical model identifier |
| Hailuo 3.0 | consumer Hailuo AI app; Minnesota Headlines FAQ ("also known as") | consumer-facing name |
| Hailuo 03 | note.com Japanese name table ("MiniMax H3 ＝ Hailuo 3.0 ＝ Hailuo 03") | stylized variant |

All three are one model. Any KB entry keyed on "Hailuo 3.0" or "Hailuo 03" alone must redirect to the merged MiniMax H3 entry. The "H" prefix marks the video/Hailuo line (vs "M" for the M2/M3 text LLMs).

## Main actors

### The labs

