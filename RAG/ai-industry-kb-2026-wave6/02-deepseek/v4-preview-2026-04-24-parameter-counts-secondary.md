---
id: ai-industry-kb-2026-wave6/02-deepseek/v4-preview-2026-04-24-parameter-counts-secondary
title: "V4 preview — 2026-04-24 parameter counts [SECONDARY]"
domain: deepseek
role: deep-dive
task: actor-profile
actors: ["AMD", "Anthropic", "Baseten", "DeepSeek", "Fireworks AI", "Huawei", "Meta", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Z.ai", "vLLM"]
dates: ["2023-11", "2026-04-24", "2026-07-31", "2026-08", "2026-08-07", "2026-08-12", "2026-08-13", "2026-09-04", "2026-09-08", "2026-09-10", "2026-09-11"]
keywords: ["agent", "agentic", "amd", "ascend", "attention", "benchmark", "benchmarks", "compute", "context window", "cost", "decode", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [637, 707]
section: "§2. DeepSeek"
delta_of: ai-industry-kb-2026
sha256: 1816f5f562ac6f35f860b6694965f93c017c33595b6950aa46170e6bb4692207
---

# V4 preview — 2026-04-24 parameter counts [SECONDARY]

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

### V4-Pro-0813 — release-state and pricing details [SECONDARY]
- **Release-state discrepancy**: the `deepseek-chat`/`deepseek-reasoner` API surface began resolving to the new checkpoint around **2026-08-12**, but the official checkpoint identifier and GA date are **2026-08-13** (0813). Corpus entries must date events by the source that claims them: "API began routing 08-12" vs "0813 GA." [SECONDARY]
- Spec: 1.6T MoE with **49B active per token**; the DSpark-augmented checkpoint adds spec tokens toward a **~1.7T physical checkpoint size** — the 1.6T/1.7T pair is a vendor/compute-vs-physical accounting nuance, not two models. Throughput reported at **~78.1 tokens/s**. [SECONDARY]
- Official API pricing after the 0813 release: **$1.32 input / $3.96 output** per million, with **$0.044 cached-input** and **half-price off-peak** during **01:00–04:00 and 06:00–10:00 UTC**. [SECONDARY]
- The 0813 GA landed alongside Moonshot's **Kimi K3** — one secondary comparison (webpronews) framed the two as the August 2026 frontier contenders. V4-Pro was DeepSeek's flagship for exactly four weeks (0813 → 0914 retirement), a tenure so short the Kimi K3 comparison was the only major head-to-head it ever got. [DIRECTIONAL]
- Vendor-claimed agent benchmarks: **Terminal-Bench 2.1: 87.9**, **CyberGym: 83.3**, **DeepSWE: 62.7** — all vendor claims until independently replicated. [SECONDARY]
- **Launch coverage** (Reuters, 2026-08-13): V4-Pro-0813 priced at **$1.32 input / $3.96 output** — ~9× input and ~14× output vs V4 Flash ($0.14/$0.28) — as DeepSeek "seeks to turn stronger benchmark performance into a premium flagship offering." The piece notes the awkward backstory: V4 Flash-0731 had **unexpectedly outperformed the April preview V4 Pro** in independent tests, forcing the official Pro to prove it was actually the flagship. [SECONDARY]
- **Pricing discrepancy**: WCCFTech and others quote **$0.435 input / $0.87 output** — ~3× below the Reuters $1.32/$3.96. The two sets are best read as different pricing tiers/windows (peak vs off-peak vs cached), not a correction; the corpus must not pick one as "the" price without the tier label. [SECONDARY]
- **Competitive context** (WCCFTech): OpenAI had just launched a price war discounting GPT-5.6 Luna 80% ($1→$0.20 input, $6→$1.20 output); DeepSeek's $0.14/$0.28 V4-Flash-0731 "eviscerated any comparative price advantage." V4-Pro-0813's Terminal-Bench 87.9 sits within a decimal point of **Fable 5's 88.0** at roughly **57× lower output-token cost**. DeepSeek was **2nd only to Anthropic in July token volume** (before V4-Flash-0731) — "I'd imagine it will be 1st for August." [SECONDARY]
- **DeepLearning.ai The Batch** (the strongest secondary): V4-Pro-0813 — 1M input / 384K output at 78.1 tok/s; 1.6T/49B (+1.7T with DSpark); adjustable reasoning (none/low/high/max, defaults high); tool calls; context caching; **53 on AA Intelligence Index (3rd among open weights)**; 10th of 115 on Arena WebDev; MIT weights; **open-source agent harness in developer preview**; API prices **increased for all models** with the release. Undisclosed: new training data/methods, knowledge cutoff. [SECONDARY]
- **Launch mechanics** (Simon Willison, via community wiki): API-only on Aug 12, **no official announcement page** — OpenRouter was the access point; benchmarks released to the **Official DeepSeek WeChat Group**, copied to a Reddit post (deleted as "low-effort"), then an **ASCII-art table on Hacker News** — "consistent with DeepSeek's pattern of informal checkpoint announcements." Willison observed **markedly different output across reasoning levels** (low/medium/high) on his pelican SVG test — variance he "had not noticed with any other model." [SECONDARY]
- **CONTRADICTION**: Artificial Analysis Index readings for V4-Pro-0813 conflict between **53** and **36** across sources — nearly certainly different AA Index revisions/dates. Per corpus policy (never mix AA Index versions), neither replaces the other; document both with their version/date or omit until confirmed. [SECONDARY]

### V4.1-Flash — architecture internals [SECONDARY]
- Depth split: **40 transformer layers = 20-layer causal encoder + 20-layer decoder** — the causal encoder-decoder structure the wave3 verification track labeled "CED." [SECONDARY]
- MoE: **1 shared + 384 routed experts, 6 routed active per token** — a much wider expert pool than V4 Flash with a modest active count, pushing specialization over per-token compute. [SECONDARY]
- **Size contradiction**: vendor headline = **552B compute backbone**; community counting adds a **196B Engram conditional memory** module for a **748B physical-checkpoint** number. This is a real accounting contradiction (compute vs physical), not two interchangeable sizes — the corpus must keep the two figures on separate ledger lines. [SECONDARY]
- Active parameters differ by phase: **8B active in prefill, 16B in decode**. [SECONDARY]
- Attention: **CSA2 (compressed sparse attention v2)**. KV cache: **FP4 at ~890 bytes/token — under 1GB for 1M tokens** — the engineering fact that makes the 1M context window deployable. [SECONDARY]
- Context/output: native **1M input, 384K output**; image input reported as possible (not firmly confirmed). [SECONDARY]
- Pricing (peak): **$0.30 input / $1.20 output**; off-peak **$0.15 / $0.60**; cache **$0.006 peak / $0.003 off-peak** — among the cheapest frontier-grade agent APIs in the corpus. [SECONDARY]
- Vendor-claimed benchmarks (label as claims): GPQA-Diamond **90.9**, Terminal-Bench 2.1 **90.6**, DeepSWE **74.2**, HLE **36.8**, CyberGym **88.1**. [SECONDARY]

### V4.1-Flash — GA 2026-09-10, ground-up redesign [SECONDARY]
- **2026-09-08** — two-day beta preview under the temporary model ID `deepseek-v4.1-flash-expires-on-0910`, capped at **20 concurrent requests**; the beta window closed and **GA landed 2026-09-10**. [SECONDARY]
- Scale ledger: **552B backbone + 196B Engram** (conditional memory module accessed sparsely via N-gram-hashing token lookup) — **CONTRADICTION**: secondary write-ups compute the total as **~748B** (552+196) or **~763B**. Both are secondary; the 552B/196B split is consistent across sources, only the summed headline differs. [SECONDARY]
- **Causal Encoder-Decoder (CED)**: instead of recomputing KV cache per decoder layer, the decoder's global KV is **projected from the encoder's final hidden states** — roughly **halving prefill computation**, which matters most at long context and high cache-miss rates. [SECONDARY]
- **CSA2 three modes**: Full (compute KV and select attention indices from scratch), Reindex (reuse shared KV, pick fresh indices), Reuse (borrow both KV and indices from a prior layer). Cross-layer sharing is the main reason the cache shrinks. One independent reconstruction: main KV entry = 512 channels × 4 bits (256 B) + 32 one-byte scales = 288 B; indexer-K = 128 × 4 bits (64 B) + 4 scales = 68 B; encoder 3 Full layers at 2:1 compression, decoder 1 Full layer at 1:1 → (3×288 + 3×68)/2 + 288 + 68 = **890 bytes/token**. [SECONDARY]
- **SWA Bounded Replay**: instead of persisting full sliding-window attention state to SSD, reconstruct by replaying only the most recent window — persistent cache footprint ≈ **1/8 of V4 Flash**. [SECONDARY]
- **DSpark speculative decoding**: a lightweight **3-block drafter** generating **5 draft tokens per forward pass**, trained separately after pre-training and kept aligned during post-training; it **replaces the MTP module** from DeepSeek V3. [SECONDARY]
- **FP4 KV caching in MXFP4 (E2M1)** format; the 890 B/token figure is **~1/4 of V4 Flash** and DeepSeek's headline **437× vs DeepSeek V1** — but V1 shipped November 2023 with a 4K context window, pre-MLA. The honest comparison from the same sources: **4× vs V4 Flash** (the model it actually replaces); full-cache comparison is V1 1.59 GB vs V4.1-Flash 0.93 GB = **256× the context in 41% less memory**. The corpus should cite the 4×/256× figures, not the 437×. [SECONDARY]
- Serving economics (community analysis): **510.3 GB across 48 files**, ~313 GB resident once the 196.9B-parameter engram tables move to host memory; on **8× NVIDIA H200 SXM (141 GB each)**, ~815 GB free → **~915 concurrent users each holding a full million tokens**, vs 40 for DeepSeek V3 on the same node. [COMMUNITY]
- Training: **from scratch on 45T tokens**, context extended to 1M at the **34T-token mark**. Post-training: SFT, RL, on-policy distillation; the agentic gains are credited to **large-scale automated synthesis of agent tasks and environments** — synthetic agent-task generation at scale. **Reasoning effort continuously controllable from 1 to 100** per request. [SECONDARY]
- Observed serving (third-party): **215 output tokens/s**, p50 TTFT 842 ms, p95 2.40 s, 1.2% error rate. **MIT license** — downloadable, self-hostable, fine-tunable, commercial use without royalties or usage restrictions. [SECONDARY]
- Community extreme: a single **DGX Spark (GB10)** recipe serves the full FP4 checkpoint — 73.8 GB hot-expert set (25.6% of 15,360 routed FP4 experts) resident, remainder streamed off NVMe, DSpark worth ~1.5× — full checkpoint quality, no requantization. [COMMUNITY]

### V4.1-Flash vendor benchmark table — full three-generation comparison [SECONDARY]
- KDnuggets published DeepSeek's reported three-generation table (all vendor claims):
  - **DeepSWE v1.1**: V4-Flash 54.4 → V4-Pro 62.7 → **V4.1-Flash 74.2**
  - **Terminal-Bench 2.1**: 82.7 → 87.9 → **90.6**
  - **CyberGym**: 76.7 → 83.3 → **88.1**
  - **AutomationBench**: 37.7 → 43.2 → **54.8**
  - **Agent's Last Exam**: 25.2 → 25.7 → **31.8**
- Note: these index numbers must not be compared to AA Index figures for these models — different instruments. [SECONDARY]

### V4.1-Flash — launch week ecosystem (2026-09-10/11) [SECONDARY]
- **Baseten** added V4.1-Flash to its Model APIs on **2026-09-11** — the day after launch — describing it as DeepSeek's **third open-weight flash release of 2026** and "the only model of its scale" using CED; support for Baseten's Loops training product was announced as coming soon. [SECONDARY]
- **First non-experimental native image input**: V4.1-Flash is DeepSeek's first non-experimental model with native image understanding — previously limited to the experimental **V4-Flash-Vision-Exp**. Card figures: **Chartography 78.9 vs 64.3**, **ZeroBench 49 vs 35** (new vs experimental). [SECONDARY]
- Launch-day cross-lab comparisons (vendor card): **GPQA Diamond 90.9**, **Codeforces 3471**, **HLE with tools 63.9**; **CyberGym 88.1 vs 84.5 for GPT-5.6 Sol and GLM 5.3** (per one secondary write-up). [SECONDARY]
- The size dispute, reconciled: one community analysis (r/LocalLLaMA, cited by OrcaRouter) reads the checkpoint as **748B total, ~510 GB in FP8** — the reconciliation offered is that DeepSeek's "552B" counts only the compute backbone while the download also carries a large sparsely-accessed memory table looked up rather than computed through. Treat 552B as vendor-reported, 748B as an unrefuted community reading. [COMMUNITY]

