---
id: ai-industry-kb-2026-wave6/14-google-gemini-and-gemma/gemini-3-1-pro-preview
title: "Gemini 3.1 Pro Preview"
domain: google-gemini-and-gemma
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Google", "Microsoft", "OpenRouter"]
dates: ["2026-02-19", "2026-03-26", "2026-05-04", "2026-06-14"]
keywords: ["gemini", "agents", "agi", "attribution", "benchmarks", "claude", "consumer", "copilot", "cost", "moe", "multimodal", "opus 4"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6887, 6920]
section: "§14. Google: Gemini and Gemma"
delta_of: ai-industry-kb-2026
sha256: 4c6e06e803d1606ca2bc3951e9edefa250025f5e66e296139acabd21e8001540
---

# Gemini 3.1 Pro Preview

### Gemini 3.1 Pro Preview
- Released 2026-02-19 (Preview) [SECONDARY](https://openrouter.ai/google/gemini-3.1-pro-preview) [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- Model ID `gemini-3.1-pro-preview` [SECONDARY](https://openrouter.ai/google/gemini-3.1-pro-preview)
- MoE multimodal reasoning model [SECONDARY](https://blog.wentuo.ai/en/gemini-3-1-pro-vs-claude-opus-4-6-comparison-en.html) [SECONDARY](https://iconpolls.com/blogs/gemini-31-pro-review-in-2026-release-date-price-benchmarks-model-name-user-experience-and-faqs)
- 1M input context [SECONDARY](https://blog.wentuo.ai/en/gemini-3-1-pro-vs-claude-opus-4-6-comparison-en.html)
- 64K max output [SECONDARY](https://blog.wentuo.ai/en/gemini-3-1-pro-vs-claude-opus-4-6-comparison-en.html)
- Inputs: text/images/audio/video/PDFs/code repos [SECONDARY](https://blog.wentuo.ai/en/gemini-3-1-pro-vs-claude-opus-4-6-comparison-en.html)
- Video up to 1 hour, audio up to 8.4 hours per one comparison [SECONDARY](https://blog.wentuo.ai/en/gemini-3-1-pro-vs-claude-opus-4-6-comparison-en.html)
- New thinking levels (Minimal/Low/Medium/High) exposing cost/reasoning tradeoffs to the caller [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026) [SECONDARY](https://github.com/mattnb/fireside/blob/HEAD/docs/eval-gemini-31-pro.md)
- API pricing for prompts ≤200K tokens: $2/M input, $12/M output [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- API pricing above 200K tokens: $4/M input, $18/M output [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- Context caching up to 75% discount [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- PRICE CONTRADICTION: one comparison table (dated June 14, 2026) lists $1.50/$9.00 — either a later price cut or an error; the $2/$12 figure is multiply corroborated [UNVERIFIED](https://tech-insider.org/claude-opus-4-8-vs-gpt-5-5-vs-gemini-3-1-pro-2026/)
- Vendor-reported benchmarks: ARC-AGI-2 77.1% [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- Vendor: SWE-Bench Verified 80.6% [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- Independent/standardized entries (OpenRouter): GPQA Diamond 94.4–95.3%, provider-dependent [SECONDARY](https://openrouter.ai/google/gemini-3.1-pro-preview)
- Provider split: Vertex 95.3%, AI Studio 94.8%, auto-routing 94.4% — report provider alongside the number [SECONDARY](https://openrouter.ai/google/gemini-3.1-pro-preview)
- TAU-Bench: 77.3/74.7% [SECONDARY](https://openrouter.ai/google/gemini-3.1-pro-preview)
- One source reports GPQA Diamond 94.3% and APEX-Agents 33.5% (vs Opus 4.6's 29.8%) — attribution differs from OpenRouter's 94.4%; keep separate [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- Performance: ~104–116 t/s (Artificial Analysis) [SECONDARY](https://iconpolls.com/blogs/gemini-31-pro-review-in-2026-release-date-price-benchmarks-model-name-user-experience-and-faqs)
- 112 tps P50 best provider [SECONDARY](https://openrouter.ai/google/gemini-3.1-pro-preview)
- OpenRouter weighted avg paid: $1.224 input / $10.96 output per M [SECONDARY](https://openrouter.ai/google/gemini-3.1-pro-preview)
- Gemini 3 Pro discontinued on Vertex AI 2026-03-26, forcing enterprise customers onto the 3.1 Pro track [SECONDARY](https://iconpolls.com/blogs/gemini-31-pro-review-in-2026-release-date-price-benchmarks-model-name-user-experience-and-faqs)
- Consumer rollout to Gemini app Pro/Ultra limits expanded 2026-05-04 [SECONDARY](https://iconpolls.com/blogs/gemini-31-pro-review-in-2026-release-date-price-benchmarks-model-name-user-experience-and-faqs)
- NotebookLM picked it up for paid users [SECONDARY](https://iconpolls.com/blogs/gemini-31-pro-review-in-2026-release-date-price-benchmarks-model-name-user-experience-and-faqs)
- Available in Google AI Studio (free with rate limits) [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- Available via Gemini API [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- Available on Vertex AI [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- Available in Gemini CLI [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- Available in GitHub Copilot [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- Available in Google Antigravity [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- Available in Android Studio [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- Available in VS Code (extension) [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)

