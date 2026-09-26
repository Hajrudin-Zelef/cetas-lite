---
id: ai-industry-kb-2026/02-open-weight-model-chronology/figures-and-metrics
title: "Figures and metrics"
domain: open-weight-model-chronology
role: deep-dive
task: model-release
actors: ["Alibaba", "DeepSeek", "EU", "Google", "Meta", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Samsung", "Unsloth", "Z.ai"]
dates: ["2025-04-05", "2025-04-29", "2025-08-05", "2026-02-16", "2026-03-16", "2026-03-23", "2026-04", "2026-04-02", "2026-04-16", "2026-04-20", "2026-04-24", "2026-04-29", "2026-06-03", "2026-06-16", "2026-07", "2026-07-02", "2026-07-16", "2026-07-31", "2026-08-10", "2026-08-12", "2026-08-13", "2026-08-14", "2026-08-19", "2026-08-26", "2026-08-28", "2026-08-31", "2026-09", "2026-09-02", "2026-09-10", "2026-09-18", "2026-09-21"]
keywords: ["apache", "attention", "compute", "decode", "deepseek", "fine-tuning", "fp8", "funding", "glm", "kimi", "license", "llama"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [732, 790]
section: "2. Open-Weight Model Chronology"
sha256: 8d04ecfe6592a99a8a8f755857a7a5b1b381009b8b467357b99dd160ca93f0fa
---

# Figures and metrics

- Unsloth's Dynamic Quantization 2.0 was current in Daniel Han's 2026 talks (LlamaCon 2025 session: 1.5× faster training, 50% less VRAM, 8× longer context on Llama 4 fine-tuning; July 2026 kernels/RL seminar) — **scope it "as of July 2026"**. **Dynamic v3.0 launched 2026-08-19** and supersedes 2.0; the vendor's headline speed/VRAM claims were only partly replicated (decode-speed improvements held up better than the most aggressive tail-quality claims) — do not repeat vendor numbers at face value. Quantization detail → §8.
- [UNVERIFIED] Unsloth's claimed presence at AI Engineer World's Fair 2026 (Moscone West, June 30 – July 2, 2026) — no session title, track, day, or schedule entry found; the "112 slides" figure has no source anywhere. Plausible from Han's 2024/2025 AIE workshop history, but plausible is not verified.
- [UNVERIFIED] Unsloth seed funding ~$40K with YC and Logan Kilpatrick as investors — no source; the only concrete third-party figure is ~$500K seed (Redpoint Ventures scout, Samsung NEXT) from an OSS investment scorecard (2026-03-23) — conflicting data point, not gospel. Founders Daniel and Michael Han; partners (not investors) Google, OpenAI, Meta, NVIDIA.

## Figures and metrics

### Spec table: all 2026-window open-weight releases (dense, dated)

| Release date | Model | Total / active params | Arch | Context | License |
|---|---|---|---|---|---|
| 2025-04-05 | Llama 4 Scout | 109B / 17B (16 experts) | MoE | 10M | Llama 4 Community |
| 2025-04-05 | Llama 4 Maverick | 400B / 17B (128 experts) | MoE | 1M | Llama 4 Community |
| 2025-04-29 | Qwen3-30B-A3B | 30.5B / 3.3B (128 experts, 8 active) | MoE | 32K (YaRN → 131K) | Apache 2.0 |
| 2025-08-05 | gpt-oss-120b | 117B / 5.1B | MoE, MXFP4 | 128K | Apache 2.0 |
| 2025-08-05 | gpt-oss-20b | 20.9B / 3.6B | MoE | 128K | Apache 2.0 |
| 2025-12 | Mistral Large 3 | 675B / 41B | MoE | 256K | Apache 2.0 |
| 2026-02-16/17 | Qwen3.5-397B-A17B | 397B / 17B (512 experts, top-K) | MoE, 3:1 Gated DeltaNet hybrid | 262K (YaRN → 1M) | Apache 2.0 |
| 2026-03-16 | Mistral Small 4 | 119B / 6B (128 experts, 4 active) | MoE | 256K | Apache 2.0 |
| 2026-04-02 | Gemma 4 family | 26B-A4B class | MoE | 256K | Apache 2.0 |
| 2026-04-16 | Qwen3.6-35B-A3B | 35B / 3B | MoE, 3:1 GDN hybrid | 262K | Apache 2.0 |
| 2026-04-20/21 | Kimi K2.6 | 1T / 32B (384 experts) | MoE | 256K/262K | Modified MIT |
| 2026-04-24 | DeepSeek V4-Pro (preview) | 1.6T / 49B | MoE, CSA/HCA | 1M | MIT |
| 2026-04-24 | DeepSeek V4-Flash (preview) | 284B / 13B | MoE, CSA/HCA | 1M | MIT |
| 2026-04-29/30 | Mistral Medium 3.5 | 128B dense (all active) | Dense | 256K | Modified MIT |
| 2026-06-03 | Gemma 4 12B | 12B dense | Dense, encoder-free multimodal | 256K | Apache 2.0 |
| 2026-06-16 | GLM-5.2 | 753B / ~40B (256 routed, 8+1 active) | MoE, IndexShare, MLA | 1M | MIT |
| 2026-07-16/27 | Kimi K3 | 2.8T (896 routed, 16 active) | MoE, KDA + Attention Residuals | 1M | Modified MIT |
| 2026-07-31 | DeepSeek V4-Flash-0731 | ~284–304B / ~13B | MoE | 1M | MIT |
| 2026-08-10 | Muse Glimmer 30B | 29.6–30B dense | Dense (distilled from Muse Spark) | 120K+ | Apache 2.0 |
| 2026-08-12 | Qwen3.8-Max weights | 2.4T / 95B (512 experts, 10+1 active) | MoE | 262K → 1M | Custom restrictive |
| 2026-08-13 | DeepSeek V4-Pro 0813 | 1.6T / 49B | MoE | 1M | MIT (Apr/Jul ckpts; 0813 weight status ambiguous) |
| 2026-08-14 | Qwen3.8-27B | 27.8B dense | Dense, multimodal | 262K → 1M | Apache 2.0 |
| 2026-08-26 | GLM-5.3-Flash | 320B / 18B | MoE | 1M | MIT |
| 2026-08-28/29 | GLM-5.3 | 753B / 40B (756GB native FP8, 141 files) | MoE, IndexShare | 1M | GLM-5.3 License |
| 2026-08-31 | DeepSeek V4-Flash-Vision-Exp | 305B (284B MoE + 32-layer vision) | MoE + vision encoder | 1M | MIT |
| 2026-09-02 | Qwen3.8-Max-0902 | 2.4T / 95B | MoE | 1M | Custom restrictive |
| 2026-09-10 | DeepSeek V4.1-Flash | 552B (8B prefill / 16B decode active) | Causal Encoder-Decoder, CSA2, Engram 196B | 1M (384K output) | MIT |
| 2026-09-21 | MiMo-V2.6 Flash | 309B / 15B | MoE | 1M | MIT |
| — | Llama 4 Behemoth | ~2T / 288B (16 experts) | MoE (announced, never released) | — | — |
| — | Qwen3.8-Omni-Flash (2026-09-18) | omni | API-only, not open-weight | 1M | — |

### Active-parameter ratios (the MoE economics of 2026)

- DeepSeek V4-Pro: 49B active / 1.6T total ≈ 3.0%. V4-Flash: 13B / 284B ≈ 4.6%. V4.1-Flash: 8–16B active / 552B ≈ 1.4–2.9%.
- GLM-5.2/5.3: ~40B / 753B ≈ 5.3%. GLM-5.3-Flash: 18B / 320B ≈ 5.6%.
- Kimi K2.6: 32B / 1T ≈ 3.2%. Kimi K3: 16 of 896 experts ≈ 1.8% routing (per-token compute far below the 2.8T total).
- Qwen3.5-397B-A17B: 17B / 397B ≈ 4.3%. Qwen3.6-35B-A3B: 3B / 35B ≈ 8.6%. Qwen3.8-Max: 95B / 2.4T ≈ 4.0%.
- Llama 4 Maverick: 17B / 400B ≈ 4.3%. Scout: 17B / 109B ≈ 15.6% (relatively high — the smaller-expert-count design).
- The consistent pattern: 3–10% activation ratios across the verified 2026 open releases — frontier knowledge capacity at mid-size per-token compute. No serious 2026 open release attempted frontier scale with a dense architecture.

### License tiers (September 2026 consolidation)

- **Tier 1 — OSI-permissive (Apache 2.0 / MIT), unrestricted commercial use:** Qwen3/3.5/3.6/3.8-27B (Apache 2.0); Gemma 4 (Apache 2.0); Mistral Small 4 / Large 3 (Apache 2.0); Muse Glimmer 30B (Apache 2.0); gpt-oss (Apache 2.0); DeepSeek V4 family incl. V4.1-Flash (MIT); GLM-5.3-Flash (MIT); MiMo-V2.6 (MIT). Practical rule: no legal review beyond standard open-source compliance.
- **Tier 2 — gated community/custom, free with conditions:** Llama 4 Community License (free under 700M MAU; acceptable-use policy; EU exclusion for multimodal variants — a hard blocker for European vision deployments); Kimi K2.6/K3 Modified MIT (display requirement above 100M MAU or $20M monthly revenue); Mistral Medium 3.5 modified MIT (free below $20M/month revenue); GLM-5.3 bespoke license.
- **Tier 3 — restrictive custom: read before you ship:** Qwen3.8-Max custom license (100M MAU / $20M monthly revenue display requirement; MaaS/AI-assistant businesses above $50M trailing revenue need a separate commercial license; internal use exempt). Note terms vary **within** families — Qwen3.8-27B is Apache 2.0 while Qwen3.8-Max is not; check the LICENSE file in the actual HF repository, per model, per release.
- Strategic trend (digitalapplied.com, April 2026): five of six major model families moved toward Apache 2.0 — converting mainstream open weights from "available with conditions" to "available unconditionally" — while the **largest 2026 flagships** (Qwen3.8-Max, Kimi K3, GLM-5.3) now carry the restrictive terms instead: the most capable weights have the most strings attached. EU AI Act overlay: genuinely open releases (MIT/Apache-2.0 with published details) carry a compliance advantage in the EU over both closed APIs and gated "open-weight-but-restricted" releases.

### Open–closed gap: the four independent measurements

