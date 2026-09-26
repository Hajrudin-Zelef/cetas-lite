---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/sources-and-urls
title: "Sources and URLs"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["DeepSeek", "MiniMax", "Nvidia", "SGLang", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["agentic", "alignment", "attention", "blackwell", "compute", "cost", "decode", "deepseek", "disaggregated", "fp4", "fp8", "inference"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3348, 3360]
section: "7. KV Cache & Long-Context Techniques"
sha256: ae7dfad2cd5e8e59c80f380a7091462189694f55d0de02be2b038801aa7292c1
---

# Sources and URLs

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

