---
id: ai-industry-kb-2026/11-multimodality-open-weight-omni-models/main-actors
title: "Main actors"
domain: multimodality-open-weight-omni-models
role: deep-dive
task: multimodal
actors: ["AWS", "Alibaba", "ByteDance", "China", "DeepSeek", "Google", "Huawei", "MiniMax", "Mistral", "Moonshot", "Nvidia", "SGLang", "StepFun", "Unsloth", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2026-09"]
keywords: ["agentic", "apache", "ascend", "awq", "aws", "compute", "cost", "deepseek", "diffusion", "fp8", "gemini", "gguf"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5983, 6025]
section: "11. Multimodality — Open-Weight Omni Models"
sha256: cbb59cd818466b1316b310da537ab226efda85530e838ca2ff58298f9a888430
---

# Main actors

- **MiMo-V2.6 RL cost (disclosed by Xiaomi, Sept 2026)**: Flash ~$850K, Pro ~$2.62M (combined >$3M); 30 steps each over ~750K trajectories in <6 days; 60K+ parallel Linux/Docker sandboxes; DeepSWE v1.1 rose 48.8→65.68 (Flash) and 58.4→72.57 (Pro). First livestreamed trillion-scale RL run — the training process, not just the weights, is becoming an open artifact.
- **GLM-5.3-Flash training/serving compute**: 100 trillion tokens/day trial week served entirely on ~100,000 domestic (Chinese) chips (likely Huawei Ascend).
- **Token rates**: MiMo audio ~6.25 tok/sec (10 hours of audio ≈ 225K tokens — why 1M-context omni models can ingest very long audio natively); DeepSeek V4 Flash images capped at 384 tokens/image; SGLang `--mm-process-config` defaults (image max_pixels 1,048,576; video fps 3, max_pixels 602,112, max_frames 60); MiMo video billed as fps×resolution.
- **Latency**: Qwen3-Omni first-packet audio ~211–234 ms (audio-video ~507 ms); fal H3 Max renders 5s 768p in <3s (hosted); community NVFP4 H3 on RTX 5090 ~175s per 10s clip.
- **Pricing (per 1M tokens, uncached unless noted)**: MiniMax M3 $0.30/$1.20; MiMo-Flash $0.14/$0.28, MiMo-Pro $0.435/$0.87 (cache-hit input $0.0028/$0.0036); GLM-5.3-Flash $0.15/$0.50 (launch discount $0.075/$0.25 through Sept 9); DeepSeek V4.1-Flash peak $0.30/$1.20, cache-read $0.006, off-peak half (50× hit/miss input gap); Gemini Omni Flash $0.10/sec 720p (closed, §1 cross-ref); H3 Max 5s clip $0.10→$0.40 after Sept 15 discount expiry.
- **KV-cache engineering**: DeepSeek V4.1-Flash 890 bytes/token (~4× vs V4-Flash via MXFP4 KV + CSA2; ~437× only against a 2023 4K-context V1 denominator); 1M-token context ≈ 0.93 GB KV cache; 8× H200 → ~915 concurrent full-1M sessions (vs ~40 for DeepSeek V3) — secondary analysis, unverified independently.
- **Speech tokenization rates**: discrete neural codecs at ~25 Hz semantic tokens (CosyVoice v1 VQ, v2 FSQ, 6561 speech tokens → flow-matching → 24 kHz); Qwen3-Omni multi-codebook Talker with 12 Hz acoustic compression + causal ConvNet/Code2Wav decoder replacing iterative diffusion; VoxCPM2 (MiniMax) tokenizer-free TTS preserving prosodic nuance.
- **VRAM recipes (September 2026)**:
  - Edge/phone (≤4 GB): Gemma 4 E2B-QAT 4.3 GB (Ollama `gemma4:e2b-it-qat`); text/image/audio.
  - Laptop 8–16 GB: Gemma 4 E4B-QAT 4.3 GB, Gemma 4 12B-QAT 7.2 GB, Qwen3-VL-8B 4-bit ~6 GB (MLX), Qwen3.5-9B Q4_K_M 6.6 GB — native vision, some audio.
  - Workstation (24–48 GB): Qwen3-VL-30B-A3B AWQ 4-bit (~10–12 GiB/GPU at TP2, 120–150 tok/s, ~25% below text-only from vision-encoder overhead), Qwen3-Omni-30B-A3B via vLLM-Omni (≥48 GB total, TP2); Unsloth Dynamic 2.0 GGUFs for Qwen3.x/GLM-5.3 at IQ2–Q4.
  - Server (multi-node): MiMo-V2.6-Flash 172.9 GB FP8 (SGLang TP16/DP2 or vLLM TP8), MiniMax-M3-MXFP8, Kimi K3 MXFP4 (~1.4 TB, 64+ accelerators).
  - API-only: Qwen3.8-Omni-Flash (1M ctx, text-out), DeepSeek-V4-Flash-vision-exp (384 tokens/image cap, now retired), MiMo Pro on Xiaomi API.
- **The memory equation**: total ≈ weights + KV cache + activations + **modality encoders** (vision/audio towers are a separate, often overlooked line item; SGLang's default offloads their features to CPU to save VRAM). Encoder quantization is conservative — towers are usually kept at higher precision than the LLM backbone; mmproj sidecars ship alongside GGUF quants rather than inside them. KV cache at 1M context remains the binding constraint (GQA/MLA, shared KV, HybridKV, TurboQuant 3-bit KV ÷6 memory).

## Main actors

- **Alibaba (Qwen team)** — Qwen3-Omni-30B-A3B (Apache 2.0, reference design), Qwen3.5-Omni (Plus/Flash/Light), Qwen3.8-Omni-Flash (API-only, 1M ctx, agentic video perception); Qwen3-VL family; Qwen3-TTS / CosyVoice v1/v2 (25 Hz semantic tokens, FSQ in v2).
- **Xiaomi (MiMo team)** — MiMo-V2.6 Pro (1.02T/42B, MIT) and Flash (309B/15B, MIT); dedicated 681M ViT + 308M audio tokenizer + 127M audio patch encoder; livestreamed $3M+ RL training; MiMo-Audio (vLLM-Omni coverage).
- **MiniMax** — MiniMax M3 (428B/23B, native text/image/video, license disputed); **MiniMax H3 / Hailuo 3.0 / Hailuo 03** (33B dense, first open model to top an AI video ranking; territorially restricted license); VoxCPM2 tokenizer-free TTS.
- **Google DeepMind** — Gemma 4 family (Apache 2.0): E2B/E4B (audio input at phone/laptop scale), 26B-A4B, 31B (#3 open on Arena), 12B Unified (claimed encoder-free).
- **Zhipu / Z.ai** — GLM-5.3-Flash (MIT, first native multimodal GLM; trained/served on domestic Chinese chips); GLM-5.3 flagship (text); GLM-4.6V-Flash / GLM-4.5V (vision); GLM-Image / GLM-ASR / GLM-TTS (separate modality models).
- **Moonshot AI** — Kimi K3 (2.8T/104B, native vision via MoonViT-V2, bespoke license); Kimi K2.5 (native MoonViT + PatchMerger in SGLang); Kimi-Audio (separate audio model).
- **DeepSeek** — V4 Flash (284B/13B, MIT; experimental vision retired Sept 10, 2026) and V4.1-Flash (552B MoE + CED + MXFP4 KV, MIT, reasoning dial, cache-price engineering).
- **Thinking Machines** — Inkling-Small (open weights, text/image/audio natively, SGLang + MTP speculative decoding).
- **StepFun** — Step3-VL-10B (Jan 13, 2026; strided-conv downsamplers).
- **Mistral** — Ministral 3 8B (Apache 2.0, multimodal, 6 GB at 8-bit).
- **RedNote** — dots3-note (vision+audio in llama.cpp #27524).
- **HiDream.ai (Beijing)** — HiDream-O1-Video-1.0 (closed weights, company-reported native omnimodal video).
- **ByteDance** — Seedance 2.5 (closed, reported 30s clips with built-in audio; single secondary mention, §12).
- **Serving ecosystem** — vLLM-Omni (`vllm-project/vllm-omni`, stable since v0.14.0 Jan 2026; AWS DLC `vllm:omni-cuda`; see §6 one-line cross-ref); SGLang (day-0 coverage for new MoE/multimodal releases; sglang-jax standalone multimodal subsystem; `--mm-process-config`); llama.cpp/mtmd (mature vision, experimental audio, nascent video); MLX (mlx-vlm, mlx-community, mlx-audio); Unsloth (GGUF catalog, Dynamic 2.0 quants); Ollama (one-command local multimodal); ComfyUI (H3 day-0); fal (H3 Max hosting); NVIDIA VoiceChat-11B / BAGEL (unified/open multimodal lineage, Wave 2 context).

## Timeline and context

### How 2026 omni models handle each modality — architecture patterns

**Vision encoders and vision→LLM mergers (early vs late fusion).**
- **Dedicated ViTs**: MiMo ViT 681M (28 layers), MoonViT / MoonViT-V2 (Kimi), SigLIP/AIMv2-style encoders in most VLMs.
- **Merger designs**: **pixel-shuffle** (InternVL, GLM-4.6V-Flash), **strided-conv downsamplers** (Step3-VL: two Conv2d stages, 4× total), **Perceiver resamplers** (MiniCPM-V: 64 learned queries), **PatchMerger** kernels (Kimi K2.5 in SGLang), **projector MLPs** (LLaVA-style).
- **Encoder-free**: Gemma 4 12B Unified processes all modalities without separate encoder networks — the emerging "unified" pole vs the modular encoder+projector mainstream. [UNVERIFIED coverage.]
- **Late-fusion survives inside unified models**: MiniMax H3's nominally single-stream "H3-Omni-Transformer" uses the full pretrained weights of **Qwen3-VL-32B** as its text encoder (layer-50 hidden states) — the same modularity-survives-where-it-pays pattern documented for BAGEL and NVIDIA VoiceChat-11B.
- **Early vs late fusion, 2026 verdict**: early fusion (MiniMax M3, MiMo-V2.6 trained on mixed modalities from step 0) gives the strongest cross-modal reasoning at the heaviest training cost; late fusion (modular encoder + projector, e.g. Qwen3-VL, InternVL, most GGUF quants) gives the best local-deployment story. A 457-model early-fusion scaling-laws study and FuseLIP are documented in Wave 2.

