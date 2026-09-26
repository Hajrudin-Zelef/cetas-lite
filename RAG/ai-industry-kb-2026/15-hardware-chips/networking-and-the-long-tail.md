---
id: ai-industry-kb-2026/15-hardware-chips/networking-and-the-long-tail
title: "Networking and the long tail"
domain: hardware-chips
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "China", "CoreWeave", "DeepSeek", "Google", "Groq", "Lambda", "Meta", "Microsoft", "Nebius", "Nscale", "Nvidia", "OpenAI", "SGLang", "Samsung", "TSMC", "United States", "vLLM"]
dates: ["2025-03", "2026-01", "2026-01-05", "2026-02-25", "2026-03-16", "2026-03-19", "2026-04-09", "2026-04-27", "2026-05-20", "2026-05-22", "2026-05-31", "2026-06-01", "2026-07-21", "2026-07-22", "2026-09-17", "2026-09-22", "2026-10"]
keywords: ["acquisition", "agentic", "amd", "aws", "blackwell", "compute", "consumer", "cost", "cpo", "decode", "deepseek", "disaggregated"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7441, 7490]
section: "15. Hardware & Chips"
sha256: 0a612b6ebbc6d0b428f2fe3a71b26aecd509d188059122260d6819fd9213d0b9
---

# Networking and the long tail

- Inference became the dominant workload: **half of all AI compute in 2025 → two-thirds in 2026** (Deloitte, Nov 2025 report); Lenovo's CEO forecasts the historical ~80/20 training/inference spend split inverting to ~20/80.
- Two hardware regimes, two economics: training = raw FLOPs, NVLink domains, $10+/GPU-hr multi-year contracts, NVIDIA-dominated; inference = $/token, latency, tokens-per-dollar-per-watt, served increasingly by LPUs, wafer-scale chips, and custom ASICs.
- Jon Peddie Research Q2 2026: **151 companies, 290+ AI processor products** — disaggregation reduces the odds any single architecture dominates all segments.
- The shortage did not ease: Blackwell-class lead times **3–7 months** (Q1 2026); H100 wait times 2–3 months; **B200 residual value at 158% of launch price** (Sep 16, 2026); binding constraints at HBM, foundry capacity, and power.

### Networking and the long tail

- **Scale-up (inside the rack):** NVLink 5 → NVLink 6 (3.6 TB/s per GPU, 260 TB/s per Rubin rack); 5,000-copper-cable NVLink spine; 6th-gen NVLink switch as a first-class Rubin component.
- **Scale-out (between racks):** ConnectX-9 SuperNICs (1.6 Tb/s backend per GPU at CoreWeave), BlueField-4 DPUs, Spectrum-X 102.4T CPO Ethernet; NIXL as the default KV-block transport for disaggregated serving.
- **Consumer-GPU ultra-low-cost tier:** RTX 4090 at a median **$0.52/hr**; RTX PRO 6000 (Blackwell workstation) ~$2.09–$2.28/hr — the bottom rung of the inference track, for fault-tolerant batch, not production serving.
- Broker/new-entrant layer: GPUaaS.com (~30% under hyperscale, wholesale model), Bitdeer, WhiteFiber, Thunder Compute (H100 $1.38/hr), Hyperbolic/Hyperstack (B200 floor $5.99/hr; H100 $1.90/hr).

### Geopolitical hardware events (cross-refs)

- **DOJ / Super Micro (March 19, 2026)** — indictment of three individuals for diverting ≥$2.5B of AI servers (A100/H100) to China, 2024–2025. Super Micro itself is not a defendant. Full treatment in §18.
- **NDRC / Meta–Manus (April 27, 2026)** — China blocked Meta's ~$2B acquisition of Manus/Butterfly Effect (Index No. 000013039-2026-00026), ordered unwinding: first AI-sector use of China's foreign-investment security review; "Singapore-washing" doctrine defeated. Full treatment in §18.
- **Geopolitical hardware events (cross-refs)**
- **DOJ / Super Micro (March 19, 2026)** — indictment of three individuals for diverting ≥$2.5B of AI servers (A100/H100) to China, 2024–2025. Super Micro itself is not a defendant. Full treatment in §18.
- **NDRC / Meta–Manus (April 27, 2026)** — China blocked Meta's ~$2B acquisition of Manus/Butterfly Effect (Index No. 000013039-2026-00026), ordered unwinding: first AI-sector use of China's foreign-investment security review; "Singapore-washing" doctrine defeated. Full treatment in §18.
- **Capacity megadeals** (CoreWeave–Meta ~$21B, April 9, 2026; Nebius–Meta up to $27B, March 16, 2026; OpenAI deploying Rubin at scale Q3 2026) are detailed in §14 — this section covers hardware milestones only.

### Competitive and economic framing

- **Prefill/decode is the 2026 serving pattern and now the hardware pattern:** vLLM/SGLang fleets disaggregate prefill, decode, and encoder pools with KV blocks moved by NIXL; NVIDIA productizes it as silicon — Rubin NVL72 for prefill/context, Groq 3 LPX for decode. NVIDIA claims 35× tokens-per-watt for the pairing.
- The bifurcation has a price signature: **training $/GPU-hr is flat-to-rising; inference $/token is collapsing.** The same weights cost ~40× more per token through a managed gateway at low concurrency ($1.68/1M) than on contracted NVL72 iron at hyperscale batching ($0.04/1M) — the spread is the business model of the inference track and the reason "inference arbitrage" routing exists.
- Agentic workloads consume **100×–1,000× more tokens per task** than a single chat turn (iterative reasoning, tool calls, re-querying); enterprise adoption is still in "early innings" (legal, finance, healthcare) — demand growth is not flattening even as unit shipments rise (sourcebyspec.com, Sep 20, 2026).
- **Power is the third constraint:** if scaling trends persist, a single location hosting a major training run could need up to **8 gigawatts by 2030** — "the output of eight nuclear reactors." The physical basis of the dual-track problem: a vast network of efficient inference chips alongside a shrinking number of hyper-powerful, hyper-expensive training centers.
- **First Rubin deployments are contracted to neoclouds** (CoreWeave, Nebius, Lambda, Nscale), not hyperscalers — a shift in NVIDIA's go-to-market (→ §14).
- **Measured inference economics on NVL72 (SemiAnalysis InferenceX, May 22, 2026):** GB200 NVL72 at $2.21/GPU-hr TCO reaches $0.04–$0.08/1M tokens at high concurrency (DeepSeek R1 FP4, disaggregated prefill/decode, Dynamo TRT-LLM).

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

