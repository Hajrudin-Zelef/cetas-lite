---
id: ai-industry-kb-2026-wave6/09-meta-llama-muse-spark-and-muse-glimmer/llama-4-behemoth-status-announced-never-released
title: "Llama 4 Behemoth status (announced, never released)"
domain: meta-llama-muse-spark-and-muse-glimmer
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Apple", "Meta", "Microsoft"]
dates: ["2024-07-23", "2024-09-25", "2025-04", "2025-04-05", "2025-06", "2026-01-26", "2026-02-05", "2026-08-05", "2026-08-07", "2026-08-10", "2026-08-31", "2026-09-02"]
keywords: ["llama", "agent", "agents", "apache", "benchmark", "benchmarks", "consumer", "context window", "distillation", "gpus", "inference", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4169, 4245]
section: "§9. Meta: Llama, Muse Spark, and Muse Glimmer"
sha256: 4a059f7df251b929ec7f18903b67c1ce5c3d8014b72d7182f74511f211d09382
---

# Llama 4 Behemoth status (announced, never released)

### Llama 4 Behemoth status (announced, never released)
- Llama 4 Behemoth was announced 2025-04-05 alongside Scout and Maverick but was never publicly released. [SECONDARY, S9][SECONDARY, S11]
- Preview specification: approximately 2T total parameters / 288B active, 16 experts. [SECONDARY, S9][SECONDARY, S11]
- Reported training data: 30T+ tokens of multimodal data — more than double the Llama 3 pretraining mixture. [SECONDARY, S9][SECONDARY, S39][SECONDARY, S40]
- Internal release targets slipped from April 2025 to June 2025 and then to fall 2025 or later. [SECONDARY, S9][SECONDARY, S10]
- Reported cause of delay: internal concern that Behemoth's capability gains over released models were not substantial enough; WSJ reporting added executive frustration and possible restructuring. [SECONDARY, S9][SECONDARY, S41][SECONDARY, S42]
- Claims that Behemoth was "shelved in favor of Llama 5" are speculation, not established fact. [UNVERIFIED]
- No LlamaCon 2026 event was confirmed as of the research date (no official second-edition date publicly exists). [SECONDARY, S11][SECONDARY, S39]

### Muse Spark 1.2 specification deltas (released 2026-08-05; details beyond base section)
- Muse Spark 1.2 is a coding-focused update to the Spark line, not a new base model. [SECONDARY, S13][SECONDARY, S15]
- Context window: 1M tokens. [SECONDARY, S13][SECONDARY, S15]
- Input modalities: text, image, video, audio, and PDF. [SECONDARY, S13][SECONDARY, S15]
- Two operating modes: Instant and Thinking. [SECONDARY, S13][SECONDARY, S15]
- Supports parallel tool calls and structured output. [SECONDARY, S13][SECONDARY, S15]
- Vendor-reported benchmarks: Terminal-Bench 2.1 at 82.9% and DeepSWE v1.1 at 59.3%. [VENDOR, S13 — single source]
- Those benchmark entries were not found on public leaderboards when checked 2026-08-07, two days after launch. [SECONDARY, S12 — single source]
- Standard-tier pricing does not permit Meta to use prompts/completions for model improvement; the cheaper Contributor tier explicitly permits it. [SECONDARY, S16][SECONDARY, S13]

### Muse Code agent CLI (beta; details beyond base section)
- Muse Code runs persistent asynchronous background agents. [SECONDARY, S14][SECONDARY, S16]
- It uses isolated Git worktrees to run parallel tasks without contaminating the working tree. [SECONDARY, S14][SECONDARY, S16]
- It keeps a crash-safe append-only local event log with replay-exact session recovery. [SECONDARY, S14][SECONDARY, S72][COMMUNITY, S73]
- It ships `/plan` and `/grill` slash commands. [SECONDARY, S14 — single source]
- Initial platform support: macOS and Linux (single install command; Meta developer account required; no official Windows support; no GUI or IDE integration). [SECONDARY, S14][SECONDARY, S72][COMMUNITY, S73][SECONDARY, S75]
- Muse Code left beta on 2026-08-31 with subscription plans ranging $5–$50/month; the standard tier is $1.25/$4.25 input/output ($0.15 cached) with 3,000 requests/min and 4M tokens/min caps; the contributor tier $0.10/$0.20 ($0.002 cached), capped at 60 requests/min, in exchange for training consent. [SECONDARY, S76][SECONDARY, S74][SECONDARY, S75]

### Muse Glimmer 30B specification (released 2026-08-10; details beyond base section)
- Glimmer 30B is a dense 29.6B-parameter language model paired with a 1.8B-parameter ViT-G/14 perception encoder. [SECONDARY, S17][SECONDARY, S19]
- It was distilled from Muse Spark 1.2 (per release coverage) using logit distillation. [SECONDARY, S17][SECONDARY, S19][SECONDARY, S64]
- Context window: 120K+ tokens. [SECONDARY, S17][SECONDARY, S19]
- Language support: 100+ languages. [SECONDARY, S17][SECONDARY, S19][SECONDARY, S65]
- A quantized build fits under 20GB, targeting 24–32GB consumer GPUs. [SECONDARY, S17][SECONDARY, S18][SECONDARY, S64][SECONDARY, S65]
- DFlash, its speculative decoding scheme, proposes 16-token blocks verified in parallel by the main model. [SECONDARY, S17][SECONDARY, S64][SECONDARY, S65]
- Vendor-reported throughput on RTX 5090: 74.9 → 233.4 tokens/second (3.1×) with DFlash. [VENDOR, S17][SECONDARY, S64][SECONDARY, S65]
- Vendor-reported speedups: 1.9× on Apple M5 Max, 1.6× on M4 Max; two other reports citing Meta's tests give 1.8× and 1.5× respectively — a minor vendor-transmission discrepancy preserved in the ledger. [VENDOR, S17][SECONDARY, S64][SECONDARY, S65]
- Vendor-reported benchmarks: MCP Atlas 75.5 (vs Gemma4-31B 54.2 and Qwen3.6-27B 62.5); Terminal-Bench 2.1 51.7; OSWorld-Verified 65.9 (trailing Qwen3.6-27B's 75.6); plus AIME 2026 94.7% and SWE-Bench Pro 51.2%. [VENDOR, S17][SECONDARY, S66][SECONDARY, S67]
- Open weights and inference assets are Apache 2.0 licensed, but training data and training code were not released — Glimmer is open-weight, not reproducibly open-source. [SECONDARY, S17][SECONDARY, S20]

### Muse Spark 1.3 specification (released 2026-09-02; details beyond base section)
- Context window: 1,048,576 tokens, identical to 1.2. [SECONDARY, S22][SECONDARY, S24]
- Reported maximum output: 943,718 tokens. [SECONDARY, S22 — single source]
- Vendor-reported internal efficiency vs 1.2: approximately 20% fewer tool calls and 25% fewer tokens per task. [VENDOR, S23][SECONDARY, S25]
- Cached-input pricing in the standard tier: $0.15 per million tokens. [SECONDARY, S22][SECONDARY, S44][SECONDARY, S45]
- Parameter count and architecture remain undisclosed. [DIRECTIONAL]
- Vendor benchmark card, dated to the 2026-09-02 launch: DeepSWE v1.1 75.4. [VENDOR, S22][SECONDARY, S24]
- Vendor benchmark card: Terminal-Bench 2.1 88.8. [VENDOR, S22][SECONDARY, S24]
- Vendor benchmark card: SWE-Atlas Codebase QnA 59.4. [VENDOR, S22 — single source]
- Vendor benchmark card: MRCR v2 256K–512K 98.5. [VENDOR, S22 — single source]
- Vendor benchmark card: MRCR v2 512K–1M 98.1. [VENDOR, S22 — single source]
- Vendor benchmark card: OSWorld 2.0 66.9 at `max` reasoning, 57.2 at `xhigh`. [VENDOR, S22 — single source]
- Vendor benchmark card: GDPval-AA v2 Elo 1754 at `max`, 1709 at `xhigh`. [VENDOR, S22 — single source]
- Vendor benchmark card: JobBench 64.9 at `max` vs 61.2 at `xhigh`. [VENDOR, S22 — single source]
- The `max` vs `xhigh` reasoning-mode distinction materially changes reported scores and must be pinned with every Spark 1.3 benchmark citation. [VENDOR, S22]
- Reasoning-mode changes between 1.2 and 1.3 explain part of the generational benchmark improvement, independent of base-model capability. [SECONDARY, S24 — single source]

### Meta AI app, Llama API, and platform (new to this section)
- The standalone Meta AI app launched at LlamaCon 2025 as a voice-first experience powered by Llama 4. [SECONDARY, S26][SECONDARY, S27]
- The app's Discover feed surfaces prompts and content users voluntarily share. [SECONDARY, S27][SECONDARY, S29]
- App personalization can draw on information users shared across Facebook and Instagram. [SECONDARY, S29 — single source]
- The Llama API entered limited preview at LlamaCon 2025. [SECONDARY, S29][SECONDARY, S28]
- Meta reported 1.2B cumulative Llama downloads at LlamaCon 2025. [VENDOR, S26 — single source]
- The Vibes AI video app entered standalone testing on 2026-02-05, initially in Brazil and Mexico. [SECONDARY, S30][SECONDARY, S31]
- Vibes' freemium plan remained unpriced with no public rollout date as of the research date. [SECONDARY, S30 — single source]
- Meta announced premium subscription testing for Instagram, Facebook, and WhatsApp on 2026-01-26/27 (confirmed to TechCrunch; distinct from Meta Verified). [SECONDARY, S32][SECONDARY, S50][SECONDARY, S51]


### Llama 3.1 / 3.2 supplementary specification sheets
- **Llama 3.1 405B** — Released 2024-07-23; 405B dense parameters; 126 layers; 128K context. [SECONDARY, S1][SECONDARY, S3]
- Llama 3.1 405B was the first frontier-scale open-weights model at its release, defining the "open frontier" product category Meta then exploited with the Llama 4 MoE line. [SECONDARY, S3]
- Llama 3.1 8B: 32 layers; Llama 3.1 70B: 80 layers — the layer counts scale sublinearly with parameters across the family. [SECONDARY, S1]
- **Llama 3.2 1B/3B** — Released 2024-09-25; text-only edge models designed for on-device deployment. [SECONDARY, S1]
- **Llama 3.2 11B/90B** — Released 2024-09-25; vision-capable variants accepting image input alongside text. [SECONDARY, S1]
- Llama 3.2 context: 128K tokens across all four sizes. [SECONDARY, S1]
- Llama 3.2 reported training data: ~9T tokens, down from the ~15T reported for Llama 3 — reflecting a shift toward data quality and multimodal mixture over raw token scale. [SECONDARY, S1]
- The Llama 3 family license (Llama 3 Community License) carried the same 700M-monthly-active-user commercial threshold later retained in the Llama 4 Community License. [SECONDARY, S3]

