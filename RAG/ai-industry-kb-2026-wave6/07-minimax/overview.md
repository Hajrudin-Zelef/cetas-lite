---
id: ai-industry-kb-2026-wave6/07-minimax/overview
title: "§7. MiniMax"
domain: minimax
role: deep-dive
task: actor-profile
actors: ["ByteDance", "China", "EU", "Hugging Face", "MiniMax", "United States"]
dates: ["2025-10", "2025-10-27", "2026-02", "2026-02-12", "2026-03-18", "2026-04", "2026-04-12", "2026-06", "2026-06-01", "2026-07-31", "2026-08-03"]
keywords: ["agent", "agentic", "agents", "apache", "attention", "attribution", "benchmark", "consumer", "copyright", "cost", "decode", "ipo"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3074, 3149]
section: "§7. MiniMax"
sha256: 17af719764620b28882d4ff494cf79073af27626ef8c7e9f8525d40e67edf159
---

# §7. MiniMax

Keywords: MiniMax, M2, M2.5, M2.7, M3, MiniMax M3, Hailuo 3, H3, MiniMax Community License, MiniMax Sparse Attention, MSA, Hailuo app, open weights, agent models

## Summary
- MiniMax is a Chinese AI lab whose 2025–2026 releases form a visible license-tightening sequence: M2 shipped plain MIT in October 2025 [SECONDARY], M2.5 moved to modified-MIT with UI attribution in February 2026 [SECONDARY], M2.7 moved to non-commercial terms in March/April 2026 [SECONDARY], and M3 (June 2026) ships under the custom MiniMax Community License that excludes the US, EU, UK, and South Korea from local deployment [SECONDARY].
- **MiniMax M3** launched **June 1, 2026** [SECONDARY] — 428B total / ~23B active MoE, 1M context, MiniMax Sparse Attention (MSA), positioned as an agentic/coding model; API pricing standard $0.60/$2.40 per million input/output tokens, with a launch promotion at $0.30/$1.20 [SECONDARY].
- **Hailuo 3 (H3)** is the third-generation video-line branding: API + Hailuo app release 2026-07-31, open weights on Hugging Face 2026-08-03 [SECONDARY]; 33B dense single-stream omn-modal transformer [SECONDARY].

## Key dated facts
### MiniMax M2 — the permissive start
- **2025-10-27** — MiniMax M2 released [SECONDARY] (one outlier source said Oct 23; the corpus retains Oct 27).
- 230B total / 10B active MoE, plain MIT license [SECONDARY].
- Competing-on-speed-and-cost launch coverage frames M2 as an efficiency/agents play [SECONDARY].

### M2.5 — modified-MIT with attribution
- **2026-02-12** — M2.5 released: 229.9B / 9.8B active, ~192–196K context [SECONDARY].
- Modified MIT requiring prominent "MiniMax M2.5" UI attribution [SECONDARY].
- A Lightning variant was covered alongside the main M2.5 release [SECONDARY].

### M2.7 — self-evolving agent model, non-commercial
- **2026-03-18** — M2.7 announced; API-only at first [SECONDARY].
- Weights followed around **2026-04-12** [SECONDARY]; 230B / ~10B, 200K context [SECONDARY].
- Non-commercial modified license: commercial use needs prior written authorization [SECONDARY].
- After community pushback, MiniMax walked "open source" back to "open weights" for M2.7 [SECONDARY].
- M2.7 scored 56.22% on SWE-bench Pro and 57.0% on Terminal-Bench 2.1 [SECONDARY; benchmark figures are vendor-reported, not standardized boards].

### M3 — sparse attention, 1M context, June 2026
- **2026-06-01** — MiniMax M3 released; the brief's March 13 date is corrected to June 1 [SECONDARY].
- 428B total / ~23B active MoE [SECONDARY]; **1M-token context** [SECONDARY].
- MiniMax Sparse Attention (MSA): vendor materials claim ~9× prefill / ~15× decode at 1M tokens [VENDOR]; the R&D-lead diagram that underlies the claim gave 9.7×/15.6× [VENDOR]; a 109B-testbed technical paper reported 14.2×/7.6× [VENDOR]. All three figures are vendor numbers from different measurements — do not merge them.
- API pricing: standard **$0.60/$2.40** per million input/output tokens; launch promotion **$0.30/$1.20** [SECONDARY].
- One Part-4 line mislabeled M3 as Apache 2.0; the resolved fact is the custom **MiniMax Community License** [SECONDARY].
- MSA vendor-figure audit: the 9.7×/15.6× pair came from an R&D-lead diagram and was rounded to 9×/15× in the shipped launch card; the 14.2×/7.6× pair came from a technical paper measured on a 109B testbed — three measurements, three vendor contexts, never merge them into one claim [VENDOR].
- Pricing audit from the master contradictions log: report both M3 figures with labels — standard $0.60/$2.40, launch-week promotion $0.30/$1.20 [SECONDARY].
- The M2.7 "open source" claim was contradicted by all available evidence (license text non-commercial) before MiniMax retreated to "open weights" wording [SECONDARY].

### M2.5 Lightning and launch coverage
- A **Lightning** variant of M2.5 was covered alongside the main release, framed around speed [SECONDARY].
- Launch coverage framed the M2 line as competing on speed and cost — an efficiency/agent-model positioning that carried through to M3's sparse-attention story [SECONDARY].

### Hailuo 3 (H3) — video-line third generation
- **2026-07-31** — Hailuo 3 API + Hailuo app release [SECONDARY].
- **2026-08-03** — H3 open weights on Hugging Face [SECONDARY].
- 33B dense single-stream omn-modal transformer [SECONDARY]; H3 is the branding of the third Hailuo generation [SECONDARY].

### The Community License — geo-exclusion
- The **MiniMax Community License excludes local deployment in the US, EU, UK, and South Korea** [SECONDARY].
- Cross-check: Llama's Community License restricts *who* (EU multimodal only); MiniMax's restricts *where*, far more aggressively — the two "community licenses" are materially different instruments, neither OSI-approved [SECONDARY].


### New verified facts — expansion (continued — MiniMax company-profile notes and a founding-date conflict)

- IPO-week coverage describes MiniMax as **"established in 2022"**, targeting **"the consumer market, particularly outside China, with its generative AI tools for speech, music and video, as well as text"** [SECONDARY] (techxplore.com).
- This **conflicts with the vibe-investing community guide's "founded in Shanghai in late 2021"** — record both; the one-year gap likely reflects incorporation vs operating history, but neither source disambiguates [SECONDARY contradiction].
- The H3 story's source trail includes **South China Morning Post** ("MiniMax challenges ByteDance with low price, open weights for new H3 model"), **Tech Times** (weights-exclusion and copyright-suit coverage), **Crypto Briefing** (lawsuit coverage), **Segmind** (release/API pricing explainer), and the **Artificial Analysis Video Generation Leaderboard** — corroborating the launch-day, lawsuit, and benchmark facts from independent outlets [SECONDARY] (justbeingresourceful.com).


### New verified facts — expansion (continued — MiniMax H3 deep-dive (launch, licensing catch, lawsuit, pricing))

### Launch and capability (justbeingresourceful.com)
- **Announced July 31, 2026** (teased the day before under **#MiniMaxH3**); described as one omni-modal model that **jointly understands text, images, video, and audio as a single context**, then generates video with **native stereo sound at up to 2K resolution and 15 seconds** in length [SECONDARY].
- Launch demo: reference the **camera movement from a video clip**, have the **character in a photo sing**, **match the vocals to an audio file** — H3 took all three inputs and produced one coherent video; no separate motion-reference tool, lip-sync pass, or manual audio mixing [SECONDARY].
- Usage guidance: **write the relationship, not just the scene** — "use the camera movement from this clip, apply it to this character, match this voice" works better than a generic scene description, because relational instruction is what H3 was trained to parse [SECONDARY].
- Commercial positioning: advertising, branding, e-commerce, product design, UI/UX mockups, gaming assets, film opening titles; noted strengths: **instruction following, accurate on-screen text and brand rendering, video-to-video motion transfer** [SECONDARY].
- This resolves the §7 audio-reference question: **a reference audio file is a first-class input alongside image and video** — the "audio-reference-alone" rule from earlier sources is superseded by the launch documentation [SECONDARY].
- Access: **H3 replaced Hailuo 2.3 as the default model in the Hailuo AI consumer app (hailuoai.video)**; API via MiniMax's platform plus third-party hosts **Segmind and fal** [SECONDARY].

### The licensing catch (H3 Community License)
- Weights published to Hugging Face **August 3, 2026** — what shipped is **H3-Base, a 33.1B-parameter model generating natively at 768p** [SECONDARY].
- The license defines an **"Applicable Territory" for local deployment that specifically excludes the United States, the European Union, the United Kingdom, and South Korea** — users there are not licensed to run, modify, or deploy outputs from locally hosted H3 weights [SECONDARY].
- **Ryan Lee, MiniMax's Head of Developer Relations, confirmed publicly** that the US exclusion ties directly to the company's active copyright litigation [SECONDARY].
- Second stated reason: video models sit in a messier regulatory spot (EU AI Act, evolving UK/South Korean rules, unsettled US landscape); once weights are public the company cannot enforce safeguards downstream, so license restriction became the alternative to delaying release [SECONDARY].
- **The 2K output in every promotional clip comes from H3-Regenerate-2K — a separate module that was NOT open-sourced and remains API-only**; even a fully licensed self-hosted setup must call MiniMax's servers for the marketed resolution [SECONDARY].
- Free for organizations under **$20M annual revenue** (with **"MiniMax H3" in the product interface**), subject to individual licensing for excluded territories; **bans using H3 outputs to train or improve a competing model** — a restriction that applies everywhere [SECONDARY].

