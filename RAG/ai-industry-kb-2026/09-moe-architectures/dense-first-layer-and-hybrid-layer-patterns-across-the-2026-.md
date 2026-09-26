---
id: ai-industry-kb-2026/09-moe-architectures/dense-first-layer-and-hybrid-layer-patterns-across-the-2026-
title: "Dense-first-layer and hybrid-layer patterns across the 2026 fleet"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "China", "DeepSeek", "MiniMax", "Moonshot", "OpenRouter", "Z.ai"]
dates: ["2025-08-05", "2026-04-16", "2026-04-24", "2026-05", "2026-06-01"]
keywords: ["agents", "attention", "compute", "cost", "cost per token", "decode", "deepseek", "fp4", "glm", "inference", "kimi", "llama"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4808, 4849]
section: "9. MoE Architectures"
sha256: a6c78ae569a9e1b168df6334a38843f56ee265dc52a9e6168077272354628c99
---

# Dense-first-layer and hybrid-layer patterns across the 2026 fleet

- **Standard token-choice top-k (the 2026 default):** router scores s(x) = W_g·x → softmax → select top-k → renormalize over selected experts; layer output = Σ renormalized_gates × expert(x). Noise added during training for exploration/stability; capacity factor C = ⌈α·T·k/N⌉ caps tokens per expert and can drop overflow tokens.
- **Shared experts (universal by 2026):** one or more experts process ALL tokens alongside routed experts, capturing common knowledge and reducing redundancy among specialists. Counts: 1 shared (MiniMax M3, Qwen3.6, Gemma 4, Qwen3.8-Max, DeepSeek-V4), 2 shared (Kimi K3), 0 shared (gpt-oss, MiMo-V2/V2.6, GLM-5.3-Flash — the three outliers).
- **Fine-grained expert segmentation (DeepSeekMoE → industry standard):** split large experts into many small ones (expert intermediate dims 512–1536 in 2026 models); the number of achievable expert combinations explodes, routing decisions get more granular, compute efficiency improves per flop on the routing side — at the bandwidth cost documented above.
- **Load balancing — the field moved away from auxiliary losses:** traditional auxiliary loss (L = α·N·Σ f_i·P_i) causes interference gradients (~23% degradation cited in one 2026 analysis). State of the art by 2026: **DeepSeek-V3-style dynamic bias** (learnable bias terms adjusted from actual load, no gradient interference — retained in V4 with a small sequence-wise balance loss); **expert-choice routing** (experts pick top-k tokens — intrinsically balanced); **QB/quantile balancing** (Kimi K3: each expert's bias set to the router-score quantile matching its target load, single global histogram/all-reduce); orthogonality + variance losses on routing scores in fine-grained designs.
- **V4 routing specifics [PARTIALLY VERIFIED]:** affinity changes from sigmoid to `sqrt(softplus(...))`; first three blocks use deterministic token-ID hash routing instead of learned routing; routed expert weights deployed in FP4.
- **Latent routing (Kimi K3, the 2026 frontier):** **Stable LatentMoE** — the full hidden state is projected into a narrower normalized latent space (ℓ = 3,584 vs 7,168) before expert dispatch, keeping activation gradients stable across 896 experts and roughly halving routed-expert activation traffic and expert compute. This is top-k routing relocated into a compressed space, not a different routing paradigm — and it is the concrete counterexample to any claim that fine granularity is bandwidth-free.
- **2026 research frontier (papers, not yet production):** **MoUE — Mixture of Universal Experts** (arXiv 2603.04971, Mar 2026) — layer-agnostic universal expert pool shared across layers ("virtual width" from depth), up to +1.3% over matched MoEs and +4.2% when converting existing MoE checkpoints; **FLAME** (arXiv 2605.09355, May 2026) — adaptive MoE for continual multimodal learning (shared + task-specific experts); **phase-aware routing** (Yang et al. 2026) — routing at environment-step rather than token granularity for RL agents, enforcing temporal consistency across multi-step plans.

### Dense-first-layer and hybrid-layer patterns across the 2026 fleet

- MiniMax M3: first 3 layers dense (stabilize early representations before sparse routing).
- Kimi K3: 93 layers with 1 dense layer; 69 KDA + 24 Gated MLA layers.
- MiMo-V2-Flash/V2.6: first block is global attention with a dense FFN, interleaved thereafter; 5:1 SWA/GA ratio with 128-token window.
- Qwen3.6: 10 repetitions of [3× Gated DeltaNet linear attention → MoE + 1× gated standard attention → MoE] (30 linear + 10 attention layers).
- Qwen3.8-Max: Gated DeltaNet (128 linear heads for V / 16 for QK) interleaved with gated attention (64 Q / 4 KV heads, dim 256).
- GLM-5.3-Flash: `glm5_next` — hybrid sparse + linear attention stack with Manifold-Constrained Hyper-Connections (mHC).
- DeepSeek-V4: CSA/HCA hybrid attention replacing the earlier MLA design (reconstructed, [PARTIALLY VERIFIED]).
- Gemma 4: pure-attention MoE, no SSM — sliding-window (1024) alternating with global attention; the only 2026 open MoE with no recurrence or linear-attention component.
- MTP training heads double as speculative-decoding drafters (MiMo-V2: accepted length 2.8–3.6, 2.0–2.6× speedup [VENDOR]; MiMo-V2.6: 5-layer DFlash-style drafter, 7 tokens ahead; Qwen3.6 and the DeepSeek lineage also train MTP heads; M3's 7 MTP modules appear in community notes but weights are absent from shipped checkpoints [COMMUNITY]).

### The bandwidth correction in numbers

- MoE Parallel Folding (arXiv 2504.14960v2): fine-grained MoE models have **lower training efficiency than coarse-grained MoE across tested parallelism strategies** — more experts and more active experts per token increase expert-dispatch communication volume, smaller experts reduce GEMM efficiency, more expert activations increase activation memory, and larger model-parallel groups add communication.
- Kimi K3's priced-and-paid response: latent width ℓ=3,584 vs full width 7,168 halves routed-expert activation traffic and expert compute; per-expert 33.0M params at MXFP4 (E8M0 scale per 32 weights), 2.72T routed parameters total; Quantile Balancing establishes balance thresholds with a single global histogram/all-reduce instead of fixed-step bias updates.
- Inference-side dispatch cost per token ≈ k × hidden-dim × 2 (dispatch + combine), amortized across batch — which is why MoE serving wins from batching, and why ultra-sparse models (<1% activation density) can lose to dense under expert parallelism (measured −7–12% on ROCm for Llama-4-Maverick-class sparsity; full serving guidance in part 09b).
- Bottom line for the RAG: record architecture claims (expert counts, routing formulas, latent widths) and systems claims (bandwidth, VRAM, speedups) as separate facts with separate evidence grades — the 2026 literature shows they move in opposite directions for fine-grained MoE.

### Context-window and modality trajectories

- Context: 128K (gpt-oss, 2025-08) → 256K (MiMo-V2-Flash 2025-12, Gemma 4 ~2026-04) → 262K native / ~1M via YaRN (Qwen3.6, 2026-04-16) → 1M native on every 2026 flagship from DeepSeek-V4 (2026-04-24) onward. GLM-5.3-Flash's config-grounded 1,048,576 (OpenRouter listing up to 1,310,720) sits against an unclear 300K figure in one model-card eval — unresolved.
- Modalities: text-only (gpt-oss, DeepSeek-V4, GLM-5.3 flagship) → text+image (Qwen3.6 built-in vision encoder, Gemma 4's ~550M separate vision encoder) → text+image+video trained from step 0 (MiniMax M3) → text+image+video (Qwen3.8-Max, GLM-5.3-Flash) → full omnimodality including audio (MiMo-V2.6: 681M ViT + 308M audio tokenizer + 127M audio patch encoder, all entering the same model). Kimi K3: native text+image via MoonViT-V2 401M, video documented on the first-party API only.
- Reading the trajectory: multimodality moved from bolted-on adapters (pre-2026) to step-0 mixed training (M3, 2026-06-01) to single-model omnimodality (MiMo-V2.6, ~2026-09) — the MoE backbone stayed constant while the input modalities widened.

### How to read the A-suffix (naming convention reference)

- "35B-A3B" = 35B total parameters, ~3B active per token; "26B-A4B" = 25.2B total, 3.8B active. The suffix is the number that determines per-token FLOPs and decode-time memory bandwidth; the prefix is the number that determines VRAM residency.
- Qwen3.5/3.6/3.7/3.8 are successive flagship generations, not point releases of Qwen3; Gemma 4's "E" prefix (E2B/E4B) denotes effective parameters for the dense siblings. The convention was established by gpt-oss (2025-08-05) and is now field-standard.

## Main actors

### The six Chinese labs that produced every 2026 frontier open MoE

