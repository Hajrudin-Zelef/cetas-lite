---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-7
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Cohere", "DeepSeek", "Google", "Moonshot", "xAI"]
dates: ["2026-09-27"]
keywords: ["apache", "attention", "cohere", "deepseek", "embedding", "exploit", "gemini", "gpu", "grok", "grok 4", "kimi", "kv cache"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [701, 809]
sha256: 39d9f4d816d6c5295fdba5e686de2eb58feceb317cc36d865635f3ea7e82c7e6
---

# Encyclopédie des modèles IA — Volume 3

Google DeepMind : **Veo 3.1** (`veo-3.1-generate-preview`) reste la version officielle au 27/09/2026 — sortie le 15 octobre 2025, update 4K/1080p en janvier 2026, variante **Veo 3.1 Lite** en public preview depuis le printemps 2026. **Pas de Veo 4** : à Google I/O 2026, Google a cassé la séquence en lançant une nouvelle marque, **Gemini Omni** (avec **Gemini Omni Flash**), au lieu d'un Veo 4. Prix API Veo : NON VÉRIFIÉ dans cette recherche.

C'est un signal produit : Google ne vend plus « un modèle vidéo » mais une **marque multimodale unifiée** (Gemini Omni) où la vidéo est une modalité parmi d'autres — même logique que Nano Banana pour l'image (section 51).

## 51. Nano Banana : ce que c'est vraiment

« Nano Banana » n'est pas un modèle figé : c'est le nom que Google donne à ses **capacités natives de génération d'images dans Gemini**. Historiquement, c'était le nom de code (stealth) de `gemini-2.5-flash-image`. Aujourd'hui c'est une famille de trois modèles, tous fermés, facturés **par image** (pas par token) via la Gemini API.

Points communs à la famille :

- Toutes les images générées portent un watermark **SynthID** (tatouage invisible de Google, guide officiel).
- Les IDs preview (`gemini-3.1-flash-image-preview`, `gemini-3-pro-image-preview`) ont été dépréciés le 25 juin 2026.
- Une seule clé API Gemini achète les trois modèles — la tarification se fait par image générée.

## 52. Nano Banana 2 — le « workhorse »

`gemini-3.1-flash-image` — Preview le 26 février 2026, stable depuis le 25 juin 2026. Le généraliste par défaut.

| Champ | Valeur |
|---|---|
| Sorties | 512px / 1K / 2K / 4K ; ratios standard + extrêmes (1:4, 4:1, 1:8, 8:1) |
| Références | Jusqu'à 14 (10 objets + 4 personnages) |
| Features | System prompts ; **search grounding** ; entrée PDF ; latence de l'ordre de la seconde |
| Contexte | 32K tokens |
| Prix | **$0,067 (1K) / $0,101 (2K) / $0,151 (4K)** par image |

Le search grounding (le modèle peut chercher sur le web avant de générer) est distinctif : tu peux demander « génère le schéma unifilaire d'un onduleur 40 kVA aux normes actuelles » et il ira vérifier avant de dessiner. L'entrée PDF permet de donner un document source à illustrer.

## 53. Nano Banana Pro

`gemini-3-pro-image` — Preview novembre 2025, version premium actuelle. Intégré à **Adobe Firefly et Photoshop**.

| Champ | Valeur |
|---|---|
| Sorties | 1K / 2K / 4K ; ratios standard + wide |
| Références | Jusqu'à 14 images (6 objets + 5 personnages + 3 styles) |
| Features | Raisonnement LLM complet avant génération (« thinking mode ») ; system prompts ; search grounding ; **typographie quasi parfaite** ; watermark SynthID |
| Prix | **$0,134 par image en 1K ou 2K ; $0,24 en 4K** |

Le « thinking mode » avant génération : le modèle raisonne (en texte) sur ce qu'il va dessiner avant de le dessiner — d'où la typographie quasi parfaite, le point faible historique de la génération d'images (textes déformés dans les schémas). Pour des visuels techniques avec du texte (étiquettes de schémas, synoptiques), c'est le modèle adapté.

## 54. Nano Banana 2 Lite

`gemini-3.1-flash-lite-image` — juin 2026, la plus récente de la famille. La plus rapide et la moins chère.

| Champ | Valeur |
|---|---|
| Sortie | **1K uniquement** |
| Usage | Brouillons rapides, génération/édition simple |
| Limites | **Pas** de search grounding ; pas optimisé pour références multiples ni éditions multi-tours |
| Prix | **$0,0336 par image** |

Positionnement clair : le brouillon à 3 centimes avant la version finale à 13 centimes. Pour itérer sur un visuel (10 brouillons + 2 finales Pro 2K), le coût total est d'environ $0,60.

## 55. Prix et usages : que choisir

| Besoin | Modèle | Coût par image |
|---|---|---|
| Brouillon rapide, itération | Nano Banana 2 Lite | $0,0336 (1K) |
| Usage général, PDF en entrée | Nano Banana 2 | $0,067–0,151 |
| Texte dans l'image, qualité max | Nano Banana Pro | $0,134–0,24 |
| Schéma technique avec étiquettes | Nano Banana Pro (typo quasi parfaite) | $0,134 (1K/2K) |

Ordre de grandeur : 1000 images « workhorse » en 1K = $67. C'est un coût marginal pour de la doc illustrée — à comparer au temps humain de production d'un schéma.

## 56. Bilan Partie A : qui choisir pour quoi

| Besoin | Premier choix | Alternative | Pourquoi |
|---|---|---|---|
| Embedding FR, self-host, licence propre | Qwen3-Embedding-8B | BGE-M3 | Apache 2.0, scores MMTEB publiés |
| Embedding léger, 32K, perso uniquement | Jina v5 small | nomic v2 MoE | 239M, Matryoshka, mais NC |
| Embedding multimodal (PDF/images) | Gemini Embedding 2 | Cohere Embed v4 | PDF natif, 100+ langues |
| Rerank qualité max, perso | jina-reranker-v3.5 | jina-reranker-v3 | Listwise 131K, BEIR 63,20 |
| Rerank licence propre | Qwen3-Reranker-0.6B/8B | mxbai-rerank-large-v2 | Apache 2.0 |
| Rerank sans infra | Cohere Rerank 4.0 Fast | Voyage rerank-2.5-lite | API, pas de GPU |
| VLM local léger (légender images) | Qwen3.5-2B (ONNX) | Gemma 4 E2B/E4B | 2B, OCRBench 84,5+, Apache 2.0 |
| VLM open puissant | Qwen3-VL-235B-A22B | Nemotron 3 Nano Omni | Apache 2.0 vs OpenMDW |
| VLM API multimodal | Gemini 3.1 Pro | Grok 4.5 | 1M contexte, T+I+A+V |
| Vidéo API qualité max | Seedance 2.0 | Kling 3.0 | Elo 1269 indépendant |
| Vidéo édition | Runway Aleph | Seedance 2.5 (Video Edit) | Édition vidéo-par-vidéo |
| Vidéo poids ouverts | LTX 2.3 | Wan 2.2 (Apache 2.0) | 20 s vs licence permissive |
| Image génération/édition | Nano Banana 2 | Nano Banana Pro | $0,067 vs typo parfaite |

*Fin de la Partie A. La Partie B commence : le transversal technique.*

---

# PARTIE B — TRANSVERSAL TECHNIQUE

*Le cœur utile : comprendre ce qui se passe sous le capot pour bien choisir, bien dimensionner, bien servir.*

## 57. Le KV cache, c'est quoi

Un LLM génère token par token. Pour prédire le token n°1000, il doit relire les 999 précédents — et le mécanisme d'attention recalcule, à chaque étape, les relations entre le nouveau token et tous les précédents. Sans optimisation, chaque nouveau token coûterait O(n²) : la génération ralentirait à mesure que le texte s'allonge.

Le **KV cache** (Key-Value cache) est l'optimisation qui évite ça : à chaque étape, le modèle calcule les vecteurs **Key** (K) et **Value** (V) du nouveau token pour chaque couche d'attention, et les **stocke**. À l'étape suivante, il ne recalcule que K et V du token courant ; pour tous les précédents, il relit le cache. Coût par token : O(n) au lieu de O(n²). C'est ce qui rend la génération fluide.

En bref : le KV cache, c'est la **mémoire de la conversation** du modèle, stockée en VRAM, couche par couche, token par token. Plus le contexte est long, plus elle grossit — linéairement. C'est le poste n°1 de consommation VRAM en inférence longue, devant les poids du modèle eux-mêmes.

## 58. Pourquoi le KV cache explose avec le contexte long

Le cache grandit avec trois facteurs multipliés :

1. **Le nombre de tokens** (la longueur du contexte) — linéaire, le facteur dominant.
2. **Le nombre de couches** du modèle (80 pour un 70B, 32 pour un 7B).
3. **La taille des vecteurs K/V par couche** (nombre de têtes KV × dimension par tête × 2 pour K+V × octets par élément).

Doubler le contexte = doubler le cache. Passer de 4K à 128K tokens = **×32 sur le cache**. C'est pour ça que les fenêtres de 1M tokens (Gemini 3.1 Pro, Kimi K3, DeepSeek V4) sont un exploit d'infrastructure autant qu'un exploit de modélisation : il faut stocker et relire des dizaines de gigaoctets de cache à chaque token généré.

