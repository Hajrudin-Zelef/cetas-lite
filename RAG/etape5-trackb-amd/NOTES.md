# NOTES — corpus `etape5-trackb-amd`

## Génération

Mode auto (`RAG/_tools/build_rag.py`) avec `first_is_content: true` (H1 unique = contenu),
cibles **45–90 lignes**. Dossier `00-amd` (`folder_name`), fichiers par titre H2.

Source sans ancre HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

« Step 5 — Track B : AMD » (2026).

- Instinct MI350X / MI355X (CDNA 4) : specs, disponibilité, benchmarks ;
- roadmap MI400 : statut, specs annoncées, calendrier ;
- EPYC Venice 9006 (6e gén, Zen 6), annonce 2026.

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
