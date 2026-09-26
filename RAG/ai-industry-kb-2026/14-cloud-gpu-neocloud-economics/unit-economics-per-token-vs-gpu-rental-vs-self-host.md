---
id: ai-industry-kb-2026/14-cloud-gpu-neocloud-economics/unit-economics-per-token-vs-gpu-rental-vs-self-host
title: "Unit economics: per-token vs GPU rental vs self-host"
domain: cloud-gpu-neocloud-economics
role: deep-dive
task: finance
actors: ["Alibaba", "Baseten", "Cerebras", "China", "CoreWeave", "Crusoe", "DeepSeek", "EU", "Fireworks AI", "Groq", "Microsoft", "Mistral", "Moonshot", "OpenRouter", "Z.ai", "vLLM"]
dates: ["2026-05-14"]
keywords: ["gpu", "compute", "consumer", "cost", "deepseek", "glm", "gpus", "ipo", "kimi", "latency", "llama", "mistral"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7090, 7109]
section: "14. Cloud GPU & Neocloud Economics"
sha256: a5fb08ae780f0141eeac6eef44c95939a5f576a5d0f8868a01129780ed91f4d1
---

# Unit economics: per-token vs GPU rental vs self-host

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

