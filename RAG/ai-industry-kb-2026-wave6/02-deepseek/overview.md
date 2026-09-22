---
id: ai-industry-kb-2026-wave6/02-deepseek/overview
title: "§2. DeepSeek"
domain: deepseek
role: deep-dive
task: actor-profile
actors: ["Anthropic", "China", "DeepSeek", "Google", "Huawei", "OpenAI", "OpenRouter", "SGLang", "Xiaomi", "vLLM", "xAI"]
dates: ["2025-09-29", "2025-11", "2025-12", "2025-12-01", "2026-01", "2026-01-21", "2026-03", "2026-03-11", "2026-04", "2026-07-24", "2026-08-13", "2026-09", "2026-09-09", "2026-09-10", "2026-09-22"]
keywords: ["deepseek", "agentic", "alignment", "ascend", "attention", "attribution", "benchmark", "benchmarks", "claude", "compute", "context window", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [565, 636]
section: "§2. DeepSeek"
sha256: 7033f88ff9f42546e3370323649c3194104617f3876fae04f4b7178fe0505d28
---

# §2. DeepSeek

Keywords: deepseek v4.1 flash, v4-pro-0813, deepseek v4 pro, deepseek-chat alias retired, deepseek-reasoner alias retired, engram, conditional memory, causal encoder-decoder, ced, csa2, deepseek r2 unreleased, deepseek r2 not released, model1 github leak, hunter alpha xiaomi mimo, v4-lite rumor, peak pricing, off-peak pricing, 8b active prefill, 16b decode, 1m context, 384k output, 552b backbone, 196b engram, aime 92.7 fabrication, deepseek v4 chronology

## Summary

DeepSeek's 2026 story in this corpus is the V4 line's arc from leak-chasing to formal releases, plus the definitively non-released R2. Around **2026-01-21** a `MODEL1` GitHub identifier leaked; a **2026-03-11** "Hunter Alpha" speculation turned out to be Xiaomi's MiMo-V2-Pro, not DeepSeek; V4-Lite never shipped and survives only as community rumor and SEO noise. DeepSeek retired the `deepseek-chat`/`deepseek-reasoner` aliases on **2026-07-24 (15:59 UTC)**, formally released **V4-Pro-0813** on **2026-08-13**, and moved V4.1-Flash from beta to early access (**2026-09-09**) to GA (**2026-09-10**) with time-of-day peak/off-peak pricing. [SECONDARY]

Architecturally, V4.1-Flash pairs a 552B backbone with a 196B "Engram" conditional memory, activates 8B on prefill / 16B on decode, and runs 1M context with 384K output; **CED definitively means "Causal Encoder-Decoder"** (quoted from DeepSeek's own HF README — the "Compressed Expert Dispatch" expansion is rejected). DeepSeek **R2 had not released by 2026-09-22**; claims of an April 2026 32B MIT model scoring AIME 92.7% are `[UNVERIFIED]` fabrication-risk and must not be cited as a release. [VENDOR]

V4.1-Flash is also the corpus's cost-per-intelligence anchor: at **$0.27 per AA Index task** (Sept 2026) it was the cheapest frontier-class API measured — a dated price-card fact, not a permanent one. [SECONDARY]

## Key dated facts

### Leak season (January–March 2026)
- **2026-01-21** — A `MODEL1` GitHub leak surfaced ahead of the V4 window; DeepSeek did not attach it to any official release. [SECONDARY]
- Coverage of the leak (technode, dataconomy, techbriefly, all 2026-01-21) treated it as a V4 precursor — none of it was vendor-confirmed. [SECONDARY]
- **2026-03-11** — "Hunter Alpha" speculation circulated as a possible DeepSeek V4 sighting; it was later revealed to be **Xiaomi's MiMo-V2-Pro**, not a DeepSeek model — the episode is a correction, not DeepSeek news. [SECONDARY]
- **V4-Lite** never shipped; mentions survive only as `[COMMUNITY]` rumor and SEO-noise content. No release date, no weights, no vendor acknowledgement. [COMMUNITY]

### Alias retirement and V4-Pro-0813
- **2026-07-24, 15:59 UTC** — DeepSeek retired the `deepseek-chat` and `deepseek-reasoner` API aliases, forcing migration to versioned V4 models. [VENDOR]
- The retirement was a hard cutoff at 15:59 UTC for both aliases simultaneously — not a deprecation window — which is why the timestamp travels with the fact. [VENDOR]
- **2026-08-13** — DeepSeek formally released **V4-Pro-0813**, the pinned production checkpoint of the V4 Pro line. [VENDOR]
- Reception note: secondary coverage (scmp) described V4-Pro-0813 as struggling on general benchmarks but shining on cybersecurity — reception, not vendor claims. [SECONDARY]
- V4-Pro-0813 scores **36 on the Artificial Analysis Intelligence Index v4.2/v4.3-era figures** (wave2.1/05); V4.1-Flash scores **40 on v4.3** — both figures are version-pinned and must not be compared against scores from v4.1.1 or v4.2 without the methodology caveat. [SECONDARY]

### V4.1-Flash: beta → early access → GA
- **Early September 2026** — V4.1-Flash ran in beta. [VENDOR]
- **2026-09-09** — V4.1-Flash early access opened. [VENDOR]
- **2026-09-10** — V4.1-Flash reached **general availability** with the peak/off-peak pricing schedule below. [VENDOR]
- V4.1-Flash: **552B backbone + 196B Engram conditional memory**; **8B active on prefill / 16B on decode**; **1M-token context**; **384K-token output**; CSA2 attention; CED ("Causal Encoder-Decoder"). [VENDOR]
- Peak pricing belongs **specifically to V4.1-Flash**: peak **$0.30/$1.20** per M input/output; off-peak **$0.15/$0.60**; cache reads **$0.006 peak / $0.003 off-peak**. [VENDOR]
- The 196B Engram memory is a *conditional* memory layer on top of the 552B backbone — it is not simply "a 748B model." [VENDOR]
- The Engram technique was reported in secondary coverage as early as **January 2026** (TechWire Asia) — the architecture detail preceded the V4.1-Flash GA by eight months. [SECONDARY]
- At inference, V4.1-Flash activates 8B (prefill) / 16B (decode) — despite the 552B backbone it behaves like a small model per token, which is the basis of its peak/off-peak pricing. [VENDOR]
- 1M-token context is the steady-state spec shared across the V4 line (V4-Pro and V4.1-Flash) — the context figure does not vary by checkpoint. [VENDOR]

### CED resolution: Causal Encoder-Decoder
- DeepSeek's HF README for DeepSeek-V4.1-Flash states **CED = "Causal Encoder-Decoder"**; the alternative expansion "Compressed Expert Dispatch" circulating in community notes is **rejected** against the vendor's own text. [VENDOR]
- The resolution matters for downstream citations: any downstream file expanding CED as "Compressed Expert Dispatch" is propagating an error. [DIRECTIONAL]

### DeepSeek R2: not released (resolution)
- DeepSeek **R2 had not released by 2026-09-22** — no weights on HF, no API, no official acknowledgement of a shipping model. [SECONDARY]
- Claims of an **April 2026 32B MIT-licensed model with AIME 92.7%** are **[UNVERIFIED]** with fabrication risk; the corpus flags them as invented specificity and they must not be cited as a release anywhere. [UNVERIFIED]
- R2 remains relevant only as a *non-release*: a case study in how SEO-grade "release guides" manufacture detail for an unreleased model. [DIRECTIONAL]


### New verified facts — expansion

### V3.2 — December 2025 baseline, size contradiction resolved [SECONDARY]
- **2025-12-01** — DeepSeek V3.2 was released in December 2025, so it is the corpus's **baseline**, not a new 2026 model. Every 2026 number below is measured against it. [SECONDARY]
- **CONTRADICTION RESOLVED**: the 671B-vs-685B dispute now has clean attribution — **671B = DeepSeek's official V3.2 model card** (fe-static.deepseek.com); **685B = DeepLearning.ai's The Batch** (V3.2-Exp coverage). Both agree on **37B active per token**. Cite 671B as the vendor figure, 685B as the secondary figure. [VENDOR]
- Official card: V3.2 uses **exactly the same architecture as V3.2-Exp**; vs V3.1-Terminus the **only modification is DSA via continued training**. Text in/out; **128K context**; MIT license for open-repository assets. [VENDOR]
- **V3.2-Speciale**: gold-medal-level on **IMO 2025, Chinese MO 2025, IOI 2025**, competitive gold at **ICPC World Finals 2025**. Post-training: **GRPO at >10% of pre-training compute**; contest specialists' cases released for external verification. [SECONDARY]
- **First DeepSeek model with thinking integrated into tool use**: thinking/non-thinking tool modes; internal reasoning **persists across tool calls**, reset only on new user messages. [SECONDARY]
- DSA economics: near-linear **O(kL)** attention; **~50% lower long-context API cost** at V3.1-Terminus-level quality. [SECONDARY]
- V3.2-Exp training recipe: lightning indexer learned from **2.1B tokens**; all parameters fine-tuned on **~100B tokens**; **five specialists distilled** (reasoning, math, coding, agentic coding, agentic search). Inputs >7K tokens: **2–3× faster**. API **$0.28/$0.028/$0.42**; 128K in / 8K out. [SECONDARY]
- The V3.2-Exp code **supports Huawei AI chips** (other Chinese designers adapting) — the domestic-silicon hedge predates the R2 Ascend saga by months. [SECONDARY]

### V3.2-Exp — 2025-09-29, the sparse-attention experiment [SECONDARY]
- **2025-09-29** — DeepSeek released **V3.2-Exp** (TechCrunch): an experimental model whose single architectural change was **DeepSeek Sparse Attention** — a "lightning indexer" that prioritizes excerpts from the context window plus a "fine-grained token selection system" that picks specific tokens from those excerpts into a limited attention window. [SECONDARY]
- The research framing matters: one researcher called it "a controlled experiment release — the kind we need more of in AI," because DeepSeek **swapped one component (attention) and held everything else constant**, letting the community study sparsity's effect on large-scale training as a testbed. [SECONDARY]
- Pricing: **$0.028 per million input tokens** — a 50% cut from V3.1-Terminus — "even when approaching the context limit of 128,000 tokens." Community reaction: "half the cost for the same results," with developers reporting low latency and high throughput via OpenRouter, crediting the sparse kernels. [SECONDARY]
- The official **V3.2** followed on **2025-12-01** alongside a specialized variant **DeepSeek-V3.2-Speciale**; one syndicated launch write-up cites official V3.2 input pricing **as low as $0.07/M** — consistent with the $0.14/$0.28 vs $0.27/$0.40 dispute above being a tier/cache artifact, not a single number. [SECONDARY]
- Benchmark deltas vs V3.1-Terminus (independent review): **+0.9 on AIME 2025**, **+75 points on Codeforces** — the consolidated RL fine-tuning (merging reasoning, tool-use, alignment) gave V3.2-Exp a more balanced skillset. [SECONDARY]
- Open tooling at launch: **TileLang** kernels (readability), **DeepGEMM** (indexer logit kernels), **FlashMLA** (sparse attention CUDA kernels), plus **vLLM and SGLang support on day one** — cited as evidence DeepSeek was "cultivating broad community integration rather than locking down distribution." [SECONDARY]
- Self-hosting economics: the 685B model needs **~700 GB VRAM at FP8** (8–10× H100 80GB); 4-bit quantization brings it to **~386 GB** (5–6× H100). API price comparisons at launch: GPT-5 Standard $1.25/$10, Claude Opus 4.1 $15/$75, DeepSeek V3.2-Exp $0.028 input — a 45×–500× gap. [SECONDARY]
- Competitive context: V3.2-Exp shipped two months before a packed November 2025 window (GPT-5.1, Grok 4.1, Gemini 3 Pro, Claude Opus 4.5 within six days) — DeepSeek's December official release was read as proof Chinese labs could match frontier pace despite hardware restrictions. [SECONDARY]
- Vendor-chart results reported for **V3.2-Exp** (label as vendor claims, not independent): MMLU-Pro **85.0**, AIME 2025 **89.3**, GPQA-Diamond **79.9** (vs **80.7** for V3.1-Terminus), Codeforces **2121**, BrowseComp **40.1**. Note the inverted V3.2 < V3.1-Terminus GPQA result — one secondary read is that V3.1-Terminus was the stronger release on that axis, which tempers the "successor always wins" reading. [SECONDARY]
- **Pricing contradiction**: vendor/third-party listings conflict between **$0.14 input / $0.28 output** and **$0.27 input / $0.40 output** per million tokens. Likely explanations are first-party-vs-third-party pricing, cache-hit accounting, or date/cache mismatches — resolve only against DeepSeek's own pricing page or changelog. [SECONDARY]

