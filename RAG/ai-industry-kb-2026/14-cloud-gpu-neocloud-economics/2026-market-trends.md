---
id: ai-industry-kb-2026/14-cloud-gpu-neocloud-economics/2026-market-trends
title: "2026 market trends"
domain: cloud-gpu-neocloud-economics
role: deep-dive
task: finance
actors: ["AMD", "Anthropic", "Broadcom", "Cerebras", "China", "CoreWeave", "Crusoe", "Fireworks AI", "Groq", "Lambda", "Meta", "Microsoft", "Nebius", "Nvidia", "OpenAI", "OpenRouter", "SpaceX", "Together AI", "xAI"]
dates: ["2025-11-11", "2025-11-18", "2026-03-16", "2026-05-14", "2026-08-28", "2026-09"]
keywords: ["acquisition", "agentic", "amd", "blackwell", "capex", "compute", "consumer", "cost", "deflation", "dram", "fine-tuning", "foundry"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7154, 7204]
section: "14. Cloud GPU & Neocloud Economics"
sha256: 14b0130bf16ea029cc4093937cd69d36915ccc908cd3ea1989818bb2c9115b90
---

# 2026 market trends

- **Structural shortage across HBM, foundry, and power** (sourcebyspec.com, Sep 20, 2026): demand concentrated in inference + training silicon.
- **Lead times (Mar 2026 sourcing data):** new H100 wait times shortened 6+ months → **2–3 months** through 2025; but **Blackwell-class lead times moved the opposite direction: 3–7 months by Q1 2026**; allocation "highly unstable across distributors."
- **HBM is the binding constraint:** DRAM prices rose ~172% through 2025 as HBM production displaced consumer DDR5; NVIDIA raised board-partner kit costs twice in 2026 (May, July). Silicon Analysts estimates B200 manufacturing COGS at **~$6,400**, with HBM3e now **45% of the bill of materials** (vs 41% on H100) — the memory-cost thesis quantified.
- **Purchase-price snapshot (Mar 2026):** L40S 48GB $8,610–$8,900; RTX 6000 Ada 48GB $7,400–$7,800; RTX PRO 6000 96GB $9,450–$9,800; new H100 80GB SXM $25,000–$35,000 (used $18,000–$22,000); A100 80GB discontinued new, $12,000–$18,000 used. Highest-pull SKUs: H200 NVL PCIe, H100 NVL PCIe, L40S, L4, RTX 4000 Ada.
- **B200 residual value: 158% of launch price** (Silicon Data, Sep 16, 2026) — installed Blackwell valued *above* new list ~1 year after mass availability; A100/H100 residuals also above straight-line depreciation.
- **Demand visibility:** NVIDIA record $215.9B FY2026 Data Center revenue (Mar 2026); four largest hyperscalers' combined ~$630–700B 2026 capex plans; OpenAI's ~26 GW secured pipeline (NVIDIA $100B/10 GW, AMD 6 GW, Broadcom); Stargate $500B/10 GW; xAI Colossus 2 (1M-GPU scale target); China's approval of H100 sales licenses (~$25B annual revenue opportunity, Sep 2026 reporting).

### 2026 market trends

- **The split market is the central fact of 2026** (GPUSmith/SemiAnalysis): cheap abundant spot (H100 $1.47–$2.20/hr) coexists with tight committed capacity; 1-year contract pricing rose ~40% Oct 2025–Mar 2026.
- **The two-year H100 deflation (~58% since early 2024) has stalled in 2026** — rates are flat-to-rising, not falling; buyers advised to lock pricing sooner rather than later.
- **HBM memory scarcity, not GPU dies, is the binding price driver** — the transmission from memory contract prices to GPU rental rates is "one of the most important cost signals in the AI economy" (Silicon Data Q1 2026).
- **Blackwell premium dynamics:** each new NVIDIA architecture launches at ~2× the prior generation's median, then compresses as neocloud supply catches up; B200 resale at 158% of launch shows the premium persisting through September 2026.
- **Compute as an asset class:** OCPI listed on Bloomberg; forward curves, utilization indices, and secondary markets (Ornn Compute) now exist; GPU price volatility (B200 +56% annualized) is traded.
- **Training-track demand overhang:** OpenAI 26 GW, Stargate $500B, xAI Colossus 2, Anthropic leasing SpaceX Colossus 1 (220K GPUs, secondary source) — multi-year contracted demand with no short-run price elasticity, behind the committed-capacity price rises; the inference track's elasticity is what enables $0.03/1M tokens.
- **Inference arbitrage is a market of its own:** "LLM providers are commodities, and the only winning strategy is Inference Arbitrage" (2026 practitioner analysis) — smart routers (e.g., Fireworks primary + DeepInfra burst-fallback ~60% cheaper, per Apr 2026 research) treat per-token prices as a tradeable spread.

## Main actors

**Vast.ai** — Peer-to-peer GPU marketplace. H100 $1.40–$2.10/hr, spot ~$1.47; B200 from $3.44/hr; B300 from $5.44/hr. Consistently cheapest in published comparisons; hosts are independent operators (reliability varies, checkpoint mandatory). Best for: fault-tolerant batch, experiments, budget-maximizing individuals.

**RunPod** — Vetted neocloud, community + secure tiers. H100 $2.89/hr pods (per-second, 1-min minimum); H200 $4.59/hr (low capacity); B200 $6.79/hr (low capacity, max 1); committed $4.78/hr (6-month) / $4.67/hr (1-year); serverless endpoints H100 $4.55/hr flex, B200 workers $8.64/hr. FlashBoot: 48% of cold starts under 200ms. Usually the cheapest *reliable* option.

**Lambda Labs** — Simplicity-first neocloud. H100 $3.29/hr (PCIe), B200 $6.69–$6.99/hr; no egress fees; ML-ready images. Financings: $1.5B Series E 2025-11-18, $350M early-2026 round, $1B private debt 2026-08-28 (incl. $926M GB300 loan), $3B pre-IPO talks at $12B+ (2026-08). Sacra pricing note: H100 PCIe ~$2.49/hr vs ~$4.25 at CoreWeave.

**Spheron** — Transparent catalog + marketplace; per-minute billing (20-min minimum), ~2-min deploys; publishes annotated cross-provider comparison. H100 $2.64–$2.75 (spot $2.10–$2.20); MI300X $3.59/hr. Best for: price discovery, fast deploys.

**CoreWeave** — Enterprise neocloud, Kubernetes-native. H100 $6.16/hr; H200 $6.31/hr; B200 $8.60/hr; GB200 NVL72 $10.50/hr on-demand (Aug 21, 2026). Reserve-only on some SKUs; bilateral lab contracts (OpenAI, Meta — $21B expanded deal Apr 2026, Meta cumulative past $35B). Risk: ~$51.6B debt vs $5.5B cash (Sep 2026); NVIDIA $2B injection at $87.20/share (Jan 2026).

**Nebius** — Enterprise neocloud (ex-Yandex infra). B200 $7.15/hr; HGX B300 $7.85/hr (preemptible $4.30). Deals: Microsoft $17.4B (Sep 2025); Meta $3B (2025-11-11) superseded by up-to-$27B five-year deal (2026-03-16); Microsoft $33B+ neocloud program includes $19.4B with Nebius for 100K+ GB300. Q2 deals >$20M/MW annual contract value; first capacity auction; 300K-chip buildout plan (Aug 2026). Vera Rubin NVL72 from H2 2026 via AI Cloud / Token Factory. Names CoreWeave, Crusoe, Lambda as key competitors.

**Crusoe** — Neocloud; named among popular 2026 providers; raised $1.375B at $10B valuation (2026-08). MI355X page still "contact sales." No current per-GPU list price captured.

**Hyperbolic** — Low-cost specialist. B200 floor $5.99/hr (Sep 2026 11-provider comparison); $3.50 on the gpusmith ladder. Best for: price-sensitive Blackwell inference.

**Hyperstack** — H100 **$1.90/hr** on-demand; B200 $6.00/hr. Among cheapest vetted neoclouds.

**Vultr** — Indie cloud gone aggressive on GPUs. GH200 (1 GPU, 96 GB) $2,913/mo ($4.335/hr) on 36-month prepaid reserved; 8× H100 node $13,432/mo reserved; MI300X/MI325X from $2.00/GPU/hr; fractional MI300X $2.29/hr; launched MI355X Cloud GPU (~Aug 8, 2026). Flagship SKUs quoted on 36-month 100%-prepaid reserved contracts — verify on deploy screen.

**Modal** — Serverless GPUs, per-second metering (CPU/RAM/GPU decoupled). H100 $3.95/hr, H200 $4.54/hr, B200 $6.25/hr base (preemptible); ×3 non-preemptible; ×1.25–2.5 non-default regions; $30/mo free credits. Best for: ephemeral batch, bursty inference (watch `min_containers` cost traps).

**Fireworks AI** — Per-token serverless (size-tiered $0.10–$0.90/1M + MoE tiers) + dedicated GPU ($2.90–$12/hr by SKU) + batch at 50% off. Ex-Meta PyTorch team; function-calling reliability and tail-latency differentiators. Best for: production agentic systems, strict JSON/tool-call workloads.

**Together AI** — 200+ open models, per-token serverless + fine-tuning + dedicated endpoints. Best for: model evaluation across many OSS models, fine-tune-then-deploy on one platform.

**Groq** — LPU inference. Prior-gen LPU pricing: Llama 3.1 8B $0.05/$0.08 (~840 tok/s); 70B $0.59/$0.79. GroqCloud pushing flagships to enterprise tiering while gpt-oss OSS weights remain the self-serve front door. NVIDIA licensed Groq's LPU tech for ~$20B (Dec 24, 2025, non-exclusive — not an acquisition); Groq 3 LPX in full production (Hot Chips, Aug 24, 2026); Nebius first cloud adopter.

**Cerebras** — Wafer-scale WSE-3. ~$0.10/1M (8B) to ~$0.60–$0.85/$1.20 (70B); 1,700–3,000 tok/s. IPO May 14, 2026 (CBRS, $185, $5.55B raised). Best for: maximum tokens/sec.

**OpenRouter** — Aggregator/router across 400+ models; deepest routing bench for gpt-oss-120b (20 providers). Best for: model fallback/routing, single API key.

