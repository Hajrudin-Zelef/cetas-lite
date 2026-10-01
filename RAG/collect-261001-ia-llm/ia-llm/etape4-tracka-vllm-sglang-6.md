---
id: collect-261001-ia-llm/ia-llm/etape4-tracka-vllm-sglang-6
title: "Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "AWS", "Alibaba", "Ant", "Cohere", "Crusoe", "DeepSeek", "Google", "Huawei", "Hugging Face", "Intel", "Lambda", "Meta", "Mistral", "Nebius", "Nvidia", "SGLang", "vLLM"]
dates: ["2025-01-27", "2026-05", "2026-06-05", "2026-09-01", "2026-09-04", "2026-09-22"]
keywords: ["inference", "sglang", "vllm", "accelerator", "acquisition", "amd", "apache", "ascend", "awq", "aws", "benchmarks", "blackwell"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [471, 563]
sha256: f6bb7c09f8527d514545e106332254a2470eb2b283e582cbc2e727fe447a1a58
---

# Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)

| Backend | Status | Evidence |
|---|---|---|
| NVIDIA CUDA | Primary; Hopper/Blackwell (+opt-in Rubin), Ada/Lovelace, GB10/GB300 | Release notes per-version kernel work; CUDA 13.0 default wheels [official] |
| AMD ROCm | First-class community path; AITER kernels; MI300X/MI325X/MI350X/MI355X (gfx90a/gfx942/gfx950); dual-stream decode w/ hipgraphs; W4A4 asm GEMM; TheRock 7.14 preview docker | v0.29/v0.30 release notes; official `vllm-openai-rocm` images [official] |
| Google TPU | Via `tpu-inference` plugin (vllm-project); v5e/v6e/v7x; native FP8 on v6e/Ironwood, INT8 on v5e | Google Cloud docs; dev.to/Google posts Aug 2026 [official][vendor-reported] |
| Intel HPU (Gaudi) | Out-of-tree plugin **vllm-gaudi** (now under vllm-project org), v0.26.0 tracking upstream v0.26.0 + Gaudi SW v1.24.1 + PT 2.11; Gaudi 2/3; INC/AWQ/GPTQ/ModelOpt/compressed-tensors; NIXL connector | vllm-gaudi release notes [official] |
| Intel XPU | Official wheels + `vllm-openai-xpu` docker; INC int4 W4A8 linear backend; AutoRound MXFP8 MoE; Data Center GPU Max 1550+ | v0.29/v0.30 notes [official] |
| CPU | In-tree; AVX512/AMX kernels (DeepSeek V2/V3/R1 MLA end-to-end on AMX); Int8 MoE via zentorch on AMD Zen; macOS/arm64 wheels | v0.29/v0.30 notes [official] |
| Huawei Ascend NPU | Out-of-tree **vllm-ascend** (vllm-project org), tracking upstream releases | vllm-ascend release notes [official] |
| Others (community) | Iluvatar BI-V150, MetaX C500X, Rebellions NPU via llm-d; AWS Neuron/Trainium referenced | llm-d accelerator docs [secondary] |

---

## 9. Ecosystem: repos, contributors, sponsors, forks

### 9.1 Repository scale (GitHub API, 2026-09-22) [official]
- vllm-project/vllm: **92,444 stars / 22,525 forks**
- sgl-project/sglang: 36,323 / 9,061 (SGLang comparison baseline)
- LMCache/LMCache: 11,892 / 1,938
- llm-d/llm-d: 4,624 / 786
- vllm-project/vllm-ascend: 2,873 / 2,347
- vllm-project/production-stack: 2,627 / 507
- vllm-project/tpu-inference: 438 / 322
- vllm-project/vllm-gaudi: 57 / 151
- vllm-project/dllm-plugin: 29 / 10 (diffusion-LLM plugin; LLaDA2.0 benchmarks May 2026)

### 9.2 Contributors & governance [official][secondary]
- Origins: UC Berkeley Sky Computing Lab + LMSYS; Woosuk Kwon initiated the project (scheduler/model
  runner). [official][secondary]
- V1 effort "mainly driven together" by **UC Berkeley, Neural Magic (now Red Hat), Anyscale, and Roblox**
  (vLLM blog, 2025-01-27). [official]
- Current contribution is heavily multi-vendor: NVIDIA, AMD, Intel, Red Hat, Google, Anyscale, Roblox
  engineers appear across release notes (kernel work, plugins, connectors); weekly merge volume in Sep
  2026 ≈ 358 PRs/week from 183 contributors (2026-09-04 week); v0.30.0: 762 commits / 315 contributors /
  104 new. [official][secondary]
- No Apache/CNCF/LF foundation umbrella; community-led open source governance (third-party assessment,
  Jul 2026). [secondary]
- UC Berkeley Sky Computing Lab and LMCache Lab (UChicago) are founding academic supporters of llm-d,
  described as "originators of vLLM" / "originators of LMCache". [secondary]

### 9.3 Sponsors (from the project's README sponsors section) [official]
- **Cash donations:** a16z, Dropbox, Sequoia Capital, Skywork AI, ZhenFund.
- **Compute resources:** Alibaba Cloud, AMD, Anyscale, AWS, Crusoe Cloud, Databricks, DeepInfra, Google
  Cloud, Intel, Lambda Lab, Nebius, Novita AI, NVIDIA, Replicate, Roblox, RunPod, Trainy, UC Berkeley,
  UC San Diego.
- **Slack sponsor:** Anyscale. Fundraising via OpenCollective (opencollective.com/vllm).
- ⚠️ Sponsor list read from README copies in GitHub forks (crawled Sep 2026); presumed current — the
  authoritative copy is https://github.com/vllm-project/vllm#sponsors. [official — verify before quoting verbatim]

### 9.4 Notable forks / related projects [official][secondary]
- **vllm-ascend** (Huawei Ascend NPU, vllm-project org) — out-of-tree plugin w/ own release line. [official]
- **vllm-gaudi** (Intel Gaudi, vllm-project org) — succeeded the deprecated HabanaAI/vllm-fork (EOL 2025-11). [official]
- **candle-vllm** (ericlbuehler) — Rust/Candle port with GPTQ/Marlin, FP8 KV, Metal/CPU, MCP support. [secondary]
- **vLLM SR / router ("Themis")** (vllm-project org) — serving-router with replay, observability, Redis/Valkey/Qdrant
  storage backends, long-context routing (v0.3 post, 2026-06-05). [official]
- **LMCache** (LMCache Lab, UChicago) — KV-cache reuse/offload layer integrated with vLLM. [secondary]
- **llm-d** (Red Hat/Google/IBM/NVIDIA) — Kubernetes-native distributed inference on top of vLLM/SGLang. [secondary]

---

## 10. Enterprise support & large deployments

### 10.1 Enterprise support options
- **Red Hat AI Inference Server** — enterprise-grade, commercially supported, security-hardened vLLM
  distribution (container image; standalone or in RHEL AI / OpenShift AI); built on vLLM + **Neural Magic**
  technologies (Red Hat acquired Neural Magic; LLM Compressor for quantization/pruning; pre-optimized model
  registry on Hugging Face; validated models; multi-accelerator: NVIDIA, AMD, Intel Gaudi, Google TPU);
  MLPerf inference results published with Supermicro on OpenShift AI + vLLM runtime (Red Hat blog).
  [vendor-reported][secondary]
- **nm-vllm** — Neural Magic's enterprise distribution of vLLM: stable builds with bug fixes and selected
  model backporting, **enterprise support with SLAs**, Kubernetes reference architectures, pre-optimized
  model registry (Red Hat Developer article; "nm-vllm" branding predates the acquisition — ⚠️ current
  naming under Red Hat AI should be verified). [vendor-reported]
- **VMware AI Factory (VCF)** — vLLM-based inference runtime as the managed model-serving layer for private
  AI on VMware Cloud Foundation (announced ~2026-09-01; techtimes). [secondary]
- ⚠️ No evidence found of a first-party "vLLM Enterprise" SKU from the vLLM project itself; enterprise
  support is delivered by vendors (Red Hat et al.). [unverified — absence of evidence]

### 10.2 Known production users (as reported by third parties; not officially confirmed by the companies)
- **Meta, LinkedIn, Mistral, Hugging Face** — named as running vLLM in production (third-party reading
  list, evaluated May 2026). [secondary — unverified by the companies]
- **Uber, LinkedIn** — named in a 2026 self-hosting guide as production users. [secondary — unverified]
- **Tesla** (ML Platform team) and **Cohere** (model serving + RL) — listed as llm-d users (2025/2026).
  [secondary]
- **JPMorgan Chase, Goldman Sachs** — cited as running AI workloads on Kubernetes with the vLLM-inclusive
  stack (third-party guide). [secondary — unverified]
- **Lambda, Mistral AI** — llm-d launch partners supporting the vLLM-based distributed-serving ecosystem.
  [secondary]

---

## 11. Competitive positioning vs SGLang (2026)

