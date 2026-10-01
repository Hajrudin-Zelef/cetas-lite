---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-8
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Meta", "Mistral", "Moonshot", "Together AI"]
dates: []
keywords: ["attention", "benchmarks", "deepseek", "fp8", "gqa", "int4", "kimi", "kv cache", "llama", "mai", "mistral", "quantization"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [810, 930]
sha256: d10e67321990cb60db2cc8fd654b91a99e24cdbe16c76c0c4beb5216857c2e99
---

# Encyclopédie des modèles IA — Volume 3

Conséquence directe : en serving, le nombre de requêtes simultanées (le *batch*) est limité non par le calcul mais par la **VRAM disponible pour les caches**. Un serveur qui tient 100 requêtes à 4K tokens n'en tiendra que ~3 à 128K tokens. C'est le premier goulot du RAG à gros contexte.

## 59. La formule du KV cache (et comment la calculer toi-même)

La formule standard, par token :

```
cache_par_token = 2 × n_couches × n_têtes_KV × dim_tête × octets
```

- `2` : Key + Value.
- `n_couches` : nombre de couches du transformer.
- `n_têtes_KV` : nombre de têtes de clés/valeurs (**pas** le nombre de têtes de requêtes — voir GQA, section 62).
- `dim_tête` : dimension par tête (souvent 128).
- `octets` : 2 en FP16/BF16, 1 en FP8/INT8.

Et pour une séquence :

```
cache_total = cache_par_token × n_tokens
```

Exemple de lecture d'une fiche modèle : « 32 couches, GQA 8 KV heads, head_dim 128 » → par token en FP16 : 2 × 32 × 8 × 128 × 2 = 131 072 octets = **128 Kio par token**. À 32K tokens : 4 Gio. À 128K : 16 Gio. Le calcul tient sur un coin de table — fais-le toujours avant de choisir une fenêtre de contexte.

Note : les modèles à attention hybride (DeepSeek V4, Qwen4-Exp) ou à MLA n'appliquent pas cette formule naïvement — voir sections 63–64.

## 60. Exemple chiffré : Llama 70B à 128K tokens

Prenons l'archétype du 70B (type Llama 3.1 70B — chiffres d'architecture publics et stables) :

| Paramètre | Valeur |
|---|---|
| Couches | 80 |
| Têtes KV (GQA) | 8 |
| dim_tête | 128 |
| Précision cache | FP16 (2 octets) |

Calcul par token : 2 × 80 × 8 × 128 × 2 = **327 680 octets ≈ 320 Kio/token**.

| Contexte | KV cache |
|---|---|
| 4K tokens | ~1,25 Gio |
| 32K tokens | ~10 Gio |
| 128K tokens | ~40 Gio |

Les poids du modèle en FP16 font ~140 Gio. À 4K de contexte, le cache est négligeable (1 %). À 128K, il ajoute **40 Gio** — soit l'équivalent d'un second petit modèle. Et ça, c'est **par requête** : avec 4 requêtes simultanées à 128K, il faut 160 Gio de cache + 140 Gio de poids = 300 Gio de VRAM. D'où les serveurs 8× H100/H200 pour servir du 70B en long contexte.

## 61. Exemple chiffré : Mistral 7B

L'archétype du petit modèle servable en local (chiffres d'architecture publics) :

| Paramètre | Valeur |
|---|---|
| Couches | 32 |
| Têtes KV (GQA) | 8 |
| dim_tête | 128 |
| Précision cache | FP16 |

Par token : 2 × 32 × 8 × 128 × 2 = **131 072 octets = 128 Kio/token**.

| Contexte | KV cache | Poids FP16 | Total |
|---|---|---|---|
| 4K | 0,5 Gio | ~14 Gio | ~14,5 Gio |
| 32K | 4 Gio | ~14 Gio | ~18 Gio |
| 128K | 16 Gio | ~14 Gio | ~30 Gio |

Lecture : un 7B en FP16 tient sur une RTX 4090 (24 Go) jusqu'à ~32K de contexte confortablement. À 128K, le cache seul (16 Go) + les poids (14 Go) saturent la carte. En **INT4** (poids ~4 Go), la même carte tient le 128K — mais le cache reste en FP16 sauf quantization dédiée (section 66).

## 62. MHA → MQA → GQA : l'effet sur le KV cache

L'évolution des mécanismes d'attention est d'abord une histoire de **réduction du KV cache** :

- **MHA** (Multi-Head Attention, le transformer originel) : chaque tête de requête a sa propre tête K et V. Pour 32 têtes : 32 K + 32 V par couche. Cache maximal.
- **MQA** (Multi-Query Attention) : **une seule** tête K/V partagée par toutes les têtes de requête. Divise le cache par le nombre de têtes (×32 dans l'exemple). Inconvénient : légère perte de qualité.
- **GQA** (Grouped-Query Attention) : compromis — les têtes de requête sont groupées (ex : 8 groupes → 8 têtes KV). Divise le cache par 4 par rapport à MHA avec une perte de qualité quasi nulle. C'est le standard 2024–2026 (Llama 3/4, Mistral, Qwen, Gemma).

Chiffres : sur un 70B à 80 couches, passer de MHA-64 têtes à GQA-8 têtes divise le cache par 8 : 320 Kio/token → 40 Kio/token en théorie. En pratique tous les gros modèles 2026 sont en GQA (ou mieux) — le MHA pur n'existe plus qu'en musée.

## 63. MLA (DeepSeek) : compresser le cache au lieu de le partager

Le **MLA** (Multi-head Latent Attention), introduit par DeepSeek-V2/V3, prend une autre voie : au lieu de réduire le nombre de têtes KV, il **compresse** K et V en un vecteur latent de faible dimension, stocké dans le cache, puis décompressé à la volée par couche.

Effet : le cache ne stocke plus `n_têtes_KV × dim_tête` nombres par couche mais un latent compact — DeepSeek annonçait une réduction d'environ **93 %** du cache par rapport au MHA équivalent sur V3. Le coût : une décompression à chaque étape (calcul en plus, mémoire en moins).

La famille **DeepSeek V4** (avril 2026) abandonne le MLA pur pour une **attention hybride** (section 64) — signe que le MLA seul avait des limites (complexité d'implémentation, interactions avec le parallélisme d'experts). Mais le principe « cache latent compressé » reste une des deux grandes directions 2026, avec le partage de têtes (GQA).

Kimi K2 (1T/32B, section 75) utilise aussi le MLA : KV cache compressé + INT4 natif, d'où son déploiement possible sur 8× H200 malgré le trillion de paramètres.

## 64. Attention hybride : la direction DeepSeek V4

DeepSeek V4 (24 avril 2026, preview ; V4-Pro GA le 13 août 2026) remplace le MLA pur par trois mécanismes combinés :

| Mécanisme | Rôle | Effet cache |
|---|---|---|
| **CSA** (Compressed Sliding Attention), stride 4 | Attention locale compressée | Cache compressé |
| **HCA** (Hybrid Chunk Attention), stride 128 | Attention par blocs | Cache dense ~8K entrées |
| **SWA** (Sliding Window Attention), ~128 | Fenêtre glissante | Cache exact, borné |

D'après la model card : le KV cache de V4 vaut **~10 % de celui de V3.2**. Together AI (mai 2026) documente 3 layouts de cache simultanés, avec une capacité de **~3,7M tokens par nœud HGX B200** en politique optimisée.

Même philosophie chez Qwen4-Exp : 36 couches Gated DeltaNet (attention linéaire, pas de cache KV classique) + 12 couches d'attention standard (ratio 3:1). La règle 2026 : **ne payer l'attention quadratique que sur une minorité de couches**.

## 65. Sliding window et NoPE : l'approche Llama 4

Meta a pris une troisième voie sur Llama 4 (2025, famille toujours servie en 2026) :

- **Attention chunkée 8K** : l'attention standard ne porte que sur des blocs de 8K tokens.
- **Couches NoPE** (No Positional Encoding) : des couches sans encodage positionnel, moins gourmandes en cache.

Résultat revendiqué : Llama 4 Maverick tient sur **un seul H100 en INT4** — impensable pour un modèle de cette classe en attention dense. Le compromis : la qualité sur les dépendances très longue distance peut souffrir (les couches NoPE « voient » moins bien les positions lointaines).

Pour ton dimensionnement : quand une fiche dit « sliding window » ou « chunked attention », le KV cache **ne grandit plus linéairement** au-delà de la fenêtre — il plafonne. C'est une excellente nouvelle pour le serving, à vérifier par des tests de rappel long-contexte.

## 66. Quantization du KV cache (FP8)

On quantize les poids depuis longtemps (section 81) ; quantizer le **cache** est plus récent et plus délicat — le cache contient des activations, pas des poids, avec des outliers qui rendent la quantization naïve destructive.

L'état de l'art 2026 : **FP8 pour le KV cache**. Le format E4M3 (4 bits d'exposant, 3 de mantisse) couvre la dynamique des activations sans calibration lourde. Effet : **÷2 sur la taille du cache** par rapport au FP16, avec une perte de qualité mesurée comme négligeable sur les benchmarks standard.

Exemples concrets :

