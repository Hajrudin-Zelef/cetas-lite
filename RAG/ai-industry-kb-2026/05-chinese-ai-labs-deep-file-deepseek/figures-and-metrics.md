---
id: ai-industry-kb-2026/05-chinese-ai-labs-deep-file-deepseek/figures-and-metrics
title: "Figures and metrics"
domain: chinese-ai-labs-deep-file-deepseek
role: deep-dive
task: actor-profile
actors: ["CISA", "China", "DeepSeek", "Hugging Face", "SGLang", "vLLM"]
dates: ["2024-12-26", "2025-01-20", "2025-03-25", "2025-05-28", "2025-08-21", "2025-09-22", "2025-09-29", "2025-12-01", "2026-02-11", "2026-02-15", "2026-03-09", "2026-04-24", "2026-06", "2026-06-01", "2026-06-02", "2026-07-24", "2026-07-31", "2026-08-12", "2026-08-13", "2026-08-16", "2026-08-21", "2026-09-08", "2026-09-09", "2026-09-10", "2026-09-14", "2026-09-22"]
keywords: ["agent", "apache", "attention", "compute", "deepseek", "fp4", "fp8", "kv cache", "license", "memory", "moe", "open weights"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1794, 1831]
section: "5. Chinese AI Labs — Deep File: DeepSeek"
sha256: cc729cd4ac7ef4c6169554b2c24688a262c3e8616b95d74220a035997be2d53f
---

# Figures and metrics

- 2024-12-26 — DeepSeek-V3 GA (prior lineage anchor). 2025-01-20 — DeepSeek-R1 GA. 2025-03-25 — V3-0324 checkpoint. 2025-05-28 — R1-0528 checkpoint. 2025-08-21 — V3.1 GA. 2025-09-22 — V3.1-Terminus checkpoint. 2025-09-29 — V3.2-Exp. 2025-12-01 — V3.2 GA.
- 2026-02-11 — Silent production upgrade: V3.2-class model context 128K → 1M tokens (community-detected) [COMMUNITY]. See Timeline for the quiet-rollout pattern.
- ~2026-02-15 — **First V4 mentions/rumors circulate**; pre-reporting/rumor phase — **explicitly NOT a release**, no model, no weights, no pricing. Multiple release windows (mid-February, Lunar New Year, early March) pass with no launch (nxcode FAQ).
- 2026-03-09 — **DeepSeek-V4-Lite ("Sealion-Lite" / "Healer Alpha" / "0302")** surfaces on DeepSeek's website as leak/soft-preview under NDA framing; ~200B params **reported by 36kr only, unconfirmed** [UNVERIFIED figure; secondary]. 36kr's report also says V4 Lite does NOT use the Engram conditional memory system.
- 2026-04-24 — **DeepSeek-V4 Preview (official)**: V4-Pro (1.6T total / 49B active) + V4-Flash (284B / 13B active); true 1M context; 384K max output; **MIT** open weights on HF/ModelScope; DeepSeek API changelog entry also announces `deepseek-chat`/`deepseek-reasoner` retirement in 3 months. Early speculation had expected Apache 2.0 — the shipped license is MIT.
- 2026-06-01 — **Permanent 75% flagship API price cut** announced (Digitimes, June 2, 2026); promo pricing ran through May 31; **NOT a V4.1 announcement**. V4-Pro list becomes $0.435/$0.87 (roninforge backfill).
- 2026-07-24 15:59 UTC — **`deepseek-chat` / `deepseek-reasoner` routing aliases retired** (aliases of `deepseek-v4-flash` non-thinking/thinking modes, resolved under the hood since April 24); no engine change. No grace alias or soft redirect announced.
- 2026-07-31 — **V4-Flash-0731**: re-post-trained checkpoint, same architecture (DeepSeek changelog).
- 2026-08-12 — Vals.ai evaluates pre-release **V4-Pro-0813**: Vals Index 52.37% (#18), **+9.48 pts vs V4 (42.89%)** — the brief's "+10.6" is contradicted. AA Intelligence Index 53 (+8 over April preview, v4.1-era scale).
- 2026-08-13 — **V4-Pro-0813 GA checkpoint (stealth)**: brief website statement on "significantly enhanced agent capabilities," removed by Thursday afternoon; no tech blog. **A V4-line revision, NOT V4.1.**
- 2026-08-16 — Peak/off-peak API pricing takes effect (wave2.1/06, techtimes): V4-Pro output $3.96 peak / $1.98 off-peak (355%/128% above the old $0.87 flat); V4-Flash output $1.32/$0.66; cache-hit input up to +1,100%.
- 2026-08-21 — **V4-Flash-Vision-Exp** image-input preview (secondary date dispute: Aug 21 vs Aug 31; immaterial, attribute per-source).
- 2026-09-08 — **V4.1-Flash two-day public beta**: <100-word notice in DeepSeek's official user group with model id `deepseek-v4.1-flash-expires-on-0910` (~12:00 PM); probe endpoint, not product.
- 2026-09-10 12:00 Beijing / 04:00 UTC — **V4.1-Flash GA**: 552B MoE, **Causal Encoder-Decoder (CED)**, CSA2, 890 B/token KV, native vision, MIT weights, model id `deepseek-flash`; pricing $0.15/$0.60 off-peak live immediately [VENDOR].
- 2026-09-14 (planned, conflicted) — DeepSeek's Sept-10 notice: all `deepseek-v4-pro` requests route to V4.1-Flash at Flash billing from 12:00 Beijing; **APIMaster's revised coverage reports V4 Pro API service continues unchanged after Sept 14, no routing** — both kept, conflict flagged.
- 2026-09-09 — vLLM ships a dedicated V4.1-Flash preview container — one day *ahead* of GA (orcarouter.ai). SGLang + Miles joint day-0 post September 10; Ollama lists `deepseek-v4.1-flash` within a day. Day-0 "support" = preview image + integration work, not stable pip releases.
- 2026-09-22 — Consolidation cutoff: no newer DeepSeek model release found in sources consulted; next milestone per the Sept-10 notice is V4.1-Pro (not yet shipped).

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

