---
id: labs-hyperscalers-2026/00-labs-hyperscalers/6-microsoft
title: "§6 — MICROSOFT"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["Anthropic", "Baseten", "EU", "Fireworks AI", "Hugging Face", "Microsoft", "Mistral", "OpenAI", "OpenRouter"]
dates: ["2026-03", "2026-04", "2026-05", "2026-06", "2026-06-02", "2026-07", "2026-08"]
keywords: ["accelerator", "agent", "agentic", "agents", "agi", "benchmark", "benchmarks", "chatgpt", "claude", "compute", "copilot", "cost"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [868, 913]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 791493bd9f342595429c3834345940af898ed2ef9f5d5119fe562637b41272d1
---

# §6 — MICROSOFT

## §6 — MICROSOFT

> **2026 at a glance — Microsoft:** the seven-model MAI family (Jun 2, Build) declared independence from OpenAI; Microsoft–OpenAI restructured (Apr 27, exclusivity ends, ~27% equity); multibillion-dollar Mistral EU compute deal (Jul 21); Maia silicon and the Aion supercomputing program; Foundry as the multi-model router.

### 6.1 MAI model family — seven models in 2026
- Microsoft released **seven "MAI" (Microsoft AI) models** during the window — its first serious in-house frontier-model push, reducing dependence on OpenAI. [secondary: Track B]
- **MAI-Voice-1 / MAI-1-preview** lineage (2025) extended with new 2026 releases spanning text, reasoning, and multimodal. [secondary]
- **Phi-4-reasoning-vision-15B** — open-weights reasoning + vision small model (15B params), notable for the open ecosystem. [secondary]

### 6.2 "Aion" & Foundry
- **Aion** — Microsoft's AI supercomputing / infrastructure program referenced in Track B (naming [secondary]).
- **Azure AI Foundry** — continued as Microsoft's model-router/agent platform; added third-party frontier models through 2026. [secondary]
- **Microsoft 365 Copilot**: now powered by **GPT-5.5 (then GPT-5.6)** as the preferred model per OpenAI (July 2026); Copilot continued as Microsoft's primary AI revenue surface. [vendor-reported via press] http://indianexpress.com/article/technology/artificial-intelligence/openai-gpt-5-6-chatgpt-work-ai-agents-productivity-10779911/

### 6.3 The OpenAI restructuring (27 April 2026)
- Microsoft–OpenAI deal **restructured 27 April 2026**: exclusivity ended; Microsoft IP license continues **non-exclusively through 2032**; Microsoft's revenue share TO OpenAI eliminated; OpenAI's 20%-to-Microsoft revenue share capped; AGI clause removed; Microsoft retains **~27% equity** ($11.8B of $13B funded as of 31 Mar). [secondary] https://www.bowenaistrategiesgroup.com/blog/microsoft-openai-restructure-april-2026.html
- Microsoft participated (undisclosed amount) in OpenAI's $122B March 2026 round. [secondary]
- **Maia** — Microsoft's in-house AI accelerator program continued (Maia 100 lineage; next-gen in 2026). [secondary]

### 6.4 Partnerships & enterprise
- **DoD classified-network AI deployment (1 May 2026)** — Microsoft among eight companies (see §2.4). [secondary]
- **"A Call for Collective Action on Cyber Defense"** (27 Aug 2026) — Microsoft signatory. [secondary]
- Microsoft remained OpenAI's largest cloud partner while diversifying with MAI models and third-party hosting. [secondary]

---

### 6.5 MAI model family — detailed (Build 2026, June 2)

- **Launch:** June 2, 2026, at Microsoft Build 2026 (June 2–3, Fort Mason, San Francisco), unveiled by Microsoft AI CEO Mustafa Suleyman. Framed as the first output of Microsoft AI's in-house "AI Superintelligence Team," trained from scratch on clean/commercially licensed data with no third-party distillation [official][secondary] https://venturebeat.com/technology/microsoft-ai-chief-says-company-was-set-free-from-openai-to-pursue-superintelligence.
- **MAI-Thinking-1** (flagship reasoning model): 35 billion active parameters, 256K token context, low token cost. Microsoft reported, via independent rater Surge AI, blind-test preference over Claude Sonnet 4.6; matches Claude Opus 4.6 on SWE-Bench Pro; **97.0% on AIME 2025 and 94.5% on AIME 2026** [vendor-reported][secondary]. Opened in private preview on Azure Foundry (now "Microsoft Foundry") on June 2, 2026 [official]. All benchmark/preference claims are vendor-reported; independent reproduction not found.
- **MAI-Code-1-Flash**: 5-billion-parameter inference-efficient agentic coding model; 51% on SWE-Bench Pro (Microsoft-reported). Integrated into GitHub Copilot and VS Code at launch [official][secondary].
- **MAI-Image-2.5** (+ MAI-Image-2.5-Flash): text-to-image and image editing; ranked #3 on the Arena AI text-to-image leaderboard and #2 on image-to-image at launch (Microsoft-reported), claimed to surpass "Nano Banana Pro" [vendor-reported][secondary]. Live in PowerPoint; rolling out on OneDrive; on Foundry with claimed market-leading quality-per-dollar [official].
- **MAI-Transcribe-1.5**: speech transcription, claimed state-of-the-art accuracy across 43 languages, 5× faster than competing models (Microsoft-reported); streaming support "coming soon" at launch [vendor-reported][secondary].
- **MAI-Voice-2** (+ MAI-Voice-2-Flash): multilingual speech synthesis for 15+ languages, voice adaptation from a short sample, built-in safeguards against misuse. MAI-Voice-2-Flash was "coming soon" at launch (June 2) [official].
- **OpenRouter availability — verified:** Microsoft's official Build recap states MAI models are also distributed via **Fireworks AI, Baseten, and OpenRouter** (in addition to Foundry). Developers can tune weights themselves through third-party platforms for the first time [official].
- A secondary report (Aug 10, 2026) claims MAI-Voice-2-Flash delivers up to **89% GPU-cost reduction** vs. the OpenAI model it replaced in Dynamics 365 Contact Center, and MAI-Image-2.5-Pro up to **84%** vs. OpenAI's GPT-Image-2 for image generation in PowerPoint [secondary][vendor-reported] — claimed, not independently measured.
- **Business model:** models distributed through Microsoft Foundry and Microsoft first-party products; developers can tune weights via OpenRouter, Fireworks, Baseten [official].
- Suleyman's public stance (interviews ~June 2026): Microsoft was "set free from our contract with OpenAI about six months ago to formally pursue superintelligence"; goal of state-of-the-art models across text/images/audio by 2027; "Humanist Superintelligence" framing [secondary].

### 6.6 Phi family (Microsoft Research)

- **Phi-4-reasoning-vision-15B** released **March 4–5, 2026**: 15B-parameter open-weight multimodal (vision-language) model. Hybrid reasoning modes ('hybrid', 'think', 'no-think') that choose when to use chain-of-thought; trained efficiently on 240 B200 GPUs over 4 days; 200B tokens. Benchmarks: **75.2 MathVista-MINI, 54.3 MMMU-VAL** (vendor-reported). Available on Hugging Face / Azure AI Foundry [secondary] https://www.neowin.net/news/microsoft-releases-phi-4-15b-an-open-weight-ai-model-that-chooses-when-to-think/.
- **Earlier Phi-4 family (timeline context, pre-window):** Phi-4 (14B, Dec 2024), Phi-4-mini (3.8B, Feb 2025), Phi-4-multimodal-instruct (5.6B, Feb 2025), Phi-4-reasoning / reasoning-plus / mini-reasoning (Apr 2025) — all MIT-licensed. **No Phi-5 announcement found in window; no Phi-5 release in 2026 confirmed.**
- **Phi lead personnel (context):** Sébastien Bubeck, key Phi researcher, left Microsoft for OpenAI (reported Dec 2025) [secondary].
- **Project Polaris (secondary only):** mirrors claim Microsoft announced Project Polaris, an in-house mixture-of-experts coding model to replace GPT-4 Turbo as the default GitHub Copilot model starting August 2026, with a 3-month fallback window; runs on Maia accelerators; companion fine-tuning service "Turing Forge"; claimed to outperform GPT-4 Turbo on HumanEval and MBPP [unverified][secondary]. Not corroborated by an official Microsoft announcement in collected sources.

