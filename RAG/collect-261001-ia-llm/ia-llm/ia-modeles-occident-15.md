---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-15
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Baseten", "Google", "Hugging Face", "Nvidia", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-04-02", "2026-06-19", "2026-09-27"]
keywords: ["agent", "alignment", "apache", "attention", "benchmarks", "blackwell", "consumer", "embedding", "embeddings", "fine-tuning", "gemini", "gguf"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [1254, 1326]
sha256: 38aff76b62f8738b41c1ddf4bf7987cdc9623c3614a1f09ce7ae5e0af7919911
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Gemma 4 — 4 variantes : **Effective 2B (E2B), Effective 4B (E4B), 26B-A4B MoE, 31B Dense** |
| Sortie | **2 avril 2026** (blog officiel Google) |
| Statut | disponible (poids ouverts) |
| Architecture | • **26B-A4B** : MoE — **26B total / 3,8B actifs**, 128 experts (8 + 1 partagé par token) • **31B** : **dense**, tous paramètres actifs • **E4B** : dense **7,9B total / 4,5B effectifs**, Per-Layer Embeddings (PLE) • **E2B** : dense **5,1B total / 2,3B effectifs**. Vocabulaire SentencePiece 262K ; 140+ langues |
| Attention | secondaire : sliding-window attention (fenêtre 1024, ratio 5 SWA : 1 globale) — **non vérifié par source officielle** |
| Contexte | **256K** (12B/26B/31B) · **128K** (E2B/E4B) |
| Features | nativement multimodal (images/vidéo à résolution variable sur toute la famille ; **audio natif sur E2B/E4B** : ASR, transcription) ; agent-native : function calling natif, JSON structuré strict, instructions système natives ; reasoning configurable (mode thinking) ; **MTP** (multi-token prediction) : têtes de draft pour décodage spéculatif sans modèle draft séparé ; construit sur la même recherche que **Gemini 3** |
| Fine-tuning | ✅ **supporté et encouragé officiellement** : « sized to run and fine-tune efficiently » ; support day-one Hugging Face (Transformers, **TRL**), Unsloth, Vertex AI, Colab, TPU. LoRA via ces frameworks (standard, non détaillé officiellement) |
| Poids / licence | **ouverts sous Apache 2.0** (**rupture majeure** : Gemma 1–3 utilisaient les Gemma Terms of Use custom). Checkpoints officiels : BF16, **NVFP4** (Blackwell, quasi-identique au 8-bit), **GGUF Q4_0 QAT officiels Google** |
| KV cache / inférence | SWA 5:1 (KV cache à croissance lente, secondaire) ; MTP speculative decoding ; 26B-A4B : le plus rapide (~110 tok/s gen en test local secondaire) ; 31B BF16 tient sur 1×H100 80GB ; 26B MoE quantifié sur GPU consumer 24GB |
| Benchmarks clés | officiel : 31B = **#3 des modèles ouverts** sur Arena AI text leaderboard ; 26B = **#6** ; « outcompete des modèles 20× leur taille ». Tech report (secondaire, 31B IT thinking) : Arena **1452**, AIME 2026 **89,2**, GPQA Diamond **84,3**, MMLU Pro **85,2**, LiveCodeBench v6 **80,0**, τ2-bench (tool use) **86,4**. Table Arena au 19/06/2026 : 31B Elo 1451±8 (#43), 26B-A4B 1438±8 (#61) |
| Prix API | pas d'API payante dédiée ; accessible via **Google AI Studio** (31B, 26B) et **Gemini API** (`gemma-4-26b-a4b-it`, `gemma-4-31b-it` depuis le 2 avril). Tarifs : **non vérifiés au 27/09/2026** |
| Déploiement | Hugging Face (`google/gemma-4-*`), Ollama, Kaggle, LM Studio, Docker, vLLM, llama.cpp, SGLang, MLX, NVIDIA NIM, NeMo, LiteRT-LM, Baseten, Vertex AI / Cloud Run / GKE / TPU ; AI Edge Gallery (E4B/E2B) ; AICore Developer Preview (Android, forward-compat Gemini Nano 4) |

## 112. Gemma 4 E2B et E4B : les petits multimodaux on-device

- **E2B** : dense, 5,1B total / 2,3B effectifs, contexte 128K.
- **E4B** : dense, 7,9B total / 4,5B effectifs (Per-Layer Embeddings), contexte 128K.
- Les deux : **audio natif** (ASR, transcription), images/vidéo à résolution variable, pensés on-device (AI Edge Gallery, AICore Android). Licence **Apache 2.0**. C'est la réponse de Google au besoin « petit modèle multimodal local » — là où Gemini 3.x reste cloud-only fermé.

## 113. Gemma 4 26B-A4B (MoE) et 31B (dense) : les têtes de série

- **26B-A4B** : MoE 26B total / 3,8B actifs (128 experts, 8 + 1 partagé par token), contexte 256K. Le plus rapide de la famille (~110 tok/s en test local secondaire) ; tient quantifié sur GPU consumer 24GB.
- **31B** : dense 31B, contexte 256K, tient en BF16 sur 1×H100 80GB. #3 des modèles ouverts sur Arena AI text (officiel).
- Tous deux : Apache 2.0, MTP speculative decoding, checkpoints NVFP4 officiels et GGUF Q4_0 QAT officiels Google.

## 114. Gemma 4 12B — déclinaison juin 2026 ⚠️

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Gemma 4 12B (`google/gemma-4-12b`, `google/gemma-4-12b-it`) |
| Sortie | 3–4 juin 2026 (**sources secondaires uniquement** — pas d'annonce officielle retrouvée) |
| Statut | disponible (HF, Ollama incl. `gemma4:12b-mlx`, LM Studio) |
| Architecture | **encoder-free** : décodeur unifié sans encodeur vision/audio séparé (~11,95B params, 48 couches, vocab 262K) ; projections linéaires directes vers l'espace d'embedding du LLM ; texte+image+audio+vidéo |
| Contexte | 256K (secondaire) |
| Features | ASR natif de qualité transcription (template de prompt ASR fourni), multimodal unifié ; tourne sur laptop ~16GB (quantifié, texte) |
| Poids / licence | Apache 2.0 (secondaire) |
| Benchmarks | secondaire : AIME 2026 77,5, GPQA Diamond 78,8, MMLU Pro 77,2, LiveCodeBench v6 72,0 ; proche du 26B MoE à <50 % de l'empreinte mémoire (vendor/secondaire) |

## 115. Tableau Gemma 4

| Variante | Sortie | Arch. | Contexte | Licence | Notes |
|---|---|---|---|---|---|
| Gemma 4 E2B | 02/04/2026 | dense 5,1B (2,3B eff.) | 128K | Apache 2.0 | audio natif, on-device |
| Gemma 4 E4B | 02/04/2026 | dense 7,9B (4,5B eff.) | 128K | Apache 2.0 | audio natif, on-device |
| Gemma 4 26B-A4B | 02/04/2026 | MoE 26B/3,8B | 256K | Apache 2.0 | #6 ouverts Arena ; le plus rapide |
| Gemma 4 31B | 02/04/2026 | dense 31B | 256K | Apache 2.0 | #3 ouverts Arena |
| Gemma 4 12B | 06/2026 (secondaires) | encoder-free ~12B | 256K | Apache 2.0 | multimodal unifié |

---

# PARTIE IX — IBM Granite (open-weight entreprise)

## 116. IBM Granite 4.1 (3B / 8B / 30B) — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Granite-4.1-3B / 4.1-8B / 4.1-30B (variantes `-Base` et `-Instruct`) |
| Sortie | **29 avril 2026** |
| Statut | disponible (Hugging Face) |
| Architecture | **dense** decoder-only, 3B / 8B / 30B params |
| Contexte | « long-context » (taille exacte **non vérifiée au 27/09/2026**) |
| Features | Instruct : SFT + RL alignment ; **tool calling natif**, instruction following, chat ; **FIM** (fill-in-the-middle) code completions ; code, RAG, multilingue (12 langues : en, de, es, fr, ja, pt, ar, cz, it, ko, nl, zh) |
| Fine-tuning | poids ouverts → fine-tuning libre ; IBM encourage le fine-tuning (docs Granite) |
| Poids / licence | **ouverts, Apache 2.0** ; HF : collection `ibm-granite/granite-41-language-models` ; repo GitHub `ibm-granite/granite-4.1-language-models` ; GGUF |
| KV cache / inférence | **non vérifié au 27/09/2026** ; serveurs : vLLM, Ollama, llama.cpp |
| Benchmarks clés | **non vérifié via sources primaires** (model cards évoquent des evals sans chiffres exploitables ici) |
| Prix API | **non vérifié** (disponible via IBM watsonx — tarifs non vérifiés) |

## 117. IBM Granite 4.2 (3B / 8B / 30B) — reasoning ✅

