---
id: collect-huggingface/huggingface/unsloth-gemma-4-12b-it-gguf
title: "gemma-4-12b-it-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["China", "Google", "Hugging Face", "Unsloth"]
dates: ["2026-09-23"]
keywords: ["gguf", "apache", "attention", "benchmarks", "context window", "embedding", "inference", "latency", "license", "llama", "llama.cpp", "memory"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-gemma-4-12b-it-GGUF.md
source_anchor: ""
source_lines: [1, 53]
sha256: 81802388fbd32be895a98efe2f86b1fdafa505c7d588a1c5f547d136920e3ca9
---

# gemma-4-12b-it-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/gemma-4-12b-it-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth's Dynamic 2.0 GGUF quantization repository for Google DeepMind's Gemma 4 12B Unified instruction-tuned model. Gemma 4 12B Unified is part of the Gemma 4 family of open multimodal models and is notable for its encoder-free architecture: other Gemma 4 models use dedicated encoders to process multimodal data, but the 12B Unified model projects raw image patches and audio waveforms directly into the LLM's embedding space through lightweight linear layers. This unified approach makes it a true "omni" model—the same GGUF files handle text, images, and audio—and reduces multimodal latency while allowing the whole model to be fine-tuned in one pass. The 12B model has 11.95B parameters, 48 layers, a 1024-token sliding window, a 256K-token context window, a 262K vocabulary, and supports text, image, and audio (video as frames). It uses hybrid attention interleaving local sliding-window attention with full global attention, with unified Keys/Values and Proportional RoPE (p-RoPE) on global layers for long-context memory efficiency. Gemma 4 adds configurable thinking modes (enabled via `<|think|>` in the system prompt), native system-prompt support, variable image resolution (token budgets 70/140/280/560/1120), and native function calling. Benchmarks for the 12B Unified: MMLU Pro 77.2%, AIME 2026 77.5%, LiveCodeBench v6 72.0%, Codeforces ELO 1659, GPQA Diamond 78.8%, Tau2 69.0%, MMMLU 83.4%, MMMU Pro 69.1%, MATH-Vision 79.7%, CoVoST 38.5 (excl. Chinese), FLEURS 0.069, and MRCR v2 128k 43.4%. Recommended sampling: temperature 1.0, top_p 0.95, top_k 64. Quant sizes range from UD-IQ2_M (4.21GB) through UD-Q4_K_XL (7.37GB), Q8_0 (12.7GB), to BF16 (23.8GB); Multi-Token Prediction (MTP) drafter files are included. License is Apache 2.0; 1,099,648 monthly downloads.

## Key points

- Unsloth Dynamic 2.0 GGUF quantization of Gemma 4 12B Unified (encoder-free, omni multimodal).
- Same GGUF files handle text, image, and audio; video as frames; 256K context; 11.95B params.
- Hybrid attention (1024-token sliding window + global), unified K/V, p-RoPE; 262K vocabulary.
- Configurable thinking mode, native system prompts, variable image resolution, native function calling.
- Benchmarks: MMLU Pro 77.2%, GPQA-D 78.8%, AIME 2026 77.5%, LiveCodeBench v6 72.0%, MMMU Pro 69.1%.
- Quant sizes from UD-IQ2_M 4.21GB to BF16 23.8GB; recommended UD-Q4_K_XL 7.37GB; MTP drafters included.
- Run with llama.cpp using `--jinja`; mmproj auto-downloads with `-hf`; supports vision/audio.
- Apache 2.0 license; recommended sampling temperature 1.0, top_p 0.95, top_k 64.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | gemma-4-12b-it-GGUF |
| Base model | google/gemma-4-12B-it |
| Architecture | Gemma 4 12B Unified (encoder-free dense, hybrid attention) |
| Parameters | 11.95B |
| Layers | 48 |
| Sliding window | 1024 tokens |
| Context length | 256K tokens |
| Vocabulary | 262K |
| Modalities | Text, Image, Audio (video as frames) |
| Quantization | Unsloth Dynamic 2.0 GGUF (2-bit to 16-bit) |
| Recommended quant | UD-Q4_K_XL (7.37GB) |
| Smallest quant | UD-IQ2_M (4.21GB) |
| Largest quant | BF16 (23.8GB) |
| MTP drafters | MTP Q8_0 (465MB), MTP BF16 (862MB) |
| Sampling | temperature 1.0, top_p 0.95, top_k 64 |
| License | Apache 2.0 |
| Key benchmarks | MMLU Pro 77.2%; GPQA-D 78.8%; AIME 2026 77.5%; LiveCodeBench v6 72.0% |
| Downloads/month | 1,099,648 |

## Why this source matters for the RAG

This card documents a popular local-deployment path for Gemma 4 12B Unified, emphasizing its encoder-free omni modality (text/image/audio in one GGUF) and Unsloth Dynamic 2.0 quantization. It is a key reference for on-device multimodal inference and GGUF quantization trade-offs.
