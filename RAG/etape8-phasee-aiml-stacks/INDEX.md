# INDEX — Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)

Corpus `etape8-phasee-aiml-stacks` · **13 fichiers** · 881 lignes source · ~5808 mots · partition exacte de `docs/RAG/etape8_phaseE_aiml_stacks.md`.

## Mode d'emploi

1. Filtrer dans `manifest.json` (ou les tableaux ci-dessous) sur `domain`, `task`, `actors`, `dates` ou `keywords`.
2. Ouvrir 1 à 3 fichiers ciblés ; chaque fichier est une unité thématique auto-suffisante avec un en-tête YAML.
3. Pour un événement répété dans plusieurs sections, préférer le fichier marqué `canonical_for` (voir la table Événements canoniques).

## Domaines (dossiers → fichiers)

### `00-aiml-stacks/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)](00-aiml-stacks/overview.md) | 1–47 | deep-dive | reference |
| 02 | [1. PyTorch — the 2.6/2.7/2.8 era](00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md) | 48–109 | deep-dive | reference |
| 03 | [2. TensorFlow, Keras 3, JAX, and MindSpore](00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md) | 110–186 | deep-dive | reference |
| 04 | [3. The Hugging Face ecosystem](00-aiml-stacks/3-the-hugging-face-ecosystem.md) | 187–266 | deep-dive | platform |
| 05 | [4. LLM application frameworks: LangChain, LangGraph, LlamaIndex, Haystack, DSPy](00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md) | 267–355 | deep-dive | reference |
| 06 | [5. Vector databases — developer-angle comparison](00-aiml-stacks/5-vector-databases-developer-angle-comparison.md) | 356–403 | deep-dive | reference |
| 07 | [6. Kernel compilers and model formats: Triton, ONNX, TensorRT, OpenVINO](00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md) | 404–475 | deep-dive | quantization |
| 08 | [7. Quantization tooling: bitsandbytes, GPTQ, AWQ, GGUF](00-aiml-stacks/7-quantization-tooling-bitsandbytes-gptq-awq-gguf.md) | 476–534 | deep-dive | quantization |
| 09 | [8. GPU/accelerator programming interfaces: CUDA Python, Numba, CuPy, ROCm, MLX](00-aiml-stacks/8-gpu-accelerator-programming-interfaces-cuda-python-numba-c.md) | 535–607 | deep-dive | hardware |
| 10 | [9. Distributed training systems: DeepSpeed, FSDP, Megatron-LM, TorchTitan](00-aiml-stacks/9-distributed-training-systems-deepspeed-fsdp-megatron-lm-to.md) | 608–661 | deep-dive | training |
| 11 | [10. Experiment tracking and evaluation: W&B, MLflow, lm-eval, Inspect](00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md) | 662–734 | deep-dive | reference |
| 12 | [11. End-to-end stack decision guide (2026)](00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md) | 735–809 | deep-dive | reference |
| 13 | [13. Glossary](00-aiml-stacks/13-glossary.md) | 810–881 | deep-dive | reference |

## Par tâche

- **hardware** — [8. GPU/accelerator programming interfaces: CUDA Python, Numba, CuPy, ROCm, MLX](00-aiml-stacks/8-gpu-accelerator-programming-interfaces-cuda-python-numba-c.md)
- **platform** — [3. The Hugging Face ecosystem](00-aiml-stacks/3-the-hugging-face-ecosystem.md)
- **quantization** — [6. Kernel compilers and model formats: Triton, ONNX, TensorRT, OpenVINO](00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md), [7. Quantization tooling: bitsandbytes, GPTQ, AWQ, GGUF](00-aiml-stacks/7-quantization-tooling-bitsandbytes-gptq-awq-gguf.md)
- **reference** — [Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)](00-aiml-stacks/overview.md), [1. PyTorch — the 2.6/2.7/2.8 era](00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md), [2. TensorFlow, Keras 3, JAX, and MindSpore](00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md), [4. LLM application frameworks: LangChain, LangGraph, LlamaIndex, Haystack, DSPy](00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md), [5. Vector databases — developer-angle comparison](00-aiml-stacks/5-vector-databases-developer-angle-comparison.md), [10. Experiment tracking and evaluation: W&B, MLflow, lm-eval, Inspect](00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md), [11. End-to-end stack decision guide (2026)](00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md), [13. Glossary](00-aiml-stacks/13-glossary.md)
- **training** — [9. Distributed training systems: DeepSpeed, FSDP, Megatron-LM, TorchTitan](00-aiml-stacks/9-distributed-training-systems-deepspeed-fsdp-megatron-lm-to.md)

## Par acteur

- **AMD** (6) — [00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md](00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md), [00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md](00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md), [00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md](00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md), [00-aiml-stacks/7-quantization-tooling-bitsandbytes-gptq-awq-gguf.md](00-aiml-stacks/7-quantization-tooling-bitsandbytes-gptq-awq-gguf.md), [00-aiml-stacks/8-gpu-accelerator-programming-interfaces-cuda-python-numba-c.md](00-aiml-stacks/8-gpu-accelerator-programming-interfaces-cuda-python-numba-c.md), [00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md](00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md)
- **Alibaba** (1) — [00-aiml-stacks/3-the-hugging-face-ecosystem.md](00-aiml-stacks/3-the-hugging-face-ecosystem.md)
- **Apple** (3) — [00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md](00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md), [00-aiml-stacks/8-gpu-accelerator-programming-interfaces-cuda-python-numba-c.md](00-aiml-stacks/8-gpu-accelerator-programming-interfaces-cuda-python-numba-c.md), [00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md](00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md)
- **China** (1) — [00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md](00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md)
- **CoreWeave** (2) — [00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md](00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md), [00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md](00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md)
- **Google** (3) — [00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md](00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md), [00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md](00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md), [00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md](00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md)
- **Huawei** (1) — [00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md](00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md)
- **Hugging Face** (3) — [00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md](00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md), [00-aiml-stacks/3-the-hugging-face-ecosystem.md](00-aiml-stacks/3-the-hugging-face-ecosystem.md), [00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md](00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md)
- **Intel** (2) — [00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md](00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md), [00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md](00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md)
- **Meta** (1) — [00-aiml-stacks/9-distributed-training-systems-deepspeed-fsdp-megatron-lm-to.md](00-aiml-stacks/9-distributed-training-systems-deepspeed-fsdp-megatron-lm-to.md)
- **Microsoft** (1) — [00-aiml-stacks/9-distributed-training-systems-deepspeed-fsdp-megatron-lm-to.md](00-aiml-stacks/9-distributed-training-systems-deepspeed-fsdp-megatron-lm-to.md)
- **Nvidia** (7) — [00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md](00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md), [00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md](00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md), [00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md](00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md), [00-aiml-stacks/7-quantization-tooling-bitsandbytes-gptq-awq-gguf.md](00-aiml-stacks/7-quantization-tooling-bitsandbytes-gptq-awq-gguf.md), [00-aiml-stacks/8-gpu-accelerator-programming-interfaces-cuda-python-numba-c.md](00-aiml-stacks/8-gpu-accelerator-programming-interfaces-cuda-python-numba-c.md), [00-aiml-stacks/9-distributed-training-systems-deepspeed-fsdp-megatron-lm-to.md](00-aiml-stacks/9-distributed-training-systems-deepspeed-fsdp-megatron-lm-to.md), [00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md](00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md)
- **OpenAI** (3) — [00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md](00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md), [00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md](00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md), [00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md](00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md)
- **SGLang** (1) — [00-aiml-stacks/overview.md](00-aiml-stacks/overview.md)
- **TensorRT-LLM** (3) — [00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md](00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md), [00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md](00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md), [00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md](00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md)
- **vLLM** (3) — [00-aiml-stacks/overview.md](00-aiml-stacks/overview.md), [00-aiml-stacks/7-quantization-tooling-bitsandbytes-gptq-awq-gguf.md](00-aiml-stacks/7-quantization-tooling-bitsandbytes-gptq-awq-gguf.md), [00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md](00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md)

## Par date

- **2025-06** — [00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md](00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md)
- **2025-10** — [00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md](00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md)
- **2025-11** — [00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md](00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md)
- **2025-12** — [00-aiml-stacks/3-the-hugging-face-ecosystem.md](00-aiml-stacks/3-the-hugging-face-ecosystem.md)
- **2025-12-18** — [00-aiml-stacks/3-the-hugging-face-ecosystem.md](00-aiml-stacks/3-the-hugging-face-ecosystem.md), [00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md](00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md)
- **2026-01** — [00-aiml-stacks/3-the-hugging-face-ecosystem.md](00-aiml-stacks/3-the-hugging-face-ecosystem.md)
- **2026-01-26** — [00-aiml-stacks/3-the-hugging-face-ecosystem.md](00-aiml-stacks/3-the-hugging-face-ecosystem.md)
- **2026-02** — [00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md](00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md)
- **2026-03** — [00-aiml-stacks/5-vector-databases-developer-angle-comparison.md](00-aiml-stacks/5-vector-databases-developer-angle-comparison.md)
- **2026-04** — [00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md](00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md)
- **2026-05** — [00-aiml-stacks/3-the-hugging-face-ecosystem.md](00-aiml-stacks/3-the-hugging-face-ecosystem.md)
- **2026-05-08** — [00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md](00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md)
- **2026-05-14** — [00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md](00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md)
- **2026-05-22** — [00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md](00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md)
- **2026-06** — [00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md](00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md), [00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md](00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md)
- **2026-06-22** — [00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md](00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md)
- **2026-06-29** — [00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md](00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md)
- **2026-07** — [00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md](00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md)
- **2026-08-24** — [00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md](00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md)
- **2026-08-29** — [00-aiml-stacks/8-gpu-accelerator-programming-interfaces-cuda-python-numba-c.md](00-aiml-stacks/8-gpu-accelerator-programming-interfaces-cuda-python-numba-c.md), [00-aiml-stacks/13-glossary.md](00-aiml-stacks/13-glossary.md)
- **2026-09** — [00-aiml-stacks/7-quantization-tooling-bitsandbytes-gptq-awq-gguf.md](00-aiml-stacks/7-quantization-tooling-bitsandbytes-gptq-awq-gguf.md)
- **2026-09-18** — [00-aiml-stacks/3-the-hugging-face-ecosystem.md](00-aiml-stacks/3-the-hugging-face-ecosystem.md)
- **2026-09-22** — [00-aiml-stacks/overview.md](00-aiml-stacks/overview.md), [00-aiml-stacks/13-glossary.md](00-aiml-stacks/13-glossary.md)

## Carte de couverture (lignes source)

| plage | fichier |
|---|---|
| 1–47 | etape8-phasee-aiml-stacks/00-aiml-stacks/overview.md |
| 48–109 | etape8-phasee-aiml-stacks/00-aiml-stacks/1-pytorch-the-2-6-2-7-2-8-era.md |
| 110–186 | etape8-phasee-aiml-stacks/00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore.md |
| 187–266 | etape8-phasee-aiml-stacks/00-aiml-stacks/3-the-hugging-face-ecosystem.md |
| 267–355 | etape8-phasee-aiml-stacks/00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-.md |
| 356–403 | etape8-phasee-aiml-stacks/00-aiml-stacks/5-vector-databases-developer-angle-comparison.md |
| 404–475 | etape8-phasee-aiml-stacks/00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op.md |
| 476–534 | etape8-phasee-aiml-stacks/00-aiml-stacks/7-quantization-tooling-bitsandbytes-gptq-awq-gguf.md |
| 535–607 | etape8-phasee-aiml-stacks/00-aiml-stacks/8-gpu-accelerator-programming-interfaces-cuda-python-numba-c.md |
| 608–661 | etape8-phasee-aiml-stacks/00-aiml-stacks/9-distributed-training-systems-deepspeed-fsdp-megatron-lm-to.md |
| 662–734 | etape8-phasee-aiml-stacks/00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins.md |
| 735–809 | etape8-phasee-aiml-stacks/00-aiml-stacks/11-end-to-end-stack-decision-guide-2026.md |
| 810–881 | etape8-phasee-aiml-stacks/00-aiml-stacks/13-glossary.md |

