# NOTES — corpus `collect-250926-servers-hardware`

## Génération

Mode **files** (`RAG/_tools/build_rag.py`) : **une source = un chunk**, copie verbatim.

- source : `docs/RAG/clean4` (196 fichiers, extraits de `RAGclean4.tar`)
- dossier interne : `servers-hardware` · domaine unique : `servers-hardware`
- total : ~195 257 mots · 19 327 lignes · 1,7 Mo

## Contenu

Collecte 4 (25/09/2026), quatre familles de sources :

- **Serveurs & centre de calcul** (~89 fiches) : Supermicro, Dell PowerEdge, HPE ProLiant,
  ASUS, Gigabyte, Mitac, Pegatron, Compal, Lenovo, Hyve/Ampere, IBM Power11, CXL,
  NVLink Fusion, GB200/GB300 NVL72, Instinct MI325X, Gaudi 2, El Capitan, xAI Colossus.
- **Docs Ollama** (35 fiches) : pages d'usage (modelfile, tool-calling, embeddings,
  cloud, MLX, blog) — le cœur de l'inférence locale OpenAI-compatible.
- **Docs Unsloth** (31 fiches) : fine-tuning, QAT, LoRA, GGUFs dynamiques, vLLM/SGLang,
  guides macOS/Linux/Windows/Docker.
- **Docs opencode & agents** (31 fiches) : agents, skills, plugins, outils, permissions,
  policies, TUI, websearch, context/compaction, go/zen.

Quelques fiches nommées par hash (`34c989ee`, `960535b9`, `eb1c9b1e`) correspondent à des
fragments de documentation sans H1 (tableaux de dépannage, pages d'installation) : titre
YAML = nom de fichier, texte non réécrit.

## Titres

Les noms avec espaces et majuscules (docs Unsloth/Ollama) sont conservés tels quels dans
l'index ; le fichier de sortie est slugifié (`3x-faster-llm-training-...`). Aucune
réécriture du texte source.

## Provenance

`source:` dans l'en-tête YAML = chemin dans `docs/RAG/clean4` (espace perso, gitignoré).
Cette ligne est une **référence morte** à l'exécution : le runtime ne la lit jamais
(strip du front-matter au chargement), seule `verify_rag.py` la relit à la génération.

## Non audité

Aucun défaut de source recensé ; les fichiers restés en FR sont laissés verbatim.

## Vérification

`verify_rag.py` : `ALL CHECKS PASSED` — correspondance 1:1 fichier ↔ chunk,
sha256, corps verbatim.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
