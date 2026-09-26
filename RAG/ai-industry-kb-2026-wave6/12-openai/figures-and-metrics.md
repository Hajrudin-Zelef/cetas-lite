---
id: ai-industry-kb-2026-wave6/12-openai/figures-and-metrics
title: "Figures and metrics"
domain: openai
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Cerebras", "OpenAI", "OpenRouter"]
dates: ["2025-12-11", "2026-02-05", "2026-02-12", "2026-03-05", "2026-03-31", "2026-04-23", "2026-07-09", "2026-08-10", "2026-08-11", "2026-09-02", "2026-09-03"]
keywords: ["agent", "agi", "alignment", "astra", "benchmark", "benchmarks", "claude", "cyber", "cybersecurity", "fable 5", "funding", "gpt-5.6"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6040, 6096]
section: "§12. OpenAI"
delta_of: ai-industry-kb-2026
sha256: dbe8200942aaa79b824c2219e694e39a73d15df88d0542a476e10145c7994722
---

# Figures and metrics

## Figures and metrics
| Model | Launch | Input / Output per 1M | Note |
|---|---|---|---|
| GPT-5.2 | 2025-12-11 | — | 2025 line closer [SECONDARY] |
| GPT-5.3-Codex | 2026-02-05 | — | ~25% faster, coding records [SECONDARY] |
| GPT-5.3-Codex-Spark | 2026-02-12 (preview) | — | Preview tier [SECONDARY] |
| GPT-5.4 | 2026-03-05 | — | Native computer use [SECONDARY] |
| GPT-5.4-mini | Sept 2026 | $0.75 / $4.50 | Legacy band [SECONDARY] |
| GPT-5.5 / 5.5 Pro | 2026-04-23 | $5/$30 · Pro $30/$180 | Codename "Spud" confirmed [SECONDARY/VENDOR] |
| GPT-5.6 Luna | 2026-07-09 (GA) | $0.20 / $1.20 | Was $1/$6 at GA; 20× under Sol [VENDOR] |
| GPT-5.6 Terra | 2026-09-02 | $2.00 / $12.00 | Mid tier [VENDOR] |
| GPT-5.6 Sol | 2026-07-09 (GA) | $4.00 / $20.00 promo (std $5/$30) | Cut 20/33% Aug 21, promo thru Nov 21 [VENDOR] |
| GPT-5.6-Cyber | 2026-08-10 | — | Daybreak Blue/Red tiers [SECONDARY] |
| GPT-6 Astra | 2026-09-03 | $10.00 / $50.00 (cache $1) | 272K cliff: 2× in/cache, 1.5× out [VENDOR/SECONDARY] |
| GPT-4o | legacy | $2.50 / $10.00 (cache $1.25) | Legacy tier [SECONDARY] |

- Benchmark snapshot (all [SECONDARY], version-pinned): AA Index v4.3 — Astra (max) **53** (tie Fable 5.1), $3.26/task; SWE-bench Verified — Sol Max **96.2%** (saturated board); TB 4.0 — Sol 37.3%, Astra 58.18–60% (per-snapshot); LMArena text — Sol 1514; factuality-weighted Arena — GPT-5.5 up 13 places to #7.
- HF snapshot: gpt-oss-120b **4.3M+** downloads [COMMUNITY].


### New verified metrics — expansion

- GPT-4.1 SWE-bench Verified: 54.6% (vendor-reported via launch coverage) [SECONDARY](http://en.zicos.com/tech/i32171379-OpenAI-Unveils-Coding-Focused-GPT-41-While-Phasing-Out-GPT-45.html)
- GPT-4.1 Video-MME: 72% vs GPT-4o's 65.3% (vendor-reported) [SECONDARY](http://en.zicos.com/tech/i32171379-OpenAI-Unveils-Coding-Focused-GPT-41-While-Phasing-Out-GPT-45.html)
- GPT-4.1 MMMU: 75% (per one launch-day table), matching GPT-4.5 and above GPT-4o's 69% [SECONDARY](http://techtarget.com/whatis/feature/GPT-41-explained-Everything-you-need-to-know)
- GPT-5.2 GDPval: 70.9% (vendor-reported) vs 38.8% for GPT-5.1; errors 30% less frequent (vendor claim) [SECONDARY](https://www.pymnts.com/artificial-intelligence-2/2025/openai-says-new-ai-model-gpt-5-2-unlocks-even-more-economic-value/)
- GPT-5.3-Codex OpenRouter standardized harness: GPQA Diamond 91.5% [SECONDARY](https://openrouter.ai/openai/gpt-5.3-codex)
- GPT-5.3-Codex OpenRouter standardized harness: HLE 42.5% [SECONDARY](https://openrouter.ai/openai/gpt-5.3-codex)
- GPT-5.3-Codex OpenRouter standardized harness: Terminal-Bench Hard 53.0% [SECONDARY](https://openrouter.ai/openai/gpt-5.3-codex)
- GPT-5.3-Codex-Spark throughput: >1,000 tokens/sec on Cerebras WSE-3 (vendor-reported via launch coverage) [SECONDARY](https://www.cometapi.com/what-is-gpt-5-3-codex-spark-how-to-use-it/) [SECONDARY](https://www.technobezz.com/news/openai-launches-gpt-53-codex-spark-on-cerebras-hardware)
- GPT-5.4 vendor benchmarks (all vendor claims, not independent): OSWorld-Verified 75.0% [SECONDARY](https://github.com/sbley/claude-code-agent-test/issues/381)
- GPT-5.4 vendor: WebArena-Verified 67.3% [SECONDARY](https://github.com/sbley/claude-code-agent-test/issues/381)
- GPT-5.4 vendor: GDPval 83.0% [SECONDARY](https://github.com/sbley/claude-code-agent-test/issues/381)
- GPT-5.4 vendor: SWE-bench Pro 57.7% [SECONDARY](https://github.com/sbley/claude-code-agent-test/issues/381)
- GPT-5.4 vendor: Terminal-Bench 2.0 75.1% [SECONDARY](https://github.com/sbley/claude-code-agent-test/issues/381)
- GPT-5.4 vendor: BrowseComp 82.7% [SECONDARY](https://github.com/sbley/claude-code-agent-test/issues/381)
- GPT-5.4 Tool Search vendor claim: 47% token reduction over 250 Scale MCP Atlas tasks across 36 servers [SECONDARY](https://almcorp.com/blog/gpt-5-4/)
- GPT-5.5 vendor benchmarks (all vendor claims): Terminal-Bench 2.0 82.7% [SECONDARY](https://github.com/mattrobenolt/pi-configs/blob/HEAD/skills/model-guide/research/gpt-5.5.md) [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- GPT-5.5 vendor: SWE-bench Pro 58.6% [SECONDARY](https://github.com/mattrobenolt/pi-configs/blob/HEAD/skills/model-guide/research/gpt-5.5.md) [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- GPT-5.5 vendor: GDPval 84.9% [SECONDARY](https://github.com/mattrobenolt/pi-configs/blob/HEAD/skills/model-guide/research/gpt-5.5.md) [SECONDARY](https://earlyterms.com/term/gpt-5-5)
- GPT-5.5 vendor: FrontierMath tiers 1–3 51.7%, Tier 4 35.4% [SECONDARY](https://github.com/mattrobenolt/pi-configs/blob/HEAD/skills/model-guide/research/gpt-5.5.md)
- GPT-5.5 vendor: MRCR v2 at 1M context 74.0% vs GPT-5.4's 36.6% [SECONDARY](https://github.com/mattrobenolt/pi-configs/blob/HEAD/skills/model-guide/research/gpt-5.5.md)
- GPT-5.6 Sol vendor Coding Agent Index 80.0 vs independent/versioned scores ~59 and 66.6 (not the same benchmark, not interchangeable) [SECONDARY](http://datafloq.com/gpt-5-6-pricing-explained-why-sol-terra-and-luna-turn-ai-buying-into-a-routing-problem/)
- GPT-5.6-Cyber vendor Advanced Cybersecurity Completion Rate: 95% (5.6-Cyber) vs 57.3% (5.5-Cyber) vs 1.5% (standard Sol) vs ~2% (Daybreak Blue) [SECONDARY](https://pondero.ai/news/2026-08-11-openai-daybreak-gpt56-cyber/)
- GPT-6 Astra vendor benchmarks (vendor claims only): ARC-AGI-3 99.9% (vendor harness) vs 62.7% (standard harness, per independent reporting) [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide) [SECONDARY](https://www.magicboat.net/mobile/en/blog/gpt-6-astra-vs-claude-fable-5-1-which-is-better-for-creators)
- GPT-6 Astra vendor: OSWorld 2.0 72.6% [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra vendor: Terminal-Bench Science 0.1 64.6% [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra vendor: Terminal-Bench 4.0 57.9% [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra vendor: DeepSWE v1.1 74.1% [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra vendor: FrontierMath Tier 4 v2 97.6% [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra vendor: HLE with tools 57.2% [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT Image 1.5 vendor-derived: prompt alignment 91.2% [SECONDARY](https://www.red94.net/news/89210-ai-image-generator-wars-heat-up-as-openai-unveils-gpt-image-1-5-promising-4x-fas/)
- GPT Image 1.5 vendor-derived: diagram/flowchart 96.9 [SECONDARY](https://www.red94.net/news/89210-ai-image-generator-wars-heat-up-as-openai-unveils-gpt-image-1-5-promising-4x-fas/)
- GPT Image 1.5 vendor-derived: single-turn BinaryEval edit 100% [SECONDARY](https://www.red94.net/news/89210-ai-image-generator-wars-heat-up-as-openai-unveils-gpt-image-1-5-promising-4x-fas/)
- GPT Image 1.5 vendor-derived: visual quality 89.96% single-turn / 89.46% multi-turn [SECONDARY](https://www.red94.net/news/89210-ai-image-generator-wars-heat-up-as-openai-unveils-gpt-image-1-5-promising-4x-fas/)
- OpenAI operating scale (2026-03-31 funding close): $2B/mo revenue, 900M weekly users, 50M paid, 15B API tokens/min, enterprise >40% of revenue [SECONDARY](https://www.advisorperspectives.com/articles/2026/04/01/openai-valued-852-billion-completing-122-billion-round)

