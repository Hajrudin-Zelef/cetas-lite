---
id: ai-industry-kb-2026-wave6/18-multimodal-generation-video-image-and-audio/image-generation-gpt-image-2
title: "Image generation — GPT Image 2"
domain: multimodal-generation-video-image-and-audio
role: deep-dive
task: multimodal
actors: ["OpenAI"]
dates: ["2023-03", "2026-03", "2026-03-19", "2026-04", "2026-04-21", "2026-04-22", "2026-05", "2026-06-11", "2026-07", "2026-08", "2026-10-23", "2026-12-01"]
keywords: ["chatgpt", "claude", "cost", "diffusion", "disclosure", "gemini", "gpu", "pricing", "research", "text-to-image"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8843, 8864]
section: "§18. Multimodal Generation: Video, Image, and Audio"
delta_of: ai-industry-kb-2026
sha256: fb5be1399dd29f615a2253b4b586c6dea397b6c75ff4ab49e8517ccfb04d4258
---

# Image generation — GPT Image 2

- [SECONDARY] Midjourney V8 Alpha launched March 2026 (release-notes coverage dated 2026-03-19) as an early public test of the V8 model; V8 is reported ~5x faster than before and understands aesthetics via personalization, style references (srefs) and moodboards. (https://aiforautomation.io/news/2026-03-19-midjourney-v8-5x-faster-4x-more-expensive ; https://releasebot.io/updates/midjourney)
- [SECONDARY] V8 introduced `--hd` mode rendering natively at 2K resolution and `--q 4` mode for extra coherence; premium jobs (--hd, --q 4, srefs, moodboards) cost 4x the standard generation at launch. (https://releasebot.io/updates/midjourney ; https://aiforautomation.io/news/2026-03-19-midjourney-v8-5x-faster-4x-more-expensive)
- [SECONDARY] At the V8 Alpha launch, Relax mode was not yet supported — Midjourney said it was building a new server cluster for Relax plus cheaper render modes, with no timeline. (https://releasebot.io/updates/midjourney ; https://aiforautomation.io/news/2026-03-19-midjourney-v8-5x-faster-4x-more-expensive)
- (single community coverage) [COMMUNITY] V8.1 became the default model on June 11, 2026, and V8.2 followed in July 2026 as a refinement pass (bolder aesthetics, fewer weak outputs, better personalization). (https://github.com/linda-mhmd/ai-solutions-wiki/blob/HEAD/content/comparisons/midjourney-vs-dalle-vs-stable-diffusion.md)
- (single secondary coverage) [SECONDARY] CONTRADICTION — Midjourney V8.1/V8.2 default and release dates conflict across catalogs — preserved pending an authoritative changelog check.
- [SECONDARY] Text rendering in images improved in V8 (specify text in "quotes"); the community guide notes Latin-alphabet text remains optimal and Arabic text can be inserted with quotation marks but is less accurate. (https://aiforautomation.io/news/2026-03-19-midjourney-v8-5x-faster-4x-more-expensive ; https://github.com/dayahimour/ai-dayahimour/blob/HEAD/src/content/blog-en/midjourney-v8.md)
- [SECONDARY] 2026 subscription structure: Basic $10/month (3.3 Fast GPU hours), Standard $30/month (15 Fast GPU hours + unlimited Relax), Pro $60/month (30 Fast GPU hours + Stealth mode), Mega $120/month (60 Fast GPU hours); annual billing is ~20% cheaper; there has never been a free tier since the trial was removed in March 2023. (https://github.com/linda-mhmd/ai-solutions-wiki/blob/HEAD/content/comparisons/midjourney-vs-dalle-vs-stable-diffusion.md ; https://aiwisepicks.com/tools/midjourney/)
- [SECONDARY] Midjourney offers no official API — third-party wrappers exist but violate the terms of service; Discord is now an optional sign-in rather than a requirement, with most users on the browser interface at midjourney.com. (https://lushbinary.com/blog/ai-image-generation-comparison-midjourney-gpt-flux/ ; https://github.com/linda-mhmd/ai-solutions-wiki/blob/HEAD/content/comparisons/midjourney-vs-dalle-vs-stable-diffusion.md)

### Image generation — GPT Image 2

- [COMMUNITY] `gpt-image-2` snapshot `gpt-image-2-2026-04-21` became the public evidence of release on April 21, 2026 (ChatGPT/Codex), with direct API general availability in early May 2026 and OpenAI developer docs showing the model page, Images API examples and pricing rows. (https://blog.laozhang.ai/en/posts/gpt-image-2-api-release-date ; https://dev.to/tokenmixai/gpt-image-2-api-developer-guide-pricing-thinking-mode-and-production-integration-2026-28p5)
- [COMMUNITY] Per-token pricing via OpenAI's official pricing page: $5/$10 per M text input/output tokens, $8/$2 per M image input/cached input tokens, $30 per M image output tokens; indicative per-image costs at 1024×1024: low $0.006, medium $0.053, high $0.211; portrait 1024×1536: $0.005/$0.041/$0.165. (https://dev.to/tokenmixai/gpt-image-2-api-developer-guide-pricing-thinking-mode-and-production-integration-2026-28p5 ; https://blog.laozhang.ai/en/posts/gpt-image-2-api-release-date ; https://github.com/shaharsha/claude-skills/blob/HEAD/skills/image-generation/reference/pricing.md)
- (single community coverage) [COMMUNITY] Size constraints: free-form sizes with each edge a multiple of 16, aspect ratio ≤ 3:1, max edge < 3840px, total pixels 655,360–8,294,400; endpoint `POST /v1/images/generations` with `size` free-form and `quality: low|medium|high|auto`. (https://github.com/backblaze-labs/genblaze/blob/HEAD/docs/exec-plans/completed/openai-image-model-expansion.md ; https://github.com/arthur-diego/linkedin-rag-content/blob/HEAD/docs/research/image-models-2026.md)
- (single community coverage) [COMMUNITY] Two modes: `instant` (default) and `thinking` (opt-in, with web-search grounding); up to 8 images per call with character/object continuity. (https://dev.to/tokenmixai/gpt-image-2-api-developer-guide-pricing-thinking-mode-and-production-integration-2026-28p5)
- (single community coverage) [COMMUNITY] Per the backblaze-labs SDK audit (sourced to developers.openai.com + openai-python SDK reference, 2026-04-22): gpt-image-2 supports edits with mask but no transparent background, unlike gpt-image-1.5 which supports it; gpt-image-2 responds in b64_json. (https://github.com/backblaze-labs/genblaze/blob/HEAD/docs/exec-plans/completed/openai-image-model-expansion.md)
- (single community coverage) [COMMUNITY] CONTRADICTION — An April 2026 SDK-derived audit (backblaze-labs/genblaze) stated OpenAI had not published public per-image rates for gpt-image-2 (`cost_usd=None` until disclosed), while later community pricing references (April 2026) and the laozhang release-watch summary report official per-token/per-image rows. The disclosure appears to have happened in stages; both states are preserved with their dates.
- (single community coverage) [COMMUNITY] Independent Aug-2026 research ranks gpt-image-2 #1 on LMArena/Artificial Analysis text-to-image and editing arenas, with ~99% text accuracy on a scientific-diagram head-to-head including 8pt sub-labels, versus ~95% for Nano Banana Pro (`gemini-3-pro-image`, $0.134/image) with typos on small labels. (https://github.com/arthur-diego/linkedin-rag-content/blob/HEAD/docs/research/image-models-2026.md)
- (single community coverage) [COMMUNITY] The OpenAI image lineup as of August 2026: gpt-image-2 flagship (thinking mode, multilingual text, up to 3840px edge), gpt-image-1.5 (API removal reported for 2026-12-01), gpt-image-1 (retirement reported for 2026-10-23), gpt-image-1-mini (cheap tier, weak text fidelity). (https://github.com/arthur-diego/linkedin-rag-content/blob/HEAD/docs/research/image-models-2026.md)

### Image generation — Ideogram 4.0

