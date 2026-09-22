---
id: ai-industry-kb-2026/02-open-weight-model-chronology/figures-and-metrics
title: "Figures and metrics"
domain: open-weight-model-chronology
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "EU", "Hugging Face", "Meta", "Microsoft", "Mistral", "Moonshot", "OpenAI", "United States", "Z.ai"]
dates: ["2025-01", "2025-04-05", "2025-04-29", "2025-08-05", "2025-10", "2026-02-16", "2026-03", "2026-03-16", "2026-04", "2026-04-02", "2026-04-16", "2026-04-20", "2026-04-24", "2026-04-29", "2026-05-29", "2026-06-03", "2026-06-16", "2026-07-16", "2026-07-17", "2026-07-30", "2026-07-31", "2026-08-10", "2026-08-12", "2026-08-13", "2026-08-14", "2026-08-26", "2026-08-28", "2026-08-31", "2026-09", "2026-09-02", "2026-09-04", "2026-09-10", "2026-09-15", "2026-09-18", "2026-09-21"]
keywords: ["agentic", "agents", "apache", "astra", "attention", "benchmark", "benchmarks", "claude", "compute", "cost", "cyber", "decode"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [736, 819]
section: "2. Open-Weight Model Chronology"
sha256: 357577383eedeba12a591eb2c866335541ec8ad5d8d0a4c773255974318d0bf3
---

# Figures and metrics

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

- **Epoch AI Capabilities Index:** average open-weight lag ~4 months / 8 ECI points (May 29, 2026 update; 3.5 months / 7 points in October 2025 — the gap widened slightly as GPT-5.x pulled ahead; convergence is directional, not monotonic). Six months under a stricter criterion.
- **Mozilla State of Open Source AI, 2nd ed. (September 15, 2026):** Chinese open-weight models 4.4 months behind the US frontier (read off the Artificial Analysis Intelligence Index; Kimi K3 within 3 index points of Anthropic's Fable 5 at ~30% of the cost). Mozilla's practical conclusion: most organizations should use open models as the default for the majority of work — paying for the frontier buys "a four-month head start at five times the cost."
- **Arena AI crowdsourced Elo (September 2026):** closed–open gap 29 points — Claude Opus 5 Max 1505 vs Kimi K3 trailing ~30; in January 2025 the gap had briefly hit zero (parity). Human-preference Elo rewards polish/instruction-following/reliability, favoring closed models more than task benchmarks do.
- **AISI open-weight gap report (July 17, 2026):** 4–7 months on cyber benchmarks — 4 months on narrow tasks, up to 7 on autonomous multi-step attack ranges. Once weights are distributed they cannot be recalled; DeepSeek V4-Pro's refusals on cyber tasks were overcome by simply retrying failed requests.
- **Where the gap remains vs. where it closed (Sept 2026):** coding benchmarks — closed or open-weight leads (Kimi K3 #1 Frontend Code Arena; GLM-5.3 #1 on Ed-o-meter); agentic computer use — near-parity (Qwen3.8-Max 86.1 vs Fable 5's 85.0 on OSWorld-Verified [VENDOR]); hardest reasoning — closed leads (Fable 5 HLE 53.3 vs Qwen3.8-Max 43.6); long-horizon autonomous tasks — closed leads; video/interleaved multimodality — closed leads; 1M-context coherence — closed leads (Claude claimed 76% at 1M, independently unverified); adversarial robustness — closed leads decisively.

### Artificial Analysis Intelligence Index — independent open-weight ranking snapshots

- **v4.1 scale (summer 2026):** GLM-5.3 60 (9th of 186/187 models, alongside proprietary frontier); Kimi K3 57 (July 30, 2026 — #1 among open-weight models at the time); V4-Flash-0731 50 (3rd among open weights); Muse Spark 1.2 57 (vendor-adjacent) vs 54 (independent evaluation — methodology difference, flag not average); GLM-5.2 51 (at release, open-model lead).
- **v4.3 scale (rebased September 4, 2026 — GPQA-Diamond dropped as solved, held-out evals added, 40% private-test weighting; absolute scores fell and prior-version numbers are no longer comparable):** GLM-5.3 ≈ 45, Kimi K3 ≈ 44, GLM-5.3-Flash 42, Qwen3.8 (2.4T A95B) 40, DeepSeek V4 Pro 0813 36 — vs closed frontier Claude Fable 5.1 = GPT-6 Astra = 53. Qwen3.8-Max is 4th among open weights, not the leader ("dominance" is vendor-reported and benchmark-specific).
- **Size class (v4.3-era, independent):** Qwen3.8-27B 52 (size-class leader), Qwen3.6-27B 38, Muse Glimmer 30B 35.

### Hugging Face distribution scale (2026)

- 2.4M+ public models, 730k+ datasets, ~1M Spaces, 50,000+ organizations, 11M users (Hugging Face's March 2026 DOE filing). The second million models landed in roughly a third of the time the first took.
- Transformers at 113M+ monthly downloads (late 2025); gpt-oss-120b at 4.3M+ downloads — among the most-downloaded open-weight models ever.
- Every major 2026 open-weight release lands on HF within days: Qwen3.8-Max (Aug 12), Kimi K3 (July 27 — "appeared on Hugging Face at 2:00 AM Beijing time"), DeepSeek V4 family, GLM-5.3-Flash (Aug 26).
- Curation signal vs raw count: Epoch tracks 1,340 open-weight models (964 language models); ATOM ~1,500 mainline open language models; 95 of 170 models on the AA Intelligence Index are open-weight — against 2.4M+ total HF models, model *selection* is now the hard problem.

### Key vendor-reported benchmark points (all self-reported unless marked independent)

- **Qwen3.8-Max (Alibaba, self-reported):** Terminal-Bench 2.1 86.6, PaperBench 93.0, GPQA Diamond 92.6, HLE 43.6, OSWorld-Verified 86.1, Agents' Last Exam 52.4, DeepSWE 1.1 56.6 (vs 21.6 for Qwen3.7-Max), ERQA 77.8 (Fable 5: 70.0), MobileWorld 77.8, JobBench 53.4 (vs 31.3).
- **Kimi K3 (Moonshot, self-reported unless marked):** Frontend Code Arena 1,679 Elo — #1 (independent, crowdsourced); Terminal-Bench 2.1 88.3; FrontierSWE 81.2; SWE-Bench Verified 76.8%; GPQA-Diamond 93.5; AIME 2025 96.1%; BrowseComp 91.2; HLE-Full 43.5; GDPval-AA v2 1687 vs Opus 4.8's 1600.
- **DeepSeek V4 family (DeepSeek, self-reported):** LiveCodeBench 93.5, GPQA Diamond 90.1, SWE-Verified ~80, Codeforces 3206; V4-Flash-0731 Terminal-Bench 2.1 82.7 (up from 61.8 for the April preview); DSA attention: ~27% inference compute and 10% KV cache vs V3.2 at the same length.
- **DeepSeek V4.1-Flash (DeepSeek, self-reported):** GPQA Diamond 90.9, Terminal-Bench 2.1 90.6, Codeforces 3471, DeepSWE v1.1 74.2% (harness swing 65.5–74.2%); concedes Terminal-Bench 3.0 30.0 / 4.0 31.2 / HLE 36.8.
- **GLM-5.3 (Z.ai, self-reported):** Terminal-Bench 3.0 28.3 (from 4.6), DeepSWE v1.1 66.9, Z.ai Code Bench 34.5% at ~75K output tokens (vs Opus 4.8's 29.5% at ~120K), CyberGym 84.5%, ExploitBench 54.4% (from 24.4%); 2,436 vulnerabilities found across 269 projects.
- **Muse Glimmer 30B (Meta, self-reported):** SWE-Bench Pro 51.2%, AIME 94.7, GPQA 83.5, MCP Atlas 75.5 (vs 54.2 Gemma4-31B / 62.5 Qwen3.6-27B); trails Qwen3.6-27B on Terminal-Bench 2.1 (51.7 vs 60.7) and OSWorld-Verified.
- **Mistral Medium 3.5 (Mistral, vendor-announced):** SWE-Bench Verified 77.6% (ahead of Qwen3.5 397B and Devstral 2), τ³-Telecom 91.4.

