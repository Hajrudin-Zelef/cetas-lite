---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-25
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "AWS", "Anthropic", "Fireworks AI", "Google", "Hugging Face", "Intel", "Microsoft", "Nvidia", "OpenAI", "Oracle", "Qualcomm"]
dates: ["2025-10", "2026-04-27", "2026-06-02", "2026-07", "2026-07-29", "2026-08"]
keywords: ["agent", "agentic", "agents", "agi", "amd", "aws", "bedrock", "benchmarks", "chatgpt", "claude", "copilot", "fine-tuning"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [909, 945]
sha256: dba5fd2457e318c79109a3826b176c6bfd563c82676523cf9335168ac8f1fc09
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

- **Phi-4-reasoning-vision-15B** released **March 4–5, 2026**: 15B-parameter open-weight multimodal (vision-language) model. Hybrid reasoning modes ('hybrid', 'think', 'no-think') that choose when to use chain-of-thought; trained efficiently on 240 B200 GPUs over 4 days; 200B tokens. Benchmarks: **75.2 MathVista-MINI, 54.3 MMMU-VAL** (vendor-reported). Available on Hugging Face / Azure AI Foundry [secondary] https://www.neowin.net/news/microsoft-releases-phi-4-15b-an-open-weight-ai-model-that-chooses-when-to-think/.
- **Earlier Phi-4 family (timeline context, pre-window):** Phi-4 (14B, Dec 2024), Phi-4-mini (3.8B, Feb 2025), Phi-4-multimodal-instruct (5.6B, Feb 2025), Phi-4-reasoning / reasoning-plus / mini-reasoning (Apr 2025) — all MIT-licensed. **No Phi-5 announcement found in window; no Phi-5 release in 2026 confirmed.**
- **Phi lead personnel (context):** Sébastien Bubeck, key Phi researcher, left Microsoft for OpenAI (reported Dec 2025) [secondary].
- **Project Polaris (secondary only):** mirrors claim Microsoft announced Project Polaris, an in-house mixture-of-experts coding model to replace GPT-4 Turbo as the default GitHub Copilot model starting August 2026, with a 3-month fallback window; runs on Maia accelerators; companion fine-tuning service "Turing Forge"; claimed to outperform GPT-4 Turbo on HumanEval and MBPP [unverified][secondary]. Not corroborated by an official Microsoft announcement in collected sources.

### 6.7 On-device models: Microsoft Aion 1.0 / Copilot+ PC pivot (Build 2026)

- **Aion 1.0** announced **June 2, 2026** at Build: on-device small language model family for Windows 11. **Aion 1.0 Instruct** (preview at launch): small SLM for summarization and rewriting. **Aion 1.0 Plan**: 14B reasoning and tool-calling model, 32K context, "coming in the following months." Open weights scheduled for **Hugging Face in July 2026** (whether shipped by Sept 22 not confirmed) [secondary] https://runaihome.com/blog/microsoft-aion-1-0-windows-local-ai-copilot-pc-2026/.
- Hardware: NPU path requires Copilot+ PC with ≥40 TOPS NPU (Qualcomm Snapdragon X Elite 45 TOPS; Intel Lunar Lake 45–48 TOPS); AMD Ryzen AI (50 TOPS) qualified on hardware but Aion NPU support deferred; runs via Windows ML / Windows AI Foundry; also runs on discrete GPUs [secondary].
- **Strategic pivot at Build 2026:** Microsoft de-emphasized the Copilot+ PC brand. CEO Satya Nadella told developers "you now have the full scope of GPUs that you can get to" when writing AI software for Windows ML; the agentic future of Windows was presented as running locally across a wider device range, including NVIDIA-powered hardware — not requiring a Copilot+ PC NPU [secondary].
- **New hardware at Build 2026:** Surface Laptop Ultra and **Surface RTX Spark Dev Box** (NVIDIA RTX Spark chip, targeting professional AI workloads well above 40 TOPS) [secondary].

### 6.8 Microsoft Foundry / Azure AI (2026)

- **Foundry catalog scale:** 11,000+ models from multiple providers per Nadella at Build 2026 (Foundry documentation counts >1,900 core models; counts above 10,000 in third-party tallies include open-weight variants) [official][secondary]. Catalog size figures vary by counting method.
- **Fireworks AI generally available on Foundry** (announced June 2, 2026): single-platform experience with enterprise governance and Azure data residency [official].
- **Frontier Tuning** (private preview, June 2, 2026): applies RL within the customer's compliance boundary to tune agents on their own data/workflows [official].
- **Anthropic models on Foundry:** Claude Sonnet 4.5, Haiku 4.5, and Opus 4.1 available in public preview in Microsoft Foundry with Azure billing (per Anthropic announcement; reported in Aug 2026 coverage) [secondary].
- **Foundry new models added (2026 digest):** NVIDIA Nemotron-3-Super-120B-A12B, IBM Granite-4.0-1b-Speech, Sarvam-105B; voice-native agents in public preview; user-scoped persistent memory for agents [secondary].
- **Build 2026 platform announcements:** Foundry Agent Service deep-dives; **Microsoft IQ context layer** (Work IQ, Web IQ, Foundry IQ); **Agent 365** governance SDK (Entra/Defender extended to agents); **Azure HorizonDB** (GPU-accelerated Fabric); **Project Solara** (hardware) [secondary].

### 6.9 Microsoft–OpenAI relationship — detailed 2026 record

**October 2025 restructuring (context):** OpenAI restructured into a public-benefit corporation (OpenAI Group PBC) valued at **$500B**; Microsoft holds ~**27%** stake (~$135B) on $13.8B invested — ~10× return [independent — Al Jazeera, Reuters].

**April 27, 2026 — commercial terms renegotiated (joint announcement):**
- **Microsoft is no longer the exclusive licensee** of OpenAI models; OpenAI may serve products on AWS, Google Cloud, Oracle; Azure remains "primary cloud partner" — products ship first on Azure "unless Microsoft cannot and chooses not to support the necessary capabilities" [secondary; AP via SRN News].
- **Revenue share:** Microsoft no longer pays OpenAI a share on Azure-resold AI products (eliminated). OpenAI continues paying Microsoft a share through 2030 — 20%, now subject to a fixed total cap (cap amount undisclosed) [secondary].
- **AGI clause removed:** the provision that would have terminated Microsoft's IP rights upon OpenAI's board declaring AGI was eliminated; Microsoft's IP license now non-exclusive with a hard expiration in **2032** [secondary].
- **Equity unchanged:** 27% stake retained [secondary].
- **Immediate consequence:** within ~24 hours, OpenAI models (GPT-5.5, GPT-5.4 preview; Codex limited preview) became available on **Amazon Bedrock**; Amazon separately disclosed a planned **$50B** investment in OpenAI [secondary] https://srnnews.com/microsoft-cuts-openai-revenue-share-in-a-fresh-step-to-loosen-their-ai-alliance/.
- Wedbush analyst Dan Ives: agreement "puts OpenAI on a strong path forward to going public through IPO" (reportedly targeted late 2026); also important for Microsoft's "tech independence" in Copilot [secondary].

**OpenAI fundraising context in 2026:** Feb 27, 2026 — OpenAI announced a **$110B funding round** ($50B Amazon + $30B SoftBank + $30B Nvidia). Pre-money valuation reported as $730B (FT/GulfNews) vs. $840B (aiforautomation) — inconsistent [secondary]. OpenAI subsequently closed a **$122B round at $852B post-money** [secondary]; as of **Sept 16, 2026**, FT reports talks of another round at ~**$1.2T** valuation ahead of a possible IPO [secondary via Reuters]. ChatGPT approaching **1B weekly active users** (reported July 29, 2026) [secondary].

### 6.10 Microsoft investments, personnel, infrastructure (2026)

