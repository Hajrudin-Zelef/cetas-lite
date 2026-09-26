---
id: ai-industry-kb-2026/05-chinese-ai-labs-deep-file-deepseek/sources-and-urls
title: "Sources and URLs"
domain: chinese-ai-labs-deep-file-deepseek
role: deep-dive
task: actor-profile
actors: ["Alibaba", "DeepSeek", "Moonshot"]
dates: ["2026-03", "2026-03-27", "2026-08-12", "2026-09-22"]
keywords: ["agent", "agentic", "apache", "attention", "benchmark", "benchmarks", "cost", "cybersecurity", "decode", "deepseek", "inference", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2135, 2192]
section: "5. Chinese AI Labs — Deep File: DeepSeek"
sha256: 10decb73660ef3a67bcebe2bac301f23ade85f021f543f7b226d0be7b9639b1c
---

# Sources and URLs

- **The brief's three conflated events.** June 1 (price cut), Aug 13 (V4 checkpoint), Sept 10 (V4.1-Flash GA) were fused into "V4.1 announced June 1 → shipped Aug 13" — the consolidation must keep all three dated and separate; each has its own evidence chain (Digitimes, Vals.ai/SCMP, WeChat/model card).
- **The +10.6 figure had no source.** Wave 3.1's quoted-phrase search found no support for it; Vals.ai's own dated page records +9.48. Unsourced brief figures must be checked against the primary index page, not averaged or hedged.
- **Alias vs model discipline generalizes.** The July 24 lesson (routing aliases ≠ models) applies to the Sept 14 question too: whether `deepseek-v4-pro` routes to V4.1-Flash or keeps serving is a routing decision, not a model existence claim.
- **Checkpoint transparency as a norm.** 0731, 0813, `-expires-on-0910` — DeepSeek's 2026 naming embeds dates in model ids. RAG retrieval should parse date-stamped ids as version evidence.
- **The Engram kill-vs-keep split.** Killed for V4 (zero mentions in the 58-page report — buzzgrewal); kept as [VENDOR] for V4.1-Flash (196B module in the model card via secondary). The same term, two evidence grades, two verdicts — the RAG must not collapse them.
- **MIT across the whole 2026 line.** V4 Preview, V4.1-Flash — both MIT. The license is the stable constant of the line; everything else (price, architecture, routing) moves.
- **Watch V4.1-Pro.** The Sept-10 notice makes it the endpoint of the V4-Pro routing plan; it has not shipped as of 2026-09-22. Its arrival (or non-arrival) resolves the routing conflict.
- **The compatibility-routing pattern.** `deepseek-v4-flash` and `deepseek-v4-flash-vision-exp` keep answering through compatibility routing to V4.1-Flash — a soft landing that contrasts with the July 24 hard cut. The RAG should note which identifiers are hard-cut vs soft-routed.
- **The 96-hour GA-to-routing gap.** Sept 10 GA → Sept 14 planned routing change gives integrators a 4-day window; the pattern (GA first, routing change days later) mirrors the April 24 → July 24 three-month alias clock at compressed scale.
- **Rumor-phase license speculation as a general hazard.** The March 2026 Apache 2.0 expectation was wrong for the entire V4 family; license claims from pre-release coverage should be re-verified at GA for every lab, not just DeepSeek.
- **The 36kr single-source rule.** The V4-Lite ~200B figure has exactly one origin (36kr via awesomeagents). Single-sourced scale figures must be flagged as such; the RAG should record the source count alongside the figure.
- **Press-rank vs eval-rank.** Vals.ai's own #18 (2026-08-12) vs press's 12th (mid-September) are two different measurements of a moving board — not a discrepancy to reconcile, but a rank-drift fact to record.
- **The vendor's own concessions are the strongest independent check.** DeepSeek's report conceding TB-3.0 30.0 vs 43.3, TB-4.0 31.2 vs 51.8, HLE 36.8 vs 56.3 is more informative than any launch headline — always pair the launch table with the conceded evals.
- **Agentic post-training as the 2026 differentiator.** V4.1's post-training story ("large-scale automated synthesis of agent tasks and environments with progressive scaling of data, tasks, and rollouts") and the agentic benchmark headlines (TB-2.1 90.6, DeepSWE 74.2%) mark the shift from raw capability to agent-task synthesis as the post-training frontier — with the harness-sensitivity caveat attached.
- **Peak/off-peak as architecture-aligned pricing.** V4.1-Flash's asymmetric 8B-prefill/16B-decode activation split prices prefill and decode phases differently at the silicon level; the 2× peak multiplier prices them differently at the billing level. The RAG's economics section (§9/§14) should connect the two.
- **The 139-day structural cycle.** April 24 (preview) → September 10 (new architecture family) is the fastest such cycle in DeepSeek's history; the 2026 velocity (98 → 13 → 8 → 20-day checkpoint intervals) is itself a data point for lab-velocity comparisons.
- **What the file does not claim.** No full V4.1 model (only the Flash variant has shipped); no V4.1-Pro; no independent verification of 890 B/token, CSA2 mechanics, or Engram behavior; no resolution of the Sept-14 routing conflict; no Llama/Qwen/Moonshot coverage (other parts). Boundaries kept.

## Sources and URLs

- [COMMUNITY] https://github.com/ez0000001000000/blog/blob/HEAD/blog/content/deepseek-v4-leak.mdx
- [COMMUNITY] https://github.com/murataslan1/cursor-ai-tips/blob/HEAD/tips/cursor-26-features.md
- [SECONDARY] https://www.nxcode.io/resources/news/deepseek-v4-release-specs-benchmarks-2026
- [COMMUNITY] https://awesomeagents.ai/news/deepseek-v4-lite-leaked-1m-context-multimodal/
- [SECONDARY] https://kissapi.ai/blog/deepseek-v4-api-access-guide-2026.html
- [SECONDARY] https://aiforautomation.io/news/2026-03-27-deepseek-v4-trillion-params-apache-license
- [SECONDARY] https://www.geeky-gadgets.com/deepseek-v4-benchmarks-leak/
- [COMMUNITY] https://github.com/uwuclxdy/ai-pricelog/blob/HEAD/state/announce/deepseek/updates.md
- [SECONDARY] https://www.techi.com/deepseek-chat-reasoner-retirement-v4-migration/
- [COMMUNITY] https://github.com/prism-shadow/agenthub/blob/HEAD/llmsdk_docs/deepseek_v4/quickstart.md
- [COMMUNITY] https://github.com/davi0015/void/blob/HEAD/docs/note-deepseek.md
- [COMMUNITY] https://github.com/roninforge/ai-price-index/blob/HEAD/data/backfill/deepseek.md
- [SECONDARY] https://benchlm.ai/providers/deepseek
- [SECONDARY] https://www.vals.ai/models/deepseek_deepseek-v4-pro-0813
- [SECONDARY] https://www.scmp.com/tech/big-tech/article/3363895/deepseeks-updated-v4-pro-ai-model-struggles-benchmarks-shines-cybersecurity
- [SECONDARY] https://www.coinlive.com/news/deepseek-s-updated-v4-pro-falls-short-on-general-benchmarks
- [SECONDARY] https://www.nationpress.com/sciencetech/deepseek-v4-pro-ranks-12th-on-vals-index
- [SECONDARY] https://en.it-daily.net/shortnews-en/deepseek-v4-pro-0813-benchmarks
- [DIRECTIONAL] https://miraflow.ai/blog/deepseek-v4-pro-0813-architecture-benchmarks-2026
- [SECONDARY] https://www.Digitimes.Com/tag/hardware/00413567.html
- [SECONDARY] https://cellcog.ai/blog/deepseek-v4-1-flash-release-date/
- [SECONDARY] https://getshint.com/deepseek-v4-1-flash-beta-specs-access/
- [SECONDARY] https://www.orcarouter.ai/blog/deepseek-v4-1-flash-inference-stack
- [SECONDARY] https://www.davidborish.com/post/deepseek-s-v4-1-flash-cuts-memory-cost-fourfold-which-changes-the-self-hosting-math
- [SECONDARY] https://oliver-foster.medium.com/deepseek-v4-1-flash-the-release-that-retired-its-own-flagship-1e883e151bb6
- [SECONDARY] https://magicshot.ai/news/deepseek-v4-1-flash-lands-with-native-vision-cheaper-tokens
- [SECONDARY] https://apimaster.ai/blog/deepseek-v4-1-flash-api-pricing-2026
- [SECONDARY] https://gnana70.medium.com/deepseek-v4-1-flash-the-architecture-benchmarks-and-real-cost-c18709aa4b8c
- [COMMUNITY] https://github.com/973bearwallow-coder/hermes-workspace/blob/HEAD/memory/model_scout_2026-06-15.md
- [SECONDARY] https://buzzgrewal.medium.com/how-ai-finally-killed-quadratic-attention-nsa-mamba-3-and-the-architectures-making-million-token-f011129c8dfa
- [SECONDARY] https://flowtivity.ai/blog/deepseek-v4-1-flash-benchmarks/
- [SECONDARY] https://temperaturezero.com/2026/09/10/deepseek-v4-1-flash-inference-kv-cache-benchmark-analysis/
- [SECONDARY] https://aicraftjournal.com/articles/deepseek-v4-1-flash-peak-0-30-out-1-20-mit-v4-pro-sep-14
- [SECONDARY] https://gotechpresso.com/blog/deepseek-v4-1-flash
- [DIRECTIONAL] https://miraflow.ai/blog/deepseek-v4-1-flash-causal-encoder-decoder-explained-2026
- [SECONDARY] https://datanorth.ai/news/deepseek-releases-deepseek-v4-1-flash
- [SECONDARY] https://aiwith.me/blog/deepseek-v4-1-flash-launch/

