---
id: collect-240926-huggingface/huggingface/unsloth-gemma-4-26b-a4b-it-gguf-hugging-face-1
title: "Read our How to Run Gemma 4 Guide!"
domain: huggingface
role: reference
task: reference
actors: ["Google", "Hugging Face", "Unsloth"]
dates: []
keywords: ["agentic", "agents", "apache", "attention", "benchmarks", "consumer", "context window", "embedding", "embeddings", "gguf", "gpus", "inference"]
source: docs/RAG/clean_en/huggingface/unsloth-gemma-4-26b-a4b-it-gguf-hugging-face.md
source_anchor: ""
source_lines: [1, 107]
sha256: 1e27ffa44be9b7a53d1a133f79227474aa5a93c2ee7c7b899f34e8a549f17568
---

# Read our How to Run Gemma 4 Guide!

<!-- source: https://huggingface.co/unsloth/gemma-4-26B-A4B-it-GGUF -->

# Read our How to Run Gemma 4 Guide!

    *See Unsloth Dynamic 2.0 GGUFs for our quantization benchmarks.*
  

- **Jun 9 Update:** Added MTP support. See our MTP Guide.
- **Apr 11 Update:** Re-download for Google's latest chat template and llama.cpp fixes.
- Gemma 4 can now be run and fine-tuned in Unsloth Studio. Read our guide.
- See all versions of Gemma 4 (GGUF, 16-bit etc.) in our collection.
- Example of Gemma 4 E4B (4-bit GGUF) running in Unsloth Studio with tool-calling:

    Hugging Face |
    GitHub |
    Launch Blog |
    Documentation
    

    **License**: Apache 2.0 | **Authors**: Google DeepMind

Gemma is a family of open models built by Google DeepMind. Gemma 4 models are multimodal, handling text and image input (with audio supported on small models) and generating text output. This release includes open-weights models in both pre-trained and instruction-tuned variants. Gemma 4 features a context window of up to 256K tokens and maintains multilingual support in over 140 languages.

Featuring both Dense and Mixture-of-Experts (MoE) architectures, Gemma 4 is well-suited for tasks like text generation, coding, and reasoning. The models are available in four distinct sizes: **E2B**, **E4B**, **26B A4B**, and **31B**. Their diverse sizes make them deployable in environments ranging from high-end phones to laptops and servers, democratizing access to state-of-the-art AI.

Gemma 4 introduces key **capability and architectural advancements**:

- **Reasoning** – All models in the family are designed as highly capable reasoners, with configurable thinking modes.
- **Extended Multimodalities** – Processes Text, Image with variable aspect ratio and resolution support (all models), Video, and Audio (featured natively on the E2B and E4B models).
- **Diverse & Efficient Architectures** – Offers Dense and Mixture-of-Experts (MoE) variants of different sizes for scalable deployment.
- **Optimized for On-Device** – Smaller models are specifically designed for efficient local execution on laptops and mobile devices.
- **Increased Context Window** – The small models feature a 128K context window, while the medium models support 256K.
- **Enhanced Coding & Agentic Capabilities** – Achieves notable improvements in coding benchmarks alongside native function-calling support, powering highly capable autonomous agents.
- **Native System Prompt Support** – Gemma 4 introduces native support for the`system` role, enabling more structured and controllable conversations.

Gemma 4 models are designed to deliver frontier-level performance at each size, targeting deployment scenarios from mobile and edge devices (E2B, E4B) to consumer GPUs and workstations (26B A4B, 31B). They are well-suited for reasoning, agentic workflows, coding, and multimodal understanding.

The models employ a hybrid attention mechanism that interleaves local sliding window attention with full global attention, ensuring the final layer is always global. This hybrid design delivers the processing speed and low memory footprint of a lightweight model without sacrificing the deep awareness required for complex, long-context tasks. To optimize memory for long contexts, global layers feature unified Keys and Values, and apply Proportional RoPE (p-RoPE).

| Property | E2B | E4B | 31B Dense | 
|---|---|---|---|
| **Total Parameters** | 2.3B effective (5.1B with embeddings) | 4.5B effective (8B with embeddings) | 30.7B | 
| **Layers** | 35 | 42 | 60 | 
| **Sliding Window** | 512 tokens | 512 tokens | 1024 tokens | 
| **Context Length** | 128K tokens | 128K tokens | 256K tokens | 
| **Vocabulary Size** | 262K | 262K | 262K | 
| **Supported Modalities** | Text, Image, Audio | Text, Image, Audio | Text, Image | 
| **Vision Encoder Parameters** | *~150M* | *~150M* | *~550M* | 
| **Audio Encoder Parameters** | *~300M* | *~300M* | No Audio | 

The "E" in E2B and E4B stands for "effective" parameters. The smaller models incorporate Per-Layer Embeddings (PLE) to maximize parameter efficiency in on-device deployments. Rather than adding more layers or parameters to the model, PLE gives each decoder layer its own small embedding for every token. These embedding tables are large but are only used for quick lookups, which is why the effective parameter count is much smaller than the total.

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

|  | Gemma 4 31B | Gemma 4 26B A4B | Gemma 4 E4B | Gemma 4 E2B | Gemma 3 27B (no think) | 
|---|---|---|---|---|---|
| MMLU Pro | 85.2% | 82.6% | 69.4% | 60.0% | 67.6% | 
| AIME 2026 no tools | 89.2% | 88.3% | 42.5% | 37.5% | 20.8% | 
| LiveCodeBench v6 | 80.0% | 77.1% | 52.0% | 44.0% | 29.1% | 
| Codeforces ELO | 2150 | 1718 | 940 | 633 | 110 | 
| GPQA Diamond | 84.3% | 82.3% | 58.6% | 43.4% | 42.4% | 
| Tau2 (average over 3) | 76.9% | 68.2% | 42.2% | 24.5% | 16.2% | 
| HLE no tools | 19.5% | 8.7% | - | - | - | 
| HLE with search | 26.5% | 17.2% | - | - | - | 
| BigBench Extra Hard | 74.4% | 64.8% | 33.1% | 21.9% | 19.3% | 
| MMMLU | 88.4% | 86.3% | 76.6% | 67.4% | 70.7% | 
| **Vision** |  |  |  |  |  | 
| MMMU Pro | 76.9% | 73.8% | 52.6% | 44.2% | 49.7% | 
| OmniDocBench 1.5 (average edit distance, lower is better) | 0.131 | 0.149 | 0.181 | 0.290 | 0.365 | 
| MATH-Vision | 85.6% | 82.4% | 59.5% | 52.4% | 46.0% | 
| MedXPertQA MM | 61.3% | 58.1% | 28.7% | 23.5% | - | 
| **Audio** |  |  |  |  |  | 
| CoVoST | - | - | 35.54 | 33.47 | - | 
| FLEURS (lower is better) | - | - | 0.08 | 0.09 | - | 
| **Long Context** |  |  |  |  |  | 
| MRCR v2 8 needle 128k (average) | 66.4% | 44.1% | 25.4% | 19.1% | 13.5% | 

Gemma 4 models handle a broad range of tasks across text, vision, and audio. Key capabilities include:

- **Thinking** – Built-in reasoning mode that lets the model think step-by-step before answering.
- **Long Context** – Context windows of up to 128K tokens (E2B/E4B) and 256K tokens (26B A4B/31B).
- **Image Understanding** – Object detection, Document/PDF parsing, screen and UI understanding, chart comprehension, OCR (including multilingual), handwriting recognition, and pointing. Images can be processed at variable aspect ratios and resolutions.
- **Video Understanding** – Analyze video by processing sequences of frames.
- **Interleaved Multimodal Input** – Freely mix text and images in any order within a single prompt.
- **Function Calling** – Native support for structured tool use, enabling agentic workflows.
- **Coding** – Code generation, completion, and correction.
- **Multilingual** – Out-of-the-box support for 35+ languages, pre-trained on 140+ languages.
- **Audio** (E2B and E4B only) – Automatic speech recognition (ASR) and speech-to-translated-text translation across multiple languages.

You can use all Gemma 4 models with the latest version of Transformers. To get started, install the necessary dependencies in your environment:

`pip install -U transformers torch accelerate`

