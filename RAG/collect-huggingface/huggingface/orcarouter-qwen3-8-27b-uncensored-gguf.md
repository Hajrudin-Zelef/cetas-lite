---
id: collect-huggingface/huggingface/orcarouter-qwen3-8-27b-uncensored-gguf
title: "Qwen3.8-27B-Uncensored-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["AMD", "Alibaba", "China", "Hugging Face", "OpenAI", "Perplexity", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "qwen", "apache", "attention", "fp8", "gpu", "license", "llama", "llama.cpp", "parameters", "perplexity", "quantization"]
source: docs/RAG/Collect RAG/03_huggingface/orcarouter-Qwen3.8-27B-Uncensored-GGUF.md
source_anchor: ""
source_lines: [1, 53]
sha256: f9729969cdb43a2ae8107e9dfa4f8ab83bc3b1d4f3cc2cafb9837b97f41bcff7
---

# Qwen3.8-27B-Uncensored-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/orcarouter/Qwen3.8-27B-Uncensored-GGUF
- **Site** : Hugging Face
- **Type** : Model card (uncensored/abliterated GGUF)
- **Language** : en
- **Verification status** : ✅ reachable (gated repository; card content public)
- **Collection date** : 2026-09-23

## Full summary

This is OrcaRouter's GGUF conversion (2-bit → 16-bit) of Qwen3.8-27B-Uncensored, an abliterated (refusal-removed) build of Qwen's Qwen3.8-27B — a 27B dense hybrid-attention (Gated DeltaNet linear + full attention) native vision-language model with reasoning, tool-calling, and an MTP (`nextn`) speculative-decoding head. Licensed Apache-2.0, architecture `qwen35`, 27B params, with a separate `mmproj` file (0.9 GB) that restores vision. Downloads: 230,871/month. The repository is gated (must accept conditions). GDN linear-attention layers are stored as SSM-style tensors (`ssm_*`), full-attention layers as `attn_*`, and the MTP head as block `nextn.*`; a recent llama.cpp build (post-2026-05, with the `qwen35` hybrid-GDN architecture and MTP/nextn head) is required.

Files: standard K-quants Q2_K (10.9GB) to Q8_0 (27.1GB) and F16 (54.7GB, split), with Q4_K_M (16.8GB) recommended as the default; IQ quants built with an importance matrix on English + Chinese calibration text — IQ4_XS 15.3GB (best low-bit pick), IQ3_M 12.8GB, IQ3_XXS 11.6GB, IQ2_M 10.5GB, IQ2_XXS 8.9GB. All quants preserve the MTP head and GDN architecture; vision needs the separate mmproj.

Usage: llama.cpp (`--jinja -c 8192`), OpenAI-compatible server with `--mmproj` for vision, `--jinja` for Qwen tool template, thinking on by default (toggle via `chat_template_kwargs.enable_thinking`, reasoning returned in `reasoning_content`, give ≥2048 max_tokens), and MTP speculative decoding enabled via llama.cpp's MTP/speculative path. The same quants are published as 16 Ollama tags (q2_K to q8_0) with mmproj bundled. Evaluation (vs official FP8 base, vLLM, rule-based refusal classifier): harmful-prompt refusal collapses from 64–99% to 0–6% (e.g., AdvBench 99%→0%, HarmBench 98.7%→2.7%, thinking-off; and AdvBench 66.7%→1.7%, HarmBench 46.7%→0% thinking-on); benign over-refusal drops (XSTest-safe 5.6%→0.4% no-think); capability stays within ±1.3 pts (MMLU +0.4, MMLU-Pro −0.8, GSM8K −1.3, CMMLU −0.6), WikiText-2 perplexity 6.96. Runs on CPU/CUDA/Metal/ROCm via llama.cpp; Q4_K_M fits a 24GB GPU.

## Key points

- Abliterated GGUF quants (2-bit→16-bit) of Qwen3.8-27B; Apache-2.0; gated download.
- 27B dense hybrid attention (GDN + full attention), reasoning, tool-calling, MTP head; `qwen35`.
- Standard K-quants plus IQ quants with English+Chinese importance matrix; Q4_K_M recommended default.
- Separate mmproj (0.9GB) restores vision; all quants preserve MTP/nextn.
- Harmful-prompt refusal 64–99% → 0–6%; benign over-refusal 5.6%→0.4%; capability within ±1.3 pts.
- Runs via llama.cpp (CPU/CUDA/Metal/ROCm), Ollama (16 tags), LM Studio, Jan, vLLM.
- Reasoning/thinking on by default; tool calling via --jinja; vision via --mmproj.
- 230,871 downloads/month; research-only disclaimer.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | orcarouter |
| Model name | Qwen3.8-27B-Uncensored-GGUF |
| Base model | Qwen/Qwen3.8-27B (abliterated) |
| Architecture | qwen35 (hybrid GDN + attention, MTP head) |
| Parameters | 27B dense |
| License | Apache-2.0 (gated) |
| K-quants | Q2_K 10.9GB → Q8_0 27.1GB; F16 54.7GB; Q4_K_M 16.8GB recommended |
| IQ quants | IQ2_XXS 8.9GB → IQ4_XS 15.3GB |
| Vision | mmproj f16 0.9GB |
| Refusal (off) | AdvBench 99%→0%; HarmBench 98.7%→2.7%; MaliciousInstruct 99%→0% |
| Refusal (on) | AdvBench 66.7%→1.7%; HarmBench 46.7%→0% |
| Capability | MMLU +0.4; GSM8K −1.3; CMMLU −0.6 (within ±1.3 pts) |
| Perplexity | WikiText-2 6.96 |
| Serving | llama.cpp, Ollama (16 tags), vLLM, LM Studio, Jan |
| Downloads/month | 230,871 |

## Why this source matters for the RAG

This card provides detailed, quantified evidence on the effects of abliteration plus quantization of a modern reasoning/vision model, including per-quant file sizing and measured safety/capability trade-offs. It is valuable for retrieval on refusal-mechanism research, GGUF quantization quality, and uncensored-model red-teaming.
