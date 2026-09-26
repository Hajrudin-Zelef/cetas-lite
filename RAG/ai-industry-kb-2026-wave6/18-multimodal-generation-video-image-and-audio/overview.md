---
id: ai-industry-kb-2026-wave6/18-multimodal-generation-video-image-and-audio/overview
title: "§18. Multimodal Generation: Video, Image, and Audio"
domain: multimodal-generation-video-image-and-audio
role: deep-dive
task: multimodal
actors: ["Alibaba", "Apple", "ByteDance", "EU", "Google", "Groq", "Meta", "MiniMax", "OpenAI"]
dates: ["2025-03-26", "2025-05-20", "2025-07", "2025-07-28", "2025-08-27", "2025-09", "2025-09-19", "2025-09-30", "2025-11-25", "2025-12", "2025-12-03", "2025-12-31", "2026-01-15", "2026-01-26", "2026-02-17", "2026-03-17", "2026-03-30", "2026-04", "2026-04-14", "2026-04-21", "2026-04-26", "2026-04-30", "2026-05", "2026-06", "2026-06-03", "2026-06-10", "2026-07", "2026-07-03", "2026-07-21", "2026-07-24", "2026-07-30", "2026-08-31", "2026-09", "2026-09-22", "2026-09-24"]
keywords: ["multimodal", "apache", "chatgpt", "consumer", "gemini", "gpu", "license", "moe", "omni", "open weights", "open-weight", "pricing"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8749, 8787]
section: "§18. Multimodal Generation: Video, Image, and Audio"
delta_of: ai-industry-kb-2026
sha256: 03a15454577c25586921ca66ec9d1b0424969f5d43e9ef412a104d3318bc5852
---

# §18. Multimodal Generation: Video, Image, and Audio

Keywords: multimodal generation, video generation, image generation, text-to-speech, Sora 2, Runway Gen-4.5, ElevenLabs, Suno v5.5, FLUX.2, Ideogram 4, Midjourney V8, Luma Ray 3, Cartesia Sonic-3.6, Dia, TTS, voice cloning, music generation, image editing, video editing, audio tags

## Summary
- **Sora 2 launched 2025-09-30** — the brief's 2026 date is corrected to 2025 [SECONDARY]; the consumer app/web reportedly ended 2026-04-26 and the API end was reported for 2026-09-24, both single-source-family and needing corroboration [SECONDARY].
- **Runway has no Gen-5 as of 2026-09-22**: Gen-4.5 (December 2025) remains the flagship, with Aleph as the in-context video-editing line [SECONDARY].
- **ElevenLabs v4 was preview-only** (ElevenSummit Warsaw, June 2026); Eleven v3 remains the flagship speech model — no GA, model ID, or endpoint as of September 2026 [SECONDARY].
- **PlayAI/PlayHT is dead**: Meta's July 2025 acqui-hire preceded the December 31, 2025 shutdown; Groq retired and repointed the engine [SECONDARY].
- **Ideogram 4** (June 3, 2026) is a 9.3B typography-layout image model whose code/pipeline is Apache 2.0 but whose weights sit under a non-commercial Ideogram agreement — the "Apache 2.0" shorthand for the weights is contradicted [SECONDARY].
- **Midjourney's V8.1 alpha date is unresolved** (2026-04-14 vs 2026-04-30); both dates are retained pending resolution [SECONDARY]; V8.2 became the default 2026-07-24 [SECONDARY].
- **Hunyuan3D 3.0 means free hosted access, not open weights** — the open checkpoints concentrate in the 2.x/Omni lines; the Hy3 Preview license split is Tencent Hy Community License for the preview vs Apache 2.0 for full Hy3/Hy4 Preview [SECONDARY].
- Image generation consolidates toward editing-capable models (FLUX.2 [dev] reference editing, GPT Image 2 editing endpoints, Photon multi-image referencing), while audio consolidates on licensed-data business models (Suno, Udio, ElevenLabs Music v2) [DIRECTIONAL].

## Key dated facts
### Image generation
- **2025-03-26** — Ideogram 3.0 launched with Ideogram Canvas (closed, paid tiers) [SECONDARY].
- **2025-05-20** — Google Imagen 4 announced at I/O: standard $0.04/image, Ultra $0.06, Fast $0.02; 2K output; SynthID watermarking [SECONDARY]; an Imagen 4 retirement date remains [UNVERIFIED] — no precise deprecation found.
- **2025-11-25** — FLUX.2 [dev] released: 32B text-to-image with single- and multi-reference editing; BFL license terms [VENDOR]; the **FLUX.2 [klein]** family (**2026-01-15**): 4B and 9B four-step distilled models; klein 4B Apache 2.0, 9B different terms, dev downloadable open-weight but commercial use requires a BFL license [VENDOR].
- **2026-02-17** — Recraft V4 reportedly launched: four variants (raster 1024², Pro 2048², Vector, Vector Pro); editable SVG output as the differentiator; V4.1 reportedly May 2026 [SECONDARY]; Elo 1172 / 72% as a single data point [DIRECTIONAL].
- **2026-03-17** — Midjourney V8 alpha [SECONDARY]; **V8.1 alpha April 2026** — date discrepancy **2026-04-14 vs 2026-04-30**, both retained [SECONDARY]; V8.1 default 2026-06-10/11; V8.2 default 2026-07-24; GPU-native rewrite, ~5× faster, native 2K via `--hd`, improved text; no public first-party REST API [SECONDARY].
- **2026-04-21** — OpenAI GPT Image 2 released (snapshot gpt-image-2-2026-04-21): generation + editing endpoints; $5/$8/$30 per 1M tokens (text-in/image-in/image-out); 1024² images at $0.006/$0.053/$0.211 [SECONDARY]; the GPT Image 2.5 "Sunburst" claim is [UNVERIFIED] single weak source.
- **2026-06-03** — Ideogram 4 released: 9.3B single-stream DiT, Qwen3-VL-8B text encoder, native 2K, structured JSON prompts, bounding boxes, color control; ~0.97 English OCR, 47.9% typography preference [DIRECTIONAL]; code/pipeline Apache 2.0 but weights under a non-commercial Ideogram agreement [SECONDARY].
- Luma Photon / Photon Flash: 1080p images $0.015 ($0.002 Flash), character consistency, multi-image referencing; no 2026 version bump [SECONDARY]; text rendering weak per review coverage [DIRECTIONAL].

### Video generation
- **2025-07-28** — Wan 2.2 released (Apache 2.0 open-weight): T2V-A14B/I2V-A14B MoE, TI2V-5B; up to 5s 720p [SECONDARY]; Wan 2.2 Animate (open weights/code) 2025-09-19 [SECONDARY]; Wan 2.5/2.6 remain closed/API-first [UNVERIFIED].
- **2025-08-27** — PixVerse V5: 360p–1080p, 5/8s, start/end frames, effects, Turbo/HD tiers [SECONDARY]; V6 (2026-03-30) single weak source [UNVERIFIED]; V5 modality coverage internally contradicted (image-only vs T2V/I2V) [SECONDARY].
- **2025-09-30** — **Sora 2 launched** (not 2026 — corrected from the brief): synchronized audio/dialogue, improved physics, Cameos; iOS/social app + Pro for ChatGPT Pro [SECONDARY].
- **2025-12-03** — HunyuanVideo released: 13B+ open code/weights under the Tencent Hy Community License (bounded below 100M monthly users, outside EU/UK/SK); no 2025–2026 descendant found [SECONDARY].
- **2026-01-26** — Luma Ray3.14: native 1080p, 4× faster/3× cheaper at 720p; dropped Character Reference and HDR/EXR; "reasoning" video model line [SECONDARY]; by September 2026 Ray 3.2 is current (5/10s T2V/I2V, 1080p, Modify up to 20s, 16 keyframes, 16-bit EXR/ACES, no generated audio) [SECONDARY].
- **2026-04-26** — Sora consumer app/web reportedly ended [SECONDARY] — single-source-family, needs corroboration; API end reported for 2026-09-24 [SECONDARY] — single-source-family.
- **2026-07-03** — Vidu S1 unveiled: real-time voice-driven interactive video from a single image, 540p/25fps, peak 42fps [SECONDARY] — syndicated coverage flagged, needs an independent source.
- Pika 2.5 flagship: 480p/720p/1080p, 5 or 10s clips; Pikaframes 2–5 keyframes, 20–25s; 2026-08-31 pricing: free 80 credits/mo, paid from $10/mo, a 5s clip at 12/20/40 credits by tier [SECONDARY].
- Kling 3.0/O3: 5–15s clips, multi-shot, native 4K reported [SECONDARY]; Kling 2.6→3.0/O3 line; **the Kling pricing block is contradicted across the corpus — do not cite any Kling price** [DIRECTIONAL].
- Runway Gen-4/Gen-4 Turbo + Aleph in-context video editing; API credits at 5/12/15 per second by model [SECONDARY]; **no Gen-5 as of 2026-09-22** — Gen-4.5 (December 2025) is the flagship [SECONDARY]; Aleph v1 sunset 2026-07-30, Aleph 2.0 replacement [UNVERIFIED]; Runway pivoting toward world models, $5.3B valuation per May 2026 coverage [SECONDARY].
- Hailuo 02: September 2025 integration coverage already named MiniMax-Hailuo-02; by July 2026 Hailuo 2.3 is the flagship (NCR architecture, native 1080p, T2V/I2V, 10s, no native audio) [SECONDARY].
- Pointer: MiniMax H3/Hailuo 3 video detail lives in §7; ByteDance Seedance 2.5 in §17; Google Gemini Omni/Veo in §14; Qwen Image 3.0 in wave6/04 (July 21, 2026, closed/API-only) [DIRECTIONAL].

