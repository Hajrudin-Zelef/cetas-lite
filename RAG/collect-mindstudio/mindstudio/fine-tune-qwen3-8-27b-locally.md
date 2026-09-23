---
id: collect-mindstudio/mindstudio/fine-tune-qwen3-8-27b-locally
title: "How to Fine-Tune Qwen3 27B Locally: LoRA, QLoRA and GGUF Guide"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Hugging Face", "Unsloth"]
dates: ["2026-09-23"]
keywords: ["gguf", "lora", "qlora", "consumer", "fine-tuning", "gpu", "gpus", "memory", "parameters", "quantization", "training"]
source: docs/RAG/Collect RAG/02_mindstudio/fine-tune-qwen3-8-27b-locally.md
source_anchor: ""
source_lines: [1, 55]
sha256: a8c2b01fc47e35de5995589f74bfabe22c7ab500f46c557f3d48528b46ec663d
---

# How to Fine-Tune Qwen3 27B Locally: LoRA, QLoRA and GGUF Guide

## Metadata

- **Source** : https://www.mindstudio.ai/blog/fine-tune-qwen3-8-27b-locally
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This guide explains how to fine-tune **Qwen3 27B** on a single GPU using **Unsloth**, **LoRA/QLoRA**, and export to **GGUF**, including dataset creation steps. Fine-tuning a 27B-parameter model on one GPU is possible because of two techniques working together. **LoRA** freezes the original model and trains small adapter layers instead of the full weight set. **QLoRA** adds a compression step, shrinking the frozen model to **4-bit precision** before attaching those adapters. Combined, this cuts VRAM use enough that a demonstrated run consumed roughly **19 GB of VRAM on a 48 GB card**, meaning it likely fits on smaller commodity GPUs too.

A fine-tuning dataset is a collection of examples teaching the model how to respond to specific prompts. The common format is **JSONL**, where each line is a separate JSON object typically containing an **instruction** (the prompt), an optional **input**, and an **output** (the desired answer). Building one manually is straightforward: write a line for each fact or behavior. There's no strict minimum — a demonstrated example used only a few dozen entries to noticeably shift answers on a specific topic. For production-grade results, aim for several hundred to a few thousand examples. Raw notes or documents can be fed to a local LLM to reformat into instruction/output pairs, followed by manual review.

Inside the model, LoRA freezes every original weight and inserts small adapter layers at specific points. Only the adapters (often just a few million parameters vs billions) get updated. Two settings control behavior: **rank** (commonly around 16) determines adapter size — smaller values keep training fast and light — and **alpha** is a scaling factor determining how strongly the adapter's changes influence output. Adapters attach to the transformer's **query, key, and value projections**, the **output projection**, and the **feed-forward layers**, letting small cheap updates reach every important part of the model.

QLoRA adds compression before adapters attach: instead of keeping the frozen base model in full precision, it shrinks it to **4-bit**, cutting memory footprint to roughly a quarter. The compressed model stays frozen in this compact form while the adapters train in higher precision. This is what makes fine-tuning a 27B model feasible on a single consumer/prosumer GPU rather than a multi-GPU cluster.

Training runs through **supervised fine-tuning (SFT)** on labeled question-answer pairs. Key settings: **batch size and gradient accumulation** (a small batch with accumulation simulates a larger batch without more memory), **warm-up steps** (gradually ramps the learning rate), **epochs** (three passes were enough for a small demo dataset), and **learning rate and optimizer** (a memory-efficient optimizer applies updates without unnecessary overhead). During training, **loss** should trend downward and **gradient norm** provides a stability check.

After training, LoRA adapter weights are saved separately from the base model. Exporting to **GGUF** with quantization such as **Q4KM** packages everything so tools like **Ollama** can run it locally without the training environment. The clearest validation is a direct comparison: ask the base model a question tied to the training data (typically no knowledge), then ask the same question with the trained adapter attached (a noticeably different, more accurate answer). The tradeoff is that dataset quality and size directly determine usefulness. Unsloth is one option; the same approach works with Hugging Face **TRL with PEFT**, **Axolotl**, or **torchtune**.

## Key points

- Unsloth + LoRA/QLoRA makes fine-tuning Qwen3 27B light on VRAM (demo: ~19 GB on a 48 GB card).
- LoRA freezes base weights and trains small adapters touching <1% of parameters.
- QLoRA compresses the frozen base to 4-bit, roughly quartering memory and enabling single-GPU training.
- Dataset format matters more than size: JSONL with instruction/input/output; a few dozen examples for a demo, hundreds to thousands for production.
- Adapters attach to query/key/value/output projections and feed-forward layers; rank ~16 and alpha control behavior.
- Export to GGUF (Q4KM) to run via Ollama; validate with a base-vs-fine-tuned comparison.
- Alternatives to Unsloth: HF TRL + PEFT, Axolotl, torchtune.

## Technical data / figures

| Item | Value |
|---|---|
| Base model | Qwen3 27B |
| Demo VRAM usage | ~19 GB on a 48 GB GPU |
| QLoRA precision | 4-bit |
| Dataset format | JSONL (instruction / input / output) |
| Demo dataset size | A few dozen entries |
| Production dataset size | Hundreds to thousands |
| LoRA rank | ~16 (typical) |
| Adapter targets | Q, K, V, output projections, feed-forward |
| Epochs (demo) | 3 |
| Export format | GGUF (Q4KM) |
| Runtime after export | Ollama |
| Alternative tools | TRL + PEFT, Axolotl, torchtune |

## Why this source matters for the RAG

It is a practical, end-to-end reference on local fine-tuning of a widely used open model, covering dataset construction, LoRA/QLoRA mechanics, training settings, and GGUF export. It supports questions about private, single-GPU fine-tuning without API costs.
