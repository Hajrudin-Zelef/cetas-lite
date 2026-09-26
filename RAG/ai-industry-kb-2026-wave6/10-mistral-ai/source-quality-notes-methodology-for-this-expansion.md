---
id: ai-industry-kb-2026-wave6/10-mistral-ai/source-quality-notes-methodology-for-this-expansion
title: "Source-quality notes (methodology for this expansion)"
domain: mistral-ai
role: deep-dive
task: actor-profile
actors: ["Hugging Face", "Mistral", "Nvidia", "Samsung"]
dates: ["2025-07", "2025-12-02", "2025-12-09", "2026-05-22", "2026-08-31"]
keywords: ["agent", "agents", "apache", "arr", "benchmarks", "funding", "gpu", "gpus", "gqa", "inference", "intel", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5104, 5148]
section: "§10. Mistral AI"
delta_of: ai-industry-kb-2026
sha256: d5c3805baa4abc1b1098246779372c519530ca8b6a9b9aa5b7c5f8b11e2ee8c3
---

# Source-quality notes (methodology for this expansion)

- S43 — Secondary. VentureBeat on Devstral 2 launch (123B Modified MIT, verbatim $20M clause, Small 2 Apache 2.0): https://venturebeat.com/ai/mistral-launches-powerful-devstral-2-coding-model-including-open-source
- S44 — Secondary. TechCrunch on Devstral 2 + Vibe CLI (2025-12-09, pricing, ≥4 H100s): https://techcrunch.com/2025/12/09/mistral-ai-surfs-vibe-coding-tailwinds-with-new-coding-models/
- S45 — Secondary. Geeky Gadgets on Devstral 2 overview (specs, Vibe limits): https://www.geeky-gadgets.com/mistral-devstral-2-overview/
- S46 — Secondary. SiliconANGLE on Devstral 2 pricing and licensing: https://www.siliconangle.com/2025/12/09/mistral-ais-devstral-2-open-weights-vibe-coding-model-built-rival-best-proprietary-systems/
- S47 — Community/secondary. dev.to developer guide on Large 3 (specs, benchmarks, pricing, 3000 H200s): https://dev.to/jangwook_kim_e31e7291ad98/mistral-large-3-the-675b-open-weight-moe-model-developer-guide-250a
- S48 — Secondary. TechCrunch on Mistral 3 family launch (2025-12-02, Ministral 3 nine models): https://techcrunch.com/2025/12/02/mistral-closes-in-on-big-ai-rivals-with-mistral-3-open-weight-frontier-and-small-models/
- S49 — Community. agentone free-tier research on Large 3 (Elo, self-hosting, free tiers): https://github.com/dirk-makerhafen/agentone/blob/HEAD/docs/free-tier-research/models/mistral-large-3.md

- S50 — Secondary. The AI Track on Magistral launch (London Tech Week, Small/Medium, Macron backing): https://theaitrack.com/mistral-magistral-europe-reasoning-ai-launch/
- S51 — Secondary. AIRevolution on Magistral announcement (benchmarks, multilingual, 10× throughput): http://airevolution.poltextlab.com/mistral-ai-unveils-its-first-reasoning-model-10x-faster-than-competitors/
- S52 — Secondary. VentureBeat on Magistral 1.2 updates (vision, benchmarks, API IDs): https://venturebeat.com/ai/mistrals-updated-magistral-small-1-2-reasoning-model-can-analyze-images-and
- S53 — Secondary. dev.ua on Magistral Small/Medium 1.2 (quantized on RTX 4090/MacBook): https://dev.ua/en/news/mistral-1758260437
- S54 — Community. AllAboutAI on Magistral multimodal benchmarks (MMMU, MathVista): https://www.allaboutai.com/ai-agents/magistral/
- S55 — Secondary. FourWeekMBA on Series D (€3B, >€21B, Samsung led, syndicate): https://fourweekmba.com/ai-mistral-series-d-samsung-nvidia-asml-sovereign-ai/
- S56 — Secondary. HPCwire AIwire on Series D (frontier research, infrastructure, 125+ customers): https://www.hpcwire.com/aiwire/2026/09/08/mistral-raises-e3b-to-expand-ai-research-and-infrastructure/
- S57 — Secondary. TechBriefly on Series D ($830M debt, 1GW by 2030, Luxembourg): https://techbriefly.com/2026/09/08/mistral-ai-raises-3-billion-samsung-led-funding-round/

- S37 — Community. Braintrust proxy issue (Medium 3/3.1 deprecated 2026-05-22, retired 2026-08-31, Medium 3.5 replacement): https://github.com/braintrustdata/braintrust-proxy/issues/922
- S38 — Community. Aitooltier sweep on Medium 3.1 retirement 2026-08-31: https://github.com/cmcmath84/aitooltier/commit/1ae148ac76919ecedfa633b1ef76471251a58e5a
- S39 — Secondary. The-Decoder on Medium 3.5 (SWE 77.6, T3-Telecom 91.4, replaces 3.1/Magistral/Devstral 2, Modified MIT): https://the-decoder.com/mistrals-new-flagship-medium-3-5-folds-chat-reasoning-and-code-into-one-model/
- S40 — Secondary. AI Dev Weekly on Medium 3.5 (Apr 29, dense 128B, 256K, Vibe remote agents): https://medium.com/@meulenjoske/ai-dev-weekly-8-mistral-medium-3-5-2bc7c0a7fcc7
- S41 — Vendor-adjacent. NVIDIA HF page for Mistral-Medium-3.5-128B-NVFP4 (replaces 3.1/Magistral, Modified MIT): https://huggingface.co/nvidia/Mistral-Medium-3.5-128B-NVFP4
- S42 — Secondary. Medium on Mistral Large 3 (Dec 2 2025, 675B/41B, 128 experts/layer, 256K): https://medium.com/@chiwai.kiriba/why-mistral-is-the-future-of-open-weight-intelligence-5fb967963643
- S43 — Vendor-adjacent. NVIDIA/Marktechpost on Large 3 inference (128 experts/layer, NVFP4, GB200 NVL72): https://www.marktechpost.com/2025/12/02/nvidia-and-mistral-ai-bring-10x-faster-inference-for-the-mistral-3-family-on-gb200-nvl72-gpu-systems/
- S44 — Secondary. Medium review of Mistral Large 3 (675B, ~40B active, ~3,000 H200 GPUs, AA Index placement): https://medium.com/@leucopsis/mistral-large-3-2512-review-7788c779a5e4
- S45 — Community. Red Hat deep dive on Mistral (Large 3 Apache 2.0, deployment, limitations): https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/mistral-ai-deep-dive.md
- S46 — Community. working-set research on Medium 3.5 (88 layers, GQA 12:1, dense 128B, config-derived params): https://github.com/t0msilver/working-set/blob/HEAD/research/model_mistral_medium35.md

- S45 — Community. Red Hat deep dive on Mistral AI (timeline, Emmi/Koyeb M&A, product specs, ARR): https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/mistral-ai-deep-dive.md
- S47 — Secondary. CRN on Mistral $830M debt for 13,800 GB300s (44 MW, Paris): https://www.crn.com/news/data-center/2026/mistral-ai-raises-830m-for-nvidia-data-center-with-13-800-gpus
- S48 — Secondary. Blockonomi on Mistral debt financing (seven banks, Bruyeres-le-Chatel, Q2 2026): https://blockonomi.com/mistral-ai-secures-830m-debt-financing-to-power-new-paris-data-center-with-nvidia-gpus/
- S49 — Secondary. Pulse on Mistral acquiring Emmi AI (Austrian LEM startup, physics simulation): https://www.pulse.bot/ai/deals/mistral-to-acquire-austrian-ai-startup-emmi-ai-78e5ec6b-19c3-4eb3-af90-c1df771f7d34/
- S50 — Secondary. AlternativeStack on Devstral Medium/Small 1.1 (61.6%/53.6% SWE-Bench): https://alternativestack.com/news/mistral-doubles-down-on-code-ai-devstral-medium-open-source-small-11-challenge-giants
- S51 — Secondary. MLQ on Devstral 2/Vibe CLI (53.6%/61.6% lineage, enterprise positioning): https://mlq.ai/news/mistral-ai-launches-devstral-2-coding-model-and-mistral-vibe-cli-for-vibe-coding-workflows/
- S52 — Secondary. Vocal.media on Devstral Medium (61.6% SWE-Bench, July 2025): https://vocal.media/futurism/how-devstral-s-ai-coding-agent-is-revolutionizing-development-in-2025-and-how-you-can-use-it
- S53 — Secondary. Onmine on Devstral launch (24B, Apache 2.0, $0.10/$0.30, OpenHands): https://onmine.io/mistral-ai-launches-devstral-powerful-new-open-source-swe-agent-model-that-runs-on-laptops/
- S54 — Secondary. Neurosignal on Devstral (finetuned from Small 3.1, 128K ctx, dogfooding): https://neurosignal.tech/mistral-ai-launches-devstral-powerful-new-open-source-swe-agent-model-that/

### Source-quality notes (methodology for this expansion)
- S31 (Awesome-Mistral) is a community-maintained catalog and the sole source for several retirement dates and Hugging Face repository names; retirement dates are described as published API dates from Mistral's official catalog.
- Medium 3.5's 77.6 SWE-Bench Verified figure appears only in S5 and is carried as UNVERIFIED.
- The Modified MIT verbatim carve-out is quoted from the LICENSE file via SingularityByte and corroborated by a battlecard citing Mistral's help article; the $20M/month figure appears in both. [SECONDARY, S36][SECONDARY, S41]
- Series D figures (€3B, >€21B) are corroborated across S13–S16 with consistent investor lists; the "largest European tech equity round" characterization is Mistral's own via S15.
- OCR 4's win-rate and OlmOCRBench figures are vendor claims via S35, with Mistral's own "directional rather than definitive" caveat preserved.

