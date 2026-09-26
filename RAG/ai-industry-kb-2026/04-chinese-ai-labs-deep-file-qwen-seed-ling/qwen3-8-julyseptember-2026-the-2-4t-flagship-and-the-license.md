---
id: ai-industry-kb-2026/04-chinese-ai-labs-deep-file-qwen-seed-ling/qwen3-8-julyseptember-2026-the-2-4t-flagship-and-the-license
title: "Qwen3.8 (July–September 2026) — the 2.4T flagship and the license question"
domain: chinese-ai-labs-deep-file-qwen-seed-ling
role: deep-dive
task: licenses
actors: ["Alibaba", "China", "Google", "Hugging Face", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-05", "2026-05-19", "2026-05-20", "2026-05-21", "2026-06-02", "2026-06-24", "2026-07-15", "2026-07-19", "2026-07-21", "2026-07-23", "2026-07-27", "2026-08-03", "2026-08-12", "2026-08-26", "2026-09", "2026-09-02", "2026-09-03", "2026-09-04"]
keywords: ["license", "qwen", "agent", "apache", "benchmark", "benchmarks", "blackwell", "cost", "datacenter", "fp8", "gguf", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1620, 1649]
section: "4. Chinese AI Labs — Deep File: Qwen, Seed, Ling"
sha256: 878b625f4f239dc3a11761d9b7e1184a53a50a0524fde46883641dfe34656c3f
---

# Qwen3.8 (July–September 2026) — the 2.4T flagship and the license question

Every Qwen3.7 release was proprietary or hosted: Qwen3.7-Max unveiled 2026-05-20 at the Alibaba Cloud Summit/Apsara Conference in Hangzhou (API public 2026-05-21/23; 1M context, 65,536 output tokens; AA Index 56.6); Qwen3.7-Plus officially released 2026-06-02 (preview 2026-05-19 on the Qianwen site) as a multimodal agent-oriented model unifying GUI and CLI operations in a single closed loop — visual perception + screen reading + terminal code execution — claiming top-5 globally and #1 in China on Vision Arena [VENDOR]; scale known only through a third-party benchmark table (17B active [COMMUNITY]); Qwen3.7 Flash went public 2026-07-27 per OpenRouter (the QwenCloud `qwen3.7-flash-2026-07-15` identifier is a snapshot, not the launch date — cite OpenRouter for availability, QwenCloud separately for technical limits).

### Qwen3.8 (July–September 2026) — the 2.4T flagship and the license question

Announced 2026-07-19 at the World AI Conference in Shanghai, two days after Moonshot's Kimi K3 (2.8T; Alibaba holds ~36% of Moonshot per MLQ), Qwen3.8 is a 2.4-trillion-parameter sparse MoE — the first Qwen multimodal model above 1T params. The per-token active count was NOT disclosed at announcement (the most-flagged gap in independent technical reviews). Qwen3.8-Max-Preview was available immediately through the Token Plan, Qoder, and QoderWork at 10% of standard pricing (text/image/video/document input). General availability landed 2026-08-03 (Bloomberg, Quartz, MarkTechPost, zeronoise.ai digest, bitrouter spec): 2.4T total / ~95B active [third-party-reported], 1M context (991K max input, 131K max output), 512 experts (10 routed + 1 shared per token), 92 layers, flat pricing $2/M input / $6/M output (implicit cache reads $0.25/M, explicit $0.17/M; bitrouter CNY 12/36 base at ~7.2 CNY/USD).

#### The 2026-08-12 weight drop, dissected

Repositories `Qwen/Qwen3.8-2.4T-A95B` and `Qwen/Qwen3.8-2.4T-A95B-FP8` (official fine-grained FP8 quant, block size 128, mixed BF16/F8_E4M3) published on Hugging Face and ModelScope on 2026-08-12; NVIDIA's August 12 deployment blog confirms "Alibaba released the open weights." Three dimensions must not be collapsed:

1. **Weights published**: real — BF16 + official FP8 + community Unsloth GGUF quants [COMMUNITY]. First 2T-class downloadable model.
2. **License permissiveness**: NOT Apache 2.0. The HF repo lists a custom "qwen3.8-max" license — past 100M MAU or $20M monthly revenue you must display the model name; MaaS/AI-assistant businesses past $50M trailing revenue need a separate commercial license; internal use exempt (aicodingdir.com, explainx.ai). Reporting says Alibaba plans revenue-sharing with large commercial users of the weights. Consolidation must classify this as "open-weight with custom restrictive license," alongside Llama Community and the GLM-5.3 License — never in an Apache-2.0/MIT index.
3. **Feature parity**: the open checkpoint is text-only and thinking-mode-only — it drops vision input and the native 1M context the hosted API offers (explainx.ai; HF discussion thread opened within hours calling it a "huge disappointment," comparing it to DLC paywalling). "Open weights" here means the capability ceiling is public and auditable; the full product experience is API-gated. Community commentary (thecherrycreeknews.com) frames it: an available model is a service rented on the vendor's terms; an open-weight model is an artifact you hold.

The developer-relevant open artifact of the drop is really Qwen3.8-27B, not the 2.4T Max: 27B dense native multimodal, Apache 2.0, 262K native context extensible to 1M, supported in Transformers/vLLM/SGLang/Docker Model Runner with an official FP8 quant — ~14–16 GB VRAM at 4-bit, RTX 4090/5090 class (aicodingdir.com). The Max remains a datacenter-only artifact: BF16 ~4.89 TB, FP8 roughly half, cheapest usable quant ~450 GB RAM+VRAM; NVIDIA reference deployment on a GB300 NVL72 rack (72 Blackwell Ultra GPUs), >4,000 tok/s per GPU in FP8.

### Qwen3.8-Flash (2026-08-26) and Qwen3.8-Max-0902 (2026-09-02)

Qwen3.8-Flash (2026-08-26): multimodal MoE, 6B active, open weights, 262K→1M via YaRN, $0.16/$0.47 per M — training cost ~1/9 of Qwen3.7-Plus [VENDOR]; completes the 3.8 family at the efficient end. Qwen3.8-Flash-Next (same day): hosted-only variant on QwenCloud/Model Studio, a Qwen4-generation 125B/6B-active architecture preview with Alibaba-reported SWE-bench Pro 62.5 / GPQA Diamond 91.7 [VENDOR]. Qwen3.8-Max-0902 (2026-09-02): newer Max checkpoint, 4th preliminary in the September 5 Code Arena WebDev table — canonical home wave 2.1/05; recap only. The Max lineage reads: Qwen3-Max (1T, Sept 2025, estimated) → Qwen3.6-Max-Preview (Apr 2026, closed) → Qwen3.7-Max (May 2026, closed) → Qwen3.8-Max-Preview (Jul 19, 2026) → Qwen3.8-Max GA (Aug 3, 2026) → Qwen3.8-Max-0902 (Sept 2, 2026).

### Qwen Image 3.0 (2026-07-21) — closed departure

`qwen-image-3.0` / `qwen-image-3.0-pro` shipped on Model Studio (also OpenRouter) 2026-07-21 with no open weights, no technical report, and no published benchmarks — a departure from Qwen-Image 1.0's Apache 2.0 launch; closed, API-only (tech-insider.org). Stated differentiation vs Google's Nano Banana 2 (Sept 2026) and OpenAI's GPT Image 2.5 (Sept 8, 2026) is dense text/structured layouts, not aesthetics: 4,500-token instruction input, legible text down to ~10 px, dense single-pass multi-panel layouts (infographics, newspapers, storyboards), 12 languages, LaTeX math notation, web-page/game-interface/livestream-overlay simulation; max 2048×2048; 1–3 reference images for editing. Framed as "industrial-grade visual foundation" — "not just pursuing good-looking, it is pursuing useful." A 302.AI head-to-head (Aug 11, 2026) compared it against GPT-Image-2 in infographic/layout scenarios [COMMUNITY]; no public eval suite exists [UNVERIFIED].

### Ling 3.0 Flash (July–September 2026) — the agent-economy cost floor

InclusionAI (Ant Group's AI lab) shipped Ling 3.0 Flash 2026-07-23: 124B / ~5.1B active MoE, 262K context, 32K max output, reasoning + tool calling, text in/out (no vision), distributed as a free tier on OpenRouter (`inclusionai/ling-3.0-flash:free`) — which explains its benchmark presence as a new entrant. Intelligence 22.6 / coding 50.6–55.6 on aggregators [COMMUNITY]: modest vs frontier, competitive on cost. Siblings: Ling 3.0 Flash VL (vision, September 2026, AI Index 41.2), Ling 3.0 Flash Sante (2026-09-04, free variant, secondary-sourced), and Ling 3.0 Flash Fin (2026-09-03): the first finance-specialized model in the Ant Ling family, built with financial institutions and domain experts, 124B/5.1B active, 256K context, 236K max output, function calling + reasoning, $0.06/M input / $0.18/M output — the cheapest finance-domain model of the window [SECONDARY vendor-adjacent]. The Ling line, paired with Qwen3.8-Flash ($0.16/$0.47), is the evidence base for 2026's bifurcation: datacenter flagships vs ultra-cheap agent workhorses.

### Seed 2.1 Turbo (2026-06-24) — dated with reserve

