---
id: labs-hyperscalers-2026/00-labs-hyperscalers/1-6-anthropic-benchmark-snapshot-as-reported-do-not-mix-vend
title: "1.6 Anthropic benchmark snapshot (as reported — do not mix vendors' scaffolds)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: benchmark
actors: ["AWS", "Anthropic", "Hugging Face", "Microsoft", "OpenAI"]
dates: ["2025-08-31", "2026-01", "2026-02", "2026-02-05", "2026-03", "2026-03-05", "2026-03-06", "2026-09-10"]
keywords: ["benchmark", "agentic", "agents", "agi", "astra", "benchmarks", "chatgpt", "claude", "context window", "cost", "cyber", "cybersecurity"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [276, 323]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 78ac389cffaa2a638b2fe3fdc1e320e1f01f049b4b734fd642d7232573b0fdb3
---

# 1.6 Anthropic benchmark snapshot (as reported — do not mix vendors' scaffolds)

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

#### GPT-5.4 — Thinking and Pro — March 5, 2026
- **GPT-5.4** launched **Thursday 5 March 2026** (TechCrunch dated 2026/03/05), billed as "our most capable and efficient frontier model for professional work," combining coding + reasoning + native computer use in one model. Ships in three variants: standard, **GPT-5.4 Thinking** (reasoning) and **GPT-5.4 Pro** (high performance). [independent] https://techcrunch.com/2026/03/05/openai-launches-gpt-5-4-with-pro-and-thinking-versions/
- API context window **1 million tokens** — by far OpenAI's largest at the time. [independent] https://techcrunch.com/2026/03/05/openai-launches-gpt-5-4-with-pro-and-thinking-versions/
- API pricing (standard): GPT-5.4 **$2.50 / 1M in, $15 / 1M out**; mini $0.75/$4.50; nano $0.20/$1.25. [secondary] https://devtk.ai/en/blog/openai-api-pricing-guide-2026/
- Benchmarks (vendor-reported): GDPval 83.0% (record; matched/exceeded industry professionals in 83% of comparisons vs 70.9% for GPT-5.2); OSWorld-Verified 75.0% and WebArena Verified (records); SWE-Bench Pro 57.7%; BrowseComp 82.7% (Pro 89.3%); Toolathlon 54.6%; factuality: individual claims 33% less likely false vs GPT-5.2, full responses 18% less likely to contain errors; spreadsheet-modeling internal benchmark 87.3% vs 68.4%. Top score on Mercor APEX-Agents (per Mercor CEO Brendan Foody). [vendor-reported via press] https://www.businesstoday.in/technology/news/story/openai-releases-gpt-54-model-with-advanced-reasoning-coding-and-native-computer-use-519373-2026-03-06 ; https://itbrief.co.nz/story/openai-unveils-gpt-5-4-with-advanced-computer-use-tools
- New image input detail levels: "original" up to **10.24M total pixels / 6000px max dimension**; "high" raised to 2.56M pixels / 2048px. [independent] https://itbrief.co.nz/story/openai-unveils-gpt-5-4-with-advanced-computer-use-tools
- "/fast" mode in Codex: up to **1.5x token speed** without quality loss. Experimental Codex skill "Playwright (Interactive)" for visual web/Electron debugging released alongside. ChatGPT for Excel add-in launched with the model. [secondary] https://the-decoder.com/openai-launches-gpt-5-4-thinking-and-pro-combining-coding-reasoning-and-computer-use-in-one-model/

