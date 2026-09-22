---
id: etape4-tracka-vllm-sglang/03-part-3-synthesis-vllm-vs-sglang-2026-state/3-4-shared-ecosystem-used-by-both
title: "3.4 Shared ecosystem (used by both)"
domain: part-3-synthesis-vllm-vs-sglang-2026-state
role: deep-dive
task: reference
actors: ["AMD", "AWS", "DeepSeek", "Google", "Huawei", "Hugging Face", "Intel", "Meta", "Mistral", "Nvidia", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM", "xAI"]
dates: ["2026-09-22"]
keywords: ["agentic", "amd", "ascend", "attention", "aws", "benchmark", "cost", "decode", "deepseek", "diffusion", "disclosure", "glm"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [1114, 1171]
section: "PART 3 — SYNTHESIS: vLLM vs SGLang (2026 state)"
sha256: 621dd73eb882f5336fffa9919aaab325508a39f31560bd69b23072d2aae469b5
---

# 3.4 Shared ecosystem (used by both)

## 3.4 Shared ecosystem (used by both)

- **FlashInfer** (0.6.18): attention/kernel library required/used by both engines; MXFP4/MXFP8 MoE
  backends, TRT-LLM kernels, CuTe DSL [official].
- **NIXL**: NVIDIA's P/D-disaggregation KV-transfer library; vLLM NIXL connectors (v0.29 NIXL P/D DCP
  for MLA), SGLang Elastic NIXL-EP, NIXL RDMA backends [official].
- **Mooncake / MooncakeStore**: KV-cache transfer & L3 storage; vLLM Mooncake connectors (decode KV
  saving, encoder-cache sharing), SGLang MooncakeStore as HiCache L3 + MoE A2A backend; dependency
  `mooncake==0.3.13` (SGLang v0.5.19) [official].
- **DeepEP / DeepEP v2**: DeepSeek's expert-parallel all-to-all; both engines support (vLLM Elastic EP
  with DeepEP v2; SGLang `--moe-a2a-backend deepep_v2` in v0.5.19) [official].
- **Spec decode family**: EAGLE3 (both), DSpark (both, Inferact's open-source block-diffusion), DFlash2
  (both), MTP multi-token prediction [official].
- **llm-d** (4,624 stars): Kubernetes-native distributed inference founded by Red Hat/Google/IBM/NVIDIA;
  vLLM is the default model server, SGLang configs validated per release [secondary].
- **LMCache** (11,892 stars): KV-cache reuse/offload layer integrated with vLLM [secondary].
- **Hugging Face Transformers** 5.x (vLLM upgraded 5.15.0 in v0.28; SGLang 5.3.0 in v0.5.10) [official].
- **AMD AITER / Quark**: AMD kernel stack leveraged by both on ROCm (vLLM: Quark NVFP4; SGLang: Lean
  attention, MoRI/Mori-EP) [official].

## 3.5 Practitioner guidance (convergent across 2026 sources) [secondary][independent]

- **Choose SGLang** when prefix reuse dominates (agentic/RAG/multi-turn), for DeepSeek-family day-0
  kernels, or on AMD MI300X/MI355X fleets (strong independent cost numbers, e.g. up to ~40% $/M-token
  undercut vs B200 on GLM-5) [independent][secondary].
- **Choose vLLM** for the broadest ecosystem and hardware breadth (NVIDIA/AMD/Intel XPU+Gaudi/TPU/CPU/
  Ascend), fastest new-model support, and formal enterprise channels (Red Hat AI Inference Server,
  VMware AI Factory) [secondary].
- **TensorRT-LLM** leads NVIDIA-only peak long-context throughput at the cost of a compile step [independent].
- The headline "29% gap" collapses to ~0 on unique-prompt workloads; always qualify benchmark claims
  by workload shape [secondary][independent].

## 3.6 Unified uncertainty log

1. ⚠️ **vLLM v0.30.0 (released 2026-09-22, research day):** adoption/bug reports not yet available;
   breaking changes (scale-out opt-in, `g_idx` removal, gRPC entrypoint deprecation) carry early-adopter risk. [official]
2. ⚠️ **MRV1 removal targeted at vLLM v0.32** — date unknown; sequence parallelism / elastic EP /
   custom logits processors still fall back to MRV1. [official]
3. ⚠️ **SGLang security fixes** (6 Aug-2026 CERT/CC vulns): fix/landed-version status as of 2026-09-22
   unverified; repo lacked SECURITY.md at disclosure time. [secondary][unverified]
4. ⚠️ **Production-user claims** (vLLM: Meta/LinkedIn/Mistral/HF/Uber; SGLang: xAI/400k GPUs/trillions
   tokens) are third-party or self-reported; no independent audit. [unverified]
5. ⚠️ **AWS Trainium** support in vLLM 2026 — not confirmed from an authoritative source. [unverified]
6. ⚠️ **TGI maintenance-mode claim** (TECHSY, Dec 2025) — unverified.
7. ⚠️ **RadixArk $100M seed figure** — single-sourced; only the ~$400M valuation is TechCrunch-sourced. [unverified]
8. ⚠️ **vLLM README sponsor list** — read from forked README copies; verify against canonical before quoting. [official]
9. ⚠️ **SGLang v0.5.18 highlights** — release-notes body not pulled in the research pass. [official]
10. ⚠️ PR-reported micro-benchmark percentages in both engines' release notes (e.g. "−2.94% E2E TPOT",
    "3.43×") are PR-author claims, not audited results. [official — vendor-reported]
11. ⚠️ **SGLang SM120 INT4 correctness issue (#21132)** — community report predates v0.5.16 NVFP4 fixes;
    current status unverified. [secondary]
12. ⚠️ **SGLang HPU (Gaudi)** — experimental only; "not supported on Gaudi 3" per a May-2026 industry guide. [secondary]

---

*End of synthesis. Parts 1–2 above are the complete original research passes (kept verbatim); this Part 3
adds only the cross-engine comparison. All dated claims carry provenance tags; see §3.6 for open questions.*

