---
id: ai-industry-kb-2026-wave6/08-mimo-and-xiaomi/main-actors
title: "Main actors"
domain: mimo-and-xiaomi
role: deep-dive
task: actor-profile
actors: ["China", "DeepSeek", "Hugging Face", "OpenRouter", "United States", "Xiaomi"]
dates: ["2025-04-30", "2025-05-30", "2026-03-11", "2026-03-18", "2026-03-19", "2026-04-22", "2026-04-23", "2026-05-27", "2026-06", "2026-06-30", "2026-09-21", "2026-09-22"]
keywords: ["agent", "agentic", "agents", "apache", "attention", "benchmark", "benchmarks", "cost", "deepseek", "disclosure", "inference", "leaderboard"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3945, 4087]
section: "§8. MiMo and Xiaomi"
sha256: dcfb7b30725c807e54e4de44ae931d34b4f3c7f8ed5f5ceac364bb9af277e56e
---

# Main actors

## Main actors
- **Xiaomi** — MiMo's owner; positions MiMo inside its "Human × Car × Home" ecosystem [SECONDARY].
- **MiMo** — Xiaomi's model family (text, omni, TTS variants) [SECONDARY].
- **Artificial Analysis** — the independent benchmark operator whose v4.3 Index scored V2.6-Pro at 46.32 [SECONDARY].
- **The "Hunter Alpha" watchers** — the community that wrongly suspected the stealth codename was DeepSeek V4; it was MiMo-V2-Pro [SECONDARY].

## Timeline and context
- **2025-04-30** — MiMo-7B family released under MIT (weights) / Apache 2.0 (repo code) [SECONDARY].
- **2026-03-18** — V2-Pro, V2-Omni, V2-TTS launched (proprietary); "Hunter Alpha"→V2-Pro, "Healer Alpha"→V2-Omni [SECONDARY].
- **2026-04-22** — V2.5 (310B, MIT) and V2.5-Pro (1.02T/42B, MIT) released; 1M context [SECONDARY].
- **2026-04-23** — V2.5-TTS-Series released, API-only, initially limited-time free [SECONDARY].
- **2026-09-21/22** — V2.6-Pro (natively omnimodal, 1.02T/42B, MIT, Live-RL-trained) and V2.6-Flash (310B/15B, MIT) released [SECONDARY].


### New verified timeline entries — expansion

- 2026-03-11: Anonymous "Hunter Alpha" appears on OpenRouter, tops daily usage charts. [SECONDARY, S27]
- 2026-03-18: Xiaomi reveals Hunter Alpha was MiMo-V2-Pro; stock +5.8%. [SECONDARY, S27][SECONDARY, S26]


- 2025-04-30: MiMo-7B base release (base section; anchor for the RL update below). [SECONDARY, S2][SECONDARY, S4]
- 2025-05-30: MiMo-7B-RL-0530 checkpoint posted to Hugging Face with 6M-instance SFT and 48K RL context. [VENDOR, S1]
- 2026-03-18: MiMo V2-Pro, V2-Omni, and V2-TTS released (base section; specification details added above). [SECONDARY, S5][SECONDARY, S6][SECONDARY, S7]
- 2026-03-19: Lei Jun announces ≥CNY60B / ~$8.7B AI investment over three years; Xiaomi shares rise ~5%. [SECONDARY, S9][SECONDARY, S10][SECONDARY, S11][SECONDARY, S12]
- 2026-04-22: MiMo V2.5 and V2.5-Pro released (base section; architecture and pricing deltas added above). [SECONDARY, S13][SECONDARY, S19]
- 2026-04-23: V2.5-TTS released (base section anchor). [SECONDARY, S17]
- 2026-05-27: Provider rate-card revision takes effect — V2.5-Pro $0.435/$0.87, V2.5 $0.14/$0.28, cache-hit tiers introduced. [SECONDARY, S14][SECONDARY, S16]
- 2026-06-30 (Beijing time): Legacy V2 model identifiers fully retired after deprecation/auto-routing period. [SECONDARY, S16]
- 2026-09-21 ~15:39 UTC: V2.6-Pro and V2.6-Flash Hugging Face checkpoints appear, ~18 seconds apart. [COMMUNITY, S24]
- 2026-09-21/22: V2.6-Pro and V2.6-Flash officially released with Live RL training disclosure, technical report, environments, and code. [SECONDARY, S21][SECONDARY, S22]
- 2026-09-22 (research date): MiClaw smartphone agent in closed beta; V2-TTS API variants still limited-time free per provider logs. [SECONDARY, S8][SECONDARY, S16]

## Implications
1. MiMo is the 2026 counterexample to license tightening: Xiaomi ships its strongest generations (V2.5-Pro, V2.6-Pro, V2.6-Flash) under plain MIT while the proprietary line (V2-Pro/V2-Omni/V2-TTS) stays closed — a two-track strategy, open weights for community reach and proprietary for product control [DIRECTIONAL].
2. V2.6-Pro's public "Live RL" (30 steps, ~750K trajectories, $2.62M) is unusually transparent post-training disclosure; treat cost/trajectory figures as secondary-sourced claims, not audited facts [SECONDARY/DIRECTIONAL].
3. AA v4.3 46.32 positions V2.6-Pro at the top of open-weight rankings on that index version — pin the methodology version whenever citing it, and never compare across v4.1.1/v4.2/v4.3 [SECONDARY].
4. Artifact-scope discipline matters for MiMo-7B: weights MIT vs repository code Apache 2.0 — quote the right artifact [SECONDARY].
5. V2.5-Pro and V2.6-Pro share a parameter profile but are distinct generations; consolidation or benchmarking work must not merge them [SECONDARY].
6. The "top open-weights model" framing of the V2.6 launch is press language — the citable, version-pinned figure is AA v4.3 46.32, nothing else [SECONDARY].


### New verified implications — expansion

- The MiMo-7B-RL-0530 update shows Xiaomi iterating on small reasoning models through post-training scale (12× SFT data, 1.5× RL context) rather than base-model retraining — the largest gains landed on math benchmarks (+11.9/+14.8 AIME points) while code gains were smaller (+3.1/+2.9 LiveCodeBench). [VENDOR, S1]
- Xiaomi's pricing trajectory is aggressively downward: V2-Pro launched at $1/$3 (≤256K) and $2/$6 (256K–1M), while V2.5-Pro fell to $0.435/$0.87 within five weeks of its own launch — a deliberate undercutting pattern consistent with the reported agent-pricing pressure narrative around Hunter Alpha. [SECONDARY, S5][SECONDARY, S14][SECONDARY, S8]
- The V2.6 "Live RL" disclosure — asynchronous GRPO at ~3.5B tokens per step with released environments and code — is unusually transparent for a frontier open-weights lab and directly targets the reproducibility gap that closed labs maintain. [VENDOR, S21][SECONDARY, S22]
- The ~$0.85M (Flash) and ~$2.62M (Pro) reported RL training costs reframe V2.6 as a post-training story: the base models already existed, and the headline capability came from a comparatively cheap RL phase — a pattern that favors labs with strong infrastructure over labs with the largest pretraining budgets. [SECONDARY, S24]
- The V2 deprecation hygiene is notable: legacy identifiers were auto-routed to V2.5 endpoints before the hard 2026-06-30 retirement, a migration pattern that reduces breaking changes for API consumers. [COMMUNITY, S16] The Beijing-time cutoff itself is single-sourced to provider logs.
- The V2.5 token-count tension (48T for the 310B model vs 27T for the 1.02T Pro) is consistent with a strategy of over-training smaller models for inference efficiency, but neither vendor nor secondary sources state this explicitly, so it remains a hypothesis, not a claim.
- V2-TTS's >100M-hour speech pretraining claim, if taken at face value, would make it one of the largest speech-pretraining runs publicly claimed by any lab. [VENDOR — single source at launch reporting; treat as a vendor claim until the technical report is verified.]
- The 6:1 sliding-window-to-global attention ratio with a 128-token local window in V2.5-Pro is one of the most specific open-weights attention disclosures of 2026, useful for inference-engine implementers even though layer counts remain undisclosed.
- MiClaw's closed beta is Xiaomi's first disclosed attempt at a system-level smartphone agent, distinct from app-level assistants — its integration into Xiaomi OS rather than a standalone app suggests the company views the OS, not the chatbot, as the agent platform.

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

- S35 — Vendor. MiMo-VL technical report PDF (architecture, 4-stage 2.4T-token training): https://github.com/XiaomiMiMo/MiMo-VL/raw/refs/heads/main/MiMo-VL-Technical-Report.pdf
- S36 — Secondary. Hugging Face paper page 2506.03569 (MiMo-VL-7B benchmarks, MORL): https://huggingface.co/papers/2506.03569
- S37 — Secondary. Dev.to complete guide to the MiMo V2 series (pricing tables, TTS dialects, integrations): https://dev.to/czmilo/mimo-v2-series-complete-guide-2026-mimo-v2-pro-mimo-v2-omni-and-mimo-v2-tts-xiaomis-agent-era-10k0
- S38 — Secondary. Quasa on MiMo-V2 family benchmarks and Omni/TTS capabilities: https://quasa.io/media/xiaomi-unleashes-mimo-v2-family-trillion-parameter-agent-powerhouse-mimo-v2-pro-ex-hunter-alpha-multimodal-omni-and-expressive-tts-hit-the-scene
- S39 — Community/secondary. cc-compatible-models pricing and integration catalog: https://github.com/Alorse/cc-compatible-models
- S40 — Community/secondary. Timbal Xiaomi MiMo provider docs (specs, pricing): https://github.com/timbal-ai/timbal/blob/HEAD/docs/models/xiaomi.mdx
- S41 — Community/secondary. wgnr-ai Xiaomi MiMo provider (V2.5-Pro benchmarks, experts, AA 54): https://github.com/wgnr-ai/xiaomi-mimo-provider

- S42 — Secondary. Geeky Gadgets on V2.5-Pro pricing and MIT licensing: https://www.geeky-gadgets.com/xiaomi-mimo-v2-5-pro-open-source-ai/

- S43 — Secondary. VentureBeat on MiMo-V2-Pro (AA 49, GDPval-AA 1426, hallucination/omniscience, 77M tokens): https://venturebeat.com/technology/xiaomi-stuns-with-new-mimo-v2-pro-llm-nearing-gpt-5-2-opus-4-6-performance
- S44 — Secondary. MEXC News on MiMo-V2-Pro (stock +5.8%, SWE 78%, $1/$3, closed source): https://www.mexc.co/en-IN/news/989975
- S45 — Secondary. Saerio on MiMo-V2-Pro (GDPval-AA 1426, AA 49, hallucination 30%): https://news.saerio.com/xiaomi-stuns-with-new-mimo-v2-pro-llm-nearing-gpt-5-2-opus-4-6-performance-at-a-fraction-of-the-cost/
- S46 — Secondary. Wikipedia Xiaomi MiMo (launch 2026-03-18, Healer Alpha, TTS dialects, licensing): https://en.wikipedia.org/wiki/Xiaomi_MiMo
- S47 — Secondary. Quasa on MiMo-V2 family launch (Mar 18–19 2026, Hunter Alpha >1T tokens, free week): https://quasa.io/media/xiaomi-unleashes-mimo-v2-family-trillion-parameter-agent-powerhouse-mimo-v2-pro-ex-hunter-alpha-multimodal-omni-and-expressive-tts-hit-the-scene
- S48 — Community. dev.to complete MiMo-V2 family guide (Omni benchmarks, TTS dialects, agent-stack roles): https://dev.to/czmilo/mimo-v2-series-complete-guide-2026-mimo-v2-pro-mimo-v2-omni-and-mimo-v2-tts-xiaomis-agent-era-10k0
- S49 — Secondary. Toolnavs on MiMo-V2 launch (Mar 18 2026, 1M ctx, 1T/42B, free first week): https://toolnavs.com/article/1224-xiaomi-released-mimo-v2-pro-omni-and-tts-betting-on-the-agent-full-stack-model
- S50 — Secondary. Qoo10 on Xiaomi AI push (Hunter Alpha 1.5T tokens, MiMo Code June 2026, V2.5-Pro 1.02T): https://www.qoo10.co.id/en/gadget/111533/xiaomis-ai-push-is-spreading-fast-and-its-newest-models-are-hard/

- S51 — Secondary. Unite.AI on MiMo-V2.6 (DeepSWE 71.9, GDPVal 1673, pricing, 30 RL steps): https://www.unite.ai/xiaomis-new-flagship-model-leads-open-weight-rankings-with-a-score-of-46/
- S52 — Secondary. VentureBeat on MiMo-V2.6-Pro (benchmark suite, open weights, 7,000+ RL environments): https://venturebeat.com/technology/better-than-deepseek-xiaomis-mimo-v2-6-pro-debuts-as-the-top-open-weights-model-in-the-world-alongside-cheaper-v2-6-flash
- S53 — Secondary. OfficeChai on MiMo-V2.6-Pro benchmarks (AA v4.3 46.32, detailed comparison): https://officechai.com/ai/xiaomi-mimo-v-2-6-pro-benchmarks/
- S54 — Secondary. TestingCatalog on MiMo-V2.6 open source (30 RL steps, 750K trajectories, $2.62M/$850K): https://www.testingcatalog.com/xiaomi-open-sources-mimo-v2-6-pro-and-flash-models/
- S55 — Secondary. TheOutpost on V2.6-Pro top open-source (AA 46, MIT license, $0.13/Index task): https://theoutpost.ai/news-story/xiaomi-mi-mo-v2-6-pro-debuts-as-top-open-source-ai-model-scoring-46-on-intelligence-index-31152/
- S56 — Secondary. LapaasVoice on MiMo-V2.6 launch (Sept 22 disclosure, 30 RL steps, training costs): https://lapaasvoice.com/xiaomi-mimo-v2-6-agent-models-2/

### Source-quality notes (methodology for this expansion)
- S5 and S6 are launch-day secondary reports; their pricing figures are treated as dated launch rate cards, superseded by S14/S16 for post-2026-05-27 rates.
- S14, S15, and S16 are provider/integration documentation and logs, reflecting post-2026-05-27 pricing and the V2 deprecation; they do not replace vendor launch claims.
- S1 is the vendor README itself; S2 summarizes it and adds the 130K-problem RL detail, so the two corroborate the 0530 benchmark deltas but share a single vendor origin.
- S24 mixes vendor disclosures with community observation; checkpoint timestamps and the 18-second upload gap are community-observed, while RL configuration and costs trace to vendor disclosures.
- Artificial Analysis scores are cited only where the methodology version (v4.3) could be pinned; versionless scores (V2-Pro 49, V2.5-Pro 54/26) were excluded rather than mixed across versions.
- Investment figures (S9–S12) are secondary financial-press reports of Lei Jun's 2026-03-19 announcement; the CNY60B figure appears in all three outlets.

