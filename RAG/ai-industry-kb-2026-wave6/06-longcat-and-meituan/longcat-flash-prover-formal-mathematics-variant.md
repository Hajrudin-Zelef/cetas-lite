---
id: ai-industry-kb-2026-wave6/06-longcat-and-meituan/longcat-flash-prover-formal-mathematics-variant
title: "LongCat-Flash-Prover — formal-mathematics variant"
domain: longcat-and-meituan
role: deep-dive
task: reference
actors: ["AMD", "China", "Huawei", "Hugging Face", "LongCat", "Meituan", "Nvidia", "OpenRouter", "SGLang"]
dates: ["2026-06-30", "2026-07-01", "2026-07-05"]
keywords: ["agent", "agentic", "amd", "ascend", "asic", "attention", "attribution", "cost", "decode", "disaggregated", "embedding", "gpu"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2797, 2823]
section: "§6. LongCat and Meituan"
delta_of: ai-industry-kb-2026
sha256: 7da86abe6ec1187aed538fefa72515ef6cf9bd44c4d3a4b25336ac13df367696
---

# LongCat-Flash-Prover — formal-mathematics variant

### LongCat-Flash-Prover — formal-mathematics variant
- **LongCat-Flash-Prover**: 560B-class **open-weight MoE specialized for Lean4** theorem proving [COMMUNITY] (github.com/brndngln/longcat-flash-prover).
- Training framework: **Hybrid-Experts Iteration Framework** [SECONDARY] (paper notes #4995).
- RL method: **HisPO** (history-aware policy optimization) with **sequence-level and token-level gradient masking** to counter policy staleness and train/inference mismatch [SECONDARY] (paper notes #4995).
- Anti-reward-hacking: **theorem consistency and legality checks** in the reward pipeline [SECONDARY] (paper notes #4995).
- Vendor results: **MiniF2F-Test 97.1% with 72 attempts; ProverBench 70.8%; PutnamBench 41.5% with ≤220 attempts** — all [VENDOR] (paper notes #4995).
- Implication: the Prover variant shows the LongCat program extending beyond agentic coding into formal mathematics — a second pillar alongside the coding flagship [DIRECTIONAL].

### LongCat-2.0 — architecture deep spec
- **Shortcut-connected MoE (ScMoE)**; dynamic active **33B–56B per token, average ~48B** of 1.6T total [SECONDARY] (marktechpost.com, 2026-07-05; cryptobriefing.com).
- **Zero-computation experts** with **PID-controller adjustment** of the dynamic expert budget [SECONDARY] (marktechpost.com, 2026-07-05; venturebeat.com).
- **LongCat Sparse Attention** with three indexing mechanisms: **Streaming-aware, Cross-Layer, and Hierarchical indexing** — the source of the native 1M-token context, cutting long-context cost from quadratic toward linear [SECONDARY] (marktechpost.com, 2026-07-05).
- **135B-parameter N-gram embedding module**: expands the core embedding space roughly **100×**, capturing dense local token relationships and accelerating large-batch inference by reducing memory I/O bottlenecks [SECONDARY] (venturebeat.com).
- Post-training: **MOPD (Multi-Teacher Optimization via Mixture of Specialized Experts)** — fuses three teacher groups instead of blending raw human feedback into one reward function: **Agent Experts** (structural execution, tool invocation, multi-turn API parameter parsing, self-correcting loops), **Reasoning Experts** (multi-hop logic, chain-of-thought engineering, math, STEM), **Interaction Experts** [VENDOR via secondary] (venturebeat.com, 2026-07).
- Systems stack: **six-dimensional parallelism, disaggregated prefill/decode, super-kernels, L2 weight prefetch** [SECONDARY] (marktechpost.com, 2026-07-05; venturebeat.com).

### LongCat-2.0 — training, hardware, and release mechanics
- Training-token report: **35T+ tokens** appears in secondary coverage; the vendor's own materials reviewed did not confirm the figure in the sources checked — record as **[VENDOR, unconfirmed in reviewed sources]** and do not present as established [DIRECTIONAL].
- Hardware: training and inference ran on a **50,000-card domestic AI ASIC cluster with no NVIDIA hardware** (no A100/H100, no AMD MI300X); Meituan claims it is a **first of its scale for domestic Chinese hardware** [VENDOR via secondary] (cryptobriefing.com; marktechpost.com, 2026-07-05).
- Hardware-identity conflict: one French secondary source explicitly names **Huawei Ascend 910B**, but stronger sources say only "domestic chips/ASICs" — **do not promote "Ascend 910B" beyond a single-source [UNVERIFIED] attribution** [DIRECTIONAL] (ayinedjimi-consultants.fr PDF vs cryptobriefing.com/marktechpost.com).
- Unveiled **June 30, 2026** (cryptobriefing "reveals" piece); weights release **July 1, 2026** [SECONDARY] (cryptobriefing.com).
- License: **MIT** [SECONDARY] (marktechpost.com, 2026-07-05).
- Weights on Hugging Face under the **`meituan-longcat`** organization; GitHub repo **`meituan-longcat/LongCat-2.0`** [VENDOR] (github.com/meituan-longcat/LongCat-2.0).
- Chat access: official website **https://longcat.ai/** [VENDOR] (LongCat-2.0 GitHub README).
- Deployment: **GPU via SGLang cookbook; NPU via SGLang-FluentLLM** — the NPU path is a first-class vendor-supported target, consistent with the domestic-chip training story [VENDOR] (LongCat-2.0 GitHub README).
- The model had been **leading OpenRouter** (by usage) ahead of the open release — community traction preceded the weights drop [SECONDARY] (venturebeat.com headline reporting).

