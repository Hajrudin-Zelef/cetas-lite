---
id: ai-industry-kb-2026/09-moe-architectures/serving-moe-at-scale-expert-parallelism-load-balancing-quant
title: "Serving MoE at scale: expert parallelism, load balancing, quantization"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["Alibaba", "DeepSeek", "Meta", "Nvidia", "Z.ai"]
dates: []
keywords: ["moe", "quantization", "attention", "blackwell", "compute", "cost", "decode", "deepseek", "fp4", "fp8", "glm", "gpu"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5041, 5051]
section: "9. MoE Architectures"
sha256: bc622a23e11ad3f9cc0471491153d50d3a9df712bd7fe9f9d58460a9dc69d5b2
---

# Serving MoE at scale: expert parallelism, load balancing, quantization

- **What MoE saves:** per-token FLOPs (∝ active params) and decode-time memory bandwidth (each token reads only active/total of weights). Canonical illustrations: DeepSeek-V3 671B/37B reads ~5% of weights per token, matching dense-30B-class throughput on H800; Qwen3.6-35B-A3B measured at **61 tok/s vs 7 tok/s for a 27B dense model on the same RTX 5070** (practitioner report, 2026-04 era) — 3B active vs 27B dense, ~9× less per-token compute. At batch=1, decode is memory-bandwidth-bound, so MoE wins strongly.
- **What MoE does NOT save — verified across sources:**
  1. **VRAM for weights scales with TOTAL params.** All experts must reside on GPU because the router can select any expert for the next token; PCIe/NVLink on-demand loading latency far exceeds expert compute latency. Measured serving profile (OLMoE-1B-7B vs OLMo-1B dense, same family): active params agree within 8.9%, but resident memory differs **5.9×**; quantization scales both columns equally (ratio invariant: 5.40× at fp32 and int4) — **quantizing a MoE yields a smaller MoE, never a dense-sized model**. A probe of Switch-Base-8 showed all 8 experts receiving tokens within 223 tokens (worst-layer coverage 100%) — even top-1 routing keeps every expert hot.
  2. **KV cache is attention's cost, not FFN's.** MoE replaces the FFN; KV cache size = f(layers × KV heads × head dim × sequence length). "Proportional to layer count" with no pure MoE savings — VERIFIED (NVIDIA technical blog, community serving profiles, MoE tutorials). Every 2026 KV reduction is attention-side: MLA (DeepSeek-V2: KV cache ~5% of LLaMA-3-70B), GQA (fewer KV heads), sliding-window hybrids (MiMo: ~6× at 5:1 SWA/GA; GLM-5.3-Flash: 4.4× claimed via hybrid sparse+linear attention), shared KV cache (Gemma 4), HybridKV compression (up to 7.9×), TurboQuant (3-bit KV, ÷6 memory).
  3. **Less room for KV cache than a same-size dense model.** Because all experts occupy VRAM, the KV pool left over is smaller — NVIDIA's technical blog flags this explicitly; it is why MoE serving leans so hard on KV-compressing attention.
- **Batch-size behavior:** as batch grows, tokens collectively touch most experts, so the bandwidth advantage narrows (per-token work reduction persists); MoE holds a throughput advantage across batch sizes but its latency margin over a well-optimized dense model compresses at high concurrency. Conversely, MoE serving wins from batching because cross-token routing amortizes all-to-all cost.
- **$/MTok picture, mid-2026 practitioner ballpark:** Qwen3-MoE 235B-A22B FP4 on 8× B200 ≈ $0.30/MTok at ~1200 tok/s/GPU; DeepSeek V3.1 FP4 + MTP on NVL72 ≈ $0.25/MTok; vs Llama 3.3 70B FP8 on 4× H100 ≈ $0.40/MTok. Reading: MoE on Blackwell is the cost winner once cluster scale exists; for single-replica deployment, dense Hopper still wins.
- **Artificial Analysis table (Aug 2026) shows the same shape at matched ~30B total:** Nemotron 3.5 Lightning (30B MoE + Mamba-2, 3B active) at $0.22/M output and 235–494 t/s vs Gemma 4 31B dense at $0.40/M and 37–222 t/s — 4–5× the speed at ~1/14th the price, but less than half the capability score (24 vs 30). The 2026 consensus: MoE buys speed and $/token at some capability cost per total-param; at matched ACTIVE params, MoE's larger total capacity gives it a higher capability ceiling — which is why every frontier open-weight release is now MoE.

### Serving MoE at scale: expert parallelism, load balancing, quantization

