# NOTES — corpus `frontier-models-2026`

## Génération

Produit en **mode auto** (`RAG/_tools/build_rag.py`, clé `mode: auto`) :

- partition par titres **H1 → H2 → H3**, sans jamais couper au milieu d'un paragraphe ;
- cibles de taille : 50–110 lignes par chunk (fusion des petits blocs, découpe des gros
  aux H3, et à défaut sur lignes vides) ;
- un dossier par document H1 ; un fichier par bloc (nommé d'après son premier titre,
  `overview` pour l'entête de section).

La source ne contient **aucune ancre** HTML (`<a id=…>`) : l'index n'a pas de table
ancre→fichier. La navigation repose sur `INDEX.md` (domaines, tâches, acteurs, dates) et
`manifest.json`.

## Contenu

Fiches de recherche concaténées (série « VOLET 1 — Vague 1 ») :

- vague de février 2026 : Claude Sonnet 5 « Fennec », GPT-5.3-Codex, Gemini 3.1 Pro ;
- Grok 4.20, DeepSeek V4 / V4.1 Flash, GLM-5.2 / 5.3 / 5.3-Flash ;
- modèles open-weight chinois (MiMo-V2.6, Kimi K3, Qwen3.8 Max, MiniMax M2.5/M3) ;
- comparaisons transverses, tendance globale, mises en garde méthodologiques.

Les chiffres de benchmark sont **annoncés par les éditeurs** (vendor-reported), sauf
mention contraire dans le texte ; les labels de provenance sont conservés verbatim.

## Non audité

Contenu non relu section par section : aucun défaut de source n'est recensé ici. La
fidélité est garantie par le vérificateur (concaténation == source), pas par une relecture
sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
