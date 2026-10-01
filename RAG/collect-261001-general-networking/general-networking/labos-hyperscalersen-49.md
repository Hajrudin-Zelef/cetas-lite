---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-49
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "Cerebras", "Cohere", "Google", "Hugging Face", "Meta", "Microsoft", "Mistral", "Nscale", "Nvidia", "OpenAI", "Sakana", "SpaceX", "United States", "xAI"]
dates: ["2026-05", "2026-05-28", "2026-08", "2026-09-21"]
keywords: ["acquisition", "agent", "agents", "apache", "astra", "aws", "benchmark", "capex", "cohere", "compute", "cyber", "fugu"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [2515, 2616]
sha256: 819f18b5ff52aa2a1ce4baa9b23b61fd5da0dd1eca6ba0bf53d28c675d460144
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

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

### 16.13 Capability frontier map — who leads what (Sep 2026, as reported)

> "Leads" = best vendor-reported or independent score on the most-cited benchmark in that lane. All vendor-reported unless noted.

| Lane | Leader (Sep 2026) | Evidence |
|---|---|---|
| Raw coding (SWE-style) | OpenAI GPT-5.6 Sol | T-Bench 2.1 88.8% (vendor) |
| Long-horizon agents | xAI Grok 4.6 → 4.7 | GDPVal-AA v2 Elo 1753; CursorBench gains (vendor) |
| Computer use | Google Gemini 3.6 Flash | OSWorld-Verified 83.0% (official table) |
| Cheap frontier intelligence | Google Gemini 3.6 Flash / xAI Grok 4.7 | $1.50/$7.50 and $2/$6 per 1M |
| Open weights (overall) | Thinking Machines Inkling | 975B/41B MoE, Apache 2.0 |
| Open weights (enterprise) | Cohere Command A+ | 218B/25B MoE, Apache 2.0, 48 languages |
| Open weights (small) | Mistral Small 4 | 119B MoE, $0.15/$0.60 API |
| Reasoning (math) | Microsoft MAI-Thinking-1 | AIME 2025 97.0% (vendor) |
| Speech-to-text | Mistral Voxtral Transcribe 2 | 13 langs, sub-200ms Realtime |
| Text-to-speech | Mistral Voxtral TTS / xAI Custom Voices | 9 langs open / 28 langs commercial |
| Image generation | MAI-Image-2.5 (claimed) / Nano Banana lineage | Arena #3 text-to-image (vendor-reported) |
| Life sciences | OpenAI GPT-Rosalind | BixBench 0.751 (vendor) |
| Multi-agent orchestration | Sakana Fugu Max / Ultra v2 | Self-reported Pareto claims |
| Document understanding | Mistral OCR 4 | 170 languages, bounding boxes |
| Formal proof (Lean 4) | Mistral Leanstral | FLTEval pass@16 31.9 (vendor) |
| Cyber offense (eval) | OpenAI GPT-6 Astra | First "Critical" classification (vendor) |

**Read:** no single lab leads everywhere — the frontier is a patchwork of lane leaders. OpenAI holds coding/cyber, Google holds cheap intelligence and computer use, xAI holds long-horizon agents, Mistral holds speech/document, and the open-weights crown is split three ways. [research finding]

### 16.14 Reading Artificial Analysis scores correctly — worked example

The single most common error in 2026 AI reporting was comparing Artificial Analysis Intelligence Index scores across index revisions. Worked example from this file:

- **Grok 4.6: 61** — measured August 2026, on the index revision then current.
- **Grok 4.7: 46** — measured September 21, 2026, on **index v4.3.2** (a different revision).
- **Naive reading:** "Grok got worse" (61 → 46). **Correct reading:** the two numbers are not comparable; the task mix, harness, and scoring changed between revisions. Same-day AA data actually showed Grok 4.7 *improving* on 4.6 (Coding Agent Index 56, +9; AA-Briefcase Elo 1,657, +111).

**Rules applied throughout this file:**
1. Every AA score carries its measurement date/revision context.
2. Cross-revision comparisons are explicitly flagged, never implied.
3. Same-revision, same-harness tables (e.g., §16.5 benchmark clusters) are the only valid comparison surface.
4. Vendor-reported numbers are never mixed into independent-measurement tables without a provenance break.

This discipline is why §16.2 and §16.5 exist as separate tables: §16.2 shows dated snapshots *with* revision warnings; §16.5 shows same-harness clusters where comparison is legitimate. Any downstream RAG use should preserve that separation — do not let a retriever flatten the two tables into one ranking.

## §17 — M&A TRACKER (Feb–Sep 2026)

> Status taxonomy: **Closed** · **Signed / pending close** · **Announced** · **Reported / unconfirmed** · **Disputed**.

### Closed
| Date | Acquirer → Target | Terms | Note |
|---|---|---|---|
| 2026 | xAI ↔ SpaceX merger | — | Treated as closed per Track C; "SpaceXAI" appears in 18 Sep 2026 filings |
| 2026 | Cerebras IPO | public listing | Landmark AI-silicon liquidity event |
| May 2026 | OpenAI Deployment Co. → Tomoro | — | ~150 forward-deployed engineers in-house |
| 11 Jun 2026 | OpenAI → Ona (ex-Gitpod) | undisclosed; pending regulatory at announcement | 2M developers served; team joins Codex group |

### Signed / pending close
| Date | Parties | Terms | Expected close |
|---|---|---|---|
| 3 Sep 2026 | **NVIDIA → Hugging Face** (definitive agreement) | — | H1 2027 (regulatory approvals pending) |
| 16 Sep 2026 | **Cohere × Aleph Alpha** (definitive merger agreement) | — | pending regulatory approval |

### Announced
| Date | Parties | Note |
|---|---|---|
| 2026 | Mistral AI Series D (€3B) | financing, not acquisition — tracked in §16.1 |

### Reported / unconfirmed
| Date reported | Claim | Status |
|---|---|---|
| 14 Sep 2026 | **OpenAI acquired Glass Imaging (>$300M)** — WSJ | **Neither company confirmed** [unverified] |
| 2026 | **OpenAI acquired Anysphere (Cursor)** | **Disputed/unverified** — no confirmed close as of 22 Sep 2026 |
| 2026 | **OpenAI acquired TBPN (talk show)** + an unnamed open-source dev-tools startup | Names/terms not returned — gap |
| 2026 | **$100B AWS–OpenAI cloud deal** tied to Amazon investment | [unverified] — vs. $38B/7-yr Project Rainier reporting |

### Disputed
| Claim | Status |
|---|---|
| Cursor/Anysphere acquisition by OpenAI | Reports circulated; **treat as [unverified]** — Cursor operating independently in-window |
| Glass Imaging acquisition | WSJ-reported 14 Sep 2026; unconfirmed by both parties |

---
### M&A deal records — full narratives (Feb–Sep 2026)

