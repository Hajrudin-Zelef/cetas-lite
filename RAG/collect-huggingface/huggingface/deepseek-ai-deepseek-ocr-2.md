---
id: collect-huggingface/huggingface/deepseek-ai-deepseek-ocr-2
title: "DeepSeek-OCR-2 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["DeepSeek", "Hugging Face", "Nvidia", "vLLM"]
dates: ["2026-09-23"]
keywords: ["deepseek", "apache", "attention", "benchmark", "benchmarks", "gpus", "inference", "license", "nvidia", "parameters", "safetensors", "vllm"]
source: docs/RAG/Collect RAG/03_huggingface/deepseek-ai-DeepSeek-OCR-2.md
source_anchor: ""
source_lines: [1, 50]
sha256: 7c6e386ad7c343b53f02ec680a62444e04bc03f926c3caadcc8d07c0f2315c57
---

# DeepSeek-OCR-2 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/deepseek-ai/DeepSeek-OCR-2
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

DeepSeek-OCR 2 ("Visual Causal Flow") is DeepSeek-AI's second-generation OCR / document understanding model, an Image-Text-to-Text vision-language model with 3B parameters in BF16, licensed under Apache-2.0, and tagged multilingual, deepseek_vl_v2, vision-language, and ocr. It explores more human-like visual encoding ("Explore more human-like visual encoding" being the tagline), succeeding the original DeepSeek-OCR which introduced Contexts Optical Compression (arXiv 2510.18234). DeepSeek-OCR-2 is documented in the paper "DeepSeek-OCR 2: Visual Causal Flow" (arXiv 2601.20552). The model has 863,675 monthly downloads and supports feature-extraction-style usage in addition to OCR.

Usage is demonstrated with Hugging Face Transformers on NVIDIA GPUs (tested with Python 3.12.9 + CUDA 11.8, torch 2.6.0, transformers 4.46.3, tokenizers 0.20.3, flash-attn 2.7.3). Example code loads the model via `AutoModel.from_pretrained(model_name, _attn_implementation='flash_attention_2', trust_remote_code=True, use_safetensors=True)` in bfloat16 and calls `model.infer(tokenizer, prompt=..., image_file=..., output_path=..., base_size=1024, image_size=768, crop_mode=True, save_results=True)`. vLLM acceleration and PDF processing guidance are referenced on the GitHub repo.

The model supports dynamic resolution: default (0-6)x768x768 + 1x1024x1024 tiles, producing (0-6)x144 + 256 visual tokens. Two main prompts are documented: for documents, `<image>\n<|grounding|>Convert the document to markdown.`; for OCR without layout preservation, `<image>\nFree OCR.` Evaluation results on the Hub include olmOCR-bench overall 76.3 (Arxiv Math 82, Old Scans Math 72) and ParseBench mean 41.2 (deepseekocr2_vllm pipeline). Acknowledgement is given to DeepSeek-OCR, Vary, GOT-OCR2.0, MinerU, PaddleOCR, and the OmniDocBench benchmark. Two citations are provided (2025 and 2026 arXiv papers by Wei, Sun, Li).

## Key points

- Second-generation DeepSeek OCR model ("Visual Causal Flow"), 3B params, Image-Text-to-Text, Apache-2.0.
- Multilingual; tagged deepseek_vl_v2, vision-language, ocr, feature-extraction.
- Dynamic resolution: (0-6)x768x768 + 1x1024x1024 tiles → (0-6)x144 + 256 visual tokens.
- Main prompts: `<|grounding|>Convert the document to markdown.` and `Free OCR.`
- Runs via Transformers (flash-attention-2, bfloat16) with an `infer()` helper; vLLM for acceleration.
- Benchmarks: olmOCR-bench overall 76.3 (Arxiv Math 82, Old Scans Math 72); ParseBench mean 41.2.
- Acknowledges DeepSeek-OCR, Vary, GOT-OCR2.0, MinerU, PaddleOCR, OmniDocBench.
- 863,675 downloads/month; arXiv 2601.20552 and 2510.18234.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | deepseek-ai |
| Model name | DeepSeek-OCR-2 |
| Type | OCR / vision-language (Image-Text-to-Text) |
| Parameters | 3B |
| Weight format | BF16 |
| License | Apache-2.0 |
| Languages | Multilingual |
| Dynamic resolution | (0-6)x768x768 + 1x1024x1024; (0-6)x144 + 256 visual tokens |
| Serving | Transformers (flash-attn 2.7.3), vLLM |
| Key benchmarks | olmOCR-bench 76.3 overall; ParseBench mean 41.2 |
| arXiv | 2601.20552 (DeepSeek-OCR 2); 2510.18234 (DeepSeek-OCR) |
| Downloads/month | 863,675 |

## Why this source matters for the RAG

This card documents a lightweight, open OCR/document-understanding model with concrete inference recipes and resolution mechanics, useful for retrieval on document parsing, OCR pipelines, and vision-language model usage. Its small size and clear API make it a practical reference for deployment-oriented searches.
