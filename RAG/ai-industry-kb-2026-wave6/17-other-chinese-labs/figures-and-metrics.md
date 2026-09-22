---
id: ai-industry-kb-2026-wave6/17-other-chinese-labs/figures-and-metrics
title: "Figures and metrics"
domain: other-chinese-labs
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Falcon", "Huawei", "MiniMax", "Moonshot", "OpenAI", "OpenRouter", "Poolside", "StepFun", "Z.ai"]
dates: ["2024-10", "2026-01-05", "2026-01-22", "2026-05", "2026-05-08", "2026-06", "2026-06-12", "2026-06-23", "2026-06-24", "2026-06-30", "2026-07", "2026-07-21", "2026-07-23", "2026-07-31", "2026-08-28", "2026-09-20", "2026-09-27", "2026-10-15"]
keywords: ["apache", "ascend", "benchmark", "benchmarks", "claude", "cost", "deepseek", "distribution", "fp8", "gguf", "glm", "gpt-5.6"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8436, 8554]
section: "§17. Other Chinese Labs"
sha256: c32d174411b5b3e15ed3ecb8d3020197da0349549b018bb4eb0b79203a243f5e
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


| Hy4 Preview backbone / experts | 78 layers (1 dense + 77 MoE); 256 routed + 1 shared, top-8 | S41 (single source) | [SECONDARY] |
| Hy4 Preview MTP layer | 10B total / 0.7B active | S41 (single source) | [SECONDARY] |
| Hy4 Preview API pricing | $0.834 / $2.501 / $0.042 per M (in/out/cache-hit) | S41, S42 | [SECONDARY] |
| Hy4 Preview measured throughput | ~40 tok/s, 3.49 s avg latency (OpenRouter) | S42 (single source) | [SECONDARY] |
| Hy4 Preview BF16 weight size | ~1.56 TB | S42 (single source) | [SECONDARY] |
| Hy4 Preview GGUF builds | STQ1_0 213.66 GiB; Q4_K_M 435.20 GiB; UD-IQ1_M 219.83 GiB | S42 (single source) | [SECONDARY] |
| Hy4 Preview blind eval (vendor) | 2.99/4.00 vs GLM-5.3 2.92, Kimi K3 2.94 (163 experts, 203 tasks) | S41 (single source) | [VENDOR] |
| openPangu 2.0 training | 34T tokens on Ascend NPUs | S45 (single source) | [SECONDARY] |
| openPangu Flash CUDA build (community) | ~56.9 GB resident memory | S78 (single source) | [COMMUNITY] |
| Falcon H1R 7B | DeepConf reasoning model | S47, S48 | [SECONDARY] |
| Falcon-H1 Arabic sizes | 3B / 7B / 34B; up to 256K context | S48 (single source) | [VENDOR] |
| Falcon-H1 Arabic OALL (vendor) | 61.87% / 71.47% / 75.36% (3B/7B/34B) | S48 (single source) | [VENDOR] |
| Poolside deal new terms | $12B pre-money (for $1B); offers to 109/~115 staff; co-founders stay | S50, S51 | [SECONDARY] |
| Poolside fee distribution | to existing investors by end of 2027 | S51 (single source) | [SECONDARY] |
| Poolside raised since 2023 | $1.6B | S51 (single source) | [SECONDARY] |
| Poolside Laguna XS 2.1 / M.1 | 33B/3B, 225B/23B (OpenMDW-1.1) | S51 (single source) | [SECONDARY] |
| Poolside Laguna reported benchmarks | M.1 72.5% SWE-bench Verified; XS 2.1 70.9%; S 2.1 70.2% Terminal-Bench 2.1 | S51 (single source) | [VENDOR] |
| ERNIE 5.0 architecture | ~2.4T native-multimodal MoE; <3% active/query | S55 (single source) | [SECONDARY] |
| ERNIE 5.1 claimed efficiency | 94% lower pre-training cost; total 1/3, active 1/2; −35% latency | S55, S56 | [VENDOR] |
| ERNIE 5.1 third-source pricing | $3/M in, $12/M out (Qianfan API) | S57 (single source) | [SECONDARY] |
| Ling-3.0-Flash license (community) | MIT attributed (was [UNVERIFIED] in base) | S59, S60 | [COMMUNITY] |
| Ling-3.0-Flash context ceiling | scalable to 1M (native 256K) | S59 (single source) | [COMMUNITY] |
| Ling-3.0-Flash AA Index / Omniscience | 38; hallucination 97% → 44% | S58 (single source) | [SECONDARY] |
| Ling-3.0-Flash API pricing | $0.075/M in, $0.22/M out | S59 (single source) | [COMMUNITY] |
| Ling FinFIRST benchmark | 123 tasks / 701 criteria / 12,300 rubric points | S61 (single source) | [VENDOR] |
| Ling-3.0-tiny | 7.9B / 1.3B, fully local | S61 (single source) | [VENDOR] |
| Step 5 Preview active density | ~4.5% (27B/600B) | S63 (single source) | [SECONDARY] |
| Step 5 Preview AA Index | 44; ~99.8 tok/s; $0.71/index task | S63, S64 | [SECONDARY] |
| Step 5 Preview vs GPT-5.6 Sol | 44 vs 47 (max) / 44 (matched speed); 7.4× cheaper output | S66 (single source) | [SECONDARY] |
| MiniMax IPO | HK$4.8–5.54B raised; closed HK$345 (+109.1–109.9%) | S69, S70, S71 | [SECONDARY] |
| MiniMax IPO oversubscription | 1,837× public, 37× international | S71 (single source) | [SECONDARY] |
| Zhipu IPO | HK$4.35B (~$552M); +3.3% open, +13% close day one | S69, S72, S74 | [SECONDARY] |
| MiniMax Stock Connect (Aug 2026) | +22.8% Aug 6; HK$10.6B mainland inflow; 9.7% stake | S73 (single source) | [SECONDARY] |

