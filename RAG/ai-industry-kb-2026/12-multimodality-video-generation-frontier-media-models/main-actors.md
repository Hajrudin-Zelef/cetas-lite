---
id: ai-industry-kb-2026/12-multimodality-video-generation-frontier-media-models/main-actors
title: "Main actors"
domain: multimodality-video-generation-frontier-media-models
role: deep-dive
task: multimodal
actors: ["Apple", "ByteDance", "China", "EU", "Google", "Hugging Face", "Meta", "MiniMax", "OpenAI", "United States", "xAI"]
dates: ["2025-08-05", "2025-09-30", "2025-12-16", "2026-01-29", "2026-02-05", "2026-02-12", "2026-03-24", "2026-03-29", "2026-04", "2026-04-26", "2026-05", "2026-05-12", "2026-05-19", "2026-05-21", "2026-06-23", "2026-06-30", "2026-07-01", "2026-07-02", "2026-07-31", "2026-08-02", "2026-08-05", "2026-08-31", "2026-09", "2026-09-02", "2026-09-11", "2026-09-15", "2026-09-17", "2026-09-22", "2026-09-24"]
keywords: ["agent", "agentic", "agents", "arr", "attribution", "benchmark", "benchmarks", "chatgpt", "compute", "consumer", "copyright", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6377, 6504]
section: "12. Multimodality — Video Generation & Frontier Media Models"
sha256: 120954ec7d93670f6a36aab424657e59040cb24fbf9f15f8d5b597ce2082e1de
---

# Main actors

## Main actors

- **OpenAI:** Sora 2 (consumer product discontinued 2026-04-26, API cutoff 2026-09-24 planned; model survives in ChatGPT paid tiers; team → world-simulation research); GPT Image 1.5 family (native image generation); DALL-E 2/3 retired 2026-05-12.
- **Google / DeepMind:** Gemini Omni Flash (any-to-any video benchmark-setter, $0.10/s); Veo 3.1 (sparse verification — distributed, no surfaced 2026 launch record); Nano Banana image family (2 Lite entry tier ~2026-07-01/02; Pro via Gemini 3 Pro Image); Project Genie (2026 consumer beta + Street View grounding; no Genie 4); Agentic Video Understanding; SynthID watermarking and C2PA credentials by default on Omni Flash.
- **ByteDance Seed:** Seedance 2.0 (China launch 2026-02-12) and 2.5 (announced 2026-06-23, Dreamina 2026-07-31) — 30s single-pass, native 4K claims [VENDOR/secondary]; international launch paused over copyright disputes.
- **Kuaishou:** Kling 3.0 (global launch 2026-02-05) and Kling 3.0 Omni (O3) — 4K60, 6-shot storyboarding, 5-language lip sync, Motion Control AI; reported $240M ARR Dec 2025 [UNVERIFIED]; the revenue counterpoint to Sora.
- **Runway:** Gen-4.5 — commercially distributed via Adobe Firefly; claimed beaten by Seedance 2.0 on AA at launch [UNVERIFIED].
- **MiniMax:** H3 — open-weight video leader on at least one public ranking under a restrictive community license (US/EU/UK/SK excluded from local deployment; → §11).
- **HiDream.ai (Beijing):** HiDream-O1-Video-1.0 — company-reported native omnimodal video entrant (2026-09-17) [VENDOR].
- **Adobe Firefly:** aggregation model — 30+ video models by April 2026, rents inference, sidesteps the training-compute bill.
- **xAI:** Grok Imagine (image/video/voice via separate Imagine API) → §1; assembly-not-any-to-any stack.
- **Meta Movie Gen:** no primary material surfaced in this pass sufficient to add — deliberately omitted.
- **Dreamina (ByteDance):** consumer distribution surface for Seedance 2.5's global launch (2026-07-31) — the venue, not a separate model; international rollout is where the copyright-gating question applies.
- **Novita AI:** third-party API distributor documenting Kling 3.0's global launch (2026-02-05) and its per-second pricing — representative of the B2B distribution layer that made Kling's economics work.
- **Hypereal/Hyperreal:** commercial PR distributor for Seedance 2.0 API availability — representative of the commercial-secondary layer behind ByteDance's API pricing claims.
- **fal:** post-trained H3 Max variant (5s 768p in under 3s; 75% launch discount expired 2026-09-15, price $0.10→$0.40) — the post-training-as-a-service layer on top of open-weight video models.
- **the-decoder:** independent secondary outlet whose reporting made H3 "the first open model to top an AI video ranking" — the single best-corroborated leaderboard claim in this section.
- **ngram.com / CIOL:** analytical voices on the Sora shutdown — ngram's deprecation-page analysis (empty "replacement" column), CIOL's cost/revenue figures [UNVERIFIED] and the "side quests" reporting chain.
- **The likeness-rights estates (MLK Jr., Robin Williams):** not labs, but active constraints — their complaints shaped Sora's moderation posture and remain the concrete case law-adjacent evidence that likeness moderation is a product-cost factor in video generation.
- **Waymo** [UNVERIFIED connection]: claimed user of a fine-tuned Genie variant for driving simulation — secondary blogs only, no primary confirmation; included because it illustrates the world-model-to-simulation use case Genie is positioned for.

- **OpenAI Applications org (Fidji Simo):** the internal actor behind the Sora shutdown call — Simo's "side quests" memo to employees and Altman's redirection decision; the org-level evidence that video was deprioritized by the applications business, not by research.
- **Disney:** the reported $1B franchise-content partnership affected by the shutdown (WSJ: informed shortly before the announcement, no money changed hands) — the enterprise-partnership casualty that shows shutdown blast radius beyond consumers.
- **Volcano Engine (ByteDance):** the announcement and distribution venue for Seedance 2.0 (Feb 2026 launch tracking) and Seedance 2.5 (FORCE conference 2026-06-23) — ByteDance's cloud/API arm as the primary-source channel, in the absence of ByteDance-primary English materials.
- **Google Labs:** the distribution channel for Project Genie's consumer beta (2026-01-29) — Labs-as-launch-vehicle for the $249.99/mo-gated experiment.
- **C2PA / SynthID provenance standards:** Omni Flash shipped with C2PA credentials and SynthID watermarks on by default — provenance-by-default as a 2026 product norm for flagship video models.
- **ABNewswire (PR syndication):** the syndication channel for Kuaishou's September 2026 "AI Director" PR wave — representative of how company positioning becomes "industry trend" coverage in this domain (cf. the wave3 narrative audit, out of scope here).
- **humai.blog:** single-source origin of the "Sora resources redirected to Spud" claim [UNVERIFIED] — recorded as an actor only to pin the provenance of that one claim.
- **Novita AI and Hypereal/Hyperreal:** the third-party API-distribution layer (Kling 3.0 pricing/launch documentation; Seedance 2.0 API PR) — the commercial-secondary layer that makes Chinese labs' models globally consumable and is the actual source of most public per-second pricing.
- **Google AI Ultra ($249.99/mo):** not a model but a business model — the subscription gate behind Project Genie's consumer beta and its May 2026 global expansion; the price point at which interactive world models are currently monetized (two orders of magnitude above typical chatbot tiers).

## Timeline and context

### 12.13 Pre-2026 anchors (context, not 2026 events)

- **2025-08-05:** Google DeepMind announces Genie 3 (interactive world model; 720p/24fps; promptable world events).
- **2025-09-30:** OpenAI launches Sora 2 as a standalone iOS app (synchronized audio, Cameos, Sora 2 Pro; US/Canada first). Not 2026.
- **2025-12-16:** gpt-image-1.5 GA (~4x faster, 20% cheaper, RGBA-native).
- **2025-12:** Kling reportedly at $240M ARR [UNVERIFIED, single secondary compilation].

### 12.14 Dated release timeline (Feb → Sep 2026)

| Date (2026) | Event |
|---|---|
| 2026-01-29 | **Project Genie** consumer beta via Google Labs, gated on Google AI Ultra ($249.99/mo) |
| 2026-02-05 | **Kling 3.0** global launch (15s clips, native multilingual audio) |
| 2026-02-12 | **Seedance 2.0** China launch (quad-modal refs, native audio, 2K) |
| 2026-02 (month) | Last month with all three consumer video products live: Sora (still up), Kling 3.0 (new), Seedance 2.0 (China) |
| 2026-03-24 | **Sora shutdown announced** (official Sora account); Sora 2 API deprecation notice to developers |
| 2026-04 (month) | **Adobe Firefly passes 30+ video models** (Veo 3.1, Runway Gen-4.5, Kling 3.0 distributed) — aggregation as a business model |
| 2026-04-26 | **Sora web/app discontinued** [VERIFIED — OpenAI Help Center] |
| 2026-08-31 | OpenAI Help Center discontinuation article updated (~2026-08-31) — confirms 2026-04-26 web/app death, 2026-09-24 API death, sora.chatgpt.com/sunset export, permanent data deletion |
| 2026-09-02 | Vendor-exit migration guide (dev.to) retrieved quoting OpenAI's provider-owned discontinuation notice — API end 2026-09-24 restated |
| 2026-05-12 | **DALL-E 2/3 sunset** — pipeline-era generation retired for GPT Image family |
| 2026-05-19 | **Google I/O 2026** — Gemini Omni (Flash) unveiled ("Nano Banana for video") |
| 2026-05-21 | **Project Genie expansion** — Street View/Maps Imagery Grounding (US first), broader global AI Ultra access |
| 2026-06-23 | **Seedance 2.5 announced** at Volcano Engine FORCE (30s single-pass, native 4K, 50 references [VENDOR/secondary]) |
| 2026-06-30 | **Gemini Omni Flash API GA** ($0.10/s 720p; $1.50/M input tokens) |
| ~2026-07-01/02 | **Gemini Omni Flash developer rollout + Nano Banana 2 Lite launch** (corrected dates — not I/O) |
| 2026-07-31 | **Seedance 2.5 global launch on Dreamina**; MiniMax H3 API launch (Hailuo) |
| 2026-08-02/03 | **MiniMax H3 open weights** land on Hugging Face (sources differ by one day) |
| 2026-08 (month) | **StreamArena** released (streaming video-agent benchmark) |
| 2026-09 (~early) | **Google Agentic Video Understanding** announced (−88% tokens, +~7% accuracy) |
| 2026-09-11/12 | **Kling 3.0 "AI Director" PR wave** — repositioning of the February model, not a new release |
| 2026-09-15 | fal H3 Max 75% launch discount expires ($0.10→$0.40 per 5s clip) |
| 2026-09-17 | **HiDream-O1-Video-1.0 announced** (native omnimodal video, company-reported) |
| 2026-09-24 | **Sora API permanent shutdown** — PLANNED, future relative to 2026-09-22; reverification required |

### 12.15 What 2026 changed (context)

- The bottleneck moved from quality to unit economics: the only consumer video product that died was the one with the worst reported economics.
- The commercial leaders are ByteDance (Seedance), Kuaishou (Kling), and Google (Veo/Genie) — two Chinese labs and one US lab — while OpenAI, the 2024–2025 video flagship, exited the consumer product entirely.
- Generation commoditized into aggregation (Firefly, 30+ models) while the any-to-any frontier moved to single-pass synchronized audio-video (Omni Flash, H3, HiDream claims) and agentic perception (AVU) moved the economics of video *understanding*.
- **The "announced vs shipped vs PR" trichotomy** became the section's governing discipline: announced (Genie beta dates, Seedance 2.5 specs), shipped (Omni Flash API GA, H3 weights, Kling pricing via Novita), PR (Kling's September "AI Director" wave, HiDream's launch PR). Secondary roundups routinely promote PR to "launch" and announcements to "availability"; the dated-facts section above demotes each to its evidence level.
- **China-first launch sequencing** is now a structural feature of the video race: Seedance 2.0 (China-only, Feb 2026), Seedance 2.5 (Dreamina global, Jul 2026), Kling 3.0 (global from day one, Feb 2026), H3 (open weights global, Aug 2026, but license-restricted in the West). "Global launch" means different availability sets for different labs — the availability geography must be recorded alongside the date.
- **The 30s/4K/lip-sync bundle** (Seedance 2.5 claims, Kling 3.0 O3 PR) is the 2026 answer to "what does the next video model add": not a new architecture announcement, but longer single-pass duration, higher native resolution, and better audio fidelity — capability stacking on the existing diffusion-transformer substrate, marketed through repositioning waves rather than version numbers.

## Implications

1. **Technology leadership ≠ product viability.** Sora's 2026 exit is the strongest evidence the AI video market has produced in two years (ngram framing): the model was state-of-the-art; the consumer product was not viable. Every viability claim in this domain must be paired with per-second or per-task economics, not just benchmark positions.
2. **Three surviving business models:** B2B/API monetization with working unit economics (Kling, Seedance), aggregation of others' models (Adobe Firefly — rents inference, skips the training bill), and premium subscription gating (Project Genie at $249.99/mo). No surviving consumer video product in 2026 was free-at-scale.
3. **Any-to-any is the structural direction of generation.** Single-pass synchronized audio-video generation (Omni Flash shipped; H3 shipped open-weight; HiDream company-reported) replaces the separate-Foley-pipeline world; conversational multi-turn editing replaces detection→edit→re-render chains. But the shipped-vs-roadmap distinction matters: Omni Flash's audio-reference inputs and scene extension were not supported at preview despite the any-to-any headline.
4. **Agentic perception is the 2026 cost story for video understanding.** Google's AVU (−88% tokens, −66% cost, +~7% accuracy) treats compute allocation as a first-class product variable; benchmarks are following (StreamArena for streaming agents, CFD for caption-once/frames-on-demand). Expect procurement to be decided on accuracy-vs-cost Pareto frontiers, not raw accuracy.
5. **"Open" video now needs a license qualifier.** MiniMax H3 leads a public ranking while barring US/EU/UK/South Korea from local deployment — capability leadership and deployability leadership are different facts and must be stated separately (→ §11).
6. **Attribution hygiene stays material.** The three most misdated items in this section — Sora 2's launch (2025-09-30, not 2026), Nano Banana 2 Lite (early July, not I/O), Seedance 2.5 (June/July, not the week of H3's weights), Kling's September wave (repositioning, not a launch) — are all cases where secondary roundups conflated announcements with availability or PR with product. The 2026-09-24 Sora API cutoff is the one fact that post-dates the writing and must be reverified before any downstream claim that Sora's API is gone.
7. **Likeness and copyright moved from legal footnote to launch-gating variable.** Seedance 2.0's international pause and Sora's estate-driven likeness restrictions show moderation infrastructure is now a release blocker and a cost factor — not an afterthought — for consumer video products.
8. **World models vs video generators split into distinct product categories.** Genie (playable runtime worlds, no exportable assets) is a simulation surface, not a video generator; the Sora team's pivot to a world-simulation research unit confirms the category split from the vendor side too. Conflating the two inflates "video model" timelines.
9. **Resolution and clip length are now pricing levers, not just spec rows.** Omni Flash at $0.10/s (720p) vs Veo 3.1 Standard at $0.40/s (1080p), and Kling's with/without-audio price split ($0.168/s vs $0.252/s Standard), show the 2026 price card is tiered by output tier — cost modeling must use the tiered price, not a single $/s figure.
10. **The assembly-vs-native verdict is layer-specific, not absolute.** Native single-pass generation won the core perception-generation loop (DALL-E sunset, Omni Flash vs Foley pipelines); modular assembly survives at product-integration (xAI's separate backends, ColPali/CLAP/SigLIP RAG heads) and where it is rationally cheaper. "The bricolage ended" should never be stated without the layer qualifier.
11. **Provenance-by-default is now a launch norm.** Omni Flash shipped with C2PA credentials and SynthID watermarks on by default; likeness estates are actively shaping product policy. Any 2026-era video-model capability comparison that ignores the moderation/provenance layer is missing a release-blocking variable.
12. **Post-training and distribution layers capture open-weight video value.** H3's capability is open-weight; its commercial realization runs through fal's post-trained H3 Max (3s renders) and community quants — the value accrues to the distribution layer, while the restrictive license bars Western-lab local deployment. Open weights ≠ open deployment.
13. **Price parity is a positioning signal, not just a number.** Omni Flash's $0.10/s matching Veo 3.1 Fast, and Nano Banana 2 Lite's commodity pricing vs the Pro tier, show 2026 labs using price cards to declare architectural succession (any-to-any as default, entry tiers as commodity) — read the price card as product strategy.
14. **Clip length is the next pricing battlefield.** Seedance 2.5's 30s single-pass and per-clip pricing (~$0.14/15s) vs per-second API cards (Kling, Omni Flash) are two incompatible billing units coexisting in 2026 — which unit wins determines how the 30s/4K capability bundle gets monetized. A 2026–2027 watch item.

## Sources and URLs

- Sora 2 launch (2025-09-30): https://venturebeat.com/ai/openai-debuts-sora-2-ai-video-generator-app-with-sound-and-self-insertion
- [PRIMARY] Sora discontinuation (Help Center, VERIFIED): https://help.openai.com/en/articles/20001152-what-to-know-about-the-sora-discontinuation
- Sora shutdown economics analysis (secondary): https://www.ngram.com/blog/openai-sora-shutdown-ai-video-economics
- Sora shutdown reporting (secondary; $5.4B vs $2.1M figure UNVERIFIED): https://www.ciol.com/tech-buzz/openai-shuts-sora-video-app-costs-usage-decline-11436718
- Sora shutdown analysis (secondary): https://www.devflokers.com/blog/openai-sora-shutdown-reasons-whats-next-2026
- Sora API shutdown migration guide (secondary; quotes OpenAI discontinuation notice, API end 2026-09-24): https://dev.to/lucas_apimart/openai-sora-api-shutdown-production-migration-guide-2ho1
- Sora 2 API limits and shutdown timeline (secondary, commercial): https://www.glbgpt.com/hub/openai-sora-2-availability/
- Genie 3 baseline + Project Genie consumer beta 2026-01-29 (secondary): https://soravideo.art/blog/genie-3-google-deepmind-world-model
- Project Genie 2026-05-21 Street View expansion (secondary): https://pulse2.com/google-deepmind-project-genie-expands-with-street-view-powered-real-world-simulation/
- Genie 3 explainer (secondary): https://yourstory.com/ai-story/deepmind-genie-3-ai-world-model
- Seedance timeline 2.0/2.5 (COMMUNITY README): https://github.com/ikaijua/awesome-aitools/blob/HEAD/docs/seedance/README.md
- Seedance 2.0 API availability (commercial PR, secondary): https://www.feedlotmagazine.com/online_features/press_releases/how-to-use-seedance-2-0-api-hypereal-ai-now-offers-bytedance-s-most-powerful/article_413a3db8-7eba-5048-a3c1-11089878d982.html
- Seedance 2.0 analysis (COMMUNITY blog; AA #1 claim UNVERIFIED): https://github.com/heyuan110/heyuan110.github.io/blob/HEAD/content/posts/ai/2026-03-29-seedance-2-bytedance-ai-video/index.md
- AI video generation tools state, Aug 2026 (deep-research compilation, SECONDARY): https://github.com/baditaflorin/0deepresearch/blob/HEAD/content/posts/2026-08-05-state-of-ai-video-generation-tools.md
- Kling 3.0 global launch 2026-02-05 + API pricing (secondary): https://medium.com/@marketing_novita.ai/kling-3-0-now-on-novita-ai-cinematic-ai-video-generation-at-scale-82aee1eeca15
- Kling 3.0 "AI Director" Sept 2026 PR wave (COMPANY PR): https://investor.wedbush.com/wedbush/article/abnewswire-2026-9-12-kling-ai-launches-kling-30-kuaishous-ai-director-platform-delivers-native-4k-60fps-video-with-multi-shot-storyboarding-motion-control-and-5-language-lip-sync
- Gemini Omni Flash + Nano Banana 2 Lite alerts (secondary): http://aiweekly.co/alerts/google-launches-gemini-omni-flash-and-nano-banana-2-lite
- Omni Flash / Nano Banana 2 Lite developer launch (secondary): https://m.dailyhunt.in/news/india/english/the+mobile+indian+english-epaper-mblinden/google+launches+nano+banana+2+lite+gemini+omni+flash+for+developers+also+announced-newsid-n718034876
- Nano Banana 2 Lite + Omni Flash dev rollout ~2026-07-02 (secondary; corrects I/O misdating): https://theaiinsider.tech/2026/07/02/google-expands-ai-media-and-agent-capabilities-with-nano-banana-2-lite-gemini-omni-flash-and-spark-for-mac/
- Google I/O 2026 live blog (primary-adjacent; keynote 2026-05-19): https://www.livemint.com/technology/tech-news/google-io-2026-live-updates-sundar-pichai-keynote-gemini-ai-android-xr-launches-announcements-highlights-19-may-2026-11779199501541.html
- Gemini Omni Flash / Nano Banana 2 Lite overview (secondary): https://trypencil.com/blog/articles/gemini-omni-flash-nano-banana-2-lite
- Nano Banana 2 Lite release news (secondary): https://completeaitraining.com/news/google-releases-nano-banana-2-lite-for-image-generation-and/
- Nano Banana 2 Lite pricing $0.034/1K images (secondary): http://gigazine.net/gsc_news/en/20260701-google-nano-banana-2-lite-gemini-omni-flash/
- HiDream-O1-Video-1.0 launch PR (VENDOR): https://www.channelnewsasia.com/media-release/hidream-unveils-hidream-o1-video-10-native-omnimodal-video-model-built-physical-consistency-6391321
- HiDream-O1-Video-1.0 launch PR (VENDOR): https://www.malaymail.com/news/money/mediaoutreach/2026/09/17/hidream-unveils-hidream-o1-video-10-a-native-omnimodal-video-model-built-for-physical-consistency/488140
- MiniMax H3 first open model to top an AI video ranking (secondary): https://the-decoder.com/chinas-minimax-h3-is-the-first-open-model-to-top-an-ai-video-ranking/
- MiniMax H3 overview (secondary): https://explainx.ai/blog/minimax-h3-open-video-model-hailuo-july-2026
- MiniMax H3 local hardware guide (secondary): https://runaihome.com/blog/minimax-h3-open-weights-local-ai-hardware-guide-2026/
- MiniMax H3 local setup guide (secondary): https://localaimaster.com/blog/minimax-h3-local-setup-guide
- MiniMax H3 architecture/modes (COMMUNITY skill card): https://github.com/ryannel/skills/blob/HEAD/skills/generative-media/minimax-h3/SKILL.md
- fal H3 Max speed/pricing, Sept 15 discount expiry (secondary): https://startupfortune.com/minimaxs-h3-max-video-model-topped-sora-and-veo-then-its-price-quadrupled/

