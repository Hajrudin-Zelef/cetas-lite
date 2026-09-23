---
id: collect-huggingface/huggingface/google-bert-bert-base-uncased
title: "bert-base-uncased - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Google", "Hugging Face"]
dates: ["2026-09-23"]
keywords: ["apache", "embeddings", "fine-tuning", "license", "parameters", "research", "tpu", "training"]
source: docs/RAG/Collect RAG/03_huggingface/google-bert-bert-base-uncased.md
source_anchor: ""
source_lines: [1, 47]
sha256: cbafc0248e3016826575a49e4390459a1afe15b60895503c846c36a66fe6820d
---

# bert-base-uncased - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/google-bert/bert-base-uncased
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

bert-base-uncased is the reference English BERT base model from Google, a Transformers encoder pre-trained in a self-supervised fashion with two objectives: masked language modeling (MLM) and next sentence prediction (NSP). It has 110M parameters (0.1B hub size), is Apache 2.0 licensed, and remains one of the most used models on the hub (~45.5 million downloads/month, 6,951 fine-tunes, 142 adapters, 32 quantizations). It was introduced in "BERT: Pre-training of Deep Bidirectional Transformers for Language Understanding" (arXiv 1810.04805) and released via google-research/bert. Because the original team did not publish a model card, this card was written by the Hugging Face team. "Uncased" means the model does not distinguish between "english" and "English", and it strips accent markers.

The model is intended primarily to be fine-tuned on downstream tasks that use a whole sentence (potentially masked) for decisions: sequence classification, token classification, question answering. It is not meant for text generation (GPT-style models are recommended for that). Pre-training data is BookCorpus (11,038 unpublished books) plus English Wikipedia (excluding lists, tables, headers). Inputs are lowercased and WordPiece-tokenized with a vocabulary of 30,000, formatted as `[CLS] Sentence A [SEP] Sentence B [SEP]` (sentence B is a real continuation 50% of the time). Masking: 15% of tokens masked, of which 80% become `[MASK]`, 10% a random token, 10% unchanged. Training used 4 cloud TPUs in pod configuration (16 TPU chips) for 1M steps, batch 256, Adam with lr 1e-4, β1=0.9, β2=0.999, weight decay 0.01, warmup 10k steps, seq length 128 for 90% of steps and 512 for the remaining 10%.

GLUE test results after fine-tuning: MNLI-(m/mm) 84.6/83.4, QQP 71.2, QNLI 90.5, SST-2 93.5, CoLA 52.1, STS-B 85.8, MRPC 88.9, RTE 66.4, average 79.6. The card also documents known gender bias (e.g., "The man worked as a [MASK]" predicts carpenter/waiter; "The woman worked as a [MASK]" predicts nurse/waitress/maid/prostitute), which affects all fine-tuned versions.

## Key points

- Google's BERT base, English uncased, 110M params, MLM + NSP pre-training.
- BookCorpus + English Wikipedia; WordPiece vocab 30k.
- Primarily for fine-tuning on classification, tagging, QA.
- GLUE average 79.6 (MNLI 84.6/83.4, SST-2 93.5, QNLI 90.5, CoLA 52.1).
- Apache 2.0; ~45.5M downloads/month; 6,951 fine-tunes.
- Documented gender bias affecting downstream models.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | google-bert |
| Model name | bert-base-uncased |
| Architecture | BERT base encoder (12 layers, 768 hidden) |
| Params | 110M |
| Pre-training | MLM + NSP on BookCorpus + Wikipedia |
| Tokenizer | WordPiece, vocab 30,000, uncased |
| License | Apache 2.0 |
| Languages | English |
| GLUE test | Avg 79.6 (MNLI 84.6/83.4, QQP 71.2, QNLI 90.5, SST-2 93.5, CoLA 52.1, STS-B 85.8, MRPC 88.9, RTE 66.4) |
| Paper | arXiv 1810.04805 |
| Downloads/month | ~45,461,966 |

## Why this source matters for the RAG

This card is the authoritative reference for BERT base, the foundational encoder behind most classic embeddings, classification and QA systems used in retrieval. It provides citable pre-training details, GLUE figures, licensing, and known-bias caveats that are essential for accurate model-selection and background knowledge in RAG documentation.
