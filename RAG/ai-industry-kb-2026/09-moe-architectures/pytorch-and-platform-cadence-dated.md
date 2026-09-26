---
id: ai-industry-kb-2026/09-moe-architectures/pytorch-and-platform-cadence-dated
title: "PyTorch and platform cadence (dated)"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "China", "DeepSeek", "Huawei", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-02-10", "2026-03-18", "2026-04", "2026-04-16", "2026-04-24", "2026-05", "2026-05-06", "2026-05-13", "2026-05-22", "2026-06-01", "2026-06-13", "2026-07-08", "2026-07-22", "2026-07-24", "2026-07-27", "2026-08-03", "2026-08-08", "2026-08-16", "2026-08-19", "2026-08-25", "2026-08-26", "2026-08-28", "2026-09-02", "2026-09-07", "2026-09-09", "2026-09-15", "2026-09-17", "2026-09-22", "2026-10-28"]
keywords: ["agentic", "amd", "apache", "ascend", "benchmarks", "compute", "consumer", "deepseek", "dpo", "fp8", "gguf", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5295, 5316]
section: "9. MoE Architectures"
sha256: 7cca44b265b219355e05bb65459502f1117ba93b4d1a478b54ce266fdea7a691
---

# PyTorch and platform cadence (dated)

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

