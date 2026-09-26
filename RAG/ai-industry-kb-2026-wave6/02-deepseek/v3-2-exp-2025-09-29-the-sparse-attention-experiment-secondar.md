---
id: ai-industry-kb-2026-wave6/02-deepseek/v3-2-exp-2025-09-29-the-sparse-attention-experiment-secondar
title: "V3.2-Exp — 2025-09-29, the sparse-attention experiment [SECONDARY]"
domain: deepseek
role: deep-dive
task: actor-profile
actors: ["AMD", "Anthropic", "China", "DeepSeek", "Fireworks AI", "Google", "Huawei", "OpenAI", "OpenRouter", "SGLang", "vLLM", "xAI"]
dates: ["2025-09-29", "2025-11", "2025-12-01", "2026-04-24", "2026-07-31", "2026-08-07", "2026-09-04"]
keywords: ["attention", "agentic", "alignment", "amd", "ascend", "benchmark", "claude", "context window", "cost", "deepseek", "distribution", "fine-tuning"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [625, 656]
section: "§2. DeepSeek"
delta_of: ai-industry-kb-2026
sha256: 811bda385c055e8e40faa82c4fa84432cb08e3a93ba4200300e166da95b492aa
---

# V3.2-Exp — 2025-09-29, the sparse-attention experiment [SECONDARY]

### V3.2-Exp — 2025-09-29, the sparse-attention experiment [SECONDARY]
- **2025-09-29** — DeepSeek released **V3.2-Exp** (TechCrunch): an experimental model whose single architectural change was **DeepSeek Sparse Attention** — a "lightning indexer" that prioritizes excerpts from the context window plus a "fine-grained token selection system" that picks specific tokens from those excerpts into a limited attention window. [SECONDARY]
- The research framing matters: one researcher called it "a controlled experiment release — the kind we need more of in AI," because DeepSeek **swapped one component (attention) and held everything else constant**, letting the community study sparsity's effect on large-scale training as a testbed. [SECONDARY]
- Pricing: **$0.028 per million input tokens** — a 50% cut from V3.1-Terminus — "even when approaching the context limit of 128,000 tokens." Community reaction: "half the cost for the same results," with developers reporting low latency and high throughput via OpenRouter, crediting the sparse kernels. [SECONDARY]
- The official **V3.2** followed on **2025-12-01** alongside a specialized variant **DeepSeek-V3.2-Speciale**; one syndicated launch write-up cites official V3.2 input pricing **as low as $0.07/M** — consistent with the $0.14/$0.28 vs $0.27/$0.40 dispute above being a tier/cache artifact, not a single number. [SECONDARY]
- Benchmark deltas vs V3.1-Terminus (independent review): **+0.9 on AIME 2025**, **+75 points on Codeforces** — the consolidated RL fine-tuning (merging reasoning, tool-use, alignment) gave V3.2-Exp a more balanced skillset. [SECONDARY]
- Open tooling at launch: **TileLang** kernels (readability), **DeepGEMM** (indexer logit kernels), **FlashMLA** (sparse attention CUDA kernels), plus **vLLM and SGLang support on day one** — cited as evidence DeepSeek was "cultivating broad community integration rather than locking down distribution." [SECONDARY]
- Self-hosting economics: the 685B model needs **~700 GB VRAM at FP8** (8–10× H100 80GB); 4-bit quantization brings it to **~386 GB** (5–6× H100). API price comparisons at launch: GPT-5 Standard $1.25/$10, Claude Opus 4.1 $15/$75, DeepSeek V3.2-Exp $0.028 input — a 45×–500× gap. [SECONDARY]
- Competitive context: V3.2-Exp shipped two months before a packed November 2025 window (GPT-5.1, Grok 4.1, Gemini 3 Pro, Claude Opus 4.5 within six days) — DeepSeek's December official release was read as proof Chinese labs could match frontier pace despite hardware restrictions. [SECONDARY]
- Vendor-chart results reported for **V3.2-Exp** (label as vendor claims, not independent): MMLU-Pro **85.0**, AIME 2025 **89.3**, GPQA-Diamond **79.9** (vs **80.7** for V3.1-Terminus), Codeforces **2121**, BrowseComp **40.1**. Note the inverted V3.2 < V3.1-Terminus GPQA result — one secondary read is that V3.1-Terminus was the stronger release on that axis, which tempers the "successor always wins" reading. [SECONDARY]
- **Pricing contradiction**: vendor/third-party listings conflict between **$0.14 input / $0.28 output** and **$0.27 input / $0.40 output** per million tokens. Likely explanations are first-party-vs-third-party pricing, cache-hit accounting, or date/cache mismatches — resolve only against DeepSeek's own pricing page or changelog. [SECONDARY]

### V4 preview — 2026-04-24 parameter counts [SECONDARY]
- The 2026-04-24 preview pinned the V4 family to **1.6T total / 49B active** for **V4 Pro** and **284B total / ~13B active** for **V4 Flash**, with **1M context / 384K output** and MIT licensing. [SECONDARY]
- V4 Flash-0731 (2026-07-31) kept the same 284B/~13B footprint but added **hybrid compressed attention** and **DSpark speculative decoding**. [SECONDARY]
- Third-party write-ups (morphllm) fill in: V4 Pro trained on **2 trillion tokens**; V4 Flash positioned as the efficient tier with **2×–7× acceleration** claims; DeepSeek positioned V4 as a unified family spanning the frontier (Pro) and cost (Flash) ends. [SECONDARY]
- The V4 preview's 1M-token context (Pro) and 384K (Flash) were the context ceiling for the family until V4.1-Flash's 1M/256K split with off-peak pricing. [SECONDARY]

### V4 Flash-0731 — evidence-brief details [SECONDARY]
- **2026-07-31** — official V4 Flash release, superseding the April/June preview. **Identical architecture to the preview** — 284B total / ~13B active MoE, 1M context — but the HF repo reads **304B params** (base + drafter); DeepSeek's changelog says it "was only re-post-trained": every gain is downstream of pretraining. MIT weights. [SECONDARY]
- New API surface: **native Responses API** ("specifically adapted for Codex") and an **`/anthropic` endpoint**. Pricing unchanged: **$0.14/$0.28 per 1M**; first-party cache hit **$0.0028/M** (~98% discount; Fireworks cache read $0.028, 10× worse). [SECONDARY]
- Fireworks route: `fireworks/accounts/fireworks/models/deepseek-v4-flash-0731` — 1M ctx, 384K max output, thinking levels minimal/low/medium/high/max, text-only. The undated `deepseek-v4-flash` alias was silently re-pointed at this checkpoint on release day. [SECONDARY]
- One community evidence brief (mattrobenolt, 2026-08-07): architecture pinned to 284B / **21B active** (vs the ~13B vendor figure — active-count discrepancy, flag it), 256 routed experts top-6 + shared, hash routing on first 3 layers, MLA attention (latent KV), mHC hyper-connections, DSA lightning indexer + KV compressor, DSpark speculative module attached. [COMMUNITY]

### V4 architecture — hybrid attention, mHC, Muon [SECONDARY]
- **Three-layer hybrid attention** (per community technical write-ups): **CSA** (Compressed Sparse Attention, m=4, DSA — queries attend only to top-k of compressed KV); **HCA** (Heavily Compressed Attention, up to 128× compression, dense — coarse global context); **SWA** (Sliding Window, 128 — local fine-grained dependencies). Outputs integrated via an undisclosed gating mechanism. [SECONDARY]
- **Manifold-Constrained Hyper-Connections (mHC)**: residual mapping constrained onto the manifold of **doubly stochastic matrices** (rows/columns sum to 1, elements ≥ 0); spectral norm ≤ 1 — suppresses vanishing/exploding gradients in ultra-deep 1M-context training, like V3's auxiliary-loss-free load balancing for depth scaling. [SECONDARY]
- Claimed efficiency vs V3.2 at 1M context: **~27% inference FLOPs / ~10% KV cache**. **Muon optimizer**; **32T+ pre-training tokens**. Three reasoning modes: **Non-think / Think High / Think Max** (≥384K context for Think Max). Recommended generation: temperature=1.0, top_p=1.0. [SECONDARY]
- Checkpoint format: Instruct ships as **FP4 MoE experts + FP8 attention/dense** (one mixed-precision checkpoint covers every FP4-capable GPU); `*-Base` repos ship pure FP8 and are for further pre-training only. [SECONDARY]
- Verified serving matrix (per one community cookbook): 4×GB300 (TP=4); B200/B300/H200 FP4 (TP=8); GB200 2-node (TP=8); H100 2-node (TP=16); **MI355X** (AMD). vLLM-Ascend docs: w8a8-mtp quantized variant runs on 1× Atlas 800 A3 (128GB×8) or A2 (64GB×8) node. [COMMUNITY]
- Community serving extreme: **Config-I hybrid GGUF** (TQ3_1S attn/down + Q2_0 experts, 2.88 bpw, 95 GiB) on Metal (M5 Max 128GB) + CUDA (DGX Spark GB10) — "a 284B agentic model that genuinely fits and runs on one 128GB box at 2.88 bpw." One DGX Spark recipe (EXL3 3.0 bpw, NVFP4-DS-MLA 432-byte KV records, DSpark K5, 384K context, ~107 GB checkpoint) froze its benchmarked config 2026-09-04. [COMMUNITY]

