---
id: etape4-tracka-vllm-sglang/01-part-1-vllm/4-quantization-support
title: "4. Quantization support"
domain: part-1-vllm
role: deep-dive
task: quantization
actors: ["AMD", "Alibaba", "Anthropic", "DeepSeek", "Intel", "MiniMax", "Moonshot", "Nvidia", "vLLM"]
dates: ["2026-02-12"]
keywords: ["quantization", "amd", "awq", "blackwell", "cohere", "decode", "deepseek", "embedding", "fp4", "fp8", "gguf", "gptq"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [291, 364]
section: "PART 1 — vLLM"
sha256: 3e1dabe36962f4e61b8458a20e7bfc7f174a90402d1cd1d6361290056443ed87
---

# 4. Quantization support

## 4. Quantization support

Single dispatch point `vllm/model_executor/layers/quantization/__init__.py` (`--quantization` flag).
State as of Aug–Sep 2026 [secondary, verified against source tree by third-party researchers]:

| Format | Scheme / notes |
|---|---|
| FP8 | `fp8` W8A8 (per-tensor/per-channel; compressed-tensors FP8 — "Hopper default"); `modelopt` (ModelOpt-exported FP8); KV cache FP8 (`fp8_ds_mla`, `nvfp4_fp8_ds_mla`); FP8 MoE kernels (block-FP8 fused shared experts) |
| NVFP4 / MXFP4 | `modelopt_fp4` (native on SM100/103 Blackwell; emulated from SM75); `mxfp4` / `gpt_oss_mxfp4` (GPT-OSS ships MXFP4); Kimi-K3 NVFP4; FlashInfer CuTeDSL NVFP4 W4A16 default over Marlin on SM100/103 (v0.30); DeepSeek V4.1 whole KV in MXFP8 (v0.30) |
| MXFP8 | `mxfp8` online (v0.19+); `modelopt_mxfp8`; FlashInfer TRT-LLM MXFP8 linear backend (v0.29); Humming MoE MXFP4 weights + block-FP8 activations (v0.29) |
| AWQ / GPTQ | `awq_marlin`, `gptq_marlin` W4A16 (Marlin kernels; recommended); legacy `awq`/`gptq` unfused deprecated; ⚠️ GPTQ `g_idx` activation ordering **removed in v0.30** |
| INT8 | `int8_per_channel_weight_only`; Marlin INT4/INT8; `experts_int8` (use int8_per_channel_weight_only) |
| Online quantization | `online` + shorthands (`fp8_per_tensor`, `fp8_per_block`, `fp8_per_channel`, `mxfp8`, `int8_per_channel_weight_only`, `nvfp4_per_token`) — quantize from BF16 checkpoint at load; **targeted online quantization** via `quantization_config.targets` (v0.30) |
| compressed-tensors | neuralmagic/Red Hat llm-compressor output; SpinQuant + QuaRot integrated via llm-compressor (vLLM v0.11+) [secondary] |
| AMD | `quark` (ROCm path); AMD Quark NVFP4 (v0.28); W4A4 preshuffled asm GEMM default on ROCm (+15% Llama-3.3-70B MXFP4, PR-reported) |
| Intel | `inc` / AutoRound (Intel Neural Compressor; AutoRound 2/3/5/6/7-bit on CUDA, v0.30); XPU INC int4 W4A8 linear backend (v0.29) |
| Others | torchao, GGUF (plugin, experimental/under-optimized), ModelOpt mixed (`modelopt_mixed`), Humming (W4A8 FP4×FP8, Ant Group), CPU AWQ (`cpu_awq` legacy) |
| Deprecated/removed | `fbgemm_fp8`, `fp_quant`, `moe_wna16` (legacy); **bitsandbytes → out-of-tree plugin** (v0.28, breaking) |

Per-quantization **linear backend overrides** added in v0.30 (`--linear-backend` choices: auto, cutlass,
deep_gemm, flashinfer_cutlass, flashinfer_cudnn, marlin, machete, humming, torch, triton, emulation).
[secondary; release notes official]

---

## 5. 2026 feature surface (serving features)

### 5.1 Structured outputs / JSON mode [official][secondary]
- Guided decoding constrains generation to JSON schemas, regex, grammars, choices
  (`guided_json`, `guided_regex`, `guided_choice`, `guided_grammar`).
- Structured-outputs engine: default backend `auto` selects **xgrammar** or guidance/llguidance; <5% latency
  impact claimed by third-party guides [secondary].
- 2026-02-12: generic `structured_outputs` enabled for the **Responses API** (grammars/regexes/choices
  beyond JSON schema; #33709). [secondary]
- v0.30: XGrammar V4.1 schema constraints for strict tool parameters (DeepSeek-V4.1); XGrammar termination
  in batches fixed (v0.29); GPT-OSS Harmony strict grammar. [official]

### 5.2 Reasoning (multi-step) serving [official][secondary]
- Reasoning parsers (token-attributed text in Rust frontend, 2026-09; `reasoning_token_count` support),
  `reasoning_effort` prompts (DeepSeek V4, v0.28), `enable_thinking` defaults (Gemma4), `thinking_token_budget`
  (MRV2, v0.28), reasoning-end detection scoped to current turn, Kimi K3 reasoning/tool parsers (v0.30).
- ⚠️ `reasoning_content` output removal was a documented **breaking client change** in v0.28.

### 5.3 LoRA serving [official]
- Multi-LoRA serving: multiple adapters on one base model, per-request selection via `--lora-modules`
  / `--enable-lora` / client `model=name`; supports dense and MoE layers.
- v0.29: LoRA for DeepSeek V4, Qwen3-Omni multimodal LoRA, LLaVA-NeXT tower/connector LoRA, LFM2-VL,
  Qwen3.5 embedding modules, partial LoRA on Qwen3.5/3.6 GatedDeltaNet; fp32 `lm_head` via `head_dtype`
  extended to LoRA path (v0.26). v0.30: rsLoRA scaling in MoE expert packing, `--lora-modules name=path`
  with `=` in path. [official]

### 5.4 Multi-modal (vision/audio/video) [official]
- Images, **video** (torchcodec backend, MP4 edit-list-aware frame sampling, encoder CUDA graphs for
  MiniCPM-V 2.5/2.6/4.0, video embeds via Python frontend v0.29), **audio** (torchcodec audio decoding,
  torchaudio resampler, `audio_backend` selectable in `--media-io-kwargs`, `/v1/audio/transcriptions`,
  `/v1/realtime` WebSocket), encoder cache for chunked-prefill VLMs, encoder-cache sharing over NIXL and
  Mooncake (v0.30), Qwen3-VL-235B / MiniMax-M3 / Gemma-4 / Llama-4-era VL support, multimodal LoRA.
- Rust frontend gained multimodal video + audio (v0.26). [official]

### 5.5 Prefix caching
- APC always-on in V1; Mamba internal prefill checkpoints (9–25% TTFT, PR-reported, v0.29); hybrid prefix
  caching redesigned for Kimi K3's recurrent KDA state (benefits all hybrid linear models, per vLLM blog);
  K3 DCP partial prefix cache hits (v0.29); K3 internal prefix checkpoints with partial prefix caching +
  spec decode (v0.30); Mooncake hybrid DCP prefix caching. [official]

### 5.6 Other 2026 service features
- **Watermarking** (Gumbel-max, v0.30); **Fast Start** IPC weight cache (v0.30); **RL weight sync**
  (`sharded_rdt` P2P over NIXL/Ray, v0.29); **scale-out endpoints** opt-in via `--enable-scale-out` (v0.30);
  admission control flags; `/v1/messages/render` (Anthropic) + `/cohere/v2/chat/render`; pooling/embedding
  models (ModernBERT, BGE-M3, Qwen3-Embedding on TPU); tool calling (Seed-OSS parser, Gemma4/Harmony
  strict grammars). [official]

---

