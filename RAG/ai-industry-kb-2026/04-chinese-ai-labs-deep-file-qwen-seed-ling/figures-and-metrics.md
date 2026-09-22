---
id: ai-industry-kb-2026/04-chinese-ai-labs-deep-file-qwen-seed-ling/figures-and-metrics
title: "Figures and metrics"
domain: chinese-ai-labs-deep-file-qwen-seed-ling
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "ByteDance", "China", "Google", "Moonshot", "Nvidia", "OpenRouter", "SGLang", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-02-16", "2026-03-30", "2026-04-02", "2026-04-16", "2026-04-21", "2026-05-19", "2026-05-20", "2026-06-02", "2026-06-24", "2026-07-19", "2026-07-21", "2026-07-23", "2026-07-27", "2026-08", "2026-08-03", "2026-08-12", "2026-08-14", "2026-08-26", "2026-09", "2026-09-02", "2026-09-03", "2026-09-04", "2026-09-22"]
keywords: ["agent", "agentic", "agents", "apache", "attention", "benchmark", "benchmarks", "blackwell", "claude", "compute", "cost", "datacenter"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1435, 1566]
section: "4. Chinese AI Labs — Deep File: Qwen, Seed, Ling"
sha256: 30dee09db599a9a3840399b067a872b28ecb81c0aac58d273fb42e4b2ea34ae7
---

# Figures and metrics

## Figures and metrics

### Specification table (Qwen/Alibaba line, 2026)

| Model | Date | Total / active | Context | License | API pricing (in/out per 1M) |
|---|---|---|---|---|---|
| Qwen3.5-397B-A17B | 2026-02-16 | 397B / 17B MoE | 262K (YaRN → 1M) | Apache 2.0 | $0.40 in (Alibaba Cloud) |
| Qwen3.5-Plus | 2026-02-16 | not disclosed | 1M default | Proprietary | not disclosed |
| Qwen3.5-Omni (Plus/Flash/Light) | 2026-03-30/31 | not disclosed | 256K | Proprietary (API-only) | not disclosed |
| Qwen3.6-Plus | 2026-04-02 | not disclosed | 1M | Proprietary | ¥2/¥12 (Bailian) |
| Qwen3.6-35B-A3B | 2026-04-16 | 35B / 3B MoE | 262K class | Open weights | community-hosted |
| Qwen3.6-Max-Preview | 2026-04-21 | not disclosed | 1M | Proprietary | preview |
| Qwen3.7-Max | 2026-05-20 | not disclosed | 1M, 65K output | Proprietary | not disclosed |
| Qwen3.7-Plus | 2026-06-02 | ~17B active [COMMUNITY] | not disclosed | Proprietary (API-only) | not disclosed |
| Qwen3.7 Flash | 2026-07-27 | not disclosed | not disclosed | Hosted | not disclosed |
| Qwen3.8-Max-Preview | 2026-07-19 | 2.4T (active undisclosed at launch) | 1M | Proprietary preview | 10% of standard pricing |
| Qwen Image 3.0 / Pro | 2026-07-21 | not disclosed | n/a (image model) | Proprietary (API-only) | not disclosed |
| Qwen3.8-Max | 2026-08-03 | 2.4T / 95B active [third-party-reported] | 1M (991K in / 131K out) | Proprietary API; weights: custom qwen3.8-max | $2.00 / $6.00; cache reads $0.25/$0.17 |
| Qwen3.8-2.4T-A95B (+FP8) | 2026-08-12 | 2.4T / 95B active [third-party-reported] | 262K native (extensible to 1,010,000) | Custom qwen3.8-max | downloadable weights |
| Qwen3.8-27B | 2026-08-14 | 27B dense | 262K native (extensible 1M) | Apache 2.0 | ~14–16 GB VRAM at 4-bit |
| Qwen3.8-Flash | 2026-08-26 | 6B active | 262K native → 1M via YaRN | Open weights | $0.16 / $0.47 |
| Qwen3.8-Flash-Next | 2026-08-26 | 125B / 6B active | not disclosed | Hosted (Model Studio) | not disclosed |
| Qwen3.8-Max-0902 | 2026-09-02 | 2.4T / 95B | 1M | Custom qwen3.8-max | $2.00 / $6.00 |

### Model-by-model license map (the open-weight spectrum)

- Apache 2.0 (permissive, true open weights): Qwen3.5-397B-A17B (2026-02-16), Qwen3.6-35B-A3B (2026-04-16), Qwen3.8-27B (2026-08-14). Pattern: mid-tier models, ecosystem adoption plays.
- Custom restrictive (weights downloadable, commercial terms attached): Qwen3.8-Max / Qwen3.8-2.4T-A95B (2026-08-12) under the bespoke "qwen3.8-max" license — name display above 100M MAU or $20M monthly revenue; separate commercial license for MaaS/AI-assistant past $50M trailing revenue; internal use exempt; reported revenue-sharing plans for large commercial users. Peer class: Llama Community, GLM-5.3 License.
- Proprietary hosted API (no weights): Qwen3.5-Plus, Qwen3.5-Omni, Qwen3.6-Plus, Qwen3.7-Max, Qwen3.7-Plus, Qwen3.8-Max-Preview, Qwen3.8-Flash-Next, Qwen Image 3.0/Pro.
- Closed/proprietary, terms undisclosed: Seed 2.1 Turbo (no license published anywhere reviewed).
- Undisclosed pricing, hosted: Qwen3.5-Plus, Qwen3.5-Omni tiers, Qwen3.7-Max, Qwen3.7-Plus, Qwen3.7 Flash, Qwen3.8-Flash-Next, Qwen Image 3.0/Pro.

### Seed 2.1 Turbo and Ling 3.0 Flash figures

- Seed 2.1 Turbo (2026-06-24): 262K context (262,144); text/image/video in → text out; reasoning (full effort ladder, `none` disables thinking); tool calling and structured output; streaming; closed/proprietary — parameter count, training cutoff, and license all undisclosed [UNVERIFIED]. Pricing $0.50/$2.50 per M tokens (OpenRouter/NanoGPT, verified current as of 2026-09-22). Positioned as the faster, lower-cost variant for high-throughput/latency-sensitive agent workloads, not a flagship reasoning tier.
- Ling 3.0 Flash (2026-07-23): 124B total / ~5.1B active MoE; 262,144-token context; 32K max output; reasoning + tool calling; text in/out (no vision). Free tier on OpenRouter (`inclusionai/ling-3.0-flash:free`).
- Ling 3.0 Flash Fin (2026-09-03): 124B/5.1B active; 256K context; 236K max output; finance-specialized (Ant Group + financial institutions + domain experts); $0.06/M input, $0.18/M output [SECONDARY vendor-adjacent]; first finance-enhanced model in the Ant Ling family.
- Ling 3.0 Flash VL (September 2026): vision variant, text/image/video input; AI Index 41.2 [COMMUNITY]. Ling 3.0 Flash Sante (2026-09-04): free variant [UNVERIFIED provenance detail — secondary only].

### Architecture figures

- Qwen3.5-397B-A17B: hybrid Gated Delta Networks + sparse MoE; 10 routed experts + 1 shared expert per token (~4.3% activation); 512 experts.
- Qwen3.8-Max: 2.4T total / ~95B active per token (~4% activation, third-party-reported, not officially disclosed); 92 layers; hidden dim 8192; padded token embedding 248,320; 512 experts with 10 routed + 1 shared per token; 262,144 native context, extensible to 1,010,000; text/image/video input → text output. Active count at announcement was NOT disclosed — flagged by multiple independent technical reviews; 2.4T describes disk size, not per-query compute cost.
- Qwen3.8-Flash: multimodal MoE, 6B active; Gated DeltaNet + Qwen Sparse Attention, Gated Residual (4 branches), n-gram embeddings, refined Muon optimizer; training cost ~1/9 of Qwen3.7-Plus [VENDOR]; 262K native → 1M via YaRN; 131K max output; open weights; API $0.16/$0.47 per M.
- Qwen3.8-Flash-Next: 125B / 6B active; Qwen4-generation architecture preview; hosted-only (QwenCloud/Model Studio).
- Qwen3.5-Omni: Thinker–Talker architecture with Hybrid-Attention MoE, native end-to-end omnimodal (not cascaded); 256K context; >10 hours audio per request; 100M+ hours audiovisual pretraining; 113 languages for speech recognition, 36 for speech generation; voice cloning via API; tiers Plus / Flash / Light.

### Benchmark reality checks (kept in this section per dedup rule)

#### Qwen3.8-Max — vendor-reported vs independent

- Alibaba self-reported [VENDOR]: Terminal-Bench 2.1 **86.6**, PaperBench **93.0**, OSWorld-Verified **86.1**, Agents' Last Exam **52.4**, SWE-bench Pro **67.7**. Briefly topped the AA Agentic Index before a methodology update.
- Independent composite: Artificial Analysis Intelligence Index v4.3 (Sept 4, 2026): Qwen3.8-Max ≈ **40** (4th open) vs GLM-5.3 ≈ 45, Kimi K3 ≈ 44, GLM-5.3-Flash 42 [DIRECTIONAL].
- "Second only to Fable 5" at the 2026-07-19 announcement: Alibaba's own claim [VENDOR]; no benchmark table, model card, or license published at announcement; independent test results did not exist.
- Exact 95B active figure: widely reported but never officially disclosed by Alibaba — classify as third-party-reported [COMMUNITY], not primary. AA Index lists "Qwen3.8 (2.4T A95B)."

#### Qwen3.8-Flash-Next — vendor-reported vs community

- Alibaba-reported [VENDOR]: SWE-bench Pro **62.5**, SWE-bench Multilingual **81.0**, CoWorkBench **73.9**, Toolathlon-Verified **73.5**, GPQA Diamond **91.7**, HLE **35.9**.
- Independently compared by Zenn: DeepSWE 1.1 **58.7** [COMMUNITY]. Medium analysis (James Anderson) explores how a 6B-active model compares against 17B-active competitors [COMMUNITY].

#### Qwen3.8-27B

- Terminal-Bench 2.1 **73.0** [VENDOR]; AA Index 52 (size-class leader) [DIRECTIONAL].
- Vendor claim "outperforms Qwen3.7-Plus overall" and beats Claude Opus 4.6 Max in some benchmarks [VENDOR] (via gigazine.net secondary).
- The developer-relevant open artifact: 4-bit at ~14–16 GB VRAM (RTX 4090/5090 class), per aicodingdir.com framing — the 2.4T Max is datacenter-only.

#### Qwen3.7-Plus, Qwen3.6-27B, Qwen3.7-Max

- Qwen3.7-Plus: SWE-bench Pro **55.8** as reported in a third-party benchmark table [COMMUNITY]; vendor claim "top 5 globally, #1 in China on Vision Arena" [VENDOR, secondary].
- Qwen3.6-27B: Terminal-Bench 2.1 **60.7**, AA Index 38 [COMMUNITY]; release date unpinned [DATE UNVERIFIED].
- Qwen3.7-Max: Artificial Analysis Index 56.6 [DIRECTIONAL, secondary].

#### Qwen3.5-Omni vs Gemini 3.1 Pro

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

