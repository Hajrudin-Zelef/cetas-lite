---
id: labs-hyperscalers-2026/00-labs-hyperscalers/part-12
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026) (part 12)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["Anthropic", "ExploitGym", "Microsoft", "OpenAI", "United States"]
dates: ["2026-04", "2026-04-16", "2026-04-22", "2026-07"]
keywords: ["agent", "agentic", "agents", "benchmarks", "chatgpt", "claude", "copilot", "gpt-5.6", "incident", "latency", "pricing", "protein"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [324, 340]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 04e56e9766dfd70be78f48a441cb03ac886be2a031b63ec57e08c6405029901b
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026) (part 12)

#### GPT-5.5 — April 22, 2026
- **GPT-5.5** released **Wednesday 22 April 2026** — agentic reasoning, designed for multi-part autonomous tasks across software tools. Maintains GPT-5.4 per-token latency with "significantly higher intelligence" and fewer tokens per task. [secondary] https://world.infonasional.com/openai-launches-gpt-5-5-agentic
- Benchmarks (vendor-reported): Terminal-Bench 2.0 82.7% (vs 75.1%); GDPval 84.9%; OSWorld-Verified 78.7%; BrowseComp 84.4% (5.5 Pro 90.1%); FrontierMath Tier 1–3 51.7%, Tier 4 35.4%; CyberGym 81.8%; Expert-SWE (internal) 73.1%. [secondary] https://world.infonasional.com/openai-launches-gpt-5-5-agentic
- API pricing: GPT-5.5 **$5.00 / 1M in, $30.00 / 1M out**; GPT-5.5 Pro $30.00/$180.00. [secondary] https://devtk.ai/en/blog/openai-api-pricing-guide-2026/
- Available to Plus/Pro/Business/Enterprise tiers. [secondary] https://www.gradually.ai/en/codex-statistics/
- GPT-5.5 was noted as outranking on ExploitGym with 120 successes (vs Claude Mythos Preview 157) — detail from incident timeline. [secondary] https://github.com/kzinmr/ai-topics/blob/HEAD/wiki/events/openai-huggingface-incident-july-2026.md
- GPT-5.5 is now the preferred model powering Microsoft 365 Copilot (per OpenAI, July 2026) — later GPT-5.6. [vendor-reported via press] http://indianexpress.com/article/technology/artificial-intelligence/openai-gpt-5-6-chatgpt-work-ai-agents-productivity-10779911/

#### GPT-Rosalind — first life-sciences model — April 16, 2026
- **GPT-Rosalind** unveiled **16 April 2026** (TechTarget dated 20 Apr 2026 for Thursday announcement — note the 16 Apr date is the announcement date per Axios/AwesomeAgents) — OpenAI's first purpose-built life-sciences / drug-discovery reasoning model, named after chemist Rosalind Franklin. [independent] https://www.techtarget.com/pharmalifesciences/news/366641922/OpenAI-debuts-AI-model-GPT-Rosalind-to-speed-up-drug-discovery
- Built on OpenAI's newest internal models; multi-step reasoning across chemistry, genomics, protein engineering; enterprise-grade security controls. [vendor-reported via press] https://www.techtarget.com/pharmalifesciences/news/366641922/OpenAI-debuts-AI-model-GPT-Rosalind-to-speed-up-drug-discovery
- Benchmarks (vendor-reported): BixBench **0.751 pass rate**; beats GPT-5.4 on **6 of 11 LABBench2** tasks (strongest on CloningQA); outranked human experts on RNA prediction per one report. [vendor-reported via press] https://awesomeagents.ai/news/openai-gpt-rosalind-life-sciences-model/ ; https://megaoneai.com/spotlight/openai-launches-gpt-rosalind-first-dedicated-life-sciences-ai-model/
- Access: **research preview** for eligible institutions — ChatGPT, Codex and API — initially **qualified US Enterprise customers**. Early partners: **Amgen, Moderna, Thermo Fisher Scientific, Allen Institute, Dyno Therapeutics**. [independent] https://www.techtarget.com/pharmalifesciences/news/366641922/OpenAI-debuts-AI-model-GPT-Rosalind-to-speed-up-drug-discovery ; https://awesomeagents.ai/news/openai-gpt-rosalind-life-sciences-model/
- Companion **Life Sciences research plugin for Codex** released simultaneously — connects to 50+ scientific tools/databases. [independent] https://www.techtarget.com/pharmalifesciences/news/366641922/OpenAI-debuts-AI-model-GPT-Rosalind-to-speed-up-drug-discovery
- Official announcement: https://openai.com/index/introducing-gpt-rosalind/ (referenced, not fetched) — cited via https://github.com/sbley/claude-code-agent-test/issues/363
- Context: launched the day before Kevin Weil's departure (see §2.4); OpenAI for Science's final output before the division was dissolved. [secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/24-openai-executive-departures-april-2026.md

