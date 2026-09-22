---
id: ai-industry-kb-2026/10-training-fine-tuning/figures-and-metrics
title: "Figures and metrics"
domain: training-fine-tuning
role: deep-dive
task: training
actors: ["AMD", "Alibaba", "Apple", "Hugging Face", "Mistral", "Nvidia", "Samsung", "Unsloth"]
dates: ["2026-02", "2026-05", "2026-07"]
keywords: ["accelerator", "amd", "benchmark", "compute", "dpo", "fine-tuning", "gpu", "llama", "lora", "memory", "mistral", "moe"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5529, 5601]
section: "10. Training & Fine-Tuning"
sha256: c4d131fc646863fed527a2f19ebd9a54b23e8cf3915e7a6035b39a1c898bfc52
---

# Figures and metrics

## Figures and metrics

### Training mechanics: where the memory savings come from

Unsloth wraps Transformers, PEFT, and TRL and monkey-patches optimized paths rather than replacing the trainer API — a design that makes it a drop-in accelerator for standard HF/TRL workflows. Cross-confirmed across Unsloth's docs, the May 2026 NVIDIA blog, the July 2026 MarkTechPost comparison, and community skill documentation:

1. **Hand-written Triton kernels** for LoRA/QLoRA forward and backward passes, with manually derived backpropagation steps — no approximations, hence the "0% accuracy degradation" claim (exact optimization). Hugging Face's Unsloth–TRL blog describes this as "overwriting some parts of the modeling code with optimized operations" and "manually deriving backpropagation steps and rewriting all PyTorch modules into Triton kernels."
2. **Split-LoRA for MoE** — avoids materializing LoRA deltas for all experts before the MoE matmul. The key MoE-specific memory win and what the February 2026 MoE release optimized, alongside the bincount-based GPT-OSS routing fix.
3. **Fused / cut cross-entropy** — avoids materializing the full vocabulary-logit tensor. The torchtune paper uses the same "linear cross-entropy" idea and notes it reduces peak memory during loss computation, especially for large vocabularies. The independent July 2026 loss-head analysis confirms this is the actual axis of differentiation between frameworks.
4. **Custom gradient checkpointing** — double-buffered checkpoint reload overlapping copies and backward compute (+8.4% on 8B, +6.7% on 14B, +4.6% on 32B per the NVIDIA blog).
5. **Padding-free packing** — packed-sequence metadata cached across layers (+43.3% forward, +5.8% backward, +14.3% per batch on Qwen3-14B QLoRA SFT).
6. **Optimized RoPE/MLP paths and reduced memory fragmentation** — pre-quantized 4-bit model loading reportedly cuts fragmentation by ~500 MB (per the HF TRL integration blog), allowing larger batches.

What Unsloth does *not* do: it is strongest on **1–2 GPU PEFT/local workflows**. Multi-node, full-parameter, and expert-parallel training remain the domain of Axolotl, torchtune/FSDP2, DeepSpeed, or Megatron-style systems — and Unsloth's own pricing tiers (per July 2026 community docs: Pro/Enterprise add multi-GPU/multi-node and full-parameter training) reflect that boundary. A July 2026 community reality-check summarizes the headline honestly: "~2x faster training with ~70% less VRAM, no accuracy loss, on a single GPU with the free core. Treat that as a *class* of improvement, not a contract."

### Vendor benchmark anchors (unsloth-zoo README; Alpaca, batch 2, grad-accum 4, rank 32, QLoRA)

| Model | VRAM | Speed vs HF+FA2 | VRAM reduction | Longer context |
|---|---|---|---|---|
| Llama 3.3 (70B) | 80 GB | 2x | >75% | 13x |
| Llama 3.1 (8B) | 80 GB | 2x | >70% | 12x |

Context-length translation (MarkTechPost, independent): on an 80 GB A100, **Llama 3.3 70B max context — Unsloth 89,389 tokens vs Transformers+FA2 6,916**; Llama 3.1 8B: 16 GB → 40,724 vs 2,551; 24 GB → 78,475 vs 5,789; 48 GB → 191,728 vs 15,502; 80 GB → 342,733 vs 28,454.

Per-model notebook table (vendor, unsloth-zoo) — the source of the GRPO-specific 80% figure, showing workload-dependence:

| Model | Speed vs baseline | VRAM reduction |
|---|---|---|
| Qwen3 (14B) | 2x | 70% |
| Qwen3 (4B) GRPO | 2x | 80% |
| Gemma 3n (4B) | 1.5x | 50% |
| Mistral v0.3 (7B) | 2.2x | 75% |
| Llama 3.1 (8B) | 2x | 70% |
| Llama 3.2 Vision (11B) | 2x | 50% |

### Older independent anchor (Hugging Face's own Unsloth–TRL blog; 59 runs, T4/A100, rank 16, baseline Transformers 4.36)

| Setup | Speedup vs HF | VRAM reduction |
|---|---|---|
| CodeLlama 34b (A100 40GB, Slim Orca) | 1.94x | −22.7% |
| Llama-2 7b (A100 40GB, Slim Orca) | 1.87x | −39.3% |
| Mistral 7b (A100 40GB, Slim Orca) | 1.88x | −65.9% |
| TinyLlama 1.1b (A100 40GB, Alpaca) | 2.74x | −57.8% |
| DPO Zephyr (A100 40GB, Ultra Chat) | 1.88x | −11.6% |
| TinyLlama 1.1b (free Colab T4, Alpaca) | 3.87x | −73.8% |

Note the variance: VRAM savings range from ~12% to ~74% depending on model and setup — the argument for always carrying the config with the headline. The blog's summary: "up to 2.7x faster and up to 74% less memory" on free-tier Colab.

### Independent balance: the torchtune paper (May 2026; single-H100 LoRA ablations, Qwen3, Alpaca, seq 2048, microbatch 2, r=16/α=16; memory GB / throughput tokens/s/GPU)

| Model | torchtune | Axolotl | Unsloth |
|---|---|---|---|
| Qwen3 0.6B | 2.6 / 3,292 | 6.2 / 1,973 | **2.3** / 2,502 |
| Qwen3 1.7B | 4.6 / 3,610 | 9.4 / 2,233 | **4.4** / 3,284 |
| Qwen3 4B | 9.3 / 2,616 | 17.7 / 1,605 | **8.9** / 1,826 |
| Qwen3 8B | 17.2 / 2,745 | 27.9 / 1,609 | **16.8** / 1,836 |

Reading: **Unsloth is the memory leader at every size** — confirming its VRAM advantage independently. **torchtune with torch.compile leads on throughput at 3 of 4 sizes** — the vendor "2x faster" claim is relative to the Hugging Face + FlashAttention 2 baseline, not an absolute speed crown. The same paper's DPO comparison (single GH200 96 GB) found **Axolotl OOM where torchtune fit** with standard AdamW (91.02 GB / 782.2 tok/s on Qwen3 8B; Axolotl + 8-bit optimizer: 69.66 GB / 185.6 tok/s) — framework choice changes *feasibility*, not just speed.

### Claim-by-claim verdicts

| Claim | Verdict | Evidence grade |
|---|---|---|
| "12x faster MoE training" | True as a vendor headline on B200-class hardware; **no independent reproduction found** | [VENDOR-REPORTED, HARDWARE-SPECIFIC] |
| VRAM reduction 70–80% on 70B+ fine-tuning | >75% is real but narrow (Llama 3.3 70B QLoRA, Alpaca, bs2, GA4, r32, 80 GB); ~70% is the general marketing figure; ~80% is GRPO-specific (Qwen3 4B row); HF's own docs show 22.7–65.9% variance across models | [PARTIALLY VERIFIED, WORKLOAD-SPECIFIC] |
| A 70B+ model fine-tunes on one A100/H100 | **Verified for QLoRA** (Llama 3.3 70B, single 80 GB GPU). Excludes full-parameter (Adam optimizer states alone exceed 80 GB) and excludes every-70B-architecture | [VERIFIED FOR QLORA] |
| Practical MoE/70B fine-tuning on RTX 4090/5090 (incl. clusters) | 8B-class QLoRA on RTX 4090 verified (< 2 h, Llama 3.1 8B); gpt-oss-20b in 12.8 GB partially verified; Qwen3-30B-A3B 16-bit LoRA needs 63 GB (> 24 GB). **70B on a single 24 GB card: [UNVERIFIED]; 4090/5090 clusters as 70B+ MoE training: [UNVERIFIED]** | [PARTIALLY VERIFIED] with unverified core |
| MoE training 3–5x (July 2026) | Vendor release note, superseded in specificity by the May NVIDIA-blog component measurements | [UNVERIFIED → RECORD WITH DATE] |
| MTP-by-default "up to 2x faster generation" (Sept 2026) | Vendor release note; no independent benchmark found | [UNVERIFIED] |
| 30x faster follow-up turns on Mac / 25% faster gated-delta on Apple Silicon / AMD Vulkan gains (Sept 2026) | Vendor release notes; no independent reproduction found | [UNVERIFIED] |
| AI Engineer World's Fair 2026 appearance + 112-slide deck | No session/schedule evidence; figure has no source | [UNVERIFIED] |
| Seed ~$40K; YC + Logan Kilpatrick investors | No source; contradicted by a third-party ~$500K (Redpoint scout, Samsung NEXT) figure | [UNVERIFIED], keep both flagged |

