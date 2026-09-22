---
id: ai-industry-kb-2026-wave6/10-mistral-ai/mistral-small-4-released-2026-03-16-details-beyond-base-sect
title: "Mistral Small 4 (released 2026-03-16; details beyond base section)"
domain: mistral-ai
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "China", "DeepSeek", "Google", "Hugging Face", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Unsloth", "Z.ai", "vLLM"]
dates: ["2025-03-30", "2025-06-10", "2025-07-10", "2025-08", "2025-09", "2025-12", "2025-12-02", "2025-12-09", "2026-03-16", "2026-06-23", "2026-07-07", "2026-07-10", "2026-07-31", "2026-09-21"]
keywords: ["mistral", "agentic", "alignment", "apache", "benchmark", "benchmarks", "blackwell", "claude", "context window", "cost", "deepseek", "distribution"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4719, 4795]
section: "§10. Mistral AI"
delta_of: ai-industry-kb-2026
sha256: 7dcad33284408fc522a3e88f6c186fd97a3698b9b44391a32974fbddfbee08f6
---

# Mistral Small 4 (released 2026-03-16; details beyond base section)

### Mistral Small 4 (released 2026-03-16; details beyond base section)
- Mistral Small 4 is a 119B-total mixture-of-experts model with approximately 6–6.5B active parameters per token. [SECONDARY, S9][SECONDARY, S10][SECONDARY, S31]
- Expert configuration: 128 experts per layer (~half of comparable models like DeepSeek-R1 per NVIDIA), 675B total parameters, ~41B active per token (16:1 total-to-active ratio); 256K context; multilingual 40+ languages; trained from scratch on ~3,000 NVIDIA H200 GPUs. [SECONDARY, S9][SECONDARY, S42][SECONDARY, S43][SECONDARY, S44]
- Advertised context: 256K tokens in stronger sources; some catalog mirrors report 128K (documented contradiction). [SECONDARY, S31][SECONDARY, S9]
- Input modalities: text and image; unified instruct/reasoning/multimodal/agentic-coding model. [SECONDARY, S9][SECONDARY, S10]
- Configurable reasoning via `reasoning_effort="none"` or `"high"`. [SECONDARY, S9 — single source]
- Vendor performance claims: 40% lower completion latency and 3× throughput versus Mistral Small 3. [VENDOR, S9 — single source]
- Independent coverage: Large 3 scores below DeepSeek-v3.2, Kimi K2-Thinking, and GLM-4.6 but above OLMo 3 and Llama 4 Maverick on the AA Intelligence Index; open-weight, Apache 2.0, vLLM/Ollama/llama.cpp support with NVFP4 for Blackwell. [SECONDARY, S44][SECONDARY, S45]
- Distribution: Mistral API, Mistral AI Studio, Hugging Face, NVIDIA NIM. [SECONDARY, S10][SECONDARY, S9]
- Secondary hardware guidance for practical self-hosting: approximately four H100 GPUs, two H200 GPUs, or one DGX B200. [SECONDARY, S12 — single source]
- Hugging Face repository: mistralai/Mistral-Small-4-119B-2603. [SECONDARY, S31 — single source]

### Ministral 3 family (released 2025-12-02 — correction to existing §10)
- The existing §10's "Ministral 3 (2026)" single-source/UNVERIFIED entry is corrected: the Ministral 3 family shipped 2025-12-02. [SECONDARY, S31][SECONDARY, S13]
- Three sizes: 3B, 8B, and 14B dense edge models, all with vision capabilities. [SECONDARY, S31][SECONDARY, S13]
- License: Apache 2.0 across the family. [SECONDARY, S31][SECONDARY, S13]
- Context: 131K–262K tokens depending on size and source. [SECONDARY, S13 — single source]
- The 14B reasoning variant hits 85% on AIME 2025; Ministral 3 is dense 14B/8B/3B, Apache 2.0. [SECONDARY, S13][COMMUNITY, S45]
- Positioning: low-latency, cost-sensitive, and edge applications including mobile and browser deployment. [SECONDARY, S31][SECONDARY, S13]
- Hugging Face repositories: mistralai/Ministral-3-3B-Instruct-2512, Ministral-3-8B-Instruct-2512, Ministral-3-14B-Instruct-2512. [SECONDARY, S31 — single source]

### Devstral 2507 (released 2025-07-10; details beyond base section)
- Devstral 2507 shipped 2025-07-10 in two variants: Devstral Small 1.1 (open-source) and Devstral Medium (API-based). [SECONDARY, S24][SECONDARY, S21]
- Devstral Small 1.1: 24B parameters, Apache 2.0 license. [SECONDARY, S21][SECONDARY, S25]
- Devstral Small 1.1 was fine-tuned from the Mistral-Small-3.1 architecture. [SECONDARY, S22][SECONDARY, S23]
- Devstral Small 1.1 is text-only — it omits the vision encoder present in Small 3.1. [SECONDARY, S23 — single source]
- Context window: 128K tokens. [SECONDARY, S22][SECONDARY, S23]
- SWE-Bench Verified: Small 1.1 scored 53.6% (OpenHands scaffold), up from 46.8% for Small 1.0. [SECONDARY, S21][SECONDARY, S24]
- The 53.6% score made Small 1.1 the #1 open model on SWE-Bench Verified as of 2026-07-10; 24B params, finetuned from Small 3.1 via RL + safety alignment, Apache 2.0, 128K context, Tekken tokenizer; API $0.10/$0.30 per M. [SECONDARY, S24][SECONDARY, S50][SECONDARY, S53][SECONDARY, S54]
- Comparison points on the same table: GPT-4.1-mini 23.6%, Claude 3.5 Haiku 40.6%, SWE-smith-LM 32B 40.2%, Skywork SWE 38.0%, DeepSWE 42.2% (R2E-Gym). [SECONDARY, S21 — single source]
- Developed in collaboration with All Hands AI (OpenHands scaffold). [SECONDARY, S22][SECONDARY, S25]
- Runs on a single NVIDIA RTX 4090 or a Mac with 32GB RAM. [SECONDARY, S22][SECONDARY, S23]
- Distribution: Hugging Face, Ollama, Kaggle, Unsloth, LM Studio. [SECONDARY, S22][SECONDARY, S23]
- API listing `devstral-small-2505`: $0.10 per million input tokens, $0.30 per million output tokens. [SECONDARY, S23 — single source]
- Devstral Medium: API-based, enterprise-focused (private infra deployment, custom fine-tuning), scored 61.6% on SWE-Bench Verified — outperforming Gemini 2.5 Pro and GPT-4.1 at 25% of the cost per one secondary. [SECONDARY, S24][SECONDARY, S50][SECONDARY, S51][SECONDARY, S52]
- Devstral 2 (123B, version 2512) launched 2025-12-09 as the next generation of the coding-model line — 123B parameters, 256K-token context, Modified MIT license with the same $20M/month carve-out as Medium 3.5 (covering base model, derivatives, fine-tunes, and redistributed variants regardless of host). [SECONDARY, S32][SECONDARY, S44][SECONDARY, S43]
- Devstral 2 2512 appears in current OpenRouter model listings as a Mistral text model. [SECONDARY, S1 — single source]

### Magistral reasoning family (launched 2025-06-10)
- Mistral announced Magistral on 2025-06-10 at London Tech Week — its first family of reasoning models with step-by-step chain-of-thought. [SECONDARY, S26][SECONDARY, S30][SECONDARY, S51]
- Two variants: Magistral Small (open-source) and Magistral Medium (enterprise). [SECONDARY, S26][SECONDARY, S30]
- Magistral Small: 24B parameters, Apache 2.0, downloadable from Hugging Face. [SECONDARY, S26][SECONDARY, S30][SECONDARY, S51]
- Magistral Medium: enterprise-grade; preview in Le Chat and on La Plateforme; also on Amazon SageMaker, with IBM WatsonX, Azure AI, and Google Cloud Marketplace planned. [SECONDARY, S26][SECONDARY, S30][SECONDARY, S51]
- Magistral Medium vendor benchmarks: 73.6% on AIME 2024, rising to 90% with majority voting at 64 samples. [VENDOR, S26][SECONDARY, S27][SECONDARY, S51]
- Magistral Small vendor benchmarks: 70.7%, rising to 83.3% (with majority voting, per reporting). [VENDOR, S27][SECONDARY, S51]
- Multilingual chain-of-thought reasoning in English, French, Spanish, German, Italian, Arabic, Russian, and Simplified Chinese. [SECONDARY, S26][SECONDARY, S27]
- "Think mode" and "Flash Answers" features in Le Chat; Flash Answers claimed 10× faster token throughput than most competitors. [VENDOR, S26 — single source]
- French President Emmanuel Macron publicly backed the launch, making it the first European reasoning model with head-of-state political support. [SECONDARY, S27 — single source]
- Magistral Small 1.2 (September 2025, HF Magistral-Small-2509): open 24B multimodal reasoning model, 128K context, `[THINK]` tokens; now legacy with published API retirement 2026-07-31. [SECONDARY, S31 — single source]
- Per current catalog guidance, Magistral-style reasoning now lives in Small 4 and Medium 3.5 rather than the standalone Magistral line (reasoning via a `reasoning_effort` "high"/"none" toggle; HF config notes a long-context entry fix in commit c4be198050fb). [SECONDARY, S31][SECONDARY, S39][COMMUNITY, S46]
- Catalog pricing: Magistral Medium $2/M input, $8/M output; Magistral Small $0.50/M input, $1.50/M output; both 128K context. [SECONDARY, S19 — single source]
- TechCrunch's launch assessment noted Magistral Medium underperformed Gemini 2.5 Pro and Claude Opus 4 on GPQA Diamond/AIME and trailed Gemini 2.5 Pro on LiveCodeBench. [SECONDARY, S30][SECONDARY, S51]

### Mistral OCR generations (new to this section)
- Mistral OCR 3 launched in December 2025, unveiled as a Tuesday release during Mistral's December product offensive. [SECONDARY, S32][SECONDARY, S33]
- OCR 3 pricing: $2 per 1,000 pages, with a 50% discount for batch processing. [SECONDARY, S32][SECONDARY, S35]
- OCR 3 vendor claim: 74% win rate against competing products on forms, scanned documents, complex tables, and handwritten content. [VENDOR, S32 — single source]
- Mistral OCR 4 launched 2026-06-23. [SECONDARY, S35][SECONDARY, S34]
- OCR 4 pricing: $4 per 1,000 pages, dropping to $2 per 1,000 pages with the Batch API. [SECONDARY, S35][SECONDARY, S34]
- OCR 4 returns bounding boxes, typed-block classification (titles, tables, equations, signatures), and per-page and per-word confidence scores alongside extracted text. [SECONDARY, S35][SECONDARY, S34]
- OCR 4 covers 170 languages across 10 language groups. [SECONDARY, S35][SECONDARY, S34]
- OCR 4 ships as a single container for fully air-gapped on-premise deployment. [SECONDARY, S35][SECONDARY, S34]
- OCR 4 vendor claims: 72% average win rate in blind human evaluations against leading systems; top score on OlmOCRBench (85.20). [VENDOR, S35 — single source]
- Mistral itself notes the OCR benchmark figures are directional rather than definitive. [VENDOR, S35 — single source]
- OCR 4 distribution: Mistral API and Studio, Amazon SageMaker, Microsoft Foundry. [SECONDARY, S34 — single source]
- Mistral OCR 4.1 exists as a proprietary document-parsing model with bounding boxes, block labels, block-level confidence, 170 languages, at $4/1K pages. [SECONDARY, S31 — single source]
- Mistral held an OCR 4 production webinar on 2026-07-07 at 18:00 CET. [SECONDARY, S33 — single source]
- Anaqua, a legal software firm, measured roughly 4× faster throughput per page with OCR 4 versus its prior provider. [SECONDARY, S34 — single source]

### Codestral and model catalog pricing (new to this section)
- Codestral catalog pricing: $0.30 per million input tokens, $0.90 per million output tokens, 256K context. [SECONDARY, S19 — single source]
- A second provider listing shows Codestral-latest at $0.33/$0.99 with 131K context, updated 2026-09-21, added August 2025, 80+ languages supported. [SECONDARY, S18 — single source]
- Devstral Medium catalog pricing: approximately $0.50/$1.50. [SECONDARY, S19 — single source]
- Shieldstral 1.0 exists in Mistral's model catalog as a content-moderation model entry. [VENDOR, S2 — single source]
- Retired legacy entries: Magistral Small 1.2 and Mistral Small 3.2 (24B, v3.2, repo 2506) both carry published API retirement 2026-07-31. [SECONDARY, S31 — single source]
- Mixtral 8x22B (141B total / 39B active) carries published API retirement 2025-03-30. [SECONDARY, S31 — single source]

