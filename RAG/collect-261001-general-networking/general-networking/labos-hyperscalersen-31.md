---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-31
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Apple", "Google", "Microsoft", "Nvidia", "OpenAI", "SpaceX", "xAI"]
dates: ["2024-05", "2026-05", "2026-06", "2026-08"]
keywords: ["acquisition", "agent", "agents", "bedrock", "compute", "consumer", "copilot", "distribution", "energy", "foundry", "funding", "gemini"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1182, 1232]
sha256: 04fbc0300095cae673d282607b9b4fdf1521c1330d8cb3a6fe83dec513c151a0
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

- **Colossus 1** (Memphis, TN): ~230,000 GPUs (incl. 30k GB200s), ~500 MW — operational [secondary] https://introl.com/blog/xai-colossus-2-gigawatt-expansion-555k-gpus-january-2026.
- **Colossus 2** (Memphis/Southaven): ~550,000 GB200s/GB300s targeting ~1 GW; "world's first gigawatt-scale AI training supercluster" per company descriptions [secondary].
- **Dec 30, 2025:** Musk announced purchase of a third Memphis building ("MACROHARDRR"), bringing the complex toward **~2 GW total capacity** and **555,000+ GPUs (~$18B spend)** [secondary, citing Musk posts; Reuters reported ultimate target of 2 GW / 1,000,000 GPUs].
- **Power:** purpose-built private power — acquired former Duke Energy natural gas plant site in Southaven, MS (mid-2025); gas turbines + Tesla Megapacks + small solar; near TVA Southaven gas plant [secondary].
- **Regulatory/legal (2026):** on Apr 14, 2026, the Mississippi NAACP + environmental groups (SELC/Earthjustice) filed a **Clean Air Act lawsuit** alleging **27 unpermitted methane gas turbines** at Colossus 2 in Southaven, MS — a ~200 MW makeshift plant feeding the training cluster with no federal air permit; framed as the first major federal test of AI compute vs the Clean Air Act [secondary] https://tech-insider.org/xai-colossus-2-naacp-lawsuit-illegal-gas-turbines-memphis-2026/. Outcome as of Sep 2026 not found — **unresolved**.
- **Financing:** Bloomberg reported (late 2025) a ~$20B equity+debt package for Colossus 2 with **Nvidia as strategic investor (up to $2B)** via an SPV that buys Nvidia GPUs and leases them to xAI; $7–8B equity + up to $12B debt [secondary: Tom's Hardware citing Bloomberg] https://www.tomshardware.com/pc-components/gpus/nvidia-backs-20-billion-xai-chip-deal.
- **Vendors:** Dell ~$5B GPU-server deal reported; Cisco networking hardware for Colossus II buildout; Introl as infrastructure contractor [secondary].
- **Second site:** xAI reportedly built a second data center in Atlanta (~$700M in chips/equipment) [secondary — thin sourcing, treat as unverified].
- **Forward guidance:** possible expansion toward 3 GW by late 2026 discussed [secondary]; one vendor-adjacent source claims 1M GPUs target [unverified].

### 8.11 Funding & valuation (2026) — detailed

| Event | Date | Amount | Valuation | Provenance |
|---|---|---|---|---|
| Series E close | Jan 6, 2026 | $20B (upsized from $15B) | $230B | [secondary: CNBC via coinheadlines] |
| Investors | — | — | — | Nvidia & Cisco (strategic), Valor Equity Partners, StepStone, Fidelity, Qatar Investment Authority, MGX (Abu Dhabi), Baron Capital [secondary] |
| SpaceX acquires xAI | Feb 2, 2026 | All-stock | $1.25T combined ($1T SpaceX / $250B xAI) | [secondary: Reuters, Bloomberg] |
| SPCX Nasdaq IPO | Jun 12, 2026 | **$75B raised** (555.56M Class A @ $135) | ~$1.77T at IPO price; closed day 1 at $160.95 (+19.2%), ~$2.1–2.2T market cap | [secondary: capital.com, zacks.com, ET] |
| Tesla investment | — | Tesla disclosed plans to invest $2B in xAI | — | [secondary — thin, date unspecified, treat as unverified] |
| SPCX status | late Jul 2026 | — | ~$1.49T market cap | [secondary] |

- IPO details: largest IPO in history (prior record Saudi Aramco ~$29.4B); 4× oversubscribed; ~22.5–30% retail allocation; Goldman Sachs lead underwriter; Musk holds ~82.4% voting power, "controlled company"; dual listing Nasdaq Global Select + Nasdaq Texas; options trading from Jun 16 [secondary]. Musk became "world's first trillionaire" on IPO day on combined SpaceX+Tesla stakes [secondary]. Prospectus nuance: SpaceX filed $41.3B accumulated losses since 2002; prospectus describes a satellite-telecom company funding an AI buildout with a launch business attached [secondary].
- Pre-2026 context (RAG continuity): Seed Dec 2023 $134.7M; Series B May 2024 $6B @ $24B; Series C Dec 2024 $6B @ $50B; xAI acquires X Mar 2025 (all-stock, $113B combined: $80B xAI + $33B X); Sep 2025 $10B raise @ $200B [secondary].

### 8.12 Grok products (2026)

- **Grok Build CLI:** launched May 2026 (grok-build-0.1 model at $1.00/$2.00); mid-June 2026 updates: Cursor Composer 2.5 renamed **Grok Composer** inside Grok Build; Build Mode agent surface reached GA in August 2026; versions 0.2.57/0.2.60 shipped Jun 16–21 [secondary].
- **Grok Bot (Aug 11, 2026):** "team of always-on agents that have their own computer, work inside tools and apps, keep working around the clock" [secondary]. Grok 4.7 trained to "natively understand the Grok Bot harness" [vendor-reported].
- **Grok TTS API:** $4.20 per 1M chars; realtime voice API $3/hour; Custom Voices (Apr 30, 2026): ~1 min of recorded speech → production voice clone in <2 min; 80+ built-in voices, 28 languages; two-stage speaker verification [secondary].
- **SuperGrok / SuperGrok Heavy:** $30/mo SuperGrok; **$300/mo SuperGrok Heavy** (launched with Grok 4, Jul 2025) — Heavy gives early access to top models (e.g., Grok 4.3 beta Apr 2026) [secondary]. X Premium+ remains the bundled path for X users [secondary].
- **Consumer availability:** Grok 4.x ships in grok.com, X app, iOS/Android apps; free tier with usage limits; Auto mode + manual model picker; voice mode in mobile apps [secondary].
- **Grok Imagine:** image/video generation; Grok Imagine 1.5 Fast (June 2026) cut 720p video gen from ~40s to ~25s [secondary].
- **Enterprise distribution (2026):** GitHub Copilot, Amazon Bedrock, Microsoft Foundry, Gemini Enterprise Agent Platform [secondary].
- Regulatory headwinds: Grok investigated in multiple jurisdictions over non-consensual explicit deepfakes incl. minors (Jan 2026, post-merger) [secondary].

### 8.13 Personnel — xAI founder exodus

- **Elon Musk:** founder; remains in charge post-merger (SpaceXAI); rang SPCX opening bell from Texas with Gwynne Shotwell on the Nasdaq floor, Jun 12, 2026 [secondary].
- **Founder exodus:** **6 of 12 xAI co-founders have left** (as of Feb 2026 — "exactly half"), departures described as amicable [secondary: TechCrunch, the-decoder] https://techcrunch.com/2026/02/10/nearly-half-of-xais-founding-team-has-now-left-the-company/:
  - Kyle Kosic (infra lead) → OpenAI, mid-2024
  - Christian Szegedy (Google veteran) → Feb 2025
  - Igor Babuschkin → Aug 2025; launched Babuschkin Ventures (AI-safety investing)
  - Greg Yang (Microsoft Research alum) → Jan 2026, health reasons
  - Yuhuai "Tony" Wu (foundational models/reasoning, reported directly to Musk) → Feb 10, 2026
  - Jimmy Ba (U. Toronto/Geoffrey Hinton lineage, reported directly to Musk) → Feb 11, 2026
- **Other exits:** X CEO Linda Yaccarino (Jul 2025); xAI head of legal Robert Keele (early 2026) [secondary].
- **Post-Cursor-acquisition:** Michael Truell (Cursor co-founder) announced Grok 4.6 on X (Aug 12, 2026) as part of SpaceXAI; Eric Jiang leads the Grok 4.3 project; one source cites a **$60B Anysphere/Cursor acquisition** [secondary — acquisition price/date unconfirmed by primary outlets; flag unverified].
- **2026 hires:** no well-documented executive hires found in the sources reviewed — research gap.

### 8.14 xAI key sources

