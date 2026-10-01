---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-5
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["DeepSeek", "Mistral", "SGLang", "vLLM"]
dates: ["2026-09-27"]
keywords: ["datacenter", "gpu", "attention", "compute", "containment", "deepseek", "fine-tuning", "fp4", "fp8", "gguf", "gqa", "int4"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [538, 692]
sha256: 2dde84f92df8df85959bddc97ef02cc3e757122f71e496f9632c0d8d269240d9
---

# Les GPU datacenter / IA — présent et futur vérifié

**Différences clés** :
- **DLSS** : propriétaire, Tensor Cores requis, meilleure qualité historique, génération
  d'images verrouillée à la génération de GPU en cours (politique commerciale assumée).
- **FSR** : ouvert, fonctionne sur tous les GPU ; FSR 4 passe au ML et ferme l'écart
  qualité avec DLSS (tests 2025-2026 : quasi-équivalent en mode Quality).
- **XeSS** : deux chemins — XMX (Arc, qualité max) et DP4a (tous GPU, qualité réduite).

**Usage pro / datacenter : marginal.** Aucun lien avec le training ou l'inférence LLM.
Deux exceptions honnêtes : (1) les Tensor Cores qui accélèrent DLSS sont les mêmes qui
accélèrent le FP8/FP4 — c'est le même silicium, pas la même fonction ; (2) cloud gaming
(GeForce NOW) où DLSS réduit le coût GPU par stream. Pour un datacenter IA : **ignorer**.
## 22. Dimensionnement VRAM pour les LLM

### 22.1. Règle de calcul n°1 : les poids

```
VRAM_poids = N_paramètres × octets_par_paramètre
```

| Précision | Octets/param | 7B | 13B | 70B | 405B |
|---|---|---|---|---|---|
| FP32 | 4 | 28 Go | 52 Go | 280 Go | 1 620 Go |
| FP16 / BF16 | 2 | 14 Go | 26 Go | 140 Go | 810 Go |
| FP8 | 1 | 7 Go | 13 Go | 70 Go | 405 Go |
| INT8 | 1 | 7 Go | 13 Go | 70 Go | 405 Go |
| INT4 / Q4 | 0,5 | 3,5 Go | 6,5 Go | 35 Go | 202 Go |
| INT2 / Q2 (expérimental) | 0,25 | 1,75 Go | 3,25 Go | 17,5 Go | 101 Go |

Ajouter **10–20 % de marge** (overhead framework, fragmentation, activations) :
multiplier par **1,15** en pratique.

### 22.2. Règle de calcul n°2 : mémoire totale d'inférence

```
VRAM_totale = VRAM_poids × 1,15 + KV_cache + marge_serveur (2–4 Go)
```

En training, ajouter : gradients (1× poids), états optimiseur Adam (2× poids en FP32
= 8 octets/param), activations. Règle empirique training :

```
VRAM_training ≈ N_param × (2 + 4 + 8) × 1,3 ≈ N_param × 18 octets (Adam, mixed precision)
```

Soit **~126 Go pour 7B** en full fine-tuning Adam — d'où l'intérêt de LoRA/QLoRA
qui divise par 3 à 5.

## 23. Le KV cache : le second compteur

### 23.1. Formule

À chaque token généré, le modèle stocke les clés/valeurs d'attention de tous les
tokens précédents :

```
KV_cache = 2 × n_couches × n_têtes_kv × dim_tête × n_tokens × octets
```

Le facteur 2 = clés + valeurs. En GQA (Grouped-Query Attention, standard depuis
Llama 3), n_têtes_kv < n_têtes_q, ce qui divise le cache par 4 à 8 vs MHA classique.

### 23.2. Ordres de grandeur (BF16, contexte 8K)

| Modèle | Couches | Têtes KV | KV cache / token | Cache @ 8K tokens | Cache @ 128K |
|---|---|---|---|---|---|
| Llama 3.1 8B (GQA 8) | 32 | 8 | ~0,5 Mo | ~4 Go | ~64 Go |
| Llama 3.1 70B (GQA 8) | 80 | 8 | ~1,25 Mo | ~10 Go | ~160 Go |
| Llama 3.1 405B (GQA 8) | 126 | 8 | ~2 Mo | ~16 Go | ~256 Go |
| Mixtral 8×22B (GQA) | 56 | 8 | ~1 Mo | ~8 Go | ~128 Go |

**Leçon** : à 128K de contexte, le KV cache d'un 70B (160 Go) **dépasse les poids FP16
(140 Go)**. Le contexte long est un problème de *mémoire*, pas de compute.

### 23.3. Leviers pour réduire le KV cache

- **Quantification du cache** (KV FP8 / INT8) : ÷2, supporté par vLLM/SGLang.
- **Sliding window / attention locale** (Mistral) : cache borné.
- **Éviction** (StreamingLLM, H2O) : garde les tokens « importants » — perte qualité.
- **MLA** (Multi-head Latent Attention, DeepSeek) : compression apprise du cache.
- **Prefix caching** : partage le cache des prompts système entre requêtes (vLLM).

## 24. Exemples chiffrés de 7B à 405B

Hypothèses : poids + 15 % overhead + KV cache BF16 @ 8K, 1 utilisateur (batch 1).

### 24.1. Llama 3.1 8B

| Config | Poids | KV cache | Total | Tient sur |
|---|---|---|---|---|
| BF16 | 16 Go | 4 Go | ~20 Go | RTX 4090 (24 Go) ✓ |
| FP8 | 8 Go | 4 Go | ~12 Go | RTX 4090 ✓, L40S ✓ |
| Q4 (GGUF) | 4,5 Go | 4 Go | ~9 Go | Tout GPU ≥ 12 Go ✓ |

### 24.2. Llama 3.1 70B

| Config | Poids | KV cache | Total | Tient sur |
|---|---|---|---|---|
| BF16 | 161 Go | 10 Go | ~171 Go | **2× H100 NVL** ou 1× MI300X (192 Go) ✓ |
| FP8 | 80 Go | 10 Go | ~90 Go | **1× RTX PRO 6000 (96 Go)** ✓, 1× H100 96 Go ✓ |
| Q4 | 40 Go | 10 Go | ~50 Go | 2× RTX 5090 (64 Go) ✓, 1× RTX PRO 6000 ✓ |

Le 70B FP8 sur **une seule** RTX PRO 6000 96 Go : c'est le « sweet spot » 2025-2026
pour l'inférence 70B économique.

### 24.3. Llama 3.1 405B

| Config | Poids | KV cache | Total | Tient sur |
|---|---|---|---|---|
| BF16 | 931 Go | 16 Go | ~947 Go | 8× H200 (1,1 To) ✓ / 12× H100 |
| FP8 | 466 Go | 16 Go | ~482 Go | **4× H100** (320 Go : non) → **8× H100** ✓ / 4× H200 (564 Go) ✓ |
| Q4 | 233 Go | 16 Go | ~249 Go | 2× H200 (282 Go) ✓ / 4× RTX PRO 6000 (384 Go) ✓ |

### 24.4. DeepSeek R1 / V3 671B (MoE)

- Poids : 671B params mais **37B actifs par token** (MoE). En Q4 : ~335 Go.
- Tient sur **4× H200 NVL** (564 Go) ou 8× H100, voire 8× RTX PRO 6000 (768 Go).
- Le MoE change la donne : le VRAM suit les params totaux, le compute suit les
  params actifs. **Dimensionner la mémoire sur le total, le compute sur l'actif.**

### 24.5. Schéma ASCII : répartition VRAM typique (70B FP8, batch 8, ctx 32K)

```
┌─────────────────────────────────────────────────┐
│ GPU 96 Go (RTX PRO 6000) — Llama 70B FP8         │
│                                                 │
│ ████████████████████████░░░░░░░░░░░░░░░░  80 Go  │ poids FP8
│ ██████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   6 Go  │ KV cache (8 req × 32K)
│ ██░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   3 Go  │ overhead framework
│ ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   7 Go  │ libre (marge)
└─────────────────────────────────────────────────┘
```

## 25. Impact du contexte long (128K → 1M tokens)

| Modèle | Poids FP8 | KV @ 128K (BF16) | KV @ 1M (BF16) | GPU requis @ 128K | GPU requis @ 1M |
|---|---|---|---|---|---|
| 8B | 8 Go | 64 Go | 512 Go | 1× RTX PRO 6000 | 8× H100 |
| 70B | 80 Go | 160 Go | 1 280 Go | 2× H200 / 1× MI350X | 8× B200 (1,5 To) ✓ limite |
| 405B | 466 Go | 256 Go | 2 048 Go | 8× H100 (640 Go : non) → 8× H200 | 16× B200 (rack) |

**Règle terrain** : au-delà de 128K, quantifier le KV cache en FP8 (÷2) devient
obligatoire, sinon le cache mange tout. À 1M tokens, même un 8B exige un nœud 8 GPU.

## 26. Refroidissement : air vs liquide par modèle

### 26.1. Seuils physiques

- **Air classique** (CRAC/CRAH, allée chaude/froide) : viable jusqu'à **15–20 kW/rack**.
- **Air haute densité** (ventilateurs 80 mm à 20 000 tr/min, containment) : jusqu'à
  **30–40 kW/rack**, au prix d'un bruit et d'une conso ventilateurs énormes.
- **Liquide direct-to-chip (DLC)** : plaques froides sur GPU/CPU, **40–120 kW/rack**.
- **Immersion** : 100–250 kW/rack, niche (maintenance compliquée).

### 26.2. Par GPU (vérifié le 27/09/2026)

