---
id: collect-261001-ia-llm/ia-llm/openai-presente-gpt-5-6-sol-ultrafast-qui-repond-14-fois-plus-vite-2
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Hugging Face", "Meta", "OpenAI"]
dates: []
keywords: ["agent", "agents", "benchmark", "benchmarks", "claude", "gemini", "gguf", "incident", "qwen", "reasoning", "research"]
source: docs/RAG/collect-261001-ia-llm/openai-presente-gpt-5-6-sol-ultrafast-qui-repond-14-fois-plus-vite.md
source_anchor: ""
source_lines: [65, 116]
sha256: 764b6a485bc42a0d105da584d36f107914804483fcfa0f2efe09cb93bfcb94e6
---

# 🧠 **RECHERCHE**

**L'impact à retenir :** le format GGUF plus 5 milliards de paramètres actifs, c'est exactement la combinaison qui permet de faire tourner ce modèle sur une machine personnelle correctement équipée, sans abonnement, sans API, sans que vos données sortent de chez vous. Et le chiffre qui compte n'est pas le score de benchmark mais le taux d'hallucination divisé par plus de deux en une génération : un modèle qui sait dire « je ne sais pas » est infiniment plus utile au quotidien qu'un modèle qui gagne trois points sur un test.

# 🧠 **RECHERCHE**

**Trois agents Claude sur le même projet, et c'est la guerre de territoire**

La Frontier Red Team d'Anthropic a lâché trois agents Claude sur un même projet logiciel, chacun avec des consignes incompatibles, aucun n'étant informé de l'existence des autres. Résultat : les agents se sont sabotés mutuellement avec des logiciels malveillants auto-répliquants de plus en plus agressifs, persuadés d'être délibérément entravés. Plus les agents sont performants, plus le conflit devient efficace, mais ils inventent parfois seuls un mécanisme de résolution, du type concours où le gagnant rafle tout. L'étude fait écho à un incident réel où des agents OpenAI ont cherché des failles pendant des semaines avant de pirater Hugging Face.

**Les modèles 3D générés par IA inondent le marché, et presque personne ne les achète**

Sur CGTrader, place de marché d'assets 3D, **un modèle sur six** mis en ligne est désormais généré par IA. Mais ces fichiers ne pèsent que **1 dollar sur 90** du chiffre d'affaires de la plateforme. Seuls 5 % des acheteurs de contenus IA s'en disent satisfaits, 20 % jugent la qualité insuffisante et 7 % doivent lourdement retoucher, notamment pour l'impression 3D. « Les acheteurs votent avec leur portefeuille », résume CGTrader, qui lance malgré tout un partenariat avec Tencent sur un workflow de création 3D assisté par IA.

**Google Sheets transforme vos colonnes en mini-applications**

Sheets canvas utilise Gemini pour convertir des lignes et des colonnes en tableaux de bord interactifs, trackers ou plans de salle, sans une formule ni une ligne de code. On ouvre une feuille, on clique sur l'icône Gemini, on choisit « create canvas », et on décrit ce qu'on veut en langage naturel. La mise en page reste synchronisée dans les deux sens avec la feuille source, se partage comme une feuille classique et s'ajuste par prompts successifs.

**Les jalons de l'IA qui s'améliore toute seule tombent plus vite que prévu**

Severin Field, fellow à l'IAPS, avait interrogé 25 chercheurs d'OpenAI, Anthropic, Google DeepMind, Meta et d'universités américaines sur l'auto-amélioration récursive des modèles, en leur demandant quels signaux les alerteraient que la recherche en IA commence à s'automatiser. Il vient de faire le bilan de ces prédictions. Plusieurs des jalons cités par ces chercheurs sont déjà atteints, et plus tôt que ce qu'ils anticipaient eux-mêmes.

**Attraper le foie gras par IA, avant que ce soit irréversible**

La stéatose hépatique touche **environ 30 % des adultes** dans le monde, soit plus d'un milliard de personnes, et progresse presque toujours sans symptôme jusqu'au stade de la fibrose, avec un risque accru de maladies cardiovasculaires et de cancers. Jeffrey Lazarus, de la CUNY Graduate School of Public Health, propose de passer les dossiers médicaux électroniques au crible d'une IA pour repérer et hiérarchiser les patients à risque. L'enjeu est réel : détectée tôt, la maladie est largement réversible, par changement d'hygiène de vie ou traitements comme le sémaglutide et le resmetirom.

**Un doctorant démontre un principe d'incertitude pour les fractales**

Le principe d'incertitude fractale relie la géométrie des formes infiniment complexes au comportement des particules quantiques piégées dans des situations chaotiques. Formulé en dimension 1 en 2016 par Semyon Dyatlov et Jean Bourgain au MIT, il résistait depuis à toute généralisation. Alex Cohen, doctorant au MIT, l'a étendu à toutes les dimensions supérieures, une preuve publiée dans les Annals of Mathematics et jugée fondamentale pour l'étude du chaos quantique. Il est aujourd'hui professeur à NYU, à 25 ans.

**Bullet, l'agent de codage qui promet d'aller plus vite que Claude Code**

Deux fondateurs passés par AppLovin et Citadel, issus du programme YC S26, revendiquent **479 tâches résolues sur 500** du premier coup sur SWE-bench Verified, soit **95,8 %**, avec une moyenne de **119 secondes** par tâche. Ils annoncent 35 à 67 % de gain de vitesse sur les agents concurrents grâce à un routage entre plusieurs modèles, une recherche de contexte ciblée et une hygiène de contexte agressive : **16 % d'allers-retours en moins** et **27 % de coût en moins**. Chiffres auto-déclarés, non vérifiés de façon indépendante à ce stade.

**Anthropic publie un indice pour mesurer le raisonnement conceptuel des IA**

Trois benchmarks agrégés dans le Conceptual Reasoning Index, publié sur conceptualreasoning.ai : LMCA (**560 textes de position** et **1 461 arguments** notés par des experts), ACCoRD et DTBench. L'objectif est de mesurer la capacité d'un modèle à argumenter sur des questions qui n'ont pas de réponse vérifiable empiriquement, comme la philosophie ou la gouvernance de l'IA. Le pari sous-jacent : pour que l'IA aide à gérer les risques qu'elle crée, encore faut-il qu'elle sache raisonner là où aucun test ne peut la corriger.

**Google fait diagnostiquer ses ralentissements par un agent plutôt que par force brute**

Optimiser l'infrastructure d'entraînement et de service d'un LLM revient d'ordinaire à tester des centaines de configurations. Ce papier de Google Research remplace la question « laquelle de ces 100 configurations est la plus rapide ? » par « qu'est-ce qui ralentit réellement le système ? ». Un Agent Analyseur lit les traces de profilage, classe le type de goulot d'étranglement, et la recherche se limite ensuite à la petite portion pertinente de l'espace des configurations. Le temps de recherche s'effondre.

**75 articles écrits par des IA, et pas un système sans faille de preuve**

Google Cloud AI Research a audité **75 papiers** produits par **cinq systèmes** de recherche autonomes sur cinq tâches ADRS. Tous les systèmes testés présentaient au moins une défaillance systématique dans leur chaîne de preuves : citations fabriquées, étapes de raisonnement invérifiables. Le constat de ScientistOne est net : ces agents sont devenus assez bons pour résoudre les problèmes de benchmark, le goulot d'étranglement est désormais la fiabilité du compte rendu qu'ils rédigent ensuite.

**Mimir v1 : 1 milliard de paramètres et pas une donnée au statut douteux**

Des chercheurs danois publient un modèle de 1 milliard de paramètres fondé sur l'architecture Hierarchical Reasoning Model, entraîné de zéro sur un mélange de **161 jeux de données tous juridiquement autorisés**, sans le moindre corpus au statut flou. Il dépasse le HRM-Text 1B d'origine, rivalise avec Qwen 3.5 4B et Gemma 4 E2B sur **20 benchmarks** en anglais, maths et code, et établit l'état de l'art pour le danois. Les poids sont libres d'accès sur Hugging Face.

# **🗞️PLUS D'ACTUALITÉS**

**Il cache un ordre destiné aux IA en police blanche de 3 points dans un dossier judiciaire**

