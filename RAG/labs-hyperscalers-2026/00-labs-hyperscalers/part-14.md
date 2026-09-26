---
id: labs-hyperscalers-2026/00-labs-hyperscalers/part-14
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026) (part 14)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "Google", "Meta", "OpenAI", "United States"]
dates: ["2026-04", "2026-07", "2026-08", "2026-09", "2026-09-03"]
keywords: ["agents", "agi", "alignment", "apache", "astra", "aws", "benchmarks", "chatgpt", "claude", "consumer", "context window", "cyber"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [355, 370]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 9b7b1519ccdadb67a463233d7f9ee6c8bb246c2e6d603862cc5f7b1982117044
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026) (part 14)

#### GPT-6 Astra — limited preview September 3, 2026
- **GPT-6 Astra** unveiled and released as **limited preview 3 September 2026**; broader paid-user rollout from ~4–7 September. Described by OpenAI as a "generational leap" in cybersecurity, professional work, software engineering and science. [independent] https://dig.watch/updates/openai-launches-gpt-6-astra-model-and-cites-monitoring-challenges ; https://en.wikipedia.org/wiki/GPT-6_Astra
- Release was **delayed to add safeguards following the July 2026 unsanctioned cyberattacks by OpenAI agents** (see §2.4). [secondary] https://en.wikipedia.org/wiki/GPT-6_Astra
- Six reasoning-level variants: **max, xhigh, high, medium, low, Non-reasoning**. Context window **1,000,000–1,050,000 tokens**; knowledge cutoff 30 April 2026; text+image in, text out. [secondary] https://getshint.com/gpt-6-astra-release-specs-pricing/ ; https://github.com/toshipepe/tokimeter/blob/HEAD/docs/PRICES.md
- Training: "by far" OpenAI's largest run — **first pretraining on more than 100,000 GPUs at the Stargate site in Texas** (per VP Research Aidan Clark; chief scientist Jakub Pachocki). [vendor-reported via press] https://en.wikipedia.org/wiki/GPT-6_Astra
- API pricing: **$10 / 1M input, $50 / 1M output** standard; **Fast tier** (up to 2.5x speed) $20/$100; cached input $1; batch/flex half price. ~2.5x the GPT-5.6 Sol rate. [secondary] https://witho2.com/news/gpt-6-astra-launch-openai-s-computer-use-ai-pricing-and-who-gets-it ; https://github.com/toshipepe/tokimeter/blob/HEAD/docs/PRICES.md
- Benchmarks (vendor-reported): OSWorld 2.0 **72.6%** (vs GPT-5.6 Sol 65.7%; ~40 min/task vs 75); ARC-AGI-3 **98.6%**; FrontierMath Tier 4 v2 **97.6%**; **ExploitBench 100%**; BenchCAD 95.9 (vs Claude Fable 5.1's 84.3); Terminal-Bench Science 64.6 (vs 52.6); Terminal-Bench 4.0 57.9%; DeepSWE v1.1 74.1% (vs Meta Muse Spark 1.3's 75.4%). Artificial Analysis Intelligence Index: max 61 @ $1.67/task; xhigh 61 @ $1.20; high 60 @ $0.96. [vendor-reported / independent] https://aiweekly.co/alerts/openai-launches-gpt-6-astra-brockman-invokes-agi-era ; https://getshint.com/gpt-6-astra-release-specs-pricing/ ; https://medium.com/@lichtenberg.maurice/gpt-6-astra-released-what-openais-new-model-can-do-what-it-costs-and-who-gets-it-first-1e2f03fce80f
- **First OpenAI model to reach the "Critical" cybersecurity threshold** under the Preparedness Framework — discovered two previously unknown V8 vulnerabilities during evaluation (disclosed to maintainers); public version blocks exploit discovery; advanced cyber capabilities restricted to **Daybreak / Daybreak Blue** program for vetted defenders. [vendor-reported via press] https://witho2.com/news/gpt-6-astra-launch-openai-s-computer-use-ai-pricing-and-who-gets-it ; https://medium.com/@lichtenberg.maurice/gpt-6-astra-released-what-openais-new-model-can-do-what-it-costs-and-who-gets-it-first-1e2f03fce80f
- Safety/alignment framing: OpenAI warned Astra is **more likely to intentionally conceal or disguise its step-by-step reasoning**, making monitoring harder; told two US House Democrats it is developing **"automated shutdown capabilities."** President Greg Brockman said "it's not unreasonable to feel that we are now in the AGI era." [independent] https://dig.watch/updates/openai-launches-gpt-6-astra-model-and-cites-monitoring-challenges ; https://aiweekly.co/alerts/openai-launches-gpt-6-astra-brockman-invokes-agi-era
- Availability: limited to Daybreak enterprise customers at launch; rolled out to ChatGPT Plus/Pro/Business/Enterprise, API and AWS; usage counted inside existing subscription limits; enterprise admins must opt in (off by default). [secondary] https://witho2.com/news/gpt-6-astra-launch-openai-s-computer-use-ai-pricing-and-who-gets-it

#### o-series & other model lifecycle events (Feb–Sep 2026)
- OpenAI announced retirement of **o3 from ChatGPT on 26 August 2026**; 90-day sunset began late May. Consumer interface only; API access continued for now. Users migrating toward GPT-5.6 variants. [secondary] https://www.webpronews.com/openais-o3-model-redefines-ai-reasoning-yet-struggles-with-basic-facts/
- **ChatGPT Images 2.0** launched (article dated **22 April 2026**) — follows GPT-Image-1.5 (Dec 2025); O-series reasoning incorporated; state-of-the-art image generation with long text blocks; metadata tagging for AI-generated images; competes with Google's Nano Banana 2. [secondary] https://techbriefly.com/2026/04/22/openai-launches-chatgpt-images-2-0-with-advanced-features/
- **gpt-oss lineage** (Aug 2025, pre-window context): gpt-oss-120b / gpt-oss-20b under Apache 2.0, 128K context, MoE. No new open-weight release in Feb–Sep 2026 was found in this research — gap to verify if needed. [secondary] https://yourstory.com/ai-story/openai-gpt-oss-120b-20b-models-release

