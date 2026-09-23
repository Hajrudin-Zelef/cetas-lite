---
id: collect-huggingface/huggingface/google-t5-t5-small
title: "T5-small - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Google", "Hugging Face"]
dates: ["2026-09-23"]
keywords: ["apache", "fine-tuning", "license", "tpu", "training"]
source: docs/RAG/Collect RAG/03_huggingface/google-t5-t5-small.md
source_anchor: ""
source_lines: [1, 47]
sha256: 056caf583b41a720f8b1ca31b1ca99a7bc06754c8774498dbb522b6d412324c6
---

# T5-small - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/google-t5/t5-small
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

T5-small is the small checkpoint of the Text-To-Text Transfer Transformer (T5) family from Google, a 60-million-parameter model that reframes all NLP tasks as text-to-text problems: input and output are always text strings, so the same model, loss function and hyperparameters apply to machine translation, summarization, question answering, classification and even regression (by predicting a string representation of a number). The hub reports 60.5M params, Apache 2.0 license, and ~24.6 million downloads/month. It supports English, French, Romanian and German, and is tagged for translation, summarization, and text2text-generation. The model was developed by Colin Raffel et al. (paper: "Exploring the Limits of Transfer Learning with a Unified Text-to-Text Transformer", JMLR 2020) and is documented with a full Hugging Face-written model card.

Pre-training used the Colossal Clean Crawled Corpus (C4) with a multi-task mixture of unsupervised denoising objectives (C4 and Wiki-DPR) and supervised text-to-text tasks (CoLA, SST-2, MRPC, STS-B, QQP, MNLI, QNLI, RTE, CB, COPA, WiC, MultiRC, ReCoRD, BoolQ). The developers evaluated the framework on 24 tasks; full T5-small results are in Table 14 of the paper. Environmental-impact details note training on Google Cloud TPU Pods. Bias/risks sections are marked "more information needed" — the card acknowledges limits of documentation. The model card was authored by the Hugging Face team.

Usage is straightforward: `AutoTokenizer` + `AutoModelForSeq2SeqLM` (or the dedicated `T5Tokenizer`/`T5Model` for a forward pass with decoder input). Note the card warns that the `translation` pipeline type is no longer supported in transformers v5 and users must load the model directly or stay on v4.x. The model has 2,328 fine-tunes, 76 adapters and 20 quantizations downstream, and appears in 100+ Spaces (including MusicGen, YourMT3, translation demos). As a small, efficient seq2seq backbone, it is widely used for summarization and lightweight translation.

## Key points

- Google's T5 small: 60M params, unified text-to-text framework for all NLP tasks.
- Supports English, French, Romanian, German; Apache 2.0.
- Pre-trained on C4 (denoising) + supervised GLUE-style and QA tasks (24-task evaluation).
- Widely fine-tuned (2,328 downstream fine-tunes) for summarization, translation, classification.
- Note: `translation` pipeline deprecated in transformers v5.
- ~24.6M downloads/month; Paper: JMLR 2020.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | google-t5 |
| Model name | t5-small |
| Architecture | Text-to-text Transformer (encoder-decoder) |
| Params | 60.5M |
| Pre-training | C4 + Wiki-DPR (denoising) + supervised text2text mixture |
| Languages | English, French, Romanian, German |
| License | Apache 2.0 |
| Tasks | Translation, summarization, text2text-generation |
| Evaluation | 24 tasks (paper Table 14) |
| Paper | JMLR 21(140), 2020 |
| Downloads/month | ~24,644,038 |

## Why this source matters for the RAG

This card is the reference source for T5-small, a compact and ubiquitous seq2seq backbone used for summarization, translation and classification in many pipelines. It provides citable architecture, pre-training, licensing and usage details (including the transformers v5 pipeline caveat), supporting accurate retrieval on small text-to-text models and fine-tuning baselines.
