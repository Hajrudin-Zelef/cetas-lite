---
id: collect-huggingface/huggingface/coherelabs-north-small-translate-1-0
title: "North-Small-Translate-1.0 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["China", "Cohere", "Hugging Face", "vLLM"]
dates: ["2026-09-23"]
keywords: ["agentic", "attention", "benchmarks", "cohere", "embeddings", "fp8", "license", "memory", "moe", "nvfp4", "open-weight", "parameters"]
source: docs/RAG/Collect RAG/03_huggingface/CohereLabs-North-Small-Translate-1.0.md
source_anchor: ""
source_lines: [1, 54]
sha256: 39460196cd79472b53c8f272a6661f33f5156b3c14b0cccabe0d7c9ac364d6ef
---

# North-Small-Translate-1.0 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/CohereLabs/North-Small-Translate-1.0
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

North Small Translate is an open-weights research release from Cohere and Cohere Labs of a sparse Mixture-of-Experts model specialized for high-quality machine translation across 50 languages. It has 25B active parameters and 218B total parameters, with 128 experts of which 8 are activated per token, alongside shared experts applied to every token. The attention layers interleave sliding-window attention (window size 4096) with RoPE against global attention layers without positional embeddings, in a 3:1 ratio, as first introduced in Command A. The router applies a sigmoid activation over expert logits and normalizes over the selected top-k. The model was post-trained specifically for translation quality. Architecture tag: cohere2_moe; decoder-only sparse MoE Transformer, text-in/text-out. Context length: 16K input and 16K output.

Languages: 50 languages including English, Arabic, Chinese (Traditional/Simplified), Hindi, Japanese, Korean, Russian, Spanish, French, German, and many others. Evaluation: WMT26 all-languages score of 83.60, increasing to 84.36 with the agentic multi-pass translation workflow.

Usage: Transformers (install from source repo with the necessary changes; note the "translation" pipeline type is deprecated in transformers v5, so load the model directly or use transformers<5.0.0). Recommended greedy decoding (production setting). On smaller nodes, pass max_memory with device_map="auto" (e.g., 8 x H100 with 70GiB each) and PYTORCH_CUDA_ALLOC_CONF=expandable_segments:True. The model wraps replies in structural markers (<|START_TEXT|>, <|END_TEXT|>) that are not registered special tokens; parse completions with Cohere's melody library (PyFilter cmd4 start_in_answer). The chat template applies a default system instruction replaceable via platform_instruction_override. vLLM serving requires cohere_melody>=0.9.0 and the cohere_command4 tool-call/reasoning parsers.

License: CC BY-NC 4.0 (non-commercial) plus Cohere Labs Acceptable Use Policy; commercial use via Cohere sales. Access requires accepting the license gated form. Quantizations: BF16 (4 x B200 / 8 x H100), FP8 (2 x B200 / 4 x H100), NVFP4 W4A16 (1 x B200 / 2 x H100), all production-served. Hub: 218B params BF16, ~122 downloads/month.

## Key points

- Open-weights translation MoE: 25B active / 218B total, 128 experts (8 per token) + shared experts.
- Decoder-only sparse MoE Transformer; interleaved SWA/global attention 3:1 (Command A design).
- Post-trained specifically for translation; 50 languages, 16K in / 16K out context.
- WMT26 score 83.60 (84.36 with agentic multi-pass workflow).
- Gated access: CC BY-NC 4.0 license + Acceptable Use Policy (non-commercial).
- Reply wrapped in structural markers; parse with Cohere melody library.
- Production quantizations: BF16, FP8, NVFP4 W4A16 (all served by Cohere).
- Greedy decoding recommended; vLLM with cohere_command4 parsers.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | CohereLabs |
| Model name | North-Small-Translate-1.0 |
| Architecture | Decoder-only sparse MoE (cohere2_moe) |
| Total params | 218B |
| Activated params | 25B |
| Experts | 128 (8 per token) + shared experts |
| Attention | SWA (window 4096, RoPE) / global 3:1 |
| Context length | 16K input / 16K output |
| Languages | 50 |
| License | CC BY-NC 4.0 + Acceptable Use Policy |
| Weight formats | BF16 (also FP8, NVFP4 W4A16) |
| Benchmarks | WMT26 83.60 (84.36 multi-pass) |
| Hardware (BF16) | 4 x B200 or 8 x H100 |
| Downloads/month | 122 |

## Why this source matters for the RAG

This card documents a frontier open-weight translation-specific MoE with detailed architecture, quantization tiers, and production serving practices (marker parsing, greedy decoding, memory config). It is essential reference material for retrieval on machine translation, sparse MoE architectures, and gated non-commercial model releases.
