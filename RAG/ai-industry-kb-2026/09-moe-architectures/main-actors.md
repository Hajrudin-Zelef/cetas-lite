---
id: ai-industry-kb-2026/09-moe-architectures/main-actors
title: "Main actors"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Huawei", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "Unsloth", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2025-08-05", "2025-12-16", "2026-04-16", "2026-04-24", "2026-05", "2026-06-01", "2026-06-13", "2026-07-16", "2026-07-27", "2026-08-03", "2026-08-08", "2026-08-25", "2026-08-26", "2026-08-28", "2026-09-22"]
keywords: ["agentic", "agents", "amd", "apache", "ascend", "attention", "benchmark", "benchmarks", "compute", "consumer", "cost", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4846, 4938]
section: "9. MoE Architectures"
sha256: 871b192d277c4ca5d71bc3cf340ada9bcc1bcc34ac4061fc8e327fe512fa76bd
---

# Main actors

## Main actors

### The six Chinese labs that produced every 2026 frontier open MoE

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

- **2017–2021:** Sparsely-Gated MoE (Shazeer et al.) showed capacity could scale far beyond dense with fixed compute; Switch Transformer (2021) simplified to top-1 routing and demonstrated 1.6T-parameter training.
- **2023:** Mixtral 8x7B (46.7B total / 12.9B active, 32K context, Apache 2.0) proved open-weight MoE could rival dense LLaMA 2 70B at 6× the inference speed; Mixtral 8x22B (141B/39B) followed.
- **2024:** DeepSeek-V2 introduced MLA and DeepSeekMoE fine-grained segmentation (160 experts at 2.5B-scale demos); DeepSeek-V3 (Dec 2024, 671B/37B) became the open reference for auxiliary-loss-free load balancing, MTP heads, and multi-plane FP8 training.
- **2025:** gpt-oss (Aug) normalized open-weight MoE + native sub-8-bit quantization + permissive licensing; MiMo-V2-Flash (Dec) pushed 300B-class MoE into production serving at 150 tok/s with speculative self-decoding.
- **2026:** the field consolidated on the total/active naming (A-suffix), 1M context as standard, hybrid attention stacks, fine-grained pools (128–896), shared experts, and day-0 open serving infrastructure.

### Month-by-month release narrative (2026)

- **April:** the densest release month. Qwen3.6-35B-A3B (2026-04-16) proved a 3B-active MoE could do agentic coding; Gemma 4 26B-A4B (~2026-04) gave the West its only 2026 open MoE; DeepSeek V4 preview (2026-04-24, same day as GPT-5.5) set the H1 ceiling at 1.6T/49B and launched on Huawei Ascend first.
- **May:** a quiet month for weights — the field digested April's releases; no major open-weight MoE shipped.
- **June:** MiniMax M3 (2026-06-01) combined reasoning + agentic coding + native multimodality in one open release; GLM-5.2 (2026-06-13/17) established Z.ai's coding-plan tier (base later disclosed via GLM-5.3).
- **July:** Kimi K3 announced 2026-07-16, weights 2026-07-27 — the 2.8T release that reset the open-weight ceiling; MoonEP + FlashKDA + AgentEnv open-sourced alongside, making July the month the model and its serving system shipped as one artifact.
- **August:** Qwen3.8-Max GA 2026-08-03, weights 2026-08-08 (first Max-class Qwen with open weights); GLM-5.3-Flash revealed 2026-08-26 after its "Ox Alpha" stealth week (served entirely on domestic Chinese chips); GLM-5.3 flagship weights 2026-08-28.
- **September:** MiMo-V2.6-Flash/Pro (~2026-09, single-sourced) pushed the MoE frontier to native omnimodality including audio; knowledge cutoff 2026-09-22.

### The editorial correction, stated once for the record

- The draft synthesis "2026 = the year MoE becomes the default architecture on all serious models" is too absolute and must not be carried forward. What the evidence supports: MoE became the default architecture for **very large open-weight models** — every frontier open-weight release since mid-2025 is sparse MoE, and the 300B+ class is now exclusively sparse in the open.
- What the evidence does not support: dense counterexamples survive at the small/workstation tier (Gemma 4 31B dense, Qwen3.8-27B dense, gpt-oss-20b at the edge) and won the one controlled matched-active-params study at ~150M scale (dense beat MoE on almost every benchmark and ran 3× faster — routing overhead dominates at small scale [COMMUNITY]); closed labs' flagship architectures remain undisclosed, so no statement about "all serious models" is checkable.
- The defensible formulation: **2026 is the year sparse MoE won the open frontier at scale; dense remains competitive below ~35B-class active compute and at the edge.**

### What 2026 settled — and what it did not

- **Settled:** the total/active naming convention (A-suffix); fine-grained expert pools (128–896) as the 300B+ default; shared experts as near-universal (1–2, with three documented outliers); auxiliary-loss-free load balancing as the production standard; 1M context as the flagship table stake; hybrid attention stacks as the way to make it servable; MTP training heads doubling as speculative drafters; day-0 SGLang/vLLM support; permissive licenses (Apache 2.0/MIT) covering the majority of open-weight MoE releases; the KV cache as attention's cost, not MoE's.
- **Not settled:** DeepSeek-V4's license (MIT vs Apache 2.0) and its reconstructed architecture numbers ([PARTIALLY VERIFIED]); Kimi K3's active params (104B vs ~50B); GLM-5.3-Flash's context (1M config vs 300K eval mention); whether MoUE's virtual width holds at 2T+ params and 1M context; whether any MoE-aware KV scheme beats attention-side hybrids; dynamic per-token top-k in production; 100B+ MoE on consumer hardware (NVMe-bandwidth routing prediction still open); whether latent-space routing spreads beyond Moonshot; whether dense strikes back below ~35B-class; what the closed labs are actually running.
- **The one structural open question for 2027:** whether the open-weight MoE frontier stays a Chinese-lab monopoly or a Western 100B+ open MoE appears — the entire 2026 leadership map hangs on this.

### Open research questions for the next wave

- **Dynamic expert count:** can top-k adapt per token (easy tokens → fewer experts) without destabilizing training? Early 2026 work on confidence-gated routing suggests 20–40% compute savings [DIRECTIONAL].
- **Universal experts at scale:** MoUE's virtual-width idea is proven at research scale; does it hold at 2T+ params and 1M context?
- **KV cache for 1M-context MoE:** all 2026 KV savings are attention-side; is there a principled MoE-aware KV scheme (e.g., expert-conditioned KV sharing) that beats MLA/GQA/SWA hybrids?
- **Routing for agents:** phase-aware and step-granularity routing is promising for multi-step tool use but unproven in production serving stacks.
- **MoE on the edge:** expert offloading (Gemma 4) and MXFP4 (gpt-oss) show the path; whether 100B+ MoE becomes viable on consumer hardware depends on NVMe-bandwidth routing prediction, still an open problem.
- **Latent routing generalization:** Kimi K3's Stable LatentMoE makes 16-of-896 routing trainable; does latent-space routing spread to other labs' designs in 2027?
- **Fine-tuning-aware routing design:** ESFT (DeepSeek AI + Northwestern, arXiv 2407.01906) shows routing distributions concentrate per task and that finer-grained experts favor expert-specialized fine-tuning — do future MoE designs optimize the routing geometry for post-training specialization, not just pretraining perplexity?

### The 2026 open-weight MoE leadership map

- **China produced every frontier open-weight MoE release of 2026:** DeepSeek (Hangzhou), Alibaba/Qwen (Hangzhou), Moonshot AI (Beijing), Z.ai (Beijing), MiniMax (Shanghai), Xiaomi/MiMo (Beijing). Google (Gemma 4) is the only non-Chinese lab with a 2026 open MoE release; OpenAI's gpt-oss (2025) is the Western baseline.
- **Co-designed serving stacks are the moat:** DeepEP (DeepSeek), MoonEP + FlashKDA (Moonshot), MSA kernels (MiniMax) — each lab ships the model with the communication library and kernels needed to serve it. Day-0 SGLang/vLLM support is now standard practice (MiMo-V2-Flash contributed inference code upstream).
- **Domestic-chip milestones:** DeepSeek-V4 launched on Huawei Ascend chips first (no NVIDIA required); GLM-5.3-Flash spent its stealth week ("Ox Alpha") topping OpenRouter charts served entirely on domestically produced Chinese AI chips; Kimi K3 runs on AMD MI355X.
- **Licensing shape:** the most permissive standard licenses (Apache 2.0, MIT) now cover the majority of open-weight MoE releases; custom licenses cluster at the largest/capability-leading models (MiniMax Community License, GLM-5.3 flagship glm-5.3 license, Qwen3.8-Max non-Apache open license) — the full license landscape is consolidated in part 09b.
- **For RAG/knowledge-base relevance:** the leadership map matters because the models RAG pipelines will run on through 2027 are overwhelmingly Chinese-lab open MoE with co-designed serving stacks — evaluate serving maturity (EP libraries, day-0 engine support, kernel availability) alongside weights when adopting.

