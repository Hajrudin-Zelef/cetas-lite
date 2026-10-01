---
id: collect-261001-ia-llm/ia-llm/etape4-tracka-vllm-sglang-5
title: "Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "AWS", "Alibaba", "Broadcom", "DeepSeek", "Google", "Huawei", "Intel", "Moonshot", "Nebius", "Nvidia", "Oracle", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM"]
dates: ["2026-03-06", "2026-05", "2026-06-30", "2026-08-21", "2026-08-26", "2026-09-01", "2026-09-22"]
keywords: ["inference", "sglang", "vllm", "accelerator", "agentic", "amd", "ascend", "awq", "aws", "benchmark", "benchmarks", "cost"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [384, 470]
sha256: 0b25f7069165251bb8f8016da76eeba5980eedb7dcb39b85691f70ebea648be5
---

# Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)

### 6.2 Independent measurements
- **SemiAnalysis InferenceX (Aug 2026)** [independent]:
  - DeepSeek V4 Pro 0813 (1.6T params / 49B active), agentic-coding replay (p50 88k in / p90 272k in tokens):
    MI355X **SGLang** matched B200 **vLLM** on perf/$ **until 2026-08-21**, when vLLM optimizations from
    Inferact and NVIDIA pushed B200 ahead — "a close race rather than a settled one". AMD published its own
    DeepSeek V4 vLLM optimization list (vllm-project/vllm#52911).
  - AMD MI355X Kimi-K2.5 MXFP4: one vLLM PR (#35850, merged 2026-03-06, shipped v0.18) moved 6.6 → **78.9
    tok/s/user** (8k/1k workload), 12.0× interactivity at low batch, 7.7× peak throughput, 15× at
    iso-throughput; peak **2,687 tok/s/GPU** (25 days from baseline to full effect).
  - Kimi K3: between 40–60 s E2E latency, MI355X **ATOM** (AMD vendor engine) beats even GB300 NVL72 vLLM on
    perf/$ — caveat: ATOM is a vendor engine, vLLM is the open comparison.
- **dstack DeepSeek-R1 benchmark (2026)** [independent]:
  - H200: TensorRT-LLM highest online throughput (4,176 tok/s); **vLLM led at concurrencies <128** in online
    throughput + E2E latency; offline: SGLang 6,311 tok/s.
  - MI300X: **vLLM outperformed SGLang in both online and offline throughput and E2E latency**; best online
    4,574 tok/s (vLLM); SGLang better only at concurrencies <32.
- **L40S head-to-head (Medium, Aug 2026)** [independent]: long-context workload — TensorRT-LLM ~2× throughput
  (75.9 vs ~40 tok/s) and 7.5× faster avg TTFT (8.6 s vs vLLM 63.8 s / SGLang 68.6 s); vLLM slight edge over
  SGLang in throughput (41.0 vs 38.2 tok/s) and TTFT; vLLM ITL 68.0 ms vs SGLang 72.6 ms. Quantization sweep
  on vLLM: FP8 vs AWQ vs GPTQ (figures only, no extracted numbers).
- **Single-L4 reproducibility study (inference-bench, GitHub)** [independent]:
  - Short regime (c=64, Qwen2.5-7B): vLLM AWQ **976 tok/s** > SGLang FP16 914 > vLLM FP16 831 > SGLang AWQ 506.
  - A100 sweep (vLLM v0.20.1 vs SGLang, May 2026, Qwen2.5-7B): vLLM FP16 3,102 tok/s @ c=64 vs SGLang 2,141;
    Marlin kernels +70% in v0.20.1 (177 vs 104 tok/s @ c=1 vs v0.8.5); SGLang Marlin collapsed at c=64
    (2,231 vs vLLM 4,762) — "opposite of the L4 pattern where SGLang wins at c=64".
- **ai-lab-benchmarks (Aug 2026, production model Qwen3.6-35B-A3B)** [independent]: NVFP4/vLLM vs GGUF/llama.cpp —
  4.7× throughput (47.1 vs 10.0 items/s), 3.4× faster prompt reading (13,916 vs 4,068 tok/s at 8k), generation
  201.5 vs 167 tok/s; translation quality identical (chrF++ 69.33 vs 69.34). Author noted SGLang setup friction
  (mem-fraction 0.80 vs vLLM 0.90; NVFP4 MoE runner backend flag) — "real setup cost".
- **Architecture-aware AMD study (arXiv 2603.10031, Feb 2026, MI325X ×8, vLLM v0.14.1)** [independent]:
  Llama-3.1-405B 15,944 tok/s vs DeepSeek V3.2 15,343 tok/s peak (text-only); Qwen3-VL-235B 47,873 tok/s
  (vision, incl. image tokens); Kimi-K2.5 7,327 tok/s; AITER required for competitive MLA (+3–5% at high
  concurrency on Llama-405B, MoE/MLA speedups larger).
- ⚠️ Legacy A100 numbers still circulating (17.12 req/s OPT-13B vs HF TGI 0.71 req/s) are from the
  original vLLM paper era — illustrative of PagedAttention's original gains, not 2026 hardware. [secondary]
- ⚠️ Viral "29% gap" claims (SGLang ~16,200 vs vLLM ~12,500 tok/s on H100 Llama-3.1-8B): the primary source
  (RunPod) attributes the gap almost entirely to **prefix-heavy** traffic (RadixAttention prefix reuse);
  on unique prompts the engines are "within a few percent". Do not quote the 29% without the workload
  qualifier. [secondary]

---

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

