---
id: ai-industry-kb-2026-wave6/14-google-gemini-and-gemma/nano-banana-2-gemini-3-1-flash-image
title: "Nano Banana 2 (Gemini 3.1 Flash Image)"
domain: google-gemini-and-gemma
role: deep-dive
task: actor-profile
actors: ["Google", "United States"]
dates: ["2025-01", "2025-06", "2025-08", "2025-08-05", "2026-01-29", "2026-02-26", "2026-03-15", "2026-04", "2026-06", "2026-09-07"]
keywords: ["gemini", "agent", "agents", "claude", "cost", "leaderboard", "memory", "moe", "pricing", "reasoning", "research", "text-to-image"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6955, 7033]
section: "§14. Google: Gemini and Gemma"
delta_of: ai-industry-kb-2026
sha256: 6c737961313d6dcfcb16264d0ff4993b2a05c3abcf654487c338f5baba5d7d35
---

# Nano Banana 2 (Gemini 3.1 Flash Image)

### Nano Banana 2 (Gemini 3.1 Flash Image)
- Launched 2026-02-26 [SECONDARY](https://techcrunch.com/2026/02/26/google-launches-nano-banana-2-model-with-faster-image-generation/) [SECONDARY](https://github.com/the-focus-ai/nano-banana-cli/blob/HEAD/reports/2026-03-15-google-gemini-image-video-generation-models.md)
- Formally Gemini 3.1 Flash Image (preview ID `gemini-3.1-flash-image-preview`) [SECONDARY](https://github.com/the-focus-ai/nano-banana-cli/blob/HEAD/reports/2026-03-15-google-gemini-image-video-generation-models.md)
- Replaces the prior Flash image model as default in the Gemini app [SECONDARY](https://techcrunch.com/2026/02/26/google-launches-nano-banana-2-model-with-faster-image-generation/)
- Default in app across Fast, Thinking, Pro modes [SECONDARY](https://github.com/the-focus-ai/nano-banana-cli/blob/HEAD/reports/2026-03-15-google-gemini-image-video-generation-models.md)
- Default in Flow [SECONDARY](https://github.com/the-focus-ai/nano-banana-cli/blob/HEAD/reports/2026-03-15-google-gemini-image-video-generation-models.md)
- MoE transformer based on Gemini 3 Flash (rendering model undisclosed) [SECONDARY](https://www.deeplearning.ai/the-batch/nano-banana-2-aka-gemini-3-1-flash-image-makes-edits-easier-and-faster)
- Reported performance: ~2.9–4× faster than Nano Banana Pro [SECONDARY](https://medium.com/@aiamplify22/google-nano-banana-2-everything-you-need-to-know-about-the-fastest-ai-image-generator-in-2026-cd32321b74ed)
- ~50% cheaper per image than Pro [SECONDARY](https://medium.com/@aiamplify22/google-nano-banana-2-everything-you-need-to-know-about-the-fastest-ai-image-generator-in-2026-cd32321b74ed)
- ~95% of Pro's visual quality per one review (all secondary) [SECONDARY](https://medium.com/@aiamplify22/google-nano-banana-2-everything-you-need-to-know-about-the-fastest-ai-image-generator-in-2026-cd32321b74ed)
- Resolutions 512px square to 4K [SECONDARY](https://www.deeplearning.ai/the-batch/nano-banana-2-aka-gemini-3-1-flash-image-makes-edits-easier-and-faster)
- 14 aspect ratios incl. extreme 1:4/4:1 (and 1:8/8:1 per one source) [SECONDARY](https://www.deeplearning.ai/the-batch/nano-banana-2-aka-gemini-3-1-flash-image-makes-edits-easier-and-faster) [SECONDARY](https://github.com/nothingfancyai/skills/blob/HEAD/plugins/nano-banana-prompting/skills/nano-banana-prompting/SKILL.md)
- 4–6 s per image [SECONDARY](https://www.deeplearning.ai/the-batch/nano-banana-2-aka-gemini-3-1-flash-image-makes-edits-easier-and-faster)
- API pricing: $0.50/M input tokens [SECONDARY](https://www.deeplearning.ai/the-batch/nano-banana-2-aka-gemini-3-1-flash-image-makes-edits-easier-and-faster)
- Per image: $0.045 (512px) [SECONDARY](https://www.deeplearning.ai/the-batch/nano-banana-2-aka-gemini-3-1-flash-image-makes-edits-easier-and-faster)
- Per image: $0.067 (1K) [SECONDARY](https://www.deeplearning.ai/the-batch/nano-banana-2-aka-gemini-3-1-flash-image-makes-edits-easier-and-faster)
- Per image: $0.101 (2K) [SECONDARY](https://www.deeplearning.ai/the-batch/nano-banana-2-aka-gemini-3-1-flash-image-makes-edits-easier-and-faster)
- Per image: $0.151 (4K) [SECONDARY](https://www.deeplearning.ai/the-batch/nano-banana-2-aka-gemini-3-1-flash-image-makes-edits-easier-and-faster)
- Image search grounding free for first 5,000 queries/month, then $14 per 1,000 queries [SECONDARY](https://www.deeplearning.ai/the-batch/nano-banana-2-aka-gemini-3-1-flash-image-makes-edits-easier-and-faster)
- INPUT CONTEXT CONTRADICTION: DeepLearning.AI says up to 1M input tokens; a prompting-skill source says 131,072 max — unresolved; do not cite a single figure without the caveat [UNVERIFIED](https://www.deeplearning.ai/the-batch/nano-banana-2-aka-gemini-3-1-flash-image-makes-edits-easier-and-faster) [UNVERIFIED](https://github.com/nothingfancyai/skills/blob/HEAD/plugins/nano-banana-prompting/skills/nano-banana-prompting/SKILL.md)
- Character consistency up to 5 characters [SECONDARY](https://news.saerio.com/google-launches-nano-banana-2-model-with-faster-image-generation/) [SECONDARY](https://github.com/nothingfancyai/skills/blob/HEAD/plugins/nano-banana-prompting/skills/nano-banana-prompting/SKILL.md)
- Up to 14 reference objects per workflow [SECONDARY](https://github.com/nothingfancyai/skills/blob/HEAD/plugins/nano-banana-prompting/skills/nano-banana-prompting/SKILL.md)
- Two reasoning levels (minimal/high) [SECONDARY](https://github.com/nothingfancyai/skills/blob/HEAD/plugins/nano-banana-prompting/skills/nano-banana-prompting/SKILL.md)
- Multilingual text rendering [SECONDARY](https://news.saerio.com/google-launches-nano-banana-2-model-with-faster-image-generation/)
- Knowledge cutoff January 2025 per one source [SECONDARY](https://github.com/nothingfancyai/skills/blob/HEAD/plugins/nano-banana-prompting/skills/nano-banana-prompting/SKILL.md)
- Arena standing: leads Arena.ai text-to-image human-preference leaderboard [SECONDARY](https://www.deeplearning.ai/the-batch/nano-banana-2-aka-gemini-3-1-flash-image-makes-edits-easier-and-faster)
- Ranks 2nd–3rd on Artificial Analysis Text-to-Image and Image-Editing leaderboards after GPT Image 1.5 and Nano Banana Pro [SECONDARY](https://www.deeplearning.ai/the-batch/nano-banana-2-aka-gemini-3-1-flash-image-makes-edits-easier-and-faster)
- SynthID watermarking + C2PA Content Credentials by default [SECONDARY](https://techfundingnews.com/nano-banana-2-rolls-out-is-it-geminis-fastest-image-model/)
- SynthID verification in Gemini hit 20M+ uses since November [SECONDARY](https://techfundingnews.com/nano-banana-2-rolls-out-is-it-geminis-fastest-image-model/)
- Default for image generation in Google Search AI Mode / Lens across 141 countries on Google app and web (desktop + mobile) [SECONDARY](https://techcrunch.com/2026/02/26/google-launches-nano-banana-2-model-with-faster-image-generation/)
- Pro subscribers can still manually select Nano Banana Pro for higher-accuracy regeneration via the three-dot menu [SECONDARY](https://news.saerio.com/google-launches-nano-banana-2-model-with-faster-image-generation/)

### Veo 3.1
- API pricing per second of video (audio bundled): Standard $0.40 (720p/1080p) [SECONDARY](https://www.scriptbyai.com/gpt-gemini-claude-pricing/) [SECONDARY](https://github.com/achint08/90210/blob/HEAD/docs/superpowers/research/veo-3.1-deep-dive.md)
- Standard 4K: $0.60/s [SECONDARY](https://www.scriptbyai.com/gpt-gemini-claude-pricing/) [SECONDARY](https://github.com/achint08/90210/blob/HEAD/docs/superpowers/research/veo-3.1-deep-dive.md)
- Fast tier: $0.15/s per one pricing page [SECONDARY](https://www.scriptbyai.com/gpt-gemini-claude-pricing/)
- Fast tier: $0.10/s per a May-2026 deep-dive (CONTRADICTION — note both; direct Gemini API) [SECONDARY](https://github.com/achint08/90210/blob/HEAD/docs/superpowers/research/veo-3.1-deep-dive.md)
- Lite/Light: $0.05 (720p) / $0.08 (1080p), no Extend, no 4K [SECONDARY](https://genrates.com/reports/state-of-ai-video-pricing-aug-2026.pdf)
- Extend cost = 7× the per-second rate (no surcharge) [SECONDARY](https://github.com/achint08/90210/blob/HEAD/docs/superpowers/research/veo-3.1-deep-dive.md) [SECONDARY](https://zenn.dev/saan/articles/f26ad506df2ebf)
- First/last-frame and reference modes priced as standard seconds [SECONDARY](https://github.com/achint08/90210/blob/HEAD/docs/superpowers/research/veo-3.1-deep-dive.md)
- Billed only for successfully generated videos [SECONDARY](https://github.com/achint08/90210/blob/HEAD/docs/superpowers/research/veo-3.1-deep-dive.md)
- Video retention: 2 days server-side [SECONDARY](https://github.com/achint08/90210/blob/HEAD/docs/superpowers/research/veo-3.1-deep-dive.md) [SECONDARY](https://zenn.dev/saan/articles/f26ad506df2ebf)
- Within-model spread: $0.030/s (Lite 720p no audio) to $0.60/s (Standard 4K with audio) — a 20× range on the same first-party API, verified Aug 1 2026 [SECONDARY](https://genrates.com/reports/state-of-ai-video-pricing-aug-2026.pdf)
- Native audio: ambience, SFX, short dialogue [SECONDARY](https://www.archyde.com/google-launches-ai-video-model-veo-3-1-enhancing-enterprise-capabilities-in-flow-and-api-integration/)
- Up to 3 reference images [SECONDARY](https://www.archyde.com/google-launches-ai-video-model-veo-3-1-enhancing-enterprise-capabilities-in-flow-and-api-integration/)
- First/last-frame interpolation [SECONDARY](https://www.archyde.com/google-launches-ai-video-model-veo-3-1-enhancing-enterprise-capabilities-in-flow-and-api-integration/)
- Scene extension: 8s clips extendable to 2.5+ minutes via up to 20 extensions [SECONDARY](https://www.archyde.com/google-launches-ai-video-model-veo-3-1-enhancing-enterprise-capabilities-in-flow-and-api-integration/)
- Insert/Remove scene functions [SECONDARY](https://www.archyde.com/google-launches-ai-video-model-veo-3-1-enhancing-enterprise-capabilities-in-flow-and-api-integration/)
- Gemini 2.0 Flash/Flash-Lite scheduled for retirement June 2026 [SECONDARY](https://github.com/the-focus-ai/nano-banana-cli/blob/HEAD/reports/2026-03-15-google-gemini-image-video-generation-models.md)
- Veo 2 preview/exp endpoints scheduled for removal April 2026 [SECONDARY](https://github.com/the-focus-ai/nano-banana-cli/blob/HEAD/reports/2026-03-15-google-gemini-image-video-generation-models.md)

### Imagen 4
- GA pricing per image: Imagen 4 Fast $0.02 [SECONDARY](https://www.scriptbyai.com/gpt-gemini-claude-pricing/)
- Standard $0.04 [SECONDARY](https://www.scriptbyai.com/gpt-gemini-claude-pricing/)
- Ultra $0.06 [SECONDARY](https://www.scriptbyai.com/gpt-gemini-claude-pricing/)
- Imagen 3: $0.03 [SECONDARY](https://www.scriptbyai.com/gpt-gemini-claude-pricing/)
- Timeline: preview June 2025 → GA August 2025 (Fast/Standard/Ultra) [SECONDARY](https://github.com/the-focus-ai/nano-banana-cli/blob/HEAD/reports/2026-03-15-google-gemini-image-video-generation-models.md)
- Preview models deprecated Nov 2025 [SECONDARY](https://github.com/the-focus-ai/nano-banana-cli/blob/HEAD/reports/2026-03-15-google-gemini-image-video-generation-models.md)
- Preview-06-06 models removed Feb 2026 [SECONDARY](https://github.com/the-focus-ai/nano-banana-cli/blob/HEAD/reports/2026-03-15-google-gemini-image-video-generation-models.md)

### Genie 3 / Project Genie
- Genie 3: 11B-parameter autoregressive transformer world model [SECONDARY](https://justoborn.com/deepmind-genie-3-worlds/) [SECONDARY](https://almcorp.com/blog/google-deepmind-project-genie-technical-analysis-applications/)
- Generates real-time interactive 3D worlds from text prompts at 720p, 24 fps [SECONDARY](https://justoborn.com/deepmind-genie-3-worlds/)
- Spatial memory ~1 minute [SECONDARY](https://justoborn.com/deepmind-genie-3-worlds/) [SECONDARY](https://almcorp.com/blog/google-deepmind-project-genie-technical-analysis-applications/)
- Announced 2025-08-05 [SECONDARY](https://justoborn.com/deepmind-genie-3-worlds/)
- Project Genie launched 2026-01-29 at labs.google/projectgenie [SECONDARY](https://soravideo.art/blog/genie-3-google-deepmind-world-model) [SECONDARY](https://mlq.ai/news/google-launches-project-genie-ai-prototype-for-generating-interactive-game-worlds/)
- Three modes: world sketching, exploration, remixing [SECONDARY](https://soravideo.art/blog/genie-3-google-deepmind-world-model)
- Nano Banana Pro generates preview images [SECONDARY](https://mlq.ai/news/google-launches-project-genie-ai-prototype-for-generating-interactive-game-worlds/)
- Sessions capped at 60 seconds [SECONDARY](https://soravideo.art/blog/genie-3-google-deepmind-world-model) [SECONDARY](https://mlq.ai/news/google-launches-project-genie-ai-prototype-for-generating-interactive-game-worlds/)
- Access: Google AI Ultra subscribers only [SECONDARY](https://almcorp.com/blog/google-deepmind-project-genie-technical-analysis-applications/)
- US only [SECONDARY](https://almcorp.com/blog/google-deepmind-project-genie-technical-analysis-applications/)
- 18+ [SECONDARY](https://almcorp.com/blog/google-deepmind-project-genie-technical-analysis-applications/)
- $249.99/mo plan context for Ultra at that time [SECONDARY](https://almcorp.com/blog/google-deepmind-project-genie-technical-analysis-applications/)
- Market reaction: Unity Software −21% and Take-Two −9% in one trading day [SECONDARY](https://abit.ee/en/games/google-project-genie-unity-artificial-intelligence-gaming-industry-stocks-game-generation-en)
- Google blocked requests mentioning well-known characters ("third-party rights holders") [SECONDARY](https://abit.ee/en/games/google-project-genie-unity-artificial-intelligence-gaming-industry-stocks-game-generation-en)
- DeepMind's SIMA agent has been tested inside Genie 3-generated worlds — positioned as unlimited synthetic training environments for AI agents/robots [SECONDARY](https://justoborn.com/deepmind-genie-3-worlds/)
- No public weights/API [SECONDARY](https://thetoolsverse.com/tools/genie-3-google)
- One review verified on 2026-09-07 that deepmind.google/models/genie contains no download/access/waitlist path — research model, not a deployable product [SECONDARY](https://thetoolsverse.com/tools/genie-3-google)

