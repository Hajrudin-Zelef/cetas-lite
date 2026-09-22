---
id: ai-industry-kb-2026/15-hardware-chips/main-actors
title: "Main actors"
domain: hardware-chips
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Anthropic", "Broadcom", "Cerebras", "China", "CoreWeave", "Crusoe", "DeepSeek", "EU", "Google", "Groq", "Lambda", "Meta", "Microsoft", "Nebius", "Nscale", "Nvidia", "OpenAI", "Samsung", "TSMC", "United States"]
dates: ["2025-03", "2025-04", "2025-06", "2025-10", "2025-12-24", "2026-01", "2026-01-05", "2026-02", "2026-02-25", "2026-03", "2026-03-16", "2026-03-19", "2026-04-17", "2026-04-27", "2026-05", "2026-05-14", "2026-05-20", "2026-05-22", "2026-05-31", "2026-06-01", "2026-06-30", "2026-07", "2026-08-18", "2026-08-21", "2026-08-24", "2026-09", "2026-09-16", "2026-09-17", "2026-09-20", "2026-09-21", "2026-09-22"]
keywords: ["accelerator", "acquisition", "amd", "asic", "aws", "benchmarks", "blackwell", "capex", "compute", "cost", "cpo", "datacenter"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7633, 7745]
section: "15. Hardware & Chips"
sha256: 346bf002647780c976644866c13cd74f0f3582949420b81add247ffe119ff24c
---

# Main actors

## Main actors

### Platform owners

- **NVIDIA** — Vera Rubin NVL72 (72 Rubin GPUs + 36 Vera CPUs, NVLink 6, HBM4); six co-designed platform chips; "full production" declared CES 2026 (Jan 5) and re-confirmed Jul 21, 2026; ~$1T Blackwell+Rubin bookings visibility through 2027 claimed at GTC 2026; $20B non-exclusive Groq licensing (Dec 24, 2025); Rubin CPX cancelled at GTC 2026; roadmap defense vs SemiAnalysis Ultra/Kyber slip report. Jensen Huang (CEO), Ian Buck (hyperscale/HPC VP).
- **AMD** — Helios rack (72 MI455X, 31 TB HBM4, 1.4 PB/s, 2.9/1.4 EF [VENDOR]; EPYC Venice Zen-6, Pensando Vulcano; launches later 2026); MI440X (enterprise), MI430X (scientific/sovereign), MI500 (2027); MI300X thinly listed at neoclouds. Lisa Su (CEO; inaugural CES 2026 keynote; "world's best AI rack"; yotta-scale framing: 10 yottaflops by 2030).
- **Broadcom** — custom XPU business ($20B+ FY2025, Q1 FY2026 $8.4B +106% YoY) serving Google, Meta, OpenAI; 10-GW OpenAI accelerator + networking deal projected >$100B; "the quiet inference giant" and the clearest proof of the training/inference bifurcation.

### Inference-silicon disruptors

- **Groq** — LPU architecture originator; ~$20B non-exclusive NVIDIA license (not acquisition); founders Ross/Madra joined NVIDIA; company independent under CEO Simon Edwards; GroqCloud continues on prior-generation LPUs.
- **Cerebras** — public since May 14, 2026 (Nasdaq CBRS, $185, $5.55B); WSE-3 wafer-scale architecture (4T transistors, 44 GB SRAM); Artificial Analysis-verified 1,990.8 tok/s on Gemma 4 31B; AWS Marketplace GA; reported OpenAI $20B partnership [SECONDARY].
- **Etched** — Sohu transformer ASIC (TSMC N4P); stealth exit Jun 30, 2026; first rack to Jane Street (Jul 2026); $700M at $21B (Aug 18, 2026); >$1B claimed contracts; **no independent benchmarks as of Sep 22, 2026**. Kleiner Perkins' Mamoon Hamid (Aug 2026): "Inference is becoming one of the most important infrastructure markets in AI, and the winners will be measured by tokens per dollar and per watt."

### Hyperscaler custom-silicon programs

- **Google** — TPU v7 "Ironwood" ramp 2026; up-to-1M-TPU / 1-GW Anthropic commitment (announced Oct 2025); largest custom-silicon commitment by an AI lab to date.
- **Amazon** — Trainium 3 volume production 2026; validated by Anthropic and OpenAI (validation the earlier generations lacked); EC2 Trn2/Trn3.
- **Microsoft** — Maia 200 ("Braga", 216 GB HBM3e) deployed for Azure AI inference; testing first Vera Rubin racks with deployments in Georgia and Wisconsin and the Fairwater AI Superfactory; named Rubin cloud cohort.
- **Meta** — MTIA 500 (3 nm roadmap), internal-only.

### Neoclouds and the first Rubin cohort

- **CoreWeave** — first production Rubin NVL72 rack (Dell delivery May 31, 2026); industry-first rack-scale validation Jun 1, 2026 (delivery-to-production <6.5 h); measured DeepSeek R1 at 10× tokens/sec/MW vs Blackwell; ~$21B Meta contract (Apr 9, 2026, term through Dec 2032) including first Rubin deployments → §14.
- **Nebius** — first cloud adopter of Groq 3 LPX; Vera Rubin NVL72 US+Europe from H2 2026 via AI Cloud/Token Factory; up-to-$27B Meta contract (Mar 16, 2026) → §14; names CoreWeave, Crusoe, Lambda as key competitors in filings.
- **Dell** — delivered the first production Vera Rubin NVL72 rack (PowerEdge XE9812); Michael Dell announced on X with photos + diagnostic logs; under-6.5-hour bring-up.
- **Lambda, Nscale, Crusoe** — named in the CES 2026 early cloud cohort; Crusoe holds the cheapest enterprise-listed GB200 NVL72 rate ($16.50/GPU-hr, Mar 2026).
- **OpenAI** — anchor Rubin customer: deploying Vera Rubin at scale in Q3 2026 (Bloomberg-sourced); 6 GW AMD commitment; 10-GW Broadcom networking deal; $20B Broadcom-linked projection.

### Supply chain and manufacturing

- **TSMC** — N3P (Rubin), N4P (Sohu first-pass success); Samsung manufactures Groq 3 LPX; **Samsung and SK Hynix** supply HBM4 for Rubin; >350 factories across 30 countries (150 in Taiwan) in the Rubin supply chain.
- **HBM as the binding constraint:** DRAM +172% through 2025; GDDR7 diverted to datacenter memory; NVIDIA raised board-partner kit costs twice in 2026.

### Analysts, benchmarkers, brokers

- **SemiAnalysis** — InferenceX benchmarks (DeepSeek R1 FP4 on GB200 NVL72, May 22, 2026); Rubin Ultra/Kyber slip-to-2028 report (Jul 2026) vs NVIDIA's "roadmap intact."
- **Artificial Analysis** — third-party measurements: Gemma 4 31B at 3,431 tok/s on Groq 3 LPX; 1,990.8 tok/s on Cerebras WSE-3.
- **Jon Peddie Research** — Q2 2026 AI Processor report (151 companies, 290+ products).
- **Deloitte** — inference share half (2025) → two-thirds (2026) (Nov 2025 report).
- **Brokers/entrants:** GPUaaS.com, Bitdeer, WhiteFiber, Thunder Compute, Hyperbolic, Hyperstack.
- **Lenovo** — CEO Yuanqing Yang at CES 2026: the ~80/20 training/inference spend split will invert to ~20/80; three new inference servers launched; EVP Ashley Gorakhpurwalla on inference capex following adoption curves.

### Brokers, low-cost specialists, and benchmarkers

- **GPUaaS.com** — wholesale broker model: GB200 NVL clusters from 20+ vetted providers at "~30% under hyperscale," quotes in under 24h, no buyer fees. The broker layer exists because NVL72-class capacity is quote-only everywhere.
- **Bitdeer** — Bitcoin-miner-turned-AI-cloud with power/datacenter depth; markets H100–GB200 access with elastic scaling and pay-as-you-go.
- **WhiteFiber** — new-entrant HPC cloud (NYC) marketing B200-vs-GB200 NVL72 guidance and next-gen reservations.
- **Thunder Compute** — H100 $1.38/hr on-demand, among the cheapest vetted listings (May 2026).
- **Hyperbolic / Hyperstack** — low-cost specialists: Hyperbolic holds the B200 floor ($5.99/hr, Sep 2026 11-provider comparison); Hyperstack H100 $1.90/hr on-demand.
- Pattern: the long tail competes on price for fungible GPU-hours (inference track); the enterprise tier competes on allocation, SLAs, and early access to new silicon (training track).
- **SemiAnalysis** — InferenceX DeepSeek R1 benchmarks; Rubin Ultra/Kyber slip report; measured ~$0.20/1M for B200-served DeepSeek R1.
- **Artificial Analysis** — third-party speed measurements (Groq 3 LPX, Cerebras WSE-3).
- **Jon Peddie Research / Deloitte / Silicon Data** — structure, workload-share, and residual-value data.
- **AMD sovereign/scientific:** MI430X referenced for Oak Ridge Discovery and France's Alice Recoque exascale.

## Timeline and context

### Roots (before February 2026)

- **March 2025:** NVIDIA announces the Vera Rubin platform at GTC 2025 (rack named NVL144). The 3.6 EF FP4 / 1.2 EF FP8 / ~3.3× GB300 figures date to this roadmap reveal.
- **April 2025:** Google announces TPU v7 "Ironwood" at Cloud Next 2025 — 2026 becomes the ramp year, not the unveiling.
- **June 2025:** AMD first teases Helios at Advancing AI; Meta's ~$2B Manus acquisition announced Dec 30, 2025 → §18.
- **October 2025:** Google–Anthropic up-to-1M-TPU / 1-GW commitment announced.
- **December 24, 2025:** NVIDIA–Groq ~$20B non-exclusive LPU licensing deal.
- **2024–2025:** GB200 NVL72 ships (2024), GB300 NVL72 ships (2025); Cerebras's 2024 IPO attempt withdrawn over CFIUS review of the G42 stake (since divested).

### January 2026 — CES production confirmations

- **January 5, 2026:** Jensen Huang (CES keynote): Rubin in "full production"; six chips back from manufacturing; H2 2026 customer deployments; NVIDIA claims up to 3.5× training / 5× inference vs Blackwell, up to 10× lower inference token cost, 4× fewer GPUs for MoE training [VENDOR].
- **January 5/6, 2026:** Lisa Su (inaugural CES keynote): detailed Helios reveal — 72 MI455X, 31 TB HBM4, 1.4 PB/s, 2.9/1.4 EF [VENDOR]; EPYC Venice; Pensando Vulcano; launches later 2026. Partners on stage: OpenAI, Luma AI, Liquid AI, World Labs, Blue Origin, AstraZeneca, Illumina.
- **January 2026:** Lenovo's CEO forecasts the 80/20 training/inference spend split inverting to 20/80; first Vera Rubin production units reported reaching US/EU datacenters.

### February–March 2026 — samples, GTC, pricing bands

- **February 25, 2026:** first VR200 samples shipped (FY26 Q4 earnings call).
- **March 16, 2026 (GTC):** "seven new chips in full production" (six Rubin + Groq 3 LPU); ~$1T bookings visibility through 2027; Rubin CPX NVL put on hold, replaced by Groq 3 LPX; BlueField-4 DPU, Spectrum-X 102.4T CPO networking in the full-stack push.
- **March 2026:** GB200 NVL72 rental band $10.50–$27/GPU-equiv/hr; NVIDIA closes FY2026 with record $215.9B Data Center revenue; Vera Rubin NVL72 rack press pricing $5–7M.
- **March 19, 2026:** DOJ indictment (Super Micro-linked individuals) unsealed in Manhattan → §18.

### April–May 2026 — IPOs, contracts, guidance

- **April 17, 2026:** Cerebras files new S-1 (CBRS).
- **April 27, 2026:** NDRC prohibits Meta–Manus (Index No. 000013039-2026-00026), orders unwinding → §18.
- **May 13–14, 2026:** Cerebras IPO — $185/share, $5.55B raised, largest IPO of 2026.
- **May 20, 2026:** NVIDIA guides production shipments from fiscal Q3 (Aug–Oct), ramp fiscal Q4 — the contractual backdrop for the CoreWeave–Meta and Nebius–Meta Rubin clauses → §14.
- **May 22, 2026:** SemiAnalysis InferenceX — DeepSeek R1 FP4 on GB200 NVL72 at $0.04–$0.08/1M tokens at extreme concurrency.
- **May 31, 2026:** Dell delivers the first production Vera Rubin NVL72 rack to CoreWeave (PowerEdge XE9812).

### June–July 2026 — first bring-ups, shipments

- **June 1, 2026:** CoreWeave completes the industry's first Vera Rubin NVL72 rack-scale validation (Livingston, NJ) — delivery-to-production in under 6.5 hours; 10× tokens/sec/MW (DeepSeek R1, CoreWeave-measured) vs Blackwell.
- **June 30, 2026:** Etched exits stealth — Sohu chip (TSMC N4P), rack system, >$1B contracts claimed.
- **July 2026:** first Etched Sohu rack to Jane Street (production deployment); Broadcom custom-XPU revenue reported at $20B+ FY2025; Ian Buck (Jul 21): "absolutely in full production"; Jul 22: Rubin NVL72 confirmed in full production, OpenAI to deploy at scale in Q3 2026 (Bloomberg-sourced); SemiAnalysis reports possible Rubin Ultra/Kyber slip to 2028 — NVIDIA says roadmap intact [SECONDARY].
- Rack pricing datapoint (one buyer, third-party): **$7–8M per Vera Rubin NVL72 rack** [UNVERIFIED single secondary].

### August–September 2026 — LPX production, valuations, status

- **August 18, 2026:** Etched raises $700M at $21B (Jane Street leads) — valuation doubled in under a month.
- **August 21, 2026:** neocloud snapshot — CoreWeave GB200 NVL72 repriced to $10.50/GPU-hr; GB300 NVL72 contact-only; B200 $7.15 (Nebius), $6.69 (Lambda); H200 $4.29, MI300X $3.45 (Crusoe).
- **August 24, 2026 (Hot Chips):** Groq 3 LPX in full production; Nebius first cloud adopter; third-party benchmarks (Artificial Analysis) published.
- **September 16, 2026:** B200 residual value 158% of launch (Silicon Data).
- **~September 17, 2026:** CoreWeave links hundreds of Rubin GPUs into a multi-rack Spectrum-X cluster.
- **September 20, 2026:** structural shortage confirmed (HBM/foundry/power); Huang defends capex optimism — GPUs bought for both training and inference.
- **September 21, 2026:** DeepInfra price floor — gpt-oss-120b $0.04/$0.17 per 1M.
- **September 22, 2026:** Rubin NVL72 rental rates still unpublished (contract-only); Etched benchmarks still absent; Helios still pre-launch ("later 2026"); Rubin on schedule within the original H2 2026 window — no acceleration vs plan.

### Demand backdrop

- **September 2026 (reported, → §14):** China approves H100 sales licenses — ~$25B annual revenue opportunity.
- **September 20, 2026 (investor meetings):** Huang — enterprises and cloud providers buy GPUs for both training and inference; moving beyond prototypes requires "sustained, high-volume compute capacity."
- Committed-capacity price rise +~40% (Oct 2025–Mar 2026): none of the contracted training demand is price-elastic in the short run — the demand wall behind the rise, while the inference track's elasticity enables $0.03/1M tokens.

