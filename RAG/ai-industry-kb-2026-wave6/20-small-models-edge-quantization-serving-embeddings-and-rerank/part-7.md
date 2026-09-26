---
id: ai-industry-kb-2026-wave6/20-small-models-edge-quantization-serving-embeddings-and-rerank/part-7
title: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers (part 7)"
domain: small-models-edge-quantization-serving-embeddings-and-rerank
role: deep-dive
task: quantization
actors: ["Alibaba", "Apple", "China", "Google", "Meta", "Mistral", "Nvidia"]
dates: ["2025-11-30", "2026-07", "2026-08-31", "2026-09-22"]
keywords: ["embedding", "quantization", "reranker", "agent", "agentic", "agents", "apache", "blackwell", "cost", "fine-tuning", "inference", "latency"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9765, 9794]
section: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers"
delta_of: ai-industry-kb-2026
sha256: 9d64834473b73e9bca92d199f94d0e5be0d39f02d3b4dd132b44fa997c003844
---

# §20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers (part 7)

- Mistral Small 4: 131K context, $0.15/$0.60 per 1M, Apache 2.0 open weights, multimodal with vision built in — the current low-cost open route [SECONDARY]. Sources: https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/ and https://curlscape.com/blog/mistral-api-pricing-2026
- Mistral Large 3: 128K context, open-weight flagship, vision built in; API $0.50/$1.50 per 1M with $0.05 cached input [SECONDARY]. Sources: https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/ and https://www.aipricing.guru/mistral-ai-pricing/
- Mistral Large 3 price history: older secondary pages (secondtalent, smashingapps) list $2/$6 — the current official-synced rate (aipricing.guru 2026-09-22 dataset, curlscape July 2026, devtk Aug 2026) is $0.50/$1.50; the $2/$6 rows are stale, not a second tier [SECONDARY — contradiction resolved by date]. Sources: https://www.aipricing.guru/mistral-ai-pricing/ versus https://www.secondtalent.com/resources/every-mistral-ai-model-explained-compared/
- Mistral Medium 3.5: 131K context, $1.50/$7.50 per 1M, $0.15 cached input — the balanced production tier; Medium 3 is legacy, retired August 31, 2026 [SECONDARY]. Source: https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/
- Mistral Small 3.1 retired November 30, 2025 — use Small 4 [SECONDARY]. Source: https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/
- Codestral: 256K context, $0.30/$0.90 per 1M, specialized for low-latency code completion and generation [SECONDARY]. Sources: https://aiworldtoday.com/guides/mistral-ai-pricing and https://curlscape.com/blog/mistral-api-pricing-2026
- Devstral 2: $0.40/$2.00 per 1M, 256K context, agentic coding; Devstral Small 2: $0.10/$0.30 per 1M, 256K, lightweight coding agents [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026
- Magistral Medium: $2.00/$5.00 per 1M, 128K, deep step-by-step reasoning; Magistral Small: ~$0.10/$0.30 per 1M, 128K, budget reasoning [SECONDARY]. Sources: https://www.secondtalent.com/resources/every-mistral-ai-model-explained-compared/ and https://curlscape.com/blog/mistral-api-pricing-2026
- Ministral 3 edge line: 3B at $0.10/$0.10 (128K), 8B at $0.15/$0.15 (256K), 14B at $0.20/$0.20 (256K) — on-device and cheap volume [SECONDARY]. Sources: https://curlscape.com/blog/mistral-api-pricing-2026 and https://www.aipricing.guru/mistral-ai-pricing/
- Mistral Nemo: Apache 2.0 open lightweight; price conflict — secondtalent lists $0.02/$0.10 while curlscape lists $0.15/$0.15; flag as unresolved, check live page [SECONDARY — contradiction]. Sources: https://www.secondtalent.com/resources/every-mistral-ai-model-explained-compared/ versus https://curlscape.com/blog/mistral-api-pricing-2026
- Mistral-Nemo-Instruct-2407 confirmed Apache 2.0 in a community license audit [COMMUNITY]. Source: https://github.com/ufal/atrium-project/issues/9
- Pixtral Large retired as a standalone model — vision is now built into Large 3, Medium 3.5, and Small 4 [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026
- Mistral specialized services: OCR 4 at $4/1,000 pages (OCR) and $5/1,000 (Document AI); Voxtral TTS $16 per million characters; Voxtral Mini Transcribe 2 $0.003/minute audio; Voxtral Small $0.004/min audio plus $0.10/$0.40 text; Codestral Embed $0.15/M; Mistral Embed $0.10/M; Moderation $0.10/M [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026
- Mistral agent tool charges: web search $30/1,000 calls, code execution $30/1,000 calls, image generation $100/1,000 images, premium news $50/1,000 calls, document libraries $3/1K pages OCR + $1/1M indexing + $0.01/call [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026
- Le Chat plans (2026): Free $0 (~25 messages/day soft cap, Small model); Pro $14.99/month (~150/day, all models, Codestral, Flash Answers, No Telemetry Mode); Team $24/user/month (5-user minimum, SSO, shared workspaces) [SECONDARY]. Source: https://www.smashingapps.com/mistral-ai-review/
- Gemma 4 license confirmation: google/gemma-4-31B-it and google/gemma-4-26B-A4B-it are Apache 2.0 — Google's transition to Apache 2.0 for Gemma 4 eliminates the custom Gemma usage policy [COMMUNITY]. Source: https://github.com/ufal/atrium-project/issues/9
- Qwen license confirmations: Qwen/Qwen3.6-27B Apache 2.0, Qwen/Qwen3.5-9B Apache 2.0, Qwen/Qwen3-235B-A22B-Instruct-2507 Apache 2.0 — unrestricted fine-tuning, integration, and commercial monetization per the audit's reading [COMMUNITY]. Source: https://github.com/ufal/atrium-project/issues/9
- The same audit notes downstream platform runtimes may impose terms differing from the weights license (Mistral-Nemo-Instruct-2407 row) — the weights license and the serving platform's terms are separate documents [COMMUNITY]. Source: https://github.com/ufal/atrium-project/issues/9
- onnx-community Qwen3-0.6B-ONNX and Qwen3-1.7B-ONNX inherit Apache 2.0 from Qwen/Qwen3 (confirmed via README YAML front matter) — quantized/converted redistributions keep the upstream license [COMMUNITY]. Source: https://github.com/mikeschirtzinger/silent-notetaker/blob/HEAD/docs/research/model-licenses.md
- CyberSecEval (Meta): Phi-4 and Gemma 3 27B post the best appropriate-refusal scores (94% and 92%) among small models tested — relevant for SOC deployments where over-permissive models assist attackers [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/small-language-models-2026.pdf
- CoreML-LLM v1.2.0 ships an EmbeddingGemma-300M Core ML conversion at 99.80% ANE utilization, 295MB, with Matryoshka 768/512/256/128 exposed — on-device embedding on the Apple Neural Engine [COMMUNITY]. Source: https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md
- EmbeddingGemma-300M's quantization-aware training holds RAM under 200MB — the sub-500M retrieval tier is now deployable inside mobile/edge memory budgets [SECONDARY]. Source: https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md
- ritrieve_zh_v1 (0.3B): 72.71 STS aggregate with the strongest Chinese-subset scores in the Qwen table (85.98/72.86/63.92 on three columns) — small language-specific embedders beat generalist giants on their home turf [SECONDARY]. Source: https://github.com/QwenLM/Qwen3-Embedding
- BGE-reranker-v2-m3: MTEB-R 57.03, CMTEB-R 72.16, MMTEB-R 58.36 — strongest of the compared rerankers on Chinese retrieval (CMTEB-R) despite trailing on MTEB-R [SECONDARY]. Source: https://github.com/QwenLM/Qwen3-Embedding
- gte-multilingual-reranker-base (0.3B): MTEB-R 59.51, CMTEB-R 74.08, MLDR 66.33 — the 0.3B class remains competitive on Chinese and long-doc retrieval [SECONDARY]. Source: https://github.com/QwenLM/Qwen3-Embedding
- Jina-multilingual-reranker-v2-base (0.3B): MTEB-R 58.22, CMTEB-R 63.37, MMTEB-R 63.73 [SECONDARY]. Source: https://github.com/QwenLM/Qwen3-Embedding
- NVFP4 W4A4 requires Blackwell-native paths, but W4A16 via Marlin support is broader — the same format label covers two different hardware gates [COMMUNITY]. Source: https://github.com/allanschramm/local-model-autotuning/blob/HEAD/docs/discovery/inference-engines-landscape.md
- SmolLM3's shipped config caps context at 65,536 tokens; the 128K figure requires YaRN rope scaling — the config default and the headline number differ [SECONDARY]. Source: https://atomic.chat/models/smollm3-3b
- SmolLM3 has also seen Arabic, Chinese, and Russian data but at lower token counts — multilingual claims beyond the six native languages should be qualified [SECONDARY]. Source: https://atomic.chat/models/smollm3-3b
- The ayinedjimi small-model table's Gemma 3 27B row (MMLU 87.3%, HumanEval 84.2%, MATH 83.6%) is the strongest 2025-era small-model math/code combination in that compilation [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/small-language-models-2026.pdf
