---
id: etape8-phasee-aiml-stacks/00-aiml-stacks/11-end-to-end-stack-decision-guide-2026
title: "11. End-to-end stack decision guide (2026)"
domain: step-8-phase-e-ai-ml-software-stacks-developer-facing
role: deep-dive
task: reference
actors: ["AMD", "Apple", "CoreWeave", "Google", "Hugging Face", "Nvidia", "TensorRT-LLM"]
dates: ["2025-12-18", "2026-06"]
keywords: ["acquisition", "agents", "amd", "benchmarks", "dpo", "gpu", "nvidia", "pricing", "qlora", "reasoning", "research", "tensorrt"]
source: docs/RAG/etape8_phaseE_aiml_stacks.md
source_anchor: ""
source_lines: [735, 809]
section: "Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)"
sha256: 277a30773d09c75260dc05ba34875469a42123a92563a69332df21731a59f49e
---

# 11. End-to-end stack decision guide (2026)

## 11. End-to-end stack decision guide (2026)

Typical developer-facing stacks observed in the window, composed from the
pieces above:

1. **Open LLM research/fine-tune**: PyTorch 2.8 + Transformers v5 + PEFT/
   QLoRA (bitsandbytes) + Accelerate/DeepSpeed + W&B or MLflow tracking
   [secondary].
2. **Reasoning-model post-training**: same base + TRL (SFT/DPO/RL) +
   lm-eval + Inspect for eval gates [secondary].
3. **RAG product**: LlamaIndex or Haystack ingestion + Qdrant/Weaviate/
   Milvus/pgvector + LangGraph agents + LangSmith/Langfuse observability
   [secondary].
4. **Prompt-optimized LLM program**: DSPy 3.x compiled against the target
   provider via LiteLLM, MLflow hooks for tracking [secondary].
5. **GPU kernel work**: Triton 3.x via `torch.compile`/Inductor; TensorRT-LLM
   or ONNX Runtime 1.30 for serving; TensorRT engines for max NVIDIA
   throughput [secondary].
6. **TPU training**: JAX 0.10.x + Flax/MaxText + Optax/Orbax on NGC or
   Google Cloud TPU (Ironwood/TPU v7 generation in 2026) [secondary].
7. **Apple Silicon local**: MLX 0.32.x + mlx-lm; PyTorch MPS as fallback
   [secondary].
8. **AMD training**: ROCm 7 PyTorch builds; watch Triton CDNA kernel-path
   lag for custom ops [secondary].

---

## 12. Conflicts & Gaps Register

### Conflicts

- **C-01 — Hugging Face Hub totals**: 2.4M+/730K+/1M (A), 1.8M/450K/720K
  (B), 2M+/500K+/1M (C) models/datasets/Spaces. Do not quote any total
  without a dated Hub API snapshot [secondary].
- **C-02 — Transformers v5 release date**: December 18, 2025 vs January
  26, 2026 across community guides. The major-version event is certain;
  the exact landing date is not [secondary].
- **C-03 — DSPy GitHub stars**: 22K+ vs 32.4K vs 33.7K across secondary
  sources. Popularity claims need a dated repo snapshot [secondary].
- **C-04 — W&B acquisition price**: ~$1.4B vs ~$1.7B reported. The
  CoreWeave acquisition itself is consistently reported; the price is
  not [secondary].

### Gaps (no verified 2026 data captured in this pass)

- **G-01**: Official PyTorch 2.6/2.7/2.8 release notes (features and the
  4,164-commit/585-contributor figures are mirror-reported).
- **G-02**: MindSpore 2026 release version.
- **G-03**: Diffusers 2026 version.
- **G-04**: CVE-2026-4539 (LangChain) details.
- **G-05**: Haystack 2.x exact 2026 version.
- **G-06**: Weaviate 2026 version.
- **G-07**: Milvus 2026 version.
- **G-08**: TensorRT / TensorRT-LLM 2026 versions.
- **G-09**: OpenVINO 2026 feature-level release notes (version 2026.2
  confirmed via Keras README).
- **G-10**: AutoGPTQ maintenance status vs GPTQModel succession (community
  claim, unverified against repos).
- **G-11**: cuda-python 2026 version.
- **G-12**: Numba 2026 version.
- **G-13**: CuPy 2026 version.
- **G-14**: Current (2026) DeepSpeed benchmarks; old "10×" claims are
  era-labeled ZeRO-2 material.
- **G-15**: Megatron-LM 2026 version.
- **G-16**: TorchTitan 2026 release state.
- **G-17**: W&B 2026 platform version/pricing tiers.
- **G-18**: Inspect AI 2026 version.
- **G-19**: Keras 3.15.1 date/version beyond the Wikipedia snapshot;
  cross-check PyPI before quoting.
- **G-20**: LangChain 1.3.8 / langchain-core 1.4.6 (June 2026 claim,
  unverified); LlamaIndex 0.14.22–0.14.24 series (secondary reports,
  unverified against PyPI).

---

