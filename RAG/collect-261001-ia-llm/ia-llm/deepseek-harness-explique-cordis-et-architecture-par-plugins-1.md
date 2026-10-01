---
id: collect-261001-ia-llm/ia-llm/deepseek-harness-explique-cordis-et-architecture-par-plugins-1
title: "deepseek-harness-explique-cordis-et-architecture-par-plugins"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "OpenAI"]
dates: []
keywords: ["deepseek", "agent", "agents", "claude", "open source"]
source: docs/RAG/collect-261001-ia-llm/deepseek-harness-explique-cordis-et-architecture-par-plugins.md
source_anchor: ""
source_lines: [1, 87]
sha256: e4f2e2bb5796e9130f59b9785411402e95b0f6a15e4ae931f509437b640a4f74
---

# deepseek-harness-explique-cordis-et-architecture-par-plugins

Cursus

DeepSeek Harness est conçu pour exécuter une tâche, pas seulement répondre à une question. Il s’agit d’un environnement d’exécution d’agent open source qui connecte un modèle à votre dépôt, votre terminal, vos outils et l’historique de session. Demandez-lui de corriger un bug : il peut inspecter des fichiers, modifier du code, lancer des tests et réagir quand une commande échoue. Un simple appel modèle ne peut pas faire tout cela seul.

La partie la plus originale se trouve sous ce flux de travail. DeepSeek Harness expose l’adaptateur de modèle, les outils, les sessions, le bac à sable et même la boucle de l’agent sous forme de plugins coordonnés par Cordis. Le modèle n’est qu’un des composants de l’agent, pas le produit en soi.

Ce n’est pas un logiciel abouti. Harness est encore en préversion développeur ; ses API peuvent changer entre deux versions et son propre avis de sécurité indique qu’aucun audit n’a été réalisé. J’aborderai ces limites en même temps que l’architecture et les différences avec Claude Code, Codex et OpenCode.

## En bref

- **Ce que c’est :** DeepSeek Harness est un environnement d’exécution d’agent open source, pas un modèle. Il fournit au modèle des outils, des sessions, un bac à sable et une boucle d’agent.
- **Conception centrale :** Cordis expose l’adaptateur de modèle, les outils, le magasin de sessions, le bac à sable et la boucle de l’agent comme des plugins interchangeables.
- **Sessions :** Un journal d’événements en ajout-only permet la reprise, le fork, la recherche, la relecture et la vue Trajectory.
- **Modes :** Standard, PTC, Minimal et Creator modifient les outils accessibles à l’agent et la manière d’y accéder.
- **Principale différence :** DeepSeek Harness permet aux développeurs de remplacer des composants bas niveau que Claude Code, Codex et OpenCode laissent fixes.
- **Principale limite :** cela reste une préversion développeur sans audit de sécurité, et ses API peuvent évoluer entre versions.

## Introduction aux agents d'intelligence artificielle

## Qu’est-ce que DeepSeek Harness ?

DeepSeek Harness, abrégé en `dsh`, est un agent harness open source de DeepSeek AI sous licence MIT. Il s’intercale entre un modèle de langage et le monde extérieur, en fournissant des outils, des sessions, un bac à sable et la boucle qui fait avancer la tâche.

Chez DeepSeek, l’équation est **« Agent = Modèle + Harness ».** Le modèle gère le raisonnement et la génération. Le harness correspond à tout ce qui permet à ce raisonnement d’agir sur un système de fichiers réel et de poursuivre sans que vous ayez à réexpliquer la tâche à chaque étape.

Il s’appuie sur Cordis, un framework de plugins antérieur à DeepSeek Harness. Cordis permet de remplacer ces briques de manière indépendante via la configuration. Je reviendrai plus loin sur le coût de ce choix.

Avec ce cadre en tête, voici deux idées reçues fréquentes.

### DeepSeek Harness n’est pas un modèle d’IA

Comme indiqué plus haut, le modèle et l’environnement d’exécution sont deux couches distinctes. Cette séparation permet de changer de fournisseur sans toucher aux outils ni au paramétrage des sessions. Le même runtime peut utiliser DeepSeek, Anthropic, OpenAI ou un endpoint compatible OpenAI.

### DeepSeek Harness va au-delà d’un assistant de code

Le mode Standard peut donner l’impression d’un assistant de développement, mais ce n’est qu’une configuration. Comme nous le verrons, les modes Minimal et Creator changent les outils accessibles à l’agent. Construire une nouvelle configuration demande toujours du travail d’ingénierie ; les développeurs ont accès aux pièces.

## Comment Cordis organise les plugins de DeepSeek Harness

Comme mentionné, Cordis est le framework de plugins sous DeepSeek Harness. Il permet à chaque composant de demander un service sans lier le code à un fournisseur unique.

Cordis vient de l’écosystème du chatbot Koishi et a été construit par un développeur connu sous le nom de Shigma ; DeepSeek le distribue et l’étend. Les auteurs décrivent la conception dans leur article A Programming Paradigm for Spatiotemporal Composability.

Ces bases mènent au slogan principal du projet et à deux notions de Cordis. Les noms paraissent académiques, mais le fonctionnement reste simple.

### « Tout est un plugin »

La documentation d’architecture de DeepSeek explique que vous étendez `dsh` en montant un plugin à côté des autres. Les adaptateurs de modèle, les outils, les sessions, les bacs à sable, le stockage, l’ordonnancement, la boucle d’agent et l’interface sont tous des plugins.

Pris au pied de la lettre, le slogan va trop loin. Cordis reste en dessous des plugins. Il les charge et les décharge, vérifie leurs dépendances et pilote les événements qu’ils utilisent pour communiquer. Cordis est indispensable, ce n’est pas une pièce optionnelle de plus.

### La composabilité spatiale gère les dépendances entre plugins

Un plugin déclare les services dont il a besoin sans exiger une séquence de démarrage écrite à la main. Il s’active quand ces services existent et se désactive si un service requis disparaît. Ses dépendances dictent quand il peut s’exécuter.

DeepSeek appelle cela composabilité spatiale. Les dépendances indiquent à Cordis où se place un composant, évitant aux développeurs d’ordonner le démarrage à la main.

### La composabilité temporelle annule les effets des plugins

Cordis suit aussi les enregistrements tels que les écouteurs d’événements, sections de prompt et schémas d’outils. Retirer un plugin supprime ces effets au lieu de laisser des écouteurs orphelins. Cela n’annule pas une action externe comme une commande shell ; la réversibilité ne s’applique qu’aux effets suivis par Cordis.

## Architecture de DeepSeek Harness : comment s’assemble le runtime

Une instance en cours d’exécution est un arbre de plugins construit à partir de paramètres chargés dans un ordre défini. Ces paramètres déterminent les briques actives.

Cordis connecte tous les plugins remplaçables du runtime. Image de l’auteur.

### Les services Cordis permettent aux plugins de se découvrir

Cordis fournit un annuaire partagé de services. Les plugins utilisent des clés stables comme `ctx.tools`, `ctx.llm` et `ctx.sessions` au lieu d’importer le code d’un fournisseur. Un outil qui appelle `ctx.llm` n’a pas besoin de savoir quel adaptateur de modèle se trouve derrière.

### Presets d’agent et profils d’exécution pilotent des couches différentes

Si tout est remplaçable, il faut tout de même décider quoi monter pour une exécution donnée, et DeepSeek Harness répond à deux niveaux faciles à confondre.

Version courte : un profil contrôle la façon dont le programme démarre, tandis qu’un preset contrôle ce que l’agent peut faire. Si vous n’utilisez que l’application web, vous pouvez ignorer les deux sous-sections suivantes.

#### Profils d’exécution

Un profil d’exécution (`web`, `headless`, `sdk`, `sdk-minimal` et `acp` sont fournis comme modèles) décide comment l’application se lance et quels bundles de plugins Cordis sont empilés au démarrage. La plupart des lecteurs ne toucheront à ce niveau qu’en lançant `dsh web` ou une commande similaire.

#### Presets d’agent

Un preset d’agent (Standard, PTC, Minimal ou Creator) décide de ce qu’une session active peut utiliser. Un fichier de patch peut changer le preset sans modifier le code source d’Harness.

### La boucle d’agent coordonne tours, étapes et appels d’outils

