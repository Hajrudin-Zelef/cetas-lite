---
id: collect-261001-ia-llm/ia-llm/etape4-tracka-vllm-sglang-15
title: "Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Google", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2024-01-17", "2024-12-04", "2026-01-01", "2026-01-09", "2026-01-16", "2026-01-21", "2026-01-23", "2026-02-19", "2026-02-24", "2026-03-28", "2026-04-06", "2026-05-05", "2026-05-16", "2026-06-13", "2026-06-26", "2026-07-10", "2026-07-25", "2026-08-08", "2026-08-22", "2026-09-04", "2026-09-05", "2026-09-18", "2026-09-22"]
keywords: ["inference", "sglang", "vllm", "agentic", "agents", "benchmark", "benchmarks", "blackwell", "compute", "copilot", "decode", "deepseek"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [942, 1022]
sha256: 1b938701022e36af28a10349dff4181eef008cd4bad3e371b622c418fcf034c3
---

# Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)

- https://github.com/sgl-project/sglang (README, releases)
- https://api.github.com/repos/sgl-project/sglang (repo stats)
- https://api.github.com/repos/sgl-project/sglang/releases?per_page=20 (release dates)
- https://github.com/sgl-project/sglang/releases/tag/v0.5.20
- https://github.com/sgl-project/sglang/releases/tag/v0.5.19
- https://github.com/sgl-project/sglang/releases/tag/v0.5.18
- https://github.com/sgl-project/sglang/releases/tag/gateway-v0.3.1
- https://github.com/sgl-project/sglang/releases/tag/v0.4.1
- https://github.com/sgl-project/sglang-omni
- https://github.com/sgl-project/sglang/pull/2357 and /pull/2121 (HPU)
- https://lmsys.org/blog/2024-01-17-sglang/ · https://lmsys.org/blog/2024-12-04-sglang-v0-4/ · https://lmsys.org/blog/2026-01-16-sglang-diffusion/ · https://lmsys.org/blog/2026-02-19-gb300-longctx/
- https://techcrunch.com/2026/01/21/sources-project-sglang-spins-out-as-radixark-with-400m-valuation-as-inference-market-explodes/
- https://techfundingnews.com/radixark-sglang-spinoff-400m-valuation-ai-inference/
- https://www.beri.net/article/vllm-vs-tensorrt-llm-vs-sglang-inference-runtime-2026
- https://particula.tech/blog/sglang-vs-vllm-inference-engine-comparison
- https://blog.premai.io/vllm-vs-sglang-vs-lmdeploy-fastest-llm-inference-engine-in-2026/
- https://github.com/semianalysisai/inferencex-app (InferenceX MI355X Qwen3.5 + GLM-5 posts)
- https://github.com/dstackai/dstack (DeepSeek-R1 H200 benchmark)
- https://www.marktechpost.com/2025/11/07/comparing-the-top-6-inference-runtimes-for-llm-serving-in-2025/
- https://www.spheron.network/blog/intel-gaudi-3-vs-nvidia-h200-b200-llm-inference-2026/
- https://github.com/jpezzulli/sglang-rtxpro6000/blob/HEAD/RESULTS.md
- https://github.com/lEWFkRAD/qwen38-rtx-pro-6000
- https://github.com/wrg-11/wrg-sigma-rules/blob/HEAD/docs/detection-notes/sglang-2026-08-disclosure-series-detection-2026-09-04.md
- https://github.com/rohitg00/llm-d/blob/HEAD/guides/pd-disaggregation/README.md
- https://github.com/hygon-ai/sglang-das/blob/HEAD/docs/docs/advanced_features/sgl_model_gateway.mdx
- https://github.com/fuzzifikation/vllm-copilot/blob/HEAD/docs/sglang-compat-plan.md
- https://github.com/vroomfondel/dgxarley/blob/HEAD/SGLANG_TP_EP_MOE_UPSTREAM_BUG.md
- https://github.com/hitechcloud-vietnam/dgxarley/blob/HEAD/SGLANG_v0.5.10_VERSION_CHANGES.md
- https://github.com/allanschramm/local-model-autotuning/blob/HEAD/docs/discovery/sglang-inference-engine.md
- https://github.com/mudler/localai/blob/HEAD/.agents/sglang-backend.md
- https://github.com/profsynapse/synaptic-tuner/blob/HEAD/docs/preparation/vllm-vs-sglang-inference-serving-research.md
- https://github.com/brendanmckeag/sglang-vllm-benchmark/blob/HEAD/README.md
- https://github.com/randomchaos7800-hub/inference-research/blob/HEAD/tower/gdn-blackwell/sglang-vs-vllm-sm120.md

## 12. Open questions / uncertainties

1. **v0.5.18 highlights** — release-notes body not pulled; only size (710 PRs / 212 contributors) confirmed [official].
2. **v0.5.17 Kimi-K3 details** — release body truncated in API fetch; "2.8T-parameter multimodal LatentMoE (896 experts, top-16…)" is a partial quote.
3. **Gateway-v0.3.1 date** (2026-01-09) is [secondary]-sourced only.
4. **RadixArk $100M seed figure** — single-sourced summary; only the ~$400M valuation is TechCrunch-sourced.
5. **Security fixes** — six Aug-2026 CERT/CC-coordinated issues; fix/landed-version status as of 2026-09-22 not verified.
6. **No Artificial Analysis, ServeTheHome, or SiliconANGLE** SGLang benchmarks/articles were found in this pass.
7. **Adoption/GPU/token counts** ("400,000 GPUs", "trillions of tokens/day") are project self-reports — no independent audit found.
8. **SM120 INT4 correctness issue (#21132)** — community report predates v0.5.16 NVFP4 fixes; current status not verified.
9. **vLLM star/fork counts** and **SGLang enterprise pricing/SLAs** were not pulled (out of scope for this pass).

---

## 13. 2026 news & blog timeline (official unless noted)

- **2026-01-01** — v0.5.7 released (day-0 Mimo-V2-Flash) [official].
- **2026-01-09** — SGLang Model Gateway v0.3.1: 10–12× faster cache-aware routing, 99% memory reduction [official; date [secondary]].
- **2026-01-16** — lmsys.org blog: "SGLang-Diffusion: Two Months In" — up to 2.5× faster than its Nov-2025 initial release [official].
- **2026-01-21** — TechCrunch: "Project SGLang spins out as RadixArk with $400M valuation" (round led by Accel; announced Aug 2025; founders Ying Sheng + Banghua Zhu) [secondary].
- **2026-01-23** — v0.5.8: diffusion speedups up to 1.5× [official].
- **2026-02-19** — lmsys.org blog: "Unlocking 25x Inference Performance with SGLang on NVIDIA GB300 NVL72" [official].
- **2026-02-24** — v0.5.9: LoRA weight-load/compute overlap (−78% TTFT) [official].
- **2026-03-28** — v0.5.10rc0 prerelease [official].
- **2026-04** — lmsys.org blog: "DeepSeek-V4 on Day 0: From Fast Inference to Verified RL with SGLang and Miles" [official].
- **2026-04-06** — v0.5.10: piecewise CUDA graphs default, Elastic NIXL-EP, HiSparse, FlashInfer MXFP8, transformers 5.3.0 [official].
- **2026-05-05** — v0.5.11: CUDA 13 + Torch 2.11 default [official].
- **2026-05-16** — v0.5.12: DeepSeek-V4 day-0 support [official].
- **2026-06** — lmsys.org blog: "The next generation of speculative decoding: DFlash and Spec V2" [official].
- **2026-06** — lmsys.org blog: day-0 support for Nemotron 3 Ultra / Nemotron 3 Super / Higgs Audio v3 TTS [official].
- **2026-06-13** — v0.5.13: Nemotron 3 Ultra day-0 [official].
- **2026-06-26** — v0.5.14: GLM-5.2, LFM2.5, Kimi-K2.7-Code [official].
- **2026-07** — lmsys.org blog: "Serving GLM5.2 NVFP4 agentic workloads with SGLang: Reaching 500 TPS in two weeks" [official].
- **2026-07** — lmsys.org blog: "SGLang and Miles add day-0 support for Kimi K3" [official].
- **2026-07** — lmsys.org blog: "RadixArk and Google bring full SGLang features to TPUs" [official].
- **2026-07-10** — v0.5.15: GLM-5.2 NVFP4 production tuning [official].
- **2026-07-25** — v0.5.16: DSpark speculative decoding [official].
- **2026-08** — SGLang provides day-0 support for OpenAI gpt-oss (noted 2025-08 in README news; gpt-oss PD-decode prefix reuse for SWA hybrids landed v0.5.19) [official].
- **2026-08** — CERT/CC-coordinated disclosure of six SGLang vulnerabilities; mitigations published by community (pickle IPC, dumper port, API key, network isolation) [secondary].
- **2026-08-08** — v0.5.17: Kimi K3 day-0 [official].
- **2026-08-22** — v0.5.18 [official].
- **2026-09-05** — v0.5.19: beam search, DeepEP v2, unified radix tree default [official].
- **2026-09-18** — v0.5.20: RL sampling masks, SGLang Simulator, CUDA 12 retired [official].
- **2026-09** — SGLang-Omni v0.1.6 on PyPI; day-0 AuK/AuK-Flash audio models [official].

## 14. Practical deployment notes (from docs & community, 2026)

