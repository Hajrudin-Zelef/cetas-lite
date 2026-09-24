---
id: collect-240926-huggingface/huggingface/vontra-mimo-v2-6-flash-rl-mlx-4bit-mtp-hugging-face
title: "MiMo-V2.6-Flash-RL MLX 4-bit MTP"
domain: huggingface
role: reference
task: reference
actors: ["Apple", "Xiaomi"]
dates: []
keywords: ["attention", "decode", "embedding", "fp8", "kv cache", "license", "memory", "mit license", "moe", "mxfp4", "parameters", "quantization"]
source: docs/RAG/clean_en/huggingface/vontra-mimo-v2-6-flash-rl-mlx-4bit-mtp-hugging-face.md
source_anchor: ""
source_lines: [1, 66]
sha256: 4188d88de440d5ebaf9731f946a7a15650d3d6d2155db7b6a29ccdd7b94aa6c2
---

# MiMo-V2.6-Flash-RL MLX 4-bit MTP

<!-- source: https://huggingface.co/Vontra/MiMo-V2.6-Flash-RL-MLX-4bit-MTP -->

# MiMo-V2.6-Flash-RL MLX 4-bit MTP

A tested Apple Silicon conversion of XiaomiMiMo/MiMo-V2.6-Flash-RL, published by Vontra.

This is the text backbone of MiMo-V2.6-Flash-RL converted for MLX. The dense projections use 4-bit affine quantization with group size 64, while the model's native MXFP4 MoE experts remain in their original group-size-32 format. The resulting main model averages 4.257 bits per weight.

The checkpoint includes its native three-layer MTP payload in `mtp/model_mtp.safetensors`, converted to 4-bit affine weights. It also carries the upstream five-layer DFlash drafter, vision encoder, audio encoder, and audio tokenizer so those assets do not need a second download.

| Native component | Path | Format | 
|---|---|---|
| MTP predictor | `mtp/model_mtp.safetensors` | MLX 4-bit affine | 
| DFlash drafter | `dflash/model.safetensors` | Upstream BF16 | 
| Vision encoder | `omnimodal/vision_encoder.safetensors` | Upstream BF16 | 
| Audio encoder | `omnimodal/audio_encoder.safetensors` | Upstream BF16 | 
| Audio tokenizer | `audio_tokenizer/model.safetensors` | Upstream weights | 

Current oMLX and MLX text generation run the target model correctly but do not automatically execute MiMo's MTP, DFlash, vision, or audio paths. The speed figures below are serial text decode measurements. The auxiliary tensors and configs are packaged for MiMo-aware runtimes and ongoing MLX integration, not advertised as working oMLX controls.

Tested on a 256 GB M3 Ultra Mac Studio with oMLX 0.7.0.dev2, MLX 0.32.2, and mlx-lm 0.31.3.

| Test | Result | 
|---|---|
| Sustained generation, 128-token decode | 59.4 tok/s average | 
| Prompt processing, 512 tokens | 477.1 tok/s | 
| Prompt processing, 2,048 tokens | 562.8 tok/s | 
| Peak unified memory, short context | 164.3 GB | 
| Peak unified memory, 2,048-token prompt | 166.8 GB | 
| Quantized text model size on disk | about 154 GiB | 
| Complete repository size | about 160 GiB | 

The model produced correct arithmetic, a clear factual explanation, and coherent Python in repeated smoke tests. A 256 GB Mac is recommended so there is room for the model, KV cache, and the rest of the system.

```
pip install -U "mlx-lm>=0.31.3"
python -m mlx_lm generate \
  --model Vontra/MiMo-V2.6-Flash-RL-MLX-4bit-MTP \
  --prompt "Write a Python function that checks whether an integer is prime." \
  --max-tokens 256 \
  --temp 0.6
```
The upstream tokenizer chat template is included. In oMLX, download or select `Vontra/MiMo-V2.6-Flash-RL-MLX-4bit-MTP` as an MLX model.

MiMo-V2.6 stores fused attention tensors in checkpoint tensor-parallel order and pads the FP8 scale grid separately for each shard. This conversion reconstructs those shards before quantization. Skipping that step produces a model that loads but returns broken output.

The quantized MTP payload lives in its own `mtp/` directory with a manifest describing its tensors. It is not a standalone drafter for `mlx_lm.generate --draft-model` today. The upstream DFlash payload retains its trained mask embedding and corrected JSON config; an MLX smoke test matched serial greedy output, but it did not beat serial decode in the current experimental runtime.

The vision and audio tensors are kept outside the root text-model index so `mlx_lm` and oMLX continue to load the tested text checkpoint unchanged. `omnimodal/manifest.json` records every auxiliary path and its upstream source.

Xiaomi describes MiMo-V2.6-Flash-RL as a sparse 309B-parameter MoE with 15B active parameters, 48 transformer layers, 256 routed experts, and eight active experts per token. The full upstream release supports a one-million-token context and omnimodal inputs. See the original model card for architecture details, evaluations, deployment recipes, intended use, and limitations.

The upstream model is released under the MIT license. All model architecture, training, tokenizer work, and original branding belong to the Xiaomi MiMo team. This repository contains a community MLX conversion and measured Apple Silicon results.

```
@misc{mimo2026v26flash,
  title={MiMo-V2.6-Flash-RL},
  author={{Xiaomi MiMo Team}},
  year={2026},
  howpublished={\url{https://huggingface.co/XiaomiMiMo/MiMo-V2.6-Flash-RL}},
}
```
Follow Vontra for new Apple Silicon releases and fixes.

- Downloads last month
- 1,688
