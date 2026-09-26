---
id: ai-industry-kb-2026-wave6/20-small-models-edge-quantization-serving-embeddings-and-rerank/part-12
title: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers (part 12)"
domain: small-models-edge-quantization-serving-embeddings-and-rerank
role: deep-dive
task: quantization
actors: ["DeepSeek", "Mistral", "United States", "vLLM"]
dates: ["2023-12", "2024-09-25", "2026-09-08"]
keywords: ["embeddings", "agentic", "apache", "attention", "bedrock", "compute", "deepseek", "distribution", "inference", "intel", "license", "llama"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10028, 10058]
section: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers"
delta_of: ai-industry-kb-2026
sha256: 76e87ccdabab1ac20e32c0e55dcfa0d177384a167db6f67bfa7b9e4b7b8904dc
---

# §20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers (part 12)

- Llama 3.2 1B/3B verified (gap closed): 1B is 1.23B actual parameters, 3B is 3.21B actual; both 128K context; trained on up to 9T tokens; December 2023 knowledge cutoff; Grouped-Query Attention; released September 25, 2024; governed by the Llama 3.2 Community License (custom commercial) [SECONDARY]. Sources: https://huggingface.co/meta-llama/Llama-3.2-3B and https://github.com/budecosystem/bud-connect/blob/HEAD/scripts/model_docs_workflow/raw_docs/bedrock/meta_llama_models.md
- Llama 3.2 1B/3B are text-only (text in/text out) instruction-tuned models optimized for multilingual dialogue, agentic retrieval, and summarization; the quantized variants drop to 8K context [SECONDARY]. Source: https://hf.global-rail.com/meta-llama/Llama-3.2-1B/resolve/main/README.md?download=true
- Llama 3.2 4-bit VRAM: 1B ~0.75GB, 3B ~1.75GB — the 1B runs on a 4GB Quadro P1000, the 3B on an 8GB RTX 3060 Ti [SECONDARY]. Source: https://github.com/budecosystem/bud-connect/blob/HEAD/scripts/model_docs_workflow/raw_docs/bedrock/meta_llama_models.md
- Llama 3.2 family also includes 11B and 90B Vision (multimodal text+image), all 128K context [SECONDARY]. Source: https://github.com/budecosystem/bud-connect/blob/HEAD/scripts/model_docs_workflow/raw_docs/bedrock/meta_llama_models.md
- Mistral 3 family (launched as a 10-model release): Mistral Large 3 plus nine Ministral 3 variants (3 sizes × 3 use cases: base, chat, reasoning) — all Apache 2.0, all open weights [SECONDARY]. Sources: https://airevolution.poltextlab.com/mistral-launches-its-new-10-model-release/ and https://docs.redhat.com/en/documentation/red_hat_ai_inference/3.5/pdf/inference_serving_mistral_3_models/Red_Hat_AI_Inference-3.5-Inference_serving_Mistral_3_models-en-US.pdf?utm_source=openai
- Mistral Large 3 architecture: 675B total / 41B active sparse MoE, DeepSeek-V3-style with fewer larger experts, top-4 expert selection with softmax routing, Llama 4 RoPE scaling; ~2.5B-parameter vision encoder fused natively (not bolted on) [SECONDARY]. Sources: https://airevolution.poltextlab.com/mistral-launches-its-new-10-model-release/ and https://medium.com/@leucopsis/mistral-large-3-2512-review-7788c779a5e4
- Mistral Large 3 debuted at #2 in LMArena's open-source non-reasoning category — "amongst the first open frontier models with multimodal and multilingual capabilities" [SECONDARY]. Source: https://airevolution.poltextlab.com/mistral-launches-its-new-10-model-release/
- Context-length contradiction ledger: older secondary sources list Large 3 at 128K, Small 4 at 131K, Medium 3.5 at 131K; a 2026-09-08 catalog check lists Large 3 at 256k, Small 4 at 256k (119B MoE/6.5B active), Medium 3.5 at 256k (128B dense) — whether this is a context extension or a correction is unresolved; do not quote a single figure [SECONDARY — contradiction]. Sources: https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/ versus https://github.com/samouraiworld/awesome-mistral/blob/HEAD/README.md
- Mistral Small 4 is a 119B MoE with 6.5B active parameters per the September catalog check — not a dense small model [SECONDARY]. Source: https://github.com/samouraiworld/awesome-mistral/blob/HEAD/README.md
- Mistral Medium 3.5 is 128B dense per the September catalog check [SECONDARY]. Source: https://github.com/samouraiworld/awesome-mistral/blob/HEAD/README.md
- Ministral 3 all carry built-in vision encoders and 256K context (3B/8B/14B) per Red Hat docs — the 3B at 128K in older pricing pages is superseded [SECONDARY]. Source: https://docs.redhat.com/en/documentation/red_hat_ai_inference_server/3.3/pdf/inference_serving_mistral_3_models/Red_Hat_AI_Inference_Server-3.3-Inference_serving_Mistral_3_models-en-US.pdf
- Red Hat's assessment: "All Mistral 3 models are released under the Apache 2.0 license with open weights" and are fully compatible with upstream vLLM, requiring no custom forks [SECONDARY]. Source: https://docs.redhat.com/en/documentation/red_hat_ai_inference/3.5/pdf/inference_serving_mistral_3_models/Red_Hat_AI_Inference-3.5-Inference_serving_Mistral_3_models-en-US.pdf?utm_source=openai
- Mistral's strategy per an OSS foundations analysis: "open weights, proprietary platform" — weights (Apache 2.0) are the distribution layer; revenue comes from La Plateforme API, Forge training, and Mistral Compute [SECONDARY]. Source: https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/mistral-ai-deep-dive.md
- Devstral 2 is co-developed with All Hands AI (OpenHands) and is Apache 2.0; Codestral is Apache 2.0 (latest); Voxtral TTS is CC BY-NC 4.0; Mistral OCR is proprietary API-only [SECONDARY]. Source: https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/mistral-ai-deep-dive.md
- Mistral OCR 4.1 (not 4.0): $4/1K pages, 170 languages, bounding boxes and block-level confidence — the versioned name matters [SECONDARY]. Source: https://github.com/samouraiworld/awesome-mistral/blob/HEAD/README.md

**Additional facts, fifth tranche:**

- Llama 3.2 11B Vision: multimodal text+image instruction-tuned model, 128K context, Llama 3.2 Community License — the vision-capable sibling of the 1B/3B text models [SECONDARY]. Source: https://github.com/budecosystem/bud-connect/blob/HEAD/scripts/model_docs_workflow/raw_docs/bedrock/meta_llama_models.md
- Llama 3.2 supported languages (all sizes): English, German, French, Italian, Portuguese, Hindi, Spanish, Thai — eight officially [SECONDARY]. Source: https://github.com/budecosystem/bud-connect/blob/HEAD/scripts/model_docs_workflow/raw_docs/bedrock/meta_llama_models.md
- Llama 3.2 uses Grouped-Query Attention and shared input/output embeddings for inference efficiency [SECONDARY]. Source: https://huggingface.co/meta-llama/Llama-3.2-3B
- Mistral OCR 4.1: $4 per 1,000 pages, 170 languages, bounding boxes and block-level confidence scores — proprietary API-only, no weights [SECONDARY]. Source: https://github.com/samouraiworld/awesome-mistral/blob/HEAD/README.md
- Devstral 2: co-developed with All Hands AI (OpenHands), Apache 2.0, agentic coding focus [SECONDARY]. Source: https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/mistral-ai-deep-dive.md
- Codestral (latest): Apache 2.0 — Mistral's code model joined the Apache line [SECONDARY]. Source: https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/mistral-ai-deep-dive.md
- The "Mistral 3" 10-model release: Large 3 plus nine Ministral 3 variants (3B/8B/14B × base/chat/reasoning) — the reasoning variants target chain-of-thought workloads at the edge [SECONDARY]. Source: https://airevolution.poltextlab.com/mistral-launches-its-new-10-model-release/

**Additional facts, sixth tranche:**

- Llama 3.2 90B Vision: the largest Llama 3.2, multimodal, 128K context, Llama 3.2 Community License [SECONDARY]. Source: https://github.com/budecosystem/bud-connect/blob/HEAD/scripts/model_docs_workflow/raw_docs/bedrock/meta_llama_models.md
- The 8K-context quantized Llama 3.2 variants are text-only — vision requires the full 128K models [SECONDARY]. Source: https://hf.global-rail.com/meta-llama/Llama-3.2-1B/resolve/main/README.md?download=true

