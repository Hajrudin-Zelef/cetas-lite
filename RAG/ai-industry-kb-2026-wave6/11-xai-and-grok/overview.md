---
id: ai-industry-kb-2026-wave6/11-xai-and-grok/overview
title: "§11. xAI and Grok"
domain: xai-and-grok
role: deep-dive
task: actor-profile
actors: ["Google", "OpenAI", "Oracle", "xAI"]
dates: ["2025-07-09", "2025-08", "2025-08-10", "2025-08-28", "2025-09", "2026-03-16", "2026-04-15", "2026-04-30", "2026-05-15", "2026-07-08", "2026-07-15", "2026-07-29", "2026-08-12", "2026-09-17", "2026-09-21", "2026-09-22"]
keywords: ["grok", "agent", "agents", "agi", "apache", "astra", "benchmark", "compute", "cost", "gemini", "gpt-6", "gpus"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5149, 5222]
section: "§11. xAI and Grok"
sha256: de3ce907a657ecb4e1a22a3ed105802f6c98f96597c5a99aad6675c26159511f
---

# §11. xAI and Grok

Keywords: xAI, Grok, Grok 4.3, Grok 4.5, Grok 4.6, Grok 4.7, Grok 5, Grok Code Fast 1, Grok Build, V9 base, Voice Think Fast, Voice Transcribe, context-window pricing, Sonic, 200K cliff

## Summary
- xAI's 2026 cadence ran **Grok 4.3 (2026-04-30) → Grok 4.5 (2026-07-08/09) → Grok 4.6 (2026-08-12) → Grok 4.7 (2026-09-21)** — **there is no Grok 4.4**; xAI skipped it, and no release was found between 4.3 and 4.5 [SECONDARY].
- **Grok 4.3**: 1M context, $1.25/$2.50 per million input/output, 84% cache-hit discount [SECONDARY].
- **Grok 4.5**: public API **2026-07-08/09** (task resolution: July 8); new V9 base [SECONDARY].
- **Grok 4.6**: 2026-08-12, post-training update on the V9 base, 500K context, $2/$6 [SECONDARY].
- **Grok 4.7**: 2026-09-21, 500K context, **$2/$6 at ≤200K tokens, $4/$12 above 200K** — the 2026 context-window pricing cliff [SECONDARY]; DeepSWE v1.1 71.0% [VENDOR]; Terminal-Bench 4.0 26% [SECONDARY, alextech].
- **Grok 5 had not shipped as of 2026-09-22**; missed targets are vendor intent/secondary reporting, not a release [UNVERIFIED].
- **Grok Code Fast 1** is August 2025 (announced Aug 28, codename Sonic), not 2026; no 2026 successor was found [SECONDARY]; its benchmark claims are [VENDOR].
- **Grok Build CLI** went to beta **2026-05-15** and was open-sourced **2026-07-15** as `xai-org/grok-build` under Apache 2.0 (Rust) [SECONDARY].
- The 2026 voice timeline: Text-to-Speech GA 2026-03-16; Speech-to-Text GA 2026-04-15; Voice Think Fast 2.0 2026-07-29; Voice Transcribe 2.0 2026-09-17 [SECONDARY]; one source gives Apr 17–19 for unversioned STT/TTS — treat as rollout-date nuance, not a contradiction.

## Key dated facts
### The 4.x cadence — and the missing 4.4
- **2026-04-30** — Grok 4.3 released: 1M context, $1.25/$2.50 per million input/output, 84% cache-hit discount [SECONDARY].
- **No Grok 4.4** — xAI skipped the number; no release exists between 4.3 and 4.5 [SECONDARY].
- **2026-07-08/09** — Grok 4.5 public API release (task resolution: July 8); new **V9 base** [SECONDARY].
- **2026-08-12** — Grok 4.6: post-training update on the V9 base, 500K context, $2/$6 pricing [SECONDARY].
- **2026-09-21** — Grok 4.7: 500K context; **$2/$6 at ≤200K tokens, $4/$12 above 200K** [SECONDARY].
- Grok 4.7 DeepSWE v1.1 71.0% [VENDOR]; Terminal-Bench 4.0 26% [SECONDARY].
- xAI's context-window cliffs (4.6/4.7 long-context pricing from 200K) are standard 2026 practice — cf. Gemini 3.1 Pro's 200K surcharge and GPT-6 Astra's 272K cliff [SECONDARY].

### Grok Code Fast 1 — 2025, not 2026
- **2025-08-28** — Grok Code Fast 1 announced (codename **Sonic**) — the brief's 2026 date is corrected to 2025 [SECONDARY].
- No 2026 successor was found in the sources [SECONDARY].
- Benchmark claims for Grok Code Fast 1 are [VENDOR].

### Grok Build — open-sourced coding agent
- **2026-05-15** — Grok Build CLI beta [SECONDARY].
- **2026-07-15** — open-sourced as **xai-org/grok-build**, Apache 2.0, written in Rust [SECONDARY].

### Voice timeline 2026
- **2026-03-16** — Text-to-Speech GA [SECONDARY].
- **2026-04-15** — Speech-to-Text GA [SECONDARY].
- **2026-07-29** — Voice Think Fast 2.0 [SECONDARY].
- **2026-09-17** — Voice Transcribe 2.0 [SECONDARY].
- One source gives **Apr 17–19** for unversioned STT/TTS GA — treat as rollout-date nuance between announcement and general availability, not a hard contradiction [SECONDARY].
- Voice is the xAI line where GA dates, not model numbers, are the citable facts — no versioned voice model releases appear in the corpus beyond the 2.0 line [SECONDARY].

### Grok 5 — not shipped
- **Grok 5 had not shipped as of 2026-09-22** [UNVERIFIED].
- Elon Musk quote-posted support for the July 24 open-weights letter ("This has my full support") without signing it [SECONDARY] — vendor intent/secondary reporting, not a release commitment.
- The gadgetfee piece framing Grok 5 as an "AGI candidate" is secondary speculation about safety implications, not evidence of a release [SECONDARY].

### Corrections applied (wave-6 discipline)
- **Grok Code Fast 1:** the brief's 2026 date is corrected to **August 28, 2025** (codename Sonic); no 2026 successor exists in the corpus [SECONDARY].
- **Grok 4.4:** does not exist — xAI skipped it; any 4.4 release claim is fabrication [SECONDARY].
- **Grok 4.5 date:** sources give July 8/9; the task resolution is **July 8** — cite July 8 with the 8/9 nuance in a footnote, not a single-day assertion without qualification [SECONDARY].
- **Grok 5:** missed targets are vendor intent/secondary reporting, not a release; status as of 2026-09-22 is unshipped [UNVERIFIED].
- **STT/TTS dates:** Apr 15 (GA per one source) vs Apr 17–19 (per another) is rollout nuance — cite the range [SECONDARY].

### Benchmark-reading notes for the Grok line
- Grok 4.7's Terminal-Bench 4.0 score (26%) sits on the TB 4.0 scale (Sept 1–2, 2026 release) — **TB 4.0 scores are NOT comparable with TB 2.1-era scores**; cite the version whenever quoting [SECONDARY].
- Grok 4.7's DeepSWE v1.1 71.0% is [VENDOR] — do not mix with standardized boards [VENDOR].
- Grok 4.6's cache-read price ($0.50/M) and 4.3's 84% cache-hit discount show xAI pushing the industry-wide 2026 pattern of near-90% cache discounts to make long-context reasoning affordable [SECONDARY].


### New verified facts — expansion

### Grok 4 / 4 Heavy / 4 Fast (July–September 2025; details beyond base section)
- Grok 4 was released 2025-07-09 as a single-agent model with native tool use and real-time search; Grok 4 Heavy shipped the same day as a multi-agent ensemble configuration. [SECONDARY, S3][SECONDARY, S6][SECONDARY, S7]
- Grok 4 was trained with reinforcement learning at pretraining scale on xAI's Colossus supercluster using 200,000 GPUs — 100× the training of Grok 2, with RL compute alone reported at 10× any other model's. [SECONDARY, S2][SECONDARY, S44][SECONDARY, S45]
- On 2025-08-10/11 xAI made Grok 4 freely available to all users worldwide with "generous usage limits" for a limited time, days after OpenAI's GPT-5 launch. [SECONDARY, S4][SECONDARY, S70][SECONDARY, S71][SECONDARY, S72][SECONDARY, S73]
- The free rollout offered Auto mode (dynamic routing between heavier/lighter reasoning paths) and Expert mode (all prompts through the advanced model); free tier reports cite up to 5 prompts per 12 hours. [SECONDARY, S4][SECONDARY, S70][SECONDARY, S71][SECONDARY, S73]
- Grok 4 Heavy remained restricted to SuperGrok Heavy subscribers during the free period. [SECONDARY, S4][SECONDARY, S63][SECONDARY, S70][SECONDARY, S72]
- Subscription tiers at the time: Heavy at $300/month for high-usage professionals/researchers (later a $99/mo promo for new subscribers), SuperGrok at $30/month for creators and casual users; later additions: SuperGrok Lite $10/mo and SuperGrok Plus $100/mo. [SECONDARY, S5][SECONDARY, S63][SECONDARY, S60]
- Grok 4's launch introduced full audio interaction via the Eve voice agent, plus the Ani (anime-style) and Rudi (red panda) companion agents. [SECONDARY, S5 — single source]
- Grok 4 Fast launched September 2025 as a cost-optimized version supporting up to 2M-token context windows. [SECONDARY, S4 — single source]
- Vendor/third-party benchmark claims for Grok 4: 25.4% on Humanity's Last Exam without tools and 44.4% for Grok 4 Heavy with tools (the 38–44% range cited in other coverage sits between these two configurations); ARC-AGI-2 at 16.2%, the only model in three months above 10%. [SECONDARY, S6][SECONDARY, S47]
- Oracle integrated Grok 4 as a first-class model in Oracle Cloud Infrastructure, citing enterprise data extraction, code generation, and domain summarization. [SECONDARY, S6 — single source]

