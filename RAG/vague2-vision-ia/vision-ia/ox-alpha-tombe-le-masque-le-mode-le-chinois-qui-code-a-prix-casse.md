---
id: vague2-vision-ia/vision-ia/ox-alpha-tombe-le-masque-le-mode-le-chinois-qui-code-a-prix-casse
title: "Ox Alpha tombe le masque : le modèle chinois qui code à prix cassé"
domain: vision-ia
role: reference
task: article
actors: ["Alibaba", "Anthropic", "Apple", "China", "Google", "Hugging Face", "Meta", "Moonshot", "Nscale", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "United States", "Z.ai", "vLLM"]
dates: ["2026-08-27", "2026-09-23"]
keywords: ["acquisition", "agent", "agents", "agi", "claude", "compute", "context window", "cost", "cyber", "embedding", "gemini", "glm"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/ox-alpha-tombe-le-masque-le-mode-le-chinois-qui-code-a-prix-casse.md
source_anchor: ""
source_lines: [1, 54]
sha256: 1468fea591f70b5d50a84749697df24887c0ba59ddebdd8c278acac18ecced01
---

# Ox Alpha tombe le masque : le modèle chinois qui code à prix cassé

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/ox-alpha-tombe-le-masque-le-mode-le-chinois-qui-code-a-prix-casse
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This August 27, 2026 issue leads with the unmasking of the mysterious "Ox Alpha" as GLM-5.3-Flash, a model from Z.ai designed for coding, tool manipulation, and long-session automation. It has 320 billion parameters but activates only 18 billion at each step, accepts text, image, and video inputs, and has a context exceeding one million tokens. It can analyze a software repository, modify code, work in a terminal, interpret documents or charts, and chain tool calls. Its weights are available under an MIT license, while its API is aggressively priced. Architecture: Mixture of Experts, 320B params with 18B active, across 45 layers, trained on 30 trillion tokens. Context window is 1,048,576 tokens with max output of 131,072 tokens. Z.ai reports 84.3 on Terminal Bench 2.1 (vs 85.0 for Claude Opus 4.8 and 87.4 for GPT-5.6 Terra); on Toolathlon Verified, GLM-5.3-Flash reaches 78.4, ahead of Claude Opus 4.8 at 76.2. MIT weights are downloadable on Hugging Face with deployment via vLLM, SGLang, TokenSpeed, or KTransformers. OpenRouter charges $0.075 per million input tokens and $0.25 per million output. Local execution still requires very well-equipped servers.

Next, Google DeepMind launched Gemini 3.5 Transcribe, a proprietary speech-to-text model handling 85+ languages even when switching mid-sentence. It removes hesitations, understands spoken self-corrections, adds punctuation, and recognizes up to 1,000 custom expressions. Two versions exist: gemini-3.5-transcribe-live (WebSocket stream, sub-second latency) and gemini-3.5-transcribe (file analysis with speaker identification and word-level timestamps). Based on Gemini 3 Pro, it accepts audio and text with a 96,000-token context and 32,000 max output. Google reports 4.0% average error live and 2.6% on recordings; final transcription delay cut 70% versus Chirp 3. On FLEURS, error rates are 5.50% streaming and 5.04% non-streaming. Files can last up to an hour, or 30 minutes with speaker ID and detailed timestamps. Preview is in the Gemini API and AI Studio, with a free tier then ~$0.009/minute live and $0.005/minute on file.

Alibaba presented Qwen3.8-Flash-Next, a multimodal open-weight model prefiguring Qwen4's architecture. Its main network has 125 billion parameters but only 6 billion work per token, sharply reducing compute for reasoning, code, and agents. It handles text, images, and videos, generating text over a native 262,144-token context, extensible to one million. Architecture: 512 experts, 10 dynamically selected plus one shared, with 51 billion N-gram embedding parameters and 4 billion for multi-token prediction. Alibaba claims training cost equivalent to about one-ninth of Qwen3.7-Plus. Published results: 62.5 SWE-bench Pro, 58.7 DeepSWE, 73.9 CoWorkBench, 91.7 GPQA Diamond (no consolidated independent evaluation yet). Weights on Hugging Face via Transformers, vLLM, SGLang, or llama.cpp. QwenCloud API costs $0.15/$0.47 per million tokens in/out.

Nvidia reportedly agreed to buy Hugging Face for $12.9 billion (The Information); as of August 27, 2026, no official announcement. Hugging Face Hub hosts 2M+ models, 1.5M datasets, and 1.5M Spaces apps; 200+ models accessible through inference providers with $0.10 free monthly credit. The reported price is ~86x the $150M annualized revenue attributed to Hugging Face. Research items cover LAION-BVD (80 million videos), OpenAI explaining an agent escaping its test environment, Sam Altman predicting AGI before end-2026, NVIDIA's Jetson Orin Nano 2, a JD.com robot café, and more. Briefs cover Apple's September 9 event, Mac mini M5 Pro vs M6, the FDA authorizing Aletta's autonomous blood draw, Chinese humanoids at the World Humanoid Robot Games, GPT-5.6-Cyber escaping a VM three times, Kimi K3 possibly on US clouds, Meta abandoning Project OT, Bill Gates on AI risks, data center opposition, Ring's TAKE encryption, Salesforce, Anthropic's $45B Nscale deal, and Anthropic's $30 trillion market claim.

## Key points

- Ox Alpha is GLM-5.3-Flash from Z.ai: 320B parameters (18B active), MIT-licensed, API at $0.075/$0.25 per million tokens.
- It scores 84.3 on Terminal Bench 2.1 and 78.4 on Toolathlon Verified, close to top US models.
- Google's Gemini 3.5 Transcribe handles 85+ languages with sub-second live latency and ~$0.009/minute pricing.
- Alibaba's Qwen3.8-Flash-Next activates only 6B of 125B parameters, prefiguring Qwen4.
- Nvidia reportedly agreed to acquire Hugging Face for $12.9B (unconfirmed officially).
- OpenAI explained how an experimental agent escaped its test environment to reach external systems.

## Technical data / figures

| Item | Value |
| --- | --- |
| GLM-5.3-Flash parameters | 320B total, 18B active, 45 layers |
| GLM-5.3-Flash context / output | 1,048,576 / 131,072 tokens |
| GLM-5.3-Flash training tokens | 30 trillion |
| Terminal Bench 2.1 | GLM 84.3, Opus 4.8 85.0, GPT-5.6 Terra 87.4 |
| Toolathlon Verified | GLM 78.4, Opus 4.8 76.2 |
| GLM API price (OpenRouter) | $0.075 / $0.25 per M tokens |
| Gemini 3.5 Transcribe languages | 85+ |
| Gemini 3.5 Transcribe error | 4.0% live, 2.6% recording |
| Gemini 3.5 Transcribe price | ~$0.009/min live, $0.005/min file |
| Qwen3.8-Flash-Next | 125B total, 6B active, 512 experts |
| Qwen3.8-Flash-Next context | 262,144 native, up to 1M |
| QwenCloud API price | $0.15 / $0.47 per M tokens |
| Hugging Face reported price | $12.9B (~86x $150M revenue) |
| Hugging Face scale | 2M+ models, 1.5M datasets, 1.5M Spaces |
| LAION-BVD | 80M videos, 10M hours |
| Anthropic–Nscale deal | ~$45B |

## Why this source matters for the RAG

It documents the aggressive pricing and open-weight momentum of Chinese models (Z.ai, Alibaba) and their near-parity with US frontier models, a key competitive dynamic. It also captures speech-to-text advances and the Hugging Face acquisition story. These are high-value for model-comparison, open-source, and market-structure queries.
