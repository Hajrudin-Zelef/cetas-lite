---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-11
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Google", "Moonshot", "Nvidia", "SGLang", "vLLM"]
dates: []
keywords: ["awq", "benchmarks", "blackwell", "datacenter", "deepseek", "fp4", "fp8", "gguf", "gptq", "gpu", "gqa", "int4"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [1114, 1216]
sha256: df0e205f3d0536eff1952a00e0717cd09052624193f6a84c4992c4dfddc62d94
---

# Encyclopédie des modèles IA — Volume 3

1. **Le poste VRAM n°1 en RAG long-contexte, c'est le cache, pas les poids.** Dimensionne le cache d'abord (formule section 59).
2. **Choisis des modèles GQA ou mieux.** En 2026 c'est le défaut, mais vérifie sur les petits modèles (certains 7B « rapides » sont encore en MHA).
3. **Le FP8 pour le cache** (vLLM/SGLang) double ta capacité utile — active-le dès que tu dépasses 32K de contexte.
4. **Sers en vLLM/SGLang**, pas en Ollama, dès que tu as du batch ou des prompts système partagés (PagedAttention + partage de préfixes).
5. **Le MoE est un choix serveur**, pas un choix local solo — sauf les petits MoE bien quantizés.

## 81. Quantization : le principe

La quantization réduit la précision des nombres du modèle : au lieu de stocker chaque paramètre en FP16 (16 bits), on le stocke en INT8 (8 bits), INT4 (4 bits), voire FP4. Effet direct : **÷2, ÷4 sur la taille** — un 70B passe de 140 Go (FP16) à ~70 Go (INT8) à ~35–40 Go (INT4).

Le principe : les poids d'un réseau sont redondants. On peut les arrondir agressivement sans casser la qualité — jusqu'à un point de rupture. L'art de la quantization, c'est de choisir **quels nombres arrondir, et comment compenser l'erreur**.

Deux grandes familles :

- **PTQ** (Post-Training Quantization) : on quantize un modèle déjà entraîné, sans réentraînement. Rapide, mais avec une calibration (un petit jeu de données pour mesurer les plages de valeurs). GPTQ, AWQ, GGUF sont des PTQ.
- **QAT** (Quantization-Aware Training) : on entraîne (ou on fine-tune) en simulant la quantization. Meilleure qualité, mais coûteux. DeepSeek V4 (FP4 natif), Kimi K2 (INT4 natif), les GGUF QAT officiels de Google (Gemma) : la tendance 2026 est au **natif quantizé** — le modèle naît déjà en basse précision.

En 2026, la question n'est plus « faut-il quantizer ? » mais « à quel niveau la qualité devient-elle inacceptable pour mon usage ? ».

## 82. GPTQ : la référence historique GPU

**GPTQ** (Frantar et al., 2022) : la méthode qui a rendu le local possible. Principe : quantization couche par couche, en **compensant l'erreur** — quand on arrondit un poids, on ajuste les poids suivants pour absorber l'erreur d'arrondi (approximation du Hessien, d'où le nom « Generative Pre-trained Transformer Quantization »).

Caractéristiques :

- Formats : INT4/INT8, avec **groupes** (group-size 32/64/128 : les poids sont quantizés par petits groupes, chacun avec sa propre échelle — plus le groupe est petit, meilleure la qualité, plus gros le fichier).
- Calibration : nécessite un jeu de calibration (typiquement 128 échantillons de C4/WikiText) — la qualité dépend du domaine de calibration.
- Usage : bibliothèques AutoGPTQ, intégré à vLLM, Transformers, text-generation-webui.
- Point fort : excellent en INT4 sur les modèles 7B–70B, écosystème mature.

En 2026, GPTQ reste le standard pour servir du quantizé sur **GPU NVIDIA en production** (vLLM). Son successeur spirituel pour le Blackwell : le NVFP4 natif (section 85).

## 83. AWQ : protéger les poids qui comptent

**AWQ** (Activation-aware Weight Quantization, MIT, 2023) : constat — 1 % des poids sont « saillants » (salient) : ils ont un impact disproportionné sur les activations. AWQ identifie ces poids via les activations observées et les **protège** (mise à l'échelle par canal avant quantization).

Différences avec GPTQ :

| Critère | GPTQ | AWQ |
|---|---|---|
| Principe | Compense l'erreur couche par couche | Protège les poids saillants par scaling |
| Calibration | 128 échantillons, sensible au domaine | Plus robuste, moins sensible |
| Vitesse d'inférence | Bonne | Légèrement meilleure (kernels optimisés) |
| Écosystème 2026 | vLLM, mature | vLLM,llama.cpp, mature |

En pratique 2026 : les deux donnent des résultats très proches en INT4. AWQ a un léger avantage de simplicité (pas de problème de « sur-ajustement à la calibration »). Si un modèle propose les deux formats, prends celui qui est le mieux supporté par ton moteur d'inférence.

## 84. GGUF et les quants K : le format du local

**GGUF** (GPT-Generated Unified Format) : le format de **llama.cpp**. Un seul fichier contient le modèle quantizé + les métadonnées + le vocabulaire. C'est le format du local : Ollama, LM Studio, llama.cpp, GPT4All le lisent.

La jungle des quants (du plus gros au plus petit, pour un 7B) :

| Quant | Taille (7B) | Qualité | Usage |
|---|---|---|---|
| Q8_0 | ~7,2 Go | Quasi-identique au FP16 | Référence « sans perte visible » |
| Q6_K | ~5,5 Go | Excellente | Bon compromis si la VRAM suit |
| Q5_K_M | ~4,8 Go | Très bonne | — |
| **Q4_K_M** | **~4,1 Go** | **Bonne (défaut recommandé)** | **Le standard du local** |
| Q4_K_S | ~3,9 Go | Correcte | Un peu moins bon que M |
| Q3_K_M | ~3,3 Go | Dégradation sensible | Dépannage |
| Q2_K | ~2,7 Go | Fortement dégradée | À éviter sauf contrainte extrême |

Le suffixe `_K_M` vs `_K_S` : M (medium) alloue plus de bits aux couches importantes que S (small). Règle : **Q4_K_M par défaut**, Q5_K_M si tu as la place, Q6_K/Q8_0 pour les usages exigeants (code, raisonnement fin).

Chiffres réels tirés des recherches : Qwen3.8-27B en Q4_K_M = **18 Go** (Ollama `qwen3.8:27b`, tourne sur 1× GPU 24 Go) ; Qwen3.5-122B-A10B en Q4_K_M ≈ **71 Go**, Q8_0 ≈ 121 Go ; MiMo Distill-9B en Q4_K_M : 16 Go VRAM suffisants.

## 85. FP8 / NVFP4 : la quantization du datacenter

Le **FP8** existe en deux variantes : E4M3 (précision, pour les poids et activations) et E5M2 (dynamique, pour les gradients). En inférence, E4M3 domine : ÷2 vs FP16 avec une perte quasi nulle, **sans calibration** dans la plupart des cas.

**NVFP4** : le format 4-bit de NVIDIA pour Blackwell (B200/B300). Micro-scaling par blocs : chaque petit bloc de valeurs a son propre facteur d'échelle en FP8, ce qui préserve la dynamique locale. Google annonce ses checkpoints Gemma NVFP4 comme « quasi-identiques au 8-bit ».

État des lieux 2026 :

| Format | Hardware cible | Modèles natifs (recherche) |
|---|---|---|
| FP8 (E4M3) | H100/H200 (support), tous en émulation | Qwen3.5-122B (128,4 Go), DeepSeek V4, Llama 4 Maverick |
| NVFP4 | Blackwell B200/B300/GB300 | DeepSeek V4, Gemma 4 (checkpoints officiels Google) |
| INT4 (natif) | Tous | Kimi K2/K2.6/K2.7 (livré nativement INT4, pas de release FP16) |

La bascule 2026 : les labs **livrent directement en basse précision** (QAT ou entraînement natif) au lieu de laisser la communauté quantizer après coup. Pour toi, ça veut dire : préfère le checkpoint officiel FP8/INT4 au PTQ communautaire quand il existe — la qualité est meilleure à taille égale.

## 86. INT4 vs INT8 vs FP16 : la qualité en pratique

Ordres de grandeur de dégradation (perplexité / benchmarks agrégés — varie selon le modèle et la méthode) :

| Format | Taille relative | Perte typique | Verdict |
|---|---|---|---|
| FP16/BF16 | 100 % | 0 (référence) | Référence qualité |
| INT8 / FP8 | 50 % | < 1 % | Transparent dans 95 % des usages |
| INT4 (GPTQ/AWQ, groupe 128) | ~25–30 % | 1–5 % | Bon pour chat, RAG, résumé |
| INT4 (GGUF Q4_K_M) | ~30 % | 2–6 % | Le standard local, OK pour RAG |
| INT3 / Q3_K | ~20 % | 5–15 % | Dégradation visible (raisonnement, code) |
| INT2 | ~15 % | > 20 % | Inutilisable sauf démo |

Deux nuances importantes :

1. **La taille du modèle protège.** Un 70B en INT4 reste meilleur qu'un 7B en FP16 : la redondance des gros modèles absorbe la quantization. Règle : à budget VRAM fixe, un gros modèle quantizé bat souvent un petit modèle en pleine précision.
2. **Le domaine compte.** La quantization dégrade d'abord le raisonnement fin, le code, les langues rares. Pour un RAG français technique (réponses factuelles, synthèse), l'INT4 est largement suffisant. Pour de la génération de code critique, reste en FP8/INT8.

## 87. Quoi choisir pour du local : l'arbre de décision

