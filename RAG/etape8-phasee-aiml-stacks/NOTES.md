# NOTES — corpus `etape8-phasee-aiml-stacks`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-aiml-stacks` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)` — 881 lignes, 13 chunks.

- Scope
- Methodology and provenance
- Provenance legend
- 1. PyTorch — the 2.6/2.7/2.8 era
- 2. TensorFlow, Keras 3, JAX, and MindSpore
- 3. The Hugging Face ecosystem
- 4. LLM application frameworks: LangChain, LangGraph, LlamaIndex, Haystack, DSPy
- 5. Vector databases — developer-angle comparison
- 6. Kernel compilers and model formats: Triton, ONNX, TensorRT, OpenVINO
- 7. Quantization tooling: bitsandbytes, GPTQ, AWQ, GGUF
- 8. GPU/accelerator programming interfaces: CUDA Python, Numba, CuPy, ROCm, MLX
- 9. Distributed training systems: DeepSpeed, FSDP, Megatron-LM, TorchTitan
- 10. Experiment tracking and evaluation: W&B, MLflow, lm-eval, Inspect
- 11. End-to-end stack decision guide (2026)
- 12. Conflicts & Gaps Register
- 13. Glossary
- 14. Source index (verbatim URLs, access 2026-09-22)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
