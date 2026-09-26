---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/compression-zoo-one-liners-dated-for-rag-retrieval
title: "Compression-zoo one-liners (dated, for RAG retrieval)"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["DeepSeek", "Huawei", "Hugging Face", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2025-05", "2026-03-21", "2026-08-26"]
keywords: ["ascend", "attention", "decode", "deepseek", "fp4", "fp8", "inference", "sglang", "tensorrt", "tensorrt-llm", "tpu", "vllm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3176, 3186]
section: "7. KV Cache & Long-Context Techniques"
sha256: b186617484ee632120effb34c5a70000e2534a7da5d925652c72cdaf27f0bb38
---

# Compression-zoo one-liners (dated, for RAG retrieval)

| Engine | First native MLA support | 2026 status |
|---|---|---|
| **vLLM** | Late 2024 (v0.6.x era; Triton MLA path for DeepSeek-V2) — exact minor [UNVERIFIED] | v0.23.0 (2026): DeepSeek-V4 production hardening (TRTLLM-gen kernel, sparse MLA metadata decoupled from V3.2); v0.28.0 (tagged 2026-08-26): sparse MLA end-to-end (plain decode, MTP, speculative); 2026 priority: `FLASHINFER_MLA_SPARSE` with FP8 KV |
| **SGLang** | v0.3, Sep 2024 ("7× Faster DeepSeek MLA") | PR #6109 (May 2025): FlashMLA + FP8 KV + MTP; v0.5.6 (Dec 2025): FP4 support for MHA+MLA KV caches; SGLang-JAX (TPU): MLA via FlashAttention Pallas kernel, no extra flag |
| **TensorRT-LLM** | Supported for DeepSeek-V3/V3.2/R1-class models — exact release [UNVERIFIED] | The TRTLLM-generation attention kernel was adopted *back* into vLLM v0.23.0 for DeepSeek-V4 |
| **Hugging Face TGI** | Gaudi branch only (v3.3.x, late 2025) — never mainline CUDA | **Archived read-only 2026-03-21** (maintenance mode) |

→ Part 06 (inference engines): full per-release kernel tables, backend-priority logic, fp8_e4m3-vs-fp8_e5m2 engine defaults, and the vLLM-Ascend / XPU / metal out-of-tree ports.

### Compression-zoo one-liners (dated, for RAG retrieval)

