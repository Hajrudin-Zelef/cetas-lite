---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/gemma-4-family-snapshot-community-verified-mid-2026
title: "Gemma 4 family snapshot (community-verified, mid-2026)"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["Alibaba", "DeepSeek", "Google", "Hugging Face", "MiniMax", "Moonshot", "Nvidia", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM"]
dates: ["2024-05", "2025-05", "2026-03-21", "2026-03-25", "2026-04-02", "2026-04-07", "2026-04-15", "2026-04-24", "2026-05-11", "2026-06", "2026-06-01", "2026-06-11", "2026-07", "2026-07-16", "2026-08-14", "2026-08-26", "2026-08-29", "2026-09-22"]
keywords: ["agent", "agentic", "apache", "attention", "benchmark", "benchmarks", "context window", "cost", "decode", "deepseek", "embedding", "embeddings"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3230, 3304]
section: "7. KV Cache & Long-Context Techniques"
sha256: b8e860d6fd9af5ed31a763c1ce4dd6232b25b5ed64d39e002e0cb6e1ee36901b
---

# Gemma 4 family snapshot (community-verified, mid-2026)

### Gemma 4 family snapshot (community-verified, mid-2026)

| Variant | Params | Arch | Context | Sliding window | Modalities |
|---|---|---|---|---|---|
| E2B | ~2.3B eff (5.1B w/ embeddings) | dense | 128K | 512 | text, image, audio (USM conformer) |
| E4B | ~4.5B (8B w/ embeddings) | dense | 128K | 512 | text, image, audio |
| 12B Unified | 11.95B dense, encoder-free, 262K vocab | dense | 256K | 1024 | text, image |
| 26B-A4B | 26B MoE, 4B active | MoE | 256K | 1024 | text, image |
| 31B | 30.7B text + 0.55B vision, dense, 60 layers, GQA 2:1 | dense | 256K | 1024 | text, image, video |

- E2B quantizes to ~1.3 GB at Q4 (2–3 GB RAM) — the mobile tier. The 12B Unified's 262,144-token vocabulary (tied embeddings ≈ 2.8 GB of the weight budget) is itself a 2026 design note: giant vocabs trade embedding memory for token efficiency on multilingual/agentic workloads. Community-verified KV math for the 31B at 256K context: ~40.8 GiB BF16 (dominated by the 10 global layers; the 50 local layers hold a constant ~0.78 GiB floor); FP8 KV → ~20.4 GiB. **Most of the cache is O(1) in context length; only global layers are O(n).** [COMMUNITY]

### DeepSeek-V4 serving specs (secondary-sourced unless noted)

| Variant | Total / active | Context | Output | License |
|---|---|---|---|---|
| V4-Pro | 1.6T / 49B (reported; treat exact totals as vendor-adjacent) | 1M | 384K | MIT |
| V4-Flash | 284B / 13B (reported) | 1M | 384K | MIT |

- Release 2026-04-24 (some secondary sources say 04-23 — immaterial disagreement). Carries the DSA lineage forward as hybrid **CSA + HCA**: V4-Pro at 1M tokens = **27% of V3.2's inference FLOPs and 10% of its KV-cache size**; V4-Flash = **10% FLOPs / 7% KV** [VENDOR]. MLA retained through the whole lineage (V2→V3→R1→V4) — the strongest endorsement of the mechanism.
- MLA serving caveat documented in 2026 recipes: MLA requires **DP+EP** for KV-cache efficiency (avoids 8× duplication across TP ranks) — a serving-topology constraint, not a dtype one.

### MiniMax M3 detail (vendor-reported where marked)

- **Coding/agent benchmarks [VENDOR]:** 59.0% SWE-Bench Pro, 66.0% Terminal-Bench 2.1, 74.2% MCP Atlas; ~100T training tokens on interleaved text/image sequences. Independent: Artificial Analysis ranked M3 #1 open-weight on its Intelligence Index (score 55, June 2026).
- **Agentic long-context training:** MiniMax built a simulator framework mimicking real developer behavior (refining requirements, discussing approaches, reacting to intermediate results, carrying tasks across multiple contexts) — multi-turn collaboration during training, not single prompts. Internal test: M3 ran autonomously ~24 h optimizing an FP8 GEMM CUDA kernel, lifting hardware utilization from 7.6% to 71.3% (9.4×) with zero human intervention (reported via SambaNova) [VENDOR, via SambaNova].
- **Two reasoning modes:** **thinking** (complex/agentic tasks) and **non-thinking** (latency-sensitive chat/completion) — the mode switch determines how much trace (and hence KV) the reasoning-aware compressors (ThinKV-class) have to work with.
- **Philosophy vs MLA, recorded verbatim for the consolidation:** MSA operates on **uncompressed** KV data, preserving long-context retrieval accuracy at the cost of slightly higher memory — MiniMax's argument is that this matters for tasks where subtle details distributed across very long documents must be recovered precisely. This is the "fidelity vs memory" axis the consolidated document should frame the MLA-vs-MSA choice on.


### More dated facts: models that carry the attention story

- **2026-04-24 — DeepSeek V4 (Pro/Flash + bases) released** (some secondary sources say 04-23 — immaterial): 1M context, MIT license, DSA-derived CSA+HCA; MLA retained through the whole V2→V3→R1→V4 lineage. V4-Pro = 1.6T total / ~37–49B active per token (exact totals vendor-adjacent); V4-Flash = 284B / 13B. Artificial Analysis (2026-09-22): V4 family variants among top-scored DeepSeek models.
- **2026-06 (mid) — MiniMax M3 quality signals:** Artificial Analysis Intelligence Index score 55 (#1 open-weight, June 2026); vendor-reported coding figures 59.0% SWE-Bench Pro, 66.0% Terminal-Bench 2.1, 74.2% MCP Atlas; ~100T training tokens on interleaved text/image sequences; weights + tech report (arXiv:2606.13392) published ~June 11, 2026.
- **2026-07-16 — Moonshot launch chart:** Kimi K3 at ~2.8T total parameters, routed pool 896 experts / top-16 (vs K2's 384 / top-8); Intelligence Index 60 (tied best open weight with GLM-5.3). Kimi K2 series = MLA adopters; K2.7 Code measured at 256K context in independent coding benchmarks (June 2026) [COMMUNITY]. Exact K3 maximum context window [UNVERIFIED] — do not carry an implied 1M claim for Kimi without a source.
- **2026-08-14 / 2026-08-29 — GLM-5.3:** announced 2026-08-14; weights held two weeks for a safety review (unexpected offensive-security scores) then published 2026-08-29 — 756 GB native FP8 across 141 files, bespoke GLM-5.3 license; Intelligence Index 60 (tied best open weight); needs H200-class hardware. MLA adopter per the MLA census (temperature2.com). Z.AI's two-week safety hold is a 2026 pattern: capability-gated releases even in the open-weight world.
- **2026-08-26 — GLM-5.3-Flash** (MIT license, 320B/18B, natively multimodal): the one-server-deployable variant (fits one 8-GPU node); Intelligence Index 57.
- **2026-08 — Qwen3.8-27B** (27B dense): context spec not verified in the sources checked — do not carry a 1M implication for Qwen.
- **General caveat:** "1M context" claims are frequently *vendor-reported usable-window* figures; independent needle-in-haystack validation at the full 1M length remains rare across all vendors. Most published NIH-at-1M scores are vendor-run.

### Architectural choice matrix: which KV strategy for which goal

| Goal | Best 2026 answer | Why | Cost |
|---|---|---|---|
| Max cache cut, no quality loss | MLA (pretraining) | 93.3% cut, quality ≥ MHA; 8+ families adopted | One-way door; MLA-aware kernels needed |
| Max retrieval fidelity at 1M | Sparse attention (MSA) | Uncompressed KV blocks; vendor-claimed NIH parity | Higher memory than MLA; vendor-specific |
| Cheap long context on small models | Shared KV + SWA interleave (Gemma 4 style) | Most cache O(1) in seq len; 128–256K on edge GPUs | Only ~1/6 layers see full context |
| Longest context, no retraining | TurboQuant 3-bit (serving) | 6× memory, lossless NIH, any model, no calibration | Not yet upstream in vLLM; needs integration |
| Extreme batch throughput, stable workload | KVQuant calibrated sub-4-bit | Beats fixed grids; evaluated to 10M ctx | Calibration step; workload-specific |
| Multimodal (image/video tokens) | HybridKV | 7.9× on MLLMs; head-heterogeneous beats uniform budgets | Research-stage; Qwen2.5-VL validated |
| Agentic loops, shared prompts | Prefix/persistent caching | 95%+ hit rates; ~9× effective input-cost cut | → sibling 07b (gateway + hash-convention complexity) |
| Lowest latency decode at scale | P/D disaggregation + decode hardware | Hardware-level P/D split | → sibling 07b |


## Main actors

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

