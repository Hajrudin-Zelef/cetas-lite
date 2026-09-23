---
id: collect-huggingface/huggingface/unsloth-qwen3-coder-30b-a3b-instruct-gguf
title: "Qwen3-Coder-30B-A3B-Instruct-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "OpenAI", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "qwen", "agent", "agentic", "agents", "apache", "cost", "gqa", "license", "llama", "llama.cpp", "moe"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-Qwen3-Coder-30B-A3B-Instruct-GGUF.md
source_anchor: ""
source_lines: [1, 58]
sha256: cf0acff9566a03e5b4f646ff37c2146c6adb33cb3e7acb688b4a0239436e1cb8
---

# Qwen3-Coder-30B-A3B-Instruct-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/Qwen3-Coder-30B-A3B-Instruct-GGUF
- **Site** : Hugging Face
- **Type** : Model card (quantized GGUF)
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth AI's GGUF quantization of Qwen3-Coder-30B-A3B-Instruct, a streamlined agentic coding MoE from the Qwen team. It is a Text Generation model, tagged GGUF/unsloth/qwen3/qwen/imatrix/conversational, licensed Apache-2.0, with 31B params on the Hub, architecture `qwen3moe`, and 12,450,427 monthly downloads — one of the most-downloaded coding GGUF repos. Qwen3-Coder-30B-A3B-Instruct has 30.5B total / 3.3B activated parameters, 48 layers, GQA (32 Q / 4 KV heads), 128 experts with 8 activated, and 262,144-token native context (extendable to 1M with YaRN). It supports only non-thinking mode and does not generate ` thinking\n response` blocks.

Unsloth Dynamic 2.0 quants are offered: 1-bit UD-IQ1_S 8.91GB / UD-TQ1_0 8.01GB / UD-IQ1_M 9.63GB; 2-bit UD-IQ2_XXS 10.3GB / Q2_K 11.3GB / UD-IQ2_M 10.8GB / Q2_K_L 11.3GB / UD-Q2_K_XL 11.8GB; 3-bit UD-IQ3_XXS 12.8GB / Q3_K_S 13.3GB / Q3_K_M 14.7GB / UD-Q3_K_XL 13.8GB; 4-bit IQ4_XS 16.4GB / Q4_K_S 17.5GB / IQ4_NL 17.3GB / Q4_0 17.4GB / Q4_1 19.2GB / Q4_K_M 18.6GB / UD-Q4_K_XL 17.7GB; 5-bit Q5_K_S 21.1GB / Q5_K_M 21.7GB / UD-Q5_K_XL 21.7GB; 6-bit Q6_K 25.1GB / UD-Q6_K_XL 26.3GB; 8-bit Q8_0 32.5GB / UD-Q8_K_XL 36GB; 16-bit BF16 61.1GB. Recommended default: UD-Q4_K_XL.

The card reproduces the Qwen3-Coder highlights: significant performance among open models on agentic coding and agentic browser-use, 256K native long context optimized for repository-scale understanding, and a specially designed function-call format supporting Qwen Code, CLINE, and other agent platforms. Quickstart notes that `transformers<4.51.0` fails with `KeyError: 'qwen3_moe'` and shows standard Transformers usage (max_new_tokens 65536; reduce context to 32,768 on OOM). Agentic coding works via an OpenAI-compatible endpoint with tools. Best practices: temperature 0.7, top_p 0.8, top_k 20, repetition_penalty 1.05; recommended output length 65,536 tokens. Citation references the Qwen3 Technical Report (arXiv 2505.09388). Run via llama.cpp, Ollama, LM Studio, Jan, vLLM, SGLang, Unsloth Desktop, Pi, Docker Model Runner, Lemonade, Hermes Agent, Atomic Chat, OpenClaw.

## Key points

- Unsloth Dynamic 2.0 GGUF of Qwen3-Coder-30B-A3B-Instruct; Apache-2.0; architecture qwen3moe.
- MoE: 30.5B total / 3.3B activated; 128 experts, 8 activated; 48 layers.
- 262,144 native context (up to 1M with YaRN); non-thinking mode only.
- Quant range: UD-TQ1_0 8.01GB to Q8_0 32.5GB and BF16 61.1GB; UD-Q4_K_XL recommended.
- Strong agentic coding / browser-use performance; function-call format for Qwen Code, CLINE.
- Requires transformers >= 4.51.0 for Transformers loading; OOM mitigation: context 32,768.
- Best practices: temp 0.7, top_p 0.8, top_k 20, rep penalty 1.05; 65,536-token output.
- 12.45M downloads/month; Qwen3 Technical Report arXiv 2505.09388.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | Qwen3-Coder-30B-A3B-Instruct-GGUF |
| Type | GGUF quantizations (Unsloth Dynamic 2.0) |
| Base model | Qwen/Qwen3-Coder-30B-A3B-Instruct |
| Total / active params | 30.5B / 3.3B (hub 31B) |
| Architecture | qwen3moe |
| Layers | 48 |
| MoE | 128 experts; 8 activated |
| GQA | 32 Q / 4 KV heads |
| Context length | 262,144 native; up to 1M (YaRN) |
| Thinking mode | Not supported (non-thinking only) |
| License | Apache-2.0 |
| 1-bit | UD-TQ1_0 8.01GB; UD-IQ1_S 8.91GB; UD-IQ1_M 9.63GB |
| 2-bit | UD-IQ2_XXS 10.3GB; UD-Q2_K_XL 11.8GB |
| 3-bit | UD-IQ3_XXS 12.8GB; UD-Q3_K_XL 13.8GB |
| 4-bit | IQ4_XS 16.4GB; Q4_K_M 18.6GB; UD-Q4_K_XL 17.7GB |
| 8-bit | Q8_0 32.5GB; UD-Q8_K_XL 36GB |
| 16-bit | BF16 61.1GB |
| Sampling | temp 0.7, top_p 0.8, top_k 20, rep_penalty 1.05 |
| Downloads/month | 12,450,427 |

## Why this source matters for the RAG

This card is the primary local-deployment reference for a leading open agentic coding MoE, covering the full quantization range and serving options at a very cost-efficient 3.3B-active scale. It is essential for retrieval on coding agents, MoE quantizations, and repository-scale code understanding.
