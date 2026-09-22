---
id: ai-industry-kb-2026/14-cloud-gpu-neocloud-economics/overview
title: "14. Cloud GPU & Neocloud Economics"
domain: cloud-gpu-neocloud-economics
role: deep-dive
task: finance
actors: ["AMD", "Anthropic", "CoreWeave", "Crusoe", "DeepSeek", "Lambda", "Meta", "Microsoft", "Nebius", "Nscale", "Nvidia"]
dates: ["2025-03", "2025-09", "2025-11-11", "2025-11-18", "2026-01-09", "2026-02", "2026-03", "2026-03-16", "2026-04-09", "2026-05-23", "2026-08", "2026-08-16", "2026-08-25", "2026-08-28", "2026-09-20", "2026-09-22", "2032-12"]
keywords: ["gpu", "neocloud", "amd", "benchmark", "blackwell", "compute", "deepseek", "deflation", "ethernet", "gpus", "hbm", "hyperscaler"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6932, 7010]
section: "14. Cloud GPU & Neocloud Economics"
sha256: 2041e680cfef338346261410de623dfc7c8e114a30d8a709065507e0ed091d07
---

# 14. Cloud GPU & Neocloud Economics
Keywords: cloud GPU rental, neocloud, GPU price index, on-demand GPU pricing, spot GPU, serverless inference, per-token pricing, unit economics, inference arbitrage, CoreWeave, Nebius, Lambda Labs, Vast.ai, RunPod, DeepInfra, B200, H100, H200, GB200 NVL72, Vera Rubin, DeepSeek V4 pricing, Nscale IPO, inference spend share, HBM shortage

## Summary

As of September 22, 2026, the cloud GPU market is a **split market**: spot and marketplace capacity is cheap and abundant, while committed/contracted capacity is tight and rising. The first half of 2026 also brought hyperscaler-scale neocloud contracts that dwarfed everything before them — CoreWeave–Meta ~$21B (announced April 9, 2026, through December 2032) and Nebius–Meta up-to-$27B (March 16, 2026) — while the neocloud IPO pipeline opened its second leg with Nscale's $30B NYSE filing (September 20, 2026).

- **On-demand rental (Sept 2026):** H100 SXM settled index (OCPI) **$2.79/hr** (Sep 21); H200 **$5.07/hr**; B200 **$7.71/hr** (up 14.3% in 30 days); B300 median ~$7.92/hr; MI300X ~$2.00–$3.59/hr; A100 settled **$0.99/hr**.
- **Spot/preemptible:** 25–45% below on-demand where offered (H100 spot from ~$1.47/hr on Vast.ai); reserved 1-yr/3-yr cuts 30–50% below on-demand; hyperscalers charge 3–6× neocloud rates for identical silicon.
- **Serverless per-token inference for open weights is "a few cents per million tokens" — verified** (DeepInfra gpt-oss-20b $0.03/$0.14, gpt-oss-120b $0.04/$0.17, Llama-3.1-8B $0.03/$0.05, all per 1M input/output, Sep 2026).
- **DeepSeek ended flat pricing on August 16, 2026**, moving V4 to peak/off-peak (V4-Pro output peaks at $3.96/1M, 355% above the old $0.87 flat) — repricing the direct-vs-gateway math; DeepInfra's flat rates are a genuine hedge against peak hours.
- **Unit economics rule of thumb:** per-token serverless wins for spiky/low volume; a dedicated H100 wins above ~$70–144/day of sustained token spend; rent-vs-buy breakeven sits at 40–77% sustained utilization.
- **The shortage did not ease:** HBM memory — not GPU dies — is the binding constraint; B200 resale hit 158% of launch price (Sep 2026); the two-year H100 deflation (~58% since early 2024) has stalled; 1-year contract pricing rose ~40% (Oct 2025–Mar 2026).
- **Inference overtook training in 2026 cloud spend — but the brief's "~70%" is CONTRADICTED at global level:** Gartner (Aug 2026, AI-optimized IaaS forecast) puts 2026 at **$42B total, inference $23.3B vs training $19B = ~55% inference**, rising to 59% in 2027. The 70–80% figure only holds for the narrower scope of GPU cloud spend among production AI teams (secondary compilations) — compute share (~2/3 per Deloitte) is not spend share.
- **Neocloud market projection:** Mordor Intelligence estimates **$35.22B (2026) → $236.53B (2031) at 46.37% CAGR** — a [DIRECTIONAL] commercial forecast from one firm, not an audited market result.
- New supply-side facts: **MI355X rental moved from "contact sales" to bookable** (Vultr launched ~Aug 8, 2026; TensorWave $2.95/GPU-hr on-demand; DigitalOcean spot $4.50); **Lium (Bittensor Subnet 51)** is verified as a live decentralized GPU-rental entrant; **Vera Rubin NVL72 still has no public rental price** as of Sep 2026.
- Lambda Labs' $1.5B raise was **November 18, 2025** (Series E, ~$5.9B post-money), not 2026; its genuinely-2026 news is **~$1B in private debt (Aug 28, 2026)** including a $926M GB300 loan, plus reported $3B pre-IPO talks at $12B+ valuation.

## Key dated facts

### Contract megadeals (first half of 2026)

- **2026-04-09:** CoreWeave and Meta announce an expanded AI infrastructure agreement worth **~$21B**, capacity contracted through **December 2032**, including the first NVIDIA Vera Rubin deployments. It builds on the September 2025 ~$14.2B agreement; Morningstar describes it as CoreWeave's largest-ever customer contract, bringing Meta's cumulative CoreWeave commitment past **$35B**. The superlative "largest AI cloud contract in history" is **NOT demonstrated industry-wide** — the parallel Nebius–Meta deal (up-to-$27B, March 2026) is comparable in scale, so no categorical industry record can be sustained.
- **2026-03-16:** Nebius and Meta sign a new five-year AI capacity deal — **$12B of dedicated capacity plus up to $15B of additional future capacity**, i.e. up-to-**$27B total**, positioned around Vera Rubin deployments. This **supersedes** Nebius's original $3B Meta deal, which was announced **November 11, 2025** (not February 2026 as a source brief misdated).
- **2025-09 (context):** Nebius–Microsoft **$17.4B** agreement (some reports cite expandability to ~$19.4B by 2031) — Microsoft was the first mover on large neocloud contracts; Meta followed and then overtook it.
- **2026-02:** NVIDIA invests **$2B in Nebius** (~8.3% stake), widely read as a vendor backstop for the neocloud buildout.

### Lambda Labs financing chronology (corrected dates)

- **2025-11-18:** Lambda raises **over $1.5B Series E**, led by TWG Global with USIT and existing investors, at **~$5.9B post-money** (Bloomberg had whispered $4B–$5B before the round).
- **2025-11:** Lambda's **multibillion-dollar Microsoft agreement** to deploy AI infrastructure powered by tens of thousands of NVIDIA GPUs (GB300 NVL72-class systems included).
- **2026-01-09:** SiliconAngle reports Lambda raising a **$350M round** in early 2026.
- **2026-08-25:** Bloomberg reports talks to raise **up to $3B in pre-IPO financing at a $12B+ valuation**; multiple term sheets received, terms not finalized as of Sep 2026.
- **2026-08-28:** Lambda secures **~$1B in private debt** to buy more NVIDIA chips (TechCrunch) — including a **$926M loan to finance GB300 GPUs** for an NVIDIA deployment; proceeds also support Microsoft-related capacity. Peer comp: **Crusoe raised $1.375B at a $10B valuation** in the same window.

### Neocloud leverage and IPO pipeline (Sep 2026)

- **2026-09-20:** Nscale files for a **$30B NYSE IPO** — disclosures include a **$1B NVIDIA note subscription** plus $10B authorization for future notes/shares, a **$45B Anthropic deal** [SECONDARY — from S-1 reporting via tech press, not independently audited], H1 2026 revenue **$140.6M** against **$103.4B total contract value**, $12.4B total liabilities, 52% of H1 revenue from one customer, and material weaknesses in internal financial controls. Context: CoreWeave IPO'd March 2025 and was at $2.575B Q2 2026 revenue (+112% YoY) — Nscale is the second-wave neocloud IPO.
- **Microsoft has committed $33B+ in neocloud capacity agreements**, the largest being **$19.4B with Nebius for 100K+ GB300 chips** (Sep 2026). Analyst framing: "everyone wants Nvidia GPUs, and neoclouds have them" — neoclouds undercut hyperscaler GPU pricing by **60–70%**.
- Nebius's disclosed unit economics (Q2 2026): deals averaged **>$20M/MW annual contract value** (2026 base near $12M/MW); Q3 short-term capacity clearing **>$40M/MW**; the company ran its **first capacity auction** — the committed tier is being price-discovered upward in public.
- Counter-read: CoreWeave's financial risk is quantified at **~$51.6B total debt vs $5.5B cash** (Sep 2026); short-seller Jim Chanos warns GPU-backed debt structures (CoreWeave, Lambda's $500M GPU-backed loan + $1.5B NVIDIA leasing arrangement, $20B+ sector-wide) are vulnerable if residual values fall. NVIDIA's **$2B CoreWeave injection at $87.20/share (Jan 2026)** is widely read as a vendor backstop.

### DeepSeek V4 peak/off-peak repricing (2026-08-16)

- DeepSeek ended flat per-token pricing on **August 16, 2026**, replacing the permanent-discount regime in effect since May 23, 2026. Peak windows: **01:00–04:00 and 06:00–10:00 UTC** (7 of 24 hours); off-peak bills at half price. Ten days' notice was given Aug 6; final figures published Aug 13 alongside the V4-Pro-0813 GA release.

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

