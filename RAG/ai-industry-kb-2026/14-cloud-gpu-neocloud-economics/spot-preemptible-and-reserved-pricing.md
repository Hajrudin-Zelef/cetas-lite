---
id: ai-industry-kb-2026/14-cloud-gpu-neocloud-economics/spot-preemptible-and-reserved-pricing
title: "Spot, preemptible, and reserved pricing"
domain: cloud-gpu-neocloud-economics
role: deep-dive
task: finance
actors: ["AMD", "AWS", "Alibaba", "Baseten", "CoreWeave", "Crusoe", "DeepSeek", "Fireworks AI", "Groq", "Lambda", "Meta", "Microsoft", "Moonshot", "Nebius", "Nvidia", "Oracle", "Perplexity", "Together AI", "Z.ai"]
dates: ["2025-07", "2026-07", "2026-09"]
keywords: ["pricing", "amd", "aws", "benchmark", "blackwell", "compute", "deepseek", "embeddings", "fine-tuning", "fp4", "fp8", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7036, 7089]
section: "14. Cloud GPU & Neocloud Economics"
sha256: 0839064b17bac7e466a8d897712cb0951416a9e97bd4c0918ca714b243a6bc0d
---

# Spot, preemptible, and reserved pricing

**B200 price ladder detail (gpusmith GPU Price Index, crawled ~Aug 30, 2026):** Spheron **$2.71** (lowest live marketplace rate, 8-GPU HGX B200 NVLink 5.0 nodes); Vultr and Hyperbolic **$3.50**; Vast.ai **$5.64**; RunPod **$5.89**; Modal **$6.25**; Lambda **$6.69**; Nebius **$7.15**; Oracle $14.00 (8 GPUs); AWS p6-b200 $14.13 (8 GPUs); Azure NDsrGB200NDRv6 $27.04 (4 GPUs) — a ~10× spread from cheapest marketplace to top hyperscaler tier.
- RunPod B200 product page (updated ~Sep 13, 2026): on-demand **$6.79/hr** (per-second, no commitment); committed **$4.78/hr (6-month)** / **$4.67/hr (1-year)**; serverless B200 workers **$8.64/hr** (pay only while requests run).
- Repricing over one year: Lambda Labs B200 on-demand was **$3.79/hr in July 2025** vs $6.69 in 2026 — ~77% increase as demand outran supply.

**B300 (Blackwell Ultra) — emerging:**
- AIMultiple cohort median **~$7.92/hr**; range $5.44 (Vast.ai) to $18.00 (Oracle Cloud). Spheron $10.21 (Sep 21, 2026); Modal $7.10 (~$0.001972/sec); RunPod $7.89 listed with no capacity.
- GB300 NVL72: contact sales at CoreWeave/Nebius; HGX B300: Nebius **$7.85/hr** (preemptible $4.30), CoreWeave spot $4.48; B300 spot from $5.01/hr (Spheron, Sep 2026).
- B300 rentable on Modal (community benchmark repo, **Sep 2, 2026**): `gpu=B300` arrives as NVIDIA B300 SXM6 AC, ~275 GB memory, compute capability 10.3, PCIe gen 6, scheduled in 1.21 s.

**MI300X (192 GB HBM3) — the memory alternative:**
- Vultr: **from $2.00/GPU/hr** (MI300X and MI325X; fractional $2.29, full $3.50; preemptible-style instances cheaper still).
- Spheron: **$3.59/hr on-demand** (Sep 21, 2026), no spot tier. Crusoe: $3.45/hr (Aug 21, 2026).
- The ~$2.40/hr figure cited in some briefs is consistent with Vultr's $2.00–$3.50 range — right ballpark, provider-dependent. On some marketplaces MI300X is the *more expensive* hourly rate than H100; the premium must be earned through memory (192 GB vs 80 GB), not raw kernel throughput. CUDA/ROCm tuning maturity still favors NVIDIA for custom-kernel work.
- MI300X pricing is **softening as MI350X/MI355X roll out** (dstack, Sep 2026).

**Prior generations and workstation cards:**
- A100 SXM4 (80 GB): OCPI settled **$0.99/hr** (Sep 21); neocloud band ~$1.79; Spheron $1.48 (spot $1.10); RunPod $1.59; Modal $2.50 (80 GB) / $2.10 (40 GB). Standard for 70B QLoRA fine-tuning (~$34–$51 per full run).
- RTX 4090: median **$0.52/hr** — cheapest training-class card on the AIMultiple index.
- RTX PRO 6000 (Blackwell workstation): Spheron $2.28 (spot $1.18); RunPod $2.09; CoreWeave $2.50; Modal $3.03.
- L40S: Modal $1.95. L4: Modal $0.80. T4: Modal $0.59.

**Marketplace price points snapshot (madebyagents.com, daily-updated, Sep 22, 2026)** — low-anchored floors, volatile by the hour; treat as dated floor indicators, not guarantees [COMMUNITY]:
- H200 NVL (RunPod Community): $0.50/hr ($365/mo). H100 NVL (Vast.ai spot): $1.00/hr ($730/mo). H100 SXM (Vast.ai spot): $1.07/hr ($779/mo). H200 SXM (Vast.ai spot): $1.32/hr ($961/mo). H100 PCIe (RunPod Community): $1.99/hr ($1,453/mo).
- AMD MI325X (DigitalOcean On-Demand): $3.80/hr ($2,774/mo). AMD MI355X (DigitalOcean Spot): $4.50/hr ($3,285/mo).
- B200 (Vast.ai spot): $5.00/hr ($3,650/mo). B300 (RunPod Community): $6.94/hr ($5,066/mo).
- RTX PRO 6000 Blackwell (GPU Mart monthly): $1.09/hr ($799/mo).
- Broader ranges (gpusmith PDF, July 2026): A100 **$0.13–$5.04/hr**, B200 **$2.69–$16.11/hr** — the spread is the market. 7× spread on identical silicon: cheapest H100 (Vast.ai $1.73) vs Oracle OCI bare-metal ($10.00).

### Spot, preemptible, and reserved pricing

- **Spot discounts of 25–45%** vs on-demand are standard where offered: H100 spot $1.47 (Vast.ai) / $2.10–$2.20 (Spheron) vs $2.64–$2.89 on-demand; H200 spot $2.77 vs $4.80 (Spheron); B200 spot $5.37 vs $9.30 (Spheron); A100 spot $1.10 vs $1.48.
- **Modal** inverts the framing: published base rates are preemptible (spot); **non-preemptible execution costs 3×** base, and region-pinned workloads add 1.5–1.75×. Combined multipliers can reach ~3.75× the headline rate.
- **Reserved/committed tiers** cut 30–50% below on-demand at the same provider (Mercatus, May/Sep 2026): hyperscalers $2.20–$3.80/hr (1-yr) / $2.20–$3.00 (3-yr); Tier-1 specialty $2.00–$2.80 (1-yr) / $1.70–$2.30 (3-yr); long-tail/regional $1.60–$2.10 (1-yr) / **$1.30–$1.80/hr (3-yr)** — close to owned-and-amortized economics.
- Marketplace reliability caveat: Vast.ai/Spheron hosts are independent operators — hosts can vanish mid-job; checkpointing is mandatory for long runs.

### Serverless inference: per-token pricing (September 2026, USD per 1M tokens in/out)

**Fireworks AI** (ex-Meta PyTorch team; Cursor, Perplexity, Sourcegraph customers):
- Size-tiered serverless (published): <$4B $0.10/1M; 4–16B $0.20/1M; >16B dense $0.90/1M; MoE ≤56B $0.50/1M; MoE 56.1–176B $1.20/1M.
- Named: DeepSeek V3 family $0.56/$1.68; DeepSeek V4 Flash **$0.14/$0.28** (did not inherit DeepSeek's Aug-16 peak regime — useful for seeing how gateways absorb lab repricing); gpt-oss-20b $0.07/$0.30; gpt-oss-120b $0.15/$0.60; DeepSeek-V4-Pro $1.74/$3.48 (projected).
- Batch API: 50% of serverless (~24h turnaround); cached input tokens 50%. Dedicated GPU deployments (per-second): A100-80GB $2.90/hr; H100/H200 $6–7/hr; B200 $9–10/hr; B300 $12/hr.
- Differentiators: function-calling reliability (92%+ multi-tool accuracy), P99 latency only 3.9× P50, FP8/FP4 quantization tiers, SOC2/HIPAA. $1 free starter credit.

**Together AI:** gpt-oss-20b $0.05/$0.20; Llama-3.3-70B-Instruct-Turbo $1.04/$1.04; DeepSeek-V4-Pro $2.10/$4.40 (projected); Qwen3.6-Plus $0.50/$3.00; Llama 4 Scout $0.18/$0.59; Muse-Glimmer-30B $0.35/$1.50 with $0.04 cached input; embeddings from $0.02/1M. 200+ open models, fine-tuning (LoRA/full), dedicated endpoints billed per GPU-hour.

**DeepInfra — the flat-rate counter-position (Sep 2026):**
- Bills **one flat rate around the clock**: V4-Pro **$1.30/$2.60** (cached input $0.10); V4-Flash **$0.08/$0.18** (cached input $0.016); flex tier at **0.8×** (Pro run at $2.08/1M output).
- **Cut rates mid-Aug 2026**: an llmgateway audit (~mid-Aug) found live metadata at $1.30/$2.60 vs catalogue $1.74/$3.48 for Pro, and $0.08/$0.18 vs $0.14/$0.28 for Flash — catalogues carrying launch prices overstate by ~a third.
- gpt-oss-120b three speed tiers on identical weights: **$0.037/$0.17 base**, **$0.15/$0.60 Turbo**, **$0.20/$0.95 Ultra** — the market's cleanest latency-priced spread.
- Price floor (Sep 2026): gpt-oss-20b **$0.03/$0.14**; gpt-oss-120b **$0.04/$0.17**; Llama-3.1-8B **$0.03/$0.05**; DeepSeek-V4-Pro $1.30/$2.60 vs **$1.74/$3.48 at Together, Fireworks, and Baseten**; GLM-5.2 $0.75/$2.40 vs $1.40/$4.40 at Fireworks/Novita; **Kimi K3 undercuts the $3.00/$15.00 everyone else charges**.
- Scale facts [VENDOR] (company-published, Sep 2026): founded 2022 by the team behind imo (200M+ users); owns its own GPU infrastructure across **9 datacenters**; processes **~5 trillion tokens/week**; **$107M Series B in 2026 with NVIDIA among the investors**; dedicated GPUs at **$0.89/A100-hr and $2.20/H100-hr**, billed by the minute. Rate-limit caveat: ~200 concurrent requests per account — 429s under spike traffic.

**Groq (LPU):** Llama 3.1 8B Instant $0.05/$0.08 (~840 tok/s); gpt-oss-120b $0.15/$0.60 (~478–493 tok/s); Llama 3.3 70B $0.59/$0.79 (~294–750 tok/s). Small catalog (~6 models), no self-serve fine-tuning. GroqCloud console fixtures (crawled ~Sep 17, 2026): gpt-oss-20b $0.075/$0.30 (1,000 tok/s); Llama 3.1 8B and Llama 3.3 70B moved to **enterprise/"Contact Sales"** tiering — Groq pushing flagships upmarket while gpt-oss OSS weights remain the self-serve front door. No public per-token pricing for Groq-3-LPX-served models.

