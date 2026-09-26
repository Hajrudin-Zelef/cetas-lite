---
id: ai-industry-kb-2026/11-multimodality-open-weight-omni-models/benchmarks-open-omni-vs-closed-frontier-2026-the-pattern
title: "Benchmarks: open omni vs closed frontier (2026) — the pattern"
domain: multimodality-open-weight-omni-models
role: deep-dive
task: multimodal
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "MiniMax", "Moonshot", "OpenAI", "SGLang", "Z.ai"]
dates: ["2026-09"]
keywords: ["benchmarks", "omni", "agent", "agentic", "agents", "claude", "deepseek", "diffusion", "full-duplex", "gemini", "glm", "inference"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6026, 6054]
section: "11. Multimodality — Open-Weight Omni Models"
sha256: 8ee0541e461438806173c5b48daba567707813ea56dfbb329b39271126413f81
---

# Benchmarks: open omni vs closed frontier (2026) — the pattern

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

