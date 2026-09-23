---
id: collect-huggingface/huggingface/zai-org-glm-4-7-flash
title: "GLM-4.7-Flash - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "Nvidia", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["glm", "agent", "agentic", "attention", "benchmarks", "blackwell", "gpus", "license", "mit license", "moe", "open-weight", "parameters"]
source: docs/RAG/Collect RAG/03_huggingface/zai-org-GLM-4.7-Flash.md
source_anchor: ""
source_lines: [1, 52]
sha256: 33e0727461295b9d55b0cd1d27b513a39c51ada0d19c5382eb711341658cfc97
---

# GLM-4.7-Flash - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/zai-org/GLM-4.7-Flash
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

GLM-4.7-Flash is a 30B-A3B Mixture-of-Experts language model from Z.ai, positioned as the strongest model in the 30B class and a lightweight deployment option balancing performance and efficiency. It is licensed under MIT. The card references the GLM-4.7 technical blog and the GLM-4.5 technical report (arXiv 2508.06471) for architecture details.

It is a text-generation model with 31B reported parameters, in BF16/F32 tensors. Benchmarks: AIME 25 91.6, GPQA 75.2, LiveCodeBench v6 64.0, HLE 14.4, SWE-bench Verified 59.2, τ²-Bench 79.5, BrowseComp 42.8. Compared against Qwen3-30B-A3B-Thinking-2507 and GPT-OSS-20B, it leads on GPQA, HLE, SWE-bench Verified, τ²-Bench, and BrowseComp, showing strong agentic capability for its size.

Evaluation parameters: default temperature 1.0, top-p 0.95, max new tokens 131,072; Terminal Bench/SWE-bench Verified use temperature 0.7, top-p 1.0, max tokens 16,384; τ²-Bench uses temperature 0. It supports a "Preserved Thinking" mode for multi-turn agentic tasks.

Deployment: vLLM and SGLang (main branch only). Example commands use tensor-parallel-size 4, MTP speculative decoding, `--tool-call-parser glm47`, and `--reasoning-parser glm45`. For Blackwell GPUs, SGLang needs `--attention-backend triton`. It also works with Transformers. ~1.84M downloads per month; 102 community quantization variants.

## Key points

- 30B-A3B MoE model; MIT license; "strongest in the 30B class".
- Strong agentic results: SWE-bench Verified 59.2, τ²-Bench 79.5, BrowseComp 42.8.
- AIME 25: 91.6; GPQA: 75.2; LiveCodeBench v6: 64.0.
- Supports Preserved Thinking mode for multi-turn agent tasks.
- Served via vLLM / SGLang (main branch); glm47 tool-call parser.
- ~1.84M monthly downloads; 102 community quants available.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 30B (31B reported) |
| Active parameters | 3B (A3B MoE) |
| Architecture | MoE (glm4_moe_lite) |
| License | MIT |
| Precision | BF16 / F32 |
| Default sampling | temp=1.0, top_p=0.95, max_new_tokens=131072 |
| AIME 25 | 91.6 |
| GPQA | 75.2 |
| LiveCodeBench v6 | 64.0 |
| HLE | 14.4 |
| SWE-bench Verified | 59.2 |
| τ²-Bench | 79.5 |
| BrowseComp | 42.8 |
| Monthly downloads | ~1.84M |

## Why this source matters for the RAG

GLM-4.7-Flash anchors the lightweight-MoE segment of the GLM family in the knowledge base, with concrete agentic benchmarks and deployment commands for vLLM/SGLang. Its comparison tables against Qwen3-30B-A3B and GPT-OSS-20B are valuable for model-selection questions. It helps track the evolution from GLM-4.7 to GLM-5.x open-weight models.
