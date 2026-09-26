---
id: ai-industry-kb-2026/03-chinese-ai-labs-deep-file-glm-kimi-minimax/individuals
title: "Individuals"
domain: chinese-ai-labs-deep-file-glm-kimi-minimax
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "ByteDance", "China", "DeepSeek", "MiniMax", "Moonshot", "OpenAI", "United States", "Z.ai"]
dates: ["2026-01", "2026-06-23", "2026-07", "2026-08-14", "2026-09", "2026-09-22"]
keywords: ["agent", "agentic", "attention", "benchmark", "claude", "cost", "cyber", "deepseek", "fable 5", "fp8", "glm", "ipo"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1237, 1286]
section: "3. Chinese AI Labs — Deep File: GLM, Kimi, MiniMax"
sha256: 0fcd043e9e5dd3b7f4547f40fc6e14a63db5d2913d267d4601104ae2f13f205a
---

# Individuals

- **Z.ai / Zhipu AI** (Beijing): the GLM line (GLM-5.2, GLM-5.3, GLM-5.3-Flash); IndexShare; the GLM Coding Plan subscription and ZCode agent; OpenVuln/VulnHunter. Listed publicly in Hong Kong before MiniMax (Reuters notes MiniMax was second of the "AI tigers" to list, following Z.AI).
- **Moonshot AI** (Beijing): the Kimi line (K2.5 → K2.6 → K3); Modified MIT open-weight licensing with the 100M-MAU / $20M-revenue display threshold; MoonViT vision encoder; Stable LatentMoE and Quantile Balancing as 2026 MoE design contributions.
- **MiniMax** (Shanghai, founded 2022): one of China's "AI tigers" — the well-funded startup group competing with domestic tech giants and US labs. Publicly listed in Hong Kong in **January 2026**. The "H" (video/Hailuo) vs "M" (text LLM) naming convention disambiguates H3 (video) from the M2/M3 text models. [UNVERIFIED] 2.7T-parameter LLM reportedly in development (Reuters, early July 2026).
- **ByteDance Seed** (Volcano Engine): Seed 2.1 family (base, Pro, Turbo) launched at the FORCE conference 2026-06-23/24; Seedance video models (Seedance 2.5 canonical in wave3/06 §1.3). Positioned by secondary review around "productivity delivery" — end-to-end software delivery from planning to verification — rather than raw benchmark leadership.

### Individuals

- **Yang Zhilin** — Moonshot AI founder (one secondary source notes he was a student of Tang Jie; biographical, low relevance).
- **Tang Jie** — the professor who later co-founded Z.AI — the biographical link between the two labs.

### Benchmark counterparties (for comparative framing, not lab actors)

- Claude Opus 4.6 (53.4% SWE-Bench Pro — the number K2.6 beat); Claude Opus 4.8 (the model K3 beats on GDPval-AA v2, benchmark-dependently); GPT-5.4 (57.7% SWE-Bench Pro); Anthropic's Fable 5 / Mythos 5 (the export-control suspension that GLM-5.2's launch followed); DeepSeek and Qwen3.8-Max (the open-weight frontier context for the 2.7T rumor tier: Kimi K3 2.8T, Qwen3.8-Max 2.4T).

### Business context (Reuters launch-day coverage, MiniMax H3)

- MiniMax framed H3 as stepping up competition in a market **"led by rivals ByteDance and Kuaishou"** — consistent with wave3/06's "China-led, closed-model race" thesis for text-to-video.
- The H3 launch commercial claim (2K video at <1/3 rival cost) [VENDOR] and the January 2026 HK listing together mark MiniMax as the capitalized challenger in the video-model race.

### Positioning notes

- **ByteDance's thesis** (secondary review): "productivity delivery" — end-to-end software delivery from planning to verification — rather than raw benchmark leadership. This is how the RAG should frame ByteDance's 2026 lab strategy versus the frontier-chasing posture of Moonshot and Z.ai.
- **Kuaishou** appears alongside ByteDance as the market co-leader in Reuters' framing of the video-model race — relevant context for H3's competitive positioning, not a model actor in this file.
- **Z.ai vs Moonshot founder link**: Moonshot's Yang Zhilin was reportedly a student of Tang Jie, who later co-founded Z.AI — the two labs' leadership shares an academic lineage (low relevance, recorded for completeness).
- **DeepSeek and Qwen** are the open-weight frontier reference points around these releases: DeepSeek Sparse Attention is GLM-5.2's attention substrate, and Qwen3.8-Max (2.4T) defines the tier the rumored MiniMax 2.7T LLM would enter.

## Timeline and context

The April-to-September arc for these three labs runs as a single competitive sequence. **April**: Moonshot opens the Kimi K2.6 preview (04-13) and reaches GA (04-20/21) with 1T/32B open weights, 384 experts, a 300-sub-agent swarm mode, and SWE-Bench Pro 58.6% — ahead of Claude Opus 4.6's 53.4%, a data point the brief undersold as "near." **June**: Z.ai ships GLM-5.2 (weights 06-16) the day after the US export-control order suspending Anthropic's Fable 5 / Mythos 5 — the open-weight counter-move, with 1M context, IndexShare, and the "best open-weight text model" crown valid only at that moment [DIRECTIONAL]; ByteDance launches the Seed 2.1 family at FORCE (06-23/24), the Turbo variant resting on a single trade report. **July**: Moonshot launches Kimi K3 (07-16; weights ~07-27) at 2.8T with the bandwidth-aware latent-MoE design, and MiniMax launches H3 = Hailuo 3.0 (07-31) — the first open model to top an AI video ranking — while Reuters reports the [UNVERIFIED] 2.7T MiniMax LLM in development. **August**: Z.ai announces GLM-5.3 (08-14) on the same 743B base, all gains from scaled post-training — the clearest 2026 instance of the post-training-scaling thesis (Terminal-Bench 3.0: 4.6 → 28.3 without new pretraining); the cyber capability outgrew Z.ai's training expectations, triggering the GLM line's first weight hold (~2 weeks; weights 08-28/29 under a bespoke license, a step down from MIT). **September 22**: the MiniMax 2.7T rumor and the Moonshot HK IPO rumor remain unconfirmed watch items.

Exact dates live in §2 (Key dated facts); the rumor-phase rule (pre-2026-08-14 GLM-5.3 mentions stay [UNVERIFIED]) applies to anything cited from the August speculation window.

### Month-by-month spine (April → September 2026)

- **April**: Kimi K2.6 — the agent-swarm open-weight release; 1T/32B, Modified MIT, SWE-Bench Pro ahead of Opus 4.6. Sets the 2026 template: trillion-scale open weights with agentic modes.
- **June**: GLM-5.2 (06-16) — the post-Anthropic-export-control counter-move; 753B/40B, 1M context, IndexShare; the "best open-weight text model" moment [DIRECTIONAL]. Seed 2.1 family at FORCE (06-23/24).
- **July**: Kimi K3 (07-16 launch; ~07-27 weights) — the bandwidth-aware 2.8T push (latent routing, Quantile Balancing, MXFP4). MiniMax H3 (07-31) — the open-video milestone; Reuters' [UNVERIFIED] 2.7T LLM report.
- **August**: GLM-5.3 — the post-training-scaling demonstration and the cyber-driven weight hold; rumor phase → 08-14 announcement → 08-18 pricing → 08-28/29 weights.
- **September**: verification cutoff 2026-09-22 — both watch items (MiniMax 2.7T, Moonshot HK IPO filing) still unconfirmed.

## Implications

### For the knowledge base design

- **Alias discipline is load-bearing.** Three names for one MiniMax model (H3 / Hailuo 3.0 / "Hailuo 03"), two numbers for one GLM-5.2 total ("744B" shorthand vs 753.3B full), three phases for one K2.6 release (preview / GA-day-1 / GA-day-2), and two dates for one Seed 2.1 Turbo launch (FORCE 06-24 vs tracker first-seen 08-10/12) — the KB must merge these into single entries with phase footnotes, or retrieval will triple-count. The consolidation notes (§8) carry the merge rules.
- **Counting conventions must be footnoted, not picked.** The GLM-5.2 case (backbone vs +MTP vs active vs FP8-build shorthand) is the template: pin the canonical total (753B/40B) and record the convention behind each competing figure.
- **Temporal bounds on superlatives.** "Best open-weight text model," "beats Opus 4.8," "first open model to top an AI video ranking" — each is [DIRECTIONAL] and valid only at its moment; the KB should date every leadership claim.

### Competitive implications

