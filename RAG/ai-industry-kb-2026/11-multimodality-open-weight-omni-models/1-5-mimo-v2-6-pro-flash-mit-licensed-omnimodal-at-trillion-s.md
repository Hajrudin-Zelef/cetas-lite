---
id: ai-industry-kb-2026/11-multimodality-open-weight-omni-models/1-5-mimo-v2-6-pro-flash-mit-licensed-omnimodal-at-trillion-s
title: "1.5 MiMo-V2.6 Pro / Flash — MIT-licensed omnimodal at trillion scale (Xiaomi, September 21, 2026)"
domain: multimodality-open-weight-omni-models
role: deep-dive
task: multimodal
actors: ["Anthropic", "China", "ExploitGym", "Google", "Huawei", "Hugging Face", "MiniMax", "OpenAI", "OpenRouter", "SGLang", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2025-12", "2026-04-02", "2026-06-01", "2026-06-03", "2026-08-20", "2026-08-23", "2026-09", "2026-09-21", "2026-09-22"]
keywords: ["license", "omni", "agentic", "apache", "ascend", "attention", "benchmarks", "claude", "compute", "consumer", "context window", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5802, 5847]
section: "11. Multimodality — Open-Weight Omni Models"
sha256: 4eceaf4bd4413ef14d0d1a48664a484536069dc02e09c4cf45f192faa69e8b13
---

# 1.5 MiMo-V2.6 Pro / Flash — MIT-licensed omnimodal at trillion scale (Xiaomi, September 21, 2026)

- **Release**: June 1, 2026; weights on Hugging Face by June 7 (`MiniMaxAI/MiniMax-M3`, plus `MiniMaxAI/MiniMax-M3-MXFP8`); technical report on arXiv June 11.
- **Scale**: **428B total / ~23B active** per token (256 fine-grained experts); **1M-token context** via MiniMax Sparse Attention (MSA) — per-token compute at 1M context ~1/20 of the previous generation; 9× prefill / 15× decode speedups vs M2. [VENDOR.]
- **Modalities**: **native text, image, video input** — trained on mixed text/image/video from step 0 (not bolted on); can operate a desktop from screenshots. **No audio input.**
- **License**: reported as **Apache 2.0** by several trackers (best-of-ai, Morph) but as **MiniMax Community License** (commercially restricted custom license) by others — treat as disputed; verify against the current HF repo before commercial use.
- **Benchmarks** [VENDOR]: **80.5% SWE-bench Verified**; 59% SWE-Bench Pro / 66% Terminal-Bench 2.1; matches Claude Sonnet 4.6 on real-world agentic benchmarks (Morph, checked Sept 7, 2026); **LMArena 1340**; Artificial Analysis Intelligence Index v4.3 (September 2026).
- **Positioning**: billed as the first open-weight release combining frontier coding, 1M context, and native multimodality. API: $0.30/$1.20 per 1M input/output tokens.
- **Sparse attention economics**: MiniMax Sparse Attention (MSA) holds per-token compute at 1M context to ~1/20 of the previous generation — the architectural reason a 428B model can serve 1M-token native-video prompts at all.
- **Sibling**: MiniMax H3 is the video-generation sibling (see §1.8) — cross-reference the two; M3 is the text-centric omni MoE, H3 the native video model.

### 1.5 MiMo-V2.6 Pro / Flash — MIT-licensed omnimodal at trillion scale (Xiaomi, September 21, 2026)

- **Release**: **September 21, 2026, 15:39 UTC** — checkpoints `XiaomiMiMo/MiMo-V2.6-Flash-RL` and `XiaomiMiMo/MiMo-V2.6-Pro-RL` published ungated on Hugging Face with technical report (notably without the usual vendor launch choreography).
- **Scale**: Pro **1.02T total / 42B active**; Flash **309B total / 15B active** (sparse MoE; Flash: 48 layers — 39 sliding-window + 9 global, hidden 4096, 256 routed experts / 8 active, 128-token sliding window, no shared experts).
- **Modalities**: **native omnimodal — text, image, video, audio**, 1M-token context, up to 128K output tokens.
- **Dedicated encoders** (all modalities enter one model): **681M-parameter MiMo ViT** (28 layers: 24 sliding-window + 4 full), **308M audio tokenizer**, **127M audio patch encoder**. Audio token rate ≈ **6.25 tokens/sec**; video via `fps` 0.1–10 (default 2) + `media_resolution`.
- **License**: **MIT** — commercially usable, fine-tunable, redistributable, no revenue gate, no research clause.
- **Training transparency**: Xiaomi livestreamed the RL run — Flash and Pro each completed 30 RL steps over ~750K trajectories in under 6 days, costing ~$850K / $2.62M (combined >$3M); 60K+ parallel Linux/Docker sandboxes; DeepSWE v1.1 rose 48.8→65.68 (Flash) and 58.4→72.57 (Pro). Training environments and RL code promised open-source but not yet shipped as of September 22, 2026.
- **Weights**: Flash published in **FP8 (e4m3)** — 172.9 GB across 65 shards (~1 byte/param); includes a DFlash-style multi-token-prediction drafter (model card says 5 layers; shipped config sets `num_nextn_predict_layers: 3` — trust the config, not the card, when sizing memory).
- **Benchmarks** (vendor-run, not independently reproduced): Intelligence Index v4.3 **46.32 (Pro)** — claimed top open model; DeepSWE v1.1 Flash 67.9 / Pro 71.9 (vs Claude Opus 5 74.0, GPT-5.6 Sol 73.0); Terminal-Bench 2.1 Flash **87.6** (vs Opus 5 89.1); AutomationBench Flash **52.3** (vs Opus 5 50.3, GPT-5.6 Sol 45.8); MiMo VisualCoding Flash 71.5 / Pro 72.3 (vs Opus 5 70.0); CyberGym Flash **95.1** (vs Pro 94.0) but weak on ExploitGym (6.0) / ExploitBench (25.3).
- **Serving**: card recommends **SGLang TP16/DP2** with EAGLE-style speculative path, or **vLLM TP8** (stable vLLM may lag the architecture) — a multi-node job; "Xiaomi removed the licence barrier and left the hardware barrier exactly where it was."
- **API pricing**: Flash $0.14/$0.28, Pro $0.435/$0.87 per 1M input/output tokens (uncached); cache-hit input $0.0028/$0.0036.

### 1.6 Gemma 4 — Apache 2.0 omni family, audio on compacts (Google DeepMind, April 2, 2026)

- **Release**: **April 2, 2026** (repos dated March 31); first Gemma generation under **Apache 2.0**; built from Gemini 3 research. 400M+ downloads of the Gemma family to date; 100K+ derivatives on HF.
- **Launch lineup** (all natively multimodal, 140+ languages):
  - **E2B**: ~2.3B effective params, 128K ctx — **text/image/audio** (phones, Raspberry Pi, browsers).
  - **E4B**: ~4.5B effective, Per-Layer Embeddings (PLE), 128K ctx — **text/image/audio** (laptops, mini-PCs).
  - **26B-A4B MoE**: 26B total / ~4B active (25.2B/3.8B per some cards), 256K ctx — text/image (single consumer GPU).
  - **31B Dense**: 30.7B, 256K ctx — text/image; ranked **#3 globally among open models** on the Arena AI text leaderboard; Codeforces ELO 110 → 2150 reported. [VENDOR — codeforces figure.]
- **Audio on compacts confirmed**: the two smallest variants accept native audio input (speech recognition) — a first for open models at this parameter count; larger variants are text/image (+video/OCR inputs per Google's model card).
- **Gemma 4 12B Unified** (June 3, 2026): reported as **encoder-free** multimodal — text, image, **audio, video** without separate encoder networks; runs on 16 GB RAM/VRAM laptops; Apache 2.0. [UNVERIFIED — primary-source confidence: exact per-modality coverage varies across secondary sources; reported modalities range from text/image/audio to full text/image/audio/video.]
- **Architecture notes**: alternating local sliding-window + global full attention; PLE (parallel lower-dimensional conditioning pathway per layer) drives the intelligence-per-parameter ratio.

### 1.7 GLM-5.3-Flash ("Ox Alpha") — first native multimodal GLM (Zhipu/Z.ai, August 26–27, 2026)

- **Identity**: stealth model `stealth/ox-alpha` on OpenRouter (August 20, 2026, free for a week) confirmed by Zhipu to Bloomberg on August 26 as **GLM-5.3-Flash**; weights released August 27 under **MIT** (`zai-org/GLM-5.3-Flash`).
- **Scale**: **320B total / 18B active** MoE; **1M-token context** (131K output limit); hybrid sparse + linear attention (first in the GLM main series); pretrained on 30T multimodal tokens.
- **Modalities**: **natively multimodal — text, image, video**; audio support for the 1M-context window appears only in inconsistent launch coverage. [UNVERIFIED — audio input is not consistently reported; text/image/video is well supported, audio should not be asserted without qualification.]
- **Training compute**: served its 100-trillion-tokens/day trial week entirely on **~100,000 domestic (Chinese) chips** — likely Huawei Ascend.
- **Benchmarks** (vendor): DeepSWE 63.4 (model card); "approaches Claude Opus 4.8" on coding/agentic; Ed Yau's independent harness (August 23, 2026): **100% pass rate / 9.3/10 rubric** across 28 practical tasks at $0.28 total cost vs $1.43 for GPT-5.5.
- **Serving**: deployable on **vLLM, SGLang, KTransformers**; KV cache 4.4× smaller than GLM-5.3. API: $0.15/$0.50 per 1M in/out (launch discount $0.075/$0.25 through Sept 9).
- **Sibling**: GLM-5.3 flagship 744B/40B (weights opened August 28, new GLM-5.3 License; SWE-bench Pro 77.8% [VENDOR]) is text-focused; GLM-4.6V-Flash (December 2025) and GLM-4.5V cover the vision line; GLM-Image / GLM-ASR / GLM-TTS are separate modality models.

### 1.8 Other notable 2026 omni / multimodal open releases

