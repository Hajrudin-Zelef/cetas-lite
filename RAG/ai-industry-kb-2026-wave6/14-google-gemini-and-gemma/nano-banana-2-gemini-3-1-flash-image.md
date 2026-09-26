---
id: ai-industry-kb-2026-wave6/14-google-gemini-and-gemma/nano-banana-2-gemini-3-1-flash-image
title: "Nano Banana 2 (Gemini 3.1 Flash Image)"
domain: google-gemini-and-gemma
role: deep-dive
task: actor-profile
actors: ["Google"]
dates: ["2025-01", "2026-02-26", "2026-03-15"]
keywords: ["gemini", "leaderboard", "moe", "pricing", "reasoning", "text-to-image", "watermarking"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6955, 6986]
section: "§14. Google: Gemini and Gemma"
delta_of: ai-industry-kb-2026
sha256: 8dc53cb5fa5d51dc49f4e6c1796755a11ae5e5624f2c25bddb08120ce0d5c2fe
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

