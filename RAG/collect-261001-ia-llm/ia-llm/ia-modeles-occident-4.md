---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-4
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Apple", "Meta", "Microsoft", "OpenAI", "Stripe", "United States"]
dates: ["2026-02-05", "2026-02-12", "2026-03-03", "2026-03-05", "2026-03-17", "2026-04-14", "2026-04-23", "2026-06-22", "2026-07-09", "2026-08-10", "2026-09-03", "2026-09-22", "2026-09-24", "2026-09-27"]
keywords: ["agent", "agentic", "agents", "apache", "astra", "benchmarks", "claude", "cyber", "exploit", "fable 5", "gpt-5.6", "gpt-6"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [257, 366]
sha256: f2b184d3558fd24424714204125020dd0bfdfd7cb510e984cef4b7462527ab25
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Modèle | Sortie | Statut | Contexte | Prix in/out par MTok | Notes |
|---|---|---|---|---|---|
| GPT-5.3-Codex | 05/02/2026 | dispo (LTS) | non vérifié | $1,75/$14 (source unique) | agentic coding |
| GPT-5.3-Codex-Spark | 12/02/2026 | preview | 128K | — | pair programmer temps réel |
| GPT-5.3 Instant | 03/03/2026 | dispo | — | ~$0,30/$1,20 | tier pas cher |
| GPT-5.4 | 05/03/2026 | dispo (retiré Codex login 31/08) | 1,05M | $2,50/$15 | computer use natif |
| GPT-5.4 Pro | 05/03/2026 | dispo | 1,05M | $30/$180 | hardest reasoning |
| GPT-5.4 mini / nano | 17/03/2026 | dispo | — | $0,75/$4,50 ; $0,20/$1,25 | — |
| GPT-5.4-Cyber | 14/04/2026 | trusted access | — | — | Daybreak |
| GPT-5.5 | 23/04/2026 | dispo | 1M (400K Codex) | $5/$30 | réentraîné de zéro |
| GPT-5.5 Pro | 23/04/2026 | dispo | 1M | $30/$180 | — |
| GPT-5.5-Cyber | 22/06/2026 | Daybreak | — | — | — |
| GPT-5.6 Sol | 09/07/2026 (GA) | dispo | 1,05M | $4/$20 (promo 21/08) | flagship juil–sept |
| GPT-5.6 Terra | 09/07/2026 (GA) | dispo | 1,05M | $2/$12 | balanced |
| GPT-5.6 Luna | 09/07/2026 (GA) | dispo | ~1M | $0,20/$1,20 | cheap |
| GPT-5.6-Cyber | 10/08/2026 | Daybreak Red | — | — | — |
| **GPT-6 Astra** | **03/09/2026** | **dispo — flagship** | **1,05M** | **$10/$50** | Critical cyber ; computer use |
| **GPT-6 Sol** | **22/09/2026** | dispo | 1,05M | **$2/$10** | coding/agents |
| **GPT-6 Luna** | **22/09/2026** | dispo | 1,05M | **$0,10/$0,50** | volume |
| « GPT Pro » standalone | — | **N'EXISTE PAS** | — | — | abonnement + variantes -Pro |
| GPT-6 Terra | — | **NON SORTI** | — | — | existence future non vérifiée |

---

# PARTIE II — Meta : Muse (agent) et Muse Spark (modèles)

## 25. Muse : le produit agent personnel de Meta ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Muse — « personal AI agent » |
| Sortie | **8–9 septembre 2026** (annonce newsroom Meta) |
| Statut | disponible — **US uniquement au lancement** ; iOS, Android, web (muse.ai), WhatsApp ; 18+ ; intégration lunettes IA Meta prévue |
| Modèle sous-jacent | **Muse Spark 1.3** |
| Architecture produit | tourne dans une **Muse Secure VM** (cloud computer dédié avec son propre navigateur) ; **Sentinel** = processus séparé, seule autorité pour les actions connecteurs et l'egress réseau (kernel-level) ; 5 couches anti-prompt-injection revendiquées (training, harness labeling, classifieurs parallèles, approbation humaine pour data sortante, bornes déterministes) ; credentials dans `hatch-authd` (jamais lisible par l'agent) ; bug bounty public jusqu'à $300K (dont $130K pour exploit prompt-injection) ; achats via Stripe Link (cartes virtuelles single-use) |
| Tarifs | gratuit jusqu'à **100M tokens/semaine** ; plans payants **$20 et $100/mois** |
| Connect 24/09/2026 | Muse Realtime Avatar (vidéo live), intégration smart glasses (wake word), **Mac desktop control** (computer use comparable OpenAI/Anthropic), shopping partners |

## 26. Sécurité Muse : Sentinel, Secure VM et le pari anti-injection

Le design sécurité de Muse est le plus documenté du marché en 2026 : l'agent ne touche jamais directement au réseau — **Sentinel**, processus séparé au niveau kernel, est la seule autorité pour les actions connecteurs et l'egress. Cinq couches anti-prompt-injection revendiquées, bug bounty jusqu'à $300K ($130K pour un exploit d'injection). Les credentials vivent dans `hatch-authd`, illisibles par l'agent. C'est une réponse directe au problème n°1 des agents 2026 : l'injection de prompt via le contenu web.

## 27. Muse Spark 1.0 ✅

| Champ | Valeur vérifiée |
|---|---|
| Sortie | **8–9 avril 2026** |
| Statut | disponible — premier modèle du stack reconstruit |
| Architecture | **non vérifié au 27/09/2026** (dense/MoE, paramètres non publiés par Meta) |
| Features | multimodal natif (vision chain-of-thought, tool use, orchestration multi-agents) ; mode « Contemplating » : **58 % Humanity's Last Exam, 38 % FrontierScience Research** ; « >10× plus efficient que Llama 4 Maverick » (revendiqué) |
| Poids | **fermés** |
| Déploiement | propulse Meta AI (Facebook, Instagram, WhatsApp) |

## 28. Muse Spark 1.1 ✅

| Champ | Valeur vérifiée |
|---|---|
| Sortie | **9 juillet 2026** |
| Statut | disponible |
| Contexte | **1M tokens** |
| Features | Meta Model API (public preview, OpenAI-compatible) ; inputs text+image+video+document, output text ; planning mode, goal conditioning, subagent delegation, context compaction |
| Poids | fermés |
| Prix API | **$1,25 in / $4,25 out** par MTok (+$20 crédits gratuits) ; tier **« Contributor » $0,10 / $0,20** (opt-in : usage des requêtes pour l'entraînement) |

## 29. Muse Spark 1.2 ✅

| Champ | Valeur vérifiée |
|---|---|
| Sortie | **5–6 août 2026** |
| Statut | disponible |
| Contexte | 1M tokens |
| Features | + **Muse Code** (bêta, terminal agent macOS/Linux, co-entraîné avec le harness) ; weights promis mais **non publiés au 27/09/2026** |
| Poids | fermés |
| Prix API | $1,25 / $4,25 |
| Benchmarks clés (Meta-déclarés) | AA Intelligence Index : 57 (ex-aequo avec GPT-5.6 Terra) |

## 30. Muse Spark 1.3 ✅ (modèle actuel de Muse)

| Champ | Valeur vérifiée |
|---|---|
| Sortie | **2 septembre 2026** |
| Statut | disponible — **modèle actuel de l'agent Muse** |
| Contexte | 1M tokens |
| Features | ~25 % de tokens en moins, ~20 % d'appels outils en moins vs 1.2 ; retrieval long-contexte 98,5 % (256–512K) / 98,1 % (512K–1M) |
| Poids | fermés |
| Prix API | inchangés : $1,25 / $4,25 |
| Benchmarks clés (Meta-déclarés uniquement dans les sources consultées) | **DeepSWE 1.1 : 75,4 %**, **Terminal-Bench 2.1 : 88,8 %** ; AA Intelligence Index : **62** (derrière Claude Fable 5.1 et Opus 5) ; LiveBench 5e à $0,219/tâche |

## 31. Muse Glimmer 30B : open-weights, source unique ⚠️

**Muse Glimmer** : modèle agentic 30B open-weights, Apache 2.0, daté du 10 août 2026 — source unique (dim-s/prompt-atlas), **non vérifié au 27/09/2026** au-delà de cette source. Si confirmé, ce serait la seule ouverture de poids de la lignée Muse Spark.

## 32. « Watermelon » : nom de code non vérifié ❓

Nom de code interne du prochain modèle Meta (architecture plus large), cité par Alexandr Wang — **non vérifié au 27/09/2026**, pas de date, pas d'annonce.

## 33. Tarifs Muse Spark et le tier « Contributor »

Le pricing $1,25/$4,25 par MTok place Muse Spark 1.x en frontal avec GPT-6 Sol ($2/$10) et sous Claude Opus 5.5 ($4/$20). L'originalité : le tier **Contributor à $0,10/$0,20** — l'utilisateur accepte que ses requêtes servent à l'entraînement contre un prix cassé. C'est la première offre grand public qui monétise explicitement le consentement data.

## 34. Stratégie Meta : l'abandon du Llama open-weight

Depuis avril 2026, Meta a pivoté : **abandon du Llama open-weights** au profit de modèles propriétaires sous la marque **Muse** (Meta Superintelligence Labs). Llama 4 Scout/Maverick (avril 2025) restent les derniers poids ouverts Meta. Une source secondaire (codersera, ferme de contenu) affirme que Behemoth aurait été « shelved » — **non vérifié au 27/09/2026**, à traiter avec scepticisme. Le signal stratégique reste : Meta ne joue plus l'open-weight.

## 35. Distinguer « Muse » des autres produits nommés Muse ⚠️

**« Muse » = l'agent personnel Meta** (produit), **« Muse Spark » = la famille de modèles** qui le propulse. Ne pas confondre avec d'autres produits nommés Muse (éditeurs, jeux) — aucun autre « Muse » vérifié dans cette recherche. En septembre 2026, « Muse » sans précision désigne l'agent Meta dans la presse tech.

## 36. Tableau récapitulatif Meta Muse (sept. 2026)

