---
id: ai-industry-kb-2026/15-hardware-chips/january-2026-ces-production-confirmations
title: "January 2026 — CES production confirmations"
domain: hardware-chips
role: deep-dive
task: hardware
actors: ["AMD", "Anthropic", "Broadcom", "Cerebras", "China", "CoreWeave", "Crusoe", "DeepSeek", "EU", "Google", "Groq", "Lambda", "Meta", "Nebius", "Nvidia", "OpenAI", "TSMC", "United States"]
dates: ["2025-03", "2025-04", "2025-06", "2025-10", "2025-12-24", "2026-01", "2026-01-05", "2026-02-25", "2026-03", "2026-03-16", "2026-03-19", "2026-04-17", "2026-04-27", "2026-05", "2026-05-20", "2026-05-22", "2026-05-31", "2026-06-01", "2026-06-30", "2026-07", "2026-08-18", "2026-08-21", "2026-08-24", "2026-09", "2026-09-16", "2026-09-17", "2026-09-20", "2026-09-21", "2026-09-22"]
keywords: ["acquisition", "amd", "benchmarks", "blackwell", "capex", "compute", "cost", "cpo", "deepseek", "foundry", "fp4", "fp8"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7693, 7745]
section: "15. Hardware & Chips"
sha256: 0c26f664e82014cfc1dfad0e6bb1e312b34c7e066a93ad012cb628b911239bef
---

# January 2026 — CES production confirmations

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

