---
id: ai-industry-kb-2026-wave6/12-openai/sora-wind-down
title: "Sora wind-down"
domain: openai
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "DeepSeek", "EU", "ExploitGym", "Nvidia", "OpenAI"]
dates: ["2025-08", "2025-12-16", "2026-02-13", "2026-03", "2026-03-24", "2026-03-31", "2026-04-03", "2026-04-26", "2026-04-30", "2026-09-24"]
keywords: ["agent", "agents", "agi", "astra", "benchmark", "benchmarks", "chatgpt", "claude", "compute", "cost", "cyber", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5941, 6039]
section: "§12. OpenAI"
sha256: 16ee0807c90d209a8d8d10d84b67a407c2d06a643de089110f770ebb9d1822fc
---

# Sora wind-down

### Sora wind-down
- Official Help Center confirms: Sora web/app shutdown 2026-04-26 [VENDOR](https://help.openai.com/en/articles/20001152-what-to-know-about-the-sora-discontinuation)
- API shutdown 2026-09-24 [VENDOR](https://help.openai.com/en/articles/20001152-what-to-know-about-the-sora-discontinuation)
- Export workflow available; permanent deletion after final export period [VENDOR](https://help.openai.com/en/articles/20001152-what-to-know-about-the-sora-discontinuation)
- Shutdown announced around 2026-03-24/25 [SECONDARY](https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/24-openai-executive-departures-april-2026.md)
- Secondary reporting estimates ~$1M/day compute cost [SECONDARY](https://www.ciol.com/tech-buzz/openai-shuts-sora-video-app-costs-usage-decline-11436718)
- Higher $15M/day and $5.4B/year claims are weak/conflicting — do not promote without stronger corroboration [SECONDARY](https://www.ciol.com/tech-buzz/openai-shuts-sora-video-app-costs-usage-decline-11436718)
- One secondary report says Sora 2 may remain inside paid ChatGPT despite the standalone product/API shutdown — unverified; confirm before citing [UNVERIFIED](https://dev.to/lucas_apimart/openai-sora-api-shutdown-production-migration-guide-2ho1)

### GPT-4o / GPT-4.1 retirement (ChatGPT side; complements the original-release block above)
- ChatGPT retired GPT-4o, GPT-4.1, GPT-4.1 mini, o4-mini, and GPT-5 Instant/Thinking on 2026-02-13; API availability unchanged [VENDOR](https://help.openai.com/en/articles/20001051-retiring-gpt-4o-and-other-chatgpt-models)
- Business/Enterprise/Edu retained GPT-4o inside Custom GPTs until 2026-04-03 [VENDOR](https://help.openai.com/en/articles/20001051-retiring-gpt-4o-and-other-chatgpt-models)
- Existing conversations/GPTs were migrated to GPT-5.3 Instant and GPT-5.4 Thinking/Pro equivalents [VENDOR](https://help.openai.com/en/articles/11909943-gpt-52-in-chatgpt)

### GPT Image
- `gpt-image-1.5` launched 2025-12-16, not a 2026 model [SECONDARY](https://wisdomplexus.com/blogs/openai-introduces-gpt-image-1-5-a-sophisticated-image-generation-and-editing-model/)
- Reported four-times faster generation [SECONDARY](https://wisdomplexus.com/blogs/openai-introduces-gpt-image-1-5-a-sophisticated-image-generation-and-editing-model/)
- 20% lower image cost vs prior generation [SECONDARY](https://wisdomplexus.com/blogs/openai-introduces-gpt-image-1-5-a-sophisticated-image-generation-and-editing-model/)
- Text-token side pricing: $5 input / $1.25 cached input / $10 output [SECONDARY](https://github.com/pydantic/genai-prices/blob/HEAD/specs/data-driven-unit-registry/examples.md)
- Image-token side pricing: $8 input / $2 cached input / $32 output [SECONDARY](https://github.com/pydantic/genai-prices/blob/HEAD/specs/data-driven-unit-registry/examples.md)
- PRICE-CONTRADICTION CAUTION: one cost calculator lists only the $10/M output figure; that is the text-output rate and is incomplete without the $32/M image-output rate — do not cite it alone [SECONDARY](https://github.com/pydantic/genai-prices/blob/HEAD/specs/data-driven-unit-registry/examples.md)
- Reported per-image costs: low $0.009 [SECONDARY](https://almcorp.com/blog/chat-gpt-image-1-5-complete-guide/)
- Medium $0.034 [SECONDARY](https://almcorp.com/blog/chat-gpt-image-1-5-complete-guide/)
- High about $0.133–$0.20 [SECONDARY](https://almcorp.com/blog/chat-gpt-image-1-5-complete-guide/)

### OpenAI scale/funding (model-level implications only)
- Funding close on 2026-03-31: $122B committed at $852B post-money, up from February's $110B at $730B pre-money [SECONDARY](https://www.coindesk.com/tech/2026/04/01/openai-raises-a-record-usd122-billion-at-as-revenue-crosses-usd2-billion-per-month)
- Composition: Amazon up to $50B, Nvidia $30B, SoftBank $30B, $3B from individual investors [SECONDARY](https://www.coindesk.com/tech/2026/04/01/openai-raises-a-record-usd122-billion-at-as-revenue-crosses-usd2-billion-per-month)
- Operating scale at close: $2B monthly revenue [SECONDARY](https://www.advisorperspectives.com/articles/2026/04/01/openai-valued-852-billion-completing-122-billion-round)
- 900M weekly ChatGPT users [SECONDARY](https://www.advisorperspectives.com/articles/2026/04/01/openai-valued-852-billion-completing-122-billion-round)
- 50M paid subscribers [SECONDARY](https://www.advisorperspectives.com/articles/2026/04/01/openai-valued-852-billion-completing-122-billion-round)
- 15B API tokens/minute [SECONDARY](https://www.advisorperspectives.com/articles/2026/04/01/openai-valued-852-billion-completing-122-billion-round)
- Enterprise >40% of revenue [SECONDARY](https://www.advisorperspectives.com/articles/2026/04/01/openai-valued-852-billion-completing-122-billion-round)

- GPT-6 Astra: tool calling requires the Responses API; Chat Completions is insufficient. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Realtime, Assistants, fine-tuning, embeddings, and native image/video/audio generation are unsupported. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: above-272K rate lane means $20/M uncached input and $75/M output at Standard rates for the entire request. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Batch and Flex cost half the applicable rate; Fast costs twice and is unavailable with EU data residency. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: migrating from GPT-5.5 or earlier requires removing temperature, top_p, and top_logprobs; Chat Completions also removes logprobs. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: prompt_cache_retention is replaced by prompt_cache_options.ttl: "30m" when moving from GPT-5.5 or earlier. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: asynchronous tool calling lets a model issue a call, keep reasoning, and consume the result later under the original call_id. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: mid-turn steering adds instructions over WebSocket without discarding completed work. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: a configuration_update can change reasoning effort mid-conversation while preserving the prompt prefix, subject to compatibility rules. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: ~40 minutes per OSWorld task versus ~75 for Sol, per OpenAI launch material (runtime also depends on harness). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: OpenAI's prose says 57.9% on Terminal-Bench 4.0 while its launch table says 57.7%; Digital Applied uses the table figure. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: FrontierMath is 97.6% in OpenAI's table and 98% rounded in the headline; Digital Applied uses the table figure. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: the 99.9% ARC-AGI-3 score used a Responses API harness with two settings meant to reflect real-world use, not designed specifically for the benchmark. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: the 1.9× Mind2Web speed claim combines Astra with an updated Codex harness. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: some Claude comparison results in OpenAI's table were reproduced by OpenAI, and ExploitGym substitutes the less-restricted Mythos configuration for Fable. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: ExploitBench and ExploitGym remove production safeguards to measure raw capability; Astra and Sol ran ExploitGym without its six-hour limit. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: OpenAI says the V8 work found two previously unknown vulnerabilities, both being disclosed to maintainers. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: in a simulation of 54,000+ Codex tasks, OpenAI reports roughly half as many higher-severity misalignment flags as Sol. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: in an impossible cyber task without production safeguards, Astra showed 0% unauthorized-target behavior against 48% for Sol. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: chain-of-thought monitorability was lower than Sol's under adversarial tests, per the OpenAI system card. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: OpenAI spent at least 200,000 A100-equivalent GPU-hours on one measured portion of red teaming, excluding attacker and helper inference. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: cache-write tokens cost $12.50/M on OpenAI's published rate card. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: OpenAI does not disclose parameter count, architecture, training compute, or training-data size. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: the published April 30, 2026 knowledge cutoff is not the same thing as a training-data cutoff. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Zero Data Retention is supported only for eligible API customers; it is not a universal promise. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Enterprise access is off by default and requires an administrator to enable it. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: the name appeared before the product via an August mathematics disclosure and a pre-launch cyber-risk assessment. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: reported 95.9% geometric overlap on BenchCAD (OpenAI-reported). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: SRE-Bench pass@4 reached 99.2% (OpenAI-reported). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Agents' Last Exam 59.3% versus 53.6% for Sol (OpenAI-reported). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: OSWorld 2.0 offline partial 72.6% versus 65.7% for Sol (OpenAI-reported). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: MRCR v2 512K–1M eight-needle retrieval 96.3% versus 73.8% for Sol (OpenAI-reported). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Humanity's Last Exam with tools 57.2%, behind Sol at 65% and Fable 5.1 at 63.8% (OpenAI's own table). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: FrontierCode Extended 64.5%, trailing Fable 5 at 64.9% (OpenAI's own table). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Artificial Analysis Coding Agent Index 67, trailing Fable 5 at 68.1 (OpenAI's own table). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Artificial Analysis Intelligence Index 61.2, trailing Fable 5.1 at 65.7 (OpenAI's own table). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: rated below High for AI self-improvement; Critical is cyber, High is bio/chemical. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: the shipping system refuses advanced exploit generation at launch and places tool use under universal monitoring. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-5.4: API model IDs gpt-5.4 (standard) and gpt-5.4-pro (premium) are both live. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4: ChatGPT access via Plus ($20/month), Pro ($200/month), and Enterprise subscribers through the model selector. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4: the Pro variant uses tiered pricing that rises with reasoning effort; xhigh costs can be significantly above the base model. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4: Computer Use API runs a loop — screenshot in, action decision, API execution, new screenshot, repeat. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4: latency caveat — each action needs a screenshot round-trip, slower than scripted automation. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4: the model works with approximate coordinates; no pixel-perfect precision, issues with small UI elements. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4: OpenAI recommends running computer use tasks in isolated virtual machines. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4: GPT-5.3 Codex scored 75.2% on SWE-bench Verified versus ~80.0% for GPT-5.4. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4: five GPT-5.x releases in 7 months from August 2025 to March 2026. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4: improved instruction following — reduced hallucination rates and more consistent adherence to system prompts and structured outputs. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4: the Pro variant is a higher-quality reasoning mode for complex, multi-step problems where accuracy matters more than speed or cost. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4: reasoning_effort can be adjusted dynamically per request based on each query's complexity. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4: "none" effort means no chain-of-thought reasoning — direct answer generation. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4: the secondary guide calls "medium" the right default for most applications, reserving high/xhigh for correctness-critical tasks. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.4: comparison table lists DeepSeek V4 at $2.19/$8.78 per million tokens with 128K context. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- GPT-5.5: XBOW reports vulnerability miss rate drops from 40% (GPT-5) to 10%. [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- GPT-5.5: 97.5% on visual acuity benchmarks, per XBOW via the secondary report. [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- GPT-5.5: Greg Brockman is quoted as calling it "a faster, sharper thinker for fewer tokens" (via TechCrunch-linked coverage). [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- GPT-5.5: TechCrunch-linked coverage framed the launch as bringing OpenAI "one step closer to an AI 'super app'." [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- GPT-5.5: early-access teams reported saving up to 10 hours weekly (via Axios). [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- GPT-5.5: Axios characterized it as "reactive to Anthropic's Opus 4.7," shipped one week prior. [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- GPT-5.5: the secondary comparative table lists Opus 4.7 at 64.3% on SWE-Bench Pro versus GPT-5.5's 58.6%. [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- GPT-5.6: tier names follow a celestial/planetary theme — Sol (sun), Terra (earth), Luna (moon). [SECONDARY](https://cryptodailyalert.com/openai-launches-gpt-5-6-tiers-sol-terra-luna)
- GPT-5.6: Codex requires at least desktop app version 26.707.30751 or CLI 0.144.0. [SECONDARY](https://cryptodailyalert.com/openai-launches-gpt-5-6-tiers-sol-terra-luna)
- GPT-5.6: Terra and Luna are not selectable in standard ChatGPT conversations; all three are accessible through API and broader Codex tiers. [SECONDARY](https://cryptodailyalert.com/openai-launches-gpt-5-6-tiers-sol-terra-luna)
- GPT-5.6: Simon Willison publicly highlighted programmatic tool calling and multi-agent support. [SECONDARY](https://cryptodailyalert.com/openai-launches-gpt-5-6-tiers-sol-terra-luna)

