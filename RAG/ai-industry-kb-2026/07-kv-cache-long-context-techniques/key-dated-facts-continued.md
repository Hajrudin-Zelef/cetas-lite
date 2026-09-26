---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/key-dated-facts-continued
title: "Key dated facts (continued)"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["DeepSeek", "MiniMax", "Nvidia", "vLLM"]
dates: ["2026-03", "2026-05-12", "2026-06-16", "2026-08-12", "2026-09", "2026-09-08"]
keywords: ["attention", "benchmark", "consumer", "cost", "decode", "deepseek", "disaggregated", "dram", "fp4", "fp8", "gguf", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3361, 3422]
section: "7. KV Cache & Long-Context Techniques"
sha256: 899c6d67293827c766304a139c9fc2d2a70db14e684a956e7426ffaff6309a86
---

# Key dated facts (continued)

- [COMMUNITY] https://temperature2.com/p/2026-09-08-did-you-know-multi-head-latent-attention/
- [COMMUNITY] https://vizuara.medium.com/what-exactly-is-multi-head-latent-attention-mla-da06e42f997f
- [COMMUNITY] https://github.com/ccomkhj/ccomkhj.github.io/blob/HEAD/_posts/2026-08-12-VariationOfMHA.md
- [COMMUNITY] https://medium.com/@htasoftware/inside-deepseeks-secret-weapon-multi-head-latent-attention-mla-explained-f359af04d38a
- [VENDOR] https://arxiv.org/abs/2604.05887
- https://aclanthology.org/2026.acl-long.594/
- https://arxiv.org/pdf/2604.05887v1.pdf
- [VENDOR] https://github.com/defilantech/llmkube/issues/308
- [COMMUNITY] https://github.com/captainbotgit/turboquant-mlx
- [COMMUNITY] https://www.marktechpost.com/2026/04/29/top-10-kv-cache-compression-techniques-for-llm-inference-reducing-memory-overhead-across-eviction-quantization-and-low-rank-methods/
- [DIRECTIONAL] https://cdn.prod.website-files.com/61a74a0b89162dfadb5acf21/69c6588d7694de16d5e965dd_TurboQuant%20%26%20The%20Memory%20Trade%20%E2%80%94%20Lighthouse%20Canton.pdf
- [COMMUNITY] https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_4.md
- [COMMUNITY] https://github.com/kostadis/dgx-fun/blob/HEAD/gemma4-31b-dense-spec.md
- [COMMUNITY] https://github.com/elizaos/eliza/issues/9033
- [VENDOR] https://huggingface.co/unsloth/MiniMax-M3-GGUF
- [VENDOR] https://the-decoder.com/minimax-m3-open-weight-model-with-a-million-token-context-challenges-proprietary-leaders/
- [VENDOR] https://www.techtimes.com/articles/318622/20260618/minimax-m3-takes-open-weight-ai-lead-sparse-attention-architecture-now-verified.htm
- [VENDOR] https://sambanova.ai/blog/minimax-m3-running-fastest-on-sambacloud
- https://github.com/sgl-project/sglang/blob/HEAD/docs/docs/advanced_features/quantized_kv_cache.mdx
- https://github.com/futuremls-lab/oscar/blob/HEAD/sglang-research/docs/advanced_features/quantized_kv_cache.md
- https://github.com/sgl-project/sglang/pull/21253
- https://github.com/sgl-project/sglang/pull/6109
- https://github.com/sgl-project/sglang/pull/30514
- [VENDOR] https://inferencex.semianalysis.com/blog/sglang-0-5-6-b200-deepseek-r1-fp4-up-to-1-8x
- https://github.com/vllm-project/vllm/pull/48250
- https://github.com/vllm-project/vllm/releases/
- [VENDOR] https://alphasignal.ai/news/vllm-v0-23-0-ships-deepseek-v4-production-hardening-and-56-throughput-boost
- [COMMUNITY] https://github.com/intel/containers/blob/HEAD/dockerfiles/vllm/release_notes/0.21.0-xpu.md
- [COMMUNITY] https://github.com/smart-lty/my-skills/blob/HEAD/model-pr-optimization-history/sglang/deepseek-v3-r1/README.en.md
- https://github.com/huggingface/text-generation-inference/releases
- https://huggingface.co/deepseek-ai/DeepSeek-V2-Lite-Chat
- [COMMUNITY] https://github.com/iopsystems/llm-calc/blob/HEAD/docs/superpowers/specs/2026-05-12-dsa-design.md
- [COMMUNITY] https://github.com/yusanxy/llm_flops/blob/HEAD/operators/references/deepseek_v32_dsa_sparse_attention/README.md
- [COMMUNITY] https://github.com/suzeai/transformers/blob/HEAD/docs/source/en/model_doc/glm_moe_dsa.md
- [COMMUNITY] https://earlyterms.com/term/indexshare
- [VENDOR] https://felloai.com/it/deepseek-v4/
- [COMMUNITY] https://github.com/neetx/ai-research-radar/blob/HEAD/reports/2026-06-16.md
- [COMMUNITY] https://github.com/pestopoppa/epyc-root/blob/HEAD/wiki/ssm-hybrid.md
- [COMMUNITY] https://github.com/mtgibbs/pi-cluster/blob/HEAD/docs/model-eval-2026-05.md
- [COMMUNITY] https://github.com/flexinfer/flexinfer/commit/39941ba27bc269ee7e2a7091abf82031477908b2
- [COMMUNITY] https://github.com/chtho-like/rosellm/blob/HEAD/docs/multimodal/kimi-k3.md
- [COMMUNITY] https://github.com/smfworks/smfworks-site/blob/HEAD/content/drj/kimi-k2-7-code-vs-minimax-m3-coding-benchmark.md
- [COMMUNITY] https://github.com/pinggy-io/pinggy_website/blob/HEAD/content/blog/best_open_source_self_hosted_llms_for_coding.md
- [COMMUNITY] https://github.com/maximiliankhan/openbeast/blob/HEAD/research/lowrank/prior-art/arxiv-kv-holistic.md
- [COMMUNITY] https://github.com/alkinun/speck/blob/HEAD/research/literature/16_deepseek_v2_mla.md
- [COMMUNITY] https://github.com/maraja/llm-evolution/blob/HEAD/09-the-cost-revolution-and-global-competition/01-deepseek_v2_and_mla.md

## Key dated facts (continued)

### Disaggregated KV storage: the cache as infrastructure (Wave 2 evidence)

- **LMCache v0.3.15 (March 2026)** is the reference 2026 implementation of disaggregated, engine-independent KV storage: a standalone daemon managing the KV cache outside the inference-engine process (GPU → pinned CPU DRAM → local disk/NVMe-GDS → remote), so cache survives engine crashes and restarts. vLLM integration runs through `KVTransferConfig` (`kv_connector="LMCacheConnectorV1"`, `kv_role="kv_both"`), requiring vLLM v1. Chunk granularity defaults to 256 tokens; the canonical CPU-offload config pins up to ~20 GB of host memory (`max_local_cpu_size: 20`).
- The documented LMCache CPU-offloading demo (quickstart docs, updated September 2026) reports a **7.43× speedup on the second run** with shared prefixes — where vLLM's in-GPU prefix caching alone scored 1.00× because the KV cache exceeded GPU memory and could not be reused at all.
- Eviction is LRU by default with pluggable LFU/FIFO. A controller API adds Pin/Unpin, Move, Compress (CacheGen), Clear, and Lookup.
- `min_retrieve_tokens` is the explicit pull-vs-recompute crossover: below a prefix-length floor, cold prefill is faster than restoring a small hit from a remote tier, so retrieval is skipped.
- LMCache extends reuse **beyond prefixes** via **CacheBlend** — cached KV blocks are blended at any prompt position, with selective token recompute for quality recovery — and supports KV transfer from prefill to decode workers over NVLink/RDMA/TCP via NIXL. It also exposes a pluggable SERDE interface for compression, token dropping, and custom serialization.
- Backend zoo (Sep 2026): Redis/Valkey (no RDMA), Mooncake Store (RDMA), InfiniStore, S3-compatible object storage, NIXL, GDS (GPUDirect Storage). The onnx-light-cpu roadmap (Aug 2026) tracks future INT8/FP8 KV for CPU-resident decode — the convergence of offload and quantization.
- **KTransformers** (kvcache-ai) pushes tiering to the consumer extreme: chunked prefill loads **one layer of KV into GPU memory at a time** and parks it on CPU/DRAM; during decode a **sparse CPU attention operator runs directly on the CPU-resident KV** instead of swapping. Reported outcome: local 128K–1M-token contexts feasible on consumer hardware (Sep 2026 design notes).
- **CacheGen** (UChicago, NSDI'24; arXiv:2310.07240) remains the reference for KV *encoding for the network* (not the GPU): custom quantization plus arithmetic coding exploiting token-wise locality and per-layer sensitivity, GPU-accelerated pipelined decoding, chunked streaming at multiple compression levels with fallback to raw text + recompute when bandwidth collapses. Measured: 3.5–4.3× less bandwidth at equal quality, 3.6–3.9× faster TTFT than the quantization baseline, 2.7–4.3× faster than text reload at <2% accuracy drop. LMCache's "Compress (CacheGen)" controller hook is its 2026 production descendant.

### RDMA-shared KV and the prefill-cost collapse at 1M+ tokens (Wave 2, VERIFIED)

