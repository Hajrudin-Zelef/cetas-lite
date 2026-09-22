# NOTES — corpus `etape4-trackb-local-inference`

## Génération

Mode auto (`RAG/_tools/build_rag.py`) avec `first_is_content: true` (H1 unique = contenu),
cibles **45–90 lignes**. Dossier `00-local-inference` (`folder_name`), fichiers par titre H2.

Source sans ancre HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

« Step 4 — Track B : pile d'inférence locale » (2026).

- **llama.cpp** : train de versions, format GGUF, backends, performances (flash attention,
  décodage spéculatif), fork `ik_llama.cpp`, autres forks, mode serveur (`llama-server`),
  outillage de quantification, incertitudes, sources ;
- **Ollama** : société/financement, timeline 2026, bibliothèque de modèles, API, offres
  Cloud/Turbo et historique de prix, fonctionnalités entreprise, adoption, app desktop,
  quantification, backends matériels, positionnement concurrentiel ;
- **LM Studio** + paysage matériel d'inférence locale + benchmarks de quantification GGUF,
  appendices et journal de vérification.

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
