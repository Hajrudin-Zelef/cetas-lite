---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/figures-and-metrics-continued
title: "Figures and metrics (continued)"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM"]
dates: ["2025-05", "2026-02-19", "2026-03-21", "2026-04-15", "2026-04-24", "2026-05-28", "2026-06", "2026-06-01", "2026-06-13", "2026-08", "2026-08-26", "2026-09-17"]
keywords: ["accelerator", "agent", "agentic", "amd", "apache", "attention", "benchmarks", "claude", "consumer", "cost", "decode", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3521, 3638]
section: "7. KV Cache & Long-Context Techniques"
sha256: 28c415e10bf341ab797bdc102ccc03cb0cec006eaebf3a97e12e73c66e7754da
---

# Figures and metrics (continued)

## Figures and metrics (continued)

### A. Disaggregated-KV measured figures

| System | What was measured | Figure | Date / source |
|---|---|---|---|
| LMCache CPU offload | 2nd-run speedup, shared prefix, KV > GPU memory (vLLM alone: 1.00×) | **7.43×** | Sep 2026, LMCache docs |
| CacheGen | TTFT vs quantization baseline, 7B–70B models | **3.6–3.9×** faster; 3.5–4.3× less bandwidth | NSDI'24 |
| Lightbits RDMA paging | Turn-2 TTFT, Llama-70B FP16 KV | 8K: 1199→**38.8 ms**; 32K: 4935→**88.7 ms**; 128K: ~23 s→**~300 ms** | 2026-09-17 |
| Lightbits + FarmGPU/ScaleFlux | TTFT, 131K–1M tokens | **100–280×**; **>1,000× at 10M** | Sep 2026 |
| Lightbits | 10.5M-token session restore on 8×H100 (vs 1h42m prefill rebuild) | **seconds, bit-identical output** | 2026-09-17 |
| llm-d p2p KV sharing | Prefill-latency delta, gpt-oss-120b, RDMA/IB | **−55.8% @2K → −88.2% @49K** (pull wins at every length) | Sep 2026 |
| Mooncake Store vs Redis | TTFT, vLLM P/D disaggregation (Tsinghua ToS 2025) | **TCP: 17–22%; RDMA: 26–33%** | ToS 2025 |
| Mooncake (Kimi prod) | Capacity at same SLOs (A800 / H800) | **+115% / +107%** | Paper era (arXiv:2407.00079) |
| Mooncake (simulated) | Throughput at matched SLOs | **50–525%** (conv. workload +498% @100 ms TBT) | Paper era |
| SGLang disagg. | 96 H100s, DeepSeek model | **52,300 input + 22,300 output tok/s per node** | 2026 survey |
| KTransformers | Local consumer-hardware context via CPU sparse attention | **128K–1M tokens** | Sep 2026 |
| vLLM v0.26.0 | Release scale | 411 commits, 212 contributors | Release notes |
| DeepSeek-R1 FP8-KV | KV moved over RDMA for an 8,192-token prefill | **~290 MB** | 2026 survey |
| NVIDIA Wide-EP | Per-GPU output throughput gain from EP8→EP64 (32→4 experts/GPU, HBM freed for KV) | **up to 2.28×** [VENDOR] | 2026 |

### B. Per-engine KV dtype support matrix (Wave 2.1 — carried forward here)

| Engine | `--kv-cache-dtype` values | Notes |
|---|---|---|
| **vLLM** (2026 mainline) | `auto \| fp8 \| fp8_e4m3 \| fp8_e5m2` | Production recipes pin `fp8_e4m3`; `--kv-cache-dtype-skip-layers`; `--mamba-cache-dtype` has no FP8 option |
| **SGLang** (docs, Sep 2026) | `fp8_e5m2 \| fp8_e4m3 \| fp4_e2m1 \| nvfp4 \| fp4_mx_block16 \| int8 \| int4 \| int2` | Broadest 2026 matrix; FP4 experimental (MXFP4 E2M1, per-block-16); INT via Triton, dynamic affine per-token scales |
| **TensorRT-LLM** | FP8 KV documented | Exact introduction release [UNVERIFIED]; "documented, version unpinned" |
| TGI | — | Archived 2026-03-21; no mainline-CUDA MLA/KV work (Gaudi branch only) — the counter-case |

SGLang capacity accounting: **FP4 ≈ 1.78× more tokens than FP8, ≈ 3.56× more tokens than BF16** (scale overhead included).
FP8-KV production benchmarks (vLLM, Apr 2026): **ITL slope 54% of BF16 on H100; break-even ~7K tokens; sub-0.3% accuracy cost**.
Hopper landmine: **128K NIAH 91% → 13%** in early FP8-KV implementations; fix = two-level FP32 accumulation (flash-attention#96/#91).

### C. Quantized-KV production-adoption datapoints

| Milestone | Detail |
|---|---|
| Feb 2025 | DeepSeek FlashMLA: custom FP8 KV format (**656 bytes/token**), the origin point of FP8 KV in production |
| May 2025 (SGLang PR #6109) | FlashMLA + FP8 KV + MTP: **~30% serving speedup, KV cache halved** on DeepSeek-V3/R1-class models |
| Dec 2025–Jan 2026 (SemiAnalysis InferenceX) | DeepSeek-R1 NVFP4 on B200 via SGLang 0.5.6: **907 tok/s/GPU** at concurrency 4 (1.79× over 0.5.5) |
| 2026 vLLM recipes | DeepSeek-R1 (8× GH200/H200, DP+EP) and DeepSeek-V4-Flash (4× GH200): **`--kv-cache-dtype fp8_e4m3`** with MLA backends |
| TurboQuant (Mar 2026) | 3-bit KV: **≥6× memory cut, up to 8× attention speedup (H100)**, 99.5% fidelity, lossless NIH |
| Engine divergence (AMD fork) [SECONDARY] | SGLang fp8_e4m3: 2× KV pool + faster decode (fused kernels); vLLM-ROCm fp8: ~0.5× decode vs fp16 (unfused Triton dequant) |

### D. The 2026 quantization ladder — one user, one request (worked sizing)

Llama-3.3-70B-class (80 layers, GQA-8, head_dim 128), per-request KV:

| Context | BF16 | FP8 (2×) | INT4 (4×) | TurboQuant 3-bit (~5.3×) |
|---|---|---|---|---|
| 4K | ~1.3 GB | ~0.7 GB | ~0.3 GB | ~0.2 GB |
| 128K | ~42 GB | ~21 GB | ~10.5 GB | ~8 GB |
| 1M | ~336 GB | ~168 GB | ~84 GB | ~56 GB |

SNIA SDC25 anchors: Qwen3-8B **150 KB/token FP16** (147 GB at 1M); LLaMA-3.3-70B **330 KB/token** (327 GB at 1M); LLaMA-3.1-405B **516 KB/token** (516 GB at 1M). Reading the table: at 1M tokens *no single lever* fits on one 80GB GPU — the 2026 answer is stacking (GQA + FP8 + prefix reuse + disaggregation) or architectural choices made at training time. (Architectural levers — MLA's ~93.3% cut, shared KV, sparse attention — are covered in part 07a; hybrid O(1)-state cuts of 8–10× are covered there as well.)

### E. Per-accelerator HBM: what 1M tokens costs (vendor-listed specs, Sep 2026)

1M-token KV for a Llama-70B-class model in FP16 (~40 GB) as a fraction of physical HBM:

| Accelerator | HBM | Bandwidth | One user's 1M-token FP16 KV as fraction |
|---|---|---|---|
| H100 SXM | 80 GB HBM3 | ~3.35 TB/s | **~50%** |
| H200 | 141 GB HBM3e | ~4.8 TB/s | ~28% |
| B200 | 192 GB HBM3e | ~8 TB/s | ~21% |
| GB200 (2×B200) | 384 GB HBM3e | ~16 TB/s | ~10% |
| MI300X | 192 GB HBM3 | ~5.3 TB/s | ~21% |
| MI325X | 256 GB HBM3e | ~6 TB/s | ~16% |
| MI350X/355X | 288 GB HBM3e | ~8 TB/s | ~14% |

Two readings: (1) on an H100, FP16 KV for *one* 1M-token user eats half the GPU — FP8 KV (→~21%) and hybrid 8–10× cuts (→~5%) are the difference between serving and not serving; (2) the HBM-generation race (HBM3 → HBM3e → HBM4) is the physical substrate of the supercycle pricing — each generation step is a capacity auction the inference industry must win.

### F. 1M-context leaderboard snapshot (flagship tiers, Sep 2026)

| Model | Input context | Release | Max output | License / note |
|---|---|---|---|---|
| Gemini 3.1 Pro | 1M | 2026-02-19 (preview) | 64–66K | Closed |
| DeepSeek V4-Pro | 1M | 2026-04-24 | 384K | MIT; 1.6T/49B; 27% FLOPs / 10% KV of V3.2 [VENDOR] |
| DeepSeek V4-Flash | 1M | 2026-04-24 | — | MIT; 284B/13B; 10% FLOPs / 7% KV [VENDOR] |
| Claude Opus 4.8 | 1M | 2026-05-28 | 128K | Closed; Foundry may cap at 200K on some tiers; secondary-sourced |
| GLM-5.2 | 1M | ~2026-06-13 (API) / ~06-17 (weights) | 65K | IndexShare 2.9× FLOP cut [VENDOR] |
| NVIDIA Nemotron-3-Ultra | 1M | June 2026 | — | OpenMDW-1.1; Mamba-2+MoE+Attn hybrid |
| GPT-5.6 family | 1.05M | August 2026 | — | With pricing cut |
| MiniMax M3 | 1M (512K floor) | 2026-06-01 | — | Open weights; sparse attention |
| GLM-5.3-Flash | 1M | 2026-08-26 | — | MIT; 320B/18B; single 8-GPU node |
| Qwen3.6-35B-A3B | **256K** (not 1M) | ~2026-04-15 | — | Apache-2.0; Gated-DeltaNet hybrid — the tier exception |
| Gemini 3.5 | [UNVERIFIED] | — | — | No citable source |

SGLang capacity accounting: **FP4 ≈ 1.78× more tokens than FP8, ≈ 3.56× more tokens than BF16** (scale overhead included).
FP8-KV production benchmarks (vLLM, Apr 2026): **ITL slope 54% of BF16 on H100; break-even ~7K tokens; sub-0.3% accuracy cost**.
Hopper landmine: **128K NIAH 91% → 13%** in early FP8-KV implementations; fix = two-level FP32 accumulation (flash-attention#96/#91).

### G. Transfer-decision matrix and KV-event plumbing (measured anchors)

| Topology | Transfer choice | Measured anchor |
|---|---|---|
| Co-located P/D (same node) | None — keep KV local | N/A (baseline) |
| Cross-node P/D | NIXL / RDMA transport | llm-d: −55.8%→−88.2% prefill deltas; pull beats push at every length |
| Remote / cold tiers | CacheGen-compressed bitstreams over object storage | 3.5–4.3× less bandwidth; fallback to text+recompute |
| Open KV standard (aspirational) | Open KV Cache API (deterministic block identity, tenant/session scoping) | No engines interoperate yet — "a cache index is a convention; an interface is a contract" |

Digest conventions (the interop hazard): requests carry **SHA-256 (or XXH3-128) hashes of canonical KV block digests**; vLLM truncates to the **last 8 bytes big-endian**, SGLang uses the **first 8** — the documented mis-slice produced 0% cache overlap before being fixed. KV events flow over **ZMQ PUB sockets** (one per prefill endpoint in vLLM; one per DP rank in SGLang) with msgpack envelopes.

### H. Prefix-reuse economics — worked example (OpenCode-style /data, June 2026)

Agent loop: 20 calls against a 100K-token repo context. Full prefill once (~100K uncached tokens) + ~1.9M cached tokens; at typical **10:1 cached:uncached pricing → ~9× effective input-cost cut** vs no caching. Multiplicative composition: a 6×-compressed (TurboQuant-class) cache reused across 20 agent turns ≈ **~100× prefill saving** vs naive per-call full prefill. This is why providers price cached input tokens separately and production traces show **95%+ hit ratios** on agentic workloads.

### I. Prefill→decode KV transfer cost math (Wave 1, §7.3)

| Request (Llama 3 70B, BF16) | KV moved P→D | Reading |
|---|---|---|
| 4K context | **~13.4 GB** | For short prompts, transfer latency can dominate total TTFT |
| 128K context | ~430 GB | Long contexts: the price of disaggregation is the transfer itself |
| Same at FP8 KV | ~215 GB | Compression halves the disaggregation tax — the coupling |

This is why KV *compression* and disaggregation are coupled design problems, not separate ones: a 6×-smaller cache is a 6×-cheaper transfer, and every transfer-cost figure above shrinks multiplicatively with the dtype and compression choices in §B–§D.

