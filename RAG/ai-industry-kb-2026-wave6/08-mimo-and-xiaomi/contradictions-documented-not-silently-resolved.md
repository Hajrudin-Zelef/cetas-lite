---
id: ai-industry-kb-2026-wave6/08-mimo-and-xiaomi/contradictions-documented-not-silently-resolved
title: "Contradictions documented (not silently resolved)"
domain: mimo-and-xiaomi
role: deep-dive
task: reference
actors: ["Hugging Face", "United States", "Xiaomi"]
dates: ["2025-04-30", "2025-05-30", "2026-05-27"]
keywords: ["agent", "agentic", "agents", "benchmark", "benchmarks", "deepseek", "leaderboard", "license", "mit license", "omni", "open-weight", "pricing"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3998, 4053]
section: "§8. MiMo and Xiaomi"
delta_of: ai-industry-kb-2026
sha256: be401946c28a99609653af565b9092dc78df48fb85c873c8ef7a34801815a183
---

# Contradictions documented (not silently resolved)

### Contradictions documented (not silently resolved)
- V2.5-Pro pricing: the base section records $0.80/$3.20; launch reporting records $1/$3 (≤256K) and $2/$6 (256K–1M); post-2026-05-27 provider logs record $0.435/$0.87. All three are dated price points on a declining curve, not a single timeless price — the expansion preserves all three with dates. [SECONDARY, S5][SECONDARY, S14]
- V2.6-Pro DeepSWE v1.1: 72.57 (post-training held-out figure) vs 71.9 (vendor launch table). The ~0.7-point gap likely reflects different evaluation settings or checkpoint snapshots; both are retained with provenance. [VENDOR, S21][VENDOR, S25]
- V2.5-Pro Artificial Analysis score: one source reports 54, another reports 26 on a September leaderboard. Without pinned AA methodology versions, neither figure is carried into the expansion. [excluded per methodology-version rule]
- V2-Pro AA score of 49 appears in weak secondary coverage without a methodology version and is excluded for the same reason. [excluded per methodology-version rule]
- V2.5 training tokens: 48T reported for V2.5 vs 27T for V2.5-Pro. Counterintuitive (the smaller model trained on more tokens); both figures retained as reported, flagged as an unresolved tension. [SECONDARY, S13]
- Hunter Alpha token volume: reports range from >1T to >1.5T depending on report date; treated as a growing cumulative figure over the stealth period. [SECONDARY, S8]

## Sources and URLs
- https://en.wikipedia.org/wiki/Xiaomi_MiMo
- https://www.infoworld.com/article/4164184/xiaomi-releases-mit‑licensed-mimo-models-for-long‑running-ai-agents.html
- https://mimo.xiaomi.com/mimo-v2-6
- https://www.unite.ai/xiaomis-new-flagship-model-leads-open-weight-rankings-with-a-score-of-46/
- https://venturebeat.com/technology/better-than-deepseek-xiaomis-mimo-v2-6-pro-debuts-as-the-top-open-weights-model-in-the-world-alongside-cheaper-v2-6-flash
- https://www.testingcatalog.com/xiaomi-open-sources-mimo-v2-6-pro-and-flash-models/
- https://mimo.mi.com/docs/en-US/updates/model


### New sources — expansion

- S1 — Vendor. MiMo-7B-RL Hugging Face README (2025-05-30 update notes, benchmark deltas): https://huggingface.co/XiaomiMiMo/MiMo-7B-RL/blob/main/README.md
- S2 — Secondary. SiliconANGLE on MiMo-7B release and RL data scale (2025-04-30): https://siliconangle.com/2025/04/30/china-ai-rising-xiaomi-releases-new-mimo-7b-models-deepseek-upgrades-prover-math-ai/
- S3 — Secondary. Dao Insights on Xiaomi's first open-source reasoning LLM: https://daoinsights.com/news/xiaomi-launches-its-first-open-sourced-reasoning-llm/
- S4 — Community/tertiary. Wikipedia overview of Xiaomi MiMo: https://en.wikipedia.org/wiki/Xiaomi_MiMo
- S5 — Secondary. BigGo on MiMo V2-Pro launch, pricing, integrations: https://biggo.com/news/202603190623_Xiaomi_MiMo_V2_Pro_AI_Model_Launch
- S6 — Secondary. The Tech Outlook on V2-Pro, V2-Omni, V2-TTS launch details: https://www.thetechoutlook.com/new-release/gadgets-release/smart-watch/xiaomi-watch-s5-with-21-day-ultra-long-battery-life-launched-in-china-xiaomi-also-introduces-new-mimo-v2-pro-mimo-v2-omni-and-mimo-v2-tts-models/
- S7 — Secondary. The Decoder on the three-model MiMo V2 launch: https://the-decoder.com/xiaomi-launches-three-mimo-ai-models-to-power-agents-robots-and-voice/
- S8 — Secondary. ChinaBizInsider on Hunter Alpha and agent-pricing pressure: https://chinabizinsider.com/xiaomi-claims-viral-hunter-alpha-models-as-mimo-v2-trio-pressuring-chinas-agent-ai-pricing/
- S9 — Secondary. Brand Equity / Economic Times on ≥CNY60B AI investment: https://brandequity.economictimes.indiatimes.com/news/business-of-brands/xiaomi-to-invest-at-least-8-7-billion-in-ai-over-next-three-years-ceo-says/129699831
- S10 — Secondary. SRN News on the $8.7B AI investment announcement: https://srnnews.com/xiaomi-to-invest-at-least-8-7-billion-in-ai-over-next-three-years-ceo-says/
- S11 — Secondary. Yicai Global on investment, R&D budget, five-year core-tech plan: https://www.yicaiglobal.com/news/chinas-xiaomi-to-invest-over-usd87-billion-in-ai-in-next-three-years-ceo-says
- S12 — Secondary. Parameter.io on the ~5% stock surge (1810.HK): https://parameter.io/xiaomi-1810-hk-stock-surges-5-following-8-7-billion-ai-investment-announcement/
- S13 — Secondary. BigGo Finance on V2.5/V2.5-Pro technical details: https://finance.biggo.com/news/unfg0p0BpwxG186N_ZKM
- S14 — Secondary. LiteLLM-rs provider docs for Xiaomi MiMo (pricing, endpoints): https://github.com/majiayu000/litellm-rs/blob/HEAD/docs/providers/xiaomi-mimo.md
- S15 — Secondary. AI Director Toolkit Xiaomi MiMo provider SKILL: https://github.com/branchingjade/aidirectortoolkit/blob/HEAD/mlops/xiaomi-mimo-provider/SKILL.md
- S16 — Secondary. LLMCapa provider update log (V2 deprecation, 2026-05-27 rates, ASR/TTS): https://github.com/awaku7/llmcapa/blob/HEAD/provider_update_log.md
- S17 — Secondary. Raycast MiMo TTS integration repo: https://github.com/semantic-craft/raycast-mimo-tts
- S18 — Secondary. AIMadeTools complete guide to the MiMo V2.5 series: https://www.aimadetools.com/blog/mimo-v2-5-series-complete-guide/
- S19 — Secondary. VentureBeat on MiMo V2.5/V2.5-Pro efficiency and agentic coding: https://venturebeat.com/ai/open-source-xiaomi-mimo-v2-5-and-v2-5-pro-are-among-the-most-efficient-and-affordable-at-agentic-claw-tasks
- S20 — Secondary. Unite.AI on V2.6-Pro and the AA v4.3 open-weight ranking: https://www.unite.ai/xiaomis-new-flagship-model-leads-open-weight-rankings-with-a-score-of-46/
- S21 — Secondary. VentureBeat on V2.6-Pro/V2.6-Flash launch and Live RL: https://venturebeat.com/technology/better-than-deepseek-xiaomis-mimo-v2-6-pro-debuts-as-the-top-open-weights-model-in-the-world-alongside-cheaper-v2-6-flash
- S22 — Secondary. Forkast News on V2.6 release timing and open-weights strategy: https://forkast.news/xiaomis-mimo-v2-6-ships-open-weights-at-frontier-class-performance-and-the-timing-is-not-an-accident/
- S23 — Secondary. OrcaRouter blog on V2.6-Flash release: https://www.orcarouter.ai/blog/xiaomi-mimo-v2-6-flash-release
- S24 — Secondary/community mix. ByteIota on the public V2.6 training run, costs, checkpoint timing: https://byteiota.com/xiaomi-mimo-v2-6-watch-a-1t-ai-model-train-live/
- S25 — Secondary. OfficeChai on V2.6-Pro benchmark table: https://officechai.com/ai/xiaomi-mimo-v-2-6-pro-benchmarks/

- S26 — Secondary. Decrypt on MiMo-V2-Pro review, benchmarks, pricing, stock move: https://decrypt.co/362633/xiaomi-mimo-v2-pro-review-so-good-mistaken-deepseek-v4
- S27 — Secondary. Particula on Hunter Alpha stealth launch, pricing FAQ, Luo Fuli: https://particula.tech/blog/xiaomi-mimo-v2-pro-hunter-alpha-explained
- S28 — Secondary. DotDotNews on V2-Pro tiered pricing and free cache writes: https://english.dotdotnews.com/a/202603/19/AP69bb9870e4b0c32d4f6c4de5.html
- S29 — Vendor. Official Xiaomi MiMo-V2.5-Pro spec page (architecture, training, post-training): https://mimo.xiaomi.com/mimo-v2-5-pro
- S30 — Vendor. Hugging Face model card XiaomiMiMo/MiMo-V2.5-Pro: https://huggingface.co/XiaomiMiMo/MiMo-V2.5-Pro
- S31 — Secondary. FoneArena on V2.5/V2.5-Pro architecture and training pipeline: https://www.fonearena.com/blog/481186/xiaomi-mimo-2-5-features.html
- S32 — Community/secondary. AgentOne free-tier research doc (MIT license, Elo, free tiers, self-hosting): https://github.com/dirk-makerhafen/agentone/blob/HEAD/docs/free-tier-research/models/mimo-v2.5-pro.md
- S33 — Secondary. BigGo Finance on V2.5-Pro efficiency, GDPVal, hardware day-0 support: https://finance.biggo.com/news/202605010126_Xiaomi-MiMo-V2.5-open-source-AI-model-token-efficient
- S34 — Community/secondary. sglang-jax cookbook recipe for V2.5-Pro on TPU: https://github.com/sgl-project/sglang-jax/blob/HEAD/docs/cookbook/autoregressive/Xiaomi/MiMo-V2.5-Pro.md

