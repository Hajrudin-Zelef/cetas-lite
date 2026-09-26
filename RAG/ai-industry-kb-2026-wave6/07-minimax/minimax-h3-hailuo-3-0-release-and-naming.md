---
id: ai-industry-kb-2026-wave6/07-minimax/minimax-h3-hailuo-3-0-release-and-naming
title: "MiniMax H3 / Hailuo 3.0 — release and naming"
domain: minimax
role: deep-dive
task: actor-profile
actors: ["Anthropic", "China", "DeepSeek", "EU", "Hugging Face", "MiniMax", "Moonshot", "Stability AI", "United States"]
dates: ["2025-10-27", "2026-02-12", "2026-03-18", "2026-05-26", "2026-06-01", "2026-07-31", "2026-08-02", "2026-08-03", "2026-08-08", "2026-08-26", "2026-09", "2026-09-19"]
keywords: ["attention", "attribution", "benchmarks", "compute", "consumer", "decode", "deepseek", "distillation", "distribution", "gguf", "gpu", "int4"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3296, 3365]
section: "§7. MiniMax"
delta_of: ai-industry-kb-2026
sha256: 35d55ed7abf7524d534be43ca6cb80ab7fb6810976408bdc12a82aec2175f9d4
---

# MiniMax H3 / Hailuo 3.0 — release and naming

### MiniMax H3 / Hailuo 3.0 — release and naming
- **Announced July 31, 2026** (teased under #MiniMaxH3 the day before); **H3 is the official model name, Hailuo 3.0 / Hailuo 03 is the consumer-app alias**, and the API labels the endpoint **"Hailuo-03"**; unrelated to Kuaishou's Kling O3 [SECONDARY] (huggingface.co/blog/ResterChed/minimax-h3-hailuo-3-0; pixo.video; evolink.ai).
- **Weights released on Hugging Face August 3, 2026**; **ComfyUI support merged the same day**; within 24 hours the community had **GGUF, INT4, and NVFP4** quants up and a ComfyUI maintainer rendering on an **RTX 3060** [SECONDARY] (runaihome.com, 2026-08-08).
- **33B parameters**; omni-modal video model: text, image, video, and audio treated as one unified context; generates **4–15 second clips at up to 2K (1440p), 24 fps, with native stereo audio in a single pass** — dialogue, effects, and room tone generated with the picture, plus **voice transfer** from a reference recording [SECONDARY] (runaihome.com; huggingface.co/blog/ResterChed; pixo.video).
- Successor to **Hailuo 2.3** (which topped out at 1080p/~10s from prompt or single image, silent); H3 adds mixed image-video-audio references, native stereo audio, voice transfer, and sentence-level editing [SECONDARY] (huggingface.co/blog/ResterChed; auraai.app).
- **Six aspect ratios** including 21:9 and 4:3 (Hailuo 2.3 offered 16:9, 9:16, 1:1) [SECONDARY] (auraai.app).
- Third-party distribution within the first week across creative platforms (EvoLink, Pixo, Aura AI) — "raw access to this model is spreading fast" [SECONDARY] (pixo.video).
- **Stability AI's H3Max** is a separate **paid speed-optimized fine-tune of H3** — a cross-vendor derivative worth tracking separately from the base model [SECONDARY] (github.com/minecraft9101010/ai-tools-mcp card, 2026-08-02).

### MiniMax H3 — license (conflict recorded)
- Announced license: **MiniMax Community License** — commercial use for organizations **under $20M revenue**, with **prominent attribution** [SECONDARY] (huggingface.co/blog/ResterChed, citing Artificial Analysis; justbeingresourceful.com).
- RunAIHome reports the H3 license **excludes the US, EU, UK, and South Korea from local deployment** of the weights [SECONDARY] (runaihome.com, 2026-08-08 — single source for the territory list; the exact license file was not reviewed in this pass, mark [UNVERIFIED] for legal reliance).
- The existing §7 records a general geographic-license exclusion for MiniMax; this file adds the **specific claimed territory list (US/EU/UK/SK)** and the **$20M revenue cap** as the two concrete clauses in circulation [DIRECTIONAL].
- Hardware note: quantized builds run on a single consumer GPU — NVFP4 build finishes a 10-second clip in **~175s on an RTX 5090 32GB**; community renders observed on RTX 3060-class hardware [SECONDARY] (runaihome.com).

### MiniMax H3 — API pricing and input limits
- Official pricing: **$0.13/second of 2K video ($7.80/min)**; **768p tier $0.09/s** (closed beta); the 768p tier costs **8/13 of the 2K rate** (a 4-second 768p clip starts at $0.304 vs $0.494 for 2K) [SECONDARY] (huggingface.co/blog/ResterChed; evolink.ai).
- Reference pricing: **reference audio free**; **first 5 reference images free, then $0.04 each**; **input video billed by duration at the output tier rate** [SECONDARY] (huggingface.co/blog/ResterChed).
- Input limits: **≤9 reference images, ≤3 reference video clips (2–15s each, ≤15s total), ≤3 reference audio clips** (audio references can be used on their own per one source — the HF FAQ says audio cannot be sent alone; record the discrepancy as [UNVERIFIED] detail), **12 files max**, prompt **≤7,000 characters** [SECONDARY] (huggingface.co/blog/ResterChed; evolink.ai).
- MiniMax's own claim: 2K generation costs **less than a third of mainstream competitors**, 768p less than half of competitors' 720p — unnamed comparators; one independent breakdown puts H3 2K at ~$7.80/min vs **~$20–22/min for Kling 3.0 and Seedance 2.0 at 1080p** [VENDOR claim; SECONDARY breakdown] (justbeingresourceful.com, 2026-08-26).
- Third-party hosted pricing (for reference, not vendor list): EvoLink 4s/10s/15s 2K clips at $0.494/$1.235/$1.8525; Aura AI flat 2–3 credits per clip across H3 Turbo / H3 Max Turbo / H3 Turbo Reference [SECONDARY] (evolink.ai; auraai.app).

### MiniMax in the Anthropic allegations
- MiniMax is one of the three Chinese labs named in Anthropic's alleged large-scale distillation campaigns (with DeepSeek and Moonshot): **16M+ interactions through ~24K fraudulent accounts** overall, per Anthropic [SECONDARY] (computerworld.com; infoworld.com).
- Same framing discipline as §5: **Anthropic's allegation**, not established fact [DIRECTIONAL].

### License-context reading
- The September 2026 aifoss.dev licensing survey's framework applies directly to MiniMax: "open weights" ≠ commercial-use rights; the survey's blacklist categories (revenue caps like the $20M clause, MAU caps, regional exclusions) are exactly the clause types MiniMax uses — read the checkpoint LICENSE file and save a dated copy, since license files change in place [SECONDARY] (aifoss.dev, 2026-09-19).

## Figures and metrics
| Model | Params (total/active) | Context | License |
|---|---|---|---|
| M2 (2025-10-27) | 230B / 10B | — | MIT [SECONDARY] |
| M2.5 (2026-02-12) | 229.9B / 9.8B | ~192–196K | Modified-MIT w/ UI attribution [SECONDARY] |
| M2.7 (2026-03-18 ann.) | 230B / ~10B | 200K | Non-commercial modified [SECONDARY] |
| M3 (2026-06-01) | 428B / ~23B | 1M | MiniMax Community License [SECONDARY] |
| H3 (2026-07-31) | 33B dense | — | Open weights (Hailuo line) [SECONDARY] |

- M3 API pricing: $0.60/$2.40 per million input/output standard; $0.30/$1.20 launch promotion [SECONDARY].
- MSA speedup claims (all [VENDOR], different measurements): 9.7× prefill / 15.6× decode (R&D-lead diagram) → rounded 9×/15× in the shipped card; 14.2×/7.6× on a 109B testbed technical paper.
- M2.7 benchmarks (vendor-reported): 56.22% SWE-bench Pro, 57.0% Terminal-Bench 2.1 [SECONDARY].


### New verified metrics — expansion (continued — H3 figures)

- July 31 announce; Aug 3 weights; 33.1B H3-Base @ 768p; 2K/15s via API-only Regenerate-2K [SECONDARY].
- $7.80/min vs $20.16/$22.45 (Kling 3.0/Seedance 2.0); 2K <1/3, 768p <1/2 of competitors [VENDOR/SECONDARY].
- AA: #1 video editing; top-3 T2V/I2V [SECONDARY].
- $20M revenue free tier; 4-territory exclusion; Sept 2025 lawsuit; May 26, 2026 discovery ruling [SECONDARY].


### New verified metrics — expansion (continued — M3 launch-day figures)

- Launch day: +5% → −12.38% close, HK$907.5 peak, HK$1.49B turnover [SECONDARY].
- Kilo audit: 13/17 @ $0.07 vs Opus 4.8 13/17 @ $1.30 / 15/17 @ $3.39 [SECONDARY].
- CUDA kernel: 7.6% → 71.3% utilization (9.4×), 147 submissions, ~2,000 tool calls, 24h [VENDOR].
- ICLR 2025 reproduction: ~12h, 18 commits, 23 figures [VENDOR].
- MSA: 1/20th compute at 1M; 15.6× decode; 9×+ prefill; 4×+ over Flash-Sparse-Attention [VENDOR].
- Token Plan: $20/$50/$120 per month [SECONDARY].


### New verified metrics — expansion (continued — M2.5/M2.7 figures)

- M2.5: 230B/10B; 197K ctx; SWE-Bench Verified 80.2; Droid 79.7; OpenCode 76.1; BrowseComp 76.3; BFCL 76.8 [VENDOR]; 1.8pp behind Opus 4.7 at ~1/17 price [COMMUNITY].
- M2.7: 10B active; SWE-Pro 56.22; Terminal Bench 2 57.0; SWE-Bench Verified 78; SWE Multilingual 76.5; GDPval-AA Elo 1495; AA-Omniscience +1 (vs M2.5 −40); hallucination 34% (vs 46%/50%); MM Claw 97%; medal rate 66.6%; 204.8K ctx / 131.1K out; ~3× Opus 4.6 throughput [VENDOR/COMMUNITY].
- Speech 2.6 / Music 2.6: 40 languages, 250ms latency [COMMUNITY].


### New verified metrics — expansion

