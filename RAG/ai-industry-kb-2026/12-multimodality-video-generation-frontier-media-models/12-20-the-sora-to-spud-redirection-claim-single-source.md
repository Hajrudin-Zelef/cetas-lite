---
id: ai-industry-kb-2026/12-multimodality-video-generation-frontier-media-models/12-20-the-sora-to-spud-redirection-claim-single-source
title: "12.20 The Sora-to-Spud redirection claim (single-source)"
domain: multimodality-video-generation-frontier-media-models
role: deep-dive
task: multimodal
actors: ["Anthropic", "EU", "Google", "MiniMax", "OpenAI"]
dates: ["2025-12", "2026-03-24", "2026-04", "2026-09", "2026-09-17", "2026-09-22"]
keywords: ["arr", "claude", "compute", "consumer", "copyright", "fable 5", "gemini", "inference", "ipo", "leaderboard", "omni", "pretraining"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6281, 6327]
section: "12. Multimodality — Video Generation & Frontier Media Models"
sha256: 38d9f996d727abd9ab6205acc8b62db0617fcf166d2d2c186fba78a772f0e65c
---

# 12.20 The Sora-to-Spud redirection claim (single-source)

- **Shipped with synchronized audio in a single pass:** Gemini Omni Flash (video + synced audio out of one forward pass — dialogue with lip-sync, timed SFX, ambience — per wave2's architecture description); MiniMax H3 (native 32 kHz stereo audio in one pass, 4–15s, 24fps); Sora 2 at launch (natively synchronized audio — product now dead); Seedance 2.0/2.5 claims native synchronized audio [VENDOR/secondary]; Kling 3.0 native multilingual audio and 5-language lip sync (product-distributed, PR-corroborated).
- **Claimed, not independently reproduced:** HiDream-O1-Video-1.0's synchronized audio (lip-sync dialogue, timed SFX, ambience) [VENDOR, company-reported 2026-09-17, no independent reproduction as of 2026-09-22].
- **Explicitly roadmap, not shipped at preview:** Gemini Omni Flash's audio *references* as inputs and scene extension — Google's own API schema docs mark them unsupported at preview. The any-to-any headline is therefore accurate for outputs (video+audio out) and aspirational for inputs (audio-in) at the September 2026 state.
- **Worlds, not video:** Project Genie produces playable runtime worlds (real-time interactive, shareable links, remix function) — not exportable video files or 3D assets. It must not be counted as a video generator in any leaderboard or capability comparison.
- Net: by Sept 2026 synchronized audio was a shipped commodity in video generation (multiple vendors), any-to-any *input* breadth was still uneven, and the remaining frontier differentiators were clip length (Seedance 2.5's claimed 30s single-pass), resolution (native 4K), and conversational editing depth — not audio presence.

### 12.20 The Sora-to-Spud redirection claim (single-source)

- One secondary source (humai.blog) claims OpenAI discontinued Sora video generation and redirected resources to "Spud" (the GPT-5.5 pretraining run completed 2026-03-24) — **[SINGLE-SOURCE, UNVERIFIED]**.
- It is consistent in timing with the 2026-03-24 shutdown announcement and with OpenAI's compute-reallocation framing ("side quests"), but no second outlet or OpenAI-primary source confirms the specific resource-redirection claim. Keep flagged; do not state as fact.
- The confirmed redirect target is generic: "higher-priority products" (per the Simo/Altman reporting chain), not a named model.

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

