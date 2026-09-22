# NOTES — corpus `open-local-models-2026`

## Génération

Produit en **mode auto** (`RAG/_tools/build_rag.py`, clé `mode: auto`) :

- partition par titres **H1 → H2 → H3**, sans jamais couper au milieu d'un paragraphe ;
- cibles de taille : 50–110 lignes par chunk ;
- un dossier par document H1 ; un fichier par bloc (nommé d'après son premier titre,
  `overview` pour l'entête de section).

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

Trois documents concaténés (série « ÉTAPE 1 — Open / Local AI Models ») :

1. **Piste chinoise** : Qwen, DeepSeek, Kimi, GLM, MiMo, MiniMax, comparaison
   inter-familles, sources clés ;
2. **Piste occidentale** : Meta Llama / Muse Spark, Poolside, Mistral, Nvidia Nemotron,
   Microsoft Phi, Google Gemma, autres publications open-weight, tableau maître ;
3. **Inférence locale (pratique)** : moteurs (Ollama, llama.cpp, LM Studio), besoins
   matériels et calcul de VRAM, quantification, vLLM vs SGLang, coût auto-hébergement vs
   API, données OpenRouter, notes de méthode.

La source contient **2 blocs de code** (fences ```` ``` ````) conservés verbatim.

## Non audité

Contenu non relu section par section : aucun défaut de source recensé. Fidélité garantie
par le vérificateur (concaténation == source).

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
