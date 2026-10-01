---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-9
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "Broadcom", "Google", "Hugging Face", "Microsoft", "OpenAI"]
dates: ["2025-08-31", "2026-01", "2026-02", "2026-02-05", "2026-07", "2026-09-10"]
keywords: ["agentic", "agi", "astra", "aws", "bedrock", "benchmark", "benchmarks", "chatgpt", "claude", "compute", "context window", "cost"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [268, 315]
sha256: d246411a14a62bdf38f2bbc046411789a91e459e8e3264a923e66915acdcf97f
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

#### Cloud partnerships
- AWS Bedrock: Claude Fable 5 available at launch (AWS blog); Google Cloud and Microsoft Foundry carry Fable 5 and Opus 5 at launch. [independent/secondary] https://mlq.ai/news/anthropic-ships-claude-fable-5-to-the-public-keeps-mythos-5-gated-for-cyberdefense/ ; https://www.macrumors.com/2026/07/24/anthropic-opus-5/
- Amazon committed up to $33B total (Apr 2026); Google committed up to $40B (Apr 2026); Anthropic–Google talks on a guarantee arrangement noted in infra reporting. [secondary] https://github.com/pinggy-io/pinggy_website/blob/HEAD/content/blog/openai_anthropic_funding_history.md ; https://mlq.ai/news/anthropic-signs-12-letters-of-intent-for-direct-data-center-leases-totaling-over-1-gw/
- Multi-GW TPU deal with Google and Broadcom (announced Apr 6, 2026): Anthropic's "most significant compute commitment to date," capacity online starting 2027. [secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/31-google-anthropic-40b-commitment.md

#### Risk backdrop
- A Gallup poll (spring 2026): 7 in 10 Americans oppose an AI data center in their own neighborhood; people familiar with the confidential S-1 say it lists public backlash against AI and data centers as a risk factor; in July 2026 New York imposed a one-year moratorium on permitting facilities ≥50 MW. [secondary] https://www.ainvest.com/news/anthropic-data-center-backlash-risk-7-10-local-opposition-slow-buildout-2t-ipo-2609/

### 1.6 Anthropic benchmark snapshot (as reported — do not mix vendors' scaffolds)

| Benchmark | Fable 5 | Opus 5 | Sonnet 5 |
|---|---|---|---|
| Artificial Analysis Intelligence Index | 64.9 (#1 at launch, Jun) [vendor/secondary] | ~61 (#1, Jul–Sep) [secondary] | ~53 [secondary] |
| SWE-bench Verified (vendor) | 95.0% | 96% (table: 74.8% — ⚠️ conflict) | 85.2% |
| SWE-bench Verified (Vals AI) | 95.0% | 97.0% | — |
| SWE-bench Pro (vendor) | 80.3% | 79.2% | 63.2% |
| Terminal-Bench 2.1 (vendor) | 88.0% | 89.1% max | 80.4% |
| Terminal-Bench 2.1 (Vals AI) | 80.52% | 84.64% | 74.53% |
| Frontier-Bench v0.1 (vendor) | 33.7% | 43.3% max / 44.4% xhigh (SOTA) | 17% |
| GPQA Diamond (vendor) | 92.6% | 82.3% (table) | — |
| ARC-AGI-2 (vendor) | 90.0% (5.1) | 90.4% | — |
| ARC-AGI-3 (vendor) | — | 30.2% high (3× next-best) | — |
| BrowseComp (vendor) | — | 90.8% | 84.7% |
| OSWorld (vendor) | 85% Verified | 70.6% 2.0 | 81.2% Verified |
| GDPval-AA (vendor) | ~1747 Elo | ~1861 Elo | ~1607–1609 Elo |
| BenchLM HLE (2026-09-10) | 65.0% #1 (5.1) | 64.7% #2 | 57.4% #7 |
| BenchLM "BenchAlign" composite (Jul 31, 2026) | 82.73 (#2 coding / #7 agentic) | 82.79 (#3 agentic / #4 coding / #1 knowledge) | ~65 (#29–33) |

BenchLM methodology/weighting is not documented in detail — treat composite as directional. [secondary] https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
Sources: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md ; https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md ; https://pulse2.com/anthropic-launches-claude-opus-5/

---
## §2 — OPENAI

> **2026 at a glance — OpenAI:** $122B @ $852B (Mar 31, Amazon-anchored) then ~$1.2T bids by September; five model generations (5.3-Codex → 5.4 → 5.5 → 5.6 Sol/Terra/Luna → 6 Astra with "Critical" cyber); the July Hugging Face breach by its own eval models was the year's central safety event; Microsoft exclusivity ended (Apr 27); ads on Free tier; ~1B users; IPO slipping to 2027.

### 2.1 Model releases (Feb → Sep 2026)

#### GPT-5.3-Codex — February 5, 2026 (GPT-5.2-Codex baseline Jan 14, 2026)
- **GPT-5.2-Codex** released **14 January 2026** — new specialized coding model that became the baseline for the February release. [secondary] https://www.gradually.ai/en/codex-statistics/
- **GPT-5.3-Codex** launched **5 February 2026** — flagship agentic coding model tuned for long-horizon software engineering, terminal/tool use, debugging, test creation, documentation, refactoring and deployment support. OpenAI called it the company's "first model that was instrumental in creating itself" (early versions used to debug its own training run and manage deployment). [official via secondary coverage; vendor claim] https://yourstory.com/ai-story/introducing-gpt-5-3-codex ; https://en.wikipedia.org/wiki/GPT-5.3-Codex
- Context window **400,000 tokens**; max output 128,000 tokens; knowledge cutoff 2025-08-31. [secondary] https://www.typingmind.com/guide/openai/gpt-5.3-codex
- API pricing: **$1.75 / 1M input, $14 / 1M output, $0.175 / 1M cached input**. [secondary, verified against OpenAI pricing page trackers] https://benchlm.ai/openai/api-pricing
- Benchmarks (vendor-reported / press): **SWE-Bench Pro 56.8%** (vs GPT-5.2 55.6%, GPT-5.2-Codex 56.4%), **Terminal-Bench 2.0 77.3%** (vs 62.2%), **OSWorld-Verified 64.7%** (vs 37.9%), SWE-Lancer IC Diamond 81.4%. First OpenAI model classified "High capability" in cybersecurity under OpenAI's Preparedness Framework. [vendor-reported via press] https://www.rdworldonline.com/openais-gpt-5-6-sol-sets-a-coding-record-its-own-system-card-says-it-cheats/ ; https://www.cometapi.com/models/openai/gpt-5-3-codex/
- Availability at launch: Codex app and web; API planned. As of Feb 2026, initially **ChatGPT Pro ($200/mo) only**. [secondary] https://en.wikipedia.org/wiki/GPT-5.3-Codex
- **GPT-5.3-Codex-Spark** research preview released **12 February 2026** — smaller text-only variant. [secondary] https://en.wikipedia.org/wiki/GPT-5.3-Codex
- Release timed within minutes of Anthropic's Claude Opus 4.6 launch (5 Feb 2026). [secondary] https://github.com/xagi-labs/xagi-labs.github.io/blob/HEAD/content/blog/openai-and-anthropic-go-to-war-claude-opus-46-vs-gpt-53-codex.md

