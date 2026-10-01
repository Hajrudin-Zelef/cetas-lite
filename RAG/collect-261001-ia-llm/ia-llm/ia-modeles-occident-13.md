---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-13
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Hugging Face", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "SGLang", "United States", "vLLM"]
dates: ["2025-04-05", "2026-05-22", "2026-09-14", "2026-09-16", "2026-09-27"]
keywords: ["agent", "agents", "apache", "attention", "benchmarks", "blackwell", "claude", "compute", "datacenter", "deepseek", "distillation", "embeddings"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [1084, 1162]
sha256: 0d13e3ca952444b89b80a9c28123000617a52c3b0ca380380299f7a3ad976df7
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Llama 4 Maverick (`meta-llama/Llama-4-Maverick-17B-128E` / `-Instruct` / `-FP8`) |
| Sortie | **5 avril 2025** (annonce officielle Meta) |
| Statut | disponible (poids ouverts) |
| Architecture | **MoE — 17B paramètres actifs / 400B total**, **128 experts routés + 1 expert partagé** (top-1 par token), couches denses/MoE alternées ; iRoPE ; config couches/hidden exacte : **non vérifiée** (source secondaire oumi : 80 couches, 8192 hidden, 64 Q / 8 KV — non confirmé) |
| Contexte | **1M tokens** |
| Features | nativement multimodal (texte+images ; jusqu'à 5 images en entrée), 12 langues, distillation (codistillation) depuis Behemoth |
| Fine-tuning | poids ouverts ; recettes LoRA/full existantes |
| Poids / licence | **ouverts sous Llama 4 Community License** (custom) ; BF16, FP8 |
| KV cache / inférence | tourne sur un seul hôte NVIDIA H100 DGX (8×H100) ; vLLM, Ollama, transformers ≥ 4.51.0 |
| Benchmarks clés | vendor : bat GPT-4o et Gemini 2.0 Flash sur benchmarks larges ; comparable à DeepSeek v3/v3.1 en raisonnement/coding avec <50 % des paramètres actifs ; version chat expérimentale : ELO 1417 sur LMArena (réserve § 93) |
| Prix API | **non vérifié au 27/09/2026** |
| Déploiement | HF, vLLM, Ollama, cloud partners |

## 96. Llama 4 Behemoth : jamais publié ❌

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Llama 4 Behemoth |
| Sortie | **non trouvé comme release** — preview uniquement le 5 avril 2025, « still training » |
| Statut | ❌ **non publié** (teacher model interne). Vérifié le 14/09/2026 : aucun checkpoint Behemoth dans l'org HF meta-llama, aucune model card dédiée |
| Architecture | **MoE — 288B paramètres actifs / ~2T total**, 16 experts, multimodal (source : blog Meta officiel) |
| Contexte | **non vérifié au 27/09/2026** |
| Benchmarks | vendor (Meta, modèle encore en training au lancement) : > GPT-4.5, Claude Sonnet 3.7, Gemini 2.0 Pro sur MATH-500 et GPQA Diamond |
| Note | une source secondaire (codersera, ferme de contenu) affirme que Meta aurait « shelved » Behemoth et pivoté vers Muse Spark / Muse Glimmer — **non vérifié, source unique de faible qualité, traité avec scepticisme** |

## 97. Tableau Llama 4

| Modèle | Sortie | Statut | Contexte | Arch. | Licence | Notes |
|---|---|---|---|---|---|---|
| Llama 4 Scout | 05/04/2025 | dispo | 10M | MoE 17B/109B | Llama 4 Community (custom) | multimodal natif |
| Llama 4 Maverick | 05/04/2025 | dispo | 1M | MoE 17B/400B | Llama 4 Community (custom) | codistillé depuis Behemoth |
| Llama 4 Behemoth | preview 05/04/2025 | **jamais publié** | non vérifié | MoE 288B/~2T | — | teacher interne |

---

# PARTIE VII — Mistral AI (France)

## 98. Mistral : un Européen dans la course 2026

Calendrier officiel des sorties Mistral fév→sept 2026 (changelog ai-pricelog + docs officielles) : **4 fév** : Voxtral Mini Transcribe 2 + Realtime · **12 mars** : Mistral Moderation 2603 / Moderation 2 · **16 mars** : **Mistral Small 4** · **23 mars** : **Voxtral TTS** + **Leanstral** · **28 avril** : **Mistral Medium 3.5** (poids ouverts) · **26–28 mai** : **Vibe** (Le Chat rebrandé en agent unifié) · **8 juillet** : **Robostral Navigate** (8B, robotique). Côté infra : **Mistral Compute** (18 000 GPU NVIDIA Blackwell, datacenter 44 MW près de Paris — secondaire) ; acquisitions : **Koyeb** (fév. 2026, infra IA), **Emmi AI** (mai 2026, ~€300M, modèles neuraux de substitution pour simulation physique industrielle) — secondaires.

## 99. Mistral Small 4 — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | Mistral Small 4 (`mistral-small-2603`) |
| Sortie | **16 mars 2026** (annoncé à NVIDIA GTC 2026) |
| Statut | disponible (API Mistral, Hugging Face, NVIDIA NIM) |
| Architecture | **MoE — 119B total / ~6–6,5B actifs par token (8B avec embeddings)** ; 128 experts, 4 actifs par token ; couches/attention : **non vérifié au 27/09/2026** |
| Contexte | **256K tokens** (262 144) |
| Features | modèle « unifié » : remplace 3 lignées (Magistral=raisonnement, Pixtral=vision, Devstral=coding) ; raisonnement configurable par requête (`reasoning_effort="none"`/`"high"`) ; multimodal texte+image en entrée ; multilingue |
| Fine-tuning | poids ouverts (Apache 2.0) → fine-tuning libre ; support API : **non vérifié** |
| Poids / licence | **ouverts, Apache 2.0** ; variante officielle NVFP4 : `mistralai/Mistral-Small-4-119B-2603-NVFP4` (collab Mistral × vLLM × Red Hat, llm-compressor) |
| KV cache / inférence | 40 % de latence en moins et 3× throughput vs Mistral Small 3 ; self-host : 4×H100 / 2×H200 / 1×DGX B200 min (secondaire) ; vLLM, SGLang, llama.cpp, Ollama ; NVIDIA NIM day-0 |
| Benchmarks clés | vendor/secondaires : GPQA Diamond **71,2 %** (vs GPT-4o-mini 40,2 %), MMLU-Pro 78,0 % ; bat GPT-OSS 120B sur LiveCodeBench avec 20 % de sortie en moins (vendor) |
| Prix API | **non vérifié au 27/09/2026** |
| Déploiement | Mistral API (`mistral-small-2603`), HF, NVIDIA NIM, vLLM, llama.cpp ; intégré à Firefox Smart Window (Mozilla, 16/09/2026, US/CA) |

## 100. Mistral Medium 3.5 — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID | Mistral Medium 3.5 (`mistral-medium-3-5-26-04` / `mistral-medium-3.5`) |
| Sortie | **28 avril 2026** (model card ; news Mistral datée du 22/05/2026 « powering remote coding agents in Vibe ») |
| Statut | disponible (poids ouverts + API) |
| Architecture | **dense — 128B paramètres** (tous actifs) ; ~88 couches, 8 têtes KV, head_dim 128, standard « Mistral 3 family » (secondaire) ; tokenizer Tekken |
| Contexte | **256K tokens** |
| Features | consolide Medium 3.1 (instruct) + Magistral (raisonnement) + Devstral 2 (coding) en un seul modèle avec toggle `reasoning_effort` ; multimodal texte+image ; function calling natif, agents, structured outputs, prefix, FIM, OCR, batching |
| Fine-tuning | poids ouverts → fine-tuning libre ; API fine-tuning : **non vérifié** |
| Poids / licence | **ouverts sous Modified MIT** (open weights avec carve-out pour les très grosses entreprises à hauts revenus) ; HF : `mistralai/Mistral-Medium-3.5-128B` (gated), GGUF publics (bartowski) ; variante **EAGLE** (décodage spéculatif) |
| KV cache / inférence | self-host : dès 4 GPUs (secondaire) ; vLLM tensor-parallel 8 recommandé (secondaire) |
| Benchmarks clés | vendor : SWE-bench Verified **77,6 %** ; τ³-Telecom 91,4 (secondaire) |
| Prix API | **$1,50 / $7,50** par MTok — plusieurs sources secondaires concordantes, mais prix officiel non vérifié directement → **à confirmer sur mistral.ai/pricing** |
| Déploiement | Mistral API, Hugging Face, moteur des agents distants Vibe |

## 101. Robostral Navigate — fiche ✅ (robotique, poids non publiés)

