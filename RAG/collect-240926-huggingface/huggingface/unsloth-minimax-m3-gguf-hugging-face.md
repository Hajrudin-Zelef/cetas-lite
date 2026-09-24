---
id: collect-240926-huggingface/huggingface/unsloth-minimax-m3-gguf-hugging-face
title: "Read our How to Run MiniMax M3 Guide!"
domain: huggingface
role: reference
task: reference
actors: ["MiniMax", "Unsloth"]
dates: []
keywords: ["agentic", "attention", "benchmarks", "compute", "decode", "gguf", "gpus", "inference", "latency", "license", "llama", "llama.cpp"]
source: docs/RAG/clean_en/huggingface/unsloth-minimax-m3-gguf-hugging-face.md
source_anchor: ""
source_lines: [1, 66]
sha256: de8bee66d88bddf85508f8960fae750c797dff9379583c76e840da7acf45d590
---

# Read our How to Run MiniMax M3 Guide!

<!-- source: https://huggingface.co/unsloth/MiniMax-M3-GGUF -->

# Read our How to Run MiniMax M3 Guide!

    *See Unsloth Dynamic 2.0 GGUFs for our quantization benchmarks.*
  

- EXPERIMENTAL GGUF / support for MiniMax-M3
- **Jun 12 Update:** You can now run MiniMax M3 in Unsloth Studio. See our Guide.
- Example of MiniMax M3 (5-bit GGUF) running in Unsloth Studio:

**EXPERIMENTAL GGUF / support for MiniMax-M3 in llama.cpp:**

MiniMax-M3 support in llama.cpp is preliminary and not yet in a released build. To run these GGUFs, build llama.cpp from PR #24523:

```
git clone https://github.com/ggml-org/llama.cpp
cd llama.cpp
git fetch origin pull/24523/head:minimax-m3
git checkout minimax-m3
cmake -B build -DGGML_CUDA=ON
cmake --build build --config Release -j --target llama-cli llama-server
```
Then run a quant. The model is large (~428B params), so offload across GPUs with `-ngl 99` or keep the weights in CPU RAM:

```
./build/bin/llama-cli -hf unsloth/MiniMax-M3-GGUF:UD-IQ1_M
```
Note: MiniMax Sparse Attention is not supported yet, so inference falls back to dense attention.

**Highlights:**

- **Native Multimodality:** M3 undergoes mixed-modality training from the very first step, enabling deeper semantic fusion across text, image, and video.
- **Context Scaling via Sparse Attention:** M3 introduces MiniMax Sparse Attention (MSA) to improve long context efficiency. M3 delivers 9× prefill and 15× decode speedups compared to M2 at 1M context, reducing per-token compute to 1/20.
- **Coding & Cowork Capability:** M3 achieves frontier-level performance across long-horizon agentic benchmarks, excelling in both coding and cowork.

| Architecture | MoE + MSA (MiniMax Sparse Attention) | 
| Total Parameters | ~428B | 
| Activated Parameters | ~23B | 
| Experts | 128 (4 active per token) | 
| Layers | 60 | 
| Context Length | 1M tokens | 
| Modalities | Text, Image, Video | 
| Precision | bfloat16 | 
| Transformers | ≥ 4.52.4 ( `trust_remote_code=True` ) | 
| License | MiniMax Community License | 

M3 supports two reasoning modes:

- **thinking** — for complex reasoning, agentic tasks, and long-horizon collaboration.
- **non-thinking** — for latency-sensitive scenarios such as chat and code completion.

Download the model:

```
hf download MiniMaxAI/MiniMax-M3 --local-dir MiniMax-M3
```
You can also get model weights from ModelScope.

We recommend the following parameters for best performance: `temperature=1.0`, `top_p=0.95`, `top_k=40`. Default system prompt:

```
You are a helpful assistant. Your name is MiniMax-M3 and was built by MiniMax.
```
- Downloads last month
- 6,273
