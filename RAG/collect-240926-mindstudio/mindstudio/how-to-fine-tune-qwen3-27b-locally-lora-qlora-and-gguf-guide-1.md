---
id: collect-240926-mindstudio/mindstudio/how-to-fine-tune-qwen3-27b-locally-lora-qlora-and-gguf-guide-1
title: "how-to-fine-tune-qwen3-27b-locally-lora-qlora-and-gguf-guide"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Unsloth"]
dates: []
keywords: ["gguf", "lora", "qlora", "agent", "agents", "attention", "compute", "consumer", "fine-tuning", "gpu", "gpus", "memory"]
source: docs/RAG/clean_en/mindstudio/how-to-fine-tune-qwen3-27b-locally-lora-qlora-and-gguf-guide.md
source_anchor: ""
source_lines: [1, 70]
sha256: 51f092c21892f771403f725756f6d2922e0de4ec46d20cbb179c4b2aba998a35
---

# how-to-fine-tune-qwen3-27b-locally-lora-qlora-and-gguf-guide

<!-- source: https://www.mindstudio.ai/blog/fine-tune-qwen3-8-27b-locally -->

## What does it take to fine-tune Qwen3 27B on one GPU?

Fine-tuning a 27 billion parameter model like Qwen3 on a single GPU is possible because of two techniques working together: LoRA and QLoRA, wrapped inside a library called Unsloth. LoRA freezes the original model and trains small adapter layers instead of the full weight set. QLoRA adds a compression step, shrinking the frozen model down to 4-bit precision before attaching those adapters. Combined, this cuts VRAM use enough that a demonstrated run consumed roughly 19GB of VRAM on a 48GB card, meaning it likely fits on smaller commodity GPUs too.

## TL;DR

- **Unsloth** is the library that makes fine-tuning large models like Qwen3 27B dramatically lighter on VRAM, letting a single GPU handle a job that would otherwise need multiple high-end cards.
- **LoRA adapters** freeze the billions of original model weights and train small inserted layers instead, touching less than 1% of the model’s parameters.
- **QLoRA** compresses the frozen base model to 4-bit precision before training, roughly quartering its memory footprint and making the 27B model fit on one card.
- **Dataset format matters more than size** : a JSONL file with instruction/input/output fields per line is enough to get started, though production work calls for hundreds to thousands of examples.
- **Target modules** like query, key, value, and feed-forward projections inside the transformer’s attention blocks are where LoRA adapters get attached, touching every important part of the model with minimal overhead.
- **GGUF export** with Q4KM quantization lets the fine-tuned model run locally through tools like Ollama, without needing to keep the training stack around.
- **A before-and-after test** is the simplest way to confirm training worked: ask the base model and the fine-tuned model the same question and compare answers.

## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

## What is a dataset for fine-tuning and how do you build one?

A fine-tuning dataset is a collection of examples that teach the model how to respond to specific prompts. The common format is JSONL, a text file where each line is a separate JSON object. A typical structure includes an instruction (the question or prompt), an optional input, and an output (the answer you want the model to learn).

Building one manually is straightforward: write a line for each fact or behavior you want the model to pick up, pairing a question with the answer you’d like it to give. There’s no strict minimum. A demonstrated example used only a few dozen entries to noticeably shift how a model answered questions about a specific topic. For anything beyond a demo or proof of concept, aim much higher, several hundred to a few thousand examples for a production-grade result.

You don’t have to write every line by hand. Feeding raw notes or documents to a local LLM and asking it to reformat that content into instruction/output pairs works well as a starting point, followed by manual review to clean up inconsistencies.

## How does LoRA actually work inside the model?

A large language model like Qwen3 27B stores its learned behavior in billions of internal numbers called weights, set during pretraining. Fully fine-tuning a model means updating all of those weights, which requires serious compute, multiple high-end GPUs and substantial time.

LoRA takes a different approach. It freezes every original weight and instead inserts small adapter layers at specific points inside the model architecture. These adapters are tiny compared to the base model, often just a few million parameters instead of billions. Only the adapters get updated during training. The frozen model provides the foundation, and the adapters learn the specific adjustments needed to shift outputs toward the training data.

Two settings control how these adapters behave. Rank (commonly set around 16) determines the size of the adapter layers, smaller values keep training fast and light. Alpha is a scaling factor that determines how strongly the adapter’s learned changes influence the model’s output.

The adapters attach to specific parts of the transformer’s internals: the query, key, and value projections that drive attention (how the model decides which words matter for a given token), the output projection that recombines attention results, and the feed-forward layers that handle deeper processing after attention. Attaching adapters across all of these lets small, cheap updates reach every meaningfully important part of the model.

## What does QLoRA add on top of LoRA?

QLoRA adds a compression step before the adapters even get attached. Instead of keeping the frozen base model in full precision, QLoRA shrinks it down to 4-bit precision, cutting its memory footprint by roughly a quarter of the original. The compressed model stays frozen in this compact form while the LoRA adapters train on top of it in higher precision.

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

That compression is what makes it feasible to fine-tune a 27 billion parameter model on a single consumer or prosumer GPU rather than a multi-GPU cluster. Without it, even LoRA’s lightweight adapter approach wouldn’t be enough to fit a model this size into typical VRAM budgets.

## What happens during the actual training run?

Once the dataset and LoRA/QLoRA configuration are set, training runs through a supervised fine-tuning process (SFT), where the model learns from labeled question-and-answer pairs. Key settings that shape a run include:

- **Batch size and gradient accumulation**: a small batch size combined with gradient accumulation effectively simulates a larger batch without requiring more memory.
- **Warm-up steps**: gradually ramps up the learning rate at the start of training rather than applying full-strength updates immediately.
- **Epochs**: the number of full passes over the dataset. Three passes was enough for a small demonstration dataset to show a clear shift in the model’s answers.
- **Learning rate and optimizer**: the learning rate controls how large each adjustment step is, while a memory-efficient optimizer applies those updates without adding unnecessary overhead.

Throughout training, loss (a measure of how wrong the model’s predictions are) should trend downward, and gradient norm (the size of each update step) provides a stability check. Watching these two metrics in real time is the simplest way to confirm training is progressing rather than stalling or diverging.

## Why export to GGUF, and how do you test the results?

After training completes, the LoRA adapter weights get saved separately from the base model. To actually use the fine-tuned model outside a training script, exporting to GGUF format (with quantization such as Q4KM) packages everything into a form that tools like Ollama can load and run locally, without needing the original training environment.

The clearest way to validate a fine-tuning run is a direct comparison. Load the original base model and ask it a question tied to your training data, it will typically show no knowledge of that specific content. Then load the same base model with the trained LoRA adapter attached and ask the identical question. A successful fine-tune produces a noticeably different, more accurate answer that reflects the training data, confirming the adapter is doing its job.

## Is fine-tuning Qwen3 27B locally worth it?

