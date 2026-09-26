---
id: ai-industry-kb-2026/02-open-weight-model-chronology/mistral-three-releases-three-dates-april-11-grouping-correct
title: "Mistral — three releases, three dates (April-11-grouping correction)"
domain: open-weight-model-chronology
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Unsloth", "Xiaomi", "Z.ai", "xAI"]
dates: ["2025-05", "2025-08", "2025-08-05", "2025-10", "2025-12", "2026-02", "2026-03", "2026-03-16", "2026-04", "2026-04-02", "2026-04-11", "2026-04-13", "2026-04-20", "2026-04-29", "2026-06-03", "2026-07-16", "2026-07-27", "2026-07-30", "2026-08-31", "2026-09-21"]
keywords: ["mistral", "agent", "agentic", "apache", "attention", "benchmark", "benchmarks", "claude", "cost", "datacenter", "deepseek", "fable 5"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [708, 731]
section: "2. Open-Weight Model Chronology"
sha256: f3aea69a0ab18125744c067744f29af2a1578f4dcafa5ec6b5a2786f18b706d3
---

# Mistral — three releases, three dates (April-11-grouping correction)

- **2026-04-13 — Kimi K2.6 preview opens (eight-day preview); GA 2026-04-20/21.** 1T total / 32B active MoE, 384 experts, Modified MIT (commercial use/fine-tuning/resale with no royalties; operators above 100M MAU or $20M monthly revenue must display "Kimi K2.6" in the product UI), 256K/262K context (256K shorthand vs exact 262,144), native text/image/video input. [SECONDARY] Launch-coverage benchmarks: SWE-Bench Pro 58.6% vs GPT-5.4 57.7% vs Claude Opus 4.6 53.4% — ahead of Opus 4.6, not merely "close"; SWE-Bench Verified 80.2% (10-run avg); Terminal-Bench 2.0 66.7%; HLE-Full with tools 54.0% vs GPT-5.4 52.1% vs Opus 4.6 53.0%; AIME 2026 96.4%; OSWorld-Verified 73.1%.
- **2026-07-16 — Kimi K3 launched; open weights ~2026-07-27.** 2.8T total — the first "3T-class" open model, an order-of-magnitude scale jump over the spring 2026 400B–750B class (Maverick 400B, GLM-5.2 753B, Qwen3.5 397B). Sparse MoE: 896 routed experts with 16 active per token; Kimi Delta Attention (KDA, hybrid linear attention — up to 6.3× faster decoding at million-token contexts [VENDOR]); Attention Residuals (~25% higher training efficiency at <2% cost [VENDOR]); 1M native context; native visual understanding (text+image); always-on thinking mode default. Modified MIT license (one curated list flags a 100M MAU threshold on the Kimi line — verify the license file before commercial deployment). [VENDOR] Launch figures: Frontend Code Arena 1,679 Elo — #1 (ahead of Claude Fable 5 at 1,631 and GPT-5.6 Sol at 1,618; #1 in 6 of 7 front-end domains); Terminal-Bench 2.1 88.3 vs Opus 4.8 84.6; FrontierSWE 81.2 vs Opus 4.8 66.7; SWE-Bench Verified 76.8%; GPQA-Diamond 93.5; AIME 2025 96.1%; BrowseComp 91.2; HLE-Full 43.5 (behind frontier); Artificial Analysis Intelligence Index 57 (July 30, 2026) — #1 among open-weight models; independent Groundtruth geological-reasoning benchmark (Aug 28, 2026): highest aggregate score 91% across OpenAI/Anthropic/DeepSeek/Google/xAI/Moonshot models (confidence intervals overlapped). API ~$3/M input / $15/M output. Moonshot's own honest caveat: K3 "still trails the most powerful proprietary models" overall. Inference cost: 1,390 GB for weights alone at INT4 — datacenter-class serving, not a laptop model.
- **"Kimi K3 beats Opus 4.8" — benchmark-dependent [VENDOR].** What the evidence supports: GDPval-AA v2 1687 vs 1600 (Kimi K3 vs Claude Opus 4.8); AA Intelligence ~57, 4th overall, ahead of Opus 4.8 in some AA snapshots; Frontend Code Arena leadership and several agentic-coding wins. But Kimi K3 trailed the strongest closed models on the hardest reasoning evaluations. Artificial Analysis aggregates multiple suites into its Intelligence Index and individual suite scores can diverge from the headline index — always name the suite. A blanket "Kimi K3 beats Opus 4.8" is misleading.
- Kimi-Linear side note (architecture deep dive → §5): `moonshotai/Kimi-Linear-48B-A3B-Instruct` refines the hybrid recipe with Kimi Delta Attention (channel-wise gating per feature dimension, replacing Qwen3-Next's scalar gate) and swaps gated full-attention layers for MLA — a GDN+MLA hybrid distinct from pure DeltaNet.

### Mistral — three releases, three dates (April-11-grouping correction)

- **December 2025 — Mistral Large 3 (CORRECTION: not April 2026).** HF repo `mistralai/Mistral-Large-3-675B-Instruct-2512` (the "2512" suffix = December 2025): 675B total / 41B active MoE flagship, 256K context, text+image input, 80+ languages, Apache 2.0. A teased reasoning variant of Large 3 had not shipped as of April 2026.
- **2026-03-16 — Mistral Small 4 (CORRECTION: not April 11).** 119B total / 6B active (128 experts, 4 active per token), Apache 2.0, 256K context, multimodal (text+images). First Mistral model to unify Magistral (reasoning), Pixtral (multimodal) and Devstral (agentic coding) capabilities; configurable reasoning effort. Shipped amid a six-products-in-15-days run (Leanstral code agent Apache 2.0, NVIDIA Nemotron Coalition partnership, Mistral Forge Mar 17, Voxtral TTS Mar 23, Spaces CLI Mar 31).
- **2026-04-29/30 — Mistral Medium 3.5 (CORRECTION: "Mistral Medium 3.1" is the wrong name).** Sources split 29 vs 30; Mistral docs card `mistral-medium-3-5-26-04`. "Mistral Medium 3.1" exists but is an **August 2025** refresh of Medium 3 ("Medium 3.1 25.08") — not an April 2026 release. Dense 128B (all active), 256K context, vision encoder trained from scratch, configurable per-request reasoning effort, **modified MIT license** (open weights on HF, free below $20M/month revenue), $1.50/$7.50 per M tokens, self-hosted on ~4 GPUs. "First flagship merged model" — collapses three lines into one: Mistral Medium 3.1 (instruction-following), Magistral (reasoning), Devstral 2 (agentic coding); retired Magistral from Le Chat and Devstral 2 from Vibe CLI. [SECONDARY] Vendor-announced: SWE-Bench Verified 77.6% (ahead of Qwen3.5 397B and Devstral 2), τ³-Telecom 91.4. Became the new default in Mistral Vibe and Le Chat. (The original closed Mistral Medium 3, May 2025, is deprecated with retirement scheduled August 31, 2026.)
- The "April 11, 2026 grouped release wave" (Mistral Small 4 + Medium 3.1 + Large 3; Grok 4.1/4.2; DeepSeek V3.2) never happened as stated — its components belong to December 2025, February 2026, March 2026, April 29–30, July 23, and August 7–12, 2026. Real April 2026 events: xAI STT/TTS APIs (Apr 17–19), Mistral Medium 3.5 (Apr 29–30), Kimi K2.6 preview/GA (Apr 13 → 20–21).

### Google — Gemma 4 (date correction) and the gpt-oss baseline

- **2026-04-02 — Gemma 4 open-model family launched.** Gemma 4 moved from custom Gemma Terms (Gemma 1–3) to Apache 2.0, removing Google's own licensing friction.
- **2026-06-03 — Gemma 4 12B announced (CORRECTION: not at Google I/O, May 19).** Dense 12B unified encoder-free multimodal model — text/image/**native audio input** (vision/audio flow directly into the LLM backbone), 256K context, runs on 16GB VRAM/unified memory, Apache 2.0 (Google's first straight Apache-2.0 Gemma), performance near the Gemma 4 26B MoE, with MTP drafters.
- **2025-08-05 — gpt-oss (OpenAI), one year on.** OpenAI's first open-weight LLMs since GPT-2 (2019), Apache 2.0: gpt-oss-120b (117B total / 5.1B active MoE, native MXFP4 ~65GB, single 80GB GPU) and gpt-oss-20b (20.9B / 3.6B active, ~13GB, 35–42 tok/s on a 16GB GPU, 128K context, 60.7% SWE-Bench Verified). 4.3M+ HF downloads — among the most-downloaded open-weight models ever. GPT-OSS-Safeguard followed October 2025. (2025 context anchor; chronology baseline.)

### MiMo (Xiaomi) — September

- **2026-09-21 — MiMo-V2.6 Pro/Flash (MIT).** Flash: 309B total / 15B active MoE, 1M context. (Chronology entry only — per coverage brief; Pro-tier specs not carried here.)

### Unsloth — quantization tooling, superseded version note

