---
id: ai-industry-kb-2026/08-quantization-model-formats/niche-formats-nf4-qtip
title: "Niche formats: NF4, QTIP"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "Alibaba", "ByteDance", "China", "DeepSeek", "Hugging Face", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-08-21"]
keywords: ["amd", "attention", "awq", "benchmarks", "compute", "consumer", "cost", "decode", "deepseek", "diffusion", "disaggregated", "distillation"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3943, 3997]
section: "8. Quantization & Model Formats"
sha256: 53d90ae12572d9d70bc686e0df013c1b0d7f28690c2602e5d5c26e3e1af3ea7e
---

# Niche formats: NF4, QTIP

### Niche formats: NF4, QTIP

- **NF4** (NormalFloat4): 4-bit index into a fixed 16-entry codebook of normal-distribution quantiles, plus a per-block absmax scale. Introduced with **QLoRA** (Dettmers et al., 2023) for fine-tuning; remains the highest-value pure-software 4-bit path with no hardware dependency (bitsandbytes). Dequant kernel: `codebook[nibble] * blockScale`.
- **QTIP** (Cornell RelaxML): EXL3 is its streamlined variant (one-step quantization, fused Viterbi kernel). EXL3 ~4 bits/weight formats for consumer NVIDIA GPUs; source-build on Windows.

### vLLM and GGUF: out-of-tree, experimental

- **vLLM's main branch contains zero GGUF code.** In-tree GGUF support was deprecated (tracked in RFC vllm-project/vllm#39583) and moved to the out-of-tree plugin **`vllm-project/vllm-gguf-plugin`** (`pip install vllm-gguf-plugin`). vLLM's own documentation describes GGUF support as **"highly experimental and under-optimized."**
- Loading: `vllm serve <hf_repo>:<quant_type>` or a local `.gguf` path; `--tokenizer <base-model>` recommended because GGUF tokenizer conversion is slow/unstable; `--tensor-parallel-size` works.
- The plugin accepts K-quants (Q2_K–Q6_K), I-quants (IQ1_S–IQ4_XS), file-level types with no tensor type (e.g. `IQ2_M`, `MXFP4_MOE`), and dash-prefixed custom names such as **`UD-Q4_K_XL`** (Unsloth Dynamic names load through it).
- Execution uses Triton dequantization kernels (dequantize to FP16/BF16 before GEMM) plus llama.cpp-style CUDA vec-dot kernels with quantized (q8_1) activations. **Hardware: GPU-only** (CUDA or ROCm toolkit required). No CPU path.
- An experimental **`prism-ternary` branch** adds Prism Q1_0 / Q2_0 group-128 tensor layouts, packed ternary embeddings and output heads, CUDA dequantization and MMVQ kernels to serve Bonsai 1.7B/4B/27B GGUFs (tested on RTX 2070 sm_75 and RTX 5060 Ti sm_120).
- **SGLang** also loads GGUF via `--load-format gguf`, but NVIDIA-only.
- Practical guidance (Sept 2026): **reach for GGUF only with a specific reason** (e.g. you only have a GGUF checkpoint and cannot reconvert). GGUF's natural home is llama.cpp or Ollama. Choosing vLLM for a GGUF-only workload means fighting an experimental out-of-tree plugin.

### llama.cpp in 2026

- **Version 0.2.0**: major version bump with enforced semantic versioning via CMake and a `release.sh` script (upstream digest, 2026-08-21).
- **New architectures**: Kimi K3 (hybrid KDA linear/recurrent + MLA full attention, 1024 latent-MoE experts, "situ" activation), BailingMoE3 (ByteDance, speculative decoding + Q-LoRA), GraniteSWA/GraniteMoeSWA (IBM, sliding-window attention, per-layer RoPE/NoPE mixing), dots3-note (DSA-iSWA dynamic sparse attention KV cache).
- **Backends**: ROCm 6.4.4 and Vulkan matured — Qwen3-235B-A22B hits **1,132 tok/s on AMD Radeon 8060S with ROCm 6.4.4 (3.47× over ROCm 7.0.1)** [COMMUNITY]; `ggml_rope_set_offset()` now spans CPU, Metal, CUDA, Vulkan, OpenCL, SYCL, WebGPU, Hexagon.
- **Speculative decoding**: DSpark speculators accept SpecForge-exported drafts; `--models-dir` auto-discovers MTP assistant models; owned draft KVarN caches for EAGLE3; audited Qwen MTP; DFlash1/2.
- **Multimodal**: `--mmproj-device` places the vision projector on a chosen GPU; mtmd audio/TTS bindings in llama-cpp-python 0.3.47.
- llama.cpp remains the only runtime spanning CPU, CUDA, Metal, HIP, Vulkan and SYCL for GGUF.

### Edge and ARM GGUF variants

- mradermacher's 2026 cards expose ARM-targeted Q4_0 layouts: `Q4_0_4_4` — "fast on arm, low quality"; `Q4_0_4_8` — "fast on arm+i8mm, low quality"; `Q4_0_8_8` — "fast on arm+sve, low quality". They trade the generic Q4_0 layout for ARM-specific vectorization (i8mm matrix-multiply, SVE) — same ~5.5–5.6 GB class as Q4_0 for a 9B model but measurably faster on phones and ARM laptops.

### MoE × quantization: why the combination wins

- Mixture-of-Experts models are the biggest beneficiaries of quantization: only a fraction of parameters are active per token (e.g. GLM-5.2: 40B active of 744B; MiniMax M3: ~23B of 428B), so compressing the **inactive** expert weights costs almost nothing in quality-per-FLOP while dividing memory by 4–8×.
- Production pattern (2026): **MXFP4/NVFP4 QAT applied to MoE experts** (DeepSeek-V4: MXFP4 QAT on experts + indexer QK path), FP8 for attention blocks, and router/gate layers kept unquantized for stability (`--moe` flag in llm-compressor pipelines).
- Verified footprints: **GLM-5.2** (744B total) at **UD-IQ2_M = 239 GB** disk → 256 GB unified-memory Mac; **DeepSeek-V4-Flash** (284B) at **NVFP4 = 172 GB** → B200/B300 native.
- Caveat: the **KV cache does not shrink with MoE sparsity** — it stays proportional to layer count and context length. A quantized 400B+ MoE still needs its full KV budget at long context, which is why FP8/4-bit KV and disaggregated prefill/decode matter as much as weight formats (see 08b).

### Quantization-aware training (QAT) at 4-bit

- Post-training quantization works well down to ~INT8/FP8; at INT4 the model never saw rounding error during training, so errors compound. **QAT** trains with simulated quantization (straight-through estimator), producing weights robust to their own rounding.
- Measured payoff at INT4: **PTQ quality drop 3–8% vs QAT <1%**. Cost: full retraining compute, and the model is committed to one quantization target.
- Production examples (2026): **Kimi K2.6 ships INT4 natively via QAT (~2× speedup)**; **DeepSeek-V4 applies FP4 (MXFP4) QAT to MoE experts plus the indexer QK path** — exactly the parameter-heavy parts.
- NVIDIA's QAD (quantization-aware distillation) datapoint: frozen BF16 teacher + KL divergence between token distributions, **no replay of SFT/RL stages** — Nemotron-3-Nano-30B NVFP4 reaches **99.4% of the BF16 baseline** on reasoning/coding benchmarks [VENDOR]. Important as a *method*: QAD sidesteps the SFT/RL replay cost of QAT.

### vLLM GGUF plugin: tested coverage and CI

- Brief-claim verification: "Experimental GGUF support via dedicated plugin (vllm-gguf-plugin); direct loading from Hugging Face; tensor parallel; base tokenizer recommended" — **VERIFIED on all four points**.
- Tested model coverage in plugin CI: Qwen 2.5 (Q6_K), Qwen 3 (Q8_0), Phi 3.5 (IQ4_XS), GPT-2 / StableLM (Q4_K_M), Gemma 3 (Q4_0), OLMoE (Q4_0), plus vision/diffusion (Gemma 3, Z-Image-Turbo, FLUX.2-klein). The README's tested quant list is Q6_K / Q8_0 / IQ4_XS / Q4_K_M / Q4_0.

### Practical quantization workflows

- **GGUF via llama.cpp**: (1) compute an importance matrix — `llama-imatrix` on representative calibration text; (2) quantize with it — `llama-quantize --imatrix imatrix.dat model-f16.gguf model-q4_k_m.gguf Q4_K_M`; (3) for 1–3-bit mixes the tool warns if no imatrix is supplied — do not skip it; (4) serve with `llama-server`/`llama-cli` (`-ngl 99` for full GPU offload, `--jinja` for the embedded chat template); warm up with dummy inferences because GGUF uses mmap (first inference page-faults; SSD/NVMe strongly preferred).
- **FP8 / AWQ / GPTQ via llm-compressor (vLLM stack)**: one-command pipeline (`vllm-quant-pipeline`): `bash run.sh Qwen/Qwen3-8B --method fp8 --verify` — FP8 needs no calibration; AWQ/GPTQ use a bundled mixed English/Chinese/code corpus. MoE-aware: `--moe` keeps router/gate layers unquantized for stability. Publish path: one command emits a Hugging Face repo with a generated model card (compressed-tensors safetensors).
- **Unsloth quants**: training/fine-tuning side — `model.save_pretrained_gguf("out.gguf", quantization_method="q4_k_m")` and Dynamic variants from the Unsloth package; consumption side — Dynamic quants are prebuilt `UD-` artifacts; the per-layer schedule is proprietary and not reproducible from the OSS tree.

### Deployment decision matrix (2026-09)

