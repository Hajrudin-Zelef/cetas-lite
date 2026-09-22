# NOTES — corpus `etape4-trackc-cuda-rocm-pytorch`

## Génération

Mode auto (`RAG/_tools/build_rag.py`) avec `first_is_content: true` (H1 unique = contenu),
cibles **45–90 lignes**. Dossier `00-cuda-rocm-pytorch` (`folder_name`), fichiers par titre H2.

Source sans ancre HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

« Step 4 — Track C : CUDA / ROCm / PyTorch — pile de calcul GPU » (2026).

- CUDA ;
- AMD ROCm ;
- PyTorch 2.x en 2026.

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
