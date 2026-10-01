---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-2
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Cohere", "Google", "vLLM"]
dates: []
keywords: ["apache", "attention", "aws", "bedrock", "benchmark", "benchmarks", "cohere", "embedding", "embeddings", "gemini", "gpu", "moe"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [121, 212]
sha256: 8acdd55c6fd4bcc230af2a966e7ac933d723c0de752365e0ae910e2600ca4b32
---

# Encyclopédie des modèles IA — Volume 3

| Champ | Valeur |
|---|---|
| Nom API | `embed-multimodal-v4` |
| Dimensions | 256 / 512 / 1024 / 1536 |
| Contexte | **128K tokens** — le plus long du comparatif |
| Multimodal | Texte + image |
| Langues | 100+ |
| Prix | $0,12/M tokens repéré sur AWS Bedrock ; grille directe Cohere NON VÉRIFIÉE |
| Licence | Fermé, API uniquement (Cohere, Bedrock, Azure, OCI) |

Le contexte 128K est unique : tu peux y faire passer un document entier (un chapitre de guide de 2500 lignes ≈ 30–40K tokens) sans chunking préalable, et laisser le modèle gérer. En pratique, les longs contextes d'embedding dégradent la précision au-delà d'une certaine longueur (voir BGE-M3 qui tombe à 0,920 à 8K dans le test Zilliz) — 128K de contexte ne veut pas dire 128K de précision uniforme.

## 8. Qwen3-Embedding-8B

Sorti le 5 juin 2025 — **hors période**, leader open-weight encore actif au snapshot. Licence **Apache 2.0** : usage commercial autorisé.

| Champ | Valeur |
|---|---|
| Paramètres | 8B |
| Dimensions | 4096 (MRL — réductible) |
| Contexte | 32K |
| Langues | 100+ |
| Benchmarks (carte officielle) | MTEB multilingue 70,58 ; MTEB English v2 75,22 ; CMTEB 73,84 |
| Prix | Self-host (gratuit hors infra) |
| Déploiement | vLLM / Transformers |

C'est la référence open-weight « gros modèle » : les scores MTEB viennent de la carte officielle (donc fournisseur, mais standardisés et reproductibles). 8B paramètres en embedding, c'est ~16 Go en FP16 : il faut un GPU 24 Go pour le servir confortablement.

## 9. Autres embeddings à connaître

Quatre mentions courtes, utiles comme points de repère :

- **Jina Embeddings v4** (2025) : 3,8B paramètres, dense 2048 réductible à 128, multi-vecteur 128/token, texte+image, 30+ langues, contexte 32K selon Qdrant, CC-BY-NC-4.0. Le multi-vecteur (un vecteur par token, style ColBERT) donne un retrieval plus fin mais un index beaucoup plus lourd.
- **BGE-M3** (BAAI) : 568M paramètres, multilingue, **Apache 2.0**. Encore la baseline Apache-2.0 en 2026. Mais le test Zilliz 2026 le fait dégrader à 0,920 à 8K — son contexte nominal de 8192 tokens n'est pas uniformément fiable.
- **nomic-embed-text-v2-moe** : 475M paramètres totaux / 305M actifs (MoE), 768 dimensions, Apache 2.0, ~100 langues, MIRACL 65,80. Limite dure : **contexte 512 tokens seulement**. Le MoE appliqué aux embeddings : moins de calcul par token pour une qualité donnée.
- **Snowflake arctic-embed-l-v2.0** : base XLM-R, Apache 2.0, 74 langues, contexte 8192. L'option « entreprise sobre » : Snowflake le maintient pour sa plateforme, licence permissive.

## 10. Tableau comparatif des embeddings

| Modèle | Taille | Dims | Contexte | Multilingue | Multimodal | Licence | Prix ordre de grandeur |
|---|---|---|---|---|---|---|---|
| Nemotron-3-Embed-8B | 8B | 4096 | 32K | 34 langues | Texte+visuel | OpenMDW-1.1 | Self-host |
| Jina v5 small | 239M | 1024→32 | 32K | Oui | Non | CC BY-NC 4.0 | Self-host / API |
| Jina v5 nano | 677M | 1024→32 | 32K | Oui | Non | CC BY-NC 4.0 | Self-host / API |
| Gemini Embedding 2 | Fermé | 128–3072 | 8K | 100+ | T/I/PDF/A/V | Propriétaire | ~$0,20/M (non vérifié) |
| Voyage 4 large | Fermé | 256–2048 | 32K | Oui | Non | Propriétaire | $0,12/M |
| Voyage 4 lite | Fermé | 256–2048 | 32K | Oui | Non | Propriétaire | $0,02/M |
| Cohere Embed v4 | Fermé | 256–1536 | 128K | 100+ | Texte+image | Propriétaire | ~$0,12/M (Bedrock) |
| Qwen3-Embedding-8B | 8B | 4096 (MRL) | 32K | 100+ | Non | Apache 2.0 | Self-host |
| BGE-M3 | 568M | 1024 | 8K | Oui | Non | Apache 2.0 | Self-host |
| nomic v2 MoE | 475M/305M | 768 | 512 | ~100 | Non | Apache 2.0 | Self-host |

Lecture du tableau : si tu veux du **gratuit + commercial**, le trio est Qwen3-Embedding-8B (gros, précis), BGE-M3 (léger, baseline), nomic v2 MoE (léger, mais 512 tokens). Si tu veux du **léger + 32K**, Jina v5 — mais la licence NC bloque le commercial.

## 11. Focus : le français dans les embeddings

C'est ton point d'attention depuis le début (noté le 26/09) : ton corpus est en français, tes questions sont en français. Trois constats :

1. **Les benchmarks anglophones dominent.** MTEB English v2 est la vitrine, mais ton usage réel se joue sur MMTEB (multilingue) et MIRACL (18 langues dont le français). Un modèle à 75 en anglais peut être médiocre en français. Exemple : Qwen3-Embedding-8B affiche 70,58 en MTEB multilingue contre 75,22 en anglais — l'écart existe et il est normal.
2. **Le français est une langue « bien servie ».** Tous les modèles du comparatif couvrent le français (34 langues pour Nemotron, 100+ pour les autres). Le risque n'est pas l'absence de français, c'est la **qualité relative** : les modèles entraînés majoritairement sur de l'anglais + chinois (Qwen) peuvent sous-performer sur les tournures techniques françaises (« onduleur », « parafoudre », « consignation électrique »).
3. **Le test qui compte, c'est le tien.** Prends 50 questions réelles en français sur ton corpus (ex : « quelle est la procédure de remplacement du kit de fusion sur Kyocera ? »), mesure le rappel@5 de chaque embedding candidat. Un après-midi de test vaut mieux que tous les leaderboards.

En 2026, la recommandation pragmatique pour du français : privilégier les modèles avec un vrai score **MMTEB Retrieval** publié (Nemotron 75,45, Qwen3-Embedding 70,58) plutôt qu'un score anglais seul.

## 12. Lire un benchmark d'embeddings sans se faire avoir

Petit décodeur des acronymes que tu croiseras :

| Benchmark | Ce qu'il mesure | Piège |
|---|---|---|
| MTEB (English v2) | 41 tâches en anglais : retrieval, clustering, classification, STS | Anglais uniquement — ne dit rien du français |
| MMTEB | Version multilingue de MTEB | Le score agrégé masque les écarts par langue |
| MIRACL | Retrieval multilingue (18 langues) | Qualité variable des jeux par langue |
| BEIR | Retrieval zero-shot (18 datasets) | Certains datasets sont vieux / saturés |
| RTEB | Nouveau benchmark retrieval (2026) | Jeune — voir l'alerte Voyage (section 6) |
| ViDoRe | Retrieval visuel sur documents | Pertinent si tu indexes des PDF scannés |

Règles de lecture :

- **Un score fournisseur n'est pas un mensonge, c'est une vitrine.** Le fournisseur choisit le benchmark qui l'arrange et le split qui l'arrange. Le cas Voyage 4 (split privé RTEB contesté) montre jusqu'où ça peut aller.
- **Compare à tâche égale.** « 78,46 RTEB » vs « 70,58 MMTEB » : ce ne sont pas les mêmes épreuves, la comparaison directe n'a aucun sens.
- **Le needle-in-a-haystack n'est pas du retrieval.** Le 1,000 de Gemini Embedding 2 (Zilliz) mesure la capacité à retrouver une aiguille, pas à classer 50 passages pertinents par ordre d'utilité. Les deux comptent, ce n'est pas la même chose.
- **Ton corpus > leurs benchmarks.** Tes guides font 1600–4200 lignes de français technique avec des tableaux, des codes d'erreur, des procédures. Aucun benchmark public ne ressemble à ça.

## 13. Rerankers : pourquoi un second passage

Le pipeline classique : l'embedding retrouve vite les 50–100 candidats (recherche vectorielle approximative, quelques millisecondes), puis le **reranker** relit chaque paire (question, passage) en profondeur et re-classe. C'est un cross-encoder : il voit la question ET le passage ensemble, là où l'embedding les a encodés séparément.

Coût : le rerank est 10 à 100× plus lent que la recherche vectorielle — d'où l'architecture en deux temps (beaucoup de candidats pas chers, peu de candidats chers). Gain typique : +5 à +15 points de nDCG@10 sur le top-5 final, ce qui se traduit directement en meilleures réponses du LLM.

Trois familles :

