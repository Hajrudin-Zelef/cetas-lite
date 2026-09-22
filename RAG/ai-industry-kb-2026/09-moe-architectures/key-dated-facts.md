---
id: ai-industry-kb-2026/09-moe-architectures/key-dated-facts
title: "Key dated facts"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Huawei", "Hugging Face", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "Together AI", "Unsloth", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2025-08-05", "2025-12", "2025-12-16", "2026-04-16", "2026-04-24", "2026-06-01", "2026-06-13", "2026-07-16", "2026-07-24", "2026-07-27", "2026-08", "2026-08-03", "2026-08-08", "2026-08-13", "2026-08-22", "2026-08-25", "2026-08-26", "2026-08-28", "2026-09-22"]
keywords: ["agent", "agentic", "amd", "apache", "ascend", "attention", "benchmark", "benchmarks", "blackwell", "claude", "compute", "consumer"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4593, 4756]
section: "9. MoE Architectures"
sha256: a5890288cf7c5abe8e9ad6c4cbe7bcce7e71b91a4f0e2560c0ec5cfefc290bcf
---

# Key dated facts

## Key dated facts

### Release timeline: the open-weight MoE explosion (2025-08 → 2026-09)

| Date | Model | Total / Active | Context | License | Significance |
|---|---|---|---|---|---|
| 2025-08-05 | OpenAI gpt-oss-120b / 20b | 117B / 5.1B ; 21B / 3.6B | 128K | Apache 2.0 | First OpenAI open weights since GPT-2 (2019); MXFP4 native; template every 2026 release follows |
| 2025-12-16 | Xiaomi MiMo-V2-Flash | 309B / 15B | 256K | MIT | First 300B-class open MoE; 150 tok/s via MTP self-speculative decoding; day-0 SGLang support |
| 2026-04-16 | Alibaba Qwen3.6-35B-A3B | 35B / 3B | 262K (1M w/ YaRN) | Apache 2.0 | Hybrid Gated DeltaNet + gated attention; 256 experts (8+1 shared); 73.4% SWE-bench Verified; runs on RTX 4090 |
| 2026-04-24 | DeepSeek V4-Pro / V4-Flash | 1.6T / 49B ; 284B / 13B | 1M | MIT (later sources; April coverage cited Apache 2.0) | Largest open MoE of H1 2026; 33T-token pretraining; launched on Huawei Ascend chips first |
| ~2026-04 | Google Gemma 4 26B-A4B | 25.2B / 3.8B | 256K | Apache 2.0 | Pure-attention MoE (no SSM); 128 experts (8+1 shared); the Western open-MoE anchor of 2026 |
| 2026-06-01 | MiniMax M3 | 428B / 23B | 1M | MiniMax Community License | First open-weight model combining reasoning + agentic coding + native multimodality (text/image/video from step 0) |
| 2026-06-13/17 | Z.ai GLM-5.2 | undisclosed (base reused in GLM-5.3: 753B) | 1M | custom | Coding-plan flagship; coding agent flagship tier |
| 2026-07-16/27 | Moonshot Kimi K3 | 2.8T / 104B | 1M | open-weight | Largest open-weight model ever released (1.56 TB checkpoint); 896 experts (16+2 shared); MoonEP + FlashKDA open-sourced |
| 2026-08-03/08 | Alibaba Qwen3.8-Max | 2.4T / 95B | 1M | non-Apache open license | First Max-class Qwen with open weights; 512 experts (10+1 shared); 92 layers |
| 2026-08-25/26 | Z.ai GLM-5.3-Flash | 320B / 18B | 1M | MIT | First natively multimodal GLM-5; 288 experts (8 active); stealth-launched as "Ox Alpha" on OpenRouter |
| 2026-08-28 | Z.ai GLM-5.3 (flagship) | 753B (MoE base from 5.2) | 1M | custom glm-5.3 license | Post-training upgrade of 5.2 base; text-only |
| ~2026-09 | Xiaomi MiMo-V2.6-Flash / Pro | 309B / 15B ; 1.02T / 42B | 1M | MIT | Native omnimodal (text/image/video/audio); 5:1 SWA/GA hybrid; 5-layer MTP drafter |

- Naming convention, now standard: "A3B"/"A4B" suffixes denote **active** parameters per token (35B-A3B = 35B total, 3B active); Qwen3.5/3.6/3.7/3.8 are successive flagship generations, not point releases of Qwen3; Gemma 4's "E" prefix (E2B/E4B) denotes effective parameters for the dense siblings.

### Release-event granularity: announcement vs weights vs GA

- 2025-08-05: OpenAI gpt-oss-120b/20b — weights and announcement same day (arXiv 2508.10925 model card); the event that set the open-MoE template.
- 2025-12-16: Xiaomi MiMo-V2-Flash — weights and announcement same day; day-0 SGLang support shipped with inference code contributed upstream.
- 2026-04-16: Alibaba Qwen3.6-35B-A3B — weights on Hugging Face, HF model card + Together AI serving pages as the verification anchors; Apache 2.0 from day one.
- ~2026-04: Google Gemma 4 26B-A4B — HF model card + gemma4.dev as anchors; the family (E2B/E4B dense, 26B-A4B MoE, 31B dense) released as one licensing event (all Apache 2.0).
- 2026-04-24: DeepSeek V4 preview (same day as OpenAI GPT-5.5) — two variants announced together; V4-Pro reached GA 2026-08-13 (V4-Pro-0813 checkpoint); legacy aliases retired 2026-07-24.
- 2026-06-01: MiniMax M3 — announcement + weights; served on NVIDIA Blackwell and AMD Instinct via SGLang from launch.
- 2026-06-13/17: Z.ai GLM-5.2 — launch window (two dates in coverage); params undisclosed; base later disclosed through GLM-5.3 (753B).
- 2026-07-16: Moonshot announces Kimi K3 (2.8T) with technical report → weights 2026-07-27 (11 days later); MoonEP, FlashKDA, and AgentEnv open-sourced alongside.
- 2026-08-03: Alibaba Qwen3.8-Max GA → weights 2026-08-08 (5 days later), two HF repos (bf16 + FP8); FP8 out-downloaded bf16 by 2026-08-22 (21,400 vs 17,386) [COMMUNITY].
- 2026-08-25/26: Z.ai GLM-5.3-Flash — after ~1 week as the anonymous "Ox Alpha" atop OpenRouter coding charts, revealed 2026-08-26; weights `zai-org/GLM-5.3-Flash`; MIT.
- 2026-08-28: Z.ai GLM-5.3 flagship weights — post-training upgrade of the 5.2 base, custom license.
- ~2026-09: Xiaomi MiMo-V2.6-Flash/Pro — reported via OrcaRouter release analysis (single source); no Xiaomi first-party announcement directly verified [UNVERIFIED].
- Pattern note: by 2026, GA and weights are decoupled by days, and stealth leaderboard launches ("Ox Alpha") precede formal reveals — release tracking must watch anonymous leaderboards, not only press releases.

### Release cadence statistics (2025-08-05 → 2026-09-22)

- 14 tracked release events in ~13.5 months (gpt-oss 120b/20b, MiMo-V2-Flash, Qwen3.6-35B-A3B, Gemma 4 26B-A4B, DeepSeek-V4 Pro/Flash, MiniMax M3, GLM-5.2, Kimi K3, Qwen3.8-Max, GLM-5.3-Flash, GLM-5.3, MiMo-V2.6-Flash/Pro) — roughly one frontier open MoE per month, accelerating to three in August 2026 alone.
- Total-parameter ceiling: 117B (2025-08) → 309B (2025-12) → 428B (2026-06-01) → 1.6T (2026-04-24, DeepSeek-V4-Pro — note the ceiling jumped before M3 despite the later date) → 2.8T (2026-07-27, Kimi K3). The ceiling roughly doubled every ~4 months.
- Active-parameter ceiling: 5.1B → 15B (2025-12) → 23B (2026-06) → 49B (2026-04-24) → 104B (2026-07-27). Active compute grew ~20× in 11.5 months while sparsity ratios stayed in the 12:1–33:1 band — the field scales capacity and per-token compute together, not sparsity.
- Context ceiling: 128K (2025-08) → 256K (2025-12 / ~2026-04) → 1M native (2026-04-24 onward); 1M is universal across every 2026 flagship.
- Expert-pool ceiling: 32 (2025-08) → 128 (2025-08, the 120b) → 256 (2026-04) → 512 (2026-08) → 896 (2026-07-27, out of chronological order — Kimi K3 leapfrogged Qwen3.8-Max's 512 before its release).
- Lab concentration: of the 12 distinct 2026 releases, 10 are from Chinese labs (DeepSeek, Alibaba ×2, Moonshot, Z.ai ×3, MiniMax, Xiaomi ×2), 1 from Google, 0 from Western labs beyond Google.

### Architectural firsts per release (2025-08 → 2026-09)

- gpt-oss-120b/20b (2025-08-05): first open-weight MoE from OpenAI; native MXFP4; the sparse-MoE + sub-8-bit + permissive-license template; the A-suffix naming convention.
- MiMo-V2-Flash (2025-12-16): first 300B-class open MoE; MTP self-speculative decoding at 150 tok/s; learnable attention-sink bias with 5:1 SWA/GA.
- Qwen3.6-35B-A3B (2026-04-16): first hybrid Gated DeltaNet + gated attention open MoE; YaRN-extended 1M context on a hybrid linear/standard stack; the 3B-active coding MoE.
- Gemma 4 26B-A4B (~2026-04): the pure-attention MoE (no SSM, no linear attention); the expert-offloading case study (8 GB consumer GPU).
- DeepSeek-V4 (2026-04-24): largest open MoE of H1 2026; first flagship launched on Huawei Ascend; reconstructed innovations: sqrt(softplus) routing affinity, deterministic hash routing for the first three blocks, FP4 expert weights.
- MiniMax M3 (2026-06-01): first open release combining reasoning + agentic coding + native multimodality (text/image/video from step 0); MSA block-sparse attention.
- Kimi K3 (2026-07-27 weights): largest open-weight release ever (1.56 TB); Stable LatentMoE latent-space routing; Quantile Balancing; SiTU activation; MXFP4-native 2.8T; MoonEP + FlashKDA co-shipped.
- Qwen3.8-Max (2026-08-08 weights): first Max-class Qwen with open weights; 512-expert pool; 95B active — the highest per-token compute of any Alibaba open release.
- GLM-5.3-Flash (2026-08-26 reveal): first natively multimodal GLM-5; first "Ox Alpha"-style stealth launch on OpenRouter; `glm5_next` hybrid sparse+linear stack with mHC; served entirely on domestic Chinese chips during stealth.
- GLM-5.3 (2026-08-28): the post-training-only flagship upgrade — same 753B base as 5.2, gains from post-training alone; the 2026 case study that post-training, not architecture, decides capability deltas.
- MiMo-V2.6-Flash/Pro (~2026-09): first native omnimodal open MoE (text/image/video/audio in one model); 5-layer DFlash-style MTP drafter.

### Model deep dives

#### gpt-oss-120b / 20b — the 2025 open-weight MoE baseline (released 2025-08-05)

- **Specs (verified via arXiv model card 2508.10925):** gpt-oss-120b: 116.83B total (114.71B MLP + 0.96B attention + 1.16B embed/unembed), 5.13B active per token; 36 layers; 128 experts with 4 active per token (no shared expert); 128K context; 60.8 GiB checkpoint. gpt-oss-20b: 20.91B total, 3.61B active; 24 layers; 32 experts, 4 active; 12.8 GiB checkpoint.
- MXFP4 native quantization; Apache 2.0. First OpenAI open weights since GPT-2 (2019).
- Historical role: established the template every 2026 release follows — sparse MoE + native sub-8-bit quantization + permissive licensing — and the "active parameter" A-suffix naming convention now standard across Qwen/Gemma.

#### Xiaomi MiMo-V2-Flash — first 300B-class open MoE (released 2025-12-16)

- **Specs:** 309B total / 15B activated per token; 48 layers (39 sliding-window + 9 global attention); hidden size 4096; 256 routed experts with 8 activated, **no shared experts**; 128-token sliding window; first block is global attention with dense FFN, interleaved thereafter; learnable attention-sink bias → ~6× KV-cache reduction vs full attention [VENDOR].
- Trained on 27T tokens in FP8 mixed precision; 150 tok/s inference via MTP self-speculative decoding (accepted length 2.8–3.6, 2.0–2.6× effective speedup) [VENDOR]; day-0 SGLang support with inference code contributed upstream.
- 256K context; MIT license; the December 2025 proof that 300B-class MoE could be production-served at 150 tok/s.

#### Alibaba Qwen3.6-35B-A3B — 35B/3B hybrid-attention coding MoE (released 2026-04-16)

- **Specs (verified via HF model card, Together AI, community analyses):** 35B total / ~3B active per token; 40 layers; hidden dim 2048; **256 experts with 8 routed + 1 shared active per token**; expert intermediate dim 512.
- Hybrid attention: 10 repetitions of [3× (Gated DeltaNet → MoE) + 1× (Gated Attention → MoE)] — 30 Gated DeltaNet linear-attention layers + 10 gated attention layers (16 Q heads, 2 KV heads, head dim 256). DeltaNet uses 32 linear attention heads for V and 16 for QK (head dim 128).
- Trained with multi-token prediction (MTP); 262,144 tokens native context, extensible to ~1,010,000 via YaRN RoPE scaling; vision encoder built in (image-text-to-text); Apache 2.0.
- Local inference [COMMUNITY]: Unsloth UD-Q4_K_XL GGUF ~21 GB runs on a single 24 GB GPU (RTX 4090); 20.9 GB Q4 quant runs on MacBook Pro; vLLM ≥ 0.19.0 required (`Qwen3MoeSparseMoeBlock` support); practitioner measurement 61 tok/s on RTX 5070 vs 7 tok/s for a 27B dense model on the same hardware — the canonical active-compute illustration (3B active vs 27B dense, 9× less per-token compute).

#### Google Gemma 4 26B-A4B — 25.2B/3.8B pure-attention MoE (released ~2026-04)

- **Specs (verified via HF model card, community architecture breakdowns):** 25.2B total / 3.8B active per token; 30 layers; hidden dim 2560; **128 experts + 1 shared, 8 routed + 1 shared = 9 active per token**; 32 Q heads, 8 KV heads (GQA); 256K native context; 262K-token vocabulary; ~550M vision encoder (separate, multimodal).
- Attention design: hybrid local/global — sliding-window attention (1024-token window) alternating with global attention; local layers use 8 KV heads (head dim 256), global layers 2 KV heads (head dim 512). Pure-attention MoE with no SSM/recurrence (unlike Qwen3.6's DeltaNet), making the full forward pass parallelizable — prefill scales to full GPU TFLOPS with no sequential bottleneck.
- Expert offloading case study [COMMUNITY]: only ~4.2–4.8 GB must stay hot in VRAM (embeddings ~1.3 GB, per-layer attention+router ~2.4 GB, shared expert ~50 MB); 128 routed experts (~11 GB in Q4_K_M) can sit on NVMe with ~24 MB loaded per forward pass (8 experts × ~3 MB) — servable on an 8 GB RTX 3070 with flash offloading.
- Family context: Gemma 4 ships four sizes — E2B (dense, 2.3B effective), E4B (dense, 4.5B effective), 26B-A4B (MoE), 31B (**dense**, 30.7B) — all Apache 2.0, a licensing shift from earlier Gemma-specific terms. The dense 31B sibling is itself the counterexample to "MoE is the default everywhere" claims.

#### DeepSeek V4 — Pro 1.6T/49B + Flash 284B/13B (released 2026-04-24)

- **Specs (verified via official announcement, model cards, API docs):** V4-Pro: 1.6T total / 49B activated, pre-trained on 33T tokens; V4-Flash: 284B total / 13B activated, 32T tokens. Both: 1M-token native context, up to 384K output tokens, dual thinking/non-thinking modes (high/max/non-think effort levels), text-only.
- Previewed 2026-04-24 (same day as OpenAI GPT-5.5); V4-Pro reached GA 2026-08-13 (V4-Pro-0813 checkpoint); legacy `deepseek-chat`/`deepseek-reasoner` aliases retired 2026-07-24, routing to V4-Flash during grace period.
- Notable industry milestone: launched on **Huawei Ascend chips first** (no NVIDIA required). License: MIT in later sources; early April coverage cited Apache 2.0 — unresolved discrepancy.
- **Reconstructed architecture [PARTIALLY VERIFIED]** (secondary technical analyses and a mirrored snapshot of the V4 technical report; no official DeepSeek report directly verified): V4-Flash — 43 layers, 256 routed experts, 6 active per token + 1 shared; V4-Pro — 61 layers, 384 routed experts, 6 active per token + 1 shared. Every block is MoE; the first three blocks use deterministic token-ID hash routing instead of learned routing. Auxiliary-loss-free load balancing retained from the V3 lineage with a small additional sequence-wise balance loss. Routing affinity changed from sigmoid to `sqrt(softplus(...))`. Routed expert weights reportedly deployed in FP4. V4 replaces the earlier MLA design with CSA/HCA hybrid attention (attention-side detail consolidated in part 09b).

#### MiniMax M3 — 428B/23B natively multimodal MoE (released 2026-06-01)

- **Specs (verified via HF model card, SGLang cookbook, Artificial Analysis):** ~428B total / ~23B activated per token; 60 layers; hidden size 6144; 64 attention heads with `num_key_value_heads=4` (GQA); **128 routed experts with 4 active per token + 1 shared expert**; first 3 layers dense. 1M-token context via MiniMax Sparse Attention (MSA), a block-sparse "lightning indexer" attention (top-k 128-token blocks) keeping decode cost roughly flat in context length — MiniMax reports ~9× prefill and ~15× decode speedup over M2 at 1M context, ~1/20 the per-token compute of the previous generation [VENDOR].
- Multimodality: trained on mixed text, image, and video **from step 0** (not a bolted-on adapter); accepts interleaved text and images; positioned as the first open-weight release combining reasoning, agentic coding, and native multimodality.
- Reasoning/tool use: three reasoning modes via `thinking` param (enabled/adaptive/disabled); chain of thought wrapped in `<mm:think>...</mm:think>`; native XML-namespace tool calling parsed to OpenAI `tool_calls`; served with `--reasoning-parser auto` / `--tool-call-parser auto` in SGLang.
- Serving notes: FP8/MXFP8 and NVFP4 checkpoints exist (`MiniMaxAI/MiniMax-M3-MXFP8`, ~440 GB); served on NVIDIA Blackwell and AMD Instinct via SGLang, and on Hopper (H200) in bfloat16. Community test logs report 7 multi-token-prediction modules in the semi-analysis model card but MTP/NEXTN weights absent from shipped checkpoints (`speculative_enabled: false` mandatory) [COMMUNITY].
- License: MiniMax Community License (custom). Reported to match Claude Sonnet 4.6 on real-world agentic benchmarks via Morph [VENDOR, UNVERIFIED independently].

#### Z.ai GLM-5.2 — coding-plan flagship (released 2026-06-13/17)

- Total parameters undisclosed; the MoE base was reused in GLM-5.3 (753B); 1M context; custom license; coding-plan flagship positioning. Documented here because its base is the one disclosed in GLM-5.3 (see below) — the 5.2→5.3 lineage is the 2026 case study of post-training upgrades on an identical MoE base.

#### Moonshot Kimi K3 — 2.8T/104B frontier MoE (announced 2026-07-16, weights 2026-07-27)

- **Specs (verified via Moonshot technical report, HF model card, multiple analyses):** 2.8 trillion total parameters ("3T-class") / 104B activated per forward pass; **896 routing experts with 16 active per token + 2 shared experts**; 93 layers (1 dense); 69 KDA + 24 Gated MLA layers; 1,048,576-token context; MoonViT-V2 401M vision encoder (native text+image; video documented on first-party API); MXFP4 weights / MXFP8 activations; 1.5609 TB checkpoint — the largest open-weight release ever.
- **Architectural innovations (primary technical report):** **Kimi Delta Attention (KDA)** — up to 6.3× faster decoding [VENDOR]; **Attention Residuals (AttnRes)** — ~25% higher training efficiency [VENDOR]; **Stable LatentMoE** — latent-space routing keeping activation gradients stable across 896 experts; the full hidden state is projected into a narrower latent space before expert dispatch: reported latent width ℓ = 3,584 vs full model width 7,168, roughly **halving routed-expert activation traffic and expert compute**; SiTU (Sigmoid Tanh Unit) activation; per-head Muon optimization; **Quantile Balancing (QB)** replacing fixed-step expert-bias adjustment — each expert's bias is set to the router-score quantile matching its target load, and a single global histogram/all-reduce establishes the quantile thresholds. Reported outcome: ~2.5× more intelligence per unit of compute vs Kimi K2.5 at ~3× the parameter count [VENDOR].
- **Per-expert geometry:** w1/w3 shapes [3072, 3584], w2 [3584, 3072], MXFP4 with E8M0 scale per 32 weights, 33.0M parameters per expert, 2.72T routed parameters total. SiTU-GLU bounds activations (β₁=4, β₂=25, ‖f(x)‖∞ ≤ 100) to stabilize low-precision training; RMSNorm applied before the up-projection.
- Infrastructure open-sourced alongside: **MoonEP** (MoE expert-parallelism communication library attacking cross-node EP transfer bottlenecks), FlashKDA (optimized KDA kernel, 1.72–2.22× faster prefill on H20 [VENDOR]), AgentEnv. Deployment note: Moonshot recommends 64+ accelerators for production self-hosting; runs on AMD GPUs (MI355X serving benchmarks exist) [COMMUNITY].
- **Discrepancy note:** one Medium analysis states "~50B active per token (16 of 896 experts)"; the technical-report-based sources consistently state 104B active. 104B is treated as authoritative-leaning, the ~50B figure as [UNVERIFIED].

#### Alibaba Qwen3.8-Max — 2.4T/95B flagship MoE (GA 2026-08-03, weights 2026-08-08)

- **Specs (verified via HF repos `Qwen/Qwen3.8-2.4T-A95B` and `-FP8`, release coverage):** 2.4T total / 95B activated per token (sparsity ratio ~1:25); 92 layers; **512 experts with 10 routed + 1 shared active**; hidden dim 8192; vocab 248,320; hybrid attention (Gated DeltaNet — 128 linear heads for V / 16 for QK — interleaved with gated attention: 64 Q heads, 4 KV heads, dim 256); 1M-token context (991K max input, 131K max output, 262K reasoning budget); text/image/video input.
- First Max-class Qwen with downloadable weights; non-Apache open license.
- Benchmarks [VENDOR unless noted]: 93.0 PaperBench; GPQA Diamond 92.6%; SWE-bench Pro 67.7%; FrontierSWE 73.5%; vendor table (~40 rows, 20 methodology footnotes) claims competitiveness with OpenAI/Anthropic flagships — treat cross-vendor comparisons as [DIRECTIONAL].
- Adoption (2026-08-22): 17,386 downloads / 1,141 likes (bf16), 21,400 downloads (FP8) — FP8 leading, indicating serving-motivated fetchers [COMMUNITY].
- Comparison point: activates ~95B of 2.4T vs Kimi K3's 104B of 2.8T — Alibaba fires nearly 2× the per-token compute of Moonshot's flagship at similar total scale.

#### Z.ai GLM-5.3-Flash — 320B/18B natively multimodal MoE (released 2026-08-25/26)

- **Specs (verified via Z.ai announcement, model card config, independent analyses):** 320B total (~321B in repo metadata) / 18B active per token; 45 layers; hidden size 4096; **288 routed experts with 8 active per token (no shared expert reported)**; architecture registered as `glm5_next` — hybrid sparse + linear attention stack, the first for the GLM series; Manifold-Constrained Hyper-Connections (mHC) for scaling efficiency. Pre-trained on a 30T-token multimodal corpus.
- Context/modalities: 1,048,576-token context (from `max_position_embeddings` in config.json; OpenRouter lists up to 1,310,720; LumaDock model-card evals cite 300K — unresolved); text+image+video input, text output; reasoning always on (`reasoning_effort` low/high/max, default max).
- Launch: spent its first week topping OpenRouter coding charts anonymously as **"Ox Alpha"** (served entirely on domestically produced Chinese AI chips), revealed 2026-08-26; weights `zai-org/GLM-5.3-Flash` on Hugging Face; MIT license.
- Checkpoints: default FP8 (~306–331 GiB on disk); BF16 repo also published.
- Economics signal [VENDOR/community pricing pages]: beats GLM-5.2 across benchmarks at ~1/10 the price — the full $/MTok ladder is consolidated in part 09b.

#### Z.ai GLM-5.3 (flagship) — 753B post-training upgrade (weights 2026-08-28)

- Reuses GLM-5.2's 753B MoE base with post-training gains; text-only; 1M context; custom `glm-5.3` license (not MIT); ~756 GB FP8 weights.
- The flagship-vs-Flash comparison is the cleanest 2026 illustration of the MoE efficiency trade inside one lab lineage: a few points of measured intelligence for ~9× the serving weight cost — details in part 09b.

#### Xiaomi MiMo-V2.6-Flash / Pro — 309B/15B omnimodal MoE (~2026-09)

- **MiMo-V2.6-Flash (verified via OrcaRouter release analysis, Sept 2026):** 309B total / 15B activated per token; 48 layers (39 sliding-window + 9 global attention); hidden size 4096; 256 routed experts with 8 activated, no shared experts; 128-token sliding window; first block is global attention with dense FFN, interleaved thereafter. Native omnimodal: 681M-param MiMo ViT vision encoder, 308M audio tokenizer, 127M audio patch encoder — text, image, video, and audio enter the same model. 1M-token claimed context. Five-layer MTP drafter (DFlash-style, 7 tokens ahead per forward pass). Published weights: 172.9 GB across 65 shards in FP8 (e4m3). MIT license.
- **MiMo-V2.6-Pro:** 1.02T total / 42B activated; MIT license.
- **Provenance caveat:** the "V2.6" details come from a single detailed third-party release analysis (OrcaRouter, Sept 2026); a Xiaomi first-party announcement was not directly opened in this research pass — the 309B/15B/1M/MIT figures are consistent with the V2-Flash lineage but single-sourced: [UNVERIFIED as definitive; verify against first-party channels before citing].

### Vendor-claims ledger: what is vendor-sourced vs independently verified

- [VENDOR] MiniMax M3 "~9× prefill / ~15× decode over M2 at 1M context; ~1/20 per-token compute of the previous generation" (MSA claims); the Morph-sourced "matches Claude Sonnet 4.6 on real-world agentic benchmarks" — [UNVERIFIED independently].
- [VENDOR] MiMo-V2-Flash "150 tok/s via MTP self-speculative decoding; accepted length 2.8–3.6; 2.0–2.6× effective speedup; learnable sink bias → ~6× KV reduction; trained on 27T tokens in FP8"; the 150 tok/s figure is independently plausible given MTP self-speculation but no third-party reproduction was found in this pass — [DIRECTIONAL].
- [VENDOR] Kimi K3 "up to 6.3× faster decoding (KDA); ~25% higher training efficiency (AttnRes); 1.72–2.22× prefill on H20 (FlashKDA); ~2.5× scaling efficiency over K2"; the technical-report mechanisms (latent width, QB, MXFP4 geometry) are primary-source and independently inspected by multiple community analyses — [COMMUNITY]-confirmed as present in the weights/config.
- [VENDOR] Qwen3.8-Max benchmark table (~40 rows, 20 methodology footnotes; 93.0 PaperBench, 92.6% GPQA Diamond, 67.7% SWE-bench Pro, 73.5% FrontierSWE) — vendor-chosen harnesses; treat cross-vendor comparisons as [DIRECTIONAL]; third-party anchors (Artificial Analysis, community reproductions) are the trustworthy cross-model numbers.
- [VENDOR] GLM-5.3-Flash "hybrid sparse+linear attention → 4.4× smaller KV cache"; "beats GLM-5.2 at ~1/10 the price" — directional until independently measured.
- [COMMUNITY] Qwen3.6 61 tok/s vs 7 tok/s dense on RTX 5070; Gemma 4 expert-offloading profile (~4.2–4.8 GB hot set); Qwen3.8-Max FP8-out-downloading-bf16 adoption signal — practitioner-measured, hardware-specific, not lab-verified.
- [UNVERIFIED] The ~50B-active figure for Kimi K3 (one Medium analysis) against the 104B technical-report consensus; the DeepSeek-V4 secondary architecture reconstruction (treat every number as [PARTIALLY VERIFIED]); MiMo-V2.6 figures (single source); DeepSeek-V4 license (MIT vs Apache 2.0); GLM-5.3-Flash 300K-context mention; M3's 7 MTP modules (community inference notes, absent from official card).

### Specification anchors: where each model's specs were verified

- gpt-oss-120b/20b: arXiv 2508.10925 model card (parameter breakdown 114.71B MLP / 0.96B attention / 1.16B embed-unembed for the 120b; checkpoint sizes 60.8 GiB / 12.8 GiB).
- Qwen3.6-35B-A3B: Hugging Face model card + Together AI model page (`together.ai/models/qwen3-6-35b-a3b-fp8`) + community GGUF/quant tables (knightli.com, allthings.how).
- Gemma 4 26B-A4B: HF model card + gemma4.dev model page + community offload architecture doc (alexchen31337/gemma4-moe-offload).
- MiniMax M3: HF model card + SGLang cookbook (solrex/sglang fork) + Artificial Analysis intelligence-index entries + Morph model page.
- DeepSeek-V4: official announcement + API docs for benchmarks/pricing; architecture numbers only via secondary reconstructions (unofficial HF technical-report mirror + community case studies) — [PARTIALLY VERIFIED].
- Kimi K3: Moonshot primary technical report (`MoonshotAI/Kimi-K3` repo PDF) + HF model card + multiple independent community architecture inspections.
- Qwen3.8-Max: HF repos `Qwen/Qwen3.8-2.4T-A95B` and `-FP8` (download/like counts as of 2026-08-22) + release coverage.
- GLM-5.3-Flash: Z.ai announcement + model card `config.json` (`max_position_embeddings` = 1,048,576) + HF repo `zai-org/GLM-5.3-Flash` + OpenRouter "Ox Alpha" chart history.
- MiMo-V2-Flash: Xiaomi announcement + upstream SGLang inference code; MiMo-V2.6-Flash/Pro: OrcaRouter release analysis only — [UNVERIFIED as definitive].

