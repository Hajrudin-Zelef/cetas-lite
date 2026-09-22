---
id: ai-industry-kb-2026/10-training-fine-tuning/overview
title: "10. Training & Fine-Tuning"
domain: training-fine-tuning
role: deep-dive
task: training
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Google", "Hugging Face", "Meta", "Nvidia", "OpenAI", "Samsung", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-02", "2026-02-09", "2026-03", "2026-03-23", "2026-04", "2026-05", "2026-05-06", "2026-07", "2026-07-02", "2026-07-22", "2026-08", "2026-08-12", "2026-08-19", "2026-08-28", "2026-08-30", "2026-09", "2026-09-02", "2026-09-09", "2026-09-15", "2026-09-17", "2026-09-22"]
keywords: ["fine-tuning", "training", "amd", "apache", "attention", "attribution", "awq", "benchmark", "benchmarks", "compute", "consumer", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5387, 5528]
section: "10. Training & Fine-Tuning"
sha256: 5e10330caced44d303edfcc1a8ff993ec378ac2f41840802b79277ffe5a0e5f4
---

# 10. Training & Fine-Tuning
Keywords: unsloth, fine-tuning, qlora, moe training, split-lora, cut cross-entropy, triton kernels, vram reduction, single-gpu training, llama 3.3 70b, rtx 4090, benchmark methodology, unsloth studio, save_pretrained_gguf, nvfp4 export, peft, esft, dora, fsdp2, axolotl, torchtune, deepspeed, liger kernel, grpo, dynamic v3.0, qwen3.8-27b, loss head, moe kernel fix, grouped gemm

## Summary

This section covers Unsloth's training and fine-tuning story from February to 22 September 2026, plus the 2026 PEFT and distributed-training landscape in which it sits. Unsloth is the community default for 1–2-GPU PEFT and local fine-tuning workflows; its differentiation is memory-first (Triton kernels, split-LoRA for MoE, cut/fused cross-entropy, padding-free packing), not an absolute speed crown.

- **The "12x faster MoE training" headline** (February 2026 release) is [VENDOR]-reported and hardware-specific: Unsloth's own MoE post attributes the top multipliers (up to ~7–12x) to B200-class hardware. No independent reproduction of the 12x figure was found. The auditable component-level evidence comes from the May 6, 2026 Unsloth×NVIDIA collaboration blog (packed-sequence metadata caching, double-buffered checkpoint reload, bincount MoE routing).
- **Memory savings come from six mechanisms**: hand-written Triton kernels with exact backpropagation ("0% accuracy degradation" claim); split-LoRA for MoE (avoids materializing LoRA deltas for all experts before the MoE matmul); fused/cut cross-entropy (avoids materializing the full logit tensor); custom gradient checkpointing; padding-free packing; pre-quantized 4-bit loading that cuts fragmentation.
- **Claim verdicts**: 70–80% VRAM reduction is [PARTIALLY VERIFIED, WORKLOAD-SPECIFIC] (~70% general marketing; >75% only for Llama 3.3 70B QLoRA in a specific vendor config; ~80% is GRPO-specific). Single-GPU 70B fine-tuning is [VERIFIED FOR QLORA] (Llama 3.3 70B on one 80 GB GPU) — not full-parameter. 70B on a single 24 GB RTX 4090/5090 remains [UNVERIFIED]; RTX 4090/5090 *clusters* as a 70B+ MoE training claim is [UNVERIFIED].
- **September 2026** brought four Studio beta releases (v0.1.805→v0.1.808): MTP-by-default for Qwen3.8-Flash and GLM-5.3-Flash, AMD Vulkan-by-default, MLX/Apple-Silicon training gains, and a downward revision of the diffusion speedup (Sept-17 changelog "2x" → v0.1.808-beta "1.2–1.7x"). The GGUF post-training export API (`save_pretrained_gguf`, one call, 23-entry quant list, OOM guard) was verified against TRL integration docs updated ~2026-09-17; NVFP4 export has existed in Studio since July 2026. A mid-September 2026 commit fixed a silent correctness bug in the MoE Triton grouped-GEMM path (down-LoRA deltas added to the wrong tokens).
- **Corrections recorded**: Unsloth's presence at AI Engineer World's Fair 2026 with a 112-slide deck is [UNVERIFIED]; the ~$40K seed / YC / Logan Kilpatrick investor claim is [UNVERIFIED] (a third-party figure says ~$500K via Redpoint Ventures scout + Samsung NEXT — both kept flagged); "Dynamic Quantization 2.0" was current in July 2026 but is SUPERSEDED by Dynamic v3.0 (2026-08-19).
- **The 2026 framework split**: Unsloth (single/dual-GPU PEFT, lowest memory) vs Axolotl (YAML multi-GPU, FSDP2/EP, DeepEP expert parallelism, NVFP4 MoE LoRA) vs torchtune (native-PyTorch hackable research, DTensor, torch.compile throughput lead at most tested sizes) vs DeepSpeed (max scale). FSDP2 is the 2026 default distributed substrate.
- **Benchmark methodology**: Unsloth figures are vs Hugging Face + FlashAttention 2 QLoRA — not vs torchtune with torch.compile (which leads on throughput) and not workload-universal. The independent torchtune paper (May 2026) and a July 2026 independent loss-head analysis confirm Unsloth's VRAM leadership while contextualizing the speed claims.

## Key dated facts

### February 2026: the MoE breakthrough and the "12x faster" claim

- In its **February 2026 release**, Unsloth announced a dedicated MoE training pipeline, headlined as up to **12x faster MoE training** with maintained accuracy. Unsloth's own MoE post attributes the top multipliers (up to ~7–12x) to B200-class hardware. [VENDOR]
- April 2026 secondary reporting (when Unsloth stood at ~61,000 GitHub stars) lists as recently shipped: **Unsloth Studio** — web UI building data recipes from PDFs, CSVs, and DOCX files, then fine-tuning without writing code; **MoE model optimization** ("12x faster training for Mixture of Experts architectures"); **FP8 training** on consumer hardware; **500K context-length training** on consumer hardware via custom Triton kernels and padding-free optimization.
- The same snapshot characterizes Unsloth as "the default tool for local fine-tuning," supporting NVIDIA GPUs, CPUs, and Apple MLX, handling Gemma 4, Qwen 3.5, Llama 3.1, and 500+ other models, under a **dual license (Apache 2.0 for core, AGPL-3.0 for Studio)**.
- Practical example from April 2026 reporting: a legal team fine-tuning Llama 3.1 8B on case-law summaries — full QLoRA reportedly under 2 hours on an RTX 4090.
- **2026-02-09**: independent practitioner comparison "Unsloth vs Standard Training" (Pelin Balci, Medium) published, covering speed, memory, *and loss*, with a reproducible notebook — the earliest independent check of Unsloth's loss-parity claim.
- Current Unsloth docs (as recorded in Wave 1 research) phrase the MoE claim as **"MoE LLM training 12x faster with 35% less VRAM"** for DeepSeek, GLM, Qwen, and gpt-oss families. [VENDOR]
- A July 2026 community review of docs.unsloth.ai warns explicitly: "Do not quote a single number as gospel — cite the docs page you read and hedge," because figures are model-, GPU-, and config-specific against a standard HF+FA2 QLoRA baseline.
- **Assessment of the 12x figure: [VENDOR-REPORTED, HARDWARE-SPECIFIC].** Mechanisms are plausible and partly measured (split-LoRA, routing-metadata caching, bincount routing; component numbers from the NVIDIA collaboration). **No independent reproduction of the headline 12x multiplier was found.**

### May 2026: the auditable evidence — Unsloth × NVIDIA (2026-05-06)

The May 6, 2026 collaboration blog ("How to Make LLM Training Faster with Unsloth and NVIDIA," signed by Daniel & Michael Han) decomposes gains into three concrete, measured optimizations:

1. **Packed-sequence metadata caching** — caching rebuilt metadata and avoiding CPU–GPU resynchronization across layers. Measured on Qwen3-14B QLoRA SFT: **+43.3% forward, +5.8% backward, +14.3% per batch**.
2. **Double-buffered checkpoint reload** — overlapping copies with backward compute instead of serializing on one buffer. Measured: **+8.4% on 8B, +6.7% on 14B, +4.6% on 32B**.
3. **GPT-OSS MoE routing with `bincount`** — grouping tokens once with argsort/bincount instead of repeating the sort per expert, removing repeated synchronization from per-expert dynamic indexing. Measured: **~10–15% in team validation; +23% forward and +13% backward in the targeted routing path**.

The engineering lesson the authors draw: once the math kernels are optimized, further speedups come from eliminating redundant glue work and overlapping the unavoidable work. An independent commenter on the announcement highlighted the same point: the MoE routing fix kills a hidden bottleneck.

### July 2026: independent framework comparison and the loss-head analysis

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

```python
model.save_pretrained_gguf("model_gguf", tokenizer, quantization_method="q4_k_m")
#multi-quant in ONE call:
model.save_pretrained_gguf("model_gguf", tokenizer,
                            quantization_method=["q4_k_m", "q8_0", "f16"])
model.push_to_hub_gguf("username/repo", tokenizer,
                       quantization_method=["q4_k_m", "q8_0", "f16"])
```

- **Full accepted `quantization_method` list (23 entries, verbatim)**: `not_quantized`, `fast_quantized`, `quantized`, `f32`, `f16`, `q8_0`, `q4_k_m`, `q5_k_m`, `q2_k`, `q3_k_l`, `q3_k_m`, `q3_k_s`, `q4_0`, `q4_1`, `q4_k_s`, `q4_k`, `q5_k`, `q5_0`, `q5_1`, `q5_k_s`, `q6_k`, `iq2_xxs`, `iq2_xs`, `iq3_xxs`, `q3_k_xs`.
- **Documented practical guidance**: `q4_k_m` = general default (~4.5 bits/weight, small quality hit); `q5_k_m` = better quality, larger; `q8_0` = near-lossless; `f16` = reference/downstream-requant base; `iq2_xxs`/`iq2_xs`/`iq3_xxs` = extreme compression, quality must be verified. **Apple-silicon note**: `q8_K_XL` upcasts layers to BF16 and is slower than plain bf16 on Macs — use `q8_0` on Macs instead.
- **Save-time OOM guard**: saving large models can OOM even when training fit; `maximum_memory_usage` (default 0.75) can be lowered on the save call. Merging LoRA into a GGUF is the memory peak most users actually hit, and Unsloth exposes a throttle for it.
- Companion save methods: `merged_16bit`, `merged_4bit`, `lora`-only, `push_to_hub_merged`.
- **Limitation**: no `iq4_xs`-style UD quants appear in this official list — the Dynamic v3.0 `UD-*` quant family (see §8) is a separate publishing pipeline (`unsloth/` GGUF org on HF), not a `save_pretrained_gguf` argument. The "one-command LoRA merging to quantized GGUF K-quants" claim is accurate for K-quants/i-quants but does not extend to Unsloth's proprietary UD quants via this API.

### NVFP4 export: available since July 2026 — export-and-serve, not training

- Unsloth's NVFP4 export (v0.1.481-beta, July 2026) predates and parallels the Axolotl NVFP4 MoE LoRA work (July 2026) — present them as parallel tracks, not one source.
- Axolotl's July track: NVFP4 (4-bit) MoE LoRA via ScatterMoE (W4A16) and SonicMoE (W4A4), including lossless adapter merge back into a plain NVFP4 checkpoint.
- **Nuance to keep**: the community `unsloth-cli` fine-tuning docs (September 2026) are explicit that on Thor (JetPack R38.2.2), LFM2.5-1.2B `nvfp4` and `awq` exports were live-tested for **serving** via `vllm-openai:v0.29.0-aarch64` — **not for training**. All Unsloth/NVFP4 claims to date are export-and-serve claims, not NVFP4 *training* claims.

### Qwen3.8-27B fine-tuning guide (September 2026)

Unsloth's official Qwen3.8 guide (docs page live September 2026) documents per-model figures weaker than the generic headline — useful evidence for the benchmark-methodology section:

- **Qwen3.8-27B** (dense 27B unified vision-language model; native text/image/video; thinking controls; 262K context; uses the `qwen3_5` architecture; requires Transformers v5).
- **Vendor training figures for this model: ~1.5x faster with ~50% less VRAM than FA2 setups** (no accuracy loss) — *below* the 2x/70% headline, illustrating that "2x/70%" is a class of improvement, not a contract.
- **VRAM tiers, explicitly stated: QLoRA works with 24GB; LoRA needs >36GB; full fine-tuning (FFT) works but uses 4x more VRAM.** This is Unsloth's own documentation confirming that 27B-class LoRA already exceeds 24GB — reinforcing the verdict that 70B on a single 24 GB card is not demonstrated.
- **New kernel detail: "Flash Linear Attention kernels"** (Gated DeltaNet) for Qwen3.8 training; first run takes longer while kernels compile.
- **Memory mechanics**: `offload_embedding=True` keeps the large untied input embedding in RAM, reducing resident VRAM (auto-disabled on unsupported platforms).
- **RL path**: Qwen3.8 RL uses the Qwen3.5 RL path with fast vLLM inference disabled.
- Export follows the one-call API: `save_pretrained_gguf("qwen38_gguf", tokenizer, quantization_method="q4_k_m")`, plus `push_to_hub_gguf`, merged-16bit for vLLM, or LoRA-adapters-only save.

### MoE-on-modest-hardware: documented fitting mechanisms (not new benchmarks)

- **MoE expert layers → system memory** (Studio v0.1.501-beta): automatic GPU placement pushes MoE expert layers into RAM so larger models fit; also added split-model-across-GPUs, tensor parallelism, per-model-per-quant hardware profiles, improved GGUF memory estimates. This is a Studio **placement/inference** feature — keep it in a separate evidence bucket from training kernels.
- **Apple Silicon**: fine-tune large MoE models (text or images) via MLX (v0.1.805/806-beta); up to **30x faster follow-up turns** in long Qwen chats [VENDOR].
- **Training/export fixes** (v0.1.501-beta): text-only training with multimodal models no longer truncates long examples before packing; **fine-tuned Qwen3.5 and Qwen3.6 MTP models now export correctly to GGUF** (previously broken); Windows GGUF-export final-write permission error fixed.
- Also in v0.1.501-beta: opt-in **MCP endpoint** (`UNSLOTH_STUDIO_ENABLE_MCP=1`) letting AI clients inspect models and training history, start/stop training, load checkpoints, validate recipes, and export GGUFs; dynamic NVFP4 models shipped in the preceding Studio release.
- Adjacent 2026 evidence that consumer training keeps pushing downward: a "Soup"-style layer-streaming approach (August 2026) fine-tunes 8B models on a 4 GB laptop GPU.

### Model-publishing cadence and third-party tooling (September 2026)

- Unsloth's Hugging Face organization (`unsloth/`, 1,374 models) shows uploads within days of 2026-09-22, including `unsloth/Laguna-S-2.1-GGUF` and the Dynamic v2.0/v3.0 lines (`unsloth/DeepSeek-V4-Flash-GGUF`, `unsloth/GLM-5.2-GGUF`, `unsloth/gemma-4-*-it-NVFP4`). **Dynamic v3.0 remains the current UD quant generation.**
- GLM-5.2's Dynamic GGUF was presented by a secondary source (LinkedIn) as 84% size reduction (~239 GB at 2-bit, ~217 GB at 1-bit) with ~82%/~76% accuracy retention and ~98% at 4-bit, runnable on 256 GB unified-memory machines — [UNVERIFIED] (vendor-originated, no independent benchmark).
- The community-built **`unsloth-cli`** wrapper (`agentculture/unsloth-cli`, releases **0.7.0 and 0.7.1, 2026-09-15**) — not an Unsloth product — operationalizes the export workflow: container-backed multi-format export (`sloth export --format merged-16bit | merged-4bit | gguf | awq | nvfp4` with `--quant`, `--calib`, `--calib-samples`, `--base`, `--force`, `--dry-run`, `--keep-intermediate`; disk estimates, no-clobber, atomic `.partial` output, provenance via `export.json` + per-adapter `exports.json`); **quantization-loss check** (`sloth eval --model DIR` scores merged/quantized outputs).
- Community economics datapoint: a third-party DGX Spark project (~May 2026) claims 7.67x LoRA / 8.35x full-FT speedup for Qwen3.5 (0.8B–27B) on a GB10 DGX Spark vs stock Unsloth, with wall-clock parity to a rented H100 (~20 h) at $0 vs ~$50 — [COMMUNITY, UNVERIFIED]; the author prices DGX Spark self-hosting breakeven at ~79 days of continuous training vs $2.49/h H100 rental.

### Correction ledger: World's Fair, funding, versioning

- **AI Engineer World's Fair 2026** (Moscone West, San Francisco, June 30 – July 2, 2026; ~300 speakers, 29 tracks) is a verified event. Daniel Han has documented AIE appearances in 2024 and 2025, a LlamaCon 2025 Dynamic Quantization session, and a mid-July 2026 kernels/RL seminar — making a 2026 World's Fair appearance plausible. But **no session title, track, day, or schedule entry confirms Unsloth presented at the 2026 Fair**, and the **"112 slides"** figure has no supporting source. **Verdict: [UNVERIFIED]** — do not state as fact.
- **Unsloth seed funding ~$40K**: [UNVERIFIED]. No primary or reliable secondary source confirms a ~$40K raise, YC as an investor, or Logan Kilpatrick as an investor. The only concrete funding figure found is a third-party analyst scorecard (agentvc-index, 2026-03-23): **~$500K seed (Redpoint Ventures scout, Samsung NEXT)** — which contradicts $40K and names different investors. Keep both figures flagged; treat neither as established.
- Confirmed context: founded by brothers **Daniel and Michael Han** (Australia-based); Daniel has an NVIDIA background; ~57K GitHub stars and 150M+ downloads (March 2026 figures); Google, OpenAI, Meta, and NVIDIA are documented **partners** (official fine-tuning partnerships), not investors; commercial products are Unsloth Pro and Unsloth Studio.
- **Dynamic Quantization 2.0**: current as of Daniel Han's July 2026 talks (LlamaCon 2025 session: 1.5x faster training, 50% less VRAM, 8x longer context for Llama 4) — but **superseded by Dynamic v3.0 on 2026-08-19**. Any sentence presenting 2.0 as current must carry a "as of July 2026" scope and a pointer to §8.

