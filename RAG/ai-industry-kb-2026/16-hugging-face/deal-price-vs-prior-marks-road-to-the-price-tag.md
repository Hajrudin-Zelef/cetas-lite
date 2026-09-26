---
id: ai-industry-kb-2026/16-hugging-face/deal-price-vs-prior-marks-road-to-the-price-tag
title: "Deal price vs prior marks (road to the price tag)"
domain: hugging-face
role: deep-dive
task: funding-deals
actors: ["Alibaba", "DeepSeek", "Moonshot", "Nvidia", "OpenRouter", "Z.ai"]
dates: ["2023-08", "2026-05", "2026-06", "2026-08", "2026-09-02"]
keywords: ["acquisition", "arr", "deepseek", "glm", "inference", "kimi", "nvidia", "open weights", "open-weight", "revenue", "valuation"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [8012, 8038]
section: "16. Hugging Face"
sha256: f6e266f5ad8773bfce04455292d370665bddde3307750217669a71663fc4c6f8
---

# Deal price vs prior marks (road to the price tag)

- The second million models landed in roughly **a third of the time** the first took — growth is accelerating, with a new repository every ~7 seconds [VENDOR, Delangue via fintech-radar].
- Epoch's curated snapshot tracks only **1,340 open-weight models** (964 of them language models); ATOM tracks ~1,500 mainline open language models; **95 of 170** models on the Artificial Analysis Intelligence Index are open-weight — against 2.4M–3M total Hub models. The raw count includes derivatives, conversions, quantizations, and fine-tunes: **model selection, not model count, is the hard problem** (Wave 1).
- Release-day behavior is part of the metric: every major 2026 open-weight release landed on HF within days (Kimi K3, Qwen3.8-Max, DeepSeek V4 family, GLM-5.3-Flash, gpt-oss) — the Hub is where open weights are *discovered*, which is what NVIDIA is buying.

### Deal price vs prior marks (road to the price tag)

| Mark | Valuation | Multiple of |
|---|---|---|
| August 2023 venture round ($235M, Salesforce Ventures-led) | $4.5B | 1.0× |
| Late 2025 rejected NVIDIA investment offer ($500M) | $7B | ~1.6× |
| Agreed acquisition price (2026-09-02/03) | $12.9303B | ~2.9× |

Revenue context: $81M ARR end of 2025 → ~$100M (June 2026) → ~$150M annualized (August 2026), all
Sacra estimates [DIRECTIONAL, unaudited]. $12.93B is also ~10% of one year's NVIDIA free cash flow
(~$127B TTM per ainvest) [DIRECTIONAL] — "pocket change" financially; the price buys leverage over
the AI software stack, not cash flows.

### Free-tier inference ecosystem (2026 reference)

The lowest-friction way to test any Hub model remains "click Inference API on any model page, get a
curl command, testing in 30 seconds" (May 2026 practitioner reference). Free-tier examples:
Cloudflare Workers AI (~80 free models including DeepSeek-R1 distills), Zhipu/Z.AI Flash tiers,
LLM7.io, Kluster AI. Competitive context: Vercel AI Gateway and OpenRouter play the same single-key
routing game; **HF's edge is zero-config coupling to the Hub** — every model page is one click from
a routed API call.
| Companion finding (footnote) | **CVE-2026-0599** — TGI v3.3.6, unauthenticated DoS via unbounded image fetching during VLM input validation (High severity; SentinelOne analysis; upgrade + reverse-proxy request limits recommended); documented by sable.somoswilab; related to the HF security picture but distinct from the LeRobot flaw |

