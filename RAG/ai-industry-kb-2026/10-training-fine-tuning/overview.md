---
id: ai-industry-kb-2026/10-training-fine-tuning/overview
title: "10. Training & Fine-Tuning"
domain: training-fine-tuning
role: deep-dive
task: training
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Hugging Face", "Nvidia", "Samsung", "Unsloth", "Z.ai"]
dates: ["2026-02", "2026-02-09", "2026-04", "2026-05", "2026-05-06", "2026-07", "2026-08-19", "2026-09", "2026-09-17"]
keywords: ["fine-tuning", "training", "amd", "apache", "benchmark", "compute", "consumer", "deepseek", "diffusion", "fp8", "gguf", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5387, 5426]
section: "10. Training & Fine-Tuning"
sha256: aa321d7d64a576d7e441769179245fc72a9b230f1d4a227309a51e4a7060fe13
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

