---
id: ai-industry-kb-2026/06-inference-engines/execution-stack-facts-2026-verified-from-release-notes
title: "Execution stack facts (2026, verified from release notes)"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "AWS", "Alibaba", "DeepSeek", "Google", "Huawei", "Intel", "Microsoft", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2026-04", "2026-07", "2026-08-17", "2026-08-20"]
keywords: ["amd", "ascend", "attention", "awq", "aws", "benchmark", "cost", "datacenter", "decode", "deepseek", "diffusion", "disaggregated"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2327, 2352]
section: "6. Inference Engines"
sha256: 129bdd1d5bbac9d60b430e2b968a7c5aed8ca8c363205295737fe60b751e55ea
---

# Execution stack facts (2026, verified from release notes)

- **TPU (Google):** `vllm-project/tpu-inference` plugin (formerly `vllm-tpu`), JAX/XLA backend by Google, installed via `uv pip install vllm-tpu`. Compatible generations: **v7x (Ironwood), v6e (Trillium), v5e recommended; v3/v4/v5p experimental**. Support matrices updated 2026-08-20/27: chunked prefill ✅, KV cache offload ✅, LoRA ✅, Eagle3/Ngram speculative decoding ✅, TP/DP/EP/SP ✅, PP ✅ (single- and multi-host), MXFP4/FP8/INT4/INT8 quants ✅ — **no MXFP8 support anywhere in the TPU codebase** (verified by code search). April 2026 community report: Gemma 3 27B, Llama 3.1 8B, Llama 3.3 70B fully passing; **Gemma 4 variants FAILING on nightly** (shared-layer KV models broken — operators serve plain checkpoints instead of QAT ones). AccelMark runner: multi-chip and streaming/async skipped (`is_async_output_supported = False`) — TPU backend strongest for **offline/batch**, weaker for interactive serving.
- **AWS Inferentia/Trainium:** two paths — **vllm-neuron plugin (Beta)** (independent of NxD Inference) and **NxD Inference (neuronx-distributed-inference)** plugging into vLLM's Plugin System (keeps vLLM input processing, scheduling, output processing). **Neuron SDK 2.32.0 released 2026-08-17**. Neuron docs: runtime **"implements vLLM V1 API compatibility on Trainium and Inferentia with optimizations for large-scale inference workloads"** — expert parallelism for MoE, disaggregated inference architectures, speculative decoding (Eagle V1), INT8/FP8 quantization, prefix caching, multimodal Llama 4 Scout/Maverick. Supported models: Llama 2/3.1/3.3, Llama 4 Scout/Maverick, Qwen 2.5, Qwen 3, plus custom models onboarded to NxD Inference. Instances: Inf2, Trn1, Trn2. Independent evidence: AWS sample GenAI-on-EKS starter kit lists "Neuron Support: AWS Inferentia2/Trainium accelerators" in its vLLM component doc, with a `qwen3-8b-neuron` config entry. Caveats: plugin still Beta; older docs note PagedAttention-on-Neuron lagged CUDA (FP16/BF16 only, no pre-built wheels — build from source) [DIRECTIONAL].
- **AMD/ROCm:** MI300X/MI325X/MI355 class functional; community forks document BF16 safetensors, GPT-OSS MXFP4 via built-in Triton MXFP4 MoE path, FP8 KV cache (~halving KV memory), hybrid RDNA W4A16 AWQ INT4 paths on gfx11/gfx12; GGUF explicitly excluded in those forks.
- **Other:** vllm-ascend (Huawei), vllm-spyre (IBM Spyre), vllm-gaudi (Intel Gaudi), vllm-openvino; Intel GPU/CPU and XPU functional per the vLLM V1 guide; FlagOS tracks vLLM against a **20+ chip test base**.

### Execution stack facts (2026, verified from release notes)

- Engine core: V1 only (v0.16.0+); API server path multiprocess `AsyncLLM` over ZMQ; single-process mode unsupported on the API path.
- Execution: Model Runner V2 default for dense models (v0.25.0+); full CUDA-graph capture of decode steps (record once, replay per step); PagedAttention deleted (#47361) in favor of the V1/MRv2 backends.
- Modeling: Transformers v5 backend at parity with native vLLM speed (v0.25.0), FP8 MoE; DeepSeek V4 with DSA attention backend, MTP, KV Pool; Hunyuan v3; Granite 4.1 Vision built-in.
- Attention backends: FlashAttention 4 default MLA prefill (SM90+, head-dim 512, paged-KV); FlashAttention 3/4 prefill with TurboQuant 2-bit KV cache (4x capacity); DeepSeek MLA on SM90+.
- Speculative decoding: dynamic speculative decoding with full CUDA graphs; universal speculative decoding across heterogeneous vocabularies (TLI #38174); drafters DSpark (#46995), DFlash (#46770, #46853); DFlash2 in v0.28.0.
- Structured/tool-call output: XGrammar default backend (v0.7+); Streaming Parser Engine (#46610): unified tool-call/reasoning parsing (Kimi k2.5/k2.6/k2.7, seed_oss, DeepSeek V4 parsers).
- Online quantization frontend (v0.20.0): experts_int8 folded into FP8 online path; MXFP8 moved to the new frontend; MiniMax-M3 NVFP4 support. (Dtype mechanics → §8.)
- Observability/ops: Rust frontend with HTTPS/mTLS and DP supervisor; profiler control routes; CUDA 13.0 default wheels; Python 3.14; PyTorch 2.11.

### The TensorRT-LLM comparator (2026 evidence, vLLM-framed)

- **2026 release notes (current):** B300/GB300 support; dedicated **disaggregated-serving** performance tests (`test_perf.py`); `benchmark_serving` for multimodal models; NIM perf test cases; **KV Cache Connector** documentation; LoRA guidance and AutoDeploy docs; deployment guides for **GPT-OSS, DeepSeek-R1, VDR 1.0**; tech blogs on guided+speculative decoding and ADP balance strategy; base images `nvcr.io/nvidia/pytorch:25.10-py3` / `tritonserver:25.10-py3`; dependency snapshot PyTorch 2.9.0, ModelOpt 0.37, xgrammar 0.1.25, transformers 4.56.0, NIXL 0.5.0.
- **v1.3.0rc21–rc23 (July 2026):** incremental line — documentation generation, shared-expert combine fusion, paged MQA logits decode tuning, fused RMSNorm/RoPE, EAGLE3 dynamic-tree kernels; **day-0 support for OpenAI GPT-OSS-120B/20B (Aug 5, 2026)**; **the AutoDeploy backend is deprecated** as of rc21 — release notes signal a shift to "more robust, future-proof integration methods". Ongoing per the notes: Gemma 4 multimodal, Kimi K2.5 vision, GPT-OSS/Ministral3/Nemotron/DeepSeek enablement, KV reuse v2 manager (`cache_salt_id` support), disaggregated serving and transfer-path improvements, FP4/FP8 decode kernels, VisualGen/diffusion pipelines (SageAttention for Wan/FLUX).
- **Positioning:** strongest on datacenter NVIDIA silicon with deep quantization (FP8 headline path — community skill docs cite ~2x throughput vs FP16 on Llama 3-70B), custom CUDA kernels and graph compilation (5–20% over vLLM single-GPU per community guides), NIM packaging, and guided+speculative decoding — **at the cost of a required model-conversion step and higher operational complexity**. With NIM 2.0 moving LLM/VLM NIMs to vLLM, TRT-LLM's center of gravity shifts to embedding/speech/edge NIMs, Triton backends, and the maximum-throughput datacenter niche.
- **Standard production stacks:** TensorRT-LLM = Client → Triton Inference Server → tensorrtllm_backend (Triton handles HTTP/gRPC, queuing, health checks, metrics; TRT-LLM does inference with in-flight batching), plus the full-lifecycle **NeMo → TensorRT-LLM → NIM** pipeline. vLLM = proprietary NIM orchestration over OSS vLLM (NIM 2.0) or Red Hat AI Inference Server (enterprise).
- **Decision rule (2026 practitioner consensus):** switch to TensorRT-LLM only when peak NVIDIA-datacenter throughput justifies the operational overhead; default remains vLLM (general) or SGLang (prefix-heavy).

### llama.cpp v0.4.0 and the GGUF authority relationship

