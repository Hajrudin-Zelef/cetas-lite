---
id: collect-261001-ia-llm/ia-llm/etape4-tracka-vllm-sglang-17
title: "Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "AWS", "DeepSeek", "Google", "Huawei", "Hugging Face", "Intel", "Meta", "Mistral", "Moonshot", "Nvidia", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM", "xAI"]
dates: ["2026-01-01", "2026-01-09", "2026-01-16", "2026-01-20", "2026-01-21", "2026-01-23", "2026-02-19", "2026-02-24", "2026-02-25", "2026-04-06", "2026-04-27", "2026-05-05", "2026-05-16", "2026-06-13", "2026-07-10", "2026-07-25", "2026-07-27", "2026-08-08", "2026-08-21", "2026-08-22", "2026-08-26", "2026-09-01", "2026-09-05", "2026-09-09", "2026-09-18", "2026-09-22"]
keywords: ["inference", "sglang", "vllm", "agentic", "amd", "ascend", "attention", "aws", "benchmark", "compute", "cost", "decode"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [1083, 1171]
sha256: 19bf3e9055ed331ac5c113429af71e74761987413f56dbcb88c8cc532656add3
---

# Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)

| Date | vLLM | SGLang |
|---|---|---|
| 2026-01-01 | — | v0.5.7 (day-0 Mimo-V2-Flash) [official] |
| 2026-01-09 | — | SGLang Model Gateway v0.3.1: 10–12× faster cache-aware routing [official; date secondary] |
| 2026-01-16 | — | lmsys blog: SGLang-Diffusion 2.5× faster than Nov-2025 release [official] |
| 2026-01-20 | **v0.14.0** — first 2026 release [official] | — |
| 2026-01-21 | — | TechCrunch: SGLang spins out as **RadixArk** (~$400M valuation, Accel) [secondary] |
| 2026-01-23 | — | v0.5.8 (diffusion speedups up to 1.5×) [official] |
| 2026-02-19 | — | lmsys blog: 25× inference perf on GB300 NVL72 [official] |
| 2026-02-24 | — | v0.5.9 (LoRA load/compute overlap, −78% TTFT) [official] |
| 2026-02-25 | v0.16.0 — V1 confirmed as only engine [secondary] | — |
| 2026-04-06 | — | v0.5.10 (piecewise CUDA graphs default, Elastic NIXL-EP, HiSparse) [official] |
| 2026-04-27 | v0.20.0 [official] | — |
| 2026-05-05 | — | v0.5.11 (CUDA 13 + Torch 2.11 default) [official] |
| 2026-05-16 | — | **v0.5.12 — DeepSeek-V4 day-0** [official] |
| 2026-06-13 | — | v0.5.13 (Nemotron 3 Ultra day-0) [official] |
| 2026-07-10 | v0.25.0 [official] | v0.5.15 (GLM-5.2 NVFP4 production tuning) [official] |
| 2026-07-25 | — | v0.5.16 (DSpark confidence-driven spec decode) [official] |
| 2026-07-27 | **v0.26.0** — Inkling family; DeepSeek-V4 perf push; vLLM Kimi-K3 blog: 370 tok/s w/ DSpark (3.14×), 16× GB300 NVL72 [official] | lmsys blog: RadixArk + Google bring full SGLang to TPUs [official] |
| 2026-08-08 | — | **v0.5.17 — Kimi K3 day-0** [official] |
| 2026-08-21 | vLLM optimizations push B200 ahead of MI355X-SGLang on DeepSeek V4 perf/$ (SemiAnalysis) [independent] | — |
| 2026-08-22 | — | v0.5.18 (710 PRs) [official; highlights not pulled] |
| 2026-08-26 | **v0.28.0** — Kimi-K3 perf push, DeepSeek V4 sparse MLA, KV disk offload, bitsandbytes→out-of-tree [breaking] [official] | — |
| 2026-08-26 | Google Cloud: vLLM TPU embedding pipelines GA on Ironwood [vendor-reported] | — |
| 2026-08 | — | CERT/CC-coordinated disclosure of 6 SGLang vulnerabilities [secondary] |
| 2026-09-01 | VMware AI Factory ships vLLM-based runtime [secondary] | — |
| 2026-09-05 | — | **v0.5.19** — beam search, DeepEP v2, unified radix tree default, AMD Lean attention [official] |
| 2026-09-09 | **v0.29.0** — Model Runner V2 default, 10 archs removed, RL weight sync [breaking] [official] | — |
| 2026-09-18 | — | **v0.5.20** — RL sampling masks, SGLang Simulator, CUDA 12 retired, XPU/MUSA images [official] |
| 2026-09-22 | **v0.30.0** — Fast Start IPC weight cache, Gumbel-max watermarking, HiSparse, DeepSeek-V4.1-Flash [breaking] [official] | — |

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

