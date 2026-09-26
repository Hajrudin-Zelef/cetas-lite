---
id: ai-industry-kb-2026/02-open-weight-model-chronology/overview
title: "2. Open-Weight Model Chronology"
domain: open-weight-model-chronology
role: deep-dive
task: model-release
actors: ["AMD", "Alibaba", "Anthropic", "China", "DeepSeek", "EU", "Google", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "United States", "Z.ai"]
dates: ["2025-04", "2025-04-05", "2025-04-29", "2025-06", "2026-02", "2026-02-16", "2026-04", "2026-04-24", "2026-05", "2026-06-03", "2026-06-16", "2026-07-09", "2026-07-16", "2026-07-17", "2026-07-27", "2026-08-05", "2026-08-10", "2026-08-12", "2026-08-14", "2026-09", "2026-09-02", "2026-09-10", "2026-09-22"]
keywords: ["open-weight", "agent", "agentic", "amd", "apache", "benchmark", "benchmarks", "claude", "cost", "cyber", "deepseek", "diffusion"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [649, 678]
section: "2. Open-Weight Model Chronology"
sha256: d2823e1623b1db626932dbb5bfd63016962b2a271329581c57eaaf63d6b5ae18
---

# 2. Open-Weight Model Chronology
Keywords: llama 4, muse glimmer 30b, qwen3.5, qwen3.6, qwen3.7, qwen3.8, deepseek v4, deepseek v4.1-flash, glm-5.2, glm-5.3, kimi k2.6, kimi k3, mistral small 4, mistral large 3, mistral medium 3.5, gemma 4 12b, mimo-v2.6, apache 2.0 license, mit license, open-weight gap, artificial analysis index, mixture of experts, gated deltanet, unsloth dynamic quantization, 2026 release chronology

## Summary

This section is a chronology of open-weight model releases from February 2026 through September 22, 2026, consolidated from Wave 3's chronology verification, Wave 2.1's September delta, Wave 3.1's February-timeline and Mistral/Meta fill-in research, and Wave 1's frontier-vs-open-weight gap measurement. Deep-dive material (DeepSeek architecture, GLM/Kimi/MiniMax internals, Qwen architecture) lives in §3–5; inference engines in §6; quantization in §8; pricing in §14 — this section keeps 2–4 lines per model, dated facts only.

The 2026 open-weight year has a clear arc: Alibaba opened the year with Qwen3.5-397B-A17B (2026-02-16/17, Apache 2.0), DeepSeek's V4 preview (2026-04-24, MIT) reset the frontier-cost baseline, Z.ai's GLM-5.2 (2026-06-16) briefly held the top independent open-weight score, Moonshot's Kimi K3 (2026-07-16 launch, weights ~2026-07-27) pushed open weights into the 3T class at 2.8T parameters, Alibaba's Qwen3.8-Max weights (2026-08-12) made a Max-tier flagship downloadable for the first time, and DeepSeek's V4.1-Flash (2026-09-10) closed the window with a new Causal Encoder-Decoder architecture while retiring V4-Pro from the API.

Three corrections from earlier waves are load-bearing for the whole knowledge base and must not regress. First, **Meta Llama 4 Scout/Maverick launched 2025-04-05, not April 2026** — the April-2026 dating was off by exactly one year; Behemoth (~2T total, 288B active) was announced alongside them, postponed twice, and remains unreleased, effectively shelved (never formally cancelled). Second, **Qwen3-30B-A3B is a 2025-04-29 SKU, not a February 2026 release** — the mid-February 2026 Alibaba release was Qwen3.5-397B-A17B; the brief attached the wrong model name to the right week. Third, **GLM-5.2 is 753B total / ~40B active** — the "744B" figure is the FP8-build/VRAM shorthand, not the architectural total.

Two additional dates were corrected in verification: Gemma 4 12B was announced 2026-06-03 (not at Google I/O 2026 on May 19), and Qwen3.8-27B's weights shipped 2026-08-14 (not August 12 — that date belongs to Qwen3.8-Max's weights, which carry a restrictive custom license, unlike the Apache-2.0 27B sibling). Qwen3.8-Max's weights are text-only and thinking-only — not the hosted multimodal product.

License fragmentation is the second axis of the chronology. 2026 saw the mainstream open-weight lines converge on permissive OSI terms — Apache 2.0 (Qwen3.5/3.6/3.8-27B, Gemma 4, Mistral Small 4/Large 3, Meta's Muse Glimmer 30B — Meta's first straight Apache-2.0 model) and MIT (DeepSeek V4 family, GLM-5.3-Flash, MiMo-V2.6) — while the largest 2026 flagships carried gated terms: Qwen3.8-Max's bespoke restrictive license (100M MAU / $20M monthly revenue display requirement; separate commercial license for MaaS above $50M trailing revenue), Kimi K2.6/K3's Modified MIT (100M MAU / $20M monthly revenue display threshold), GLM-5.3's bespoke GLM-5.3 License, and Llama 4's Community License (700M MAU cutoff, EU exclusion for multimodal variants).

MoE is the architectural default for frontier-scale open-weight releases in 2026 (DeepSeek V4, GLM-5.x, Kimi K2.6/K3, Qwen3.5/3.6-35B-A3B, Llama 4), with 3–10% active-parameter ratios delivering frontier knowledge capacity at mid-size inference cost. Dense architectures persist at the small/efficient end: Qwen3.8-27B (Apache 2.0, AA Index 52, size-class leader) and Meta's Muse Glimmer 30B (dense 30B, Apache 2.0) are the counterexamples that break any unqualified "MoE everywhere" claim.

The measured open–closed gap as of September 2026: ~4 months / 8 ECI points (Epoch AI, May 2026 update), 4.4 months for Chinese open weights vs the US frontier (Mozilla State of Open Source AI, Sept 15, 2026), 29 LMArena Elo points (Sept 2026, Claude Opus 5 Max 1505 vs Kimi K3 trailing ~30), and 4–7 months on cyber benchmarks (AISI, July 17, 2026). On coding and agentic benchmarks, open weights outright win individual leaderboards (Kimi K3 #1 Frontend Code Arena; GLM-5.3 100% pass on independent Ed-o-meter).

## Key dated facts

### Meta — Llama 4 (2025 correction) and Muse Glimmer (2026)

- **2025-04-05 — Llama 4 Scout and Maverick released (CORRECTION: not April 2026).** Scout: 109B total / 17B active / 16 experts / 10M-token context; Maverick: 400B total / 17B active / 128 experts / 1M-token context. Early-fusion multimodal MoE under the Llama 4 Community License (free under 700M MAU; acceptable-use policy; EU exclusion for multimodal variants). Meta's first open MoE models; Maverick was co-distilled from Behemoth, Scout trained from scratch. As of September 2026, Maverick sits at 1352 on LMArena text — no longer competitive at the open-weight frontier.
- **Llama 4 Behemoth — announced 2025-04-05, never released, effectively shelved (never formally cancelled).** Previewed at ~2T total / ~288B active / 16 experts, "still training"; internal target slipped April 2025 → June 2025 → fall 2025 or later; Meta engineers reportedly questioned whether gains over predecessors justified shipping. Its April 2025 preview benchmark figures (MMLU Pro 82.2, LiveCodeBench 49.4, MMMU 76.1) are stale and must not be quoted as current. Do not list Behemoth as a 2026 open-weight model.
- **2026-08-10 — Meta Muse Glimmer 30B released.** Meta Superintelligence Labs' real 2026 open-weight release: dense 30B (29.6B by one count), distilled from closed Muse Spark via logit distillation, native multimodal (dedicated perception encoder; text+images; 100+ languages), 120K+ context, Apache 2.0 — Meta's first straight Apache-2.0 model (no MAU cutoff, no acceptable-use policy; no training data or code released, so "open weight" not "open source"). Bundled DFlash block-diffusion speculative drafter (16 tokens/forward pass; ~3.1× speedup on RTX 5090: 74.9 → 233.4 tok/s). K-Quant-17GB build (~17GB, 1.0% avg degradation across 15 benchmarks); reference hardware RTX 5090 / M5-M4 Max; community runs on 20GB AMD 7900XT and 32GB Mac Mini via Ollama. [VENDOR] Meta-reported: SWE-Bench Pro 51.2%, AIME 94.7, GPQA 83.5; MCP Atlas 75.5; independent AA Index 35 in its size class (behind Qwen3.8-27B's 52 and Qwen3.6-27B's 38). Muse Spark 1.2's open weights, announced for "soon" on release day, remain "future work" as of 2026-09-22.
- Muse Spark lineage for context (closed, §1-adjacent): original April 2026 → 1.1 July 9, 2026 → 1.2 August 5, 2026 (co-launched with Muse Code terminal agent; 1M context; $1.25/$4.25 per M tokens) → 1.3 September 2, 2026 (closed, via Muse Code and Meta Model API).

### Alibaba — the Qwen cadence (open vs proprietary tiers)

