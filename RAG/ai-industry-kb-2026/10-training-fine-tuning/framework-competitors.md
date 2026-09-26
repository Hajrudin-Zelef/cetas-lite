---
id: ai-industry-kb-2026/10-training-fine-tuning/framework-competitors
title: "Framework competitors"
domain: training-fine-tuning
role: deep-dive
task: training
actors: ["Alibaba", "DeepSeek", "Google", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-02", "2026-02-09", "2026-05", "2026-07", "2026-07-22", "2026-09", "2026-09-15", "2026-09-17"]
keywords: ["apache", "attention", "benchmarks", "bitnet", "consumer", "deepseek", "dpo", "embedding", "fine-tuning", "flash attention", "fp8", "gguf"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5606, 5639]
section: "10. Training & Fine-Tuning"
sha256: ca17fe1c1346dae58713584c2b8246907a5d298c2f7cd46c22ccda8fc566aad0
---

# Framework competitors

- **Daniel Han and Michael Han** (brothers, Australia-based; Daniel has an NVIDIA background) — founders. Documented official fine-tuning partnerships with Google, OpenAI, Meta, and NVIDIA (partners, not investors).
- **Three products by September 2026**: Unsloth Desktop (Tauri app), Unsloth Studio (web UI), Unsloth Core (code library). Model hub ships per-quant downloads with thinking toggle and `reasoning_effort` controls.
- **Dual license**: Apache 2.0 for core, AGPL-3.0 for Studio. Commercial tiers Pro/Enterprise add multi-GPU/multi-node and full-parameter training — reflecting the boundary between Unsloth's core strength (1–2 GPU PEFT) and the multi-node world.
- **RL integration**: plugs into TRL's `GRPOTrainer`/`GRPOConfig` with a built-in vLLM engine (`fast_inference=True`) for rollouts; the 2026-09-17 changelog confirms GRPO with Qwen3.5 plus latest TRL/vLLM. 2026 RL items: 7x longer-context RL via new batching; FP8 & Vision (VLM) GRPO on consumer GPUs; DoRA support; embedding fine-tuning 1.8–3.3x faster. Docs cite ~80% less VRAM for GRPO vs standard setup (flagged "verify" even by community skill docs — treat as vendor claim).
- **Model catalog (Sept 2026)**: Qwen3.8, Qwen3.6 (MTP 1.4–2.2x faster inference, NVFP4), Kimi K3, Kimi K2.7 Code, MiniMax M3, MiniMax-H3, GLM-5.3-Flash, GLM-5.2 (Dynamic GGUF), DeepSeek-V4-Flash, Gemma 4, DiffusionGemma, Qwen3.8-Flash-Next, Muse Glimmer (Meta), Ornith, FLUX, Wan. Catalog scope caveat: verified current Unsloth materials confirm the named families plus NVFP4/GGUF availability in *parts* of the catalog — availability is family/model/hardware-specific, and "Nemotron" was not confirmed in searched official materials. Do not present the catalog as uniform.

### Framework competitors

- **Axolotl** — the YAML-driven multi-GPU fine-tuning framework. 2026 MoE cadence (from its README "Latest Updates"): **February** — ScatterMoE LoRA (LoRA directly on MoE expert weights via custom Triton kernels), SageAttention, GDPO; **March** — `quantize_moe_experts` (greatly reduces VRAM, FSDP2-compatible), new-model support (Mistral Small 4, Qwen3.5/MoE, GLM-4.7-Flash); **April** — Async GRPO (up to 58% faster steps), Flash Attention 4, SonicMoE fused LoRA, EBFT, uv-first packaging; **June** — Expert Parallelism via DeepEP for distributed MoE training, Tinker-compatible APIs for remote training, context parallelism for hybrid SSM models, BitNet 1.58-bit fine-tuning; **July** — NVFP4 (4-bit) MoE LoRA via ScatterMoE (W4A16) and SonicMoE (W4A4), including lossless adapter merge back into a plain NVFP4 checkpoint. Axolotl owns the multi-GPU research lane alongside torchtune.
- **torchtune** (Meta/PyTorch ecosystem; May 2026 paper, arXiv 2605.21442v1) — positions itself between high-level automated trainers and kernel-specialized systems: modular YAML recipes, in-backward optimizer fusion (reducing gradient-buffer lifetime), composable DTensor parallelism (FSDP2 / tensor / sequence / loss / context parallelism plus a custom expert-parallel plan for MoE), linear cross-entropy, and an **asynchronous GRPO recipe** (Ray-coordinated queue + replay buffer, vLLM rollout workers, on-policy and bounded-off-policy modes — a system design, with head-to-head comparisons left to future work). Its stated conclusion: competitive or superior efficiency vs Axolotl and Unsloth while remaining transparent and hackable.
- **DeepSpeed** (Microsoft) — the max-scale path. February 2026 blog introduced **AutoEP** (automatic expert-parallelism configuration); the MoE tutorial supports expert parallelism combined with data parallelism, ZeRO, and model parallelism (docs refreshed ~September 2026); communication-optimization notes identify MoE all-reduce penalties of **3–12 seconds** in problematic large-scale layouts, with multi-rank bucketing as the fix.
- **Liger Kernel** — open-source Triton kernels integrated into TRL, FSDP2, and DeepSpeed; published end-to-end benchmarks on 4× A100 80 GB (Alpaca, BF16, seq 512): Llama 3 8B +42.8% throughput / −54.8% memory; Qwen2 +25.5% / −56.8%; Gemma +11.9% / −51.8%; Mistral +27% / −21%; Phi-3 +17% / −13%. Trainer support for SFT, DPO, GRPO, KTO, GKD. **Complementary** to Unsloth (open, integrated into the HF stack) rather than a drop-in substitute.

### PEFT method researchers (adapter-quality axis)

- **ESFT — Expert-Specialized Fine-Tuning** (DeepSeek AI + Northwestern, arXiv 2407.01906, EMNLP 2024). The MoE-native PEFT method: score expert affinity on task data, fine-tune only the most relevant experts per layer, freeze the rest. Reported: storage ↓ up to 90%, training time ↓ up to 30% vs full-parameter fine-tuning; matching or exceeding full-FT on math/code while retaining general-task performance better than LoRA. Key architectural finding for this topic: **finer-grained experts are more advantageous for ESFT**. Official code: github.com/shahils01/ESFT.
- **DoRA** (arXiv 2402.09353): decomposes weights into magnitude + direction, applies low-rank adaptation to directional components; mergeable at inference with no serving overhead. Reported: LLaMA-7B 78.1 vs LoRA 74.7; LLaMA-13B 81.5 vs 80.5 (average scores); partial-DoRA can beat LoRA with under half the trainable parameters in reported tests.
- **LoRA-FA** (ICLR 2026 submission; arXiv 2511.04021): at r=64, scored 5.7/28.1/57.0 (MT-Bench/HumanEval/GSM8K) with gains over LoRA of +0.5/+19.1/+14.4; DoRA scored 5.9/19.0/52.2 at the same rank.
- **RoRA** reportedly beats DoRA; **Dual LoRA** (ICLR 2026) reportedly outperforms DoRA with fewer parameters; **PiSSA** is best framed as an initialization/quality method, not a memory-system breakthrough.
- **PEFT-Arena** (May 2026, arXiv): stability–plasticity comparison across LoRA, AdaLoRA, DoRA, VeRA, PiSSA, MiLoRA, OFT, IA3, and full fine-tuning. OFT variants may offer the better retention/adaptation Pareto — do not rank methods by downstream accuracy alone.

### Community and third-party ecosystem

- **Chew Loong Nian** (July 2026) — independent practitioner whose first-principles loss-head analysis confirmed the cut-cross-entropy mechanism story.
- **Pelin Balci** (2026-02-09) — early independent comparison including loss parity.
- **agentculture/unsloth-cli** (0.7.0/0.7.1, 2026-09-15) — container-backed multi-format export, quantization-loss eval (`sloth eval`).
- **megastood DGX Spark project** (~May 2026) — [COMMUNITY, UNVERIFIED] speedups for Qwen3.5 on GB10 vs stock Unsloth; self-hosting breakeven analysis.
- **Community skill-documentation authors** (e.g. rsc-harness, fabrik, unsloth-mcp-server, local-model-autotuning) — the ecosystem's de-facto hedge culture ("do not quote a single number as gospel"; treat 80%-GRPO as vendor claim).
- **MarkTechPost** (2026-07-22) — the independent 4-framework comparison.

## Timeline and context

### Full trajectory, February → September 2026

