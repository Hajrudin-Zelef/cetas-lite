---
id: ai-industry-kb-2026-wave6/09-meta-llama-muse-spark-and-muse-glimmer/meta-ai-platform-supplementary-facts
title: "Meta AI platform supplementary facts"
domain: meta-llama-muse-spark-and-muse-glimmer
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "EU", "Hugging Face", "Meta", "Microsoft", "OpenAI"]
dates: ["2024-12-06", "2025-04-05", "2026-01", "2026-01-26", "2026-02-05", "2026-08", "2026-08-05", "2026-08-10", "2026-09-02"]
keywords: ["agent", "agentic", "apache", "aws", "bedrock", "benchmarks", "compute", "consumer", "context window", "distillation", "distribution", "fp8"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4246, 4314]
section: "§9. Meta: Llama, Muse Spark, and Muse Glimmer"
delta_of: ai-industry-kb-2026
sha256: 9825e5a460b09b4af7956235e7822ddf0cacdd67af8b837a4c412a78985b203f
---

# Meta AI platform supplementary facts

### Meta AI platform supplementary facts
- LlamaCon 2025 was Meta's first dedicated AI developer conference, and its agenda centered on undercutting OpenAI on API pricing and open-model distribution. [SECONDARY, S26][SECONDARY, S28]
- The Meta AI app's voice-first design made it the first Meta surface where voice, not text, was the default interaction mode. [SECONDARY, S27]
- The Discover feed's voluntary-sharing model means shared prompts are user-opted-in content, not telemetry — a distinction relevant to the contributor-pricing data-use debate. [SECONDARY, S27][SECONDARY, S29]
- The Llama API's limited-preview status at LlamaCon 2025 positioned it as a direct competitor to OpenAI's and Anthropic's developer platforms rather than a research artifact. [SECONDARY, S29][SECONDARY, S28]
- Vibes was spun out of the Meta AI app's AI-video tab into a standalone app, following the pattern of unbundling successful in-app features. [SECONDARY, S30]
- The Vibes standalone test began 2026-02-05 in Brazil and Mexico — emerging-market-first rollout, consistent with Meta's distribution strategy for new consumer AI surfaces. [SECONDARY, S30][SECONDARY, S31]
- Vibes' freemium plan was unpriced with no public rollout date as of the research date, indicating monetization was still undecided seven months into testing. [SECONDARY, S30]
- The January 2026 premium-subscription tests covered Instagram, Facebook, and WhatsApp simultaneously — the first time Meta tested paid tiers across all three flagship apps at once. [SECONDARY, S32]
- The subscription tests were announced 2026-01-26/27, roughly one week before the Vibes standalone test began, suggesting a coordinated Q1 2026 monetization push. [SECONDARY, S32][SECONDARY, S30]

### Consolidated specification sheets (full-field reference)
- **Llama 3.3 70B** — Released 2024-12-06/07 (timezone-dependent dating). [SECONDARY, S2][SECONDARY, S4]
- Architecture: dense transformer, GQA; 70B parameters; layer count not publicly disclosed. [SECONDARY, S4]
- Context: 128K tokens; text-only; instruction-tuned release (no base-weights release for 3.3). [SECONDARY, S2][SECONDARY, S4]
- Languages: 8 supported. [SECONDARY, S2]
- Training tokens/compute: not publicly disclosed. [DIRECTIONAL]
- License: Llama 3 Community License (base §9 anchor; 700M-MAU threshold and EU multimodal exclusion apply to the family). [SECONDARY, S3]
- Benchmarks (vendor): MMLU 86.0; HumanEval 88.4. [VENDOR, S2]
- **Llama 4 Scout** — Released 2025-04-05. [SECONDARY, S5][SECONDARY, S8]
- Architecture: natively multimodal early-fusion MoE; 16 experts. [SECONDARY, S5][SECONDARY, S8]
- Parameters: 109B total / 17B active. [SECONDARY, S5][SECONDARY, S7]
- Context: 10M tokens. [SECONDARY, S5][SECONDARY, S8]
- Quantization: INT4 build fits a single H100; BF16/FP8 checkpoints reported. [SECONDARY, S5]
- Training tokens/compute: not publicly disclosed; distillation-from-Behemoth disputed (see contradiction). [DIRECTIONAL]
- License: Llama 4 Community License (base §9 anchor). [SECONDARY, S6]
- Distribution: Hugging Face, AWS Bedrock, Azure (same-day); WhatsApp/Messenger/Instagram in 40 countries. [SECONDARY, S6][SECONDARY, S7]
- **Llama 4 Maverick** — Released 2025-04-05. [SECONDARY, S5][SECONDARY, S8]
- Architecture: natively multimodal early-fusion MoE; 128 experts. [SECONDARY, S5][SECONDARY, S8]
- Parameters: ~400B total / 17B active. [SECONDARY, S5][SECONDARY, S7]
- Context: 1M tokens. [SECONDARY, S5][SECONDARY, S8]
- Training tokens/compute: not publicly disclosed. [DIRECTIONAL]
- License: Llama 4 Community License (base §9 anchor). [SECONDARY, S6]
- Distribution: Hugging Face, AWS Bedrock, Azure (same-day). [SECONDARY, S6][SECONDARY, S7]
- **Llama 4 Behemoth** — Announced 2025-04-05; never released. [SECONDARY, S9][SECONDARY, S11]
- Preview spec: ~2T total / 288B active; 16 experts; 30T+ training tokens. [SECONDARY, S9][SECONDARY, S11]
- Status: delayed over capability-concern reports; "shelved" unconfirmed. [SECONDARY, S9][UNVERIFIED]
- **Muse Spark 1.2** — Released 2026-08-05 (base §9 anchor). [SECONDARY, S13]
- Type: coding-focused update, not a new base model. [SECONDARY, S13][SECONDARY, S15]
- Context: 1M tokens; input modalities text/image/video/audio/PDF. [SECONDARY, S13][SECONDARY, S15]
- Modes: Instant and Thinking; parallel tool calls; structured output. [SECONDARY, S13][SECONDARY, S15]
- License/pricing: closed model; standard and contributor tiers (base §9 anchor). [SECONDARY, S16]
- Benchmarks (vendor): Terminal-Bench 2.1 82.9%; DeepSWE v1.1 59.3%. [VENDOR, S13]
- **Muse Glimmer 30B** — Released 2026-08-10 (base §9 anchor). [SECONDARY, S17][SECONDARY, S18]
- Architecture: dense 29.6B LM + 1.8B ViT-G/14 encoder; distilled from Spark via logit distillation. [SECONDARY, S17][SECONDARY, S19]
- Context: 120K+ tokens; 100+ languages. [SECONDARY, S17][SECONDARY, S19]
- Local footprint: quantized build <20GB; targets 24–32GB GPUs. [SECONDARY, S17][SECONDARY, S18]
- Decoding: DFlash with 16-token proposal blocks. [SECONDARY, S17]
- License: Apache 2.0 for weights and inference assets; training data/code not released. [SECONDARY, S17][SECONDARY, S20]
- Benchmarks (vendor): MCP Atlas 75.5; Terminal-Bench 2.1 51.7; OSWorld-Verified 65.9. [VENDOR, S17]
- **Muse Spark 1.3** — Released 2026-09-02 (base §9 anchor). [SECONDARY, S22][SECONDARY, S23]
- Architecture/parameters: undisclosed. [DIRECTIONAL]
- Context: 1,048,576 tokens; max output 943,718. [SECONDARY, S22][SECONDARY, S24]
- Efficiency vs 1.2 (vendor): ~20% fewer tool calls, ~25% fewer tokens. [VENDOR, S23][SECONDARY, S25]
- Pricing: cached input $0.15/M (standard tier); base pricing in base §9. [SECONDARY, S22]
- Benchmarks: full dated vendor card listed above; `max` vs `xhigh` modes pinned. [VENDOR, S22]


### Muse brand and product-line facts
- "Muse" is Meta's AI agent product brand spanning Muse Spark (flagship coding model), Muse Glimmer (local 30B agent model), and Muse Code (terminal agent CLI). [SECONDARY, S14][SECONDARY, S17]
- Muse Spark is positioned as an agentic coding model rather than a general chat model; both the 1.2 update and 1.3 release were framed around coding-agent efficiency. [SECONDARY, S23][SECONDARY, S13]
- Muse Glimmer's release (2026-08-10) came five days after Spark 1.2 (2026-08-05), making August 2026 a two-release month for the Muse line. [SECONDARY, S17][SECONDARY, S13]
- Muse Spark 1.3 followed Glimmer by 23 days (2026-09-02), completing a three-release sequence in under a month. [SECONDARY, S22][SECONDARY, S17]
- The Spark 1.2 → 1.3 efficiency claim (20% fewer tool calls, 25% fewer tokens) is the only vendor-quantified agent-efficiency delta in Meta's 2026 disclosures. [VENDOR, S23]
- Spark 1.3's 943,718-token maximum output is unusually close to its 1,048,576-token context window (~90%), implying most of the context budget can be converted to output. [SECONDARY, S22]
- Muse Code's `/grill` command name is a distinctive product detail corroborated in launch coverage; it has no analogue in competing agent CLIs covered in this research. [SECONDARY, S14]
- The contributor-pricing data-use term applies to the Muse API pricing tiers generally, not only to Spark — it is a platform-level policy. [SECONDARY, S16]


