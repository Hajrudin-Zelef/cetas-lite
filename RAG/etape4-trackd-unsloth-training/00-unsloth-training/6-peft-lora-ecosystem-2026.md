---
id: etape4-trackd-unsloth-training/00-unsloth-training/6-peft-lora-ecosystem-2026
title: "6. PEFT / LoRA ECOSYSTEM (2026)"
domain: step-4-track-d-unsloth-training-fine-tuning-tooling-2026
role: deep-dive
task: training
actors: ["AMD", "Alibaba", "Apple", "China", "DeepSeek", "Google", "Meta", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Samsung", "TensorRT-LLM", "Unsloth", "Z.ai"]
dates: ["2026-04-20", "2026-07-14", "2026-08-31", "2026-09-07"]
keywords: ["lora", "agentic", "attention", "awq", "benchmark", "benchmarks", "blackwell", "consumer", "decode", "deepseek", "disaggregated", "disaggregated serving"]
source: docs/RAG/etape4_trackD_unsloth_training.md
source_anchor: ""
source_lines: [120, 201]
section: "Step 4 — Track D: Unsloth + Training / Fine-Tuning Tooling (2026)"
sha256: 02390b56a9e5f2b6cdf30dde8bc91fc2f6ebd98d525838ee53432e74dcd7506f
---

# 6. PEFT / LoRA ECOSYSTEM (2026)

## 6. PEFT / LoRA ECOSYSTEM (2026)

- **DoRA** supported in Unsloth [secondary]; **LoRA/QLoRA** remain the default adapter method; **NVFP4 LoRA** (ScatterMoE W4A16 / SonicMoE W4A4) landed in Axolotl Jul 2026 [official/secondary].
- **LoRA directly on MoE expert weights** via ScatterMoE custom Triton kernels (Axolotl, Feb 2026) [secondary].
- **Liger Kernels** integrated in LLaMA-Factory (Qwen 3.5/3-Next) and **Liger 0.8.0** supported by TRL v1.4 GRPO [secondary].
- **NeMo Gym** (NVIDIA) integration in Axolotl (Apr 2026) [secondary].
- **Merged-adapter export**: Axolotl Jul 2026 supports lossless merge of NVFP4 MoE LoRA back into plain NVFP4 checkpoints [secondary].

---

## 7. DISTRIBUTED TRAINING FRAMEWORKS

### 7.1 DeepSpeed
- **v0.19.4 patch release** (~Aug 6, 2026) [official — GitHub release]: highlights — **AutoTP: ZeRO stage 3 inference with tensor parallelism**; HF `colwise_gather_output` for `lm_head` replace; AutoEP rank splits derived from per-expert count exchange; Tutel support for shared-MoE k≠1; DeepCompile compiler-state scoping; NVMe write warning [official].
- **Release 0.19.5** docs dated **Aug 06, 2026** (readthedocs PDF) [official]: covers AutoEP (automatic expert parallelism), MoE, Transformer kernels, pipeline parallelism, optimizers.
- **ZeRO-Inference refresh**: up to **20X throughput speedup** via weight quantization + KV-cache offloading (DeepSpeed ≥ 0.10.3) [official — deepspeedexamples README].
- AutoTP/AutoEP direction: automatic parallelism-plan derivation (HF `tp_plan` support) — reduces manual sharding config [official].

### 7.2 FSDP2
- FSDP2 is the 2026 default distributed strategy in the PyTorch ecosystem: Axolotl ships **FSDP2** for MoE expert quantization and distributed Muon pretraining; LLaMA-Factory v0.9.5 adds **FSDP2** support; torchtune builds on latest PyTorch distributed APIs; Axolotl added **PyTorch 2.13** support (v0.19.0) [official/secondary].
- ⚠️ No standalone "FSDP2 2026 announcement" located; evidence is via downstream adoption — treat specific claims per-project.

### 7.3 Megatron-LM
- **No 2026 Megatron-LM release notes located** in this window. Context: NVIDIA's Megatron-LM repo development continued (TransformerEngine, Megatron-Bridge), and DeepSpeed docs still reference Megatron-LM model init patterns, but no dated 2026 milestone was verified. → **[gap — flag for re-research]**.

---

## 8. INFERENCE EXTRAS

### 8.1 TensorRT-LLM
- **v1.2.0** (~Mar 2026) [official]: beta support for K-EXAONE, Nemotron Nano V3, Qwen3-Next, Qwen3-VL; expanded Blackwell/Hopper/Ampere enablement (B300/GB200/GB300, SM120/SM121/SM103); FP8/NVFP4/MXFP4/INT4-AWQ; MTP>1 speculative decoding for DeepSeek v3.2; disaggregated serving (service discovery, request cancellation, NIXL-LibFabric, Mooncake); FlashInfer batched sampling; EPLB expert parallelism; **DGX Spark (GB10) beta support**.
- **v1.2.1** (2026-04-20): multi-arch container incl. **arm64 on NGC** (`nvcr.io/nvidia/tensorrt-llm/release:1.2.1`) [secondary].
- **v1.3.0 RC** (Sep 2026): "Enable FlashInfer GDN decoding kernel for Qwen3.5"; Mamba-hybrid support (from 0.19.0-era work); Qwen3.6-27B (GDN hybrid) **not yet** in the stable support matrix [secondary].
- Infra: TensorRT 10.10.0, CUDA 12.9.0, PyTorch 2.7.0 wheels, ModelOpt 0.29.0; XQA kernels open-sourced; **V100 support removed** (breaking change) [official/secondary].
- Notable fixes in window: LLaMA 4 CUDA-graph illegal memory access; Qwen3-MoE attention DP bug [official].

### 8.2 ONNX Runtime GenAI
- **No 2026 ONNX Runtime GenAI release notes located** in this window → **[gap — flag for re-research]**. (Last verified releases predate the window; Microsoft's focus in 2026 reporting has been on Foundry/TRT-LLM-class stacks.)

### 8.3 llama.cpp server (llama-server)
- No dedicated 2026 llama-server milestone captured in this window → **[gap]**. (llama.cpp itself continues active development; server-side OpenAI-compatible endpoint is stable. Not separately verified — flag.)

### 8.4 ExLlamaV3 + TabbyAPI
- **ExLlamaV3** (`turboderp-org/exllamav3`) is the live project; **ExLlamaV2 is archived** ("Development continues on ExLlamaV3") [official].
- **v1.4.5 (2026-08-31)**: win_amd64 wheels cp310–cp313, cu128/torch 2.10.0; **v1.4.6** in use by Sep 3, 2026 (community benchmark) [independent].
- Format: **EXL3** — streamlined variant of **QTIP** (Cornell RelaxML); one-step quantization (fused Viterbi kernel), minutes for small models on one RTX 4090; coherent at **1.6 bpw** (Llama-3.1-70B in <16 GB); Marlin-inspired GEMM ≈ memory-bound at 4 bpw on RTX 4090; v1.0.0 (2026-07-14) greatly improved Ampere GEMM/GEMV, removed flash-attn-2/xformers deps [independent — community doc citing GitHub API].
- Features: tensor-parallel + expert-parallel for consumer setups; 2–8-bit KV-cache quantization; speculative decoding; LoRA; multimodal; ROCm on the to-do list (not yet) [independent].
- Model support (Sep 2026): Llama 3/4, Qwen 2/2.5/3 (+MoE, Next, VL, 3.5), Mistral/Mixtral, Phi3/4, Gemma, GLM-5.3-Flash, DeepSeek family, Kimi, NanoChat, Nemotron-H/3, Olmo 3.1, Step 3.5/3.7, SmolLM, SolarOpen, Seed-OSS [independent].
- **TabbyAPI** (`theroyallab/tabbyAPI`, ~1.3k★, AGPL-3.0): active; security issue #448 (CORS allowlist) **closed 2026-09-07** by maintainer; startup script manages ExLlamaV3 prerequisites; OpenAI-compatible serving [independent].
- Community benchmark (Sep 3, 2026, 4× CMP 170HX, GLM-5.3-Flash EXL3 4.05bpw via TabbyAPI): **25.2 tok/s single-concurrency decode**, 44.6 tok/s aggregate at c=8 [independent — measured].

---

## 9. FUNDING / M&A IN THE TOOLING SPACE

- **Unsloth AI**: ~$500K pre-seed (YC, Redpoint scout, Samsung NEXT, Pioneer Fund); 12 investors incl. Lightspeed VP, Transpose, Brightwing per PitchBook. No confirmed larger round as of Sep 22, 2026 [secondary; see §1.1 caveats].
- **Axolotl AI Cloud**: no funding news located in window [gap].
- **Thinking Machines' Tinker** (managed fine-tuning API, Oct 2025 launch; Inkling model Jul 2026) is the managed-cloud alternative positioning against Unsloth Core/Axolotl self-hosted — relevant competitive context, covered in Step 3 Track C [cross-ref].
- **No 2026 M&A** found in the fine-tuning/inference-tooling space (Unsloth, Axolotl, LLaMA-Factory, TRL, torchtune, DeepSpeed, ExLlamaV3 all independent as of Sep 22, 2026) [secondary — absence of evidence; not proof].

---

## 10. UNCERTAINTY / VERIFICATION LOG

1. Unsloth speed/VRAM claims (2–5x, 70% less, 12x MoE, 7x long-ctx RL) are **vendor-only** — no independent reproduction located. [open]
2. Unsloth raise beyond ~$500K pre-seed — unconfirmed; Crunchbase/PitchBook details paywalled. [open]
3. Unsloth "official fine-tuning partner" of Google/OpenAI/Meta/NVIDIA — no official announcements located. [unverified]
4. Unsloth Apple-Silicon/MLX native support — conflicting community claims. [unverified]
5. Axolotl funding/company financials — no data. [gap]
6. torchtune 2026 releases — not located. [gap]
7. Megatron-LM 2026 milestones — not located. [gap]
8. ONNX Runtime GenAI 2026 releases — not located. [gap]
9. llama-server 2026 milestones — not located. [gap]
10. TRL version ~v1.8 claim — community doc only; latest directly confirmed is v1.4 (Sep 2026). [secondary]
11. LLaMA-Factory specifics — sourced from auto-generated AI China News posts only. [secondary]
12. Unsloth Desktop feature list (MCP, Deep Research Mode, etc.) — community skill cards; not independently confirmed. [secondary]

## 11. COLLECTION METADATA
- Method: web search (browser.search), 8 query batches, Sep 22, 2026.
- Priority sources hit: official GitHub release pages (Axolotl v0.19.0, DeepSpeed v0.19.4, TensorRT-LLM v1.2.0), official docs (TRL CLI, torchtune PyPI, Unsloth docs), HF blog (Aug 2026 agentic RL), community benchmarks (ExLlamaV3/TabbyAPI Sep 2026).
- No SiliconANGLE piece specific to this tooling niche was located in-window; inference-tooling coverage came from GitHub/primary sources instead.
- Nothing was sent externally; read-only research + local file write.
