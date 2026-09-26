---
id: ai-industry-kb-2026/06-inference-engines/tensorrt-llm-2026-state
title: "TensorRT-LLM — 2026 state"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["Alibaba", "Apple", "DeepSeek", "Hugging Face", "Meta", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "Unsloth", "vLLM"]
dates: ["2025-10", "2025-12", "2025-12-15", "2026-07", "2026-07-14", "2026-08", "2026-08-05", "2026-09", "2026-09-05"]
keywords: ["tensorrt", "tensorrt-llm", "agentic", "attention", "awq", "benchmark", "blackwell", "consumer", "cost", "datacenter", "decode", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2752, 2788]
section: "6. Inference Engines"
sha256: 468825c9ba83acd8b2bd4ae1a00cc7f91d325f577b63889d9e7a745dff1a6f10
---

# TensorRT-LLM — 2026 state

- Origin: DeepSeek-V2's MLA paper reported 93.3% KV-cache reduction and 5.76x throughput vs standard MHA — if per-token KV shrinks ~15x, batch capacity (KV-bound in decode) grows several-fold, making ~4x batch the natural arithmetic consequence.
- [PARTIALLY VERIFIED]: direction and mechanism are confirmed by the paper figures and by both engines' MLA-native paths (SGLang: DP attention up to 1.9x decode throughput, FP4 MLA KV caches, FlashInfer SM120 sparse-MLA kernels; vLLM: FlashAttention 4 default MLA prefill since v0.20.0 — see 06a), but no independent framework-level benchmark was found measuring "4x batch size per GPU" as a stated, reproduced result.
- Present "4x batch" as a derived planning figure, not a measured engine benchmark — published figures are 5.76x throughput (paper) and 1.9x decode (SGLang DP attention), which are different metrics.
- What would settle it: a controlled measurement — same GPU class (e.g. 8×H100), same MLA model, same precision, engine-only variable (vLLM MLA vs a non-MLA baseline serving path) — measuring maximum sustainable concurrent batch before OOM or SLO breach, not throughput.
- The 2026 compounding pattern: MLA compresses the KV cache, FP8/FP4 quantizes it, disaggregation moves it, RadixAttention shares it — four independent multipliers on cost-per-token (KV technique details in §7).

### TensorRT-LLM — 2026 state

- Release notes (current): B300/GB300 support; dedicated disaggregated-serving performance tests (`test_perf.py`); `benchmark_serving` for multimodal models; NIM perf test cases; KV Cache Connector documentation; LoRA guidance and AutoDeploy docs; deployment guides for GPT-OSS, DeepSeek-R1, VDR 1.0; tech blogs on guided+speculative decoding and ADP balance strategy.
- Dependency snapshot: base images `nvcr.io/nvidia/pytorch:25.10-py3` / `tritonserver:25.10-py3`; PyTorch 2.9.0, ModelOpt 0.37, xgrammar 0.1.25, transformers 4.56.0, NIXL 0.5.0.
- Quantization: FP8 is the headline path (~2x throughput vs FP16 on Llama 3-70B per community docs [COMMUNITY]); GPTQ/AWQ supported; NIM packaging standardizes the deployment artifact.
- v1.3.0rc21–rc23 (July 2026): documentation generation, shared-expert combine fusion, paged MQA logits decode tuning, fused RMSNorm/RoPE, EAGLE3 dynamic-tree kernels — plus day-0 support for OpenAI GPT-OSS-120B / GPT-OSS-20B (2026-08-05); the AutoDeploy backend is deprecated as of rc21, signaling a shift to "more robust, future-proof integration methods."
- **Gaps carried into consolidation:** the MLA introduction version is NOT pinned; the FP8-KV introduction version is NOT pinned — keep both as open version gaps, do not backfill by guessing.
- Cautionary evidence (independent, August 2026, Qwen3-Coder-30B-A3B, concurrency 1/8/32/64): vLLM led at 1 (12.3 vs 10.8), 8 (56.3 vs 44.4), and 64 (129.0 vs 115.1); TRT-LLM led only at 32 (120.5 vs 119.1, +1%) — published 10–25% throughput claims did not hold on workstation Blackwell, likely because sm_120 lacks FlashAttention-4; always measure on your own hardware class.
- Positioning: strongest on datacenter NVIDIA with deep quantization, NIM packaging, and guided decoding; with NIM 2.0 moving LLM/VLM NIMs to vLLM (see 06a), TRT-LLM's center of gravity shifts to embedding/speech/edge NIMs and the maximum-throughput datacenter niche — not the default serving choice.
- The standard production stack remains Client → Triton Inference Server → tensorrtllm_backend (Triton handling HTTP/gRPC, queuing, health checks, metrics; TRT-LLM doing inference with in-flight batching), plus the full-lifecycle NeMo → TensorRT-LLM → NIM pipeline.
- 2026 decision rule: switch to TensorRT-LLM only when peak NVIDIA-datacenter throughput justifies the model-conversion overhead and higher operational complexity; default remains vLLM (general) or SGLang (prefix-heavy).

### llama.cpp server — 2026 state

- v0.4.0 (~2026-09-05): server changes — per-slot context limit, `data:` URLs for media, `preserve_reasoning` enabled by default, reject prefilled assistant tool calls, synthetic speculative-decoding acceptance options, pytest-xdist server tests.
- UI changes: Chat Form Actions UX, MCP overrides replaced by tool policy, Settings/MCP moved to dialogs, grouped agentic response text copy; ggml bumped v0.22.0 → v0.23.0 (sparse flash attention, RPC event/async APIs, Apple RDMA transport).
- Model router (announced 2025-12-15): multi-process design — each model in its own process (crash isolation), dynamic loading/switching via OpenAI-compatible API, automatic GGUF discovery, on-demand loading, LRU eviction (default max 4 models) — Ollama-like convenience in the C/C++ stack.
- 2026 role: the GGUF reference runtime and the CPU/edge/consumer-GPU default; Unsloth Desktop/Studio build on llama.cpp binaries (signed Windows binaries); vLLM's GGUF plugin reuses llama.cpp-style CUDA kernels (`mmvq.cuh`, `mmq.cuh`); when vLLM's plugin hits a model-family bug, the first diagnostic is always "does it run in llama.cpp?"

### Emerging and adjacent runtimes (2026)

- TileRT: new decode-focused runtime; the 2026-07-14 vLLM blog documents production pairing — vLLM owns prefill/scheduling/APIs, TileRT owns decode — targeting latency-bound workloads (agentic loops, coding assistants, real-time voice) where per-user speed is the scaling dimension.
- Tiny-vLLM: community-launched lightweight C++/CUDA framework for resource-constrained inference (edge, research budgets); Hacker News traction (~129 upvotes per dev.to coverage).
- Neutree 1.2 / Flex Engine (Arcfra, September 2026): enterprise platform unifying vLLM and SGLang under one gateway, adding a proprietary Flex Engine for non-LLM models (MinerU, PaddleOCR, selected ML models), automatic KV-cache/GPU-memory calculation from model structure, and project-based API key management.
- Hugging Face TGI: entered maintenance mode in December 2025 (bug fixes only) — effectively ceding the serving-engine race to vLLM and SGLang.
- ExecuTorch 1.0 GA (October 2025): Meta's on-device inference (Instagram, WhatsApp, Quest 3) — adjacent, not datacenter.
- Orchestration layer (Dynamo vs llm-d, Triton folded into Dynamo, Grove operator, AIBrix, KServe's LLMInferenceService): covered in sibling part 06a — this part keeps only the engine-level implication that the competitive surface has moved above the engines.
- What didn't survive 2026: Hugging Face TGI (maintenance mode) effectively conceded the race; ExecuTorch stays adjacent (on-device, not datacenter); the 2026 pattern is disaggregation everywhere — prefill/decode/encoder pools sized independently, KV blocks moved by NIXL — with llm-d orchestrating on Kubernetes (see 06a).

### Engine comparison matrix (September 2026) — the SGLang-relevant reading

