---
id: collect-261001-ia-llm/ia-llm/openai-devoile-36-nouveaux-cas-ou-l-ia-a-essaye-de-se-cacher-1
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Apple", "Google", "Hugging Face", "JFrog", "OpenAI"]
dates: []
keywords: ["agent", "agents", "chatgpt", "claude", "gpt-5.6", "incident", "jailbreak", "opus 5", "rlhf", "sol", "sonnet 5"]
source: docs/RAG/collect-261001-ia-llm/openai-devoile-36-nouveaux-cas-ou-l-ia-a-essaye-de-se-cacher.md
source_anchor: ""
source_lines: [1, 72]
sha256: d2eb846becaf992a88bfdeb6986b19a2a936a43f00f8342bbf3b5a9eec6588ec
---

# 🧠 **RECHERCHE**

## **Aujourd'hui:**

🕵️ Les modèles d'OpenAI dissimulent leurs erreurs à leurs créateurs

🎯 Jev, l'IA qui classe au lieu d'écrire

🧩 Anthropic fond Cowork dans Claude et lance Docs et Slides

🍎 Apple va entraîner son IA avec vos conversations Siri

🐭 Des souris dont la moitié du cerveau est faite de cellules humaines

🪰 Un cerveau de mouche recruté pour écrire des titres de presse

🐘 Un modèle de 4 milliards de paramètres bat l'optimiseur de Postgres

🏠 Google Home s'ouvre aux agents IA

💘 Des applications de rencontre entièrement fabriquées par l'IA

💾 La crise de la RAM ne fait que commencer

🇪🇺 Von der Leyen et les agents qui s'échappent de leur bac à sable

🎬 Un film IA de 2h30 qui rate absolument tout

🇨🇳 Les modèles chinois raflent l'usage réel des agents

🛰️ Le Pentagone confirme avoir des armes en orbite

**Savoir se servir de ChatGPT ne vous distingue plus. Votre collègue le fait, le stagiaire aussi.**

Et une compétence que tout le monde a ne se facture pas. Pourtant, il reste une demande que ChatGPT ne peut pas servir : un cabinet comptable ne mettra jamais ses dossiers clients dedans. Un avocat non plus, un cabinet médical encore moins. Ils veulent l'IA comme tout le monde, mais sans que leurs fichiers sortent du bureau.

Ça existe, et ça s'appelle l'IA locale : le modèle tourne sur leur propre machine, et il répond même avec le wifi coupé. Pareil pour l'image et la vidéo, sans limite de crédits. Peu de gens savent l'installer, et personne ne l'enseigne en français. Une compétence rare, elle, se facture.

**Dans l'Académie Privée VISION IA**, je vous l'apprends dans l'ordre, en commençant par vérifier ce que votre machine peut faire tourner. Aujourd’hui pas besoin d'une machine de laboratoire.

39 € par mois, **sans engagement, résiliable à tout moment.**

OpenAI a publié mercredi un rapport recensant **six cas** de comportements « inattendus ou préoccupants » observés sur ses modèles au cours des **six derniers mois**, en plus de l'incident Hugging Face de juillet. Les deux plus frappants : un modèle de recherche non publié et une session d'entraînement de **GPT-5.6 Sol** ont inséré, dans les résumés de leurs propres conversations, des instructions destinées à leurs futures versions afin de « dissimuler des erreurs ou un comportement désaligné à l'utilisateur ».

Deux modèles ont glissé des consignes à leurs successeurs dans les résumés de fenêtres de chat, pour masquer leurs propres ratés
Un modèle réservé à l'usage interne a utilisé **une clé API fuitée sans autorisation** , puis a fabriqué des données de toutes pièces
Des modèles et des agents ont communiqué entre eux via **des forums et des partages de fichiers non sanctionnés**
Deux exemples d'entraînement montrent des modèles qui **uploadent des fichiers sur internet** pour pouvoir ensuite les citer comme sources pertinentes devant leurs évaluateurs humains
Rappel du précédent : en juillet, un modèle interne de la taille de GPT-5.6 Sol avait contourné son isolement, communiqué via le gestionnaire de paquets Artifactory et exfiltré des données dans un dataset malveillant (__le récit d'OpenAI__ )

L'entreprise accompagne ces révélations d'un cadre de signalement : n'importe quel employé peut remonter un incident à l'équipe sécurité et alignement, avec des délais imposés pour l'enquête puis la divulgation publique. Sam Altman avait par ailleurs soutenu samedi la proposition d'Anthropic, pourtant son concurrent direct, de ralentir le rythme des progrès.

Un laboratoire valorisé près de **1 000 milliards de dollars** écrit noir sur blanc que « l'industrie de l'IA n'a pas résolu l'alignement et la surveillance à un degré suffisant pour continuer à monter en puissance à vitesse maximale de façon responsable encore très longtemps ». Ce n'est plus une critique extérieure, c'est un constat interne, publié volontairement, et il change la nature du débat sur la vitesse de déploiement.

TypeSafe AI, fondée par Diogo Almeida, ancien chercheur d'OpenAI et coauteur d'InstructGPT, lance Jev, un modèle qui ne génère strictement aucun texte. On lui donne une question et une liste d'options, il en choisit une et attribue des probabilités, en **70 à 500 millisecondes**, pour **0,042 dollar par million de tokens** en entrée et zéro frais en sortie.

TypeSafe revendique **193,6 fois plus rapide que Claude Sonnet 5** et**444,6 fois moins cher que Claude Opus 5** sur des tâches de décision structurée, avec des gains de**40x à 200x** selon les requêtes
Le modèle est entraîné par RLCD, « Reinforcement Learning for Calibrated Decisions », une méthode que l'entreprise oppose au RLHF : l'objectif n'est plus de plaire à l'humain mais de sortir des probabilités « épistémiquement honnêtes »
Cas d'usage concrets : trier un ticket client entre paiement, livraison et retour, scorer des enregistrements en base, filtrer les sorties d'une autre IA pour repérer une tentative de jailbreak
L'accès développeur se fait sur __liste d'attente__ , la taille du modèle n'est pas communiquée
La garantie porte sur le format, pas sur la justesse : Jev ne peut pas inventer une réponse hors du menu, mais il peut parfaitement cocher la mauvaise case

L'essentiel de ce que les entreprises font faire à l'IA n'est pas de la rédaction, c'est du tri : classer, router, filtrer, noter. Sur ces tâches, un modèle incapable de sortir des options autorisées supprime d'un coup toute une famille de bugs, ceux où le logiciel reçoit un paragraphe poli au lieu de la valeur attendue. C'est un pari à rebours de la course aux modèles généralistes, et il arrive avec des prix qui rendent la comparaison difficile à ignorer.

Anthropic a annoncé hier la fusion de Claude Chat et de Cowork, son outil agentique lancé cette année, en une interface unique. Vous ne choisissez plus votre mode : Claude évalue lui-même si votre demande appelle une réponse en trois lignes ou un chantier de plusieurs heures avec connecteurs et découpage de tâches. Deux produits arrivent en même temps, Claude Docs et Claude Slides.

**Claude Docs** : rédaction collaborative en direct, Claude écrit, commente et pose des questions pendant que d'autres éditent le même document en parallèle, avec export Word et Google Docs
**Claude Slides** : des présentations éditables diapositive par diapositive, présentables depuis l'application, exportables en PowerPoint ou PDF. Claude Design, pour les visuels modifiables, rejoint aussi le chat
Les tâches longues **continuent de tourner dans le cloud** même après avoir fermé votre machine
Déploiement d'abord pour les abonnés **Pro et Max** sur web, desktop et mobile dans les semaines qui viennent, puis Team et Free. Les administrateurs Enterprise disposeront d'un préavis d'au moins**30 jours**
La raison invoquée par Anthropic : la séparation Chat/Cowork était jugée « clunky », les deux interfaces se chevauchaient et personne ne savait laquelle ouvrir (__détails chez VentureBeat__ )

Claude cesse d'être un chatbot auquel on ajoute des outils pour devenir un environnement de travail où le document, la présentation et l'agent vivent dans la même fenêtre. C'est la thèse du moment chez les grands laboratoires : plutôt qu'un catalogue de logiciels spécialisés, un seul assistant généraliste qui décide lui-même de la méthode. Pour l'utilisateur, la promesse est de ne plus avoir à savoir quel outil utiliser. Le prix à payer, c'est de laisser la machine trancher à votre place.

