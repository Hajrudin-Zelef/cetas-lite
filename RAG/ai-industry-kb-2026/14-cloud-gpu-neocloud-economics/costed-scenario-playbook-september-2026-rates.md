---
id: ai-industry-kb-2026/14-cloud-gpu-neocloud-economics/costed-scenario-playbook-september-2026-rates
title: "Costed scenario playbook (September 2026 rates)"
domain: cloud-gpu-neocloud-economics
role: deep-dive
task: finance
actors: ["AWS", "CoreWeave", "Crusoe", "DeepSeek", "Fireworks AI", "Lambda", "Meta", "Microsoft", "Nebius", "vLLM"]
dates: ["2026-05", "2026-08", "2026-09"]
keywords: ["cost", "aws", "benchmark", "compute", "deepseek", "disaggregated", "fp4", "gguf", "gpu", "gpus", "hyperscaler", "inference"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7110, 7153]
section: "14. Cloud GPU & Neocloud Economics"
sha256: 4382a4a2c74ca904c6d951222e5bb8326c93d352332770f1182077b28fc8e425
---

# Costed scenario playbook (September 2026 rates)

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

