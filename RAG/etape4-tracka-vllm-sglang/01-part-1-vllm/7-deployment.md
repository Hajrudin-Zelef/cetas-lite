---
id: etape4-tracka-vllm-sglang/01-part-1-vllm/7-deployment
title: "7. Deployment"
domain: part-1-vllm
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Alibaba", "Broadcom", "DeepSeek", "Google", "Huawei", "Intel", "Moonshot", "Nebius", "Nvidia", "Oracle", "Z.ai", "vLLM"]
dates: ["2026-06-30", "2026-08-26", "2026-09-01", "2026-09-22"]
keywords: ["accelerator", "amd", "ascend", "awq", "aws", "blackwell", "decode", "deepseek", "embedding", "fp8", "glm", "gptq"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [426, 483]
section: "PART 1 — vLLM"
sha256: cb1552bc5d7f3128780413d3d4d41daa450ca352064c782ff2c0d0d892a86b27
---

# 7. Deployment

## 7. Deployment

### 7.1 Official artifacts [official]
- **PyPI:** `pip install vllm` (CUDA 13.0); `uv pip install vllm --torch-backend=auto`; ROCm wheels at
  `https://wheels.vllm.ai/rocm/<ver>/rocm723`; XPU via `https://wheels.vllm.ai/<ver>/xpu`; CPU wheels
  (x86_64, arm64, macOS); CUDA 12.9/13.0 wheels for x86_64 and arm64.
- **Docker (v0.30.0):** `vllm/vllm-openai:v0.30.0` (CUDA 13.0), `:v0.30.0-cu129`, `:v0.30.0-ubuntu2404`,
  `:v0.30.0-cu129-ubuntu2404`, **`vllm/vllm-openai-rocm:v0.30.0`**, **`vllm/vllm-openai-cpu:v0.30.0`**,
  **`vllm/vllm-openai-xpu:v0.30.0`**; opt-in Rubin builds for CUDA 13.4/13.5 (v0.29).
- **Hardware plugins (out-of-tree):** vllm-ascend (Huawei Ascend NPU; release line tracks upstream,
  e.g. v0.22.1rc1 2026-06-30 with Mooncake DeepSeek-V4/HCCL weight transfer/Ascend 950 W8A8),
  vllm-gaudi (Intel Gaudi 2/3; v0.26.0 on upstream v0.26.0 + Gaudi software v1.24.1),
  tpu-inference (Google TPU v5e/v6e/v7x; v0.28.0 pinned in vLLM v0.29 deps), XPU support in-tree-ish
  (official wheels + docker), CPU backend in-tree.

### 7.2 Kubernetes [official][secondary]
- **vLLM Production Stack** (https://github.com/vllm-project/production-stack; 2,627 stars / 507 forks
  as of 2026-09-22 [official]): Helm chart (`helm repo add vllm https://vllm-project.github.io/production-stack;
  helm install vllm vllm/vllm-stack -f values.yaml`) installing serving engine + KV-cache-aware **request
  router** behind ClusterIP; Prometheus/Grafana observability; autoscaling guidance via KEDA on
  `vllm:num_requests_waiting`; deploy guides incl. Oracle OKE and Nebius; **Kubernetes operator with
  `VLLMRuntime` and `VLLMRouter` CRDs** (production-stack.ai API group). [official][secondary]
- Ecosystem: **llm-d** (Red Hat/Google/IBM/NVIDIA founded; 4,624 stars; P/D disaggregation at scale, GKE
  well-lit paths, NIXL KV transfer; vLLM as default model server; AMD/Intel/TPU/Rebellions accelerator
  maintainers), **AIBrix** (4.7k stars; LoRA management, SLO-aware autoscaling), **KubeAI** (1.2k stars,
  lightweight, scale-from-zero), Ray Serve + vLLM (KubeRay) for multi-node/multi-model. [secondary]
- Docs: https://docs.vllm.ai (official); vLLM recipes tool: https://recipes.vllm.ai (hardware × model ×
  feature recipes, e.g. Kimi-K2.5 on MI355X with tool_calling/reasoning/encoder_parallel). [official]

### 7.3 Cloud integrations [official][vendor-reported][secondary]
- **Google Cloud:** vLLM on TPU via `tpu-inference` plugin (JAX + PyTorch unified runtime; TPU v5e/v6e/v7x,
  Trillium, Ironwood); GKE deployment guides; 2026-08-26 native TPU support for production embedding
  pipelines (see §6.1).
- **AWS:** Trainium support exists (named in comparison tables; no dedicated 2026 announcement found —
  ⚠️ status should be confirmed in docs). [unverified]
- **Broadcom/VMware:** VMware AI Factory (announced ~2026-09-01, [secondary]) ships a **vLLM-based
  inference runtime** for VCF private AI, >150 models runnable, 5 fully validated as managed services
  (NVIDIA Nemotron 3, Google DeepMind Gemma 4, NEC cotomi, Alibaba Qwen 3.8-27B, Z.ai GLM 5.2); AMD
  MI350 + ROCm validated path with zero-touch provisioning.
- **Red Hat OpenShift AI:** see §9.

---

## 8. Hardware backends (2026 status)

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

