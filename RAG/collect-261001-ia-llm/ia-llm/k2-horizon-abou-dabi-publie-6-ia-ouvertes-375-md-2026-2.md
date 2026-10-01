---
id: collect-261001-ia-llm/ia-llm/k2-horizon-abou-dabi-publie-6-ia-ouvertes-375-md-2026-2
title: "k2-horizon-abou-dabi-publie-6-ia-ouvertes-375-md-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Z.ai", "vLLM", "xAI"]
dates: []
keywords: ["apache", "claude", "deepseek", "gemini", "glm", "gpt-5.6", "grok", "grok 4", "kimi", "mai", "mistral", "nvidia"]
source: docs/RAG/collect-261001-ia-llm/k2-horizon-abou-dabi-publie-6-ia-ouvertes-375-md-2026.md
source_anchor: ""
source_lines: [42, 86]
sha256: e02d8a64f1d83fbdac54caa4cddd409492631d26e68230219ac8d4a8a85fc0fb
---

# k2-horizon-abou-dabi-publie-6-ia-ouvertes-375-md-2026

Sur le papier, Kimi K3 dépasse largement K2 Horizon 375B-A23B en nombre de paramètres bruts. Mais la comparaison n’est pas si simple : IFM mise sur la transparence totale (poids, données, code, recettes) là où Moonshot AI publie surtout des poids et une documentation technique plus classique. Les deux approches répondent à des besoins différents, celle d’IFM s’adressant davantage aux chercheurs et institutions qui veulent auditer et reproduire un modèle de bout en bout.

## Comparatif face aux autres familles de modèles ouverts

Le marché des modèles à poids ouverts s’est considérablement densifié en 2026. Rien qu’au mois d’août, neuf lancements distincts et vérifiables ont eu lieu, provenant de six laboratoires différents, dont trois avec des poids ouverts. Voici comment K2 Horizon se positionne face aux principales alternatives disponibles début septembre 2026.

| Famille | Éditeur | Taille (total / actifs) | Contexte | Licence | 
|---|---|---|---|---|
| K2 Horizon 375B-A23B | IFM (MBZUAI) | 375 Md / 23 Md | 524 000 tokens | Apache 2.0 | 
| Kimi K3 | Moonshot AI | 2,8 billions | 1 000 000 tokens | Modified MIT | 
| DeepSeek V4 Pro-0813 | DeepSeek | 1,7 billion | 1 000 000 tokens | MIT | 
| Qwen3.8-27B | Alibaba | 27 Md dense | 262 144 tokens (ext. 1M) | Apache 2.0 | 
| GLM-5.3 | Zhipu AI | 753 Md | 1 000 000 tokens | MIT | 
| Mistral Large 3 | Mistral AI | Non communiqué | 1 000 000 tokens | Apache 2.0 | 

Ce tableau illustre une tendance de fond : sur le seul critère du nombre de paramètres, K2 Horizon 375B-A23B reste modeste face aux mastodontes chinois. Mais IFM ne joue pas la même partition. L’institut mise sur l’amplitude de la gamme (six tailles couvrant tout le spectre matériel, d’une montre à un cluster de serveurs) et sur la profondeur de la divulgation, deux arguments que ni Moonshot AI ni DeepSeek ne mettent autant en avant. Pour situer ces choix dans le paysage plus large des modèles ouverts récents, notre comparatif Qwen3.8-Max et notre analyse du trio DeepSeek V4, Qwen3.8 Max et GLM-5.3 détaillent les rapports de force actuels.

## Disponibilité, outils et absence de tarification officielle

Les six modèles K2 Horizon sont disponibles en téléchargement libre sur Hugging Face, avec un support dès le premier jour pour les frameworks d’inférence vLLM, SGLang et Ollama. Cette compatibilité immédiate avec les outils les plus utilisés par la communauté open source facilite grandement l’adoption par les équipes techniques qui veulent tester rapidement les modèles sur leur propre infrastructure.

Aucune tarification d’API commerciale de premier niveau n’a été communiquée par IFM au moment du lancement. Le modèle économique repose donc, pour l’instant, sur l’auto-hébergement : les entreprises et les chercheurs qui veulent utiliser K2 Horizon doivent déployer les modèles sur leur propre infrastructure ou passer par un fournisseur d’inférence tiers, plutôt que de payer un tarif au token facturé directement par IFM. Pour une entreprise française soumise à des contraintes de résidence des données, cette absence de dépendance à une API propriétaire hébergée à l’étranger peut constituer un argument commercial en soi.

## Ce que disent les responsables d’IFM

Eric Xing, président de MBZUAI et fondateur de l’Institute of Foundation Models, a défendu la philosophie du lancement en expliquant : “K2 Horizon is fully open, making it possible for other researchers, developers and institutions to develop and inspect code, training data and methodology” (“K2 Horizon est entièrement ouvert, ce qui permet à d’autres chercheurs, développeurs et institutions de développer et d’inspecter le code, les données d’entraînement et la méthodologie”), selon des propos rapportés par The National.

Hector Liu, directeur du laboratoire de recherche d’IFM dans la Silicon Valley, a précisé la logique de conception derrière chaque taille de modèle : “Every model is built to compete with the best open models at its size, and everyone ships with the weights, code, training data and methodology behind it” (“Chaque modèle est conçu pour rivaliser avec les meilleurs modèles ouverts de sa catégorie de taille, et chacun est livré avec les poids, le code, les données d’entraînement et la méthodologie qui les sous-tendent”), également cité par The National.

Ces deux prises de parole insistent sur le même point : IFM ne cherche pas à revendiquer le titre de meilleur modèle toutes catégories confondues, mais à imposer un standard de transparence que peu d’acteurs de cette taille acceptent aujourd’hui d’appliquer. Une analyse détaillée du lancement est également disponible sur MarkTechPost.

## Le contexte : une cadence de sorties devenue effrénée

Pour comprendre pourquoi K2 Horizon fait parler de lui, il faut le replacer dans le rythme de sorties de modèles observé depuis le printemps 2026. Entre fin mai et début juin, plusieurs modèles majeurs sont sortis coup sur coup : MiniMax M2.5, M2.7 et M3 Highspeed, Nvidia Nemotron 3 Ultra 550B et Google Gemma 4 12B, portant la cadence à un “LLM majeur” toutes les 48 heures selon certains observateurs spécialisés. Le mois d’août 2026 a confirmé cette tendance avec neuf lancements distincts, incluant GLM-5.3, Qwen3.8-27B, Gemini 3.7 Flash, Grok 4.6 et Qwen3.8-Max.

Dans ce contexte de surproduction de modèles, la stratégie d’IFM tranche par son absence de course à la taille brute. L’institut préfère miser sur une gamme complète et une documentation exhaustive, un pari qui rappelle les débuts de la vague open source occidentale mais appliqué à une échelle et une transparence inédites.

## Un signal fort pour le débat français sur l’IA souveraine

La présence d’un laboratoire IFM à Paris change la donne pour les acteurs français qui cherchent des alternatives crédibles aux modèles fermés américains. Les nouvelles obligations européennes visent en priorité les modèles les plus récents disponibles en France, à savoir GPT-5.6, Claude Opus 5, DeepSeek V4-Pro et les nouveaux modèles de Google, tous désormais dans le viseur réglementaire de Bruxelles. Dans ce climat, un modèle entièrement ouvert et documenté comme K2 Horizon, produit par un institut ayant pignon sur rue à Paris, coche plusieurs cases recherchées par les administrations et les entreprises soucieuses de conformité et de traçabilité des données d’entraînement.

Cela dit, les sources disponibles ne montrent pas encore d’adoption confirmée par un ministère français, une institution européenne ou une grande entreprise du CAC 40. Ce qui est documenté avec certitude, c’est la présence opérationnelle d’IFM sur le sol français et une empreinte de lancement mondiale, pas encore des métriques d’adoption concrètes en Europe. Le parallèle avec LLMs4Europe, le programme doté de 20 millions d’euros lancé par la Commission européenne fin août 2026 avec plus de 70 partenaires, est frappant : les deux initiatives misent sur l’ouverture complète comme réponse à la dépendance technologique, mais l’une est un consortium public européen quand l’autre est un institut de recherche financé depuis le Golfe.

## Impact sur le marché des modèles ouverts européens

