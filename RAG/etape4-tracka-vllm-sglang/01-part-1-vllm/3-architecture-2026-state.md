---
id: etape4-tracka-vllm-sglang/01-part-1-vllm/3-architecture-2026-state
title: "3. Architecture (2026 state)"
domain: part-1-vllm
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "DeepSeek", "Huawei", "Moonshot", "OpenAI", "vLLM"]
dates: ["2026-09-09"]
keywords: ["amd", "ascend", "decode", "deepseek", "diffusion", "disaggregated", "fp8", "gpu", "gqa", "kimi", "kv cache", "multimodal"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [240, 290]
section: "PART 1 — vLLM"
sha256: b2494226bbd4992231aca1543c20ccd750913fcbd45954cac43fc15b4401342b
---

# 3. Architecture (2026 state)

## 3. Architecture (2026 state)

### 3.1 Core ideas (unchanged foundations)

- **PagedAttention** (SOSP 2023): KV cache split into fixed-size physical blocks (typical 16 tokens),
  per-request page table maps logical→physical blocks; custom CUDA kernels gather non-contiguous KV.
  [official][secondary]
- **Continuous batching / iteration-level scheduling:** requests join/leave between forward passes.
  [official]
- **Prefix caching (Automatic Prefix Caching, APC):** block-level hash-chained KV reuse; in V1,
  **always-on with near-zero overhead** (constant-time eviction). In v0.29: Mamba internal prefill
  checkpoints (+9–25% TTFT, PR-reported); `prefix_cache_retention_interval` CLI arg defaulting to 0;
  deterministic `NONE_HASH`. [official][secondary]

### 3.2 V1 engine → Model Runner V2 (2026)

- **V1** (only engine since V0 removal, v0.11+; confirmed as sole engine in v0.16.0, Feb 2026):
  isolated EngineCore process (scheduler + executor) separate from API server; async single-step
  scheduling with unified token budget `{request_id: num_tokens}`; persistent batch (cached tensors,
  incremental diffs); stateful workers (only diffs sent); chunked prefill **on by default**; APC
  **on by default**; piecewise CUDA graphs + torch.compile; async multimodal preprocessing with
  encoder cache. [official][secondary]
- **Model Runner V2** (MRV2): became **default for all models in v0.29.0 (2026-09-09)**; MRV1 deprecated,
  targeted removal in **v0.32**; fallback to MRV1 when sequence parallelism / dual-batch overlap /
  elastic EP / custom logits processors / certain spec-decode methods are configured. Features in
  v0.30: dual-batch overlap in eager + FULL CUDA graphs, PP speculative decoding, online acceptance
  estimator for adaptive verification, GC-frozen graph capture (12s→2s capture, 28.9s→8.2s init on H200,
  PR-reported). [official]
- **Tensor parallelism** (multi-GPU; ~7.2× throughput on 8× A100 per one secondary explainer),
  **pipeline parallelism**, cross-node PP via InfiniBand; **data parallelism**; **expert parallelism**
  incl. **Elastic EP** (reconfiguration reusing CUDA graphs; DeepEP v2, FlashInfer one-sided All2All).
  [official][secondary]
- **Disaggregated prefill/decode:** first-class support via **NIXL connector** and **Mooncake** connectors;
  v0.29 added NIXL P/D DCP for MLA models, Mooncake Store decode KV saving (`save_decode_cache`), hybrid
  DCP prefix caching, EC connectors (P2P NIXL + CPU). **Context parallelism:** DCP (decode) and PCP
  (prefill) for sparse-MLA models; FlashInfer native CP for MLA decode; FlashMLA sparse DCP on Hopper.
  [official]
- **KV-cache management 2026:** KV offloading & tiered secondary storage (object-store secondary tier with
  workload identity, DP-replica-aware tiering, EC connectors incl. CPU offload); disk offload for large
  contexts (v0.28); HiSparse host-resident tier for sparse-MLA decode (v0.30); FP8/C8 KV cache (incl. GQA,
  DeepSeek-V3.1 PD scenario in Ascend plugin); MXFP8/NVFP4 KV caches (`nvfp4_fp8_ds_mla`, `fp8_ds_mla`);
  SSD support for multiple DP ranks in Mooncake offload dirs (experimental, Ascend plugin). [official]
- **Speculative decoding (2026):** EAGLE3 (incl. for Sarvam MLA), DSpark (block-diffusion; 3.14× on Kimi K3,
  GB300 NVL72, per vLLM blog), DFlash2 (local convolution + candidate selector), MTP multi-token prediction
  (Qwen3.8-Flash-Next native MTP, Nemotron, PLaMo3 EAGLE-3/DFlash), P-Eagle and PARD (stable parallel spec
  methods per Ascend plugin notes), DSpark on AMD and XPU (v0.26), adaptive verification (online acceptance
  estimator, v0.30), per-request acceptance stats in OpenAI responses (`--per-request-spec-decode-metrics`,
  v0.29). [official][secondary]

---

