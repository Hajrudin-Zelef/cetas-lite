---
id: collect-261001-ia-llm/ia-llm/frameworks-dagents-ia-des-outils-pour-des-systemes-plus-intelligents-2
title: "frameworks-dagents-ia-des-outils-pour-des-systemes-plus-intelligents"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Google", "Lambda", "Microsoft", "OpenAI"]
dates: []
keywords: ["agent", "agents", "agentic", "aws", "gemini", "incident", "open source", "valuation"]
source: docs/RAG/collect-261001-ia-llm/frameworks-dagents-ia-des-outils-pour-des-systemes-plus-intelligents.md
source_anchor: ""
source_lines: [101, 209]
sha256: a10a2fe283cffd20eff1dc90033e47419ef6005768e270034ac40e6c2a817094
---

# frameworks-dagents-ia-des-outils-pour-des-systemes-plus-intelligents

Intégré à LangChain, il donne accès à tout l’écosystème d’outils. Inconvénient : une courbe d’apprentissage marquée. Mais une fois maîtrisé, vous débloquez une logique très sophistiquée.

Explorez nos cours Designing Agentic Systems with LangChain et le nouveau Multi-Agent Systems with LangGraph pour apprendre à construire des systèmes multi-agents avec des patterns agentiques.

## Systèmes multi-agents avec LangGraph

### 3. AutoGen

AutoGen est la réponse de Microsoft à l’orchestration d’agents, et il ne vient pas les mains vides. Son architecture en couches comprend des fonctions Core, AgentChat pour la messagerie inter-agents, et des Extensions avancées.

Point fort : AutoGen Studio — un environnement low-code pour concevoir visuellement des agents. Particulièrement utile si vous travaillez avec Azure ou Teams.

Il gère aussi la messagerie asynchrone entre agents, pratique pour des workflows dynamiques. Seul bémol : il peut paraître trop « ingénieré » pour des projets simples.

### 4. Agno (par Phidata)

Agno suit la philosophie « less is more » : un framework minimaliste, syntaxe Python épurée, déploiement cloud intégré, idéal pour prototyper vite.

Envie d’un bot d’analyse de sentiment sur X/Twitter via AWS Lambda ? C’est un bon point de départ. Il s’intègre facilement aux LLMs et à des APIs comme DuckDuckGo ou Yahoo Finance : parfait pour des dashboards, agents ou outils internes en prototype.

Limite : communauté plus restreinte, donc moins de plug-ins et de guides en cas de blocage.

### 5. Atomic Agents

Atomic Agents s’adresse à ceux qui veulent un contrôle total, sans pilotage automatique. L’architecture, claire et modulaire, demande plus de mise en place au départ, mais offre ensuite une flexibilité maximale.

Pas de couche d’orchestration « magique » en coulisses : ce que vous construisez est ce que vous obtenez. Idéal pour des équipes focalisées sur la maintenabilité et la performance long terme (outils d’entreprise, R&D).

En revanche, pour un prototype à livrer vite, ce n’est pas le plus adapté.

### 6. OpenAI Agents SDK

OpenAI Agents SDK est l’évolution prêt-à-la-production de Swarm. Framework léger pour créer des workflows multi-agents, il est agnostique côté fournisseur et prend en charge les APIs OpenAI Responses et Chat Completions ainsi que 100+ autres LLMs.

Son idée clé : les « handoffs », où un agent termine une tâche et la transmet au suivant, comme un relais. Idéal pour des applis ou démos qui exigent confidentialité et rapidité, sans lourde infrastructure.

### 7. LlamaIndex

Lancé comme GPT Index, LlamaIndex aidait d’abord les LLMs à exploiter des données structurées — et s’est désormais aventuré dans le monde des agents IA.

Excellent pour extraire des insights à partir de documents, tableaux et bases, il gère désormais les transitions entre requêtes et des workflows simples. Autres cas d’usage : applications multimodales et agents autonomes capables de rechercher et d’agir.

Moins fort en orchestration que CrewAI ou LangGraph, il excelle dès que l’enjeu est l’interaction poussée avec la donnée plutôt que des comportements d’agents complexes.

### 8. Semantic Kernel

Encore un Microsoft : Semantic Kernel est le framework d’entreprise pour agents, compatible Python, C# et Java, conçu pour intégrer des LLMs dans des systèmes comme des ERP et CRM. Il automatise des processus en combinant prompts et APIs existantes pour exécuter des actions.

Il gère mémoire, planification et « skills » (outils) et passe à l’échelle dans le cloud. Si vous avez besoin d’une intégration stricte à l’existant, c’est probablement le meilleur choix. En solo ou petite équipe, il peut toutefois paraître un peu lourd pour démarrer vite.

### 9. Google ADK (Agent Development Kit)

Pour finir : vous pouvez créer des agents IA avec Google ADK. Si vous ne connaissez pas encore, Google ADK est un framework open source pour construire des agents modulaires et intelligents, avec un fort accent sur l’orchestration et l’usage d’outils. Il prend en charge les LLMs (Gemini ou autres), les agents personnalisés et des agents de workflow.

ADK gère le paradigme « agent-comme-outil » (où des agents peuvent en appeler d’autres) et propose des workflows avec évaluation. C’est un excellent choix pour les développeurs qui veulent bâtir des assistants sérieux, multi-étapes.

## Créer des agents d'intelligence artificielle avec Google ADK

## Comment choisir le bon framework

Avant de vous lancer, clarifiez l’objectif final : que doivent concrètement faire vos agents ? Le bon framework doit s’aligner sur votre stack, votre budget et vos besoins à court comme à long terme.

*Critères pour sélectionner le bon framework. Source : Napkin IA*

### Facilité d’usage

Pour démarrer, Agno ou CrewAI sont d’excellents points d’entrée. Documentation claire, frictions minimales. La syntaxe épurée d’Agno permet de prototyper très vite, tandis que l’organisation par rôles de CrewAI est intuitive, même sans être développeur. Les deux prennent en charge la complexité en coulisses pour vous laisser viser le résultat.

### Complexité du workflow

Besoin de boucles, conditions ou branches ? LangGraph et LangChain sont de bonnes options. LangGraph gère des flux à états complexes : parfait pour des chatbots qui retentent une étape ou escaladent un incident.

Et si vous coordonnez une équipe d’agents aux rôles différents, le modèle de délégation de CrewAI fonctionne très bien.

### Personnalisation

Si vous voulez tout contrôler dans le détail, Atomic Agents est fait pour vous. Aucune abstraction cachée : flexibilité brute pour les équipes de dév. Idéal quand le plug-and-play ne suffit pas (ex. : coordination de drones en situation de crise). LangGraph offre un juste milieu : vous ajoutez votre logique tout en profitant d’outils intégrés pour le courant.

### Besoins d’intégration

Si votre projet dépend d’APIs ou d’outils d’entreprise, et si votre organisation est déjà sur l’écosystème Microsoft, Semantic Kernel s’impose.

CrewAI est aussi solide : connecteurs pratiques pour des APIs comme X. Si votre agent doit agréger de multiples sources, ces frameworks gèrent l’intégration pour vous.

### Environnement de déploiement

Pour des applis locales dans des secteurs sensibles (santé, banque) où la confidentialité est critique, l’Agent SDK d’OpenAI permet de rester on-device/on-prem sans cloud. Si vous visez le cloud, Agno propose le support intégré d’AWS, GCP et des workflows serverless : quasi « déploiement en un clic » en production.

### Performance/évolutivité

Si vous gérez un trafic important ou de gros volumes de données, LangGraph brille par sa capacité à piloter des milliers de workflows simultanés — idéal pour la détection de fraude. AutoGen de Microsoft est conçu pour l’échelle entreprise, avec support des systèmes distribués : parfait pour de grands réseaux et chaînes d’approvisionnement.

### Sécurité et confidentialité

Dans le cloud, Semantic Kernel apporte une sécurité de niveau entreprise. Avec CrewAI, la sécurité dépend de votre déploiement : chiffrement et contrôles d’accès restent à votre charge.

Voici un récapitulatif rapide des critères et des meilleurs choix probables.

| **#** | **Facteur** | **Meilleurs choix** | 
| 1 | Débutant-friendly | CrewAI, Agno | 
| 2 | Workflows complexes | LangGraph, AutoGen | 
| 3 | Personnalisation poussée | Atomic Agents, LangGraph | 
| 4 | Intégration API/outils | Semantic Kernel, CrewAI | 
| 5 | Déploiement cloud | Agno, Semantic Kernel | 
| 6 | Local/axé confidentialité | OpenAI’s Agent SDK | 
| 7 | Applications à l’échelle | AutoGen, LangGraph | 

## Tableau comparatif : en un coup d’œil

