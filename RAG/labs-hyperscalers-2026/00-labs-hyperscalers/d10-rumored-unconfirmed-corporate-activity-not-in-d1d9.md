---
id: labs-hyperscalers-2026/00-labs-hyperscalers/d10-rumored-unconfirmed-corporate-activity-not-in-d1d9
title: "D10. Rumored / unconfirmed corporate activity (not in D1–D9)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "Cohere", "EU", "Google", "Hugging Face", "Meta", "Microsoft", "Mistral", "Nscale", "Nvidia", "OpenAI", "OpenRouter", "Perplexity", "Sakana", "SpaceX", "Stripe", "United States", "xAI"]
dates: ["2026-05-28", "2026-09"]
keywords: ["agent", "agents", "apache", "astra", "benchmarks", "capex", "claude", "cohere", "compute", "consumer", "cost", "distribution"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [2480, 2535]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: f4c73876155ed4b6951ec4f876048c76fc676cfc242bf32aec806acf237d2374
---

# D10. Rumored / unconfirmed corporate activity (not in D1–D9)

### D10. Rumored / unconfirmed corporate activity (not in D1–D9)

| Date surfaced | Claim | Status |
|---|---|---|
| Jun 2026 | Mistral **$3.5B raise at ~$23B** (leaked internal timeline) | Superseded by Sep 8 Series D (€3B @ €21B+) — do not double-count |
| Aug 2026 | **Nvidia → Perplexity** round participation at $30B+ | Talks reported (The Information); unconfirmed |
| Sep 2026 | Thinking Machines **$5–6B at ≥$40B** (Accel leading; Nvidia ~$2.5B) | Talks reported (Forkast); unconfirmed |
| 2026 | **Tesla → xAI $2B** investment plan | Thin sourcing; date unspecified |
| 2026 | **OpenAI → TBPN + unnamed dev-tools startup** | Names/terms not returned |
| Jan 2026 | Thinking Machines **$50–60B follow-on** | Collapsed (see D8) |

### M&A scoreboard — 2026 AI consolidation wave

> Five structural themes: (1) **compute-merges-AI** (SpaceX–xAI); (2) **sovereign stacks** (Cohere–Aleph Alpha, Mistral–Microsoft/HUMAIN); (3) **distribution buys models' routes** (Stripe–OpenRouter, NVIDIA–Hugging Face); (4) **acqui-hires for agents** (Ona, Tomoro, Emmi, Reliant); (5) **failed mega-rounds** (Thinking Machines $50–60B talks collapse). 2026 is the year AI M&A shifted from talent-acquisitions to balance-sheet-scale combinations. [research finding]

### 16.10 Pricing methodology & limitations

- **Vendor-reported vs measured:** per-token list prices are vendor-published; $/task figures (Artificial Analysis) are measured on that index's harness at a point in time and embed that revision's task mix. Never blend the two without labeling.
- **Discounts omitted:** cached-input, batch-API (typically 50% off), committed-use, and education/nonprofit discounts are noted only where material (Grok 4.3 cached $0.20; Luna cut; OCR 4 batch $2). Effective prices for scaled buyers are lower than list.
- **Currency:** Mistral consumer pricing in EUR (€14.99 Pro); enterprise/API in USD. Series D in EUR (€3B @ €21B+ ≈ $24B at Sep 2026 rates).
- **Time-sensitivity:** the steepest 2026 moves — Grok 4.3 output −83% (Apr 30), Luna input −80% (Jul 30), Anthropic Opus/Sonnet input +60%/+33% (Aug 6, the rare increase) — show pricing is a strategic weapon, not a cost curve. Re-verify any figure before citing beyond Sep 2026.
- **Gaps:** no public per-token list for NVIDIA NIM, Muse Spark API, Command A+, Inkling, or Jamba Mini 2 as of Sep 22, 2026 (see §18).

### 16.11 Open-weight vs closed — September 2026 scoreboard

| Dimension | Closed frontier | Open weights |
|---|---|---|
| Top models | GPT-6 Astra, Claude Fable 5/5.1, Gemini 3.x, Grok 4.7, Muse Spark 1.3 | Inkling (975B/41B), Command A+ (218B/25B), Mistral Large 3 (675B/41B), Nemotron 3 Ultra (500B), Mistral Medium 3.5 (128B dense) |
| Best-reported coding | GPT-5.6 Sol T-Bench 2.1 88.8% (vendor) | Medium 3.5 SWE-Bench Verified 77.6% (vendor) |
| Cheapest frontier-class API | Gemini 3.6 Flash $1.50/$7.50; Grok 4.7 $2/$6 | Mistral Small 4 $0.15/$0.60 (API); most weights $0 self-host |
| Licenses | Proprietary | Apache 2.0 dominates (Inkling, Command A+, Small 4, Large 3, Jamba); NVIDIA Open Model License (Nemotron); modified MIT (Medium 3.5) |
| Momentum | Export controls (Jun 12) push some buyers toward open | Sakana Fugu routes *around* closed models for "supply-chain resilience"; EU sovereign AI prefers open |

**Read:** September 2026 is the high-water mark for open weights — five Apache 2.0 models at ≥100B total params, two above 500B. But the very top of coding/agent benchmarks is still closed, and export controls are accelerating open-weight adoption outside the US. [research finding]

### 16.12 Compute & data-center scoreboard (Sep 2026)

| Owner | Site / program | Scale | Status |
|---|---|---|---|
| xAI (SpaceXAI) | Colossus 1, Memphis | ~230k GPUs, ~500 MW | Operating |
| xAI (SpaceXAI) | Colossus 2, Memphis/Southaven | ~550k GB200/GB300, ~1 GW | Building; Clean Air Act suit filed Apr 14, 2026 |
| xAI (SpaceXAI) | Memphis complex total | ~555k+ GPUs, ~2 GW target (1M GPUs discussed) | Expanding |
| Anthropic + US gov | US compute partnership | $200B / 10 GW | Announced Feb 12, 2026 |
| Anthropic + UK/Nscale | Stargate UK | £31B / 4 GW | Announced Feb 12, 2026 |
| Anthropic + Nvidia + Fujitsu | UK Stargate | $200B | Announced Aug 18, 2026 |
| Anthropic + Microsoft | Azure Florida | $50B / 1 GW (2026–2030) | Announced May 28, 2026 |
| Thinking Machines + Nvidia | Vera Rubin procurement | ≥1 GW from 2027 | Signed Mar 30, 2026 |
| Mistral | Essonne + Les Ulis + Sweden | <200 MW today → 1 GW by 2030 | Operating/building |
| Schwarz Group | German campus (StackIT) | €11–13B, up to 100k chips | Planned (tied to Cohere merger) |
| Meta | Fleet expansion | Hundreds of thousands of GPUs; historic-high capex | Ongoing |
| Google | TPU fleet (Ironwood GA Mar 31) | 7th-gen; internal + Cloud | GA |
| Amazon | Trainium3 / Project Rainier | $1.80/chip-hour quoted; 2 GW OpenAI commit | Deploying |
| Microsoft | Aion supercomputing; Maia silicon | Next-gen in 2026 | Building |

**Read:** announced 2026 commitments exceed **15 GW** of AI compute across these programs — the bottleneck has moved from models to power and silicon supply. [research finding — summed from dated entries]

