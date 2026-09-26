---
id: ai-industry-kb-2026-wave6/11-xai-and-grok/figures-and-metrics
title: "Figures and metrics"
domain: xai-and-grok
role: deep-dive
task: actor-profile
actors: ["Nvidia", "OpenAI", "SpaceX", "xAI"]
dates: ["2025-08-28", "2026-01-06", "2026-02-02", "2026-03-16", "2026-04-15", "2026-04-30", "2026-05", "2026-05-15", "2026-06", "2026-07-08", "2026-07-15", "2026-07-29", "2026-08-12", "2026-09-17", "2026-09-21", "2026-09-22"]
keywords: ["acquisition", "agent", "agents", "agi", "apache", "arr", "astra", "benchmark", "compute", "funding", "gpt-6", "gpu"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5455, 5538]
section: "§11. xAI and Grok"
delta_of: ai-industry-kb-2026
sha256: 3f9647953bc0ba207af154fa1ab59827051515928e317d4c636a4525e9dc6482
---

# Figures and metrics

## Figures and metrics
| Model | Release | Context | Price $/1M in/out |
|---|---|---|---|
| Grok 4.3 | 2026-04-30 | 1M | $1.25/$2.50; 84% cache-hit discount [SECONDARY] |
| Grok 4.5 | 2026-07-08/09 | — | — [SECONDARY] |
| Grok 4.6 | 2026-08-12 | 500K | $2/$6 (short ctx); $0.50 cache [SECONDARY] |
| Grok 4.7 | 2026-09-21 | 500K | $2/$6 ≤200K; $4/$12 >200K [SECONDARY] |
| Grok Code Fast 1 | 2025-08-28 | — | — [SECONDARY] |
| Grok 5 | — | — | not shipped as of 2026-09-22 [UNVERIFIED] |

- Grok 4.7: DeepSWE v1.1 71.0% [VENDOR]; Terminal-Bench 4.0 26% [SECONDARY].
- Grok 4.7 pricing (full): $2/$6 ≤200K tokens; $4/$12 above 200K [SECONDARY].
- Grok Build: `xai-org/grok-build`, Apache 2.0, Rust [SECONDARY].


### New verified metrics — expansion

### xAI / Imagine additional figures
- Series E: $20B at $230B; investors: Valor, StepStone, Fidelity, QIA, MGX, Baron; NVIDIA up to $2B; Cisco. [SECONDARY, S53][SECONDARY, S54][SECONDARY, S56]
- Colossus I+II: 1M+ H100 equivalents end 2025; 600M MAU (X + Grok). [SECONDARY, S54][SECONDARY, S55]
- SpaceX-xAI: $1.25T combined (SpaceX $1T + xAI $250B); date 2026-02-02 vs Sept 2026 breaking report. [SECONDARY, S15][SECONDARY, S58]
- Doosan: 5× 380MW turbines → 600,000+ GB200 NVL72-equivalents. [SECONDARY, S55]
- Imagine: 34M images first week; 6–15s clips <20s; free ~3/day; SuperGrok 500/day. [SECONDARY, S59][SECONDARY, S60][SECONDARY, S62]

### Grok 4 / 4.1 / 4.20 additional figures
- Grok 4: 200K GPUs; 200M H100 GPU hours; 100× training vs Grok 2; 15× compute; AA Index 73; HLE 25.4% / 44.4% (Heavy, tools); ARC-AGI-2 16.2%. [SECONDARY, S44][SECONDARY, S45][SECONDARY, S46][SECONDARY, S47]
- xAI Series C: $6B (Dec 2024) at $24B valuation. [SECONDARY, S46]
- Grok 4.1: EQ-Bench3 1586 (vs 1206); FActScore 2.97% (vs 9.89%); 65% stealth preference; free 5–10 daily queries. [SECONDARY, S48][SECONDARY, S49][SECONDARY, S52]
- Grok 4.20: 4 agents; 256K (up to 2M) ctx; trained Jan 2026 completion; $30/mo SuperGrok access. [SECONDARY, S50]


### Model benchmark figures (all vendor-reported unless noted)
- Grok 4: Humanity's Last Exam ~38–44% (third-party/vendor claims). [SECONDARY, S6]
- Grok 4.1: hallucinations 12.09% → 4.22% (−65%); LMArena 1483 Elo. [VENDOR, S1]
- Grok 4.3: AA Intelligence Index 38 (v4.1). [SECONDARY, S12]
- Grok 4.5: TB 2.1 83.3%; SWE-Bench Pro 64.7%; SWE Marathon 29.0% pass@1; 4.2× token efficiency (15,954 vs 67,020 tokens). [VENDOR, S8][VENDOR, S10][VENDOR, S11]
- Grok 4.5: AA Intelligence Index 54 (v4.1, high reasoning), #4 overall at release. [SECONDARY, S12]
- Grok 4.6: CursorBench v3.2 69.9%; DeepSWE v1.1 65.9%; FrontierCode v1.1 Extended 61.3%; APEX-Agents 57.5%. [VENDOR, S25]
- Grok 4.7: CursorBench 4.0 46.3%; DeepSWE v1.1 71.0%; EEBench 64.0%; AA Briefcase v1.1 1,657; Harvey Legal 19.6%; HealthBench Pro 56.7%; TB 4.0 38.0%; GDPval Elo 1,695. [VENDOR, S43]
- Grok 4.7: AA Intelligence Index 46 (xhigh); Coding Agent Index 56 (with Grok Build). [SECONDARY, S41][SECONDARY, S38]
- Grok 4.7 token intensity: ~81K output tokens/Index task (xhigh) vs 36K (4.6 high) vs 27K (GPT-6 Astra max). [SECONDARY, S41]
- Grok 4.7 safeguards: 3.3% risky-prompt pass-through on HackerBench v0.3 (internal). [SECONDARY, S42]
- Imagine video: AA Video Arena image-to-video Elo 1404 ±6 (late May 2026), #1, displacing Seedance 2.0. [SECONDARY, S36]
- Imagine Image 2.0: #2 globally text-to-image and image editing (Aug 7, 2026 Arena snapshot). [SECONDARY, S32]

### Pricing figures
- Grok 4 (2025): Heavy $300/mo; SuperGrok $30/mo. [SECONDARY, S5]
- Grok 4.5: $2/$6 per 1M; premium faster variant $4/$18; CursorBench $1.51/task. [SECONDARY, S8][SECONDARY, S10][SECONDARY, S11]
- Grok 4.6: $2/$6; cached $0.50/M; long-context pricing ≥200K tokens. [VENDOR, S30]
- Grok 4.7: $2/$6 (<200K), $4/$12 (≥200K), cached $0.50/$1.00; Fast 2× standard, not on public API. [VENDOR, S44][SECONDARY, S40]
- Voice API: Agent $0.08/min; TTS $15.00/1M chars; STT batch $0.10/hr; streaming $0.20/hr. [VENDOR, S44]

### Compute and corporate figures
- Colossus 1: 230K GPUs (150K H100 / 50K H200 / 30K GB200); ~500 MW; 122-day build. [SECONDARY, S14]
- Colossus 2: 550K GB200/GB300; >1 GW (targets 1.5 GW, then 2 GW); 2 GPUs/node → >1M accelerators site-wide. [SECONDARY, S14]
- Memphis complex: ~555K GPUs for ~$18B (Feb 2026); ~2 GW with third building; 1M-GPU target Q2–Q3 2026. [SECONDARY, S13][SECONDARY, S15]
- 2030 target: 50M H100-equivalent GPUs. [SECONDARY, S21]
- Series E: $20B at $230B (2026-01-06); SpaceX acquisition all-stock @ ~$1.25T combined (2026-02-02). [SECONDARY, S21][SECONDARY, S15][SECONDARY, S20]
- Funding history: seed $134.7M (Jul 2023); Series C $6B @ $24B (Nov 2024); X merger @ ~$113B (Mar/Apr 2025); total ~$26.7B raised. [SECONDARY, S20]
- Scale: ~$500M ARR early 2026 (est., $2B FY26 target); ~600M MAU (X + Grok apps). [SECONDARY, S20][SECONDARY, S22]
- Cursor/Anysphere acquisition: $60B, announced mid-June 2026, close expected Q3 2026. [SECONDARY, S10]

## Main actors
- **xAI** — lab behind the Grok line; skipped Grok 4.4; cut context windows from 1M (4.3) to 500K (4.6/4.7) while layering context-window pricing cliffs [SECONDARY].
- **Elon Musk** — publicly backed the July 24 open-weights letter via quote-post ("This has my full support") without signing [SECONDARY].
- **Artificial Analysis / alextech** — secondary trackers carrying Grok benchmark figures (TB 4.0 26%, AA mirrors) [SECONDARY].

## Timeline and context
- **2025-08-28** — Grok Code Fast 1 announced (codename Sonic) [SECONDARY].
- **2026-03-16** — xAI Text-to-Speech GA [SECONDARY].
- **2026-04-15** — Speech-to-Text GA [SECONDARY].
- **2026-04-30** — Grok 4.3 released (1M context, $1.25/$2.50, 84% cache-hit discount) [SECONDARY].
- **2026-05-15** — Grok Build CLI beta [SECONDARY].
- **2026-07-08/09** — Grok 4.5 public API (V9 base) [SECONDARY].
- **2026-07-15** — Grok Build open-sourced (`xai-org/grok-build`, Apache 2.0, Rust) [SECONDARY].
- **2026-07-29** — Voice Think Fast 2.0 [SECONDARY].
- **2026-08-12** — Grok 4.6 (post-training on V9 base, 500K, $2/$6) [SECONDARY].
- **2026-09-17** — Voice Transcribe 2.0 [SECONDARY].
- **2026-09-21** — Grok 4.7 (500K; $2/$6 ≤200K, $4/$12 above) [SECONDARY].
- **2026-09-22** — Grok 5 still unshipped [UNVERIFIED].


### New verified timeline entries — expansion

