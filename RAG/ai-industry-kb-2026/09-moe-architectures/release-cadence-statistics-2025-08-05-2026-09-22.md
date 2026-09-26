---
id: ai-industry-kb-2026/09-moe-architectures/release-cadence-statistics-2025-08-05-2026-09-22
title: "Release cadence statistics (2025-08-05 → 2026-09-22)"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "China", "DeepSeek", "Google", "Huawei", "Hugging Face", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "Together AI", "Xiaomi", "Z.ai"]
dates: ["2025-08-05", "2025-12", "2025-12-16", "2026-04-16", "2026-04-24", "2026-06-01", "2026-06-13", "2026-07-16", "2026-07-24", "2026-07-27", "2026-08", "2026-08-03", "2026-08-08", "2026-08-13", "2026-08-22", "2026-08-25", "2026-08-26", "2026-08-28", "2026-09-22"]
keywords: ["agentic", "amd", "apache", "ascend", "attention", "blackwell", "compute", "consumer", "deepseek", "fp4", "fp8", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4616, 4668]
section: "9. MoE Architectures"
sha256: 6b9619fb6d0143d2bc12b2ff45614951f014bf23b3fb6fa31d1f41dd8d9d84b6
---

# Release cadence statistics (2025-08-05 → 2026-09-22)

- 2025-08-05: OpenAI gpt-oss-120b/20b — weights and announcement same day (arXiv 2508.10925 model card); the event that set the open-MoE template.
- 2025-12-16: Xiaomi MiMo-V2-Flash — weights and announcement same day; day-0 SGLang support shipped with inference code contributed upstream.
- 2026-04-16: Alibaba Qwen3.6-35B-A3B — weights on Hugging Face, HF model card + Together AI serving pages as the verification anchors; Apache 2.0 from day one.
- ~2026-04: Google Gemma 4 26B-A4B — HF model card + gemma4.dev as anchors; the family (E2B/E4B dense, 26B-A4B MoE, 31B dense) released as one licensing event (all Apache 2.0).
- 2026-04-24: DeepSeek V4 preview (same day as OpenAI GPT-5.5) — two variants announced together; V4-Pro reached GA 2026-08-13 (V4-Pro-0813 checkpoint); legacy aliases retired 2026-07-24.
- 2026-06-01: MiniMax M3 — announcement + weights; served on NVIDIA Blackwell and AMD Instinct via SGLang from launch.
- 2026-06-13/17: Z.ai GLM-5.2 — launch window (two dates in coverage); params undisclosed; base later disclosed through GLM-5.3 (753B).
- 2026-07-16: Moonshot announces Kimi K3 (2.8T) with technical report → weights 2026-07-27 (11 days later); MoonEP, FlashKDA, and AgentEnv open-sourced alongside.
- 2026-08-03: Alibaba Qwen3.8-Max GA → weights 2026-08-08 (5 days later), two HF repos (bf16 + FP8); FP8 out-downloaded bf16 by 2026-08-22 (21,400 vs 17,386) [COMMUNITY].
- 2026-08-25/26: Z.ai GLM-5.3-Flash — after ~1 week as the anonymous "Ox Alpha" atop OpenRouter coding charts, revealed 2026-08-26; weights `zai-org/GLM-5.3-Flash`; MIT.
- 2026-08-28: Z.ai GLM-5.3 flagship weights — post-training upgrade of the 5.2 base, custom license.
- ~2026-09: Xiaomi MiMo-V2.6-Flash/Pro — reported via OrcaRouter release analysis (single source); no Xiaomi first-party announcement directly verified [UNVERIFIED].
- Pattern note: by 2026, GA and weights are decoupled by days, and stealth leaderboard launches ("Ox Alpha") precede formal reveals — release tracking must watch anonymous leaderboards, not only press releases.

### Release cadence statistics (2025-08-05 → 2026-09-22)

- 14 tracked release events in ~13.5 months (gpt-oss 120b/20b, MiMo-V2-Flash, Qwen3.6-35B-A3B, Gemma 4 26B-A4B, DeepSeek-V4 Pro/Flash, MiniMax M3, GLM-5.2, Kimi K3, Qwen3.8-Max, GLM-5.3-Flash, GLM-5.3, MiMo-V2.6-Flash/Pro) — roughly one frontier open MoE per month, accelerating to three in August 2026 alone.
- Total-parameter ceiling: 117B (2025-08) → 309B (2025-12) → 428B (2026-06-01) → 1.6T (2026-04-24, DeepSeek-V4-Pro — note the ceiling jumped before M3 despite the later date) → 2.8T (2026-07-27, Kimi K3). The ceiling roughly doubled every ~4 months.
- Active-parameter ceiling: 5.1B → 15B (2025-12) → 23B (2026-06) → 49B (2026-04-24) → 104B (2026-07-27). Active compute grew ~20× in 11.5 months while sparsity ratios stayed in the 12:1–33:1 band — the field scales capacity and per-token compute together, not sparsity.
- Context ceiling: 128K (2025-08) → 256K (2025-12 / ~2026-04) → 1M native (2026-04-24 onward); 1M is universal across every 2026 flagship.
- Expert-pool ceiling: 32 (2025-08) → 128 (2025-08, the 120b) → 256 (2026-04) → 512 (2026-08) → 896 (2026-07-27, out of chronological order — Kimi K3 leapfrogged Qwen3.8-Max's 512 before its release).
- Lab concentration: of the 12 distinct 2026 releases, 10 are from Chinese labs (DeepSeek, Alibaba ×2, Moonshot, Z.ai ×3, MiniMax, Xiaomi ×2), 1 from Google, 0 from Western labs beyond Google.

### Architectural firsts per release (2025-08 → 2026-09)

- gpt-oss-120b/20b (2025-08-05): first open-weight MoE from OpenAI; native MXFP4; the sparse-MoE + sub-8-bit + permissive-license template; the A-suffix naming convention.
- MiMo-V2-Flash (2025-12-16): first 300B-class open MoE; MTP self-speculative decoding at 150 tok/s; learnable attention-sink bias with 5:1 SWA/GA.
- Qwen3.6-35B-A3B (2026-04-16): first hybrid Gated DeltaNet + gated attention open MoE; YaRN-extended 1M context on a hybrid linear/standard stack; the 3B-active coding MoE.
- Gemma 4 26B-A4B (~2026-04): the pure-attention MoE (no SSM, no linear attention); the expert-offloading case study (8 GB consumer GPU).
- DeepSeek-V4 (2026-04-24): largest open MoE of H1 2026; first flagship launched on Huawei Ascend; reconstructed innovations: sqrt(softplus) routing affinity, deterministic hash routing for the first three blocks, FP4 expert weights.
- MiniMax M3 (2026-06-01): first open release combining reasoning + agentic coding + native multimodality (text/image/video from step 0); MSA block-sparse attention.
- Kimi K3 (2026-07-27 weights): largest open-weight release ever (1.56 TB); Stable LatentMoE latent-space routing; Quantile Balancing; SiTU activation; MXFP4-native 2.8T; MoonEP + FlashKDA co-shipped.
- Qwen3.8-Max (2026-08-08 weights): first Max-class Qwen with open weights; 512-expert pool; 95B active — the highest per-token compute of any Alibaba open release.
- GLM-5.3-Flash (2026-08-26 reveal): first natively multimodal GLM-5; first "Ox Alpha"-style stealth launch on OpenRouter; `glm5_next` hybrid sparse+linear stack with mHC; served entirely on domestic Chinese chips during stealth.
- GLM-5.3 (2026-08-28): the post-training-only flagship upgrade — same 753B base as 5.2, gains from post-training alone; the 2026 case study that post-training, not architecture, decides capability deltas.
- MiMo-V2.6-Flash/Pro (~2026-09): first native omnimodal open MoE (text/image/video/audio in one model); 5-layer DFlash-style MTP drafter.

### Model deep dives

#### gpt-oss-120b / 20b — the 2025 open-weight MoE baseline (released 2025-08-05)

- **Specs (verified via arXiv model card 2508.10925):** gpt-oss-120b: 116.83B total (114.71B MLP + 0.96B attention + 1.16B embed/unembed), 5.13B active per token; 36 layers; 128 experts with 4 active per token (no shared expert); 128K context; 60.8 GiB checkpoint. gpt-oss-20b: 20.91B total, 3.61B active; 24 layers; 32 experts, 4 active; 12.8 GiB checkpoint.
- MXFP4 native quantization; Apache 2.0. First OpenAI open weights since GPT-2 (2019).
- Historical role: established the template every 2026 release follows — sparse MoE + native sub-8-bit quantization + permissive licensing — and the "active parameter" A-suffix naming convention now standard across Qwen/Gemma.

#### Xiaomi MiMo-V2-Flash — first 300B-class open MoE (released 2025-12-16)

- **Specs:** 309B total / 15B activated per token; 48 layers (39 sliding-window + 9 global attention); hidden size 4096; 256 routed experts with 8 activated, **no shared experts**; 128-token sliding window; first block is global attention with dense FFN, interleaved thereafter; learnable attention-sink bias → ~6× KV-cache reduction vs full attention [VENDOR].
- Trained on 27T tokens in FP8 mixed precision; 150 tok/s inference via MTP self-speculative decoding (accepted length 2.8–3.6, 2.0–2.6× effective speedup) [VENDOR]; day-0 SGLang support with inference code contributed upstream.
- 256K context; MIT license; the December 2025 proof that 300B-class MoE could be production-served at 150 tok/s.

#### Alibaba Qwen3.6-35B-A3B — 35B/3B hybrid-attention coding MoE (released 2026-04-16)

