---
id: ai-industry-kb-2026/10-training-fine-tuning/august-2026-adoption-signal-and-versioning-shift
title: "August 2026: adoption signal and versioning shift"
domain: training-fine-tuning
role: deep-dive
task: training
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Hugging Face", "Meta", "OpenAI", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-07", "2026-07-22", "2026-08", "2026-08-12", "2026-08-19", "2026-08-28", "2026-08-30", "2026-09", "2026-09-02", "2026-09-09", "2026-09-17", "2026-09-22"]
keywords: ["amd", "attribution", "benchmark", "deepseek", "diffusion", "dpo", "fine-tuning", "fp8", "gguf", "glm", "gpu", "inference"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5427, 5473]
section: "10. Training & Fine-Tuning"
sha256: 65f706a16f00e3a8f7d79187798300889aed6ddec54cc5e270295d9fea50a79b
---

# August 2026: adoption signal and versioning shift

- **2026-07-22**: independent 4-framework comparison (MarkTechPost: Unsloth vs Axolotl vs TRL vs Llama-Factory) attributes Unsloth's MoE advantage to the **split-LoRA formulation** that avoids materializing LoRA deltas for all experts before the MoE matmul. Reported figures: **gpt-oss-20b fine-tunable within 12.8 GB** (secondary-reported); **Qwen3-30B-A3B 16-bit LoRA at 63 GB**; on a **B200** with gpt-oss at 8K context, Unsloth 47.43 GB vs Transformers v5 73.80 GB; at 16K context, Unsloth 55.13 GB while Transformers v5 OOMs.
- **July 2026**: independent practitioner analysis (Chew Loong Nian, Medium) prices every byte of a LoRA fine-tuning step from first principles and independently validates the memory-mechanism attribution. For Llama 3.1 8B, **87.3% of the memory that scales with sequence length is the cross-entropy loss head** — "one tensor that exists for a few microseconds, produces a single scalar, and never appears in your training script." For gpt-oss-20b it is 95.2%; at 16K context that one tensor wants **49.1 GB** while the 4-bit model fits in 12.8 GB — the loss function is ~4x the size of the thing being trained.
- The same author's 120-line calculator **predicts the Transformers+FA2 baseline at 27,713 tokens on an 80 GB GPU**, close to Unsloth's published 28,454 — an independent check that Unsloth's published context-length tables are in the right ballpark. His conclusion: "The speed numbers are real but modest and model-dependent. The memory numbers are enormous and structural, and they come almost entirely from what each framework does with this one tensor" — i.e., cut/fused cross-entropy is the actual axis of differentiation between Unsloth, Axolotl, TRL, and LLaMA-Factory. [INDEPENDENT ANALYTICAL, SECONDARY-REPORTED]
- **July 2026 (v0.1.481-beta, "DeepSeek-V4 + NVFP4 Exporting")**: Unsloth Studio can export to **NVFP4, FP8, and imatrix GGUFs** after training; DeepSeek-V4-Flash supported with NVFP4 export. The same release notes claim **GRPO 1.3x faster** and **MoE training 3–5x faster** in Unsloth core [VENDOR] — an intermediate datapoint between the February "12x" headline and the measured September reality.
- **No new independent benchmark** of Unsloth fine-tuning speed or VRAM vs baselines (HF/TRL/Axolotl/torchtune) published between the July 22 MarkTechPost comparison and 2026-09-22 was found.

### August 2026: adoption signal and versioning shift

- **2026-08-12 → 2026-08-30**: Unsloth GitHub stars grew 70.4K → 72.2K → 72.9K → 73.4K → 75.2K (~400–570 stars/day), supporting characterization as the community default for local fine-tuning (not, by itself, enterprise or multi-node dominance).
- **Dynamic v3.0 launched 2026-08-19**, superseding Dynamic Quantization 2.0 (which was current as of Daniel Han's July 2026 talks). Format mechanics of the UD quant family are covered in §8 (cross-ref; not duplicated here).
- **2026-08-28/29**: GLM-5.3 weights shipped, superseding GLM-5.2 — the model Unsloth's early-September MTP release would target (see below).

### September 2026: Studio beta releases v0.1.805 → v0.1.808

GitHub release pages (indexed early-to-mid September 2026; the vendor does not stamp exact dates on these beta pages, so dates below are approximate) document two thematic pairs:

**v0.1.805-beta / v0.1.806-beta — "2x Faster Qwen3.8-Flash + GLM-5.3-Flash MTP"**
- **Multi-token prediction (MTP) enabled by default** for Qwen3.8-Flash-Next and GLM-5.3-Flash, delivering **up to 2x faster generation** [VENDOR]; disable-able. Ships with downloadable Unsloth GGUFs (`unsloth/Qwen3.8-Flash-Next`, `unsloth/GLM-5.3-Flash`) and vendor guides.
- **170+ training, chat, hardware and performance improvements** in the release (per vendor count).
- **MLX / Apple Silicon training gains**: "Fine-tune both large MoE models with text or images on Apple Silicon using MLX"; long Qwen chats up to **30x faster follow-up turns on Mac** [VENDOR]; MLX models use full context size with longer batched generation; MLX models servable through Unsloth's OpenAI-compatible API.
- **Improved multi-GPU planning, memory fitting, and split-model training** — adjacent to the "modest GPU" claims, but documented as planning/fitting features, not benchmarked speedups.

**v0.1.807-beta / v0.1.808-beta — "Large Performance Gains + Fixes"**
- **AMD by default uses Vulkan instead of ROCm: +20% prefill, +23% faster prompt processing, +8% faster generation on Strix Halo** [VENDOR]; Strix iGPU BIOS guidance claims 3x faster inference with more iGPU VRAM; gibberish issues on Strix/iGPUs fixed (reported upstream to AMD).
- **INT8/FP8 diffusion 1.2x–1.7x faster** (all models accelerated) — this **revises downward** the official Sept-17 changelog's "2x faster" line. The consolidation should use 1.2–1.7x (the later, more precise figure) and treat the earlier 2x as superseded.
- **Over 250 bug fixes, 60% smaller binaries**, 2x faster updates, signed Windows llama.cpp binaries (reduces Smart App Control false positives).
- **Apple Silicon**: gated-delta models train up to **25% faster** [VENDOR]; quantized MLX KV caches use up to **74% less prompt memory** [VENDOR]; **DoRA fine-tuning and more DPO loss types now available on Apple Silicon**; more multimodal models trainable from text-only datasets.
- PyTorch updated to 2.11 (from 2.10); 2.14 "will be soon" — consistent with the PyTorch 2.14 release (2026-09-02), whose silent clamp/min/max boundary-subgradient change (1→0) is a genuine training-reproducibility hazard for anyone upgrading mid-experiment.

**Backend versioning scheme**: a 2026-09-02 PR bumps `unsloth/_version.py` to **`2026.9.2`**, stamped on the v0.1.806-beta desktop release (PyPI files are immutable, so each desktop release needs a new backend stamp). Cadence: `2026.6.9` (June) → `2026.7.2` (July) → `2026.8.3` (August) → `2026.9.1/2026.9.2` (September). Install pins moved in lockstep (`unsloth>=2026.5.8` → `>=2026.8.3` across the summer betas).

**2026-09-17**: official changelog entry "Docker + MultiUser + AMD Support" — multi-user Docker images, AMD RDNA1/2, INT8/FP8 diffusion inference, ARM64 Windows CUDA, GRPO with Qwen3.5 + latest TRL/vLLM. **As of 2026-09-22, no newer releases were found** — the official changelog's latest entry remains 2026-09-17 and the newest beta tags remain v0.1.808-beta.

### MoE training-kernel correctness fix (2026-09-09)

A 2026-09-09 commit in `danielhanchen/unsloth-zoo-staging` ("Fix down-LoRA row order in the Triton grouped-GEMM MoE forward", PR #887) is the only **new Unsloth MoE-kernel change** found between the Wave 2 research cutoff and 2026-09-22:

- **The bug**: in the Triton grouped-GEMM MoE forward, the second grouped GEMM runs with `permute_y=True`, so its output returns to token order while the down-LoRA delta stays expert-sorted. The delta was therefore **added to the wrong tokens and gradients were silently wrong** — a correctness (not performance) issue.
- **The fix**: scatter the delta through `gather_indices` so rows land on the tokens they belong to.
- **Companion hardening** (moved to its own PR): dtype-aware MoE backend selection — `torch._grouped_mm` is bf16-only, so auto-selection keys on forward activation dtype (`grouped_mm` only for bf16, otherwise falling through to `unsloth_triton` / the native loop), with an explicit `UNSLOTH_MOE_BACKEND=grouped_mm` override.

For the consolidation: this is evidence that Unsloth's MoE path was still under active correctness maintenance in September 2026, and it further qualifies any undated "MoE is production-solid" reading.

### GGUF export after training: one-call API (verified ~2026-09-17)

The "direct GGUF/Ollama export" claim is real and documented as a first-class, post-training API — not a workaround. Verified against the Hugging Face TRL Unsloth-integration docs and community skill documents:

