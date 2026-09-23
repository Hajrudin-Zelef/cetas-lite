---
id: vague2-datacamp/datacamp/llama-4
title: "Llama 4 de Meta : fonctionnalités, accès, fonctionnement et plus encore"
domain: datacamp
role: reference
task: article
actors: ["Alibaba", "DeepSeek", "Google", "Hugging Face", "Meta", "Microsoft", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["llama", "benchmark", "benchmarks", "consumer", "context window", "deepseek", "distillation", "fine-tuning", "gemini", "gpu", "inference", "license"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/llama-4.md
source_anchor: ""
source_lines: [1, 61]
sha256: 08f784d6329c9f378405ee63996fe1bd297ce31d7a7e1e5eaf4fb515ffafc5d1
---

# Llama 4 de Meta : fonctionnalités, accès, fonctionnement et plus encore

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/llama-4
- **Site** : DataCamp
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article introduces Meta's Llama 4 family, comprising two released models—Llama 4 Scout and Llama 4 Maverick—and one still in training, Llama 4 Behemoth. Scout and Maverick are open-weight under Meta's usual license, with a notable restriction: services exceeding 700 million monthly active users require a separate license from Meta, which may or may not be granted. Scout supports a 10-million-token context window, the largest of any publicly released model. Maverick is a generalist targeting GPT-4o, Gemini 2.0 Flash, and DeepSeek-V3. Behemoth serves as a high-capacity teacher model.

Llama 4 introduces a mixture-of-experts (MoE) architecture, activating only the components needed for a given task. It arrives in a far more competitive open-weight field than Llama 2/3 faced, with DeepSeek strong on reasoning, Alibaba's Qwen on multilingual/code, Google's Gemma on efficiency, and OpenAI announcing an open-weight model.

**Llama 4 Scout**: lightest of the suite, runs on a single H100 GPU, 10M-token context. It has 17B active parameters across 16 experts (~109B total), pre-/post-trained with a 256K context but generalizing far beyond (to be verified). MoE activates only a subset of parameters per token vs dense models. Multimodal (text, image, video pre-training with early fusion); outperforms previous Llama models on image-rich tasks (visual grounding, VQA).

**Llama 4 Maverick**: generalist, 17B active parameters, larger MoE (128 experts, ~400B total), runs on a single H100 DGX or distributed inference. Post-training mixed lightweight supervised fine-tuning, online RL, and direct preference optimization, dropping >50% of "easy" training examples. Co-distilled from Behemoth.

**Llama 4 Behemoth**: most powerful/largest Meta model, still training, not a reasoning model like DeepSeek-R1 or o3, and not a direct-use product—rather a teacher for distilling Scout and Maverick. 288B active parameters across 16 experts, ~2 trillion total. Required new training infrastructure: asynchronous RL, difficulty-based curriculum sampling, and a new distillation loss balancing soft/hard targets. Post-training dropped >95% of SFT examples.

**Benchmarks** (Meta self-reported): Scout scores 88.8 ChartQA, 94.4 DocVQA, 69.4 MMMU, 70.7 MathVista, 32.8 LiveCodeBench, 74.3 MMLU Pro, 57.2 GPQA Diamond, and strong MTOB long-context results. Maverick scores 73.4 MMMU, 73.7 MathVista, 90.0 ChartQA, 94.4 DocVQA, 43.4 LiveCodeBench, 80.5 MMLU Pro, 69.8 GPQA Diamond, 84.6 Multilingual MMLU. Behemoth scores 95.0 MATH-500, 82.2 MMLU Pro, 73.7 GPQA Diamond, 85.8 Multilingual MMLU, 76.1 MMMU, 49.4 LiveCodeBench. Access: download from Meta's Llama site or Hugging Face; use via Meta AI on WhatsApp, Messenger, Instagram, Facebook (Meta account required, no standalone API endpoint). Fine-tuning and local running are supported (Scout needs a high-end GPU; quantized versions may run on consumer hardware).

## Key points

- Llama 4 includes Scout and Maverick (released, open-weight) and Behemoth (in training).
- Open-weight license restricts services with >700M monthly active users (separate Meta license needed).
- Scout: single H100, 10M-token context (largest public), 17B active/~109B total, 16 experts, multimodal.
- Maverick: generalist, 17B active/~400B total, 128 experts, co-distilled from Behemoth.
- Behemoth: ~2T total params, 288B active, teacher model, new async RL infrastructure.
- Uses MoE architecture vs prior dense Llama models; adds native multimodal training.
- No official Meta API; third-party providers may offer access; fine-tuning supported.

## Technical data / figures

| Model | Active params | Total params | Experts | Context | Status |
| --- | --- | --- | --- | --- | --- |
| Llama 4 Scout | 17B | ~109B | 16 | 10M tokens | Released |
| Llama 4 Maverick | 17B | ~400B | 128 | (not stated) | Released |
| Llama 4 Behemoth | 288B | ~2T | 16 | (not stated) | In training |

| Benchmark | Scout | Maverick | Behemoth |
| --- | --- | --- | --- |
| MMMU | 69.4 | 73.4 | 76.1 |
| MathVista | 70.7 | 73.7 | — |
| ChartQA | 88.8 | 90.0 | — |
| DocVQA (test) | 94.4 | 94.4 | — |
| LiveCodeBench | 32.8 | 43.4 | 49.4 |
| MMLU Pro | 74.3 | 80.5 | 82.2 |
| GPQA Diamond | 57.2 | 69.8 | 73.7 |
| MATH-500 | — | — | 95.0 |
| Multilingual MMLU | — | 84.6 | 85.8 |

- License threshold: 700 million monthly active users.
- Scout hardware: single H100 GPU.

## Why this source matters for the RAG

It provides a detailed overview of Meta's Llama 4 family, including architecture, parameter counts, benchmark scores, licensing, and access, useful for open-weight LLM questions. It also situates Llama 4 within the competitive open-weight landscape.
