---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-1
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cohere", "DeepSeek", "Google", "Huawei", "Hugging Face", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "vLLM"]
dates: ["2026-09-27"]
keywords: ["apache", "arr", "attention", "benchmark", "benchmarks", "claude", "cohere", "deepseek", "embedding", "embeddings", "fine-tuning", "gemini"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [1, 120]
sha256: f61dbe560478ae90148d28dfb7af29357bd2a2442fd81023036124ae51841b26
---

# Encyclopédie des modèles IA — Volume 3
## Modèles spécialisés + transversal technique

*Rédigé le 27 septembre 2026. Snapshot arrêté au 27/09/2026.*
*Volumes 1 et 2 : familles généralistes (GPT, Claude, Gemini, DeepSeek, Qwen, Kimi, Llama, Mistral…).*
*Ce volume : embeddings, rerankers, vision, vidéo, image — puis tout le transversal technique (KV cache, MoE, attention, quantization, fine-tuning, dimensionnement, RAG).*

**Règle d'honnêteté appliquée partout :** les faits issus des recherches portent leur source ; ce qui n'est pas vérifié est marqué `NON VÉRIFIÉ (sept 2026)`. Aucune spec inventée.

---

# PARTIE A — MODÈLES SPÉCIALISÉS

## 1. Mode d'emploi de ce volume

Ce volume a deux jambes. La première (sections 2 à 56) cartographie les modèles spécialisés repérés entre février et septembre 2026 : embeddings, rerankers, vision, vidéo, génération d'image. La seconde (sections 57 à 129) est le transversal technique : tout ce qu'il faut comprendre pour servir des modèles en production et construire un RAG solide.

Si tu ne lis qu'une chose : la Partie B, sections 94 à 102 (choisir les briques d'un RAG) et 103 à 108 (dimensionnement VRAM). C'est le cœur utile pour ton pipeline.

Conventions utilisées dans tout le document :

| Mention | Signification |
|---|---|
| Prix sans mention | Grille officielle du fournisseur (quand la source l'indique) |
| `NON VÉRIFIÉ (sept 2026)` | Info vue dans une source secondaire, non recoupée avec une source primaire |
| `hors période` | Modèle sorti avant février 2026 mais encore leader au snapshot |
| Benchmarks « fournisseur » | Chiffres publiés par le créateur du modèle — à lire avec l'esprit critique de la section 119 |

## 2. Embeddings : à quoi ça sert dans un RAG

Un modèle d'embedding convertit un texte (ou une image, un PDF, un audio) en vecteur de nombres — typiquement 768 à 4096 dimensions. Deux textes de sens proche donnent deux vecteurs proches. C'est la brique de base de la recherche sémantique : tu vectorises ton corpus une fois, tu vectorises la question de l'utilisateur, et tu cherches les vecteurs les plus proches (cosinus, produit scalaire).

Dans ton pipeline actuel, c'est `text-embedding-3-small` d'OpenAI qui fait ce travail. Le choix de l'embedding détermine le plafond de qualité de tout ton RAG : si le retrieval rate le bon chunk, même le meilleur LLM générateur ne pourra pas répondre.

Trois dimensions à comparer entre modèles d'embeddings :

1. **Qualité de retrieval** — mesurée sur MTEB / MMTEB (multilingue), BEIR, MIRACL. Le score seul ne suffit pas (voir section 12).
2. **Contexte d'entrée** — 512 tokens (nomic v2) vs 32K (Jina v5, Qwen3-Embedding) vs 128K (Cohere v4). Si tes chunks font 2000 tokens, un modèle à 512 tronque.
3. **Coût et licence** — API fermée (facturation au million de tokens) vs poids ouverts (Apache 2.0 = usage commercial OK, CC BY-NC = pas de commercial).

## 3. NVIDIA Nemotron-3-Embed-8B-BF16

Le flagship embedding open-weight de NVIDIA, sorti mi-juillet 2026 (jour exact NON VÉRIFIÉ).

| Champ | Valeur |
|---|---|
| Paramètres | ~8B |
| Dimensions | 4096 |
| Contexte | 32 768 tokens |
| Langues | 34 |
| Multimodal | Oui — retrieval texte + visuel (texte sur ViDoRe-V3) |
| Benchmarks (fournisseur, blog HF NVIDIA) | RTEB 78,46 ; MMTEB Retrieval 75,45 ; ViDoRe-V3 (texte) 60,60 |
| Prix API | NON VÉRIFIÉ — modèle pensé pour le self-host |
| Licence | Poids ouverts, **OpenMDW-1.1** |
| Déploiement | PyTorch / vLLM / self-host |

Point d'attention : le score RTEB de 78,46 est annoncé pour la *famille* Nemotron-3-Embed, pas forcément ce checkpoint précis. À vérifier sur la carte Hugging Face avant de le citer comme score du modèle.

Pour ton usage : 8B paramètres, c'est lourd pour un embedding (à comparer aux 239M de Jina v5 small). En self-host il faut une vraie carte GPU. L'intérêt, c'est le retrieval visuel natif : utile si ton corpus contient des schémas, captures d'écrans de CLI, topologies réseau en image.

## 4. Jina Embeddings v5 (small / nano)

Sortie le 23 février 2026, annoncée par Elastic (Jina AI est passée sous pavillon Elastic). C'est la famille « petit mais costaud » du moment.

| Champ | Valeur |
|---|---|
| Variantes | `jina-embeddings-v5-text-small` (239M paramètres) / `jina-embeddings-v5-text-nano` (677M paramètres) |
| Base | Qwen3-0.6B + adaptateurs LoRA par tâche (pré-fusionnés) |
| Dimensions | 1024, réductible par Matryoshka à 32 / 64 / 128 / 256 / 512 / 768 |
| Contexte | 32 768 tokens |
| Langues | Multilingue |
| Benchmarks (fournisseur) | « Best-in-class » MMTEB parmi les modèles de taille comparable ; bat des 7B–14B selon le communiqué |
| Licence | Poids ouverts, **CC BY-NC 4.0** (non commerciale) |
| Déploiement | vLLM, llama.cpp, MLX, API Jina, Elastic Inference Service |

Deux remarques honnêtes :

- Les chiffres 239M (small) / 677M (nano) viennent du communiqué Business Wire et **inversent la logique habituelle** des noms small/nano. Repris tels quels, à vérifier sur la carte HF officielle.
- La licence CC BY-NC 4.0 interdit l'usage commercial. Pour un RAG personnel c'est parfait ; pour un usage pro en entreprise, c'est un frein — voir la section 95 pour les alternatives Apache 2.0.

Le point fort pour toi : 32K de contexte avec un modèle de quelques centaines de millions de paramètres, ça tourne sur CPU ou petit GPU, et Matryoshka permet de stocker des vecteurs à 256 dimensions au lieu de 1024 (4× moins de stockage et de RAM pour l'index).

## 5. Gemini Embedding 2

L'embedding multimodal de Google, disponible via l'API Gemini.

| Champ | Valeur |
|---|---|
| Dimensions | 128 à 3072 (réglable) |
| Contexte texte | 8192 tokens |
| Multimodal | Texte, image, PDF, audio, vidéo |
| Langues | 100+ annoncées |
| Benchmark (indépendant, Zilliz/Milvus, sept 2026) | Needle-in-a-Haystack : **1,000 de précision** à 1K/4K/8K/16K/32K — seul modèle testé sur toute la plage, seul à atteindre 32K |
| Prix (source OpenRouter, secondaire) | Texte $0,20/M tokens ; image/fichier $0,45/M ; audio $6,50/M ; vidéo $12/M ; batch −50 % — NON VÉRIFIÉ contre la grille officielle Google |
| Licence | Fermé, API uniquement |

Incertitude sur la date : annonce le 10 mars 2026 selon une source, sortie repérée le 20 mai 2026 selon une autre. Le test Zilliz est le seul benchmark **indépendant** du lot : une précision parfaite sur toute la plage 1K–32K, c'est remarquable, mais c'est un seul test (needle-in-a-haystack), pas un MTEB complet.

Pour ton usage : l'entrée PDF native est intéressante (tes datasheets Huawei, manuels Kyocera), mais le modèle est fermé et le prix non confirmé officiellement.

## 6. Voyage 4 (large / standard / lite)

Sorti en janvier 2026 — **hors borne basse** de la fenêtre de recherche (février 2026), inclus car leader du snapshot au 27/09/2026.

| Champ | Valeur |
|---|---|
| Variantes | voyage-4-large / voyage-4 / voyage-4-lite |
| Dimensions | 256 / 512 / 1024 / 2048 |
| Contexte | 32K |
| Benchmarks | Revendications Voyage sur RTEB **contestées** : problème d'accès au split privé du benchmark signalé par des tiers — à ne citer que comme revendication fournisseur |
| Prix (docs officielles Voyage) | large $0,12/M ; standard $0,06/M ; lite $0,02/M tokens |
| Bonus | 200M tokens gratuits annoncés |
| Licence | Fermé, API uniquement |

Voyage AI s'est spécialisé dans les embeddings « retrieval-first », avec une bonne réputation sur le code et le multilingue. Mais l'alerte sur les benchmarks RTEB est sérieuse : quand un tiers signale un problème d'accès au split privé d'un benchmark, les scores deviennent suspects (risque de contamination ou de sur-optimisation). Esprit critique, section 119.

## 7. Cohere Embed v4

Sorti en avril 2025 — **hors période**, mais encore un leader actif au snapshot.

