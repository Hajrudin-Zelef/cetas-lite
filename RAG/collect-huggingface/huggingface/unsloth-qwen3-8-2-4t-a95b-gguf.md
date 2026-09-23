---
id: collect-huggingface/huggingface/unsloth-qwen3-8-2-4t-a95b-gguf
title: "Qwen3.8-2.4T-A95B-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "qwen", "agent", "agents", "attention", "benchmark", "benchmarks", "license", "llama", "llama.cpp", "memory", "moe"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-Qwen3.8-2.4T-A95B-GGUF.md
source_anchor: ""
source_lines: [1, 56]
sha256: 5fbc52662200361685d49dd54e22f7db5feac1c99370217bf2477f8a27602d2c
---

# Qwen3.8-2.4T-A95B-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/Qwen3.8-2.4T-A95B-GGUF
- **Site** : Hugging Face
- **Type** : Model card (quantized GGUF)
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth AI's GGUF quantization of Qwen/Qwen3.8-2.4T-A95B, the first Qwen-Max-class open release — a text-only Causal Language Model with 2.4T total / 95B activated parameters, licensed under qwen3.8-max, architecture `qwen35moe`, 2.4T params on the hub, and 11,107 monthly downloads. Unsloth Dynamic 2.0 quants are used, and the card notes new sub-1-bit IQ1 data-types (Q1_0 / IQ1_XXXXS) for Qwen 3.8, and that Qwen3.8 runs in Unsloth Desktop with thinking toggles.

Quantization sizes are extremely large, reflecting the 2.4T-parameter scale: 1-bit UD-IQ1_S 508GB / UD-Q1_0 397GB / UD-IQ1_M 564GB; 2-bit UD-IQ2_XXS 657GB / UD-IQ2_XS 731GB; 3-bit UD-IQ3_XXS 956GB; 4-bit UD-IQ4_XS 1.31TB; 8-bit Q8_0 2.6TB; 16-bit BF16 4.89TB. The default run examples use the BF16 tag (e.g., `llama serve -hf unsloth/Qwen3.8-2.4T-A95B-GGUF:BF16`).

The card reproduces the full Qwen3.8-2.4T-A95B model overview (92 layers, hidden 8192, 248,320 padded vocab, Gated DeltaNet + Gated Attention hybrid layout, 512 MoE experts with 10 routed + 1 shared, MTP, 262,144 native context extensible to 1,010,000) and the complete Qwen3.8-Max benchmark table (Terminal Bench 2.1 86.6, SWE-bench Pro 67.7, DeepSWE 1.1 56.6, FrontierSWE 73.5, PaperBench 93.0, AndroidBench 75.1, QwenSWEBench 80.7, QwenReactBench 1724, CoWorkBench 74.8, JobBench 53.4, SkillsBench 70.2, Agents' Last Exam 27.0/52.4, Toolathlon Verified 72.5, WideSearch 81.9, GPQA Diamond 92.6, HLE 43.6, IFBench 82.8, $OneMillion-Bench 52.5, HealthBench 60.2, PLawBench 73.2, PRBench-Legal/Finance 57.6/58.3, MRCR v2 256K 92.9, LongBench v2 66.3). API notes: text-only, thinking mode required (cannot be disabled), `reasoning_effort` xhigh/medium/low, `preserve_thinking` on by default, recommended sampling temp 1.0 / top_p 0.95 / top_k 20, and output-split best practice (262,144 reasoning / 131,072 final within 1M context). Serving via llama.cpp, Ollama, LM Studio, Jan, vLLM, SGLang, Unsloth Desktop, Pi, Docker Model Runner, Lemonade, Hermes Agent, Atomic Chat, OpenClaw; Transformers loading also shown.

## Key points

- Unsloth Dynamic 2.0 GGUF of the 2.4T-param Qwen3.8-Max-class model (architecture qwen35moe), qwen3.8-max license.
- Text-only, thinking-mode-required model; 95B activated.
- Enormous quants: UD-Q1_0 397GB up to Q8_0 2.6TB and BF16 4.89TB.
- New sub-1-bit IQ1 data-types (Q1_0 / IQ1_XXXXS) introduced for Qwen 3.8.
- Reproduces full Qwen3.8-Max benchmarks (Terminal Bench 2.1 86.6, SWE Pro 67.7, PaperBench 93.0, GPQA-D 92.6).
- Default run example uses BF16 tag; runs in llama.cpp, Ollama, vLLM, SGLang, Unsloth Desktop, etc.
- Recommended sampling temp 1.0, top_p 0.95, top_k 20; output split 262,144/131,072.
- 11,107 downloads/month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | Qwen3.8-2.4T-A95B-GGUF |
| Type | GGUF quantizations (Unsloth Dynamic 2.0) |
| Base model | Qwen/Qwen3.8-2.4T-A95B |
| Total params | 2.4T |
| Activated params | 95B |
| Architecture | qwen35moe |
| License | qwen3.8-max |
| 1-bit | UD-Q1_0 397GB; UD-IQ1_S 508GB; UD-IQ1_M 564GB |
| 2-bit | UD-IQ2_XXS 657GB; UD-IQ2_XS 731GB |
| 3-bit | UD-IQ3_XXS 956GB |
| 4-bit | UD-IQ4_XS 1.31TB |
| 8-bit | Q8_0 2.6TB |
| 16-bit | BF16 4.89TB |
| Context | 262,144 native; up to 1,010,000 |
| Key benchmarks | Terminal Bench 2.1 86.6; SWE Pro 67.7; PaperBench 93.0; GPQA-D 92.6; MRCR v2 92.9 |
| Serving | llama.cpp, Ollama, vLLM, SGLang, Unsloth Desktop, etc. |
| Downloads/month | 11,107 |

## Why this source matters for the RAG

This card maps the practical quantization/sizing space for running the largest open Qwen model (2.4T params) locally, including new sub-1-bit formats, and reproduces its full benchmark suite. It is essential for retrieval on extreme-scale MoE deployment, memory/hardware planning, and frontier open-model capabilities.
