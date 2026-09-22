---
id: open-local-models-2026/11-part-9-master-comparison-table-sept-2026-snapshot/overview
title: "PART 9 — MASTER COMPARISON TABLE (Sept 2026 snapshot)"
domain: part-9-master-comparison-table-sept-2026-snapshot
role: deep-dive
task: reference
actors: ["Google", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "Poolside", "United States"]
dates: ["2024-12-12", "2025-04-05", "2025-11-20", "2025-12-02", "2025-12-15", "2026-03-11", "2026-03-16", "2026-04-02", "2026-04-29", "2026-06-04", "2026-07-21", "2026-08-10", "2026-09-02"]
keywords: ["apache", "benchmarks", "license", "llama", "mistral", "muse", "muse spark", "nvidia", "scout"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [678, 703]
section: "PART 9 — MASTER COMPARISON TABLE (Sept 2026 snapshot)"
sha256: 8d4fe36a6fdb14e68757dfcb1b98f1d26945b7ef5f429abe639bb6e51271c4c9
---

# PART 9 — MASTER COMPARISON TABLE (Sept 2026 snapshot)

| Model | Vendor | Release | Params | Context | License | HF weights | AA Intelligence Index | SWE-bench Verified | API in/out per 1M |
|---|---|---|---|---|---|---|---|---|---|
| Llama 4 Maverick | Meta | 2025-04-05 | 400B / 17B act | 1M | Community License | ✅ | 9.3 | 21.0% | $0.20 / $0.80 |
| Llama 4 Scout | Meta | 2025-04-05 | 109B / 17B act | 10M | Community License | ✅ | 6.5 | — | $0.19 / $0.68 |
| Muse Spark 1.3 | Meta | 2026-09-02 | undisclosed | 1M | Proprietary (API-only) | ❌ | 61 (xhigh) / 62 (max) | — | $1.25 / $4.25 |
| Muse Glimmer 30B | Meta | 2026-08-10 | 30B dense | 120K+ | Apache 2.0 | ✅ | — | 76% (vendor) | n/a (local) |
| Laguna S 2.1 | Poolside | 2026-07-21 | 118B / 8B act | 1M | OpenMDW-1.1 | ✅ | — (coding) | 59.4% Pro (vendor) | $0.10 / $0.20 |
| Mistral Large 3 | Mistral | 2025-12-02 | 675B / 41B act | 256K | Apache 2.0 | ✅ | — | — | $0.50 / $1.50 |
| Mistral Medium 3.5 | Mistral | 2026-04-29 | 128B dense | 256K | Modified MIT | ✅ | — | 77.6% (vendor) | via La Plateforme |
| Mistral Small 4 | Mistral | 2026-03-16 | 119B / 6B act | 262K | Apache 2.0 | ✅ | — | — | via La Plateforme |
| Nemotron 3 Ultra | NVIDIA | 2026-06-04 | 550B / 55B act | 1M | OpenMDW-1.1 | ✅ | **48** (top US open) | 70.7 (vendor) | $0.50 / $2.50 |
| Nemotron 3 Super | NVIDIA | 2026-03-11 | 120B / 12B act | 1M | NVIDIA Open Model License | ✅ | 36 | — | via NIM |
| Nemotron 3 Nano | NVIDIA | 2025-12-15 | 31.6B / 3.2B act | 1M | NVIDIA Open Model License | ✅ | — | — | via NIM |
| Gemma 4 31B | Google | 2026-04-02 | 30.7B dense | 256K | Apache 2.0 | ✅ | 39 | ~54% (vals.ai subset) | Free (AI Studio) |
| Gemma 4 26B-A4B | Google | 2026-04-02 | 25.2B / 3.8B act | 256K | Apache 2.0 | ✅ | — | — | Free (AI Studio) |
| OLMo 3 Think 32B | Ai2 | 2025-11-20 | 32B | 65K | Apache 2.0 | ✅ | — | — | self-host |
| Granite 4.0 H 1B | IBM | 2025-10 | 1B | long | Apache 2.0 | ✅ | — | — | self-host |
| Phi-4 | Microsoft | 2024-12-12 | 14B dense | 16K | MIT | ✅ | — | — | self-host |
| gpt-oss-120b | OpenAI | 2025-08 | 120B | 128K | Apache 2.0 | ✅ | 33 | — | self-host |

*AA Intelligence Index values are from AA v4.x snapshots (methodology-version dependent); vendor benchmarks are marked where relevant.*

---

