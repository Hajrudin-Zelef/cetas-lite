---
id: ai-industry-kb-2026/04-chinese-ai-labs-deep-file-qwen-seed-ling/what-the-missing-numbers-tell-us-disclosure-gaps-as-data
title: "What the missing numbers tell us (disclosure gaps as data)"
domain: chinese-ai-labs-deep-file-qwen-seed-ling
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "Apple", "ByteDance", "China", "Google", "Hugging Face", "Moonshot", "Nvidia", "OpenRouter", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-02-16", "2026-04-02", "2026-05-19", "2026-06-02", "2026-06-24", "2026-07-15", "2026-07-17", "2026-07-19", "2026-07-23", "2026-08", "2026-08-03", "2026-08-12", "2026-08-26", "2026-09", "2026-09-03", "2026-09-22"]
keywords: ["disclosure", "agent", "agents", "benchmark", "benchmarks", "blackwell", "claude", "compute", "cost", "datacenter", "distribution", "fp8"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1510, 1574]
section: "4. Chinese AI Labs — Deep File: Qwen, Seed, Ling"
sha256: 41d1fb978b7c5bb327267622e359586eeb26f96dd8b6562e37bd724f59e1bb86
---

# What the missing numbers tell us (disclosure gaps as data)

- Beats Gemini 3.1 Pro on general audio understanding/reasoning/translation; matches on audio-visual comprehension [VENDOR].

#### Qwen Image 3.0

- No published benchmarks at launch [UNVERIFIED] — performance claims rest on feature descriptions (4,500-token instruction input; text legible down to ~10 px; dense single-pass multi-panel layouts; 12 languages; LaTeX math) and a single independent review (302.AI head-to-head vs GPT-Image-2 in infographic/layout scenarios, Aug 11, 2026) [COMMUNITY]. Treat ranking claims with caution.

#### Ling 3.0 Flash

- Intelligence 22.6 (outperforms ~74% of tracked models on one aggregator); coding 50.6–55.6 depending on aggregator [COMMUNITY] — modest vs frontier, competitive on cost. Ling 3.0 Flash Fin evaluated on FinFIRST, FinSearchComp Verified, FinCRAFT, Finance Agent, APEX-Agents, SpreadsheetBench, τ³-Banking [VENDOR — vendor-reported suite].

#### Seed 2.1 Turbo

- No independent benchmark scores published [UNVERIFIED]. ByteDance's launch materials used chart-only superlatives with no absolute numbers in prose. Seed 2.1 Pro (sibling): preview ranked 8th on Code Arena at 1539, framed by ByteDance as "level with Claude Opus 4.6" — secondary, chart-sourced, treat with caution.

### What the missing numbers tell us (disclosure gaps as data)

- Qwen3.8-Max active-per-token count: NOT disclosed at the 2026-07-19 announcement (flagged by multiple independent technical reviews); "95B" is third-party-reported post-launch, never Alibaba-confirmed — 2.4T describes disk size, not per-query compute cost.
- Qwen3.5-Plus and Qwen3.6-Plus total/active splits: never separately disclosed; Qwen3.7-Max and Qwen3.7-Plus totals: undisclosed; Qwen3.7-Plus scale known only via a third-party benchmark table (17B active [COMMUNITY]).
- Qwen3.7 Flash: no public specs at all (date from OpenRouter, limits from the QwenCloud snapshot ID) — a hosted efficiency tier with no datasheet.
- Seed 2.1 Turbo: parameter count, training cutoff, and license terms all undisclosed [UNVERIFIED] — the least-documented flagship-adjacent model in this part.
- Qwen3.6-27B: exists only in third-party benchmark tables (Terminal-Bench 2.1 60.7, AA Index 38); no release date, no model card, no vendor acknowledgment pinned [DATE UNVERIFIED].
- Qwen Image 3.0: no weights, no technical report, no benchmarks at launch [UNVERIFIED] — the closed-API image flagship is the least transparent release in the Qwen 2026 line.
- Pattern: disclosure depth is inversely correlated with product maturity — mid-tier open weights come with configs and community replication; flagship hosted/API products come with marketing claims and chart images. The RAG should read missing fields as a strategic signal, not a research failure.

### Training and serving economics

- Qwen3.8-Flash training cost ≈ 1/9 of Qwen3.7-Plus [VENDOR] — the efficiency-tier pitch is cost-of-training, not just cost-of-inference.
- MoE activation rates: Qwen3.5-397B-A17B ~4.3% (17B of 397B, 10 routed + 1 shared); Qwen3.8-Max ~4.0% (95B of 2.4T, third-party-reported); Ling 3.0 Flash ~4.1% (5.1B of 124B) — the Chinese 2026 MoE convention clusters at ~4% activation.
- Official FP8 quant of Qwen3.8-Max: fine-grained, block size 128, mixed BF16/F8_E4M3, claimed "nearly identical" to BF16 by Alibaba; halves the ~4.89 TB BF16 footprint. Community path: Unsloth GGUF quants 1.31 TB (4-bit UD-IQ4_XS) down to ~397–564 GB (1-bit) [COMMUNITY].
- The real deployability ladder, August 2026: Qwen3.8-27B (~14–16 GB VRAM at 4-bit, single RTX 4090/5090) ≫ Qwen3.8-Flash (6B active, open weights, API $0.16/$0.47) ≫ Ling 3.0 Flash (free tier) ≫ Qwen3.8-Max (datacenter-only, GB300 NVL72 rack, >4K tok/s per GPU FP8). Consolidation should present the ladder, not the flagship alone.

### Qwen3.7-Plus and Qwen3.6-27B — the thin files

Two entries in this part are documented from fragments, and the RAG should carry them thin rather than padded:

- Qwen3.7-Plus (official 2026-06-02, preview 2026-05-19): proprietary, not open-source; multimodal agent-oriented (GUI + CLI in one closed loop; visual perception, screen reading, terminal code execution); vendor claim "top 5 globally, #1 in China on Vision Arena" [VENDOR]; scale known only via a third-party benchmark table (17B active [COMMUNITY], SWE-bench Pro 55.8 [COMMUNITY]). No vendor spec sheet was found in research — the model is real but thinly documented.
- Qwen3.6-27B: dense variant; exists only in third-party benchmark tables (Terminal-Bench 2.1 60.7, AA Index 38 [COMMUNITY]); no release date pinned, no vendor acknowledgment found [DATE UNVERIFIED]. Do not invent a date; carry the gap visibly.

### Pricing chronology (February → September 2026)

- 2026-02-16: Qwen3.5-397B-A17B $0.40/M input on Alibaba Cloud — the open-tier hosted price anchor for the year.
- 2026-04-02: Qwen3.6-Plus ¥2/¥12 (Bailian) — domestic pricing, ~$0.28/$1.67 equivalent.
- 2026-07-19 → 2026-08-03: Qwen3.8-Max-Preview at 10% of standard pricing via Token Plan, Qoder, QoderWork — the preview discount window.
- 2026-08-03: Qwen3.8-Max GA at flat $2.00/$6.00 (no context-length tiers); implicit cache reads $0.25/M, explicit cache reads $0.17/M; bitrouter spec CNY 12/36 base at ~7.2 CNY/USD.
- 2026-08-26: Qwen3.8-Flash $0.16/$0.47 — the order-of-magnitude efficiency step inside the same generation.
- 2026-06-24 → 2026-09-22: Seed 2.1 Turbo $0.50/$2.50 (OpenRouter/NanoGPT) — ByteDance's mid-tier agent pricing, stable across the window.
- 2026-07-23 → 2026-09-22: Ling 3.0 Flash free tier on OpenRouter — the zero-marginal-cost agent model.
- 2026-09-03: Ling 3.0 Flash Fin $0.06/$0.18 — cheapest domain-specialized model of the window; ~33× cheaper input than Qwen3.8-Max, ~33× cheaper output.
- Read together: the 2026 Chinese pricing ladder spans two orders of magnitude ($0.06 → $6.00 per M output), and the cheapest tier is a 124B-class MoE, not a small model — scale and price decoupled.

### Deployment economics

- Qwen3.8-Max weights: BF16 ~4.89 TB; FP8 roughly half; cheapest usable quant ~450 GB combined RAM+VRAM (Unsloth community GGUF from 1.31 TB at 4-bit UD-IQ4_XS down to ~397–564 GB at 1-bit) [COMMUNITY]; lossless BF16 ~4.9 TB.
- NVIDIA reference deployment (August 12, 2026 blog): GB300 NVL72 rack (72 Blackwell Ultra GPUs), >4,000 tok/s per GPU in FP8.
- vLLM/SGLang one-line serving documented (`vllm serve "Qwen/Qwen3.8-2.4T-A95B-FP8"`); TokenSpeed also recommended. Hardware is the constraint, not software.
- Pricing stack, August–September 2026 [SECONDARY/vendor-adjacent]: Qwen3.8-Max $2.00/$6.00 (implicit cache reads $0.25/M, explicit cache reads $0.17/M; bitrouter spec: CNY 12/36 base at ~7.2 CNY/USD); Qwen3.8-Flash $0.16/$0.47; Seed 2.1 Turbo $0.50/$2.50; Ling 3.0 Flash Fin $0.06/$0.18; Ling 3.0 Flash free tier; Qwen3.8-Max-Preview at 10% of standard pricing (July 19–Aug 3 window); Qwen3.6-Plus ¥2/¥12 on Bailian; Qwen3.5-397B-A17B $0.40/M input on Alibaba Cloud.

## Main actors

### Alibaba Group / Alibaba Cloud / Qwen team

- Developer Shuai Bai named in connection with the Qwen3.8 2.4T announcement (first Qwen multimodal model above 1T params). Model Studio (QwenCloud) is the hosted-API surface; ModelScope is the domestic model hub mirroring Hugging Face publications. Token Plan, Qoder, and QoderWork carried the Qwen3.8-Max-Preview at 10% pricing. Beijing approved Apple Intelligence powered by Qwen in China on 2026-07-15 (memeburn.com) — distribution context. Alibaba holds ~36% of Moonshot (MLQ) — the ownership link to the Kimi line, relevant to the WAIC-week sequencing (Kimi K3 2026-07-17 → Qwen3.8 2026-07-19).

### Ant Group / InclusionAI

