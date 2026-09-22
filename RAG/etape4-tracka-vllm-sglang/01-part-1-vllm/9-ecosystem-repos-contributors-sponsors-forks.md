---
id: etape4-tracka-vllm-sglang/01-part-1-vllm/9-ecosystem-repos-contributors-sponsors-forks
title: "9. Ecosystem: repos, contributors, sponsors, forks"
domain: part-1-vllm
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Alibaba", "Ant", "Crusoe", "Google", "Huawei", "Intel", "Lambda", "Nebius", "Nvidia", "SGLang", "vLLM"]
dates: ["2025-01-27", "2026-05", "2026-06-05", "2026-09-04", "2026-09-22"]
keywords: ["amd", "apache", "ascend", "aws", "benchmarks", "compute", "diffusion", "fp8", "governance", "gptq", "inference", "intel"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [484, 530]
section: "PART 1 — vLLM"
sha256: 839845ec56972115c43f66b0e68aeb20b8b203710c0fe1ab942de18560b130f6
---

# 9. Ecosystem: repos, contributors, sponsors, forks

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

