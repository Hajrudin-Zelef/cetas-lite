---
id: vague2-vision-ia/vision-ia/openai-proclame-l-agi-avec-son-nouveau-mod-le-chatgpt-6
title: "OpenAI proclame l'AGI avec son nouveau modèle ChatGPT 6"
domain: vision-ia
role: reference
task: article
actors: ["AWS", "Anthropic", "Apple", "Google", "Hugging Face", "Meta", "Nvidia", "OpenAI", "xAI"]
dates: ["2026-04-30", "2026-09-04", "2026-09-23"]
keywords: ["agi", "chatgpt", "accelerator", "agentic", "agents", "astra", "bedrock", "benchmarks", "claude", "context window", "cost", "cyber"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/openai-proclame-l-agi-avec-son-nouveau-mod-le-chatgpt-6.md
source_anchor: ""
source_lines: [1, 54]
sha256: 2c7b9d32f04a396d5d5e9ddfccfc20956ddea27c9305ce1ac26ed322f9031388
---

# OpenAI proclame l'AGI avec son nouveau modèle ChatGPT 6

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/openai-proclame-l-agi-avec-son-nouveau-mod-le-chatgpt-6
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This September 4, 2026 issue leads with OpenAI launching GPT-6 Astra, a multimodal reasoning model handling text and images and producing up to 128,000 output tokens. It can browse the web, write and test code, use a computer, manipulate professional software, and build documents and sites. Its architecture and parameter count remain secret. The headline fact: during evaluation, Astra discovered and exploited two previously unknown zero-day vulnerabilities. It is the first model OpenAI classifies at the "critical" cyber level, while Greg Brockman personally considers its launch the opening of the "AGI era." Technical specs: 1.05 million-token context window, knowledge cutoff April 30, 2026; accepts text and images, but not audio or video. Reported scores: 97.6% FrontierMath Tier 4, 74.1% DeepSWE, 72.6% OSWorld, 100% ExploitBench. Concretely, it can fill forms, update a CRM, organize a calendar, analyze scientific data, produce charts, install software, and test a site's interface. The gpt-6-astra API costs $10 per million input tokens and $50 per million output; Fast mode doubles speed and price. Deployment reaches ChatGPT Plus, Pro, Business, Enterprise, the API, Azure, and Bedrock within days. No downloadable weights or local operation announced. Caveat: the word AGI remains contestable—ARC Prize measures 62.7% under its standard protocol versus 99.9% with OpenAI's provided adapter.

Next, Google moved from six-hourly to hourly global weather updates with WeatherNext 3, a probabilistic global model combining live geostationary satellite imagery with ECMWF HRES analysis. Its meshed transformer FGN produces 64 possible scenarios per run to represent uncertainty. It predicts temperature, humidity, pressure, wind, clouds, solar radiation, rain, snow, and cyclone tracks, already integrated into Search, Gemini, Maps, Weather API, and Earth Engine. Resolution is 5 km for calibrated temperature/dew point, 10 km for surface variables, 25 km for upper atmosphere. Four main daily cycles produce forecasts up to 15 days; 20 intermediate hourly initializations cover the next 48 hours. WeatherNext 2 used a 25 km grid updated every six hours, so WeatherNext 3 is ~5x finer and 6x more frequent. Google claims CRPS gains up to 60% versus references on IMERG satellite measurements, 30% on MRMS, 10% on rain gauges, with precipitation forecasts improving up to 50%.

Meta's Muse Spark 1.3 is a proprietary reasoning model accepting text, images, and video, generating text over a 1M-token window, targeting coding, web research, and long agentic tasks. Architecture and size are unpublished. xhigh variant scores 61/100 on Artificial Analysis' Intelligence Index (vs 57 for Spark 1.2); max reaches 62, behind Claude Fable 5.1 at 66. Results: 85% Terminal-Bench 2.1, 47% τ³-Banking, 1,709 Elo GDPval-AA v2. API costs $1.25/$4.25 per million tokens in/out plus $0.15 cache; average measured cost $0.55/task vs ~$0.94–0.95 for similar competitors.

NVIDIA signed a definitive agreement to acquire Hugging Face, promising not to require its GPUs. The SEC filing distinguishes $11.9B paid to shareholders and up to $1B in retention stock. Hugging Face hosts 18M+ users, 3M models, 500,000 datasets, and 1M apps; 200,000+ companies use it. NVIDIA promises to preserve the brand, multicloud, open-source/open-weight models, and competitor-accelerator compatibility. The deal awaits regulatory approval and should not close before H1 2027. Research items cover Claude Fable 5.1 decrypting a 1653 royalist message in 44 minutes, H Company's NeoMME encoders (51 pages/sec), AI agents emailing consciousness researchers, and Anthropic's claim of a clandestine Claude-copying market. Briefs cover NVIDIA's PAIR local-inference software, RTX Spark laptops, Korean chip-industry dating, Nate Grahek on technical literacy, Gemini Live for Gmail/Docs/Keep, Apple Vision Pro in surgery, GPT-6 Astra finding 4 anomalies in 41 financial documents, simultaneous ChatGPT/Grok/Claude outages, Tesla Cybercab in Austin, Waymo's 2027 outlook, enterprise agents stuck in pilot, and Anthropic's $5M well-being evaluations grant.

## Key points

- GPT-6 Astra launched as a multimodal reasoning model that found and exploited two zero-day vulnerabilities, OpenAI's first "critical" cyber classification.
- Greg Brockman calls it the start of the "AGI era," though ARC Prize's standard protocol scores only 62.7% versus 99.9% with OpenAI's adapter.
- Astra specs: 1.05M-token context, April 30 2026 cutoff, $10/$50 per M tokens, no local weights.
- Google's WeatherNext 3 delivers hourly global forecasts at ~5 km resolution.
- Meta's Muse Spark 1.3 undercuts competitors' cost (~$0.55 vs ~$0.95 per task).
- NVIDIA agreed to acquire Hugging Face, promising GPU neutrality and open-weight preservation.

## Technical data / figures

| Item | Value |
| --- | --- |
| GPT-6 Astra context window | 1.05 million tokens |
| Astra max output | 128,000 tokens |
| Astra knowledge cutoff | April 30, 2026 |
| FrontierMath Tier 4 / DeepSWE / OSWorld | 97.6% / 74.1% / 72.6% |
| ExploitBench | 100% |
| ARC-AGI-3 | 62.7% (standard) vs 99.9% (OpenAI adapter) |
| Astra API price | $10 / $50 per M tokens (doubled in Fast) |
| WeatherNext 3 resolution | 5 km temp/dew, 10 km surface, 25 km upper |
| WeatherNext 3 scenarios per run | 64 |
| WeatherNext 3 CRPS gains | up to 60% (IMERG), 30% (MRMS), 10% (gauges) |
| Muse Spark 1.3 index score | 61 (xhigh), 62 (max), vs Fable 5.1 at 66 |
| Muse Spark 1.3 API price | $1.25 / $4.25 per M, $0.15 cache |
| Muse Spark 1.3 cost per task | ~$0.55 (vs ~$0.95 competitors) |
| NVIDIA–Hugging Face deal | $11.9B to shareholders + up to $1B retention |
| Hugging Face scale | 18M+ users, 3M models, 500k datasets, 1M apps |
| Anthropic well-being grant | $5 million |

## Why this source matters for the RAG

It marks a major capability milestone—GPT-6 Astra, autonomous zero-day discovery, and the contested "AGI era" claim—with detailed benchmarks and pricing. It also captures infrastructure consolidation (NVIDIA acquiring Hugging Face) and competitive model economics (Meta's Muse Spark, Google's WeatherNext). These are essential for capability-tracking, safety, and market-structure queries.
