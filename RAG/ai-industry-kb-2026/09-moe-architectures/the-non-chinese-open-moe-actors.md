---
id: ai-industry-kb-2026/09-moe-architectures/the-non-chinese-open-moe-actors
title: "The non-Chinese open-MoE actors"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Huawei", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "Unsloth", "Xiaomi", "Z.ai"]
dates: ["2025-08-05", "2025-12-16", "2026-04-16", "2026-04-24", "2026-05", "2026-06-01", "2026-06-13", "2026-07-27", "2026-08-03", "2026-08-25", "2026-08-28"]
keywords: ["moe", "agentic", "agents", "amd", "apache", "ascend", "attention", "benchmarks", "consumer", "deepseek", "fine-tuning", "fp8"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4850, 4893]
section: "9. MoE Architectures"
sha256: 7907db5ae625452a3b6e7a357f8c132171914e64f1d33348f2ec7ea68645e810
---

# The non-Chinese open-MoE actors

- **DeepSeek (Hangzhou):** V4-Pro 1.6T/49B + V4-Flash 284B/13B (2026-04-24); DeepSeekMoE fine-grained segmentation (2024) is the innovation the whole 2026 fleet inherits; DeepSeek-V3 (2024-12, 671B/37B) set the reference for auxiliary-loss-free dynamic-bias load balancing, MTP heads, and multi-plane FP8 training; DeepEP is the co-designed communication library the field's EP practice derives from. Launched V4 on Huawei Ascend first — no NVIDIA required.
- **Alibaba / Qwen (Hangzhou):** Qwen3.6-35B-A3B (2026-04-16, Apache 2.0, the 3B-active coding MoE that runs on an RTX 4090); Qwen3.8-Max 2.4T/95B (2026-08-03/08, first Max-class Qwen with open weights); Gated DeltaNet linear-attention lineage; 33T-token-scale pretraining culture.
- **Moonshot AI (Beijing):** Kimi K3 2.8T/104B (weights 2026-07-27) — largest open-weight release ever; Stable LatentMoE, Quantile Balancing, Kimi Delta Attention, AttnRes, SiTU; open-sourced MoonEP + FlashKDA + AgentEnv alongside the weights; AMD MI355X serving benchmarks.
- **Z.ai (Beijing):** GLM-5.2 (2026-06-13/17) → GLM-5.3-Flash 320B/18B MIT (2026-08-25/26, the "Ox Alpha" stealth launch on OpenRouter served entirely on domestically produced Chinese chips) → GLM-5.3 flagship 753B custom license (2026-08-28); glm5_next hybrid sparse+linear attention + mHC; first natively multimodal GLM-5.
- **MiniMax (Shanghai):** M3 428B/23B (2026-06-01, MiniMax Community License) — first open release combining reasoning, agentic coding, and native text/image/video multimodality from step 0; MiniMax Sparse Attention (MSA) block-sparse indexer with its co-designed kernels.
- **Xiaomi / MiMo (Beijing):** MiMo-V2-Flash 309B/15B (2025-12-16, first 300B-class open MoE, 150 tok/s via MTP self-speculative decoding) → MiMo-V2.6-Flash/Pro 309B/15B + 1.02T/42B (~2026-09, native omnimodal incl. audio); day-0 SGLang culture (inference code contributed upstream).

### The non-Chinese open-MoE actors

- **Google:** Gemma 4 26B-A4B (~2026-04, Apache 2.0) — the only non-Chinese 2026 open MoE release; pure-attention (no SSM) design; ships alongside dense E2B/E4B/31B siblings; the 31B dense model doubles as the dense-counterexample in the leadership map.
- **OpenAI:** gpt-oss-120b/20b (2025-08-05, Apache 2.0) — the Western open-MoE baseline and template-setter (sparse MoE + native MXFP4 + permissive license); first open weights since GPT-2 (2019).
- **NVIDIA:** Nemotron 3.5 Lightning (30B/3B MoE+Mamba-2 hybrid, $0.22/M output efficiency reference point) [VENDOR]; the NVIDIA technical blog's dense-vs-MoE active-parameter analysis is the most-cited neutral systems reference of the year.
### The undisclosed tier: what the leadership map does not cover

- The leadership map in this part is strictly an **open-weight** map. The architectures inside closed labs (OpenAI, Anthropic, Google's non-Gemma flagships) remain undisclosed — no statement in this part extends "MoE won" or "dense won" to them.
- Consequence for the RAG: "default architecture" claims must always carry their scope (open-weight, ≥300B-class, 2025-08 → 2026-09). Unscoped versions are how the too-absolute editorial synthesis entered the draft.
- **Western non-actors (2026):** no Western lab released a 100B+ open MoE in 2026 — a structural fact of the leadership map, not a claim about closed-lab architectures (undisclosed).

### Release counts by lab (2025-08 → 2026-09) and quarterly cadence

- Release events per lab: Z.ai 3 (GLM-5.2, GLM-5.3-Flash, GLM-5.3); Alibaba 2 (Qwen3.6, Qwen3.8-Max); Xiaomi 2 (MiMo-V2-Flash, MiMo-V2.6); DeepSeek 1 event / 2 variants (V4-Pro, V4-Flash); Moonshot 1; MiniMax 1; Google 1 event / 4 sizes; OpenAI 1 event / 2 variants. No lab released more than 3; the frontier is a multi-lab race, not a single-lab streak.
- Quarterly cadence 2026: Q1 — quiet for weights (digestion of the 2025-12 MiMo release); Q2 — 5 events (Qwen3.6, Gemma 4, DeepSeek-V4, MiniMax M3, GLM-5.2); Q3 — 5 events (Kimi K3, Qwen3.8-Max, GLM-5.3-Flash, GLM-5.3, MiMo-V2.6). August alone had 3. The open-weight MoE frontier shipped roughly one flagship per month through 2026.

### The 2026 routing-design pattern book: who contributed what

- **DeepSeek:** fine-grained segmentation (DeepSeekMoE, 2024); auxiliary-loss-free dynamic-bias load balancing (V3 lineage, retained in V4 with a sequence-wise balance loss); MTP training heads; DeepEP dispatch. The lab whose 2024 innovations are the 2026 defaults.
- **Moonshot:** latent-space routing (Stable LatentMoE — top-k relocated into a 3,584-wide normalized latent space); quantile balancing with a single global all-reduce; SiTU-bounded activations for low-precision stability; MXFP4-native expert weights.
- **Alibaba/Qwen:** Gated DeltaNet linear-attention interleaves (3:1 in Qwen3.6, hybrid in Qwen3.8-Max); the small-tier MoE proof point (35B/3B); YaRN-extended 1M context on a hybrid stack.
- **MiniMax:** block-sparse MiniMax Sparse Attention with co-designed kernels; dense-first-3-layers pattern at 428B scale; always-on shared expert + native XML-namespace tool calling co-designed with routing.
- **Z.ai:** Manifold-Constrained Hyper-Connections for scaling efficiency at 320B; `glm5_next` hybrid sparse+linear stack; the no-shared-expert 288-expert design (with MiMo, one of two 2026 labs to ship zero shared experts).
- **Xiaomi/MiMo:** learnable attention-sink bias for 5:1 SWA/GA hybrids (~6× KV reduction [VENDOR]); DFlash-style multi-layer MTP drafters (5 layers, 7 tokens ahead); day-0 open serving culture.
- **Google:** the pure-attention MoE counterexample (no SSM, no linear attention); SWA(1024)/global hybrid; offload-friendly small-expert geometry enabling 8 GB consumer-GPU serving.
- **OpenAI:** the 2025 template — sparse MoE + native MXFP4 + Apache 2.0; the A-suffix naming convention the field adopted.

### Infrastructure and research actors

- **Communication libraries (co-designed with models):** DeepEP (DeepSeek lineage), MoonEP (Moonshot, 2026-07), llm-d wide-expert-parallelism documentation (DeepEP over NVSHMEM, GPU-initiated RDMA, sparse dispatch/combine over InfiniBand/RoCE while avoiding KV replication).
- **Research frontier (papers, not yet production):** **Stable LatentMoE** (Moonshot/Kimi K3) — latent-space routing with stable activation gradients across 896 experts; **MoUE — Mixture of Universal Experts** (arXiv 2603.04971, Mar 2026) — reusing a layer-agnostic universal expert pool across layers ("virtual width" from depth), up to +1.3% over matched MoEs and +4.2% when converting existing MoE checkpoints; **FLAME** (arXiv 2605.09355, May 2026) — adaptive MoE for continual multimodal learning with shared + task-specific experts; **MoE+Mamba-2 hybrids** (Nemotron 3.5 Lightning); **phase-aware MoE routing** (Yang et al. 2026) — routing at environment-step rather than token granularity for RL agents, enforcing temporal consistency across multi-step plans.
- **Fine-tuning frameworks** (Unsloth / Axolotl / torchtune / DeepSpeed; ESFT as the MoE-native PEFT method): consolidated in part 09b — referenced here only as actors that make the 2026 MoE fleet trainable, not as release facts.

## Timeline and context

### How MoE got here

