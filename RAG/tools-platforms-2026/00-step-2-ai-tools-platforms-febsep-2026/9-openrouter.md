---
id: tools-platforms-2026/00-step-2-ai-tools-platforms-febsep-2026/9-openrouter
title: "9. OpenRouter"
domain: step-2-ai-tools-platforms-febsep-2026
role: deep-dive
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Microsoft", "MiniMax", "Nvidia", "OpenAI", "OpenRouter", "Stripe", "xAI"]
dates: ["2025-06", "2026-03", "2026-05-26", "2026-06-30", "2026-07", "2026-07-31", "2026-08", "2026-08-12", "2026-08-19"]
keywords: ["acquisition", "agentic", "attribution", "aws", "benchmarks", "claude", "cost", "deepseek", "disclosure", "embeddings", "funding", "gemini"]
source: docs/RAG/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [399, 460]
section: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
sha256: 4ba9589226a93716e0736c95647fae6aa8e76291464fe849e29636484d26808b
---

# 9. OpenRouter

## 9. OpenRouter

### 9.1 Company overview (2026)
- Founded 2023 by Alex Atallah and Louis Vichy. Team ~90 people as of August 2026. **[independent — SiliconANGLE, Aug 19 2026]**
- Unified, OpenAI-compatible API gateway aggregating 400–500+ models from 80+ inference providers behind a single endpoint. **[official — catalog/API docs; independent]**.
- OpenRouter.ai catalog exposes per-model pricing, supported parameters, context length, top provider data, architecture/modality info, `huggingFaceId`, creation date, knowledge cutoff, and supported voices. **[official — API docs/OpenAPI]**

### 9.2 Funding milestones — 2026
**Series B — May 26, 2026 ($113M, ~$1.3B post-money)**
- $113M Series B led by **CapitalG** (Alphabet's independent growth fund). **[official — openrouter.ai/announcements/series-b; independent — TechCrunch, SiliconANGLE]**
- Participants: NVentures (NVIDIA's venture arm), ServiceNow Ventures, MongoDB Ventures, Snowflake Ventures, Databricks Ventures, plus existing investors Andreessen Horowitz and Menlo Ventures. **[independent — SiliconANGLE]**
- Valuation ~$1.3B post-money, per The New York Times (company did not disclose). Up from ~$547M post-money after the $40M Series A in June 2025 (PitchBook). **[independent — TechCrunch, NYT]**
- Series A (June 2025, $40M) was led by a16z and Menlo Ventures with Sequoia participation. **[independent — TechCrunch]**
- Reported scale at Series B: ~25 trillion tokens/week (≈100T/month), 5× growth in six months; 8M+ global users. **[official/company-reported — Series B announcement; independent]**
- Company quote (CEO Alex Atallah): "Running inference at scale is fundamentally a multimodel problem. The era of picking a single model is over." **[independent — SiliconANGLE]**

**Stripe acquisition — announced August 19, 2026 (pending close)**
- Stripe announced a definitive agreement to acquire OpenRouter. Neither company disclosed terms; closing expected within weeks of the announcement. **[independent — SiliconANGLE, Payments Dive]**
- Reported valuations disagree: NYT ~$7.5B ($1.5B to founders, $6B to investors); Axios >$8B (mostly stock); Bloomberg >$7B; earlier WSJ reporting said talks opened nearer $10B. **[independent — NYT/Axios/Bloomberg/WSJ via SiliconANGLE]**
- Scale reported at acquisition: >10M developers/companies, >10T tokens/day (>200T/month), 400+ models, 80+ providers. **[independent — SiliconANGLE; secondary]**
- Revenue: ~$50M annualized in March 2026 (per ainvest analysis), up from ~$19M at end-2025; one source (cryptobriefing, Aug 2026) put annualized revenue at ~$140M in July 2026 growing 29% MoM. ⚠️ These figures conflict; treat both as secondary estimates, not company disclosures. **[secondary]**
- Customers cited: NVIDIA, Zoom, Lovable. **[independent — SiliconANGLE]**
- Stripe rationale: CEO Patrick Collison called tokens "the central currency" of AI companies; OpenRouter to sit in Stripe's agentic-commerce/payments stack for AI cost optimization and routing. **[independent — Payments Dive, SiliconANGLE]**
- TechCrunch reported Stripe outbid Databricks for the deal. **[independent — TechCrunch via zubnet.ai]** ⚠️ Second-hand attribution; mark as press-reported.

### 9.3 API & product features (changelog, 2026)
Sourced from the official API changelog (`openrouterteam/docs`, changelog.mdx — generated from OpenAPI diffs, human-reviewed for breaking changes) **[official]**:

| Date | Change |
|---|---|
| Aug 19, 2026 | BYOK API: manage credential restrictions (`allowed_models`, `allowed_user_ids`, `allowed_api_key_hashes`) via create/update |
| Jul 29, 2026 | `GET /benchmarks` updated; `UnifiedBenchmarksORItem` schema added |
| Jul 28, 2026 | Breaking (no action needed): `analyst_model` replaces `judge_model` in Fusion streaming events; `POST /responses` OpenAPI tag renamed `Responses`; Web search `blocked_domains`; guardrail creation gains `enable_free_model_publication` / `enable_free_model_training` / `enable_paid_model_training`; auto-router plugins gain `cost_tier` |
| Jul 27, 2026 | `POST /embeddings` request schema updated; `Quantization` and `ReasoningFormat` enums extended |
| Jul 25, 2026 | **Responses API GA** (was `beta.responses`; old SDK namespace kept as deprecated alias); provider options add `claude-on-aws`; BYOK provider enum extended; `voyageai` provider option (Jul 24) |
| Jul 24, 2026 | Anthropic Messages tool blocks widened to support MCP tools (`mcp_tool_reference`, `mcp_toolset_reference`); encrypted compaction content support |
| Jul 8, 2026 | Image streaming chunk events (`POST /images`) |
| Jul 7, 2026 | `GET /models` gains filters: `max_age_days`, `max_agentic_index`, `max_coding_index`, `max_intelligence_index`, `max_output_price`, `max_tool_success_rate` |

Broader API surface (all **[official]** via OpenAPI):
- OpenAI Chat Completions-compatible, Anthropic Messages passthrough, Responses API, presets, Files API, embeddings, rerank, images, audio (speech + transcriptions), video, benchmarks, guardrails, ZDR endpoints.
- Audio: `POST /audio/speech` (19 TTS models verified in July 2026 community test: Fish Audio, MiniMax, Deepgram Aura-2, Gemini 3.1 Flash TTS, Grok, MAI-Voice-2, Qwen, Kokoro, Orpheus, CSM, Zonos, Voxtral) with per-model `supported_voices` arrays. **[secondary — community-verified, not official]**
- Model-routing aliases: `openrouter/free` (auto-routed free models, guaranteed $0) and `openrouter/auto` (best-available, may cost). **[secondary — verified against docs by community]**

### 9.4 Model catalog additions (mid-2026)
Notable new catalog entries observed Aug 2026 **[secondary — community-compiled from live API, prices per 1M tokens]**:

| OpenRouter ID | Catalog created | Context | Input/output price |
|---|---|---|---|
| `anthropic/claude-sonnet-5` | 2026-06-30 | 1M | $2.00 / $10.00 |
| `deepseek/deepseek-v4-flash-0731` | 2026-07-31 | ~1.3M | $0.14 / $0.28 |
| `deepseek/deepseek-v4-pro-0813` | 2026-08-12 | ~1M | $0.66 / $1.98 base (UTC time-band overrides up to $1.32/$3.96 reported) |
| `x-ai/grok-4.6` | 2026-08-12 | 500K | $2.00 / $6.00 ($4.00/$12.00 over 200K prompts) |

Catalog also carries per-model `supported_parameters`, `supported_voices`, quantizations, and tags; endpoints expose provider, quantization, and pricing. **[official]**

### 9.5 Pricing model
- OpenRouter takes ~5–5.5% platform fee on inference spend routed through it (company-reported figures cited in press analyses); credits are non-refundable and do not expire. **[secondary — press analyses; exact current fee terms should be verified on openrouter.ai/docs]**
- "About 5% of the inference spending that runs through the platform stays with the company." **[independent — SiliconANGLE]**
- Pricing is per-model/per-provider with transparent catalog pricing; free `:free` model variants are offered by upstream providers but availability is unstable and controlled by those providers. **[official catalog + secondary]**
- 90%+ gross margin claimed (no own inference infra) — per ainvest analysis. **[secondary]** ⚠️ Analyst estimate, not company disclosure.

