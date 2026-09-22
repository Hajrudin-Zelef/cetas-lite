# NOTES — corpus `labs-grok-platforms-2026`

## Génération

Produit en **mode auto** (`RAG/_tools/build_rag.py`, clé `mode: auto`) :

- partition par titres **H1 → H2 → H3**, sans jamais couper au milieu d'un paragraphe ;
- cibles de taille : 50–110 lignes par chunk ;
- un dossier par document H1 ; un fichier par bloc (nommé d'après son premier titre,
  `overview` pour l'entête de section).

La source ne contient **aucune ancre** HTML : pas de table ancre→fichier ; navigation via
`INDEX.md` et `manifest.json`.

## Contenu

Fiches de recherche concaténées (série « VOLET 1 — Vague 2 ») :

- Anthropic : Claude Opus 5.5, Opus 5, rumeur Opus 5.2, Fable 5 / 5.1, Claude Code ;
- xAI : Grok 4.7, rachat de xAI par SpaceX, intégration au Pentagone, datacenters orbitaux ;
- OpenAI : progression GPT-5.x, GPT-6 « Astra », procès antitrust de septembre 2026 ;
- Google / Gemini 3.x (dont incident de sécurité de mai 2026) ;
- plateformes : OpenRouter, Hugging Face, vLLM, SGLang.

Chiffres de benchmark **annoncés par les éditeurs**, sauf mention contraire ; les labels de
provenance sont conservés verbatim.

## Non audité

Contenu non relu section par section : aucun défaut de source recensé. Fidélité garantie
par le vérificateur (concaténation == source).

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
