---
id: collect-240926-huggingface/huggingface/coherelabs-north-small-translate-1-0-w4a16-hugging-face
title: "This is for H100, adjust tp for your device"
domain: huggingface
role: reference
task: reference
actors: ["China", "Cohere", "Hugging Face", "Nvidia", "vLLM"]
dates: []
keywords: ["agentic", "attention", "blackwell", "cohere", "embeddings", "fp4", "fp8", "gpu", "gpus", "license", "moe", "nvfp4"]
source: docs/RAG/clean_en/huggingface/coherelabs-north-small-translate-1-0-w4a16-hugging-face.md
source_anchor: ""
source_lines: [1, 88]
sha256: b4ec04395dc10df014647eabf6d7745256ec3fe7a8d849b0cbfaa474ef2db5bd
---

# This is for H100, adjust tp for your device

<!-- source: https://huggingface.co/CohereLabs/North-Small-Translate-1.0-w4a16 -->

## You need to agree to share your contact information to access this model

This repository is publicly accessible, but you have to accept the conditions to access its files and content.

By submitting this form, you agree to the License Agreement and acknowledge that the information you provide will be collected, used, and shared in accordance with Cohere’s Privacy Policy. You’ll receive email updates about Cohere Labs and Cohere research, events, products and services. You can unsubscribe at any time.

Log in or Sign Up to review the conditions and access this model content.

North Small Translate is an open weights research release of a sparse Mixture-of-Experts model with 25 billion active parameters and 218 billion total parameters, specialized for high-quality machine translation across 50 languages. This repository holds the **NVFP4 W4A16** checkpoint.

Developed by: Cohere and Cohere Labs

- Point of Contact: **Cohere Labs**
- License: CC BY-NC 4.0, requires also adhering to **Cohere Lab's Acceptable Use Policy**
- Model: North-Small-Translate-1.0-w4a16
- Model Size: 25B active parameters, 218B total parameters
- Context length: 16K input & 16K output
- Quantization: NVFP4 W4A16

**Try North Small Translate**

You can try out North Small Translate before downloading the weights in our hosted Hugging Face Space.

**Available quantizations**

The following quantizations are available, with example minimum GPU requirements.

| Quantization | Blackwell | Hopper | 
|---|---|---|
| BF16 (16-bit) | 4 x B200 | 8 x H100 | 
| FP8 (8-bit) | 2 x B200 | 4 x H100 | 
| NVFP4 W4A16 (4-bit weights) | 1 x B200 | 2 x H100 | 

All three variants are the checkpoints Cohere serves in production for this model.

**Usage**

To use this model in transformers, please use our BF16 model weights. This NVFP4 W4A16 checkpoint is designed to be served with vLLM and is not compatible with transformers, which has no native 4-bit support for this format.

**vLLM**

Accurate response parsing requires installing Cohere's `melody` library.

```
uv pip install vllm
uv pip install cohere_melody>=0.9.0
```
Then the vLLM server can be started with the following command:

```
# This is for H100, adjust tp for your device
vllm serve CohereLabs/North-Small-Translate-1.0-w4a16 \
  -tp 2 \
  --max-model-len 32768 \
  --tool-call-parser cohere_command4 \
  --reasoning-parser cohere_command4 \
  --enable-auto-tool-choice
```
We recommend greedy decoding, which is what the production deployment of this model uses.

**System instructions**

The chat template applies a default system instruction that names the model and sets its safety defaults. You can replace it by passing your own `platform_instruction_override` when applying the chat template, for example through vLLM's `chat_template_kwargs`. Per-conversation instructions can also be supplied as an ordinary `system` message, which is appended after the default instruction rather than replacing it.

**Input**: Text only.

**Output**: Model generates text.

**Model Architecture**: North Small Translate is a decoder-only sparse Mixture-of-Experts Transformer model. With 25B active parameters and 218B total parameters, it has 128 experts, of which 8 are activated per token, alongside shared experts applied to every token. The attention layers interleave sliding-window attention layers (window size 4096) using Rotary Positional Embeddings with global attention layers without positional embeddings, in a 3:1 ratio, as first introduced in Command A. The router applies a sigmoid activation over the expert logits and normalizes over the selected top-k. The model was post-trained specifically for translation quality.

**Quantization Methodology:** This checkpoint uses NVFP4 W4A16 quantization — 4-bit weights with 16-bit activations — in the `compressed-tensors` format, with a group size of 16 and FP8 scales. We quantize the MoE experts only, keeping the attention projections, the routers and the output head at higher precision. Since the experts hold the overwhelming majority of the model's parameters, this captures nearly all of the savings while leaving the most precision-sensitive parts untouched, reducing the footprint from roughly 437GB to 131GB. Because only the weights are quantized, this format does not require native FP4 hardware and runs on pre-Blackwell GPUs such as Hopper and Ada.

**Languages covered:** The model supports translation across 50 languages: English, Albanian, Arabic, Bulgarian, Bengali, Catalan, Czech, Danish, German, Greek, Spanish, Estonian, Persian, Finnish, Filipino, French, Irish, Hebrew, Hindi, Croatian, Hungarian, Indonesian, Icelandic, Italian, Japanese, Korean, Lithuanian, Latvian, Malay, Maltese, Dutch, Norwegian, Punjabi, Polish, Portuguese, Romanian, Russian, Slovak, Slovenian, Serbian, Swedish, Tamil, Telugu, Thai, Turkish, Ukrainian, Urdu, Vietnamese, Traditional Chinese, Simplified Chinese.

**Context Length:** North Small Translate supports a context length of 16K input & 16K output.

**Figure 1. WMT26 performance across all evaluated languages.** North Small Translate scores 83.60, increasing to 84.36 with the agentic multi-pass translation workflow. Figure and evaluation details are from the North Small Translate launch blog post.

For errors or additional questions about details in this model card, contact labs@cohere.com.

We hope that the release of this model will make community-based research efforts more accessible, by releasing the weights of a highly performant translation model to researchers all over the world. This model is governed by a CC BY-NC 4.0 License (Non-Commercial) with an acceptable use addendum, *and also requires adhering to Cohere Lab's Acceptable Use Policy*. If you are interested in commercial use, please contact Cohere's Sales team.

You can use North Small Translate in our dedicated Hugging Face Space.

- Downloads last month
- 31
