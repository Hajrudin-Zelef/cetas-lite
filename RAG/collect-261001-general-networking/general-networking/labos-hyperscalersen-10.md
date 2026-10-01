---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-10
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "ExploitGym", "Microsoft", "OpenAI", "United States"]
dates: ["2026-03", "2026-03-05", "2026-03-06", "2026-04", "2026-04-16", "2026-04-22", "2026-07"]
keywords: ["agent", "agentic", "agents", "benchmark", "benchmarks", "chatgpt", "claude", "context window", "copilot", "gpt-5.6", "incident", "latency"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [316, 340]
sha256: 8c50ae6caaa40aa67e38d78a46bd1086d14f5616e99cc5ff05202ee63bac0628
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

#### GPT-5.4 — Thinking and Pro — March 5, 2026
- **GPT-5.4** launched **Thursday 5 March 2026** (TechCrunch dated 2026/03/05), billed as "our most capable and efficient frontier model for professional work," combining coding + reasoning + native computer use in one model. Ships in three variants: standard, **GPT-5.4 Thinking** (reasoning) and **GPT-5.4 Pro** (high performance). [independent] https://techcrunch.com/2026/03/05/openai-launches-gpt-5-4-with-pro-and-thinking-versions/
- API context window **1 million tokens** — by far OpenAI's largest at the time. [independent] https://techcrunch.com/2026/03/05/openai-launches-gpt-5-4-with-pro-and-thinking-versions/
- API pricing (standard): GPT-5.4 **$2.50 / 1M in, $15 / 1M out**; mini $0.75/$4.50; nano $0.20/$1.25. [secondary] https://devtk.ai/en/blog/openai-api-pricing-guide-2026/
- Benchmarks (vendor-reported): GDPval 83.0% (record; matched/exceeded industry professionals in 83% of comparisons vs 70.9% for GPT-5.2); OSWorld-Verified 75.0% and WebArena Verified (records); SWE-Bench Pro 57.7%; BrowseComp 82.7% (Pro 89.3%); Toolathlon 54.6%; factuality: individual claims 33% less likely false vs GPT-5.2, full responses 18% less likely to contain errors; spreadsheet-modeling internal benchmark 87.3% vs 68.4%. Top score on Mercor APEX-Agents (per Mercor CEO Brendan Foody). [vendor-reported via press] https://www.businesstoday.in/technology/news/story/openai-releases-gpt-54-model-with-advanced-reasoning-coding-and-native-computer-use-519373-2026-03-06 ; https://itbrief.co.nz/story/openai-unveils-gpt-5-4-with-advanced-computer-use-tools
- New image input detail levels: "original" up to **10.24M total pixels / 6000px max dimension**; "high" raised to 2.56M pixels / 2048px. [independent] https://itbrief.co.nz/story/openai-unveils-gpt-5-4-with-advanced-computer-use-tools
- "/fast" mode in Codex: up to **1.5x token speed** without quality loss. Experimental Codex skill "Playwright (Interactive)" for visual web/Electron debugging released alongside. ChatGPT for Excel add-in launched with the model. [secondary] https://the-decoder.com/openai-launches-gpt-5-4-thinking-and-pro-combining-coding-reasoning-and-computer-use-in-one-model/

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

