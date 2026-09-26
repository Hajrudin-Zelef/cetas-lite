---
id: ai-industry-kb-2026-wave6/12-openai/gpt-5-5-spud-full-spec-sheet
title: "GPT-5.5 \"Spud\" — full spec sheet"
domain: openai
role: deep-dive
task: actor-profile
actors: ["Cerebras", "Google", "Microsoft", "OpenAI"]
dates: ["2026-02-16", "2026-04-23", "2026-04-24", "2026-06", "2026-07-02", "2026-07-09", "2026-07-10", "2026-08-11", "2026-09-01"]
keywords: ["agent", "attribution", "benchmark", "chatgpt", "context window", "copilot", "cyber", "cybersecurity", "disclosure", "exploit", "gpt-5.6", "gpt-6"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5865, 5910]
section: "§12. OpenAI"
delta_of: ai-industry-kb-2026
sha256: 3aae522c8fc17aab588d2cbf4ab5c162fe7d30da4964b48178c88ce5df85962b
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

