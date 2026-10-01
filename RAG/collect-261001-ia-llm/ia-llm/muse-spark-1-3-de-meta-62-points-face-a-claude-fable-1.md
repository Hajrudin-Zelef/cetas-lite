---
id: collect-261001-ia-llm/ia-llm/muse-spark-1-3-de-meta-62-points-face-a-claude-fable-1
title: "muse-spark-1-3-de-meta-62-points-face-a-claude-fable"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Meta", "OpenAI"]
dates: []
keywords: ["claude", "muse", "agent", "agents", "benchmark", "benchmarks", "fable 5", "gemini", "gpt-5.6", "llama", "multimodal", "muse spark"]
source: docs/RAG/collect-261001-ia-llm/muse-spark-1-3-de-meta-62-points-face-a-claude-fable.md
source_anchor: ""
source_lines: [1, 46]
sha256: c41d873053ec02ce870fd8894c5478fa120eeb71c2410af75b7456088fc40d8f
---

# muse-spark-1-3-de-meta-62-points-face-a-claude-fable

Le 2 septembre 2026, Meta a mis en ligne **Muse Spark 1.3**, la quatrième itération de sa ligne de modèles de raisonnement en cinq mois. Le lancement s’est glissé dans une fenêtre particulièrement chargée : un jour après Claude Fable 5.1 d’Anthropic et le même jour qu’une nouvelle version Gemini Flash de Google, selon Coursiv. Score de 62 points sur l’Intelligence Index d’Artificial Analysis, 75,4 % sur le benchmark de programmation DeepSWE 1.1, fenêtre de contexte proche du million de tokens : Meta affiche, pour la première fois depuis longtemps, des chiffres qui la rapprochent réellement d’Anthropic et d’OpenAI plutôt que de simplement les suivre de loin. Reste à savoir ce que ce rattrapage technique change concrètement pour les entreprises et les développeurs en France et en Europe, où le groupe reste sous le coup d’une procédure judiciaire liée à l’entraînement de ses modèles Llama.

## Muse Spark 1.3 : ce que Meta a réellement publié le 2 septembre 2026

Selon l’annonce officielle publiée sur le blog de recherche de Meta, Introducing Muse Spark 1.3, le modèle est disponible depuis le mercredi 2 septembre 2026 via deux canaux : Muse Code, l’agent de programmation en ligne de commande de Meta, et l’API Meta Model destinée aux développeurs tiers. Le déploiement plus large vers Facebook, Instagram et l’assistant Meta AI doit suivre dans les semaines suivantes, sans date précise communiquée par l’entreprise.

Muse Spark 1.3 est un modèle propriétaire à poids fermés, multimodal (texte, image, vidéo), pensé pour les tâches agentiques et les workflows de programmation longs. Meta met en avant trois axes d’amélioration : une meilleure gestion des tâches réparties sur plusieurs étapes dans une même conversation, une collaboration plus active avec l’utilisateur (poser des questions de clarification, signaler un blocage, demander confirmation avant une action à conséquence), et une réduction des hallucinations grâce à un meilleur calibrage sur ses propres limites. Ces trois points, documentés dans l’annonce officielle, correspondent directement aux reproches les plus courants adressés aux agents IA d’entreprise : actions non vérifiées, réponses trop confiantes, et perte de contexte sur les tâches longues.

Il s’agit de la quatrième mise à jour de la famille Muse Spark en cinq mois, un rythme de publication que TechTimes qualifie de “première version à afficher des scores de benchmark qui la placent à distance réelle des deux modèles phares d’Anthropic”. Contrairement à la ligne Llama, ouverte et téléchargeable, Muse Spark reste fermée : aucun poids n’est publié, et le nombre exact de paramètres n’a pas été communiqué par Meta.

## Les scores qui comptent : Intelligence Index, DeepSWE et contexte long

D’après Artificial Analysis, cabinet de référence pour le benchmarking indépendant des grands modèles de langage, Muse Spark 1.3 en configuration “max” obtient **62 points sur l’Intelligence Index**, un score qui le place juste derrière Claude Fable 5.1 et Claude Opus 5, mais devant l’ensemble des autres modèles évalués à ce jour. Le site titre d’ailleurs son analyse “Meta atteint la frontière”, une formule qui résume assez bien le changement de statut du modèle par rapport aux générations précédentes de Muse Spark.

Sur le benchmark de programmation DeepSWE 1.1, qui évalue la capacité d’un modèle à résoudre des tâches d’ingénierie logicielle réelles, Muse Spark 1.3 obtient **75,4 %**, soit un bond de 16 points par rapport à la version précédente selon les chiffres relayés par TechTimes. Sur le test de raisonnement à contexte long MRCR, le modèle atteint **98,5 %**, avec de bonnes performances jusqu’à des contextes d’un million de tokens. Plusieurs analyses techniques, dont celle d’eesel, rapportent également que Muse Spark 1.3 dépasse GPT-5.6 Sol et Claude Opus 5 sur certains benchmarks de code spécifiques, bien que l’écart reste étroit et dépende fortement du type de tâche testée.

Sur l’efficacité opérationnelle, un critère de plus en plus scruté par les équipes qui déploient des agents IA en production, Meta annonce environ 20 % d’appels d’outils en moins et 25 % de tokens consommés en moins par rapport à Muse Spark 1.2, pour un niveau de tâche comparable. Concrètement, cela signifie des factures d’API plus prévisibles pour les entreprises qui font tourner des agents en continu, un argument que Meta pousse activement dans sa communication aux développeurs.

| Modèle | Éditeur | Date | Score Intelligence Index | 
|---|---|---|---|
| Claude Fable 5.1 | Anthropic | 1er septembre 2026 | 66 points | 
| Claude Opus 5 | Anthropic | 24 juillet 2026 (relevé le 15 août) | 63,1 points | 
| Muse Spark 1.3 (max) | Meta | 2 septembre 2026 | 62 points | 
| Muse Spark 1.3 (xhigh) | Meta | 2 septembre 2026 | 61 points | 
| Quasar 438B | Multiverse Computing | 2 septembre 2026 | 43 points (2 langues) | 

Ce tableau illustre un point souvent oublié dans la course aux annonces : un score d’Intelligence Index n’a de sens que comparé sur une même échelle et à une date donnée, le classement bougeant parfois de plusieurs points en quelques semaines. Meta, qui occupait une position nettement plus modeste avec Muse Spark 1.2, entre ici directement dans le trio de tête mondial, un mouvement que même les équipes internes d’Artificial Analysis n’anticipaient pas à ce rythme.

## Fiche technique complète : fenêtre de contexte, tarifs et gains d’efficacité

Meta a communiqué une fiche technique précise pour Muse Spark 1.3, disponible via l’API Meta Model. La fenêtre de contexte atteint 1 048 576 tokens en entrée, avec une sortie maximale de 943 718 tokens, des chiffres confirmés par plusieurs sources techniques indépendantes dont Flowtivity. Côté tarification, Meta propose deux paliers distincts, une nouveauté par rapport aux versions précédentes de Muse Spark.

| Caractéristique | Valeur | 
|---|---|
| Fenêtre de contexte | 1 048 576 tokens | 
| Sortie maximale | 943 718 tokens | 
| Score DeepSWE 1.1 | 75,4 % (+16 points vs Muse Spark 1.2) | 
| Score MRCR (contexte long) | 98,5 % | 
| Tarif standard (données non réutilisées) | 1,25 $ / 1M tokens entrée, 4,25 $ / 1M tokens sortie | 
| Tarif “contributeur” (données réutilisées pour l’entraînement) | 0,10 $ / 1M tokens entrée, 0,20 $ / 1M tokens sortie | 
| Gain d’efficacité vs Muse Spark 1.2 | -20 % d’appels d’outils, -25 % de tokens consommés | 

Le palier “contributeur”, jusqu’à vingt fois moins cher que le tarif standard, pose une question simple mais structurante pour toute entreprise européenne : accepter que ses prompts et résultats alimentent l’entraînement des futurs modèles Meta, en échange d’un coût d’exploitation divisé par dix à vingt. Pour un cabinet de conseil ou une entreprise manipulant des données clients confidentielles, l’option est difficilement compatible avec le RGPD sans anonymisation préalable poussée. Le tarif standard, identique à celui de Muse Spark 1.2, garantit en revanche que les données ne sont pas réutilisées, un point que Meta a choisi de maintenir stable malgré le saut de performance du modèle.

## Une fenêtre de lancement extraordinairement dense en septembre 2026

