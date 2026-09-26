---
id: ai-industry-kb-2026/14-cloud-gpu-neocloud-economics/sources-and-urls
title: "Sources and URLs"
domain: cloud-gpu-neocloud-economics
role: deep-dive
task: finance
actors: ["Alibaba", "Cerebras", "China", "CoreWeave", "DeepSeek", "Fireworks AI", "Groq", "Meta", "Microsoft", "Moonshot", "Nebius", "Nvidia", "OpenAI", "Z.ai", "vLLM", "xAI"]
dates: []
keywords: ["cost", "deepseek", "deflation", "glm", "gpu", "gpus", "hbm", "hyperscaler", "inference", "ipo", "kimi", "latency"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7282, 7295]
section: "14. Cloud GPU & Neocloud Economics"
sha256: 4c9800651abc145404b50832e7fa23355b82e7c2b11d42b8c490ae9cecf84e88
---

# Sources and URLs

1. **For inference buyers (Sep 2026):** default to per-token serverless for variable workloads; the decision that saves the most money is *which provider tier* (neocloud vs hyperscaler), not which GPU — a 40–60% swing for identical silicon. Budget for rate-limit engineering on price-leader providers (200 concurrent caps, 429s under spikes).
2. **For self-hosters:** a single rented H100/H200 with vLLM + quantized open-weight models is the price/performance sweet spot (~$0.27/1M tokens effective at saturation vs ~$1.04/1M managed); go dedicated (rented GPU) once sustained token spend exceeds ~$70–144/day depending on managed vs raw.
3. **For capacity planners:** 2026 is not a buyer's market for *committed* capacity — lock 1–3 year rates if utilization justifies it; use spot/marketplace for everything fault-tolerant. The 40% committed-price run-up (Oct 2025–Mar 2026) is the warning. Note that Nebius now runs capacity *auctions* — the committed tier is price-discovered upward.
4. **For DeepSeek-direct users:** direct pricing is no longer a single number. Peak windows (01:00–04:00, 06:00–10:00 UTC) bill up to 3.55× the old flat rate; flat-rate gateways (DeepInfra) are a genuine hedge. Treat any undated DeepSeek price as stale.
5. **For RAG/knowledge-base purposes:** GPU prices are volatile (±5% daily moves observed); any cost model should reference a dated index value (OCPI) rather than a static figure, refreshed at least monthly. Inference $/token is a *flow* price that reprices monthly; GPU $/hr is increasingly a *contract* price. Any cost model mixing the two without dates is wrong.
6. **Watch:** B200 price trajectory (+14.3%/30d — if neocloud supply catches up, expect compression toward the H200 band through 2027); HBM supply expansion (the true price driver); Cerebras post-IPO (CBRS) as a signal for LPU economics going mainstream; Cerebras WSE-3 loses on perf-per-dollar at scale despite single-stream latency wins (academic modeling) — track tokens-per-dollar-per-watt, not peak FLOPS.
7. **Demand-side overhang:** OpenAI's 26 GW secured pipeline, Stargate's $500B/10 GW buildout, xAI Colossus 2 (1M GPUs), Microsoft's $33B+ in neocloud commitments — multi-year demand visibility means no near-term return to 2024-style deflation for contracted capacity.
8. **Decision shortcut:** variable/low volume → per-token serverless (Fireworks/Together/Groq); steady high volume → dedicated GPU (rented H100/H200); strategic/long-term → 1–3 yr reserved neocloud or buy; latency-critical → LPU (Cerebras/Groq); non-real-time bulk → batch APIs at 50% off.
9. **Chinese-lab price pressure is structural:** Kimi/Qwen/GLM/Zhipu open-weight flagships at $0.45–$1.00 input / $2.25–$4.00 output per 1M via third-party gateways undercut Western frontier APIs 3–10× at comparable capability tiers — this compresses the "frontier premium" that justified hyperscaler inference pricing.
10. **Data freshness protocol:** this file's prices are anchored to September 13–22, 2026 sources. Re-verify against OCPI (data.ornn.com) and the Mercatus GPU Index before any procurement decision; treat any undated $/hr figure in the RAG as suspect.
11. **Strategic read on the contract wave:** neoclouds are now contracting at hyperscaler scale (>$20B single-customer commitments), and Meta is diversifying across multiple neoclouds (CoreWeave $21B+, Nebius up-to-$27B) rather than concentrating. The first Vera Rubin deployments are contracted to neoclouds, not hyperscalers — a shift in NVIDIA's go-to-market and in the hierarchy of cloud computing toward purpose-built GPU clouds.

## Sources and URLs

