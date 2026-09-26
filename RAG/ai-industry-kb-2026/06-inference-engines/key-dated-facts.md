---
id: ai-industry-kb-2026/06-inference-engines/key-dated-facts
title: "Key dated facts"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "AWS", "DeepSeek", "Google", "Huawei", "Hugging Face", "Intel", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "SGLang", "Stripe", "TensorRT-LLM", "vLLM"]
dates: ["2025-10", "2026-01", "2026-01-22", "2026-03-20", "2026-05", "2026-06", "2026-08", "2026-08-17", "2026-08-29", "2026-09", "2026-09-22", "2026-10-20"]
keywords: ["accelerator", "agentic", "amd", "apache", "ascend", "aws", "blackwell", "capex", "compute", "cost", "deepseek", "disaggregated"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2212, 2240]
section: "6. Inference Engines"
sha256: 1d819c6b52e29efca4dd506f4d6b2e9ed67a3ce689b1fc9620d2603c702e2e32
---

# Key dated facts

**Heterogeneous clusters are a verified vLLM claim.** vLLM is the only open engine in 2026 with functional paths on NVIDIA (CUDA 13 default wheels, Blackwell SM120, H200, B200/GB200), AMD (ROCm — MI300X/MI325X/MI355 class; BF16 safetensors, MXFP4 MoE paths, FP8 KV cache on community forks), Google TPU (the `vllm-project/tpu-inference` plugin, formerly `vllm-tpu` — a JAX/XLA backend by Google; chunked prefill, KV offload, LoRA, TP/DP/EP/SP all supported, no MXFP8, strongest for offline/batch, weaker for interactive streaming), Intel (XPU/Gaudi plugins), Huawei Ascend (vllm-ascend), and **AWS Trainium/Inferentia** (the **vllm-neuron plugin, Beta**; NxD Inference — neuronx-distributed-inference — plugging into vLLM's Plugin System; Neuron SDK **2.32.0, 2026-08-17**; V1 API compatibility with expert parallelism for MoE, disaggregated inference architectures, Eagle V1 speculative decoding, INT8/FP8 quantization, prefix caching, multimodal Llama 4 Scout/Maverick; instances Inf2/Trn1/Trn2). Caveats: the Neuron plugin is still Beta; TPU interactive serving and Gemma 4 QAT checkpoints have documented gaps.

**Above the engine, 2026 created a commercial layer and an orchestration battle.** Engines converged on OpenAI-compatible APIs, XGrammar structured generation, and NIXL KV transfer — so proprietary value moved up: Inferact commercializes vLLM directly ($150M/$800M, January 2026), Red Hat/IBM sells a hardened vLLM distribution, NVIDIA packages vLLM in NIM 2.0 and competes on Dynamo orchestration, and RadixArk ($100M/$400M, May 2026) commercializes SGLang. The 2026 procurement picture is therefore **two commercial stacks** (vLLM/Red Hat/NVIDIA/Inferact vs SGLang/RadixArk) sitting on an increasingly interchangeable open-engine substrate, with the orchestration philosophies (Dynamo's vendor-integrated, performance-maximalist path vs llm-d's Kubernetes-native, hardware-neutral path) as the remaining strategic choice.

## Key dated facts

### The spend flip: inference overtakes training

- **2026-01-22 (context):** GPT-4-level inference cost reported at ~$0.40/M tokens vs ~$20/M in late 2022 [COMMUNITY]; inference's share of AI compute reported rising from one-third (2023) to two-thirds (2026); agentic workloads generate 5x–50x more tokens per interaction [COMMUNITY].
- **2026-03-20:** the "Training is dead, inference is the real AI war" discourse crystallizes the spend-flip narrative in industry blogs [COMMUNITY].
- **2026-08 (reported):** **Gartner: global AI-optimized IaaS spending surging 96% to $42B in 2026**; **inference workload spending ($23.3B) surpasses training ($19B)**; 55% of AI-optimized IaaS supports inference in 2026, forecast **59% in 2027** (Gartner principal analyst Hardeep Singh cites agentic AI multistep execution as the driver). Earlier October 2025 forecast had been $37.5B with 55% inference — the August 2026 update is an upward revision.
- **2026-08 (reported):** inference accounts for **60–80% of AI GPU spend for production teams**, projected at roughly two-thirds of all AI accelerator spending in 2026; the five largest hyperscalers on track for **$757.7B in 2026 capex (+60% YoY)**, the third straight year of 60%+ growth. Analyst framing: "Nvidia built its dominance on training... it is weakening in the segment where the majority of spending is now flowing."
- **2026-08:** why inference won — agentic multiplication (one user request becomes dozens of model calls), context-length growth (128K–1M contexts turn every long-context request into a prefill bill), test-time compute (reasoning models shift FLOPs from training to serving), and Jevons-style demand (collapsed unit cost bought more usage, not less) [COMMUNITY synthesis].
- **2026-09 (reported):** ISG: **65% of organizations piloting open-weight AI models; nearly 20% have deployed a local LLM** — the demand-side tailwind for open serving engines.

### vLLM commercialization and enterprise milestones

- **2026-01-22:** **Inferact launches** with a **$150M seed round at an $800M valuation** to commercialize vLLM; round co-led by **Andreessen Horowitz and Lightspeed Venture Partners**; formed by vLLM creators, CEO Simon Mo. The $150M figure is unusually large for a seed but consistently reported (TechCrunch, Pulse2, AI Insider).
- **2025-05 (announced, current through 2026):** Red Hat AI Inference Server — hardened vLLM distribution with Neural Magic compression; Red Hat claim: **"two to four times more token production with pre-optimized models"** [VENDOR] (Brian Stevens, SVP/AI CTO, Red Hat). AI Validated Models published on Hugging Face for compatibility-tested checkpoints. SiliconANGLE's read: Red Hat intends to shape the inference layer as it did containers/Kubernetes with OpenShift.
- **2026-08 (reported, third-party analysis):** **NIM 2.0 uses vLLM as its sole LLM/VLM inference backend** — "an upstream-first architectural shift from NIM 1.x": proprietary `nim-llm` orchestration + `nimlib` (model licensing, hardware-aware profile selection, health endpoints) over OSS vLLM (Apache 2.0). Backend map: **LLM NIM (v2.0) → vLLM (sole backend)**; **VLM NIM → vLLM 0.19 (sole backend)**; Embedding NIM → TensorRT + Triton; Speech/Biology NIM → Triton + custom backends; Edge NIM → TensorRT + Triton (no vLLM). Needs NVIDIA primary confirmation before being presented as NVIDIA's stated position.
- **2026-08-29:** Futurum Group analysis: vLLM is the **"de facto open-source LLM inference engine"**, framed at PyTorch Conference 2026 as production infrastructure and a strategic counterweight to proprietary serving. AI platforms market sized **$181.3B in 2026, growing to $496.9B by 2030 at 28.7% CAGR**.
- **2026-10-20–21 (post-cutoff):** PyTorch Conference NA 2026, San Jose — vLLM featured across KV cache, disaggregated serving, hardware portability, MoE inference and production deployment tracks, with contributors from Red Hat, IBM, NVIDIA, Mistral AI, Amazon, Huawei, Meta and Google. Only the program and advance analysis are verifiable as of 2026-09-22.
- **2026 (reported):** Stripe — **73% cost reduction serving 50M daily API calls on one-third the GPU fleet** via vLLM (the single most-cited enterprise proof point; secondary-reported via bbw9n).
- **2026 (community guides):** self-hosted vLLM serving cited at **$0.50–1.00 per million tokens** [COMMUNITY, directional].
- **2026-09 (reported):** 51% of enterprises pursue a balanced mix of in-house and vendor AI solutions; 63.9% deploy on provider-managed infrastructure — but seek portability that proprietary serving APIs do not provide [COMMUNITY].
- **2026 (reported):** one production homelab operator **parked their SGLang stack for DeepSeek V4 Flash in June 2026 and serves it via vLLM instead**, because "SGLang's strict-config loader rejects the NVFP4 quant" — while keeping SGLang for models llama.cpp can't serve. Engine choice in 2026 is per-model, not per-religion [COMMUNITY].

### vLLM release arc (February → September 2026)

