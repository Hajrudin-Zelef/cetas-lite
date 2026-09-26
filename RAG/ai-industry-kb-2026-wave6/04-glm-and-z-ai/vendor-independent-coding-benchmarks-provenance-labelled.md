---
id: ai-industry-kb-2026-wave6/04-glm-and-z-ai/vendor-independent-coding-benchmarks-provenance-labelled
title: "Vendor/independent coding benchmarks (provenance-labelled)"
domain: glm-and-z-ai
role: deep-dive
task: benchmark
actors: ["Anthropic", "Baseten", "DeepSeek", "Fireworks AI", "OpenAI", "OpenRouter", "Z.ai"]
dates: ["2026-07", "2026-08-26", "2026-08-28", "2026-08-29", "2026-09-09"]
keywords: ["benchmark", "benchmarks", "agentic", "attention", "claude", "compute", "cost", "cyber", "deepseek", "fable 5", "fine-tuning", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1858, 1904]
section: "§4. GLM and Z.ai"
delta_of: ai-industry-kb-2026
sha256: a1fb28b5c4800954eab4d2e60d983503dc3969bd1d4fe7fe680e56f1aa8527aa
---

# Vendor/independent coding benchmarks (provenance-labelled)

### Vendor/independent coding benchmarks (provenance-labelled)
- **SWE-bench Pro 62.1** vs GPT-5.5 58.6 vs Claude Opus 4.8 69.2 [SECONDARY] (techtimes.com; cryptobriefing.com).
- **Terminal-Bench 2.1: 81.0** — best open-source result at the time, near Claude Opus 4.8's 85.0 [SECONDARY] (cryptobriefing.com; toknow.ai).
- **FrontierSWE 74.4** vs Claude Opus 4.8 75.1 vs GPT-5.5 72.6 — evaluated by **independent third-party firm Proximal**, not Z.ai [SECONDARY] (techtimes.com).
- **MCP-Atlas 77.0** — nearly on par with Claude Opus 4.8's 77.8 [SECONDARY] (techtimes.com).
- **Artificial Analysis Intelligence Index: 51** — ranked **first among open-weights models** in AA's 9-eval composite; AA clocked **168.8 output tokens/sec** but flagged it as **token-hungry (~43K output tokens per task)**, inflating real cost above sticker price [SECONDARY] (pilot-shell blog).
- **AA-Briefcase (agentic knowledge work): Elo 1266 at $2.40/task** — between GPT-5.5 and Opus 4.8 (1356 at $10.40); Claude Fable 5 far ahead at 1587 [SECONDARY] (pilot-shell blog).
- **Semgrep IDOR cyber benchmark: 39% F1** (prompt-only, Pydantic-AI) — edging Claude Code on Opus 4.6 (37%) and Opus 4.8 (28%) at **~$0.17 per vulnerability**; Semgrep's caveat: "one task, one dataset, one run"; Sonnet 5 not tested [SECONDARY] (pilot-shell blog).
- **Code Arena: 2nd globally** on front-end web development, trailing only Claude Fable 5; **beat Claude Fable 5 on the crowdsourced Design Arena** design-task benchmark with **Elo 1360** [SECONDARY] (indianexpress.com).
- **Databricks July 2026 enterprise test**: multi-million-line internal coding test, GLM-5.2 on par with Claude Opus 4.8 at **$1.28/task vs $1.94 — ~34% cheaper** for equivalent quality [SECONDARY] (cryptobriefing.com).
- Price-to-performance: about **1/6 the cost of GPT-5.5** while surpassing it on several benchmarks [SECONDARY] (medium/@wenmingtech).
- GLM-5.3 vs 5.2 on SWE-bench Pro (independent table, starred = vendor/cited): **5.3 at 64.6 vs 5.2 at 62.1** — post-training-only improvement direction [SECONDARY] (intelligentliving.co, 2026-09).

### GLM-5.3-Flash — deep spec from the pi-configs evidence brief (2026-08-29)
- **A new base model, not a 5.3 post-training refresh** — this distinction matters for §4's model-lineage accuracy [COMMUNITY] (pi-configs evidence brief, citing z.ai blog and arXiv 2602.15763).
- Hybrid sparse+linear attention: **linear attention for local dependencies via state modeling, sparse indexer for global retrieval**; **IndexPool compresses 4 indexer key vectors into 1** [COMMUNITY] (pi-configs brief).
- Efficiency vs GLM-5.3: **3.0× less attention compute, 4.4× smaller KV cache** [COMMUNITY] (pi-configs brief).
- **mHC (Manifold-Constrained Hyper-Connections)**; **30T-token multimodal pre-train**; paper **arXiv 2602.15763**; **45 layers vs GLM-4.5's 92** [COMMUNITY] (pi-configs brief).
- Modalities: **vision + text**; confirmed live on Fireworks with `image_url` blocks accepted (alertmanager screenshot read correctly in testing) [COMMUNITY] (pi-configs brief, live probe 2026-08-29).
- Context enforcement on the Fireworks route: **prompt + max_tokens ≤ 1,048,576** — a 917,505-token prompt with max_tokens 131072 fails with a 400 that looks like a model failure [COMMUNITY] (pi-configs brief).
- **Thinking is always on**; Fireworks enum `low/medium/high/xhigh/max`; `none` rejected ("reasoning cannot be disabled"); no `minimal` [COMMUNITY] (pi-configs brief, probed 2026-08-29).
- Weights note: brief dated 2026-08-29 says **"MIT promised, not shipped at Fireworks launch"** — in tension with day-one-HF-weights reporting; record as a **source difference on weights timing**, not a contradiction of the license [COMMUNITY vs SECONDARY].
- Live on Fireworks **2026-08-29** as `fireworks/accounts/fireworks/models/glm-5p3-flash`; brief status: **"default-seat challenger"** [COMMUNITY] (pi-configs brief).
- Vendor claims in the brief (all [VENDOR]): **AA Intelligence Index 57 at $0.045/task** ("intelligence previously only available at roughly 10× the cost"); **outperforms GLM-5.2 across benchmarks at one-tenth the price**; **approaching Claude Opus 4.8** on coding/agentic benchmarks; **GLM Coding Plan quota: 3× the points of 5.3-proper**.
- Independent local eval (2026-08-29, Fireworks, max effort): routine pack **12/12 vs DeepSeek Flash 0731** ($0.0199 vs $0.0240, 19s vs 25s, 10.2K vs 22.4K output tokens); hard pack **4/4** ($0.0087 vs $0.0162); **needle ~925K: 2/2 PASS at $0.139 each** — 5.2 does this at ~$1.30, **5.3-proper cannot (hangs)**; vision 2/2 PASS; guess-vs-abstain **5/5 with 0 fabrications**; strict json_schema + thinking **10/10** [COMMUNITY] (pi-configs brief).
- Throughput (controlled stream, 2026-08-29): **~50–64 tok/s steady-state**, TTFT 0.13s warm / 1–4s cold; agentic wall-clock 36.5 tok/s (vs DeepSeek 81, 5.3-proper 22) [COMMUNITY] (pi-configs brief).
- OpenRouter field data: **50 tps / 1.03s TTFT / 99.84% completion** over the observed window [COMMUNITY] (pi-configs brief).
- Z.ai's own same-harness scorecard: **Flash DeepSWE 63.4 vs 5.3-proper 66.9; Terminal-Bench 2.1 84.3 vs 88.2; AutomationBench 48.8; Toolathlon 78.4; GDPval-AA 1773** — beating Opus 4.8 (1582), K3 (1685), Sol (1728); only Opus 5 (1852) higher [VENDOR via community brief].
- Per-success economics: Flash ~$0.15–0.25/task at 63.4% vs proper ~$1.50–2/task at 66.9% — the ~9× token premium buys ~3.5–4 pass points, favoring Flash **~8–10× per success** [COMMUNITY analysis] (pi-configs brief).
- Cross-provider throughput spread on identical weights — **Baseten 109 / Friendli 77 / Fireworks 50 tps** — proves throughput is **deployment-bound, not model-bound** [COMMUNITY] (pi-configs brief).
- Prefill already fast: **~23K tok/s on the 925K needle** [COMMUNITY] (pi-configs brief).
- Competitive notes: cheaper than **DeepSeek V4 Flash 0731** on both axes after its 2026-08-29 hike to $0.22/$0.66; DeepSeek retains a **384K output cap (vs 64K)** and **$0.007 cache-read (vs $0.029)** [COMMUNITY] (pi-configs brief).
- Positioning: the **only sub-$1 model on the route serving ~925K context** — takes 5.2's "marathon" role at ~1/9th the cost where <64K output suffices [COMMUNITY] (pi-configs brief).


### New verified metrics — expansion

### GLM-5.3-Flash — API pricing and launch promotion
- Standard API pricing: **$0.15/M input / $0.029/M cached input / $0.50/M output** tokens [SECONDARY] (model-guide research note; memeburn.com).
- Launch promotion: **$0.075/M input / $0.25/M output** through **September 9, 2026** [SECONDARY] (model-guide research note; startupfortune.com pricing discussion).
- Z.ai's positioning: Flash **beats GLM-5.2 across benchmarks and real workloads at roughly one-tenth the price** [VENDOR] (MarkTechPost, 2026-08-26, reporting Z.ai).
- Flash lands **within half a point of Claude Opus 4.8** on Z.ai's internal coding benchmark [VENDOR] (MarkTechPost, 2026-08-26; Medium/noahkenji283, 2026-08-28 — vendor's internal benchmark, not independently rerun).

### GLM-5.3-Flash — independent/community benchmark note
- On **KingBench**, GLM-5.3-Flash scored **63/80**, slightly **below its anonymous Ox Alpha score** — evidence that anonymous-preview scores and branded-release scores can diverge [COMMUNITY] (YouTube/AISeeKing breakdown, 2026-08).
- Strengths noted in community testing: reasoning, math, agentic coding, and local fine-tuning tasks [COMMUNITY] (YouTube/AISeeKing).

