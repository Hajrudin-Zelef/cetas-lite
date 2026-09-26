---
id: ai-industry-kb-2026/10-training-fine-tuning/distributed-training-in-2026-fsdp2-as-the-default
title: "Distributed training in 2026: FSDP2 as the default"
domain: training-fine-tuning
role: deep-dive
task: training
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Hugging Face", "Nvidia", "Unsloth", "Z.ai"]
dates: ["2026-02", "2026-02-09", "2026-02-10", "2026-03-18", "2026-05", "2026-05-06", "2026-05-13", "2026-06", "2026-07-08", "2026-07-22", "2026-08-12", "2026-08-19", "2026-08-30", "2026-09", "2026-09-02", "2026-09-09", "2026-09-15", "2026-09-17", "2026-09-22"]
keywords: ["training", "amd", "attention", "benchmark", "consumer", "deepseek", "diffusion", "dpo", "fine-tuning", "fp8", "gguf", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5640, 5698]
section: "10. Training & Fine-Tuning"
sha256: 3dcd0186e6a63e47260b96df84e3c7ae90397d91582d6bb6a01f6231d4f9f4c3
---

# Distributed training in 2026: FSDP2 as the default

| Date | Event |
|---|---|
| 2026-02-09 | Pelin Balci publishes independent "Unsloth vs Standard Training" comparison (speed, memory, loss) |
| 2026-02-10 | DeepSpeed AutoEP blog (automatic expert-parallelism configuration) |
| 2026-02 | Unsloth February release: MoE training pipeline, "12x faster MoE training" headline [VENDOR] |
| 2026-03-18 | PyTorch 2.11 released |
| 2026-04 | Unsloth ~61K GitHub stars; "default tool for local fine-tuning"; Studio (data recipes from PDFs/CSVs/DOCXs); FP8 training; 500K-context training on consumer hardware; Llama 3.1 8B QLoRA < 2 h on RTX 4090 |
| 2026-05-06 | Unsloth×NVIDIA collaboration blog: 3 measured optimizations (packed metadata +43.3% fwd / double-buffered checkpointing +4.6–8.4% / bincount MoE routing +23% fwd +13% bwd) |
| 2026-05 | torchtune paper (arXiv 2605.21442v1): Unsloth = memory leader at every tested size; torchtune+torch.compile = throughput leader at 3 of 4 sizes |
| 2026-05 | PEFT-Arena: OFT variants may offer the better retention/adaptation Pareto |
| 2026-05-13 | PyTorch 2.12 released |
| ~2026-05 | Community DGX Spark project claims 7.67x LoRA / 8.35x full-FT for Qwen3.5 on GB10 vs stock Unsloth [COMMUNITY, UNVERIFIED] |
| 2026-07-08 | PyTorch 2.13 released |
| 2026-07 | Studio v0.1.481-beta: NVFP4/FP8/imatrix GGUF export; DeepSeek-V4-Flash; "GRPO 1.3x", "MoE training 3–5x" [VENDOR] |
| 2026-07 | Chew Loong Nian's loss-head analysis: 87.3% of sequence-scaling memory in Llama 3.1 8B LoRA is the cross-entropy head; 95.2% for gpt-oss-20b; 49.1 GB at 16K context |
| 2026-07 | Axolotl NVFP4 MoE LoRA (ScatterMoE W4A16, SonicMoE W4A4); Unsloth NVFP4 export — parallel tracks |
| 2026-07-22 | MarkTechPost 4-framework comparison: gpt-oss-20b in 12.8 GB; Qwen3-30B-A3B 16-bit LoRA 63 GB; B200 8K 47.43 vs 73.80 GB, 16K 55.13 vs OOM |
| 2026-08-12 → 2026-08-30 | Stars 70.4K → 75.2K (~400–570/day) |
| 2026-08 | "Soup"-style layer-streaming approach fine-tunes 8B models on a 4 GB laptop GPU |
| 2026-08-19 | **Dynamic v3.0 launched — supersedes Dynamic Quantization 2.0** (format details: §8) |
| 2026-09-02 | PyTorch 2.14 released (silent clamp/min/max boundary-subgradient change — training-reproducibility hazard); `unsloth/_version.py` bumped to 2026.9.2 |
| ~2026-09-early | Studio v0.1.805/806-beta: MTP-by-default for Qwen3.8-Flash-Next + GLM-5.3-Flash ("up to 2x faster generation" [VENDOR]); 170+ improvements; MLX MoE training |
| 2026-09-09 | MoE Triton grouped-GEMM correctness fix (down-LoRA row order; scatter via `gather_indices`) + dtype-aware backend selection |
| ~2026-09-mid | Studio v0.1.807/808-beta: AMD Vulkan by default; INT8/FP8 diffusion 1.2–1.7x (downward revision of the Sept-17 "2x"); gated-delta training +25% on Apple Silicon [VENDOR]; DoRA/DPO on MLX |
| 2026-09-15 | Community `unsloth-cli` 0.7.0/0.7.1 (container export, quantization-loss eval) [COMMUNITY] |
| 2026-09-17 | Official changelog "Docker + MultiUser + AMD Support" (latest entry as of 2026-09-22) |
| 2026-09-17 | TRL Unsloth-integration docs (GGUF one-call API, 23-entry quant list, OOM guard) verified current |
| September 2026 | Qwen3.8-27B fine-tuning guide: ~1.5x/50% per-model figures; QLoRA 24 GB / LoRA >36 GB / FFT 4x VRAM |
| 2026-09-22 | No releases newer than the 2026-09-17 changelog found; HF `unsloth/` org (1,374 models) shows uploads within days |

### Distributed training in 2026: FSDP2 as the default

- **FSDP2 is the 2026 default distributed substrate** for multi-GPU fine-tuning (torchtune's DTensor stack; Axolotl's FSDP2-compatible `quantize_moe_experts`; Liger's FSDP2 integration).
- **Axolotl (June 2026)** added **Expert Parallelism (EP) for distributed MoE training via DeepEP**, documented under its N-dimensional parallelism docs — moving DeepEP from the pretraining world into the fine-tuning YAML-config world.
- **torchtune (May 2026 paper)** builds its parallelism stack on PyTorch DTensor: FSDP2, tensor and sequence parallel, **a custom expert-parallel plan for MoE models**, **loss parallel** (sharding output features over the vocabulary dimension across the TP mesh so full logits are never materialized), context parallelism via Ring Attention. The same recipes scale from a single H100 to multi-node FSDP2 clusters without rewriting the training loop.
- Together these confirm that **expert parallelism is becoming a configurable primitive in fine-tuning frameworks**, not just a pretraining-systems concern.

### The framework-selection rule (falls out of the 2026 evidence)

- **Single/dual-GPU PEFT and local workflows → Unsloth** (lowest memory, simplest).
- **YAML-driven multi-GPU with FSDP2/EP and broad model coverage → Axolotl**.
- **Hackable native-PyTorch research with custom parallelism → torchtune**.
- **Trillion-parameter or heavily customized scale → DeepSpeed/Megatron-style**.

### Benchmark methodology: reading Unsloth numbers correctly

1. **The baseline dominates the headline.** Unsloth's figures are vs Hugging Face + FlashAttention 2 QLoRA. Against that baseline: 1.87–2.74x speedups and 12–74% VRAM savings (HF's own blog, 59 runs). Against torchtune with torch.compile (May 2026 paper), Unsloth wins on memory but loses on throughput at most sizes. Always record the baseline.
2. **Report the config, not just the model.** The >75% VRAM figure is Llama 3.3 70B, Alpaca, batch 2, grad-accum 4, rank 32, QLoRA on all linear layers, 80 GB GPU. Change any of these and the figure moves.
3. **Separate the three VRAM figures.** ~70% = general marketing claim (still current); >75% = Llama 3.3 70B QLoRA vendor benchmark; ~80% = GRPO-specific (Qwen3 4B row). The 80%-as-universal reading is unverified.
4. **Separate MoE-specific from general figures.** The 12x MoE headline is hardware-specific (B200 in Unsloth's own post) and vendor-reported; the component-level measurements (NVIDIA blog: +43.3% forward packed metadata, +8.4%/+6.7%/+4.6% checkpoint reload, +23%/+13% bincount routing) are the auditable part.
5. **Vendor vs independent.** Vendor: unsloth-zoo tables, Unsloth docs/blog. Independent: HF's Unsloth–TRL blog, torchtune paper, MarkTechPost comparison, Balci's loss-aware comparison, Chew's loss-head analysis. **No independent reproduction of the 12x MoE multiplier or the 12.8 GB gpt-oss-20b figure was found.**
6. **Loss parity is claimed, and partly checked.** Unsloth claims 0% accuracy degradation (exact optimization, no approximations). The February 2026 Balci comparison is the independent practitioner check that included loss; the HF blog predates the MoE work but also reported 0% degradation for standard QLoRA.

### Fine-grained MoE and bandwidth — cross-ref to §9

The MoE architecture question (whether fine-grained experts hurt inter-GPU bandwidth) is resolved in §9: fine granularity improves routing specialization and combinatorial flexibility, but MoE Parallel Folding (arXiv 2504.14960v2) shows fine-grained MoE has *lower* training efficiency than coarse-grained MoE across tested parallelism strategies — bandwidth is preserved only by co-designed mitigations (DeepEP-class dispatch, latent-space routing, Quantile Balancing). The claim that fine granularity improves routing "without hurting inter-GPU bandwidth" is **not verified and contradicted as stated**; this section does not duplicate that analysis.

## Implications

