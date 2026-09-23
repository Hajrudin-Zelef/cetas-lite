---
id: etape8-phasee-aiml-stacks/00-aiml-stacks/13-glossary
title: "13. Glossary"
domain: step-8-phase-e-ai-ml-software-stacks-developer-facing
role: deep-dive
task: reference
actors: []
dates: ["2026-08-29", "2026-09-22"]
keywords: ["agent", "agents", "awq", "benchmark", "fine-tuning", "gguf", "gptq", "gpu", "ipo", "llama", "lora", "mcp"]
source: docs/RAG/etape8_phaseE_aiml_stacks.md
source_anchor: ""
source_lines: [810, 881]
section: "Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)"
sha256: f874e0531fbfa994eb276459a5c6cd43dff8ff539bf3d195fea0a5fd5d29eb14
---

# 13. Glossary

## 13. Glossary

- **DDP** — DistributedDataParallel: replicated-model data parallelism.
- **FSDP** — Fully Sharded Data Parallel: ZeRO-style sharding in PyTorch.
- **ZeRO** — Zero Redundancy Optimizer: DeepSpeed's sharding stages.
- **DTensor** — PyTorch's distributed tensor abstraction.
- **DCP** — Distributed Checkpoint: reshardable checkpoint I/O.
- **QLoRA** — 4-bit quantized base + LoRA adapters (bitsandbytes + PEFT).
- **LoRA** — Low-Rank Adaptation: parameter-efficient fine-tuning.
- **TRL** — Transformer Reinforcement Learning (HF post-training lib).
- **PEFT** — Parameter-Efficient Fine-Tuning (HF adapter lib).
- **GEPA** — reflective prompt-evolution optimizer absorbed into DSPy 3.2.
- **HNSW** — Hierarchical Navigable Small World: ANN index algorithm.
- **ANN** — Approximate Nearest Neighbor search.
- **IR** — Intermediate Representation (OpenVINO model format).
- **MoE** — Mixture of Experts.
- **MFU** — Model FLOPs Utilization.
- **AOT** — Ahead-of-time compilation.

---

## 14. Source index (verbatim URLs, access 2026-09-22)

- `https://github.com/tunguz/pytorch_hardware_may_2026/blob/HEAD/reference_audit_may_2026.md`
- `https://ngontinh24.com/article/pytorch-2-8-release-blog-pytorch`
- `https://github.com/tensorflow/tensorflow/releases`
- `https://github.com/keras-team/keras/blob/HEAD/README.md`
- `https://pypi.org/project/keras-nightly/3.16.0.dev2026090305/`
- `http://en.wikipedia.org/wiki/Keras`
- `https://kanerika.com/blogs/keras-vs-tensorflow/`
- `https://github.com/nvidia/jax-toolbox/blob/HEAD/README.md`
- `https://github.com/nvidia/jax-toolbox/blob/HEAD/docs/reference/staging-containers.md`
- `https://github.com/jax-ml/jax/issues/38145`
- `https://siteproxy.ruqli.workers.dev:443/https/docs.nvidia.com/deeplearning/frameworks/pdf/JAX-Release-Notes.pdf`
- `https://github.com/tenstorrent/tt-vscode-toolkit/blob/HEAD/content/lessons/tt-xla-jax.md`
- `https://medium.com/@vici0549/migrating-to-transformers-5-guide-1e90058a7633`
- `https://pub.towardsai.net/transformers-v5-hugging-faces-next-big-leap-in-simple-and-powerful-ai-models-0b8b4dbc2216?gi=c435d15ac7a2`
- `https://github.com/zucchini-nlp/transformers/blob/HEAD/MIGRATION_GUIDE_V5.md`
- `https://github.com/advisories/GHSA-69w3-r845-3855`
- `https://www.digitalapplied.com/blog/most-downloaded-open-ai-models-hugging-face-september-2026`
- `https://memeburn.com/open-weight-ai-model-statistics-2026/`
- `https://www.metacto.com/blogs/what-is-hugging-face-a-guide-to-the-ai-community-and-its-tools`
- `https://aifoss.dev/blog/langchain-vs-llamaindex-vs-haystack-2026/`
- `https://media.patentllm.org/news/dev-stack/llamaindex-v0-14-24-langchain-fireworks-v1-6-0-github-copilo-20260824`
- `https://www.piwheels.org/project/llama-index-llms-google-genai/`
- `https://github.com/stephendenisedwards/micro-x-agent-loop-python/blob/HEAD/documentation/docs/research/dspy-architecture.md`
- `https://github.com/intertwine/dspy-agent-skills/blob/HEAD/articles/01-why-dspy-agent-skills.md`
- `https://medium.com/@rahulponnusamy/forget-prompt-engineering-why-dspy-is-the-future-of-ai-programming-04bdd4909e4d`
- `https://medium.com/actian-for-developers/vector-database-news-2026-5d707384be6b`
- `https://www.techtarget.com/data-technologies/news/366649627/Qdrant-builds-dataset-to-benchmark-vector-retrieval-at-scale`
- `https://www.spheron.network/blog/openai-triton-kernel-gpu-cloud-2026/`
- `https://pypi.org/project/onnxruntime/1.30.0/`
- `https://www.marktechpost.com/2026/09/18/gguf-vs-gptq-vs-awq-vs-exl2-llm-model-formats-explained-2026/`
- `https://www.premai.io/blog/llm-quantization-guide-gguf-vs-awq-vs-gptq-vs-bitsandbytes-compared-2026/`
- `https://github.com/furyhawk/bitsandbytes`
- `https://deepspeed.readthedocs.io/_/downloads/en/latest/pdf/`
- `https://www.microsoft.com/en-us/research/?p=658659`
- `https://github.com/solaius/ai-asset-registry/blob/HEAD/agents/agent-registry/research/05-mlflow-upstream.md`
- `https://github.com/kkruglik/mlflow-mcp/blob/HEAD/CHANGELOG.md`
- `https://github.com/eleutherai/lm-evaluation-harness/blob/HEAD/README.md`
- `https://github.com/EleutherAI/lm-evaluation-harness/releases/tag/v0.4.11`
- `https://github.com/EleutherAI/lm-evaluation-harness/releases/tag/v0.4.9.2`
- `https://read.getsuperintel.com/p/exclusive-interview-with-phil-gurbacki-vp-of-product-weights-biases-at-coreweave`
- `https://siliconangle.com/2025/03/04/ipo-bound-coreweave-buy-ai-developer-weights-biases-reported-1-4b/`
- `https://pitchbook.com/news/articles/coreweave-acquires-ai-developer-platform-weights-biases`
- `https://github.com/defai-digital/ax-engine/blob/HEAD/docs/performance/mlx-0.32.2-admission-2026-08-29.md`
- `https://www.macrumors.com/2026/06/09/apple-outlines-major-ai-and-developer-tool-updates/`

---

*End of Phase E file. Cutoff 2026-09-22. Sole writer; no other project file
modified.*
