---
id: ai-industry-kb-2026-wave6/08-mimo-and-xiaomi/overview
title: "§8. MiMo and Xiaomi"
domain: mimo-and-xiaomi
role: deep-dive
task: actor-profile
actors: ["DeepSeek", "Hugging Face", "Xiaomi"]
dates: ["2025-04-30", "2025-05-30", "2026-03", "2026-03-18", "2026-04", "2026-04-22", "2026-04-23", "2026-09", "2026-09-21"]
keywords: ["agents", "alignment", "apache", "attention", "context window", "cost", "deepseek", "fine-tuning", "licenses", "multimodal", "omni", "open weights"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3564, 3624]
section: "§8. MiMo and Xiaomi"
delta_of: ai-industry-kb-2026
sha256: 09bb20585ddac6357d2f6c2b33f26a2b444fe8bdb54f020b394005e058ddda84
---

# §8. MiMo and Xiaomi

Keywords: MiMo, Xiaomi, MiMo-7B, MiMo-V2-Pro, MiMo-V2-Omni, MiMo-V2.5, MiMo-V2.5-Pro, MiMo-V2.6-Pro, MiMo-V2.6-Flash, MiMo-V2.5-TTS, Hunter Alpha, Healer Alpha, Live RL, Human x Car x Home, open weights

## Summary
- **MiMo** is Xiaomi's AI model family. It opened with the **MiMo-7B** family on **2025-04-30** — MIT-licensed weights (repository code Apache 2.0), pretrained on ~25T tokens, with RL variants following (incl. RL-0530) [SECONDARY].
- The 2026 releases came in three generations: proprietary omnimodal flagships **V2-Pro / V2-Omni / V2-TTS (2026-03-18)** [SECONDARY]; open-weight **V2.5 / V2.5-Pro (2026-04-22)** — V2.5 at 310B, V2.5-Pro at 1.02T/42B active, both MIT with 1M context [SECONDARY]; and **V2.6-Pro / V2.6-Flash (2026-09-21/22)**, both MIT [SECONDARY].
- **V2.6-Pro** — natively omnimodal, 1.02T/42B active — was trained with ~6 days of public "Live RL" (~30 steps, ~750K trajectories, ~$2.62M cost) [SECONDARY]; V2.6-Flash is the cheaper sibling at 310B/15B active [SECONDARY].
- MiMo-V2.5-Pro is priced at $0.80/$3.20 per million input/output with an 80% cache discount ($0.16) [SECONDARY].
- MiMo is positioned inside Xiaomi's **"Human × Car × Home"** ecosystem strategy [SECONDARY].

## Key dated facts
### MiMo-7B — the MIT start
- **2025-04-30** — MiMo-7B family released [SECONDARY].
- Base pretrained on ~25T tokens; RL variants followed, including RL-0530 [SECONDARY].
- Weights under MIT; the repository code under Apache 2.0 — keep artifact scope explicit [SECONDARY].

### March 2026 — proprietary omnimodal flagships
- **2026-03-18** — MiMo-V2-Pro, V2-Omni, and V2-TTS launched, proprietary licenses [SECONDARY].
- **"Hunter Alpha"** was V2-Pro's stealth codename and was wrongly suspected to be DeepSeek V4 [SECONDARY].
- **"Healer Alpha"** was V2-Omni's stealth codename [SECONDARY].

### April 2026 — V2.5 generation, open weights
- **2026-04-22** — MiMo-V2.5 and V2.5-Pro released [SECONDARY].
- V2.5: 310B parameters, MIT [SECONDARY].
- V2.5-Pro: 1.02T total / 42B active, MIT [SECONDARY].
- 1M context window, with a 256K standard tier and extended 1M [SECONDARY].
- **2026-04-23** — V2.5-TTS-Series released: API-only, initially limited-time free [SECONDARY/VENDOR]; variants `mimo-v2.5-tts`, `mimo-v2.5-tts-voicedesign`, `mimo-v2.5-tts-voiceclone` [VENDOR].
- MiMo-V2.5-Pro API pricing: $0.80/$3.20 per million input/output; 80% off cache reads ($0.16) [SECONDARY].

### September 2026 — V2.6 generation
- **2026-09-21/22** — V2.6-Pro announced (Sept 21) and covered via the Xiaomi blog (Sept 22) [SECONDARY]; V2.6-Flash released the same window [SECONDARY].
- V2.6-Pro is a distinct generation from V2.5-Pro even though both use the 1.02T/42B profile [SECONDARY].
- V2.6-Pro is natively omnimodal and used ~6 days of public **"Live RL"** — 30 steps, ~750K trajectories, ~$2.62M cost [SECONDARY].
- V2.6-Flash: 310B / 15B active, MIT [SECONDARY].

### Coverage and positioning context
- The V2.6 launch was framed by secondary coverage as "the top open-weights model in the world" alongside the cheaper V2.6-Flash [SECONDARY] — ranking language that belongs to the press, not to a standardized board; the AA v4.3 46.32 figure is the only version-pinned independent score in this section.
- InfoWorld framed the MIT-licensed MiMo models around long-running AI agents — the positioning thread that runs from MiMo-7B (2025) through the 2026 generations [SECONDARY].
- The stealth-codename episode ("Hunter Alpha" misidentified as DeepSeek V4) shows how much market attention attaches to unidentified strong open models in 2026; the correction was that it was Xiaomi's V2-Pro all along [SECONDARY].
- MiMo's documentation hub (mimo.mi.com) and the Xiaomi-hosted model page (mimo.xiaomi.com/mimo-v2-6) are the vendor-primary access points for the V2.6 generation [VENDOR].
- V2.5 vs V2.6 generation discipline: V2.5-Pro and V2.6-Pro share the 1.02T/42B profile but are distinct generations — do not treat V2.6-Pro as a re-release or checkpoint of V2.5-Pro [SECONDARY].
- V2.5-TTS note: three API-only voice variants launched limited-time free on Apr 23, 2026; voice is the product line where MiMo stays API-gated while text weights go open [SECONDARY].
- The 256K standard / 1M extended context tiering on V2.5 predates the 2026 industry-wide move to context-window pricing cliffs; MiMo's rate card instead prices V2.5-Pro flat at $0.80/$3.20 with an 80% cache discount [SECONDARY].


### New verified facts — expansion

### MiMo-7B-RL-0530 post-training update (2025-05-30)
- MiMo-7B-RL-0530 was posted to Hugging Face on 2025-05-30, one month after the MiMo-7B base release, as an updated reasoning checkpoint rather than a new base model. [VENDOR, S1][SECONDARY, S2]
- The SFT (supervised fine-tuning) dataset behind the RL update grew from approximately 500K instances to approximately 6M instances. [VENDOR, S1][SECONDARY, S2]
- The RL training context window increased from 32K tokens to 48K tokens between the base release and the 0530 update. [VENDOR, S1]
- AIME 2024 accuracy rose from 68.2 to 80.1 between the base model and the 0530 RL checkpoint. [VENDOR, S1][SECONDARY, S2]
- AIME 2025 accuracy rose from 55.4 to 70.2 across the same update. [VENDOR, S1][SECONDARY, S2]
- LiveCodeBench v5 rose from 57.8 to 60.9. [VENDOR, S1]
- LiveCodeBench v6 rose from 49.3 to 52.2. [VENDOR, S1]
- GPQA Diamond rose from 54.4 to 60.6. [VENDOR, S1]
- The RL stage consumed approximately 130,000 curated math and code problems. [SECONDARY, S2][SECONDARY, S4]
- MiMo-VL-7B, the vision-language variant, was trained on approximately 2.4T tokens across four distinct training stages (projector warmup, vision-language alignment, general multimodal pretraining, long-context SFT). [VENDOR, S35][SECONDARY, S36][SECONDARY, S2]
- A MiMo-Audio-7B audio-language variant also exists, built for voice conversion, style transfer, and speech editing. [SECONDARY, S2][SECONDARY, S4]

