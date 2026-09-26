---
id: ai-industry-kb-2026/09-moe-architectures/vendor-claims-ledger-what-is-vendor-sourced-vs-independently
title: "Vendor-claims ledger: what is vendor-sourced vs independently verified"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Anthropic", "China", "Hugging Face", "Moonshot", "OpenAI", "OpenRouter", "Xiaomi", "Z.ai"]
dates: ["2026-08-03", "2026-08-08", "2026-08-22", "2026-08-25", "2026-08-26", "2026-08-28"]
keywords: ["amd", "apache", "attention", "benchmarks", "compute", "cost", "embeddings", "fp8", "glm", "gpus", "kimi", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4702, 4736]
section: "9. MoE Architectures"
sha256: 3d963e861f4dae20c1bc7af3ac80fe8013e125e9595c0517bff4b3e8efdc03cd
---

# Vendor-claims ledger: what is vendor-sourced vs independently verified

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

