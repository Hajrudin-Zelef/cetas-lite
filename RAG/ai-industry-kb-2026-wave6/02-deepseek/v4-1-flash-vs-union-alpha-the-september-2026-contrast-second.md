---
id: ai-industry-kb-2026-wave6/02-deepseek/v4-1-flash-vs-union-alpha-the-september-2026-contrast-second
title: "V4.1-Flash vs Union Alpha — the September 2026 contrast [SECONDARY]"
domain: deepseek
role: deep-dive
task: actor-profile
actors: ["Baseten", "DeepSeek", "Meta", "Nvidia", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-07-24", "2026-07-25", "2026-07-28", "2026-09", "2026-09-11", "2026-09-14", "2026-09-16"]
keywords: ["awq", "benchmark", "blackwell", "compute", "cost", "deepseek", "distribution", "fp4", "fp8", "glm", "gpt-5.6", "gptq"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [722, 755]
section: "§2. DeepSeek"
delta_of: ai-industry-kb-2026
sha256: 332e015cb3e37208fe648cac80976eb5b1011df9088d7b5dd8f52f98e4f4f8bd
---

# V4.1-Flash vs Union Alpha — the September 2026 contrast [SECONDARY]

### V4.1-Flash vs Union Alpha — the September 2026 contrast [SECONDARY]
- One comparison piece (OrcaRouter) frames V4.1-Flash against **Union Alpha** (launched 2026-09-16, six days after V4.1-Flash) as opposite objects: V4.1-Flash has leaderboard entries, vendor benchmark tables, and a model card; Union Alpha (at the time) had none of the three. [SECONDARY]
- Spec contrast: V4.1-Flash 1M context / 384K max output vs Union Alpha 262K / 131K; V4.1-Flash text+image input (native vision) vs Union Alpha's narrower surface. [SECONDARY]
- The piece's verdict pattern — open weights + published card + reproducible serving recipes vs closed announcement — is the corpus's template for grading September 2026 releases by evidence quality, not headline claims. [DIRECTIONAL]
### V4.1-Flash — audited revision and technical report [SECONDARY]
- One community audit (alkinun/speck) pins the release: repository `deepseek-ai/DeepSeek-V4.1-Flash`, audited revision **`df42c109f1defefcbfcedbe7d905718a12266e40`**, technical report `DeepSeek_V41_Tech_Report.pdf` (1,809,802 bytes, SHA-256 `ba68e2e4…`). The Hub API counts **~484.6B stored tensor elements** because routed-expert FP4 values are packed — not the logical model size. The artifact occupies **~510 GB**. [COMMUNITY]
- CSA2 roles, per the report: **Full** (produce main KV and indexer K, compute new Top-K selection); **Reindex** (reuse main KV and indexer K, fresh depth-specific selection); **Reuse** (reuse representation and latest selection). [SECONDARY]
- **Engram internals** (per secondary technical write-ups): two modules, half the 196B each; looks up patterns spanning **2, 3, and 4 tokens**; **8 hash heads**; ~16M entries per head; **FP8 tables**; lookup address predictable from input so retrieval from host memory can start while the transformer computes. [SECONDARY]

### V4.1-Flash — hardware reality [SECONDARY]
- The checkpoint is **510 GB across 48 shards** (476 GiB); V4 Flash was 166.9 GB on two H200s — the replacement is **3× the size**, and day-one recipes from vLLM, SGLang, and NVIDIA start at **four Blackwell-class GPUs or eight H200s**. "The model's clever memory tricks are real, but they're on the cache side. The weights are the problem." [SECONDARY]
- GPU memory floor: **~614 GB with headroom** (vLLM recipe). Precision as shipped: **FP8 dense, FP4 routed experts, FP8 Engram** — "the easy compression is already spent," so no "just quantize it" path; GPTQ/AWQ re-quantization is neither necessary nor recommended. [SECONDARY]
- Engram is **~183–189 GiB** of the total and **not optional** — every token touches it; SGLang can move it to host RAM, the one route to fewer GPUs. [SECONDARY]
- Context: **1,048,576 tokens**; max output 384K; DSpark built in (5 draft tokens/step in recipes); text+image input (day-one recipes serve text only). [SECONDARY]
- **Baseten** added V4.1-Flash to its Model APIs on **2026-09-11** — the day after launch — describing it as DeepSeek's **third open-weight flash release of 2026** and "the only model of its scale" using CED; support for Baseten's Loops training product was announced as coming soon. [SECONDARY]
- **First non-experimental native image input**: V4.1-Flash is DeepSeek's first non-experimental model with native image understanding — previously limited to the experimental **V4-Flash-Vision-Exp**. Card figures: **Chartography 78.9 vs 64.3**, **ZeroBench 49 vs 35** (new vs experimental). [SECONDARY]
- Launch-day cross-lab comparisons (vendor card): **GPQA Diamond 90.9**, **Codeforces 3471**, **HLE with tools 63.9**; **CyberGym 88.1 vs 84.5 for GPT-5.6 Sol and GLM 5.3** (per one secondary write-up). [SECONDARY]
- The size dispute, reconciled: one community analysis (r/LocalLLaMA, cited by OrcaRouter) reads the checkpoint as **748B total, ~510 GB in FP8** — the reconciliation offered is that DeepSeek's "552B" counts only the compute backbone while the download also carries a large sparsely-accessed memory table looked up rather than computed through. Treat 552B as vendor-reported, 748B as an unrefuted community reading. [COMMUNITY]

### V4 Pro retirement — 2026-09-14 [SECONDARY]
- DeepSeek retired the ~4×-more-expensive **V4 Pro on 2026-09-14**, telling API users: "V4.1 Flash has comprehensively surpassed V4 Pro across all key metrics, including performance, cost, speed, and task completion time." Legacy **V4 Flash and V4 Pro requests are being routed to V4.1 Flash at Flash pricing** — the corpus must treat `deepseek-chat`/`deepseek-reasoner` traffic as landing on V4.1-Flash after this date. [SECONDARY]
- Migration mechanics: from **04:00 UTC on 2026-09-14**, the `deepseek-v4-pro` endpoint stopped serving V4 Pro and started serving V4.1-Flash at Flash prices — DeepSeek said this holds until a future V4.1 Pro arrives. [SECONDARY]
- The migration table (one secondary source): V4 Pro input ~$0.435 / output ~$0.87 → V4.1-Flash off-peak $0.15/$0.60; context 1M both; **HLE (no tools) 42.7 → 36.8** — a regression; V4.1-Flash is weaker on unassisted HLE even as it wins on tool-using HLE (63.9). [SECONDARY]
- Off-peak output price cut vs V4-Pro: **~70%** ($1.98 → $0.60 per million). [SECONDARY]
- The operational critique (one digest): the silent migration means any prompt tuned against Pro now runs against a different model — "your request semantics, latency profile, and output distribution changed without a config change on your side." The source does not report measured throughput or quality deltas against Pro; treat benchmark claims as vendor statements until independent evals land. [SECONDARY]
- The framing from one commentary: "The 'budget' model got good enough that the 'premium' model no longer earns its markup" — V4.1-Flash (284B→552B, nearly double the previous Flash) is "not small"; the flash label now means serving economics, not model size. [SECONDARY]

### V4.1-Flash pricing — peak/off-peak refined [SECONDARY]
- Off-peak: **$0.15 input / $0.60 output / $0.003 cached input** per million. Peak (Mon–Fri, **01:00–04:00 and 06:00–10:00 UTC**): **$0.30 / $1.20 / $0.006** — exactly doubled. **Weekends are always off-peak.** [SECONDARY]

### Alias retirement — cutoff, then compatibility routing [SECONDARY]
- Official cutoff: `deepseek-chat` and `deepseek-reasoner` aliases retired **2026-07-24 15:59 UTC**. [SECONDARY]
- Third-party observations: **HTTP 400 on 2026-07-25** (aliases dead), then **HTTP 200 compatibility routing on 2026-07-28** — but `/models` still omitted the aliases. Treat post-07-28 behavior as **unsupported compatibility**, not a reversal of retirement: the retirement stands, the routing is best-effort. [SECONDARY]

