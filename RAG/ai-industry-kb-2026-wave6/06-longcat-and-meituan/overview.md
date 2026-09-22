---
id: ai-industry-kb-2026-wave6/06-longcat-and-meituan/overview
title: "§6. LongCat and Meituan"
domain: longcat-and-meituan
role: deep-dive
task: reference
actors: ["China", "DeepSeek", "Huawei", "LongCat", "Meituan", "Nvidia", "OpenRouter", "Z.ai"]
dates: ["2025-09-22", "2025-11", "2026-01-14", "2026-05-29", "2026-06-12", "2026-06-29", "2026-06-30", "2026-07-05", "2026-08-13", "2026-08-28"]
keywords: ["agent", "apache", "ascend", "attention", "compute", "cost", "decode", "deepseek", "disaggregated", "embedding", "glm", "inference"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2603, 2674]
section: "§6. LongCat and Meituan"
sha256: ed9ec550a5c279072be5f42a998a608f0197abdf998234d8eae5fe7203ef86e3
---

# §6. LongCat and Meituan

Keywords: longcat, longcat-2.0, longcat flash thinking, thinking-2601, longcat flash omni, owl alpha, meituan, 10.1 trillion tokens, 1.6t 48b, 1m context, 50000 domestic accelerators, no nvidia hardware, six-model sunset, may 29 2026, tencent hy4 preview, huawei openpangu 2.0, openpangu pro 505b, openpangu flash 92b, 512k context, apache 2.0 hy4, tencent hy community license, hy3 apache 2.0, chinese open-weight flagships

## Summary

Meituan's LongCat line in 2026 runs from service consolidation to a trillion-parameter open release. **LongCat-Flash-Thinking** first released **2025-09-22/23**; its **Thinking-2601** refresh is effective **2026-01-14** per Meituan's own changelog (public reporting also references Jan 16 — the changelog is canonical). **LongCat-Flash-Omni** (around November 2025): 560B/27B, 128K, MIT. Meituan sunset its six-model LongCat Flash service on **2026-05-29**. The **Owl Alpha** stealth episode — revealed by Meituan on **2026-06-29** — was reported at roughly **10.1 trillion tokens/month** (trillion, not billion). **LongCat-2.0** was revealed **2026-06-30**, with weights and inference code under MIT on **2026-07-05**: 1.6T total / ~48B active, 1M context. Training on 50,000+ domestic accelerators with no NVIDIA hardware is `[VENDOR]`-only and not independently verified. [VENDOR]

This section also carries the "Other 2026 Chinese open-weight flagships": **Tencent Hy4 Preview (2026-08-28)** — 770B/49B, >1M context, Apache 2.0; and **Huawei openPangu 2.0** — announced 2026-06-12, staged release from 2026-06-30, Pro 505B/18B and Flash 92B/6B, 512K — whose official license was not verified, so it is **not** labeled Apache/MIT. [VENDOR]

## Key dated facts

### LongCat-Flash-Thinking and Thinking-2601
- **2025-09-22/23** — LongCat-Flash-Thinking original release (LongCat chat platform model docs). [VENDOR]
- **2026-01-14** — Thinking-2601 effective date per Meituan's own changelog; public reporting also references January 16. The vendor changelog is treated as canonical here. [VENDOR]
- LongCat-Flash-Thinking-2601 is a checkpoint refresh of the Thinking line, not a new architecture generation. [VENDOR]

### LongCat-Flash-Omni (~November 2025)
- **~November 2025** — LongCat-Flash-Omni: 560B total / 27B active, 128K context, **MIT**; real-time audio-visual interaction. [VENDOR]
- Omni's real-time audio-visual interaction is the differentiator against the text-only Thinking line — Meituan runs modality-specialized lines, not one omni model. [VENDOR]

### The six-model service sunset
- **2026-05-29** — Meituan sunset the six-model LongCat Flash service, consolidating the LongCat offering ahead of the 2.0 reveal. [VENDOR]

### Owl Alpha: the stealth episode
- **2026-06-29** — Meituan revealed the **Owl Alpha** stealth episode: an unreleased checkpoint that had been serving traffic under a codename. [SECONDARY]
- Reported volume: approximately **10.1 trillion tokens/month** — trillion, not billion; the corpus flags the unit explicitly because secondary writeups garbled it. [SECONDARY]
- Secondary coverage characterized Owl Alpha as a stealth model that had been "quietly topping OpenRouter all along" — a reception framing, not a vendor claim; the rankings remain [VENDOR]/secondary as sourced. [SECONDARY]
- Owl Alpha's rankings and volume figures are `[VENDOR]`/secondary as sourced — treat as vendor-reported, not independently measured. [VENDOR]

### LongCat-2.0: reveal and MIT release
- **2026-06-30** — LongCat-2.0 revealed. [VENDOR]
- **2026-07-05** — LongCat-2.0 weights and inference code released under **MIT**. [VENDOR]
- 1.6T total / ~48B active MoE; **1M-token context**. [VENDOR]
- Training on **50,000+ domestic accelerators with no NVIDIA hardware** is a vendor claim only — **not independently verified**. [VENDOR]
- The 50K-chip domestic-training claim, if true, puts LongCat-2.0 alongside DeepSeek V4 (Ascend) and GLM-5.1 (Ascend 910B) as the 2026 domestic-silicon training set; the "if true" is load-bearing. [DIRECTIONAL]

### Other 2026 Chinese open-weight flagships (full treatment in §17)

- **Tencent Hy4 Preview** (released 2026-08-28; announced 2026-08-13): 770B total / 49B active, >1M context, **Apache 2.0** — vs Hy3 Preview under the Tencent Hy Community License. [VENDOR/SECONDARY]
- **Huawei openPangu 2.0** (announced 2026-06-12; staged release from 2026-06-30): Pro 505B/18B, Flash 92B/6B, 512K context — official license **not verified** (do not label Apache/MIT). [SECONDARY]


### New verified facts — expansion (continued — LongCat-2.0 retail and docs notes)

- The tamago-labs reviewer accessed LongCat-2.0 through **OpenCode at $5/month (now $10/month)** — a concrete data point on the aggregator-subscription pricing tier that sits between token packs and direct API [COMMUNITY] (medium.com/tamago-labs).
- The official model blog is **`https://longcat.ai/blog/longcat-2.0/`** — cited as the source for the tech recap in both marktechpost and tamago-labs coverage; the vendor's canonical technical narrative lives there [SECONDARY].
- The ChangeLog's enterprise entries point to dedicated **FAQ-Invoices** and **FAQ-Account Verification & Benefits** documents — the platform's billing documentation surface is expanding alongside the enterprise tier [VENDOR] (longcat.chat/platform/docs/ChangeLog.html).


### New verified facts — expansion (continued — LongCat-2.0 architecture deep-dive (marktechpost, cryptobriefing, mbrukman, tamago-labs))

### Architecture: how a 1.6T model stays cheap to run (marktechpost)
- **Zero-computation experts**: simple tokens (e.g. punctuation) route to a zero-computation expert and return unchanged; complex tokens engage more expert capacity. A **PID controller adjusts expert bias** to hold the average in range — producing the **33B–56B dynamic activation window** instead of a fixed cost [SECONDARY] (marktechpost.com).
- The MoE backbone uses a **shortcut-connected design (ScMoE)** for higher throughput [SECONDARY] (marktechpost.com).
- **LongCat Sparse Attention (LSA)** is described as an **evolution of DeepSeek Sparse Attention (DSA)** — standard attention's quadratic scaling is dropped closer to linear via three orthogonal indexing methods [SECONDARY] (marktechpost.com):
  1. **Streaming-aware Indexing**: turns fragmented memory reads into contiguous blocks.
  2. **Cross-Layer Indexing**: reuses attention saliency across adjacent layers.
  3. **Hierarchical Indexing**: coarse-to-fine two-stage filtering.
  Together they sustain the 1M-token window "without a memory wall" [SECONDARY].
- **N-gram Embedding: a 135B-parameter module** sitting orthogonal to the MoE experts in sparse dimensions; captures dense local token relationships and reduces memory I/O during large-batch decoding [SECONDARY] (marktechpost.com).
- **Post-training (MOPD)**: a dedicated pipeline fusing **three teacher expert groups — Agent, Reasoning, Interaction** — into one unified model [SECONDARY] (marktechpost.com).
- Serving stack: **6D parallelism**, **prefill-decode disaggregated architecture**, **"super kernels"**, and **L2-cache weight prefetching** to hide I/O latency [SECONDARY] (marktechpost.com).
- The stability claim is framed as mattering **on non-Nvidia hardware, where tooling is less mature** [SECONDARY] (marktechpost.com).

### Training infrastructure (mbrukman GitHub technical notes)
- **Determinism & reliability**: enforced determinism for training reproducibility; numerical reliability via optimized foundational operators; **automated monitoring for seamless fault recovery** to secure stable production operations [SECONDARY] (github.com/mbrukman/longcat-2.0).
- **Training at scale**: 6D parallelism integrated with **super-node architectures**; multi-dimensional memory optimizations; **pioneering large-scale deployment of a customized Muon optimizer** [SECONDARY].
- **Long-context training**: optimized in-house operators; context scaled to 1M via an **all-gather-based CP (context-parallelism) scheme**; compute-communication overlap to minimize synchronization overhead [SECONDARY].
- Inference optimization on domestic superpod accelerators: **indexer pipelining and KV-cache parallelism** to mitigate KV-cache overhead; **explicit per-core control** for fully parallel dense+MoE execution; super kernels; L2 weight prefetching; high-speed interconnects for scale-up/out [SECONDARY].
- Serving: prefill-decode disaggregation with tailored schemes — **CPP and SP for prefill, KVP and large EP for decode** — plus **asynchronous load balancing** against stage-specific bottlenecks [SECONDARY].

