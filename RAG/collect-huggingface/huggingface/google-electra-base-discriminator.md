---
id: collect-huggingface/huggingface/google-electra-base-discriminator
title: "electra-base-discriminator - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Google", "Hugging Face"]
dates: ["2026-09-23"]
keywords: ["apache", "compute", "gpu", "license", "training"]
source: docs/RAG/Collect RAG/03_huggingface/google-electra-base-discriminator.md
source_anchor: ""
source_lines: [1, 46]
sha256: 7dc7eb4ab61272a078ba2bab2c8636a53f374cd604170a6ae6e5835bf216674d
---

# electra-base-discriminator - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/google/electra-base-discriminator
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

electra-base-discriminator is Google's base-size ELECTRA model, a self-supervised text encoder pre-trained with the ELECTRA method described in "ELECTRA: Pre-training Text Encoders as Discriminators Rather Than Generators" (OpenReview). Instead of masking tokens like BERT's MLM, ELECTRA trains a discriminator to distinguish "real" input tokens from "fake" tokens replaced by a small generator network, in a GAN-inspired setup. This makes pre-training sample-efficient: strong results are achievable at small scale even on a single GPU, and at large scale ELECTRA reaches state-of-the-art results on SQuAD 2.0. The repository contains code to pre-train ELECTRA and to fine-tune it on classification (GLUE), QA (SQuAD) and sequence tagging (text chunking) tasks.

This specific checkpoint is the "base" English discriminator, tagged as a Transformers / PyTorch / TensorFlow / JAX / Rust model for pre-training/feature extraction, under the Apache 2.0 license, with ~47.4 million downloads per month. The model card provides a usage example with `ElectraForPreTraining` and `ElectraTokenizerFast`, demonstrating how the discriminator flags replaced tokens: for the sentence "The quick brown fox jumps over the lazy dog" versus the fake "The quick brown fox fake over the lazy dog", the discriminator output is thresholded (`torch.round((torch.sign(outputs) + 1) / 2)`) to identify the replaced "fake" token. This "replaced token detection" objective is the core innovation of ELECTRA and makes it useful as a pretrained encoder backbone for downstream classification and QA tasks.

On the hub, the model has 3 adapters, 87 fine-tunes and 4 quantizations, appears in 32 Spaces (including clinical NLP and hallucination-detection demos), and is part of the Google "ELECTRA release" collection. The referenced paper listed on the page is the original GAN paper (arXiv 1406.2661).

## Key points

- Google's base-size ELECTRA discriminator: GAN-style "replaced token detection" pre-training.
- More compute-efficient than MLM pre-training; SOTA on SQuAD 2.0 at large scale.
- Apache 2.0; ~47.4M downloads/month.
- Usable via ElectraForPreTraining / ElectraTokenizerFast (PyTorch, TF, JAX, Rust).
- Fine-tunes on classification (GLUE), QA (SQuAD), sequence tagging.
- English model; 87 fine-tunes, 4 quantizations, 32 Spaces.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | google |
| Model name | electra-base-discriminator |
| Architecture | ELECTRA base discriminator (Transformer encoder) |
| Pre-training | Replaced-token detection (discriminator vs. generator) |
| License | Apache 2.0 |
| Languages | English |
| Tasks | Feature extraction / pre-training; fine-tunable (GLUE, SQuAD, chunking) |
| Frameworks | Transformers, PyTorch, TF, JAX, Rust |
| Downloads/month | ~47,371,933 |
| Paper | ELECTRA (OpenReview); GAN (arXiv 1406.2661) |

## Why this source matters for the RAG

This card documents the reference ELECTRA discriminator, an efficient alternative to BERT-style MLM encoders that underpins many downstream classification and QA systems. It provides authoritative architecture, licensing, and usage details (including how to run replaced-token detection), supporting accurate retrieval on encoder pre-training methods and model selection for classification/reranking backbones.
