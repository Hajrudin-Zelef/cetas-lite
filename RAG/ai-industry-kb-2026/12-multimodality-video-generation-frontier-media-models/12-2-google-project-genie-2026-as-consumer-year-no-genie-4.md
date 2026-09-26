---
id: ai-industry-kb-2026/12-multimodality-video-generation-frontier-media-models/12-2-google-project-genie-2026-as-consumer-year-no-genie-4
title: "12.2 Google Project Genie: 2026 as consumer year — no Genie 4"
domain: multimodality-video-generation-frontier-media-models
role: deep-dive
task: multimodal
actors: ["Apple", "ByteDance", "Google", "OpenAI", "United States"]
dates: ["2025-08-05", "2025-09-30", "2026-01-29", "2026-02-12", "2026-03-24", "2026-04-26", "2026-05", "2026-05-21", "2026-06-23", "2026-07-31", "2026-08-31", "2026-09-22", "2026-09-24", "2026-09-30"]
keywords: ["consumer", "acquisition", "chatgpt", "compute", "cost", "distribution", "ipo", "omni", "research", "revenue", "robotics", "text-to-video"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6174, 6204]
section: "12. Multimodality — Video Generation & Frontier Media Models"
sha256: 49ebae588c139066c9e04b1b34a07f16dfc81a72d9015cb0ea5de8d840e0f4e3
---

# 12.2 Google Project Genie: 2026 as consumer year — no Genie 4

- **Launch (2025-09-30 — explicitly NOT 2026).** VentureBeat's launch report confirms Sora 2 debuted as a standalone iOS app on 2025-09-30; one secondary page dated the launch "2026-09-30," a one-year-off error that must not propagate.
- Headline features at launch: text-to-video with natively synchronized audio; a **"Cameos"** feature inserting the user's likeness; initial US/Canada availability; free limited access; **Sora 2 Pro** for ChatGPT Pro subscribers.
- API snapshots `sora-2` / `sora-2-pro`: new generations up to 20 seconds, up to six extensions, 120 seconds maximum total length; 720p for Sora 2, up to 1080p portrait/landscape for 2 Pro.
- Strict content policies: only under-18-suitable content; no real people or public figures; human-likeness character uploads blocked by default; no copyrighted characters or music.
- The free limited access tier and the Cameos feature were the consumer-acquisition mechanics; Sora 2 Pro (ChatGPT Pro subscribers) was the monetization attempt — the reported unit economics (~$5.4B compute vs ~$2.1M lifetime revenue) show it failed at both ends, which is why the shutdown is attributed to cost structure rather than to demand.
- **Shutdown announced 2026-03-24** via the official Sora account; API developers were notified of the Sora 2 deprecation the same day.
- **The Sora web and app experiences were discontinued on 2026-04-26** [VERIFIED — OpenAI Help Center, "What to know about the Sora discontinuation," updated ~2026-08-31].
- **The Videos API was scheduled for permanent shutdown on 2026-09-24** — planned, not retrospective, relative to the 2026-09-22 writing date; maintenance-only until then. **This single date post-dates the knowledge cutoff and must be reverified.**
- Users could export content at `sora.chatgpt.com/sunset`; after discontinuation and any final export window, all data associated with Sora use is permanently deleted.
- **Why it happened (secondary, corroborated across outlets — CIOL citing the official line):** applications CEO Fidji Simo told employees the company "could no longer afford what she called 'side quests'"; Sam Altman made the shutdown call and redirected compute to higher-priority products.
- A reported **$1B partnership with Disney** (franchise content integration) was affected; per the Wall Street Journal, Disney was informed shortly before the public announcement and **no money had changed hands**.
- Likeness-moderation disputes ran through Sora's life: the estate of Martin Luther King Jr. pushed OpenAI to restrict videos using his likeness (Washington Post); unauthorized Robin Williams videos drew estate complaints.
- **Unit economics (secondary, UNVERIFIED methodology):** CIOL reports the product was "spending about $5.4 billion annually in compute while generating roughly $2.1 million in lifetime revenue." No outlet disputes the direction; the numbers are a reported claim, not an established fact. 1M+ downloads in the first five days (ngram) did not change the equation.
- **Precision on "Sora is dead":** the Sora 2 *model* remains available inside ChatGPT's paid tiers; only the standalone product, dedicated infrastructure, and public API end. The Sora team continues as a **research unit focused on world simulation**, with applications in robotics and physical-environment modelling — the world-model ambition survives as research, not as a product. No successor has been named; ngram's analysis notes the official deprecations page's **"replacement" column is empty**.
- **IPO context (analyst framing, not OpenAI's stated reason):** CIOL reports OpenAI was preparing an IPO at a ~$300B reported valuation and frames the shutdown as removing a loss-making product ahead of it; the official reason given was compute reallocation.

### 12.2 Google Project Genie: 2026 as consumer year — no Genie 4

- **Baseline (2025-08-05):** DeepMind announced Genie 3 — an interactive world model generating 720p/24fps explorable environments lasting several minutes, with "promptable world events" letting the user describe scenarios and alter the world in real time.
- **2026-01-29: consumer beta via Google Labs** [secondary, consistent across sources]. Project Genie opened to consumers as a Google Labs experiment, gated behind the **Google AI Ultra plan ($249.99/mo)**.
- Reported properties: real-time interactive worlds from text prompts; a **"remix"** function that turns one world into another; worlds shareable via public link.
- The remix function is the Genie-native editing primitive: world-to-world transformation in place, rather than video's frame-level or clip-level editing. Combined with promptable world events (altering the world in real time by describing scenarios), it makes Genie a runtime-editing surface — the interactive counterpart to Omni Flash's conversational video editing, in a different product category.
- Google's stated motivation for the May 2026 grounding expansion (planning and training, not just play) reframes Genie from entertainment experiment to simulation infrastructure — the category where the [UNVERIFIED] Waymo claim would sit if confirmed.
- Key distinction: Genie produces **playable runtime worlds, not exportable 3D assets or video files** — a simulation surface closer to a game engine than to a video generator.
- **2026-05-21: Street View/Maps Imagery Grounding expansion** [secondary — Pulse2]. Project Genie gained grounding on real-world imagery (Street View/Maps), letting users generate simulations rooted in actual locations; location grounding initially covered **US locations**; Google AI Ultra access expanded to more countries globally.
- Google's stated motivation: simulations anchored to the real world make Genie useful for planning and training, not just play.
- **Waymo connection [UNVERIFIED]:** secondary blogs claim Waymo uses a fine-tuned Genie variant for autonomous-driving simulation; no primary DeepMind/Waymo confirmation was found — plausible given the world-model framing, but treat as rumor until a first-party source appears.
- **No Genie 4.** As of 2026-09-22, no new numbered Genie model was found. Contrast with the video-generator race, where model versions kept climbing: Genie's 2026 story is distribution + grounding, not model iteration.

### 12.3 ByteDance Seedance: 2.0 (2026-02-12) and 2.5 (announced 2026-06-23; Dreamina 2026-07-31)

