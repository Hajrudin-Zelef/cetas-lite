---
id: collect-huggingface/huggingface/zai-org-glm-5-3-flash
title: "GLM-5.3-Flash - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Hugging Face", "SGLang", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["glm", "agentic", "attention", "benchmark", "benchmarks", "claude", "compute", "cost", "deepseek", "license", "mit license", "moe"]
source: docs/RAG/Collect RAG/03_huggingface/zai-org-GLM-5.3-Flash.md
source_anchor: ""
source_lines: [1, 52]
sha256: f4467b0d0c556af728a7883c8c14e416e7653ddb42567ec84d582fcc84ad9767
---

# GLM-5.3-Flash - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/zai-org/GLM-5.3-Flash
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

GLM-5.3-Flash is the first natively multimodal model of the GLM-5 series, released by Z.ai under the MIT license. It has 320B total parameters with only 18B active parameters, outperforming GLM-5.2 across benchmarks and real-world workloads at roughly one-tenth the price, while approaching Claude Opus 4.8 on coding and agentic benchmarks. The card references the GLM-5 Technical Report (arXiv 2602.15763) and the GLM-5.3-Flash blog.

The model starts from a newly trained base model with a redesigned architecture and training recipe. For the first time in the GLM series, it uses a hybrid architecture combining sparse and linear attention, sharply reducing long-context serving costs while preserving precise long-context capabilities. It adopts Manifold-Constrained Hyper-Connections (mHC) to improve scaling efficiency, and was trained on a 30T-token multimodal pre-training corpus, delivering more intelligence with less compute.

Deployment is supported across SGLang, vLLM, TokenSpeed, Transformers, KTransformers, and Unsloth. GLM-5.3-Flash controls its thinking budget via the `reasoning_effort` parameter with three levels (low, high, max; defaults to max). In the chat template, `clear_thinking` defaults to false; for chat scenarios it should be passed as true. HF-reported model size is 321B params, tensor types BF16/F8_E4M3/F32.

Evaluation results: Terminal-Bench 2.1 84.3, Deep-SWE 63.4, and ExtractBench 80.75 (mean), 96.3 (short), 51.56 (medium) are reported directly on the card. The benchmarks are run under careful settings (temperature, context windows up to 1M, max generation up to 163,840 tokens). It has ~3.77M monthly downloads and 118 community quantization variants.

## Key points

- First natively multimodal GLM-5 model: 320B total / 18B active MoE.
- Hybrid sparse + linear attention; Manifold-Constrained Hyper-Connections (mHC).
- MIT license; outperforms GLM-5.2 at ~1/10th the price.
- Approaches Claude Opus 4.8 on coding/agentic benchmarks.
- `reasoning_effort` controls thinking budget (low/high/max).
- Served via SGLang, vLLM, TokenSpeed, Transformers, KTransformers, Unsloth.
- Terminal-Bench 2.1: 84.3; Deep-SWE: 63.4; ExtractBench mean: 80.75.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 320B (321B reported) |
| Active parameters | 18B |
| Architecture | Hybrid sparse + linear attention, MoE, mHC |
| Multimodality | Native (image-text-to-text) |
| Training corpus | 30T-token multimodal pre-training |
| License | MIT |
| Precision | BF16 / F8_E4M3 / F32 |
| reasoning_effort | low / high / max (default max) |
| Terminal-Bench 2.1 | 84.3 |
| Deep-SWE | 63.4 |
| ExtractBench (mean / short / medium) | 80.75 / 96.3 / 51.56 |
| Community quants | 118 models |
| Monthly downloads | ~3.77M |

## Why this source matters for the RAG

GLM-5.3-Flash is a landmark open-weight multimodal MoE with an extremely strong cost/performance ratio, making it central to the GLM family coverage in the knowledge base. Its hybrid attention and mHC architecture and benchmark footnotes provide authoritative, current facts on frontier Chinese open models. It supports comparisons with Qwen, DeepSeek, and other open-weight models across coding and agentic tasks.
