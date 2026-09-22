---
id: ai-industry-kb-2026/08-quantization-model-formats/first-independent-dynamic-v3-0-replication-test-2026-08-27
title: "First independent Dynamic v3.0 replication test (2026-08-27)"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "Alibaba", "Apple", "Hugging Face", "Intel", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM"]
dates: ["2025-05", "2026-07-14", "2026-08-18", "2026-08-27", "2026-08-31", "2026-09", "2026-09-08", "2026-09-09"]
keywords: ["amd", "apache", "awq", "benchmark", "benchmarks", "blackwell", "consumer", "cost", "decode", "distillation", "distribution", "fine-tuning"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3861, 3942]
section: "8. Quantization & Model Formats"
sha256: f650e314a5174cbb5b7ed89ba6e7ca94eb6c8a8cf71a3efdb5dfe942c4bbdf06
---

# First independent Dynamic v3.0 replication test (2026-08-27)

### First independent Dynamic v3.0 replication test (2026-08-27)

- Wave 2's open item ("no independent replication of the Divergence-300/KL metrics") is now partially addressed by a community benchmark (srmiles/local-llm-benchmarks, 2026-08-27): Dynamic v3.0 vs v2.0 on Qwen3.8-27B, matched quant names, real hardware, Q8_0 quant as referee.
- Verdict: **"yes on accuracy, no on throughput, and one file is unusable."** v3.0 is measurably closer to the Q8_0 referee on mean, median and tail KLD and on RMS Δp, *while being smaller* — corroborating the direction of Unsloth's ">10% better top-1" claim (KLD-based, not benchmark-score-based).
- But it is **not uniformly better**: it loses on 99.0% Δp at both sizes, and on top-1 agreement at 4-bit (consistent with Wave 2's note that v3.0's published top-1 agreement figures are graph readings only).
- v3.0 **decodes 9.5% slower** at Q4_K_XL (pooled medians over an interleaved A/B, Welch t = 4.05), alongside a much more heterogeneous block-type mix — the accuracy gain trades against decode speed.
- [COMMUNITY] **v3.0's `UD-Q3_K_XL` wedges llama-server**: a token id of `-1` reaches the batch, initialization fails, and with a drafter attached every subsequent request 500s until restart. Unsloth re-published the Qwen3.8-27B-GGUF files after launch ("a new version of Dynamic v3.0," per the HF discussion), so this file's provenance matters — may be fixed in later re-uploads.
- Net: direction corroborated on KL divergence with independent hardware; Divergence-300 specifically still unreplicated; decode-speed regression (~9.5% at Q4_K_XL) measured; one UD variant shipped with a serving-breaking defect.

### EXL2 → ExLlamaV3 / EXL3: end of an era

- **ExLlamaV2 is archived** — the README states: "This project is archived for now. Development continues on ExLlamaV3." The 2026 format landscape treats EXL2 as legacy: existing EXL2 checkpoints (e.g. LoneStriker's `6.0bpw` series, thousands of HF repos) still run under TabbyAPI, but no new development lands on the V2 codebase.
- EXL2's technical signature: same optimization method as GPTQ, 2/3/4/5/6/8-bit support, **any average bitrate between 2 and 8 bpw** via per-layer and per-column mixed precision — the converter quantizes each matrix several ways, measures error against calibration data, and picks the combination minimizing worst-case error at the target bitrate. Hence filenames like `4.65bpw` rather than `4-bit`. EXL2 renames tensors so every model looks like a Llama variant internally — the format never ported to other frameworks (TabbyAPI is the official server, OpenAI-compatible API).
- **ExLlamaV3** (turboderp-org/exllamav3) is the active successor — "optimized quantization and inference library for running LLMs locally on modern consumer-class GPUs." **v1.4.5 published 2026-08-31** with prebuilt Windows wheels (cp310–cp313, cu128/torch 2.10.0); September brought a fast bug-fix cadence **v1.4.6 → v1.4.9** (CUDA-graph capture fixes, devctx workspace hardening, COOP_AUTOTUNE_VERSION 3→4; no format changes). **v1.5.0 referenced** with fused routed-expert launch (no speed change measured: 52.05 vs 52.22 tok/s on GB10).
- **EXL3** is a streamlined variant of **QTIP** (Cornell RelaxML): one-step quantization (fused Viterbi kernel), minutes for small models on a single RTX 4090, **coherent at 1.6 bpw** [VENDOR] (Llama-3.1-70B in <16 GB). The GEMM kernel is Marlin-inspired and **memory-bound at 4 bpw on RTX 4090**; v1.0.0 (2026-07-14) greatly improved GEMM/GEMV on Ampere and dropped the `flash-attn-2`/`xformers` dependencies. V3 adds: tensor-parallel + expert-parallel for consumer setups, 2–8-bit KV cache quantization, speculative decoding, LoRA, multimodal support, v1.4.3 partial CPU expert offload, v1.4.4 vision-tower quantization + vision offload from system RAM, v1.4.5 improved MoE MTP performance.
- Verdict from the 2026 inference-engine shootouts: ExLlamaV3 is **the only credible raw-throughput challenger to llama.cpp on consumer NVIDIA GPUs** (memory-bound GEMM design) — sharpened by the September shootout below to **single-stream short-context decode only** — but EXL3 has no HF repo-store distribution: GGUF remains the portable choice and llama.cpp the baseline.
- The **vLLM-EXL3 serving plugin** (vcruz305/vllm-exl3, v0.4.2, 2026-09-09) documents real serving-integration fixes: dense EXL3 calls no longer launch exllamav3's cooperative trellis GEMM (rows 17–144 take the reconstruct+hgemm path) — on the vLLM nightly V2 model runner the cooperative GEMM **wedged the engine**; a 4-worker MTP stress that wedged within 3 minutes ran clean for 45 minutes after the routing fix, with decode speed unchanged. EXL3-in-vLLM is real but friction-heavy (third-party plugin).
- [COMMUNITY] Production datapoint: Qwen3.8-Flash-Next EXL3 4.05bpw on 4× RTX 3090 via TabbyAPI — 108 tok/s single-stream decode, 144.5 tok/s eight-stream aggregate, three concurrent 260k-token sessions in 96 GB VRAM; tool-calling quality 89/100 vs 88/100 baseline (inside ±2.5 noise).

### vLLM (NVFP4/FP8) vs ExLlamaV3 (EXL3) on RTX PRO 6000 (Sept 2026) [COMMUNITY]

- The most complete public format/engine shootout: MarcoPizeta's reproducible benchmark (04–08 Sept 2026), Qwen3.8-Flash-Next (125B-A6B + 51B n-gram/PLE table) on a single RTX PRO 6000 Blackwell 96 GB — vLLM with `primitive-ai/Qwen3.8-Flash-Next-mixed-NVFP4-FP8` + INT4 PLE offload + MTP vs ExLlamaV3/TabbyAPI with `turboderp/Qwen3.8-Flash-Next-exl3` (4.05bpw; 5.05bpw also tested). A v3 extension (08/09/2026) added SGLang with the calibrated `RadixArk/Qwen3.8-Flash-Next-NVFP4` checkpoint (`--quantization modelopt_fp4`, `--fp4-gemm-backend flashinfer_cutlass`, NEXTN speculative steps 3).
- Decode throughput (mean of 3 runs; same prompts per engine):

| Input × concurrency | vLLM MTP on | EXL3 4.05 MTP on |
|---|---|---:|
| 1k × 1 | 101.6 tok/s (TTFT 259 ms) | **168.9** tok/s (283 ms) |
| 1k × 4 | **296.4** tok/s (179 ms) | 221.4 tok/s (638 ms) |
| 8k × 1 | **140.5** tok/s (193 ms) | 129.1 tok/s (527 ms) |
| 8k × 4 | **282.1** tok/s (368 ms) | 125.2 tok/s (1466 ms), ok 56/60 |
| 32k × 4 | **150.7** tok/s (1412 ms) | 16.1 tok/s (5270 ms), **ok 21/60** |

- **Prefill: vLLM ~2.7× faster and flat from 8k to 128k** (28–29.6k tok/s vs EXL3 ~10–11.6k tok/s); at 128k only 5/9 EXL3 requests per config completed. The TabbyAPI failures are non-JSON SSE chunk errors on concurrent long-prompt ingestion — a TabbyAPI/ExLlamaV3 serving-path limit, not an MTP or bpw issue.
- EXL3 wins **only short single-user chat** (169–173 tok/s at 1k, loads in 1–1.5 min vs 5–6); vLLM wins from 8k context up and at any concurrency > 1, is stable at 32k × 4/8/16, and needs only ~10–12 GB host RAM vs ~45 GB for TabbyAPI's pinned n-gram table.
- Median TPOT at c=1 with MTP: EXL3 **5.5 ms** vs vLLM 6.1 ms — ExLlamaV3's pure decode step is faster; it loses on prefill and scaling. MTP is worth +30% on vLLM, +60% on EXL3 at c=1. **5.05bpw runs at the same speed as 4.05bpw.**
- **Quality: quantization is not the differentiator.** Flash-Next vs incumbent Qwen3.8-27B on tool-eval-bench: 27B scores 91.0 ± 1.5 vs Flash-Next 87.9 ± 1.9 (vLLM, reasoning_effort=low); 27B BF16 scores 89.0 (≤ NVFP4 91.0), EXL3 5.05bpw 86.9 ≈ 4.05bpw 87.4. On AIME 2025+2026 they are equal (48/60 vs 50/60).
- **SGLang result (v3)**: Flash-Next quality 88.0 low / 88.9 thinking-off, `tool_choice="required"` enforced 7–8/8, warm prefix TTFT 210 ms, zero server-side errors — but with a file-backed (not pinned-RAM) n-gram table on the 64 GB host: **37–51 tok/s flat across all points**; cookbook pinned-RAM numbers (148 tok/s single, 613 at 16) not reproduced due to host-RAM constraints.
- Caveats: single workstation; TabbyAPI 21/60 failure mode may be TabbyAPI-version-specific rather than EXL3-intrinsic. Implication for the "only credible raw-TPS challenger to llama.cpp" verdict: keep, sharpened — EXL3's challenge is single-stream short-context decode; on prefill (2.7× deficit), concurrency scaling, and operational footprint, the vLLM NVFP4/FP8 stack dominates. EXL3's bottleneck is serving-path maturity, not kernel speed.

### AWQ and GPTQ in 2026: tooling turnover, deployment nuance

- **AutoAWQ is officially deprecated.** Last tested setup: Torch 2.6.0 + Transformers 4.51.3. The AWQ quantization workflow now lives in **llm-compressor** (vLLM project) — the recommended path — and MLX-LM also supports AWQ on Apple Silicon.
- **GPTQModel** (ModelCloud) is the AutoGPTQ successor: 2/3/4/8-bit quantization with latest-model support; **5,000+ GPTQ repos on Hugging Face** reference it. Hugging Face's quantization overview now lists GPTQModel (not AutoGPTQ) as the GPTQ entry.
- Methods recap: **AWQ** (Activation-aware Weight Quantization) scales weights by activation importance before quantization; consensus best INT4 PTQ quality; W4A16 typical, group size 128; official kernel is the default vLLM path, with **Marlin/Machete** kernels available when large-batch throughput matters. **GPTQ** quantizes weights column-by-column, compensating accumulated rounding error via the inverse Hessian; requires calibration data (~5–30 minutes for 7B); slightly worse quality than AWQ on most benchmarks but **widest hardware compatibility** (including via GGUF); default kernel ExLlamaV2.
- Kernel gates (audited against vLLM v0.26.0 source): **Marlin's floor is SM75 (Turing), not SM80** (`if device_capability < 75: return []` — RTX 20-series not excluded); **Marlin-24 act-order kernels removed from vLLM main** (existed at tag v0.9.0); **Machete requires exactly SM90 (Hopper)** — handles odd layer dims (e.g. vision towers) that Marlin rejects; Conch targets SM80+; CUTLASS native FP8 GEMM needs cc ≥ 89; FP8 *weights* (W8A16) via Marlin-FP8 works from SM75.
- The old `VLLM_NVFP4_GEMM_BACKEND` environment variable is **deprecated**, replaced by `--linear-backend` (vLLM auto-selects CUTLASS/FlashInfer/Marlin at load).
- vLLM quantization→backend mapping (Isambard, Aug 2026): FP8 → `deep_gemm`/`hpc`/`flashinfer_cutlass`; **INT4 AWQ/GPTQ → `marlin`**; NVFP4 → `flashinfer_cutedsl`; W4A8 → `humming`.
- Hard-won deployment rules: Marlin rejects symmetric AWQ `uint4` with `zero_point=false` (AutoRound quants) on Ampere — use a compressed-tensors quant; Marlin needs layer dims divisible by 128 (multimodal vision MLPs with odd dims need Hopper/Machete); AWQ `quant_method` requires `--dtype float16` (bf16 rejected at config validation); vLLM's `--quantization awq_marlin` flag is explicitly required for AWQ checkpoints (config.json alone insufficient).
- On AMD, AWQ/GPTQ/Marlin remain NVIDIA-only — the ROCm paths use FP8 via llm-compressor, GGUF, or AMD Quark.
- Quality consensus (2026): AWQ holds slightly lower perplexity than GPTQ at equal bits; GGUF Q4_K_M sits between them depending on model. A French 2026 industry guide puts well-calibrated 4-bit degradation at **2–4 points on MMLU for 7B+ models**.
- **Verifying the brief's "AWQ as the production standard" claim**: both engines ship AWQ paths — vLLM supports awq/awq_marlin/auto_awq; SGLang supports awq plus awq_marlin, gptq_marlin, and more. A mid-2026 comparison calls AWQ **"the more commonly used of the two [vs GPTQ] for vLLM deployments"** — the "standard" part is supported on vLLM; **6,500+ AWQ repos** on the Hugging Face Hub.
- Nuances against the brief's framing: **workflow friction on 2026 models** — AutoAWQ was archived in May 2025 and never got Qwen3.5 support (hard `TypeError`); the **llm-compressor pip release pins `transformers<=4.57.6`**, predating Qwen3.5 support (transformers 5.x) — source install required; AWQ's layer-by-layer scale search **OOMs on a 40 GB A100** for hybrid architectures (fix: `offload_device=cpu` on the AWQModifier). New-model coverage is where the "standard" frays: AWQ tooling lags architectures, FP8/MXFP4 paths need no calibration at all.
- [COMMUNITY] A 4× RTX PRO 6000 Blackwell shootout (vLLM 0.17.0rc1, TP4) had **AWQ beating both NVFP4 recipes at every concurrency** — C=128: AWQ 3519 vs 3232/3220 tok/s with MTP; 2796 vs 2294/2291 without MTP (single rig; NVFP4 path/backend version-sensitive). Quality deltas on a 19-question math set were mixed.
- Net verdict: **PARTIALLY VERIFIED / NUANCED**. AWQ is the dominant *compat* INT4 path on both engines and the most-used vLLM INT4 format — but not where new quantization work targets (llm-compressor/AutoRound/ModelOpt/FP8/NVFP4 per the decision matrix), and "preserving salient activations" is the algorithm's 2023-era technical edge, not a 2026 differentiator over FP4/FP8.

### The 2026 tooling verdict

- **AutoRound** (Intel Neural Compressor): signSGD-based rounding-optimization PTQ, tunes per-weight rounding against a small calibration set at GPTQ-comparable cost. Intel's lm-eval comparisons: superior to GPTQ (30/32 configs) and AWQ (27/32), OmniQuant (16/16). Independent September 2026 test (Intel Arc Pro B70 cookbook, Ornith-1.5-35B-A3B, vLLM-XPU): logprob parity vs BF16 — AutoRound **91.4% top-1 agreement, KL 0.277** vs GPTQ 90.5% / 0.319 — equal-or-best on all primary metrics. Composes with expert pruning: **REAP (50% expert pruning) + AutoRound W4A16 took GLM-4.7 from 700 GB to 92 GB (~6.5× total compression)**. Status: **strongest 4-bit PTQ**.
- **HQQ / gemlite** (mobiusml): data-free, no calibration, quantizes the largest models in minutes; 1/2/3/4/8 bits; linear dequantization compatible with optimized CUDA/Triton kernels; PEFT-compatible; recommended start `nbits=4, group_size=64, axis=1`. **HQQ+** adds trainable low-rank adapters to recover quality at lower bits. [COMMUNITY/DIRECTIONAL] ~2× speedup on H100/A100 at 4-bit with ~11–13% perplexity increase on Llama-2 (community research compilation, not peer-reviewed). **gemlite** (Triton kernel library): 7–8× faster prefill, 3–6× faster decode vs default PyTorch AO kernels; ships MXFP formats end-to-end — A16W8_MXFP, A16W4_MXFP, A8W8_MXFP_dynamic, **A4W4_MXFP_dynamic**, **A4W4_NVFP_dynamic (Blackwell)** — plus HQQ weight-only (4/2-bit) and 1.58-bit ternary; automatic kernel selection per matrix shape with GPU-specific autotuned configs.
- **AQLM** (Egiazarian et al., arXiv:2401.06166): consistently cited in 2026 reports as **best-in-class at 2-bit** [UNVERIFIED — evidence base remains the original 2024 paper plus derivative citations; no 2026 independent production comparison], without production tooling or kernel ecosystem — the research reference for the 2-bit tier.
- **NVIDIA ModelOpt** (`pip install nvidia-modelopt`, Apache 2.0): the toolkit turning a checkpoint into a deployable quantized artifact for TensorRT-LLM, vLLM, and SGLang. **0.46.0 (2026-08-18)** highlights: NVFP4 and FP8 **PTQ recipes** with projection-output quantizers; **Learned Scale Quantization (LSQ) and Dual-LSQ** for quantization-aware distillation; **fused Triton fast path for `local_hessian` NVFP4 weight-scale search — ~34× faster** (bit-exact); **NVFP4 Four-Over-Six (4/6)** and MXFP8 checkpoint support in vLLM via `hf_quant_config.json`. 0.45.0: **pruning and quantization compose multiplicatively** — Minitron structured pruning + FP8 PTQ on Nemotron-3-Nano-30B = 2.6× throughput and 2.6× memory. Supports INT8, FP8, INT4, NF4, NVFP4 (weight-only, channel, block schemes).
- **optimum-quanto: maintenance mode** — the definitive loser call. Its README: "Major new features or breaking changes are unlikely to be merged. For production-ready quantization features or active development, consider alternative projects such as bitsandbytes or torchAO." Remains functional (int2/int4/int8/float8 weights, int8/float8 activations, CUDA kernels) but no longer where HF invests.
- **bitsandbytes**: refactored to multi-backend — **ROCm (AMD) and Intel CPU implementations mature**, Intel XPU in progress, Apple Silicon expected Q4 2026/Q1 2027. NF4 + QLoRA remains the standard single-GPU fine-tuning path.
- **torchAO** (pytorch/ao): HF's recommended active alternative alongside bitsandbytes; torchao int4-weight is the only int4 path on Apple Silicon (Metal).
- The 4-bit PTQ crown moved from AWQ to AutoRound — anyone still calibrating with AutoAWQ is on a deprecated path; the 2026 recommendation is AutoRound (accuracy) or HQQ (speed/data-free), quantized through llm-compressor or ModelOpt.

### Tooling verdict table

| Tool | 2026 status | Niche |
|---|---|---|
| llm-compressor | **Active, recommended** | AWQ/FP8/NVFP4/MXFP4 for vLLM; MoE-aware `--moe` |
| AutoRound | **Strongest 4-bit PTQ** | Best measured accuracy; composes with pruning |
| HQQ / gemlite | Active | Data-free/fast quant; Triton kernels incl. MXFP4/NVFP4 |
| GPTQModel | **Active** | AutoGPTQ successor; 5,000+ HF repos |
| NVIDIA ModelOpt | **Active** | Blackwell-native; NVFP4/FP8 PTQ+QAT recipes |
| bitsandbytes | Active | NF4/QLoRA fine-tuning; multi-backend refactor |
| torchAO | Active | HF-recommended; Metal int4 |
| AutoAWQ | **Deprecated** | Migrated to llm-compressor |
| optimum-quanto | **Maintenance mode** | HF redirects to bitsandbytes/torchAO |
| AQLM | Research | Best-in-class 2-bit, no production ecosystem |
| ExLlamaV2 | **Archived** | Legacy EXL2 checkpoints only |

