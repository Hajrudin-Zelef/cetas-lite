---
id: labs-hyperscalers-2026/00-labs-hyperscalers/5-10-ai-server-chips-private-cloud-compute-expansion
title: "5.10 AI server chips & Private Cloud Compute expansion"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: funding-deals
actors: ["Anthropic", "Apple", "Broadcom", "EU", "Google", "OpenAI", "TSMC", "United States"]
dates: ["2025-10", "2026-01-12", "2026-06-08", "2026-07-06", "2026-08"]
keywords: ["compute", "3nm", "capex", "chatgpt", "cost", "distribution", "gemini", "inference", "nvidia", "nvlink", "parameters", "reasoning"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [838, 867]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 99e77777c8c770746d1e03a4f8728c970b2a8e9e532c20ce5d80b61f853ad284
---

# 5.10 AI server chips & Private Cloud Compute expansion

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


