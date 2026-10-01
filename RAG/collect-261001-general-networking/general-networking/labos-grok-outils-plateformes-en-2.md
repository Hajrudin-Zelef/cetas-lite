---
id: collect-261001-general-networking/general-networking/labos-grok-outils-plateformes-en-2
title: "VOLET 1 — Vague 2 (EN)"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "EU", "Glasswing", "Microsoft", "OpenAI", "United States"]
dates: ["2026-01", "2026-03", "2026-05", "2026-06", "2026-06-09", "2026-07-01", "2026-07-24", "2026-09", "2026-09-30"]
keywords: ["agent", "agentic", "agi", "alignment", "astra", "benchmark", "benchmarks", "claude", "cost", "cyber", "cybersecurity", "distillation"]
source: docs/RAG/collect-261001-general-networking/Labos, Grok, outils & plateformes_EN.md
source_anchor: ""
source_lines: [77, 141]
sha256: 6e3e1110bac1adc43971743bc779f2beadf86281d6a6447d76add9db6e4092c9
---

# VOLET 1 — Vague 2 (EN)

Context needed to interpret Opus 5.5 deltas:
- **Released July 24, 2026** at **$5/$25 per 1M** — identical to Opus 4.8 ("the price line did not move but the capability line did").
- **1M-token context, 128K max output, May 2026 knowledge cutoff** (most current Claude at the time). API ID `claude-opus-5`. Default on Claude Max; strongest on Claude Pro.
- **Adaptive thinking on by default** (new vs 4.8); effort tiers low/medium/high/xhigh/**max**; Fast mode ~2.5× speed at $10/$50.
- Key vendor figures: **SWE-bench Verified 96.0%** (highest publicly verified at the time), Frontier-Bench v0.1 **43.3%** (2× Opus 4.8's 21.1%), ARC-AGI-3 **30.2%** (vs 1.5% for 4.8), OSWorld 2.0 **70.6%**, AutomationBench 26.0%, GDPval-AA v2 Elo **1,861**.
- Independent: **BenchLM #1 across 215 models (85.88)** — above Mythos 5 (83.01), Fable 5 (82.76), GPT-5.6 Sol (81.46); **AA Intelligence Index 61**, one point above Fable 5.
- vs Fable 5: effectively tied/leading on repo-level coding (SWE-bench Pro 79.2% vs 80.3% — Fable's only coding win), ahead on Frontier-Bench (+9.6pp), at half the price.

---

# 3. Rumors: Claude Opus 5.2 (never publicly released)

## 3.1 What was observed
- **Sept 14–15, 2026:** multiple developers reported Opus 5 requests inside Claude Code being silently routed to a different backend checkpoint identified as **Opus 5.2** — gray-scale testing pattern (same pattern as past Anthropic pre-releases: limited surfaces first — Claude Code, Cursor, Vertex — before API GA).
- **Sept 16, 2026:** model slug `claude-opus-5-2` surfaced in **Microsoft Foundry** with a config listing reasoning effort tiers **low / medium / high / xhigh / max**.
- **Prediction markets:** ~80% chance of a next Opus by September 30, 2026 (Polymarket).
- **Tester reports (consistent across accounts):** noticeably faster than Opus 5, less "overthinking," more direct output; no longer "lazy" — persists on long tasks instead of stopping partway (matters for per-token billing).

## 3.2 Status as of 22 Sept 2026: superseded
- Anthropic never published a model card, API ID, pricing, or benchmarks for Opus 5.2; the catalog still listed `claude-opus-5` as latest Opus until the 5.5 launch.
- **On Sept 22, Anthropic shipped Opus 5.5 instead** — the rumored 5.2 never materialized as a public release. Community analysis (YouTube forensic coverage) suggests the Sept 14–15 routing was early testing of what became 5.5 (internal checkpoint "Wafer-EAP"), and that Anthropic **skipped 5.1 and 5.2 on the Opus tier entirely** (the only 5.1 release was Fable 5.1 — a different tier).
- Predicted pricing had been $5/$25 (Opus-tier stability); actual 5.5 launched at $4/$20.
- RAG note: treat "Opus 5.2" as an unreleased internal checkpoint, not a product. The rumor trail is useful for release-cadence modeling (Anthropic tests in Claude Code ~1 week before launch).

---

# 4. Claude Fable 5 (9 June 2026) & Fable 5.1 (1 Sept 2026)

## 4.1 Origin: the March 2026 "Capybara/Mythos" leak
- ~27 March 2026: ~3,000 unpublished Anthropic assets (draft blog posts, images, internal docs) exposed via a misconfigured CMS data store; discovered by LayerX Security / Cambridge researchers. Revealed a new **"Capybara" tier** with first model **"Mythos"** — sitting above Opus. Anthropic confirmed as human error; no weight breach.

## 4.2 Fable 5 — the Mythos-class public tier
- **Released June 9, 2026.** API ID `claude-fable-5`. First "Mythos-class" model generally available — Anthropic's most powerful public model at launch.
- **The split:** Fable 5 and restricted twin **Claude Mythos 5** share the same underlying weights; Fable ships with safety classifiers intercepting cybersecurity, biology, chemistry, and model-distillation queries and silently rerouting them (cyber → Opus 4.8, bio → Opus 5), firing in <5% of sessions. Mythos 5 (no classifiers) is restricted to **Project Glasswing** partners and limited US government channels.
- **Specs:** 1M context, 128K max output, January 2026 knowledge cutoff, adaptive thinking always on.
- **Pricing:** **$10/$50 per 1M** — 2× Opus 4.8. Cache reads $1/MTok.
- **Benchmarks (vendor):** SWE-bench Verified **~95%**, SWE-bench Pro **80.3%** (launch SOTA; +11pp over Opus 4.8's 69.2%, +21pp over GPT-5.5's 58.6%), GPQA Diamond **92.6%**, Terminal-Bench 2.1 **88.0%**, τ²-Bench **98.5%**, HLE with tools **64.5%**, AA Intelligence Index **64.9** (#1 public at launch).
- **Export-control incident:** pulled globally **June 12 – July 1, 2026** under a US export-control directive following a jailbreak report; restored behind stricter classifiers. Subscription bundling ended July 7 (moved to metered usage credits).
- Reversed a hidden anti-distillation guardrail within 24h of launch after community discovery (documented in the 319-page system card).

## 4.3 Fable 5.1 — released 1 September 2026 (with Mythos 5.1)
- Same underlying weights as Mythos 5.1; Fable = safeguarded public version. API ID `claude-fable-5-1`. Anthropic's fifth Claude 5-gen release in <4 months.
- **Pricing:** sticker unchanged **$10/$50**; **cache reads cut 75% → $0.25/MTok** — Anthropic estimates ~25% savings on typical workloads, up to ~45% on complex agentic pipelines.
- **Benchmarks (vendor, Sept 1 post):**
  | Benchmark | Fable 5.1 | Fable 5 | Opus 5 |
  |---|---|---|---|
  | Terminal-Bench-Science 0.1 | **52.6%** | 24.7% | 29.0% |
  | Terminal-Bench 4.0 | **55.8%** | 42.0% | 52.3% |
  | GDPval-AA v2 | **1853** | 1723 | 1824 |
  | OSWorld 2.0 | 77.9% / 41.7% strict | 72.9% / 36.1% | 75.4% / 39.6% |
  | HLE (with tools) | **65.0%** | 63.8% | 63.6% |
  | AutomationBench | **31.4%** | 17.1% | 26.9% |
  | CursorBench 3.2.0 | **73.4%** | 70.5% | 70.0% |
  | GPQA Diamond | **93.7%** | — | — |
  | FrontierMath Tier 4 v2 | 87.8% | — | — |
  | ARC-AGI-2 | **90%** at 32% lower cost/task than Fable 5 | 89.2% ($5.45/task) | — |
  | AA Intelligence Index (max) | **66** | — | — |
- Standout: science-agent leap (24.7% → 52.6%); robotic pick-and-place 5% → 40% (8×); #1 on BenchAlign (alignment/refusal benchmark).
- vs GPT-6 Astra (Sept 3, same $10/$50): Astra leads on FrontierMath (97.6%), AutomationBench (41.4%), ExploitBench (100% vs 70%), cost/task ($1.67 vs $3.76); Fable 5.1 leads on HLE+tools (65.0%), AA Index (66 vs 61).
- Other: reduced false-positive bio/cyber safeguards (~60% fewer cyber false positives in Claude Code); **invisible text watermarking** with detection API (private preview, EU-law eligible orgs); **Enterprise Frontier Safeguards (EFS)** rolling out fall 2026 (ZDR-like privacy with misuse detection; customer-controlled infra).

---

# 5. Claude Code — evolution through 2026

