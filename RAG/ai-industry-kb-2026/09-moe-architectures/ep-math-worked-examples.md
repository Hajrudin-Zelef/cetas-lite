---
id: ai-industry-kb-2026/09-moe-architectures/ep-math-worked-examples
title: "EP math worked examples"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["Alibaba", "China", "DeepSeek", "Google", "MiniMax", "Mistral", "Moonshot", "OpenAI", "Unsloth", "Xiaomi", "Z.ai"]
dates: ["2026-05", "2026-07", "2026-08-12", "2026-08-30"]
keywords: ["apache", "attention", "benchmark", "compute", "cost", "cost per token", "deepseek", "dpo", "flash attention", "glm", "gpu", "inference"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5260, 5284]
section: "9. MoE Architectures"
sha256: 97b584d37ad25d3aedf2d72cc09843ba0ddce496a52136ea12cabe432b7a3127
---

# EP math worked examples

- **Training-system benchmark anchors (dated).** torchtune paper, May 2026 (single-H100 LoRA, Qwen3, Alpaca, seq 2048, microbatch 2, r=16/α=16 — memory GB / throughput tok/s/GPU): 0.6B — torchtune 2.6/3,292, Axolotl 6.2/1,973, Unsloth 2.3/2,502; 1.7B — 4.6/3,610, 9.4/2,233, 4.4/3,284; 4B — 9.3/2,616, 17.7/1,605, 8.9/1,826; 8B — 17.2/2,745, 27.9/1,609, 16.8/1,836. Reading: Unsloth the memory leader at every size (independent); torchtune+torch.compile throughput-leading at 3 of 4 sizes. DPO comparison (single GH200 96 GB): Axolotl OOM where torchtune fit with standard AdamW (91.02 GB / 782.2 tok/s on Qwen3 8B; Axolotl + 8-bit optimizer 69.66 GB / 185.6 tok/s).
- HF's own Unsloth–TRL blog (59 runs, T4/A100, Transformers 4.36 baseline): speedups 1.87–2.74×, VRAM from −11.6% (DPO Zephyr) to −73.8% (TinyLlama on free T4); blog summary "up to 2.7x faster, up to 74% less memory."
- unsloth-zoo vendor benchmark (Llama 3.3 70B QLoRA, one 80 GB GPU, Alpaca, bs2, GA4, r32): 2x speed, >75% VRAM reduction, 13x longer context vs HF+FA2. Context-length translation (independent, 80 GB A100): Llama 3.3 70B max context — Unsloth 89,389 tokens vs Transformers+FA2 6,916; Llama 3.1 8B: 16 GB → 40,724 vs 2,551; 24 GB → 78,475 vs 5,789; 48 GB → 191,728 vs 15,502; 80 GB → 342,733 vs 28,454.
- Per-model notebook table (vendor): Qwen3 14B 2x/70%; Qwen3 4B GRPO 2x/80%; Gemma 3n 4B 1.5x/50%; Mistral v0.3 7B 2.2x/75%; Llama 3.1 8B 2x/70%; Llama 3.2 Vision 11B 2x/50%. The spread (50–80% VRAM, 1.5–2.2x speedup) is itself the methodology lesson: quote the row, not the headline.
- Liger Kernel end-to-end (4× A100 80 GB, Alpaca, BF16, seq 512): Llama 3 8B +42.8% throughput / −54.8% memory; Qwen2 +25.5%/−56.8%; Gemma +11.9%/−51.8%; Mistral +27%/−21%; Phi-3 +17%/−13%. TRL docs: ~+20% multi-GPU throughput, ~−60% memory, up to 4× longer context.
- ESFT (DeepSeek AI + Northwestern, arXiv 2407.01906): storage ↓ up to 90%, training time ↓ up to 30% vs full-parameter FT; matches/exceeds full FT on math/code; finer-grained experts more advantageous for ESFT.
- Axolotl 2026 MoE cadence: February ScatterMoE LoRA + SageAttention; March `quantize_moe_experts` (GLM-4.7-Flash reserved memory ~127 GiB → ~23 GiB); April Async GRPO (up to 58% faster steps), Flash Attention 4, SonicMoE; June DeepEP-based EP; July NVFP4 MoE LoRA.
- Unsloth star trajectory: ~61K (Apr 2026) → 70.4K (2026-08-12) → 72.2K → 72.9K → 73.4K → 75.2K (2026-08-30), ~400–570/day — developer mindshare, not enterprise dominance.
- July 2026 independent byte-level analysis (Chew Loong Nian): for Llama 3.1 8B LoRA, 87.3% of sequence-scaled memory is the cross-entropy loss head (95.2% for gpt-oss-20b; at 16K context that one tensor wants 49.1 GB vs a 12.8 GB 4-bit model) — cut/fused cross-entropy is the real framework differentiator; his calculator predicted the Transformers+FA2 baseline at 27,713 tokens on 80 GB, close to Unsloth's published 28,454.

### EP math worked examples

- **Training-side:** at 64-way EP with V3-class dims (tokens T, active experts k, hidden D), per-layer all-to-all ≈ 2 × T × k × D × (G−1)/G reaches TB-scale per layer per step — the dominant training cost alongside optimizer memory (~8× weights for Adam-style). DeepSeek-V3's answer: 16-way pipeline × 64-way expert × ZeRO-1 DP on 2048 H800, DualPipe, DeepEP with compute-communication overlap.
- **Inference-side:** dispatch cost per token ≈ k × hidden-dim × 2 (dispatch + combine), amortized across batch — why MoE serving wins from batching and why ultra-sparse (<1% activation) models can be 7–12% slower than dense under EP while standard (≥3%) MoE wins.
- **Granularity arithmetic:** C(64,8) ≈ 4.4 billion achievable expert combinations vs C(16,2) = 120 — the routing-flexibility argument for fine-grained experts, which must be weighed against the MoE Parallel Folding finding (lower training efficiency than coarse-grained across tested strategies: dispatch volume ↑, GEMM efficiency ↓, activation memory ↑).
- **Kimi K3 latent arithmetic:** hidden 7,168 → latent 3,584 halves routed-expert activation traffic and compute; 16-of-896 routing (sparsity ~27:1); 2.72T routed params at 33.0M per expert; MXFP4 with E8M0 scale per 32 weights.
- **Residency arithmetic:** OLMoE profile — active params agree within 8.9% vs OLMo-1B dense, resident memory differs 5.9×; ratio 5.40× invariant at fp32 and int4. Quantization shrinks the numerator and denominator together; the total/active ratio is the invariant, and it is what prices the fleet.

- **License and release counts:** 7 permissive (Apache 2.0: Qwen3.6, Gemma 4 family, gpt-oss-120b/20b; MIT: Kimi K3, GLM-5.3-Flash, MiMo-V2-Flash/V2.6) vs 3 custom (M3, GLM-5.3, Qwen3.8-Max) — permissive licenses cover the majority of 2026 open MoE releases.
- **Geography:** 6 Chinese labs produced every frontier open-weight MoE release of 2026 (DeepSeek, Alibaba/Qwen, Moonshot, Z.ai, MiniMax, Xiaomi); 1 non-Chinese lab (Google, Gemma 4); Western baseline from 2025 (OpenAI gpt-oss).
- **Release velocity 2025-08 → 2026-09:** 2 releases (2025) → the 2026 cascade: 04-16 Qwen3.6, 04-24 DeepSeek-V4, ~04 Gemma 4, 06-01 MiniMax M3, 06-13/17 GLM-5.2, 07-27 Kimi K3, 08-03 Qwen3.8-Max, 08-25/26 GLM-5.3-Flash, 08-28 GLM-5.3, ~09 MiMo-V2.6.
- **Sparsity-ratio ladder:** ~6.6:1 (Gemma 4) → ~12:1 (Qwen3.6) → ~18.6:1 (MiniMax M3) → ~20.6:1 (MiMo-V2-Flash) → ~21.8:1 (V4-Flash) → ~25:1 (Qwen3.8-Max) → ~27:1 (Kimi K3, disputed ~56:1 on the ~50B reading). Total params span 21B–2.8T; active 3B–104B — two orders of magnitude in total, ~35× in active.

## Main actors (continued)

