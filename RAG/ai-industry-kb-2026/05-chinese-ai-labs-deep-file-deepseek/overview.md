---
id: ai-industry-kb-2026/05-chinese-ai-labs-deep-file-deepseek/overview
title: "5. Chinese AI Labs — Deep File: DeepSeek"
domain: chinese-ai-labs-deep-file-deepseek
role: deep-dive
task: actor-profile
actors: ["China", "DeepSeek", "Moonshot", "SGLang", "vLLM"]
dates: ["2026-02-15", "2026-03", "2026-03-09", "2026-04-24", "2026-06-01", "2026-06-02", "2026-07-24", "2026-08-12", "2026-08-13", "2026-08-16", "2026-09-08", "2026-09-10"]
keywords: ["deepseek", "apache", "attention", "compute", "cost", "decode", "fp4", "kimi", "kv cache", "license", "memory", "mit license"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1766, 1793]
section: "5. Chinese AI Labs — Deep File: DeepSeek"
sha256: 4d3aecc1667af789292a616d03b04e2884da90b547d69daa9650abb18b16c27f
---

# 5. Chinese AI Labs — Deep File: DeepSeek
Keywords: DeepSeek V4, DeepSeek-V4-Lite, DeepSeek V4-Pro-0813, DeepSeek V4.1-Flash, Causal Encoder-Decoder, CED, Compressed Sparse Attention 2, CSA2, Engram, MIT license, 552B MoE, 1.6T MoE, 1M context, Sealion-Lite, Healer Alpha, deepseek-flash, deepseek-chat, deepseek-reasoner, Vals Index, SWE-bench Verified 96.40%, peak off-peak pricing, stealth release, CSA, DSA, MoE fine-grained experts, MLA, FP4 KV cache, Aikido Security, beta to GA, routing aliases

## Summary

- DeepSeek's 2026 arc runs **rumor → leak → official preview → rapid checkpoints → new architecture family**: the first V4 mentions circulate around **2026-02-15** as a pure rumor phase (NOT a release); a **V4-Lite ("Sealion-Lite" / "Healer Alpha" / "0302")** surfaces on DeepSeek's website on **2026-03-09** as a leak/soft-preview (~200B parameters reported by 36kr, unconfirmed); the official **DeepSeek-V4 Preview** (V4-Pro 1.6T/49B active + V4-Flash 284B/13B active, 1M context, MIT weights) lands on **2026-04-24**; and **V4.1-Flash** (552B MoE, Causal Encoder-Decoder, CSA2) ships via a two-day beta **2026-09-08→09-10** to GA on **2026-09-10**.
- Three dates that prior briefs misread are corrected here with primary-adjacent evidence: **2026-06-01 was a permanent 75% flagship API price cut, not a V4.1 announcement** (Digitimes, June 2, 2026); **2026-08-13 was the V4-Pro-0813 GA checkpoint, a V4-line revision, not V4.1**; and the brief's **"+10.6 points on the Vals Index" is contradicted — Vals.ai's own dated evaluation records +9.48 points** (52.37% vs 42.89%).
- The most important terminology kill in the file: **"Compressed Expert Dispatch" does not exist**. CED is **Causal Encoder-Decoder** (20-layer causal encoder + 20-layer decoder; decoder KV projected from encoder final hidden states). A quoted-phrase web search for "Compressed Expert Dispatch" + DeepSeek returned zero relevant hits; it is plausibly an AI-generated expansion of the acronym produced during the rumor cycle. The consolidation draft must be grepped for it.
- Architecture facts split by evidence grade. For V4: DeepSeek-style fine-grained MoE (Flash: 284B/13B, 256 routed experts, 6 active; Pro: 1.6T/49B, 384 routed experts, 6 active; both +1 shared expert), first three blocks hash-routed, aux-loss-free balancing with sqrt(softplus) routing affinity, FP4 expert weights, hybrid CSA/HCA attention — **all secondary reconstructions, no directly verified official report**. For V4.1-Flash: 552B MoE with **Causal Encoder-Decoder**, **CSA2** (successor of V4's DSA sparse-attention lineage), FP4-E2M1 KV caching, **890 bytes/token global KV cache ≈ 1/4 of V4-Flash** (~75% reduction — the brief's "75% compression" maps to this vendor claim), 8B active prefill / 16B active decode, 1M context, 384K max output, MIT weights — all **[VENDOR]** via the model card and secondary coverage; **no independent reproduction exists yet**.
- **DSA (DeepSeek Sparse Attention)** belongs to V4's attention lineage and reduces **attention compute** — it must NOT be equated with KV-cache reduction. KV-cache reduction in the V4 line comes from the CSA/HCA hybrid (CSA ~4× compression, HCA ~128× compression of the KV); in V4.1 it comes from the CED structure plus FP4 KV caching and SWA Bounded Replay (persistent footprint ~1/8 of V4-Flash).
- **Engram** enters the story as a **March 2026 rumor term** (36kr via awesomeagents: a "conditional memory system" co-developed with Peking University, O(1) memory narrative, 97%-at-1M needle-in-a-haystack rumors). DeepSeek's official 58-page V4 technical report mentions "Engram" and "O(1)" **zero times** (buzzgrewal analysis). Engram reappears as a 196B-parameter conditional memory module in the V4.1 model card via secondary coverage — kept here as **[VENDOR]** with the rumor-provenance flag; do not present it as established DeepSeek technology.
- Serving economics are price-war material, kept brief in this file (full economics → §9/§14 of the final document): permanent 75% flagship cut from 2026-06-01 (V4-Pro list $0.435/$0.87 from that date); peak/off-peak pricing from 2026-08-16 (V4-Pro output $3.96 peak / $1.98 off-peak); V4.1-Flash GA pricing **$0.15 input / $0.60 output per 1M tokens off-peak, 2× peak**, effective 2026-09-10. DeepSeek claims V4.1-Flash "comprehensively surpassed V4 Pro across performance, cost, speed, and task completion time" [VENDOR] and moved to retire the bigger model in favor of the smaller, cheaper one — an unusual flagship-retirement inversion.
- The retirement pipeline is naming, not engines: `deepseek-chat` and `deepseek-reasoner` were always **routing aliases** (non-thinking / thinking modes of the same engine, resolving to `deepseek-v4-flash` through the preview period), created at the April 24 announcement and **retired 2026-07-24 at 15:59 UTC** exactly as announced. A September 10 notice planned `deepseek-v4-pro` → V4.1-Flash routing from September 14; APIMaster's updated coverage reports a **revised notice keeping V4 Pro service unchanged after September 14** — **the conflict is flagged and kept open, not resolved**.
- Third-party evaluation anchors: Vals.ai evaluated the pre-release V4-Pro-0813 on 2026-08-12 (Vals Index 52.37%, #18 at the time; SWE-bench Verified 96.40% #2 of 82 — the highest open-weight score, ahead of Kimi K3 at 93.40%); Artificial Analysis Intelligence Index at **53, +8 over the April preview** (v4.1-era scale); Aikido Security found V4-Pro-0813 led all tested models in vulnerability-detection count (poor precision); mid-September press reports **12th on the Vals Index** — the index recomputes live, so rank must always be paired with its date.
- Ecosystem day-0 for V4.1-Flash was unusually fast and informative: vLLM preview container tagged September 9 (a day *ahead* of GA), SGLang + Miles joint day-0 post September 10 (explicit about which architecture parts are expensive), Ollama listing within a day. Orcarouter's caution stands: day-0 "support" meant a preview image plus integration work, not a stable pip release anywhere.
- Release-cadence takeaway: April 24 (preview) → July 31 (Flash-0731) → August 13 (Pro-0813) → August 21 (Flash-Vision-Exp) → September 10 (V4.1-Flash). Intervals make V4.1-Flash the fastest structural update in DeepSeek's history and the first previewed through an expiring test model id (`deepseek-v4.1-flash-expires-on-0910`).
- License note for the RAG: March 2026 speculative coverage repeatedly expected **Apache 2.0** for the V4 family; the shipped license is **MIT** (V4, V4.1-Flash). Undated "Apache 2.0" claims for V4-family weights are rumor-phase residue and must not be carried forward.

### Claim-verdict table (wave 3.1 re-verification)

- Claim 1 — 2026-03-09 DeepSeek-V4-Lite (200B) released quietly: **VERIFIED as soft-launch/leak**; 200B unconfirmed [UNVERIFIED figure].
- Claim 2 — 2026-04-24 V4 Pro (1.6T/49B) + V4 Flash (284B/13B), official preview, MIT weights: **VERIFIED** (recap of wave3/02 Claim 1; MIT confirmed, correcting Apache 2.0 speculation).
- Claim 3 — 2026-07-24 definitive retirement of deepseek-chat / deepseek-reasoner: **VERIFIED** (alias retirement at 15:59 UTC; engine unchanged) — with the nuance that they were routing aliases, not models.
- Claim 4 — 2026-06-01 announced → 2026-08-13 V4.1 confirmed, 552B, "Compressed Expert Dispatch": **CONTRADICTED on dates and term** — June 1 was the 75% price cut; Aug 13 was the V4-Pro-0813 checkpoint; V4.1-Flash GA = 2026-09-10; CED = Causal Encoder-Decoder.
- Claim 5 — 2026-08-12 V4 Pro "0813" evaluated, +10.6 pts vs initial V4 on Vals Index: **VERIFIED WITH CORRECTION — +9.48 pts, not +10.6**.

## Key dated facts

