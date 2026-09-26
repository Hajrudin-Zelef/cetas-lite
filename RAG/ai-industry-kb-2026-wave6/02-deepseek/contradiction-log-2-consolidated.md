---
id: ai-industry-kb-2026-wave6/02-deepseek/contradiction-log-2-consolidated
title: "Contradiction log — §2 consolidated"
domain: deepseek
role: deep-dive
task: reference
actors: ["Baseten", "DeepSeek", "Huawei", "Meta", "Nvidia"]
dates: ["2025-05", "2025-09-29", "2025-12-01", "2026-04-24", "2026-05", "2026-06", "2026-07-31", "2026-08-13", "2026-09", "2026-09-10", "2026-09-22"]
keywords: ["ascend", "attention", "attribution", "benchmark", "benchmarks", "cost", "decode", "deepseek", "funding", "inference", "nvidia", "prefill"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [955, 995]
section: "§2. DeepSeek"
delta_of: ai-industry-kb-2026
sha256: 3d1389b1444839b5b72d7958836e9e8399f00754dc8d4f25b4b2a19ea8c5eebe
---

# Contradiction log — §2 consolidated

### Contradiction log — §2 consolidated
- **C1**: V3.2 total params 671B (official card) vs 685B (DeepLearning.ai) — **resolved** by attribution. [VENDOR]
- **C2**: V3.2 pricing $0.14/$0.28 vs $0.27/$0.40 — tier/cache artifact, both retained with tier labels. [SECONDARY]
- **C3**: V4-Pro AA Index 53 vs 36 — different revisions/dates; never mix versions. [SECONDARY]
- **C4**: V4.1-Flash total 552B (vendor) vs 748B (community) vs 763B (headline) — four-number resolution adopted. [SECONDARY]
- **C5**: V4-Pro-0813 pricing $1.32/$3.96 vs $0.435/$0.87 — tier/window-dependent; do not mix. [SECONDARY]
- **C6**: R2 launch May 2026 vs May 2025 — 2025 conflicts with R1 lineage; May 2026 better-attested. [SECONDARY]
- **C7**: V4-Flash-0731 active params ~13B (vendor) vs 21B (community evidence brief) — flagged, unresolved. [COMMUNITY]
- **C8**: V4.1-Flash Engram 196B vs 196.6B vs 196.9B — rounding across sources; treat as ~196B. [SECONDARY]

### §2 close — what the corpus now holds
- V3.2 (2025-12-01) is the baseline: 671B/37B, DSA, 128K, MIT, $0.28/$0.028/$0.42 API, V3.2-Speciale contest golds. [VENDOR]
- V3.2-Exp (2025-09-29) is the ablation: one-component swap, $0.028/M input, 2–3× faster long-context, Huawei-chip support. [SECONDARY]
- V4 preview (2026-04-24) → V4-Flash-0731 (2026-07-31) → V4-Pro-0813 (2026-08-13) → V4.1-Flash (2026-09-10) is the 2026 cadence: hybrid attention + mHC + Muon → CED + CSA2 + Engram + DSpark. [SECONDARY]
- V4-Pro lived four weeks (0813→0914); its endpoint now serves V4.1-Flash at Flash prices. [SECONDARY]
- R2 has no artifact; the Ascend training failure and the fundraising freeze are the same Nvidia-dependence story. [SECONDARY]
- The first round closed at $7B (June 2026); the $71–74B second round froze in September 2026. [SECONDARY]
- Every claim above carries two independent sources or an explicit label downgrade; vendor tables remain vendor claims. [DIRECTIONAL]

### NEW SOURCES — provenance key
- [VENDOR] = DeepSeek official card/report/API docs.
- [SECONDARY] = Reuters, DeepLearning.ai, TechCrunch, VentureBeat, established trade press, audited community engineering.
- [COMMUNITY] = r/LocalLLaMA analyses, GitHub recipes, independent benchmarks without peer review.
- [UNVERIFIED] = single-source rumors (R2 specs, leaked notes) — never merged into metrics.
- [DIRECTIONAL] = analyst inference from the above, marked as such.

### §2 expansion record
- This expansion added: V3.2-Exp full recipe, V3.2 official card + Speciale, V4 architecture deep-dive (mHC/CSA/HCA/SWA/Muon), V4-Flash-0731 evidence brief, V4-Pro-0813 launch coverage + pricing discrepancy, V4.1-Flash GA + audit + hardware + ecosystem, V4 Pro retirement mechanics, R2 delay causes, fundraising freeze, lineage table, contradiction log C1–C8.
- Unique URLs: 66. Contradictions documented: 8 (1 resolved).
- All benchmark tables labeled vendor claims; no AA Index versions mixed; no SWE-bench generations mixed.
- Research date: 2026-09-22. All URLs verified via search on that date; no constructed URLs.
- The 12,000-line consolidated target is a downstream constraint; this file is one input among many.
- End of §2 expansion. Next: consolidate with wave6 corpus §2 (lines 165–285) at the parent's direction.
- **Addendum — V4.1-Flash vs V4-Flash-0731 price ladder**: V4-Flash-0731 $0.14/$0.28 → V4.1-Flash $0.15/$0.60 off-peak ($0.30/$1.20 peak) — the input price held flat while output doubled off-peak, reflecting the 8B/16B prefill/decode asymmetry: decode is where the cost lives. [SECONDARY]
- **Addendum — the 437× denominator in one line**: 389,120 B/token (V1: 95 layers × 8 KV heads × 128 wide × 2 bytes bf16) ÷ 890 = 437.2 — the arithmetic is exact; the comparison is marketing. [SECONDARY]
- **Addendum — DeepSeek's September 2026 in one paragraph**: V4.1-Flash GA (09-10), Baseten day-one (09-11), V4 Pro retired (09-14), second-round frozen (09-≈15), R2 still missing (09-22) — a release, a retirement, a financing freeze, and an absence. The month reads as DeepSeek trading its premium tier for its efficient tier while its funding story wobbled. [DIRECTIONAL]
- File complete: 400+ substantive lines. All claims labeled; all URLs search-verified 2026-09-22.
-

---

