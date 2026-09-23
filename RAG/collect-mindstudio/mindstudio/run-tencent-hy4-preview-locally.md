---
id: collect-mindstudio/mindstudio/run-tencent-hy4-preview-locally
title: "How to Run Tencent's Hy4 Preview Locally with vLLM or SGLang"
domain: mindstudio
role: reference
task: article
actors: ["DeepSeek", "Hugging Face", "Moonshot", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["sglang", "vllm", "apache", "attention", "consumer", "deepseek", "fp8", "glm", "gpu", "gpus", "inference", "kimi"]
source: docs/RAG/Collect RAG/02_mindstudio/run-tencent-hy4-preview-locally.md
source_anchor: ""
source_lines: [1, 56]
sha256: 662bc89658d797179b7026a41e8b9b7da0a6d1229173529bf618ced61fb310f5
---

# How to Run Tencent's Hy4 Preview Locally with vLLM or SGLang

## Metadata

- **Source** : https://www.mindstudio.ai/blog/run-tencent-hy4-preview-locally
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article is a practical guide to deploying **Tencent's Hy4 preview**, Tencent's newest open-weight flagship model, locally using **FP8 weights**, **tensor parallelism**, and prebuilt **vLLM** or **SGLang** Docker images.

Hy4 preview is a **Mixture-of-Experts (MoE)** architecture with **770 billion total parameters** and **49 billion activated per token**, shipped under **Apache 2.0** in both a full-precision release and an **FP8 quantized** version. It is designed to be served through standard inference engines rather than run as a raw checkpoint, making it a multi-GPU deployment exercise rather than a laptop chatbot. The architecture spans 78 layers where the first is dense and the remaining 77 use MoE with **256 routed experts plus 1 shared expert**. It uses **Gated DeepSeek Sparse Attention (Gated DSA)** with IndexCache for cross-layer index reuse, plus identity Hyper-Connections (iHC) on the residual pathway — choices adapted from DeepSeek and GLM. Weights are distributed via Hugging Face, ModelScope, GitCode, and CNB.

Both official deployment recipes default to **8-way tensor parallelism** (`--tensor-parallel-size 8` / `--tp-size 8`), meaning an 8-GPU node is the documented baseline; this is not a single-consumer-GPU model. The vLLM path uses the `vllm/vllm-openai:hy4-preview` image with `--speculative-config '{"num_speculative_tokens":3,"method":"mtp"}'`, the required `FLASHMLA_SPARSE` attention backend, and `hy_v4` tool-call and reasoning parsers. The SGLang path uses `lmsysorg/sglang:hy4-preview` (multi-arch, x86 and Arm) with `auto` parsers and `NEXTN` speculative algorithm (`--speculative-num-steps 3`, `--speculative-eagle-topk 1`, `--speculative-num-draft-tokens 4`). Both expose an OpenAI-compatible API on port 8000.

The model supports speculative decoding out of the box via a built-in **MTP (multi-token prediction)** layer of roughly **10B total / 0.7B activated** parameters. Context length is rated at **1 million tokens**, and it defaults to a "high" reasoning effort mode that can be disabled via `extra_body={"chat_template_kwargs": {"reasoning_effort": "no_think"}}` for faster responses. Tencent recommends `temperature=0.9` and `top_p=1.0`.

On capability, Tencent's internal blind evaluation — **163 experts rating 203 real engineering tasks** — placed Hy4 preview slightly ahead of **GLM 5.3** (2.99 vs 2.92 average, winning 46.8% of head-to-heads) and slightly ahead of **Kimi K3** (2.99 vs 2.94, winning 51.2%). Tencent describes it as the largest generation-over-generation improvement for this model line. However, it is explicitly labeled a preview with known limitations: a tendency to over-reason and over-verify on simpler tasks. Tencent provides a full finetuning pipeline and the AngelSlim quantization/compression toolkit.

## Key points

- Hy4 preview is a 770B MoE model with 49B active parameters per token, Apache 2.0, released as full precision and FP8.
- 78 layers: first dense, remaining 77 MoE with 256 routed experts + 1 shared expert.
- Gated DeepSeek Sparse Attention (Gated DSA) + IndexCache and identity Hyper-Connections (iHC).
- Official deployment uses 8-way tensor parallelism via vLLM or SGLang Docker images.
- Built-in MTP layer (~10B total / 0.7B active) enables speculative decoding.
- 1M-token context; defaults to "high" reasoning effort, togglable for faster answers.
- Blind eval: slightly ahead of GLM 5.3 and Kimi K3 across 203 engineering tasks.
- Explicit preview caveats: over-reasoning and over-verification on simple tasks.

## Technical data / figures

| Item | Value |
|---|---|
| Total parameters | 770B (MoE) |
| Active parameters per token | 49B |
| Layers | 78 (1 dense + 77 MoE) |
| Experts | 256 routed + 1 shared |
| License | Apache 2.0 |
| Checkpoints | Hy4-preview (full) / Hy4-preview-FP8 |
| Context length | 1,000,000 tokens |
| MTP layer | ~10B total / 0.7B active |
| Speculative decoding | vLLM: mtp, 3 tokens; SGLang: NEXTN, 3 steps / topk 1 / 4 draft tokens |
| Baseline deployment | 8 GPUs, tensor parallel |
| vLLM attention backend | FLASHMLA_SPARSE |
| Defaults | temperature=0.9, top_p=1.0 |
| Eval vs GLM 5.3 | 2.99 vs 2.92 (win 46.8%) |
| Eval vs Kimi K3 | 2.99 vs 2.94 (win 51.2%) |

## Why this source matters for the RAG

It provides an up-to-date, deployment-focused reference on a frontier-scale open-weight MoE model, including exact Docker commands, architectural details, and evaluation figures. This is directly useful for answering questions about self-hosting large MoE models and about the vLLM/SGLang serving ecosystem.
