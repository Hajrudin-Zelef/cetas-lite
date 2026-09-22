---
id: ai-industry-kb-2026/14-cloud-gpu-neocloud-economics/main-actors
title: "Main actors"
domain: cloud-gpu-neocloud-economics
role: deep-dive
task: finance
actors: ["AWS", "Anthropic", "Baseten", "Cerebras", "CoreWeave", "Crusoe", "DeepSeek", "EU", "Fireworks AI", "Google", "Groq", "Lambda", "Meta", "Microsoft", "Moonshot", "Nebius", "Nscale", "Nvidia", "OpenAI", "OpenRouter", "Oracle", "Together AI"]
dates: ["2025-07", "2025-10", "2025-10-15", "2025-11-11", "2025-11-18", "2025-12", "2026-01", "2026-01-09", "2026-03", "2026-03-16", "2026-04-09", "2026-05-14", "2026-05-23", "2026-06-18", "2026-08-08", "2026-08-16", "2026-08-25", "2026-08-28", "2026-09", "2026-09-11", "2026-09-13", "2026-09-16", "2026-09-17", "2026-09-20", "2026-09-21", "2026-09-22"]
keywords: ["acquisition", "agentic", "aws", "benchmark", "blackwell", "compute", "cost", "deepseek", "deflation", "fine-tuning", "gpu", "gpus"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7171, 7279]
section: "14. Cloud GPU & Neocloud Economics"
sha256: 90bf15c7ae9ec046e9010bf801db2835abf329a1f51d18d23d766375d77ef29c
---

# Main actors

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

**DeepInfra / Novita / Hyperbolic** — Low-cost per-token specialists. gpt-oss-120b ~$0.03–$0.10/$0.14–$0.50. Best for: cheapest managed open-weight tokens (watch concurrency caps and 429s).

**Baseten / Replicate / Beam** — Managed deployment platforms at premium per-GPU-hour (Baseten H100 $6.50, Replicate H100 $5.49) — the premium buys capacity management, region-locking, enterprise support.

**Nscale** — Second-wave neocloud IPO: filed $30B NYSE (Sep 20, 2026); $1B NVIDIA note; $45B Anthropic deal [SECONDARY]; H1 revenue $140.6M vs $103.4B TCV.

**Lium (datura-ai)** — Decentralized GPU-rental network, Bittensor Subnet 51; CLI + Python SDK; TAO-funded accounts; top subnet rank ~Sep 11, 2026. New entrant, live.

**OVHcloud / DigitalOcean / Thunder Compute** — Mid-tier: OVHcloud H100 ~$3.00/hr (EU/CA); DigitalOcean H100 $3.39/hr, MI325X $3.80/hr, MI355X spot $4.50/hr; Thunder Compute H100 $1.38/hr (NVL variant).

**Hyperscalers** — AWS, Google Cloud, Oracle, Azure, Cloudflare Workers AI: 3–6× neocloud for identical silicon (GCP B200 $16.11/hr ceiling); Azure NCadsH100v5 $6.98/hr; Cloudflare Workers AI for edge-adjacent serverless inference (Kimi K2.5 $0.60/$3.00).

## Timeline and context

| Date | Event |
|---|---|
| Early 2024 | Peak scarcity — H100 on-demand above $7–$8/hr on neoclouds. |
| Late 2024–mid 2025 | Steady deflation as Hopper supply normalized; H100 cohort median fell toward ~$3/hr. |
| 2025-09 | CoreWeave–Meta ~$14.2B agreement; Nebius–Microsoft $17.4B deal. |
| 2025-10-15 | Oct 2025: local bottom for committed capacity — SemiAnalysis 1-year H100 index at $1.70/hr. |
| 2025-11-11 | Nebius–Meta $3B deal (NOT Feb 2026). |
| 2025-11-18 | Lambda raises over $1.5B Series E (~$5.9B post-money); Lambda–Microsoft multibillion GPU deal. |
| 2025-11 | Deloitte report: inference ~half of AI compute in 2025, projected two-thirds in 2026. |
| 2025-12 | SDB200RT B200 rental benchmark launched; Dec 2025 10% counter-seasonal spot spike. |
| 2026-01-09 | SiliconAngle reports Lambda $350M round in early 2026. |
| 2026-01 | CES 2026: Lenovo CEO forecasts training/inference spend split inverting to 20/80. |
| 2026-02 | NVIDIA invests $2B in Nebius (~8.3% stake). |
| 2026-03-16 | Nebius–Meta up-to-$27B five-year deal (supersedes $3B deal). |
| 2026-03 | SDB200RT surged +24% in a single month (HBM3e passthrough, GTC repricing, model-release wave). |
| 2026-04-09 | CoreWeave–Meta expanded ~$21B agreement (through Dec 2032); Meta cumulative past $35B. |
| 2026-05-23 | DeepSeek permanent-discount regime begins (ends Aug 16). |
| 2026-06-18 | B200 settled low of $4.22/hr (OCPI 118-day window) — cheapest Blackwell hour on record. |
| 2026-07 | Lambda B200 was $3.79/hr in July 2025 vs $6.69 in 2026 — ~77% repricing over one year. |
| 2026-08-08 | Vultr launches MI355X Cloud GPU availability (~Aug 8). |
| 2026-08-16 | DeepSeek V4 moves to peak/off-peak pricing (+355–371% peak output increases). |
| 2026-08-25 | Bloomberg reports Lambda $3B pre-IPO talks at $12B+ valuation. |
| 2026-08-28 | Lambda secures $1B private debt (incl. $926M GB300 loan). |
| 2026-09-11 | Lium flips Chutes as Bittensor's top subnet (~Sep 11). |
| 2026-09-13 | Mercatus snapshot: H100 $3.84 avg (27 providers), H200 $4.43 (20), B200 $6.39 (12). |
| 2026-09-16 | B200 residual value 158% of launch (Silicon Data). |
| 2026-09-17 | CoreWeave multi-rack Rubin cluster milestone (~Sep 17). |
| 2026-09-20 | Nscale files $30B NYSE IPO. |
| 2026-09-21 | OCPI settled: H100 $2.79 (+4.0% 30d), H200 $5.07, B200 $7.71 (+14.3% 30d, 118-day high); DeepInfra gpt-oss-120b at $0.04/$0.17. |
| 2026-09-22 | Research date: Vera Rubin NVL72 still has no public rental price; committed capacity tight, spot abundant. |

### Price history 2024 → September 2026 (narrative)

Early 2024 was peak scarcity (H100 above $7–$8/hr). Late 2024–mid 2025 brought steady deflation to ~$3/hr as Hopper supply normalized. October 2025 marked the local bottom for *committed* capacity ($1.70/hr 1-year index). December 2025–January 2026 saw a 10% counter-seasonal spot spike. March 2026 brought the SDB200RT +24% single-month surge. Mid-2026: H100 cohort median ~$2.99/hr (down ~57–58% from the early-2024 peak), B200 neocloud band $5.29–$7.05/hr. September 2026: OCPI settled H100 $2.79, H200 $5.07, B200 $7.71 (118-day high). **Net 2026 direction:** spot/marketplace flat-to-soft; committed/contract pricing up ~40% (Oct 2025–Mar 2026); Blackwell spot premiums rising into September. The deflation era is over; the market is now bifurcated.

### Reading GPU price data: why published ranges are so wide

Any single "$/hr" figure is meaningless without four qualifiers — this explains the 5–10× spreads in published comparisons:

1. **Listings vs transactions:** provider rate cards vs transaction-settled indices (OCPI). A listed $2.89/hr you cannot book (no capacity) is not a competing rate — Spheron's table explicitly annotates "listed but no capacity" cells.
2. **Contract commitment:** on-demand vs 1-yr vs 3-yr reserved vs 36-month prepaid (Vultr's headline rates are 36-month prepaid back-calculated to hourly).
3. **Configuration:** single GPU vs 8-GPU node pricing; PCIe vs SXM variants; NVL variants.
4. **Billing granularity and add-ons:** per-second (Modal, RunPod) vs per-minute (Jarvislabs, Spheron, 20-min minimum) vs hourly; storage, egress (Lambda: none), CPU/RAM metered separately (Modal).
5. **Multipliers:** Modal's region (1.5–1.75×) and non-preemptible (3×) multipliers; combined up to ~3.75× headline.
6. **Capacity reality:** "low capacity" / "max 1" / "no capacity" annotations mean the effective market price is the *bookable* rate, often one tier up.

RAG rule: store GPU prices as **(value, date, provider, commitment tier, source)** tuples; prefer dated index settlements (OCPI) over static figures; refresh at least monthly (±5% daily moves observed in Sep 2026).

### Pricing-terms glossary

- **On-demand:** pay per hour/second, no commitment; the headline rate most comparisons quote.
- **Spot / preemptible:** spare capacity at 25–45% discount; reclaimable anytime (Vast.ai, Spheron) or the default base tier (Modal).
- **Reserved (1-yr / 3-yr):** committed capacity at 30–50% below on-demand; 36-month prepaid (Vultr) back-calculates to the lowest hourly figures published.
- **Per-token (serverless inference):** billed per 1M input/output tokens; input/output priced separately; output typically 3–7× input.
- **Batch API:** async (~24h) inference at ~50% of real-time serverless (Fireworks, Together).
- **Prompt caching:** discounted input tokens on cache hits (50–90%+ off; Together as low as $0.04/1M).
- **OCPI:** Ornn Compute Price Index — daily-settled, transaction-based $/GPU-hr benchmark, Bloomberg-listed.
- **Crossover point:** daily token spend at which a dedicated GPU becomes cheaper than per-token billing (~$70/day raw rental, ~$144/day managed, for H100-class).
- **Utilization breakeven:** sustained GPU utilization above which buying beats renting (40–77% per 2026 studies).
- **Inference arbitrage:** routing identical open-weight models across providers to capture per-token price spreads.

