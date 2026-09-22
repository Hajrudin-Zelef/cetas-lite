---
id: ai-industry-kb-2026/09-moe-architectures/main-actors-continued
title: "Main actors (continued)"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "China", "DeepSeek", "Google", "Huawei", "Hugging Face", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "TensorRT-LLM", "Unsloth", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2026-02-09", "2026-02-10", "2026-03-18", "2026-04", "2026-04-16", "2026-04-24", "2026-05", "2026-05-06", "2026-05-13", "2026-05-22", "2026-06", "2026-06-01", "2026-06-13", "2026-07", "2026-07-08", "2026-07-22", "2026-07-24", "2026-07-27", "2026-08-03", "2026-08-08", "2026-08-16", "2026-08-19", "2026-08-25", "2026-08-26", "2026-08-28", "2026-09", "2026-09-01", "2026-09-02", "2026-09-07", "2026-09-09", "2026-09-15", "2026-09-17", "2026-09-22", "2026-10-28"]
keywords: ["agentic", "agents", "amd", "apache", "ascend", "attention", "attribution", "awq", "benchmark", "benchmarks", "blackwell", "claude"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5283, 5386]
section: "9. MoE Architectures"
sha256: 4243ed4d63c9c3f53e3b89e53b41d16e558e26439ef0e90b7addd36439ab33d4
---

# Main actors (continued)

## Main actors (continued)

- **Labs and their MoE economics posture:** DeepSeek (Hangzhou) — the open training-systems reference (V3 recipe, DeepEP, DualPipe, aux-loss-free balancing, V4's CSA/HCA + FP4 experts); Alibaba/Qwen (Hangzhou) — Gated DeltaNet linear-attention hybrids and the largest serving-motivated FP8 adoption; Moonshot AI (Beijing) — bandwidth-aware fine-grained MoE (Stable LatentMoE, Quantile Balancing, MoonEP, FlashKDA) and the most expensive Chinese-lab API ($3.00/M input for K3); Z.ai (Beijing) — the Flash-tier price disruptor ($0.15/M input, MIT license, stealth "Ox Alpha" launch on domestic chips); MiniMax (Shanghai) — MSA sparse attention and NVFP4 serving; Xiaomi/MiMo (Beijing) — SWA/GA hybrids and MTP self-speculative decoding at 150 tok/s. Google (Gemma 4) is the only non-Chinese lab with a 2026 open MoE release; OpenAI's gpt-oss (Aug 2025, Apache 2.0, MXFP4 native) is the Western baseline; NVIDIA's Nemotron 3.5 Lightning (30B/3B MoE+Mamba-2) is the Western efficiency reference point. Co-designed serving stacks are the moat: DeepEP, MoonEP + FlashKDA, MSA kernels, MLA/DeepEP lineage — model and serving system designed together, day-0 SGLang/vLLM support standard.
- **Serving-stack vendors (2026):** vLLM (EP flags, EPLB on vLLM-Ascend, v0.15.0 MoE kernels Marlin/NVFP4-CUTLASS/FP8/INT8 non-gated); SGLang (DeepEP integration, upstream-contributed model support, reasoning/tool-call parser auto-detection); NVIDIA Dynamo (vLLM/SGLang/TensorRT-LLM runtimes, KV-aware routing, disaggregated prefill/decode); Moonshot (MoonEP open-sourced 2026-07, FlashKDA kernels); Huawei Ascend (V4 launch platform — first major model debuting on non-NVIDIA chips; vLLM-Ascend EPLB + W8A8/W4A8/MXFP4/MXFP8 on Ascend 950); AMD (ROCm EP topology guidance; K3 MI355X serving benchmarks).
- **Training-systems actors:** Unsloth (Feb 2026 MoE pipeline; May 2026 NVIDIA collab; Studio betas Sept 2026; 73K+ GitHub stars by Aug 30, ~400–570/day — developer-mindshare signal, not enterprise-dominance signal); Axolotl (ScatterMoE/SonicMoE LoRA, `quantize_moe_experts`, DeepEP-based EP, NVFP4 MoE LoRA — multi-GPU research workflows); torchtune (May 2026 paper: DTensor FSDP2 + custom expert-parallel plan + loss parallel + Ring-Attention context parallel; independent comparison finding Unsloth the memory leader but torchtune throughput-leading at 3 of 4 sizes); DeepSpeed (max-scale path: ZeRO family, AutoEP Feb 2026, MoE all-reduce comm-opt); Megatron-style systems (trillion-param custom scale); Liger Kernel (open Triton kernels complementary to Unsloth, TRL/FSDP2/DeepSpeed integrations). Research actors: MoE Parallel Folding authors (arXiv 2504.14960v2 — the fine-granularity efficiency correction); ESFT authors (DeepSeek AI + Northwestern, arXiv 2407.01906 — expert-specialized fine-tuning, finer experts MORE advantageous for ESFT); MoUE (arXiv 2603.04971, universal experts); FLAME (arXiv 2605.09355, continual multimodal MoE); llm-d project (wide-EP documentation); Chew Loong Nian (July 2026 independent analysis: 87.3% of sequence-scaled memory in Llama 3.1 8B LoRA is the cross-entropy loss head — independently validating cut/fused cross-entropy as the true framework differentiator).

- **Research-frontier actors (dated):** MoUE — Mixture of Universal Experts (arXiv 2603.04971, Mar 2026) — layer-agnostic universal expert pool shared across layers ("virtual width" from depth), +1.3% over matched MoEs, +4.2% when converting existing MoE checkpoints; FLAME (arXiv 2605.09355, May 2026) — adaptive MoE for continual multimodal learning with shared + task-specific experts; LoRA-FA (arXiv 2511.04021, ICLR 2026 submission) and DoRA (arXiv 2402.09353) — the adapter-quality axis; PEFT-Arena (May 2026) — stability–plasticity comparison across LoRA, AdaLoRA, DoRA, VeRA, PiSSA, MiLoRA, OFT, IA3; RoRA and Dual LoRA (ICLR 2026) reportedly beating DoRA; phase-aware MoE routing for RL agents (2026) — routing at environment-step granularity for temporal consistency; the MoE Parallel Folding authors (arXiv 2504.14960v2) and the two-phase EP dispatch analysis (arXiv 2503.04398) — the systems literature that corrected the fine-granularity narrative.

- **Community and independent analysts (dated):** Pelin Balci (2026-02-09 independent Unsloth-vs-standard comparison with loss — the independent practitioner check that included loss parity); Chew Loong Nian (July 2026 byte-level loss-head analysis — the cross-entropy memory attribution); the torchtune paper authors (May 2026, arXiv 2605.21442v1 — the only independent 3-framework torchtune/Axolotl/Unsloth benchmark); MarkTechPost (2026-07-22 4-framework comparison — the B200 memory figures); the karpathy-wiki synthesis (2026-09-01 arXiv AI search — among the archived secondary sources); OliverSundaram/MoE-Study (2026 controlled from-scratch study — the small-scale dense-beats-MoE result); malcomzww/moe-vs-dense-serving-profile (the 5.9× residency measurement); solrex SGLang cookbook and semianalysisai inference notes (MiniMax M3 serving details); alexchen31337 gemma4-moe-offload (the Gemma 4 offload blueprint); the unsloth-cli maintainers (agentculture, 0.7.0/0.7.1, 2026-09-15 — container-backed export + quantization-loss eval). GitHub-AI-radar daily reports (Aug 2026) tracked Unsloth's star trajectory — the adoption signal.

## Timeline and context (continued)

- **2023:** Mixtral 8x7B (46.7B/12.9B) proves open MoE economics — dense-70B-class capability at ~6× inference speed; Mixtral 8x22B (141B/39B) follows.
- **2024:** DeepSeek-V2 introduces MLA (KV cache ~5% of LLaMA-3-70B) and DeepSeekMoE fine-grained segmentation; DeepSeek-V3 (Dec 2024, 671B/37B) becomes the open reference for aux-loss-free balancing, MTP heads, FP8 multi-plane training.
- **2025:** gpt-oss (Aug 2025) normalizes open MoE + native sub-8-bit quantization + the A-suffix "active parameter" naming; MiMo-V2-Flash (Dec 2025) pushes 300B-class MoE to 150 tok/s production serving with MTP self-speculative decoding.
- **2026-02:** Unsloth MoE pipeline launch ("12x faster" [VENDOR]); DeepSpeed AutoEP blog (2026-02-10); DeepSeek-V4 previewed 2026-04-24 (same day as GPT-5.5) with 33T/32T-token pretraining; Qwen3.6-35B-A3B (2026-04-16) establishes the 3B-active agentic-coding tier (73.4% SWE-bench Verified); Gemma 4 (~2026-04) ships the expert-offloading consumer-GPU blueprint.
- **2026-05:** Unsloth×NVIDIA collab blog (2026-05-06) publishes the auditable component numbers; DeepSeek makes V4-Pro's 75% discount permanent (2026-05-22); torchtune paper (arXiv 2605.21442v1) independently benchmarks Unsloth vs Axolotl vs torchtune; PyTorch 2.12 (2026-05-13).
- **2026-06:** pre-registered matched-compute ablation (MoE 4.0× on perplexity, 5.72 vs 22.66); Axolotl adds DeepEP-based EP for distributed MoE training; MiniMax M3 (2026-06-01) launches with MSA; GLM-5.2 (2026-06-13/17).
- **2026-07:** Kimi K3 weights (2026-07-27) — 2.8T/104B, the bandwidth-aware fine-grained blueprint (Stable LatentMoE, Quantile Balancing, MXFP4, SiTU, per-head Muon); MoonEP + FlashKDA open-sourced; Unsloth v0.1.481-beta adds NVFP4/FP8/imatrix-GGUF export; Axolotl NVFP4 MoE LoRA; MarkTechPost 4-framework comparison (2026-07-22); `deepseek-chat`/`deepseek-reasoner` aliases retired 2026-07-24 (→ V4-Flash).
- **2026-08:** Qwen3.8-Max GA (2026-08-03) + open weights (2026-08-08, non-Apache license) — 2.4T/95B, ~1:25 sparsity; GLM-5.3-Flash stealth-launch (2026-08-25/26) as "Ox Alpha" then revealed — $0.15/M input, MIT; GLM-5.3 flagship weights (2026-08-28) — 753B base, custom license, $1.40/M; Unsloth Dynamic v3.0 launch (2026-08-19); peak/off-peak V4 tiers from 2026-08-16; Qwen3.8-Max FP8 repo becomes the most-downloaded variant (serving-motivated fetchers).
- **2026-09:** GLM-5.3-Flash promo ends 2026-09-09 ($0.075 → $0.15 input); Unsloth Studio betas v0.1.805–808 (MTP-by-default for Qwen3.8-Flash/GLM-5.3-Flash, MoE expert offload to system RAM, Vulkan-by-default on AMD, MLX DoRA/DPO) in the first half of the month; MoE kernel correctness fix (2026-09-09, unsloth-zoo-staging PR #887); MiniMax M3 80.5% SWE-bench Verified (AA, 2026-09-07); PyTorch 2.14 (2026-09-02) with the silent clamp-gradient change; MiMo-V2.6-Flash/Pro (~2026-09) — 309B/15B Flash, 1.02T/42B Pro, MIT, omnimodal; unsloth-cli 0.7.0/0.7.1 (2026-09-15) operationalizes container-backed multi-format export + quantization-loss eval. No Unsloth releases after the 2026-09-17 official changelog entry as of 2026-09-22.

- **Licensing timeline (dated):** Gemma 4 (~2026-04) shifts the Gemma line to Apache 2.0 from earlier Gemma-specific terms; GLM-5.3-Flash (2026-08-26) is MIT — the cheapest Flash-tier serving model ($0.15/M input) is also the most permissively licensed; Qwen3.8-Max (2026-08-08) ships a non-Apache house license; GLM-5.3 flagship (2026-08-28) ships a custom `glm-5.3` license; MiniMax M3 (2026-06-01) ships a MiniMax Community License; DeepSeek-V4's license (MIT in later sources vs Apache 2.0 in April 2026 coverage) is unresolved — do not cite without checking the first-party model card. Trend: Apache 2.0/MIT cover the majority of 2026 open MoE releases; custom licenses cluster at the capability-leading models.
- **Domestic-chip serving milestones (dated):** DeepSeek-V4 previewed 2026-04-24 launching on Huawei Ascend chips first (no NVIDIA required); GLM-5.3-Flash spent its first week topping OpenRouter coding charts anonymously as "Ox Alpha" served entirely on domestically produced Chinese AI chips (revealed 2026-08-26); Kimi K3 runs on AMD MI355X (serving benchmarks exist); Unsloth Studio v0.1.807/808-beta (~early Sept 2026) made AMD Vulkan-by-default (+20% prefill, +23% prompt processing, +8% generation on Strix Halo [VENDOR]) and RDNA1/2 supported (Sept-17 changelog). Serving economics are no longer NVIDIA-exclusive.

### PyTorch and platform cadence (dated)

- **PyTorch 2.11 (2026-03-18) → 2.12 (2026-05-13) → 2.13 (2026-07-08) → 2.14 (2026-09-02)** — brisk 2026 cadence; 2.14 silently changed clamp/min/max boundary subgradients (1→0), a genuine training-reproducibility hazard [DIRECTIONAL — independent dev-blog source]. 2.15 scheduled for 2026-10-28.
- **Unsloth backend versioning:** year.month scheme — 2026.6.9 (June) → 2026.7.2 (July) → 2026.8.3 (August) → 2026.9.1/2026.9.2 (September, PR of 2026-09-02, stamped on v0.1.806-beta).
- **Hardware-side dates that move serving economics:** MoonEP open-sourced 2026-07; DeepSpeed AutoEP blog 2026-02-10; vLLM v0.15.0 MoE kernels (2026); Unsloth Studio AMD RDNA1/2 + Vulkan-by-default (Sept 2026 betas); Unsloth official changelog's latest entry 2026-09-17 ("Docker + MultiUser + AMD Support"); no Unsloth releases after 2026-09-17 as of 2026-09-22.
- **Framework refresh cadence:** Axolotl's MoE cadence (Feb ScatterMoE → Mar quantize_moe_experts → Apr Async GRPO/SonicMoE → Jun DeepEP EP → Jul NVFP4 MoE LoRA); torchtune May 2026 paper; DeepSpeed MoE tutorial refresh ~Sept 2026.

## Implications (continued)

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

- NVIDIA dense-vs-MoE technical blog: https://developer.nvidia.com/blog/dense-vs-moe-models-active-parameters-throughput-and-when-to-choose-each/ [VENDOR]
- MoE KV-cache economics research notes: https://github.com/eduardogbg/mnemo/blob/HEAD/docs/research/dag-kv-cache-economics.md
- MoE inference optimization (Spheron): https://www.spheron.network/blog/moe-inference-optimization-gpu-cloud/
- vLLM-Ascend expert parallelism load balancer: https://github.com/vllm-hust/vllm-ascend-hust/blob/HEAD/docs/source/user_guide/feature_guide/expert_parallelism_load_balancer.md
- vLLM Mixtral MoE optimization case: https://github.com/ineshreddy249/vllm-mixtral-moe-optimization/blob/HEAD/README.md
- ROCm vLLM inference optimization (EP topology guidance): https://github.com/yusrildestroy1/rocm/blob/HEAD/docs/how-to/rocm-for-ai/inference-optimization/vllm-optimization.rst
- vLLM distributed inference cookbook: https://github.com/wahajsayyed/ai-system-engineering/blob/HEAD/vLLM-cookbook/chapter-07-distributed-inference.md
- MoE tutorial (ESFT, aux-loss discussion): https://github.com/wanshuiyin/auto-claude-code-research-in-sleep/blob/HEAD/docs/tutorials/moe_tutorial_en.md
- MoE at Blackwell lecture notes: https://github.com/kingarthurv1/ai-hardware-engineer-roadmap/blob/HEAD/Phase%205%20-%20Advanced%20Topics%20and%20Specialization/7.%20ML%20Systems%20Engineering/AI%20Inference%20Engineer%202026/Part%203%20-%20MoE%20at%20Blackwell/Lecture-01.md [COMMUNITY]
- MoE vs dense serving profile (OLMoE): https://github.com/malcomzww/moe-vs-dense-serving-profile [COMMUNITY]
- Controlled MoE vs dense study: https://github.com/oliversundaram/moe-study [COMMUNITY]
- GLM-5.3-Flash release (MarkTechPost, 2026-08-26): https://www.marktechpost.com/2026/08/26/z-ai-releases-glm-5-3-flash-a-320b-a18b-natively-multimodal-moe-with-a-1m-token-context/
- GLM-5.3-Flash model page (Atomic): https://atomic.chat/models/glm-5-3-flash
- GLM-5.3-Flash analysis (LumaDock): https://lumadock.com/blog/glm-5-3-flash
- Kimi K3 breakdown (Medium, secondary): https://medium.com/@jamilxt/kimi-k3-is-now-open-weight-the-full-breakdown-of-the-2-8t-frontier-model-65e7033aa917 [COMMUNITY]
- Kimi K3 concept wiki: https://github.com/kzinmr/ai-topics/blob/HEAD/wiki/concepts/kimi-k3.md [COMMUNITY]
- Kimi K3 technical report (primary, Moonshot): https://raw.githubusercontent.com/MoonshotAI/Kimi-K3/master/k3_tech_report.pdf
- Kimi K3 overview (ai-stack): https://ai-stack.ai/en/kimi-k3-moonshot-ai-open-source
- MoE Parallel Folding (arXiv 2504.14960v2): https://arxiv.org/pdf/2504.14960v2.pdf
- Expert-parallel dispatch two-phase analysis (arXiv 2503.04398): https://arxiv.org/pdf/2503.04398
- llm-d wide expert parallelism docs: https://github.com/llm-d/llm-d.github.io/blob/HEAD/versioned_docs/version-0.7/well-lit-paths/wide-expert-parallelism.md
- torchtune paper, May 2026 (independent framework comparison): https://arxiv.org/pdf/2605.21442v1.pdf
- MarkTechPost 4-framework comparison, 2026-07-22: https://www.marktechpost.com/2026/07/22/unsloth-vs-axolotl-vs-trl-vs-llama-factory-a-fine-tuning-framework-comparison-on-speed-vram-and-multi-gpu/
- Hugging Face Unsloth–TRL benchmark blog: https://huggingface.co/blog/unsloth-trl
- DeepSpeed communication optimization (MoE all-reduce penalties): https://github.com/deepspeedai/deepspeed/blob/HEAD/blogs/comm-opt/README.md
- DeepSpeed AutoEP blog, 2026-02-10: https://deepspeed.ai/blog/2026/02/10/AutoEP/
- Unsloth Studio v0.1.501-beta release-body fixture (MoE expert offload to system RAM): https://github.com/unslothai/unsloth/blob/HEAD/tests/studio/fixtures/release_bodies/v0.1.501-beta.md
- Unsloth Qwen3.8 fine-tuning guide (Sept 2026; VRAM tiers, Flash Linear Attention kernels): https://unsloth.ai/docs/models/qwen3.8/train [VENDOR]
- Unsloth official changelog (latest entry 2026-09-17): https://unsloth.ai/docs/new/changelog [VENDOR]

