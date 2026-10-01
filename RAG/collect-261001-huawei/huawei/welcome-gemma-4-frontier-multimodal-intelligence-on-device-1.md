---
id: collect-261001-huawei/huawei/welcome-gemma-4-frontier-multimodal-intelligence-on-device-1
title: "Cette image montre un personnage anthropomorphe ressemblant à un lapin, vêtu d'un manteau bleu et d'un pantalon beige, se tenant sur un chemin de terre dans un paysage rural idyllique ..."
domain: huawei
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["agentic", "apache", "attention", "benchmarks", "compute", "consumer", "context window", "cost", "embedding", "embeddings", "fine-tuning", "inference"]
source: docs/RAG/collect-261001-huawei/welcome-gemma-4-frontier-multimodal-intelligence-on-device.md
source_anchor: ""
source_lines: [1, 66]
sha256: ad7fedeb46ca05cf21922f3cdd23056faa97969b9ed2267d5c05957144f60137
---

# Cette image montre un personnage anthropomorphe ressemblant à un lapin, vêtu d'un manteau bleu et d'un pantalon beige, se tenant sur un chemin de terre dans un paysage rural idyllique ...

These models are the real deal: truly open with Apache 2 licenses, high quality with pareto frontier arena scores, multimodal including audio, and sizes you can use *everywhere* including on-device. Gemma 4 builds on advances from previous families and makes them click together. In our tests with pre-release checkpoints we have been impressed by their capabilities, to the extent that we struggled to find good fine-tuning examples because they are *so good* out of the box.

We collaborated with Google and the community to make them available everywhere: transformers, llama.cpp, MLX, WebGPU, Rust; you name it. This blog post will show you how to build with your favorite tools so let us know what you think!

Similar to Gemma-3n, Gemma 4 supports image, text, and audio inputs, and generates text responses. The text decoder is based on the Gemma model with support for long context windows. The image encoder is similar to the one from Gemma 3 but with two crucial improvements: variable aspect ratios, and configurable number of image token inputs to find your sweet spot between speed, memory, and quality. All models support images (or video) and text inputs, while the small variants (E2B, E4B) and the 12B Unified model support audio as well.

Gemma 4 comes in five sizes, all base and instruction fine-tuned:

| Model | Parameter Size | Context Window | Checkpoints | 
|---|---|---|---|
| Gemma 4 E2B | 2.3B effective, 5.1B with embeddings | 128k | base, IT | 
| Gemma 4 E4B | 4.5B effective, 8B with embeddings | 128k | base, IT | 
| Gemma 4 12B Unified | 11.95B dense, encoder-free | 256K | base, IT | 
| Gemma 4 31B | 31B dense model | 256K | base, IT | 
| Gemma 4 26B A4B | mixture-of-experts with 4B activated/26B total parameters | 256K | base, IT | 

Gemma 4 leverages several architecture components used in previous Gemma versions and other open models, and leaves out complex or inconclusive features such as Altup. The combination is a mix designed to be highly compatible across libraries and devices, that can efficiently support long context and agentic use cases, whilst being ideal for quantization.

As shown in the benchmarks above, this feature mix (combined with the training data and recipe) enables the 31B dense model to achieve an estimated LMArena score (text only) of 1452, while the 26B MoE reaches 1441 with just 4B active parameters 🤯. As we'll see, multimodal operation is comparatively as good as text generation, at least in informal and subjective tests.

These are the main architecture characteristics in Gemma 4:

- Alternating **local sliding-window** and**global full-context** attention layers. Smaller dense models use sliding windows of 512 tokens while larger models use 1024 tokens.
- **Dual RoPE** configurations: standard RoPE for sliding layers, pruned RoPE for global layers, to enable longer context.
- **Per-Layer Embeddings (PLE)** : a second embedding table that feeds a small residual signal into every decoder layer.
- **Shared KV Cache** : the last N layers of the model reuse key-value states from earlier layers, eliminating redundant KV projections.
- **Vision encoder (except 12B)** : uses learned 2D positions and multidimensional RoPE. Preserves the original aspect ratios and can encode images to a few different token budgets (70, 140, 280, 560, 1120).
- **Audio encoder (E2B, E4B)** : USM-style conformer with the same base architecture as the one in Gemma-3n.
- **Unified encoder-free multimodality (12B only)** : the 12B Unified variant drops the separate vision and audio encoders and projects raw image patches and audio waveforms directly into the LLM's embedding space. See Unified Multimodal (12B).

One of the most distinctive features in smaller Gemma 4 models is Per-Layer Embeddings (PLE), which was introduced previously in Gemma-3n. In a standard transformer, each token gets a single embedding vector at input, and the same initial representation is what the residual stream builds on across all layers, forcing the embedding to frontload everything the model might need. PLE adds a parallel, lower-dimensional conditioning pathway alongside the main residual stream. For each token, it produces a small dedicated vector for every layer by combining two signals: a token-identity component (from an embedding lookup) and a context-aware component (from a learned projection of the main embeddings). Each decoder layer then uses its corresponding vector to modulate the hidden states via a lightweight residual block after attention and feed-forward. This gives each layer its own channel to receive token-specific information only when it becomes relevant, rather than requiring everything to be packed into a single upfront embedding. Because the PLE dimension is much smaller than the main hidden size, this adds meaningful per-layer specialization at modest parameter cost. For multimodal inputs (images, audio, video), PLE is computed before soft tokens are merged into the embedding sequence — since PLE relies on token IDs that are lost once multimodal features replace the placeholders. Multimodal positions use the pad token ID, effectively receiving neutral per-layer signals.

The **shared KV cache** is an efficiency optimization that reduces both compute and memory during inference. The last `num_kv_shared_layers` layers of the model don't compute their own key and value projections. Instead, they **reuse** the K and V tensors from the last non-shared layer of the same attention type (sliding or full).

In practice, this has a minimal impact on quality while being much more efficient (in terms of both memory and compute) for long context generation and on-device use.

The Gemma 4 12B Unified model has no separate vision or audio encoder. Instead, raw image patches and audio waveforms are projected directly into the LLM's embedding space through lightweight linear layers, and all modalities flow into a single decoder-only transformer. This reduces multimodal latency and allows the whole model to be fine-tuned in one pass. The checkpoint size makes it deployment-friendly on consumer hardware.

We saw in our tests that Gemma 4 supports comprehensive multimodal capabilities out of the box. We don't know what was the training mix, but we had success using it for tasks such as OCR, speech-to-text, object detection, or pointing. It also supports text-only and multimodal function calling, reasoning, code completion and correction.

Here, we show a few inference examples across different model sizes. You can run them conveniently with this notebook. We encourage you to try the demos and share them below this blog!

Gemma 4 was trained with a specific convention to interleave input modalities:

- **Image content goes *before* the text** in your prompt.
- **Audio content goes *after* the text** in your prompt.

To be specific, this is a correct fragment that prepares inputs for the chat template:

```
image_url = "https://huggingface.co/datasets/huggingface/documentation-images/resolve/0052a70beed5bf71b92610a43a52df6d286cd5f3/diffusers/rabbit.jpg"
audio_url = "instructions.m4a"      # A recording of "describe this image"
messages = [
    {
        "role": "user",
        "content": [
            {"type": "image", "image": image_url},
            {"type": "text", "text": "Answer in French."},
            {"type": "audio", "audio": audio_url},
        ]
    }
]
# Cette image montre un personnage anthropomorphe ressemblant à un lapin, vêtu d'un manteau bleu et d'un pantalon beige, se tenant sur un chemin de terre dans un paysage rural idyllique ...
```
The audio inference snippets in this post don't follow this convention, as they were written before this guidance was confirmed. We are keeping them as originally written so result tables remain reproducible. For your own code, prefer the order above.

