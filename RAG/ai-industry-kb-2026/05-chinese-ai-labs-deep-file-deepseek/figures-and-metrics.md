---
id: ai-industry-kb-2026/05-chinese-ai-labs-deep-file-deepseek/figures-and-metrics
title: "Figures and metrics"
domain: chinese-ai-labs-deep-file-deepseek
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "CISA", "DeepSeek", "Hugging Face", "Moonshot", "OpenAI"]
dates: ["2026-03", "2026-03-09", "2026-04-24", "2026-05-31", "2026-06", "2026-06-01", "2026-08-12", "2026-08-13", "2026-08-16", "2026-09-10"]
keywords: ["agent", "agentic", "apache", "attention", "benchmark", "benchmarks", "claude", "compute", "cost", "cybersecurity", "decode", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1812, 1952]
section: "5. Chinese AI Labs — Deep File: DeepSeek"
sha256: ab7d6f127e3ea067996b297fd7ac2f6b89d1c914fdc37a3cd0d511a6f4184018
---

# Figures and metrics

## Figures and metrics

### V4-Lite (2026-03-09) — all [UNVERIFIED]/secondary

- Reported scale ~200B parameters (36kr via awesomeagents; no official parameter count exists).
- 1M context; codenames "Sealion-Lite" / "Healer Alpha" / "0302" (multiple secondary trackers).
- 36kr's report additionally said V4 Lite **does not use the Engram conditional memory system** DeepSeek co-developed with Peking University — implying the Engram narrative was attached to the full V4 at that time [rumor-phase secondary].

### V4 Preview (2026-04-24) — verified against DeepSeek changelog + cadence tables

- **V4-Pro**: 1.6T total / **49B active per token**; 61 layers; 384 routed experts, 6 active, 1 shared (secondary reconstruction [UNVERIFIED]).
- **V4-Flash**: 284B total / **13B active**; 43 layers; 256 routed experts, 6 active, 1 shared (secondary reconstruction [UNVERIFIED]).
- Context 1M tokens; max output 384K; MIT weights on Hugging Face and ModelScope; FP4/FP8 mixed precision (hermes-workspace scout, June 2026) [COMMUNITY].
- V4-Pro pricing: **$0.435 input / $0.87 output per 1M** (cache-miss) from 2026-06-01 (roninforge price backfill) [COMMUNITY]; 0813 GA priced **$0.44/$0.87** (it-daily).
- Peak/off-peak from 2026-08-16: V4-Pro output **$3.96 peak / $1.98 off-peak** (355%/128% above the old $0.87 flat); V4-Flash output $1.32/$0.66; cache-hit input increases up to **1,100%** (wave2.1/06, techtimes).
- V4 architecture reconstruction details (secondary, [UNVERIFIED]): every block is MoE; **first three blocks use deterministic token-ID hash routing** instead of learned routing; aux-loss-free load balancing retained with a small sequence-wise balance loss; routing affinity changed from sigmoid to **sqrt(softplus(...))**; routed expert weights reportedly deployed in FP4.
- V4 hybrid attention (buzzgrewal's analysis of the 58-page report): **CSA (Compressed Sparse Attention)** — compresses the KV cache ~4×, applies lightning-indexer-style top-k over compressed blocks (NSA/DSA lineage); **HCA (Heavily Compressed Attention)** — compresses the KV ~128× and runs dense attention over compressed blocks. DSA (DeepSeek Sparse Attention) **reduces attention compute**; do not equate it with KV-cache reduction. The report uses a "hash routing" trick only for assigning tokens to MoE experts, unrelated to memory.

### V4-Pro-0813 evaluation (2026-08-12/13)

- Vals.ai (Aug 12, pre-release eval): **Vals Index 52.37%**, **#18** at evaluation time, **+9.48 pts vs DeepSeek V4 (42.89%)**. The brief's "+10.6" has no source and is contradicted.
- **SWE-bench Verified 96.40% (#2 of 82)** — highest-scoring open-weight model on the board, ahead of Kimi K3 (93.40%); ~$0.02 per test run vs $1.29 for the closed leader.
- **ProofBench 49.00%** (up from 10.00%, #45 → #15); **Legal Research Bench 40.87%** (up from 23.08%, #27 → #11); Harvey's Legal Agent Benchmark 7.50% (#10 of 43).
- Weaknesses: **Terminal-Bench 2.1 54.68%** (#33 of 52; 28.89% on hard tasks); **EMB 52.80%** (#24 of 37); no temperature parameter accepted; evaluations run at max reasoning effort.
- Artificial Analysis Intelligence Index **53** — an 8-point gain over the April preview (miraflow synthesis; v4.1-era scale); behind GPT-5.6 Terra (-4) and Kimi K3 (-7); among open weights, second only to Kimi K2.6 per miraflow's read. On the rebased **v4.3 scale** (methodology changed Sept 4), V4 Pro 0813 ≈ **36** — not comparable across versions.
- Rank drift: mid-September press (coinlive, nationpress, SCMP, ~40 days before Sept 22) reports **12th on the Vals Index** — the index is live and recomputed; record rank with its date, never as a static fact.
- Aikido Security (Belgian cybersecurity firm): V4-Pro-0813 **outperformed every other model tested at finding system vulnerabilities** — significantly more detections than Claude Opus 5 or Qwen 3.8 — but with poor precision (coinlive, SCMP, it-daily) [SECONDARY-level reporting].
- Pricing criticism: coinlive calls ~$0.44/M input "somewhat expensive" vs a 33-cent market input median; miraflow frames the GA as a ~**3.6× price increase over promotional April-preview pricing** for the 8-point AA-index gain [DIRECTIONAL framing — attribute it].

### V4.1-Flash GA (2026-09-10) — all [VENDOR] unless noted

- **552B-parameter** MoE backbone; **8B active per token in prefill, 16B in decode** (asymmetric split).
- **Causal Encoder-Decoder (CED)**: 40-layer Transformer split into a **20-layer causal encoder + 20-layer decoder**; decoder's global KV cache projected from the encoder's final hidden states. "Compressed Expert Dispatch" is not a DeepSeek term — never use it.
- **CSA2** (Compressed Sparse Attention 2): new attention kernel of the family, successor lineage to V4's DSA.
- **Global KV cache 890 bytes/token** — ≈ **1/4 of V4-Flash** at the same context length (~75% reduction — this is the brief's "up to 75% vs previous versions" claim, verified as stated by DeepSeek) and ~1/437 of DeepSeek-V1 on the vendor's own chart.
- **FP4 KV caching in E2M1 format** plus shared attention indices across layers; **SWA Bounded Replay** cuts the **persistent** KV-cache footprint to ~1/8 of V4-Flash.
- **Engram**: 196B-parameter conditional memory module accessed sparsely via token-based lookup — [VENDOR] for V4.1-Flash, with the March-2026 rumor-provenance flag (the V4 report mentions "Engram"/"O(1)" zero times).
- MoE layout: **1 shared expert + 384 routed experts, 6 active per token**.
- Context **1M tokens** (1,048,576 per third-party trackers); **384K max output**; native vision (text + images in, text out); continuously controllable reasoning effort dial **1–100**.
- Training: from scratch on **45T tokens** (mixed text + images); context extended to 1M at the 34T-token mark; post-training SFT + RL + on-policy distillation; agentic gains attributed to large-scale automated synthesis of agent tasks/environments [VENDOR].
- Vendor launch benchmarks [VENDOR]: GPQA Diamond **90.9**; Terminal-Bench 2.1 **90.6**; Codeforces **3471**; DeepSWE v1.1 **74.2%** vs Claude Opus 5's 74.0% (harness swing 65.5%–74.2%, Techpresso — harness-sensitive); the vendor's own report concedes harder evals: Terminal-Bench 3.0 **30.0 vs 43.3** (Opus 5), Terminal-Bench 4.0 **31.2 vs 51.8**, HLE **36.8 vs 56.3** (63.9 with tools) — strength is tool-using agentic execution, not unaided abstract reasoning.
- Pricing (effective 2026-09-10): **$0.15 input / $0.60 output per 1M off-peak**; peak 2× ($0.30/$1.20); weekends off-peak; APIMaster quotes CNY 1 / CNY 4 off-peak. Replaces V4-Flash pricing ($0.14/$0.27–0.28) — output roughly doubled off-peak vs old Flash tier.

### V4-Pro-0813 benchmark ledger (dated, attributed)

- Vals Index: **52.37%**, #18 on 2026-08-12 (eval date); **+9.48 pts vs DeepSeek V4 (42.89%)** — vals.ai. Brief's "+10.6" contradicted, no source found.
- Vals Index rank drift: **12th** in mid-September press (coinlive, nationpress, SCMP, ~40 days before Sept 22) — live index; keep rank+date paired.
- SWE-bench Verified: **96.40%**, #2 of 82 — highest open-weight score, ahead of Kimi K3 (93.40%); ~$0.02 per test run vs $1.29 for the closed leader (vals.ai).
- ProofBench: **49.00%** (up from 10.00%, #45 → #15) — vals.ai.
- Legal Research Bench: **40.87%** (up from 23.08%, #27 → #11) — vals.ai.
- Harvey's Legal Agent Benchmark: **7.50%** (#10 of 43) — vals.ai.
- Terminal-Bench 2.1: **54.68%** (#33 of 52; 28.89% on hard tasks) — vals.ai; a relative weakness.
- EMB: **52.80%** (#24 of 37) — vals.ai; relative weakness.
- Eval constraints: no temperature parameter accepted; evaluations run at max reasoning effort (vals.ai).
- Artificial Analysis Intelligence Index: **53** (v4.1-era scale), **+8 over the April preview** (miraflow synthesis); behind GPT-5.6 Terra (-4) and Kimi K3 (-7); among open weights second only to Kimi K2.6 per miraflow's read.
- AA Intelligence Index v4.3 (rebased Sept 4, 2026): V4 Pro 0813 ≈ **36** — not comparable to v4.1-era figures.
- Aikido Security: led all tested models in vulnerability-detection count; significantly more detections than Claude Opus 5 or Qwen 3.8; poor precision (coinlive, SCMP, it-daily) [SECONDARY reporting].
- GA pricing: **$0.44/$0.87** per 1M (it-daily); coinlive "somewhat expensive" vs 33-cent market input median; miraflow: ~**3.6× price increase over promotional April-preview pricing** for the 8-point AA gain [DIRECTIONAL — attribute].

### V4.1-Flash vendor benchmark ledger (all [VENDOR])

- GPQA Diamond: **90.9** — DeepSeek launch table.
- Terminal-Bench 2.1: **90.6** — DeepSeek launch table.
- Codeforces: **3471** — DeepSeek launch table.
- DeepSWE v1.1: **74.2%** vs Claude Opus 5's 74.0% — a 0.2-point margin presented as matching-frontier evidence; harness-sensitive swing **65.5%–74.2%** (Techpresso).
- Terminal-Bench 3.0: **30.0 vs Opus 5's 43.3** — conceded in the vendor's own technical report.
- Terminal-Bench 4.0: **31.2 vs 51.8** — conceded in the vendor's own technical report.
- HLE: **36.8 vs 56.3** (Opus 5) / 44.5 (GPT-5.6 Sol); **63.9 with tool access** — the architecture's strength is tool-using agentic execution, not unaided abstract reasoning.
- Status: no independent reproduction of any launch figure yet — treat as vendor claims.

### Pricing ledger (dated)

- 2026-04-24 → 2026-05-31: promotional V4-preview pricing (roninforge backfill) [COMMUNITY].
- 2026-06-01: permanent **75% flagship price cut** announced (Digitimes); V4-Pro list **$0.435/$0.87** cache-miss input/output per 1M (roninforge price backfill) [COMMUNITY].
- 2026-08-13: V4-Pro-0813 GA priced **$0.44/$0.87** (it-daily).
- 2026-08-16: peak/off-peak takes effect — V4-Pro output **$3.96 peak / $1.98 off-peak** (355%/128% above the old $0.87 flat); V4-Flash output **$1.32/$0.66**; cache-hit input increases up to **1,100%** (wave2.1/06, techtimes).
- 2026-09-10 04:00 UTC: V4.1-Flash pricing live — **$0.15 input / $0.60 output off-peak**; peak 2× ($0.30/$1.20); weekends off-peak; CNY 1 / CNY 4 off-peak (APIMaster); replaces V4-Flash pricing ($0.14/$0.27–0.28).
- Full serving/economics analysis lives in §9/§14 of the final document — this ledger is the dated skeleton only.

### V4-Lite fact sheet (2026-03-09) — all [UNVERIFIED]/secondary

- Event: a "V4 Lite" update **appeared on DeepSeek's website on March 9, 2026** (nxcode FAQ as of March 2026; ez0000001000000's V4-leak roundup: "reportedly been released/surfaced as of March 9, 2026").
- Codenames (secondary trackers only): "Sealion-Lite", "Healer Alpha", "0302".
- Scale: ~**200B parameters** — reported by 36kr (relayed by awesomeagents); **no official parameter count exists**.
- Context: 1M tokens (awesomeagents) — extends the Feb 11 silent 1M upgrade pattern.
- Framing: "leaked under NDA" (awesomeagents), "soft-launched" (aiforautomation.io), "NDA-testing" — not a GA release; treated by March coverage (kissapi.ai, geeky-gadgets) as a preview of V4-family improvements ahead of the full V4.
- Engram note: 36kr's report said V4 Lite **does not use the Engram conditional memory system** — implying the Engram narrative was attached to the full V4 at that time (see Engram provenance ledger).
- Wave 3 topic 2 and Wave 2.1 do not document V4-Lite — this file is the sole record; keep the leak framing and the unconfirmed figure together.

### V4 architecture reconstruction fact sheet [UNVERIFIED — secondary reconstructions, no directly verified official report]

- Variant table: V4-Flash = 284B total / 13B active, 43 layers, 256 routed experts, 6 active, 1 shared; V4-Pro = 1.6T / 49B active, 61 layers, 384 routed experts, 6 active, 1 shared.
- DeepSeek-style fine-grained MoE: every block is MoE; fine granularity raises routing specialization and combinatorial flexibility (cf. wave2/03 §3 — systems literature notes the bandwidth/training-efficiency cost of fine granularity, preserved only by co-designed mitigations).
- **First three blocks use deterministic token-ID hash routing** instead of learned routing.
- Aux-loss-free load balancing retained from the V3 lineage, with a small additional **sequence-wise balance loss**.
- Routing affinity changes from sigmoid to **sqrt(softplus(...))**.
- Routed expert weights reportedly deployed in **FP4**.
- Attention: **CSA (Compressed Sparse Attention)** — compresses KV cache ~4×, lightning-indexer-style top-k over compressed blocks (NSA/DSA lineage); **HCA (Heavily Compressed Attention)** — compresses KV ~128×, dense attention over compressed blocks.
- **DSA = DeepSeek Sparse Attention**: reduces **attention compute** — do NOT equate with KV-cache reduction.
- The report's "hash routing" trick is used only for assigning tokens to MoE experts — nothing to do with memory.
- MLA (Multi-Head Latent Attention): V4 is described as replacing the earlier MLA design with the CSA/HCA hybrid attention (peripheral to this topic; keep the lineage note).
- Evidence grade: consistent across multiple independent secondary analyses and aligned with the V3 lineage (aux-loss-free dynamic bias), but each numeric detail is [UNVERIFIED] pending the official report.

### V4.1-Flash architecture fact sheet [VENDOR — model card via secondary coverage, no independent reproduction]

- Backbone: **552B-parameter MoE**; **8B active per token in prefill, 16B in decode** (asymmetric split).
- **CED = Causal Encoder-Decoder**: 40-layer Transformer = **20-layer causal encoder + 20-layer decoder**; decoder's global KV cache projected from the encoder's final hidden states.
- **CSA2** (Compressed Sparse Attention 2): new attention kernel of the family; successor lineage to V4's DSA.
- Global KV cache: **890 bytes/token** ≈ 1/4 of V4-Flash (~75% reduction — the brief's "up to 75%" claim, verified as stated by DeepSeek); ~1/437 of DeepSeek-V1 on the vendor chart.
- **FP4 KV caching in E2M1 format**; shared attention indices across layers.
- **SWA Bounded Replay**: persistent KV-cache footprint ≈ **1/8 of V4-Flash**.
- **Engram**: 196B-parameter conditional memory module, sparse token-based lookup — [VENDOR] with March-2026 rumor-provenance flag.
- MoE layout per layer: **1 shared expert + 384 routed experts, 6 routed active per token**.
- Context: **1M tokens** (1,048,576 per third-party trackers); max output **384K**; native vision (text + images in, text out; image tokens billed as text tokens).
- Reasoning: continuously controllable effort dial **1–100** per request.
- Training: from scratch on **45T tokens** (mixed text + images); context extended to 1M at the 34T-token mark; post-training SFT + RL + on-policy distillation; agentic gains attributed to large-scale automated agent-task synthesis [VENDOR].
- License: **MIT** weights on Hugging Face; technical report published alongside.
- Positioning per DeepSeek: smallest model of an entirely new architecture family — not a V4 checkpoint refresh [VENDOR].
- Retirement inversion: DeepSeek claims V4.1-Flash "comprehensively surpassed V4 Pro across performance, cost, speed, and task completion time" and moves to retire the bigger model in favor of the smaller, cheaper one — justified by the claim that memory, not parameter count, is the binding serving constraint [VENDOR].

### License and weight-distribution facts

- V4 family (Preview, 2026-04-24): **MIT** open weights on Hugging Face and ModelScope — verified against wave3/02; BenchLM's September directory lists V4 weights as MIT open weight.
- V4.1-Flash (2026-09-10): **MIT** weights on Hugging Face; technical report published alongside.
- March 2026 speculative coverage (kissapi.ai, aiforautomation.io, books.brightlearn.ai) repeatedly expected **Apache 2.0** for the V4 family — never materialized; do not carry forward.
- The MIT regime covers self-host, fine-tune, commercial deployment, and redistribution without restrictions (multiple secondary sources confirm).

### Evidence-grade legend used in this file

- [VENDOR]: DeepSeek's own claims (changelog, pricing page, WeChat announcement, model card) — via primary-adjacent or secondary relay.
- [UNVERIFIED]: figures with no official source (V4-Lite 200B; V4 architecture reconstructions; V4.1 benchmark/architecture claims pending independent reproduction).
- [COMMUNITY]: community-detected or community-maintained evidence (Feb 11 silent upgrade; ai-pricelog/roninforge backfills; hermes-workspace scout).
- [DIRECTIONAL]: secondary framing/analysis (miraflow's price-increase framing; orcarouter's day-0 caution).
- [SECONDARY-level reporting]: press relay (SCMP, coinlive, nationpress, it-daily) for events like the Aikido finding and stealth GA mechanics.

### Engram provenance ledger

- **March 2026 (rumor phase)**: "Engram" / "conditional memory system" co-developed with Peking University enters discourse via 36kr reporting (relayed by awesomeagents); O(1) memory narrative, "97% needle-in-a-haystack at 1M", "1M context costs roughly the same as 128K" (buzzgrewal's summary of the rumor cycle). V4 Lite reportedly does NOT use Engram (36kr) [UNVERIFIED].
- **2026-04-24 (V4 report)**: the official 58-page V4 technical report mentions "Engram" and "O(1)" **zero times** (buzzgrewal). V4's actual memory design: CSA (~4× KV compression, top-k over compressed blocks, NSA/DSA lineage), HCA (~128× KV compression, dense attention over compressed blocks). The report's "hash routing" is token→expert assignment only, unrelated to memory.
- **2026-09-10 (V4.1 model card via secondary)**: Engram reappears as a **196B-parameter conditional memory module** accessed sparsely via token-based lookup — keep as [VENDOR], never as established DeepSeek technology.

