---
id: collect-261001-ia-llm/ia-llm/frameworks-dagents-ia-des-outils-pour-des-systemes-plus-intelligents-1
title: "frameworks-dagents-ia-des-outils-pour-des-systemes-plus-intelligents"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["agent", "agents", "open source"]
source: docs/RAG/collect-261001-ia-llm/frameworks-dagents-ia-des-outils-pour-des-systemes-plus-intelligents.md
source_anchor: ""
source_lines: [1, 100]
sha256: 9519a49b3570967a0e9820d3d8f3119132d7355e75f97489be868ce5f2b15d22
---

# frameworks-dagents-ia-des-outils-pour-des-systemes-plus-intelligents

Cours

Imaginez un collaborateur qui ne dort jamais, ne se plaint pas et s’améliore en continu. De la science-fiction ? C’est précisément ce que les agents IA apportent. Ces cerveaux numériques planifient des tâches, récupèrent des informations, dialoguent avec les utilisateurs, se connectent à des API — et peuvent même coopérer entre eux pour résoudre des problèmes d’envergure.

Bien sûr, vous pourriez en construire un de A à Z, mais la plupart des équipes s’appuient sur des frameworks d’agents IA : des boîtes à outils prééquipées qui gèrent la mémoire, l’orchestration — bref, l’essentiel. Dans cet article, nous expliquons ce que sont ces frameworks, comment ils fonctionnent et comment choisir celui qui convient à votre cas d’usage.

## Introduction aux agents d'intelligence artificielle

## Qu’est-ce qu’un framework d’agent IA ?

Au fond, les agents IA sont des programmes capables de percevoir, planifier et agir. Conçus pour analyser un objectif, le décomposer en étapes et passer à l’action — seuls ou à plusieurs.

Qu’il s’agisse de répondre à une question, lancer une recherche ou collaborer avec un autre agent, ils sont utilisés dans de nombreuses industries et opèrent avec un niveau d’autonomie surprenant.

Il existe plusieurs façons de construire un agent IA à partir de zéro. On peut les développer en Python, avec React ou d’autres stacks. Mais repartir de la feuille blanche, c’est une toute autre histoire : il faut gérer la mémoire, planifier les tâches, brancher les outils, coordonner, traiter les erreurs… et bien plus encore. Autant de pièces à assembler sans faux pas.

#### Frameworks d’agents IA = cerveau + boîte à outils

Ces frameworks facilitent grandement la création et la mise à l’échelle des agents IA.

*Composants d’un framework d’agent IA. Source : Napkin IA*

Ces frameworks apportent :

- Architecture : une structure d’interaction entre agents.
- Mémoire : rappel à court et à long terme.
- Modèles : les grands modèles de langage (LLMs) sont le cœur des agents, leur donnant compréhension, raisonnement et capacité d’action.
- Boîtes à outils : APIs, moteurs de recherche, interpréteurs — tout le nécessaire pour agir.
- Couche d’orchestration : coordination des tâches et de la collaboration (notamment en multi-agents).
- Intégrations : connexion à LangChain, OpenAI, Azure, Slack, etc.

En bref : ils vous font gagner du temps, vous épargnent des migraines et réduisent la complexité.

## Comment fonctionnent les agents IA en pratique

Pour comprendre un agent IA, partons de son cycle de décision le plus simple.

Tout commence par la planification des tâches. Face à une demande — « Résumez cet article » ou « Trouvez un vol à moins de 1 000 $ » — l’agent la décompose en étapes. Il se crée une to-do list pour décider quoi faire, dans quel ordre, et s’il peut tout gérer seul ou doit solliciter d’autres agents.

*Flux de travail d’un agent IA. Source : Napkin IA*

Vient ensuite le function calling, où l’agent choisit les outils ou APIs à mobiliser. Qu’il s’agisse de naviguer sur le web, consulter la météo ou interroger une base, il appelle la bonne fonction — un peu comme vous alternez entre applications pour boucler votre liste de tâches.

Puis vient la phase d’exécution, où l’agent réalise l’action : exécuter du code, récupérer des données, envoyer un e-mail ou rédiger une réponse. Il interagit avec ses outils et systèmes pour mener à bien la mission.

Enfin, la boucle de feedback. Qu’il ait réussi, partiellement réussi ou échoué, l’agent évalue le résultat, met à jour sa mémoire, ajuste ses prochaines étapes ou demande des précisions à l’utilisateur. Ce retour d’expérience l’affûtte au fil du temps.

Selon la complexité, l’agent parcourt ce cycle une fois ou itère jusqu’au bon résultat.

### Systèmes mono-agent et multi-agents

Les agents IA se classent selon leurs rôles, capacités et environnements. Une distinction clé tient au nombre d’agents impliqués.

Les systèmes mono-agent opèrent de manière indépendante pour un objectif précis. Ils s’appuient sur des outils et ressources externes pour agir efficacement dans divers environnements. Idéals pour des buts bien définis qui ne nécessitent pas de coordination. En général, un seul modèle fondationnel suffit.

Les systèmes multi-agents mobilisent plusieurs agents qui collaborent ou entrent en compétition pour des objectifs partagés ou individuels. Ils capitalisent sur des compétences et rôles variés, parfaits pour les problèmes complexes. Ils peuvent aussi simuler des comportements proches de l’humain (communication, interactions). Chaque agent peut reposer sur un modèle fondationnel différent, adapté à sa fonction.

### Fonctionnalités communes des frameworks d’agents IA

La plupart des frameworks partagent un socle de fonctionnalités : le minimum vital pour que les agents livrent des résultats.

*Fonctionnalités communes des frameworks d’agents IA. Source : auteur*

D’abord, la mémoire persistante, qui permet de conserver le contexte. Plutôt que de tout réinitialiser, l’agent capitalise sur l’historique, ce qui le rend plus pertinent, réactif et presque « humain ».

Ensuite, la génération augmentée par recherche (RAG). En clair, l’agent va chercher en temps réel la bonne information dans des sources externes : documents, bases, web. Il ne se limite plus à son entraînement et s’appuie sur un savoir à jour et spécifique au domaine.

Puis l’usage d’outils, qui fait toute la différence. Un agent n’est pas seulement là pour répondre : il agit. Appeler une API, faire des calculs, scraper un site, déclencher un job backend : c’est ce qui transforme un simple chatbot en assistant opérationnel.

Enfin, la collaboration entre agents. Les agents se partagent des tâches, se transmettent des mises à jour ou résolvent chacun une partie d’un problème global. Comme dans une entreprise : l’un fait la recherche, l’autre rédige le rapport. À plusieurs, ils traitent des missions qu’un seul agent ne pourrait pas mener.

## Frameworks d’agents IA populaires

L’offre évolue vite, et choisir le bon framework peut ressembler à un choix de garniture sur une pizza : il y en a pour tous les goûts, mais trop d’options, ça déroute.

Simplifions : voici les frameworks qui font le plus parler d’eux, comment ils fonctionnent, leurs atouts et leurs terrains de jeu idéaux.

*Principales options de frameworks d’agents IA : source : Napkin IA*

### 1. CrewAI

CrewAI, c’est comme monter votre équipe de super-héros : chaque agent a un rôle (« Researcher », « Writer », « Analyst ») et travaille en équipe.

L’approche orientée rôle assigne les tâches à chaque agent, et le système se charge de la coordination.

Ses atouts :

- Prise en main simple : il masque la complexité d’ingénierie pour vous laisser vous concentrer sur les résultats.
- Open source, compatible avec des LLMs majeurs (OpenAI, Anthropic) et RAG intégré.

Idéal pour les entreprises qui déploient des bots de service client ou automatisent des flux de recherche intensive.

Limite : ça peut sembler contraint si vous avez besoin d’agents très flexibles dans leurs rôles, adaptables à la volée.

### 2. LangGraph

Si vous voulez que votre agent stratégise et improvise ses étapes, LangGraph est un excellent choix. Basé sur la théorie des graphes, il permet boucles, branches et détours. Parfait pour des workflows non linéaires.

Utile par exemple dans l’hospitalité, où un assistant de voyage évalue plusieurs options ou vérifie le système avant de répondre. Dans l’assurance, un agent qui compare des offres personnalisées doit jongler avec de nombreuses variables : LangGraph s’y prête bien.

