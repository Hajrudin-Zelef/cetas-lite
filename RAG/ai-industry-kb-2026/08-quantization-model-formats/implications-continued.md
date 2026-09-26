---
id: ai-industry-kb-2026/08-quantization-model-formats/implications-continued
title: "Implications (continued)"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Hugging Face", "Intel", "Meta", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "Qualcomm", "Unsloth", "Z.ai", "vLLM"]
dates: ["2025-05", "2026-02-13", "2026-02-20", "2026-03-04", "2026-05-11", "2026-05-13", "2026-05-28", "2026-07-15", "2026-07-26", "2026-08-15", "2026-08-18", "2026-08-19", "2026-08-27", "2026-09", "2026-09-18", "2026-09-22"]
keywords: ["agentic", "amd", "apache", "awq", "benchmark", "benchmarks", "bitnet", "blackwell", "compute", "consumer", "cost", "datacenter"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4462, 4510]
section: "8. Quantization & Model Formats"
sha256: 64d98daa128ddf01d8e2c385520c25bb2fd2b28cfa2ed7f9f711b9d4645c0de3
---

# Implications (continued)

| Date | Event |
|---|---|
| 2023-05 | QLoRA / NF4 (Dettmers et al.) — 4-bit index into a normal-distribution codebook; fine-tuning standard |
| 2023-09 | **OCP MX specification v1.0** (AMD, Arm, Intel, Meta, Microsoft, NVIDIA, Qualcomm) |
| 2024 | BitNet b1.58 paper (Ma et al., Microsoft Research + UCAS): ternary training from scratch |
| 2025-04 | **BitNet b1.58 2B4T** released (2B params, 4T tokens, MIT) |
| 2025-05 | AutoAWQ archived (May 2025) — never got Qwen3.5 support; the AWQ workflow migrates to llm-compressor |
| 2025-08 | **OpenAI GPT-OSS** ships natively in **MXFP4** (120B on single H100) |
| 2026-01 | BitNet CPU-optimization update (+1.15–2.1×); arXiv 2601.09527: NVFP4 consumer-Blackwell economics (2× RTX 5090, Qwen3-8B at $0.002/MTok) |
| 2026-02-13 | vLLM project blog: DeepSeek-V3.2/R1 NVFP4 on GB300 (R1: 22,476 TGS prefill-only) |
| 2026-02-20 | **ggml.ai joins Hugging Face** announced (Discussion ggml-org/llama.cpp #19759; MIT preserved, team autonomy) |
| Feb 2026 | DeltaKV (arXiv:2602.08005): 3.4× KV compression, near-lossless; ARKV (arXiv:2603.08727): 4× KV at ~97% accuracy |
| 2026-03-04 | KVQuant pre-RoPE quantization paper finalized (arXiv:2401.18079) |
| 2026-04 | SAW-INT4 (arXiv:2604.19157): system-aware token-wise INT4 KV, no calibration |
| 2026-05-11 | **Red Hat AI's vLLM TurboQuant evaluation published**; PR #38479 opens the vLLM integration path; FP8 confirmed best KV default |
| 2026-05-13 | NVIDIA posts **Kimi-K2.6-NVFP4** on Hugging Face (1T params, 32B active; ModelOpt 0.44.0; vLLM on B200) |
| 2026-05-28 | canada-quant measurements: NVFP4→Marlin fallback on consumer Blackwell forfeits FP4 FLOPS |
| Jul 2026 | **Unsloth starts NVFP4 export** ("Dynamic Unsloth NVFP4 Quants" + NVFP4-GGUF hybrids) |
| 2026-07-15 | **PrismML Bonsai-27B** announced (Apache 2.0): 27B at 3.9 GB (1-bit), 11 tok/s on iPhone 17 Pro |
| 2026-07-26 | Community FP4 KV route (FA2 prefill + XQA decode) validated on RTX 5090 at 1.6× fp8 KV pool |
| 2026-08-18 | massed-compute publishes Nemotron-3.5-Lightning-30B-A3B-NVFP4 benchmarks (L40S/A100/PRO 6000) |
| 2026-08-27 | srmiles publishes first independent Dynamic v3.0-vs-v2.0 test (KLD vs Q8_0 referee; 9.5% slower decode; UD-Q3_K_XL llama-server bug) |
| Sep 2026 | rtx6kpro NVFP4-vs-AWQ decode shootout on 4× RTX PRO 6000: AWQ wins at every concurrency; OpenVINO 2026.2 ships INT4 KV on Intel GPU |
| 2026-09-18 | MarkTechPost overview: **AutoAWQ officially deprecated**; llm-compressor is the AWQ path |
| 2026-06 | blokz survey documents DeltaKV/ARKV/INT4 KV field state |
| 2026-08-15 | **Qwen3.8-27B launch** (8.3M downloads first week) — Dynamic v3.0 launch vehicle |
| 2026-08-19 | **Unsloth Dynamic v3.0** official release (>10% top-1 accuracy claim; 5.1M quant downloads in 5 days) |
| Sep 2026 | Independent AutoRound-vs-GPTQ logprob-parity test (91.4% vs 90.5% top-1 agreement); FP4-vs-AWQ decode shootouts on Blackwell rigs |
| 2026-09-22 | Hugging Face confirms optimum-quanto is in **maintenance mode** (redirects to bitsandbytes/torchAO) |

Context framing: by September 2026 the precision stack has split cleanly by hardware — **FP8 the default on Hopper/Ada/Blackwell, INT4 (AWQ/GPTQ/Marlin) the pre-Hopper and memory-constrained standard, FP4 the Blackwell-native frontier** with the two-level-scaling anatomy (NVFP4) or the open spec (MXFP4). FP4 is the production norm on Blackwell and experimental/limited elsewhere — weights via the Marlin fallback, KV cache only on datacenter Blackwell in stock stacks.

## Implications (continued)

1. **Format choice is now a hardware decision, not a quality debate.** FP8 is the default where Hopper/Ada/Blackwell tensor cores exist; INT4 (AWQ best quality, GPTQ widest compat, Marlin kernels) owns pre-Hopper and single-GPU-constrained serving; NVFP4/MXFP4 are the Blackwell-native future.
2. **The 4-bit quality war is over at the top end**: AWQ holds best-PTQ quality historically, AutoRound is the strongest measured 4-bit PTQ method in 2026, QAT (Kimi K2.6, DeepSeek-V4) pushes INT4/FP4 under 1% loss, and microscaling FP4 formats structurally beat INT4 at equal bit-width.
3. **FP4's bottleneck is no longer math but deployment surface**: the formulas are verified (−0.1 MMLU at 671B scale; ~4.5 bits/element), while kernel coverage (SM100/103 native, SM120 software dequant, Marlin fallback elsewhere), KV-cache gating, and checkout-time W4A16-vs-W4A4 confusion are the practical limits.
4. **Local-first AI crossed a threshold in 2026**: 27B-class models on phones (Bonsai) and 70B on laptops (Q4_K_M) mean serious agentic/coding workloads run privately, offline, with zero per-token cost — quantization moved the break-even of self-hosting far below cloud pricing.
5. **Watch the KV cache, not just the weights**: at long context, KV dominates memory — 4-bit KV (Bonsai, OpenVINO 2026.2, SAW-INT4/DeltaKV/ARKV), FP8 KV (production default), and emerging FP4 KV paths decide what context lengths are actually deployable; the 8B KV-vs-weights crossover moved from 128K to ~512K at INT4.
6. **MoE + quantization is the combination that wins at scale**: compressing inactive experts is nearly free in quality-per-FLOP (GLM-5.2 744B at 239 GB disk, DeepSeek-V4-Flash 284B at 172 GB), but the KV cache does not shrink with sparsity — a 400B+ MoE's KV budget is the remaining constraint.
7. **Standardization tension**: FP8 and MXFP4 are open or cross-vendor (OCP spec, AMD/Intel coverage); NVFP4 is NVIDIA-proprietary and Blackwell-gated. At scale, quantization labor costs (calibration, QAT retraining, per-hardware kernel validation) push buyers toward the standardized formats — FP8 now, MXFP4 as the open 4-bit candidate.
8. **Measure what matters**: perplexity deltas are the first check, KL divergence the stricter distribution test, per-task benchmark retention the reliable signal — aggregate scores can lie; QAT vs PTQ is the largest single quality lever at 4-bit.

9. **Per-format precision is not the whole story**: per-layer allocation (Unsloth Dynamic schedules, KVQuant's non-uniform per-layer bits) matters as much as the nominal bit-width — the same nominal "4-bit" can sit ±0.1 perplexity apart on schedule alone; Dynamic v3.0's 9.5% decode slowdown is the cost of that heterogeneity.
10. **FP8 became the unit of currency**: benchmark deltas are quoted against FP8 (DeepSeek-R1-FP4's −0.1 MMLU is vs the FP8 base), making FP8 the 2026 reference point the way FP16 was in 2024 — every finer format must justify itself against it.
11. **The NVFP4/MXFP4 split is a bet on openness**: MXFP4's OCP standard plus AMD/Intel coverage versus NVFP4's tcgen05 moat; whoever wins the cross-vendor 4-bit default decides whether Blackwell-era checkpoints are portable or NVIDIA-locked. MXFP4's llm-compressor path (no calibration needed) is its sharpest adoption edge.

Watchlist carried into the final document:

