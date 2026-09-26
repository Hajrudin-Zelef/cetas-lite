---
id: ai-industry-kb-2026-wave6/05-kimi-and-moonshot-ai/kimi-k3-license-conflict-recorded-unresolved
title: "Kimi K3 — license conflict (recorded, unresolved)"
domain: kimi-and-moonshot-ai
role: deep-dive
task: licenses
actors: ["Anthropic", "Baseten", "China", "DeepSeek", "Fireworks AI", "LongCat", "Meituan", "Moonshot", "Nebius", "OpenAI", "OpenRouter", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-01", "2026-01-27", "2026-03", "2026-04", "2026-04-20", "2026-05", "2026-06-12", "2026-07", "2026-07-16", "2026-09-07"]
keywords: ["kimi", "license", "arr", "attention", "benchmark", "benchmarks", "claude", "deepseek", "distillation", "funding", "glm", "mit license"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2313, 2399]
section: "§5. Kimi and Moonshot AI"
delta_of: ai-industry-kb-2026
sha256: a1cb7a0063f803ff47112cd5b3dcfe0d651c20bf8806f5d3f605eeaae67a60f5
---

# Kimi K3 — license conflict (recorded, unresolved)

### Kimi K3 — license conflict (recorded, unresolved)
- Some coverage describes a **custom "Kimi K3 License"** [SECONDARY] (ai-stack.ai).
- Other coverage describes a **modified MIT license with a separate commercial agreement required after $20M MaaS revenue over 12 consecutive months** [SECONDARY] (felloai.com).
- The existing §5 records "Modified MIT thresholds" — this file records the conflict explicitly: **do not state a single K3 license characterization until the license file on the exact checkpoint is read**; both characterizations are in circulation [DIRECTIONAL].

### MuonClip optimizer — technical detail
- MuonClip = **Muon + weight decay + RMS matching + QK-Clip** [SECONDARY] (github.com/rocm/maxtext Run_Kimi.md; wilsonwu-ai/scaling-open-models).
- QK-Clip mechanism: **scales offending attention-head Q/K projections to cap maximum logits** [SECONDARY] (rocm/maxtext notes; scaling-open-models).
- Vendor claim: **stable 15.5T-token training with no loss spike** on K2 [VENDOR] (reported via secondary coverage — vendor number).

### Kimi For Coding / Kimi Code product layer
- **Kimi Code CLI launched January 2026**, crossed **6,400+ GitHub stars**, and shipped with **K2.6 as its default backend** [SECONDARY] (medium.com/@tentenco, 2026-04).
- **`kimi-for-coding`** model ID now serves **Kimi K2.8 Preview** (model ID unchanged since 2026-09; underlying model upgraded); described as overall close to K3 with 1M context and low/high/max effort levels, available to all memberships [COMMUNITY] (benedictking/ccx kimi.md — community-maintained provider doc, mark [UNVERIFIED] for vendor confirmation).
- **`kimi-for-coding-highspeed`** = Kimi K2.7 Code HighSpeed [COMMUNITY] (benedictking/ccx kimi.md).
- Kimi Code and Moonshot are **separate providers**: keys are not interchangeable, endpoints differ, and model refs differ (`moonshot/...` vs `kimi/...`) [COMMUNITY] (gemmaclaw/gemmaclaw moonshot.md).

### API endpoints, tiers, and rate-limit mechanics
- Endpoints: global OpenAI-compatible **`https://api.moonshot.ai/v1`**; China pay-as-you-go **`https://api.moonshot.cn/v1`**; Kimi Code Anthropic **`https://api.kimi.com/coding/`**, OpenAI **`https://api.kimi.com/coding/v1`** [COMMUNITY] (benedictking/ccx kimi.md; dev.to).
- Membership tiers observed: **Moderato** and **Allegretto**; `k3` serves 256K context on Moderato and up to **1M on Allegretto+**, with low/high/max effort levels [COMMUNITY] (benedictking/ccx kimi.md).
- Rate-limit gotcha: Kimi books tokens against the caller's limit based on **input + max_completion_tokens at request time, regardless of actual generation** — not on tokens actually consumed [COMMUNITY] (dev.to, 2026-09).
- Account activation: creating a key is free but calls require **prepaid credits — a $1 minimum activates the account** [COMMUNITY] (dev.to, 2026-09).
- K3 pay-as-you-go: **$3.00 input / $0.30 cache-hit / $15.00 output** per 1M tokens — consistent with existing §5 [SECONDARY] (saygm.com; dev.to).
- K2.5 pay-as-you-go: **$0.60 / $0.10 cached / $3.00** [SECONDARY] (saygm.com; gemmaclaw docs; cloudprice.net lists $0.6/$3.00/$0.1).
- K2.5 launch-era pricing in April 2026 coverage: **$0.44 in / $2.00 out** (Medium comparison) and **$0.45/$2.25** (cloudprice version table) — older regional/third-party figures; use the saygm.com figures as the current list reference [DIRECTIONAL].

### Moonshot funding timeline (2026)
- **Early 2026: $700M round at $10B valuation** [SECONDARY] (clay.com dossier; mlq.ai).
- **March 2026: $18B valuation** reported in Chinese AI funding coverage — a dated/intermediate report, not directly combinable with the May round [SECONDARY] (theagenttimes.com).
- **May 2026: ~$2B round led by Meituan's Long-Z Investment at $20B+ valuation**; reported participants **Tsinghua Capital, China Mobile, CPE Yuanfeng**; coverage reports **$3.9B raised in six months** and **ARR above $200M in April** [SECONDARY] (mlq.ai; clay.com dossier).
- **$30B valuation figure** in later coverage refers to **funding talks, not a closed round** — do not present it as a valuation [SECONDARY] (ainvest.com, 2026).
- Cross-link: Meituan's Long-Z Investment vehicle appears both here and in the LongCat/Meituan section — Meituan is simultaneously building its own lab and bankrolling Moonshot [DIRECTIONAL].

### Anthropic distillation allegations (framing discipline)
- Anthropic alleged **large-scale distillation campaigns** targeting Claude by three Chinese labs including Moonshot [SECONDARY] (computerworld.com; infoworld.com, 2026).
- Figures as alleged by Anthropic: **Moonshot generated 3.4M+ Claude exchanges** targeting reasoning, tool use, coding, and computer vision; **16M+ interactions overall through ~24K fraudulent accounts** across the three labs [SECONDARY] (computerworld.com).
- Phrase strictly as **Anthropic's allegation**; monitoring evidence was not independently verified; no admission or adjudication is recorded in the sources reviewed [DIRECTIONAL].

## Figures and metrics

| Model | Date | Size (total/active) | Context | License | Provenance |
|---|---|---|---|---|---|
| Kimi K2 | 2025-07 (mid) | 1T / 32B | 128K | Modified MIT | [VENDOR] |
| Kimi K2.5 | 2026-01-27 | 1T / 32B | 256K (200K at launch) | Modified MIT | [VENDOR] |
| Kimi K2.6 | 2026-04-20 | 1T / 32B (384 experts) | 256K | Modified MIT | [VENDOR] |
| Kimi K2.7 Code | 2026-06-12 | post-train of K2.6 | 256K | Modified MIT | [VENDOR] |
| Kimi K3 | 2026-07-16 (API) / 07-27 (weights) | 2.8T | — | Modified MIT | [SECONDARY] |

| Benchmark | Score | Date/version | Provenance |
|---|---|---|---|
| Kimi K2.7 Code independent benchmarks at launch | none (vendor-only) | 2026-06-12 | [VENDOR] |
| Kimi K3, AA Intelligence Index | 43.8 | v4.3, 2026-09-07 | [SECONDARY] |
| Kimi K3, LMArena Frontend Code | 1,679 Elo (#1, first open top) | Sept 2026 | [SECONDARY] |
| Kimi K3, AA $/Index-task | $2.00 | Sept 2026 | [SECONDARY] |
| Kimi K3 API price | $3.00/$15.00 per M in/out | 2026-07-16 price card | [VENDOR] |
| Kimi commercial surface | free K3 quota (gateway) · Kimi Code from $19/mo | 2026 | [SECONDARY] |

**No cross-version comparison:** the 43.8 is v4.3-pinned; Baseten's "60" for K3 is the AA percentage-presentation scale, not the Index scale — never mix. [DIRECTIONAL]


### New verified metrics — expansion (continued — K3 license/economics figures)

- 96 shards / ~1.56 TB; MXFP4 QAT; 104B active [SECONDARY].
- $0.27 vs $2.00 per AA task (7.4×); $0.41 vs $17.59 Vals (~40×); 15 vs 90 min [SECONDARY].
- 214.4 vs 34.7 tok/s (6.2×); 98% vs 90% cache; $0.003 vs $0.30 cache-read [SECONDARY].
- Day-0: Together/Fireworks/Nebius/vLLM/SGLang/Cursor/Devin; $5.40 blended [SECONDARY].


### New verified metrics — expansion (continued — ecosystem pricing snapshot (July 2026, OpenRouter))

- Kimi K3 $3/$15; GLM 5.2 $0.97/$3.06; DeepSeek V4 Pro $0.44/$0.87; V4 Flash $0.10/$0.20 [COMMUNITY] (model-router).


### New verified metrics — expansion

### Current pay-as-you-go price table (list prices, per 1M tokens)
| Model | Input | Cached input | Output | Source |
|---|---|---|---|---|
| K2.5 | $0.60 | $0.10 | $3.00 | [SECONDARY] saygm.com; gemmaclaw docs |
| K2.6 | $0.95 | $0.16 | $4.00 | [SECONDARY] saygm.com; gemmaclaw docs |
| K2.7 Code | $0.95 | $0.19 | $4.00 | [SECONDARY] saygm.com |
| K2.7 Code Highspeed | $1.90 | $0.38 | $8.00 | [SECONDARY] saygm.com |
| K3 | $3.00 | $0.30 | $15.00 | [SECONDARY] saygm.com; dev.to |

- Price evolution (not contradiction): K2.6 launched April 2026 at $0.60/$2.50 (in/out) [SECONDARY] (Medium/@tentenco); current list $0.95/$4.00 [SECONDARY] (saygm.com). K2.5 launch-era figures $0.44/$2.00 [SECONDARY] (Medium/@endlesslyimprovisng) vs current $0.60/$3.00.
- Highspeed multiplier: ~5–6× faster output at 2× the token price vs K2.7 Code base [SECONDARY/COMMUNITY] (saygm.com; ccx docs).
- K3 costs 5× K2.5 input price ($3.00 vs $0.60) — the premium tier step [DIRECTIONAL].

