---
id: ai-industry-kb-2026/09-moe-architectures/moe-ssm-hybrids-and-alternative-sparsity-dated
title: "MoE+SSM hybrids and alternative sparsity (dated)"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Anthropic", "CISA", "China", "DeepSeek", "Google", "Meta", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Unsloth", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2025-08-05", "2025-12-16", "2026-04", "2026-04-16", "2026-04-24", "2026-05", "2026-05-06", "2026-05-22", "2026-06-01", "2026-07", "2026-07-24", "2026-07-27", "2026-07-28", "2026-08-03", "2026-08-12", "2026-08-16", "2026-08-22", "2026-08-25", "2026-08-28", "2026-08-30", "2026-09", "2026-09-02", "2026-09-07", "2026-09-09", "2026-09-22"]
keywords: ["moe", "agent", "agentic", "agents", "amd", "apache", "attention", "awq", "benchmark", "benchmarks", "blackwell", "claude"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5173, 5282]
section: "9. MoE Architectures"
sha256: c13a501e0d20c784b828b2ccd7edaab9dd278374a2299740f7beb5344fb8f924
---

# MoE+SSM hybrids and alternative sparsity (dated)

### MoE+SSM hybrids and alternative sparsity (dated)

- **Mixtral 8x7B (2023):** 46.7B/12.9B, Apache 2.0 — the original open MoE economics proof: dense-70B-class capability at ~6× inference speed; the AWQ 4-bit vLLM case (90 GB → 26 GB VRAM, 2.1× throughput, −75% TTFT) is the canonical local-serving datum.
- **Nemotron 3.5 Lightning (2026):** 30B/3B MoE + Mamba-2 hybrid with speculative decoding — 235–494 t/s at $0.22/M output (AA, Aug 2026) vs Gemma 4 31B dense at 37–222 t/s and $0.40/M; capability 24 vs 30. The MoE+SSM hybrid is the 2026 efficiency frontier at ~30B total.
- **Llama-4-Maverick-class ultra-sparse MoE (<1% activation density):** EP is 7–12% SLOWER (ROCm measurements) — all-to-all overhead exceeds sparse-compute savings. Sparsity has a floor below which parallelism loses; the 2026 serving rule is ≥3% density for EP to win.
- **MoUE (arXiv 2603.04971, Mar 2026):** reusing a layer-agnostic universal expert pool across layers ("virtual width" from depth) — up to +1.3% over matched MoEs, +4.2% when converting existing MoE checkpoints. Research-scale only; the open question is whether it holds at 2T+ params and 1M context.
- **FLAME (arXiv 2605.09355, May 2026):** adaptive MoE for continual multimodal learning with shared + task-specific experts — the 2026 direction for specialization without full retraining.
- **Phase-aware routing (2026):** routing at environment-step granularity for RL agents (temporal consistency across multi-step plans) — promising for agentic workloads, unproven in production serving stacks.
- **Confidence-gated dynamic expert count (early 2026 work):** easy tokens → fewer experts, suggesting 20–40% compute savings — the one untapped compute axis beyond the active/total ratio itself.

### RAG-operator takeaways

- **$/token economics favor MoE for retrieval-augmented pipelines:** GLM-5.3-Flash at $0.075/M input (promo, through 2026-09-09) to $0.15/M; DeepSeek-V4-Flash cache-hit pricing at fractions of a cent — long retrieved contexts are cheap per token.
- **1M context reduces chunking pressure:** all 2026 flagships accept 1M tokens natively, so "stuff the whole corpus section" beats fine-grained retrieval for many workloads — but KV residency caps concurrent long sessions, making prefix caching and session batching the operative optimizations.
- **Active params predict per-token latency, total params predict fit:** size the GPU fleet on total params (quantized), the latency SLO on active params.
- **Tool-use MoE models are the RAG agent substrate:** Qwen3.6-35B-A3B's 37.0 MCPMark (2× Gemma 4-31B) and DeepSeek-V4's 67% agentic pass rate show sparse models now lead on function calling — the RAG agent loop (retrieve → call tools → synthesize) runs best on exactly these models.
- **Prefix/cache-aware routing:** for RAG serving, the KV cache is the bottleneck, not the experts — invest in prefix caching, HybridKV-style compression, and attention-efficient models rather than denser retrieval.

### Serving-relevant spec snapshot (dated, economics columns)

Total / active per token / context / KV-attention design / license / release:

- gpt-oss-120b: 117B / 5.1B / 128K / standard attention / Apache 2.0 / 2025-08-05
- gpt-oss-20b: 21B / 3.6B / 128K / standard attention / Apache 2.0 / 2025-08-05
- MiMo-V2-Flash: 309B / 15B / 256K / 5:1 SWA(128)/global + sink bias / MIT / 2025-12-16
- Qwen3.6-35B-A3B: 35B / 3B / 262K native (1M w/ YaRN) / 3:1 Gated DeltaNet:gated attention / Apache 2.0 / 2026-04-16
- Gemma 4 26B-A4B: 25.2B / 3.8B / 256K / SWA(1024)+global hybrid, shared KV / Apache 2.0 / ~2026-04
- DeepSeek-V4-Pro: 1.6T / 49B / 1M (384K max output) / NSA lineage / MIT (disputed — see §27 list) / 2026-04-24
- DeepSeek-V4-Flash: 284B / 13B / 1M / NSA + DSA + token-wise compression / MIT (disputed) / 2026-04-24
- MiniMax M3: 428B / 23B / 1M / MSA block-sparse + GQA(64Q/4KV) / MiniMax Community License / 2026-06-01
- Kimi K3: 2.8T / 104B (disputed ~50B — see §27 list) / 1M / 69 KDA + 24 Gated MLA / open-weight / 2026-07-27
- Qwen3.8-Max: 2.4T / 95B / 1M / Gated DeltaNet + gated attention (64Q/4KV) / non-Apache open license / 2026-08-03
- GLM-5.3-Flash: 320B / 18B / 1M (config; 300K cited in one eval — see §27 list) / hybrid sparse+linear + mHC / MIT / 2026-08-25/26
- GLM-5.3 (flagship): 753B / undisclosed / 1M / MoE base from 5.2 / custom glm-5.3 license / 2026-08-28
- MiMo-V2.6-Flash: 309B / 15B / 1M / 5:1 SWA(128)/global / MIT / ~2026-09
- MiMo-V2.6-Pro: 1.02T / 42B / 1M / same family / MIT / ~2026-09
- Nemotron 3.5 Lightning: 30B / 3B / — / MoE + Mamba-2 hybrid / — / 2026 (AA efficiency reference)

### Pricing mechanics notes (dated)

- **2026-05-22:** DeepSeek makes V4-Pro's 75% discount permanent ($0.435/$0.003625/$0.87 input/cached/output).
- **2026-07-24:** legacy `deepseek-chat`/`deepseek-reasoner` aliases retired, routing to V4-Flash during the grace period — model-name continuity matters for cost tracking.
- **2026-08-16:** V4 peak/off-peak tiers introduced (Pro $1.32/$0.66 input; Flash $0.44/$0.22 input) — time-of-day becomes a pricing dimension for MoE serving.
- **2026-09-09:** GLM-5.3-Flash launch promo ends ($0.075/$0.015/$0.25 → $0.15/$0.03/$0.50 input/cached/output).
- **2026-08-22 adoption signal:** Qwen3.8-Max HF downloads 17,386 (bf16) / 1,141 likes; FP8 repo 21,400 downloads — the serving-motivated variant leads.
- **The price ladder mirrors active compute almost monotonically:** Kimi K3 (104B active) and Qwen3.8-Max (95B active) at $3.00/$2.00 per 1M input; 3B-active models (Qwen3.6, Nemotron Lightning) at the bottom. Capability per dollar is dominated by post-training, not architecture.

### RAG session-economics patterns

- **KV residency caps concurrent long-context sessions.** All experts occupy VRAM, so the KV pool left over is smaller than in a same-size dense model — for RAG this means session-level batching and prefix caching are the operative optimizations, not model shrinking.
- **Retrieved-context prefixes are the cost lever.** DeepSeek-V4-Pro cache-hit pricing ($0.003625/M) vs miss ($0.435/M) is a ~120× spread; GLM-5.3-Flash cached $0.03/M vs $0.15/M. Stable system-prompt + retrieval-template prefixes make MoE's cheap long context actually cheap.
- **1M context changes chunking economics:** "stuff the whole corpus section" beats fine-grained retrieval for many workloads — but only where the model's post-training actually uses long context (the attention-side innovations, not MoE sparsity, decide).
- **Tool-use-first model selection for RAG agents:** Qwen3.6-35B-A3B's 37.0 MCPMark (2× Gemma 4-31B) and DeepSeek-V4's 67% agentic pass rate put sparse models ahead on function calling — the retrieve → call tools → synthesize loop runs best on exactly these models.
- **Self-host vs API breakpoint (2026):** self-hosted ballpark $0.25–0.30/MTok on Blackwell (Qwen3-MoE 235B-A22B FP4 on 8× B200; DeepSeek V3.1 FP4+MTP on NVL72) vs API Flash-tier $0.075–$0.15/M input — at Flash-tier prices, API undercuts self-host for most RAG workloads unless residency or data-sovereignty demands otherwise; at flagship tier ($1.40–$3.00/M), self-host breaks even at scale.

### Unresolved discrepancies and unverified claims (do not cite without checking)

- **Kimi K3 active params:** 104B (technical-report-based sources, majority) vs ~50B (one Medium analysis, 2026-07-28). The 16-of-896-experts figure is consistent across sources; the implied per-expert size differs. Treated as 104B (authoritative-leaning) with the dispute noted.
- **DeepSeek-V4 license:** MIT (later mid-2026 coverage) vs Apache 2.0 (April 2026 launch coverage). First-party model card not directly opened in this research pass — verify before citing.
- **GLM-5.3-Flash context:** config.json `max_position_embeddings` = 1,048,576 and OpenRouter listing up to 1,310,720 vs LumaDock model-card eval citing 300K context. The 1M figure is config-grounded; the 300K figure's basis is unclear.
- **MiniMax M3 "matches Claude Sonnet 4.6 on real-world agentic benchmarks":** sourced to Morph's vendor-adjacent model page; no independent benchmark reproduction found. Marked UNVERIFIED.
- **MiMo-V2.6 release details:** come from a single detailed third-party release analysis (OrcaRouter, Sept 2026); Xiaomi first-party announcement not directly opened. The 309B/15B/1M/MIT figures are consistent with the V2-Flash lineage but single-sourced — verify against first-party channels before citing as definitive.
- **MiniMax M3 "7 MTP modules":** appears in community inference notes (MTP/NEXTN weights absent from shipped checkpoints); not confirmed in the official model card. Treat as community-reported.
- **Unsloth "12x faster MoE training"** (Feb 2026), **"MoE training 3–5x faster"** (July 2026 v0.1.481-beta), **"up to 2x faster" MTP-by-default for Qwen3.8-Flash/GLM-5.3-Flash** (Sept 2026 v0.1.805/806-beta), **GLM-5.2 Dynamic GGUF retention figures** (~82%/~76% at 2-bit/~1-bit, ~98% at 4-bit; LinkedIn secondary, vendor-originated): all [VENDOR-REPORTED], no independent reproduction found as of 2026-09-22.
- **Unsloth Studio beta release dates (v0.1.805–808):** presented as "first half of September 2026" from listing recency; the vendor does not stamp exact dates on these beta pages — [UNVERIFIED].
- **PyTorch 2.14 (released 2026-09-02)** silently changed clamp/min/max boundary subgradients (1→0): a genuine training-reproducibility hazard for anyone upgrading mid-experiment. [DIRECTIONAL — independent dev-blog source, not a first-party changelog line.]

## Figures and metrics (continued)

- **Training FLOPs anchor (DeepSeekMoE published ratios):** 16B MoE ≈ 7B dense at ~40% compute; 145B MoE ≈ 67B dense at 28.5% compute — ≈2.5–3.5× compute savings at matched capability.
- **DeepSeek-V3 training footprint:** 16-way pipeline × 64-way expert × ZeRO-1 data parallel; 2048 H800 GPUs; DualPipe; DeepEP all-to-all; FP8 mixed precision. EP dispatch bandwidth ∝ T·k·D per layer — TB-scale per layer at 64-way EP, V3-class.
- **Pretraining token budgets (dated):** MiMo-V2-Flash 27T tokens (2025-12-16); GLM-5.3-Flash 30T multimodal tokens (2026-08); DeepSeek-V4-Pro 33T / V4-Flash 32T (2026-04-24).
- **Aux-loss degradation:** ~23% cited in 2026 analyses of interference gradients from the classic auxiliary loss L = α·N·Σ f_i·P_i — the production field moved to dynamic bias / expert-choice / quantile balancing.
- **Unsloth×NVIDIA component measurements (2026-05-06):** packed-sequence metadata caching +43.3% forward, +5.8% backward, +14.3% per batch (Qwen3-14B QLoRA SFT); double-buffered checkpoint reload +8.4% (8B) / +6.7% (14B) / +4.6% (32B); GPT-OSS bincount routing +23% forward / +13% backward in the targeted path. Vendor headline timeline: "12x faster MoE" (Feb 2026, B200-specific) → "MoE training 3–5x faster" (July 2026) → component numbers (May 2026) as the auditable evidence.
- **Active-compute throughput illustration:** Qwen3.6-35B-A3B (3B active) 61 tok/s vs 27B dense 7 tok/s on the same RTX 5070 (~9×). DeepSeek-V3 671B/37B reads ~5% of weights per token (dense-30B-class throughput on H800).
- **Active/total residency gap:** OLMoE-1B-7B vs OLMo-1B — active params agree within 8.9%, resident memory differs 5.9×; ratio invariant 5.40× at fp32 and int4. Switch-Base-8: all 8 experts receive tokens within 223 tokens (worst-layer coverage 100%).
- **KV-cache reduction factors (attention-side, 2026):** MLA ~5% of LLaMA-3-70B (DeepSeek-V2); MiMo 5:1 SWA/GA ~6× reduction; GLM-5.3-Flash hybrid sparse+linear 4.4× claimed; HybridKV compression up to 7.9×; TurboQuant 3-bit KV ÷6 memory.
- **MoE speculative decoding:** MiMo-V2-Flash accepted length 2.8–3.6 tokens/forward → 2.0–2.6× speedup; MiMo-V2.6-Flash 5-layer DFlash drafter, 7 tokens ahead; MTP-by-default for Qwen3.8-Flash/GLM-5.3-Flash claims up to 2x generation [VENDOR]. ROCm asymmetry: MTP near-useless on MoE at some configs while dense 27B went 45→62 t/s.
- **EP topology guidance (AMD ROCm):** ≤128 concurrent requests → TP=8 +40–86% throughput; ≥512 → DP=8+EP +16–47% (7,114 TPS DeepSeek-R1 at 1024 concurrency); crossover ~256–512; ultra-sparse (<1% activation) → EP 7–12% SLOWER; standard MoE (≥3% density) → EP helps.
- **DeepSpeed comm-opt:** MoE all-reduce penalties 3–12 s in problematic large-scale layouts; multi-rank bucketing + rank placement as fix.
- **$/MTok self-hosted ballpark (mid-2026):** Qwen3-MoE 235B-A22B FP4 on 8× B200 ≈ $0.30/MTok at ~1200 tok/s/GPU; DeepSeek V3.1 FP4+MTP on NVL72 ≈ $0.25/MTok; Llama 3.3 70B FP8 on 4× H100 ≈ $0.40/MTok. AA Aug 2026: Nemotron 3.5 Lightning $0.22/M output at 235–494 t/s vs Gemma 4 31B dense $0.40/M at 37–222 t/s (capability 24 vs 30).
- **API price ladder (2026-09, $/M input / cached / output):** DeepSeek-V4-Pro $0.435/$0.003625/$0.87 (peak tiers $1.32/$0.66 from 2026-08-16); V4-Flash $0.44/$0.22 peak/off-peak input; GLM-5.3-Flash $0.15/$0.03/$0.50 (promo $0.075/$0.015/$0.25 through 2026-09-09); GLM-5.3 flagship $1.40/$4.40; MiniMax M3 $0.30; MiMo-V2-Flash $0.10/$0.30; Qwen3.8-Max $2.00/$0.25 cache/$6.00; Kimi K3 $3.00/$0.30/$15.00.
- **Quantization footprint (2026 MoE fleet):** GLM-5.3-Flash FP8 ~306–331 GiB; MiMo-V2.6-Flash FP8 172.9 GB (65 shards); MiniMax M3 NVFP4 ~245 GB → ~61 GB/GPU at TP4; M3 MXFP8 ~440 GB; gpt-oss MXFP4 native 60.8 GiB (120b) / 12.8 GiB (20b); Kimi K3 per-expert 33.0M params, MXFP4 E8M0 scale per 32 weights; Mixtral 8x7B AWQ 4-bit 90 GB → 26 GB VRAM, 2.1× throughput, −75% TTFT; Qwen3.8-Max FP8 repo 21,400 downloads vs 17,386 bf16 (2026-08-22).
- **Fine-granularity scale ladder (2026):** expert counts 128 (M3, Gemma 4) → 256 (Qwen3.6, MiMo-V2.6, V4-Flash) → 288 (GLM-5.3-Flash) → 384 (V4-Pro) → 512 (Qwen3.8-Max) → 896 (Kimi K3); top-k 4–16; combinatorial capacity C(64,8) ≈ 4.4B vs C(16,2) = 120. Kimi K3 latent routing: hidden 7,168 → latent 3,584, ~halving routed traffic/compute; sparsity ratio ~27:1 (K3), ~25:1 (Qwen3.8-Max), ~12:1 (Qwen3.6), ~6.6:1 (Gemma 4).
- **Scoreboard anchors (with dates):** SWE-bench Verified — M3 80.5% (AA, 2026-09-07), Qwen3.6 73.4% (2026-04), MiMo-V2-Flash 73.4% (2025-12), Gemma 4 17.4% (community); LiveCodeBench — V4-Pro-Max 93.5%, V4-Flash 91.6%, Qwen3.6 80.4% v6, Codeforces 3,052; agentic — V4-Pro 67% vs Sonnet 4.5 47% vs Opus 4.6 Thinking 80%; GPQA Diamond — K3 93.5%, Qwen3.8-Max 92.6%, Qwen3.6 84.1%; MCPMark — Qwen3.6 37.0 vs Gemma 4-31B 18.1; AA Index — K3 57.1 (#4 globally), GLM-5.3 60 vs 5.3-Flash 57 vs 5.2 53.
- **Modest-hardware datums:** Gemma 4 offload hot set ~4.5 GB, ~24 MB/forward from NVMe, 8 GB RTX 3070 viable; Qwen3.6 UD-Q4_K_XL ~21 GB on 24 GB RTX 4090; MacBook Pro 20.9 GB Q4; gpt-oss-20b 12.8 GB [PARTIALLY VERIFIED]; Axolotl `quantize_moe_experts` GLM-4.7-Flash ~127 GiB → ~23 GiB (secondary); Unsloth `save_pretrained_gguf` OOM guard default 0.75; MLX quantized KV caches "up to 74% less prompt memory" [VENDOR].

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

