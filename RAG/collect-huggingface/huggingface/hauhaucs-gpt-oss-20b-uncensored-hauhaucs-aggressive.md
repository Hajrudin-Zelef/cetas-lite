---
id: collect-huggingface/huggingface/hauhaucs-gpt-oss-20b-uncensored-hauhaucs-aggressive
title: "GPT-OSS-20B-Uncensored-HauhauCS-Aggressive - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["apache", "benchmark", "benchmarks", "fine-tuning", "gguf", "license", "llama", "llama.cpp", "moe", "mxfp4", "open-weight", "packaging"]
source: docs/RAG/Collect RAG/03_huggingface/HauhauCS-GPT-OSS-20B-Uncensored-HauhauCS-Aggressive.md
source_anchor: ""
source_lines: [1, 44]
sha256: 7d50d3d7bf530d30af42693c714f46e11a49b74ca163a007ea6423876af95d78
---

# GPT-OSS-20B-Uncensored-HauhauCS-Aggressive - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/HauhauCS/GPT-OSS-20B-Uncensored-HauhauCS-Aggressive
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is a community "uncensored" (abliterated) derivative of OpenAI's gpt-oss-20b, produced by HauhauCS. It is the aggressive variant of the uncensored line, tuned harder for fewer refusals; the author also publishes a more conservative "Balanced" variant. The model is distributed as an MXFP4 GGUF, the native 4-bit quantization of the gpt-oss family, at roughly 12.1 GB, and works with llama.cpp, Ollama, LM Studio, Jan, Docker Model Runner and any other GGUF-capable runtime. The base model is the 21B-parameter gpt-oss-20b Mixture-of-Experts (MoE) architecture with 3.6B active parameters, originally released by OpenAI under Apache 2.0 and trained on the harmony response format. The abliteration was performed with custom tooling to reduce refusal behavior, making it suitable for users who want maximum uncensoring. The card notes LM Studio compatibility with Reasoning Effort custom buttons if the model is placed in `LM Models\lmstudio-community\gpt-oss-20b-GGUF\`. There are no benchmark figures published on the card, and the license is listed simply as "other", since the derivative carries its own terms. The repository has 63 likes and 7,273 downloads in the last month. This is a niche, community fine-tune rather than an official release, and users should evaluate safety implications carefully.

## Key points

- Community uncensored (abliterated) version of OpenAI gpt-oss-20b; aggressive variant tuned for fewer refusals.
- Distributed as MXFP4 GGUF, ~12.1 GB, the native 4-bit format of gpt-oss.
- Base architecture: 21B total / 3.6B active MoE.
- Runs on llama.cpp, Ollama, LM Studio, Jan, Docker Model Runner.
- A separate Balanced variant exists for less aggressive uncensoring.
- License listed as "other"; no benchmarks published.
- 7,273 downloads last month; 63 likes.

## Technical data / figures

| Attribute | Value |
|---|---|
| Base model | GPT-OSS 20B (OpenAI gpt-oss-20b) |
| Total / active params | 21B / 3.6B (MoE) |
| Architecture | gpt-oss |
| Format | GGUF (MXFP4) |
| Size | 12.1 GB |
| Language | English |
| Tags | uncensored, abliterated, mxfp4, conversational |
| License | other |
| Runtimes | llama.cpp, Ollama, LM Studio, Jan, Docker Model Runner |
| Downloads last month | 7,273 |
| Likes | 63 |

## Why this source matters for the RAG

This card documents a popular community abliteration workflow and its practical GGUF packaging, illustrating the ecosystem of uncensored derivatives that exists around open-weight models. It is useful context for discussions of model safety, fine-tuning and local deployment, even though it carries no formal benchmarks.
