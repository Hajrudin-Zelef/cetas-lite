---
id: labs-hyperscalers-2026/00-labs-hyperscalers/5-apple
title: "§5 — APPLE"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["Anthropic", "Apple", "Broadcom", "China", "EU", "Google", "Nvidia", "OpenAI", "TSMC", "United States"]
dates: ["2025-10", "2026-01-12", "2026-06", "2026-06-08", "2026-07-06", "2026-08", "2026-09-14"]
keywords: ["3nm", "capex", "chatgpt", "compute", "cost", "distillation", "distribution", "dram", "gemini", "gpu", "gpus", "inference"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [778, 867]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 87c8c0e0a3d9ea2ff823b7074666d3d97d5e4215000160624138b9146ee2306a
---

# §5 — APPLE

## §5 — APPLE

> **2026 at a glance — Apple:** no frontier model of its own; the Siri AI rebuild shipped with iOS 27 (Sep 14, 12 GB RAM gating) on AFM 3, with Google Gemini handling complex server-side queries — a landmark partnership; most cautious frontier-adjacent player.

### 5.1 Apple Intelligence & Siri (2026)
- **iOS 27** (announced WWDC June 2026, shipping Sept 2026) — **Siri AI** rebuild: the long-delayed personalized Siri, rebuilt on new Apple foundation models with deeper app-intent integration. [secondary: Track B]
- **Apple Foundation Models v3 (AFM 3)** — third-generation on-device/server model family powering Apple Intelligence features in 2026. [secondary]
- **12 GB RAM gating**: Apple Intelligence features gated to devices with **≥12 GB RAM** — effectively iPhone 17 Pro and newer / M-series Macs & iPads; older devices excluded, a notable segmentation decision. [secondary]

### 5.2 The Google Gemini partnership
- Apple confirmed a **partnership with Google to power parts of Siri / Apple Intelligence with Gemini models** — a landmark admission that Apple's in-house models lagged frontier competitors. [secondary: Track B]
- Structure reported: Gemini handles complex queries server-side; Apple retains on-device AFM for privacy-sensitive tasks. Commercial terms not disclosed. [secondary]

### 5.3 Hardware & silicon
- **"Baltra"** — Apple AI server chip / Private Cloud Compute silicon program referenced in Track B roadmaps ([secondary/unverified naming]).
- **M-series / A-series**: continued on-device NPU scaling for Apple Intelligence (Track B).
- iPhone 17 lineup (Sept 2026) shipped with Apple Intelligence as a headline feature. [secondary]

### 5.4 Strategy notes
- Apple remained the **most cautious frontier-adjacent player**: no frontier-scale foundation model release of its own in Feb–Sep 2026; strategy = on-device + private cloud + Gemini partnership. [secondary]
- Regulatory: EU DMA interoperability obligations continued to shape Apple Intelligence rollout in Europe (Track B).

---

### 5.5 Apple Intelligence 2026 — detailed (from Track B)

- **WWDC 2026 (June 8, 2026):** Apple unveiled the next generation of Apple Intelligence and the rebuilt **Siri AI** for iOS 27, iPadOS 27, macOS 27 ("Golden Gate" per press), watchOS 27, visionOS 27 [secondary] https://www.eweek.com/news/apple-ai-features-2026/.
- **Shipping:** iOS 27 released **September 14, 2026**; Apple support page "How to get the next generation of Apple Intelligence" (published Sep 14) says features are automatically available on eligible devices with the 27.0 releases [secondary].
- **Feature set (press summaries of Apple's own materials):** Siri AI full rebuild (natural conversation, personal-context understanding, onscreen awareness, system-wide app actions, 20-turn conversations, Camera Siri mode, message/email lookup, drafting help, photo actions, web answers); app actions across Messages/Music/Reminders/Calendar/Mail/Photos; on-screen Visual Intelligence; Photos (Spatial Reframing, Extend, upgraded Clean Up); photorealistic Image Playground (planned SynthID support); Safari smart tab topics + "Notify Me" website alerts; Call Context (codes/reservations); Health Insights tab; AI-generated subtitles for any video; accessibility upgrades (Smarter VoiceOver, natural-language Magnifier, Voice Control, Accessibility Reader); system-wide AI for Shortcuts and Home; AI-assisted parental controls; Wallet bill-splitting; self-healing Passwords [secondary].
- **Dedicated Siri app:** lets users revisit conversations and continue tasks across devices, synced via iCloud [secondary]. **ChatGPT integration continues** (OpenAI chatbot usable via Writing Tools and Siri) [secondary].

### 5.6 Device gating — 12 GB unified-memory requirement (detailed)

Apple confirmed its most advanced on-device model requires **≥12 GB unified memory**, gating flagship features to [secondary] https://www.gsmarena.com/apples_most_advanced_ondevice_ai_features_will_only_work_on_select_devices-news-73187.php:
- iPhone: iPhone 17 Pro, iPhone 17 Pro Max, iPhone Air (all 12 GB); standard iPhone 17 (8 GB) excluded from the advanced tier. Baseline Apple Intelligence: iPhone 15 Pro/Pro Max, all iPhone 16 models, all iPhone 17 models.
- iPad: M4 or newer with ≥12 GB unified memory. Mac: M3 or later with ≥12 GB RAM. Apple Vision Pro with M5.
- Storage footprint: Apple Intelligence can take **up to 14 GB on iPhone 17 Pro/Pro Max/Air**, up to 8 GB on other supported iPhones (per Apple's Sep 14, 2026 support figures) [secondary].
- "iPhone Air" naming/details come from press (Apple's Sep 2026 event coverage) — treat exact configurations as [secondary].

### 5.7 Siri AI rollout staging

- Launches **English-first as a waitlisted beta** (join via Settings → Apple Intelligence); more languages to follow; **32 locales** supported over time [secondary].
- **Blocked at launch:** iPhone and iPad in the **EU** (DMA regulatory compliance pending); available in EU on Mac, Vision Pro, and Watch (watchOS 27 requires paired iPhone, so Watch EU availability is effectively constrained). **Mainland China:** unavailable pending regulatory approval [secondary].
- Privacy routing: simple requests on-device; complex tasks via **Private Cloud Compute** (data not stored/shared per Apple); heaviest requests handled by a Gemini-powered cloud model (see §5.8) [secondary].

### 5.8 Apple Foundation Models, 3rd generation (AFM 3) — detailed

Apple's most substantive 2026 ML research publication, per Apple ML Research blog (technical report behind the Siri AI launch): **five models** [secondary reporting on the official report]:
1. **AFM 3 Core** — next-gen ~3B-parameter dense on-device model (exposed to developers via the Foundation Models framework, Swift API; mixed 2-bit/4-bit palletization, avg 3.7 bits/weight; ANE-native inference ~30 tok/s on iPhone 15 Pro; third-party LoRA adapters supported).
2. **AFM 3 Core Advanced** — Apple's most powerful on-device model; **20B-parameter sparse** model using **Instruction-Following Pruning (IFP)** (a lightweight predictor selects which FFN rows/columns to activate; only 1–4B params active per request; full model in flash, selected experts in DRAM); natively multimodal — powers expressive Siri voices and higher-accuracy dictation (the features gated behind 12 GB RAM).
3. **AFM 3 Cloud** — Private Cloud Compute text model.
4. **ADM 3 Cloud** — Private Cloud Compute image-generation model.
5. **AFM 3 Cloud Pro** — Apple's most capable model; runs on **Nvidia GPUs hosted in Google Cloud** under an extended Private Cloud Compute architecture (Apple's own statement per press).
- Built **with Google**: Apple says models are custom builds for Apple Silicon, trained on proprietary data, and refined using outputs from Gemini frontier models (distillation-style). Apple AI VP Amar Subramanya: *"custom builds for Apple Silicon, trained using proprietary data, and refined using outputs from Gemini frontier models."* Craig Federighi: *"The amount of the Google Assistant we use is none"* — no Google branding in the interface [secondary].
- IFP technique itself published by Apple Research in Jan 2025, shipped in production with AFM 3 Core Advanced [secondary].
- Exact parameter counts (20B) and the five-model taxonomy come from developer press reads of the WWDC-day materials, not the fetched Apple ML blog page itself [secondary].
- Foundation Models framework for developers: exposes the ~3B on-device LLM (iOS 26+/macOS 26+) with custom LoRA adapter training via Apple's Python toolkit; ANE-native, ~1/10th the power of GPU alternatives [secondary].

### 5.9 Apple–Google Gemini partnership, "Project Campos" — detailed

- **Announced January 12, 2026:** joint Apple–Google statement — multi-year, **non-exclusive** partnership; "Apple Foundation Models will be based on Google's Gemini models and cloud technology." Apple chose Google after "careful evaluation" (also evaluated OpenAI and Anthropic); Google called Gemini "the most capable foundation" for Apple Foundation Models [secondary] https://tech-insider.org/apple-google-gemini-siri-deal-1-billion-2026/.
- **Financial terms:** reported at **~$1B/year** (reportedly up to $5B over the term); neither company officially confirmed the figure [secondary].
- **Technical scope:** custom Gemini model built for Apple, described as **1.2 trillion parameters** (~8× Apple's previous in-house cloud models, reported as ~150B) — optimized for Siri's frequent tasks (summarization, planning). The model weights run **inside Apple's Private Cloud Compute infrastructure**, not on Google's servers; Apple strips identifying information, Google reportedly never sees user identity. One outlet claims the model in shipping iOS 27 Siri is "Gemini 2.5 Pro" while others describe a custom 1.2T-param model — the naming is inconsistent across sources; treat exact model name as [unverified] [secondary/unverified].
- **Timeline:** first Gemini-powered Siri features with **iOS 26.4 (spring 2026)** under internal codename **Project Campos** (on-screen awareness, multi-step cross-app automation, complex reasoning); full rebuilt Siri AI as iOS 27's chatbot-grade assistant (WWDC June 8, 2026; shipping Sep 14, 2026) [secondary].
- Apple states no user data is shared with Google; training used Gemini models plus Google TPUs (lab-side, during development). Partnership framed as a bridge until Apple's own autonomous capability arrives (~2027 per some reporting — [unverified]) [secondary].

### 5.10 AI server chips & Private Cloud Compute expansion

- **Baltra** — Apple's first dedicated AI server chip, **co-developed with Broadcom**, purpose-built for AI inference (not repurposed Mac silicon), reportedly on **TSMC's N3P 3nm-class process**; for Apple's own infrastructure only (not sold externally) [secondary] https://tech-insider.org/apple-ai-servers-m8-ultra-nvidia-nvlink-2026/.
- **Timeline:** mass production targeted **H2 2026** (Ming-Chi Kuo / TechSpot reporting); new Apple-run data centers built around it coming online **2027**; some reports note the schedule slipped from an earlier internal target — Apple has confirmed none of these dates [secondary/unverified].
- **Apple–Broadcom (July 6, 2026):** regulatory filing disclosing the supply agreement now runs **through 2031**, covering custom ASICs plus RF/Wi-Fi/Bluetooth components for multiple product generations; financial terms private; Apple ~20% of Broadcom revenue [secondary].
- **Bridge hardware:** Apple announced the **M5 Ultra chip in late August 2026**; reporting says a beefed-up M5 Ultra-based server platform is bridging until Baltra/M7-Ultra server platforms are ready [secondary].
- **Background:** "ACDC" (Apple Chips in Data Centers) project reported since 2024; M8 Ultra server-class plans also reported [secondary].
- **US manufacturing / data centers:** Apple's $50B US plan (announced Feb 2025) included a **Houston, Texas factory** for AI servers — operating ahead of schedule since **October 2025**, shipping US-manufactured servers for the Apple Intelligence platform. Additional data-center capacity expansion in **North Carolina, Iowa, Oregon, Arizona, Nevada** [secondary].
- No verified new AI-executive hires or named AI-capex figures for Apple in the fetched sources for this window (Apple discloses little; the Google deal ~$1B/yr and $50B US manufacturing plan are the closest verified-scale investments) [research finding].

### 5.7 Apple–Google Gemini partnership — what each side gets (synthesis)

| | Apple gets | Google gets |
|---|---|---|
| Capability | Frontier server-side reasoning for complex Siri queries without building its own frontier model | Distribution to ~2B+ Apple devices; default-assistant footprint |
| Privacy story | On-device AFM 3 + Private Cloud Compute for sensitive tasks; Gemini only for escalated queries | Operates inside Apple's privacy frame (reported structure) |
| Strategic cost | Admission that in-house models lagged; dependency on a direct competitor | Revenue share / commercial terms undisclosed |
| Competitive effect | Siri AI ships on time (iOS 27, Sep 14, 2026) | Gemini becomes the intelligence layer of two mobile ecosystems (Android + iPhone) |

- **ChatGPT integration continues** in parallel (Writing Tools, Siri) — Apple is multi-sourcing foundation models, not single-sourcing [secondary].
- **Regulatory overlay:** EU DMA interoperability obligations shaped the Apple Intelligence rollout in Europe [secondary].
- **Open question:** whether the partnership extends to future AFM generations or remains a stopgap — commercial terms undisclosed [secondary].


