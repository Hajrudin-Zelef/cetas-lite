---
id: collect-261001-ia-llm/ia-llm/guerre-des-prix-ia-2026-gpt-5-6-80-opus-5-kimi-k3-1
title: "Estimation simplifiee du cout mensuel (500M tokens entree, 100M tokens sortie)"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Moonshot", "OpenAI"]
dates: []
keywords: ["benchmarks", "claude", "deepseek", "fable 5", "gemini", "gpt-5.6", "kimi", "luna", "open source", "opus 4", "opus 5", "sol"]
source: docs/RAG/collect-261001-ia-llm/guerre-des-prix-ia-2026-gpt-5-6-80-opus-5-kimi-k3.md
source_anchor: ""
source_lines: [1, 34]
sha256: 7b2fbc40748ccc3902b84422d22de0f41164e8ee70d295bf603aff287f2a575c
---

# Estimation simplifiee du cout mensuel (500M tokens entree, 100M tokens sortie)

En l’espace de dix jours, entre le 21 et le 31 juillet 2026, quatre des plus grands noms de l’intelligence artificielle générative ont revu leurs grilles tarifaires à la baisse. OpenAI a coupé le prix de GPT-5.6 Luna de 80 %, Google a lancé une nouvelle gamme Flash agressive, Moonshot AI a ouvert les poids de Kimi K3 et DeepSeek a mis à jour son modèle V4-Flash sans toucher à son prix déjà cassé. Cette guerre des prix IA, la plus intense depuis le lancement de GPT-4 en 2023, redessine l’équation économique de millions de développeurs, d’entreprises et d’agences européennes qui bâtissent des produits sur ces modèles. Voici ce qui s’est passé, pourquoi cela arrive maintenant, et ce que cela signifie pour la France et le reste de l’Europe.

## Ce qui s’est passé : la chronologie de l’été 2026

La séquence a démarré le 9 juillet 2026, quand OpenAI a lancé publiquement la famille GPT-5.6, déclinée en trois niveaux : Sol (le modèle phare), Terra (le milieu de gamme) et Luna (l’entrée de gamme rapide). Trois semaines plus tard à peine, le 30 juillet, l’entreprise a annoncé une baisse de prix spectaculaire sur deux des trois modèles. Entre ces deux dates, Anthropic a lancé Claude Opus 5 le 24 juillet, Google DeepMind a dévoilé Gemini 3.6 Flash et Gemini 3.5 Flash-Lite le 21 juillet, Moonshot AI a ouvert les poids de Kimi K3 le 27 juillet, et DeepSeek a mis en ligne V4-Flash 0731 le 31 juillet. Cinq lancements ou révisions tarifaires majeures en l’espace de trois semaines : la fréquence à elle seule illustre l’intensité de la compétition actuelle sur le marché des grands modèles de langage.

Ce calendrier resserré n’est pas un hasard. Chaque fournisseur observait les annonces des autres avant d’ajuster sa propre grille. Selon des informations reprises par The Verge, OpenAI aurait envisagé dès juin 2026 des baisses de prix agressives pour freiner la migration de sa clientèle entreprise vers Anthropic, qui gagnait du terrain sur les cas d’usage de codage agentique.

## GPT-5.6 Luna perd 80 % de son prix : le détail des chiffres

Le geste le plus spectaculaire de cette guerre des prix IA vient d’OpenAI. Le 30 juillet 2026, l’entreprise a annoncé que GPT-5.6 Luna, son modèle d’entrée de gamme optimisé pour la vitesse, passait de 1,00 $ à 0,20 $ par million de tokens en entrée, et de 6,00 $ à 1,20 $ par million de tokens en sortie. Une baisse de 80 % sur toute la ligne, confirmée par plusieurs médias dont CNBC. GPT-5.6 Terra, le niveau intermédiaire, a lui aussi baissé, mais de façon plus modeste : de 2,50 $/15,00 $ à 2,00 $/12,00 $ par million de tokens, soit une réduction de 20 %. Le modèle phare, GPT-5.6 Sol, conserve son tarif initial de 5 $ en entrée et 30 $ en sortie par million de tokens, avec en complément un nouveau mode « Fast » facturé deux fois plus cher pour une vitesse environ 2,5 fois supérieure, sans gain d’intelligence supplémentaire.

OpenAI a justifié cette baisse par des gains d’efficacité obtenus lors de l’entraînement de GPT-5.6, expliquant notamment que le modèle avait contribué à réécrire et optimiser son propre code d’inférence en production. Une explication technique plausible, mais qui occulte la vraie raison structurelle : une vague de modèles ouverts, moins chers, qui grignotent des parts de marché sur les usages à faible marge. Cette baisse concerne directement les développeurs français qui utilisent l’API OpenAI pour des tâches de classification, de résumé ou de support client à fort volume, où le coût par requête pèse lourd dans l’équation économique.

## Claude Opus 5 : Anthropic mise sur la puissance, pas la baisse de prix

Face à cette agressivité tarifaire, Anthropic a choisi une stratégie différente. Claude Opus 5, lancé le 24 juillet 2026, conserve exactement le même tarif que son prédécesseur Opus 4.8 : 5 $ par million de tokens en entrée et 25 $ par million de tokens en sortie, comme le confirme la page de tarification officielle d’Anthropic. Plutôt que de baisser le prix affiché, l’entreprise a préféré augmenter la capacité livrée pour ce même prix : Opus 5 dépasse nettement les performances d’Opus 4.8 sur les benchmarks de codage et de raisonnement, tout en restant deux fois moins cher que Claude Fable 5, le modèle le plus premium du catalogue, facturé 10 $/50 $ par million de tokens.

Anthropic complète cette offre avec plusieurs paliers tarifaires : le traitement par lots (batch) tombe à 2,50 $/12,50 $ par million de tokens, le mode rapide grimpe à 10 $/50 $, et la mise en cache des tokens d’entrée descend à seulement 0,50 $ par million. Cette segmentation fine permet à l’entreprise de défendre sa marge sur le prix affiché tout en offrant, dans les faits, des réductions substantielles aux clients capables d’adapter leur architecture technique. C’est une manière d’entrer dans la guerre des prix IA sans jamais annoncer de baisse de prix officielle, une nuance stratégique qui compte pour les équipes qui budgétisent leurs déploiements sur plusieurs trimestres. Plus de détails dans notre analyse complète du classement de Claude Opus 5.

## Kimi K3 : Moonshot AI ouvre ses poids et casse les prix

Le troisième front de cette bataille tarifaire vient de Chine. Moonshot AI a mis à disposition les poids complets de Kimi K3 le 27 juillet 2026, sous une licence permissive qui permet un hébergement local ou chez un fournisseur cloud tiers. Sur l’API officielle, telle que documentée sur la plateforme Moonshot, le modèle est facturé 3,00 $ par million de tokens en entrée (0,30 $ seulement en cas de lecture depuis le cache), et 15,00 $ par million de tokens en sortie, pour une fenêtre de contexte allant jusqu’à 1 million de tokens sans surcoût lié à la longueur du contexte.

Ce positionnement, proche du tarif de Claude Sonnet 5, mais avec des poids ouverts, place Kimi K3 dans une catégorie particulière : celle des modèles capables de rivaliser avec les meilleures offres fermées sur les tâches de codage agentique et de navigation web, tout en offrant l’option d’un hébergement souverain. Pour les entreprises européennes soumises au RGPD et soucieuses de maîtriser la localisation de leurs données, cette ouverture des poids change la donne, même si Kimi K3 reste qualifié de modèle « à poids ouverts » plutôt que pleinement open source, faute de publication du code d’entraînement et des données.

## DeepSeek V4-Flash 0731 : la pression open source continue

Le 31 juillet 2026, DeepSeek a publié une mise à jour de son modèle V4-Flash, baptisée 0731 en référence à sa date de sortie. Le tarif reste inchangé par rapport à la version précédente : environ 0,14 $ par million de tokens en entrée (0,0028 $ seulement pour les tokens lus depuis le cache, soit une remise d’environ 98 %) et 0,28 $ par million de tokens en sortie, selon la documentation tarifaire officielle de DeepSeek. Ce qui change, c’est la capacité : la nouvelle version améliore sensiblement les performances sur les tâches agentiques et l’usage d’outils, sans que le prix augmente d’un centime.

À moins d’un tiers du coût des modèles « pro » propriétaires les plus proches en performance, DeepSeek V4-Flash 0731 accentue la pression déjà exercée sur les acteurs américains. C’est précisément ce type de modèle, cité par plusieurs analystes comme la véritable cause de la baisse de GPT-5.6 Luna, qui pousse désormais OpenAI, Anthropic et Google à revoir leurs paliers d’entrée de gamme pour ne pas perdre les usages à fort volume et faible marge unitaire. Pour aller plus loin, voir notre comparatif DeepSeek V4 vs Claude Fable 5 vs GPT-5.6.

## Gemini 3.6 Flash et 3.5 Flash-Lite : Google verrouille le milieu de gamme

