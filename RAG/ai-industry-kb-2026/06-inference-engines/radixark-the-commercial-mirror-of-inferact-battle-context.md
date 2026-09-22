---
id: ai-industry-kb-2026/06-inference-engines/radixark-the-commercial-mirror-of-inferact-battle-context
title: "RadixArk: the commercial mirror of Inferact (battle context)"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "AWS", "Alibaba", "Apple", "Baseten", "Broadcom", "Cohere", "CoreWeave", "DeepSeek", "Google", "Groq", "Hugging Face", "Intel", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nebius", "Nvidia", "OpenAI", "Oracle", "SGLang", "Stripe", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM", "xAI"]
dates: ["2025-12-15", "2026-01", "2026-01-22", "2026-01-31", "2026-02-02", "2026-02-28", "2026-03-02", "2026-03-28", "2026-04-24", "2026-04-28", "2026-05-04", "2026-05-05", "2026-05-07", "2026-06-05", "2026-06-12", "2026-07-12", "2026-07-14", "2026-07-24", "2026-08-05", "2026-08-13", "2026-08-17", "2026-08-20", "2026-08-26", "2026-08-29", "2026-09", "2026-09-01", "2026-09-04", "2026-09-05", "2026-09-09", "2026-09-10", "2026-09-12", "2026-09-17", "2026-09-22"]
keywords: ["accelerator", "acquisition", "agent", "agentic", "agents", "amd", "apache", "attention", "attribution", "aws", "benchmark", "benchmarks"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2414, 2529]
section: "6. Inference Engines"
sha256: e0bcdb2e9f6871e17b169aaf4a9e2ac78c2a6862b861b11a41dbfd0b7df3b601
---

# RadixArk: the commercial mirror of Inferact (battle context)

### RadixArk: the commercial mirror of Inferact (battle context)

- **2026-05-05: RadixArk formally launches** — SGLang's commercial spinoff — with a **$100M seed at $400M post-money valuation**: led by Accel, co-led by Spark Capital; participants NVentures (NVIDIA's VC arm), AMD, MediaTek, Salience Capital, HOF Capital, Walden Catalyst, LDV Partners, WTT Investment; angels include the CEOs of Intel (Lip-Bu Tan) and Broadcom (Hock Tan), xAI co-founder Igor Babuschkin, OpenAI co-founder John Schulman, PyTorch creator Soumith Chintala, Hugging Face co-founder Thomas Wolf, Datadog co-founder Olivier Pomel, Anyscale co-founder Robert Nishihara. Founders: **Ying Sheng and Banghua Zhu** (AI infrastructure veterans from xAI and NVIDIA). The $400M figure was already in press in January 2026 (TechCrunch's 2026-01-22 Inferact piece) — **the "January launch" framing is wrong; formal launch = 2026-05-05**.
- SGLang remains Apache 2.0; RadixArk will "steward SGLang" while building a commercial end-to-end platform (training, fine-tuning, RL, inference at scale); introduced **Miles**, its own large-scale RL framework. Pitch language: "hundreds of thousands of GPUs worldwide", "trillions of tokens daily" [UNVERIFIED].
- **Why it matters here:** RadixArk narrows the enterprise-packaging gap Wave 2 identified as vLLM's structural advantage — the consolidation should reframe the 2026 commercial picture as **two commercial stacks** (Red Hat/IBM/NVIDIA vLLM vs RadixArk SGLang), though as of September 2026 RadixArk is still early relative to Red Hat's shipped enterprise server. (Full SGLang track → 06b.)

### Validation discipline: what the 2026 evidence does and does not support

- **No GPU execution was performed in any research wave** — all benchmark figures are cited with attribution, not reproduced; community benchmark figures (+29% SGLang vs vLLM, 6.4x RAG, RadixAttention token savings, Unsloth UD within 0.8 pt of FP weights) carry methodologies not independently audited.
- **Version numbers rot within weeks:** vLLM shipped roughly a minor per month; SGLang v0.5.18 → v0.5.19 in two weeks. Every "latest version" statement in this part is date-stamped; re-check against upstream release notes before operational decisions.
- **The engines trade wins by workload, not quality:** SGLang +29% on prefix-heavy load (same-kernel control test); vLLM wins fastest TTFT and 100+ concurrency on GPT-OSS-120B tests; an independent Aug 2026 workstation-Blackwell measurement contradicted published TRT-LLM deltas. "Universal benchmark winner" claims are unsupported by design.
- **Vendor marketing ratios to keep attributed:** Dynamo 7x throughput-per-GPU / 2x TTFT / −80% SLA violations (NVIDIA); XGrammar "80x" (MLC-adjacent); Red Hat 2–4x token production; llm-d 57x TTFT (project); Kimi-K3 release-note figures (1.5–3x kernel speedups, ~60% DSpark TTFT, ~17 GiB/GPU savings); vLLM-Omni 91.4% JCT (paper); Stripe 73% (secondary-reported).

## Figures and metrics

- **Gartner (Aug 2026):** AI-optimized IaaS **$42B in 2026 (+96% YoY)**; inference **$23.3B** vs training **$19B**; 55% inference share in 2026 → **59% in 2027**.
- **Gartner (Oct 2025 vintage):** $37.5B AI-optimized IaaS forecast for 2026, 55% inference — upward revision within ten months.
- **Production-team economics (Aug 2026 analysis):** inference **60–80% of AI GPU spend** for production teams; roughly **two-thirds of all AI accelerator spending** in 2026.
- **Hyperscaler capex:** **$757.7B in 2026 (+60% YoY)** — third straight year of 60%+ growth.
- **GPT-4-level inference unit cost:** ~$20/M tokens (late 2022) → ~$0.40/M tokens (early 2026) [COMMUNITY].
- **Stripe (reported 2026):** **73% cost reduction**, 50M daily API calls, one-third the GPU fleet via vLLM.
- **Self-hosted vLLM cost:** **$0.50–1.00 per million tokens** [COMMUNITY guides, directional].
- **Decision math:** with GPT-4-level inference at ~$0.40/M tokens and agentic workloads multiplying tokens 5–50x per interaction, a 29% throughput delta (SGLang on prefix-heavy load, sibling 06b) or a fleet reduction (Stripe's 73%) translates into seven-figure annual differences for high-volume deployments — which is why 2026 procurement treats the inference engine as infrastructure, not tooling [COMMUNITY synthesis].
- **Disaggregation dividend:** **4–10x cost-per-token reduction** vs aggregated serving [COMMUNITY practitioner analysis]; Dynamo-claimed 7x throughput-per-GPU on DeepSeek R1/GB200 [VENDOR].
- **Red Hat AI Inference Server:** **2–4x token production** with pre-optimized models [VENDOR].
- **llm-d v0.5 benchmarks:** up to **57x faster TTFT, 2x throughput** vs round-robin under high prefix reuse (8 pods/16×H100) [VENDOR]; operators' own measurements: ~25% over defaults, 2–3x tokens/s/GPU with prefix-cache-hit routing, 3–5x cost-per-token on chat-shaped workloads [COMMUNITY].
- **vLLM release cadence:** roughly one minor version per month in 2026 (v0.15.1 → v0.29.0 in eight months). v0.20.0: 752 commits/320 contributors; v0.25.0: 558 commits/232 contributors; v0.26.0: 411 commits/212 contributors; v0.28.0: 584 commits/270 contributors (76 new).
- **GitHub scale:** vLLM ~50K–75K stars (sources disagree; present as range), ~2,400 contributors, 100+ model architectures, Apache 2.0; issue response 12 hours–3 days [COMMUNITY].
- **Kimi-K3 (v0.28.0 release notes) [VENDOR]:** combined all-gathers 1.5–3x kernel-level speedup; adaptive speculative token budget ~60% better DSpark TTFT; shared-expert sharding ~17 GiB/GPU memory savings.
- **vLLM-Omni paper (arXiv:2602.02204):** up to **91.4% lower JCT** vs baselines [UNVERIFIED as independently reproduced].
- **vLLM-Omni diffusion wall-clock (H200, Wan2.2-I2V) [COMMUNITY]:** 133.94 s (v0.16.0 retro) → 93.67 s (v0.18.0) → 79.19 s (v0.20.0).
- **New defaults (v0.28.0):** `max_num_batched_tokens` 8192 → 16384; Blackwell CUDA-graph capture default raised to 1024.
- **Independent TRT-LLM counter-benchmark (Aug 2026, Qwen3-Coder-30B-A3B, workstation Blackwell) [COMMUNITY]:** vLLM led at concurrency 1 (12.3 vs 10.8), 8 (56.3 vs 44.4), 64 (129.0 vs 115.1); TRT-LLM led only at 32 (120.5 vs 119.1, +1%). Peak: vLLM 129.0 vs TRT-LLM 120.5. Published 10–25% TRT-LLM claims did not hold on workstation Blackwell (likely because sm_120 lacks FlashAttention-4); vLLM kept climbing to 64 concurrent while TRT-LLM saturated at 32.
- **Head-to-head counter-evidence on the vLLM side (AIMultiple/GPT-OSS-120B tests) [COMMUNITY]:** vLLM wins **fastest TTFT** across concurrency levels in GPT-OSS-120B tests and leads at **100+ concurrent requests** — the workload-dependent correction to any "SGLang always faster" reading (SGLang's +29% holds on prefix-heavy load, tested with identical FlashInfer kernels, i.e. the difference is orchestration overhead, not kernel performance).
- **AI platforms market:** $181.3B (2026) → $496.9B (2030), 28.7% CAGR [COMMUNITY].
- **The 1M-context serving proof:** GLM-5.3-Flash NVFP4 at 1M-token context on DGX Spark is the longest-context vLLM deployment documented this cycle [COMMUNITY]; the SM121 kernel patch it required shows community patching still carries part of the day-0 burden.
- **Enterprise adoption (PyTorch Conference 2026/Futurum) [COMMUNITY]:** 51% of enterprises pursue balanced in-house/vendor AI; 63.9% deploy on provider-managed infrastructure.
- **FlagOS chip coverage:** vLLM's multi-hardware effort tracked against a **20+ chip test base** — the concrete mechanism behind the "any accelerator" claim [COMMUNITY].
- **Agentic token multiplication:** one user request becomes dozens of model calls; agents generate 5x–50x more tokens per interaction than single-turn chat [COMMUNITY] — the workload shape that made prefix-cache efficiency and structured tool-call generation the 2026 differentiators rather than raw prefill FLOPS.
- **Context-length production costs:** 128K–1M token contexts moved from flagship demos to production RAG and code agents; every long-context request is a prefill bill paid at inference time — a structural driver of the spend flip alongside test-time compute.
- **The MLA compounding pattern (2026):** MLA compresses the KV cache, FP8/FP4 quantizes it, disaggregation moves it, and RadixAttention shares it — four independent multipliers on cost-per-token [COMMUNITY synthesis].
- **New entrants (2026, adjacent):** TileRT (per-user decode-speed runtime pairing with vLLM prefill); Tiny-vLLM (community C++/CUDA lightweight engine, Hacker News traction); Neutree 1.2 / Flex Engine (Arcfra, September 2026 — enterprise platform unifying vLLM and SGLang under one gateway, proprietary Flex Engine for non-LLM models, automatic KV-cache/GPU-memory calculation, project-based API keys).

## Main actors

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

| Date | Event |
|---|---|
| 2025-03 | NVIDIA Dynamo launched at GTC 2025 — inference orchestration (disaggregated P/D, KV-aware routing) |
| 2025-12 | Hugging Face TGI enters maintenance mode (bug fixes only) — serving race ceded to vLLM/SGLang |
| 2025-12 | DeepSeek-V3.2 introduces DeepSeek Sparse Attention (lightning indexer, top-K 2048) — the DSA lineage vLLM serves in 2026 |
| 2025-12-15 | llama.cpp model router announced (multi-process, LRU eviction) |
| 2025-05 | Red Hat AI Inference Server announced — hardened vLLM + Neural Magic compression |
| 2025-11 | vLLM-Omni officially released under the vLLM community (any-to-any multimodal serving) |
| 2026-01-22 | **Inferact launches**: $150M seed at $800M valuation (a16z + Lightspeed), formed by vLLM creators, CEO Simon Mo, to commercialize vLLM |
| 2026-01-31 | vLLM-Omni v0.14.0 — first stable Wan2.2 diffusion pipelines |
| 2026-02 | vLLM v0.15.1 — PyTorch 2.10, RTX Blackwell (SM120), H200 optimization |
| 2026-02-02 | vLLM-Omni paper arXiv:2602.02204 — up to 91.4% lower JCT claimed [UNVERIFIED] |
| 2026-02-28 | vLLM-Omni v0.16.0 — `/v1/videos` API |
| 2026-03 | **vLLM v0.16.0** — V1 engine lockdown (legacy V0 core removed); `--use-v1=0`-style flags dead |
| 2026-03-02 | Forensic code review confirms V0 removal in v0.16.0 at binary level |
| 2026-03-28 | vLLM-Omni v0.18.0 — IPC −17.5% on Wan2.2 |
| 2026-04-24 | DeepSeek V4 (Pro/Flash + bases): 1.6T/49B and 284B/13B, 1M context, MIT license, DSA-derived CSA+HCA — vLLM's v0.20.0/v0.26.0 sparse-MLA target |
| 2026-04-28 | AWS DLC v1.1.0 — LMCache bidirectional NIXL cache probe for P/D disaggregation (vLLM 0.19.1) |
| 2026-06 | GLM-5.2: IndexShare (shared sparse-attention indexer every 4 layers, 2.9× FLOP cut at 1M), 1M context — API ~06-13, weights ~06-17 [VENDOR] |
| 2026-06-12 | NVIDIA Nemotron-3-Ultra-550B-A55B: Mamba-2 + MoE + Attention hybrid ("LatentMoE"), MTP, 1M context, NVFP4 — vLLM serves SSM-hybrid layers with prefix caching for Mamba (v0.25.0) |
| 2026-07 | TensorRT-LLM **v1.3.0rc21–rc23**: EAGLE3 dynamic-tree kernels, fused RMSNorm/RoPE, **AutoDeploy backend deprecated** |
| 2026-05-04 | XGrammar-2 released — Structural Tag protocol (agent structured generation) |
| 2026-05-05 | **RadixArk formal launch**: $100M seed at $400M post-money (Accel-led, Spark Capital co-lead; NVentures, AMD, MediaTek) — SGLang commercial steward |
| 2026-05-07 | vLLM-Omni v0.20.0 — fused DiT + CI JSON |
| 2026-07 | RadixArk expands Google TPU partnership (SGLang-JAX as the commercial vehicle) |
| 2026-06 | **vLLM v0.20.0** — CUDA 13.0 default, PyTorch 2.11, Transformers v5, DeepSeek V4 initial support, FA4 default MLA prefill, TurboQuant 2-bit KV, online quantization frontend |
| 2026-06-05 | Distributed inference best-practices post documents the KV connector landscape |
| 2026-06 | One operator parks SGLang for DeepSeek V4 Flash, serves via vLLM (NVFP4 loader issue) — per-model engine choice |
| 2026-07-12 | Community analysis: "vLLM 0.25 Deletes PagedAttention" (kenashe) |
| 2026-07 | **vLLM v0.24.0** — in-tree GGUF removed → `vllm-gguf-plugin`; **v0.25.0** — PagedAttention deleted, Model Runner V2 default, Transformers backend parity, Streaming Parser Engine, universal speculative decoding (TLI) |
| 2026-07-14 | vLLM blog: vLLM prefill + TileRT decode production pairing |
| ~2026-07 | **NVIDIA Dynamo 1.0** ships — disaggregated prefill/decode orchestration open-sourced |
| 2026-08 | vLLM-Omni v0.28.0 line — MiniMax H3 production-ready, unified AR/DiT paged KV runtime, full-duplex MiniCPM-o; VeRL-Omni v0.2.0 |
| 2026-08 (reported) | Independent benchmark: vLLM beats TensorRT-LLM at peak concurrency on workstation Blackwell (Qwen3-Coder-30B-A3B) |
| 2026-08 (reported) | **Gartner: inference $23.3B > training $19B** in 2026 AI-optimized IaaS ($42B, +96%) |
| 2026-08 (reported) | **NIM 2.0: vLLM sole LLM/VLM backend** (third-party analysis) |
| 2026-08-17 | AWS Neuron SDK 2.32.0 — vLLM V1 API compatibility on Trainium/Inferentia |
| 2026-08-05 | TensorRT-LLM **day-0 support for OpenAI GPT-OSS-120B/20B** (v1.3.0rc line) |
| 2026-08-13 | DeepSeek V4-Pro GA checkpoint (V4-Pro-0813); legacy `deepseek-chat`/`deepseek-reasoner` aliases retired 2026-07-24 [COMMUNITY] |
| **2026-08-26** | **vLLM v0.28.0 tagged** (core release): Kimi-K3 push, DeepSeek V4 end-to-end sparse MLA, Model Runner V2 E/P/D disaggregation, tiered KV offload with disk tier |
| 2026-08-29 | Futurum Group: vLLM is the "de facto open-source LLM inference engine" |
| 2026-09-01 | First community production upgrade to vLLM v0.28.0 recorded |
| 2026-09-04 | SGLang v0.5.19 on PyPI (786 PRs) — the engines race to serve the same late-2026 flagships (→ 06b) |
| 2026-09-05 | llama.cpp v0.4.0 (ggml v0.23.0: sparse flash attention, Apple RDMA transport) — GGUF reference runtime |
| 2026-09-09 | Community vLLM v0.29.0 production attempt recorded [DIRECTIONAL] |
| 2026-09-10 | Unsloth v0.1.808-beta — diffusion 1.2–1.7x (INT8/FP8), AMD via Vulkan +20%, PyTorch 2.11 |
| 2026-09-12/13 | Working ROCm vLLM 0.27.0 + GGUF plugin report (operator: safetensors recommended for reliability) |
| 2026-09-17 | Unsloth official changelog — multi-user Docker, AMD RDNA1/2, FP8/INT8 diffusion 2x claimed, ARM64 Windows CUDA, GRPO Qwen3.5 + latest TRL/vLLM |
| 2026-09 (month) | ISG: 65% piloting open-weight AI, ~20% with local LLM deployed |
| 2026-09-22 | Research date for all waves |

