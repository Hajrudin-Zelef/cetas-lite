---
id: ai-industry-kb-2026/09-moe-architectures/serving-procurement-guidance-consolidated-dated
title: "Serving procurement guidance (consolidated, dated)"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "China", "DeepSeek", "Huawei", "MiniMax", "Moonshot", "Nvidia", "SGLang", "Unsloth", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2026-05-22", "2026-06", "2026-07", "2026-08-16", "2026-08-19", "2026-09", "2026-09-09", "2026-10-28"]
keywords: ["agentic", "agents", "amd", "apache", "ascend", "attention", "awq", "benchmarks", "compute", "consumer", "cost", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5317, 5356]
section: "9. MoE Architectures"
sha256: a2645f828373077e855751320ecac8eb27de7d2034f808290577fff8c8e3480e
---

# Serving procurement guidance (consolidated, dated)

- **The scarce resource shifted from FLOPs to VRAM + KV cache.** Active params keep falling (3B active can now do agentic coding at 73% SWE-bench), but total params keep rising (2.8T), so serving cost is dominated by weight residency and KV memory — hence the parallel wars in quantization (FP8/NVFP4/MXFP4) and attention (MLA/GQA/SWA/linear hybrids).
- **1M context is table stakes for open MoE flagships** (all 2026 releases), but only economically servable because of attention-side innovations — not MoE itself. This is the single most load-bearing economic fact in the consolidation: MoE never reduced KV cache; every KV saving is from attention design.
- **The total/active ratio is the master pricing variable.** API price ladders, self-host $/MTok ballparks, and quantization footprints all track active compute monotonically; capability per dollar tracks post-training. Any cost model built on total params alone (the number in the headline) misprices by 5.9× (the measured residency gap at matched active params).
- **Quantization is now a first-class release artifact, not an afterthought.** FP8 ships as the default checkpoint (GLM-5.3-Flash), NVFP4/MXFP4 are native formats (M3, K3, gpt-oss), and the most-downloaded variant of Qwen3.8-Max is FP8 — fetchers vote for serving with their downloads. But quantization never changes the total/active ratio: it shrinks the whole MoE, including its VRAM bill.
- **Expert parallelism is a commodity config, not a research project.** EP moved from pretraining-systems lore (DeepEP, DualPipe) into fine-tuning YAML configs (Axolotl June 2026), framework primitives (torchtune's DTensor expert-parallel plan), auto-config (DeepSpeed AutoEP), and serving flags (`--enable-expert-parallel`) — with documented topology rules (TP for latency at low concurrency, DP+EP for throughput at high concurrency, skip EP below 3% activation density).
- **The fine-granularity correction reframes vendor roadmaps.** More experts is not free: dispatch volume, GEMM efficiency, activation memory, and MP-group communication all degrade with finer granularity unless co-designed mitigations (DeepEP, latent routing, quantile balancing, topology-aware parallelism) are budgeted. Kimi K3 is the proof that fine granularity is *payable*, not free.
- **China leads the open-weight MoE frontier in 2026** (DeepSeek, Alibaba/Qwen, Moonshot, Z.ai, MiniMax, Xiaomi) with day-0 SGLang/vLLM support and co-designed serving stacks — the model and the serving system are now designed together. Domestic-chip milestones (V4 on Huawei Ascend first; GLM-5.3-Flash's stealth week on Chinese chips; K3 on AMD MI355X) mean serving economics are no longer NVIDIA-exclusive.
- **Licensing liberalized around the frontier.** Apache 2.0 and MIT now cover the majority of open-weight MoE releases; custom licenses cluster at the capability-leading models. For RAG operators this means the cheapest long-context serving tier (Flash-class models at $0.075–$0.15/M input) is also the most permissively licensed.
- **For training-system buyers:** attach hardware + dataset + batch + rank + baseline to every speedup/VRAM figure before acting; treat vendor headline multipliers as the top of a range, the component-level numbers (NVIDIA collab blog) as the plannable part, and independent comparisons (torchtune paper, MarkTechPost, HF blog) as the calibration. The September 2026 MoE kernel correctness fix is a reminder that "production-solid" claims about fast-moving MoE paths age in weeks.
- **Open research economics:** dynamic expert count (confidence-gated routing, early-2026 work suggesting 20–40% compute savings); MoUE universal experts at 2T+ scale; principled MoE-aware KV schemes (expert-conditioned KV sharing) as the one untapped KV axis; NVMe-bandwidth routing prediction for 100B+ MoE on consumer hardware; MoE on the edge via MXFP4 + offloading.

### Serving procurement guidance (consolidated, dated)

- **Size the fleet on total params (quantized), the latency SLO on active params.** The 5.9× residency gap at matched active params (OLMoE profile) means any cost model built on the headline number misprices by that factor.
- **Concurrency topology:** ≤128 concurrent requests → TP for latency (+40–86% throughput); ≥512 → DP+EP for throughput (+16–47%); crossover ~256–512; skip EP entirely below 3% activation density (−7–12% measured). MLA/MQA-family models → DP+EP mandatory to avoid KV duplication.
- **Format selection by tier (2026 practice):** 300B+ flagship serving → FP8 default (GLM-5.3-Flash) or NVFP4 (MiniMax M3); local/workstation → AWQ/Marlin 4-bit (90 GB → 26 GB VRAM, 2.1× throughput, −75% TTFT in the documented Mixtral case); consumer GPU → expert offloading to NVMe (Gemma 4 blueprint: ~4.5 GB hot, ~24 MB/forward); CPU/offload → GGUF Q4 and below.
- **KV is the binding constraint, not experts.** Invest in prefix caching, HybridKV-style compression (up to 7.9×), attention-efficient model selection, and MTP drafter heads (2.0–2.6× accepted-length speedup) before buying more GPU memory for the experts themselves.
- **Cache-hit economics:** DeepSeek-V4-Pro cache-hit input pricing ($0.003625/M) and GLM-5.3-Flash cached pricing ($0.03/M) make prefix-cache-friendly RAG pipelines — repeated system prompts, retrieved-context prefixes — the dominant cost lever, ahead of model choice.
- **Time-sensitive economics:** promo windows move fast (GLM-5.3-Flash $0.075 promo ended 2026-09-09; V4-Pro 75% discount made permanent 2026-05-22; peak/off-peak tiers from 2026-08-16). Any fixed price cited from before mid-2026 is stale.

### Open research questions (economic framing)

- **Dynamic expert count:** can top-k adapt per token (easy tokens → fewer experts) without destabilizing training? Early 2026 work on confidence-gated routing suggests 20–40% compute savings — the largest untapped per-token axis.
- **Universal experts at scale:** MoUE's virtual-width idea is proven at research scale; does it hold at 2T+ params and 1M context?
- **KV cache for 1M-context MoE:** all 2026 KV savings are attention-side; is there a principled MoE-aware KV scheme (e.g., expert-conditioned KV sharing) that beats MLA/GQA/SWA hybrids? This is the one axis where MoE itself could finally reduce KV cost.
- **Routing for agents:** phase-aware and step-granularity routing is promising for multi-step tool use but unproven in production serving stacks.
- **MoE on the edge:** expert offloading (Gemma 4) and MXFP4 (gpt-oss) show the path; whether 100B+ MoE becomes viable on consumer hardware depends on NVMe-bandwidth routing prediction, still an open problem.
- **Latent routing diffusion:** Kimi K3's Stable LatentMoE (hidden 7,168 → latent 3,584) is the pattern to watch spreading to other labs — routing in a narrower space halves dispatch traffic and expert compute at the cost of a projection.

### Watch items (late 2026)

- **Latent-space routing diffusion:** whether other labs adopt Kimi K3-style latent routing (hidden → narrower latent before dispatch) to pay for finer granularity.
- **NVFP4 maturity:** Axolotl's NVFP4 MoE LoRA (July 2026) and Unsloth's NVFP4 export (July 2026) are parallel tracks; watch for NVFP4 *training* claims (all current claims are export-and-serve).
- **EP as a standard fine-tuning primitive:** Axolotl's DeepEP integration (June 2026) and torchtune's expert-parallel plan suggest EP is leaving the pretraining world — watch for it in mainstream trainers.
- **MoE-aware KV:** the one untapped KV axis — whether expert-conditioned KV sharing can beat the current MLA/GQA/SWA-hybrid stack.
- **Dynamic UD quants:** Unsloth Dynamic v3.0 (launched 2026-08-19) is current; GLM-5.2's Dynamic GGUF secondary figures (2-bit ~239 GB / ~82% retention) need independent benchmarks.
- **PyTorch 2.15 (scheduled 2026-10-28):** watch for further silent behavioral changes after 2.14's clamp-gradient incident.

## Sources and URLs (continued)

