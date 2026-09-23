---
id: collect-huggingface/huggingface/unsloth-minimax-h3-gguf
title: "MiniMax-H3-GGUF (unsloth) - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "MiniMax", "Unsloth"]
dates: ["2026-09-23"]
keywords: ["gguf", "attribution", "diffusion", "fp8", "gpus", "license", "omni", "quantization", "video generation"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-MiniMax-H3-GGUF.md
source_anchor: ""
source_lines: [1, 50]
sha256: 7a58a9b97b032c95c85f4f4b842def45f91ccbb0e24c679490ce09859ffbdcab
---

# MiniMax-H3-GGUF (unsloth) - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/MiniMax-H3-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth's GGUF quantization release of MiniMaxAI/MiniMax-H3, enabling local execution of the omni-modal video generation system on most platforms, including stable-diffusion.cpp (sd-cli) and Unsloth. MiniMax H3 produces video with native stereo audio, up to 15 seconds at 24 FPS with 32 kHz stereo audio. The repository contains both halves of the runtime: the two denoisers (H3-Base variants) and the Qwen3-VL text encoder they require.

Two denoiser checkpoints are provided as separate GGUF files, not settings: `fl2va_pruned` (H3-Base first-and-last-frame variant: text plus zero, one, or two frames) and `ref2va_pruned` (reference variant: text plus reference pictures, videos, and audio). Both are quantized at the same rungs. File listing (fl2va_pruned): Q2_K 6.26 GiB, UD-Q2_K_XL 7.51 GiB, Q3_K 8.16 GiB, UD-Q3_K_XL 8.90 GiB, Q4_K 10.64 GiB, Q5_0 12.97 GiB, Q6_K 15.45 GiB, Q8_0 19.97 GiB. ref2va_pruned equivalents: Q2_K 6.22 GiB, Q3_K 8.12 GiB, Q4_K 10.60 GiB, Q5_0 12.94 GiB, Q6_K 15.42 GiB, Q8_0 19.94 GiB. Shared text encoder `qwen3vl_32b_minimax_h3`: Q2_K_M 12.20 GiB and Q4_K_M 16.97 GiB. The UD- rungs are Unsloth Dynamic 2.0 mixed-precision builds; pair the Q2_K_M text encoder with the two smallest denoisers and Q4_K_M with everything else. VAEs are shared and taken from Comfy-Org/MiniMax-H3.

Example run (sd-cli) requires three non-optional flags: `--mode vid_gen`, explicit `--cfg-scale 1.0` (H3 is CFG-free/distilled and aborts above 1.0 while the default is 7.0), and `--backend te=cpu` (keeps the 12 GB text encoder off the card); add `--offload-to-cpu` for smaller GPUs. License: MiniMax H3 Community License Agreement; files are Model Derivatives (quantization changes numerics) so a NOTICE lists every change with attribution. Not an official MiniMax product. Pre-quantized PyTorch checkpoints are in unsloth/MiniMax-H3-FP8. Hub reports 20B params and 1.13M monthly downloads.

## Key points

- Community GGUF quantization of MiniMax-H3 for local use via sd-cli / stablediffusion.cpp / Unsloth.
- Both denoisers shipped: fl2va_pruned and ref2va_pruned (separate checkpoints, not settings).
- Shared Qwen3-VL-32B text encoder (Q2_K_M / Q4_K_M GGUF); VAEs from Comfy-Org/MiniMax-H3.
- Quants from Q2_K (~6.2 GiB) to Q8_0 (~20 GiB), including Unsloth Dynamic UD- rungs.
- H3 generates 32 kHz stereo audio jointly with video (audio part of model output).
- Required sd-cli flags: --mode vid_gen, --cfg-scale 1.0, --backend te=cpu.
- MiniMax H3 Community License; NOTICE lists derivative changes.
- Hub reports 20B params, 1.13M downloads/month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | MiniMax-H3-GGUF |
| Base model | MiniMaxAI/MiniMax-H3 |
| Format | GGUF (stable-diffusion.cpp compatible) |
| Denoisers | fl2va_pruned, ref2va_pruned |
| Text encoder | qwen3vl_32b_minimax_h3 (Q2_K_M 12.2 GiB / Q4_K_M 17.0 GiB) |
| Quant rungs | Q2_K to Q8_0 + UD-Q2_K_XL / UD-Q3_K_XL |
| Denoiser sizes | ~6.2 GiB (Q2_K) to ~19.9 GiB (Q8_0) |
| Hub-reported params | 20B |
| Output | Video + 32 kHz stereo audio, up to 15s / 24 FPS |
| License | MiniMax H3 Community License |
| Downloads/month | 1,125,973 |

## Why this source matters for the RAG

This card documents practical local quantization of a frontier audio-video generation model, with exact file sizes, quantization rungs, and critical runtime flags, plus derivative-license handling. It is valuable reference for retrieval on GGUF quantization, local video-model deployment, and community conversion practices.
