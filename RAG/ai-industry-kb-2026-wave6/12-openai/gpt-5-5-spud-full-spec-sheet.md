---
id: ai-industry-kb-2026-wave6/12-openai/gpt-5-5-spud-full-spec-sheet
title: "GPT-5.5 \"Spud\" — full spec sheet"
domain: openai
role: deep-dive
task: actor-profile
actors: ["AWS", "Cerebras", "Google", "Microsoft", "OpenAI"]
dates: ["2026-02-16", "2026-04-23", "2026-04-24", "2026-04-30", "2026-06", "2026-07-02", "2026-07-09", "2026-07-10", "2026-08-11", "2026-09-01"]
keywords: ["agent", "agentic", "agi", "astra", "attribution", "aws", "bedrock", "benchmark", "benchmarks", "chatgpt", "claude", "context window"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5865, 5940]
section: "§12. OpenAI"
delta_of: ai-industry-kb-2026
sha256: 7193229711a1606b04457e825c29de463fa55dd11396da7af3bc4faf08a8b6ce
---

# GPT-5.5 "Spud" — full spec sheet

### GPT-5.5 "Spud" — full spec sheet
- GPT-5.5 reportedly the first fully retrained OpenAI base since GPT-4.5; 5.1–5.4 releases characterized as post-training updates [SECONDARY](https://github.com/mattrobenolt/pi-configs/blob/HEAD/skills/model-guide/research/gpt-5.5.md)
- A directly opened aggregator page corroborates the "first fully retrained base since GPT-4.5" framing, citing OpenAI's own launch page [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- Internal codename "Spud" [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- Modalities CONFLICT: reported official spec is text+image input, text output; secondary "natively omnimodal" claims adding audio/video contradict it and appear unreliable [UNVERIFIED](https://earlyterms.com/term/gpt-5-5)
- Context 1M tokens [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- Max output 128K tokens [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- ChatGPT/Codex release 2026-04-23; API release 2026-04-24 [SECONDARY](https://blog.buildfastwithai.com/gpt-5-5-review-2026)
- Released simultaneously to ChatGPT Plus, Pro, Business, Enterprise and Codex on launch day [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- API pricing: $5/M input, $30/M output; GPT-5.5 Pro $30/$180 (from a directly opened aggregator page) [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- Four-tier commercial structure reported: standard, Pro, half-rate batch, 2.5× priority [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- Vendor claim: matches GPT-5.4's per-token latency while delivering higher intelligence with fewer tokens [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- Preparedness Framework biosecurity risk rated 'High' for GPT-5.5 (per the same aggregator page) [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- GPT-5.5-Cyber was in use with Codex Security in "Patch the Planet" by June 2026 [SECONDARY](https://github.com/mattrobenolt/pi-configs/blob/HEAD/skills/model-guide/research/gpt-5.5.md)

### GPT-5.6 Sol/Terra/Luna — full spec sheet
- Three-tier family: Sol (large), Terra (medium), Luna (small) [SECONDARY](https://cryptodailyalert.com/openai-launches-gpt-5-6-tiers-sol-terra-luna) [SECONDARY](https://theplanettools.ai/blog/gpt-5-6-sol-terra-luna-not-gpt-6-explained-2026)
- LAUNCH-DATE CONTRADICTION: one source reports launch 2026-07-02; the page above cites OpenAI's launch page as dated 2026-07-10; GitHub confirmed Copilot availability 2026-07-09 — do not cite a single date without the caveat [SECONDARY](https://www.testingcatalog.com/openai-launches-gpt-5-6-sol-terra-and-luna-on-apps-and-api/) [SECONDARY](https://cryptodailyalert.com/openai-launches-gpt-5-6-tiers-sol-terra-luna)
- Launch-day pricing, Sol: $5 input / $0.50 cache read / $6.25 cache write / $30 output [SECONDARY](https://cryptodailyalert.com/openai-launches-gpt-5-6-tiers-sol-terra-luna) [SECONDARY](https://theplanettools.ai/blog/gpt-5-6-sol-terra-luna-not-gpt-6-explained-2026)
- Launch-day pricing, Terra: $2.50 / $0.25 / $3.125 / $15 [SECONDARY](https://cryptodailyalert.com/openai-launches-gpt-5-6-tiers-sol-terra-luna)
- Launch-day pricing, Luna: $1 / $0.10 / $1.25 / $6 [SECONDARY](https://cryptodailyalert.com/openai-launches-gpt-5-6-tiers-sol-terra-luna)
- Cache write priced at 1.25× uncached input [SECONDARY](https://kie.ai/blog/gpt-5-6-sol-terra-luna-deep-dive) [SECONDARY](https://github.com/octopollux/machine-violet/commit/2f42165dcbc6e1f461b6445eabef77c91b8024aa)
- Cache read carries a 90% discount [SECONDARY](https://kie.ai/blog/gpt-5-6-sol-terra-luna-deep-dive)
- Explicit cache breakpoints and 30-minute minimum cache life documented [SECONDARY](https://kie.ai/blog/gpt-5-6-sol-terra-luna-deep-dive) [SECONDARY](https://github.com/octopollux/machine-violet/commit/2f42165dcbc6e1f461b6445eabef77c91b8024aa)
- Post-cut rates (Aug–Sep 2026): Sol $4/$20, Terra $2/$12, Luna $0.20/$1.20 — differ materially from launch; date every rate [SECONDARY](http://datafloq.com/gpt-5-6-pricing-explained-why-sol-terra-and-luna-turn-ai-buying-into-a-routing-problem/)
- All three models reportedly ~1.05M context, 128K max output, knowledge cutoff 2026-02-16 [SECONDARY](https://www.testingcatalog.com/openai-launches-gpt-5-6-sol-terra-and-luna-on-apps-and-api/)
- Per Simon Willison's notes (via the launch coverage): a 1-million-token context window with 128,000 max output [SECONDARY](https://cryptodailyalert.com/openai-launches-gpt-5-6-tiers-sol-terra-luna)
- Programmatic Tool Calling executing in-memory JavaScript [SECONDARY](https://kie.ai/blog/gpt-5-6-sol-terra-luna-deep-dive)
- Multi-agent Responses API (beta) [SECONDARY](https://kie.ai/blog/gpt-5-6-sol-terra-luna-deep-dive)
- Six reasoning levels across the three models (per Simon Willison's notes) [SECONDARY](https://cryptodailyalert.com/openai-launches-gpt-5-6-tiers-sol-terra-luna)
- Reported Max/Ultra modes [SECONDARY](https://kie.ai/blog/gpt-5-6-sol-terra-luna-deep-dive)
- GitHub Copilot integration began 2026-07-09: Sol for Copilot Pro+/Max/Business/Enterprise; Terra and Luna across Pro through Enterprise SKUs [SECONDARY](https://cryptodailyalert.com/openai-launches-gpt-5-6-tiers-sol-terra-luna)
- Business/Enterprise admin opt-in required for Copilot (policy off by default) [SECONDARY](https://cryptodailyalert.com/openai-launches-gpt-5-6-tiers-sol-terra-luna)
- Benchmark attribution conflict: vendor Coding Agent Index score 80.0 for Sol vs independent/versioned scores ~59 and 66.6 — not the same benchmark, not interchangeable [SECONDARY](http://datafloq.com/gpt-5-6-pricing-explained-why-sol-terra-and-luna-turn-ai-buying-into-a-routing-problem/)
- One retirement-focused source says GPT-5.6 Sol Ultrafast runs on Cerebras at up to 750 output tokens/sec — single-source [SECONDARY](https://aiidelist.com/blog/gpt-5-3-codex-spark-retirement)

### GPT-5.6-Cyber
- Built on GPT-5.6 Sol [SECONDARY](https://www.thebrief.news/en/standard/article/10441/openai-splits-daybreak-cybersecurity-program-into-blue-and-red-tiers-launches-gpt-56-cyber) [SECONDARY](https://www.csoonline.com/article/4207896/openai-launches-gpt-5-6-cyber-as-ai-narrows-vulnerability-response-window.html)
- Restricted to Daybreak Red [SECONDARY](https://www.thebrief.news/en/standard/article/10441/openai-splits-daybreak-cybersecurity-program-into-blue-and-red-tiers-launches-gpt-56-cyber)
- Daybreak Blue uses guarded general models while Red grants specialized exploit/vulnerability models [SECONDARY](https://www.thebrief.news/en/standard/article/10441/openai-splits-daybreak-cybersecurity-program-into-blue-and-red-tiers-launches-gpt-56-cyber) [SECONDARY](https://www.csoonline.com/article/4207896/openai-launches-gpt-5-6-cyber-as-ai-narrows-vulnerability-response-window.html)
- Vendor-reported Advanced Cybersecurity Completion Rate: 95% for 5.6-Cyber vs 57.3% for 5.5-Cyber [SECONDARY](https://pondero.ai/news/2026-08-11-openai-daybreak-gpt56-cyber/)
- Same vendor metric: 1.5% for standard Sol, ~2% for Daybreak Blue [SECONDARY](https://pondero.ai/news/2026-08-11-openai-daybreak-gpt56-cyber/)
- Found two V8 flaws reportedly chainable to a heap-sandbox escape; coordinated disclosure to Google [SECONDARY](https://www.business-standard.com/technology/tech-news/openai-daybreak-gpt-5-6-cyber-ai-cybersecurity-breaches-126081100597_1.html)
- OpenAI's preparedness rating for this model family is "High," not "Critical" [SECONDARY](https://www.freepressjournal.in/tech/openai-expands-daybreak-cybersecurity-programme-launches-more-permissive-gpt-56-cyber-model)
- Identity verification, monitoring, and mandatory hardware security keys required from 2026-09-01 [SECONDARY](https://www.csoonline.com/article/4207896/openai-launches-gpt-5-6-cyber-as-ai-narrows-vulnerability-response-window.html)

### GPT-6 Astra — full spec sheet
- GPT-6 Astra: proprietary reasoning/computer-use model; parameters undisclosed [SECONDARY](https://www.userightai.com/models/gpt-6-astra) [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- Model ID `gpt-6-astra` [SECONDARY](https://www.userightai.com/models/gpt-6-astra)
- Context 1,050,000 total: 922K input + 128K max output [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- Knowledge cutoff 2026-04-30 [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- Inputs: text + images; output: text only [SECONDARY](https://www.userightai.com/models/gpt-6-astra)
- Pricing: $10/M input [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- Cache read: $1/M [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- Output: $50/M [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- Five-minute cache write: $12.50/M [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- Above 272K input: the entire request is reportedly repriced at $20 input / $75 output [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- Batch API: half-price [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- Fast tier: $20/M input, $100/M output [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- ChatGPT rollout covered Plus, Pro, Business, Enterprise; Enterprise access off by default [SECONDARY](https://witho2.com/news/gpt-6-astra-launch-openai-s-computer-use-ai-pricing-and-who-gets-it)
- Integrations reported: Azure, AWS/Bedrock, GitHub Copilot [SECONDARY](https://davidandgoliath.ai/daily-ai-briefing/openai-gpt-6-astra-enterprise-launch)
- Preparedness rating "Critical" for cybersecurity; restricted advanced capabilities gated through Daybreak [SECONDARY](https://digikestra.com/blog/gpt-6-astra-agi)
- ARC-AGI-3 CONTRADICTION: vendor harness reports 99.9%, but the standard ARC-AGI-3 harness reportedly scores 62.7% — different harnesses, both vendor-derived, not interchangeable [SECONDARY](https://www.magicboat.net/mobile/en/blog/gpt-6-astra-vs-claude-fable-5-1-which-is-better-for-creators)
- Artificial Analysis Intelligence Index figures for Astra conflict across versions (v4.1.1 vs v4.3, percentages vs later revisions); keep each version's figure separate — never merge index versions [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)

### ChatGPT Work
- ChatGPT Work runs GPT-5.6 with the Codex agent across files, apps, browser, local desktop, and connected services [SECONDARY](https://www.macrumors.com/2026/07/09/openai-chatgpt-work/) [SECONDARY](https://siliconangle.com/2026/07/09/openai-debuts-chatgpt-work-agentic-tool-automating-business-workflows/)
- Web/mobile rollout began for Pro/Enterprise/Edu on 2026-07-09 [SECONDARY](https://www.macrumors.com/2026/07/09/openai-chatgpt-work/)
- Plus/Business access came later [SECONDARY](https://www.macrumors.com/2026/07/09/openai-chatgpt-work/)
- Unified Mac/Windows desktop app launched globally with Chat, Work, and Codex surfaces [SECONDARY](https://www.digitalapplied.com/blog/chatgpt-work-openai-agent-launch-2026)
- The former desktop app was renamed ChatGPT Classic [SECONDARY](https://www.digitalapplied.com/blog/chatgpt-work-openai-agent-launch-2026)
- Atlas browser sunset began with functionality folded into Work [SECONDARY](http://ppc.land/openai-kills-atlas-browser-folds-it-into-new-chatgpt-work-agent/)
- Work can schedule background work [SECONDARY](https://siliconangle.com/2026/07/09/openai-debuts-chatgpt-work-agentic-tool-automating-business-workflows/)
- Output types: spreadsheets, slides, documents, dashboards, websites, and apps [SECONDARY](https://siliconangle.com/2026/07/09/openai-debuts-chatgpt-work-agentic-tool-automating-business-workflows/)
- Codex reportedly had over 5M weekly users, with over 1M outside software development [SECONDARY](https://www.digitalapplied.com/blog/chatgpt-work-openai-agent-launch-2026)

