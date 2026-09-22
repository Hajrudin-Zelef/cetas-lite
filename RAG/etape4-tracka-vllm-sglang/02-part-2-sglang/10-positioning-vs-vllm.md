---
id: etape4-tracka-vllm-sglang/02-part-2-sglang/10-positioning-vs-vllm
title: "10. Positioning vs vLLM"
domain: part-2-sglang
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "DeepSeek", "Huawei", "Intel", "Nvidia", "SGLang", "Z.ai", "vLLM"]
dates: ["2024-01-17", "2024-12-04", "2026-01-16", "2026-02-19", "2026-09-04", "2026-09-22"]
keywords: ["vllm", "agentic", "agents", "amd", "ascend", "benchmark", "blackwell", "copilot", "cost", "deepseek", "diffusion", "disclosure"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [919, 975]
section: "PART 2 — SGLang"
sha256: 66c0a720aa3a7d1b69b47b9d85f80b35667c0674b0d73a38c153bccf1fe246a4
---

# 10. Positioning vs vLLM

## 10. Positioning vs vLLM

| Dimension | SGLang | vLLM |
|---|---|---|
| Core scheduling idea | **RadixAttention**: radix-tree prefix KV reuse across requests (automatic) | **PagedAttention**: block-level KV management; prefix caching opt-in (`--enable-prefix-caching`) |
| Sweet spot | Prefix-heavy: agents, multi-turn chat, RAG, few-shot evals (up to 6.4× claimed) | General production, broadest model coverage, most mature ecosystem |
| Raw throughput (H100, Llama 3.1 8B, high concurrency) | ~16,200 tok/s [secondary] | ~12,500 tok/s [secondary] |
| Single-stream unique prompts | Roughly parity to −12% [secondary] | Slightly ahead in one old test [secondary] |
| DeepSeek models | Day-0 support (V3/R1/V4), optimized MLA/DSA/MTP, claimed 3.1× on V3 [official/secondary] | Supported; SGLang is the "officially recommended engine" per one comparison [secondary] |
| Speculative decoding | EAGLE/DFlash2/DSpark/KDA/NEXTN/MTP + beam search (v0.5.19) | EAGLE3 etc. |
| PD disaggregation | First-class (Mooncake/NIXL, DCP) | Available; llm-d covers both |
| Hardware breadth | NVIDIA + AMD (strong) + Intel XPU/CPU + TPU + Ascend + MUSA; HPU experimental | NVIDIA + AMD + TPU + XPU + HPU (Habana fork, mature-ish) |
| Structured output | XGrammar, compressed-FSM JSON (3× claim), frontend DSL | Outlines/XGrammar guidance |
| Governance | LMSYS non-profit + RadixArk commercial spin-out ($400M) | PyTorch Foundation (neutral IP), formal security process; Red Hat sells supported vLLM/llm-d |
| GitHub traction | 36.3K stars / 9.1K forks (2026-09-22) | (not pulled in this pass) |
| Default port | 30000 | 8000 |

**Practitioner summary:** choose SGLang when prefix reuse dominates (agentic/RAG/multi-turn), when serving DeepSeek-family models (day-0 kernels), or when AMD MI300X/MI355X is the fleet (feature parity + strong independent cost numbers in 2026). Choose vLLM for the broadest ecosystem, formal enterprise support channels (e.g., Red Hat), and Gaudi/HPU fleets. The raw-throughput gap is real at high concurrency but collapses on unique-prompt workloads.

---

## 11. Source list

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

