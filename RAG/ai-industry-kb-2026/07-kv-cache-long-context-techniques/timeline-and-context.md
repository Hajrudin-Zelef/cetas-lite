---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/timeline-and-context
title: "Timeline and context"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["Alibaba", "DeepSeek", "Google", "Hugging Face", "MiniMax", "Mistral", "Moonshot", "Nvidia", "SGLang", "TensorRT-LLM", "United States", "Z.ai", "vLLM"]
dates: ["2024-05", "2025-03-16", "2025-05", "2025-10-29", "2026-01-22", "2026-03-21", "2026-03-25", "2026-03-31", "2026-04-02", "2026-04-04", "2026-04-07", "2026-04-15", "2026-04-24", "2026-05-07", "2026-05-11", "2026-06-01", "2026-06-12", "2026-07", "2026-08-14", "2026-08-26", "2026-08-29", "2026-09-08", "2026-09-17", "2026-09-22"]
keywords: ["apache", "attention", "benchmark", "benchmarks", "compute", "deepseek", "fp4", "fp8", "glm", "gpu", "gqa", "inference"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3286, 3347]
section: "7. KV Cache & Long-Context Techniques"
sha256: 9a5382970edbcea4cbb61f2729c9ded6ba78c0b596754c037a1d3c1301c35e36
---

# Timeline and context

| Actor | Role in this part | Key dated signal |
|---|---|---|
| **DeepSeek** | MLA origin; sparse-attention lineage (DSA→CSA/HCA); FP8-KV origin (FlashMLA) | May 2024 V2 paper (93.3%); Feb 2025 FlashMLA (656 B/token); Dec 2025 V3.2 DSA; Apr 2026 V4 (27% FLOPs / 10% KV vs V3.2) |
| **Google Research** | **TurboQuant inventor** (NOT Red Hat); PolarQuant (AISTATS 2026); Gemma 4 shared KV | 2026-03-25 TurboQuant announcement (ICLR 2026); 2026-04-02 Gemma 4 |
| **Red Hat** | Independent vLLM evaluation of TurboQuant (llmkube#308, 6×/8× figures) — evaluator, not inventor | 2026-05-11 independent evaluation |
| **Moonshot AI (MiniMax lineage / Kimi)** | MiniMax M3 + MSA; Kimi K2 MLA adopters; K2.7 Code 256K independent benchmark | 2026-06-01 M3 launch (arXiv:2606.13392); Kimi K3 Intelligence Index 60 (tied best open weight, July 2026) |
| **Zhipu (Z.AI)** | GLM-5 / GLM-5.3 MLA adopters; GLM-5.2 IndexShare | GLM-5.3 announced 2026-08-14, weights 2026-08-29 (safety hold); GLM-5.3-Flash 2026-08-26 (MIT, 1M); IndexShare Jun 2026 |
| **NVIDIA** | TensorRT-LLM MLA; Nemotron-3-Ultra-550B-A55B (Mamba-2+MoE+Attention "LatentMoE"); TRTLLM-gen kernel fed back into vLLM v0.23.0; NVFP4 KV path | Jun 2026 Nemotron-3-Ultra (arXiv:2606.15007); OpenMDW-1.1 license |
| **Alibaba (Qwen)** | Gated DeltaNet hybrids: Qwen3-Next (zero-KV recurrent layers), Qwen3.6-35B-A3B, Qwen3.5-MoE | Qwen3.5-MoE Feb 2026; Qwen3.6-35B-A3B ~2026-04-15 (Apache-2.0) |
| **IBM** | Granite-4.0-H-Small: hybrid Mamba-2 / Transformer MoE (32B / 9B active) | 2026 |
| **AI21** | Jamba-1.5 (1:7 attention:Mamba hybrid; 4–9 GB KV at 256K) | 2026 |
| **Zyphra** | Zamba2-VL (Mamba2–Transformer hybrid VLM, ~10× TTFT cut) | Jun 2026 |
| **vLLM project / community** | MLA backends (`FLASHMLA`, `FLASHINFER_MLA` + sparse, `TRITON_ATTN`, `CUTLASS_MLA`, `XLA`); fp8_e4m3 production recipes; sparse MLA end-to-end in v0.28.0 | v0.6.x late 2024 (MLA, [UNVERIFIED] exact minor); v0.23.0 2026; v0.28.0 tagged 2026-08-26 |
| **SGLang project / community** | Earliest MLA serving (v0.3, Sep 2024); FlashMLA+FP8-KV+MTP (PR #6109, May 2025); broadest 2026 KV-dtype matrix incl. experimental FP4 (v0.5.6, Dec 2025) | 2026-09 SGLang quantized-KV docs |
| **Zhejiang University / Hangzhou institutes** | HybridKV (VERIFIED, ACL 2026): 7.9× multimodal KV reduction | arXiv:2604.05887 (2026-04-07); ACL 2026 long paper 2026.acl-long.594 |
| **Sebastian Raschka (2026 architecture review)** | Re-summarized (Aug 2026): MLA modeling quality matches/slightly beats MHA, unlike GQA which measurably underperforms MHA head-to-head | Aug 2026 re-summary |
| **SemiAnalysis (InferenceX)** | Measured NVFP4 on B200: DeepSeek-R1 907 tok/s/GPU via SGLang 0.5.6 | Dec 2025–Jan 2026 |
| **Hugging Face (TGI)** | Counter-case: MLA never reached mainline CUDA; archived 2026-03-21 | TGI v3.3.x Gaudi-branch only (late 2025); archived 2026-03-21 |

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

