---
id: ai-industry-kb-2026/03-chinese-ai-labs-deep-file-glm-kimi-minimax/benchmarks-all-figures-vendor-unless-otherwise-labeled
title: "Benchmarks (all figures [VENDOR] unless otherwise labeled)"
domain: chinese-ai-labs-deep-file-glm-kimi-minimax
role: deep-dive
task: benchmark
actors: ["Anthropic", "ByteDance", "EU", "Hugging Face", "MiniMax", "Moonshot", "United States", "Z.ai"]
dates: []
keywords: ["benchmark", "benchmarks", "agent", "agents", "attention", "claude", "compute", "cost", "fine-tuning", "glm", "int4", "kimi"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1097, 1135]
section: "3. Chinese AI Labs — Deep File: GLM, Kimi, MiniMax"
sha256: eef74870ce18c18a255adef180da9668c98f3d8cefad6341b662d15e3a660411
---

# Benchmarks (all figures [VENDOR] unless otherwise labeled)

- **Scale**: 1T total / 32B active, verified across all sources. Successor to Kimi K2.5 on the same trillion-parameter MoE backbone.
- **Expert layout**: 384 total experts, **8 routed per token + 1 shared expert** (always active); 61 layers (1 dense); MLA attention (7,168 hidden); SwiGLU; 160K vocabulary; native INT4 quantization.
- **Context**: 256K tokens / 262,144 exact (256K is the shorthand) with automatic context compression on long sessions; up to **98,304 output tokens** for reasoning tasks.
- **License**: **Modified MIT** open weights on Hugging Face (`moonshotai/Kimi-K2.6`) — permits commercial use, fine-tuning, and resale with no royalties; operators above **100M MAU or $20M monthly revenue** must display "Kimi K2.6" in the product UI — a gating threshold far more permissive than community-license norms.
- **Modalities**: native multimodal — text, image, **video** input via the **MoonViT vision encoder (400M parameters)**; text output.
- **Variants / agent modes**: Instant (low latency), Thinking (default), Agent, Agent Swarm. Headline capability: up to **300 sub-agents / 4,000+ coordinated steps** per task (3× the K2.5 cap).
- **Agent-mode ladder detail** (secondary launch coverage): Instant trades depth for latency on quick queries; Thinking is the default reasoning mode; Agent runs tool-using single-agent trajectories; **Agent Swarm** coordinates up to 300 sub-agents across 4,000+ steps — the mode behind the BrowseComp jump from 83.2% (single-agent) to 86.3% (swarm). K2.6 is the first open-weight model family to productize a sub-agent swarm tier as a named mode rather than an API pattern.
- **Context handling**: automatic context compression on long sessions; the 98,304-token output ceiling applies to reasoning tasks, well above typical chat-output limits.

#### Kimi K3 — architecture detail only (chronology canonical in wave3/02 Claim 4)

- **Scale**: 2.8T total / 104B active (canonical per wave1/wave3 recording); **896 routed experts + 2 shared experts, 16 routed active per token** — the most extreme fine granularity in a shipped open model (sparsity 56).
- **Stable LatentMoE**: the full hidden state is projected into a narrower latent space before expert dispatch — reported latent width **ℓ = 3,584 vs full model width 7,168** — approximately **halving routed-expert activation traffic and expert compute** relative to operating at full width. The latent space is normalized (Normalized LatentMoE).
- **Per-expert geometry** (from architecture inspections of the technical report): w1/w3 shapes [3072, 3584], w2 [3584, 3072], **MXFP4 with E8M0 scale per 32 weights** — see Consolidation notes for the flagged internal inconsistency in the 33.0M-per-expert figure.
- **SiTU-GLU activation** (β₁=4, β₂=25, ‖f(x)‖∞ ≤ 100) bounds activations to stabilize low-precision training; **RMSNorm** applied before the up-projection.
- **Quantile Balancing** replaces fixed-step expert-bias adjustment: each expert's bias is set to the router-score quantile matching its target load, and a **single global histogram / all-reduce** establishes the quantile thresholds.
- Reported outcome: **~2.5× scaling efficiency over Kimi K2** per the technical report [VENDOR].
- Design thesis: K3 does not refute the bandwidth cost of fine granularity — it *prices and pays* it with compression (latent routing, MXFP4) and single-all-reduce balancing, making top-16-of-896 routing trainable. This is the concrete 2026 counter-model to any claim that fine-grained MoE is bandwidth-free.
- "Beats Opus 4.8" is **benchmark-dependent**: GDPval-AA v2 1687 vs 1600; AA Intelligence ~57 (full evidence in wave3/02 Claim 4).

#### MiniMax H3 — spec summary (detail canonical in wave2.1/07 §1.1)

- **33B dense H3-Omni-Transformer**; text/image/video/audio in → video + **native 32 kHz stereo audio** out (4–15s clips, 768p local / 2K via API).
- **First open model to top an AI video ranking** (#1 Video Editing with Audio on Artificial Analysis).
- **"MiniMax H3 Community License"** with **US / EU / UK / South Korea territorial exclusion** from local deployment — the primary-license restriction for the KB's licensing index.
- 2K upscaling module and H3-Context-IR remain closed.
- Naming convention confirmed by secondary sources: **"H" = video/Hailuo line, "M" = text LLM line** (M2/M3 are the text models; H3 is the video model).

#### Seed 2.1 Turbo — specs

- **262K-token context** (262,144); input text/image/video → text output; reasoning supported (full effort ladder; `none` disables thinking); tool calling and structured output; streaming.
- **Closed/proprietary**: no license published, **parameter count undisclosed**, no training cutoff disclosed [UNVERIFIED on all three].
- Positioned as the faster, lower-cost variant for high-throughput / latency-sensitive agent workloads; the Pro is the flagship-class tier.
- ByteDance's launch framing used chart-only figures with superlatives and no absolute numbers in prose ("leading scores," "currently leads") — one journal transcribed Seed 2.1 Pro preview at **8th on Code Arena, 1539**, framed as "level with Claude Opus 4.6" (chart-sourced secondary — treat with caution). **No independent benchmark scores exist for Turbo specifically** [UNVERIFIED].

### Benchmarks (all figures [VENDOR] unless otherwise labeled)

#### GLM-5.2

