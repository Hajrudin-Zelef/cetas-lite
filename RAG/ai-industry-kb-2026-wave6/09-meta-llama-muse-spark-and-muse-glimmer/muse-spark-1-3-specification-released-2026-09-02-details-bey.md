---
id: ai-industry-kb-2026-wave6/09-meta-llama-muse-spark-and-muse-glimmer/muse-spark-1-3-specification-released-2026-09-02-details-bey
title: "Muse Spark 1.3 specification (released 2026-09-02; details beyond base section)"
domain: meta-llama-muse-spark-and-muse-glimmer
role: deep-dive
task: actor-profile
actors: ["Meta"]
dates: ["2024-07-23", "2024-09-25", "2026-01-26", "2026-02-05", "2026-09-02"]
keywords: ["muse", "muse spark", "benchmark", "context window", "license", "llama", "moe", "multimodal", "parameters", "pricing", "reasoning", "research"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4208, 4245]
section: "§9. Meta: Llama, Muse Spark, and Muse Glimmer"
delta_of: ai-industry-kb-2026
sha256: efa595a00ae0847c98053c6aba523dd8a1d48e57e6810a1c12cc3ccac66e31b1
---

# Muse Spark 1.3 specification (released 2026-09-02; details beyond base section)

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

