---
id: ai-industry-kb-2026-wave6/17-other-chinese-labs/figures-and-metrics
title: "Figures and metrics"
domain: other-chinese-labs
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Falcon", "Moonshot", "OpenRouter", "StepFun"]
dates: ["2024-10", "2026-01-05", "2026-01-22", "2026-05", "2026-05-08", "2026-06", "2026-06-12", "2026-06-23", "2026-06-24", "2026-06-30", "2026-07", "2026-07-21", "2026-07-23", "2026-07-31", "2026-08-28", "2026-09-20", "2026-09-27", "2026-10-15"]
keywords: ["apache", "claude", "cost", "deepseek", "fp8", "ipo", "kimi", "license", "memory", "moe", "open-weight", "opus 4"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8436, 8520]
section: "§17. Other Chinese Labs"
delta_of: ai-industry-kb-2026
sha256: b0e2b0a6fb519e751fd8ef8399a2af29a3e809f3c827add4f2b4841c75b95ae9
---

# Figures and metrics

## Figures and metrics
| Item | Size / detail | Date | Note |
|---|---|---|---|
| StepFun Step 5 Preview | API $1.00/$2.70 (cache $0.05) | 2026-09-20 | API-only; weights promised 2026-10-15 [SECONDARY] |
| ERNIE 5.0 | GA; $0.85/$3.40 | 2026-01-22 | Pricing conflicts with 5.1 [SECONDARY] |
| ERNIE 5.1 | $0.59/$2.65 | ~2026-05-08 | Version- vs market-driven unclear [SECONDARY] |
| Seed 2.1 Turbo | — | 2026-06-24 | [SECONDARY] |
| Seedance 2.5 | announced → global | 2026-06-23 → 2026-07-31 | Video line [SECONDARY] |
| Ling 3.0 Flash | 124B / ~5.1B | 2026-07-23 | License [UNVERIFIED] [SECONDARY] |
| Falcon H1R / H1 Arabic | — | 2026-01-05 | [SECONDARY] |
| Falcon Perception + OCR | — | ~2026-04 | Date approximate [SECONDARY] |
| Falcon Perception-300M | 300M | 2026-07 | [SECONDARY] |
| Laguna S 2.1 | 118B / 8B; OpenMDW-1.1; $0.09/$0.18 | 2026-07-21 | Cheapest open-weight API [VENDOR] |
| Hy4 Preview | 770B / 49B; Apache 2.0 | 2026-08-28 | Preview-gated [SECONDARY] |
| openPangu 2.0 Pro | 505B / 18B; 512K | ann. 2026-06-12 | License unverified [SECONDARY] |
| openPangu 2.0 Flash | 92B / 6B; 512K | staged from 2026-06-30 | License unverified [SECONDARY] |
| Sonar endpoints | retirement | 2026-09-27 | Chat-completions [SECONDARY] |

- DeepSeek R2: not released — one line only; §2 owns it.


### New verified metrics — expansion

| Claim | Value | Sources | Label |
|---|---|---|---|
| 01.AI audited 2025 revenue | 250M yuan (~$34M) | S1 (single source) | [SECONDARY] |
| 01.AI contract orders (May 2026) | >1.5B yuan (~$207M) | S1 (single source) | [SECONDARY] |
| 01.AI FY target | 2B yuan | S1 (single source) | [SECONDARY] |
| 01.AI recurring subscription share | ~half of 2026 contract volume | S1 (single source) | [SECONDARY] |
| 01.AI non-China business share | ~50% | S3 (single source) | [SECONDARY] |
| 01.AI IPO target | Hong Kong, 2027 (after Dec fiscal year-end) | S2, S3 | [SECONDARY] |
| 01.AI founding → unicorn | 6 vs 8 months (contradiction) | S8, S1 | [SECONDARY] |
| Baichuan-M3 total params | 235B (Qwen3-235B-A22B base) | S9, S10 | [SECONDARY] |
| Baichuan-M3 HealthBench-Hard | 44.4 | S9, S10 | [VENDOR] |
| Baichuan-M3 HealthBench Total | 65.1 | S9 (single source) | [VENDOR] |
| Baichuan-M3 hallucination rate | 3.5% (claimed industry-lowest) | S9, S10 | [VENDOR] |
| Baichuan-M3 SCAN-bench inquiry lead | +12.4 pts over 2nd place | S9, S12 | [VENDOR] |
| Baichuan-M3 speculative decoding | Gated Eagle3, claimed 96% speedup | S11 (single source) | [VENDOR] |
| Baichuan-M3 W4 memory | 26% of baseline; 2× RTX 4090 deployable | S11, S9 | [VENDOR] |
| Baichuan-M3 knowledge cutoff | October 2024 (inherited from Qwen3) | S13 (single source) | [VENDOR] |
| dots3-note total / active | 280B / 16B | S17, S15 | [SECONDARY] |
| dots3-note context | 512K tokens | S17, S19 | [SECONDARY] |
| dots3-note announcement | Aug 14, 2026, 02:00:31 UTC | S15 (single source) | [SECONDARY] |
| dots3-note license | Apache 2.0 | S15 (single source) | [SECONDARY] |
| dots3-note MoE structure | 256 routed + 1 shared, top-8; 1 dense + 45 MoE layers | S17 (single source) | [VENDOR] |
| dots3-note hidden size | 5120 (FFN 13824 dense / 1536 per expert) | S17 (single source) | [VENDOR] |
| dots3-note vision encoder | MoE ViT 7B total / 1.2B active | S17 (single source) | [VENDOR] |
| dots3-note audio encoder | 800M dense | S17 (single source) | [VENDOR] |
| dots3-note vocab | 152K | S17 (single source) | [VENDOR] |
| dots3-note Terminal-Bench 2.1 (SemiAnalysis via source) | 75.1 (+4.9 vs U.S. open models) | S16 (single source) | [DIRECTIONAL] |
| dots3-note OpenRouter price | $0 (free, rate-limited) | S19 (single source) | [SECONDARY] |
| Doubao-Seed-2.1 Pro input price | 6 yuan/M tokens (~$0.88) | S22, S23 | [SECONDARY] |
| Doubao-Seed-2.1 Pro output price | 30 yuan/M tokens (~$4.41) | S22, S23 | [SECONDARY] |
| Doubao-Seed-2.1 Pro cache-hit price | 1.2 yuan/M tokens (~$0.18) | S22, S23 | [SECONDARY] |
| Doubao-Seed-2.1 Turbo price | half of Pro | S22, S23 | [SECONDARY] |
| Doubao-Seed-2.1 Pro claimed TCO vs Claude Opus 4.6 | ~80% lower | S22 (single source) | [VENDOR] |
| Doubao daily token calls (June 2026) | >180 trillion | S23 (single source) | [SECONDARY] |
| Doubao YoY call growth | >10x | S23 (single source) | [SECONDARY] |
| Volcano Engine China public cloud MaaS share | 49.5% (IDC) | S23 (single source) | [SECONDARY] |
| Doubao-Seed-Code price | RMB 9.9 first month (~63% below market avg) | S25 (single source) | [SECONDARY] |
| Doubao-Seed-Code SWE-Bench Verified | 78.8% | S25 (single source) | [SECONDARY] |
| Doubao-Seed-Code context per query | 256,000 words | S25 (single source) | [SECONDARY] |
| Doubao 2.1 Pro 0915: Luanti fix rate | 83% of 1,000 issues (~387K-line repo, ~36h) | S24 (single source) | [VENDOR] |
| Doubao 2.1 Pro 0915: image/video token reduction | >30% vs prior generation | S24 (single source) | [VENDOR] |
| Step 3.7 Flash total / active | 198B / ~11B | S27, S28 | [SECONDARY] |
| Step 3.7 Flash backbone / vision encoder | 196B LM + 1.8B vision | S27, S28 | [SECONDARY] |
| Step 3.7 Flash context | 256K | S27, S28 | [SECONDARY] |
| Step 3.7 Flash throughput | up to 400 tok/s | S27 (single source) | [VENDOR] |
| Step 3.7 Flash pricing | $0.20 / $0.04 / $1.15 per M (miss/hit/out) | S28, S76 | [SECONDARY][VENDOR] |
| Step 3.7 Flash local minimum | 120 GB unified memory/VRAM | S28, S76 | [SECONDARY][VENDOR] |
| Step 3.7 Flash ClawEval-1.1 | 67.1 (vs 59.8 2nd) | S27, S30 | [VENDOR][COMMUNITY] |
| Step 3.7 Flash SWE-Bench PRO | 56.3 (2nd place) | S27, S30 | [VENDOR][COMMUNITY] |
| Step 3.7 Flash Advisor SWE-Bench Verified | 76.3–76.5% at $0.19/task | S28, S27 | [VENDOR] |
| Step 3.7 Flash real per-task cost (independent test) | ~$2.66/1,000 vs Opus 4.6 $4.05 | S31 (single source) | [SECONDARY] |
| Step 3.7 Flash independent coding score | 8/9 tasks | S31 (single source) | [SECONDARY] |
| Hunyuan3D 3.0 geometric resolution | 1536³ (3.6B voxels, 3x accuracy claim) | S32, S33 | [SECONDARY] |
| Hunyuan3D 3.0 announcement | Sept 16, 2025 (not 2026) | S32, S33 | [SECONDARY] |
| Hunyuan3D 3.1 | ~Feb 2026, closed hosted | S35 (single source) | [COMMUNITY] |
| HY3D-Bench | 252K+ watertight meshes (~11 TB), 1,252 categories | S35 (single source) | [COMMUNITY] |
| Hunyuan3D Engine free tier | 20 generations/day individuals; 200 API credits | S33, S35 | [SECONDARY] |
| Hunyuan3D-World 1.0-Lite VRAM | <17 GB (FP8, 35% memory cut) | S36 (single source) | [SECONDARY] |
| Moonshot Kimi K3 launch | July 2026 | S38 (single source) | [SECONDARY] |
| Moonshot Kimi subscription pricing | 49–699 yuan/month ($7.31–$104.23) | S38 (single source) | [SECONDARY] |


