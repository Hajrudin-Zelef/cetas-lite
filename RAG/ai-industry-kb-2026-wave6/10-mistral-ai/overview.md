---
id: ai-industry-kb-2026-wave6/10-mistral-ai/overview
title: "§10. Mistral AI"
domain: mistral-ai
role: deep-dive
task: actor-profile
actors: ["AWS", "Alibaba", "China", "EU", "Hugging Face", "Mistral", "Nvidia"]
dates: ["2024-10", "2025-06-30", "2025-07-10", "2025-08", "2025-08-12", "2025-08-13", "2025-12", "2025-12-02", "2025-12-09", "2026-03-16", "2026-04", "2026-04-11", "2026-04-28", "2026-04-29", "2026-05-22", "2026-07-24", "2026-08-31"]
keywords: ["mistral", "agentic", "agents", "apache", "aws", "bedrock", "benchmark", "benchmarks", "context window", "distillation", "distribution", "foundry"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4640, 4718]
section: "§10. Mistral AI"
delta_of: ai-industry-kb-2026
sha256: 50a8d12f103251c027ca29fa9bc76eda61b26fcf411cfbfd8fb963cd0a7b68b6
---

# §10. Mistral AI

Keywords: Mistral, Mistral Large 3, Mistral Medium 3.1, Mistral Medium 3.5, Mistral Small 4, Devstral, Devstral 2507, Devstral 2, Devstral Small 2, Ministral 3, Mistral Vibe CLI, Apache 2.0, modified MIT

## Summary
- **Mistral Large 3** released **December 2025** [SECONDARY] — the brief's "April 11, 2026" date belongs to a different release. The real April 2026 release was **Medium 3.5** (2026-04-29/30) [SECONDARY]; **Medium 3.1** was an August 2025 refresh, not an April 2026 release [SECONDARY].
- **Mistral Small 4** released **2026-03-16**: 119B / 6B active, Apache 2.0, 256K context [SECONDARY].
- **Devstral 2507** (2025-07-10): Devstral Small 1.1 (24B, Apache 2.0, local) and Devstral Medium (proprietary) [VENDOR].
- **Devstral 2** released **2025-12-09/10** (not 2026): Devstral 2 (123B dense, 256K, modified MIT) and Devstral Small 2 (24B, 256K, Apache 2.0), with the Mistral Vibe CLI alongside [SECONDARY]; benchmark claims are [VENDOR].
- **"Ministral 3" (2026)** is a single-source claim kept [UNVERIFIED]; it is not merged with the established Ministral 3B/8B line (Oct 2024).
- Mistral signed the July 24, 2026 "Open Weights and American AI Leadership" letter among the 25 signatories [SECONDARY].

## Key dated facts
### The general line — date corrections
- **2025-12** — Mistral Large 3 released (December 2025, not April 11, 2026) [SECONDARY].
- **2025-08** — Medium 3.1 was a refresh of that month, not an April 2026 release [SECONDARY].
- **2026-03-16** — Mistral Small 4 released: 119B / 6B active, Apache 2.0, 256K context [SECONDARY].
- **2026-04-29/30** — Medium 3.5 released — the real April 2026 release the brief confused with Large 3 [SECONDARY].

### Devstral — the coding line
- **2025-07-10** — Devstral 2507 released [VENDOR].
- Devstral Small 1.1: 24B, Apache 2.0, local deployment [VENDOR].
- Devstral Medium: proprietary [VENDOR].
- **2025-12-09/10** — Devstral 2 released — the brief's 2026 date is corrected to 2025 [SECONDARY].
- Devstral 2: 123B dense, 256K context, modified MIT license [SECONDARY].
- Devstral Small 2: 24B, 256K context, Apache 2.0 [SECONDARY].
- Mistral Vibe CLI shipped alongside Devstral 2 [SECONDARY].
- Performance/benchmark claims for the Devstral 2 generation are [VENDOR] self-reported.

### Ministral 3 (2026) — unverified
- "Ministral 3" in 2026 appears in a single secondary source (aigazine.net) [SECONDARY-only] and is kept **[UNVERIFIED]** in the corpus.
- It is not merged with the established October 2024 Ministral 3B/8B line; treat them as distinct claims [DIRECTIONAL].
- The single source frames the 2026 models around "cascade distillation" — unattested elsewhere in the corpus; do not cite the technique claim beyond the single-source label [SECONDARY-only, UNVERIFIED].

### Coverage context
- The Devstral 2 launch was covered as riding the "vibe coding" tailwinds of late 2025 (TechCrunch, Dec 9, 2025) — the coding-model wave that also carried the Mistral Vibe CLI's release [SECONDARY].
- Devstral 2's 123B dense / modified-MIT pairing is the Mistral pattern in miniature: the most capable coding weights carry commercial strings, while the 24B Small sibling ships Apache 2.0 for local deployment [SECONDARY].
- Mistral's general line and coding line diverge on licensing the same way the industry does: permissive small/local (Small 4, Devstral Small 2) vs gated or proprietary flagship tiers (Devstral 2 modified-MIT, Devstral Medium proprietary, Medium 3.x) [SECONDARY/DIRECTIONAL].
- As a French/EU lab, Mistral's Apache-2.0 small-model line is the EU-domiciled counterpart to the Chinese permissive wave and the Western 2026 Apache firsts — small, local, permissive [DIRECTIONAL].

### Policy posture
- **2026-07-24** — Mistral is among the 25 signatories of the "Open Weights and American AI Leadership" letter [SECONDARY].


### New verified facts — expansion

### Mistral Medium 3.1 refresh (August 2025; details beyond base section)
- Mistral Medium 3.1 was released 2025-08-13 per catalog records, with Mistral's own docs page dating the announcement 2025-08-12. [SECONDARY, S1][VENDOR, S2]
- The model is a refresh of Mistral Medium 3, positioned as a frontier-class multimodal enterprise model with improved tone and response consistency. [SECONDARY, S2][SECONDARY, S1]
- Context window: 131,072 tokens per catalog records (128K per Mistral docs — minor catalog rounding difference). [SECONDARY, S1][SECONDARY, S2]
- Reported output token limit: 104.9K tokens. [SECONDARY, S3 — single source]
- Knowledge cutoff: 2025-06-30. [SECONDARY, S1][SECONDARY, S3]
- API pricing: $0.40 per million input tokens, $2.00 per million output tokens. [SECONDARY, S1][SECONDARY, S3]
- Cache-read pricing: $0.04 per million tokens. [SECONDARY, S1][SECONDARY, S3]
- Input modalities: text, image, and files including PDFs. [SECONDARY, S1][SECONDARY, S2]
- Capabilities: function calling, structured outputs, predicted outputs, document QnA, prefix prompts, batching, agents/conversations, built-in tools. [SECONDARY, S2][SECONDARY, S1]
- Mistral's docs page marks Medium 3.1 (and Medium 3) as deprecated with deprecation 2026-05-22 and retirement 2026-08-31, naming Mistral Medium 3.5 as the replacement; Medium 3.1 retired as scheduled on 2026-08-31. [VENDOR, S2][COMMUNITY, S37][COMMUNITY, S38]
- The model version identifier is v25.08. [VENDOR, S2 — single source]

### Mistral Medium 3.5 (official changelog date 2026-04-28)
- The official platform changelog dates Mistral Medium 3.5 to 2026-04-28, correcting the Apr 29/30 dating in the existing §10. [VENDOR, S4]
- Medium 3.5 is a 128B dense model with 256K-token context. [SECONDARY, S5][SECONDARY, S31]
- It unifies instruction-following, reasoning, and coding with configurable `reasoning_effort`. [SECONDARY, S31][SECONDARY, S5]
- It is multimodal (text and vision input). [SECONDARY, S5][SECONDARY, S31]
- Open weights are released under the Modified MIT license. [SECONDARY, S5][SECONDARY, S31]
- The Modified MIT LICENSE carve-out, quoted directly from the repo's LICENSE file: "You are not authorized to exercise any rights under this license if the global consolidated monthly revenue of your company (or that of your employer) exceeds $20 million." Above that bar, users email Mistral for a commercial license (~$20M/month ≈ $240M/year); everyone below may use, fine-tune, and redistribute commercially. [SECONDARY, S36][SECONDARY, S41]
- Vendor-reported: 77.6% on SWE-Bench Verified and 91.4% on T3-Telecom; vendor claims it outperforms Devstral 2 and Qwen3.5 397B A17B on coding and agentic benchmarks. [VENDOR, S37][SECONDARY, S5]
- Medium 3.5 is a dense 128B model (256K context) released 2026-04-29 that merges chat, reasoning, and code into one set of weights, replacing Medium 3.1, Magistral (Le Chat), and Devstral 2 (Vibe CLI); weights on Hugging Face (including an NVIDIA NVFP4 build) under Modified MIT. [SECONDARY, S31][SECONDARY, S39][SECONDARY, S40][SECONDARY, S41]

### Mistral Large 3 (released 2025-12-02; details beyond base section)
- Mistral Large 3 was released 2025-12-02 as a 675B-total / 41B-active sparse mixture-of-experts model. [SECONDARY, S31][SECONDARY, S7][SECONDARY, S48]
- Context window: 256K tokens. [SECONDARY, S31][SECONDARY, S8]
- The model includes reasoning and vision capabilities. [SECONDARY, S31][SECONDARY, S7]
- License: Apache 2.0 — corroborated across multiple newer sources, contrary to one comparison table's vague "commercial terms" label. [SECONDARY, S31][SECONDARY, S7]
- Reported API pricing: $0.50 per million input tokens, $1.50 per million output tokens. [SECONDARY, S8][SECONDARY, S7]
- Distribution: Mistral API, AWS Bedrock, Azure AI Foundry. [SECONDARY, S7 — single source]
- Hugging Face repository: mistralai/Mistral-Large-3-675B-Instruct-2512. [SECONDARY, S31][SECONDARY, S47][SECONDARY, S49]
- The 2512 suffix in the repository name matches the December 2025 (25/12) release dating. [SECONDARY, S31]

