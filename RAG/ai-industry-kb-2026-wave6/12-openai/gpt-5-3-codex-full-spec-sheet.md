---
id: ai-industry-kb-2026-wave6/12-openai/gpt-5-3-codex-full-spec-sheet
title: "GPT-5.3-Codex — full spec sheet"
domain: openai
role: deep-dive
task: actor-profile
actors: ["Cerebras", "Microsoft", "Nvidia", "OpenAI", "OpenRouter"]
dates: ["2025-08", "2025-08-31", "2026-02-05", "2026-02-12", "2026-02-24", "2026-03-17", "2026-09-11"]
keywords: ["chatgpt", "claude", "context window", "foundry", "glm", "gpt-5.6", "gpus", "inference", "latency", "mcp", "multimodal", "nvidia"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5814, 5864]
section: "§12. OpenAI"
delta_of: ai-industry-kb-2026
sha256: a93086f2ace2f60b0b88896ba05494fc1bb6d7efce31822d306fa7d1be08bd85
---

# GPT-5.3-Codex — full spec sheet

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

