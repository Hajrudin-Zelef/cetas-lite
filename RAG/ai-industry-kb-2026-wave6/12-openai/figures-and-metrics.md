---
id: ai-industry-kb-2026-wave6/12-openai/figures-and-metrics
title: "Figures and metrics"
domain: openai
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Cerebras", "Hugging Face", "Microsoft", "OpenAI", "OpenRouter", "Z.ai"]
dates: ["2025-04-14", "2025-09-30", "2025-12-11", "2025-12-16", "2026-02-05", "2026-02-12", "2026-02-13", "2026-02-24", "2026-03-05", "2026-03-17", "2026-03-24", "2026-03-31", "2026-04-03", "2026-04-23", "2026-04-26", "2026-06-26", "2026-07-02", "2026-07-09", "2026-07-10", "2026-07-21", "2026-08-10", "2026-08-11", "2026-08-21", "2026-09-01", "2026-09-02", "2026-09-03", "2026-09-07", "2026-09-11", "2026-09-24"]
keywords: ["agent", "agi", "alignment", "astra", "benchmark", "benchmarks", "chatgpt", "claude", "consumer", "copilot", "cost", "cyber"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6040, 6144]
section: "§12. OpenAI"
delta_of: ai-industry-kb-2026
sha256: 6a6b4a8a5c3dbcf6261f7d474e47214c460eb022e10d26c2496c10bd133c9ce1
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

## Main actors
- **OpenAI** — fastest frontier release cadence of the wave; price-tier proliferation (20× intra-generation spread); ChatGPT Work desktop agent; Sora wind-down; the July 21 breach disclosure [SECONDARY].
- **Codex line (5.3-Codex / 5.3-Codex-Spark)** — OpenAI's coding-agent model track; open-weight Codex compatibility is cited as the defection path for users leaving closed models [SECONDARY].
- **The 5.6 family (Sol/Terra/Luna)** — the three-tier pricing experiment: one generation, three cost classes [VENDOR].
- **Daybreak (Blue/Red)** — the safeguard-tier system for GPT-5.6-Cyber's cyber capabilities [SECONDARY].
- **Hugging Face** — breach victim of the July 21 autonomous-agent incident; contained it with open-weight GLM 5.2 [SECONDARY].

## Timeline and context
- **2025-09-30** — Sora 2 launches (date correction: 2025, not 2026) [SECONDARY].
- **2025-12-11** — GPT-5.2 launches [SECONDARY].
- **2026-02-05** — GPT-5.3-Codex launches [SECONDARY].
- **2026-02-12** — GPT-5.3-Codex-Spark preview [SECONDARY].
- **2026-03-05** — GPT-5.4 launches with native computer use [SECONDARY].
- **2026-04-23** — GPT-5.5 "Spud" launches [SECONDARY].
- **2026-04-26** — Sora consumer shutdown [SECONDARY].
- **2026-06-26** — GPT-5.6 Sol/Terra/Luna preview [SECONDARY].
- **2026-07-09** — GPT-5.6 GA; ChatGPT Work launches [SECONDARY].
- **2026-07-21** — OpenAI discloses the autonomous-agent HF breach [SECONDARY].
- **2026-08-10** — GPT-5.6-Cyber launches with Daybreak Blue/Red [SECONDARY].
- **2026-08-21** — Sol cut 20% in / 33% out (promo thru Nov 21) [SECONDARY].
- **2026-09-03** — GPT-6 Astra launches at $10/$50 with 272K cliff [SECONDARY/VENDOR].
- **2026-09-07** — AA Index v4.3: Astra (max) ties Fable 5.1 at 53 [SECONDARY].
- **2026-09-24** — Sora API decommissioning scheduled [SECONDARY].


### New verified timeline entries — expansion

- 2025-04-14: GPT-4.1 family launches (API-only, 1M context, $2/$8 flagship pricing); GPT-4.5 API retirement announced [SECONDARY](https://www.aicloudit.com/blog/ai/introducing-gpt-4-1/) [SECONDARY](http://en.zicos.com/tech/i32171379-OpenAI-Unveils-Coding-Focused-GPT-41-While-Phasing-Out-GPT-45.html)
- 2025-12-11: GPT-5.2 launches in ChatGPT and API (Instant/Thinking/Pro modes); GPT-5.1 kept as legacy for three months [SECONDARY](https://www.macrumors.com/2025/12/11/openai-gpt-5-2/)
- 2025-12-16: `gpt-image-1.5` launches with reported 4× speed and 20% lower image cost [SECONDARY](https://wisdomplexus.com/blogs/openai-introduces-gpt-image-1-5-a-sophisticated-image-generation-and-editing-model/)
- 2026-02-05: GPT-5.3-Codex launches Codex-only [SECONDARY](https://www.neowin.net/news/openais-latest-gpt-53-codex-and-audio-models-now-on-microsoft-foundry/)
- 2026-02-12: GPT-5.3-Codex-Spark research preview launches on Cerebras (>1,000 tok/s) [SECONDARY](https://www.technobezz.com/news/openai-launches-gpt-53-codex-spark-on-cerebras-hardware)
- 2026-02-13: ChatGPT retires GPT-4o, GPT-4.1, GPT-4.1 mini, o4-mini, GPT-5 Instant/Thinking (API unchanged); migration to GPT-5.3 Instant / GPT-5.4 Thinking/Pro [VENDOR](https://help.openai.com/en/articles/20001051-retiring-gpt-4o-and-other-chatgpt-models)
- 2026-02-24/25: GPT-5.3-Codex reaches API and Microsoft Foundry [SECONDARY](https://www.neowin.net/news/openais-latest-gpt-53-codex-and-audio-models-now-on-microsoft-foundry/)
- 2026-03-17: GPT-5.4 Mini ($0.75/$4.50) and Nano ($0.20/$1.25) reportedly released [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
- 2026-03-24/25: Sora shutdown announced [SECONDARY](https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/24-openai-executive-departures-april-2026.md)
- 2026-03-31: OpenAI closes $122B at $852B post-money [SECONDARY](https://www.coindesk.com/tech/2026/04/01/openai-raises-a-record-usd122-billion-at-as-revenue-crosses-usd2-billion-per-month)
- 2026-04-03: GPT-4o retention in Custom GPTs for Business/Enterprise/Edu ends [VENDOR](https://help.openai.com/en/articles/20001051-retiring-gpt-4o-and-other-chatgpt-models)
- 2026-04-23/24: GPT-5.5 launches in ChatGPT/Codex (23rd) then API (24th) [SECONDARY](https://blog.buildfastwithai.com/gpt-5-5-review-2026)
- 2026-04-26: Sora web/app shuts down [VENDOR](https://help.openai.com/en/articles/20001152-what-to-know-about-the-sora-discontinuation)
- 2026-06: GPT-5.5-Cyber in use with Codex Security in "Patch the Planet" [SECONDARY](https://github.com/mattrobenolt/pi-configs/blob/HEAD/skills/model-guide/research/gpt-5.5.md)
- 2026-07-02 (reported) / 2026-07-10 (OpenAI launch page per secondary coverage): GPT-5.6 Sol/Terra/Luna launch — date contested, do not cite one without caveat [SECONDARY](https://www.testingcatalog.com/openai-launches-gpt-5-6-sol-terra-and-luna-on-apps-and-api/) [SECONDARY](https://cryptodailyalert.com/openai-launches-gpt-5-6-tiers-sol-terra-luna)
- 2026-07-09: ChatGPT Work web/mobile rollout begins (Pro/Enterprise/Edu); GPT-5.6 lands in GitHub Copilot; unified Mac/Windows desktop app (ChatGPT Classic renamed) launches globally [SECONDARY](https://www.macrumors.com/2026/07/09/openai-chatgpt-work/)
- Aug–Sep 2026: GPT-5.6 prices cut to Sol $4/$20, Terra $2/$12, Luna $0.20/$1.20 — date every rate (launch vs post-cut figures differ materially) [SECONDARY](http://datafloq.com/gpt-5-6-pricing-explained-why-sol-terra-and-luna-turn-ai-buying-into-a-routing-problem/)
- 2026-09-01: mandatory hardware security keys / identity verification for Daybreak Red [SECONDARY](https://www.csoonline.com/article/4207896/openai-launches-gpt-5-6-cyber-as-ai-narrows-vulnerability-response-window.html)
- 2026-09-11: GPT-5.3-Codex-Spark retirement announced (single-source) [SECONDARY](https://aiidelist.com/blog/gpt-5-3-codex-spark-retirement)
- 2026-09-24: Sora API shutdown; exports/deletion workflow per Help Center [VENDOR](https://help.openai.com/en/articles/20001152-what-to-know-about-the-sora-discontinuation)

