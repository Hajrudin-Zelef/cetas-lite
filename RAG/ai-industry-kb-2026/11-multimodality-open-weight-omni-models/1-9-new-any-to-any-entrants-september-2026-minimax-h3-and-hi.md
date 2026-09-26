---
id: ai-industry-kb-2026/11-multimodality-open-weight-omni-models/1-9-new-any-to-any-entrants-september-2026-minimax-h3-and-hi
title: "1.9 New any-to-any entrants (September 2026): MiniMax H3 and HiDream-O1-Video-1.0"
domain: multimodality-open-weight-omni-models
role: deep-dive
task: multimodal
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "EU", "Google", "Hugging Face", "MiniMax", "Mistral", "Moonshot", "SGLang", "StepFun", "United States", "Z.ai", "vLLM"]
dates: ["2026-01-13", "2026-07-31", "2026-09", "2026-09-10", "2026-09-15", "2026-09-17", "2026-09-22"]
keywords: ["agentic", "apache", "attention", "awq", "benchmarks", "claude", "consumer", "deepseek", "diffusion", "gemini", "gguf", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5848, 5872]
section: "11. Multimodality — Open-Weight Omni Models"
sha256: ed6d9107b08a779fcd6e4c658f6b6a56d84951165c3fa23ee5932285afb129cb
---

# 1.9 New any-to-any entrants (September 2026): MiniMax H3 and HiDream-O1-Video-1.0

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

