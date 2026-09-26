---
id: ai-industry-kb-2026/05-chinese-ai-labs-deep-file-deepseek/v4-architecture-reconstruction-fact-sheet-unverified-seconda
title: "V4 architecture reconstruction fact sheet [UNVERIFIED — secondary reconstructions, no directly verified official report]"
domain: chinese-ai-labs-deep-file-deepseek
role: deep-dive
task: actor-profile
actors: ["CISA", "DeepSeek", "Hugging Face"]
dates: ["2026-03", "2026-03-09", "2026-04-24", "2026-09-10"]
keywords: ["agent", "agentic", "apache", "attention", "benchmark", "compute", "cost", "decode", "deepseek", "distillation", "distribution", "fp4"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1893, 1952]
section: "5. Chinese AI Labs — Deep File: DeepSeek"
sha256: 2c55938aba6804711b716418f39f08573254e04e18a105db504359dbd07e6e25
---

# V4 architecture reconstruction fact sheet [UNVERIFIED — secondary reconstructions, no directly verified official report]

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

