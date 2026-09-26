---
id: ai-industry-kb-2026-wave6/04-glm-and-z-ai/glm-5v-turbo-closed-commercial-counterpart
title: "GLM-5V-Turbo — closed commercial counterpart"
domain: glm-and-z-ai
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Huawei", "Hugging Face", "StepFun", "United States", "Z.ai"]
dates: ["2026-01-08", "2026-02-14", "2026-04-01", "2026-04-07", "2026-06", "2026-06-13", "2026-07-02", "2026-08", "2026-08-14", "2026-08-18", "2026-08-26", "2026-09", "2026-09-01", "2026-09-07", "2026-09-20"]
keywords: ["glm", "agent", "ascend", "attention", "benchmark", "benchmarks", "claude", "context window", "fable 5", "fine-tuning", "ipo", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1792, 1857]
section: "§4. GLM and Z.ai"
delta_of: ai-industry-kb-2026
sha256: f4c15519bd6a53af6b23b980438b30b2ee976313335372a7754329b3eebf57c7
---

# GLM-5V-Turbo — closed commercial counterpart

### GLM-5V-Turbo — closed commercial counterpart
- GLM-5V-Turbo (April 1, 2026) is **closed-source and API-only**: no downloadable weights, no MIT license, no self-hosting or fine-tuning [SECONDARY] (techtimes.com, 2026-07-02).
- Vision handled by the **CogViT** vision encoder for native image/video/document-layout processing [SECONDARY] (techtimes.com; MarkTechPost 2026-04-01; github.com/kyegomez/cogvit reference implementation).
- Reported spec: **200K context window, up to 128K max output**; **$1.20/M input / $4.00/M output** tokens [SECONDARY] (cometapi.com review, 2026).
- The open/closed split is deliberate product architecture: Z.ai kept vision capability in a closed commercial product while the open-weight flagship stayed text-only — a point of friction with the self-hosting community that adopted GLM-5.2 for its MIT license [SECONDARY] (techtimes.com, 2026-07-02; emergent.sh GLM-5.3 guide notes the same community demand for vision in an open checkpoint).

### Access routes and pricing-table mechanics (August 2026)
- GLM-5.3 per-token API access opened **August 18, 2026** at **$1.40 input / $4.40 output / $0.26 cached input per 1M tokens — identical to the GLM-5.2 rate** [SECONDARY] (emergent.sh, 2026-08).
- As of August 2026, Z.ai's **public pricing page still listed GLM-5.2 as its top row**; the confirmed 5.3 rate came through the API and gateways rather than a dedicated pricing-table line — re-verify before relying on it [SECONDARY] (emergent.sh, 2026-08).
- GLM Coding Plan: subscription for use inside Claude Code and Z.ai's **ZCode** development environment; **from $18/month on a points-based quota with an off-peak discount** [SECONDARY] (emergent.sh, 2026-08).
- deeplearning.ai/The Batch describes the Coding Plan range as **"$18 to $168 per month"**, which differs from the $18/$72/$160 tier figures in other sources — record as a **pricing-source conflict**, not a price change [SECONDARY] (deeplearning.ai vs llm-coding-benchmark pricing doc and pilot-shell blog).

### Z.ai / Zhipu capital-market facts
- Hong Kong IPO: **January 8, 2026**, priced at **HK$116.20/share**; ~**$558M raised** at ~**HK$51B / $6.5–6.7B** IPO valuation [SECONDARY] (caproasia.com, 2026-02-14; existing §4).
- First-day close **HK$131.50**, implying roughly **$7.4B** market value [SECONDARY] (caproasia.com, 2026-02-14).
- By **February 14, 2026**, shares were **+317.3% from the IPO price**; the company was founded in **2019**; February reporting also noted **Shanghai IPO plans** after the Hong Kong listing [SECONDARY] (caproasia.com, 2026-02-14).
- By **late June 2026** shares had risen **more than 2,000% from the January listing** as GLM-5.2 drew attention for approaching leading US models in coding and agent tasks [SECONDARY] (kr-asia.com, 2026-08).
- On **August 14, 2026** (GLM-5.3 launch day) the Hong Kong-listed shares **closed down 3.6%** — no launch rally, contrasting with the GLM-5.2 summer enthusiasm [SECONDARY] (kr-asia.com, 2026-08).
- Late-August 2026 market snapshot: ticker **2513**; market cap roughly **HK$500B (~$64B)** against about **HK$3.2B of consensus 2026 revenue and no earnings**; ainvest frames the valuation as resting on the claim that Z.ai leads the open-weights frontier [SECONDARY] (ainvest.com, 2026-08 — single source for the exact figures, mark [UNVERIFIED] for financial use).
- The company bills itself as the **world's first listed foundation-model company** — company self-description, not an independent classification [SECONDARY] (ainvest.com, 2026-08).

### Competitive board snapshot (Artificial Analysis, September 2026)
- On the independent board, **Step 5 Preview** (StepFun, announced September 20, 2026) and **GLM-5.3** sat one point apart: **44 vs 45 on AA Intelligence Index v4.3.2**, GLM-5.3 scored as "max" effort and Step 5 Preview carrying no effort label [SECONDARY] (orcarouter.ai, 2026-09-20).
- StepFun's Hugging Face repository contained only a **.gitattributes** file at announcement; the BF16 checkpoint was slated for **October 15** — described as a "top-three open-source result" but not yet open source at announcement, unlike GLM-5.3's downloadable weights [SECONDARY] (orcarouter.ai, 2026-09-20).
- Do not mix this AA **v4.3.2** reading with the v4.3 readings elsewhere in §4 [DIRECTIONAL].

## Figures and metrics

| Model | Date | Size (total/active) | Context | License | Provenance |
|---|---|---|---|---|---|
| GLM-5V-Turbo | 2026-04-01 | 744B / ~40B | 200K | closed / API-only | [VENDOR] |
| GLM-5.1 | 2026-04-07 | 744B / 40B | — | MIT | [VENDOR] |
| GLM-5.2 | 2026-06-13 (API); 06-16/17 (weights) | 744B-class | 1M | MIT | [VENDOR] |
| GLM-5.3-Flash | 2026 (stealth "Ox Alpha" → Flash) | same base as 5.2 | 1M | MIT | [SECONDARY] |
| GLM-5.3 | 2026 | identical base to 5.2 | 1M | GLM-5.3 License (bespoke) | [VENDOR] |

| Benchmark/price | Figure | Date | Provenance |
|---|---|---|---|
| GLM-5.3, AA Intelligence Index | ≈44.9 | v4.3, 2026-09-07 | [SECONDARY] |
| GLM-5.3-Flash, AA Intelligence Index | 42 | v4.3, 2026-09-07 | [SECONDARY] |
| GLM-5.3 API | $1.40/$4.40 per M in/out, 81% cache discount | 2026-08-18 | [VENDOR] |
| GLM-5.3-Flash API | $0.15/M input | 2026-08-26 | [VENDOR] |
| GLM Coding Plan | from $18/mo | 2026-06 | [VENDOR] |
| GLM-5V-Turbo benchmarks | vendor-only (no independent replication) | 2026-04 | [VENDOR] |
| Terminal-Bench 4.0, GLM-5.3 | 41.8% | official, 2026-09-01/02 | [SECONDARY] |
| TB 4.0 official leader (Fable 5.1) | 57.9%±3.8 | official, 2026-09-01/02 | [SECONDARY] |
| TB 4.0 vendor-reported (Mythos 5.1) | 60.9% | Anthropic self-report | [VENDOR] |

**No cross-version comparison:** GLM-5.3's 44.9 is v4.3-pinned; do not compare against v4.1.1/v4.2-era scores. TB 4.0 (41.8%) is not comparable with TB 2.1-era scores. [DIRECTIONAL]


### New verified metrics — expansion (continued — Z.ai capital figures)

- IPO: HK$4.35B ($558M); HK$116.2 offer → HK$131.5 close (+13.1%); ~$7.4B day-1 value [SECONDARY].
- Peak: HK$2,980; >HK$1T (~$128B) cap; +2,000% since IPO [SECONDARY].
- Follow-on: ~$4B via 19.8M shares @ HK$1,588–1,698; CICC sole coordinator [SECONDARY].
- 2024 revenue: 312.4M yuan (~$46M) [SECONDARY].


### New verified metrics — expansion (continued — GLM-5V-Turbo figures)

- 744B-A40B (753B on HF cards); CogViT; 200K ctx / 128K out; 30+ RL tasks; Design2Code 94.8; WebVoyager/AndroidWorld #1; SpeedBench 221.2 tok/s (#5); $1.20/$4.00; 100K Ascend 910B [VENDOR via secondary / SECONDARY].


### New verified metrics — expansion (continued — GLM-5.2 benchmark tables)

