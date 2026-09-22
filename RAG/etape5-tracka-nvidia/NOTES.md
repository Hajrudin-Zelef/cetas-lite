# NOTES — corpus `etape5-tracka-nvidia`

## Génération

Mode auto (`RAG/_tools/build_rag.py`) avec `first_is_content: true` (H1 unique = contenu),
cibles **45–90 lignes**. Dossier `00-nvidia` (`folder_name`), fichiers par titre H2.

Source sans ancre HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

« Step 5 — Track A : Nvidia » (2026).

- GPUs : B200 / B300, plateforme Vera Rubin ;
- H100 / H200 : disponibilité et tendances de prix 2026 ;
- systèmes DGX ;
- réseau : NICs, DPUs, switches, NVLink.

Chiffres souvent **annoncés par le constructeur** ; les tags de provenance du texte sont
conservés verbatim.

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
