---
id: frontier-models-2026/09-rag-research-brief-volet1-wave-1-part-3-3/3-qwen3-8-max-alibaba-qwen3-x-family
title: "3. QWEN3.8 MAX (Alibaba) + Qwen3.x family"
domain: rag-research-brief-volet1-wave-1-part-3-3
role: deep-dive
task: actor-profile
actors: ["AWS", "Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Hugging Face", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "United States", "Xiaomi", "Z.ai", "xAI"]
dates: ["2026-05", "2026-06", "2026-06-01", "2026-07-16", "2026-07-19", "2026-08-03", "2026-09-02", "2026-09-10", "2026-09-21"]
keywords: ["agent", "agentic", "agents", "apache", "astra", "attention", "bedrock", "benchmark", "benchmarks", "claude", "compute", "cost"]
source: docs/RAG/Grands titres IA modèlesEN.md
source_anchor: ""
source_lines: [455, 510]
section: "RAG Research Brief — VOLET1 / Wave 1 (Part 3/3)"
sha256: 5e088108993e5f3699a18d5411f622bf57b4d8684b69edc3ae2d05b849bcc70e
---

# 3. QWEN3.8 MAX (Alibaba) + Qwen3.x family

## 3. QWEN3.8 MAX (Alibaba) + Qwen3.x family

**Release timeline:** Qwen3-Max (est. 1T): Sept 2025. Qwen3.6-Max-Preview (closed): Apr 2026. Qwen3.7-Max (closed): May 2026. **Qwen3.8-Max-Preview: July 19, 2026 (WAIC, no benchmark table). Qwen3.8-Max GA: August 3, 2026. Snapshot Qwen3.8-Max-0902: September 2, 2026** (post-training focused on coding/long-horizon agents).

**Architecture:** 2.4T total / ~95B active sparse MoE (Qwen3.5 architecture lineage). Alibaba's first multimodal model above 1T params. Native: 1M-token context, 131K max output, up to 262K reasoning tokens; text/image/video in, text out. Reasoning effort low/medium/xhigh (default xhigh). RL-environment scaling (~4,000 internal envs cited as generational-jump mechanism).

**Benchmarks:**
- AA Intelligence Index v4.3: 40 (Aug 3 GA) → **45 (0902 snapshot)** — reclaimed #1 Chinese-model spot mid-Sept, ahead of GLM-5.3 (44.9) and Kimi K3 (43.8), before MiMo-V2.6-Pro hit 46.
- Earlier v4.1.1: 58.1 (hosted) / 57.7 (open 2.4T-A95B weights) — effectively tied; vs Claude Opus 5 max 63.1, Fable 5 62.1, GPT-5.6 Sol 60.9, Kimi K3 max 59.7, GLM-5.3 max 59.5 on same version.
- Code Arena WebDev: 0902 snapshot ranked 1st (1,691), ahead of Claude Opus 5 Max.
- Caveat: 0902's intelligence gain came with verbosity inflation — reasoning tokens/task 44K → 71K, output 63K → 108K tokens; cost/task $2.67 → $5.41 at unchanged $2/$6 per M pricing.

**License & price:** API $2.00/M input / $6.00/M output; cache read $0.25 (10%), cache write $2.50; flat rate across full 1M context (no tiered long-context pricing, unlike Gemini 3.1 Pro). Open weights: PROMISED ("week of Aug 10" incl. Qwen3.8-27B) but not shipped as of research date; AA classifies the model proprietary. Note: do NOT confuse Qwen3.8-Max (2.4T flagship) with Qwen3-8B (8.2B dense Apache-2.0 model, Apr 2025) — search engines conflate them.

**Positioning:** Alibaba's flagship; strongest on web-dev/coding-agent tasks among Chinese models in Sept 2026; priced well above MiMo/DeepSeek Flash tiers — the "premium Chinese open-weight" play vs Xiaomi/Moonshot cost leadership.

---

## 4. MINIMAX M2.5 / M3

**MiniMax M3 — release June 1, 2026** (weights on Hugging Face by ~June 7–13; tech report arXiv June 11).
- Architecture: 428B total / ~22–23B active MoE, 128 experts (4 active); MiniMax Sparse Attention (MSA): GQA + sparse attention cutting per-token compute at 1M context to ~1/20 of prior gen; ~15.6× faster decoding, ~9.7× faster prefill at 1M. Native multimodal (text/image/video from step 0, ~100T interleaved training tokens) — structural edge over text-only DeepSeek V4, GLM-5.2, Qwen3.6. 1M context.
- Benchmarks: **SWE-bench Verified 80.5%**; SWE-bench Pro 59.0% (vendor — edged GPT-5.5's 58.6% at launch, first open-weight model to do so); Terminal-Bench 2.1 66.0%; MCP Atlas 74.2%; OSWorld-Verified 70.06%; LiveCodeBench 82.2; BrowseComp 83.5 (beats Claude Opus 4.7's 79.3). AA caveat: low attempt rate on AA-Omniscience (30.9% — low hallucination, low accuracy).
- Price: official $0.60/M in / $2.40/M out (promo $0.30/$1.20); Morph serving $0.255/$1.02 (256K). Among cheapest models above 80% on SWE-bench Verified.
- License: announced open-weight; license terms flagged as potentially restrictive for commercial use (confirm before self-host reliance).

**MiniMax M2.5:** prior generation; per user brief, M2.5/M3 class holds 70–75%+ SWE-bench Verified across snapshots; M2.5 established MiniMax in the agentic-coding tier before M3's June leap. (M2.5-specific independent figures are sparse in Sept-2026 coverage — treat the 70–75% band as snapshot-dependent, vendor/aggregator-reported.)

**Positioning:** M3 was "the first open-weight model to combine frontier coding + true 1M context + native multimodality." Directly collapses the cost/control case for closed APIs: ~1/30th the cost of Claude Fable 5-class, ~1/10th of Claude Opus 4.8 on coding workflows.

---

## 5. GLOBAL TREND — Chinese open-weights vs closed frontier (Sept 2026)

1. **Chinese labs own the open-weight coding/agent leaderboard.** Sept 2026 AA open rankings: MiMo-V2.6-Pro (46) > GLM-5.3 (45) > Qwen3.8-Max-0902 (45) > Kimi K3 (44) — all Chinese. SWE-bench Verified 78–80.6% cluster: DeepSeek V4 Pro 80.6%, MiniMax M3 80.5%, Kimi K2.6 80.2%, GLM-5.1 ~78%, MiMo-V2.5 Pro 78.0%.
2. **The efficiency frontier moved to Flash-tier models:** DeepSeek V4.1 Flash (Sept 10, 2026): 552B MoE (8B/16B active), 1M ctx, MIT, $0.14/$0.28 — near-frontier coding at the cheapest frontier-class price. MiMo-V2.6-Flash and GLM-5.3-FlashX play the same game.
3. **Closed models still lead on the hardest evals:** Claude Opus 5.x / Fable 5.1, GPT-5.x–6 (Sol/Astra), Gemini 3.x, Grok 4.x lead on Terminal-Bench 4.0 (Opus 5: 49.0 vs MiMo 34.9), ProgramBench competitive programming (37.0 vs 26.5), exploit-generation (ExploitBench 78.5 vs 47.9), and top AA Index absolute scores (Opus 5 max 63.1 vs best open ~58–60 on v4.1.1 scale).
4. **Cost asymmetry is the story:** open-weight Chinese models deliver ~90–95% of closed-frontier agentic coding at 1/10 to 1/60 the price (MiMo-Pro $0.13/task vs dollars-per-task closed equivalents; M3 at 1/30th of Fable 5-class).
5. **Transparency as strategy:** Xiaomi's live-streamed RL training (public dashboard, published costs, open-sourced RL tooling/environments) contrasts with US labs' closed training — reproducibility as competitive positioning. NVIDIA's Nemotron 3 Ultra (June 2026, 550B/55B, 71.9% Verified, free on OpenRouter) is the first Western open model partially breaking the Chinese monopoly in this tier.
6. **Caveats for RAG:** (a) AA Intelligence Index versions (v4.1.1 vs v4.3/v4.3.2) are NOT comparable across versions — always record the version; (b) vendor-reported SWE/DeepSWE figures are provisional until independently reproduced; (c) "open-weight" ≠ "open-source" — K3's MaaS revenue clause, Qwen3.8-Max's unshipped weights, and MiniMax's license terms all restrict the label.

---

## 6. Quick-reference table (Sept 2026 snapshots)

| Model | Release | Params (total/active) | Context | AA Index (v4.3) | SWE-bench Verified | License | API $/M (in/out) |
|---|---|---|---|---|---|---|---|
| MiMo-V2.6-Pro (Xiaomi) | 2026-09-21 | 1.02T / 42B | 1M, omni | 46 | n/r (DeepSWE 71.9) | MIT | 0.435 / 0.87 |
| MiMo-V2.6-Flash | 2026-09-21 | 309B / 15B | 1M, omni | — | n/r (DeepSWE 67.9) | MIT | 0.14 / 0.28 |
| Kimi K3 (Moonshot) | 2026-07-16 (wt late Jul) | 2.8T / ~50B | 1M | 44 | — (K2.6: 80.2%) | Open-weight* | via Bedrock/API |
| Qwen3.8-Max-0902 (Alibaba) | 2026-09-02 (GA Aug 3) | 2.4T / 95B | 1M | 45 | — | Proprietary (wt promised) | 2.00 / 6.00 |
| MiniMax M3 | 2026-06-01 | 428B / 23B | 1M, native multi | — | 80.5% | Open-weight‡ | 0.60 / 2.40 |
| DeepSeek V4.1 Flash | 2026-09-10 | 552B / 8–16B | 1M | 39 | 79.0% | MIT | 0.14 / 0.28 |
| GLM-5.3 (Z.ai) | 2026 (mid) | — | 1M (5.2+) | 45 | ~78% (5.1) | MIT | — / 3.20 |

\* K3: "open weight" license with >$20M/yr MaaS separate-agreement clause. ‡ M3: confirm commercial terms.
