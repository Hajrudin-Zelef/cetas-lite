---
id: ai-industry-kb-2026/06-inference-engines/mla-and-attention-support-on-vllm-2026
title: "MLA and attention support on vLLM (2026)"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "AWS", "Alibaba", "DeepSeek", "Hugging Face", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-06", "2026-07-28", "2026-08", "2026-09-22"]
keywords: ["attention", "vllm", "amd", "benchmark", "blackwell", "datacenter", "decode", "deepseek", "diffusion", "gguf", "glm", "gpu"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2264, 2295]
section: "6. Inference Engines"
sha256: 57149f1c7f81e213a7ffc66740b3df3ba285cf1ca11e1313a840147674a740ab
---

# MLA and attention support on vLLM (2026)

- **Mechanism:** generate multiple candidate tokens per KV-cache load via a small drafter model (EAGLE, NEXTN, DSpark, DFlash), then verify with the large model.
- **v0.25.0:** **universal speculative decoding for heterogeneous vocabularies (TLI, #38174)** — drafter and target can use different vocabularies; plus DSpark (#46995) and DFlash (#46770, #46853) drafters; dynamic speculative decoding compatible with full CUDA graphs.
- **v0.28.0:** **DFlash2** with local convolution and a candidate selector; **DSpark confidence-scheduled verification**; async scheduling auto-enabled for draft models; adaptive speculative token budget delivering ~60% better DSpark TTFT [VENDOR] (Kimi-K3 push).
- **v0.20.0 → v0.28.0 lineage:** MTP (multi-token prediction) supported on newer models (Qwen3.5/3.6/3.8, DeepSeek V4); reasoning-effort prompts and mappings for DeepSeek V4; sparse top-k metadata kernel optimizations.
- **Neuron path:** speculative decoding support (Eagle V1) documented on Trainium/Inferentia; **TPU path:** Eagle3/Ngram speculative decoding supported in `tpu-inference` [DIRECTIONAL — code matrices, not benchmarked].

### MLA and attention support on vLLM (2026)

- **v0.20.0 (June 2026):** **FlashAttention 4 is the default MLA prefill backend** (SM90+, head-dim 512, paged-KV); DeepSeek MLA supported on SM90+.
- **v0.28.0 (August 2026):** **sparse MLA works end-to-end for DeepSeek V4** (plain decode, MTP, DSpark speculative decoding); the GLM-5.3-Flash NVFP4 1M-context DGX Spark deployment required extending vLLM's SM90 NoPE sparse-MLA backend to SM121 [COMMUNITY patch — evidence of day-0 MLA-class model support with per-architecture tuning]; ROCm enablement on gfx11 and gfx950.
- **Hardware reality check:** MLA kernels are SM90+-first (Hopper); SM120 (Blackwell workstation) support arrived through 2026 via FlashInfer with community patching (NaN bugs in specific batch ranges on SM121 bisected and fixed in nightlies); datacenter Hopper/Blackwell mature, workstation silicon still being tuned [COMMUNITY].

### MLA "4x batch size per GPU" — claim audit

- **Origin:** DeepSeek-V2's MLA paper reported **93.3% KV-cache reduction** and **5.76x throughput** vs standard MHA. A ~4x batch-size increase at fixed HBM is the natural arithmetic consequence: if per-token KV shrinks ~15x, batch capacity (KV-bound in decode) grows several-fold.
- **[UNVERIFIED]:** direction and mechanism confirmed by the paper figures and both engines' MLA-native paths, but **no independent framework-level benchmark was found that measures "4x batch size per GPU" as a stated, reproduced result** — published figures are 5.76x throughput (paper) and 1.9x decode (SGLang DP attention), different metrics. Present "4x batch" as a derived planning figure, not a measured engine benchmark.
- **What would settle it:** a controlled measurement — same GPU class (e.g. 8×H100), same MLA model, same precision, engine-only variable (vLLM MLA vs a non-MLA baseline serving path) — measuring **maximum sustainable concurrent batch before OOM or SLO breach**, not throughput. Until published, the honest formulation is: *MLA's 93.3% KV reduction removes the decode batch ceiling by roughly an order of magnitude in the KV-bound regime; operators report ~4x as a conservative planning figure.*

### The GGUF plugin — status as of 2026-09-22

- Repository `vllm-project/vllm-gguf-plugin`; install via `pip install vllm-gguf-plugin` or from git. Loading: `vllm serve <hf_repo>:<quant_type>` for direct Hugging Face loading, or a local `.gguf` path. The HF repository `unsloth/Qwen3-0.6B-GGUF` exists (quants including Q4_K_M listed); the `:Q4_K_M` addressing suffix is the plugin's documented form. [UNVERIFIED: no live execution of the exact example command performed in this research — repo, quant, and syntax each verified separately.]
- `--tokenizer <base-model>` recommended (GGUF tokenizer conversion slow/unstable); `--hf-config-path` fallback for architectures Hugging Face cannot convert. `--tensor-parallel-size` works with the plugin.
- Accepted quant types: GGML tensor types, file-level types with no tensor type (`IQ2_M`, `MXFP4_MOE`), extended suffixes (`Q4_K_M` → `Q4_K`), dash-prefixed custom names (`UD-Q4_K_XL` — the Unsloth Dynamic naming).
- Execution kernels: Triton dequant kernels (standard `Q4_0`–`Q8_1`, K-quants `Q2_K`–`Q6_K`, IQ quants `IQ1_S`–`IQ4_XS`) dequantizing to FP16/BF16 before GEMM, plus llama.cpp-style CUDA vec-dot kernels with q8_1 activations (`csrc/gguf/gguf_kernel.cu`, `mmvq.cuh`, `mmq.cuh`).
- Hardware: **GPU-only** — CUDA or ROCm toolkit required; the plugin builds a HIP/CUDA kernel (`_C_gguf`) against the exact installed PyTorch (`--no-build-isolation` matters). x86/Arm CPU not supported. One documented ROCm build compiled ~12 GPU architectures by default (gfx90a, gfx942, gfx950, gfx1100–1103, gfx1150–1153, gfx1200, gfx1201), taking 15–20 minutes; `PYTORCH_ROCM_ARCH` (e.g. `gfx1201`) cuts this to minutes. On macOS the plugin does not install (builds a CUDA extension, imports Triton) — `vllm-project/vllm-metal` re-implements the layer locally.
- Model coverage: generic GGUF weight adapter, not an allowlist — "but not every vLLM-supported architecture works". Plugin CI covers Qwen 2.5 (Q6_K), Qwen 3 (Q8_0), Phi 3.5 (IQ4_XS), GPT-2 (Q4_K_M), StableLM (Q4_K_M), Gemma 3 (Q4_0), OLMoE (Q4_0), plus VLM/diffusion (Gemma 3, Z-Image-Turbo, FLUX.2-klein).
- **Maturity caveats (2026-09-22):** young enough that per-model-family bugs are expected. Documented gap: **no Gated-DeltaNet hybrid support as of 2026-07-28** (tracking issue #80 — Qwen3.5/3.6 MoE architectures; without a registered quant method, `FusedMoE` dequantizes all experts to bf16 — ~160 GB needed vs 121 GB available on one reported setup). Practitioner guidance: serve safetensors if reliability matters today.
- AMD/RDNA community forks (e.g. `vllm_for_amd` guides) explicitly exclude GGUF ("prefer native Safetensors checkpoints") — the plugin path targets datacenter/ROCm.
- llama.cpp remains the GGUF reference runtime and the format authority — "does it run in llama.cpp?" is the first diagnostic when the plugin hits a model-family bug.

### vLLM-Omni — dated facts

