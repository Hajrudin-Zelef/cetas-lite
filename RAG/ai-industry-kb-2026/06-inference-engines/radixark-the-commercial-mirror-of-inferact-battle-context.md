---
id: ai-industry-kb-2026/06-inference-engines/radixark-the-commercial-mirror-of-inferact-battle-context
title: "RadixArk: the commercial mirror of Inferact (battle context)"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "AWS", "Broadcom", "Cohere", "Hugging Face", "Intel", "Meta", "Mistral", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Stripe", "Unsloth", "vLLM", "xAI"]
dates: ["2026-01", "2026-01-22", "2026-05-05", "2026-09"]
keywords: ["amd", "apache", "attribution", "benchmark", "blackwell", "cohere", "distribution", "fine-tuning", "gpu", "gpus", "inference", "intel"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2406, 2428]
section: "6. Inference Engines"
sha256: 08c9689a994841555f3b35caf904621dc069e814201269baab874b919b2a5334
---

# RadixArk: the commercial mirror of Inferact (battle context)

- **vLLM production-stack** (vLLM project): Helm-based Kubernetes deployment with a request router, multi-instance management, and KV-aware routing.
- **AIBrix** (vLLM project): StormService + RoleSet CRDs for high-density LoRA serving, gateway, autoscaling, and P/D disaggregation.
- **KServe** frames Dynamo as an *alternative backend* to its llm-d-based `LLMInferenceService`: Dynamo brings NIXL RDMA transfer, SGLang/TRT-LLM backends, and 1.0-level maturity; llm-d brings Gateway API standardization and K8s-native RBAC/multitenancy (alpha-stage in 2026). Positioning matters: "not replacement."
- **New K8s deployment decision rule (2026 practitioner consensus):** NVIDIA fleet → Dynamo (SLO planner, ModelExpress, AIConfigurator). Multi-hardware or K8s-standard fleet → llm-d (Gateway API, TPU/XPU neutrality).
- **Migration reality:** both vLLM and SGLang implement OpenAI-compatible APIs — client-side migration is near-zero effort; server-side is medium effort (CLI flags, middleware, metrics names, LoRA APIs). Practitioner recommendation: **start with vLLM, abstract the backend behind a gateway from day one** so it stays swappable.
- **Named production shapes:** Meta, Amazon (Rufus), Stripe, Mistral AI, Cohere, Anyscale, Roblox on vLLM; Moonshot AI's Kimi K2 at 10x inference speedup on GB200 via Dynamo [VENDOR]; Mistral Large 3 at 10x faster inference via Dynamo [VENDOR]; Dell PowerScale + NIXL at 19x faster TTFT [VENDOR].
- **Enterprise procurement (2026):** regulated-industry buyers choose vLLM via Red Hat's support SLAs and AI Validated Models; Azure productizes SGLang inside first-party AMD endpoints (Azure's product, not SGLang's); as of September 2026 no Red Hat- or NIM-equivalent commercial SGLang distribution had surfaced — the commercial-packaging asymmetry is vLLM's structural enterprise advantage, now partly answered by RadixArk (below).

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

