---
id: ai-industry-kb-2026/09-moe-architectures/specification-anchors-where-each-model-s-specs-were-verified
title: "Specification anchors: where each model's specs were verified"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["Alibaba", "Anthropic", "DeepSeek", "Hugging Face", "MiniMax", "Moonshot", "OpenRouter", "SGLang", "Together AI", "Xiaomi", "Z.ai"]
dates: ["2026-08-22"]
keywords: ["agentic", "apache", "attention", "benchmark", "benchmarks", "claude", "compute", "decode", "deepseek", "embeddings", "fp8", "gguf"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4737, 4756]
section: "9. MoE Architectures"
sha256: e4218fffe69f9cc6052dee9cc37240637a05a9da0c22706d61e689caa9197d8f
---

# Specification anchors: where each model's specs were verified

- [VENDOR] MiniMax M3 "~9× prefill / ~15× decode over M2 at 1M context; ~1/20 per-token compute of the previous generation" (MSA claims); the Morph-sourced "matches Claude Sonnet 4.6 on real-world agentic benchmarks" — [UNVERIFIED independently].
- [VENDOR] MiMo-V2-Flash "150 tok/s via MTP self-speculative decoding; accepted length 2.8–3.6; 2.0–2.6× effective speedup; learnable sink bias → ~6× KV reduction; trained on 27T tokens in FP8"; the 150 tok/s figure is independently plausible given MTP self-speculation but no third-party reproduction was found in this pass — [DIRECTIONAL].
- [VENDOR] Kimi K3 "up to 6.3× faster decoding (KDA); ~25% higher training efficiency (AttnRes); 1.72–2.22× prefill on H20 (FlashKDA); ~2.5× scaling efficiency over K2"; the technical-report mechanisms (latent width, QB, MXFP4 geometry) are primary-source and independently inspected by multiple community analyses — [COMMUNITY]-confirmed as present in the weights/config.
- [VENDOR] Qwen3.8-Max benchmark table (~40 rows, 20 methodology footnotes; 93.0 PaperBench, 92.6% GPQA Diamond, 67.7% SWE-bench Pro, 73.5% FrontierSWE) — vendor-chosen harnesses; treat cross-vendor comparisons as [DIRECTIONAL]; third-party anchors (Artificial Analysis, community reproductions) are the trustworthy cross-model numbers.
- [VENDOR] GLM-5.3-Flash "hybrid sparse+linear attention → 4.4× smaller KV cache"; "beats GLM-5.2 at ~1/10 the price" — directional until independently measured.
- [COMMUNITY] Qwen3.6 61 tok/s vs 7 tok/s dense on RTX 5070; Gemma 4 expert-offloading profile (~4.2–4.8 GB hot set); Qwen3.8-Max FP8-out-downloading-bf16 adoption signal — practitioner-measured, hardware-specific, not lab-verified.
- [UNVERIFIED] The ~50B-active figure for Kimi K3 (one Medium analysis) against the 104B technical-report consensus; the DeepSeek-V4 secondary architecture reconstruction (treat every number as [PARTIALLY VERIFIED]); MiMo-V2.6 figures (single source); DeepSeek-V4 license (MIT vs Apache 2.0); GLM-5.3-Flash 300K-context mention; M3's 7 MTP modules (community inference notes, absent from official card).

### Specification anchors: where each model's specs were verified

- gpt-oss-120b/20b: arXiv 2508.10925 model card (parameter breakdown 114.71B MLP / 0.96B attention / 1.16B embed-unembed for the 120b; checkpoint sizes 60.8 GiB / 12.8 GiB).
- Qwen3.6-35B-A3B: Hugging Face model card + Together AI model page (`together.ai/models/qwen3-6-35b-a3b-fp8`) + community GGUF/quant tables (knightli.com, allthings.how).
- Gemma 4 26B-A4B: HF model card + gemma4.dev model page + community offload architecture doc (alexchen31337/gemma4-moe-offload).
- MiniMax M3: HF model card + SGLang cookbook (solrex/sglang fork) + Artificial Analysis intelligence-index entries + Morph model page.
- DeepSeek-V4: official announcement + API docs for benchmarks/pricing; architecture numbers only via secondary reconstructions (unofficial HF technical-report mirror + community case studies) — [PARTIALLY VERIFIED].
- Kimi K3: Moonshot primary technical report (`MoonshotAI/Kimi-K3` repo PDF) + HF model card + multiple independent community architecture inspections.
- Qwen3.8-Max: HF repos `Qwen/Qwen3.8-2.4T-A95B` and `-FP8` (download/like counts as of 2026-08-22) + release coverage.
- GLM-5.3-Flash: Z.ai announcement + model card `config.json` (`max_position_embeddings` = 1,048,576) + HF repo `zai-org/GLM-5.3-Flash` + OpenRouter "Ox Alpha" chart history.
- MiMo-V2-Flash: Xiaomi announcement + upstream SGLang inference code; MiMo-V2.6-Flash/Pro: OrcaRouter release analysis only — [UNVERIFIED as definitive].

