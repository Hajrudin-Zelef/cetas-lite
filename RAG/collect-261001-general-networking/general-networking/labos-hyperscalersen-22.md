---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-22
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Apple", "Broadcom", "EU", "Google", "Meta", "Microsoft", "Nvidia", "OpenAI", "TSMC"]
dates: ["2025-04", "2026-04-08", "2026-06", "2026-06-08", "2026-07-09", "2026-07-28", "2026-09", "2026-09-14", "2026-09-18"]
keywords: ["accelerator", "amd", "blackwell", "capex", "chatgpt", "compute", "cost", "energy", "gemini", "gpu", "gpus", "inference"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [759, 810]
sha256: 6c615f3cc6b8e46046b6ed6515048408b5de1f5c51e03ecda8d173b5dc215b75
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

- **Capex guidance 2026:** initially $115–135B; raised (April) to **$125–145B** including finance-lease payments; Q2 reporting raised the floor again to **$130–145B**. ~70% of capex is now AI infrastructure. Q2 2026 revenue **$60.8B (+28% YoY)**; profit missed expectations; shares fell on margin concerns [secondary].
- **Compute targets (Reuters, July 9, 2026, internal comms):** ~7 GW of compute capacity by end of 2026 (1 GW deployed H1 2026 + 2.5 GW planned H2), doubling to **14 GW in 2027** [secondary].
- **MTIA "Iris":** next MTIA accelerator reportedly begins production **September 2026**; co-developed with **Broadcom**, manufactured by **TSMC**; plan to refresh chips roughly every 6 months through 2027 [secondary — Reuters via DataFLOQ].
- **AMD deal (Feb 2026):** up to **$60B over 5 years** for AMD AI chips, per Reuters reporting (also described as "parallel 6 gigawatts of AMD GPUs") [secondary — primary Reuters URL not directly fetched].
- **Nvidia:** remains Meta's primary GPU supplier; training compute reportedly on 500,000+ Blackwell B200-class GPUs [unverified]; Meta also rents large Google TPU clusters [secondary/unverified].
- **BlackRock JV (July 28, 2026):** 1-GW AI data-center campus in **El Paso, Texas**, ~$14B total cost; BlackRock funds hold 80% equity, Meta 20%; capacity online from 2028. Separate 1-GW facility in **Alberta** announced alongside [secondary].
- **Tulsa:** $1B AI-optimized data center announced (1,000+ construction jobs, ~100 permanent); water-efficient cooling, clean-energy matching [secondary].
- **Corning (Jan 2026):** multi-year **$6B** agreement for fiber-optic cables for data centers [secondary].
- **"Tents":** Meta experimenting with lightweight rapidly-deployable data-center structures — permanent facilities take 12–24 months vs. weeks for tents; signal is speed-to-capacity [secondary].
- Power strategy: direct gas-turbine deals, long-term nuclear PPAs, on-site battery storage; behind-the-meter generation toward GW-scale campuses [secondary].
- IDC (via secondary): global AI infrastructure spend projected **$497B in 2026 (+56% YoY)** [secondary].

### 4.12 Llama status — detailed (from Track B)

- **Last verified Llama-family release:** Llama 4 Scout / Maverick, April 2025. No Llama 4.x point release verified in Feb–Sep 2026.
- **Llama 5: UNVERIFIED — conflicting claims. Do not treat as a shipped product fact.** Several low-quality aggregator pages claim Llama 5 launched **April 8, 2026** at a "Meta AI Connect summit" with "600B+ parameters" and "5M-token context," citing a CNBC article whose actual title references the Muse Spark line, not a Llama launch [unverified]. A separate aggregator narrative ("Llama 5 Avocado," Aug–Sep 2026) says Meta's next open model project finished basic training in early 2026 but launch was pushed back [unverified]. A structured Meta-Llama tracking page (verification stamp 2026-09-18) states explicitly: *"No Llama 4.1, 4.5 or 5 exists on any first-party source"* [secondary].
- **AI pricing tracker (Aug 10, 2026 checkpoint):** Meta announced it will *resume* open-source model releases "soon" but published no model name, checkpoint, license, hardware target, or endpoint; hosted Llama API pricing remains third-party inference providers only — Meta publishes weights, not a first-party token price [secondary] https://www.aipricing.guru/meta-llama-pricing/.
- **Meta AI app / Reality Labs:** Meta AI app and meta.ai added "Thinking" mode powered by Muse Spark 1.1 (July 9, 2026) [vendor-reported/secondary]. No verified 2026 Reality Labs AI-specific launches found in fetched sources. **Meta Connect 2026 scheduled Sept 23–24, 2026** — after this report's cutoff; re-verify post-Connect [secondary].

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

