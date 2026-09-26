---
id: ai-industry-kb-2026/12-multimodality-video-generation-frontier-media-models/12-4-kuaishou-kling-3-0-global-launch-2026-02-05-september-p
title: "12.4 Kuaishou Kling 3.0: global launch 2026-02-05; September PR wave ≠ a new model"
domain: multimodality-video-generation-frontier-media-models
role: deep-dive
task: multimodal
actors: ["ByteDance", "China", "Google", "OpenAI"]
dates: ["2025-12", "2026-02-05", "2026-02-12", "2026-04", "2026-05-19", "2026-06-23", "2026-07-01", "2026-07-31"]
keywords: ["arr", "consumer", "copyright", "diffusion", "distribution", "gemini", "inference", "latency", "leaderboard", "license", "multimodal", "omni"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6205, 6228]
section: "12. Multimodality — Video Generation & Frontier Media Models"
sha256: 2fcd5afc7606e268e31e3692b57fb96ed5e8d900486b0ec55e574ccc97cb6a73
---

# 12.4 Kuaishou Kling 3.0: global launch 2026-02-05; September PR wave ≠ a new model

- **Seedance 2.0 — launched 2026-02-12 (China).** ByteDance Seed's announcement via Volcano Engine (community-tracked timeline): a **quad-modal video generator** (text/image/video/audio references), with native synchronized audio, up to **2K output**, and multi-shot control.
- API distribution began via third-party platforms (e.g., a Hypereal/Hyperreal commercial PR); one secondary PR prices API access at roughly $0.14 per 15-second clip — commercial-secondary, treat pricing as unconfirmed.
- A community analysis claims Seedance 2.0 hit **#1 on the Artificial Analysis text-to-video leaderboard** at launch, beating Veo 3, Sora 2, and Runway Gen-4.5 — **[UNVERIFIED]:** the Artificial Analysis page itself was not re-opened in this pass; the claim reaches us via a secondary blog quoting the leaderboard.
- The launch was initially China-focused; one deep-research compilation reports the **planned international launch was put on hold after copyright disputes and demands for stronger safeguards** — echoed independently by CIOL ("Seedance platform has paused expansion due to copyright concerns").
- Quad-modal referencing (text/image/video/audio) is Seedance's distinctive input-breadth claim: up to 50 multimodal references in 2.5, with region-level editing and Maya/Blender plugins positioning it for production VFX pipelines rather than consumer prompting — the B2B-production framing that distinguishes ByteDance's video play from Sora's consumer play.
- The Seedance arc (2.0 China-only → 2.5 Dreamina global) vs Kling 3.0 (global from day one) vs H3 (open weights, license-restricted) makes this section's availability-geography point concrete: identical "launch" language covers three different distribution realities.
- **Seedance 2.5 — announced 2026-06-23, global Dreamina launch 2026-07-31 [VENDOR/secondary].** A community-maintained README (ikaijua/awesome-aitools) gives the tightest available timeline: announced at the Volcano Engine **FORCE conference (2026-06-23)**; launched globally on ByteDance's **Dreamina** platform on **2026-07-31**.
- Claimed specs: **30-second single-pass** video generation, **native 4K**, up to **50 multimodal references**, region-level editing, plugins for Maya and Blender. No ByteDance-primary English source was found; the specs are vendor claims relayed by community docs — keep the [UNVERIFIED]/[VENDOR] marks.
- **This corrects wave2.1 §1.3's dating** ("released the same week H3's weights dropped," ~early Aug 2026): 2.5 was already out when H3's weights dropped, not released that week. The underlying claim (30s clips with built-in audio) is consistent; the timing is corrected.

### 12.4 Kuaishou Kling 3.0: global launch 2026-02-05; September PR wave ≠ a new model

- **Global launch 2026-02-05** (per a Novita AI integration post, April 2026): up to **15-second clips** (vs 10s in prior versions); **native multilingual audio** (Chinese, English, Japanese, Korean, Spanish, with accents and dialects); stronger element consistency (characters, objects, scenes coherent across frames); better text/logo preservation; more photorealistic faces and natural body motion.
- **API pricing via Novita:** Standard **$0.168/s without audio, $0.252/s with audio**; Pro **$0.224/s without audio, $0.336/s with audio**.
- **September 11–12, 2026 "AI Director" PR wave** [COMPANY PR, syndicated via ABNewswire]. Kuaishou positioned Kling 3.0 as a cinematic **"AI Director paradigm"**: native **4K at 60fps**, multi-shot storyboarding with **up to 6 connected shots**.
- **Motion Control AI:** 6-axis camera control (pan, tilt, roll, zoom, horizontal, vertical with −10 to 10 intensity values); Motion Brush region painting; character motion transfer from reference video; a Motion Library of pre-built templates; full API support for all motion parameters.
- Native **lip sync in 5 languages** (EN, ZH, JA, KO, ES); unified multimodal **MVL architecture** (vs the older Diffusion Transformer 3D-VAE); 16-bit HDR; up to 4 character reference images.
- A second variant, **Kling 3.0 Omni (O3)**: multimodal-transformer variant with an integrated 48kHz audio synthesis engine, **35% lower inference latency** than the V3 base, physics-based object interaction, real-time latent-space editing [COMPANY PR only — no independent verification].
- The September PR's Motion Control AI stack (6-axis camera, Motion Brush region painting, motion transfer from reference video, Motion Library templates, full API parameter support) is the most detailed public motion-control surface of any 2026 video model — but it arrives as PR copy, not as independently tested capability; treat the feature list as claimed, the positioning ("AI Director") as marketing.
- **Caution:** the September PRs carry **no new version number** — this reads as a **re-positioning/re-launch of the February model** (director-grade positioning), not a confirmed Kling 3.1. Do not cite Sept 2026 as a new model release.
- **Economics (secondary):** a deep-research compilation reports Kling reached **$240M ARR in December 2025** [UNVERIFIED — single secondary compilation]. If true, Kling is the clearest counterpoint to Sora: a video product with nine-figure run-rate while OpenAI's consumer product failed on unit economics.

### 12.5 Gemini Omni Flash at Google I/O 2026-05-19; Nano Banana 2 Lite ~2026-07-01/02

