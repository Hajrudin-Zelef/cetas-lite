---
id: ai-industry-kb-2026/11-multimodality-open-weight-omni-models/overview
title: "11. Multimodality — Open-Weight Omni Models"
domain: multimodality-open-weight-omni-models
role: deep-dive
task: multimodal
actors: ["Alibaba", "Apple", "ByteDance", "EU", "Google", "Hugging Face", "MiniMax", "Mistral", "Moonshot", "OpenAI", "SGLang", "United States", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2025-09", "2026-01", "2026-03-30", "2026-04-02", "2026-06-01", "2026-06-03", "2026-08", "2026-09", "2026-09-17", "2026-09-18", "2026-09-21", "2026-09-22"]
keywords: ["multimodal", "omni", "open-weight", "agentic", "apache", "benchmark", "benchmarks", "diffusion", "gemini", "glm", "gpu", "kimi"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5756, 5801]
section: "11. Multimodality — Open-Weight Omni Models"
sha256: e5439abfa8b8c8b6fe689804e14dfebc153e5cf4adc816f05f6457e632fe563f
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

