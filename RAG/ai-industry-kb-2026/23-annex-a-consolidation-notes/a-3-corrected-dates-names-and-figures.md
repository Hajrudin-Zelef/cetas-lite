---
id: ai-industry-kb-2026/23-annex-a-consolidation-notes/a-3-corrected-dates-names-and-figures
title: "A.3 Corrected dates, names, and figures"
domain: appendix
role: appendix
task: reference
actors: ["Anthropic", "Cerebras", "DeepSeek", "Moonshot", "OpenAI", "Samsung", "SpaceX", "Unsloth", "xAI"]
dates: ["2026-03-23", "2026-04", "2026-05", "2026-06-08", "2026-06-16", "2026-09-14"]
keywords: ["apache", "deepseek", "disclosure", "funding", "grok", "grok 4", "ipo", "kimi", "license", "mxfp4", "parameters", "quantization"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [11065, 11087]
section: "Annex A — Consolidation notes"
sha256: b1cd3b667230d18af965acf39a6f73c7412e8bb1703e75c6e8876069fbf4c220
---

# A.3 Corrected dates, names, and figures

1. **DeepSeek-V4 license: MIT vs Apache 2.0** — §9. Later sources report MIT; April 2026 coverage reports Apache 2.0; the first-party card was not directly opened. Settled by: direct reading of the official DeepSeek V4 model card. (Shipped MIT is the current working record; rumor-phase Apache 2.0 claims are residue, but the first-party confirmation was never obtained.)
2. **V4-Pro → V4.1-Flash routing notice (Sept 10) vs APIMaster revision** — §5. The Sept-10 notice says V4-Pro routes to V4.1-Flash from Sept 14; APIMaster's reported revision says V4 Pro service continues unchanged. Flagged conflict, not resolved arbitrarily. Settled by: provider routing-behavior verification after 2026-09-14, or a second vendor notice.
3. **Kimi K3 active params: 104B (technical-report consensus) vs ~50B (single Medium analysis)** — §9. Working value 104B with the dispute noted (do-not-cite-without-checking). Settled by: Moonshot publishing the exact active-parameter accounting from the primary technical report.
4. **Kimi K3 per-expert geometry: 33.0M/expert vs 2.72T routed total** — §3. Architecture inspections report w1/w3 [3072, 3584], w2 [3584, 3072] = 11.0M parameters each (~33.0M per expert), while also citing 2.72T routed parameters total — but 896 × 33.0M ≈ 29.6B, not 2.72T (2.72T / 896 ≈ 3.0B per expert). The 33.0M figure is internally inconsistent and treated as a reconstruction artifact; shapes and MXFP4 E8M0 quantization are the reliable part. Settled by: primary-source correction of the per-expert accounting. (Active-count cross-check, 104B vs naive 16/896 × 2.8T ≈ 50B routed-active, is likewise undecomposed — do not present a decomposition as fact.)
5. **Unsloth funding: ~$40K/YC/Kilpatrick (unsourced) vs ~$500K (Redpoint scout, Samsung NEXT, 2026-03-23 scorecard)** — §2, §10. The $40K figure has no source; the only concrete third-party figure (~$500K) contradicts it. Both stay [UNVERIFIED] and must not be presented as facts. Settled by: a primary disclosure or a verifiable filing. (The AI Engineer World's Fair 2026 presence, June 30 – July 2, and the "112 slides" deck are likewise [UNVERIFIED].)
6. **Cursor $60B acquirer identity: xAI vs SpaceX vs "SpaceXAI"** — §13, §19. Wave 1 flagged xAI; June 16, 2026 reports say SpaceX (expected close Q3 2026) [UNVERIFIED]; September reports use "SpaceXAI" for Grok 4.7 — press-level naming drift. Settled by: a confirmed deal announcement naming the acquirer. (§20 separately records "xAI absorbed by SpaceX (Feb 2, 2026)" as [UNVERIFIED detail] pending source confirmation — do not merge the two without further sourcing.)
7. **Robotics application splits: SAG (70%+ industrial) vs Counterpoint (>60% entertainment/research)** — §22. Reported as incompatible, not blended. Settled by: a third independent tracker with a disclosed, comparable methodology.
8. **Cerebras IPO proceeds: $5.55B vs up-to-$3.5B estimates** — §20. Intentionally unresolved per wave-source verdicts. Settled by: the final IPO prospectus / post-IPO disclosure reconciliation.
9. **Anthropic May 2026 amount/valuation** — §20. Intentionally unresolved per wave-source verdicts. Settled by: a confirmed filing or company statement. (Anthropic fundraising figures were not present in the source files and were not invented.)
10. **Etched round lettering and cumulative totals** — §20. Intentionally unresolved per wave-source verdicts. Settled by: company-confirmed round data.
11. **xAI SPV/debt/equity split** — §20. Intentionally unresolved per wave-source verdicts. Settled by: confirmed deal documentation.
12. **OpenAI tender per-person cap: $50M vs $30M** — §20. Intentionally unresolved per wave-source verdicts. Settled by: tender-offer documentation.
13. **Helsing Series E lead: Dragoneer vs Goldman Sachs** — §20. Intentionally unresolved per wave-source verdicts. Settled by: confirmed round reporting.
14. **OpenAI confidential filing date: June 8, 2026 [VENDOR] vs May 22 [UNVERIFIED single digest]** — §21. Working record uses June 8 everywhere, but the discrepancy is flagged, not adjudicated. Settled by: EDGAR or company-announced confirmation of the filing date.

---

## A.3 Corrected dates, names, and figures

Deduped across fragments; each row lists the corrected value with all § references where the correction was applied or flagged.

### Dates

