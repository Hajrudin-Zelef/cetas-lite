---
id: labs-hyperscalers-2026/00-labs-hyperscalers/9-mistral-ai
title: "§9 — MISTRAL AI"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: actor-profile
actors: ["AWS", "Alibaba", "Anthropic", "EU", "Microsoft", "Mistral", "OpenAI"]
dates: ["2026-06"]
keywords: ["mistral", "agent", "agentic", "agents", "apache", "benchmarks", "claude", "compute", "consumer", "foundry", "funding", "gpus"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1250, 1298]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 8ae4b6011204acf72000e39851bb886ceaa784daa081475961b1fb214de6a205
---

# §9 — MISTRAL AI

## §9 — MISTRAL AI

> **2026 at a glance — Mistral:** monthly model drops (Voxtral T2, Small 4, Medium 3.5, OCR 4) mostly Apache 2.0; Le Chat → Vibe rebrand (May 28); €3B Series D @ €21B+ (Sep 8, Europe's biggest private tech round); Microsoft multibillion EU compute deal (Jul 21); HUMAIN/Saudi pact (Aug 24); Airbus/BMW/EDF 5-year deals; 1 GW-by-2030 sovereign compute plan.

### 9.1 Model releases (Feb → Sep 2026)
- **Mistral Small 4** — 2026 release of the efficient small-model line. [secondary: Track C]
- **Mistral Medium 3.5** — mid-tier release. [secondary]
- **Voxtral** — Mistral's voice/audio model line (speech-to-text / audio intelligence). [secondary]
- **Mistral OCR 4** — document-understanding / OCR model, 4th generation. [secondary]
- **Le Chat** — continued as Mistral's consumer assistant surface; enterprise "Le Chat Enterprise" push. [secondary]

### 9.2 "Vibe" rebrand
- Mistral underwent a **"Vibe" rebrand** during the window (Track C) — new brand identity for the Le Chat / consumer surface. [secondary]

### 9.3 Funding — €3B Series D
- **€3 billion Series D** — Mistral's landmark 2026 raise, cementing its position as Europe's AI champion. [secondary: Track C]
- Valuation and lead investors per Track C sources (details in §16 funding scoreboard).

### 9.4 Sovereign infrastructure & partnerships
- **Sovereign AI infrastructure**: Mistral positioned as the EU's sovereign-AI champion; partnerships with European governments and enterprises. [secondary]
- **Partnerships**: continued alliances with Microsoft (Azure), and European industrial partners (Track C).
- Data-center / compute partnerships in Europe (Track C).

### 9.5 Strategy notes
- Open-weights strategy continued alongside proprietary frontier offerings. [secondary]
- EU AI Act: Mistral as both GPAI provider and the EU's political showcase for "European AI." [secondary]

---

### 9.6 Mistral model releases — detailed (from Track C)

**Voxtral Transcribe 2 — Feb 4, 2026:** Two speech-to-text models: **Voxtral Mini Transcribe V2** (batch) and **Voxtral Realtime** (live, streaming architecture). API IDs: `voxtral-mini-2602`, `voxtral-mini-transcribe-realtime-2602`. 13 languages, speaker diarization, word-level timestamps, context biasing; Realtime latency configurable down to sub-200ms; 4B params (Realtime). Open weights under Apache 2.0 [independent: VentureBeat, TechCrunch]. Transcription API pricing reported at ~$0.003/minute [secondary].

**Mistral Moderation 2603 — Mar 12, 2026:** New moderation model; custom guardrails added for Agents and Conversations API calls [secondary].

**Mistral Small 4 (`mistral-small-2603`) — Mar 16, 2026:** 119B total-parameter **sparse MoE** (128 experts, 4 active per token; ~6–6.5B active/token), 256k context, native image input / text output, per-request `reasoning_effort` (none → high), 40% lower latency and 3× throughput vs Small 3. Consolidates Magistral (reasoning), Pixtral (vision), Devstral (coding) capabilities in one checkpoint [vendor-reported]. **Open weights, Apache 2.0** (HF: `mistralai/Mistral-Small-4-119B-2603`, plus NVFP4 and eagle-head speculative-decoding variants). API price: **$0.15 / $0.60 per 1M input/output tokens** [secondary]. Reported benchmarks (vendor): GPQA Diamond 71.2%, MMLU-Pro 78.0% (vs GPT-4o-mini 40.2%/64.8%) [vendor-reported]. SiliconANGLE covered the release Mar 17, 2026 [independent].

**Leanstral (`labs-leanstral-2603`) — Mar 16, 2026:** First open-source AI agent for **Lean 4 formal proof engineering**; same 119B-class architecture as Small 4; claimed FLTEval pass@16 of 31.9 vs Claude Sonnet 4.6's 23.9 [secondary; vendor benchmarks]. License conflict: one roundup says it is the only major 2026 Mistral release *not* Apache 2.0; other sources say Apache 2.0. Verify against official blog [unverified].

**Voxtral TTS (`voxtral-tts-2603`) — ~Mar 23–26, 2026:** 4B-parameter streaming text-to-speech, open weights; 9 languages (EN, FR, DE, ES, NL, PT, IT, HI, AR); zero-shot voice cloning from <5s sample; time-to-first-audio 90ms, 6× real-time factor [independent: TechCrunch, AI Business] https://techcrunch.com/2026/03/26/mistral-releases-a-new-open-source-model-for-speech-generation/. Based on Ministral 3B backbone [independent].

**Mistral Medium 3.5 (`mistral-medium-3-5`) — Apr 28–29, 2026:** 128B **dense** model, 256k context, unifying instruction-following, reasoning and coding; configurable `reasoning_effort` [vendor-reported]. Benchmarks (vendor): **77.6% SWE-Bench Verified**, ahead of Devstral 2 and Qwen3.5 397B A17B on coding/agentic evals per vendor charts [vendor-reported; outside criticism noted — ML professor Pedro Domingos publicly questioned readiness — [secondary]]. **Open weights under a modified MIT license** (not Apache 2.0); deployable on as few as 4 GPUs [vendor-reported]. Launched alongside remote coding agents in Mistral Vibe CLI and Work Mode in Le Chat [independent: TestingCatalog, TechGlimmer, ODSC] https://www.testingcatalog.com/mistral-ai-unveils-medium-3-5-model-and-work-mode-for-le-chat/.

**OCR 4 — ~Jun 2026:** Document-understanding model: structured extraction with bounding boxes, block classification, inline confidence scores, **170 languages**; self-host option; available via Mistral API, Document AI (no-code in Mistral Studio), Amazon SageMaker, Microsoft Foundry; Snowflake Parse Document support planned [secondary] https://www.fonearena.com/blog/485781/mistral-ocr-4-features.html. Pricing: **$4 / 1,000 pages** API; **$2** batch API (50% off); Document AI $5 / 1,000 pages [secondary].

**Robostral Navigate — Jul 2026 [unverified]:** Described as Mistral's first robotics model, formalizing the "physical AI" push. Sourced only from an internal research doc mirrored on GitHub (marked "not for public repo") [unverified] — treat cautiously.

**What did NOT ship in-window:** No new Mixtral (8x22B/8x7B successors) announced Feb–Sep 2026 [unverified negative]. **Mistral Large 3** (675B MoE, 41B active, Apache 2.0) was released **Dec 2, 2025** — before the window; remains the flagship open-weight model [secondary]. **Devstral 2** released Dec 2025; not a 2026 release [independent]. A viral June 2026 hoax ("Le Chaton Fat" meme) was **fake** — no such model exists [independent: ExplainX] https://explainx.ai/blog/le-chaton-fat-mistral-ai-viral-hoax-meme-2026.

