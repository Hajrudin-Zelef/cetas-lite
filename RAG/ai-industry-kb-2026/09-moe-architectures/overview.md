---
id: ai-industry-kb-2026/09-moe-architectures/overview
title: "9. MoE Architectures"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["Alibaba", "China", "DeepSeek", "Google", "MiniMax", "Moonshot", "OpenAI", "SGLang", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2025-08-05", "2026-04-24", "2026-07-27", "2026-09"]
keywords: ["moe", "attention", "benchmark", "compute", "cost", "deepseek", "fp4", "glm", "gpu", "inference", "kimi", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4578, 4592]
section: "9. MoE Architectures"
sha256: 7f521b51382dcae337128cf10fa2ae26e6df502995241b11d6db42772d3502c9
---

# 9. MoE Architectures
Keywords: mixture of experts, MoE, sparse activation, open-weight, expert routing, fine-grained experts, shared experts, top-k routing, latent routing, load balancing, DeepSeekMoE, DeepSeek-V4, Kimi K3, Qwen3.8-Max, MiniMax M3, GLM-5.3-Flash, MiMo-V2.6, Qwen3.6-35B-A3B, Gemma 4, gpt-oss, Stable LatentMoE, Quantile Balancing, expert parallelism, Mixture of Universal Experts, FLAME, hybrid attention, Mamba-2 hybrid, native multimodality, auxiliary-loss-free balancing, sparsity ratio

## Summary

- From 2025-08-05 (OpenAI gpt-oss, 117B/5.1B) to September 2026, open-weight MoE releases scaled from ~100B-class to the 2.8T/104B Kimi K3 (weights 2026-07-27) — the largest open-weight release ever, a 1.56 TB checkpoint — with total parameter counts now routinely in the hundreds of billions to trillions paired with per-token active counts of 3B–104B.
- The uniform 2026 formula: MoE decouples capacity (total params, paid once in VRAM) from compute (active params, paid per token). A 428B/23B model does roughly the per-token FLOPs of a ~20B dense model while drawing on a 428B knowledge pool. Every 2026 flagship ships with 1M-token context.
- **Correction 1 (editorial scope):** "2026 = the year MoE becomes the default architecture on all serious models" is too absolute. The accurate statement: MoE became the default architecture for **very large open-weight models** (the 300B+ class and every 2026 frontier open release). Dense counterexamples persist: Gemma 4 ships a 31B dense sibling, Qwen3.8 has a 27B dense tier, gpt-oss-20b serves the edge as dense-scale 21B/3.6B, the 2026 controlled ~150M-active study had dense beating MoE 3× at small scale, and the architectures inside undisclosed closed labs remain unknown.
- **Correction 2 (fine granularity and bandwidth):** fine-grained experts (128–896 small experts, top-8 to top-16) genuinely improve routing specialization — the combinatorics go from C(16,2)=120 achievable combinations to C(64,8)≈4.4 billion, and the ESFT paper measures task routing concentrating more strongly with finer granularity — but published systems research (MoE Parallel Folding, arXiv 2504.14960v2) shows fine-grained MoE has **lower training efficiency than coarse-grained MoE** across tested parallelism strategies: more experts and more active experts per token raise dispatch volume, shrink GEMM efficiency, and grow activation memory. "Fine granularity without hurting inter-GPU bandwidth" is **[UNVERIFIED / contradicted as stated]**; bandwidth is preserved only by co-designed mitigations.
- Kimi K3 is the counter-model that prices and pays the bandwidth cost instead of denying it: Stable LatentMoE projects the 7,168-wide hidden state into a 3,584-wide latent space before dispatch (roughly halving routed-expert activation traffic and expert compute), Quantile Balancing replaces fixed-step bias adjustment with a single global all-reduce, and MXFP4 experts (E8M0 scale per 32 weights) compress the 2.72T routed parameters.
- DeepSeek V4 (2026-04-24) is the flagship DeepSeek-style MoE: V4-Pro 1.6T/49B on 33T tokens, V4-Flash 284B/13B on 32T tokens, 1M context, 384K max output, dual thinking modes — and the architecture reconstruction ([PARTIALLY VERIFIED], secondary sources only, no official report directly verified): 256/384 routed experts with 6 active per token + 1 shared, first three blocks on deterministic token-ID hash routing, auxiliary-loss-free balancing retained with a small sequence-wise balance loss, routing affinity changed to sqrt(softplus(...)), routed expert weights deployed in FP4.
- Architectural consensus for 2026 open MoE: fine-grained expert pools (128–896), always-on shared experts (1–2; MiMo and GLM-5.3-Flash notably use 0), auxiliary-loss-free load balancing (DeepSeek-V3-style dynamic bias), hybrid attention stacks (linear + gated + sliding-window + global) to make 1M context servable, MTP training heads doubling as speculative decoders.
- The 2026 open-weight MoE leadership map is China: DeepSeek (Hangzhou), Alibaba/Qwen (Hangzhou), Moonshot (Beijing), Z.ai (Beijing), MiniMax (Shanghai), Xiaomi (Beijing) produced every frontier open MoE of 2026; Google (Gemma 4) is the only non-Chinese lab with a 2026 open MoE release; OpenAI's gpt-oss (2025) is the Western baseline. No Western lab has released a 100B+ open MoE in 2026.
- Cross-references to sibling part 09b: training economics, inference economics (active-vs-total split, KV-cache nuance), the serving cookbook (vLLM/SGLang EP, EPLB), quantization recipes, pricing tables, the license landscape, and the benchmark scoreboard are consolidated there. This part covers releases, specifications, routing theory, and the leadership map.

