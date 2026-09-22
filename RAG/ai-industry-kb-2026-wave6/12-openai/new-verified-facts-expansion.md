---
id: ai-industry-kb-2026-wave6/12-openai/new-verified-facts-expansion
title: "New verified facts — expansion"
domain: openai
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Cerebras", "Google", "Microsoft", "Nvidia", "OpenAI", "OpenRouter"]
dates: ["2024-06", "2025-04-14", "2025-05", "2025-07-14", "2025-08", "2025-08-31", "2025-12-11", "2026-02-05", "2026-02-12", "2026-02-13", "2026-02-24", "2026-03-17", "2026-09-11"]
keywords: ["chatgpt", "claude", "context window", "cost", "foundry", "gemini", "glm", "gpt-5.6", "gpus", "inference", "latency", "mcp"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5779, 5864]
section: "§12. OpenAI"
sha256: 01c04255c8d58adcc10a32577fb68d673a0754c839cdb2a33ac96682d4f5d302
---

# New verified facts — expansion

### New verified facts — expansion

### GPT-4.1 — original release spec sheet (absent from baseline §12, which covers only the 2026 retirement)
- GPT-4.1 family launched 2025-04-14 with three variants: GPT-4.1, GPT-4.1 mini, GPT-4.1 nano [SECONDARY](https://www.aicloudit.com/blog/ai/introducing-gpt-4-1/) [SECONDARY](http://techtarget.com/whatis/feature/GPT-41-explained-Everything-you-need-to-know)
- API-only at launch — not available in ChatGPT; described as trained exclusively for developers [SECONDARY](https://beebom.com/openai-launch-gpt-4-1-models-not-coming-chatgpt/)
- All three variants carry a 1M-token context window (~750,000 words), including nano [SECONDARY](http://en.zicos.com/tech/i32171379-OpenAI-Unveils-Coding-Focused-GPT-41-While-Phasing-Out-GPT-45.html) [SECONDARY](http://techtarget.com/whatis/feature/GPT-41-explained-Everything-you-need-to-know)
- Knowledge cutoff June 2024 across the family [SECONDARY](http://en.zicos.com/tech/i32171379-OpenAI-Unveils-Coding-Focused-GPT-41-While-Phasing-Out-GPT-45.html) [SECONDARY](http://techtarget.com/whatis/feature/GPT-41-explained-Everything-you-need-to-know)
- Flagship API pricing at launch: $2.00/M input, $8.00/M output [SECONDARY](http://en.zicos.com/tech/i32171379-OpenAI-Unveils-Coding-Focused-GPT-41-While-Phasing-Out-GPT-45.html) [SECONDARY](https://www.techloy.com/openai-launches-gpt-4-1-and-it-is-focused-on-performance-and-code/)
- Mini pricing: $0.40/M input, $1.60/M output [SECONDARY](https://www.gadgetreview.com/openais-gpt-4-1-hits-chatgpt-better-code) [SECONDARY](https://www.aicloudit.com/blog/ai/introducing-gpt-4-1/)
- Nano pricing: $0.10/M input, $0.40/M output — described by OpenAI as its cheapest and fastest model ever at launch [SECONDARY](https://www.gadgetreview.com/openais-gpt-4-1-hits-chatgpt-better-code) [SECONDARY](https://www.aicloudit.com/blog/ai/introducing-gpt-4-1/)
- GPT-4.1 scored 54.6% on SWE-bench Verified (vendor-reported via launch coverage) vs 33% for GPT-4o and 38% for GPT-4.5 [SECONDARY](http://en.zicos.com/tech/i32171379-OpenAI-Unveils-Coding-Focused-GPT-41-While-Phasing-Out-GPT-45.html) [SECONDARY](https://www.techloy.com/openai-launches-gpt-4-1-and-it-is-focused-on-performance-and-code/)
- Same vendor run: Gemini 2.5 Pro 63.8% and Claude 3.7 Sonnet 62.3% on SWE-bench Verified — GPT-4.1 trailed both [SECONDARY](http://en.zicos.com/tech/i32171379-OpenAI-Unveils-Coding-Focused-GPT-41-While-Phasing-Out-GPT-45.html)
- Vendor-reported code-speed claims: 40% faster code generation, 21% better coding performance, up to 80% lower cost vs GPT-4o [SECONDARY](https://www.aicloudit.com/blog/ai/introducing-gpt-4-1/)
- Video-MME long-video comprehension: 72% for GPT-4.1 vs 65.3% for GPT-4o (vendor-reported via launch coverage) [SECONDARY](http://en.zicos.com/tech/i32171379-OpenAI-Unveils-Coding-Focused-GPT-41-While-Phasing-Out-GPT-45.html)
- Instruction following: +10.5-point gain on Scale's MultiChallenge vs GPT-4o; verbosity reduced ~50% (vendor-reported) [SECONDARY](https://newyorkdawn.com/openai-brings-gpt-4-1-and-4-1-mini-to-chatgpt-what-enterprises-ought-to-know/)
- MMMU multimodal: 75% (per one launch-day comparison table), matching GPT-4.5's 75% and above GPT-4o's 69% [SECONDARY](http://techtarget.com/whatis/feature/GPT-41-explained-Everything-you-need-to-know)
- Tested pre-launch on OpenRouter under the codename "Quasar Alpha" [SECONDARY](https://beebom.com/openai-launch-gpt-4-1-models-not-coming-chatgpt/)
- Modalities: text + image input (text and image), unlike GPT-4.5's text-only input per the same table [SECONDARY](https://beebom.com/openai-launch-gpt-4-1-models-not-coming-chatgpt/) [SECONDARY](http://techtarget.com/whatis/feature/GPT-41-explained-Everything-you-need-to-know)
- OpenAI released no full safety report for GPT-4.1, arguing it is not a "frontier model" — a decision that drew criticism from researchers [SECONDARY](https://www.gadgetreview.com/openais-gpt-4-1-hits-chatgpt-better-code)
- GPT-4.5's API retirement was announced on the same day, effective 2025-07-14 [SECONDARY](http://en.zicos.com/tech/i32171379-OpenAI-Unveils-Coding-Focused-GPT-41-While-Phasing-Out-GPT-45.html) [SECONDARY](https://beebom.com/openai-launch-gpt-4-1-models-not-coming-chatgpt/)
- One later source reports GPT-4.1 and 4.1 mini eventually reached ChatGPT (May 2025) — single-source, treat as [SECONDARY] only [SECONDARY](https://newyorkdawn.com/openai-brings-gpt-4-1-and-4-1-mini-to-chatgpt-what-enterprises-ought-to-know/)
- NOTE: baseline §12 covers only the 2026-02-13 ChatGPT retirement of GPT-4.1/mini; none of the above original-release specs appear there.

### GPT-5.2 — full spec sheet
- GPT-5.2 released 2025-12-11 [SECONDARY](https://www.macrumors.com/2025/12/11/openai-gpt-5-2/) [SECONDARY](https://github.com/petehsu/llm-docs/blob/HEAD/data/docs/OpenAI/docs/changelog.md)
- Available same day in ChatGPT and via API as Instant, Thinking, and Pro modes [SECONDARY](https://www.macrumors.com/2025/12/11/openai-gpt-5-2/)
- Proprietary multimodal transformer; architecture not disclosed; no open weights [SECONDARY](https://www.macrumors.com/2025/12/11/openai-gpt-5-2/)
- Introduced the `xhigh` reasoning effort level [SECONDARY](https://www.macrumors.com/2025/12/11/openai-gpt-5-2/)
- Added concise reasoning summaries [SECONDARY](https://www.macrumors.com/2025/12/11/openai-gpt-5-2/)
- Introduced context compaction [SECONDARY](https://www.macrumors.com/2025/12/11/openai-gpt-5-2/)
- Vendor-reported GDPval: 70.9% for GPT-5.2 vs 38.8% for GPT-5.1 — vendor claim, not independent confirmation [SECONDARY](https://www.pymnts.com/artificial-intelligence-2/2025/openai-says-new-ai-model-gpt-5-2-unlocks-even-more-economic-value/)
- Vendor claim: 30% fewer errors than GPT-5.1 [SECONDARY](https://www.pymnts.com/artificial-intelligence-2/2025/openai-says-new-ai-model-gpt-5-2-unlocks-even-more-economic-value/)
- GPT-5.1 kept as a legacy ChatGPT model for three months after the GPT-5.2 launch [SECONDARY](https://www.macrumors.com/2025/12/11/openai-gpt-5-2/)
- GPT-5.2 was a post-training update on the GPT-5 base, not a full base-model retrain [SECONDARY](https://www.theweek.in/news/sci-tech/2025/12/12/openai-s-flagship-gpt-5-2-for-chatgpt-is-here-all-you-need-to-know.html)

### GPT-5.3-Codex — full spec sheet
- GPT-5.3-Codex: proprietary dense-transformer architecture; parameter count undisclosed [SECONDARY](https://www.datacamp.com/blog/glm-5-vs-gpt-5-3-codex) [SECONDARY](https://llm-stats.com/models/gpt-5.3-codex)
- Context window 400K tokens [SECONDARY](https://llm-stats.com/models/gpt-5.3-codex) [SECONDARY](https://www.cometapi.com/models/openai/gpt-5-3-codex/)
- Max output 128K tokens [SECONDARY](https://llm-stats.com/models/gpt-5.3-codex) [SECONDARY](https://www.cometapi.com/models/openai/gpt-5-3-codex/)
- Inputs: text + images [SECONDARY](https://llm-stats.com/models/gpt-5.3-codex) [SECONDARY](https://www.cometapi.com/models/openai/gpt-5-3-codex/)
- API pricing: $1.75/M input [SECONDARY](https://openrouter.ai/openai/gpt-5.3-codex)
- Cached input: $0.175/M [SECONDARY](https://openrouter.ai/openai/gpt-5.3-codex)
- Output: $14/M [SECONDARY](https://openrouter.ai/openai/gpt-5.3-codex)
- Initially Codex-only on 2026-02-05 [SECONDARY](https://www.neowin.net/news/openais-latest-gpt-53-codex-and-audio-models-now-on-microsoft-foundry/)
- API and Microsoft Foundry availability around 2026-02-24/25 [SECONDARY](https://www.neowin.net/news/openais-latest-gpt-53-codex-and-audio-models-now-on-microsoft-foundry/)
- Supports mid-task steering and longer-running execution [SECONDARY](https://www.datacamp.com/blog/glm-5-vs-gpt-5-3-codex)
- Reported ~25% faster than GPT-5.2-class Codex behavior [SECONDARY](https://www.datacamp.com/blog/glm-5-vs-gpt-5-3-codex)
- OpenRouter standardized (independent, standardized-harness) entries, clearly dated and non-vendor: GPQA Diamond 91.5% [SECONDARY](https://openrouter.ai/openai/gpt-5.3-codex)
- Same listing: HLE 42.5% [SECONDARY](https://openrouter.ai/openai/gpt-5.3-codex)
- Same listing: Terminal-Bench Hard 53.0% [SECONDARY](https://openrouter.ai/openai/gpt-5.3-codex)

### GPT-5.3-Codex-Spark — full spec sheet (absent from baseline §12)
- Launched 2026-02-12 as a research preview [SECONDARY](https://www.technobezz.com/news/openai-launches-gpt-53-codex-spark-on-cerebras-hardware) [SECONDARY](https://particle.news/story/openai-debuts-gpt-53-codex-spark-for-real-time-coding-on-cerebras-chips)
- Speed-optimized variant of GPT-5.3-Codex explicitly built for real-time interactive coding [SECONDARY](https://www.cometapi.com/what-is-gpt-5-3-codex-spark-how-to-use-it/)
- Generation speed: >1,000 tokens per second on the low-latency serving path [SECONDARY](https://www.cometapi.com/what-is-gpt-5-3-codex-spark-how-to-use-it/) [SECONDARY](https://www.technobezz.com/news/openai-launches-gpt-53-codex-spark-on-cerebras-hardware)
- Roughly 15× faster than the full GPT-5.3-Codex (~70 tok/s baseline per one review) [SECONDARY](https://felixng.vercel.app/en/ai-news/gpt53-codex-spark-1000-tokens-per-second)
- Runs on Cerebras Wafer Scale Engine 3 (WSE-3) hardware — OpenAI's first production deployment off Nvidia GPUs [SECONDARY](https://felixng.vercel.app/en/ai-news/gpt53-codex-spark-1000-tokens-per-second) [SECONDARY](https://wccftech.com/openais-latest-codex-model-runs-on-cerebras-infrastructure-hinting-at-a-serious-second-option-for-ai-inference-beyond-nvidia/)
- Context window 128K tokens — much smaller than the full GPT-5.3-Codex's 400K [SECONDARY](https://www.cometapi.com/what-is-gpt-5-3-codex-spark-how-to-use-it/) [SECONDARY](https://www.technobezz.com/news/openai-launches-gpt-53-codex-spark-on-cerebras-hardware)
- Text-only at launch; no multimodal inputs [SECONDARY](https://www.cometapi.com/what-is-gpt-5-3-codex-spark-how-to-use-it/)
- Research-preview access for ChatGPT Pro users via the Codex app, CLI, and VS Code extension [SECONDARY](https://www.cometapi.com/what-is-gpt-5-3-codex-spark-how-to-use-it/) [SECONDARY](https://www.technobezz.com/news/openai-launches-gpt-53-codex-spark-on-cerebras-hardware)
- Limited API access for select design partners under separate rate limits; Spark usage does not count against standard Codex quotas [SECONDARY](https://www.cometapi.com/what-is-gpt-5-3-codex-spark-how-to-use-it/) [SECONDARY](https://particle.news/story/openai-debuts-gpt-53-codex-spark-for-real-time-coding-on-cerebras-chips)
- Vendor serving-stack claims: 80% lower client-server roundtrip overhead, 30% lower per-token overhead, 50% faster time-to-first-token [SECONDARY](https://particle.news/story/openai-debuts-gpt-53-codex-spark-for-real-time-coding-on-cerebras-chips)
- Reported API pricing $5.00/M input and $15.00/M output — single source [SECONDARY](https://creati.ai/ai-news/2026-02-12/openai-gpt-5-3-codex-spark-cerebras-partnership/)
- RETIREMENT: one secondary source says OpenAI announced Spark's retirement on 2026-09-11 (~7 months post-launch), citing declining usage — single-source [SECONDARY](https://aiidelist.com/blog/gpt-5-3-codex-spark-retirement)
- Same retirement source says the low-latency idea continued into GPT-5.6 Sol Ultrafast (up to 750 output tokens/sec on Cerebras) — single-source [SECONDARY](https://aiidelist.com/blog/gpt-5-3-codex-spark-retirement)

### GPT-5.4 — full spec sheet
- GPT-5.4: proprietary model; parameter count undisclosed [SECONDARY](https://github.com/baltwross/ross-claude-toolkit/blob/HEAD/skills/frontier-model-landscape/SKILL.md)
- Total context 1,050,000 tokens [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026) [SECONDARY](https://almcorp.com/blog/gpt-5-4/)
- Typically described as 922K input + 128K output [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- CONTEXT CONTRADICTION: one source describes "272K standard" context while others describe 1.05M maximum; per pricing documentation 272K is a price threshold, not necessarily a hard context limit [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- Knowledge cutoff CONTESTED: one source explicitly labels August 2025 "assumed"; another states 2025-08-31 as fact — treat as unconfirmed [UNVERIFIED](https://almcorp.com/blog/gpt-5-4/)
- Input pricing: $2.50/M [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026) [SECONDARY](https://aitoolbriefing.com/reviews/gpt-5-4-review-2026/)
- Cached input: $0.25/M [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- Output: $15/M [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026) [SECONDARY](https://aitoolbriefing.com/reviews/gpt-5-4-review-2026/)
- GPT-5.4 Pro: $30/M input, $180/M output [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- Above the 272K input threshold: reported $5 input / $22.50 output for standard — a tiered price cliff [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- Above 272K for Pro: reported $60/$270 — same cliff structure at the Pro tier [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- Capabilities: native computer use [SECONDARY](https://almcorp.com/blog/gpt-5-4/)
- Tool search support [SECONDARY](https://almcorp.com/blog/gpt-5-4/)
- MCP support [SECONDARY](https://almcorp.com/blog/gpt-5-4/)
- Adaptive reasoning with five reasoning efforts [SECONDARY](https://almcorp.com/blog/gpt-5-4/)
- Vendor claim: Tool Search reduced tokens by 47% over 250 Scale MCP Atlas tasks across 36 servers [SECONDARY](https://almcorp.com/blog/gpt-5-4/)
- GPT-5.4 Mini reportedly released 2026-03-17 at $0.75/M input, $4.50/M output [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4 Nano reportedly released 2026-03-17 at $0.20/M input, $1.25/M output [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)

