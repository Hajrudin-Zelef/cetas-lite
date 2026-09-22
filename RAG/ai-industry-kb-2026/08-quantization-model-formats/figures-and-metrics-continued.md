---
id: ai-industry-kb-2026/08-quantization-model-formats/figures-and-metrics-continued
title: "Figures and metrics (continued)"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AWS", "Alibaba", "Apple", "DeepSeek", "Intel", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "Perplexity", "Z.ai", "vLLM"]
dates: ["2026-01", "2026-08-18", "2026-09"]
keywords: ["attention", "awq", "aws", "benchmark", "benchmarks", "blackwell", "compute", "consumer", "cost", "decode", "deepseek", "disaggregated"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4346, 4431]
section: "8. Quantization & Model Formats"
sha256: 0026eeea7bc836cd8e06443c53fee4f715e1d721cafa09ba09baf7d04c21376c
---

# Figures and metrics (continued)

## Figures and metrics (continued)

### Density and accuracy arithmetic

- True density is **bits per weight (bpw)** — a "4-bit" quant is often ~4.5–5.2 bpw with scales included. Notation: **W4A16 / W8A8 / W4A4** = weight-bits + activation-bits.
- FP16→FP8: 2×; FP16→INT4: 4× (practical ~4.5–5.2 bpw); FP16→INT2: 8×; FP32→MXFP4: 7.5× (136 bits vs 1,024 bits per 32-element block).
- FP4 spacing advantage: floating-point levels are logarithmic; INT4's uniform grid must either clip outliers or expand the scale and quantize most values to zero — block-level shared scales give FP4 consistently better signal-to-noise at equal bit-width.
- MXFP4 block math: 32 elements × 4 bits + 8-bit E8M0 scale = 136 bits per block. NVFP4 element math: ≈ 4 + 8/16 bits + negligible per-tensor scale ≈ **4.5 bits/element**.
- NVIDIA's official NVFP4 guidance: **~3.5× memory vs FP16, ~1.8× vs FP8, usually <1% accuracy loss** (build.nvidia.com).
- GB10 measured: **160 GB → 40 GB** (4×) for NVFP4 W4A4 MoE inference; DeepSeek-V3.1 671B → ~192 GB via 1-bit UD-TQ1_0.

### Measured per-type accuracy (GGUF per-type table, Artefact2 — bpw, median KL vs FP16, ln(PPL ratio))

| Quant | bpw | KL median | ln(PPL ratio) |
|---|---|---:|---:|
| IQ1_S | 1.78 | 0.5495 | 0.9235 (~+152% PPL — "for the desperate") |
| IQ2_XS | 2.43 | 0.1146 | 0.2046 (~+23% PPL) |
| Q2_K | 3.00 | 0.0588 | 0.1103 |
| Q3_K_S | 3.50 | 0.0304 | 0.0511 |
| IQ3_S | 3.52 | 0.0205 | 0.0306 — beats Q3_K_S at equal bpw on every column |
| Q3_K_M | 3.89 | 0.0171 | 0.0258 |
| IQ4_XS | 4.32 | 0.0088 | 0.0079 — within noise of Q4_K_M at ~0.5 fewer bits/weight |
| Q4_K_M | 4.83 | 0.0075 | 0.0060 |
| Q5_K_M | 5.67 | 0.0043 | 0.0005 |
| Q6_K | 6.57 | 0.0032 | −0.0008 |
| Q8_0 | 8.57 | ~0.00 | +0.03% perplexity vs FP16 |

- Readings: the cliff is at 2-bit (IQ2_XS: 19.4% of top tokens differing); **IQ-quants dominate their size class**; Q4_K_M remains the default quality/size trade-off.
- Perplexity-delta rule of thumb: <0.1 ≈ indistinguishable, 0.1–0.2 ≈ production-safe, >0.5 ≈ noticeable.

### Quality retention benchmarks

- DeepSeek-R1-FP4 (671B): MMLU 90.7 vs 90.8 (−0.1), GSM8K 96.1 vs 96.3 (−0.2) vs the FP8 base.
- Bonsai-27B: 1-bit **>90%**, ternary **>95%** of full-precision benchmark performance.
- GLM-5.2 (744B total) Dynamic quants: 1-bit ~76.2% top-1 accuracy at 86% smaller; 2-bit ~82% at 84% smaller.
- Nemotron-3-Nano-30B NVFP4 via QAD: **99.4%** of the BF16 baseline [VENDOR].
- AutoRound (Intel): **91.4% top-1 logprob agreement** vs BF16 (KL 0.277) vs GPTQ 90.5% / 0.319 — September 2026 independent test.
- QAT vs PTQ at INT4: **3–8% drop (PTQ) vs <1% (QAT)** — the largest single quality lever at 4-bit.
- French 2026 industry guide: well-calibrated 4-bit degrades **2–4 MMLU points for 7B+ models** — acceptable for most production.
- INT4 KV cache: **4× reduction with 1–3% degradation**; ARKV 4× at ~97% baseline accuracy; DeltaKV ~3.4× near-lossless.
- Caveat carried: aggregate benchmark scores can lie — a format can hold MMLU while degrading long-form coherence; per-task evaluation (especially coding) is the reliable signal. KL divergence catches distribution shifts perplexity averages hide.

### What model sizes now run where (VRAM impact)

| Model scale | Format | Footprint | Where it runs (verified) |
|---|---|---|---|
| 70B dense | GGUF Q4_K_M | ~40 GB | 64 GB laptop (M3 Max: 10–15 tok/s); single 48 GB GPU via INT4 |
| 70B dense | FP8 | ~70 GB | Single H100 (30–50% faster than FP16) |
| 70B+ dense | NVFP4/MXFP4 | ~35–40 GB | Single RTX PRO 6000 96 GB (Blackwell) |
| 27B dense | Bonsai 1-bit | **3.9 GB weights / 4.2 GB peak** | iPhone 17 Pro (11 tok/s); 16 GB M4 Mac mini (21.5 tok/s) |
| 27B dense | Ternary Q2_0_g128 | ~7.2 GB | Consumer laptop; full 262K ctx at ~9.4 GB peak |
| 284B MoE (DeepSeek-V4-Flash) | NVFP4 | 172 GB | B200/B300 native; 2-GPU or high-RAM hosts |
| 744B MoE (GLM-5.2, 40B active) | UD-IQ2_M | **239 GB disk** | 256 GB unified-memory Mac; or 24 GB GPU + 256 GB RAM offload |

### Cost arithmetic — what quantization saves

- **Halving the GPU count halves the hourly bill**: FP16 Llama 3.3 70B needs 2× A100 (140 GB); Q4_K_M fits on 1× A100 (~42 GB) at similar batch-1 throughput — roughly **2× cheaper per hour** for the same workload.
- **OpenAI's MXFP4 deployment** of GPT-OSS is credited with **~75% inference-cost reduction** versus the hardware FP16 would have required (120B on a single H100 instead of a multi-GPU node).
- **FP8 on H100**: 30–50% throughput gain at equal quality → ~1.3–1.5× more tokens per GPU-hour, i.e. a direct **25–35% cost-per-token cut**.
- **AWS Blackwell rental economics** (NextPlatform sizing): FP8 halves cost-per-teraflop vs FP16; FP4 halves it again — a model moved to FP4 needs **one quarter of the machinery at one quarter of the cost**, or the same money trains a 4× larger model [DIRECTIONAL].
- Consumer-Blackwell economics (arXiv 2601.09527, January 2026): **2× RTX 5090, Qwen3-8B NVFP4 at concurrency 128–256 → 4,400+ TPS at $0.002/MTok**, sustaining 380M tokens/day at under $800/year electricity; RAG at 8K context on 1× RTX 5090: 411 TPS at $0.029/MTok vs $1,900–15,000/month for equivalent cloud APIs [COMMUNITY — electricity-price and utilization assumptions apply].
- Private-Blackwell RAG study: self-hosted $0.001–$0.04/MTok (electricity only); hardware break-even under 4 months at 30M tokens/day [COMMUNITY/SECONDARY].
- Nemotron-3.5-Lightning-30B-A3B-NVFP4 (massed-compute, 2026-08-18): L40S 1345.4 tok/s @ $0.182/1M out tokens; RTX PRO 6000 2541.1 tok/s @ $0.239/1M; A100 2100.8 tok/s @ $0.182/1M — best $/token tied on L40S/A100, native Blackwell path wins raw throughput [COMMUNITY].
- **Local-first economics**: a 27B model at 4.2 GB (Bonsai) or a 70B at 40 GB (Q4_K_M) runs on hardware the user already owns — marginal inference cost drops to electricity; quantization moved the break-even point of self-hosting far below cloud per-token pricing for steady workloads.
- The counterweight: **quantization labor is not free** — GPTQ calibration runs, QAT retraining, and per-hardware kernel validation (e.g. the sm_120 FP4-KV saga) are engineering costs that favor standardized formats (FP8, MXFP4) over bespoke ones at scale.

### MoE × quantization — why the combination wins

- Mixture-of-Experts models are quantization's biggest beneficiaries: only a fraction of parameters are active per token (GLM-5.2: 40B active of 744B; Kimi K2.6: 32B active of 1T; MiniMax M3: ~23B of 428B), so compressing the **inactive** expert weights costs almost nothing in quality-per-FLOP while dividing memory by 4–8×.
- Production pattern (2026): **MXFP4/NVFP4 QAT applied to MoE experts** (DeepSeek-V4: MXFP4 QAT on experts + indexer QK path), FP8 for attention blocks, and router/gate layers kept unquantized for stability (`--moe` flag in llm-compressor pipelines).
- Verified footprints: **GLM-5.2** (744B total) at **UD-IQ2_M = 239 GB** disk → 256 GB unified-memory Mac; **DeepSeek-V4-Flash** (284B) at **NVFP4 = 172 GB** → B200/B300 native.
- Caveat: the **KV cache does not shrink with MoE sparsity** — it stays proportional to layer count and context length. A quantized 400B+ MoE still needs its full KV budget at long context, which is why FP8/4-bit KV and disaggregated prefill/decode matter as much as weight formats.
### Throughput tables (measured, September 2026)

- Qwen-Coder 32B FP8: RTX 6000 Ada ~100–120 tok/s (1.7–2× vs FP16); RTX 4090 ~60–80 tok/s (emulated, no speedup).
- Llama 3.3 70B, A100 80GB, batch=1 (estimated): Q4_K_M ~20–30 tok/s at ~42 GB VRAM; Q5_K_M ~28–36 tok/s at ~48 GB; Q8_0 ~20–25 tok/s at ~75 GB; FP16 needs 2× A100 (~30–40 tok/s, 140 GB) — Q4_K_M on one A100 costs roughly half per hour of FP16 on two.
- DeepSeek-V3.2 NVFP4+TP2 (GB300): **7,360 TGS prefill-only**, 2,816 TGS mixed (ISL=2k/OSL=1k); DeepSeek-R1 NVFP4+EP2 (2× GB300): **22,476 TGS prefill-only**, 3,072 TGS mixed — 8× prefill, 10–20× mixed-context vs Hopper.
- Nemotron-3.5-Lightning-30B-A3B-NVFP4 on vLLM 0.27.1 (2026-08-18): L40S 1345.4 tok/s @ $0.182/1M out tokens; **RTX PRO 6000 Blackwell 2541.1 tok/s** @ $0.239/1M; A100 2100.8 tok/s @ $0.182/1M [COMMUNITY].
- RTX PRO 6000 shootout (Qwen3.8-Flash-Next, vLLM 0.17.0rc1, TP4, Sept 2026): AWQ 3519 vs NVFP4 3232/3220 tok/s at C=128 with MTP; 2796 vs 2294/2291 without MTP — AWQ won decode at every concurrency on this rig [COMMUNITY].
- Qwen3.6-35B-A3B NVFP4 GEMM on consumer Blackwell (s0me1-dev SM120 patches): 175 tok/s [COMMUNITY].
- 2× RTX 5090, Qwen3-8B NVFP4, concurrency 128–256: **4,400+ TPS at $0.002/MTok**; 380M tokens/day under $800/year electricity [COMMUNITY].
- Private-Blackwell RAG (vLLM + AIPerf): NVFP4 **1.6× throughput vs BF16**, 41% energy reduction, 2–4% quality loss [COMMUNITY/SECONDARY].
- INT4 KV + fused flash-attention kernel: **3.9× kernel speedup** to the bandwidth roofline; attention MAE 3e-08 vs FP32; 42/42 GPU tests passing [COMMUNITY].
- NVIDIA official figures: NVFP4 ≈ 2× FP8 GEMM compute; B200 FP4 ≈ 8× FP16 tensor-core throughput [VENDOR]; FP8 30–50% speedup over FP16 on H100.
- Hardware sizing rules: **RTX PRO 6000 (Blackwell, 96 GB)** holds 70B+ models in FP4 on a single card; **RTX 5090 (32 GB)** is the budget entry point for 14B–32B FP4 inference; Ornith-1.5-35B-A3B-NVFP4 loads on an 8 GB-class Ada GPU via the Marlin fallback (~4.5 bits/value, W4A16-class speed) vs 21.7 GB for the GGUF Q4_K_M of the same model.

