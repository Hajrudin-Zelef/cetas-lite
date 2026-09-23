---
id: etape8-phasee-aiml-stacks/00-aiml-stacks/2-tensorflow-keras-3-jax-and-mindspore
title: "2. TensorFlow, Keras 3, JAX, and MindSpore"
domain: step-8-phase-e-ai-ml-software-stacks-developer-facing
role: deep-dive
task: reference
actors: ["AMD", "China", "Google", "Huawei", "Nvidia", "TensorRT-LLM"]
dates: ["2026-07"]
keywords: ["ascend", "attention", "gpu", "inference", "nvfp4", "nvidia", "research", "tensorrt", "tpu", "training"]
source: docs/RAG/etape8_phaseE_aiml_stacks.md
source_anchor: ""
source_lines: [110, 186]
section: "Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)"
sha256: 87c32f80a77ee564d8ce4d9a90318e07e1b062f50b54ee0e25ea7bc7153628bd
---

# 2. TensorFlow, Keras 3, JAX, and MindSpore

## 2. TensorFlow, Keras 3, JAX, and MindSpore

### 2.1 TensorFlow 2.20 / 2.21

- The official TensorFlow GitHub release page lists 2.20.0 with: an optional
  `tensorflow-io-gcs-filesystem` package, a `tf.data` minimum-parallelism
  tuning knob (`autotune.min_parallelism`), and TensorFlow Lite deprecation
  in favor of migration to LiteRT [official].
- A separate repository reference reports 2.21.0, but it was not verified
  against the official release page; treat 2.21.0 as [unverified].
- `tf.lite` deprecation toward LiteRT (the `google-ai-edge` / LiteRT project)
  continues the multi-year migration of the mobile/edge story out of the
  core TensorFlow package [official for 2.20 direction; secondary for naming].
- TensorFlow remains strong in production serving (TF Serving, TFLite/LiteRT
  edge) while ceding most new LLM research to PyTorch [secondary].

### 2.2 Keras 3 — the multi-backend front end

- Keras 3 is a full rewrite that runs on TensorFlow, JAX, PyTorch, and
  (inference-only) OpenVINO backends, selected via the `KERAS_BACKEND`
  environment variable or `~/.keras/keras.json` [official].
- Wikipedia records the stable release as **3.15.1 on 29 July 2026**
  [secondary].
- The Keras README backend-compatibility table lists minimum backend
  versions for the latest 3.x: TensorFlow 2.16.1, JAX 0.4.20, PyTorch 2.1.0,
  and OpenVINO 2026.2.0 [official].
- The OpenVINO backend is inference-only (`model.predict()`), designed for
  running predictions rather than training [official].
- Keras 3 is the default Keras for TensorFlow ≥2.16 (where `tf.keras`
  resolves to Keras 3); Keras 2 remains available via the `tf_keras`
  maintenance package [secondary].
- Keras 3 models can be instantiated as PyTorch modules, exported as
  TensorFlow SavedModels, or expressed as stateless JAX functions, which is
  the framework's portability pitch [secondary].

### 2.3 JAX and the XLA ecosystem

- JAX is Google's NumPy-on-XLA library pairing composable function
  transformations (`grad`, `jit`, `vmap`, `pmap`) with XLA compilation; it is
  the primary framework for TPU training at Google and in the open-source
  MaxText/Flax ecosystem [secondary].
- NVIDIA's JAX-Toolbox staging notes show the 2026 cadence: NGC 26.04 ships
  JAX 0.9.2, 26.05 ships 0.10.0, 26.06 ships 0.10.1, and 26.07 ships **0.10.2**
  with CUDA 13.3.x, cuDNN-backed flex attention, AoT compilation support,
  `nvfp4` kernel codegen, and multi-stream collectives [secondary].
- A community issue report confirms the existence of JAX 0.8.2 alongside
  0.7.2 (both referenced in a May-2026 regression report) [secondary].
- Flax 0.11.2 appears as the companion neural-network library in the
  NVIDIA JAX container (25.10 release notes) [secondary].
- Key distributed primitives: `jax.pjit` / `jax.jit` with explicit sharding
  via `jax.sharding.PartitionSpec` and `Mesh`, plus `jax.distributed` for
  multi-host coordination [secondary].
- Optax (optimizers), Orbax (checkpointing), and Chex (testing utilities)
  round out the standard JAX training stack [secondary].

### 2.4 MindSpore

- MindSpore is Huawei's open-source deep-learning framework, positioned for
  Ascend NPU hardware alongside CPU/GPU support [secondary].
- No 2026 release version could be verified from an official source during
  this research pass; the framework is noted here for completeness and
  marked as a gap (G-02) [unverified].
- Its practical relevance is concentrated in China-market and Huawei-cloud
  workloads rather than the global open-source LLM mainstream [secondary].

### 2.5 Framework positioning matrix (2026)

| Framework | Primary hardware | LLM research share | Production serving | Notes |
|---|---|---|---|---|
| PyTorch 2.x | NVIDIA GPU (CUDA), ROCm, MPS | Dominant | TorchServe, ONNX/TensorRT export | Default for HF/DeepSpeed/TorchTitan [secondary] |
| JAX 0.10.x | TPU, NVIDIA GPU | Strong (Google, open TPU) | XLA AOT, TF Serving | Best TPU story; `pjit` sharding [secondary] |
| TensorFlow 2.20 | TPU, GPU, edge | Legacy/maintenance | TF Serving, LiteRT | `tf.lite` deprecated → LiteRT [official] |
| Keras 3.15 | Backend-agnostic | Growing (portability) | Per-backend export | Multi-backend API, OpenVINO inference [official] |
| MindSpore | Ascend NPU | Niche (regional) | Huawei cloud | Version unverified (G-02) [unverified] |

---

