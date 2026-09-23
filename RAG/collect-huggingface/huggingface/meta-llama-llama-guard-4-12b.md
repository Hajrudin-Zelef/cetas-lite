---
id: collect-huggingface/huggingface/meta-llama-llama-guard-4-12b
title: "Llama-Guard-4-12B - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Meta", "Microsoft"]
dates: ["2026-09-23"]
keywords: ["llama", "gpu", "license", "moe", "multimodal", "parameters", "pretraining", "pruning", "scout", "training"]
source: docs/RAG/Collect RAG/03_huggingface/meta-llama-Llama-Guard-4-12B.md
source_anchor: ""
source_lines: [1, 48]
sha256: faeb6db8da5b18ba51f4711bc71b0f02c42e1f07a5367975c13774d6c9fe368e
---

# Llama-Guard-4-12B - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/meta-llama/Llama-Guard-4-12B
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Llama Guard 4 is a natively multimodal safety classifier with 12 billion parameters, trained jointly on text and multiple images. It is a dense architecture pruned from Meta's Llama 4 Scout pre-trained model and fine-tuned for content safety classification. Like earlier versions, it acts as an LLM: it generates text indicating whether a prompt or response is safe or unsafe, and if unsafe, lists the violated content categories. It can classify content in both LLM inputs (prompt classification) and LLM responses (response classification). Llama Guard 4 is aligned to the standardized MLCommons hazards taxonomy and consolidates the previous Llama Guard 3-8B and Llama Guard 3-11B-vision models into a single classifier supporting English and multilingual text as well as mixed text-and-image prompts, and unlike the vision predecessor it supports multiple images per prompt. It uses an early-fusion transformer with dense layers to keep size small enough to run on a single GPU, and shares the same tokenizer and vision encoder as Llama 4 Scout and Maverick. The model is integrated into the Llama Moderations API for text and images. To build it, Meta took the pre-trained Llama 4 Scout checkpoint (one shared dense expert plus sixteen routed experts per MoE layer), pruned all routed experts and router layers, retained only the shared expert (turning MoE into a dense feedforward layer), and performed no additional pretraining. Post-training blended Llama Guard 3-8B and 3-11B-vision data with new multi-image data (2-5 images) and multilingual data, at roughly a 3:1 text-only to multimodal ratio. It predicts 14 hazard categories (S1-S14), including a text-only Code Interpreter Abuse category. Output-filtering evaluation versus Llama Guard 3 shows English recall 69% / FPR 11% / F1 61%, multilingual 43/3/51, single-image 41/9/38, and multi-image 61/9/52. The model is BF16 (12B params), licensed under Llama 4 Community License, with 45,595 monthly downloads.

## Key points

- 12B-parameter natively multimodal safety classifier (dense), pruned from Llama 4 Scout MoE.
- Acts as an LLM, classifying prompts (input filtering) and responses (output filtering) as safe/unsafe.
- Supports English + multilingual text and mixed/multiple image inputs; shares Scout/Maverick tokenizer and vision encoder.
- Aligned to the MLCommons taxonomy; predicts 14 hazard categories S1-S14.
- Consolidated successor to Llama Guard 3-8B and Llama Guard 3-11B-vision; integrated into the Llama Moderations API.
- Runs on a single GPU; BF16 weights; Llama 4 Community License.
- Output-filtering gains over LG3: multi-image recall +20%, single-image +10%, English F1 +8%.
- Requires a preview Transformers build (`v4.51.3-LlamaGuard-preview`) plus hf_xet.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | meta-llama |
| Model name | Llama-Guard-4-12B |
| Architecture | Early-fusion dense transformer (pruned Llama 4 Scout) |
| Parameters | 12B |
| Inputs | Multilingual text + up to multiple images |
| Hazard categories | S1-S14 (MLCommons taxonomy) |
| Filtering modes | Input filtering, output filtering, or both |
| Precision | BF16 |
| Hardware | Single GPU |
| License | Llama 4 Community License |
| Eval (output filtering) | English R 69% / FPR 11% / F1 61%; multi-image R 61% / F1 52% |
| Getting started | `pip install git+...@v4.51.3-LlamaGuard-preview hf_xet` |
| Downloads/month | 45,595 |
| arXiv | 2503.05731, 2407.21783 |

## Why this source matters for the RAG

This card documents Meta's flagship multimodal safety classifier, including its pruning-from-MoE construction and MLCommons hazard taxonomy. It is a key reference for AI safety moderation, input/output filtering pipelines, and guardrail models.
