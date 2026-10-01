---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-23
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cohere", "DeepSeek", "Google", "Hugging Face", "Meta", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Z.ai", "vLLM", "xAI"]
dates: ["2026-09-24", "2026-09-27"]
keywords: ["apache", "arr", "attribution", "claude", "cohere", "deepseek", "distribution", "embedding", "embeddings", "gemini", "gguf", "glm"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [2309, 2407]
sha256: 1dd5b570d5479593c711dd83654c0dba4c568f4a8a64b737d7a9664c73e60422
---

# Encyclopédie des modèles IA — Volume 3

| Licence | Vrai open-source (OSI) ? | Usage commercial | Exemples (ce volume) | Piège |
|---|---|---|---|---|
| **Apache 2.0** | Oui | Oui, libre | Qwen3-Embedding, Qwen3-Reranker, Qwen3-VL, Qwen3.5-Omni, Gemma 4, BGE-M3, mxbai (?), Granite | Aucun — la plus sûre |
| **MIT** | Oui | Oui, libre | DeepSeek V4, MiMo/Ling | Aucun — la plus permissive |
| **CC BY-NC 4.0** | Non | **Interdit** | Jina v5, jina-reranker-v3/v3.5 | « Non commercial » inclut l'usage interne d'entreprise selon l'interprétation stricte |
| **Llama Community License** | Non (source-available) | Oui sous 700M MAU | Llama 5 (5V) | Au-delà de 700M d'utilisateurs mensuels : licence Meta obligatoire |
| **NVIDIA Open Model License** | Non | Oui (conditions) | Nemotron 3 Nano Omni, Nemotron-3-Embed (OpenMDW-1.1) | Lire les conditions d'attribution et de redistribution |
| **Licence custom** | Non | Selon le texte | Kimi K3, Mistral Modified MIT | Chaque texte est différent — à lire |
| **Plafond d'ARR** | Non | Oui sous 10M$ d'ARR | LTX 2.3 | Au-delà : renégocier |
| **Propriétaire / fermé** | Non | Via API (CGU) | Gemini, Voyage, Cohere, OpenAI, xAI | Pas de poids, dépendance fournisseur |

Règles pratiques :

1. **Usage personnel** : tout est permis, même le NC. Ton RAG perso peut utiliser Jina sans souci.
2. **Usage entreprise** : Apache 2.0 / MIT par défaut. Le CC BY-NC est exclu sauf avis juridique. Les licences custom se lisent ligne par ligne.
3. **Redistribution** (revendre un produit basé sur le modèle) : Apache 2.0 exige la conservation des notices ; MIT quasi rien ; les customs varient.
4. **Les poids ≠ le code.** Une licence de poids peut restreindre ce que la licence du code autorise — vérifie les deux quand tu déploies un serveur (vLLM est Apache 2.0, mais les poids que tu sers ont leur propre licence).

## 145. Le décodage spéculatif : générer plus vite sans perdre en qualité

La génération est séquentielle : 1 token par passe. Le **décodage spéculatif** casse ça : un petit modèle « draft » propose K tokens d'avance, le grand modèle les **vérifie en une seule passe** (le prefill parallèle est peu coûteux). Les tokens acceptés sont gardés, les refusés sont régénérés.

| Méthode | Principe | Où on la voit (2026) |
|---|---|---|
| Draft model classique | Un petit modèle (ex : 1B) propose, le grand vérifie | vLLM, SGLang (config manuelle) |
| **EAGLE** | Le draft est une tête légère entraînée sur les états du grand modèle (pas un modèle séparé) | Mistral Medium 3.5 (variante EAGLE citée), vLLM |
| **MTP** (Multi-Token Prediction) | Le modèle prédit nativement plusieurs tokens par passe (tête dédiée) | Qwen4-Exp (tête MTP 4B), Qwen3.8-27B (tête MTP incluse → spec decoding par défaut sur Ollama) |

Gain typique : ×1,5 à ×3 sur la vitesse de génération, **sans perte de qualité** (la vérification garantit la même distribution que le grand modèle seul). Condition : le draft doit être bon — si le petit modèle propose n'importe quoi, tout est refusé et on perd du temps.

Pour ton RAG : le décodage spéculatif accélère la phase de génération (les 800 tokens de réponse), pas le prefill (les 20K tokens de chunks — déjà parallélisé). Sur des réponses longues et structurées (tes formats Diagnostic → Vérification), le gain est maximal car le texte est prévisible.

## 146. Index des modèles cités dans ce volume

Rappel alphabétique — statut au 27/09/2026, voir les sections détaillées pour les réserves.

| Modèle | Section | Type | Licence / accès |
|---|---|---|---|
| Alibaba Wan 2.2 → 2.7 | 43 | Vidéo | Apache 2.0 (≤ 2.2) / fermé (au-delà) |
| BGE-M3 / bge-reranker-v2-m3 | 9, 20 | Embedding / rerank | Apache 2.0 |
| Claude Opus 4.8 | 34 | Vision (mention) | Fermé — source secondaire uniquement |
| Cohere Embed v4 | 7 | Embedding | Fermé (API) |
| Cohere Rerank 4.0 Pro/Fast | 15 | Rerank | Fermé (API) |
| DeepSeek V4 / V4-Pro / Flash | 64, 75 | LLM MoE | MIT (poids ouverts) |
| Gemini 3.1 Pro | 27 | VLM | Fermé (API) |
| Gemini Embedding 2 | 5 | Embedding | Fermé (API) |
| Gemini Omni Flash | 50 | Vidéo | Fermé |
| Gemma 4 31B | 26 | VLM | Apache 2.0 |
| GLM-4.6V | 34 | Vision (mention) | Poids téléchargeables |
| GPT-5.6 Sol / GPT-6 Sol | 34 | Vision (mention) | Fermé — non confirmé par OpenAI |
| Grok 4.5 | 33 | VLM | Fermé (API) |
| Grok Imagine 1.5 | 48 | Vidéo (mention) | Fermé |
| Jina Embeddings v5 | 4 | Embedding | CC BY-NC 4.0 |
| jina-reranker-v3 / v3.5 | 14, 17 | Rerank | CC BY-NC |
| Kimi K2 / K2.6 / K2.7 | 75 | LLM MoE | Modified MIT / custom selon version |
| Kimi K3 | 25 | VLM MoE | Licence custom (pas OSI) |
| Kling 3.0 | 41 | Vidéo | Fermé (API) |
| Llama 4 (Maverick) | 65 | LLM | Llama 4 Community License |
| Llama 5 / 5V | 29 | VLM | Community License — source unique |
| LTX 2.3 | 47 | Vidéo | Poids ouverts, plafond 10M$ ARR |
| Luma Ray3.2 | 44 | Vidéo | Fermé (API) |
| MiMo-V2-Omni | 31 | VLM | Propriétaire — source unique |
| MiniMax Hailuo 2.3 | 45 | Vidéo | Fermé (API) |
| Mistral Medium 3.5 | 100 | LLM | Modified MIT |
| mxbai-rerank-large-v2 | 19 | Rerank | Apache 2.0 (à confirmer) |
| Nano Banana 2 / Pro / Lite | 52–54 | Image | Fermé (API) |
| Nemotron 3 Nano Omni | 24 | VLM MoE | NVIDIA Open Model License |
| Nemotron-3-Embed-8B | 3 | Embedding | OpenMDW-1.1 |
| nomic-embed-text-v2-moe | 9 | Embedding | Apache 2.0 |
| PixVerse V6 | 46 | Vidéo | Fermé (API) |
| Qwen3-Embedding-8B | 8 | Embedding | Apache 2.0 |
| Qwen3-Reranker (0.6B/4B/8B) | 18 | Rerank | Apache 2.0 |
| Qwen3-VL-235B-A22B | 30 | VLM | Apache 2.0 |
| Qwen3.5-Omni (0.8B/2B) | 32 | VLM | Apache 2.0 |
| Qwen3.6-35B-A3B | 72–73 | LLM MoE | Apache 2.0 |
| Qwen3.8-27B | 84, 100 | LLM dense | Apache 2.0 |
| Qwen4-Exp | 28 | VLM MoE | Poids ouverts, licence non vérifiée |
| Runway Gen-4.5 / Aleph | 42 | Vidéo | Fermé (API) |
| Seedance 2.0 / 2.5 | 39–40 | Vidéo | Fermé (API) |
| Snowflake arctic-embed-l-v2.0 | 9 | Embedding | Apache 2.0 |
| Sora 2 / Sora 2 Pro | 38 | Vidéo | **API fermée le 24/09/2026** |
| Veo 3.1 | 50 | Vidéo | Fermé (API) |
| Vidu Q3 Pro | 48 | Vidéo (mention) | Fermé |

*Fin du volume 3 — 146 sections.*

## 147. Travaux pratiques : 5 exercices pour ancrer le transversal

Pour passer de la lecture à la maîtrise — chaque exercice prend 15 à 45 minutes, avec du matériel gratuit.

**TP 1 — Calculer un KV cache à la main (15 min).**
Prends la fiche d'un modèle que tu envisages (ex : Qwen3.8-27B : 48 couches, GQA, head_dim 128 — vérifie les chiffres sur la carte Hugging Face). Applique la formule de la section 59 pour 8K, 32K et 128K tokens. Puis additionne les poids en Q4_K_M (18 Go) et déduis la taille max du batch sur ta carte. Objectif : que le calcul devienne un réflexe avant tout choix de modèle.

**TP 2 — Comparer deux quants du même modèle (30 min).**
Via Ollama : `ollama pull qwen2.5:7b-instruct-q4_K_M` puis la variante `q8_0` (ou via llama.cpp deux GGUF). Pose les mêmes 10 questions techniques en français aux deux, note les différences (précision des chiffres, qualité du français, hallucinations). Objectif : calibrer ton œil sur ce que la quantization change *vraiment* — et ce qu'elle ne change pas.

**TP 3 — Mesurer l'effet du chunking (45 min).**
Prends un de tes guides (ex : le guide onduleurs, 1602 lignes). Indexe-le deux fois : (a) chunks de 1000 tokens à taille fixe, (b) chunks par sections `##`. Pose 20 questions, compte combien de fois la bonne section arrive dans le top-5 pour chaque méthode. Objectif : constater sur TES données l'écart entre les deux stratégies (section 109).

