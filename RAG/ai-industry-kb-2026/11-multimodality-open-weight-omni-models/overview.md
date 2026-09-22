---
id: ai-industry-kb-2026/11-multimodality-open-weight-omni-models/overview
title: "11. Multimodality — Open-Weight Omni Models"
domain: multimodality-open-weight-omni-models
role: deep-dive
task: multimodal
actors: ["Alibaba", "Anthropic", "Apple", "ByteDance", "China", "DeepSeek", "EU", "ExploitGym", "Google", "Huawei", "Hugging Face", "MiniMax", "Mistral", "Moonshot", "OpenAI", "OpenRouter", "SGLang", "StepFun", "United States", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2025-09", "2025-12", "2026-01", "2026-01-13", "2026-03-30", "2026-04-02", "2026-06-01", "2026-06-03", "2026-07-31", "2026-08", "2026-08-20", "2026-08-23", "2026-09", "2026-09-10", "2026-09-15", "2026-09-17", "2026-09-18", "2026-09-21", "2026-09-22"]
keywords: ["multimodal", "omni", "open-weight", "agentic", "apache", "ascend", "attention", "awq", "benchmark", "benchmarks", "claude", "compute"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5756, 5892]
section: "11. Multimodality — Open-Weight Omni Models"
sha256: ff0f406987c86ddf313b42b40abd1c38040d3ff1aefda520966b84653d77a92f
---

# 11. Multimodality — Open-Weight Omni Models
Keywords: open-weight omni models, native multimodality, Qwen3-Omni, MiniMax M3, MiniMax H3, MiMo-V2.6, Gemma 4, GLM-5.3-Flash, Kimi K3, HiDream-O1-Video-1.0, Thinker-Talker, vision encoders, audio tokenizers, video understanding, any-to-any scoreboard, vLLM-Omni, SGLang, llama.cpp mtmd, VRAM, quantization, benchmarks, open vs closed gap

## Summary

By September 22, 2026 the open-weight ecosystem completed its move from "vision bolted onto a text model" to **native omni models** trained on mixed text/image/video/audio from step zero. Four flagship releases define the native open-weight omni landscape:

1. **Qwen3-Omni-30B-A3B** (Alibaba, September 2025) — the reference open omni design, Apache 2.0, Thinker-Talker MoE, text/image/audio/video in and streaming speech out.
2. **MiniMax M3** (June 1, 2026) — 428B/23B MoE, native text/image/video from step 0, 1M context; no audio input; license disputed (Apache 2.0 reported by some trackers, MiniMax Community License per others).
3. **MiMo-V2.6 Pro/Flash** (Xiaomi, September 21, 2026) — first MIT-licensed omnimodal models at 1.02T/42B (Pro) and 309B/15B (Flash), dedicated vision + audio encoders, 1M context, checkpoints published ungated on Hugging Face with a livestreamed RL training run.
4. **Gemma 4** (Google DeepMind, April 2, 2026) — Apache 2.0 family; the two compact variants (E2B/E4B) accept native audio input, a first at that parameter count; **Gemma 4 12B Unified** (June 3, 2026) claims an encoder-free text/image/audio/video design. [UNVERIFIED — per-modality coverage varies across secondary sources.]

A fifth native-multimodal open release joined in August 2026: **GLM-5.3-Flash** (Zhipu, 320B/18B, MIT, 1M context), stealth-tested as "Ox Alpha". The benchmark picture is two-sided: open omni models now beat closed models on most audio understanding tasks (Qwen3-Omni: overall SOTA on 22 of 36 audio benchmarks versus Gemini-2.5-Pro and GPT-4o-Transcribe [VENDOR]) but trail 10–17× on hard agentic visual reasoning (Scale Labs VTB: GPT-5-think 18.44% APR vs open 1.16–1.65%). The serving stack consolidated around **vLLM-Omni** (stable since January 2026), **SGLang** (day-0 coverage of new releases), **llama.cpp/mtmd** (mature vision, experimental audio, nascent video), and the **MLX** ecosystem on Apple Silicon.

September 2026 added the first open-weight entries on the **any-to-any** board: **MiniMax H3** (Hailuo 3.0 = Hailuo 03 — one model, three names; open weights August 2026) became the first open model to top an AI video ranking (#1 Video Editing on Artificial Analysis) with native synchronized stereo audio output — under a restrictive community license excluding the US, EU, UK, and South Korea from local deployment. **HiDream-O1-Video-1.0** (September 17, 2026, company-reported) claims native omnimodal video generation with synchronized audio but is closed-weights; closed frontier omni models (Gemini Omni etc.) are covered in §1 and video-generation models in §12 — this section covers only the open-weight omni lineage and the September 2026 any-to-any entrants with their licensing reality.

### 1.1 Qwen3-Omni-30B-A3B — the reference open omni design (Alibaba, September 2025)

- **Release**: September 2025 (announced ~September 23–25, 2025); weights and code on GitHub (`QwenLM/Qwen3-Omni`) under **Apache 2.0**; free commercial use.
- **Modalities**: text, images, audio, video **in**; text plus **streaming speech** out (real-time voice).
- **Architecture**: **Thinker-Talker MoE** — 30B total / 3B active per token. The *Thinker* handles perception/reasoning across all modalities; the *Talker* autoregressively predicts discrete speech codecs with a **multi-codebook scheme** and replaces block-wise diffusion with a lightweight **causal ConvNet**, enabling streaming from the first codec frame.
- **Latency**: ~211–234 ms end-to-end cold-start first-packet latency for audio; ~507 ms for audio-video tasks. [VENDOR — technical report.]
- **Languages**: text in 119+ languages; speech understanding in 19; speech generation in 10 languages/dialects (expanded in successors).
- **Benchmarks** [VENDOR, technical report]: open-source SOTA on **32 of 36** audio and audio-visual benchmarks; **overall SOTA on 22 of 36**, outperforming closed models Gemini-2.5-Pro, ByteDance Seed-ASR, and GPT-4o-Transcribe. Competitive with Gemini 2.5 Pro on mixed-media understanding and voice output.
- **Variants**: `Qwen3-Omni-30B-A3B`, `-Thinking` (explicit reasoning over any modality), `-Captioner` (fine-tuned audio captioner, low-hallucination detailed captions).
- **Serving**: supported in **vLLM-Omni** (`vllm serve Qwen/Qwen3-Omni-30B-A3B-Instruct --omni`); requires **≥48 GB total GPU memory** across tensor-parallel workers (TP=2 typical); expert parallelism for multi-node.

### 1.2 Qwen3.5-Omni — scaled-up successor (Alibaba, March 30, 2026)

- **Release**: March 30, 2026.
- **Variants**: Plus (30B-A3B MoE), Flash (lightweight MoE), Light (dense model, **open weights**).
- **Context**: 256K tokens; **10+ hours of continuous audio**; **400+ seconds of 720p video at 1 FPS sampling**.
- **Languages**: speech recognition in **113 languages/dialects** (up from 19); speech generation in 36 (up from 10).
- **Training**: 100M+ hours of audio/video data; SOTA claimed on 215 audio/video understanding benchmarks. [VENDOR.]
- **Emergent capability**: "Audio-Visual Vibe Coding" reported by the team.

### 1.3 Qwen3.8-Omni-Flash — 1M-context agentic omni, API-only (Alibaba, September 18, 2026)

- **Release**: September 18, 2026, via Qianwen AI Platform / Alibaba Cloud Model Studio / Qwen Studio.
- **Modalities**: text, image, audio, video in; **text out only** (speech generation delegated to Qwen3.5-Omni; media production via tool calls — Qwen-MM-Plugins, Qwen-Live Harness).
- **Context**: 1M tokens (991K max input / 131K max output listed on QwenCloud; 262K max reasoning length); thinking on by default (`reasoning_effort: xhigh`, settable to `none`).
- **Key idea**: **agentic perception for long video** — the model decides what to watch/hear via coarse-to-fine evidence gathering instead of reading the whole file. On **OmniVideoBench**: accuracy 63.4 → **67.8**, token use 145,736 → **79,117** (−45.7%). [VENDOR.]
- **Open weights**: **none announced at launch — self-hosting is not an option** (important: the Qwen3.8-Omni line is currently the exception to Alibaba's open-weights omni tradition).

### 1.4 MiniMax M3 — native video from step 0 (MiniMax, June 1, 2026)

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

- **Kimi K3** (Moonshot AI, mid-2026): **2.8T total / 104B active** "Stable LatentMoE" (896 experts, 16 active), 1M ctx (Kimi Delta Attention), **native vision via MoonViT-V2**; MXFP4 weights / MXFP8 activations (~1.4 TB resident); Intelligence Index **60** (#1 open weight at publication, trails Claude Opus 5 at 63); LiveBench Coding 81.45 / Agentic Coding 57.58; Kimi K3 License (bespoke, MaaS resale gate); needs 64+ accelerators to self-host. (Vision confirmed; audio input not confirmed for K3 — Kimi-Audio is a separate model.)
- **Kimi K2.5** (Moonshot AI): native MoonViT + PatchMerger kernels in SGLang; "native multimodal vision" per Moonshot.
- **DeepSeek V4 Flash** (284B/13B MoE, MIT): experimental vision variant `deepseek-v4-flash-vision-exp` (API) — images capped at **384 tokens/image** regardless of resolution, JPEG/PNG/GIF/WebP, up to 600 images/request. Retired September 10, 2026: model IDs now route to DeepSeek V4.1-Flash, which folds multimodal into the base model (45T multimodal tokens, 7:1 text-to-multimodal). [VENDOR.]
- **DeepSeek V4.1-Flash** (September 10, 2026): MIT open weights (48 safetensors shards, ~510 GB); **MXFP4 KV cache at 890 bytes/token** (E2M1, NVFP4 standard) via CSA2 — ~4× reduction vs V4-Flash (the "437×" headline uses a 2023 4K-context V1 denominator); built-in reasoning dial (effort 25→100: 67.1%→76.3% reasoning average at ~2.5× output tokens); 1M-token context costs ~0.93 GB of KV cache. Price card: peak $0.30/$1.20 per million uncached in/out tokens; cache-read $0.006; off-peak half; ~50× gap between cache-hit and cache-miss input — engineered to force prefix reuse in agentic workloads. [VENDOR — architecture/price per DeepSeek's report; the 437× denominator critique is secondary analysis.]
- **Qwen3-VL family** (Alibaba): 8B (DocVQA **96.1%**, OCRBench ~**896**, ScreenSpot **94.4%**, 256K ctx), 30B-A3B (AWQ-4bit ~19 GB on disk; vLLM TP=2, ~10–12 GiB/GPU + KV), 235B-A22B open weights; **Qwen3.5-9B** (Apache 2.0, native vision, hybrid attention, small KV) fits 8 GB VRAM at Q4_K_M.
- **Ministral 3 8B** (Mistral): Apache 2.0, **multimodal**, 256K ctx, 6.0 GB at 8-bit.
- **Thinking Machines Inkling-Small**: open weights (BF16 + NVFP4), **text/image/audio natively**, variable reasoning effort; served on SGLang with MTP speculative decoding.
- **dots3-note** (RedNote): vision+audio support added to llama.cpp (#27524).
- **Step3-VL-10B** (StepFun, January 13, 2026): 10.2B dense, 1.8B ViT + two-stage strided-conv downsamplers (4×), tech report arXiv 2601.09668.
- **Ovis2.5-9B**, **InternVL3.5-8B**, **GLM-4.6V-Flash**: established open VLM lines still current in 2026.

### 1.9 New any-to-any entrants (September 2026): MiniMax H3 and HiDream-O1-Video-1.0

- **MiniMax H3 = Hailuo 3.0 = Hailuo 03 — ONE model, three names.** (Naming discipline matters for the RAG: trackers, the HF repo, and press use the three names interchangeably; they refer to the same release.)
- **Release**: API launch July 31, 2026 (model ID `MiniMax-H3`, Hailuo AI app); open weights on Hugging Face **August 2–3, 2026** (sources differ by one day). Day-0 ComfyUI support; community produced GGUF (14.2 GB Q4), INT4, NVFP4 quants and an MLX 8-bit port within days — reportedly the fastest open-weights-to-consumer-hardware pipeline tracked in 2026. [COMMUNITY — quant availability.]
- **Architecture**: 33B-parameter dense single-stream "H3-Omni-Transformer"; jointly understands text, images, video, audio and generates video with **native 32 kHz stereo audio in one pass** (4–15s clips, 24fps, 768p local canvas / 2K via hosted API). Three modules: H3-Context-IR (input preprocessing), H3-Base (generation), H3-Regenerate-2K (upscaler). Text understanding uses the full pretrained weights of **Qwen3-VL-32B** as the text encoder (hidden states from layer 50) — a late-fusion-style pretrained encoder inside a nominally unified model (modularity survives where it pays).
- **Generation modes**: FL2VA (text-to-video with optional first/last-frame conditioning) and Ref2VA (reference-to-video: up to 9 images, 3 videos, 3 audio clips per prompt). Instruction-based editing in plain language.
- **Results**: #1 Video Editing (With Audio), #2 Text-to-Video, #3 Image-to-Video on Artificial Analysis; ~1476 Arena.ai image-to-video Elo; described by the-decoder as the **first open model to top an AI video ranking**. [COMMUNITY — leaderboard-sourced.]
- **Deployability**: official BF16 SGLang recipe targets 4 GPUs; community NVFP4 quants run on a single RTX 5090 (~175s per 10s clip); ComfyUI maintainer demonstrated renders on an RTX 3060. fal's post-trained **H3 Max** variant renders 5s 768p in under 3s; its 75% launch discount expired September 15, 2026, quadrupling the 5s-clip price $0.10→$0.40.
- **License (critical)**: "MiniMax H3 Community License Agreement" — open weights, not Apache. Territorial exclusion: **local deployment rights exclude the US, EU, UK, and South Korea**; companies above $20M/year revenue need separate written authorization; training other models on H3 outputs is prohibited. Two pieces remain closed: the 2K upscaling module and H3-Context-IR. For the RAG: this forces a *territorially-restricted community license* category in the license taxonomy — "open weights" alone is no longer a sufficient label.
- **HiDream-O1-Video-1.0 ("HiDream V1")** (HiDream.ai, Beijing, September 17, 2026): announced "native omnimodal video generation model" — text/image/video in; 1080p clips of 5–20 seconds with **natively synchronized audio** (dialogue with lip-sync, timed SFX, ambience); duration is part of narrative planning (model selects 5–20s based on event unfolding). Company claims a unified **UiT (Unified Transformer)** foundation jointly modeling text, video, audio, post-trained via **Diffusion Reinforcement Learning** with a multimodal reward model. Benchmarks company-reported only: #4 on the Artificial Analysis Image-to-Video Leaderboard (With Audio); #8 on Arena.ai Image-to-Video. **Independent reproduction: none found as of September 22, 2026 — treat all figures as company-reported.** [VENDOR/UNVERIFIED.]
- Dedup note: full video-generation model coverage (Sora, Seedance, Kling, Genie) lives in §12; H3 and HiDream-O1 appear here only as the open-weight any-to-any entrants touching the omni lineage. Closed frontier omni models (Gemini Omni Flash etc.) → §1.

### 1.10 License and provenance caveats

- **MiniMax M3 license is disputed**: best-of-ai and Morph report **Apache 2.0**; the agentone research file reports **MiniMax Community License** (commercially restricted). The HF repo (`MiniMaxAI/MiniMax-M3`) is the tiebreaker — verify the current license file before commercial use.
- **MiniMax H3 license**: territorially restricted — US, EU, UK, South Korea excluded from local deployment; not Apache-grade despite "open weights".
- **Qwen3.8-Omni-Flash**: no open weights were announced at launch (September 18, 2026); do not classify the 1M-context omni flagship as open-weight. The Qwen omni line otherwise remains open (Qwen3-Omni Apache 2.0, Qwen3.5-Omni Light dense open).
- **Kimi K3**: open weights under a bespoke Kimi K3 License with a MaaS resale gate — downloadable but not fully open-source in the OSI sense.
- **Vendor vs independent**: MiMo-V2.6, MiniMax M3, Qwen3.5-Omni, GLM-5.3-Flash, DeepSeek V4.1-Flash, HiDream-O1-Video-1.0 headline numbers are vendor-run. Independent confirmations exist for GLM-5.3-Flash (Ed Yau: 100% pass / 9.3/10 over 28 tasks, $0.28 vs $1.43 for GPT-5.5), MiniMax M3 (Morph fact-check Sept 7, 2026), Gemma 4 31B (LMArena #3). MiMo-V2.6's benchmark independence is pending as of September 22, 2026.
- **"Native" vs "modular"**: "native omni" should not be read as "single encoder" — MiMo-V2.6 (ViT + audio tokenizer + patch encoder), MoonViT-family models, Qwen3-Omni (Thinker/Talker), and MiniMax H3 (Qwen3-VL-32B as text encoder) are all natively multimodal *and* multi-encoder. Only Gemma 4 12B Unified claims an encoder-free single-transformer design.

### 1.11 Any-to-any scoreboard update (September 2026)

Wave 1 §1-era snapshot named Gemini Omni Flash as the only closed any-to-any model. September 2026 additions change the board:

| Model | Lab | Access | Inputs | Outputs | Status / caveat |
|---|---|---|---|---|---|
| Gemini Omni Flash | Google | Closed API | text, image, audio, video | video + synchronized audio (10s clips) | Still the benchmark-setter; $0.10/sec 720p (see §1) |
| HiDream-O1-Video-1.0 | HiDream.ai | API (closed weights) | text, image, video | 1080p video (5–20s) + synchronized audio | Company-reported only (Sept 17) [VENDOR] |
| Seedance 2.5 | ByteDance | Closed | text (reported) | 30s clips with built-in audio | Single secondary mention [UNVERIFIED]; full coverage in §12 |
| MiniMax H3 | MiniMax | Open weights (restrictive license) | text, image, video, audio | video + native stereo audio (4–15s, 768p local) | First open model to top an AI video ranking; US/EU/UK/KR excluded from local deployment |

The open side now has a genuine any-to-any video entry on at least one public leaderboard, but under a license that disqualifies it as Apache-grade open weights. The thesis "closed labs lead productized any-to-any output" survives, narrowed: it is now a *productization-and-licensing* lead, not a capability lead.

