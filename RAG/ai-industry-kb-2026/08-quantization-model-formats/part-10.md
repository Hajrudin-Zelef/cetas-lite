---
id: ai-industry-kb-2026/08-quantization-model-formats/part-10
title: "8. Quantization & Model Formats (part 10)"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Intel", "Nvidia", "SGLang", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-03-07", "2026-08-18", "2026-08-19", "2026-08-20", "2026-08-27", "2026-08-29", "2026-08-31", "2026-09-05", "2026-09-08", "2026-09-09", "2026-09-18"]
keywords: ["quantization", "awq", "benchmarks", "blackwell", "consumer", "datacenter", "decode", "deepseek", "fp4", "fp8", "gguf", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4153, 4195]
section: "8. Quantization & Model Formats"
sha256: 5d2913427baf12b32ec40ca9e48d8a0bcb9bf101d8979323740b0801ed878d79
---

# 8. Quantization & Model Formats (part 10)

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
