---
id: ai-industry-kb-2026/06-inference-engines/timeline-and-context
title: "Timeline and context"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "AWS", "Baseten", "Cohere", "CoreWeave", "DeepSeek", "Google", "Groq", "Hugging Face", "Meta", "Microsoft", "MiniMax", "Mistral", "Nebius", "Nvidia", "Oracle", "SGLang", "Stripe", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM"]
dates: ["2025-12-15", "2026-01-22", "2026-04-28", "2026-06-05", "2026-07-12", "2026-07-14", "2026-08-17", "2026-08-20", "2026-09-05", "2026-09-12", "2026-09-17"]
keywords: ["acquisition", "apache", "aws", "benchmark", "cohere", "cost", "decode", "deepseek", "diffusion", "disaggregated", "fp8", "gguf"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2459, 2478]
section: "6. Inference Engines"
sha256: b513f34a0575dccd07ffde59af4941b6db0d6a6df77f232531c550ad72451e51
---

# Timeline and context

- **vLLM project** (vllm-project; Apache 2.0): the engine itself — ~50K–75K GitHub stars, ~2,400 contributors, 100+ model architectures; roughly one minor release per month in 2026; V1 lockdown (v0.16.0), PagedAttention removal (v0.25.0), Model Runner V2 default (v0.25.0), GGUF split to out-of-tree plugin (v0.24.0). Reference engine for llm-d's cache-aware routing and NIXL KV transfer.
- **Inferact:** formed by vLLM creators, CEO **Simon Mo**; launched **2026-01-22** with **$150M seed at $800M valuation** (a16z + Lightspeed co-led) to commercialize vLLM — opens the 2026 "inference commercialization" storyline. Unusually large seed, consistently reported.
- **Red Hat / IBM:** AI Inference Server — hardened vLLM + Neural Magic compression (Neural Magic acquired by Red Hat), 2–4x token-production claim [VENDOR]; AI Validated Models on Hugging Face; llm-d co-led with Google Cloud/IBM/NVIDIA/CoreWeave; K8s-native, hardware-neutral orchestration philosophy.
- **NVIDIA:** the pivot that validates vLLM's standard status — **NIM 2.0 uses vLLM as sole LLM/VLM backend** (third-party analysis; upstream-first shift from NIM 1.x; needs NVIDIA primary confirmation); **Dynamo** (1.0, mid-2026) as the vendor-integrated orchestration layer with 7x throughput-per-GPU vendor claims; folds Triton into Dynamo ("Dynamo-Triton"); **NIXL** as the transfer fabric shared by all players; NVIDIA contributed to llm-d. Rubin platform (full production, shipping H2 2026) integrates a post-Groq-acquisition inference processor (prefill→Rubin GPU / decode→Groq 3 LPU via Dynamo) [COMMUNITY].
- **AWS:** Deep Learning Containers ship **vLLM-Omni images** (CUDA 13.0, PyTorch, NCCL, Python 3.12) bundled with FlashInfer, DeepEP, LMCache + NIXL, runai-model-streamer, EFA/OpenMPI, espeak-ng/ffmpeg; DLC v1.1.0 (2026-04-28) bundled LMCache 0.4.5.dev0 for bidirectional P/D cache probing; `awslabs/ai-on-eks` Dynamo blueprints for disaggregated vLLM/SGLang/TRT-LLM. **Neuron team**: vllm-neuron plugin (Beta) + NxD Inference, V1 API compatibility, Neuron SDK 2.32.0 (2026-08-17), Inf2/Trn1/Trn2.
- **Google:** `vllm-project/tpu-inference` — the JAX/XLA TPU plugin (formerly `vllm-tpu`); machine-readable support matrices updated 2026-08-20/27; TPU generations v7x (Ironwood)/v6e (Trillium)/v5e recommended.
- **Mistral AI, Cohere, Meta, Amazon (Rufus), Anyscale, Roblox:** named 2026 flagship production users of vLLM [COMMUNITY surveys].
- **Microsoft Azure / Oracle Cloud / Google Cloud / AWS / Nebius:** named cloud homes for vLLM and SGLang deployments [COMMUNITY].
- **Enterprise buyers:** Stripe (73% cost reduction, 50M daily API calls); regulated-industry buyers choosing vLLM via Red Hat's support SLAs and validated model catalogs.
- **TileRT:** new decode-focused runtime; official vLLM blog pairing (2026-07-14) — vLLM owns prefill/scheduling/APIs, TileRT owns decode for latency-bound workloads. Thesis: per-user decode speed as its own scaling dimension.
- **Community validators:** terrytangyuan (KV-connector production guidance, 2026-06-05); kenashe ("vLLM 0.25 Deletes PagedAttention", 2026-07-12); kender242 (forensic V0-removal review of v0.16.0); marianvid (independent Aug-2026 vLLM vs TensorRT-LLM benchmark); operators documenting ROCm 0.27.0 + GGUF plugin builds (2026-09-12/13); hsliuustc0106 (vLLM-Omni diffusion cookbook archaeology).
- **DeepEP (expert parallelism kernels):** high-throughput, low-latency all-to-all dispatch/combine with FP8 support (V2 on NCCL GIN); UCCL-EP extends GPU-driven EP on heterogeneous hardware; bundled with AWS vLLM-Omni images alongside FlashInfer for large-MoE serving (DeepSeek V3.2/V4, MiniMax M3, GLM-5).
- **Adjacent engines (cross-referenced, not covered here):** SGLang / RadixArk → 06b; TensorRT-LLM → 06c sibling; llama.cpp (v0.4.0, 2026-09-05; model router announced 2025-12-15; GGUF reference runtime); Unsloth (Dynamic v3.0 GGUF, Desktop/Studio platform, 2026-09-17 release); TileRT, Tiny-vLLM (Hacker News traction 2026), Arcfra Neutree 1.2/Flex Engine (2026-09), ExecuTorch 1.0 GA (Oct 2025, on-device), Hugging Face TGI in maintenance mode since 2025-12.
- **Standard-setting backdrop:** llm-d (Kubernetes-native, hardware-neutral), Dynamo (NVIDIA-integrated, performance-maximalist), KServe (frames Dynamo as an alternative backend to its llm-d-based `LLMInferenceService`), Grove (Dynamo's K8s operator), AIBrix (vLLM project's RoleSets), vLLM production-stack (Helm + KV-aware routing).
- **Neural Magic (acquired by Red Hat):** compression technologies behind the Red Hat AI Inference Server's 2–4x token-production claim [VENDOR]; built-in LLM compression tools that shrink base and fine-tuned models "while preserving accuracy."
- **Futurum Group:** the analysis house that framed vLLM as "production infrastructure" at PyTorch Conference 2026 — the enterprise evidence layer (economic crossover → production-infrastructure status → demand growth → commercial packaging) now reads end-to-end.
- **CoreWeave / Nebius / DataCrunch / Voltage Park / Baseten / RunPod / Novita:** 2026 GPU-cloud and serverless names appearing in SGLang adoption surveys — the rental-market side of the engine battle (GPU pricing mechanics → §14).

## Timeline and context

