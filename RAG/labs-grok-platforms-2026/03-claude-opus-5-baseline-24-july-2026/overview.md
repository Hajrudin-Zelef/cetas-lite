---
id: labs-grok-platforms-2026/03-claude-opus-5-baseline-24-july-2026/overview
title: "2. Claude Opus 5 — baseline (24 July 2026)"
domain: claude-opus-5-baseline-24-july-2026
role: deep-dive
task: reference
actors: ["Anthropic", "OpenAI"]
dates: ["2026-05", "2026-07", "2026-07-24"]
keywords: ["claude", "opus 5", "agi", "fable 5", "gpt-5.6", "mythos 5", "opus 4", "sol"]
source: docs/RAG/Labos, Grok, outils & plateformes_EN.md
source_anchor: ""
source_lines: [75, 86]
section: "2. Claude Opus 5 — baseline (24 July 2026)"
sha256: 08dad640374036aabfbc4a7449876c6f887e8f3b4449475cdf3f6e7010fade5d
---

# 2. Claude Opus 5 — baseline (24 July 2026)

Context needed to interpret Opus 5.5 deltas:
- **Released July 24, 2026** at **$5/$25 per 1M** — identical to Opus 4.8 ("the price line did not move but the capability line did").
- **1M-token context, 128K max output, May 2026 knowledge cutoff** (most current Claude at the time). API ID `claude-opus-5`. Default on Claude Max; strongest on Claude Pro.
- **Adaptive thinking on by default** (new vs 4.8); effort tiers low/medium/high/xhigh/**max**; Fast mode ~2.5× speed at $10/$50.
- Key vendor figures: **SWE-bench Verified 96.0%** (highest publicly verified at the time), Frontier-Bench v0.1 **43.3%** (2× Opus 4.8's 21.1%), ARC-AGI-3 **30.2%** (vs 1.5% for 4.8), OSWorld 2.0 **70.6%**, AutomationBench 26.0%, GDPval-AA v2 Elo **1,861**.
- Independent: **BenchLM #1 across 215 models (85.88)** — above Mythos 5 (83.01), Fable 5 (82.76), GPT-5.6 Sol (81.46); **AA Intelligence Index 61**, one point above Fable 5.
- vs Fable 5: effectively tied/leading on repo-level coding (SWE-bench Pro 79.2% vs 80.3% — Fable's only coding win), ahead on Frontier-Bench (+9.6pp), at half the price.

---

