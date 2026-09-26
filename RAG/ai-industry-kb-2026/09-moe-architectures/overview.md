---
id: ai-industry-kb-2026/09-moe-architectures/overview
title: "9. MoE Architectures"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["Alibaba", "China", "DeepSeek", "Google", "Huawei", "MiniMax", "Moonshot", "OpenAI", "OpenRouter", "SGLang", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2025-08-05", "2025-12-16", "2026-04-16", "2026-04-24", "2026-06-01", "2026-06-13", "2026-07-16", "2026-07-27", "2026-08-03", "2026-08-25", "2026-08-28", "2026-09"]
keywords: ["moe", "agent", "agentic", "apache", "ascend", "attention", "benchmark", "compute", "cost", "deepseek", "fp4", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4578, 4615]
section: "9. MoE Architectures"
sha256: 4407ff44edfd4562a39429654f6b09fa0291be65f105d96569e42f5d188f2642
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

## Key dated facts

### Release timeline: the open-weight MoE explosion (2025-08 → 2026-09)

| Date | Model | Total / Active | Context | License | Significance |
|---|---|---|---|---|---|
| 2025-08-05 | OpenAI gpt-oss-120b / 20b | 117B / 5.1B ; 21B / 3.6B | 128K | Apache 2.0 | First OpenAI open weights since GPT-2 (2019); MXFP4 native; template every 2026 release follows |
| 2025-12-16 | Xiaomi MiMo-V2-Flash | 309B / 15B | 256K | MIT | First 300B-class open MoE; 150 tok/s via MTP self-speculative decoding; day-0 SGLang support |
| 2026-04-16 | Alibaba Qwen3.6-35B-A3B | 35B / 3B | 262K (1M w/ YaRN) | Apache 2.0 | Hybrid Gated DeltaNet + gated attention; 256 experts (8+1 shared); 73.4% SWE-bench Verified; runs on RTX 4090 |
| 2026-04-24 | DeepSeek V4-Pro / V4-Flash | 1.6T / 49B ; 284B / 13B | 1M | MIT (later sources; April coverage cited Apache 2.0) | Largest open MoE of H1 2026; 33T-token pretraining; launched on Huawei Ascend chips first |
| ~2026-04 | Google Gemma 4 26B-A4B | 25.2B / 3.8B | 256K | Apache 2.0 | Pure-attention MoE (no SSM); 128 experts (8+1 shared); the Western open-MoE anchor of 2026 |
| 2026-06-01 | MiniMax M3 | 428B / 23B | 1M | MiniMax Community License | First open-weight model combining reasoning + agentic coding + native multimodality (text/image/video from step 0) |
| 2026-06-13/17 | Z.ai GLM-5.2 | undisclosed (base reused in GLM-5.3: 753B) | 1M | custom | Coding-plan flagship; coding agent flagship tier |
| 2026-07-16/27 | Moonshot Kimi K3 | 2.8T / 104B | 1M | open-weight | Largest open-weight model ever released (1.56 TB checkpoint); 896 experts (16+2 shared); MoonEP + FlashKDA open-sourced |
| 2026-08-03/08 | Alibaba Qwen3.8-Max | 2.4T / 95B | 1M | non-Apache open license | First Max-class Qwen with open weights; 512 experts (10+1 shared); 92 layers |
| 2026-08-25/26 | Z.ai GLM-5.3-Flash | 320B / 18B | 1M | MIT | First natively multimodal GLM-5; 288 experts (8 active); stealth-launched as "Ox Alpha" on OpenRouter |
| 2026-08-28 | Z.ai GLM-5.3 (flagship) | 753B (MoE base from 5.2) | 1M | custom glm-5.3 license | Post-training upgrade of 5.2 base; text-only |
| ~2026-09 | Xiaomi MiMo-V2.6-Flash / Pro | 309B / 15B ; 1.02T / 42B | 1M | MIT | Native omnimodal (text/image/video/audio); 5:1 SWA/GA hybrid; 5-layer MTP drafter |

- Naming convention, now standard: "A3B"/"A4B" suffixes denote **active** parameters per token (35B-A3B = 35B total, 3B active); Qwen3.5/3.6/3.7/3.8 are successive flagship generations, not point releases of Qwen3; Gemma 4's "E" prefix (E2B/E4B) denotes effective parameters for the dense siblings.

### Release-event granularity: announcement vs weights vs GA

