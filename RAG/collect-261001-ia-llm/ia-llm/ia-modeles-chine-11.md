---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-11
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "Hugging Face", "LongCat", "Meituan", "MiniMax", "OpenAI", "OpenRouter", "SGLang"]
dates: ["2025-08-30", "2025-09-22", "2026-04-02", "2026-06-30", "2026-09-27"]
keywords: ["agent", "agentic", "asic", "attention", "benchmarks", "embedding", "fine-tuning", "gemini", "inference", "moe", "multimodal", "open weights"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [995, 1079]
sha256: e7f3ae18434e39148d97edb4f25838d14c8a77cc11afc9d67dd06e41edc772ca
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

- **Nom exact** : LongCat-Flash-Chat (HF `meituan-longcat/LongCat-Flash-Chat`). **Sortie : 30 août 2025**
  (annonce ; blog SGLang 1er sept.). Statut : disponible, open weights. Antérieur à la période, mais c'est
  la fondation.
- **Architecture** : MoE, **560B total / 18,6–31,3B actifs (~27B en moyenne)**. **512 experts FFN + 256
  experts « zero-computation »**. **ScMoE** (Shortcut-Connected MoE : overlap calcul/communication),
  **MLA**, biais d'expert contrôlé par PID, hyperparameter transfer.
- **Contexte** : **128K tokens** (131 072).
- **Fonctionnalités** : non-thinking, agent/tool-use (focus), >100 tokens/s en inférence.
- **Poids** : **ouverts, MIT**.
- **Benchmarks** : MMLU 89,7 %, MATH-500 96,4 %, GPQA-Diamond 73,2 %, τ²-Bench Telecom 73,7 %.
- **Prix API** : ~$0,69 / 1M tokens (Caixin, 2025 — probablement obsolète au 27/09/2026).
- **Déploiement** : SGLang (blog officiel lm-sys), OpenRouter.

## 79. LongCat-Flash-Thinking (septembre 2025)

- **Nom exact** : LongCat-Flash-Thinking. **Sortie : 22 septembre 2025**. Statut : disponible, open weights
  (Hugging Face + GitHub, démo en ligne).
- **Architecture** : MoE reasoning, **560B total** (même base que Flash). Cold-start CoT long + RL large
  échelle ; schéma d'entraînement **« domain-parallel »** (fusion de modèles experts par domaine : STEM,
  Code, Agentic) ; framework RL **DORA** (Dynamic ORchestration for Asynchronous rollout, >3× speedup).
- **Fonctionnalités** : reasoning avec efficacité revendiquée : -64,5 % de tokens moyens sur AIME-25
  (19 653 → 6 965) sans perte de précision.
- **Poids** : ouverts (licence exacte non vérifiée — probablement MIT comme Flash, non confirmé).
- **Contexte / prix / déploiement** : non vérifiés au 27/09/2026.

## 80. LongCat-Next (avril 2026, le multimodal natif)

- **Nom exact** : LongCat-Next. **Sortie : 2 avril 2026** (article ; papier arXiv 2603.27538, mars 2026).
  Statut : disponible, open weights.
- **Architecture** : modèle **nativement multimodal** (texte, vision, audio) sous paradigme **DiNA**
  (Discrete Native Autoregression) — toutes les modalités en tokens discrets dans un espace partagé.
  Tokenizer visuel **dNaViT** (RVQ 8 couches). Backbone **LongCat-Flash-Lite MoE : ~68,5–75B total / ~3B
  actifs** (sources : 74B / 75B / 68,5B — incohérence mineure non vérifiée), attention MLA, N-gram
  over-embedding, tokenizers RVQ natifs vision+audio.
- **Fonctionnalités** : compréhension + **génération d'images**, compréhension audio, synthèse vocale /
  clonage de voix (24 kHz, streaming), tool calling. Pas de génération vidéo (compréhension seule).
- **Poids** : **ouverts, MIT** (poids, tokenizers, code d'inférence).
- **Contexte / prix / benchmarks** : non vérifiés au 27/09/2026.
- **Déploiement** : SGLang (adaptation officielle `LongCat-Next-inference`) ; build communautaire
  w8a8_int8 ~90 Go sur DGX Spark.

## 81. LongCat-2.0 (juin 2026, le 1,6T domestique)

- **Nom exact** : LongCat-2.0. **Sortie : 30 juin 2026**. Statut : annoncé + open-sourcé (MIT) — une source
  (5 juillet 2026) disait « weights coming soon » ; au 27/09/2026 le repo/blog officiel existe
  (longcat.ai/blog/longcat-2.0/), mais la disponibilité effective des poids reste **non vérifiée**.
- **Architecture** : MoE, **1,6T total / ~48B actifs (33–56B)**. Contexte **1M natif**, **LongCat Sparse
  Attention**, sortie max 128K. Entraîné et servi sur **supercalculateurs d'ASIC chinois domestiques**
  (cluster ~50 000 cartes — claim Meituan). « Super kernels » + L2-cache weight prefetching.
- **Fonctionnalités** : agentic coding.
- **Poids** : **MIT** (selon table marktechpost).
- **Benchmarks** (Meituan, non indépendants) : SWE-bench Pro 59,5 (vs GPT-5.5 58,6), Terminal-Bench 2.1
  70,8, SWE-bench Multilingual 77,3 ; « comparable à Gemini 3.1 Pro ».
- **Prix API / déploiement / fine-tuning** : non vérifiés au 27/09/2026.

## 82. LongCat : ScMoE et l'entraînement « domain-parallel »

Deux idées originales chez Meituan. **ScMoE** (Shortcut-Connected MoE) : des connexions « shortcut » entre
experts permettent de recouvrir (overlap) le calcul et la communication inter-experts — le goulot
classique du MoE à grande échelle. Les **experts « zero-computation »** (256 dans Flash-Chat) sont des
experts qui ne calculent rien mais routent : ils servent de « mémoire de routage » à coût nul. Côté
entraînement, le **domain-parallel** fusionne des modèles experts par domaine (STEM, Code, Agentic) au
lieu d'un seul run monolithique, et le framework **DORA** accélère le RL asynchrone (>3× speedup). C'est
une école d'ingénierie pragmatique : moins de théorie de l'attention, plus d'optimisation système.

## 83. LongCat : tableau comparatif

| Modèle | Date | Archi | Params (total/actifs) | Contexte | Prix API ($/1M in-out) | Licence |
|---|---|---|---|---|---|---|
| LongCat-Flash-Chat | 30/08/2025 | MoE, ScMoE, MLA | 560B / ~27B | 128K | ~0,69 (2025, obsolète ?) | MIT |
| LongCat-Flash-Thinking | 22/09/2025 | MoE reasoning, DORA | 560B / non vérifié | non vérifié | non vérifié | Ouverts (prob. MIT) |
| LongCat-Next | 02/04/2026 | MoE multimodal DiNA | ~68,5–75B / ~3B | non vérifié | non vérifié | MIT |
| LongCat-2.0 | 30/06/2026 | MoE, LongCat Sparse Attn | 1,6T / ~48B | 1M natif | non vérifié | MIT (poids à confirmer) |

## 84. MiniMax : portrait de l'entreprise

MiniMax (Shanghai) est le laboratoire chinois le plus « produit » du volume : modèles de langue (série M),
génération vidéo/audio (série H), API internationale. En 2026, la série **M2** (M2 → M2.7) a conquis le
titre de « roi de l'open-source » sur l'agentique, **M3** (1er juin 2026) est passé au 1M de contexte
multimodal, et **H3** (juillet 2026) a ouvert la génération vidéo+audio — avec une **licence territoriale
restrictive** inédite qui mérite sa section. Dérivé notable : **humain-m3**, modèle arabe co-développé
avec HUMAIN (Arabie saoudite).

## 85. MiniMax-M2 (série : M2, M2.1, M2.5, M2.7)

