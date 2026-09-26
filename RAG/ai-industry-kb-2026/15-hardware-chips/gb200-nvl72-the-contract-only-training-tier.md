---
id: ai-industry-kb-2026/15-hardware-chips/gb200-nvl72-the-contract-only-training-tier
title: "GB200 NVL72: the contract-only training tier"
domain: hardware-chips
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Anthropic", "Broadcom", "Cerebras", "CoreWeave", "Crusoe", "DeepSeek", "Google", "Groq", "Lambda", "Meta", "Microsoft", "Nebius", "Nvidia", "OpenAI", "Samsung", "TSMC"]
dates: ["2025-04", "2025-06", "2025-10", "2025-12-24", "2026-03", "2026-03-16", "2026-04-17", "2026-05", "2026-05-22", "2026-06-30", "2026-07", "2026-08-18", "2026-08-21", "2026-08-24", "2026-09-22"]
keywords: ["training", "acquisition", "asic", "aws", "benchmarks", "blackwell", "compute", "decode", "deepseek", "disaggregated", "fp4", "fp8"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7491, 7538]
section: "15. Hardware & Chips"
sha256: f66a00b234e98c902beb4e50c93288a355c1bda2c5c6fb86289cd20aea494845
---

# GB200 NVL72: the contract-only training tier

- **June 2025 (Advancing AI):** Helios first teased.
- **January 5/6, 2026 (CES, Lisa Su inaugural keynote):** detailed Helios reveal — "world's best AI rack," rack on stage. **72 MI455X**, 31 TB HBM4, 1.4 PB/s, 2.9 FP4 EF / 1.4 FP8 EF [VENDOR]; EPYC Venice (Zen 6, 256 cores/512 threads, 2 nm), Pensando Vulcano NICs; liquid-cooled, ~7,000 lbs; **launches later in 2026**. Also: MI440X (enterprise 8-GPU), MI500 preview (2027, CDNA 6, 2 nm, HBM4E).
- **2026 rentals:** MI300X at Crusoe $3.45/hr (Aug 21), Vultr from $2.00/hr, Spheron $3.59/hr (no spot tier) — bookable but thinly listed vs NVIDIA SKUs. MI355X "contact sales" only [UNVERIFIED for rental availability].

### GB200 NVL72: the contract-only training tier

- **2024–2025:** GB200 NVL72 first shipped 2024; GB300 NVL72 (Blackwell Ultra, 288 GB HBM3e/GPU, ~50% perf uplift over GB200) shipped 2025.
- **March 2026 survey:** GB200 NVL72 rents at **$10.50–$27/GPU-equivalent/hr** ($756–$1,944/hr per full rack); CoreWeave $27.00, Nebius $22.50, Lambda $22.50, Crusoe $16.50 (cheapest enterprise listing); hyperscalers $18+/hr on-demand, contract-only.
- **August 21, 2026 snapshot:** CoreWeave GB200 NVL72 repriced to **$10.50/GPU-hr on-demand** — dispersion and downward repricing through 2026; GB300 NVL72 = contact sales.
- **May 22, 2026 (SemiAnalysis InferenceX):** DeepSeek R1 FP4 disaggregated on GB200 NVL72 reaches **$0.04–$0.08/1M tokens at high concurrency** (16,130 concurrent requests → $0.04/1M; 4 → $10.12) — the per-token collapse curve at hyperscale batching.
- Purchase price: **$2–3M per GB200 NVL72 rack**; not rentable as single GPUs (exposed at Superchip or rack level).

### NVIDIA–Groq: the LPU inside the stack

- **December 24, 2025:** NVIDIA signs **~$20B non-exclusive technology licensing** with Groq; Jonathan Ross and Sunny Madra join NVIDIA; Groq stays independent under CEO Simon Edwards. Corrects "acquisition" phrasing — this is the template for how incumbents absorb disruptive inference hardware.
- **March 16, 2026 (GTC):** Groq 3 LPU unveiled as the 7th Vera Rubin chip; LPX rack = 256 interconnected LPUs (Samsung manufacturing); 128 GB SRAM per rack, 40 PB/s on-chip bandwidth, 640 TB/s scale-up, 315 PFLOPS FP8. Internal codename LP30; Groq skipped a "Groq 2."
- **August 24, 2026 (Hot Chips):** **Groq 3 LPX in full production**; Nebius first cloud adopter (Token Factory). Artificial Analysis measured Gemma 4 31B at 100K context — **3,431 output tok/s** (fastest ever recorded); NVIDIA SPEED-Bench median 4,767 tok/s. NVIDIA claims 35× tokens-per-watt improvement of Rubin GPUs paired with LPX racks.
- Design split: **Vera Rubin NVL72 handles prefill/context and KV-cache construction; Groq 3 LPX handles decode.** Ian Buck: "The LPU is only one-five-hundredth of the [GPU's] capacity per chip, but the bandwidth is exceptional."

### Cerebras: public since May 2026

- **April 17, 2026:** new S-1 filed (Nasdaq CBRS) after the 2024 attempt was withdrawn over CFIUS review of the G42 stake (since divested).
- **May 13–14, 2026: IPO priced at $185/share (above the $150–160 range); $5.55B raised — the largest IPO of 2026.** Reported valuation conflict: ~$56.4B fully diluted at pricing vs ~$95B first-day — flagged, sources conflict.
- WSE-3 (5 nm): 4 trillion transistors, 900,000 AI cores, 44 GB on-chip SRAM, 125 petaflops; CS-3 system ~$2–3M estimated, 23 kW liquid-cooled; MemoryX + SwarmX for wafer-scale clusters. AWS Marketplace integration (GA, EDP-billable). Reported: OpenAI $20B partnership for inference capacity [SECONDARY, not primary-confirmed].
- **July 2026 (Artificial Analysis):** Gemma 4 31B at **1,990.8 tok/s**, Llama 4 Scout >2,600 tok/s — ~19× the fastest GPU solutions tested. Independent academic modeling counterpoint: B200 offers 1.5–3× better performance-per-watt-per-dollar at scale.

### Etched Sohu: shipping, not yet independently benchmarked

- **June 30, 2026:** stealth exit; Sohu transformer-inference ASIC + rack-scale system; TSMC N4P first-pass silicon success; >$1B in signed customer contracts claimed (described as "contracts and demand," not recognized revenue).
- **July 2026:** first rack delivered to **Jane Street** and deployed into production ("pleased with the early results"); Jane Street is also a major investor.
- **August 18, 2026:** **$700M raised at $21B valuation** — doubled in under a month from the $10.3B Series C (July); led by Jane Street with Kleiner Perkins, Sequoia, a16z, Tiger Global; **$1.9B raised to date**; >400 employees. Taiwan factory opened; gigawatt-scale capacity target by 2027.
- **As of September 22, 2026: no third-party benchmarks published** — all throughput/latency/efficiency claims are company-reported. Analyst caution (Cerity Partners, Aug 2026): "semiconductor history is littered with brilliant chips that never became great businesses."

### Broadcom custom XPU and the broker layer

- **July 2026 (reported):** Broadcom custom AI XPU business at **$20B+ in FY2025**; Q1 FY2026 AI semiconductor revenue $8.4B (+106% YoY); Q2 FY2026 record $22.19B total (+48% YoY); customers Google, Meta, OpenAI simultaneously.
- Brokers/new entrants 2026: **GPUaaS.com** (wholesale broker, "~30% under hyperscale," quotes <24h, no buyer fees), **Bitdeer** (Bitcoin-miner-turned-AI-cloud, H100–GB200), **WhiteFiber** (NYC HPC cloud), **Thunder Compute** (H100 $1.38/hr on-demand, May 2026), **Hyperbolic/Hyperstack** (B200 floor $5.99/hr; H100 $1.90/hr).
- Nebius names CoreWeave, Crusoe, and Lambda as key competitors in filings — the neocloud tier now defines its own competitive set independent of hyperscalers.

### Hyperscaler silicon milestones

- **April 2025 (Google Cloud Next):** TPU v7 "Ironwood" announced — 2026 is the production ramp year. Google–Anthropic commitment: up to **1M TPU units / 1 GW** of dedicated compute through 2026 (announced October 2025) — the largest custom-silicon commitment by an AI lab to date.
- **2026:** Amazon Trainium 3 volume production; EC2 Trn2/Trn3 instances; Anthropic and OpenAI workloads validated on Trainium 3.
- **2026:** Microsoft Maia 200 ("Braga") deployed in Azure data centers for Azure AI inference; limited select-customer access via dedicated inference instances.
- Note on source quality: several 2026 "silicon shift" pieces (financialcontent tokenring) are low-quality aggregators with no named primary source — used only as [SECONDARY] directional color.

### Supply, demand, and the unabated shortage

