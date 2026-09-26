---
id: ai-industry-kb-2026/06-inference-engines/main-actors-continued
title: "Main actors (continued)"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "AWS", "Alibaba", "Baseten", "DeepSeek", "Google", "Hugging Face", "Intel", "LongCat", "Meta", "Microsoft", "Moonshot", "Nebius", "Nvidia", "OpenAI", "Oracle", "SGLang", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM", "xAI"]
dates: ["2025-03", "2025-10-29", "2025-12-15", "2026-01-22", "2026-03-15", "2026-05-04", "2026-05-05", "2026-07", "2026-07-14", "2026-08-05", "2026-08-20", "2026-09-04", "2026-09-05", "2026-09-20", "2026-09-22"]
keywords: ["amd", "apache", "attention", "aws", "decode", "deepseek", "diffusion", "disaggregated", "distribution", "fp4", "fp8", "gguf"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2882, 2941]
section: "6. Inference Engines"
sha256: f055d322050eb76be5153d3221932f7d66589c7808338bec96184707da7fa200
---

# Main actors (continued)

## Main actors (continued)

### SGLang and RadixArk

- SGLang team (LMSYS/UC Berkeley, Chatbot Arena lineage); RadixArk founders Ying Sheng and Banghua Zhu (xAI and NVIDIA backgrounds); SGLang continuity as Apache 2.0.
- RadixArk investors: Accel (lead), Spark Capital (co-lead); NVentures (NVIDIA), AMD, MediaTek, Salience Capital, A&E Investments, HOF Capital, Walden Catalyst Ventures, LDV Partners, WTT Investment.
- RadixArk angels: Igor Babuschkin, Lip-Bu Tan, Hock Tan, John Schulman, Soumith Chintala, Olivier Pomel, Thomas Wolf, William Fedus, Robert Nishihara, Eric Zelikman, Logan Kilpatrick.
- a16z OSS AI Grant funder of SGLang; SGLang in the PyTorch ecosystem since March 2025.

### Frontier adopters

- xAI (Grok 3 on SGLang), Microsoft Azure (DeepSeek R1 on AMD), Oracle Cloud, Google Cloud, AWS, Nebius, DataCrunch, Voltage Park, Cursor (code completion), LinkedIn (AI features); NVIDIA, AMD, Intel as both vendors and users; Baseten, RunPod, Novita.
- Academia: Stanford, UC Berkeley, UCLA, MIT, UW, Tsinghua.
- RL stack: verl, AReaL, Miles (RadixArk's own framework), slime, Tunix; SGLang's OME Kubernetes operator.

### Adjacent engine actors

- MLC (XGrammar-2 authors); xAI, Databricks, DeepSeek (XGrammar-2 adopters).
- Arcfra (Neutree 1.2 / Flex Engine); TileRT; Tiny-vLLM community; Hugging Face (TGI maintenance); Meta (ExecuTorch 1.0 GA).
- Ant Group (MOVA cookbook for SGLang Diffusion); Google (SGLang-JAX TPU partnership, July 2026 expansion).
- Kernel and fabric vendors inside SGLang's stack: FlashInfer (BatchPrefill/BatchDecode, sparse-MLA kernels), DeepEP (MoE all-to-all), MORI on AMD, NIXL/Mooncake (P/D transfer) — the 2026 engine is as much a composition of specialized kernels as a codebase.
- Project tooling: sgl-router (cache-aware routing), sgl-kernel (ops), OME (K8s operator), sglang-omni (any-to-any track) — the SGLang surface extends well beyond the core server.

### Dedup pointers for this part

- vLLM core internals, vllm-gguf-plugin, vLLM-Omni, Inferact, NIM 2.0, Dynamo vs llm-d, spend flip, per-token economics → sibling part 06a (one-line cross-refs only in this part).
- KV-cache techniques (radix-tree deep mechanics, paged-KV internals, tiered offload, FP8 KV dtype, APC mechanics) → §7 authoritative; this part retains only engine-level outcome figures.
- TPU hardware specifics → §15; Unsloth platform internals → the Unsloth-track source (§4 of wave 1) where needed.
## Timeline and context (continued)

| Date | Event |
|---|---|
| 2025-08 | RadixArk first announced as a startup (~$400M valuation, Accel-led round reported) — precursor to the formal launch |
| 2025-10-29 | SGLang-JAX announced (native JAX/XLA TPU engine) |
| 2025-12 | DeepSeek-V3.2 introduces DeepSeek Sparse Attention (lightning indexer, top-K 2048); Hugging Face TGI enters maintenance mode |
| 2025-12-15 | llama.cpp model router announced (multi-process, LRU eviction, OpenAI-compatible dynamic loading) |
| 2026-01-22 | TechCrunch reports SGLang commercialization "seeking a $400M valuation" — the first $400M press datapoint; Inferact launches for vLLM the same day ($150M seed, $800M valuation — see 06a) |
| 2026-02 | SGLang blog: "Unlocking 25x Inference Performance with SGLang on NVIDIA GB300 NVL72" [VENDOR headline, baseline unaudited] |
| 2026-03-15 | Enterprise comparison published: vLLM vs SGLang for production (privocto) — 4.7x multi-turn speedup, 99.8% JSON validity [DIRECTIONAL] |
| 2026-04 | SGLang day-zero support for DeepSeek-V4 |
| 2026-05-04 | XGrammar-2 released — Structural Tag protocol; adopted by xAI, Databricks, DeepSeek |
| 2026-05-05 | **RadixArk formal launch**: $100M seed at $400M post-money, led by Accel, co-led by Spark Capital; NVentures, AMD, MediaTek participate |
| 2026-06 | SGLang FlashInfer SM120 sparse-MLA decode PR #27455 (~2.2–3.7x TPOT); Nemotron and Higgs Audio added to SGLang; DeepSeek V4 day-zero GLM-5.2 IndexShare ships elsewhere as sparse-attention context |
| 2026-06 | One operator parks SGLang for DeepSeek V4 Flash (NVFP4 loader issue), serves via vLLM — per-model engine choice |
| 2026-07 | SGLang day-zero support for Kimi K3; RadixArk expands Google TPU partnership (SGLang-JAX as commercial vehicle) |
| 2026-07-14 | vLLM blog announces vLLM prefill + TileRT decode pairing |
| 2026-07 | TensorRT-LLM v1.3.0rc21–rc23; AutoDeploy backend deprecated (rc21) |
| 2026-08-05 | TensorRT-LLM day-0 support for OpenAI GPT-OSS-120B/20B |
| 2026-08-20 | SGLang v0.5.18 tagged |
| 2026-09-04 | **SGLang v0.5.19 on PyPI** (786 PRs, 214 contributors): Qwen3.8, dots3.note, Ling-3.0, Spark2.5, MiniCPM-SALA, Granite 4.2, LongCat-Image-Edit |
| ~2026-09-05 | llama.cpp server v0.4.0 |
| 2026-09 | Neutree 1.2 / Flex Engine (Arcfra): vLLM + SGLang under one gateway |
| 2026-09-20 | Architecture review: up to 2.5x throughput gain under strict JSON constraints (jump-forward decoding) — the origin of the misread "2.5x cache hit rate" |
| 2026-09-22 | Research cutoff for all wave documents |

- Context arc: SGLang spent 2026 converting its 2025 scheduler advantages (RadixAttention, zero-overhead scheduling) into a full platform story — diffusion serving, a K8s operator (OME), a native TPU engine (SGLang-JAX), and a funded commercial steward (RadixArk) — while TensorRT-LLM converged onto the same disaggregated/NIXL patterns as the OSS engines and its commercial center of gravity shifted toward NIM packaging of non-LLM workloads.
- The two-engine convergence pattern of 2026: each engine absorbs the other's differentiators (SGLang adds FP4/quantized KV, multi-hardware breadth, K8s ops; vLLM absorbs disaggregation, diffusion, decode-speed — see 06a) while deepening its own moat — scheduler/prefix-affinity for SGLang, enterprise distribution for vLLM.

## Implications (continued)

