---
id: ai-industry-kb-2026-wave6/14-google-gemini-and-gemma/veo-3-1
title: "Veo 3.1"
domain: google-gemini-and-gemma
role: deep-dive
task: actor-profile
actors: ["Google", "United States"]
dates: ["2025-06", "2025-08", "2025-08-05", "2026-01-29", "2026-03-15", "2026-04", "2026-06", "2026-09-07"]
keywords: ["agent", "agents", "claude", "cost", "gemini", "memory", "pricing", "research", "training"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6987, 7033]
section: "§14. Google: Gemini and Gemma"
delta_of: ai-industry-kb-2026
sha256: 8dcb7f3a301a87bd68ef1d4c7edcc544eec68c15e53647d6d0ac65078149e01f
---

# Veo 3.1

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

