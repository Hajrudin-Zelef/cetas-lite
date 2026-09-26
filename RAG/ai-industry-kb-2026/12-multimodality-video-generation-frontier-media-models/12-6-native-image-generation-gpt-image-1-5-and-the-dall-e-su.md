---
id: ai-industry-kb-2026/12-multimodality-video-generation-frontier-media-models/12-6-native-image-generation-gpt-image-1-5-and-the-dall-e-su
title: "12.6 Native image generation: GPT Image 1.5 and the DALL-E sunset"
domain: multimodality-video-generation-frontier-media-models
role: deep-dive
task: multimodal
actors: ["Google", "MiniMax", "OpenAI"]
dates: ["2025-12-16", "2026-05-12", "2026-05-19", "2026-06-30", "2026-07", "2026-07-01"]
keywords: ["agent", "consumer", "diffusion", "gemini", "multimodal", "omni", "pricing", "revenue", "text-to-image", "video generation", "watermarking"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6229, 6250]
section: "12. Multimodality — Video Generation & Frontier Media Models"
sha256: 88b8c9071211c906dd62dd790ef417ab9a3b03409f23b5fc69d4d0f7e8574dd4
---

# 12.6 Native image generation: GPT Image 1.5 and the DALL-E sunset

- **Google I/O 2026 opened 2026-05-19** (livemint live blog: keynote May 19, 10:00 am PT). Gemini Omni was unveiled there — "Gemini Omni is coming to Gemini app from today for Google's paid subscribers," described as "Nano Banana for video."
- The model: **unified multimodal video generation** — text, image, and video inputs combined to generate and conversationally edit video (character swaps, relighting, style transfers via natural-language prompts, scene consistency, on-screen text/graphics sync). Priced at **$0.10/second of video output**, matching Veo 3.1 Fast; 10-second generations at preview.
- Architecture per wave2: text/image/audio/video in, video with **synchronized audio out** (dialogue with lip-sync, timed SFX, ambience) in a **single forward pass**; conversational multi-turn editing; reference stacking (7 images + 3 video clips); 1M context; SynthID watermarking; 10-second clips at launch (deployment decision, not architectural cap).
- **API GA 2026-06-30:** $0.10/sec of 720p output (5,792 output tokens/sec at $17.50/M video tokens), $1.50/M input tokens, 50% batch discount. Consumer access via Gemini app, Flow, YouTube Shorts/Create.
- **Nuance on shipped capability [wave3 verification]:** Google's own API schema accepted video references up to 3 seconds but did not correctly process them at preview; **audio references and scene extension were explicitly not supported**. Synchronized native audio generation was the stated direction, not the shipped capability — keep this distinction.
- **Developer rollout came later:** Gemini Omni Flash (`gemini-omni-flash-preview`) reached the Gemini API, Google AI Studio, and the Gemini Enterprise Agent Platform around **late June / early July 2026**, with C2PA credentials and SynthID watermarks on by default.
- The I/O unveiling framed Omni as "Nano Banana for video" — Google's own analogy positions the model as the video-side counterpart of its image family, i.e., the any-to-any flagship is marketed as a sibling product, not as a Veo successor. That framing matters for §3/§12 cross-references: Veo 3.1 remains the specialist video model; Omni Flash is the conversational any-to-any one.
- Pricing nuance: the $0.10/s headline matches Veo 3.1 Fast exactly — Google priced its any-to-any flagship at its specialist model's fast tier, a deliberate price-parity signal that any-to-any is not a premium upsell but the new default.
- **Nano Banana 2 Lite [MISDATING CORRECTED]:** the new entry tier of the Nano Banana image family (built on Gemini 3.1 Flash Lite), below Nano Banana 2 and Pro; ~4-second text-to-image at ~1K, $0.034 per 1K images ($0.25/$1.50 per M tokens per gigazine). Launched in the **developer rollout of ~2026-07-01/02**, alongside Gemini Omni Flash's wider release — **not at Google I/O** (evidence: theaiinsider.tech roundup dated 2026/07/02; aiweekly's "today" piece ≈ 83 days before Sept 22 ≈ July 1). The **original** Nano Banana (Gemini 2.5 Flash Image) predates it; Google explicitly encouraged developers on the original Nano Banana to upgrade.

### 12.6 Native image generation: GPT Image 1.5 and the DALL-E sunset

- **gpt-image-1.5 (GA 2025-12-16):** preserves 5 input images with high fidelity (vs 1), ~4x faster generation (~4s for 1024x1024), 20% lower API pricing ($0.011–$0.167 per 1024x1024 by quality; Mini variant $0.005–$0.036). The only widely-deployed production image API emitting real RGBA PNG/WebP with populated alpha channel.
- **DALL-E 2 and DALL-E 3 deprecated with sunset 2026-05-12** — the strongest product evidence of pipeline-era generation retired in favor of natively multimodal generation.
- Google's counterpoint: Gemini 3 Pro Image ("Nano Banana Pro") competes on text rendering and follow-up instruction adherence; 1B+ images generated via Gemini within weeks of launch.
- **Why gpt-image-1.5 matters structurally:** the ~4x speed gain and 20% price cut are the economics of native multimodality (one fused trunk doing text understanding + image synthesis) vs the DALL-E pipeline (separate diffusion stack with its own captioning/conditioning overhead). The RGBA-native alpha channel is the concrete capability proof: the model emits transparency as a native output property, not as a post-processed mask.
- The DALL-E 2/3 sunset (2026-05-12) is the single strongest product evidence in this knowledge base of the "assembly bricolage" ending: OpenAI retired working, revenue-generating image products to force migration onto the native multimodal family — a deliberate product sacrifice for architectural consolidation, the inverse of Sora (where the architecture survived and the product died).
- **Nano Banana family map (Sept 2026):** original Nano Banana (Gemini 2.5 Flash Image) → Nano Banana 2 → Nano Banana Pro (Gemini 3 Pro Image; text rendering and instruction adherence; 1B+ images via Gemini within weeks) → **Nano Banana 2 Lite** (Gemini 3.1 Flash Lite; entry tier; ~2026-07-01/02; 4s text-to-image at ~1K; $0.034/1K images). Google explicitly encouraged developers on the original Nano Banana to upgrade — a managed migration inside the image family, mirroring the DALL-E→GPT Image forced migration on the OpenAI side.
- Nano Banana 2 Lite pricing ($0.25/$1.50 per M tokens, gigazine) puts entry-tier image generation at roughly an order of magnitude below gpt-image-1.5's low end ($0.011 per 1024x1024) on a per-image basis — the 2026 image market is tiered: flagship quality (Pro/Image 1.5) vs commodity entry (2 Lite), with the middle occupied by Mini variants (gpt-image-1 Mini $0.005–$0.036).

### 12.7 HiDream-O1-Video-1.0 and MiniMax H3 (capability notes; licensing → §11)

