---
id: ai-industry-kb-2026/10-training-fine-tuning/nvfp4-export-available-since-july-2026-export-and-serve-not-
title: "NVFP4 export: available since July 2026 — export-and-serve, not training"
domain: training-fine-tuning
role: deep-dive
task: training
actors: ["Alibaba", "Apple", "DeepSeek", "Hugging Face", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-05", "2026-07", "2026-08", "2026-09", "2026-09-15", "2026-09-22"]
keywords: ["nvfp4", "training", "attention", "awq", "benchmark", "benchmarks", "consumer", "deepseek", "embedding", "fine-tuning", "funding", "gguf"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5474, 5523]
section: "10. Training & Fine-Tuning"
sha256: 1c8dd26fd4e374cc9f9387c7f61e8bbab897ba3469809910758e02b7fdcd36ad
---

# NVFP4 export: available since July 2026 — export-and-serve, not training

```python
model.save_pretrained_gguf("model_gguf", tokenizer, quantization_method="q4_k_m")
#multi-quant in ONE call:
model.save_pretrained_gguf("model_gguf", tokenizer,
                            quantization_method=["q4_k_m", "q8_0", "f16"])
model.push_to_hub_gguf("username/repo", tokenizer,
                       quantization_method=["q4_k_m", "q8_0", "f16"])
```

- **Full accepted `quantization_method` list (23 entries, verbatim)**: `not_quantized`, `fast_quantized`, `quantized`, `f32`, `f16`, `q8_0`, `q4_k_m`, `q5_k_m`, `q2_k`, `q3_k_l`, `q3_k_m`, `q3_k_s`, `q4_0`, `q4_1`, `q4_k_s`, `q4_k`, `q5_k`, `q5_0`, `q5_1`, `q5_k_s`, `q6_k`, `iq2_xxs`, `iq2_xs`, `iq3_xxs`, `q3_k_xs`.
- **Documented practical guidance**: `q4_k_m` = general default (~4.5 bits/weight, small quality hit); `q5_k_m` = better quality, larger; `q8_0` = near-lossless; `f16` = reference/downstream-requant base; `iq2_xxs`/`iq2_xs`/`iq3_xxs` = extreme compression, quality must be verified. **Apple-silicon note**: `q8_K_XL` upcasts layers to BF16 and is slower than plain bf16 on Macs — use `q8_0` on Macs instead.
- **Save-time OOM guard**: saving large models can OOM even when training fit; `maximum_memory_usage` (default 0.75) can be lowered on the save call. Merging LoRA into a GGUF is the memory peak most users actually hit, and Unsloth exposes a throttle for it.
- Companion save methods: `merged_16bit`, `merged_4bit`, `lora`-only, `push_to_hub_merged`.
- **Limitation**: no `iq4_xs`-style UD quants appear in this official list — the Dynamic v3.0 `UD-*` quant family (see §8) is a separate publishing pipeline (`unsloth/` GGUF org on HF), not a `save_pretrained_gguf` argument. The "one-command LoRA merging to quantized GGUF K-quants" claim is accurate for K-quants/i-quants but does not extend to Unsloth's proprietary UD quants via this API.

### NVFP4 export: available since July 2026 — export-and-serve, not training

- Unsloth's NVFP4 export (v0.1.481-beta, July 2026) predates and parallels the Axolotl NVFP4 MoE LoRA work (July 2026) — present them as parallel tracks, not one source.
- Axolotl's July track: NVFP4 (4-bit) MoE LoRA via ScatterMoE (W4A16) and SonicMoE (W4A4), including lossless adapter merge back into a plain NVFP4 checkpoint.
- **Nuance to keep**: the community `unsloth-cli` fine-tuning docs (September 2026) are explicit that on Thor (JetPack R38.2.2), LFM2.5-1.2B `nvfp4` and `awq` exports were live-tested for **serving** via `vllm-openai:v0.29.0-aarch64` — **not for training**. All Unsloth/NVFP4 claims to date are export-and-serve claims, not NVFP4 *training* claims.

### Qwen3.8-27B fine-tuning guide (September 2026)

Unsloth's official Qwen3.8 guide (docs page live September 2026) documents per-model figures weaker than the generic headline — useful evidence for the benchmark-methodology section:

- **Qwen3.8-27B** (dense 27B unified vision-language model; native text/image/video; thinking controls; 262K context; uses the `qwen3_5` architecture; requires Transformers v5).
- **Vendor training figures for this model: ~1.5x faster with ~50% less VRAM than FA2 setups** (no accuracy loss) — *below* the 2x/70% headline, illustrating that "2x/70%" is a class of improvement, not a contract.
- **VRAM tiers, explicitly stated: QLoRA works with 24GB; LoRA needs >36GB; full fine-tuning (FFT) works but uses 4x more VRAM.** This is Unsloth's own documentation confirming that 27B-class LoRA already exceeds 24GB — reinforcing the verdict that 70B on a single 24 GB card is not demonstrated.
- **New kernel detail: "Flash Linear Attention kernels"** (Gated DeltaNet) for Qwen3.8 training; first run takes longer while kernels compile.
- **Memory mechanics**: `offload_embedding=True` keeps the large untied input embedding in RAM, reducing resident VRAM (auto-disabled on unsupported platforms).
- **RL path**: Qwen3.8 RL uses the Qwen3.5 RL path with fast vLLM inference disabled.
- Export follows the one-call API: `save_pretrained_gguf("qwen38_gguf", tokenizer, quantization_method="q4_k_m")`, plus `push_to_hub_gguf`, merged-16bit for vLLM, or LoRA-adapters-only save.

### MoE-on-modest-hardware: documented fitting mechanisms (not new benchmarks)

- **MoE expert layers → system memory** (Studio v0.1.501-beta): automatic GPU placement pushes MoE expert layers into RAM so larger models fit; also added split-model-across-GPUs, tensor parallelism, per-model-per-quant hardware profiles, improved GGUF memory estimates. This is a Studio **placement/inference** feature — keep it in a separate evidence bucket from training kernels.
- **Apple Silicon**: fine-tune large MoE models (text or images) via MLX (v0.1.805/806-beta); up to **30x faster follow-up turns** in long Qwen chats [VENDOR].
- **Training/export fixes** (v0.1.501-beta): text-only training with multimodal models no longer truncates long examples before packing; **fine-tuned Qwen3.5 and Qwen3.6 MTP models now export correctly to GGUF** (previously broken); Windows GGUF-export final-write permission error fixed.
- Also in v0.1.501-beta: opt-in **MCP endpoint** (`UNSLOTH_STUDIO_ENABLE_MCP=1`) letting AI clients inspect models and training history, start/stop training, load checkpoints, validate recipes, and export GGUFs; dynamic NVFP4 models shipped in the preceding Studio release.
- Adjacent 2026 evidence that consumer training keeps pushing downward: a "Soup"-style layer-streaming approach (August 2026) fine-tunes 8B models on a 4 GB laptop GPU.

### Model-publishing cadence and third-party tooling (September 2026)

- Unsloth's Hugging Face organization (`unsloth/`, 1,374 models) shows uploads within days of 2026-09-22, including `unsloth/Laguna-S-2.1-GGUF` and the Dynamic v2.0/v3.0 lines (`unsloth/DeepSeek-V4-Flash-GGUF`, `unsloth/GLM-5.2-GGUF`, `unsloth/gemma-4-*-it-NVFP4`). **Dynamic v3.0 remains the current UD quant generation.**
- GLM-5.2's Dynamic GGUF was presented by a secondary source (LinkedIn) as 84% size reduction (~239 GB at 2-bit, ~217 GB at 1-bit) with ~82%/~76% accuracy retention and ~98% at 4-bit, runnable on 256 GB unified-memory machines — [UNVERIFIED] (vendor-originated, no independent benchmark).
- The community-built **`unsloth-cli`** wrapper (`agentculture/unsloth-cli`, releases **0.7.0 and 0.7.1, 2026-09-15**) — not an Unsloth product — operationalizes the export workflow: container-backed multi-format export (`sloth export --format merged-16bit | merged-4bit | gguf | awq | nvfp4` with `--quant`, `--calib`, `--calib-samples`, `--base`, `--force`, `--dry-run`, `--keep-intermediate`; disk estimates, no-clobber, atomic `.partial` output, provenance via `export.json` + per-adapter `exports.json`); **quantization-loss check** (`sloth eval --model DIR` scores merged/quantized outputs).
- Community economics datapoint: a third-party DGX Spark project (~May 2026) claims 7.67x LoRA / 8.35x full-FT speedup for Qwen3.5 (0.8B–27B) on a GB10 DGX Spark vs stock Unsloth, with wall-clock parity to a rented H100 (~20 h) at $0 vs ~$50 — [COMMUNITY, UNVERIFIED]; the author prices DGX Spark self-hosting breakeven at ~79 days of continuous training vs $2.49/h H100 rental.

### Correction ledger: World's Fair, funding, versioning

