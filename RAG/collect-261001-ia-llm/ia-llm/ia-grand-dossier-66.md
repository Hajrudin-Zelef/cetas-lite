---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-66
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Cerebras", "DeepSeek", "OpenAI", "OpenRouter", "xAI"]
dates: []
keywords: ["arr", "bedrock", "benchmarks", "compute", "cost", "deepseek", "distillation", "gpu", "hyperscaler", "kv cache", "llama", "llama.cpp"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [5498, 5540]
sha256: d5e075ae7b9cd7c59e259e5a91bb307c07f1a14dc470c61915d4550c687456d3
---

# IA — Le grand dossier

> **Corrigé.** Camp « ralentissement » : (a) **rendements décroissants mesurés** — GPT-4.5 « Orion » (2025) puis GPT-5 (2025) n'apportent plus de rupture pour des coûts explosifs (presse spécialisée) ; (b) **mur des données** — Epoch AI (Villalobos et al.) : stock de texte de qualité épuisé entre 2026 et 2032 ; Musk (2025) : « somme cumulée des connaissances humaines épuisée » ; (c) **coûts** — GPT-4 > 100 M$ (Altman), coût total de développement 2-3× le run final (Epoch AI). Camp « continuation » : (a) **Gwern** (« The Scaling Hypothesis ») : les courbes ne fléchissent pas quand on les trace bien, les Cassandre ont toujours eu tort ; (b) **le scaling a changé d'axe** — test-time compute (Snell et al., 2024), **Bitter Lesson** de Sutton (2019) : apprentissage + search scalent indéfiniment ; les sauts o1/R1 (AIME 13 % → 83 % → 96,7 %) sans rupture de training le prouvent. (Sections 4.2-4.3.)

**Q7.** Expliquez le principe du Mixture of Experts (MoE) avec l'exemple chiffré de DeepSeek-V3, et dites pourquoi c'est une réponse au coût de l'inférence.

> **Corrigé.** En MoE, chaque token n'est traité que par **quelques experts** (sous-réseaux) parmi des dizaines/centaines, sélectionnés par un routeur. DeepSeek-V3 : **671B paramètres totaux mais seulement 37B actifs par token** (~5,5 %). On obtient la **capacité** d'un modèle géant pour le **coût d'inférence** d'un modèle moyen — d'où son rôle central dans la baisse des coûts et la réponse au mur du training : la performance ne vient plus de « tout activer », mais de « bien router ». (Sections 4.4, 3.2.)

**Q8.** Qu'est-ce que le « 80/20 reversal » décrit par les analystes en 2026, et quelle décision d'infrastructure en découle pour un sysadmin ?

> **Corrigé.** Le « 80/20 reversal » (Deloitte/Lenovo, analystes 2026) : ~80 % du compute allait au **training** en 2024, ~80 % irait à **l'inférence** fin 2026. Décision sysadmin : **ne plus dimensionner pour l'entraînement** (risque de clusters « stranded ») mais pour le **serving** — flotte d'inférence élastique, optimisation du coût par token (quantification, batching, routage petit/gros modèle), FinOps IA. (Sections 4.2, 6.1.)

**Q9.** Pourquoi le sur-entraînement volontaire (ex. : Llama 3 8B à ~1 875 tokens/paramètre) est-il économiquement rationnel alors qu'il contredit l'optimum Chinchilla ?

> **Corrigé.** Chinchilla optimise le **coût du training** (one-shot). Mais le training est un coût **unique** et l'inférence un coût **récurrent** sur toute la vie du modèle : un petit modèle sur-entraîné (Llama 3 8B : 15 000B tokens, ~1 875 tok/param, ~100× Chinchilla) est **moins cher à servir à chaque requête, pour toujours**. L'optimum économique total ≠ l'optimum training-compute. En 2025, la moyenne des modèles ouverts est à ~300 tok/param (~15× Chinchilla). (Section 4.4.)

**Q10.** Un collègue affirme : « Puisque le training ne scale plus, le progrès de l'IA est terminé. » En 5 lignes, réfutez en mobilisant au moins 3 déplacements du progrès décrits dans ce dossier.

> **Corrigé.** Exemple de réfutation : le progrès s'est **déplacé**, pas arrêté. (1) **Test-time compute** : o1/R1 ont fait des sauts discontinus (AIME 13 % → 96,7 %) sans scaling du training. (2) **Efficacité** : MoE, distillation (raisonnement o1-class dans un 7B local), quantification — plus d'intelligence par dollar et par watt. (3) **RL + données synthétiques/vérifiables** : R1-Zero apprend à raisonner par pur RL, sans données humaines — la voie AlphaGo Zero appliquée au langage. La fin du scaling **naïf du pré-entraînement** n'est pas la fin du progrès : c'est la fin du gaspillage. (Section 4.7.)

---

### Quiz — Partie 2 : Providers, passerelles, IA locale (10 questions)

**Q1.** Quelle est la différence entre un *lab* et un *provider*, et pourquoi est-ce important pour choisir une offre ?
> **Corrigé.** Le lab entraîne le modèle (OpenAI, DeepSeek…), le provider le sert via une API (Together, DeepInfra, Bedrock…). Un même modèle est servi par dix providers à dix prix : on ne choisit donc pas « un lab » mais un couple (modèle, hébergeur) optimisé en coût/vitesse/qualité — d'où les routeurs.

**Q2.** Citez trois leviers tarifaires qui divisent par deux (ou plus) une facture d'API, avec un exemple chiffré chacun.
> **Corrigé.** (1) **Batch API** : -50 % (OpenAI, Together, DeepInfra) — ex. indexation RAG nocturne. (2) **Prompt caching** : lecture à -90 % — ex. chez Anthropic un préfixe système réutilisé passe de 3,00 $ à 0,30 $/1M (Sonnet 5). (3) **Choisir l'hébergeur le moins cher du même modèle** : ex. GPT-OSS-120B de ~0,09 $ (DeepInfra) à 0,35 $/1M in (Cerebras) selon la vitesse voulue.

**Q3.** Pourquoi OpenRouter facture-t-il 5,5 % sur les achats de crédits alors qu'il « ne prend pas de marge sur le token » ?
> **Corrigé.** Son modèle économique n'est pas une marge par token (tu paies le prix public du provider sous-jacent) mais une commission sur le flux financier (achat de crédits, min 0,80 $) et +5 % en BYO-key. C'est le prix de l'agrégation : une clé, 400+ modèles, routage et fallbacks natifs.

**Q4.** Dans une config LiteLLM, à quoi servent respectivement `router_settings.routing_strategy`, la section `fallbacks`, et les clés virtuelles ?
> **Corrigé.** `routing_strategy` = comment choisir un deployment dans un groupe (ex. `usage-based-routing-v2` : latence+charge+erreurs ; `cost-based-routing` : le moins cher). `fallbacks` = où basculer quand un groupe échoue (ex. `triage → raisonnement`, avec variantes par type d'erreur : 429, contexte dépassé, refus de modération). Les clés virtuelles = contrôle d'accès et **budgets par app/équipe** (plafond $, modèles autorisés, RPM/TPM) + spend logs pour le pilotage.

**Q5.** Un collègue veut faire tourner un 70B Q4_K_M sur sa RTX 5090. Que lui répondez-vous, chiffres à l'appui ?
> **Corrigé.** Impossible en full-GPU : le fichier fait ~40–43 Go et il faut ~48 Go avec KV cache + overhead, pour 32 Go de VRAM sur la 5090. Options : 2×24 Go (offload), carte 48 Go (RTX 6000 Ada), Mac ≥64 Go unifiés, ou descendre à un 32B Q4 (~22–24 Go, plafond confortable de la 5090). En dernier recours : offload partiel CPU via `--n-gpu-layers` (llama.cpp), lent (quelques tok/s).

**Q6.** Q4_K_M vs Q8_0 vs FP16 : donnez l'ordre de grandeur de la mémoire pour un 8B et la perte de qualité associée.
> **Corrigé.** 8B : Q4_K_M ≈ 5 Go de fichier / ~6–8 Go en mémoire ; Q8_0 ≈ 8–9 Go / ~11 Go ; FP16 ≈ 16 Go / ~19 Go. Qualité : Q8 ≈ 100 % du FP16, Q4 ≈ 92–98 % (perplexité +0,05–0,2 pt, benchmarks -1–2 %). En prod : Q4 par défaut, Q8 si la VRAM le permet sur tâche exigeante.

**Q7.** Citez les quatre raisons de passer en IA locale et, pour chacune, un contre-exemple où le cloud reste meilleur.
> **Corrigé.** (1) **Coût** à fort volume — mais sous ~1M tokens/jour le cloud est moins cher (pas d'amortissement). (2) **Confidentialité** — mais avec DPA + zero-retention contractuelle le cloud convient aux données non sensibles. (3) **Offline/résilience** — mais la dispo d'un GPU perso < SLA 99,9 % d'un hyperscaler. (4) **Souveraineté/reproductibilité** — mais la qualité max reste aux frontier cloud (un 32B local ≈ 70–85 % d'un frontier).

