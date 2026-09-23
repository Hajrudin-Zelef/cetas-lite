---
id: collect-huggingface/huggingface/unsloth-minimax-m3-gguf
title: "MiniMax-M3-GGUF (unsloth) - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "MiniMax", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "agent", "attention", "compute", "decode", "gpus", "inference", "license", "llama", "llama.cpp", "moe", "multimodal"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-MiniMax-M3-GGUF.md
source_anchor: ""
source_lines: [1, 54]
sha256: 4eae63161ba510325ee750f5fe04828d77b56e2034e453e36c9d7e64cc2ab613
---

# MiniMax-M3-GGUF (unsloth) - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/MiniMax-M3-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth's GGUF quantization release of MiniMaxAI/MiniMax-M3, providing experimental llama.cpp support for the ~428B-parameter multimodal MoE. It is an Image-Text-to-Text model (Transformers + GGUF libraries) that can run locally via llama.cpp, vLLM, SGLang, Ollama, LM Studio, Jan, Unsloth Desktop, and various agent tools (Pi, Hermes, OpenClaw, Lemonade). Architecture: MoE + MiniMax Sparse Attention (MSA), ~428B total / ~23B activated, 128 experts with 4 active per token, 60 layers, 1M-token context, text/image/video modalities, bfloat16 precision, transformers >= 4.52.4 with trust_remote_code. M3 delivers 9x prefill and 15x decode speedups over M2 at 1M context (per-token compute reduced to 1/20). Two reasoning modes: thinking and non-thinking.

Important caveat: MiniMax-M3 support in llama.cpp is preliminary and not yet in a released build; the card instructs building llama.cpp from PR #24523 (checkout minimax-m3 branch) before running the GGUFs. Also, MiniMax Sparse Attention is not yet supported, so inference falls back to dense attention. The model is large (~428B params), so offload across GPUs with -ngl 99 or keep weights in CPU RAM.

Quantization lineup (hub-reported sizes): 1-bit UD-IQ1_M (128 GB), 2-bit UD-IQ2_XXS (134 GB), UD-IQ2_M (134 GB), UD-Q2_K_XL (143 GB), 3-bit UD-IQ3_XXS (159 GB), UD-IQ3_S (175 GB), UD-Q3_K_M (195 GB), UD-Q3_K_XL (195 GB), 4-bit UD-IQ4_XS (208 GB), UD-IQ4_NL (212 GB), UD-Q4_K_S (248 GB), MXFP4_MOE (256 GB), UD-Q4_K_M (264 GB), UD-Q4_K_XL (265 GB), 5-bit UD-Q5_K_S (299 GB), UD-Q5_K_M/XL (318 GB), 6-bit UD-Q6_K (354 GB), UD-Q6_K_XL (387 GB), 8-bit Q8_0 (453 GB), UD-Q8_K_XL (464 GB), 16-bit BF16 (852 GB). License: MiniMax Community License. Recommended inference parameters: temperature=1.0, top_p=0.95, top_k=40. Hub reports 426B params and 6,196 monthly downloads.

## Key points

- Experimental GGUF quantization of MiniMax-M3 for llama.cpp (requires building PR #24523).
- ~428B total / ~23B activated MoE, 128 experts (4 active), 60 layers, 1M context.
- Multimodal (text/image/video); transformers>=4.52.4 trust_remote_code.
- MiniMax Sparse Attention not yet supported -> dense-attention fallback.
- Quants from 1-bit UD-IQ1_M (128 GB) to 16-bit BF16 (852 GB).
- Runs via llama.cpp, vLLM, SGLang, Ollama, LM Studio, Unsloth Desktop, agent tools.
- MiniMax Community License; temp 1.0, top_p 0.95, top_k 40 recommended.
- Hub reports 426B params, 6,196 downloads/month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | MiniMax-M3-GGUF |
| Base model | MiniMaxAI/MiniMax-M3 |
| Architecture | MoE + MiniMax Sparse Attention (MSA) |
| Total params | ~428B (hub: 426B) |
| Activated params | ~23B |
| Experts | 128 (4 active per token) |
| Layers | 60 |
| Context length | 1M tokens |
| Modalities | Text, Image, Video |
| Precision | bfloat16 |
| Quant rungs | UD-IQ1_M (128 GB) to BF16 (852 GB) |
| License | MiniMax Community License |
| llama.cpp status | Experimental (PR #24523) |
| Sampling | temperature=1.0, top_p=0.95, top_k=40 |
| Downloads/month | 6,196 |

## Why this source matters for the RAG

This card documents the practical path to running a 428B multimodal MoE locally via GGUF quantization, including exact rung sizes, experimental llama.cpp status, and sparse-attention fallback behavior. It is valuable reference for retrieval on GGUF quantization, llama.cpp compatibility, and local deployment of frontier-scale MoEs.
