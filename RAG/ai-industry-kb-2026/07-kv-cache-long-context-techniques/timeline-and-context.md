---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/timeline-and-context
title: "Timeline and context"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["Alibaba", "DeepSeek", "Google", "Hugging Face", "MiniMax", "Mistral", "Moonshot", "Nvidia", "SGLang", "TensorRT-LLM", "United States", "Z.ai", "vLLM"]
dates: ["2024-05", "2025-03-16", "2025-10-29", "2026-01-22", "2026-03-21", "2026-03-25", "2026-03-31", "2026-04-02", "2026-04-04", "2026-04-07", "2026-04-24", "2026-05-07", "2026-05-11", "2026-05-12", "2026-06-01", "2026-06-12", "2026-06-16", "2026-08-12", "2026-08-26", "2026-08-29", "2026-09-08", "2026-09-17", "2026-09-22"]
keywords: ["agentic", "alignment", "attention", "benchmark", "benchmarks", "blackwell", "compute", "cost", "decode", "deepseek", "disaggregated", "fp4"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3305, 3407]
section: "7. KV Cache & Long-Context Techniques"
sha256: dc32d9211c93676123d5af536eca5aea95d528853be4739d524b217b0e66ba4b
---

# Timeline and context

## Timeline and context

| Date | Event | Relevance |
|---|---|---|
| May 2024 | DeepSeek-V2 paper (arXiv:2405.04434) introduces MLA | 93.3% KV reduction, 5.76× throughput; the architectural shot heard around inference |
| Sep 2023–2024 | Llama 3, Mistral standardize GQA | GQA becomes the open-weight default (~88% vs MHA) |
| Nov 2023 | H2O (NeurIPS 2023) | Heavy-hitter eviction; 29× throughput on OPT models |
| 2024 | KIVI (ICML 2024) | 2-bit KV quantization; 2.6× combined memory cut |
| 2024 | GEAR (arXiv:2403.05527) | X ≈ D̂ + L + S error-correction canon for KV quantization |
| Feb 2025 | DeepSeek FlashMLA | Custom FP8 KV (656 bytes/token); origin point of FP8 KV in production |
| Sep 2024 | SGLang v0.3: "7× Faster DeepSeek MLA" | Earliest MLA serving on a major engine |
| Late 2024 | vLLM v0.6.x-era: DeepSeek-V2 MLA (Triton path) | MLA in vLLM; exact minor [UNVERIFIED] |
| 2025-03-16 | SGLang FlashMLA backend (PR #4079) | MLA backend family in SGLang |
| 2025-05 | SGLang PR #6109: FlashMLA + FP8 KV + MTP | ~30% speedup, KV halved |
| 2025-10-29 | KVLinC (arXiv:2510.05373) | Hadamard rotation + trainable linear-correction adapters |
| 2025-11 | Kitty (arXiv:2511.18643) | 2-bit KV with dynamic channel-wise precision boost |
| 2025-12 | DeepSeek-V3.2 introduces DSA (lightning indexer, top-K 2048) | Sparse attention as compute lever — not automatically a cache lever |
| 2025-12 | SGLang v0.5.6: MHA+MLA KV caches refactored for FP4 | FP4/NVFP4 KV on SM100; experimental |
| 2026-01-22 | R-KV v4 (arXiv:2505.24133v4) | Reasoning-aware compression: 10% cache, ~100% performance, 6.6× throughput |
| 2026-01 | vLLM v0.7.x hardening of DeepSeek-V3 MLA | Production-grade MLA serving |
| 2026-02 | Qwen3.5-MoE ships (Gated-DeltaNet/Mamba + 256-expert MoE hybrid) | Recurrent+MoE in production open weights |
| 2026-03-21 | Hugging Face TGI archived read-only | MLA never mainline CUDA → "all engines" framing = 3/4 |
| 2026-03-25 | Google Research blog announces TurboQuant (ICLR 2026) | 3-bit KV, 6× memory, 8× attention speedup; semiconductor selloff |
| 2026-03-31 | "Thin Keys, Full Values" (arXiv:2603.04427) | 75% key-cache retrofit savings, no pretraining door needed |
| 2026-04-02 | Gemma 4 ships (shared KV cache, interleaved SWA/global) | Cross-layer KV sharing in a mainstream open family |
| 2026-04-07 | HybridKV (arXiv:2604.05887) | 7.9× KV cut for multimodal models |
| 2026-04 | vLLM FP8-KV benchmarks | 54% ITL slope, break-even ~7K tokens, sub-0.3% accuracy |
| 2026-04-04 | FoveatedKV independent benchmark | 75% KV cut, ≤3% quality loss (top-10% fp16, rest fp8/INT4) |
| 2026-04-24 | DeepSeek V4 (Pro/Flash; DSA-lineage CSA+HCA; 1M context; MIT) | 27% FLOPs / 10% KV vs V3.2 at 1M (V4-Pro); 10%/7% (V4-Flash) [VENDOR] |
| 2026-05-07 | ThinKV v2 (ICLR 2026 Oral per Sep survey) | Thought-adaptive: <5% KV cache, 5.8× throughput |
| 2026-05-11 | Red Hat independent vLLM evaluation of TurboQuant (llmkube#308) | Confirms 6×/8×; Red Hat is the evaluator, not the inventor |
| 2026-06-01 | MiniMax M3 launches (MSA) | Sparse-attention alternative to MLA; uncompressed KV, NIH-fidelity argument |
| 2026-06 | GLM-5.2 ships IndexShare | Shared indexer every 4 layers; 2.9× FLOP cut at 1M; 1M input/65K output |
| 2026-06-12 | NVIDIA Nemotron-3-Ultra-550B-A55B (arXiv:2606.15007) | First US-lab frontier-scale open-weight Mamba-2-hybrid MoE |
| 2026-07 | ACL 2026 (San Diego): HybridKV long paper (2026.acl-long.594) | Peer-reviewed confirmation of the 7.9× claim |
| 2026-08-26 | vLLM v0.28.0 tagged | Sparse MLA end-to-end; tiered KV offload to disk |
| 2026-08-29 | GLM-5.3 weights published (safety hold lifted) | MLA-adopter family at frontier-open scale |
| 2026-09-08 | temperature2.com MLA analysis (8+ adopter families; 14.2× toy arithmetic) | MLA census: Kimi K2, GLM-5, six+ more |
| 2026-09-17 | Lightbits RDMA-paged KV post (10.5M-token session restore) | Sibling 07b territory — cross-ref only |
| 2026-09-22 | Consolidation cutoff | FP8 E4M3 = production KV dtype; MLA census ≥8 families; FP4 KV experimental |

## Implications

1. **KV-cache is the 2026 inference battleground.** Weight quantization (FP8/INT4) is commoditized; the frontier moved to the cache: MLA-style architectural compression, 3-bit online quantization (TurboQuant), cross-layer sharing (Gemma 4), sparse attention (MSA/DSA/IndexShare), and hybrid O(1)-state layers.
2. **Pretraining locks in your cache economics.** MLA, shared KV, MSA, Gated DeltaNet must be chosen before training — one-way doors. Serving teams inherit these decisions; model selection *is* infrastructure selection. The consolation for the installed base: post-hoc low-rank retrofits ("Thin Keys, Full Values," 75% key savings, <1% pretraining data) give MLA-class savings without the pretraining door.
3. **Default to FP8 E4M3 KV in production serving — with a validation gate.** ~2× capacity, 54% ITL slope on H100, break-even ~7K tokens, <0.3% accuracy cost. But: verify FlashAttention-3 two-level FP32 accumulation (`flash-attention#96`, `#91`) is present on Hopper *before* routing traffic, and run a 128K NIAH probe after every KV-dtype change — the 2026 91%→13% collapse was silent. Pin `fp8_e4m3`, not bare `fp8`, in DeepSeek-class serving recipes.
4. **Do not overclaim FP4/NVFP4.** They remain Blackwell/SM100-native or experimental (SGLang v0.5.6 experimental flag; SGLang FP4 ≈ 1.78× tokens vs FP8, 3.56× vs BF16 on paper). INT4 KV quality is method- and setup-dependent. The production ladder reads: BF16 → FP8 (free) → INT4 (4×) → TurboQuant 3-bit (5.3×).
5. **Do not equate sparse attention with KV-cache reduction.** DSA/MSA/IndexShare reduce attention *compute*; unless paired with explicit compression (V4's CSA/HCA), the cache must retain full-fidelity KVs for future indexer selection. Audit vendor claims on which quantity shrank.
6. **For retrieval-critical long context, prefer architectures that keep KV uncompressed — or validate yourself.** MiniMax's MSA argument (uncompressed blocks → NIH fidelity) vs MLA's latent compression is an unresolved fidelity tradeoff; independent NIH-at-1M validation remains rare across vendors. For agentic/legal/medical workloads, token-eviction methods permanently discard context — prefer *lossless* compression (TurboQuant-class) where compliance demands it.
7. **Reasoning traces are the new KV growth vector.** ThinKV (<5% cache, 5.8× throughput) and R-KV (10% cache, 6.6× throughput) show the 2026 frontier moving from prompt-cache compression to *thought*-cache compression — directly relevant as agentic workloads push reasoning traces past 100K tokens.
8. **Cite mechanisms, not slogans.** "Dynamic drift correction" names no documented technique — the real mechanisms are GEAR's X ≈ D̂ + L + S, MiKV's runtime Dynamic Outlier Awareness, ResQ's PCA-gated splits, WKVQuant's 2D alignment, and TurboQuant's QJL unbiased residual correction. Name one of them.
9. **Hybrids are the capacity answer; attention remains the recall answer.** 8–10× KV cuts are measured (Jamba-1.5; worked hybrid arithmetic), but exact-recall quality is the documented cost; the converged 2026 design is hybrid-by-ratio, not SSM-purist. Pure-SSM retrieval parity stays [UNVERIFIED] until needle-style evaluations.
10. **Cross-refs to sibling coverage (one line each).** → **Part 06 (inference engines):** full per-release MLA kernel matrices (vLLM backend-priority rules, SGLang PR-level KV-dtype history, TensorRT-LLM pinning). → **Part 07b:** disaggregated/tiered KV storage (LMCache 7.43×, Mooncake, CacheGen), prefill/decode disaggregation, RDMA-shared KV (Lightbits, llm-d −88.2%), prefix caching (RadixAttention, vLLM), and 1M-context product status.

## Sources and URLs

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

