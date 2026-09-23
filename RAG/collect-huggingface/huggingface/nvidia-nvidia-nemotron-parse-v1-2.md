---
id: collect-huggingface/huggingface/nvidia-nvidia-nemotron-parse-v1-2
title: "NVIDIA-Nemotron-Parse-v1.2 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Nvidia", "TensorRT-LLM", "vLLM"]
dates: ["2026-02-17", "2026-09-23"]
keywords: ["nvidia", "agentic", "blackwell", "embedding", "license", "parameters", "tensorrt", "tensorrt-llm", "training", "vllm"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-NVIDIA-Nemotron-Parse-v1.2.md
source_anchor: ""
source_lines: [1, 45]
sha256: 7c690b6733993e1d602725b329d3b1f4de824db16e7371d948b2c3ab99c2066b
---

# NVIDIA-Nemotron-Parse-v1.2 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/NVIDIA-Nemotron-Parse-v1.2
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

NVIDIA Nemotron Parse v1.2 is a document-understanding model that extracts text and table elements with spatial grounding from document images. Given an image, it produces structured annotations including formatted text, bounding boxes and semantic classes, ordered according to the document's reading flow. It overcomes the limitations of traditional OCR on complex document layouts, enabling downstream benefits such as training data for LLMs, and improved extractor, curator, retriever and agentic applications. The architecture is a transformer-based vision-encoder-decoder: a ViT-H vision encoder (C-RADIO), a 1D-convolution adapter layer that compresses the latent space from 1280 to 320 tokens, an mBart decoder with 10 blocks, and under 1B total parameters. Input is RGB images (max 2048x1664, min 1024x1280) plus a prompt string; output is text encoding content plus bounding boxes and classes. v1.2 adds a fourth prompt-token category — Text-in-picture prompts (`<predict_text_in_pic>` / `<predict_no_text_in_pic>`) — controls whether text is extracted from embedded pictures, and now follows natural reading order across all semantic classes (footnotes, page-footers, tables, pictures, captions). It can classify objects such as title, section, caption, index, footnote, lists, tables, bibliography and image. Table output supports latex, HTML, markdown, JSON and CSV. It runs on TensorRT-LLM and vLLM (Ampere, Blackwell, Hopper, Turing; Linux) and includes two shared logits processors for repetition-stopping and table-structure enforcement. Training used millions of image-text items from document/table datasets (several TB) mixing public, synthetic and third-party OCR data. It is governed by the NVIDIA Nemotron Open Model License (tokenizer CC-BY-4.0), is ready for commercial use, and was released February 17, 2026. 66,712 downloads last month.

## Key points

- Document parsing VLM: extracts formatted text, tables and bounding boxes with reading-order grounding.
- Architecture: ViT-H (C-RADIO) encoder + 1D-conv adapter (1280→320 tokens) + mBart 10-block decoder; <1B params.
- v1.2 adds Text-in-picture prompt tokens and natural reading order for all semantic classes.
- Input max 2048x1664; output text with classes like title, section, table, footnote, bibliography, image.
- Table formats: latex, HTML, markdown, JSON, CSV.
- Runtimes: TensorRT-LLM, vLLM; Ampere/Blackwell/Hopper/Turing.
- License: NVIDIA Nemotron Open Model License; released 2026-02-17; 0.9B params.

## Technical data / figures

| Attribute | Value |
|---|---|
| Parameters | <1B (0.9B reported) |
| Architecture | Transformer vision-encoder-decoder (ViT-H + mBart 10 blocks) |
| Task | Image-Text-to-Text (document parse / OCR / feature extraction) |
| Input image | RGB, max 2048x1664, min 1024x1280 |
| Prompt tokens | predict_bbox, predict_classes, output_markdown/no_text, text_in_pic/no_text_in_pic |
| Output | Text + bounding boxes + semantic classes |
| Table formats | latex, HTML, markdown, json, csv |
| Runtimes | TensorRT-LLM, vLLM |
| Hardware | Ampere, Blackwell, Hopper, Turing (Linux) |
| Training data | Millions of image-text items, several TB |
| License | NVIDIA Nemotron Open Model License (tokenizer CC-BY-4.0) |
| Release date | 2026-02-17 |

## Why this source matters for the RAG

This card documents the document-parsing model that converts unstructured documents into machine-usable text, tables and layout structure, which is directly relevant to building document-processing and retrieval pipelines. It complements the embedding, VLM and LLM entries to cover the full NVIDIA RAG stack.
