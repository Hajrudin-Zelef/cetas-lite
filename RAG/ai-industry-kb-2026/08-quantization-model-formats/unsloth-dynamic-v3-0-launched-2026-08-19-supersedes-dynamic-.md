---
id: ai-industry-kb-2026/08-quantization-model-formats/unsloth-dynamic-v3-0-launched-2026-08-19-supersedes-dynamic-
title: "Unsloth Dynamic v3.0 (launched 2026-08-19 — supersedes Dynamic 2.0)"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Meta", "Nvidia", "Unsloth", "Z.ai"]
dates: ["2026-08-15", "2026-08-19", "2026-08-20", "2026-08-27", "2026-09-22"]
keywords: ["agentic", "benchmark", "benchmarks", "blackwell", "deepseek", "distillation", "distribution", "embeddings", "gguf", "glm", "gpu", "inference"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3809, 3862]
section: "8. Quantization & Model Formats"
sha256: 591241adadf2296968b8624c62ce439ed7bd39b88540c961e3ea09d55cb79232
---

# Unsloth Dynamic v3.0 (launched 2026-08-19 — supersedes Dynamic 2.0)

| Quant | bpw | KL median | KL q99 | Top tokens differ | ln(PPL ratio) |
|---|---|---:|---:|---:|---:|
| IQ1_S | 1.78 | 0.5495 | 5.5174 | 0.3840 | 0.9235 |
| IQ2_XXS | 2.20 | 0.1751 | 2.4983 | 0.2313 | 0.2988 |
| IQ2_XS | 2.43 | 0.1146 | 1.7693 | 0.1943 | 0.2046 |
| IQ2_S | 2.55 | 0.0949 | 1.6284 | 0.1806 | 0.1722 |
| IQ2_M | 2.76 | 0.0702 | 1.0935 | 0.1557 | 0.1223 |
| Q2_K | 3.00 | 0.0588 | 1.0337 | 0.1492 | 0.1103 |
| IQ3_XXS | 3.21 | 0.0330 | 0.5492 | 0.1137 | 0.0589 |
| IQ3_XS | 3.32 | 0.0296 | 0.4550 | 0.1071 | 0.0458 |
| Q3_K_S | 3.50 | 0.0304 | 0.4481 | 0.1068 | 0.0511 |
| IQ3_S | 3.52 | 0.0205 | 0.3018 | 0.0895 | 0.0306 |
| IQ3_M | 3.63 | 0.0186 | 0.2740 | 0.0859 | 0.0268 |
| Q3_K_M | 3.89 | 0.0171 | 0.2546 | 0.0839 | 0.0258 |
| IQ4_XS | 4.32 | 0.0088 | 0.1082 | 0.0606 | 0.0079 |
| IQ4_NL | 4.56 | 0.0085 | 0.1077 | 0.0605 | 0.0074 |
| Q4_K_S | 4.57 | 0.0083 | 0.1012 | 0.0600 | 0.0081 |
| Q4_K_M | 4.83 | 0.0075 | 0.0885 | 0.0576 | 0.0060 |
| Q5_K_M | 5.67 | 0.0043 | 0.0368 | 0.0444 | 0.0005 |
| Q6_K | 6.57 | 0.0032 | 0.0222 | 0.0394 | −0.0008 |

- Readings: **IQ-quants dominate their size class** — IQ3_S (3.52 bpw) beats Q3_K_M (3.89 bpw) on every column while being smaller; IQ4_XS (4.32 bpw) is within noise of Q4_K_M (4.83 bpw) at ~0.5 fewer bits/weight. The cliff is at 2-bit: IQ2_XS shows ln(PPL ratio) 0.2046 (roughly +23% perplexity) and 19% of top tokens differing; IQ1_S at 0.9235 (≈+152% PPL) is "for the desperate" only.
- The community ladder (mradermacher cards — DeepSeek-V2.5, QwQ-32B, Qwen2.5-14B, Gemma-2-9B — converges on the same ordering): i1-IQ1_S/IQ1_M "for the desperate"; IQ2_XXS…IQ2_M usable only on large models where redundancy absorbs the damage; **i1-IQ3_S "beats Q3_K\*" at equal file size**; i1-IQ4_XS the IQ entry point to the ~4-bit tier; **i1-Q4_K_S "optimal size/speed/quality"**; i1-Q4_K_M "fast, recommended"; i1-Q6_K "practically like static Q6_K". Restated rule: **"IQ-quants are often preferable over similar sized non-IQ quants."** All IQ types require an imatrix for best results (the `i1-` prefix marks imatrix-weighted quants).
- Scale datapoint: DeepSeek-V2.5 (236B-class) at i1-IQ4_XS = 125.7 GB with ~0.8% perplexity increase; Qwen2.5-14B at i1-Q4_K_M = 9.1 GB.
- GGUF-side quality cascade at ~4-bit (community consensus): **UD-Q4_K_XL / UD-Q4_K_M (Unsloth Dynamic) > Q4_K_M > IQ4_XS**.

### Unsloth Dynamic v3.0 (launched 2026-08-19 — supersedes Dynamic 2.0)

- Wave 1 documented Dynamic 2.0 as current; as of **2026-08-19** the current generation is **Dynamic v3.0**, announced alongside **Qwen3.8-27B** (launched 2026-08-15).
- Mechanism: per-layer importance-weighted quantization assigning GGUF quant *types* per tensor (each type has a different block structure, scale format and codebook), instead of one uniform bit-width — conceptually: embeddings and LM head at Q8_0, sensitive early layers at Q6_K, robust middle layers at Q4_K, late layers at Q5_K.
- [VENDOR] Headline claims: **>10% better top-1 accuracy at the same size than every other quantization provider**, measured on Qwen3.8-27B; Qwen3.8-27B v3.0 quants recorded **5.1 million downloads in 5 days**; stronger results on **Divergence-300 @32 and KL Divergence** vs Dynamic 2.0; **1-bit quants retaining 77% accuracy**, runnable on 8 GB RAM.
- Mechanism upgrades over v2.0: (1) a **much higher-quality imatrix calibration dataset from diverse sources** (>1.5M tokens), explicitly refined for **agentic coding, chat, and multilingual performance** — "instruct models have unique chat templates, and using text-only calibration datasets is not effective for instruct models"; (2) **improved layer-selection strategies** — per-layer dynamic bit allocation across the full GGUF type set (Q2_K, Q3_K_M, Q4_K_M, Q5_K, Q6_K, Q8_0, IQ1_S, IQ2_XXS, etc.); (3) combinations of multiple quantization techniques.
- Unsloth stresses everything is **pure PTQ — no QAT, no quantization-aware distillation (QAD)** — "resulting in a lower risk of overfitting," with the imatrix published for community testing.
- Independent technical analysis (malaiwah, 2026-08-20): per-layer allocation over the GGUF type set; corroborating finding that calibration data must match the model's inference distribution — text-only calibration fails for instruct models, matching the GDN `cache=None` tracing results. The headline `UD-Q2_K_XL` of Qwen3.8-27B: **~9.83 GB**, reported to generate working HTML programs (a task where conventional quants break down) with a single minor JS bug.
- Independent reverse-engineering (orka.py project, measured on LFM2.5-2.6B, wikitext-2): the per-layer schedule is **not in Unsloth's open-source tree** — the OSS repo delegates to llama.cpp's `llama-quantize`; the schedule lives in the proprietary `unsloth_zo` package. Dynamic quants are consumed as prebuilt artifacts, not reproduced from open code.
- Dynamic quants are named with a `UD-` prefix (e.g. `UD-Q4_K_M`, `UD-Q4_K_XL`, `UD-IQ2_M`) and are distributed across Unsloth's **1,374-model** Hugging Face catalog (GGUF + NVFP4).
- The UD variant ladder:

| Variant | Bits | Use case |
|---|---|---|
| `UD-TQ1_0` | ~1-bit | Extreme compression; DeepSeek-V3.1 671B → ~192 GB |
| `UD-IQ2_M` | ~2-bit | Accuracy-maximizing 2-bit (slower conversion) |
| `UD-Q2_K_XL` | ~2.7-bit | **Recommended** size/accuracy balance (e.g. GLM-5.1 → ~220 GB) |
| `UD-Q3_K_XL` | ~3-bit | Middle ground (Qwen3.8-27B ~92.4% top-1 BF16 agreement) |
| `UD-Q4_K_XL` | ~4-bit | Highest quality (Qwen3.8-27B ~96% top-1 BF16 agreement, 16.69 GiB) |

- (Top-1 agreement figures are third-party graph readings of Unsloth's Qwen3.8 fidelity curve — quant-fidelity proxies, not downstream benchmark scores. [DIRECTIONAL])
- Operational notes: Unsloth deleted the non-dynamic quants from the Qwen3.8-27B-GGUF repo — Dynamic v3.0 is now the only flavor published there; MTP (multi-token prediction) models are published as **separate `mtp.gguf` files** (unresolved as of 2026-09-22); model coverage for v3.0 includes Qwen3.8, Meta Muse, Glimmer, and DeepSeek-V4-Pro.
- Unsloth also ships **"Dynamic Unsloth NVFP4 Quants"** ("faster and more accurate") for the NVIDIA stack — e.g. `unsloth/gemma-4-31B-it-NVFP4`, `gemma-4-26B-A4B-it-NVFP4`, `gemma-4-12b-it-NVFP4` — and hybrid **NVFP4-GGUF** containers (e.g. Qwen3.6 NVFP4-GGUF): a GGUF container holding NVFP4-packed weights, accelerated only via Unsloth kernels on Blackwell.
- Community benchmark compilation (Spheron, Sept 2026; figures described as **directional estimates from community benchmarks, not direct measurements** — Llama 3.3 70B): Dynamic 2.0 (Q4) +0.12 PPL delta vs BF16 vs +0.18 for uniform Q4_K_M — **Dynamic 2.0 closed about one third of the gap between uniform Q4_K_M and Q5_K_M at the same file size and memory footprint as Q4_K_M**. Unsloth's own "SOTA on MMLU/KL divergence" claims for v2.0 had no independent benchmarks as of Sept 2026 [PARTIALLY VERIFIED — vendor claim].
- Extreme-compression showcase: Z.ai's **GLM-5.2** (744B total parameters, 40B active, 1M context) ships Unsloth Dynamic GGUFs — Dynamic 1-bit reaches ~76.2% top-1 accuracy while 86% smaller; Dynamic 2-bit reaches ~82% accuracy while 84% smaller; **UD-IQ2_M uses 239 GB of disk** and fits on a **256 GB unified-memory Mac**, or runs with a 24 GB GPU plus 256 GB RAM via MoE offloading.

### First independent Dynamic v3.0 replication test (2026-08-27)

