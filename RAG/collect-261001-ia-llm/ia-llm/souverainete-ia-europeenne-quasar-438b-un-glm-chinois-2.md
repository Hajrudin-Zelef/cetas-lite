---
id: collect-261001-ia-llm/ia-llm/souverainete-ia-europeenne-quasar-438b-un-glm-chinois-2
title: "souverainete-ia-europeenne-quasar-438b-un-glm-chinois"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "EU", "Mistral", "Nvidia", "Z.ai"]
dates: ["2026-09-10"]
keywords: ["glm", "apache", "benchmark", "benchmarks", "claude", "cost", "datacenter", "decode", "deepseek", "fable 5", "mistral", "moe"]
source: docs/RAG/collect-261001-ia-llm/souverainete-ia-europeenne-quasar-438b-un-glm-chinois.md
source_anchor: ""
source_lines: [39, 84]
sha256: 9f3407165f14987d8ee34fc8f79b9266c48e1fad7c3fd55b23ac04c7822e3233
---

# souverainete-ia-europeenne-quasar-438b-un-glm-chinois

L’affaire Quasar 438B et la décision de Mistral AI mettent en lumière deux définitions concurrentes de la souveraineté IA européenne. La première, défendue implicitement par Multiverse Computing lors du lancement de Quasar, associe souveraineté à performance affichée sous une bannière européenne, quelle que soit l’origine réelle de la technologie sous-jacente. La seconde, plus proche de la position de Mistral AI, distingue la souveraineté d’infrastructure, c’est-à-dire l’hébergement et le traitement des données sur le sol européen sous droit européen, de la souveraineté d’architecture, qui exigerait que le modèle lui-même soit conçu et entraîné en Europe.

Cette ambiguïté n’est pas nouvelle. Elle traverse déjà le débat sur le Cloud Act américain et les exigences de résidence des données imposées par le RGPD. Mais l’affaire Quasar 438B l’étend pour la première fois de manière aussi visible au niveau du modèle d’IA lui-même, et non plus seulement à l’infrastructure qui l’héberge. Un modèle peut être hébergé en Europe, sur des serveurs européens, tout en étant construit à partir d’une architecture et de données entièrement extérieures au continent. La question posée aux régulateurs devient alors : la souveraineté se mesure-t-elle à l’adresse du datacenter ou à l’origine du code et des données d’entraînement ?

## Ce que disent les benchmarks : Quasar, GLM-5.2 et Mistral face à face

Au-delà de la polémique sur l’origine des modèles, les chiffres de performance publiés en septembre 2026 dessinent un paysage compétitif resserré. Le tableau suivant compare les principaux modèles cités dans le débat sur la souveraineté IA européenne, à partir des données publiées par BenchLM, Startup Fortune et ThursdAI.

| Modèle | Origine | Éditeur | Score / indice de référence | Paramètres | 
|---|---|---|---|---|
| Quasar 438B | Espagne (revendiqué) | Multiverse Computing | 43 (Artificial Analysis Intelligence Index v4.1.1) | 438 Md (modèle compressé) | 
| Mistral Medium 3.5 | France | Mistral AI | 30 sur le même indice (-13 pts vs Quasar) | Non communiqué | 
| Ministral 3 14B (Reasoning) | France | Mistral AI | 47,2 (classement européen BenchLM, 10/09/2026) | 14 Md | 
| GLM-5.2 | Chine | Z.ai | Dépasse Claude Opus 4.7 sur plusieurs tests | Non communiqué | 
| DeepSeek V4.1 Flash | Chine | DeepSeek | 90,6 (Terminal-Bench 2.1) / 74,2 (DeepSWE 1.1) | 552 Md (8 Md actifs prefill / 16 Md decode) | 
| SWE-2 | États-Unis | Cognition | 50 (Frontier Code 1.1) / 92,8 (Terminal-Bench 2.1) | Non communiqué | 

Ce tableau illustre un point souvent négligé dans la couverture médiatique de ces classements : les indices utilisés (Artificial Analysis Intelligence Index, BenchLM, Terminal-Bench) ne mesurent pas les mêmes capacités et ne sont pas directement comparables entre eux. BenchLM revendique le suivi de plus de 232 modèles et la comparaison de plus de 489 grands modèles de langage à travers 435 bancs d’essai différents, avec 19 nouvelles publications de modèles recensées sur le seul mois de septembre 2026. Sur son classement général BenchAlign, c’est Claude Fable 5.1, d’Anthropic, qui occupe la première place avec un score de 84,61, loin devant l’ensemble des modèles européens cités dans cet article.

## Panorama des modèles IA européens open source en 2026

La controverse Quasar 438B intervient alors que l’Europe dispose déjà d’un tissu de modèles ouverts authentiquement développés sur le continent, recensés notamment par l’European Open Source AI Index. Le tableau ci-dessous présente les principaux projets nationaux et européens actifs en septembre 2026.

| Modèle | Pays / Zone | Organisation | Licence | Dernière version connue | 
|---|---|---|---|---|
| Mistral Large 3 | France | Mistral AI | Apache 2.0 | Décembre 2025 (675 Md MoE, 41 Md actifs) | 
| Apertus | Suisse | ETH Zurich / EPFL / CSCS | Ouverte (OSI) | 2026 | 
| ALIA | Espagne | Gouvernement espagnol / BSC | Ouverte (OSI) | 2026 | 
| Teuken-7B | Allemagne | OpenGPT-X | Ouverte (OSI) | 2026 | 
| Bielik | Pologne | Consortium SpeakLeash | Ouverte (OSI) | 2026 | 
| Velvet | Italie | iGenius | Ouverte (OSI) | 2026 | 
| EuroLLM-22B | UE (consortium) | OpenEuroLLM | Ouverte (OSI) | 2026 | 
| Amália | Portugal | Gouvernement portugais | Ouverte | Juillet 2026 (coût : 7 M€) | 

Ce panorama, mis à jour le 25 août 2026 par mrkt30.com, montre que Mistral Large 3 reste, selon la cartographie disponible, le seul modèle européen à l’échelle frontière avec des poids ouverts. Tous les autres projets listés opèrent à une échelle nettement plus modeste, souvent centrée sur des besoins linguistiques ou administratifs nationaux plutôt que sur la compétition frontale avec les meilleurs modèles généralistes mondiaux. C’est précisément cet écart d’échelle qui explique pourquoi un modèle comme Quasar 438B, prétendant combler ce vide avec 438 milliards de paramètres, a suscité un tel enthousiasme initial avant que la controverse n’éclate.

## Contexte historique : dix ans de dépendance technologique

Le débat sur la souveraineté IA européenne ne date pas de 2026. Il prolonge une inquiétude plus ancienne sur la dépendance du continent envers les infrastructures cloud américaines, puis envers les puces d’entraînement conçues par Nvidia. L’arrivée de Mistral AI en 2023 avait été présentée comme la première réponse crédible à cette dépendance sur le terrain des modèles de langage. Trois ans plus tard, l’entreprise reste effectivement la référence européenne, mais l’épisode Quasar 438B et l’ouverture de sa plateforme à GLM-5.2 montrent que la course ne se joue plus seulement entre l’Europe et les États-Unis : la Chine s’est imposée comme un troisième pôle incontournable, avec DeepSeek et désormais Z.ai en position de force sur les modèles ouverts.

Face à cette double pression américaine et chinoise, plusieurs États européens ont lancé leurs propres initiatives nationales plutôt que d’attendre une réponse continentale unifiée. Le Portugal a dévoilé Amália le 1er juillet 2026 pour un coût de 7 millions d’euros, selon Actu IA, un montant dérisoire comparé aux budgets d’entraînement des modèles frontière américains ou chinois, mais suffisant pour produire le premier grand modèle de langage ouvert en portugais européen. Cette approche low-cost et ciblée contraste avec la stratégie de Multiverse Computing, qui a préféré revendiquer une performance de pointe généraliste plutôt qu’une spécialisation nationale ou linguistique.

## Le cadre réglementaire : AI Act et benchmark EU MMLU

L’affaire Quasar 438B intervient alors que la Commission européenne renforce ses propres outils d’évaluation des modèles d’IA. Le benchmark EU MMLU, mis à jour le 22 juillet 2026, est désormais disponible en 16 langues officielles de l’Union, dont le français, l’allemand, le néerlandais et le portugais, selon la direction générale de la traduction de la Commission européenne. L’objectif affiché est de proposer une grille d’évaluation multilingue indépendante des laboratoires eux-mêmes, précisément pour éviter les situations où un fournisseur communique son propre score sans possibilité de vérification externe standardisée.

