# NOTES — corpus `etape5-trackc-huawei-intel`

## Génération

Mode auto (`RAG/_tools/build_rag.py`) avec `first_is_content: true` (H1 unique = contenu),
cibles **45–90 lignes**. Dossier `00-huawei-intel` (`folder_name`), fichiers par titre H2.

Source sans ancre HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

« Step 5 — Track C : Huawei, matériel chinois, CPUs Intel » (février → septembre 2026).

- Huawei ;
- puces IA chinoises : Cambricon, Moore Threads, Biren, Metax ;
- Intel (CPUs).

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
