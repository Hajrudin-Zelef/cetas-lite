---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-9
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Google", "Moonshot", "Nvidia", "Oracle", "SGLang", "vLLM"]
dates: []
keywords: ["attention", "blackwell", "deepseek", "fp4", "fp8", "glm", "gqa", "int4", "kimi", "kv cache", "llama", "llama.cpp"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [931, 1027]
sha256: 1a7266c5a54c557df7f70189299d81e1ffd262596d66acc5f1b800c33ec4ed27
---

# Encyclopédie des modèles IA — Volume 3

- **Qwen3.5-122B-A10B** : checkpoint **FP8 officiel** (128,4 Go en FP8, voir qwen_kimi_glm.md) — poids ET cache en FP8, déploiement vLLM/SGLang day-0.
- **DeepSeek V4** : formats BF16 / FP8 / **FP4 natifs** (NVFP4 sur Blackwell) — la quantization est conçue dès l'entraînement (QAT), pas appliquée après coup.
- **Gemma 4** : checkpoints **NVFP4** officiels Google — « quasi-identique au 8-bit » selon Google.

Règle pratique : si ton serveur est limité par le cache (long contexte, gros batch), passer le cache en FP8 équivaut à **doubler ta VRAM utile pour le cache**. vLLM et SGLang le supportent en option (`--kv-cache-dtype fp8`).

## 67. Eviction : StreamingLLM, H2O — jeter pour survivre

Quand le cache déborde, une autre stratégie : **ne pas tout garder**. Deux méthodes académiques devenues classiques :

- **StreamingLLM** (MIT, 2023) : garde les premiers tokens (le « attention sink » — le modèle s'ancre dessus) + une fenêtre glissante des derniers tokens, jette le milieu. Permet une génération « infinie » en mémoire constante. Limite : si l'info utile est au milieu du contexte, elle est perdue.
- **H2O** (Heavy-Hitter Oracle, 2023) : garde les tokens « heavy hitters » (ceux qui reçoivent le plus d'attention cumulée) + la fenêtre récente, évince le reste. Plus fin que StreamingLLM, avec un budget de cache fixe.

En 2026, ces méthodes sont intégrées aux moteurs d'inférence comme **politiques d'éviction** quand le cache est plein, plutôt que comme mode par défaut. Pour un RAG : l'éviction est dangereuse si tes chunks pertinents sont au milieu du prompt (le fameux « lost in the middle » — les modèles performent moins bien sur l'info au milieu d'un long contexte). Préfère un cache entier + un bon tri des chunks (les plus pertinents en premier/dernier) plutôt que l'éviction.

## 68. PagedAttention : la mémoire virtuelle du LLM

**PagedAttention** (vLLM, 2023) est l'innovation système la plus importante de l'inférence moderne : au lieu d'allouer au KV cache un bloc contigu par requête (avec du gaspillage quand les requêtes ont des longueurs différentes), il découpe le cache en **pages** (blocs) allouées à la demande, comme la mémoire virtuelle d'un OS.

Effets concrets :

- **Zéro gaspillage** : plus de pré-allocation « au max de la fenêtre » par requête.
- **Batching continu** : les requêtes entrent et sortent du batch à chaque étape de génération, au lieu d'attendre la fin du batch (vs batching statique).
- **Partage de préfixes** : si 10 requêtes partagent le même prompt système, le cache du préfixe est stocké **une fois** — énorme pour un RAG où le prompt système + les instructions sont identiques à chaque appel.

C'est pour ça que **vLLM bat Ollama/llama.cpp en débit serveur** d'un facteur 2 à 10× : ce n'est pas le modèle qui change, c'est la gestion mémoire. Si tu sers un LLM en production (même petite), vLLM ou SGLang (qui a son équivalent) est le choix par défaut.

## 69. Tableau : VRAM du KV cache selon modèle et contexte

Hypothèses : GQA-8, head_dim 128, cache FP16. Poids FP16 inclus pour le total.

| Modèle | Par token | 4K | 32K | 128K | Poids FP16 | Total à 32K |
|---|---|---|---|---|---|---|
| 7B (32 couches) | 128 Kio | 0,5 Go | 4 Go | 16 Go | ~14 Go | ~18 Go |
| 13B (40 couches) | 160 Kio | 0,6 Go | 5 Go | 20 Go | ~26 Go | ~31 Go |
| 70B (80 couches) | 320 Kio | 1,25 Go | 10 Go | 40 Go | ~140 Go | ~150 Go |
| 8×70B MoE type Mixtral (32 couches, 8 experts) | 128 Kio | 0,5 Go | 4 Go | 16 Go | ~90 Go | ~94 Go |

Lecture :

- En dessous de 32K, le cache est un poste secondaire. Au-delà, il devient le poste dominant du dimensionnement.
- Le MoE ne change rien au cache par token (le cache dépend des couches d'attention, pas des experts FFN) — mais les poids MoE sont bien plus lourds (section 74).
- En FP8, divise la colonne cache par 2. En INT4 pour les poids, divise la colonne poids par ~4.

## 70. Ce que le KV cache change pour ton RAG

Traduction opérationnelle pour ton pipeline :

1. **Le nombre de chunks injectés dans le prompt a un coût VRAM direct.** 20 chunks de 1000 tokens = 20K tokens de contexte = pour un 7B, 2,5 Go de cache par requête. ×10 requêtes simultanées = 25 Go. Dimensionne en fonction.
2. **Le prompt système partagé est gratuit (ou presque) avec vLLM.** Grâce au partage de préfixes (section 68), tes instructions RAG identiques à chaque appel ne coûtent qu'une fois. Raison de plus pour servir en vLLM.
3. **Ne mets pas 1M de tokens « au cas où ».** La tentation des fenêtres géantes (section 115) : balancer tout le corpus dans le prompt. Le cache rend ça hors de prix — le retrieval + rerank reste 10 à 100× moins cher par question.
4. **Trie les chunks par pertinence décroissante.** Le « lost in the middle » + le coût du cache convergent vers la même recommandation : peu de chunks, bien classés, plutôt qu'une marée de contexte.

## 71. MoE : le principe

**MoE** (Mixture of Experts) : au lieu d'un seul gros réseau feed-forward (FFN) par couche, le modèle contient **N experts** (des FFN plus petits) + un **routeur** (un petit réseau qui décide, pour chaque token, quels experts consulter).

Analogie : un hôpital. Le routeur est l'accueil : il envoie le patient (token) vers 2 spécialistes sur 200 (les experts), pas vers tout l'hôpital. Chaque expert se spécialise (un expert « code Python », un expert « français technique », un expert « raisonnement math » — en pratique les spécialisations émergent, elles ne sont pas programmées).

L'intérêt : **découpler la capacité (paramètres totaux) du coût de calcul (paramètres actifs)**. Un Mixtral 8×7B a 47B paramètres mais n'en active que ~13B par token — le coût d'inférence d'un 13B avec la capacité d'un 47B.

## 72. Routage des experts : top-k, partagés, granularité

Le routeur attribue un score à chaque expert pour chaque token, et sélectionne les **top-k** :

| Modèle (2026) | Experts totaux | Sélection | Expert(s) partagé(s) |
|---|---|---|---|
| Mixtral 8×7B (référence 2023) | 8 | top-2 | 0 |
| DeepSeek V4-Flash | 256 | top-6 | 1 |
| Kimi K2 / K2.6 / K2.7-Code | 384 | top-8 | 1 |
| Kimi K3 | 896 | top-16 | 2 |
| Qwen4-Exp | 512 | top-10 | 1 |
| Qwen3.6-35B-A3B | 256 | — | — |

Deux subtilités :

- **L'expert partagé** (shared expert) traite *tous* les tokens : il capture le savoir commun (grammaire, faits généraux), pendant que les experts routés se spécialisent. DeepSeek l'a introduit sur V3, tout le monde l'a copié.
- **La granularité fine** : la tendance 2026 est « beaucoup de petits experts, beaucoup de sélectionnés » (K3 : 896 experts, top-16) plutôt que « peu de gros experts » (Mixtral : 8, top-2). Plus de granularité = routage plus précis = meilleure qualité à budget de calcul égal.

Le risque du routage : le **déséquilibre de charge** (load imbalance) — si le routeur envoie 80 % des tokens vers 10 % des experts, les autres dorment et le calcul est gaspillé. Les entraînements modernes ajoutent une perte d'équilibrage (auxiliary loss) pour forcer la répartition.

## 73. Actifs vs totaux : lire une fiche modèle

La nomenclature 2026 s'est standardisée : `Nom-Total-AActifs`. Exemples tirés des recherches :

| Nom | Total | Actifs/token | Ratio |
|---|---|---|---|
| Qwen3.5-122B-A10B | 122B | 10B | 8 % |
| Qwen3.6-35B-A3B | 35B | 3B | 9 % |
| Qwen3.5-397B-A17B | 397B | 17B | 4 % |
| Kimi K2.6 | 1T | 32B | 3 % |
| Kimi K3 | 2,8T | 104B | 4 % |
| Qwen4-Exp | 125B | 6B | 5 % |
| DeepSeek V4-Pro | ~1,6T | 49B | 3 % |
| Nemotron 3 Nano Omni | 30B | 3B | 10 % |
| nomic-embed-text-v2-moe | 475M | 305M | 64 % |

