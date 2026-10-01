---
id: collect-261001-ia-llm/ia-llm/gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces-1
title: "gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "OpenAI", "SpaceX", "xAI"]
dates: []
keywords: ["benchmark", "benchmarks", "gemini", "agent", "agents", "astra", "gemini 3.8", "gpt-6", "gpt-live", "grok", "voice"]
source: docs/RAG/collect-261001-ia-llm/gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces.md
source_anchor: ""
source_lines: [1, 54]
sha256: 994e65f44a0469f7856f15375c8ae2eb6beaa83e41ca0831f26da987dea7aaa0
---

# gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces

Cours

La voix est le terrain de jeu du moment. Sur le Speech to Speech Index d’Artificial Analysis, GPT-Live-1 Astra d’OpenAI et Grok Voice Think Fast 2.0 de SpaceXAI se tenaient en tête à 0,2 point d’écart, et OpenAI a lancé GPT-6 Astra début septembre. Le 15 septembre, Google a répondu avec deux nouveaux modèles speech-to-speech : Gemini 3.8 Live et Gemini 3.8 Live Extended Thinking.

Extended Thinking prend la première place de cet index avec un score de 82,6 et mène le benchmark agentique τ-Voice à 68,6 %. Le modèle 3.8 Live de base est le plus économique : dans le test de coût d’Artificial Analysis, il traite une heure d’audio en entrée pour 0,84 $, le plus bas de tous les modèles chartés par Google. Les deux sont disponibles dès aujourd’hui dans l’API Gemini, au même prix que leur prédécesseur, Gemini 3.1 Flash Live.

Dans cet article, nous passons en revue toutes les nouveautés de Gemini 3.8 Live : fonctionnalités, benchmarks, différences entre variantes et coûts d’exécution. Consultez aussi nos guides pour bien démarrer avec l’API Gemini Live et créer un assistant vocal avec GPT-Live-1.

## En bref

- Gemini 3.8 Live et 3.8 Live Extended Thinking sont les nouveaux modèles speech-to-speech de Google pour agents vocaux, remplaçant Gemini 3.1 Flash Live.
- Extended Thinking raisonne et appelle des outils en arrière-plan tout en parlant, et prend désormais la tête des classements speech-to-speech et τ-Voice d’Artificial Analysis.
- Le modèle de base est rapide et peu coûteux mais faible sur les tâches agentiques multi-étapes : choisissez Extended Thinking dès que les outils prennent plus qu’un instant à répondre.
- La tarification est inchangée par rapport à la génération précédente : migrer ne coûte rien au-delà du changement d’identifiant de modèle.
- Si vous êtes sur Gemini 3.1 Flash Live, passez à la nouvelle version. Si vous utilisez la pile temps réel d’OpenAI, l’écart de prix à lui seul justifie un essai d’une après-midi.

## Qu’est-ce que Gemini 3.8 Live ?

Gemini 3.8 Live est le modèle speech-to-speech natif de Google pour agents vocaux en temps réel, sorti le 15 septembre 2026 aux côtés d’un jumeau plus performant en raisonnement, Gemini 3.8 Live Extended Thinking. Les deux fonctionnent via l’API Gemini Live sur WebSocket, acceptent audio, vidéo, images et texte en entrée, et renvoient de l’audio. Google positionne 3.8 Live comme le choix par défaut pour le dialogue à faible latence et Extended Thinking pour les tâches complexes et multi-étapes.

Le saut générationnel par rapport à Gemini 3.1 Flash Live concerne la gestion du « travail qui n’est pas de la parole ». Les appels d’outils et d’API s’exécutent désormais en arrière-plan par défaut pendant que le modèle poursuit la conversation, et Extended Thinking ajoute par-dessus un raisonnement de fond configurable. Google présente ce duo comme une rupture par rapport à ses précédents modèles live et comme un remplacement des pipelines en cascade qui enchaînent reconnaissance vocale, LLM et synthèse vocale.

## Fonctionnalités clés de Gemini 3.8 Live

Google met en avant cinq capacités pour les nouveaux modèles, et les deux qui comptent le plus pour quiconque construit un agent vocal sont celles qui modifient la boucle conversationnelle : les appels d’outils asynchrones et le raisonnement en arrière-plan.

### Continuer à parler pendant l’exécution des outils

Les deux modèles exécutent des appels de fonctions et des requêtes d’API en arrière-plan tout en diffusant l’audio à l’utilisateur. Sur Gemini 3.8 Live, l’exécution asynchrone est désormais le mode par défaut pour l’appel de fonctions. Vous pouvez toujours forcer l’ancien comportement synchrone avec `behavior: BLOCKING` sur une déclaration d’outil, et le modèle prend en charge la planification de fonctions avec les options `SILENT`, `WHEN_IDLE` et `INTERRUPTED` qui déterminent quand un résultat est énoncé.

Extended Thinking va plus loin : il exige des outils non bloquants et renvoie une erreur si vous en déclarez un bloquant. Concrètement, un appelant qui demande une modification de réservation entend d’abord un accusé de réception, puis une mise à jour d’avancement, puis le résultat, au lieu du silence que produit un pipeline en cascade pendant l’attente d’une API. Sur Gemini 3.1 Flash Live, l’appel de fonctions n’était qu’en séquentiel, et le modèle ne commençait pas à répondre avant que vous ne renvoyiez le résultat de l’outil.

### Raisonner en arrière-plan sans se taire

Gemini 3.8 Live Extended Thinking raisonne et parle en même temps. La documentation de Google décrit le modèle utilisant des signaux verbaux précoces comme « Laissez-moi vérifier » pour accuser réception, puis narrant la progression à mesure que le travail de fond multi-étapes s’exécute. Vous contrôlez la profondeur avec `thinking_level` à `low`, `medium` ou `high` ; le niveau `MINIMAL` présent sur 3.1 Flash Live disparaît.

Cela change le protocole, et c’est selon moi le point de migration le plus important de la version. Parce que le modèle peut parler plusieurs fois pendant une même requête, `turnComplete: true` ne signifie plus qu’il est au repos.

Votre client doit surveiller un nouveau champ `interaction_status`, qui indique `IN_PROGRESS` tant que le raisonnement ou les outils tournent et `IDLE` quand la tâche est entièrement terminée. Sans cela, votre interface repassera en mode « écoute » alors que le modèle est encore en plein travail.

### Voir ce que voit l’utilisateur

Gemini 3.8 Live ancre le dialogue dans l’entrée visuelle en direct, en traitant des images vidéo quasi en temps réel pour qu’un agent réagisse à ce que l’utilisateur regarde autant qu’à ce qu’il dit. Les démonstrations de Google incluent une partie d’échecs en direct et un onboarding employé où le modèle répond aux questions sur ce qui s’affiche à l’écran.

La vidéo n’est pas gratuite pour autant. La couverture d’un tour envoie désormais par défaut l’activité audio plus toutes les images vidéo au modèle ; si votre application n’a pas besoin de vision, envoyez des frames seulement lorsqu’elles sont utiles, pour le budget de contexte comme pour le coût. Les sessions audio+vidéo sont aussi limitées à 2 minutes avant d’avoir besoin de gestion de session pour les prolonger, contre 15 minutes pour les sessions audio seules.

### Ne plus se tromper sur les codes et numéros de dossier

Google appelle cela la « précision alphanumérique » : analyser correctement des chaînes comme des codes de confirmation, numéros de dossier ou identifiants techniques dictés à voix haute. Qiconque a déjà construit un agent vocal pour le support ou la logistique sait que c’est là que les piles en cascade trébuchent : une erreur de reconnaissance vocale au début est irrécupérable en aval. Un modèle vocal natif qui entend l’énoncé complet en contexte a ici un avantage structurel.

### Changer de langue en plein milieu d’une phrase

Gemini 3.8 Live détecte et bascule entre 97 langues prises en charge au cours d’une conversation, avec ce que Google appelle une cohérence d’accent. Les deux modèles prennent aussi en charge ce que Google nomme des mises à jour de contenu incrémentales : vous pouvez injecter des données structurées dans la session à tout moment, avec un rôle explicite `user` ou `model`, et le modèle les fusionne avec l’audio en direct. C’est ainsi que vous injectez une fiche client ou un statut de commande dans un appel en cours sans casser le flux.

