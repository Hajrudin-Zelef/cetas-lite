---
id: ai-industry-kb-2026-wave6/16-cohere/part-10
title: "§16. Cohere (part 10)"
domain: cohere
role: deep-dive
task: actor-profile
actors: ["AWS", "Cohere", "CoreWeave", "Hugging Face", "Mistral", "Nvidia", "OpenAI"]
dates: ["2025-03", "2025-03-17", "2025-04-15", "2025-08-21", "2026-03", "2026-04-24", "2026-05", "2026-05-20", "2026-05-21", "2026-05-24", "2026-07-30"]
keywords: ["cohere", "agent", "agentic", "apache", "arr", "attention", "bedrock", "benchmarks", "blackwell", "claude", "context window", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8191, 8227]
section: "§16. Cohere"
delta_of: ai-industry-kb-2026
sha256: 86f396fcc04f2a7c6b208b22daa48fd51a6d491eab1efcdeaa6caac58fc2e75e
---

# §16. Cohere (part 10)

- S1 — AwesomeAgents, "Cohere Command A+" spec sheet (218B/25B, 128K/64K, Apache 2.0, May 20, 2026): https://awesomeagents.ai/models/cohere-command-a-plus/
- S2 — TokenCost, "Cohere Command A+ Pricing: Apache 2.0 on Two H100s" ($2.50/$10 mirrored rates, no Cohere-published row): https://tokencost.app/blog/cohere-command-a-plus-pricing
- S3 — FutureTweets, "Mistral Large 3 vs Nova 2 Sonic vs Command A+ [2026]" (context comparison, pricing caveats): https://futuretweets.com/mistral-large-3-vs-nova-2-vs-command-a-plus-2026/
- S4 — VentureBeat, "Cohere cracks lossless quantization and native citations with first full Apache 2.0 licensed open model Command A+": https://venturebeat.com/technology/cohere-cracks-lossless-quantization-and-native-citations-with-first-full-apache-2-0-licensed-open-model-command-a
- S5 — CloudRadix, "Cohere Command A+: Native Citations Reset AI Procurement" (May 21, 2026, W4A4 on 2×H100 or 1×B200): https://cloudradix.com/blog/cohere-command-a-plus-native-citations-open-weight-ai-procurement-aeo-2026/
- S6 — AI Insight Lab, "The Open-Source Sovereign AI Decision" deployment memo (May 24, 2026): https://aiinsightlab.cloud/api/memos/cohere-command-a-plus/pdf
- S7 — FutureTweets FAQ extract on Command A+ pricing status (no published per-token price, free up to rate limits): https://futuretweets.com/mistral-large-3-vs-nova-2-vs-command-a-plus-2026/
- S8 — GitHub issue a5c-ai/babysitter#1382, "Track Cohere Rerank v4 and Transcribe model coverage" (model IDs rerank-v4.0-pro/fast, cohere-transcribe-03-2026, command-a-plus-05-2026 from docs.cohere.com): https://github.com/a5c-ai/babysitter/issues/1382
- S9 — Wealth Engine, "Cohere $20 Billion Valuation: Aleph Alpha Merger Signed at 83x ARR" (Sept 18, 2026; $240M ARR, 70% margins, 85% private deployments): https://wealthengine.blog/2026/09/18/cohere-20-billion-valuation/
- S10 — pedro-bright/the-ledger, "32-cohere-aleph-alpha-merger.md" (Sept 16 definitive agreement, 90/10 split, €500M Schwarz Digits, Gomez/Andrulis): https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/32-cohere-aleph-alpha-merger.md
- S11 — CryptoBriefing, "Cohere in advanced talks to raise up to $3B as enterprise AI race heats up" (Sept 2025 $7B valuation, $240M ARR doubling from $100M): https://cryptobriefing.com/cohere-advanced-talks-raise-3-billion/
- S12 — PitchBook, Cohere company profile (Series E $2.5B in progress, April 24, 2026): https://pitchbook.com/profiles/company/443089-09
- S13 — MENAFN/GlobeNewsWire, "Cohere And Carahsoft Partner To Bring Secure, Sovereign AI Deployment Solutions To The Public Sector" (July 30, 2026; SEWP V, ITES-SW2, NASPO, TIPS, OMNIA): https://menafn.com/1111467348/Cohere-And-Carahsoft-Partner-To-Bring-Secure-Sovereign-AI-Deployment-Solutions-To-The-Public-Sector
- S50 — NVIDIA blog, "Thousands of NVIDIA Grace Blackwell GPUs Now Live at CoreWeave" (Cohere GB200 NVL72: up to 3x training performance on 100B-param models): https://blogs.nvidia.com/blog/coreweave-grace-blackwell-gb200-nvl72/
- S53 — Hugging Face model card, CohereLabs/c4ai-command-a-03-2025 (111B, 256K configured 128K on HF, CC-BY-NC + Acceptable Use Policy): https://huggingface.co/CohereLabs/c4ai-command-a-03-2025
- S54 — Community AI-model tracker, "Cohere Command A Family Release - March 2025" (Command A Vision Jul 2025, 112B; family release dates): https://github.com/sbley/claude-code-agent-test/issues/62
- S13b — ExecutiveBiz, "Carahsoft to Distribute Cohere AI Models to Government" (FedRAMP High May 2026 via Second Front Systems, air-gapped/on-prem): https://www.executivebiz.com/articles/cohere-carahsoft-public-sector-ai-distribution-partnership
- S13c — Carahsoft official news, "Cohere and Carahsoft Partner to Bring Secure, Sovereign AI Deployment Solutions to the Public Sector" (vendor-side corroboration, July 30, 2026): https://www.carahsoft.com/news/cohere-and-carahsoft-partner-to-bring-secure-sovereign-ai-deployment-solutions-to-the-public-sector-2026
- S14 — Unite.AI, "OpenText and Cohere Partner on Agentic AI for Regulated Sectors" (Sept 16, 2026, ALL IN AI): https://www.unite.ai/opentext-and-cohere-partner-on-agentic-ai-for-regulated-sectors/
- S15 — newswire.ca, "OpenText, Cohere Partner to Combine Trusted Data with Agentic AI" (Sept 17, 2026; Antoun quote; deployment options): https://www.newswire.ca/news-releases/opentext-cohere-partner-to-combine-trusted-data-with-agentic-ai-897974014.html
- S16 — ts2.tech, "Cohere Confirms Series E Interest as Reported Talks Target a $20 Billion Valuation" (no revenue disclosed, no audited figures): https://ts2.tech/en/cohere-confirms-series-e-interest-as-reported-talks-target-a-20-billion-valuation/
- S17 — ecorpit.com, "Cohere Command A vs Command R+ (2026): migrate?" (111B, 256K, 150% throughput, $2.50/$10, last updated July 30, 2026): https://ecorpit.com/cohere-command-a-vs-command-r-plus-migration-2026/
- S17b — runapi-builder/awesome-ai-tools-curated-llm-apis (Command A 256K / $2.50/$10, March 2026 snapshot): https://github.com/runapi-builder/awesome-ai-tools-curated-llm-apis
- S18 — metacto.com, "Cohere API Pricing 2026" (Command A/R+ $2.50/$10, legacy 04-2024 $3/$15, legacy customer tiers; R7B $0.0375/$0.15, no prompt caching, trial key limits): https://www.metacto.com/blogs/cohere-pricing-explained-a-deep-dive-into-integration-development-costs
- S19 — Cohere docs changelog, "2025-08-21 Command A Reasoning" (vendor): https://docs.cohere.com/changelog/2025-08-21-command-a-reasoning
- S20 — Hugging Face, CohereLabs/command-a-reasoning-08-2025: https://huggingface.co/CohereLabs/command-a-reasoning-08-2025
- S21 — VentureBeat, "Cohere's Rerank 4 quadruples the context window" (2025; 32K context): https://venturebeat.com/ai/coheres-rerank-4-quadruples-the-context-window-to-cut-agent-errors-and-boost
- S22 — Vercel AI Gateway, "Embed v4.0 API & Pricing" (April 15, 2025; 65.2 MTEB; dims; $0.12/M): https://vercel.com/ai-gateway/models/embed-v4.0
- S23 — app.ailog.fr, "Cohere Embed v4: The First Production Multimodal Embedding" (pricing $0.12/$0.47, cloud availability): https://app.ailog.fr/en/blog/news/cohere-embed-v4-multimodal
- S24 — eesel.ai, "Cohere AI pricing in 2026: A complete guide to real costs" (Bedrock pared back, provisioned-throughput pricing): https://www.eesel.ai/blog/cohere-ai-pricing
- S26 — hal-xp/hal-o knowledge mirror, Cohere provider notes (v2 OpenAI-compatible API, pricing, free-tier limits): https://github.com/hal-xp/hal-o
- S27 — sbley/claude-code-agent-test#423, "Cohere Command A+ - First Fully Apache 2.0 Open-Source Flagship Model" (218B/25B, May 20–21, 2026): https://github.com/sbley/claude-code-agent-test/issues/423
- S28 — Dataconomy, "Cohere's 111B-parameter AI Model Can Run On Just Two GPUs" (March 17, 2025; attention architecture, 50% cost-reduction claim): https://dataconomy.com/2025/03/17/cohere-111b-parameter-ai-model-can-run-on-just-two-gpus/
- S29 — VentureBeat, "Cohere targets global enterprises with new highly multilingual Command A model requiring only 2 GPUs" (March 2025; 73 tok/s, ADI2 24.7, $2.50/$10 pricing): https://venturebeat.com/ai/cohere-targets-global-enterprises-with-new-highly-multilingual-command-a-model-requiring-only-2-gpus
- S30 — Hugging Face, CohereLabs/c4ai-command-a-03-2025 model card (111B, CC-BY-NC, 256K/128K config note; vendor): https://huggingface.co/CohereLabs/c4ai-command-a-03-2025
- S30b — Cohere docs, Command A Reasoning (reasoning toggle parameter; vendor): https://docs.cohere.com/docs/command-a-reasoning
- S31 — VentureBeat, "Don't sleep on Cohere: Command A Reasoning" (token budget, reasoning toggle, vendor benchmarks vs R1 0528): https://venturebeat.com/ai/dont-sleep-on-cohere-command-a-reasoning-its-first-reasoning-model-is-built-for-enterprise-customer-service-and-more
