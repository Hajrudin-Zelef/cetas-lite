---
id: ai-industry-kb-2026/06-inference-engines/sources-and-urls
title: "Sources and URLs"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Huawei", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Stripe", "TensorRT-LLM", "Unsloth", "vLLM"]
dates: ["2026-05", "2026-08"]
keywords: ["accelerator", "amd", "ascend", "attention", "benchmark", "blackwell", "cost", "decode", "dram", "gguf", "gpu", "hbm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2532, 2548]
section: "6. Inference Engines"
sha256: ea8668b0bc61ccd60a7f6eeb8fc8e6b6776dcb0f1bb3ee6bd4aa5077b2ea4b45
---

# Sources and URLs

1. **Budget on inference, not training:** with inference at 55%+ of AI-optimized IaaS spend (Gartner 2026) and GPT-4-level inference at ~$0.40/M tokens, engine choice is a first-order cost decision — Stripe's 73% fleet reduction and the 4–10x disaggregation dividend dwarf most model-selection debates.
2. **Default to vLLM for general GPU serving; keep the backend swappable:** broadest model support, disaggregation-ready, deepest operator knowledge base, commercial SLAs via Red Hat/IBM, NIM packaging, and the reference backend for llm-d. Abstract the backend behind a gateway from day one — 2026 engine deltas are workload-specific (SGLang wins prefix-heavy load; vLLM wins TTFT and 100+ concurrency).
3. **Re-validate ops tooling after the 2026 rewrites:** V0 removal (v0.16), PagedAttention deletion (v0.25), MRv2 default (v0.25), GGUF split (v0.24) mean monitoring, profiling, and deployment tooling built on older vLLM assumptions need re-validation — the engine kept its API surface while replacing almost everything underneath.
4. **Stop describing PagedAttention as the current mechanism:** v0.25.0 removed it from the main internal path (legacy only). The 2026 serving story is Model Runner V2, sparse attention (DSA/CSA), KV offloading/tiering, and disaggregation.
5. **Disaggregation is the architecture, not an option:** prefill/decode split is the default for new production deployments; size pools independently and treat KV transfer as a production data path (measure end-to-end latency including queue time, alert on tail latency, verify RDMA driver health per node).
6. **NIXL is the shared data plane:** vLLM, Dynamo, SGLang, TensorRT-LLM, LMCache, and llm-d all move KV on NIXL — orchestration choice (Dynamo vs llm-d) no longer forces a data-plane choice. Think of the KV cache as a cluster-wide resource with its own scheduling concerns (HBM→DRAM→NVMe tiers, global prefix index), not per-worker memory.
7. **Treat vLLM v0.28.0 dates carefully:** it is the core release (Kimi-K3, sparse MLA, MRv2, tiered offload), not a "vLLM-Omni" variant; and V1-only dates to v0.16.0 (March), not v0.28.0. Wave-2-era material mixing these up must be re-attributed.
8. **Commercialize the inference layer deliberately:** the Inferact ($150M/$800M) and NIM 2.0/RHIS moves show the open engine is now the monetizable layer — proprietary value accrues at orchestration (Dynamo), packaging (NIM, RHIS), and silicon, not at the serving kernel.
9. **Heterogeneous clusters are real for vLLM, but weight by maturity:** CUDA > ROCm > TPU (offline/batch) > Neuron (Beta). The "one engine, any accelerator" claim survives scrutiny at the API level; production guidance should weight the Neuron and TPU paths below CUDA/ROCm.
10. **Benchmark hygiene decides engine selection:** independent August 2026 measurements contradicted published TensorRT-LLM throughput claims on workstation Blackwell; always measure on your own hardware class and workload shape before choosing an engine — published deltas do not transfer across GPU generations.
11. **Two commercial stacks by end of 2026:** the open engine is no longer a neutral substrate — Red Hat/IBM/NVIDIA/Inferact commercialize vLLM; RadixArk ($100M, May 2026) commercializes SGLang. Procurement should evaluate the commercial stack (support SLAs, validated models, hosting tiers) alongside the engine, because the commercial surfaces now diverge while the engines converge on OpenAI-compatible APIs, XGrammar, and NIXL.
12. **Watch time-to-serve, not architecture:** vLLM v0.28.0 and SGLang v0.5.19 both raced to serve the same late-2026 flagships (Qwen3.8, Kimi-K3) across NVIDIA/AMD/Ascend — engine differentiation is increasingly measured in *time-to-serve* for new models, not architectural features.
13. **GGUF on GPU is fragmented in 2026 — plan for it:** vLLM's plugin (young, GPU-only, per-family bugs), Unsloth's Dynamic v3.0 pipeline (local-first, AMD/Windows coverage the plugin lacks), llama.cpp as the reference. Pin plugin commits, validate per model family, and keep llama.cpp as the first diagnostic.
14. **Version everything, date everything:** vLLM moved roughly a minor per month in 2026 and the GGUF plugin is pinned from git main in working ROCm builds — "latest" is a date, not a version. Bake the date into every version claim in the consolidated document.

## Sources and URLs

