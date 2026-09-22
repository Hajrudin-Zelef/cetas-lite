---
id: ai-industry-kb-2026/15-hardware-chips/key-dated-facts
title: "Key dated facts"
domain: hardware-chips
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Anthropic", "Broadcom", "Cerebras", "CoreWeave", "Crusoe", "DeepSeek", "Google", "Groq", "Lambda", "Meta", "Microsoft", "Nebius", "Nscale", "Nvidia", "OpenAI", "Samsung", "SpaceX", "TSMC", "United States", "xAI"]
dates: ["2025-03", "2025-04", "2025-06", "2025-10", "2025-12-24", "2026-01", "2026-01-05", "2026-02-25", "2026-03", "2026-03-16", "2026-03-19", "2026-04-17", "2026-04-27", "2026-05", "2026-05-20", "2026-05-22", "2026-05-31", "2026-06-01", "2026-06-30", "2026-07", "2026-07-21", "2026-07-22", "2026-08-18", "2026-08-21", "2026-08-24", "2026-09-16", "2026-09-17", "2026-09-20", "2026-09-22", "2026-10"]
keywords: ["acquisition", "amd", "asic", "aws", "benchmarks", "blackwell", "claude", "compute", "cost", "decode", "deepseek", "disaggregated"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7471, 7557]
section: "15. Hardware & Chips"
sha256: 0e23fa77995e681bf30b118ca551662dc1107545ae677c01a672438603bd178a
---

# Key dated facts

## Key dated facts

### Vera Rubin: announcement → production → deployments

- **March 2025 (GTC):** Vera Rubin platform announced; rack named **NVL144** (die-count naming).
- **January 5, 2026 (CES):** Jensen Huang declares Rubin in "full production"; all six chips back from manufacturing; customer deployments guided H2 2026. TSMC N3P process; HBM4 from Samsung and SK Hynix; supply chain of >350 factories across 30 countries.
- **January 2026:** Early cloud cohort named — AWS, Google Cloud, Microsoft Azure, OCI, plus neoclouds CoreWeave, Lambda, Nebius, Nscale. Microsoft was already testing first racks; deployments planned for Georgia and Wisconsin data centers and the Fairwater AI Superfactory.
- **February 25, 2026 (FY26 Q4 earnings call):** first **VR200 samples** shipped to customers.
- **March 16, 2026 (GTC):** "seven new chips in full production" (six Rubin-platform chips + Groq 3 LPU); ~$1T Blackwell+Rubin bookings visibility through 2027 claimed; **Rubin CPX NVL rack put on hold/cancelled**, replaced by Groq 3 LPX.
- **May 20, 2026 (FY27 Q1 earnings call):** production shipments begin **fiscal Q3 (August–October 2026)**; volume ramp in fiscal Q4.
- **May 31, 2026:** **Dell delivers the world's first production Vera Rubin NVL72 rack** (PowerEdge XE9812) to CoreWeave; Michael Dell announced with photos and diagnostic logs.
- **June 1, 2026:** **CoreWeave completes industry-first rack-scale operational validation** — delivery-to-production in under 6.5 hours at Livingston, NJ. Rack = 72 Rubin GPUs + 36 Vera CPUs, NVLink 6 (3.6 TB/s per GPU, 260 TB/s scale-up), 1.6 Tb/s backend per GPU.
- **July 21, 2026:** Ian Buck (NVIDIA hyperscale/HPC VP): "We are absolutely in full production right now. All the major customers are deploying these systems" — OpenAI, Azure, Google Cloud, CoreWeave, Meta, Dell named.
- **July 22, 2026:** Vera Rubin NVL72 confirmed in full production; OpenAI to deploy at scale in Q3 2026 (Bloomberg-sourced).
- **H2 2026:** Nebius offers Vera Rubin NVL72 in the US and Europe via AI Cloud and Token Factory; complements existing GB200/GB300 NVL72 capacity.
- **~September 17, 2026:** CoreWeave links hundreds of Rubin GPUs into a multi-rack Spectrum-X cluster (cross-ref wave2.1/06; not re-verified here).
- **September 22, 2026:** no cloud provider had published hourly rental rates for Vera Rubin NVL72 — access remains contract/quote-only.

### AMD: Helios and the 2026 lineup

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

- **Q1 2026 sourcing:** new H100 wait times 6+ months → **2–3 months**; Blackwell-class lead times moved the opposite way: **3–7 months**; allocation "highly unstable across distributors."
- **September 16, 2026 (Silicon Data):** **B200 residual value at 158% of launch price** — the market values installed Blackwell above new list ~1 year after mass availability; A100/H100 residuals also above straight-line depreciation.
- **September 20, 2026:** structural shortage confirmed across HBM, foundry, and power; Jensen Huang says enterprises and clouds buy GPUs for both training and inference, requiring "sustained, high-volume compute capacity."
- Utilization paradox: field reports of ~9% average GPU utilization (single jobs pinned to whole A100s); the 2026 fix is GPU sharing (MIG partitions, time-slicing, Karpenter autoscaling on spot), not more hardware.
- Demand wall: OpenAI ~26 GW secured pipeline (NVIDIA $100B/10 GW, AMD 6 GW, Broadcom); Stargate $500B/10 GW; xAI Colossus 2 (1M GPUs); NVIDIA FY2026 record **$215.9B** Data Center revenue (Mar 2026).
- Inference price floor verified (Sep 2026): DeepInfra gpt-oss-20b **$0.03/$0.14** per 1M (input/output); gpt-oss-120b $0.04/$0.17; Llama-3.1-8B $0.03/$0.05; SemiAnalysis: ~$0.20/1M for B200-served DeepSeek R1. Inference costs fell **280-fold in two years** (Jan 2026, citing Deloitte/SDxCentral) — Jevons paradox in effect.
- **Export-control hardware events (see §18):** DOJ/Super Micro indictment unsealed March 19, 2026 (Manhattan federal court); NDRC prohibited Meta–Manus April 27, 2026.

### Additional dated datapoints

- **GB200 vs 8×B200 decision rule (2026):** for models under ~100B parameters, 8×B200 HGX nodes are "almost always sufficient and far more cost-effective"; NVL72 earns its keep at 200B+ parameter training needing the 130 TB/s all-to-all domain, or 671B-scale reasoning models that must hold the full model in one rack.
- **AMD rental strategy note:** MI300X's value prop in rental is memory-per-dollar (192 GB vs 80 GB H100 — a 70B FP16 model fits on one card), not kernel throughput; ROCm tuning maturity still trails CUDA for custom-kernel work.
- **Cerebras serving modes:** **native** (model fits on wafer — full speed advantage) vs **Weight Streaming** (larger models stream weights from MemoryX — narrows the advantage); AWS Marketplace integration reached GA (provision via console, bill against EDPs; pay-per-token and dedicated reservations).
- **GroqCloud continuity:** GroqCloud per-token pricing in 2026 ($0.05/$0.08 for 8B-class, $0.59/$0.79 for 70B-class) reflects the **prior LPU generation**; the LPX generation's first production route is enterprise/cloud (Nebius), not the public API.
- **October 2025:** Google–Anthropic up-to-1M-TPU / 1-GW commitment — frames why Rubin faces non-NVIDIA competition at the frontier-training tier. **Reported:** Anthropic leasing SpaceX Colossus 1 (220K GPUs) with Claude rate limits doubled on the back of it [SECONDARY].
- **xAI Colossus 2:** 1M-GPU scale target (→ §14 for capacity detail).
- **March 2026 purchase-price snapshot (sourcebyspec.com):** L40S 48GB $8,610–$8,900; RTX 6000 Ada 48GB $7,400–$7,800; RTX PRO 6000 96GB $9,450–$9,800; new H100 80GB SXM $25,000–$35,000 (used $18,000–$22,000); A100 80GB discontinued new, $12,000–$18,000 used. Highest-pull SKUs: H200 NVL PCIe, H100 NVL PCIe, L40S, L4, RTX 4000 Ada.
- **Spelling and structure corrections:** brief wrote "Sohgo" → verified product name **Sohu**; brief's "Groq acquisition" → **non-exclusive license**.

