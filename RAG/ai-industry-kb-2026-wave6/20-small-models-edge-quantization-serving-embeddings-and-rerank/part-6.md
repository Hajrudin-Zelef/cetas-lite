---
id: ai-industry-kb-2026-wave6/20-small-models-edge-quantization-serving-embeddings-and-rerank/part-6
title: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers (part 6)"
domain: small-models-edge-quantization-serving-embeddings-and-rerank
role: deep-dive
task: quantization
actors: ["Alibaba", "Nvidia", "SGLang", "Z.ai", "vLLM"]
dates: []
keywords: ["quantization", "apache", "awq", "blackwell", "consumer", "copyright", "cost", "gguf", "glm", "gptq", "gpu", "inference"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9757, 9764]
section: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers"
delta_of: ai-industry-kb-2026
sha256: d7d94f4b69286b29f1e639bc6ce8be274131691a1b9c7b4d043f6b4c10ef4f87
---

# §20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers (part 6)

- Hardware portability: GGUF runs broadly on CPU/GPU via llama.cpp; AWQ/GPTQ need CUDA kernels; NVFP4 W4A4 requires Blackwell-native paths (W4A16 Marlin support is broader); EXL3/QTIP targets ~4 bpw on consumer NVIDIA via ExLlamaV3 [COMMUNITY]. Sources: https://github.com/allanschramm/local-model-autotuning/blob/HEAD/docs/discovery/inference-engines-landscape.md and https://github.com/kekzl/imp/blob/HEAD/docs/quantization.md
- No universal quality-loss percentages exist: published degradation figures are directional and implementation/model-dependent — never quote a fixed "% quality loss" for a quantization format [DIRECTIONAL]. Source: https://github.com/xuehaosun/lb_eval/blob/HEAD/openclaw_config/workspace/skills/auto_quant/SKILL.md
- Z.ai's GLM-5.3-Flash serving stack: custom SGLang-based engine separating encoding, prefill, and decoding, claiming 3x end-to-end serving improvement across tens of thousands of domestic accelerators [VENDOR]. Source: https://medium.com/@noahkenji283/glm-5-3-flash-looks-like-a-giant-model-release-but-the-real-story-is-cost-context-and-deployment-f38d62113e4f
- GLM-5.3-Flash local serving support: SGLang, vLLM, TokenSpeed, KTransformers [VENDOR]. Source: https://medium.com/@noahkenji283/glm-5-3-flash-looks-like-a-giant-model-release-but-the-real-story-is-cost-context-and-deployment-f38d62113e4f
- The Ollama registry blob for qwen3:0.6b-q8_0 carries the full Apache 2.0 text with "Copyright 2024 Alibaba Cloud" — the license text travels with the weights even in registry-quantized form [SECONDARY]. Source: https://registry.ollama.com/library/qwen3:0.6b-q8_0/blobs/d18a5cc71b84
- A community license-audit file records: NVIDIA nemotron-speech-streaming-en-0.6b under NVIDIA Open Model License (NOTICE: "Licensed by NVIDIA Corporation under the NVIDIA Open Model License"); Voxtral-Mini-4B-Realtime-2602-ONNX under Apache 2.0; FunAudioLLM/SenseVoiceSmall under FunASR Model Open Source License v1.1 (ambiguous commercial scope — pending legal review) [COMMUNITY]. Source: https://github.com/mikeschirtzinger/silent-notetaker/blob/HEAD/docs/research/model-licenses.md


