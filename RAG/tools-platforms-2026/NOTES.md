# NOTES — corpus `tools-platforms-2026`

## Génération

Produit en **mode auto** (`RAG/_tools/build_rag.py`, clé `mode: auto`), avec l'option
`first_is_content: true` : la source ne contient **qu'un seul titre H1**, qui est du
contenu (et non un entête de front-matter). Le dossier prend donc son nom, et les chunks
sont répartis par titres H2/H3 à l'intérieur.

- partition par titres **H1 → H2 → H3**, sans jamais couper au milieu d'un paragraphe ;
- cibles de taille : 50–110 lignes par chunk ;
- un fichier par bloc (nommé d'après son premier titre, `overview` pour l'entête).

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

Rapport consolidé « Step 2 — AI Tools & Platforms (Feb–Sep 2026) » :

- agents de code et IDE : Claude Code, OpenCode, OpenClaw, Cursor, VS Code / GitHub
  Copilot, ZCode, Grok Build, DeepSeek Harness ;
- plateformes de distribution : OpenRouter, Hugging Face ;
- autres outils : GitHub Spec Kit, Google Antigravity 2.0, AWS Kiro, Cline, Huawei
  Cloud CodeArts Agent, Augment Code Cosmos ;
- comparaison inter-outils, points de vérification ouverts, métadonnées de collecte.

Les prix sont des **instantanés datés** ; les éléments `[secondary/unverified]` et
`[vendor-reported]` sont marqués dans le texte et conservés verbatim.

## Non audité

Contenu non relu section par section : aucun défaut de source recensé. Fidélité garantie
par le vérificateur (concaténation == source).

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
