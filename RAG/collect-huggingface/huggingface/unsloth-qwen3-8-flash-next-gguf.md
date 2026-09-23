---
id: collect-huggingface/huggingface/unsloth-qwen3-8-flash-next-gguf
title: "Qwen3.8-Flash-Next-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "qwen", "agent", "agents", "attention", "benchmark", "benchmarks", "embedding", "inference", "license", "llama", "llama.cpp"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-Qwen3.8-Flash-Next-GGUF.md
source_anchor: ""
source_lines: [1, 57]
sha256: 4e58ce5fe34e238a7ce82dab3223de2db7d403d5726c48a9a2ded310b0b91d64
---

# Qwen3.8-Flash-Next-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/Qwen3.8-Flash-Next-GGUF
- **Site** : Hugging Face
- **Type** : Model card (quantized GGUF)
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth AI's GGUF quantization of Qwen/Qwen3.8-Flash-Next, the experimental preview of the architecture that will underpin Qwen4 (hybrid attention with Qwen Sparse Attention, Gated Residual, n-gram embedding, MoE with 512 experts / 10 routed + 1 shared). The repo is tagged Image-Text-to-Text, GGUF, unsloth, conversational; licensed qwen-community-1.0; hub size 177B params, architecture `qwen4exp`; 1,735,774 downloads/month. Unsloth Dynamic 3.0 is used, with MTP now available for 1.3–1.7x faster inference in Unsloth.

Quantization variants: 1-bit UD-IQ1_S 72.5GB / UD-IQ1_M 74.5GB; 2-bit UD-Q2_K_XL 78.9GB; 3-bit UD-IQ3_XXS 82GB / UD-Q3_K_XL 90GB; 4-bit UD-IQ4_XS 93.7GB / MTP Q4_K_M 2.79GB and 1.91GB / UD-Q4_K_XL 111GB; 5-bit UD-Q5_K_XL 158GB; 6-bit UD-Q6_K_XL 169GB; 8-bit Q8_0 192GB / MTP Q8_0 2.79GB; 16-bit BF16 354GB / MTP BF16 7.77GB and 5.23GB. Separate MTP draft files are provided for accelerated decoding. The recommended run command uses UD-Q4_K_XL.

The card reproduces the full Qwen3.8-Flash-Next model overview (125B with 6B activated, plus 51B n-gram embedding and 4B MTP; 48 layers; hidden 2560; Gated DeltaNet + QSA; 512 MoE experts; Gated Residual 4 branches; 262,144 native context extensible to 1M) and benchmark tables (DeepSWE 1.1 58.7, SWE-bench Pro 62.5, SWE-bench Multilingual 81.0, CoWorkBench 73.9, JobBench 55.7, Agents' Last Exam 51.2, Toolathlon Verified 73.5, GPQA Diamond 91.7, LiveCodeBench v6 91.9; VL: ClawEval-MM 64.4, RecreationBench 49.9, AndroidWorld 84.5, OSWorld 2.0, Vision2Web 64.0, ERQA 72.3, LVBench 76.6, RealWorldQA 88.5, MathVision 95.7 with CI, CharXiv 90.6 with CI). API usage notes cover `enable_thinking`/`preserve_thinking`/`reasoning_effort` and sampling params (thinking temp 1.0/top_p 0.95; instruct temp 0.7/top_p 0.80). Best practices include the 262,144/131,072 output split, YaRN scaling for >262K context, and the video preprocessor `longest_edge` setting. Run via llama.cpp, Unsloth Desktop (thinking controls), Ollama, LM Studio, Jan, vLLM, Pi, Docker Model Runner, Lemonade, Hermes Agent, Atomic Chat, OpenClaw.

## Key points

- Unsloth Dynamic 3.0 GGUF of the Qwen4-preview architecture model (architecture `qwen4exp`), qwen-community-1.0.
- Hub size 177B params (125B active-core + 51B n-gram embedding + MTP); 6B activated.
- Variants from 1-bit UD-IQ1_S (72.5GB) to 8-bit Q8_0 (192GB) and BF16 (354GB).
- Dedicated MTP draft files (Q4_K_M/Q8_0/BF16) for 1.3–1.7x faster inference.
- Reproduces full benchmark suite (DeepSWE 58.7, SWE Pro 62.5, GPQA-D 91.7, AndroidWorld 84.5).
- Native context 262,144, up to 1M via YaRN.
- Unsloth Desktop thinking controls; supports llama.cpp, Ollama, LM Studio, Jan, vLLM.
- 1.74M downloads/month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | Qwen3.8-Flash-Next-GGUF |
| Type | GGUF quantizations (Unsloth Dynamic 3.0) |
| Base model | Qwen/Qwen3.8-Flash-Next |
| Hub size | 177B params |
| Architecture | qwen4exp |
| License | qwen-community-1.0 |
| 1-bit | UD-IQ1_S 72.5GB; UD-IQ1_M 74.5GB |
| 2-bit | UD-Q2_K_XL 78.9GB |
| 3-bit | UD-IQ3_XXS 82GB; UD-Q3_K_XL 90GB |
| 4-bit | UD-IQ4_XS 93.7GB; UD-Q4_K_XL 111GB; MTP Q4_K_M 2.79/1.91GB |
| 5-bit | UD-Q5_K_XL 158GB |
| 6-bit | UD-Q6_K_XL 169GB |
| 8-bit | Q8_0 192GB; MTP Q8_0 2.79GB |
| 16-bit | BF16 354GB; MTP BF16 7.77/5.23GB |
| Context | 262,144 native; up to 1M (YaRN) |
| Key benchmarks | DeepSWE 58.7; SWE Pro 62.5; GPQA-D 91.7; AndroidWorld 84.5; RealWorldQA 88.5 |
| Serving | llama.cpp, Unsloth Desktop, Ollama, LM Studio, Jan, vLLM, etc. |
| Downloads/month | 1,735,774 |

## Why this source matters for the RAG

This card is the reference for running the Qwen4-preview architecture locally, giving the complete quantization and MTP-acceleration picture for a 177B MoE multimodal model. It is valuable for retrieval on local deployment of large MoE models, GGUF quantization, and sparse-attention architecture performance.
