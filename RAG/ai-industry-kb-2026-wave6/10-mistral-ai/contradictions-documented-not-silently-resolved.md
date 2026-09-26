---
id: ai-industry-kb-2026-wave6/10-mistral-ai/contradictions-documented-not-silently-resolved
title: "Contradictions documented (not silently resolved)"
domain: mistral-ai
role: deep-dive
task: reference
actors: ["EU", "Mistral", "Nvidia", "OpenAI", "United States"]
dates: ["2025-12", "2025-12-02", "2026-04-28", "2026-05", "2026-09", "2026-09-21", "2026-09-22"]
keywords: ["agents", "apache", "blackwell", "chatgpt", "compute", "gpus", "hyperscaler", "license", "mistral", "multimodal", "pricing", "research"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5030, 5047]
section: "§10. Mistral AI"
delta_of: ai-industry-kb-2026
sha256: a0ae9a6b74da176fe20aab63451b062d3c674d5a7339c3eb0225b0ded048c240
---

# Contradictions documented (not silently resolved)

- Mistral's December 2025 product offensive (Large 3, Ministral 3, Devstral 2, OCR 3) was a coordinated open-weights push timed against American rivals' capital raises — and the September 2026 Series D (€3B at >€21B post-money, ~1.8× step-up from Series C's €11.7B) retroactively funded that strategy. [SECONDARY, S32][SECONDARY, S13][SECONDARY, S55][SECONDARY, S56]
- The license map now reads as a deliberate portfolio: Apache 2.0 for edge/small/open models (Ministral 3, Small 4, Large 3, Devstral, Magistral Small, Robostral), Modified MIT for the dense flagship (Medium 3.5), proprietary for enterprise APIs (Medium 3.1, Magistral Medium, OCR). The expansion's per-model license fields preserve this structure. [SECONDARY, S31][SECONDARY, S5]
- The Ministral 3 correction matters beyond dating: a December 2025 Apache 2.0 edge family with vision at 3B–14B reframes Mistral's edge story as a year old, not a 2026 development — the existing §10's "unverified 2026" framing understated it. [SECONDARY, S31][SECONDARY, S13]
- OCR 4's single-container air-gapped deployment plus EU incorporation is Mistral's sharpest sovereignty differentiator: unlike US providers offering EU data residency under US law, documents never leave the customer's infrastructure. The techtimes analysis explicitly ties this to GDPR and financial-services procurement. [SECONDARY, S34][SECONDARY, S35]
- Mistral Compute (18,000 Blackwell GPUs, 44 MW near Paris) marks Mistral's entry into the infrastructure layer — a necessary move if the €1B 2026 revenue target is to be served on sovereign capacity rather than rented hyperscaler GPUs. [SECONDARY, S13][SECONDARY, S33]
- The Vibe rebrand (May 2026) folds chat, work automation, and coding agents into one license — the same bundle logic as OpenAI's ChatGPT Work — but Mistral's self-host story is the wedge for European enterprise buyers wary of US provider dependency. [SECONDARY, S35][SECONDARY, S17]
- The Codestral pricing discrepancy ($0.30/$0.90 at 256K vs $0.33/$0.99 at 131K) likely reflects deployment/version differences between provider listings rather than a price change; neither listing is dated as a rate revision. [SECONDARY, S19][SECONDARY, S18]

### Contradictions documented (not silently resolved)
- Medium 3.5 release date: official changelog 2026-04-28 vs existing §10's Apr 29/30 — the expansion adopts 2026-04-28 as authoritative. [VENDOR, S4]
- Small 4 context: 256K (stronger sources, current catalog) vs 128K (some mirrors) — 256K carried as primary, 128K noted as mirror variance. [SECONDARY, S31][SECONDARY, S9]
- Large 3 modality: multimodal per multiple sources vs text-only in one comparison table — multimodal carried. [SECONDARY, S31][SECONDARY, S7]
- Large 3 license: Apache 2.0 (multiple newer sources) vs vague "commercial terms" (one table) — Apache 2.0 carried. [SECONDARY, S31][SECONDARY, S7]
- Ministral 3: existing §10 "Ministral 3 (2026), unverified single-source" vs new research "Ministral 3 family, 2025-12-02, 3B/8B/14B, Apache 2.0" — the expansion corrects the record with two corroborating sources. [SECONDARY, S31][SECONDARY, S13]
- Medium 3.5 SWE-Bench Verified 77.6: single-sourced, no corroboration found — carried as UNVERIFIED, not as fact. [UNVERIFIED, S5]
- Modified MIT (resolved 2026-09-22): verbatim carve-out extracted from the LICENSE file — $20M global consolidated monthly revenue threshold, above which rights cannot be exercised and a commercial license from Mistral is required. [SECONDARY, S36][SECONDARY, S41]
- Codestral context/pricing: 131K at $0.33/$0.99 (updated 2026-09-21) vs 256K at $0.30/$0.90 — likely version/deployment differences; both preserved. [SECONDARY, S18][SECONDARY, S19]

