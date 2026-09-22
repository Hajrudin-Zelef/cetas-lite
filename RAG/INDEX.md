# INDEX — RAG

Corpus RAG de référence pour Cetas. Chaque corpus est une partition exacte de sa source, avec index et manifest.

Un corpus `delta` ne contient que les faits nouveaux/corrigés d'une source déjà couverte par un corpus `base` ; il ne la remplace pas. `delta_of` indique la base visée.

| corpus | titre | relation | fichiers | source | index |
|---|---|---|---|---|---|
| `briefing-ia-2026` | AI News 2026 — Reference Dossier |  | 126 | `docs/RAG/briefing-ia-2026-en.md` | [INDEX](briefing-ia-2026/INDEX.md) |
| `briefing-general-tech-2026` | General Tech News 2026 — Hardware, Infrastructure & Consumer Tech |  | 112 | `docs/RAG/briefing-general-tech-2026-en.md` | [INDEX](briefing-general-tech-2026/INDEX.md) |
| `ai-industry-kb-2026` | AI Industry Knowledge Base 2026 | base | 123 | `docs/RAG/ai-industry-knowledge-base-2026.md` | [INDEX](ai-industry-kb-2026/INDEX.md) |
| `ai-industry-kb-2026-wave6` | AI Industry Knowledge Base 2026 — Wave 6 Consolidation | delta de `ai-industry-kb-2026` | 136 | `docs/RAG/ai-industry-knowledge-base-2026-wave6.md` | [INDEX](ai-industry-kb-2026-wave6/INDEX.md) |
| `frontier-models-2026` | Frontier AI Models 2026 — Vague 1 (EN) |  | 12 | `docs/RAG/Grands titres IA modèlesEN.md` | [INDEX](frontier-models-2026/INDEX.md) |
| `labs-grok-platforms-2026` | Labs, Grok, Tools & Platforms 2026 — Vague 2 (EN) |  | 17 | `docs/RAG/Labos, Grok, outils & plateformes_EN.md` | [INDEX](labs-grok-platforms-2026/INDEX.md) |
| `open-local-models-2026` | Open / Local AI Models 2026 (EN) |  | 26 | `docs/RAG/Modèles IA open  locauxEN.md` | [INDEX](open-local-models-2026/INDEX.md) |
| `tools-platforms-2026` | AI Tools & Platforms 2026 (Step 2) |  | 8 | `docs/RAG/Outils & plateformes IAEN.md` | [INDEX](tools-platforms-2026/INDEX.md) |
| `etape4-trackd-unsloth-training` | Step 4 — Track D : Unsloth & outillage d'entraînement/fine-tuning 2026 |  | 3 | `docs/RAG/etape4_trackD_unsloth_training.md` | [INDEX](etape4-trackd-unsloth-training/INDEX.md) |
| `etape5-tracka-nvidia` | Step 5 — Track A : Nvidia (2026) |  | 4 | `docs/RAG/etape5_trackA_nvidia.md` | [INDEX](etape5-tracka-nvidia/INDEX.md) |
| `etape5-trackc-huawei-intel` | Step 5 — Track C : Huawei & Intel (2026) |  | 3 | `docs/RAG/etape5_trackC_huawei_intel.md` | [INDEX](etape5-trackc-huawei-intel/INDEX.md) |
| `etape4-trackc-cuda-rocm-pytorch` | Step 4 — Track C : CUDA / ROCm / PyTorch (2026) |  | 3 | `docs/RAG/etape4_trackC_cuda_rocm_pytorch.md` | [INDEX](etape4-trackc-cuda-rocm-pytorch/INDEX.md) |
| `etape5-trackb-amd` | Step 5 — Track B : AMD (2026) |  | 4 | `docs/RAG/etape5_trackB_amd.md` | [INDEX](etape5-trackb-amd/INDEX.md) |
| `etape4-tracka-vllm-sglang` | Step 4 — Track A : vLLM + SGLang (2026) |  | 20 | `docs/RAG/etape4_trackA_vllm_sglang.md` | [INDEX](etape4-tracka-vllm-sglang/INDEX.md) |
| `etape4-trackb-local-inference` | Step 4 — Track B : pile d'inférence locale (llama.cpp, Ollama, LM Studio) |  | 17 | `docs/RAG/etape4_trackB_local_inference.md` | [INDEX](etape4-trackb-local-inference/INDEX.md) |
| `etape5-trackd-servers` | Step 5 — Track D : serveurs IA, marché et réseau datacenter (2026) |  | 24 | `docs/RAG/etape5_trackD_servers.md` | [INDEX](etape5-trackd-servers/INDEX.md) |
| `labs-hyperscalers-2026` | Step 3 — Labs & Hyperscalers (2026) |  | 47 | `docs/RAG/Labos  hyperscalersEN.md` | [INDEX](labs-hyperscalers-2026/INDEX.md) |

Voir aussi [README](README.md) et `manifest.json`.
