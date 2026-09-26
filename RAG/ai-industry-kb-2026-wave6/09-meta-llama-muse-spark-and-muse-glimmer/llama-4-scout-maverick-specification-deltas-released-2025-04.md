---
id: ai-industry-kb-2026-wave6/09-meta-llama-muse-spark-and-muse-glimmer/llama-4-scout-maverick-specification-deltas-released-2025-04
title: "Llama 4 Scout / Maverick specification deltas (released 2025-04-05)"
domain: meta-llama-muse-spark-and-muse-glimmer
role: deep-dive
task: actor-profile
actors: ["AWS", "Alibaba", "Apple", "EU", "Hugging Face", "Meta", "Microsoft", "United States"]
dates: ["2025-04", "2025-04-05", "2025-06", "2026-08-05", "2026-08-07", "2026-08-10", "2026-08-31"]
keywords: ["llama", "scout", "agent", "agents", "apache", "aws", "bedrock", "benchmark", "benchmarks", "consumer", "context window", "distillation"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4155, 4207]
section: "§9. Meta: Llama, Muse Spark, and Muse Glimmer"
delta_of: ai-industry-kb-2026
sha256: d19532a004b9d05feb4767e55bd07f6c7273b0f4034d6ce97976cfeb91a9b242
---

# Llama 4 Scout / Maverick specification deltas (released 2025-04-05)

### Llama 4 Scout / Maverick specification deltas (released 2025-04-05)
- Llama 4 Scout and Maverick were released 2025-04-05 as natively multimodal, early-fusion mixture-of-experts models. [SECONDARY, S5][SECONDARY, S8]
- Llama 4 Scout: 109B total parameters with 17B active, 16 experts. [SECONDARY, S5][SECONDARY, S7]
- Llama 4 Maverick: approximately 400B total parameters with 17B active, 128 experts. [SECONDARY, S5][SECONDARY, S7]
- Scout's context window is 10M tokens. [SECONDARY, S5][SECONDARY, S8]
- Maverick's context window is 1M tokens. [SECONDARY, S5][SECONDARY, S8]
- Scout can fit on a single H100 GPU using INT4 quantization, per Meta's launch materials. [SECONDARY, S5][SECONDARY, S40]
- BF16 and FP8 checkpoints of the Llama 4 models were also reported. [SECONDARY, S5 — single source]
- Same-day availability covered Hugging Face, AWS Bedrock, and Azure. [SECONDARY, S6][SECONDARY, S7]
- Scout and Maverick were integrated into Meta AI on WhatsApp, Messenger, and Instagram across 40 countries at launch (free downloads via Llama.com/Hugging Face; EU companies excluded by license). [SECONDARY, S6][SECONDARY, S63][SECONDARY, S64][SECONDARY, S65]
- The multimodal rollout was initially limited to US English. [SECONDARY, S6][SECONDARY, S63][SECONDARY, S64]
- The LMArena submission that generated controversy was the experimental checkpoint `Llama-4-Maverick-03-26-Experimental`, not the public release checkpoint; it scored Elo 1417 and ranked 2nd, while the unmodified Maverick later ranked below months-old rivals. [SECONDARY, S8][SECONDARY, S36][SECONDARY, S38]
- Scout's training provenance is disputed: some summaries report both Scout and Maverick partially distilled from Behemoth, while another reports Maverick was co-distilled and Scout trained from scratch. [SECONDARY, S5 vs S6 — documented contradiction]

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

