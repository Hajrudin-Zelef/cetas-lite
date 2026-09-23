---
id: collect-huggingface/huggingface/unsloth-kimi-k3-gguf
title: "Kimi K3 GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Moonshot", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "kimi", "agent", "agentic", "attention", "license", "llama", "llama.cpp", "moe", "multimodal", "mxfp4", "open-weight"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-Kimi-K3-GGUF.md
source_anchor: ""
source_lines: [1, 46]
sha256: 2782222489c8b0c09d7417fd9b650ce715a31f4a2b0c89e6f1c0f5bdbf58044e
---

# Kimi K3 GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/Kimi-K3-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

unsloth/Kimi-K3-GGUF is Unsloth's GGUF conversion of Moonshot AI's Kimi K3, the 2.8T-parameter (104B active) open-weight multimodal agentic model, using Unsloth Dynamic 2.0 quantization. The card embeds the full Kimi K3 model card, so the architecture is unchanged: MoE built on Kimi Delta Attention (KDA) and Attention Residuals (AttnRes) with a Stable LatentMoE framework activating 16 of 896 experts; 93 layers (1 dense + 69 KDA + 24 Gated MLA); 1M-token context; native text/image/video understanding via the MoonViT-V2 vision encoder (401M); MXFP4 weights / MXFP8 activations; vocabulary 160K; Kimi K3 License. The GGUF release adds vision support and thinking toggles for High and Max in Unsloth Studio.

The key added content is the quantization guidance. To run Kimi K3 in full-precision lossless, users should use Q8 (UD-Q8_K_XL), which is ~1.56 TB and only ~50 GB bigger than the Q4 (UD-Q4_K_XL at 1.51 TB). The hardware-compatibility table lists the full Dynamic 2.0 quant lineup: 1-bit UD-IQ1_S 594 GB, UD-Q1_0 466 GB, UD-TQ1_0 509 GB, UD-IQ1_M 649 GB; 2-bit UD-IQ2_XXS 711 GB, UD-TQ2_0 551 GB, UD-Q2_K_XL 861 GB; 4-bit UD-Q4_K_XL 1.51 TB; 8-bit UD-Q8_K_XL 1.56 TB. Downloads are ~457K/month, with a GGUF `kimi-k3` architecture tag and a chat template.

Usage: run with Unsloth's llama.cpp PR fork or in Unsloth Studio; llama.cpp commands are provided (e.g., `llama serve -hf unsloth/Kimi-K3-GGUF:UD-Q4_K_XL`), plus Ollama, LM Studio, Jan, vLLM, SGLang, Pi, OpenClaw, Hermes, and Docker Model Runner. As with the base model, thinking is always enabled with `reasoning_effort` ("low"/"high"/"max"), preserved-thinking mode requires passing the full assistant message (including `reasoning_content` and `tool_calls`) back, and the recommended agent framework is Kimi Code CLI. A dedicated "How to Run Kimi K3" guide is linked (unsloth.ai/docs/models/kimi-k3), and the base model is moonshotai/Kimi-K3 (51 quantizations).

## Key points

- Unsloth Dynamic 2.0 GGUF of Kimi K3 (2.8T-A104B, multimodal, 1M context).
- Full-precision lossless = UD-Q8_K_XL (~1.56 TB), ~50 GB bigger than Q4.
- Quant lineup from 466 GB (Q1_0) to 1.56 TB (Q8_K_XL).
- Vision supported; High/Max thinking toggles in Unsloth Studio.
- Kimi K3 License; runs via llama.cpp PR fork / Unsloth Studio / Ollama / vLLM / SGLang.
- ~457K downloads/month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | Kimi-K3-GGUF |
| Base model | moonshotai/Kimi-K3 |
| Architecture | kimi-k3 (MoE: 69 KDA + 24 Gated MLA, 896 experts, 16 active) |
| Total / active params | 2.8T / 104B |
| Context length | 1M tokens |
| Modalities | Text, Image, Video (MoonViT-V2) |
| License | kimi-k3 |
| Quantizations | UD-IQ1_S 594 GB, UD-Q1_0 466 GB, UD-TQ1_0 509 GB, UD-IQ1_M 649 GB, UD-IQ2_XXS 711 GB, UD-TQ2_0 551 GB, UD-Q2_K_XL 861 GB, UD-Q4_K_XL 1.51 TB, UD-Q8_K_XL 1.56 TB |
| Downloads/month | ~457,281 |

## Why this source matters for the RAG

This card is the practical companion to the Kimi K3 release, documenting exact GGUF file sizes across the Dynamic 2.0 quant range and lossless-Q8 guidance for a 3T-class model, plus local-serving commands and thinking toggles. It provides essential, citable data on quantization overhead and local deployment of frontier multimodal MoEs.
