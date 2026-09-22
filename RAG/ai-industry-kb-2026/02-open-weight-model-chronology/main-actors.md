---
id: ai-industry-kb-2026/02-open-weight-model-chronology/main-actors
title: "Main actors"
domain: open-weight-model-chronology
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "EU", "Google", "Huawei", "Hugging Face", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "United States", "Unsloth", "Xiaomi", "Z.ai", "vLLM", "xAI"]
dates: ["2025-04-05", "2025-04-29", "2025-08-05", "2026-02-10", "2026-02-16", "2026-03-16", "2026-04", "2026-04-02", "2026-04-11", "2026-04-13", "2026-04-16", "2026-04-20", "2026-04-21", "2026-04-24", "2026-04-29", "2026-05-20", "2026-06-03", "2026-06-16", "2026-06-30", "2026-07-16", "2026-07-27", "2026-07-31", "2026-08-10", "2026-08-12", "2026-08-13", "2026-08-14", "2026-08-18", "2026-08-19", "2026-08-20", "2026-08-26", "2026-08-28", "2026-08-31", "2026-09", "2026-09-02", "2026-09-10", "2026-09-14", "2026-09-18", "2026-09-21"]
keywords: ["agent", "agentic", "apache", "ascend", "attention", "benchmark", "claude", "copilot", "cost", "cyber", "deepseek", "distribution"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [820, 900]
section: "2. Open-Weight Model Chronology"
sha256: 4aaffa997faf76c8bf0599726a1874eee6c869f2a8bf2c99a327e7fc693b211d
---

# Main actors

## Main actors

| Lab | Country | 2026 open-weight releases (this chronology) | License posture |
|---|---|---|---|
| Alibaba (Qwen team) | China | Qwen3.5-397B-A17B (Feb 16/17), Qwen3.6-35B-A3B (Apr 16), Qwen3.8-Max weights (Aug 12), Qwen3.8-27B (Aug 14), Qwen3.8-Max-0902 (Sep 2) | Apache 2.0 on open tiers; proprietary Plus/Max API tiers; custom restrictive license for Qwen3.8-Max |
| DeepSeek | China | V4-Pro / V4-Flash (Apr 24 preview), V4-Flash-0731, V4-Pro 0813, V4-Flash-Vision-Exp (Aug 31), V4.1-Flash (Sep 10); V4-Pro retired from API Sep 14 | MIT across the line (0813 weight status ambiguous) |
| Moonshot AI | China | Kimi K2.6 (Apr 13 preview → Apr 20/21 GA), Kimi K3 (Jul 16 launch → Jul 27 weights) | Modified MIT (display threshold above 100M MAU / $20M revenue) |
| Z.ai | China | GLM-5.2 (Jun 16), GLM-5.3 (announced Aug 14, weights Aug 28/29), GLM-5.3-Flash (Aug 26) | MIT for 5.2 and 5.3-Flash; bespoke GLM-5.3 License for 5.3 |
| Meta (Superintelligence Labs) | US | Muse Glimmer 30B (Aug 10); Llama 4 Scout/Maverick date-corrected to Apr 5, 2025 | Glimmer: Apache 2.0 (first for Meta); Llama 4: Community License |
| Mistral AI | France/EU | Small 4 (Mar 16), Medium 3.5 (Apr 29/30); Large 3 (Dec 2025) | Apache 2.0 (Small 4, Large 3); modified MIT (Medium 3.5) |
| Google | US | Gemma 4 family (Apr 2), Gemma 4 12B (Jun 3) | Apache 2.0 (first straight Apache Gemma) |
| OpenAI | US | gpt-oss-120b/20b (Aug 5, 2025 — one-year baseline) | Apache 2.0 |
| Xiaomi (MiMo) | China | MiMo-V2.6 Pro/Flash (Sep 21) | MIT |
| Unsloth (Daniel & Michael Han) | Australia/US | Tooling only: Dynamic Quantization 2.0 (current as of Jul 2026), v3.0 (Aug 19); GGUF/NVFP4 distribution via HF | Tooling vendor, not a model lab |

Notable absences (no new open-weight releases found for these labs in the August 22 – September 22 window): Moonshot beyond Kimi K3 (July 16); MiniMax beyond M3 (June 1); Mistral beyond Large 3; no GLM-5.4; no Kimi K4; no Llama 5.

**Anthropic and frontier posture (context, not open-weight releases):** Anthropic keeps flagships closed (Opus 4.6 Feb 5, Sonnet 4.6 Feb 17, Opus 4.8, Fable 5/5.1) and monetizes the workflow layer (Claude Code: agent teams, context compaction, adaptive thinking); xAI ships Grok closed (4.5 Jul 8 → 4.6 Aug 12 → 4.7 Sep 21) with distribution through GitHub Copilot plans on launch day; the industry pattern is "closed flagship + open flank" (OpenAI gpt-oss, Google Gemma 4) rather than opening flagships.

**China's open-weight strategy as structural posture:** MIT/Apache-licensed releases trained without Nvidia hardware (DeepSeek V4 family trained on Huawei Ascend silicon), priced at 1/20th–1/89th of frontier closed models on rate cards, distributed free on Hugging Face — export controls on chips are being routed around by open weights themselves. Mozilla's September 2026 read: Chinese open-weight models are 4.4 months behind the US frontier at ~30% of the cost.

## Timeline and context

### Month-by-month open-weight chronology (February → September 2026)

| Date | Event |
|---|---|
| 2026-02-16/17 | Qwen3.5-397B-A17B (Apache 2.0, 397B/17B MoE) — Alibaba opens the 2026 open-weight year; Chinese New Year Eve |
| 2026-03-16 | Mistral Small 4 (Apache 2.0, 119B/6B MoE); Mistral's six-products-in-15-days run |
| 2026-04-02 | Qwen3.6-Plus (proprietary API, agentic coding, 1M ctx) |
| 2026-04-02 | Gemma 4 family (Apache 2.0) |
| 2026-04-13 | Kimi K2.6 preview opens (eight-day preview) |
| 2026-04-16 | Qwen3.6-35B-A3B (open weights, 35B/3B MoE) |
| 2026-04-20/21 | Kimi K2.6 GA (1T/32B, Modified MIT) |
| 2026-04-21 | Qwen3.6-Max-Preview |
| 2026-04-24 | DeepSeek V4 Preview: V4-Pro (1.6T/49B) + V4-Flash (284B/13B), MIT, 1M ctx |
| 2026-04-29/30 | Mistral Medium 3.5 (dense 128B, modified MIT) — the real April 2026 Mistral flagship, not "Medium 3.1" |
| 2026-05-20 | Qwen3.7-Max unveiled, Hangzhou (proprietary); API public May 21–23 |
| 2026-06-03 | Gemma 4 12B (Apache 2.0, encoder-free multimodal) — not at Google I/O (corrects claim 7) |
| 2026-06-16 | GLM-5.2 weights (753B/40B — not 744B; MIT; IndexShare; AA v4.1 51 at release) |
| 2026-06-30–07-02 | AI Engineer World's Fair 2026, Moscone West, SF (event verified; Unsloth presence unverified) |
| 2026-07-16 | Kimi K3 launched (2.8T, first 3T-class open model) |
| 2026-07-27 | Kimi K3 open weights (Modified MIT) |
| 2026-07-31 | DeepSeek V4-Flash-0731 checkpoint (Terminal-Bench 2.1 82.7, AA Index 50) |
| 2026-08-10 | Meta Muse Glimmer 30B (Apache 2.0) — Meta's actual open release of the window |
| 2026-08-12 | Qwen3.8-Max weights (2.4T/95B, custom restrictive license, text-only) — first downloadable Max-tier Qwen |
| 2026-08-13 | DeepSeek V4-Pro 0813 GA revision (weight status ambiguous) |
| 2026-08-14 | Z.ai announces GLM-5.3 (743B); weights held ~2 weeks for safety review |
| 2026-08-14 | Qwen3.8-27B weights (Apache 2.0, 27.8B dense, AA Index 52) |
| 2026-08-18 | GLM-5.3 API opens ($1.40/$4.40 per M); GLM Coding Plan from $18/mo |
| 2026-08-19 | Unsloth Dynamic Quantization v3.0 (supersedes 2.0) |
| 2026-08-20 | "Ox Alpha" (later revealed as GLM-5.3-Flash) tops OpenRouter usage charts anonymously |
| 2026-08-26 | GLM-5.3-Flash weights (320B MoE, MIT) |
| 2026-08-28/29 | GLM-5.3 weights shipped (`zai-org/GLM-5.3`) — safety hold resolved as delay, not reversal |
| 2026-08-31 | DeepSeek V4-Flash-Vision-Exp (305B, MIT) — first V4 vision weights |
| 2026-09-02 | Qwen3.8-Max-0902 (newer checkpoint, 1M ctx); Muse Spark 1.3 shipped (closed, weights pending) |
| 2026-09-10 | DeepSeek V4.1-Flash GA (552B CED/CSA2/Engram, MIT) |
| 2026-09-14 | V4-Pro retired from the DeepSeek API (all traffic routed to V4.1-Flash) |
| 2026-09-18 | Qwen3.8-Omni-Flash (API-only) |
| 2026-09-21 | MiMo-V2.6 Pro/Flash (MIT) |
| 2026-09-21 | xAI launches Grok 4.7 ($2/$6 per M) — closed, context only |

### 2025 context anchors (kept minimal — the corrections that matter)

- 2025-04-05 — Llama 4 Scout/Maverick (2026 dating corrected).
- 2025-04-29 — Qwen3 family; Qwen3-30B-A3B SKU (2026-02-10 dating contradicted).
- 2025-08-05 — OpenAI gpt-oss (one-year anniversary baseline, Aug 5, 2026: 4.3M+ HF downloads).
- 2025-12 — Mistral Large 3 (April 2026 dating corrected); DeepSeek V3.2 (Dec 1, 2025 — the "April 11, 2026 DeepSeek V3.2" claim contradicted; V3.2-Exp Sep 29, 2025 introduced DeepSeek Sparse Attention).

### The mid-February DeepSeek V4 window (context, not a release)

DeepSeek V4's planned mid-February launch (reported by The Information) failed on the Spring Festival window and slipped to the April 24 preview. The surrounding ecosystem noise — fabricated benchmark screenshots (Epoch AI: FrontierMath scores fabricated), the "MODEL1" GitHub leak parsed as Engram architecture, and the "Hunter Alpha" mystery model (revealed as Xiaomi's) — is documented so that the RAG's retrieval on "DeepSeek V4 February" returns the confusion narrative rather than a phantom release.

### The GLM-5.3 safety hold as a governance datapoint

Z.ai's ~2-week self-imposed hold (announced August 14, 2026) on GLM-5.3 weights because cyber capabilities grew faster than expected is the first open-release delay of its kind in the chronology. It resolved as a delay, not a policy reversal — "the first crack in release-openly-immediately did not escalate." It belongs in any discussion of open-weight governance alongside the AISI finding that distributed weights cannot be recalled and that refusals are retry-strippable.

### Cross-wave pointers (dedup)

- DeepSeek architecture internals (DSA, V3.2 lineage, V4 sparse stack) → §5. Qwen Gated-DeltaNet hybrid mapping and Kimi-Linear KDA detail → §4 (Qwen) and §5 (Kimi). GLM-5.2/5.3 internals → §3. Inference engines (vLLM/SGLang/TGI/llama.cpp serving these models) → §6. Quantization formats (GGUF/NVFP4/MXFP4/FP8, Unsloth v3.0, K-Quant builds) → §8. Full pricing tables → §14. This section carries only the dated release facts, spec figures, and licensing terms needed for the timeline.

