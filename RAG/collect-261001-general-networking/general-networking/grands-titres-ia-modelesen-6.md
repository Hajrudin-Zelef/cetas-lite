---
id: collect-261001-general-networking/general-networking/grands-titres-ia-modelesen-6
title: "VOLET 1 — Vague 1 (EN)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Hugging Face", "MiniMax", "Moonshot", "OpenAI", "OpenRouter", "Xiaomi", "Z.ai"]
dates: ["2025-07-11", "2026-05", "2026-06-01", "2026-06-12", "2026-07", "2026-07-16", "2026-07-19", "2026-08-03", "2026-09-02", "2026-09-18"]
keywords: ["agent", "agentic", "agents", "apache", "attention", "aws", "bedrock", "benchmark", "benchmarks", "claude", "compute", "cost"]
source: docs/RAG/collect-261001-general-networking/Grands titres IA modèlesEN.md
source_anchor: ""
source_lines: [436, 488]
sha256: 05a08a00d04a75a4b7fa9c8106e371500d0f1bad3003544733dac35e0ee65cf9
---

# VOLET 1 — Vague 1 (EN)

**Release timeline:**
- Kimi K2 (1.04T / 32.6B active, 128K ctx): July 11, 2025. K2-Instruct-0905 (256K): Sept 2025. K2 Thinking: Nov 2025. K2.5 (multimodal MoE, 1T/32B, 256K): Jan 27, 2026. K2.6: Apr 20, 2026. K2.7 Code: June 12, 2026.
- **Kimi K3: hosted service July 16, 2026; downloadable weights late July 2026 (~July 27); Amazon Bedrock GA September 18, 2026.**

**Architecture:** 2.8T total parameters — "first open model to reach 2.8T" per Moonshot. MoE "Stable LatentMoE": 16 of 896 experts active (~50B active). Innovations: Kimi Delta Attention + Attention Residuals (long-sequence/depth information flow); quantization-aware training from SFT onward (MXFP4 weights, MXFP8 activations). ~2.5× scaling-efficiency gain over K2. Native vision; 1M-token context.

**Benchmarks:**
- Artificial Analysis Intelligence Index: K3 (max) ~44 (v4.3.2) in Sept 2026 snapshots; earlier v4.1.1 readings ~59.7–60. Index-version changes make cross-version comparison invalid — note methodology version when citing.
- BenchAlign ~74.4 (per user brief; Chinese coding/agent composite where K3 frequently leads Chinese open models).
- Vendor claim: frontier-level across Moonshot's suite, consistently beating other tested open models; still trails proprietary Claude Fable 5 and GPT-5.6 Sol overall.
- K2.6: 80.2% SWE-bench Verified, 58.6% SWE-bench Pro (vendor).
- Known limitations (Moonshot): trained in "preserved thinking history" mode — unstable if harness drops thinking history; long-horizon bias can cause unexpected autonomous decisions on ambiguous instructions.

**License & price:** K2 series: modified MIT. K3: new "open weight" license — large MaaS businesses (>$20M/yr revenue) must sign a separate agreement; NOT traditional open source. Weights 1.56 TB on Hugging Face. Commercial access via OpenRouter, Moonshot API, and now AWS Bedrock (first open-weight model on Bedrock with explicit prompt caching; runs in Bedrock security/compliance boundary).

**Positioning:** Moonshot is the most open-weight-forward Chinese frontier lab and the only one repeatedly setting the open-model size ceiling (K2 → K2.5 → K3). K3 competes as the top Chinese open model on several composites but was overtaken by MiMo-V2.6-Pro (46) and Qwen3.8-Max-0902 (45) on AA Index v4.3 in Sept 2026.

---

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

