---
id: ai-industry-kb-2026/12-multimodality-video-generation-frontier-media-models/12-8-veo-3-1-runway-gen-4-5-and-deliberate-omissions
title: "12.8 Veo 3.1, Runway Gen-4.5, and deliberate omissions"
domain: multimodality-video-generation-frontier-media-models
role: deep-dive
task: multimodal
actors: ["Alibaba", "ByteDance", "China", "EU", "Google", "Hugging Face", "Meta", "MiniMax", "Nvidia", "SGLang", "United States", "xAI"]
dates: ["2026-04", "2026-05-12", "2026-07-31", "2026-09-17", "2026-09-22"]
keywords: ["agent", "agentic", "agents", "apache", "benchmark", "benchmarks", "compute", "cost", "diffusion", "distribution", "gemini", "gpus"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6251, 6280]
section: "12. Multimodality — Video Generation & Frontier Media Models"
sha256: 7b9de9431a3aeaae23d2e729fc68bd4ff26069de2dac20132665e23d7de9dd5b
---

# 12.8 Veo 3.1, Runway Gen-4.5, and deliberate omissions

- **HiDream-O1-Video-1.0 ("HiDream V1") — announced 2026-09-17, company-reported** [VENDOR]. HiDream.ai (Beijing) announced a "native omnimodal video generation model": text/image/video in; 1080p clips of 5–20 seconds with **natively synchronized audio** (dialogue with lip-sync, timed SFX, ambience). Duration is part of narrative planning rather than a fixed clip length — the model selects 5–20s based on event unfolding and pacing.
- Architecture (company claims): a unified **UiT (Unified Transformer)** foundation jointly modeling text, video, and audio in one generation process; post-training via **Diffusion Reinforcement Learning** with a multimodal reward model (visual quality, semantics, motion, physical plausibility, sound). Framed as one of four model families in a "native omnimodal world model portfolio" — the other three families were not detailed in the announcement material found.
- Benchmarks (company-reported, PR-sourced): #4 globally on the Artificial Analysis Image-to-Video Leaderboard (With Audio); #8 on Arena.ai Image-to-Video. **Independent reproduction: none found as of 2026-09-22 — treat all figures as company-reported.**
- **MiniMax H3 (API 2026-07-31 via Hailuo; open weights on Hugging Face Aug 2–3, 2026 — sources differ by one day):** 33B-parameter dense single-stream **H3-Omni-Transformer**; jointly understands text, images, video, audio and generates video with **native 32 kHz stereo audio in one pass** (4–15s clips, 24fps, 768p local canvas / 2K via hosted API). Generation modes: FL2VA (text-to-video with optional first/last-frame conditioning) and Ref2VA (reference-to-video: up to 9 images, 3 videos, 3 audio clips per prompt); instruction-based editing in plain language.
- Results: **#1 Video Editing (With Audio), #2 Text-to-Video, #3 Image-to-Video on Artificial Analysis**; ~1476 Arena.ai image-to-video Elo; described by the-decoder as the **first open model to top an AI video ranking**.
- Deployability: official BF16 SGLang recipe targets 4 GPUs; community NVFP4 quants run on a single RTX 5090 (~175s per 10s clip); fal's post-trained **H3 Max** variant renders 5s 768p in under 3s; its 75% launch discount expired Sept 15, 2026, quadrupling the 5s-clip price $0.10→$0.40.
- Architecture note (community skill card): H3's text understanding uses the full pretrained weights of **Qwen3-VL-32B** as the text encoder (hidden states from layer 50) — a late-fusion-style pretrained encoder inside a nominally unified model, the same modularity-survives-where-it-pays pattern documented for BAGEL and NVIDIA VoiceChat-11B. Unified at the product surface, modular in the trunk — the honest engineering description of 2026 "any-to-any" models.
- License (critical, detailed in §11): "MiniMax H3 Community License Agreement" — open weights, not Apache; **local deployment rights exclude the US, EU, UK, and South Korea**; companies above $20M/year revenue need separate written authorization; training other models on H3 outputs is prohibited. Wave2's thesis — "closed labs lead productized any-to-any output" — survives narrowed: it is now a *productization-and-licensing* lead, not a capability lead.

### 12.8 Veo 3.1, Runway Gen-4.5, and deliberate omissions

- **Veo 3.1:** exists and is commercially distributed. Adobe Firefly offered 30+ models by April 2026 including Veo 3.1, Runway Gen-4.5, and Kling 3.0 (per the baditaflorin deep-research compilation). No dedicated primary Google announcement was surfaced in this pass — no 2026 launch date is claimed.
- **Runway Gen-4.5:** commercially distributed (same Firefly evidence). The heyuan110 Seedance analysis asserts Seedance 2.0 beat it on the Artificial Analysis leaderboard at launch — that claim inherits the leaderboard's [UNVERIFIED] status. No further independent detail was found.
- **What the Firefly evidence actually proves:** Veo 3.1 and Gen-4.5 are commercially real (distributed through Adobe's 30-model catalog by April 2026) but primary-announcement-poor in the 2026 record — a sign that by 2026, Google and Runway were shipping video models through distribution partnerships rather than flagship launch events. Contrast with Kuaishou (dedicated PR waves) and ByteDance (Volcano Engine conferences).
- **Deliberate restraint:** this pass found no primary material on Luma, Pika, or Meta's Movie Gen sufficient to add — they stay out rather than get padded.
- xAI's Grok Imagine is covered in §1 (agents/Grok coverage is §13's); cross-referencing rather than duplicating.

### 12.9 Agentic video understanding and the end of the assembly bricolage

- **Google Agentic Video Understanding (announced 2026-09, ~20 days before 2026-09-22):** concentrates compute on query-relevant segments — adjustable frame rate, sub-second state-change detection, anomaly-triggered densification. Token consumption cut up to **88%**, accuracy up ~**7%**, cost reduced up to **66%**; Gemini 3.7 Flash with the feature sits on the accuracy-vs-cost Pareto frontier on **1H-VideoQA**. Static uniform processing retained only for videos under ~5 minutes.
- **StreamArena (2026-08):** first benchmark for continuous, interactive, long-horizon *streaming* video agents — 243 full-length videos averaging 88.8 minutes, 3,646 manually validated open-ended tasks (774 proactive monitoring tasks with ground-truth trigger times).
- **CFD — Caption-once, Frames-on-Demand** (arXiv 2609.11899): LVBench 52.9 overall, beating agent-based MemVid (44.4), VCA (41.3), VideoTree (28.8), VideoAgent (29.3); competitive with dense end-to-end AdaReTaKe-72B (53.3).
- **LVBench** (anchor benchmark): 103 manually curated YouTube videos, 117 hours total (avg 4,101s), 1,549 human-annotated QA pairs across 6 capability axes.
- **CASTLE / CVPR 2026 EgoVis:** 3rd-place solution used hierarchical knowledge-graph retrieval over 600 hours of synchronized ego+exo video (15 cameras, 4 days, 185 long-context QA samples). EgoLife pushed "very long" to 50 hours of egocentric week-long video.
- The benchmark frontier moved from minutes (LVBench's 4,101s average) to days (EgoLife's 50 hours) within 2026 — that is the scale agentic video understanding (AVU-style compute allocation: −88% tokens on query-relevant segments) is designed for; dense uniform processing is retained only for sub-5-minute content.
- **Evidence that natively-trained models displace stitched pipelines:** (1) DALL-E 2/3 sunset (2026-05-12), superseded by the natively multimodal GPT Image family; (2) Gemini Omni Flash generating dialogue with lip-sync, timed SFX, and ambience in one forward pass — positioned against a "separate Foley pipeline" world; (3) conversational any-to-any editing through chat with shared cross-modal context instead of chaining detection → edit → re-render tools.
- **Counterpoint — assembly persists where cheaper** [PARTIALLY VERIFIED]: xAI's stack is separate backends (Grok 4.6 + Imagine/Aurora + Voice API); the RAG/agent stack (ColPali, CLAP, SigLIP heads with score fusion) is still fundamentally modular. Honest formulation: native training replaced assembly for the *core* perception-generation loop; pipelines survive at the product-integration layer.

### 12.19 What shipped vs what was roadmap in any-to-any video (Sept 2026)

