---
id: collect-261001-ia-llm/ia-llm/creer-un-agent-ia-avec-mistral-12-etapes-2026-1
title: "Créer le dossier du projet"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Glasswing", "Microsoft", "Mistral", "OpenAI", "xAI"]
dates: []
keywords: ["agent", "agents", "chatgpt", "claude", "gpt-5.6", "grok", "mai", "mcp", "mistral", "model context protocol", "multimodal", "open source"]
source: docs/RAG/collect-261001-ia-llm/creer-un-agent-ia-avec-mistral-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 49]
sha256: 16d6ffb265e5ec42935d115bccb6bc1c9487abe47a3dff698b1cb3a6f8dc98c1
---

# Créer le dossier du projet

Construire un **agent IA** n’est plus réservé aux laboratoires de la Silicon Valley : en juin 2026, Anthropic a étendu son programme d’entreprise Project Glasswing à 150 organisations dans plus de 15 pays avec Claude Mythos, Microsoft dévoilait le 2 juin 2026 à sa conférence Build son agent autonome Scout pour Microsoft 365, et OpenAI a lancé le 9 juillet 2026 ChatGPT Work, un agent transverse capable d’exécuter des tâches planifiées dans plusieurs applications. Cet emballement se lit aussi dans les chiffres des investisseurs : les startups de l’IA agentique ont levé 6,42 milliards de dollars sur l’ensemble de 2025 selon le cabinet Tracxn (avril 2026), et déjà 12 milliards de dollars sur le seul premier trimestre 2026 d’après Callsphere.ai (mars 2026), soit le triple des 4 milliards de dollars levés au T1 2025. Depuis le lancement de l’**API Agents** de Mistral le 27 mai 2025, n’importe quel développeur français peut, lui aussi, assembler en quelques heures un assistant autonome capable de chercher sur le web, d’exécuter du code Python, d’interroger ses propres documents et d’appeler des fonctions métier – le tout sur une plateforme européenne. Dans ce tutoriel pas à pas, nous allons créer un **agent IA** complet et fonctionnel en 12 étapes, à partir d’une page blanche, en moins de 30 minutes de code effectif.

L’objectif n’est pas de produire un énième chatbot « bonjour, comment puis-je vous aider ». Nous allons bâtir un **agent IA** outillé : un assistant de veille technologique qui combine recherche web en temps réel, un connecteur maison (function calling) vers une API externe, une mémoire conversationnelle persistante et un mécanisme de délégation entre agents (handoff). À la fin, vous disposerez d’un projet Python complet, réutilisable, et de la grille de lecture nécessaire pour le passer en production. Toutes les commandes, tous les blocs de code et tous les exemples de sortie sont fournis. Aucune étape n’est sautée.

## Agent IA vs chatbot classique : ce qui change en 2026

Avant d’écrire une ligne de code, il faut clarifier le vocabulaire, car la confusion entre « chatbot » et « **agent IA** » est la première source d’erreurs de conception. Un chatbot classique repose sur l’API de complétion de chat : vous envoyez une liste de messages, le modèle répond du texte, point final. Il ne sait rien faire d’autre que générer des mots. Si vous lui demandez la météo à Lyon, il inventera une réponse plausible mais fausse, car il n’a aucun moyen d’accéder à une donnée fraîche.

Un **agent IA**, lui, est un système qui perçoit un contexte, décide d’une action, exécute un ou plusieurs outils, observe le résultat, puis recommence cette boucle jusqu’à atteindre l’objectif. La différence fondamentale tient en un mot : l’*agentivité*. L’agent ne se contente pas de prédire le prochain mot ; il orchestre des appels d’outils (web, code, fonctions, documents) et conserve un état entre les tours de conversation. C’est précisément ce que formalise l’API Agents de Mistral, présentée par l’entreprise comme « l’épine dorsale des plateformes agentiques de niveau entreprise ».

| Critère | Chatbot (Chat Completion) | Agent IA (Agents API) | 
|---|---|---|
| Capacité principale | Génère du texte | Décide et agit via des outils | 
| Accès à des données fraîches | Non (hallucine) | Oui (recherche web intégrée) | 
| Exécution de code | Non | Oui (sandbox Python) | 
| Mémoire entre les tours | Manuelle (vous gérez l’historique) | Persistante (conversations à état) | 
| Appel de fonctions métier | Possible mais à orchestrer soi-même | Natif (function calling + connecteurs) | 
| Délégation entre agents | Non | Oui (handoffs) | 
| Cas d’usage type | FAQ, support de premier niveau | Veille, automatisation, analyse | 

Cette distinction n’est pas théorique. Dans la pratique, choisir l’API Agents plutôt que la simple complétion vous évite de réimplémenter à la main une machinerie complexe : la boucle de raisonnement, le stockage de l’historique, l’exécution sécurisée du code et la gestion des appels d’outils. Mistral fournit ces briques « déployées et prêtes à l’emploi », pour reprendre les termes de l’annonce officielle. Si vous voulez creuser l’approche concurrente multi-agents avec un framework Python dédié, notre tutoriel CrewAI détaille une logique d’orchestration différente ; côté outillage open source, Vercel a de son côté publié le 27 juin 2026 son propre framework *Eve* pour construire et faire tourner des agents IA, signe que la brique agentique s’industrialise sur tout l’écosystème – une course à laquelle s’est aussi jointe xAI, avec son agent de code Grok Build lancé le 8 juillet 2026 et facturé au million de tokens.

## Les modèles Mistral et l’API Agents : panorama 2026

La Plateforme – le nom de la couche développeur de Mistral – expose en 2026 une famille de modèles spécialisés. Pour construire un **agent IA** efficace, il faut choisir le bon modèle selon la tâche : un modèle puissant et multimodal pour le raisonnement et l’orchestration, un modèle plus léger et économique pour les tâches simples à fort volume. Le tableau ci-dessous récapitule les modèles documentés sur La Plateforme et leur usage recommandé.

| Modèle | Alias API (latest) | Usage recommandé | Type | 
|---|---|---|---|
| Mistral Medium 3.5 | mistral-medium-latest | Orchestration d’agents, raisonnement, code | Frontier multimodal | 
| Mistral Small 4 | mistral-small-latest | Tâches à fort volume, faible coût | Hybride instruct/raisonnement | 
| Magistral | magistral-latest | Raisonnement explicite étape par étape | Reasoning | 
| Devstral 2 | devstral-latest | Agents de code, génie logiciel | Code | 
| Pixtral | pixtral-latest | Compréhension d’images et de documents | Vision | 
| Voxtral | voxtral-latest | Transcription audio, synthèse vocale | Audio | 
| OCR 3 | – | Extraction de texte, Document AI | OCR | 

Pour notre **agent IA**, nous utiliserons `mistral-medium-latest` comme cerveau orchestrateur. Ce modèle, dans sa version Mistral Medium 3.5, affiche selon nos mesures un score de 77,6 % sur SWE-Bench Verified – un détail que nous documentons dans notre analyse de Mistral Medium 3.5. Sur ce même segment des agents de code, la concurrence s’est encore resserrée depuis qu’OpenAI a fait passer sa famille Codex/GPT-5.6 en disponibilité générale le 9 juillet 2026, avec une tarification annoncée au million de tokens. Pour les tâches secondaires (résumés courts, classification), nous basculerons sur `mistral-small-latest`, nettement plus économique. Cette stratégie de routage entre modèles est l’un des leviers de coût les plus puissants quand on industrialise un agent.

### Les connecteurs intégrés de l’API Agents

L’API Agents de Mistral fournit nativement plusieurs outils que votre **agent IA** peut appeler à la demande, sans que vous ayez à les héberger :

- **Recherche web** (`web_search` ) : interroge le web et des sources d’actualité réputées pour des informations fraîches. Une variante premium donne accès à des agences de presse selon la documentation.
- **Exécution de code** (`code_interpreter` ) : exécute du Python dans un bac à sable sécurisé pour le calcul, l’analyse et les graphiques.
- **Génération d’images** (`image_generation` ) : produit des visuels, propulsée par le modèle FLUX1.1 [pro] Ultra de Black Forest Labs.
- **Bibliothèque de documents** (`document_library` ) : un RAG hébergé sur des documents stockés dans Mistral Cloud.
- **Outils MCP** : compatibilité avec le Model Context Protocol pour brancher vos propres serveurs d’outils.
- **Function calling** : déclaration de fonctions maison que l’agent appelle avec des arguments structurés.

