---
id: ai-industry-kb-2026/08-quantization-model-formats/main-actors-continued
title: "Main actors (continued)"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "AWS", "Alibaba", "Apple", "DeepSeek", "Falcon", "Google", "Hugging Face", "Intel", "Meta", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "Qualcomm", "SGLang", "TII", "Unsloth", "Z.ai", "vLLM"]
dates: ["2023-09", "2025-05", "2026-02-13", "2026-02-20", "2026-03-04", "2026-03-17", "2026-05", "2026-05-11", "2026-05-13", "2026-05-28", "2026-07", "2026-07-15", "2026-07-26", "2026-08-15", "2026-08-18", "2026-08-19", "2026-08-27", "2026-09", "2026-09-18", "2026-09-22"]
keywords: ["accelerator", "agentic", "amd", "apache", "attention", "attribution", "awq", "aws", "benchmark", "benchmarks", "bitnet", "blackwell"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4432, 4577]
section: "8. Quantization & Model Formats"
sha256: c8de6ae5f59ef5b8869d3508a2a0a2b1c176fd98a1d8fe7ecfe2276209569878
---

# Main actors (continued)

## Main actors (continued)

- **NVIDIA** — NVFP4 format owner; tcgen05 FP4 tensor cores on Blackwell (SM100/103); claimed ~2× FP8 GEMM compute and B200 ≈ 8× FP16 [VENDOR]; ModelOpt (0.46.0, 2026-08-18: Triton fast path 34×, LSQ/Dual-LSQ, NVFP4 4/6, MXFP8); DeepSeek-R1-FP4 (671B) production proof; Kimi-K2.6-NVFP4 posted 2026-05-13; Nemotron-3-Nano-30B NVFP4 via QAD (99.4% of BF16) [VENDOR]; official NVFP4 guidance (3.5× vs FP16, 1.8× vs FP8, <1% loss).
- **OpenAI** — GPT-OSS ships natively in MXFP4 (2025-08): GPT-OSS-120B on a single H100, GPT-OSS-20B in 16 GB; MXFP4 credited with ~75% inference-cost reduction.
- **Open Compute Project (OCP)** — MX specification v1.0 (September 2023): AMD, Arm, Intel, Meta, Microsoft, NVIDIA, Qualcomm; E2M1 blocks of 32 with E8M0 scales; the open standard behind MXFP4.
- **AMD** — Instinct MI355X: Dell measured up to 6.1× inference throughput with MXFP4 on PowerEdge XE9785L; MI350 partial MXFP4 support; AMD Quark on ROCm paths.
- **Intel** — AutoRound: strongest measured 4-bit PTQ (lm-eval 30/32 vs GPTQ, 27/32 vs AWQ; 91.4% top-1 logprob agreement); **OpenVINO 2026.2**: INT4 KV-cache compression on Intel GPU (BY_CHANNEL keys group-16, BY_TOKEN values on paged attention).
- **Microsoft Research** — BitNet: natively trained 1.58-bit ternary models (b1.58 2B4T, MIT); at 3B matches FP16 LLaMA perplexity/zero-shot; the 100B-on-single-CPU claim remains theoretical.
- **Moonshot AI** — Kimi K2.6 (1T total, 32B active): ships INT4 natively via QAT (~2× speedup); NVIDIA's Kimi-K2.6-NVFP4 (2026-05-13) makes it a Blackwell-format reference.
- **DeepSeek** — DeepSeek-R1-FP4 (671B: MMLU −0.1, GSM8K −0.2 vs FP8 base); DeepSeek-V3.2/R1 NVFP4 on GB300 (vLLM blog 2026-02-13; R1: 22,476 TGS prefill-only on 2× GB300); DeepSeek-V4: FP4 (MXFP4) QAT on MoE experts + indexer QK path; community DeepSeek-V4-Flash-NVFP4-FP8-MTP quant (May 2026, 172 GB).
- **Unsloth** — NVFP4 export since July 2026 ("Dynamic Unsloth NVFP4 Quants": gemma-4-31B-it-NVFP4, gemma-4-26B-A4B-it-NVFP4, gemma-4-12b-it-NVFP4; hybrid NVFP4-GGUF containers such as Qwen3.6 NVFP4-GGUF, accelerated only via Unsloth kernels on Blackwell); 1,374-model Hugging Face catalog (GGUF + NVFP4).
- **PrismML** — Bonsai-27B (2026-07-15, Apache 2.0): first 27B model on a phone; 1-bit "Binary g128" (3.9 GB, >90% retention, 11 tok/s iPhone 17 Pro) and ternary (7.2 GB, >95% retention); near-lossless 4-bit KV enabling 262K context on-device.
- **Meta** — Qwen3.8-27B (2026-08-15) as the Dynamic v3.0 launch vehicle; Meta Muse covered by v3.0 quants.
- **Alibaba / Qwen** — Qwen-Coder 32B: the FP8 accuracy reference datapoint (79–80% vs ~80% FP16); Qwen3.6-27B/35B-A3B NVFP4 pre-quantized checkpoints.
- **Hugging Face** — absorbed ggml.ai (announced 2026-02-20; MIT license and team autonomy preserved), cementing GGUF as the de-facto local format; 6,500+ AWQ repos on the Hub; optimum-quanto in maintenance mode (confirmed 2026-09-22, redirects to bitsandbytes/torchAO); quantization overview lists GPTQModel as the GPTQ entry.
- **Red Hat** — independently evaluated TurboQuant and drove its vLLM integration (May 2026; PR #38479): FP8 remains the best KV-cache default; publishes quantized checkpoints (e.g. Qwen3.5-122B-A10B-NVFP4).
- **Google** — TurboQuant: the original method (arXiv:2504.19874, ICLR 2026); attribution reconciled against earlier Red-Hat-as-inventor wording.
- **Community / independent actors** — srmiles (first independent Dynamic v3.0 replication test, 2026-08-27); MarcoPizeta (vLLM/SGLang/EXL3 shootout, Sept 2026); rtx6kpro (AWQ > NVFP4 decode shootout, Sept 2026); avarok (NVFP4 W4A4 MoE on GB10, production-ready); canada-quant (2026-05-28 Marlin-fallback measurements); 0xsero (Blackwell dtype matrices); club-3090 (consumer-Blackwell dtype matrix); massed-compute (Nemotron NVFP4 benchmarks, 2026-08-18); TII (Falcon3-10B-Instruct-1.58bit, Falcon-E 1B/3B ternary); mobiusml (HQQ data-free quantization, gemlite Triton kernels with MXFP4/NVFP4 dynamic).

- **TII** — Falcon3-10B-Instruct-1.58bit (~2 GB GGUF) and Falcon-E 1B/3B (natively trained ternary from scratch on ~1.5T tokens; 665 MB / 999 MB disk).
- **Dell** — measured up to 6.1× inference throughput with MXFP4 on PowerEdge XE9785L + AMD Instinct MI355X [VENDOR].
- **Artefact2** — GGUF quantization overview: the community-standard per-type KL/PPL table measured on real hardware [COMMUNITY].
- **mradermacher** — per-type model cards converging on the IQ-quant ladder ("IQ-quants are often preferable over similar sized non-IQ quants"; DeepSeek-V2.5 236B-class i1-IQ4_XS = 125.7 GB at ~0.8% perplexity increase) [COMMUNITY].
- **Intel Arc Pro B70 cookbook (sergiob)** — independent AutoRound-vs-GPTQ logprob-parity test (Ornith-1.5-35B-A3B, vLLM-XPU): AutoRound 91.4% top-1 agreement / KL 0.277 vs GPTQ 90.5% / 0.319 [COMMUNITY].
- **Red Hat AI engineers (Eldar Kurtic, Michael Goin, Alexandre Marques)** — TurboQuant vLLM evaluation, published 2026-05-11; drove the vLLM integration (PR #38479).
- **0xsero / club-3090 / massed-compute / canada-quant / srmiles / s0me1-dev** — community dtype matrices (NVFP4 KV datacenter-only; consumer-Blackwell W4A16-vs-W4A4 guidance), throughput/$ benchmarks, the 2026-05-28 Marlin-fallback measurement, the first Dynamic v3.0 independent test, and SM120 NVFP4 GEMM patches [COMMUNITY].
- **NextPlatform** — AWS Blackwell rental economics analysis: FP8 halves cost-per-teraflop vs FP16; FP4 halves it again [ANALYST].

## Timeline and context (continued)

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

- **Dynamic v3.0 independence**: partially addressed by the 2026-08-27 community test (KLD corroborated, 9.5% decode-speed regression measured, one UD variant shipped a llama-server bug) — Divergence-300 specifically and benchmark-score replication still open.
- **FP4 KV on consumer Blackwell**: stock vLLM path datacenter-only; community FA2+XQA route (1.6× fp8 KV pool, July 2026) needs wider validation before production use.
- **BitNet at scale**: the 100B-on-a-single-CPU claim is theoretical; whether natively-trained ternary models scale past 3B without quality collapse is unanswered (Bonsai's converted 1-bit at 27B is the parallel bet).
- **MXFP4 beyond NVIDIA**: Dell's 6.1× datapoint on AMD MI355X suggests the OCP standard could become the cross-vendor 4-bit default; AMD next-gen and Intel accelerator coverage is the thing to track.
- **Sub-2-bit training economics**: Bonsai proved conversion-to-1-bit at 27B with >90% retention; native 1-bit training (BitNet-style) at 27B+ has not been demonstrated publicly.
- **NVFP4 W4A16 vs W4A4 confusion at download time**: checkpoint naming and model-card flags remain the failure point — always check the card before downloading.
- **AWQ's tooling tail**: the format remains the dominant compat INT4 path, but new-model coverage (Qwen3.5+), llm-compressor pinning, and calibration-OOM workarounds show where the "standard" frays.

- **NVFP4 W4A16 vs W4A4 confusion at download time**: checkpoint naming and model-card flags remain the failure point — always check the card before downloading; the industry needs machine-readable dtype flags.
- **AutoRound vs llm-compressor adoption split**: AutoRound is the strongest measured 4-bit PTQ, but Marlin rejects its symmetric `uint4`/`zero_point=false` form on Ampere — adoption hinges on compressed-tensors export paths.
- **INT4 KV in vLLM/SGLang mainline**: OpenVINO 2026.2 shipped it on Intel GPU; whether vLLM/SGLang mainline adopt INT4 KV (vs the FP8 default) decides long-context economics on Hopper hardware.
- **MXFP4 kernel coverage on AMD next-gen and Intel accelerators** beyond the single MI355X datapoint — the open-standard bet's progress metric.
- **Consumer-Blackwell W4A4 maturity**: GB10 production-ready and RTX 5090 software-dequant routes exist, but the stock vLLM NVFP4 KV/attention path remains datacenter-only — watch whether SM120 gets first-class FP4 KV in 2027.
- **Quantization labor vs format lock-in**: QAT's <1%-loss promise commits a model to one quantization target; the retraining cost means PTQ-plus-good-kernels (FP8, MXFP4 via llm-compressor) keeps winning on economics unless the deployment is pinned to one format for years.

## Sources and URLs (continued)

- https://github.com/open-hadis/rvllm/blob/HEAD/docs/boost/07-fp8-quantization.md [OPEN-SOURCE]
- https://github.com/mithudso/skills/blob/HEAD/llm-quantization-strategies/SKILL.md [COMMUNITY]
- https://github.com/s-samarth/datasciencepreparation/blob/HEAD/LLM/docs/inference-arch/frontier-techniques.md [COMMUNITY]
- https://github.com/harshuljain13/llm-inference-at-scale/blob/HEAD/content/05_optimization/04.1_quantization/quantization.md [COMMUNITY]
- https://github.com/xinhaoc/ferret/blob/HEAD/docs/architecture/blackwell-b200.md [COMMUNITY]
- https://github.com/allanschramm/local-model-autotuning/blob/HEAD/docs/discovery/nvfp4-quantization.md [COMMUNITY]
- https://github.com/noonghunna/club-3090/blob/HEAD/docs/DTYPE_MATRIX.md [COMMUNITY]
- https://www.gpuyard.com/tutorials/howto/deploy-llm-blackwell-fp4/ [COMMUNITY]
- https://github.com/canada-quant/dsv4-flash-nvfp4-fp8-mtp [COMMUNITY]
- https://www.opencompute.org/documents/ocpmx-v1-0-spec-final.pdf [STANDARDS]
- https://github.com/jino-rohit/llm-compressor/blob/HEAD/docs/guides/compression_schemes.md [OPEN-SOURCE]
- https://huggingface.co/blog/RakshitAralimatti/learn-ai-with-me [COMMUNITY]
- https://www.redpacketsecurity.com/how-openai-used-a-new-data-type-to-cut-inference-costs-by-75/ [COMMUNITY]
- https://infohub.delltechnologies.com/static/media/client/7phukh/DAM_49b1e26b-f10f-418c-9697-02a9c31fb7d7.pdf [VENDOR]
- https://github.com/avifenesh/memra/blob/HEAD/research/fp4-act-scoping-20260806/BRIEF.md [COMMUNITY]
- https://startupfortune.com/nvidia-puts-kimi-k26-on-a-faster-path-to-blackwell-inference/ [COMMUNITY]
- https://github.com/0xsero/blackwell-gpu-wiki/blob/HEAD/docs/fundamentals/number-formats.md [COMMUNITY]
- https://blog.avarok.net/nvfp4-w4a4-moe-inference-on-nvidia-blackwell-gb10-1a83e85d0f9e?gi=b51174f3e648 [COMMUNITY]
- https://arxiv.org/pdf/2601.09527v1 [ACADEMIC]
- https://github.com/0xsero/blackwell-gpu-wiki/blob/HEAD/docs/case-studies/generic-moe-on-consumer-blackwell.md [COMMUNITY]
- http://build.nvidia.com/spark/nvfp4-quantization [VENDOR]
- http://quantumzeitgeist.com/faster-blackwell-gpus-rag-inference-private/ [COMMUNITY]
- https://www.pulse.bot/ai/news/nvidia-ai-brings-nemotron-3-nano-30b-to-nvfp4-with-quantization-aware-distillation-qad-for-efficient-5bb90468-5981-4007-bd25-0fb90468-5981-4007-bd25022/ [VENDOR]
- https://www.nextplatform.com/cloud/2025/07/10/sizing-up-aws-blackwell-gpu-systems-against-prior-gpus-and-trainiums/1651108 [ANALYST]
- https://github.com/vllm-project/vllm-project.github.io/blob/HEAD/_posts/2026-02-13-gb300-deepseek.md [VENDOR]
- https://github.com/massed-compute/gpu-benchmark/blob/HEAD/nemotron-3.5-lightning-30b/nemotron-3.5-lightning-30b.md [COMMUNITY]
- https://github.com/kubesimplify/website/blob/HEAD/content/blog/day-4-quantization-demystified-bf16-fp8-nvfp4-mxfp4-int4-gguf-and-why-it-all-matters.md [COMMUNITY]
- https://github.com/local-inference-lab/rtx6kpro/blob/HEAD/benchmarks/nvfp4-quantization-comparison.md [COMMUNITY]
- https://github.com/wanshuiyin/aris-in-ai-offer/blob/HEAD/docs/tutorials/quantization_tutorial_en.md [COMMUNITY]
- https://github.com/0-co/company/blob/HEAD/research/bitnet-deep-dive-2026-03-17.md [COMMUNITY]
- https://www.freshlab.es/blog/microsoft-bitnet-local-llm-cpu-no-gpu [COMMUNITY]
- https://emelia.io/hub/bitnet-1bit-llm-cpu-inference [COMMUNITY]
- https://medium.com/@kkhushi/the-era-of-1-bit-llms-all-large-language-models-are-in-1-58-bits-2f113032a9fe [COMMUNITY]
- https://huggingface.co/prism-ml/Ternary-Bonsai-27B-gguf [VENDOR]
- https://dev.to/pneumetron/prism-ml-introduces-bonsai-27b-a-1-bit-llm-for-on-device-inference-2hhj [COMMUNITY]
- https://www.manilatimes.net/2026/07/15/tmt-newswire/plentisoft/prismml-announces-1-bit-bonsai-27b-the-first-27b-model-to-run-on-a-phone/2384502 [COMMUNITY]
- https://medium.com/macoclock/bonsai-27b-runs-on-a-16-gb-m4-mac-mini-with-4-2-gb-of-ram-1-bit-quantization-with-mlx-662a30587822 [COMMUNITY]
- https://medium.com/openvino-toolkit/int4-kv-cache-compression-for-llm-inference-on-intel-gpu-new-in-openvino-2026-2-d71d03c27897 [VENDOR]
- https://github.com/dragonshadows1978/ai-atlasforge/blob/HEAD/KV_CACHE_QUANTIZATION_TECHNICAL_REFERENCE.md [COMMUNITY]
- https://github.com/blokzdev/blokz/blob/HEAD/content/articles/2026/06/the-memory-wall-kv-cache-and-why-long-context-ai-cant-decentralize/index.mdx [COMMUNITY]
- https://arxiv.org/pdf/2604.19157v1 [ACADEMIC]
- https://ayinedjimi-consultants.fr/static/pdf/quantization-llm-2026-gguf-gptq.pdf [COMMUNITY]
- https://huggingface.co/docs/transformers/v4.49.0/quantization/overview [VENDOR]
- https://medium.com/intel-analytics-software/autoround-sota-weight-only-quantization-algorithm-for-llms-across-hardware-platforms-99fe6eac2861 [VENDOR]
- https://github.com/sergiiob/intel-arc-pro-b70-inference-cookbook/blob/HEAD/docs/ornith15-35a3/AUTOROUND-VS-GPTQ.md [COMMUNITY]
- https://github.com/nvidia/model-optimizer/blob/HEAD/CHANGELOG.rst [VENDOR]
- https://github.com/vllm-project/vllm-project.github.io/blob/HEAD/_posts/2026-05-11-turboquant.md [COMMUNITY]
- https://github.com/vllm-project/vllm/pull/38479 [OPEN-SOURCE]
- https://www.marktechpost.com/2026/09/18/gguf-vs-gptq-vs-awq-vs-exl2-llm-model-formats-explained-2026/ [COMMUNITY]

