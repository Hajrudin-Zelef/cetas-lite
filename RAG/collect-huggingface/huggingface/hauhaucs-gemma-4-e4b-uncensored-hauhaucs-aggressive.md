---
id: collect-huggingface/huggingface/hauhaucs-gemma-4-e4b-uncensored-hauhaucs-aggressive
title: "Gemma-4-E4B-Uncensored-HauhauCS-Aggressive - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Google", "Hugging Face"]
dates: ["2026-09-23"]
keywords: ["attention", "context window", "embeddings", "gguf", "inference", "license", "llama", "llama.cpp", "memory", "multimodal", "parameters", "quantization"]
source: docs/RAG/Collect RAG/03_huggingface/HauhauCS-Gemma-4-E4B-Uncensored-HauhauCS-Aggressive.md
source_anchor: ""
source_lines: [1, 52]
sha256: d19bcd5987d16fe26698c4b86dcdf399f95970007a8f1f3c2e3c3ab493945688
---

# Gemma-4-E4B-Uncensored-HauhauCS-Aggressive - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/HauhauCS/Gemma-4-E4B-Uncensored-HauhauCS-Aggressive
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is HauhauCS's uncensored ("abliterated") GGUF release of Google's Gemma 4 E4B instruction-tuned model, distributed under the Gemma license. The author states that no changes were made to datasets or capabilities—the model is fully functional and "100% of what the original authors intended," but with refusals removed. The card claims 0/465 refusals in testing. This "Aggressive" variant is the stronger uncensoring: the model is fully unlocked and won't refuse prompts, though it may occasionally append short disclaimers baked into the base model training (not refusals). A more conservative "Balanced" variant is referenced as forthcoming. The model is natively multimodal (text, image, video, audio), with 4B parameters, 42 layers using mixed sliding-window (512) plus full attention, a 131K context window, and 18 KV-shared layers for memory efficiency; it is based on google/gemma-4-e4b-it. A companion mmproj (f16, 945 MB) file enables vision/audio. The repository provides many GGUF quantizations, all generated with an importance matrix (imatrix) to preserve quality on the abliterated weights. Notable files and sizes: Q8_K_P (9.4 BPW, 7.6 GB), Q6_K_P (7.0 BPW, 5.9 GB), Q5_K_P (6.1 BPW, 5.5 GB), Q5_K_M (5.7 BPW, 5.4 GB), Q4_K_P (5.2 BPW, 5.1 GB), Q4_K_M (4.8 BPW, 5.0 GB), IQ4_XS (4.3 BPW, 4.8 GB), Q3_K_P (4.1 BPW, 4.6 GB), Q3_K_M (3.9 BPW, 4.6 GB), IQ3_M (3.7 BPW, 4.4 GB), and Q2_K_P (3.5 BPW, 4.2 GB). K_P ("Perfect") quants are custom HauhauCS profiles that selectively preserve quality, effectively bumping quality by 1-2 quant levels at only ~5-15% larger size, and remain fully llama.cpp-compatible. Recommended settings from the Gemma 4 authors: temperature 1.0, top_p 0.95, top_k 64; use the `--jinja` flag and the mmproj file for vision/audio. The hub reports 8B params (including embeddings) and 1,679,485 monthly downloads.

## Key points

- Uncensored/abliterated GGUF build of Google Gemma 4 E4B-IT by HauhauCS; claimed 0/465 refusals.
- "Aggressive" variant: fully unlocked, won't refuse; may append short baked-in disclaimers.
- 4B parameters, 42 layers, mixed sliding-window (512) + full attention, 131K context, 18 KV-shared layers.
- Natively multimodal (text, image, video, audio); based on google/gemma-4-e4b-it.
- Extensive GGUF range Q2_K_P to Q8_K_P; separate mmproj f16 file (945 MB) for vision/audio.
- Custom K_P ("Perfect") quants improve quality 1-2 levels at ~5-15% larger size; imatrix-based.
- Recommended sampling temperature 1.0, top_p 0.95, top_k 64; use `--jinja`.
- Gemma license; 1,679,485 downloads/month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | HauhauCS |
| Model name | Gemma-4-E4B-Uncensored-HauhauCS-Aggressive |
| Base model | google/gemma-4-e4b-it |
| Architecture | gemma4 (dense, hybrid attention) |
| Parameters | 4B (hub: 8B with embeddings) |
| Layers | 42 |
| Sliding window | 512 tokens |
| Context length | 131K tokens |
| Modalities | Text, Image, Video, Audio |
| KV shared layers | 18 |
| Quantization | GGUF, Q2_K_P to Q8_K_P + IQ variants |
| Flagship quant | Q8_K_P (9.4 BPW, 7.6 GB) |
| Smallest quant | Q2_K_P (3.5 BPW, 4.2 GB) |
| mmproj | f16, 945 MB (vision/audio) |
| Recommended sampling | temperature 1.0, top_p 0.95, top_k 64 |
| Refusals | 0/465 claimed |
| License | Gemma |
| Downloads/month | 1,679,485 |

## Why this source matters for the RAG

This card documents a popular community uncensoring of Gemma 4 E4B distributed as GGUF, illustrating abliteration, custom K_P quantization, and multimodal local deployment. It is a key reference for uncensored/abliterated models and quantized on-device multimodal inference.
