---
id: ai-industry-kb-2026/11-multimodality-open-weight-omni-models/main-actors
title: "Main actors"
domain: multimodality-open-weight-omni-models
role: deep-dive
task: multimodal
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "Apple", "ByteDance", "China", "DeepSeek", "Google", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "SGLang", "StepFun", "Unsloth", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2026-01", "2026-09"]
keywords: ["agent", "agentic", "agents", "apache", "aws", "benchmarks", "claude", "cost", "deepseek", "diffusion", "embeddings", "fine-tuning"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5998, 6069]
section: "11. Multimodality — Open-Weight Omni Models"
sha256: 6f98fd54494b8e081b98f97054b03239720696e13c7fc91e6891fb1e49e40ad8
---

# Main actors

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

**Audio: tokenization approaches.**
- **Discrete neural codecs** (dominant): VQ (CosyVoice v1) / **FSQ (v2)** quantization bottlenecks, ~25 Hz semantic tokens (6561 speech tokens; LLM Qwen2.5-0.5B → flow-matching → HiFi/HiFT vocoder at 24 kHz).
- **Multi-codebook** (Qwen3-Omni/Qwen3-TTS): Talker predicts **multiple codebooks** autoregressively; 12 Hz tokenizer for acoustic compression; lightweight **causal ConvNet / Code2Wav** decoder replaces iterative diffusion — the key to 211–234 ms streaming first-packet latency.
- **MiMo approach**: separate **308M audio tokenizer** + **127M audio patch encoder** feeding the same backbone as vision/text; API token rate ≈ 6.25 tok/sec.
- **Whisper-style encoders**: still used for ASR-side understanding in llama.cpp's mtmd (Whisper, MERaLiON-2) and MLX audio pipelines.
- **Tokenizer-free** (TTS frontier): VoxCPM2 skips discrete codec tokens entirely to preserve prosodic nuance.
- Honest summary (sota-reference, Sept 2026): open models closed most of the gap on short-clip understanding; everyone — open and closed — is weak on long-audio, temporal, and multi-hop reasoning. Closed frontier audio leaders (MMAR): Gemini and GPT-4o-audio still lead; open Qwen2.5-Omni and Kimi-Audio are "within striking distance"; Audio Flamingo 3 (developer-reported) beats Gemini Pro 1.5/2.5 and GPT-4o-audio on its suite.

**Video understanding.**
- **Frame sampling is universal**: 1 FPS (Qwen3.5-Omni: 400+ s of 720p), 0.1–10 FPS configurable (MiMo), fps 3 default in SGLang.
- **Agentic perception** (Qwen3.8-Omni-Flash): question-driven coarse-to-fine evidence gathering over long video — 63.4→67.8 on OmniVideoBench at −45.7% tokens. This is the 2026 answer to the "watch the whole file" bottleneck; it stays API-only (no weights).
- **Native video input** (no frame-extraction hacks): MiniMax M3, MiMo-V2.6, GLM-5.3-Flash, Gemma 4 12B Unified, Qwen3-Omni/3.5-Omni.
- **Engine support**: llama.cpp added `--video-*` CLI flags (#24318) and exports video helper symbols, but end-to-end video input is **not yet qualified** in current releases — needs `LLAMA_SUBPROCESS`/`MTMD_VIDEO` builds plus FFmpeg/ffprobe packaging; the documented workaround is extracting frames and sending them as images.

**Streaming speech output (Talker architectures).**
- **Talker → Code2Wav** (Qwen3-TTS): multi-codebook discrete token prediction → non-DiT parallel decoder → waveform; replaces iterative diffusion for low first-packet latency.
- **Thinker-Talker split** is the standard pattern for low-latency voice: Thinker (perception/reasoning) feeds Talker (streaming synthesis).
- **Implication for RAG**: streaming speech out (Qwen3-Omni 211–234 ms first packet) makes open-weight voice agents deployable without closed TTS APIs — the Talker stage, not the Thinker, is the latency-critical path. Full-duplex speech state is covered in Wave 2/§1-adjacent sections.

### Benchmarks: open omni vs closed frontier (2026) — the pattern

- **Audio understanding — open leads or ties**: Qwen3-Omni 22/36 overall SOTA vs closed [VENDOR]; peer-reviewed MMAR still favors Gemini/GPT-4o-audio but open models are within striking distance.
- **Image/document understanding — open competitive**: Qwen3-VL-8B (DocVQA 96.1%, OCRBench ~896, ScreenSpot 94.4%); MiniMax M3 the only >80% SWE-bench-Verified model with native image+video reading; MiMo VisualCoding Flash/Pro 71.5/72.3 vs Claude Opus 5 70.0 [VENDOR]; Gemma 4 31B #3 open on Arena.
- **Video — agentic methods move the needle**: Qwen3.8-Omni-Flash 63.4→67.8 on OmniVideoBench at −45.7% tokens [VENDOR, API-only]; Qwen3.5-Omni claims SOTA on 215 audio/video benchmarks [VENDOR].
- **Hard agentic visual tasks — the remaining gap**: Scale VTB — GPT-5-think 18.44% APR vs open 1.16–1.65% (10–17×); 70–82% of failures are visual perception errors (the bottleneck is perception, not logic). Takeaway: on standard VQA/OCR/audio benchmarks open omni is at parity or ahead; on long-horizon visual-agent tasks the closed frontier still dominates. Hard-agent evals (TB4.0, HLE) also show a clear frontier gap even for strong open models (DeepSeek V4.1-Flash TB4.0 31.2 vs Opus 5 51.8 [VENDOR]).
- **Provenance rule**: tag every score (`vendor | independent | community | directional`); never compare across harnesses; treat vendor decimals as decoration.

### Inference engines — omni serving status (September 2026)

- **vLLM-Omni**: repo `vllm-project/vllm-omni`; community release Nov 2025; **v0.14.0 (Jan 2026)** first stable; **v0.16.0 (Feb 2026)** rebased on upstream vLLM v0.16.0; paper arXiv 2602.02204. Serves text, image, video, audio **and diffusion models** (DiT image/video) from one framework, one API, one deployment; OpenAI-compatible routes (`/v1/chat/completions`, `/v1/audio/speech`, `/v1/audio/generate`, `/v1/images/generations`, `/v1/videos`). Model coverage: Qwen3-Omni / Qwen3-TTS, Bagel, MiMo-Audio, GLM-Image, diffusion stack; platforms CUDA/ROCm/NPU/XPU; pipelined AR + diffusion execution. Multimodal flags: `--mm-encoder-tp-mode data` (vision tower data-parallel, avoids all-reduce), `--enable-vit-cuda-graph`. AWS Deep Learning Containers ship `vllm:omni-cuda` images. (Cross-ref §6 — one line: vLLM-Omni is the dedicated omni serving framework; production-grade as of January 2026.)
- **SGLang — broadest day-0 multimodal coverage**: `--enable-multimodal`; `--mm-process-config` for per-modality input limits; `--keep-mm-feature-on-device` (latency vs VRAM tradeoff; default moves feature tensors to CPU). **sglang-jax** has a standalone multimodal subsystem: text-to-image/video (Wan), vision-language (Qwen2.5-VL), **omni (Qwen3-Omni)**, audio (MiMo Audio); in-model path for regular AR multimodal models. Day-0 wins: Qwen3.6, Kimi K2.6, GLM-5.1, MiniMax M2.5/M2.6; native Kimi K2.5 (MoonViT + PatchMerger); Inkling-Small with MTP speculative decoding. Caveat: new hybrid checkpoints (e.g. DeepSeek-V4-Flash + MoonViT combos) sometimes need external model packages + narrow patches. **RadixAttention note**: multimodal mode auto-disables the radix cache — visual token sequences almost never share prefixes. Multi-node omni: Qwen3.8-Max (2.4T/95B, Aug 2026) runs on SGLang across 8×A800 with DeepEP MoE expert parallelism; MiMo-V2.6-Flash on SGLang TP16/DP2.
- **llama.cpp / GGUF — mature vision, experimental audio, nascent video**: `tools/mtmd` subsystem; vision encoder runs as a separate model, embeddings injected into the LLM token stream; **mmproj sidecars**; supported: LLaVA, MiniCPM-V, Pixtral, **Gemma 3/4**, Qwen2.5-VL, dots3-note, Qwen3.8-Flash-Next (qwen4exp) image handling. Audio: Whisper / MERaLiON-2 encoders; Gemma 4 E2B's GGUF projector reports **vision+audio** via mtmd; OpenAI-style `input_audio` message parts for Ultravox/Qwen2.5-Omni-class models. Video: `--video-*` CLI flags added (#24318); end-to-end video **not qualified** — requires `MTMD_VIDEO` builds + FFmpeg; workaround: extract frames, send as images. Quantization: MoE experts quantize like all weights; **no expert-parallelism fabric** (wide-EP is the llm-d/NVIDIA Dynamo layer); `vllm-gguf-plugin` brings experimental GGUF to vLLM. Unsloth Dynamic 2.0 GGUFs (Qwen3.x, GLM-5.3, Kimi, DeepSeek-V4, Gemma 4); UD-IQ2_XXS squeezes models like Qwen3.8-27B into 8 GB.
- **MLX / Apple Silicon and local tooling**: mlx-vlm (Blaizzy) vision fine-tuning + inference on M-series (Gemma 4 Vision, Qwen2.5-VL); mlx-community quants (Qwen3-VL-4B 3-bit ~3 GB, 8B 4-bit ~6 GB, 30B-A3B 6-bit ~20 GB); Whisper/Parakeet STT and Kokoro/Chatterbox TTS via mlx-audio; H3 MLX 8-bit port within days of weights. Unsloth claims image/video diffusion and multimodal training generally; Apple-Silicon video-understanding support specifically is community-led (mlx-vlm), not a confirmed Unsloth product capability. [UNVERIFIED — as narrowed in source.]
- **Ollama**: `gemma4:12b-it-qat` (7.2 GB), `gemma4:e2b-it-qat` (4.3 GB), `qwen3.5:9b` (6.6 GB, native vision) — one-command local multimodal.

### Practical deployment — choosing what to serve

- **Rule 1**: check the engine coverage matrix before choosing a model, not after — vLLM-Omni for production omni + diffusion, SGLang for fastest day-0 support, llama.cpp/mtmd for local/edge (video still the gap).
- **Rule 2**: "open weights" ≠ "runs locally" — MiMo-Pro (MIT) and MiniMax M3 are data-center class; Gemma 4 E2B/E4B are phone-class. Qwen3.8-Omni-Flash has no weights at all.
- **Rule 3**: budget the modality encoders as a separate line item (ViT/audio towers often kept at higher precision; SGLang offloads their features to CPU by default).
- **Rule 4**: KV cache at 1M context is the binding constraint — prefer GQA/MLA/shared-KV/HybridKV/TurboQuant architectures for long omni prompts; DeepSeek V4.1-Flash's 890 bytes/token MXFP4 KV is the current state of the art for cheap context carrying.
- **Rule 5**: for token pricing, audio ~6.25 tok/sec and video fps×resolution dominate long-context omni costs; pair `$ per token` with `tokens per task` (reasoning-agent fan-out can multiply the billed surface even as unit prices fall).
- **Rule 6**: verify the license file in the HF repo before commercial use — M3's license is disputed, H3's excludes major territories, K3 has a MaaS resale gate, and H3 forbids training other models on its outputs.

