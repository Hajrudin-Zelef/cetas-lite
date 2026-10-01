---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-3
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Apple", "Cohere", "DeepSeek", "Google", "Microsoft", "Moonshot", "vLLM"]
dates: []
keywords: ["apache", "attention", "benchmark", "benchmarks", "cohere", "deepseek", "embedding", "embeddings", "foundry", "gemini", "kimi", "omni"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [213, 346]
sha256: 90a032f17e9e2db485dfbca63a056c3d62c7eb1dace123f6c46429cf21754315
---

# Encyclopédie des modèles IA — Volume 3

- **Pointwise** : score chaque paire indépendamment (0–1). Simple, parallélisable.
- **Pairwise** : compare les passages deux à deux. Plus précis, plus cher.
- **Listwise** : lit toute la liste d'un coup et ordonne (jina-reranker-v3/v3.5). Le plus moderne, contexte long requis.

## 14. jina-reranker-v3.5

Le flagship reranker de juillet 2026 (papier arXiv 2607.18152). Le plus avancé techniquement du lot.

| Champ | Valeur |
|---|---|
| Paramètres | 0,6B |
| Type | **Listwise** — ~64 documents par passe |
| Attention | Hybride : 3 couches sliding-window + 2 globales (dernière = lecture LBNL) |
| Contexte | **131K tokens** |
| Langues | 93 |
| Benchmarks (papier, équipe Jina) | BEIR **63,20 nDCG@10** (≈ un modèle 4B à ~7× moins de paramètres) ; +9,6 pts vs v3 sur retrieval semi-structuré ; latence listwise réduite jusqu'à 1,56× |
| Licence | Poids ouverts, **non commerciale** (CC-BY-NC) |
| Déploiement | API Jina, self-host, portage MLX (Apple Silicon) |

Le chiffre qui compte : 63,20 BEIR avec 0,6B paramètres, là où il fallait ~4B avant. L'attention hybride (sliding-window + globale) est la même philosophie que DeepSeek V4 (section 64) : payer l'attention globale seulement où elle sert.

Limite : licence non commerciale, comme les embeddings v5. Pour un RAG perso, aucun problème.

## 15. Cohere Rerank 4.0 (Pro / Fast)

Sorti le 11 décembre 2025 — hors période, leader actif au snapshot.

| Champ | Valeur |
|---|---|
| Variantes | Pro (précision max) / Fast (latence réduite) |
| Contexte | 32K tokens |
| Langues | 100+ |
| Positionnement | Pipelines RAG agentiques, texte et semi-structuré |
| Prix (source Microsoft/Azure) | Pro : $2,50 / 1000 search units ; Fast : $2,00 / 1000 search units |
| Licence | Fermé, API uniquement (Cohere, Azure AI Foundry, OCI) |

L'unité de facturation (« search unit ») est propre à Cohere : ce n'est pas du « par million de tokens ». En pratique, une search unit couvre une requête + ses documents à reranker. Pour chiffrer, il faut tester avec ton volume réel.

## 16. Voyage rerank-2.5 / rerank-2.5-lite

Août 2025 — hors période, leader actif.

| Champ | Valeur |
|---|---|
| Contexte | 32K |
| Capacité annoncée | Jusqu'à 1000 documents par appel (fournisseur) |
| Prix (secondaire) | $0,05/M et $0,02/M repérés — unité exacte NON VÉRIFIÉE |
| Licence | Fermé, API uniquement |

1000 documents par appel, c'est énorme pour un reranker (la norme est 20–100). Si c'est vrai en pratique, ça permet de reranker large sans multiplier les appels. Mais le prix et l'unité de facturation ne sont pas confirmés — à vérifier sur docs.voyageai.com avant tout chiffrage.

## 17. jina-reranker-v3

Octobre 2025 — la version précédente du flagship Jina, toujours pertinente.

| Champ | Valeur |
|---|---|
| Paramètres | 0,6B, listwise, contexte 131K, ~64 documents |
| Benchmarks (carte officielle / papier) | BEIR 61,94 ; MIRACL 66,83 ; MKQA 67,92 ; CoIR 70,64 (nDCG@10) |
| Licence | Poids ouverts, CC-BY-NC |

L'écart v3 → v3.5 : +1,26 pt BEIR (61,94 → 63,20) et +9,6 pts sur le semi-structuré. Le semi-structuré (tableaux, JSON, documents avec structure) progresse plus vite que le texte pur — logique : c'est là que les modèles avaient le plus de marge.

## 18. Qwen3-Reranker-8B / 4B / 0.6B

Juin 2025 — hors période. La famille open-weight **Apache 2.0**.

| Champ | Valeur |
|---|---|
| Tailles | 0.6B → 4B → 8B (choisir selon le budget latence) |
| Contexte | 32K |
| Langues | 100+ |
| Prix | Self-host, gratuit hors infra |
| Déploiement | vLLM / Transformers |

C'est l'option « licence propre » : Apache 2.0, trois tailles pour arbitrer qualité/latence. Le 0.6B concurrence directement les Jina en taille ; le 8B vise la qualité max en self-host. Les chiffres comparatifs affichés sur la carte Jina v3 sont à recouper avec la carte officielle Qwen (NON VÉRIFIÉ sept 2026).

## 19. mxbai-rerank-large-v2

Mixedbread — le challenger open-weight.

| Champ | Valeur |
|---|---|
| Paramètres | 1,5B |
| Contexte | 8192 (signalé, NON VÉRIFIÉ) |
| Benchmark (carte Jina v3, secondaire) | BEIR 61,44 — à recouper |
| Licence | **Apache 2.0** (selon sources secondaires) |

Si la licence Apache 2.0 se confirme, c'est le meilleur rapport qualité/licence du lot open : 61,44 BEIR (proche des Jina) en 1,5B paramètres, utilisable commercialement.

## 20. Mentions courtes (rerankers)

- **BAAI bge-reranker-v2-m3** : 568M paramètres, 8192 tokens, multilingue, Apache 2.0. BEIR 56,51 / MIRACL 69,32 (carte Jina, secondaire). Le compagnon naturel de BGE-M3 : même équipe, même licence, pipeline 100 % Apache 2.0.
- **jina-reranker-m0** (avril 2025) : toujours la seule option multimodale de Jina au snapshot — reranker capable de prendre des images en entrée. Pas de successeur repéré.
- **Non vérifiés** : `mxbai-rerank-v3.1-listwise`, `ctxl-rerank-v2-instruct-multilingual` (Contextual AI), ReasonRank — repérés mais NON TROUVÉS (aucune fiche primaire ouverte). Ne pas les intégrer à une architecture sans vérification.

## 21. Tableau comparatif des rerankers

| Modèle | Taille | Type | Contexte | Langues | Licence | Prix |
|---|---|---|---|---|---|---|
| jina-reranker-v3.5 | 0,6B | Listwise | 131K | 93 | CC BY-NC | API au token / self-host |
| Cohere Rerank 4.0 Pro | Fermé | — | 32K | 100+ | Propriétaire | $2,50 / 1000 search units |
| Cohere Rerank 4.0 Fast | Fermé | — | 32K | 100+ | Propriétaire | $2,00 / 1000 search units |
| Voyage rerank-2.5 | Fermé | — | 32K | — | Propriétaire | ~$0,05/M (non vérifié) |
| jina-reranker-v3 | 0,6B | Listwise | 131K | Multi | CC BY-NC | API au token / self-host |
| Qwen3-Reranker-8B | 8B | — | 32K | 100+ | Apache 2.0 | Self-host |
| Qwen3-Reranker-0.6B | 0,6B | — | 32K | 100+ | Apache 2.0 | Self-host |
| mxbai-rerank-large-v2 | 1,5B | — | 8K | — | Apache 2.0 (?) | Self-host |
| bge-reranker-v2-m3 | 568M | — | 8K | Multi | Apache 2.0 | Self-host |

## 22. Quand le rerank vaut le coup (et quand non)

Le rerank n'est pas gratuit : latence (+200 ms à +2 s par requête) et coût (un second modèle à servir ou à payer). Il vaut le coup quand :

- **Ton top-5 est bruité.** Si la bonne réponse est souvent en position 8–30 du retrieval, le rerank la remonte. C'est le cas typique des corpus techniques où le vocabulaire se répète (« port », « VLAN », « PoE » apparaissent partout).
- **Tes chunks sont longs.** Plus le chunk est long, plus l'embedding dilue le signal. Le reranker, lui, lit le passage en entier avec la question.
- **Tu as du semi-structuré.** Tableaux de specs, tableaux CLI : le rerank listwise (Jina v3.5) y gagne +9,6 pts.

Il ne vaut pas le coup quand :

- Ton retrieval est déjà excellent (rappel@5 > 95 % sur tes tests) — le rerank n'ajoute rien.
- Ta latence est critique (chat temps réel) et ton budget serré.
- Ton corpus est petit (< 1000 chunks) : un bon embedding + un bon chunking suffisent souvent.

Règle pratique : implémente le rerank comme une **option activable**, mesure avec et sans sur tes 50 questions de test (section 11), et ne le garde que si le gain est mesurable.

## 23. Vision : les VLM en 2026

Les modèles vision-langage (VLM) prennent des images — et souvent de la vidéo et de l'audio — en entrée, en plus du texte. En 2026, trois tendances :

1. **L'omni-modalité** : un seul modèle pour texte + image + audio + vidéo (Nemotron 3 Nano Omni, Qwen3.5-Omni, Gemini 3.1 Pro). Fini les encodeurs bricolés à côté.
2. **La vision native** : les patchs d'image sont tokenisés directement dans le flux du transformer, pas via un encodeur CLIP figé. Meilleure compréhension des documents denses (tableaux, schémas).
3. **La taille n'est plus le sujet** : Qwen3.5-Omni sort en 0.8B/2B pour l'edge, pendant que Kimi K3 pousse à 2,8T paramètres totaux. Le marché se polarise : minuscule ou gigantesque.

