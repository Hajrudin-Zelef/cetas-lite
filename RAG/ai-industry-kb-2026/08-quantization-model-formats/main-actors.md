---
id: ai-industry-kb-2026/08-quantization-model-formats/main-actors
title: "Main actors"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Hugging Face", "Intel", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM"]
dates: ["2022-10-31", "2023-08-22", "2026-02-20", "2026-03-07", "2026-05", "2026-05-13", "2026-07-06", "2026-07-14", "2026-08-15", "2026-08-18", "2026-08-19", "2026-08-20", "2026-08-21", "2026-08-27", "2026-08-29", "2026-08-31", "2026-09", "2026-09-05", "2026-09-08", "2026-09-09", "2026-09-18", "2026-09-22"]
keywords: ["accelerator", "acquisition", "amd", "awq", "benchmark", "benchmarks", "bitnet", "blackwell", "compute", "consumer", "datacenter", "decode"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4089, 4216]
section: "8. Quantization & Model Formats"
sha256: 52202bc1fe5bf45967274bdc8784996cfa35fb5caf5dd7f35e649bf982c44226
---

# Main actors

## Main actors

| Actor | Role in the 2026 format/tooling story |
|---|---|
| **Hugging Face** | Absorbed ggml.ai (2026-02-20); hosts tens of thousands of GGUF checkpoints with metadata viewer and JS parser; redirects optimum-quanto users to bitsandbytes/torchAO |
| **ggml.ai / Georgi Gerganov** | llama.cpp and GGUF creator; team now HF employees with full technical leadership and MIT licensing preserved; 126,000+ GitHub stars by 2026 |
| **Unsloth** | Dynamic v3.0 selective-layer GGUF quants (UD- prefix, 1,374-model catalog, pure PTQ); proprietary `unsloth_zo` schedule; NVFP4-GGUF hybrid containers |
| **NVIDIA** | NVFP4 format; ModelOpt 0.46.0 Blackwell-native toolchain; DeepSeek-R1-FP4 and Kimi-K2.6-NVFP4 production proofs; builds FP4 Tensor-Core hardware (SM100/103/120) |
| **turboderp (ExLlamaV3)** | Active successor to archived ExLlamaV2; EXL3/QTIP format; v1.4.6–v1.4.9 September 2026 patch cadence; vLLM-EXL3 third-party serving plugin |
| **vLLM project** | Datacenter inference stack; llm-compressor (FP8/AWQ/NVFP4/MXFP4); kernel gates (Marlin SM75, Machete SM90, NVFP4→flashinfer_cutedsl); `--linear-backend` replaces `VLLM_NVFP4_GEMM_BACKEND` |
| **SGLang** | Parallel datacenter stack: awq, awq_marlin, gptq_marlin, modelopt_fp4, mxfp4, fp8; `--load-format gguf` (NVIDIA-only) |
| **Intel** | AutoRound — strongest measured 4-bit PTQ (91.4% logprob top-1 agreement vs BF16); Neural Compressor; OpenVINO 2026.2 INT4 KV on Intel GPU; vLLM-XPU validation |
| **mobiusml** | HQQ (data-free/fast 1/2/3/4/8-bit) and gemlite (Triton kernels incl. MXFP4/NVFP4 dynamic); HQQ+ low-rank adapters |
| **ModelCloud** | GPTQModel — AutoGPTQ successor, 5,000+ HF repos reference it |
| **Red Hat AI** | May 2026 vLLM evaluation of TurboQuant KV-cache variants (see 08b); publishes quantized checkpoints (e.g. Qwen3.5-122B-A10B-NVFP4) |
| **Microsoft Research** | BitNet natively-trained 1.58-bit ternary models (see 08b); contributes to OCP MX spec |
| **PrismML** | Bonsai-27B 1-bit/ternary GGUFs (see 08b); `prism-ternary` experimental vLLM-GGUF plugin branch |
| **mradermacher** (community) | i1- imatrix GGUF cards documenting the per-type ladder across DeepSeek-V2.5, QwQ-32B, Qwen2.5-14B, Gemma-2-9B; ARM edge variants Q4_0_4_4 / Q4_0_4_8 / Q4_0_8_8 |
| **Artefact2** (community) | Reference per-type KL/PPL table used as the GGUF accuracy standard |
| **MarcoPizeta** (community) | Reproducible vLLM-vs-EXL3 shootout on RTX PRO 6000 (Sept 2026) |
| **AMD** | ROCm path (ROCm 6.4.4 beats 7.0.1 3.47× on llama.cpp); MI355X + MXFP4 6.1× measured (Dell); AMD Quark on ROCm |
| **Open Compute Project** | MX specification v1.0 (Sept 2023) — the open microscaling standard behind MXFP4 |

## Timeline and context

| Date | Event |
|---|---|
| 2022-10-31 | GPTQ paper (Frantar, Ashkboos, Hoefler, Alistarh) on arXiv; ICLR 2023 |
| 2023 | AWQ paper (Lin et al., arXiv 2306.00978); QLoRA / NF4 (Dettmers et al.); OCP MX specification v1.0 (Sept 2023) |
| 2023-08-22 | **GGUF v1** released by llama.cpp (Georgi Gerganov), superseding GGML |
| 2025-08 | OpenAI GPT-OSS ships natively in **MXFP4** (120B on single H100) |
| 2026-01 | arXiv 2601.09527: NVFP4 consumer-Blackwell economics (2× RTX 5090, Qwen3-8B at $0.002/MTok) |
| 2026-02-20 | **Hugging Face absorbs ggml.ai** (Discussion #19759; MIT license, team autonomy) |
| 2026-03-07 | Community ADRs standardize on GGUF+Ollama for heterogeneous fleets |
| 2026-05-13 | NVIDIA posts **Kimi-K2.6-NVFP4** (1T params, 32B active) on Hugging Face, ModelOpt 0.44.0 |
| 2026-07-14 | ExLlamaV3 v1.0.0: Marlin-inspired GEMM, drops flash-attn-2/xformers deps |
| 2026-08-15 | Qwen3.8-27B launch |
| 2026-08-18 | NVIDIA ModelOpt **0.46.0** (Triton fast path 34×, LSQ/Dual-LSQ, NVFP4 4/6, MXFP8) |
| 2026-08-19 | Unsloth **Dynamic v3.0** official release — SUPERSEDES Dynamic 2.0 |
| 2026-08-21 | llama.cpp upstream digest: v0.2.0 semver, Kimi K3, ROCm/Vulkan, DSpark, --mmproj-device |
| 2026-08-27 | First independent Dynamic v3.0-vs-v2.0 test (srmiles): direction confirmed on KLD; 9.5% slower decode; UD-Q3_K_XL llama-server bug |
| 2026-08-31 | ExLlamaV3 **v1.4.5**; **ExLlamaV2 archived** (end of the EXL2 era) |
| early Sep 2026 | ExLlamaV3 v1.4.6–v1.4.9 (bug-fix cadence, no format changes) |
| 2026-09-09 | vllm-exl3 plugin v0.4.2 fixes dense EXL3 dispatch (engine-wedging cooperative trellis GEMM) |
| 04–08 Sep 2026 | MarcoPizeta vLLM-vs-EXL3 shootout on RTX PRO 6000: EXL3 wins single-stream short decode only; vLLM prefills ~2.7× faster |
| Sep 2026 | AWQ beats both NVFP4 recipes on decode at every concurrency (4× PRO 6000 Blackwell, [COMMUNITY]); AutoRound-vs-GPTQ logprob-parity test (91.4% vs 90.5%); practitioner logs: AWQ needs explicit `--quantization awq_marlin` in vLLM |
| 2026-09-18 | MarkTechPost overview: **AutoAWQ officially deprecated**; llm-compressor is the AWQ path |
| 2026-09-22 | HF confirms optimum-quanto is in **maintenance mode** (redirects to bitsandbytes/torchAO) |

## Implications

1. **Format choice is now a hardware decision, not a quality debate.** GGUF K-quants (preferably UD- or imatrix variants) own CPU/Apple/consumer/heterogeneous hardware; FP8 owns Hopper/Ada/Blackwell serving; NVFP4 is the Blackwell-native future (see 08b for the FP8/FP4 production arc).
2. **The 4-bit PTQ crown moved from AWQ to AutoRound** — anyone still calibrating with AutoAWQ is on a deprecated path; the 2026 recommendation is AutoRound (accuracy) or HQQ (speed/data-free), quantized through llm-compressor or ModelOpt.
3. **The brief's "AWQ as production standard" is true only in the compat sense**: AWQ remains the most-used vLLM INT4 format and the more-commonly-used of AWQ-vs-GPTQ, but its tooling lags new architectures (no Qwen3.5 in AutoAWQ; llm-compressor pins transformers ≤4.57.6), and new work should not target it.
4. **Watch Dynamic v3.0's independence claims**: the >10% top-1 accuracy claim now has KLD-based independent corroboration (direction only, one hardware setup, one model), but Divergence-300 is unreplicated and a 9.5% decode-speed regression was measured — accuracy gains are not free.
5. **EXL3's bottleneck is serving-path maturity, not kernel speed**: the vLLM-EXL3 plugin needed engine-wedging fixes as recently as 2026-09-09, and TabbyAPI/ExLlamaV3 failed 39/60 requests at 32k×4 where vLLM stayed stable. Keep EXL3 for single-stream short-context decode on consumer NVIDIA; use vLLM's stack for everything batched.
6. **Don't conflate the two FP4s**: NVFP4 (block-16, FP8-E4M3 scales, FP32 tensor scale, NVIDIA-proprietary, ~4.5 bits/elem) vs MXFP4 (block-32, E8M0, open spec, AMD/Intel paths). Checklist: MXFP4 is Hopper+ (no Ampere kernel — gpt-oss-120b does not run on A100); NVFP4 W4A4 needs Blackwell; W4A16 rides the Marlin fallback.
7. **The tooling loser list is a procurement signal**: deprecated AutoAWQ, maintenance-mode optimum-quanto, archived ExLlamaV2 — pin new pipelines to llm-compressor, ModelOpt, GPTQModel, AutoRound, HQQ, bitsandbytes, or torchAO.
8. **KV-cache quantization is now a first-class lever alongside weight quantization** — it moves the INT4 crossover context from ~128K to ~512K and decides what context lengths are actually deployable; detail in 08b.
9. **AMD and Intel are no longer second-class quant citizens** — bitsandbytes multi-backend (ROCm/Intel CPU mature), OpenVINO INT4 KV on Intel GPU, MXFP4 6.1× on AMD MI355X, ROCm 6.4.4 beating 7.0.1 by 3.47× on llama.cpp.
10. **Watch the open items, not the marketing**: Divergence-300 replication for Dynamic v3.0 (KLD corroborated, benchmark-score replication missing); whether the vLLM GGUF plugin graduates from "highly experimental and under-optimized" to a supported path or remains a conversion-constrained bridge; MXFP4 coverage beyond NVIDIA (MI355X datapoint vs AMD next-gen and Intel accelerator plans); and whether mainline llama.cpp adopts TurboQuant-style KV compression (08b), which would change long-context economics on consumer hardware.

## Sources and URLs

- [VENDOR] GGUF format overview — https://en.wikipedia.org/wiki/GGUF
- [COMMUNITY] GGUF vs GPTQ vs AWQ vs EXL2 (MarkTechPost, 2026-09-18; AutoAWQ deprecated, EXL2 bpw naming, TabbyAPI) — https://www.marktechpost.com/2026/09/18/gguf-vs-gptq-vs-awq-vs-exl2-llm-model-formats-explained-2026/
- [COMMUNITY] vLLM GGUF plugin (vllm-project) — https://github.com/estridell/vllm-gguf-plugin
- [COMMUNITY] vLLM quantization deep-dive (GGUF out-of-tree migration, RFC #39583) — https://github.com/allanschramm/local-model-autotuning/blob/HEAD/docs/discovery/vllm-quant-deep-dive.md
- [COMMUNITY] vLLM course best-practices (GGUF "highly experimental and under-optimized") — https://github.com/akgaur12/developer-notes/blob/HEAD/AI-ML/vllm-course/21-best-practices.md
- [COMMUNITY] Inference engines landscape (quantization equivalence table) — https://github.com/allanschramm/local-model-autotuning/blob/HEAD/docs/discovery/inference-engines-landscape.md
- [COMMUNITY] GGUF Dynamic Quantization on GPU Cloud (Spheron blog, Dynamic 2.0 benchmarks) — https://www.spheron.network/blog/gguf-dynamic-quantization-gpu-cloud/
- [VENDOR] Unsloth Hugging Face org (1,374 models, Dynamic 2.0 + NVFP4) — https://huggingface.co/unsloth
- [COMMUNITY] GLM-5.2 local deployment with Unsloth Dynamic GGUFs (DailySynapse) — https://dailysynapse.com/news/glm-5-2-can-run-locally-with-unsloth-dynamic-ggufs/
- [COMMUNITY] orka.py commit (Unsloth OSS delegates to llama-quantize; Dynamic schedule proprietary) — https://github.com/orkait/orka.py/commit/f32b24f99fdf0d96985085a560ea8272c3380847
- [COMMUNITY] LocoLLM ADR-0006 (GGUF + Ollama as inference standard, 2026-03-07) — https://github.com/michael-borck/loco-llm/blob/HEAD/src/content/docs/adr/0006-gguf-ollama-inference-standard.md
- [COMMUNITY] Local LLM capability matrix (quantization strategy table) — https://github.com/hhalperin/ai-briefcase/blob/HEAD/docs/design/tier1/1.11-llm-inference/local-llm-capability-matrix.md
- [COMMUNITY] Fastest TPS inference engine (ExLlamaV2 archived → ExLlamaV3 v1.4.5 2026-08-31, EXL3/QTIP 1.6 bpw) — https://github.com/allanschramm/local-model-autotuning/blob/HEAD/docs/discovery/fastest-tps-inference-engine.md
- [VENDOR] Unsloth Dynamic v3.0 announcement (BigGo, 2026-08-19; >10% top-1, 5.1M downloads, UD-Q2_K_XL 9.83GB, pure PTQ) — https://finance.biggo.com/news/8d1ede07-a4cd-4497-b4c5-ab2e8e16bbdc
- [COMMUNITY] Unsloth Dynamic 3.0 GGUFs released (AIToolly, 2026-08-20) — https://aitoolly.com/ai-news/article/2026-08-20-unsloth-dynamic-30-ggufs-released-delivering-10-better-accuracy-for-local-llm-quantization
- [COMMUNITY] Unsloth Dynamic 3.0 technical analysis (malaiwah, 2026-08-20; per-layer type allocation, imatrix >1.5M tokens) — https://github.com/malaiwah/qwen38-27b-exl3/blob/HEAD/receipts/unsloth-dynamic3-comparison-2026-08-20.md
- [VENDOR] Unsloth Dynamic 3.0 docs (official; Divergence-300, KL, methodology) — https://unsloth.ai/docs/basics/dynamic-3.0-ggufs
- [COMMUNITY] Qwen3.8-27B-GGUF discussion (unsloth, 2026-08-29; 1-bit 77% accuracy, MTP separation, non-dynamic quants deleted) — https://huggingface.co/unsloth/Qwen3.8-27B-GGUF/discussions/74
- [COMMUNITY] Qwen3.8-27B GGUF quantization research doc (UD-Q4_K_XL 16.69 GiB, top-1 agreement readings) — https://github.com/xenodeve/qwen-3.8-27b-tuning/blob/HEAD/docs/researchs/Qwen3.8-27B_Optimization_Research_Docs/02-UNSLOTH-GGUF-QUANTIZATION.md
- [COMMUNITY] Unsloth Dynamic GGUFs wiki (UD variant ladder) — https://github.com/1bit-systems/1bit-systems/blob/HEAD/docs/wiki/unsloth-dynamic-ggufs.md
- [COMMUNITY] First independent Dynamic v3.0-vs-v2.0 test (srmiles, 2026-08-27; KLD vs Q8_0, 9.5% slower decode, UD-Q3_K_XL llama-server bug) — https://github.com/srmiles/local-llm-benchmarks/blob/HEAD/models/tested/2026-08-27-unsloth-dynamic-3-vs-2.md
- [COMMUNITY] GGUF quantizations overview (Artefact2; KL/PPL per-type table) — https://gist.github.com/Artefact2/b5f810600771265fc1e39442288e8ec9
- [VENDOR] DeepSeek-V2.5 i1-GGUF card (mradermacher; IQ ladder, imatrix quants, i1-IQ4_XS 125.7GB) — https://huggingface.co/mradermacher/DeepSeek-V2.5-1210-i1-GGUF
- [VENDOR] vLLM-EXL3 plugin changelog (vcruz305, v0.4.2, 2026-09-09; dense dispatch fix, GB10 52.05 vs 52.22 tok/s) — https://github.com/vcruz305/vllm-exl3/blob/HEAD/CHANGELOG.md
- [COMMUNITY] ExLlamaV3 v1.4.8 pin intake (vllm-starter; CUDA-graph capture fixes, workspace hardening) — https://github.com/mykhailotamarin/vllm-starter/commit/08558fd0b59c8647681f873209f9c48a943afa5b
- [COMMUNITY] ExLlamaV3 v1.4.7→v1.4.9 pin intake (69 commits, COOP_AUTOTUNE_VERSION 3→4) — https://github.com/reederey87/glm53-flash-exl3-2x-dgx-spark/blob/HEAD/docs/06-improvement-plan.md
- [COMMUNITY] vLLM (NVFP4/FP8) vs ExLlamaV3 (EXL3) vs SGLang shootout — Qwen3.8-Flash-Next on RTX PRO 6000 (MarcoPizeta, v2 05/09/2026, v3 08/09/2026) — https://github.com/MarcoPizeta/flash-next-rtxpro6000-bench
- [COMMUNITY] Qwen3.8-Flash-Next EXL3 on 4× RTX 3090 via TabbyAPI (108 tok/s single-stream, 3× 260k sessions) — https://github.com/creslinux/exllamav3/blob/perf/ls-opt/CHANGELOG.md
- [COMMUNITY] NVFP4 vs MXFP4 anatomy (block-16/FP8-E4M3/FP32-tensor-scale vs block-32/E8M0; B200 ~8× FP16 claimed) — https://github.com/wanshuiyin/aris-in-ai-offer/blob/HEAD/docs/tutorials/quantization_tutorial_en.md
- [VENDOR] NVIDIA official NVFP4 guidance (3.5× vs FP16, 1.8× vs FP8, <1% loss) — http://build.nvidia.com/spark/nvfp4-quantization
- [VENDOR] NVFP4 quantization DGX Station playbook (TensorRT Model Optimizer, DeepSeek-R1-Distill-Llama-8B) — https://github.com/gumbiidigital/dgx-spark-playbooks/blob/HEAD/nvidia/station-nvfp4-quantization/README.md
- [COMMUNITY] NVFP4 vs AWQ decode shootout on 4× RTX PRO 6000 (rtx6kpro; AWQ 3519 vs NVFP4 3232 tok/s at C=128) — https://github.com/local-inference-lab/rtx6kpro/blob/HEAD/benchmarks/nvfp4-quantization-comparison.md
- [COMMUNITY] FP4 training paradox (kubesimplify blog; BF16 beat NVFP4 on DGX Spark nanochat recipe) — https://github.com/kubesimplify/website/blob/HEAD/content/blog/day-4-quantization-demystified-bf16-fp8-nvfp4-mxfp4-int4-gguf-and-why-it-all-matters.md
- [COMMUNITY] Consumer-Blackwell dtype matrix (club-3090; NVFP4 weights-only, nvfp4 KV datacenter-only) — https://github.com/noonghunna/club-3090/blob/HEAD/docs/DTYPE_MATRIX.md
- [VENDOR] vLLM docs: AutoAWQ deprecated → llm-compressor — https://docs.vllm.ai/en/latest/features/quantization/auto_awq/
- [VENDOR] AutoAWQ GitHub (deprecated notice, adopted by vLLM project) — https://github.com/atone/AutoAWQ
- [COMMUNITY] AWQ quantization practice on Qwen3.5 (Medium, Apr 2026; AutoAWQ no Qwen3.5 support, llm-compressor transformers pin, A100 OOM fix) — https://medium.com/@ishaafsalman/quantizing-qwen3-5-9b-to-int8-rtn-vs-awq-on-a-hybrid-vlm-architecture-51507bcad773
- [COMMUNITY] Mid-2026 vLLM vs SGLang format support comparison (AWQ "more commonly used" for vLLM deployments) — https://github.com/andysingal/llmops/blob/HEAD/vLLM.md
- [COMMUNITY] AWQ Marlin practitioner fix log (Sept 2026; explicit `--quantization awq_marlin` required) — https://github.com/peuqui/aifred-intelligence/blob/HEAD/docs/vllm/VLLM_FIX_SUMMARY.md
- [VENDOR] AutoRound algorithm (Intel; 30/32 vs GPTQ, 27/32 vs AWQ on lm-eval) — https://medium.com/intel-analytics-software/autoround-sota-weight-only-quantization-algorithm-for-llms-across-hardware-platforms-99fe6eac2861
- [COMMUNITY] AutoRound vs GPTQ (sergiiob, 2026-09; 91.4% vs 90.5% top-1 agreement, KL 0.277 vs 0.319) — https://github.com/sergiiob/intel-arc-pro-b70-inference-cookbook/blob/HEAD/docs/ornith15-35a3/AUTOROUND-VS-GPTQ.md
- [VENDOR] ModelOpt CHANGELOG (0.46.0, 2026-08-18; Triton fast path 34×, LSQ, NVFP4 4/6, MXFP8) — https://github.com/nvidia/model-optimizer/blob/HEAD/CHANGELOG.rst
- [VENDOR] nvidia-modelopt on PyPI (0.46.0, 2026-08-18) — https://pypi.org/project/nvidia-modelopt/
- [COMMUNITY] ModelOpt 0.45.0 deep dive (SMF Works; pruning+quant 2.6× compose, Nemotron-3-Nano-30B) — https://github.com/smfworks/aiclearinghouse-site/blob/HEAD/content/blog/2026-07-06-nvidia-model-optimizer-0-45-deep-dive.md
- [COMMUNITY] ModelOpt checkpoint formats in vLLM (FP8 variants, NVFP4, MXFP8; --linear-backend) — https://github.com/guqiong96/lvllm/blob/HEAD/docs/features/quantization/modelopt.md
- [VENDOR] optimum-quanto README (maintenance mode; HF redirects to bitsandbytes/torchAO) — https://github.com/huggingface/optimum-quanto/blob/HEAD/README.md
- [VENDOR] Hugging Face quantization overview (GPTQModel vs AutoGPTQ, bitsandbytes multi-backend, torchao Metal) — https://huggingface.co/docs/transformers/v4.49.0/quantization/overview
- [COMMUNITY] vLLM v0.26.0 capability gates (Marlin SM75 floor, Machete SM90-exact, Conch SM80+, CUTLASS FP8 ≥89) — https://github.com/adargit008/mcgyvr/blob/HEAD/okf/must-read/touching-engine.md
- [COMMUNITY] vLLM quant deployment rules (MXFP4 Hopper+ only, NVFP4 W4A16 via Marlin on Ampere, AWQ --dtype float16) — https://github.com/billyhargroveofficial/billy-skills-hub/blob/HEAD/cloud-gpu-vllm-uncensored/SKILL.md
- [COMMUNITY] vLLM backend selection (Isambard, Aug 2026; quant→backend mapping, W4A8 humming, NVFP4 flashinfer_cutedsl) — https://github.com/ai4ci/ivllm/blob/HEAD/design/references/vllm-backend-selection.md
- [VENDOR] HQQ (mobiusml; data-free, 1/2/3/4/8-bit, HQQ+, PEFT) — https://github.com/mobiusml/hqq/blob/HEAD/Readme.md
- [VENDOR] gemlite (mobiusml; Triton kernels, MXFP4/NVFP4 dynamic, 7–8× prefill, 3–6× decode) — https://deepwiki.com/mobiusml/gemlite
- [COMMUNITY] HQQ quantization benchmarks (ai-atlasforge; A100 numbers, INT8 KV for 8K+ ctx) — https://github.com/dragonshadows1978/ai-atlasforge/blob/HEAD/HQQ_QUANTIZATION_BENCHMARKS_RESEARCH.md
- [VENDOR] llama.cpp upstream digest (2026-08-21; v0.2.0 semver, Kimi K3, ROCm/Vulkan, DSpark, --mmproj-device) — https://github.com/licjon/cl-llama-cpp/blob/HEAD/docs/upstream-digest.md
- [COMMUNITY] llama.cpp ROCm 6.4.4 (dasroot; Qwen3-235B-A22B 1132 tok/s on Radeon 8060S, 3.47× vs ROCm 7.0.1) — https://dasroot.net/posts/2026/01/running-local-llms-python-ollama-llama-cpp-transformers/
- [COMMUNITY] Production quantization wiki (epyc-root; Q4_K_M production standard, REAP+AutoRound GLM-4.7 700→92GB) — https://github.com/pestopoppa/epyc-root/blob/HEAD/wiki/quantization.md
- [COMMUNITY] Quantization inference guide 2026 (GPTQ/AWQ derivations, INT3/INT2 pitfalls) — https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Inference_Optimization/quantization_inference.md
- [COMMUNITY] LLM quantization guide 2026 (ayinedjimi; 4-bit MMLU 2–4 pts, AWQ vs GPTQ perplexity) — https://ayinedjimi-consultants.fr/static/pdf/quantization-llm-2026-gguf-gptq.pdf
- [VENDOR] Simon Willison — ggml.ai joins Hugging Face (2026-02-20) — https://simonwillison.net/2026/Feb/20/ggmlai-joins-hugging-face/
- [VENDOR] GitHub — ggml.ai joins Hugging Face, llama.cpp Discussion #19759 — https://github.com/ggml-org/llama.cpp/discussions/19759
- [COMMUNITY] InsiderLLM — what the Hugging Face acquisition means — https://insiderllm.com/guides/llamacpp-hugging-face-ggml-acquisition/
- [UNVERIFIED] NVIDIA offered $12.9B to acquire Hugging Face (single-source secondary outlet) — no corroboration; claim treated as unverified
- [VENDOR] KDnuggets — Quantization and pruning methods to make your LLM leaner — https://www.kdnuggets.com/quantization-and-pruning-methods-to-make-your-llm-leaner

