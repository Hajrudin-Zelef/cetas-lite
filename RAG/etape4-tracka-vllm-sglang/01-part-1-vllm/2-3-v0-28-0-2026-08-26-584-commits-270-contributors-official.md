---
id: etape4-tracka-vllm-sglang/01-part-1-vllm/2-3-v0-28-0-2026-08-26-584-commits-270-contributors-official
title: "2.3 v0.28.0 — 2026-08-26 (584 commits, 270 contributors) [official][secondary]"
domain: part-1-vllm
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "DeepSeek", "Meta", "Mistral", "Moonshot", "Nvidia", "vLLM"]
dates: ["2026-02-25", "2026-07-27", "2026-08-26", "2026-09-04"]
keywords: ["amd", "attention", "decode", "deepseek", "diffusion", "fp8", "gpu", "kimi", "kv cache", "lora", "moe", "multimodal"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [188, 239]
section: "PART 1 — vLLM"
sha256: 2952120d0a0f4118cd60d5db6cf812a5d060aa2c8671fa6f635009562f67ca34
---

# 2.3 v0.28.0 — 2026-08-26 (584 commits, 270 contributors) [official][secondary]

### 2.3 v0.28.0 — 2026-08-26 (584 commits, 270 contributors) [official][secondary]

- **Kimi-K3 performance push:** Decode Context Parallel (DCP), fused FlashKDA decode/prefill kernels,
  SiTU activation for MegaMoE, GEMM-RS for sequence parallelism, combined all-gathers (1.5–3× kernel-level
  speedup reported), adaptive speculative token budget (~60% better DSpark TTFT reported), optional
  shared-expert sharding (~17 GiB saved per GPU reported); runs on ROCm with the V2 model runner.
- **DeepSeek V4:** sparse MLA works end-to-end for plain decode, MTP, and DSpark spec decode; AMD Quark
  NVFP4 support; reasoning-effort prompts; ROCm enablement on gfx11 and gfx950.
- **Tiered KV cache offloading:** KV cache offload to disk (object-store secondary tier, DP-replica-aware
  tiering, EC connectors incl. CPU offload); `max_num_batched_tokens` default raised 8192 → 16384.
- **New models:** Muse Glimmer, Ling 3.0 Flash (BF16/FP8/MXFP4), Dots3 NOTE.
- ⚠️ **Breaking changes:** bitsandbytes support migrated to an **out-of-tree plugin**; deprecated
  `calculate_kv_scales` runtime KV-scale calculation, `override_attention_dtype`, MoE legacy code removed;
  KV-offload tiering metrics renamed `kv_offload_tiering_block_{queries,hits}` → `..._chunk_...`;
  Transformers 5.15.0 upgrade (breaking environment change); `reasoning_content` output removal documented
  as a breaking client change.
- Spec decode: DFlash2 with local convolution + candidate selector, DSpark confidence-scheduled
  verification, async scheduling auto-enabled for draft models.
- Model Runner V2 maturation: E/P/D disaggregation, weight offloading, multi-layer MTP KV cache, encoder
  CUDA graphs, `thinking_token_budget` support.

### 2.4 v0.26.0 — 2026-07-27 (411 commits, 212 contributors, 61 new) [official]

- **New Inkling model family** with full support stack: base modeling, piecewise CUDA graph, Hopper FA4
  relative attention, MTP=1 speculative decoding, LoRA, standard ModelOpt NVFP4 quantization.
- **DeepSeek-V4 performance push:** specialized routing kernel (2.94% E2E TPOT, PR-reported), `fused_topk_bias`
  (1.5–2× kernel, PR-reported), redundant repeat/copy removal (1.8% E2E TPOT, PR-reported); ROCm two-stage
  compressor for HCA prefill; sparse decode/prefill optimizations; DSpark speculative decoding on AMD and XPU.
- **`head_dtype`**: fp32 `lm_head` for generation models (also extended to LoRA path, ROCm `torch.mm` fast
  path) — improves generation-head accuracy.
- **Flexible attention backends:** per-KV-cache-group attention backend selection; sliding-window support as
  explicit backend capability (hybrid-model support).
- **KV offloading & tiered secondary storage matured:** offloading metrics, tier-owned event handling,
  object-store secondary tier with workload identity, DP-replica-aware tiering, encoder-cache (EC)
  connectors incl. CPU offloading.
- **Rust frontend:** multimodal video + audio, Seed-OSS tool parser, native `vllm-bench` port.
- Transformers 5.13.0; more models migrated to Transformers backend (Olmo/Olmo2, MistralLarge3, HunyuanVL).
- Source: https://github.com/vllm-project/vllm/releases/tag/v0.26.0 [official]

### 2.5 Notable earlier-2026 items (from secondary summaries) [secondary][unverified for exact PR claims]

- v0.16.0 (2026-02-25): V1 is the only engine — V0 fully removed (as of v0.11+, V0 code fully removed). [secondary]
- 2026-09-04 weekly summary: Qwen3.8-Flash-Next support added (#53896), EC Connector P2P NIXL + CPU EC
  connector (#47941), "no breaking changes" that week. [secondary]
- Kimi K3 blog (vLLM blog, 2026-07-27): 2.8T-parameter multimodal MoE (16 of 896 experts/token, 1M-token
  context); vLLM serves it at 118 tok/s w/o spec decode, **370 tok/s (3.14×) with DSpark on 16× NVIDIA
  GB300 NVL72**; DSpark (block-diffusion spec decode, trained with vLLM + TorchSpec, open-sourced by
  **Inferact**); hybrid prefix caching redesigned over recurrent KDA state. [official]
  (Source: https://github.com/vllm-project/vllm-project.github.io/blob/HEAD/_posts/2026-07-27-k3.md)

---

