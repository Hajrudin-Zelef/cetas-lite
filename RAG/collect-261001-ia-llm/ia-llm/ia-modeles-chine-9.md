---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-9
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Ant", "Anthropic", "DeepSeek", "Hugging Face", "Xiaomi", "vLLM"]
dates: ["2025-04-30", "2025-09-15", "2025-12-17", "2026-03-18", "2026-04-22", "2026-09-21", "2026-09-27"]
keywords: ["agent", "attention", "benchmarks", "claude", "deepseek", "distillation", "gqa", "licenses", "mai", "moe", "omni", "open weights"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [797, 890]
sha256: 0d595036d2dbfe12233783d56fb305fda09305c0ee053d2db4b6d33624852c5a
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

Anthropic accuse Xiaomi d'avoir utilisé **Claude à grande échelle pour la collecte de données**
(distillation non autorisée) en septembre 2026 — allégation **non tranchée** au 27/09/2026. C'est la
deuxième affaire « distillation sauvage » de l'année côté chinois après les soupçons récurrents visant
d'autres labos. Enjeu : les conditions d'utilisation des API frontier interdisent l'entraînement de
modèles concurrents sur leurs sorties, mais la preuve est techniquement difficile (détection par
« fingerprints » stylistiques, canaris). À suivre : si l'accusation se confirme, elle fragilisera la
légitimité open-weight de toute la lignée MiMo.

## 62. MiMo : tableau comparatif

| Modèle | Date | Archi | Params (total/actifs) | Contexte | Prix API ($/1M in-out) | Licence |
|---|---|---|---|---|---|---|
| MiMo-7B | 30/04/2025 | Dense, MTP, reasoning RL | 7B | non vérifié | — | MIT |
| MiMo-V2-Flash | 17/12/2025 | MoE, hybride 5:1 | 309B / 15B | non vérifié | non vérifié | MIT |
| MiMo-V2-Pro | 18/03/2026 | MoE | >1T / 42B | non vérifié | non vérifié | Fermé |
| MiMo-V2.5 | 22/04/2026 | MoE | 310B / actifs non vérifiés | non vérifié | non vérifié | MIT |
| MiMo-V2.5-Pro | 22/04/2026 | MoE | 1,02T / actifs non vérifiés | non vérifié | non vérifié | MIT |
| MiMo-V2.6-Pro-RL | 21/09/2026 | MoE, SW + globale, omnimodal | 1,02T / 42B | 1M | 0,435 / 0,87 (tiers) | MIT |
| MiMo-V2.6-Flash-RL | 21/09/2026 | MoE, omnimodal | 309B / 15B | 1M | 0,14 / 0,28 (tiers) | MIT |
| MiMo-V2.6-Distill-Qwen-9B | 21/09/2026 | Dense distillé | 9B | non vérifié | — | MIT |

## 63. MiMo : ce que la lignée raconte

Xiaomi a fait en 18 mois le trajet inverse d'Alibaba : parti de l'open-weight pur (7B), passé par une
parenthèse propriétaire (V2-Pro, mars 2026), revenu à l'ouvert avec une radicalité nouvelle (RL public,
7 000 environnements publiés). La constante technique : le MoE sparse à ratio ~20–24× (309B/15B, 1,02T/42B),
l'omnimodalité native, et une équipe issue de DeepSeek qui applique les recettes DeepSeek avec les moyens
de Xiaomi. Le Distill-Qwen-9B est le geste le plus intéressant : distiller un trillion-class vers 9B
utilisables en local, c'est admettre que la valeur n'est plus dans la taille mais dans la recette.

## 64. Ant Group / inclusionAI : portrait

« Ling » est la famille de LLM open-source d'**Ant Group** (le lab s'appelle **inclusionAI**), aussi
appelée **BaiLing / Bailing**. Site : ant-ling.com. Org Hugging Face : `inclusionAI`. GitHub :
`inclusionAI/ling-cookbook`. Ant est l'acteur « efficacité » du paysage chinois : là où les autres
empilent les billions de paramètres, Ling optimise les **paramètres actifs par token** (5,1B actifs pour
le 3.0-flash à 124B totaux) et l'attention linéaire. C'est aussi le seul à décliner nativement une
version **finance** (Fin) co-développée avec des institutions financières.

## 65. Les trois séries Ling / Ring / Ming

La nomenclature Ant tient en trois séries. **Ling** : les modèles « non-thinking » (réponse directe,
efficacité). **Ring** : les modèles « thinking » (raisonnement explicite, CoT long). **Ming** : la série
multimodale omni (texte, vision, audio, musique). Un même socle technique (familles `BailingMoeV2`,
`BailingMoeV3`) décliné selon le besoin : Ling-2.5-1T (flagship non-thinking) et Ring-2.5-1T (premier
thinking à attention linéaire hybride) partagent la génération 2.5 ; Ming-Flash-Omni-2.0 porte la
multimodalité. Retenir le triplet, c'est ne plus se perdre dans les noms.

## 66. Ling 2.0 (octobre 2025, la fondation)

- **Noms exacts** : Ling-mini-2.0, Ling-flash-2.0, Ling-1T (+ variantes base). **Sortie : octobre 2025**.
  Statut : disponible, open weights.
- **Architecture** : `BailingMoeV2ForCausalLM` (HF `model_type: bailing_moe`), MoE. Attention **GQA**
  (16 têtes / 4 KV, head_dim 128) + QK-RMSNorm + demi-RoPE. Routing style DeepSeek-V3 : sigmoid, biais par
  expert, MoE groupé (n_group=8, topk_group=4, top-8 experts), 1 expert partagé, `first_k_dense_replace`
  (premières couches denses).
  - mini-2.0 : **16B total / ~1,4B actifs** (20 couches, 256 experts).
  - flash-2.0 : **100B total / ~6,1B actifs** (32 couches, 256 experts).
  - Ling-1T : **1T total / ~50B actifs** (80 couches).
- **Poids** : ouverts (licence exacte non vérifiée pour chaque variante ; flash-2.0 cité MIT par une source
  tierce). **Contexte / prix / benchmarks** : non vérifiés au 27/09/2026.
- **Déploiement** : vLLM (support `bailing_moe.py`, PR #24627 mergée le 2025-09-15).

## 67. Ling-2.5-1T / Ring-2.5-1T (février 2026)

- **Noms exacts** : Ling-2.5-1T, Ring-2.5-1T. **Sortie : 16 février 2026** (annonce Business Wire). Statut :
  disponible, open licenses (Hugging Face + ModelScope).
- **Architecture** : **1T paramètres** ; Ling-2.5-1T : flagship non-thinking, efficacité de raisonnement ;
  Ring-2.5-1T : **premier modèle thinking à architecture linéaire hybride** (hybrid linear attention).
- **Contexte** : **1M tokens** (Ling-2.5-1T).
- **Fonctionnalités** : agent natif, alignement de préférences fin, token usage réduit (AIME 2026 : ~5 890
  tokens vs 15–23K pour les frontier thinking models) — l'argument « raisonner moins pour raisonner mieux ».
- **Benchmarks** (Ant) : Ring-2.5-1T — IMO 2025 : 35/42 (niveau médaille d'or), CMO 2025 : 105/126
  (> cutoff équipe nationale chinoise).
- **Poids** : ouverts (licence exacte non vérifiée).

## 68. Ming-Flash-Omni-2.0 (février 2026)

- **Nom exact** : Ming-Flash-Omni-2.0. **Sortie : 11 février 2026**. Série multimodale Ming.
- Présenté comme le **premier modèle unifiant parole, audio et musique dans une seule architecture**.
- Specs détaillées : non vérifiées au 27/09/2026. C'est la fiche la plus mince du volume côté Ant :
  l'annonce existe (même Business Wire que Ling-2.5), le détail technique n'a pas filtré.

## 69. Ling-2.6-flash (avril 2026)

- **Nom exact** : Ling-2.6-flash. **Sortie : avril 2026**. Statut : disponible (API ; poids annoncés mais
  **non encore publiés en mai 2026** selon lilting.ch — statut des poids au 27/09/2026 : non vérifié).
- **Architecture** : MoE, **104B total / 7,4B actifs**.
- **Fonctionnalités** : agent/tool-use, production à haut volume — le « cheval de trait » de la gamme.
- C'est le chaînon entre la génération 2.x (GQA classique) et la 3.0 (KDA + MLA) : mêmes ordres de
  grandeur, architecture en transition.

## 70. Ling-3.0-flash (juillet 2026, le flagship efficacité)

