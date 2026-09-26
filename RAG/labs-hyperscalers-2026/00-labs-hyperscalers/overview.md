---
id: labs-hyperscalers-2026/00-labs-hyperscalers/overview
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Anthropic", "Apple", "Broadcom", "Cerebras", "China", "Cohere", "CoreWeave", "EU", "Fireworks AI", "Fluidstack", "Google", "Groq", "Hugging Face", "Meta", "Microsoft", "Mistral", "Nebius", "Nscale", "Nvidia", "OpenAI", "Oracle", "Perplexity", "Sakana", "SpaceX", "TSMC", "Together AI", "United States", "xAI"]
dates: ["2026-02-01", "2026-04-27", "2026-05-13", "2026-09-22"]
keywords: ["acquisition", "agent", "agents", "amd", "antitrust", "apache", "astra", "backlog", "bedrock", "benchmark", "benchmarks", "capex"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1, 44]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 46be7e69353ffced4d6cecba3350af0e016c0d58591a9ed7beb37b296a4ec5a3
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)
## Final consolidated research report (English) — compiled September 22, 2026

**Project:** RAG data-collection, Step 3 (Labos / hyperscalers)
**Coverage window:** February 1, 2026 → September 22, 2026
**Language:** English
**Consolidation date:** September 22, 2026
**Source tracks merged:** Track A (Anthropic + OpenAI), Track B (Google/DeepMind + Meta + Apple + Microsoft + Amazon), Track C (xAI/Grok + Mistral + independent labs), Track D (NVIDIA NIM + Groq + Cerebras + SambaNova + Together AI + Fireworks AI + Nebius + CoreWeave + "Cabreras" + FreeLLMAPI)
**Status:** Research snapshot. Prices, valuations, benchmark figures and star counts are dated snapshots; re-verify against official pages before operational use.

### Scope
- **§1 Anthropic** — Claude model releases (Opus 4.6→Opus 5.5, Sonnet 4.6→Sonnet 5, Fable 5/5.1, Mythos 5), funding (Series G $30B → Series H $65B, $380B→$965B valuations, $14B→$100B revenue run rate), products (Claude Code, Cowork, Managed Agents), controversies (Fable export-control suspension, $1.5B copyright settlement, safety resignations, Amodei's "Pace the Frontier" essay, Sept 18 antitrust class action), infrastructure (Fluidstack $50B, Nscale $45B, 12+ LOIs >1 GW), IPO (confidential S-1 filed Jun 1, 2026; target Nov 2026).
- **§2 OpenAI** — GPT-5.3-Codex→GPT-5.6 Sol/Terra/Luna→GPT-6 Astra, $122B round at $852B, $25B→$40B+ revenue run rate, ChatGPT tiers, Codex, ChatGPT Work, Atlas/Sora shutdowns, April executive exodus, Shazeer hire, Microsoft restructure, Bedrock, DoD deals, Hugging Face sandbox-escape incident, Nippon Life suit, EU AI Act incident report, EO 14409 gating, Stargate $500B, Oracle $300B, AMD 6 GW, Broadcom.
- **§3 Google / DeepMind** — Gemini 3.1 Pro / 3.5 Flash / 3.6 Flash, Ironwood TPU 7th gen GA, TPU 8t/8i announced, Gemini Enterprise Agent Platform (Vertex AI renamed), Alphabet capex $195–205B, $40B Anthropic investment commitment (reported), Isomorphic Labs $2.1B Series B, leadership overhaul (Hassabis → Alphabet Chief Scientist), personnel departures.
- **§4 Meta** — Llama 5 unverified; Muse Spark 1.0→1.3 (first paid Meta API), Muse Glimmer 30B open weights (Apache 2.0), Zuckerberg open-source essay, FAIR research, capex $130–145B, compute 7 GW→14 GW, MTIA "Iris" chip.
- **§5 Apple** — Apple Intelligence 2026 / iOS 27, Siri AI rebuild, AFM 3 (five models), Apple–Google Gemini partnership (~$1B/yr reported), Baltra AI server chip (Broadcom/TSMC), 12 GB device gating, EU/China launch constraints.
- **§6 Microsoft** — MAI model family (7 models, Build 2026), Phi-4-reasoning-vision-15B, Aion 1.0 on-device SLM, Microsoft Foundry, Microsoft–OpenAI April 27, 2026 restructure, capex ~$175–190B, Maia 200 in production, Maia 300 targeted.
- **§7 Amazon** — Nova 2 family consolidation (Premier/Omni/Reel/Canvas wound down; Frontier Model Research under Pieter Abbeel), OpenAI models on Bedrock (GA Jun 1, 2026), Grok on Bedrock, Trainium3 GA, $50B OpenAI investment (closed Jul 31, 2026), Alexa+ GA (US Feb 4; India Sep 16–17), capex ~$220B.
- **§8 xAI / Grok** — Grok 4.20→4.7 releases, xAI–SpaceX merger (Feb 2, 2026; combined $1.25T), SPCX IPO ($75B raised, Jun 12, 2026), Colossus compute cluster (~2 GW), Colossus 2 Clean Air Act lawsuit, SuperGrok tiers, Grok Bot, Grok Build.
- **§9 Mistral AI** — Voxtral Transcribe 2, Small 4 (119B MoE), Medium 3.5 (128B dense), OCR 4, Le Chat→Vibe rebrand, €3B Series D at >€21B valuation (Sep 8, 2026), $830M debt raise, Airbus/BMW/EDF/Microsoft/HUMAIN/Cloudera partnerships, sovereign infrastructure (Essonne 13,800 GB300s).
- **§10 Independent labs** — Thinking Machines (Inkling 975B/41B MoE open weights, Tinker), Sakana AI (Fugu Max/Ultra v2 orchestration, AI Scientist in Nature), Cohere (Command A+ 218B/25B MoE Apache 2.0; definitive merger with Aleph Alpha signed Sep 16, 2026), AI21 (Jamba Mini 2, Gateway beta), Aleph Alpha (PhariaAI pivot), Perplexity (>$750M annualized revenue; Nvidia talks for $30B+ round).
- **§11 NVIDIA NIM** — Build free API (100+ models), Nemotron 3 family (Ultra/Super/Nano/Mini-4B), Llama Nemotron, NIM for RTX AI PCs, DGX Spark, AI Enterprise licensing ($4,500/GPU/yr), Vera Rubin production, NVIDIA–Groq licensing context, NVIDIA→Hugging Face acquisition implications.
- **§12 Groq** — ~$20B NVIDIA non-exclusive licensing deal (late Dec 2025), Groq 3 LPU (1,500 tok/s target), Groq 3 LPX in full production, GroqCloud pricing ($0.05–$0.90/M input), $650M raise + reported $350M Series A (Aug 2026), neocloud pivot.
- **§13 Cerebras, SambaNova, Together AI, Fireworks AI, Nebius, CoreWeave** — Cerebras IPO ($5.55B, May 13, 2026; $185/share), OpenAI $20B+ Cerebras contract, CS-4 (Aug 2026); SambaNova Series F $1B @ $11B (Jul 8, 2026), SN50; Together AI Series C $800M @ $8.3B (Jul 1, 2026), IBM $240M cluster; Fireworks Series D $1.505B @ $17.5B (Jul 16, 2026); Nebius NVIDIA $2B warrant, $4.0B + $5.75B converts, Meta ~$27B orders; CoreWeave $3B converts + 35M-share ATM (Sep 17, 2026), $104.2B backlog.
- **§14 "Cabreras" disambiguation + FreeLLMAPI** — "Cabreras" is a misspelling of Cerebras; FreeLLMAPI is a community self-hosted free-tier aggregator (ToS risk), not a provider.
- **§15 Master chronological timeline** (150+ dated entries, Feb 1 → Sep 22, 2026).
- **§16 Cross-lab comparisons** — funding/valuation race, independent benchmarks, API pricing, personnel moves, infrastructure scoreboard.
- **§17 M&A tracker** (Feb–Sep 2026).
- **§18 Master uncertainty / verification log.**
- **§19 Collection metadata.**

### Provenance legend (used throughout)
- **[official]** — vendor's own blog, docs, changelog, repository, earnings statement, or regulatory filing.
- **[vendor-reported]** — figure claimed by the vendor (benchmarks, pricing, user counts) without independent audit.
- **[independent]** — reputable third-party press/analysis (Bloomberg, Reuters, CNBC, NYT, WSJ, FT, AP, SiliconANGLE, VentureBeat, The Information, The Register, SemiAnalysis, TechCrunch, TechTarget, CSA).
- **[secondary]** — lower-tier press, blogs, analyst summaries, community trackers; useful but unverified.
- **[unverified]** — single-source or conflicting claims; treat as uncertain.

### Conventions kept
- Every price, valuation, version, and benchmark figure is dated or versioned; no undated numbers presented as current.
- Vendor-reported benchmark numbers produced on each lab's own scaffolding are NOT directly compared across vendors; cross-lab ranking prefers Artificial Analysis Intelligence Index (with index version), Vals AI, BenchLM, METR, and SecureBio figures.
- Conflicting figures are presented side-by-side rather than silently reconciled.
- Artificial Analysis Intelligence Index scores are stored with their index version (different 2026 index revisions are not comparable).

