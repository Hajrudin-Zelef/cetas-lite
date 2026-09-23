---
id: collect-huggingface/huggingface/zai-org-glm-ocr
title: "GLM-OCR - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["glm", "apache", "benchmarks", "license", "licenses", "mit license", "multimodal", "parameters", "sglang", "throughput", "training", "vllm"]
source: docs/RAG/Collect RAG/03_huggingface/zai-org-GLM-OCR.md
source_anchor: ""
source_lines: [1, 46]
sha256: ae7132f09a9cfb3d821654d647ef72adc03371e06e9a128ab1b6b0c4b32933fa
---

# GLM-OCR - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/zai-org/GLM-OCR
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

GLM-OCR is a multimodal OCR model from Z.ai for complex document understanding, built on the GLM-V encoder-decoder architecture and released under the MIT license. It introduces Multi-Token Prediction (MTP) loss and stable full-task reinforcement learning to improve training efficiency, recognition accuracy, and generalization. The model integrates the CogViT visual encoder, a lightweight cross-modal connector with efficient token downsampling, and a GLM-0.5B language decoder, combined with a two-stage pipeline of layout analysis and parallel recognition based on PP-DocLayout-V3.

With only 0.9B parameters (1B reported on HF, BF16), GLM-OCR achieves a score of 94.62 on OmniDocBench V1.5, ranking #1 overall, with state-of-the-art results across formula recognition, table recognition, and information extraction. It is optimized for real-world business scenarios: complex tables, code-heavy documents, seals, and other challenging layouts. Throughput is 1.86 pages/second for PDF documents and 0.67 images/second for images under identical hardware conditions, significantly outperforming comparable models.

Deployment supports vLLM, SGLang, and Ollama, making it suitable for high-concurrency and edge deployments. It supports 8 languages. The official SDK integrates PP-DocLayoutV3 for complete document parsing (layout analysis + structured output), and supports two prompt scenarios: document parsing (text/formula/table recognition) and information extraction (strict JSON schema). The complete OCR pipeline uses PP-DocLayoutV3, which is Apache-2.0 licensed; both licenses must be respected. Citation: arXiv 2603.10910. ~1.69M monthly downloads.

## Key points

- Multimodal OCR for complex documents; only 0.9B parameters.
- No.1 on OmniDocBench V1.5 (94.62); SOTA formula/table recognition and info extraction.
- GLM-V encoder-decoder with MTP loss and full-task RL; CogViT visual encoder.
- 1.86 pages/sec (PDF) and 0.67 images/sec throughput.
- Served via vLLM, SGLang, Ollama; official SDK with PP-DocLayoutV3.
- MIT license (pipeline component PP-DocLayoutV3 is Apache-2.0).

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 0.9B (1B reported) |
| Architecture | GLM-V encoder-decoder (CogViT + cross-modal connector + GLM-0.5B decoder) |
| Pipeline | PP-DocLayoutV3 layout analysis + parallel recognition |
| Languages | 8 |
| License | MIT (PP-DocLayoutV3: Apache-2.0) |
| Precision | BF16 |
| OmniDocBench V1.5 | 94.62 (No.1) |
| Throughput | 1.86 pages/s (PDF), 0.67 images/s |
| Frameworks | vLLM, SGLang, Ollama, Transformers |
| Monthly downloads | ~1.69M |

## Why this source matters for the RAG

GLM-OCR is a high-value model for document ingestion inside RAG pipelines, converting PDFs, tables, and formulas into clean structured text. Its small footprint, edge-deployability, and top document-understanding benchmarks are directly actionable for RAG document processing decisions. It grounds the knowledge base on modern OCR and document-parsing tooling.
