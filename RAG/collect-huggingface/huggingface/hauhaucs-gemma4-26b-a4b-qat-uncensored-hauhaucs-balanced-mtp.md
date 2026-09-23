---
id: collect-huggingface/huggingface/hauhaucs-gemma4-26b-a4b-qat-uncensored-hauhaucs-balanced-mtp
title: "Gemma4-26B-A4B-QAT-Uncensored-HauhauCS-Balanced-MTP - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Google", "Hugging Face", "Unsloth"]
dates: ["2026-09-23"]
keywords: ["agentic", "benchmarks", "context window", "gguf", "gpu", "license", "llama", "llama.cpp", "moe", "multimodal", "parameters", "quantization"]
source: docs/RAG/Collect RAG/03_huggingface/HauhauCS-Gemma4-26B-A4B-QAT-Uncensored-HauhauCS-Balanced-MTP.md
source_anchor: ""
source_lines: [1, 49]
sha256: e67830d39ab3d81be3338b036394f4a61805e0dc0d43a29bab3ef02d38e3f3ee
---

# Gemma4-26B-A4B-QAT-Uncensored-HauhauCS-Balanced-MTP - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/HauhauCS/Gemma4-26B-A4B-QAT-Uncensored-HauhauCS-Balanced-MTP
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is HauhauCS's uncensored GGUF release of Google's Gemma 4 26B A4B QAT instruction-tuned model, distributed under the Gemma license. The author states there are no changes to datasets or capabilities—the model is fully functional, "100% of what the original authors intended, just without the refusals"—and it is built from the official QAT weights so the 4-bit quant stays close to full-precision quality. The card claims 0/465 refusals (tested with both automated and manual refusal benchmarks), noting a small number of edge-case prompts deflect on the first ask but comply on re-ask. The "Balanced" variant is recommended (99%+ of users) and uses optimized full uncensoring tuned especially for agentic coding, reasoning, creative writing, and reliability-critical tasks; it reasons before answering and stays dependable and on-instruction. An "Aggressive" variant is not currently required. The release ships with an MTP (multi-token-prediction) draft head for speculative decoding, claimed to be roughly 35% faster generation with identical output (the model verifies every drafted token), tested via llama.cpp. The underlying model is a Gemma 4 26B-A4B MoE with 128 experts (8 active per token) and a 256K (262144) context window, supporting vision (image input) via an mmproj file. Available files: the Q4_K_M text GGUF (16.8 GB), the mmproj BF16 vision projector (1.2 GB), and the MTP speculative drafter `mtp-gemma-4-26B-A4B-it.gguf` (252 MB). The author notes Gemma 4 is QAT-trained for ~4-bit, so Q4_K_M is the sweet spot and higher-precision quants add size without real quality gain. Recommended sampling is tuned for this build: temperature 0.6, top_k 64, top_p 0.9, min_p 0.05, repeat_penalty 1.1 (not the stock Gemma defaults). It works with llama.cpp, LM Studio, Jan, koboldcpp, and other GGUF runtimes; a caveat is noted that Gemma 4 can crash under LM Studio's tensor-split mode, so a single GPU (layer-split or priority order) is advised. The MTP draft head comes from Unsloth's Gemma 4 release. The hub reports 25B params and 354,490 monthly downloads.

## Key points

- Uncensored GGUF of Google Gemma 4 26B A4B (QAT) by HauhauCS; claimed 0/465 refusals.
- Built from official QAT weights, so 4-bit quality stays close to full precision.
- "Balanced" variant tuned for agentic coding, reasoning, creative writing, and reliability.
- Ships an MTP draft head for speculative decoding, ~35% faster with identical output.
- MoE: 128 experts (8 active per token), 256K context; vision via mmproj file.
- Files: Q4_K_M 16.8 GB, mmproj BF16 1.2 GB, MTP drafter 252 MB.
- Custom sampling recommended: temperature 0.6, top_k 64, top_p 0.9, min_p 0.05, repeat_penalty 1.1.
- Gemma license; works with llama.cpp, LM Studio, Jan, koboldcpp; single-GPU recommended.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | HauhauCS |
| Model name | Gemma4-26B-A4B-QAT-Uncensored-HauhauCS-Balanced-MTP |
| Base model | google/gemma-4-26B-A4B-it (QAT) |
| Architecture | gemma4 MoE (128 experts, 8 active) |
| Parameters | 26B-A4B (hub: 25B) |
| Context length | 256K (262144) |
| Modalities | Text, Image (vision via mmproj) |
| Quantization | GGUF Q4_K_M (16.8 GB) |
| mmproj | BF16 vision projector (1.2 GB) |
| MTP drafter | mtp-gemma-4-26B-A4B-it.gguf (252 MB) |
| Speculative decoding | `--spec-type draft-mtp`, ~35% faster |
| Refusals | 0/465 claimed |
| Recommended sampling | temp 0.6, top_k 64, top_p 0.9, min_p 0.05, repeat_penalty 1.1 |
| License | Gemma |
| Downloads/month | 354,490 |

## Why this source matters for the RAG

This card documents a community uncensoring of Gemma 4 26B A4B QAT combined with MTP speculative decoding, illustrating abliteration on QAT weights and the Balanced uncensoring philosophy. It is a key reference for uncensored MoE models, speculative decoding, and local multimodal deployment.
