---
id: collect-huggingface/huggingface/unsloth-qwen3-8-27b-gguf
title: "Qwen3.8-27B-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "Unsloth"]
dates: ["2026-09-23"]
keywords: ["gguf", "qwen", "agent", "agentic", "apache", "attention", "license", "llama", "llama.cpp", "multimodal", "parameters", "quantization"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-Qwen3.8-27B-GGUF.md
source_anchor: ""
source_lines: [1, 56]
sha256: af1267fb4e4717178ad31570e2fc86c25a254851f250ba21830d9b49d9caabec
---

# Qwen3.8-27B-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/Qwen3.8-27B-GGUF
- **Site** : Hugging Face
- **Type** : Model card (quantized GGUF)
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth AI's GGUF quantization of Qwen/Qwen3.8-27B, the flagship dense multimodal model of the Qwen3.8 generation (native vision-language, image and video understanding, flexible thinking control). The repo is tagged GGUF, qwen3_5, unsloth, imatrix, conversational, licensed Apache-2.0, with 27B params and architecture `qwen35`, and 7,134,167 monthly downloads — one of the most-downloaded quantization repos. It is part of Unsloth Dynamic 3.0, which the card claims delivers >10% top-1% better accuracy at the same size than other providers.

The full quantization ladder is listed: 1-bit UD-IQ1_S 6.19GB / UD-IQ1_M 6.73GB; 2-bit UD-IQ2_XXS 7.27GB / UD-IQ2_S 8.37GB / UD-Q2_K_XL 9.83GB; 3-bit UD-IQ3_XXS 10.9GB / UD-IQ3_S 12GB / UD-Q3_K_XL 13.1GB; 4-bit UD-IQ4_XS 14.3GB / UD-Q4_K_S 15.4GB / MTP Q4_0 1.37GB / Q4_0 16.1GB / Q4_1 17.5GB / UD-Q4_K_M 16.5GB / UD-Q4_K_XL 17.6GB; 5-bit UD-Q5_K_S 18.7GB / UD-Q5_K_M 19.8GB / UD-Q5_K_XL 20.9GB; 6-bit UD-Q6_K 22GB / UD-Q6_K_M 23.1GB / UD-Q6_K_L 24.2GB / UD-Q6_K_XL 25.3GB; 8-bit Q8_0 29GB / UD-Q8_K_XL 31.5GB; and BF16 54.7GB. The recommended default in run examples is UD-Q4_K_M.

The card reproduces the Qwen3.8-27B overview (64 layers, hidden 5120, Gated DeltaNet + Gated Attention, 262,144 native context extensible to 1M) and best practices: thinking-mode sampling (temp 1.0, top_p 0.95, top_k 20) vs instruct mode (temp 0.7, top_p 0.80, presence_penalty 1.5); recommend 262,144-token reasoning output and 131,072-token final response within 1M context; YaRN RoPE scaling for >262K sequences; and the video preprocessor `longest_edge` recommendation (469,762,048 for ~224k video tokens). Unsloth-specific extras include Developer Role support for agentic tools like Codex, tool-calling improvements for nested objects, and Thinking toggles in Unsloth Desktop. Run options cover llama.cpp (`llama serve -hf unsloth/Qwen3.8-27B-GGUF:UD-Q4_K_M`), LM Studio, Jan, Ollama, Unsloth Desktop, Pi, Docker Model Runner, Lemonade, Hermes Agent, Atomic Chat, and OpenClaw.

## Key points

- Unsloth Dynamic 3.0 GGUF quantization of Qwen3.8-27B (27B dense, architecture qwen35), Apache-2.0.
- Huge quantization range: 1-bit UD-IQ1_S (6.19GB) up to 8-bit UD-Q8_K_XL (31.5GB) and BF16 (54.7GB).
- Includes MTP Q4_0 (1.37GB) draft/embedded MTP file; 7.13M downloads/month.
- Native vision-language, 262,144 context (up to 1M with YaRN).
- Unsloth claims >10% top-1% accuracy advantage over other quants at the same size.
- Developer Role support (Codex etc.) and improved tool-calling in this build.
- Best-practice sampling: thinking temp 1.0/top_p 0.95; instruct temp 0.7/top_p 0.80.
- Runs across llama.cpp, Ollama, LM Studio, Jan, Unsloth Desktop, Lemonade, Hermes, OpenClaw, Pi.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | Qwen3.8-27B-GGUF |
| Type | GGUF quantizations (Unsloth Dynamic 3.0) |
| Base model | Qwen/Qwen3.8-27B |
| Parameters | 27B |
| Architecture | qwen35 |
| License | Apache-2.0 |
| 1-bit | UD-IQ1_S 6.19GB; UD-IQ1_M 6.73GB |
| 2-bit | UD-IQ2_XXS 7.27GB; UD-IQ2_S 8.37GB; UD-Q2_K_XL 9.83GB |
| 3-bit | UD-IQ3_XXS 10.9GB; UD-IQ3_S 12GB; UD-Q3_K_XL 13.1GB |
| 4-bit | UD-IQ4_XS 14.3GB; UD-Q4_K_M 16.5GB; UD-Q4_K_XL 17.6GB; MTP Q4_0 1.37GB |
| 5-bit | UD-Q5_K_S 18.7GB to UD-Q5_K_XL 20.9GB |
| 6-bit | UD-Q6_K 22GB to UD-Q6_K_XL 25.3GB |
| 8-bit | Q8_0 29GB; UD-Q8_K_XL 31.5GB |
| 16-bit | BF16 54.7GB |
| Context | 262,144 native; up to 1M (YaRN) |
| Serving | llama.cpp, Ollama, LM Studio, Jan, Unsloth Desktop, etc. |
| Downloads/month | 7,134,167 |

## Why this source matters for the RAG

This card is the reference for locally running the flagship Qwen3.8-27B, mapping the complete size/quality trade-off space and the MTP acceleration file. It is essential for retrieval on GGUF quantization, local multimodal deployment, and hardware sizing.
