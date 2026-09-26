---
id: labs-grok-platforms-2026/11-part-a-xai-grok/overview
title: "PART A — xAI / GROK"
domain: part-a-xai-grok
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Google", "Nvidia", "OpenAI", "SpaceX", "Xiaomi", "xAI"]
dates: ["2024-05", "2025-03", "2025-07", "2026-02-02", "2026-05", "2026-06", "2026-09-21"]
keywords: ["grok", "acquisition", "agent", "arr", "astra", "benchmark", "benchmarks", "compute", "context window", "cost", "fable 5", "gemini"]
source: docs/RAG/Labos, Grok, outils & plateformes_EN.md
source_anchor: ""
source_lines: [216, 304]
section: "PART A — xAI / GROK"
sha256: 159309b5a49cf32a542bc1dc320ac3ae82994aedf782c054546da23b642ae21e
---

# PART A — xAI / GROK

## 1. Grok 4.7 — released September 21, 2026

### 1.1 Release & positioning
- **Released September 21, 2026** by SpaceXAI (post-merger branding of xAI), replacing Grok 4.6 as the flagship model. Announced via the @SpaceXAI account the same day: "Grok 4.7 works longer on difficult tasks, checks its work more carefully, and comes with our strongest safeguards to date."
- Pitched explicitly as a **coding-and-knowledge-work model**, not a general-purpose upgrade — launch benchmarks led with CursorBench 4.0 and DeepSWE v1.1 (long-running, multi-file coding) rather than MMLU/GPQA.
- Available at launch in **Cursor, Grok Build, the xAI API (model ID `grok-4.7`), third-party coding harnesses, model routers, and cloud platforms**; Grok Build offers free trial access.

### 1.2 Specs
| Property | Value |
|---|---|
| Base model | New, larger base than Grok 4.6 |
| Context window | 500,000 tokens |
| Knowledge cutoff | May 2026 |
| Modalities | Text + image input, text output |
| Reasoning effort | low / medium / high (default) / **xhigh** (new top tier) |
| APIs | Responses API, Chat Completions |
| Tools | Function calling, web search, X search, code execution |
| Training | Longer RL run on harder multi-hour tasks; self-verification; native Grok Bot harness training |

### 1.3 Benchmarks (vendor-reported launch table: Grok 4.7 xHigh vs Grok 4.6 High, GPT-5.6 Sol Max, Fable 5.1 Max)
| Benchmark | Grok 4.7 xHigh | Grok 4.6 High | GPT-5.6 Sol Max | Fable 5.1 Max |
|---|---|---|---|---|
| CursorBench 4.0 | **46.3%** | 40.4% | 41.7% | **51.8%** |
| DeepSWE v1.1 | 71.0%* | 65.2% | **72.7%** | 70.0% |
| EEBench | **64.0%** | 53.0% | 39.4% | 56.4% |
| AA Briefcase v1.1 | 1,657 | 1,546 | 1,487 | **1,678** |
| Terminal-Bench 4.0 | 38.0% | 20.3% | 37.3% | **57.9%** |
| Harvey Legal Agent | **19.6%** | 15.8% | 2.5% | 6.7% |
| HealthBench Professional | 56.7% | 48.5% | 60.5% | **62.1%** |
*\*DeepSWE run at high effort. GDPval professional knowledge work: Grok 4.7 xhigh 1,695 Elo (up from 1,605 for 4.6); Fable 5.1 max 1,735; GPT-6 Astra max 1,542.*
- **Independent check (Artificial Analysis Intelligence Index): Grok 4.7 = 46 vs Grok 4.6 = 44** — a smaller step than vendor charts suggest, and a caveat: at max reasoning, Grok 4.7 consumed ~240M output tokens per index task vs ~94M for 4.6 (different settings, but per-task cost rises).
- Notable: Terminal-Bench 4.0 nearly doubled (20.3% → 38.0%); EEBench top of table (64.0%); Harvey legal agent 19.6% vs Fable 5.1's 6.7%.

### 1.4 Pricing & availability
- **$2.00 / 1M input, $6.00 / 1M output** (below 200K tokens), $0.50 / 1M cached input — **identical to Grok 4.6** (a version upgrade with no price increase).
- Above 200K tokens the whole request bills at $4 / $1 / $12.
- **Grok 4.7 Fast**: 2× output speed at 2× price; **not available through the public API**.
- Better document/presentation generation claimed (GDPval/AA Briefcase) for lawyers, nurses, financial analysts.

### 1.5 Safety stack
- Entirely new safeguard stack (not an iteration of 4.6's). Internal **HackerBench v0.3** (risky-prompt pass-through): **3.3%** — framed as a meaningful drop in unsafe completions, though no full system card at Anthropic/OpenAI's September-launch level of detail.

### 1.6 vs predecessors (4.6 / 4.20)
- **Grok 4.20** (Feb 17, 2026, public beta): four-agent "council" architecture (Grok/Harper/Benjamin/Lucas), adversarial consensus, hallucination 12% → 4.2%, 256K–2M context, continuous weekly learning — from Vague 1 research.
- **Grok 4.6** (mid-2026): 500K context baseline, CursorBench 40.4%, Terminal-Bench 4.0 20.3%.
- **Grok 4.7**: larger base model, longer RL on multi-hour tasks, new **xhigh** reasoning tier, new safety stack, same price/speed as 4.6. Independent AA Index: 46 (tied with Xiaomi MiMo-V2.6-Pro on v4.3 — best open-weight score, and Grok 4.7's 46 ties it as a closed model).

## 2. SpaceX acquired xAI — February 2, 2026

### 2.1 Deal terms
- **Announced February 2, 2026** in an internal SpaceX memo (Musk); structured as an **all-stock share exchange** — the largest merger on record at announcement.
- **Combined valuation ~$1.25 trillion**: SpaceX ~$1 trillion + xAI ~$250 billion.
- **Exchange ratio: 1 xAI share → 0.1433 SpaceX shares** (xAI stock $75.46/share; SpaceX $526.59/share per bank documents).
- Assets combined: **SpaceX, Starlink, xAI/Grok, and the X social platform** (xAI had acquired X in March 2025 for a combined $113B: $80B xAI + $33B X).
- **Tesla excluded** — a public company, merging it raised legal/governance complications. Tesla had previously invested ~$2B in xAI, converting into an indirect slice of the merged company.
- **Tax-free reorganization**: xAI shareholders defer taxes until sale; executed via two Nevada intermediary companies, avoiding triggering xAI's debt covenants (xAI inherited $12B of X debt in 2025 and had taken on ≥$5B more since).

### 2.2 Valuation trajectory (blended from CNBC, Bloomberg, Forbes, WSJ reporting)
| Date | Milestone | Valuation |
|---|---|---|
| Dec 2023 | Seed ($134.7M) | undisclosed |
| May 2024 | Series B ($6B) | $24B |
| Dec 2024 | Series C ($6B) | $50B |
| Mar 2025 | xAI acquires X (stock swap) | $113B combined |
| Sep 2025 | Equity raise ($10B) | $200B |
| Jan 2026 | Series E ($20B; Valor, Fidelity, QIA, MGX, Nvidia, Cisco) | **$230B** |
| Feb 2, 2026 | SpaceX acquisition (all-stock) | **$1.25T combined** |
| Jun 12, 2026 | SPCX Nasdaq IPO ($75B raised) | $135/share, ~$2.1T intraday ATH; IPO targeting $1.75T market cap |

Context: xAI generated only ~$107M revenue in the September quarter (net loss $1.46B, widening from $1B) — roughly **~460× revenue multiple** at the $230B Series E, vs ~20× for Anthropic ($965B / $47B ARR) and ~34× for OpenAI ($852B / $25B ARR) per July-2026 estimates.

### 2.3 Strategic rationale
1. **Capital**: xAI was burning ~$1B/month racing OpenAI/Anthropic; folding into SpaceX gave it a far larger balance sheet (SpaceX: ~$8B profit on $15–16B revenue).
2. **Orbital data centers**: Musk's memo framed space-based compute as the strategic core — the thesis that space becomes the cheapest place to run AI compute within 2–3 years.
3. **Vertical integration**: "the most ambitious vertically integrated innovation engine on (and off) Earth" — rockets (Starship), satellite internet (Starlink), AI (Grok), social data (X) under one cap table.
4. **IPO staging**: the merger set up the largest IPO in history (June 2026, $40–80B raise range reported pre-IPO).

## 3. Pentagon integration of Grok

### 3.1 The deal
- **$200M contract** (initial value, with potential expansion to $800M) — part of a broader July 2025 DoD initiative awarding $200M each to **xAI, OpenAI, Google, and Anthropic** to develop an "AI arsenal" for critical national security challenges, aligned with the DoD's Responsible AI Strategy.
- Signed **December 22** (2025); deployment **began early 2026** into **GenAI.mil**, the Pentagon's centralized generative-AI platform.
- **Scale: ~3 million military and civilian personnel** get access.
- **Security tier: Impact Level 5 (IL5)** — cleared for Controlled Unclassified Information (CUI) in daily workflows (research, writing, analysis, training).
- **Differentiator: real-time X data feed** — military planners get AI-assisted live social-sentiment and breaking-news intelligence from X, a capability rival models can't offer. Announced Feb 23, 2026 confirmation of Grok supporting rapid geospatial/signals-intelligence analysis for commanders.
- Alphabet's Gemini was the first model on the platform; Grok joined a multi-vendor roster.

