---
id: ai-industry-kb-2026/15-hardware-chips/figures-and-metrics
title: "Figures and metrics"
domain: hardware-chips
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "Broadcom", "Cerebras", "China", "CoreWeave", "Crusoe", "DeepSeek", "Fireworks AI", "Google", "Groq", "Meta", "Microsoft", "Nebius", "Nvidia", "OpenAI", "Samsung", "TSMC"]
dates: ["2026-05-14", "2026-05-22"]
keywords: ["accelerator", "amd", "asic", "benchmark", "benchmarks", "blackwell", "compute", "consumer", "cost", "datacenter", "deepseek", "dram"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7558, 7632]
section: "15. Hardware & Chips"
sha256: cc2506a0ee5cb84080e008f9cae3c00a0a0a3473b2872a84539802875907e0c4
---

# Figures and metrics

## Figures and metrics

### Rack-scale system specs (company-stated [VENDOR] unless noted)

| System | Rack config | Stated performance | Memory | Interconnect | Status Sep 22, 2026 |
|---|---|---|---|---|---|
| NVIDIA Vera Rubin NVL72 | 72 Rubin GPUs + 36 Vera CPUs | 3.6 EF FP4 / 1.2 EF FP8 (~3.3× GB300 NVL72) | 75 TB HBM4, 13 TB/s | NVLink 6, 260 TB/s | Full production; cloud cohort deploying H2 2026; no public rental price |
| NVIDIA GB200 NVL72 | 72 B200 + 36 Grace | 1.44 EF FP4 (sparsity) | 13.4–13.5 TB unified | NVLink 5, 130 TB/s | Shipping 2024–2026; $10.50–$27/GPU-hr; contract-only tier |
| NVIDIA GB300 NVL72 | 72 B300 + 36 Grace | ~50% perf uplift over GB200 | 288 GB HBM3e/GPU | NVLink 5 | Shipping since 2025; contact sales |
| AMD Helios | 72 MI455X + EPYC Venice | 2.9 FP4 EF / 1.4 FP8 EF | 31 TB HBM4, 1.4 PB/s | Pensando Vulcano (scale-out) | Detailed CES 2026; launches later 2026 [VENDOR] |
| Rubin Ultra NVL576 ("Kyber") | 576 Rubin GPU dies | 15 EF FP4 / 5 EF FP8 | HBM4e | NVLink 7 | 2027 plan; SemiAnalysis reported possible slip to 2028; NVIDIA says roadmap intact [SECONDARY] |

- Architectural convergence is the real story: both camps settled on **72-accelerator liquid-cooled racks, HBM4, and rack-as-the-unit-of-compute**. Peak FLOPS numbers are not head-to-head comparable (different precisions and measurement conditions) — present side by side, never as a benchmark ranking.
- Physical engineering (Rubin): 1.3M components, ~1,300 chips per single-wide 3rd-gen MGX rack (~4,000 lbs); NVLink spine with 5,000 copper cables (2+ miles); cable-free, hose-free, fanless compute trays; 47-minute truck-to-power-on claimed; dry closed-loop liquid cooling; 800VDC sidecar power. ~1.36 metric tons and ~120 kW for GB200 NVL72 (direct-to-chip liquid cooling mandatory).

### Inference silicon specs

| Chip | Architecture | Key specs | Status |
|---|---|---|---|
| NVIDIA Groq 3 LPX rack | 256 LPU chips (Samsung) | Per-chip: 500 MB SRAM, 150 TB/s SRAM BW, 2.5 TB/s scale-up; per-rack: 128 GB SRAM, 40 PB/s, 640 TB/s, 315 PFLOPS FP8 | Full production announced Aug 24, 2026; Nebius first adopter |
| Cerebras WSE-3 | Wafer-scale (5 nm) | 4T transistors, 900K AI cores, 44 GB SRAM, 125 petaflops; CS-3 ~$2–3M est., 23 kW | Public since May 14, 2026 IPO |
| Etched Sohu | Transformer ASIC (TSMC N4P) | Rack-scale inference system; SOTA claims on Llama/DeepSeek/Qwen/Mamba [company only] | Shipping since Jun–Jul 2026; no independent benchmarks |
| Google TPU v7 "Ironwood" | Hyperscaler ASIC | 4.6 PFLOPS dense FP8, 192 GB HBM3E, 7.37 TB/s; 9,216-chip Superpod → 42.5 EF [SECONDARY] | Ramp 2026 |
| Amazon Trainium 3 | Hyperscaler ASIC (3 nm) | 2.52 PFLOPS MXFP8, 128–144 GB HBM3E; EC2 Trn2/Trn3 | Volume production 2026; validated by Anthropic + OpenAI |
| Microsoft Maia 200 "Braga" | Hyperscaler ASIC (3 nm) | 216 GB HBM3e at 7 TB/s; 3× FP4 perf of Trainium 3 [VENDOR] | Azure deployment 2026, select-customer inference instances |
| Meta MTIA 500 | Hyperscaler ASIC (3 nm roadmap) | Internal recommendation-model benchmarks — not comparable to LLM inference figures | Internal-only, no external availability |

### Pricing: purchase, rental, residual

- Rubin NVL72 rack: **$7–8M** (one buyer, third-party; [UNVERIFIED single secondary]); earlier press $5–7M; vs ~$5M for a GB300 rack and $2–3M for a GB200 rack.
- GB200 NVL72 rental: **$10.50–$27/GPU-equiv/hr** (Mar 2026 survey); CoreWeave repriced to **$10.50** on-demand (Aug 21, 2026). Hyperscalers $18+/hr on-demand, contract-only. Wholesale/secondary $7.20–$10.00/GPU-hr.
- MI300X rental: Crusoe **$3.45/hr** (Aug 21, 2026); Vultr **from $2.00/hr** (fractional $2.29, full $3.50); Spheron **$3.59/hr on-demand** (no spot tier) (Sep 21, 2026).
- HGX B300: Nebius $7.85/hr (preemptible $4.30), CoreWeave spot $4.48; B300 spot from $5.01/hr (Spheron, Sep 2026). H100 spot $1.38–$1.90/hr (Thunder Compute, Hyperstack).
- **B200 residual: 158% of launch** (Sep 16, 2026) — above new list ~1 year after mass availability.
- **Rubin NVL72 rental pricing: no public $/GPU-hr listing as of Sep 22, 2026** — all access contract/quote-based → [UNVERIFIED] for public rental rates. MI355X rental: "contact sales" only → [UNVERIFIED].

### Per-token price floor (verified, Sep 2026)

- **DeepInfra gpt-oss-20b: $0.03/$0.14** per 1M (input/output); gpt-oss-120b $0.04/$0.17; Llama-3.1-8B $0.03/$0.05 — "a few cents per million" verified for the small-model class.
- GroqCloud prior-generation LPU pricing: $0.05/$0.08 (8B-class), $0.59/$0.79 (70B-class) — no public Groq-3-served per-token pricing found; first production route is enterprise/cloud (Nebius), not the public API.
- SemiAnalysis InferenceX (May 22, 2026): DeepSeek R1 FP4 on GB200 NVL72 at $2.21/GPU-hr TCO → **$0.04–$0.08/1M tokens at high concurrency**; the concurrency curve (4 → $10.12; 24 → $2.23; 180 → $0.53; 2,253 → $0.08; 16,130 → $0.04) shows why per-token prices collapse at hyperscale batching.
- SemiAnalysis estimate: ~$0.20/1M tokens for B200-served DeepSeek R1 (Sep 2026). DeepSeek V4 direct: $0.07/$0.28.
- Caution: cheapest providers carry rate-limit traps — DeepInfra caps ~200 concurrent requests; hitting ceilings yields 429s under spike traffic.

### Market structure figures

- Inference share of AI compute: **~1/2 (2025) → ~2/3 (2026)** (Deloitte, Nov 2025).
- AI processor vendors: **151 companies, 290+ products** (JPR Q2 2026, Jul 2026).
- Inference-optimized chips: **>$50B in 2026**; frontier-class chips **>$200B** (early-2026 analyst analyses).
- Committed-capacity prices: **+~40%, Oct 2025–Mar 2026** (tech-insider.org, Sep 2026).
- Blackwell-class lead times: **3–7 months** (Q1 2026); H100 waits: 2–3 months.
- Cerebras IPO: **$185/share, $5.55B raised** (largest IPO of 2026); valuation reported $56.4B (at pricing, fully diluted) vs ~$95B (first-day) — conflict flagged.
- Etched: **$700M at $21B** (Aug 18, 2026), $1.9B raised to date, >$1B signed contracts claimed (not recognized revenue).
- Broadcom: custom XPU **$20B+ FY2025**; Q1 FY2026 AI semi revenue $8.4B (+106% YoY).
- NVIDIA–Groq: **~$20B non-exclusive licensing** (Dec 24, 2025). NVIDIA FY2026 Data Center revenue: **$215.9B** (Mar 2026).
- Rubin GPUs 2026: **~5.7M units** sell-side estimate [UNVERIFIED]. Viral "1,000 racks/day" — **do not cite** (no NVIDIA statement, no named analyst).
- Rubin GPUs 2026: **~5.7M units** sell-side estimate [UNVERIFIED]. Viral "1,000 racks/day" — **do not cite** (no NVIDIA statement, no named analyst).
- DRAM prices rose ~172% through 2025 as HBM displaced consumer DDR5; NVIDIA raised board-partner kit costs twice in 2026 (May, July).

### Consumer-GPU and low-cost tier (Wave 1 §1.6, Sep 2026, cross-referenced)

- **RTX 4090:** median **$0.52/hr** — cheapest training-class card on the AIMultiple index; workhorse of Vast.ai-style marketplaces for quantized inference.
- **RTX PRO 6000 (Blackwell workstation):** Spheron $2.28/hr (spot $1.18); RunPod $2.09/hr — 2026 bridge product between consumer and datacenter Blackwell.
- **L40S / L4 / T4:** Modal $1.95 / $0.80 / $0.59 per hr — classic inference-serving SKUs for smaller models and embedding workloads.
- **Marketplace reliability caveat:** Vast.ai/Spheron hosts are independent operators — hosts can vanish mid-job; checkpointing is mandatory. This tier is for fault-tolerant batch, not production serving.
- **Strategic read:** the consumer tier is where the inference track's price elasticity is highest. It does not serve the training track at all (no NVLink domains, no InfiniBand).

### Serverless provider tiering (per-token, 2026)

- **Price leaders:** DeepInfra, Hyperbolic, Novita — cheapest $/1M on identical open weights; floor $0.03/1M input (gpt-oss-20b, DeepInfra). Trade-off: concurrency caps and 429s under spikes.
- **Latency leaders:** Groq (prior-gen LPU: $0.05/$0.08 for 8B), Cerebras (WSE: $0.10 for 8B, but 1,990+ tok/s on 31B-class). Trade-off: higher $/1M at 70B class vs price leaders.
- **Enterprise gateways:** Fireworks, Together — higher per-token prices but primary-routing reliability, private endpoints, VPC/on-prem options (documented arbitrage pattern: Fireworks primary + DeepInfra burst-fallback, Apr 2026).
- **Direct labs:** DeepSeek ($0.07/$0.28 V4), OpenAI ($0.10/$0.50 gpt-oss-20b) — price-setters; third-party gateways undercut or overcut them depending on quantization and batching.
- **Cost-replication thesis:** Chinese-lab flagships at $0.45–$1.00 input / $2.25–$4.00 output per 1M via third-party gateways undercut Western frontier APIs 3–10×.

