---
id: ai-industry-kb-2026/08-quantization-model-formats/fp4-nvfp4-and-mxfp4-the-2026-frontier
title: "FP4 — NVFP4 and MXFP4, the 2026 frontier"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Intel", "Moonshot", "Nvidia", "SGLang", "vLLM"]
dates: ["2025-05", "2026-09"]
keywords: ["fp4", "mxfp4", "nvfp4", "awq", "benchmarks", "blackwell", "compute", "cost", "decode", "deepseek", "gptq", "gpu"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4234, 4243]
section: "8. Quantization & Model Formats"
sha256: 175df3342baa2b471f95b20a8c2f79711b0cea938de0861d4119f0b548cadb3a
---

# FP4 — NVFP4 and MXFP4, the 2026 frontier

- **FP16→INT4 (GPTQ/AWQ) = 4× memory reduction, 1–3% quality loss** — the format that puts a 70B dense model on a single 48 GB GPU.
- AWQ (Lin et al., arXiv 2306.00978): scales weights by activation importance before quantization; W4A16, group size 128; consensus best INT4 PTQ quality; official kernel is the default vLLM path, Marlin/Machete when large-batch throughput matters.
- GPTQ: column-by-column inverse-Hessian error compensation; requires calibration (~5–30 minutes for a 7B model); slightly worse quality than AWQ on most benchmarks but widest hardware compatibility.
- Marlin family (fused dequantize+GEMM for weight-only 4-bit): **Marlin's floor is SM75 (Turing), not SM80** — audited in vLLM v0.26.0 source (`if device_capability < 75`); **Marlin-24 act-order kernels were removed from vLLM main**; Machete is SM90-exact (Hopper); Conch targets SM80+.
- QAT vs PTQ at INT4: **PTQ quality drop 3–8% vs QAT <1%**; cost is full retraining compute plus format lock-in. **Kimi K2.6 ships INT4 natively via QAT (~2× speedup)**; **DeepSeek-V4 applies FP4 (MXFP4) QAT to MoE experts plus the indexer QK path** — the parameter-heavy parts.
- 2026 tooling shakeout (detail in part 08a): **AutoAWQ is officially deprecated** — the AWQ workflow moved to llm-compressor; AutoGPTQ's successor is **GPTQModel** (5,000+ GPTQ repos on Hugging Face reference it); **AutoRound (Intel)** is the strongest measured 4-bit PTQ (beats GPTQ 30/32 configs and AWQ 27/32 on lm-eval; September 2026 independent logprob test: 91.4% top-1 agreement vs BF16, KL 0.277, vs GPTQ 90.5% / 0.319); **HQQ** owns the data-free/fast niche (1/2/3/4/8 bits, no calibration).
- AWQ deployment nuance (September 2026): vLLM and SGLang both ship awq/awq_marlin (and gptq_marlin) paths, and AWQ is described as "the more commonly used of the two" vs GPTQ on vLLM deployments — but workflow friction on 2026 models is real (AutoAWQ was archived in May 2025 and never got Qwen3.5 support; llm-compressor's pip release pins `transformers<=4.57.6`, predating Qwen3.5 support; AWQ's layer-by-layer scale search OOMs on a 40 GB A100 — fix: `offload_device=cpu`); and a community 4× RTX PRO 6000 Blackwell shootout had **AWQ beating both NVFP4 recipes on decode throughput at every concurrency** (C=128: AWQ 3519 vs 3232/3220 tok/s with MTP) [COMMUNITY, single rig — NVFP4 backend/kernel version-sensitive]. Net: dominant *compat* INT4 path, not always the fastest.

### FP4 — NVFP4 and MXFP4, the 2026 frontier

