---
id: ai-industry-kb-2026/14-cloud-gpu-neocloud-economics/mi355x-rental-becomes-bookable-augsep-2026
title: "MI355X rental becomes bookable (Aug–Sep 2026)"
domain: cloud-gpu-neocloud-economics
role: deep-dive
task: finance
actors: ["AMD", "AWS", "CoreWeave", "Crusoe", "DeepSeek", "Google", "Lambda", "Meta", "Microsoft", "Nebius", "Nvidia", "Oracle"]
dates: ["2026-08", "2026-09"]
keywords: ["accelerator", "amd", "aws", "benchmark", "blackwell", "compute", "cost", "deepseek", "ethernet", "gpu", "gpus", "hbm3"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6978, 7035]
section: "14. Cloud GPU & Neocloud Economics"
sha256: 1bb44acf9d0e89ba00516b2e35db0097e46967d5a3e08eb771d07af7d80fea4a
---

# MI355X rental becomes bookable (Aug–Sep 2026)

| Model | Peak output /1M | Off-peak output /1M | Peak input (cache-miss) /1M | Off-peak input /1M |
|---|---|---|---|---|
| V4-Pro | **$3.96** (was $0.87; +355%) | **$1.98** (was $0.87; +128%) | $1.32 (was $0.435) | $0.66 |
| V4-Flash | **$1.32** (was $0.28; +371%) | **$0.66** (was $0.28; +136%) | $0.44 (was $0.14) | $0.22 |

- Cache-hit input increases were steeper still — up to **1,100%** on some tiers (InfoWorld via techtimes.com). DeepSeek framed the change as "allocating resources more reasonably."
- RAG note: any undated "DeepSeek V4 $0.07/$0.28" figure is the **pre-Aug-16 promotional flat rate** (off-peak Flash class) and is now stale. `deepseek-chat` and `deepseek-reasoner` are deprecated aliases billing at Flash rates.

### MI355X rental becomes bookable (Aug–Sep 2026)

- **~Aug 8, 2026:** Vultr launches AMD Instinct MI355X Cloud GPU availability — 8-GPU Cloud GPU plans and bare-metal servers, ROCm-enabled images, deployable via control panel or API.
- Live price points (Sep 2026): TensorWave **$2.95/GPU-hr on-demand**; DigitalOcean **$4.50/GPU-hr spot**. Crusoe's page remains "contact sales" (provider-dependent, same pattern MI300X followed in 2025).
- Market context: dstack's Sep 2026 state-of-cloud-GPU note observes **AMD MI300X pricing softening as MI350X/MI355X roll out**, with some neoclouds undercutting H100/H200 on $/GPU-hr while offering more memory per GPU — the AMD memory-per-dollar story moving to 288 GB HBM3e cards.

### Vera Rubin NVL72: still no public rental price (Sep 2026)

- Confirmed unchanged: **no cloud provider has published an hourly rate for Vera Rubin NVL72** as of Sep 22, 2026. CoreWeave has not said whether its multi-rack cluster is generally available, which customers have access, or how capacity is priced. Treat any Rubin $/hr figure as speculation.
- **~Sep 17, 2026:** CoreWeave linked **hundreds of Rubin GPUs** into a multi-rack cluster with Spectrum-X Ethernet, non-blocking fabric rated for ~128,000 GPUs — the step from single-rack bring-up (Jun 1, 2026) to production-scale cluster.
- Rack-purchase reporting needs care: March-2026 press put Vera Rubin NVL72 at **$5–7M per rack**; a July-2026 $7–8M figure refers to the later **NVL144 Rubin Ultra** generation, not Vera Rubin NVL72.
- Nebius offers Vera Rubin NVL72 from **H2 2026** via AI Cloud and **Nebius Token Factory** (inference/post-training platform), complementing existing GB200/GB300 NVL72 capacity.

### Lium verified: Bittensor Subnet 51 GPU rental (new entrant)

- Verified as a live provider (github.com/datura-ai/lium, active repo, Sep 2026): **Lium** is a decentralized GPU-rental network — **Subnet 51 (SN51) on the Bittensor protocol**, developed by Datura-ai.
- Mechanics: CLI + Python SDK rent containerized GPU "pods" (docs example: 1×A100 at $1.20/hr, pod ready in ~48s); accounts funded with **TAO** from a Bittensor wallet; miners rewarded on performance/uptime validated by Bittensor validators.
- Position: benchmark repos ran **36 runs across B200, H200, H100, RTX PRO 6000 Blackwell, RTX 5090**; Lium **flipped Chutes as Bittensor's top subnet ~Sep 11, 2026**. Vs peers: Chutes (SN64) sells serverless inference, Targon (SN4) sells confidential compute — Lium sells the GPU-hours that feed both.
- Caveat [COMMUNITY]: subnet emissions rankings reflect network incentives, not audited external revenue; independent pricing comparisons vs centralized neoclouds are not yet available.

### Nebius buildout and competitive framing (Aug–Sep 2026)

- August 2026: Reuters reports Nebius's planned GPU buildout including **300,000 NVIDIA chips** and orders spanning **B300/GB300 and Vera Rubin** systems — consistent with the Meta and Microsoft commitments being real and active.
- In filings, Nebius names **CoreWeave, Crusoe, and Lambda** as key competitors alongside general-purpose clouds — the neocloud tier now defines its own competitive set independent of hyperscalers.

## Figures and metrics

### On-demand GPU rental prices (September 2026, USD per GPU-hour, single-GPU on-demand unless noted)

All figures need four qualifiers — listing vs transaction, contract commitment, configuration (PCIe vs SXM, single GPU vs 8-GPU node), billing granularity — see §Timeline and context / "Reading GPU price data." Anchors: transaction-settled indices (Ornn OCPI, Bloomberg-listed), which reflect executed trades rather than listings.

**H100 SXM (80 GB HBM3) — the reference price:**
- OCPI-H100 settled index: **$2.79/hr** (Sep 21, 2026), settled daily 4:00 PM ET; +5.2% (1 day), +1.4% (7 days), +4.0% (30 days). Three-month public window: $2.31 (low) to $3.17 (high).
- 27-provider volume-weighted average: **$3.84/hr** (Mercatus, Sep 13, 2026). AIMultiple cohort median: **~$2.99/hr** (mid-2026), down from above $7 in early 2024 — a ~57–58% two-year decline.
- Provider listings (Sep 2026): Vast.ai $1.40–$2.10 (spot ~$1.47); Hyperstack **$1.90**; Thunder Compute **$1.38**; Spheron $2.64–$2.75; Jarvislabs $2.69; RunPod $2.89 (per-second, 1-min minimums); OVHcloud ~$3.00 (PCIe); Lambda $3.29 (PCIe, no egress fees); DigitalOcean $3.39; Modal $3.95 (per-second, preemptible base); CoreWeave $6.16 (enterprise tier); AWS $3.22–$6.88; Google Cloud up to $11.68 (A3 Mega); Azure NCadsH100v5 $6.98.
- Monthly cost at index: ~$2,007.60/GPU-month continuous ($66.92/GPU-day). H100 listed at 36–46 providers — broadest availability of any current accelerator.
- An 8-GPU H100 node running 24/7: **$10,500/month at the cheapest provider vs $58,400/month at the most expensive** — identical hardware (Mercatus, Sep 13, 2026).

**H200 (141 GB HBM3e) — the memory play:**
- Average **$4.43/hr** across 20 providers (Mercatus, Sep 2026); OCPI-H200 settled **$5.07/hr** (Sep 21).
- Premium over H100: ~15–40% on neoclouds; up to 2×+ at hyperscalers. Provider listings: FluidStack from $2.30/hr (8-GPU nodes); RunPod $4.59 (low capacity); Spheron $4.80; CoreWeave $6.31; AWS $7.91; Google Cloud ~$10.60 (A3 Ultra).
- Why it matters: 141 GB HBM3e at the same 700W as H100 — one H200 serves a 70B model in FP16 where two H100s were needed; roughly triples effective KV-cache headroom. For memory-bound inference and long-context serving, the premium typically pays for itself.
- Contract signal: SemiAnalysis 1-year H200 rental index rose ~40% Oct 2025–Mar 2026; ~half of tracked specialists reported **no Hopper-class capacity coming off contract** in that window.

**B200 (180 GB HBM3e) — the premium tier:**
- Average **$6.39/hr** across 12 providers (Mercatus, Sep 2026); OCPI-B200 settled **$7.71/hr** (Sep 21): +7.2% (7 days), **+14.3% (30 days)**. 118-day window (May 27–Sep 21, 2026): low $4.22 (Jun 18), high $7.71 (Sep 20).
- 11-platform comparison (Sep 2026): Hyperbolic **$5.99** (floor), Hyperstack $6.00, Modal $6.25, Lambda $6.69–$6.99, RunPod $6.79 (low capacity, max 1), Nebius $7.15, Vultr $8.50, CoreWeave $8.60, Spheron $9.30 (spot **$5.37**), Oracle $14.00, AWS $14.24, Google Cloud **$16.11** (ceiling).
- Narrower specialist index: $5.29–$7.05/hr band; absolute lowest documented rate $5.29/hr (Lambda Labs, GridStackHub Sep 2026 index). Long-term reserved (36-month) contracts dropped as low as **~$2.25/hr**.
- Economics: one B200 hour costs **2.76× one H100 hour** (OCPI) while buying 2.28× the dense BF16 throughput and 2.25× the memory. Hardware context: B200 SXM volume list $30,000–$40,000/GPU; street/spot $45,000–$55,000; 8-GPU HGX server $400,000–$500,000; resale value **158% of launch** (Sep 2026). Liquid cooling mandatory (deployment constraint for small operators).

