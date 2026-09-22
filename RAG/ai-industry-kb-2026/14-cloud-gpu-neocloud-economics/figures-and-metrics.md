---
id: ai-industry-kb-2026/14-cloud-gpu-neocloud-economics/figures-and-metrics
title: "Figures and metrics"
domain: cloud-gpu-neocloud-economics
role: deep-dive
task: finance
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "Baseten", "Broadcom", "Cerebras", "China", "CoreWeave", "Crusoe", "DeepSeek", "EU", "Fireworks AI", "Google", "Groq", "Lambda", "Meta", "Microsoft", "Mistral", "Moonshot", "Nebius", "Nvidia", "OpenAI", "OpenRouter", "Oracle", "Perplexity", "SpaceX", "Together AI", "Z.ai", "vLLM", "xAI"]
dates: ["2025-07", "2026-05", "2026-05-14", "2026-07", "2026-08", "2026-09"]
keywords: ["accelerator", "amd", "aws", "benchmark", "blackwell", "capex", "compute", "consumer", "cost", "deepseek", "deflation", "disaggregated"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7011, 7170]
section: "14. Cloud GPU & Neocloud Economics"
sha256: c4f1e010f70d8e8b800fb2c3f35425bbb9a1c1b95de707bb5a09af3469fafb61
---

# Figures and metrics

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

**Cerebras (WSE-3 wafer-scale):** ~$0.10/1M (8B) to ~$0.60–$0.85/$1.20 (70B); GPT-OSS-120B $0.35/$0.75; ~1,700–3,000 tok/s (fastest in market for comparable sizes). Free tier 1M tokens/day; paid from $10 deposit. Corporate: **IPO completed May 14, 2026, Nasdaq CBRS — priced at $185/share, $5.55B raised**, the largest IPO of 2026 (valuation reported as ~$56.4B fully diluted at pricing vs ~$95B first-day — sources conflict).

**OpenRouter (aggregator):** routes gpt-oss-120b from ~$0.09/$0.45; 400+ models; 50–1,000 free requests/day tier. 20 providers now serve gpt-oss-120b on OpenRouter (AkashML, CoreWeave, DekaLLM, DigitalOcean, Crusoe, NovitaAI, Mancer + others) — the deepest routing bench of any open-weight model in 2026.

**Modal (serverless GPUs, per-second):** B200 $6.25, H200 $4.54, H100 $3.95, A100-80GB $2.50, L40S $1.95, L4 $0.80, T4 $0.59 (preemptible base; ×3 non-preemptible; ×1.25–2.5 non-default regions). $30/month recurring free compute credit. Cost traps: `min_containers=1` on A100-80GB burns the $30 credit in ~12h; one warm H100 24/7 ≈ $2,800+/mo at base rates.

**Other per-token providers:** Novita AI ~$0.10/$0.50; DeepSeek direct API (post-peak: see §Key dated facts); Mistral Direct Devstral Small 2 $0.10/$0.30 (EU); Replicate H100 $5.49/hr private; Baseten H100 $6.50/hr; Cloudflare Workers AI (Kimi K2.5 $0.60/$3.00 — edge-adjacent, not raw rental).

**Chinese-lab open-weight families (Sep 2026):** Kimi K2.5 cross-provider table — DeepInfra/OpenRouter cheapest at $0.45/$2.25 (cache $0.07), Together $0.50/$2.80, Moonshot direct/Fireworks/Novita/Baseten $0.60/$3.00, Azure $0.60–$0.66/$3.00–$3.30. Kimi K3 (flagship, 2.8T params, 1M context): **$3/$15 per 1M**. Qwen3-Max $0.78/$3.90; Qwen3-Coder-Plus $0.65/$3.25; GLM-5 $0.95/$2.55. DeepSeek R1 $0.70/$2.50. Pattern: Chinese-lab open-weight flagships cluster at **$0.45–$1.00 input / $2.25–$4.00 output** via third-party gateways — undercutting Western frontier APIs 3–10× at comparable capability tiers.

**Pricing regularities:** output tokens cost 3–7× input; batch/async APIs ~50% off real-time (Fireworks, Together); prompt caching cuts input 50–90%+ (Together cached input as low as $0.04/1M).

**gpt-oss-120b (high-reasoning config) via Artificial Analysis (Sep 2026)** — cheapest blended: **CoreWeave $0.04/1M, DeepInfra $0.05/1M, Novita $0.07/1M**; speed leaders: Cerebras 1,763 tok/s, Groq 473.8 tok/s, DeepInfra Turbo 400.4 tok/s — the latency/price stratification measured on one model across 18 providers. **Llama 4 Scout, two tracks:** Groq $0.05/$0.08 (prior-gen LPU) vs Together $0.18/$0.59 (GPU) — ~3.6× input / ~7.4× output undercut by the LPU track.

### Unit economics: per-token vs GPU rental vs self-host

**The crossover rule (serverless vs dedicated GPU):** a managed dedicated H100 at ~$6/hr = **$144/day**. Serverless per-token is cheaper until daily token spend on the same model exceeds ~$144; above that, with steady traffic, dedicated capacity wins on cost + guaranteed availability. On raw rental ($2.89/hr RunPod H100 = $69/day), the crossover drops to roughly **$70/day** — but you then operate the stack (vLLM, autoscaling, monitoring) yourself.

**Rent vs buy:** independent 2026 studies put breakeven at **40–77% sustained GPU utilization** — below it renting wins; above it buying or long-term reserved wins. Most teams measuring real utilization land at 35–55% — renting territory. Consumer cards differ: an RTX 4090 used several hours daily can save $1,000+ over three years vs renting.

**Shopping providers beats everything else:** same H100 SXM5 silicon prices at 2.5× between hyperscaler and long-tail provider; savings from shopping providers routinely **40–60%**. Reserved 3-year long-tail H100 at $1.30–$1.80/hr approaches owned-cluster economics.

**Self-host open-weight (vLLM + quantized GGUF):** a single H100 ($2.89/hr RunPod) serves 70B-class models; quantized 70B at 4-bit needs ~38–40 GB VRAM. Self-host wins at high sustained utilization; per-token inference-as-a-service wins for variable/low volume.

**Worked conversions (GPU-hours → $/1M tokens):**
- Self-hosted 70B on RunPod H100 ($2.89/hr) at 3,000 tok/s sustained: 10.8M tokens/hr → **~$0.27/1M blended** vs $1.04/1M for Llama-3.3-70B-Turbo on Together — ~4× cheaper *if* saturated; at 25% utilization the advantage evaporates ($1.08/1M effective).
- Modal H100 ($3.95/hr) community benchmark: 1,000 samples × 200 tokens on 7–8B BF16 cost **$0.98 in 15 min at 224 tok/s** vs $1.12 on A100-40GB (32 min) and $2.22 on L4 (167 min) — faster GPUs are often cheaper per completed job despite higher $/hr.
- Managed dedicated break-even: dedicated H100-class endpoint at $6/hr = $4,320/month; at Together's $1.04/1M blended for 70B, that's ~4.15B tokens/month (~1,600 tok/s sustained 24/7). Below ~50B tokens/month equivalent, per-token wins.
- The 40–60% provider-shopping rule dominates all of the above: moving an 8-GPU H100 node from $58.4k/mo (hyperscaler ceiling) to $10.5k/mo (long-tail floor) saves more than any throughput optimization.

**One model family, three tracks (DeepSeek-class, per 1M tokens, 2026):** training-contract track — GB200 NVL72, DeepSeek R1 FP4 disaggregated at 16k concurrency, **$0.04/1M blended** (SemiAnalysis InferenceX, May 2026); neocloud inference — B200-served DeepSeek R1 estimated **~$0.20/1M** (SemiAnalysis via tech-insider.org, Sep 2026); serverless direct — DeepSeek V4 API (see peak/off-peak above); serverless gateway — DeepSeek V3 via DeepInfra $0.49/$0.89, via Fireworks $0.56/$1.68. The same weights cost **~40× more per token** through a managed gateway at low concurrency than on contracted NVL72 iron at hyperscale batching.

### Costed scenario playbook (September 2026 rates)

Monthly 24/7 cost = hourly rate × 730.

**Reference: 1× H100 running 24/7 for a month:** Vast.ai (mid) $1.75/hr → ~$1,278; RunPod $2.89 → ~$2,110; Lambda $3.29 → ~$2,402; Modal $3.95 → ~$2,884; AWS (mid) ~$5.00 → ~$3,650; CoreWeave $6.16 → ~$4,497; GCP (A3 Mega) $11.68 → ~$8,526.

**Scenario A — Indie dev, 10M tokens/day** (70% 8B-class @ $0.10/1M + 30% 70B @ $1.04/1M, Together-style pricing): 7M×$0.10 + 3M×$1.04 = **$3.82/day ≈ $115/month** serverless. A dedicated H100 ($2,110/mo) would be 18× more expensive. **Verdict: serverless, no contest.**

**Scenario B — Startup, 500M tokens/day of 70B** at $1.04/1M blended: $520/day = **$15,600/month** serverless. Self-hosted: 500M/day ≈ 5,800 tok/s average → 2× H100 on RunPod ($2.89/hr each) = **$4,220/month** at ~73% utilization. **Verdict: dedicated rental wins ~3.7×;** even a managed dedicated endpoint (~$6/hr/GPU = $8,760/mo for 2×) wins ~1.8×.

**Scenario C — Batch processing, 5B tokens/month, 70B-class:** Fireworks real-time $0.90/1M = $4,500/mo → **batch API (50% off) = $2,250/mo**. Always route non-real-time bulk through batch.

**Scenario D — Fine-tune then serve (70B):** one-time QLoRA run $34–51 on A100 80GB; then serve on 1× H100 ($2,110/mo RunPod) at ~$0.27/1M effective at saturation vs $1.04/1M managed. **Payback of the fine-tune vs managed premium: days** at any serious volume.

**Scenario E — 8-GPU H100 node, 24/7:** cheapest bookable ~$10,500/mo vs hyperscaler ceiling ~$58,400/mo. Provider selection remains the single biggest cost lever in GPU infrastructure.

### The "70% inference spend" contradiction — stated explicitly

- The claim "inference ≈ 70% of AI cloud spend" appears in source briefs but is **CONTRADICTED at global market level** by the most robust audited figure available: **Gartner, August 2026** (AI-optimized IaaS forecast): 2026 total **$42B** (+96% growth), inference **$23.3B** vs training **$19B** — i.e. **~55% inference**, the first year inference exceeds training, forecast at **59% in 2027**.
- Where "~70%" comes from (narrower scopes, secondary compilations): CloudZero (2026): 70–80% of total GPU cloud spend **for teams running AI in production**; ainvest/ai-tool-discovery: 55–80% of enterprise AI GPU spend, "60–80% of AI compute spending in production systems," 2026 forecasts of 70–90%. These are vendor-adjacent compilations, not single audited measurements.
- Compute share ≠ spend share: Deloitte put inference at ~2/3 of AI **compute** in 2026; SeekingAlpha wrote "roughly two-thirds of all AI compute, up from approximately one-third in 2023." Compute share is not cloud-spend share — this distinction is the core of the reconciliation.
- Reconciliation for the knowledge base: global AI-optimized IaaS 2026 — **55% inference** (Gartner, audited forecast, reference value); 2027 — 59%; AI compute share 2026 — ~2/3 (Deloitte); GPU cloud spend among production teams — 70–80% (secondary compilations). The brief's "~70% de la dépense cloud IA" overstates the global figure and conflates compute share, production GPU-spend share, and total-market spend share.
- Acceptable phrasing: "inference exceeded training in 2026 (55% of AI-optimized IaaS per Gartner), and represents ~70% of GPU cloud spend among production AI teams per secondary compilations."

### Neocloud market projection — $35.22B → $236.53B [DIRECTIONAL]

- Mordor Intelligence's neocloud market report estimates **$24.07B in 2025**, expected to reach **$35.22B in 2026**, projecting **$236.53B by 2031** — a **46.37% CAGR over 2026–2031**. Drivers cited: GPU scarcity, hyperscaler and frontier-lab demand, the shift toward gigawatt-scale "AI factories."
- Caveats: this is a **commercial forecast from one research firm**, not a consensus figure or an audited result. Scope definitions vary — Synergy Research Group publishes figures in the same order of magnitude for AI-cloud-adjacent segments but with **different boundaries and timelines** (approaching ~$400B by 2031 in broader scopes). Always cite the source with the number; do not present it as consensus.
- The headline contracts (Meta–CoreWeave $21B, Meta–Nebius up-to-$27B, Microsoft–Nebius $17.4B, Microsoft–Lambda multibillion, Crusoe $1.375B raise) are **commitments and financings, not market-size audits** — they support the direction of the forecast, not its precision.

### Supply/demand trajectory Feb→Sep 2026: the shortage did not ease

- **Structural shortage across HBM, foundry, and power** (sourcebyspec.com, Sep 20, 2026): demand concentrated in inference + training silicon.
- **Lead times (Mar 2026 sourcing data):** new H100 wait times shortened 6+ months → **2–3 months** through 2025; but **Blackwell-class lead times moved the opposite direction: 3–7 months by Q1 2026**; allocation "highly unstable across distributors."
- **HBM is the binding constraint:** DRAM prices rose ~172% through 2025 as HBM production displaced consumer DDR5; NVIDIA raised board-partner kit costs twice in 2026 (May, July). Silicon Analysts estimates B200 manufacturing COGS at **~$6,400**, with HBM3e now **45% of the bill of materials** (vs 41% on H100) — the memory-cost thesis quantified.
- **Purchase-price snapshot (Mar 2026):** L40S 48GB $8,610–$8,900; RTX 6000 Ada 48GB $7,400–$7,800; RTX PRO 6000 96GB $9,450–$9,800; new H100 80GB SXM $25,000–$35,000 (used $18,000–$22,000); A100 80GB discontinued new, $12,000–$18,000 used. Highest-pull SKUs: H200 NVL PCIe, H100 NVL PCIe, L40S, L4, RTX 4000 Ada.
- **B200 residual value: 158% of launch price** (Silicon Data, Sep 16, 2026) — installed Blackwell valued *above* new list ~1 year after mass availability; A100/H100 residuals also above straight-line depreciation.
- **Demand visibility:** NVIDIA record $215.9B FY2026 Data Center revenue (Mar 2026); four largest hyperscalers' combined ~$630–700B 2026 capex plans; OpenAI's ~26 GW secured pipeline (NVIDIA $100B/10 GW, AMD 6 GW, Broadcom); Stargate $500B/10 GW; xAI Colossus 2 (1M-GPU scale target); China's approval of H100 sales licenses (~$25B annual revenue opportunity, Sep 2026 reporting).

### 2026 market trends

- **The split market is the central fact of 2026** (GPUSmith/SemiAnalysis): cheap abundant spot (H100 $1.47–$2.20/hr) coexists with tight committed capacity; 1-year contract pricing rose ~40% Oct 2025–Mar 2026.
- **The two-year H100 deflation (~58% since early 2024) has stalled in 2026** — rates are flat-to-rising, not falling; buyers advised to lock pricing sooner rather than later.
- **HBM memory scarcity, not GPU dies, is the binding price driver** — the transmission from memory contract prices to GPU rental rates is "one of the most important cost signals in the AI economy" (Silicon Data Q1 2026).
- **Blackwell premium dynamics:** each new NVIDIA architecture launches at ~2× the prior generation's median, then compresses as neocloud supply catches up; B200 resale at 158% of launch shows the premium persisting through September 2026.
- **Compute as an asset class:** OCPI listed on Bloomberg; forward curves, utilization indices, and secondary markets (Ornn Compute) now exist; GPU price volatility (B200 +56% annualized) is traded.
- **Training-track demand overhang:** OpenAI 26 GW, Stargate $500B, xAI Colossus 2, Anthropic leasing SpaceX Colossus 1 (220K GPUs, secondary source) — multi-year contracted demand with no short-run price elasticity, behind the committed-capacity price rises; the inference track's elasticity is what enables $0.03/1M tokens.
- **Inference arbitrage is a market of its own:** "LLM providers are commodities, and the only winning strategy is Inference Arbitrage" (2026 practitioner analysis) — smart routers (e.g., Fireworks primary + DeepInfra burst-fallback ~60% cheaper, per Apr 2026 research) treat per-token prices as a tradeable spread.

