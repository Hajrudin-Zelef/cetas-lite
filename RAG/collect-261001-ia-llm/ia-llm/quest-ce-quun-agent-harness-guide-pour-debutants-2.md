---
id: collect-261001-ia-llm/ia-llm/quest-ce-quun-agent-harness-guide-pour-debutants-2
title: "quest-ce-quun-agent-harness-guide-pour-debutants"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft", "OpenAI"]
dates: []
keywords: ["agent", "agents", "aws", "bedrock", "claude", "gemini", "mcp", "open source", "sandbox", "valuation"]
source: docs/RAG/collect-261001-ia-llm/quest-ce-quun-agent-harness-guide-pour-debutants.md
source_anchor: ""
source_lines: [81, 161]
sha256: 4825c2489ad3f0e660a6917a6318d07b624460c796dde5fc2a98f6d4d9184c2d
---

# quest-ce-quun-agent-harness-guide-pour-debutants

Un runtime d’agent est la couche qui aide un agent à s’exécuter de manière fiable dans le temps. Il gère l’exécution durable, la persistance de l’état, les relances, l’humain dans la boucle et le streaming. LangGraph, Temporal et Inngest en sont des exemples. Harrison Chase propose cette analogie : si Node.js est le runtime et Express le framework, un harness ressemble à Next.js.

### Qu’est-ce qui distingue un harness ?

Un harness opère à plus haut niveau qu’un framework. Là où un framework fournit des composants, un harness arrive généralement avec davantage de choix déjà posés : outils, planification, accès au système de fichiers et gestion du contexte.

## Cas d’usage d’un agent harness : code, recherche, data et entreprise

On retrouve les mêmes briques sur des métiers très différents, mais c’est la combinaison qui change. Un agent de développement et un agent de workflow en entreprise ont tous deux besoin d’un harness, mais n’en sollicitent pas les mêmes parties. Ces catégories ne sont pas des standards formels, ce sont des manières pratiques de voir comment une même idée s’adapte au travail à réaliser.

### Harness pour agents de développement

Les agents de développement sont un bon exemple actuel, car le harness y est très visible. Pour produire un travail utile, un agent a besoin d’accès fichiers, de contexte git, d’exécution terminal, de lancement de tests, d’installation de dépendances et de règles projet. Claude Code et Codex illustrent ce schéma : ils reposent fortement sur du code de harness, pas sur un simple appel d’API à un modèle.

La différence entre un bon et un moyen harness de développement se voit dans les détails : comment il se remet d’un test échoué, s’il peut annuler une mauvaise modification, ou la propreté avec laquelle il expose l’historique git au modèle. C’est là que se concentre l’essentiel de l’effort d’ingénierie.

### Harness pour agents de recherche

Les agents de recherche ont besoin d’un autre arsenal : recherche web, suivi des sources, prise de notes, gestion des citations et synthèse. Le harness gère la façon de stocker les résultats, d’attribuer les sources, et de découper de longs documents pour éviter de saturer la fenêtre de contexte d’un seul coup.

### Harness pour agents d’analyse de données

Les agents data ont besoin d’accès à des jeux de données, à des bases SQL, à des environnements Python et au contexte de schéma pour connaître tables et colonnes disponibles avant d’écrire des requêtes. Le harness applique aussi des frontières de permission, essentielles quand l’agent peut toucher des données de production.

### Harness pour workflows d’entreprise

Les déploiements en entreprise ajoutent une couche d’exigences : authentification, journaux d’audit, circuits d’approbation, contrôle d’accès par rôle et connexions aux systèmes internes. AWS AgentCore en est un exemple managé, avec identité, réseau VPC et observabilité intégrés. Microsoft Agent Framework couvre des besoins similaires pour les équipes sur Azure ou dans des environnements .NET.

## Outils pour construire des systèmes d’agent harness en 2026

Une poignée de produits reviennent le plus souvent à mi-2026. Ils se situent à différents niveaux du spectre framework–runtime–harness, et les frontières évoluent encore.

### LangChain Deep Agents

LangChain Deep Agents est le harness open source de LangChain, construit sur LangGraph comme runtime. Il inclut un outil de planification, un système de fichiers virtuel, le lancement de sous-agents, la compression automatique du contexte et un middleware pour l’approbation humaine et la détection de données personnelles. Agnostique au modèle, il prend en charge les endpoints compatibles OpenAI et se connecte à des bacs à sable comme Modal, Runloop et Daytona pour l’exécution de code.

### Anthropic Agent SDK

L’Anthropic Agent SDK (nom du package : `claude-agent-sdk`) a été extrait de Claude Code et publié en option autonome. Il inclut une boucle d’agent intégrée, des outils pour l’exécution bash, la lecture/écriture de fichiers, la recherche web, l’intégration MCP et la compaction du contexte. Il fonctionne uniquement avec les modèles Claude, via l’API d’Anthropic, Amazon Bedrock, Vertex AI et Azure.

### OpenAI Agents SDK

Comme mentionné plus haut, OpenAI Agents SDK a franchi le cap du framework vers le harness à mesure que ses fonctionnalités ont grandi. La version d’avril 2026 a ajouté l’exécution en sandbox native, la compaction de mémoire et des outils système de fichiers. Disponible en Python et TypeScript, le SDK gère l’usage d’outils, les handoffs entre agents et les garde-fous.

### Google Agent Development Kit

Google ADK prend en charge l’orchestration multi-agents avec des classes intégrées pour des structures séquentielles, parallèles et en boucle. Il inclut des outils d’évaluation, fonctionne avec Vertex AI pour les déploiements managés et supporte MCP pour la connexion aux outils. Disponible en Python, Java, TypeScript et Go, il est optimisé pour les modèles Gemini mais se veut agnostique au modèle.

### Microsoft Agent Framework

Microsoft Agent Framework est la voie de migration actuelle de Microsoft pour les projets AutoGen. Il prend en charge Python et .NET, fonctionne avec Azure AI et inclut la prise en charge MCP pour la connectivité des outils.

### CrewAI

CrewAI adopte une approche basée sur les rôles pour les systèmes multi-agents. Vous définissez des agents avec des rôles spécifiques, assignez des tâches, configurez des équipes et déclarez mémoire et garde-fous. Il convient aux problèmes qui se mappent naturellement à une équipe de spécialistes.

### Temporal et Inngest

Ce ne sont pas des agent harness à eux seuls. Ce sont des plateformes d’exécution durable qui gèrent ce qui se passe lorsqu’une tâche d’agent doit tourner des heures ou des jours sans perdre l’état. En cas d’échec, le moteur rejoue depuis le dernier point de contrôle réussi plutôt que de tout recommencer.

## Défis et arbitrages avec un agent harness

Ajouter un harness élargit les capacités du système, mais chaque outil, permission et agent supplémentaire ouvre une nouvelle voie de défaillance. À mesure que les tâches s’allongent, les garde-fous, le traçage et l’état durable cessent d’être optionnels et deviennent ce qui rend un long run récupérable.

Il existe aussi un risque de couplage qui surprend les équipes. LangChain a rapporté une hausse de 10 à 20 points sur un sous-ensemble de tau2-bench après l’ajout de profils de harness spécifiques au modèle. Artificial Analysis va dans le même sens dans son Coding Agent Index : les résultats des agents de développement dépendent du modèle et du harness ensemble, avec des variations marquées de coût, de jetons et de temps par tâche selon les combinaisons. Le modèle n’a pas changé. Les prompts, outils et middleware autour, si. Et ce profilage relève du travail de harness.

## Avez-vous vraiment besoin d’un agent harness ?

Voici une façon directe d’évaluer si vous en avez besoin.

Vous avez probablement besoin d’un harness si votre système remplit une ou plusieurs de ces conditions :

- Il doit utiliser des outils externes
- Il doit se souvenir des avancées entre les sessions
- Il doit exécuter du code dans un environnement réel
- Il coordonne plus d’un agent
- Il doit se remettre de défaillances partielles sans perdre le travail
- Il requiert une approbation humaine

Vous n’avez probablement pas besoin d’un harness si la tâche est un workflow prévisible où chaque étape est prédéfinie.

