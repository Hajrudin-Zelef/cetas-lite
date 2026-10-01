---
id: collect-261001-ia-llm/ia-llm/llms4europe-70-partenaires-20-m-2026-3
title: "llms4europe-70-partenaires-20-m-2026"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Mistral"]
dates: []
keywords: ["amd", "attention", "fine-tuning", "mistral", "valuation"]
source: docs/RAG/collect-261001-ia-llm/llms4europe-70-partenaires-20-m-2026.md
source_anchor: ""
source_lines: [61, 96]
sha256: 023f958b7b4da9517c222c2501a6ebff9d6e09bc13baff898c07c55a8aebffbc
---

# llms4europe-70-partenaires-20-m-2026

Sur le papier, Mistral et les projets financés par Digital Europe ne sont pas en concurrence directe : l’un vend des modèles propriétaires à des clients privés et publics, les autres produisent des modèles ouverts financés par des fonds publics. Dans les faits, ils se disputent pourtant les mêmes ressources rares, chercheurs qualifiés en traitement du langage, accès prioritaire aux clusters EuroHPC, et surtout l’attention politique des gouvernements nationaux, qui doivent arbitrer entre soutenir un champion privé national et financer un consortium académique paneuropéen. La décision, fin 2025, d’attribuer une part réduite de l’accès à EuroHPC au projet Domyn plutôt qu’à Mistral, un épisode que *Tech Insider* avait également couvert, montre que ces arbitrages sont loin d’être automatiques.

## Contexte historique : la course européenne aux LLM depuis 2023

La séquence LLMs4Europe s’inscrit dans une accélération continue depuis 2023. L’Union a d’abord misé sur la réglementation, avec l’adoption de l’AI Act, avant de comprendre que la seule régulation ne suffirait pas à combler l’écart technologique face aux laboratoires américains et chinois. Le programme Digital Europe a ensuite financé une série de projets de plus en plus ambitieux : TildeOpen LLM, un modèle de 30 milliards de paramètres optimisé pour les langues européennes issu du programme European AI Grand Challenge, puis OpenEuroLLM au premier semestre 2025, puis LLMs4Europe et ALT-EDIC4EU. En parallèle, plusieurs États membres ont lancé leurs propres initiatives nationales, comme le Portugal avec son modèle Amália, financé à hauteur de 5,5 millions d’euros, que *Tech Insider* avait présenté comme en avance sur la France dans ce domaine spécifique.

Ce qui distingue la période actuelle des tentatives précédentes, c’est la tentative de mise en cohérence entre ces initiatives éparses. Plutôt que de laisser chaque État membre financer son propre modèle national sans coordination, la Commission structure désormais un écosystème autour d’ALT-EDIC, avec une répartition des rôles plus lisible entre fondation, adaptation sectorielle et réutilisation institutionnelle. Reste à savoir si cette architecture tiendra la distance face à des laboratoires privés qui lèvent, eux, des dizaines de milliards de dollars en quelques mois.

## Impact sur le marché : PME, chercheurs, administrations et développeurs

Pour une PME française ou européenne, l’intérêt immédiat de LLMs4Europe tient moins au modèle final qu’à la porte d’entrée qu’ouvrent les appels FSTP du printemps 2026. Une entreprise du secteur du tourisme ou de l’énergie disposant de données linguistiques ou métier de qualité pourra candidater pour contribuer au fine-tuning des modèles sectoriels, avec un financement européen à la clé, sans avoir à monter un projet de recherche fondamentale complet. C’est un point d’entrée nettement plus accessible que les grands appels à projets EuroHPC, historiquement réservés à des consortiums académiques lourds.

Pour les développeurs et les équipes techniques, l’impact dépendra largement de la politique de licence retenue pour les modèles issus du fine-tuning, un point que les documents publics consultés ne précisent pas encore. Si LLMs4Europe suit la philosophie d’ouverture revendiquée par OpenEuroLLM, avec des données et des poids pleinement inspectables, les équipes pourront potentiellement héberger localement des modèles adaptés aux secteurs de l’énergie ou des télécoms sans dépendre d’une API facturée à l’usage, un argument déjà mis en avant par les promoteurs du projet frère. Pour les administrations publiques, en particulier dans les collectivités locales et les guichets citoyens, cette architecture pourrait offrir une alternative crédible aux solutions américaines dans des usages jugés sensibles, à condition que la performance des modèles sectoriels tienne la comparaison avec les meilleurs modèles commerciaux du marché.

## Comparatif chiffré des grands projets d’IA souveraine européenne

Le tableau suivant synthétise les données officielles disponibles sur les principaux projets financés par l’Union autour des grands modèles de langage, tels que documentés par la Commission européenne et l’ALT-EDIC à la date du 2 septembre 2026.

| Projet | Coordination | Budget | Partenaires | Focus principal | 
|---|---|---|---|---|
| LLMs4EU | ALT-EDIC | 20 M€ (subvention, cofinancement 50 %) | 70+ organisations | Données linguistiques et fine-tuning sectoriel (5 secteurs) | 
| ALT-EDIC4EU | ALT-EDIC | 4 M€ (action de coordination, 100 %) | Écosystème ALT-EDIC | Coordination de l’infrastructure linguistique européenne | 
| OpenEuroLLM | Université Charles (Jan Hajič) / AMD Silo AI (Peter Sarlin) | 37,4 M€ dont 20,65 M€ de fonds UE | 20 institutions | Modèles de fondation multilingues, toutes langues officielles UE | 
| LLM Institutionnel | DG Traduction, Commission européenne | Non communiqué | Institutions et entités juridiques UE | Modèle interne réutilisable, publié le 16 juillet 2026 | 
| Écosystème combiné | ALT-EDIC | ~80 M€ + ressources EuroHPC | 90+ organisations | Ensemble de la chaîne, de la fondation à l’usage sectoriel | 

Ce qui ressort de cette comparaison, c’est la dissymétrie assumée entre les projets : OpenEuroLLM concentre le budget individuel le plus élevé pour l’entraînement de modèles de fondation, tandis que LLMs4Europe mobilise un budget comparable mais orienté vers l’aval, la donnée et l’adaptation métier. Aucun de ces montants, pris isolément, ne rivalise avec les investissements des laboratoires américains sur un seul modèle frontière, ce qui explique pourquoi la Commission insiste sur la complémentarité et la mutualisation des ressources EuroHPC plutôt que sur la course à la taille du modèle.

## Limites, risques et zones d’ombre du projet

Plusieurs éléments restent flous à ce stade et méritent d’être suivis dans les prochains mois. D’abord, la liste complète des 70 partenaires n’est pas encore publique, ce qui empêche d’évaluer précisément la répartition géographique du consortium et le poids réel de chaque pays membre. Ensuite, ni le nombre exact de langues couvertes, ni la taille en paramètres des modèles fine-tunés, ni le volume de ressources de calcul dédiées spécifiquement à LLMs4Europe (par opposition à l’enveloppe EuroHPC partagée avec OpenEuroLLM) n’ont été communiqués dans les documents consultés. Ce niveau de flou est habituel à ce stade d’un projet européen pluriannuel, mais il complique toute évaluation sérieuse de l’ambition technique réelle par rapport aux moyens engagés.

Autre point de vigilance : aucune réaction publique d’analystes indépendants, de laboratoires concurrents ou de la presse technologique française n’a encore émergé sur ce projet spécifique au moment de la publication de cet article, ce qui distingue LLMs4Europe de sujets plus médiatisés comme les grands modèles commerciaux. Un projet financé à 20 millions d’euros et porté par un consortium administratif, aussi large soit-il, reste structurellement moins agile qu’un laboratoire privé pour itérer rapidement sur ses modèles, un écart que la Commission elle-même reconnaît implicitement en misant sur la mutualisation plutôt que sur la vitesse pure.

## Cinq prédictions pour la suite (2026-2028)

