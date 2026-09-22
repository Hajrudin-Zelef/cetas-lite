---
id: etape4-tracka-vllm-sglang/02-part-2-sglang/overview
title: "PART 2 — SGLang"
domain: part-2-sglang
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Alibaba", "Baidu", "Baseten", "ByteDance", "DeepSeek", "Google", "Intel", "Microsoft", "Moonshot", "Nebius", "Nvidia", "Oracle", "SGLang", "Z.ai", "vLLM", "xAI"]
dates: ["2024-01-08", "2026-01-01", "2026-01-09", "2026-01-21", "2026-01-23", "2026-02-24", "2026-03-28", "2026-04-06", "2026-04-09", "2026-05-05", "2026-05-16", "2026-05-26", "2026-06-13", "2026-06-26", "2026-07-10", "2026-07-14", "2026-07-25", "2026-08-08", "2026-08-22", "2026-09", "2026-09-05", "2026-09-18", "2026-09-22"]
keywords: ["sglang", "amd", "apache", "attention", "aws", "blackwell", "compute", "deepseek", "diffusion", "glm", "governance", "gpu"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [639, 698]
section: "PART 2 — SGLang"
sha256: d1099875bcc86743972866b11a49c7192bdad095e705dd12dbf5213406fb1f5f
---

# PART 2 — SGLang

# SGLang — Inference-Serving Framework (LMSYS) — Research Report

> **Research date:** 2026-09-22 (UTC)
> **Project:** RAG data collection, Step 4 (AI infra — inference/training)
> **Scope:** SGLang news and state as of September 2026
> **Provenance legend:** `[official]` = sgl-project GitHub / docs.sglang.io / lmsys.org blog; `[vendor-reported]` = company marketing claims; `[independent]` = third-party measurements (SemiAnalysis InferenceX, dstack, community); `[secondary]` = press/blogs summarizing others; `[unverified]` = could not be confirmed.
> Every fact below carries its date. Uncertainties are marked explicitly.

---

## 1. Project identity & ecosystem snapshot

- **What it is:** SGLang is a high-performance serving framework for large language models and multimodal models, "designed to deliver low-latency and high-throughput inference across a wide range of setups, from a single GPU to large distributed clusters" [official, GitHub README, viewed 2026-09-22](https://github.com/sgl-project/sglang).
- **Repository:** `github.com/sgl-project/sglang`, Apache-2.0 license, default branch `main`, created 2024-01-08 [official, GitHub API, 2026-09-22](https://api.github.com/repos/sgl-project/sglang).
- **GitHub traction at research date:** **36,323 stars**, **9,061 forks**, 184 subscribers, 5,359 open issues [official, GitHub API, `updated_at` 2026-09-22T19:08:12Z](https://api.github.com/repos/sgl-project/sglang). Last push at research date: 2026-09-22T19:29:47Z.
- **Governance:** Hosted under the non-profit open-source organization **LMSYS** [official]. Originated in 2023 inside the UC Berkeley lab of Databricks co-founder Ion Stoica [secondary, TechCrunch 2026-01-21](https://techcrunch.com/2026/01/21/sources-project-sglang-spins-out-as-radixark-with-400m-valuation-as-inference-market-explodes/).
- **Scale claims (project-stated):** "deployed at large scale, generating trillions of tokens in production each day"; "deployments running on over 400,000 GPUs worldwide"; "the de facto industry standard" [official, README]. Treat exact token/GPU counts as **[unverified]** (self-reported, no independent audit found).
- **Design heritage:** README acknowledges reusing design/code from Guidance, vLLM, LightLLM, FlashInfer, Outlines, LMQL [official].

### Named adopters (README, 2026-09-22) [official]
xAI, NVIDIA, AMD, Intel, LinkedIn, Cursor, Oracle Cloud, Google Cloud, Microsoft Azure, AWS, Atlas Cloud, Voltage Park, Nebius, DataCrunch, Novita, RunPod, InnoMatrix, Modal, MIT, UCLA, University of Washington, Stanford, UC Berkeley, Tsinghua University, Baseten, Baidu, AntGroup, Alibaba, Tencent. (Note: an older README revision listed "Jam & Tea Studios" and fewer names; the current list is the one above.) Adopters are self-declared by the project — treat the *production* nature of each as [vendor-reported]/[unverified] unless independently confirmed.

### RL / post-training role
SGLang is a "proven rollout backend used for training many frontier models", with native RL integrations in **AReaL, Miles, slime, Tunix, verl** [official, README, 2026-09-22]. A ByteDance-iaas fork README (crawled 2026-09-18) mirrors the same list [secondary](https://github.com/bytedance-iaas/sglang/blob/HEAD/README.md). A Jan-2026 research note described SGLang as an "RL/post-training rollout backend: verl, AReaL, slime, Tunix, Miles" [secondary](https://github.com/allanschramm/local-model-autotuning/blob/HEAD/docs/discovery/sglang-inference-engine.md).

---

## 2. 2026 release timeline (official GitHub releases)

All dates below are `published_at` from the GitHub Releases API [official](https://api.github.com/repos/sgl-project/sglang/releases?per_page=20), fetched 2026-09-22.

| Version | Release date | Size | Headline |
|---|---|---|---|
| **v0.5.20** | **2026-09-18** | 713 PRs, 237 contributors | RL sampling masks, SGLang Simulator, CUDA 12 lane retired, XPU/MUSA release images |
| v0.5.19 | 2026-09-05 | 786 PRs, 214 contributors | Beam search, DeepEP v2, LayerNorm SP, unified radix tree by default, Lean attention on AMD |
| v0.5.18 | 2026-08-22 | 710 PRs, 212 contributors | (highlights not pulled in this pass — see uncertainties) |
| v0.5.17 | 2026-08-08 | 582 PRs, 194 contributors | **Kimi K3 day-0 support** |
| v0.5.16 | 2026-07-25 | 574 PRs, 169 contributors | DSpark confidence-driven speculative decoding |
| v0.5.15.post1 | 2026-07-14 | patch | GLM-5.2 / DSA launch fixes |
| v0.5.15 | 2026-07-10 | — | GLM-5.2 NVFP4 production tuning on Blackwell |
| v0.5.14 | 2026-06-26 | — | New models: GLM-5.2, LFM2.5, Kimi-K2.7-Code |
| v0.5.13 | 2026-06-13 | — | Nemotron 3 Ultra day-0 support |
| v0.5.12.post1 | 2026-05-26 | stability patch | 12 cherry-picked fixes, mostly DeepSeek-V4 |
| v0.5.12 | 2026-05-16 | — | **DeepSeek-V4 day-0 support** |
| v0.5.11 | 2026-05-05 | — | **CUDA 13 + Torch 2.11** become default; transformers upgrade path |
| v0.5.10.post1 | 2026-04-09 | patch | FlashInfer 0.6.7.post2 → 0.6.7.post3 (JIT cubin downloader fix) |
| v0.5.10 | 2026-04-06 | — | Piecewise CUDA graph default; Elastic NIXL-EP; HiSparse; FlashInfer MXFP8 kernels |
| v0.5.10rc0 | 2026-03-28 | prerelease | v0.5.10 release candidate |
| v0.5.9 | 2026-02-24 | — | LoRA weight-load/compute overlap (−78% TTFT) |
| v0.5.8 | 2026-01-23 | — | Diffusion speedups (up to 1.5× across major diffusion models) |
| v0.5.7 | 2026-01-01 | — | Day-0: Mimo-V2-Flash; other new models |
| gateway-v0.3.1 | ~2026-01-09 [secondary] | — | SGLang Model Gateway: 10–12× faster cache-aware routing, 99% less memory |

Notes:
- v0.5.7 was published 2026-01-01T10:01:57Z [official]. The v0.5.0–v0.5.6 series dates to late 2025 and is out of this report's scope; not pulled.
- Cadence in 2026 is roughly one minor release every 2–4 weeks, with `.post1` stability patches between.
- "gateway-v0.3.1" date (2026-01-09) comes from a [secondary] research note, not the releases API — the release notes themselves say "SMG v0.3.1 Released!" with "10-12x performance improvement and 99% memory reduction in cache-aware routing" [official](https://github.com/sgl-project/sglang/releases/tag/gateway-v0.3.1).

