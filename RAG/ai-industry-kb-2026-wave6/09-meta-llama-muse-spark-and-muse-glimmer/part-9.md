---
id: ai-industry-kb-2026-wave6/09-meta-llama-muse-spark-and-muse-glimmer/part-9
title: "§9. Meta: Llama, Muse Spark, and Muse Glimmer (part 9)"
domain: meta-llama-muse-spark-and-muse-glimmer
role: deep-dive
task: actor-profile
actors: ["Intel", "Meta", "Microsoft"]
dates: ["2025-05", "2026-08-09", "2026-08-10"]
keywords: ["llama", "muse", "muse spark", "agent", "agentic", "benchmark", "benchmarks", "capex", "gqa", "intel", "moe", "multimodal"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4525, 4569]
section: "§9. Meta: Llama, Muse Spark, and Muse Glimmer"
delta_of: ai-industry-kb-2026
sha256: d3d0e742c23c5b8fb0e5fd7f9029627ebddd80b793322d85b5c1a0e4702bd8ed
---

# §9. Meta: Llama, Muse Spark, and Muse Glimmer (part 9)

- S1 — Tertiary (Wikipedia mirror). Llama family overview, 3.x specs, layers, token counts: https://pengen.diewe.workers.dev/empiree/will-https-en.wikipedia.org/wiki/Llama_(language_model)
- S2 — Secondary. Beebom on Llama 3.3 70B matching 405B performance: https://beebom.com/meta-llama-3-3-70b-model-matches-405b-performance/
- S3 — Secondary. Novita comparison guide across Llama 3 models: https://medium.com/@marketing_novita.ai/which-llama-3-model-is-right-for-you-a-comparison-guide-87de6c017c48
- S4 — Secondary. ITBear on the Llama 3.3 unveiling: https://www.itbear.com/technews/meta-unveils-llama-3-3-a-70b-parameter-ai-model-that-matches-the-performance-of-a-405b-version/
- S5 — Secondary. AIWiki on Llama 4 Scout/Maverick specs: https://aiwiki.ai/wiki/llama_4_scout_maverick
- S6 — Secondary. AICerts on Llama 4 launch and open-source multimodal positioning: https://www.aicerts.ai/news/metas-llama-4-and-the-rise-of-open-source-multimodal-ai/
- S7 — Secondary. ETV Bharat on the Llama 4 launch (Scout/Maverick/Behemoth): https://www.etvbharat.com/en/!technology/meta-introduces-new-llama-4-models-scout-maverick-behemoth-enn25040702037
- S8 — Secondary. TestingCatalog on Llama 4 10M context, MoE, and the LMArena checkpoint controversy: https://www.testingcatalog.com/llama-4-brings-10m-token-context-and-moe-architecture-with-3-new-models/
- S9 — Secondary. Computerworld on the Behemoth pause and capability concerns: https://www.computerworld.com/article/3987990/meta-hits-pause-on-llama-4-behemoth-ai-model-amid-capability-concerns.html
- S10 — Secondary. TMTPost on Behemoth timeline: http://en.tmtpost.com/post/7567992
- S11 — Secondary. TokenMix on Behemoth still-unreleased status in 2026: https://tokenmix.ai/blog/llama-4-behemoth-still-training-2026
- S12 — Secondary. AI-Intel brief 2026-08-09 (Spark 1.2 board check): https://github.com/krish-shahh/ai-intel/blob/HEAD/briefs/2026-08-09-evening.md
- S13 — Secondary. TheAIRankings on Muse Spark 1.2: https://theairankings.com/meta/muse-spark/
- S14 — Secondary. Neomanex on Muse Code and Spark 1.2 launch: https://neomanex.com/news/meta-muse-code-spark-1-2-launch
- S15 — Secondary. Mixpeek model page for Muse Spark: https://mixpeek.com/model/meta/muse-spark
- S16 — Secondary. Concord News Now on Muse Code and contributor-pricing data use: http://news.concordnewsnow.com/story/609579/meta-enters-the-terminal-coding-agent-race-with-muse-code-as-contributor-pricing-raises-new-datause-questions.html
- S17 — Secondary. SingularityByte on Muse Glimmer 30B specs and DFlash: https://singularitybyte.com/models/meta-muse-glimmer-30b-local-agent-model.html
- S18 — Secondary. Forkast on Glimmer and local-agent strategy: https://forkast.news/metas-muse-glimmer-30b-signals-a-shift-toward-local-agent-dominance/
- S19 — Secondary. ExplainX on Glimmer open-weight 30B agentic specs: https://explainx.ai/blog/meta-muse-glimmer-open-weight-30b-agentic-model-2026
- S20 — Secondary. Tech-Insider on Glimmer as open-weight model: https://tech-insider.org/meta-muse-glimmer-open-weight-ai-model-2026/
- S21 — Secondary. Temperature2 on the 2026-08-10 Glimmer release: http://temperature2.com/p/2026-08-10-meta-muse-glimmer-open-weight-model/
- S22 — Secondary. Shattered on the Spark 1.3 benchmark card: https://shattered.io/meta-muse-spark-1-3-agentic-coding-model-2026/
- S23 — Secondary. MarkTechPost on Spark 1.3 efficiency gains: https://www.marktechpost.com/2026/09/03/meta-ai-released-muse-spark-1-3-an-agentic-coding-model-that-uses-20-fewer-tool-calls-and-25-fewer-tokens-than-muse-spark-1-2/
- S24 — Secondary. Eesel on Spark 1.3: https://www.eesel.ai/blog/muse-spark-1-3
- S25 — Secondary. 36Kr Europe on Spark 1.3: http://eu.36kr.com/en/p/3967796330082819
- S26 — Secondary. YourStory on LlamaCon 2025 launches: https://yourstory.com/ai-story/meta-llamacon-2025-mark-zuckerberg-ai-launch
- S27 — Secondary. Neowin on the Meta AI app: https://www.neowin.net/amp/meta-ai-gets-a-dedicated-app-with-voice-first-experience-powered-by-llama-4/
- S28 — Secondary. TechCrunch on LlamaCon 2025: https://techcrunch.com/2025/04/29/metas-llamacon-was-all-about-undercutting-openai/embed/
- S29 — Secondary. The Decoder on the Meta AI app and Llama API: http://the-decoder.com/meta-launches-ai-assistant-app-and-llama-api-platform/
- S30 — Secondary. TechBuzz on Vibes standalone testing: https://www.techbuzz.ai/articles/meta-spins-out-vibes-as-standalone-ai-video-app
- S31 — Secondary. NextAI Press on Vibes AI video tooling: https://medium.com/@nextaipress/vibes-ai-metas-free-tool-for-ai-video-creation-2026-78f6d57eefca
- S32 — Secondary. PetaPixel on premium subscription trials: https://petapixel.com/2026/01/27/meta-will-trial-premium-subscriptions-for-instagram-and-facebook/

- S33 — Secondary. Medium on Llama 3.3 70B benchmarks, GQA, pricing (Dec 2024): https://medium.com/@hdeep0100/unveiling-metas-llama-3-3-70b-redefining-efficiency-and-performance-in-ai-a1e76a30be3f
- S34 — Vendor. Meta Llama 3.3 70B Instruct HF README benchmark table: https://huggingface.co/meta-llama/Llama-3.3-70B-Instruct/resolve/main/README.md?download=true
- S35 — Vendor. meta-llama/llama-models Llama 3.3 model card (cutoff, benchmarks): https://github.com/meta-llama/llama-models/blob/main/models/llama3_3/MODEL_CARD.md
- S36 — Secondary. Cameron Wolfe on Llama 4 LMArena controversy and long-context retrieval: https://cameronrwolfe.substack.com/p/llama-4
- S37 — Secondary. Beebom on LMArena/Meta benchmark manipulation allegations: https://beebom.com/meta-llama-4-benchmark-manipulation-not-first-time/
- S38 — Secondary. Neowin on unmodified Maverick ranking below rivals: https://www.neowin.net/news/unmodified-llama-4-maverick-ranks-below-rivals-following-meta-cheating-allegations/
- S39 — Secondary. TokenMix on Behemoth status, scale, and LlamaCon 2026 (Apr 2026): https://tokenmix.ai/blog/llama-4-behemoth-still-training-2026
- S40 — Secondary. The Deep View on Llama 4 training (CO2, costs, 30T tokens): https://archive.thedeepview.com/p/meta-developed-new-techniques-to-train-llama-4-the-difference-is-minimal
- S41 — Secondary. FxVerify/WSJ on Behemoth delay, attrition, $72B capex: https://fxverify.com/ja/news/meta-delays-rollout-of-flagship-ai-model-amid-internal-performance-concerns-11421
- S42 — Community/secondary. llama_blog citing Axios on Behemoth delay (May 2025): https://github.com/mattscar/llama_blog/blob/HEAD/src/content/blog/20250520_behemoth_delay.md
- S43 — Secondary. prompthackers on Llama 3.3 pricing and benchmarks: https://www.prompthackers.co/compare/llama-3.3-70b/gpt-4

