---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-16
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Hugging Face", "Nvidia", "OpenAI", "OpenRouter", "Poolside", "SGLang", "vLLM"]
dates: ["2025-10-28", "2026-01-01", "2026-07-02", "2026-07-09", "2026-07-21", "2026-07-28", "2026-09-27"]
keywords: ["agent", "agentic", "apache", "benchmarks", "datacenter", "embeddings", "fine-tuning", "fp8", "gguf", "gpu", "gqa", "kv cache"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [1327, 1406]
sha256: e8e75b2899183ec0b1ec00618ffd96a2382f787e4f18a79915e50309ace95836
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Granite-4.2-3B / 4.2-8B / 4.2-30B (`ibm-granite/granite-4.2-*-instruct`) |
| Sortie | **25 août 2026** |
| Statut | disponible (Hugging Face, Ollama, GitHub — le jour même) |
| Architecture | **dense** decoder-only reasoning ; 3B / 8B / 30B. 3B : 3,66B params, 40 couches, hidden 2560, **GQA 40:8**, vocab 100k, rope_theta 1e7, embeddings non liés (secondaire) |
| Contexte | phase de pré-training long-contexte ciblant **512K** ; limites servies par carte à vérifier |
| Features | **thinking / non-thinking toggle** (+ mode low-effort pour prompts faciles) ; **tool calling natif** format OpenAI function-calling ; **agentic RL** (8B et 30B uniquement) entraîné dans de vrais sandboxes (coding, terminal, web search) ; foundational RL sur les 3 tailles (math, science, code, tool use) + RLHF ; **CodeAlchemy** : ~1T tokens de code synthétique ; 5 phases de pré-training (~15T tokens from scratch) documentées ; poids signés cryptographiquement (positionnement entreprise/régulé) |
| Fine-tuning | poids ouverts → fine-tuning libre |
| Poids / licence | **ouverts, Apache 2.0** ; GGUF officiels ; Ollama : `granite4.2:8b` |
| KV cache / inférence | vLLM (`--enable-auto-tool-choice --tool-call-parser granite`), SGLang, Ollama, LiteRT-LM (on-device) ; 3B laptop, 8B sweet spot agent single-GPU (24–32GB), 30B multi-GPU/DGX Spark |
| Benchmarks clés | positionné « enterprise agent workflows » (math, science, tool use, planification multi-étapes) plutôt que headline SWE-bench — **pas de sweep SWE-bench indépendant vérifié** |
| Prix API | **non vérifié au 27/09/2026** |

## 118. Positionnement Granite : l'open-weight « entreprise régulée »

Granite ne joue pas la course aux benchmarks grand public : IBM vise les DSI et les secteurs régulés — poids **signés cryptographiquement**, 5 phases de pré-training documentées (~15T tokens from scratch), tool calling au format OpenAI, agentic RL en sandbox réel, licence Apache 2.0 sans ambiguïté. C'est l'open-weight « boring but safe » : 3B pour le laptop, 8B pour l'agent single-GPU, 30B pour le datacenter.

## 119. Autres Granite notables (fiches courtes)

- **Granite 4.0 / 4.0-H 350M** (28/10/2025 — hors période) : instruct denses/hybrides 352M/340M, Apache 2.0. (secondaire)
- **Granite 3.3 8B Math PRM v2** (01/01/2026 — hors période) : process reward model génératif pour supervision math, Apache 2.0, SOTA best-of-N dans sa classe (model card via strix-halo).
- **Modèle code Granite dédié en 2026 : NON TROUVÉ.** Aucune release « Granite 4.x Code » identifiée après 2 requêtes (« Granite code 2026 », « granite-4 code instruct »). Le code est couvert par : les instruct généralistes (FIM code completions), CodeAlchemy (~1T tokens de code synthétique dans Granite 4.2), et l'agentic RL code/terminal. Les lignées code historiques (granite-3.x-code) n'ont pas de successeur 2026 vérifié.

---

# PARTIE X — Poolside (Laguna, modèles de code open-weight)

## 120. Poolside : la « Model Factory » du code

Poolside (startup fondée par Jason Warner, ex-CTO GitHub ; co-CEO : Jason Warner & Eiso Kant) industrialise l'entraînement : cadence de sortie annoncée ~5 semaines (« Model Factory » industrialisée). Spécialité : **modèles de code agentic open-weight**, optimiseur **Muon** (pas AdamW), KV cache quantifié **FP8**, SWA à per-head gating. Levées de fonds citées : chiffres contradictoires entre sources ($626M / $2Mds / $3Mds — **à traiter avec prudence**). **Aucun nouveau modèle après S 2.1 (21/07/2026)** trouvé au 27/09/2026 — seulement un checkpoint mis à jour d'S 2.1 en août 2026.

## 121. Laguna XS.2 — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Sortie | **28 avril 2026** (premiers modèles publics de Poolside) |
| Statut | disponible (supersedé par XS 2.1 le 02/07/2026) ; tier API gratuit fermé le 09/07/2026 |
| Architecture | **MoE : 33B totaux / 3B actifs par token**, 256 experts + 1 expert partagé, 40 couches (30 SWA avec gating per-head + 10 globales) ; entraîné from scratch sur **30T tokens** ; optimiseur **Muon** ; pré-entraînement commencé ~5 semaines avant la sortie |
| Contexte | 131 072 tokens entrée / 8 192 max sortie |
| Features | raisonnement activé par défaut (`enable_thinking`) ; tool use agentique ; agent terminal `pool` (protocole ACP, open-source) ; sandbox cloud `Shimmer` |
| Fine-tuning | **non vérifié au 27/09/2026** |
| Poids | **ouverts — Apache 2.0** (Hugging Face) ; tourne sur un seul GPU / Mac 36 Go |
| KV cache / optimisations | KV cache quantifié **FP8** ; SWA per-head gating choisie pour réduire la mémoire KV ; DFlash draft models pour speculative decoding |
| Benchmarks (vendor, avril 2026) | SWE-Bench Verified : 68,2 % (table intro) — une table ultérieure cite 64 %/65,4 % ; SWE-Bench Pro : 44,5 % ; Terminal-Bench 2.0 : 30,1 % |
| Prix API | $0/token pendant la période gratuite limitée (API Poolside + OpenRouter) |
| Déploiement | Hugging Face, Ollama, API Poolside, OpenRouter ; intégrations : Agent Client Protocol |

## 122. Laguna M.1 — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Sortie | **28 avril 2026** (avec XS.2) |
| Statut | disponible (tier gratuit OpenRouter fermé le 28/07/2026) ; entreprise |
| Architecture | **MoE : 225B totaux / 23B actifs par token** ; entraîné from scratch sur **30T tokens** avec **6 144 GPU NVIDIA H200** ; optimiseur Muon ; boucle RL asynchrone (transferts de poids <5 s via RDMA) |
| Contexte | 256K (selon table tiers, juil. 2026) |
| Features | agentic coding long-horizon ; agent `pool` ; sandbox `Shimmer` |
| Fine-tuning | **non vérifié au 27/09/2026** |
| Poids | **propriétaire** — poids partagés « sur demande » (pas de licence ouverte confirmée) |
| Benchmarks (vendor, avril 2026) | SWE-Bench Verified : 72,5 % (table intro) — table S 2.1 : 63,1 % multilingual ; SWE-Bench Pro : 46,9 % (intro) / 49,2 % (juil. 2026) ; Terminal-Bench 2.1 : 44,9 % |
| Prix API | $0/token pendant la période gratuite limitée |
| Déploiement | API Poolside, OpenRouter (free tier jusqu'au 28/07/2026) |

## 123. Laguna XS 2.1 — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Sortie | **2 juillet 2026** |
| Statut | disponible (supersedé comme « most capable » par S 2.1 le 21/07/2026) |
| Architecture | **MoE : 33B totaux / 3B actifs par token** ; mêmes données de pré-entraînement que XS 2.1/S 2.1 |
| Contexte | 256K |
| Features | agentic coding ; DFlash speculators (speculative decoding, 2× vitesse locale) |
| Fine-tuning | **non vérifié au 27/09/2026** |
| Poids | **ouverts — OpenMDW-1.1** (licence Linux Foundation) |
| Benchmarks (vendor) | SWE-Bench Multilingual : 63,1 % (+5,4 pp vs XS.2) ; Terminal-Bench 2.0 : 37,5 % |
| Prix API | **non vérifié au 27/09/2026** |
| Déploiement | Hugging Face, API Poolside, OpenRouter |

## 124. Laguna S 2.1 — fiche ✅ (modèle Poolside le plus récent au 27/09/2026)

