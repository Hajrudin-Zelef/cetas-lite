---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-12
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Apple", "Fireworks AI", "Google", "Groq", "Hugging Face", "Meta", "Microsoft", "Mistral", "OpenAI", "OpenRouter", "vLLM", "xAI"]
dates: ["2025-08-26", "2026-05-01", "2026-05-15", "2026-07-08", "2026-08-08", "2026-08-12", "2026-08-18", "2026-09-14", "2026-09-21", "2026-09-27"]
keywords: ["agent", "agents", "apache", "attention", "aws", "bedrock", "benchmarks", "claude", "copilot", "embeddings", "fable 5", "fine-tuning"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [1003, 1083]
sha256: 025a45007967bfc46c4aad4a42739cdadb18ffc2a88a31531851011925cf7107
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | `grok-4.6` (+ variante fast non documentée par ID public) |
| Sortie | **12 août 2026** (rollout partiel dès le 7 août ; Musk confirme l'élargissement le 11 août) |
| Statut | **disponible** — flagship jusqu'au 21 sept. 2026 ; 2ᵉ modèle xAI sur Amazon Bedrock (18/08/2026) |
| Architecture | upgrade de post-training (SFT + RL étendus) **sur la même base V9 que 4.5**, pas une nouvelle fondation ; ~1.5T params selon un brief tiers — **non vérifié officiellement au 27/09/2026** |
| Contexte | **500K tokens** |
| Features | long-running agents (sessions longues, auto-test du travail) ; raisonnement configurable à 4 niveaux : low / medium / high / **xhigh** ; texte+images (JPEG/PNG, max 20 Mo/image, nombre illimité) en entrée ; structured outputs ; Converse API + Chat Completions + Responses sur Bedrock |
| Poids | fermés |
| KV cache | cache d'entrée : `prompt_cache_key` / header `x-grok-conv-id` requis pour des hits fiables (sinon les requêtes se dispersent) ; cache-read $0,50/M (<200K) |
| Benchmarks (vendor) | AA Intelligence Index : 61 (ex æquo GPT-5.6 Sol Max) ; DeepSWE 1.1 : 65,9 % ; AA-Briefcase : 1577 (top) |
| Prix API | **$2 / $6** (<200K tokens), **$4 / $12** (≥200K) par MTok ; cache-read $0,50 / $1,00 ; variante fast au double |
| Déploiement | API xAI, app Grok, Cursor, Grok Build, OpenRouter, Vercel, Cloudflare, Amazon Bedrock (18/08/2026), GitHub Copilot, Microsoft Foundry, Gemini Enterprise Agent Platform |

## 90. Grok 4.7 ✅ (flagship actuel au 27/09/2026)

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | `grok-4.7` |
| Sortie | **21 septembre 2026** (annoncé par Musk le 1er sept. : « dans dix jours » ; sorti effectivement le 21) |
| Statut | **disponible** — flagship actuel (app Grok, Cursor, Grok Build, API xAI) |
| Architecture | **nouvelle base plus large (~2.1T params selon sources médias — non vérifié officiellement)** ; RL plus long pondéré sur tâches de plusieurs heures ; stack de safeguards reconstruite |
| Contexte | **non vérifié officiellement au 27/09/2026** |
| Features | coding + knowledge work long-horizon ; auto-vérification ; meilleure gestion long-contexte ; entraîné nativement sur le harness Grok Bot ; création de documents/présentations |
| Poids | fermés |
| Benchmarks (vendor xAI + analyses tiers) | GDPval : 1695 Elo (vs Claude Fable 5.1 : 1735) ; AA-Briefcase : 1657 (vs Fable 5.1 : 1678) ; gains sur CursorBench 4.0, Terminal-Bench, DeepSWE (chiffres non détaillés dans les sources consultées) ; safeguards : top LatchBio et HackerBench (biosécurité / jailbreak) |
| Prix API | **$2 / $6** par MTok (même prix que 4.6) |
| Déploiement | app Grok, Cursor, Grok Build, API xAI ; Grok Bot (agents always-on, ~418K utilisateurs hebdo au 14/09/2026) |

## 91. Grok Imagine / Aurora : image + vidéo ✅

| Champ | Valeur vérifiée |
|---|---|
| Produit | **Grok Imagine** — suite de génération/édition d'images et de vidéos ; moteur image : **Aurora** ; API image actuelle : `grok-imagine-image` (Aurora) |
| Dates | **Imagine Image 2.0** : **7–8 août 2026** (annoncé par @grok le 8 août ; GA web + iOS/Android) ; API développeur : « coming soon » au 9 sept. 2026 ; migration : `grok-imagine-image-pro` → `grok-imagine-image-quality` le 15 mai 2026 → `grok-imagine-image-2.0` le 2 nov. 2026 (calendrier xAI) ; `grok-2-image-1212` déprécié le 24 fév. 2026 |
| Statut | disponible (Quality Mode sur grok.com/imagine, apps) |
| Features | Image 2.0 : typographie/layout planifiés (petit texte net), édition précise (magic wand, segmentation, suppression de fond avec transparence), multi-références (jusqu'à 5 images), smart resize (9 ratios), templates (produit, headshots, icônes, game assets), consistance de style personnage/lieu/props → génération de « mondes » pour la vidéo |
| Benchmarks | Arena text-to-image : **1320** (#2 mondial) ; Arena image editing : **1439** (#2) — derrière GPT-Image-2 (OpenAI) dans les deux cas |
| Prix API | **non vérifié au 27/09/2026** |
| Déploiement | web (grok.com/imagine), iOS, Android ; API « coming soon » |

## 92. Tableau récapitulatif xAI Grok (sept. 2026)

| Modèle | Sortie | Statut | Contexte | Prix in/out par MTok | Notes |
|---|---|---|---|---|---|
| Grok 3 / 3 Mini | 02/2025 | retirés 15/05/2026 | 131K | — | redirects → 4.3 |
| Grok Code Fast 1 | 26/08/2025 | retiré 15/05/2026 | 256K | $0,20/$1,50 | « sonic » |
| Grok 4.1 / Fast | 11/2025 | Fast retirés 15/05/2026 | 2M (Fast) | — | — |
| Grok 4.3 | 01/05/2026 | dispo | 1M | $1,25/$2,50 | flagship mai–juil |
| Grok 4.5 | 08/07/2026 | dispo | 500K ? | $2/$6 | base V9 |
| Grok 4.6 | 12/08/2026 | dispo | 500K | $2/$6 (<200K) | Bedrock 18/08 |
| **Grok 4.7** | **21/09/2026** | **dispo — flagship** | **non vérifié** | **$2/$6** | ~2.1T (médias) |
| Grok Imagine 2.0 | 08/08/2026 | dispo | — | non vérifié | moteur Aurora |

---

# PARTIE VI — Meta Llama 4 (hors période, mais à documenter)

## 93. Contexte : Llama 4, dernière salve open-weight de Meta

Llama 4 Scout et Maverick sont sortis le **5 avril 2025** — hors période de la mission (fév→sept 2026), et **aucune nouvelle déclinaison Llama 4 n'est sortie fév→sept 2026**. Ils restent néanmoins les derniers poids ouverts majeurs de Meta avant son pivot vers les modèles fermés Muse Spark (§ 34). ⚠️ Réserve transversale : Yann LeCun a admis (Financial Times) que Meta a utilisé différentes versions de Maverick/Scout sur différents benchmarks (« fudged a little bit ») — transparence contestée.

## 94. Llama 4 Scout — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Llama 4 Scout (`meta-llama/Llama-4-Scout-17B-16E` / `-Instruct`) |
| Sortie | **5 avril 2025** (annonce officielle Meta) |
| Statut | disponible (poids ouverts) |
| Architecture | **MoE — 17B paramètres actifs / 109B total**, **16 experts** (routage top-1 + 1 expert partagé), couches denses et MoE alternées ; 48 couches, 40 têtes query / **8 têtes KV (GQA)**, head_dim 128 ; architecture **iRoPE** : couches d'attention sans embeddings positionnels entrelacées (toutes les 4 couches = NoPE), attention chunkée locale 8K + temperature scaling à l'inférence ; QK-norm (RMS) ; rope_theta 500000 |
| Contexte | **10M tokens** (pré/post-entraîné sur 256K, généralisation démontrée via needle-in-haystack sur 10M) |
| Features | nativement multimodal (early fusion texte+vision, encodeur vision dérivé de MetaCLIP ; images/vidéo en entrée), 12 langues en fine-tuning (ar, en, fr, de, hi, id, it, pt, es, tagalog, th, vi), pré-entraîné sur 200 langues ; >30T tokens d'entraînement (conflit : 30T+ blog Meta vs 40T oumi — non tranché) ; training FP8 |
| Fine-tuning | poids ouverts → fine-tuning libre ; recettes LoRA/full/QLoRA existent (ex. oumi-ai/oumi) |
| Poids / licence | **ouverts sous Llama 4 Community License** (custom, PAS Apache 2.0) ; formats BF16, int4/int8 à la volée, FP8 |
| KV cache / inférence | attention chunkée 8K + couches NoPE (réduit le KV cache) ; fits on single H100 (Int4) ; vLLM, Ollama, llama.cpp, transformers ≥ 4.51.0 ; Llama API (OpenAI-compatible) |
| Benchmarks clés | vendor (Meta) : meilleur multimodal de sa classe ; > Gemma 3, Gemini 2.0 Flash-Lite, Mistral 3.1 ; meilleur que tous les Llama précédents (réserve § 93) |
| Prix API | **non vérifié au 27/09/2026** |
| Déploiement | Hugging Face, llama.com, vLLM, Ollama, llama.cpp, partenaires cloud (AWS, Azure, Google Cloud, Groq, Together, Fireworks, DeepInfra, etc.) |

## 95. Llama 4 Maverick — fiche ✅

