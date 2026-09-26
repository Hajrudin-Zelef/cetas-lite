---
id: collect-240926-huggingface/huggingface/unsloth-gemma-4-31b-it-qat-gguf-hugging-face-1
title: "Read our How to Run Gemma 4 QAT Guide!"
domain: huggingface
role: reference
task: reference
actors: ["Google", "Hugging Face", "Unsloth", "vLLM"]
dates: []
keywords: ["agentic", "agents", "apache", "attention", "benchmarks", "consumer", "context window", "embedding", "embeddings", "gguf", "gpus", "inference"]
source: docs/RAG/clean_en/huggingface/unsloth-gemma-4-31b-it-qat-gguf-hugging-face.md
source_anchor: ""
source_lines: [1, 87]
sha256: 698de8d08041a59f77b7052e182ae8a1929cce640cdd8fd6ad7cbe24e0564617
---

# Read our How to Run Gemma 4 QAT Guide!

<!-- source: https://huggingface.co/unsloth/gemma-4-31B-it-qat-GGUF -->

# Read our How to Run Gemma 4 QAT Guide!

    *See Unsloth Dynamic 2.0 GGUFs for our quantization benchmarks.*
  

- **Jun 9 Update:** Added MTP support. See our MTP Guide.
- Gemma 4 can now be run and fine-tuned in Unsloth Studio. Read our guide.
- See all versions of Gemma 4 QAT (GGUF, 16-bit etc.) in our collection.
- Example of Gemma 4 E4B (4-bit GGUF) running in Unsloth Studio with tool-calling:

This model ships a Multi-Token Prediction drafter at the repo root (`mtp-gemma-4-31B-it.gguf`, a near-lossless smart Q4_0). A recent llama.cpp auto-discovers it from `-hf`, so you do not pass `--model-draft`:

```
./build/bin/llama-server \
  -hf unsloth/gemma-4-31B-it-qat-GGUF:UD-Q4_K_XL \
  --spec-type draft-mtp --spec-draft-n-max 4 \
  -ngl 999 -fa on
```
The drafter shares the target's KV cache and does not change the output (the target verifies every drafted token). See the `MTP/` folder for the other precisions and explicit usage.

    Hugging Face |
    GitHub |
    Launch Blog |
    Documentation
    

    **License**: Apache 2.0 | **Authors**: Google DeepMind

This model card is for the new versions of the Gemma 4 family optimized with Quantization-Aware Training (QAT), which allows preserving similar quality to bfloat16 while dramatically reducing the memory requirements to load the model. Four versions of the QAT checkpoints are available:


**Unquantized QAT checkpoints** (Q4_0): Half-precision weights extracted from the QAT pipeline, ideal for custom downstream compilation and research. Available for Gemma 4 E2B, E4B, 12B, 26B A4B, and 31B, and their drafter models.
**GGUF** (Q4_0): Ready-to-deploy formats for broad ecosystem compatibility. Available for Gemma 4 E2B, E4B, 12B, 26B A4B, and 31B.
**Mobile-optimized** (wNa8o8): A custom schema engineered explicitly for mobile hardware efficiency. It features targeted 2-bit decoding layers, optimized KV caches, and static activations to maximize VRAM savings. Available for Gemma 4 E2B and E4B.
**Compressed Tensors** (w4a16): QAT checkpoints serialized in the compressed-tensors format for native, optimized inference with vLLM. Available for Gemma 4 E2B, E4B, 12B, and 31B.

Gemma is a family of open models built by Google DeepMind. Gemma 4 models are multimodal, handling text and image input (with audio supported on E2B, E4B, and 12B) and generating text output. This release includes open-weights models in both pre-trained and instruction-tuned variants. Gemma 4 features a context window of up to 256K tokens and maintains multilingual support in over 140 languages.

Featuring both Dense and Mixture-of-Experts (MoE) architectures, Gemma 4 is well-suited for tasks like text generation, coding, and reasoning. The models are available in five distinct sizes: **E2B**, **E4B**, **12B**, **26B A4B**, and **31B**. Their diverse sizes make them deployable in environments ranging from high-end phones to laptops and servers, democratizing access to state-of-the-art AI.

Gemma 4 introduces key **capability and architectural advancements**:

- **Reasoning** – All models in the family are designed as highly capable reasoners, with configurable thinking modes.
- **Extended Multimodalities** – Processes Text, Image with variable aspect ratio and resolution support (all models), Video, and Audio (featured natively on the E2B, E4B, and 12B models).
- **Diverse & Efficient Architectures** – Offers Dense and Mixture-of-Experts (MoE) variants of different sizes for scalable deployment.
- **Optimized for On-Device** – Smaller models are specifically designed for efficient local execution on laptops and mobile devices.
- **Increased Context Window** – The small models feature a 128K context window, while the medium models support 256K.
- **Enhanced Coding & Agentic Capabilities** – Achieves notable improvements in coding benchmarks alongside native function-calling support, powering highly capable autonomous agents.
- **Native System Prompt Support** – Gemma 4 introduces native support for the`system` role, enabling more structured and controllable conversations.

Gemma 4 models are designed to deliver frontier-level performance at each size, targeting deployment scenarios from mobile and edge devices (E2B, E4B) to consumer GPUs and workstations (12B, 26B A4B, 31B). They are well-suited for reasoning, agentic workflows, coding, and multimodal understanding.

The models employ a hybrid attention mechanism that interleaves local sliding window attention with full global attention, ensuring the final layer is always global. This hybrid design delivers the processing speed and low memory footprint of a lightweight model without sacrificing the deep awareness required for complex, long-context tasks. To optimize memory for long contexts, global layers feature unified Keys and Values, and apply Proportional RoPE (p-RoPE).

| Property | E2B | E4B | 12B Unified | 31B Dense | 
|---|---|---|---|---|
| **Total Parameters** | 2.3B effective (5.1B with embeddings) | 4.5B effective (8B with embeddings) | 11.95B | 30.7B | 
| **Layers** | 35 | 42 | 48 | 60 | 
| **Sliding Window** | 512 tokens | 512 tokens | 1024 tokens | 1024 tokens | 
| **Context Length** | 128K tokens | 128K tokens | 256K tokens | 256K tokens | 
| **Vocabulary Size** | 262K | 262K | 262K | 262K | 
| **Supported Modalities** | Text, Image, Audio | Text, Image, Audio | Text, Image, Audio | Text, Image | 
| **Vision Encoder Parameters** | *~150M* | *~150M* | - | *~550M* | 
| **Audio Encoder Parameters** | *~300M* | *~300M* | - | No Audio | 

The "E" in E2B and E4B stands for "effective" parameters. The smaller models incorporate Per-Layer Embeddings (PLE) to maximize parameter efficiency in on-device deployments. Rather than adding more layers or parameters to the model, PLE gives each decoder layer its own small embedding for every token. These embedding tables are large but are only used for quick lookups, which is why the effective parameter count is much smaller than the total.

The "Unified" in Gemma 4 12B Unified refers to its encoder-free architecture. Other Gemma 4 models use dedicated encoders to process multimodal data before passing it to the LLM. Gemma 4 12B eliminates these encoders entirely, projecting raw image patches and audio waveforms directly into the LLM's embedding space through lightweight linear layers. This unified approach means all modalities flow straight into a single decoder-only transformer, reducing multimodal latency and allowing the entire model to be fine-tuned in one pass.

| Property | 26B A4B MoE | 
|---|---|
| **Total Parameters** | 25.2B | 
| **Active Parameters** | 3.8B | 
| **Layers** | 30 | 
| **Sliding Window** | 1024 tokens | 
| **Context Length** | 256K tokens | 
| **Vocabulary Size** | 262K | 
| **Expert Count** | 8 active / 128 total and 1 shared | 
| **Supported Modalities** | Text, Image | 
| **Vision Encoder Parameters** | *~550M* | 

The "A" in 26B A4B stands for "active parameters" in contrast to the total number of parameters the model contains. By only activating a 4B subset of parameters during inference, the Mixture-of-Experts model runs much faster than its 26B total might suggest. This makes it an excellent choice for fast inference compared to the dense 31B model since it runs almost as fast as a 4B-parameter model.

These models were evaluated against a large collection of different datasets and metrics to cover different aspects of text generation. Evaluation results marked in the table are for instruction-tuned models.

