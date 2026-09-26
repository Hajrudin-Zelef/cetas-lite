---
id: ai-industry-kb-2026/15-hardware-chips/main-actors
title: "Main actors"
domain: hardware-chips
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Anthropic", "Broadcom", "Cerebras", "China", "CoreWeave", "Crusoe", "DeepSeek", "Fireworks AI", "Google", "Groq", "Lambda", "Meta", "Microsoft", "Nebius", "Nscale", "Nvidia", "OpenAI", "Samsung", "TSMC", "United States"]
dates: ["2026-02", "2026-05", "2026-05-14", "2026-05-22", "2026-05-31"]
keywords: ["accelerator", "acquisition", "amd", "asic", "aws", "benchmarks", "blackwell", "capex", "compute", "cost", "datacenter", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7627, 7692]
section: "15. Hardware & Chips"
sha256: 309a56166ceb70d95db5787b0e5bb51e6bdfdc97b7b2cfb9db67fa4fa318f659
---

# Main actors

- **Price leaders:** DeepInfra, Hyperbolic, Novita — cheapest $/1M on identical open weights; floor $0.03/1M input (gpt-oss-20b, DeepInfra). Trade-off: concurrency caps and 429s under spikes.
- **Latency leaders:** Groq (prior-gen LPU: $0.05/$0.08 for 8B), Cerebras (WSE: $0.10 for 8B, but 1,990+ tok/s on 31B-class). Trade-off: higher $/1M at 70B class vs price leaders.
- **Enterprise gateways:** Fireworks, Together — higher per-token prices but primary-routing reliability, private endpoints, VPC/on-prem options (documented arbitrage pattern: Fireworks primary + DeepInfra burst-fallback, Apr 2026).
- **Direct labs:** DeepSeek ($0.07/$0.28 V4), OpenAI ($0.10/$0.50 gpt-oss-20b) — price-setters; third-party gateways undercut or overcut them depending on quantization and batching.
- **Cost-replication thesis:** Chinese-lab flagships at $0.45–$1.00 input / $2.25–$4.00 output per 1M via third-party gateways undercut Western frontier APIs 3–10×.

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

