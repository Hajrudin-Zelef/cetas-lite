---
id: ai-industry-kb-2026/12-multimodality-video-generation-frontier-media-models/figures-and-metrics
title: "Figures and metrics"
domain: multimodality-video-generation-frontier-media-models
role: deep-dive
task: multimodal
actors: ["Anthropic", "ByteDance", "China", "EU", "Google", "MiniMax", "OpenAI"]
dates: ["2025-12", "2026-02", "2026-02-05", "2026-02-12", "2026-03-24", "2026-04", "2026-04-26", "2026-09"]
keywords: ["agentic", "arr", "claude", "compute", "consumer", "copyright", "cost", "fable 5", "gemini", "inference", "ipo", "leaderboard"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6293, 6376]
section: "12. Multimodality — Video Generation & Frontier Media Models"
sha256: 7a65115a1dc10dd7c83793ff74a2f26140f2925aca9c55d788e503d18585bbcf
---

# Figures and metrics

## Figures and metrics

### 12.10 Economic divergence: Sora vs Kling vs Firefly (2026)

- **Sora (dead consumer product):** reported ~$5.4B annual compute spend vs ~$2.1M lifetime revenue (CIOL, UNVERIFIED methodology — but no outlet disputes the direction). 1M+ downloads in the first five days (ngram) did not change the equation; ngram frames it as "the most useful data point the AI video market has produced in two years."
- **Kling (alive, growing):** reported $240M ARR in December 2025 (secondary compilation, UNVERIFIED) — the clearest counterpoint to Sora: a video product with nine-figure run-rate while OpenAI's consumer product failed.
- **Adobe Firefly (alive, aggregating):** 30+ video models by April 2026, including Veo 3.1, Runway Gen-4.5, and Kling 3.0 — Adobe rents inference, doesn't train foundation models. The aggregation model sidesteps the training-compute bill entirely.
- **The lesson for the knowledge base:** the only consumer video product that died in 2026 was the one with the worst reported economics. Survivors monetize B2B/API (Kling, Seedance), aggregate others' models (Firefly), or gate behind expensive subscriptions (Genie at $249.99/mo). "Technology leadership ≠ product viability" is now an evidenced claim, not a slogan.
- **IPO context (secondary interpretation, not OpenAI's stated reason):** CIOL reports OpenAI was preparing an IPO at a ~$300B reported valuation and frames the shutdown as removing a loss-making product ahead of it. Treat as analyst framing; the official reason given was compute reallocation ("side quests").

### 12.11 Reading the September 2026 video leaderboards

- Leaderboard positions in this domain are almost entirely secondary- or vendor-sourced: Seedance 2.0's claimed #1 on Artificial Analysis text-to-video at launch reaches us via a community blog quoting the leaderboard (AA page not re-opened — [UNVERIFIED]); HiDream's #4/#8 are company PR with no independent reproduction [VENDOR].
- The one leaderboard claim with independent secondary corroboration: MiniMax H3's #1 on Artificial Analysis Video Editing (With Audio), reported by the-decoder as the first open model to top an AI video ranking (secondary corroboration, not a re-open of the leaderboard itself).
- Wave2's price comparators remain the consistent anchor across sources: Gemini Omni Flash $0.10/s (720p) vs Veo 3.1 Standard $0.40/s (1080p) vs Sora 2 $0.10–$0.70/s — the per-second price gap between 720p and 1080p output tiers is itself a product variable (quality tiering as pricing lever).

### 12.12 Safety, watermarking, and content policy in video generation

- **Sora 2 launch policies:** only under-18-suitable content; no real people or public figures; human-likeness character uploads blocked by default; no copyrighted characters or music. The Cameos likeness feature coexisted with strict likeness moderation — and likeness disputes (MLK Jr. estate pushing restrictions, Washington Post; unauthorized Robin Williams videos) ran through the product's life.
- **Gemini Omni Flash:** C2PA credentials and SynthID watermarks on by default at developer rollout — provenance-by-default as a product decision at launch, not a retrofit.
- **Seedance 2.0's international launch** was reportedly put on hold over copyright disputes and demands for stronger safeguards (CIOL; deep-research compilation) — copyright and safeguards as a launch-gating variable in 2026, not just a legal footnote.
- **Claude Fable 5.1** introduced invisible watermarking of generated text with a detection API in private preview for EU-law eligible orgs — the watermarking move extended to text flagships in the same window, but is a §3 matter; recorded here only for the media-provenance pattern.

### 12.13 Per-second video API economics

- Gemini Omni Flash: **$0.10/sec** at 720p (5,792 output tokens/sec at $17.50/M video tokens); $1.50/M input tokens; 50% batch discount. Matches Veo 3.1 Fast ($0.10/s) on price point.
- Veo 3.1 Standard: **$0.40/sec** at 1080p; Sora 2: $0.10–$0.70/sec (wave2 comparators).
- Kling 3.0 (Novita): Standard $0.168/s without audio / $0.252/s with audio; Pro $0.224/s without audio / $0.336/s with audio.
- The audio multiplier is consistent across vendors: Kling's with-audio price is 1.5x the without-audio price at both tiers ($0.168→$0.252, $0.224→$0.336); Omni Flash's 50% batch discount halves the $0.10/s and $1.50/M-input figures for batch workloads — the effective price depends on billing mode as much as on the headline rate.
- Seedance 2.0 API: ~$0.14 per 15-second clip [UNVERIFIED — commercial-secondary PR].
- Price-per-second is not price-per-task: Seedance's per-clip pricing (~$0.14/15s) scales with clip length (a 30s 2.5 clip ≠ 2× a 15s clip necessarily), while Kling/Novita and Omni Flash bill per second with audio/resolution multipliers. The two billing units (per-clip vs per-second) coexist in 2026 — task-level economics require the clip-length multiplier, not a single $/s figure.
- Sora 2 API: new generations up to 20s; up to six extensions; 120s maximum total length; 720p for Sora 2, up to 1080p portrait/landscape for 2 Pro.

### 12.16 Revenue, spend, and scale claims

- Sora: reported ~$5.4B annual compute spend vs ~$2.1M lifetime revenue [UNVERIFIED — CIOL, no named methodology]; 1M+ downloads in first five days (ngram).
- Kling: reported **$240M ARR in December 2025** [UNVERIFIED — single secondary deep-research compilation].
- Google AI Ultra plan: **$249.99/mo** (Project Genie consumer gate).
- Adobe Firefly: **30+ video models** by April 2026 (Veo 3.1, Runway Gen-4.5, Kling 3.0 among them).
- gpt-image-1.5: $0.011–$0.167 per 1024x1024 by quality; Mini variant $0.005–$0.036; ~4s per 1024x1024 (~4x faster than prior); 5 input images preserved with high fidelity; 1B+ images generated via Gemini (Nano Banana Pro counterpoint).
- Nano Banana 2 Lite: $0.034 per 1K images; ~4s text-to-image at ~1K; $0.25/$1.50 per M tokens (gigazine).
- fal H3 Max: 5s 768p render in under 3s; launch discount expiry Sept 15, 2026 quadrupled 5s-clip price $0.10→$0.40; community NVFP4 quants ~175s per 10s clip on a single RTX 5090.

### 12.17 Video-understanding and leaderboard figures

- Google Agentic Video Understanding (Sept 2026): **−88% tokens, +~7% accuracy, −66% cost**; Gemini 3.7 Flash on the accuracy-vs-cost Pareto frontier on 1H-VideoQA.
- StreamArena (2026-08): 243 full-length videos (avg 88.8 min), 3,646 manually validated open-ended tasks, 774 proactive-monitoring tasks with ground-truth trigger times.
- LVBench: 103 videos, 117 hours total, avg 4,101s, 1,549 human QA pairs, 6 capability axes.
- CFD on LVBench: **52.9** overall vs MemVid 44.4 / VCA 41.3 / VideoTree 28.8 / VideoAgent 29.3 / AdaReTaKe-72B 53.3.
- Leaderboard claims (secondary/unverified unless noted): Seedance 2.0 claimed #1 Artificial Analysis text-to-video at launch [UNVERIFIED — secondary blog quoting the leaderboard]; MiniMax H3 #1 AA Video Editing (With Audio), #2 Text-to-Video, #3 Image-to-Video; H3 ~1476 Arena.ai image-to-video Elo; HiDream-O1-Video-1.0 company-reported #4 AA Image-to-Video (With Audio), #8 Arena.ai [VENDOR — no independent reproduction].
- Genie 3: 720p/24fps explorable worlds lasting several minutes; worlds shareable via public link; real-time interactive, not exportable files.

### 12.18 Clip-spec comparison (flagship video outputs, Sept 2026)

- **Sora 2 / 2 Pro** (discontinued as API): new generations up to 20s, up to six extensions, 120s max total; 720p (2 Pro up to 1080p portrait/landscape); synchronized audio at launch.
- **Gemini Omni Flash:** 10s clips at preview (deployment decision, not architectural cap); video + synchronized audio in a single forward pass; $0.10/s at 720p.
- **Seedance 2.5** [VENDOR/secondary]: 30s single-pass clips, native 4K, up to 50 multimodal references, region-level editing, Maya/Blender plugins.
- **Kling 3.0:** 15s clips; 4K at 60fps (director positioning); multi-shot storyboarding up to 6 connected shots; 5-language lip sync; 48kHz audio synthesis engine in the O3 variant [PR].
- **HiDream-O1-Video-1.0** [VENDOR]: 1080p, 5–20s clips (duration chosen by narrative planning), natively synchronized audio (lip-sync dialogue, timed SFX, ambience).
- **MiniMax H3:** 4–15s clips, 24fps, 768p local canvas / 2K via hosted API; native 32 kHz stereo audio in one pass; Ref2VA up to 9 images + 3 videos + 3 audio clips per prompt.
- Direction of travel across the board: longer single-pass clips (Seedance 2.5's 30s), higher native resolution (4K), and synchronized audio as a default rather than a premium — with per-second pricing tiered by resolution and audio inclusion.

### 12.25 Genie vs video generators: the category distinction

- **Project Genie is a simulation surface, not a video generator:** real-time interactive worlds from text prompts, the "remix" function turning one world into another, shareable via public link, Street View/Maps grounding for real-world-rooted simulations — but no exportable 3D assets and no video files.
- Google's stated 2026 motivation (planning and training, not just play) plus the [UNVERIFIED] Waymo fine-tuned-variant claim point to the same category: world models as simulation infrastructure, closer to a game engine than to a text-to-video product.
- The Sora team's reported continuation as a world-simulation research unit (robotics, physical-environment modelling) is the mirror image: OpenAI exited the video *product* while keeping the world-model *research*.
- For the knowledge base: never place Genie on a text-to-video leaderboard, never count "no Genie 4" as a video-model gap, and never describe the Sora team continuation as video-product continuity. Three distinct categories — video generation (Sora/Seedance/Kling/Veo/Omni Flash/H3/HiDream), interactive world models (Genie), world-simulation research (Sora team post-shutdown).

### 12.26 The February 2026 competitive moment (pre-shutdown context)

- **2026-02-05:** Kling 3.0 launches globally with 15s clips and native multilingual audio. **2026-02-12:** Seedance 2.0 launches in China with quad-modal references, native audio, and 2K output. Sora's web/app was still live (shutdown announced 2026-03-24, effective 2026-04-26).
- February 2026 is therefore the last month all three — OpenAI (Sora), Kuaishou (Kling), ByteDance (Seedance) — had live consumer video products simultaneously. By July, the board had no Sora, a Dreamina-launched Seedance 2.5, an open-weight H3 topping an editing ranking, and an Omni Flash API at $0.10/s.
- The February→September arc is the cleanest 2026 illustration of the section's thesis: the technology leadership stayed in place (Sora 2 was still best-in-class at its death per contemporary coverage), the product leadership moved to labs with working unit economics.

### 12.27 Audio and multilingual claims compared (Sept 2026)

- **Kling 3.0:** native multilingual audio at launch (ZH/EN/JA/KO/ES with accents and dialects); September PR wave adds native 5-language lip sync (EN, ZH, JA, KO, ES) and the O3 variant's 48kHz integrated audio synthesis engine [PR].
- **MiniMax H3:** native 32 kHz stereo audio in a single generation pass (4–15s clips); Ref2VA accepts up to 3 audio clips per prompt as references — audio as both input and output modality.
- **Gemini Omni Flash:** dialogue with lip-sync, timed SFX, and ambience generated in one forward pass (shipped); audio *references* as inputs unsupported at preview (roadmap).
- **HiDream-O1-Video-1.0** [VENDOR]: natively synchronized audio (lip-sync dialogue, timed SFX, ambience) claimed at 1080p, 5–20s.
- **Seedance 2.0/2.5** [VENDOR/secondary]: native synchronized audio (2.0); 30s single-pass clips with built-in audio (2.5).
- Pattern: by Sept 2026 "synchronized audio" was table stakes in video generation claims; the differentiators were audio *quality tiers* (48kHz synthesis engines, stereo 32kHz), audio *as input* (H3's reference audio; Omni Flash's unsupported-at-preview audio refs), and multilingual lip sync breadth (Kling's 5 languages).

