---
id: ai-industry-kb-2026-wave6/21-licenses-and-open-weight-politics/part-6
title: "§21. Licenses and Open-Weight Politics (part 6)"
domain: licenses-and-open-weight-politics
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "Cohere", "DeepSeek", "EU", "Google", "LongCat", "Meta", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "United States", "Z.ai", "xAI"]
dates: []
keywords: ["license", "licenses", "open-weight", "agents", "apache", "astra", "attribution", "awq", "benchmarks", "claude", "cohere", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10192, 10243]
section: "§21. Licenses and Open-Weight Politics"
delta_of: ai-industry-kb-2026
sha256: 619806f3dac8e3b291a9d8ecc3e1825d3226d90c6b321f39943830cb9bc4fd2b
---

# §21. Licenses and Open-Weight Politics (part 6)

- GLM-5.3-Flash's "previously previewed as ox-alpha" stealth testing on OpenRouter/OpenCode (community forensics Aug 21–22: tokenizer match, Z.AI error codes, Java stack trace) preceded the MIT release — open weights can follow closed preview with no license continuity obligation [SECONDARY]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- Zhipu's Hong Kong shares closed up more than 12% (HK$1,160) on the GLM-5.3-Flash/MIT launch news — markets price open-weight releases [SECONDARY]. Source: https://pasqualepillitteri.it/en/news/13263/ox-alpha-revealed-zhipu-glm-5-3-flash-mit


Model | License | Source basis
GPT-6 Astra, GPT-5.6 Sol/Terra/Luna, GPT-5.5, GPT-5.4, GPT-5.3-Codex | Proprietary API-only, no weights | https://devtk.ai/en/blog/openai-api-pricing-guide-2026/
Claude Opus 5, Sonnet 5, Fable 5, Haiku 4.5, Mythos 5/5.1 | Proprietary API-only, no weights | https://www.morphllm.com/claude-code-pricing
Gemini 3.8/3.7/3.6/3.5 Flash, 3.1 Pro, 2.5 Flash, 3.5 Flash-Lite | Proprietary API-only, no weights | https://venturebeat.com/technology/googles-gemini-3-7-flash-targets-coding-and-agents-with-a-50-introductory-price-cut
Grok 4.5, Grok 4.6, Grok Code Fast 1 | Proprietary API-only, no weights | https://www.aipricing.guru/xai-grok-pricing/
Qwen3 0.6B–32B dense, 30B-A3B, Qwen3-235B-A22B | Apache 2.0 | https://registry.ollama.com/library/qwen3:0.6b-q8_0/blobs/d18a5cc71b84
Qwen3-Coder-30B-A3B-Instruct (+1M variant) | Apache 2.0 | https://huggingface.co/unsloth/Qwen3-Coder-30B-A3B-Instruct-1M-GGUF
Qwen3.5-9B, Qwen3.6-27B, Qwen3.7 Max/Plus/Turbo | Apache 2.0 | https://github.com/ufal/atrium-project/issues/9, https://github.com/kzinmr/ai-topics/blob/HEAD/wiki/comparisons/llm-api-pricing.md
Qwen3.8-Max | Custom text-only weights (not Apache) | wave6/03-model-weights-wave3.md
Qwen3-Embedding 0.6B/4B/8B, Qwen3-Reranker 0.6B/4B/8B, Qwen3-VL-Embedding | Apache 2.0 | https://github.com/QwenLM/Qwen3-Embedding
Qwen3.5-Omni Light | UNVERIFIED — no LICENSE confirmation | (gap)
DeepSeek V4 Pro, V4 Flash, V4.1 Flash | MIT | https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
GLM-5.3-Flash | MIT | https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
GLM-5.3 (non-Flash) | Closed at launch; staged behind safety review, license unpublished | https://emergent.sh/learn/what-is-glm-5-3
GLM-5.2 | MIT | https://www.orcarouter.ai/blog/deepseek-v4-1-flash-vs-kimi-k3
Kimi K3 | Modified MIT (100M MAU attribution; $20M/mo revenue prong unresolved) | https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/ vs https://www.orcarouter.ai/blog/deepseek-v4-1-flash-vs-kimi-k3
MiniMax M3 | Community License, geo-restricted local deployment | https://github.com/utensils/mold/blob/HEAD/docs/qualification/minimax-h3.md
MiniMax H3 | H3 Community License (EU/UK/SK/US excluded, $20M revenue, Exhibit A) | https://huggingface.co/OpenVDN/vdn-minimax-h3/blob/main/README.md
MiniMax Music 3 | Community license, NO territory exclusion, $20M revenue trigger | https://digitalmatters.me/artificial-intelligence-ai/minimax-music-3/
LongCat-2.0 | MIT | https://github.com/XUEXUE-XING/minimax-h3-webui/pull/12
MiMo V2.5 | MIT | wave6/03-model-weights-wave3.md
Llama 3.2 11B | Llama 3 Community License | https://ayinedjimi-consultants.fr/static/pdf/small-language-models-2026.pdf
Meta Muse Glimmer | Apache 2.0 | https://digitalmatters.me/artificial-intelligence-ai/minimax-music-3/
Phi-4 family (14B, mini, multimodal, reasoning, reasoning-vision) | MIT | https://techcommunity.microsoft.com/blog/educatordeveloperblog/phi-4-small-language-models-that-pack-a-punch/4464167
SmolLM3-3B | Apache 2.0 | https://github.com/huggingface/blog/blob/HEAD/smollm3.md
Gemma 3 (1B/4B/12B/27B/270M) | Gemma Terms of Use (custom) | https://themunicheye.com/google-launches-gemma-3-open-source-ai-model-13085
Gemma 4 (31B-it, 26B-A4B-it) | Apache 2.0 | https://github.com/ufal/atrium-project/issues/9
NVIDIA Nemotron 3/3.5 | NVIDIA Open Model License | wave6/02-model-weights-delta.md
Mistral Small 4 | Apache 2.0 | https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/
Mistral Nemo | Apache 2.0 | https://github.com/ufal/atrium-project/issues/9
Mistral Large 3 | UNVERIFIED — open-weight status confirmed, license text not pulled in this pass | https://aiworldtoday.com/guides/mistral-ai-pricing vs https://www.secondtalent.com/resources/every-mistral-ai-model-explained-compared/
Ministral 3 (3B/8B/14B) | Open (edge) — exact license text not pulled | https://curlscape.com/blog/mistral-api-pricing-2026
Mistral Medium 3.5, Codestral, Magistral | Proprietary API-only | https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/
Command A+ (218B/25B) | Apache 2.0 | wave6/02-model-weights-delta.md
Cohere Embed v4 | Proprietary API-only | https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md
Jina v3 | CC BY-NC 4.0 | https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md
Jina v4 | Qwen Research (non-commercial) | https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md
Jina v5-omni-small | CC BY-NC-4.0 | https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md
KaLM-Embedding-Gemma3-12B | Custom (non-standard) | https://www.codesota.com/benchmarks/mteb
BGE-large-en-v1.5 | Apache 2.0 | https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
nomic-embed-text | Apache 2.0 | https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
EmbeddingGemma-300M | UNVERIFIED — license not confirmed in this pass | (gap)
OpenAI text-embedding-3 | Proprietary API-only | https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
Qwen2.5-14B-Instruct-AWQ | Apache 2.0 (redistribution) | https://github.com/ufal/atrium-project/issues/9
FunAudioLLM/SenseVoiceSmall | FunASR Model Open Source License v1.1 — commercial scope ambiguous, do not treat as MIT | https://github.com/mikeschirtzinger/silent-notetaker/blob/HEAD/docs/research/model-licenses.md
Heimdall qwen3-0.6b purpose build | Apache 2.0 (inherited) | https://github.com/mikeschirtzinger/silent-notetaker/blob/HEAD/docs/research/model-licenses.md


